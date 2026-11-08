from pathlib import Path
import subprocess,json,os,time,datetime,hashlib,signal,sys
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path(__file__).resolve().parent;PIN=ROOT/'.fireweed/bin/fireweed';HASH='20278504a8ea89fcd5f23d1de9a6ec849389fd55707054992c59967eec65e2b1';DEADLINE=datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp();env=dict(os.environ,PATH=str(PIN.parent)+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache');kind=sys.argv[1];names={'workflow':('0183-reserved-ids-review-fail-first.md','reserved-ids-review-fail-first'),'runner':('0184-verification-stop-diagnostics.md','verification-stop-diagnostics')};name,slug=names[kind];git=lambda *a:subprocess.check_output(['git','-C',str(ROOT),*a],text=True,timeout=60).strip()
assert not (OUT/'active-owner.json').exists();assert hashlib.sha256(PIN.read_bytes()).hexdigest()==HASH;assert not git('status','--porcelain');assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');assert sorted(p.name for p in (ROOT/'changes').glob('*.md'))[0]==name
if kind=='runner':assert (ROOT/'changes/archive/0183-reserved-ids-review-fail-first.md').exists()
ctx=subprocess.run([str(PIN),'context',slug,'--json'],cwd=ROOT,env=env,capture_output=True,timeout=30);(OUT/(kind+'-pre-next-context.stdout')).write_bytes(ctx.stdout);(OUT/(kind+'-pre-next-context.stderr')).write_bytes(ctx.stderr);assert ctx.returncode==1
start=datetime.datetime.now(datetime.timezone.utc).isoformat();argv=[str(PIN),'next','--allow-unmerged','--json'];before=git('rev-parse','HEAD')
with (OUT/(kind+'-owner.stdout')).open('xb',buffering=0) as a,(OUT/(kind+'-owner.stderr')).open('xb',buffering=0) as b:
 p=subprocess.Popen(argv,cwd=ROOT,env=env,stdout=a,stderr=b,start_new_session=True);(OUT/'active-owner.json').write_text(json.dumps({'kind':kind,'slug':slug,'pid':p.pid,'argv':argv,'pin_sha256':HASH,'start':start,'head_before':before},indent=2))
 while p.poll() is None:
  if time.time()>=DEADLINE:
   os.killpg(p.pid,signal.SIGINT)
   try:p.wait(timeout=10)
   except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=10)
   raise RuntimeError('owner deadline reached')
  time.sleep(.2)
 result={'kind':kind,'exit':p.returncode,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'head_before':before,'head':git('rev-parse','HEAD')};(OUT/(kind+'-owner-result.json')).write_text(json.dumps(result,indent=2));assert p.returncode==0,result
assert (ROOT/'changes/archive'/name).exists();assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');assert not git('status','--porcelain');assert hashlib.sha256(PIN.read_bytes()).hexdigest()==HASH;(OUT/'active-owner.json').unlink();(OUT/(kind+'-complete.json')).write_text(json.dumps({'source':git('rev-parse','HEAD'),'archive':name,'old_pin_unchanged':True,'no_shared_claim':True,'native_workflow_drills':0},indent=2));print(json.dumps(result),flush=True)
if kind=='runner':
 assert not list((ROOT/'changes').glob('*.md'));(OUT/'both-complete.json').write_text(json.dumps({'source':git('rev-parse','HEAD'),'both_applied':True,'old_pin_unchanged':True,'no_queue_or_shared_claim':True,'native_workflow_drills':0},indent=2));print('Both approved changes are archived; validation and preparation remain.',flush=True)
