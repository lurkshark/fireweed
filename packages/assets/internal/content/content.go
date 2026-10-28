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
Read the change, projected spec and public interface. Add tests for every assigned added, modified or fixed spec ID in the evals zone. Exercise observable behavior through the public interface. Run the tests against the current implementation and record genuine fail-first evidence: at least one test for each ID must fail, unless the change marks it already_satisfied. Keep implementation out of this step and do not commit.`,
	"audit": `# Audit evals
Read the change, spec and evals only. Report only exact spec clauses with no test assertion; do not raise findings about helpers, fixtures, observers, unnamed edge cases or conventions. Quote each missing clause verbatim and name its spec ID. Re-audits must limit findings to IDs from the prior findings; passed IDs stay passed. If there are no findings, say so and list checked IDs. Edit only the change's Audit section and do not commit.`,
	"impl": `# Implement
Read the change, projected spec, public interface and evals. Implement only the assigned behavior in the module's disposable code. Keep the public interface, evals, specs and protected files unchanged. Run fireweed verify to check the work; fireweed owns phase commits, so do not commit.`,
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
	"C04":   "The spec item is non-normative because its body lacks MUST or SHOULD wording. Rewrite the spec with a MUST or SHOULD requirement that states the expected behavior.",
	"C05":   "A spec ID uses a foreign module prefix or duplicates another ID. Change a foreign ID to use the owning module prefix, and assign a unique unused ID for a duplicate.",
	"C06":   "The spec has an invalid review or quarantine tag. Correct the tag to use a valid review or quarantine form, or remove it.",
	"C07":   "The change has a missing Motivation section or its body is empty. Add a Motivation explaining the user-visible or design need for the change.",
	"P2-01": "Genuine fail-first evidence is missing for an assigned spec ID: its new eval did not fail against the pre-implementation code. Write or correct a test assertion for the specified behavior, rerun the eval before implementing it, and record the failing result.",
}

func Explain(code string) (string, error) {
	if body, ok := explanations[code]; ok {
		return body, nil
	}
	return "", fmt.Errorf("unknown diagnostic code %q", code)
}
