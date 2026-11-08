from pathlib import Path
import subprocess,json,hashlib,os,shutil,datetime
O=Path(__file__).resolve().parent;S=O/'run';R=Path('/Users/joel/Workspace/fireweed');C=S/'caller';A=json.loads((O/'native-analysis.json').read_text());entry=json.loads((O/'entry.json').read_text());ref=A['retained_ref'];oid=A['retained_oid'];g=lambda root,*a:subprocess.check_output(['git','-C',str(root),*a],text=True,timeout=60).strip();sha=lambda p:hashlib.sha256(p.read_bytes()).hexdigest()
assert json.loads((S/'terminal.json').read_text())=={'exit':1,'emergency':False,'rejected':True};assert (O/'inspector-result.json').exists();assert not g(R,'status','--porcelain') and not g(R,'for-each-ref','--format=%(refname)','refs/fireweed/claims') and not list((R/'changes').glob('*.md'))
pids={'inspector':json.loads((O/'inspector-owner.json').read_text())['pid'],'native_controller':json.loads((S/'controller.json').read_text())['pid']}
for n in [1,2]:pids[f'worker-{n}']=json.loads((S/f'gate/worker-{n}.json').read_text())['pid']
checks={}
for name,pid in pids.items():
 try:os.killpg(pid,0);exists=True
 except ProcessLookupError:exists=False
 checks[name]={'pid':pid,'group_exists':exists};assert not exists,(name,pid)
(O/'process-cleanup.json').write_text(json.dumps({'owned_groups':checks,'escaped_or_reparented_processes_not_certified':True},indent=2))
for n,h in entry['pins'].items():assert sha(R/n)==h,n
assert g(C,'rev-parse',ref)==oid;assert not g(C,'for-each-ref','--format=%(refname)','refs/fireweed/claims');assert g(C,'worktree','list','--porcelain').count('worktree ')==1;assert g(C,'diff','--name-only').splitlines()==['.fireweed/drills/workflow.md'];assert not g(C,'ls-files','--others','--exclude-standard');assert sha(C/'.fireweed/drills/workflow.md')==A['native_report_sha256']
first=json.loads((O/'first-worker-review.json').read_text())['fresh_output'];paths=g(C,'ls-tree','-r','--name-only',ref,'--','packages/workflow/internal').splitlines();wanted={p:h for p,h in first.items() if p.startswith('packages/workflow/internal/')};assert set(paths)==set(wanted)
for p,h in wanted.items():b=subprocess.check_output(['git','-C',str(C),'show',ref+':'+p],timeout=60);assert hashlib.sha256(b).hexdigest()==h,p
changed=g(C,'diff','--name-only',A['setup_child'],ref).splitlines();assert all(p.startswith('packages/workflow/internal/') for p in changed),changed;assert g(C,'rev-parse',ref+'^')==A['setup_child'];assert not g(R,'for-each-ref','--format=%(refname)',ref)
# Retain only the native evidence ref/objects. No checkout, merge, fetch or source apply.
pack=O/'retained-native.pack'
with pack.open('xb') as f:
 p=subprocess.run(['git','-C',str(C),'pack-objects','--stdout','--revs'],input=(ref+'\n^'+A['candidate']+'\n').encode(),stdout=f,stderr=subprocess.PIPE,timeout=120);(O/'pack.stderr').write_bytes(p.stderr);assert p.returncode==0
with pack.open('rb') as f:
 p=subprocess.run(['git','-C',str(R),'unpack-objects','-r'],stdin=f,capture_output=True,timeout=120);(O/'unpack.stdout').write_bytes(p.stdout);(O/'unpack.stderr').write_bytes(p.stderr);assert p.returncode==0,p.stderr
subprocess.run(['git','-C',str(R),'update-ref',ref,oid,'0'*40],check=True,timeout=30);assert g(R,'rev-parse',ref)==oid
receipt={'ref':ref,'oid':oid,'parent':A['setup_child'],'candidate':A['candidate'],'pack_sha256':sha(pack),'proof_tree_equals_first_verified_fresh_code':True,'second_rejected_partial_code_not_in_proof':True,'changed_paths':changed,'source_checkout_unchanged':not g(R,'status','--porcelain'),'no_fetch_push_checkout_merge_or_product_source_apply':True};(O/'retained-ref.json').write_text(json.dumps(receipt,indent=2));pack.unlink()
# Native already removed its verifier and input tree. Remove only this idle caller clone.
input_path=Path(json.loads((S/'gate/ready-1.json').read_text())['input']);verifier=Path(json.loads((S/'gate/preflight-1.json').read_text())['verifier']);assert not input_path.exists() and not verifier.exists();assert C.resolve()==(O/'run/caller').resolve() and C.is_relative_to(O);shutil.rmtree(C);assert not C.exists();(O/'cleanup.json').write_text(json.dumps({'native_removed_input':str(input_path),'native_removed_verifier':str(verifier),'owned_caller_removed':str(C),'no_claim_or_linked_worktree_before_caller_removal':True,'owned_groups_exited':True,'retained_ref_verified_before_removal':True,'native_invocations':1,'no_further_release_after_rejection':not (S/'gate/release-3').exists(),'time':datetime.datetime.now(datetime.timezone.utc).isoformat()},indent=2));print(json.dumps(receipt),flush=True)
