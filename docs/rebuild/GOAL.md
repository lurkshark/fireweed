# Restart prompt

Complete native regeneration proofs for `runner`, `app`, and the modules
listed for proof renewal in [PROGRESS.md](PROGRESS.md). Read AGENTS.md,
SYSTEM.md (including its scope guardrails), PROGRESS.md and the regeneration
gate in [ACCEPTANCE.md](ACCEPTANCE.md). Check `fireweed status` for the queue
and claims, and confirm the controller pin recorded in PROGRESS.

Await the owner's instruction to start drill execution. Follow the
dependency-first order in PROGRESS and reuse proofs whose retained refs and
coverage still satisfy acceptance. For each selected module, use
`fireweed drill <module>` with the configured worker, commit the native report,
retain its rebuilt-code ref, and confirm scratch cleanup. Review own and
dependent eval results, contradicted Decisions, and gaps with the owner.
Stop when the selected work is done or needs an owner decision and update
PROGRESS with the new state.
