// Project hooks for Node's TypeScript test runner and Fireweed module boundaries.
import fs from 'node:fs';
import path from 'node:path';
import { spawnSync } from 'node:child_process';
import { run } from 'node:test';

const slash = value => value.split(path.sep).join('/');
const esc = value => String(value).replace(/[&<>"']/g, c => ({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&apos;'}[c]));
function walk(dir) {
  if (!fs.existsSync(dir)) return [];
  return fs.readdirSync(dir, {withFileTypes:true}).flatMap(e => {
    const file = path.join(dir, e.name);
    return e.isDirectory() ? walk(file) : [file];
  });
}
function graph() {
  const bin = process.env.FIREWEED_BIN || 'fireweed';
  const run = spawnSync(bin, ['modules', '--json'], {encoding:'utf8'});
  if (run.status !== 0 || run.error) throw Error(`query modules: ${run.stderr || run.error}`);
  return JSON.parse(run.stdout).data;
}
function matches(file, patterns) {
  return (patterns || []).some(p => p.endsWith('/**') ? file.startsWith(p.slice(0,-3) + '/') : p.includes('*') ? new RegExp('^' + p.split('*').map(x => x.replace(/[.*+?^${}()|[\]\\]/g,'\\$&')).join('[^/]*') + '$').test(file) : p === file);
}
function moduleFor(imported, sourceFile, modules) {
  for (const target of modules) {
    const root = slash(path.normalize(target.dir));
    const candidate = imported.startsWith('.') ? slash(path.normalize(path.join(path.dirname(sourceFile), imported))) : slash(path.normalize(imported));
    if (candidate === root || candidate.startsWith(root + '/')) return {target, rel:candidate.slice(root.length).replace(/^\//,'')};
    if (imported === target.name) return {target, rel:'src/index.ts'};
    if (imported.startsWith(target.name + '/')) return {target, rel:imported.slice(target.name.length+1)};
  }
  return null;
}
function boundaries() {
  const modules = graph();
  let violations = [];
  for (const source of modules) {
    for (const file of walk(source.dir).filter(x => /\.(?:ts|tsx|js|mjs)$/.test(x) && !x.includes('/node_modules/'))) {
      const rel = slash(path.relative(source.dir, file));
      const fromEval = matches(rel, source.zones?.evals);
      const body = fs.readFileSync(file,'utf8');
      const imports = /(?:\bimport\s*(?:[^'";]*?\s+from\s*)?|\bexport\s+[^'";]*?\s+from\s*|\brequire\s*\(|\bimport\s*\()\s*['"]([^'"]+)['"]/g;
      for (const hit of body.matchAll(imports)) {
        const found = moduleFor(hit[1], file, modules);
        if (!found) continue;
        const {target, rel:toRel} = found;
        const declared = source.name === target.name || (source.depends_on || []).includes(target.name);
        if (!declared || (fromEval && !matches(toRel, target.zones?.interface)))
          violations.push(`${source.name}: ${slash(file)} imports ${hit[1]} (module ${target.name}) outside declared dependencies or retained interface`);
      }
    }
  }
  for (const v of violations.sort()) console.log(v);
  if (violations.length) throw Error(`${violations.length} boundary violation(s)`);
}
function evals(dir, report) {
  const files = walk(path.join(dir,'evals')).filter(x => /\.test\.[cm]?[jt]s$/.test(x)).sort();
  const cases = [];
  let failed = false;
  for (const file of files) {
    const child = spawnSync(process.execPath, [process.argv[1], 'child', file], {encoding:'utf8'});
    const started = new Map();
    let fileFailed = false;
    const identity = data => JSON.stringify([data.name, data.file, data.line, data.column, data.nesting]);
    for (const line of (child.stdout || '').split(/\r?\n/)) {
      let event;
      try { event = JSON.parse(line); } catch { continue; }
      const {type, data} = event;
      if (type === 'test:dequeue' && data.file) started.set(identity(data), data);
      if (type !== 'test:complete') continue;
      // Node represents a failed module load as a file-level test without a source location.
      if (!data.file) {
        fileFailed = true;
        cases.push({name:'CompileOrCollectionFailure', classname:'collection', time:'0', state:'fail', detail:`${file}: ${data.error || child.stderr || 'test file failed to load'}`});
        continue;
      }
      started.delete(identity(data));
      const state = data.details?.passed ? (data.skip || data.todo ? 'skip' : 'pass') : 'fail';
      if (state === 'fail') fileFailed = true;
      cases.push({name:data.name, classname:slash(file), time:String((data.details?.duration_ms || 0)/1000), state, detail:data.error || ''});
    }
    for (const data of started.values()) {
      fileFailed = true;
      cases.push({name:data.name, classname:slash(file), time:'0', state:'fail', detail:'test started but did not finish'});
    }
    if ((child.status !== 0 || child.error) && !fileFailed) {
      fileFailed = true;
      cases.push({name:'CompileOrCollectionFailure', classname:'collection', time:'0', state:'fail', detail:`${file}: ${child.stderr || child.error || child.stdout}`});
    }
    failed ||= fileFailed;
    if (child.stderr) process.stderr.write(child.stderr);
  }
  if (!files.length) { failed = true; cases.push({name:'CollectionFailure',classname:'collection',time:'0',state:'fail',detail:`${dir}: no eval files found`}); }
  const xml = '<testsuite>' + cases.map(c => `<testcase name="${esc(c.name)}" classname="${esc(c.classname)}" time="${esc(c.time)}">${c.state === 'fail' ? `<failure>${esc(c.detail)}</failure>` : c.state === 'skip' ? '<skipped/>' : ''}</testcase>`).join('') + '</testsuite>';
  fs.writeFileSync(report, xml);
  if (failed) throw Error('TypeScript eval failed');
}
function child(file) {
  const stream = run({files:[file], isolation:'none'});
  let failed = false;
  for (const type of ['test:dequeue', 'test:complete']) stream.on(type, data => {
    if (type === 'test:complete' && !data.details?.passed) failed = true;
    fs.writeSync(1, JSON.stringify({type, data:{...data, error:data.details?.error ? String(data.details.error.stack || data.details.error) : ''}}) + '\n');
  });
  stream.on('end', () => { if (failed) process.exitCode = 1; });
  stream.resume();
}
try {
  if (process.argv[2] === 'eval') evals(process.argv[3],process.argv[4]);
  else if (process.argv[2] === 'boundaries') boundaries();
  else if (process.argv[2] === 'child') child(process.argv[3]);
  else throw Error('usage: hooks.mjs eval <module-dir> <report> | boundaries');
} catch (error) { console.error(error.message); process.exitCode = 1; }
