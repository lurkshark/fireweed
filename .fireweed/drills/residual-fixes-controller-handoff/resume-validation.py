from pathlib import Path
import subprocess,os,json,time,datetime,hashlib,signal,re
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path('/private/tmp/fireweed-residual-fixes-27wddhxj');V=OUT/'validation';OLD='c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada';PREFIX='12c571cc53c7e03b051811477adc83e5518f157d';DEADLINE=datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp();env=dict(os.environ,PATH=str(ROOT/'.fireweed/bin')+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache');steps=[]
def git(root,*args):return subprocess.check_output(['git','-C',str(root),*args],text=True,timeout=60).strip()
def run(stage,argv,cwd,bound=120,expected=0):
 assert time.time()<DEADLINE,'owner deadline reached';start=datetime.datetime.now(datetime.timezone.utc).isoformat()
 with (V/(stage+'.stdout')).open('wb',buffering=0) as out,(V/(stage+'.stderr')).open('wb',buffering=0) as err:
  p=subprocess.Popen(argv,cwd=cwd,env=env,stdout=out,stderr=err,start_new_session=True);(V/'active.json').write_text(json.dumps({'stage':stage,'pid':p.pid,'argv':argv,'start':start},indent=2))
  try:code=p.wait(timeout=min(bound,max(1,DEADLINE-time.time())))
  except subprocess.TimeoutExpired:
   os.killpg(p.pid,signal.SIGINT)
   try:p.wait(timeout=10)
   except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=10)
   raise
 steps.append({'stage':stage,'argv':argv,'cwd':str(cwd),'bound_seconds':bound,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'exit':code,'expected':expected});(V/'results.json').write_text(json.dumps(steps,indent=2));(V/'active.json').unlink();assert code==expected,(stage,code,expected);print(stage+' passed',flush=True)
complete=json.loads((OUT/'both-complete.json').read_text());source=complete['source'];SRC=V/'source';BIN=V/'fireweed';assert not (V/'active.json').exists() and not (OUT/'active-owner.json').exists();assert git(ROOT,'rev-parse','HEAD')==source and not git(ROOT,'status','--porcelain') and not git(SRC,'status','--porcelain');assert hashlib.sha256((ROOT/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==OLD;assert not git(ROOT,'for-each-ref','--format=%(refname)','refs/fireweed/claims');steps=json.loads((V/'results.json').read_text());assert steps[-1]['stage']=='normal' and steps[-1]['exit']==1;steps[-1]['stage']='normal-initial';steps[-1]['stdout']='initial-normal-failure/normal.stdout';steps[-1]['stderr']='initial-normal-failure/normal.stderr';focused=json.loads((V/'focused-result.json').read_text());assert focused['exit']==0;steps.append(focused)
run('normal',['go','test','./...'],SRC,900)
run('race',['go','test','-race','./...'],SRC,900)
run('vet',['go','vet','./...'],SRC,180)
run('check',[str(BIN),'check','--json'],SRC,60)
run('doctor',[str(BIN),'doctor','--no-agent','--json'],SRC,180)
delta=PREFIX+'..'+source
run('history-delta',[str(BIN),'check-commits',delta,'--json'],SRC,300)
run('ci',[str(BIN),'ci',delta,'--json'],SRC,1500)
run('version',[str(BIN),'version'],SRC,30)
run('build-info',['go','version','-m',str(BIN)],SRC,30)
info=(V/'build-info.stdout').read_text();assert 'vcs.revision='+source in info and 'vcs.modified=false' in info
for mode in ['unstarted','recorded']:
 fixture=V/('cli-'+mode);fixture.mkdir();files={'SYSTEM.md':'# Public CLI fixture\n','fireweed.yaml':"fireweed: '>=0.1, <0.2'\nversion: 1\nbase: main\nmodules: [packages/*]\nmodule_defaults:\n  interface: [api.go]\n  evals: [evals/**]\n  disposable: [internal/**]\nhooks:\n  eval: [sh, -c, 'exit 1']\n  boundaries: [sh, -c, 'exit 0']\n",'packages/alpha/SPEC.md':'---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n','packages/alpha/api.go':'package alpha\n','packages/alpha/internal/ready':'existing implementation\n','changes/0001-cli.md':'---\nmodule: alpha\nbases: {}\nregenerate: true\n'+('phase: impl\nrun: fixture-recorded\n' if mode=='recorded' else '')+'---\n\n## Motivation\nExercise the approved regenerate entry state.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n'}
 for name,body in files.items():p=fixture/name;p.parent.mkdir(parents=True,exist_ok=True);p.write_text(body)
 run(mode+'-init',['git','init','-b','main'],fixture)
 run(mode+'-add',['git','add','-A'],fixture)
 run(mode+'-commit',['git','-c','user.name=Fireweed Fixture','-c','user.email=fixture@localhost','commit','-m','public regenerate CLI fixture'],fixture)
 started=False
 try:
  run(mode+'-start',[str(BIN),'start','cli','--json'],fixture,60);started=True
  exists=(fixture/'packages/alpha/internal/ready').exists();assert exists==(mode=='recorded'),(mode,exists)
  assert (fixture/'packages/alpha/api.go').read_text()=='package alpha\n'
  change=(fixture/'changes/0001-cli.md').read_text();assert re.search(r'^phase: impl$',change,re.M)
  if mode=='recorded':assert re.search(r'^run: fixture-recorded$',change,re.M)
  run(mode+'-context',[str(BIN),'context','cli','--json'],fixture,60)
  run(mode+'-regenerate-book',[str(BIN),'playbook','regenerate'],fixture,30)
  run(mode+'-impl-book',[str(BIN),'playbook','impl'],fixture,30)
  context=json.loads((V/(mode+'-context.stdout')).read_text())['data'];assert isinstance(context,str);assert (V/(mode+'-regenerate-book.stdout')).read_text().strip() in context;assert (V/(mode+'-impl-book.stdout')).read_text().strip() not in context
 finally:
  if started:run(mode+'-release',[str(BIN),'release','cli','--json'],fixture,60)
 assert not git(fixture,'for-each-ref','--format=%(refname)','refs/fireweed/claims')
assert not git(SRC,'status','--porcelain');assert hashlib.sha256((ROOT/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==OLD;assert not git(ROOT,'for-each-ref','--format=%(refname)','refs/fireweed/claims');assert not list((ROOT/'changes').glob('*.md'))
meta={'source':source,'candidate':source,'original_baseline':'79cf33a7690082f28c06bf41a72719d282bc00fe','history_delta':delta,'reused_valid_prefix':'be35552b706aa3000ff8f11306d129ff1d701b85..12c571cc53c7e03b051811477adc83e5518f157d (composed committed WF2 and CHK-034 delta evidence)','sha256':hashlib.sha256(BIN.read_bytes()).hexdigest(),'all_checks_passed':True,'compiled_cli_unstarted_deletes_recorded_preserves_context_regenerate':True,'all_fixture_claims_released':True,'pin_not_yet_changed':True,'native_drills':0,'initial_normal_failure':'CHK-083 cancel did not return within four-second testcase wait; unchanged source; exact focused replay passed; one full ordinary retry retained','initial_failure_evidence':'initial-normal-failure/'};(V/'complete.json').write_text(json.dumps(meta,indent=2));print(json.dumps(meta),flush=True)
