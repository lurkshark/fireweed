# Current milestone

**Regeneration is complete as of 2026-11-10.** All eight modules have passing
native reports, and the owner accepted each proof. [PROGRESS.md](PROGRESS.md)
records the evidence and deferred gaps.

**Next: queued changes 0202–0233.** They stop drill retries from seeing
dependent eval output and shrink drill reports (0202), give three dependent
expectations a home in the owning module's spec (0203–0205), repair evals
that pin other modules' wording or depend on timing (0206–0208), make
lower-module errors carry their codes so the app stops parsing message text
(0209–0212, 0214), rename workflow's `*_stub.go` files (0213), and make P08
warn on chains of changes that rework the same spec IDs (0215–0217), and then
remove the code labels from error text (0218–0223). 0224–0231 address
deferred drill gaps: the design's warning for SPEC.md files outside module
globs, fenced-code section boundaries, evals that pin wording, Doctor probe
names and double verification of successful phases. 0232–0233 add the
design's C17–C19 spec-shape warnings. They run only after the owner merges
them to main.

For future rebuild work, read AGENTS.md, SYSTEM.md, PROGRESS.md and the
applicable gates in [ACCEPTANCE.md](ACCEPTANCE.md#full-completion-gates-still-required).
Reuse valid proof. If a relevant change invalidates a proof, await owner
approval before repeating its drill with the verified controller.
