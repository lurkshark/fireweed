from pathlib import Path
import subprocess,json,shutil,hashlib
R=Path('/Users/joel/Workspace/fireweed');O=Path('/private/tmp/fireweed-residual-fixes-27wddhxj');V=O/'validation';P=json.loads((O/'promotion.json').read_text());assert not (V/'active.json').exists();assert hashlib.sha256((R/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==P['sha256'];removed=[]
for n in ['cli-unstarted','cli-recorded','source']:
 p=V/n;git=lambda *a:subprocess.check_output(['git','-C',str(p),*a],text=True,timeout=60).strip()
 assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims');assert git('worktree','list','--porcelain').count('worktree ')==1
 if n=='source':assert not git('status','--porcelain') and git('rev-parse','HEAD')==P['source']
 shutil.rmtree(p);removed.append(n)
receipt={'removed_owned_repositories':removed,'all_fixture_claims_released_before_removal':True,'clean_source_clone_before_removal':True,'new_binary_retained_outside_and_primary_unchanged':True};(O/'validation-cleanup.json').write_text(json.dumps(receipt,indent=2));print(json.dumps(receipt))
