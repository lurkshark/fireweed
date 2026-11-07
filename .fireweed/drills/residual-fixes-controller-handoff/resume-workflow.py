from pathlib import Path
import subprocess,json,os,time,datetime,signal,hashlib
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path('/private/tmp/fireweed-residual-fixes-27wddhxj');PIN=ROOT/'.fireweed/bin/fireweed';HASH='c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada';DEADLINE=datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp();env=dict(os.environ,PATH=str(PIN.parent)+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache');git=lambda *a:subprocess.check_output(['git','-C',str(ROOT),*a],text=True,timeout=60).strip()
assert hashlib.sha256(PIN.read_bytes()).hexdigest()==HASH;assert git('diff','--name-only')=='changes/0181-recorded-regenerate-preservation.md';assert json.loads((OUT/'audit-scaffold-verify.json').read_text())['diagnostics'][0]['code']=='P09';old=json.loads((OUT/'active-owner.json').read_text())
try:os.killpg(old['pid'],0)
except ProcessLookupError:pass
else:raise AssertionError('old owning process still active')
(OUT/'failed-owner-1.json').write_text(json.dumps(old,indent=2));start=datetime.datetime.now(datetime.timezone.utc).isoformat();argv=[str(PIN),'next','--allow-unmerged','--json']
with (OUT/'workflow-resume.stdout').open('xb',buffering=0) as a,(OUT/'workflow-resume.stderr').open('xb',buffering=0) as b:
 p=subprocess.Popen(argv,cwd=ROOT,env=env,stdout=a,stderr=b,start_new_session=True);(OUT/'active-owner.json').write_text(json.dumps({'pid':p.pid,'slug':'recorded-regenerate-preservation','argv':argv,'pin_sha256':HASH,'start':start,'resume_same_run':True,'audit_invocation':2},indent=2))
 while p.poll() is None:
  if time.time()>=DEADLINE:
   os.killpg(p.pid,signal.SIGINT)
   try:p.wait(timeout=10)
   except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=10)
   raise RuntimeError('owner deadline reached')
  time.sleep(.2)
 result={'exit':p.returncode,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'head':git('rev-parse','HEAD')};(OUT/'workflow-resume-result.json').write_text(json.dumps(result,indent=2));assert p.returncode==0,result
assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');assert not git('status','--porcelain');assert (ROOT/'changes/archive/0181-recorded-regenerate-preservation.md').exists();assert hashlib.sha256(PIN.read_bytes()).hexdigest()==HASH;(OUT/'active-owner.json').unlink();(OUT/'workflow-complete.json').write_text(json.dumps({'source':git('rev-parse','HEAD'),'applied':True,'same_run':'20261107T214650-d187de5ae2adf49791730c83','same_controller':True,'rejected_audit_preserved':True,'successful_audit_invocation':2},indent=2));print(json.dumps(result),flush=True)
