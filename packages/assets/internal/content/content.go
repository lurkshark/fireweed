package content

import "fmt"

func Version() string { return "0.1.0" }

const help = `fireweed 0.1.0

Setup:        init, doctor
Authoring:    new-module, new-change, stamp
Reading:      list, status, show, modules, context, explain, playbook
Checks:       check, coverage, check-commits, ci, pr-body
Lifecycle:    start, verify, done, next, run, release, commits, abandon
Regeneration: drill
Maintenance:  gc, migrate, agents-md
Information:  version
`

func Help() string { return help }

var playbooks = map[string]string{
	"author-change": `# Author change
Read the design, SYSTEM.md and the module spec before writing a change. Give its Motivation a user-visible need or a clause of the design or SYSTEM.md; a probe showing a theoretical failure is insufficient. State each observable behavior at the module interface and keep the change focused. If a requirement costs more than it is worth, narrow or remove the spec and use already_satisfied for behavior the current implementation already meets. Do not commit during a worker step.`,
	"evals": `# Write evals
Read the change, projected spec and public interface. Add tests for every assigned added, modified or fixed spec ID in the evals zone. Exercise observable behavior through the public interface. Run the tests against the current implementation and record genuine fail-first evidence: at least one test for each ID must fail, unless the change marks it already_satisfied. Keep implementation out of this step and do not commit.

For eval_fix, correct only evals, retain coverage of listed IDs, and require all module evals to pass; fail-first is not required.`,
	"audit": `# Audit evals
Read the change, spec and evals only. Report only exact spec clauses with no test assertion; do not raise findings about helpers, fixtures, observers, unnamed edge cases or conventions. Quote each missing clause verbatim and name its spec ID. Re-audits must limit findings to IDs from the prior findings; passed IDs stay passed. If there are no findings, say so and list checked IDs. Edit only the change's Audit section and do not commit.

For a passing audit, use this example format:

## Audit

No findings
Checked IDs: AST-009, AST-010

Replace these illustrative IDs with the actual checked IDs.

For eval_fix, check that corrected tests still assert the listed specs.`,
	"impl": `# Implement
Read the change, projected spec, public interface and evals. Implement only the assigned behavior in the module's disposable code. Keep the public interface, evals, specs and protected files unchanged. Run fireweed verify to check the work; fireweed owns phase commits, so do not commit.

For eval_fix, edit only the current change file and leave implementation unchanged.`,
	"regenerate": `# Regenerate
Read the module spec, Decisions, public interface and evals. Rebuild disposable implementation files from those sources after the controller removes the old files. Run the module's evals and dependent evals through fireweed verify. Do not consult deleted code, prior implementations or history, and do not commit.`,
}

func Playbook(name string) (string, error) {
	if body, ok := playbooks[name]; ok {
		return body, nil
	}
	return "", fmt.Errorf("unknown playbook %q", name)
}

var explanations = map[string]string{
	"C01":   "The project configuration is invalid: fireweed.yaml is missing or has an invalid required setting. Edit the configuration file to correct the reported field, then run fireweed check again.",
	"C02":   "The module graph is invalid because of an empty module name, a duplicate module name, an empty module prefix, a duplicate module prefix, or a dependency that names an unknown module or creates a cycle. Correct module names and prefixes in SPEC.md so they are nonempty and unique. Correct the module dependency in SPEC.md so it names an existing module and the graph is acyclic.",
	"C03":   "A document or spec item is unreadable or malformed. Repair the document syntax or restore the missing file, then check it again.",
	"C04":   "The spec item is non-normative because its body lacks MUST or SHOULD wording. Rewrite the spec with a MUST or SHOULD requirement that states the expected behavior.",
	"C05":   "A spec ID uses a foreign module prefix or duplicates another ID, or a change has a duplicate edit identity. Change a foreign ID to use the owning module prefix, assign a unique unused ID for a duplicate spec ID, or remove or combine the repeated edit.",
	"C06":   "The spec has an invalid review or quarantine tag. Correct the tag to use a valid review or quarantine form, or remove it.",
	"C07":   "The change has a missing Motivation section or its body is empty, or the pending change module target is missing or invalid. Add a Motivation explaining the user-visible or design need, and set the pending change module to an existing module. An eval_fix change mixed with fixes or spec edits is also invalid. Separate eval-fix work from fixes or spec changes.",
	"C08":   "An edit base is missing or stale, or its target ID is absent from the preceding projection. Stamp or update the base against the preceding projection, and correct or remove an edit whose target is absent.",
	"C09":   "A change references an undeclared handle. Declare the handle in Added, or correct or remove the reference.",
	"C10":   "Required coverage is missing a mapped non-skipped testcase for an enforced spec ID. Add and map a non-skipped testcase to that ID, then rerun the eval; the testcase may pass or fail while coverage is established.",
	"C11":   "An eval hook's JUnit report or testcase mapping is invalid or missing, or its exit status disagrees with the reported results; hook launch or setup may also have failed. Fix the hook, report or mapping, then rerun the eval and check that the exit status matches the results.",
	"C12":   "Module metadata or a zone definition is invalid or missing, or local recipe metadata is invalid. Correct the metadata field or define a valid zone in the module specification; repair invalid local recipe metadata in recipe.yaml.",
	"C13":   "A duplicate queue slug or an unparseable numeric queue prefix makes the queue identity invalid; tied numeric positions are allowed. Choose an unused slug with a valid numeric prefix.",
	"C14":   "Doctor could not prepare or clean up its scratch worktree, found a missing or invalid hook, template or canary, found a canary that failed to detect its violation, or encountered a recovery failure after a canary run. Fix the scratch worktree setup if reported, or repair the named hook, template or canary, then rerun doctor to confirm detection and recovery.",
	"C15":   "A worker or audit agent probe failed because its command failed or is unavailable, its model is invalid or unavailable, or login failed; a usage limit may also stop the probe. Check and fix the configured command, model and login, or wait for the usage limit to reset and check access before rerunning doctor. Warning: the audit uses a shared agent command, leaving a shared-command blind spot; configure an independent audit command.",
	"C16":   "Managed guidance in AGENTS.md is missing or stale, or a playbook override has a missing or stale embedded_base. Run agents-md --write to restore the managed guidance; review and update the playbook override's embedded_base against the current embedded playbook.",
	"E01":   "An enforced testcase failed or reported an error. Fix the failing test or implementation, then rerun the eval.",
	"X01":   "A boundary violation was reported for an import crossing a module dependency or interface boundary. Fix the import or declare the required dependency.",
	"X02":   "A required boundaries hook is missing, or a configured boundary hook failed to complete. Configure a boundaries hook for the module; fix and rerun the hook if it fails or times out.",
	"X03":   "Warning: a module declares a boundaries opt-out with a reason. Review the reason and remove the opt-out when it no longer applies; configure a boundary hook if enforcement is needed.",
	"X04":   "Warning: no boundary enforcement is required, and the boundaries hook is absent. No action is needed; configure a boundaries hook if a module later needs enforcement.",
	"U01":   "An invalid argument or option, or an unsupported command, was requested. Check fireweed help and use the supported command syntax and options.",
	"U02":   "An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.",
	"P01":   "A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.",
	"P02":   "A change claim is missing, held by another owner, or unreadable or undecodable because its claim data is invalid. If no claim exists, start an eligible change. If claim data is unreadable, undecodable, or invalid, seek approved human review and recovery before retrying. If another owner holds the claim, wait for the holder to release it or coordinate with that holder to resolve ownership before retrying.",
	"P03":   "The working state has an unapproved alteration. Restore the approved state or have the alteration reviewed and approved through the normal workflow.",
	"P04":   "Warning: approval for this change is not recorded on the base branch. Obtain approval through review before unattended execution.",
	"P05":   "A change base is stale or mismatches the intended state. Refresh the base through the approved workflow, or restore the intended state through that workflow.",
	"P06":   "The change is blocked by unresolved Questions. Review the Questions and resolve them before continuing the change.",
	"P07":   "The shared-base unapproved-work bound has been reached. Review the shared work and synchronize the shared base after approval before starting more unapproved changes.",
	"P09":   "An audit submission is incomplete when No findings or the Checked IDs list is missing. Include No findings and a complete Checked IDs list in the audit submission.",
	"P10":   "Non-blocking regenerate warning: no passing module drill is newer than the module's latest applied implementation. Run a module drill, then review and commit its report.",
	"P2-01": "Genuine fail-first evidence is missing for an assigned spec ID: its new eval did not fail against the pre-implementation code. Write or correct a test assertion for the specified behavior, rerun the eval before implementing it, and record the failing result.",
	"Z01":   "A forbidden phase-zone edit changed a file outside the authorized zone. Restore the forbidden edit or move the work to the appropriate authorized phase and zone.",
}

func Explain(code string) (string, error) {
	if body, ok := explanations[code]; ok {
		return body, nil
	}
	return "", fmt.Errorf("unknown diagnostic code %q", code)
}
