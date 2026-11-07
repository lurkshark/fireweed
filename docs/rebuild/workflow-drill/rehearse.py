#!/usr/bin/env python3
"""Bounded external plumbing rehearsal. Never invokes codex or native Fireweed."""
import hashlib
import json
import os
from pathlib import Path
import signal
import subprocess
import sys
import tempfile
import threading
import time

import guard
import inspector
from control import CANDIDATE, group_exists, record, release, signal_group

HERE = Path(__file__).resolve()
PREFIX_OUT = b'worker-stdout\x00\r\n'
PREFIX_ERR = b'worker-stderr\xff\r\n'
BULK = bytes(range(256)) * 1025


def wait_for(check, seconds=8):
    until = time.monotonic() + seconds
    while time.monotonic() < until:
        if check():
            return
        time.sleep(.02)
    raise AssertionError('bounded fake wait expired')


def worker_command(root, mode):
    return [sys.executable, str(HERE), '_worker', str(root), mode]


def worker(root, mode):
    os.write(1, PREFIX_OUT)
    os.write(2, PREFIX_ERR)
    if mode == 'emergency':
        signal.signal(signal.SIGTERM, signal.SIG_IGN)
        time.sleep(12)
        return 99
    if mode in ('normal', 'retry'):
        wait_for(lambda: (root / 'finish-worker').exists())
        os.write(1, BULK)
        os.write(2, BULK[::-1])
        return 0
    child = subprocess.Popen([sys.executable, str(HERE), '_descendant', str(root), mode])
    record(root / 'descendant.json', {'pid': child.pid})

    def stop(sig, frame):
        signal.signal(signal.SIGTERM, signal.SIG_IGN)
        os.write(1, b'worker-term-tail\x00\n')
        os.write(2, b'worker-term-error\xff\n')
        child.wait(timeout=2)
        raise SystemExit(23)

    signal.signal(signal.SIGTERM, stop)
    # Even a failed rehearsal is bounded independently of the inspector.
    time.sleep(12)
    signal_group(os.getpgrp(), signal.SIGKILL)


def descendant(root):
    def stop(sig, frame):
        signal.signal(signal.SIGTERM, signal.SIG_IGN)
        os.write(1, b'descendant-term-tail\n')
        raise SystemExit(0)
    signal.signal(signal.SIGTERM, stop)
    record(root / 'descendant-ready.json', {'ready': True})
    time.sleep(12)


def fake_controller(root, mode):
    if mode in ('emergency', 'reporting-timeout'):
        signal.signal(signal.SIGINT, signal.SIG_IGN)
    results = []
    for n in range(1, 2 if mode == 'emergency' else 3):
        child = subprocess.Popen([sys.executable, str(HERE), '_wrap', str(root), mode],
                                 cwd=root / 'input', stdout=subprocess.PIPE, stderr=subprocess.PIPE)

        def copy(pipe, target):
            with target.open('wb', buffering=0) as f:
                while True:
                    data = os.read(pipe.fileno(), 65536)
                    if not data:
                        break
                    f.write(data)
                    os.fsync(f.fileno())
            pipe.close()

        threads = [threading.Thread(target=copy, args=(pipe, root / f'forwarded-{n}.{name}'))
                   for name, pipe in [('stdout', child.stdout), ('stderr', child.stderr)]]
        for t in threads:
            t.start()
        code = child.wait(timeout=12)
        for t in threads:
            t.join(timeout=2)
            assert not t.is_alive()
        results.append(code)
        if code == 0 and (mode != 'retry' or n == 2):
            break
    record(root / 'fake-report.json', {'wrapper_exits': results,
                                     'fixture_only': True})
    if mode in ('emergency', 'reporting-timeout'):
        time.sleep(12)
    return 0 if mode in ('normal', 'retry') else 1


def fake_inspector(root, mode):
    def preflight(n, stopped):
        assert not stopped()
        p = root / f'fake-preflight-{n}.json'
        calls = json.loads(p.read_text())['calls'] if p.exists() else 0
        record(p, {'calls': calls + 1, 'fixture_only': True})
    return inspector.monitor(root, [sys.executable, str(HERE), '_controller', str(root), mode],
                             preflight, deadline=time.time() + (2 if mode == 'emergency' else 10),
                             term=.3, drain=.5, publish=.6, command=worker_command(root, mode))


def case(base, mode):
    root = base / mode
    root.mkdir()
    (root / 'caller').mkdir()
    (root / 'input').mkdir()
    record(root / 'expected.json', {'input': {}})
    start = time.monotonic()
    child = subprocess.Popen([sys.executable, str(HERE), '_inspect', str(root), mode],
                             stdout=(root / 'fixture.stdout').open('wb'),
                             stderr=(root / 'fixture.stderr').open('wb'), start_new_session=True)
    try:
        gate = root / 'gate'
        ready_path = gate / 'ready-1.json'
        wait_for(lambda: ready_path.exists() and (root / 'fake-preflight-1.json').exists())
        assert not (gate / 'worker-1.json').exists(), 'worker ran before review'
        ready = json.loads(ready_path.read_text())
        review = {'result': 'reject' if mode == 'rejection' else 'pass',
                  'input': ready['input'], 'candidate': CANDIDATE,
                  'actual_retained_inputs_reviewed': True}
        record(gate / 'reviewed-1.json', review)
        live = False
        if mode != 'rejection':
            out = gate / 'attempt-1/stdout.bin'
            err = gate / 'attempt-1/stderr.bin'
            wait_for(lambda: out.exists() and err.exists() and out.read_bytes() == PREFIX_OUT
                     and err.read_bytes() == PREFIX_ERR)
            wait_for(lambda: (root / 'forwarded-1.stdout').read_bytes() == PREFIX_OUT and
                     (root / 'forwarded-1.stderr').read_bytes() == PREFIX_ERR)
            live = group_exists(json.loads((gate / 'worker-1.json').read_text())['pid'])
            assert live, 'durability/forwarding not observed before worker completion'
            assert json.loads((root / 'fake-preflight-1.json').read_text())['calls'] == 2
            if mode in ('normal', 'retry'):
                (root / 'finish-worker').touch()
                if mode == 'retry':
                    wait_for(lambda: (gate / 'ready-2.json').exists() and
                             (root / 'fake-preflight-2.json').exists())
                    assert not (gate / 'worker-2.json').exists()
                    second = json.loads((gate / 'ready-2.json').read_text())
                    record(gate / 'reviewed-2.json', {
                        'result': 'pass', 'input': second['input'], 'candidate': CANDIDATE,
                        'full_first_commands_scripts_patches_reviewed': True,
                        'full_native_feedback_reviewed': True,
                        'second_prompt_only_original_plus_native_feedback': True,
                        'fresh_output_matches_reviewed_trace': True,
                        'first_streams_sha256': {s: inspector.sha(gate / f'attempt-1/{s}.bin')
                                                 for s in ('stdout', 'stderr')}})
            else:
                if mode != 'emergency':
                    wait_for(lambda: (root / 'descendant-ready.json').exists())
                if mode in ('interruption', 'reporting-timeout'):
                    child.send_signal(signal.SIGINT)
        code = child.wait(timeout=8)
        expected_exit = 0 if mode in ('normal', 'retry') else 247 if mode in ('emergency', 'reporting-timeout') else 1
        assert code == expected_exit, (mode, code, (root / 'fixture.stderr').read_text())
        terminal = json.loads((root / 'terminal.json').read_text())
        assert terminal['emergency'] == (mode in ('emergency', 'reporting-timeout'))
        if mode == 'rejection':
            assert not list(gate.glob('worker-*.json'))
            assert not list(gate.glob('release-*'))
            # A stale/later approval cannot reopen an aborted gate.
            record(gate / 'reviewed-1.json', dict(review, result='pass'))
            try:
                release(gate, 1, ready_path.read_bytes())
            except RuntimeError:
                pass
            else:
                raise AssertionError('release succeeded after rejection')
        else:
            finished = json.loads((gate / 'finished-1.json').read_text())
            assert finished['complete_forwarding'], finished
            for stream in ('stdout', 'stderr'):
                raw = (gate / f'attempt-1/{stream}.bin').read_bytes()
                assert raw == (root / f'forwarded-1.{stream}').read_bytes()
            if mode in ('normal', 'retry'):
                assert (gate / 'attempt-1/stdout.bin').read_bytes() == PREFIX_OUT + BULK
                assert (gate / 'attempt-1/stderr.bin').read_bytes() == PREFIX_ERR + BULK[::-1]
                if mode == 'retry':
                    assert json.loads((root / 'fake-preflight-2.json').read_text())['calls'] == 2
                    for s in ('stdout', 'stderr'):
                        assert (gate / f'attempt-2/{s}.bin').read_bytes() == (root / f'forwarded-2.{s}').read_bytes()
            elif mode != 'emergency':
                assert b'worker-term-tail' in (gate / 'attempt-1/stdout.bin').read_bytes()
                assert b'descendant-term-tail' in (gate / 'attempt-1/stdout.bin').read_bytes()
                assert b'worker-term-error' in (gate / 'attempt-1/stderr.bin').read_bytes()
                assert not (gate / 'release-2').exists()
            else:
                assert finished['worker_exit'] == -signal.SIGKILL
                assert not (gate / 'release-2').exists()
        pids = [json.loads(p.read_text())['pid'] for p in gate.glob('worker-*.json')]
        pids.append(json.loads((root / 'controller.json').read_text())['pid'])
        if (root / 'descendant.json').exists():
            pids.append(json.loads((root / 'descendant.json').read_text())['pid'])
        for pid in pids:
            wait_for(lambda: not group_exists(pid), seconds=2)
            # Descendants share the worker's group, so check their individual PID too.
            try:
                os.kill(pid, 0)
            except ProcessLookupError:
                pass
            else:
                raise AssertionError(f'owned child {pid} survived')
        report = json.loads((root / 'fake-report.json').read_text())
        assert report['wrapper_exits'] == ([0] if mode == 'normal' else [0, 0] if mode == 'retry' else
                                           [125, 125] if mode == 'rejection' else
                                           [130] if mode == 'emergency' else [130, 125])
        (root / 'input').rmdir()
        if mode != 'rejection':
            assert (gate / 'attempt-1/stdout.bin').is_file()
            assert (gate / 'attempt-1/stderr.bin').is_file()
        return {'case': mode, 'pass': True, 'live_logs_and_forwarding': live,
                'logs_survive_input_removal': True,
                'all_owned_processes_exited': True, 'terminal': terminal,
                'fake_controller_report': report, 'seconds': round(time.monotonic() - start, 3),
                'streams': {name: {'bytes': len((gate / f'attempt-1/{name}.bin').read_bytes()),
                                  'sha256': hashlib.sha256((gate / f'attempt-1/{name}.bin').read_bytes()).hexdigest()}
                            for name in ('stdout', 'stderr')} if mode != 'rejection' else {}}
    finally:
        signal_group(child.pid, signal.SIGKILL)
        child.wait(timeout=2)
        if (root / 'gate').exists():
            for p in (root / 'gate').glob('worker-*.json'):
                signal_group(json.loads(p.read_text())['pid'], signal.SIGKILL)


def main():
    if len(sys.argv) > 1 and sys.argv[1].startswith('_'):
        action, root, mode = sys.argv[1], Path(sys.argv[2]), sys.argv[3]
        if action == '_worker':
            return worker(root, mode)
        if action == '_descendant':
            return descendant(root)
        if action == '_wrap':
            prompt = '# Rebuild workflow\nfixture'
            if (root / 'gate/ready-1.json').exists():
                prompt += '\n\n# Previous verification failed\nfixture'
            return guard.run(root / 'gate', worker_command(root, mode) + [prompt],
                             hold=4, term=.3, drain=.5)
        if action == '_controller':
            return fake_controller(root, mode)
        if action == '_inspect':
            return fake_inspector(root, mode)
        raise AssertionError(action)
    base = Path(tempfile.mkdtemp(prefix='fireweed-wrapper-rehearsal-', dir='/private/tmp'))
    results = [case(base, mode) for mode in ['normal', 'retry', 'interruption', 'rejection',
                                          'reporting-timeout', 'emergency']]
    invalid = base / 'invalid-command'
    run = subprocess.run([sys.executable, str(HERE.parent / 'guard.py'), str(invalid),
                          sys.executable, '-c', 'raise SystemExit("must not execute")'],
                         capture_output=True, timeout=3)
    assert run.returncode == 125 and run.stdout == b'' and run.stderr == b''
    assert (invalid / 'abort').exists() and not list(invalid.glob('worker-*.json'))
    record(base / 'results.json', {'native_drills': 0, 'cases': results, 'command_rejection': 'pass'})
    print(base)
    print(json.dumps(results, indent=2))


if __name__ == '__main__':
    sys.exit(main())
