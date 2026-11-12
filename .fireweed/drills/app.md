---
date: 2026-11-10T20:15:33.930973Z
base_commit: 2ab6c98e9daed015b22dd8819f537f077c8edd79
agent: codex
model: gpt-6.1-sol
result: pass
---

## Own evals
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
- app | Test_APP_014_MaintenanceCommandsDispatch | APP-014 | pass
- app | Test_APP_014_MaintenanceCommandsDispatch/gc | APP-014 | pass
- app | Test_APP_014_MaintenanceCommandsDispatch/migrate | APP-014 | pass
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
- app | Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases | APP-147, APP-148 | pass
- app | Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp | APP-147, APP-148 | pass
- app | Test_APP_148_APP_149_CheckFixRejectsBeforeMutation | APP-148, APP-149 | pass
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
- app | Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty | APP-150, APP-151 | pass
- app | Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/plain | APP-150, APP-151 | pass
- app | Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/JSON | APP-150, APP-151 | pass
- app | Test_APP_150_APP_152_GCRejectsUsageBeforePruning | APP-150, APP-152 | pass
- app | Test_APP_152_GCConfigurationAndGitFailures | APP-152 | pass
- app | Test_APP_152_GCConfigurationAndGitFailures/malformed-config | APP-152 | pass
- app | Test_APP_152_GCConfigurationAndGitFailures/invalid-retention | APP-152 | pass
- app | Test_APP_041_APP_042_APP_043_APP_044_APP_045_APP_046_APP_050_APP_053_APP_054_APP_059_GoInit | APP-041, APP-042, APP-043, APP-044, APP-045, APP-046, APP-050, APP-053, APP-054, APP-059 | pass
- app | Test_APP_043_APP_044_APP_047_APP_051_NoneInit | APP-043, APP-044, APP-047, APP-051 | pass
- app | Test_APP_050_APP_052_NoneDefaultApp | APP-050, APP-052 | pass
- app | Test_APP_053_InitFileOnlyWorkerCommand | APP-053 | pass
- app | Test_APP_054_InitFileOnlyAuditCommand | APP-054 | pass
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
- app | Test_APP_153_APP_154_MigrateNearestRootAndRepeatedNoOp | APP-153, APP-154 | pass
- app | Test_APP_153_APP_154_MigrateEmptyCurrentProject | APP-153, APP-154 | pass
- app | Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures | APP-154, APP-155 | pass
- app | Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/static-check | APP-154, APP-155 | pass
- app | Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/malformed-config | APP-154, APP-155 | pass
- app | Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/malformed-module | APP-154, APP-155 | pass
- app | Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/unsupported-version | APP-154, APP-155 | pass
- app | Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/missing-version | APP-154, APP-155 | pass
- app | Test_APP_155_MigrateRejectsArgumentsBeforeMutation | APP-155 | pass
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
- app | Test_APP_143_APP_144_StaleRecoveryDiffAndReadOnly | APP-143, APP-144 | pass
- app | Test_APP_144_StaleUnavailableHistoryAndEmptyView | APP-144 | pass
- app | Test_APP_143_StaleViewChoosesNearestProject | APP-143 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures | APP-145, APP-146 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/unknown | APP-145, APP-146 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/ambiguous | APP-145, APP-146 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/bad-bases | APP-145, APP-146 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/missing-config | APP-145, APP-146 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/malformed-config | APP-145, APP-146 | pass
- app | Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/invalid-graph | APP-145, APP-146 | pass
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
ok  	github.com/lurkshark/fireweed/packages/app/evals	261.296s


## Dependent evals

## Contradicted Decisions
None found.
## Gaps
- **APP-004 and APP-012, app.** The retained workflow API resolves a claim through the live queue. If the pending document or module configuration can no longer be parsed, the app cannot get the approved snapshot through that API. The specs do not define how to combine approved metadata with a projection requested for a different change.
- **APP-037, APP-078 and APP-085, app.** Some project errors carry their codes privately. The CLI recovers those codes by matching the cause against public `project.Check` diagnostics. The retained interfaces do not provide a general way to read a code from those errors after wrapping.
- **APP-067, app.** The retained eval expects short Doctor probe names such as `baseline` and `recovery`. The checks API includes the hook scope in the name. JSON uses the short name and keeps the scope in the detail.
- **APP-103, app.** The specs leave the accepted module name and prefix grammar implicit. The implementation uses lowercase letter-led names with digits and hyphenated segments, and uppercase letter-led prefixes with optional digits.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt the Fireweed app CLI in `packages/app/internal/entry`. The target eval hook passed all **323 cases**. The changes are uncommitted. [Eval report](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/.fireweed/runs/drill/report.xml)

The controller’s `verify` returned P02 because the drill workspace has no held change.

[DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/DRILL_NOTES.md) records the approved-snapshot lookup limitation, diagnostic recovery, Doctor JSON naming, and module-name validation. It records no contradicted Decisions.

Reading additional input from stdin...
2026-11-10T20:15:36.243712Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1c71a-0cec-7803-9eff-7dbaaa28ea73
--------
user
# Rebuild app

## Target spec and Decisions

---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
    - runner
id_prefix: APP
name: app
---

# app

## Specs

### APP-001
The CLI MUST run help, version, explain <code> and playbook <name> successfully
outside a project, returning the corresponding assets text on stdout with exit
0 and empty stderr. These commands MUST also work in a directory containing an
invalid fireweed.yaml.

### APP-002
Each information command MUST accept --json and emit one JSON object with schema
1, string data equal to its plain text result, and an empty diagnostics array.
It MUST return exit 0 and leave stderr empty.

### APP-003
Unknown commands or options and missing or extra arguments MUST return exit 2 with a nonempty corrective U01 diagnostic.

### APP-004
modules MUST report the validated graph, using the held approved projected snapshot while this worktree owns a change.

### APP-005
new-change <module> <slug> MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0.

### APP-006
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned.

### APP-007
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail.

### APP-008
Without --history or --stale, show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected.

### APP-009
Without --from-attempt, start <change> MUST discover the root, invoke
workflow.Start and print first-step workflow.Context with exit 0, assigning IDs,
run and phase without committing.

### APP-010
verify [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-011
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics.

### APP-012
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable.

### APP-013
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective U02 diagnostic.

### APP-014
gc and migrate listed by help MUST dispatch supported maintenance operations instead of reporting that they are unimplemented or under construction.

### APP-015
Failed --json requests MUST use the schema-1 envelope with null data and a diagnostic containing code and message.

### APP-016
modules MUST discover the nearest ancestor fireweed.yaml.

### APP-017
modules --json MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics on success.

### APP-018
modules MUST return exit 0 on success or exit 2 for missing or invalid project inputs, with a corrective failure diagnostic: C01 for configuration, C02 for graph, C03 for document syntax or C12 for metadata errors.

### APP-019
new-change MUST return exit 2 with corrective diagnostics U01 for invalid arguments or slug, U02 for an unknown module and C13 for duplicate slugs, preserving existing files.

### APP-020
new-change --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

### APP-021
stamp MUST return exit 2 for usage or configuration errors and exit 1 for stamping rejections, with code/message diagnostics.

### APP-022
stamp --json MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### APP-023
check MUST return exit 2 for usage or configuration errors with code/message diagnostics.

### APP-024
check --json MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### APP-025
Without --history, show --json MUST return identical text as schema-1 string data with empty diagnostics and exit 0 on success.

### APP-026
show MUST return exit 2 with code/message diagnostics for unknown or ambiguous module or item, invalid arguments or configuration errors.

### APP-027
start MUST preserve reported warnings and return exit 1 for lifecycle rejection or exit 2 for usage or configuration errors, with corrective code/message diagnostics.

### APP-028
start --json MUST return schema 1 with context string data on success or null data on failure.

### APP-029
done [change] MUST invoke full workflow.Done with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-030
verify --quick MUST use quick checks without committing.

### APP-031
verify and done --json MUST return schema 1 with result data on success or null on failure and reported code/message diagnostics.

### APP-032
context --json MUST return identical workflow.Context text as schema-1 string data.

### APP-033
context MUST return exit 2 for usage or configuration errors and exit 1 for held-state lookup failures, with corrective code/message diagnostics and null JSON data.

### APP-034
coverage MUST map results to IDs and require coverage and pass, returning exit 0 on success or exit 1 on failed checks.

### APP-035
coverage MUST return exit 2 with corrective diagnostics for usage, configuration or unknown module/change errors.

### APP-036
coverage --json MUST return schema 1 with report data on success, null on failure and code/message diagnostics.

### APP-037
Existing static failures exposed by the CLI MUST retain their D006 cause: C01 configuration, C02 graph, C03 syntax, C04 normative text, C05 identities, C06 tags, C07 change structure, C08 bases, C09 handles, C12 metadata or C13 queue identity.

### APP-038
CLI lookup rejections for unknown names or ambiguous projected change-local handles MUST use U02. Invalid graphs, module spec identities and other static validation failures MUST retain their D006 static codes.

### APP-039
Existing operational precondition failures without a more specific D006 cause MUST report P01.

### APP-040
For workflow or checks failures carrying an assigned code, the CLI MUST preserve that code and reported warnings.

### APP-041
init MUST create fireweed.yaml in the current directory with format version 1, a binary range >=MAJOR.MINOR, <MAJOR.(MINOR+1) derived from assets.Version, base main, remote origin and modules [packages/*].

### APP-042
init MUST write the selected recipe's module_defaults, hooks, protected and worker_writable values into the generated configuration.

### APP-043
init MUST create an editable SYSTEM.md project scaffold when absent and retain an existing SYSTEM.md unchanged.

### APP-044
init MUST create or update the fireweed managed AGENTS.md block, pointing to author-change and context, while preserving text outside the markers.

### APP-045
init --recipe <name> MUST accept any currently available recipe returned by assets.BuiltinRecipe(name) and copy its Files into their project-relative destinations without changing their contents.

### APP-046
Successful init MUST print the selected recipe's instructions on stdout and return exit 0 with empty stderr.

### APP-047
init --recipe none MUST write minimal setup without language hooks, templates or canaries and print the eval, boundaries and setup hook contract.

### APP-048
init --recipe ./directory MUST load recipe.yaml module_defaults, hooks, protected and worker_writable metadata and recipe.md instructions from that directory.

### APP-049
init with a local recipe MUST copy the directory's files/ contents into matching project-relative destinations and copy its instructions to .fireweed/recipes/local.md.

### APP-050
By default init MUST create a normal packages/app module named app with prefix APP and no dependencies, rendering available module templates with name app, prefix APP and package app.

### APP-051
init --no-app MUST skip creation of the app module.

### APP-052
When the selected recipe supplies no module templates, init MUST create a minimal app SPEC.md with empty Specs and Decisions sections.

### APP-053
init MUST configure codex or claude argv command presets, or preserve a custom agent command template containing {prompt} or {prompt_file} as a shell string, in agent.command.

### APP-054
init MUST enable audit by default and write a separately selected --audit-agent command when supplied, accepting the same presets and custom template forms as --agent, otherwise leaving audit.command unset for agent fallback.

### APP-055
init MUST reject missing --recipe or --agent values, malformed options or extra positional arguments with corrective U01 and exit 2 before creating files; missing-input messages MUST list valid recipe and agent choices.

### APP-056
init MUST return U02 and exit 2 for an unknown recipe name, C03 for unreadable or malformed local recipe documents, and C12 for invalid local recipe metadata types, without creating setup files.

### APP-057
init MUST refuse an existing fireweed.yaml with P01 and exit 2, preserving existing project files.

### APP-058
init MUST refuse conflicting generated recipe or module files with P01 and exit 2 rather than overwrite them.

### APP-059
init MUST retain existing .gitignore entries and ensure .fireweed/runs/ is ignored so lifecycle reports do not become forbidden worktree edits.

### APP-060
init MUST copy setup and print instructions without executing configured agents or hooks.

### APP-061
init --json MUST return schema 1 with instruction text as string data and empty diagnostics on success, or null data and code/message diagnostics on failure, with empty stderr.

### APP-062
doctor MUST discover the nearest ancestor fireweed.yaml and invoke checks.Doctor for that project root.

### APP-063
doctor MUST supply the actual invoking executable as DoctorRequest.Bin so scratch hooks query the running NEW CLI.

### APP-064
doctor --no-agent MUST skip configured worker and audit commands while retaining hook canary and recovery checks.

### APP-065
Plain doctor output MUST identify each reported probe's outcome and retain reported diagnostic codes, warning status and named failure details.

### APP-066
doctor MUST return exit 0 when no error diagnostics remain, including warning-only results, exit 1 for static or probe failures, and exit 2 for usage or configuration-loading errors.

### APP-067
doctor --json MUST emit schema 1 with an object containing a checks array on success or null data on failure, retaining code/message/path/warning diagnostics and leaving stderr empty.

### APP-068
agents-md --write MUST discover the nearest ancestor fireweed.yaml and refresh AGENTS.md at that project root.

### APP-069
agents-md --write MUST regenerate the current fireweed managed block while preserving text outside its markers exactly.

### APP-070
agents-md --write MUST create AGENTS.md when absent or append the managed block when neither marker exists.

### APP-071
agents-md --write MUST reject unpaired managed-block markers with corrective P01 and exit 2, preserving the existing text.

### APP-072
agents-md --write MUST print confirmation and exit 0 on success, or exit 2 with corrective U01 for usage, C01 for missing/invalid fireweed.yaml or P01 for operational file/setup failures.

### APP-073
agents-md --write --json MUST emit schema 1 with confirmation string data and empty diagnostics on success, or null data and code/message diagnostics on failure, leaving stderr empty.

### APP-074
next MUST discover the nearest ancestor fireweed.yaml and invoke runner.Next
for that project, supplying the actual invoking executable as Bin.

### APP-075
next MUST accept --allow-unmerged, passing true only when present. It MUST
reject extra positional arguments and unsupported options with corrective U01
and exit 2; --allow-unmerged MUST be rejected for unrelated implemented commands.

### APP-076
Plain next MUST identify the applied change or empty queue after successful
runner completion. On a clean usage stop it MUST name the held change and show
the wait/budget reason with instructions for resuming it.

### APP-077
next MUST return exit0 for an applied change or empty queue, exit1 for ordinary
runner execution or lifecycle failure, exit2 for usage or configuration errors,
and exit3 for a clean usage stop.

### APP-078
next failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-079
next --json MUST emit schema1 and leave stderr empty. It MUST use object data
containing change, run, phase and empty on success or a clean usage stop, with
a nonempty reason reporting completed wait, next wait and budget on a stop.
On failure it MUST use null data with code/message diagnostics.

### APP-080
run MUST discover the nearest ancestor fireweed.yaml and invoke runner.Run for
that project with the actual invoking executable as Bin.

### APP-081
run --max N MUST accept a positive integer and supply it as Max, using zero
when --max is absent. It MUST reject missing, noninteger or nonpositive values,
extra positional arguments and unsupported options with corrective U01 and
exit 2; --max MUST be rejected for unrelated implemented commands.

### APP-082
run MUST accept --allow-unmerged and supply AllowUnmerged true only when the
flag is present, using false by default.

### APP-083
After successful runner completion, plain run MUST identify applied changes
and their count, including whether the queue is empty or Max left queued work.
On a clean usage stop it MUST retain the applied names/count and show the
wait/budget reason with instructions for resuming.

### APP-084
run MUST return exit0 for successful queue completion, Max termination or an
empty queue, exit1 for execution or lifecycle failure including any earlier
blocked change, exit2 for usage or configuration errors, and exit3 for a clean
time or usage stop.

### APP-085
run failure diagnostics MUST retain existing cause codes from wrapped
dependency errors and include a corrective action.

### APP-086
run --json MUST emit schema1 and leave stderr empty. It MUST use object data
containing applied change/run/phase summaries, empty and limited on success or
a clean usage stop, with a nonempty reason reporting completed wait, next wait
and budget on a stop. On failure it MUST use null data with code/message
diagnostics.

### APP-087
start <change> --from-attempt N MUST discover the project root and invoke
workflow.StartFromAttempt with the selected change and N. On success it MUST
return workflow.Context in the existing plain or JSON start shape with exit 0.

### APP-088
For start, --from-attempt MUST require one positive integer and reject missing,
noninteger, nonpositive or repeated values with corrective U01 and exit 2 before
mutation. Its value diagnostics MUST describe valid start usage. Other commands
MUST reject --from-attempt with U01 and exit 2.

### APP-089
run --for D MUST require one positive Go duration and pass it as RunRequest.For,
using zero when absent. It MUST reject missing, malformed, nonpositive or repeated
values and use by other commands with corrective U01 and exit2 before mutation,
showing valid run usage.

### APP-090
On a clean runner time stop, run MUST return exit3 with retained applied summaries
and a nonempty reason naming the time limit. Plain output MUST include the
applied count and resume guidance. JSON MUST use the existing schema1 run object
with reason and empty/limited flags, leaving stderr empty.

### APP-091
release <change> MUST discover the nearest project root and invoke
workflow.Release for the named change. On success it MUST print confirmation
with exit0, preserving work, runtime state and HEAD while dropping only the
owned claim.

### APP-092
release MUST return exit1 with the workflow cause code and corrective action
when claim release fails, including missing or foreign claims. It MUST return
exit2 with a corrective diagnostic for usage or project configuration errors.

### APP-093
release --json MUST emit schema1 with confirmation string data and empty
diagnostics on success, or null data with code/message diagnostics on failure.
It MUST leave stderr empty.

### APP-094
commits <change> MUST discover the nearest project root and invoke
workflow.Commits. Plain success MUST print the returned object IDs one per line
in returned order with exit0, returning empty output when the list is empty.

### APP-095
commits MUST return exit2 with corrective U01 for usage errors, U02 for selector
errors, and the existing cause code for project configuration errors. Other
workflow failures MUST return exit1 retaining their cause code and a corrective
action.

### APP-096
commits --json MUST emit schema1 with an array of returned object IDs, including
an empty array for no commits, and empty diagnostics on success. On failure it
MUST emit null data with code/message diagnostics. It MUST leave stderr empty.

### APP-097
abandon <change> MUST discover the nearest project root and invoke workflow.Abandon for the named change.
Plain success MUST print confirmation naming the change and returned full commit object ID with exit0.

### APP-098
abandon MUST return exit2 with corrective U01 for usage errors, U02 for selector errors, and the existing cause code for project configuration errors.
Other workflow failures MUST return exit1 retaining their cause code and a corrective action.

### APP-099
abandon --json MUST emit schema1 with confirmation string data and empty diagnostics on success.
On failure it MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-100
new-module <name> --prefix <P> MUST discover the nearest project root and create the named module from .fireweed/templates/module in a new directory matched by the configured module globs.
It MUST use that project-owned copy rather than replace it with embedded templates.

### APP-101
Module creation MUST render template contents using name, prefix, package, module_dir and go_module substitutions, removing a terminal .tmpl suffix from output filenames.
The resulting SPEC.md MUST declare the requested module name and prefix and be loadable as part of the project graph.

### APP-102
new-module MUST reject an existing module name, existing ID prefix or occupied destination before writing module files, preserving the existing module and destination contents.

### APP-103
new-module MUST return exit2 with corrective U01 for invalid arguments and options, the existing cause code for project errors, C03 for missing or malformed templates, and P01 for conflicting module identity or destination.
It MUST reject requests before writing module files when these errors are present.

### APP-104
Successful new-module MUST return exit0 and print the repository-relative module directory.
With --json it MUST emit schema1 string data equal to that directory and empty diagnostics; failures MUST emit null data and code/message diagnostics with empty stderr.

### APP-105
check-commits [range] MUST discover the nearest project root and call checks.History with the supplied range, or an empty range when omitted to use the configured local base.
Plain success MUST return exit0, report the number of checked commits, and list each squashed change with its full commit object ID and change path.

### APP-106
check-commits MUST return exit1 for history-rule violations and report every returned H01–H04 diagnostic with its affected path, full commit object ID and corrective action.

### APP-107
check-commits MUST return exit2 with corrective U01 for invalid arguments/options and the existing cause code for project configuration errors.
Operational history failures MUST return exit1 with their existing cause code, or P01 when uncoded, and a corrective action.

### APP-108
Successful check-commits --json MUST emit schema1 object data with commits in returned order and squashed entries containing commit and change, using empty arrays when appropriate and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-109
ci [range] MUST discover the nearest project root and run static project checks plus history checks on the supplied range, or the configured local-base range when omitted.
Static or history errors MUST prevent a successful CI result.

### APP-110
ci MUST run evals for every discovered module and require passing coverage of its current specs, regardless of verify_dependents.
The configured boundary checks MUST also pass for CI to succeed.

### APP-111
ci MUST run doctor canaries by default with agent probes disabled.
With --no-doctor it MUST skip the doctor canaries while retaining the other deterministic checks.

### APP-112
Successful ci MUST return exit0 with a confirmation.
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Deterministic check failures MUST return exit1 retaining their cause codes and corrective guidance.

### APP-113
Successful ci --json MUST emit schema1 confirmation-string data and collected diagnostics, including warnings.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-114
list MUST discover the nearest project root and list project.Queue entries in returned order.
Plain success MUST return exit0 and identify every queued change by slug, path and module, with empty output for an empty queue.

### APP-115
Each listed entry MUST report its phase, using queued when the phase is absent.
It MUST report the blocked reason and audit rounds, using an empty reason and zero rounds when absent.

### APP-116
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing cause code for configuration.
Queue failures MUST return exit1 with their existing cause code and corrective guidance.

### APP-117
Successful list --json MUST emit schema1 ordered-array data with slug, path, module, phase, blocked and audit_rounds fields, including an empty array for an empty queue, and empty diagnostics.
Failures MUST emit null data with code/message diagnostics.
It MUST leave stderr empty.

### APP-118
status MUST discover the nearest project root and expose workflow.Status for it.
Successful status MUST return exit0, including empty queues and warning-only reports.

### APP-119
Plain status MUST identify every queued entry in returned order by slug, path and module, with its phase, blocked reason, audit rounds, waiting change and blocking rule.
It MUST report an empty queue explicitly when there are no entries.

### APP-120
Plain status MUST identify discovered module names, prefixes, directories, dependencies and resolved zones.
It MUST report quarantines with module/ID/reason, boundary opt-outs with module/reason, each module's drift, the unapproved count and each archived change's phase-commit count from the returned report.

### APP-121
Plain status MUST show every returned module cost with its JUnit seconds, testcase counts per spec, eval lines and disposable lines.
Modules without a recorded cost MUST not receive a fabricated measurement.

### APP-122
Successful status --json MUST emit schema1 object data containing modules, queue, quarantines, boundary_opt_outs, drift_since_drill, unapproved, phase_commits and metrics.
Empty array/map containers MUST appear as arrays/maps rather than null.
It MUST leave stderr empty.

### APP-123
JSON module rows MUST use the existing name, id_prefix, dir, depends_on and zones shape.
Queue rows MUST use slug, path, module, phase, blocked, audit_rounds, waiting_on and waiting_rule.
Quarantine and opt-out rows MUST use module/id/reason and module/reason fields respectively.

### APP-124
JSON metrics MUST use module, seconds, cases_per_spec, eval_lines and disposable_lines fields.
The cases_per_spec, drift_since_drill and phase_commits maps MUST retain their returned identifier/count associations.

### APP-125
status MUST retain the returned diagnostic codes and messages, including budget and P08 warnings.
JSON diagnostics MUST retain path when present and warning status; plain output MUST show each diagnostic code and message.

### APP-126
Usage and project configuration errors MUST return exit2 with corrective guidance, using U01 for usage and the existing configuration cause code.
Status failures MUST return exit1 with corrective guidance, retaining assigned cause codes and using P01 for an unavailable base ref.
JSON failures MUST use null data with code/message diagnostics and empty stderr.

### APP-127
pr-body MUST discover the nearest project root and expose workflow.PRBody for the supplied explicit range.
Requested endpoint data MUST remain authoritative when later HEAD or working-tree metadata differs.

### APP-128
Plain pr-body MUST print the report Body and show every returned diagnostic code and message.
Successful reports MUST return exit0, including empty selections and warning-only reports.

### APP-129
Successful pr-body --json MUST emit schema1 with string data equal to the report Body and retained diagnostic codes, messages, warning status and path when present.
It MUST leave stderr empty.

### APP-130
Missing or extra ranges and invalid options MUST return exit2 with U01 and corrective guidance.
Project discovery or endpoint configuration-loading failures MUST return exit2 with the existing cause code or C01 and corrective guidance.
Unresolvable ranges MUST return exit1 with P01, the range resolution cause and corrective guidance.

### APP-131
JSON pr-body failures MUST emit null data with cause-code/message diagnostics and empty stderr.

### APP-132
show <module> [ID] --history MUST discover the nearest project root, call project.SpecHistory with the optional ID filter and return its entries in order with exit0.
Plain output MUST include each archive path, module, matching IDs, implementation date, Motivation, added and removed Decisions and complete Provenance.
Empty history MUST print an explicit no-history message with exit0.

### APP-133
Successful show --history --json MUST emit schema1 with ordered array data, empty diagnostics and empty stderr.
Each row MUST contain path, module, ids, motivation, decisions, removed_decisions, provenance and implementation_date, the latter an RFC3339 timestamp or an empty string when undated.
Empty history MUST use an empty array.

### APP-134
--history MUST be accepted only with show and MUST reject combination with --projected, using exit2 and a corrective U01 diagnostic for either misuse.

### APP-135
History MUST return exit2 for lookup failures with U02 naming the rejected module or ID, or for missing or malformed project configuration with C01.
Other project.SpecHistory failures MUST return exit1 retaining their cause code and corrective guidance.
Failure output MUST contain code/message diagnostics: JSON uses null data and empty stderr; plain output uses empty stdout and stderr diagnostics.

### APP-136
drill <module> MUST discover the nearest project root and invoke runner.Drill with that Root, the requested Module and the actual invoking executable as Bin. It MUST leave Timeout unset so configured backend bounds apply.

### APP-137
A successful drill MUST return exit 0 with empty stderr and plain output identifying the module, passing result, project-relative native report path and retained rebuilt-code ref.

### APP-138
Successful drill --json MUST return exit 0 with schema 1 object data containing module, result, report, ref and base_commit string fields associated with the runner result. It MUST emit empty diagnostics and empty stderr.

### APP-139
drill MUST reject missing or extra module arguments with exit 2 and corrective U01, and an unknown module with exit 2 and corrective U02 naming it. These rejections MUST invoke neither setup/eval hooks nor agents.

### APP-140
drill project-discovery or configuration-loading failures MUST return exit 2 with corrective diagnostics retaining their cause code or C01 when uncoded. JSON MUST use schema 1 with null data and empty stderr; plain failures MUST use empty stdout and stderr diagnostics.

### APP-141
Ordinary runner.Drill failures MUST return exit 1 with corrective diagnostics retaining the cause code or P01 when uncoded. Diagnostics MUST identify the native report and rebuilt ref when those artifacts exist. JSON MUST use schema 1 with null data and empty stderr; plain failures use empty stdout and stderr diagnostics.

### APP-142
A clean runner.IsUsageStop error with drill Result stopped MUST return exit 3 with the stop reason and any native report/ref identifiers. JSON MUST use schema 1 object data containing module, result, report, ref, base_commit and reason, with empty diagnostics and stderr; plain output MUST name the stopped module and reason.

### APP-143
show --stale <change> MUST discover the nearest project root and query project.StaleSpecs for the pending change selector.
For each recovered entry in returned order, it MUST identify the spec and display a line diff from Before to Current, distinguishing removed and added lines.

### APP-144
For an unrecovered stale entry, show --stale MUST identify the spec, display its RecordedHash and Current text, and indicate that stamped text was unavailable.
It MUST NOT present an unverified historical text as a recovered diff.

### APP-145
The CLI MUST accept show --stale with exactly one change selector and reject its combination with --history or --projected.
Use of --stale outside show or invalid stale-view arguments MUST return exit 2 with corrective U01 guidance describing valid show --stale usage.

### APP-146
show --stale MUST return exit 2 for change lookup failures, retaining U02 for an unknown selector or C13 for an ambiguous slug and naming the rejected selector.
Project discovery or configuration errors MUST return exit 2 retaining their cause code or C01 when uncoded; other query failures MUST return exit 1 retaining their cause code or P01.

### APP-147
check --fix MUST fill missing Modified and Removed bases across the pending queue using project.StampMissing before ordinary project.Check diagnostics.
Every existing bases entry, including valid and stale hashes, MUST remain unchanged; stale conflicts MUST still produce C08 and exit 1.

### APP-148
check --fix MUST use check's existing plain/JSON success and diagnostic shapes, returning exit 0 only when no error remains.
Stamping rejections MUST return exit 1 with their cause code and corrective guidance; project discovery or configuration errors MUST retain exit 2.

### APP-149
The CLI MUST accept check --fix without positional arguments and reject --fix with other commands.
Invalid fix-command usage MUST return exit 2 with corrective U01 guidance describing valid check --fix usage before mutation.

### APP-150
gc MUST discover the nearest project root, invoke workflow.GC and return exit 0 on success. Plain output MUST list the sorted full pruned ref names or explicitly report that no refs were pruned.

### APP-151
Successful gc --json MUST emit schema 1 object data containing a pruned array of sorted full deleted ref names, including an empty array when nothing was pruned, with empty diagnostics and stderr.

### APP-152
gc MUST return exit 2 for project discovery or configuration failures, retaining the cause code or C01 when uncoded. Other cleanup failures MUST return exit 1 retaining the cause code or P01 when uncoded. Failures MUST give corrective diagnostics, with null data and empty stderr in JSON mode.

### APP-153
migrate MUST discover the nearest project root and validate a format-1 project with project.Check before reporting already current with exit 0. The format-1 operation MUST leave working files, index, HEAD and refs unchanged.

### APP-154
migrate --json MUST emit schema 1 with string data equal to the plain confirmation on success or null data on failure. It MUST retain diagnostic codes and messages, with empty stderr.

### APP-155
migrate MUST return exit 2 for usage, project discovery or loading failures and unsupported format versions, using U01 for usage, C01 for unsupported formats and the existing loading cause code otherwise. Static validation errors MUST return exit 1 with their existing cause codes. All failures MUST include corrective guidance.

## Decisions


[0074-doctor-setup-commands] Doctor CLI leaves request Timeout unset, using the backend's existing defaults;
no new timeout flag is introduced. Successful doctor JSON checks use fields name,
module, passed and detail, with diagnostics in the outer envelope.

[0112-new-module-command] Creation uses the first configured module glob ending in a single wildcard
component under a literal parent, placing the requested name in that component;
a configuration with no such destination is reported as C01 before writing.
The package substitution replaces name hyphens with underscores; go_module is
read from go.mod when present and is otherwise empty.

[0127-show-history-command] History JSON uses lowercase keys and preserves project query section strings;
plain output labels each section and identifies undated entries explicitly.

- init-prompt-file-commands: Local recipes use recipe.yaml metadata, recipe.md instructions and a files/ project-relative payload; none uses empty language defaults and a generic module scaffold.


## Retained interfaces and files

packages/app/SPEC.md
packages/app/api.go

## Retained evals

packages/app/evals/.gitkeep
packages/app/evals/abandon_cli_test.go
packages/app/evals/app_test.go
packages/app/evals/authoring_commands_test.go
packages/app/evals/check_commits_cli_test.go
packages/app/evals/check_fix_cli_test.go
packages/app/evals/ci_cli_test.go
packages/app/evals/commits_cli_test.go
packages/app/evals/diagnostic_causes_test.go
packages/app/evals/doctor_setup_test.go
packages/app/evals/drill_cli_test.go
packages/app/evals/gc_cli_test.go
packages/app/evals/init_test.go
packages/app/evals/lifecycle_cli_test.go
packages/app/evals/list_cli_test.go
packages/app/evals/migrate_cli_test.go
packages/app/evals/new_module_cli_test.go
packages/app/evals/next_cli_test.go
packages/app/evals/pr_body_cli_test.go
packages/app/evals/project_commands_test.go
packages/app/evals/release_cli_test.go
packages/app/evals/run_cli_test.go
packages/app/evals/run_for_cli_test.go
packages/app/evals/saved_attempt_cli_test.go
packages/app/evals/show_history_cli_test.go
packages/app/evals/show_stale_cli_test.go
packages/app/evals/status_cli_test.go
packages/app/evals/usage_cli_test.go

Rebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.

Run the target's eval hook from the workspace root as often as needed.

~~~sh
mkdir -p .fireweed/runs/drill && sh -c 'go run .fireweed/hooks.go eval '"'"'packages/app'"'"' '"'"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/.fireweed/runs/drill/report.xml'"'"''
~~~

Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module.

codex
I’ll read the rebuild guardrails and retained app interfaces, then implement the CLI in the disposable paths and run the app eval hook.
exec
/bin/zsh -lc 'pwd; git status --short; ls -la; cat SYSTEM.md; cat docs/rebuild/PROGRESS.md; cat docs/rebuild/DECISIONS.md; ls -la packages/app; ls -la .fireweed/bin' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 632ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
total 280
drwx------@   17 joel  staff     544 Nov 10 12:15 .
drwx------@ 6453 joel  staff  206496 Nov 10 12:15 ..
drwxr-xr-x@    8 joel  staff     256 Nov 10 12:15 .fireweed
drwxr-xr-x@    9 joel  staff     288 Nov 10 12:15 .git
-rw-r--r--@    1 joel  staff      56 Nov 10 12:15 .gitignore
-rw-r--r--@    1 joel  staff     859 Nov 10 12:15 AGENTS.md
-rw-r--r--@    1 joel  staff   35149 Nov 10 12:15 LICENSE
-rw-r--r--@    1 joel  staff     100 Nov 10 12:15 Makefile
-rw-r--r--@    1 joel  staff    3399 Nov 10 12:15 README.md
-rw-r--r--@    1 joel  staff    8073 Nov 10 12:15 SYSTEM.md
drwxr-xr-x@    3 joel  staff      96 Nov 10 12:15 docs
-rw-r--r--@    1 joel  staff   61761 Nov 10 12:15 fireweed-design-v1.md
-rw-r--r--@    1 joel  staff    1289 Nov 10 12:15 fireweed.yaml
-rw-r--r--@    1 joel  staff      79 Nov 10 12:15 go.mod
-rw-r--r--@    1 joel  staff     151 Nov 10 12:15 go.sum
-rw-r--r--@    1 joel  staff     146 Nov 10 12:15 main.go
drwxr-xr-x@   10 joel  staff     320 Nov 10 12:15 packages
# Fireweed

Fireweed is a local, language-agnostic Go CLI for checked spec-driven development and regeneration. The unchanged fireweed-design-v1.md is product authority; docs/rebuild/GOAL.md is the restart prompt and ACCEPTANCE.md defines ordered milestones and gates, adding no product scope. Record relevant discrepancies briefly in docs/rebuild/DECISIONS.md; add evals only for uncovered design behavior.

## Modules

contracts owns public data, Markdown grammar, IDs, hashes and zones. gitstore owns safe local Git and process primitives. assets owns embedded recipes, playbooks and versions. project owns discovery, projection, validation and authoring. checks owns JUnit coverage, hooks, boundaries, history and doctor. workflow owns claims, lifecycle, provenance and context. runner owns agents, queues, retries, recovery and drills. app composes the CLI. SPEC.md frontmatter is the only graph source.

## Conventions

Module path: github.com/lurkshark/fireweed. Thin api.go facades and public contracts are retained interfaces; all behavior belongs under internal/. Evals import public packages only and use temporary repositories, explicit fixture Git identities, deterministic local hooks and bounded fake agents. No runtime network calls or private implementation tests. Keep dependencies minimal (YAML parser only initially). Never import a public parent from its own internal implementation; shared data belongs in contracts.

Fixture notes: canonicalize macOS temporary paths before comparing symlinked locations; parse Git stdout separately from stderr and use `git --no-optional-locks status` for read-only status. Avoid case-only ref names on case-insensitive filesystems. Keep shared facade types in retained model packages to avoid import cycles. Accept the supplied macOS sh diagnosis: fixtures receiving long prompt arguments use Bash or stdin; do not reinterpret it as a timeout or investigate it again.

## Controller and product

The controller is `.fireweed/bin/fireweed`, a separate build of this repository, so a worker's build of the product under test (`bin/fireweed`) never replaces the controller it runs under. PROGRESS records the commit it was built from. Rebuild it from clean committed main with `go build -trimpath -o .fireweed/bin/fireweed .` only when no run or claim is active, then record the commit in PROGRESS. Put `.fireweed/bin` first on worker PATH. FIREWEED_BIN names the executable invoking each hook, and scratch setup copies it into the scratch tree.

## Agent configuration

Workers and regeneration use exactly `codex exec --ephemeral --ignore-user-config -m gpt-6.1-sol -c 'model_reasoning_effort="high"' -s workspace-write {prompt}`. Separate restricted auditors use exactly `codex exec --ephemeral --ignore-user-config -m gpt-6-astra -c 'model_reasoning_effort="medium"' -s workspace-write {prompt}`. Do not silently change the model or effort or retry indefinitely after an access failure.

The repository owner grants standing authorization for the configured workers and separate auditors to send the context needed for owner-approved Fireweed changes and drills to OpenAI Codex at `https://chatgpt.com/backend-api/codex`. Worker context may include relevant tracked source, interfaces, Specs, evals, change files, operational instructions, playbooks and verification feedback. Auditor context stays restricted to the change, projected spec, target evals and audit playbook. For an approved drill, the regeneration worker may use the Fireweed-built reconstruction workspace, including the target interfaces, Specs and evals, dependency implementations, and required tracked setup and hooks. Keep credentials and unrelated local data out of the submitted context. This authorization covers Codex access across sessions. Obtain owner approval for the underlying change or drill and observe its review stops.

## Author and worker policy

Authors may commit change files and protected setup or docs on main, between active changes and in separate commits. Only fireweed applies Specs and Decisions, makes phase commits and archives. Workers follow context and zones, never commit, never weaken enforcement, never invoke lifecycle mutators (start, done, release, next, run, abandon) and never edit run, phase or Provenance. Audit stays enabled with separate restricted context. No push, publication or release tags without explicit owner instruction.

## Scope guardrails

- **Scope.** Implement only an owner request, a demonstrated failure or a concrete unmet design clause. Agent suggestions and drill gaps authorize nothing until the owner agrees. Stop and report on a blocked change, an access failure or a decision that needs the owner, and stop when the selected work is done.
- **Reconstruction inputs.** Never use `rebuild-v1` or deleted code as implementation oracles; old spec IDs have no authority. Drills use exactly the input workspace the product builds (RUN-053); nothing else is added or policed.
- **Conventions.** The product never contacts the network or runs fetch/push. Read commands run no hooks/agents and change no files, index or refs; output is deterministic. Check these by review and one app-level PATH-shim eval recording Git/hook/agent invocations, not module specs.
- **Evals.** Use plain Go public-interface tests, temporary repositories, fixture hooks and fake agents. Inject faults through public inputs. No process instrumentation: cgo, syscall/I/O interposers, loader tricks, sandbox profiles, re-signed tools, argv-classifying Git proxies, fsmonitor traps or byte-coverage observers.
- **Specs.** One observable behavior per item, at most three normative clauses and about 80 words. No aliasing, backing-storage, purity or syscall requirements, and no tuning constants or internal mechanics unless a caller can observe them.
- **Environment and processes.** Support macOS/Linux on amd64/arm64 with ordinary full SHA-1 clones. Windows, other architectures, SHA-256/partial/promisor clones, replacement refs, fsmonitor and filters get no special code/evals. Hooks/agents have timeouts and process groups; cancellation kills the group. Escaped/daemonized/reparented processes are a documented limit. No native objects, assembly, cgo or embedded/generated architecture binaries.
- **Cost and corrections.** Own Verify stays under ten minutes, full implementation Verify under 35; aim for at most 15 cases per spec. Do not raise bounds or add instrumentation. After two unsuccessful corrections of one issue, try one simpler approach; if that fails, report and stop.
- **Audit.** Findings identify spec clauses with no asserting test, not helpers, platforms, unnamed edge cases or conventions. Re-audit only IDs with previous findings. At most two findings round trips; anything further becomes Questions and blocks. No replacement change to reset rounds.
- **Authoring.** Read the author-change playbook and projected specs; run context before each step. Motivation names a design/SYSTEM clause or user need. Prefer narrowing with legitimate already_satisfied over extra hardening. Never edit Specs or Decisions by hand. Frozen-eval corrections use audited eval_fix.
- **Enforcement.** P07 and audit bounds stay active. Never alter refs, bounds or requirements to get past a check.
- **Decisions.** Record only missing behavioral choices in one or two sentences; keep module Decisions under about 300 words. Process, fixture and recovery notes belong in change Notes, not Decisions.
- **Drills.** A drill is a measurement, not a gate to force through. Run `fireweed drill <module>`, commit the report it writes, and turn gaps the owner agrees with into ordinary changes. Do not respond to a failing drill with drill-specific features, wrappers, supervisors or prompt tuning, and do not repeat a drill until a change addresses its gaps.
- **Evidence.** Evidence is what fireweed commits: archived changes with Provenance and drill reports in `.fireweed/drills/<module>.md`. Do not write separate handoff, receipt, hash or review files. Keep PROGRESS to the current state and next step.
# Rebuild progress

## Current state (2026-11-10)

WF1 is complete. Fireweed applied and archived both changes after separate
audits with no findings and passing implementation verification:

- [0195-run-argv-contract](../../changes/archive/0195-run-argv-contract.md)
  (`eb4a7f2`) documents `gitstore.Request.Args[0]` as the executable and
  adds regression coverage. All 925 cases passed.
- [0196-drill-self-checked-reconstruction](../../changes/archive/0196-drill-self-checked-reconstruction.md)
  (`084fbd4`) gives the drill worker a buildable input workspace with
  dependency implementations and its own eval hook. All 485 cases passed.

Fireweed released both claims after applying the changes.

WF2 is complete. `go test ./...`, `go vet ./...` and `fireweed ci` passed
on the WF2 controller source commit `6cc2629`. The first Go test run alongside CI
failed a timed app eval, which passed on the standalone retry. CI ran with
Doctor enabled in a temporary local clone over `65bf933..6cc2629`.

WF3 is complete. With owner approval, the recorded controller ran
`fireweed drill workflow` once at `4754026` with the configured Codex worker.
The drill passed after its configured retry. All 666 cases passed (181
workflow, 323 app, 162 runner) and the worker reported no contradicted
Decisions. Fireweed cleaned up the scratch worktree and retained reconstruction
`38aa8fc` under `refs/fireweed/drills/workflow/20261109T224519.080152000`.
The [native report](../../.fireweed/drills/workflow.md) is committed in `c62b5be`.

WF4 is complete. The owner reviewed the report and approved these dispositions:

- **FLOW-110 / APP-126.** Closed. APP-126 specifies P01 for an unavailable
  base and its eval passed on the drill retry.
- **FLOW-052 / checks.** Deferred raw boundary-hook output retention. The
  current boundary API returns diagnostics only.
- **FLOW-027 / checks.** Fireweed applied and archived
  [0197-boundary-hook-nonzero-exit](../../changes/archive/0197-boundary-hook-nonzero-exit.md)
  in `d45652a`. CHK-043 now rejects a nonzero exit without a module-attributed
  violation. Both regressions failed first, the separate audit found no
  gaps, and all 787 implementation cases passed. `fireweed check` and
  `fireweed check-commits a0e25a2..d45652a` passed. The queue is empty and
  Fireweed released the claim.

The owner reaffirmed standing Codex authorization for the configured workers
and restricted auditors on owner-approved changes and drills. Commit `a4e020a`
records the scope, including reconstruction inputs, and destination in
[SYSTEM.md](../../SYSTEM.md#agent-configuration).

## Controller and candidate

The installed `.fireweed/bin/fireweed` was rebuilt with `-trimpath` from clean
main with an empty queue and no claims. Go build metadata confirms the
revision and a clean source tree.

Controller source: `e1e1869131833778f1f8f1a23cf62b0d3b3869b3`

With `GOCACHE=/tmp/fireweed-rebuild-cache`, the fresh macOS arm64 candidate
clone passed the product build, `go test ./...`,
`go test -race ./...`, `go vet ./...`, `fireweed check`, and standalone
`fireweed doctor` including canaries, recovery and both agent probes.
`fireweed check-commits` passed 1,304 commits over
`be35552b706aa3000ff8f11306d129ff1d701b85..e1e1869131833778f1f8f1a23cf62b0d3b3869b3`.
CI with Doctor enabled over the same range passed on the full retry. CI
verified passing evals and coverage for app, assets, checks, contracts,
gitstore, project, runner and workflow. The first CI run failed on the
deadline-sensitive `Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied`
eval, which passed in isolation before the full retry.

## Drill proof inventory

**Contracts proof accepted.** The verified controller ran
`fireweed drill contracts` at `e7f1554`. All 1,119 cases passed on the first
attempt, including 64 contracts cases and 1,055 transitive dependent cases.
The [native report](../../.fireweed/drills/contracts.md) is committed in
`24f9d94`. Fireweed cleaned up both scratch locations and retained
`30c7cdfd356d7068f841179c9c4a1fb1321be6c4` at
`refs/fireweed/drills/contracts/20261110T034341.899515000`.
The proof changes only disposable contracts paths.

The report lists no contradicted Decisions. The owner deferred all three gaps:

- `CON-002` / `CON-009` leave fenced headings and level-one section boundaries
  implicit. The reconstruction handles these differently from the current code.
- `CON-008` leaves equality for YAML NaN and offset timestamp keys ambiguous.
- `CON-017` through `CON-019` specify codes and, for duplicates, a subject.
  The evals also pin error wording or prefixes.

**Renewed gitstore proof accepted.** The verified controller
ran `fireweed drill gitstore` at `c3863e5`. All 933 cases passed on the first
attempt, including 13 gitstore cases and 920 transitive dependent cases. Both
GIT-010 expected-absence forms passed. The
[native report](../../.fireweed/drills/gitstore.md) is committed in `c687766`.
Fireweed cleaned up both scratch locations and retained
`d3230d81b89f327ec5a4dfabfcd6381951d7b235` at
`refs/fireweed/drills/gitstore/20261110T060020.420365000`. The proof changes
only disposable gitstore paths.

The owner-approved GIT-010 changes
[0198](../../changes/archive/0198-casref-empty-expected-delete.md) (`feabbc7`)
and [0199](../../changes/archive/0199-casref-zero-expected-delete.md) (`3fa1918`)
passed fail-first, separate audits, and full verification. The earlier report
at `0b34909` and proof `3382ac8` remain in history and at the earlier drill ref.
That reconstruction failed the explicit-zero eval added after its drill.

The renewed report lists no contradicted Decisions. The owner deferred its
remaining gaps:

- `GIT-001` and `GIT-003` leave unborn HEAD and precedence for overlapping
  tracked edits implicit.
- `GIT-005` leaves environment inheritance and executable PATH lookup implicit.
- `GIT-007` leaves literal path syntax and index state after a failed commit
  implicit. Lifecycle callers use the `:(literal)` spelling.
- `GIT-008` through `GIT-010` leave abbreviated blob IDs, ref naming, and
  symbolic-ref lookup and mutation implicit.

**Renewed assets proof accepted.** The verified controller ran
`fireweed drill assets` at `34f9f54`. All 1,049 cases passed on the configured
retry, including 125 assets cases and 924 transitive dependent cases. All four
Doctor cases that failed in the earlier drill passed. The
[native report](../../.fireweed/drills/assets.md) is committed in `db34f44`.
Fireweed cleaned up both scratch locations and retained
`fe1a898a1944436fa7fc066c4cf3781f26ef43cb` at
`refs/fireweed/drills/assets/20261110T080441.699697000`. Its changes are
confined to disposable assets paths.

The earlier failed report remains in `84f7138` and reconstruction `69b7d7c`
remains at `refs/fireweed/drills/assets/20261110T063418.808553000`. It passed
125 assets cases and failed four dependent Doctor cases on `TestMissingIDCanary`.
The owner deferred its other gaps on baseline coverage, mixed Go package zones,
expected-failure JUnit status, and static import resolution.

Fireweed applied [0200](../../changes/archive/0200-doctor-go-missing-id-canary.md)
in `f854288`. CHK-059 now matches the supplied Go test declaration and its
subtests, including names without an underscore after `Test`. The change
passed fail-first, a separate audit with no findings, and 791 implementation
cases. All four previously failing Doctor evals passed with the retained
assets reconstruction `69b7d7c` and the current checks code in a temporary clone.

The owner accepted the proof and approved proceeding to project. The renewed
report lists no contradicted Decisions. The retry restored the
`# Write evals` heading expected by workflow and app, which the current product
already supplies. The report identifies the heading as implicit in AST-002
and the missing-ID handoff to Doctor as implicit in AST-058. The remaining
observations revisit the deferred mixed-zone, JUnit and import-resolution gaps.

**Renewed project proof accepted.** The verified controller ran
`fireweed drill project` at `73e269d`. All 925 cases passed on the configured
retry, including 134 project and 791 dependent cases. `PRJ-069` and the
original `APP-116`/`APP-117` regression both passed. The
[native report](../../.fireweed/drills/project.md) is committed in `acc1c26`.
Fireweed cleaned up both scratch locations and retained
`0b51fc13ad29206da8d3a5ca762b2a0b06688036` at
`refs/fireweed/drills/project/20261110T164427.736325000`. Its changes are
confined to disposable project paths.

The earlier failed report remains in `dfbb207` and reconstruction `917fb02d`
remains at `refs/fireweed/drills/project/20261110T144729.916008000`. That
reconstruction passed 133 project and 790 dependent cases but returned C13
instead of C03 in the app list eval for a regular-file `changes` path.

Fireweed applied [0201](../../changes/archive/0201-queue-path-diagnostic.md)
in `73e269d`. The owner-approved amendment in `fde499d` narrowed `PRJ-069` to
the existing public behavior. `Queue` rejects the regular-file path and
identifies `changes` in its error. `Check` emits a non-warning C03 diagnostic
identifying `changes`. The amendment retained the same run and audit phase.
The separate audit found no gaps and all 925 implementation cases passed.
Fireweed released the claim. The earlier P09 audit submission remains at
`refs/fireweed/attempts/queue-path-diagnostic/1` (`dc223bd`).

The renewed report lists no contradicted Decisions. The first attempt passed
all project cases and the retry addressed downstream diagnostic mismatches.
The owner deferred all remaining gaps. They cover app diagnostic codes and
wording, duplicate C08 diagnostics, partial projections, numbering past 9999,
intermediate graph validation, stale-view and history behavior, and the
design's warning for specs outside configured globs.

**Renewed checks proof accepted.** The verified controller ran
`fireweed drill checks` at `9e81519`. All 791 cases passed on the configured
retry, including 125 checks and 666 dependent cases. All three CHK-043 cases
and the app CI boundary-failure eval passed. The
[native report](../../.fireweed/drills/checks.md) is committed in `8369532`.
Fireweed cleaned up both scratch locations and retained
`3f0279d42e2bcca4efd7641d6a3a5ebd8d40ace2` at
`refs/fireweed/drills/checks/20261110T174227.798503000`. Its changes are
confined to disposable checks paths.

The retry fixed a JUnit parser ambiguity that dropped failure details from
numbered test labels. The `RUN-044` Unicode excerpt eval passed after that
correction. Two Doctor fixtures could not build the executable in the
worker's reconstruction workspace because of a missing module checksum.
Both fixtures passed in the controller's full scratch checkout.

The report lists no contradicted Decisions. The owner deferred all remaining
gaps. They cover ID matching for numbered testcase labels, shell quoting around
hook placeholders, boundary rule attribution, canary identity, probe output
ordering, and fixture availability in the reconstruction workspace.

**Reuse workflow.** The retained ref and commit `38aa8fc` match the committed
passing report and change only workflow's disposable files. Since that report,
`0197` changed the checks dependency. The retained workflow reconstruction
passed the `FLOW-027` and `FLOW-062` evals and the app CI boundary-failure eval
against the current candidate in a temporary clone.

**Runner proof accepted.** The verified controller ran
`fireweed drill runner` at `5b59f35`. All 485 cases passed on the configured
retry, including 162 runner and 323 app cases. The
[native report](../../.fireweed/drills/runner.md) is committed in `9390435`.
Fireweed cleaned up both scratch locations and retained
`521e048e7c39879c2e6600ab1158171395847ff9` at
`refs/fireweed/drills/runner/20261110T192409.383763000`. Its changes are
confined to disposable runner paths.

The first attempt passed all runner cases. Three app cases failed because
the approval refusal did not use P01. The retry preserved the typed approval
error, emitted P01, and passed `APP-075`, `APP-077`, and `APP-082`.

The report lists no contradicted Decisions. The owner deferred all three
gaps. They cover shell quoting for already-quoted placeholders
(`RUN-003`), repeated verification on successful phases (`RUN-014`), and
the P01 diagnostic for approval refusals (`RUN-041`).

**Two historical passing refs remain missing.** The prior passing reports
remain in Git history.
The following refs under `refs/fireweed/drills/<module>/` are absent from this
checkout and GitHub origin advertises none. Their exact commits were unavailable
across 119 local object stores, including the historical candidate directories.

| Module | Ref suffix | Expected commit |
| --- | --- | --- |
| project | `20261102T031647.104747000` | `481bc00c1da360d3e8bf76cd5ba04053d2dccaeb` |
| checks | `20261107T002339.020159000` | `c0ca0397e3781410e795f27b52de2b2105350559` |

The renewed project and checks proofs supersede these historical reports.
`app` needs its first native report and retained ref.

## Next step

Run the owner-approved `app` drill with the verified controller, commit the
native report, retain the proof ref, and stop for owner review. This is the
remaining drill in the [regeneration goal](GOAL.md).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
# Rebuild decisions

Retained design clarifications. These do not add milestone scope or replace the unchanged design. Standing
process and fixture guidance is consolidated in SYSTEM.md. Historical cutover,
controller comparisons and handoff notes remain in commit `c34c3cb`, never
reconstruction inputs.

## D005 — Initial product diagnostic allocation

The design requires stable codes and explanations but does not assign a meaning
to P2-01. AST-003 assigns it to missing genuine fail-first evidence.

## D006 — Diagnostic allocation (human review 3)

Codes identify a stable cause. Their prefix does not determine severity; the
diagnostic carries its warning status. Assignments below apply to the
product. This allocation did not itself change the controller pin.

| Range | Area and fixed assignments |
| --- | --- |
| C01–C99 | Deterministic configuration and static checks. C01 configuration discovery/decoding; C02 module graph; C03 document/item syntax; C04 normative text; C05 ID ownership/uniqueness; C06 spec tags; C07 change structure/Motivation; C08 bases; C09 handles; C10 coverage; C11 eval-hook/JUnit contract; C12 metadata/zone declarations; C13 queue identity. C14–C16 doctor diagnostics are allocated in D007. C17 spec shape, C18 copied conventions, C19 Decision length retain design meanings. |
| P01–P99, P2-01 | Lifecycle and approval. P01 preconditions; P02 claims; P03 approved state; P04 approval warning; P05 stale base; P06 blocked/Questions; P07 unapproved-work limit; P08 repeated-module warning; P09 audit submission. Existing P2-01 remains missing genuine fail-first evidence. |
| Z01–Z99 | Phase-zone permissions. Existing Z01 remains a forbidden path edit. |
| X01–X99 | Boundaries. X01 unsuppressed violation; X02 required hook missing or unable to complete; X03 reasoned opt-out; X04 boundary enforcement unnecessary. B05 is retired. |
| U01–U99 | CLI usage and lookup. U01 invalid arguments/options; U02 unknown lookup name. The legacy CLI_* labels are retired. |
| E01–E99 | Required eval results. E01 an enforced testcase failed where passing results are required. |
| H01–H04 | History, with the design's exact meanings. |
| B01–B04 | Budget/report warnings only, with the design's exact meanings. No B05 allocation. |

Unlisted numbers are reserved. Every emitted code must have an accurate cause
and corrective action in explain. Wrong label assertions are repaired through
eval_fix without losing behavior assertions or ID coverage; implementation/code
allocation changes use normal spec changes. Existing product Specs and Decisions
are never edited by hand.

## D007 — Doctor diagnostics

Doctor uses the previously reserved C14–C16 slots:
C14 for doctor hook/canary setup, detection and recovery failures; C15 for agent
probe failures or usage limits, and the shared-auditor-command warning; C16 for
stale or missing managed guidance and playbook override base metadata. Severity
remains a property of the diagnostic. D006's existing assignments remain unchanged.

## D008 — Regenerate warning

The regenerate lifecycle feature assigns P10 to the warning that its module
lacks a passing drill dated after the latest applied implementation Provenance.
This allocation grants no drill authorization; PROGRESS records current
status and GOAL links the planned milestone.
total 80
drwxr-xr-x@  5 joel  staff    160 Nov 10 12:15 .
drwxr-xr-x@ 10 joel  staff    320 Nov 10 12:15 ..
-rw-r--r--@  1 joel  staff  33788 Nov 10 12:15 SPEC.md
-rw-r--r--@  1 joel  staff    253 Nov 10 12:15 api.go
drwxr-xr-x@ 30 joel  staff    960 Nov 10 12:15 evals
total 13512
drwxr-xr-x@ 3 joel  staff       96 Nov 10 12:15 .
drwxr-xr-x@ 8 joel  staff      256 Nov 10 12:15 ..
-rwxr-xr-x@ 1 joel  staff  6914850 Nov 10 12:15 fireweed

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
fireweed playbook author-change
git ls-files packages/app; cat main.go packages/app/api.go; for x in packages/*/api.go; do echo ====; cat "$x"; done
rg '"'"'^### APP-|''^'"\\[' packages/app/SPEC.md | tail -35" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 exited 1 in 1618ms:
# Author change
Read the design, SYSTEM.md and the module spec before writing a change. Give its Motivation a user-visible need or a clause of the design or SYSTEM.md; a probe showing a theoretical failure is insufficient. State each observable behavior at the module interface and keep the change focused. If a requirement costs more than it is worth, narrow or remove the spec and use already_satisfied for behavior the current implementation already meets. Do not commit during a worker step.
packages/app/SPEC.md
packages/app/api.go
packages/app/evals/.gitkeep
packages/app/evals/abandon_cli_test.go
packages/app/evals/app_test.go
packages/app/evals/authoring_commands_test.go
packages/app/evals/check_commits_cli_test.go
packages/app/evals/check_fix_cli_test.go
packages/app/evals/ci_cli_test.go
packages/app/evals/commits_cli_test.go
packages/app/evals/diagnostic_causes_test.go
packages/app/evals/doctor_setup_test.go
packages/app/evals/drill_cli_test.go
packages/app/evals/gc_cli_test.go
packages/app/evals/init_test.go
packages/app/evals/lifecycle_cli_test.go
packages/app/evals/list_cli_test.go
packages/app/evals/migrate_cli_test.go
packages/app/evals/new_module_cli_test.go
packages/app/evals/next_cli_test.go
packages/app/evals/pr_body_cli_test.go
packages/app/evals/project_commands_test.go
packages/app/evals/release_cli_test.go
packages/app/evals/run_cli_test.go
packages/app/evals/run_for_cli_test.go
packages/app/evals/saved_attempt_cli_test.go
packages/app/evals/show_history_cli_test.go
packages/app/evals/show_stale_cli_test.go
packages/app/evals/status_cli_test.go
packages/app/evals/usage_cli_test.go
package main

import (
	"github.com/lurkshark/fireweed/packages/app"
	"os"
)

func main() { os.Exit(app.Run(os.Args[1:], os.Stdout, os.Stderr)) }
package app

import (
	"github.com/lurkshark/fireweed/packages/app/internal/entry"
	"io"
)

// Run executes the fireweed CLI and returns its process exit status.
func Run(args []string, out, errOut io.Writer) int { return entry.Run(args, out, errOut) }
zsh:3: === not found

exec
/bin/zsh -lc "for x in packages/*/api.go; do printf '\\n%s\\n' \""'$x"; cat "$x"; done; rg --files packages/{project,workflow,checks,runner,contracts,assets} | rg '"'/(model|types|api)'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc "rg -n 'func Test|func (run|require|assert|cli|build)' packages/app/evals; cat fireweed.yaml; ls .fireweed; export PATH=\""'$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed context app' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 670ms:

packages/app/api.go
package app

import (
	"github.com/lurkshark/fireweed/packages/app/internal/entry"
	"io"
)

// Run executes the fireweed CLI and returns its process exit status.
func Run(args []string, out, errOut io.Writer) int { return entry.Run(args, out, errOut) }

packages/assets/api.go
package assets

import "github.com/lurkshark/fireweed/packages/assets/internal/content"
import "github.com/lurkshark/fireweed/packages/assets/internal/recipe"
import "github.com/lurkshark/fireweed/packages/assets/model"

type Recipe = model.Recipe

func Version() string { return content.Version() }

func Help() string { return content.Help() }

func Playbook(name string) (string, error) { return content.Playbook(name) }

func Explain(code string) (string, error) { return content.Explain(code) }

func BuiltinRecipe(name string) (Recipe, error) { return recipe.BuiltinRecipe(name) }

packages/checks/api.go
package checks

import (
	"context"

	"github.com/lurkshark/fireweed/packages/checks/internal/engine"
	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case = model.Case
type Report = model.Report
type HookRequest = model.HookRequest
type BoundaryRequest = model.BoundaryRequest
type ZoneRequest = model.ZoneRequest
type DoctorRequest = model.DoctorRequest
type DoctorCheck = model.DoctorCheck
type DoctorReport = model.DoctorReport
type HistoryReport = model.HistoryReport
type SquashedChange = model.SquashedChange

func ParseJUnit(data []byte, specs []contracts.Item) (Report, error) {
	return engine.ParseJUnit(data, specs)
}
func RequireCoverage(report Report, specs []contracts.Item) error {
	return engine.RequireCoverage(report, specs)
}
func RequireFailFirst(report Report, ids, already []string) error {
	return engine.RequireFailFirst(report, ids, already)
}
func RequirePass(report Report, specs []contracts.Item) error {
	return engine.RequirePass(report, specs)
}
func Evaluate(ctx context.Context, request HookRequest) (Report, error) {
	return engine.Evaluate(ctx, request)
}
func Boundaries(ctx context.Context, request BoundaryRequest) ([]project.Diagnostic, error) {
	return engine.Boundaries(ctx, request)
}
func Zones(request ZoneRequest) []project.Diagnostic { return engine.Zones(request) }
func Doctor(ctx context.Context, request DoctorRequest) DoctorReport {
	return engine.Doctor(ctx, request)
}
func History(root, revisionRange string) (HistoryReport, error) {
	return engine.History(root, revisionRange)
}

packages/contracts/api.go
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

packages/gitstore/api.go
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
func WriteBlob(root string, data []byte) (string, error) { return local.WriteBlob(root, data) }
func ReadBlob(root, oid string) ([]byte, error)          { return local.ReadBlob(root, oid) }
func Ref(root, name string) (string, bool, error)        { return local.Ref(root, name) }
func CASRef(root, name, old, new string) error           { return local.CASRef(root, name, old, new) }

// Run launches request.Args as the complete argv without prepending git.
func Run(ctx context.Context, request Request) (Result, error) { return local.Run(ctx, request) }

packages/project/api.go
package project

import (
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/internal/history"
	"github.com/lurkshark/fireweed/packages/project/internal/queue"
	"github.com/lurkshark/fireweed/packages/project/model"
)

type Snapshot = model.Snapshot
type Module = model.Module
type Change = model.Change
type SpecHistoryEntry = model.SpecHistoryEntry
type StaleSpec = model.StaleSpec
type Diagnostic = model.Diagnostic

func Load(root string) (Snapshot, error) { return fixture.Load(root) }

func LoadAt(root, revision string) (Snapshot, error) { return fixture.LoadAt(root, revision) }

func ValidateGraph(modules []Module) error { return fixture.ValidateGraph(modules) }

func NewChange(root, module, slug string) (string, error) {
	return fixture.NewChange(root, module, slug)
}

func Queue(root string) ([]Change, error) { return queue.Queue(root) }

func Project(root, through string) (Snapshot, error) { return queue.Project(root, through) }

func Stamp(root, change string) error { return queue.Stamp(root, change) }

func StampMissing(root, change string) error { return queue.StampMissing(root, change) }

func StaleSpecs(root, change string) ([]StaleSpec, error) { return queue.StaleSpecs(root, change) }

func Check(root string) []Diagnostic { return queue.Check(root) }

func SpecHistory(root, module, id string) ([]SpecHistoryEntry, error) {
	return history.SpecHistory(root, module, id)
}

packages/runner/api.go
package runner

import (
	"context"

	"github.com/lurkshark/fireweed/packages/runner/internal/drill"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/internal/run"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

type AgentRequest = model.AgentRequest
type AgentResult = model.AgentResult
type NextRequest = model.NextRequest
type NextResult = model.NextResult
type StepResult = model.StepResult
type RunRequest = model.RunRequest
type RunResult = model.RunResult
type DrillRequest = model.DrillRequest
type DrillResult = model.DrillResult

func InvokeAgent(ctx context.Context, request AgentRequest) (AgentResult, error) {
	return invoke.Agent(ctx, request)
}

func Next(ctx context.Context, request NextRequest) (NextResult, error) {
	return next.Run(ctx, request)
}

func Run(ctx context.Context, request RunRequest) (RunResult, error) {
	return run.Execute(ctx, request)
}

func Drill(ctx context.Context, request DrillRequest) (DrillResult, error) {
	return drill.Run(ctx, request)
}

func IsUsageStop(err error) bool {
	return next.IsUsageStop(err)
}

func IsApprovalRequired(err error) bool {
	return next.IsApprovalRequired(err)
}

func IsTimeStop(err error) bool {
	return run.IsTimeStop(err)
}

packages/workflow/api.go
package workflow

import (
	"context"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type State = model.State
type Request = model.Request
type Result = model.Result
type ModuleCost = model.ModuleCost
type CostSnapshot = model.CostSnapshot
type StatusReport = model.StatusReport
type PRBodyReport = model.PRBodyReport
type StatusEntry = model.StatusEntry
type Quarantine = model.Quarantine
type BoundaryOptOut = model.BoundaryOptOut
type BlockRequest = model.BlockRequest
type BlockResult = model.BlockResult
type GCRequest = model.GCRequest
type GCResult = model.GCResult

func Start(root, change string) (State, error) { return lifecycle.Start(root, change) }
func StartFromAttempt(root, change string, attempt int) (State, error) {
	return lifecycle.StartFromAttempt(root, change, attempt)
}
func Eligible(root string) ([]project.Change, error) { return lifecycle.Eligible(root) }
func Current(root, change string) (State, error)     { return lifecycle.Current(root, change) }
func Commits(root, change string) ([]string, error)  { return lifecycle.Commits(root, change) }
func Abandon(root, change string) (string, error)    { return lifecycle.Abandon(root, change) }
func Release(root, change string) error              { return lifecycle.Release(root, change) }
func Verify(ctx context.Context, request Request) (Result, error) {
	return lifecycle.Verify(ctx, request)
}
func Done(ctx context.Context, request Request) (Result, error) {
	return lifecycle.Done(ctx, request)
}
func Block(request BlockRequest) (BlockResult, error) {
	return lifecycle.Block(request)
}
func Context(root, change string, full bool) (string, error) {
	return lifecycle.Context(root, change, full)
}
func Costs(root string) ([]CostSnapshot, error) { return lifecycle.Costs(root) }
func Status(root string) (StatusReport, error)  { return lifecycle.Status(root) }
func PRBody(root, revisionRange string) (PRBodyReport, error) {
	return lifecycle.PRBody(root, revisionRange)
}
func IsEndpointConfigError(err error) bool { return lifecycle.IsEndpointConfigError(err) }
func GC(ctx context.Context, request GCRequest) (GCResult, error) {
	return lifecycle.GC(ctx, request)
}
packages/checks/api.go
packages/workflow/api.go
packages/project/api.go
packages/checks/model/model.go
packages/project/model/model.go
packages/workflow/model/model.go
packages/assets/api.go
packages/runner/api.go
packages/contracts/api.go
packages/assets/model/model.go
packages/contracts/model/model.go
packages/runner/model/model.go
packages/assets/internal/recipe/files/typescript-pnpm-index.ts.tmpl
packages/assets/internal/recipe/files/typescript-pnpm-hooks.mjs
packages/assets/internal/recipe/files/typescript-pnpm-package.json.tmpl
packages/assets/internal/recipe/files/typescript-pnpm-canaries.yaml
packages/assets/internal/recipe/files/api.go.tmpl
packages/assets/internal/recipe/files/typescript-pnpm.md
packages/assets/internal/recipe/files/typescript-pnpm-SPEC.md

 exited 1 in 769ms:
packages/app/evals/saved_attempt_cli_test.go:17:func runSavedAttemptApp(t *testing.T, dir string, args ...string) (int, string, string) {
packages/app/evals/saved_attempt_cli_test.go:94:func assertSavedAttemptUnchanged(t *testing.T, root, head, change string) {
packages/app/evals/saved_attempt_cli_test.go:111:func Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext(t *testing.T) {
packages/app/evals/saved_attempt_cli_test.go:173:func Test_APP_087_MissingSavedAttemptRetainsLifecycleFailureShape(t *testing.T) {
packages/app/evals/saved_attempt_cli_test.go:190:func Test_APP_088_FromAttemptUsageRejectsBeforeMutation(t *testing.T) {
packages/app/evals/abandon_cli_test.go:30:func Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON(t *testing.T) {
packages/app/evals/abandon_cli_test.go:128:func Test_APP_098_APP_099_AbandonFailures(t *testing.T) {
packages/app/evals/run_cli_test.go:48:func Test_${prefix}_002_Added(t *testing.T) { if $module.Value() != "new" { t.Fatal("expected new value") } }
packages/app/evals/run_cli_test.go:111:		appPut(t, root, "packages/"+m.name+"/evals/existing_test.go", fmt.Sprintf("package evals\nimport (\"testing\"; \"example.invalid/run-cli/packages/%s\")\nfunc Test_%s_001_Value(t *testing.T) { if %s.Value() == \"\" { t.Fatal(\"empty\") } }\n", m.name, m.prefix, m.name))
packages/app/evals/run_cli_test.go:128:func runCLICalls(t *testing.T, f runCLISetup) []string {
packages/app/evals/run_cli_test.go:140:func runCLIData(t *testing.T, stdout string) (map[string]json.RawMessage, []map[string]json.RawMessage) {
packages/app/evals/run_cli_test.go:154:func runCLIFlag(t *testing.T, data map[string]json.RawMessage, key string, want bool) {
packages/app/evals/run_cli_test.go:162:func runCLISummaries(t *testing.T, applied []map[string]json.RawMessage, names ...string) {
packages/app/evals/run_cli_test.go:180:func runCLISuccess(t *testing.T, status int, out, errOut string, jsonMode bool) {
packages/app/evals/run_cli_test.go:187:func runCLIPlain(t *testing.T, out string, count int, name, state string) {
packages/app/evals/run_cli_test.go:205:func runCLIReport(t *testing.T, f runCLISetup, name, prefix, phase, want string) {
packages/app/evals/run_cli_test.go:227:func runCLIAssertCalls(t *testing.T, f runCLISetup, bin string, names ...string) {
packages/app/evals/run_cli_test.go:240:func runCLIAssertCall(t *testing.T, call, module, phase, root, bin string) {
packages/app/evals/run_cli_test.go:252:func Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation(t *testing.T) {
packages/app/evals/run_cli_test.go:314:func Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON(t *testing.T) {
packages/app/evals/run_cli_test.go:338:func Test_APP_083_RunPlainMultiple(t *testing.T) {
packages/app/evals/run_cli_test.go:350:func Test_APP_086_RunJSONLimited(t *testing.T) {
packages/app/evals/run_cli_test.go:362:func Test_APP_082_RunAllowUnmerged(t *testing.T) {
packages/app/evals/run_cli_test.go:376:func Test_APP_081_APP_084_RunUsageAndConfiguration(t *testing.T) {
packages/app/evals/run_cli_test.go:392:func Test_APP_084_APP_085_APP_086_RunWrappedFailure(t *testing.T) {
packages/app/evals/migrate_cli_test.go:99:func Test_APP_153_APP_154_MigrateNearestRootAndRepeatedNoOp(t *testing.T) {
packages/app/evals/migrate_cli_test.go:133:func Test_APP_153_APP_154_MigrateEmptyCurrentProject(t *testing.T) {
packages/app/evals/migrate_cli_test.go:151:func Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures(t *testing.T) {
packages/app/evals/migrate_cli_test.go:240:func Test_APP_155_MigrateRejectsArgumentsBeforeMutation(t *testing.T) {
packages/app/evals/authoring_commands_test.go:65:func Test_APP_006_APP_022_StampAliasesAndUnassignedHandles(t *testing.T) {
packages/app/evals/authoring_commands_test.go:108:func Test_APP_021_APP_022_StampErrorExits(t *testing.T) {
packages/app/evals/authoring_commands_test.go:141:func Test_APP_007_APP_024_CheckValidAndAllDiagnostics(t *testing.T) {
packages/app/evals/authoring_commands_test.go:198:func Test_APP_023_APP_024_CheckUsageAndConfiguration(t *testing.T) {
packages/app/evals/authoring_commands_test.go:225:func Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity(t *testing.T) {
packages/app/evals/authoring_commands_test.go:269:func Test_APP_026_ShowNamedErrors(t *testing.T) {
packages/app/evals/status_cli_test.go:246:func Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport(t *testing.T) {
packages/app/evals/status_cli_test.go:424:func Test_APP_118_APP_119_APP_122_StatusNearestRootEmptyContainers(t *testing.T) {
packages/app/evals/status_cli_test.go:470:func Test_APP_126_StatusUsageConfigurationAndOperationalFailures(t *testing.T) {
packages/app/evals/show_history_cli_test.go:54:func assertAppHistoryJSON(t *testing.T, output string, entries []project.SpecHistoryEntry) {
packages/app/evals/show_history_cli_test.go:103:func Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter(t *testing.T) {
packages/app/evals/show_history_cli_test.go:177:func Test_APP_132_APP_133_ShowHistoryNearestRootAndEmpty(t *testing.T) {
packages/app/evals/show_history_cli_test.go:215:func assertAppHistoryFailure(t *testing.T, status, wantStatus int, stdout, stderr string, jsonMode bool, code, named string) {
packages/app/evals/show_history_cli_test.go:237:func Test_APP_134_ShowHistoryOptionMisuse(t *testing.T) {
packages/app/evals/show_history_cli_test.go:261:func Test_APP_135_ShowHistoryLookupAndArchiveFailures(t *testing.T) {
packages/app/evals/next_cli_test.go:36:func Test_ALP_002_NewValue(t *testing.T) { if alpha.Value() != "new" { t.Fatal("expected new value") } }
packages/app/evals/next_cli_test.go:95:	appPut(t, root, "packages/alpha/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-cli/packages/alpha\")\nfunc Test_ALP_001_Value(t *testing.T) { if alpha.Value() == \"\" { t.Fatal(\"empty\") } }\n")
packages/app/evals/next_cli_test.go:111:func runNextCLI(t *testing.T, bin, dir string, args ...string) (int, string, string) {
packages/app/evals/next_cli_test.go:161:func Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange(t *testing.T) {
packages/app/evals/next_cli_test.go:232:func Test_APP_075_NextAllowUnmergedAndUsage(t *testing.T) {
packages/app/evals/next_cli_test.go:250:func Test_APP_076_NextPlainEmptyQueue(t *testing.T) {
packages/app/evals/next_cli_test.go:259:func Test_APP_077_NextExitClasses(t *testing.T) {
packages/app/evals/next_cli_test.go:276:func Test_APP_078_NextWrappedCauseAndAction(t *testing.T) {
packages/app/evals/next_cli_test.go:288:func Test_APP_079_NextJSONSuccessAndFailure(t *testing.T) {
packages/app/evals/new_module_cli_test.go:57:func Test_APP_100_APP_101_APP_104_NewModuleFromNearestProjectTemplates(t *testing.T) {
packages/app/evals/new_module_cli_test.go:127:func Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents(t *testing.T) {
packages/app/evals/new_module_cli_test.go:173:func Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors(t *testing.T) {
packages/app/evals/new_module_cli_test.go:218:func Test_APP_101_APP_104_NewModuleEmptyGoModule(t *testing.T) {
packages/app/evals/release_cli_test.go:30:func Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork(t *testing.T) {
packages/app/evals/release_cli_test.go:111:func Test_APP_092_APP_093_ReleaseForeignClaimKeepsOwner(t *testing.T) {
packages/app/evals/release_cli_test.go:157:func Test_APP_092_APP_093_ReleaseUsageAndConfiguration(t *testing.T) {
packages/app/evals/commits_cli_test.go:69:func Test_APP_094_APP_096_CommitsPendingAndArchivedOrder(t *testing.T) {
packages/app/evals/commits_cli_test.go:110:func Test_APP_095_APP_096_CommitsFailures(t *testing.T) {
packages/app/evals/list_cli_test.go:52:func Test_APP_114_APP_115_APP_117_ListQueueOrderAndFields(t *testing.T) {
packages/app/evals/list_cli_test.go:134:func Test_APP_114_APP_117_ListNearestRootAndEmptyQueue(t *testing.T) {
packages/app/evals/list_cli_test.go:183:func Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors(t *testing.T) {
packages/app/evals/diagnostic_causes_test.go:53:func Test_APP_037_StaticCauses(t *testing.T) {
packages/app/evals/diagnostic_causes_test.go:110:func Test_APP_037_APP_038_ShowStaticCauses(t *testing.T) {
packages/app/evals/diagnostic_causes_test.go:132:func Test_APP_038_LookupCauses(t *testing.T) {
packages/app/evals/diagnostic_causes_test.go:161:func Test_APP_039_BlockedNewChangeOutputPath(t *testing.T) {
packages/app/evals/diagnostic_causes_test.go:194:func Test_APP_040_ForwardedCheckCode(t *testing.T) {
packages/app/evals/pr_body_cli_test.go:127:func Test_APP_127_APP_128_APP_129_PRBodyExplicitHistoricalEndpoint(t *testing.T) {
packages/app/evals/pr_body_cli_test.go:132:func Test_APP_128_APP_129_PRBodyValidEmptyRange(t *testing.T) {
packages/app/evals/pr_body_cli_test.go:146:func Test_APP_130_APP_131_PRBodyFailures(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:19:func runLifecycleBinary(t *testing.T, bin, dir string, args ...string) (int, string, string) {
packages/app/evals/lifecycle_cli_test.go:103:func cliHold(t *testing.T, root string) {
packages/app/evals/lifecycle_cli_test.go:111:func cliGit(t *testing.T, root string, args ...string) string {
packages/app/evals/lifecycle_cli_test.go:123:func cliCommit(t *testing.T, root, message string) {
packages/app/evals/lifecycle_cli_test.go:129:func cliPhase(t *testing.T, root, want string) {
packages/app/evals/lifecycle_cli_test.go:137:func cliString(t *testing.T, output string) string {
packages/app/evals/lifecycle_cli_test.go:147:func cliJSONResult(t *testing.T, output string) map[string]json.RawMessage {
packages/app/evals/lifecycle_cli_test.go:160:func cliResultPhase(t *testing.T, output, phase string) {
packages/app/evals/lifecycle_cli_test.go:181:func Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:223:func Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:259:func Test_APP_004_HeldModulesUseApprovedProjection(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:297:func Test_APP_011_APP_032_APP_033_ContextCompactFullAndErrors(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:344:func Test_APP_012_APP_034_APP_035_APP_036_CoverageCurrentProjectionAndErrors(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:399:func Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:469:func Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits(t *testing.T) {
packages/app/evals/lifecycle_cli_test.go:514:func Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit(t *testing.T) {
packages/app/evals/run_for_cli_test.go:15:func runForBetaWorker(t *testing.T, f runCLISetup, blockedAlpha bool) {
packages/app/evals/run_for_cli_test.go:47:func runForReason(t *testing.T, reason, limit string) {
packages/app/evals/run_for_cli_test.go:54:func runForCalls(t *testing.T, f runCLISetup, bin string, want ...string) {
packages/app/evals/run_for_cli_test.go:69:func runForBetaHeld(t *testing.T, f runCLISetup, head string) (workflow.State, string) {
packages/app/evals/run_for_cli_test.go:88:func Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes(t *testing.T) {
packages/app/evals/run_for_cli_test.go:167:func Test_APP_090_APP_084_RunForPlainStop(t *testing.T) {
packages/app/evals/run_for_cli_test.go:205:func Test_APP_084_RunForPriorBlockStillFails(t *testing.T) {
packages/app/evals/run_for_cli_test.go:251:func Test_APP_089_RunForInvalidOptionsBeforeMutation(t *testing.T) {
packages/app/evals/project_commands_test.go:58:func runAppBinary(t *testing.T, bin, dir string, args ...string) (int, string, string) {
packages/app/evals/project_commands_test.go:75:func assertModuleJSON(t *testing.T, output string, names []string) {
packages/app/evals/project_commands_test.go:154:func Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor(t *testing.T) {
packages/app/evals/project_commands_test.go:196:func Test_APP_018_ModulesConfigurationErrors(t *testing.T) {
packages/app/evals/project_commands_test.go:259:func assertChangeSkeleton(t *testing.T, root, rel, module string) {
packages/app/evals/project_commands_test.go:290:func Test_APP_005_APP_020_NewChangePlainAndJSON(t *testing.T) {
packages/app/evals/project_commands_test.go:354:func Test_APP_019_NewChangeErrorsPreserveFiles(t *testing.T) {
packages/app/evals/project_commands_test.go:407:func Test_APP_020_NewChangeJSONFailureEnvelope(t *testing.T) {
packages/app/evals/doctor_setup_test.go:131:func Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI(t *testing.T) {
packages/app/evals/doctor_setup_test.go:215:func Test_APP_062_DoctorUsesCloserProject(t *testing.T) {
packages/app/evals/doctor_setup_test.go:231:func Test_APP_065_APP_066_APP_067_DoctorFailures(t *testing.T) {
packages/app/evals/doctor_setup_test.go:316:func Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite(t *testing.T) {
packages/app/evals/gc_cli_test.go:75:func Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty(t *testing.T) {
packages/app/evals/gc_cli_test.go:166:func Test_APP_150_APP_152_GCRejectsUsageBeforePruning(t *testing.T) {
packages/app/evals/gc_cli_test.go:186:func Test_APP_152_GCConfigurationAndGitFailures(t *testing.T) {
packages/app/evals/check_commits_cli_test.go:63:func Test_APP_105_APP_108_CheckCommitsRangesAndRecordedAuthorWork(t *testing.T) {
packages/app/evals/check_commits_cli_test.go:114:func Test_APP_105_APP_108_CheckCommitsSquashAndEmptyJSON(t *testing.T) {
packages/app/evals/check_commits_cli_test.go:162:func Test_APP_106_APP_108_CheckCommitsReportsEveryHistoryRule(t *testing.T) {
packages/app/evals/check_commits_cli_test.go:230:func Test_APP_107_APP_108_CheckCommitsFailures(t *testing.T) {
packages/app/evals/ci_cli_test.go:18:func TestMain(m *testing.M) {
packages/app/evals/ci_cli_test.go:127:func Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules(t *testing.T) {
packages/app/evals/ci_cli_test.go:169:func Test_APP_110_APP_112_APP_113_CIEvalAndBoundaryFailures(t *testing.T) {
packages/app/evals/ci_cli_test.go:184:func Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration(t *testing.T) {
packages/app/evals/ci_cli_test.go:195:func Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip(t *testing.T) {
packages/app/evals/ci_cli_test.go:202:	appPut(t, root, "packages/alpha/evals/alpha_test.go", "package alpha\nimport \"testing\"\nfunc Test_ALP_001_Fixture(t *testing.T) {}\n")
packages/app/evals/drill_cli_test.go:154:func Test_APP_136_APP_137_DrillPlainNearestProjectAndInvokingBin(t *testing.T) {
packages/app/evals/drill_cli_test.go:189:func Test_APP_138_DrillJSONSuccess(t *testing.T) {
packages/app/evals/drill_cli_test.go:209:func Test_APP_139_DrillArgumentAndModuleRejectionsDoNotInvokeHooks(t *testing.T) {
packages/app/evals/drill_cli_test.go:246:func Test_APP_140_DrillProjectAndConfigurationFailures(t *testing.T) {
packages/app/evals/drill_cli_test.go:263:func Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts(t *testing.T) {
packages/app/evals/drill_cli_test.go:306:func Test_APP_141_DrillSetupFailureHasNoArtifacts(t *testing.T) {
packages/app/evals/drill_cli_test.go:326:func Test_APP_142_DrillCleanUsageStopPlainAndJSON(t *testing.T) {
packages/app/evals/drill_cli_test.go:370:func Test_APP_136_DrillUsesConfiguredTimeout(t *testing.T) {
packages/app/evals/app_test.go:43:func Test_APP_001_StandaloneBinary(t *testing.T) {
packages/app/evals/app_test.go:70:func buildAppBinary(t *testing.T) string {
packages/app/evals/app_test.go:87:func Test_APP_002_InformationJSON(t *testing.T) {
packages/app/evals/app_test.go:116:func Test_APP_003_InvalidRequests(t *testing.T) {
packages/app/evals/app_test.go:165:func Test_APP_014_MaintenanceCommandsDispatch(t *testing.T) {
packages/app/evals/app_test.go:214:func Test_APP_015_FailedJSON(t *testing.T) {
packages/app/evals/app_test.go:222:func assertFailedJSON(t *testing.T, stdout string, wantCode ...string) string {
packages/app/evals/app_test.go:245:func Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents(t *testing.T) {
packages/app/evals/app_test.go:318:func run(args ...string) (int, string, string) {
packages/app/evals/app_test.go:324:func assertInformationOutput(t *testing.T, stdout, stderr, want string) {
packages/app/evals/init_test.go:78:func assertInitOutput(t *testing.T, status int, stdout, stderr, instructions string) {
packages/app/evals/init_test.go:88:func assertNoInitConfig(t *testing.T, root string) {
packages/app/evals/init_test.go:150:func Test_APP_041_APP_042_APP_043_APP_044_APP_045_APP_046_APP_050_APP_053_APP_054_APP_059_GoInit(t *testing.T) {
packages/app/evals/init_test.go:264:func Test_APP_043_APP_044_APP_047_APP_051_NoneInit(t *testing.T) {
packages/app/evals/init_test.go:303:func Test_APP_050_APP_052_NoneDefaultApp(t *testing.T) {
packages/app/evals/init_test.go:321:func Test_APP_053_InitFileOnlyWorkerCommand(t *testing.T) {
packages/app/evals/init_test.go:345:func Test_APP_054_InitFileOnlyAuditCommand(t *testing.T) {
packages/app/evals/init_test.go:369:func Test_APP_042_APP_048_APP_049_APP_050_APP_053_APP_054_APP_060_LocalInit(t *testing.T) {
packages/app/evals/init_test.go:412:func Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors(t *testing.T) {
packages/app/evals/init_test.go:471:func Test_APP_057_APP_058_InitConflicts(t *testing.T) {
packages/app/evals/init_test.go:499:func Test_APP_061_InitJSONSuccess(t *testing.T) {
packages/app/evals/usage_cli_test.go:35:func runCLIUsageAgent(t *testing.T, f runCLISetup) {
packages/app/evals/usage_cli_test.go:53:func cliUsagePolicy(t *testing.T, root, budget string) {
packages/app/evals/usage_cli_test.go:60:func cliStopEnvelope(t *testing.T, output, wait, budget string) (map[string]json.RawMessage, string) {
packages/app/evals/usage_cli_test.go:89:func cliStopReason(t *testing.T, reason, wait, budget string) {
packages/app/evals/usage_cli_test.go:99:func cliHeldState(t *testing.T, root, slug, prefix, head string) workflow.State {
packages/app/evals/usage_cli_test.go:121:func Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume(t *testing.T) {
packages/app/evals/usage_cli_test.go:207:func Test_APP_076_APP_077_NextPlainUsageStop(t *testing.T) {
packages/app/evals/usage_cli_test.go:234:func Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume(t *testing.T) {
packages/app/evals/usage_cli_test.go:314:func Test_APP_083_APP_084_RunPlainSharedUsageStop(t *testing.T) {
packages/app/evals/usage_cli_test.go:337:func Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop(t *testing.T) {
packages/app/evals/usage_cli_test.go:407:func Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation(t *testing.T) {
packages/app/evals/check_fix_cli_test.go:104:func Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases(t *testing.T) {
packages/app/evals/check_fix_cli_test.go:156:func Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp(t *testing.T) {
packages/app/evals/check_fix_cli_test.go:179:func Test_APP_148_APP_149_CheckFixRejectsBeforeMutation(t *testing.T) {
packages/app/evals/show_stale_cli_test.go:48:func Test_APP_143_APP_144_StaleRecoveryDiffAndReadOnly(t *testing.T) {
packages/app/evals/show_stale_cli_test.go:125:func Test_APP_144_StaleUnavailableHistoryAndEmptyView(t *testing.T) {
packages/app/evals/show_stale_cli_test.go:143:func Test_APP_143_StaleViewChoosesNearestProject(t *testing.T) {
packages/app/evals/show_stale_cli_test.go:157:func Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures(t *testing.T) {
fireweed: '>=0.1, <0.2'
version: 1
remote: origin
modules:
    - packages/*
module_defaults:
    interface:
        - api.go
    evals:
        - evals/**
    disposable:
        - internal/**
hooks:
    eval: go run .fireweed/hooks.go eval {module_dir} {report}
    boundaries: go run .fireweed/hooks.go boundaries
    setup: sh .fireweed/setup.sh
protected:
    - .github/**
    - docs/**
    - fireweed-design-v1.md
    - go.mod
    - main.go
    - Makefile
    - scripts/**
    - .goreleaser.yaml
    - README.md
    - LICENSE
    - .gitignore
worker_writable:
    - go.sum
agent:
    command:
        - codex
        - exec
        - --ephemeral
        - --ignore-user-config
        - -m
        - gpt-6.1-sol
        - -c
        - model_reasoning_effort="high"
        - -s
        - workspace-write
        - '{prompt}'
audit:
    enabled: true
    max_rounds: 2
    command:
        - codex
        - exec
        - --ephemeral
        - --ignore-user-config
        - -m
        - gpt-6-astra
        - -c
        - model_reasoning_effort="medium"
        - -s
        - workspace-write
        - '{prompt}'
authoring:
    max_unapproved: 10
attempts: 2
timeout: 90m
verify_dependents: direct
limit_patterns:
    - usage limit
    - rate limit
limit_wait: 30m
limit_budget: 4h
bin
canaries
hooks.go
recipe.md
setup.sh
templates
U02: change "app" not found; no valid change slugs. Use a change held by this worktree and retry context.

exec
/bin/zsh -lc "cat packages/app/evals/app_test.go; sed -n '75,154p' packages/app/evals/project_commands_test.go; sed -n '181,399p' packages/app/evals/lifecycle_cli_test.go; cat .fireweed/hooks.go | head -170; rg -n '"'^### APP-|''^'"\\[' packages/app/SPEC.md | tail -30" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc "cat packages/{contracts,project,workflow,runner,checks,assets}/model/*.go; rg -n 'func (Load|LoadAt|Check|StampMissing|Project|Find|Discover)|Approved|Snapshot|func .*Config' packages/{project,workflow}/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 639ms:
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
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// Snapshot is the project data read from a repository root.
type Snapshot struct {
	Root    string
	Config  map[string]any
	Modules []Module
}

// Module is a discovered module and its declared document and zones.
type Module struct {
	Name      string
	IDPrefix  string
	Dir       string
	DependsOn []string
	Zones     contracts.Zones
	Document  contracts.Document
}

// Change is one pending change document and its spec edits.
type Change struct {
	Number   int
	Slug     string
	Path     string
	Module   string
	Document contracts.Document
	Added    []contracts.Item
	Modified []contracts.Item
	Removed  []contracts.Item
}

// SpecHistoryEntry describes one archived change affecting selected specs.
type SpecHistoryEntry struct {
	Path               string
	Module             string
	IDs                []string
	Motivation         string
	Decisions          string
	RemovedDecisions   string
	Provenance         string
	ImplementationDate time.Time
}

// StaleSpec describes a recorded base that differs from the preceding projection.
type StaleSpec struct {
	ID           string
	RecordedHash string
	Before       string
	Current      string
	Recovered    bool
}

// Diagnostic reports a static authoring check.
type Diagnostic struct {
	Code    string
	Message string
	Path    string
	Warning bool
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
)

type State struct {
	Change      project.Change
	Snapshot    project.Snapshot
	Run         string
	Phase       string
	Head        string
	PhaseBase   string
	ClaimOID    string
	Diagnostics []project.Diagnostic
}

type Request struct {
	Root    string
	Change  string
	Bin     string
	Agent   string
	Quick   bool
	Timeout time.Duration
}

type Result struct {
	State       State
	Reports     map[string]checks.Report
	Costs       []ModuleCost
	Diagnostics []project.Diagnostic
}

type ModuleCost struct {
	Module          string
	Seconds         float64
	CasesPerSpec    map[string]int
	EvalLines       int
	DisposableLines int
}

type CostSnapshot struct {
	Run         string
	Phase       string
	RecordedAt  time.Time
	Modules     []ModuleCost
	Diagnostics []project.Diagnostic
}

type StatusReport struct {
	Modules         []project.Module
	Queue           []StatusEntry
	Quarantines     []Quarantine
	BoundaryOptOuts []BoundaryOptOut
	DriftSinceDrill map[string]int
	Unapproved      int
	PhaseCommits    map[string]int
	Costs           []ModuleCost
	Diagnostics     []project.Diagnostic
}

type PRBodyReport struct {
	Body        string
	Diagnostics []project.Diagnostic
}

type StatusEntry struct {
	Slug        string
	Path        string
	Module      string
	Phase       string
	Blocked     string
	WaitingOn   string
	WaitingRule string
	AuditRounds int
}

type Quarantine struct {
	Module string
	ID     string
	Reason string
}

type BoundaryOptOut struct {
	Module string
	Reason string
}

type BlockRequest struct {
	Root   string
	Change string
	Reason string
}

type BlockResult struct {
	State      State
	AttemptRef string
	AttemptOID string
}

type GCRequest struct {
	Root string
	Now  time.Time
}

type GCResult struct {
	Pruned []string
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

type AgentRequest struct {
	Snapshot project.Snapshot
	Role     string
	Prompt   string
	Bin      string
	Timeout  time.Duration
	LogRoot  string
	Run      string
	Phase    string
	Attempt  int
}

type AgentResult struct {
	Stdout     string
	Stderr     string
	ExitCode   int
	UsageLimit bool
	LogPath    string
}

type NextRequest struct {
	Root          string
	Bin           string
	AllowUnmerged bool
	Timeout       time.Duration
}

type StepResult struct {
	Phase string
	Role  string
	Agent AgentResult
}

type NextResult struct {
	State workflow.State
	Empty bool
	Steps []StepResult
}

type RunRequest struct {
	Root          string
	Bin           string
	AllowUnmerged bool
	Timeout       time.Duration
	Max           int
	For           time.Duration
}

type RunResult struct {
	Applied []NextResult
	Empty   bool
	Limited bool
}

type DrillRequest struct {
	Root    string
	Module  string
	Bin     string
	Timeout time.Duration
}

type DrillResult struct {
	Module     string
	Base       string
	Date       string
	Agent      string
	Model      string
	Result     string
	ReportPath string
	Ref        string
	Own        checks.Report
	Dependents map[string]checks.Report
	Steps      []AgentResult
}
package model

import (
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
)

type Case struct {
	Name      string
	Classname string
	IDs       []string
	Result    string
	Failure   string
	Seconds   float64
}

type Report struct {
	Cases    []Case
	Seconds  float64
	HookExit int
	Stdout   string
	Stderr   string
}

type HookRequest struct {
	Snapshot project.Snapshot
	Module   string
	Bin      string
	Timeout  time.Duration
}

type BoundaryRequest struct {
	Snapshot project.Snapshot
	Bin      string
	Timeout  time.Duration
}

type ZoneRequest struct {
	Snapshot          project.Snapshot
	Module            string
	Phase             string
	ChangePath        string
	Base              string
	InterfaceDeclared bool
	Changes           []gitstore.Change
}

type DoctorRequest struct {
	Root    string
	Bin     string
	NoAgent bool
	Timeout time.Duration
}

type DoctorCheck struct {
	Name   string
	Module string
	Passed bool
	Detail string
}

type DoctorReport struct {
	Checks      []DoctorCheck
	Diagnostics []project.Diagnostic
}

type HistoryReport struct {
	Commits     []string
	Squashed    []SquashedChange
	Diagnostics []project.Diagnostic
}

type SquashedChange struct {
	Commit string
	Change string
}
package model

import "github.com/lurkshark/fireweed/packages/contracts"

// Recipe is the data copied into a project during setup.
type Recipe struct {
	Name           string
	Instructions   string
	ModuleDefaults contracts.Zones
	Hooks          map[string]any
	Protected      []string
	WorkerWritable []string
	Files          map[string]string
}
packages/project/internal/queue/project.go:21:	snapshot model.Snapshot
packages/project/internal/queue/project.go:25:func newProjection(snapshot model.Snapshot) (*projection, error) {
packages/project/internal/queue/project.go:163:func (p *projection) result() model.Snapshot {
packages/project/internal/queue/project.go:182:func Project(root, through string) (model.Snapshot, error) {
packages/project/internal/queue/project.go:185:		return model.Snapshot{}, err
packages/project/internal/queue/project.go:189:		return model.Snapshot{}, err
packages/project/internal/queue/project.go:195:			return model.Snapshot{}, err
packages/project/internal/queue/project.go:201:		return model.Snapshot{}, err
packages/project/internal/queue/project.go:205:			return model.Snapshot{}, err
packages/project/internal/queue/project.go:209:		return model.Snapshot{}, err
packages/workflow/internal/lifecycle/stub.go:33:	Snapshot  project.Snapshot  `json:"snapshot"`
packages/workflow/internal/lifecycle/stub.go:182:	return model.State{Change: ch, Snapshot: c.Snapshot, Run: c.Run, Phase: c.Phase, Head: head, PhaseBase: c.PhaseBase, ClaimOID: oid}
packages/workflow/internal/lifecycle/stub.go:218:func assign(root string, ch project.Change, snap project.Snapshot) (project.Change, error) {
packages/workflow/internal/lifecycle/stub.go:235:	baseSnapshot, _ := project.Load(root)
packages/workflow/internal/lifecycle/stub.go:236:	for _, m := range baseSnapshot.Modules {
packages/workflow/internal/lifecycle/stub.go:375:	c := claim{Worktree: root, Change: ch.Slug, Run: run, Phase: phase, PhaseBase: repo.Head, Snapshot: snap}
packages/workflow/internal/lifecycle/stub.go:408:	c.Snapshot = snap
packages/workflow/internal/lifecycle/stub.go:419:func mustLoad(root string) project.Snapshot { s, _ := project.Load(root); return s }
packages/workflow/internal/lifecycle/stub.go:470:	c, _, e := existing(s.Snapshot.Root, s.Change.Slug)
packages/workflow/internal/lifecycle/stub.go:478:func module(s project.Snapshot, name string) project.Module {
packages/workflow/internal/lifecycle/stub.go:557:	m := module(c.Snapshot, s.Change.Module)
packages/workflow/internal/lifecycle/stub.go:573:func dependents(s project.Snapshot, name, mode string) []string {
packages/workflow/internal/lifecycle/stub.go:602:	dir := filepath.Join(s.Snapshot.Root, ".fireweed", "runs", s.Run, phase)
packages/workflow/internal/lifecycle/stub.go:623:	if e = hard(project.Check(s.Snapshot.Root)); e != nil {
packages/workflow/internal/lifecycle/stub.go:629:	diff, e := stepChanges(s.Snapshot.Root, c.PhaseBase)
packages/workflow/internal/lifecycle/stub.go:633:	r.Diagnostics = checks.Zones(checks.ZoneRequest{Snapshot: c.Snapshot, Module: s.Change.Module, Phase: c.Phase, ChangePath: s.Change.Path, Base: c.PhaseBase, InterfaceDeclared: strings.TrimSpace(c.Sections["Interface"]) != "", Changes: diff})
packages/workflow/internal/lifecycle/stub.go:644:		mode := str(c.Snapshot.Config["verify_dependents"])
packages/workflow/internal/lifecycle/stub.go:648:		names = append(names, dependents(c.Snapshot, s.Change.Module, mode)...)
packages/workflow/internal/lifecycle/stub.go:652:		rep, err := checks.Evaluate(ctx, checks.HookRequest{Snapshot: c.Snapshot, Module: name, Bin: req.Bin, Timeout: req.Timeout})
packages/workflow/internal/lifecycle/stub.go:655:		cost, costErr := measureModuleCost(c.Snapshot, name, rep)
packages/workflow/internal/lifecycle/stub.go:666:		items := specs(module(c.Snapshot, name))
packages/workflow/internal/lifecycle/stub.go:700:		ds, err := checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: c.Snapshot, Bin: req.Bin, Timeout: req.Timeout})
packages/workflow/internal/lifecycle/stub.go:728:	warnings := costWarnings(r.Costs, c.Snapshot.Config)
packages/workflow/internal/lifecycle/stub.go:761:	root := s.Snapshot.Root
packages/workflow/internal/lifecycle/stub.go:797:			for _, it := range specs(module(c.Snapshot, ch.Module)) {
packages/workflow/internal/lifecycle/stub.go:839:		if untestable || round >= num(mapval(c.Snapshot.Config["audit"])["max_rounds"], 2) {
packages/workflow/internal/lifecycle/stub.go:857:	for _, id := range append(selectedIDs(ch, specs(module(c.Snapshot, ch.Module))), evalFixIDs(ch)...) {
packages/workflow/internal/lifecycle/stub.go:884:	root := s.Snapshot.Root
packages/workflow/internal/lifecycle/stub.go:885:	m := module(c.Snapshot, ch.Module)
packages/workflow/internal/lifecycle/stub.go:987:	fresh, e := Current(s.Snapshot.Root, s.Change.Slug)
packages/workflow/internal/lifecycle/stub.go:991:	c, oid, e := existing(s.Snapshot.Root, s.Change.Slug)
packages/workflow/internal/lifecycle/stub.go:998:	if e = hard(project.Check(s.Snapshot.Root)); e != nil {
packages/workflow/internal/lifecycle/stub.go:1001:	diffAfter, e := stepChanges(s.Snapshot.Root, c.PhaseBase)
packages/workflow/internal/lifecycle/stub.go:1005:	postZones := checks.Zones(checks.ZoneRequest{Snapshot: c.Snapshot, Module: fresh.Change.Module, Phase: c.Phase, ChangePath: fresh.Change.Path, Base: c.PhaseBase, InterfaceDeclared: strings.TrimSpace(c.Sections["Interface"]) != "", Changes: diffAfter})
packages/workflow/internal/lifecycle/stub.go:1019:		if mapval(c.Snapshot.Config["audit"])["enabled"] == false {
packages/workflow/internal/lifecycle/stub.go:1036:	if e = unstageRunArtifacts(s.Snapshot.Root, c.PhaseBase); e != nil {
packages/workflow/internal/lifecycle/stub.go:1060:		if e = write(s.Snapshot.Root, ch.Path, []byte(rendered)); e != nil {
packages/workflow/internal/lifecycle/stub.go:1063:		head, e := gitstore.Commit(s.Snapshot.Root, subject, paths)
packages/workflow/internal/lifecycle/stub.go:1067:		if e = gitstore.CASRef(s.Snapshot.Root, ref(ch.Slug), oid, ""); e != nil {
packages/workflow/internal/lifecycle/stub.go:1084:	raw, e := read(s.Snapshot.Root, s.Change.Path)
packages/workflow/internal/lifecycle/stub.go:1090:	m := module(s.Snapshot, s.Change.Module)
packages/workflow/internal/lifecycle/stub.go:1099:		_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
packages/workflow/internal/lifecycle/stub.go:1103:			rel, _ := filepath.Rel(file(s.Snapshot.Root, m.Dir), path)
packages/workflow/internal/lifecycle/stub.go:1115:	sys, e := read(s.Snapshot.Root, "SYSTEM.md")
packages/workflow/internal/lifecycle/stub.go:1130:	_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
packages/workflow/internal/lifecycle/stub.go:1134:		rel, _ := filepath.Rel(s.Snapshot.Root, path)
packages/project/internal/queue/check.go:15:func Check(root string) []model.Diagnostic {
packages/project/internal/queue/check.go:112:func modulePrefix(snapshot model.Snapshot, name string) string {
packages/workflow/internal/lifecycle/block_stub.go:123:	root := s.Snapshot.Root
packages/workflow/internal/lifecycle/regenerate.go:17:func removeDisposable(root string, snapshot project.Snapshot, name string) error {
packages/project/internal/queue/queue.go:27:func readQueue(root string, snapshot model.Snapshot) ([]model.Change, error) {
packages/workflow/internal/lifecycle/status_costs.go:55:func selectedCostWarning(index int, snapshot model.CostSnapshot, selected map[string]int, warning project.Diagnostic) bool {
packages/project/internal/queue/stale.go:108:func StampMissing(root, target string) error {
packages/project/internal/fixture/load_at.go:17:func LoadAt(root, revision string) (model.Snapshot, error) {
packages/project/internal/fixture/load_at.go:20:		return model.Snapshot{}, fmt.Errorf("find repository root: %w", err)
packages/project/internal/fixture/load_at.go:24:		return model.Snapshot{}, fmt.Errorf("canonical repository root: %w", err)
packages/project/internal/fixture/load_at.go:28:		return model.Snapshot{}, fmt.Errorf("resolve revision %q: %w", revision, err)
packages/project/internal/fixture/load_at.go:34:		return model.Snapshot{}, fmt.Errorf("read fireweed.yaml at %s: %w", commit, err)
packages/project/internal/fixture/load_at.go:38:		return model.Snapshot{}, err
packages/project/internal/fixture/load_at.go:43:		return model.Snapshot{}, fmt.Errorf("list module directories at %s: %w", commit, err)
packages/project/internal/fixture/load_at.go:45:	snapshot := model.Snapshot{Root: canonical, Config: config}
packages/project/internal/fixture/load_at.go:64:			return model.Snapshot{}, fmt.Errorf("read %s at %s: %w", path, commit, err)
packages/project/internal/fixture/load_at.go:68:			return model.Snapshot{}, err
packages/workflow/internal/lifecycle/endpoint_config_error_stub.go:7:func (e *endpointConfigError) Error() string { return e.cause.Error() }
packages/workflow/internal/lifecycle/endpoint_config_error_stub.go:8:func (e *endpointConfigError) Unwrap() error { return e.cause }
packages/workflow/internal/lifecycle/endpoint_config_error_stub.go:10:func IsEndpointConfigError(err error) bool {
packages/workflow/internal/lifecycle/pr_body_stub.go:120:	cache := map[string]project.Snapshot{resolved[1]: endpoint}
packages/workflow/internal/lifecycle/pr_body_stub.go:229:func prRootProtected(snapshot project.Snapshot, file string) bool {
packages/workflow/internal/lifecycle/pr_body_stub.go:241:func prRenderChange(body *strings.Builder, root string, change *prChange, snapshots map[string]project.Snapshot) error {
packages/workflow/internal/lifecycle/pr_body_stub.go:387:func prExceptions(body *strings.Builder, endpoint project.Snapshot, paths []string, changes map[string]*prChange) error {
packages/workflow/internal/lifecycle/pr_body_stub.go:419:func prCosts(body *strings.Builder, root string, endpoint project.Snapshot, result *model.PRBodyReport) error {
packages/workflow/internal/lifecycle/pr_body_stub.go:474:func prArtifacts(body *strings.Builder, root, base, endpoint string, snapshot project.Snapshot, result *model.PRBodyReport) error {
packages/workflow/internal/lifecycle/costs_stub.go:20:func measureModuleCost(snapshot project.Snapshot, name string, report checks.Report) (model.ModuleCost, error) {
packages/workflow/internal/lifecycle/costs_stub.go:123:	snapshot := model.CostSnapshot{Run: state.Run, Phase: phase, RecordedAt: time.Now().UTC(), Modules: costs, Diagnostics: warnings}
packages/workflow/internal/lifecycle/costs_stub.go:128:	path := costPath(state.Snapshot.Root, state.Run, phase)
packages/workflow/internal/lifecycle/costs_stub.go:135:func Costs(root string) ([]model.CostSnapshot, error) {
packages/workflow/internal/lifecycle/costs_stub.go:139:		return []model.CostSnapshot{}, nil
packages/workflow/internal/lifecycle/costs_stub.go:144:	snapshots := []model.CostSnapshot{}
packages/workflow/internal/lifecycle/costs_stub.go:164:			var snapshot model.CostSnapshot
packages/project/internal/fixture/fixture.go:27:func LoadCode(err error) string {
packages/project/internal/fixture/fixture.go:38:func LoadPath(err error) string {
packages/project/internal/fixture/fixture.go:59:func Load(root string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:62:		return model.Snapshot{}, fmt.Errorf("read fireweed.yaml: %w", err)
packages/project/internal/fixture/fixture.go:66:		return model.Snapshot{}, err
packages/project/internal/fixture/fixture.go:69:	snapshot := model.Snapshot{Root: root, Config: config}
packages/project/internal/fixture/fixture.go:100:		return model.Snapshot{}, err
packages/project/internal/fixture/fixture.go:105:func parseConfig(data []byte) (map[string]any, []string, map[string]any, error) {

 succeeded in 649ms:
package evals_test

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/app"
	"github.com/lurkshark/fireweed/packages/assets"
)

type informationCase struct {
	name string
	args []string
	want string
}

func informationCases(t *testing.T) []informationCase {
	t.Helper()
	explanation, err := assets.Explain("C01")
	if err != nil {
		t.Fatal(err)
	}
	playbook, err := assets.Playbook("evals")
	if err != nil {
		t.Fatal(err)
	}
	return []informationCase{
		{"help", []string{"help"}, assets.Help()},
		{"version", []string{"version"}, assets.Version()},
		{"explain", []string{"explain", "C01"}, explanation},
		{"playbook", []string{"playbook", "evals"}, playbook},
	}
}

func Test_APP_001_StandaloneBinary(t *testing.T) {
	bin := buildAppBinary(t)

	outside := t.TempDir()
	malformed := t.TempDir()
	if err := os.WriteFile(filepath.Join(malformed, "fireweed.yaml"), []byte("[malformed yaml\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, location := range []struct {
		name string
		dir  string
	}{{"outside-project", outside}, {"invalid-config", malformed}} {
		for _, tc := range informationCases(t) {
			t.Run(location.name+"/"+tc.name, func(t *testing.T) {
				cmd := exec.Command(bin, tc.args...)
				cmd.Dir = location.dir
				var stdout, stderr bytes.Buffer
				cmd.Stdout, cmd.Stderr = &stdout, &stderr
				if err := cmd.Run(); err != nil {
					t.Errorf("%v: %v", tc.args, err)
				}
				assertInformationOutput(t, stdout.String(), stderr.String(), tc.want)
			})
		}
	}
}

func buildAppBinary(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate eval source")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Dir = root
	build.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build NEW binary: %v\n%s", err, output)
	}
	return bin
}

func Test_APP_002_InformationJSON(t *testing.T) {
	for _, tc := range informationCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			plainStatus, plain, plainErr := run(tc.args...)
			if plainStatus != 0 || plainErr != "" {
				t.Errorf("plain status=%d stderr=%q; want 0 and empty", plainStatus, plainErr)
			}
			assertInformationOutput(t, plain, plainErr, tc.want)
			args := append(append([]string{}, tc.args...), "--json")
			status, stdout, stderr := run(args...)
			if status != 0 || stderr != "" {
				t.Errorf("status=%d stderr=%q; want 0 and empty", status, stderr)
			}
			envelope := decodeEnvelope(t, stdout)
			var data string
			if err := json.Unmarshal(envelope["data"], &data); err != nil {
				t.Fatalf("data must be a string: %v", err)
			}
			if data != plain {
				t.Errorf("data=%q; want plain result %q", data, plain)
			}
			var diagnostics []json.RawMessage
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
				t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
			}
		})
	}
}

func Test_APP_003_InvalidRequests(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		bad  string
		code string
	}{
		{"unknown-command", []string{"not-a-command"}, "not-a-command", "U01"},
		{"unknown-option", []string{"help", "--bogus"}, "--bogus", "U01"},
		{"missing-argument", []string{"explain"}, "explain", "U01"},
		{"extra-argument", []string{"version", "surplus"}, "version", "U01"},
		{"new-change-arity", []string{"new-change", "alpha"}, "new-change", "U01"},
		{"stamp-arity", []string{"stamp"}, "stamp", "U01"},
		{"check-arity", []string{"check", "extra"}, "check", "U01"},
		{"show-arity", []string{"show"}, "show", "U01"},
		{"start-arity", []string{"start"}, "start", "U01"},
		{"context-arity", []string{"context"}, "context", "U01"},
		{"coverage-arity", []string{"coverage"}, "coverage", "U01"},
		{"verify-arity", []string{"verify", "manual", "extra"}, "verify", "U01"},
		{"APP_013_unknown-explanation", []string{"explain", "NO-SUCH-CODE"}, "NO-SUCH-CODE", "U02"},
		{"APP_013_unknown-playbook", []string{"playbook", "no-such-playbook"}, "no-such-playbook", "U02"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, format := range []struct {
				name string
				json bool
			}{{"plain", false}, {"JSON", true}} {
				args := append([]string(nil), tc.args...)
				if format.json {
					args = append(args, "--json")
				}
				status, stdout, stderr := run(args...)
				if status != 2 {
					t.Errorf("%s status=%d; want 2", format.name, status)
				}
				diagnostic := stdout + stderr
				if format.json {
					diagnostic = assertFailedJSON(t, stdout, tc.code)
				} else if !strings.HasPrefix(strings.TrimSpace(diagnostic), tc.code+":") {
					t.Errorf("%s diagnostic %q must have code %s", format.name, diagnostic, tc.code)
				}
				if !strings.Contains(diagnostic, tc.bad) || !corrective(diagnostic) {
					t.Errorf("%s diagnostic %q must identify %q and suggest a correction", format.name, diagnostic, tc.bad)
				}
			}
		})
	}
}

func Test_APP_014_MaintenanceCommandsDispatch(t *testing.T) {
	bin := buildAppBinary(t)
	root := migrateProject(t)
	for _, command := range []string{"gc", "migrate"} {
		t.Run(command, func(t *testing.T) {
			if !strings.Contains(assets.Help(), command) {
				t.Fatalf("%q is not listed by help", command)
			}
			status, plain, stderr := runAppBinary(t, bin, root, command)
			if status != 0 || stderr != "" || strings.TrimSpace(plain) == "" {
				t.Fatalf("%s plain: status=%d stdout=%q stderr=%q", command, status, plain, stderr)
			}
			if regexp.MustCompile(`(?i)unimplemented|not implemented|under construction`).MatchString(plain) {
				t.Errorf("%s still reported a placeholder: %q", command, plain)
			}
			status, output, stderr := runAppBinary(t, bin, root, command, "--json")
			if status != 0 || stderr != "" {
				t.Fatalf("%s JSON: status=%d stdout=%q stderr=%q", command, status, output, stderr)
			}
			env := decodeEnvelope(t, output)
			if string(env["data"]) == "null" {
				t.Errorf("%s JSON data is null", command)
			}
			var diagnostics []json.RawMessage
			if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
				t.Errorf("%s JSON diagnostics=%s: %v", command, env["diagnostics"], err)
			}
			for _, args := range [][]string{{command, "extra"}, {command, "--bogus"}} {
				for _, jsonMode := range []bool{false, true} {
					call := append([]string(nil), args...)
					if jsonMode {
						call = append(call, "--json")
					}
					status, stdout, errOut := runAppBinary(t, bin, root, call...)
					message := commandFailure(t, status, 2, stdout, errOut, jsonMode)
					if jsonMode {
						message = assertFailedJSON(t, stdout, "U01")
					} else if !strings.HasPrefix(strings.TrimSpace(errOut), "U01:") {
						t.Errorf("%s invalid plain diagnostic=%q; want U01", command, errOut)
					}
					if !corrective(message) {
						t.Errorf("%s invalid diagnostic lacks guidance: %q", command, message)
					}
				}
			}
		})
	}
}

func Test_APP_015_FailedJSON(t *testing.T) {
	status, stdout, _ := run("init", "--json")
	if status == 0 {
		t.Error("failed request returned exit 0")
	}
	assertFailedJSON(t, stdout)
}

func assertFailedJSON(t *testing.T, stdout string, wantCode ...string) string {
	t.Helper()
	envelope := decodeEnvelope(t, stdout)
	if string(envelope["data"]) != "null" {
		t.Errorf("data=%s; want null", envelope["data"])
	}
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
		t.Fatalf("diagnostics=%s; want nonempty array: %v", envelope["diagnostics"], err)
	}
	if strings.TrimSpace(diagnostics[0].Code) == "" || strings.TrimSpace(diagnostics[0].Message) == "" {
		t.Errorf("diagnostic needs code and message: %+v", diagnostics[0])
	}
	if len(wantCode) > 0 && diagnostics[0].Code != wantCode[0] {
		t.Errorf("diagnostic code=%q; want %q", diagnostics[0].Code, wantCode[0])
	}
	return diagnostics[0].Message
}

// The app-level convention probe uses one PATH recorder for external commands.
func Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	appPut(t, root, "packages/alpha/SPEC.md", authoringSpec)
	authoringChange(t, root, "0001-first.md", "Users need a queued item.\n", "### new:queued\nQueued MUST work.\n", "")
	nested := filepath.Join(root, "packages", "alpha", "nested")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	binDir := t.TempDir()
	logPath := filepath.Join(t.TempDir(), "invocations")
	recorder := []byte("#!/bin/sh\nprintf '%s\\n' \"$0 $*\" >> \"$FIREWEED_RECORDER\"\nexit 77\n")
	for _, name := range []string{"git", "hook", "codex"} {
		if err := os.WriteFile(filepath.Join(binDir, name), recorder, 0700); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("FIREWEED_RECORDER", logPath)
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	for _, tc := range informationCases(t) {
		status, stdout, stderr := runAppBinary(t, bin, nested, tc.args...)
		if status != 0 {
			t.Errorf("%s status=%d stderr=%q", tc.name, status, stderr)
		}
		assertInformationOutput(t, stdout, stderr, tc.want)
	}
	runAppBinary(t, bin, nested, "modules", "--json")
	for _, args := range [][]string{
		{"list"}, {"list", "--json"},
		{"show", "alpha"}, {"show", "alpha", "ALP-001"},
		{"show", "alpha", "ALP-001", "--history"}, {"show", "alpha", "ALP-001", "--history", "--json"},
		{"show", "alpha", "--projected"}, {"show", "alpha", "new:queued", "--projected"},
	} {
		runAppBinary(t, bin, nested, args...)
	}
	if calls, err := os.ReadFile(logPath); err == nil {
		t.Errorf("read commands invoked external tools: %s", calls)
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
	// Set up the repository with real Git before recording status reads.
	gitPath := filepath.Join(binDir, "git")
	if err := os.Remove(gitPath); err != nil {
		t.Fatal(err)
	}
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "read command fixture")
	// Status may read Git. This recorder forwards every invocation without
	// inspecting its arguments; the existing hook and agent sentinels remain.
	gitRecorder := []byte("#!/bin/sh\nprintf '%s\\n' \"$0 $*\" >> \"$FIREWEED_RECORDER\"\nexec \"$FIREWEED_REAL_GIT\" \"$@\"\n")
	if err := os.WriteFile(filepath.Join(binDir, "git"), gitRecorder, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("FIREWEED_REAL_GIT", realGit)
	for _, args := range [][]string{{"status"}, {"status", "--json"}, {"pr-body", "HEAD..HEAD"}, {"pr-body", "HEAD..HEAD", "--json"}} {
		runAppBinary(t, bin, nested, args...)
	}
	calls, err := os.ReadFile(logPath)
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(calls), "\n") {
		if strings.Contains(line, "/hook ") || strings.Contains(line, "/codex ") ||
			strings.Contains(" "+line+" ", " fetch ") || strings.Contains(" "+line+" ", " push ") {
			t.Errorf("status read invoked hook, agent, fetch or push: %q", line)
		}
	}
}

func run(args ...string) (int, string, string) {
	var stdout, stderr bytes.Buffer
	status := app.Run(args, &stdout, &stderr)
	return status, stdout.String(), stderr.String()
}

func assertInformationOutput(t *testing.T, stdout, stderr, want string) {
	t.Helper()
	if stderr != "" {
		t.Errorf("stderr=%q; want empty", stderr)
	}
	if strings.TrimSuffix(stdout, "\n") != strings.TrimSuffix(want, "\n") {
		t.Errorf("stdout=%q; want assets text %q", stdout, want)
	}
}

func decodeEnvelope(t *testing.T, output string) map[string]json.RawMessage {
	t.Helper()
	dec := json.NewDecoder(strings.NewReader(output))
	var envelope map[string]json.RawMessage
	if err := dec.Decode(&envelope); err != nil || envelope == nil {
		t.Fatalf("stdout=%q; want JSON object: %v", output, err)
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		t.Fatalf("stdout=%q; want exactly one JSON object: %v", output, err)
	}
	var schema int
	if err := json.Unmarshal(envelope["schema"], &schema); err != nil || schema != 1 {
		t.Errorf("schema=%s; want 1", envelope["schema"])
	}
	return envelope
}

func corrective(message string) bool {
	return regexp.MustCompile(`(?i)\b(help|usage|valid|available|expected|required|requires|try|use|choose|provide|specify|correct|remove|rename)\b`).MatchString(message)
}
func assertModuleJSON(t *testing.T, output string, names []string) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
	}
	var modules []struct {
		Name      string   `json:"name"`
		IDPrefix  string   `json:"id_prefix"`
		Dir       string   `json:"dir"`
		DependsOn []string `json:"depends_on"`
		Zones     struct {
			Interface  []string `json:"interface"`
			Evals      []string `json:"evals"`
			Disposable []string `json:"disposable"`
			Preserve   []string `json:"preserve"`
		} `json:"zones"`
	}
	if err := json.Unmarshal(envelope["data"], &modules); err != nil || modules == nil {
		t.Fatalf("data=%s; want module array: %v", envelope["data"], err)
	}
	var rawModules []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &rawModules); err != nil {
		t.Fatalf("decode module fields: %v", err)
	}
	got := map[string]bool{}
	for i, module := range modules {
		for _, key := range []string{"name", "id_prefix", "dir", "depends_on", "zones"} {
			if _, ok := rawModules[i][key]; !ok {
				t.Errorf("module %q missing %s field", module.Name, key)
			}
		}
		var zoneFields map[string]json.RawMessage
		if err := json.Unmarshal(rawModules[i]["zones"], &zoneFields); err != nil {
			t.Errorf("module %q zones=%s: %v", module.Name, rawModules[i]["zones"], err)
		}
		for _, key := range []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"} {
			if _, ok := zoneFields[key]; !ok {
				t.Errorf("module %q missing resolved %s zone", module.Name, key)
			}
		}
		for _, key := range []string{"protected", "append_only"} {
			var paths []json.RawMessage
			if err := json.Unmarshal(zoneFields[key], &paths); err != nil || paths == nil {
				t.Errorf("module %q %s zone=%s; want array", module.Name, key, zoneFields[key])
			}
		}
		got[module.Name] = true
		switch module.Name {
		case "alpha":
			if module.IDPrefix != "ALP" || module.Dir != "packages/alpha" || len(module.DependsOn) != 0 || module.Zones.Evals == nil || len(module.Zones.Evals) != 0 {
				t.Errorf("alpha fields/zones = %+v", module)
			}
		case "beta":
			if module.IDPrefix != "BET" || module.Dir != "packages/beta" || !reflect.DeepEqual(module.DependsOn, []string{"alpha"}) || !reflect.DeepEqual(module.Zones.Evals, []string{"evals/**"}) {
				t.Errorf("beta fields/zones = %+v", module)
			}
		case "gamma":
			if module.IDPrefix != "GAM" || module.Dir != "packages/gamma" || len(module.DependsOn) != 0 {
				t.Errorf("gamma fields = %+v", module)
			}
		default:
			t.Errorf("unexpected module %+v", module)
		}
		if !reflect.DeepEqual(module.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(module.Zones.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(module.Zones.Preserve, []string{"data/**"}) {
			t.Errorf("unresolved default zones for %s: %+v", module.Name, module.Zones)
		}
	}
	if len(modules) != len(names) {
		t.Errorf("modules=%d; want %d", len(modules), len(names))
	}
	for _, name := range names {
		if !got[name] {
			t.Errorf("missing module %q: %s", name, envelope["data"])
		}
	}
}

func Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor(t *testing.T) {
func Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, true)
	before := cliGit(t, root, "rev-parse", "HEAD")
	status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), "start", "manual")
	if status != 0 || stderr != "" || !strings.Contains(plain, "# Write evals") || !strings.Contains(plain, "ALP-002") {
		t.Fatalf("start: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	if after := cliGit(t, root, "rev-parse", "HEAD"); after != before {
		t.Errorf("start committed: %s -> %s", before, after)
	}
	state, err := workflow.Current(root, "manual")
	if err != nil || state.Phase != "evals" || state.Run == "" {
		t.Fatalf("start state=%#v, %v", state, err)
	}
	change, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(change), "run:") || !strings.Contains(string(change), "phase: evals") || !strings.Contains(string(change), "### ALP-002") || strings.Contains(string(change), "new:extra") {
		t.Errorf("assigned change=%q, %v", change, err)
	}
	expected, err := workflow.Context(root, "manual", false)
	if err != nil || plain != expected {
		t.Errorf("first-step context differs from workflow.Context: %v", err)
	}
	for _, tc := range []struct {
		args []string
		want int
	}{
		{[]string{"start"}, 2},
		{[]string{"start", "manual", "extra"}, 2},
		{[]string{"start", "missing"}, 1},
		{[]string{"start", "manual"}, 1},
	} {
		args := append(append([]string{}, tc.args...), "--json")
		status, out, errOut := runLifecycleBinary(t, bin, root, args...)
		lifecycleCommandFailure(t, status, tc.want, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runLifecycleBinary(t, bin, bad, "start", "manual", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
}

func Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		root := lifecycleCLIProject(t, bin, false)
		cliGit(t, root, "checkout", "-b", "feature")
		cliGit(t, root, "checkout", "main")
		cliGit(t, root, "rm", "changes/0001-manual.md")
		cliCommit(t, root, "base without queued change")
		cliGit(t, root, "update-ref", "refs/remotes/origin/main", cliGit(t, root, "rev-parse", "HEAD"))
		cliGit(t, root, "checkout", "feature")
		args := []string{"start", "manual"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, output, stderr := runLifecycleBinary(t, bin, root, args...)
		if status != 0 {
			t.Fatalf("start with approval warning (JSON=%v): %d %q %q", jsonMode, status, output, stderr)
		}
		if jsonMode {
			if stderr != "" || !strings.Contains(cliString(t, output), "# Write evals") {
				t.Errorf("JSON start omitted context or wrote stderr: %q %q", output, stderr)
			}
			var diagnostics []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
				Warning bool   `json:"warning"`
			}
			if err := json.Unmarshal(decodeEnvelope(t, output)["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 || diagnostics[0].Code != "P04" || diagnostics[0].Message == "" || !diagnostics[0].Warning {
				t.Errorf("approval warning missing: %s, %v", output, err)
			}
		} else if !strings.Contains(output, "# Write evals") || !strings.Contains(output+stderr, "P04") {
			t.Errorf("plain start omitted context or P04 warning: %q %q", output, stderr)
		}
	}
}

func Test_APP_004_HeldModulesUseApprovedProjection(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "projected report")
	cliHold(t, root)
	changePath := filepath.Join(root, "changes/0001-manual.md")
	change, err := os.ReadFile(changePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, root, "changes/0001-manual.md", strings.Replace(string(change), "interface: [api.go, extra.go]", "interface: [api.go, unapproved.go]", 1))
	var status int
	var stderr string
	status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), "modules")
	if status != 0 || stderr != "" || !strings.Contains(plain, "extra.go") || strings.Contains(plain, "unapproved.go") {
		t.Errorf("held plain graph: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := runLifecycleBinary(t, bin, root, "modules", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("held JSON graph: exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	envelope := decodeEnvelope(t, output)
	var modules []struct {
		Name  string `json:"name"`
		Zones struct {
			Interface []string `json:"interface"`
		} `json:"zones"`
	}
	if err := json.Unmarshal(envelope["data"], &modules); err != nil || len(modules) != 2 || modules[0].Name != "alpha" || !strings.Contains(strings.Join(modules[0].Zones.Interface, ","), "extra.go") || strings.Contains(strings.Join(modules[0].Zones.Interface, ","), "unapproved.go") {
		t.Errorf("held JSON graph=%s, %v", envelope["data"], err)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) != 0 {
		t.Errorf("held diagnostics=%s, %v", envelope["diagnostics"], err)
	}
}

func Test_APP_011_APP_032_APP_033_ContextCompactFullAndErrors(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	appPut(t, root, "packages/alpha/internal/context-note.txt", "Full context marker.\n")
	cliCommit(t, root, "context content")
	cliHold(t, root)
	for _, full := range []bool{false, true} {
		args := []string{"context", "manual"}
		if full {
			args = append(args, "--full")
		}
		status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), args...)
		if status != 0 || stderr != "" {
			t.Fatalf("context %v: %d %q", full, status, stderr)
		}
		expected, err := workflow.Context(root, "manual", full)
		if err != nil || plain != expected {
			t.Errorf("context %v differs from public workflow: %v", full, err)
		}
		status, output, stderr := runLifecycleBinary(t, bin, root, append(args, "--json")...)
		if status != 0 || stderr != "" || cliString(t, output) != plain {
			t.Errorf("context JSON %v: %d %q %q", full, status, output, stderr)
		}
		var diagnostics []json.RawMessage
		diagnosticJSON := decodeEnvelope(t, output)["diagnostics"]
		if err := json.Unmarshal(diagnosticJSON, &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
			t.Errorf("context diagnostics=%s; want empty array: %v", diagnosticJSON, err)
		}
		if full && !strings.Contains(plain, "Full context marker.") || !full && strings.Contains(plain, "Full context marker.") {
			t.Errorf("context full selection: %v", full)
		}
	}
	for _, tc := range []struct {
		args []string
		want int
	}{
		{[]string{"context"}, 2}, {[]string{"context", "manual", "extra"}, 2}, {[]string{"context", "missing"}, 1},
	} {
		status, out, errOut := runLifecycleBinary(t, bin, root, append(tc.args, "--json")...)
		lifecycleCommandFailure(t, status, tc.want, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runLifecycleBinary(t, bin, bad, "context", "manual", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
}

func Test_APP_012_APP_034_APP_035_APP_036_CoverageCurrentProjectionAndErrors(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, false)
	status, plain, stderr := runLifecycleBinary(t, bin, root, "coverage", "alpha")
	if status != 0 || stderr != "" || !strings.Contains(plain, "ALP-001") {
		t.Errorf("plain current coverage: %d %q %q", status, plain, stderr)
	}
	status, output, stderr := runLifecycleBinary(t, bin, root, "coverage", "alpha", "--json")
	if status != 0 || stderr != "" || !strings.Contains(output, "ALP-001") {
		t.Errorf("current coverage: %d %q %q", status, output, stderr)
	}
	cliJSONResult(t, output)
	status, out, errOut := runLifecycleBinary(t, bin, root, "coverage", "alpha", "--with-change", "manual", "--json")
	lifecycleCommandFailure(t, status, 1, out, errOut, true) // projected ALP-002 is uncovered
	appPut(t, root, ".fireweed/emit-new", "yes\n")
	cliCommit(t, root, "projected JUnit case")
	cliHold(t, root)
	status, out, errOut = runLifecycleBinary(t, bin, root, "coverage", "alpha", "--with-change", "manual", "--json")
	lifecycleCommandFailure(t, status, 1, out, errOut, true) // mapped ALP-002 fails before implementation
	appPut(t, root, "packages/alpha/internal/ready", "ready\n")
	status, output, stderr = runLifecycleBinary(t, bin, root, "coverage", "alpha", "--with-change", "manual", "--json")
	if status != 0 || stderr != "" || !strings.Contains(output, "ALP-002") {
		t.Errorf("projected coverage: %d %q %q", status, output, stderr)
	}
	cliJSONResult(t, output)
	held := lifecycleCLIProject(t, bin, false)
	appPut(t, held, ".fireweed/emit-new", "yes\n")
	appPut(t, held, ".fireweed/no-report.sh", "#!/bin/sh\nexit 7\n")
	cliCommit(t, held, "projected JUnit case")
	cliHold(t, held)
	appPut(t, held, "packages/alpha/internal/ready", "ready\n")
	changePath := filepath.Join(held, "changes/0001-manual.md")
	change, err := os.ReadFile(changePath)
	if err != nil {
		t.Fatal(err)
	}
	appPut(t, held, "changes/0001-manual.md", strings.Replace(string(change), "  interface: [api.go, extra.go]\n", "  interface: [api.go, extra.go]\n  hooks:\n    eval: [sh, .fireweed/no-report.sh]\n", 1))
	if _, err := project.Project(held, "manual"); err != nil {
		t.Fatalf("live hook override must be a valid projection: %v", err)
	}
	status, output, stderr = runLifecycleBinary(t, bin, held, "coverage", "alpha", "--with-change", "manual", "--json")
	if status != 0 || stderr != "" || !strings.Contains(output, "ALP-002") {
		t.Errorf("held approved coverage: %d %q %q", status, output, stderr)
	}
	cliJSONResult(t, output)
	for _, args := range [][]string{{"coverage"}, {"coverage", "missing"}, {"coverage", "alpha", "--with-change", "missing"}, {"coverage", "alpha", "extra"}} {
		status, out, errOut := runLifecycleBinary(t, bin, root, append(args, "--json")...)
		lifecycleCommandFailure(t, status, 2, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut = runLifecycleBinary(t, bin, bad, "coverage", "alpha", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
}

func Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks(t *testing.T) {
// Protected, project-owned Go hook. No third-party dependencies.
package main
import (
 "bytes"
 "encoding/json"
 "encoding/xml"
 "fmt"
 "go/parser"
 "go/token"
 "io/fs"
 "os"
 "os/exec"
 "path/filepath"
 "sort"
 "strconv"
 "strings"
)
type Case struct { Name string `xml:"name,attr"`; Class string `xml:"classname,attr"`; Failure *string `xml:"failure,omitempty"`; Skipped *string `xml:"skipped,omitempty"` }
type Suite struct { XMLName xml.Name `xml:"testsuite"`; Cases []Case `xml:"testcase"` }
func main() {
 if len(os.Args)<2 { os.Exit(2) }
 if os.Args[1]=="eval" {
  cmd:=exec.Command("go","test","-json","./"+os.Args[2]+"/..."); var output bytes.Buffer; cmd.Stdout=&output; cmd.Stderr=os.Stderr; err:=cmd.Run()
  dec:=json.NewDecoder(&output); suite:=Suite{}; indexes:=map[string]int{}; logs:=map[string]string{}; finished:=map[string]bool{}; failedPkgs:=map[string]bool{}; pkgHasFailure:=map[string]bool{}
  for { var ev struct { Action,Package,Test,Output string }; if dec.Decode(&ev)!=nil { break }; if ev.Test=="" { if ev.Output!="" { fmt.Print(ev.Output) }; if ev.Action=="fail" { failedPkgs[ev.Package]=true }; continue }; key:=ev.Package+"/"+ev.Test
   if ev.Action=="run" { indexes[key]=len(suite.Cases);suite.Cases=append(suite.Cases,Case{Name:ev.Test,Class:ev.Package}) }
   logs[key]+=ev.Output
   if ev.Action=="pass"||ev.Action=="fail"||ev.Action=="skip" { finished[key]=true }
   if ev.Action=="fail" { if i,ok:=indexes[key];ok { text:=logs[key];suite.Cases[i].Failure=&text;pkgHasFailure[ev.Package]=true } }
   if ev.Action=="skip" { if i,ok:=indexes[key];ok { text:=logs[key];suite.Cases[i].Skipped=&text } }
  }
  // A test that started but never finished (package timeout, crash) is failed evidence, never a pass.
  for key,i:=range indexes { if !finished[key] { text:=logs[key]+"\nfireweed: test did not finish (timeout, crash or interrupted run)\n";suite.Cases[i].Failure=&text;pkgHasFailure[suite.Cases[i].Class]=true } }
  // A package that failed without any failing test (build or setup failure) gets an unmapped failing case.
  pkgs:=[]string{};for pkg:=range failedPkgs { if !pkgHasFailure[pkg] { pkgs=append(pkgs,pkg) } };sort.Strings(pkgs)
  for _,pkg:=range pkgs { text:="fireweed: package failed without a failing test (build, setup or collection failure)\n";suite.Cases=append(suite.Cases,Case{Name:"PackageFailure",Class:pkg,Failure:&text}) }
  b,e:=xml.Marshal(suite);if e!=nil {panic(e)};if e=os.WriteFile(os.Args[3],b,0644);e!=nil{panic(e)};if err!=nil {os.Exit(1)};return
 }
 bin:=os.Getenv("FIREWEED_BIN");if bin==""{bin="fireweed"};cmd:=exec.Command(bin,"modules","--json");data,err:=cmd.Output();if err!=nil{panic(err)}
 var envelope struct { Data []struct { Name,Dir string; DependsOn []string `json:"depends_on"` } };if err=json.Unmarshal(data,&envelope);err!=nil{panic(err)}
 mod,err:=os.ReadFile("go.mod");if err!=nil{panic(err)};modulePath:="";for _,line:=range strings.Split(string(mod),"\n"){f:=strings.Fields(line);if len(f)==2&&f[0]=="module"{modulePath=strings.Trim(f[1],"\"")}}
 violations:=0
 for _,m:=range envelope.Data { err=filepath.WalkDir(m.Dir,func(file string,d fs.DirEntry,err error)error{
  if err!=nil{return err};if d.IsDir()||!strings.HasSuffix(file,".go"){return nil};tree,err:=parser.ParseFile(token.NewFileSet(),file,nil,parser.ImportsOnly);if err!=nil{return err}
  for _,imp:=range tree.Imports { imported,_:=strconv.Unquote(imp.Path.Value);for _,target:=range envelope.Data { base:=modulePath+"/"+target.Dir;if imported!=base&&!strings.HasPrefix(imported,base+"/"){continue};allowed:=m.Name==target.Name;for _,dep:=range m.DependsOn{if dep==target.Name{allowed=true}}
   if !allowed || (strings.Contains(filepath.ToSlash(file),"/evals/")&&imported!=base) || (m.Name!=target.Name&&imported!=base) {fmt.Printf("%s: %s imports %s outside its declared interface/dependencies\n",m.Name,file,imported);violations++}
  }};return nil });if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(2)} }
 if violations>0{os.Exit(1)}
}
499:### APP-129
503:### APP-130
508:### APP-131
511:### APP-132
516:### APP-133
521:### APP-134
524:### APP-135
529:### APP-136
532:### APP-137
535:### APP-138
538:### APP-139
541:### APP-140
544:### APP-141
547:### APP-142
550:### APP-143
554:### APP-144
558:### APP-145
562:### APP-146
566:### APP-147
570:### APP-148
574:### APP-149
578:### APP-150
581:### APP-151
584:### APP-152
587:### APP-153
590:### APP-154
593:### APP-155
599:[0074-doctor-setup-commands] Doctor CLI leaves request Timeout unset, using the backend's existing defaults;
603:[0112-new-module-command] Creation uses the first configured module glob ending in a single wildcard
609:[0127-show-history-command] History JSON uses lowercase keys and preserves project query section strings;

exec
/bin/zsh -lc "sed -n '1,225p' packages/project/internal/fixture/fixture.go; cat packages/project/internal/queue/check.go; sed -n '180,217p' packages/project/internal/queue/project.go; rg -n 'func (Current|state|existing)|CodedError|func .*Error' packages/workflow/internal/lifecycle/stub.go; cat packages/assets/internal/content/*go | head -170; cat packages/assets/internal/recipe/*go | head -235" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc "sed -n '116,245p' packages/app/evals/app_test.go; sed -n '145,220p' packages/app/evals/run_cli_test.go; sed -n '54,103p' packages/app/evals/show_history_cli_test.go; rg -n 'Contains|HasPrefix|EqualFold|"'!= "|corrective'"\\(' packages/app/evals/{check_fix_cli_test,show_stale_cli_test,status_cli_test,migrate_cli_test,doctor_setup_test,check_commits_cli_test}.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 565ms:
func Test_APP_003_InvalidRequests(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		bad  string
		code string
	}{
		{"unknown-command", []string{"not-a-command"}, "not-a-command", "U01"},
		{"unknown-option", []string{"help", "--bogus"}, "--bogus", "U01"},
		{"missing-argument", []string{"explain"}, "explain", "U01"},
		{"extra-argument", []string{"version", "surplus"}, "version", "U01"},
		{"new-change-arity", []string{"new-change", "alpha"}, "new-change", "U01"},
		{"stamp-arity", []string{"stamp"}, "stamp", "U01"},
		{"check-arity", []string{"check", "extra"}, "check", "U01"},
		{"show-arity", []string{"show"}, "show", "U01"},
		{"start-arity", []string{"start"}, "start", "U01"},
		{"context-arity", []string{"context"}, "context", "U01"},
		{"coverage-arity", []string{"coverage"}, "coverage", "U01"},
		{"verify-arity", []string{"verify", "manual", "extra"}, "verify", "U01"},
		{"APP_013_unknown-explanation", []string{"explain", "NO-SUCH-CODE"}, "NO-SUCH-CODE", "U02"},
		{"APP_013_unknown-playbook", []string{"playbook", "no-such-playbook"}, "no-such-playbook", "U02"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, format := range []struct {
				name string
				json bool
			}{{"plain", false}, {"JSON", true}} {
				args := append([]string(nil), tc.args...)
				if format.json {
					args = append(args, "--json")
				}
				status, stdout, stderr := run(args...)
				if status != 2 {
					t.Errorf("%s status=%d; want 2", format.name, status)
				}
				diagnostic := stdout + stderr
				if format.json {
					diagnostic = assertFailedJSON(t, stdout, tc.code)
				} else if !strings.HasPrefix(strings.TrimSpace(diagnostic), tc.code+":") {
					t.Errorf("%s diagnostic %q must have code %s", format.name, diagnostic, tc.code)
				}
				if !strings.Contains(diagnostic, tc.bad) || !corrective(diagnostic) {
					t.Errorf("%s diagnostic %q must identify %q and suggest a correction", format.name, diagnostic, tc.bad)
				}
			}
		})
	}
}

func Test_APP_014_MaintenanceCommandsDispatch(t *testing.T) {
	bin := buildAppBinary(t)
	root := migrateProject(t)
	for _, command := range []string{"gc", "migrate"} {
		t.Run(command, func(t *testing.T) {
			if !strings.Contains(assets.Help(), command) {
				t.Fatalf("%q is not listed by help", command)
			}
			status, plain, stderr := runAppBinary(t, bin, root, command)
			if status != 0 || stderr != "" || strings.TrimSpace(plain) == "" {
				t.Fatalf("%s plain: status=%d stdout=%q stderr=%q", command, status, plain, stderr)
			}
			if regexp.MustCompile(`(?i)unimplemented|not implemented|under construction`).MatchString(plain) {
				t.Errorf("%s still reported a placeholder: %q", command, plain)
			}
			status, output, stderr := runAppBinary(t, bin, root, command, "--json")
			if status != 0 || stderr != "" {
				t.Fatalf("%s JSON: status=%d stdout=%q stderr=%q", command, status, output, stderr)
			}
			env := decodeEnvelope(t, output)
			if string(env["data"]) == "null" {
				t.Errorf("%s JSON data is null", command)
			}
			var diagnostics []json.RawMessage
			if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
				t.Errorf("%s JSON diagnostics=%s: %v", command, env["diagnostics"], err)
			}
			for _, args := range [][]string{{command, "extra"}, {command, "--bogus"}} {
				for _, jsonMode := range []bool{false, true} {
					call := append([]string(nil), args...)
					if jsonMode {
						call = append(call, "--json")
					}
					status, stdout, errOut := runAppBinary(t, bin, root, call...)
					message := commandFailure(t, status, 2, stdout, errOut, jsonMode)
					if jsonMode {
						message = assertFailedJSON(t, stdout, "U01")
					} else if !strings.HasPrefix(strings.TrimSpace(errOut), "U01:") {
						t.Errorf("%s invalid plain diagnostic=%q; want U01", command, errOut)
					}
					if !corrective(message) {
						t.Errorf("%s invalid diagnostic lacks guidance: %q", command, message)
					}
				}
			}
		})
	}
}

func Test_APP_015_FailedJSON(t *testing.T) {
	status, stdout, _ := run("init", "--json")
	if status == 0 {
		t.Error("failed request returned exit 0")
	}
	assertFailedJSON(t, stdout)
}

func assertFailedJSON(t *testing.T, stdout string, wantCode ...string) string {
	t.Helper()
	envelope := decodeEnvelope(t, stdout)
	if string(envelope["data"]) != "null" {
		t.Errorf("data=%s; want null", envelope["data"])
	}
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
		t.Fatalf("diagnostics=%s; want nonempty array: %v", envelope["diagnostics"], err)
	}
	if strings.TrimSpace(diagnostics[0].Code) == "" || strings.TrimSpace(diagnostics[0].Message) == "" {
		t.Errorf("diagnostic needs code and message: %+v", diagnostics[0])
	}
	if len(wantCode) > 0 && diagnostics[0].Code != wantCode[0] {
		t.Errorf("diagnostic code=%q; want %q", diagnostics[0].Code, wantCode[0])
	}
	return diagnostics[0].Message
}

// The app-level convention probe uses one PATH recorder for external commands.
func Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents(t *testing.T) {
		t.Fatalf("data=%s: %v", env["data"], err)
	}
	var applied []map[string]json.RawMessage
	if err := json.Unmarshal(data["applied"], &applied); err != nil || applied == nil {
		t.Fatalf("applied=%s: %v", data["applied"], err)
	}
	return data, applied
}

func runCLIFlag(t *testing.T, data map[string]json.RawMessage, key string, want bool) {
	t.Helper()
	var got bool
	if err := json.Unmarshal(data[key], &got); err != nil || got != want {
		t.Errorf("%s=%s; want %v: %v", key, data[key], want, err)
	}
}

func runCLISummaries(t *testing.T, applied []map[string]json.RawMessage, names ...string) {
	t.Helper()
	if len(applied) != len(names) {
		t.Fatalf("applied=%v; want %v", applied, names)
	}
	for i, name := range names {
		var change, run, phase string
		for field, target := range map[string]*string{"change": &change, "run": &run, "phase": &phase} {
			if err := json.Unmarshal(applied[i][field], target); err != nil {
				t.Errorf("applied[%d].%s=%s: %v", i, field, applied[i][field], err)
			}
		}
		if !strings.Contains(change, name) || run == "" || phase != "applied" {
			t.Errorf("applied[%d] = change %q run %q phase %q", i, change, run, phase)
		}
	}
}

func runCLISuccess(t *testing.T, status int, out, errOut string, jsonMode bool) {
	t.Helper()
	if status != 0 || (jsonMode && errOut != "") {
		t.Fatalf("run exit=%d stdout=%q stderr=%q", status, out, errOut)
	}
}

func runCLIPlain(t *testing.T, out string, count int, name, state string) {
	t.Helper()
	lower := strings.ToLower(out)
	if name != "" && !strings.Contains(lower, name) {
		t.Errorf("plain run lacks %s: %q", name, out)
	}
	counted := regexp.MustCompile(fmt.Sprintf(`(?i)(?:applied|count)\D*%d\b|\b%d\s+changes?\b`, count, count))
	if !counted.MatchString(out) {
		t.Errorf("plain run lacks applied count %d: %q", count, out)
	}
	if state == "empty" && !strings.Contains(lower, "empty") {
		t.Errorf("plain run lacks empty queue: %q", out)
	}
	if state == "limited" && !strings.Contains(lower, "limit") && !strings.Contains(lower, "remain") && !strings.Contains(lower, "queued") {
		t.Errorf("plain run lacks queued work: %q", out)
	}
}

func runCLIReport(t *testing.T, f runCLISetup, name, prefix, phase, want string) {
	t.Helper()
	paths, err := filepath.Glob(filepath.Join(f.root, ".fireweed/runs/*/", phase, name+".json"))
	if err != nil || len(paths) != 1 {
		t.Fatalf("%s %s reports=%v: %v", name, phase, paths, err)
	}
	var report checks.Report
	if err := json.Unmarshal(mustReadNextCLI(t, paths[0]), &report); err != nil {
		t.Fatal(err)
	}
	for _, c := range report.Cases {
		if c.Result == want {
			for _, id := range c.IDs {
				if id == prefix+"-002" {
					return
				}
func assertAppHistoryJSON(t *testing.T, output string, entries []project.SpecHistoryEntry) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
	}
	var rows []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &rows); err != nil || rows == nil || len(rows) != len(entries) {
		t.Fatalf("data=%s; want %d history rows: %v", envelope["data"], len(entries), err)
	}
	for i, row := range rows {
		want := entries[i]
		for key := range row {
			if key != strings.ToLower(key) {
				t.Errorf("row %d has non-lowercase key %q", i, key)
			}
		}
		for key, value := range map[string]string{
			"path": want.Path, "module": want.Module, "motivation": want.Motivation,
			"decisions": want.Decisions, "removed_decisions": want.RemovedDecisions,
			"provenance": want.Provenance,
		} {
			var got string
			if err := json.Unmarshal(row[key], &got); err != nil || got != value {
				t.Errorf("row %d %s=%s; want %q: %v", i, key, row[key], value, err)
			}
		}
		var gotIDs []string
		if err := json.Unmarshal(row["ids"], &gotIDs); err != nil || !reflect.DeepEqual(gotIDs, want.IDs) {
			t.Errorf("row %d ids=%s; want %v: %v", i, row["ids"], want.IDs, err)
		}
		var date string
		if err := json.Unmarshal(row["implementation_date"], &date); err != nil {
			t.Errorf("row %d implementation_date=%s; want string: %v", i, row["implementation_date"], err)
		}
		if want.ImplementationDate.IsZero() {
			if date != "" {
				t.Errorf("row %d undated implementation_date=%q; want empty", i, date)
			}
		} else {
			parsed, err := time.Parse(time.RFC3339, date)
			if err != nil || !parsed.Equal(want.ImplementationDate) {
				t.Errorf("row %d implementation_date=%q; want RFC3339 %s: %v", i, date, want.ImplementationDate.Format(time.RFC3339), err)
			}
		}
	}
}

func Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter(t *testing.T) {
packages/app/evals/check_fix_cli_test.go:35:		if module.Name != "alpha" {
packages/app/evals/check_fix_cli_test.go:77:		if stderr != "" {
packages/app/evals/check_fix_cli_test.go:86:		if stderr != "" {
packages/app/evals/check_fix_cli_test.go:90:		if !corrective(message) {
packages/app/evals/check_fix_cli_test.go:95:	if out != "" {
packages/app/evals/check_fix_cli_test.go:98:	if wantCode != "" && !strings.Contains(stderr, wantCode+":") || !corrective(stderr) {
packages/app/evals/check_fix_cli_test.go:120:	if !strings.Contains(plain, "stale base for ALP-001") || strings.Contains(plain, "missing base for") {
packages/app/evals/check_fix_cli_test.go:132:	if !reflect.DeepEqual(bases, want) || doc.Frontmatter["custom"] != "keep-first" || doc.Body != originalFirst.Body {
packages/app/evals/check_fix_cli_test.go:140:	if !reflect.DeepEqual(secondBases, map[string]any{"ALP-001": secondOne, "ALP-003": secondThree}) || secondDoc.Frontmatter["custom"] != "keep-second" || secondDoc.Body != originalSecond.Body {
packages/app/evals/check_fix_cli_test.go:151:	if !strings.Contains(message, "stale base for ALP-001") || !bytes.Equal(afterFirst, staleCLIText(t, root, first)) || !bytes.Equal(afterSecond, staleCLIText(t, root, second)) {
packages/app/evals/check_fix_cli_test.go:163:	if !strings.Contains(plain, "Check passed") {
packages/app/evals/check_fix_cli_test.go:167:	if !reflect.DeepEqual(bases, map[string]any{"ALP-001": checkFixHash(t, root, "", "ALP-001"), "ALP-002": checkFixHash(t, root, "", "ALP-002")}) || doc.Frontmatter["custom"] != "retain" || !strings.Contains(doc.Body, "Keep human text.") {
packages/app/evals/status_cli_test.go:261:			if field != "" && !strings.Contains(plain, field) {
packages/app/evals/status_cli_test.go:268:			if !strings.Contains(plain, field) {
packages/app/evals/status_cli_test.go:273:			if !strings.Contains(plain, dep) {
packages/app/evals/status_cli_test.go:279:				if !strings.Contains(plain, zone) {
packages/app/evals/status_cli_test.go:287:			if !strings.Contains(plain, field) {
packages/app/evals/status_cli_test.go:293:		if !strings.Contains(plain, o.Module) || !strings.Contains(plain, o.Reason) {
packages/app/evals/status_cli_test.go:297:	if !strings.Contains(plain, strconv.Itoa(report.Unapproved)) {
packages/app/evals/status_cli_test.go:301:		if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
packages/app/evals/status_cli_test.go:306:		if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
packages/app/evals/status_cli_test.go:312:			if !strings.Contains(plain, field) {
packages/app/evals/status_cli_test.go:317:			if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
packages/app/evals/status_cli_test.go:346:		if !strings.Contains(plain+stderr, d.Code) || !strings.Contains(plain+stderr, d.Message) {
packages/app/evals/status_cli_test.go:351:	if status != 0 || stderr != "" {
packages/app/evals/status_cli_test.go:418:		} else if d.Path != "" {
packages/app/evals/status_cli_test.go:443:	if status != 0 || stderr != "" || !strings.Contains(strings.ToLower(plain), "empty") {
packages/app/evals/status_cli_test.go:447:	if status != 0 || stderr != "" {
packages/app/evals/status_cli_test.go:455:		if string(data[key]) != "[]" {
packages/app/evals/status_cli_test.go:460:		if string(data[key]) == "null" || !strings.HasPrefix(string(data[key]), "{") {
packages/app/evals/status_cli_test.go:474:		if status != wantExit || !strings.Contains(out+errOut, code+":") || !strings.Contains(out+errOut, cause) || !corrective(out+errOut) {
packages/app/evals/status_cli_test.go:478:		if status != wantExit || errOut != "" {
packages/app/evals/status_cli_test.go:482:		if !strings.Contains(message, cause) || !corrective(message) {
packages/app/evals/check_commits_cli_test.go:82:		if err != nil || !reflect.DeepEqual(report.Commits, tc.want) || len(report.Diagnostics) != 0 || len(report.Squashed) != 1 || report.Squashed[0].Commit != specOID || report.Squashed[0].Change != "changes/archive/0001-queued.md" {
packages/app/evals/check_commits_cli_test.go:89:	if status != 0 || stderr != "" || !regexp.MustCompile(`\b3\b`).MatchString(plain) {
packages/app/evals/check_commits_cli_test.go:93:	if status != 0 || stderr != "" || !regexp.MustCompile(`\b2\b`).MatchString(plain) {
packages/app/evals/check_commits_cli_test.go:97:	if status != 0 || stderr != "" {
packages/app/evals/check_commits_cli_test.go:105:	if status != 0 || stderr != "" {
packages/app/evals/check_commits_cli_test.go:131:	if status != 0 || stderr != "" {
packages/app/evals/check_commits_cli_test.go:139:	if status != 0 || stderr != "" || !regexp.MustCompile(`\b1\b`).MatchString(plain) || strings.Contains(strings.ToLower(plain), "warning") {
packages/app/evals/check_commits_cli_test.go:143:		if !strings.Contains(plain, oid) || !strings.Contains(plain, path) {
packages/app/evals/check_commits_cli_test.go:148:	if status != 0 || stderr != "" {
packages/app/evals/check_commits_cli_test.go:182:		if want[d.Code] != d.Path || !strings.Contains(d.Message, oid) {
packages/app/evals/check_commits_cli_test.go:194:			if strings.Contains(line, code) && strings.Contains(line, path) && strings.Contains(line, oid) && corrective(line) {
packages/app/evals/check_commits_cli_test.go:203:	if status != 1 || stderr != "" {
packages/app/evals/check_commits_cli_test.go:207:	if string(envelope["data"]) != "null" {
packages/app/evals/check_commits_cli_test.go:220:			if d.Code == code && d.Path == path && strings.Contains(d.Message, oid) && corrective(d.Message) {
packages/app/evals/check_commits_cli_test.go:245:			message := commandFailure(t, status, tc.status, output, stderr, strings.Contains(strings.Join(tc.args, " "), "--json"))
packages/app/evals/check_commits_cli_test.go:246:			if strings.Contains(strings.Join(tc.args, " "), "--json") {
packages/app/evals/check_commits_cli_test.go:248:			} else if !strings.Contains(output+stderr, tc.code+":") {
packages/app/evals/check_commits_cli_test.go:251:			if !corrective(message) {
packages/app/evals/check_commits_cli_test.go:260:	if !corrective(message) {
packages/app/evals/check_commits_cli_test.go:273:		if !(coded && code == "C03" || strings.HasPrefix(err.Error(), "C03:")) || !strings.Contains(err.Error(), "packages/alpha/SPEC.md") {
packages/app/evals/check_commits_cli_test.go:280:	if !corrective(message) {
packages/app/evals/migrate_cli_test.go:118:		if status != 0 || stderr != "" || plain != migrateConfirmation {
packages/app/evals/migrate_cli_test.go:122:		if status != 0 || stderr != "" {
packages/app/evals/migrate_cli_test.go:141:	if status != 0 || stderr != "" || plain != migrateConfirmation {
packages/app/evals/migrate_cli_test.go:145:	if status != 0 || stderr != "" || commandStringData(t, output) != plain {
packages/app/evals/migrate_cli_test.go:195:					if stderr != "" {
packages/app/evals/migrate_cli_test.go:198:					if string(decodeEnvelope(t, out)["data"]) != "null" {
packages/app/evals/migrate_cli_test.go:206:						if i >= len(got) || got[i].Code != want.Code || !strings.Contains(got[i].Message, want.Message) || got[i].Warning != want.Warning || got[i].Path != want.Path {
packages/app/evals/migrate_cli_test.go:211:						if !corrective(d.Message) {
packages/app/evals/migrate_cli_test.go:215:				} else if out != "" || !strings.Contains(stderr, tc.code+":") || !corrective(stderr) {
packages/app/evals/migrate_cli_test.go:231:		} else if !strings.HasPrefix(strings.TrimSpace(stderr), "C01:") {
packages/app/evals/migrate_cli_test.go:234:		if !corrective(message) {
packages/app/evals/migrate_cli_test.go:254:			} else if out != "" || !strings.HasPrefix(strings.TrimSpace(stderr), "U01:") {
packages/app/evals/migrate_cli_test.go:257:			if !corrective(message) {
packages/app/evals/show_stale_cli_test.go:38:	if status != 0 || stderr != "" {
packages/app/evals/show_stale_cli_test.go:42:	if status != 0 || stderr != "" || commandStringData(t, output) != plain {
packages/app/evals/show_stale_cli_test.go:100:		if !strings.Contains(recovered, marker) {
packages/app/evals/show_stale_cli_test.go:104:	if strings.Contains(recovered, "Target MUST propose different text") || strings.Contains(recovered, "deadbeef") {
packages/app/evals/show_stale_cli_test.go:108:		if !strings.Contains(strings.ToLower(fallback), strings.ToLower(marker)) {
packages/app/evals/show_stale_cli_test.go:112:	if strings.Contains(fallback, "Two MUST use old text") || strings.Contains(fallback, "-Two MUST") || strings.Contains(fallback, "+Two MUST") {
packages/app/evals/show_stale_cli_test.go:131:		if !strings.Contains(strings.ToLower(plain), strings.ToLower(marker)) {
packages/app/evals/show_stale_cli_test.go:135:	if strings.Contains(plain, "-One MUST") || strings.Contains(plain, "+One MUST") || strings.Contains(plain, "Proposed MUST differ") {
packages/app/evals/show_stale_cli_test.go:152:	if !strings.Contains(plain, "innerhash") || !strings.Contains(plain, "Inner MUST work.") || strings.Contains(plain, "outerhash") || strings.Contains(plain, "One MUST work.") {
packages/app/evals/show_stale_cli_test.go:169:			if jsonMode && !strings.Contains(strings.Join(call, " "), "--json") {
packages/app/evals/show_stale_cli_test.go:173:			assertAppHistoryFailure(t, status, 2, out, stderr, strings.Contains(strings.Join(call, " "), "--json"), "U01", "show --stale")
packages/app/evals/doctor_setup_test.go:65:	if status != 0 || errOut != "" {
packages/app/evals/doctor_setup_test.go:124:		if strings.Contains(line, name) && strings.Contains(strings.ToLower(line), outcome) {
packages/app/evals/doctor_setup_test.go:145:	if !strings.Contains(probeLog, "setup:") || !strings.Contains(probeLog, bin) || strings.Contains(probeLog, "worker:") || strings.Contains(probeLog, "audit:") {
packages/app/evals/doctor_setup_test.go:149:	if status != 0 || errOut != "" {
packages/app/evals/doctor_setup_test.go:180:	if strings.Contains(doctorCLIFile(t, log), "worker:") || strings.Contains(doctorCLIFile(t, log), "audit:") {
packages/app/evals/doctor_setup_test.go:187:	if status != 0 || !strings.Contains(plain+errOut, "C16") || !strings.Contains(strings.ToLower(plain+errOut), "warn") || !strings.Contains(plain+errOut, "AGENTS.md") {
packages/app/evals/doctor_setup_test.go:191:	if status != 0 || errOut != "" {
packages/app/evals/doctor_setup_test.go:197:		if string(d["code"]) == `"C16"` && string(d["warning"]) == "true" && strings.Contains(string(d["message"]), "AGENTS.md") {
packages/app/evals/doctor_setup_test.go:207:	if status != 0 || !strings.Contains(plain+errOut, "C16") {
packages/app/evals/doctor_setup_test.go:210:	if got := doctorCLIFile(t, log); !strings.Contains(got, "worker:") || !strings.Contains(got, "audit:") {
packages/app/evals/doctor_setup_test.go:223:	if got := doctorCLIFile(t, innerLog); !strings.Contains(got, "setup:") || !strings.Contains(got, bin) {
packages/app/evals/doctor_setup_test.go:226:	if got := doctorCLIFile(t, outerLog); got != "" {
packages/app/evals/doctor_setup_test.go:238:	if status != 1 || !strings.Contains(plain+errOut, "C14") || !strings.Contains(plain+errOut, "named-setup-failure") || !doctorCLIPlainOutcome(plain+errOut, "setup", "fail") {
packages/app/evals/doctor_setup_test.go:242:	if status != 1 || errOut != "" {
packages/app/evals/doctor_setup_test.go:246:	if string(envelope["data"]) != "null" {
packages/app/evals/doctor_setup_test.go:253:	if string(diags[0]["code"]) != `"C14"` || !strings.Contains(string(diags[0]["message"]), "named-setup-failure") {
packages/app/evals/doctor_setup_test.go:259:	if status != 1 || !strings.Contains(plain+errOut, "C05") {
packages/app/evals/doctor_setup_test.go:263:	if status != 1 || errOut != "" {
packages/app/evals/doctor_setup_test.go:267:	if string(envelope["data"]) != "null" {
packages/app/evals/doctor_setup_test.go:273:	if string(diags[0]["code"]) != `"C05"` || !strings.Contains(string(diags[0]["path"]), "SPEC.md") || string(diags[0]["warning"]) != "false" || len(diags[0]["message"]) == 0 {
packages/app/evals/doctor_setup_test.go:287:			if status != 2 || !strings.Contains(stdout+stderr, "U01") || !corrective(stdout+stderr) {
packages/app/evals/doctor_setup_test.go:291:			if status != 2 || stderr != "" {
packages/app/evals/doctor_setup_test.go:299:	if status != 2 || !strings.Contains(plain+errOut, "C01") {
packages/app/evals/doctor_setup_test.go:303:	if status != 2 || errOut != "" {
packages/app/evals/doctor_setup_test.go:311:	if status != 2 || !strings.Contains(plain+errOut, "C01") {
packages/app/evals/doctor_setup_test.go:329:	if agents != "Before.\n\n"+doctorCLIManagedBlock+"\n\nAfter.\n" {
packages/app/evals/doctor_setup_test.go:336:	if status != 0 || stderr != "" {
packages/app/evals/doctor_setup_test.go:354:	if status != 0 || stderr != "" {
packages/app/evals/doctor_setup_test.go:357:	if got := readInitFile(t, closer, "AGENTS.md"); got != "Closer before.\n"+doctorCLIManagedBlock+"\nCloser after.\n" {
packages/app/evals/doctor_setup_test.go:384:			if !strings.HasPrefix(got, tc.initial) || strings.Count(got, "<!-- fireweed:start -->") != 1 || strings.Count(got, "<!-- fireweed:end -->") != 1 || !strings.Contains(got, doctorCLIManagedBlock) {
packages/app/evals/doctor_setup_test.go:393:		if status != 2 || !strings.Contains(stdout+stderr, "P01") || !corrective(stdout+stderr) || readInitFile(t, root, "AGENTS.md") != malformed {
packages/app/evals/doctor_setup_test.go:397:		if status != 2 || stderr != "" {
packages/app/evals/doctor_setup_test.go:405:		if status != 2 || !strings.Contains(stdout+stderr, "U01") || !corrective(stdout+stderr) {
packages/app/evals/doctor_setup_test.go:409:		if status != 2 || errOut != "" {
packages/app/evals/doctor_setup_test.go:414:	if got := readInitFile(t, root, "AGENTS.md"); got != "Personal.\n<!-- fireweed:end -->\n" {
packages/app/evals/doctor_setup_test.go:424:	if status != 2 || !strings.Contains(stdout+stderr, "P01") || !corrective(stdout+stderr) {
packages/app/evals/doctor_setup_test.go:428:	if status != 2 || stderr != "" {
packages/app/evals/doctor_setup_test.go:438:	if status != 2 || !strings.Contains(stdout+stderr, "C01") {
packages/app/evals/doctor_setup_test.go:442:	if status != 2 || stderr != "" {
packages/app/evals/doctor_setup_test.go:450:	if status != 2 || !strings.Contains(stdout+stderr, "C01") {

 succeeded in 573ms:
package fixture

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/model"
	"gopkg.in/yaml.v3"
)

type loadFailure struct {
	code string
	path string
	err  error
}

func (e *loadFailure) Error() string { return e.err.Error() }
func (e *loadFailure) Unwrap() error { return e.err }

// LoadCode identifies the diagnostic cause retained by a Load error.
func LoadCode(err error) string {
	var failure *loadFailure
	if errors.As(err, &failure) {
		return failure.code
	}
	return "C01"
}

func loadCause(code string, err error) error { return &loadFailure{code: code, err: err} }

// LoadPath identifies the document responsible for a Load failure.
func LoadPath(err error) string {
	var failure *loadFailure
	if errors.As(err, &failure) && failure.path != "" {
		return failure.path
	}
	return "fireweed.yaml"
}

func moduleCause(rel string, err error) error {
	code, _, ok := contracts.ErrorCode(err)
	if !ok {
		code = "C03"
	}
	path := rel + "/SPEC.md"
	return &loadFailure{
		code: code,
		path: path,
		err:  fmt.Errorf("parse %s: %w. Use valid frontmatter and unique spec item headings, then retry", path, err),
	}
}

func Load(root string) (model.Snapshot, error) {
	data, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
	if err != nil {
		return model.Snapshot{}, fmt.Errorf("read fireweed.yaml: %w", err)
	}
	config, patterns, defaults, err := parseConfig(data)
	if err != nil {
		return model.Snapshot{}, err
	}

	snapshot := model.Snapshot{Root: root, Config: config}
	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() {
			return nil
		}
		if entry.Name() == ".git" {
			return filepath.SkipDir
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." || !matchesModule(patterns, rel) {
			return nil
		}
		data, err := os.ReadFile(filepath.Join(path, "SPEC.md"))
		if err != nil {
			return &loadFailure{code: "C03", path: rel + "/SPEC.md", err: fmt.Errorf("read %s/SPEC.md: %w", rel, err)}
		}
		module, err := parseModule(data, defaults, rel)
		if err != nil {
			return err
		}
		snapshot.Modules = append(snapshot.Modules, module)
		return nil
	})
	if err != nil {
		return model.Snapshot{}, err
	}
	return snapshot, nil
}

func parseConfig(data []byte) (map[string]any, []string, map[string]any, error) {
	var config map[string]any
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, nil, nil, fmt.Errorf("parse fireweed.yaml: %w", err)
	}
	if config == nil {
		return nil, nil, nil, fmt.Errorf("fireweed.yaml must be a mapping")
	}
	patterns, err := stringList(config["modules"], "fireweed.yaml modules")
	if err != nil {
		return nil, nil, nil, err
	}
	defaults := map[string]any{}
	if value, ok := config["module_defaults"]; ok {
		defaults, ok = value.(map[string]any)
		if !ok {
			return nil, nil, nil, fmt.Errorf("fireweed.yaml module_defaults must be a mapping")
		}
	}
	return config, patterns, defaults, nil
}

func parseModule(data []byte, defaults map[string]any, rel string) (model.Module, error) {
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return model.Module{}, moduleCause(rel, err)
	}
	section, _ := contracts.Section(doc.Body, "Specs")
	if _, err := contracts.ParseItems(section); err != nil {
		return model.Module{}, moduleCause(rel, err)
	}
	name, ok := doc.Frontmatter["name"].(string)
	if !ok {
		return model.Module{}, loadCause("C12", fmt.Errorf("%s/SPEC.md name must be a string", rel))
	}
	prefix, ok := doc.Frontmatter["id_prefix"].(string)
	if !ok {
		return model.Module{}, loadCause("C12", fmt.Errorf("%s/SPEC.md id_prefix must be a string", rel))
	}
	var dependencies []string
	if value, ok := doc.Frontmatter["depends_on"]; ok {
		dependencies, err = stringList(value, rel+"/SPEC.md depends_on")
		if err != nil {
			return model.Module{}, loadCause("C12", err)
		}
	}
	zones, err := moduleZones(doc.Frontmatter, defaults, rel)
	if err != nil {
		return model.Module{}, loadCause("C12", err)
	}
	return model.Module{
		Name: name, IDPrefix: prefix, Dir: rel, DependsOn: dependencies,
		Zones: zones, Document: doc,
	}, nil
}

func matchesModule(patterns []string, dir string) bool {
	for _, pattern := range patterns {
		if contracts.MatchGlob(pattern, dir) {
			return true
		}
	}
	return false
}

func stringList(value any, label string) ([]string, error) {
	items, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("%s must be a list of strings", label)
	}
	result := make([]string, 0, len(items))
	for _, item := range items {
		text, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("%s must be a list of strings", label)
		}
		result = append(result, text)
	}
	return result, nil
}

func moduleZones(frontmatter, defaults map[string]any, dir string) (contracts.Zones, error) {
	var zones contracts.Zones
	fields := []struct {
		name string
		into *[]string
	}{
		{"protected", &zones.Protected},
		{"evals", &zones.Evals},
		{"interface", &zones.Interface},
		{"append_only", &zones.AppendOnly},
		{"preserve", &zones.Preserve},
		{"disposable", &zones.Disposable},
	}
	for _, field := range fields {
		value, exists := frontmatter[field.name]
		label := dir + "/SPEC.md " + field.name
		if !exists {
			value, exists = defaults[field.name]
			label = "fireweed.yaml module_defaults " + field.name
		}
		if exists {
			list, err := stringList(value, label)
			if err != nil {
				return contracts.Zones{}, err
			}
			*field.into = list
		}
	}
	return zones, nil
}

func ValidateGraph(modules []model.Module) error {
	byName := make(map[string]model.Module, len(modules))
	prefixes := make(map[string]string, len(modules))
	for _, module := range modules {
		if module.Name == "" {
			return fmt.Errorf("empty module name")
		}
		if module.IDPrefix == "" {
			return fmt.Errorf("empty ID prefix for module %s", module.Name)
package queue

import (
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var handleReference = regexp.MustCompile(`new:[a-z0-9-]+`)
var normative = regexp.MustCompile(`\b(MUST|SHOULD)\b`)

func Check(root string) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	addCode := func(code, path, message string) {
		diagnostics = append(diagnostics, model.Diagnostic{Code: code, Path: path, Message: message})
	}
	snapshot, err := fixture.Load(root)
	if err != nil {
		addCode(fixture.LoadCode(err), fixture.LoadPath(err), err.Error())
		return diagnostics
	}
	if err := fixture.ValidateGraph(snapshot.Modules); err != nil {
		addCode("C02", "fireweed.yaml", err.Error())
	}
	p, err := newProjection(snapshot)
	if err != nil {
		addCode(errorCode(err, "C03"), "SPEC.md", err.Error())
		return diagnostics
	}
	for _, module := range snapshot.Modules {
		path := module.Dir + "/SPEC.md"
		for _, entry := range p.items[module.Name] {
			item := entry.item
			if !strings.HasPrefix(item.ID, module.IDPrefix+"-") {
				addCode("C05", path, "foreign spec ID "+item.ID)
			}
			checkItem(path, item, true, addCode)
		}
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		addCode(errorCode(err, "C03"), "changes", err.Error())
		return diagnostics
	}
	for _, change := range changes {
		path := change.Path
		motivation, exists := contracts.Section(change.Document.Body, "Motivation")
		if !exists || strings.TrimSpace(motivation) == "" {
			addCode("C07", path, "missing Motivation")
		}
		if _, hasEvalFix := change.Document.Frontmatter["eval_fix"]; hasEvalFix {
			_, hasFixes := change.Document.Frontmatter["fixes"]
			if hasFixes || len(change.Added) > 0 || len(change.Modified) > 0 || len(change.Removed) > 0 {
				addCode("C07", path, "eval_fix cannot be combined with fixes or spec edits")
			}
		}
		declared, seen := map[string]bool{}, map[string]bool{}
		for _, item := range change.Added {
			if strings.HasPrefix(item.ID, "new:") {
				declared[item.ID] = true
			}
		}
		for _, group := range []struct {
			items     []contracts.Item
			normative bool
		}{
			{change.Added, true}, {change.Modified, true}, {change.Removed, false},
		} {
			for _, item := range group.items {
				if seen[item.ID] {
					addCode("C05", path, "duplicate edit identity "+item.ID)
				}
				seen[item.ID] = true
				if !strings.HasPrefix(item.ID, "new:") && !strings.HasPrefix(item.ID, modulePrefix(snapshot, change.Module)+"-") {
					addCode("C05", path, "foreign module ID "+item.ID)
				}
				checkItem(path, item, group.normative, addCode)
			}
		}
		for _, ref := range handleReference.FindAllString(change.Document.Body, -1) {
			if !declared[ref] {
				addCode("C09", path, "reference to undeclared handle "+ref)
			}
		}
		var bases map[string]any
		if raw, exists := change.Document.Frontmatter["bases"]; exists {
			bases, _ = raw.(map[string]any)
		}
		for _, item := range append(append([]contracts.Item{}, change.Modified...), change.Removed...) {
			base, ok := p.find(change.Module, item.ID)
			if !ok {
				addCode("C08", path, "ID "+item.ID+" absent from preceding projection")
				continue
			}
			got, ok := bases[item.ID].(string)
			if !ok || got == "" {
				addCode("C08", path, "missing base for "+item.ID)
			} else if got != contracts.SpecHash(base) {
				addCode("C08", path, "stale base for "+item.ID)
			}
		}
		if err := p.apply(change); err != nil {
			addCode(errorCode(err, "C02"), path, err.Error())
		}
	}
	return diagnostics
}

func modulePrefix(snapshot model.Snapshot, name string) string {
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module.IDPrefix
		}
	}
	return ""
}

func checkItem(path string, item contracts.Item, requireNormative bool, addCode func(string, string, string)) {
	if tag := item.Tag; tag != "" && tag != "review" && !(strings.HasPrefix(tag, "quarantine: ") && strings.TrimSpace(strings.TrimPrefix(tag, "quarantine: ")) != "") {
		addCode("C06", path, "invalid review/quarantine tag "+tag)
	}
	if requireNormative && !normative.MatchString(item.Body) {
		addCode("C04", path, "non-normative spec body for "+item.ID)
	}
}
}

func Project(root, through string) (model.Snapshot, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Snapshot{}, err
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		return model.Snapshot{}, err
	}
	limit := len(changes)
	if through != "" {
		index, err := resolve(changes, through)
		if err != nil {
			return model.Snapshot{}, err
		}
		limit = index + 1
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return model.Snapshot{}, err
	}
	for _, change := range changes[:limit] {
		if err := p.apply(change); err != nil {
			return model.Snapshot{}, err
		}
	}
	if err := fixture.ValidateGraph(p.snapshot.Modules); err != nil {
		return model.Snapshot{}, err
	}
	return p.result(), nil
}

func Stamp(root, target string) error {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return err
163:func existing(root, slug string) (claim, string, error) {
181:func state(root string, ch project.Change, c claim, oid, head string) model.State {
420:func Current(root, target string) (model.State, error) {
package content

import "fmt"

func Version() string { return "0.1.0" }

const help = `fireweed 0.1.0

Setup:        init, doctor
Authoring:    new-module, new-change, stamp
Reading:      list, status, show, modules, context, explain, playbook
Checks:       check, coverage, check-commits, ci, pr-body
Lifecycle:    start, verify, done, next, run, release, commits, abandon
Regeneration: drill
Maintenance:  gc, migrate, agents-md
Information:  version

gc prunes stale claims, attempt refs for archived changes, and drill refs older than gc.drill_max_age.
gc.drill_max_age is a positive Go duration in fireweed.yaml.
gc.drill_max_age has a default of 720h (30 days).
`

func Help() string { return help }

var playbooks = map[string]string{
	"author-change": `# Author change
Read the design, SYSTEM.md and the module spec before writing a change. Give its Motivation a user-visible need or a clause of the design or SYSTEM.md; a probe showing a theoretical failure is insufficient. State each observable behavior at the module interface and keep the change focused. If a requirement costs more than it is worth, narrow or remove the spec and use already_satisfied for behavior the current implementation already meets. Do not commit during a worker step.`,
	"evals": `# Write evals
Read the change, projected spec and public interface. Add tests for every assigned added, modified or fixed spec ID in the evals zone. Exercise observable behavior through the public interface. Run the tests against the current implementation and record genuine fail-first evidence: at least one test for each ID must fail, unless the change marks it already_satisfied. Keep implementation out of this step and do not commit.

Keep temporary stub behavior in disposable files that remain implementation-writable. Implementation must be able to replace each stub without editing frozen interface files. Before completing evals, identify the implementation-writable files for each missing behavior; create a disposable stub file if needed.

For eval_fix, correct only evals, retain coverage of listed IDs, and require all module evals to pass; fail-first is not required.

Decisions contain only previously unspecified behavioral choices. Write one or two sentences per choice in Decisions. Put verification, debugging and recovery records in Notes.
Decision: When a lookup matches more than one item, return an ambiguity error.
Note: Verification passed after rerunning the module tests.`,
	"audit": `# Audit evals
Read the change, spec and evals only. Report only exact spec clauses with no test assertion; do not raise findings about helpers, fixtures, observers, unnamed edge cases or conventions. Quote each missing clause verbatim and name its spec ID. Re-audits must limit findings to IDs from the prior findings; passed IDs stay passed. If there are no findings, say so and list checked IDs. Edit only the change's Audit section and do not commit.

For a passing audit, use this example format:

## Audit

No findings
Checked IDs: AST-009, AST-010

Replace these illustrative IDs with the actual checked IDs.

For eval_fix, check that corrected tests still assert the listed specs.`,
	"impl": `# Implement
Read the change, projected spec, public interface and evals. Implement only the assigned behavior in the module's disposable code. Keep the public interface, evals, specs and protected files unchanged. Run fireweed verify to check the work; fireweed owns phase commits, so do not commit.

For eval_fix, edit only the current change file and leave implementation unchanged.

Decisions contain only previously unspecified behavioral choices. Write one or two sentences per choice in Decisions. Put verification, debugging and recovery records in Notes.
Decision: When a lookup matches more than one item, return an ambiguity error.
Note: Verification passed after rerunning the module tests.`,
	"regenerate": `# Regenerate
Read the module spec, Decisions, public interface and evals. Rebuild disposable implementation files from those sources after the controller removes the old files. Run the module's evals and dependent evals through fireweed verify. Do not consult deleted code, prior implementations or history, and do not commit.

Decisions contain only previously unspecified behavioral choices. Write one or two sentences per choice in Decisions. Put verification, debugging and recovery records in Notes.
Decision: When a lookup matches more than one item, return an ambiguity error.
Note: Verification passed after rerunning the module tests.`,
}

func Playbook(name string) (string, error) {
	if body, ok := playbooks[name]; ok {
		return body, nil
	}
	return "", fmt.Errorf("unknown playbook %q", name)
}

var explanations = map[string]string{
	"B01":   "Budget warning: summed JUnit testcase seconds for evaluated modules exceed budgets.verify_time (default 15m). Review expensive or redundant eval work while preserving required assertions.",
	"B02":   "Budget warning: a spec's testcase count exceeds budgets.cases_per_spec (default 15). Consolidate redundant cases while preserving required behavioral assertions.",
	"B03":   "Budget warning: eval lines divided by disposable lines exceed budgets.eval_ratio (default 4); positive eval lines with zero disposable lines also trigger this warning. Simplify redundant eval or fixture code while preserving required assertions.",
	"B04":   "Warning: pr-body reports binary files or files larger than 256 KB newly added to a disposable zone as opaque additions. Review these opaque additions and remove or replace unnecessary generated or vendored artifacts.",
	"C01":   "The project configuration is invalid: fireweed.yaml is missing or has an invalid required setting. Edit the configuration file to correct the reported field, then run fireweed check again.",
	"C02":   "The module graph is invalid because of an empty module name, a duplicate module name, an empty module prefix, a duplicate module prefix, or a dependency that names an unknown module or creates a cycle. Correct module names and prefixes in SPEC.md so they are nonempty and unique. Correct the module dependency in SPEC.md so it names an existing module and the graph is acyclic.",
	"C03":   "A document or spec item is unreadable or malformed. Repair the document syntax or restore the missing file, then check it again.",
	"C04":   "The spec item is non-normative because its body lacks MUST or SHOULD wording. Rewrite the spec with a MUST or SHOULD requirement that states the expected behavior.",
	"C05":   "A spec ID uses a foreign module prefix or duplicates another ID, or a change has a duplicate edit identity. Change a foreign ID to use the owning module prefix, assign a unique unused ID for a duplicate spec ID, or remove or combine the repeated edit.",
	"C06":   "The spec has an invalid review or quarantine tag. Correct the tag to use a valid review or quarantine form, or remove it.",
	"C07":   "The change has a missing Motivation section or its body is empty, or the pending change module target is missing or invalid. Add a Motivation explaining the user-visible or design need, and set the pending change module to an existing module. An eval_fix change mixed with fixes or spec edits is also invalid. Separate eval-fix work from fixes or spec changes.",
	"C08":   "An edit base is missing or stale, or its target ID is absent from the preceding projection. Stamp or update the base against the preceding projection, and correct or remove an edit whose target is absent.",
	"C09":   "A change references an undeclared handle. Declare the handle in Added, or correct or remove the reference.",
	"C10":   "Required coverage is missing a mapped non-skipped testcase for an enforced spec ID. Add and map a non-skipped testcase to that ID, then rerun the eval; the testcase may pass or fail while coverage is established.",
	"C11":   "An eval hook's JUnit report or testcase mapping is invalid or missing, or its exit status disagrees with the reported results; hook launch or setup may also have failed. Fix the hook, report or mapping, then rerun the eval and check that the exit status matches the results.",
	"C12":   "Module metadata or a zone definition is invalid or missing, or local recipe metadata is invalid. Correct the metadata field or define a valid zone in the module specification; repair invalid local recipe metadata in recipe.yaml.",
	"C13":   "A duplicate queue slug or an unparseable numeric queue prefix makes the queue identity invalid; tied numeric positions are allowed. Choose an unused slug with a valid numeric prefix.",
	"C14":   "Doctor could not prepare or clean up its scratch worktree, found a missing or invalid hook, template or canary, found a canary that failed to detect its violation, or encountered a recovery failure after a canary run. Fix the scratch worktree setup if reported, or repair the named hook, template or canary, then rerun doctor to confirm detection and recovery.",
	"C15":   "A worker or audit agent probe failed because its command failed or is unavailable, its model is invalid or unavailable, or login failed; a usage limit may also stop the probe. Check and fix the configured command, model and login, or wait for the usage limit to reset and check access before rerunning doctor. Warning: the audit uses a shared agent command, leaving a shared-command blind spot; configure an independent audit command.",
	"C16":   "Managed guidance in AGENTS.md is missing or stale, or a playbook override has a missing or stale embedded_base. Run agents-md --write to restore the managed guidance; review and update the playbook override's embedded_base against the current embedded playbook.",
	"E01":   "An enforced testcase failed or reported an error. Fix the failing test or implementation, then rerun the eval.",
	"H01":   "A worker phase commit changed a protected path or a project playbook override. Restore the edit. Separate human-owned setup and human-owned override work into an author commit outside worker phases.",
	"H02":   "An impl-to-applied commit changed evals or interface files. Make eval edits and interface edits in the evals step before implementation. Rework the affected phase history before rerunning history checks.",
	"H03":   "Modifying an existing append-only path, deleting an existing append-only path, or renaming an existing append-only path violates its append-only rule. Restore the original path and original content. Add a new file for new data.",
	"H04":   "Modifying an archived change file violates archive history. Deleting an archived change file is allowed for a revert. Restore the archive and author follow-up work as a pending change.",
	"X01":   "A boundary violation was reported for an import crossing a module dependency or interface boundary. Fix the import or declare the required dependency.",
	"X02":   "A required boundaries hook is missing, or a configured boundary hook failed to complete. Configure a boundaries hook for the module; fix and rerun the hook if it fails or times out.",
	"X03":   "Warning: a module declares a boundaries opt-out with a reason. Review the reason and remove the opt-out when it no longer applies; configure a boundary hook if enforcement is needed.",
	"X04":   "Warning: no boundary enforcement is required, and the boundaries hook is absent. No action is needed; configure a boundaries hook if a module later needs enforcement.",
	"U01":   "An invalid argument or option, or an unsupported command, was requested. Check fireweed help and use the supported command syntax and options.",
	"U02":   "An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.",
	"P01":   "A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.",
	"P02":   "A change claim is missing, held by another owner, or unreadable or undecodable because its claim data is invalid. If no claim exists, start an eligible change. If claim data is unreadable, undecodable, or invalid, seek approved human review and recovery before retrying. If another owner holds the claim, wait for the holder to release it or coordinate with that holder to resolve ownership before retrying.",
	"P03":   "The working state has an unapproved alteration. Restore the approved state or have the alteration reviewed and approved through the normal workflow.",
	"P04":   "Warning: approval for this change is not recorded on the base branch. Obtain approval through review before unattended execution.",
	"P05":   "A change base is stale or mismatches the intended state. Refresh the base through the approved workflow, or restore the intended state through that workflow.",
	"P06":   "The change is blocked by unresolved Questions. Review the Questions and resolve them before continuing the change.",
	"P07":   "The shared-base unapproved-work bound has been reached. Review the shared work and synchronize the shared base after approval before starting more unapproved changes.",
	"P08":   "Warning: at least three consecutive applied changes target the same module. Review the sequence's scope with the project owner before continuing.",
	"P09":   "An audit submission is incomplete when No findings or the Checked IDs list is missing. Include No findings and a complete Checked IDs list in the audit submission.",
	"P10":   "Non-blocking regenerate warning: no passing module drill is newer than the module's latest applied implementation. Run a module drill, then review and commit its report.",
	"P2-01": "Genuine fail-first evidence is missing for an assigned spec ID: its new eval did not fail against the pre-implementation code. Write or correct a test assertion for the specified behavior, rerun the eval before implementing it, and record the failing result.",
	"Z01":   "A forbidden phase-zone edit changed a file outside the authorized zone. Restore the forbidden edit or move the work to the appropriate authorized phase and zone.",
}

func Explain(code string) (string, error) {
	if body, ok := explanations[code]; ok {
		return body, nil
	}
	return "", fmt.Errorf("unknown diagnostic code %q", code)
}
package recipe

import (
	"embed"
	"fmt"

	"github.com/lurkshark/fireweed/packages/assets/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

//go:embed files/*
var bundle embed.FS

func BuiltinRecipe(name string) (model.Recipe, error) {
	if name != "go" && name != "typescript-pnpm" && name != "python-uv" {
		return model.Recipe{}, fmt.Errorf("unknown recipe %q", name)
	}
	if name != "go" {
		return additionalRecipe(name)
	}
	paths := map[string]string{
		".fireweed/recipes/go.md":                "files/go.md",
		".fireweed/hooks.go":                     "files/hooks.go",
		".fireweed/templates/module/SPEC.md":     "files/SPEC.md",
		".fireweed/templates/module/api.go.tmpl": "files/api.go.tmpl",
		".fireweed/canaries/canaries.yaml":       "files/canaries.yaml",
	}
	files := make(map[string]string, len(paths))
	for output, resource := range paths {
		data, err := bundle.ReadFile(resource)
		if err != nil {
			return model.Recipe{}, fmt.Errorf("read Go recipe %s: %w", resource, err)
		}
		files[output] = string(data)
	}
	return model.Recipe{
		Name:         "go",
		Instructions: files[".fireweed/recipes/go.md"],
		ModuleDefaults: contracts.Zones{
			Interface:  []string{"api.go"},
			Evals:      []string{"evals/**"},
			Disposable: []string{"internal/**"},
		},
		Hooks: map[string]any{
			"eval":       "go run .fireweed/hooks.go eval {module_dir} {report}",
			"boundaries": "go run .fireweed/hooks.go boundaries",
			"setup":      "go mod download",
		},
		Protected:      []string{".fireweed/**", "fireweed.yaml", "AGENTS.md", "SYSTEM.md"},
		WorkerWritable: []string{"go.sum"},
		Files:          files,
	}, nil
}

func additionalRecipe(name string) (model.Recipe, error) {
	paths := map[string]string{
		".fireweed/recipes/" + name + ".md": "files/" + name + ".md",
		".fireweed/templates/module/SPEC.md": "files/" + name + "-SPEC.md",
		".fireweed/canaries/canaries.yaml": "files/" + name + "-canaries.yaml",
	}
	var zones contracts.Zones
	var hooks map[string]any
	var writable []string
	if name == "typescript-pnpm" {
		paths[".fireweed/hooks.mjs"] = "files/typescript-pnpm-hooks.mjs"
		paths[".fireweed/templates/module/package.json.tmpl"] = "files/typescript-pnpm-package.json.tmpl"
		paths[".fireweed/templates/module/src/index.ts.tmpl"] = "files/typescript-pnpm-index.ts.tmpl"
		zones = contracts.Zones{Interface: []string{"src/index.ts"}, Evals: []string{"evals/**"}, Disposable: []string{"src/internal/**"}}
		hooks = map[string]any{"eval": "pnpm exec node .fireweed/hooks.mjs eval {module_dir} {report}", "boundaries": "pnpm exec node .fireweed/hooks.mjs boundaries", "setup": "pnpm install --offline"}
		writable = []string{"pnpm-lock.yaml"}
	} else {
		paths[".fireweed/hooks.py"] = "files/python-uv-hooks.py"
		paths[".fireweed/templates/module/pyproject.toml.tmpl"] = "files/python-uv-pyproject.toml.tmpl"
		paths[".fireweed/templates/module/__init__.py.tmpl"] = "files/python-uv-init.py.tmpl"
		zones = contracts.Zones{Interface: []string{"__init__.py"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}}
		hooks = map[string]any{"eval": "uv run --no-project python .fireweed/hooks.py eval {module_dir} {report}", "boundaries": "uv run --no-project python .fireweed/hooks.py boundaries", "setup": "uv sync --offline"}
		writable = []string{"uv.lock"}
	}
	files := make(map[string]string, len(paths))
	for output, resource := range paths {
		data, err := bundle.ReadFile(resource)
		if err != nil {
			return model.Recipe{}, fmt.Errorf("read %s recipe %s: %w", name, resource, err)
		}
		files[output] = string(data)
	}
	return model.Recipe{Name: name, Instructions: files[".fireweed/recipes/"+name+".md"], ModuleDefaults: zones, Hooks: hooks,
		Protected: []string{".fireweed/**", "fireweed.yaml", "AGENTS.md", "SYSTEM.md"}, WorkerWritable: writable, Files: files}, nil
}

exec
/bin/zsh -lc "sed -n '420,468p' packages/workflow/internal/lifecycle/stub.go; sed -n '1,120p' packages/project/internal/queue/project.go; sed -n '115,214p' packages/app/evals/app_test.go; rg -n 'migrateConfirmation|doctorCLIManagedBlock|Check passed|Contains\\(' packages/app/evals/diagnostic_causes_test.go; rg -n 'func (Next|Run|Execute|Drill)|type .*Error|invalid|project.Load' packages/runner/internal/{next,run,drill}/*.go | head -100; cat go.mod; cat .fireweed/setup.sh" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 482ms:
func Current(root, target string) (model.State, error) {
	repo, e := gitstore.Inspect(root)
	if e != nil {
		return model.State{}, e
	}
	root = repo.Root
	if target == "" {
		q, e := project.Queue(root)
		if e != nil {
			return model.State{}, e
		}
		for _, ch := range q {
			c, oid, e := existing(root, ch.Slug)
			if e == nil && c.Worktree == root && c.Change == ch.Slug {
				return state(root, ch, c, oid, repo.Head), nil
			}
		}
		return model.State{}, errors.New("P02: worktree holds no change")
	}
	ch, e := changeBy(root, target)
	if e != nil {
		return model.State{}, e
	}
	c, oid, e := existing(root, ch.Slug)
	if e != nil {
		return model.State{}, e
	}
	if c.Worktree != root || c.Change != ch.Slug {
		return model.State{}, errors.New("P02: claim belongs to another worktree")
	}
	return state(root, ch, c, oid, repo.Head), nil
}
func validate(s model.State, c claim) error {
	if s.Change.Slug != c.Change || str(s.Change.Document.Frontmatter["run"]) != c.Run || str(s.Change.Document.Frontmatter["phase"]) != c.Phase {
		return errors.New("P03: approved run or phase changed")
	}
	a, sections := author(s.Change.Document)
	if !equal(a, c.Author) || !equal(sections, c.Sections) {
		return errors.New("P03: approved author metadata changed")
	}
	if !equal(auditState(s.Change.Document), c.Audit) {
		return errors.New("P03: approved audit state changed")
	}
	return nil
}
func loadCurrent(root, target string) (model.State, claim, error) {
	s, e := Current(root, target)
	if e != nil {
		return s, claim{}, e
package queue

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

// An owner distinguishes equal handle text declared by separate changes.
type ownedItem struct {
	item  contracts.Item
	owner string
}

type projection struct {
	snapshot model.Snapshot
	items    map[string][]ownedItem
}

func newProjection(snapshot model.Snapshot) (*projection, error) {
	p := &projection{snapshot: snapshot, items: map[string][]ownedItem{}}
	for _, module := range snapshot.Modules {
		section, _ := contracts.Section(module.Document.Body, "Specs")
		items, err := contracts.ParseItems(section)
		if err != nil {
			return nil, withCause(itemErrorCode(err), fmt.Errorf("%s/SPEC.md: %w", module.Dir, err))
		}
		for _, item := range items {
			p.items[module.Name] = append(p.items[module.Name], ownedItem{item: item})
		}
	}
	return p, nil
}

func (p *projection) find(module, id string) (contracts.Item, bool) {
	for _, entry := range p.items[module] {
		if entry.item.ID == id && entry.owner == "" {
			return entry.item, true
		}
	}
	return contracts.Item{}, false
}

func (p *projection) apply(change model.Change) error {
	index := -1
	for i := range p.snapshot.Modules {
		if p.snapshot.Modules[i].Name == change.Module {
			index = i
			break
		}
	}
	if index < 0 {
		return withCause("C07", fmt.Errorf("unknown module %q", change.Module))
	}
	if err := applySet(&p.snapshot.Modules[index], change.Document.Frontmatter["set"]); err != nil {
		return withCause("C12", fmt.Errorf("%s: %w", change.Path, err))
	}
	items := p.items[change.Module]
	for _, edit := range change.Modified {
		found := false
		for i := range items {
			if items[i].item.ID == edit.ID && items[i].owner == "" {
				items[i].item = edit
				found = true
				break
			}
		}
		if !found {
			return withCause("C08", fmt.Errorf("%s: Modified ID %s absent from preceding projection", change.Path, edit.ID))
		}
	}
	for _, edit := range change.Removed {
		found := false
		for i := range items {
			if items[i].item.ID == edit.ID && items[i].owner == "" {
				items = append(items[:i], items[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			return withCause("C08", fmt.Errorf("%s: Removed ID %s absent from preceding projection", change.Path, edit.ID))
		}
	}
	for _, edit := range change.Added {
		owner := ""
		if strings.HasPrefix(edit.ID, "new:") {
			owner = change.Path
		}
		items = append(items, ownedItem{item: edit, owner: owner})
	}
	p.items[change.Module] = items
	if err := fixture.ValidateGraph(p.snapshot.Modules); err != nil {
		return withCause("C02", err)
	}
	return nil
}

func applySet(module *model.Module, raw any) error {
	if raw == nil {
		return nil
	}
	settings, ok := raw.(map[string]any)
	if !ok {
		return fmt.Errorf("set must be a mapping")
	}
	front := make(map[string]any, len(module.Document.Frontmatter)+len(settings))
	for key, value := range module.Document.Frontmatter {
		front[key] = value
	}
	for key, value := range settings {
		switch key {
		case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
			list, err := asStrings(value)
			if err != nil {

func Test_APP_003_InvalidRequests(t *testing.T) {
	for _, tc := range []struct {
		name string
		args []string
		bad  string
		code string
	}{
		{"unknown-command", []string{"not-a-command"}, "not-a-command", "U01"},
		{"unknown-option", []string{"help", "--bogus"}, "--bogus", "U01"},
		{"missing-argument", []string{"explain"}, "explain", "U01"},
		{"extra-argument", []string{"version", "surplus"}, "version", "U01"},
		{"new-change-arity", []string{"new-change", "alpha"}, "new-change", "U01"},
		{"stamp-arity", []string{"stamp"}, "stamp", "U01"},
		{"check-arity", []string{"check", "extra"}, "check", "U01"},
		{"show-arity", []string{"show"}, "show", "U01"},
		{"start-arity", []string{"start"}, "start", "U01"},
		{"context-arity", []string{"context"}, "context", "U01"},
		{"coverage-arity", []string{"coverage"}, "coverage", "U01"},
		{"verify-arity", []string{"verify", "manual", "extra"}, "verify", "U01"},
		{"APP_013_unknown-explanation", []string{"explain", "NO-SUCH-CODE"}, "NO-SUCH-CODE", "U02"},
		{"APP_013_unknown-playbook", []string{"playbook", "no-such-playbook"}, "no-such-playbook", "U02"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, format := range []struct {
				name string
				json bool
			}{{"plain", false}, {"JSON", true}} {
				args := append([]string(nil), tc.args...)
				if format.json {
					args = append(args, "--json")
				}
				status, stdout, stderr := run(args...)
				if status != 2 {
					t.Errorf("%s status=%d; want 2", format.name, status)
				}
				diagnostic := stdout + stderr
				if format.json {
					diagnostic = assertFailedJSON(t, stdout, tc.code)
				} else if !strings.HasPrefix(strings.TrimSpace(diagnostic), tc.code+":") {
					t.Errorf("%s diagnostic %q must have code %s", format.name, diagnostic, tc.code)
				}
				if !strings.Contains(diagnostic, tc.bad) || !corrective(diagnostic) {
					t.Errorf("%s diagnostic %q must identify %q and suggest a correction", format.name, diagnostic, tc.bad)
				}
			}
		})
	}
}

func Test_APP_014_MaintenanceCommandsDispatch(t *testing.T) {
	bin := buildAppBinary(t)
	root := migrateProject(t)
	for _, command := range []string{"gc", "migrate"} {
		t.Run(command, func(t *testing.T) {
			if !strings.Contains(assets.Help(), command) {
				t.Fatalf("%q is not listed by help", command)
			}
			status, plain, stderr := runAppBinary(t, bin, root, command)
			if status != 0 || stderr != "" || strings.TrimSpace(plain) == "" {
				t.Fatalf("%s plain: status=%d stdout=%q stderr=%q", command, status, plain, stderr)
			}
			if regexp.MustCompile(`(?i)unimplemented|not implemented|under construction`).MatchString(plain) {
				t.Errorf("%s still reported a placeholder: %q", command, plain)
			}
			status, output, stderr := runAppBinary(t, bin, root, command, "--json")
			if status != 0 || stderr != "" {
				t.Fatalf("%s JSON: status=%d stdout=%q stderr=%q", command, status, output, stderr)
			}
			env := decodeEnvelope(t, output)
			if string(env["data"]) == "null" {
				t.Errorf("%s JSON data is null", command)
			}
			var diagnostics []json.RawMessage
			if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
				t.Errorf("%s JSON diagnostics=%s: %v", command, env["diagnostics"], err)
			}
			for _, args := range [][]string{{command, "extra"}, {command, "--bogus"}} {
				for _, jsonMode := range []bool{false, true} {
					call := append([]string(nil), args...)
					if jsonMode {
						call = append(call, "--json")
					}
					status, stdout, errOut := runAppBinary(t, bin, root, call...)
					message := commandFailure(t, status, 2, stdout, errOut, jsonMode)
					if jsonMode {
						message = assertFailedJSON(t, stdout, "U01")
					} else if !strings.HasPrefix(strings.TrimSpace(errOut), "U01:") {
						t.Errorf("%s invalid plain diagnostic=%q; want U01", command, errOut)
					}
					if !corrective(message) {
						t.Errorf("%s invalid diagnostic lacks guidance: %q", command, message)
					}
				}
			}
		})
	}
}

func Test_APP_015_FailedJSON(t *testing.T) {
38:				if diagnostic.Code == code && strings.Contains(diagnostic.Message, cause) {
45:		} else if !strings.Contains(stdout+stderr, code+": ") || !strings.Contains(stdout+stderr, cause) {
47:		} else if strings.Contains(stdout+stderr, code+": "+code+":") {
182:		if !strings.Contains(message, "create changes/0003-blocked.md") || !corrective(message) {
packages/runner/internal/run/stub.go:14:func Execute(ctx context.Context, request model.RunRequest) (model.RunResult, error) {
packages/runner/internal/next/usage.go:46:		return 0, fmt.Errorf("C01: invalid %s %q: %w", name, text, err)
packages/runner/internal/next/usage.go:58:type TimeStopError struct{ Limit time.Duration }
packages/runner/internal/next/usage.go:102:type UsageStopError struct {
packages/runner/internal/next/run.go:24:func Run(ctx context.Context, request model.NextRequest) (model.NextResult, error) {
packages/runner/internal/next/run.go:29:func RunWithBudget(ctx context.Context, request model.NextRequest, budget *UsageBudget) (model.NextResult, error) {
packages/runner/internal/next/run.go:48:		snapshot, err := project.Load(request.Root)
packages/runner/internal/next/run.go:234:type BlockedError struct {
packages/runner/internal/next/run.go:341:	snapshot, err := project.Load(root)
packages/runner/internal/next/approval.go:8:type approvalRequiredError struct {
packages/runner/internal/drill/drill.go:26:func Run(ctx context.Context, request model.DrillRequest) (result model.DrillResult, err error) {
packages/runner/internal/drill/drill.go:29:	snapshot, err := project.Load(request.Root)
packages/runner/internal/drill/drill.go:160:			failure = fmt.Errorf("invalid reconstruction edits: %w", validationErr)
packages/runner/internal/drill/drill.go:325:				return fmt.Errorf("invalid setup argv")
packages/runner/internal/drill/drill.go:332:		return fmt.Errorf("invalid setup hook")
packages/runner/internal/drill/drill.go:585:		return "", fmt.Errorf("C11: missing or invalid eval hook for %s", target.Name)
packages/runner/internal/drill/drill.go:856:			return 0, fmt.Errorf("C01: invalid %s: %w", key, err)
packages/runner/internal/drill/drill.go:869:		return 0, 0, fmt.Errorf("C01: invalid usage wait or budget")
module github.com/lurkshark/fireweed

go 1.23

require gopkg.in/yaml.v3 v3.0.1
#!/bin/sh
set -eu
# FIREWEED_BIN is supplied by the invoking controller, including in scratch trees.
mkdir -p .fireweed/bin
if [ "$(cd .fireweed/bin && pwd -P)/fireweed" != "$FIREWEED_BIN" ]; then
    controller_copy=$(mktemp .fireweed/bin/fireweed.XXXXXX)
    trap 'rm -f "$controller_copy"' EXIT HUP INT TERM
    cp "$FIREWEED_BIN" "$controller_copy"
    chmod 755 "$controller_copy"
    mv -f "$controller_copy" .fireweed/bin/fireweed
    trap - EXIT HUP INT TERM
fi

exec
/bin/zsh -lc "rg -n 'func NewChange|func History|func Doctor|func .*config|func resolve' packages/{project,checks}/internal; sed -n '150,220p' packages/app/evals/lifecycle_cli_test.go; sed -n '53,160p' packages/app/evals/diagnostic_causes_test.go; rg -n 'migrateConfirmation|doctorCLIManagedBlock' packages/app/evals/{doctor_setup_test,migrate_cli_test}.go; rg -n 'env|proxy|GOMOD' .fireweed/hooks.go; sed -n '105,125p' packages/runner/internal/next/usage.go; cat packages/runner/internal/run/stub.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 503ms:
packages/checks/internal/engine/doctor.go:74:func Doctor(ctx context.Context, request model.DoctorRequest) model.DoctorReport {
packages/project/internal/queue/queue.go:95:func resolve(changes []model.Change, target string) (int, error) {
packages/checks/internal/engine/history.go:247:func History(root, revisionRange string) (model.HistoryReport, error) {
packages/project/internal/fixture/fixture.go:272:func NewChange(root, module, slug string) (string, error) {
	if string(envelope["data"]) == "null" {
		t.Fatalf("result data is null: %s", output)
	}
	var result map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &result); err != nil || result == nil {
		t.Fatalf("result data=%s: %v", envelope["data"], err)
	}
	return result
}

func cliResultPhase(t *testing.T, output, phase string) {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var result struct {
		State struct {
			Phase string `json:"Phase"`
		} `json:"State"`
	}
	if err := json.Unmarshal(envelope["data"], &result); err != nil || result.State.Phase != phase {
		t.Errorf("result phase=%s; want %s: %v", envelope["data"], phase, err)
	}
}

func lifecycleCommandFailure(t *testing.T, status, want int, stdout, stderr string, jsonMode bool) {
	t.Helper()
	message := commandFailure(t, status, want, stdout, stderr, jsonMode)
	if !corrective(message) {
		t.Errorf("diagnostic %q must suggest a correction", message)
	}
}

func Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection(t *testing.T) {
	bin := buildAppBinary(t)
	root := lifecycleCLIProject(t, bin, true)
	before := cliGit(t, root, "rev-parse", "HEAD")
	status, plain, stderr := runLifecycleBinary(t, bin, appNestedDir(t, root), "start", "manual")
	if status != 0 || stderr != "" || !strings.Contains(plain, "# Write evals") || !strings.Contains(plain, "ALP-002") {
		t.Fatalf("start: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	if after := cliGit(t, root, "rev-parse", "HEAD"); after != before {
		t.Errorf("start committed: %s -> %s", before, after)
	}
	state, err := workflow.Current(root, "manual")
	if err != nil || state.Phase != "evals" || state.Run == "" {
		t.Fatalf("start state=%#v, %v", state, err)
	}
	change, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(change), "run:") || !strings.Contains(string(change), "phase: evals") || !strings.Contains(string(change), "### ALP-002") || strings.Contains(string(change), "new:extra") {
		t.Errorf("assigned change=%q, %v", change, err)
	}
	expected, err := workflow.Context(root, "manual", false)
	if err != nil || plain != expected {
		t.Errorf("first-step context differs from workflow.Context: %v", err)
	}
	for _, tc := range []struct {
		args []string
		want int
	}{
		{[]string{"start"}, 2},
		{[]string{"start", "manual", "extra"}, 2},
		{[]string{"start", "missing"}, 1},
		{[]string{"start", "manual"}, 1},
	} {
		args := append(append([]string{}, tc.args...), "--json")
		status, out, errOut := runLifecycleBinary(t, bin, root, args...)
		lifecycleCommandFailure(t, status, tc.want, out, errOut, true)
	}
	bad := lifecycleCLIProject(t, bin, false)
	appPut(t, bad, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runLifecycleBinary(t, bin, bad, "start", "manual", "--json")
	lifecycleCommandFailure(t, status, 2, out, errOut, true)
func Test_APP_037_StaticCauses(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code, cause string
		args              []string
		status            int
		edit              func(*testing.T, string)
	}{
		{"configuration", "C01", "fireweed.yaml", []string{"modules"}, 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"graph", "C02", "missing module", []string{"modules"}, 2, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
		{"syntax", "C03", "malformed item heading", []string{"coverage", "alpha"}, 2, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "### ALP-001", "### broken item", 1))
		}},
		{"normative", "C04", "non-normative", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need a statement.\n", "### new:first\nA statement without a normative verb.\n", "")
		}},
		{"identity", "C05", "duplicate item ID", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need two items.\n", "### new:one\nOne MUST work.\n### new:one\nOne MUST remain unique.\n", "")
		}},
		{"tag", "C06", "invalid review/quarantine tag", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need an item.\n", "### new:one [invalid]\nOne MUST work.\n", "")
		}},
		{"structure", "C07", "missing Motivation", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "", "### new:one\nOne MUST work.\n", "")
		}},
		{"base", "C08", "absent from preceding projection", []string{"stamp", "revise"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-revise.md", "Users need a revision.\n", "", "### ALP-999\nMissing MUST change.\n")
		}},
		{"handle", "C09", "undeclared handle", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need new:missing.\n", "### new:one\nOne MUST work.\n", "")
		}},
		{"metadata", "C12", "set must be a mapping", []string{"show", "alpha", "--projected"}, 2, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "Users need a change.\n", "", "")
			path := filepath.Join(root, "changes/0001-first.md")
			body, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			appPut(t, root, "changes/0001-first.md", strings.Replace(string(body), "module: alpha\n", "module: alpha\nset: invalid\n", 1))
		}},
		{"queue-identity", "C13", "ambiguous change slug", []string{"check"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-repeat.md", "Users need one item.\n", "### new:one\nOne MUST work.\n", "")
			authoringChange(t, root, "0002-repeat.md", "Users need another item.\n", "### new:two\nTwo MUST work.\n", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			tc.edit(t, root)
			appCause(t, bin, root, tc.args, tc.status, tc.code, tc.cause)
		})
	}
}

func Test_APP_037_APP_038_ShowStaticCauses(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code, cause string
		args              []string
		edit              func(*testing.T, string)
	}{
		{"duplicate-graph-name", "C02", "duplicate module name", []string{"show", "alpha"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: alpha\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
		}},
		{"duplicate-fixed-spec-id", "C05", "duplicate item ID", []string{"show", "alpha", "ALP-002"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", authoringSpec+"\n### ALP-001\nDuplicate MUST fail.\n\n### ALP-002\nDistinct MUST remain selectable.\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			tc.edit(t, root)
			appCause(t, bin, root, tc.args, 2, tc.code, tc.cause)
		})
	}
}

func Test_APP_038_LookupCauses(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, cause string
		args        []string
		status      int
		edit        func(*testing.T, string)
	}{
		{"author-module", "unknown module", []string{"new-change", "unknown", "valid"}, 2, nil},
		{"stamp-change", "not found", []string{"stamp", "absent"}, 1, nil},
		{"show-module", "unknown module", []string{"show", "absent"}, 2, nil},
		{"show-item", "unknown item", []string{"show", "alpha", "ALP-999"}, 2, nil},
		{"show-ambiguous-item", "duplicate item ID", []string{"show", "alpha", "new:shared", "--projected"}, 2, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "First need.\n", "### new:shared\nFirst MUST work.\n", "")
			authoringChange(t, root, "0002-second.md", "Second need.\n", "### new:shared\nSecond MUST work.\n", "")
		}},
		{"coverage-module", "Unknown module", []string{"coverage", "missing"}, 2, nil},
		{"coverage-change", "not found", []string{"coverage", "alpha", "--with-change", "missing"}, 2, nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			appCause(t, bin, root, tc.args, tc.status, "U02", tc.cause)
		})
	}
}

packages/app/evals/doctor_setup_test.go:15:const doctorCLIManagedBlock = "<!-- fireweed:start -->\nThis project uses fireweed. Specs and change files are the source of truth.\nRun `fireweed playbook author-change` before authoring changes and\n`fireweed context <change>` before each step.\n<!-- fireweed:end -->"
packages/app/evals/doctor_setup_test.go:329:	if agents != "Before.\n\n"+doctorCLIManagedBlock+"\n\nAfter.\n" {
packages/app/evals/doctor_setup_test.go:357:	if got := readInitFile(t, closer, "AGENTS.md"); got != "Closer before.\n"+doctorCLIManagedBlock+"\nCloser after.\n" {
packages/app/evals/doctor_setup_test.go:384:			if !strings.HasPrefix(got, tc.initial) || strings.Count(got, "<!-- fireweed:start -->") != 1 || strings.Count(got, "<!-- fireweed:end -->") != 1 || !strings.Contains(got, doctorCLIManagedBlock) {
packages/app/evals/migrate_cli_test.go:15:const migrateConfirmation = "Project format 1 is already current.\n"
packages/app/evals/migrate_cli_test.go:118:		if status != 0 || stderr != "" || plain != migrateConfirmation {
packages/app/evals/migrate_cli_test.go:141:	if status != 0 || stderr != "" || plain != migrateConfirmation {
39: bin:=os.Getenv("FIREWEED_BIN");if bin==""{bin="fireweed"};cmd:=exec.Command(bin,"modules","--json");data,err:=cmd.Output();if err!=nil{panic(err)}
40: var envelope struct { Data []struct { Name,Dir string; DependsOn []string `json:"depends_on"` } };if err=json.Unmarshal(data,&envelope);err!=nil{panic(err)}
43: for _,m:=range envelope.Data { err=filepath.WalkDir(m.Dir,func(file string,d fs.DirEntry,err error)error{
45:  for _,imp:=range tree.Imports { imported,_:=strconv.Unquote(imp.Path.Value);for _,target:=range envelope.Data { base:=modulePath+"/"+target.Dir;if imported!=base&&!strings.HasPrefix(imported,base+"/"){continue};allowed:=m.Name==target.Name;for _,dep:=range m.DependsOn{if dep==target.Name{allowed=true}}
	Budget    time.Duration
}

func (e *UsageStopError) Error() string {
	return fmt.Sprintf("usage limit: completed wait %s; next wait %s exceeds budget %s", e.Completed, e.Next, e.Budget)
}

// CombinedUsageFailure preserves an earlier block as the primary failure.
type CombinedUsageFailure struct {
	Blocked error
	Stop    error
}

func (e *CombinedUsageFailure) Error() string   { return fmt.Sprintf("%v; %v", e.Blocked, e.Stop) }
func (e *CombinedUsageFailure) Unwrap() []error { return []error{e.Blocked, e.Stop} }

func IsUsageStop(err error) bool {
	var combined *CombinedUsageFailure
	if errors.As(err, &combined) {
		return false
	}
package run

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func Execute(ctx context.Context, request model.RunRequest) (model.RunResult, error) {
	var result model.RunResult
	budget := &next.UsageBudget{}
	if request.For > 0 {
		stop := &next.TimeStopError{Limit: request.For}
		deadline := time.Now().Add(request.For)
		var cancel context.CancelFunc
		ctx, cancel = context.WithDeadlineCause(ctx, deadline, stop)
		defer cancel()
		budget.SetTimeLimit(deadline, stop)
	}
	nextRequest := model.NextRequest{
		Root:          request.Root,
		Bin:           request.Bin,
		AllowUnmerged: request.AllowUnmerged,
		Timeout:       request.Timeout,
	}
	var blockedErr error
	for {
		change, err := next.RunWithBudget(ctx, nextRequest, budget)
		if err != nil {
			var blocked *next.BlockedError
			if errors.As(err, &blocked) {
				blockedErr = errors.Join(blockedErr, err)
				continue
			}
			if blockedErr != nil && next.IsUsageStop(err) {
				return result, &next.CombinedUsageFailure{Blocked: blockedErr, Stop: err}
			}
			if blockedErr != nil && IsTimeStop(err) {
				return result, &combinedTimeFailure{blocked: blockedErr, stop: err}
			}
			return result, errors.Join(blockedErr, err)
		}
		if change.Empty {
			queue, err := project.Queue(request.Root)
			if err != nil {
				return result, errors.Join(blockedErr, fmt.Errorf("read change queue: %w", err))
			}
			result.Empty = len(queue) == 0
			return result, blockedErr
		}
		result.Applied = append(result.Applied, change)
		if request.Max > 0 && len(result.Applied) == request.Max {
			queue, err := project.Queue(request.Root)
			if err != nil {
				return result, errors.Join(blockedErr, fmt.Errorf("read change queue: %w", err))
			}
			result.Empty = len(queue) == 0
			result.Limited = !result.Empty
			return result, blockedErr
		}
	}
}

exec
/bin/zsh -lc "sed -n '272,355p' packages/project/internal/fixture/fixture.go; sed -n '74,135p' packages/checks/internal/engine/doctor.go; sed -n '247,295p' packages/checks/internal/engine/history.go; sed -n '154,195p' packages/app/evals/project_commands_test.go; sed -n '225,310p' packages/app/evals/authoring_commands_test.go; sed -n '118,160p' packages/app/evals/next_cli_test.go; sed -n '32,60p' packages/runner/internal/next/run.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 489ms:
func NewChange(root, module, slug string) (string, error) {
	if !slugPattern.MatchString(slug) {
		return "", fmt.Errorf("invalid change slug %q", slug)
	}
	snapshot, err := Load(root)
	if err != nil {
		return "", err
	}
	found := false
	for _, item := range snapshot.Modules {
		if item.Name == module {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("unknown module %q", module)
	}
	maxNumber := 0
	for _, directory := range []string{"changes", "changes/archive"} {
		entries, err := os.ReadDir(filepath.Join(root, directory))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return "", fmt.Errorf("read %s: %w", directory, err)
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			parts := changeNamePattern.FindStringSubmatch(entry.Name())
			if parts == nil {
				continue
			}
			if parts[2] == slug {
				return "", fmt.Errorf("change slug %q already exists in %s", slug, directory)
			}
			number, err := strconv.Atoi(parts[1])
			if err != nil {
				return "", fmt.Errorf("invalid change number in %s/%s: %w", directory, entry.Name(), err)
			}
			if number > maxNumber {
				maxNumber = number
			}
		}
	}
	path := fmt.Sprintf("changes/%04d-%s.md", maxNumber+1, slug)
	body := "\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"
	content, err := contracts.RenderDocument(contracts.Document{Frontmatter: map[string]any{"module": module}, Body: body})
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Join(root, "changes"), 0755); err != nil {
		return "", fmt.Errorf("create changes directory: %w", err)
	}
	file, err := os.OpenFile(filepath.Join(root, filepath.FromSlash(path)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", path, err)
	}
	defer file.Close()
	if _, err := file.WriteString(content); err != nil {
		return "", fmt.Errorf("write %s: %w", path, err)
	}
	return path, nil
}
func Doctor(ctx context.Context, request model.DoctorRequest) model.DoctorReport {
	if ctx == nil {
		ctx = context.Background()
	}
	d := &doctorState{ctx: ctx, req: request, root: request.Root}
	d.report.Diagnostics = append(d.report.Diagnostics, project.Check(request.Root)...)
	for _, v := range d.report.Diagnostics {
		if !v.Warning {
			return d.report
		}
	}
	d.guidance()
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	snapshot, err := project.Load(request.Root)
	if err != nil {
		d.issue("C14", "project: "+err.Error(), "", false)
		return d.report
	}
	d.snapshot = snapshot
	d.boundaryWarnings(snapshot)
	if err = d.readRecipe(request.Root); err != nil {
		d.issue("C14", err.Error(), "", false)
		return d.report
	}
	base := filepath.Join(repo.CommonDir, "fireweed", "worktrees")
	if err = os.MkdirAll(base, 0755); err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	id, err := randomHex(8)
	if err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	scratch := filepath.Join(base, "doctor-"+id)
	if _, err = d.git(request.Root, "worktree", "add", "--detach", scratch, repo.Head); err != nil {
		d.issue("C14", "scratch worktree: "+err.Error(), "", false)
		return d.report
	}
	defer func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		if _, e := d.gitWithContext(cleanupCtx, request.Root, "worktree", "remove", "--force", scratch); e != nil {
			d.issue("C14", "scratch cleanup: "+e.Error(), "", false)
		}
		_ = os.RemoveAll(scratch)
	}()
	if err = d.copyCurrent(request.Root, scratch); err != nil {
		d.issue("C14", "scratch copy: "+err.Error(), "", false)
		return d.report
	}
	d.root = scratch
	setup := mapValue(snapshot.Config, "hooks")["setup"]
	if setup != nil {
		p, e := runHook(ctx, scratch, setup, nil, request.Bin, request.Timeout)
		if e != nil || p.ExitCode != 0 {
			d.fail("setup", "", "setup", fmt.Sprintf("hook %v: %v; exit %d; %s %s", setup, e, p.ExitCode, p.Stdout, p.Stderr))
			return d.report
func History(root, revisionRange string) (model.HistoryReport, error) {
	var report model.HistoryReport
	commits, err := historyRange(root, revisionRange)
	if err != nil {
		return report, err
	}
	report.Commits = commits
	cache := map[string]historyState{}
	load := func(commit string) (historyState, error) {
		if state, ok := cache[commit]; ok {
			return state, nil
		}
		state, err := historyLoad(root, commit)
		if err == nil {
			cache[commit] = state
		}
		return state, err
	}
	for _, commit := range commits {
		after, err := load(commit)
		if err != nil {
			return model.HistoryReport{}, err
		}
		parentIDs, err := historyParents(root, commit)
		if err != nil {
			return model.HistoryReport{}, err
		}
		parents := make([]historyState, len(parentIDs))
		var diffs []historyDiff
		for i, parentID := range parentIDs {
			parents[i], err = load(parentID)
			if err != nil {
				return model.HistoryReport{}, err
			}
			part, err := historyDiffs(root, parentID, commit)
			if err != nil {
				return model.HistoryReport{}, err
			}
			diffs = append(diffs, part...)
		}
		if len(parentIDs) == 0 {
			diffs, err = historyDiffs(root, "4b825dc642cb6eb9a060e54bf8d69288fbee4904", commit)
			if err != nil {
				return model.HistoryReport{}, err
			}
		}
		if len(parentIDs) > 1 {
			combined, err := historyCombined(root, commit)
			if err != nil {
func Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	nested := appNestedDir(t, root)
	t.Run("plain", func(t *testing.T) {
		status, stdout, stderr := runAppBinary(t, bin, nested, "modules")
		if status != 0 || stderr != "" {
			t.Errorf("status=%d stderr=%q; want success", status, stderr)
		}
		for _, value := range []string{"alpha", "ALP", "packages/alpha", "beta", "BET", "packages/beta", "depends", "api.go", "evals/**", "internal/**", "data/**"} {
			if !strings.Contains(stdout, value) {
				t.Errorf("plain graph %q omits %q", stdout, value)
			}
		}
	})
	t.Run("JSON", func(t *testing.T) {
		status, stdout, stderr := runAppBinary(t, bin, nested, "modules", "--json")
		if status != 0 || stderr != "" {
			t.Errorf("status=%d stderr=%q; want success", status, stderr)
		}
		assertModuleJSON(t, stdout, []string{"alpha", "beta"})
	})
	inner := filepath.Join(root, "workspace", "inner")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: []\n---\n# Gamma\n\n## Specs\n")
	deep := filepath.Join(inner, "packages", "gamma", "deep")
	if err := os.MkdirAll(deep, 0755); err != nil {
		t.Fatal(err)
	}
	t.Run("nearest-ancestor", func(t *testing.T) {
		status, stdout, stderr := runAppBinary(t, bin, deep, "modules")
		if status != 0 || stderr != "" || !strings.Contains(stdout, "gamma") || strings.Contains(stdout, "alpha") || strings.Contains(stdout, "beta") {
			t.Errorf("nearest config: status=%d stdout=%q stderr=%q", status, stdout, stderr)
		}
		status, stdout, stderr = runAppBinary(t, bin, deep, "modules", "--json")
		if status != 0 || stderr != "" {
			t.Errorf("nearest JSON: status=%d stderr=%q", status, stderr)
		}
		assertModuleJSON(t, stdout, []string{"gamma"})
	})
}

func Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	authoringChange(t, root, "0001-first.md", "Users need the first item.\n", "### new:first\nFirst MUST work.\n", "")
	authoringChange(t, root, "0002-second.md", "Users need the second item.\n", "### new:second\nSecond MUST work.\n", "### ALP-001\nAlpha MUST have projected text.\n")
	for _, tc := range []struct {
		name     string
		args     []string
		contains []string
		absent   []string
	}{
		{"current-full", []string{"show", "alpha"}, []string{"# Alpha", "### ALP-001", "Alpha MUST keep its current text."}, []string{"new:first", "new:second", "projected text"}},
		{"current-item", []string{"show", "alpha", "ALP-001"}, []string{"### ALP-001", "Alpha MUST keep its current text."}, []string{"# Alpha", "new:first"}},
		{"projected-full", []string{"show", "alpha", "--projected"}, []string{"# Alpha", "### new:first", "### new:second", "Alpha MUST have projected text."}, []string{"Alpha MUST keep its current text."}},
		{"projected-item", []string{"show", "alpha", "new:second", "--projected"}, []string{"### new:second", "Second MUST work."}, []string{"# Alpha", "new:first"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			dir := appNestedDir(t, root)
			status, plain, stderr := runAppBinary(t, bin, dir, tc.args...)
			if status != 0 || stderr != "" {
				t.Fatalf("show: status=%d stdout=%q stderr=%q", status, plain, stderr)
			}
			for _, token := range tc.contains {
				if !strings.Contains(plain, token) {
					t.Errorf("show omitted %q: %q", token, plain)
				}
			}
			for _, token := range tc.absent {
				if strings.Contains(plain, token) {
					t.Errorf("show unexpectedly included %q: %q", token, plain)
				}
			}
			args := append(append([]string{}, tc.args...), "--json")
			status, output, stderr := runAppBinary(t, bin, dir, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("show JSON: status=%d stdout=%q stderr=%q", status, output, stderr)
			}
			if data := commandStringData(t, output); data != plain {
				t.Errorf("JSON text=%q; want plain text %q", data, plain)
			}
		})
	}
}

func Test_APP_026_ShowNamedErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		args []string
		edit func(*testing.T, string)
	}{
		{"unknown-module", []string{"show", "absent"}, nil},
		{"unknown-item", []string{"show", "alpha", "ALP-999"}, nil},
		{"ambiguous-item", []string{"show", "alpha", "new:shared", "--projected"}, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-first.md", "First need.\n", "### new:shared\nFirst MUST work.\n", "")
			authoringChange(t, root, "0002-second.md", "Second need.\n", "### new:shared\nSecond MUST work.\n", "")
		}},
		{"ambiguous-module", []string{"show", "alpha"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: alpha\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
		}},
		{"missing-module", []string{"show"}, nil},
		{"extra-argument", []string{"show", "alpha", "ALP-001", "extra"}, nil},
		{"invalid-config", []string{"show", "alpha"}, func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			for _, jsonMode := range []bool{false, true} {
				args := append([]string{}, tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				message := commandFailure(t, status, 2, stdout, stderr, jsonMode)
				if tc.name == "unknown-module" && !strings.Contains(message, "absent") ||
					tc.name == "unknown-item" && !strings.Contains(message, "ALP-999") ||
					tc.name == "ambiguous-item" && !strings.Contains(message, "new:shared") {
					t.Errorf("diagnostic %q does not identify %s", message, tc.name)
				}
			}
		})
	}
}
	cmd.Stdout, cmd.Stderr = &out, &errOut
	err := cmd.Run()
	if ctx.Err() != nil {
		t.Fatalf("NEW %v timed out: %v", args, ctx.Err())
	}
	if err == nil {
		return 0, out.String(), errOut.String()
	}
	var exit *exec.ExitError
	if !errors.As(err, &exit) {
		t.Fatalf("NEW %v: %v", args, err)
	}
	return exit.ExitCode(), out.String(), errOut.String()
}

func nextCLIDiagnostic(t *testing.T, status, want int, out, errOut string, jsonMode bool, code string) {
	t.Helper()
	if status != want {
		t.Errorf("exit=%d, want %d; stdout=%q stderr=%q", status, want, out, errOut)
	}
	if jsonMode {
		if errOut != "" {
			t.Errorf("JSON stderr=%q", errOut)
		}
		env := decodeEnvelope(t, out)
		if string(env["data"]) != "null" {
			t.Errorf("failure data=%s", env["data"])
		}
		var ds []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		}
		if err := json.Unmarshal(env["diagnostics"], &ds); err != nil || len(ds) == 0 {
			t.Fatalf("diagnostics=%s: %v", env["diagnostics"], err)
		}
		if ds[0].Code != code || !corrective(ds[0].Message) {
			t.Errorf("diagnostic=%+v; want %s and action", ds[0], code)
		}
	} else if out != "" || !strings.Contains(errOut, code+":") || !corrective(errOut) {
		t.Errorf("plain diagnostic stdout=%q stderr=%q; want %s and action", out, errOut, code)
	}
}

		return result, stop
	}
	state, currentErr := workflow.Current(request.Root, "")
	if currentErr != nil && !strings.Contains(currentErr.Error(), "P02: worktree holds no change") {
		return result, fmt.Errorf("read held change: %w", currentErr)
	}
	if currentErr != nil {
		queue, err := workflow.Eligible(request.Root)
		if err != nil {
			return result, fmt.Errorf("read eligible changes: %w", err)
		}
		if len(queue) == 0 {
			result.Empty = true
			return result, nil
		}
		selected := &queue[0]
		snapshot, err := project.Load(request.Root)
		if err != nil {
			return result, fmt.Errorf("load usage policy: %w", err)
		}
		if _, err := readUsagePolicy(snapshot.Config); err != nil {
			return result, err
		}
		if !request.AllowUnmerged {
			if err := requireApproval(request.Root, selected.Path); err != nil {
				return result, err
			}
		}
		state, err = workflow.Start(request.Root, selected.Path)

exec
/bin/zsh -lc "sed -n '140,225p' packages/app/evals/authoring_commands_test.go; sed -n '245,318p' packages/app/evals/app_test.go; sed -n '196,258p' packages/app/evals/project_commands_test.go; sed -n '1,95p' packages/project/internal/queue/queue.go; sed -n '1,68p' packages/project/internal/history/*.go; sed -n '89,125p' packages/app/evals/show_stale_cli_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 480ms:

func Test_APP_007_APP_024_CheckValidAndAllDiagnostics(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	authoringChange(t, root, "0001-valid.md", "Users need this addition.\n", "### new:valid\nValid MUST work.\n", "")
	for _, jsonMode := range []bool{false, true} {
		args := []string{"check"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
		if status != 0 || stderr != "" {
			t.Errorf("valid check: status=%d stdout=%q stderr=%q", status, stdout, stderr)
		}
		if jsonMode {
			envelope := decodeEnvelope(t, stdout)
			var summary string
			if err := json.Unmarshal(envelope["data"], &summary); err != nil || strings.TrimSpace(summary) == "" {
				t.Error("JSON success needs a summary")
			}
			var diagnostics []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
				t.Errorf("diagnostics=%s; want array: %v", envelope["diagnostics"], err)
			}
			for _, diagnostic := range diagnostics {
				if diagnostic.Code == "" || diagnostic.Message == "" {
					t.Errorf("diagnostic needs code and message: %+v", diagnostic)
				}
			}
		} else if strings.TrimSpace(stdout) == "" {
			t.Error("plain success needs a summary")
		}
	}

	bad := authoringFixture(t)
	authoringChange(t, bad, "0001-bad.md", "", "### new:bad\nA statement without a normative verb.\n", "")
	status, plainOut, plainErr := runAppBinary(t, bin, appNestedDir(t, bad), "check")
	commandFailure(t, status, 1, plainOut, plainErr, false)
	status, jsonOut, jsonErr := runAppBinary(t, bin, appNestedDir(t, bad), "check", "--json")
	commandFailure(t, status, 1, jsonOut, jsonErr, true)
	envelope := decodeEnvelope(t, jsonOut)
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || len(diagnostics) < 2 {
		t.Fatalf("diagnostics=%s; want both independent errors: %v", envelope["diagnostics"], err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == "" || diagnostic.Message == "" || !strings.Contains(plainOut+plainErr, diagnostic.Message) {
			t.Errorf("diagnostic missing code/message or plain equivalent: %+v plain=%q", diagnostic, plainOut+plainErr)
		}
	}
}

func Test_APP_023_APP_024_CheckUsageAndConfiguration(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		args []string
		edit func(*testing.T, string)
	}{
		{"extra-argument", []string{"check", "extra"}, nil},
		{"invalid-config", []string{"check"}, func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			for _, jsonMode := range []bool{false, true} {
				args := append([]string{}, tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				commandFailure(t, status, 2, stdout, stderr, jsonMode)
			}
		})
	}
}

func Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity(t *testing.T) {
func Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	appPut(t, root, "packages/alpha/SPEC.md", authoringSpec)
	authoringChange(t, root, "0001-first.md", "Users need a queued item.\n", "### new:queued\nQueued MUST work.\n", "")
	nested := filepath.Join(root, "packages", "alpha", "nested")
	if err := os.MkdirAll(nested, 0755); err != nil {
		t.Fatal(err)
	}
	binDir := t.TempDir()
	logPath := filepath.Join(t.TempDir(), "invocations")
	recorder := []byte("#!/bin/sh\nprintf '%s\\n' \"$0 $*\" >> \"$FIREWEED_RECORDER\"\nexit 77\n")
	for _, name := range []string{"git", "hook", "codex"} {
		if err := os.WriteFile(filepath.Join(binDir, name), recorder, 0700); err != nil {
			t.Fatal(err)
		}
	}
	t.Setenv("FIREWEED_RECORDER", logPath)
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))
	for _, tc := range informationCases(t) {
		status, stdout, stderr := runAppBinary(t, bin, nested, tc.args...)
		if status != 0 {
			t.Errorf("%s status=%d stderr=%q", tc.name, status, stderr)
		}
		assertInformationOutput(t, stdout, stderr, tc.want)
	}
	runAppBinary(t, bin, nested, "modules", "--json")
	for _, args := range [][]string{
		{"list"}, {"list", "--json"},
		{"show", "alpha"}, {"show", "alpha", "ALP-001"},
		{"show", "alpha", "ALP-001", "--history"}, {"show", "alpha", "ALP-001", "--history", "--json"},
		{"show", "alpha", "--projected"}, {"show", "alpha", "new:queued", "--projected"},
	} {
		runAppBinary(t, bin, nested, args...)
	}
	if calls, err := os.ReadFile(logPath); err == nil {
		t.Errorf("read commands invoked external tools: %s", calls)
	} else if !os.IsNotExist(err) {
		t.Fatal(err)
	}
	// Set up the repository with real Git before recording status reads.
	gitPath := filepath.Join(binDir, "git")
	if err := os.Remove(gitPath); err != nil {
		t.Fatal(err)
	}
	realGit, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "read command fixture")
	// Status may read Git. This recorder forwards every invocation without
	// inspecting its arguments; the existing hook and agent sentinels remain.
	gitRecorder := []byte("#!/bin/sh\nprintf '%s\\n' \"$0 $*\" >> \"$FIREWEED_RECORDER\"\nexec \"$FIREWEED_REAL_GIT\" \"$@\"\n")
	if err := os.WriteFile(filepath.Join(binDir, "git"), gitRecorder, 0700); err != nil {
		t.Fatal(err)
	}
	t.Setenv("FIREWEED_REAL_GIT", realGit)
	for _, args := range [][]string{{"status"}, {"status", "--json"}, {"pr-body", "HEAD..HEAD"}, {"pr-body", "HEAD..HEAD", "--json"}} {
		runAppBinary(t, bin, nested, args...)
	}
	calls, err := os.ReadFile(logPath)
	if err != nil && !os.IsNotExist(err) {
		t.Fatal(err)
	}
	for _, line := range strings.Split(string(calls), "\n") {
		if strings.Contains(line, "/hook ") || strings.Contains(line, "/codex ") ||
			strings.Contains(" "+line+" ", " fetch ") || strings.Contains(" "+line+" ", " push ") {
			t.Errorf("status read invoked hook, agent, fetch or push: %q", line)
		}
	}
}

func run(args ...string) (int, string, string) {
func Test_APP_018_ModulesConfigurationErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name string
		code string
		edit func(*testing.T, string)
	}{
		{"missing", "C01", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed", "C01", func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"invalid-graph", "C02", func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
		{"invalid-document", "C03", func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing\n---\n# Beta\n\n## Specs\n")
		}},
		{"invalid-metadata", "C12", func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: invalid\n---\n# Beta\n\n## Specs\n")
		}},
	} {
		for _, jsonMode := range []bool{false, true} {
			format := "plain"
			if jsonMode {
				format = "JSON"
			}
			t.Run(tc.name+"/"+format, func(t *testing.T) {
				root := appProjectFixture(t)
				tc.edit(t, root)
				args := []string{"modules"}
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				if status != 2 {
					t.Errorf("status=%d; want 2", status)
				}
				if jsonMode {
					if stderr != "" {
						t.Errorf("JSON stderr=%q; want empty", stderr)
					}
					diagnostic := assertFailedJSON(t, stdout, tc.code)
					var result struct {
						Diagnostics []struct {
							Code string `json:"code"`
						} `json:"diagnostics"`
					}
					if err := json.Unmarshal([]byte(stdout), &result); err != nil || len(result.Diagnostics) == 0 || result.Diagnostics[0].Code != tc.code {
						t.Errorf("JSON diagnostic needs %s: %s", tc.code, stdout)
					}
					if !corrective(diagnostic) {
						t.Errorf("diagnostic %q needs correction", diagnostic)
					}
				} else if !strings.HasPrefix(strings.TrimSpace(stdout+stderr), tc.code+":") || !corrective(stdout+stderr) {
					t.Errorf("plain diagnostic=%q; want corrective %s", stdout+stderr, tc.code)
				}
			})
		}
	}
}

package queue

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var changeName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

func Queue(root string) ([]model.Change, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	return readQueue(root, snapshot)
}

func readQueue(root string, snapshot model.Snapshot) ([]model.Change, error) {
	known := map[string]bool{}
	for _, module := range snapshot.Modules {
		known[module.Name] = true
	}
	entries, err := os.ReadDir(filepath.Join(root, "changes"))
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, withCause("C03", fmt.Errorf("read changes: %w", err))
	}
	var changes []model.Change
	slugs := map[string]bool{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		parts := changeName.FindStringSubmatch(entry.Name())
		if parts == nil {
			continue
		}
		number, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, withCause("C13", fmt.Errorf("invalid change number %s: %w", entry.Name(), err))
		}
		if slugs[parts[2]] {
			return nil, withCause("C13", fmt.Errorf("ambiguous change slug %q", parts[2]))
		}
		slugs[parts[2]] = true
		path := "changes/" + entry.Name()
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
		if err != nil {
			return nil, withCause("C03", fmt.Errorf("read %s: %w", path, err))
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return nil, withCause("C03", fmt.Errorf("parse %s: %w", path, err))
		}
		module, ok := doc.Frontmatter["module"].(string)
		if !ok || !known[module] {
			return nil, withCause("C07", fmt.Errorf("%s names unknown module %q", path, module))
		}
		change := model.Change{Number: number, Slug: parts[2], Path: path, Module: module, Document: doc}
		for _, section := range []struct {
			name string
			into *[]contracts.Item
		}{
			{"Added", &change.Added}, {"Modified", &change.Modified}, {"Removed", &change.Removed},
		} {
			body, _ := contracts.Section(doc.Body, section.name)
			items, err := contracts.ParseItems(body)
			if err != nil {
				return nil, withCause(itemErrorCode(err), fmt.Errorf("%s %s: %w", path, section.name, err))
			}
			*section.into = items
		}
		changes = append(changes, change)
	}
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].Number != changes[j].Number {
			return changes[i].Number < changes[j].Number
		}
		return changes[i].Slug < changes[j].Slug
	})
	return changes, nil
}

func resolve(changes []model.Change, target string) (int, error) {
package history

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fixture.LoadCode(err), err)
	}

	var target *model.Module
	for i := range snapshot.Modules {
		if snapshot.Modules[i].Name == module {
			target = &snapshot.Modules[i]
			break
		}
	}
	if target == nil {
		return nil, fmt.Errorf("U02: unknown module %q", module)
	}

	result := make([]model.SpecHistoryEntry, 0)
	archiveDir := filepath.Join(root, "changes", "archive")
	files, err := os.ReadDir(archiveDir)
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("C03: read changes/archive: %w", err)
	}
	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".md" {
			continue
		}
		path := "changes/archive/" + file.Name()
		entry, err := readArchive(root, path)
		if err != nil {
			return nil, err
		}
		if entry.Module != module {
			continue
		}
		if id != "" {
			if !contains(entry.IDs, id) {
				continue
			}
			entry.IDs = []string{id}
		} else if len(entry.IDs) == 0 {
			continue
		}
		result = append(result, entry)
	}

	if id != "" && len(result) == 0 {
		body, _ := contracts.Section(target.Document.Body, "Specs")
		items, err := contracts.ParseItems(body)
		if err != nil {
			return nil, fmt.Errorf("C03: %s/SPEC.md Specs: %w", target.Dir, err)
		}
		found := false
	for _, alias := range []string{"0002-target.md", path} {
		if got := staleCLISuccess(t, bin, dir, alias); got != plain {
			t.Errorf("alias %q differs from slug output", alias)
		}
	}
	first, second := strings.Index(plain, "ALP-001"), strings.Index(plain, "ALP-002")
	if first < 0 || second <= first {
		t.Fatalf("stale entries are missing or out of order: %q", plain)
	}
	recovered, fallback := plain[first:second], plain[second:]
	for _, marker := range []string{"-### ALP-001 [review]", "-One MUST use old text.  ", "+### ALP-001 [quarantine: changed]", "+One MUST use current text."} {
		if !strings.Contains(recovered, marker) {
			t.Errorf("recovered diff lacks %q: %q", marker, recovered)
		}
	}
	if strings.Contains(recovered, "Target MUST propose different text") || strings.Contains(recovered, "deadbeef") {
		t.Errorf("recovered diff used proposed edit or another item's hash: %q", recovered)
	}
	for _, marker := range []string{"deadbeef", "Two MUST use current text.", "unavailable"} {
		if !strings.Contains(strings.ToLower(fallback), strings.ToLower(marker)) {
			t.Errorf("unrecovered entry lacks %q: %q", marker, fallback)
		}
	}
	if strings.Contains(fallback, "Two MUST use old text") || strings.Contains(fallback, "-Two MUST") || strings.Contains(fallback, "+Two MUST") {
		t.Errorf("unverified history appeared as a diff: %q", fallback)
	}
	for file, before := range beforeFiles {
		if !bytes.Equal(before, staleCLIText(t, root, file)) {
			t.Errorf("show changed %s", file)
		}
	}
	if cliGit(t, root, "rev-parse", "HEAD") != head || cliGit(t, root, "for-each-ref", "--format=%(refname) %(objectname)") != refs {
		t.Error("show changed HEAD or refs")
	}
}

func Test_APP_144_StaleUnavailableHistoryAndEmptyView(t *testing.T) {

codex
The controller has no held `app` change in this reconstruction workspace, so `fireweed context app` returned U02. I’ll use the retained spec, API and evals for this drill.
exec
/bin/zsh -lc "mkdir -p packages/app/internal/entry
cat > packages/app/internal/entry/entry.go <<'EOF'
package entry

import (
	\"context\"
	\"encoding/json\"
	\"fmt\"
	\"io\"
	\"os\"
	\"path/filepath\"
	\"regexp\"
	\"slices\"
	\"strconv\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/assets\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
)

type command struct {
	name string
	args []string
	flags map[string]string
	json bool
}

func (c command) has(flag string) bool { _, ok := c.flags[flag]; return ok }
func (c command) arg(i int) string {
	if i < len(c.args) { return c.args[i] }
	return \"\"
}

type syntax struct {
	min, max int
	usage string
	flags string
}

var commands = map[string]syntax{
	\"help\": {}, \"version\": {},
	\"explain\": {1,1,\"<code>\",\"\"}, \"playbook\": {1,1,\"<name>\",\"\"},
	\"modules\": {}, \"new-change\": {2,2,\"<module> <slug>\",\"\"},
	\"stamp\": {1,1,\"<change>\",\"\"}, \"check\": {0,0,\"[--fix]\",\"--fix\"},
	\"show\": {1,2,\"<module> [ID] [--history | --projected] or fireweed show --stale <change>\",\"--history --projected --stale\"},
	\"start\": {1,1,\"<change> [--from-attempt N]\",\"--from-attempt\"},
	\"context\": {1,1,\"<change> [--full]\",\"--full\"},
	\"verify\": {0,1,\"[change] [--quick]\",\"--quick\"}, \"done\": {0,1,\"[change]\",\"\"},
	\"coverage\": {1,1,\"<module> [--with-change <change>]\",\"--with-change\"},
	\"gc\": {}, \"migrate\": {},
	\"init\": {0,0,\"[--recipe NAME] [--agent COMMAND] [--audit-agent COMMAND] [--no-app]\",\"--recipe --agent --audit-agent --no-app\"},
	\"doctor\": {0,0,\"[--no-agent]\",\"--no-agent\"},
	\"agents-md\": {0,0,\"--write\",\"--write\"},
	\"next\": {0,0,\"[--allow-unmerged]\",\"--allow-unmerged\"},
	\"run\": {0,0,\"[--max N] [--for D] [--allow-unmerged]\",\"--max --for --allow-unmerged\"},
	\"release\": {1,1,\"<change>\",\"\"}, \"commits\": {1,1,\"<change>\",\"\"}, \"abandon\": {1,1,\"<change>\",\"\"},
	\"new-module\": {1,1,\"<name> --prefix <P>\",\"--prefix\"},
	\"check-commits\": {0,1,\"[range]\",\"\"}, \"ci\": {0,1,\"[range] [--no-doctor]\",\"--no-doctor\"},
	\"list\": {}, \"status\": {}, \"pr-body\": {1,1,\"<range>\",\"\"}, \"drill\": {1,1,\"<module>\",\"\"},
}

var valueFlags = strings.Fields(\"--from-attempt --with-change --recipe --agent --audit-agent --max --for --prefix\")

func parse(args []string) (command, error) {
	c := command{name: \"help\", json: slices.Contains(args, \"--json\"), flags: map[string]string{}}
	for len(args)>0 && args[0]==\"--json\" { args=args[1:] }
	if len(args)==0 { return c,nil }
	c.name, args = args[0], args[1:]
	switch c.name {
	case \"--help\", \"-h\": c.name=\"help\"
	case \"--version\": c.name=\"version\"
	}
	s, ok := commands[c.name]
	if "'!ok { return c, fmt.Errorf("unknown command %q. Use fireweed help for valid commands",c.name) }
	usage := strings.TrimSpace("fireweed "+c.name+" "+s.usage)
	if slices.Contains(args,"--stale") { usage="fireweed show --stale <change>" }
	if slices.Contains(args,"--fix") { usage="fireweed check --fix" }
	bad := func(message string, params ...any) (command,error) {
		msg:=fmt.Sprintf(message,params...)
		if c.name=="init" { msg+=". Valid recipes: "+recipeChoices+"; valid agents: codex, claude, or a command containing {prompt} or {prompt_file}" }
		return c,fmt.Errorf("%s. Usage: %s",msg,usage)
	}
	for i:=0;i<len(args);i++ {
		arg:=args[i]
		if arg=="--json" { continue }
		if !strings.HasPrefix(arg,"-") {
			if strings.TrimSpace(arg)=="" { return bad("empty argument") }
			c.args=append(c.args,arg)
			continue
		}
		if !slices.Contains(strings.Fields(s.flags),arg) { return bad("unknown option %q for %s",arg,c.name) }
		if c.has(arg) { return bad("repeated option %s",arg) }
		value:=""
		if slices.Contains(valueFlags,arg) {
			if i+1==len(args) || strings.HasPrefix(args[i+1],"--") || strings.TrimSpace(args[i+1])=="" { return bad("%s requires a value",arg) }
			i++
			value=args[i]
			switch arg {
			case "--max","--from-attempt":
				if n,e:=strconv.Atoi(value); e!=nil || n<=0 { return bad("%s requires a positive integer, got %q",arg,value) }
			case "--for":
				if d,e:=time.ParseDuration(value); e!=nil || d<=0 { return bad("%s requires a positive Go duration, got %q",arg,value) }
			}
		}
		c.flags[arg]=value
	}
	if len(c.args)<s.min || len(c.args)>s.max { return bad("invalid arguments for %s",c.name) }
	if c.has("--stale") && (len(c.args)!=1 || c.has("--history") || c.has("--projected")) { return bad("--stale requires one change and excludes --history and --projected") }
	if c.has("--history") && c.has("--projected") { return bad("--history and --projected cannot be combined") }
	if c.name=="agents-md" && !c.has("--write") { return bad("agents-md requires --write") }
	if c.name=="new-module" && !c.has("--prefix") { return bad("new-module requires --prefix") }
	return c,nil
}

type diagnostic struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Path string `json:"path,omitempty"`
	Warning bool `json:"warning"`
}

type response struct {
	exit int
	data any
	plain string
	diagnostics []diagnostic
}

func stringResult(s string) response { return response{data:s,plain:s} }
func confirmation(s string) response { return stringResult(s+"'"\\n\") }

func corrective(code, message string) string {
	if help,err:=assets.Explain(code); err==nil { return message+\". \"+help }
	return message+\". Correct the reported problem and retry.\"
}

func reportDiagnostics(ds []project.Diagnostic, guide bool) []diagnostic {
	result:=make([]diagnostic,0,len(ds))
	for _,d:=range ds {
		message:=d.Message
		if guide && "'!d.Warning { message=corrective(d.Code,message) }
		result=append(result,diagnostic{d.Code,message,d.Path,d.Warning})
	}
	return result
}

func hasErrors(ds []project.Diagnostic) bool {
	return slices.ContainsFunc(ds,func(d project.Diagnostic) bool {return !d.Warning})
}

var causePattern = regexp.MustCompile(`(?:''^|['"\\s:])([A-Z][0-9]{2}|P2-01):\\s+"'`)

func cause(err error, fallback string) (string,string) {
	message:=err.Error()
	if code,_,ok:=contracts.ErrorCode(err); ok { return code,strings.TrimPrefix(message,code+": ") }
	if match:=causePattern.FindStringSubmatch(message); match!=nil { return match[1],strings.TrimPrefix(message,match[1]+": ") }
	return fallback,message
}

func reject(exit int, fallback string, err error) response {
	code,message:=cause(err,fallback)
	return response{exit:exit,diagnostics:[]diagnostic{{Code:code,Message:corrective(code,message)}}}
}

func usageError(message string) response {
	return response{exit:2,diagnostics:[]diagnostic{{Code:"U01",Message:message}}}
}

func discover() (string,error) {
	dir,err:=os.Getwd()
	if err!=nil { return "",err }
	for {
		_,err:=os.Lstat(filepath.Join(dir,"fireweed.yaml"))
		if err==nil { return dir,nil }
		if !os.IsNotExist(err) { return "",err }
		parent:=filepath.Dir(dir)
		if parent==dir { return "",fmt.Errorf("no fireweed.yaml found in this directory or its ancestors") }
		dir=parent
	}
}

func loadFailure(root string, err error) response {
	ds:=project.Check(root)
	if len(ds)>0 { return response{exit:2, diagnostics:reportDiagnostics(ds[:1],true)} }
	return reject(2,"C01",err)
}

func Run(args []string, out, errOut io.Writer) int {
	c,err:=parse(args)
	var r response
	if err!=nil { r=usageError(err.Error()) } else { r=execute(c) }
	if c.json {
		if r.diagnostics==nil { r.diagnostics=[]diagnostic{} }
		_ = json.NewEncoder(out).Encode(struct {
			Schema int `json:"schema"`
			Data any `json:"data"`
			Diagnostics []diagnostic `json:"diagnostics"`
		}{1,r.data,r.diagnostics})
	} else {
		fmt.Fprint(out,r.plain)
		for _,d:=range r.diagnostics {
			fmt.Fprintf(errOut,"%s: ",d.Code)
			if d.Warning { fmt.Fprint(errOut,"warning: ") }
			if d.Path!="" { fmt.Fprintf(errOut,"%s: ",d.Path) }
			fmt.Fprintln(errOut,d.Message)
		}
	}
	return r.exit
}

func execute(c command) response {
	switch c.name {
	case "help": return stringResult(assets.Help())
	case "version": return confirmation(assets.Version())
	case "explain","playbook":
		lookup:=assets.Explain
		if c.name=="playbook" { lookup=assets.Playbook }
		s,err:=lookup(c.arg(0))
		if err!=nil { return reject(2,"U02",err) }
		if !strings.HasSuffix(s,"'"\\n\") { s+=\"\\n\" }
		return stringResult(s)
	case \"init\":
		return initialize(c)
	}
	root,err:=discover()
	if err"'!=nil { return reject(2,"C01",err) }
	if c.name=="pr-body" { return prBody(root,c.arg(0)) }

	snap,err:=project.Load(root)
	if err!=nil {
		failure:=loadFailure(root,err)
		if (c.name=="check" || c.name=="doctor" || c.name=="ci") && failure.diagnostics[0].Code!="C01" {
			failure.exit=1
		}
		return failure
	}
	ctx:=context.Background()
	bin,err:=os.Executable()
	if err!=nil { return reject(2,"P01",err) }
	switch c.name {
	case "modules","new-change","stamp","check","show","coverage","list","migrate":
		return projectCommand(ctx,c,snap,bin)
	case "start","context","verify","done","release","commits","abandon","gc":
		return workflowCommand(ctx,c,root,bin)
	case "next","run","drill":
		return runnerCommand(ctx,c,snap,bin)
	case "doctor","check-commits","ci":
		return checksCommand(ctx,c,snap,bin)
	case "agents-md":
		return writeAgents(root)
	case "new-module":
		return newModule(c,snap)
	case "status":
		return status(root)
	}
	panic("unhandled command "+c.name)
}
EOF
sed -n '"'87,116p' packages/app/evals/app_test.go; sed -n '95,128p' packages/project/internal/queue/queue.go; sed -n '145,240p' packages/app/evals/migrate_cli_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 662ms:
func Test_APP_002_InformationJSON(t *testing.T) {
	for _, tc := range informationCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			plainStatus, plain, plainErr := run(tc.args...)
			if plainStatus != 0 || plainErr != "" {
				t.Errorf("plain status=%d stderr=%q; want 0 and empty", plainStatus, plainErr)
			}
			assertInformationOutput(t, plain, plainErr, tc.want)
			args := append(append([]string{}, tc.args...), "--json")
			status, stdout, stderr := run(args...)
			if status != 0 || stderr != "" {
				t.Errorf("status=%d stderr=%q; want 0 and empty", status, stderr)
			}
			envelope := decodeEnvelope(t, stdout)
			var data string
			if err := json.Unmarshal(envelope["data"], &data); err != nil {
				t.Fatalf("data must be a string: %v", err)
			}
			if data != plain {
				t.Errorf("data=%q; want plain result %q", data, plain)
			}
			var diagnostics []json.RawMessage
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
				t.Errorf("diagnostics=%s; want empty array", envelope["diagnostics"])
			}
		})
	}
}

func Test_APP_003_InvalidRequests(t *testing.T) {
func resolve(changes []model.Change, target string) (int, error) {
	for i, change := range changes {
		if target == change.Slug || target == filepath.Base(change.Path) || target == change.Path {
			return i, nil
		}
	}
	slugs := make([]string, len(changes))
	for i, change := range changes {
		slugs[i] = change.Slug
	}
	guidance := "no pending changes"
	if len(slugs) > 0 {
		guidance = "valid pending change slugs: " + strings.Join(slugs, ", ")
	}
	return -1, contracts.CodedError{Code: "U02", Subject: target, Cause: fmt.Errorf("pending change %q not found; %s", target, guidance)}
}
	if status != 0 || stderr != "" || commandStringData(t, output) != plain {
		t.Errorf("empty JSON migrate: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	migrateUnchanged(t, root, before)
}

func Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code string
		status     int
		edit       func(*testing.T, string)
	}{
		{"static-check", "C04", 1, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "Alpha MUST keep its current text.", "Alpha has no normative verb.", 1))
		}},
		{"malformed-config", "C01", 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"malformed-module", "C03", 2, func(t *testing.T, root string) {
			appPut(t, root, "packages/alpha/SPEC.md", strings.Replace(authoringSpec, "### ALP-001", "### invalid heading", 1))
		}},
		{"unsupported-version", "C01", 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "version: 1", "version: 2", 1))
		}},
		{"missing-version", "C01", 2, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "version: 1\n", "", 1))
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := migrateProject(t)
			tc.edit(t, root)
			before := migrateCapture(t, root)
			var expected []project.Diagnostic
			if tc.name == "static-check" {
				expected = project.Check(root)
				if len(expected) == 0 || expected[0].Code != tc.code || expected[0].Warning {
					t.Fatalf("static fixture did not produce %s: %+v", tc.code, expected)
				}
			}
			for _, jsonMode := range []bool{false, true} {
				args := []string{"migrate"}
				if jsonMode {
					args = append(args, "--json")
				}
				status, out, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				if status != tc.status {
					t.Errorf("%s JSON=%v: status=%d stdout=%q stderr=%q; want %d", tc.name, jsonMode, status, out, stderr, tc.status)
				}
				if jsonMode {
					if stderr != "" {
						t.Errorf("JSON stderr=%q", stderr)
					}
					if string(decodeEnvelope(t, out)["data"]) != "null" {
						t.Errorf("failure data=%s; want null", decodeEnvelope(t, out)["data"])
					}
					got := migrateDiagnostics(t, out)
					if len(got) == 0 || got[0].Code != tc.code {
						t.Errorf("JSON diagnostics=%+v; want %s", got, tc.code)
					}
					for i, want := range expected {
						if i >= len(got) || got[i].Code != want.Code || !strings.Contains(got[i].Message, want.Message) || got[i].Warning != want.Warning || got[i].Path != want.Path {
							t.Errorf("static diagnostic %d=%+v; want cause %+v", i, got, want)
						}
					}
					for _, d := range got {
						if !corrective(d.Message) {
							t.Errorf("diagnostic lacks corrective guidance: %+v", d)
						}
					}
				} else if out != "" || !strings.Contains(stderr, tc.code+":") || !corrective(stderr) {
					t.Errorf("plain failure stdout=%q stderr=%q; want corrective %s", out, stderr, tc.code)
				}
				migrateUnchanged(t, root, before)
			}
		})
	}
	for _, jsonMode := range []bool{false, true} {
		args := []string{"migrate"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, out, stderr := runAppBinary(t, bin, t.TempDir(), args...)
		message := commandFailure(t, status, 2, out, stderr, jsonMode)
		if jsonMode {
			message = assertFailedJSON(t, out, "C01")
		} else if !strings.HasPrefix(strings.TrimSpace(stderr), "C01:") {
			t.Errorf("discovery stderr=%q; want C01", stderr)
		}
		if !corrective(message) {
			t.Errorf("discovery diagnostic lacks guidance: %q", message)
		}
	}
}

func Test_APP_155_MigrateRejectsArgumentsBeforeMutation(t *testing.T) {

exec
/bin/zsh -lc "sed -n '130,180p' packages/project/internal/queue/project.go; rg -n 'func Require|func Evaluate|quarantine' packages/checks/internal/engine/*.go; sed -n '1,125p' packages/app/evals/list_cli_test.go; rg -n 'Contains|== \"' packages/app/evals/{release_cli_test,abandon_cli_test,gc_cli_test,pr_body_cli_test,usage_cli_test,run_for_cli_test}.go | head -105; rg -n 'func .*Approved|func .*claim|func .*ref' packages/workflow/internal/lifecycle/stub.go | head -30" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 482ms:
			case "disposable":
				module.Zones.Disposable = list
			case "preserve":
				module.Zones.Preserve = list
			case "append_only":
				module.Zones.AppendOnly = list
			}
		case "boundaries_opt_out", "hooks":
		default:
			return fmt.Errorf("set cannot change %q", key)
		}
		front[key] = value
	}
	module.Document.Frontmatter = front
	return nil
}

func asStrings(value any) ([]string, error) {
	values, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("expected a string list")
	}
	result := make([]string, 0, len(values))
	for _, value := range values {
		text, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("expected a string list")
		}
		result = append(result, text)
	}
	return result, nil
}

func (p *projection) result() model.Snapshot {
	for i := range p.snapshot.Modules {
		module := &p.snapshot.Modules[i]
		var body strings.Builder
		body.WriteByte('\n')
		for j, entry := range p.items[module.Name] {
			if j > 0 && !strings.HasSuffix(body.String(), "\n\n") {
				body.WriteByte('\n')
			}
			body.WriteString(contracts.RenderItem(entry.item))
			if !strings.HasSuffix(body.String(), "\n") {
				body.WriteByte('\n')
			}
		}
		module.Document.Body = contracts.ReplaceSection(module.Document.Body, "Specs", body.String())
	}
	return p.snapshot
}
packages/checks/internal/engine/hooks.go:105:func Evaluate(ctx context.Context, request model.HookRequest) (model.Report, error) {
packages/checks/internal/engine/junit.go:149:func RequireCoverage(report model.Report, specs []contracts.Item) error {
packages/checks/internal/engine/junit.go:151:		if exempt(spec.Tag, "review") || exempt(spec.Tag, "quarantine") {
packages/checks/internal/engine/junit.go:168:func RequireFailFirst(report model.Report, ids, already []string) error {
packages/checks/internal/engine/junit.go:193:func RequirePass(report model.Report, specs []contracts.Item) error {
packages/checks/internal/engine/junit.go:195:		if exempt(spec.Tag, "quarantine") {
package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func listRun(t *testing.T, dir string, args ...string) (int, string, string) {
	t.Helper()
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(previous); err != nil {
			t.Fatalf("restore working directory: %v", err)
		}
	}()
	return run(args...)
}

func listChange(t *testing.T, root, filename, module, front string) {
	t.Helper()
	appPut(t, root, "changes/"+filename, "---\nmodule: "+module+"\n"+front+"---\n\n## Motivation\nUsers need this queued change.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
}

func listJSON(t *testing.T, output string) []map[string]json.RawMessage {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want empty array: %v", envelope["diagnostics"], err)
	}
	var entries []map[string]json.RawMessage
	if err := json.Unmarshal(envelope["data"], &entries); err != nil || entries == nil {
		t.Fatalf("data=%s; want array: %v", envelope["data"], err)
	}
	return entries
}

func Test_APP_114_APP_115_APP_117_ListQueueOrderAndFields(t *testing.T) {
	root := appProjectFixture(t)
	listChange(t, root, "0002-zed.md", "alpha", "")
	listChange(t, root, "0001-first.md", "beta", "phase: evals\naudit_rounds: 1\n")
	listChange(t, root, "0002-ace.md", "alpha", "phase: audit\nblocked: review required\naudit_rounds: 2\n")
	want := []struct {
		slug, path, module, phase, blocked string
		rounds                             int
	}{
		{"first", "changes/0001-first.md", "beta", "evals", "", 1},
		{"ace", "changes/0002-ace.md", "alpha", "audit", "review required", 2},
		{"zed", "changes/0002-zed.md", "alpha", "queued", "", 0},
	}
	queued, err := project.Queue(root)
	if err != nil || len(queued) != len(want) {
		t.Fatalf("public Queue fixture: %v, entries=%v", err, queued)
	}
	for i, entry := range queued {
		if entry.Slug != want[i].slug || entry.Path != want[i].path || entry.Module != want[i].module {
			t.Fatalf("public Queue order[%d]=%+v", i, entry)
		}
	}
	status, plain, stderr := listRun(t, appNestedDir(t, root), "list")
	if status != 0 || stderr != "" {
		t.Fatalf("plain list exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	previousPath := -1
	for _, entry := range want {
		for _, field := range []string{entry.slug, entry.path, entry.module, entry.phase} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain list=%q missing %q", plain, field)
			}
		}
		if strings.Count(plain, entry.slug) < 2 {
			t.Errorf("plain list=%q must identify slug %q separately from its path", plain, entry.slug)
		}
		position := strings.Index(plain, entry.path)
		if position <= previousPath {
			t.Errorf("plain path %q at %d after previous at %d; output=%q", entry.path, position, previousPath, plain)
		}
		previousPath = position
		if entry.blocked != "" && !strings.Contains(plain, entry.blocked) {
			t.Errorf("plain list=%q missing blocked reason %q", plain, entry.blocked)
		}
		if entry.blocked == "" {
			found := false
			for _, row := range strings.Split(plain, "\n") {
				if strings.Contains(row, entry.path) && strings.Contains(row, `blocked=""`) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("plain list=%q missing blocked=\"\" on row for %q", plain, entry.path)
			}
		}
		if !regexp.MustCompile(`\b` + strconv.Itoa(entry.rounds) + `\b`).MatchString(plain) {
			t.Errorf("plain list=%q missing audit rounds %d", plain, entry.rounds)
		}
	}
	status, output, stderr := listRun(t, root, "list", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("JSON list exit=%d stdout=%q stderr=%q", status, output, stderr)
	}
	entries := listJSON(t, output)
	if len(entries) != len(want) {
		t.Fatalf("JSON entries=%s; want %d", output, len(want))
	}
	for i, entry := range want {
		fields := entries[i]
		for key, expected := range map[string]any{
			"slug": entry.slug, "path": entry.path, "module": entry.module,
			"phase": entry.phase, "blocked": entry.blocked, "audit_rounds": float64(entry.rounds),
		} {
packages/app/evals/run_for_cli_test.go:21:		if !strings.Contains(script, selection) {
packages/app/evals/run_for_cli_test.go:27:	if !strings.Contains(script, anchor) {
packages/app/evals/run_for_cli_test.go:49:	if strings.TrimSpace(reason) == "" || !strings.Contains(reason, limit) || !regexp.MustCompile(`(?i)time|duration|deadline`).MatchString(reason) {
packages/app/evals/run_for_cli_test.go:74:	if queued != original || !strings.Contains(original, "run: "+held.Run) || !strings.Contains(original, "phase: evals") || !strings.Contains(original, "### BET-002") {
packages/app/evals/run_for_cli_test.go:78:	if claim == "" || claim != held.ClaimOID {
packages/app/evals/run_for_cli_test.go:82:	if !strings.Contains(partial, "TIME_PARTIAL_WORK") {
packages/app/evals/run_for_cli_test.go:135:	if err := json.Unmarshal(decodeEnvelope(t, out)["data"], &resumed); err != nil || !strings.Contains(resumed.Change, "beta") || resumed.Run != held.Run || resumed.Phase != "applied" || resumed.Empty {
packages/app/evals/run_for_cli_test.go:138:	if got := string(mustReadNextCLI(t, filepath.Join(f.logs, "resumed-beta.md"))); !strings.Contains(got, "run: "+held.Run) || !strings.Contains(got, "phase: evals") || !strings.Contains(got, "### BET-002") {
packages/app/evals/run_for_cli_test.go:146:		if !strings.Contains(archive, marker) {
packages/app/evals/run_for_cli_test.go:173:	if !strings.Contains(script, anchor) {
packages/app/evals/run_for_cli_test.go:193:	if err != nil || state.Phase != "evals" || state.Run == "" || state.ClaimOID == "" || len(state.Change.Added) != 1 || state.Change.Added[0].ID != "ALP-002" {
packages/app/evals/run_for_cli_test.go:199:	if work := string(mustReadNextCLI(t, filepath.Join(f.root, "packages/alpha/evals/partial_test.go"))); !strings.Contains(work, "TIME_PARTIAL_WORK") {
packages/app/evals/run_for_cli_test.go:230:		if !strings.Contains(diagnostics[0].Message, marker) {
packages/app/evals/run_for_cli_test.go:236:	if !strings.Contains(alpha, "blocked:") || !strings.Contains(alpha, "phase: evals") || !strings.Contains(alpha, "### ALP-002") {
packages/app/evals/run_for_cli_test.go:240:	if saved := cliGit(t, f.root, "for-each-ref", "--format=%(objectname)", ref); saved == "" || cliGit(t, f.root, "cat-file", "-t", saved) != "commit" {
packages/app/evals/pr_body_cli_test.go:66:	if !strings.Contains(report.Body, "0001-review") || !strings.Contains(report.Body, "opaque.dat") || !strings.Contains(report.Body, "SAVED B01") || strings.Contains(report.Body, "later HEAD") {
packages/app/evals/pr_body_cli_test.go:74:				if code == "B04" && d.Path != "" && d.Path != "packages/alpha/internal/opaque.dat" {
packages/app/evals/pr_body_cli_test.go:92:	if !strings.Contains(plain, report.Body) {
packages/app/evals/pr_body_cli_test.go:96:		if !strings.Contains(plain+plainErr, diagnostic.Code) || !strings.Contains(plain+plainErr, diagnostic.Message) {
packages/app/evals/pr_body_cli_test.go:140:	if strings.TrimSpace(report.Body) == "" || !(strings.Contains(lowerBody, "none") || strings.Contains(lowerBody, "no change") || strings.Contains(lowerBody, "no selected changes")) {
packages/app/evals/pr_body_cli_test.go:157:	if _, err := workflow.PRBody(badConfig, badRange); err == nil || !strings.Contains(err.Error(), "fireweed.yaml") {
packages/app/evals/pr_body_cli_test.go:161:	if _, err := workflow.PRBody(root, unavailableRange); err == nil || !strings.Contains(err.Error(), "P01") {
packages/app/evals/pr_body_cli_test.go:189:				if !corrective(message) || (tc.cause != "" && !strings.Contains(message, tc.cause)) {
packages/app/evals/gc_cli_test.go:37:	if out == "" {
packages/app/evals/gc_cli_test.go:144:			} else if !strings.Contains(strings.ToLower(out), "no refs") || !strings.Contains(strings.ToLower(out), "pruned") {
packages/app/evals/gc_cli_test.go:153:	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
packages/app/evals/release_cli_test.go:41:			if err != nil || started.ClaimOID == "" || started.Run == "" || started.Phase != "evals" || len(started.Change.Added) != 1 {
packages/app/evals/release_cli_test.go:49:			if !strings.Contains(string(change), "run: "+started.Run) || !strings.Contains(string(change), "phase: evals") || !strings.Contains(string(change), "### "+started.Change.Added[0].ID) {
packages/app/evals/release_cli_test.go:61:			if !strings.Contains(beforeRefs, claimLine) {
packages/app/evals/release_cli_test.go:76:			if !strings.Contains(strings.ToLower(confirmation), "releas") || !strings.Contains(confirmation, "manual") {
packages/app/evals/release_cli_test.go:116:	if err != nil || owner.ClaimOID == "" {
packages/app/evals/abandon_cli_test.go:17:	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
packages/app/evals/abandon_cli_test.go:69:				if err != nil || started.ClaimOID == "" || started.Phase != "evals" {
packages/app/evals/abandon_cli_test.go:77:				if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID == "" || current.Phase != "impl" {
packages/app/evals/abandon_cli_test.go:101:			if !strings.Contains(strings.ToLower(confirmation), "abandon") || !strings.Contains(confirmation, "manual") || !strings.Contains(confirmation, oid) {
packages/app/evals/abandon_cli_test.go:164:	if err != nil || owner.ClaimOID == "" {
packages/app/evals/usage_cli_test.go:29:	if !strings.Contains(script, anchor) {
packages/app/evals/usage_cli_test.go:47:	if !strings.Contains(script, anchor) {
packages/app/evals/usage_cli_test.go:82:		if !d.Warning || d.Code == "" || d.Message == "" {
packages/app/evals/usage_cli_test.go:93:		if !strings.Contains(lower, strings.ToLower(marker)) {
packages/app/evals/usage_cli_test.go:105:	if state.Phase != "evals" || state.Run == "" || state.ClaimOID == "" || state.Head != head || len(state.Change.Added) != 1 || state.Change.Added[0].ID != prefix+"-002" {
packages/app/evals/usage_cli_test.go:115:	if strings.Contains(change, "blocked:") || !strings.Contains(change, "phase: evals") || !strings.Contains(change, "run: "+state.Run) {
packages/app/evals/usage_cli_test.go:144:		if err := json.Unmarshal(decodeEnvelope(t, out)["data"], &got); err != nil || !strings.Contains(got.Change, "next") || got.Run == "" || got.Phase != "evals" || got.Empty {
packages/app/evals/usage_cli_test.go:150:	if !strings.Contains(partial, "USAGE_PARTIAL_WORK") {
packages/app/evals/usage_cli_test.go:165:	if err := json.Unmarshal(env["data"], &resumed); err != nil || resumed.Run != held.Run || resumed.Phase != "applied" || !strings.Contains(resumed.Change, "next") || resumed.Empty {
packages/app/evals/usage_cli_test.go:170:		if !strings.Contains(archive, marker) {
packages/app/evals/usage_cli_test.go:186:	} else if body := string(mustReadNextCLI(t, matches[0])); !strings.Contains(body, "ALP-002") || !strings.Contains(body, "fail") {
packages/app/evals/usage_cli_test.go:192:	} else if body := string(mustReadNextCLI(t, matches[0])); !strings.Contains(body, "ALP-002") || !strings.Contains(body, "pass") {
packages/app/evals/usage_cli_test.go:200:			if len(fields) != 3 || fields[0] != phase || fields[2] != bin || fields[1] == "" || (phase != "audit" && fields[1] != root) {
packages/app/evals/usage_cli_test.go:224:	if !strings.Contains(strings.ToLower(out), "next") || !regexp.MustCompile(`(?i)resum|retry|run next`).MatchString(out) {
packages/app/evals/usage_cli_test.go:227:	if !strings.Contains(out, "held-quota-work") {
packages/app/evals/usage_cli_test.go:261:	if !strings.Contains(alphaArchive, "phase: applied") || !strings.Contains(alphaArchive, "### ALP-002") {
packages/app/evals/usage_cli_test.go:272:	if !strings.Contains(partial, "USAGE_PARTIAL_WORK") {
packages/app/evals/usage_cli_test.go:292:	if err := json.Unmarshal(env["data"], &resumed); err != nil || !strings.Contains(resumed.Change, "beta") || resumed.Run != held.Run || resumed.Phase != "applied" || resumed.Empty {
packages/app/evals/usage_cli_test.go:295:	if archive := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/archive/0002-beta.md"))); !strings.Contains(archive, "### BET-002") || !strings.Contains(archive, "run: "+held.Run) {
packages/app/evals/usage_cli_test.go:356:	if !strings.Contains(script, anchor) {
packages/app/evals/usage_cli_test.go:360:	if !strings.Contains(script, selection) {
packages/app/evals/usage_cli_test.go:377:		if !strings.Contains(message, marker) {
packages/app/evals/usage_cli_test.go:382:	if !strings.Contains(alpha, "blocked:") || !strings.Contains(alpha, "phase: evals") || !strings.Contains(alpha, "### ALP-002") {
packages/app/evals/usage_cli_test.go:386:	if saved := cliGit(t, f.root, "for-each-ref", "--format=%(objectname)", ref); saved == "" {
packages/app/evals/usage_cli_test.go:393:	if !strings.Contains(partial, "USAGE_PARTIAL_WORK") {
39:func ref(slug string) string                 { return "refs/fireweed/claims/" + slug }
163:func existing(root, slug string) (claim, string, error) {
181:func state(root string, ch project.Change, c claim, oid, head string) model.State {
184:func storeClaim(root string, c claim, old string) (string, error) {
452:func validate(s model.State, c claim) error {
465:func loadCurrent(root, target string) (model.State, claim, error) {
553:func evalFixZones(s model.State, c claim, diff []gitstore.Change) []project.Diagnostic {
760:func commitPhase(s model.State, c claim, oid string, ch project.Change, paths []string, subject, next string) (model.State, error) {
787:func auditProcess(s model.State, c claim) (project.Change, string, bool, error) {
883:func apply(s model.State, c claim, ch project.Change, oid string) (model.State, error) {

exec
/bin/zsh -lc "cat > packages/app/internal/entry/project.go <<'EOF'
package entry

import (
	\"context\"
	\"fmt\"
	\"os\"
	\"path/filepath\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow\"
)

func array[T any](items []T) []T {
	if items == nil { return []T{} }
	return items
}

type zoneRow struct {
	Protected []string "'`json:"protected"`
	Evals []string `json:"evals"`
	Interface []string `json:"interface"`
	AppendOnly []string `json:"append_only"`
	Preserve []string `json:"preserve"`
	Disposable []string `json:"disposable"`
}

type moduleRow struct {
	Name string `json:"name"`
	IDPrefix string `json:"id_prefix"`
	Dir string `json:"dir"`
	DependsOn []string `json:"depends_on"`
	Zones zoneRow `json:"zones"`
}

func moduleRows(modules []project.Module) ([]moduleRow,string) {
	rows:=make([]moduleRow,0,len(modules))
	var plain strings.Builder
	for _,m:=range modules {
		z:=m.Zones
		rows=append(rows,moduleRow{m.Name,m.IDPrefix,m.Dir,array(m.DependsOn),zoneRow{array(z.Protected),array(z.Evals),array(z.Interface),array(z.AppendOnly),array(z.Preserve),array(z.Disposable)}})
		fmt.Fprintf(&plain,"%s (%s) %s depends_on=%v'"\\n  protected=%v evals=%v interface=%v append_only=%v preserve=%v disposable=%v\\n\",
			m.Name,m.IDPrefix,m.Dir,m.DependsOn,z.Protected,z.Evals,z.Interface,z.AppendOnly,z.Preserve,z.Disposable)
	}
	return rows,plain.String()
}

func findModule(s project.Snapshot, name string) (project.Module,error) {
	for _,m:=range s.Modules {
		if m.Name==name { return m,nil }
	}
	return project.Module{},fmt.Errorf(\"Unknown module %q\",name)
}

func held(root, selector string) (workflow.State,bool) {
	// Worktree claims are Git refs. A project can also be used before Git init.
	for dir:=root;;dir=filepath.Dir(dir) {
		if _,err:=os.Lstat(filepath.Join(dir,\".git\")); err==nil {
			state,err:=workflow.Current(root,selector)
			return state,err==nil
		}
		if filepath.Dir(dir)==dir { return workflow.State{},false }
	}
}

func checkResult(ds []project.Diagnostic, text string) response {
	if hasErrors(ds) { return response{exit:1,diagnostics:reportDiagnostics(ds,true)} }
	r:=confirmation(text)
	r.diagnostics=reportDiagnostics(ds,false)
	return r
}

func projectCommand(ctx context.Context,c command,snap project.Snapshot,bin string) response {
	root:=snap.Root
	arg:=c.arg(0)
	switch c.name {
	case \"modules\":
		if state,ok:=held(root,\"\"); ok { snap=state.Snapshot }
		if err:=project.ValidateGraph(snap.Modules);err"'!=nil { return reject(2,"C02",err) }
		rows,plain:=moduleRows(snap.Modules)
		return response{data:rows,plain:plain}
	case "new-change":
		path,err:=project.NewChange(root,arg,c.arg(1))
		if err!=nil {
			code:="P01"
			switch message:=err.Error(); {
			case strings.HasPrefix(message,"invalid change slug"): code="U01"
			case strings.HasPrefix(message,"unknown module"): code="U02"
			case strings.HasPrefix(message,"invalid change number") || strings.HasPrefix(message,"change slug"): code="C13"
			}
			return reject(2,code,err)
		}
		return response{data:path,plain:path+"'"\\n\"}
	case \"stamp\":
		if err:=project.Stamp(root,arg);err"'!=nil { return reject(1,"P01",err) }
		return confirmation("Stamped "+arg+".")
	case "check":
		if c.has("--fix") {
			if err:=project.StampMissing(root,"");err!=nil { return reject(1,"P01",err) }
		}
		return checkResult(project.Check(root),"Check passed.")
	case "migrate":
		if version:=snap.Config["version"]; version!=1 { return reject(2,"C01",fmt.Errorf("unsupported project format version %v; expected version 1",version)) }
		return checkResult(project.Check(root),"Project format 1 is already current.")
	case "show":
		switch {
		case c.has("--history"): return specHistory(root,arg,c.arg(1))
		case c.has("--stale"): return stale(root,arg)
		default: return show(snap,arg,c.arg(1),c.has("--projected"))
		}
	case "coverage": return coverage(ctx,snap,arg,c.flags["--with-change"],bin)
	case "list": return list(root)
	}
	panic("unhandled project command "+c.name)
}

func specItems(m project.Module) ([]contracts.Item,error) {
	body,_:=contracts.Section(m.Document.Body,"Specs")
	return contracts.ParseItems(body)
}

func show(snap project.Snapshot,module,id string,projected bool) response {
	if err:=project.ValidateGraph(snap.Modules);err!=nil { return reject(2,"C02",err) }
	if projected {
		var err error
		snap,err=project.Project(snap.Root,"")
		if err!=nil { return reject(2,"C02",err) }
	}
	m,err:=findModule(snap,module)
	if err!=nil { return reject(2,"U02",fmt.Errorf("unknown module %q",module)) }
	if id=="" {
		if !projected {
			data,err:=os.ReadFile(filepath.Join(snap.Root,m.Dir,"SPEC.md"))
			if err!=nil { return reject(2,"C03",err) }
			return stringResult(string(data))
		}
		data,err:=contracts.RenderDocument(m.Document)
		if err!=nil { return reject(2,"C03",err) }
		return stringResult(data)
	}
	items,err:=specItems(m)
	if err!=nil && projected {
		// Projection can contain the same local handle from different changes.
		// Validate each matching item on its own to preserve those identities.
		body,_:=contracts.Section(m.Document.Body,"Specs")
		var selected []contracts.Item
		chunks:=strings.Split("'"\\n\"+body,\"\\n### \")
		for _,chunk:=range chunks[1:] {
			part,parseErr:=contracts.ParseItems(\"### \"+chunk)
			if parseErr==nil && len(part)==1 && part[0].ID==id { selected=append(selected,part[0]) }
		}
		if len(selected)>1 && strings.HasPrefix(id,\"new:\") { return reject(2,\"U02\",fmt.Errorf(\"duplicate item ID %q is ambiguous across changes\",id)) }
		code,subject,_:=contracts.ErrorCode(err)
		if code"'!="C05" || !strings.HasPrefix(subject,"new:") { return reject(2,"C03",err) }
		items=selected
		err=nil
	}
	if err!=nil { return reject(2,"C03",err) }
	for _,item:=range items {
		if item.ID==id { return stringResult(contracts.RenderItem(item)) }
	}
	return reject(2,"U02",fmt.Errorf("unknown item %q in module %q",id,module))
}

func evaluate(ctx context.Context, snap project.Snapshot, m project.Module, bin string) (checks.Report,error) {
	items,err:=specItems(m)
	if err!=nil { return checks.Report{},err }
	report,err:=checks.Evaluate(ctx,checks.HookRequest{Snapshot:snap,Module:m.Name,Bin:bin})
	if err!=nil { return report,err }
	if err=checks.RequireCoverage(report,items);err!=nil { return report,err }
	return report,checks.RequirePass(report,items)
}

func coverage(ctx context.Context,snap project.Snapshot,module,through,bin string) response {
	if err:=project.ValidateGraph(snap.Modules);err!=nil { return reject(2,"C02",err) }
	if through!="" {
		if state,ok:=held(snap.Root,through);ok { snap=state.Snapshot } else {
			var err error
			snap,err=project.Project(snap.Root,through)
			if err!=nil { return reject(2,"C02",err) }
		}
	} else if state,ok:=held(snap.Root,"");ok {
		// Preserve current specs while running hooks with approved configuration.
		snap.Config=state.Snapshot.Config
		for i,m:=range snap.Modules {
			if approved,err:=findModule(state.Snapshot,m.Name);err==nil {
				front:=approved.Document.Frontmatter
				snap.Modules[i].Document.Frontmatter=front
				snap.Modules[i].Zones=approved.Zones
				snap.Modules[i].DependsOn=approved.DependsOn
			}
		}
	}
	m,err:=findModule(snap,module)
	if err!=nil { return reject(2,"U02",err) }
	report,err:=evaluate(ctx,snap,m,bin)
	if err!=nil { return reject(1,"C11",err) }
	var plain strings.Builder
	fmt.Fprintf(&plain,"Coverage for %s'"\\n\",module)
	for _,test:=range report.Cases {
		fmt.Fprintf(&plain,\"%s %s %s\\n\",strings.Join(test.IDs,\", \"),test.Result,test.Name)
	}
	return response{data:report,plain:plain.String()}
}

type queueRow struct {
	Slug string "'`json:"slug"`
	Path string `json:"path"`
	Module string `json:"module"`
	Phase string `json:"phase"`
	Blocked string `json:"blocked"`
	AuditRounds int `json:"audit_rounds"`
}

func (r queueRow) line() string {
	return fmt.Sprintf("%s %s module=%s phase=%s blocked=%q audit_rounds=%d",r.Slug,r.Path,r.Module,r.Phase,r.Blocked,r.AuditRounds)
}

func queuePhase(phase string) string {
	if phase=="" { return "queued" }
	return phase
}

func list(root string) response {
	changes,err:=project.Queue(root)
	if err!=nil { return reject(1,"P01",err) }
	rows:=make([]queueRow,0,len(changes))
	var plain strings.Builder
	for _,change:=range changes {
		f:=change.Document.Frontmatter
		phase,_:=f["phase"].(string)
		blocked,_:=f["blocked"].(string)
		rounds,_:=f["audit_rounds"].(int)
		row:=queueRow{change.Slug,change.Path,change.Module,queuePhase(phase),blocked,rounds}
		rows=append(rows,row)
		fmt.Fprintln(&plain,row.line())
	}
	return response{data:rows,plain:plain.String()}
}

type historyRow struct {
	Path string `json:"path"`
	Module string `json:"module"`
	IDs []string `json:"ids"`
	Motivation string `json:"motivation"`
	Decisions string `json:"decisions"`
	RemovedDecisions string `json:"removed_decisions"`
	Provenance string `json:"provenance"`
	Date string `json:"implementation_date"`
}

func specHistory(root,module,id string) response {
	history,err:=project.SpecHistory(root,module,id)
	if err!=nil {
		code,_:=cause(err,"P01")
		exit:=1
		if code=="U02" || code=="C01" { exit=2 }
		return reject(exit,code,err)
	}
	rows:=make([]historyRow,0,len(history))
	var plain strings.Builder
	if len(history)==0 { fmt.Fprintln(&plain,"No history found.") }
	for _,h:=range history {
		date:=""
		if !h.ImplementationDate.IsZero() { date=h.ImplementationDate.Format(time.RFC3339) }
		rows=append(rows,historyRow{h.Path,h.Module,array(h.IDs),h.Motivation,h.Decisions,h.RemovedDecisions,h.Provenance,date})
		if date=="" { date="undated" }
		fmt.Fprintf(&plain,"%s (%s) %s'"\\nImplementation date: %s\\nMotivation\\n%s\\nDecisions\\n%s\\nRemoved Decisions\\n%s\\nProvenance\\n%s\\n\\n\",
			h.Path,h.Module,strings.Join(h.IDs,\", \"),date,h.Motivation,h.Decisions,h.RemovedDecisions,h.Provenance)
	}
	return response{data:rows,plain:plain.String()}
}

func stale(root,selector string) response {
	items,err:=project.StaleSpecs(root,selector)
	if err"'!=nil {
		code,_:=cause(err,"P01")
		exit:=1
		if code=="U02" || code=="C13" || code=="C01" { exit=2 }
		return reject(exit,code,fmt.Errorf("show stale for %q: %w",selector,err))
	}
	var plain strings.Builder
	if len(items)==0 { fmt.Fprintln(&plain,"No stale specs.") }
	for _,item:=range items {
		fmt.Fprintln(&plain,item.ID)
		if item.Recovered {
			writeDiff(&plain,item.Before,item.Current)
		} else {
			fmt.Fprintf(&plain,"Stamped text unavailable (recorded hash %s)'"\\n%s\\n\",item.RecordedHash,item.Current)
		}
	}
	return stringResult(plain.String())
}

func writeDiff(out *strings.Builder,before,current string) {
	a:=strings.Split(strings.TrimSuffix(before,\"\\n\"),\"\\n\")
	b:=strings.Split(strings.TrimSuffix(current,\"\\n\"),\"\\n\")
	// Compute the longest common subsequence, then walk it in line order.
	stride:=len(b)+1
	lengths:=make([]int,(len(a)+1)*stride)
	for i:=len(a)-1;i>=0;i-- {
		for j:=len(b)-1;j>=0;j-- {
			at:=i*stride+j
			if a[i]==b[j] { lengths[at]=1+lengths[at+stride+1] } else { lengths[at]=max(lengths[at+stride],lengths[at+1]) }
		}
	}
	for i,j:=0,0;i<len(a)||j<len(b); {
		switch {
		case i<len(a)&&j<len(b)&&a[i]==b[j]:
			fmt.Fprintln(out,\" \"+a[i]); i++;j++
		case j==len(b)||(i<len(a)&&lengths[(i+1)*stride+j]>=lengths[i*stride+j+1]):
			fmt.Fprintln(out,\"-\"+a[i]);i++
		default:
			fmt.Fprintln(out,\"+\"+b[j]);j++
		}
	}
}
EOF
sed -n '60,90p' packages/app/evals/usage_cli_test.go; sed -n '190,226p' packages/app/evals/gc_cli_test.go; sed -n '1,92p' packages/workflow/internal/lifecycle/gc*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 457ms:
func cliStopEnvelope(t *testing.T, output, wait, budget string) (map[string]json.RawMessage, string) {
	t.Helper()
	env := decodeEnvelope(t, output)
	var data map[string]json.RawMessage
	if err := json.Unmarshal(env["data"], &data); err != nil || data == nil {
		t.Errorf("clean stop data=%s; want object: %v", env["data"], err)
		return nil, ""
	}
	var reason string
	if err := json.Unmarshal(data["reason"], &reason); err != nil {
		t.Errorf("clean stop reason=%s: %v", data["reason"], err)
	}
	cliStopReason(t, reason, wait, budget)
	var diagnostics []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Warning bool   `json:"warning"`
	}
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Errorf("clean stop diagnostics=%s: %v", env["diagnostics"], err)
	}
	for _, d := range diagnostics {
		if !d.Warning || d.Code == "" || d.Message == "" {
			t.Errorf("clean stop has error/invalid diagnostic: %+v", d)
		}
	}
	return data, reason
}

func cliStopReason(t *testing.T, reason, wait, budget string) {
	t.Helper()
	}{
		{"malformed-config", "modules: [broken\n"},
		{"invalid-retention", appFixtureConfig + "gc:\n  drill_max_age: 0s\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root, _, nested := gcCLIProject(t)
			oldDrill := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			gcCLIRef(t, root, oldDrill, cliGit(t, root, "rev-parse", "HEAD"))
			before := gcCLIRefs(t, root)
			appPut(t, root, "fireweed.yaml", tc.config)
			for _, jsonMode := range []bool{false, true} {
				args := []string{"gc"}
				if jsonMode {
					args = append(args, "--json")
				}
				gcCLIFailure(t, bin, nested, args, 2, "C01")
				if got := gcCLIRefs(t, root); !reflect.DeepEqual(got, before) {
					t.Errorf("failed gc changed refs: before=%v after=%v", before, got)
				}
			}
		})
	}
	noGit := appProjectFixture(t)
	for _, jsonMode := range []bool{false, true} {
		args := []string{"gc"}
		if jsonMode {
			args = append(args, "--json")
		}
		gcCLIFailure(t, bin, appNestedDir(t, noGit), args, 1, "P01")
	}
	for _, jsonMode := range []bool{false, true} {
		args := []string{"gc"}
		if jsonMode {
			args = append(args, "--json")
		}
		gcCLIFailure(t, bin, t.TempDir(), args, 2, "C01")
	}
package lifecycle

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

const gcDrillDate = "20060102T150405.000000000"

type gcRef struct{ name, oid string }

func gcGit(ctx context.Context, root string, args ...string) (string, error) {
	r, err := gitstore.Run(ctx, gitstore.Request{
		Args:    append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...),
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0", "GIT_OPTIONAL_LOCKS": "0"},
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, r.Stderr)
	}
	return r.Stdout, nil
}

func gcMaxAge(config map[string]any) (time.Duration, error) {
	raw, exists := config["gc"]
	if !exists {
		return 720 * time.Hour, nil
	}
	settings, ok := raw.(map[string]any)
	if !ok {
		return 0, fmt.Errorf("C01: gc must be a mapping")
	}
	raw, exists = settings["drill_max_age"]
	if !exists {
		return 720 * time.Hour, nil
	}
	value, ok := raw.(string)
	if !ok {
		return 0, fmt.Errorf("C01: gc.drill_max_age must be a positive Go duration")
	}
	age, err := time.ParseDuration(value)
	if err != nil || age <= 0 {
		return 0, fmt.Errorf("C01: gc.drill_max_age must be a positive Go duration: %q", value)
	}
	return age, nil
}

func gcRefs(ctx context.Context, root string) ([]gcRef, error) {
	out, err := gcGit(ctx, root, "for-each-ref", "--format=%(refname) %(objectname)",
		"refs/fireweed/claims", "refs/fireweed/attempts", "refs/fireweed/drills")
	if err != nil {
		return nil, err
	}
	var refs []gcRef
	for _, line := range strings.Split(strings.TrimSpace(out), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			refs = append(refs, gcRef{fields[0], fields[1]})
		}
	}
	return refs, nil
}

func gcRegistered(ctx context.Context, root string) (map[string]bool, error) {
	out, err := gcGit(ctx, root, "worktree", "list", "--porcelain")
	if err != nil {
		return nil, err
	}
	registered := map[string]bool{}
	for _, line := range strings.Split(out, "\n") {
		if strings.HasPrefix(line, "worktree ") {
			registered[strings.TrimPrefix(line, "worktree ")] = true
		}
	}
	return registered, nil
}

func gcArchivedSlugs(root string) (map[string]bool, error) {
	archived := map[string]bool{}
	for _, directory := range []string{"changes/archive", "changes"} {
		entries, err := os.ReadDir(filepath.Join(root, directory))

exec
/bin/zsh -lc "sed -n '610,745p' packages/workflow/internal/lifecycle/stub.go; rg -n 'func (Release|Abandon|Commits)|U02|C01' packages/workflow/internal/lifecycle/{release*,abandon*,commits*}.go; sed -n '255,326p' packages/app/evals/drill_cli_test.go; sed -n '376,430p' packages/app/evals/run_cli_test.go; sed -n '115,220p' packages/app/evals/ci_cli_test.go; sed -n '135,175p' packages/app/evals/gc_cli_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 488ms:
	p := reportPath(s, phase, name)
	_ = os.MkdirAll(filepath.Dir(p), 0755)
	_ = os.WriteFile(p, b, 0644)
}
func Verify(ctx context.Context, req model.Request) (model.Result, error) {
	s, c, e := loadCurrent(req.Root, req.Change)
	r := model.Result{State: s, Reports: map[string]checks.Report{}}
	if e != nil {
		return r, e
	}
	if e = validate(s, c); e != nil {
		return r, e
	}
	if e = hard(project.Check(s.Snapshot.Root)); e != nil {
		return r, e
	}
	if c.Phase != "evals" && c.Phase != "audit" && c.Phase != "impl" {
		return r, fmt.Errorf("P01: unsupported phase %s", c.Phase)
	}
	diff, e := stepChanges(s.Snapshot.Root, c.PhaseBase)
	if e != nil {
		return r, e
	}
	r.Diagnostics = checks.Zones(checks.ZoneRequest{Snapshot: c.Snapshot, Module: s.Change.Module, Phase: c.Phase, ChangePath: s.Change.Path, Base: c.PhaseBase, InterfaceDeclared: strings.TrimSpace(c.Sections["Interface"]) != "", Changes: diff})
	r.Diagnostics = append(r.Diagnostics, evalFixZones(s, c, diff)...)
	if e = hard(r.Diagnostics); e != nil {
		return r, e
	}
	if c.Phase == "audit" {
		_, _, _, e = auditProcess(s, c)
		return r, e
	}
	names := []string{s.Change.Module}
	if c.Phase == "impl" && !req.Quick {
		mode := str(c.Snapshot.Config["verify_dependents"])
		if mode == "" {
			mode = "direct"
		}
		names = append(names, dependents(c.Snapshot, s.Change.Module, mode)...)
	}
	var first error
	for _, name := range names {
		rep, err := checks.Evaluate(ctx, checks.HookRequest{Snapshot: c.Snapshot, Module: name, Bin: req.Bin, Timeout: req.Timeout})
		r.Reports[name] = rep
		saveReport(s, c.Phase, name, rep)
		cost, costErr := measureModuleCost(c.Snapshot, name, rep)
		if costErr != nil {
			return r, costErr
		}
		r.Costs = append(r.Costs, cost)
		if err != nil {
			if first == nil {
				first = err
			}
			continue
		}
		items := specs(module(c.Snapshot, name))
		if name == s.Change.Module {
			if err = checks.RequireCoverage(rep, items); err != nil && first == nil {
				first = err
			}
			if c.Phase == "evals" {
				if ids := evalFixIDs(s.Change); len(ids) > 0 {
					for _, id := range ids {
						if err = checks.RequireCoverage(rep, []contracts.Item{{ID: id}}); err != nil && first == nil {
							first = err
						}
					}
					for _, testcase := range rep.Cases {
						if testcase.Result != "pass" && first == nil {
							first = fmt.Errorf("E01: eval_fix testcase %q did not pass", testcase.Name)
						}
					}
				} else {
					if err = checks.RequireFailFirst(rep, selectedIDs(s.Change, items), list(s.Change.Document.Frontmatter["already_satisfied"])); err != nil {
						r.Diagnostics = append(r.Diagnostics, diag("P2-01", err.Error(), false))
						if first == nil {
							first = err
						}
					}
				}
			}
		}
		if c.Phase == "impl" {
			if err = checks.RequirePass(rep, items); err != nil && first == nil {
				first = err
			}
		}
	}
	if c.Phase == "impl" && !req.Quick {
		ds, err := checks.Boundaries(ctx, checks.BoundaryRequest{Snapshot: c.Snapshot, Bin: req.Bin, Timeout: req.Timeout})
		r.Diagnostics = append(r.Diagnostics, ds...)
		if err != nil && first == nil {
			for _, d := range ds {
				if d.Warning || d.Code == "" {
					continue
				}
				if d.Code == "X02" && strings.HasPrefix(err.Error(), d.Code+": ") {
					first = err
				} else if d.Code == "X02" || strings.Contains(err.Error(), d.Message) {
					first = fmt.Errorf("%s: %w", d.Code, err)
				} else {
					first = fmt.Errorf("%s: %s: %w", d.Code, d.Message, err)
				}
				break
			}
			if first == nil {
				if strings.HasPrefix(err.Error(), "X02: ") {
					first = err
				} else {
					first = fmt.Errorf("X02: %w", err)
				}
			}
		}
		if err = hard(ds); err != nil && first == nil {
			first = err
		}
	}
	warnings := costWarnings(r.Costs, c.Snapshot.Config)
	r.Diagnostics = append(r.Diagnostics, warnings...)
	if err := saveCosts(s, c.Phase, r.Costs, warnings); err != nil {
		return r, err
	}
	return r, first
}
func provenance(ch project.Change, phase, agent string, reports map[string]checks.Report) (project.Change, error) {
	if agent == "" {
		agent = "manual"
	}
	playbook, e := assets.Playbook(phase)
	if e != nil {
		return ch, e
	}
	hash := sha256.Sum256([]byte(playbook))
	total, passing, failing := 0, 0, 0
	for _, r := range reports {
packages/workflow/internal/lifecycle/release_stub.go:9:func Release(root, change string) error {
packages/workflow/internal/lifecycle/commits_stub.go:69:	return "", fmt.Errorf("U02: change %q is ambiguous", selector)
packages/workflow/internal/lifecycle/commits_stub.go:108:func Commits(root, change string) ([]string, error) {
packages/workflow/internal/lifecycle/abandon_stub.go:14:func Abandon(root, change string) (string, error) {
packages/workflow/internal/lifecycle/abandon_stub.go:47:		return "", fmt.Errorf("U02: change %q not found", change)
	drillCLINoInvocation(t, f)
	graph := drillCLIProject(t, bin)
	appPut(t, graph.root, "packages/beta/SPEC.md", strings.Replace(string(mustReadNextCLI(t, filepath.Join(graph.root, "packages/beta/SPEC.md"))), "depends_on: [alpha]", "depends_on: [missing]", 1))
	status, out, stderr = runNextCLI(t, bin, graph.nested, "drill", "alpha", "--json")
	drillCLIFailure(t, status, 2, out, stderr, true, "C02", "missing")
	drillCLINoInvocation(t, graph)
}

func Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, mode, marker string
		jsonMode           bool
	}{
		{"own-plain", "own-fail", "own fixture failure", false},
		{"dependent-json", "dependent-fail", "dependent fixture failure", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := drillCLIProject(t, bin)
			if tc.mode == "dependent-fail" {
				appPut(t, f.observations, "dependent-fail", "yes\n")
			} else {
				appPut(t, f.observations, "mode", tc.mode+"\n")
			}
			args := []string{"drill", "alpha"}
			if tc.jsonMode {
				args = append(args, "--json")
			}
			status, out, stderr := runNextCLI(t, bin, f.nested, args...)
			drillCLIFailure(t, status, 1, out, stderr, tc.jsonMode, "E01", "")
			ref := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha")
			if ref == "" {
				t.Fatal("failed drill did not retain rebuilt ref")
			}
			for _, artifact := range []string{".fireweed/drills/alpha.md", ref} {
				if !strings.Contains(out+stderr, artifact) {
					t.Errorf("failure diagnostic omits artifact %s: %q %q", artifact, out, stderr)
				}
			}
			body := string(mustReadNextCLI(t, filepath.Join(f.root, ".fireweed/drills/alpha.md")))
			if !strings.Contains(body, "result: fail") || !strings.Contains(body, tc.marker) {
				t.Errorf("failed native report=%q", body)
			}
			if got := cliGit(t, f.root, "show", ref+":packages/alpha/internal/value.txt"); got != map[string]string{"own-fail": "wrong", "dependent-fail": "new"}[tc.mode] {
				t.Errorf("retained ref value=%q", got)
			}
			drillCLICleaned(t, f)
		})
	}
}

func Test_APP_141_DrillSetupFailureHasNoArtifacts(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	appPut(t, f.observations, "setup-fail", "yes\n")
	status, out, stderr := runNextCLI(t, bin, f.nested, "drill", "alpha", "--json")
	drillCLIFailure(t, status, 1, out, stderr, true, "P01", "FIXTURE_SETUP_FAILURE")
	for _, artifact := range []string{".fireweed/drills/alpha.md", "refs/fireweed/drills/alpha/"} {
		if strings.Contains(out, artifact) {
			t.Errorf("setup failure invented artifact %s: %q", artifact, out)
		}
	}
	if _, err := os.Stat(filepath.Join(f.root, ".fireweed/drills/alpha.md")); !os.IsNotExist(err) {
		t.Errorf("setup failure wrote report: %v", err)
	}
	if refs := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"); refs != "" {
		t.Errorf("setup failure retained ref: %q", refs)
	}
	drillCLICleaned(t, f)
}

func Test_APP_142_DrillCleanUsageStopPlainAndJSON(t *testing.T) {
func Test_APP_081_APP_084_RunUsageAndConfiguration(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 0, false)
	for _, args := range [][]string{
		{"run", "--max"}, {"run", "--max", "oops"}, {"run", "--max", "0"},
		{"run", "--max", "-1"}, {"run", "extra"}, {"run", "--bogus"},
		{"modules", "--max", "1"}, {"modules", "--allow-unmerged"},
	} {
		status, out, errOut := runNextCLI(t, bin, f.root, args...)
		nextCLIDiagnostic(t, status, 2, out, errOut, false, "U01")
	}
	appPut(t, f.root, "fireweed.yaml", "modules: [broken\n")
	status, out, errOut := runNextCLI(t, bin, f.root, "run")
	nextCLIDiagnostic(t, status, 2, out, errOut, false, "C01")
}

func Test_APP_084_APP_085_APP_086_RunWrappedFailure(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		f := newRunCLISetup(t, bin, 1, false)
		appPut(t, f.root, ".fireweed/omit-added", "yes\n")
		cliCommit(t, f.root, "omit projected coverage")
		args := []string{"run"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, out, errOut := runNextCLI(t, bin, f.root, args...)
		nextCLIDiagnostic(t, status, 1, out, errOut, jsonMode, "C10")
		if !strings.Contains(out+errOut, "ALP-002") {
			t.Errorf("wrapped cause lost ID: stdout=%q stderr=%q", out, errOut)
		}
	}
	f := newRunCLISetup(t, bin, 0, false)
	status, out, errOut := runNextCLI(t, bin, f.root, "run", "--max", "bad", "--json")
	nextCLIDiagnostic(t, status, 2, out, errOut, true, "U01")
}
			found = true
			if !corrective(d.Message) {
				t.Errorf("CI %s diagnostic lacks corrective action: %q", code, d.Message)
			}
		}
	}
	if !found {
		t.Errorf("CI failure lacks %s diagnostic: %s", code, output)
	}
	return output
}

func Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules(t *testing.T) {
	root, nested, base := ciCLIFixture(t)
	status, plain, stderr := runSavedAttemptApp(t, nested, "ci", "--no-doctor")
	if status != 0 || stderr != "" || strings.TrimSpace(plain) == "" {
		t.Errorf("configured-base CI: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	log := doctorCLIFile(t, filepath.Join(root, ".fireweed", "hooks.log"))
	for _, marker := range []string{"eval:alpha:", "eval:beta:", "boundaries:"} {
		if !strings.Contains(log, marker) {
			t.Errorf("CI omitted %s with verify_dependents none: %q", marker, log)
		}
	}
	status, output, stderr := runSavedAttemptApp(t, nested, "ci", base+"..HEAD", "--no-doctor", "--json")
	if status != 0 || stderr != "" || ciCLISuccess(t, output) != strings.TrimSpace(plain) {
		t.Errorf("explicit-range JSON CI: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	if diagnostics := ciCLIDiagnostics(t, output); len(diagnostics) != 0 {
		t.Errorf("clean CI diagnostics=%+v", diagnostics)
	}

	pending := "changes/0001-work.md"
	appPut(t, root, pending, strings.Replace(checkCommitsChange, "%s", "phase: impl\n", 1))
	cliCommit(t, root, "ready for application")
	if err := os.MkdirAll(filepath.Join(root, "changes", "archive"), 0755); err != nil {
		t.Fatal(err)
	}
	cliGit(t, root, "mv", pending, "changes/archive/0001-work.md")
	appPut(t, root, "changes/archive/0001-work.md", strings.Replace(checkCommitsChange, "%s", "phase: applied\n", 1))
	appPut(t, root, "locked/keep.txt", "changed in application\n")
	cliCommit(t, root, "application with protected edit")
	if report, err := checks.History(root, base+"..HEAD"); err != nil || len(report.Diagnostics) == 0 || report.Diagnostics[0].Code != "H01" {
		t.Fatalf("history failure fixture: %+v, %v", report, err)
	}
	ciCLIFailure(t, nested, "H01", 1, "ci", "--no-doctor", "--json")
	status, output, stderr = runSavedAttemptApp(t, nested, "ci", "HEAD..HEAD", "--no-doctor", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("explicit empty range should exclude earlier violation: status=%d stdout=%q stderr=%q", status, output, stderr)
	} else {
		ciCLISuccess(t, output)
	}
}

func Test_APP_110_APP_112_APP_113_CIEvalAndBoundaryFailures(t *testing.T) {
	root, nested, _ := ciCLIFixture(t)
	appPut(t, root, ".fireweed/beta.mode", "missing\n")
	ciCLIFailure(t, nested, "C10", 1, "ci", "--no-doctor", "--json")
	appPut(t, root, ".fireweed/beta.mode", "failing\n")
	ciCLIFailure(t, nested, "E01", 1, "ci", "--no-doctor", "--json")
	status, plain, stderr := runSavedAttemptApp(t, nested, "ci", "--no-doctor")
	if status != 1 || !strings.Contains(plain+stderr, "E01") || !corrective(plain+stderr) {
		t.Errorf("plain failing eval: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	appPut(t, root, ".fireweed/beta.mode", "pass\n")
	appPut(t, root, ".fireweed/boundary-fail", "yes\n")
	ciCLIFailure(t, nested, "X01", 1, "ci", "--no-doctor", "--json")
}

func Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration(t *testing.T) {
	root, nested, _ := ciCLIFixture(t)
	for _, args := range [][]string{{"ci", "one", "two", "--no-doctor", "--json"}, {"ci", "--unexpected", "--json"}} {
		ciCLIFailure(t, nested, "U01", 2, args...)
	}
	appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST pass.\n\n### BET-001\nDuplicate MUST fail.\n")
	ciCLIFailure(t, nested, "C05", 1, "ci", "--no-doctor", "--json")
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	ciCLIFailure(t, nested, "C01", 2, "ci", "--no-doctor", "--json")
}

func Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip(t *testing.T) {
	t.Setenv("FIREWEED_APP_EVAL_CLI", "1")
	bin, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	root, nested, log := doctorCLIFixture(t, bin)
	appPut(t, root, "packages/alpha/evals/alpha_test.go", "package alpha\nimport \"testing\"\nfunc Test_ALP_001_Fixture(t *testing.T) {}\n")
	appPut(t, root, "AGENTS.md", "<!-- fireweed:start -->\nstale\n<!-- fireweed:end -->\n")
	status, output, stderr := runSavedAttemptApp(t, nested, "ci", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("default doctor CI: status=%d stdout=%q stderr=%q", status, output, stderr)
	}
	ciCLISuccess(t, output)
	foundWarning := false
	for _, d := range ciCLIDiagnostics(t, output) {
		if d.Code == "C16" && d.Warning && strings.Contains(d.Message, "AGENTS.md") {
			foundWarning = true
		}
	}
	if !foundWarning {
		t.Errorf("CI lost doctor warning: %s", output)
	}
	if got := doctorCLIFile(t, log); !strings.Contains(got, "setup:") || !strings.Contains(got, bin) || strings.Contains(got, "worker:") || strings.Contains(got, "audit:") {
		t.Errorf("CI doctor did not run canaries with NoAgent: %q", got)
	}
			}
			status, out, stderr = runAppBinary(t, bin, nested, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("second gc: exit=%d stdout=%q stderr=%q", status, out, stderr)
			}
			if jsonMode {
				if got := gcCLIJSONPruned(t, out); len(got) != 0 {
					t.Errorf("second gc pruned=%v; want []", got)
				}
			} else if !strings.Contains(strings.ToLower(out), "no refs") || !strings.Contains(strings.ToLower(out), "pruned") {
				t.Errorf("empty plain gc must explicitly report no refs pruned: %q", out)
			}
		})
	}
}

func gcCLIFailure(t *testing.T, bin, dir string, args []string, wantStatus int, code string) {
	t.Helper()
	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
	status, out, stderr := runAppBinary(t, bin, dir, args...)
	message := commandFailure(t, status, wantStatus, out, stderr, jsonMode)
	if jsonMode {
		message = assertFailedJSON(t, out, code)
	} else if out != "" || !strings.HasPrefix(strings.TrimSpace(stderr), code+":") {
		t.Errorf("plain failure stdout=%q stderr=%q; want %s on stderr", out, stderr, code)
	}
	if !corrective(message) {
		t.Errorf("gc diagnostic lacks correction: %q", message)
	}
}

func Test_APP_150_APP_152_GCRejectsUsageBeforePruning(t *testing.T) {
	bin := buildAppBinary(t)
	root, _, nested := gcCLIProject(t)
	oldDrill := "refs/fireweed/drills/alpha/20200101T000000.000000000"
	gcCLIRef(t, root, oldDrill, cliGit(t, root, "rev-parse", "HEAD"))
	before := gcCLIRefs(t, root)
	for _, args := range [][]string{{"gc", "extra"}, {"gc", "--bogus"}, {"gc", "--max", "1"}} {
		for _, jsonMode := range []bool{false, true} {
			call := append([]string(nil), args...)
			if jsonMode {

exec
/bin/zsh -lc "rg '\"(P01|P02|P06|C10|U01|U02|H0|E01|Z01)' packages/assets/internal/content/*.go; sed -n '91,119p' packages/app/evals/ci_cli_test.go; sed -n '100,138p' packages/app/evals/gc_cli_test.go; sed -n '165,210p' packages/app/evals/drill_cli_test.go; sed -n '326,371p' packages/app/evals/drill_cli_test.go; sed -n '1,70p' packages/workflow/internal/lifecycle/release_stub.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 501ms:
	"C10":   "Required coverage is missing a mapped non-skipped testcase for an enforced spec ID. Add and map a non-skipped testcase to that ID, then rerun the eval; the testcase may pass or fail while coverage is established.",
	"E01":   "An enforced testcase failed or reported an error. Fix the failing test or implementation, then rerun the eval.",
	"H01":   "A worker phase commit changed a protected path or a project playbook override. Restore the edit. Separate human-owned setup and human-owned override work into an author commit outside worker phases.",
	"H02":   "An impl-to-applied commit changed evals or interface files. Make eval edits and interface edits in the evals step before implementation. Rework the affected phase history before rerunning history checks.",
	"H03":   "Modifying an existing append-only path, deleting an existing append-only path, or renaming an existing append-only path violates its append-only rule. Restore the original path and original content. Add a new file for new data.",
	"H04":   "Modifying an archived change file violates archive history. Deleting an archived change file is allowed for a revert. Restore the archive and author follow-up work as a pending change.",
	"U01":   "An invalid argument or option, or an unsupported command, was requested. Check fireweed help and use the supported command syntax and options.",
	"U02":   "An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.",
	"P01":   "A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.",
	"P02":   "A change claim is missing, held by another owner, or unreadable or undecodable because its claim data is invalid. If no claim exists, start an eligible change. If claim data is unreadable, undecodable, or invalid, seek approved human review and recovery before retrying. If another owner holds the claim, wait for the holder to release it or coordinate with that holder to resolve ownership before retrying.",
	"P06":   "The change is blocked by unresolved Questions. Review the Questions and resolve them before continuing the change.",
	"Z01":   "A forbidden phase-zone edit changed a file outside the authorized zone. Restore the forbidden edit or move the work to the appropriate authorized phase and zone.",
	Warning       bool
} {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []struct {
		Code, Message string
		Warning       bool
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("CI diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	return diagnostics
}

func ciCLIFailure(t *testing.T, nested, code string, status int, args ...string) string {
	t.Helper()
	got, output, stderr := runSavedAttemptApp(t, nested, args...)
	if got != status || stderr != "" {
		t.Errorf("ci %v: status=%d stdout=%q stderr=%q; want %d and empty stderr", args, got, output, stderr, status)
	}
	assertFailedJSON(t, output)
	found := false
	for _, d := range ciCLIDiagnostics(t, output) {
		if d.Code == code {
			found = true
			if !corrective(d.Message) {
				t.Errorf("CI %s diagnostic lacks corrective action: %q", code, d.Message)
			}
		}
			beforeHead := cliGit(t, root, "rev-parse", "HEAD")
			beforeIndex := cliGit(t, root, "write-tree")
			beforeStatus := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all")
			want := []string{attempt, staleClaim, oldDrill}
			sort.Strings(want)
			args := []string{"gc"}
			if jsonMode {
				args = append(args, "--json")
			}
			status, out, stderr := runAppBinary(t, bin, nested, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("gc: exit=%d stdout=%q stderr=%q", status, out, stderr)
			}
			if jsonMode {
				if got := gcCLIJSONPruned(t, out); !reflect.DeepEqual(got, want) {
					t.Errorf("JSON pruned=%v; want sorted full names %v", got, want)
				}
			} else {
				position := -1
				for _, ref := range want {
					at := strings.Index(out, ref)
					if at <= position || strings.Count(out, ref) != 1 {
						t.Errorf("plain output %q does not list %q once in order", out, ref)
					}
					position = at
				}
			}
			if got := gcCLIRefs(t, root); !reflect.DeepEqual(got, []string{"refs/fireweed/claims/linked-live", "refs/fireweed/claims/main-live", futureDrill, "refs/fireweed/other/keep"}) {
				t.Errorf("gc retained refs=%v", got)
			}
			if got := gcCLIRefs(t, outer); !reflect.DeepEqual(got, []string{oldDrill}) {
				t.Errorf("gc changed enclosing project refs=%v", got)
			}
			if cliGit(t, root, "rev-parse", "HEAD") != beforeHead || cliGit(t, root, "write-tree") != beforeIndex || cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all") != beforeStatus {
				t.Error("gc changed HEAD, index or working files")
			}
			status, out, stderr = runAppBinary(t, bin, nested, args...)
			if status != 0 || stderr != "" {
				t.Fatalf("second gc: exit=%d stdout=%q stderr=%q", status, out, stderr)
			t.Errorf("plain drill output %q lacks %q", out, marker)
		}
	}
	if got := drillCLIObserve(t, f, "setup"); !strings.Contains(got, bin) {
		t.Errorf("setup FIREWEED_BIN=%q; want %s", got, bin)
	}
	if got := drillCLIObserve(t, f, "agent"); !strings.Contains(got, bin) {
		t.Errorf("agent FIREWEED_BIN=%q; want %s", got, bin)
	}
	if got := drillCLIObserve(t, f, "evals"); !strings.Contains(got, "alpha|"+bin) || !strings.Contains(got, "beta|"+bin) {
		t.Errorf("own/dependent evals and Bin=%q", got)
	}
	if body := string(mustReadNextCLI(t, filepath.Join(f.root, report))); !strings.Contains(body, "base_commit: "+f.base) || !strings.Contains(body, "result: pass") {
		t.Errorf("native report=%q", body)
	}
	if got := cliGit(t, f.root, "show", ref+":packages/alpha/internal/value.txt"); got != "new" {
		t.Errorf("retained rebuilt code=%q", got)
	}
	if got := cliGit(t, f.root, "rev-parse", "HEAD"); got != f.base {
		t.Errorf("drill moved invoking HEAD: %s -> %s", f.base, got)
	}
	drillCLICleaned(t, f)
}

func Test_APP_138_DrillJSONSuccess(t *testing.T) {
	bin := buildAppBinary(t)
	f := drillCLIProject(t, bin)
	status, out, stderr := runNextCLI(t, bin, f.nested, "drill", "alpha", "--json")
	if status != 0 || stderr != "" {
		t.Fatalf("JSON drill: status=%d stdout=%q stderr=%q", status, out, stderr)
	}
	data := drillCLIJSONData(t, out, false)
	if data["module"] != "alpha" || data["result"] != "pass" || data["report"] != ".fireweed/drills/alpha.md" || data["base_commit"] != f.base || !strings.HasPrefix(data["ref"], "refs/fireweed/drills/alpha/") {
		t.Errorf("JSON drill result=%v", data)
	}
	if body := string(mustReadNextCLI(t, filepath.Join(f.root, data["report"]))); !strings.Contains(body, "result: pass") {
		t.Errorf("JSON report path does not identify native report: %q", body)
	}
	if got := cliGit(t, f.root, "show", data["ref"]+":packages/alpha/internal/value.txt"); got != "new" {
		t.Errorf("JSON ref does not identify rebuilt code: %q", got)
	}
	drillCLICleaned(t, f)
}

func Test_APP_139_DrillArgumentAndModuleRejectionsDoNotInvokeHooks(t *testing.T) {
	bin := buildAppBinary(t)
func Test_APP_142_DrillCleanUsageStopPlainAndJSON(t *testing.T) {
	bin := buildAppBinary(t)
	for _, jsonMode := range []bool{false, true} {
		f := drillCLIProject(t, bin)
		if jsonMode {
			config := string(mustReadNextCLI(t, filepath.Join(f.root, "fireweed.yaml")))
			appPut(t, f.root, "fireweed.yaml", strings.Replace(config, "attempts: 1", "attempts: 2", 1))
			cliCommit(t, f.root, "allow verification then usage stop")
			f.base = cliGit(t, f.root, "rev-parse", "HEAD")
			appPut(t, f.observations, "mode", "usage-after-failure\n")
		} else {
			appPut(t, f.observations, "mode", "usage\n")
		}
		args := []string{"drill", "alpha"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, out, stderr := runNextCLI(t, bin, f.nested, args...)
		if status != 3 || stderr != "" {
			t.Errorf("clean stop exit=%d stdout=%q stderr=%q", status, out, stderr)
		}
		if jsonMode {
			data := drillCLIJSONData(t, out, true)
			if data["module"] != "alpha" || data["result"] != "stopped" || data["report"] != ".fireweed/drills/alpha.md" || !strings.HasPrefix(data["ref"], "refs/fireweed/drills/alpha/") || data["base_commit"] != f.base || !strings.Contains(data["reason"], "budget") {
				t.Errorf("clean stop data=%v", data)
			}
			if ref := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"); ref != data["ref"] {
				t.Errorf("stop ref=%q; want JSON ref %q", ref, data["ref"])
			}
		} else if !strings.Contains(out, "alpha") || !strings.Contains(strings.ToLower(out), "stopped") || !strings.Contains(out, "budget") || !strings.Contains(out, ".fireweed/drills/alpha.md") {
			t.Errorf("plain clean stop omits module, reason, or report: %q", out)
		}
		if body := string(mustReadNextCLI(t, filepath.Join(f.root, ".fireweed/drills/alpha.md"))); !strings.Contains(body, "result: stopped") {
			t.Errorf("stopped report=%q", body)
		}
		if !jsonMode {
			if ref := cliGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills/alpha"); ref != "" {
				t.Errorf("usage stop retained ref: %q", ref)
			}
		}
		drillCLICleaned(t, f)
	}
}

func Test_APP_136_DrillUsesConfiguredTimeout(t *testing.T) {
	bin := buildAppBinary(t)
package lifecycle

import (
	"fmt"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

func Release(root, change string) error {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return err
	}
	current, err := Current(repo.Root, change)
	if err != nil {
		return err
	}
	if err := gitstore.CASRef(repo.Root, ref(current.Change.Slug), current.ClaimOID, ""); err != nil {
		return fmt.Errorf("P02: release claim: %w", err)
	}
	return nil
}

exec
/bin/zsh -lc "sed -i '' 's/return message+\"\\. \"+help/return message+\". To correct this, follow the guidance for \"+code+\". \"+help/' packages/app/internal/entry/entry.go
cat > packages/app/internal/entry/lifecycle.go <<'EOF'
package entry

import (
	\"context\"
	\"fmt\"
	\"path/filepath\"
	\"slices\"
	\"strconv\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/runner\"
	\"github.com/lurkshark/fireweed/packages/workflow\"
)

func lines(items []string) string {
	if len(items)==0 { return \"\" }
	return strings.Join(items,\"\\n\")+\"\\n\"
}

func operationFailure(err error,lookupIsUsage bool) response {
	code,_:=cause(err,\"P01\")
	exit:=1
	if code==\"C01\" || (lookupIsUsage && code==\"U02\") { exit=2 }
	return reject(exit,code,err)
}

func workflowCommand(ctx context.Context,c command,root,bin string) response {
	arg:=c.arg(0)
	switch c.name {
	case \"start\":
		var state workflow.State
		var err error
		if c.has(\"--from-attempt\") {
			n,_:=strconv.Atoi(c.flags[\"--from-attempt\"])
			state,err=workflow.StartFromAttempt(root,arg,n)
		} else { state,err=workflow.Start(root,arg) }
		var r response
		if err"'!=nil { r=operationFailure(err,false) } else {
			plain,err:=workflow.Context(root,state.Change.Slug,false)
			if err!=nil { r=operationFailure(err,false) } else { r=stringResult(plain) }
		}
		r.diagnostics=append(r.diagnostics,reportDiagnostics(state.Diagnostics,false)...)
		return r
	case "context":
		plain,err:=workflow.Context(root,arg,c.has("--full"))
		if err!=nil { return operationFailure(err,false) }
		return stringResult(plain)
	case "verify","done":
		run:=workflow.Verify
		if c.name=="done" { run=workflow.Done }
		result,err:=run(ctx,workflow.Request{Root:root,Change:arg,Bin:bin,Quick:c.has("--quick")})
		ds:=reportDiagnostics(result.Diagnostics,err!=nil)
		if err!=nil {
			r:=operationFailure(err,false)
			code,message:=cause(err,"P01")
			if slices.ContainsFunc(ds,func(d diagnostic)bool{return !d.Warning && d.Code==code && strings.Contains(d.Message,message)}) {
				r.diagnostics=ds
			} else { r.diagnostics=append(r.diagnostics,ds...) }
			return r
		}
		return response{data:result,plain:fmt.Sprintf("%s is in phase %s.'"\\n\",result.State.Change.Slug,result.State.Phase),diagnostics:ds}
	case \"release\":
		if err:=workflow.Release(root,arg);err"'!=nil { return operationFailure(err,false) }
		return confirmation("Released "+arg+".")
	case "commits":
		commits,err:=workflow.Commits(root,arg)
		if err!=nil { return operationFailure(err,true) }
		return response{data:array(commits),plain:lines(commits)}
	case "abandon":
		oid,err:=workflow.Abandon(root,arg)
		if err!=nil { return operationFailure(err,true) }
		return confirmation(fmt.Sprintf("Abandoned %s in commit %s.",arg,oid))
	case "gc":
		result,err:=workflow.GC(ctx,workflow.GCRequest{Root:root})
		if err!=nil { return operationFailure(err,false) }
		pruned:=array(result.Pruned)
		slices.Sort(pruned)
		plain:=lines(pruned)
		if len(pruned)==0 { plain="No refs pruned.'"\\n\" }
		return response{data:map[string][]string{\"pruned\":pruned},plain:plain}
	}
	panic(\"unhandled lifecycle command \"+c.name)
}

type nextRow struct {
	Change string "'`json:"change"`
	Run string `json:"run"`
	Phase string `json:"phase"`
}

func summary(result runner.NextResult) nextRow {
	return nextRow{result.State.Change.Slug,result.State.Run,result.State.Phase}
}

func runnerCommand(ctx context.Context,c command,snap project.Snapshot,bin string) response {
	root:=snap.Root
	allow:=c.has("--allow-unmerged")
	switch c.name {
	case "drill":
		return drill(ctx,snap,c.arg(0),bin)
	case "next":
		result,err:=runner.Next(ctx,runner.NextRequest{Root:root,Bin:bin,AllowUnmerged:allow})
		data:=struct {
			nextRow
			Empty bool `json:"empty"`
			Reason string `json:"reason,omitempty"`
		}{nextRow:summary(result),Empty:result.Empty}
		ds:=reportDiagnostics(result.State.Diagnostics,false)
		if err!=nil {
			if runner.IsUsageStop(err) {
				data.Reason=err.Error()
				return response{exit:3,data:data,plain:fmt.Sprintf("Stopped with %s held: %s. Resume with fireweed next.'"\\n\",data.Change,data.Reason),diagnostics:ds}
			}
			r:=operationFailure(err,false)
			r.diagnostics=append(r.diagnostics,ds...)
			return r
		}
		plain:=\"Queue is empty.\\n\"
		if "'!result.Empty { plain="Applied "+data.Change+".'"\\n\" }
		return response{data:data,plain:plain,diagnostics:ds}
	case \"run\":
		max,_:=strconv.Atoi(c.flags[\"--max\"])
		duration,_:=time.ParseDuration(c.flags[\"--for\"])
		result,err:=runner.Run(ctx,runner.RunRequest{Root:root,Bin:bin,AllowUnmerged:allow,Max:max,For:duration})
		data:=struct {
			Applied []nextRow "'`json:"applied"`
			Empty bool `json:"empty"`
			Limited bool `json:"limited"`
			Reason string `json:"reason,omitempty"`
		}{Applied:make([]nextRow,0,len(result.Applied)),Empty:result.Empty,Limited:result.Limited}
		var names []string
		var ds []diagnostic
		for _,applied:=range result.Applied {
			row:=summary(applied)
			data.Applied=append(data.Applied,row)
			names=append(names,row.Change)
			ds=append(ds,reportDiagnostics(applied.State.Diagnostics,false)...)
		}
		plain:=fmt.Sprintf("Applied %d changes",len(names))
		if len(names)>0 { plain+=": "+strings.Join(names,", ") }
		plain+=".'"\\n\"
		if err"'!=nil {
			if runner.IsUsageStop(err) || runner.IsTimeStop(err) {
				data.Reason=err.Error()
				return response{exit:3,data:data,plain:plain+"Stopped: "+data.Reason+". Resume with fireweed run.'"\\n\",diagnostics:ds}
			}
			r:=operationFailure(err,false)
			r.diagnostics=append(r.diagnostics,ds...)
			return r
		}
		if result.Empty { plain+=\"Queue is empty.\\n\" }
		if result.Limited { plain+=\"Change limit reached with queued work remaining.\\n\" }
		return response{data:data,plain:plain,diagnostics:ds}
	}
	panic(\"unhandled runner command \"+c.name)
}

func drill(ctx context.Context,snap project.Snapshot,module,bin string) response {
	if err:=project.ValidateGraph(snap.Modules);err"'!=nil { return reject(2,"C02",err) }
	if _,err:=findModule(snap,module);err!=nil { return reject(2,"U02",err) }
	result,err:=runner.Drill(ctx,runner.DrillRequest{Root:snap.Root,Module:module,Bin:bin})
	report:=result.ReportPath
	if filepath.IsAbs(report) {
		if rel,e:=filepath.Rel(snap.Root,report);e==nil { report=filepath.ToSlash(rel) }
	}
	data:=struct {
		Module string `json:"module"`
		Result string `json:"result"`
		Report string `json:"report"`
		Ref string `json:"ref"`
		Base string `json:"base_commit"`
		Reason string `json:"reason,omitempty"`
	}{Module:result.Module,Result:result.Result,Report:report,Ref:result.Ref,Base:result.Base}
	var artifacts string
	if report!="" { artifacts+=" Report: "+report+"." }
	if result.Ref!="" { artifacts+=" Rebuilt ref: "+result.Ref+"." }
	if err!=nil {
		if result.Result=="stopped" && runner.IsUsageStop(err) {
			data.Reason=err.Error()
			return response{exit:3,data:data,plain:fmt.Sprintf("Drill for %s stopped: %s.%s'"\\n\",module,data.Reason,artifacts)}
		}
		if artifacts"'!="" { err=fmt.Errorf("%w.%s",err,artifacts) }
		return operationFailure(err,false)
	}
	return response{data:data,plain:fmt.Sprintf("Drill for %s: %s.%s'"\\n\",module,result.Result,artifacts)}
}
EOF
sed -n '240,289p' packages/app/evals/check_commits_cli_test.go; sed -n '26,61p' packages/runner/internal/drill/drill.go; rg -n 'ReportPath' packages/runner/internal/drill/drill.go; rg -n 'func hard' packages/workflow/internal/lifecycle/stub.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 496ms:
		{"option", []string{"check-commits", "--bogus"}, 2, "U01"},
		{"bad-range", []string{"check-commits", "missing-ref..HEAD", "--json"}, 1, "P01"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status, output, stderr := runSavedAttemptApp(t, nested, tc.args...)
			message := commandFailure(t, status, tc.status, output, stderr, strings.Contains(strings.Join(tc.args, " "), "--json"))
			if strings.Contains(strings.Join(tc.args, " "), "--json") {
				message = assertFailedJSON(t, output, tc.code)
			} else if !strings.Contains(output+stderr, tc.code+":") {
				t.Errorf("diagnostic=%q; want %s", output+stderr, tc.code)
			}
			if !corrective(message) {
				t.Errorf("diagnostic %q needs corrective action", message)
			}
		})
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	status, output, stderr := runSavedAttemptApp(t, nested, "check-commits", "--json")
	message := commandFailure(t, status, 2, output, stderr, true)
	message = assertFailedJSON(t, output, "C01")
	if !corrective(message) {
		t.Errorf("configuration diagnostic %q needs corrective action", message)
	}

	codedRoot, _, codedBase := checkCommitsFixture(t)
	appPut(t, codedRoot, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n")
	cliCommit(t, codedRoot, "record malformed SPEC")
	cliGit(t, codedRoot, "checkout", "HEAD^", "--", "packages/alpha/SPEC.md")
	cliCommit(t, codedRoot, "restore current SPEC")
	if _, err := checks.History(codedRoot, codedBase+"..HEAD"); err == nil {
		t.Fatal("coded operational fixture: want malformed recorded SPEC")
	} else {
		code, _, coded := contracts.ErrorCode(err)
		if !(coded && code == "C03" || strings.HasPrefix(err.Error(), "C03:")) || !strings.Contains(err.Error(), "packages/alpha/SPEC.md") {
			t.Fatalf("coded operational fixture: want C03 for malformed recorded SPEC path, got %v", err)
		}
	}
	status, output, stderr = runSavedAttemptApp(t, appNestedDir(t, codedRoot), "check-commits", codedBase+"..HEAD", "--json")
	message = commandFailure(t, status, 1, output, stderr, true)
	message = assertFailedJSON(t, output, "C03")
	if !corrective(message) {
		t.Errorf("coded operational diagnostic %q needs corrective action", message)
	}
}
func Run(ctx context.Context, request model.DrillRequest) (result model.DrillResult, err error) {
	result.Module = request.Module
	result.Dependents = make(map[string]checks.Report)
	snapshot, err := project.Load(request.Root)
	if err != nil {
		return result, err
	}
	target := findModule(snapshot, request.Module)
	if target == nil {
		return result, fmt.Errorf("U02: unknown module %q", request.Module)
	}
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return result, err
	}
	result.Base = repo.Head
	status, err := git(repo.Root, "--no-optional-locks", "status", "--porcelain", "-uall")
	if err != nil {
		return result, err
	}
	if status != "" {
		return result, fmt.Errorf("drill requires a clean invoking checkout: dirty worktree")
	}
	claims, err := git(repo.Root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims")
	if err != nil {
		return result, err
	}
	if claims != "" {
		return result, fmt.Errorf("drill requires no live shared change claim: %s", claims)
	}

	result.Date = time.Now().UTC().Format("20060102T150405.000000000")
	result.Result = "fail"
	result.ReportPath = filepath.ToSlash(filepath.Join(".fireweed", "drills", request.Module+".md"))
	result.Agent = agentName(snapshot.Config)
	result.Model = modelName(snapshot.Config)
59:	result.ReportPath = filepath.ToSlash(filepath.Join(".fireweed", "drills", request.Module+".md"))
97:				err = errors.Join(err, markReportFailure(repo.Root, result.ReportPath, cleanup))
813:	path := filepath.Join(root, filepath.FromSlash(result.ReportPath))
49:func hard(ds []project.Diagnostic) error {

codex
The shared CLI dispatch, diagnostics and project commands are in place. I’ll finish wiring the workflow and runner results, then check the setup commands against the retained evals.
exec
/bin/zsh -lc "cat > packages/app/internal/entry/reports.go <<'EOF'
package entry

import (
	\"context\"
	\"fmt\"
	\"maps\"
	\"slices\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow\"
)

type doctorRow struct {
	Name string "'`json:"name"`
	Module string `json:"module"`
	Passed bool `json:"passed"`
	Detail string `json:"detail"`
}

func doctorResult(report checks.DoctorReport) response {
	rows:=make([]doctorRow,0,len(report.Checks))
	var plain strings.Builder
	for _,check:=range report.Checks {
		rows=append(rows,doctorRow{check.Name,check.Module,check.Passed,check.Detail})
		outcome:="passed"
		if !check.Passed { outcome="failed" }
		fmt.Fprintf(&plain,"%s %s: %s %s'"\\n\",check.Name,check.Module,outcome,check.Detail)
	}
	r:=response{data:map[string][]doctorRow{\"checks\":rows},plain:plain.String(),diagnostics:reportDiagnostics(report.Diagnostics,true)}
	if hasErrors(report.Diagnostics) { r.exit=1;r.data=nil }
	return r
}

type squashRow struct {
	Commit string "'`json:"commit"`
	Change string `json:"change"`
}

func checksCommand(ctx context.Context,c command,snap project.Snapshot,bin string) response {
	root:=snap.Root
	if c.name=="doctor" {
		return doctorResult(checks.Doctor(ctx,checks.DoctorRequest{Root:root,Bin:bin,NoAgent:c.has("--no-agent")}))
	}
	if c.name=="ci" { return ci(ctx,snap,bin,c.arg(0),c.has("--no-doctor")) }
	report,err:=checks.History(root,c.arg(0))
	if err!=nil { return reject(1,"P01",err) }
	if hasErrors(report.Diagnostics) { return response{exit:1,diagnostics:reportDiagnostics(report.Diagnostics,true)} }
	var plain strings.Builder
	fmt.Fprintf(&plain,"Checked %d commits.'"\\n\",len(report.Commits))
	rows:=make([]squashRow,0,len(report.Squashed))
	for _,item:=range report.Squashed {
		rows=append(rows,squashRow{item.Commit,item.Change})
		fmt.Fprintf(&plain,\"Squashed %s %s\\n\",item.Commit,item.Change)
	}
	return response{plain:plain.String(),data:struct {
		Commits []string "'`json:"commits"`
		Squashed []squashRow `json:"squashed"`
	}{array(report.Commits),rows},diagnostics:reportDiagnostics(report.Diagnostics,false)}
}

func ci(ctx context.Context,snap project.Snapshot,bin,revisionRange string,noDoctor bool) response {
	ds:=project.Check(snap.Root)
	if !hasErrors(ds) {
		history,err:=checks.History(snap.Root,revisionRange)
		if err!=nil {
			r:=reject(1,"P01",err)
			r.diagnostics=append(r.diagnostics,reportDiagnostics(ds,false)...)
			return r
		}
		ds=append(ds,history.Diagnostics...)
	}
	if hasErrors(ds) { return response{exit:1,diagnostics:reportDiagnostics(ds,true)} }
	for _,m:=range snap.Modules {
		if _,err:=evaluate(ctx,snap,m,bin);err!=nil {
			code,message:=cause(err,"C11")
			ds=append(ds,project.Diagnostic{Code:code,Message:m.Name+": "+message})
		}
	}
	boundaries,err:=checks.Boundaries(ctx,checks.BoundaryRequest{Snapshot:snap,Bin:bin})
	ds=append(ds,boundaries...)
	if err!=nil && !hasErrors(boundaries) {
		code,message:=cause(err,"X02")
		ds=append(ds,project.Diagnostic{Code:code,Message:message})
	}
	if !noDoctor {
		report:=checks.Doctor(ctx,checks.DoctorRequest{Root:snap.Root,Bin:bin,NoAgent:true})
		ds=append(ds,report.Diagnostics...)
	}
	r:=checkResult(ds,"CI passed.")
	if r.exit==0 { r.data=strings.TrimSuffix(r.plain,"'"\\n\") }
	return r
}

type statusQueueRow struct {
	queueRow
	WaitingOn string "'`json:"waiting_on"`
	WaitingRule string `json:"waiting_rule"`
}

type quarantineRow struct {
	Module string `json:"module"`
	ID string `json:"id"`
	Reason string `json:"reason"`
}

type optOutRow struct {
	Module string `json:"module"`
	Reason string `json:"reason"`
}

type costRow struct {
	Module string `json:"module"`
	Seconds float64 `json:"seconds"`
	CasesPerSpec map[string]int `json:"cases_per_spec"`
	EvalLines int `json:"eval_lines"`
	DisposableLines int `json:"disposable_lines"`
}

func countMap(m map[string]int) map[string]int {
	if m==nil { return map[string]int{} }
	return m
}

func counts(out *strings.Builder,label string,items map[string]int) {
	for _,name:=range slices.Sorted(maps.Keys(items)) {
		fmt.Fprintf(out,"%s %s=%d'"\\n\",label,name,items[name])
	}
}

func status(root string) response {
	report,err:=workflow.Status(root)
	if err"'!=nil { return operationFailure(err,false) }
	modules,moduleText:=moduleRows(report.Modules)
	data:=struct {
		Modules []moduleRow `json:"modules"`
		Queue []statusQueueRow `json:"queue"`
		Quarantines []quarantineRow `json:"quarantines"`
		BoundaryOptOuts []optOutRow `json:"boundary_opt_outs"`
		Drift map[string]int `json:"drift_since_drill"`
		Unapproved int `json:"unapproved"`
		PhaseCommits map[string]int `json:"phase_commits"`
		Metrics []costRow `json:"metrics"`
	}{
		Modules:modules,Queue:make([]statusQueueRow,0,len(report.Queue)),
		Quarantines:make([]quarantineRow,0,len(report.Quarantines)),
		BoundaryOptOuts:make([]optOutRow,0,len(report.BoundaryOptOuts)),
		Drift:countMap(report.DriftSinceDrill),Unapproved:report.Unapproved,
		PhaseCommits:countMap(report.PhaseCommits),Metrics:make([]costRow,0,len(report.Costs)),
	}
	var plain strings.Builder
	if len(report.Queue)==0 { fmt.Fprintln(&plain,"Queue is empty.") }
	for _,item:=range report.Queue {
		row:=statusQueueRow{queueRow{item.Slug,item.Path,item.Module,queuePhase(item.Phase),item.Blocked,item.AuditRounds},item.WaitingOn,item.WaitingRule}
		data.Queue=append(data.Queue,row)
		fmt.Fprintf(&plain,"%s waiting_on=%q waiting_rule=%q'"\\n\",row.line(),row.WaitingOn,row.WaitingRule)
	}
	plain.WriteString(moduleText)
	for _,q:=range report.Quarantines {
		data.Quarantines=append(data.Quarantines,quarantineRow(q))
		fmt.Fprintf(&plain,\"Quarantine %s %s: %s\\n\",q.Module,q.ID,q.Reason)
	}
	for _,o:=range report.BoundaryOptOuts {
		data.BoundaryOptOuts=append(data.BoundaryOptOuts,optOutRow(o))
		fmt.Fprintf(&plain,\"Boundary opt-out %s: %s\\n\",o.Module,o.Reason)
	}
	counts(&plain,\"Drift since drill\",data.Drift)
	fmt.Fprintf(&plain,\"Unapproved changes: %d\\n\",report.Unapproved)
	counts(&plain,\"Phase commits\",data.PhaseCommits)
	for _,cost:=range report.Costs {
		data.Metrics=append(data.Metrics,costRow{cost.Module,cost.Seconds,countMap(cost.CasesPerSpec),cost.EvalLines,cost.DisposableLines})
		fmt.Fprintf(&plain,\"Cost %s: JUnit seconds=%g eval lines=%d disposable lines=%d\\n\",cost.Module,cost.Seconds,cost.EvalLines,cost.DisposableLines)
		counts(&plain,\"Testcases\",cost.CasesPerSpec)
	}
	return response{data:data,plain:plain.String(),diagnostics:reportDiagnostics(report.Diagnostics,false)}
}

func prBody(root,revisionRange string) response {
	report,err:=workflow.PRBody(root,revisionRange)
	if err"'!=nil {
		if workflow.IsEndpointConfigError(err) { return reject(2,"C01",err) }
		return reject(1,"P01",err)
	}
	r:=stringResult(report.Body)
	r.diagnostics=reportDiagnostics(report.Diagnostics,false)
	return r
}
EOF
sed -n '"'1,264p' packages/app/evals/init_test.go; sed -n '264,545p' packages/app/evals/init_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 468ms:
package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"gopkg.in/yaml.v3"
)

type initConfig struct {
	Fireweed       string   `yaml:"fireweed"`
	Version        int      `yaml:"version"`
	Base           string   `yaml:"base"`
	Remote         string   `yaml:"remote"`
	Modules        []string `yaml:"modules"`
	ModuleDefaults struct {
		Protected  []string `yaml:"protected"`
		Interface  []string `yaml:"interface"`
		Evals      []string `yaml:"evals"`
		AppendOnly []string `yaml:"append_only"`
		Preserve   []string `yaml:"preserve"`
		Disposable []string `yaml:"disposable"`
	} `yaml:"module_defaults"`
	Hooks          map[string]any `yaml:"hooks"`
	Protected      []string       `yaml:"protected"`
	WorkerWritable []string       `yaml:"worker_writable"`
	Agent          struct {
		Command any `yaml:"command"`
	} `yaml:"agent"`
	Audit struct {
		Enabled bool `yaml:"enabled"`
		Command any  `yaml:"command"`
	} `yaml:"audit"`
}

func readInitFile(t *testing.T, root, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
	if err != nil {
		t.Errorf("read %s: %v", name, err)
		return ""
	}
	return string(data)
}

func readInitConfig(t *testing.T, root string) initConfig {
	t.Helper()
	var config initConfig
	if err := yaml.Unmarshal([]byte(readInitFile(t, root, "fireweed.yaml")), &config); err != nil {
		t.Errorf("parse fireweed.yaml: %v", err)
	}
	return config
}

func initCommand(value any) []string {
	items, ok := value.([]any)
	if !ok {
		return nil
	}
	result := make([]string, 0, len(items))
	for _, item := range items {
		value, ok := item.(string)
		if !ok {
			return nil
		}
		result = append(result, value)
	}
	return result
}

func assertInitOutput(t *testing.T, status int, stdout, stderr, instructions string) {
	t.Helper()
	if status != 0 || stderr != "" {
		t.Errorf("init status=%d stderr=%q; want 0 and empty", status, stderr)
	}
	if strings.TrimSuffix(stdout, "\n") != strings.TrimSuffix(instructions, "\n") {
		t.Errorf("init stdout=%q; want instructions %q", stdout, instructions)
	}
}

func assertNoInitConfig(t *testing.T, root string) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, "fireweed.yaml")); !os.IsNotExist(err) {
		t.Errorf("rejected init created fireweed.yaml or returned unexpected stat error: %v", err)
	}
}

func initVersionParts(value string) ([3]int, bool) {
	var parts [3]int
	fields := strings.Split(value, ".")
	if len(fields) == 0 || len(fields) > len(parts) {
		return parts, false
	}
	for i, field := range fields {
		part, err := strconv.Atoi(field)
		if err != nil || part < 0 {
			return parts, false
		}
		parts[i] = part
	}
	return parts, true
}

func initRangeContainsVersion(versionRange, version string) bool {
	want, ok := initVersionParts(version)
	if !ok || strings.TrimSpace(versionRange) == "" {
		return false
	}
	for _, clause := range strings.Split(versionRange, ",") {
		clause = strings.TrimSpace(clause)
		operator := ""
		for _, candidate := range []string{">=", "<=", ">", "<", "="} {
			if strings.HasPrefix(clause, candidate) {
				operator = candidate
				break
			}
		}
		if operator == "" {
			return false
		}
		bound, ok := initVersionParts(strings.TrimSpace(strings.TrimPrefix(clause, operator)))
		if !ok {
			return false
		}
		comparison := 0
		for i := range want {
			if want[i] < bound[i] {
				comparison = -1
				break
			}
			if want[i] > bound[i] {
				comparison = 1
				break
			}
		}
		if !(operator == ">=" && comparison >= 0 || operator == "<=" && comparison <= 0 || operator == ">" && comparison > 0 || operator == "<" && comparison < 0 || operator == "=" && comparison == 0) {
			return false
		}
	}
	return true
}

func Test_APP_041_APP_042_APP_043_APP_044_APP_045_APP_046_APP_050_APP_053_APP_054_APP_059_GoInit(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	appPut(t, root, "go.mod", "module example.test/initial\n\ngo 1.23\n")
	appPut(t, root, "SYSTEM.md", "# Existing project instructions\n")
	appPut(t, root, "AGENTS.md", "Before the block.\n\n<!-- fireweed:start -->\nold directions\n<!-- fireweed:end -->\n\nAfter the block.\n")
	appPut(t, root, ".gitignore", "custom-entry\n")
	recipe, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "codex")
	assertInitOutput(t, status, stdout, stderr, recipe.Instructions) // APP-046
	config := readInitConfig(t, root)
	if config.Version != 1 || !initRangeContainsVersion(config.Fireweed, assets.Version()) || config.Base != "main" || config.Remote != "origin" || !reflect.DeepEqual(config.Modules, []string{"packages/*"}) {
		t.Errorf("APP-041 configuration header = %+v", config)
	}
	version, ok := initVersionParts(assets.Version())
	if !ok {
		t.Fatalf("APP-041 assets.Version %q is not numeric", assets.Version())
	}
	wantRange := fmt.Sprintf(">=%d.%d, <%d.%d", version[0], version[1], version[0], version[1]+1)
	if config.Fireweed != wantRange {
		t.Errorf("APP-041 binary range = %q; want exact minor window %q", config.Fireweed, wantRange)
	}
	patch := fmt.Sprintf("%d.%d.%d", version[0], version[1], version[2]+1)
	if !initRangeContainsVersion(config.Fireweed, patch) {
		t.Errorf("APP-041 binary range %q excludes later patch %s", config.Fireweed, patch)
	}
	nextMinor := fmt.Sprintf("%d.%d.0", version[0], version[1]+1)
	if initRangeContainsVersion(config.Fireweed, nextMinor) {
		t.Errorf("APP-041 binary range %q accepts next minor %s", config.Fireweed, nextMinor)
	}
	if !reflect.DeepEqual(config.ModuleDefaults.Protected, recipe.ModuleDefaults.Protected) || !reflect.DeepEqual(config.ModuleDefaults.Interface, recipe.ModuleDefaults.Interface) || !reflect.DeepEqual(config.ModuleDefaults.Evals, recipe.ModuleDefaults.Evals) || !reflect.DeepEqual(config.ModuleDefaults.Disposable, recipe.ModuleDefaults.Disposable) || !reflect.DeepEqual(config.Hooks, recipe.Hooks) || !reflect.DeepEqual(config.Protected, recipe.Protected) || !reflect.DeepEqual(config.WorkerWritable, recipe.WorkerWritable) {
		t.Errorf("APP-042 recipe metadata differs: %+v", config)
	}
	if got := readInitFile(t, root, "SYSTEM.md"); got != "# Existing project instructions\n" {
		t.Errorf("APP-043 SYSTEM.md changed: %q", got)
	}
	agents := readInitFile(t, root, "AGENTS.md")
	for _, required := range []string{"Before the block.", "After the block.", "<!-- fireweed:start -->", "<!-- fireweed:end -->", "playbook author-change", "context"} {
		if !strings.Contains(agents, required) {
			t.Errorf("APP-044 AGENTS.md missing %q: %q", required, agents)
		}
	}
	if strings.Contains(agents, "old directions") || strings.Count(agents, "<!-- fireweed:start -->") != 1 {
		t.Errorf("APP-044 managed block was not replaced once: %q", agents)
	}
	if !strings.HasPrefix(agents, "Before the block.\n\n<!-- fireweed:start -->") || !strings.HasSuffix(agents, "<!-- fireweed:end -->\n\nAfter the block.\n") {
		t.Errorf("APP-044 text outside managed block changed: %q", agents)
	}
	for name, want := range recipe.Files {
		if got := readInitFile(t, root, name); got != want {
			t.Errorf("APP-045 %s differs from embedded recipe", name)
		}
	}
	spec := readInitFile(t, root, "packages/app/SPEC.md")
	if !strings.Contains(spec, "name: app") || !strings.Contains(spec, "id_prefix: APP") || !strings.Contains(spec, "depends_on: []") {
		t.Errorf("APP-050 rendered app module metadata: %q", spec)
	}
	if !strings.Contains(readInitFile(t, root, "packages/app/api.go"), "package app") {
		t.Error("APP-050 Go interface template was not rendered")
	}
	for _, directory := range []string{"packages/app/evals", "packages/app/internal"} {
		if info, err := os.Stat(filepath.Join(root, filepath.FromSlash(directory))); err != nil || !info.IsDir() {
			t.Errorf("APP-050 Go layout lacks directory %s: %v", directory, err)
		}
	}
	if _, err := os.Stat(filepath.Join(root, "packages/app/api.go.tmpl")); !os.IsNotExist(err) {
		t.Errorf("APP-050 left template filename: %v", err)
	}
	if got := initCommand(config.Agent.Command); !reflect.DeepEqual(got, []string{"codex", "exec", "{prompt}"}) {
		t.Errorf("APP-053 codex command = %q", got)
	}
	if !config.Audit.Enabled || config.Audit.Command != nil {
		t.Errorf("APP-054 audit = %+v; want enabled with agent fallback", config.Audit)
	}
	var rawConfig map[string]any
	if err := yaml.Unmarshal([]byte(readInitFile(t, root, "fireweed.yaml")), &rawConfig); err == nil {
		if audit, ok := rawConfig["audit"].(map[string]any); ok {
			if _, present := audit["command"]; present {
				t.Error("APP-054 audit.command must be unset for agent fallback")
			}
		}
	}
	if ignore := readInitFile(t, root, ".gitignore"); !strings.Contains(ignore, "custom-entry") || !strings.Contains(ignore, ".fireweed/runs/") {
		t.Errorf("APP-059 .gitignore = %q", ignore)
	}
	if status, stdout, stderr := runAppBinary(t, bin, root, "modules", "--json"); status != 0 || stderr != "" {
		t.Errorf("APP-050 generated Go graph modules status=%d stdout=%q stderr=%q", status, stdout, stderr)
	} else {
		envelope := decodeEnvelope(t, stdout)
		var modules []struct {
			Name      string   `json:"name"`
			IDPrefix  string   `json:"id_prefix"`
			Dir       string   `json:"dir"`
			DependsOn []string `json:"depends_on"`
			Zones     struct {
				Interface  []string `json:"interface"`
				Evals      []string `json:"evals"`
				Disposable []string `json:"disposable"`
			} `json:"zones"`
		}
		if err := json.Unmarshal(envelope["data"], &modules); err != nil || len(modules) != 1 {
			t.Errorf("APP-050 generated modules data=%s: %v", envelope["data"], err)
		} else if module := modules[0]; module.Name != "app" || module.IDPrefix != "APP" || module.Dir != "packages/app" || len(module.DependsOn) != 0 || !reflect.DeepEqual(module.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(module.Zones.Evals, []string{"evals/**"}) || !reflect.DeepEqual(module.Zones.Disposable, []string{"internal/**"}) {
			t.Errorf("APP-050 generated app graph = %+v", module)
		}
	}
	if status, stdout, stderr := runAppBinary(t, bin, root, "check"); status != 0 || stderr != "" {
		t.Errorf("APP-050 generated Go graph check status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
}

func Test_APP_043_APP_044_APP_047_APP_051_NoneInit(t *testing.T) {
func Test_APP_043_APP_044_APP_047_APP_051_NoneInit(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", "claude", "--audit-agent", "codex", "--no-app")
	if status != 0 || stderr != "" {
		t.Errorf("none init status=%d stderr=%q", status, stderr)
	}
	for _, word := range []string{"eval", "boundaries", "setup"} {
		if !strings.Contains(strings.ToLower(stdout), word) {
			t.Errorf("APP-047 instructions omit %q: %q", word, stdout)
		}
	}
	config := readInitConfig(t, root)
	if len(config.Hooks) != 0 || len(config.ModuleDefaults.Protected)+len(config.ModuleDefaults.Interface)+len(config.ModuleDefaults.Evals)+len(config.ModuleDefaults.AppendOnly)+len(config.ModuleDefaults.Preserve)+len(config.ModuleDefaults.Disposable) != 0 {
		t.Errorf("APP-047 none has language defaults or hooks: %+v", config)
	}
	if _, err := os.Stat(filepath.Join(root, ".fireweed/templates")); !os.IsNotExist(err) {
		t.Errorf("APP-047 none created templates: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, ".fireweed/canaries")); !os.IsNotExist(err) {
		t.Errorf("APP-047 none created canaries: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/app")); !os.IsNotExist(err) {
		t.Errorf("APP-051 --no-app created module: %v", err)
	}
	if got := initCommand(config.Agent.Command); !reflect.DeepEqual(got, []string{"claude", "-p", "{prompt}", "--permission-mode", "acceptEdits"}) {
		t.Errorf("APP-053 claude command = %q", got)
	}
	if got := initCommand(config.Audit.Command); !config.Audit.Enabled || !reflect.DeepEqual(got, []string{"codex", "exec", "{prompt}"}) {
		t.Errorf("APP-054 audit preset = %+v", config.Audit)
	}
	if system := readInitFile(t, root, "SYSTEM.md"); strings.TrimSpace(system) == "" {
		t.Error("APP-043 missing editable SYSTEM.md scaffold")
	}
	if agents := readInitFile(t, root, "AGENTS.md"); !strings.Contains(agents, "playbook author-change") || !strings.Contains(agents, "context") {
		t.Errorf("APP-044 missing managed guidance: %q", agents)
	}
}

func Test_APP_050_APP_052_NoneDefaultApp(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	status, _, stderr := runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", "codex")
	if status != 0 || stderr != "" {
		t.Errorf("none init status=%d stderr=%q", status, stderr)
	}
	spec := readInitFile(t, root, "packages/app/SPEC.md")
	for _, part := range []string{"name: app", "id_prefix: APP", "depends_on: []", "## Specs", "## Decisions"} {
		if !strings.Contains(spec, part) {
			t.Errorf("APP-050/052 minimal SPEC lacks %q: %q", part, spec)
		}
	}
	if strings.Contains(spec, "### ") {
		t.Errorf("APP-052 minimal SPEC is not empty: %q", spec)
	}
}

func Test_APP_053_InitFileOnlyWorkerCommand(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	agent := "touch FILE_WORKER_RAN && printf '%s' {prompt_file}"
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", agent, "--no-app")
	if status != 0 || stderr != "" || !strings.Contains(stdout, "eval") {
		t.Errorf("APP-053 file-only worker init: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if status != 0 {
		assertNoInitConfig(t, root)
		return
	}
	config := readInitConfig(t, root)
	if config.Agent.Command != agent {
		t.Errorf("APP-053 file-only worker command = %#v; want exact shell template %q", config.Agent.Command, agent)
	}
	if !config.Audit.Enabled || config.Audit.Command != nil {
		t.Errorf("APP-054 audit fallback changed: %+v", config.Audit)
	}
	if _, err := os.Stat(filepath.Join(root, "FILE_WORKER_RAN")); !os.IsNotExist(err) {
		t.Errorf("init executed worker command: %v", err)
	}
}

func Test_APP_054_InitFileOnlyAuditCommand(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	auditor := "touch FILE_AUDITOR_RAN && printf '%s' {prompt_file}"
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "none", "--agent", "codex", "--audit-agent", auditor, "--no-app")
	if status != 0 || stderr != "" || !strings.Contains(stdout, "eval") {
		t.Errorf("APP-054 file-only audit init: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if status != 0 {
		assertNoInitConfig(t, root)
		return
	}
	config := readInitConfig(t, root)
	if got := initCommand(config.Agent.Command); !reflect.DeepEqual(got, []string{"codex", "exec", "{prompt}"}) {
		t.Errorf("APP-053 worker preset changed: %q", got)
	}
	if !config.Audit.Enabled || config.Audit.Command != auditor {
		t.Errorf("APP-054 file-only audit command = %+v; want exact shell template %q", config.Audit, auditor)
	}
	if _, err := os.Stat(filepath.Join(root, "FILE_AUDITOR_RAN")); !os.IsNotExist(err) {
		t.Errorf("init executed audit command: %v", err)
	}
}

func Test_APP_042_APP_048_APP_049_APP_050_APP_053_APP_054_APP_060_LocalInit(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	recipeDir := filepath.Join(root, "recipe")
	instructions := "# Local setup\nConfigure the eval, boundaries and setup hooks.\n"
	appPut(t, recipeDir, "recipe.yaml", "module_defaults:\n  protected: [schema.lock]\n  interface: [api.txt]\n  evals: [evals/**]\n  disposable: [internal/**]\n  preserve: [data/**]\n  append_only: [log/**]\nhooks:\n  eval: 'touch HOOK_RAN'\n  boundaries: 'touch BOUNDARIES_RAN'\n  setup: 'touch SETUP_RAN'\nprotected: [.fireweed/**, SYSTEM.md]\nworker_writable: [local.lock]\n")
	appPut(t, recipeDir, "recipe.md", instructions)
	appPut(t, recipeDir, "files/assets/banner.txt", "sentinel payload\n")
	appPut(t, recipeDir, "files/.fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n## Decisions\n")
	appPut(t, recipeDir, "files/.fireweed/templates/module/api.txt.tmpl", "package={{package}} name={{name}} prefix={{prefix}}\n")
	agent := "touch AGENT_RAN && printf '%s' {prompt}"
	auditor := "touch AUDITOR_RAN && printf '%s' {prompt}"
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "./recipe", "--agent", agent, "--audit-agent", auditor)
	assertInitOutput(t, status, stdout, stderr, instructions)
	config := readInitConfig(t, root)
	if !reflect.DeepEqual(config.ModuleDefaults.Protected, []string{"schema.lock"}) || !reflect.DeepEqual(config.ModuleDefaults.Interface, []string{"api.txt"}) || !reflect.DeepEqual(config.ModuleDefaults.Evals, []string{"evals/**"}) || !reflect.DeepEqual(config.ModuleDefaults.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(config.ModuleDefaults.Preserve, []string{"data/**"}) || !reflect.DeepEqual(config.ModuleDefaults.AppendOnly, []string{"log/**"}) || config.Hooks["eval"] != "touch HOOK_RAN" || config.Hooks["boundaries"] != "touch BOUNDARIES_RAN" || config.Hooks["setup"] != "touch SETUP_RAN" || !reflect.DeepEqual(config.Protected, []string{".fireweed/**", "SYSTEM.md"}) || !reflect.DeepEqual(config.WorkerWritable, []string{"local.lock"}) {
		t.Errorf("APP-042/048 local metadata differs: %+v", config)
	}
	if got := readInitFile(t, root, "assets/banner.txt"); got != "sentinel payload\n" {
		t.Errorf("APP-049 local payload = %q", got)
	}
	if got := readInitFile(t, root, ".fireweed/recipes/local.md"); got != instructions {
		t.Errorf("APP-049 copied instructions = %q", got)
	}
	if got := readInitFile(t, root, "packages/app/api.txt"); got != "package=app name=app prefix=APP\n" {
		t.Errorf("APP-050 rendered template = %q", got)
	}
	if spec := readInitFile(t, root, "packages/app/SPEC.md"); !strings.Contains(spec, "name: app") || !strings.Contains(spec, "id_prefix: APP") {
		t.Errorf("APP-050 rendered SPEC = %q", spec)
	}
	if config.Agent.Command != agent {
		t.Errorf("APP-053 custom command = %#v", config.Agent.Command)
	}
	if !config.Audit.Enabled || config.Audit.Command != auditor {
		t.Errorf("APP-054 custom audit command = %+v", config.Audit)
	}
	for _, marker := range []string{"HOOK_RAN", "BOUNDARIES_RAN", "SETUP_RAN", "AGENT_RAN", "AUDITOR_RAN"} {
		if _, err := os.Stat(filepath.Join(root, marker)); !os.IsNotExist(err) {
			t.Errorf("APP-060 init executed %s: %v", marker, err)
		}
	}
}

func Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name    string
		args    []string
		code    string
		choices bool
	}{
		{"missing-recipe", []string{"--agent", "codex"}, "U01", true},
		{"missing-agent", []string{"--recipe", "go"}, "U01", true},
		{"missing-value", []string{"--recipe", "go", "--agent"}, "U01", true},
		{"invalid-agent", []string{"--recipe", "go", "--agent", "fake-agent"}, "U01", false},
		{"malformed-option", []string{"--recipe", "go", "--agent", "codex", "--bad"}, "U01", false},
		{"extra-position", []string{"--recipe", "go", "--agent", "codex", "extra"}, "U01", false},
		{"unknown-recipe", []string{"--recipe", "ruby", "--agent", "codex"}, "U02", false},
		{"malformed-document", []string{"--recipe", "./recipe", "--agent", "codex"}, "C03", false},
		{"missing-document", []string{"--recipe", "./recipe", "--agent", "codex"}, "C03", false},
		{"invalid-metadata-type", []string{"--recipe", "./recipe", "--agent", "codex"}, "C12", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			switch tc.name {
			case "malformed-document":
				appPut(t, root, "recipe/recipe.yaml", "module_defaults: [broken\n")
				appPut(t, root, "recipe/recipe.md", "instructions\n")
			case "missing-document":
				appPut(t, root, "recipe/recipe.yaml", "module_defaults: {}\nhooks: {}\nprotected: []\nworker_writable: []\n")
			case "invalid-metadata-type":
				appPut(t, root, "recipe/recipe.yaml", "module_defaults: [wrong-type]\nhooks: {}\nprotected: []\nworker_writable: []\n")
				appPut(t, root, "recipe/recipe.md", "instructions\n")
			}
			args := append([]string{"init"}, tc.args...)
			status, stdout, stderr := runAppBinary(t, bin, root, args...)
			if status != 2 || !strings.HasPrefix(strings.TrimSpace(stdout+stderr), tc.code+":") || !corrective(stdout+stderr) {
				t.Errorf("status=%d output=%q; want corrective %s", status, stdout+stderr, tc.code)
			}
			if tc.choices {
				for _, choice := range []string{"go", "none", "codex", "claude"} {
					if !strings.Contains(stdout+stderr, choice) {
						t.Errorf("APP-055 missing-input diagnostic lacks %q: %q", choice, stdout+stderr)
					}
				}
			}
			assertNoInitConfig(t, root)
			for _, name := range []string{"SYSTEM.md", "AGENTS.md", ".gitignore", ".fireweed", "packages/app"} {
				if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
					t.Errorf("rejected init created %s: %v", name, err)
				}
			}
			jsonArgs := append(args, "--json")
			jsonStatus, jsonOut, jsonErr := runAppBinary(t, bin, root, jsonArgs...)
			if jsonStatus != 2 || jsonErr != "" {
				t.Errorf("APP-061 JSON status=%d stderr=%q", jsonStatus, jsonErr)
			}
			assertFailedJSON(t, jsonOut, tc.code)
		})
	}
}

func Test_APP_057_APP_058_InitConflicts(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct{ name, path, body string }{
		{"config", "fireweed.yaml", "existing config\n"},
		{"recipe", ".fireweed/hooks.go", "existing hook\n"},
		{"module", "packages/app/SPEC.md", "existing spec\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			appPut(t, root, tc.path, tc.body)
			appPut(t, root, "SYSTEM.md", "keep system\n")
			status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "codex")
			if status != 2 || !strings.HasPrefix(strings.TrimSpace(stdout+stderr), "P01:") {
				t.Errorf("status=%d output=%q; want P01", status, stdout+stderr)
			}
			if got := readInitFile(t, root, tc.path); got != tc.body {
				t.Errorf("conflict file changed: %q", got)
			}
			if got := readInitFile(t, root, "SYSTEM.md"); got != "keep system\n" {
				t.Errorf("SYSTEM.md changed: %q", got)
			}
			if tc.name != "config" {
				assertNoInitConfig(t, root)
			}
		})
	}
}

func Test_APP_061_InitJSONSuccess(t *testing.T) {
	bin := buildAppBinary(t)
	root := t.TempDir()
	status, stdout, stderr := runAppBinary(t, bin, root, "init", "--recipe", "go", "--agent", "codex", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("JSON init status=%d stderr=%q", status, stderr)
	}
	envelope := decodeEnvelope(t, stdout)
	var instructions string
	if err := json.Unmarshal(envelope["data"], &instructions); err != nil {
		t.Errorf("data is not instruction string: %v", err)
	}
	recipe, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	if instructions != recipe.Instructions {
		t.Errorf("JSON instructions differ from recipe: %q", instructions)
	}
	var diagnostics []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil || len(diagnostics) != 0 {
		t.Errorf("diagnostics=%s; want []", envelope["diagnostics"])
	}
}

exec
/bin/zsh -lc "cat packages/assets/internal/recipe/recipe.go; cat packages/assets/internal/recipe/files/{api.go.tmpl,SPEC.md,typescript-pnpm-SPEC.md}; sed -n '1,250p' packages/app/evals/new_module_cli_test.go; sed -n '315,454p' packages/app/evals/doctor_setup_test.go; rg -n 'init|agents-md' fireweed-design-v1.md | tail -28" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 472ms:
package recipe

import (
	"embed"
	"fmt"

	"github.com/lurkshark/fireweed/packages/assets/model"
	"github.com/lurkshark/fireweed/packages/contracts"
)

//go:embed files/*
var bundle embed.FS

func BuiltinRecipe(name string) (model.Recipe, error) {
	if name != "go" && name != "typescript-pnpm" && name != "python-uv" {
		return model.Recipe{}, fmt.Errorf("unknown recipe %q", name)
	}
	if name != "go" {
		return additionalRecipe(name)
	}
	paths := map[string]string{
		".fireweed/recipes/go.md":                "files/go.md",
		".fireweed/hooks.go":                     "files/hooks.go",
		".fireweed/templates/module/SPEC.md":     "files/SPEC.md",
		".fireweed/templates/module/api.go.tmpl": "files/api.go.tmpl",
		".fireweed/canaries/canaries.yaml":       "files/canaries.yaml",
	}
	files := make(map[string]string, len(paths))
	for output, resource := range paths {
		data, err := bundle.ReadFile(resource)
		if err != nil {
			return model.Recipe{}, fmt.Errorf("read Go recipe %s: %w", resource, err)
		}
		files[output] = string(data)
	}
	return model.Recipe{
		Name:         "go",
		Instructions: files[".fireweed/recipes/go.md"],
		ModuleDefaults: contracts.Zones{
			Interface:  []string{"api.go"},
			Evals:      []string{"evals/**"},
			Disposable: []string{"internal/**"},
		},
		Hooks: map[string]any{
			"eval":       "go run .fireweed/hooks.go eval {module_dir} {report}",
			"boundaries": "go run .fireweed/hooks.go boundaries",
			"setup":      "go mod download",
		},
		Protected:      []string{".fireweed/**", "fireweed.yaml", "AGENTS.md", "SYSTEM.md"},
		WorkerWritable: []string{"go.sum"},
		Files:          files,
	}, nil
}

func additionalRecipe(name string) (model.Recipe, error) {
	paths := map[string]string{
		".fireweed/recipes/" + name + ".md": "files/" + name + ".md",
		".fireweed/templates/module/SPEC.md": "files/" + name + "-SPEC.md",
		".fireweed/canaries/canaries.yaml": "files/" + name + "-canaries.yaml",
	}
	var zones contracts.Zones
	var hooks map[string]any
	var writable []string
	if name == "typescript-pnpm" {
		paths[".fireweed/hooks.mjs"] = "files/typescript-pnpm-hooks.mjs"
		paths[".fireweed/templates/module/package.json.tmpl"] = "files/typescript-pnpm-package.json.tmpl"
		paths[".fireweed/templates/module/src/index.ts.tmpl"] = "files/typescript-pnpm-index.ts.tmpl"
		zones = contracts.Zones{Interface: []string{"src/index.ts"}, Evals: []string{"evals/**"}, Disposable: []string{"src/internal/**"}}
		hooks = map[string]any{"eval": "pnpm exec node .fireweed/hooks.mjs eval {module_dir} {report}", "boundaries": "pnpm exec node .fireweed/hooks.mjs boundaries", "setup": "pnpm install --offline"}
		writable = []string{"pnpm-lock.yaml"}
	} else {
		paths[".fireweed/hooks.py"] = "files/python-uv-hooks.py"
		paths[".fireweed/templates/module/pyproject.toml.tmpl"] = "files/python-uv-pyproject.toml.tmpl"
		paths[".fireweed/templates/module/__init__.py.tmpl"] = "files/python-uv-init.py.tmpl"
		zones = contracts.Zones{Interface: []string{"__init__.py"}, Evals: []string{"evals/**"}, Disposable: []string{"internal/**"}}
		hooks = map[string]any{"eval": "uv run --no-project python .fireweed/hooks.py eval {module_dir} {report}", "boundaries": "uv run --no-project python .fireweed/hooks.py boundaries", "setup": "uv sync --offline"}
		writable = []string{"uv.lock"}
	}
	files := make(map[string]string, len(paths))
	for output, resource := range paths {
		data, err := bundle.ReadFile(resource)
		if err != nil {
			return model.Recipe{}, fmt.Errorf("read %s recipe %s: %w", name, resource, err)
		}
		files[output] = string(data)
	}
	return model.Recipe{Name: name, Instructions: files[".fireweed/recipes/"+name+".md"], ModuleDefaults: zones, Hooks: hooks,
		Protected: []string{".fireweed/**", "fireweed.yaml", "AGENTS.md", "SYSTEM.md"}, WorkerWritable: writable, Files: files}, nil
}
package {{package}}
---
name: {{name}}
id_prefix: {{prefix}}
depends_on: []
---

# {{name}}

## Specs

When naming Go eval tests, a letter immediately before a spec prefix prevents ID matching. Test_GRT_001_Name matches GRT-001; TestGRT001 does not match GRT-001.

## Decisions
---
name: {{name}}
id_prefix: {{prefix}}
depends_on: []
---

# {{name}}

Module directory: {{module_dir}}

## Specs

## Decisions
package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
)

func newModuleTemplates(t *testing.T, root string) {
	t.Helper()
	recipe, err := assets.BuiltinRecipe("go")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"SPEC.md", "api.go.tmpl"} {
		path := ".fireweed/templates/module/" + name
		appPut(t, root, path, recipe.Files[path])
	}
}

func newModuleAbsent(t *testing.T, root, relative string) {
	t.Helper()
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(relative))); !os.IsNotExist(err) {
		t.Errorf("rejected request wrote %s or stat failed: %v", relative, err)
	}
}

func newModuleFailure(t *testing.T, bin, dir string, args []string, code string) {
	t.Helper()
	for _, jsonMode := range []bool{false, true} {
		call := append([]string(nil), args...)
		if jsonMode {
			call = append(call, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, dir, call...)
		if status != 2 {
			t.Errorf("%v: status=%d, want 2; stdout=%q stderr=%q", call, status, stdout, stderr)
		}
		if jsonMode {
			if stderr != "" {
				t.Errorf("%v: JSON stderr=%q, want empty", call, stderr)
			}
			message := assertFailedJSON(t, stdout, code)
			if !corrective(message) {
				t.Errorf("%v: diagnostic %q needs corrective action", call, message)
			}
		} else if !strings.HasPrefix(strings.TrimSpace(stdout+stderr), code+":") || !corrective(stdout+stderr) {
			t.Errorf("%v: diagnostic=%q, want corrective %s", call, stdout+stderr, code)
		}
	}
}

func Test_APP_100_APP_101_APP_104_NewModuleFromNearestProjectTemplates(t *testing.T) {
	bin := buildAppBinary(t)
	outer := appProjectFixture(t)
	newModuleTemplates(t, outer)
	appPut(t, outer, ".fireweed/templates/module/README.txt.tmpl", "outer template\n")

	inner := filepath.Join(outer, "nested-project")
	appPut(t, inner, "fireweed.yaml", strings.Replace(appFixtureConfig, "packages/*", "components/*", 1))
	appPut(t, inner, "components/base/SPEC.md", "---\nname: base\nid_prefix: BAS\ndepends_on: []\n---\n# base\n\n## Specs\n")
	appPut(t, inner, "go.mod", "module example.test/nearest\n\ngo 1.23\n")
	newModuleTemplates(t, inner)
	appPut(t, inner, ".fireweed/templates/module/SPEC.md", "---\nname: {{name}}\nid_prefix: {{prefix}}\ndepends_on: []\n---\n# {{name}}\n\n## Specs\n\n## Decisions\n\nProject-owned SPEC marker.\n")
	appPut(t, inner, ".fireweed/templates/module/api.go.tmpl", "package {{package}}\n// Project-owned API: {{name}} {{prefix}}\n")
	appPut(t, inner, ".fireweed/templates/module/README.txt.tmpl", "name={{name}}\nprefix={{prefix}}\npackage={{package}}\nmodule_dir={{module_dir}}\ngo_module={{go_module}}\n")
	from := filepath.Join(inner, "components", "base", "deep")
	if err := os.MkdirAll(from, 0755); err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr := runAppBinary(t, bin, from, "new-module", "quartz-core", "--prefix", "QTZ")
	if status != 0 || stderr != "" || strings.TrimSpace(stdout) != "components/quartz-core" {
		t.Errorf("plain creation: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	newModuleAbsent(t, outer, "packages/quartz-core")
	if got := readInitFile(t, inner, "components/quartz-core/SPEC.md"); !strings.Contains(got, "name: quartz-core\nid_prefix: QTZ") || !strings.Contains(got, "Project-owned SPEC marker.") {
		t.Errorf("rendered SPEC identity/project template: %q", got)
	}
	if got := readInitFile(t, inner, "components/quartz-core/api.go"); got != "package quartz_core\n// Project-owned API: quartz-core QTZ\n" {
		t.Errorf("rendered API=%q", got)
	}
	if got := readInitFile(t, inner, "components/quartz-core/README.txt"); got != "name=quartz-core\nprefix=QTZ\npackage=quartz_core\nmodule_dir=components/quartz-core\ngo_module=example.test/nearest\n" {
		t.Errorf("rendered substitutions=%q", got)
	}
	newModuleAbsent(t, inner, "components/quartz-core/api.go.tmpl")
	newModuleAbsent(t, inner, "components/quartz-core/README.txt.tmpl")
	snapshot, err := project.Load(inner)
	if err != nil {
		t.Fatalf("created module is not loadable: %v", err)
	}
	found := false
	for _, module := range snapshot.Modules {
		if module.Name == "quartz-core" {
			found = true
			if module.IDPrefix != "QTZ" || module.Dir != "components/quartz-core" {
				t.Errorf("created graph module=%+v", module)
			}
		}
	}
	if !found {
		t.Errorf("created module absent from project graph: %+v", snapshot.Modules)
	}
	if err := project.ValidateGraph(snapshot.Modules); err != nil {
		t.Errorf("created graph invalid: %v", err)
	}
	status, stdout, stderr = runAppBinary(t, bin, from, "modules", "--json")
	if status != 0 || stderr != "" || !strings.Contains(stdout, `"name":"quartz-core"`) {
		t.Errorf("CLI cannot discover created module: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}

	status, stdout, stderr = runAppBinary(t, bin, from, "new-module", "second", "--prefix", "SEC", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("JSON creation: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if got := commandStringData(t, stdout); got != "components/second" {
		t.Errorf("JSON data=%q, want repository-relative module directory", got)
	}
	if got := readInitFile(t, inner, "components/second/README.txt"); !strings.Contains(got, "module_dir=components/second\n") || !strings.Contains(got, "go_module=example.test/nearest\n") {
		t.Errorf("second module substitutions=%q", got)
	}
}

func Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, module, prefix, target string
		prepare                      func(*testing.T, string)
	}{
		{"existing-name", "alpha", "NEW", "components/alpha", func(t *testing.T, root string) {}},
		{"existing-prefix", "fresh", "ALP", "components/fresh", func(t *testing.T, root string) {}},
		{"occupied-directory", "fresh", "FRS", "components/fresh", func(t *testing.T, root string) {
			appPut(t, root, "components/fresh/SPEC.md", "---\nname: occupied\nid_prefix: OCC\ndepends_on: []\n---\n# occupied\n\n## Specs\n")
			appPut(t, root, "components/fresh/keep.txt", "keep occupied directory\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, `"packages/*"`, `"components/*", "packages/*"`, 1))
			appPut(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# alpha\n\n## Specs\n")
			newModuleTemplates(t, root)
			tc.prepare(t, root)
			if _, err := project.Load(root); err != nil {
				t.Fatalf("conflict fixture must have a valid existing graph: %v", err)
			}
			before := readInitFile(t, root, "packages/alpha/SPEC.md")
			occupiedBefore := ""
			if tc.name == "occupied-directory" {
				occupiedBefore = readInitFile(t, root, "components/fresh/SPEC.md")
			}
			newModuleFailure(t, bin, root, []string{"new-module", tc.module, "--prefix", tc.prefix}, "P01")
			if got := readInitFile(t, root, "packages/alpha/SPEC.md"); got != before {
				t.Errorf("existing module changed: %q", got)
			}
			if tc.name == "occupied-directory" {
				if got := readInitFile(t, root, "components/fresh/SPEC.md"); got != occupiedBefore {
					t.Errorf("occupied module changed: %q", got)
				}
				if got := readInitFile(t, root, "components/fresh/keep.txt"); got != "keep occupied directory\n" {
					t.Errorf("occupied contents changed: %q", got)
				}
				newModuleAbsent(t, root, "components/fresh/api.go")
			} else {
				newModuleAbsent(t, root, tc.target)
			}
		})
	}
}

func Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name, code, target string
		args               []string
		prepare            func(*testing.T, string)
	}{
		{"missing-name", "U01", "packages/new", []string{"new-module", "--prefix", "NEW"}, nil},
		{"missing-prefix", "U01", "packages/new", []string{"new-module", "new"}, nil},
		{"invalid-name", "U01", "packages/new", []string{"new-module", "bad/name", "--prefix", "NEW"}, nil},
		{"invalid-prefix", "U01", "packages/new", []string{"new-module", "new", "--prefix", "bad-prefix"}, nil},
		{"repeated-prefix", "U01", "packages/new", []string{"new-module", "new", "--prefix", "NEW", "--prefix", "AGAIN"}, nil},
		{"extra-argument", "U01", "packages/new", []string{"new-module", "new", "--prefix", "NEW", "extra"}, nil},
		{"unknown-option", "U01", "packages/new", []string{"new-module", "new", "--prefix", "NEW", "--bogus"}, nil},
		{"missing-template", "C03", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			if err := os.RemoveAll(filepath.Join(root, ".fireweed/templates/module")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed-template", "C03", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, ".fireweed/templates/module/SPEC.md", "---\nname: [broken\n---\n")
		}},
		{"invalid-configuration", "C01", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"invalid-existing-graph", "C02", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
		{"no-configured-destination", "C01", "packages/new", []string{"new-module", "new", "--prefix", "NEW"}, func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", strings.Replace(appFixtureConfig, "packages/*", "packages/*/internal", 1))
			appPut(t, root, "packages/alpha/internal/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# alpha\n\n## Specs\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := appProjectFixture(t)
			newModuleTemplates(t, root)
			if tc.prepare != nil {
				tc.prepare(t, root)
			}
			newModuleFailure(t, bin, root, tc.args, tc.code)
			newModuleAbsent(t, root, tc.target)
		})
	}
}

func Test_APP_101_APP_104_NewModuleEmptyGoModule(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	newModuleTemplates(t, root)
	appPut(t, root, ".fireweed/templates/module/README.tmpl", "go_module={{go_module}}\nmodule_dir={{module_dir}}\n")
	status, stdout, stderr := runAppBinary(t, bin, root, "new-module", "empty-mod", "--prefix", "EMP", "--json")
	if status != 0 || stderr != "" || commandStringData(t, stdout) != "packages/empty-mod" {
		t.Errorf("creation without go.mod: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if got := readInitFile(t, root, "packages/empty-mod/README"); got != "go_module=\nmodule_dir=packages/empty-mod\n" {
		t.Errorf("empty go_module substitution=%q", got)
	}
}

func Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite(t *testing.T) {
	bin := buildAppBinary(t)
	root := appProjectFixture(t)
	nested := appNestedDir(t, root)
	system := "# Existing system\n"
	appPut(t, root, "SYSTEM.md", system)
	before := "Before.\n\n<!-- fireweed:start -->\nold guidance\n<!-- fireweed:end -->\n\nAfter.\n"
	appPut(t, root, "AGENTS.md", before)
	status, stdout, stderr := runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 0 || strings.TrimSpace(stdout) == "" {
		t.Errorf("APP-068/072 status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	agents := readInitFile(t, root, "AGENTS.md")
	if agents != "Before.\n\n"+doctorCLIManagedBlock+"\n\nAfter.\n" {
		t.Errorf("APP-069 managed guidance=%q", agents)
	}
	if got := readInitFile(t, root, "SYSTEM.md"); got != system {
		t.Errorf("APP-068 changed SYSTEM.md: %q", got)
	}
	status, output, stderr := runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
	if status != 0 || stderr != "" {
		t.Errorf("APP-073 status=%d stderr=%q", status, stderr)
	}
	envelope := decodeEnvelope(t, output)
	var confirmation string
	if err := json.Unmarshal(envelope["data"], &confirmation); err != nil || confirmation != strings.TrimSpace(stdout) {
		t.Errorf("APP-073 confirmation=%s: %v", envelope["data"], err)
	}
	var diags []json.RawMessage
	if err := json.Unmarshal(envelope["diagnostics"], &diags); err != nil || diags == nil || len(diags) != 0 {
		t.Errorf("APP-073 diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	outerGuidance := agents
	closer := filepath.Join(root, "closer")
	appPut(t, closer, "fireweed.yaml", appFixtureConfig)
	appPut(t, closer, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n")
	appPut(t, closer, "AGENTS.md", "Closer before.\n<!-- fireweed:start -->\nstale\n<!-- fireweed:end -->\nCloser after.\n")
	status, stdout, stderr = runAppBinary(t, bin, appNestedDir(t, closer), "agents-md", "--write")
	if status != 0 || stderr != "" {
		t.Errorf("APP-068 closer write status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	if got := readInitFile(t, closer, "AGENTS.md"); got != "Closer before.\n"+doctorCLIManagedBlock+"\nCloser after.\n" {
		t.Errorf("APP-068 closer guidance=%q", got)
	}
	if got := readInitFile(t, root, "AGENTS.md"); got != outerGuidance {
		t.Errorf("APP-068 farther project guidance changed: %q", got)
	}

	for _, tc := range []struct {
		name, initial string
		absent        bool
	}{
		{"no-block", "Personal instructions.\n", false},
		{"missing-file", "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.absent {
				if err := os.Remove(filepath.Join(root, "AGENTS.md")); err != nil {
					t.Fatal(err)
				}
			} else {
				appPut(t, root, "AGENTS.md", tc.initial)
			}
			status, stdout, stderr := runAppBinary(t, bin, nested, "agents-md", "--write")
			if status != 0 || strings.TrimSpace(stdout) == "" {
				t.Errorf("APP-070/072 status=%d stdout=%q stderr=%q", status, stdout, stderr)
			}
			got := readInitFile(t, root, "AGENTS.md")
			if !strings.HasPrefix(got, tc.initial) || strings.Count(got, "<!-- fireweed:start -->") != 1 || strings.Count(got, "<!-- fireweed:end -->") != 1 || !strings.Contains(got, doctorCLIManagedBlock) {
				t.Errorf("APP-070 content=%q", got)
			}
		})
	}

	for _, malformed := range []string{"Personal.\n<!-- fireweed:start -->\nold\n", "Personal.\n<!-- fireweed:end -->\n"} {
		appPut(t, root, "AGENTS.md", malformed)
		status, stdout, stderr := runAppBinary(t, bin, nested, "agents-md", "--write")
		if status != 2 || !strings.Contains(stdout+stderr, "P01") || !corrective(stdout+stderr) || readInitFile(t, root, "AGENTS.md") != malformed {
			t.Errorf("APP-071/072 malformed status=%d output=%q", status, stdout+stderr)
		}
		status, output, stderr := runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
		if status != 2 || stderr != "" {
			t.Errorf("APP-073 malformed JSON status=%d stderr=%q", status, stderr)
		}
		assertFailedJSON(t, output, "P01")
	}

	for _, args := range [][]string{{"agents-md"}, {"agents-md", "--bad"}, {"agents-md", "--write", "extra"}, {"doctor", "--write"}} {
		status, stdout, stderr := runAppBinary(t, bin, nested, args...)
		if status != 2 || !strings.Contains(stdout+stderr, "U01") || !corrective(stdout+stderr) {
			t.Errorf("APP-072 usage %v: status=%d output=%q", args, status, stdout+stderr)
		}
		status, output, errOut := runAppBinary(t, bin, nested, append(args, "--json")...)
		if status != 2 || errOut != "" {
			t.Errorf("APP-073 usage JSON %v: status=%d stderr=%q", args, status, errOut)
		}
		assertFailedJSON(t, output, "U01")
	}
	if got := readInitFile(t, root, "AGENTS.md"); got != "Personal.\n<!-- fireweed:end -->\n" {
		t.Errorf("APP-072 usage changed guidance: %q", got)
	}
	if err := os.Remove(filepath.Join(root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(root, "AGENTS.md"), 0755); err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr = runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 2 || !strings.Contains(stdout+stderr, "P01") || !corrective(stdout+stderr) {
		t.Errorf("APP-072 operational failure status=%d output=%q", status, stdout+stderr)
	}
	status, output, stderr = runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
	if status != 2 || stderr != "" {
		t.Errorf("APP-073 operational JSON status=%d stderr=%q", status, stderr)
	}
	assertFailedJSON(t, output, "P01")
	if err := os.Remove(filepath.Join(root, "AGENTS.md")); err != nil {
		t.Fatal(err)
	}

	appPut(t, root, "fireweed.yaml", "[invalid yaml\n")
	status, stdout, stderr = runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 2 || !strings.Contains(stdout+stderr, "C01") {
		t.Errorf("APP-072 config status=%d output=%q", status, stdout+stderr)
	}
	status, output, stderr = runAppBinary(t, bin, nested, "agents-md", "--write", "--json")
	if status != 2 || stderr != "" {
		t.Errorf("APP-073 config JSON status=%d stderr=%q", status, stderr)
	}
	assertFailedJSON(t, output, "C01")
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
	status, stdout, stderr = runAppBinary(t, bin, nested, "agents-md", "--write")
	if status != 2 || !strings.Contains(stdout+stderr, "C01") {
		t.Errorf("APP-072 missing config status=%d output=%q", status, stdout+stderr)
	}
}
34:A project created with `fireweed init` has one config file, one project document, a change queue, a protected `.fireweed/` directory and its modules. Only the module layout is a default; everything else is fixed.
51:  app/                   # composition-root module created by init
288:fireweed keeps a finding only if the ID exists and the quote is an exact substring of that spec. If `audit.command` is not set, the worker's command is used and doctor warns that the auditor shares the worker's blind spots; `init` takes an `--audit-agent` flag.
374:**Recipes** package an opinionated setup for one stack: a document the agent follows, module templates, canary snippets and default hook commands. The first three are TypeScript with pnpm workspaces, Python with uv, and Go, covering three enforcement styles: package exports, an import linter, and Go's internal directories plus a dependency linter such as depguard. Recipes are embedded in the binary, and `init` copies the chosen recipe into `.fireweed/`. After that the copy belongs to the project, because the agent adapts it to the repository and an automatic upgrade would break that work. A third-party or homemade recipe is a local directory passed as `--recipe ./path`. For an unsupported stack, `--recipe none` writes the minimum and prints the hook contract. The agent then implements the hooks and also writes the module templates and canaries, without which doctor cannot run.
427:| Setup | `init --recipe <r> --agent <a> [--no-app]` | Create config, `SYSTEM.md`, the `AGENTS.md` block, the recipe copy and the `app` module |
446:| Maintenance | `gc`, `migrate`, `agents-md --write` | Prune refs, upgrade file formats, refresh the `AGENTS.md` block |
457:## Distribution and init
517:1. The agent runs `fireweed init --recipe typescript-pnpm --agent claude`. This writes the config, `SYSTEM.md`, the `AGENTS.md` block, the recipe copy and an `app` composition-root module (`--no-app` skips it), then prints the recipe's instructions.
524:**AGENTS.md.** `init` writes a short block between markers telling any agent that the project uses fireweed and pointing it at `fireweed playbook author-change` and `fireweed context`. `fireweed agents-md --write` regenerates only that block, and doctor warns when it is out of date.
566:| Attempt limits only inside `next` and `run` | Replaced | An agent driving `start`, `verify` and `done` itself looped evals and audit indefinitely; `audit_rounds` lives in the change file so the limit holds however a change is driven |

exec
/bin/zsh -lc "cat > packages/app/internal/entry/setup.go <<'EOF'
package entry

import (
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"maps\"
	\"os\"
	\"path/filepath\"
	\"regexp\"
	\"slices\"
	\"strconv\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/assets\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"gopkg.in/yaml.v3\"
)

const recipeChoices = \"go, typescript-pnpm, python-uv, none, or ./directory\"
const startMarker = \"<"'!-- fireweed:start -->"
const endMarker = "<!-- fireweed:end -->"
const managedBlock = startMarker + "'"\\nThis project uses fireweed. Specs and change files are the source of truth.\\n\" +
	\"Run "'`fireweed playbook author-change` before authoring changes and'"\\n\" +
	\""'`fireweed context <change>` before each step.'"\\n\" + endMarker

const systemScaffold = "'`# Project instructions

Describe the project, its module responsibilities, and the conventions workers
should follow. Document setup, evals and boundary checks here.
`

const hookContract = `# Fireweed hook contract

- eval receives {module_dir} and {report}. Write JUnit XML to {report} and
  return a nonzero status when tests fail.
- boundaries checks module dependency and interface rules. Report violations
  by module and return a nonzero status on failure.
- setup prepares a scratch checkout for the eval and boundaries hooks.
`

const moduleTemplateDir = ".fireweed/templates/module"
const minimalSpec = "---'"\\nname: {{name}}\\nid_prefix: {{prefix}}\\ndepends_on: []\\n---\\n\\n# {{name}}\\n\\n## Specs\\n\\n## Decisions\\n\"

func commandTemplate(value string) (any,error) {
	switch value {
	case \"codex\": return []string{\"codex\",\"exec\",\"{prompt}\"},nil
	case \"claude\": return []string{\"claude\",\"-p\",\"{prompt}\",\"--permission-mode\",\"acceptEdits\"},nil
	}
	if strings.Contains(value,\"{prompt}\") || strings.Contains(value,\"{prompt_file}\") { return value,nil }
	return nil,fmt.Errorf(\"invalid agent command %q\",value)
}

func initUsage(detail string) response {
	return usageError(detail+\". Usage: fireweed init --recipe NAME --agent COMMAND [--audit-agent COMMAND] [--no-app]. Valid recipes: \"+recipeChoices+\". Valid agents: codex, claude, or a shell command containing {prompt} or {prompt_file}.\")
}

func appendLine(text,line string) string {
	if text"'!="" && !strings.HasSuffix(text,"'"\\n\") { text+=\"\\n\" }
	return text+line+\"\\n\"
}

func managedText(text string) (string,error) {
	start,end:=strings.Index(text,startMarker),strings.Index(text,endMarker)
	if start<0 && end<0 { return appendLine(text,managedBlock),nil }
	if start<0 || end<start || strings.Count(text,startMarker)"'!=1 || strings.Count(text,endMarker)!=1 {
		return "",fmt.Errorf("AGENTS.md has invalid managed-block markers; use one paired start and end marker")
	}
	return text[:start]+managedBlock+text[end+len(endMarker):],nil
}

type editableFile struct {
	name string
	old string
	mode fs.FileMode
	exists bool
}

func readEditable(root,name string) (editableFile,error) {
	file:=editableFile{name:name,mode:0644}
	path:=filepath.Join(root,name)
	info,err:=os.Lstat(path)
	if os.IsNotExist(err) { return file,nil }
	if err!=nil { return file,err }
	if !info.Mode().IsRegular() { return file,fmt.Errorf("%s must be a regular file",name) }
	file.mode,file.exists=info.Mode().Perm(),true
	data,err:=os.ReadFile(path)
	file.old=string(data)
	return file,err
}

func writeAgents(root string) response {
	file,err:=readEditable(root,"AGENTS.md")
	if err!=nil { return reject(2,"P01",err) }
	content,err:=managedText(file.old)
	if err!=nil { return reject(2,"P01",err) }
	if err=writeSetup(root,map[string]string{file.name:content},map[string]editableFile{file.name:file});err!=nil { return reject(2,"P01",err) }
	return response{data:"Updated AGENTS.md.",plain:"Updated AGENTS.md.'"\\n\"}
}

func versionRange() (string,error) {
	parts:=strings.Split(assets.Version(),\".\")
	if len(parts)<2 { return \"\",fmt.Errorf(\"invalid binary version %q\",assets.Version()) }
	major,err:=strconv.Atoi(parts[0])
	if err"'!=nil { return "",err }
	minor,err:=strconv.Atoi(parts[1])
	if err!=nil { return "",err }
	return fmt.Sprintf(">=%d.%d, <%d.%d",major,minor,major,minor+1),nil
}

func initialize(c command) response {
	if !c.has("--recipe") || !c.has("--agent") { return initUsage("--recipe and --agent are required") }
	agent,err:=commandTemplate(c.flags["--agent"])
	if err!=nil { return initUsage(err.Error()) }
	audit:=map[string]any{"enabled":true}
	if c.has("--audit-agent") {
		audit["command"],err=commandTemplate(c.flags["--audit-agent"])
		if err!=nil { return initUsage(err.Error()) }
	}
	root,err:=os.Getwd()
	if err!=nil { return reject(2,"P01",err) }
	if err=vacant(root,"fireweed.yaml");err!=nil { return reject(2,"P01",err) }
	recipe,err:=loadRecipe(c.flags["--recipe"])
	if err!=nil { return reject(2,"U02",err) }
	version,err:=versionRange()
	if err!=nil { return reject(2,"P01",err) }
	z:=recipe.ModuleDefaults
	defaults:=map[string][]string{}
	for key,value:=range map[string][]string{"protected":z.Protected,"evals":z.Evals,"interface":z.Interface,"append_only":z.AppendOnly,"preserve":z.Preserve,"disposable":z.Disposable} {
		if value!=nil { defaults[key]=value }
	}
	config,err:=yaml.Marshal(map[string]any{
		"fireweed":version,"version":1,"base":"main","remote":"origin","modules":[]string{"packages/*"},
		"module_defaults":defaults,"hooks":recipe.Hooks,"protected":recipe.Protected,"worker_writable":recipe.WorkerWritable,
		"agent":map[string]any{"command":agent},"audit":audit,
	})
	if err!=nil { return reject(2,"P01",err) }
	files:=make(map[string]string)
	maps.Copy(files,recipe.Files)
	for _,reserved:=range []string{"fireweed.yaml","AGENTS.md","SYSTEM.md",".gitignore"} {
		if _,exists:=files[reserved];exists { return reject(2,"P01",fmt.Errorf("recipe payload conflicts with %s",reserved)) }
	}
	var dirs []string
	if !c.has("--no-app") {
		templates:=map[string]string{}
		for path,body:=range recipe.Files {
			if name,ok:=strings.CutPrefix(path,moduleTemplateDir+"/");ok { templates[name]=body }
		}
		if len(templates)==0 { templates["SPEC.md"]=minimalSpec }
		goModule,err:=goModuleName(root)
		if err!=nil { return reject(2,"P01",err) }
		generated,_,err:=renderModule(templates,"app","APP","packages/app",goModule)
		if err!=nil { return reject(2,"C03",err) }
		for path,body:=range generated {
			if _,exists:=files[path];exists { return reject(2,"P01",fmt.Errorf("recipe payload conflicts with generated %s",path)) }
			files[path]=body
		}
		dirs=literalZoneDirs("packages/app",append(slices.Clone(z.Evals),z.Disposable...))
	}
	edit:=map[string]editableFile{}
	for _,name:=range []string{"AGENTS.md",".gitignore"} {
		file,err:=readEditable(root,name)
		if err!=nil { return reject(2,"P01",err) }
		edit[name]=file
	}
	files["AGENTS.md"],err=managedText(edit["AGENTS.md"].old)
	if err!=nil { return reject(2,"P01",err) }
	ignore:=edit[".gitignore"].old
	if !slices.Contains(strings.Split(strings.ReplaceAll(ignore,"'"\\r\\n\",\"\\n\"),\"\\n\"),\".fireweed/runs/\") {
		ignore=appendLine(ignore,\".fireweed/runs/\")
	}
	files[\".gitignore\"]=ignore
	if _,err:=os.Lstat(filepath.Join(root,\"SYSTEM.md\"));os.IsNotExist(err) {
		files[\"SYSTEM.md\"]=systemScaffold
	} else if err"'!=nil { return reject(2,"P01",err) }
	files["fireweed.yaml"]=string(config)
	if err:=writeSetup(root,files,edit,dirs...);err!=nil { return reject(2,"P01",err) }
	return stringResult(recipe.Instructions)
}

func localRecipeError(code string,err error) error {
	return contracts.CodedError{Code:code,Cause:err}
}

func loadRecipe(name string) (assets.Recipe,error) {
	if name=="none" { return assets.Recipe{Name:name,Instructions:hookContract},nil }
	if !strings.HasPrefix(name,"./") && !strings.HasPrefix(name,"../") && !filepath.IsAbs(name) { return assets.BuiltinRecipe(name) }
	meta,err:=os.ReadFile(filepath.Join(name,"recipe.yaml"))
	if err!=nil { return assets.Recipe{},localRecipeError("C03",err) }
	var doc yaml.Node
	if err=yaml.Unmarshal(meta,&doc);err!=nil { return assets.Recipe{},localRecipeError("C03",fmt.Errorf("parse recipe.yaml: %w",err)) }
	var raw map[string]any
	if len(doc.Content)==0 || doc.Content[0].Kind!=yaml.MappingNode { return assets.Recipe{},localRecipeError("C12",fmt.Errorf("recipe.yaml must be a mapping")) }
	if err=doc.Decode(&raw);err!=nil { return assets.Recipe{},localRecipeError("C03",fmt.Errorf("parse recipe.yaml: %w",err)) }
	recipe:=assets.Recipe{Name:"local"}
	if err=recipeMetadata(&recipe,raw);err!=nil { return assets.Recipe{},localRecipeError("C12",err) }
	body,err:=os.ReadFile(filepath.Join(name,"recipe.md"))
	if err!=nil { return assets.Recipe{},localRecipeError("C03",err) }
	recipe.Instructions=string(body)
	recipe.Files,err=readTree(filepath.Join(name,"files"))
	if err!=nil {
		if !errors.Is(err,fs.ErrNotExist) { return assets.Recipe{},localRecipeError("C03",err) }
		recipe.Files=map[string]string{}
	}
	const instructionsPath=".fireweed/recipes/local.md"
	if _,exists:=recipe.Files[instructionsPath];exists { return assets.Recipe{},localRecipeError("P01",fmt.Errorf("recipe payload conflicts with %s",instructionsPath)) }
	recipe.Files[instructionsPath]=recipe.Instructions
	return recipe,nil
}

func stringList(raw any) ([]string,bool) {
	items,ok:=raw.([]any)
	if !ok { return nil,false }
	result:=make([]string,0,len(items))
	for _,v:=range items {
		text,ok:=v.(string)
		if !ok { return nil,false }
		result=append(result,text)
	}
	return result,true
}

func recipeMetadata(recipe *assets.Recipe,raw map[string]any) error {
	defaults:=map[string]any{}
	if v,exists:=raw["module_defaults"];exists {
		var ok bool
		defaults,ok=v.(map[string]any)
		if !ok { return fmt.Errorf("recipe.yaml module_defaults must be a mapping") }
	}
	z:=&recipe.ModuleDefaults
	for _,field:=range []struct{
		from map[string]any
		key string
		to *[]string
	}{
		{raw,"protected",&recipe.Protected},{raw,"worker_writable",&recipe.WorkerWritable},
		{defaults,"protected",&z.Protected},{defaults,"evals",&z.Evals},{defaults,"interface",&z.Interface},
		{defaults,"append_only",&z.AppendOnly},{defaults,"preserve",&z.Preserve},{defaults,"disposable",&z.Disposable},
	} {
		if v,exists:=field.from[field.key];exists {
			list,ok:=stringList(v)
			if !ok { return fmt.Errorf("recipe.yaml %s must be a list of strings",field.key) }
			*field.to=list
		}
	}
	if v,exists:=raw["hooks"];exists {
		var ok bool
		recipe.Hooks,ok=v.(map[string]any)
		if !ok { return fmt.Errorf("recipe.yaml hooks must be a mapping") }
		for name,value:=range recipe.Hooks {
			if _,ok:=value.(string);ok { continue }
			if argv,ok:=stringList(value);!ok || len(argv)==0 { return fmt.Errorf("recipe.yaml hook %s must be a shell string or argv list",name) }
		}
	}
	return nil
}

func readTree(dir string) (map[string]string,error) {
	files:=map[string]string{}
	err:=filepath.WalkDir(dir,func(path string,entry fs.DirEntry,err error) error {
		if err!=nil { return err }
		if entry.IsDir() { return nil }
		if !entry.Type().IsRegular() { return fmt.Errorf("%s must be a regular file",path) }
		rel,err:=filepath.Rel(dir,path)
		if err!=nil { return err }
		data,err:=os.ReadFile(path)
		if err!=nil { return err }
		files[filepath.ToSlash(rel)]=string(data)
		return nil
	})
	return files,err
}

func vacant(root,relative string) error {
	if _,err:=os.Lstat(filepath.Join(root,relative));os.IsNotExist(err) { return nil } else if err!=nil { return err }
	return fmt.Errorf("%s already exists",relative)
}

func safeRelative(path string) bool {
	return filepath.IsLocal(path) && filepath.Clean(path)==path && path!="."
}

func ensureDirectory(root,dir string,created *[]string) error {
	if dir=="." { return nil }
	if !safeRelative(dir) { return fmt.Errorf("invalid project directory %q",dir) }
	if err:=ensureDirectory(root,filepath.Dir(dir),created);err!=nil { return err }
	full:=filepath.Join(root,dir)
	info,err:=os.Lstat(full)
	if os.IsNotExist(err) {
		if created==nil { return nil }
		if err=os.Mkdir(full,0755);err!=nil { return err }
		*created=append(*created,full)
		return nil
	}
	if err!=nil { return err }
	if !info.IsDir() { return fmt.Errorf("%s is not a project directory",dir) }
	return nil
}

func replaceFile(path,body string,mode fs.FileMode) error {
	file,err:=os.CreateTemp(filepath.Dir(path),".fireweed-write-*")
	if err!=nil { return err }
	defer os.Remove(file.Name())
	if err=file.Chmod(mode);err!=nil { _=file.Close();return err }
	_,writeErr:=file.WriteString(body)
	if err=errors.Join(writeErr,file.Close());err!=nil { return err }
	return os.Rename(file.Name(),path)
}

func writeSetup(root string,files map[string]string,editable map[string]editableFile,dirs ...string) (err error) {
	names:=slices.Sorted(maps.Keys(files))
	for _,name:=range names {
		if !safeRelative(name) { return fmt.Errorf("invalid generated path %q",name) }
		if err:=ensureDirectory(root,filepath.Dir(name),nil);err!=nil { return err }
		if old,ok:=editable[name];!ok || !old.exists {
			if err:=vacant(root,name);err!=nil { return err }
		}
		for dir:=filepath.Dir(name);dir!=".";dir=filepath.Dir(dir) {
			if _,conflicts:=files[dir];conflicts { return fmt.Errorf("generated file %s also needs to be a directory",dir) }
		}
	}
	for _,dir:=range dirs {
		if err:=ensureDirectory(root,dir,nil);err!=nil { return err }
	}
	// Publish the configuration only after all its companion files exist.
	if i:=slices.Index(names,"fireweed.yaml");i>=0 {
		names=append(slices.Delete(names,i,i+1),"fireweed.yaml")
	}
	var created,createdDirs []string
	var replaced []editableFile
	defer func(){
		if err==nil { return }
		for _,file:=range replaced { err=errors.Join(err,replaceFile(filepath.Join(root,file.name),file.old,file.mode)) }
		for _,file:=range created { err=errors.Join(err,os.Remove(file)) }
		for i:=len(createdDirs)-1;i>=0;i-- { err=errors.Join(err,os.Remove(createdDirs[i])) }
	}()
	for _,dir:=range dirs {
		if err=ensureDirectory(root,dir,&createdDirs);err!=nil { return err }
	}
	for _,name:=range names {
		if err=ensureDirectory(root,filepath.Dir(name),&createdDirs);err!=nil { return err }
		path:=filepath.Join(root,name)
		if old:=editable[name];old.exists {
			if old.old==files[name] { continue }
			if err=replaceFile(path,files[name],old.mode);err!=nil { return err }
			replaced=append(replaced,old)
			continue
		}
		file,err:=os.OpenFile(path,os.O_CREATE|os.O_EXCL|os.O_WRONLY,0644)
		if err!=nil { return err }
		created=append(created,path)
		_,writeErr:=file.WriteString(files[name])
		if err=errors.Join(writeErr,file.Close());err!=nil { return err }
	}
	return nil
}

func literalZoneDirs(module string,patterns []string) []string {
	var dirs []string
	for _,pattern:=range patterns {
		if dir,ok:=strings.CutSuffix(pattern,"/**");ok && safeRelative(dir) && !strings.ContainsAny(dir,"*?[") {
			dirs=append(dirs,filepath.Join(module,dir))
		}
	}
	return dirs
}

func goModuleName(root string) (string,error) {
	data,err:=os.ReadFile(filepath.Join(root,"go.mod"))
	if os.IsNotExist(err) { return "",nil }
	if err!=nil { return "",err }
	for _,line:=range strings.Split(string(data),"'"\\n\") {
		fields:=strings.Fields(line)
		if len(fields)>=2 && fields[0]==\"module\" {
			if value,err:=strconv.Unquote(fields[1]);err==nil { return value,nil }
			return fields[1],nil
		}
	}
	return \"\",nil
}

var moduleName=regexp.MustCompile("'`''^[a-z][a-z0-9]*(-[a-z0-9]+)*$`)
var modulePrefix=regexp.MustCompile(`''^[A-Z][A-Z0-9]*$`)

func moduleDestination(patterns []string,name string) string {
	for _,pattern:=range patterns {
		if filepath.Base(pattern)!="*" { continue }
		parent:=filepath.Dir(pattern)
		if parent!="." && (!safeRelative(parent) || strings.ContainsAny(parent,"*?[")) { continue }
		dir:=filepath.ToSlash(filepath.Join(parent,name))
		if contracts.MatchGlob(pattern,dir) { return dir }
	}
	return ""
}

func renderModule(templates map[string]string,name,prefix,dir,goModule string) (map[string]string,project.Module,error) {
	files:=map[string]string{}
	render:=strings.NewReplacer("{{name}}",name,"{{prefix}}",prefix,"{{package}}",strings.ReplaceAll(name,"-","_"),"{{module_dir}}",dir,"{{go_module}}",goModule)
	for path,source:=range templates {
		out:=strings.TrimSuffix(path,".tmpl")
		if !safeRelative(out) { return nil,project.Module{},fmt.Errorf("invalid module template path %q",path) }
		out=filepath.Join(dir,out)
		if _,exists:=files[out];exists { return nil,project.Module{},fmt.Errorf("multiple templates generate %s",out) }
		files[out]=render.Replace(source)
	}
	spec:=filepath.Join(dir,"SPEC.md")
	body,exists:=files[spec]
	if !exists { return nil,project.Module{},fmt.Errorf("module templates must generate SPEC.md") }
	doc,err:=contracts.ParseDocument(body)
	if err!=nil { return nil,project.Module{},fmt.Errorf("rendered SPEC.md: %v",err) }
	if doc.Frontmatter["name"]!=name || doc.Frontmatter["id_prefix"]!=prefix { return nil,project.Module{},fmt.Errorf("rendered SPEC.md must declare name %q and id_prefix %q",name,prefix) }
	m:=project.Module{Name:name,IDPrefix:prefix,Dir:dir,Document:doc}
	if raw,exists:=doc.Frontmatter["depends_on"];exists {
		var ok bool
		m.DependsOn,ok=stringList(raw)
		if !ok { return nil,project.Module{},fmt.Errorf("rendered SPEC.md depends_on must be a list of strings") }
	}
	for _,key:=range []string{"protected","evals","interface","append_only","preserve","disposable"} {
		if value,exists:=doc.Frontmatter[key];exists {
			if _,ok:=stringList(value);!ok { return nil,project.Module{},fmt.Errorf("rendered SPEC.md %s must be a list of strings",key) }
		}
	}
	if _,err:=specItems(m);err!=nil { return nil,project.Module{},fmt.Errorf("rendered SPEC.md: %v",err) }
	return files,m,nil
}

func newModule(c command,snap project.Snapshot) response {
	name,prefix:=c.arg(0),c.flags["--prefix"]
	if err:=project.ValidateGraph(snap.Modules);err!=nil { return reject(2,"C02",err) }
	patterns,_:=stringList(snap.Config["modules"])
	dir:=moduleDestination(patterns,name)
	if dir=="" { return reject(2,"C01",fmt.Errorf("modules must include a glob ending in a single wildcard under a literal parent")) }
	for _,m:=range snap.Modules {
		if m.Name==name || m.IDPrefix==prefix { return reject(2,"P01",fmt.Errorf("module name %q or ID prefix %q already exists",name,prefix)) }
	}
	if err:=vacant(snap.Root,dir);err!=nil { return reject(2,"P01",err) }
	templates,err:=readTree(filepath.Join(snap.Root,moduleTemplateDir))
	if err!=nil { return reject(2,"C03",err) }
	goModule,err:=goModuleName(snap.Root)
	if err!=nil { return reject(2,"P01",err) }
	files,m,err:=renderModule(templates,name,prefix,dir,goModule)
	if err!=nil { return reject(2,"C03",err) }
	if err=project.ValidateGraph(append(slices.Clone(snap.Modules),m));err!=nil { return reject(2,"C03",fmt.Errorf("rendered module graph: %v",err)) }
	if err=writeSetup(snap.Root,files,nil,dir);err!=nil { return reject(2,"P01",err) }
	return response{data:dir,plain:dir+"'"\\n\"}
}
EOF
python3 - <<'PY'
from pathlib import Path
p=Path('packages/app/internal/entry/entry.go')
s=p.read_text().replace('if c.name==\"new-module\" && "'!c.has("--prefix") { return bad("new-module requires --prefix") }'"','''if c.name==\"new-module\" {
		if "'!c.has("--prefix") { return bad("new-module requires --prefix") }
		if !moduleName.MatchString(c.arg(0)) { return bad("invalid module name %q",c.arg(0)) }
		if !modulePrefix.MatchString(c.flags["--prefix"]) { return bad("invalid ID prefix %q",c.flags["--prefix"]) }
	}'"''')
p.write_text(s)
p=Path('packages/app/internal/entry/lifecycle.go')
s=p.read_text().replace('\"fmt\"','\"fmt\"\\n\\t\"os\"')
s=s.replace('data:=struct {\\n\\t\\tModule string','''if report"'!="" {
		if _,e:=os.Stat(filepath.Join(snap.Root,report));e!=nil { report="" }
	}
	data:=struct {
		Module string'"''')
p.write_text(s)
PY
gofmt -w packages/app/internal
export PATH=\""'$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
mkdir -p .fireweed/runs/drill
sh -c '"'go run .fireweed/hooks.go eval \"packages/app\" \""'$1"'"' sh \""'$PWD/.fireweed/runs/drill/report.xml"' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed verify' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc "git status --short; cat .gitignore; rg -n 'func ErrorCode|func ParseItems' packages/contracts/internal; rg -n 'func requireApproval|func IsApprovalRequired' packages/runner/internal/next; sed -n '149,220p' packages/checks/internal/engine/junit.go; sed -n '105,175p' packages/checks/internal/engine/hooks.go; sed -n '515,524p' fireweed-design-v1.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 558ms:
?? packages/app/internal/
.fireweed/runs/
.fireweed/bin/
bin/
.fireweed/evidence/
packages/contracts/internal/primitives/coded_error.go:6:func ErrorCode(err error) (code, subject string, ok bool) {
packages/contracts/internal/primitives/primitives.go:100:func ParseItems(section string) ([]model.Item, error) {
packages/runner/internal/next/approval.go:16:func IsApprovalRequired(err error) bool {
packages/runner/internal/next/run.go:340:func requireApproval(root, path string) error {
func RequireCoverage(report model.Report, specs []contracts.Item) error {
	for _, spec := range specs {
		if exempt(spec.Tag, "review") || exempt(spec.Tag, "quarantine") {
			continue
		}
		covered := false
		for _, c := range report.Cases {
			if c.Result != "skip" && references(c, spec.ID) {
				covered = true
				break
			}
		}
		if !covered {
			return fmt.Errorf("C10: no coverage for %s", spec.ID)
		}
	}
	return nil
}

func RequireFailFirst(report model.Report, ids, already []string) error {
	for _, id := range ids {
		satisfied := false
		for _, done := range already {
			if done == id {
				satisfied = true
				break
			}
		}
		if satisfied {
			continue
		}
		for _, c := range report.Cases {
			if c.Result == "fail" && references(c, id) {
				satisfied = true
				break
			}
		}
		if !satisfied {
			return fmt.Errorf("P2-01: no fail-first case for %s", id)
		}
	}
	return nil
}

func RequirePass(report model.Report, specs []contracts.Item) error {
	for _, spec := range specs {
		if exempt(spec.Tag, "quarantine") {
			continue
		}
		for _, c := range report.Cases {
			if c.Result == "fail" && references(c, spec.ID) {
				return fmt.Errorf("E01: failed case %q for %s", c.Name, spec.ID)
			}
		}
	}
	return nil
}
func Evaluate(ctx context.Context, request model.HookRequest) (model.Report, error) {
	var module *project.Module
	for i := range request.Snapshot.Modules {
		if request.Snapshot.Modules[i].Name == request.Module {
			module = &request.Snapshot.Modules[i]
			break
		}
	}
	if module == nil {
		return model.Report{}, fmt.Errorf("U02: unknown module %q", request.Module)
	}
	command := mapValue(request.Snapshot.Config, "hooks")["eval"]
	if override, ok := mapValue(module.Document.Frontmatter, "hooks")["eval"]; ok {
		command = override
	}
	if command == nil {
		return model.Report{}, fmt.Errorf("C11: missing eval hook for %s", module.Name)
	}
	specs, err := moduleSpecs(*module)
	if err != nil {
		return model.Report{}, err
	}
	temp, err := os.MkdirTemp("", "fireweed-junit-")
	if err != nil {
		return model.Report{}, fmt.Errorf("C11: %w", err)
	}
	defer os.RemoveAll(temp)
	reportPath := filepath.Join(temp, "report.xml")
	values := map[string]string{"module": module.Name, "module_dir": module.Dir, "report": reportPath}
	process, err := runHook(ctx, request.Snapshot.Root, command, values, request.Bin, request.Timeout)
	report := model.Report{HookExit: process.ExitCode, Stdout: process.Stdout, Stderr: process.Stderr}
	if err != nil {
		return report, fmt.Errorf("C11: %w", err)
	}
	data, err := os.ReadFile(reportPath)
	if err != nil {
		return report, fmt.Errorf("C11: read JUnit report: %w", err)
	}
	parsed, err := ParseJUnit(data, specs)
	if err != nil {
		return report, err
	}
	report.Cases, report.Seconds = parsed.Cases, parsed.Seconds
	failed := false
	for _, c := range report.Cases {
		if c.Result == "fail" {
			failed = true
			break
		}
	}
	if (process.ExitCode != 0) != failed {
		return report, fmt.Errorf("C11: eval hook exit %d disagrees with JUnit failures", process.ExitCode)
	}
	return report, nil
}
**First run.** The expected flow starts with a developer telling their agent "set up fireweed for a TypeScript project":

1. The agent runs `fireweed init --recipe typescript-pnpm --agent claude`. This writes the config, `SYSTEM.md`, the `AGENTS.md` block, the recipe copy and an `app` composition-root module (`--no-app` skips it), then prints the recipe's instructions.
2. The agent completes the glue by following the recipe: workspace config, eval reporter, boundary checker.
3. The agent runs `fireweed doctor`, which plants canaries in throwaway modules and confirms every hook catches them. Failures name the hook that misbehaved.
4. With doctor passing, the developer fills in `SYSTEM.md` and authors the first change.

**Playbooks** for authoring changes, writing evals, auditing, implementing and regenerating are embedded in the binary, because they describe the protocol the binary verifies and must match it. Two rules in them exist to keep agents from escalating. The authoring playbook requires each change's Motivation to name a user-visible need or a clause of the design or `SYSTEM.md`; a probe showing that some failure is theoretically possible is not enough on its own. It also tells authors to narrow or remove specs, with `already_satisfied`, when a requirement turns out to cost more than it is worth. The audit playbook limits findings to spec clauses no test asserts, as described in Audit step. A project can override one in `.fireweed/playbooks/`; each override records the hash of the embedded playbook it was based on, and doctor warns when an upgrade changes that playbook underneath it.

**AGENTS.md.** `init` writes a short block between markers telling any agent that the project uses fireweed and pointing it at `fireweed playbook author-change` and `fireweed context`. `fireweed agents-md --write` regenerates only that block, and doctor warns when it is out of date.

 exited 1 in 570ms:
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.

exec
/bin/zsh -lc "sed -n '70,87p' packages/app/evals/app_test.go; sed -n '1,120p' packages/contracts/internal/primitives/coded_error.go; sed -n '92,172p' packages/contracts/internal/primitives/primitives.go; sed -n '1,97p' packages/workflow/internal/lifecycle/status_stub.go; sed -n '45,62p' packages/workflow/internal/lifecycle/stub.go; rg -n 'return .*C01|return .*C02|is not.*project' packages/workflow/internal/lifecycle/status_stub.go; ls -la .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 559ms:
func buildAppBinary(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("locate eval source")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", "..", ".."))
	bin := filepath.Join(t.TempDir(), "fireweed-new")
	build := exec.Command("go", "build", "-o", bin, ".")
	build.Dir = root
	build.Env = append(os.Environ(), "GOCACHE=/tmp/fireweed-rebuild-cache")
	if output, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build NEW binary: %v\n%s", err, output)
	}
	return bin
}

func Test_APP_002_InformationJSON(t *testing.T) {
package primitives

import "github.com/lurkshark/fireweed/packages/contracts/model"

// ErrorCode returns the first coded error encountered while unwrapping err.
func ErrorCode(err error) (code, subject string, ok bool) {
	for err != nil {
		switch coded := err.(type) {
		case model.CodedError:
			return coded.Code, coded.Subject, true
		case *model.CodedError:
			if coded != nil {
				return coded.Code, coded.Subject, true
			}
		}
		wrapped, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		err = wrapped.Unwrap()
	}
	return "", "", false
}
	if markdown != "" && !strings.HasSuffix(markdown, "\n") {
		markdown += "\n"
	}
	return markdown + "## " + name + "\n" + body
}

var itemHeading = regexp.MustCompile("^### ([A-Z0-9]+-[0-9]+|new:[a-z0-9-]+)(?: \\[([^\\[\\]]+)\\])?$")

func ParseItems(section string) ([]model.Item, error) {
	var items []model.Item
	seen := make(map[string]bool)
	bodyStart := 0
	for pos := 0; pos < len(section); {
		line, next := lineAt(section, pos)
		if strings.HasPrefix(line, "###") && !strings.HasPrefix(line, "####") {
			parts := itemHeading.FindStringSubmatch(line)
			if parts == nil {
				return nil, model.CodedError{Code: "C03", Subject: line, Cause: fmt.Errorf("malformed item heading: %q", line)}
			}
			if seen[parts[1]] {
				return nil, model.CodedError{Code: "C05", Subject: parts[1], Cause: fmt.Errorf("duplicate item ID: %s", parts[1])}
			}
			seen[parts[1]] = true
			if len(items) > 0 {
				items[len(items)-1].Body = section[bodyStart:pos]
			}
			items = append(items, model.Item{ID: parts[1], Tag: parts[2]})
			bodyStart = next
		}
		pos = next
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
	content := strings.ReplaceAll(RenderItem(item), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], " \t")
	}
	canonical := strings.TrimRight(strings.Join(lines, "\n"), "\n") + "\n"
	sum := sha256.Sum256([]byte(canonical))
	return hex.EncodeToString(sum[:4])
}

var specID = regexp.MustCompile("^([A-Za-z][A-Za-z0-9]*)[-_]?([0-9]+)$")

func MatchesID(id, testcaseName, classname string) bool {
	parts := specID.FindStringSubmatch(id)
	if parts == nil {
		return false
	}
	pattern := regexp.MustCompile("(?i)" + regexp.QuoteMeta(parts[1]) + "[-_]?" + regexp.QuoteMeta(parts[2]))
	for _, text := range []string{testcaseName, classname} {
		for _, match := range pattern.FindAllStringIndex(text, -1) {
			if match[0] > 0 && asciiLetterOrDigit(text[match[0]-1]) {
				continue
			}
			if match[1] < len(text) && text[match[1]] >= '0' && text[match[1]] <= '9' {
				continue
			}
			return true
		}
	}
	return false
}

func asciiLetterOrDigit(ch byte) bool {
	return ch >= 'A' && ch <= 'Z' || ch >= 'a' && ch <= 'z' || ch >= '0' && ch <= '9'
}
package lifecycle

import (
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type appliedChange struct {
	name, module string
	date         time.Time
}

func Status(root string) (model.StatusReport, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return model.StatusReport{}, err
	}
	root = repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return model.StatusReport{}, err
	}
	r := model.StatusReport{Modules: snapshot.Modules, DriftSinceDrill: map[string]int{}, PhaseCommits: map[string]int{}}
	for _, m := range snapshot.Modules {
		r.DriftSinceDrill[m.Name] = 0
		section, _ := contracts.Section(m.Document.Body, "Specs")
		items, err := contracts.ParseItems(section)
		if err != nil {
			return model.StatusReport{}, fmt.Errorf("%s/SPEC.md: %w", m.Dir, err)
		}
		for _, item := range items {
			if reason, ok := strings.CutPrefix(item.Tag, "quarantine:"); ok {
				r.Quarantines = append(r.Quarantines, model.Quarantine{Module: m.Name, ID: item.ID, Reason: strings.TrimSpace(reason)})
			}
		}
		if reason := str(m.Document.Frontmatter["boundaries_opt_out"]); reason != "" {
			r.BoundaryOptOuts = append(r.BoundaryOptOuts, model.BoundaryOptOut{Module: m.Name, Reason: reason})
		}
	}
	queue, err := project.Queue(root)
	if err != nil {
		return model.StatusReport{}, err
	}
	var blocked []project.Change
	for _, candidate := range queue {
		e := model.StatusEntry{Slug: candidate.Slug, Path: candidate.Path, Module: candidate.Module,
			Phase: str(candidate.Document.Frontmatter["phase"]), Blocked: str(candidate.Document.Frontmatter["blocked"]),
			AuditRounds: num(candidate.Document.Frontmatter["audit_rounds"], 0)}
		if e.Phase == "" {
			e.Phase = "queued"
		}
		if e.Blocked != "" {
			blocked = append(blocked, candidate)
		} else if len(blocked) > 0 {
			projected, err := project.Project(root, candidate.Slug)
			if err != nil {
				return model.StatusReport{}, err
			}
			graph := map[string][]string{}
			for _, m := range projected.Modules {
				graph[m.Name] = m.DependsOn
			}
			section, _ := contracts.Section(candidate.Document.Body, "Interface")
			for _, earlier := range blocked {
				if rule := blockingRule(graph, candidate.Module, earlier.Module, strings.TrimSpace(section) != ""); rule != "" {
					e.WaitingOn = strings.TrimSuffix(path.Base(earlier.Path), ".md")
					e.WaitingRule = rule
					break
				}
			}
		}
		r.Queue = append(r.Queue, e)
	}
	archives, err := appliedAtHead(root, repo.Head)
	if err != nil {
		return model.StatusReport{}, err
	}
	base, err := gitstore.BaseRef(root, str(snapshot.Config["remote"]), str(snapshot.Config["base"]))
	if err != nil {
		return model.StatusReport{}, err
	}
	basePaths, err := commitChangePaths(root, base)
	if err != nil {
		return model.StatusReport{}, err
	}
	known := map[string]bool{}
	for _, p := range basePaths {
		known[path.Base(p)] = true
		return err
	}
	return os.WriteFile(p, data, 0644)
}
func hard(ds []project.Diagnostic) error {
	for _, d := range ds {
		if !d.Warning {
			if d.Code == "C08" && strings.HasPrefix(d.Message, "stale base for ") {
				return fmt.Errorf("P05: C08: %s", d.Message)
			}
			if strings.HasPrefix(d.Message, d.Code+": ") {
				return errors.New(d.Message)
			}
			return fmt.Errorf("%s: %s", d.Code, d.Message)
		}
	}
	return nil
}
total 0
drwxr-xr-x@ 2 joel  staff  64 Nov 10 12:22 .
drwxr-xr-x@ 3 joel  staff  96 Nov 10 12:22 ..

exec
/bin/zsh -lc "sed -n '60,235p' packages/runner/internal/drill/drill.go; rg -n 'slugPattern|func num' packages/{project,workflow}/internal; sed -n '1,100p' packages/project/internal/queue/stale.go; rg -n '"'^### APP-'"' packages/app/SPEC.md | wc -l; git status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 595ms:
	result.Agent = agentName(snapshot.Config)
	result.Model = modelName(snapshot.Config)
	parent := filepath.Join(repo.CommonDir, "fireweed", "worktrees")
	if err = os.MkdirAll(parent, 0755); err != nil {
		return result, err
	}
	scratch, err := os.MkdirTemp(parent, "drill-")
	if err != nil {
		return result, err
	}
	if err = os.Remove(scratch); err != nil {
		return result, err
	}
	worktreeAdded := false
	input := ""
	defer func() {
		var cleanup error
		if input != "" {
			cleanup = errors.Join(cleanup, os.RemoveAll(input))
		}
		if worktreeAdded {
			parentInfo, statErr := os.Stat(parent)
			if statErr != nil {
				cleanup = errors.Join(cleanup, statErr)
			} else if parentInfo.Mode().Perm()&0200 == 0 {
				cleanup = errors.Join(cleanup, fmt.Errorf("scratch parent %s is not writable", parent))
			} else {
				_, removeErr := git(repo.Root, "worktree", "remove", "--force", scratch)
				cleanup = errors.Join(cleanup, removeErr)
			}
		} else {
			cleanup = errors.Join(cleanup, os.RemoveAll(scratch))
		}
		if cleanup != nil {
			err = errors.Join(err, fmt.Errorf("drill cleanup: %w", cleanup))
			result.Result = "fail"
			if len(result.Steps) > 0 {
				err = errors.Join(err, markReportFailure(repo.Root, result.ReportPath, cleanup))
			}
		}
	}()
	if _, err = git(repo.Root, "worktree", "add", "--detach", scratch, repo.Head); err != nil {
		return result, err
	}
	worktreeAdded = true
	if err = setup(ctx, scratch, snapshot.Config, request); err != nil {
		return result, err
	}

	// Delete only the target's disposable files. The verifier retains all other
	// project files, including dependent implementation and hook configuration.
	if err = removeDisposable(scratch, snapshot, *target); err != nil {
		return result, err
	}
	input, err = makeInput(scratch, repo.Head, snapshot, *target)
	if err != nil {
		return result, err
	}
	if err = setup(ctx, input, snapshot.Config, request); err != nil {
		return result, err
	}
	original, err := inventory(input)
	if err != nil {
		return result, err
	}
	eval, err := evalCommand(input, snapshot.Config, *target)
	if err != nil {
		return result, err
	}
	prompt, err := reconstructionPrompt(*target, original, eval)
	if err != nil {
		return result, err
	}
	attempts := configuredAttempts(snapshot.Config)
	wait, budget, err := usagePolicy(snapshot.Config)
	if err != nil {
		return result, err
	}
	var waited time.Duration
	feedback := ""
	accepted := false
	var failure error
	logRun, err := invoke.NewRunID()
	if err != nil {
		return result, err
	}
	for attempt := 0; attempt < attempts; {
		if err = ctx.Err(); err != nil {
			failure = err
			break
		}
		agentSnapshot := snapshot
		agentSnapshot.Root = input
		step, invokeErr := invoke.Agent(ctx, model.AgentRequest{
			Snapshot: agentSnapshot, Role: "worker", Prompt: prompt + feedback,
			Bin: request.Bin, Timeout: request.Timeout,
			LogRoot: repo.Root, Run: logRun, Phase: "drill", Attempt: attempt + 1,
		})
		result.Steps = append(result.Steps, step)
		if validationErr := validateInput(input, original); validationErr != nil {
			failure = fmt.Errorf("invalid reconstruction edits: %w", validationErr)
			break
		}
		if ctx.Err() != nil {
			failure = ctx.Err()
			break
		}
		if errors.Is(invokeErr, context.DeadlineExceeded) || errors.Is(invokeErr, context.Canceled) {
			failure = invokeErr
			break
		}
		if invoke.IsLogError(invokeErr) {
			failure = invokeErr
			break
		}
		if step.UsageLimit {
			if waited > budget || wait > budget-waited {
				failure = &next.UsageStopError{Completed: waited, Next: wait, Budget: budget}
				result.Result = "stopped"
				break
			}
			timer := time.NewTimer(wait)
			select {
			case <-ctx.Done():
				failure = ctx.Err()
			case <-timer.C:
				waited += wait
			}
			timer.Stop()
			if failure != nil {
				break
			}
			continue
		}
		attempt++
		if invokeErr != nil {
			failure = fmt.Errorf("reconstruction agent: %w", invokeErr)
			feedback = "\n\n# Previous reconstruction failure\n" + failure.Error() + "\n" + step.Stdout + "\n" + step.Stderr
			continue
		}
		if err = importDisposable(input, scratch, snapshot, *target, original); err != nil {
			failure = err
			break
		}
		accepted = true
		verification := verifyAll(ctx, snapshot, scratch, *target, request, &result)
		if verification == nil {
			failure = nil
			break
		}
		failure = verification
		feedback = "\n\n# Previous verification failed\n" + verification.Error() + failureDetails(result)
		if ctx.Err() != nil {
			failure = ctx.Err()
			break
		}
	}
	if failure == nil && !accepted {
		failure = fmt.Errorf("reconstruction attempts exhausted")
	}
	if accepted {
		result.Ref = "refs/fireweed/drills/" + request.Module + "/" + result.Date
		if retainErr := retain(scratch, repo.Root, result.Ref, *target); retainErr != nil {
			failure = errors.Join(failure, retainErr)
			result.Ref = ""
		}
	}
	if failure == nil {
		result.Result = "pass"
	}
	if len(result.Steps) > 0 {
		if reportErr := writeReport(repo.Root, input, result, failure); reportErr != nil {
			failure = errors.Join(failure, reportErr)
			result.Result = "fail"
		}
	}
packages/workflow/internal/lifecycle/stub.go:67:func num(v any, fallback int) int {
packages/project/internal/fixture/fixture.go:269:var slugPattern = regexp.MustCompile(`^[a-z0-9-]+$`)
packages/project/internal/fixture/fixture.go:273:	if !slugPattern.MatchString(slug) {
package queue

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func preceding(root, target string) (model.Change, *projection, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Change{}, nil, err
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		if errorCode(err, "") == "C13" {
			return model.Change{}, nil, contracts.CodedError{Code: "C13", Subject: target, Cause: err}
		}
		return model.Change{}, nil, err
	}
	index, err := resolve(changes, target)
	if err != nil {
		return model.Change{}, nil, err
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return model.Change{}, nil, err
	}
	for _, change := range changes[:index] {
		if err := p.apply(change); err != nil {
			return model.Change{}, nil, err
		}
	}
	return changes[index], p, nil
}

func changeBases(change model.Change) (map[string]any, error) {
	raw, exists := change.Document.Frontmatter["bases"]
	if !exists {
		return map[string]any{}, nil
	}
	bases, ok := raw.(map[string]any)
	if !ok {
		return nil, contracts.CodedError{Code: "C08", Subject: change.Path, Cause: fmt.Errorf("%s bases must be a mapping", change.Path)}
	}
	return bases, nil
}

func edited(change model.Change) []contracts.Item {
	items := make([]contracts.Item, 0, len(change.Modified)+len(change.Removed))
	items = append(items, change.Modified...)
	return append(items, change.Removed...)
}

// StaleSpecs reports recorded bases that differ from the preceding projection.
func StaleSpecs(root, target string) ([]model.StaleSpec, error) {
	change, p, err := preceding(root, target)
	if err != nil {
		return nil, err
	}
	bases, err := changeBases(change)
	if err != nil {
		return nil, err
	}
	stale := make([]model.StaleSpec, 0)
	for _, edit := range edited(change) {
		raw, exists := bases[edit.ID]
		if !exists {
			continue
		}
		recorded := fmt.Sprint(raw)
		current, ok := p.find(change.Module, edit.ID)
		if !ok {
			return nil, contracts.CodedError{Code: "C08", Subject: change.Path, Cause: fmt.Errorf("%s: ID %s absent from preceding projection", change.Path, edit.ID)}
		}
		if recorded != contracts.SpecHash(current) {
			stale = append(stale, model.StaleSpec{ID: edit.ID, RecordedHash: recorded, Current: contracts.RenderItem(current)})
		}
	}
	if len(stale) == 0 {
		return stale, nil
	}
	if before, ok := stampedProjection(root, change, bases); ok {
		for i := range stale {
			item, found := before.find(change.Module, stale[i].ID)
			if found && contracts.SpecHash(item) == stale[i].RecordedHash {
				stale[i].Before = contracts.RenderItem(item)
				stale[i].Recovered = true
     155
?? packages/app/internal/

exec
/bin/zsh -lc "sed -n '250,290p' packages/app/evals/run_for_cli_test.go; sed -n '65,82p' packages/workflow/internal/lifecycle/stub.go; sed -n '88,132p' packages/app/evals/ci_cli_test.go; rg -n 'func Write|"'^func|return'"' packages/project/internal/fixture/load_at.go | head -50; rg -n 'func Boundaries|func runHook' packages/checks/internal/engine/hooks.go; sed -n '150,197p' packages/app/evals/pr_body_cli_test.go; rg -n 'report.Root|Config' packages/workflow/internal/lifecycle/pr_body_stub.go | head -25" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 530ms:

func Test_APP_089_RunForInvalidOptionsBeforeMutation(t *testing.T) {
	bin := buildAppBinary(t)
	f := newRunCLISetup(t, bin, 1, false)
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	queued := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md")))
	statusBefore := cliGit(t, f.root, "status", "--porcelain", "--untracked-files=all")
	for _, tc := range []struct {
		name string
		args []string
	}{
		{"missing", []string{"run", "--for"}},
		{"malformed", []string{"run", "--for", "nonsense"}},
		{"zero", []string{"run", "--for", "0s"}},
		{"negative", []string{"run", "--for", "-1s"}},
		{"repeated", []string{"run", "--for", "1s", "--for", "2s"}},
		{"next", []string{"next", "--for", "1s"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			args := append(append([]string{}, tc.args...), "--json")
			status, out, stderr := runNextCLI(t, bin, f.root, args...)
			if status != 2 || stderr != "" {
				t.Errorf("invalid For exit=%d stdout=%q stderr=%q", status, out, stderr)
			}
			message := assertFailedJSON(t, out, "U01")
			if !regexp.MustCompile(`(?i)\brun\s+[^\r\n]{0,120}--for(?:\s+|=)(?:[1-9][0-9]*(?:ns|us|µs|ms|s|m|h)|<?(?:duration|d)>?)(?:\W|$)`).MatchString(message) {
				t.Errorf("diagnostic lacks valid run --for usage: %q", message)
			}
			if got := cliGit(t, f.root, "rev-parse", "HEAD"); got != head {
				t.Errorf("invalid For committed: %s", got)
			}
			if got := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md"))); got != queued {
				t.Error("invalid For changed queued file")
			}
			if got := cliGit(t, f.root, "status", "--porcelain", "--untracked-files=all"); got != statusBefore {
				t.Errorf("invalid For changed files: %q", got)
			}
			if claims := cliGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); claims != "" {
				t.Errorf("invalid For claimed work: %q", claims)
			}
			if calls := runCLICalls(t, f); len(calls) != 0 {
}
func str(v any) string { s, _ := v.(string); return s }
func num(v any, fallback int) int {
	switch x := v.(type) {
	case int:
		return x
	case float64:
		return int(x)
	case string:
		n, e := strconv.Atoi(x)
		if e == nil {
			return n
		}
	}
	return fallback
}
func mapval(v any) map[string]any { m, _ := v.(map[string]any); return m }
func list(v any) []string {

func ciCLIDiagnostics(t *testing.T, output string) []struct {
	Code, Message string
	Warning       bool
} {
	t.Helper()
	envelope := decodeEnvelope(t, output)
	var diagnostics []struct {
		Code, Message string
		Warning       bool
	}
	if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
		t.Fatalf("CI diagnostics=%s: %v", envelope["diagnostics"], err)
	}
	return diagnostics
}

func ciCLIFailure(t *testing.T, nested, code string, status int, args ...string) string {
	t.Helper()
	got, output, stderr := runSavedAttemptApp(t, nested, args...)
	if got != status || stderr != "" {
		t.Errorf("ci %v: status=%d stdout=%q stderr=%q; want %d and empty stderr", args, got, output, stderr, status)
	}
	assertFailedJSON(t, output)
	found := false
	for _, d := range ciCLIDiagnostics(t, output) {
		if d.Code == code {
			found = true
			if !corrective(d.Message) {
				t.Errorf("CI %s diagnostic lacks corrective action: %q", code, d.Message)
			}
		}
	}
	if !found {
		t.Errorf("CI failure lacks %s diagnostic: %s", code, output)
	}
	return output
}

func Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules(t *testing.T) {
	root, nested, base := ciCLIFixture(t)
	status, plain, stderr := runSavedAttemptApp(t, nested, "ci", "--no-doctor")
	if status != 0 || stderr != "" || strings.TrimSpace(plain) == "" {
		t.Errorf("configured-base CI: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
17:func LoadAt(root, revision string) (model.Snapshot, error) {
20:		return model.Snapshot{}, fmt.Errorf("find repository root: %w", err)
24:		return model.Snapshot{}, fmt.Errorf("canonical repository root: %w", err)
28:		return model.Snapshot{}, fmt.Errorf("resolve revision %q: %w", revision, err)
34:		return model.Snapshot{}, fmt.Errorf("read fireweed.yaml at %s: %w", commit, err)
38:		return model.Snapshot{}, err
43:		return model.Snapshot{}, fmt.Errorf("list module directories at %s: %w", commit, err)
64:			return model.Snapshot{}, fmt.Errorf("read %s at %s: %w", path, commit, err)
68:			return model.Snapshot{}, err
72:	return snapshot, nil
75:func readGit(root string, args ...string) ([]byte, error) {
82:		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, result.Stderr)
84:	return []byte(result.Stdout), nil
74:func runHook(ctx context.Context, root string, command any, values map[string]string, bin string, timeout time.Duration) (gitstore.Result, error) {
	cliGit(t, badConfig, "init", "-q", "-b", "main")
	cliCommit(t, badConfig, "valid base")
	badBase := cliGit(t, badConfig, "rev-parse", "HEAD")
	appPut(t, badConfig, "fireweed.yaml", "[invalid endpoint config\n")
	cliCommit(t, badConfig, "bad endpoint config")
	badEndpoint := cliGit(t, badConfig, "rev-parse", "HEAD")
	badRange := badBase + ".." + badEndpoint
	if _, err := workflow.PRBody(badConfig, badRange); err == nil || !strings.Contains(err.Error(), "fireweed.yaml") {
		t.Fatalf("public PRBody bad endpoint configuration error=%v; want configuration cause", err)
	}
	unavailableRange := strings.SplitN(revisionRange, "..", 2)[0] + "..not-a-commit"
	if _, err := workflow.PRBody(root, unavailableRange); err == nil || !strings.Contains(err.Error(), "P01") {
		t.Fatalf("public PRBody unavailable endpoint error=%v; want P01", err)
	}
	for _, tc := range []struct {
		name, dir, code, cause string
		args                   []string
		status                 int
	}{
		{"missing-range", nested, "U01", "pr-body", []string{"pr-body"}, 2},
		{"extra-range", nested, "U01", "pr-body", []string{"pr-body", revisionRange, "HEAD"}, 2},
		{"invalid-option", nested, "U01", "--bogus", []string{"pr-body", revisionRange, "--bogus"}, 2},
		{"absent-project", invalidRoot, "C01", "", []string{"pr-body", revisionRange}, 2},
		{"bad-endpoint-config", badConfig, "C01", "", []string{"pr-body", badRange}, 2},
		{"unavailable-range", nested, "P01", "not-a-commit", []string{"pr-body", unavailableRange}, 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for _, jsonMode := range []bool{false, true} {
				args := append([]string(nil), tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := prBodyAppRun(t, tc.dir, args...)
				message := commandFailure(t, status, tc.status, stdout, stderr, jsonMode)
				if jsonMode {
					message = assertFailedJSON(t, stdout, tc.code)
				} else if !strings.HasPrefix(strings.TrimSpace(stdout+stderr), tc.code+":") {
					t.Errorf("plain diagnostic=%q; want %s", stdout+stderr, tc.code)
				}
				if !corrective(message) || (tc.cause != "" && !strings.Contains(message, tc.cause)) {
					t.Errorf("diagnostic %q lacks cause %q or corrective action", message, tc.cause)
				}
				if jsonMode && stderr != "" {
					t.Errorf("JSON stderr=%q; want empty", stderr)
				}
			}
		})
	}
43:		return model.PRBodyReport{}, &endpointConfigError{cause: fmt.Errorf("load endpoint project at %s: %w", resolved[1], err)}
233:	for _, pattern := range append([]string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}, list(snapshot.Config["protected"])...) {
322:				for _, pattern := range list(snapshot.Config["worker_writable"]) {

exec
/bin/zsh -lc "sed -n '190,246p' packages/app/evals/saved_attempt_cli_test.go; sed -n '174,213p' packages/app/evals/check_fix_cli_test.go; rg -n 'func TestMain|func runSavedAttemptApp' packages/app/evals/{saved_attempt_cli_test,ci_cli_test}.go; sed -n '1,39p' packages/app/evals/ci_cli_test.go; cat packages/runner/internal/next/approval.go; rg -n 'func.*PRBody|C01|P01' packages/workflow/internal/lifecycle/pr_body_stub.go | head -20; ls .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 551ms:
func Test_APP_088_FromAttemptUsageRejectsBeforeMutation(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	head := cliGit(t, root, "rev-parse", "HEAD")
	change, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name  string
		args  []string
		value bool
	}{
		{"missing", []string{"start", "manual", "--from-attempt", "--json"}, true},
		{"noninteger", []string{"start", "manual", "--from-attempt", "abc", "--json"}, true},
		{"zero", []string{"start", "manual", "--from-attempt", "0", "--json"}, true},
		{"negative", []string{"start", "manual", "--from-attempt", "-1", "--json"}, true},
		{"repeated", []string{"start", "manual", "--from-attempt", "1", "--from-attempt", "2", "--json"}, true},
		{"context", []string{"context", "manual", "--from-attempt", "1", "--json"}, false},
		{"help", []string{"help", "--from-attempt", "1", "--json"}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			status, out, errOut := runSavedAttemptApp(t, root, tc.args...)
			if status != 2 || errOut != "" {
				t.Errorf("usage: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			message := assertFailedJSON(t, out, "U01")
			if !corrective(message) {
				t.Errorf("usage diagnostic is not corrective: %q", message)
			}
			if tc.value && (!strings.Contains(message, "start") || !strings.Contains(message, "--from-attempt") || !strings.Contains(message, "<change>")) {
				t.Errorf("value diagnostic omits valid start usage: %q", message)
			}
			assertSavedAttemptUnchanged(t, root, head, string(change))
		})
	}
}
	if got := checkFixCLI(t, bin, root, true, 0, ""); got != plain || !bytes.Equal(after, staleCLIText(t, root, path)) {
		t.Errorf("repeated fix changed output or bytes: %q", got)
	}
}

func Test_APP_148_APP_149_CheckFixRejectsBeforeMutation(t *testing.T) {
	bin := buildAppBinary(t)
	root := checkFixFixture(t)
	staleCLIChange(t, root, "changes/0001-first.md", "", "Needed.\n", "### ALP-001\nOne MUST change.\n", "")
	path := "changes/0001-first.md"
	before := staleCLIText(t, root, path)
	for _, args := range [][]string{{"check", "--fix", "extra"}, {"show", "alpha", "--fix"}, {"modules", "--fix"}, {"check", "--fix", "--bogus"}} {
		for _, jsonMode := range []bool{false, true} {
			call := append([]string(nil), args...)
			if jsonMode {
				call = append(call, "--json")
			}
			status, out, stderr := runAppBinary(t, bin, appNestedDir(t, root), call...)
			assertAppHistoryFailure(t, status, 2, out, stderr, jsonMode, "U01", "check --fix")
			if !bytes.Equal(before, staleCLIText(t, root, path)) {
				t.Error("invalid fix usage changed queued file")
			}
		}
	}
	staleCLIChange(t, root, path, "bases: [invalid]\n", "Needed.\n", "### ALP-001\nOne MUST change.\n", "")
	before = staleCLIText(t, root, path)
	for _, jsonMode := range []bool{false, true} {
		checkFixCLI(t, bin, root, jsonMode, 1, "C08")
		if !bytes.Equal(before, staleCLIText(t, root, path)) {
			t.Error("stamping rejection changed queued file")
		}
	}
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	checkFixCLI(t, bin, root, true, 2, "C01")
	if !bytes.Equal(before, staleCLIText(t, root, path)) {
		t.Error("configuration rejection changed queued file")
	}
	if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
		t.Fatal(err)
	}
packages/app/evals/saved_attempt_cli_test.go:17:func runSavedAttemptApp(t *testing.T, dir string, args ...string) (int, string, string) {
packages/app/evals/ci_cli_test.go:18:func TestMain(m *testing.M) {
package evals_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/app"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
)

// Scratch doctor hooks invoke FIREWEED_BIN modules --json. In this file the
// invoking executable is the Go test binary, dispatched through public app.Run.
func TestMain(m *testing.M) {
	if os.Getenv("FIREWEED_APP_EVAL_CLI") == "1" {
		os.Exit(app.Run(os.Args[1:], os.Stdout, os.Stderr))
	}
	os.Exit(m.Run())
}

func ciCLIFixture(t *testing.T) (root, nested, base string) {
	t.Helper()
	outer := t.TempDir()
	appPut(t, outer, "fireweed.yaml", appFixtureConfig)
	appPut(t, outer, "packages/ancestor/SPEC.md", "---\nname: ancestor\nid_prefix: ANC\ndepends_on: []\n---\n# Ancestor\n\n## Specs\n\n### ANC-001\nThe ancestor MUST pass.\n")
	appPut(t, outer, "packages/ancestor/api.go", "package ancestor\n")
	root = filepath.Join(outer, "inner")
	log := filepath.Join(root, ".fireweed", "hooks.log")
	config := strings.Replace(appFixtureConfig, "base: main", "base: review-base", 1) + `verify_dependents: none
protected: [locked/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
`
	appPut(t, root, "fireweed.yaml", config)
package next

import (
	"errors"
	"fmt"
)

type approvalRequiredError struct {
	path string
}

func (e *approvalRequiredError) Error() string {
	return fmt.Sprintf("change %s requires approval on the local or shared base", e.path)
}

func IsApprovalRequired(err error) bool {
	var approval *approvalRequiredError
	return errors.As(err, &approval)
}
28:func PRBody(root, revisionRange string) (model.PRBodyReport, error) {
31:		return model.PRBodyReport{}, fmt.Errorf("P01: invalid range %q", revisionRange)
37:			return model.PRBodyReport{}, fmt.Errorf("P01: resolve range %q: %w", revisionRange, err)
419:func prCosts(body *strings.Builder, root string, endpoint project.Snapshot, result *model.PRBodyReport) error {
474:func prArtifacts(body *strings.Builder, root, base, endpoint string, snapshot project.Snapshot, result *model.PRBodyReport) error {

exec
/bin/zsh -lc "rg -n 'name.*slug|id_prefix|prefix.*uppercase|new-module' fireweed-design-v1.md packages/project/SPEC.md | head -35; sed -n '82,133p' packages/runner/internal/next/run.go; sed -n '25,77p' packages/workflow/internal/lifecycle/pr_body_stub.go; sed -n '150,190p' packages/app/evals/show_history_cli_test.go; cat .fireweed/recipe.md | head -55; ls -l .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 574ms:
packages/project/SPEC.md:6:id_prefix: PRJ
packages/project/SPEC.md:27:Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets. An unknown target MUST return U02 naming the rejected selector and listing valid pending change slugs.
fireweed-design-v1.md:68:id_prefix: AUTH
fireweed-design-v1.md:78:Any zone field left out takes its value from `module_defaults` in `fireweed.yaml`, so a module using the recommended layout needs only `name`, `id_prefix` and `depends_on`. A Go module might instead declare `interface: [auth.go]`, `evals: ["*_eval_test.go"]` and `disposable: [internal/**]`.
fireweed-design-v1.md:99:- `id_prefix` is unique across modules.
fireweed-design-v1.md:178:**Ordering and numbering.** Files are named `NNNN-slug.md`. Numbers only order the queue and need not be unique; changes that share a number run in slug order. Two changes written in parallel without knowing about each other have no meaningful order anyway. Nothing refers to a change by its number.
fireweed-design-v1.md:191:| `set` | Author | New values for module frontmatter: `depends_on`, `interface`, `evals`, `disposable`, `preserve`, `append_only`, `boundaries_opt_out` and per-module hook overrides. `name` and `id_prefix` cannot be changed |
fireweed-design-v1.md:429:| Authoring | `new-module <name> --prefix <P>` | Create a module from the templates |
fireweed-design-v1.md:542:| Random suffix in change filenames to avoid number clashes | Rejected | Numbers only order the queue; ties sort by slug, and real collisions (IDs, specs) are handled elsewhere |
		switch phase {
		case "evals", "impl":
			prompt, err = workflow.Context(state.Snapshot.Root, state.Change.Path, true)
		case "audit":
			role = "audit"
			prompt, err = auditPrompt(state)
		default:
			return result, fmt.Errorf("unsupported next phase %q", phase)
		}
		if err != nil {
			return result, fmt.Errorf("prepare %s context: %w", phase, err)
		}
		step := workflow.Request{
			Root: state.Snapshot.Root, Change: state.Change.Path, Bin: request.Bin,
			Agent: configuredCommand(state.Snapshot.Config, role), Timeout: request.Timeout,
		}
		attempts := configuredAttempts(state.Snapshot.Config)
		feedback := ""
		for attempt := 0; attempt < attempts; {
			if stop := budget.TimeStop(ctx); stop != nil {
				return result, stop
			}
			preInvoke, inspectErr := gitstore.Inspect(state.Snapshot.Root)
			if inspectErr != nil {
				return result, fmt.Errorf("inspect held phase HEAD: %w", inspectErr)
			}
			if preInvoke.Head != state.PhaseBase {
				return result, fmt.Errorf("P05: held %s phase HEAD moved from %s to %s before agent invocation", phase, state.PhaseBase, preInvoke.Head)
			}
			invokeSnapshot := state.Snapshot
			var audit *auditWorkspace
			if role == "audit" {
				audit, err = newAuditWorkspace(state)
				if err != nil {
					return result, fmt.Errorf("prepare audit workspace: %w", err)
				}
				invokeSnapshot.Root = audit.root
			}
			agent, invokeErr := invoke.Agent(ctx, model.AgentRequest{
				Snapshot: invokeSnapshot, Role: role, Prompt: prompt + feedback,
				Bin: request.Bin, Timeout: request.Timeout,
				LogRoot: state.Snapshot.Root, Run: state.Run, Phase: phase, Attempt: attempt + 1,
			})
			result.Steps = append(result.Steps, model.StepResult{Phase: phase, Role: role, Agent: agent})
			if audit != nil {
				auditErr := audit.importAudit(state.Snapshot.Root, state.Change.Path)
				cleanupErr := os.RemoveAll(audit.root)
				if auditErr != nil || cleanupErr != nil {
					if auditErr != nil {
						auditErr = fmt.Errorf("validate audit submission: %w", auditErr)
					}
					return result, errors.Join(auditErr, cleanupErr, invokeErr)
	commits                  []prCommit
}

func PRBody(root, revisionRange string) (model.PRBodyReport, error) {
	parts := strings.Split(revisionRange, "..")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return model.PRBodyReport{}, fmt.Errorf("P01: invalid range %q", revisionRange)
	}
	var resolved [2]string
	for i, part := range parts {
		out, err := commitsGit(root, "rev-parse", "--verify", "--end-of-options", part+"^{commit}")
		if err != nil {
			return model.PRBodyReport{}, fmt.Errorf("P01: resolve range %q: %w", revisionRange, err)
		}
		resolved[i] = strings.TrimSpace(out)
	}
	endpoint, err := project.LoadAt(root, resolved[1])
	if err != nil {
		return model.PRBodyReport{}, &endpointConfigError{cause: fmt.Errorf("load endpoint project at %s: %w", resolved[1], err)}
	}
	root = endpoint.Root
	out, err := commitsGit(root, "rev-list", resolved[0]+".."+resolved[1])
	if err != nil {
		return model.PRBodyReport{}, err
	}
	inRange := map[string]bool{}
	for _, oid := range strings.Fields(out) {
		inRange[oid] = true
	}
	paths, err := commitChangePaths(root, resolved[1])
	if err != nil {
		return model.PRBodyReport{}, err
	}
	byName := map[string]*prChange{}
	for _, p := range paths {
		data, err := gitstore.ReadAt(root, resolved[1], p)
		if err != nil {
			return model.PRBodyReport{}, fmt.Errorf("read endpoint %s: %w", p, err)
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return model.PRBodyReport{}, fmt.Errorf("parse endpoint %s: %w", p, err)
		}
		name := path.Base(p)
		byName[name] = &prChange{path: p, name: name, module: str(doc.Frontmatter["module"]), kind: prKind(doc), doc: doc}
	}
	selected := map[string]bool{}
	out = ""
	if len(inRange) > 0 && len(byName) > 0 {
		// One ancestry traversal supplies retained phase context for all identities.
		out, err = commitsGit(root, "rev-list", "--topo-order", "--parents", resolved[1])
		if err != nil {
			return model.PRBodyReport{}, err
				for _, value := range append([]string{entry.Module, entry.Motivation, entry.Decisions, entry.RemovedDecisions, entry.Provenance}, entry.IDs...) {
					if value = strings.TrimSpace(value); value != "" && !strings.Contains(block, value) {
						t.Errorf("plain history %q omits metadata %q: %q", entry.Path, value, block)
					}
				}
				for _, label := range []string{"motivation", "decisions", "removed decisions", "provenance", "implementation date"} {
					if !strings.Contains(strings.ToLower(block), label) {
						t.Errorf("plain history %q omits %q label: %q", entry.Path, label, block)
					}
				}
				if entry.ImplementationDate.IsZero() {
					if !strings.Contains(strings.ToLower(block), "undated") {
						t.Errorf("plain history %q must identify undated entry: %q", entry.Path, block)
					}
				} else if !strings.Contains(block, entry.ImplementationDate.Format(time.RFC3339)) {
					t.Errorf("plain history %q omits implementation date %s: %q", entry.Path, entry.ImplementationDate, block)
				}
			}
			status, output, stderr := runAppBinary(t, bin, nested, append(args, "--json")...)
			if status != 0 || stderr != "" {
				t.Fatalf("JSON status=%d stdout=%q stderr=%q; want success", status, output, stderr)
			}
			assertAppHistoryJSON(t, output, entries)
		})
	}
}

func Test_APP_132_APP_133_ShowHistoryNearestRootAndEmpty(t *testing.T) {
	bin := buildAppBinary(t)
	outer := appHistoryFixture(t)
	inner := filepath.Join(outer, "other-project")
	appPut(t, inner, "fireweed.yaml", appFixtureConfig)
	appPut(t, inner, "packages/alpha/SPEC.md", authoringSpec)
	appHistoryArchive(t, inner, "changes/archive/0007-inner.md", "alpha", "Inner project only.\n\n",
		"### ALP-001\nInner behavior.\n", "", "", "", "- impl 2026-01-02T03:04:05Z agent=fixture\n")
	nested := appNestedDir(t, inner)
	status, plain, stderr := runAppBinary(t, bin, nested, "show", "alpha", "ALP-001", "--history")
	if status != 0 || stderr != "" || !strings.Contains(plain, "changes/archive/0007-inner.md") || strings.Contains(plain, "changes/archive/0002-revised.md") {
		t.Errorf("nearest root plain: status=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	status, output, stderr := runAppBinary(t, bin, nested, "show", "alpha", "ALP-001", "--history", "--json")
# Project-owned Go recipe

Each module exposes a thin api.go facade, keeps black-box Go tests in evals/,
and places disposable behavior under internal/. The protected Go hook runs
`go test -json ./<module>/...`, writes JUnit, and rejects unfinished tests and
package failures. Its boundary check reads the invoking FIREWEED_BIN's modules
JSON; module SPEC.md files are the only graph source. It enforces declared
cross-module dependencies and eval imports through public interfaces.

When naming Go eval tests, a letter immediately before a spec prefix prevents ID matching. Test_GRT_001_Name matches GRT-001; TestGRT001 does not match GRT-001.

Canaries use the controller's version-1 canaries.yaml contract, with Go module
and random module-name substitutions. Templates use {{name}}, {{package}} and
{{prefix}}. The setup hook copies the invoking controller into scratch trees
and exposes it as a regular, independent .fireweed/bin/fireweed copy. Launch
workers with the current tree's .fireweed/bin first on PATH. Main uses the
controller recorded in docs/rebuild/PROGRESS.md; mutable bin/fireweed is the
subject.
FIREWEED_BIN always names the invoking controller, including in scratch hooks.
Product fixtures explicitly select the product executable under test. Never fall back
automatically or switch controllers while a run or claim is active.

Go and Git are required. The controller and setup remain protected. Run doctor
after a coherent setup commit; no package downloads are part of setup.
total 0

exec
/bin/zsh -lc "sed -n '255,349p' packages/app/evals/status_cli_test.go; sed -n '82,90p' packages/app/evals/ci_cli_test.go; sed -n '117,141p' packages/app/evals/show_stale_cli_test.go; rg -n 'warning|C0' packages/app/evals/lifecycle_cli_test.go | tail -30; rg -n 'func IsTimeStop|TimeStopError.*Error' packages/runner/internal/{next,run}/*.go; git diff --check" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 529ms:
		pos := strings.Index(plain, entry.Path)
		if pos <= last {
			t.Errorf("queue path %q at %d after %d in %q", entry.Path, pos, last, plain)
		}
		last = pos
		for _, field := range []string{entry.Slug, entry.Path, entry.Module, entry.Phase, entry.Blocked, entry.WaitingOn, entry.WaitingRule, strconv.Itoa(entry.AuditRounds)} {
			if field != "" && !strings.Contains(plain, field) {
				t.Errorf("plain status missing queue field %q: %q", field, plain)
			}
		}
	}
	for _, module := range report.Modules {
		for _, field := range []string{module.Name, module.IDPrefix, module.Dir} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain status missing module field %q", field)
			}
		}
		for _, dep := range module.DependsOn {
			if !strings.Contains(plain, dep) {
				t.Errorf("plain status missing dependency %q", dep)
			}
		}
		for _, zones := range [][]string{module.Zones.Interface, module.Zones.Evals, module.Zones.Disposable, module.Zones.Preserve, module.Zones.Protected, module.Zones.AppendOnly} {
			for _, zone := range zones {
				if !strings.Contains(plain, zone) {
					t.Errorf("plain status missing resolved zone %q", zone)
				}
			}
		}
	}
	for _, q := range report.Quarantines {
		for _, field := range []string{q.Module, q.ID, q.Reason} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain status missing quarantine %q", field)
			}
		}
	}
	for _, o := range report.BoundaryOptOuts {
		if !strings.Contains(plain, o.Module) || !strings.Contains(plain, o.Reason) {
			t.Errorf("plain status missing opt-out %+v", o)
		}
	}
	if !strings.Contains(plain, strconv.Itoa(report.Unapproved)) {
		t.Errorf("plain status missing unapproved=%d", report.Unapproved)
	}
	for id, count := range report.DriftSinceDrill {
		if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
			t.Errorf("plain status missing drift %s=%d", id, count)
		}
	}
	for id, count := range report.PhaseCommits {
		if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
			t.Errorf("plain status missing phase commits %s=%d", id, count)
		}
	}
	for _, cost := range report.Costs {
		for _, field := range []string{cost.Module, strconv.FormatFloat(cost.Seconds, 'f', -1, 64), strconv.Itoa(cost.EvalLines), strconv.Itoa(cost.DisposableLines)} {
			if !strings.Contains(plain, field) {
				t.Errorf("plain status missing cost field %q", field)
			}
		}
		for id, count := range cost.CasesPerSpec {
			if !strings.Contains(plain, id) || !strings.Contains(plain, strconv.Itoa(count)) {
				t.Errorf("plain status missing cases %s=%d", id, count)
			}
		}
	}
	// Cost rows distinguish measurements from module, queue and quarantine metadata.
	costRows := map[string]int{}
	for _, line := range strings.Split(plain, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		label := strings.TrimSuffix(strings.ToLower(fields[0]), ":")
		if label == "cost" || label == "metric" {
			costRows[strings.TrimSuffix(fields[1], ":")]++
		}
	}
	if len(costRows) != len(report.Costs) {
		t.Errorf("plain status cost rows=%v; want one per recorded cost, with none for unmeasured gamma", costRows)
	}
	for _, cost := range report.Costs {
		if costRows[cost.Module] != 1 {
			t.Errorf("plain status cost rows for %s=%d; want 1", cost.Module, costRows[cost.Module])
		}
	}
	if costRows["gamma"] != 0 {
		t.Errorf("plain status fabricated gamma cost row: %v", costRows)
	}
	for _, d := range report.Diagnostics {
		if !strings.Contains(plain+stderr, d.Code) || !strings.Contains(plain+stderr, d.Message) {
			t.Errorf("plain status missing diagnostic %+v in stdout=%q stderr=%q", d, plain, stderr)
		}
	}
	var confirmation string
	if err := json.Unmarshal(envelope["data"], &confirmation); err != nil || strings.TrimSpace(confirmation) == "" {
		t.Fatalf("CI confirmation=%s: %v", envelope["data"], err)
	}
	return confirmation
}

func ciCLIDiagnostics(t *testing.T, output string) []struct {
	Code, Message string
			t.Errorf("show changed %s", file)
		}
	}
	if cliGit(t, root, "rev-parse", "HEAD") != head || cliGit(t, root, "for-each-ref", "--format=%(refname) %(objectname)") != refs {
		t.Error("show changed HEAD or refs")
	}
}

func Test_APP_144_StaleUnavailableHistoryAndEmptyView(t *testing.T) {
	bin := buildAppBinary(t)
	root := staleCLIProject(t)
	staleCLIChange(t, root, "changes/0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
	plain := staleCLISuccess(t, bin, appNestedDir(t, root), "target")
	for _, marker := range []string{"ALP-001", "deadbeef", "One MUST work.", "unavailable"} {
		if !strings.Contains(strings.ToLower(plain), strings.ToLower(marker)) {
			t.Errorf("unavailable-history output lacks %q: %q", marker, plain)
		}
	}
	if strings.Contains(plain, "-One MUST") || strings.Contains(plain, "+One MUST") || strings.Contains(plain, "Proposed MUST differ") {
		t.Errorf("unavailable history invented a recovered diff: %q", plain)
	}
	// A separate valid missing-base change has no stale entries.
	staleCLIChange(t, root, "changes/0001-target.md", "", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
	staleCLISuccess(t, bin, appNestedDir(t, root), "target")
}
239:			t.Fatalf("start with approval warning (JSON=%v): %d %q %q", jsonMode, status, output, stderr)
248:				Warning bool   `json:"warning"`
251:				t.Errorf("approval warning missing: %s, %v", output, err)
254:			t.Errorf("plain start omitted context or P04 warning: %q %q", output, stderr)
packages/runner/internal/run/time_stop.go:18:func IsTimeStop(err error) bool {
packages/runner/internal/next/usage.go:60:func (e *TimeStopError) Error() string {

exec
/bin/zsh -lc "ps -Ao pid,ppid,etime,comm | rg 'go"'$|evals.test|hooks$|fireweed-new$'"' | tail -30; ls -la .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 518ms:
zsh:1: operation not permitted: ps
total 0
drwxr-xr-x@ 2 joel  staff  64 Nov 10 12:22 .
drwxr-xr-x@ 3 joel  staff  96 Nov 10 12:22 ..

exec
/bin/zsh -lc "sed -n '1,110p' packages/app/SPEC.md; sed -n '76,124p' packages/app/evals/commits_cli_test.go; sed -n '174,201p' packages/app/evals/check_commits_cli_test.go; sed -n '218,245p' packages/app/evals/run_for_cli_test.go; sed -n '76,106p' packages/app/evals/check_fix_cli_test.go; rg -n 'unknown|ambiguous' packages/workflow/internal/lifecycle/{stub.go,commits_stub.go,abandon_stub.go} | head -30" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 514ms:
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
    - runner
id_prefix: APP
name: app
---

# app

## Specs

### APP-001
The CLI MUST run help, version, explain <code> and playbook <name> successfully
outside a project, returning the corresponding assets text on stdout with exit
0 and empty stderr. These commands MUST also work in a directory containing an
invalid fireweed.yaml.

### APP-002
Each information command MUST accept --json and emit one JSON object with schema
1, string data equal to its plain text result, and an empty diagnostics array.
It MUST return exit 0 and leave stderr empty.

### APP-003
Unknown commands or options and missing or extra arguments MUST return exit 2 with a nonempty corrective U01 diagnostic.

### APP-004
modules MUST report the validated graph, using the held approved projected snapshot while this worktree owns a change.

### APP-005
new-change <module> <slug> MUST use the discovered project root to create the project API's queued skeleton and print its repository-relative path with exit 0.

### APP-006
stamp <change> MUST discover the root, invoke project.Stamp and print confirmation with exit 0, accepting the existing change aliases and leaving Added handles unassigned.

### APP-007
check MUST invoke project.Check and exit 0 when no error diagnostics remain, or exit 1 with all diagnostics when static checks fail.

### APP-008
Without --history or --stale, show <module> [ID] MUST return the current module document or selected item, using the whole-queue project projection with --projected.

### APP-009
Without --from-attempt, start <change> MUST discover the root, invoke
workflow.Start and print first-step workflow.Context with exit 0, assigning IDs,
run and phase without committing.

### APP-010
verify [change] MUST invoke workflow with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.

### APP-011
context <change> [--full] MUST return the held workflow.Context text, using compact or full behavior as requested, with exit 0 and empty diagnostics.

### APP-012
coverage <module> MUST evaluate current specs, or the selected projection with --with-change <change>, using approved metadata while held and setting FIREWEED_BIN to the actual executable.

### APP-013
Unknown explain codes or playbook names MUST return exit 2 with a nonempty corrective U02 diagnostic.

### APP-014
gc and migrate listed by help MUST dispatch supported maintenance operations instead of reporting that they are unimplemented or under construction.

### APP-015
Failed --json requests MUST use the schema-1 envelope with null data and a diagnostic containing code and message.

### APP-016
modules MUST discover the nearest ancestor fireweed.yaml.

### APP-017
modules --json MUST return schema 1 with array data containing name, id_prefix, dir, depends_on and resolved zone arrays, plus empty diagnostics on success.

### APP-018
modules MUST return exit 0 on success or exit 2 for missing or invalid project inputs, with a corrective failure diagnostic: C01 for configuration, C02 for graph, C03 for document syntax or C12 for metadata errors.

### APP-019
new-change MUST return exit 2 with corrective diagnostics U01 for invalid arguments or slug, U02 for an unknown module and C13 for duplicate slugs, preserving existing files.

### APP-020
new-change --json MUST use schema 1 with string path data on success, or null data and code/message diagnostics on failure.

### APP-021
stamp MUST return exit 2 for usage or configuration errors and exit 1 for stamping rejections, with code/message diagnostics.

### APP-022
stamp --json MUST return schema 1 with string confirmation data on success, or null data and diagnostics on failure.

### APP-023
check MUST return exit 2 for usage or configuration errors with code/message diagnostics.

### APP-024
check --json MUST return schema 1 with string summary data on success, null data on failure and the reported diagnostics.

### APP-025
Without --history, show --json MUST return identical text as schema-1 string data with empty diagnostics and exit 0 on success.

### APP-026
show MUST return exit 2 with code/message diagnostics for unknown or ambiguous module or item, invalid arguments or configuration errors.

### APP-027
start MUST preserve reported warnings and return exit 1 for lifecycle rejection or exit 2 for usage or configuration errors, with corrective code/message diagnostics.

### APP-028
start --json MUST return schema 1 with context string data on success or null data on failure.

### APP-029
done [change] MUST invoke full workflow.Done with the actual invoking executable as Bin, returning 0 for success, 1 for failed checks or blocked state and 2 for usage or configuration errors; plain success names the resulting phase.
	impl := commitsAppPhase(t, root, path, "audit", "impl")
	wantPending := []string{impl, audit, evals}
	commitsAppSuccess(t, root, "manual", false, wantPending)
	commitsAppSuccess(t, root, "0001-manual.md", true, wantPending)
	archived := "changes/archive/0001-manual.md"
	if err := os.MkdirAll(filepath.Join(root, "changes", "archive"), 0755); err != nil {
		t.Fatal(err)
	}
	cliGit(t, root, "mv", path, archived)
	applied := commitsAppPhase(t, root, archived, "impl", "applied")
	wantArchived := []string{applied, impl, audit, evals}
	commitsAppSuccess(t, root, archived, false, wantArchived)
	commitsAppSuccess(t, root, "manual", true, wantArchived)
}

func commitsAppFailure(t *testing.T, root string, args []string, wantStatus int, code string) {
	t.Helper()
	jsonMode := len(args) > 0 && args[len(args)-1] == "--json"
	status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
	message := commandFailure(t, status, wantStatus, out, errOut, jsonMode)
	if jsonMode {
		if code == "" {
			message = assertFailedJSON(t, out)
		} else {
			message = assertFailedJSON(t, out, code)
		}
	} else if code != "" && !strings.HasPrefix(strings.TrimSpace(message), code+":") {
		t.Errorf("diagnostic=%q; want %s cause", message, code)
	}
	if !corrective(message) {
		t.Errorf("diagnostic %q must suggest a correction", message)
	}
}

func Test_APP_095_APP_096_CommitsFailures(t *testing.T) {
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	for _, tc := range []struct {
		name string
		args []string
		code string
	}{
		{"missing", []string{"commits"}, "U01"},
		{"extra-json", []string{"commits", "manual", "extra", "--json"}, "U01"},
		{"unsupported-option", []string{"commits", "manual", "--bogus"}, "U01"},
		{"unknown-selector", []string{"commits", "missing", "--json"}, "U02"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			commitsAppFailure(t, root, tc.args, 2, tc.code)
		})
	cliCommit(t, root, "application with four rule failures")
	oid := cliGit(t, root, "rev-parse", "HEAD")
	want := map[string]string{"H01": "locked/config.txt", "H02": "packages/alpha/api.go", "H03": "packages/alpha/ledger/entry.txt", "H04": "changes/archive/0000-old.md"}
	report, err := checks.History(root, base+"..HEAD")
	if err != nil || len(report.Diagnostics) != len(want) {
		t.Fatalf("four-rule history fixture: report=%+v err=%v", report, err)
	}
	for _, d := range report.Diagnostics {
		if want[d.Code] != d.Path || !strings.Contains(d.Message, oid) {
			t.Fatalf("unexpected history fixture diagnostic: %+v", d)
		}
	}
	nested := appNestedDir(t, root)
	status, plain, stderr := runSavedAttemptApp(t, nested, "check-commits", base+"..HEAD")
	if status != 1 || stderr == "" && plain == "" {
		t.Errorf("plain rules: exit=%d stdout=%q stderr=%q", status, plain, stderr)
	}
	for code, path := range want {
		found := false
		for _, line := range strings.Split(plain+stderr, "\n") {
			if strings.Contains(line, code) && strings.Contains(line, path) && strings.Contains(line, oid) && corrective(line) {
				found = true
			}
		}
		if !found {
			t.Errorf("plain diagnostics lack actionable %s, %s, %s: %q", code, path, oid, plain+stderr)
		}
	}
	nextCLIDiagnostic(t, status, 1, out, stderr, true, "C10")
	env := decodeEnvelope(t, out)
	if string(env["data"]) != "null" {
		t.Errorf("mixed failure data=%s; want null", env["data"])
	}
	var diagnostics []struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(env["diagnostics"], &diagnostics); err != nil || len(diagnostics) == 0 {
		t.Fatalf("mixed diagnostics=%s: %v", env["diagnostics"], err)
	}
	for _, marker := range []string{"ALP-002", "refs/fireweed/attempts/alpha/1", "8s"} {
		if !strings.Contains(diagnostics[0].Message, marker) {
			t.Errorf("mixed failure lacks %q: %q", marker, diagnostics[0].Message)
		}
	}
	runForReason(t, diagnostics[0].Message, "8s")
	alpha := string(mustReadNextCLI(t, filepath.Join(f.root, "changes/0001-alpha.md")))
	if !strings.Contains(alpha, "blocked:") || !strings.Contains(alpha, "phase: evals") || !strings.Contains(alpha, "### ALP-002") {
		t.Errorf("Alpha block state=%q", alpha)
	}
	ref := "refs/fireweed/attempts/alpha/1"
	if saved := cliGit(t, f.root, "for-each-ref", "--format=%(objectname)", ref); saved == "" || cliGit(t, f.root, "cat-file", "-t", saved) != "commit" {
		t.Errorf("missing Alpha saved attempt %s", ref)
	}
	head := cliGit(t, f.root, "rev-parse", "HEAD")
	if count := cliGit(t, f.root, "rev-list", "--count", base+"..HEAD"); count != "1" {
		t.Errorf("stop phase commits=%s; want Alpha's block only", count)
	if wantStatus == 0 {
		if stderr != "" {
			t.Errorf("success stderr=%q", stderr)
		}
		if jsonMode {
			return commandStringData(t, out)
		}
		return out
	}
	if jsonMode {
		if stderr != "" {
			t.Errorf("JSON failure stderr=%q", stderr)
		}
		message := assertFailedJSON(t, out, wantCode)
		if !corrective(message) {
			t.Errorf("diagnostic lacks correction: %q", message)
		}
		return message
	}
	if out != "" {
		t.Errorf("plain failure stdout=%q; want empty", out)
	}
	if wantCode != "" && !strings.Contains(stderr, wantCode+":") || !corrective(stderr) {
		t.Errorf("plain failure stderr=%q; want corrective %s", stderr, wantCode)
	}
	return stderr
}

func Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases(t *testing.T) {
	bin := buildAppBinary(t)
	root := checkFixFixture(t)
packages/workflow/internal/lifecycle/commits_stub.go:67:		return "", unknownChange(selector, slugs)
packages/workflow/internal/lifecycle/commits_stub.go:69:	return "", fmt.Errorf("U02: change %q is ambiguous", selector)
packages/workflow/internal/lifecycle/stub.go:153:	return project.Change{}, unknownChange(target, slugs)
packages/workflow/internal/lifecycle/stub.go:156:func unknownChange(target string, slugs []string) error {

codex
The app builds through the retained facade, and the eval hook is running. The controller’s `verify` command also returned P02 because this workspace has no held change, so the app hook is the verification path here.
exec
/bin/zsh -lc "sed -n '1,88p' packages/project/SPEC.md; sed -n '186,236p' packages/app/evals/list_cli_test.go; sed -n '261,325p' packages/app/evals/show_history_cli_test.go; sed -n '470,523p' packages/app/evals/status_cli_test.go; sed -n '112,168p' packages/app/evals/authoring_commands_test.go; rg -n 'type BlockedError|func \\(.*BlockedError' packages/runner/internal/next/run.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 508ms:
---
depends_on:
    - contracts
    - gitstore
    - assets
id_prefix: PRJ
interface:
    - api.go
    - model/**
name: project
---

# project

## Specs

### PRJ-001
Load MUST read fireweed.yaml and module SPEC.md files only in configured modules globs, returning module names, prefixes, dependencies, documents and zones.

### PRJ-002
ValidateGraph MUST return nil for a valid graph and an explanatory error identifying the rejection cause for an invalid graph.

### PRJ-003
NewChange MUST create a queued Markdown skeleton for an existing module with Motivation, Interface, Added, Modified, Removed and Decisions sections, containing no assigned spec IDs.

### PRJ-004
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets. An unknown target MUST return U02 naming the rejected selector and listing valid pending change slugs.

### PRJ-005
Project MUST preserve complete Added and Modified item text and remove Removed items when applying pending spec edits.

### PRJ-006
Stamp MUST write each Modified or Removed item's canonical SpecHash from the projection preceding the target change into its bases.

### PRJ-007
Check MUST return no error diagnostics for valid module specs and changes. Each failure diagnostic MUST have a nonempty code and message.

### PRJ-008
Check MUST reject a pending change without a nonempty Motivation section with diagnostic C07.

### PRJ-009
Check MUST reject a module spec or Added/Modified spec whose body contains no normative keyword with diagnostic C04.

### PRJ-010
Check MUST reject spec tags other than review or quarantine with a nonempty reason with diagnostic C06.

### PRJ-011
Check MUST reject duplicate spec IDs within a module's Specs section with diagnostic C05.

### PRJ-012
Check MUST reject module spec IDs and change edit IDs whose prefix belongs to another module with diagnostic C05.

### PRJ-013
Check MUST reject a change that uses the same edit identity more than once across Added, Modified and Removed with diagnostic C05.

### PRJ-014
Check MUST reject a pending change that references a handle absent from that change's Added entries with diagnostic C09.

### PRJ-015
Check MUST reject Modified or Removed entries whose bases hash is missing or
differs from the preceding projected spec's canonical hash with diagnostic C08.
An absent, null or empty-string entry MUST report `missing base for <ID>`;
a nonempty differing hash MUST report `stale base for <ID>`.

### PRJ-016
Check MUST reject Modified or Removed IDs absent from the projection preceding the change with diagnostic C08.

### PRJ-017
Load MUST inherit an omitted module zone field from module_defaults while retaining an explicitly empty module zone list as empty.

### PRJ-018
Load MUST return an error for missing or malformed config or module documents.

### PRJ-019
ValidateGraph MUST reject module names that are empty or duplicated in the graph.

### PRJ-020
ValidateGraph MUST reject module prefixes that are empty or duplicated in the graph.

### PRJ-021
ValidateGraph MUST reject a dependency that does not name a different existing module.

### PRJ-022
ValidateGraph MUST reject a graph containing a dependency cycle.

### PRJ-023
NewChange MUST use the next number after the largest pending or archived change number.

	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	listFailure(t, root, []string{"list"}, 2, "C01", "fireweed.yaml")
	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	listChange(t, root, "0001-repeat.md", "alpha", "")
	listChange(t, root, "0002-repeat.md", "alpha", "")
	_, cause := project.Queue(root)
	if cause == nil || !strings.Contains(cause.Error(), "ambiguous change slug") {
		t.Fatalf("public Queue fixture cause: %v", cause)
	}
	listFailure(t, root, []string{"list"}, 1, "C13", cause.Error())
	regularFile := appProjectFixture(t)
	appPut(t, regularFile, "changes", "not a directory")
	regularFile, err := filepath.EvalSymlinks(regularFile)
	if err != nil {
		t.Fatal(err)
	}
	_, cause = project.Queue(regularFile)
	if cause == nil {
		t.Fatal("public Queue accepted a regular file as changes directory")
	}
	hasC03 := false
	for _, diagnostic := range project.Check(regularFile) {
		if diagnostic.Code == "C03" && !diagnostic.Warning && strings.Contains(diagnostic.Path+diagnostic.Message, "changes") {
			hasC03 = true
			break
		}
	}
	if !hasC03 {
		t.Fatal("public Check did not report a C03 error for the regular-file changes path")
	}
	listFailure(t, regularFile, []string{"list"}, 1, "C03", cause.Error())
}
func Test_APP_135_ShowHistoryLookupAndArchiveFailures(t *testing.T) {
	bin := buildAppBinary(t)
	for _, tc := range []struct {
		name   string
		args   []string
		code   string
		exit   int
		named  string
		modify func(*testing.T, string)
	}{
		{"unknown-module", []string{"show", "absent-module", "--history"}, "U02", 2, "absent-module", nil},
		{"unknown-ID", []string{"show", "alpha", "ALP-999", "--history"}, "U02", 2, "ALP-999", nil},
		{"missing-config", []string{"show", "alpha", "--history"}, "C01", 2, "", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed-config", []string{"show", "alpha", "--history"}, "C01", 2, "", func(t *testing.T, root string) {
			appPut(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
		{"malformed-archive", []string{"show", "alpha", "--history"}, "C03", 1, "changes/archive/0001-invalid.md", func(t *testing.T, root string) {
			appPut(t, root, "changes/archive/0001-invalid.md", "---\nmodule: alpha\n---\n## Added\n### ALP-001 [broken\nBehavior.\n")
		}},
	} {
		for _, jsonMode := range []bool{false, true} {
			args := append([]string(nil), tc.args...)
			if jsonMode {
				args = append(args, "--json")
			}
			t.Run(tc.name+map[bool]string{false: "/plain", true: "/JSON"}[jsonMode], func(t *testing.T) {
				root := authoringFixture(t)
				if tc.modify != nil {
					tc.modify(t, root)
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				assertAppHistoryFailure(t, status, tc.exit, stdout, stderr, jsonMode, tc.code, tc.named)
			})
		}
	}
}
func Test_APP_126_StatusUsageConfigurationAndOperationalFailures(t *testing.T) {
	check := func(root string, args []string, wantExit int, code, cause string) {
		t.Helper()
		status, out, errOut := listRun(t, root, args...)
		if status != wantExit || !strings.Contains(out+errOut, code+":") || !strings.Contains(out+errOut, cause) || !corrective(out+errOut) {
			t.Errorf("plain %v exit=%d stdout=%q stderr=%q; want %d %s %q action", args, status, out, errOut, wantExit, code, cause)
		}
		status, out, errOut = listRun(t, root, append(append([]string{}, args...), "--json")...)
		if status != wantExit || errOut != "" {
			t.Errorf("JSON %v exit=%d stdout=%q stderr=%q", args, status, out, errOut)
		}
		message := assertFailedJSON(t, out, code)
		if !strings.Contains(message, cause) || !corrective(message) {
			t.Errorf("JSON %v message=%q; want %q and action", args, message, cause)
		}
	}
	root := appProjectFixture(t)
	check(root, []string{"status", "extra"}, 2, "U01", "extra")
	appPut(t, root, "fireweed.yaml", "modules: [broken\n")
	check(root, []string{"status"}, 2, "C01", "fireweed.yaml")
	appPut(t, root, "fireweed.yaml", appFixtureConfig)
	listChange(t, root, "0001-repeat.md", "alpha", "")
	listChange(t, root, "0002-repeat.md", "beta", "")
	cliGit(t, root, "init", "-q", "-b", "main")
	cliCommit(t, root, "duplicate queue")
	if _, err := workflow.Status(root); err == nil {
		t.Fatal("public Status accepted duplicate slugs")
	}
	check(root, []string{"status"}, 1, "C13", "repeat")
	baseRoot := appProjectFixture(t)
	appPut(t, baseRoot, "fireweed.yaml", strings.Replace(appFixtureConfig, "base: main", "base: unavailable-base", 1))
	cliGit(t, baseRoot, "init", "-q", "-b", "main")
	cliCommit(t, baseRoot, "unavailable base")
	if _, err := workflow.Status(baseRoot); err == nil {
		t.Fatal("public Status accepted unavailable base ref")
	}
	check(baseRoot, []string{"status"}, 1, "P01", "unavailable-base")
}
		args []string
		want int
		edit func(*testing.T, string)
	}{
		{"missing-argument", []string{"stamp"}, 2, nil},
		{"extra-argument", []string{"stamp", "revise", "extra"}, 2, nil},
		{"invalid-config", []string{"stamp", "revise"}, 2, func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"missing-change", []string{"stamp", "absent"}, 1, nil},
		{"absent-modified-ID", []string{"stamp", "revise"}, 1, func(t *testing.T, root string) {
			authoringChange(t, root, "0001-revise.md", "Users need a revision.\n", "", "### ALP-999\nMissing MUST change.\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := authoringFixture(t)
			if tc.edit != nil {
				tc.edit(t, root)
			}
			for _, jsonMode := range []bool{false, true} {
				args := append([]string{}, tc.args...)
				if jsonMode {
					args = append(args, "--json")
				}
				status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
				commandFailure(t, status, tc.want, stdout, stderr, jsonMode)
			}
		})
	}
}

func Test_APP_007_APP_024_CheckValidAndAllDiagnostics(t *testing.T) {
	bin := buildAppBinary(t)
	root := authoringFixture(t)
	authoringChange(t, root, "0001-valid.md", "Users need this addition.\n", "### new:valid\nValid MUST work.\n", "")
	for _, jsonMode := range []bool{false, true} {
		args := []string{"check"}
		if jsonMode {
			args = append(args, "--json")
		}
		status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), args...)
		if status != 0 || stderr != "" {
			t.Errorf("valid check: status=%d stdout=%q stderr=%q", status, stdout, stderr)
		}
		if jsonMode {
			envelope := decodeEnvelope(t, stdout)
			var summary string
			if err := json.Unmarshal(envelope["data"], &summary); err != nil || strings.TrimSpace(summary) == "" {
				t.Error("JSON success needs a summary")
			}
			var diagnostics []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			}
			if err := json.Unmarshal(envelope["diagnostics"], &diagnostics); err != nil || diagnostics == nil {
				t.Errorf("diagnostics=%s; want array: %v", envelope["diagnostics"], err)
			}
			for _, diagnostic := range diagnostics {
				if diagnostic.Code == "" || diagnostic.Message == "" {
234:type BlockedError struct {
240:func (e *BlockedError) Error() string {
244:func (e *BlockedError) Unwrap() error { return e.Cause }

exec
/bin/zsh -lc "sed -n '710,760p' packages/runner/internal/drill/drill.go; sed -n '785,832p' packages/runner/internal/drill/drill.go; rg -n 'contains.*phase|empty|next' packages/app/evals/next_cli_test.go | tail -24; ls -la .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 495ms:
			result.Dependents[name] = report
		}
		module := findModule(snapshot, name)
		items, specErr := specs(*module)
		err = errors.Join(err, specErr)
		if specErr == nil {
			err = errors.Join(err, checks.RequirePass(report, items), checks.RequireCoverage(report, items))
		}
		if err != nil {
			failures = append(failures, fmt.Errorf("%s eval: %w", name, err))
		}
	}
	return errors.Join(failures...)
}
func failureDetails(result model.DrillResult) string {
	var b strings.Builder
	write := func(name string, report checks.Report) {
		fmt.Fprintf(&b, "\n%s stdout: %s\n%s stderr: %s\n", name, report.Stdout, name, report.Stderr)
		for _, c := range report.Cases {
			if c.Result == "fail" {
				fmt.Fprintf(&b, "%s %s %s: %s\n", name, c.Classname, c.Name, c.Failure)
			}
		}
	}
	write(result.Module, result.Own)
	names := make([]string, 0, len(result.Dependents))
	for name := range result.Dependents {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		write(name, result.Dependents[name])
	}
	return b.String()
}
func retain(scratch, root, ref string, target project.Module) error {
	if _, err := git(scratch, "add", "-A", "--", target.Dir); err != nil {
		return err
	}
	if _, err := git(scratch, "-c", "user.name=Fireweed Drill", "-c", "user.email=drill@localhost", "commit", "-q", "--allow-empty", "-m", "rebuilt drill code"); err != nil {
		return err
	}
	oid, err := git(scratch, "rev-parse", "HEAD")
	if err != nil {
		return err
	}
	_, err = git(root, "update-ref", ref, oid)
	return err
}

func configuredTimeout(config map[string]any) time.Duration {
		if report.Stderr != "" {
			fmt.Fprintf(&b, "  Stderr: %s\n", report.Stderr)
		}
	}
	writeCases(result.Module, result.Own)
	b.WriteString("\n## Dependent evals\n")
	names := make([]string, 0, len(result.Dependents))
	for name := range result.Dependents {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		writeCases(name, result.Dependents[name])
	}
	var contradicted, gaps string
	if input != "" {
		if notes, err := readInputFile(input, "DRILL_NOTES.md"); err == nil && notes.mode.IsRegular() {
			contradicted, _ = contracts.Section(notes.content, "Contradicted Decisions")
			gaps, _ = contracts.Section(notes.content, "Gaps")
		}
	}
	fmt.Fprintf(&b, "\n## Contradicted Decisions\n%s\n## Gaps\n%s\n## Agent observations\n", strings.TrimSpace(contradicted), strings.TrimSpace(gaps))
	for i, step := range result.Steps {
		fmt.Fprintf(&b, "\n### Attempt %d (exit %d)\n%s\n%s\n", i+1, step.ExitCode, step.Stdout, step.Stderr)
	}
	if failure != nil {
		fmt.Fprintf(&b, "\n## Verification failures\n%s\n", failure)
	}
	path := filepath.Join(root, filepath.FromSlash(result.ReportPath))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(b.String()), 0644)
}

func markReportFailure(root, relative string, cause error) error {
	path := filepath.Join(root, filepath.FromSlash(relative))
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	report := strings.Replace(string(data), "result: pass\n", "result: fail\n", 1)
	report += fmt.Sprintf("\n## Cleanup failure\n%s\n", cause)
	return os.WriteFile(path, []byte(report), 0644)
}

func configuredAttempts(config map[string]any) int {
	switch v := config["attempts"].(type) {
264:		t.Errorf("empty queue exit=%d stdout=%q stderr=%q", status, out, errOut)
267:	status, out, errOut = runNextCLI(t, bin, root, "next")
268:	nextCLIDiagnostic(t, status, 2, out, errOut, false, "C01")
269:	status, out, errOut = runNextCLI(t, bin, root, "next", "extra")
270:	nextCLIDiagnostic(t, status, 2, out, errOut, false, "U01")
271:	failed, _ := nextCLIProject(t, bin, true, false)
272:	status, out, errOut = runNextCLI(t, bin, failed, "next")
273:	nextCLIDiagnostic(t, status, 1, out, errOut, false, "P01")
278:	root, _ := nextCLIProject(t, bin, true, true)
281:	status, out, errOut := runNextCLI(t, bin, root, "next", "--json")
282:	nextCLIDiagnostic(t, status, 1, out, errOut, true, "C10")
290:	root, _ := nextCLIProject(t, bin, false, true)
291:	status, out, errOut := runNextCLI(t, bin, root, "next", "--json")
293:		t.Fatalf("JSON empty exit=%d stdout=%q stderr=%q", status, out, errOut)
300:		Empty  bool   `json:"empty"`
304:		t.Errorf("empty data is not an object: %s: %v", env["data"], err)
306:	for _, key := range []string{"change", "run", "phase", "empty"} {
308:			t.Errorf("empty data missing %q: %s", key, env["data"])
312:		t.Errorf("empty data=%s: %v", env["data"], err)
318:	status, out, errOut = runNextCLI(t, bin, root, "next", "extra", "--json")
319:	nextCLIDiagnostic(t, status, 2, out, errOut, true, "U01")
320:	applied, _ := nextCLIProject(t, bin, true, true)
321:	status, out, errOut = runNextCLI(t, bin, applied, "next", "--json")
330:	for _, key := range []string{"change", "run", "phase", "empty"} {
total 0
drwxr-xr-x@ 2 joel  staff  64 Nov 10 12:22 .
drwxr-xr-x@ 3 joel  staff  96 Nov 10 12:22 ..

 exited 1 in 265462ms:
?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	262.163s
exit status 1

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/entry.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..351b3bf088806acf3069b87c0fcdbf7443c0508c
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
+	}
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
 	}
-	return fallback, message
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -313,10 +341,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/setup.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..351b3bf088806acf3069b87c0fcdbf7443c0508c
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
+	}
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
 	}
-	return fallback, message
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -313,10 +341,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/project.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..351b3bf088806acf3069b87c0fcdbf7443c0508c
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
+	}
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
 	}
-	return fallback, message
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -313,10 +341,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..c8c77d4fe35e4ddc2fc1e4c97fbbeda311d38468
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -127,14 +121,10 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

exec
/bin/zsh -lc 'gofmt -w packages/app/internal/entry; ls -l .fireweed/runs/drill; git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 520ms:
total 144
-rw-r--r--@ 1 joel  staff  70695 Nov 10 12:26 report.xml
?? packages/app/internal/

exec
/bin/zsh -lc "python3 - <<'PY'
import xml.etree.ElementTree as E
root=E.parse('.fireweed/runs/drill/report.xml').getroot()
cases=root.findall('testcase')
failures=[(t.attrib['name'], t.find('failure').text) for t in cases if t.find('failure') is not None]
print(len(cases),'cases;',len(failures),'failures;',len(root.findall('.//skipped')),'skipped')
for name, message in failures:
 print(name, message,sep='\\n')
PY" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 719ms:
323 cases; 13 failures; 0 skipped
Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases
=== RUN   Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases
    check_fix_cli_test.go:119: plain failure stderr="U02: pending change \"\" not found; valid pending change slugs: first, second. To correct this, follow the guidance for U02. An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.\n"; want corrective C08
    check_fix_cli_test.go:121: post-fix diagnostics="U02: pending change \"\" not found; valid pending change slugs: first, second. To correct this, follow the guidance for U02. An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.\n"; want only preserved stale conflict
    check_fix_cli_test.go:133: first change lost recorded bases or human content: bases=map[ALP-001:deadbeef ALP-002:4cb10a55 ALP-999:extra-base] doc={Frontmatter:map[bases:map[ALP-001:deadbeef ALP-002:4cb10a55 ALP-999:extra-base] custom:keep-first module:alpha] Body:
        ## Motivation
        Keep first human motivation.
        
        ## Interface
        
        ## Added
        
        ## Modified
        ### ALP-001
        One MUST change.
        
        ### ALP-003
        Three MUST change.
        
        ## Removed
        ### ALP-002
        
        ### ALP-004
        
        ## Decisions
        }
    check_fix_cli_test.go:135: changes/0002-second.md bases=<nil>
--- FAIL: Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases (0.55s)

Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp
=== RUN   Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp
    check_fix_cli_test.go:162: check --fix: exit=1 stdout="" stderr="U02: pending change \"\" not found; valid pending change slugs: first. To correct this, follow the guidance for U02. An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.\n"; want 0
    check_fix_cli_test.go:162: success stderr="U02: pending change \"\" not found; valid pending change slugs: first. To correct this, follow the guidance for U02. An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.\n"
    check_fix_cli_test.go:164: plain success=""
    check_fix_cli_test.go:166: changes/0001-first.md bases=<nil>
--- FAIL: Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp (0.56s)

Test_APP_148_APP_149_CheckFixRejectsBeforeMutation
=== RUN   Test_APP_148_APP_149_CheckFixRejectsBeforeMutation
    check_fix_cli_test.go:201: plain failure stderr="U02: pending change \"\" not found; valid pending change slugs: first. To correct this, follow the guidance for U02. An unknown lookup name was given to explain or playbook, an unknown or ambiguous requested module, item or change was given to a lookup, or a projected local handle is ambiguous. Select a valid name or code from the listed explain codes or playbook names; specify an existing, fully qualified module, item or change ID to resolve a requested lookup. An unknown recipe name was requested for init. Choose a supported recipe name or select a readable local recipe directory.\n"; want corrective C08
    check_fix_cli_test.go:201: diagnostic code="U02"; want "C08"
--- FAIL: Test_APP_148_APP_149_CheckFixRejectsBeforeMutation (0.61s)

Test_APP_037_StaticCauses
=== RUN   Test_APP_037_StaticCauses
--- FAIL: Test_APP_037_StaticCauses (0.76s)

Test_APP_037_StaticCauses/base
=== RUN   Test_APP_037_StaticCauses/base
    diagnostic_causes_test.go:105: missing C08 cause "absent from preceding projection" in plain output "P01: changes/0001-revise.md: ID ALP-999 absent from preceding projection. To correct this, follow the guidance for P01. A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.\n"
    diagnostic_causes_test.go:105: missing C08 cause "absent from preceding projection" in {"schema":1,"data":null,"diagnostics":[{"code":"P01","message":"changes/0001-revise.md: ID ALP-999 absent from preceding projection. To correct this, follow the guidance for P01. A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/base (0.01s)

Test_APP_037_StaticCauses/metadata
=== RUN   Test_APP_037_StaticCauses/metadata
    diagnostic_causes_test.go:105: missing C12 cause "set must be a mapping" in plain output "C02: changes/0001-first.md: set must be a mapping. To correct this, follow the guidance for C02. The module graph is invalid because of an empty module name, a duplicate module name, an empty module prefix, a duplicate module prefix, or a dependency that names an unknown module or creates a cycle. Correct module names and prefixes in SPEC.md so they are nonempty and unique. Correct the module dependency in SPEC.md so it names an existing module and the graph is acyclic.\n"
    diagnostic_causes_test.go:105: missing C12 cause "set must be a mapping" in {"schema":1,"data":null,"diagnostics":[{"code":"C02","message":"changes/0001-first.md: set must be a mapping. To correct this, follow the guidance for C02. The module graph is invalid because of an empty module name, a duplicate module name, an empty module prefix, a duplicate module prefix, or a dependency that names an unknown module or creates a cycle. Correct module names and prefixes in SPEC.md so they are nonempty and unique. Correct the module dependency in SPEC.md so it names an existing module and the graph is acyclic.","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/metadata (0.01s)

Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI
=== RUN   Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI
    doctor_setup_test.go:174: APP-067 missing named passing check baseline: {"schema":1,"data":{"checks":[{"name":"baseline global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval and coverage pass"},{"name":"failing_eval global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected failing_eval"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"failing_eval: eval and coverage pass"},{"name":"missing_id global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected missing_id"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"missing_id: eval and coverage pass"},{"name":"baseline global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries pass"},{"name":"undeclared global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries detected undeclared"},{"name":"recovery global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"undeclared: boundaries pass"},{"name":"internal_eval global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"boundaries detected internal_eval"},{"name":"recovery global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"internal_eval: boundaries pass"}]},"diagnostics":[]}
    doctor_setup_test.go:174: APP-067 missing named passing check undeclared: {"schema":1,"data":{"checks":[{"name":"baseline global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval and coverage pass"},{"name":"failing_eval global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected failing_eval"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"failing_eval: eval and coverage pass"},{"name":"missing_id global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected missing_id"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"missing_id: eval and coverage pass"},{"name":"baseline global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries pass"},{"name":"undeclared global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries detected undeclared"},{"name":"recovery global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"undeclared: boundaries pass"},{"name":"internal_eval global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"boundaries detected internal_eval"},{"name":"recovery global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"internal_eval: boundaries pass"}]},"diagnostics":[]}
    doctor_setup_test.go:174: APP-067 missing named passing check internal_eval: {"schema":1,"data":{"checks":[{"name":"baseline global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval and coverage pass"},{"name":"failing_eval global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected failing_eval"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"failing_eval: eval and coverage pass"},{"name":"missing_id global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected missing_id"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"missing_id: eval and coverage pass"},{"name":"baseline global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries pass"},{"name":"undeclared global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries detected undeclared"},{"name":"recovery global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"undeclared: boundaries pass"},{"name":"internal_eval global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"boundaries detected internal_eval"},{"name":"recovery global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"internal_eval: boundaries pass"}]},"diagnostics":[]}
    doctor_setup_test.go:174: APP-067 missing named passing check failing_eval: {"schema":1,"data":{"checks":[{"name":"baseline global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval and coverage pass"},{"name":"failing_eval global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected failing_eval"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"failing_eval: eval and coverage pass"},{"name":"missing_id global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected missing_id"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"missing_id: eval and coverage pass"},{"name":"baseline global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries pass"},{"name":"undeclared global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries detected undeclared"},{"name":"recovery global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"undeclared: boundaries pass"},{"name":"internal_eval global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"boundaries detected internal_eval"},{"name":"recovery global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"internal_eval: boundaries pass"}]},"diagnostics":[]}
    doctor_setup_test.go:174: APP-067 missing named passing check missing_id: {"schema":1,"data":{"checks":[{"name":"baseline global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval and coverage pass"},{"name":"failing_eval global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected failing_eval"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"failing_eval: eval and coverage pass"},{"name":"missing_id global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected missing_id"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"missing_id: eval and coverage pass"},{"name":"baseline global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries pass"},{"name":"undeclared global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries detected undeclared"},{"name":"recovery global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"undeclared: boundaries pass"},{"name":"internal_eval global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"boundaries detected internal_eval"},{"name":"recovery global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"internal_eval: boundaries pass"}]},"diagnostics":[]}
    doctor_setup_test.go:174: APP-067 missing named passing check recovery: {"schema":1,"data":{"checks":[{"name":"baseline global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval and coverage pass"},{"name":"failing_eval global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected failing_eval"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"failing_eval: eval and coverage pass"},{"name":"missing_id global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"eval detected missing_id"},{"name":"recovery global eval","module":"doctor4af2ec4cfc","passed":true,"detail":"missing_id: eval and coverage pass"},{"name":"baseline global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries pass"},{"name":"undeclared global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"boundaries detected undeclared"},{"name":"recovery global boundaries","module":"doctor4af2ec4cfc","passed":true,"detail":"undeclared: boundaries pass"},{"name":"internal_eval global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"boundaries detected internal_eval"},{"name":"recovery global boundaries","module":"doctorcfb37ee5fd","passed":true,"detail":"internal_eval: boundaries pass"}]},"diagnostics":[]}
--- FAIL: Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI (40.58s)

Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors
=== RUN   Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors
    list_cli_test.go:185: plain [list extra] exit=2 stdout="" stderr="U01: invalid arguments for list. Usage: fireweed list\n"; want U01 cause "extra" and action
    list_cli_test.go:185: JSON [list extra --json] diagnostic="invalid arguments for list. Usage: fireweed list"; want cause "extra" and action
    list_cli_test.go:195: plain [list] exit=1 stdout="" stderr="P01: ambiguous change slug \"repeat\". To correct this, follow the guidance for P01. A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.\n"; want C13 cause "ambiguous change slug \"repeat\"" and action
    list_cli_test.go:195: diagnostic code="P01"; want "C13"
    list_cli_test.go:216: plain [list] exit=1 stdout="" stderr="P01: read changes: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors4116027568/002/changes: not a directory. To correct this, follow the guidance for P01. A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.\n"; want C03 cause "read changes: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors4116027568/002/changes: not a directory" and action
    list_cli_test.go:216: diagnostic code="P01"; want "C03"
--- FAIL: Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors (0.01s)

Test_APP_089_RunForInvalidOptionsBeforeMutation
=== RUN   Test_APP_089_RunForInvalidOptionsBeforeMutation
--- FAIL: Test_APP_089_RunForInvalidOptionsBeforeMutation (0.89s)

Test_APP_089_RunForInvalidOptionsBeforeMutation/next
=== RUN   Test_APP_089_RunForInvalidOptionsBeforeMutation/next
    run_for_cli_test.go:276: diagnostic lacks valid run --for usage: "unknown option \"--for\" for next. Usage: fireweed next [--allow-unmerged]"
--- FAIL: Test_APP_089_RunForInvalidOptionsBeforeMutation/next (0.04s)

Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures
=== RUN   Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures
--- FAIL: Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures (0.64s)

Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/invalid-graph
=== RUN   Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/invalid-graph
    show_stale_cli_test.go:221: diagnostic code="U02"; want "C02"
--- FAIL: Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/invalid-graph (0.01s)

Test_APP_126_StatusUsageConfigurationAndOperationalFailures
=== RUN   Test_APP_126_StatusUsageConfigurationAndOperationalFailures
    status_cli_test.go:487: plain [status extra] exit=2 stdout="" stderr="U01: invalid arguments for status. Usage: fireweed status\n"; want 2 U01 "extra" action
    status_cli_test.go:487: JSON [status extra] message="invalid arguments for status. Usage: fireweed status"; want "extra" and action
    status_cli_test.go:498: plain [status] exit=1 stdout="" stderr="P01: ambiguous change slug \"repeat\". To correct this, follow the guidance for P01. A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.\n"; want 1 C13 "repeat" action
    status_cli_test.go:498: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_126_StatusUsageConfigurationAndOperationalFailures (0.51s)


exec
/bin/zsh -lc "rg -n 'func withCause|func errorCode|type .*Failure' packages/project/internal/queue; sed -n '103,175p' packages/project/internal/queue/stale.go; sed -n '135,181p' packages/app/evals/doctor_setup_test.go; sed -n '175,228p' packages/app/evals/show_stale_cli_test.go; sed -n '210,258p' packages/project/internal/queue/project.go; ls -l .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 483ms:
packages/project/internal/queue/cause.go:10:type causeFailure struct {
packages/project/internal/queue/cause.go:18:func withCause(code string, err error) error { return &causeFailure{code: code, err: err} }
packages/project/internal/queue/cause.go:20:func errorCode(err error, fallback string) string {
	}
	return stale, nil
}

// StampMissing fills absent bases without changing recorded values.
func StampMissing(root, target string) error {
	change, p, err := preceding(root, target)
	if err != nil {
		return err
	}
	old, err := changeBases(change)
	if err != nil {
		return err
	}
	bases := make(map[string]any, len(old)+len(edited(change)))
	for id, hash := range old {
		bases[id] = hash
	}
	missing := false
	for _, edit := range edited(change) {
		base, ok := p.find(change.Module, edit.ID)
		if !ok {
			return contracts.CodedError{Code: "C08", Subject: change.Path, Cause: fmt.Errorf("%s: ID %s absent from preceding projection", change.Path, edit.ID)}
		}
		if _, exists := bases[edit.ID]; !exists {
			bases[edit.ID] = contracts.SpecHash(base)
			missing = true
		}
	}
	if !missing {
		return nil
	}
	change.Document.Frontmatter["bases"] = bases
	content, err := contracts.RenderDocument(change.Document)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, filepath.FromSlash(change.Path)), []byte(content), 0644)
}

// stampedProjection uses the commit that introduced the current bases mapping.
func stampedProjection(root string, change model.Change, bases map[string]any) (*projection, bool) {
	out, err := staleGit(root, "log", "--format=%H", "HEAD", "--", change.Path)
	if err != nil {
		return nil, false
	}
	commits := strings.Fields(string(out))
	if len(commits) == 0 {
		return nil, false
	}
	var stamped string
	for i, commit := range commits {
		data, err := gitstore.ReadAt(root, commit, change.Path)
		if err != nil {
			return nil, false
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil {
			return nil, false
		}
		recorded, exists := doc.Frontmatter["bases"]
		if !exists {
			recorded = map[string]any{}
		}
		if i == 0 && !reflect.DeepEqual(recorded, bases) {
			return nil, false
		}
		if !reflect.DeepEqual(recorded, bases) {
			break
		}
		stamped = commit
	}
	if stamped == "" {
	if status != 0 {
		t.Fatalf("APP-062/066 doctor status=%d stdout=%q stderr=%q", status, plain, errOut)
	}
	plain += errOut
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !doctorCLIPlainOutcome(plain, name, "pass") {
			t.Errorf("APP-065 plain output lacks outcome %s: %q", name, plain)
		}
	}
	probeLog := doctorCLIFile(t, log)
	if !strings.Contains(probeLog, "setup:") || !strings.Contains(probeLog, bin) || strings.Contains(probeLog, "worker:") || strings.Contains(probeLog, "audit:") {
		t.Errorf("APP-062/063/064 scratch hook did not query NEW or agent ran: %q", probeLog)
	}
	status, jsonOut, errOut := runAppBinary(t, bin, nested, "doctor", "--no-agent", "--json")
	if status != 0 || errOut != "" {
		t.Fatalf("APP-067 JSON status=%d stdout=%q stderr=%q", status, jsonOut, errOut)
	}
	checks, diagnostics := doctorCLIChecks(t, jsonOut)
	for _, check := range checks {
		var name string
		var passed bool
		if err := json.Unmarshal(check["name"], &name); err != nil || name == "" {
			t.Errorf("APP-067 unnamed check: %v", check)
			continue
		}
		if err := json.Unmarshal(check["passed"], &passed); err != nil {
			t.Errorf("APP-067 check %s lacks outcome: %v", name, check)
			continue
		}
		outcome := "fail"
		if passed {
			outcome = "pass"
		}
		if !doctorCLIPlainOutcome(plain, name, outcome) {
			t.Errorf("APP-065 plain output lacks %s outcome for reported %s: %q", outcome, name, plain)
		}
	}
	for _, name := range []string{"baseline", "undeclared", "internal_eval", "failing_eval", "missing_id", "recovery"} {
		if !doctorCLIHasCheck(checks, name, true) {
			t.Errorf("APP-067 missing named passing check %s: %s", name, jsonOut)
		}
	}
	if len(diagnostics) != 0 {
		t.Errorf("APP-067 healthy diagnostics=%v", diagnostics)
	}
	if strings.Contains(doctorCLIFile(t, log), "worker:") || strings.Contains(doctorCLIFile(t, log), "audit:") {
		t.Error("APP-064 --no-agent ran worker or audit")
	}
	for _, tc := range []struct {
		name, selector, code string
		exit                 int
	}{
		{"unknown", "absent", "U02", 2},
		{"ambiguous", "target", "C13", 2},
		{"bad-bases", "changes/0001-target.md", "C08", 1},
	} {
		t.Run(tc.name, func(t *testing.T) {
			fixture := staleCLIProject(t)
			staleCLIChange(t, fixture, "changes/0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
			if tc.name == "ambiguous" {
				staleCLIChange(t, fixture, "changes/0002-target.md", "", "Another.\n", "", "")
			}
			if tc.name == "bad-bases" {
				staleCLIChange(t, fixture, "changes/0001-target.md", "bases: [invalid]\n", "Needed.\n", "### ALP-001\nProposed MUST differ.\n", "")
			}
			for _, jsonMode := range []bool{false, true} {
				args := []string{"show", "--stale", tc.selector}
				if jsonMode {
					args = append(args, "--json")
				}
				status, out, stderr := runAppBinary(t, bin, appNestedDir(t, fixture), args...)
				assertAppHistoryFailure(t, status, tc.exit, out, stderr, jsonMode, tc.code, tc.selector)
			}
		})
	}
	for _, bad := range []struct {
		name, code string
		edit       func(*testing.T, string)
	}{
		{"missing-config", "C01", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed-config", "C01", func(t *testing.T, root string) { appPut(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"invalid-graph", "C02", func(t *testing.T, root string) {
			appPut(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		}},
	} {
		t.Run(bad.name, func(t *testing.T) {
			fixture := staleCLIProject(t)
			bad.edit(t, fixture)
			status, out, stderr := runAppBinary(t, bin, appNestedDir(t, fixture), "show", "--stale", "target", "--json")
			assertAppHistoryFailure(t, status, 2, out, stderr, true, bad.code, "")
		})
	}
}
	}
	return p.result(), nil
}

func Stamp(root, target string) error {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return err
	}
	changes, err := readQueue(root, snapshot)
	if err != nil {
		return err
	}
	index, err := resolve(changes, target)
	if err != nil {
		return err
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return err
	}
	for _, preceding := range changes[:index] {
		if err := p.apply(preceding); err != nil {
			return err
		}
	}
	change := changes[index]
	bases := map[string]any{}
	if raw, exists := change.Document.Frontmatter["bases"]; exists {
		current, ok := raw.(map[string]any)
		if !ok {
			return contracts.CodedError{Code: "C08", Subject: change.Path, Cause: fmt.Errorf("%s bases must be a mapping", change.Path)}
		}
		for key, value := range current {
			bases[key] = value
		}
	}
	for _, edit := range append(append([]contracts.Item{}, change.Modified...), change.Removed...) {
		base, ok := p.find(change.Module, edit.ID)
		if !ok {
			return fmt.Errorf("%s: ID %s absent from preceding projection", change.Path, edit.ID)
		}
		bases[edit.ID] = contracts.SpecHash(base)
	}
	change.Document.Frontmatter["bases"] = bases
	content, err := contracts.RenderDocument(change.Document)
	if err != nil {
		return err
	}
total 144
-rw-r--r--@ 1 joel  staff  70695 Nov 10 12:26 report.xml

exec
/bin/zsh -lc "cat packages/project/internal/queue/cause.go; sed -n '100,135p' packages/app/evals/doctor_setup_test.go; sed -n '750,788p' packages/runner/internal/drill/drill.go; rg -n 'doctor.*result|func \\(.*(pass|fail)|Name:' packages/checks/internal/engine/doctor*.go | head -55; sed -n '195,254p' packages/runner/internal/next/run.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 477ms:
package queue

import (
	"errors"

	"github.com/lurkshark/fireweed/packages/contracts"
)

// causeFailure keeps a rejection's static cause without changing its message.
type causeFailure struct {
	code string
	err  error
}

func (e *causeFailure) Error() string { return e.err.Error() }
func (e *causeFailure) Unwrap() error { return e.err }

func withCause(code string, err error) error { return &causeFailure{code: code, err: err} }

func errorCode(err error, fallback string) string {
	var failure *causeFailure
	if errors.As(err, &failure) {
		return failure.code
	}
	return fallback
}

func itemErrorCode(err error) string {
	if code, _, ok := contracts.ErrorCode(err); ok {
		return code
	}
	return "C03"
}
	}
	return data.Checks, diagnostics
}

func doctorCLIHasCheck(checks []map[string]json.RawMessage, name string, passed bool) bool {
	for _, check := range checks {
		var gotName string
		var gotPassed bool
		_ = json.Unmarshal(check["name"], &gotName)
		_ = json.Unmarshal(check["passed"], &gotPassed)
		if gotName == name && gotPassed == passed {
			for _, field := range []string{"module", "detail"} {
				if _, ok := check[field]; !ok {
					return false
				}
			}
			return true
		}
	}
	return false
}

func doctorCLIPlainOutcome(output, name, outcome string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.Contains(line, name) && strings.Contains(strings.ToLower(line), outcome) {
			return true
		}
	}
	return false
}

func Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI(t *testing.T) {
	bin := buildAppBinary(t)
	root, nested, log := doctorCLIFixture(t, bin)
	status, plain, errOut := runAppBinary(t, bin, nested, "doctor", "--no-agent")
	if status != 0 {
		return err
	}
	oid, err := git(scratch, "rev-parse", "HEAD")
	if err != nil {
		return err
	}
	_, err = git(root, "update-ref", ref, oid)
	return err
}

func configuredTimeout(config map[string]any) time.Duration {
	if raw, ok := config["timeout"].(string); ok {
		if duration, err := time.ParseDuration(raw); err == nil && duration > 0 {
			return duration
		}
	}
	return 90 * time.Minute
}
func writeReport(root, input string, result model.DrillResult, failure error) error {
	date, err := time.Parse("20060102T150405.000000000", result.Date)
	if err != nil {
		return fmt.Errorf("drill report date: %w", err)
	}
	var b strings.Builder
	fmt.Fprintf(&b, "---\ndate: %s\nbase_commit: %s\nagent: %s\nmodel: %s\nresult: %s\n---\n\n## Own evals\n", date.UTC().Format(time.RFC3339Nano), result.Base, result.Agent, result.Model, result.Result)
	writeCases := func(name string, report checks.Report) {
		for _, c := range report.Cases {
			fmt.Fprintf(&b, "- %s | %s | %s | %s\n", name, c.Name, strings.Join(c.IDs, ", "), c.Result)
			if c.Failure != "" {
				fmt.Fprintf(&b, "  Failure: %s\n", c.Failure)
			}
		}
		if report.Stdout != "" {
			fmt.Fprintf(&b, "  Stdout: %s\n", report.Stdout)
		}
		if report.Stderr != "" {
			fmt.Fprintf(&b, "  Stderr: %s\n", report.Stderr)
		}
	}
64:	d.report.Checks = append(d.report.Checks, model.DoctorCheck{Name: name, Module: module, Passed: ok, Detail: detail})
66:func (d *doctorState) fail(name, module, hook, detail string) {
			if done.State.Phase != "" {
				result.State = withRecoveryWarnings(done.State, recoveryWarnings)
			}
			if stop := budget.TimeStop(ctx); stop != nil {
				return result, stop
			}
			return result, fmt.Errorf("complete %s: %w", phase, err)
		}
		state = done.State
		result.State = withRecoveryWarnings(state, recoveryWarnings)
	}
	return result, nil
}

func recoverAgentCommits(root, phaseBase string) (bool, error) {
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return false, err
	}
	if repo.Head == phaseBase {
		return false, nil
	}
	reset, err := gitstore.Run(context.Background(), gitstore.Request{
		Args:    []string{"git", "-C", root, "-c", "core.hooksPath=/dev/null", "reset", "--soft", phaseBase},
		Timeout: 30 * time.Second,
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
	})
	if err != nil {
		return false, fmt.Errorf("git reset --soft: %w: %s", err, strings.TrimSpace(reset.Stderr))
	}
	return true, nil
}

func withRecoveryWarnings(state workflow.State, warnings []project.Diagnostic) workflow.State {
	state.Diagnostics = append(state.Diagnostics, warnings...)
	return state
}

// BlockedError identifies the verification failure returned after Block succeeds.
type BlockedError struct {
	Phase      string
	AttemptRef string
	Cause      error
}

func (e *BlockedError) Error() string {
	return fmt.Sprintf("verify %s after exhausted attempts; saved work at %s: %v", e.Phase, e.AttemptRef, e.Cause)
}

func (e *BlockedError) Unwrap() error { return e.Cause }

func configuredAttempts(config map[string]any) int {
	switch attempts := config["attempts"].(type) {
	case int:
		if attempts > 0 {
			return attempts
		}
	case float64:
		if attempts > 0 && attempts <= float64(int(^uint(0)>>1)) && attempts == float64(int(attempts)) {
			return int(attempts)

codex
The first eval run passed 310 of 323 cases. The failures showed that `StampMissing` takes one change at a time, Doctor JSON expects short probe names, and several project errors expose their codes only through `project.Check`. I’ll fix those translations and rerun the affected cases.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/entry.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..35ae94741427288cb4c27a7ee15e4ff16e049ec7
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..c8c77d4fe35e4ddc2fc1e4c97fbbeda311d38468
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -127,14 +121,10 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/project.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..35ae94741427288cb4c27a7ee15e4ff16e049ec7
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..7477daa212223f4c65b1369045e486a948c3b523
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,13 +105,19 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
+			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
 			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -156,7 +155,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,6 +226,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -237,7 +237,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -291,7 +291,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +351,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/project.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..35ae94741427288cb4c27a7ee15e4ff16e049ec7
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..1c60a9b418a85152ffbf98e96c7b445ba2976756
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,14 +105,20 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
 			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
+			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
 	case "migrate":
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -149,6 +148,7 @@
 }
 
 func show(snap project.Snapshot, module, id string, projected bool) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -156,7 +156,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,6 +227,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -237,7 +238,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -291,7 +292,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +352,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/reports.go
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..35ae94741427288cb4c27a7ee15e4ff16e049ec7
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..1c60a9b418a85152ffbf98e96c7b445ba2976756
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,14 +105,20 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
 			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
+			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
 	case "migrate":
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -149,6 +148,7 @@
 }
 
 func show(snap project.Snapshot, module, id string, projected bool) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -156,7 +156,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,6 +227,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -237,7 +238,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -291,7 +292,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +352,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/reports.go b/packages/app/internal/entry/reports.go
index bda0a154eb9521941465cd715efc89dd779233ec..8ec42a61235dea977e9ab5e4c30fa784d5db62da
--- a/packages/app/internal/entry/reports.go
+++ b/packages/app/internal/entry/reports.go
@@ -23,7 +23,12 @@
 	rows := make([]doctorRow, 0, len(report.Checks))
 	var plain strings.Builder
 	for _, check := range report.Checks {
-		rows = append(rows, doctorRow{check.Name, check.Module, check.Passed, check.Detail})
+		name, scope, _ := strings.Cut(check.Name, " ")
+		detail := check.Detail
+		if scope != "" {
+			detail = scope + ": " + detail
+		}
+		rows = append(rows, doctorRow{name, check.Module, check.Passed, detail})
 		outcome := "passed"
 		if !check.Passed {
 			outcome = "failed"
@@ -149,7 +154,7 @@
 func status(root string) response {
 	report, err := workflow.Status(root)
 	if err != nil {
-		return operationFailure(err, false)
+		return operationFailure(root, err, false)
 	}
 	modules, moduleText := moduleRows(report.Modules)
 	data := struct {
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
p=Path('packages/app/internal/entry/lifecycle.go')
s=p.read_text().replace('func operationFailure(err error, lookupIsUsage bool) response {\\n\\tcode, _ := cause(err, \"P01\")','func operationFailure(root string, err error, lookupIsUsage bool) response {\\n\\tcode := projectCause(root, err, \"P01\")')
s=s.replace('operationFailure(err,', 'operationFailure(root, err,')
s=s.replace('func drill(ctx context.Context, snap project.Snapshot, module, bin string) response {','func drill(ctx context.Context, snap project.Snapshot, module, bin string) response {\\n\\troot := snap.Root')
s=s.replace('if report "'!= "" {'"\\n\\t\\tif _, e := os.Stat(filepath.Join(snap.Root, report)); e "'!= nil {'"\\n\\t\\t\\treport = \"\"\\n\\t\\t}\\n\\t}', 'if report "'!= "" && !currentDrillReport(root, report, result.Date) {'"\\n\\t\\treport = \"\"\\n\\t}')
s += '''
func currentDrillReport(root, report, date string) bool {
	when, err := time.Parse(\"20060102T150405.000000000\", date)
	if err "'!= nil {
		return false
	}
	body, err := os.ReadFile(filepath.Join(root, report))
	if err != nil {
		return false
	}
	return strings.HasPrefix(string(body), "---'"\\\\ndate: \"+when.UTC().Format(time.RFC3339Nano)+\"\\\\n\")
}
'''
p.write_text(s)
PY
gofmt -w packages/app/internal/entry
export PATH=\""'$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
go test ./packages/app/evals -run '"'Test_APP_(147|148|037|062|116|089|145|126|141|142)'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc "rg -n 'date|report|stdout|range' packages/runner/SPEC.md | tail -40; rg -n 'currentDrillReport|operationFailure|projectCause' packages/app/internal/entry; git status --short; wc -l packages/app/internal/entry/*" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
exec
/bin/zsh -lc "sed -n '1,115p' packages/checks/internal/engine/boundaries.go; rg -n 'func ValidateGraph' -A50 packages/project/internal/fixture/fixture.go; sed -n '1,58p' packages/app/evals/saved_attempt_cli_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 555ms:
package engine

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks/model"
	"github.com/lurkshark/fireweed/packages/project"
)

func boundaryReason(module project.Module) string {
	reason, _ := module.Document.Frontmatter["boundaries_opt_out"].(string)
	return strings.TrimSpace(reason)
}

func Boundaries(ctx context.Context, request model.BoundaryRequest) ([]project.Diagnostic, error) {
	snapshot := request.Snapshot
	baseline := map[string]bool{}
	data, err := os.ReadFile(filepath.Join(snapshot.Root, ".fireweed", "boundaries-baseline"))
	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("read boundaries baseline: %w", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			baseline[line] = true
		}
	}
	var diagnostics []project.Diagnostic
	var violations int
	global := mapValue(snapshot.Config, "hooks")["boundaries"]
	configured := global != nil
	for _, module := range snapshot.Modules {
		moduleHook := mapValue(module.Document.Frontmatter, "hooks")["boundaries"]
		if moduleHook != nil {
			configured = true
		}
		if global == nil && moduleHook == nil && len(module.Zones.Disposable) > 0 && boundaryReason(module) == "" {
			violations++
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X02", Message: module.Name + ": missing boundaries hook"})
		}
		if reason := boundaryReason(module); reason != "" {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X03", Message: module.Name + ": boundaries opt-out: " + reason, Warning: true})
		}
	}
	if !configured {
		if violations == 0 {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X04", Message: "no boundary enforcement required; boundaries hook absent", Warning: true})
			return diagnostics, nil
		}
		return diagnostics, fmt.Errorf("missing required boundaries hook")
	}
	commands := []struct {
		command any
		module  project.Module
	}{}
	if global != nil {
		commands = append(commands, struct {
			command any
			module  project.Module
		}{command: global})
	}
	for _, module := range snapshot.Modules {
		if command := mapValue(module.Document.Frontmatter, "hooks")["boundaries"]; command != nil {
			commands = append(commands, struct {
				command any
				module  project.Module
			}{command: command, module: module})
		}
	}
	var hookErrors []error
	for _, hook := range commands {
		values := map[string]string{"module": hook.module.Name, "module_dir": hook.module.Dir}
		result, runErr := runHook(ctx, snapshot.Root, hook.command, values, request.Bin, request.Timeout)
		if runErr != nil {
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X02", Message: "boundary hook: " + runErr.Error()})
			hookErrors = append(hookErrors, fmt.Errorf("boundary hook: %w", runErr))
			break
		}
		attributed := false
		for _, output := range []string{result.Stdout, result.Stderr} {
			for _, line := range strings.Split(output, "\n") {
				line = strings.TrimSuffix(line, "\r")
				for _, module := range snapshot.Modules {
					if strings.HasPrefix(line, module.Name+":") {
						attributed = true
						if !baseline[line] && boundaryReason(module) == "" {
							violations++
							diagnostics = append(diagnostics, project.Diagnostic{Code: "X01", Message: line})
						}
						break
					}
				}
			}
		}
		if result.ExitCode != 0 && !attributed {
			name := "global boundary hook"
			if hook.module.Name != "" {
				name = hook.module.Name + " boundary hook"
			}
			err := fmt.Errorf("%s: exit status %d without a module-attributed violation", name, result.ExitCode)
			diagnostics = append(diagnostics, project.Diagnostic{Code: "X02", Message: err.Error()})
			hookErrors = append(hookErrors, err)
		}
	}
	if len(hookErrors) == 1 {
		return diagnostics, hookErrors[0]
	}
	if len(hookErrors) > 1 {
		return diagnostics, errors.Join(hookErrors...)
	}
	if violations > 0 {
217:func ValidateGraph(modules []model.Module) error {
218-	byName := make(map[string]model.Module, len(modules))
219-	prefixes := make(map[string]string, len(modules))
220-	for _, module := range modules {
221-		if module.Name == "" {
222-			return fmt.Errorf("empty module name")
223-		}
224-		if module.IDPrefix == "" {
225-			return fmt.Errorf("empty ID prefix for module %s", module.Name)
226-		}
227-		if _, exists := byName[module.Name]; exists {
228-			return fmt.Errorf("duplicate module name %q", module.Name)
229-		}
230-		if other, exists := prefixes[module.IDPrefix]; exists {
231-			return fmt.Errorf("duplicate ID prefix %q for %s and %s", module.IDPrefix, other, module.Name)
232-		}
233-		byName[module.Name] = module
234-		prefixes[module.IDPrefix] = module.Name
235-	}
236-	state := make(map[string]uint8, len(modules))
237-	var visit func(string) error
238-	visit = func(name string) error {
239-		state[name] = 1
240-		for _, dependency := range byName[name].DependsOn {
241-			if dependency == name {
242-				return fmt.Errorf("module %s has a self-dependency", name)
243-			}
244-			if _, exists := byName[dependency]; !exists {
245-				return fmt.Errorf("module %s depends on missing module %s", name, dependency)
246-			}
247-			if state[dependency] == 1 {
248-				return fmt.Errorf("dependency cycle between %s and %s", name, dependency)
249-			}
250-			if state[dependency] == 0 {
251-				if err := visit(dependency); err != nil {
252-					return err
253-				}
254-			}
255-		}
256-		state[name] = 2
257-		return nil
258-	}
259-	for _, module := range modules {
260-		if state[module.Name] == 0 {
261-			if err := visit(module.Name); err != nil {
262-				return err
263-			}
264-		}
265-	}
266-	return nil
267-}
package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/app"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// The public app entry point discovers its project from the process directory.
// These tests do not run in parallel because this helper changes that directory.
func runSavedAttemptApp(t *testing.T, dir string, args ...string) (int, string, string) {
	t.Helper()
	previous, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chdir(previous); err != nil {
			t.Fatal(err)
		}
	}()
	var out, errOut bytes.Buffer
	status := app.Run(args, &out, &errOut)
	return status, out.String(), errOut.String()
}

func savedAttemptAppFixture(t *testing.T, attempts int) (string, workflow.State) {
	t.Helper()
	root := lifecycleCLIProject(t, "fixture-start-does-not-run-hooks", false)
	cliGit(t, root, "config", "user.name", "Fixture")
	cliGit(t, root, "config", "user.email", "fixture@example.invalid")
	appPut(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	appPut(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	cliCommit(t, root, "add baseline eval")

	started, err := workflow.Start(root, "manual")
	if err != nil || started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 {
		t.Fatalf("fixture Start: state=%#v err=%v", started, err)
	}
	for n := 1; n <= attempts; n++ {
		appPut(t, root, "packages/alpha/evals/current_test.go", "saved attempt "+string(rune('0'+n))+"\n")
		appPut(t, root, "packages/alpha/evals/new_test.go", "new attempt "+string(rune('0'+n))+"\n")
		if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
		blocked, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "eval attempts exhausted"})
		if err != nil || blocked.AttemptRef != "refs/fireweed/attempts/manual/"+string(rune('0'+n)) || blocked.AttemptOID == "" {
			t.Fatalf("fixture Block %d: result=%#v err=%v", n, blocked, err)
		}

 succeeded in 591ms:
36:InvokeAgent MUST retain the command's stdout, stderr and exit code in AgentResult.
172:that combined failure so the caller cannot report it as a clean stop.
201:Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.
204:IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.
207:On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to RUN-044.
208:It MUST select cases by lexical module name and then report order, excluding passing and skipped cases.
244:Drill MUST invoke reconstruction in a separate input workspace holding the invoking HEAD's tracked files except the target's disposable files, every direct and transitive dependent module's files, the changes directory and drill reports. Its Git metadata MUST hold one fresh commit and none of the invoking repository's history, refs or saved attempts. Drill MUST run the configured setup hook there before agent invocation.
259:After accepted reconstruction, Drill MUST run the target module's eval hook and require passing enforced cases and coverage through checks. It MUST retain the complete returned own report in Own, including failed testcase details when verification fails.
262:Drill MUST evaluate every direct and transitive dependent in lexical module-name order regardless of verify_dependents, retaining each report in Dependents. Every dependent MUST satisfy passing enforced cases and coverage for a passing drill, even when own or another dependent's evals fail.
266:After reconstruction has been invoked, Drill MUST write the current invocation's report to .fireweed/drills/<module>.md, replacing any previous report there, with date, base_commit, agent, model and result frontmatter and Own evals, Dependent evals, Contradicted Decisions and Gaps sections. Eval entries MUST identify module, testcase, spec IDs and result. Agent observations and verification failures MUST remain visible, with each gap identifying a nearest spec ID or "no spec" and target module.
269:After accepted reconstruction, Drill MUST retain the rebuilt verification tree at refs/fireweed/drills/<module>/<date>, returning that ref without moving the invoking HEAD. A passing result MUST require accepted reconstruction and passing own and all dependent evals; failure MUST return a nonnil error with Result fail.
272:Drill MUST remove its verification worktree and input workspace on success, agent or verification failure, invalid edits, usage stop and cancellation. Cleanup failure MUST be returned as an error rather than reported as a passing drill.
275:InvokeAgent MUST persist the invocation's complete stdout and stderr, including partial output on command failure, timeout or cancellation, at .fireweed/runs/<run>/<phase>-<attempt>.log under its log project root. Repeated invocations at the same path MUST append without discarding earlier records. AgentResult.LogPath MUST identify that file by an absolute path.
292:case-insensitive literal substrings, using the final 20 lines of stdout followed
295:- agent-invocation-logs: Agent logs append length-labeled stdout and stderr records after each invocation, preserving both streams' bytes even without final newlines. A log write error is returned alongside any agent error while the captured result remains available.
packages/app/internal/entry/project.go:108:			return reject(1, projectCause(root, err, "P01"), err)
packages/app/internal/entry/project.go:115:				return reject(1, projectCause(root, err, "P01"), err)
packages/app/internal/entry/project.go:119:					return reject(1, projectCause(root, err, "P01"), err)
packages/app/internal/entry/project.go:159:			return reject(2, projectCause(root, err, "C02"), err)
packages/app/internal/entry/project.go:241:				return reject(2, projectCause(root, err, "C02"), err)
packages/app/internal/entry/project.go:295:		return reject(1, projectCause(root, err, "P01"), err)
packages/app/internal/entry/project.go:355:		code := projectCause(root, err, "P01")
packages/app/internal/entry/lifecycle.go:25:func operationFailure(root string, err error, lookupIsUsage bool) response {
packages/app/internal/entry/lifecycle.go:26:	code := projectCause(root, err, "P01")
packages/app/internal/entry/lifecycle.go:48:			r = operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:52:				r = operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:62:			return operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:73:			r := operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:85:			return operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:91:			return operationFailure(root, err, true)
packages/app/internal/entry/lifecycle.go:97:			return operationFailure(root, err, true)
packages/app/internal/entry/lifecycle.go:103:			return operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:145:			r := operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:182:			r := operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:212:	if report != "" && !currentDrillReport(root, report, result.Date) {
packages/app/internal/entry/lifecycle.go:238:		return operationFailure(root, err, false)
packages/app/internal/entry/lifecycle.go:243:func currentDrillReport(root, report, date string) bool {
packages/app/internal/entry/entry.go:282:func projectCause(root string, err error, fallback string) string {
packages/app/internal/entry/reports.go:157:		return operationFailure(root, err, false)
?? packages/app/internal/
     398 packages/app/internal/entry/entry.go
     253 packages/app/internal/entry/lifecycle.go
     407 packages/app/internal/entry/project.go
     216 packages/app/internal/entry/reports.go
     699 packages/app/internal/entry/setup.go
    1973 total

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..94c057b47b68b3912e396e91f2a7dfdc4e9fb31f
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,10 @@
+## Contradicted Decisions
+
+None found.
+
+## Gaps
+
+- **APP-004 and APP-012, app.** The retained workflow API resolves a claim through the live queue. If the pending document or module configuration can no longer be parsed, the app cannot get the approved snapshot through that API. The specs do not define how to combine approved metadata with a projection requested for a different change.
+- **APP-037, APP-078 and APP-085, app.** Some project errors carry their codes privately. The CLI recovers those codes by matching the cause against public `project.Check` diagnostics. The retained interfaces do not provide a general way to read a code from those errors after wrapping.
+- **APP-067, app.** The retained eval expects short Doctor probe names such as `baseline` and `recovery`. The checks API includes the hook scope in the name. JSON uses the short name and keeps the scope in the detail.
+- **APP-103, app.** The specs leave the accepted module name and prefix grammar implicit. The implementation uses lowercase letter-led names with digits and hyphenated segments, and uppercase letter-led prefixes with optional digits.
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..35ae94741427288cb4c27a7ee15e4ff16e049ec7
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..1c60a9b418a85152ffbf98e96c7b445ba2976756
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,14 +105,20 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
 			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
+			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
 	case "migrate":
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -149,6 +148,7 @@
 }
 
 func show(snap project.Snapshot, module, id string, projected bool) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -156,7 +156,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,6 +227,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -237,7 +238,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -291,7 +292,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +352,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/reports.go b/packages/app/internal/entry/reports.go
index bda0a154eb9521941465cd715efc89dd779233ec..8ec42a61235dea977e9ab5e4c30fa784d5db62da
--- a/packages/app/internal/entry/reports.go
+++ b/packages/app/internal/entry/reports.go
@@ -23,7 +23,12 @@
 	rows := make([]doctorRow, 0, len(report.Checks))
 	var plain strings.Builder
 	for _, check := range report.Checks {
-		rows = append(rows, doctorRow{check.Name, check.Module, check.Passed, check.Detail})
+		name, scope, _ := strings.Cut(check.Name, " ")
+		detail := check.Detail
+		if scope != "" {
+			detail = scope + ": " + detail
+		}
+		rows = append(rows, doctorRow{name, check.Module, check.Passed, detail})
 		outcome := "passed"
 		if !check.Passed {
 			outcome = "failed"
@@ -149,7 +154,7 @@
 func status(root string) response {
 	report, err := workflow.Status(root)
 	if err != nil {
-		return operationFailure(err, false)
+		return operationFailure(root, err, false)
 	}
 	modules, moduleText := moduleRows(report.Modules)
 	data := struct {
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

 succeeded in 73165ms:
ok  	github.com/lurkshark/fireweed/packages/app/evals	71.191s

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/entry.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..94c057b47b68b3912e396e91f2a7dfdc4e9fb31f
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,10 @@
+## Contradicted Decisions
+
+None found.
+
+## Gaps
+
+- **APP-004 and APP-012, app.** The retained workflow API resolves a claim through the live queue. If the pending document or module configuration can no longer be parsed, the app cannot get the approved snapshot through that API. The specs do not define how to combine approved metadata with a projection requested for a different change.
+- **APP-037, APP-078 and APP-085, app.** Some project errors carry their codes privately. The CLI recovers those codes by matching the cause against public `project.Check` diagnostics. The retained interfaces do not provide a general way to read a code from those errors after wrapping.
+- **APP-067, app.** The retained eval expects short Doctor probe names such as `baseline` and `recovery`. The checks API includes the hook scope in the name. JSON uses the short name and keeps the scope in the detail.
+- **APP-103, app.** The specs leave the accepted module name and prefix grammar implicit. The implementation uses lowercase letter-led names with digits and hyphenated segments, and uppercase letter-led prefixes with optional digits.
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..160e253ba121d83bd2144687b5ad1a5f52633c02
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message || strings.HasSuffix(message, ": "+d.Message) ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..1c60a9b418a85152ffbf98e96c7b445ba2976756
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,14 +105,20 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
 			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
+			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
 	case "migrate":
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -149,6 +148,7 @@
 }
 
 func show(snap project.Snapshot, module, id string, projected bool) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -156,7 +156,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,6 +227,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -237,7 +238,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -291,7 +292,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +352,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/reports.go b/packages/app/internal/entry/reports.go
index bda0a154eb9521941465cd715efc89dd779233ec..8ec42a61235dea977e9ab5e4c30fa784d5db62da
--- a/packages/app/internal/entry/reports.go
+++ b/packages/app/internal/entry/reports.go
@@ -23,7 +23,12 @@
 	rows := make([]doctorRow, 0, len(report.Checks))
 	var plain strings.Builder
 	for _, check := range report.Checks {
-		rows = append(rows, doctorRow{check.Name, check.Module, check.Passed, check.Detail})
+		name, scope, _ := strings.Cut(check.Name, " ")
+		detail := check.Detail
+		if scope != "" {
+			detail = scope + ": " + detail
+		}
+		rows = append(rows, doctorRow{name, check.Module, check.Passed, detail})
 		outcome := "passed"
 		if !check.Passed {
 			outcome = "failed"
@@ -149,7 +154,7 @@
 func status(root string) response {
 	report, err := workflow.Status(root)
 	if err != nil {
-		return operationFailure(err, false)
+		return operationFailure(root, err, false)
 	}
 	modules, moduleText := moduleRows(report.Modules)
 	data := struct {
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/project.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..94c057b47b68b3912e396e91f2a7dfdc4e9fb31f
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,10 @@
+## Contradicted Decisions
+
+None found.
+
+## Gaps
+
+- **APP-004 and APP-012, app.** The retained workflow API resolves a claim through the live queue. If the pending document or module configuration can no longer be parsed, the app cannot get the approved snapshot through that API. The specs do not define how to combine approved metadata with a projection requested for a different change.
+- **APP-037, APP-078 and APP-085, app.** Some project errors carry their codes privately. The CLI recovers those codes by matching the cause against public `project.Check` diagnostics. The retained interfaces do not provide a general way to read a code from those errors after wrapping.
+- **APP-067, app.** The retained eval expects short Doctor probe names such as `baseline` and `recovery`. The checks API includes the hook scope in the name. JSON uses the short name and keeps the scope in the detail.
+- **APP-103, app.** The specs leave the accepted module name and prefix grammar implicit. The implementation uses lowercase letter-led names with digits and hyphenated segments, and uppercase letter-led prefixes with optional digits.
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..160e253ba121d83bd2144687b5ad1a5f52633c02
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message || strings.HasSuffix(message, ": "+d.Message) ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..9bea8df8cd95b9bc99e7437579ab58a90fb68e7d
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,14 +105,20 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
 			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
+			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
 	case "migrate":
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -149,6 +148,7 @@
 }
 
 func show(snap project.Snapshot, module, id string, projected bool) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -156,7 +156,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,9 +227,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
-	if err := project.ValidateGraph(snap.Modules); err != nil {
-		return reject(2, "C02", err)
-	}
+	root := snap.Root
 	if through != "" {
 		if state, ok := held(snap.Root, through); ok {
 			snap = state.Snapshot
@@ -237,7 +235,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -252,6 +250,9 @@
 			}
 		}
 	}
+	if err := project.ValidateGraph(snap.Modules); err != nil {
+		return reject(2, "C02", err)
+	}
 	m, err := findModule(snap, module)
 	if err != nil {
 		return reject(2, "U02", err)
@@ -291,7 +292,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +352,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/reports.go b/packages/app/internal/entry/reports.go
index bda0a154eb9521941465cd715efc89dd779233ec..8ec42a61235dea977e9ab5e4c30fa784d5db62da
--- a/packages/app/internal/entry/reports.go
+++ b/packages/app/internal/entry/reports.go
@@ -23,7 +23,12 @@
 	rows := make([]doctorRow, 0, len(report.Checks))
 	var plain strings.Builder
 	for _, check := range report.Checks {
-		rows = append(rows, doctorRow{check.Name, check.Module, check.Passed, check.Detail})
+		name, scope, _ := strings.Cut(check.Name, " ")
+		detail := check.Detail
+		if scope != "" {
+			detail = scope + ": " + detail
+		}
+		rows = append(rows, doctorRow{name, check.Module, check.Passed, detail})
 		outcome := "passed"
 		if !check.Passed {
 			outcome = "failed"
@@ -149,7 +154,7 @@
 func status(root string) response {
 	report, err := workflow.Status(root)
 	if err != nil {
-		return operationFailure(err, false)
+		return operationFailure(root, err, false)
 	}
 	modules, moduleText := moduleRows(report.Modules)
 	data := struct {
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..78f3ef445c10f206132c0a0c89e75abd63d55531
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -569,6 +569,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +672,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/packages/app/internal/entry/setup.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..94c057b47b68b3912e396e91f2a7dfdc4e9fb31f
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,10 @@
+## Contradicted Decisions
+
+None found.
+
+## Gaps
+
+- **APP-004 and APP-012, app.** The retained workflow API resolves a claim through the live queue. If the pending document or module configuration can no longer be parsed, the app cannot get the approved snapshot through that API. The specs do not define how to combine approved metadata with a projection requested for a different change.
+- **APP-037, APP-078 and APP-085, app.** Some project errors carry their codes privately. The CLI recovers those codes by matching the cause against public `project.Check` diagnostics. The retained interfaces do not provide a general way to read a code from those errors after wrapping.
+- **APP-067, app.** The retained eval expects short Doctor probe names such as `baseline` and `recovery`. The checks API includes the hook scope in the name. JSON uses the short name and keeps the scope in the detail.
+- **APP-103, app.** The specs leave the accepted module name and prefix grammar implicit. The implementation uses lowercase letter-led names with digits and hyphenated segments, and uppercase letter-led prefixes with optional digits.
diff --git a/packages/app/internal/entry/entry.go b/packages/app/internal/entry/entry.go
index 2c8fc2ffffa658f455bea344e10b2c51767373b1..160e253ba121d83bd2144687b5ad1a5f52633c02
--- a/packages/app/internal/entry/entry.go
+++ b/packages/app/internal/entry/entry.go
@@ -89,6 +89,9 @@
 	if slices.Contains(args, "--fix") {
 		usage = "fireweed check --fix"
 	}
+	if slices.Contains(args, "--for") {
+		usage = "fireweed run [--max N] [--for D] [--allow-unmerged]"
+	}
 	bad := func(message string, params ...any) (command, error) {
 		msg := fmt.Sprintf(message, params...)
 		if c.name == "init" {
@@ -135,7 +138,7 @@
 		c.flags[arg] = value
 	}
 	if len(c.args) < s.min || len(c.args) > s.max {
-		return bad("invalid arguments for %s", c.name)
+		return bad("invalid arguments for %s: %q", c.name, c.args)
 	}
 	if c.has("--stale") && (len(c.args) != 1 || c.has("--history") || c.has("--projected")) {
 		return bad("--stale requires one change and excludes --history and --projected")
@@ -157,6 +160,9 @@
 			return bad("invalid ID prefix %q", c.flags["--prefix"])
 		}
 	}
+	if c.name == "new-change" && !changeSlug.MatchString(c.arg(1)) {
+		return bad("invalid change slug %q", c.arg(1))
+	}
 	return c, nil
 }
 
@@ -200,17 +206,39 @@
 	return slices.ContainsFunc(ds, func(d project.Diagnostic) bool { return !d.Warning })
 }
 
-var causePattern = regexp.MustCompile(`(?:^|[\s:])([A-Z][0-9]{2}|P2-01):\s+`)
+var changeSlug = regexp.MustCompile(`^[a-z0-9-]+$`)
+var causePattern = regexp.MustCompile(`^([A-Z][0-9]{2}|P2-01):\s+`)
 
 func cause(err error, fallback string) (string, string) {
+	code := assignedCause(err)
 	message := err.Error()
-	if code, _, ok := contracts.ErrorCode(err); ok {
-		return code, strings.TrimPrefix(message, code+": ")
+	if code == "" {
+		return fallback, message
 	}
-	if match := causePattern.FindStringSubmatch(message); match != nil {
-		return match[1], strings.TrimPrefix(message, match[1]+": ")
+	return code, strings.TrimPrefix(message, code+": ")
+}
+
+func assignedCause(err error) string {
+	if match := causePattern.FindStringSubmatch(err.Error()); match != nil {
+		return match[1]
 	}
-	return fallback, message
+	switch e := err.(type) {
+	case contracts.CodedError:
+		return e.Code
+	case *contracts.CodedError:
+		return e.Code
+	case interface{ Unwrap() []error }:
+		for _, child := range e.Unwrap() {
+			if code := assignedCause(child); code != "" {
+				return code
+			}
+		}
+	case interface{ Unwrap() error }:
+		if child := e.Unwrap(); child != nil {
+			return assignedCause(child)
+		}
+	}
+	return ""
 }
 
 func reject(exit int, fallback string, err error) response {
@@ -251,6 +279,21 @@
 	return reject(2, "C01", err)
 }
 
+func projectCause(root string, err error, fallback string) string {
+	if code := assignedCause(err); code != "" {
+		return code
+	}
+	// Some project query errors retain their code only in Check's diagnostics.
+	message := err.Error()
+	for _, d := range project.Check(root) {
+		if d.Message == message || strings.HasSuffix(message, ": "+d.Message) ||
+			(d.Path != "" && strings.Contains(message, d.Path) && strings.Contains(message, d.Message)) {
+			return d.Code
+		}
+	}
+	return fallback
+}
+
 func Run(args []string, out, errOut io.Writer) int {
 	c, err := parse(args)
 	var r response
@@ -313,10 +356,18 @@
 	if c.name == "pr-body" {
 		return prBody(root, c.arg(0))
 	}
+	if c.name == "show" && c.has("--history") {
+		return specHistory(root, c.arg(0), c.arg(1))
+	}
 
 	snap, err := project.Load(root)
 	if err != nil {
 		failure := loadFailure(root, err)
+		if c.name == "new-module" && failure.diagnostics[0].Code != "C01" {
+			if conflict := occupiedModuleDestination(root, c.arg(0), failure.diagnostics[0].Path); conflict != nil {
+				return reject(2, "P01", conflict)
+			}
+		}
 		if (c.name == "check" || c.name == "doctor" || c.name == "ci") && failure.diagnostics[0].Code != "C01" {
 			failure.exit = 1
 		}
diff --git a/packages/app/internal/entry/project.go b/packages/app/internal/entry/project.go
index 00080493e7d09b324560bd12c1999663b06e0e31..9bea8df8cd95b9bc99e7437579ab58a90fb68e7d
--- a/packages/app/internal/entry/project.go
+++ b/packages/app/internal/entry/project.go
@@ -61,15 +61,11 @@
 
 func held(root, selector string) (workflow.State, bool) {
 	// Worktree claims are Git refs. A project can also be used before Git init.
-	for dir := root; ; dir = filepath.Dir(dir) {
-		if _, err := os.Lstat(filepath.Join(dir, ".git")); err == nil {
-			state, err := workflow.Current(root, selector)
-			return state, err == nil
-		}
-		if filepath.Dir(dir) == dir {
-			return workflow.State{}, false
-		}
+	if _, err := os.Lstat(filepath.Join(root, ".git")); err != nil {
+		return workflow.State{}, false
 	}
+	state, err := workflow.Current(root, selector)
+	return state, err == nil
 }
 
 func checkResult(ds []project.Diagnostic, text string) response {
@@ -99,8 +95,6 @@
 		if err != nil {
 			code := "P01"
 			switch message := err.Error(); {
-			case strings.HasPrefix(message, "invalid change slug"):
-				code = "U01"
 			case strings.HasPrefix(message, "unknown module"):
 				code = "U02"
 			case strings.HasPrefix(message, "invalid change number") || strings.HasPrefix(message, "change slug"):
@@ -111,14 +105,20 @@
 		return response{data: path, plain: path + "\n"}
 	case "stamp":
 		if err := project.Stamp(root, arg); err != nil {
-			return reject(1, "P01", err)
+			return reject(1, projectCause(root, err, "P01"), err)
 		}
 		return confirmation("Stamped " + arg + ".")
 	case "check":
 		if c.has("--fix") {
-			if err := project.StampMissing(root, ""); err != nil {
-				return reject(1, "P01", err)
+			changes, err := project.Queue(root)
+			if err != nil {
+				return reject(1, projectCause(root, err, "P01"), err)
 			}
+			for _, change := range changes {
+				if err := project.StampMissing(root, change.Path); err != nil {
+					return reject(1, projectCause(root, err, "P01"), err)
+				}
+			}
 		}
 		return checkResult(project.Check(root), "Check passed.")
 	case "migrate":
@@ -127,14 +127,13 @@
 		}
 		return checkResult(project.Check(root), "Project format 1 is already current.")
 	case "show":
-		switch {
-		case c.has("--history"):
-			return specHistory(root, arg, c.arg(1))
-		case c.has("--stale"):
+		if c.has("--stale") {
+			if err := project.ValidateGraph(snap.Modules); err != nil {
+				return reject(2, "C02", err)
+			}
 			return stale(root, arg)
-		default:
-			return show(snap, arg, c.arg(1), c.has("--projected"))
 		}
+		return show(snap, arg, c.arg(1), c.has("--projected"))
 	case "coverage":
 		return coverage(ctx, snap, arg, c.flags["--with-change"], bin)
 	case "list":
@@ -149,6 +148,7 @@
 }
 
 func show(snap project.Snapshot, module, id string, projected bool) response {
+	root := snap.Root
 	if err := project.ValidateGraph(snap.Modules); err != nil {
 		return reject(2, "C02", err)
 	}
@@ -156,7 +156,7 @@
 		var err error
 		snap, err = project.Project(snap.Root, "")
 		if err != nil {
-			return reject(2, "C02", err)
+			return reject(2, projectCause(root, err, "C02"), err)
 		}
 	}
 	m, err := findModule(snap, module)
@@ -227,9 +227,7 @@
 }
 
 func coverage(ctx context.Context, snap project.Snapshot, module, through, bin string) response {
-	if err := project.ValidateGraph(snap.Modules); err != nil {
-		return reject(2, "C02", err)
-	}
+	root := snap.Root
 	if through != "" {
 		if state, ok := held(snap.Root, through); ok {
 			snap = state.Snapshot
@@ -237,7 +235,7 @@
 			var err error
 			snap, err = project.Project(snap.Root, through)
 			if err != nil {
-				return reject(2, "C02", err)
+				return reject(2, projectCause(root, err, "C02"), err)
 			}
 		}
 	} else if state, ok := held(snap.Root, ""); ok {
@@ -252,6 +250,9 @@
 			}
 		}
 	}
+	if err := project.ValidateGraph(snap.Modules); err != nil {
+		return reject(2, "C02", err)
+	}
 	m, err := findModule(snap, module)
 	if err != nil {
 		return reject(2, "U02", err)
@@ -291,7 +292,7 @@
 func list(root string) response {
 	changes, err := project.Queue(root)
 	if err != nil {
-		return reject(1, "P01", err)
+		return reject(1, projectCause(root, err, "P01"), err)
 	}
 	rows := make([]queueRow, 0, len(changes))
 	var plain strings.Builder
@@ -351,7 +352,7 @@
 func stale(root, selector string) response {
 	items, err := project.StaleSpecs(root, selector)
 	if err != nil {
-		code, _ := cause(err, "P01")
+		code := projectCause(root, err, "P01")
 		exit := 1
 		if code == "U02" || code == "C13" || code == "C01" {
 			exit = 2
diff --git a/packages/app/internal/entry/reports.go b/packages/app/internal/entry/reports.go
index bda0a154eb9521941465cd715efc89dd779233ec..8ec42a61235dea977e9ab5e4c30fa784d5db62da
--- a/packages/app/internal/entry/reports.go
+++ b/packages/app/internal/entry/reports.go
@@ -23,7 +23,12 @@
 	rows := make([]doctorRow, 0, len(report.Checks))
 	var plain strings.Builder
 	for _, check := range report.Checks {
-		rows = append(rows, doctorRow{check.Name, check.Module, check.Passed, check.Detail})
+		name, scope, _ := strings.Cut(check.Name, " ")
+		detail := check.Detail
+		if scope != "" {
+			detail = scope + ": " + detail
+		}
+		rows = append(rows, doctorRow{name, check.Module, check.Passed, detail})
 		outcome := "passed"
 		if !check.Passed {
 			outcome = "failed"
@@ -149,7 +154,7 @@
 func status(root string) response {
 	report, err := workflow.Status(root)
 	if err != nil {
-		return operationFailure(err, false)
+		return operationFailure(root, err, false)
 	}
 	modules, moduleText := moduleRows(report.Modules)
 	data := struct {
diff --git a/packages/app/internal/entry/setup.go b/packages/app/internal/entry/setup.go
index ab158da7a955d3d0072ea648d1b6002d1af9731d..969f6d562ac9a9991891131540f29c520d25283e
--- a/packages/app/internal/entry/setup.go
+++ b/packages/app/internal/entry/setup.go
@@ -202,10 +202,13 @@
 		if err != nil {
 			return reject(2, "P01", err)
 		}
-		generated, _, err := renderModule(templates, "app", "APP", "packages/app", goModule)
+		generated, module, err := renderModule(templates, "app", "APP", "packages/app", goModule)
 		if err != nil {
 			return reject(2, "C03", err)
 		}
+		if len(module.DependsOn) != 0 {
+			return reject(2, "C03", fmt.Errorf("the generated app module must have no dependencies"))
+		}
 		for path, body := range generated {
 			if _, exists := files[path]; exists {
 				return reject(2, "P01", fmt.Errorf("recipe payload conflicts with generated %s", path))
@@ -356,7 +359,14 @@
 
 func readTree(dir string) (map[string]string, error) {
 	files := map[string]string{}
-	err := filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
+	info, err := os.Lstat(dir)
+	if err != nil {
+		return nil, err
+	}
+	if !info.IsDir() {
+		return nil, fmt.Errorf("%s must be a directory", dir)
+	}
+	err = filepath.WalkDir(dir, func(path string, entry fs.DirEntry, err error) error {
 		if err != nil {
 			return err
 		}
@@ -569,6 +579,29 @@
 	return ""
 }
 
+// Load can reject an occupied destination without a SPEC.md before it
+// returns the module globs. Report that destination as a collision.
+func occupiedModuleDestination(root, name, failedPath string) error {
+	if !strings.HasSuffix(failedPath, "/SPEC.md") {
+		return nil
+	}
+	raw, err := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
+	if err != nil {
+		return nil
+	}
+	var cfg struct {
+		Modules []string `yaml:"modules"`
+	}
+	if err := yaml.Unmarshal(raw, &cfg); err != nil {
+		return nil
+	}
+	dir := moduleDestination(cfg.Modules, name)
+	if dir == "" || failedPath != dir+"/SPEC.md" {
+		return nil
+	}
+	return vacant(root, dir)
+}
+
 func renderModule(templates map[string]string, name, prefix, dir, goModule string) (map[string]string, project.Module, error) {
 	files := map[string]string{}
 	render := strings.NewReplacer("{{name}}", name, "{{prefix}}", prefix, "{{package}}", strings.ReplaceAll(name, "-", "_"), "{{module_dir}}", dir, "{{go_module}}", goModule)
@@ -649,8 +682,28 @@
 	if err = project.ValidateGraph(append(slices.Clone(snap.Modules), m)); err != nil {
 		return reject(2, "C03", fmt.Errorf("rendered module graph: %v", err))
 	}
-	if err = writeSetup(snap.Root, files, nil, dir); err != nil {
+	if err = writeModule(snap.Root, dir, files); err != nil {
 		return reject(2, "P01", err)
 	}
 	return response{data: dir, plain: dir + "\n"}
 }
+
+func writeModule(root, dir string, files map[string]string) (err error) {
+	var created []string
+	defer func() {
+		if err != nil {
+			for i := len(created) - 1; i >= 0; i-- {
+				err = errors.Join(err, os.Remove(created[i]))
+			}
+		}
+	}()
+	if err = ensureDirectory(root, filepath.Dir(dir), &created); err != nil {
+		return err
+	}
+	path := filepath.Join(root, dir)
+	if err = os.Mkdir(path, 0755); err != nil {
+		return err
+	}
+	created = append(created, path)
+	return writeSetup(root, files, nil)
+}

exec
/bin/zsh -lc 'gofmt -w packages/app/internal/entry; git diff --check; git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 573ms:
?? DRILL_NOTES.md
?? packages/app/internal/

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
mkdir -p .fireweed/runs/drill && sh -c '"'go run .fireweed/hooks.go eval \"packages/app\" \""'$1"'"' sh \""'$PWD/.fireweed/runs/drill/report.xml"' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
codex
The focused rerun passed, including the previously failing diagnostic, stamping and Doctor cases. The full app eval hook is running again.
exec
/bin/zsh -lc "sed -n '70,103p' packages/app/evals/authoring_commands_test.go; sed -n '65,87p' packages/app/evals/release_cli_test.go; sed -n '91,106p' packages/app/evals/abandon_cli_test.go; git diff --exit-code; git ls-files --others --exclude-standard; find packages/app/internal -type f -maxdepth 2 | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 515ms:
			authoringChange(t, root, "0001-revise.md", "Users need a revision.\n", "### new:local\nLocal MUST stay unassigned.\n", "### ALP-001\nAlpha MUST use revised text.\n")
			status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), "stamp", alias)
			if status != 0 || stderr != "" || strings.TrimSpace(stdout) == "" || !strings.Contains(stdout, "revise") {
				t.Errorf("stamp %q: status=%d stdout=%q stderr=%q; want confirmation", alias, status, stdout, stderr)
			}
			body, err := os.ReadFile(filepath.Join(root, "changes/0001-revise.md"))
			if err != nil {
				t.Fatal(err)
			}
			doc, err := contracts.ParseDocument(string(body))
			if err != nil {
				t.Fatal(err)
			}
			bases, ok := doc.Frontmatter["bases"].(map[string]any)
			want := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST keep its current text.\n"})
			if !ok || bases["ALP-001"] != want || len(bases) != 1 {
				t.Errorf("bases=%#v; want only ALP-001=%s", doc.Frontmatter["bases"], want)
			}
			if !strings.Contains(doc.Body, "### new:local\n") || strings.Contains(doc.Body, "### ALP-002\n") {
				t.Errorf("stamp assigned Added handle: %s", doc.Body)
			}
		})
	}
	root := authoringFixture(t)
	authoringChange(t, root, "0001-added.md", "Users need an addition.\n", "### new:only\nOnly MUST work.\n", "")
	status, stdout, stderr := runAppBinary(t, bin, appNestedDir(t, root), "stamp", "added", "--json")
	if status != 0 || stderr != "" || !strings.Contains(commandStringData(t, stdout), "added") {
		t.Errorf("Added-only JSON stamp: status=%d stdout=%q stderr=%q", status, stdout, stderr)
	}
	body, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), "### new:only\n") || strings.Contains(string(body), "### ALP-002\n") {
			if tc.json {
				args = append(args, "--json")
			}
			status, out, errOut := runSavedAttemptApp(t, appNestedDir(t, root), args...)
			if status != 0 || errOut != "" {
				t.Fatalf("release: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			confirmation := out
			if tc.json {
				confirmation = commandStringData(t, out)
			}
			if !strings.Contains(strings.ToLower(confirmation), "releas") || !strings.Contains(confirmation, "manual") {
				t.Errorf("missing release confirmation: %q", confirmation)
			}
			if got := cliGit(t, root, "rev-parse", "HEAD"); got != beforeHead {
				t.Errorf("release moved HEAD: %s -> %s", beforeHead, got)
			}
			if got := cliGit(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
				t.Errorf("release changed worktree status: %q -> %q", beforeStatus, got)
			}
			for path, want := range map[string]string{
				changePath:                           wantChange,
				"packages/alpha/internal/worker.txt": "unfinished worker work\n",
				t.Fatalf("abandon: exit=%d stdout=%q stderr=%q", status, out, errOut)
			}
			confirmation := out
			if tc.json {
				confirmation = commandStringData(t, out)
			}
			oid := cliGit(t, root, "rev-parse", "HEAD")
			if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) || oid == parent {
				t.Errorf("abandon HEAD=%q; want a new full object ID after %q", oid, parent)
			}
			if !strings.Contains(strings.ToLower(confirmation), "abandon") || !strings.Contains(confirmation, "manual") || !strings.Contains(confirmation, oid) {
				t.Errorf("confirmation %q must name abandon, selected change, and exact full HEAD %s", confirmation, oid)
			}
			if got := cliGit(t, root, "rev-parse", "HEAD^"); got != parent {
				t.Errorf("abandon commit parent=%s; want %s", got, parent)
			}
DRILL_NOTES.md
packages/app/internal/entry/entry.go
packages/app/internal/entry/lifecycle.go
packages/app/internal/entry/project.go
packages/app/internal/entry/reports.go
packages/app/internal/entry/setup.go
packages/app/internal/entry/entry.go
packages/app/internal/entry/lifecycle.go
packages/app/internal/entry/project.go
packages/app/internal/entry/reports.go
packages/app/internal/entry/setup.go

codex
The workspace diff is confined to the disposable app implementation and `DRILL_NOTES.md`. The notes record the approved-snapshot limitation and the CLI’s choices for Doctor probe names and module-name validation.
 succeeded in 264329ms:
?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/app/evals	260.711s

exec
/bin/zsh -lc "python3 - <<'PY'
import xml.etree.ElementTree as ET
report = ET.parse('.fireweed/runs/drill/report.xml')
print(f'{len(report.findall(\".//testcase\"))} cases, {len(report.findall(\".//failure\"))} failures, {len(report.findall(\".//skipped\"))} skipped')
PY
git diff --exit-code
git status --short
gofmt -l packages/app/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661
 succeeded in 726ms:
323 cases, 0 failures, 0 skipped
?? DRILL_NOTES.md
?? packages/app/internal/

codex
Rebuilt the Fireweed app CLI in `packages/app/internal/entry`. The target eval hook passed all **323 cases**. The changes are uncommitted. [Eval report](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/.fireweed/runs/drill/report.xml)

The controller’s `verify` returned P02 because the drill workspace has no held change.

[DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1023077661/DRILL_NOTES.md) records the approved-snapshot lookup limitation, diagnostic recovery, Doctor JSON naming, and module-name validation. It records no contradicted Decisions.
tokens used
238,455

