from pathlib import Path
import re,json,sys,difflib,hashlib,subprocess
r=Path('/private/tmp/fireweed-workflow-corrected-controller-20261107-ngrr5qf_/run');n=int(sys.argv[1]);s=(r/f'gate/attempt-{n}/stderr.bin').read_text(errors='replace');ready=json.loads((r/f'gate/ready-{n}.json').read_text());inp=ready['input'];review=r/f'event-review-{n}';review.mkdir(exist_ok=True);statefile=review/'state.json';state=json.loads(statefile.read_text()) if statefile.exists() else {'commands':0,'snapshots':{},'events':0};remaining=24000;shown=0
commands=re.findall(r'^exec\n(.*?) in '+re.escape(inp)+r'\n',s,re.M|re.S)
for ix,c in enumerate(commands,1):
 if ix<=state['commands']:continue
 if len(c)>remaining:break
 print('COMMAND',ix,c);remaining-=len(c);state['commands']=ix
blocks=list(re.finditer(r'^diff --git a/(.+) b/([^\n]+)\n',s,re.M));events=[]
for ix,m in enumerate(blocks):
 end=blocks[ix+1].start() if ix+1<len(blocks) else len(s);body=s[m.end():end];b=re.search(r'^(?:exec|apply patch|codex|file update:|tokens used)\n',body,re.M)
 if b:body=body[:b.start()]
 elif ix+1==len(blocks) and not (r/'gate'/f'finished-{n}.json').exists():continue
 name=m.group(2)
 if 'new file mode ' in body:
  text='\n'.join(re.findall(r'^\+(?!\+\+)(.*)$',body,re.M))+'\n'
 elif (review/'baseline.json').exists():
  original=json.loads((review/'baseline.json').read_text()).get(name,'');old=original.splitlines(keepends=True);result=[];pos=0;lines=body.splitlines(keepends=True);j=0
  index=re.search(r'^index ([0-9a-f]+)\.\.',body,re.M)
  if index and len(index[1])==40:assert hashlib.sha1(b'blob '+str(len(original.encode())).encode()+b'\0'+original.encode()).hexdigest()==index[1],name
  while j<len(lines):
   h=re.match(r'@@ -(\d+)(?:,(\d+))? \+(\d+)(?:,(\d+))? @@',lines[j])
   if not h:j+=1;continue
   start=max(0,int(h[1])-1);assert start>=pos;result.extend(old[pos:start]);pos=start;j+=1
   while j<len(lines) and not lines[j].startswith('@@ '):
    line=lines[j]
    if line.startswith(' '):assert old[pos]==line[1:],(name,pos);result.append(old[pos]);pos+=1
    elif line.startswith('-'):assert old[pos]==line[1:],(name,pos);pos+=1
    elif line.startswith('+'):result.append(line[1:])
    elif line.startswith('\\ No newline'):pass
    else:break
    j+=1
  result.extend(old[pos:]);text=''.join(result)
 else:continue
 events.append((name,text,m.start()))
for ix,(name,text,offset) in enumerate(events,1):
 if ix<=state['events']:continue
 prev=state['snapshots'].get(name,'');a,b=prev,text
 if name.endswith('.go') and prev:
  results=[subprocess.run(['gofmt'],input=t,text=True,capture_output=True,timeout=10) for t in [prev,text]]
  if all(x.returncode==0 for x in results):a,b=[x.stdout for x in results]
 diff=''.join(difflib.unified_diff(a.splitlines(keepends=True),b.splitlines(keepends=True),fromfile=name+' before (format normalized)',tofile=name+' after (format normalized)',n=1))
 if diff:
  f=review/f'patch-{ix:03d}.diff';f.write_text(diff)
  pending=state.get('pending_diff',{});part_start=pending.get('offset',0) if pending.get('event')==ix else 0
  part_end=min(len(diff),part_start+remaining)
  if part_end<=part_start:break
  print('TRACE PATCH',ix,'log offset',offset,'display chars',part_start,'..',part_end,'of',len(diff));print(diff[part_start:part_end]);remaining-=part_end-part_start;shown+=1
  if part_end<len(diff):state['pending_diff']={'event':ix,'offset':part_end};break
 state.pop('pending_diff',None);state['snapshots'][name]=text;state['events']=ix
(review/'commands.json').write_text(json.dumps(commands,indent=2));statefile.write_text(json.dumps(state,indent=2));print('REVIEW',{'commands_seen':len(commands),'commands_presented':state['commands'],'trace_snapshots_seen':len(events),'trace_snapshots_presented':state['events'],'new_diffs':shown,'stderr_bytes':len(s),'gate':sorted(p.name for p in (r/'gate').iterdir() if p.suffix=='.json' or p.name=='abort')})
