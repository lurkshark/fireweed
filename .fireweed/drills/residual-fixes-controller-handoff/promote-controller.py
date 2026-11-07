from pathlib import Path
import subprocess,json,hashlib,os,tempfile,datetime,time
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path('/private/tmp/fireweed-residual-fixes-27wddhxj');V=OUT/'validation';OLD='c9167d3878a1344719574d3ee1179d42295596f2f994c31c556ca25df8aaaada';BASE='79cf33a7690082f28c06bf41a72719d282bc00fe';OLD_SOURCE='12c571cc53c7e03b051811477adc83e5518f157d'
def sha(p):return hashlib.sha256(p.read_bytes()).hexdigest()
def git(*a):return subprocess.check_output(['git','-C',str(ROOT),*a],text=True,timeout=60).strip()
meta=json.loads((V/'complete.json').read_text());steps=json.loads((V/'results.json').read_text());required={'normal','race','vet','doctor','check','history-delta','ci','build-info','unstarted-start','recorded-start','unstarted-context','recorded-context','unstarted-release','recorded-release'};assert required<={s['stage'] for s in steps};assert all(s['exit']==s['expected'] for s in steps if s['stage']!='normal-initial');initial=[s for s in steps if s['stage']=='normal-initial'];assert len(initial)==1 and initial[0]['exit']==1 and any(s['stage']=='chk083-focused' and s['exit']==0 for s in steps);assert any(s['stage']=='normal' and s['exit']==0 for s in steps);assert meta['all_checks_passed'] and meta['all_fixture_claims_released'];assert git('rev-parse','HEAD')==meta['source']==meta['candidate'];assert not git('status','--porcelain');assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');assert not list((ROOT/'changes').glob('*.md'));assert not (OUT/'active-owner.json').exists() and not (V/'active.json').exists();assert json.loads((OUT/'both-complete.json').read_text())['source']==meta['source'];assert time.time()<datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp()
for mode in ['unstarted','recorded']:
 assert not subprocess.check_output(['git','-C',str(V/('cli-'+mode)),'for-each-ref','--format=%(refname)','refs/fireweed/claims'],text=True).strip()
info=(V/'build-info.stdout').read_text();assert 'vcs.revision='+meta['source'] in info and 'vcs.modified=false' in info
pin=ROOT/'.fireweed/bin/fireweed';new=V/'fireweed';assert sha(pin)==OLD and sha(new)==meta['sha256'];assert meta['sha256']!=OLD
entry=json.loads((OUT/'entry.json').read_text())
for n,h in entry['pins'].items():assert sha(ROOT/n)==h,n
backup=pin.with_name('fireweed-12c571c')
if backup.exists():assert sha(backup)==OLD,'existing recovery differs'
else:
 with backup.open('xb') as f:f.write(pin.read_bytes());f.flush();os.fsync(f.fileno())
 backup.chmod(pin.stat().st_mode&0o777)
assert sha(backup)==OLD
fd,name=tempfile.mkstemp(prefix='fireweed.promoting.',dir=pin.parent)
with os.fdopen(fd,'wb') as f:f.write(new.read_bytes());f.flush();os.fsync(f.fileno());os.fchmod(f.fileno(),0o755)
assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((ROOT/'changes').glob('*.md'))
os.replace(name,pin);directory=os.open(pin.parent,os.O_RDONLY)
try:os.fsync(directory)
finally:os.close(directory)
assert sha(pin)==meta['sha256']
p=ROOT/'docs/rebuild/workflow-drill/control.py';s=p.read_text();assert s.count("PIN = '"+OLD+"'")==1 and s.count("CANDIDATE = '"+BASE+"'")==1;s=s.replace("PIN = '"+OLD+"'","PIN = '"+meta['sha256']+"'").replace("CANDIDATE = '"+BASE+"'","CANDIDATE = '"+meta['candidate']+"'");p.write_text(s)
p=ROOT/'docs/rebuild/workflow-drill/preflight.py';s=p.read_text();assert s.count(OLD)==1;p.write_text(s.replace(OLD,meta['sha256']))
version=subprocess.run([str(pin),'version'],cwd=ROOT,capture_output=True,text=True,timeout=30);assert version.returncode==0
idle=subprocess.run([str(pin),'verify','--json'],cwd=ROOT,capture_output=True,text=True,timeout=30);assert idle.returncode==1 and json.loads(idle.stdout)['diagnostics'][0]['code']=='P02'
receipt={'source':meta['source'],'candidate':meta['candidate'],'original_baseline':BASE,'sha256':meta['sha256'],'version':version.stdout.strip(),'retiring_source':OLD_SOURCE,'retiring_sha256':OLD,'retiring_recovery':'fireweed-12c571c','no_active_run_or_claim':True,'pin_replaced_atomically':True,'outside_metadata_fields_only':['control.PIN','control.CANDIDATE','preflight.pinhash'],'idle_check':json.loads(idle.stdout),'native_drills':0};(OUT/'promotion.json').write_text(json.dumps(receipt,indent=2));print(json.dumps(receipt),flush=True)
