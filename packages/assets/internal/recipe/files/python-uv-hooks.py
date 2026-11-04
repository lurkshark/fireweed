"""Project unittest and import-linter hooks using Fireweed's live module graph."""
import ast
import json
import os
from pathlib import Path
import re
import subprocess
import sys
import time
import unittest
import xml.etree.ElementTree as ET


def graph():
    result = subprocess.run([os.environ.get('FIREWEED_BIN', 'fireweed'), 'modules', '--json'], capture_output=True, text=True, check=True)
    return json.loads(result.stdout)['data']


def matches(file, patterns):
    from fnmatch import fnmatch
    return any(file.startswith(p[:-3] + '/') if p.endswith('/**') else fnmatch(file, p) for p in (patterns or []))


def package_name(module):
    return module['name'].replace('-', '_')


def boundaries():
    modules = graph()
    names = {package_name(m): m for m in modules}
    violations = []
    for source in modules:
        root = Path(source['dir'])
        if not root.exists():
            continue
        for file in root.rglob('*.py'):
            rel = file.relative_to(root).as_posix()
            eval_file = matches(rel, source.get('zones', {}).get('evals'))
            body = file.read_text()
            try:
                tree = ast.parse(body, filename=str(file))
            except SyntaxError as exc:
                raise RuntimeError(f'{source["name"]}: cannot parse {file}: {exc}') from exc
            imports = []
            for node in ast.walk(tree):
                if isinstance(node, ast.Import):
                    imports.extend(a.name for a in node.names)
                elif isinstance(node, ast.ImportFrom) and node.module:
                    imports.append(node.module)
            # Support an import whose package is supplied as a doctor path substitution.
            imports.extend(Path(x).name for x in re.findall(r"__import__\(['\"]([^'\"]+)['\"]\.split\(['\"]/['\"]\)\[-1\]\)", body))
            for imported in imports:
                package = imported.split('.')[0]
                target = names.get(package)
                if target is None:
                    continue
                suffix = imported[len(package):].strip('.').replace('.', '/')
                target_rel = suffix + '.py' if suffix else '__init__.py'
                declared = source['name'] == target['name'] or target['name'] in source.get('depends_on', [])
                if not declared or (eval_file and not matches(target_rel, target.get('zones', {}).get('interface'))):
                    violations.append(f'{source["name"]}: {file.as_posix()} imports {imported} (module {target["name"]}) outside declared dependencies or retained interface')
    for violation in sorted(violations):
        print(violation)
    if violations:
        raise RuntimeError(f'{len(violations)} boundary violation(s)')


class Reporter(unittest.TestResult):
    def __init__(self):
        super().__init__()
        self.starts = {}

    def emit(self, action, test, detail=''):
        print(json.dumps({'action': action, 'name': test.__class__.__name__ + '.' + test._testMethodName if hasattr(test, '_testMethodName') else str(test),
                          'classname': test.id(), 'time': max(0, time.monotonic() - self.starts.get(test.id(), time.monotonic())),
                          'detail': detail}), flush=True)

    def startTest(self, test):
        super().startTest(test)
        self.starts[test.id()] = time.monotonic()
        self.emit('run', test)

    def addSuccess(self, test):
        super().addSuccess(test)
        self.emit('pass', test)

    def addFailure(self, test, err):
        super().addFailure(test, err)
        self.emit('fail', test, self._exc_info_to_string(err, test))

    def addError(self, test, err):
        super().addError(test, err)
        self.emit('fail', test, self._exc_info_to_string(err, test))

    def addSkip(self, test, reason):
        super().addSkip(test, reason)
        self.emit('skip', test, reason)


def child(file):
    root = Path(file).parent
    loader = unittest.TestLoader()
    suite = loader.discover(str(root), pattern=Path(file).name)
    if loader.errors:
        print(json.dumps({'action': 'fail', 'name': 'CollectionFailure', 'classname': 'collection',
                          'time': 0, 'detail': '\n'.join(loader.errors)}), flush=True)
        return 1
    result = Reporter()
    suite.run(result)
    return 0 if result.wasSuccessful() else 1


def evals(directory, report):
    eval_dir = Path(directory) / 'evals'
    files = sorted(eval_dir.glob('test*.py')) if eval_dir.exists() else []
    cases = []
    failed = not files
    module_paths = [str(Path('packages').resolve())]
    env = dict(os.environ)
    env['PYTHONPATH'] = os.pathsep.join(module_paths + [env.get('PYTHONPATH', '')])
    for file in files:
        file_start = len(cases)
        result = subprocess.run([sys.executable, __file__, 'child', str(file)], capture_output=True, text=True, env=env)
        starts = {}
        finished = set()
        for line in result.stdout.splitlines():
            try:
                event = json.loads(line)
            except json.JSONDecodeError:
                continue
            identity = event['classname']
            if event['action'] == 'run':
                starts[identity] = event
            else:
                finished.add(identity)
                cases.append(event)
        for identity, event in starts.items():
            if identity not in finished:
                cases.append({**event, 'action': 'fail', 'detail': 'test started but did not finish'})
                failed = True
        if result.returncode:
            failed = True
            if not any(c['action'] == 'fail' for c in cases[file_start:]):
                cases.append({'name': 'CollectionFailure', 'classname': 'collection', 'time': 0, 'action': 'fail', 'detail': str(file) + ': ' + result.stderr + result.stdout})
        if result.stderr:
            sys.stderr.write(result.stderr)
    if not files:
        cases.append({'name': 'CollectionFailure', 'classname': 'collection', 'time': 0, 'action': 'fail', 'detail': str(eval_dir) + ': no eval files found'})
    failed = failed or any(case['action'] == 'fail' for case in cases)
    suite = ET.Element('testsuite')
    for case in cases:
        elem = ET.SubElement(suite, 'testcase', name=case['name'], classname=case['classname'], time=str(case['time']))
        if case['action'] == 'fail':
            ET.SubElement(elem, 'failure').text = case['detail']
        elif case['action'] == 'skip':
            ET.SubElement(elem, 'skipped').text = case['detail']
    ET.ElementTree(suite).write(report, encoding='unicode')
    if failed:
        raise RuntimeError('Python eval failed')


if __name__ == '__main__':
    try:
        if sys.argv[1] == 'eval':
            evals(sys.argv[2], sys.argv[3])
        elif sys.argv[1] == 'boundaries':
            boundaries()
        elif sys.argv[1] == 'child':
            sys.exit(child(sys.argv[2]))
        else:
            raise RuntimeError('usage: hooks.py eval <module-dir> <report> | boundaries')
    except Exception as exc:
        print(exc, file=sys.stderr)
        sys.exit(1)
