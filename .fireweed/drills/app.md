---
agent: opencode
base_commit: cd97c191e19627fa6e2764179bff134b8ebbd632
date: "2026-11-19T13:49:22.302703Z"
model: opencode/inkling-2-small-free#high
result: pass
---

## Own evals
- app / Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON / APP-097, APP-099 / pass
- app / Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/plain-queued / APP-097, APP-099 / pass
- app / Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned / APP-097, APP-099 / pass
- app / Test_APP_098_APP_099_AbandonFailures / APP-098, APP-099 / pass
- app / Test_APP_098_APP_099_AbandonFailures/missing / APP-098, APP-099 / pass
- app / Test_APP_098_APP_099_AbandonFailures/extra-json / APP-098, APP-099 / pass
- app / Test_APP_098_APP_099_AbandonFailures/unsupported-option / APP-098, APP-099 / pass
- app / Test_APP_098_APP_099_AbandonFailures/unknown-selector / APP-098, APP-099 / pass
- app / Test_APP_001_StandaloneBinary / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/outside-project/help / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/outside-project/version / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/outside-project/explain / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/outside-project/playbook / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/invalid-config/help / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/invalid-config/version / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/invalid-config/explain / APP-001 / pass
- app / Test_APP_001_StandaloneBinary/invalid-config/playbook / APP-001 / pass
- app / Test_APP_002_InformationJSON / APP-002 / pass
- app / Test_APP_002_InformationJSON/help / APP-002 / pass
- app / Test_APP_002_InformationJSON/version / APP-002 / pass
- app / Test_APP_002_InformationJSON/explain / APP-002 / pass
- app / Test_APP_002_InformationJSON/playbook / APP-002 / pass
- app / Test_APP_003_InvalidRequests / APP-003 / pass
- app / Test_APP_003_InvalidRequests/unknown-command / APP-003 / pass
- app / Test_APP_003_InvalidRequests/unknown-option / APP-003 / pass
- app / Test_APP_003_InvalidRequests/missing-argument / APP-003 / pass
- app / Test_APP_003_InvalidRequests/extra-argument / APP-003 / pass
- app / Test_APP_003_InvalidRequests/new-change-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/stamp-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/check-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/show-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/start-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/context-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/coverage-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/verify-arity / APP-003 / pass
- app / Test_APP_003_InvalidRequests/APP_013_unknown-explanation / APP-003, APP-013 / pass
- app / Test_APP_003_InvalidRequests/APP_013_unknown-playbook / APP-003, APP-013 / pass
- app / Test_APP_014_MaintenanceCommandsDispatch / APP-014 / pass
- app / Test_APP_014_MaintenanceCommandsDispatch/gc / APP-014 / pass
- app / Test_APP_014_MaintenanceCommandsDispatch/migrate / APP-014 / pass
- app / Test_APP_015_FailedJSON / APP-015 / pass
- app / Test_APP_001_ReadCommandsDoNotInvokeGitHooksOrAgents / APP-001 / pass
- app / Test_APP_006_APP_022_StampAliasesAndUnassignedHandles / APP-006, APP-022 / pass
- app / Test_APP_006_APP_022_StampAliasesAndUnassignedHandles/revise / APP-006, APP-022 / pass
- app / Test_APP_006_APP_022_StampAliasesAndUnassignedHandles/0001-revise.md / APP-006, APP-022 / pass
- app / Test_APP_006_APP_022_StampAliasesAndUnassignedHandles/changes/0001-revise.md / APP-006, APP-022 / pass
- app / Test_APP_021_APP_022_StampErrorExits / APP-021, APP-022 / pass
- app / Test_APP_021_APP_022_StampErrorExits/missing-argument / APP-021, APP-022 / pass
- app / Test_APP_021_APP_022_StampErrorExits/extra-argument / APP-021, APP-022 / pass
- app / Test_APP_021_APP_022_StampErrorExits/invalid-config / APP-021, APP-022 / pass
- app / Test_APP_021_APP_022_StampErrorExits/missing-change / APP-021, APP-022 / pass
- app / Test_APP_021_APP_022_StampErrorExits/absent-modified-ID / APP-021, APP-022 / pass
- app / Test_APP_007_APP_024_CheckValidAndAllDiagnostics / APP-007, APP-024 / pass
- app / Test_APP_023_APP_024_CheckUsageAndConfiguration / APP-023, APP-024 / pass
- app / Test_APP_023_APP_024_CheckUsageAndConfiguration/extra-argument / APP-023, APP-024 / pass
- app / Test_APP_023_APP_024_CheckUsageAndConfiguration/invalid-config / APP-023, APP-024 / pass
- app / Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity / APP-008, APP-025 / pass
- app / Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/current-full / APP-008, APP-025 / pass
- app / Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/current-item / APP-008, APP-025 / pass
- app / Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/projected-full / APP-008, APP-025 / pass
- app / Test_APP_008_APP_025_ShowCurrentProjectedAndJSONParity/projected-item / APP-008, APP-025 / pass
- app / Test_APP_026_ShowNamedErrors / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/unknown-module / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/unknown-item / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/ambiguous-item / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/ambiguous-module / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/missing-module / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/extra-argument / APP-026 / pass
- app / Test_APP_026_ShowNamedErrors/invalid-config / APP-026 / pass
- app / Test_APP_105_APP_108_CheckCommitsRangesAndRecordedAuthorWork / APP-105, APP-108 / pass
- app / Test_APP_105_APP_108_CheckCommitsSquashAndEmptyJSON / APP-105, APP-108 / pass
- app / Test_APP_106_APP_108_CheckCommitsReportsEveryHistoryRule / APP-106, APP-108 / pass
- app / Test_APP_107_APP_108_CheckCommitsFailures / APP-107, APP-108 / pass
- app / Test_APP_107_APP_108_CheckCommitsFailures/extra / APP-107, APP-108 / pass
- app / Test_APP_107_APP_108_CheckCommitsFailures/option / APP-107, APP-108 / pass
- app / Test_APP_107_APP_108_CheckCommitsFailures/bad-range / APP-107, APP-108 / pass
- app / Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases / APP-147, APP-148 / pass
- app / Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp / APP-147, APP-148 / pass
- app / Test_APP_148_APP_149_CheckFixRejectsBeforeMutation / APP-148, APP-149 / pass
- app / Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules / APP-109, APP-110, APP-112, APP-113 / pass
- app / Test_APP_110_APP_112_APP_113_CIEvalAndBoundaryFailures / APP-110, APP-112, APP-113 / pass
- app / Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration / APP-109, APP-112, APP-113 / pass
- app / Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip / APP-111, APP-112, APP-113 / pass
- app / Test_APP_094_APP_096_CommitsPendingAndArchivedOrder / APP-094, APP-096 / pass
- app / Test_APP_095_APP_096_CommitsFailures / APP-095, APP-096 / pass
- app / Test_APP_095_APP_096_CommitsFailures/missing / APP-095, APP-096 / pass
- app / Test_APP_095_APP_096_CommitsFailures/extra-json / APP-095, APP-096 / pass
- app / Test_APP_095_APP_096_CommitsFailures/unsupported-option / APP-095, APP-096 / pass
- app / Test_APP_095_APP_096_CommitsFailures/unknown-selector / APP-095, APP-096 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects / APP-163 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects/show_alpha_--history / APP-163 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects/check-commits_HEAD..HEAD_--json / APP-163 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects/start_manual_--json / APP-163 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects/new-change_alpha_new-work / APP-163 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects/unsupported-format / APP-163 / pass
- app / Test_APP_163_LiveCompatibilityBeforeCommandEffects/live-pin-for-history / APP-163 / pass
- app / Test_APP_165_StandaloneAndSetupPreconditions / APP-165 / pass
- app / Test_APP_164_RejectsIncompatibleApprovedConfiguration / APP-164 / pass
- app / Test_APP_164_RejectsIncompatibleApprovedConfiguration/modules / APP-164 / pass
- app / Test_APP_164_RejectsIncompatibleApprovedConfiguration/coverage / APP-164 / pass
- app / Test_APP_164_RejectsIncompatibleApprovedConfiguration/next / APP-164 / pass
- app / Test_APP_164_RejectsIncompatibleApprovedConfiguration/run / APP-164 / pass
- app / Test_APP_170_ManualVerificationRejectsIncompatibleApprovedConfiguration / APP-170 / pass
- app / Test_APP_170_ManualVerificationRejectsIncompatibleApprovedConfiguration/verify / APP-170 / pass
- app / Test_APP_170_ManualVerificationRejectsIncompatibleApprovedConfiguration/done / APP-170 / pass
- app / Test_APP_170_ManualVerificationReportsProtectedConfigEdit / APP-170 / pass
- app / Test_APP_170_ManualVerificationReportsProtectedConfigEdit/verify / APP-170 / pass
- app / Test_APP_170_ManualVerificationReportsProtectedConfigEdit/done / APP-170 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/check_--json / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/show_alpha / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/new-change_alpha_new-work / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run / APP-164 / pass
- app / Test_APP_037_StaticCauses / APP-037 / pass
- app / Test_APP_037_StaticCauses/configuration / APP-037 / pass
- app / Test_APP_037_StaticCauses/graph / APP-037 / pass
- app / Test_APP_037_StaticCauses/syntax / APP-037 / pass
- app / Test_APP_037_StaticCauses/normative / APP-037 / pass
- app / Test_APP_037_StaticCauses/identity / APP-037 / pass
- app / Test_APP_037_StaticCauses/tag / APP-037 / pass
- app / Test_APP_037_StaticCauses/structure / APP-037 / pass
- app / Test_APP_037_StaticCauses/base / APP-037 / pass
- app / Test_APP_037_StaticCauses/handle / APP-037 / pass
- app / Test_APP_037_StaticCauses/metadata / APP-037 / pass
- app / Test_APP_037_StaticCauses/queue-identity / APP-037 / pass
- app / Test_APP_037_APP_038_ShowStaticCauses / APP-037, APP-038 / pass
- app / Test_APP_037_APP_038_ShowStaticCauses/duplicate-graph-name / APP-037, APP-038 / pass
- app / Test_APP_037_APP_038_ShowStaticCauses/duplicate-fixed-spec-id / APP-037, APP-038 / pass
- app / Test_APP_038_LookupCauses / APP-038 / pass
- app / Test_APP_038_LookupCauses/author-module / APP-038 / pass
- app / Test_APP_038_LookupCauses/show-module / APP-038 / pass
- app / Test_APP_038_LookupCauses/show-item / APP-038 / pass
- app / Test_APP_038_LookupCauses/show-ambiguous-item / APP-038 / pass
- app / Test_APP_038_LookupCauses/coverage-module / APP-038 / pass
- app / Test_APP_038_LookupCauses/stamp-change / APP-038 / pass
- app / Test_APP_038_LookupCauses/coverage-change / APP-038 / pass
- app / Test_APP_039_BlockedNewChangeOutputPath / APP-039 / pass
- app / Test_APP_040_ForwardedCheckCode / APP-040 / pass
- app / Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI / APP-062, APP-063, APP-064, APP-065, APP-066, APP-067 / pass
- app / Test_APP_062_DoctorUsesCloserProject / APP-062 / pass
- app / Test_APP_065_APP_066_APP_067_DoctorFailures / APP-065, APP-066, APP-067 / pass
- app / Test_APP_065_APP_066_APP_067_DoctorFailures/extra / APP-065, APP-066, APP-067 / pass
- app / Test_APP_065_APP_066_APP_067_DoctorFailures/unknown-flag / APP-065, APP-066, APP-067 / pass
- app / Test_APP_065_APP_066_APP_067_DoctorFailures/wrong-command-flag / APP-065, APP-066, APP-067 / pass
- app / Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite / APP-068, APP-069, APP-070, APP-071, APP-072, APP-073 / pass
- app / Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite/no-block / APP-068, APP-069, APP-070, APP-071, APP-072, APP-073 / pass
- app / Test_APP_068_APP_069_APP_070_APP_071_APP_072_APP_073_AgentsMDWrite/missing-file / APP-068, APP-069, APP-070, APP-071, APP-072, APP-073 / pass
- app / Test_APP_136_APP_137_DrillPlainNearestProjectAndInvokingBin / APP-136, APP-137 / pass
- app / Test_APP_138_DrillJSONSuccess / APP-138 / pass
- app / Test_APP_139_DrillArgumentAndModuleRejectionsDoNotInvokeHooks / APP-139 / pass
- app / Test_APP_166_DrillInvokesSelectedAgent / APP-166 / pass
- app / Test_APP_166_DrillRejectsInvalidAgentOptionBeforeInvocation / APP-166 / pass
- app / Test_APP_166_DrillRejectsInvalidAgentOptionBeforeInvocation/missing / APP-166 / pass
- app / Test_APP_166_DrillRejectsInvalidAgentOptionBeforeInvocation/empty / APP-166 / pass
- app / Test_APP_166_DrillRejectsInvalidAgentOptionBeforeInvocation/repeated / APP-166 / pass
- app / Test_APP_167_DrillRejectsUnknownAgentAlias / APP-167 / pass
- app / Test_APP_140_DrillProjectAndConfigurationFailures / APP-140 / pass
- app / Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts / APP-141 / pass
- app / Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain / APP-141 / pass
- app / Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json / APP-141 / pass
- app / Test_APP_141_DrillSetupFailureHasNoArtifacts / APP-141 / pass
- app / Test_APP_142_DrillCleanUsageStopPlainAndJSON / APP-142 / pass
- app / Test_APP_136_DrillUsesConfiguredTimeout / APP-136 / pass
- app / Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty / APP-150, APP-151 / pass
- app / Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/plain / APP-150, APP-151 / pass
- app / Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/JSON / APP-150, APP-151 / pass
- app / Test_APP_150_APP_152_GCRejectsUsageBeforePruning / APP-150, APP-152 / pass
- app / Test_APP_152_GCConfigurationAndGitFailures / APP-152 / pass
- app / Test_APP_152_GCConfigurationAndGitFailures/malformed-config / APP-152 / pass
- app / Test_APP_152_GCConfigurationAndGitFailures/invalid-retention / APP-152 / pass
- app / Test_APP_156_HelpForms / APP-156 / pass
- app / Test_APP_156_HelpForms/help / APP-156 / pass
- app / Test_APP_156_HelpForms/--help / APP-156 / pass
- app / Test_APP_156_HelpForms/-h / APP-156 / pass
- app / Test_APP_156_HelpForms/help_coverage / APP-156 / pass
- app / Test_APP_156_HelpForms/agents-md_help / APP-156 / pass
- app / Test_APP_156_HelpForms/doctor_help / APP-156 / pass
- app / Test_APP_156_HelpForms/init_--help / APP-156 / pass
- app / Test_APP_156_HelpForms/init_-h / APP-156 / pass
- app / Test_APP_156_HelpForms/new-module_--help / APP-156 / pass
- app / Test_APP_156_HelpForms/new-module_-h / APP-156 / pass
- app / Test_APP_156_EachSupportedCommandHasCLIHelp / APP-156 / pass
- app / Test_APP_156_ExactHelpWordForm / APP-156 / pass
- app / Test_APP_157_HelpJSONMatchesPlain / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/help / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/--help / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/-h / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/help_coverage / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/agents-md_help / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/doctor_help / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/init_--help / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/init_-h / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/new-module_--help / APP-157 / pass
- app / Test_APP_157_HelpJSONMatchesPlain/new-module_-h / APP-157 / pass
- app / Test_APP_158_HelpFailures / APP-158 / pass
- app / Test_APP_158_HelpFailures/help_not-a-fireweed-command / APP-158 / pass
- app / Test_APP_158_HelpFailures/not-a-fireweed-command_help / APP-158 / pass
- app / Test_APP_158_HelpFailures/not-a-fireweed-command_--help / APP-158 / pass
- app / Test_APP_158_HelpFailures/not-a-fireweed-command_-h / APP-158 / pass
- app / Test_APP_158_HelpFailures/help_doctor_extra / APP-158 / pass
- app / Test_APP_158_HelpFailures/-h_doctor / APP-158 / pass
- app / Test_APP_158_HelpFailures/show_alpha_-h / APP-158 / pass
- app / Test_APP_158_HelpFailures/doctor_--help_--no-agent / APP-158 / pass
- app / Test_APP_158_HelpFailures/agents-md_--write_--help / APP-158 / pass
- app / Test_APP_158_HelpFailures/help_--bogus / APP-158 / pass
- app / Test_APP_158_HelpFailures/--help_-h / APP-158 / pass
- app / Test_APP_041_APP_042_APP_043_APP_044_APP_045_APP_046_APP_050_APP_053_APP_054_APP_059_GoInit / APP-041, APP-042, APP-043, APP-044, APP-045, APP-046, APP-050, APP-053, APP-054, APP-059 / pass
- app / Test_APP_043_APP_044_APP_047_APP_051_APP_053_APP_054_APP_161_NoneInit / APP-043, APP-044, APP-047, APP-051, APP-053, APP-054, APP-161 / pass
- app / Test_APP_050_APP_052_APP_161_NoneDefaultApp / APP-050, APP-052, APP-161 / pass
- app / Test_APP_053_APP_054_InitFileOnlyWorkerCommand / APP-053, APP-054 / pass
- app / Test_APP_053_APP_054_InitAuditCommandForms / APP-053, APP-054 / pass
- app / Test_APP_053_APP_054_InitAuditCommandForms/file-template / APP-053, APP-054 / pass
- app / Test_APP_053_APP_054_InitAuditCommandForms/claude-preset / APP-053, APP-054 / pass
- app / Test_APP_042_APP_048_APP_049_APP_050_APP_053_APP_054_APP_060_APP_161_LocalInit / APP-042, APP-048, APP-049, APP-050, APP-053, APP-054, APP-060, APP-161 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-recipe / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-agent / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-value / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/invalid-agent / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/malformed-option / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/extra-position / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/unknown-recipe / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/malformed-document / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/missing-document / APP-055, APP-056, APP-061 / pass
- app / Test_APP_055_APP_056_APP_061_InitUsageAndRecipeErrors/invalid-metadata-type / APP-055, APP-056, APP-061 / pass
- app / Test_APP_057_APP_058_InitConflicts / APP-057, APP-058 / pass
- app / Test_APP_057_APP_058_InitConflicts/config / APP-057, APP-058 / pass
- app / Test_APP_057_APP_058_InitConflicts/recipe / APP-057, APP-058 / pass
- app / Test_APP_057_APP_058_InitConflicts/module / APP-057, APP-058 / pass
- app / Test_APP_061_APP_161_BuiltinInitJSONSuccess / APP-061, APP-161 / pass
- app / Test_APP_061_APP_161_BuiltinInitJSONSuccess/go / APP-061, APP-161 / pass
- app / Test_APP_061_APP_161_BuiltinInitJSONSuccess/typescript-pnpm / APP-061, APP-161 / pass
- app / Test_APP_061_APP_161_BuiltinInitJSONSuccess/python-uv / APP-061, APP-161 / pass
- app / Test_APP_162_InitQueueDirectoryConflicts / APP-162 / pass
- app / Test_APP_162_InitQueueDirectoryConflicts/queue / APP-162 / pass
- app / Test_APP_162_InitQueueDirectoryConflicts/archive / APP-162 / pass
- app / Test_APP_009_APP_027_APP_028_StartContextAssignmentAndRejection / APP-009, APP-027, APP-028 / pass
- app / Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning / APP-009, APP-027, APP-028, APP-040 / pass
- app / Test_APP_004_APP_164_HeldModulesUseApprovedProjection / APP-004, APP-164 / pass
- app / Test_APP_011_APP_032_APP_033_ContextCompactFullAndErrors / APP-011, APP-032, APP-033 / pass
- app / Test_APP_168_ContextRecoversWithInvalidLiveConfiguration / APP-168 / pass
- app / Test_APP_168_ContextRecoversWithInvalidLiveConfiguration/malformed-yaml / APP-168 / pass
- app / Test_APP_168_ContextRecoversWithInvalidLiveConfiguration/excluding-pin / APP-168 / pass
- app / Test_APP_012_APP_034_APP_035_APP_036_CoverageCurrentProjectionAndErrors / APP-012, APP-034, APP-035, APP-036 / pass
- app / Test_APP_012_APP_160_APP_164_CoverageUsesApprovedHooksWithSelectedSpecs / APP-012, APP-160, APP-164 / pass
- app / Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks / APP-010, APP-029, APP-030, APP-031 / pass
- app / Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits / APP-010, APP-029, APP-030, APP-031 / pass
- app / Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit / APP-010, APP-029, APP-031 / pass
- app / Test_APP_114_APP_115_APP_117_ListQueueOrderAndFields / APP-114, APP-115, APP-117 / pass
- app / Test_APP_114_APP_117_ListNearestRootAndEmptyQueue / APP-114, APP-117 / pass
- app / Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors / APP-116, APP-117 / pass
- app / Test_APP_153_APP_154_MigrateNearestRootAndRepeatedNoOp / APP-153, APP-154 / pass
- app / Test_APP_153_APP_154_MigrateEmptyCurrentProject / APP-153, APP-154 / pass
- app / Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures / APP-154, APP-155 / pass
- app / Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/static-check / APP-154, APP-155 / pass
- app / Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/malformed-config / APP-154, APP-155 / pass
- app / Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/malformed-module / APP-154, APP-155 / pass
- app / Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/unsupported-version / APP-154, APP-155 / pass
- app / Test_APP_154_APP_155_MigrateStaticDiagnosticsAndFailures/missing-version / APP-154, APP-155 / pass
- app / Test_APP_155_MigrateRejectsArgumentsBeforeMutation / APP-155 / pass
- app / Test_APP_100_APP_101_APP_104_NewModuleFromNearestProjectTemplates / APP-100, APP-101, APP-104 / pass
- app / Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents / APP-102, APP-103, APP-104 / pass
- app / Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents/existing-name / APP-102, APP-103, APP-104 / pass
- app / Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents/existing-prefix / APP-102, APP-103, APP-104 / pass
- app / Test_APP_102_APP_103_APP_104_NewModuleConflictsPreserveContents/occupied-directory / APP-102, APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/missing-name / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/missing-prefix / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/APP_159_invalid-name / APP-103, APP-104, APP-159 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/APP_159_invalid-prefix / APP-103, APP-104, APP-159 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/repeated-prefix / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/extra-argument / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/unknown-option / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/missing-template / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/malformed-template / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/invalid-configuration / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/invalid-existing-graph / APP-103, APP-104 / pass
- app / Test_APP_103_APP_104_NewModuleUsageTemplatesAndProjectErrors/no-configured-destination / APP-103, APP-104 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/uppercase-name / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/repeated-name-hyphen / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/digit-led-name / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/non-ASCII-name / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/lowercase-prefix / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/prefix-hyphen / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/digit-led-prefix / APP-159 / pass
- app / Test_APP_159_NewModuleRejectsInvalidIdentities/non-ASCII-prefix / APP-159 / pass
- app / Test_APP_101_APP_104_APP_159_NewModuleEmptyGoModule / APP-101, APP-104, APP-159 / pass
- app / Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange / APP-074, APP-076, APP-077 / pass
- app / Test_APP_075_NextAllowUnmergedAndUsage / APP-075 / pass
- app / Test_APP_076_NextPlainEmptyQueue / APP-076 / pass
- app / Test_APP_077_NextExitClasses / APP-077 / pass
- app / Test_APP_078_NextWrappedCauseAndAction / APP-078 / pass
- app / Test_APP_079_NextJSONSuccessAndFailure / APP-079 / pass
- app / Test_APP_127_APP_128_APP_129_PRBodyExplicitHistoricalEndpoint / APP-127, APP-128, APP-129 / pass
- app / Test_APP_165_PRBodyWithIncompatibleRecordedPin / APP-165 / pass
- app / Test_APP_128_APP_129_PRBodyValidEmptyRange / APP-128, APP-129 / pass
- app / Test_APP_130_APP_131_PRBodyFailures / APP-130, APP-131 / pass
- app / Test_APP_130_APP_131_PRBodyFailures/missing-range / APP-130, APP-131 / pass
- app / Test_APP_130_APP_131_PRBodyFailures/extra-range / APP-130, APP-131 / pass
- app / Test_APP_130_APP_131_PRBodyFailures/invalid-option / APP-130, APP-131 / pass
- app / Test_APP_130_APP_131_PRBodyFailures/absent-project / APP-130, APP-131 / pass
- app / Test_APP_130_APP_131_PRBodyFailures/bad-endpoint-config / APP-130, APP-131 / pass
- app / Test_APP_130_APP_131_PRBodyFailures/unavailable-range / APP-130, APP-131 / pass
- app / Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor / APP-004, APP-016, APP-017, APP-018 / pass
- app / Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor/plain / APP-004, APP-016, APP-017, APP-018 / pass
- app / Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor/JSON / APP-004, APP-016, APP-017, APP-018 / pass
- app / Test_APP_004_APP_016_APP_017_APP_018_ModulesGraphAndNearestAncestor/nearest-ancestor / APP-004, APP-016, APP-017, APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/missing/plain / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/missing/JSON / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/malformed/plain / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/malformed/JSON / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/invalid-graph/plain / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/invalid-graph/JSON / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/invalid-document/plain / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/invalid-document/JSON / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/invalid-metadata/plain / APP-018 / pass
- app / Test_APP_018_ModulesConfigurationErrors/invalid-metadata/JSON / APP-018 / pass
- app / Test_APP_005_APP_020_NewChangePlainAndJSON / APP-005, APP-020 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/missing-argument/plain / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/missing-argument/JSON / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/extra-argument/plain / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/extra-argument/JSON / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/invalid-module/plain / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/invalid-module/JSON / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain / APP-019 / pass
- app / Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON / APP-019 / pass
- app / Test_APP_020_NewChangeJSONFailureEnvelope / APP-020 / pass
- app / Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork / APP-091, APP-092, APP-093 / pass
- app / Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork/plain-path / APP-091, APP-092, APP-093 / pass
- app / Test_APP_091_APP_092_APP_093_ReleaseOwnedClaimRetainsDirtyWork/json-filename / APP-091, APP-092, APP-093 / pass
- app / Test_APP_169_ReleaseRecoversWithInvalidLiveConfiguration / APP-169 / pass
- app / Test_APP_169_ReleaseRecoversWithInvalidLiveConfiguration/malformed-yaml / APP-169 / pass
- app / Test_APP_169_ReleaseRecoversWithInvalidLiveConfiguration/excluding-pin / APP-169 / pass
- app / Test_APP_092_APP_093_ReleaseForeignClaimKeepsOwner / APP-092, APP-093 / pass
- app / Test_APP_092_APP_093_ReleaseUsageAndConfiguration / APP-092, APP-093 / pass
- app / Test_APP_092_APP_093_ReleaseUsageAndConfiguration/missing-argument / APP-092, APP-093 / pass
- app / Test_APP_092_APP_093_ReleaseUsageAndConfiguration/extra-argument-json / APP-092, APP-093 / pass
- app / Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation / APP-080, APP-081, APP-083, APP-084 / pass
- app / Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON / APP-081, APP-084, APP-086 / pass
- app / Test_APP_083_RunPlainMultiple / APP-083 / pass
- app / Test_APP_086_RunJSONLimited / APP-086 / pass
- app / Test_APP_082_RunAllowUnmerged / APP-082 / pass
- app / Test_APP_081_APP_084_RunUsageAndConfiguration / APP-081, APP-084 / pass
- app / Test_APP_084_APP_085_APP_086_RunWrappedFailure / APP-084, APP-085, APP-086 / pass
- app / Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes / APP-084, APP-089, APP-090 / pass
- app / Test_APP_090_APP_084_RunForPlainStop / APP-084, APP-090 / pass
- app / Test_APP_084_RunForPriorBlockStillFails / APP-084 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation / APP-089 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation/missing / APP-089 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation/malformed / APP-089 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation/zero / APP-089 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation/negative / APP-089 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation/repeated / APP-089 / pass
- app / Test_APP_089_RunForInvalidOptionsBeforeMutation/next / APP-089 / pass
- app / Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext / APP-087 / pass
- app / Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext/plain-second-ref / APP-087 / pass
- app / Test_APP_087_StartFromRequestedAttemptRestoresWorkAndContext/json-first-ref / APP-087 / pass
- app / Test_APP_087_MissingSavedAttemptRetainsLifecycleFailureShape / APP-087 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/missing / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/noninteger / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/zero / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/negative / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/repeated / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/context / APP-088 / pass
- app / Test_APP_088_FromAttemptUsageRejectsBeforeMutation/help / APP-088 / pass
- app / Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter / APP-132, APP-133 / pass
- app / Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter/selected-ID / APP-132, APP-133 / pass
- app / Test_APP_132_APP_133_ShowHistoryOrderedMetadataAndFilter/module-wide / APP-132, APP-133 / pass
- app / Test_APP_132_APP_133_ShowHistoryNearestRootAndEmpty / APP-132, APP-133 / pass
- app / Test_APP_134_ShowHistoryOptionMisuse / APP-134 / pass
- app / Test_APP_134_ShowHistoryOptionMisuse/other-command/plain / APP-134 / pass
- app / Test_APP_134_ShowHistoryOptionMisuse/other-command/JSON / APP-134 / pass
- app / Test_APP_134_ShowHistoryOptionMisuse/projected-combination/plain / APP-134 / pass
- app / Test_APP_134_ShowHistoryOptionMisuse/projected-combination/JSON / APP-134 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/missing-config/plain / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/missing-config/JSON / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain / APP-135 / pass
- app / Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON / APP-135 / pass
- app / Test_APP_143_APP_144_StaleRecoveryDiffAndReadOnly / APP-143, APP-144 / pass
- app / Test_APP_144_StaleUnavailableHistoryAndEmptyView / APP-144 / pass
- app / Test_APP_143_StaleViewChoosesNearestProject / APP-143 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures / APP-145, APP-146 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/unknown / APP-145, APP-146 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/ambiguous / APP-145, APP-146 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/bad-bases / APP-145, APP-146 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/missing-config / APP-145, APP-146 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/malformed-config / APP-145, APP-146 / pass
- app / Test_APP_145_APP_146_StaleUsageLookupAndQueryFailures/invalid-graph / APP-145, APP-146 / pass
- app / Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport / APP-118, APP-119, APP-120, APP-121, APP-122, APP-123, APP-124, APP-125 / pass
- app / Test_APP_118_APP_119_APP_122_StatusNearestRootEmptyContainers / APP-118, APP-119, APP-122 / pass
- app / Test_APP_126_StatusUsageConfigurationAndOperationalFailures / APP-126 / pass
- app / Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume / APP-076, APP-077, APP-079 / pass
- app / Test_APP_076_APP_077_NextPlainUsageStop / APP-076, APP-077 / pass
- app / Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume / APP-083, APP-084, APP-086 / pass
- app / Test_APP_083_APP_084_RunPlainSharedUsageStop / APP-083, APP-084 / pass
- app / Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop / APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/next / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/run / APP-077, APP-079, APP-084, APP-086 / pass

## Dependent evals

## Contradicted Decisions
None. The four retained Decisions are all implemented as written:

- [0074-doctor-setup-commands] Doctor request Timeout is left unset (no new
  flag), and successful JSON checks emit `name`, `module`, `passed`, `detail`
  with diagnostics in the outer envelope; canaries are disambiguated by
  splitting the engine's `<probe> <scope>` label into the probe name and
  keeping the hook scope in `detail`, satisfying APP-067 without a new flag.
- [0112-new-module-command] The destination is the first configured module
  glob ending in a single wildcard component under a literal parent; when none
  exists the request is reported as C01 before any file is written; the
  package substitution replaces hyphens with underscores; go_module is read
  from go.mod when present and is otherwise empty.
- [0127-show-history-command] History JSON uses lowercase keys and preserves
  the project query's section strings; plain output labels each section and
  marks undated entries explicitly.
- init-prompt-file-commands Local recipes load recipe.yaml (module_defaults,
  hooks, protected, worker_writable), recipe.md (instructions) and a files/
  project-relative payload; `--recipe none` writes minimal setup with no
  hooks, templates or canaries and a generic module scaffold.

Reconciled tensions that did NOT rise to contradictions: APP-065 (warn-only
doctor is a success) coexisting with Diagnostic severities being open to the
envelope; the `envelope`'s optional `path` being emitted only when present and
required `warning: false` always serialized — both of these are demanded by
different evals and are satisfied jointly by omitting empty paths.

## Gaps
1. nearest spec APP-016, target module `app` (project/policy fix): no spec
   defines how project discovery behaves when no ancestor `fireweed.yaml`
   exists at all, even though APP-004 breaks exactly that case for held
   modules. The implementation falls back from "nearest ancestor
   fireweed.yaml" to a "nearest ancestor `.git` marker" purely so that the
   held approved snapshot can still be consulted; without it, modules/coverage
   crash with missing-config before ownership can be resolved. A spec should
   state the discovery order and whether every command or only held-eligible
   commands fall back.

2. nearest spec APP-163 (with APP-164, APP-168..170), target module `app`
   (compatibility policy): no spec enumerates which commands admit the live
   configuration vs the held approved configuration. APP-163 mandates
   admission for "commands using live project configuration" while APP-164
   prohibits live-config parsing on approved-snapshot paths, but the boundary
   matrix itself is only recoverable from the lifecycle evals. Capture it as a
   single policy (e.g. a shared `admit(command, selector)` table) in the spec,
   including the held carve-outs for context, release, verify, done, modules,
   coverage, next and run.

3. nearest spec APP-012 (with APP-160), target module `checks` + `app`: the
   spec says execution while held MUST use the approved hooks even for another
   queued change's projection, but does not say whether that projection's
   `set:` hook overrides apply to its own coverage run. The eval requires
   approved hooks to govern (the projected document's hook overrides are
   stripped) unless the projection is for the held change itself. Spec should
   state the spec-text cut explicitly so future grep-based hook resolution
   cannot drift.

4. nearest spec APP-015, target module `contracts` + `app`: the schema-1
   envelope guarantees code and message but not `subject`, `path`, or
   `warning` presence for all producers; the CLI had to rely on the
   dependency carrying selectors inside message text, which is why tests can
   only match substrings ("unknown module", "duplicate item ID"). Add a
   structured `subject` (or path) whenever ErrorCode reports one and update
   APP-015 to require it.

5. nearest spec APP-037, target module `project`: `project.Load` does not
   validate the dependency graph (it returns coded errors for configuration,
   documents, metadata but not for C02), while `project.Project` wraps bare
   ValidateGraph failures with C02 — so the CLI has to re-run ValidateGraph to
   report C02 for bare "modules"-style paths. The code assignment for graph
   failures depends on which entry point is used; APP-037 claims C02 is the
   graph code, but only some library paths deliver it. Move ValidateGraph into
   Load (with C02 coding) or document the split.

6. nearest spec APP-036, target module `checks`: `coverage --json` returns the
   raw `checks.Report`, including `stdout`/`stderr` of the eval hook process
   and `HookExit`; nothing in any spec bounds this data, so a hook that
   echoes secrets or full logs into the JUnit stream reaches consumers of the
   envelope. Decide whether the envelope projects a sanitized subset.

7. nearest spec "no spec", target module `assets` or `app`: nothing pins down
   the alias that init stores for custom `--agent`/`--audit-agent` templates in
   `agent.commands` (`worker-command`, `audit-command` here). Any future
   golden-output comparison for APP-053/054 would depend on these invented
   names; either pin them down or require only uniqueness + selectability via
   `agent.default`.

## Agent observations

### Invocation 1
Log: .fireweed/runs/20261119T134922-784504758274b8faa488e883/drill-1.log

Now let me read the evals, which define the expected CLI behavior:
Now let me look at key dependency internals — particularly error codes and runner stop reasons:
Now I have a good picture of the expected CLI. Let me check the Go environment, then start building:
Now I'll write the implementation. Starting with the shared foundation:
Now the info commands:
Now status:
Now the show command (current, projected, history, stale):
Now the authoring commands (new-change, stamp, check, coverage):
Let me fix issues in authoring.go:
Now the lifecycle commands:
Let me clean up lifecycle.go — remove dead code and fix the diagnostics ordering (error diagnostic must be first):
Now the maintenance commands, part 1 (gc, migrate, agents-md):
Now the doctor command:
Now ci and check-commits:
Now drill and pr-body:
Now new-module:
Now the biggest command — init:
Now the preflight and writer parts of init.go:
Now the dispatcher:
Now I see the right admission model (held admission only for modules/coverage/next/run/verify/done/context/release). Let me restructure:
Now add the loose root discovery (YAML or `.git` fallback) and rework modules:
Now add the modules-specific snapshot helper and remove the old `snapshotFor`:
Found it: `moduleFor` returns a typed `*contracts.CodedError` which, when nil, becomes a non-nil interface. Change to plain `error`:
Found the splitArgs bug — options stored as `false` are treated as disallowed:
Let me replicate the exact eval fixture with a scratch test to see the raw output:
The scratch fixture was missing the change file — let me redo faithfully:
Wait — the actual failing fixture has no change file?! Let me re-read the exact test source:
The stale flow must validate the config/graph *before* resolving the change. Fix:
Fix the held coverage path to take only the spec text from live projections, and improve the error message:
All tests pass individually. Let me check the workspace state and the drill hook:
Let me clean up the stray log file and write DRILL_NOTES.md:
Now let me run the target eval hook as requested (background, it's slow):
While the hook runs, let me check git status to confirm I only touched disposable paths:
Drill hook is still running. While it works, let me double-check a couple of trickier behaviors — version range handling for init and greedy stdin on the CLI:
The rebuild is complete and all tests pass individually (full suite: `ok 276s`). The target eval hook is running in the background; I'll confirm its result when it completes.
Done. The `app` module was rebuilt from its spec, Decisions, retained interfaces and evals.

**What was done**
- Rebuilt `packages/app/internal/entry/` (only disposable paths; `api.go`, `SPEC.md`, evals untouched). Result: `Run` dispatches all implemented commands with the schema‑1 envelope (`{"schema":1,"data":…,"diagnostics":[…]}`), plain output that mirrors JSON data, and corrective diagnostics with retained D006 codes.
- Coverage of all 164 APP specs: help/version/explain/playbook (works outside a project), modules (held approved snapshot while owned; live nearest‑ancestor otherwise), list, status (full rich report, exact JSON keys), new-change/stamp/check (incl. `--fix`), show (`--projected`/`--history`/`--stale` with recovered line diffs), lifecycle (start/verify/done/context/next/run/release/commits/abandon/verify `--quick`, `--from-attempt`, `--max`, `--for`, usage‑stop exit 3), coverage (approved hooks while held, live projection spec‑text for other changes, FIREWEED_BIN set, aliases for canonical/basename/path), gc/migrate/agents-md, init (builtin recipes, local recipe, `none`, all conflict/usage/format guards, queue dirs, .gitignore, managed AGENTS.md block), doctor (canaries with agents disabled, plain/JSON), new-module (project-owned templates, .tmpl stripping, grammar validation, C01/C02/C03/P01 mapping), check-commits/ci, pr-body, help forms, compatibility admission with the held‑eligibility matrix (modules/coverage/next/run/verify/done/context/release admit against the approved snapshot; everything else against live config).
- **All retained evals pass** — `go test -json ./packages/app/...`: 415 testcases, 0 failures; boundaries hook: clean.
- `.fireweed/drill-eval.xml` generated by the target eval hook — `PASS`.
- `DRILL_NOTES.md` written with separate `## Contradicted Decisions` (none; all four retained decisions implemented as written) and `## Gaps` (7 gaps, each with nearest spec ID and target module).

**Kept consistent with the spec**
- Compatibility admission runs before command effects/hooks/agents; held-adjacent commands never parse live fireweed.yaml; non‑held commands fail with C01 before anything runs.
- Held read paths resolve through `workflow.Approved` while other worktree claims are honored; read commands avoid git invocations in non‑git projects (with a `.git`‑marker fallback for held recovery).
- Git‑controlled commands (release/abandon/commits/gc) drop only the owned claim; init/`--no-app` variants, recipe unreadable/malformed (C03) vs invalid metadata types (C12), and byte‑preserving agents-md marker handling all follow the evals.

Untracked additions only (`packages/app/internal/`, `DRILL_NOTES.md`, generated `.fireweed/drill-eval.xml`); no tracked files were modified.

