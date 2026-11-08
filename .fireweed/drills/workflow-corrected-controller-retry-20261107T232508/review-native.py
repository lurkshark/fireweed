from pathlib import Path
import json,re,subprocess,collections,hashlib
R=Path(__file__).resolve().parent;ROOT=Path('/Users/joel/Workspace/fireweed');terminal=json.loads((R/'terminal.json').read_text());expected=json.loads((R/'expected.json').read_text());p=R/'caller/.fireweed/drills/workflow.md';raw=p.read_bytes();s=raw.decode();assert s.startswith('---\n');fm={k:v for k,v in re.findall(r'^([a-z_]+): (.+)$',s.split('---',2)[1],re.M)};assert fm['base_commit']==expected['base'];(R/'native-report.md').write_bytes(raw);prefix=s.split('## Agent observations',1)[0];(R/'native-fields.txt').write_text(prefix);rows=re.findall(r'^- (workflow|app|runner) \| ([^\n|]+) \| ([^\n|]+) \| ([^\n|]+)$',prefix,re.M);review={};failures=[]
for module in ['workflow','app','runner']:
 spec=subprocess.check_output(['git','-C',str(ROOT),'show',expected['candidate']+':packages/'+module+'/SPEC.md'],text=True,timeout=30).split('## Decisions',1)[0];allids=set(re.findall(r'^### ([A-Z]+-\d+)',spec,re.M));enforced={m[1] for m in re.finditer(r'^### ([A-Z]+-\d+)([^\n]*)',spec,re.M) if 'quarantine' not in m[2]};names=set();unknown=[];covered=set();counts=collections.Counter();cases=[]
 for mod,name,ids,result in rows:
  if mod!=module:continue
  name=name.strip();result=result.strip();ids=re.findall(r'\b[A-Z]+-\d+\b',ids);cases.append({'name':name,'ids':ids,'result':result});counts[result]+=1
  if name in names:unknown.append('duplicate testcase '+name)
  names.add(name)
  unknown.extend(i for i in ids if i not in allids)
  if result!='skip':covered.update(ids)
  if result in ['fail','error']:failures.append({'module':module,'name':name,'ids':ids,'result':result})
 review[module]={'counts':dict(counts),'total':len(cases),'enforced_ids':len(enforced),'covered_enforced_ids':len(enforced&covered),'missing_enforced_ids':sorted(enforced-covered),'unknown_or_duplicate_entries':unknown,'cases':cases}
headings=re.findall(r'^## .+$',prefix,re.M);tail=s[s.rfind('## Verification failures'):] if '## Verification failures' in s else '';result={'terminal':terminal,'frontmatter':fm,'native_report_sha256':hashlib.sha256(raw).hexdigest(),'headings':headings,'modules':review,'failures':failures,'verification_failure_tail':tail};(R/'native-case-review.json').write_text(json.dumps(result,indent=2));print(json.dumps({k:v for k,v in result.items() if k!='modules'},indent=2));print('MODULE SUMMARY',json.dumps({m:{k:v for k,v in d.items() if k!='cases'} for m,d in review.items()},indent=2));print('FIELD CHARS',len(prefix))
