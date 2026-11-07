from pathlib import Path
import json,sys
r=Path('/private/tmp/fireweed-workflow-fixed-20261107-0trl531y/run');n=int(sys.argv[1]);gate=r/'gate';p=gate/f'attempt-{n}/stderr.bin';data=p.read_bytes() if p.exists() else b'';review=r/f'trace-review-{n}';review.mkdir(exist_ok=True);cursor=review/'cursor.json'
if cursor.exists():start=json.loads(cursor.read_text())['offset']
else:
 prompt=json.loads((gate/f'ready-{n}.json').read_text())['argv'][-1].encode();idx=data.find(prompt)
 if idx<0:print('Waiting for exact echoed prompt');sys.exit(0)
 start=idx+len(prompt);(review/'prompt-identity.json').write_text(json.dumps({'exact_prompt_echo':True,'offset':idx,'bytes':len(prompt)}))
end=min(len(data),start+22000)
if end<len(data):
 newline=data.rfind(b'\n',start,end)
 if newline>start:end=newline+1
batch=data[start:end];num=len(list(review.glob('batch-*.txt')))+1
if batch:(review/f'batch-{num:03d}.txt').write_bytes(batch)
cursor.write_text(json.dumps({'offset':end,'observed_size':len(data),'batch':num if batch else num-1}));print(f'attempt {n}: bytes {start}..{end} of {len(data)}');print(batch.decode(errors='replace'));print('gate:',sorted(x.name for x in gate.iterdir() if x.suffix=='.json' or x.name=='abort'))
