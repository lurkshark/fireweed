---
date: 2026-11-01T17:04:11.628421Z
base_commit: 89d7dbfb23b788d3efc98ba8c7709291420d8438
agent: codex
model: gpt-6-sol
result: pass
---

## Own evals
- gitstore | Test_GIT_001_InspectWorktreesAndDetachedHead | GIT-001 | pass
- gitstore | Test_GIT_002_BaseRefPreferenceDefaultsAndMissing | GIT-002 | pass
- gitstore | Test_GIT_003_GIT_006_GIT_007_ChangesReadAtAndSelectedCommit | GIT-003, GIT-006, GIT-007 | pass
- gitstore | Test_GIT_003_UnstagedTrackedRename | GIT-003 | pass
- gitstore | Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS | GIT-004, GIT-008, GIT-009, GIT-010 | pass
- gitstore | Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS/StaleUpdatePreservesNewerRef | GIT-004, GIT-008, GIT-009, GIT-010 | pass
- gitstore | Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel | GIT-005 | pass
- gitstore | Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel/TimeoutTerminatesGroup | GIT-005 | pass
- gitstore | Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel/ContextCancellationTerminatesGroup | GIT-005 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/gitstore	[no test files]
?   	github.com/lurkshark/fireweed/packages/gitstore/internal/local	[no test files]
?   	github.com/lurkshark/fireweed/packages/gitstore/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/gitstore/evals	5.354s


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
ok  	github.com/lurkshark/fireweed/packages/app/evals	252.089s

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
ok  	github.com/lurkshark/fireweed/packages/checks/evals	227.988s

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
ok  	github.com/lurkshark/fireweed/packages/project/evals	1.771s

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
ok  	github.com/lurkshark/fireweed/packages/runner/evals	265.293s

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
ok  	github.com/lurkshark/fireweed/packages/workflow/evals	121.823s


## Contradicted Decisions
None identified.
## Gaps
None identified.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt gitstore in [packages/gitstore/internal/local](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local). The retained evals pass. [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/DRILL_NOTES.md) records no identified contradictions or gaps. Retained inputs are unchanged, and nothing was committed.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a19811-99ac-7f71-893e-fbfa28763300
--------
user
# Rebuild gitstore

Target spec and Decisions:
---
depends_on:
    - contracts
id_prefix: GIT
interface:
    - api.go
    - model/**
name: gitstore
---

# gitstore

## Specs

### GIT-001
Inspect MUST return canonical absolute worktree/common-Git directories plus HEAD and branch, from a directory inside a primary or linked worktree. The branch MUST be empty for detached HEAD, and non-repositories MUST return an error.

### GIT-002
BaseRef MUST prefer a locally known remote/base ref, falling back to the local base branch and rejecting a missing base. An omitted remote MUST mean origin. An omitted base MUST use the local remote HEAD's branch when known, otherwise main.

### GIT-003
Changes MUST report repository-relative staged and unstaged tracked edits and untracked paths against the given base, once per reported Path. Renames reported by either tracked diff MUST retain both names.

### GIT-004
WriteBlob MUST store a byte payload in local Git objects and return its object ID.

### GIT-005
Run MUST launch argv in the requested directory with the supplied environment/stdin and return captured stdout, stderr and exit code. A nonzero exit MUST return an error while preserving output. Timeout or context cancellation MUST terminate the launched process group and return an error; processes escaping that group are outside the supported contract.

### GIT-006
ReadAt MUST return a revision's file bytes or an error for absent paths.

### GIT-007
Commit MUST commit the worktree contents of explicitly selected paths, including new files, with the given subject and return the new HEAD. It MUST leave unrelated working and index changes uncommitted.

### GIT-008
ReadBlob MUST return the bytes of a local Git blob object by its object ID.

### GIT-009
Ref MUST return the named local ref's object ID and existence.

### GIT-010
CASRef MUST create, update or delete a ref only when its exact current object matches the expected old value, returning an error without overwriting a newer value on mismatch.

## Decisions

[0010-manual-lifecycle-git] Change.Status uses the leading letter of Git's name-status code for tracked changes; untracked paths use `?`.


Retained interfaces and files:
packages/gitstore/SPEC.md
packages/gitstore/api.go
packages/gitstore/model/model.go

Retained evals:
packages/gitstore/evals/.gitkeep
packages/gitstore/evals/gitstore_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.

codex
I’ll inspect the retained gitstore interfaces and the target module layout, then implement the disposable paths and run the retained evals. I’ll leave the retained inputs untouched and record spec gaps in DRILL_NOTES.md.
exec
/bin/zsh -lc "pwd && rg --files -g 'AGENTS.md' -g 'go.mod' -g 'packages/gitstore/**' -g '"'!packages/gitstore/evals/.gitkeep'"' -g '"'!**/.git/**'"' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 481ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
packages/gitstore/SPEC.md
packages/gitstore/api.go
packages/gitstore/evals/gitstore_test.go
packages/gitstore/model/model.go

exec
/bin/zsh -lc 'ls -la && ls -la packages/gitstore && find packages/gitstore -maxdepth 2 -type f -print | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc 'cat packages/gitstore/SPEC.md packages/gitstore/api.go packages/gitstore/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc 'cat packages/gitstore/evals/gitstore_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc 'find . -name AGENTS.md -print' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 673ms:
package evals_test

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

func git(t *testing.T, dir string, args ...string) string {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", append([]string{"-C", dir}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_CONFIG_GLOBAL=/dev/null", "GIT_NO_LAZY_FETCH=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func put(t *testing.T, root, name string, data []byte) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}
}

func repo(t *testing.T) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "repo")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	var err error
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	git(t, root, "init", "-q", "-b", "main", "--object-format=sha1")
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
	put(t, root, "seed.txt", []byte("seed\n"))
	git(t, root, "add", ".")
	git(t, root, "commit", "-qm", "seed")
	return root
}

func head(t *testing.T, root string) string { return git(t, root, "rev-parse", "HEAD") }

func Test_GIT_001_InspectWorktreesAndDetachedHead(t *testing.T) {
	root := repo(t)
	inside := filepath.Join(root, "nested")
	if err := os.Mkdir(inside, 0755); err != nil {
		t.Fatal(err)
	}
	primary, err := gitstore.Inspect(inside)
	if err != nil {
		t.Fatal(err)
	}
	common, err := filepath.EvalSymlinks(filepath.Join(root, ".git"))
	if err != nil {
		t.Fatal(err)
	}
	if primary.Root != root || primary.CommonDir != common || primary.Head != head(t, root) || primary.Branch != "main" {
		t.Errorf("primary = %#v", primary)
	}
	linkedRoot := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-qb", "linked", linkedRoot)
	linkedRoot, err = filepath.EvalSymlinks(linkedRoot)
	if err != nil {
		t.Fatal(err)
	}
	linked, err := gitstore.Inspect(linkedRoot)
	if err != nil {
		t.Fatal(err)
	}
	if linked.Root != linkedRoot || linked.CommonDir != common || linked.Head != head(t, root) || linked.Branch != "linked" {
		t.Errorf("linked = %#v", linked)
	}
	git(t, linkedRoot, "checkout", "--detach", "-q")
	detached, err := gitstore.Inspect(linkedRoot)
	if err != nil {
		t.Fatal(err)
	}
	if detached.Root != linkedRoot || detached.CommonDir != common || detached.Head != head(t, root) || detached.Branch != "" {
		t.Errorf("detached = %#v", detached)
	}
	if _, err := gitstore.Inspect(t.TempDir()); err == nil {
		t.Error("Inspect(non-repository) succeeded")
	}
}

func Test_GIT_002_BaseRefPreferenceDefaultsAndMissing(t *testing.T) {
	root := repo(t)
	if got, err := gitstore.BaseRef(root, "", ""); err != nil || got != "refs/heads/main" {
		t.Errorf("default main = %q, %v", got, err)
	}
	git(t, root, "branch", "release")
	if got, err := gitstore.BaseRef(root, "upstream", "release"); err != nil || got != "refs/heads/release" {
		t.Errorf("local fallback = %q, %v", got, err)
	}
	if _, err := gitstore.BaseRef(root, "upstream", "missing"); err == nil {
		t.Error("missing base succeeded")
	}
	oid := head(t, root)
	git(t, root, "update-ref", "refs/remotes/origin/main", oid)
	if got, err := gitstore.BaseRef(root, "", "main"); err != nil || got != "refs/remotes/origin/main" {
		t.Errorf("origin preference = %q, %v", got, err)
	}
	git(t, root, "update-ref", "refs/remotes/upstream/release", oid)
	if got, err := gitstore.BaseRef(root, "upstream", "release"); err != nil || got != "refs/remotes/upstream/release" {
		t.Errorf("named remote preference = %q, %v", got, err)
	}
	git(t, root, "update-ref", "refs/remotes/origin/develop", oid)
	git(t, root, "symbolic-ref", "refs/remotes/origin/HEAD", "refs/remotes/origin/develop")
	if got, err := gitstore.BaseRef(root, "", ""); err != nil || got != "refs/remotes/origin/develop" {
		t.Errorf("remote HEAD default = %q, %v", got, err)
	}
	git(t, root, "symbolic-ref", "refs/remotes/upstream/HEAD", "refs/remotes/upstream/release")
	if got, err := gitstore.BaseRef(root, "upstream", ""); err != nil || got != "refs/remotes/upstream/release" {
		t.Errorf("named remote HEAD default = %q, %v", got, err)
	}
}

func Test_GIT_003_GIT_006_GIT_007_ChangesReadAtAndSelectedCommit(t *testing.T) {
	root := repo(t)
	for _, name := range []string{"selected.txt", "staged.txt", "unstaged.txt", "old-name.txt"} {
		put(t, root, name, []byte("original\n"))
	}
	git(t, root, "add", ".")
	git(t, root, "commit", "-qm", "files")
	base := head(t, root)
	put(t, root, "committed.txt", []byte("committed after base\n"))
	git(t, root, "add", "committed.txt")
	git(t, root, "commit", "-qm", "after base")
	put(t, root, "selected.txt", []byte("selected\n"))
	put(t, root, "staged.txt", []byte("staged\n"))
	git(t, root, "add", "selected.txt", "staged.txt")
	put(t, root, "staged.txt", []byte("staged worktree\n"))
	put(t, root, "unstaged.txt", []byte("unstaged\n"))
	git(t, root, "mv", "old-name.txt", "new-name.txt")
	put(t, root, "new-name.txt", []byte("renamed worktree\n"))
	put(t, root, "untracked.txt", []byte("new\n"))
	changes, err := gitstore.Changes(root, base)
	if err != nil {
		t.Fatal(err)
	}
	byPath := make(map[string]gitstore.Change)
	for _, change := range changes {
		if change.Status == "" || filepath.IsAbs(change.Path) {
			t.Errorf("invalid change: %#v", change)
		}
		byPath[change.Path] = change
	}
	if len(changes) != len(byPath) {
		t.Errorf("duplicate reported paths: %#v", changes)
	}
	for _, path := range []string{"committed.txt", "selected.txt", "staged.txt", "unstaged.txt", "untracked.txt", "new-name.txt"} {
		if _, ok := byPath[path]; !ok {
			t.Errorf("missing change %q: %#v", path, changes)
		}
	}
	if got := byPath["new-name.txt"].OldPath; got != "old-name.txt" {
		t.Errorf("rename OldPath = %q", got)
	}
	data, err := gitstore.ReadAt(root, base, "selected.txt")
	if err != nil || !bytes.Equal(data, []byte("original\n")) {
		t.Errorf("ReadAt(base) = %q, %v", data, err)
	}
	if _, err := gitstore.ReadAt(root, base, "untracked.txt"); err == nil {
		t.Error("ReadAt(absent path) succeeded")
	}
	gotHead, err := gitstore.Commit(root, "selected only", []string{"selected.txt"})
	if err != nil || gotHead == "" || gotHead != head(t, root) {
		t.Fatalf("Commit = %q, %v", gotHead, err)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", gotHead); got != "selected.txt" {
		t.Errorf("committed paths = %q, want selected.txt", got)
	}
	if got := git(t, root, "log", "-1", "--format=%s"); got != "selected only" {
		t.Errorf("subject = %q", got)
	}
	if got := git(t, root, "show", "HEAD:selected.txt"); got != "selected" {
		t.Errorf("selected HEAD content = %q", got)
	}
	if got := git(t, root, "show", "HEAD:staged.txt"); got != "original" {
		t.Errorf("unrelated staged file was committed: %q", got)
	}
	if got := git(t, root, "show", ":staged.txt"); got != "staged" {
		t.Errorf("unrelated staged index content changed: %q", got)
	}
	if got := git(t, root, "diff", "--cached", "--name-only"); !strings.Contains(got, "staged.txt") || !strings.Contains(got, "new-name.txt") {
		t.Errorf("unrelated index changes lost: %q", got)
	}
	if data, err := os.ReadFile(filepath.Join(root, "unstaged.txt")); err != nil || string(data) != "unstaged\n" {
		t.Errorf("unrelated worktree change lost: %q, %v", data, err)
	}
	put(t, root, "selected.txt", []byte("selected index\n"))
	git(t, root, "add", "selected.txt")
	put(t, root, "selected.txt", []byte("selected worktree\n"))
	put(t, root, "new-selected.txt", []byte("new selected\n"))
	gotHead, err = gitstore.Commit(root, "selected worktree and new file", []string{"selected.txt", "new-selected.txt"})
	if err != nil || gotHead == "" || gotHead != head(t, root) {
		t.Fatalf("Commit selected worktree and new file = %q, %v", gotHead, err)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", gotHead); got != "new-selected.txt\nselected.txt" {
		t.Errorf("committed selected paths = %q, want new-selected.txt and selected.txt", got)
	}
	if got := git(t, root, "log", "-1", "--format=%s"); got != "selected worktree and new file" {
		t.Errorf("selected commit subject = %q", got)
	}
	if got := git(t, root, "show", "HEAD:selected.txt"); got != "selected worktree" {
		t.Errorf("selected worktree HEAD content = %q", got)
	}
	if got := git(t, root, "show", "HEAD:new-selected.txt"); got != "new selected" {
		t.Errorf("new selected HEAD content = %q", got)
	}
	if got := git(t, root, "show", "HEAD:staged.txt"); got != "original" {
		t.Errorf("unrelated staged file was committed: %q", got)
	}
	if got := git(t, root, "show", ":staged.txt"); got != "staged" {
		t.Errorf("unrelated staged index content changed: %q", got)
	}
	if got := git(t, root, "diff", "--cached", "--name-only"); !strings.Contains(got, "staged.txt") || !strings.Contains(got, "new-name.txt") {
		t.Errorf("unrelated index changes lost: %q", got)
	}
	if data, err := os.ReadFile(filepath.Join(root, "unstaged.txt")); err != nil || string(data) != "unstaged\n" {
		t.Errorf("unrelated worktree change lost: %q, %v", data, err)
	}
}

func Test_GIT_003_UnstagedTrackedRename(t *testing.T) {
	root := repo(t)
	base := head(t, root)
	if err := os.Rename(filepath.Join(root, "seed.txt"), filepath.Join(root, "renamed-seed.txt")); err != nil {
		t.Fatal(err)
	}
	git(t, root, "add", "-N", "renamed-seed.txt")
	changes, err := gitstore.Changes(root, base)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	for _, change := range changes {
		if change.Path == "renamed-seed.txt" {
			count++
			if change.OldPath != "seed.txt" {
				t.Errorf("unstaged rename = %#v, want old path seed.txt", change)
			}
		}
	}
	if count != 1 {
		t.Errorf("renamed-seed.txt reported %d times: %#v", count, changes)
	}
}

func Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS(t *testing.T) {
	root := repo(t)
	payload := []byte{0, 'b', '\n', 255}
	oid, err := gitstore.WriteBlob(root, payload)
	if err != nil || !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) {
		t.Fatalf("WriteBlob = %q, %v", oid, err)
	}
	if got, err := gitstore.ReadBlob(root, oid); err != nil || !bytes.Equal(got, payload) {
		t.Errorf("ReadBlob = %v, %v", got, err)
	}
	name := "refs/fireweed/eval-claim"
	if got, exists, err := gitstore.Ref(root, name); err != nil || exists || got != "" {
		t.Errorf("absent Ref = %q, %v, %v", got, exists, err)
	}
	first := head(t, root)
	git(t, root, "commit", "-q", "--allow-empty", "-m", "second")
	second := head(t, root)
	git(t, root, "commit", "-q", "--allow-empty", "-m", "third")
	third := head(t, root)
	if err := gitstore.CASRef(root, name, "", first); err != nil {
		t.Fatal(err)
	}
	if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != first {
		t.Errorf("created Ref = %q, %v, %v", got, exists, err)
	}
	if err := gitstore.CASRef(root, name, "", second); err == nil {
		t.Error("create-only overwrote existing ref")
	}
	if err := gitstore.CASRef(root, name, first, second); err != nil {
		t.Fatal(err)
	}
	t.Run("StaleUpdatePreservesNewerRef", func(t *testing.T) {
		if err := gitstore.CASRef(root, name, first, third); err == nil {
			t.Error("stale update succeeded")
		}
		if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != second {
			t.Errorf("stale update changed Ref = %q, %v, %v", got, exists, err)
		}
	})
	if err := gitstore.CASRef(root, name, first, ""); err == nil {
		t.Error("stale delete succeeded")
	}
	if got, exists, err := gitstore.Ref(root, name); err != nil || !exists || got != second {
		t.Errorf("stale operation changed Ref = %q, %v, %v", got, exists, err)
	}
	if err := gitstore.CASRef(root, name, second, ""); err != nil {
		t.Fatal(err)
	}
	if got, exists, err := gitstore.Ref(root, name); err != nil || exists || got != "" {
		t.Errorf("deleted Ref = %q, %v, %v", got, exists, err)
	}
}

func Test_GIT_005_RunCaptureExitTimeoutAndGroupCancel(t *testing.T) {
	dir := t.TempDir()
	var err error
	dir, err = filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"/bin/sh", "-c", `printf '%s|%s|' "$PWD" "$FIREWEED_EVAL"; cat; printf 'warn' >&2`},
		Dir:  dir, Env: map[string]string{"FIREWEED_EVAL": "supplied"}, Input: []byte("input"), Timeout: 3 * time.Second,
	})
	if err != nil || result.Stdout != dir+"|supplied|input" || result.Stderr != "warn" || result.ExitCode != 0 {
		t.Errorf("successful Run = %#v, %v", result, err)
	}
	result, err = gitstore.Run(context.Background(), gitstore.Request{
		Args: []string{"/bin/sh", "-c", `printf 'out'; printf 'err' >&2; exit 7`}, Dir: dir, Timeout: 3 * time.Second,
	})
	if err == nil || result.Stdout != "out" || result.Stderr != "err" || result.ExitCode != 7 {
		t.Errorf("failed Run = %#v, %v", result, err)
	}
	for _, mode := range []struct {
		name    string
		timeout time.Duration
		cancel  bool
	}{
		{name: "TimeoutTerminatesGroup", timeout: 500 * time.Millisecond},
		{name: "ContextCancellationTerminatesGroup", timeout: 3 * time.Second, cancel: true},
	} {
		t.Run(mode.name, func(t *testing.T) {
			ready := filepath.Join(dir, mode.name+"-ready")
			marker := filepath.Join(dir, mode.name+"-child-marker")
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			done := make(chan error, 1)
			start := time.Now()
			go func() {
				_, runErr := gitstore.Run(ctx, gitstore.Request{
					Args: []string{"/bin/sh", "-c", `(sleep 1; printf child > "$2") & printf ready > "$1"; wait`, "sh", ready, marker},
					Dir:  dir, Timeout: mode.timeout,
				})
				done <- runErr
			}()
			deadline := time.After(2 * time.Second)
			for {
				if _, statErr := os.Stat(ready); statErr == nil {
					break
				}
				select {
				case runErr := <-done:
					t.Fatalf("Run exited before child was ready: %v", runErr)
				case <-deadline:
					t.Fatal("child did not become ready")
				case <-time.After(10 * time.Millisecond):
				}
			}
			if mode.cancel {
				cancel()
			}
			select {
			case runErr := <-done:
				if runErr == nil {
					t.Error("terminated Run succeeded")
				}
			case <-time.After(2 * time.Second):
				t.Error("terminated Run did not return promptly")
			}
			if elapsed := time.Since(start); elapsed > 3*time.Second {
				t.Errorf("terminated Run took %s", elapsed)
			}
			time.Sleep(1200 * time.Millisecond)
			if _, statErr := os.Stat(marker); !os.IsNotExist(statErr) {
				t.Errorf("child survived group termination: marker stat = %v", statErr)
			}
		})
	}
}

 succeeded in 676ms:
---
depends_on:
    - contracts
id_prefix: GIT
interface:
    - api.go
    - model/**
name: gitstore
---

# gitstore

## Specs

### GIT-001
Inspect MUST return canonical absolute worktree/common-Git directories plus HEAD and branch, from a directory inside a primary or linked worktree. The branch MUST be empty for detached HEAD, and non-repositories MUST return an error.

### GIT-002
BaseRef MUST prefer a locally known remote/base ref, falling back to the local base branch and rejecting a missing base. An omitted remote MUST mean origin. An omitted base MUST use the local remote HEAD's branch when known, otherwise main.

### GIT-003
Changes MUST report repository-relative staged and unstaged tracked edits and untracked paths against the given base, once per reported Path. Renames reported by either tracked diff MUST retain both names.

### GIT-004
WriteBlob MUST store a byte payload in local Git objects and return its object ID.

### GIT-005
Run MUST launch argv in the requested directory with the supplied environment/stdin and return captured stdout, stderr and exit code. A nonzero exit MUST return an error while preserving output. Timeout or context cancellation MUST terminate the launched process group and return an error; processes escaping that group are outside the supported contract.

### GIT-006
ReadAt MUST return a revision's file bytes or an error for absent paths.

### GIT-007
Commit MUST commit the worktree contents of explicitly selected paths, including new files, with the given subject and return the new HEAD. It MUST leave unrelated working and index changes uncommitted.

### GIT-008
ReadBlob MUST return the bytes of a local Git blob object by its object ID.

### GIT-009
Ref MUST return the named local ref's object ID and existence.

### GIT-010
CASRef MUST create, update or delete a ref only when its exact current object matches the expected old value, returning an error without overwriting a newer value on mismatch.

## Decisions

[0010-manual-lifecycle-git] Change.Status uses the leading letter of Git's name-status code for tracked changes; untracked paths use `?`.
package gitstore

import (
	"context"

	"github.com/lurkshark/fireweed/packages/gitstore/internal/local"
	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

type Repository = model.Repository
type Change = model.Change
type Request = model.Request
type Result = model.Result

func Inspect(dir string) (Repository, error)             { return local.Inspect(dir) }
func BaseRef(root, remote, base string) (string, error)  { return local.BaseRef(root, remote, base) }
func Changes(root, base string) ([]Change, error)        { return local.Changes(root, base) }
func ReadAt(root, revision, path string) ([]byte, error) { return local.ReadAt(root, revision, path) }
func Commit(root, subject string, paths []string) (string, error) {
	return local.Commit(root, subject, paths)
}
func WriteBlob(root string, data []byte) (string, error)       { return local.WriteBlob(root, data) }
func ReadBlob(root, oid string) ([]byte, error)                { return local.ReadBlob(root, oid) }
func Ref(root, name string) (string, bool, error)              { return local.Ref(root, name) }
func CASRef(root, name, old, new string) error                 { return local.CASRef(root, name, old, new) }
func Run(ctx context.Context, request Request) (Result, error) { return local.Run(ctx, request) }
package model

import "time"

type Repository struct {
	Root      string
	CommonDir string
	Head      string
	Branch    string
}

type Change struct {
	Path    string
	OldPath string
	Status  string
}

type Request struct {
	Args    []string
	Dir     string
	Env     map[string]string
	Input   []byte
	Timeout time.Duration
}

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

 succeeded in 671ms:
 succeeded in 685ms:
total 8
drwx------@    5 joel  staff     160 Nov  1 09:04 .
drwx------@ 6804 joel  staff  217728 Nov  1 09:04 ..
drwxr-xr-x@    9 joel  staff     288 Nov  1 09:04 .git
-rw-r--r--@    1 joel  staff    1613 Nov  1 09:04 SYSTEM.md
drwxr-xr-x@    4 joel  staff     128 Nov  1 09:04 packages
total 16
drwxr-xr-x@ 6 joel  staff   192 Nov  1 09:04 .
drwxr-xr-x@ 4 joel  staff   128 Nov  1 09:04 ..
-rw-r--r--@ 1 joel  staff  2121 Nov  1 09:04 SPEC.md
-rw-r--r--@ 1 joel  staff  1288 Nov  1 09:04 api.go
drwxr-xr-x@ 4 joel  staff   128 Nov  1 09:04 evals
drwxr-xr-x@ 3 joel  staff    96 Nov  1 09:04 model
packages/gitstore/SPEC.md
packages/gitstore/api.go
packages/gitstore/evals/.gitkeep
packages/gitstore/evals/gitstore_test.go
packages/gitstore/model/model.go

exec
/bin/zsh -lc 'ls -la packages && git status --short && git rev-parse --show-toplevel' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc 'cat SYSTEM.md && find packages -maxdepth 3 -type f -print | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 520ms:
# Independent gitstore reconstruction

Reconstruct only gitstore from the supplied SPEC.md (including Decisions),
retained public interfaces/model and evals. Its only dependency is contracts;
only contracts' declared interfaces are supplied. Retained files are authority.

Stay inside the supplied input workspace. Do not read external trees, old
implementation, source history, saved attempts, binaries, bootstrap, embedded
product content or another module's internals. Do not launch nested agents.
Do not run any product executable to recover target behavior or content.
The fresh input Git repository contains only supplied inputs; do not seek
another repository or remote.

Write only packages/gitstore/internal/** and DRILL_NOTES.md. Leave all retained
inputs unchanged. Do not commit. Use ordinary Go source and gofmt. Workers use
GOCACHE=/tmp/fireweed-rebuild-cache and the current tree's .fireweed/bin first
on PATH. No worker controller invocation is needed in this input workspace.

The full project's build manifests, hooks and dependent implementations stay
in a separate verifier checkout. Their absence here is deliberate: do not
search for them or create files outside the allowed paths. The controller
compiles and runs own and all transitive dependent evals and supplies bounded
failure feedback. Do not change frozen evals/interfaces/specs to make them pass.

Record separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md.
For each actual gap, name the nearest spec ID (or "no spec") and target module.
Report contradictions and uncertainty honestly; this run authorizes no fixes.
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/SPEC.md
packages/gitstore/api.go
packages/gitstore/evals/.gitkeep
packages/gitstore/evals/gitstore_test.go
packages/gitstore/model/model.go

 succeeded in 553ms:
total 0
drwxr-xr-x@ 4 joel  staff  128 Nov  1 09:04 .
drwx------@ 5 joel  staff  160 Nov  1 09:04 ..
drwxr-xr-x@ 4 joel  staff  128 Nov  1 09:04 contracts
drwxr-xr-x@ 6 joel  staff  192 Nov  1 09:04 gitstore
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844

codex
The public API delegates to a missing `internal/local` package. I’m implementing the Git operations there, including selective commits that preserve unrelated index entries and process group termination for `Run`.
exec
/bin/zsh -lc 'mkdir -p packages/gitstore/internal/local' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 519ms:
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/git.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/run.go
diff --git a/packages/gitstore/internal/local/git.go b/packages/gitstore/internal/local/git.go
new file mode 100644
index 0000000000000000000000000000000000000000..23d319eb3c560fe3c772d9560e45ef9273a914d6
--- /dev/null
+++ b/packages/gitstore/internal/local/git.go
@@ -0,0 +1,23 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func git(root string, env map[string]string, input []byte, args ...string) ([]byte, error) {
+	argv := append([]string{"git", "-C", root}, args...)
+	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input})
+	if err != nil {
+		return []byte(result.Stdout), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
+	}
+	return []byte(result.Stdout), nil
+}
+
+func gitString(root string, args ...string) (string, error) {
+	data, err := git(root, nil, nil, args...)
+	return strings.TrimSpace(string(data)), err
+}
diff --git a/packages/gitstore/internal/local/run.go b/packages/gitstore/internal/local/run.go
new file mode 100644
index 0000000000000000000000000000000000000000..ce18829152d3e597194e31e16fdb9ac31525c0a3
--- /dev/null
+++ b/packages/gitstore/internal/local/run.go
@@ -0,0 +1,66 @@
+package local
+
+import (
+	"bytes"
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+// Run executes one command in its own process group so cancellation also
+// terminates children which remain in that group.
+func Run(ctx context.Context, request model.Request) (model.Result, error) {
+	var result model.Result
+	if len(request.Args) == 0 {
+		return result, errors.New("empty command")
+	}
+	if ctx == nil {
+		return result, errors.New("nil context")
+	}
+	if request.Timeout > 0 {
+		var cancel context.CancelFunc
+		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
+		defer cancel()
+	}
+	if err := ctx.Err(); err != nil {
+		return result, err
+	}
+
+	cmd := exec.Command(request.Args[0], request.Args[1:]...)
+	cmd.Dir = request.Dir
+	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
+	cmd.Env = os.Environ()
+	for key, value := range request.Env {
+		cmd.Env = append(cmd.Env, key+"="+value)
+	}
+	cmd.Stdin = bytes.NewReader(request.Input)
+	var stdout, stderr bytes.Buffer
+	cmd.Stdout, cmd.Stderr = &stdout, &stderr
+	if err := cmd.Start(); err != nil {
+		return result, err
+	}
+	finished := make(chan error, 1)
+	go func() { finished <- cmd.Wait() }()
+	var err error
+	select {
+	case err = <-finished:
+	case <-ctx.Done():
+		// A process that leaves this group is outside the Run contract.
+		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
+		err = <-finished
+		result.Stdout, result.Stderr = stdout.String(), stderr.String()
+		result.ExitCode = cmd.ProcessState.ExitCode()
+		return result, ctx.Err()
+	}
+	result.Stdout, result.Stderr = stdout.String(), stderr.String()
+	result.ExitCode = cmd.ProcessState.ExitCode()
+	if err != nil {
+		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
+	}
+	return result, nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/objects.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/refs.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/repository.go
diff --git a/packages/gitstore/internal/local/git.go b/packages/gitstore/internal/local/git.go
new file mode 100644
index 0000000000000000000000000000000000000000..23d319eb3c560fe3c772d9560e45ef9273a914d6
--- /dev/null
+++ b/packages/gitstore/internal/local/git.go
@@ -0,0 +1,23 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func git(root string, env map[string]string, input []byte, args ...string) ([]byte, error) {
+	argv := append([]string{"git", "-C", root}, args...)
+	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input})
+	if err != nil {
+		return []byte(result.Stdout), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
+	}
+	return []byte(result.Stdout), nil
+}
+
+func gitString(root string, args ...string) (string, error) {
+	data, err := git(root, nil, nil, args...)
+	return strings.TrimSpace(string(data)), err
+}
diff --git a/packages/gitstore/internal/local/objects.go b/packages/gitstore/internal/local/objects.go
new file mode 100644
index 0000000000000000000000000000000000000000..7889d2b9989c12d98a4b7ad8202d58cc37a8dabe
--- /dev/null
+++ b/packages/gitstore/internal/local/objects.go
@@ -0,0 +1,25 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func WriteBlob(root string, data []byte) (string, error) {
+	out, err := git(root, nil, data, "hash-object", "-w", "--stdin")
+	return strings.TrimSpace(string(out)), err
+}
+
+func ReadBlob(root, oid string) ([]byte, error) {
+	if oid == "" || strings.HasPrefix(oid, "-") {
+		return nil, fmt.Errorf("invalid object ID %q", oid)
+	}
+	return git(root, nil, nil, "cat-file", "blob", oid)
+}
+
+func ReadAt(root, revision, path string) ([]byte, error) {
+	if revision == "" || path == "" {
+		return nil, fmt.Errorf("revision and path are required")
+	}
+	return git(root, nil, nil, "show", revision+":"+path)
+}
diff --git a/packages/gitstore/internal/local/refs.go b/packages/gitstore/internal/local/refs.go
new file mode 100644
index 0000000000000000000000000000000000000000..71e9a30473c0b7625ead7172e78363ce9fed518e
--- /dev/null
+++ b/packages/gitstore/internal/local/refs.go
@@ -0,0 +1,52 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func checkedRef(root, name string) error {
+	if !strings.HasPrefix(name, "refs/") {
+		return fmt.Errorf("ref must be fully qualified: %q", name)
+	}
+	_, err := gitString(root, "check-ref-format", name)
+	return err
+}
+
+func Ref(root, name string) (string, bool, error) {
+	if err := checkedRef(root, name); err != nil {
+		return "", false, err
+	}
+	// show-ref --exists distinguishes a missing ref (2) from lookup errors,
+	// but is not available in all Git versions. A successful show-ref --verify
+	// is enough; a failed lookup is checked with rev-parse to verify the repo.
+	out, err := gitString(root, "show-ref", "--verify", "--hash", name)
+	if err == nil {
+		return out, true, nil
+	}
+	if _, repoErr := gitString(root, "rev-parse", "--git-dir"); repoErr != nil {
+		return "", false, repoErr
+	}
+	return "", false, nil
+}
+
+func CASRef(root, name, old, new string) error {
+	if err := checkedRef(root, name); err != nil {
+		return err
+	}
+	// Git's all-zero object ID is the atomic "must not exist" condition.
+	zero, err := gitString(root, "hash-object", "--stdin")
+	if err != nil {
+		return err
+	}
+	zero = strings.Repeat("0", len(zero))
+	if old == "" {
+		old = zero
+	}
+	if new == "" {
+		_, err = git(root, nil, nil, "update-ref", "-d", name, old)
+		return err
+	}
+	_, err = git(root, nil, nil, "update-ref", name, new, old)
+	return err
+}
diff --git a/packages/gitstore/internal/local/repository.go b/packages/gitstore/internal/local/repository.go
new file mode 100644
index 0000000000000000000000000000000000000000..cb40716b89d389d52886c55739615794910322fe
--- /dev/null
+++ b/packages/gitstore/internal/local/repository.go
@@ -0,0 +1,72 @@
+package local
+
+import (
+	"fmt"
+	"path/filepath"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func canonical(path string) (string, error) {
+	abs, err := filepath.Abs(path)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(abs)
+}
+
+func Inspect(dir string) (model.Repository, error) {
+	var repository model.Repository
+	root, err := gitString(dir, "rev-parse", "--show-toplevel")
+	if err != nil {
+		return repository, err
+	}
+	common, err := gitString(dir, "rev-parse", "--path-format=absolute", "--git-common-dir")
+	if err != nil {
+		return repository, err
+	}
+	if repository.Root, err = canonical(root); err != nil {
+		return repository, err
+	}
+	if repository.CommonDir, err = canonical(common); err != nil {
+		return repository, err
+	}
+	if repository.Head, err = gitString(dir, "rev-parse", "--verify", "HEAD"); err != nil {
+		return repository, err
+	}
+	// symbolic-ref exits nonzero for detached HEAD.
+	repository.Branch, _ = gitString(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
+	return repository, nil
+}
+
+func BaseRef(root, remote, base string) (string, error) {
+	if remote == "" {
+		remote = "origin"
+	}
+	if base == "" {
+		head := "refs/remotes/" + remote + "/HEAD"
+		if target, err := gitString(root, "symbolic-ref", "--quiet", head); err == nil {
+			prefix := "refs/remotes/" + remote + "/"
+			if strings.HasPrefix(target, prefix) && len(target) > len(prefix) {
+				base = strings.TrimPrefix(target, prefix)
+			}
+		}
+		if base == "" {
+			base = "main"
+		}
+	}
+	remoteRef := "refs/remotes/" + remote + "/" + base
+	localRef := "refs/heads/" + base
+	for _, name := range []string{remoteRef, localRef} {
+		if _, err := gitString(root, "check-ref-format", name); err != nil {
+			return "", fmt.Errorf("invalid base ref %q: %w", name, err)
+		}
+		if _, exists, err := Ref(root, name); err != nil {
+			return "", err
+		} else if exists {
+			return name, nil
+		}
+	}
+	return "", fmt.Errorf("base branch %q is missing", base)
+}
diff --git a/packages/gitstore/internal/local/run.go b/packages/gitstore/internal/local/run.go
new file mode 100644
index 0000000000000000000000000000000000000000..ce18829152d3e597194e31e16fdb9ac31525c0a3
--- /dev/null
+++ b/packages/gitstore/internal/local/run.go
@@ -0,0 +1,66 @@
+package local
+
+import (
+	"bytes"
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+// Run executes one command in its own process group so cancellation also
+// terminates children which remain in that group.
+func Run(ctx context.Context, request model.Request) (model.Result, error) {
+	var result model.Result
+	if len(request.Args) == 0 {
+		return result, errors.New("empty command")
+	}
+	if ctx == nil {
+		return result, errors.New("nil context")
+	}
+	if request.Timeout > 0 {
+		var cancel context.CancelFunc
+		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
+		defer cancel()
+	}
+	if err := ctx.Err(); err != nil {
+		return result, err
+	}
+
+	cmd := exec.Command(request.Args[0], request.Args[1:]...)
+	cmd.Dir = request.Dir
+	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
+	cmd.Env = os.Environ()
+	for key, value := range request.Env {
+		cmd.Env = append(cmd.Env, key+"="+value)
+	}
+	cmd.Stdin = bytes.NewReader(request.Input)
+	var stdout, stderr bytes.Buffer
+	cmd.Stdout, cmd.Stderr = &stdout, &stderr
+	if err := cmd.Start(); err != nil {
+		return result, err
+	}
+	finished := make(chan error, 1)
+	go func() { finished <- cmd.Wait() }()
+	var err error
+	select {
+	case err = <-finished:
+	case <-ctx.Done():
+		// A process that leaves this group is outside the Run contract.
+		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
+		err = <-finished
+		result.Stdout, result.Stderr = stdout.String(), stderr.String()
+		result.ExitCode = cmd.ProcessState.ExitCode()
+		return result, ctx.Err()
+	}
+	result.Stdout, result.Stderr = stdout.String(), stderr.String()
+	result.ExitCode = cmd.ProcessState.ExitCode()
+	if err != nil {
+		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
+	}
+	return result, nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/DRILL_NOTES.md
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/changes.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/commit.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..e2b89a1e2dff87ba5fb68f6c858b165eaede1f66
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,7 @@
+## Contradicted Decisions
+
+None identified.
+
+## Gaps
+
+None identified.
diff --git a/packages/gitstore/internal/local/changes.go b/packages/gitstore/internal/local/changes.go
new file mode 100644
index 0000000000000000000000000000000000000000..d009b3e1734ca6d01464877d1be435c383dc05fc
--- /dev/null
+++ b/packages/gitstore/internal/local/changes.go
@@ -0,0 +1,86 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func Changes(root, base string) ([]model.Change, error) {
+	if base == "" {
+		return nil, fmt.Errorf("base ref is required")
+	}
+	byPath := make(map[string]model.Change)
+	order := make([]string, 0)
+	add := func(change model.Change) {
+		if prior, exists := byPath[change.Path]; exists {
+			// Preserve the rename when a later diff also reports its destination.
+			if prior.OldPath != "" && change.OldPath == "" {
+				return
+			}
+		} else {
+			order = append(order, change.Path)
+		}
+		byPath[change.Path] = change
+	}
+	for _, args := range [][]string{
+		{"diff", "--cached", "--no-ext-diff", "--name-status", "-z", "-M", base, "--"},
+		{"diff", "--no-ext-diff", "--name-status", "-z", "-M", "--"},
+	} {
+		data, err := git(root, nil, nil, args...)
+		if err != nil {
+			return nil, err
+		}
+		if err := parseNameStatus(data, add); err != nil {
+			return nil, err
+		}
+	}
+	data, err := git(root, nil, nil, "ls-files", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return nil, err
+	}
+	if len(data) > 0 {
+		for _, path := range strings.Split(strings.TrimSuffix(string(data), "\x00"), "\x00") {
+			if path != "" {
+				add(model.Change{Path: path, Status: "?"})
+			}
+		}
+	}
+	changes := make([]model.Change, 0, len(order))
+	for _, path := range order {
+		changes = append(changes, byPath[path])
+	}
+	return changes, nil
+}
+
+func parseNameStatus(data []byte, add func(model.Change)) error {
+	if len(data) == 0 {
+		return nil
+	}
+	if data[len(data)-1] != 0 {
+		return fmt.Errorf("malformed Git name-status output")
+	}
+	fields := strings.Split(string(data[:len(data)-1]), "\x00")
+	for i := 0; i < len(fields); {
+		status := fields[i]
+		i++
+		if status == "" || i >= len(fields) {
+			return fmt.Errorf("malformed Git name-status entry")
+		}
+		change := model.Change{Status: status[:1]}
+		if change.Status == "R" || change.Status == "C" {
+			if i+1 >= len(fields) {
+				return fmt.Errorf("malformed Git rename entry")
+			}
+			change.OldPath = fields[i]
+			change.Path = fields[i+1]
+			i += 2
+		} else {
+			change.Path = fields[i]
+			i++
+		}
+		add(change)
+	}
+	return nil
+}
diff --git a/packages/gitstore/internal/local/commit.go b/packages/gitstore/internal/local/commit.go
new file mode 100644
index 0000000000000000000000000000000000000000..ddcdc58ed66b34546f5fc194a697f6c2cedfa584
--- /dev/null
+++ b/packages/gitstore/internal/local/commit.go
@@ -0,0 +1,64 @@
+package local
+
+import (
+	"fmt"
+	"os"
+	"path/filepath"
+	"strings"
+)
+
+func Commit(root, subject string, paths []string) (string, error) {
+	if len(paths) == 0 {
+		return "", fmt.Errorf("at least one path is required")
+	}
+	selected := make([]string, 0, len(paths))
+	for _, path := range paths {
+		clean := filepath.Clean(path)
+		if path == "" || filepath.IsAbs(path) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
+			return "", fmt.Errorf("path must be inside the worktree: %q", path)
+		}
+		// Git pathspec magic makes the selection literal, including names which
+		// contain glob characters or begin with a colon.
+		selected = append(selected, ":(literal)"+filepath.ToSlash(clean))
+	}
+	old, err := gitString(root, "rev-parse", "--verify", "HEAD")
+	if err != nil {
+		return "", err
+	}
+	index, err := os.CreateTemp("", "gitstore-index-")
+	if err != nil {
+		return "", err
+	}
+	indexPath := index.Name()
+	_ = index.Close()
+	_ = os.Remove(indexPath)
+	defer os.Remove(indexPath)
+	env := map[string]string{"GIT_INDEX_FILE": indexPath}
+	if _, err := git(root, env, nil, "read-tree", old); err != nil {
+		return "", err
+	}
+	addArgs := append([]string{"add", "-A", "--"}, selected...)
+	if _, err := git(root, env, nil, addArgs...); err != nil {
+		return "", err
+	}
+	treeData, err := git(root, env, nil, "write-tree")
+	if err != nil {
+		return "", err
+	}
+	tree := strings.TrimSpace(string(treeData))
+	commitData, err := git(root, env, nil, "commit-tree", tree, "-p", old, "-m", subject)
+	if err != nil {
+		return "", err
+	}
+	newHead := strings.TrimSpace(string(commitData))
+	if _, err := git(root, nil, nil, "update-ref", "HEAD", newHead, old); err != nil {
+		return "", err
+	}
+	// Bring only selected entries of the real index into line with the new
+	// commit. This leaves every unrelated staged entry untouched.
+	resetArgs := append([]string{"reset", "-q", "HEAD", "--"}, selected...)
+	if _, err := git(root, nil, nil, resetArgs...); err != nil {
+		return newHead, err
+	}
+	return newHead, nil
+}
diff --git a/packages/gitstore/internal/local/git.go b/packages/gitstore/internal/local/git.go
new file mode 100644
index 0000000000000000000000000000000000000000..23d319eb3c560fe3c772d9560e45ef9273a914d6
--- /dev/null
+++ b/packages/gitstore/internal/local/git.go
@@ -0,0 +1,23 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func git(root string, env map[string]string, input []byte, args ...string) ([]byte, error) {
+	argv := append([]string{"git", "-C", root}, args...)
+	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input})
+	if err != nil {
+		return []byte(result.Stdout), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
+	}
+	return []byte(result.Stdout), nil
+}
+
+func gitString(root string, args ...string) (string, error) {
+	data, err := git(root, nil, nil, args...)
+	return strings.TrimSpace(string(data)), err
+}
diff --git a/packages/gitstore/internal/local/objects.go b/packages/gitstore/internal/local/objects.go
new file mode 100644
index 0000000000000000000000000000000000000000..7889d2b9989c12d98a4b7ad8202d58cc37a8dabe
--- /dev/null
+++ b/packages/gitstore/internal/local/objects.go
@@ -0,0 +1,25 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func WriteBlob(root string, data []byte) (string, error) {
+	out, err := git(root, nil, data, "hash-object", "-w", "--stdin")
+	return strings.TrimSpace(string(out)), err
+}
+
+func ReadBlob(root, oid string) ([]byte, error) {
+	if oid == "" || strings.HasPrefix(oid, "-") {
+		return nil, fmt.Errorf("invalid object ID %q", oid)
+	}
+	return git(root, nil, nil, "cat-file", "blob", oid)
+}
+
+func ReadAt(root, revision, path string) ([]byte, error) {
+	if revision == "" || path == "" {
+		return nil, fmt.Errorf("revision and path are required")
+	}
+	return git(root, nil, nil, "show", revision+":"+path)
+}
diff --git a/packages/gitstore/internal/local/refs.go b/packages/gitstore/internal/local/refs.go
new file mode 100644
index 0000000000000000000000000000000000000000..71e9a30473c0b7625ead7172e78363ce9fed518e
--- /dev/null
+++ b/packages/gitstore/internal/local/refs.go
@@ -0,0 +1,52 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func checkedRef(root, name string) error {
+	if !strings.HasPrefix(name, "refs/") {
+		return fmt.Errorf("ref must be fully qualified: %q", name)
+	}
+	_, err := gitString(root, "check-ref-format", name)
+	return err
+}
+
+func Ref(root, name string) (string, bool, error) {
+	if err := checkedRef(root, name); err != nil {
+		return "", false, err
+	}
+	// show-ref --exists distinguishes a missing ref (2) from lookup errors,
+	// but is not available in all Git versions. A successful show-ref --verify
+	// is enough; a failed lookup is checked with rev-parse to verify the repo.
+	out, err := gitString(root, "show-ref", "--verify", "--hash", name)
+	if err == nil {
+		return out, true, nil
+	}
+	if _, repoErr := gitString(root, "rev-parse", "--git-dir"); repoErr != nil {
+		return "", false, repoErr
+	}
+	return "", false, nil
+}
+
+func CASRef(root, name, old, new string) error {
+	if err := checkedRef(root, name); err != nil {
+		return err
+	}
+	// Git's all-zero object ID is the atomic "must not exist" condition.
+	zero, err := gitString(root, "hash-object", "--stdin")
+	if err != nil {
+		return err
+	}
+	zero = strings.Repeat("0", len(zero))
+	if old == "" {
+		old = zero
+	}
+	if new == "" {
+		_, err = git(root, nil, nil, "update-ref", "-d", name, old)
+		return err
+	}
+	_, err = git(root, nil, nil, "update-ref", name, new, old)
+	return err
+}
diff --git a/packages/gitstore/internal/local/repository.go b/packages/gitstore/internal/local/repository.go
new file mode 100644
index 0000000000000000000000000000000000000000..cb40716b89d389d52886c55739615794910322fe
--- /dev/null
+++ b/packages/gitstore/internal/local/repository.go
@@ -0,0 +1,72 @@
+package local
+
+import (
+	"fmt"
+	"path/filepath"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func canonical(path string) (string, error) {
+	abs, err := filepath.Abs(path)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(abs)
+}
+
+func Inspect(dir string) (model.Repository, error) {
+	var repository model.Repository
+	root, err := gitString(dir, "rev-parse", "--show-toplevel")
+	if err != nil {
+		return repository, err
+	}
+	common, err := gitString(dir, "rev-parse", "--path-format=absolute", "--git-common-dir")
+	if err != nil {
+		return repository, err
+	}
+	if repository.Root, err = canonical(root); err != nil {
+		return repository, err
+	}
+	if repository.CommonDir, err = canonical(common); err != nil {
+		return repository, err
+	}
+	if repository.Head, err = gitString(dir, "rev-parse", "--verify", "HEAD"); err != nil {
+		return repository, err
+	}
+	// symbolic-ref exits nonzero for detached HEAD.
+	repository.Branch, _ = gitString(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
+	return repository, nil
+}
+
+func BaseRef(root, remote, base string) (string, error) {
+	if remote == "" {
+		remote = "origin"
+	}
+	if base == "" {
+		head := "refs/remotes/" + remote + "/HEAD"
+		if target, err := gitString(root, "symbolic-ref", "--quiet", head); err == nil {
+			prefix := "refs/remotes/" + remote + "/"
+			if strings.HasPrefix(target, prefix) && len(target) > len(prefix) {
+				base = strings.TrimPrefix(target, prefix)
+			}
+		}
+		if base == "" {
+			base = "main"
+		}
+	}
+	remoteRef := "refs/remotes/" + remote + "/" + base
+	localRef := "refs/heads/" + base
+	for _, name := range []string{remoteRef, localRef} {
+		if _, err := gitString(root, "check-ref-format", name); err != nil {
+			return "", fmt.Errorf("invalid base ref %q: %w", name, err)
+		}
+		if _, exists, err := Ref(root, name); err != nil {
+			return "", err
+		} else if exists {
+			return name, nil
+		}
+	}
+	return "", fmt.Errorf("base branch %q is missing", base)
+}
diff --git a/packages/gitstore/internal/local/run.go b/packages/gitstore/internal/local/run.go
new file mode 100644
index 0000000000000000000000000000000000000000..ce18829152d3e597194e31e16fdb9ac31525c0a3
--- /dev/null
+++ b/packages/gitstore/internal/local/run.go
@@ -0,0 +1,66 @@
+package local
+
+import (
+	"bytes"
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+// Run executes one command in its own process group so cancellation also
+// terminates children which remain in that group.
+func Run(ctx context.Context, request model.Request) (model.Result, error) {
+	var result model.Result
+	if len(request.Args) == 0 {
+		return result, errors.New("empty command")
+	}
+	if ctx == nil {
+		return result, errors.New("nil context")
+	}
+	if request.Timeout > 0 {
+		var cancel context.CancelFunc
+		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
+		defer cancel()
+	}
+	if err := ctx.Err(); err != nil {
+		return result, err
+	}
+
+	cmd := exec.Command(request.Args[0], request.Args[1:]...)
+	cmd.Dir = request.Dir
+	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
+	cmd.Env = os.Environ()
+	for key, value := range request.Env {
+		cmd.Env = append(cmd.Env, key+"="+value)
+	}
+	cmd.Stdin = bytes.NewReader(request.Input)
+	var stdout, stderr bytes.Buffer
+	cmd.Stdout, cmd.Stderr = &stdout, &stderr
+	if err := cmd.Start(); err != nil {
+		return result, err
+	}
+	finished := make(chan error, 1)
+	go func() { finished <- cmd.Wait() }()
+	var err error
+	select {
+	case err = <-finished:
+	case <-ctx.Done():
+		// A process that leaves this group is outside the Run contract.
+		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
+		err = <-finished
+		result.Stdout, result.Stderr = stdout.String(), stderr.String()
+		result.ExitCode = cmd.ProcessState.ExitCode()
+		return result, ctx.Err()
+	}
+	result.Stdout, result.Stderr = stdout.String(), stderr.String()
+	result.ExitCode = cmd.ProcessState.ExitCode()
+	if err != nil {
+		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
+	}
+	return result, nil
+}

codex
The core implementation is in place. I’m formatting it and running the retained evals through a temporary GOPATH layout inside the allowed disposable path, since this input deliberately has no build manifest.
exec
/bin/zsh -lc 'gofmt -w packages/gitstore/internal/local/*.go && mkdir -p packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages && ln -s ../../../../../../../../gitstore packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 606ms:
exec
/bin/zsh -lc 'readlink packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore && ls packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 470ms:
../../../../../../../../gitstore
packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore

exec
/bin/zsh -lc 'rm packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore && ln -s /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore && ls packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 514ms:
SPEC.md
api.go
evals
internal
model

exec
/bin/zsh -lc 'env GO111MODULE=off GOPATH=/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/.test-gopath GOCACHE=/tmp/fireweed-rebuild-cache go test github.com/lurkshark/fireweed/packages/gitstore/evals -count=1' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 exited 1 in 6744ms:
--- FAIL: Test_GIT_001_InspectWorktreesAndDetachedHead (0.16s)
    gitstore_test.go:71: lstat /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/evals/--path-format=absolute
        ..: no such file or directory
FAIL
FAIL	github.com/lurkshark/fireweed/packages/gitstore/evals	5.328s
FAIL

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/repository.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..e2b89a1e2dff87ba5fb68f6c858b165eaede1f66
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,7 @@
+## Contradicted Decisions
+
+None identified.
+
+## Gaps
+
+None identified.
diff --git a/packages/gitstore/internal/local/changes.go b/packages/gitstore/internal/local/changes.go
new file mode 100644
index 0000000000000000000000000000000000000000..d009b3e1734ca6d01464877d1be435c383dc05fc
--- /dev/null
+++ b/packages/gitstore/internal/local/changes.go
@@ -0,0 +1,86 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func Changes(root, base string) ([]model.Change, error) {
+	if base == "" {
+		return nil, fmt.Errorf("base ref is required")
+	}
+	byPath := make(map[string]model.Change)
+	order := make([]string, 0)
+	add := func(change model.Change) {
+		if prior, exists := byPath[change.Path]; exists {
+			// Preserve the rename when a later diff also reports its destination.
+			if prior.OldPath != "" && change.OldPath == "" {
+				return
+			}
+		} else {
+			order = append(order, change.Path)
+		}
+		byPath[change.Path] = change
+	}
+	for _, args := range [][]string{
+		{"diff", "--cached", "--no-ext-diff", "--name-status", "-z", "-M", base, "--"},
+		{"diff", "--no-ext-diff", "--name-status", "-z", "-M", "--"},
+	} {
+		data, err := git(root, nil, nil, args...)
+		if err != nil {
+			return nil, err
+		}
+		if err := parseNameStatus(data, add); err != nil {
+			return nil, err
+		}
+	}
+	data, err := git(root, nil, nil, "ls-files", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return nil, err
+	}
+	if len(data) > 0 {
+		for _, path := range strings.Split(strings.TrimSuffix(string(data), "\x00"), "\x00") {
+			if path != "" {
+				add(model.Change{Path: path, Status: "?"})
+			}
+		}
+	}
+	changes := make([]model.Change, 0, len(order))
+	for _, path := range order {
+		changes = append(changes, byPath[path])
+	}
+	return changes, nil
+}
+
+func parseNameStatus(data []byte, add func(model.Change)) error {
+	if len(data) == 0 {
+		return nil
+	}
+	if data[len(data)-1] != 0 {
+		return fmt.Errorf("malformed Git name-status output")
+	}
+	fields := strings.Split(string(data[:len(data)-1]), "\x00")
+	for i := 0; i < len(fields); {
+		status := fields[i]
+		i++
+		if status == "" || i >= len(fields) {
+			return fmt.Errorf("malformed Git name-status entry")
+		}
+		change := model.Change{Status: status[:1]}
+		if change.Status == "R" || change.Status == "C" {
+			if i+1 >= len(fields) {
+				return fmt.Errorf("malformed Git rename entry")
+			}
+			change.OldPath = fields[i]
+			change.Path = fields[i+1]
+			i += 2
+		} else {
+			change.Path = fields[i]
+			i++
+		}
+		add(change)
+	}
+	return nil
+}
diff --git a/packages/gitstore/internal/local/commit.go b/packages/gitstore/internal/local/commit.go
new file mode 100644
index 0000000000000000000000000000000000000000..ddcdc58ed66b34546f5fc194a697f6c2cedfa584
--- /dev/null
+++ b/packages/gitstore/internal/local/commit.go
@@ -0,0 +1,64 @@
+package local
+
+import (
+	"fmt"
+	"os"
+	"path/filepath"
+	"strings"
+)
+
+func Commit(root, subject string, paths []string) (string, error) {
+	if len(paths) == 0 {
+		return "", fmt.Errorf("at least one path is required")
+	}
+	selected := make([]string, 0, len(paths))
+	for _, path := range paths {
+		clean := filepath.Clean(path)
+		if path == "" || filepath.IsAbs(path) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
+			return "", fmt.Errorf("path must be inside the worktree: %q", path)
+		}
+		// Git pathspec magic makes the selection literal, including names which
+		// contain glob characters or begin with a colon.
+		selected = append(selected, ":(literal)"+filepath.ToSlash(clean))
+	}
+	old, err := gitString(root, "rev-parse", "--verify", "HEAD")
+	if err != nil {
+		return "", err
+	}
+	index, err := os.CreateTemp("", "gitstore-index-")
+	if err != nil {
+		return "", err
+	}
+	indexPath := index.Name()
+	_ = index.Close()
+	_ = os.Remove(indexPath)
+	defer os.Remove(indexPath)
+	env := map[string]string{"GIT_INDEX_FILE": indexPath}
+	if _, err := git(root, env, nil, "read-tree", old); err != nil {
+		return "", err
+	}
+	addArgs := append([]string{"add", "-A", "--"}, selected...)
+	if _, err := git(root, env, nil, addArgs...); err != nil {
+		return "", err
+	}
+	treeData, err := git(root, env, nil, "write-tree")
+	if err != nil {
+		return "", err
+	}
+	tree := strings.TrimSpace(string(treeData))
+	commitData, err := git(root, env, nil, "commit-tree", tree, "-p", old, "-m", subject)
+	if err != nil {
+		return "", err
+	}
+	newHead := strings.TrimSpace(string(commitData))
+	if _, err := git(root, nil, nil, "update-ref", "HEAD", newHead, old); err != nil {
+		return "", err
+	}
+	// Bring only selected entries of the real index into line with the new
+	// commit. This leaves every unrelated staged entry untouched.
+	resetArgs := append([]string{"reset", "-q", "HEAD", "--"}, selected...)
+	if _, err := git(root, nil, nil, resetArgs...); err != nil {
+		return newHead, err
+	}
+	return newHead, nil
+}
diff --git a/packages/gitstore/internal/local/git.go b/packages/gitstore/internal/local/git.go
new file mode 100644
index 0000000000000000000000000000000000000000..23d319eb3c560fe3c772d9560e45ef9273a914d6
--- /dev/null
+++ b/packages/gitstore/internal/local/git.go
@@ -0,0 +1,23 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func git(root string, env map[string]string, input []byte, args ...string) ([]byte, error) {
+	argv := append([]string{"git", "-C", root}, args...)
+	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input})
+	if err != nil {
+		return []byte(result.Stdout), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
+	}
+	return []byte(result.Stdout), nil
+}
+
+func gitString(root string, args ...string) (string, error) {
+	data, err := git(root, nil, nil, args...)
+	return strings.TrimSpace(string(data)), err
+}
diff --git a/packages/gitstore/internal/local/objects.go b/packages/gitstore/internal/local/objects.go
new file mode 100644
index 0000000000000000000000000000000000000000..7889d2b9989c12d98a4b7ad8202d58cc37a8dabe
--- /dev/null
+++ b/packages/gitstore/internal/local/objects.go
@@ -0,0 +1,25 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func WriteBlob(root string, data []byte) (string, error) {
+	out, err := git(root, nil, data, "hash-object", "-w", "--stdin")
+	return strings.TrimSpace(string(out)), err
+}
+
+func ReadBlob(root, oid string) ([]byte, error) {
+	if oid == "" || strings.HasPrefix(oid, "-") {
+		return nil, fmt.Errorf("invalid object ID %q", oid)
+	}
+	return git(root, nil, nil, "cat-file", "blob", oid)
+}
+
+func ReadAt(root, revision, path string) ([]byte, error) {
+	if revision == "" || path == "" {
+		return nil, fmt.Errorf("revision and path are required")
+	}
+	return git(root, nil, nil, "show", revision+":"+path)
+}
diff --git a/packages/gitstore/internal/local/refs.go b/packages/gitstore/internal/local/refs.go
new file mode 100644
index 0000000000000000000000000000000000000000..71e9a30473c0b7625ead7172e78363ce9fed518e
--- /dev/null
+++ b/packages/gitstore/internal/local/refs.go
@@ -0,0 +1,52 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func checkedRef(root, name string) error {
+	if !strings.HasPrefix(name, "refs/") {
+		return fmt.Errorf("ref must be fully qualified: %q", name)
+	}
+	_, err := gitString(root, "check-ref-format", name)
+	return err
+}
+
+func Ref(root, name string) (string, bool, error) {
+	if err := checkedRef(root, name); err != nil {
+		return "", false, err
+	}
+	// show-ref --exists distinguishes a missing ref (2) from lookup errors,
+	// but is not available in all Git versions. A successful show-ref --verify
+	// is enough; a failed lookup is checked with rev-parse to verify the repo.
+	out, err := gitString(root, "show-ref", "--verify", "--hash", name)
+	if err == nil {
+		return out, true, nil
+	}
+	if _, repoErr := gitString(root, "rev-parse", "--git-dir"); repoErr != nil {
+		return "", false, repoErr
+	}
+	return "", false, nil
+}
+
+func CASRef(root, name, old, new string) error {
+	if err := checkedRef(root, name); err != nil {
+		return err
+	}
+	// Git's all-zero object ID is the atomic "must not exist" condition.
+	zero, err := gitString(root, "hash-object", "--stdin")
+	if err != nil {
+		return err
+	}
+	zero = strings.Repeat("0", len(zero))
+	if old == "" {
+		old = zero
+	}
+	if new == "" {
+		_, err = git(root, nil, nil, "update-ref", "-d", name, old)
+		return err
+	}
+	_, err = git(root, nil, nil, "update-ref", name, new, old)
+	return err
+}
diff --git a/packages/gitstore/internal/local/repository.go b/packages/gitstore/internal/local/repository.go
new file mode 100644
index 0000000000000000000000000000000000000000..25ed03220b1baa2a0b48941d3873335b00616399
--- /dev/null
+++ b/packages/gitstore/internal/local/repository.go
@@ -0,0 +1,75 @@
+package local
+
+import (
+	"fmt"
+	"path/filepath"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func canonical(path string) (string, error) {
+	abs, err := filepath.Abs(path)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(abs)
+}
+
+func Inspect(dir string) (model.Repository, error) {
+	var repository model.Repository
+	root, err := gitString(dir, "rev-parse", "--show-toplevel")
+	if err != nil {
+		return repository, err
+	}
+	common, err := gitString(dir, "rev-parse", "--git-common-dir")
+	if err != nil {
+		return repository, err
+	}
+	if !filepath.IsAbs(common) {
+		common = filepath.Join(dir, common)
+	}
+	if repository.Root, err = canonical(root); err != nil {
+		return repository, err
+	}
+	if repository.CommonDir, err = canonical(common); err != nil {
+		return repository, err
+	}
+	if repository.Head, err = gitString(dir, "rev-parse", "--verify", "HEAD"); err != nil {
+		return repository, err
+	}
+	// symbolic-ref exits nonzero for detached HEAD.
+	repository.Branch, _ = gitString(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
+	return repository, nil
+}
+
+func BaseRef(root, remote, base string) (string, error) {
+	if remote == "" {
+		remote = "origin"
+	}
+	if base == "" {
+		head := "refs/remotes/" + remote + "/HEAD"
+		if target, err := gitString(root, "symbolic-ref", "--quiet", head); err == nil {
+			prefix := "refs/remotes/" + remote + "/"
+			if strings.HasPrefix(target, prefix) && len(target) > len(prefix) {
+				base = strings.TrimPrefix(target, prefix)
+			}
+		}
+		if base == "" {
+			base = "main"
+		}
+	}
+	remoteRef := "refs/remotes/" + remote + "/" + base
+	localRef := "refs/heads/" + base
+	for _, name := range []string{remoteRef, localRef} {
+		if _, err := gitString(root, "check-ref-format", name); err != nil {
+			return "", fmt.Errorf("invalid base ref %q: %w", name, err)
+		}
+		if _, exists, err := Ref(root, name); err != nil {
+			return "", err
+		} else if exists {
+			return name, nil
+		}
+	}
+	return "", fmt.Errorf("base branch %q is missing", base)
+}
diff --git a/packages/gitstore/internal/local/run.go b/packages/gitstore/internal/local/run.go
new file mode 100644
index 0000000000000000000000000000000000000000..ce18829152d3e597194e31e16fdb9ac31525c0a3
--- /dev/null
+++ b/packages/gitstore/internal/local/run.go
@@ -0,0 +1,66 @@
+package local
+
+import (
+	"bytes"
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+// Run executes one command in its own process group so cancellation also
+// terminates children which remain in that group.
+func Run(ctx context.Context, request model.Request) (model.Result, error) {
+	var result model.Result
+	if len(request.Args) == 0 {
+		return result, errors.New("empty command")
+	}
+	if ctx == nil {
+		return result, errors.New("nil context")
+	}
+	if request.Timeout > 0 {
+		var cancel context.CancelFunc
+		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
+		defer cancel()
+	}
+	if err := ctx.Err(); err != nil {
+		return result, err
+	}
+
+	cmd := exec.Command(request.Args[0], request.Args[1:]...)
+	cmd.Dir = request.Dir
+	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
+	cmd.Env = os.Environ()
+	for key, value := range request.Env {
+		cmd.Env = append(cmd.Env, key+"="+value)
+	}
+	cmd.Stdin = bytes.NewReader(request.Input)
+	var stdout, stderr bytes.Buffer
+	cmd.Stdout, cmd.Stderr = &stdout, &stderr
+	if err := cmd.Start(); err != nil {
+		return result, err
+	}
+	finished := make(chan error, 1)
+	go func() { finished <- cmd.Wait() }()
+	var err error
+	select {
+	case err = <-finished:
+	case <-ctx.Done():
+		// A process that leaves this group is outside the Run contract.
+		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
+		err = <-finished
+		result.Stdout, result.Stderr = stdout.String(), stderr.String()
+		result.ExitCode = cmd.ProcessState.ExitCode()
+		return result, ctx.Err()
+	}
+	result.Stdout, result.Stderr = stdout.String(), stderr.String()
+	result.ExitCode = cmd.ProcessState.ExitCode()
+	if err != nil {
+		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
+	}
+	return result, nil
+}

exec
/bin/zsh -lc 'gofmt -w packages/gitstore/internal/local/repository.go && env GO111MODULE=off GOPATH=/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/.test-gopath GOCACHE=/tmp/fireweed-rebuild-cache go test github.com/lurkshark/fireweed/packages/gitstore/evals -count=1' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 6458ms:
ok  	github.com/lurkshark/fireweed/packages/gitstore/evals	5.164s

codex
The retained evals now pass. The first run exposed an older Git behavior in this environment: `rev-parse` echoed `--path-format=absolute` instead of handling it, so I switched to resolving `--git-common-dir` myself. I’m doing a final review of the disposable files and removing the test scaffold.
exec
/bin/zsh -lc 'git status --short && git diff --stat && git diff --check' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc "sed -n '1,240p' packages/gitstore/internal/local/commit.go && sed -n '1,220p' packages/gitstore/internal/local/changes.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc "sed -n '1,200p' packages/gitstore/internal/local/repository.go && sed -n '1,200p' packages/gitstore/internal/local/refs.go && sed -n '1,180p' packages/gitstore/internal/local/run.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 587ms:
package local

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Commit(root, subject string, paths []string) (string, error) {
	if len(paths) == 0 {
		return "", fmt.Errorf("at least one path is required")
	}
	selected := make([]string, 0, len(paths))
	for _, path := range paths {
		clean := filepath.Clean(path)
		if path == "" || filepath.IsAbs(path) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
			return "", fmt.Errorf("path must be inside the worktree: %q", path)
		}
		// Git pathspec magic makes the selection literal, including names which
		// contain glob characters or begin with a colon.
		selected = append(selected, ":(literal)"+filepath.ToSlash(clean))
	}
	old, err := gitString(root, "rev-parse", "--verify", "HEAD")
	if err != nil {
		return "", err
	}
	index, err := os.CreateTemp("", "gitstore-index-")
	if err != nil {
		return "", err
	}
	indexPath := index.Name()
	_ = index.Close()
	_ = os.Remove(indexPath)
	defer os.Remove(indexPath)
	env := map[string]string{"GIT_INDEX_FILE": indexPath}
	if _, err := git(root, env, nil, "read-tree", old); err != nil {
		return "", err
	}
	addArgs := append([]string{"add", "-A", "--"}, selected...)
	if _, err := git(root, env, nil, addArgs...); err != nil {
		return "", err
	}
	treeData, err := git(root, env, nil, "write-tree")
	if err != nil {
		return "", err
	}
	tree := strings.TrimSpace(string(treeData))
	commitData, err := git(root, env, nil, "commit-tree", tree, "-p", old, "-m", subject)
	if err != nil {
		return "", err
	}
	newHead := strings.TrimSpace(string(commitData))
	if _, err := git(root, nil, nil, "update-ref", "HEAD", newHead, old); err != nil {
		return "", err
	}
	// Bring only selected entries of the real index into line with the new
	// commit. This leaves every unrelated staged entry untouched.
	resetArgs := append([]string{"reset", "-q", "HEAD", "--"}, selected...)
	if _, err := git(root, nil, nil, resetArgs...); err != nil {
		return newHead, err
	}
	return newHead, nil
}
package local

import (
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

func Changes(root, base string) ([]model.Change, error) {
	if base == "" {
		return nil, fmt.Errorf("base ref is required")
	}
	byPath := make(map[string]model.Change)
	order := make([]string, 0)
	add := func(change model.Change) {
		if prior, exists := byPath[change.Path]; exists {
			// Preserve the rename when a later diff also reports its destination.
			if prior.OldPath != "" && change.OldPath == "" {
				return
			}
		} else {
			order = append(order, change.Path)
		}
		byPath[change.Path] = change
	}
	for _, args := range [][]string{
		{"diff", "--cached", "--no-ext-diff", "--name-status", "-z", "-M", base, "--"},
		{"diff", "--no-ext-diff", "--name-status", "-z", "-M", "--"},
	} {
		data, err := git(root, nil, nil, args...)
		if err != nil {
			return nil, err
		}
		if err := parseNameStatus(data, add); err != nil {
			return nil, err
		}
	}
	data, err := git(root, nil, nil, "ls-files", "--others", "--exclude-standard", "-z")
	if err != nil {
		return nil, err
	}
	if len(data) > 0 {
		for _, path := range strings.Split(strings.TrimSuffix(string(data), "\x00"), "\x00") {
			if path != "" {
				add(model.Change{Path: path, Status: "?"})
			}
		}
	}
	changes := make([]model.Change, 0, len(order))
	for _, path := range order {
		changes = append(changes, byPath[path])
	}
	return changes, nil
}

func parseNameStatus(data []byte, add func(model.Change)) error {
	if len(data) == 0 {
		return nil
	}
	if data[len(data)-1] != 0 {
		return fmt.Errorf("malformed Git name-status output")
	}
	fields := strings.Split(string(data[:len(data)-1]), "\x00")
	for i := 0; i < len(fields); {
		status := fields[i]
		i++
		if status == "" || i >= len(fields) {
			return fmt.Errorf("malformed Git name-status entry")
		}
		change := model.Change{Status: status[:1]}
		if change.Status == "R" || change.Status == "C" {
			if i+1 >= len(fields) {
				return fmt.Errorf("malformed Git rename entry")
			}
			change.OldPath = fields[i]
			change.Path = fields[i+1]
			i += 2
		} else {
			change.Path = fields[i]
			i++
		}
		add(change)
	}
	return nil
}

 succeeded in 591ms:
package local

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

func canonical(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(abs)
}

func Inspect(dir string) (model.Repository, error) {
	var repository model.Repository
	root, err := gitString(dir, "rev-parse", "--show-toplevel")
	if err != nil {
		return repository, err
	}
	common, err := gitString(dir, "rev-parse", "--git-common-dir")
	if err != nil {
		return repository, err
	}
	if !filepath.IsAbs(common) {
		common = filepath.Join(dir, common)
	}
	if repository.Root, err = canonical(root); err != nil {
		return repository, err
	}
	if repository.CommonDir, err = canonical(common); err != nil {
		return repository, err
	}
	if repository.Head, err = gitString(dir, "rev-parse", "--verify", "HEAD"); err != nil {
		return repository, err
	}
	// symbolic-ref exits nonzero for detached HEAD.
	repository.Branch, _ = gitString(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
	return repository, nil
}

func BaseRef(root, remote, base string) (string, error) {
	if remote == "" {
		remote = "origin"
	}
	if base == "" {
		head := "refs/remotes/" + remote + "/HEAD"
		if target, err := gitString(root, "symbolic-ref", "--quiet", head); err == nil {
			prefix := "refs/remotes/" + remote + "/"
			if strings.HasPrefix(target, prefix) && len(target) > len(prefix) {
				base = strings.TrimPrefix(target, prefix)
			}
		}
		if base == "" {
			base = "main"
		}
	}
	remoteRef := "refs/remotes/" + remote + "/" + base
	localRef := "refs/heads/" + base
	for _, name := range []string{remoteRef, localRef} {
		if _, err := gitString(root, "check-ref-format", name); err != nil {
			return "", fmt.Errorf("invalid base ref %q: %w", name, err)
		}
		if _, exists, err := Ref(root, name); err != nil {
			return "", err
		} else if exists {
			return name, nil
		}
	}
	return "", fmt.Errorf("base branch %q is missing", base)
}
package local

import (
	"fmt"
	"strings"
)

func checkedRef(root, name string) error {
	if !strings.HasPrefix(name, "refs/") {
		return fmt.Errorf("ref must be fully qualified: %q", name)
	}
	_, err := gitString(root, "check-ref-format", name)
	return err
}

func Ref(root, name string) (string, bool, error) {
	if err := checkedRef(root, name); err != nil {
		return "", false, err
	}
	// show-ref --exists distinguishes a missing ref (2) from lookup errors,
	// but is not available in all Git versions. A successful show-ref --verify
	// is enough; a failed lookup is checked with rev-parse to verify the repo.
	out, err := gitString(root, "show-ref", "--verify", "--hash", name)
	if err == nil {
		return out, true, nil
	}
	if _, repoErr := gitString(root, "rev-parse", "--git-dir"); repoErr != nil {
		return "", false, repoErr
	}
	return "", false, nil
}

func CASRef(root, name, old, new string) error {
	if err := checkedRef(root, name); err != nil {
		return err
	}
	// Git's all-zero object ID is the atomic "must not exist" condition.
	zero, err := gitString(root, "hash-object", "--stdin")
	if err != nil {
		return err
	}
	zero = strings.Repeat("0", len(zero))
	if old == "" {
		old = zero
	}
	if new == "" {
		_, err = git(root, nil, nil, "update-ref", "-d", name, old)
		return err
	}
	_, err = git(root, nil, nil, "update-ref", name, new, old)
	return err
}
package local

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

// Run executes one command in its own process group so cancellation also
// terminates children which remain in that group.
func Run(ctx context.Context, request model.Request) (model.Result, error) {
	var result model.Result
	if len(request.Args) == 0 {
		return result, errors.New("empty command")
	}
	if ctx == nil {
		return result, errors.New("nil context")
	}
	if request.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
		defer cancel()
	}
	if err := ctx.Err(); err != nil {
		return result, err
	}

	cmd := exec.Command(request.Args[0], request.Args[1:]...)
	cmd.Dir = request.Dir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Env = os.Environ()
	for key, value := range request.Env {
		cmd.Env = append(cmd.Env, key+"="+value)
	}
	cmd.Stdin = bytes.NewReader(request.Input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Start(); err != nil {
		return result, err
	}
	finished := make(chan error, 1)
	go func() { finished <- cmd.Wait() }()
	var err error
	select {
	case err = <-finished:
	case <-ctx.Done():
		// A process that leaves this group is outside the Run contract.
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		err = <-finished
		result.Stdout, result.Stderr = stdout.String(), stderr.String()
		result.ExitCode = cmd.ProcessState.ExitCode()
		return result, ctx.Err()
	}
	result.Stdout, result.Stderr = stdout.String(), stderr.String()
	result.ExitCode = cmd.ProcessState.ExitCode()
	if err != nil {
		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
	}
	return result, nil
}

 succeeded in 618ms:
?? DRILL_NOTES.md
?? packages/gitstore/internal/

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/git.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/refs.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..e2b89a1e2dff87ba5fb68f6c858b165eaede1f66
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,7 @@
+## Contradicted Decisions
+
+None identified.
+
+## Gaps
+
+None identified.
diff --git a/packages/gitstore/internal/local/changes.go b/packages/gitstore/internal/local/changes.go
new file mode 100644
index 0000000000000000000000000000000000000000..d009b3e1734ca6d01464877d1be435c383dc05fc
--- /dev/null
+++ b/packages/gitstore/internal/local/changes.go
@@ -0,0 +1,86 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func Changes(root, base string) ([]model.Change, error) {
+	if base == "" {
+		return nil, fmt.Errorf("base ref is required")
+	}
+	byPath := make(map[string]model.Change)
+	order := make([]string, 0)
+	add := func(change model.Change) {
+		if prior, exists := byPath[change.Path]; exists {
+			// Preserve the rename when a later diff also reports its destination.
+			if prior.OldPath != "" && change.OldPath == "" {
+				return
+			}
+		} else {
+			order = append(order, change.Path)
+		}
+		byPath[change.Path] = change
+	}
+	for _, args := range [][]string{
+		{"diff", "--cached", "--no-ext-diff", "--name-status", "-z", "-M", base, "--"},
+		{"diff", "--no-ext-diff", "--name-status", "-z", "-M", "--"},
+	} {
+		data, err := git(root, nil, nil, args...)
+		if err != nil {
+			return nil, err
+		}
+		if err := parseNameStatus(data, add); err != nil {
+			return nil, err
+		}
+	}
+	data, err := git(root, nil, nil, "ls-files", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return nil, err
+	}
+	if len(data) > 0 {
+		for _, path := range strings.Split(strings.TrimSuffix(string(data), "\x00"), "\x00") {
+			if path != "" {
+				add(model.Change{Path: path, Status: "?"})
+			}
+		}
+	}
+	changes := make([]model.Change, 0, len(order))
+	for _, path := range order {
+		changes = append(changes, byPath[path])
+	}
+	return changes, nil
+}
+
+func parseNameStatus(data []byte, add func(model.Change)) error {
+	if len(data) == 0 {
+		return nil
+	}
+	if data[len(data)-1] != 0 {
+		return fmt.Errorf("malformed Git name-status output")
+	}
+	fields := strings.Split(string(data[:len(data)-1]), "\x00")
+	for i := 0; i < len(fields); {
+		status := fields[i]
+		i++
+		if status == "" || i >= len(fields) {
+			return fmt.Errorf("malformed Git name-status entry")
+		}
+		change := model.Change{Status: status[:1]}
+		if change.Status == "R" || change.Status == "C" {
+			if i+1 >= len(fields) {
+				return fmt.Errorf("malformed Git rename entry")
+			}
+			change.OldPath = fields[i]
+			change.Path = fields[i+1]
+			i += 2
+		} else {
+			change.Path = fields[i]
+			i++
+		}
+		add(change)
+	}
+	return nil
+}
diff --git a/packages/gitstore/internal/local/commit.go b/packages/gitstore/internal/local/commit.go
new file mode 100644
index 0000000000000000000000000000000000000000..ddcdc58ed66b34546f5fc194a697f6c2cedfa584
--- /dev/null
+++ b/packages/gitstore/internal/local/commit.go
@@ -0,0 +1,64 @@
+package local
+
+import (
+	"fmt"
+	"os"
+	"path/filepath"
+	"strings"
+)
+
+func Commit(root, subject string, paths []string) (string, error) {
+	if len(paths) == 0 {
+		return "", fmt.Errorf("at least one path is required")
+	}
+	selected := make([]string, 0, len(paths))
+	for _, path := range paths {
+		clean := filepath.Clean(path)
+		if path == "" || filepath.IsAbs(path) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
+			return "", fmt.Errorf("path must be inside the worktree: %q", path)
+		}
+		// Git pathspec magic makes the selection literal, including names which
+		// contain glob characters or begin with a colon.
+		selected = append(selected, ":(literal)"+filepath.ToSlash(clean))
+	}
+	old, err := gitString(root, "rev-parse", "--verify", "HEAD")
+	if err != nil {
+		return "", err
+	}
+	index, err := os.CreateTemp("", "gitstore-index-")
+	if err != nil {
+		return "", err
+	}
+	indexPath := index.Name()
+	_ = index.Close()
+	_ = os.Remove(indexPath)
+	defer os.Remove(indexPath)
+	env := map[string]string{"GIT_INDEX_FILE": indexPath}
+	if _, err := git(root, env, nil, "read-tree", old); err != nil {
+		return "", err
+	}
+	addArgs := append([]string{"add", "-A", "--"}, selected...)
+	if _, err := git(root, env, nil, addArgs...); err != nil {
+		return "", err
+	}
+	treeData, err := git(root, env, nil, "write-tree")
+	if err != nil {
+		return "", err
+	}
+	tree := strings.TrimSpace(string(treeData))
+	commitData, err := git(root, env, nil, "commit-tree", tree, "-p", old, "-m", subject)
+	if err != nil {
+		return "", err
+	}
+	newHead := strings.TrimSpace(string(commitData))
+	if _, err := git(root, nil, nil, "update-ref", "HEAD", newHead, old); err != nil {
+		return "", err
+	}
+	// Bring only selected entries of the real index into line with the new
+	// commit. This leaves every unrelated staged entry untouched.
+	resetArgs := append([]string{"reset", "-q", "HEAD", "--"}, selected...)
+	if _, err := git(root, nil, nil, resetArgs...); err != nil {
+		return newHead, err
+	}
+	return newHead, nil
+}
diff --git a/packages/gitstore/internal/local/git.go b/packages/gitstore/internal/local/git.go
new file mode 100644
index 0000000000000000000000000000000000000000..2678b8042275f0fd1b1dc62a4c425562b84b6ea6
--- /dev/null
+++ b/packages/gitstore/internal/local/git.go
@@ -0,0 +1,23 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func git(root string, env map[string]string, input []byte, args ...string) ([]byte, error) {
+	argv := append([]string{"git", "-C", root}, args...)
+	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input})
+	if err != nil {
+		return []byte(result.Stdout), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
+	}
+	return []byte(result.Stdout), nil
+}
+
+func gitString(root string, args ...string) (string, error) {
+	data, err := git(root, nil, nil, args...)
+	return strings.TrimSuffix(string(data), "\n"), err
+}
diff --git a/packages/gitstore/internal/local/objects.go b/packages/gitstore/internal/local/objects.go
new file mode 100644
index 0000000000000000000000000000000000000000..7889d2b9989c12d98a4b7ad8202d58cc37a8dabe
--- /dev/null
+++ b/packages/gitstore/internal/local/objects.go
@@ -0,0 +1,25 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func WriteBlob(root string, data []byte) (string, error) {
+	out, err := git(root, nil, data, "hash-object", "-w", "--stdin")
+	return strings.TrimSpace(string(out)), err
+}
+
+func ReadBlob(root, oid string) ([]byte, error) {
+	if oid == "" || strings.HasPrefix(oid, "-") {
+		return nil, fmt.Errorf("invalid object ID %q", oid)
+	}
+	return git(root, nil, nil, "cat-file", "blob", oid)
+}
+
+func ReadAt(root, revision, path string) ([]byte, error) {
+	if revision == "" || path == "" {
+		return nil, fmt.Errorf("revision and path are required")
+	}
+	return git(root, nil, nil, "show", revision+":"+path)
+}
diff --git a/packages/gitstore/internal/local/refs.go b/packages/gitstore/internal/local/refs.go
new file mode 100644
index 0000000000000000000000000000000000000000..ba28df8f4fd897c9da0a5fd359a3a8e87ca14404
--- /dev/null
+++ b/packages/gitstore/internal/local/refs.go
@@ -0,0 +1,54 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func checkedRef(root, name string) error {
+	if !strings.HasPrefix(name, "refs/") {
+		return fmt.Errorf("ref must be fully qualified: %q", name)
+	}
+	_, err := gitString(root, "check-ref-format", name)
+	return err
+}
+
+func Ref(root, name string) (string, bool, error) {
+	if err := checkedRef(root, name); err != nil {
+		return "", false, err
+	}
+	result, err := Run(context.Background(), model.Request{
+		Args: []string{"git", "-C", root, "show-ref", "--verify", "--hash", name},
+	})
+	if err == nil {
+		return strings.TrimSuffix(result.Stdout, "\n"), true, nil
+	}
+	if result.ExitCode == 1 {
+		return "", false, nil
+	}
+	return "", false, fmt.Errorf("git show-ref %q: %w: %s", name, err, strings.TrimSpace(result.Stderr))
+}
+
+func CASRef(root, name, old, new string) error {
+	if err := checkedRef(root, name); err != nil {
+		return err
+	}
+	// Git's all-zero object ID is the atomic "must not exist" condition.
+	zero, err := gitString(root, "hash-object", "--stdin")
+	if err != nil {
+		return err
+	}
+	zero = strings.Repeat("0", len(zero))
+	if old == "" {
+		old = zero
+	}
+	if new == "" {
+		_, err = git(root, nil, nil, "update-ref", "-d", name, old)
+		return err
+	}
+	_, err = git(root, nil, nil, "update-ref", name, new, old)
+	return err
+}
diff --git a/packages/gitstore/internal/local/repository.go b/packages/gitstore/internal/local/repository.go
new file mode 100644
index 0000000000000000000000000000000000000000..25ed03220b1baa2a0b48941d3873335b00616399
--- /dev/null
+++ b/packages/gitstore/internal/local/repository.go
@@ -0,0 +1,75 @@
+package local
+
+import (
+	"fmt"
+	"path/filepath"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func canonical(path string) (string, error) {
+	abs, err := filepath.Abs(path)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(abs)
+}
+
+func Inspect(dir string) (model.Repository, error) {
+	var repository model.Repository
+	root, err := gitString(dir, "rev-parse", "--show-toplevel")
+	if err != nil {
+		return repository, err
+	}
+	common, err := gitString(dir, "rev-parse", "--git-common-dir")
+	if err != nil {
+		return repository, err
+	}
+	if !filepath.IsAbs(common) {
+		common = filepath.Join(dir, common)
+	}
+	if repository.Root, err = canonical(root); err != nil {
+		return repository, err
+	}
+	if repository.CommonDir, err = canonical(common); err != nil {
+		return repository, err
+	}
+	if repository.Head, err = gitString(dir, "rev-parse", "--verify", "HEAD"); err != nil {
+		return repository, err
+	}
+	// symbolic-ref exits nonzero for detached HEAD.
+	repository.Branch, _ = gitString(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
+	return repository, nil
+}
+
+func BaseRef(root, remote, base string) (string, error) {
+	if remote == "" {
+		remote = "origin"
+	}
+	if base == "" {
+		head := "refs/remotes/" + remote + "/HEAD"
+		if target, err := gitString(root, "symbolic-ref", "--quiet", head); err == nil {
+			prefix := "refs/remotes/" + remote + "/"
+			if strings.HasPrefix(target, prefix) && len(target) > len(prefix) {
+				base = strings.TrimPrefix(target, prefix)
+			}
+		}
+		if base == "" {
+			base = "main"
+		}
+	}
+	remoteRef := "refs/remotes/" + remote + "/" + base
+	localRef := "refs/heads/" + base
+	for _, name := range []string{remoteRef, localRef} {
+		if _, err := gitString(root, "check-ref-format", name); err != nil {
+			return "", fmt.Errorf("invalid base ref %q: %w", name, err)
+		}
+		if _, exists, err := Ref(root, name); err != nil {
+			return "", err
+		} else if exists {
+			return name, nil
+		}
+	}
+	return "", fmt.Errorf("base branch %q is missing", base)
+}
diff --git a/packages/gitstore/internal/local/run.go b/packages/gitstore/internal/local/run.go
new file mode 100644
index 0000000000000000000000000000000000000000..ce18829152d3e597194e31e16fdb9ac31525c0a3
--- /dev/null
+++ b/packages/gitstore/internal/local/run.go
@@ -0,0 +1,66 @@
+package local
+
+import (
+	"bytes"
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+// Run executes one command in its own process group so cancellation also
+// terminates children which remain in that group.
+func Run(ctx context.Context, request model.Request) (model.Result, error) {
+	var result model.Result
+	if len(request.Args) == 0 {
+		return result, errors.New("empty command")
+	}
+	if ctx == nil {
+		return result, errors.New("nil context")
+	}
+	if request.Timeout > 0 {
+		var cancel context.CancelFunc
+		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
+		defer cancel()
+	}
+	if err := ctx.Err(); err != nil {
+		return result, err
+	}
+
+	cmd := exec.Command(request.Args[0], request.Args[1:]...)
+	cmd.Dir = request.Dir
+	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
+	cmd.Env = os.Environ()
+	for key, value := range request.Env {
+		cmd.Env = append(cmd.Env, key+"="+value)
+	}
+	cmd.Stdin = bytes.NewReader(request.Input)
+	var stdout, stderr bytes.Buffer
+	cmd.Stdout, cmd.Stderr = &stdout, &stderr
+	if err := cmd.Start(); err != nil {
+		return result, err
+	}
+	finished := make(chan error, 1)
+	go func() { finished <- cmd.Wait() }()
+	var err error
+	select {
+	case err = <-finished:
+	case <-ctx.Done():
+		// A process that leaves this group is outside the Run contract.
+		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
+		err = <-finished
+		result.Stdout, result.Stderr = stdout.String(), stderr.String()
+		result.ExitCode = cmd.ProcessState.ExitCode()
+		return result, ctx.Err()
+	}
+	result.Stdout, result.Stderr = stdout.String(), stderr.String()
+	result.ExitCode = cmd.ProcessState.ExitCode()
+	if err != nil {
+		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
+	}
+	return result, nil
+}

exec
/bin/zsh -lc 'gofmt -w packages/gitstore/internal/local/*.go && env GO111MODULE=off GOPATH=/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/.test-gopath GOCACHE=/tmp/fireweed-rebuild-cache go test github.com/lurkshark/fireweed/packages/gitstore/evals -count=1' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 exited 1 in 6613ms:
--- FAIL: Test_GIT_002_BaseRefPreferenceDefaultsAndMissing (0.43s)
    gitstore_test.go:109: default main = "", git show-ref "refs/remotes/origin/main": command "git" exited with code 128: exit status 128: fatal: 'refs/remotes/origin/main' - not a valid ref
    gitstore_test.go:113: local fallback = "", git show-ref "refs/remotes/upstream/release": command "git" exited with code 128: exit status 128: fatal: 'refs/remotes/upstream/release' - not a valid ref
--- FAIL: Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS (0.43s)
    gitstore_test.go:282: absent Ref = "", false, git show-ref "refs/fireweed/eval-claim": command "git" exited with code 128: exit status 128: fatal: 'refs/fireweed/eval-claim' - not a valid ref
    gitstore_test.go:319: deleted Ref = "", false, git show-ref "refs/fireweed/eval-claim": command "git" exited with code 128: exit status 128: fatal: 'refs/fireweed/eval-claim' - not a valid ref
FAIL
FAIL	github.com/lurkshark/fireweed/packages/gitstore/evals	5.096s
FAIL

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/refs.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..e2b89a1e2dff87ba5fb68f6c858b165eaede1f66
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,7 @@
+## Contradicted Decisions
+
+None identified.
+
+## Gaps
+
+None identified.
diff --git a/packages/gitstore/internal/local/changes.go b/packages/gitstore/internal/local/changes.go
new file mode 100644
index 0000000000000000000000000000000000000000..d009b3e1734ca6d01464877d1be435c383dc05fc
--- /dev/null
+++ b/packages/gitstore/internal/local/changes.go
@@ -0,0 +1,86 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func Changes(root, base string) ([]model.Change, error) {
+	if base == "" {
+		return nil, fmt.Errorf("base ref is required")
+	}
+	byPath := make(map[string]model.Change)
+	order := make([]string, 0)
+	add := func(change model.Change) {
+		if prior, exists := byPath[change.Path]; exists {
+			// Preserve the rename when a later diff also reports its destination.
+			if prior.OldPath != "" && change.OldPath == "" {
+				return
+			}
+		} else {
+			order = append(order, change.Path)
+		}
+		byPath[change.Path] = change
+	}
+	for _, args := range [][]string{
+		{"diff", "--cached", "--no-ext-diff", "--name-status", "-z", "-M", base, "--"},
+		{"diff", "--no-ext-diff", "--name-status", "-z", "-M", "--"},
+	} {
+		data, err := git(root, nil, nil, args...)
+		if err != nil {
+			return nil, err
+		}
+		if err := parseNameStatus(data, add); err != nil {
+			return nil, err
+		}
+	}
+	data, err := git(root, nil, nil, "ls-files", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return nil, err
+	}
+	if len(data) > 0 {
+		for _, path := range strings.Split(strings.TrimSuffix(string(data), "\x00"), "\x00") {
+			if path != "" {
+				add(model.Change{Path: path, Status: "?"})
+			}
+		}
+	}
+	changes := make([]model.Change, 0, len(order))
+	for _, path := range order {
+		changes = append(changes, byPath[path])
+	}
+	return changes, nil
+}
+
+func parseNameStatus(data []byte, add func(model.Change)) error {
+	if len(data) == 0 {
+		return nil
+	}
+	if data[len(data)-1] != 0 {
+		return fmt.Errorf("malformed Git name-status output")
+	}
+	fields := strings.Split(string(data[:len(data)-1]), "\x00")
+	for i := 0; i < len(fields); {
+		status := fields[i]
+		i++
+		if status == "" || i >= len(fields) {
+			return fmt.Errorf("malformed Git name-status entry")
+		}
+		change := model.Change{Status: status[:1]}
+		if change.Status == "R" || change.Status == "C" {
+			if i+1 >= len(fields) {
+				return fmt.Errorf("malformed Git rename entry")
+			}
+			change.OldPath = fields[i]
+			change.Path = fields[i+1]
+			i += 2
+		} else {
+			change.Path = fields[i]
+			i++
+		}
+		add(change)
+	}
+	return nil
+}
diff --git a/packages/gitstore/internal/local/commit.go b/packages/gitstore/internal/local/commit.go
new file mode 100644
index 0000000000000000000000000000000000000000..ddcdc58ed66b34546f5fc194a697f6c2cedfa584
--- /dev/null
+++ b/packages/gitstore/internal/local/commit.go
@@ -0,0 +1,64 @@
+package local
+
+import (
+	"fmt"
+	"os"
+	"path/filepath"
+	"strings"
+)
+
+func Commit(root, subject string, paths []string) (string, error) {
+	if len(paths) == 0 {
+		return "", fmt.Errorf("at least one path is required")
+	}
+	selected := make([]string, 0, len(paths))
+	for _, path := range paths {
+		clean := filepath.Clean(path)
+		if path == "" || filepath.IsAbs(path) || clean == "." || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
+			return "", fmt.Errorf("path must be inside the worktree: %q", path)
+		}
+		// Git pathspec magic makes the selection literal, including names which
+		// contain glob characters or begin with a colon.
+		selected = append(selected, ":(literal)"+filepath.ToSlash(clean))
+	}
+	old, err := gitString(root, "rev-parse", "--verify", "HEAD")
+	if err != nil {
+		return "", err
+	}
+	index, err := os.CreateTemp("", "gitstore-index-")
+	if err != nil {
+		return "", err
+	}
+	indexPath := index.Name()
+	_ = index.Close()
+	_ = os.Remove(indexPath)
+	defer os.Remove(indexPath)
+	env := map[string]string{"GIT_INDEX_FILE": indexPath}
+	if _, err := git(root, env, nil, "read-tree", old); err != nil {
+		return "", err
+	}
+	addArgs := append([]string{"add", "-A", "--"}, selected...)
+	if _, err := git(root, env, nil, addArgs...); err != nil {
+		return "", err
+	}
+	treeData, err := git(root, env, nil, "write-tree")
+	if err != nil {
+		return "", err
+	}
+	tree := strings.TrimSpace(string(treeData))
+	commitData, err := git(root, env, nil, "commit-tree", tree, "-p", old, "-m", subject)
+	if err != nil {
+		return "", err
+	}
+	newHead := strings.TrimSpace(string(commitData))
+	if _, err := git(root, nil, nil, "update-ref", "HEAD", newHead, old); err != nil {
+		return "", err
+	}
+	// Bring only selected entries of the real index into line with the new
+	// commit. This leaves every unrelated staged entry untouched.
+	resetArgs := append([]string{"reset", "-q", "HEAD", "--"}, selected...)
+	if _, err := git(root, nil, nil, resetArgs...); err != nil {
+		return newHead, err
+	}
+	return newHead, nil
+}
diff --git a/packages/gitstore/internal/local/git.go b/packages/gitstore/internal/local/git.go
new file mode 100644
index 0000000000000000000000000000000000000000..2678b8042275f0fd1b1dc62a4c425562b84b6ea6
--- /dev/null
+++ b/packages/gitstore/internal/local/git.go
@@ -0,0 +1,23 @@
+package local
+
+import (
+	"context"
+	"fmt"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func git(root string, env map[string]string, input []byte, args ...string) ([]byte, error) {
+	argv := append([]string{"git", "-C", root}, args...)
+	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input})
+	if err != nil {
+		return []byte(result.Stdout), fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
+	}
+	return []byte(result.Stdout), nil
+}
+
+func gitString(root string, args ...string) (string, error) {
+	data, err := git(root, nil, nil, args...)
+	return strings.TrimSuffix(string(data), "\n"), err
+}
diff --git a/packages/gitstore/internal/local/objects.go b/packages/gitstore/internal/local/objects.go
new file mode 100644
index 0000000000000000000000000000000000000000..7889d2b9989c12d98a4b7ad8202d58cc37a8dabe
--- /dev/null
+++ b/packages/gitstore/internal/local/objects.go
@@ -0,0 +1,25 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func WriteBlob(root string, data []byte) (string, error) {
+	out, err := git(root, nil, data, "hash-object", "-w", "--stdin")
+	return strings.TrimSpace(string(out)), err
+}
+
+func ReadBlob(root, oid string) ([]byte, error) {
+	if oid == "" || strings.HasPrefix(oid, "-") {
+		return nil, fmt.Errorf("invalid object ID %q", oid)
+	}
+	return git(root, nil, nil, "cat-file", "blob", oid)
+}
+
+func ReadAt(root, revision, path string) ([]byte, error) {
+	if revision == "" || path == "" {
+		return nil, fmt.Errorf("revision and path are required")
+	}
+	return git(root, nil, nil, "show", revision+":"+path)
+}
diff --git a/packages/gitstore/internal/local/refs.go b/packages/gitstore/internal/local/refs.go
new file mode 100644
index 0000000000000000000000000000000000000000..d91c313a0f3c6af146a5c0190b2a8249367f4504
--- /dev/null
+++ b/packages/gitstore/internal/local/refs.go
@@ -0,0 +1,55 @@
+package local
+
+import (
+	"fmt"
+	"strings"
+)
+
+func checkedRef(root, name string) error {
+	if !strings.HasPrefix(name, "refs/") {
+		return fmt.Errorf("ref must be fully qualified: %q", name)
+	}
+	_, err := gitString(root, "check-ref-format", name)
+	return err
+}
+
+func Ref(root, name string) (string, bool, error) {
+	if err := checkedRef(root, name); err != nil {
+		return "", false, err
+	}
+	// for-each-ref returns successfully for an absent ref on both older and
+	// newer Git versions. Match the full name because its pattern may also
+	// select refs below the requested namespace.
+	out, err := gitString(root, "for-each-ref", "--format=%(refname) %(objectname)", name)
+	if err != nil {
+		return "", false, err
+	}
+	for _, line := range strings.Split(out, "\n") {
+		fields := strings.Fields(line)
+		if len(fields) == 2 && fields[0] == name {
+			return fields[1], true, nil
+		}
+	}
+	return "", false, nil
+}
+
+func CASRef(root, name, old, new string) error {
+	if err := checkedRef(root, name); err != nil {
+		return err
+	}
+	// Git's all-zero object ID is the atomic "must not exist" condition.
+	zero, err := gitString(root, "hash-object", "--stdin")
+	if err != nil {
+		return err
+	}
+	zero = strings.Repeat("0", len(zero))
+	if old == "" {
+		old = zero
+	}
+	if new == "" {
+		_, err = git(root, nil, nil, "update-ref", "-d", name, old)
+		return err
+	}
+	_, err = git(root, nil, nil, "update-ref", name, new, old)
+	return err
+}
diff --git a/packages/gitstore/internal/local/repository.go b/packages/gitstore/internal/local/repository.go
new file mode 100644
index 0000000000000000000000000000000000000000..25ed03220b1baa2a0b48941d3873335b00616399
--- /dev/null
+++ b/packages/gitstore/internal/local/repository.go
@@ -0,0 +1,75 @@
+package local
+
+import (
+	"fmt"
+	"path/filepath"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+func canonical(path string) (string, error) {
+	abs, err := filepath.Abs(path)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(abs)
+}
+
+func Inspect(dir string) (model.Repository, error) {
+	var repository model.Repository
+	root, err := gitString(dir, "rev-parse", "--show-toplevel")
+	if err != nil {
+		return repository, err
+	}
+	common, err := gitString(dir, "rev-parse", "--git-common-dir")
+	if err != nil {
+		return repository, err
+	}
+	if !filepath.IsAbs(common) {
+		common = filepath.Join(dir, common)
+	}
+	if repository.Root, err = canonical(root); err != nil {
+		return repository, err
+	}
+	if repository.CommonDir, err = canonical(common); err != nil {
+		return repository, err
+	}
+	if repository.Head, err = gitString(dir, "rev-parse", "--verify", "HEAD"); err != nil {
+		return repository, err
+	}
+	// symbolic-ref exits nonzero for detached HEAD.
+	repository.Branch, _ = gitString(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
+	return repository, nil
+}
+
+func BaseRef(root, remote, base string) (string, error) {
+	if remote == "" {
+		remote = "origin"
+	}
+	if base == "" {
+		head := "refs/remotes/" + remote + "/HEAD"
+		if target, err := gitString(root, "symbolic-ref", "--quiet", head); err == nil {
+			prefix := "refs/remotes/" + remote + "/"
+			if strings.HasPrefix(target, prefix) && len(target) > len(prefix) {
+				base = strings.TrimPrefix(target, prefix)
+			}
+		}
+		if base == "" {
+			base = "main"
+		}
+	}
+	remoteRef := "refs/remotes/" + remote + "/" + base
+	localRef := "refs/heads/" + base
+	for _, name := range []string{remoteRef, localRef} {
+		if _, err := gitString(root, "check-ref-format", name); err != nil {
+			return "", fmt.Errorf("invalid base ref %q: %w", name, err)
+		}
+		if _, exists, err := Ref(root, name); err != nil {
+			return "", err
+		} else if exists {
+			return name, nil
+		}
+	}
+	return "", fmt.Errorf("base branch %q is missing", base)
+}
diff --git a/packages/gitstore/internal/local/run.go b/packages/gitstore/internal/local/run.go
new file mode 100644
index 0000000000000000000000000000000000000000..ce18829152d3e597194e31e16fdb9ac31525c0a3
--- /dev/null
+++ b/packages/gitstore/internal/local/run.go
@@ -0,0 +1,66 @@
+package local
+
+import (
+	"bytes"
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"os/exec"
+	"syscall"
+
+	"github.com/lurkshark/fireweed/packages/gitstore/model"
+)
+
+// Run executes one command in its own process group so cancellation also
+// terminates children which remain in that group.
+func Run(ctx context.Context, request model.Request) (model.Result, error) {
+	var result model.Result
+	if len(request.Args) == 0 {
+		return result, errors.New("empty command")
+	}
+	if ctx == nil {
+		return result, errors.New("nil context")
+	}
+	if request.Timeout > 0 {
+		var cancel context.CancelFunc
+		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
+		defer cancel()
+	}
+	if err := ctx.Err(); err != nil {
+		return result, err
+	}
+
+	cmd := exec.Command(request.Args[0], request.Args[1:]...)
+	cmd.Dir = request.Dir
+	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
+	cmd.Env = os.Environ()
+	for key, value := range request.Env {
+		cmd.Env = append(cmd.Env, key+"="+value)
+	}
+	cmd.Stdin = bytes.NewReader(request.Input)
+	var stdout, stderr bytes.Buffer
+	cmd.Stdout, cmd.Stderr = &stdout, &stderr
+	if err := cmd.Start(); err != nil {
+		return result, err
+	}
+	finished := make(chan error, 1)
+	go func() { finished <- cmd.Wait() }()
+	var err error
+	select {
+	case err = <-finished:
+	case <-ctx.Done():
+		// A process that leaves this group is outside the Run contract.
+		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
+		err = <-finished
+		result.Stdout, result.Stderr = stdout.String(), stderr.String()
+		result.ExitCode = cmd.ProcessState.ExitCode()
+		return result, ctx.Err()
+	}
+	result.Stdout, result.Stderr = stdout.String(), stderr.String()
+	result.ExitCode = cmd.ProcessState.ExitCode()
+	if err != nil {
+		return result, fmt.Errorf("command %q exited with code %d: %w", request.Args[0], result.ExitCode, err)
+	}
+	return result, nil
+}

exec
/bin/zsh -lc 'gofmt -w packages/gitstore/internal/local/refs.go && env GO111MODULE=off GOPATH=/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local/.test-gopath GOCACHE=/tmp/fireweed-rebuild-cache go test github.com/lurkshark/fireweed/packages/gitstore/evals -count=1' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 6454ms:
ok  	github.com/lurkshark/fireweed/packages/gitstore/evals	5.100s

2026-11-01T17:08:45.508181Z ERROR codex_core::tools::router: error=exec_command failed: CreateProcess { message: "Rejected(\"`/bin/zsh -lc 'rm -rf packages/gitstore/internal/local/.test-gopath && git status --short && find packages/gitstore/internal -type f -print | sort && git diff --check'` rejected: rm -f style commands are not permitted. Use a safer approach\")" }
exec
/bin/zsh -lc 'unlink packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages/gitstore && rmdir packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed/packages packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark/fireweed packages/gitstore/internal/local/.test-gopath/src/github.com/lurkshark packages/gitstore/internal/local/.test-gopath/src/github.com packages/gitstore/internal/local/.test-gopath/src packages/gitstore/internal/local/.test-gopath' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 484ms:
exec
/bin/zsh -lc 'cat DRILL_NOTES.md && git status --porcelain=v1 packages/gitstore/SPEC.md packages/gitstore/api.go packages/gitstore/model/model.go packages/gitstore/evals/gitstore_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
exec
/bin/zsh -lc 'git status --short && find packages/gitstore/internal -type f -print | sort && git diff --check' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844
 succeeded in 524ms:
## Contradicted Decisions

None identified.

## Gaps

None identified.

 succeeded in 539ms:
?? DRILL_NOTES.md
?? packages/gitstore/internal/
packages/gitstore/internal/local/changes.go
packages/gitstore/internal/local/commit.go
packages/gitstore/internal/local/git.go
packages/gitstore/internal/local/objects.go
packages/gitstore/internal/local/refs.go
packages/gitstore/internal/local/repository.go
packages/gitstore/internal/local/run.go

codex
Rebuilt gitstore in [packages/gitstore/internal/local](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/packages/gitstore/internal/local). The retained evals pass. [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-441413844/DRILL_NOTES.md) records no identified contradictions or gaps. Retained inputs are unchanged, and nothing was committed.
tokens used
58,478

