from pathlib import Path
import subprocess,json,hashlib,shutil,re,datetime
R=Path('/Users/joel/Workspace/fireweed');O=Path(__file__).resolve().parent;D=R/'.fireweed/drills/workflow-contract-diagnostics-handoff';m=json.loads((O/'promotion.json').read_text());assert m['source']==json.loads((O/'both-complete.json').read_text())['source'];assert hashlib.sha256((R/'.fireweed/bin/fireweed').read_bytes()).hexdigest()==m['sha256'];assert not (O/'active-owner.json').exists() and not (O/'validation/active.json').exists();assert (O/'cleanup.json').exists() and (O/'preservation.json').exists();assert not subprocess.check_output(['git','-C',str(R),'for-each-ref','--format=%(refname)','refs/fireweed/claims'],text=True).strip();assert not list((R/'changes').glob('*.md'));D.mkdir(exist_ok=False)
# Only ordinary evidence files, never clones, executables or worker caches.
for p in O.iterdir():
 if p.is_file() and p.suffix not in ['.md']:
  if p.suffix=='.py':q=D/'procedures'/p.name
  else:q=D/p.name
  q.parent.mkdir(parents=True,exist_ok=True);shutil.copy2(p,q)
V=D/'validation';V.mkdir()
for p in (O/'validation').iterdir():
 if p.is_file() and p.name!='fireweed':shutil.copy2(p,V/p.name)
P=D/'prepared';P.mkdir();shutil.copy2(O/'prepared-retry/expected.json',P/'expected.json');shutil.copytree(O/'prepared-retry/tooling',P/'tooling',ignore=shutil.ignore_patterns('__pycache__','README.md'))
runs=[];summaries={}
for name in ['0183-reserved-ids-review-fail-first','0184-verification-stop-diagnostics']:
 text=(R/'changes/archive'/(name+'.md')).read_text();run=re.search(r'^run: (.+)$',text,re.M)[1];runs.append(run);shutil.copytree(R/'.fireweed/runs'/run,D/'runs'/run)
 summary={'run':run,'archive':name+'.md','provenance':[s for s in text.splitlines() if s.startswith('- evals ') or s.startswith('- audit ') or s.startswith('- impl ')],'phase_reports':{},'quick_verify_ms':{}}
 for phase in ['evals','impl']:
  log=(R/'.fireweed/runs'/run/(phase+'-1.log')).read_text();found=re.search(r'succeeded in (\d+)ms:\nverify: '+phase+r'\b',log)
  summary['quick_verify_ms'][phase]=int(found[1]) if found else None
  folder=R/'.fireweed/runs'/run/phase
  for report in folder.glob('*.json'):
   if report.name=='costs.json':continue
   d=json.loads(report.read_text());cases=d.get('Cases',[]);summary['phase_reports'][phase+'/'+report.name]={'cases':len(cases),'results':{k:sum(c.get('Result')==k for c in cases) for k in sorted({c.get('Result','') for c in cases})},'hook_exit':d.get('HookExit'),'stdout':d.get('Stdout',''),'stderr':d.get('Stderr','')}
 summaries[name]=summary
(D/'phase-summary.json').write_text(json.dumps(summaries,indent=2))
files={str(p.relative_to(D)):hashlib.sha256(p.read_bytes()).hexdigest() for p in sorted(D.rglob('*')) if p.is_file()};(D/'manifest.json').write_text(json.dumps({'files':files,'native_workflow_drills':0,'created':datetime.datetime.now(datetime.timezone.utc).isoformat()},indent=2));print(json.dumps({'directory':str(D),'files':len(files),'runs':runs,'source':m['source'],'sha256':m['sha256']}),flush=True)
