#!/usr/bin/env python3
"""Outside inspector. Preparation never launches a drill; launch needs owner approval."""
import argparse
import datetime
import hashlib
import json
import os
from pathlib import Path
import signal
import shutil
import subprocess
import sys
import time

from control import (CANDIDATE, COMMAND, CONTROLLER, DEADLINE, DRAIN_SECONDS,
                     PIN, PUBLISH_SECONDS, TERM_SECONDS, group_exists, record,
                     reject, release, signal_group, worker_groups)

HERE = Path(__file__).resolve().parent


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def git(root, *args):
    return subprocess.check_output(['git', '-C', str(root), *args], text=True, timeout=60).strip()


def prepare(root):
    """Only caller setup and hash expectations; no product build or invocation."""
    source = HERE.parents[2]
    assert sha(Path(CONTROLLER)) == PIN, 'controller pin mismatch'
    assert not git(source, 'for-each-ref', '--format=%(refname)', 'refs/fireweed/claims')
    assert not list((source / 'changes').glob('*.md')), 'queued change exists'
    root.mkdir(parents=True, exist_ok=False)
    tooling = root / 'tooling'
    tooling.mkdir()
    for p in HERE.iterdir():
        if p.suffix in ('.py', '.md'):
            shutil.copy2(p, tooling / p.name)
    caller = root / 'caller'
    subprocess.run(['git', 'clone', '--no-hardlinks', '--no-checkout', str(source), str(caller)],
                   check=True, timeout=60)
    subprocess.run(['git', '-C', str(caller), 'checkout', '--detach', CANDIDATE], check=True, timeout=60)
    config = (caller / 'fireweed.yaml').read_text()
    assert '\n    wrap:' not in config
    prefix = [sys.executable, str(tooling / 'guard.py'), str(root / 'gate')]
    config = config.replace('agent:\n', 'agent:\n    wrap: ' + json.dumps(prefix) + '\n', 1)
    (caller / 'fireweed.yaml').write_text(config)
    (caller / 'SYSTEM.md').write_bytes((HERE / 'worker-system.md').read_bytes())
    # This child commit changes setup only, retaining the candidate's product.
    subprocess.run(['git', '-C', str(caller), 'add', 'SYSTEM.md', 'fireweed.yaml'], check=True, timeout=60)
    subprocess.run(['git', '-C', str(caller), 'commit', '-m',
                    'author: isolate workflow drill inputs and external gate'], check=True, timeout=60)
    pin = caller / '.fireweed/bin/fireweed'
    pin.parent.mkdir(parents=True)
    pin.write_bytes(Path(CONTROLLER).read_bytes())
    pin.chmod(0o755)
    paths = git(caller, 'ls-files').splitlines()
    deleted = [p for p in paths if p.startswith('packages/workflow/internal/')]
    inputs = ['SYSTEM.md']
    inputs += [p for p in paths if p.startswith('packages/workflow/') and
               (p in ['packages/workflow/SPEC.md', 'packages/workflow/api.go'] or
                p.startswith(('packages/workflow/model/', 'packages/workflow/evals/')))]
    for module in ['contracts', 'gitstore', 'assets', 'project', 'checks']:
        inputs += [p for p in paths if p == f'packages/{module}/api.go' or
                   p.startswith(f'packages/{module}/model/')]
    expected = {'base': git(caller, 'rev-parse', 'HEAD'), 'candidate': CANDIDATE,
                'input': {p: sha(caller / p) for p in sorted(inputs)},
                'retained': {p: sha(caller / p) for p in paths if p not in deleted},
                'deleted': deleted}
    record(root / 'expected.json', expected)
    assert git(caller, 'diff', '--name-only', CANDIDATE).splitlines() == ['SYSTEM.md', 'fireweed.yaml']
    assert not git(caller, 'status', '--porcelain')
    print('Prepared isolated caller only. No native drill launched.')


def validate_review(attempt, review, ready, gate):
    assert review['result'] == 'pass' and review['input'] == ready['input']
    assert review['candidate'] == CANDIDATE
    if attempt == 1:
        assert review['actual_retained_inputs_reviewed'] is True
    else:
        for key in ['full_first_commands_scripts_patches_reviewed', 'full_native_feedback_reviewed',
                    'second_prompt_only_original_plus_native_feedback', 'fresh_output_matches_reviewed_trace']:
            assert review[key] is True
        assert (gate / 'finished-1.json').exists()
        assert json.loads((gate / 'finished-1.json').read_text())['complete_forwarding']
        for stream in ('stdout', 'stderr'):
            assert review['first_streams_sha256'][stream] == sha(gate / f'attempt-1/{stream}.bin')


def monitor(root, argv, preflight, *, deadline, term=TERM_SECONDS,
            drain=DRAIN_SECONDS, publish=PUBLISH_SECONDS, command=COMMAND):
    """The fake rehearsal calls this same supervisor, with an explicit fake controller."""
    gate = root / 'gate'
    gate.mkdir(exist_ok=False)
    interrupted = []
    previous = {}
    for sig in (signal.SIGINT, signal.SIGTERM):
        previous[sig] = signal.signal(sig, lambda s, f: interrupted.append(s))
    env = dict(os.environ)
    env['PATH'] = str(root / 'caller/.fireweed/bin') + os.pathsep + env['PATH']
    env['GOCACHE'] = '/tmp/fireweed-rebuild-cache'
    controller = None
    checks = {}
    abort_at = None
    drained_at = None
    emergency_at = None
    controller_stop_at = None
    emergency = False
    try:
        with (root / 'native.stdout').open('xb', buffering=0) as out, \
             (root / 'native.stderr').open('xb', buffering=0) as err:
            controller = subprocess.Popen(argv, cwd=root / 'caller', env=env,
                                          stdout=out, stderr=err, start_new_session=True)
            record(root / 'controller.json', {'pid': controller.pid, 'argv': argv})

            def check_stop():
                if interrupted:
                    reject(gate, 'inspector received stop signal')
                if time.time() >= deadline:
                    reject(gate, 'original whole-run deadline reached')
                return (gate / 'abort').exists()

            while True:
                now = time.monotonic()
                check_stop()
                if not (gate / 'abort').exists() and controller.poll() is None:
                    try:
                        for attempt in (1, 2):
                            ready_path = gate / f'ready-{attempt}.json'
                            if not ready_path.exists() or (gate / f'release-{attempt}').exists():
                                continue
                            before = ready_path.read_bytes()
                            ready = json.loads(before)
                            assert ready['argv'][:-1] == command
                            prompt = ready['argv'][-1]
                            assert prompt.startswith('# Rebuild workflow\n')
                            if attempt == 1:
                                assert 'Previous verification failed' not in prompt
                                assert 'Previous reconstruction failure' not in prompt
                            if attempt not in checks:
                                if attempt == 2:
                                    first = json.loads((gate / 'finished-1.json').read_text())
                                    assert first['complete_forwarding'], 'first trace was not fully forwarded'
                                    inp = Path(ready['input'])
                                    expected = json.loads((root / 'expected.json').read_text())
                                    fresh = {str(p.relative_to(inp)): sha(p) for p in inp.rglob('*')
                                             if p.is_file() and '.git' not in p.relative_to(inp).parts
                                             and str(p.relative_to(inp)) not in expected['input']}
                                    assert all(p == 'DRILL_NOTES.md' or p.startswith('packages/workflow/internal/')
                                               for p in fresh), 'forbidden fresh input path'
                                    record(root / 'attempt-1-output.json', fresh)
                                preflight(attempt, check_stop)
                                checks[attempt] = before
                            assert checks[attempt] == before
                            review_path = gate / f'reviewed-{attempt}.json'
                            if not review_path.exists():
                                continue
                            review = json.loads(review_path.read_text())
                            validate_review(attempt, review, ready, gate)
                            preflight(attempt, check_stop)
                            assert not check_stop(), 'stopped during preflight'
                            release(gate, attempt, before)
                    except Exception as exc:
                        reject(gate, 'inspection rejected: ' + str(exc))
                if (gate / 'abort').exists():
                    if abort_at is None:
                        abort_at = now
                        record(root / 'stop.json', {'worker_stop_requested': True})
                    ready = list(gate.glob('ready-*.json'))
                    drained = all((gate / p.name.replace('ready-', 'finished-')).exists() for p in ready)
                    if drained and drained_at is None:
                        drained_at = now
                    # Give the controller a reporting window after the wrapper
                    # returns failure, without releasing any subsequent attempt.
                    if (drained_at is not None and now - drained_at >= publish or
                        drained_at is None and now - abort_at >= term + drain + 1 or
                        time.time() >= deadline):
                        emergency = True
                if controller.poll() is not None:
                    # An unexpectedly dead native controller cannot receive more
                    # output. Ask any still-live wrapper to drain before cleanup.
                    if any(group_exists(pid) for pid in worker_groups(gate)):
                        reject(gate, 'controller exited while worker group remained')
                    elif not (gate / 'abort').exists() or drained_at is not None:
                        break
                if emergency:
                    if emergency_at is None:
                        emergency_at = now
                        for pid in worker_groups(gate):
                            signal_group(pid, signal.SIGTERM)
                    if now - emergency_at >= term:
                        for pid in worker_groups(gate):
                            signal_group(pid, signal.SIGKILL)
                    # Even the whole-run emergency preserves a bounded drain
                    # window before signaling the controller's group (which
                    # includes wrappers). Hard limits still cover a wedged wrapper.
                    if controller_stop_at is None and (drained_at is not None or
                                                        now - abort_at >= term + drain + 1):
                        controller_stop_at = now
                        signal_group(controller.pid, signal.SIGINT)
                    if controller_stop_at is not None:
                        if now - controller_stop_at >= term:
                            signal_group(controller.pid, signal.SIGKILL)
                        if now - controller_stop_at >= term + drain:
                            break
                os.fsync(out.fileno())
                os.fsync(err.fileno())
                time.sleep(.05)
            controller.wait(timeout=term)
            os.fsync(out.fileno())
            os.fsync(err.fileno())
        record(root / 'terminal.json', {'exit': controller.returncode, 'emergency': emergency,
                                       'rejected': (gate / 'abort').exists()})
        return controller.returncode
    finally:
        # Applies to inspector errors too; never leave detached owned groups.
        for pid in worker_groups(gate):
            signal_group(pid, signal.SIGKILL)
        if controller is not None:
            signal_group(controller.pid, signal.SIGKILL)
            controller.wait(timeout=term)
        for sig, handler in previous.items():
            signal.signal(sig, handler)


def launch(root):
    assert sha(Path(CONTROLLER)) == PIN
    expected = json.loads((root / 'expected.json').read_text())
    assert expected['candidate'] == CANDIDATE
    assert git(root / 'caller', 'rev-parse', 'HEAD') == expected['base']
    assert not git(root / 'caller', 'status', '--porcelain')
    deadline = datetime.datetime.fromisoformat(DEADLINE).timestamp()
    assert time.time() < deadline, 'original deadline already passed'

    def preflight(attempt, stopped):
        with (root / f'preflight-{attempt}.log').open('ab', buffering=0) as log:
            child = subprocess.Popen([sys.executable, str(HERE / 'preflight.py'), str(attempt), str(root)],
                                     stdout=log, stderr=log, start_new_session=True)
            expires = time.monotonic() + 120
            try:
                while child.poll() is None:
                    if stopped() or time.monotonic() >= expires:
                        raise RuntimeError('preflight stopped or timed out')
                    time.sleep(.05)
                assert child.returncode == 0, 'live isolation preflight failed'
            finally:
                signal_group(child.pid, signal.SIGKILL)
                child.wait(timeout=TERM_SECONDS)
                os.fsync(log.fileno())

    return monitor(root, [CONTROLLER, 'drill', 'workflow', '--json'], preflight, deadline=deadline)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('action', choices=['prepare', 'launch', 'reject'])
    parser.add_argument('root', type=Path)
    parser.add_argument('reason', nargs='?', default='outside inspector rejected reconstruction')
    args = parser.parse_args()
    root = args.root.resolve()
    if args.action == 'prepare':
        prepare(root)
        return 0
    if args.action == 'reject':
        reject(root / 'gate', args.reason)
        return 0
    return launch(root)


if __name__ == '__main__':
    sys.exit(main())
