from pathlib import Path
import subprocess,json,os,time,datetime,signal,hashlib
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path('/private/tmp/fireweed-residual-fixes-27wddhxj');PIN=ROOT/'.fireweed/bin/fireweed';HASH='c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada';DEADLINE=datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp();env=dict(os.environ,PATH=str(PIN.parent)+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache');git=lambda *a:subprocess.check_output(['git','-C',str(ROOT),*a],text=True,timeout=60).strip()
assert (ROOT/'changes/archive/0181-recorded-regenerate-preservation.md').exists();assert not (OUT/'active-owner.json').exists();assert hashlib.sha256(PIN.read_bytes()).hexdigest()==HASH;assert not git('status','--porcelain');assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');start=datetime.datetime.now(datetime.timezone.utc).isoformat();argv=[str(PIN),'next','--allow-unmerged','--json']
with (OUT/'runner-owner.stdout').open('xb',buffering=0) as a,(OUT/'runner-owner.stderr').open('xb',buffering=0) as b:
 p=subprocess.Popen(argv,cwd=ROOT,env=env,stdout=a,stderr=b,start_new_session=True);(OUT/'active-owner.json').write_text(json.dumps({'pid':p.pid,'slug':'recovery-history-assertions','argv':argv,'pin_sha256':HASH,'start':start},indent=2))
 while p.poll() is None:
  if time.time()>=DEADLINE:
   os.killpg(p.pid,signal.SIGINT)
   try:p.wait(timeout=10)
   except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=10)
   raise RuntimeError('owner deadline reached')
  time.sleep(.2)
 result={'exit':p.returncode,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'head':git('rev-parse','HEAD')};(OUT/'runner-owner-result.json').write_text(json.dumps(result,indent=2));assert p.returncode==0,result
assert (ROOT/'changes/archive/0182-recovery-history-assertions.md').exists();assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');assert not list((ROOT/'changes').glob('*.md'));assert not git('status','--porcelain');assert hashlib.sha256(PIN.read_bytes()).hexdigest()==HASH;(OUT/'active-owner.json').unlink();(OUT/'both-complete.json').write_text(json.dumps({'source':git('rev-parse','HEAD'),'both_applied':True,'old_pin_unchanged':True,'no_queue_or_shared_claim':True,'native_drills':0},indent=2));print(json.dumps(result),flush=True);print('Both scoped changes applied and archived with the old pin unchanged.',flush=True)
