from pathlib import Path
import datetime,hashlib,json,os,signal,subprocess,time
ROOT=Path('/Users/joel/Workspace/fireweed')
OUT=Path('/private/tmp/fireweed-chk034-3rjojxry/validation')
OLDPIN=ROOT/'.fireweed/bin/fireweed'
DEADLINE=datetime.datetime(2026,11,7,22,51,5,tzinfo=datetime.timezone.utc).timestamp()
BASE='be35552b706aa3000ff8f11306d129ff1d701b85'
SUBJECT='79cf33a7690082f28c06bf41a72719d282bc00fe'
def git(root,*args):return subprocess.check_output(['git','-C',str(root),*args],text=True,timeout=60).strip()
def sha(p):return hashlib.sha256(p.read_bytes()).hexdigest()
assert not git(ROOT,'status','--porcelain')
assert not git(ROOT,'for-each-ref','--format=%(refname)','refs/fireweed/claims')
archive=ROOT/'changes/archive/0180-retain-failed-hook-output.md'
assert archive.is_file() and '\nphase: applied\n' in archive.read_text()
assert git(ROOT,'rev-parse',SUBJECT)==SUBJECT
assert sha(OLDPIN)=='23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f'
SOURCE=git(ROOT,'rev-parse','HEAD')
assert OUT.is_dir()
CLONE=OUT/'source';BINARY=OUT/'fireweed'
env=dict(os.environ);env.update(GOCACHE='/tmp/fireweed-rebuild-cache',GIT_NO_LAZY_FETCH='1',GOPROXY='off',GOSUMDB='off')
results=json.loads((OUT/'results.json').read_text())
def run(name,args,cwd,timeout=1500,expected=0):
 start=datetime.datetime.now(datetime.timezone.utc).isoformat();limit=min(timeout,DEADLINE-time.time());assert limit>0,'original deadline reached'
 print('START',name,flush=True)
 with (OUT/(name+'.stdout')).open('wb') as out,(OUT/(name+'.stderr')).open('wb') as err:
  child=subprocess.Popen([str(a) for a in args],cwd=cwd,env=env,stdout=out,stderr=err,start_new_session=True)
  (OUT/'active.json').write_text(json.dumps({'stage':name,'pid':child.pid,'argv':[str(a) for a in args]}))
  try:code=child.wait(timeout=limit)
  except subprocess.TimeoutExpired:
   os.killpg(child.pid,signal.SIGTERM)
   try:child.wait(timeout=5)
   except subprocess.TimeoutExpired:os.killpg(child.pid,signal.SIGKILL);child.wait(timeout=5)
   raise RuntimeError(name+' exceeded its unchanged bound/deadline')
 results.append({'stage':name,'argv':[str(a) for a in args],'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'exit':code,'expected':expected})
 (OUT/'results.json').write_text(json.dumps(results,indent=2))
 print('END',name,'exit',code,flush=True)
 if code!=expected:raise RuntimeError(name+' failed; inspect retained stdout/stderr')
 return (OUT/(name+'.stdout')).read_text()
assert not git(CLONE,'status','--porcelain')
assert git(CLONE,'rev-parse','HEAD')==SOURCE
run('history-delta',[BINARY,'check-commits',SUBJECT+'..'+SOURCE,'--json'],CLONE,300)
run('ci',[BINARY,'ci',SUBJECT+'..'+SOURCE,'--json'],CLONE)
run('version',[BINARY,'version'],CLONE,30)
run('build-info',['go','version','-m',BINARY],CLONE,30)
# Actual compiled-controller regression, with no worker/agent or Drill invocation.
FIXTURE=OUT/'diagnostic-fixture';(FIXTURE/'packages/alpha/evals').mkdir(parents=True);(FIXTURE/'packages/alpha/internal').mkdir()
(FIXTURE/'.fireweed').mkdir();(FIXTURE/'changes/archive').mkdir(parents=True)
(FIXTURE/'go.mod').write_text('module example.test/diagnostic\n\ngo 1.27.0\n')
(FIXTURE/'fireweed.yaml').write_text('''fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: [packages/*]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [go, run, .fireweed/hooks.go, eval, "{module_dir}", "{report}"]
''')
(FIXTURE/'SYSTEM.md').write_text('# Isolated compiler-diagnostic fixture\n')
(FIXTURE/'.fireweed/hooks.go').write_bytes((CLONE/'.fireweed/hooks.go').read_bytes())
(FIXTURE/'packages/alpha/SPEC.md').write_text('''---
name: alpha
id_prefix: ALP
---
# Alpha

## Specs

### ALP-001
Value MUST return text.
''')
(FIXTURE/'packages/alpha/api.go').write_text('package alpha\nfunc Value() string { return strings.TrimSpace("value") }\n')
(FIXTURE/'packages/alpha/evals/value_test.go').write_text('package evals_test\nimport ("testing"; "example.test/diagnostic/packages/alpha")\nfunc Test_ALP_001(t *testing.T) { _ = alpha.Value() }\n')
(FIXTURE/'changes/0001-diagnostic.md').write_text('''---
module: alpha
fixes: [ALP-001]
bases: {}
---
## Motivation
Verify that compiler diagnostics survive invalid JUnit in this isolated fixture.
## Interface
## Added
## Modified
## Removed
## Decisions
''')
run('fixture-init',['git','init','-b','main'],FIXTURE,30)
run('fixture-add',['git','add','.'],FIXTURE,30)
run('fixture-commit',['git','-c','user.name=Fireweed Fixture','-c','user.email=fixture@localhost','commit','-m','diagnostic fixture'],FIXTURE,30)
run('fixture-start',[BINARY,'start','diagnostic','--json'],FIXTURE,30)
run('fixture-verify',[BINARY,'verify','diagnostic','--quick','--json'],FIXTURE,60,expected=1)
reports=list((FIXTURE/'.fireweed/runs').glob('*/evals/alpha.json'));assert len(reports)==1,reports
report=json.loads(reports[0].read_text());assert 'undefined: strings' in report['Stdout'];assert report['Stderr'] and report['HookExit']==1 and not report['Cases'],report
(OUT/'diagnostic-report.json').write_text(json.dumps(report,indent=2))
run('fixture-release',[BINARY,'release','diagnostic','--json'],FIXTURE,30)
assert not git(FIXTURE,'for-each-ref','--format=%(refname)','refs/fireweed/claims')
assert not git(CLONE,'status','--porcelain')
assert sha(OLDPIN)=='23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f'
meta={'source':SOURCE,'subject':SUBJECT,'history_range':SUBJECT+'..'+SOURCE,'reused_history_prefix':BASE+'..'+SUBJECT,'sha256':sha(BINARY),'native_drills':0,'all_checks_passed':True,'compiled_controller_diagnostics':True,'pin_not_yet_changed':True}
(OUT/'complete.json').write_text(json.dumps(meta,indent=2));(OUT/'active.json').unlink();print(json.dumps(meta),flush=True)
