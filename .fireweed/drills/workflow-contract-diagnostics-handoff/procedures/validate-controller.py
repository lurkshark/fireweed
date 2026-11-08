from pathlib import Path
import subprocess,os,json,time,datetime,hashlib,signal,re
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path(__file__).resolve().parent;V=OUT/'validation';OLD='20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1';PREFIX='ccd3d6b7732d613e80a027ba7969b65adc58bc0e';DEADLINE=datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp();env=dict(os.environ,PATH=str(ROOT/'.fireweed/bin')+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache');steps=[]
def git(root,*args):return subprocess.check_output(['git','-C',str(root),*args],text=True,timeout=60).strip()
def run(stage,argv,cwd,bound=120,expected=0):
 assert time.time()<DEADLINE;start=datetime.datetime.now(datetime.timezone.utc).isoformat()
 with (V/(stage+'.stdout')).open('wb',buffering=0) as out,(V/(stage+'.stderr')).open('wb',buffering=0) as err:
  p=subprocess.Popen(argv,cwd=cwd,env=env,stdout=out,stderr=err,start_new_session=True);(V/'active.json').write_text(json.dumps({'stage':stage,'pid':p.pid,'argv':argv,'start':start},indent=2))
  try:code=p.wait(timeout=min(bound,max(1,DEADLINE-time.time())))
  except subprocess.TimeoutExpired:
   os.killpg(p.pid,signal.SIGINT)
   try:p.wait(timeout=5)
   except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=5)
   raise
 row={'stage':stage,'argv':argv,'cwd':str(cwd),'bound_seconds':bound,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'exit':code,'expected':expected};steps.append(row);(V/'results.json').write_text(json.dumps(steps,indent=2));(V/'active.json').unlink();assert code==expected,(stage,code,expected);print(stage+' passed',flush=True)
complete=json.loads((OUT/'both-complete.json').read_text());assert complete['both_applied'] and complete['old_pin_unchanged'];assert not (OUT/'active-owner.json').exists();source=complete['source'];assert git(ROOT,'rev-parse','HEAD')==source and not git(ROOT,'status','--porcelain');assert not git(ROOT,'for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((ROOT/'changes').glob('*.md'));assert hashlib.sha256((ROOT/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==OLD
V.mkdir(exist_ok=False);SRC=V/'source';BIN=V/'fireweed';run('clone',['git','clone','--no-hardlinks','--no-checkout',str(ROOT),str(SRC)],ROOT);run('checkout',['git','checkout','--detach',source],SRC);assert not git(SRC,'status','--porcelain');run('build',['go','build','-trimpath','-o',str(BIN),'.'],SRC,180);run('normal',['go','test','./...'],SRC,900);run('race',['go','test','-race','./...'],SRC,900);run('vet',['go','vet','./...'],SRC,180);run('check',[str(BIN),'check','--json'],SRC,60);run('doctor',[str(BIN),'doctor','--no-agent','--json'],SRC,180);delta=PREFIX+'..'+source;run('history-delta',[str(BIN),'check-commits',delta,'--json'],SRC,300);run('ci',[str(BIN),'ci',delta,'--json'],SRC,1500);run('version',[str(BIN),'version'],SRC,30);run('build-info',['go','version','-m',str(BIN)],SRC,30);info=(V/'build-info.stdout').read_text();assert 'vcs.revision='+source in info and 'vcs.modified=false' in info
for prior in [False,True]:
 mode='prior-removed' if prior else 'current-removed';F=V/('cli-'+mode);F.mkdir();files={'SYSTEM.md':'# Public CLI validation fixture\n','fireweed.yaml':"fireweed: '>=0.1, <0.2'\nversion: 1\nbase: main\nmodules: [packages/*]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\nhooks:\n  eval: [sh, .fireweed/eval.sh, '{module}', '{report}']\n  boundaries: [sh, -c, 'exit 0']\n",'packages/alpha/SPEC.md':'---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the old value.\n\n### ALP-002\nAlpha MUST retain a second value.\n','packages/alpha/api.go':'package alpha\n','.fireweed/eval.sh':'''#!/bin/sh
printf '<testsuite><testcase name="ALP-001"><failure message="new value not implemented"/></testcase></testsuite>' > "$2"
exit 1
''','changes/0001-cli.md':'''---
module: alpha
bases: {}
---

## Motivation
Exercise existing ID reservation and review fail-first exclusion.

## Interface

## Added

### new:extra [review]
Extra MUST be reviewed.

## Modified

### ALP-001
Alpha MUST return the revised value.

## Removed

### ALP-002

## Decisions
'''}
 if prior:files['changes/archive/0000-prior.md']='''---
module: alpha
phase: applied
bases:
  ALP-003: 00000000
---

## Motivation
Previously removed third value.

## Interface

## Added

## Modified

## Removed

### ALP-003

## Decisions
'''
 for n,body in files.items():p=F/n;p.parent.mkdir(parents=True,exist_ok=True);p.write_text(body)
 run(mode+'-init',['git','init','-b','main'],F);run(mode+'-stamp',[str(BIN),'stamp','cli','--json'],F,60);run(mode+'-add',['git','add','-A'],F);run(mode+'-commit',['git','-c','user.name=Fireweed Fixture','-c','user.email=fixture@localhost','commit','-m','approved public CLI fixture'],F);started=False
 try:
  run(mode+'-start',[str(BIN),'start','cli','--json'],F,60);started=True;change=(F/'changes/0001-cli.md').read_text();want='ALP-004' if prior else 'ALP-003';assert 'new:extra' not in change and '### '+want+' [review]' in change and '### ALP-002' in change and re.search(r'^phase: evals$',change,re.M)
  run(mode+'-verify',[str(BIN),'verify','--quick','--json'],F,90);data=json.loads((V/(mode+'-verify.stdout')).read_text());cases=[]
  def walk(x):
   if isinstance(x,dict):
    d={k.lower():v for k,v in x.items()}
    if 'name' in d and 'result' in d:cases.append(d)
    for v in x.values():walk(v)
   elif isinstance(x,list):
    for v in x:walk(v)
  walk(data);assert len(cases)==1 and cases[0]['name']=='ALP-001' and cases[0]['result']=='fail',cases
 finally:
  if started:run(mode+'-release',[str(BIN),'release','cli','--json'],F,60)
 assert not git(F,'for-each-ref','--format=%(refname)','refs/fireweed/claims')
assert not git(SRC,'status','--porcelain') and git(ROOT,'rev-parse','HEAD')==source and not git(ROOT,'status','--porcelain');assert hashlib.sha256((ROOT/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==OLD;assert not git(ROOT,'for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((ROOT/'changes').glob('*.md'));meta={'source':source,'candidate':source,'previous_subject':PREFIX,'original_baseline':'79cf33a7690082f28c06bf41a72719d282bc00fe','history_delta':delta,'reused_valid_prefix':'be35552b706aa3000ff8f11306d129ff1d701b85..ccd3d6b7732d613e80a027ba7969b65adc58bc0e (composed committed WF2, CHK-034 and residual-fix delta evidence; history runtime unchanged)','sha256':hashlib.sha256(BIN.read_bytes()).hexdigest(),'all_checks_passed':True,'compiled_cli_current_and_prior_removed_ids_reserved_review_no_fail_first':True,'all_fixture_claims_released':True,'pin_not_yet_changed':True,'native_workflow_drills':0};(V/'complete.json').write_text(json.dumps(meta,indent=2));print(json.dumps(meta),flush=True)
