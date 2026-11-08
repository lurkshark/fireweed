# Drill evidence retention

Keep one current native report per module, concise handoff/failure reviews,
controller-validation outcomes, proof/ref receipts, cleanup evidence and the
wrapper rehearsal results. The latest workflow report contains both worker
streams and native verification details; separate copies, parsed case catalogs,
source snapshots and repeated prompt/input copies are unnecessary here.

This owner-authorized cleanup removes superseded report bodies, duplicate raw
logs, copied tooling, temporary procedures, input/source snapshots, broad hash
manifests, repeated status/context output and empty files. It changes no product,
spec, eval, controller, configuration, proof/attempt ref or unrelated worktree.
No native drill or history rewrite was run. Existing failed/open milestone status
and unresolved findings remain unchanged.

The latest [workflow review](workflow-native-retry-20261108T090949/review.md)
records the completed supervised native retry, which failed after compilation
succeeded on its final attempt. The single current native report preserves all
four raw worker streams byte-for-byte; compact extraction/gate, result, proof and
cleanup receipts provide their identities without duplicate transcripts or source
snapshots. Earlier reviews and refs remain intact, including the prior
[rejected retry](workflow-contract-retry-20261108T055420/review.md).

## Controller-pin attachments

The owner authorized pruning the six expanded attachments under
`controller-pin-a26f7b9/`. Keep the [compact handoff](controller-pin-a26f7b9.md),
current pin metadata in PROGRESS and the recovery executable. Actual module
drill reports and proof/attempt refs are preserved.

The original expanded handoff, build metadata, input manifest, validation,
promotion, post-install and cleanup receipts remain at commit
`fb86abf44b5e525a7525503f9db9c763bd3d3b5e`. Retrieve an attachment with:

```sh
git show fb86abf:.fireweed/drills/controller-pin-a26f7b9/validation.json
```

## Historical evidence

Files removed in the original drill-retention pass are retained at commit
`d63c16ef3cfcd887a8969985f6cd5cdb4e1deba5`.
Retrieve exact historical bytes without checking out an old implementation:

```sh
git show d63c16ef3cfcd887a8969985f6cd5cdb4e1deba5:.fireweed/drills/PATH
git ls-tree -r --name-only d63c16ef3cfcd887a8969985f6cd5cdb4e1deba5 -- .fireweed/drills
```

Use historical reconstruction/source only as evidence, never as worker input or an
implementation oracle. Small `workflow-before-*.md` files point to the full old
reports in Git. Restoring files is optional; the commands above print them directly.
Git history and retained refs still occupy disk space; this cleanup reduces the
working directory and future checkout contents without rewriting history.

| Evidence group | Files before | Files retained |
| --- | ---: | ---: |
| chk034-controller-handoff | 69 | 9 |
| residual-fixes-controller-handoff | 158 | 8 |
| workflow-contract-diagnostics-handoff | 169 | 7 |
| workflow-contract-retry-20261108T055420 | 269 | 33 |
| workflow-corrected-controller-retry-20261107T232508 | 136 | 9 |
| workflow-failure-investigation-20261108 | 58 | 7 |
| workflow-fixed-controller-retry-20261107T202952 | 127 | 7 |
| workflow-native-retry-20261107T172136 | 47 | 6 |
| workflow-residual-investigation-20261107 | 31 | 5 |
| workflow-superseded-repeat-20261107T213809 | 30 | 4 |
| workflow-wrapper-rehearsal | 52 | 20 |

For future runs, retain one native report plus concise outcome/rejection/cleanup
receipts. Keep scratch transcripts, snapshots and copied tooling outside the repo;
commit additional detail only when it provides unique evidence needed for a live
investigation. Do not duplicate streams already embedded in a retained native report.
