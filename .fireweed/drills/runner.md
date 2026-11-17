---
agent: opencode
base_commit: f6f5997536ad8aeaa8e3b90a4a914f66d91a133f
date: "2026-11-17T07:10:31.162645Z"
model: opencode/deepseek-v4.1-flash#high
result: fail
---

## Own evals
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
- runner / Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail / RUN-058, RUN-059 / pass
- runner / Test_RUN_058_OwnCoverageRequiredAfterPassingCases / RUN-058 / pass
- runner / Test_RUN_059_DependentCoverageRequiredAfterPassingCases / RUN-059 / pass
- runner / Test_RUN_068_DependentOnlyFailuresEndDrill / RUN-068 / pass
- runner / Test_RUN_068_OwnEvalRetryOmitsDependentEvalOutput / RUN-068 / pass
- runner / Test_RUN_069_DrillObservationsIncludeStdoutAndRelativeLogPerAttempt / RUN-069 / pass
- runner / Test_RUN_060_RUN_061_DrillReportAndRetainedRef / RUN-060, RUN-061 / pass
- runner / Test_RUN_074_DrillResultAndReportIdentifySelectedCommand / RUN-074 / pass
- runner / Test_RUN_074_DrillResultAndReportIdentifySelectedCommand/default / RUN-074 / pass
- runner / Test_RUN_074_DrillResultAndReportIdentifySelectedCommand/explicit / RUN-074 / pass
- runner / Test_RUN_060_RUN_061_FailedDrillReportsFailuresAndRetainsCode / RUN-060, RUN-061 / pass
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
- runner / Test_RUN_033_OwnedResumeUsesApprovedUsageBudget / RUN-033 / pass
- runner / Test_RUN_030_UsageRetryPreservesVerificationFeedback / RUN-030 / pass
- runner / Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases / RUN-030, RUN-031 / pass
- runner / Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied / RUN-032 / pass
- runner / Test_RUN_034_UsageStopClassifierWrapping / RUN-034 / pass
- runner / Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure / RUN-035 / pass
- runner / Test_RUN_031_CancelledUsageWaitRetainsHeldWork / RUN-031 / pass

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
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration / APP-164 / fail

=== RUN   Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration
--- FAIL: Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration (1.77s)

- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/check_--json / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/show_alpha / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/new-change_alpha_new-work / APP-164 / pass
- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next / APP-164 / fail

=== RUN   Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next
    compatibility_cli_test.go:246: stop reason "usage wait stopped: completed 0s, next wait 5ms, budget 0s" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next (0.33s)

- app / Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run / APP-164 / fail

=== RUN   Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run
    compatibility_cli_test.go:246: stop reason "usage wait stopped: completed 0s, next wait 5ms, budget 0s" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run (0.35s)

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
- app / Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts / APP-141 / fail

=== RUN   Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts
--- FAIL: Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts (1.69s)

- app / Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain / APP-141 / fail

=== RUN   Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain
    drill_cli_test.go:344: plain failure stdout="" stderr="P01: own evals: failed case \"ALP-001_case\" for ALP-001. Correct the reported issue before retrying. Report: .fireweed/drills/alpha.md. Rebuilt ref: refs/fireweed/drills/alpha/2026-11-17.\n"; want empty stdout and E01
--- FAIL: Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain (0.83s)

- app / Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json / APP-141 / fail

=== RUN   Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json
    drill_cli_test.go:344: diagnostic code="P01"; want "E01"
--- FAIL: Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json (0.55s)

- app / Test_APP_141_DrillSetupFailureHasNoArtifacts / APP-141 / pass
- app / Test_APP_142_DrillCleanUsageStopPlainAndJSON / APP-142 / fail

=== RUN   Test_APP_142_DrillCleanUsageStopPlainAndJSON
    drill_cli_test.go:410: clean stop data=map[base_commit:58061d73485e99eb9895e3981029d9a3fcb44f92 module:alpha reason:usage wait stopped: completed 0s, next wait 5ms, budget 0s ref: report:.fireweed/drills/alpha.md result:stopped]
--- FAIL: Test_APP_142_DrillCleanUsageStopPlainAndJSON (1.56s)

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
- app / Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume / APP-076, APP-077, APP-079 / fail

=== RUN   Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume
    usage_cli_test.go:141: stop reason "usage wait stopped: completed 0s, next wait 5ms, budget 0s" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (4.15s)

- app / Test_APP_076_APP_077_NextPlainUsageStop / APP-076, APP-077 / fail

=== RUN   Test_APP_076_APP_077_NextPlainUsageStop
    usage_cli_test.go:240: stop reason "Stopped on held-quota-work: usage wait stopped: completed 0s, next wait 5ms, budget 0s. Use fireweed next to resume.\n" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_076_APP_077_NextPlainUsageStop (1.31s)

- app / Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume / APP-083, APP-084, APP-086 / fail

=== RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:254: stop reason "usage wait stopped: completed 5ms, next wait 5ms, budget 5ms" lacks a labeled completed wait of 5ms
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (7.21s)

- app / Test_APP_083_APP_084_RunPlainSharedUsageStop / APP-083, APP-084 / fail

=== RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:337: stop reason "Applied 1 changes.\nalpha\nStopped: usage wait stopped: completed 5ms, next wait 5ms, budget 5ms. Use fireweed run to resume.\n" lacks a labeled completed wait of 5ms
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (4.25s)

- app / Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop / APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/next / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/run / APP-077, APP-079, APP-084, APP-086 / pass

## Contradicted Decisions
None. The runner reconstruction follows the retained Decisions, including the
90-minute default invocation timeout and the length-labeled append-only agent
logs.

## Gaps
- RUN-063, module runner: the exact framing of length-labeled log records is not
  specified beyond preserving both streams and appending. The implementation
  records `stdout <bytes>` followed by `stderr <bytes>` per invocation.
- RUN-003, module runner: shell quote-state handling for escaped quotes in mixed
  command templates is outside the stated contract. The implementation tracks
  single and double quotes and backslash escapes when deciding whether to quote
  substituted values.
- no spec, module runner: the disposable internal package decomposition
  (`invoke`, `next`, `run`, `drill`, `stop`) is not specified; only `api.go` and
  `model/**` are retained public boundaries.

## Agent observations

### Invocation 1
Log: .fireweed/runs/20261117T071031-edb15b349c113463d682c8ae/drill-1.log

Reconstructed `packages/runner/internal/**` and added `DRILL_NOTES.md`.

Implemented:
- `internal/invoke` — command selection, wrappers, placeholders, prompt files, timeouts, usage-limit detection, process-group handling, and append-only invocation logs.
- `internal/next` — lifecycle driving, approvals, retries, commit recovery, blocking, audit isolation/copy-back, usage waits/budgets, diagnostics, run resumption, and stop classification.
- `internal/run` — bounded/unbounded queue execution, shared usage/time budgets, blocked-change continuation, and result flags.
- `internal/drill` — preflight, scratch/input workspaces, filtered Git inputs, reconstruction retries, eval verification, dependent evaluation, reports, refs, and cleanup.
- `internal/stop` — shared clean usage/time/blocked error classifications.

Validation:
- Full `packages/runner/evals` suite passes.
- The supplied eval hook command passes:
  - `ok github.com/lurkshark/fireweed/packages/runner/evals`


## Failures
app: failed case "Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume" for APP-076
Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration: === RUN   Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration
--- FAIL: Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration (1.77s)
Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next: === RUN   Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next
    compatibility_cli_test.go:246: stop reason "usage wait stopped: completed 0s, next wait 5ms, budget 0s" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/next (0.33s)
Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run: === RUN   Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run
    compatibility_cli_test.go:246: stop reason "usage wait stopped: completed 0s, next wait 5ms, budget 0s" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_164_HeldRunnerAdmissionUsesApprovedConfiguration/run (0.35s)
Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts: === RUN   Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts
--- FAIL: Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts (1.69s)
Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain: === RUN   Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain
    drill_cli_test.go:344: plain failure stdout="" stderr="P01: own evals: failed case \"ALP-001_case\" for ALP-001. Correct the reported issue before retrying. Report: .fireweed/drills/alpha.md. Rebuilt ref: refs/fireweed/drills/alpha/2026-11-17.\n"; want empty stdout and E01
--- FAIL: Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/own-plain (0.83s)
Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json: === RUN   Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json
    drill_cli_test.go:344: diagnostic code="P01"; want "E01"
--- FAIL: Test_APP_141_DrillOwnAndDependentFailuresRetainArtifacts/dependent-json (0.55s)
Test_APP_142_DrillCleanUsageStopPlainAndJSON: === RUN   Test_APP_142_DrillCleanUsageStopPlainAndJSON
    drill_cli_test.go:410: clean stop data=map[base_commit:58061d73485e99eb9895e3981029d9a3fcb44f92 module:alpha reason:usage wait stopped: completed 0s, next wait 5ms, budget 0s ref: report:.fireweed/drills/alpha.md result:stopped]
--- FAIL: Test_APP_142_DrillCleanUsageStopPlainAndJSON (1.56s)
Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume: === RUN   Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume
    usage_cli_test.go:141: stop reason "usage wait stopped: completed 0s, next wait 5ms, budget 0s" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (4.15s)
Test_APP_076_APP_077_NextPlainUsageStop: === RUN   Test_APP_076_APP_077_NextPlainUsageStop
    usage_cli_test.go:240: stop reason "Stopped on held-quota-work: usage wait stopped: completed 0s, next wait 5ms, budget 0s. Use fireweed next to resume.\n" lacks a labeled completed wait of 0s
--- FAIL: Test_APP_076_APP_077_NextPlainUsageStop (1.31s)
Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume: === RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:254: stop reason "usage wait stopped: completed 5ms, next wait 5ms, budget 5ms" lacks a labeled completed wait of 5ms
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (7.21s)
Test_APP_083_APP_084_RunPlainSharedUsageStop: === RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:337: stop reason "Applied 1 changes.\nalpha\nStopped: usage wait stopped: completed 5ms, next wait 5ms, budget 5ms. Use fireweed run to resume.\n" lacks a labeled completed wait of 5ms
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (4.25s)
?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	269.459s

exit status 1

