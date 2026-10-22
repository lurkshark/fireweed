# Requirement evidence matrix

Every inventory row remains incomplete until assertion-level command, workflow, comparison and regeneration evidence is present. This record accumulates actual partial coverage; it does not certify the full design.

| Requirement/clause | Actual spec / permanent assertion | Lifecycle archive | Result / remaining |
| --- | --- | --- | --- |
| R03 test ID matching | CON-001 / Test_CON001Matching: case/separators/name+class, letter/digit boundaries incl Unicode, misleading longer IDs | 0001-spec-primitives | Public API verified, end-to-end JUnit mapping pending |
| R03 grammar/tags | CON-003 / Test_CON003SpecGrammar; CON-004 / Test_CON004TagsAndHandles: ordered bodies, exactly one section, normative text, fences, malformed/duplicate IDs, review/quarantine combinations | 0001-spec-primitives | Public API verified, project validation/status pending |
| R04 bases | CON-002 / Test_CON002BaseHash: exact SHA-256 truncation, CRLF/trailing-space normalization, meaningful body/ID/tag changes | 0001-spec-primitives | Public API verified, stamp/staleness pending |
| R04 local handles | CON-004 / Test_CON004TagsAndHandles: explicit enablement, grammar, duplicate rejection | 0001-spec-primitives | Public API verified, cross-change validation/rewriting pending |

R01–R21 overall: unverified. Setup controller doctor is process evidence, not product doctor evidence. No real regeneration/self-hosting claimed.

| R02 globs/future overlap/strict zones | CON-005–CON-007 / Test_CON005GlobMatching, Test_CON006FutureOverlap, Test_CON007ZonePrecedence, Test_CON007AuditRegressions | 0002-zone-primitives | Public primitives verified; project graph/ownership/default inheritance and regeneration still pending |

Audit retry evidence is in phase commits 9bc6825 (four real clause findings), b5624fd (added missing assertions), 7d7f600 (independent pass). Implementation d7dfbd9 did not change interface/evals.
