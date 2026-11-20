---
agent: opencode
base_commit: a9ee7419fcd9c8f2f36705a4f25a518fe57d6245
date: "2026-11-19T10:07:42.456537Z"
model: opencode/inkling-2-small-free#high
result: pass
---

## Own evals
- checks / Test_CHK_005_GlobalAndModuleBoundaryHooks / CHK-005 / pass
- checks / Test_CHK_041_ExactBaselineOnly / CHK-041 / pass
- checks / Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns / CHK-039, CHK-040 / pass
- checks / Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional / CHK-044, CHK-045 / pass
- checks / Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut / CHK-041, CHK-042 / pass
- checks / Test_CHK_043_CHK_086_RequiredGlobalHookRejectsUnattributedNonzeroExit / CHK-043, CHK-086 / pass
- checks / Test_CHK_086_OverrideFailureIncludesBoundedUTF8StderrTail / CHK-086 / pass
- checks / Test_CHK_043_SilentNonzeroOverrideIsIndependentOfGlobalViolation / CHK-043 / pass
- checks / Test_CHK_043_LaunchCancellationAndTimeoutStayErrors / CHK-043 / pass
- checks / Test_CHK_065_CHK_066_DoctorAgentTailAcrossStreams / CHK-065, CHK-066 / pass
- checks / Test_CHK_065_CHK_066_DoctorAgentTailAcrossStreams/match_twentieth / CHK-065, CHK-066 / pass
- checks / Test_CHK_065_CHK_066_DoctorAgentTailAcrossStreams/match_twenty-first / CHK-065, CHK-066 / pass
- checks / Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams / CHK-081 / pass
- checks / Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput / CHK-081 / pass
- checks / Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput/failure / CHK-081 / pass
- checks / Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput/timeout / CHK-081 / pass
- checks / Test_CHK_081_DoctorPersistsCancelledProbeOutput / CHK-081 / pass
- checks / Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary / CHK-059 / pass
- checks / Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test / CHK-059 / pass
- checks / Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/subtest / CHK-059 / pass
- checks / Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/unrelated_test_with_shared_prefix / CHK-059 / pass
- checks / Test_CHK_084_CanaryNamesBeginWithProbeAndHookScope / CHK-084 / pass
- checks / Test_CHK_084_FailedRecoveryNameBeginsWithProbeAndHookScope / CHK-084 / pass
- checks / Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce / CHK-082, CHK-083 / pass
- checks / Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures / CHK-083 / pass
- checks / Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/launch / CHK-083 / pass
- checks / Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/nonzero / CHK-083 / pass
- checks / Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/timeout / CHK-083 / pass
- checks / Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/cancel / CHK-083 / pass
- checks / Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder / CHK-083 / pass
- checks / Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback / CHK-082, CHK-083 / pass
- checks / Test_CHK_049_StaticDiagnosticsStopProbes / CHK-049 / pass
- checks / Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe / CHK-050, CHK-051, CHK-052, CHK-053, CHK-055, CHK-056, CHK-057, CHK-058, CHK-059, CHK-060, CHK-061, CHK-067 / pass
- checks / Test_CHK_053_NestedModuleTemplatesRenderForBothCanaries / CHK-053 / pass
- checks / Test_CHK_051_CHK_052_SetupFailureStopsAndCleans / CHK-051, CHK-052 / pass
- checks / Test_CHK_051_CancellationCleansScratchAndStopsSetupGroup / CHK-051 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps / CHK-054 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps/missing_template / CHK-054 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps/malformed_template / CHK-054 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps/missing_baseline / CHK-054 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps/missing_internal_eval / CHK-054 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps/wrong_version / CHK-054 / pass
- checks / Test_CHK_054_MissingTemplatesAndCanaryMaps/malformed_canary / CHK-054 / pass
- checks / Test_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_ZeroHookAndRecovery / CHK-055, CHK-056, CHK-057, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific / CHK-055, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/missing_baseline_coverage / CHK-055, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/failing_eval_compile_error_without_report / CHK-055, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/failing_eval_with_unrelated_failure / CHK-055, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/missing_id_with_unrelated_failure / CHK-055, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_CHK_058_CHK_059_EvalEvidenceIsSpecific/missing_id_without_report / CHK-055, CHK-058, CHK-059 / pass
- checks / Test_CHK_055_BaselineAndRecoveryFailures / CHK-055 / pass
- checks / Test_CHK_055_BaselineAndRecoveryFailures/baseline_eval_failure / CHK-055 / pass
- checks / Test_CHK_055_BaselineAndRecoveryFailures/baseline_boundary_failure / CHK-055 / pass
- checks / Test_CHK_056_CHK_057_BoundaryViolationMustNameCanaryModule / CHK-056, CHK-057 / pass
- checks / Test_CHK_057_InternalEvalRequiresInterfaceViolation / CHK-057 / pass
- checks / Test_CHK_055_CHK_056_CHK_057_MissingRequiredBoundaryHook / CHK-055, CHK-056, CHK-057 / pass
- checks / Test_CHK_055_MissingRequiredHookBesideOverride / CHK-055 / pass
- checks / Test_CHK_055_ApplicableOverrideNeedsNoGlobalHook / CHK-055 / pass
- checks / Test_CHK_060_CHK_062_ModuleOverridesAndBoundaryWarnings / CHK-060, CHK-062 / pass
- checks / Test_CHK_060_OverridesMustDetectTheirCanaries / CHK-060 / pass
- checks / Test_CHK_060_OverridesMustDetectTheirCanaries/eval / CHK-060 / pass
- checks / Test_CHK_060_OverridesMustDetectTheirCanaries/boundary / CHK-060 / pass
- checks / Test_CHK_063_DoctorRunsNamedWorkerAndAuditInScratchThroughWrapper / CHK-063 / pass
- checks / Test_CHK_064_DoctorProbesInheritedWorkerDefaultAndWarns / CHK-064 / pass
- checks / Test_CHK_085_DoctorSelectionErrorSkipsOnlyUnresolvedRole / CHK-085 / pass
- checks / Test_CHK_085_DoctorSelectionErrorSkipsOnlyUnresolvedRole/worker / CHK-085 / pass
- checks / Test_CHK_085_DoctorSelectionErrorSkipsOnlyUnresolvedRole/audit / CHK-085 / pass
- checks / Test_CHK_063_CHK_065_CHK_066_CHK_067_DoctorAgentProbeOutcomes / CHK-063, CHK-065, CHK-066, CHK-067 / pass
- checks / Test_CHK_063_CHK_065_CHK_066_CHK_067_DoctorAgentProbeOutcomes/wrong_reply / CHK-063, CHK-065, CHK-066, CHK-067 / pass
- checks / Test_CHK_063_CHK_065_CHK_066_CHK_067_DoctorAgentProbeOutcomes/nonzero_with_reply / CHK-063, CHK-065, CHK-066, CHK-067 / pass
- checks / Test_CHK_063_CHK_065_CHK_066_CHK_067_DoctorAgentProbeOutcomes/ordinary_failure / CHK-063, CHK-065, CHK-066, CHK-067 / pass
- checks / Test_CHK_065_UnsetTimeoutBoundsAuditProbe / CHK-065 / pass
- checks / Test_CHK_068_CHK_069_ManagedGuidanceAndPlaybookBase / CHK-068, CHK-069 / pass
- checks / Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports / CHK-034 / pass
- checks / Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports/unmapped_PackageFailure / CHK-034 / pass
- checks / Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports/malformed_XML / CHK-034 / pass
- checks / Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports/missing_report / CHK-034 / pass
- checks / Test_CHK_034_EvaluateRetainsOutputOnHookTimeout / CHK-034 / pass
- checks / Test_CHK_070_HistoryExplicitRangeOrderAndFullOIDs / CHK-070 / pass
- checks / Test_CHK_070_HistoryBranchingRangeTopologicalOrder / CHK-070 / pass
- checks / Test_CHK_070_HistoryDefaultRangeUsesLocalConfiguredBase / CHK-070 / pass
- checks / Test_CHK_071_CHK_072_HistoryPhaseContentsAndDeletionControl / CHK-071, CHK-072 / pass
- checks / Test_CHK_072_CHK_073_CHK_077_HistoryPlaybookOverrideAndProjectedInterface / CHK-072, CHK-073, CHK-077 / pass
- checks / Test_CHK_071_CHK_073_HistoryAppliedImplTouchesEvalsAndInterface / CHK-071, CHK-073 / pass
- checks / Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename / CHK-074 / pass
- checks / Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename/modify / CHK-074 / pass
- checks / Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename/delete / CHK-074 / pass
- checks / Test_CHK_074_HistoryAuthorAppendOnlyModifyDeleteRename/rename / CHK-074 / pass
- checks / Test_CHK_074_HistoryMergeResolutionInMergeOnlyRange / CHK-074 / pass
- checks / Test_CHK_075_HistoryArchiveModifyAndDeletionControl / CHK-075 / pass
- checks / Test_CHK_076_HistoryMultipleSquashedChangesKeepH03H04 / CHK-076 / pass
- checks / Test_CHK_077_HistoryUsesRecordedBeforeAfterZones / CHK-077 / pass
- checks / Test_CHK_077_HistoryChecksBothRecordedZoneMaps / CHK-077 / pass
- checks / Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments / CHK-078 / pass
- checks / Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/invalid_range / CHK-078 / pass
- checks / Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/missing_config / CHK-078 / pass
- checks / Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/malformed_config / CHK-078 / pass
- checks / Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/missing_SPEC / CHK-078 / pass
- checks / Test_CHK_078_HistoryRejectsUnreadableRangeAndRecordedDocuments/malformed_change / CHK-078 / pass
- checks / Test_CHK_003_CHK_032_CHK_033_CHK_034_CHK_035_GlobalShellHookQuotesAndOnePassExpansion / CHK-003, CHK-032, CHK-033, CHK-034, CHK-035 / pass
- checks / Test_CHK_003_CHK_031_CHK_032_CHK_033_CHK_034_CHK_035_ModuleOverrideArgvAndOutput / CHK-003, CHK-031, CHK-032, CHK-033, CHK-034, CHK-035 / pass
- checks / Test_CHK_038_ExitAndReportAgreement / CHK-038 / pass
- checks / Test_CHK_038_ExitAndReportAgreement/zero-with-failure / CHK-038 / pass
- checks / Test_CHK_038_ExitAndReportAgreement/nonzero-with-pass / CHK-038 / pass
- checks / Test_CHK_038_ExitAndReportAgreement/nonzero-with-failure / CHK-038 / pass
- checks / Test_CHK_080_EvaluateRetainsParsedFailureDetails / CHK-080 / pass
- checks / Test_CHK_080_EvaluateRetainsParsedFailureDetails/nonzero_failure / CHK-080 / pass
- checks / Test_CHK_080_EvaluateRetainsParsedFailureDetails/zero_passing / CHK-080 / pass
- checks / Test_CHK_036_CHK_037_MissingReportAndLaunchFailure / CHK-036, CHK-037 / pass
- checks / Test_CHK_036_TimeoutFails / CHK-036 / pass
- checks / Test_CHK_046_EvaluateHookSetupCauses / CHK-046 / pass
- checks / Test_CHK_046_EvaluateHookSetupCauses/missing_hook / CHK-046 / pass
- checks / Test_CHK_046_EvaluateHookSetupCauses/invalid_hook_command / CHK-046 / pass
- checks / Test_CHK_046_EvaluateHookSetupCauses/unavailable_report_directory / CHK-046 / pass
- checks / Test_CHK_047_EvaluateUnknownModuleCause / CHK-047 / pass
- checks / Test_CHK_048_EvaluateModuleItemCauses / CHK-048 / pass
- checks / Test_CHK_048_EvaluateModuleItemCauses/malformed_item / CHK-048 / pass
- checks / Test_CHK_048_EvaluateModuleItemCauses/duplicate_item / CHK-048 / pass
- checks / Test_CHK_028_EvaluatePreservesJUnitDiagnostic / CHK-028 / pass
- checks / Test_CHK_003_CHK_032_ArgvPlaceholdersAreOnePass / CHK-003, CHK-032 / pass
- checks / Test_CHK_001_CHK_029_CHK_030_NestedJUnitIDsResultsAndTime / CHK-001, CHK-029, CHK-030 / pass
- checks / Test_CHK_001_CHK_029_CHK_030_SingleSuiteSkipped / CHK-001, CHK-029, CHK-030 / pass
- checks / Test_CHK_079_ParseJUnitRetainsFailureAndErrorDetails / CHK-079 / pass
- checks / Test_CHK_002_SkippedJUnitCaseDoesNotCover / CHK-002 / pass
- checks / Test_CHK_027_RejectMalformedXML / CHK-027 / pass
- checks / Test_CHK_028_RejectMissingAndUnknownIDs / CHK-028 / pass
- checks / Test_CHK_028_ParseJUnitAcceptsNumberedLabels / CHK-028 / pass
- checks / Test_CHK_028_ParseJUnitAcceptsNumberedLabels/pass / CHK-028 / pass
- checks / Test_CHK_028_ParseJUnitAcceptsNumberedLabels/fail / CHK-028 / pass
- checks / Test_CHK_028_ParseJUnitRejectsUnknownIDAlongsideKnownID / CHK-028 / pass
- checks / Test_CHK_002_CoverageRequiresCurrentNonSkippedCases / CHK-002 / pass
- checks / Test_CHK_002_CoverageExemptsReviewAndQuarantine / CHK-002 / pass
- checks / Test_CHK_025_FailFirstRequiresEachRequestedFailure / CHK-025 / pass
- checks / Test_CHK_025_FailFirstAlreadySatisfiedExemption / CHK-025 / pass
- checks / Test_CHK_026_PassRejectsNonQuarantinedFailures / CHK-026 / pass
- checks / Test_CHK_004_AuditOnlyCurrentChange / CHK-004 / pass
- checks / Test_CHK_006_CHK_007_CHK_008_CHK_009_RejectProtectedSpecForeignModuleAndChangeEveryPhase / CHK-006, CHK-007, CHK-008, CHK-009 / pass
- checks / Test_CHK_011_CHK_012_CHK_013_CHK_014_CHK_015_CHK_021_EvalsAllowedZonesAndNewStubs / CHK-011, CHK-012, CHK-013, CHK-014, CHK-015, CHK-021 / pass
- checks / Test_CHK_016_CHK_017_CHK_018_CHK_019_CHK_020_CHK_021_ImplFreezesEvalsAndInterfaces / CHK-016, CHK-017, CHK-018, CHK-019, CHK-020, CHK-021 / pass
- checks / Test_CHK_022_CHK_023_ImplAppendOnlyNewVersusExisting / CHK-022, CHK-023 / pass
- checks / Test_CHK_024_RenamesCheckBothPaths / CHK-024 / pass
- checks / Test_CHK_010_WorkerWritableOnlyEvalsAndImpl / CHK-010 / pass

## Dependent evals
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
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/failure / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/timeout / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/cancellation / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/append / RUN-063 / pass
- runner / Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput / RUN-064 / pass
- runner / Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries / RUN-065 / pass
- runner / Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries/usage_once / RUN-065 / pass
- runner / Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries/agent_retry / RUN-065 / pass
- runner / Test_RUN_066_InvokeAgentLogContextDefaultsAndExplicitFields / RUN-066 / pass
- runner / Test_RUN_041_UnapprovedChangeRequiresApproval / RUN-041 / pass
- runner / Test_RUN_041_RunPreservesApprovalRefusalCode / RUN-041 / pass
- runner / Test_RUN_042_OtherNextErrorsDoNotRequireApproval / RUN-042 / pass
- runner / Test_RUN_042_OtherNextErrorsDoNotRequireApproval/start_failure / RUN-042 / pass
- runner / Test_RUN_042_OtherNextErrorsDoNotRequireApproval/agent_failure / RUN-042 / pass
- runner / Test_RUN_021_AuditCorrectionWithinAllowance / RUN-021 / pass
- runner / Test_RUN_021_AuditCorrectionWithinAllowance/correction / RUN-021 / pass
- runner / Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry / RUN-021 / pass
- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals / RUN-021, RUN-070 / pass
- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return / RUN-021, RUN-070 / pass
- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry / RUN-021, RUN-070 / pass
- runner / Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance / RUN-021 / pass
- runner / Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock / RUN-070 / pass
- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit / RUN-013, RUN-045, RUN-046, RUN-071 / pass
- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override / RUN-013, RUN-045, RUN-046, RUN-071 / pass
- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback / RUN-013, RUN-045, RUN-046, RUN-071 / pass
- runner / Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration / RUN-013 / pass
- runner / Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry / RUN-048 / pass
- runner / Test_RUN_046_AuditCommitDoesNotEnterHeldGit / RUN-046 / pass
- runner / Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported / RUN-047, RUN-048 / pass
- runner / Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces / RUN-045, RUN-049 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically / RUN-047 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/change_outside_Audit / RUN-047 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/projected_spec / RUN-047 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/eval / RUN-047 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/added_file / RUN-047 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/deleted_file / RUN-047 / pass
- runner / Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/committed_invalid_edit / RUN-047 / pass
- runner / Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops / RUN-048, RUN-049 / pass
- runner / Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/agent_failure / RUN-048, RUN-049 / pass
- runner / Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/usage_stop / RUN-048, RUN-049 / pass
- runner / Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/time_stop / RUN-048, RUN-049 / pass
- runner / Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle / RUN-027, RUN-028 / pass
- runner / Test_RUN_027_RUN_028_RecoversCommitBeforeInvocationFailure / RUN-027, RUN-028 / pass
- runner / Test_RUN_053_DrillFiltersTrackedInputsAndSeedCommit / RUN-053 / pass
- runner / Test_RUN_053_DrillFiltersTrackedInputsAndSeedCommit/configured / RUN-053 / pass
- runner / Test_RUN_053_DrillFiltersTrackedInputsAndSeedCommit/omitted / RUN-053 / pass
- runner / Test_RUN_053_DrillFiltersTrackedInputsAndSeedCommit/empty / RUN-053 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/scalar / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/mapping / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/null_list / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/nonstring_entry / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/empty_entry / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/absolute_path / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/dot_component / RUN-075 / pass
- runner / Test_RUN_075_DrillRejectsMalformedExclusionsBeforeSetup/dot-dot_component / RUN-075 / pass
- runner / Test_RUN_076_DrillRejectsExcludedRequiredInputsBeforeSetup / RUN-076 / pass
- runner / Test_RUN_076_DrillRejectsExcludedRequiredInputsBeforeSetup/spec / RUN-076 / pass
- runner / Test_RUN_076_DrillRejectsExcludedRequiredInputsBeforeSetup/interface / RUN-076 / pass
- runner / Test_RUN_076_DrillRejectsExcludedRequiredInputsBeforeSetup/eval / RUN-076 / pass
- runner / Test_RUN_077_DrillRejectsExcludedPathRecreatedBySetup / RUN-077 / pass
- runner / Test_RUN_050_DrillPreflightRejectsWithoutEffects / RUN-050 / pass
- runner / Test_RUN_050_DrillPreflightRejectsWithoutEffects/unknown / RUN-050 / pass
- runner / Test_RUN_050_DrillPreflightRejectsWithoutEffects/dirty / RUN-050 / pass
- runner / Test_RUN_050_DrillPreflightRejectsWithoutEffects/claim / RUN-050 / pass
- runner / Test_RUN_073_DrillResolutionFailurePrecedesWorkspaceAndSetup / RUN-073 / pass
- runner / Test_RUN_073_DrillResolutionFailurePrecedesWorkspaceAndSetup/unknown_explicit_alias / RUN-073 / pass
- runner / Test_RUN_073_DrillResolutionFailurePrecedesWorkspaceAndSetup/unlisted_default / RUN-073 / pass
- runner / Test_RUN_051_SetupRunsDetachedAtHeadAndFailureStops / RUN-051 / pass
- runner / Test_RUN_052_RUN_053_RestrictedInputsAndIndependentGit / RUN-052, RUN-053 / pass
- runner / Test_RUN_054_PromptLetsWorkerRunOwnEval / RUN-054 / pass
- runner / Test_RUN_052_RUN_053_ProjectProtectedDisposableIsRetainedInputAndVerifierFile / RUN-052, RUN-053 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_api / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/delete_eval / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_preserved / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_protected_dependency_interface / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_setup / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/delete_setup / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions/add_foreign / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions/add_other / RUN-055 / pass
- runner / Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions/add_project_protected / RUN-055 / pass
- runner / Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound / RUN-056 / pass
- runner / Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound/agent_retry / RUN-056 / pass
- runner / Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound/verify_retry / RUN-056 / pass
- runner / Test_RUN_056_DrillUsesSelectedCommandAcrossRetries / RUN-056 / pass
- runner / Test_RUN_056_DrillUsesSelectedCommandAcrossRetries/agent_retry / RUN-056 / pass
- runner / Test_RUN_056_DrillUsesSelectedCommandAcrossRetries/verify_retry / RUN-056 / pass
- runner / Test_RUN_056_DrillUsesConfiguredDefault / RUN-056 / pass
- runner / Test_RUN_056_ConfiguredTimeoutBoundsDrillWorker / RUN-056 / pass
- runner / Test_RUN_057_UsageDoesNotSpendAttemptAndStopsWhenBudgetExhausted / RUN-057 / pass
- runner / Test_RUN_061_UsageStopRetainsAcceptedRebuild / RUN-061 / pass
- runner / Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail / RUN-058, RUN-059 / pass
- runner / Test_RUN_058_OwnCoverageRequiredAfterPassingCases / RUN-058 / pass
- runner / Test_RUN_059_DependentCoverageRequiredAfterPassingCases / RUN-059 / pass
- runner / Test_RUN_068_RUN_078_DependentOnlyFailuresEndDrill / RUN-068, RUN-078 / pass
- runner / Test_RUN_068_OwnEvalRetryOmitsDependentEvalOutput / RUN-068 / pass
- runner / Test_RUN_069_DrillObservationsIncludeStdoutAndRelativeLogPerAttempt / RUN-069 / pass
- runner / Test_RUN_060_RUN_061_DrillReportAndRetainedRef / RUN-060, RUN-061 / pass
- runner / Test_RUN_074_DrillResultAndReportIdentifySelectedCommand / RUN-074 / pass
- runner / Test_RUN_074_DrillResultAndReportIdentifySelectedCommand/default / RUN-074 / pass
- runner / Test_RUN_074_DrillResultAndReportIdentifySelectedCommand/explicit / RUN-074 / pass
- runner / Test_RUN_060_RUN_061_RUN_078_FailedDrillReportsFailuresAndRetainsCode / RUN-060, RUN-061, RUN-078 / pass
- runner / Test_RUN_062_CleanupOnSuccessAndFailures / RUN-062 / pass
- runner / Test_RUN_062_CleanupOnSuccessAndFailures/success / RUN-062 / pass
- runner / Test_RUN_062_CleanupOnSuccessAndFailures/agent_retry / RUN-062 / pass
- runner / Test_RUN_062_CleanupOnSuccessAndFailures/change_api / RUN-062 / pass
- runner / Test_RUN_062_CleanupOnSuccessAndFailures/usage_always / RUN-062 / pass
- runner / Test_RUN_062_CleanupOnSuccessAndFailures/timeout / RUN-062 / pass
- runner / Test_RUN_062_CleanupFailureCannotPass / RUN-062 / pass
- runner / Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput / RUN-001, RUN-002, RUN-005, RUN-006 / pass
- runner / Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput/worker / RUN-001, RUN-002, RUN-005, RUN-006 / pass
- runner / Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput/audit / RUN-001, RUN-002, RUN-005, RUN-006 / pass
- runner / Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput/audit_fallback / RUN-001, RUN-002, RUN-005, RUN-006 / pass
- runner / Test_RUN_001_InvokeAgentPropagatesResolutionError / RUN-001 / pass
- runner / Test_RUN_001_InvokeAgentPropagatesResolutionError/worker / RUN-001 / pass
- runner / Test_RUN_001_InvokeAgentPropagatesResolutionError/audit / RUN-001 / pass
- runner / Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues / RUN-003 / pass
- runner / Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues/shell / RUN-003 / pass
- runner / Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues/argv / RUN-003 / pass
- runner / Test_RUN_003_RUN_004_WrapperPrefixForBothRoles / RUN-003, RUN-004 / pass
- runner / Test_RUN_003_RUN_004_WrapperPrefixForBothRoles/worker / RUN-003, RUN-004 / pass
- runner / Test_RUN_003_RUN_004_WrapperPrefixForBothRoles/audit / RUN-003, RUN-004 / pass
- runner / Test_RUN_006_RUN_009_NonzeroExitRetainsStreamsAndCode / RUN-006, RUN-009 / pass
- runner / Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput / RUN-007, RUN-009 / pass
- runner / Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput/configured / RUN-007, RUN-009 / pass
- runner / Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput/request / RUN-007, RUN-009 / pass
- runner / Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant / RUN-007, RUN-009 / pass
- runner / Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant/timeout / RUN-007, RUN-009 / pass
- runner / Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant/cancel / RUN-007, RUN-009 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines / RUN-008 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/late_stderr_match / RUN-008 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/late_stdout_match / RUN-008 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/successful_match / RUN-008 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/outside_last_twenty / RUN-008 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/stdout_before_stderr_window / RUN-008 / pass
- runner / Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines/literal_substring / RUN-008 / pass
- runner / Test_RUN_009_LaunchFailureReturnsError / RUN-009 / pass
- runner / Test_RUN_010_EmptyQueue / RUN-010 / pass
- runner / Test_RUN_010_AllBlockedQueue / RUN-010 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/unmerged / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/start_failure / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/first_unblocked / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/unapproved_archive_bound / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/shared_base_approval / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents/allow_unmerged / RUN-011, RUN-012 / pass
- runner / Test_RUN_011_WaitingHeadSelectsIndependent / RUN-011 / pass
- runner / Test_RUN_010_EntirelyWaitingQueue / RUN-010 / pass
- runner / Test_RUN_013_RUN_014_CompleteSpecChange / RUN-013, RUN-014 / pass
- runner / Test_RUN_015_InvocationFailureRetainsHeldPhase / RUN-015 / pass
- runner / Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks / RUN-023, RUN-024 / pass
- runner / Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase / RUN-021, RUN-022 / pass
- runner / Test_RUN_021_ConfiguredThreeAttempts / RUN-021 / pass
- runner / Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks / RUN-023, RUN-024 / pass
- runner / Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce / RUN-003 / pass
- runner / Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/worker_shell / RUN-003 / pass
- runner / Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/worker_argv / RUN-003 / pass
- runner / Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/audit_shell / RUN-003 / pass
- runner / Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/audit_argv / RUN-003 / pass
- runner / Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback / RUN-003, RUN-067 / pass
- runner / Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback/worker / RUN-003, RUN-067 / pass
- runner / Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback/audit / RUN-003, RUN-067 / pass
- runner / Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback/audit_fallback / RUN-003, RUN-067 / pass
- runner / Test_RUN_067_WrapperOnlyPlaceholderCreatesPromptFile / RUN-067 / pass
- runner / Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure / RUN-067 / pass
- runner / Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/success / RUN-067 / pass
- runner / Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/nonzero / RUN-067 / pass
- runner / Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/timeout / RUN-067 / pass
- runner / Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/cancellation / RUN-067 / pass
- runner / Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/launch_failure / RUN-067 / pass
- runner / Test_RUN_067_UnselectedTemplateDoesNotRequirePromptFileIO / RUN-067 / pass
- runner / Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures / RUN-022, RUN-043 / pass
- runner / Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions / RUN-044 / pass
- runner / Test_RUN_016_UnlimitedQueueOrder / RUN-016 / pass
- runner / Test_RUN_016_NegativeMaxIsUnlimited / RUN-016 / pass
- runner / Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding / RUN-017, RUN-018, RUN-019 / pass
- runner / Test_RUN_018_InitialEmptyAndExactLimit / RUN-018 / pass
- runner / Test_RUN_018_InitialEmptyAndExactLimit/initial_empty / RUN-018 / pass
- runner / Test_RUN_018_InitialEmptyAndExactLimit/exact_limit / RUN-018 / pass
- runner / Test_RUN_020_HeldFailureRetainsEarlierResults / RUN-020 / pass
- runner / Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent / RUN-016, RUN-025, RUN-026 / pass
- runner / Test_RUN_025_RUN_026_PositiveMaxRetainsBlock / RUN-025, RUN-026 / pass
- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt / RUN-014 / pass
- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass / RUN-014 / pass
- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry / RUN-014 / pass
- runner / Test_RUN_072_RetainsStartWarning / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/verification_exhaustion / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/invocation_failure / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/usage_stop / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/Run_applied_result / RUN-072 / pass
- runner / Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied / RUN-036, RUN-037 / pass
- runner / Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded / RUN-036 / pass
- runner / Test_RUN_037_RUN_039_AgentDeadlineRetainsHeldWork / RUN-037, RUN-039 / pass
- runner / Test_RUN_037_VerificationDeadlineRetainsHeldWork / RUN-037 / pass
- runner / Test_RUN_038_UsageWaitPreflightStopsBeforeWaiting / RUN-038 / pass
- runner / Test_RUN_039_ClassifierAndParentCancellation / RUN-039 / pass
- runner / Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop / RUN-040 / pass
- runner / Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart / RUN-029 / pass
- runner / Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/bad_wait / RUN-029 / pass
- runner / Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/bad_budget / RUN-029 / pass
- runner / Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/zero_wait / RUN-029 / pass
- runner / Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/negative_wait / RUN-029 / pass
- runner / Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart/negative_budget / RUN-029 / pass
- runner / Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately / RUN-029, RUN-031 / pass
- runner / Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately/default_wait / RUN-029, RUN-031 / pass
- runner / Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately/default_budget / RUN-029, RUN-031 / pass
- runner / Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep / RUN-030, RUN-031, RUN-033 / pass
- runner / Test_RUN_033_OwnedResumeUsesApprovedSnapshot / RUN-033 / pass
- runner / Test_RUN_033_OwnedResumeUsesApprovedSnapshot/different_budget / RUN-033 / pass
- runner / Test_RUN_033_OwnedResumeUsesApprovedSnapshot/malformed_YAML / RUN-033 / pass
- runner / Test_RUN_030_UsageRetryPreservesVerificationFeedback / RUN-030 / pass
- runner / Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases / RUN-030, RUN-031 / pass
- runner / Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied / RUN-032 / pass
- runner / Test_RUN_034_UsageStopClassifierWrapping / RUN-034 / pass
- runner / Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure / RUN-035 / pass
- runner / Test_RUN_031_CancelledUsageWaitRetainsHeldWork / RUN-031 / pass
- workflow / Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange / FLOW-089, FLOW-091, FLOW-094 / pass
- workflow / Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange / FLOW-089, FLOW-091, FLOW-094 / pass
- workflow / Test_FLOW_090_AbandonRejectsDirtyWorkingTree / FLOW-090 / pass
- workflow / Test_FLOW_090_AbandonRejectsDirtyWorkingTree/tracked / FLOW-090 / pass
- workflow / Test_FLOW_090_AbandonRejectsDirtyWorkingTree/untracked / FLOW-090 / pass
- workflow / Test_FLOW_091_AbandonRejectsForeignClaim / FLOW-091 / pass
- workflow / Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending / FLOW-092 / pass
- workflow / Test_FLOW_093_FLOW_095_AbandonRetainsRealAttemptAndFreesAssignedID / FLOW-093, FLOW-095 / pass
- workflow / Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal / FLOW-096 / pass
- workflow / Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal/changes/0001-manual.md / FLOW-096 / pass
- workflow / Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal/0001-manual.md / FLOW-096 / pass
- workflow / Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal/manual / FLOW-096 / pass
- workflow / Test_FLOW_019_ClaimReloadPreservesApprovedConfigNumbers / FLOW-019 / pass
- workflow / Test_FLOW_019_ClaimReloadPreservesApprovedConfigNumbers/1.0 / FLOW-019 / pass
- workflow / Test_FLOW_019_ClaimReloadPreservesApprovedConfigNumbers/1 / FLOW-019 / pass
- workflow / Test_FLOW_019_CurrentResolvesHeldPathWithInvalidLiveConfig / FLOW-019 / pass
- workflow / Test_FLOW_019_CurrentResolvesHeldPathWithInvalidLiveConfig/malformed_YAML / FLOW-019 / pass
- workflow / Test_FLOW_019_CurrentResolvesHeldPathWithInvalidLiveConfig/incompatible_version_pin / FLOW-019 / pass
- workflow / Test_FLOW_019_FLOW_139_ApprovedSurvivesLiveEdits / FLOW-019, FLOW-139 / pass
- workflow / Test_FLOW_140_ApprovedReportsNoOwnedClaim / FLOW-140 / pass
- workflow / Test_FLOW_140_ApprovedRejectsUnreadableOrMalformedClaim / FLOW-140 / pass
- workflow / Test_FLOW_141_StartWarnsWhenBaseTipIsOutsideHead / FLOW-141 / pass
- workflow / Test_FLOW_141_StartWarnsWhenBaseTipIsOutsideHead/Start/behind_remote / FLOW-141 / pass
- workflow / Test_FLOW_141_StartWarnsWhenBaseTipIsOutsideHead/Start/diverged_from_local_fallback / FLOW-141 / pass
- workflow / Test_FLOW_141_StartWarnsWhenBaseTipIsOutsideHead/StartFromAttempt/behind_remote / FLOW-141 / pass
- workflow / Test_FLOW_141_StartWarnsWhenBaseTipIsOutsideHead/StartFromAttempt/diverged_from_local_fallback / FLOW-141 / pass
- workflow / Test_FLOW_142_StartOmitsWarningWhenHeadContainsBaseTip / FLOW-142 / pass
- workflow / Test_FLOW_142_StartOmitsWarningWhenHeadContainsBaseTip/Start/equal_to_remote / FLOW-142 / pass
- workflow / Test_FLOW_142_StartOmitsWarningWhenHeadContainsBaseTip/Start/ahead_of_local_fallback / FLOW-142 / pass
- workflow / Test_FLOW_142_StartOmitsWarningWhenHeadContainsBaseTip/StartFromAttempt/equal_to_remote / FLOW-142 / pass
- workflow / Test_FLOW_142_StartOmitsWarningWhenHeadContainsBaseTip/StartFromAttempt/ahead_of_local_fallback / FLOW-142 / pass
- workflow / Test_FLOW_058_StartRejectsStaleSpecHashWithNewerBaseTip / FLOW-058 / pass
- workflow / Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState / FLOW-063, FLOW-064, FLOW-065, FLOW-066, FLOW-067 / pass
- workflow / Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState/ignored / FLOW-063, FLOW-064, FLOW-065, FLOW-066, FLOW-067 / pass
- workflow / Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState/unignored / FLOW-063, FLOW-064, FLOW-065, FLOW-066, FLOW-067 / pass
- workflow / Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit / FLOW-066 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead / FLOW-063 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/no_claim / FLOW-063 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/foreign_claim / FLOW-063 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/author / FLOW-063 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/phase / FLOW-063 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/audit / FLOW-063 / pass
- workflow / Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead/changed_HEAD / FLOW-063 / pass
- workflow / Test_FLOW_087_FLOW_088_CommitsFollowPendingAndArchivePhaseContents / FLOW-087, FLOW-088 / pass
- workflow / Test_FLOW_087_CommitsUseFirstParentOfMerge / FLOW-087 / pass
- workflow / Test_FLOW_088_CommitsRejectMissingAndAmbiguousSelectors / FLOW-088 / pass
- workflow / Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot / FLOW-097 / pass
- workflow / Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts / FLOW-098 / pass
- workflow / Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure / FLOW-099 / pass
- workflow / Test_FLOW_099_DefaultBudgetBoundaries / FLOW-099 / pass
- workflow / Test_FLOW_099_DefaultBudgetBoundaries/equal / FLOW-099 / pass
- workflow / Test_FLOW_099_DefaultBudgetBoundaries/above / FLOW-099 / pass
- workflow / Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits / FLOW-100 / pass
- workflow / Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim / FLOW-133 / pass
- workflow / Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim/unclaimed_queued_change / FLOW-133 / pass
- workflow / Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim/empty_valid_queue / FLOW-133 / pass
- workflow / Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim/claim_owned_by_another_worktree / FLOW-133 / pass
- workflow / Test_FLOW_068_EligibleOrderAndEmptyQueues / FLOW-068 / pass
- workflow / Test_FLOW_068_EligibleOrderAndEmptyQueues/skips_waiting_and_retains_independent_order / FLOW-068 / pass
- workflow / Test_FLOW_068_EligibleOrderAndEmptyQueues/empty_queue / FLOW-068 / pass
- workflow / Test_FLOW_068_EligibleOrderAndEmptyQueues/all_waiting / FLOW-068 / pass
- workflow / Test_FLOW_069_EligibleWaitsForSameModule / FLOW-069 / pass
- workflow / Test_FLOW_070_EligibleWaitsForDirectAndTransitiveDependents / FLOW-070 / pass
- workflow / Test_FLOW_071_EligibleInterfaceChangeWaitsForBlockedDependent / FLOW-071 / pass
- workflow / Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet / FLOW-072 / pass
- workflow / Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet/preceding_set_removes_dependency / FLOW-072 / pass
- workflow / Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet/own_set_adds_or_removes_dependency / FLOW-072 / pass
- workflow / Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead / FLOW-011 / pass
- workflow / Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/empty_target_selects_first_eligible / FLOW-011 / pass
- workflow / Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/explicit_first_eligible / FLOW-011 / pass
- workflow / Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/reject_waiting / FLOW-011 / pass
- workflow / Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/reject_later-eligible / FLOW-011 / pass
- workflow / Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead/no_eligible_change / FLOW-011 / pass
- workflow / Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits / FLOW-036, FLOW-078, FLOW-079 / pass
- workflow / Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals / FLOW-079 / pass
- workflow / Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals/failing / FLOW-079 / pass
- workflow / Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals/missing-listed-coverage / FLOW-079 / pass
- workflow / Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks / FLOW-078 / pass
- workflow / Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/evals_interface / FLOW-078 / pass
- workflow / Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/impl_implementation / FLOW-078 / pass
- workflow / Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit / FLOW-036 / pass
- workflow / Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly / FLOW-128 / pass
- workflow / Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts / FLOW-129 / pass
- workflow / Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate / FLOW-130, FLOW-131 / pass
- workflow / Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/configured / FLOW-130, FLOW-131 / pass
- workflow / Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/default / FLOW-130, FLOW-131 / pass
- workflow / Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning / FLOW-131 / pass
- workflow / Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning/0s / FLOW-131 / pass
- workflow / Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning/-1h / FLOW-131 / pass
- workflow / Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning/invalid / FLOW-131 / pass
- workflow / Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState / FLOW-132 / pass
- workflow / Test_FLOW_001_FLOW_006_FLOW_007_FLOW_008_FLOW_012_FLOW_013_FLOW_014_FLOW_017_FLOW_018_FLOW_019_FLOW_053_StartAssignsLateIDsAndRetainsSnapshotWithoutCommit / FLOW-001, FLOW-006, FLOW-007, FLOW-008, FLOW-012, FLOW-013, FLOW-014, FLOW-017, FLOW-018, FLOW-019, FLOW-053 / pass
- workflow / Test_FLOW_007_StartReservesCurrentAndPreviouslyRemovedIDs / FLOW-007 / pass
- workflow / Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead / FLOW-010, FLOW-011 / pass
- workflow / Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead/tracked / FLOW-010, FLOW-011 / pass
- workflow / Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead/untracked / FLOW-010, FLOW-011 / pass
- workflow / Test_FLOW_001_FLOW_016_FLOW_017_LinkedWorktreeExcludesCompetingOwner / FLOW-001, FLOW-016, FLOW-017 / pass
- workflow / Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit / FLOW-009, FLOW-015 / pass
- workflow / Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/default-below / FLOW-009, FLOW-015 / pass
- workflow / Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/default-at / FLOW-009, FLOW-015 / pass
- workflow / Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/configured-at / FLOW-009, FLOW-015 / pass
- workflow / Test_FLOW_055_VerifyAndDoneRejectUnsupportedHeldPhase / FLOW-055 / pass
- workflow / Test_FLOW_055_UnsupportedPhasePrecedesProtectedEdit / FLOW-055 / pass
- workflow / Test_FLOW_056_FLOW_061_CurrentAndContextIdentifyMissingClaimsAndChanges / FLOW-056, FLOW-061 / pass
- workflow / Test_FLOW_058_StaleLifecycleBaseKeepsCodedCause / FLOW-058 / pass
- workflow / Test_FLOW_062_MissingLifecycleBaseKeepsC08 / FLOW-062 / pass
- workflow / Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks / FLOW-002, FLOW-020, FLOW-021 / pass
- workflow / Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks/author / FLOW-002, FLOW-020, FLOW-021 / pass
- workflow / Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks/eval-zone / FLOW-002, FLOW-020, FLOW-021 / pass
- workflow / Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst / FLOW-022, FLOW-023, FLOW-031 / pass
- workflow / Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/missing-coverage / FLOW-022, FLOW-023, FLOW-031 / pass
- workflow / Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/passing-target-unrelated-failure / FLOW-022, FLOW-023, FLOW-031 / pass
- workflow / Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/already-satisfied / FLOW-022, FLOW-023, FLOW-031 / pass
- workflow / Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst/fixes / FLOW-022, FLOW-023, FLOW-031 / pass
- workflow / Test_FLOW_023_AddedIDRequiresMappedFailingCase / FLOW-023 / pass
- workflow / Test_FLOW_023_ReviewTaggedAdditionNeedsNoFailFirstCase / FLOW-023 / pass
- workflow / Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries / FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 / pass
- workflow / Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/frozen-eval / FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 / pass
- workflow / Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/direct / FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 / pass
- workflow / Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/all / FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 / pass
- workflow / Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/none / FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 / pass
- workflow / Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/quick-boundary / FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 / pass
- workflow / Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures / FLOW-134 / pass
- workflow / Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/direct / FLOW-134 / pass
- workflow / Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/all / FLOW-134 / pass
- workflow / Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/none / FLOW-134 / pass
- workflow / Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_pass/all / FLOW-134 / pass
- workflow / Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal / FLOW-032, FLOW-033, FLOW-046, FLOW-135 / pass
- workflow / Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal/default_audit/ignored / FLOW-032, FLOW-033, FLOW-046, FLOW-135 / pass
- workflow / Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal/default_audit/unignored / FLOW-032, FLOW-033, FLOW-046, FLOW-135 / pass
- workflow / Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal/disabled_audit/ignored / FLOW-032, FLOW-033, FLOW-046, FLOW-135 / pass
- workflow / Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings / FLOW-034, FLOW-036, FLOW-041, FLOW-046 / pass
- workflow / Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_FLOW_146_AuditFindingsConvergeAndPersistTwoRoundBound / FLOW-035, FLOW-037, FLOW-038, FLOW-039, FLOW-042, FLOW-043, FLOW-046, FLOW-146 / pass
- workflow / Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion / FLOW-040, FLOW-042, FLOW-043 / pass
- workflow / Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives / FLOW-004, FLOW-028, FLOW-044, FLOW-045, FLOW-046, FLOW-047, FLOW-050 / pass
- workflow / Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead / FLOW-003, FLOW-024, FLOW-048, FLOW-049 / pass
- workflow / Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_FLOW_143_FLOW_144_ContextAndProvenanceAndRetainedReports / FLOW-005, FLOW-051, FLOW-052, FLOW-053, FLOW-143, FLOW-144 / pass
- workflow / Test_FLOW_052_DoneRetainsCostModuleReportBesideCostSnapshot / FLOW-052 / pass
- workflow / Test_FLOW_052_DoneRetainsCostModuleReportBesideCostSnapshot/costs / FLOW-052 / pass
- workflow / Test_FLOW_052_DoneRetainsCostModuleReportBesideCostSnapshot/CoStS / FLOW-052 / pass
- workflow / Test_FLOW_054_StartOperationalPreconditionsHaveP01 / FLOW-054 / pass
- workflow / Test_FLOW_056_StartContentionAndCurrentForeignClaimHaveP02 / FLOW-056 / pass
- workflow / Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks / FLOW-057 / pass
- workflow / Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/author / FLOW-057 / pass
- workflow / Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/runtime / FLOW-057 / pass
- workflow / Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/audit / FLOW-057 / pass
- workflow / Test_FLOW_059_AuditNoFindingsErrorsHaveP09 / FLOW-059 / pass
- workflow / Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim / FLOW-060 / pass
- workflow / Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses / FLOW-062 / pass
- workflow / Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics / FLOW-062 / pass
- workflow / Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly / FLOW-077 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_114_explicit_endpoints / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_115_selected_identities / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_116_provenance_and_evals / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_117_phase_file_categories / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_118_preserved_edits / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_119_review_specs / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_120_prior_corrected_audit / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint/FLOW_121_decisions / FLOW-114, FLOW-115, FLOW-116, FLOW-117, FLOW-118, FLOW-119, FLOW-120, FLOW-121 / pass
- workflow / Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts / FLOW-122, FLOW-123 / pass
- workflow / Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts/FLOW_122_endpoint_exceptions / FLOW-122, FLOW-123 / pass
- workflow / Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts/FLOW_123_latest_saved_costs / FLOW-122, FLOW-123 / pass
- workflow / Test_FLOW_126_PRBodyEndpointConfigError / FLOW-126 / pass
- workflow / Test_FLOW_127_PRBodyUnresolvableRangeIsNotEndpointConfigError / FLOW-127 / pass
- workflow / Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges / FLOW-115, FLOW-124, FLOW-125 / pass
- workflow / Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges/FLOW_124_binary_and_large_artifacts / FLOW-115, FLOW-124, FLOW-125 / pass
- workflow / Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges/FLOW_115_FLOW_125_empty_and_invalid_ranges / FLOW-115, FLOW-124, FLOW-125 / pass
- workflow / Test_FLOW_080_FLOW_081_FLOW_083_FLOW_143_FLOW_144_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive / FLOW-080, FLOW-081, FLOW-083, FLOW-143, FLOW-144 / pass
- workflow / Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate / FLOW-082 / pass
- workflow / Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/missing / FLOW-082 / pass
- workflow / Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/older_passing / FLOW-082 / pass
- workflow / Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/newer_failed / FLOW-082 / pass
- workflow / Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/newer_passing / FLOW-082 / pass
- workflow / Test_FLOW_081_RecordedRegenerateStartRetainsDisposableFilesAndState / FLOW-081 / pass
- workflow / Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook / FLOW-083 / pass
- workflow / Test_FLOW_084_ReleaseOwnedClaimPreservesDirtyWorkAndRecordedState / FLOW-084 / pass
- workflow / Test_FLOW_085_ReleaseMissingAndForeignClaimsPreserveRefs / FLOW-085 / pass
- workflow / Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors / FLOW-086 / pass
- workflow / Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors/path / FLOW-086 / pass
- workflow / Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors/filename / FLOW-086 / pass
- workflow / Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors/slug / FLOW-086 / pass
- workflow / Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptRestoresEvalsStepAndRetainsCurrentApproval / FLOW-073, FLOW-074, FLOW-075 / pass
- workflow / Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance / FLOW-073, FLOW-074, FLOW-075 / pass
- workflow / Test_FLOW_074_StartFromAttemptSelectsRequestedSavedRef / FLOW-074 / pass
- workflow / Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval / FLOW-075 / pass
- workflow / Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange / FLOW-073 / pass
- workflow / Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/dirty / FLOW-073 / pass
- workflow / Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/still_blocked / FLOW-073 / pass
- workflow / Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/not_first_eligible / FLOW-073 / pass
- workflow / Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange/foreign_claim / FLOW-073 / pass
- workflow / Test_FLOW_076_StartFromAttemptRejectsInvalidOrMissingAttemptBeforeMutation / FLOW-076 / pass
- workflow / Test_FLOW_059_FLOW_145_AuditVerifyChecksSubmissionWithoutAdvancing / FLOW-059, FLOW-145 / pass
- workflow / Test_FLOW_002_FLOW_136_AuditRejectsForbiddenEditAndDoneMarksFailure / FLOW-002, FLOW-136 / pass
- workflow / Test_FLOW_061_UnknownLifecycleSelectorsNameValidSlugs / FLOW-061 / pass
- workflow / Test_FLOW_101_FLOW_102_FLOW_103_StatusQueueAndForwardWaiting / FLOW-101, FLOW-102, FLOW-103 / pass
- workflow / Test_FLOW_104_FLOW_105_StatusInterfaceAndProjectedWaiting / FLOW-104, FLOW-105 / pass
- workflow / Test_FLOW_106_FLOW_107_FLOW_108_StatusModuleMetadata / FLOW-106, FLOW-107, FLOW-108 / pass
- workflow / Test_FLOW_109_StatusProvenanceDrift / FLOW-109 / pass
- workflow / Test_FLOW_111_StatusReworkChainWarning / FLOW-111 / pass
- workflow / Test_FLOW_111_StatusReworkChainWarning/three_linked_changes / FLOW-111 / pass
- workflow / Test_FLOW_111_StatusReworkChainWarning/unrelated_changes_before_tenth_predecessor / FLOW-111 / pass
- workflow / Test_FLOW_111_StatusReworkChainWarning/predecessor_more_than_ten_applied_changes_back / FLOW-111 / pass
- workflow / Test_FLOW_111_StatusReworkChainWarning/latest_matching_predecessor_wins / FLOW-111 / pass
- workflow / Test_FLOW_111_StatusReworkChainWarning/same_module_without_shared_IDs / FLOW-111 / pass
- workflow / Test_FLOW_110_StatusSharedBaseUsesSameArchivedDocumentIdentity / FLOW-110 / pass
- workflow / Test_FLOW_110_StatusReportsUnavailableConfiguredBaseWithEmptyArchive / FLOW-110 / pass
- workflow / Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle / FLOW-112 / pass
- workflow / Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings / FLOW-113 / pass
- workflow / Test_FLOW_136_DoneMarksFullVerificationFailureAndPreservesResult / FLOW-136 / pass
- workflow / Test_FLOW_136_DoneDistinguishesSameCodeByOrigin / FLOW-136 / pass
- workflow / Test_FLOW_137_DoneSavesWorkerDeclaredBlock / FLOW-137 / pass
- workflow / Test_FLOW_137_DoneSavesWorkerDeclaredBlock/evals / FLOW-137 / pass
- workflow / Test_FLOW_137_DoneSavesWorkerDeclaredBlock/impl / FLOW-137 / pass
- workflow / Test_FLOW_137_DoneValidatesApprovalBeforeRecordingWorkerBlock / FLOW-137 / pass
- workflow / Test_FLOW_138_ContextExplainsHowToDeclareWorkerBlock / FLOW-138 / pass
- workflow / Test_FLOW_138_ContextExplainsHowToDeclareWorkerBlock/evals / FLOW-138 / pass
- workflow / Test_FLOW_138_ContextExplainsHowToDeclareWorkerBlock/impl / FLOW-138 / pass

## Contradicted Decisions
None. Both decisions reconstruct cleanly:

- `[0071-doctor-probes]` — Doctor requests the literal `FIREWEED_DOCTOR_OK` reply;
  `limit_patterns` match case-insensitively as literal substrings within the final 20
  lines; playbook overrides are checked against the hexadecimal SHA-256 of the embedded
  `assets.Playbook(name)` text.
- `[0114-history-validation]` — a merge retaining a phase value already held by a parent
  is not treated as a new phase assignment (phase equality test); merge commits are
  inspected with `--cc` combined diffs for append-only paths; squashed entries identify
  the change's repository-relative after-path; non-merge phase transitions compare the
  sole (first) parent.

## Gaps
Each gap names the nearest spec ID and the module a fix would land in (all: `checks`,
unless noted otherwise).

- CHK-043/CHK-041 — the suppression/exit interaction is underspecified. A hook that runs
  normally (any exit code) with all module-attributed lines suppressed yields no error,
  while a hook with attribution remaining also emits only X01 diagnostics and no
  exit-status error, and a normally-exited hook without any module-attributed line at all
  becomes an X02 error with an exit-status cause. The spec says "exits nonzero without
  emitting any module-attributed violation" but not how suppressed lines count in that
  rule; I treat suppression as happening before the rule. (checks)
- CHK-086 — the spec only requires the X02 diagnostic to identify the hook as "global hook
or module override" and to include the bounded stderr tail; no canonical label exists, so
  the wording (`global boundaries hook`, `<name> boundaries override hook`) is invented.
  Tests accept several variants of it. (checks)
- CHK-077 — "the worker change's projected set values" does not say which `set` keys
  project into zone evaluation or how they merge with before/after maps. I apply
  before/after union plus set overrides for interface/evals/disposable/preserve/
  append_only/protected. (checks)
- CHK-076 — for merge commits the spec does not say which parent's change-file contents
  identify phases; I compare the first parent for phase detection and all parents for
  squash exemption. (checks)
- CHK-050/CHK-053 — the scratch worktree below the common Git directory must hold the
  caller's *current* files; I `git worktree add --detach` then copy the caller tree in
  (excluding `.git` and `.fireweed/runs`) instead of checking out HEAD, and I re-widen
  `modules` in the scratch copy's fireweed.yaml with `packages/*` so the canary modules
  are discoverable via FIREWEED_BIN's modules output (CHK-061). (checks)
- CHK-055 — the canary gate semantics are inferred from fixtures, not the spec text:
  every applicable baseline must pass before any canary is planted; any failing baseline
  blocks all canary probes across all scopes and probe kinds, and skipped canaries produce
  C14s naming the canary and (for boundary canaries) the boundaries hook. Also, scratch
  boundary enforcement (missing-hook X02 / baseline failure) is evaluated for the caller's
  real modules, never for the throwaway canary modules themselves. (checks)
- CHK-055 — failing_eval / missing_id detection rules are inferred: failing_eval needs a
  failed referencing testcase plus a nonzero hook exit; missing_id needs a mapping
  rejection (C11 with `has no known spec ID` or `references unknown spec ID`) whose
  testcase name matches the Go test function (or the bare token line) from the missing_id
  canary file. The reporting canaries' detection semantics are otherwise unspecified.
  (checks)
- CHK-079/CHK-030 — failure element text retention is implemented as message + newline +
  trimmed body in `Case.Failure`; pass/skip testcases carry empty Failure as required.
  (No gap in behavior; noted because the exact format of the joined string is invented.)
- CHK-084 — the hook-scope suffix for canary check names (`global eval`, `alpha eval`,
  `global boundaries`, `alpha boundaries`) and the per-throwaway baseline duplication
  policy (one check per scope per throwaway module) are invented to satisfy both the
  name-prefix contract and the two-distinct-throwaway-modules requirement. (checks)
- CHK-065 — the shared usage-limit wording and C15 message skeleton (final 20 output
  lines, kind label, exit status) is my invented format; only substrings are tested.
  (checks)
- CHK-050/CHK-062 — which modules count as "real" for retention/enforcement purposes is
  inferred: all modules loaded from the caller's fireweed.yaml before the scratch tree
  exists; throwaway canary modules are excluded from missing-hook enforcement. (No spec
  states this.) (checks)
- no spec — Zone edits to `.fireweed/**` are rejected only because the file falls outside
  every module and matches no `worker_writable` pattern (CHK-010); there is no explicit
  `.fireweed` protected-path rule in the spec. (checks)
- no spec — Diagnostic dedup (identical `Code/Path/Message/Warning` collapsed) is a
  reconstruction-internal choice to keep reports stable; the spec says nothing about
  duplicates. (checks)

## Agent observations

### Invocation 1
Log: .fireweed/runs/20261119T100742-d1f4b6b358997bc4e075b4d1/drill-1.log

This is a large reconstruction. Let me look at a few more support files to get the details right.
Now I understand the shape. Let me write the reconstructed implementation, starting with shared helpers.
Now the Boundaries implementation:
Let me rewrite boundaries.go cleanly — that draft had structural problems.
Now the Zones implementation:
Important correction needed: `gitstore.Run` returns a non-nil error for any nonzero exit (`*exec.ExitError`), which must be treated as normal hook completion, not a launch failure. Let me add a shared helper and fix call sites.
History needs combined-diff semantics for merges. Let me experiment with git output formats in a scratch repo (outside the workspace only for the experiment — actually let me use a temp dir under /private/var/folders path):
Now I'll write the History implementation:
Now the history helper functions:
Now the Doctor — the biggest piece. I'll write it in sections. First, the main flow and helpers:
Now the scratch worktree and setup pieces:
Now the canary probes — recipe loading, rendering, and gate logic:
Now the canary probe execution, gate logic, and orchestration:
Let me restructure Boundaries to share the per-scope hook runner with Doctor probes. First the shared boundary-running core:
Now the baseline and canary probes:
Now the agent probes and guidance checks:
Now let's run the module evals against the workspace hook:
14 failures. Let me build a small standalone debug harness to inspect what Doctor does in the real fixture (I'll clean it up after):
Two distinct bugs found already (argv one-pass, H04 on rename targets). Let me first fix them, then debug the empty-report Doctor case:
The pipeline works; now debug the ID matching and EmptyReport cases with one-off probes:
ParseJUnit is fine. Let me add temporary debug prints in Evaluate/failingEvalProbe to see the real arguments:
The embedded recipe's template SPEC.md has no spec item — the item comes from the **canary baseline map** (CHK-053: templates + version-1 canary baseline). My render only used templates. Fixing:
Now re-run the module evals to see remaining failures:
Big remaining items. Let me add temporary debug prints and run the nested-template test to see what the scratch looks like:
Found it: the nested-template test expects `{{package}}` to equal the module name, and the real Go recipe needs a valid package identifier — both are satisfied by dashless throwaway names (`doctor<hex>`). Fixing my name generation and re-checking:
Let me re-run the remaining failing tests with current code:
RealGoRecipe is likely failing at the setup step (which needs `.fireweed/doctor-untracked`). Let me replicate the exact fixture in my debug harness:
Add debug prints to the doctor flow:
ready=true but the reloaded scratch still has only alpha. Check what render() actually created:
The fixture sets `modules: [packages/alpha]`, so the throwaways are intentionally invisible to `project.Load` — Doctor must synthesize their module entries instead. Implementing:
CHK-061 requires the scratch canary modules to be discoverable through FIREWEED_BIN's modules output. The real fixture restricts `modules: [packages/alpha]`, so Doctor must widen the scratch config's glob:
Fixing check names (recovery must include the canary), override C14 formats, and missing-enforcement gate emission:
Refining missing_id detection to match the canary's own test ID (Go func name or bare token):
Now the full eval suite:
One test remains: the malformed SPEC.md template must be validated **before** the canary baseline overlays it (which masks the defect). Fixing render order:
Now let me remove all debug artifacts and helper files:
Full suite again to confirm everything passes:
All 142 tests pass. Now let me review the git status and clean up remaining details:
All tests pass and the zones simplification held. Let me double-check the full suite once more, then write DRILL_NOTES.md:
Let me do a final verification pass: file inventory, vet, and the complete module eval once more with the sanctioned command:
Final verification with the exact task-provided command, then tidy up:
Done. The `checks` module was reconstructed from its spec, decisions, retained interfaces and evals, and all 142 retained eval cases pass via the project's eval hook (`hook-exit=0`).

**What was rebuilt** — only the disposable paths `packages/checks/internal/**` (17 files, `internal/engine/**`). No other module, the interface (`api.go`, `model/**`), evals, specs, or protected files were touched.

Key pieces:
- **junit.go / coverage.go** — JUnit parse (nested suites, ID mapping via `contracts.MatchesID`, C11 on malformed XML/unknown or missing IDs, failure/error/skip classification, message+body retention) and coverage/fail-first/pass gates (C10/P2-01/E01).
- **evaluate.go** — hook resolution (module override over global), shell/argv placeholder expansion with one-pass literal/quoted substitution, `FIREWEED_BIN`, process metadata retention, JUnit parsing, exit/report agreement rules.
- **boundaries.go** — global + per-module hooks with attribution, exact baseline suppression, reasoned opt-outs, X01/X02/X03/X04 handling with the bounded 4,096-byte UTF-8 stderr tail.
- **zones.go** — protected/spec/foreign/change-file rules, per-phase zone allow/reject including disposable-absent-from-base for evals, append-only add-only for impl, rename source+destination checks, worker_writable outside modules.
- **doctor.go** — static-error stop, detached scratch worktree below the common Git dir (current files copied in), setup hook with C14-stop, scratch-only probes, two throwaway canary modules from project templates + the version-1 canary baseline, baseline → plant → probe → recovery cycles for all six canaries, agent probes through `agent.wrap` with logs under `.fireweed/runs/<run>/`, C15 usage-limit classification, C16 AGENTS.md/playbook-override checks.
- **history.go** — revision‑range commits in oldest‑first topological order, worker commits from phase-assignment diffs (not subjects), squashed-change detection, H01–H04 using before/after recorded zone maps and projected `set` values; merge commits use combined `--cc` diffs.

**DRILL_NOTES.md** documents `## Contradicted Decisions` (none found) and `## Gaps` (12 items, each naming its nearest spec ID or "no spec", target module: `checks`). Most gaps are underspecified message wording, merge-parent choices, projected-`set` zone evaluation, canary gate/threshold semantics, and scratch-modules discovery.

Everything is untracked in the drill workspace (`packages/checks/internal/**` and `DRILL_NOTES.md`); no commits were made, and only the sanctioned eval command was run.

