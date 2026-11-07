from pathlib import Path
import hashlib,json,subprocess
R=Path('/Users/joel/Workspace/fireweed');O=Path('/private/tmp/fireweed-residual-fixes-27wddhxj');E=json.loads((O/'entry.json').read_text());git=lambda *a:subprocess.check_output(['git','-C',str(R),*a],text=True,timeout=60).strip()
assert not git('for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((R/'changes').glob('*.md'))
allowed={'docs/rebuild/ACCEPTANCE.md','docs/rebuild/GOAL.md','docs/rebuild/PROGRESS.md','docs/rebuild/workflow-drill/README.md','docs/rebuild/workflow-drill/control.py','docs/rebuild/workflow-drill/preflight.py','packages/workflow/SPEC.md','packages/workflow/evals/regenerate_start_test.go','packages/runner/evals/commit_recovery_test.go','packages/runner/evals/next_test.go'}
changed=[]
for n,h in E['files'].items():
 p=R/n
 if not p.is_file() or hashlib.sha256(p.read_bytes()).hexdigest()!=h:changed.append(n)
assert set(changed)<=allowed,set(changed)-allowed
for line in E['refs'].splitlines():
 oid,ref=line.split(' ',1)
 if ref.startswith(('refs/fireweed/','refs/remotes/')) or ref=='refs/heads/codex/stub-handoff-guidance':assert git('rev-parse',ref)==oid,ref
w=git('worktree','list','--porcelain');other='worktree /Users/joel/.codex/worktrees/stub-handoff-guidance/fireweed\nHEAD b59e9cc15a7d6c82f6436ef72fe5c2c3ccded211\nbranch refs/heads/codex/stub-handoff-guidance';assert other in w and w.count('worktree ')==2
pins={}
for n,h in E['pins'].items():
 if n=='.fireweed/bin/fireweed':continue
 pins[n]=hashlib.sha256((R/n).read_bytes()).hexdigest();assert pins[n]==h,n
assert git('cat-file','-t','79cf33a7690082f28c06bf41a72719d282bc00fe')=='commit'
result={'head':git('rev-parse','HEAD'),'original_baseline':'79cf33a7690082f28c06bf41a72719d282bc00fe','changed_existing_paths':changed,'all_prior_archives_reports_and_proof_attempt_refs_unchanged':True,'unrelated_worktree_identity_unchanged':True,'recovery_hashes':pins,'no_queue_or_shared_claim':True,'configuration_and_worker_instruction_bytes_unchanged':True};(O/'preservation.json').write_text(json.dumps(result,indent=2));print(json.dumps(result,indent=2))
