from pathlib import Path
import subprocess,json,os,datetime
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path('/private/tmp/fireweed-residual-review-i2np9tg8');SRC=OUT/'failed-ref';REF='refs/fireweed/drills/workflow/20261107T202952.649365000';OID='50bd33c5526fb495d541e3290868d198ed010b78';steps=[]
def run(stage,argv,cwd,expected=0,bound=120):
 start=datetime.datetime.now(datetime.timezone.utc).isoformat()
 with (OUT/(stage+'.stdout')).open('wb') as out,(OUT/(stage+'.stderr')).open('wb') as err:
  p=subprocess.Popen(argv,cwd=cwd,env=dict(os.environ,PATH=str(ROOT/'.fireweed/bin')+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache'),stdout=out,stderr=err,start_new_session=True)
  try:code=p.wait(timeout=bound)
  except subprocess.TimeoutExpired:
   import signal
   os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=10);raise
 steps.append({'stage':stage,'argv':argv,'cwd':str(cwd),'bound_seconds':bound,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'exit':code,'expected':expected});(OUT/'steps.json').write_text(json.dumps(steps,indent=2));assert code==expected,(stage,code,expected)
run('clone',['git','clone','--no-hardlinks','--no-checkout',str(ROOT),str(SRC)],ROOT)
run('retain-failed-ref',['git','fetch','--no-tags','--no-write-fetch-head',str(ROOT),REF+':'+REF],SRC)
run('checkout',['git','checkout','--detach',OID],SRC)
assert subprocess.check_output(['git','status','--porcelain'],cwd=SRC,text=True).strip()==''
run('workflow-target',['go','test','./packages/workflow/evals','-run','^Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook$','-count=1','-timeout=60s'],SRC,expected=1,bound=90)
run('runner-target',['go','test','./packages/runner/evals','-run','^Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle$','-count=1','-timeout=90s'],SRC,expected=1,bound=120)
assert subprocess.check_output(['git','status','--porcelain'],cwd=SRC,text=True).strip()==''
(OUT/'reproduction.json').write_text(json.dumps({'failed_ref':REF,'oid':OID,'tracked_files_unchanged':True,'native_drills':0,'configured_real_agents':0,'bounded_fake_fixture_agents_only':True,'target_failures_reproduced':True},indent=2));print('Both targeted failures reproduced without altering tracked files or running a native drill.')
