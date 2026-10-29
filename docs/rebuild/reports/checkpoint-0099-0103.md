# Checkpoint 0099–0103 — BLOCK

Separate Sol/high checkpoint_0099_0103 reviewed clean source
75f05f8fe58ca40b78a0d1ce4b86f4bdfad7f90b read-only and found one actionable
fail-first bypass. No tests/product/controller commands, edits or drills were
run by the reviewer. Other reviewed areas yielded no additional finding.

## Finding and impact

Workflow stub.go596–609 chooses eval-fix passing-evals whenever eval_fix lists
an ID. Project queue/check.go48–101 does not reject simultaneous fixes or
Added/Modified specs. A change with eval_fix:[ALP-001] and fixes:[ALP-002] can
therefore pass all target evals and archive the fix without a failing testcase.
Design Kinds of change204–209 defines eval_fix as having no spec edits/fixes;
FLOW-023 and R07 preserve fail-first for fixes. This is a concrete uncovered
invalid-kind case, not a helper/style/platform request.

The valid eval-fix, regenerate and ordinary-fix CLI proofs remain valid for
their recorded inputs. They do not establish rejection of mixed kinds. Release
is a library dependency only; its CLI caller/proof stays pending. Full goal and
acceptance remain incomplete.

## Checks at the reviewed clean source

- CGO-disabled build and go vet ./... passed, empty logs. NEW records
  Go1.27.2 darwin/arm64, source75f05f8fe58ca40b78a0d1ce4b86f4bdfad7f90b,
  vcs.modified=false.
- go test -race ./... passed, no race reports: app238.013s, assets18.616s,
  checks244.399s, runner231.022s, workflow132.568s. Unchanged packages reused
  cached results.
- Controller static check passed with the documented C01 embedded-recipe
  SPEC-outside-globs warning. B02 CHK-055/16-case explanation remains.
- Explicit history9949dfe5c26e306ab00dece696f3bc32f44ff747..75f05f8fe58ca40b78a0d1ce4b86f4bdfad7f90b
  passed28 commits, zero diagnostics/squashed entries, extending prior valid
  history ranges. Original baseline remainsbe35552b706aa3000ff8f11306d129ff1d701b85.

Detailed local outputs: /tmp/fireweed-checkpoint-0099-0103-{build.log,
build-info.txt,vet.log,race.log,static.json,history.json}. Checks do not override
the quality finding. No new final CI/doctor/fresh-checkout/drill/self-hosting
claim. No pending governing change or claim; temporary proof fixtures removed.

## Scoped correction proposal — requires owner checkpoint

1. One normal project change rejects eval_fix combined with fixes or spec edits
   as invalid change structure C07. Use representative public regression inputs
   and retain valid eval-fix/fix behavior. Existing Start, Verify and Done call
   project.Check, so central validation rejects the mix before lifecycle work.
2. One scoped assets change extends C07's existing explanation with the invalid
   kind combination and action to split eval-fix from spec/fix changes. Existing
   Motivation/module causes and every assertion remain.
3. Native full verification and independent restricted audits for both changes,
   then this reviewer's scoped re-review and appropriate checks. Continue and
   push corrections only on PASS under the owner's approved correction gate.

Five applications0099–0103 remain since the last passing review; the count is
not reset. Origin/main remains7e0da2758de7a2b9fbfc37536b9dfd34118ea51c,
so shared-base P07 is also at five. No push occurred: standing authorization is
conditional on PASS, and any finding requires owner stop. An explicitly approved
push with this finding recorded would free P07 for the scoped two-change
correction gate; push alone resets no review count. No bound/ref/count change,
phase exception, new drill or broad schema-hardening program is proposed.

Owner explicitly approved: "Approve push and correction gate". The recorded
batch/finding may now be pushed, followed by the two scoped corrections, native
audits/full checks and scoped checkpoint re-review. Corrections are not yet
implemented; review remains BLOCK until re-review passes. No general bound,
phase-write or drill exception is granted.

The explicitly authorized push succeeded: origin/main7e0da27→dac2d16, with
the finding and owner correction-gate approval recorded. Review count remains
five pending the scoped corrections and passing re-review.

The owner-approved0104–0105 corrections and same reviewer scoped re-review now
PASS with actual NEW rejection and fresh checks. See [correction gate](checkpoint-corrections-0104-0105.md). The original finding remains historical.
