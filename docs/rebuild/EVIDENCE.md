# Requirement evidence matrix

Every inventory row remains incomplete until assertion-level command, workflow, comparison and regeneration evidence is present. This record accumulates actual partial coverage; it does not certify the full design.

| Requirement/clause | Actual spec / permanent assertion | Lifecycle archive | Result / remaining |
| --- | --- | --- | --- |
| R03 test ID matching | CON-001 / Test_CON001Matching: case/separators/name+class, letter/digit boundaries incl Unicode, misleading longer IDs | 0001-spec-primitives | Public API verified, end-to-end JUnit mapping pending |
| R03 grammar/tags | CON-003 / Test_CON003SpecGrammar; CON-004 / Test_CON004TagsAndHandles: ordered bodies, exactly one section, normative text, fences, malformed/duplicate IDs, review/quarantine combinations | 0001-spec-primitives | Public API verified, project validation/status pending |
| R04 bases | CON-002 / Test_CON002BaseHash: exact SHA-256 truncation, CRLF/trailing-space normalization, meaningful body/ID/tag changes | 0001-spec-primitives | Public API verified, stamp/staleness pending |
| R04 local handles | CON-004 / Test_CON004TagsAndHandles: explicit enablement, grammar, duplicate rejection | 0001-spec-primitives | Public API verified, cross-change validation/rewriting pending |

| R02 globs/future overlap/strict zones | CON-005–CON-007 / Test_CON005GlobMatching, Test_CON006FutureOverlap, Test_CON007ZonePrecedence, Test_CON007AuditRegressions | 0002-zone-primitives | Public primitives verified; project graph/ownership/default inheritance and regeneration still pending |

| R02 configuration default/empty-list distinction | CON-008–CON-009 / Test_CON008ConfigDefaults, Test_CON008EmptyCommands, Test_CON009ConfigFieldsAndIsolation, Test_CON009CustomIndependenceAndNull | 0003-config-contract | Configuration API verified; module inheritance/project discovery pending |
| R20 binary/format compatibility | CON-010 / Test_CON010ConfigVersionGate, Test_CON010VersionRegressions | 0003-config-contract | Parser gates verified; CLI enforcement/migration pending |
| R12–R13 configured controls | CON-011 / Test_CON011ConfigValidation | 0003-config-contract | Settings validation verified; actual retries/limits pending |

Audit retry evidence is in phase commits 9bc6825 (four real clause findings), b5624fd (added missing assertions), 7d7f600 (independent pass). Implementation d7dfbd9 did not change interface/evals.

R01–R21 overall: unverified. Setup/checkpoint controller CI and doctor are process evidence, not product functionality. No real regeneration/self-hosting claimed. Configuration initial evals b61d96d, findings 75045c5, retry evals 8537546, passing audit 3d67cc6, impl be9fe8b.

| R02 module defaults and metadata | CON-013 / Test_CON013ModuleInheritance, Test_CON013TypedMetadata | 0004-document-grammar | Parser verified; graph/path ownership validation pending |
| R04/R05 documents and change grammar | CON-012, CON-014–CON-016 / Test_CON012Markdown, Test_CON012IndependentRoundTrip, Test_CON014ChangeMetadata, Test_CON014ClassificationAndRequiredValues, Test_CON015ChangeSections, Test_CON015ProseAndHandleReferences, Test_CON016SetFields, Test_CON016SetAndIDValidation | 0004-document-grammar | Public parser/round-trip verified; projection, stale bases, late IDs and archive application pending |

| R06/R14 local Git reads and paths | GIT-001–GIT-004 / Test_GIT001DiscoveryAndWorktrees, Test_GIT001CanonicalAndNonmutating, Test_GIT002LocalBaseResolution, Test_GIT002NoRemoteContact, Test_GIT003CleanAndDiff, Test_GIT003IsolatedEditsAndOlderBase, Test_GIT004SafePaths, Test_GIT004CanonicalRootAndNoMutation | 0005-local-repository | Read-only primitives verified; claims/commits/scratch/recovery still pending |
