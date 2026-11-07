from pathlib import Path
import hashlib,json,os,subprocess,tempfile,time,datetime
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path('/private/tmp/fireweed-chk034-3rjojxry');V=OUT/'validation'
OLD='23e87b00217ab1a6990352393311bd150b7d60d3a4c138b75e285f518fba117f';SUBJECT='79cf33a7690082f28c06bf41a72719d282bc00fe'
def sha(p):return hashlib.sha256(p.read_bytes()).hexdigest()
def git(*args):return subprocess.check_output(['git','-C',str(ROOT),*args],text=True,timeout=60).strip()
meta=json.loads((V/'complete.json').read_text());steps=json.loads((V/'results.json').read_text());assert all(s['exit']==s['expected'] for s in steps)
required={'normal','race','vet','check','doctor','history-delta','ci','fixture-verify','fixture-release','build-info'};assert required<={s['stage'] for s in steps}
assert meta['all_checks_passed'] and meta['compiled_controller_diagnostics'] and meta['subject']==SUBJECT and meta['native_drills']==0
assert git('rev-parse','HEAD')==meta['source'] and not git('status','--porcelain')
assert git('rev-parse',SUBJECT)==SUBJECT
assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((ROOT/'changes').glob('*.md'))
assert json.loads((OUT/'terminal.json').read_text())['exit']==0
assert not (V/'active.json').exists()
info=(V/'build-info.stdout').read_text();assert 'vcs.revision='+meta['source'] in info and 'vcs.modified=false' in info
pin=ROOT/'.fireweed/bin/fireweed';new=V/'fireweed';assert sha(pin)==OLD and sha(new)==meta['sha256'] and meta['sha256']!=OLD
assert time.time()<datetime.datetime(2026,11,7,22,51,5,tzinfo=datetime.timezone.utc).timestamp(),'original deadline reached'
backup=pin.with_name('fireweed-1936346')
if backup.exists():assert sha(backup)==OLD,'existing recovery copy differs; stop'
else:
 with backup.open('xb') as f:f.write(pin.read_bytes());f.flush();os.fsync(f.fileno())
 backup.chmod(pin.stat().st_mode&0o777)
assert sha(backup)==OLD
fd,name=tempfile.mkstemp(prefix='fireweed.promoting.',dir=pin.parent)
with os.fdopen(fd,'wb') as f:f.write(new.read_bytes());f.flush();os.fsync(f.fileno());os.fchmod(f.fileno(),0o755)
assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((ROOT/'changes').glob('*.md'))
os.replace(name,pin)
dirfd=os.open(pin.parent,os.O_RDONLY)
try:os.fsync(dirfd)
finally:os.close(dirfd)
assert sha(pin)==meta['sha256']
for rel in ['docs/rebuild/workflow-drill/control.py','docs/rebuild/workflow-drill/preflight.py']:
 p=ROOT/rel;s=p.read_text();assert s.count(OLD)==1;p.write_text(s.replace(OLD,meta['sha256']))
version=subprocess.run([str(pin),'version'],cwd=ROOT,capture_output=True,text=True,timeout=30);assert version.returncode==0
idle=subprocess.run([str(pin),'verify','--json'],cwd=ROOT,capture_output=True,text=True,timeout=30);assert idle.returncode==1 and 'P02' in idle.stdout and 'worktree holds no change' in idle.stdout
(OUT/'promotion.json').write_text(json.dumps({'source':meta['source'],'subject':SUBJECT,'sha256':meta['sha256'],'version':version.stdout.strip(),'retiring_source':'193634686aa6a1c4eb32e69df1c85e1d5ce48bd7','retiring_recovery':'fireweed-1936346','no_active_run_or_claim':True,'pin_replaced_atomically':True,'outside_pin_metadata_updated_only':True,'native_drills':0,'idle_check':json.loads(idle.stdout)},indent=2))
print(json.dumps(meta));print('Primary controller promoted after validation; original candidate and prior recovery copies preserved.')
