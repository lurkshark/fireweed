from pathlib import Path
import subprocess,json,os,signal,datetime
O=Path(__file__).resolve().parent;assert json.loads((O/'full-runner-result.json').read_text())['exit']==0 and not (O/'active.json').exists();env=dict(os.environ,GOCACHE='/tmp/fireweed-rebuild-cache',PATH='/Users/joel/Workspace/fireweed/.fireweed/bin'+os.pathsep+os.environ['PATH']);rows=[]
for kind in ['candidate','failed']:
 C=O/kind
 assert not subprocess.check_output(['git','-C',str(C),'status','--porcelain'],text=True).strip()
 for module in ['workflow','runner']:
  target=C/'packages'/module/'evals/investigation_probe_test.go';assert not target.exists();target.write_bytes((O/(module+'-probe.txt')).read_bytes())
  try:
   subprocess.run(['gofmt','-w',str(target)],check=True,timeout=30);argv=['go','test','./packages/'+module+'/evals','-run','^Test_Investigation_','-count=1','-timeout=90s','-v'];name=kind+'-'+module+'-probe';start=datetime.datetime.now(datetime.timezone.utc).isoformat()
   with (O/(name+'.stdout')).open('wb') as a,(O/(name+'.stderr')).open('wb') as b:
    p=subprocess.Popen(argv,cwd=C,env=env,stdout=a,stderr=b,start_new_session=True);(O/'active.json').write_text(json.dumps({'stage':name,'pid':p.pid,'argv':argv},indent=2))
    try:code=p.wait(timeout=120)
    except subprocess.TimeoutExpired:
     os.killpg(p.pid,signal.SIGINT)
     try:p.wait(timeout=5)
     except subprocess.TimeoutExpired:os.killpg(p.pid,signal.SIGKILL);p.wait(timeout=5)
     code=p.returncode
   rows.append({'stage':name,'argv':argv,'cwd':str(C),'bound_seconds':120,'start':start,'finish':datetime.datetime.now(datetime.timezone.utc).isoformat(),'exit':code});(O/'probe-results.json').write_text(json.dumps(rows,indent=2));(O/'active.json').unlink();print(name,code,flush=True);assert code==0
  finally:target.unlink()
 assert not subprocess.check_output(['git','-C',str(C),'status','--porcelain'],text=True).strip()
(O/'probes-complete.json').write_text(json.dumps({'public_apis_only':True,'existing_fake_fixture_inputs_reused':True,'original_two_second_test_bound_unchanged':True,'one_nanosecond_counterexample_is_separate_diagnostic_not_native_failure_reproduction':True,'all_probe_claims_released':'test defers public Release and Go fixtures were removed','all_temporary_probe_source_files_removed':True,'tracked_sources_unmodified':True},indent=2))
