---
already_satisfied:
    - new:queue-malformed-config
module: project
---

## Motivation

Users running `fireweed next` or `fireweed run` need to know when their
configuration is malformed, even when they have no pending changes. In the
`project` reconstruction drill, both commands reported a successful empty
queue because `Queue` skipped configuration validation when `changes/` was
absent.

## Interface

## Added

### new:queue-malformed-config
If `fireweed.yaml` contains malformed YAML, `Queue` MUST return an error whose
`contracts.ErrorCode` is `C01`, even when the `changes/` directory is absent.

## Modified

## Removed

## Decisions

## Notes

Add an eval through the public `project.Queue` API. Use an otherwise valid
project fixture with no `changes/` directory, write malformed YAML to
`fireweed.yaml`, and assert that `contracts.ErrorCode` reports `C01`. The
current implementation loads the project before reading the queue.
