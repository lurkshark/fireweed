---
date: 2026-11-01T04:16:58.861289Z
base_commit: 0b1a36a551073c660e4093ab5aa73729c49a7fbb
agent: codex
model: gpt-6-sol
result: pass
---

## Own evals
- contracts | Test_CON_001_ParseDocument | CON-001 | pass
- contracts | Test_CON_007_RejectInvalidFrontmatter | CON-007 | pass
- contracts | Test_CON_007_RejectInvalidFrontmatter/malformed | CON-007 | pass
- contracts | Test_CON_007_RejectInvalidFrontmatter/non_mapping | CON-007 | pass
- contracts | Test_CON_007_RejectInvalidFrontmatter/unterminated | CON-007 | pass
- contracts | Test_CON_008_RenderRoundTrip | CON-008 | pass
- contracts | Test_CON_002_CON_009_Sections | CON-002, CON-009 | pass
- contracts | Test_CON_009_CON_002_AppendSection | CON-002, CON-009 | pass
- contracts | Test_CON_010_ParseItems | CON-010 | pass
- contracts | Test_CON_011_CON_012_RejectInvalidItems | CON-011, CON-012 | pass
- contracts | Test_CON_011_CON_012_RejectInvalidItems/CON_012_duplicate | CON-011, CON-012 | pass
- contracts | Test_CON_011_CON_012_RejectInvalidItems/CON_011_malformed_heading | CON-011, CON-012 | pass
- contracts | Test_CON_006_CON_010_ParseHandlesInSourceOrder | CON-006, CON-010 | pass
- contracts | Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles | CON-012, CON-015 | pass
- contracts | Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles/CON_015_empty | CON-012, CON-015 | pass
- contracts | Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles/CON_015_uppercase | CON-012, CON-015 | pass
- contracts | Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles/CON_015_underscore | CON-012, CON-015 | pass
- contracts | Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles/CON_015_punctuation | CON-012, CON-015 | pass
- contracts | Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles/CON_012_duplicate | CON-012, CON-015 | pass
- contracts | Test_CON_013_RenderItem | CON-013 | pass
- contracts | Test_CON_003_CanonicalHash | CON-003 | pass
- contracts | Test_CON_003_IDAndTagAffectHash | CON-003 | pass
- contracts | Test_CON_004_PositiveMatches | CON-004 | pass
- contracts | Test_CON_004_Boundaries | CON-004 | pass
- contracts | Test_CON_005_MatchGlob | CON-005 | pass
- contracts | Test_CON_014_ResolveZone | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/protected_first | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/root_spec | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/evals_before_interface | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/interface_before_append | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/append_before_preserve | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/preserve_before_disposable | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/disposable | CON-014 | pass
- contracts | Test_CON_014_ResolveZone/other | CON-014 | pass
- contracts | Test_CON_016_ErrorCode | CON-016 | pass
- contracts | Test_CON_016_ErrorCode/direct | CON-016 | pass
- contracts | Test_CON_016_ErrorCode/wrapped | CON-016 | pass
- contracts | Test_CON_016_ErrorCode/first_in_chain | CON-016 | pass
- contracts | Test_CON_016_ErrorCode/no_cause | CON-016 | pass
- contracts | Test_CON_016_ErrorCode/plain | CON-016 | pass
- contracts | Test_CON_016_ErrorCode/nil | CON-016 | pass
- contracts | Test_CON_017_ParseItemsDuplicateCode | CON-017 | pass
- contracts | Test_CON_017_ParseItemsDuplicateCode/assigned | CON-017 | pass
- contracts | Test_CON_017_ParseItemsDuplicateCode/handle | CON-017 | pass
- contracts | Test_CON_018_ParseItemsMalformedCode | CON-018 | pass
- contracts | Test_CON_018_ParseItemsMalformedCode/unclosed_tag | CON-018 | pass
- contracts | Test_CON_018_ParseItemsMalformedCode/invalid_handle | CON-018 | pass
- contracts | Test_CON_019_ParseDocumentFrontmatterCode | CON-019 | pass
- contracts | Test_CON_019_ParseDocumentFrontmatterCode/missing | CON-019 | pass
- contracts | Test_CON_019_ParseDocumentFrontmatterCode/unterminated | CON-019 | pass
- contracts | Test_CON_019_ParseDocumentFrontmatterCode/invalid_yaml | CON-019 | pass
- contracts | Test_CON_019_ParseDocumentFrontmatterCode/non_mapping | CON-019 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/contracts	[no test files]
?   	github.com/lurkshark/fireweed/packages/contracts/internal/primitives	[no test files]
?   	github.com/lurkshark/fireweed/packages/contracts/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/contracts/evals	0.284s


## Dependent evals
- app | Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON | APP-097, APP-099 | pass
- app | Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/plain-queued | APP-097, APP-099 | pass
- app | Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned | APP-097, APP-099 | pass
- app | Test_APP_098_APP_099_AbandonFailures | APP-098, APP-099 | pass
- app | Test_APP_098_APP_099_AbandonFailures/missing | APP-098, APP-099 | pass
- app | Test_APP_098_APP_099_AbandonFailures/extra-json | APP-098, APP-099 | pass
- app | Test_APP_098_APP_099_AbandonFailures/unsupported-option | APP-098, APP-099 | pass
- app | Test_APP_098_APP_099_AbandonFailures/unknown-selector | APP-098, APP-099 | pass
- app | Test_APP_001_StandaloneBinary | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/outside-project/help | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/outside-project/version | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/outside-project/explain | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/outside-project/playbook | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/invalid-config/help | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/invalid-config/version | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/invalid-config/explain | APP-001 | pass
- app | Test_APP_001_StandaloneBinary/invalid-config/playbook | APP-001 | pass
- app | Test_APP_002_InformationJSON | APP-002 | pass
- app | Test_APP_002_InformationJSON/help | APP-002 | pass
- app | Test_APP_002_InformationJSON/version | APP-002 | pass
- app | Test_APP_002_InformationJSON/explain | APP-002 | pass
- app | Test_APP_002_InformationJSON/playbook | APP-002 | pass
- app | Test_APP_003_InvalidRequests | APP-003 | pass
- app | Test_APP_003_InvalidRequests/unknown-command | APP-003 | pass
- app | Test_APP_003_InvalidRequests/unknown-option | APP-003 | pass
- app | Test_APP_003_InvalidRequests/missing-argument | APP-003 | pass
- app | Test_APP_003_InvalidRequests/extra-argument | APP-003 | pass
- app | Test_APP_003_InvalidRequests/new-change-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/stamp-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/check-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/show-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/start-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/context-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/coverage-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/verify-arity | APP-003 | pass
- app | Test_APP_003_InvalidRequests/APP_013_unknown-explanation | APP-003, APP-013 | pass
- app | Test_APP_003_InvalidRequests/APP_013_unknown-playbook | APP-003, APP-013 | pass
- app | Test_APP_014_UnimplementedCommands | APP-014 | pass
- app | Test_APP_014_UnimplementedCommands/gc | APP-014 | pass
- app | Test_APP_015_FailedJSON | APP-015 | pass
- app | Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents | APP-001 | pass
- app | Test_APP_006_APP_022_StampAliasesAndUnassignedHandles | APP-006, APP-022 | pass
- app | Test_APP_006_APP_022_StampAliasesAndUnassignedHandles/revise | APP-006, APP-022 | pass
- app | Test_APP_006_APP_022_StampAliasesAndUnassignedHandles/0001-revise.md | APP-006, APP-022 | pass
- app | Test_APP_006_APP_022_StampAliasesAndUnassignedHandles/changes/0001-revise.md | APP-006, APP-022 | pass
- app | Test_APP_021_APP_022_StampErrorExits | APP-021, APP-022 | pass
- app | Test_APP_021_APP_022_StampErrorExits/missing-argument | APP-021, APP-022 | pass
- app | Test_APP_021_APP_022_StampErrorExits/extra-argument | APP-021, APP-022 | pass
- app | Test_APP_021_APP_022_StampErrorExits/invalid-config | APP-021, APP-022 | pass
- app | Test_APP_021_APP_022_StampErrorExits/missing-change | APP-021, APP-022 | pass
- app | Test_APP_021_APP_022_StampErrorExits/absent-modified-ID | APP-021, APP-022 | pass
- app | Test_APP_007_APP_024_CheckValidAndAllDiagnostics | APP-007, APP-024 | pass
- app | Test_APP_023_APP_024_CheckUsageAndConfiguration | APP-023, APP-024 | pass
- app | Test_APP_023_APP_024_CheckUsageAndConfiguration/extra-argument | APP-023, APP-024 | pass
- app | Test_APP_023_APP_024_CheckUsageAndConfiguration/invalid-config | APP-023, APP-024 | pass
- app | Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity | APP-008, APP-025 | pass
- app | Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/current-full | APP-008, APP-025 | pass
- app | Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/current-item | APP-008, APP-025 | pass
- app | Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/projected-full | APP-008, APP-025 | pass
- app | Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/projected-item | APP-008, APP-025 | pass
- app | Test_APP_026_ShowNamedErrors | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/unknown-module | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/unknown-item | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/ambiguous-item | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/ambiguous-module | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/missing-module | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/extra-argument | APP-026 | pass
- app | Test_APP_026_ShowNamedErrors/invalid-config | APP-026 | pass
- app | Test_APP_105_APP_108_CheckCommitsRangesAndRecordedAuthorWork | APP-105, APP-108 | pass
- app | Test_APP_105_APP_108_CheckCommitsSquashAndEmptyJSON | APP-105, APP-108 | pass
- app | Test_APP_106_APP_108_CheckCommitsReportsEveryHistoryRule | APP-106, APP-108 | pass
- app | Test_APP_107_APP_108_CheckCommitsFailures | APP-107, APP-108 | pass
- app | Test_APP_107_APP_108_CheckCommitsFailures/extra | APP-107, APP-108 | pass
- app | Test_APP_107_APP_108_CheckCommitsFailures/option | APP-107, APP-108 | pass
- app | Test_APP_107_APP_108_CheckCommitsFailures/bad-range | APP-107, APP-108 | pass
- app | Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules | APP-109, APP-110, APP-112, APP-113 | pass
- app | Test_APP_110_APP_112_APP_113_CIEvalAndBoundaryFailures | APP-110, APP-112, APP-113 | pass
- app | Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration | APP-109, APP-112, APP-113 | pass
- app | Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip | APP-111, APP-112, APP-113 | pass
- app | Test_APP_094_APP_096_CommitsPendingAndArchivedOrder | APP-094, APP-096 | pass
- app | Test_APP_095_APP_096_CommitsFailures | APP-095, APP-096 | pass
- app | Test_APP_095_APP_096_CommitsFailures/missing | APP-095, APP-096 | pass
- app | Test_APP_095_APP_096_CommitsFailures/extra-json | APP-095, APP-096 | pass
- app | Test_APP_095_APP_096_CommitsFailures/unsupported-option | APP-095, APP-096 | pass
- app | Test_APP_095_APP_096_CommitsFailures/unknown-selector | APP-095, APP-096 | pass
- app | Test_APP_037_StaticCauses | APP-037 | pass
- app | Test_APP_037_StaticCauses/configuration | APP-037 | pass
- app | Test_APP_037_StaticCauses/graph | APP-037 | pass
- app | Test_APP_037_StaticCauses/syntax | APP-037 | pass
- app | Test_APP_037_StaticCauses/normative | APP-037 | pass
- app | Test_APP_037_StaticCauses/identity | APP-037 | pass
- app | Test_APP_037_StaticCauses/tag | APP-037 | pass
- app | Test_APP_037_StaticCauses/structure | APP-037 | pass
- app | Test_APP_037_StaticCauses/base | APP-037 | pass
- app | Test_APP_037_StaticCauses/handle | APP-037 | pass
- app | Test_APP_037_StaticCauses/metadata | APP-037 | pass
- app | Test_APP_037_StaticCauses/queue-identity | APP-037 | pass
- app | Test_APP_037_APP_038_ShowStaticCauses | APP-037, APP-038 | pass
- app | Test_APP_037_APP_038_ShowStaticCauses/duplicate-graph-name | APP-037, APP-038 | pass
- app | Test_APP_037_APP_038_ShowStaticCauses/duplicate-fixed-spec-id | APP-037, APP-038 | pass
- app | Test_APP_038_LookupCauses | APP-038 | pass
- app | Test_APP_038_LookupCauses/author-module | APP-038 | pass
- app | Test_APP_038_LookupCauses/stamp-change | APP-038 | pass
- app | Test_APP_038_LookupCauses/show-module | APP-038 | pass
- app | Test_APP_038_LookupCauses/show-item | APP-038 | pass
- app | Test_APP_038_LookupCauses/show-ambiguous-item | APP-038 | pass
- app | Test_APP_038_LookupCauses/coverage-module | APP-038 | pass
- app | Test_APP_038_LookupCauses/coverage-change | APP-038 | pass
- app | Test_APP_039_BlockedNewChangeOutputPath | APP-039 | pass
- app | Test_APP_040_ForwardedCheckCode | APP-040 | pass
- app | Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI | APP-062, APP-063, APP-064, APP-065, APP-066, APP-067 | pass
- app | Test_APP_062_DoctorUsesCloserProject | APP-062 | pass
- app | Test_APP_065_APP_066_APP_067_DoctorFailures | APP-065, APP-066, APP-067 | pass
- app | Test_APP_065_APP_066_APP_067_DoctorFailures/extra | APP-065, APP-066, APP-067 | pass
- app | Test_APP_065_APP_066_APP_067_DoctorFailures/unknown-flag | APP-065, APP-066, APP-067 | pass
- app | Test_APP_065_APP_066_APP_067_DoctorFailures/wrong-command-flag | APP-065, APP-066, APP-067 | pass
- app | Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite | APP-068, APP-069, APP-070, APP-071, APP-072, APP-073 | pass
- app | Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite/no-block | APP-068, APP-069, APP-070, APP-071, APP-072, APP-073 | pass
- app | Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite/missing-file | APP-068, APP-069, APP-070, APP-071, APP-072, APP-073 | pass
- app | Test_APP_136_APP_137_DrillPlainNearestProjectAndInvokingBin | APP-136, APP-137 | pass
- app | Test_APP_138_DrillJSONSuccess | APP-138 | pass
- app | Test_APP_139_DrillArgumentAndModuleRejectionsDoNotInvokeHooks | APP-139 | pass
- app | Test_APP_140_DrillProjectAndConfigurationFailures | APP-140 | pass
- app | Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts | APP-141 | pass
- app | Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain | APP-141 | pass
- app | Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json | APP-141 | pass
- app | Test_APP_141_DrillSetupFailureHasNoArtifacts | APP-141 | pass
- app | Test_APP_142_DrillCleanUsageStopPlainAndJSON | APP-142 | pass
- app | Test_APP_136_DrillUsesConfiguredTimeout | APP-136 | pass
- app | Test_APP_041_APP_042_APP_043_APP_044_APP_045_APP_046_APP_050_APP_053_APP_054_APP_059_GoInit | APP-041, APP-042, APP-043, APP-044, APP-045, APP-046, APP-050, APP-053, APP-054, APP-059 | pass
- app | Test_APP_043_APP_044_APP_047_APP_051_NoneInit | APP-043, APP-044, APP-047, APP-051 | pass
- app | Test_APP_050_APP_052_NoneDefaultApp | APP-050, APP-052 | pass
- app | Test_APP_042_APP_048_APP_049_APP_050_APP_053_APP_054_APP_060_LocalInit | APP-042, APP-048, APP-049, APP-050, APP-053, APP-054, APP-060 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-recipe | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-agent | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-value | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/invalid-agent | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/malformed-option | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/extra-position | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/unknown-recipe | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/malformed-document | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-document | APP-055, APP-056, APP-061 | pass
- app | Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/invalid-metadata-type | APP-055, APP-056, APP-061 | pass
- app | Test_APP_057_APP_058_InitConflicts | APP-057, APP-058 | pass
- app | Test_APP_057_APP_058_InitConflicts/config | APP-057, APP-058 | pass
- app | Test_APP_057_APP_058_InitConflicts/recipe | APP-057, APP-058 | pass
- app | Test_APP_057_APP_058_InitConflicts/module | APP-057, APP-058 | pass
- app | Test_APP_061_InitJSONSuccess | APP-061 | pass
- app | Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection | APP-009, APP-027, APP-028 | pass
- app | Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning | APP-009, APP-027, APP-028, APP-040 | pass
- app | Test_APP_004_HeldModulesUseApprovedProjection | APP-004 | pass
- app | Test_APP_011_APP_032_APP_033_ContextCompactFullAndErrors | APP-011, APP-032, APP-033 | pass
- app | Test_APP_012_APP_034_APP_035_APP_036_CoverageCurrentProjectionAndErrors | APP-012, APP-034, APP-035, APP-036 | pass
- app | Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks | APP-010, APP-029, APP-030, APP-031 | pass
- app | Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits | APP-010, APP-029, APP-030, APP-031 | pass
- app | Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit | APP-010, APP-029, APP-031 | pass
- app | Test_APP_114_APP_115_APP_117_ListQueueOrderAndFields | APP-114, APP-115, APP-117 | pass
- app | Test_APP_114_APP_117_ListNearestRootAndEmptyQueue | APP-114, APP-117 | pass
- app | Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors | APP-116, APP-117 | pass
- app | Test_APP_100_APP_101_APP_104_NewModuleFromNearestProjectTemplates | APP-100, APP-101, APP-104 | pass
- app | Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents | APP-102, APP-103, APP-104 | pass
- app | Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents/existing-name | APP-102, APP-103, APP-104 | pass
- app | Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents/existing-prefix | APP-102, APP-103, APP-104 | pass
- app | Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents/occupied-directory | APP-102, APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/missing-name | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/missing-prefix | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/invalid-name | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/invalid-prefix | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/repeated-prefix | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/extra-argument | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/unknown-option | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/missing-template | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/malformed-template | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/invalid-configuration | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/invalid-existing-graph | APP-103, APP-104 | pass
- app | Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/no-configured-destination | APP-103, APP-104 | pass
- app | Test_APP_101_APP_104_NewModuleEmptyGoModule | APP-101, APP-104 | pass
- app | Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange | APP-074, APP-076, APP-077 | pass
- app | Test_APP_075_NextAllowUnmergedAndUsage | APP-075 | pass
- app | Test_APP_076_NextPlainEmptyQueue | APP-076 | pass
- app | Test_APP_077_NextExitClasses | APP-077 | pass
- app | Test_APP_078_NextWrappedCauseAndAction | APP-078 | pass
- app | Test_APP_079_NextJSONSuccessAndFailure | APP-079 | pass
- app | Test_APP_127_APP_128_APP_129_PRBodyExplicitHistoricalEndpoint | APP-127, APP-128, APP-129 | pass
- app | Test_APP_128_APP_129_PRBodyValidEmptyRange | APP-128, APP-129 | pass
- app | Test_APP_130_APP_131_PRBodyFailures | APP-130, APP-131 | pass
- app | Test_APP_130_APP_131_PRBodyFailures/missing-range | APP-130, APP-131 | pass
- app | Test_APP_130_APP_131_PRBodyFailures/extra-range | APP-130, APP-131 | pass
- app | Test_APP_130_APP_131_PRBodyFailures/invalid-option | APP-130, APP-131 | pass
- app | Test_APP_130_APP_131_PRBodyFailures/absent-project | APP-130, APP-131 | pass
- app | Test_APP_130_APP_131_PRBodyFailures/bad-endpoint-config | APP-130, APP-131 | pass
- app | Test_APP_130_APP_131_PRBodyFailures/unavailable-range | APP-130, APP-131 | pass
- app | Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor | APP-004, APP-016, APP-017, APP-018 | pass
- app | Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor/plain | APP-004, APP-016, APP-017, APP-018 | pass
- app | Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor/JSON | APP-004, APP-016, APP-017, APP-018 | pass
- app | Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor/nearest-ancestor | APP-004, APP-016, APP-017, APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/missing/plain | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/missing/JSON | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/malformed/plain | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/malformed/JSON | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/invalid-graph/plain | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/invalid-graph/JSON | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/invalid-document/plain | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/invalid-document/JSON | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/invalid-metadata/plain | APP-018 | pass
- app | Test_APP_018_ModulesConfigurationErrors/invalid-metadata/JSON | APP-018 | pass
- app | Test_APP_005_APP_020_NewChangePlainAndJSON | APP-005, APP-020 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/missing-argument/plain | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/missing-argument/JSON | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/extra-argument/plain | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/extra-argument/JSON | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/invalid-module/plain | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/invalid-module/JSON | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain | APP-019 | pass
- app | Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON | APP-019 | pass
- app | Test_APP_020_NewChangeJSONFailureEnvelope | APP-020 | pass
- app | Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork | APP-091, APP-092, APP-093 | pass
- app | Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork/plain-path | APP-091, APP-092, APP-093 | pass
- app | Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork/json-filename | APP-091, APP-092, APP-093 | pass
- app | Test_APP_092_APP_093_ReleaseForeignClaimKeepsOwner | APP-092, APP-093 | pass
- app | Test_APP_092_APP_093_ReleaseUsageAndConfiguration | APP-092, APP-093 | pass
- app | Test_APP_092_APP_093_ReleaseUsageAndConfiguration/missing-argument | APP-092, APP-093 | pass
- app | Test_APP_092_APP_093_ReleaseUsageAndConfiguration/extra-argument-json | APP-092, APP-093 | pass
- app | Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation | APP-080, APP-081, APP-083, APP-084 | pass
- app | Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON | APP-081, APP-084, APP-086 | pass
- app | Test_APP_083_RunPlainMultiple | APP-083 | pass
- app | Test_APP_086_RunJSONLimited | APP-086 | pass
- app | Test_APP_082_RunAllowUnmerged | APP-082 | pass
- app | Test_APP_081_APP_084_RunUsageAndConfiguration | APP-081, APP-084 | pass
- app | Test_APP_084_APP_085_APP_086_RunWrappedFailure | APP-084, APP-085, APP-086 | pass
- app | Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes | APP-084, APP-089, APP-090 | pass
- app | Test_APP_090_APP_084_RunForPlainStop | APP-084, APP-090 | pass
- app | Test_APP_084_RunForPriorBlockStillFails | APP-084 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation | APP-089 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation/missing | APP-089 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation/malformed | APP-089 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation/zero | APP-089 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation/negative | APP-089 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation/repeated | APP-089 | pass
- app | Test_APP_089_RunForInvalidOptionsBeforeMutation/next | APP-089 | pass
- app | Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext | APP-087 | pass
- app | Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext/plain-second-ref | APP-087 | pass
- app | Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext/json-first-ref | APP-087 | pass
- app | Test_APP_087_MissingSavedAttemptRetainsLifecycleFailureShape | APP-087 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/missing | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/noninteger | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/zero | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/negative | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/repeated | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/context | APP-088 | pass
- app | Test_APP_088_FromAttemptUsageRejectsBeforeMutation/help | APP-088 | pass
- app | Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter | APP-132, APP-133 | pass
- app | Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter/selected-ID | APP-132, APP-133 | pass
- app | Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter/module-wide | APP-132, APP-133 | pass
- app | Test_APP_132_APP_133_ShowHistoryNearestRootAndEmpty | APP-132, APP-133 | pass
- app | Test_APP_134_ShowHistoryOptionMisuse | APP-134 | pass
- app | Test_APP_134_ShowHistoryOptionMisuse/other-command/plain | APP-134 | pass
- app | Test_APP_134_ShowHistoryOptionMisuse/other-command/JSON | APP-134 | pass
- app | Test_APP_134_ShowHistoryOptionMisuse/projected-combination/plain | APP-134 | pass
- app | Test_APP_134_ShowHistoryOptionMisuse/projected-combination/JSON | APP-134 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/missing-config/plain | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/missing-config/JSON | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON | APP-135 | pass
- app | Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport | APP-118, APP-119, APP-120, APP-121, APP-122, APP-123, APP-124, APP-125 | pass
- app | Test_APP_118_APP_119_APP_122_StatusNearestRootEmptyContainers | APP-118, APP-119, APP-122 | pass
- app | Test_APP_126_StatusUsageConfigurationAndOperationalFailures | APP-126 | pass
- app | Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume | APP-076, APP-077, APP-079 | pass
- app | Test_APP_076_APP_077_NextPlainUsageStop | APP-076, APP-077 | pass
- app | Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume | APP-083, APP-084, APP-086 | pass
- app | Test_APP_083_APP_084_RunPlainSharedUsageStop | APP-083, APP-084 | pass
- app | Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop | APP-084, APP-086 | pass
- app | Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation | APP-077, APP-079, APP-084, APP-086 | pass
- app | Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/next | APP-077, APP-079, APP-084, APP-086 | pass
- app | Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/run | APP-077, APP-079, APP-084, APP-086 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/app/evals	251.659s

- assets | Test_AST_001_Version | AST-001 | pass
- assets | Test_AST_008_Help | AST-008 | pass
- assets | Test_AST_002_Playbooks | AST-002 | pass
- assets | Test_AST_009_AST_010_AuditRules | AST-009, AST-010 | pass
- assets | Test_AST_028_AuditPassingExample | AST-028 | pass
- assets | Test_AST_011_AST_012_AuthorRules | AST-011, AST-012 | pass
- assets | Test_AST_003_Explanations | AST-003 | pass
- assets | Test_AST_003_Explanations/C01 | AST-003 | pass
- assets | Test_AST_003_Explanations/C04 | AST-003 | pass
- assets | Test_AST_003_Explanations/C05 | AST-003 | pass
- assets | Test_AST_003_Explanations/C06 | AST-003 | pass
- assets | Test_AST_003_Explanations/C07 | AST-003 | pass
- assets | Test_AST_003_Explanations/C08 | AST-003 | pass
- assets | Test_AST_003_Explanations/C09 | AST-003 | pass
- assets | Test_AST_003_Explanations/X01 | AST-003 | pass
- assets | Test_AST_003_Explanations/X02 | AST-003 | pass
- assets | Test_AST_003_Explanations/X03 | AST-003 | pass
- assets | Test_AST_003_Explanations/X04 | AST-003 | pass
- assets | Test_AST_003_Explanations/U01 | AST-003 | pass
- assets | Test_AST_003_Explanations/U02 | AST-003 | pass
- assets | Test_AST_003_Explanations/P2-01 | AST-003 | pass
- assets | Test_AST_024_StaticExplanations | AST-024 | pass
- assets | Test_AST_024_StaticExplanations/C02 | AST-024 | pass
- assets | Test_AST_024_StaticExplanations/C03 | AST-024 | pass
- assets | Test_AST_024_StaticExplanations/C12 | AST-024 | pass
- assets | Test_AST_024_StaticExplanations/C13 | AST-024 | pass
- assets | Test_AST_025_EvaluationExplanations | AST-025 | pass
- assets | Test_AST_025_EvaluationExplanations/C10 | AST-025 | pass
- assets | Test_AST_025_EvaluationExplanations/C11 | AST-025 | pass
- assets | Test_AST_025_EvaluationExplanations/E01 | AST-025 | pass
- assets | Test_AST_026_LifecycleExplanations | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P01 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P02 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P03 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P04 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P05 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P06 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P07 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/P09 | AST-026 | pass
- assets | Test_AST_026_LifecycleExplanations/Z01 | AST-026 | pass
- assets | Test_AST_027_DoctorExplanations | AST-027 | pass
- assets | Test_AST_027_DoctorExplanations/C14 | AST-027 | pass
- assets | Test_AST_027_DoctorExplanations/C15 | AST-027 | pass
- assets | Test_AST_027_DoctorExplanations/C16 | AST-027 | pass
- assets | Test_AST_045_DecisionsAndNotesGuidance | AST-045 | pass
- assets | Test_AST_045_DecisionsAndNotesGuidance/evals | AST-045 | pass
- assets | Test_AST_045_DecisionsAndNotesGuidance/impl | AST-045 | pass
- assets | Test_AST_045_DecisionsAndNotesGuidance/regenerate | AST-045 | pass
- assets | Test_AST_046_DecisionAndNoteExamples | AST-046 | pass
- assets | Test_AST_046_DecisionAndNoteExamples/evals | AST-046 | pass
- assets | Test_AST_046_DecisionAndNoteExamples/impl | AST-046 | pass
- assets | Test_AST_046_DecisionAndNoteExamples/regenerate | AST-046 | pass
- assets | Test_AST_029_EvalFixEvalsGuidance | AST-029 | pass
- assets | Test_AST_030_EvalFixAuditGuidance | AST-030 | pass
- assets | Test_AST_031_EvalFixImplGuidance | AST-031 | pass
- assets | Test_AST_004_AST_013_AST_014_AST_015_GoRecipe | AST-004, AST-013, AST-014, AST-015 | pass
- assets | Test_AST_005_EvalHookResults | AST-005 | pass
- assets | Test_AST_016_UnfinishedAndCompileFailure | AST-016 | pass
- assets | Test_AST_005_EvalOnlyRequestedModule | AST-005 | pass
- assets | Test_AST_006_AST_017_AST_019_BoundaryGraph | AST-006, AST-017, AST-019 | pass
- assets | Test_AST_006_AST_018_AST_019_EvalInterfaceZones | AST-006, AST-018, AST-019 | pass
- assets | Test_AST_006_AST_017_AST_018_AST_019_UsesModuleDirsAndWorkspacePath | AST-006, AST-017, AST-018, AST-019 | pass
- assets | Test_AST_007_AST_020_AST_021_AST_022_AST_023_TemplatesAndCanaries | AST-007, AST-020, AST-021, AST-022, AST-023 | pass
- assets | Test_AST_007_AST_020_AST_021_AST_022_AST_023_TemplatesAndCanaries/undeclared | AST-007, AST-020, AST-021, AST-022, AST-023 | pass
- assets | Test_AST_007_AST_020_AST_021_AST_022_AST_023_TemplatesAndCanaries/internal_eval | AST-007, AST-020, AST-021, AST-022, AST-023 | pass
- assets | Test_AST_007_AST_020_AST_021_AST_022_AST_023_TemplatesAndCanaries/failing_eval | AST-007, AST-020, AST-021, AST-022, AST-023 | pass
- assets | Test_AST_007_AST_020_AST_021_AST_022_AST_023_TemplatesAndCanaries/missing_id | AST-007, AST-020, AST-021, AST-022, AST-023 | pass
- assets | Test_AST_034_H01Explanation | AST-034 | pass
- assets | Test_AST_035_H02Explanation | AST-035 | pass
- assets | Test_AST_036_H03Explanation | AST-036 | pass
- assets | Test_AST_037_H04Explanation | AST-037 | pass
- assets | Test_AST_033_MixedEvalFixExplanation | AST-033 | pass
- assets | Test_AST_042_B04Explanation | AST-042 | pass
- assets | Test_AST_032_RegenerateWarningExplanation | AST-032 | pass
- assets | Test_AST_038_B01Explanation | AST-038 | pass
- assets | Test_AST_039_B02Explanation | AST-039 | pass
- assets | Test_AST_040_B03Explanation | AST-040 | pass
- assets | Test_AST_041_P08Explanation | AST-041 | pass
- assets | Test_AST_043_EvalsStubHandoff | AST-043 | pass
- assets | Test_AST_044_GoRecipeStubHandoff | AST-044 | pass
- assets | Test_AST_044_GoRecipeStubHandoff/Instructions | AST-044 | pass
- assets | Test_AST_044_GoRecipeStubHandoff/bundled_go.md | AST-044 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/assets	[no test files]
?   	github.com/lurkshark/fireweed/packages/assets/internal/content	[no test files]
?   	github.com/lurkshark/fireweed/packages/assets/internal/recipe	[no test files]
?   	github.com/lurkshark/fireweed/packages/assets/internal/recipe/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/assets/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/assets/evals	11.055s

- checks | Test_CHK_005_GlobalAndModuleBoundaryHooks | CHK-005 | pass
- checks | Test_CHK_041_ExactBaselineOnly | CHK-041 | pass
- checks | Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns | CHK-039, CHK-040 | pass
- checks | Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional | CHK-044, CHK-045 | pass
- checks | Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut | CHK-041, CHK-042 | pass
- checks | Test_CHK_043_LaunchCancellationAndTimeoutStayErrors | CHK-043 | pass
- checks | Test_CHK_049_StaticDiagnosticsStopProbes | CHK-049 | pass
- checks | Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe | CHK-050, CHK-051, CHK-052, CHK-053, CHK-055, CHK-056, CHK-057, CHK-058, CHK-059, CHK-060, CHK-061, CHK-067 | pass
- checks | Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries | CHK-053 | pass
- checks | Test_CHK_051_CHK_052_SetupFailureStopsAndCleans | CHK-051, CHK-052 | pass
- checks | Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup | CHK-051 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps | CHK-054 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps/missing_template | CHK-054 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps/malformed_template | CHK-054 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps/missing_baseline | CHK-054 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps/missing_internal_eval | CHK-054 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps/wrong_version | CHK-054 | pass
- checks | Test_CHK_054_MissingTemplatesAndCanaryMaps/malformed_canary | CHK-054 | pass
- checks | Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery | CHK-055, CHK-056, CHK-057, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific | CHK-055, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/missing_baseline_coverage | CHK-055, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/failing_eval_compile_error_without_report | CHK-055, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/failing_eval_with_unrelated_failure | CHK-055, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/missing_id_with_unrelated_failure | CHK-055, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/missing_id_without_report | CHK-055, CHK-058, CHK-059 | pass
- checks | Test_CHK_055_BaselineAndRecoveryFailures | CHK-055 | pass
- checks | Test_CHK_055_BaselineAndRecoveryFailures/baseline_eval_failure | CHK-055 | pass
- checks | Test_CHK_055_BaselineAndRecoveryFailures/baseline_boundary_failure | CHK-055 | pass
- checks | Test_CHK_055_BaselineAndRecoveryFailures/recovery_eval_failure | CHK-055 | pass
- checks | Test_CHK_055_BaselineAndRecoveryFailures/recovery_coverage_failure | CHK-055 | pass
- checks | Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule | CHK-056, CHK-057 | pass
- checks | Test_CHK_057_InternalEvalRequiresInterfaceViolation | CHK-057 | pass
- checks | Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook | CHK-055, CHK-056, CHK-057 | pass
- checks | Test_CHK_055_MissingRequiredHookBesideOverride | CHK-055 | pass
- checks | Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook | CHK-055 | pass
- checks | Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings | CHK-060, CHK-062 | pass
- checks | Test_CHK_060_OverridesMustDetectTheirCanaries | CHK-060 | pass
- checks | Test_CHK_060_OverridesMustDetectTheirCanaries/eval | CHK-060 | pass
- checks | Test_CHK_060_OverridesMustDetectTheirCanaries/boundary | CHK-060 | pass
- checks | Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes | CHK-063, CHK-064, CHK-065, CHK-066, CHK-067 | pass
- checks | Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes/wrong_reply | CHK-063, CHK-064, CHK-065, CHK-066, CHK-067 | pass
- checks | Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes/nonzero_with_reply | CHK-063, CHK-064, CHK-065, CHK-066, CHK-067 | pass
- checks | Test_CHK_063_CHK_064_CHK_065_CHK_066_CHK_067_AgentProbes/ordinary_failure | CHK-063, CHK-064, CHK-065, CHK-066, CHK-067 | pass
- checks | Test_CHK_065_UnsetTimeoutBoundsAuditProbe | CHK-065 | pass
- checks | Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase | CHK-068, CHK-069 | pass
- checks | Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs | CHK-070 | pass
- checks | Test_CHK_070_HistoryBranchingRangeTopologicalOrder | CHK-070 | pass
- checks | Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase | CHK-070 | pass
- checks | Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl | CHK-071, CHK-072 | pass
- checks | Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface | CHK-072, CHK-073, CHK-077 | pass
- checks | Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface | CHK-071, CHK-073 | pass
- checks | Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename | CHK-074 | pass
- checks | Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename/modify | CHK-074 | pass
- checks | Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename/delete | CHK-074 | pass
- checks | Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename/rename | CHK-074 | pass
- checks | Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange | CHK-074 | pass
- checks | Test_CHK_075_HistoryArchiveModifyAndDeletionControl | CHK-075 | pass
- checks | Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04 | CHK-076 | pass
- checks | Test_CHK_077_HistoryUsesRecordedBeforeAfterZones | CHK-077 | pass
- checks | Test_CHK_077_HistoryChecksBothRecordedZoneMaps | CHK-077 | pass
- checks | Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments | CHK-078 | pass
- checks | Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/invalid_range | CHK-078 | pass
- checks | Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/missing_config | CHK-078 | pass
- checks | Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/malformed_config | CHK-078 | pass
- checks | Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/missing_SPEC | CHK-078 | pass
- checks | Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/malformed_change | CHK-078 | pass
- checks | Test_CHK_003_CHK_032_CHK_033_CHK_034_CHK_035_GlobalShellHookQuotesAndOnePassExpansion | CHK-003, CHK-032, CHK-033, CHK-034, CHK-035 | pass
- checks | Test_CHK_003_CHK_031_CHK_032_CHK_033_CHK_034_CHK_035_ModuleOverrideArgvAndOutput | CHK-003, CHK-031, CHK-032, CHK-033, CHK-034, CHK-035 | pass
- checks | Test_CHK_038_ExitAndReportAgreement | CHK-038 | pass
- checks | Test_CHK_080_EvaluateRetainsParsedFailureDetails | CHK-080 | pass
- checks | Test_CHK_080_EvaluateRetainsParsedFailureDetails/nonzero_failure | CHK-080 | pass
- checks | Test_CHK_080_EvaluateRetainsParsedFailureDetails/zero_passing | CHK-080 | pass
- checks | Test_CHK_036_CHK_037_MissingReportAndLaunchFailure | CHK-036, CHK-037 | pass
- checks | Test_CHK_036_TimeoutFails | CHK-036 | pass
- checks | Test_CHK_046_EvaluateHookSetupCauses | CHK-046 | pass
- checks | Test_CHK_046_EvaluateHookSetupCauses/missing_hook | CHK-046 | pass
- checks | Test_CHK_046_EvaluateHookSetupCauses/invalid_hook_command | CHK-046 | pass
- checks | Test_CHK_047_EvaluateUnknownModuleCause | CHK-047 | pass
- checks | Test_CHK_048_EvaluateModuleItemCauses | CHK-048 | pass
- checks | Test_CHK_048_EvaluateModuleItemCauses/malformed_item | CHK-048 | pass
- checks | Test_CHK_048_EvaluateModuleItemCauses/duplicate_item | CHK-048 | pass
- checks | Test_CHK_028_EvaluatePreservesJUnitDiagnostic | CHK-028 | pass
- checks | Test_CHK_003_CHK_032_ArgvPlaceholdersAreOnePass | CHK-003, CHK-032 | pass
- checks | Test_CHK_001_CHK_029_CHK_030_NestedJUnitIDsResultsAndTime | CHK-001, CHK-029, CHK-030 | pass
- checks | Test_CHK_001_CHK_029_CHK_030_SingleSuiteSkipped | CHK-001, CHK-029, CHK-030 | pass
- checks | Test_CHK_079_ParseJUnitRetainsFailureAndErrorDetails | CHK-079 | pass
- checks | Test_CHK_002_SkippedJUnitCaseDoesNotCover | CHK-002 | pass
- checks | Test_CHK_027_RejectMalformedXML | CHK-027 | pass
- checks | Test_CHK_028_RejectMissingAndUnknownIDs | CHK-028 | pass
- checks | Test_CHK_002_CoverageRequiresCurrentNonSkippedCases | CHK-002 | pass
- checks | Test_CHK_002_CoverageExemptsReviewAndQuarantine | CHK-002 | pass
- checks | Test_CHK_025_FailFirstRequiresEachRequestedFailure | CHK-025 | pass
- checks | Test_CHK_025_FailFirstAlreadySatisfiedExemption | CHK-025 | pass
- checks | Test_CHK_026_PassRejectsNonQuarantinedFailures | CHK-026 | pass
- checks | Test_CHK_004_AuditOnlyCurrentChange | CHK-004 | pass
- checks | Test_CHK_006_CHK_007_CHK_008_CHK_009_RejectProtectedSpecForeignModuleAndChangeEveryPhase | CHK-006, CHK-007, CHK-008, CHK-009 | pass
- checks | Test_CHK_011_CHK_012_CHK_013_CHK_014_CHK_015_CHK_021_EvalsAllowedZonesAndNewStubs | CHK-011, CHK-012, CHK-013, CHK-014, CHK-015, CHK-021 | pass
- checks | Test_CHK_016_CHK_017_CHK_018_CHK_019_CHK_020_CHK_021_ImplFreezesEvalsAndInterfaces | CHK-016, CHK-017, CHK-018, CHK-019, CHK-020, CHK-021 | pass
- checks | Test_CHK_022_CHK_023_ImplAppendOnlyNewVersusExisting | CHK-022, CHK-023 | pass
- checks | Test_CHK_024_RenamesCheckBothPaths | CHK-024 | pass
- checks | Test_CHK_010_WorkerWritableOnlyEvalsAndImpl | CHK-010 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/checks	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/internal/engine	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/checks/evals	229.121s

- gitstore | Test_GIT_001_InspectWorktreesAndDetachedHead | GIT-001 | pass
- gitstore | Test_GIT_002_BaseRefPreferenceDefaultsAndMissing | GIT-002 | pass
- gitstore | Test_GIT_003_GIT_006_GIT_007_ChangesReadAtAndSelectedCommit | GIT-003, GIT-006, GIT-007 | pass
- gitstore | Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS | GIT-004, GIT-008, GIT-009, GIT-010 | pass
- gitstore | Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS/StaleUpdatePreservesNewerRef | GIT-004, GIT-008, GIT-009, GIT-010 | pass
- gitstore | Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel | GIT-005 | pass
- gitstore | Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel/TimeoutTerminatesGroup | GIT-005 | pass
- gitstore | Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel/ContextCancellationTerminatesGroup | GIT-005 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/gitstore	[no test files]
?   	github.com/lurkshark/fireweed/packages/gitstore/internal/local	[no test files]
?   	github.com/lurkshark/fireweed/packages/gitstore/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/gitstore/evals	4.456s

- project | Test_PRJ_057_StampInvalidBasesHasCode | PRJ-057 | pass
- project | Test_PRJ_058_MissingPendingChangeHasSubject | PRJ-058 | pass
- project | Test_PRJ_058_MissingPendingChangeHasSubject/Stamp | PRJ-058 | pass
- project | Test_PRJ_058_MissingPendingChangeHasSubject/Project | PRJ-058 | pass
- project | Test_PRJ_059_LoadPreservesModuleParseCode | PRJ-059 | pass
- project | Test_PRJ_059_LoadPreservesModuleParseCode/malformed_frontmatter | PRJ-059 | pass
- project | Test_PRJ_059_LoadPreservesModuleParseCode/malformed_spec_heading | PRJ-059 | pass
- project | Test_PRJ_059_LoadPreservesModuleParseCode/duplicate_spec_id | PRJ-059 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_fixes | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_added | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_modified | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_removed | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_only | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/fixes_only | PRJ-046 | pass
- project | Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/added_only | PRJ-046 | pass
- project | Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults | PRJ-001, PRJ-017 | pass
- project | Test_PRJ_018_LoadInvalidDocuments | PRJ-018 | pass
- project | Test_PRJ_018_LoadInvalidDocuments/missing_config | PRJ-018 | pass
- project | Test_PRJ_018_LoadInvalidDocuments/malformed_config | PRJ-018 | pass
- project | Test_PRJ_018_LoadInvalidDocuments/missing_module_document | PRJ-018 | pass
- project | Test_PRJ_018_LoadInvalidDocuments/malformed_module_document | PRJ-018 | pass
- project | Test_PRJ_002_Graph | PRJ-002 | pass
- project | Test_PRJ_002_Graph/PRJ-002_acyclic | PRJ-002 | pass
- project | Test_PRJ_002_Graph/PRJ-019_empty_name | PRJ-002, PRJ-019 | pass
- project | Test_PRJ_002_Graph/PRJ-020_empty_prefix | PRJ-002, PRJ-020 | pass
- project | Test_PRJ_002_Graph/PRJ-019_duplicate_name | PRJ-002, PRJ-019 | pass
- project | Test_PRJ_002_Graph/PRJ-020_duplicate_prefix | PRJ-002, PRJ-020 | pass
- project | Test_PRJ_002_Graph/PRJ-021_missing_dependency | PRJ-002, PRJ-021 | pass
- project | Test_PRJ_002_Graph/PRJ-021_self_dependency | PRJ-002, PRJ-021 | pass
- project | Test_PRJ_002_Graph/PRJ-022_cycle | PRJ-002, PRJ-022 | pass
- project | Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering | PRJ-003, PRJ-023 | pass
- project | Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite | PRJ-024 | pass
- project | Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite/pending | PRJ-024 | pass
- project | Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite/archived | PRJ-024 | pass
- project | Test_PRJ_025_PRJ_026_NewChangeInvalidArguments | PRJ-025, PRJ-026 | pass
- project | Test_PRJ_025_PRJ_026_NewChangeInvalidArguments/PRJ-025_unknown_module | PRJ-025, PRJ-026 | pass
- project | Test_PRJ_025_PRJ_026_NewChangeInvalidArguments/PRJ-026_uppercase_slug | PRJ-025, PRJ-026 | pass
- project | Test_PRJ_025_PRJ_026_NewChangeInvalidArguments/PRJ-026_underscore_slug | PRJ-025, PRJ-026 | pass
- project | Test_PRJ_025_PRJ_026_NewChangeInvalidArguments/PRJ-026_slash_slug | PRJ-025, PRJ-026 | pass
- project | Test_PRJ_025_PRJ_026_NewChangeInvalidArguments/PRJ-026_empty_slug | PRJ-025, PRJ-026 | pass
- project | Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive | PRJ-027, PRJ-028 | pass
- project | Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments | PRJ-029, PRJ-030 | pass
- project | Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments/PRJ-029_malformed | PRJ-029, PRJ-030 | pass
- project | Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments/PRJ-030_unknown_module | PRJ-029, PRJ-030 | pass
- project | Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug | PRJ-004, PRJ-031, PRJ-032 | pass
- project | Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles | PRJ-005, PRJ-032, PRJ-036 | pass
- project | Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections | PRJ-033, PRJ-034, PRJ-035 | pass
- project | Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-034_identity_name | PRJ-033, PRJ-034, PRJ-035 | pass
- project | Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-034_identity_prefix | PRJ-033, PRJ-034, PRJ-035 | pass
- project | Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-034_unknown_field | PRJ-033, PRJ-034, PRJ-035 | pass
- project | Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-035_invalid_graph | PRJ-033, PRJ-034, PRJ-035 | pass
- project | Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation | PRJ-006, PRJ-038 | pass
- project | Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID | PRJ-037, PRJ-039 | pass
- project | Test_PRJ_007_CheckValidAndDiagnostics | PRJ-007 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics | PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-008_missing_motivation | PRJ-008, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-009_nonnormative_module | PRJ-009, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-009_nonnormative_change | PRJ-009, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-010_invalid_review_tag | PRJ-010, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-010_invalid_quarantine_tag | PRJ-010, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-011_duplicate_module_id | PRJ-011, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_module_id | PRJ-012, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_change_id | PRJ-012, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-013_duplicate_edit_identity | PRJ-013, PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-014_undeclared_handle | PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-015_missing_base | PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-015_stale_base | PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-016_absent_modified | PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-016_absent_removed | PRJ-014, PRJ-015, PRJ-016 | pass
- project | Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments | PRJ-047 | pass
- project | Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones | PRJ-048 | pass
- project | Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs | PRJ-049 | pass
- project | Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs | PRJ-050 | pass
- project | Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/invalid_revision | PRJ-050 | pass
- project | Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/missing_recorded_config | PRJ-050 | pass
- project | Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/malformed_recorded_config | PRJ-050 | pass
- project | Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/missing_recorded_SPEC | PRJ-050 | pass
- project | Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications | PRJ-051 | pass
- project | Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate | PRJ-053 | pass
- project | Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata | PRJ-052 | pass
- project | Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule | PRJ-054 | pass
- project | Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec | PRJ-055 | pass
- project | Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors | PRJ-056 | pass
- project | Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/config_missing | PRJ-056 | pass
- project | Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/config_malformed | PRJ-056 | pass
- project | Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/archive_document | PRJ-056 | pass
- project | Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/archive_added | PRJ-056 | pass
- project | Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/archive_modified | PRJ-056 | pass
- project | Test_PRJ_040_CheckConfigurationCauses | PRJ-040 | pass
- project | Test_PRJ_040_CheckConfigurationCauses/missing | PRJ-040 | pass
- project | Test_PRJ_040_CheckConfigurationCauses/malformed | PRJ-040 | pass
- project | Test_PRJ_041_CheckGraphCauses | PRJ-041 | pass
- project | Test_PRJ_041_CheckGraphCauses/current | PRJ-041 | pass
- project | Test_PRJ_041_CheckGraphCauses/projected | PRJ-041 | pass
- project | Test_PRJ_042_CheckDocumentAndItemCauses | PRJ-042 | pass
- project | Test_PRJ_042_CheckDocumentAndItemCauses/missing_module | PRJ-042 | pass
- project | Test_PRJ_042_CheckDocumentAndItemCauses/malformed_module | PRJ-042 | pass
- project | Test_PRJ_042_CheckDocumentAndItemCauses/malformed_change | PRJ-042 | pass
- project | Test_PRJ_042_CheckDocumentAndItemCauses/malformed_item | PRJ-042 | pass
- project | Test_PRJ_043_CheckMetadataAndSetCauses | PRJ-043 | pass
- project | Test_PRJ_043_CheckMetadataAndSetCauses/metadata_type | PRJ-043 | pass
- project | Test_PRJ_043_CheckMetadataAndSetCauses/zone_type | PRJ-043 | pass
- project | Test_PRJ_043_CheckMetadataAndSetCauses/identity_set_key | PRJ-043 | pass
- project | Test_PRJ_043_CheckMetadataAndSetCauses/unknown_set_key | PRJ-043 | pass
- project | Test_PRJ_043_CheckMetadataAndSetCauses/set_value_type | PRJ-043 | pass
- project | Test_PRJ_044_CheckQueueIdentityCause | PRJ-044 | pass
- project | Test_PRJ_045_CheckPendingModuleTargetCause | PRJ-045 | pass
- project | Test_PRJ_013_CheckDuplicateWithinEditSection | PRJ-013 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	1.792s

- runner | Test_RUN_041_UnapprovedChangeRequiresApproval | RUN-041 | pass
- runner | Test_RUN_042_OtherNextErrorsDoNotRequireApproval | RUN-042 | pass
- runner | Test_RUN_042_OtherNextErrorsDoNotRequireApproval/start_failure | RUN-042 | pass
- runner | Test_RUN_042_OtherNextErrorsDoNotRequireApproval/agent_failure | RUN-042 | pass
- runner | Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit | RUN-013, RUN-045, RUN-046 | pass
- runner | Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration | RUN-013 | pass
- runner | Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry | RUN-048 | pass
- runner | Test_RUN_046_AuditCommitDoesNotEnterHeldGit | RUN-046 | pass
- runner | Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported | RUN-047, RUN-048 | pass
- runner | Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces | RUN-045, RUN-049 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically | RUN-047 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/change_outside_Audit | RUN-047 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/projected_spec | RUN-047 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/eval | RUN-047 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/added_file | RUN-047 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/deleted_file | RUN-047 | pass
- runner | Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/committed_invalid_edit | RUN-047 | pass
- runner | Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops | RUN-048, RUN-049 | pass
- runner | Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/agent_failure | RUN-048, RUN-049 | pass
- runner | Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/usage_stop | RUN-048, RUN-049 | pass
- runner | Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/time_stop | RUN-048, RUN-049 | pass
- runner | Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle | RUN-027, RUN-028 | pass
- runner | Test_RUN_027_RUN_028_RecoversCommitBeforeInvocationFailure | RUN-027, RUN-028 | pass
- runner | Test_RUN_050_DrillPreflightRejectsWithoutEffects | RUN-050 | pass
- runner | Test_RUN_050_DrillPreflightRejectsWithoutEffects/unknown | RUN-050 | pass
- runner | Test_RUN_050_DrillPreflightRejectsWithoutEffects/dirty | RUN-050 | pass
- runner | Test_RUN_050_DrillPreflightRejectsWithoutEffects/claim | RUN-050 | pass
- runner | Test_RUN_051_SetupRunsDetachedAtHeadAndFailureStops | RUN-051 | pass
- runner | Test_RUN_052_RUN_053_RUN_054_RestrictedInputsAndPrompt | RUN-052, RUN-053, RUN-054 | pass
- runner | Test_RUN_052_RUN_053_ProjectProtectedDisposableIsRetainedInputAndVerifierFile | RUN-052, RUN-053 | pass
- runner | Test_RUN_055_ProjectProtectedAdditionAndDependencyInterfaceEditAreRejected | RUN-055 | pass
- runner | Test_RUN_055_ProjectProtectedAdditionAndDependencyInterfaceEditAreRejected/add_project_protected | RUN-055 | pass
- runner | Test_RUN_055_ProjectProtectedAdditionAndDependencyInterfaceEditAreRejected/change_protected_dependency_interface | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes/change_api | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes/delete_eval | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes/add_foreign | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes/add_other | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionSubmissionAllowsOnlyDisposableAndNotes/change_preserved | RUN-055 | pass
- runner | Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound | RUN-056 | pass
- runner | Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound/agent_retry | RUN-056 | pass
- runner | Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound/verify_retry | RUN-056 | pass
- runner | Test_RUN_056_ConfiguredTimeoutBoundsDrillWorker | RUN-056 | pass
- runner | Test_RUN_057_UsageDoesNotSpendAttemptAndStopsWhenBudgetExhausted | RUN-057 | pass
- runner | Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail | RUN-058, RUN-059 | pass
- runner | Test_RUN_058_OwnCoverageRequiredAfterPassingCases | RUN-058 | pass
- runner | Test_RUN_059_DependentCoverageRequiredAfterPassingCases | RUN-059 | pass
- runner | Test_RUN_060_RUN_061_DrillReportAndRetainedRef | RUN-060, RUN-061 | pass
- runner | Test_RUN_060_RUN_061_FailedDrillReportsFailuresAndRetainsCode | RUN-060, RUN-061 | pass
- runner | Test_RUN_062_CleanupOnSuccessAndFailures | RUN-062 | pass
- runner | Test_RUN_062_CleanupOnSuccessAndFailures/success | RUN-062 | pass
- runner | Test_RUN_062_CleanupOnSuccessAndFailures/agent_retry | RUN-062 | pass
- runner | Test_RUN_062_CleanupOnSuccessAndFailures/change_api | RUN-062 | pass
- runner | Test_RUN_062_CleanupOnSuccessAndFailures/usage_always | RUN-062 | pass
- runner | Test_RUN_062_CleanupOnSuccessAndFailures/timeout | RUN-062 | pass
- runner | Test_RUN_062_CleanupFailureCannotPass | RUN-062 | pass
- runner | Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput | RUN-001, RUN-002, RUN-005, RUN-006 | pass
- runner | Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput/worker | RUN-001, RUN-002, RUN-005, RUN-006 | pass
- runner | Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput/audit | RUN-001, RUN-002, RUN-005, RUN-006 | pass
- runner | Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput/audit_fallback | RUN-001, RUN-002, RUN-005, RUN-006 | pass
- runner | Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues | RUN-003 | pass
- runner | Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues/shell | RUN-003 | pass
- runner | Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues/argv | RUN-003 | pass
- runner | Test_RUN_003_RUN_004_WrapperPrefixForBothRoles | RUN-003, RUN-004 | pass
- runner | Test_RUN_003_RUN_004_WrapperPrefixForBothRoles/worker | RUN-003, RUN-004 | pass
- runner | Test_RUN_003_RUN_004_WrapperPrefixForBothRoles/audit | RUN-003, RUN-004 | pass
- runner | Test_RUN_006_RUN_009_NonzeroExitRetainsStreamsAndCode | RUN-006, RUN-009 | pass
- runner | Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput | RUN-007, RUN-009 | pass
- runner | Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput/configured | RUN-007, RUN-009 | pass
- runner | Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput/request | RUN-007, RUN-009 | pass
- runner | Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant | RUN-007, RUN-009 | pass
- runner | Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant/timeout | RUN-007, RUN-009 | pass
- runner | Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant/cancel | RUN-007, RUN-009 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines | RUN-008 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/late_stderr_match | RUN-008 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/late_stdout_match | RUN-008 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/successful_match | RUN-008 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/outside_last_twenty | RUN-008 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/stdout_before_stderr_window | RUN-008 | pass
- runner | Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/literal_substring | RUN-008 | pass
- runner | Test_RUN_009_LaunchFailureReturnsError | RUN-009 | pass
- runner | Test_RUN_010_EmptyQueue | RUN-010 | pass
- runner | Test_RUN_010_AllBlockedQueue | RUN-010 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/unmerged | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/start_failure | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/first_unblocked | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/unapproved_archive_bound | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/shared_base_approval | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/allow_unmerged | RUN-011, RUN-012 | pass
- runner | Test_RUN_011_WaitingHeadSelectsIndependent | RUN-011 | pass
- runner | Test_RUN_010_EntirelyWaitingQueue | RUN-010 | pass
- runner | Test_RUN_013_RUN_014_CompleteSpecChange | RUN-013, RUN-014 | pass
- runner | Test_RUN_015_InvocationFailureRetainsHeldPhase | RUN-015 | pass
- runner | Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks | RUN-023, RUN-024 | pass
- runner | Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase | RUN-021, RUN-022 | pass
- runner | Test_RUN_021_ConfiguredThreeAttempts | RUN-021 | pass
- runner | Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks | RUN-023, RUN-024 | pass
- runner | Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules | RUN-043 | pass
- runner | Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions | RUN-044 | pass
- runner | Test_RUN_016_UnlimitedQueueOrder | RUN-016 | pass
- runner | Test_RUN_016_NegativeMaxIsUnlimited | RUN-016 | pass
- runner | Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding | RUN-017, RUN-018, RUN-019 | pass
- runner | Test_RUN_018_InitialEmptyAndExactLimit | RUN-018 | pass
- runner | Test_RUN_018_InitialEmptyAndExactLimit/initial_empty | RUN-018 | pass
- runner | Test_RUN_018_InitialEmptyAndExactLimit/exact_limit | RUN-018 | pass
- runner | Test_RUN_020_HeldFailureRetainsEarlierResults | RUN-020 | pass
- runner | Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent | RUN-016, RUN-025, RUN-026 | pass
- runner | Test_RUN_025_RUN_026_PositiveMaxRetainsBlock | RUN-025, RUN-026 | pass
- runner | Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied | RUN-036, RUN-037 | pass
- runner | Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded | RUN-036 | pass
- runner | Test_RUN_037_RUN_039_AgentDeadlineRetainsHeldWork | RUN-037, RUN-039 | pass
- runner | Test_RUN_037_VerificationDeadlineRetainsHeldWork | RUN-037 | pass
- runner | Test_RUN_038_UsageWaitPreflightStopsBeforeWaiting | RUN-038 | pass
- runner | Test_RUN_039_ClassifierAndParentCancellation | RUN-039 | pass
- runner | Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop | RUN-040 | pass
- runner | Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart | RUN-029 | pass
- runner | Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/bad_wait | RUN-029 | pass
- runner | Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/bad_budget | RUN-029 | pass
- runner | Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/zero_wait | RUN-029 | pass
- runner | Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/negative_budget | RUN-029 | pass
- runner | Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately | RUN-029, RUN-031 | pass
- runner | Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately/default_wait | RUN-029, RUN-031 | pass
- runner | Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately/default_budget | RUN-029, RUN-031 | pass
- runner | Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep | RUN-030, RUN-031, RUN-033 | pass
- runner | Test_RUN_033_OwnedResumeUsesApprovedUsageBudget | RUN-033 | pass
- runner | Test_RUN_030_UsageRetryPreservesVerificationFeedback | RUN-030 | pass
- runner | Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases | RUN-030, RUN-031 | pass
- runner | Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied | RUN-032 | pass
- runner | Test_RUN_034_UsageStopClassifierWrapping | RUN-034 | pass
- runner | Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure | RUN-035 | pass
- runner | Test_RUN_031_CancelledUsageWaitRetainsHeldWork | RUN-031 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	268.425s

- workflow | Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange | FLOW-089, FLOW-091, FLOW-094 | pass
- workflow | Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange | FLOW-089, FLOW-091, FLOW-094 | pass
- workflow | Test_FLOW_090_AbandonRejectsDirtyWorkingTree | FLOW-090 | pass
- workflow | Test_FLOW_090_AbandonRejectsDirtyWorkingTree/tracked | FLOW-090 | pass
- workflow | Test_FLOW_090_AbandonRejectsDirtyWorkingTree/untracked | FLOW-090 | pass
- workflow | Test_FLOW_091_AbandonRejectsForeignClaim | FLOW-091 | pass
- workflow | Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending | FLOW-092 | pass
- workflow | Test_FLOW_093_FLOW_095_AbandonRetainsRealAttemptAndFreesAssignedID | FLOW-093, FLOW-095 | pass
- workflow | Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal | FLOW-096 | pass
- workflow | Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal/changes/0001-manual.md | FLOW-096 | pass
- workflow | Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal/0001-manual.md | FLOW-096 | pass
- workflow | Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal/manual | FLOW-096 | pass
- workflow | Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState | FLOW-063, FLOW-064, FLOW-065, FLOW-066, FLOW-067 | pass
- workflow | Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit | FLOW-066 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead | FLOW-063 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/no_claim | FLOW-063 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/foreign_claim | FLOW-063 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/author | FLOW-063 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/phase | FLOW-063 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/audit | FLOW-063 | pass
- workflow | Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/changed_HEAD | FLOW-063 | pass
- workflow | Test_FLOW_087_FLOW_088_CommitsFollowPendingAndArchivePhaseContents | FLOW-087, FLOW-088 | pass
- workflow | Test_FLOW_087_CommitsUseFirstParentOfMerge | FLOW-087 | pass
- workflow | Test_FLOW_088_CommitsRejectMissingAndAmbiguousSelectors | FLOW-088 | pass
- workflow | Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot | FLOW-097 | pass
- workflow | Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts | FLOW-098 | pass
- workflow | Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure | FLOW-099 | pass
- workflow | Test_FLOW_099_DefaultBudgetBoundaries | FLOW-099 | pass
- workflow | Test_FLOW_099_DefaultBudgetBoundaries/equal | FLOW-099 | pass
- workflow | Test_FLOW_099_DefaultBudgetBoundaries/above | FLOW-099 | pass
- workflow | Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits | FLOW-100 | pass
- workflow | Test_FLOW_068_EligibleOrderAndEmptyQueues | FLOW-068 | pass
- workflow | Test_FLOW_068_EligibleOrderAndEmptyQueues/skips_waiting_and_retains_independent_order | FLOW-068 | pass
- workflow | Test_FLOW_068_EligibleOrderAndEmptyQueues/empty_queue | FLOW-068 | pass
- workflow | Test_FLOW_068_EligibleOrderAndEmptyQueues/all_waiting | FLOW-068 | pass
- workflow | Test_FLOW_069_EligibleWaitsForSameModule | FLOW-069 | pass
- workflow | Test_FLOW_070_EligibleWaitsForDirectAndTransitiveDependents | FLOW-070 | pass
- workflow | Test_FLOW_071_EligibleInterfaceChangeWaitsForBlockedDependent | FLOW-071 | pass
- workflow | Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet | FLOW-072 | pass
- workflow | Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet/preceding_set_removes_dependency | FLOW-072 | pass
- workflow | Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet/own_set_adds_or_removes_dependency | FLOW-072 | pass
- workflow | Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead | FLOW-011 | pass
- workflow | Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/empty_target_selects_first_eligible | FLOW-011 | pass
- workflow | Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/explicit_first_eligible | FLOW-011 | pass
- workflow | Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/reject_waiting | FLOW-011 | pass
- workflow | Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/reject_later-eligible | FLOW-011 | pass
- workflow | Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/no_eligible_change | FLOW-011 | pass
- workflow | Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits | FLOW-036, FLOW-078, FLOW-079 | pass
- workflow | Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals | FLOW-079 | pass
- workflow | Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals/failing | FLOW-079 | pass
- workflow | Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals/missing-listed-coverage | FLOW-079 | pass
- workflow | Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks | FLOW-078 | pass
- workflow | Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/evals_interface | FLOW-078 | pass
- workflow | Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/impl_implementation | FLOW-078 | pass
- workflow | Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit | FLOW-036 | pass
- workflow | Test_FLOW_001_FLOW_006_FLOW_007_FLOW_008_FLOW_012_FLOW_013_FLOW_014_FLOW_017_FLOW_018_FLOW_019_FLOW_053_StartAssignsLateIDsAndRetainsSnapshotWithoutCommit | FLOW-001, FLOW-006, FLOW-007, FLOW-008, FLOW-012, FLOW-013, FLOW-014, FLOW-017, FLOW-018, FLOW-019, FLOW-053 | pass
- workflow | Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead | FLOW-010, FLOW-011 | pass
- workflow | Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead/tracked | FLOW-010, FLOW-011 | pass
- workflow | Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead/untracked | FLOW-010, FLOW-011 | pass
- workflow | Test_FLOW_001_FLOW_016_FLOW_017_LinkedWorktreeExcludesCompetingOwner | FLOW-001, FLOW-016, FLOW-017 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/default-below | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/default-at | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/configured-at | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_055_VerifyAndDoneRejectUnsupportedHeldPhase | FLOW-055 | pass
- workflow | Test_FLOW_056_FLOW_061_CurrentAndContextIdentifyMissingClaimsAndChanges | FLOW-056, FLOW-061 | pass
- workflow | Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08 | FLOW-058, FLOW-062 | pass
- workflow | Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks | FLOW-002, FLOW-020, FLOW-021 | pass
- workflow | Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks/author | FLOW-002, FLOW-020, FLOW-021 | pass
- workflow | Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks/eval-zone | FLOW-002, FLOW-020, FLOW-021 | pass
- workflow | Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst | FLOW-022, FLOW-023, FLOW-031 | pass
- workflow | Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/missing-coverage | FLOW-022, FLOW-023, FLOW-031 | pass
- workflow | Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/passing-target-unrelated-failure | FLOW-022, FLOW-023, FLOW-031 | pass
- workflow | Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/already-satisfied | FLOW-022, FLOW-023, FLOW-031 | pass
- workflow | Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/fixes | FLOW-022, FLOW-023, FLOW-031 | pass
- workflow | Test_FLOW_023_AddedIDRequiresMappedFailingCase | FLOW-023 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/frozen-eval | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/direct | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/all | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/none | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/quick-boundary | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl | FLOW-032, FLOW-033, FLOW-046 | pass
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/false | FLOW-032, FLOW-033, FLOW-046 | pass
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/true | FLOW-032, FLOW-033, FLOW-046 | pass
- workflow | Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings | FLOW-034, FLOW-036, FLOW-041, FLOW-046 | pass
- workflow | Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound | FLOW-035, FLOW-037, FLOW-038, FLOW-039, FLOW-042, FLOW-043, FLOW-046 | pass
- workflow | Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion | FLOW-040, FLOW-042, FLOW-043 | pass
- workflow | Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives | FLOW-004, FLOW-028, FLOW-044, FLOW-045, FLOW-046, FLOW-047, FLOW-050 | pass
- workflow | Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead | FLOW-003, FLOW-024, FLOW-048, FLOW-049 | pass
- workflow | Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports | FLOW-005, FLOW-051, FLOW-052, FLOW-053 | pass
- workflow | Test_FLOW_054_StartOperationalPreconditionsHaveP01 | FLOW-054 | pass
- workflow | Test_FLOW_056_StartContentionAndCurrentForeignClaimHaveP02 | FLOW-056 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks | FLOW-057 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/author | FLOW-057 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/runtime | FLOW-057 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/audit | FLOW-057 | pass
- workflow | Test_FLOW_059_AuditNoFindingsErrorsHaveP09 | FLOW-059 | pass
- workflow | Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim | FLOW-060 | pass
- workflow | Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses | FLOW-062 | pass
- workflow | Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly | FLOW-077 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_114_explicit_endpoints | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_115_selected_identities | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_116_provenance_and_evals | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_117_phase_file_categories | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_118_preserved_edits | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_119_review_specs | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_120_prior_corrected_audit | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_121_decisions | FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 | pass
- workflow | Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts | FLOW-122, FLOW-123 | pass
- workflow | Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts/FLOW_122_endpoint_exceptions | FLOW-122, FLOW-123 | pass
- workflow | Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts/FLOW_123_latest_saved_costs | FLOW-122, FLOW-123 | pass
- workflow | Test_FLOW_126_PRBodyEndpointConfigError | FLOW-126 | pass
- workflow | Test_FLOW_127_PRBodyUnresolvableRangeIsNotEndpointConfigError | FLOW-127 | pass
- workflow | Test_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges | FLOW-124, FLOW-125 | pass
- workflow | Test_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges/FLOW_124_binary_and_large_artifacts | FLOW-124, FLOW-125 | pass
- workflow | Test_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges/FLOW_125_empty_and_invalid_ranges | FLOW-124, FLOW-125 | pass
- workflow | Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive | FLOW-080, FLOW-081, FLOW-083 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/missing | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/older_passing | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/newer_failed | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/newer_passing | FLOW-082 | pass
- workflow | Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook | FLOW-083 | pass
- workflow | Test_FLOW_084_ReleaseOwnedClaimPreservesDirtyWorkAndRecordedState | FLOW-084 | pass
- workflow | Test_FLOW_085_ReleaseMissingAndForeignClaimsPreserveRefs | FLOW-085 | pass
- workflow | Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors | FLOW-086 | pass
- workflow | Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors/path | FLOW-086 | pass
- workflow | Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors/filename | FLOW-086 | pass
- workflow | Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors/slug | FLOW-086 | pass
- workflow | Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptRestoresEvalsStepAndRetainsCurrentApproval | FLOW-073, FLOW-074, FLOW-075 | pass
- workflow | Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance | FLOW-073, FLOW-074, FLOW-075 | pass
- workflow | Test_FLOW_074_StartFromAttemptSelectsRequestedSavedRef | FLOW-074 | pass
- workflow | Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval | FLOW-075 | pass
- workflow | Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange | FLOW-073 | pass
- workflow | Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/dirty | FLOW-073 | pass
- workflow | Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/still_blocked | FLOW-073 | pass
- workflow | Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/not_first_eligible | FLOW-073 | pass
- workflow | Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/foreign_claim | FLOW-073 | pass
- workflow | Test_FLOW_076_StartFromAttemptRejectsInvalidOrMissingAttemptBeforeMutation | FLOW-076 | pass
- workflow | Test_FLOW_101_FLOW_102_FLOW_103_StatusQueueAndForwardWaiting | FLOW-101, FLOW-102, FLOW-103 | pass
- workflow | Test_FLOW_104_FLOW_105_StatusInterfaceAndProjectedWaiting | FLOW-104, FLOW-105 | pass
- workflow | Test_FLOW_106_FLOW_107_FLOW_108_StatusModuleMetadata | FLOW-106, FLOW-107, FLOW-108 | pass
- workflow | Test_FLOW_109_FLOW_111_StatusProvenanceDriftAndRecentModuleWarning | FLOW-109, FLOW-111 | pass
- workflow | Test_FLOW_110_StatusSharedBaseUsesSameArchivedDocumentIdentity | FLOW-110 | pass
- workflow | Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle | FLOW-112 | pass
- workflow | Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings | FLOW-113 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/workflow	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/workflow/evals	122.825s


## Contradicted Decisions
None identified. The retained `CodedError` implementation follows decision 0128.
## Gaps
- CON-008, target module `contracts`: The required round trip cannot hold for every possible `map[string]any` value. A nil map is not a YAML mapping, and YAML decoding does not preserve all Go concrete types. `RenderDocument` rejects a nil map and handles YAML-compatible values.
- CON-010, target module `contracts`: The assigned-ID heading grammar is not fully specified. The implementation accepts an uppercase alphanumeric prefix, a hyphen, and a nonempty decimal suffix, following the retained examples.
- CON-002, target module `contracts`: The behavior for multiple level-two sections with the same name is unspecified. `Section` and `ReplaceSection` use the first matching section.
## Agent observations

### Attempt 1 (exit 0)
Implemented the contracts primitives in [primitives.go](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go) and recorded contradictions and gaps in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/DRILL_NOTES.md). Retained inputs were left untouched, and nothing was committed.

`gofmt` passed. I could not run the evals here: `go test ./...` stops because the supplied workspace has no Go module manifest.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a19553-3209-70a2-ad49-d5a426e503bb
--------
user
# Rebuild contracts

Target spec and Decisions:
---
depends_on: []
id_prefix: CON
interface:
    - api.go
    - model/**
name: contracts
---

# contracts

## Specs

### CON-001
ParseDocument MUST decode leading --- delimited YAML mapping frontmatter and retain the remaining Markdown body.

### CON-002
Section MUST return the body and presence of a named level-two Markdown section.

### CON-003
SpecHash MUST return the first eight lowercase hex characters of SHA-256 over the rendered item, including its ID and tag, after normalizing CRLF to LF, removing spaces/tabs at line ends and canonicalizing trailing blank lines to one final newline.

### CON-004
MatchesID MUST find a spec ID in either testcase name or classname, ignoring case and treating hyphen, underscore and an omitted separator equally. It MUST reject a match preceded by a letter/digit or followed by a digit.

### CON-005
MatchGlob MUST support slash-separated paths with * within one component and ** spanning zero or more components.

### CON-006
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens.

### CON-007
ParseDocument MUST reject malformed, non-mapping or unterminated frontmatter.

### CON-008
RenderDocument MUST produce a document whose parsed frontmatter and body equal its input.

### CON-009
ReplaceSection MUST replace a named level-two Markdown section without changing other sections, appending the section when absent.

### CON-010
ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, including mixed assigned-ID and handle items.

### CON-011
ParseItems MUST reject malformed item headings.

### CON-012
ParseItems MUST reject duplicate item IDs, including duplicate handles.

### CON-013
RenderItem MUST reproduce an item's heading, tag and body.

### CON-014
ResolveZone MUST choose the first matching zone in order protected, SPEC.md, evals, interface, append_only, preserve, disposable, other, treating SPEC.md as the module's root spec path.

### CON-015
ParseItems MUST reject empty handles and handles containing characters other than lowercase letters, digits and hyphens.

### CON-016
ErrorCode MUST return the code and subject of the first CodedError in an error chain, including one wrapped with %w, and MUST report not ok when the chain has none.

### CON-017
ParseItems MUST return a CodedError with code C05 and the duplicated ID or handle as subject when a Specs section repeats an item heading.

### CON-018
ParseItems MUST return a CodedError with code C03 when an item heading is malformed.

### CON-019
ParseDocument MUST return a CodedError with code C03 when frontmatter is missing, unterminated, invalid YAML or not a mapping.

## Decisions

[0128-coded-parse-errors] CodedError exposes an optional Cause error. Error returns the cause's text when present and the code otherwise; Unwrap returns Cause.


Retained interfaces and files:
packages/contracts/SPEC.md
packages/contracts/api.go
packages/contracts/model/model.go

Retained evals:
packages/contracts/evals/.gitkeep
packages/contracts/evals/contracts_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.

codex
I’ll inspect the supplied contracts files and evals, then implement the disposable paths and run the relevant tests. I’ll keep the retained inputs untouched.
exec
/bin/zsh -lc "pwd && rg --files -g '"'!**/.git/**'"' -g '"'!**/.agents/**'"' -g '"'!**/.codex/**'"' -g '"'!**/.aws/**'"' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
 succeeded in 455ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
SYSTEM.md
packages/contracts/SPEC.md
packages/contracts/api.go
packages/contracts/evals/contracts_test.go
packages/contracts/model/model.go

exec
/bin/zsh -lc 'cat packages/contracts/api.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc 'cat SYSTEM.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc 'cat packages/contracts/evals/contracts_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc 'cat packages/contracts/SPEC.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc "find . -maxdepth 4 -type f -not -path './.git/*' -print | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc 'cat packages/contracts/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
 succeeded in 785ms:
package model

// Document is a Markdown document with YAML frontmatter.
type Document struct {
	Frontmatter map[string]any
	Body        string
}

// Item is one level-three spec item. Body excludes the heading line.
type Item struct {
	ID   string
	Tag  string
	Body string
}

// CodedError carries a stable diagnostic code while preserving the cause's text.
type CodedError struct {
	Code    string
	Subject string
	Cause   error
}

func (e CodedError) Error() string {
	if e.Cause != nil {
		return e.Cause.Error()
	}
	return e.Code
}

func (e CodedError) Unwrap() error { return e.Cause }

// Zones holds module-relative path patterns for each editable zone.
type Zones struct {
	Protected  []string
	Evals      []string
	Interface  []string
	AppendOnly []string
	Preserve   []string
	Disposable []string
}

 succeeded in 793ms:
# Independent contracts reconstruction

Reconstruct only contracts from the supplied SPEC.md (including Decisions),
retained public interfaces and evals. Contracts has no dependency interfaces.
These retained files are the behavioral authority for this drill.

Stay inside the supplied input workspace. Do not read external trees, old
implementation, source history, saved attempts, binaries, bootstrap, or another
module's internals. Do not launch nested agents. The fresh input Git repository
contains only the supplied inputs; do not seek another repository or remote.

Write only packages/contracts/internal/** and DRILL_NOTES.md. Leave all retained
inputs unchanged. Do not commit. Use ordinary Go source and gofmt.

The full project's build manifests, hooks and dependent implementations stay in
a separate verifier checkout. Their absence here is deliberate: do not search
for them or create files outside the allowed paths. The controller compiles and
runs own and all transitive dependent evals and supplies bounded failure feedback.

Record separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md.
For each actual gap, name the nearest spec ID (or "no spec") and target module.
Report contradictions and uncertainty honestly; do not modify specs, Decisions,
interfaces or evals to resolve them. This run authorizes no gap fixes.

 succeeded in 797ms:
package contracts

import (
	"github.com/lurkshark/fireweed/packages/contracts/internal/primitives"
	"github.com/lurkshark/fireweed/packages/contracts/model"
)

type Document = model.Document
type Item = model.Item
type Zones = model.Zones
type CodedError = model.CodedError

// ErrorCode returns the first coded diagnostic in an error chain.
func ErrorCode(err error) (code, subject string, ok bool) { return primitives.ErrorCode(err) }

func ParseDocument(source string) (Document, error) { return primitives.ParseDocument(source) }
func RenderDocument(doc Document) (string, error)   { return primitives.RenderDocument(doc) }

// Section returns the section body without its level-two heading.
func Section(markdown, name string) (string, bool) { return primitives.Section(markdown, name) }

// ReplaceSection replaces the body of a level-two section, or appends it.
func ReplaceSection(markdown, name, body string) string {
	return primitives.ReplaceSection(markdown, name, body)
}

func ParseItems(section string) ([]Item, error) { return primitives.ParseItems(section) }
func RenderItem(item Item) string               { return primitives.RenderItem(item) }
func SpecHash(item Item) string                 { return primitives.SpecHash(item) }

func MatchesID(id, testcaseName, classname string) bool {
	return primitives.MatchesID(id, testcaseName, classname)
}

func MatchGlob(pattern, path string) bool { return primitives.MatchGlob(pattern, path) }

// ResolveZone classifies a module-relative path.
func ResolveZone(path string, zones Zones) string { return primitives.ResolveZone(path, zones) }

 succeeded in 797ms:
package evals_test

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
)

func Test_CON_001_ParseDocument(t *testing.T) {
	source := "---\nname: contracts\ncount: 2\n---\n# Title\n\nBody  \n"
	doc, err := contracts.ParseDocument(source)
	if err != nil {
		t.Fatalf("ParseDocument: %v", err)
	}
	if !reflect.DeepEqual(doc.Frontmatter, map[string]any{"name": "contracts", "count": 2}) {
		t.Errorf("frontmatter = %#v", doc.Frontmatter)
	}
	if doc.Body != "# Title\n\nBody  \n" {
		t.Errorf("body = %q", doc.Body)
	}
}

func Test_CON_007_RejectInvalidFrontmatter(t *testing.T) {
	for _, tc := range []struct{ name, source string }{
		{"malformed", "---\nname: [broken\n---\nbody\n"},
		{"non_mapping", "---\n- first\n---\nbody\n"},
		{"unterminated", "---\nname: contracts\nbody\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := contracts.ParseDocument(tc.source); err == nil {
				t.Errorf("ParseDocument(%q) accepted invalid frontmatter", tc.source)
			}
		})
	}
}

func Test_CON_008_RenderRoundTrip(t *testing.T) {
	want := contracts.Document{
		Frontmatter: map[string]any{"name": "contracts", "enabled": true},
		Body:        "# Title\n\nBody\n",
	}
	source, err := contracts.RenderDocument(want)
	if err != nil {
		t.Fatalf("RenderDocument: %v", err)
	}
	got, err := contracts.ParseDocument(source)
	if err != nil {
		t.Fatalf("parse rendered document: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip = %#v; want %#v", got, want)
	}
}

func Test_CON_002_CON_009_Sections(t *testing.T) {
	source := "intro\n## Alpha\nA  \n\n## Beta\nB\n\n## Gamma\nC\n"
	if got, ok := contracts.Section(source, "Beta"); !ok || got != "B\n\n" {
		t.Errorf("Section(Beta) = %q, %v", got, ok)
	}
	if _, ok := contracts.Section(source, "Missing"); ok {
		t.Error("Section(Missing) was reported present")
	}
	want := "intro\n## Alpha\nA  \n\n## Beta\nnew\n## Gamma\nC\n"
	if got := contracts.ReplaceSection(source, "Beta", "new\n"); got != want {
		t.Errorf("replacement changed section or surrounding bytes:\n got %q\nwant %q", got, want)
	}
}

func Test_CON_009_CON_002_AppendSection(t *testing.T) {
	source := "# Doc\n\n## Existing\nkeep\n"
	got := contracts.ReplaceSection(source, "New", "added\n")
	if !strings.HasPrefix(got, source) {
		t.Errorf("append changed original bytes: %q", got)
	}
	if body, ok := contracts.Section(got, "New"); !ok || body != "added\n" {
		t.Errorf("appended section = %q, %v", body, ok)
	}
}

func Test_CON_010_ParseItems(t *testing.T) {
	source := "### CON-001 [draft]\nFirst line\n\nSecond line\n### CON-002\nNext\n"
	want := []contracts.Item{
		{ID: "CON-001", Tag: "draft", Body: "First line\n\nSecond line\n"},
		{ID: "CON-002", Body: "Next\n"},
	}
	got, err := contracts.ParseItems(source)
	if err != nil {
		t.Fatalf("ParseItems: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("items = %#v; want %#v", got, want)
	}
}

func Test_CON_011_CON_012_RejectInvalidItems(t *testing.T) {
	for _, tc := range []struct{ name, source string }{
		{"CON_012_duplicate", "### CON-001\nFirst\n### CON-001\nSecond\n"},
		{"CON_011_malformed_heading", "### CON-001 [draft\nBody\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := contracts.ParseItems(tc.source); err == nil {
				t.Errorf("ParseItems(%q) accepted invalid items", tc.source)
			}
		})
	}
}

func Test_CON_006_CON_010_ParseHandlesInSourceOrder(t *testing.T) {
	source := "### new:alpha-2 [draft]\nFirst line\n\nSecond line\n### CON-007\nAssigned body\n### new:7-beta\nHandle body\n"
	want := []contracts.Item{
		{ID: "new:alpha-2", Tag: "draft", Body: "First line\n\nSecond line\n"},
		{ID: "CON-007", Body: "Assigned body\n"},
		{ID: "new:7-beta", Body: "Handle body\n"},
	}
	got, err := contracts.ParseItems(source)
	if err != nil {
		t.Fatalf("ParseItems: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("items = %#v; want %#v", got, want)
	}
}

func Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles(t *testing.T) {
	for _, tc := range []struct{ name, source string }{
		{"CON_015_empty", "### new:\nBody\n"},
		{"CON_015_uppercase", "### new:Alpha\nBody\n"},
		{"CON_015_underscore", "### new:alpha_beta\nBody\n"},
		{"CON_015_punctuation", "### new:alpha.beta\nBody\n"},
		{"CON_012_duplicate", "### new:alpha-2\nFirst\n### new:alpha-2\nSecond\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := contracts.ParseItems(tc.source); err == nil {
				t.Errorf("ParseItems(%q) accepted invalid handles", tc.source)
			}
		})
	}
}

func Test_CON_013_RenderItem(t *testing.T) {
	for _, tc := range []struct {
		item contracts.Item
		want string
	}{
		{contracts.Item{ID: "CON-001", Tag: "draft", Body: "Body\n"}, "### CON-001 [draft]\nBody\n"},
		{contracts.Item{ID: "CON-002", Body: "Line one\n\nLine two\n"}, "### CON-002\nLine one\n\nLine two\n"},
	} {
		if got := contracts.RenderItem(tc.item); got != tc.want {
			t.Errorf("RenderItem(%#v) = %q; want %q", tc.item, got, tc.want)
		}
	}
}

func Test_CON_003_CanonicalHash(t *testing.T) {
	item := contracts.Item{ID: "CON-013", Tag: "draft", Body: "alpha  \r\nbeta\t\r\n\r\n\r\n"}
	canonical := "### CON-013 [draft]\nalpha\nbeta\n"
	sum := sha256.Sum256([]byte(canonical))
	want := hex.EncodeToString(sum[:])[:8]
	if got := contracts.SpecHash(item); got != want {
		t.Errorf("SpecHash = %q; want %q for canonical item %q", got, want, canonical)
	}
	if got := contracts.SpecHash(contracts.Item{ID: "CON-013", Tag: "draft", Body: "alpha\nbeta\n"}); got != want {
		t.Errorf("equivalent LF item hash = %q; want %q", got, want)
	}
}

func Test_CON_003_IDAndTagAffectHash(t *testing.T) {
	base := contracts.Item{ID: "CON-013", Tag: "draft", Body: "body\n"}
	baseHash := contracts.SpecHash(base)
	withID := base
	withID.ID = "CON-014"
	withTag := base
	withTag.Tag = "ready"
	if baseHash == contracts.SpecHash(withID) || baseHash == contracts.SpecHash(withTag) {
		t.Errorf("ID and tag must affect hash: base=%q id=%q tag=%q", baseHash, contracts.SpecHash(withID), contracts.SpecHash(withTag))
	}
}

func Test_CON_004_PositiveMatches(t *testing.T) {
	for _, tc := range []struct{ name, class string }{
		{"Test_con-001_checks", "Suite"},
		{"Test_CON_001", "Suite"},
		{"Test_CON001", "Suite"},
		{"unrelated", "checks.con_001"},
	} {
		if !contracts.MatchesID("CON-001", tc.name, tc.class) {
			t.Errorf("MatchesID(CON-001, %q, %q) = false", tc.name, tc.class)
		}
	}
}

func Test_CON_004_Boundaries(t *testing.T) {
	for _, tc := range []struct{ name, class string }{
		{"XCON-001", "Suite"},
		{"7CON_001", "Suite"},
		{"CON0012", "Suite"},
	} {
		if contracts.MatchesID("CON-001", tc.name, tc.class) {
			t.Errorf("MatchesID accepted embedded or extended ID in %q", tc.name)
		}
	}
}

func Test_CON_005_MatchGlob(t *testing.T) {
	for _, tc := range []struct {
		pattern, path string
		want          bool
	}{
		{"evals/*.go", "evals/check_test.go", true},
		{"evals/*.go", "evals/nested/check.go", false},
		{"internal/**/test.go", "internal/test.go", true},
		{"internal/**/test.go", "internal/a/b/test.go", true},
		{"internal/**/*.go", "internal/a/b/test.go", true},
		{"api.go", "sub/api.go", false},
	} {
		if got := contracts.MatchGlob(tc.pattern, tc.path); got != tc.want {
			t.Errorf("MatchGlob(%q, %q) = %v; want %v", tc.pattern, tc.path, got, tc.want)
		}
	}
}

func Test_CON_014_ResolveZone(t *testing.T) {
	for _, tc := range []struct {
		name, path, want string
		zones            contracts.Zones
	}{
		{"protected_first", "api.go", "protected", contracts.Zones{Protected: []string{"api.go"}, Interface: []string{"api.go"}}},
		{"root_spec", "SPEC.md", "spec", contracts.Zones{Evals: []string{"*"}}},
		{"evals_before_interface", "evals/check.go", "evals", contracts.Zones{Evals: []string{"evals/**"}, Interface: []string{"**"}}},
		{"interface_before_append", "api.go", "interface", contracts.Zones{Interface: []string{"api.go"}, AppendOnly: []string{"*.go"}}},
		{"append_before_preserve", "notes.md", "append_only", contracts.Zones{AppendOnly: []string{"*.md"}, Preserve: []string{"*.md"}}},
		{"preserve_before_disposable", "scratch.go", "preserve", contracts.Zones{Preserve: []string{"*.go"}, Disposable: []string{"*.go"}}},
		{"disposable", "internal/logic.go", "disposable", contracts.Zones{Disposable: []string{"internal/**"}}},
		{"other", "sub/SPEC.md", "other", contracts.Zones{}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := contracts.ResolveZone(tc.path, tc.zones); got != tc.want {
				t.Errorf("ResolveZone(%q) = %q; want %q", tc.path, got, tc.want)
			}
		})
	}
}

func Test_CON_016_ErrorCode(t *testing.T) {
	cause := errors.New("cause text")
	inner := contracts.CodedError{Code: "C03", Subject: "inner", Cause: cause}
	outer := contracts.CodedError{Code: "C05", Subject: "outer", Cause: inner}
	noCause := contracts.CodedError{Code: "C07", Subject: "standalone"}
	for _, tc := range []struct {
		name, code, subject string
		err                 error
		ok                  bool
	}{
		{"direct", "C03", "inner", inner, true},
		{"wrapped", "C03", "inner", fmt.Errorf("context: %w", &inner), true},
		{"first_in_chain", "C05", "outer", fmt.Errorf("context: %w", &outer), true},
		{"no_cause", "C07", "standalone", noCause, true},
		{"plain", "", "", fmt.Errorf("context: %w", cause), false},
		{"nil", "", "", nil, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			code, subject, ok := contracts.ErrorCode(tc.err)
			if code != tc.code || subject != tc.subject || ok != tc.ok {
				t.Errorf("ErrorCode(%v) = (%q, %q, %v); want (%q, %q, %v)", tc.err, code, subject, ok, tc.code, tc.subject, tc.ok)
			}
		})
	}
	if inner.Error() != cause.Error() || !errors.Is(inner, cause) {
		t.Errorf("CodedError did not preserve and wrap its cause: %v", inner)
	}
}

func Test_CON_017_ParseItemsDuplicateCode(t *testing.T) {
	for _, tc := range []struct{ name, id string }{
		{"assigned", "CON-001"},
		{"handle", "new:alpha-2"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			id := tc.id
			_, err := contracts.ParseItems("### " + id + "\nFirst\n### " + id + "\nSecond\n")
			if err == nil {
				t.Fatal("ParseItems accepted duplicate heading")
			}
			if !hasCodedError(err) {
				t.Fatalf("duplicate error has no CodedError: %T: %v", err, err)
			}
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "C05" || subject != id {
				t.Errorf("ErrorCode = (%q, %q, %v); want (C05, %q, true)", code, subject, ok, id)
			}
			if want := "duplicate item ID: " + id; err.Error() != want {
				t.Errorf("duplicate message = %q; want %q", err, want)
			}
		})
	}
}

func Test_CON_018_ParseItemsMalformedCode(t *testing.T) {
	for _, tc := range []struct{ name, heading string }{
		{"unclosed_tag", "### CON-001 [draft"},
		{"invalid_handle", "### new:Alpha"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			heading := tc.heading
			_, err := contracts.ParseItems(heading + "\nBody\n")
			if err == nil {
				t.Fatal("ParseItems accepted malformed heading")
			}
			if !hasCodedError(err) {
				t.Fatalf("malformed heading error has no CodedError: %T: %v", err, err)
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C03" {
				t.Errorf("ErrorCode = (%q, _, %v); want (C03, _, true)", code, ok)
			}
			if want := fmt.Sprintf("malformed item heading: %q", heading); err.Error() != want {
				t.Errorf("malformed heading message = %q; want %q", err, want)
			}
		})
	}
}

func Test_CON_019_ParseDocumentFrontmatterCode(t *testing.T) {
	for _, tc := range []struct {
		name, source, message string
	}{
		{"missing", "# Body\n", "missing YAML frontmatter"},
		{"unterminated", "---\nname: contracts\nbody\n", "unterminated YAML frontmatter"},
		{"invalid_yaml", "---\nname: [broken\n---\nbody\n", "invalid YAML frontmatter:"},
		{"non_mapping", "---\n- first\n---\nbody\n", "invalid YAML frontmatter:"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, err := contracts.ParseDocument(tc.source)
			if err == nil {
				t.Fatal("ParseDocument accepted invalid frontmatter")
			}
			if !hasCodedError(err) {
				t.Fatalf("frontmatter error has no CodedError: %T: %v", err, err)
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C03" {
				t.Errorf("ErrorCode = (%q, _, %v); want (C03, _, true)", code, ok)
			}
			if !strings.HasPrefix(err.Error(), tc.message) {
				t.Errorf("frontmatter message = %q; want prefix %q", err, tc.message)
			}
		})
	}
}

func hasCodedError(err error) bool {
	var pointer *contracts.CodedError
	if errors.As(err, &pointer) {
		return true
	}
	var value contracts.CodedError
	return errors.As(err, &value)
}

 succeeded in 800ms:
./SYSTEM.md
./packages/contracts/SPEC.md
./packages/contracts/api.go
./packages/contracts/evals/.gitkeep
./packages/contracts/evals/contracts_test.go
./packages/contracts/model/model.go

 succeeded in 821ms:
---
depends_on: []
id_prefix: CON
interface:
    - api.go
    - model/**
name: contracts
---

# contracts

## Specs

### CON-001
ParseDocument MUST decode leading --- delimited YAML mapping frontmatter and retain the remaining Markdown body.

### CON-002
Section MUST return the body and presence of a named level-two Markdown section.

### CON-003
SpecHash MUST return the first eight lowercase hex characters of SHA-256 over the rendered item, including its ID and tag, after normalizing CRLF to LF, removing spaces/tabs at line ends and canonicalizing trailing blank lines to one final newline.

### CON-004
MatchesID MUST find a spec ID in either testcase name or classname, ignoring case and treating hyphen, underscore and an omitted separator equally. It MUST reject a match preceded by a letter/digit or followed by a digit.

### CON-005
MatchGlob MUST support slash-separated paths with * within one component and ** spanning zero or more components.

### CON-006
ParseItems MUST accept new:<slug> item headings and preserve the complete handle in Item.ID, where slug contains lowercase letters, digits and hyphens.

### CON-007
ParseDocument MUST reject malformed, non-mapping or unterminated frontmatter.

### CON-008
RenderDocument MUST produce a document whose parsed frontmatter and body equal its input.

### CON-009
ReplaceSection MUST replace a named level-two Markdown section without changing other sections, appending the section when absent.

### CON-010
ParseItems MUST return level-three items in source order with ID, optional bracketed tag and complete body, including mixed assigned-ID and handle items.

### CON-011
ParseItems MUST reject malformed item headings.

### CON-012
ParseItems MUST reject duplicate item IDs, including duplicate handles.

### CON-013
RenderItem MUST reproduce an item's heading, tag and body.

### CON-014
ResolveZone MUST choose the first matching zone in order protected, SPEC.md, evals, interface, append_only, preserve, disposable, other, treating SPEC.md as the module's root spec path.

### CON-015
ParseItems MUST reject empty handles and handles containing characters other than lowercase letters, digits and hyphens.

### CON-016
ErrorCode MUST return the code and subject of the first CodedError in an error chain, including one wrapped with %w, and MUST report not ok when the chain has none.

### CON-017
ParseItems MUST return a CodedError with code C05 and the duplicated ID or handle as subject when a Specs section repeats an item heading.

### CON-018
ParseItems MUST return a CodedError with code C03 when an item heading is malformed.

### CON-019
ParseDocument MUST return a CodedError with code C03 when frontmatter is missing, unterminated, invalid YAML or not a mapping.

## Decisions

[0128-coded-parse-errors] CodedError exposes an optional Cause error. Error returns the cause's text when present and the code otherwise; Unwrap returns Cause.

codex
The retained API delegates to `internal/primitives`, so that is the implementation path. The evals also specify the parse error messages and codes; I’ll preserve those while implementing the other contract behavior.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..cc5689261e1347c96c45038c24c36137b498c791
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,382 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	if !strings.HasPrefix(source, "---\n") {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	rest := source[len("---\n"):]
+	end := strings.Index(rest, "\n---\n")
+	if end < 0 {
+		if strings.HasPrefix(rest, "---\n") {
+			end = 0
+		} else if strings.HasSuffix(rest, "\n---") {
+			end = len(rest) - len("\n---")
+		} else if rest == "---" {
+			end = 0
+		} else {
+			return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+		}
+	}
+	front := rest[:end]
+	var body string
+	if len(rest) > end+len("\n---") {
+		body = rest[end+len("\n---\n"):]
+	}
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	return i, strings.TrimSpace(trimmed[i:]), true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	for _, l := range lines(markdown) {
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	for _, l := range lines(section) {
+		if !strings.HasPrefix(l.text, "###") || strings.HasPrefix(l.text, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(l.text)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[itemsBodyStart(items, section):l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+	}
+	return items, nil
+}
+
+func itemsBodyStart(items []model.Item, section string) int {
+	// Kept separate from the item values so an item's body can remain byte exact.
+	_ = items
+	_ = section
+	return 0
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func asciiAlphaNum(ch byte) bool {
+	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) || value[j] < '0' || value[j] > '9' {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..fee12e94705f437491f21307db4b0822a9d184e4
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,381 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	if !strings.HasPrefix(source, "---\n") {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	rest := source[len("---\n"):]
+	end := strings.Index(rest, "\n---\n")
+	if end < 0 {
+		if strings.HasPrefix(rest, "---\n") {
+			end = 0
+		} else if strings.HasSuffix(rest, "\n---") {
+			end = len(rest) - len("\n---")
+		} else if rest == "---" {
+			end = 0
+		} else {
+			return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+		}
+	}
+	front := rest[:end]
+	var body string
+	if len(rest) > end+len("\n---") {
+		body = rest[end+len("\n---\n"):]
+	}
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	return i, strings.TrimSpace(trimmed[i:]), true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	for _, l := range lines(markdown) {
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	bodyStart := 0
+	for _, l := range lines(section) {
+		candidate := strings.TrimLeft(l.text, " ")
+		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(candidate)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[bodyStart:l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+		bodyStart = l.end
+	}
+	if len(items) > 0 {
+		items[len(items)-1].Body = section[bodyStart:]
+	}
+	return items, nil
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func asciiAlphaNum(ch byte) bool {
+	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) || value[j] < '0' || value[j] > '9' {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..75ce92a09e2493f7e1b210e69b798b5fe67d710f
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,377 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+			return "", "", false
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	allLines := lines(source)
+	if len(allLines) == 0 || allLines[0].text != "---" || allLines[0].end == len(source) && !strings.HasSuffix(source, "\n") {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	closing := -1
+	for i := 1; i < len(allLines); i++ {
+		if allLines[i].text == "---" {
+			closing = i
+			break
+		}
+	}
+	if closing < 0 {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+	}
+	front := source[allLines[0].end:allLines[closing].start]
+	body := source[allLines[closing].end:]
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	return i, strings.TrimSpace(trimmed[i:]), true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	for _, l := range lines(markdown) {
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	bodyStart := 0
+	for _, l := range lines(section) {
+		candidate := strings.TrimLeft(l.text, " ")
+		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(candidate)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[bodyStart:l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+		bodyStart = l.end
+	}
+	if len(items) > 0 {
+		items[len(items)-1].Body = section[bodyStart:]
+	}
+	return items, nil
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func asciiAlphaNum(ch byte) bool {
+	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) || value[j] < '0' || value[j] > '9' {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..05959e63e7adb6a70e8c419301bd90f75bd031ea
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,377 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+			return "", "", false
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	allLines := lines(source)
+	if len(allLines) == 0 || allLines[0].text != "---" {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	closing := -1
+	for i := 1; i < len(allLines); i++ {
+		if allLines[i].text == "---" {
+			closing = i
+			break
+		}
+	}
+	if closing < 0 {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+	}
+	front := source[allLines[0].end:allLines[closing].start]
+	body := source[allLines[closing].end:]
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	return i, strings.TrimSpace(trimmed[i:]), true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	for _, l := range lines(markdown) {
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	bodyStart := 0
+	for _, l := range lines(section) {
+		candidate := strings.TrimLeft(l.text, " ")
+		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(candidate)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[bodyStart:l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+		bodyStart = l.end
+	}
+	if len(items) > 0 {
+		items[len(items)-1].Body = section[bodyStart:]
+	}
+	return items, nil
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func asciiAlphaNum(ch byte) bool {
+	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) || value[j] < '0' || value[j] > '9' {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

codex
The primitives are in place. I’m reviewing the edge cases now, especially byte preservation in section and item bodies and the YAML round trip. The supplied workspace has no build manifest, so the available test run may be limited.
exec
/bin/zsh -lc "gofmt -w packages/contracts/internal/primitives/primitives.go && sed -n '1,420p' packages/contracts/internal/primitives/primitives.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
 succeeded in 483ms:
package primitives

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts/model"
	"gopkg.in/yaml.v3"
)

func ErrorCode(err error) (code, subject string, ok bool) {
	for err != nil {
		switch e := err.(type) {
		case model.CodedError:
			return e.Code, e.Subject, true
		case *model.CodedError:
			if e != nil {
				return e.Code, e.Subject, true
			}
			return "", "", false
		}
		if many, ok := err.(interface{ Unwrap() []error }); ok {
			for _, child := range many.Unwrap() {
				if code, subject, ok := ErrorCode(child); ok {
					return code, subject, true
				}
			}
			return "", "", false
		}
		err = errors.Unwrap(err)
	}
	return "", "", false
}

func ParseDocument(source string) (model.Document, error) {
	allLines := lines(source)
	if len(allLines) == 0 || allLines[0].text != "---" {
		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
	}
	closing := -1
	for i := 1; i < len(allLines); i++ {
		if allLines[i].text == "---" {
			closing = i
			break
		}
	}
	if closing < 0 {
		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
	}
	front := source[allLines[0].end:allLines[closing].start]
	body := source[allLines[closing].end:]
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
	}
	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
	}
	var mapping map[string]any
	if err := node.Content[0].Decode(&mapping); err != nil {
		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
	}
	return model.Document{Frontmatter: mapping, Body: body}, nil
}

func RenderDocument(doc model.Document) (string, error) {
	if doc.Frontmatter == nil {
		return "", errors.New("frontmatter must be a mapping")
	}
	data, err := yaml.Marshal(doc.Frontmatter)
	if err != nil {
		return "", err
	}
	return "---\n" + string(data) + "---\n" + doc.Body, nil
}

type line struct {
	start, end int
	text       string
}

func lines(source string) []line {
	var out []line
	for start := 0; start < len(source); {
		next := strings.IndexByte(source[start:], '\n')
		end := len(source)
		if next >= 0 {
			end = start + next + 1
		}
		text := source[start:end]
		text = strings.TrimSuffix(text, "\n")
		text = strings.TrimSuffix(text, "\r")
		out = append(out, line{start: start, end: end, text: text})
		start = end
	}
	return out
}

func heading(text string) (level int, name string, ok bool) {
	trimmed := strings.TrimLeft(text, " ")
	if len(text)-len(trimmed) > 3 {
		return 0, "", false
	}
	i := 0
	for i < len(trimmed) && trimmed[i] == '#' {
		i++
	}
	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
		return 0, "", false
	}
	return i, strings.TrimSpace(trimmed[i:]), true
}

func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
	start := -1
	for _, l := range lines(markdown) {
		level, title, ok := heading(l.text)
		if !ok {
			continue
		}
		if start >= 0 && level <= 2 {
			return start, l.start, true
		}
		if start < 0 && level == 2 && title == name {
			start = l.end
		}
	}
	if start >= 0 {
		return start, len(markdown), true
	}
	return 0, 0, false
}

func Section(markdown, name string) (string, bool) {
	start, end, ok := sectionBounds(markdown, name)
	if !ok {
		return "", false
	}
	return markdown[start:end], true
}

func ReplaceSection(markdown, name, body string) string {
	start, end, ok := sectionBounds(markdown, name)
	if ok {
		return markdown[:start] + body + markdown[end:]
	}
	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
		markdown += "\n"
	}
	return markdown + "## " + name + "\n" + body
}

func validItemID(id string) bool {
	if strings.HasPrefix(id, "new:") {
		slug := strings.TrimPrefix(id, "new:")
		if slug == "" {
			return false
		}
		for _, ch := range slug {
			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
				return false
			}
		}
		return true
	}
	parts := strings.Split(id, "-")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return false
	}
	for _, ch := range parts[0] {
		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
			return false
		}
	}
	for _, ch := range parts[1] {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func parseItemHeading(text string) (id, tag string, ok bool) {
	if !strings.HasPrefix(text, "### ") {
		return "", "", false
	}
	content := strings.TrimPrefix(text, "### ")
	parts := strings.SplitN(content, " ", 2)
	id = parts[0]
	if !validItemID(id) {
		return "", "", false
	}
	if len(parts) == 1 {
		return id, "", true
	}
	remaining := parts[1]
	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
		return "", "", false
	}
	tag = remaining[1 : len(remaining)-1]
	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
		return "", "", false
	}
	return id, tag, true
}

func ParseItems(section string) ([]model.Item, error) {
	var items []model.Item
	seen := make(map[string]bool)
	bodyStart := 0
	for _, l := range lines(section) {
		candidate := strings.TrimLeft(l.text, " ")
		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
			continue
		}
		id, tag, ok := parseItemHeading(candidate)
		if !ok {
			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
		}
		if seen[id] {
			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
		}
		seen[id] = true
		if len(items) > 0 {
			items[len(items)-1].Body = section[bodyStart:l.start]
		}
		items = append(items, model.Item{ID: id, Tag: tag})
		bodyStart = l.end
	}
	if len(items) > 0 {
		items[len(items)-1].Body = section[bodyStart:]
	}
	return items, nil
}

func RenderItem(item model.Item) string {
	heading := "### " + item.ID
	if item.Tag != "" {
		heading += " [" + item.Tag + "]"
	}
	return heading + "\n" + item.Body
}

func SpecHash(item model.Item) string {
	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
	parts := strings.Split(value, "\n")
	for i := range parts {
		parts[i] = strings.TrimRight(parts[i], " \t")
	}
	for len(parts) > 0 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	canonical := strings.Join(parts, "\n") + "\n"
	sum := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(sum[:])[:8]
}

func asciiAlphaNum(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
}

func matchIDIn(id, value string) bool {
	id = strings.ToLower(id)
	value = strings.ToLower(value)
	cut := len(id)
	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
		cut--
	}
	if cut == len(id) || cut == 0 {
		return false
	}
	prefix := strings.TrimRight(id[:cut], "-_")
	digits := id[cut:]
	for i := 0; i+len(prefix) <= len(value); i++ {
		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
			continue
		}
		j := i + len(prefix)
		if j < len(value) && (value[j] == '-' || value[j] == '_') {
			j++
		}
		if !strings.HasPrefix(value[j:], digits) {
			continue
		}
		j += len(digits)
		if j == len(value) || value[j] < '0' || value[j] > '9' {
			return true
		}
	}
	return false
}

func MatchesID(id, testcaseName, classname string) bool {
	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
}

func matchComponent(pattern, value string) bool {
	p, v := 0, 0
	star, retry := -1, 0
	for v < len(value) {
		if p < len(pattern) && pattern[p] == value[v] {
			p++
			v++
		} else if p < len(pattern) && pattern[p] == '*' {
			star = p
			p++
			retry = v
		} else if star >= 0 {
			p = star + 1
			retry++
			v = retry
		} else {
			return false
		}
	}
	for p < len(pattern) && pattern[p] == '*' {
		p++
	}
	return p == len(pattern)
}

func MatchGlob(pattern, path string) bool {
	pp := strings.Split(pattern, "/")
	parts := strings.Split(path, "/")
	var visit func(int, int) bool
	memo := make(map[[2]int]bool)
	known := make(map[[2]int]bool)
	visit = func(p, i int) bool {
		key := [2]int{p, i}
		if known[key] {
			return memo[key]
		}
		known[key] = true
		result := false
		if p == len(pp) {
			result = i == len(parts)
		} else if pp[p] == "**" {
			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
			result = visit(p+1, i+1)
		}
		memo[key] = result
		return result
	}
	return visit(0, 0)
}

func ResolveZone(path string, zones model.Zones) string {
	for _, pattern := range zones.Protected {
		if MatchGlob(pattern, path) {
			return "protected"
		}
	}
	if path == "SPEC.md" {
		return "spec"
	}
	for _, zone := range []struct {
		name     string
		patterns []string
	}{
		{"evals", zones.Evals},
		{"interface", zones.Interface},
		{"append_only", zones.AppendOnly},
		{"preserve", zones.Preserve},
		{"disposable", zones.Disposable},
	} {
		for _, pattern := range zone.patterns {
			if MatchGlob(pattern, path) {
				return zone.name
			}
		}
	}
	return "other"
}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..5200a8b9ece7c25708de9005ba35c4a7d40b1396
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,434 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+			return "", "", false
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	allLines := lines(source)
+	if len(allLines) == 0 || allLines[0].text != "---" {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	closing := -1
+	for i := 1; i < len(allLines); i++ {
+		if allLines[i].text == "---" {
+			closing = i
+			break
+		}
+	}
+	if closing < 0 {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+	}
+	front := source[allLines[0].end:allLines[closing].start]
+	body := source[allLines[closing].end:]
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	name = strings.TrimSpace(trimmed[i:])
+	j := len(name)
+	for j > 0 && name[j-1] == '#' {
+		j--
+	}
+	if j > 0 && j < len(name) && (name[j-1] == ' ' || name[j-1] == '\t') {
+		name = strings.TrimSpace(name[:j])
+	}
+	return i, name, true
+}
+
+type fence struct {
+	marker byte
+	width  int
+}
+
+// consume reports whether the line belongs to a fenced code block.
+func (f *fence) consume(text string) bool {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return f.marker != 0
+	}
+	if trimmed == "" {
+		return f.marker != 0
+	}
+	marker := trimmed[0]
+	if f.marker != 0 {
+		if marker == f.marker {
+			i := 0
+			for i < len(trimmed) && trimmed[i] == marker {
+				i++
+			}
+			if i >= f.width && strings.TrimSpace(trimmed[i:]) == "" {
+				*f = fence{}
+			}
+		}
+		return true
+	}
+	if marker != '`' && marker != '~' {
+		return false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == marker {
+		i++
+	}
+	if i < 3 {
+		return false
+	}
+	*f = fence{marker: marker, width: i}
+	return true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	var fenced fence
+	for _, l := range lines(markdown) {
+		if fenced.consume(l.text) {
+			continue
+		}
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	bodyStart := 0
+	var fenced fence
+	for _, l := range lines(section) {
+		if fenced.consume(l.text) {
+			continue
+		}
+		candidate := strings.TrimLeft(l.text, " ")
+		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(candidate)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[bodyStart:l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+		bodyStart = l.end
+	}
+	if len(items) > 0 {
+		items[len(items)-1].Body = section[bodyStart:]
+	}
+	return items, nil
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func asciiAlphaNum(ch byte) bool {
+	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) || value[j] < '0' || value[j] > '9' {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..f278cf6ff3e1fe8319df19c335d28c5d734216a4
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,9 @@
+## Contradicted Decisions
+
+None identified. The retained `CodedError` implementation follows decision 0128.
+
+## Gaps
+
+- CON-008, target module `contracts`: The required round trip cannot hold for every possible `map[string]any` value. A nil map is not a YAML mapping, and YAML decoding does not preserve all Go concrete types. `RenderDocument` rejects a nil map and handles YAML-compatible values.
+- CON-010, target module `contracts`: The assigned-ID heading grammar is not fully specified. The implementation accepts an uppercase alphanumeric prefix, a hyphen, and a nonempty decimal suffix, following the retained examples.
+- CON-002, target module `contracts`: The behavior for multiple level-two sections with the same name is unspecified. `Section` and `ReplaceSection` use the first matching section.
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..5200a8b9ece7c25708de9005ba35c4a7d40b1396
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,434 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+			return "", "", false
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	allLines := lines(source)
+	if len(allLines) == 0 || allLines[0].text != "---" {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	closing := -1
+	for i := 1; i < len(allLines); i++ {
+		if allLines[i].text == "---" {
+			closing = i
+			break
+		}
+	}
+	if closing < 0 {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+	}
+	front := source[allLines[0].end:allLines[closing].start]
+	body := source[allLines[closing].end:]
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	name = strings.TrimSpace(trimmed[i:])
+	j := len(name)
+	for j > 0 && name[j-1] == '#' {
+		j--
+	}
+	if j > 0 && j < len(name) && (name[j-1] == ' ' || name[j-1] == '\t') {
+		name = strings.TrimSpace(name[:j])
+	}
+	return i, name, true
+}
+
+type fence struct {
+	marker byte
+	width  int
+}
+
+// consume reports whether the line belongs to a fenced code block.
+func (f *fence) consume(text string) bool {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return f.marker != 0
+	}
+	if trimmed == "" {
+		return f.marker != 0
+	}
+	marker := trimmed[0]
+	if f.marker != 0 {
+		if marker == f.marker {
+			i := 0
+			for i < len(trimmed) && trimmed[i] == marker {
+				i++
+			}
+			if i >= f.width && strings.TrimSpace(trimmed[i:]) == "" {
+				*f = fence{}
+			}
+		}
+		return true
+	}
+	if marker != '`' && marker != '~' {
+		return false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == marker {
+		i++
+	}
+	if i < 3 {
+		return false
+	}
+	*f = fence{marker: marker, width: i}
+	return true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	var fenced fence
+	for _, l := range lines(markdown) {
+		if fenced.consume(l.text) {
+			continue
+		}
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	bodyStart := 0
+	var fenced fence
+	for _, l := range lines(section) {
+		if fenced.consume(l.text) {
+			continue
+		}
+		candidate := strings.TrimLeft(l.text, " ")
+		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(candidate)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[bodyStart:l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+		bodyStart = l.end
+	}
+	if len(items) > 0 {
+		items[len(items)-1].Body = section[bodyStart:]
+	}
+	return items, nil
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func asciiAlphaNum(ch byte) bool {
+	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9'
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if i > 0 && asciiAlphaNum(value[i-1]) || !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) || value[j] < '0' || value[j] > '9' {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

exec
/bin/zsh -lc 'find packages/contracts/internal -type f -print' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc 'gofmt -w packages/contracts/internal/primitives/primitives.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
exec
/bin/zsh -lc 'go test ./...' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
 succeeded in 604ms:
packages/contracts/internal/primitives/primitives.go

 succeeded in 615ms:
 exited 1 in 657ms:
FAIL	./... [setup failed]
# ./...
pattern ./...: directory prefix . does not contain main module or its selected dependencies
FAIL

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..f278cf6ff3e1fe8319df19c335d28c5d734216a4
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,9 @@
+## Contradicted Decisions
+
+None identified. The retained `CodedError` implementation follows decision 0128.
+
+## Gaps
+
+- CON-008, target module `contracts`: The required round trip cannot hold for every possible `map[string]any` value. A nil map is not a YAML mapping, and YAML decoding does not preserve all Go concrete types. `RenderDocument` rejects a nil map and handles YAML-compatible values.
+- CON-010, target module `contracts`: The assigned-ID heading grammar is not fully specified. The implementation accepts an uppercase alphanumeric prefix, a hyphen, and a nonempty decimal suffix, following the retained examples.
+- CON-002, target module `contracts`: The behavior for multiple level-two sections with the same name is unspecified. `Section` and `ReplaceSection` use the first matching section.
diff --git a/packages/contracts/internal/primitives/primitives.go b/packages/contracts/internal/primitives/primitives.go
new file mode 100644
index 0000000000000000000000000000000000000000..fe8bc7f35e9b763fb6f0e4f7406bd14b42f16829
--- /dev/null
+++ b/packages/contracts/internal/primitives/primitives.go
@@ -0,0 +1,442 @@
+package primitives
+
+import (
+	"crypto/sha256"
+	"encoding/hex"
+	"errors"
+	"fmt"
+	"strings"
+	"unicode"
+	"unicode/utf8"
+
+	"github.com/lurkshark/fireweed/packages/contracts/model"
+	"gopkg.in/yaml.v3"
+)
+
+func ErrorCode(err error) (code, subject string, ok bool) {
+	for err != nil {
+		switch e := err.(type) {
+		case model.CodedError:
+			return e.Code, e.Subject, true
+		case *model.CodedError:
+			if e != nil {
+				return e.Code, e.Subject, true
+			}
+			return "", "", false
+		}
+		if many, ok := err.(interface{ Unwrap() []error }); ok {
+			for _, child := range many.Unwrap() {
+				if code, subject, ok := ErrorCode(child); ok {
+					return code, subject, true
+				}
+			}
+			return "", "", false
+		}
+		err = errors.Unwrap(err)
+	}
+	return "", "", false
+}
+
+func ParseDocument(source string) (model.Document, error) {
+	allLines := lines(source)
+	if len(allLines) == 0 || allLines[0].text != "---" {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
+	}
+	closing := -1
+	for i := 1; i < len(allLines); i++ {
+		if allLines[i].text == "---" {
+			closing = i
+			break
+		}
+	}
+	if closing < 0 {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
+	}
+	front := source[allLines[0].end:allLines[closing].start]
+	body := source[allLines[closing].end:]
+	var node yaml.Node
+	if err := yaml.Unmarshal([]byte(front), &node); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	if len(node.Content) != 1 || node.Content[0].Kind != yaml.MappingNode {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("invalid YAML frontmatter: expected a mapping")}
+	}
+	var mapping map[string]any
+	if err := node.Content[0].Decode(&mapping); err != nil {
+		return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
+	}
+	return model.Document{Frontmatter: mapping, Body: body}, nil
+}
+
+func RenderDocument(doc model.Document) (string, error) {
+	if doc.Frontmatter == nil {
+		return "", errors.New("frontmatter must be a mapping")
+	}
+	data, err := yaml.Marshal(doc.Frontmatter)
+	if err != nil {
+		return "", err
+	}
+	return "---\n" + string(data) + "---\n" + doc.Body, nil
+}
+
+type line struct {
+	start, end int
+	text       string
+}
+
+func lines(source string) []line {
+	var out []line
+	for start := 0; start < len(source); {
+		next := strings.IndexByte(source[start:], '\n')
+		end := len(source)
+		if next >= 0 {
+			end = start + next + 1
+		}
+		text := source[start:end]
+		text = strings.TrimSuffix(text, "\n")
+		text = strings.TrimSuffix(text, "\r")
+		out = append(out, line{start: start, end: end, text: text})
+		start = end
+	}
+	return out
+}
+
+func heading(text string) (level int, name string, ok bool) {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return 0, "", false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == '#' {
+		i++
+	}
+	if i < 1 || i > 6 || (i < len(trimmed) && trimmed[i] != ' ' && trimmed[i] != '\t') {
+		return 0, "", false
+	}
+	name = strings.TrimSpace(trimmed[i:])
+	j := len(name)
+	for j > 0 && name[j-1] == '#' {
+		j--
+	}
+	if j > 0 && j < len(name) && (name[j-1] == ' ' || name[j-1] == '\t') {
+		name = strings.TrimSpace(name[:j])
+	}
+	return i, name, true
+}
+
+type fence struct {
+	marker byte
+	width  int
+}
+
+// consume reports whether the line belongs to a fenced code block.
+func (f *fence) consume(text string) bool {
+	trimmed := strings.TrimLeft(text, " ")
+	if len(text)-len(trimmed) > 3 {
+		return f.marker != 0
+	}
+	if trimmed == "" {
+		return f.marker != 0
+	}
+	marker := trimmed[0]
+	if f.marker != 0 {
+		if marker == f.marker {
+			i := 0
+			for i < len(trimmed) && trimmed[i] == marker {
+				i++
+			}
+			if i >= f.width && strings.TrimSpace(trimmed[i:]) == "" {
+				*f = fence{}
+			}
+		}
+		return true
+	}
+	if marker != '`' && marker != '~' {
+		return false
+	}
+	i := 0
+	for i < len(trimmed) && trimmed[i] == marker {
+		i++
+	}
+	if i < 3 {
+		return false
+	}
+	*f = fence{marker: marker, width: i}
+	return true
+}
+
+func sectionBounds(markdown, name string) (bodyStart, bodyEnd int, found bool) {
+	start := -1
+	var fenced fence
+	for _, l := range lines(markdown) {
+		if fenced.consume(l.text) {
+			continue
+		}
+		level, title, ok := heading(l.text)
+		if !ok {
+			continue
+		}
+		if start >= 0 && level <= 2 {
+			return start, l.start, true
+		}
+		if start < 0 && level == 2 && title == name {
+			start = l.end
+		}
+	}
+	if start >= 0 {
+		return start, len(markdown), true
+	}
+	return 0, 0, false
+}
+
+func Section(markdown, name string) (string, bool) {
+	start, end, ok := sectionBounds(markdown, name)
+	if !ok {
+		return "", false
+	}
+	return markdown[start:end], true
+}
+
+func ReplaceSection(markdown, name, body string) string {
+	start, end, ok := sectionBounds(markdown, name)
+	if ok {
+		return markdown[:start] + body + markdown[end:]
+	}
+	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
+		markdown += "\n"
+	}
+	return markdown + "## " + name + "\n" + body
+}
+
+func validItemID(id string) bool {
+	if strings.HasPrefix(id, "new:") {
+		slug := strings.TrimPrefix(id, "new:")
+		if slug == "" {
+			return false
+		}
+		for _, ch := range slug {
+			if !(ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9' || ch == '-') {
+				return false
+			}
+		}
+		return true
+	}
+	parts := strings.Split(id, "-")
+	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
+		return false
+	}
+	for _, ch := range parts[0] {
+		if !(ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9') {
+			return false
+		}
+	}
+	for _, ch := range parts[1] {
+		if ch < '0' || ch > '9' {
+			return false
+		}
+	}
+	return true
+}
+
+func parseItemHeading(text string) (id, tag string, ok bool) {
+	if !strings.HasPrefix(text, "### ") {
+		return "", "", false
+	}
+	content := strings.TrimPrefix(text, "### ")
+	parts := strings.SplitN(content, " ", 2)
+	id = parts[0]
+	if !validItemID(id) {
+		return "", "", false
+	}
+	if len(parts) == 1 {
+		return id, "", true
+	}
+	remaining := parts[1]
+	if len(remaining) < 3 || remaining[0] != '[' || remaining[len(remaining)-1] != ']' {
+		return "", "", false
+	}
+	tag = remaining[1 : len(remaining)-1]
+	if strings.ContainsAny(tag, "[]") || strings.TrimSpace(tag) == "" {
+		return "", "", false
+	}
+	return id, tag, true
+}
+
+func ParseItems(section string) ([]model.Item, error) {
+	var items []model.Item
+	seen := make(map[string]bool)
+	bodyStart := 0
+	var fenced fence
+	for _, l := range lines(section) {
+		if fenced.consume(l.text) {
+			continue
+		}
+		candidate := strings.TrimLeft(l.text, " ")
+		if len(l.text)-len(candidate) > 3 || !strings.HasPrefix(candidate, "###") || strings.HasPrefix(candidate, "####") {
+			continue
+		}
+		id, tag, ok := parseItemHeading(candidate)
+		if !ok {
+			return nil, model.CodedError{Code: "C03", Cause: fmt.Errorf("malformed item heading: %q", l.text)}
+		}
+		if seen[id] {
+			return nil, model.CodedError{Code: "C05", Subject: id, Cause: fmt.Errorf("duplicate item ID: %s", id)}
+		}
+		seen[id] = true
+		if len(items) > 0 {
+			items[len(items)-1].Body = section[bodyStart:l.start]
+		}
+		items = append(items, model.Item{ID: id, Tag: tag})
+		bodyStart = l.end
+	}
+	if len(items) > 0 {
+		items[len(items)-1].Body = section[bodyStart:]
+	}
+	return items, nil
+}
+
+func RenderItem(item model.Item) string {
+	heading := "### " + item.ID
+	if item.Tag != "" {
+		heading += " [" + item.Tag + "]"
+	}
+	return heading + "\n" + item.Body
+}
+
+func SpecHash(item model.Item) string {
+	value := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
+	parts := strings.Split(value, "\n")
+	for i := range parts {
+		parts[i] = strings.TrimRight(parts[i], " \t")
+	}
+	for len(parts) > 0 && parts[len(parts)-1] == "" {
+		parts = parts[:len(parts)-1]
+	}
+	canonical := strings.Join(parts, "\n") + "\n"
+	sum := sha256.Sum256([]byte(canonical))
+	return hex.EncodeToString(sum[:])[:8]
+}
+
+func matchIDIn(id, value string) bool {
+	id = strings.ToLower(id)
+	value = strings.ToLower(value)
+	cut := len(id)
+	for cut > 0 && id[cut-1] >= '0' && id[cut-1] <= '9' {
+		cut--
+	}
+	if cut == len(id) || cut == 0 {
+		return false
+	}
+	prefix := strings.TrimRight(id[:cut], "-_")
+	digits := id[cut:]
+	for i := 0; i+len(prefix) <= len(value); i++ {
+		if !strings.HasPrefix(value[i:], prefix) {
+			continue
+		}
+		if i > 0 {
+			prior, _ := utf8.DecodeLastRuneInString(value[:i])
+			if unicode.IsLetter(prior) || unicode.IsDigit(prior) {
+				continue
+			}
+		}
+		j := i + len(prefix)
+		if j < len(value) && (value[j] == '-' || value[j] == '_') {
+			j++
+		}
+		if !strings.HasPrefix(value[j:], digits) {
+			continue
+		}
+		j += len(digits)
+		if j == len(value) {
+			return true
+		}
+		next, _ := utf8.DecodeRuneInString(value[j:])
+		if !unicode.IsDigit(next) {
+			return true
+		}
+	}
+	return false
+}
+
+func MatchesID(id, testcaseName, classname string) bool {
+	return matchIDIn(id, testcaseName) || matchIDIn(id, classname)
+}
+
+func matchComponent(pattern, value string) bool {
+	p, v := 0, 0
+	star, retry := -1, 0
+	for v < len(value) {
+		if p < len(pattern) && pattern[p] == value[v] {
+			p++
+			v++
+		} else if p < len(pattern) && pattern[p] == '*' {
+			star = p
+			p++
+			retry = v
+		} else if star >= 0 {
+			p = star + 1
+			retry++
+			v = retry
+		} else {
+			return false
+		}
+	}
+	for p < len(pattern) && pattern[p] == '*' {
+		p++
+	}
+	return p == len(pattern)
+}
+
+func MatchGlob(pattern, path string) bool {
+	pp := strings.Split(pattern, "/")
+	parts := strings.Split(path, "/")
+	var visit func(int, int) bool
+	memo := make(map[[2]int]bool)
+	known := make(map[[2]int]bool)
+	visit = func(p, i int) bool {
+		key := [2]int{p, i}
+		if known[key] {
+			return memo[key]
+		}
+		known[key] = true
+		result := false
+		if p == len(pp) {
+			result = i == len(parts)
+		} else if pp[p] == "**" {
+			result = visit(p+1, i) || i < len(parts) && visit(p, i+1)
+		} else if i < len(parts) && matchComponent(pp[p], parts[i]) {
+			result = visit(p+1, i+1)
+		}
+		memo[key] = result
+		return result
+	}
+	return visit(0, 0)
+}
+
+func ResolveZone(path string, zones model.Zones) string {
+	for _, pattern := range zones.Protected {
+		if MatchGlob(pattern, path) {
+			return "protected"
+		}
+	}
+	if path == "SPEC.md" {
+		return "spec"
+	}
+	for _, zone := range []struct {
+		name     string
+		patterns []string
+	}{
+		{"evals", zones.Evals},
+		{"interface", zones.Interface},
+		{"append_only", zones.AppendOnly},
+		{"preserve", zones.Preserve},
+		{"disposable", zones.Disposable},
+	} {
+		for _, pattern := range zone.patterns {
+			if MatchGlob(pattern, path) {
+				return zone.name
+			}
+		}
+	}
+	return "other"
+}

exec
/bin/zsh -lc "gofmt -w packages/contracts/internal/primitives/primitives.go && rg -n 'TODO|itemsBodyStart|panic\\(' packages/contracts/internal/primitives DRILL_NOTES.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570
 exited 1 in 505ms:
codex
Implemented the contracts primitives in [primitives.go](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/packages/contracts/internal/primitives/primitives.go) and recorded contradictions and gaps in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1341965570/DRILL_NOTES.md). Retained inputs were left untouched, and nothing was committed.

`gofmt` passed. I could not run the evals here: `go test ./...` stops because the supplied workspace has no Go module manifest.
tokens used
33,412

