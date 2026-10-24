---
module: app
---

## Motivation

Expose the newly checked read-only workflow verification through the independent executable. Preserve its approved-source selection, full phase evidence and guarded hook behavior while making argument, stream, exit and wire-format behavior explicit.

## Interface

Expose verify [<change>] through the existing app.Run and executable entry point. No public Go signature or retained model change is required. JSON data carries Verification's change, phase, diagnostics, coverage, edits and optional findings; CLI edit entries use status, path and optional nonempty old_path, matching the established reference command wire format rather than gitstore.Edit's unspecified default Go field spelling.

## Added

### new:verify-command
verify with zero or one exact slug MUST compose public workflow.Verify from the current directory, including nested and linked SHA-1/SHA-256 checkouts, using implicit held ownership when the operand is absent. It MUST NOT independently load current project metadata, select a different claim, project again or advance the workflow. The returned data MUST retain the full public Change/Phase, ordered Coverage and every nested testcase/ID/mapping/hook exit/output/diagnostic, actual Edits and ordered Findings with ID/Clause/Text/Untestable. A nil public error with no error-severity verification diagnostics MUST return exit0, including legitimate evals fail-first, warnings alone, passing audit and valid audit findings. Any structured checking error MUST return exit1 with all available data/diagnostics intact. Approved hooks and dependent policy MUST remain authoritative despite invalid or worker-edited current metadata; evals, audit and impl MUST retain public Verify's hook selection and guard behavior.

### new:verify-wire
JSON verify MUST emit exactly one schema_version 1 canonical verify/data/non-null diagnostics envelope on stdout and empty stderr. Data MUST use change, phase, diagnostics, coverage and edits, with non-null collection values, and findings only when nonempty. Coverage and findings MUST retain public wire values; edit entries MUST use status and path plus old_path only when nonempty. Envelope diagnostics MUST equal verification diagnostics in their public order. Plain success or structured checking failure MUST render pretty JSON of that same complete wire data on stdout and coded/path-attributed diagnostics on stderr. Hook stdout/stderr MUST remain solely in Coverage.Output. Fatal public errors MUST have null data, preserve the full useful message as one error diagnostic and invent no path: leading C05/P02 MUST retain their code and exit1, leading C01/D05 MUST retain their code and exit2, other messages MUST use C01/exit2. Plain fatal failure MUST have empty stdout. Output writer failures MUST return exit2 for data, diagnostic, fatal and usage output, retaining recoverable hook changes and held state.

### new:verify-arguments
verify MUST accept --json before or after command/operand and valid --help/-h without repository/content access or hook startup. Excess operands, unknown or valued boolean flags, show-only --projected/--history and coverage-only --with-change MUST produce corrective U01 exit2 before repository/content access. -- MUST stop option parsing so subsequent flag-looking text is an operand and cannot enable JSON. A supplied invalid exact slug MUST retain public C05-before-I/O behavior rather than become an implicit selection or a command alias. Existing command/help/flag precedence MUST remain intact. Missing ownership for an operand-free call MUST remain a public checking failure rather than a prompt or a usage error.

### new:verify-safety
Apart from explicitly approved verification hooks, successful and rejected verify calls MUST remain local, bounded and noninteractive, and add no file/ignored/outside-content/index/HEAD/branch/ref/claim/common-Git/temp/cwd/environment changes beyond public Verify and output. The CLI MUST NOT independently execute setup/eval/boundary hooks, agents, bootstrap or network commands, create commits/branches/worktrees, stage/reset an index, repair hook changes or update/release claims. Current blocked, unsafe paths and forbidden worker edits MUST prevent hooks through public guards; post-hook forbidden edits and ownership/HEAD changes MUST remain visible and retained. Owned verification temporary storage MUST be removed on success and failure. FIREWEED_BIN observed by an approved hook MUST identify the actual new executable rather than an inherited controller path. Output failures MUST NOT discard or advance the held transaction.

## Decisions

Wire formatting belongs to the app composition root. The low-level public Go Edit value retains its current interface; this command renders all of its values using the existing lowercase CLI keys. Findings retain their explicit boolean untestable field. Public Verify owns checking, approved-source restoration, hook execution and locality; the CLI owns arguments, wire representation, streams and exits.

This slice does not implement Done, automatic retries, signal/time-limit exit policy, or claim that real regeneration/self-hosting is complete. Those remain separately authored work. The goal's local author authority permits manual execution with the normal unmerged warning; no human merge is asserted.

## Rejected alternatives

Reloading the dirty current project can replace committed approval or fail before returning useful guard evidence. Converting every failed testcase into a command failure would reject legitimate fail-first verification. Dropping structured data on failed checks would hide the evidence needed to fix them. Default Go edit field names would silently change the established CLI wire contract.
