---
module: checks
---

## Motivation

`fireweed doctor --json` lists each probe by a short name such as
`baseline` or `recovery`, and the app's eval for APP-067 requires those
names. Checks' DoctorCheck names include the hook scope after the probe
name, and no checks spec says what a name contains, so the app takes the
first word of a string whose format is unstated. The app drill lists this
as a gap. State the format in the module that produces it.

## Interface

None. DoctorCheck.Name's format is specified.

## Added

### new:probe-names
Each canary DoctorCheck Name MUST begin with its probe name (baseline, undeclared, internal_eval, failing_eval, missing_id or recovery), followed by a space and the hook scope when one applies.

## Modified

## Removed

## Decisions

## Notes

The current names already have this shape for the canary probes, so this
probably passes first; if it does not, the evals step reports the
mismatch. Other DoctorChecks, such as agent probes, are not affected.
Run before 0231.
