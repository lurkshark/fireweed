from pathlib import Path
import subprocess,os,json,time,datetime,hashlib,shutil,signal
ROOT=Path('/Users/joel/Workspace/fireweed');OUT=Path(__file__).resolve().parent;V=OUT/'validation';PIN=ROOT/'.fireweed/bin/fireweed';env=dict(os.environ,PATH=str(PIN.parent)+os.pathsep+os.environ['PATH'],GOCACHE='/tmp/fireweed-rebuild-cache',PYTHONDONTWRITEBYTECODE='1');entry=json.loads((OUT/'entry.json').read_text());meta=json.loads((V/'complete.json').read_text());source=meta['source'];new=meta['sha256'];old=entry['pins']['.fireweed/bin/fireweed'];oldsource='ccd3d6b7732d613e80a027ba7969b65adc58bc0e';deadline=datetime.datetime.fromisoformat('2026-11-10T22:51:05+00:00').timestamp()
def sha(p):return hashlib.sha256(Path(p).read_bytes()).hexdigest()
def git(root,*a):return subprocess.check_output(['git','-C',str(root),*a],text=True,timeout=60).strip()
def record(p,v):
 with p.open('w') as f:json.dump(v,f,indent=2);f.write('\n');f.flush();os.fsync(f.fileno())
def idle():
 assert not (OUT/'active-owner.json').exists() and not (V/'active.json').exists();assert not git(ROOT,'for-each-ref','--format=%(refname)','refs/fireweed/claims');assert not list((ROOT/'changes').glob('*.md'));assert git(ROOT,'rev-parse','HEAD')==source
assert time.time()<deadline;idle();assert not git(ROOT,'status','--porcelain');assert meta['all_checks_passed'] and meta['all_fixture_claims_released'] and meta['compiled_cli_current_and_prior_removed_ids_reserved_review_no_fail_first'];assert sha(V/'fireweed')==new
for p,h in entry['pins'].items():assert sha(ROOT/p)==h,p
required={'build','normal','race','vet','check','doctor','history-delta','ci','version','build-info','current-removed-verify','prior-removed-verify','current-removed-release','prior-removed-release'};rows=json.loads((V/'results.json').read_text());assert required<={r['stage'] for r in rows} and all(r['exit']==r['expected'] for r in rows)
assert not git(V/'source','status','--porcelain') and git(V/'source','rev-parse','HEAD')==source
for name in ['cli-current-removed','cli-prior-removed']:assert not git(V/name,'for-each-ref','--format=%(refname)','refs/fireweed/claims')
retiring=PIN.with_name('fireweed-ccd3d6b')
if retiring.exists():assert sha(retiring)==old
else:
 with retiring.open('xb') as f:f.write(PIN.read_bytes());f.flush();os.fsync(f.fileno())
 retiring.chmod(0o755)
assert sha(retiring)==old
pending=PIN.with_name('fireweed.pending');assert not pending.exists()
with pending.open('xb') as f:f.write((V/'fireweed').read_bytes());f.flush();os.fsync(f.fileno())
pending.chmod(0o755);assert sha(pending)==new;idle();assert sha(PIN)==old;os.replace(pending,PIN);fd=os.open(PIN.parent,os.O_RDONLY);os.fsync(fd);os.close(fd);assert sha(PIN)==new
for name in ['control.py','preflight.py']:
 p=ROOT/'docs/rebuild/workflow-drill'/name;text=p.read_text();assert text.count(old)==1;text=text.replace(old,new)
 if name=='control.py':assert text.count(oldsource)==1;text=text.replace(oldsource,source)
 p.write_text(text)
receipt={'source':source,'sha256':new,'retiring_source':oldsource,'retiring_path':str(retiring),'retiring_sha256':old,'idle_at_promotion':True,'old_recoveries_unchanged':True,'tooling_changes':'only outside candidate and expected pin','native_workflow_drills':0,'time':datetime.datetime.now(datetime.timezone.utc).isoformat()}
for label,args,expected in [('new-version',['version'],0),('idle-verify',['verify','--json'],1)]:
 p=subprocess.run([str(PIN),*args],cwd=ROOT,env=env,capture_output=True,timeout=30);(OUT/(label+'.stdout')).write_bytes(p.stdout);(OUT/(label+'.stderr')).write_bytes(p.stderr);assert p.returncode==expected
 if label=='idle-verify':assert 'P02' in p.stdout.decode()
record(OUT/'promotion.json',receipt)
# Prepare calls only clone/setup/hash generation, never launch or native drill.
S=OUT/'prepared-retry';p=subprocess.run(['python3',str(ROOT/'docs/rebuild/workflow-drill/inspector.py'),'prepare',str(S)],cwd=ROOT,env=env,capture_output=True,timeout=180);(OUT/'prepare.stdout').write_bytes(p.stdout);(OUT/'prepare.stderr').write_bytes(p.stderr);assert p.returncode==0
caller=S/'caller';exp=json.loads((S/'expected.json').read_text());assert exp['candidate']==source and git(caller,'rev-parse','HEAD')==exp['base'];assert git(caller,'diff','--name-only',source).splitlines()==['SYSTEM.md','fireweed.yaml'];assert git(caller,'rev-parse','HEAD^')==source;assert not git(caller,'status','--porcelain');assert sha(caller/'.fireweed/bin/fireweed')==new and sha(PIN)==new
config=(caller/'fireweed.yaml').read_text();original_config=subprocess.check_output(['git','-C',str(caller),'show',source+':fireweed.yaml'],text=True,timeout=30);lines=[line for line in config.splitlines(keepends=True) if not line.startswith('    wrap:')];assert ''.join(lines)==original_config and sum(line.startswith('    wrap:') for line in config.splitlines())==1
assert (caller/'SYSTEM.md').read_bytes()==(ROOT/'docs/rebuild/workflow-drill/worker-system.md').read_bytes()
(OUT/'setup-child.patch').write_bytes(subprocess.check_output(['git','-C',str(caller),'show','--format=fuller','--binary','HEAD'],timeout=30))
for name,h in exp['input'].items():assert sha(caller/name)==h
for name,h in exp['retained'].items():assert sha(caller/name)==h
assert not (S/'gate').exists();assert not git(caller,'for-each-ref','--format=%(refname)','refs/fireweed/claims');assert len(git(caller,'worktree','list','--porcelain').split('worktree '))-1==1
record(OUT/'preparation.json',{'candidate':source,'setup_child':exp['base'],'changed_paths':['SYSTEM.md','fireweed.yaml'],'primary_and_caller_pin_match':True,'expected_retained_hashes_match':True,'gate_created':False,'native_invocation':False,'verifier_copy_not_exercised':True,'future_live_review_still_required':True})
# Keep evidence and setup identity, remove only owned clones after claim/worktree checks.
cleanup=[]
for path in [caller,V/'source',V/'cli-current-removed',V/'cli-prior-removed']:
 assert path.is_relative_to(OUT) and (path/'.git').exists();assert not git(path,'for-each-ref','--format=%(refname)','refs/fireweed/claims');assert len(git(path,'worktree','list','--porcelain').split('worktree '))-1==1
 if path.name in ['caller','source']:assert not git(path,'status','--porcelain')
 cleanup.append({'path':str(path),'head':git(path,'rev-parse','HEAD'),'no_claim':True,'single_worktree':True});shutil.rmtree(path);assert not path.exists()
assert sha(V/'fireweed')==new and sha(PIN)==new;(V/'fireweed').unlink()
record(OUT/'cleanup.json',{'removed_owned_clones':cleanup,'validated_subject_copy_removed_after_promotion':True,'native_workflow_drills':0});idle();assert sha(PIN)==new;print(json.dumps(receipt),flush=True)
