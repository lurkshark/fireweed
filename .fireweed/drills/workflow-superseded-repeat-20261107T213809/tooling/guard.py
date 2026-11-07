#!/usr/bin/env python3
"""External agent.wrap: hold, durably spool raw streams, forward, stop/drain."""
import json
import os
from pathlib import Path
import selectors
import signal
import subprocess
import sys
import time

from control import (COMMAND, DRAIN_SECONDS, HOLD_SECONDS, TERM_SECONDS,
                     group_exists, locked, record, reject, signal_group)


def run(gate, argv, *, hold=HOLD_SECONDS, term=TERM_SECONDS, drain=DRAIN_SECONDS):
    """Rehearsals call this child supervisor with explicit fake argv, no shim."""
    gate = Path(gate).resolve()
    inp = Path.cwd().resolve()
    if gate == inp or inp in gate.parents:
        raise RuntimeError('gate/logs must be outside the reconstruction workspace')
    gate.mkdir(parents=True, exist_ok=True)
    interrupted = []
    previous = {}
    for sig in (signal.SIGINT, signal.SIGTERM):
        previous[sig] = signal.signal(sig, lambda s, f: interrupted.append(s))
    child = None
    try:
        with locked(gate):
            if (gate / 'abort').exists():
                return 125
            attempt = 1 + len(list(gate.glob('ready-*.json')))
            if attempt > 2:
                reject_reason = 'configured attempt bound exceeded'
            else:
                reject_reason = None
                record(gate / f'ready-{attempt}.json', {
                    'input': str(inp), 'argv': argv,
                    'fireweed_bin': os.environ.get('FIREWEED_BIN')})
        if reject_reason:
            reject(gate, reject_reason)
            return 125
        expires = time.monotonic() + hold
        while True:
            if interrupted or time.monotonic() >= expires:
                reject(gate, 'wrapper interrupted or preflight hold expired')
            with locked(gate):
                if (gate / 'abort').exists():
                    record(gate / f'finished-{attempt}.json', {'exit': 125, 'launched': False})
                    return 125
                if (gate / f'release-{attempt}').exists():
                    logs = gate / f'attempt-{attempt}'
                    logs.mkdir()
                    env = dict(os.environ)
                    env['PATH'] = str(inp / '.fireweed/bin') + os.pathsep + env['PATH']
                    env['GOCACHE'] = '/tmp/fireweed-rebuild-cache'
                    child = subprocess.Popen(argv, env=env, stdout=subprocess.PIPE,
                                             stderr=subprocess.PIPE, start_new_session=True)
                    record(gate / f'worker-{attempt}.json', {'pid': child.pid})
                    break
            time.sleep(.05)
        return relay(gate, logs, attempt, child, interrupted, term, drain)
    except BaseException as exc:
        reject(gate, 'wrapper failure: ' + str(exc))
        raise
    finally:
        if child is not None:
            # Also clean ordinary descendants when the leader has already exited.
            signal_group(child.pid, signal.SIGKILL)
            child.wait(timeout=TERM_SECONDS)
        for sig, handler in previous.items():
            signal.signal(sig, handler)


def relay(gate, logs, attempt, child, interrupted, term, drain):
    sel = selectors.DefaultSelector()
    streams = []
    old_blocking = {}
    stopped = None
    exited = None
    errors = []
    try:
        for name, pipe, dest in [('stdout', child.stdout, 1), ('stderr', child.stderr, 2)]:
            writer = (logs / (name + '.bin')).open('xb', buffering=0)
            reader = (logs / (name + '.bin')).open('rb', buffering=0)
            stream = dict(name=name, pipe=pipe, dest=dest, writer=writer,
                          reader=reader, received=0, forwarded=0, eof=False, broken=False)
            streams.append(stream)
            os.set_blocking(pipe.fileno(), False)
            old_blocking[dest] = os.get_blocking(dest)
            os.set_blocking(dest, False)
            sel.register(pipe, selectors.EVENT_READ, stream)
        directory = os.open(logs, os.O_RDONLY)
        try:
            os.fsync(directory)
        finally:
            os.close(directory)
        while True:
            now = time.monotonic()
            if interrupted:
                reject(gate, 'wrapper received stop signal')
            if child.poll() is not None and exited is None:
                exited = now
            if (gate / 'abort').exists() and stopped is None:
                stopped = now
                signal_group(child.pid, signal.SIGTERM)
            if exited is not None and stopped is None and now - exited >= drain:
                reject(gate, 'worker exited without closing streams or completing forwarding')
                stopped = now
                signal_group(child.pid, signal.SIGTERM)
            if stopped is not None and now - stopped >= term:
                signal_group(child.pid, signal.SIGKILL)
            for key, _ in sel.select(.05):
                s = key.data
                try:
                    data = os.read(s['pipe'].fileno(), 65536)
                except BlockingIOError:
                    continue
                if not data:
                    sel.unregister(s['pipe'])
                    s['pipe'].close()
                    s['eof'] = True
                else:
                    s['writer'].write(data)
                    os.fsync(s['writer'].fileno())
                    s['received'] += len(data)
            # Logs are the spool: a slow native reader cannot block draining or
            # grow an in-memory output queue without bound. Never decode bytes.
            for s in streams:
                if not s['broken'] and s['forwarded'] < s['received']:
                    s['reader'].seek(s['forwarded'])
                    data = s['reader'].read(min(65536, s['received'] - s['forwarded']))
                    try:
                        count = os.write(s['dest'], data)
                        s['forwarded'] += count
                    except BlockingIOError:
                        pass
                    except BrokenPipeError:
                        s['broken'] = True
                        errors.append(s['name'] + ' native pipe closed')
                        reject(gate, errors[-1])
            complete = all(s['eof'] and s['received'] == s['forwarded'] for s in streams)
            if child.poll() is not None and complete and not group_exists(child.pid):
                break
            if stopped is not None and now - stopped >= term + drain:
                errors.append('bounded stop/drain expired')
                break
        code = child.wait(timeout=term)
        ok = complete and not errors
        result = code if stopped is None and ok else 130
        record(gate / f'finished-{attempt}.json', {
            'exit': result, 'worker_exit': code, 'launched': True,
            'stopped': stopped is not None, 'complete_forwarding': complete,
            'streams': {s['name']: {'received': s['received'], 'forwarded': s['forwarded'],
                                     'eof': s['eof']} for s in streams}, 'errors': errors})
        return result
    finally:
        sel.close()
        for s in streams:
            s['pipe'].close()
            os.fsync(s['writer'].fileno())
            s['writer'].close()
            s['reader'].close()
        for fd, blocking in old_blocking.items():
            os.set_blocking(fd, blocking)


def main():
    gate, argv = Path(sys.argv[1]), sys.argv[2:]
    if len(argv) != len(COMMAND) + 1 or argv[:-1] != COMMAND:
        # Diagnostics stay outside native stdout/stderr as well.
        gate.mkdir(parents=True, exist_ok=True)
        reject(gate, 'unexpected worker command; worker not invoked')
        return 125
    try:
        return run(gate, argv)
    except BaseException:
        # run() has durably recorded the failure and cleaned up its child.
        # Never mix wrapper tracebacks into the native worker byte streams.
        return 125


if __name__ == '__main__':
    sys.exit(main())
