from pathlib import Path
import subprocess,json,hashlib
ROOT=Path('/Users/joel/Workspace/fireweed');O=Path(__file__).resolve().parent;entry=json.loads((O/'entry.json').read_text());g=lambda *a:subprocess.check_output(['git','-C',str(ROOT),*a],text=True,timeout=60).strip();sha=lambda p:hashlib.sha256(p.read_bytes()).hexdigest()
allowed={'docs/rebuild/GOAL.md','docs/rebuild/PROGRESS.md','docs/rebuild/ACCEPTANCE.md','docs/rebuild/workflow-drill/README.md','docs/rebuild/workflow-drill/control.py','docs/rebuild/workflow-drill/preflight.py','packages/workflow/SPEC.md','packages/workflow/evals/lifecycle_test.go','packages/runner/evals/time_limit_test.go'}
changed=[]
for n,h in entry['files'].items():
 p=ROOT/n
 if not p.is_file() or sha(p)!=h:changed.append(n);assert n in allowed,n
oldrefs={ref:oid for oid,ref in (line.split(' ',1) for line in entry['refs'].splitlines())};now={ref:oid for oid,ref in (line.split(' ',1) for line in g('for-each-ref','--format=%(objectname) %(refname)').splitlines())}
for ref,oid in oldrefs.items():
 if ref!='refs/heads/main':assert now.get(ref)==oid,ref
for prefix in ['refs/fireweed/drills/','refs/fireweed/attempts/']:assert {r:o for r,o in now.items() if r.startswith(prefix)}=={r:o for r,o in oldrefs.items() if r.startswith(prefix)},prefix
for n in g('diff','--name-only',entry['head']).splitlines():assert n in allowed or n in {'changes/archive/0183-reserved-ids-review-fail-first.md','changes/archive/0184-verification-stop-diagnostics.md','.fireweed/drills/workflow-contract-diagnostics-handoff.md'} or n.startswith('.fireweed/drills/workflow-contract-diagnostics-handoff/'),n
for oid in ['79cf33a7690082f28c06bf41a72719d282bc00fe','ccd3d6b7732d613e80a027ba7969b65adc58bc0e']:assert g('cat-file','-t',oid)=='commit'
assert not g('for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((ROOT/'changes').glob('*.md'))
wt=g('worktree','list','--porcelain');expected='worktree /Users/joel/.codex/worktrees/stub-handoff-guidance/fireweed\nHEAD b59e9cc15a7d6c82f6436ef72fe5c2c3ccded211\nbranch refs/heads/codex/stub-handoff-guidance';assert expected in wt and wt.count('worktree ')==2
for n,h in entry['pins'].items():
 if n!='.fireweed/bin/fireweed':assert sha(ROOT/n)==h,n
status=json.loads((O/'final-status.json').read_text());before=json.loads((O/'entry-status.json').read_text())['data'];after=status['data'];assert after['queue']==[] and after['unapproved']==before['unapproved']+2
for k,v in before['phase_commits'].items():assert after['phase_commits'][k]==v,k
for name in ['0183-reserved-ids-review-fail-first','0184-verification-stop-diagnostics']:assert after['phase_commits'][name]==3
assert set(after['phase_commits'])==set(before['phase_commits'])|{'0183-reserved-ids-review-fail-first','0184-verification-stop-diagnostics'}
assert sha(ROOT/'.fireweed/bin/fireweed-ccd3d6b')==entry['pins']['.fireweed/bin/fireweed']
result={'no_new_native_proof_or_attempt_refs':True,'old_tracked_files_preserved_except_authorized_paths':changed,'entry_refs_preserved_except_advancing_main':True,'original_and_prior_corrected_commit_objects_preserved':True,'old_archives_reports_proofs_attempts_configuration_unchanged':True,'existing_recoveries_unchanged':True,'unrelated_worktree_unchanged':True,'old_phase_counters_preserved':True,'new_phase_commits_each':3,'unapproved_before':before['unapproved'],'unapproved_after':after['unapproved'],'no_queue_or_shared_claim':True,'worktrees':wt,'native_workflow_drills':0};(O/'preservation.json').write_text(json.dumps(result,indent=2));print(json.dumps(result),flush=True)
