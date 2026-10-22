---
module: app
---

## Motivation

The public project queue validator now detects stale intent and weakening of projected enforcement. Expose those checks through the actual executable so authored changes cannot appear valid while the current module graph alone is sound.

## Interface

No public interface change. Compose the existing public project Load, Check and CheckQueue APIs in the check command. modules continues to publish current static graph diagnostics.

## Added

### new:queue-check
check MUST combine current/discovery diagnostics with public queued-change validation, including stale/missing bases, projected enforcement, handle scope and ID/reference errors. The combined diagnostics MUST sort by path, code, severity and message without repeating unchanged preexisting problems, preserving code, severity, corrective message and repository-relative path. It MUST return exit 1 and {valid:false} for any errors; a valid queue or warnings alone MUST return exit 0 and {valid:true}. Plain data MUST go to stdout with coded/path-attributed diagnostics on stderr; JSON MUST use the existing schema-1 envelope with no plaintext on either stream. From a nested directory, it MUST process numeric/slug ordered intent regardless of lexical filenames and MUST NOT modify files, ignored content or refs or run hooks, agents, the bootstrap or network commands. modules MUST retain its current-static-only behavior when queued projected errors exist.

## Modified

## Removed

## Decisions

This is composition of the existing two pure validation surfaces. CheckQueue excludes existing Check diagnostics; the CLI globally sorts their concatenation. modules remains a current graph read, so a structurally valid queue that introduces errors affects check only. Discovery errors from loading malformed queue files remain visible to both commands through Check. Existing C01 and U01 behaviors and options stay intact.

## Rejected alternatives

Checking only the final projected state would miss intermediate enforcement changes. Printing module-level paths for newly introduced queue errors would hide the authored change that needs correction.
