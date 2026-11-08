from pathlib import Path
import sys,json,re,hashlib,datetime
O=Path(__file__).resolve().parent;S=O/'run';A=int(sys.argv[1]) if len(sys.argv)>1 else 1;G=S/'gate';ready=json.loads((G/f'ready-{A}.json').read_text());I=Path(ready['input']);P=O/f'live-review-{A}';P.mkdir(exist_ok=True);state_path=P/'state.json';state=json.loads(state_path.read_text()) if state_path.exists() else {'events':{},'snapshots':{},'pending':'','emitted_chars':0};B=(G/f'attempt-{A}/stderr.bin').read_bytes();T=B.decode('utf-8',errors='replace');events=[]
for kind,pattern in [('command',r'(?m)^exec\n(.*?) in '+re.escape(str(I))+r'\n'),('patch',r'(\*\*\* Begin Patch\n.*?\*\*\* End Patch)'),('file-patch',r'(?m)^apply patch\n(.*?)(?=^(?:apply patch|exec|codex|tokens used)\n)')]:
 if kind=='file-patch' and (G/f'finished-{A}.json').exists():pattern=r'(?m)^apply patch\n(.*?)(?=^(?:apply patch|exec|codex|tokens used)\n|\Z)'
 for m in re.finditer(pattern,T,re.S):
  body=m[1];position=m.start()-1 if kind=='command' and m.start()>0 else m.start();key=kind+':'+str(position);h=hashlib.sha256(body.encode()).hexdigest()
  if key in state['events']:assert state['events'][key]==h
  else:events.append((m.start(),kind,key,h,body))
for pos,kind,key,h,body in sorted(events):
 state['events'][key]=h;number=len(state['events']);text=f'\nEVENT {number} {kind} raw_offset={pos} SHA256={h}\n{body}\n';(P/f'event-{number:04}.txt').write_text(text)
 if kind=='file-patch':
  rendered=[];parts=re.split(r'(?m)(?=^diff --git )',body);seen=state.setdefault('seen_diffs',{})
  for fragment in parts:
   dh=hashlib.sha256(fragment.encode()).hexdigest()
   if fragment.startswith('diff --git '):
    if dh in seen:rendered.append('[identical previously reviewed diff '+dh+']\n');continue
    seen[dh]=number
   rendered.append(fragment)
  text=f'\nEVENT {number} {kind} raw_offset={pos} SHA256={h}\n'+''.join(rendered)+'\n'
 state['pending']+=text
expected=json.loads((S/'expected.json').read_text())['input'];unexpected=[]
if I.exists():
 for name,h in expected.items():assert hashlib.sha256((I/name).read_bytes()).hexdigest()==h,'retained input changed: '+name
 for p in sorted(I.rglob('*')):
  if not p.is_file() or '.git' in p.relative_to(I).parts:continue
  name=str(p.relative_to(I))
  if name in expected:continue
  if not (name=='DRILL_NOTES.md' or name.startswith('packages/workflow/internal/')):unexpected.append(name)
  data=p.read_bytes();h=hashlib.sha256(data).hexdigest()
  if state['snapshots'].get(name)==h:continue
  dest=P/'snapshots'/h;dest.parent.mkdir(exist_ok=True);dest.write_bytes(data);state['snapshots'][name]=h
  (P/'snapshot-ledger.jsonl').open('a').write(json.dumps({'time':datetime.datetime.now(datetime.timezone.utc).isoformat(),'path':name,'sha256':h,'size':len(data)})+'\n')
  # Full writes/patches appear as command events. Retain snapshots independently.
limit=18000;part=state['pending'][:limit];state['pending']=state['pending'][limit:];state['emitted_chars']+=len(part);emission_no=state.get('emissions',0)+1;state['emissions']=emission_no;(P/f'emission-{emission_no:04}.txt').write_text(part);state['raw_bytes']=len(B);state['exec_markers']=len(re.findall(r'(?:^|\n)exec\n',T));state['parsed_commands']=sum(k.startswith('command:') for k in state['events']);state_path.write_text(json.dumps(state,indent=2));print(part);print('\nREVIEW STATUS',json.dumps({'attempt':A,'raw_bytes':len(B),'events':len(state['events']),'parsed_commands':state['parsed_commands'],'exec_markers':state['exec_markers'],'patch_markers':len(re.findall(r'(?m)^apply patch$',T)),'parsed_file_patches':sum(k.startswith('file-patch:') for k in state['events']),'pending_chars':len(state['pending']),'fresh_files':state['snapshots'],'unexpected':unexpected,'finished':(G/f'finished-{A}.json').exists(),'abort':(G/'abort').exists()}));print('STDOUT', (G/f'attempt-{A}/stdout.bin').read_text(errors='replace') if (G/f'attempt-{A}/stdout.bin').exists() else '')
