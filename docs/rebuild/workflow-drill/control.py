"""Outside-only gate records and ordinary child-process control. No worker inputs."""
import contextlib
import fcntl
import json
import os
from pathlib import Path
import signal
import time

COMMAND = ['codex', 'exec', '--ephemeral', '--ignore-user-config', '-m',
           'gpt-6-sol', '-c', 'model_reasoning_effort="high"', '-s', 'workspace-write']
PIN = 'c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada'
CONTROLLER = '/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed'
CANDIDATE = '79cf33a7690082f28c06bf41a72719d282bc00fe'
DEADLINE = '2026-11-07T22:51:05+00:00'
HOLD_SECONDS = 600
TERM_SECONDS = 5
DRAIN_SECONDS = 5
PUBLISH_SECONDS = 120


def record(path, value):
    path = Path(path)
    pending = path.with_name(path.name + '.pending')
    with pending.open('w') as f:
        json.dump(value, f, indent=2)
        f.write('\n')
        f.flush()
        os.fsync(f.fileno())
    pending.replace(path)
    fd = os.open(path.parent, os.O_RDONLY)
    try:
        os.fsync(fd)
    finally:
        os.close(fd)


@contextlib.contextmanager
def locked(gate):
    with (gate / 'lock').open('a') as f:
        fcntl.flock(f, fcntl.LOCK_EX)
        yield


def reject(gate, reason):
    # The same lock serializes rejection, release and worker launch. Once abort
    # is committed, no later attempt can be released or launched.
    with locked(gate):
        if not (gate / 'abort').exists():
            record(gate / 'abort', {'reason': reason, 'time': time.time()})


def signal_group(pid, sig):
    try:
        os.killpg(pid, sig)
    except ProcessLookupError:
        pass


def group_exists(pid):
    try:
        os.killpg(pid, 0)
        return True
    except ProcessLookupError:
        return False


def worker_groups(gate):
    # Only live-attempt records owned by this gate, never a process scan. Retired
    # attempts must not cause a later signal to a potentially reused process ID.
    return [json.loads(p.read_text())['pid'] for p in gate.glob('worker-*.json')
            if not (gate / p.name.replace('worker-', 'finished-')).exists()]


def release(gate, attempt, ready_bytes):
    with locked(gate):
        if (gate / 'abort').exists():
            raise RuntimeError('rejected gate cannot release')
        if (gate / f'ready-{attempt}.json').read_bytes() != ready_bytes:
            raise RuntimeError('ready record changed during review')
        record(gate / f'release-{attempt}', {'reviewed': True})
