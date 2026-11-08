from pathlib import Path
import hashlib,json,subprocess,sys
n=int(sys.argv[1])
assert n in [1,2]
root=Path(sys.argv[2]).resolve()
gate=root/'gate'
try:
 record=json.loads((gate/f'ready-{n}.json').read_text())
 inp=Path(record['input']).resolve()
 caller=(root/'caller').resolve()
 expected=json.loads((root/'expected.json').read_text())
 command=['codex','exec','--ephemeral','--ignore-user-config','-m','gpt-6.1-sol','-c','model_reasoning_effort="high"','-s','workspace-write']
 assert record['argv'][:-1]==command
 assert record['fireweed_bin']=='/Users/joel/Workspace/fireweed/.fireweed/bin/fireweed'
 def git(root,*args):return subprocess.check_output(['git','-C',str(root),*args],text=True).strip()
 paths={str(p.relative_to(inp)):p for p in inp.rglob('*') if p.is_file() and '.git' not in p.relative_to(inp).parts}
 if n==1:
  assert set(paths)==set(expected['input']), (set(paths)-set(expected['input']),set(expected['input'])-set(paths))
 else:
  fresh=json.loads((root/'attempt-1-output.json').read_text())
  assert set(paths)==set(expected['input'])|set(fresh)
  assert all(sha_path.exists() and hashlib.sha256(sha_path.read_bytes()).hexdigest()==h for f,h in fresh.items() for sha_path in [inp/f])
 sha=lambda p:hashlib.sha256(p.read_bytes()).hexdigest()
 assert all(not p.is_symlink() for p in inp.rglob('*') if '.git' not in p.relative_to(inp).parts)
 assert all(sha(paths[p])==h for p,h in expected['input'].items())
 assert (inp/'.git').is_dir() and not (inp/'.git').is_symlink()
 assert git(inp,'rev-list','--all','--count')=='1'
 status=git(inp,'status','--porcelain','--untracked-files=all')
 if n==1: assert status==''
 else: assert all(line.startswith('?? ') for line in status.splitlines()),status
 assert set(git(inp,'ls-tree','-r','--name-only','HEAD').splitlines())==set(expected['input'])
 assert git(inp,'remote','-v')==''
 assert not (inp/'.git/objects/info/alternates').exists()
 assert git(inp,'fsck','--full','--no-reflogs','--unreachable')==''
 entries=git(caller,'worktree','list','--porcelain').split('\n\n')
 detached=[Path(e.splitlines()[0].split(' ',1)[1]).resolve() for e in entries if '\ndetached' in e]
 verifiers=[p for p in detached if p != caller]
 assert all(p.parent==(caller/'.git/fireweed/worktrees').resolve() and p.name.startswith('drill-') for p in verifiers)
 assert len(verifiers)==1
 verifier=verifiers[0]
 assert git(verifier,'rev-parse','HEAD')==expected['base']
 for p,h in expected['retained'].items():assert (verifier/p).is_file() and sha(verifier/p)==h,p
 if n==1:
  assert all(not (verifier/p).exists() for p in expected['deleted'])
  changes=git(verifier,'diff','--name-status','HEAD').splitlines()
  assert sorted(changes)==sorted('D\t'+p for p in expected['deleted']),changes
 else:
  fresh_code={p:h for p,h in fresh.items() if p.startswith('packages/workflow/internal/')}
  actual={str(p.relative_to(verifier)):sha(p) for p in (verifier/'packages/workflow/internal').rglob('*') if p.is_file()}
  assert actual==fresh_code,actual
  assert all(line.split('\t')[-1].startswith('packages/workflow/internal/') for line in git(verifier,'diff','--name-status','HEAD').splitlines())
 for p in ['go.mod','go.sum','main.go','.fireweed/hooks.go','.fireweed/setup.sh']:
  assert (verifier/p).is_file(),p
 for module in ['contracts','gitstore','assets','project','checks','runner','app']:
  assert list((verifier/'packages'/module/'internal').rglob('*.go')),module
 pinhash='fea9d492473bbb20883a74fc15ea5921b161727e4ea7096b42b3a5a909d60c65'
 for p in [Path(record['fireweed_bin']),caller/'.fireweed/bin/fireweed',verifier/'.fireweed/bin/fireweed']:
  assert sha(p)==pinhash,str(p)
 summary={'input':str(inp),'verifier':str(verifier),'files':len(paths),'retained_verifier_files':len(expected['retained']),'deleted_disposable_paths':len(expected['deleted']),'input_git_commits':1,'pin':pinhash,'invocation':n,'result':'pass'}
 (gate/f'preflight-{n}.json').write_text(json.dumps(summary,indent=2))
 print(json.dumps(summary))
except BaseException:
 (gate/'abort').touch()
 raise
