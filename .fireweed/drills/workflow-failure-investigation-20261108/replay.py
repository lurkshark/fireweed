from pathlib import Path
import subprocess,json,os,signal,datetime,time
O=Path(__file__).resolve().parent;ROOT=Path('/Users/joel/Workspace/fireweed');E=json.loads((O/'entry.json').read_text());env=dict(os.environ,PATH=str(ROOT/'.fireweed/bin')+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache',GIT_NO_LAZY_FETCH='1');rows=[]
def run(name,argv,cwd,bound):
 start=datetime.datetime.now(datetime.timezone.utc).isoformat()
 with (O/(name+'.stdout')).open('wb') as a,(O/(name+'.stderr')).open('wb') as b:
  p=subprocess.Popen(argv,cwd=cwd,env=env,stdout=a,stderr=b,start_new_session=True);(O/'active.json').write_text(json.dumps({'stage':name,'pid':p.pid,'argv':argv},indent=2))
  try:code=p.wait(timeout=bound)
  except subprocess.TimeoutExpired:
   os.killpg(p.pid,signal.SIGINT)
   try:p.wait(timeout=5)
   except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=5)
   code=p.returncode
  rows.append({'stage':name,'argv':argv,'cwd':str(cwd),'bound_seconds':bound,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'exit':code});(O/'results.json').write_text(json.dumps(rows,indent=2));(O/'active.json').unlink();print(name,code,flush=True)
 return code
names=['Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound','Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives','Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses'];pattern='^('+ '|'.join(names)+')$'
for kind,oid in [('candidate',E['candidate']),('failed',E['failed_oid'])]:
 c=O/kind;assert run(kind+'-clone',['git','clone','--no-hardlinks','--no-checkout',str(ROOT),str(c)],ROOT,120)==0;assert run(kind+'-checkout',['git','checkout','--detach',oid],c,60)==0
 assert run(kind+'-workflow',['go','test','./packages/workflow/evals','-run',pattern,'-count=1','-timeout=90s','-v'],c,120) in [0,1]
 assert run(kind+'-runner',['go','test','./packages/runner/evals','-run','^Test_RUN_037_VerificationDeadlineRetainsHeldWork$','-count=1','-timeout=90s','-v'],c,120) in [0,1]
 assert not subprocess.check_output(['git','-C',str(c),'status','--porcelain'],text=True).strip()
(O/'replay-complete.json').write_text(json.dumps({'four_targeted_cases_only':True,'no_native_drill_or_real_agent_invoked':True,'candidate_and_failed_clone_tracked_files_unchanged':True},indent=2))
