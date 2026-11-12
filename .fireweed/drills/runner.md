---
date: 2026-11-10T19:24:09.383763Z
base_commit: 5b59f35f27d4e53e4061a367e1d8d7f02e62ec1a
agent: codex
model: gpt-6.1-sol
result: pass
---

## Own evals
- runner | Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend | RUN-063 | pass
- runner | Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/failure | RUN-063 | pass
- runner | Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/timeout | RUN-063 | pass
- runner | Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/cancellation | RUN-063 | pass
- runner | Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/append | RUN-063 | pass
- runner | Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput | RUN-064 | pass
- runner | Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries | RUN-065 | pass
- runner | Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries/usage_once | RUN-065 | pass
- runner | Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries/agent_retry | RUN-065 | pass
- runner | Test_RUN_066_InvokeAgentLogContextDefaultsAndExplicitFields | RUN-066 | pass
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
- runner | Test_RUN_052_RUN_053_RestrictedInputsAndIndependentGit | RUN-052, RUN-053 | pass
- runner | Test_RUN_054_PromptLetsWorkerRunOwnEval | RUN-054 | pass
- runner | Test_RUN_052_RUN_053_ProjectProtectedDisposableIsRetainedInputAndVerifierFile | RUN-052, RUN-053 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_api | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/delete_eval | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_preserved | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_protected_dependency_interface | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/change_setup | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs/delete_setup | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions/add_foreign | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions/add_other | RUN-055 | pass
- runner | Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions/add_project_protected | RUN-055 | pass
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
- runner | Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce | RUN-003 | pass
- runner | Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/worker_shell | RUN-003 | pass
- runner | Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/worker_argv | RUN-003 | pass
- runner | Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/audit_shell | RUN-003 | pass
- runner | Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce/audit_argv | RUN-003 | pass
- runner | Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback | RUN-003, RUN-067 | pass
- runner | Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback/worker | RUN-003, RUN-067 | pass
- runner | Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback/audit | RUN-003, RUN-067 | pass
- runner | Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback/audit_fallback | RUN-003, RUN-067 | pass
- runner | Test_RUN_067_WrapperOnlyPlaceholderCreatesPromptFile | RUN-067 | pass
- runner | Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure | RUN-067 | pass
- runner | Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/success | RUN-067 | pass
- runner | Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/nonzero | RUN-067 | pass
- runner | Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/timeout | RUN-067 | pass
- runner | Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/cancellation | RUN-067 | pass
- runner | Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure/launch_failure | RUN-067 | pass
- runner | Test_RUN_067_UnselectedTemplateDoesNotRequirePromptFileIO | RUN-067 | pass
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
?   	github.com/lurkshark/fireweed/packages/runner/internal/support	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	277.435s


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
ok  	github.com/lurkshark/fireweed/packages/app/evals	261.174s


## Contradicted Decisions
None observed.
## Gaps
- RUN-003, module runner. The shell template grammar leaves prequoted placeholders unspecified. The reconstruction shell-quotes each replacement, using the same convention as the eval hook.
- RUN-014, module workflow. Done verifies internally and accepts no prior Verify result. The reconstruction verifies before calling Done, so successful phases can run their eval hooks twice.
- RUN-041, module runner. The spec requires `IsApprovalRequired` to recognize the refusal. The app evals for APP-075, APP-077 and APP-082 also require P01 for the refusal. The reconstruction uses P01 and preserves the typed approval error. P04 is the warning emitted by `workflow.Start`.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt the runner under `packages/runner/internal`. All 162 retained evals passed, with all 67 runner spec IDs represented. `go vet` and the boundary hook also passed.

The controller’s `fireweed verify` returned P02 because the drill workspace had no held change. The supplied runner hook produced the passing [JUnit report](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/.fireweed/runs/drill/report.xml).

[DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/DRILL_NOTES.md) records two gaps and no contradicted Decisions.

Reading additional input from stdin...
2026-11-10T19:24:11.819736Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1c6ea-fbf2-7a82-aba5-c519a680b209
--------
user
# Rebuild runner

## Target spec and Decisions

---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
id_prefix: RUN
interface:
    - api.go
    - model/**
name: runner
---

# runner

## Specs

### RUN-001
InvokeAgent MUST select agent.command for the worker role and audit.command for the audit role, falling back to agent.command when audit.command is unset.

### RUN-002
InvokeAgent MUST run the selected command from Snapshot.Root.

### RUN-003
InvokeAgent MUST expand {prompt}, {prompt_file} and {git_common_dir} once in command templates and wrapper argv. Commands MUST accept shell strings or argv, quoting shell substitutions and preserving argv values literally.

### RUN-004
InvokeAgent MUST prepend the configured agent.wrap argv prefix for both worker and audit commands.

### RUN-005
InvokeAgent MUST set the command's FIREWEED_BIN environment variable to the request Bin.

### RUN-006
InvokeAgent MUST retain the command's stdout, stderr and exit code in AgentResult.

### RUN-007
InvokeAgent MUST use a positive request Timeout in preference to the configured timeout, using the configured timeout when the request is unset. It MUST stop the command's process group when that bound expires or its context is cancelled.

### RUN-008
InvokeAgent MUST set UsageLimit true exactly when a nonzero command exit has a configured limit_patterns match in its final 20 output lines.

### RUN-009
InvokeAgent MUST return command-launch, timeout or cancellation errors without discarding captured output.

### RUN-010
Next MUST return Empty true without invoking agents when workflow.Eligible
returns no changes, including an empty or entirely waiting queue.

### RUN-011
When workflow.Eligible returns changes, Next MUST select its first result and
start it through workflow.Start, propagating start failures without invoking
an agent.

### RUN-012
Unless AllowUnmerged is true, Next MUST reject a selected change absent from
both the local and shared base before claiming it or assigning IDs.

### RUN-013
For each phase, Next MUST invoke the worker role for evals and impl from the held checkout and the audit role from its restricted input workspace, using approved agent configuration and request Bin and Timeout.
Worker prompts MUST include full workflow context.
Audit prompts MUST include the change, projected spec, eval contents and audit playbook without implementation contents.

### RUN-014
After each successful agent invocation, Next MUST verify the current phase and
call workflow.Done only when verification passes, continuing until applied.
It MUST return the applied state and the ordered phase invocation results.

### RUN-015
For a non-usage-limit agent invocation failure, Next MUST return an error and
captured invocation results without advancing the phase or releasing its held
claim.

### RUN-016
Run without a positive Max MUST continue single-change execution until no
eligible changes remain or its For bound stops it, returning applied NextResults
in execution order.

### RUN-017
With a positive Max, Run MUST apply at most Max changes per invocation. A later
Run invocation MUST continue the remaining queue without reapplying archived
changes.

### RUN-018
On successful return, Run MUST set Empty when no queued changes remain and set
Limited exactly when a positive Max stopped it with queued work remaining.

### RUN-019
Run MUST supply its Root, Bin, AllowUnmerged and Timeout to single-change
execution for every change it processes.

### RUN-020
If single-change execution fails while the calling worktree still holds the
change, Run MUST return the error with its cause and already-applied results
retained, leaving later changes unstarted.

### RUN-021
Next MUST permit at most the configured attempts count of agent invocations
for a phase after verification failures, defaulting to two when unset. Each
phase MUST have its own attempt allowance.

### RUN-022
When phase verification fails and an attempt remains, Next MUST retain the
earlier work and retry that phase. The retry prompt MUST include the last
verification output.

### RUN-023
For exhausted phase verification with the worktree's approved claim and an
unchanged HEAD, Next MUST save and block the step through the lifecycle,
leaving a clean unclaimed checkout and a state-only block commit.

### RUN-024
After that blocking succeeds, Next MUST return the blocked lifecycle State and
ordered captured invocations without changing the run, assigned IDs or current
phase. Its error MUST retain the last verification cause and identify the saved
attempt ref.

### RUN-025
When single-change execution returns an error after successfully saving and
blocking its change with a clean unclaimed checkout, Run MUST continue processing
eligible changes. Run MUST exclude that blocked change from Applied.

### RUN-026
If a change was blocked during Run, Run MUST return an error when eligible work
is exhausted or Max is reached. That error MUST retain a blocked change's
verification cause and saved attempt ref.

### RUN-027
When an evals or impl worker creates commits, Next MUST soft-reset the held checkout's HEAD to its phase base after invocation and before verification or returning an invocation failure, retaining the step's working-tree and index work.

### RUN-028
For each such recovery, Next MUST return a P05 warning in State.Diagnostics
identifying the phase and that work was retained. It MUST preserve these warnings
across phase advancement and on an invocation-failure return.

### RUN-029
Next MUST read limit_wait and limit_budget as durations, defaulting to 30m and
4h when unset. It MUST reject malformed durations, a nonpositive wait or negative
budget with C01 before claiming a fresh change or invoking an agent.

### RUN-030
Subject to the usage budget and Run's For bound, an AgentResult with UsageLimit
true MUST make Next wait limit_wait and repeat the same phase attempt without
spending a failed attempt. It MUST retain earlier step work and captured
invocations, keeping existing verification feedback.

### RUN-031
When another usage wait would exceed limit_budget, Next MUST return a clean
usage-stop error identifying completed wait, next wait and budget. It MUST retain
the held State, work and ordered Steps without advancing, blocking, committing
or releasing the claim.

### RUN-032
Run MUST account for total usage wait across phases and changes in one
invocation. On a clean usage stop it MUST retain already-applied results and
propagate the stop classification.

### RUN-033
When the calling worktree already holds a queued step, Next MUST resume that
owned State using its approved snapshot without a fresh Start, retaining its
run, phase, assigned IDs and work. Existing approval and unapproved-work bounds
MUST remain enforced for starting another change.

### RUN-034
IsUsageStop MUST recognize clean usage-stop errors through wrapping and return
false for nil or ordinary failures.

### RUN-035
If Run blocked a change before a usage stop, it MUST return failure with the
blocking cause/ref and stop reason retained. IsUsageStop MUST return false for
that combined failure so the caller cannot report it as a clean stop.

### RUN-036
Run MUST treat For zero as unbounded and positive For as one elapsed-time
budget for the whole invocation, including agent and verification time across
phases and changes.

### RUN-037
When its For bound stops execution, Run MUST return a nonnil clean time-stop
error naming the limit, with already-applied results retained. A step interrupted
during agent invocation or verification MUST keep its work, claim, run, phase
and assigned IDs without advancing, blocking, committing or releasing that step.

### RUN-038
When a usage wait fits its usage budget but would pass Run's For deadline,
execution MUST return a clean time stop before waiting. It MUST preserve the
held step's work and identity.

### RUN-039
IsTimeStop MUST recognize clean time-stop errors through wrapping. It MUST
return false for nil, usage stops, ordinary failures and parent-context
cancellation.

### RUN-040
If Run blocked a change before a time stop, it MUST return ordinary failure
retaining the blocking cause/ref and time-stop reason. IsTimeStop MUST return
false for that combined failure.

### RUN-041
Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

### RUN-042
IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.

### RUN-043
On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to RUN-044.
It MUST select cases by lexical module name and then report order, excluding passing and skipped cases.

### RUN-044
The failed-case detail section MUST fit within 16,384 Unicode code points, including labels and truncation markers.
Each Case.Failure excerpt MUST contain at most its first 4,096 code points.
Shortened diagnostic text or omitted cases MUST be marked.

### RUN-045
Next MUST run every automatic audit invocation in a fresh temporary workspace containing only the current change, projected module SPEC.md and module eval files at project-relative paths, plus independent Git metadata.
It MUST omit the held checkout's other files, including implementation, interfaces, hooks and configuration.

### RUN-046
The audit workspace's Git metadata MUST contain none of the held repository's history, refs or saved attempts.
Audit-agent commits within this workspace MUST remain absent from the held checkout's history and refs.

### RUN-047
Next MUST accept only edits to the Audit section of the copied current change.
Edits elsewhere in that change, edits to copied specs or evals, or added or deleted files outside Git metadata MUST reject the submission without copying any workspace edits to the held checkout.

### RUN-048
When audit workspace edits are permitted, Next MUST copy its Audit section back before retrying or returning, including agent failures and time or usage stops.
Copying audit work MUST not itself advance a phase or replace other held change contents.

### RUN-049
Next MUST remove every temporary audit workspace after invocation, including after invalid submissions, agent failures and time or usage stops.

### RUN-050
Drill MUST reject an unknown module, dirty invoking checkout or any live shared change claim before setup or agent invocation, leaving the invoking checkout and refs unchanged.

### RUN-051
Drill MUST create a detached scratch worktree at the invoking HEAD under the common Git directory's fireweed/worktrees directory and run the configured setup hook there with FIREWEED_BIN set to request Bin. Setup failure MUST prevent agent invocation and return its captured failure.

### RUN-052
Drill MUST delete only target-module files classified disposable by the approved zone precedence before reconstruction. It MUST retain every other target-module file and every other module's files in the verification worktree.

### RUN-053
Drill MUST invoke reconstruction in a separate input workspace holding the invoking HEAD's tracked files except the target's disposable files, every direct and transitive dependent module's files, the changes directory and drill reports. Its Git metadata MUST hold one fresh commit and none of the invoking repository's history, refs or saved attempts. Drill MUST run the configured setup hook there before agent invocation.

### RUN-054
The reconstruction prompt MUST include the target spec and Decisions, identify its retained interfaces and evals, and require rebuilding only disposable paths without reading outside the input workspace. It MUST give a command that runs the target's eval hook in that workspace. It MUST request separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md, each gap naming the nearest spec ID or "no spec" and a target module.

### RUN-055
Drill MUST reject an attempt that modified or deleted any input file present after setup, before importing code. It MUST import only the target's rebuilt disposable files into the verification worktree, ignoring other files the worker or its eval runs created, without changing retained files or the invoking implementation.

### RUN-056
Drill MUST use the configured worker command, wrapper and timeout through InvokeAgent, retaining ordered AgentResults in Steps. Non-usage failures and verification retries MUST share the configured attempts bound, retain rebuilt work and include the preceding failure in retry prompts.

### RUN-057
Drill MUST treat usage-limit responses as the same attempt, applying configured limit_wait and limit_budget without spending a failed attempt. An exhausted usage budget MUST return an IsUsageStop error with Result stopped; cancellation or timeout MUST return an error and stop reconstruction.

### RUN-058
After accepted reconstruction, Drill MUST run the target module's eval hook and require passing enforced cases and coverage through checks. It MUST retain the complete returned own report in Own, including failed testcase details when verification fails.

### RUN-059
Drill MUST evaluate every direct and transitive dependent in lexical module-name order regardless of verify_dependents, retaining each report in Dependents. Every dependent MUST satisfy passing enforced cases and coverage for a passing drill, even when own or another dependent's evals fail.

### RUN-060

After reconstruction has been invoked, Drill MUST write the current invocation's report to .fireweed/drills/<module>.md, replacing any previous report there, with date, base_commit, agent, model and result frontmatter and Own evals, Dependent evals, Contradicted Decisions and Gaps sections. Eval entries MUST identify module, testcase, spec IDs and result. Agent observations and verification failures MUST remain visible, with each gap identifying a nearest spec ID or "no spec" and target module.

### RUN-061
After accepted reconstruction, Drill MUST retain the rebuilt verification tree at refs/fireweed/drills/<module>/<date>, returning that ref without moving the invoking HEAD. A passing result MUST require accepted reconstruction and passing own and all dependent evals; failure MUST return a nonnil error with Result fail.

### RUN-062
Drill MUST remove its verification worktree and input workspace on success, agent or verification failure, invalid edits, usage stop and cancellation. Cleanup failure MUST be returned as an error rather than reported as a passing drill.

### RUN-063
InvokeAgent MUST persist the invocation's complete stdout and stderr, including partial output on command failure, timeout or cancellation, at .fireweed/runs/<run>/<phase>-<attempt>.log under its log project root. Repeated invocations at the same path MUST append without discarding earlier records. AgentResult.LogPath MUST identify that file by an absolute path.

### RUN-064
Next MUST assign each invocation the held change's run, current phase and one-based phase attempt, writing its log under the invoking project's .fireweed/runs directory. Audit logs MUST stay in that invoking project, outside the restricted audit input workspace.

### RUN-065
Drill MUST allocate one log run identity per invocation and reuse it across reconstruction retries and usage waits. Its logs MUST use phase drill, the current one-based reconstruction attempt and the invoking project's root, outside scratch and reconstruction input workspaces.

### RUN-066
InvokeAgent MUST accept a log project root, run, phase and positive one-based attempt in AgentRequest. Unset fields MUST default to Snapshot.Root, a fresh run identifier, request Role and attempt 1, respectively.

### RUN-067
InvokeAgent MUST create one temporary file containing the exact request Prompt only when {prompt_file} occurs in the selected command or wrapper templates, and expand that placeholder to its absolute path. The file MUST remain available through command execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Decisions

[0075-bounded-agent-invocation] An unset configured timeout defaults to 90 minutes. limit_patterns are
case-insensitive literal substrings, using the final 20 lines of stdout followed
by stderr with terminal newline separators trimmed.

- agent-invocation-logs: Agent logs append length-labeled stdout and stderr records after each invocation, preserving both streams' bytes even without final newlines. A log write error is returned alongside any agent error while the captured result remains available.


## Retained interfaces and files

packages/runner/SPEC.md
packages/runner/api.go
packages/runner/model/model.go

## Retained evals

packages/runner/evals/.gitkeep
packages/runner/evals/agent_logs_test.go
packages/runner/evals/approval_test.go
packages/runner/evals/audit_workspace_test.go
packages/runner/evals/commit_recovery_test.go
packages/runner/evals/drill_test.go
packages/runner/evals/invoke_test.go
packages/runner/evals/next_test.go
packages/runner/evals/prompt_file_test.go
packages/runner/evals/retry_failure_details_test.go
packages/runner/evals/run_test.go
packages/runner/evals/time_limit_test.go
packages/runner/evals/usage_test.go

Rebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.

Run the target's eval hook from the workspace root as often as needed.

~~~sh
mkdir -p .fireweed/runs/drill && sh -c 'go run .fireweed/hooks.go eval '"'"'packages/runner'"'"' '"'"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/.fireweed/runs/drill/report.xml'"'"''
~~~

Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module.

codex
I’ll read the rebuild guardrails and retained runner interfaces, then implement the runner in its disposable paths and run the supplied eval hook.
exec
/bin/zsh -lc "pwd; rg --files --hidden -g '"'!**/.git/**'"' -g '"'!**/rebuild-v1/**'"' | head -130; printf '\\n--- SYSTEM ---\\n'; cat SYSTEM.md; printf '\\n--- PROGRESS ---\\n'; cat docs/rebuild/PROGRESS.md; printf '\\n--- runner interface ---\\n'; cat packages/runner/api.go packages/runner/model/model.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 615ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
main.go
AGENTS.md
.fireweed/canaries/canaries.yaml
.fireweed/templates/module/SPEC.md
.fireweed/templates/module/api.go.tmpl
.fireweed/templates/module/internal/.gitkeep
.fireweed/templates/module/evals/.gitkeep
.fireweed/recipe.md
.fireweed/hooks.go
.fireweed/setup.sh
SYSTEM.md
packages/gitstore/SPEC.md
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/gitstore/internal/local/local.go
packages/gitstore/evals/.gitkeep
packages/gitstore/evals/gitstore_test.go
packages/assets/SPEC.md
packages/assets/api.go
packages/assets/model/model.go
packages/assets/internal/content/content.go
packages/assets/internal/.gitkeep
packages/assets/internal/recipe/recipe.go
packages/assets/internal/recipe/files/typescript-pnpm-index.ts.tmpl
packages/assets/internal/recipe/files/typescript-pnpm-hooks.mjs
packages/assets/internal/recipe/files/typescript-pnpm-package.json.tmpl
packages/assets/internal/recipe/files/typescript-pnpm-canaries.yaml
packages/assets/internal/recipe/files/python-uv-pyproject.toml.tmpl
packages/assets/internal/recipe/files/SPEC.md
packages/assets/internal/recipe/files/python-uv-init.py.tmpl
packages/assets/internal/recipe/files/api.go.tmpl
packages/assets/internal/recipe/files/python-uv-SPEC.md
packages/assets/internal/recipe/files/typescript-pnpm.md
packages/assets/internal/recipe/files/python-uv.md
packages/assets/internal/recipe/files/typescript-pnpm-SPEC.md
packages/assets/internal/recipe/files/python-uv-hooks.py
packages/assets/internal/recipe/files/canaries.yaml
packages/assets/internal/recipe/files/python-uv-canaries.yaml
packages/assets/internal/recipe/files/hooks.go
packages/assets/internal/recipe/files/go.md
packages/assets/evals/assets_test.go
packages/assets/evals/go_recipe_id_guidance_test.go
packages/assets/evals/status_warning_explanations_test.go
packages/assets/evals/eval_fix_playbooks_test.go
packages/assets/evals/go_recipe_test.go
packages/assets/evals/additional_recipes_test.go
packages/assets/evals/regenerate_warning_explanation_test.go
packages/assets/evals/mixed_eval_fix_explanation_test.go
packages/assets/evals/.gitkeep
packages/assets/evals/history_explanations_test.go
packages/assets/evals/stub_handoff_test.go
packages/assets/evals/decision_notes_playbooks_test.go
packages/assets/evals/pr_body_artifact_explanation_test.go
packages/workflow/SPEC.md
packages/workflow/api.go
README.md
LICENSE
go.mod
fireweed-design-v1.md
fireweed.yaml
Makefile
go.sum
packages/workflow/model/model.go
packages/contracts/SPEC.md
packages/contracts/api.go
packages/workflow/evals/costs_test.go
packages/workflow/evals/pr_body_test.go
packages/workflow/evals/lifecycle_test.go
packages/workflow/evals/maintenance_start_test.go
packages/workflow/evals/gc_test.go
packages/workflow/evals/selector_audit_validation_test.go
packages/workflow/evals/current_absence_test.go
packages/workflow/evals/eval_fix_lifecycle_test.go
packages/workflow/evals/abandon_test.go
packages/workflow/evals/eligible_test.go
packages/workflow/evals/resume_test.go
packages/workflow/evals/status_test.go
packages/workflow/evals/.gitkeep
packages/workflow/evals/block_test.go
packages/workflow/evals/regenerate_start_test.go
packages/workflow/evals/commits_test.go
packages/workflow/evals/release_test.go
.gitignore
docs/rebuild/ACCEPTANCE.md
docs/rebuild/DECISIONS.md
docs/rebuild/PROGRESS.md
docs/rebuild/GOAL.md
packages/contracts/model/model.go
packages/contracts/evals/contracts_test.go
packages/contracts/evals/.gitkeep
packages/workflow/internal/.gitkeep
packages/workflow/internal/lifecycle/status_history.go
packages/workflow/internal/lifecycle/resume_stub.go
packages/workflow/internal/lifecycle/eligible_stub.go
packages/workflow/internal/lifecycle/stub.go
packages/workflow/internal/lifecycle/block_stub.go
packages/workflow/internal/lifecycle/status_stub.go
packages/workflow/internal/lifecycle/status_costs.go
packages/workflow/internal/lifecycle/commits_stub.go
packages/workflow/internal/lifecycle/endpoint_config_error_stub.go
packages/workflow/internal/lifecycle/release_stub.go
packages/workflow/internal/lifecycle/regenerate.go
packages/workflow/internal/lifecycle/pr_body_stub.go
packages/workflow/internal/lifecycle/costs_stub.go
packages/workflow/internal/lifecycle/abandon_stub.go
packages/workflow/internal/lifecycle/gc.go
packages/checks/SPEC.md
packages/checks/api.go
packages/runner/SPEC.md
packages/runner/api.go
packages/project/SPEC.md
packages/project/api.go
packages/checks/model/model.go
packages/runner/model/model.go
packages/contracts/internal/primitives/coded_error.go
packages/contracts/internal/primitives/primitives.go
packages/contracts/internal/.gitkeep
packages/project/model/model.go
packages/runner/evals/audit_workspace_test.go
packages/runner/evals/commit_recovery_test.go
packages/runner/evals/.gitkeep
packages/runner/evals/next_test.go
packages/runner/evals/usage_test.go
packages/runner/evals/approval_test.go
packages/runner/evals/prompt_file_test.go
packages/runner/evals/time_limit_test.go
packages/runner/evals/invoke_test.go
packages/runner/evals/retry_failure_details_test.go
packages/runner/evals/agent_logs_test.go
packages/runner/evals/run_test.go

--- SYSTEM ---
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

--- PROGRESS ---
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
`runner` and `app` need first native reports and retained refs.

## Next step

Run the owner-approved `runner` drill with the verified controller, commit
the native report, retain the proof ref, and stop for owner review before
`app`. The remaining drill order is `runner`, then `app` in the
[regeneration goal](GOAL.md).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.

--- runner interface ---
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

exec
/bin/zsh -lc 'for f in packages/runner/evals/*test.go; do echo =====$f; rg -n '"'"'^func |assert|Contains'"\\(|want |fatal' \""'$f"; done' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc 'cat packages/{workflow,project,checks,gitstore,contracts}/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "ls -la .fireweed .fireweed/bin packages/runner; cat fireweed.yaml .fireweed/setup.sh; rg -n 'runner|0075' docs/rebuild/DECISIONS.md; cat packages/{workflow,project,checks,gitstore,contracts,assets}/api.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 751ms:
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
	// Args is the complete argv. Args[0] names the executable and is resolved
	// through PATH when it contains no slash. Run never prepends git.
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

 exited 1 in 755ms:
zsh:1: ====packages/runner/evals/agent_logs_test.go not found

 succeeded in 779ms:
.fireweed:
total 24
drwxr-xr-x@  8 joel  staff   256 Nov 10 11:24 .
drwx------@ 17 joel  staff   544 Nov 10 11:24 ..
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:24 bin
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:24 canaries
-rw-r--r--@  1 joel  staff  3993 Nov 10 11:24 hooks.go
-rw-r--r--@  1 joel  staff  1502 Nov 10 11:24 recipe.md
-rw-r--r--@  1 joel  staff   470 Nov 10 11:24 setup.sh
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:24 templates

.fireweed/bin:
total 13512
drwxr-xr-x@ 3 joel  staff       96 Nov 10 11:24 .
drwxr-xr-x@ 8 joel  staff      256 Nov 10 11:24 ..
-rwxr-xr-x@ 1 joel  staff  6914850 Nov 10 11:24 fireweed

packages/runner:
total 48
drwxr-xr-x@  6 joel  staff    192 Nov 10 11:24 .
drwxr-xr-x@  9 joel  staff    288 Nov 10 11:24 ..
-rw-r--r--@  1 joel  staff  16717 Nov 10 11:24 SPEC.md
-rw-r--r--@  1 joel  staff   1340 Nov 10 11:24 api.go
drwxr-xr-x@ 15 joel  staff    480 Nov 10 11:24 evals
drwxr-xr-x@  3 joel  staff     96 Nov 10 11:24 model
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

exec
/bin/zsh -lc "for f in packages/runner/evals/*test.go; do printf '\\n===== %s\\n' \""'$f"; rg -n '"'"'^func |Contains'"\\(|want |Fatal' \""'$f"; done' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "rg -n '"'^func |claims/|base_branch|remote'"' packages/workflow/internal/lifecycle/{stub,resume_stub,eligible_stub,block_stub}.go; rg -n '"'^func |Context|timeout'"' packages/gitstore/internal/local/local.go; cat .fireweed/hooks.go; .fireweed/bin/fireweed verify --help; git status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 633ms:

===== packages/runner/evals/agent_logs_test.go
16:func assertLogFile(t *testing.T, path, root, run, phase string, attempt int, streams ...string) []byte {
18:	want := filepath.Join(root, ".fireweed", "runs", run, fmt.Sprintf("%s-%d.log", phase, attempt))
19:	if path != want || !filepath.IsAbs(path) {
20:		t.Errorf("log path = %q, want absolute %q", path, want)
31:		if !strings.Contains(string(data), stream) {
38:func Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend(t *testing.T) {
96:					if !strings.Contains(string(data), earlier) {
105:func Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput(t *testing.T) {
109:		t.Fatalf("Next fixture = %#v, %v", result, err)
111:	for i, want := range []struct {
132:func Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries(t *testing.T) {
139:				t.Fatalf("Drill fixture = %#v, %v", result, err)
184:func Test_RUN_066_InvokeAgentLogContextDefaultsAndExplicitFields(t *testing.T) {
191:			t.Fatalf("default invocation: %v", err)
218:		t.Fatalf("explicit invocation: %v", err)

===== packages/runner/evals/approval_test.go
10:func Test_RUN_041_UnapprovedChangeRequiresApproval(t *testing.T) {
14:		t.Fatal("Next accepted a change absent from the local and shared base")
22:func Test_RUN_042_OtherNextErrorsDoNotRequireApproval(t *testing.T) {
35:			t.Fatal("expected dirty-worktree start failure")
47:			t.Fatalf("expected agent exit 9, got %#v, %v", result, err)

===== packages/runner/evals/audit_workspace_test.go
18:func auditGitObject(root, oid string) (string, error) {
27:func auditAction(t *testing.T, f nextFixture, action string) {
31:	if !strings.Contains(script, anchor) {
32:		t.Fatal("auditor insertion point missing")
37:func auditBeforeEdit(t *testing.T, f nextFixture, action string) {
41:	if !strings.Contains(script, anchor) {
42:		t.Fatal("auditor input inspection point missing")
47:func auditWorkspaceGone(t *testing.T, f nextFixture) {
59:func Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit(t *testing.T) {
73:		t.Fatalf("isolated audit lifecycle = %#v, %v", r, err)
81:	want := []string{"./changes/0001-next.md", "./packages/alpha/SPEC.md", "./packages/alpha/evals/existing_test.go", "./packages/alpha/evals/next_test.go"}
83:		t.Errorf("audit workspace files = %q, want %q", files, want)
88:	if refs := nextRead(t, f.logDir, "audit.refs"); strings.Contains(refs, "sentinel") || strings.Contains(refs, "fireweed/") {
92:		if !strings.Contains(nextRead(t, f.logDir, item.file), item.marker) {
96:	if strings.Contains(nextRead(t, f.logDir, "audit.change"), "No findings") {
102:func Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration(t *testing.T) {
110:		t.Fatal("fixture did not change the audit command")
131:		t.Fatalf("approved audit retry = %#v, %v", r, err)
150:func Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry(t *testing.T) {
179:		t.Fatalf("audit verification retry = %#v, %v", r, err)
183:		if !strings.Contains(got, "No findings\nChecked IDs: ALP-002") || !strings.Contains(got, "One new value is needed.") {
188:	if !strings.Contains(archived, "No findings; checked ALP-002") {
194:func Test_RUN_046_AuditCommitDoesNotEnterHeldGit(t *testing.T) {
201:		t.Fatalf("audit commit lifecycle = %#v, %v", r, err)
205:		t.Fatal("auditor did not commit")
210:	if out := nextGit(t, f.root, "for-each-ref", "--format=%(objectname)"); strings.Contains(out, auditOID) {
219:func Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported(t *testing.T) {
243:		t.Fatalf("first audit failure = %#v, %v", first, firstErr)
249:		t.Fatalf("audit before Provenance lifecycle = %#v, %v", r, err)
251:	for i, want := range []struct {
256:			t.Errorf("step %d = %#v, want %#v", i, step, want)
264:	if !strings.Contains(original, "## Provenance\n") || strings.Count(candidate, inserted) != 1 ||
273:		t.Errorf("Audit copying changed raw Provenance before successful Done: got %q, want %q", retryProvenance, beforeProvenance)
275:	if audit, ok := contracts.Section(retryInput, "Audit"); !ok || !strings.Contains(audit, "No findings\nChecked IDs: ALP-002") {
282:		if got != want {
283:			t.Errorf("held %s changed outside Audit: got %q, want %q", section, got, want)
307:	if audit, ok := contracts.Section(archived, "Audit"); !ok || !strings.Contains(audit, "No findings; checked ALP-002") {
313:func Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces(t *testing.T) {
318:		t.Fatalf("two-change audit lifecycle: applied=%d error=%v", len(r.Applied), err)
328:		t.Fatalf("audit invocation roots = %#v", cwds)
340:func Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically(t *testing.T) {
361:				t.Fatalf("rejected edit removed held change: %v", readErr)
364:			if strings.Contains(held, "No findings") || strings.Contains(held, "Altered motivation") || !strings.Contains(held, "One new value is needed.") {
373:			if strings.Contains(nextRead(t, f.root, "packages/alpha/evals/next_test.go"), "altered eval") {
387:func Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops(t *testing.T) {
425:			if !strings.Contains(held, "No findings") || !strings.Contains(held, "One new value is needed.") || !strings.Contains(held, "phase: audit") {
429:				t.Errorf("audit error advanced phases %s times, want only evals", got)
442:func fastAuditHook(t *testing.T, f nextFixture) {

===== packages/runner/evals/commit_recovery_test.go
13:func assertRecoveryWarnings(t *testing.T, diagnostics []project.Diagnostic, phases ...string) {
22:		t.Errorf("P05 recovery warnings = %#v, want one per phase %v", warnings, phases)
27:		if !warnings[i].Warning || !strings.Contains(message, phase) || !strings.Contains(message, "reset") || !strings.Contains(message, "retain") {
33:func Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle(t *testing.T) {
38:		t.Fatalf("Next after fixture agent commits: %v; state=%#v; steps=%#v", err, r.State, r.Steps)
41:		t.Fatalf("incomplete lifecycle after agent commits: %#v", r)
43:	for i, want := range []struct{ phase, role, stdout string }{{"evals", "worker", "worker-evals"}, {"audit", "audit", "auditor"}, {"impl", "worker", "worker-impl"}} {
52:			if diagnostic.Code == "P05" && diagnostic.Warning && strings.Contains(strings.ToLower(diagnostic.Message), phase) && strings.Contains(strings.ToLower(diagnostic.Message), "retain") {
61:		t.Errorf("evals verification HEAD = %s, want held base %s", got, base)
69:		if !strings.Contains(nextRead(t, f.logDir, observation.file), observation.path) {
74:		t.Errorf("branch has %s commits after three checked phases, want 3", count)
79:		t.Fatalf("checked phase commits = %v; branch commits = %v; error = %v", phaseCommits, commits, commitsErr)
81:	for i, want := range []struct{ path, phase string }{
88:			t.Errorf("phase commit %s phase = %v, parse error = %v; want %s", commits[i], document.Frontmatter["phase"], parseErr, want.phase)
100:			t.Errorf("fixture agent commit %d = %q, want %s and a full OID", i, fixtureCommits[i], label)
103:		if slices.Contains(reachable, parts[1]) {
108:		t.Errorf("fixture agent commits = %q, want four", fixtureCommits)
112:		t.Errorf("archive phase = %v, parse error = %v; want applied", archive.Frontmatter["phase"], parseErr)
115:		if !strings.Contains(nextRead(t, f.root, "changes/archive/0001-next.md"), marker) {
120:		if !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/next_test.go"), marker) {
124:	if !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/second_test.go"), "SECOND_AGENT_COMMIT") || !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/staged_test.go"), "STAGED_EVAL_WORK") || !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/untracked_test.go"), "UNTRACKED_EVAL_WORK") {
127:	if !strings.Contains(nextRead(t, f.root, "packages/alpha/internal/value.go"), `return "new"`) || !strings.Contains(nextRead(t, f.root, "packages/alpha/SPEC.md"), "Alpha MUST return the new value.") {
138:func Test_RUN_027_RUN_028_RecoversCommitBeforeInvocationFailure(t *testing.T) {
143:		t.Fatalf("agent failure = %v, steps=%#v", err, r.Steps)
150:		t.Errorf("failure HEAD = %s, want held base %s", head, base)
160:	if staged := nextGit(t, f.root, "diff", "--cached", "--name-only"); !strings.Contains(staged, "packages/alpha/evals/next_test.go") || !strings.Contains(staged, "packages/alpha/evals/staged_test.go") {
163:	if unstaged := nextGit(t, f.root, "diff", "--name-only"); !strings.Contains(unstaged, "packages/alpha/evals/next_test.go") {
166:	if untracked := nextGit(t, f.root, "ls-files", "--others", "--exclude-standard"); !strings.Contains(untracked, "packages/alpha/evals/untracked_test.go") {

===== packages/runner/evals/drill_test.go
23:func newDrillFixture(t *testing.T) drillFixture {
27:		t.Fatal(err)
31:		t.Fatal(err)
228:func (f drillFixture) run(ctx context.Context, timeout time.Duration) (runner.DrillResult, error) {
232:func (f drillFixture) mode(t *testing.T, mode string) {
237:func drillProjectProtectedFixture(t *testing.T) drillFixture {
250:func drillObserved(t *testing.T, f drillFixture, path string) string {
255:func drillNoSetupOrAgent(t *testing.T, f drillFixture) {
264:func drillCleaned(t *testing.T, f drillFixture) {
272:			t.Fatal(err)
288:func Test_RUN_050_DrillPreflightRejectsWithoutEffects(t *testing.T) {
304:				t.Fatal("preflight accepted " + mode)
306:			want := map[string]string{"unknown": "missing", "dirty": "dirty", "claim": "claim"}[mode]
307:			if !strings.Contains(strings.ToLower(err.Error()), want) {
319:func Test_RUN_051_SetupRunsDetachedAtHeadAndFailureStops(t *testing.T) {
323:	if err == nil || !strings.Contains(err.Error(), "SETUP_FAILURE") || !strings.Contains(err.Error(), "17") {
332:	if !strings.HasPrefix(setupRoot, filepath.Join(common, "fireweed", "worktrees")+string(os.PathSeparator)) || strings.TrimSpace(drillObserved(t, f, "setup.branch")) != "" || strings.TrimSpace(drillObserved(t, f, "setup.head")) != f.base || !strings.Contains(drillObserved(t, f, "setup"), "/fixture/new-controller") {
338:func drillNoAgent(t *testing.T, f drillFixture) {
345:func Test_RUN_052_RUN_053_RestrictedInputsAndIndependentGit(t *testing.T) {
359:		t.Fatalf("drill = %#v, %v", r, err)
362:	want := []string{
372:		t.Errorf("reconstruction inputs = %q, want %q", got, want)
374:	for path, want := range map[string]string{
377:		if got := strings.TrimSpace(drillObserved(t, f, path)); got != want {
378:			t.Errorf("%s = %q, want %q", path, got, want)
381:	if refs := drillObserved(t, f, "input.refs"); strings.Contains(refs, "invoking-only") || strings.Contains(refs, "fireweed/") {
390:		t.Errorf("input setup = %q, %v, want setup in %s before the worker", got, readErr, agentRoot)
409:func Test_RUN_054_PromptLetsWorkerRunOwnEval(t *testing.T) {
414:	for _, want := range []string{"ALP-001", "Prefer ordered values", "packages/alpha/api.go", "packages/alpha/evals/alpha_test.go", "DRILL_NOTES.md", "Contradicted Decisions", "Gaps", "no spec", "module", "disposable", "outside"} {
415:		if !strings.Contains(prompt, want) {
420:		if strings.Contains(prompt, forbidden) {
425:		t.Fatalf("drill with worker self-check: result=%q, steps=%#v, err=%v", r.Result, r.Steps, err)
429:		t.Errorf("worker eval roots = %q, want two runs in %q", got, inputRoot)
437:func Test_RUN_052_RUN_053_ProjectProtectedDisposableIsRetainedInputAndVerifierFile(t *testing.T) {
441:	if !strings.Contains(script, anchor) {
442:		t.Fatal("fixture agent insertion point missing")
454:		t.Fatalf("protected-path drill = %#v, %v", r, err)
469:	for _, want := range []string{"./" + protected, "./packages/core/api.go"} {
470:		if !slices.Contains(files, want) {
477:func Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs(t *testing.T) {
500:func Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions(t *testing.T) {
507:				t.Fatalf("drill with stray file: result=%q, err=%v", r.Result, err)
523:func Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound(t *testing.T) {
530:				t.Fatalf("retry = %#v, %v", r, err)
532:			if !strings.Contains(drillObserved(t, f, "prompt-2"), map[string]string{"agent_retry": "WORKER_FAILURE", "verify_retry": "own verification failure"}[mode]) {
544:	if mixedErr == nil || mixedResult.Result != "fail" || len(mixedResult.Steps) != 2 || !strings.Contains(drillObserved(t, mixed, "prompt-2"), "WORKER_FAILURE") {
559:func Test_RUN_056_ConfiguredTimeoutBoundsDrillWorker(t *testing.T) {
573:func Test_RUN_057_UsageDoesNotSpendAttemptAndStopsWhenBudgetExhausted(t *testing.T) {
620:func Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail(t *testing.T) {
628:	if len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "fail" || !strings.Contains(r.Own.Cases[0].Failure, "own verification failure") {
648:func Test_RUN_058_OwnCoverageRequiredAfterPassingCases(t *testing.T) {
655:	if err == nil || r.Result != "fail" || !strings.Contains(err.Error(), "ALP-002") || len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "pass" || len(r.Dependents) != 2 {
664:func Test_RUN_059_DependentCoverageRequiredAfterPassingCases(t *testing.T) {
671:	if err == nil || r.Result != "fail" || !strings.Contains(err.Error(), "AAR-002") || len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "pass" || len(r.Dependents) != 2 || len(r.Dependents["aardvark"].Cases) != 1 || r.Dependents["aardvark"].Cases[0].Result != "pass" || len(r.Dependents["zeta"].Cases) != 1 || r.Dependents["zeta"].Cases[0].Result != "pass" {
677:func drillReportSection(report, heading string) string {
689:func drillReportLine(section string, tokens ...string) bool {
693:			if !strings.Contains(line, token) {
705:func drillEvalEntry(section, module, testcase, id, result string) bool {
707:		if strings.Contains(line, testcase) && strings.Contains(line, module) && strings.Contains(line, result) && strings.Contains(strings.ReplaceAll(line, testcase, ""), id) {
714:func Test_RUN_060_RUN_061_DrillReportAndRetainedRef(t *testing.T) {
723:	if prior := nextGit(t, f.root, "show", f.base+":"+reportPath); !strings.Contains(prior, priorMarker) {
724:		t.Fatalf("committed prior report lacks marker: %q", prior)
728:		t.Fatalf("passing result = %#v, %v", r, err)
734:	if strings.Contains(report, priorMarker) {
737:	if prior := nextGit(t, f.root, "show", f.base+":"+reportPath); !strings.Contains(prior, priorMarker) {
740:	for _, want := range []string{"date:", "base_commit: " + f.base, "agent:", "model:", "result: pass", "## Own evals", "## Dependent evals", "## Contradicted Decisions", "## Gaps", "ALP-001_case", "AAR-001_case", "ZET-001_case", "ALP-001", "AAR-001", "ZET-001", "pass", "ordered preference may be contradicted", "module alpha", "no spec", "module aardvark"} {
741:		if !strings.Contains(report, want) {
757:	if !strings.Contains(report, "worker observation") {
769:func Test_RUN_060_RUN_061_FailedDrillReportsFailuresAndRetainsCode(t *testing.T) {
774:		t.Fatalf("failed drill = %#v, %v", r, err)
780:	for _, want := range []string{"result: fail", "own verification failure", "ALP-001_case", "ALP-001", "fail", "## Contradicted Decisions", "## Gaps"} {
781:		if !strings.Contains(report, want) {
796:	if report := nextRead(t, agentFailure.root, ".fireweed/drills/alpha.md"); !strings.Contains(report, "WORKER_FAILURE") || !strings.Contains(report, "result: fail") {
802:func Test_RUN_062_CleanupOnSuccessAndFailures(t *testing.T) {
822:func Test_RUN_062_CleanupFailureCannotPass(t *testing.T) {

===== packages/runner/evals/invoke_test.go
17:func agentFixture(t *testing.T) project.Snapshot {
21:		t.Fatal(err)
31:			t.Fatalf("git %v: %v: %s", args, err, out)
37:func agentScript(t *testing.T, root, name, body string) string {
41:		t.Fatal(err)
46:func shellQuote(s string) string { return "'" + strings.ReplaceAll(s, "'", "'\\''") + "'" }
48:func request(s project.Snapshot, role, prompt string) runner.AgentRequest {
52:func Test_RUN_001_RUN_002_RUN_005_RUN_006_CommandSelectionDirectoryEnvAndOutput(t *testing.T) {
58:	for _, tc := range []struct{ name, role, want string }{
75:func Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues(t *testing.T) {
79:		t.Fatal(err)
81:	if !strings.Contains(common.CommonDir, " ") {
82:		t.Fatalf("fixture common Git directory must contain spaces: %q", common.CommonDir)
96:			want := prompt + "\n" + common.CommonDir + "\n"
97:			if err != nil || result.Stdout != want || result.ExitCode != 0 {
98:				t.Errorf("%s substitution = %#v, %v; want stdout %q", tc.name, result, err, want)
104:func Test_RUN_003_RUN_004_WrapperPrefixForBothRoles(t *testing.T) {
108:		t.Fatal(err)
122:				t.Fatalf("wrapper did not run: %v; InvokeAgent = %#v, %v", readErr, result, err)
131:func Test_RUN_006_RUN_009_NonzeroExitRetainsStreamsAndCode(t *testing.T) {
140:func Test_RUN_007_RUN_009_TimeoutSelectionAndCapturedOutput(t *testing.T) {
172:func Test_RUN_007_RUN_009_TimeoutAndCancellationKillOrdinaryDescendant(t *testing.T) {
202:					t.Fatalf("agent exited before descendant was ready: %#v, %v", got.result, got.err)
204:					t.Fatal("agent never started descendant")
217:				t.Fatal("terminated agent did not stop promptly")
227:func Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines(t *testing.T) {
232:		want          bool
244:			if result.UsageLimit != tc.want || (err != nil) != (result.ExitCode != 0) {
245:				t.Errorf("%s = %#v, %v; want UsageLimit=%v", tc.name, result, err, tc.want)
251:func Test_RUN_009_LaunchFailureReturnsError(t *testing.T) {
256:	if err == nil || !strings.Contains(err.Error(), "missing-agent") || result.Stdout != "" || result.Stderr != "" {

===== packages/runner/evals/next_test.go
21:func nextWrite(t *testing.T, root, name, body string) {
25:		t.Fatal(err)
28:		t.Fatal(err)
32:func nextGit(t *testing.T, root string, args ...string) string {
39:		t.Fatalf("git %v: %v: %s", args, err, out)
44:func nextCommit(t *testing.T, root, message string) {
50:func nextRead(t *testing.T, root, path string) string {
54:		t.Fatal(err)
61:func newNextFixture(t *testing.T, queued, approved bool, workerMode string) nextFixture {
65:		t.Fatal(err)
69:		t.Fatal(err)
103:func Test_ALP_002_NewValue(t *testing.T) {
104:    if alpha.Value() != "new" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
123:func Test_ALP_002_NewValue(t *testing.T) {
124:    if alpha.Value() != "old" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
162:func Test_ALP_002_NewValue(t *testing.T) { // EVAL_CONTENT_MARKER
163:    if alpha.Value() != "new" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
170:func Value() string { return "new" }
252:	nextWrite(t, root, "packages/alpha/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-fixture/packages/alpha\")\nfunc Test_ALP_001_Value(t *testing.T) { if alpha.Value() == \"\" { t.Fatal(\"empty\") } }\n")
267:			t.Fatalf("bad fixture: %#v", d)
273:func (f nextFixture) run(allow bool) (runner.NextResult, error) {
277:func (f nextFixture) noAgent(t *testing.T) {
284:func nextAttempts(t *testing.T, f nextFixture, attempts int) {
290:			t.Fatalf("invalid attempts fixture: %#v", d)
295:func nextInvocations(t *testing.T, f nextFixture) []string {
300:func Test_RUN_010_EmptyQueue(t *testing.T) {
310:func Test_RUN_010_AllBlockedQueue(t *testing.T) {
325:func Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents(t *testing.T) {
330:		if err == nil || !strings.Contains(strings.ToLower(err.Error()), "approv") {
333:		if nextGit(t, f.root, "rev-parse", "HEAD") != before || strings.Contains(nextRead(t, f.root, nextChange), "run:") {
345:		if err == nil || !strings.Contains(err.Error(), "P01") {
372:				t.Fatalf("bad bound fixture: %#v", d)
376:		if err == nil || !strings.Contains(err.Error(), "P07") {
403:		if currentErr != nil || state.Phase != "evals" || state.Run == "" || !strings.Contains(nextRead(t, f.root, nextChange), "ALP-002") {
409:func Test_RUN_011_WaitingHeadSelectsIndependent(t *testing.T) {
413:	if blockErr == nil || !strings.Contains(blockErr.Error(), "P2-01: no fail-first case for ALP-002") || blocked.State.Change.Slug != "alpha" || blocked.State.Phase != "evals" {
414:		t.Fatalf("fixture alpha did not block: %#v, %v", blocked, blockErr)
418:		t.Fatalf("fixture eligible queue = %#v, %v", eligible, err)
427:func Test_RUN_010_EntirelyWaitingQueue(t *testing.T) {
431:	if blockErr == nil || !strings.Contains(blockErr.Error(), "P2-01: no fail-first case for ALP-002") || blocked.State.Change.Slug != "alpha" || blocked.State.Phase != "evals" {
432:		t.Fatalf("fixture alpha did not block: %#v, %v", blocked, blockErr)
435:		t.Fatal(err)
442:		t.Fatalf("fixture entirely waiting queue = %#v, %v", eligible, err)
453:func Test_RUN_013_RUN_014_CompleteSpecChange(t *testing.T) {
458:		t.Fatalf("Next: %#v, %v", r, err)
464:		t.Fatalf("steps = %#v", r.Steps)
466:	for i, want := range []struct{ phase, role, stdout string }{{"evals", "worker", "worker-evals"}, {"audit", "audit", "auditor"}, {"impl", "worker", "worker-impl"}} {
478:		if !strings.Contains(strings.Join(strings.Fields(prompt), " "), strings.Join(strings.Fields(content), " ")) {
486:			t.Fatalf("load %s playbook: %v", phase, err)
499:		for _, want := range []struct{ category, content string }{
523:		if strings.Contains(audit, secret) {
532:		if !strings.Contains(archived, marker) {
536:	if !strings.Contains(nextRead(t, f.root, "packages/alpha/SPEC.md"), "Alpha MUST return the new value.") {
545:	if old := nextGit(t, f.root, "show", "HEAD~2:packages/alpha/internal/value.go"); !strings.Contains(old, `return "old"`) {
548:	if first := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~2"); !strings.Contains(first, "evals") {
551:	if second := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~1"); !strings.Contains(second, "audit") {
554:	if applied := nextGit(t, f.root, "show", "HEAD:changes/archive/0001-next.md"); !strings.Contains(applied, "phase: applied") {
559:	if !strings.Contains(failReport, "ALP-002") || !strings.Contains(failReport, "fail") || !strings.Contains(passReport, "ALP-002") || !strings.Contains(passReport, "pass") {
567:func Test_RUN_015_InvocationFailureRetainsHeldPhase(t *testing.T) {
572:		t.Errorf("invocation failure: error=%v steps=%d, want 1", err, len(r.Steps))
580:		t.Errorf("agent calls = %d, want 1", got)
591:func nextAssertExhaustedBlock(t *testing.T, f nextFixture, r runner.NextResult, err error, phase, before string, wantPhaseCommits int) string {
601:		t.Errorf("fixture %s phase commits = %s, want %d", phase, got, wantPhaseCommits)
604:	if !strings.Contains(held, "phase: "+phase) || !strings.Contains(held, "run: "+r.State.Run) || !strings.Contains(held, "### ALP-002") {
608:		t.Errorf("block commit parents = %d, want held %s HEAD", len(parents)-1, phase)
611:		t.Errorf("block commit paths = %q, want state only", paths)
623:	if !strings.Contains(blocked, "phase: "+phase) || !strings.Contains(blocked, "run: "+r.State.Run) || !strings.Contains(blocked, "### ALP-002") || !strings.Contains(blocked, "blocked:") {
631:	if err == nil || !strings.Contains(err.Error(), ref) {
638:		t.Errorf("attempt ref kind = %q, want commit", kind)
643:func Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks(t *testing.T) {
647:	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") {
651:		t.Errorf("evals exhaustion steps = %d, want 2", len(r.Steps))
659:		t.Errorf("agent calls = %d, want 2", got)
663:		if saved := nextGit(t, f.root, "show", ref+":"+nextChange); !strings.Contains(saved, "phase: evals") || !strings.Contains(saved, "### ALP-002") || strings.Contains(saved, "blocked:") {
669:func Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase(t *testing.T) {
673:		t.Fatalf("Next after retry = %#v, %v", r, err)
676:		t.Fatalf("retry result = %#v", r)
678:	for i, want := range []struct{ phase, role, output string }{
691:		t.Errorf("agent calls = %d, want 5", got)
703:		if !strings.Contains(nextRead(t, f.root, path), marker) {
712:		if !strings.Contains(prompt, phase.feedback) {
716:	if prompt := nextRead(t, f.logDir, "worker-impl-2.prompt"); !strings.Contains(prompt, "--- FAIL: Test_ALP_002_NewValue") {
721:func Test_RUN_021_ConfiguredThreeAttempts(t *testing.T) {
726:		t.Fatalf("Next with three attempts = %#v, %v", r, err)
729:		t.Fatalf("configured retry result = %#v", r)
731:	for i, want := range []struct{ phase, output string }{
740:		t.Errorf("agent calls = %d, want 5", got)
742:	if prompt := nextRead(t, f.logDir, "worker-impl-3.prompt"); !strings.Contains(prompt, `got "wrong-2"`) {
747:func Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks(t *testing.T) {
751:	if err == nil || !strings.Contains(err.Error(), "E01: failed case") {
755:		t.Fatalf("exhaustion steps = %d, want 4", len(r.Steps))
757:	for i, want := range []struct{ phase, role, output string }{
767:		t.Errorf("agent calls = %d, want 4", got)
771:		if got := nextGit(t, f.root, "show", ref+":packages/alpha/internal/value.go"); !strings.Contains(got, `return "wrong"`) {
774:		if got := nextGit(t, f.root, "show", ref+":packages/alpha/internal/retained.go"); !strings.Contains(got, "RETAINED_IMPL_WORK") {
778:	if got := nextRead(t, f.root, "packages/alpha/internal/value.go"); !strings.Contains(got, `return "old"`) {

===== packages/runner/evals/prompt_file_test.go
16:func promptFileTempDir(t *testing.T) string {
20:		t.Fatal(err)
26:func promptFileObservation(t *testing.T, root, name string) string {
30:		t.Fatalf("read %s: %v", name, err)
35:func promptFileRemoved(t *testing.T, path, workspaceRoot string) {
45:func Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce(t *testing.T) {
49:		t.Fatal(err)
81:				t.Errorf("expanded prompt = %q, want %q", got, prompt)
84:				t.Errorf("prompt file content = %q, want %q", got, prompt)
87:				t.Errorf("expanded git common dir = %q, want %q", got, common.CommonDir)
94:func Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback(t *testing.T) {
98:		t.Fatal(err)
120:	for _, tc := range []struct{ name, role, want string }{
131:			if err != nil || result.Stdout != tc.want || result.ExitCode != 0 {
136:					t.Errorf("%s = %q, want %q", name, got, prompt)
140:				t.Errorf("wrapper common dir = %q, want %q", got, common.CommonDir)
145:					t.Errorf("%s received separate prompt file %q, want %q", name, duplicate, path)
153:func Test_RUN_067_WrapperOnlyPlaceholderCreatesPromptFile(t *testing.T) {
173:		t.Errorf("wrapper-only prompt file content = %q, want %q", got, prompt)
181:func Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure(t *testing.T) {
244:					t.Errorf("file during %s = %q, want %q", tc.name, got, prompt)
254:				if !strings.Contains(entry.Name(), "xcrun_db") {
262:func Test_RUN_067_UnselectedTemplateDoesNotRequirePromptFileIO(t *testing.T) {

===== packages/runner/evals/retry_failure_details_test.go
16:func retryDetailFixture(t *testing.T, alphaFailed, dependentFailed string) nextFixture {
21:	nextWrite(t, f.root, "packages/aardvark/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-fixture/packages/aardvark\")\nfunc Test_AAR_001_Value(t *testing.T) { if aardvark.Value() == \"\" { t.Fatal(\"empty\") } }\n")
55:			t.Fatalf("invalid retry-detail fixture: %#v", d)
61:func retryDetailSection(t *testing.T, f nextFixture) (string, string) {
66:		t.Fatal("retry did not retain the original worker context")
71:		t.Fatalf("retry prompt lost dependent hook stderr marker: %q", retry[len(first):])
76:func Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules(t *testing.T) {
90:		t.Fatalf("Next with JUnit-only dependent failure = %#v, %v", r, err)
93:	if !strings.Contains(feedback, "E01: failed case") || !strings.Contains(feedback, retryDetailHookMarker) {
96:	for _, want := range []string{
101:		if !strings.Contains(section, want) {
106:		if strings.Contains(section, absent) {
121:func Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions(t *testing.T) {
133:		t.Fatalf("Next with large JUnit failures = %#v, %v", r, err)
137:		t.Errorf("failed-case section has %d Unicode code points, want at most 16384", n)
139:	if !strings.Contains(section, "RUN044_HEAD_00_") || !strings.Contains(section, strings.Repeat("界", 4000)) {
156:			t.Errorf("case %d failure excerpt contains at least %d Unicode code points, want at most 4096", i, n)
158:		if strings.Contains(section, fmt.Sprintf("RUN044_TAIL_%02d", i)) {
162:	if !strings.Contains(section, "RUN044_DEPENDENT_FIRST") || strings.Index(section, "RUN044_DEPENDENT_FIRST") >= strings.Index(section, "RUN044_HEAD_00_") {
165:	if strings.Contains(section, "ALP-002_long_11") {
169:	if !strings.Contains(lower, "truncat") && !strings.Contains(lower, "shorten") && !strings.Contains(section, "…") {
172:	if !strings.Contains(lower, "omitt") && !strings.Contains(lower, "remaining") && !strings.Contains(lower, "more cases") && !strings.Contains(lower, "additional cases") {

===== packages/runner/evals/run_test.go
23:func newRunFixture(t *testing.T, queued int, unmerged, failBeta bool) runFixture {
32:func newRunFixtureMode(t *testing.T, queued int, unmerged bool, mode string) runFixture {
36:		t.Fatal(err)
40:		t.Fatal(err)
65:func Test_${prefix}_002_Added(t *testing.T) {
66:  if $module.Value() != "new" { t.Fatalf("${prefix}-002: got %q", $module.Value()) }
73:func Value() string { return "new" }
128:		nextWrite(t, root, "packages/"+m.name+"/evals/existing_test.go", fmt.Sprintf("package evals\nimport (\"testing\"; \"example.invalid/run-fixture/packages/%s\")\nfunc Test_%s_001_Value(t *testing.T) { if %s.Value() == \"\" { t.Fatal(\"empty\") } }\n", m.name, m.prefix, m.name))
145:			t.Fatalf("bad fixture: %#v", d)
151:func indexRunModule(name string) int {
160:func (f runFixture) request(max int, allow bool) runner.RunRequest {
164:func runArchive(t *testing.T, f runFixture, i int) string {
169:func runInvocations(t *testing.T, f runFixture) []string {
176:		t.Fatal(err)
181:func assertRunApplied(t *testing.T, f runFixture, applied []runner.NextResult, want ...string) {
185:func assertRunAppliedSteps(t *testing.T, f runFixture, applied []runner.NextResult, stepCount int, want ...string) {
188:		t.Fatalf("Applied count = %d, want %d: %#v", len(applied), len(want), applied)
190:	for i, name := range want {
196:		if !strings.Contains(archive, "phase: applied") || !strings.Contains(archive, "No findings; checked "+runModules[indexRunModule(name)].prefix+"-002") {
204:				t.Fatal(err)
208:				t.Fatal(err)
227:func assertRunAgentLog(t *testing.T, f runFixture, want ...string) {
231:		t.Fatalf("agent calls = %#v, want %d changes", got, len(want))
233:	for i, name := range want {
253:func runBlockedDependencyFixture(t *testing.T) runFixture {
261:			t.Fatalf("invalid dependent queue: %#v", d)
267:func assertRunBlockedAlpha(t *testing.T, f runFixture, err error) {
271:	if err == nil || !strings.Contains(err.Error(), cause) || !strings.Contains(err.Error(), ref) {
272:		t.Errorf("Run blocking error = %v, want verification cause and %s", err, ref)
279:	if blocked := nextRead(t, f.root, "changes/0001-alpha.md"); !strings.Contains(blocked, "blocked:") || !strings.Contains(blocked, "phase: evals") || !strings.Contains(blocked, "### ALP-002") {
290:func assertRunBetaWaiting(t *testing.T, f runFixture) {
293:	if strings.Contains(beta, "run:") || strings.Contains(beta, "phase:") || strings.Contains(beta, "blocked:") {
301:func Test_RUN_016_UnlimitedQueueOrder(t *testing.T) {
305:		t.Fatalf("Run: %#v, %v", r, err)
317:func Test_RUN_016_NegativeMaxIsUnlimited(t *testing.T) {
321:		t.Fatalf("Run: %#v, %v", r, err)
329:func Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding(t *testing.T) {
333:		t.Fatalf("first Run: %#v, %v", first, err)
346:		t.Fatalf("continuation: %#v, %v", second, err)
361:func Test_RUN_018_InitialEmptyAndExactLimit(t *testing.T) {
377:			t.Fatalf("exact Run: %#v, %v", r, err)
387:func Test_RUN_020_HeldFailureRetainsEarlierResults(t *testing.T) {
390:	if err == nil || !strings.Contains(err.Error(), "17") {
407:	if strings.Contains(nextRead(t, f.root, "changes/0003-gamma.md"), "phase:") {
416:func Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent(t *testing.T) {
430:func Test_RUN_025_RUN_026_PositiveMaxRetainsBlock(t *testing.T) {
436:			t.Fatalf("invalid Max fixture: %#v", d)
443:	if later := nextRead(t, f.root, "changes/0004-later.md"); strings.Contains(later, "run:") || strings.Contains(later, "phase:") {
447:		t.Errorf("Max agent calls = %#v, want exhausted alpha then one applied gamma", got)

===== packages/runner/evals/time_limit_test.go
20:func timeFastHook(t *testing.T, f runFixture) {
51:func timeAgentAfterWork(t *testing.T, f runFixture, body string) {
55:	if !strings.Contains(script, anchor) {
56:		t.Fatal("fixture agent insertion point missing")
61:func timeStop(t *testing.T, err error, limit string) {
63:	if err == nil || !runner.IsTimeStop(err) || !strings.Contains(err.Error(), limit) {
64:		t.Errorf("clean time stop = %v, want classified nonnil error naming %s", err, limit)
68:func timeHeldWork(t *testing.T, f runFixture, slug, phase, marker, want, head string) workflow.State {
71:	if !strings.Contains(nextRead(t, f.root, marker), want) {
83:func Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied(t *testing.T) {
107:		t.Errorf("Applied count = %d, want one retained alpha", len(r.Applied))
114:func Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded(t *testing.T) {
121:		t.Fatalf("zero For run = %#v, %v", r, err)
126:func Test_RUN_037_RUN_039_AgentDeadlineRetainsHeldWork(t *testing.T) {
137:		t.Fatalf("start interrupted agent step: %v", startErr)
156:func Test_RUN_037_VerificationDeadlineRetainsHeldWork(t *testing.T) {
166:	if !strings.Contains(hook, anchor) {
167:		t.Fatal("fixture hook insertion point missing")
174:		t.Fatalf("start interrupted verification step: %v", startErr)
209:	if markerReadErr != nil || !strings.Contains(string(marker), "TIME_HELD_WORK") {
229:		t.Errorf("HEAD = %s, want %s: %v", actualHead, head, headErr)
236:func Test_RUN_038_UsageWaitPreflightStopsBeforeWaiting(t *testing.T) {
260:func Test_RUN_039_ClassifierAndParentCancellation(t *testing.T) {
288:func Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop(t *testing.T) {
299:	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") || !strings.Contains(err.Error(), "refs/fireweed/attempts/alpha/1") || !strings.Contains(err.Error(), "8s") || runner.IsTimeStop(err) {

===== packages/runner/evals/usage_test.go
18:func usagePolicy(t *testing.T, root, policy string) {
26:func nextUsageAgent(t *testing.T, f nextFixture, behavior string) {
54:func Test_ALP_002_NewValue(t *testing.T) { if alpha.Value() != "old" { t.Fatal("wrong") } }
70:		t.Fatal("unknown fixture behavior")
73:	if !strings.Contains(script, anchor) {
74:		t.Fatal("fixture agent insertion point missing")
79:func runUsageAgent(t *testing.T, f runFixture, targets string) {
96:	if !strings.Contains(script, anchor) {
97:		t.Fatal("run fixture agent insertion point missing")
102:func usageHeld(t *testing.T, root, slug, phase, run, head, claim string) workflow.State {
106:		t.Fatalf("held state: %v", err)
112:		t.Errorf("HEAD = %s, want %s", got, head)
120:func Test_RUN_029_InvalidUsagePolicyPrecedesFreshStart(t *testing.T) {
133:			if err == nil || !strings.Contains(err.Error(), "C01") {
134:				t.Errorf("invalid policy error = %v, want C01", err)
147:func Test_RUN_029_RUN_031_DefaultUsageDurationsStopImmediately(t *testing.T) {
158:			if err == nil || !runner.IsUsageStop(err) || !strings.Contains(err.Error(), tc.wait) || !strings.Contains(err.Error(), tc.budget) || !strings.Contains(err.Error(), "0s") {
165:			if state.Change.Added[0].ID != "ALP-002" || !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/partial_test.go"), "USAGE_PARTIAL_WORK") || len(nextInvocations(t, f)) != 1 {
172:func Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep(t *testing.T) {
179:		t.Fatalf("short-budget stop = %#v, %v", first, err)
183:	if state.Change.Added[0].ID != "ALP-002" || !strings.Contains(partialWork, "USAGE_PARTIAL_WORK") {
196:	if !strings.Contains(approvedSpec, "### ALP-002") || !strings.Contains(approvedSpec, "Alpha MUST return the new value.") {
206:	if !strings.Contains(nextRead(t, f.root, "changes/archive/0001-next.md"), "ALP-002") {
210:		t.Errorf("resume changed held partial eval work: got %q, want %q", got, partialWork)
223:func Test_RUN_033_OwnedResumeUsesApprovedUsageBudget(t *testing.T) {
230:		t.Fatalf("first approved-budget stop = %#v, %v", first, err)
240:		t.Fatalf("approved state lost budget or assigned ID: %#v", held)
243:	if !strings.Contains(partialWork, "USAGE_PARTIAL_WORK") || len(nextInvocations(t, f)) != 2 {
244:		t.Fatal("first stop lost partial work or usage invocations")
247:	if !strings.Contains(config, "limit_budget: 5ms\n") {
248:		t.Fatal("fixture budget missing")
255:		t.Fatalf("resume with approved budget = %#v, %v", second, err)
267:		t.Errorf("resume changed partial eval work: got %q, want %q", got, partialWork)
273:		t.Errorf("total usage invocations = %d, want 4", got)
276:		t.Errorf("resume committed: HEAD = %s, want %s", got, head)
280:func Test_RUN_030_UsageRetryPreservesVerificationFeedback(t *testing.T) {
286:		t.Fatalf("verification and usage retry = %#v, %v", r, err)
291:	if prompt := nextRead(t, f.logDir, "worker-evals-3.prompt"); !strings.Contains(prompt, "P2-01: no fail-first case for ALP-002") {
294:	if !strings.Contains(nextRead(t, f.root, "packages/alpha/evals/retained_test.go"), "RETAINED_EVAL_WORK") {
299:func Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases(t *testing.T) {
304:	if err == nil || !runner.IsUsageStop(err) || !strings.Contains(err.Error(), "5ms") || len(r.Steps) != 4 {
305:		t.Fatalf("cross-phase budget = %#v, %v", r, err)
307:	for i, want := range []string{"evals", "evals", "audit", "impl"} {
308:		if r.Steps[i].Phase != want {
309:			t.Errorf("step %d phase = %q, want %q", i, r.Steps[i].Phase, want)
316:	if !strings.Contains(nextRead(t, f.root, "packages/alpha/internal/value.go"), "USAGE_IMPL_WORK") {
321:func Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied(t *testing.T) {
326:	if err == nil || !runner.IsUsageStop(err) || !strings.Contains(err.Error(), "5ms") {
330:	for i, want := range []string{"evals", "evals", "audit", "impl"} {
331:		if got := r.Applied[0].Steps[i].Phase; got != want {
332:			t.Errorf("alpha step %d phase = %q, want %q", i, got, want)
339:	if state.Change.Added[0].ID != "BET-002" || !strings.Contains(nextRead(t, f.root, "packages/beta/evals/usage_test.go"), "USAGE_PARTIAL_WORK") {
342:	if got := nextRead(t, f.root, "changes/0003-gamma.md"); strings.Contains(got, "phase:") || strings.Contains(got, "run:") {
346:		t.Errorf("agent calls = %#v, want two alpha evals, audit, impl, beta evals", got)
350:func Test_RUN_034_UsageStopClassifierWrapping(t *testing.T) {
363:func Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure(t *testing.T) {
368:	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") || !strings.Contains(err.Error(), "refs/fireweed/attempts/alpha/1") || !strings.Contains(err.Error(), "5ms") || runner.IsUsageStop(err) || len(r.Applied) != 0 {
374:	if blocked := nextRead(t, f.root, "changes/0001-alpha.md"); !strings.Contains(blocked, "blocked:") || !strings.Contains(blocked, "phase: evals") {
381:func Test_RUN_031_CancelledUsageWaitRetainsHeldWork(t *testing.T) {

 succeeded in 1220ms:
packages/workflow/internal/lifecycle/stub.go:39:func ref(slug string) string                 { return "refs/fireweed/claims/" + slug }
packages/workflow/internal/lifecycle/stub.go:40:func file(root, path string) string          { return filepath.Join(root, filepath.FromSlash(path)) }
packages/workflow/internal/lifecycle/stub.go:41:func read(root, path string) ([]byte, error) { return os.ReadFile(file(root, path)) }
packages/workflow/internal/lifecycle/stub.go:42:func write(root, path string, data []byte) error {
packages/workflow/internal/lifecycle/stub.go:49:func hard(ds []project.Diagnostic) error {
packages/workflow/internal/lifecycle/stub.go:63:func diag(code, msg string, warning bool) project.Diagnostic {
packages/workflow/internal/lifecycle/stub.go:66:func str(v any) string { s, _ := v.(string); return s }
packages/workflow/internal/lifecycle/stub.go:67:func num(v any, fallback int) int {
packages/workflow/internal/lifecycle/stub.go:81:func mapval(v any) map[string]any { m, _ := v.(map[string]any); return m }
packages/workflow/internal/lifecycle/stub.go:82:func list(v any) []string {
packages/workflow/internal/lifecycle/stub.go:96:func cloneMap(m map[string]any) map[string]any {
packages/workflow/internal/lifecycle/stub.go:106:func auditState(doc contracts.Document) map[string]any {
packages/workflow/internal/lifecycle/stub.go:116:func author(doc contracts.Document) (map[string]any, map[string]string) {
packages/workflow/internal/lifecycle/stub.go:127:func equal(a, b any) bool {
packages/workflow/internal/lifecycle/stub.go:132:func id() (string, error) {
packages/workflow/internal/lifecycle/stub.go:139:func changeBy(root, target string) (project.Change, error) {
packages/workflow/internal/lifecycle/stub.go:156:func unknownChange(target string, slugs []string) error {
packages/workflow/internal/lifecycle/stub.go:163:func existing(root, slug string) (claim, string, error) {
packages/workflow/internal/lifecycle/stub.go:181:func state(root string, ch project.Change, c claim, oid, head string) model.State {
packages/workflow/internal/lifecycle/stub.go:184:func storeClaim(root string, c claim, old string) (string, error) {
packages/workflow/internal/lifecycle/stub.go:198:func baseRefs(root string, config map[string]any) (local, shared string) {
packages/workflow/internal/lifecycle/stub.go:203:	remote := str(config["remote"])
packages/workflow/internal/lifecycle/stub.go:204:	if remote == "" {
packages/workflow/internal/lifecycle/stub.go:205:		remote = "origin"
packages/workflow/internal/lifecycle/stub.go:208:	shared = remote + "/" + base
packages/workflow/internal/lifecycle/stub.go:209:	if _, ok, _ := gitstore.Ref(root, "refs/remotes/"+shared); !ok {
packages/workflow/internal/lifecycle/stub.go:214:func approved(root, rev, path string) bool { _, e := gitstore.ReadAt(root, rev, path); return e == nil }
packages/workflow/internal/lifecycle/stub.go:218:func assign(root string, ch project.Change, snap project.Snapshot) (project.Change, error) {
packages/workflow/internal/lifecycle/stub.go:295:func Start(root, target string) (model.State, error) {
packages/workflow/internal/lifecycle/stub.go:419:func mustLoad(root string) project.Snapshot { s, _ := project.Load(root); return s }
packages/workflow/internal/lifecycle/stub.go:420:func Current(root, target string) (model.State, error) {
packages/workflow/internal/lifecycle/stub.go:452:func validate(s model.State, c claim) error {
packages/workflow/internal/lifecycle/stub.go:465:func loadCurrent(root, target string) (model.State, claim, error) {
packages/workflow/internal/lifecycle/stub.go:473:func specs(m project.Module) []contracts.Item {
packages/workflow/internal/lifecycle/stub.go:478:func module(s project.Snapshot, name string) project.Module {
packages/workflow/internal/lifecycle/stub.go:486:func selectedIDs(ch project.Change, projected []contracts.Item) []string {
packages/workflow/internal/lifecycle/stub.go:502:func evalFixIDs(ch project.Change) []string {
packages/workflow/internal/lifecycle/stub.go:508:func isRunArtifact(path string) bool {
packages/workflow/internal/lifecycle/stub.go:512:func stepChanges(root, base string) ([]gitstore.Change, error) {
packages/workflow/internal/lifecycle/stub.go:534:func stepPathspecs(diff []gitstore.Change, required ...string) []string {
packages/workflow/internal/lifecycle/stub.go:553:func evalFixZones(s model.State, c claim, diff []gitstore.Change) []project.Diagnostic {
packages/workflow/internal/lifecycle/stub.go:573:func dependents(s project.Snapshot, name, mode string) []string {
packages/workflow/internal/lifecycle/stub.go:601:func reportPath(s model.State, phase, name string) string {
packages/workflow/internal/lifecycle/stub.go:608:func saveReport(s model.State, phase, name string, r checks.Report) {
packages/workflow/internal/lifecycle/stub.go:614:func Verify(ctx context.Context, req model.Request) (model.Result, error) {
packages/workflow/internal/lifecycle/stub.go:735:func provenance(ch project.Change, phase, agent string, reports map[string]checks.Report) (project.Change, error) {
packages/workflow/internal/lifecycle/stub.go:760:func commitPhase(s model.State, c claim, oid string, ch project.Change, paths []string, subject, next string) (model.State, error) {
packages/workflow/internal/lifecycle/stub.go:787:func auditProcess(s model.State, c claim) (project.Change, string, bool, error) {
packages/workflow/internal/lifecycle/stub.go:866:func sortedKeys(m map[string]bool) []string {
packages/workflow/internal/lifecycle/stub.go:874:func rollbackFiles(root string, originals map[string][]byte) {
packages/workflow/internal/lifecycle/stub.go:883:func apply(s model.State, c claim, ch project.Change, oid string) (model.State, error) {
packages/workflow/internal/lifecycle/stub.go:971:func Done(ctx context.Context, req model.Request) (model.Result, error) {
packages/workflow/internal/lifecycle/stub.go:1078:func Context(root, target string, full bool) (string, error) {
packages/workflow/internal/lifecycle/resume_stub.go:24:func savedStep(root, oid string) ([]savedFile, error) {
packages/workflow/internal/lifecycle/resume_stub.go:59:func restoredChange(current, saved contracts.Document, phase string) (string, error) {
packages/workflow/internal/lifecycle/resume_stub.go:73:func StartFromAttempt(root, change string, attempt int) (model.State, error) {
packages/workflow/internal/lifecycle/block_stub.go:15:func blockGit(root string, input []byte, args ...string) (string, error) {
packages/workflow/internal/lifecycle/block_stub.go:28:func unstageRunArtifacts(root, base string) error {
packages/workflow/internal/lifecycle/block_stub.go:35:func blockDocument(root string, s model.State, c claim) (contracts.Document, error) {
packages/workflow/internal/lifecycle/block_stub.go:81:func nextAttempt(root, slug, oid string) (string, error) {
packages/workflow/internal/lifecycle/block_stub.go:102:func Block(request model.BlockRequest) (model.BlockResult, error) {
packages/workflow/internal/lifecycle/eligible_stub.go:11:func dependsOn(graph map[string][]string, module, dependency string) bool {
packages/workflow/internal/lifecycle/eligible_stub.go:29:func blockingRule(graph map[string][]string, candidate, earlier string, interfaceChanged bool) string {
packages/workflow/internal/lifecycle/eligible_stub.go:42:func Eligible(root string) ([]project.Change, error) {
23:func Run(ctx context.Context, request model.Request) (model.Result, error) {
72:func git(root string, input []byte, env map[string]string, args ...string) ([]byte, error) {
86:func gitText(root string, args ...string) (string, error) {
91:func canonical(path string) (string, error) {
99:func Inspect(dir string) (model.Repository, error) {
128:func BaseRef(root, remote, base string) (string, error) {
151:func Changes(root, base string) ([]model.Change, error) {
206:func ReadAt(root, revision, path string) ([]byte, error) {
210:func Commit(root, subject string, paths []string) (string, error) {
252:func WriteBlob(root string, data []byte) (string, error) {
257:func ReadBlob(root, oid string) ([]byte, error) {
261:func Ref(root, name string) (string, bool, error) {
277:func CASRef(root, name, old, new string) error {
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
U01: Unknown option "--help". Use help to see available commands and options.

exec
/bin/zsh -lc "cat packages/runner/evals/prompt_file_test.go; sed -n '65,270p;420,720p' packages/runner/evals/next_test.go; sed -n '1,230p' packages/runner/evals/drill_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "sed -n '190,217p;295,420p;420,472p;614,735p;970,1135p' packages/workflow/internal/lifecycle/stub.go; cat packages/workflow/internal/lifecycle/eligible_stub.go; sed -n '1,74p' packages/gitstore/internal/local/local.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 665ms:
	if e != nil {
		return "", e
	}
	if e = gitstore.CASRef(root, ref(c.Change), old, oid); e != nil {
		return "", e
	}
	return oid, nil
}
func baseRefs(root string, config map[string]any) (local, shared string) {
	base := str(config["base"])
	if base == "" {
		base = "main"
	}
	remote := str(config["remote"])
	if remote == "" {
		remote = "origin"
	}
	local = base
	shared = remote + "/" + base
	if _, ok, _ := gitstore.Ref(root, "refs/remotes/"+shared); !ok {
		shared = local
	}
	return
}
func approved(root, rev, path string) bool { _, e := gitstore.ReadAt(root, rev, path); return e == nil }

var handle = regexp.MustCompile(`new:[a-z0-9-]+`)

func Start(root, target string) (model.State, error) {
	repo, e := gitstore.Inspect(root)
	if e != nil {
		return model.State{}, e
	}
	root = repo.Root
	changes, e := gitstore.Changes(root, repo.Head)
	if e != nil {
		return model.State{}, e
	}
	if len(changes) > 0 {
		return model.State{}, errors.New("P01: start requires clean worktree")
	}
	if e = hard(project.Check(root)); e != nil {
		return model.State{}, e
	}
	if target != "" {
		if _, e = changeBy(root, target); e != nil {
			return model.State{}, e
		}
	}
	q, e := Eligible(root)
	if e != nil {
		return model.State{}, e
	}
	if len(q) == 0 {
		return model.State{}, errors.New("P01: no eligible change")
	}
	ch := q[0]
	if target != "" && target != ch.Slug && target != ch.Path && target != filepath.Base(ch.Path) {
		return model.State{}, errors.New("P01: change is not eligible queue head")
	}
	local, shared := baseRefs(root, mustLoad(root).Config)
	ds := []project.Diagnostic{}
	if !approved(root, local, ch.Path) && !approved(root, shared, ch.Path) {
		ds = append(ds, diag("P04", "change is not approved on local or shared base", true))
	}
	limit := num(mapval(mustLoad(root).Config["authoring"])["max_unapproved"], 5)
	if limit > 0 {
		entries, _ := os.ReadDir(file(root, "changes/archive"))
		count := 0
		for _, ent := range entries {
			if ent.IsDir() || !strings.HasSuffix(ent.Name(), ".md") {
				continue
			}
			p := "changes/archive/" + ent.Name()
			if !approved(root, shared, p) && !approved(root, shared, "changes/"+ent.Name()) {
				count++
			}
		}
		if count >= limit {
			ds = append(ds, diag("P07", "shared-base unapproved archive limit reached", false))
			return model.State{Diagnostics: ds}, errors.New("P07: unapproved archive limit")
		}
	}
	snap, e := project.Project(root, ch.Slug)
	if e != nil {
		return model.State{Diagnostics: ds}, e
	}
	run := str(ch.Document.Frontmatter["run"])
	if run == "" {
		run, e = id()
		if e != nil {
			return model.State{}, e
		}
	}
	phase := str(ch.Document.Frontmatter["phase"])
	freshRegenerate := phase == "" && ch.Document.Frontmatter["regenerate"] == true
	if phase == "" {
		phase = "evals"
		_, hasEvalFix := ch.Document.Frontmatter["eval_fix"]
		if freshRegenerate || len(ch.Added) == 0 && len(ch.Modified) == 0 && len(ch.Removed) == 0 &&
			len(list(ch.Document.Frontmatter["fixes"])) == 0 &&
			!hasEvalFix && ch.Document.Frontmatter["regenerate"] != true {
			phase = "impl"
		}
	}
	if ch.Document.Frontmatter["regenerate"] == true && !recentPassingDrill(root, ch.Module) {
		ds = append(ds, diag("P10", "no passing drill newer than the module's latest applied implementation", true))
	}
	c := claim{Worktree: root, Change: ch.Slug, Run: run, Phase: phase, PhaseBase: repo.Head, Snapshot: snap}
	c.Author, c.Sections = author(ch.Document)
	c.Audit = auditState(ch.Document)
	oid, e := storeClaim(root, c, "")
	if e != nil {
		return model.State{Diagnostics: ds}, fmt.Errorf("P02: claim already held: %w", e)
	}
	rollback := func(err error) (model.State, error) {
		_ = gitstore.CASRef(root, ref(ch.Slug), oid, "")
		return model.State{Diagnostics: ds}, err
	}
	ch, e = assign(root, ch, snap)
	if e != nil {
		return rollback(e)
	}
	ch.Document.Frontmatter["run"] = run
	ch.Document.Frontmatter["phase"] = phase
	rendered, e := contracts.RenderDocument(ch.Document)
	if e != nil {
		return rollback(e)
	}
	if e = write(root, ch.Path, []byte(rendered)); e != nil {
		return rollback(e)
	}
	if freshRegenerate {
		if e = removeDisposable(root, snap, ch.Module); e != nil {
			return rollback(e)
		}
	}
	snap, e = project.Project(root, ch.Slug)
	if e != nil {
		return rollback(e)
	}
	c.Snapshot = snap
	c.Author, c.Sections = author(ch.Document)
	c.Audit = auditState(ch.Document)
	oid, e = storeClaim(root, c, oid)
	if e != nil {
		return model.State{}, fmt.Errorf("P02: update claim: %w", e)
	}
	s := state(root, ch, c, oid, repo.Head)
	s.Diagnostics = ds
	return s, nil
}
func mustLoad(root string) project.Snapshot { s, _ := project.Load(root); return s }
func Current(root, target string) (model.State, error) {
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
	}
	c, _, e := existing(s.Snapshot.Root, s.Change.Slug)
	return s, c, e
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
}
func Done(ctx context.Context, req model.Request) (model.Result, error) {
	req.Quick = false
	s, c, e := loadCurrent(req.Root, req.Change)
	result := model.Result{State: s, Reports: map[string]checks.Report{}}
	if e != nil {
		return result, e
	}
	if e = validate(s, c); e != nil {
		return result, e
	}
	result, e = Verify(ctx, req)
	if e != nil {
		return result, e
	}
	s = result.State
	// A hook can write files. Check its aftermath before committing or applying.
	fresh, e := Current(s.Snapshot.Root, s.Change.Slug)
	if e != nil {
		return result, e
	}
	c, oid, e := existing(s.Snapshot.Root, s.Change.Slug)
	if e != nil {
		return result, e
	}
	if e = validate(fresh, c); e != nil {
		return result, e
	}
	if e = hard(project.Check(s.Snapshot.Root)); e != nil {
		return result, e
	}
	diffAfter, e := stepChanges(s.Snapshot.Root, c.PhaseBase)
	if e != nil {
		return result, e
	}
	postZones := checks.Zones(checks.ZoneRequest{Snapshot: c.Snapshot, Module: fresh.Change.Module, Phase: c.Phase, ChangePath: fresh.Change.Path, Base: c.PhaseBase, InterfaceDeclared: strings.TrimSpace(c.Sections["Interface"]) != "", Changes: diffAfter})
	postZones = append(postZones, evalFixZones(fresh, c, diffAfter)...)
	result.Diagnostics = append(result.Diagnostics, postZones...)
	if e = hard(postZones); e != nil {
		return result, e
	}
	s = fresh
	result.State = fresh
	ch := s.Change
	next := ""
	blocked := false
	switch c.Phase {
	case "evals":
		next = "audit"
		if mapval(c.Snapshot.Config["audit"])["enabled"] == false {
			next = "impl"
		}
	case "audit":
		ch, next, blocked, e = auditProcess(s, c)
		if e != nil {
			return result, e
		}
	case "impl":
		next = "applied"
	default:
		return result, fmt.Errorf("P01: unsupported phase %s", c.Phase)
	}
	ch, e = provenance(ch, c.Phase, req.Agent, result.Reports)
	if e != nil {
		return result, e
	}
	if e = unstageRunArtifacts(s.Snapshot.Root, c.PhaseBase); e != nil {
		return result, e
	}
	if next == "applied" {
		result.State, e = apply(s, c, ch, oid)
		return result, e
	}
	paths := stepPathspecs(nil, ch.Path)
	if c.Phase == "evals" {
		paths = stepPathspecs(diffAfter, ch.Path)
	}
	subject := "test(" + ch.Module + "): evals " + ch.Slug
	if c.Phase == "audit" {
		subject = "docs(" + ch.Module + "): audit " + ch.Slug
	}
	if blocked {
		subject = "chore(" + ch.Module + "): block " + ch.Slug
	}
	if blocked {
		ch.Document.Frontmatter["phase"] = c.Phase
		rendered, e := contracts.RenderDocument(ch.Document)
		if e != nil {
			return result, e
		}
		if e = write(s.Snapshot.Root, ch.Path, []byte(rendered)); e != nil {
			return result, e
		}
		head, e := gitstore.Commit(s.Snapshot.Root, subject, paths)
		if e != nil {
			return result, e
		}
		if e = gitstore.CASRef(s.Snapshot.Root, ref(ch.Slug), oid, ""); e != nil {
			return result, e
		}
		result.State.Head = head
		result.State.PhaseBase = head
		result.State.Change = ch
		return result, errors.New("P06: change blocked")
	}
	result.State, e = commitPhase(s, c, oid, ch, paths, subject, next)
	return result, e
}
func Context(root, target string, full bool) (string, error) {
	s, e := Current(root, target)
	if e != nil {
		return "", e
	}
	var b strings.Builder
	raw, e := read(s.Snapshot.Root, s.Change.Path)
	if e != nil {
		return "", e
	}
	b.WriteString("# Change\n")
	b.Write(raw)
	m := module(s.Snapshot, s.Change.Module)
	spec, e := contracts.RenderDocument(m.Document)
	if e != nil {
		return "", e
	}
	b.WriteString("\n# Projected spec\n")
	b.WriteString(spec)
	b.WriteString("\n# Interface\n")
	for _, pat := range m.Zones.Interface {
		_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			rel, _ := filepath.Rel(file(s.Snapshot.Root, m.Dir), path)
			if !contracts.MatchGlob(pat, filepath.ToSlash(rel)) {
				return nil
			}
			data, e := os.ReadFile(path)
			if e == nil {
				b.WriteString("\n## " + filepath.ToSlash(filepath.Join(m.Dir, rel)) + "\n")
				b.Write(data)
			}
			return nil
		})
	}
	sys, e := read(s.Snapshot.Root, "SYSTEM.md")
	if e == nil {
		b.WriteString("\n# SYSTEM.md\n")
		b.Write(sys)
	}
	playbookName := s.Phase
	if s.Phase == "impl" && s.Change.Document.Frontmatter["regenerate"] == true {
		playbookName = "regenerate"
	}
	play, e := assets.Playbook(playbookName)
	if e == nil {
		b.WriteString("\n# Playbook\n")
		b.WriteString(play)
	}
	b.WriteString("\n# Other relevant files\n")
	_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(s.Snapshot.Root, path)
		rel = filepath.ToSlash(rel)
package lifecycle

import (
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

// dependsOn follows the dependency graph in the candidate's projected snapshot.
func dependsOn(graph map[string][]string, module, dependency string) bool {
	seen := map[string]bool{}
	var visit func(string) bool
	visit = func(name string) bool {
		if seen[name] {
			return false
		}
		seen[name] = true
		for _, next := range graph[name] {
			if next == dependency || visit(next) {
				return true
			}
		}
		return false
	}
	return visit(module)
}

func blockingRule(graph map[string][]string, candidate, earlier string, interfaceChanged bool) string {
	if candidate == earlier {
		return "same_module"
	}
	if dependsOn(graph, candidate, earlier) {
		return "depends_on_blocked"
	}
	if interfaceChanged && dependsOn(graph, earlier, candidate) {
		return "interface_dependency"
	}
	return ""
}

func Eligible(root string) ([]project.Change, error) {
	queue, err := project.Queue(root)
	if err != nil {
		return nil, err
	}
	eligible := make([]project.Change, 0, len(queue))
	var blocked []string
	for _, candidate := range queue {
		if str(candidate.Document.Frontmatter["blocked"]) != "" {
			blocked = append(blocked, candidate.Module)
			continue
		}
		projected, err := project.Project(root, candidate.Slug)
		if err != nil {
			return nil, err
		}
		graph := make(map[string][]string, len(projected.Modules))
		for _, module := range projected.Modules {
			graph[module.Name] = module.DependsOn
		}
		section, _ := contracts.Section(candidate.Document.Body, "Interface")
		interfaceChanged := strings.TrimSpace(section) != ""
		waiting := false
		for _, earlier := range blocked {
			if blockingRule(graph, candidate.Module, earlier, interfaceChanged) != "" {
				waiting = true
				break
			}
		}
		if !waiting {
			eligible = append(eligible, candidate)
		}
	}
	return eligible, nil
}
package local

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore/model"
)

const gitTimeout = 30 * time.Second

// Run starts each command in its own process group so cancellation also stops
// children that remain in that group.
func Run(ctx context.Context, request model.Request) (model.Result, error) {
	var result model.Result
	if len(request.Args) == 0 {
		return result, errors.New("empty command")
	}
	if err := ctx.Err(); err != nil {
		return result, err
	}
	if request.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, request.Timeout)
		defer cancel()
	}
	cmd := exec.Command(request.Args[0], request.Args[1:]...)
	cmd.Dir = request.Dir
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Env = os.Environ()
	keys := make([]string, 0, len(request.Env))
	for key := range request.Env {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		cmd.Env = append(cmd.Env, key+"="+request.Env[key])
	}
	cmd.Stdin = bytes.NewReader(request.Input)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	if err := cmd.Start(); err != nil {
		return result, err
	}
	done := make(chan error, 1)
	go func() { done <- cmd.Wait() }()
	var err error
	select {
	case err = <-done:
	case <-ctx.Done():
		_ = syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		err = <-done
		result = model.Result{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: cmd.ProcessState.ExitCode()}
		return result, ctx.Err()
	}
	result = model.Result{Stdout: stdout.String(), Stderr: stderr.String(), ExitCode: cmd.ProcessState.ExitCode()}
	if err != nil {
		return result, err
	}
	return result, nil
}

func git(root string, input []byte, env map[string]string, args ...string) ([]byte, error) {
	if env == nil {
		env = make(map[string]string)

 succeeded in 665ms:
package evals_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/runner"
)

func promptFileTempDir(t *testing.T) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "temporary files with spaces")
	if err := os.Mkdir(root, 0755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("TMPDIR", root)
	return root
}

func promptFileObservation(t *testing.T, root, name string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(data)
}

func promptFileRemoved(t *testing.T, path, workspaceRoot string) {
	t.Helper()
	if !filepath.IsAbs(path) || path == workspaceRoot || strings.HasPrefix(path, workspaceRoot+string(os.PathSeparator)) {
		t.Errorf("prompt file path %q is not absolute outside input workspace %q", path, workspaceRoot)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Errorf("prompt file %q remains after invocation: %v", path, err)
	}
}

func Test_RUN_003_PromptFileExpansionInShellAndArgvPreservesValuesOnce(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	observations := t.TempDir()
	script := agentScript(t, snap.Root, "prompt-file-args.sh", `printf '%s' "$2" > "$1/prompt"
printf '%s' "$3" > "$1/path"
printf '%s' "$4" > "$1/common"
cat "$3" > "$1/content" || exit 12
printf 'agent-ok'`)
	prompt := "a'b ; $(printf injected) {git_common_dir} {prompt_file}\nsecond line Ω"
	for _, tc := range []struct {
		name, role string
		command    any
	}{
		{"worker shell", "worker", shellQuote(script) + " " + shellQuote(observations) + " {prompt} {prompt_file} {git_common_dir}"},
		{"worker argv", "worker", []any{script, observations, "{prompt}", "{prompt_file}", "{git_common_dir}"}},
		{"audit shell", "audit", shellQuote(script) + " " + shellQuote(observations) + " {prompt} {prompt_file} {git_common_dir}"},
		{"audit argv", "audit", []any{script, observations, "{prompt}", "{prompt_file}", "{git_common_dir}"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			promptFileTempDir(t)
			snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf unselected"}}
			delete(snap.Config, "audit")
			if tc.role == "worker" {
				snap.Config["agent"] = map[string]any{"command": tc.command}
			} else {
				snap.Config["audit"] = map[string]any{"command": tc.command}
			}
			result, err := runner.InvokeAgent(context.Background(), request(snap, tc.role, prompt))
			if err != nil || result.Stdout != "agent-ok" || result.ExitCode != 0 {
				t.Errorf("InvokeAgent = %#v, %v", result, err)
			}
			if got := promptFileObservation(t, observations, "prompt"); got != prompt {
				t.Errorf("expanded prompt = %q, want %q", got, prompt)
			}
			if got := promptFileObservation(t, observations, "content"); got != prompt {
				t.Errorf("prompt file content = %q, want %q", got, prompt)
			}
			if got := promptFileObservation(t, observations, "common"); got != common.CommonDir {
				t.Errorf("expanded git common dir = %q, want %q", got, common.CommonDir)
			}
			promptFileRemoved(t, promptFileObservation(t, observations, "path"), snap.Root)
		})
	}
}

func Test_RUN_003_RUN_067_WrapperUsesOnePromptFileForBothRolesAndFallback(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	observations := t.TempDir()
	worker := agentScript(t, snap.Root, "prompt-file-worker.sh", `printf '%s' "$2" > "$1/command-path"
cat "$2" > "$1/command-content" || exit 12
printf worker`)
	auditor := agentScript(t, snap.Root, "prompt-file-auditor.sh", `printf '%s' "$2" > "$1/command-path"
cat "$2" > "$1/command-content" || exit 12
printf audit`)
	wrapper := agentScript(t, snap.Root, "prompt-file-wrapper.sh", `printf '%s' "$2" > "$1/prompt"
printf '%s' "$3" > "$1/path"
printf '%s' "$4" > "$1/common"
printf '%s' "$5" > "$1/second-path"
cat "$3" > "$1/content" || exit 12
shift 5
exec "$@"`)
	snap.Config["agent"] = map[string]any{
		"command": []any{worker, observations, "{prompt_file}"},
		"wrap":    []any{wrapper, observations, "{prompt}", "{prompt_file}", "{git_common_dir}", "{prompt_file}"},
	}
	snap.Config["audit"] = map[string]any{"command": []any{auditor, observations, "{prompt_file}"}}
	prompt := "wrapper 'literal' {prompt_file} {git_common_dir}\nΩ"
	for _, tc := range []struct{ name, role, want string }{
		{"worker", "worker", "worker"},
		{"audit", "audit", "audit"},
		{"audit fallback", "audit", "worker"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			promptFileTempDir(t)
			if tc.name == "audit fallback" {
				delete(snap.Config, "audit")
			}
			result, err := runner.InvokeAgent(context.Background(), request(snap, tc.role, prompt))
			if err != nil || result.Stdout != tc.want || result.ExitCode != 0 {
				t.Errorf("wrapped %s = %#v, %v", tc.role, result, err)
			}
			for _, name := range []string{"prompt", "content", "command-content"} {
				if got := promptFileObservation(t, observations, name); got != prompt {
					t.Errorf("%s = %q, want %q", name, got, prompt)
				}
			}
			if got := promptFileObservation(t, observations, "common"); got != common.CommonDir {
				t.Errorf("wrapper common dir = %q, want %q", got, common.CommonDir)
			}
			path := promptFileObservation(t, observations, "path")
			for _, name := range []string{"second-path", "command-path"} {
				if duplicate := promptFileObservation(t, observations, name); duplicate != path {
					t.Errorf("%s received separate prompt file %q, want %q", name, duplicate, path)
				}
			}
			promptFileRemoved(t, path, snap.Root)
		})
	}
}

func Test_RUN_067_WrapperOnlyPlaceholderCreatesPromptFile(t *testing.T) {
	snap := agentFixture(t)
	observations := t.TempDir()
	promptFileTempDir(t)
	wrapper := agentScript(t, snap.Root, "wrapper-only.sh", `printf '%s' "$1" > "$2/path"
cat "$1" > "$2/content" || exit 12
printf '%s' "$3" > "$2/second-path"
shift 3
exec "$@"`)
	snap.Config["agent"] = map[string]any{
		"command": []any{"/bin/sh", "-c", "printf wrapper-only"},
		"wrap":    []any{wrapper, "{prompt_file}", observations, "{prompt_file}"},
	}
	prompt := "wrapper-only 'literal' {prompt_file}\nΩ"
	result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", prompt))
	if err != nil || result.Stdout != "wrapper-only" || result.ExitCode != 0 {
		t.Errorf("wrapper-only invocation = %#v, %v", result, err)
	}
	path := promptFileObservation(t, observations, "path")
	if got := promptFileObservation(t, observations, "content"); got != prompt {
		t.Errorf("wrapper-only prompt file content = %q, want %q", got, prompt)
	}
	if second := promptFileObservation(t, observations, "second-path"); second != path {
		t.Errorf("wrapper-only placeholders used different files: %q, %q", path, second)
	}
	promptFileRemoved(t, path, snap.Root)
}

func Test_RUN_067_PromptFileLifetimeOnExitTimeoutCancellationAndLaunchFailure(t *testing.T) {
	for _, tc := range []struct{ name string }{
		{"success"}, {"nonzero"}, {"timeout"}, {"cancellation"}, {"launch failure"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap := agentFixture(t)
			observations := t.TempDir()
			tempRoot := promptFileTempDir(t)
			prompt := "exact prompt\nwith trailing newline Ω\n"
			if tc.name == "launch failure" {
				snap.Config["agent"] = map[string]any{"command": []any{filepath.Join(snap.Root, "missing-agent"), "{prompt_file}"}}
			} else {
				script := agentScript(t, snap.Root, "prompt-file-lifetime.sh", `printf '%s' "$1" > "$2/path"
test -r "$1" || exit 12
prompt_bytes=
IFS= read -r -d '' prompt_bytes < "$1" || test -n "$prompt_bytes" || exit 12
printf '%s' "$prompt_bytes" > "$2/content" || exit 12
printf ready > "$2/ready"
case "$3" in
  nonzero) printf partial; printf warning >&2; exit 9;;
  timeout|cancellation) printf partial; printf warning >&2; sleep 5;;
esac
printf success`)
				snap.Config["agent"] = map[string]any{"command": []any{"/bin/bash", script, "{prompt_file}", observations, tc.name}}
			}
			req := request(snap, "worker", prompt)
			req.Timeout = 2 * time.Second
			ctx := context.Background()
			if tc.name == "timeout" {
				req.Timeout = 150 * time.Millisecond
			}
			if tc.name == "cancellation" {
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				defer cancel()
				go func() {
					ready := filepath.Join(observations, "ready")
					for i := 0; i < 200; i++ {
						if _, err := os.Stat(ready); err == nil {
							cancel()
							return
						}
						time.Sleep(10 * time.Millisecond)
					}
					cancel()
				}()
			}
			result, err := runner.InvokeAgent(ctx, req)
			if tc.name == "success" {
				if err != nil || result.Stdout != "success" || result.ExitCode != 0 {
					t.Errorf("successful invocation = %#v, %v", result, err)
				}
			} else if err == nil {
				t.Errorf("%s returned no error: %#v", tc.name, result)
			}
			if tc.name == "cancellation" && !errors.Is(err, context.Canceled) {
				t.Errorf("cancellation error = %v", err)
			}
			if (tc.name == "nonzero" && result.ExitCode != 9) || (tc.name == "timeout" && !errors.Is(err, context.DeadlineExceeded)) {
				t.Errorf("%s result = %#v, %v", tc.name, result, err)
			}
			if tc.name != "launch failure" {
				if got := promptFileObservation(t, observations, "content"); got != prompt {
					t.Errorf("file during %s = %q, want %q", tc.name, got, prompt)
				}
				promptFileRemoved(t, promptFileObservation(t, observations, "path"), snap.Root)
			}
			entries, readErr := os.ReadDir(tempRoot)
			if readErr != nil {
				t.Errorf("read temporary directory after %s: %v", tc.name, readErr)
			}
			for _, entry := range entries {
				// macOS command-line tooling can populate this cache in TMPDIR.
				if !strings.Contains(entry.Name(), "xcrun_db") {
					t.Errorf("temporary file %q remains after %s", entry.Name(), tc.name)
				}
			}
		})
	}
}

func Test_RUN_067_UnselectedTemplateDoesNotRequirePromptFileIO(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf worker"}}
	snap.Config["audit"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf '{prompt_file}'"}}
	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "missing-temp-directory"))
	result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", "unused prompt"))
	if err != nil || result.Stdout != "worker" || result.ExitCode != 0 {
		t.Errorf("worker selected without prompt file = %#v, %v", result, err)
	}
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf '{prompt_file}'"}}
	snap.Config["audit"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf audit"}}
	result, err = runner.InvokeAgent(context.Background(), request(snap, "audit", "unused prompt"))
	if err != nil || result.Stdout != "audit" || result.ExitCode != 0 {
		t.Errorf("audit selected without prompt file = %#v, %v", result, err)
	}
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", "printf worker"}}
	delete(snap.Config, "audit")
	result, err = runner.InvokeAgent(context.Background(), request(snap, "audit", "unused prompt"))
	if err != nil || result.Stdout != "worker" || result.ExitCode != 0 {
		t.Errorf("audit fallback without prompt file = %#v, %v", result, err)
	}
}
		t.Fatal(err)
	}
	logDir, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	worker := filepath.Join(logDir, "worker.sh")
	auditor := filepath.Join(logDir, "auditor.sh")
	// The scripts are fixture agents. Only the runner may call workflow transitions.
	nextWrite(t, logDir, "worker.sh", `#!/bin/bash
phase=$(sed -n 's/^phase: //p' changes/0001-next.md | head -1)
attempt_file="$1/$phase-attempts"
attempt=0
[ ! -f "$attempt_file" ] || attempt=$(cat "$attempt_file")
attempt=$((attempt + 1))
printf '%s' "$attempt" > "$attempt_file"
printf '%s' "$3" > "$1/worker-$phase.prompt"
printf '%s' "$3" > "$1/worker-$phase-$attempt.prompt"
printf '%s|%s|%s\n' "$phase" "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
case "$2/$phase" in
  noedit/evals|exhaust-impl/impl)
    if [ "$attempt" -eq 1 ]; then
      git rev-parse HEAD > "$1/$phase-phase-head"
      git for-each-ref --format='%(refname)' refs/fireweed/claims > "$1/$phase-claim"
      cp changes/0001-next.md "$1/$phase-held-change"
    else
      git rev-parse HEAD > "$1/$phase-final-head"
    fi
    ;;
esac
if [ "$2" = exit ]; then printf 'partial worker output'; printf 'worker failure' >&2; exit 9; fi
if [ "$2" = commit-exit ]; then
  cat > packages/alpha/evals/next_test.go <<'GO'
package evals
import (
    "testing"
    "example.invalid/next-fixture/packages/alpha"
)
func Test_ALP_002_NewValue(t *testing.T) {
    if alpha.Value() != "new" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
}
GO
  git add packages/alpha/evals/next_test.go
  git commit -q -m 'fixture agent evals commit'
  printf 'package evals\n// STAGED_FAILURE_WORK\n' > packages/alpha/evals/staged_test.go
  git add packages/alpha/evals/staged_test.go
  printf '// TRACKED_FAILURE_WORK\n' >> packages/alpha/evals/next_test.go
  printf 'package evals\n// UNTRACKED_FAILURE_WORK\n' > packages/alpha/evals/untracked_test.go
  printf 'partial worker output'; printf 'worker failure' >&2; exit 9
fi
if [ "$2" = noedit ]; then printf 'no edit'; exit 0; fi
if [ "$2" = retry-both ] && [ "$phase" = evals ] && [ "$attempt" -eq 1 ]; then
  cat > packages/alpha/evals/next_test.go <<'GO'
package evals
import (
    "testing"
    "example.invalid/next-fixture/packages/alpha"
)
func Test_ALP_002_NewValue(t *testing.T) {
    if alpha.Value() != "old" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
}
GO
  printf 'package evals\n// RETAINED_EVAL_WORK\n' > packages/alpha/evals/retained_test.go
  printf 'worker-evals-1'
  exit 0
fi
if [ "$2" = retry-both ] && [ "$phase" = evals ] && [ "$attempt" -eq 2 ]; then
  if [ -f packages/alpha/evals/retained_test.go ]; then printf 'retained\n' > "$1/evals-retention"; else printf 'lost\n' > "$1/evals-retention"; fi
fi
if [ "$phase" = impl ]; then
  wrong=0
  case "$2" in
    retry-impl|retry-both) [ "$attempt" -ne 1 ] || wrong=1 ;;
    retry-third) [ "$attempt" -ge 3 ] || wrong=1 ;;
    exhaust-impl) wrong=1 ;;
  esac
  if [ "$wrong" -eq 1 ]; then
    wrong_value=wrong
    [ "$2" != retry-third ] || wrong_value="wrong-$attempt"
    printf 'package internal\nfunc Value() string { return "%s" }\n' "$wrong_value" > packages/alpha/internal/value.go
    printf 'package internal\nconst RetainedWork = "RETAINED_IMPL_WORK"\n' > packages/alpha/internal/retained.go
    printf 'worker-impl-%s' "$attempt"
    exit 0
  fi
  if [ "$attempt" -gt 1 ]; then
    if [ -f packages/alpha/internal/retained.go ]; then printf 'retained\n' > "$1/impl-retention"; else printf 'lost\n' > "$1/impl-retention"; fi
  fi
fi
sleep 0.15
case "$phase" in
  evals)
    cat > packages/alpha/evals/next_test.go <<'GO'
package evals
import (
    "testing"
    "example.invalid/next-fixture/packages/alpha"
)
func Test_ALP_002_NewValue(t *testing.T) { // EVAL_CONTENT_MARKER
    if alpha.Value() != "new" { t.Fatalf("ALP-002: got %q", alpha.Value()) }
}
GO
    ;;
  impl)
    cat > packages/alpha/internal/value.go <<'GO'
package internal
func Value() string { return "new" }
GO
    ;;
  *) exit 8 ;;
esac
if [ "$2" = commit-success ]; then
  git add -A
  git commit -q -m "fixture agent $phase first commit"
  printf '%s %s\n' "$phase-first" "$(git rev-parse HEAD)" >> "$1/fixture-commits"
  case "$phase" in
    evals)
      printf 'package evals\n// SECOND_AGENT_COMMIT\n' > packages/alpha/evals/second_test.go
      git add packages/alpha/evals/second_test.go
      git commit -q -m 'fixture agent evals second commit'
      printf '%s %s\n' 'evals-second' "$(git rev-parse HEAD)" >> "$1/fixture-commits"
      printf 'package evals\n// STAGED_EVAL_WORK\n' > packages/alpha/evals/staged_test.go
      git add packages/alpha/evals/staged_test.go
      printf '// TRACKED_EVAL_WORK\n' >> packages/alpha/evals/next_test.go
      printf 'package evals\n// UNTRACKED_EVAL_WORK\n' > packages/alpha/evals/untracked_test.go
      ;;
  esac
fi
printf 'worker-%s' "$phase"
`)
	nextWrite(t, logDir, "auditor.sh", `#!/bin/bash
printf '%s' "$2" > "$1/audit.prompt"
printf 'audit|%s|%s\n' "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
printf '%s' "$PWD" > "$1/audit.cwd"
cat >> changes/0001-next.md <<'AUDIT'

## Audit
No findings
Checked IDs: ALP-002
AUDIT
if [ -f "$1/commit-success" ]; then
  git add changes/0001-next.md
  git -c user.name=Fixture -c user.email=fixture@example.invalid commit -q -m 'fixture agent audit commit'
  printf '%s %s\n' 'audit' "$(git rev-parse HEAD)" >> "$1/fixture-commits"
fi
printf 'auditor'
`)
	if workerMode == "commit-success" {
		nextWrite(t, logDir, "commit-success", "enabled\n")
	}
	nextWrite(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	nextWrite(t, root, "go.mod", "module example.invalid/next-fixture\n\ngo 1.22\n")
	nextWrite(t, root, "SYSTEM.md", "# Fixture\nSYSTEM_CONTEXT_MARKER\n")
	nextWrite(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}", %q]
  boundaries: [sh, .fireweed/boundaries.sh]
agent:
  command: [bash, %q, %q, %q, "{prompt}"]
audit:
  command: [bash, %q, %q, "{prompt}"]
timeout: 50ms
`, logDir, worker, logDir, workerMode, auditor, logDir))
	nextWrite(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ -f packages/alpha/evals/staged_test.go ] && [ ! -f "$3/evals-observed" ]; then
  git rev-parse HEAD > "$3/evals-reset-head"
  git diff --cached --name-only > "$3/evals-staged"
  git diff --name-only > "$3/evals-unstaged"
  git ls-files --others --exclude-standard > "$3/evals-untracked"
  touch "$3/evals-observed"
fi
go test ./packages/alpha/evals > .fireweed/go-test.log 2>&1
code=$?
if [ "$code" -eq 0 ]; then result=''; else result='<failure message="Go eval failed"/>'; cat .fireweed/go-test.log; fi
printf '<testsuite><testcase name="ALP-001" classname="fixture"/><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" > "$2"
exit "$code"
`)
	nextWrite(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	nextWrite(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n## Decisions\n")
	nextWrite(t, root, "packages/alpha/api.go", "package alpha\nimport \"example.invalid/next-fixture/packages/alpha/internal\"\nfunc Value() string { return internal.Value() }\n")
	nextWrite(t, root, "packages/alpha/internal/value.go", "package internal\n// FULL_INTERNAL_SECRET\nfunc Value() string { return \"old\" }\n")
	nextWrite(t, root, "packages/alpha/evals/existing_test.go", "package evals\nimport (\"testing\"; \"example.invalid/next-fixture/packages/alpha\")\nfunc Test_ALP_001_Value(t *testing.T) { if alpha.Value() == \"\" { t.Fatal(\"empty\") } }\n")
	nextGit(t, root, "init", "-q")
	nextGit(t, root, "checkout", "-q", "-b", "main")
	nextGit(t, root, "config", "user.name", "Fixture")
	nextGit(t, root, "config", "user.email", "fixture@example.invalid")
	if queued && !approved {
		nextCommit(t, root, "base fixture")
		nextGit(t, root, "checkout", "-q", "-b", "feature")
	}
	if queued {
		nextWrite(t, root, nextChange, "---\nmodule: alpha\nbases: {}\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	}
	nextCommit(t, root, "approved fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("bad fixture: %#v", d)
		}
	}
	return nextFixture{root, logDir}
	result, err := runner.Next(context.Background(), request)
	if err != nil || result.Empty || result.State.Change.Slug != "gamma" || result.State.Phase != "applied" {
		t.Errorf("Next past waiting beta = %#v, %v", result, err)
	}
	assertRunBetaWaiting(t, f)
}

func Test_RUN_010_EntirelyWaitingQueue(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	request := runner.NextRequest{Root: f.root, Bin: "/fixture/run-bin", Timeout: 20 * time.Second}
	blocked, blockErr := runner.Next(context.Background(), request)
	if blockErr == nil || !strings.Contains(blockErr.Error(), "P2-01: no fail-first case for ALP-002") || blocked.State.Change.Slug != "alpha" || blocked.State.Phase != "evals" {
		t.Fatalf("fixture alpha did not block: %#v, %v", blocked, blockErr)
	}
	if err := os.Remove(filepath.Join(f.root, "changes/0003-gamma.md")); err != nil {
		t.Fatal(err)
	}
	nextCommit(t, f.root, "leave only waiting beta")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	invocations := runInvocations(t, f)
	eligible, err := workflow.Eligible(f.root)
	if err != nil || len(eligible) != 0 {
		t.Fatalf("fixture entirely waiting queue = %#v, %v", eligible, err)
	}
	empty, err := runner.Next(context.Background(), request)
	if err != nil || !empty.Empty || len(empty.Steps) != 0 {
		t.Errorf("Next with entirely waiting queue = %#v, %v", empty, err)
	}
	if nextGit(t, f.root, "rev-parse", "HEAD") != before || strings.Join(runInvocations(t, f), "\n") != strings.Join(invocations, "\n") {
		t.Error("entirely waiting Next changed HEAD or invoked an agent")
	}
}

func Test_RUN_013_RUN_014_CompleteSpecChange(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err != nil {
		t.Fatalf("Next: %#v, %v", r, err)
	}
	if r.Empty || r.State.Phase != "applied" || r.State.Change.Slug != "next" || r.State.Run == "" || r.State.ClaimOID != "" || r.State.Snapshot.Root != f.root {
		t.Errorf("applied result = %#v", r)
	}
	if len(r.Steps) != 3 {
		t.Fatalf("steps = %#v", r.Steps)
	}
	for i, want := range []struct{ phase, role, stdout string }{{"evals", "worker", "worker-evals"}, {"audit", "audit", "auditor"}, {"impl", "worker", "worker-impl"}} {
		got := r.Steps[i]
		if got.Phase != want.phase || got.Role != want.role || got.Agent.Stdout != want.stdout || got.Agent.ExitCode != 0 {
			t.Errorf("step %d = %#v", i, got)
		}
	}
	lines := strings.Split(strings.TrimSpace(nextRead(t, f.logDir, "invocations")), "\n")
	if len(lines) != 3 || lines[0] != "evals|"+f.root+"|/fixture/next-bin" || !strings.HasPrefix(lines[1], "audit|") || !strings.HasSuffix(lines[1], "|/fixture/next-bin") || lines[2] != "impl|"+f.root+"|/fixture/next-bin" {
		t.Errorf("invocations = %#v", lines)
	}
	assertContent := func(prompt, category, content string) {
		t.Helper()
		if !strings.Contains(strings.Join(strings.Fields(prompt), " "), strings.Join(strings.Fields(content), " ")) {
			t.Errorf("%s omitted content %q", category, content)
		}
	}
	assertPlaybook := func(prompt, phase string) {
		t.Helper()
		playbook, err := assets.Playbook(phase)
		if err != nil {
			t.Fatalf("load %s playbook: %v", phase, err)
		}
		// Check the public playbook's instructions without requiring its headings
		// or the context command's presentation and whitespace.
		for _, line := range strings.Split(playbook, "\n") {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				assertContent(prompt, phase+" playbook", line)
			}
		}
	}
	for _, phase := range []string{"evals", "impl"} {
		prompt := nextRead(t, f.logDir, "worker-"+phase+".prompt")
		for _, want := range []struct{ category, content string }{
			{"change", "One new value is needed."},
			{"change interface description", "Existing Value function."},
			{"added spec", "Alpha MUST return the new value."},
			// This retained spec occurs only in the projected spec, not the change.
			{"projected spec", "Alpha MUST return its value."},
			{"interface contents", "func Value() string { return internal.Value() }"},
			{"SYSTEM", "SYSTEM_CONTEXT_MARKER"},
			{"other implementation contents", "FULL_INTERNAL_SECRET"},
			{"other eval contents", "func Test_ALP_001_Value"},
		} {
			assertContent(prompt, phase+" worker "+want.category, want.content)
		}
		assertPlaybook(prompt, phase)
		if phase == "impl" {
			assertContent(prompt, "impl worker added eval contents", "EVAL_CONTENT_MARKER")
		}
	}
	audit := nextRead(t, f.logDir, "audit.prompt")
	for _, marker := range []string{nextChange, "One new value is needed.", "Alpha MUST return the new value.", "Alpha MUST return its value.", "ALP-002", "EVAL_CONTENT_MARKER", "func Test_ALP_001_Value"} {
		assertContent(audit, "audit context", marker)
	}
	assertPlaybook(audit, "audit")
	for _, secret := range []string{"FULL_INTERNAL_SECRET", "SYSTEM_CONTEXT_MARKER"} {
		if strings.Contains(audit, secret) {
			t.Errorf("audit context exposed %q", secret)
		}
	}
	if _, err := os.Stat(filepath.Join(f.root, nextChange)); !os.IsNotExist(err) {
		t.Errorf("pending change remains: %v", err)
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	for _, marker := range []string{"phase: applied", "### ALP-002", "No findings", "- evals ", "- audit ", "- impl ", "agent=bash"} {
		if !strings.Contains(archived, marker) {
			t.Errorf("archive omitted %q", marker)
		}
	}
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/SPEC.md"), "Alpha MUST return the new value.") {
		t.Error("SPEC not applied")
	}
	if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("claim retained: %s", refs)
	}
	if count := nextGit(t, f.root, "rev-list", "--count", before+"..HEAD"); count != "3" {
		t.Errorf("phase commit count = %s", count)
	}
	if old := nextGit(t, f.root, "show", "HEAD~2:packages/alpha/internal/value.go"); !strings.Contains(old, `return "old"`) {
		t.Errorf("evals commit changed implementation: %q", old)
	}
	if first := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~2"); !strings.Contains(first, "evals") {
		t.Errorf("first phase commit = %q", first)
	}
	if second := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~1"); !strings.Contains(second, "audit") {
		t.Errorf("second phase commit = %q", second)
	}
	if applied := nextGit(t, f.root, "show", "HEAD:changes/archive/0001-next.md"); !strings.Contains(applied, "phase: applied") {
		t.Errorf("third commit did not archive applied state: %q", applied)
	}
	failReport := nextRead(t, f.root, ".fireweed/runs/"+r.State.Run+"/evals/alpha.json")
	passReport := nextRead(t, f.root, ".fireweed/runs/"+r.State.Run+"/impl/alpha.json")
	if !strings.Contains(failReport, "ALP-002") || !strings.Contains(failReport, "fail") || !strings.Contains(passReport, "ALP-002") || !strings.Contains(passReport, "pass") {
		t.Errorf("phase reports omitted genuine fail-first/pass: evals=%q impl=%q", failReport, passReport)
	}
	if _, err := workflow.Current(f.root, "next"); err == nil {
		t.Error("applied change remains current")
	}
}

func Test_RUN_015_InvocationFailureRetainsHeldPhase(t *testing.T) {
	f := newNextFixture(t, true, true, "exit")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err == nil || len(r.Steps) != 1 {
		t.Errorf("invocation failure: error=%v steps=%d, want 1", err, len(r.Steps))
	}
	for i, step := range r.Steps {
		if step.Phase != "evals" || step.Role != "worker" || step.Agent.Stdout != "partial worker output" || step.Agent.ExitCode != 9 || step.Agent.Stderr != "worker failure" {
			t.Errorf("captured invocation %d = %#v", i, step)
		}
	}
	if got := len(nextInvocations(t, f)); got != 1 {
		t.Errorf("agent calls = %d, want 1", got)
	}
	state, currentErr := workflow.Current(f.root, "next")
	if currentErr != nil || state.Phase != "evals" || state.Run == "" || r.State.Phase != "evals" || r.State.Run != state.Run || nextGit(t, f.root, "rev-parse", "HEAD") != before {
		t.Errorf("invocation failure advanced phase: currentErr=%v currentPhase=%q resultPhase=%q runPreserved=%t", currentErr, state.Phase, r.State.Phase, r.State.Run == state.Run)
	}
	if claim := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); claim != "refs/fireweed/claims/next" || state.ClaimOID == "" {
		t.Errorf("invocation failure lost claim: ref=%q claimOIDPresent=%t", claim, state.ClaimOID != "")
	}
}

func nextAssertExhaustedBlock(t *testing.T, f nextFixture, r runner.NextResult, err error, phase, before string, wantPhaseCommits int) string {
	t.Helper()
	phaseHead := strings.TrimSpace(nextRead(t, f.logDir, phase+"-phase-head"))
	if finalHead := strings.TrimSpace(nextRead(t, f.logDir, phase+"-final-head")); finalHead != phaseHead {
		t.Errorf("fixture %s HEAD changed across attempts: first=%s final=%s", phase, phaseHead, finalHead)
	}
	if claim := strings.TrimSpace(nextRead(t, f.logDir, phase+"-claim")); claim != "refs/fireweed/claims/next" {
		t.Errorf("fixture lacked held claim at %s: %q", phase, claim)
	}
	if got := nextGit(t, f.root, "rev-list", "--count", before+".."+phaseHead); got != fmt.Sprint(wantPhaseCommits) {
		t.Errorf("fixture %s phase commits = %s, want %d", phase, got, wantPhaseCommits)
	}
	held := nextRead(t, f.logDir, phase+"-held-change")
	if !strings.Contains(held, "phase: "+phase) || !strings.Contains(held, "run: "+r.State.Run) || !strings.Contains(held, "### ALP-002") {
		t.Error("fixture held phase lacked expected phase, run or assigned ID")
	}
	if parents := strings.Fields(nextGit(t, f.root, "rev-list", "--parents", "-n", "1", "HEAD")); len(parents) != 2 || parents[1] != phaseHead {
		t.Errorf("block commit parents = %d, want held %s HEAD", len(parents)-1, phase)
	}
	if paths := nextGit(t, f.root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); paths != nextChange {
		t.Errorf("block commit paths = %q, want state only", paths)
	}
	if status := nextGit(t, f.root, "status", "--porcelain", "-uall"); status != "" {
		t.Errorf("blocked checkout dirty: %q", status)
	}
	if claim := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); claim != "" {
		t.Errorf("blocked claim retained: %q", claim)
	}
	if _, currentErr := workflow.Current(f.root, "next"); currentErr == nil {
		t.Error("blocked change still held")
	}
	blocked := nextRead(t, f.root, nextChange)
	if !strings.Contains(blocked, "phase: "+phase) || !strings.Contains(blocked, "run: "+r.State.Run) || !strings.Contains(blocked, "### ALP-002") || !strings.Contains(blocked, "blocked:") {
		t.Error("block document lost phase, run, assigned ID or reason")
	}
	reason, _ := r.State.Change.Document.Frontmatter["blocked"].(string)
	if r.State.Phase != phase || r.State.Run == "" || len(r.State.Change.Added) != 1 || r.State.Change.Added[0].ID != "ALP-002" || r.State.ClaimOID != "" || r.State.Head != nextGit(t, f.root, "rev-parse", "HEAD") || r.State.Snapshot.Root != f.root || strings.TrimSpace(reason) == "" {
		t.Errorf("returned block state: phase=%q runPresent=%t assignedID=%t claimReleased=%t headMatches=%t snapshotMatches=%t blocked=%q", r.State.Phase, r.State.Run != "", len(r.State.Change.Added) == 1 && r.State.Change.Added[0].ID == "ALP-002", r.State.ClaimOID == "", r.State.Head == nextGit(t, f.root, "rev-parse", "HEAD"), r.State.Snapshot.Root == f.root, reason)
	}
	ref := "refs/fireweed/attempts/next/1"
	if err == nil || !strings.Contains(err.Error(), ref) {
		t.Errorf("exhaustion error does not identify %s: %v", ref, err)
	}
	if oid := nextGit(t, f.root, "for-each-ref", "--format=%(objectname)", ref); oid == "" {
		t.Errorf("saved attempt ref absent: %s", ref)
		return ""
	} else if kind := nextGit(t, f.root, "cat-file", "-t", oid); kind != "commit" {
		t.Errorf("attempt ref kind = %q, want commit", kind)
	}
	return ref
}

func Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks(t *testing.T) {
	f := newNextFixture(t, true, true, "noedit")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") {
		t.Errorf("evals exhaustion cause = %v", err)
	}
	if len(r.Steps) != 2 {
		t.Errorf("evals exhaustion steps = %d, want 2", len(r.Steps))
	}
	for i, step := range r.Steps {
		if step.Phase != "evals" || step.Role != "worker" || step.Agent.Stdout != "no edit" || step.Agent.ExitCode != 0 || step.Agent.Stderr != "" {
			t.Errorf("captured invocation %d = %#v", i, step)
		}
	}
	if got := len(nextInvocations(t, f)); got != 2 {
		t.Errorf("agent calls = %d, want 2", got)
	}
	ref := nextAssertExhaustedBlock(t, f, r, err, "evals", before, 0)
	if ref != "" {
		if saved := nextGit(t, f.root, "show", ref+":"+nextChange); !strings.Contains(saved, "phase: evals") || !strings.Contains(saved, "### ALP-002") || strings.Contains(saved, "blocked:") {
			t.Error("saved evals attempt lost held unblocked change")
		}
	}
}

func Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase(t *testing.T) {
	f := newNextFixture(t, true, true, "retry-both")
	r, err := f.run(false)
	if err != nil {
		t.Fatalf("Next after retry = %#v, %v", r, err)
	}
	if r.State.Phase != "applied" || len(r.Steps) != 5 {
		t.Fatalf("retry result = %#v", r)
	}
	for i, want := range []struct{ phase, role, output string }{
		{"evals", "worker", "worker-evals-1"},
		{"evals", "worker", "worker-evals"},
		{"audit", "audit", "auditor"},
		{"impl", "worker", "worker-impl-1"},
		{"impl", "worker", "worker-impl"},
	} {
		step := r.Steps[i]
		if step.Phase != want.phase || step.Role != want.role || step.Agent.Stdout != want.output || step.Agent.ExitCode != 0 {
			t.Errorf("step %d = %#v", i, step)
		}
	}
	if got := len(nextInvocations(t, f)); got != 5 {
		t.Errorf("agent calls = %d, want 5", got)
	}
	if got := nextRead(t, f.logDir, "impl-retention"); got != "retained\n" {
		t.Errorf("work seen by second impl agent = %q", got)
	}
	if got := nextRead(t, f.logDir, "evals-retention"); got != "retained\n" {
		t.Errorf("work seen by second evals agent = %q", got)
	}
	for path, marker := range map[string]string{
		"packages/alpha/evals/retained_test.go": "RETAINED_EVAL_WORK",
		"packages/alpha/internal/retained.go":   "RETAINED_IMPL_WORK",
	} {
		if !strings.Contains(nextRead(t, f.root, path), marker) {
			t.Errorf("retained work missing from %s", path)
		}
	}
	for _, phase := range []struct{ name, feedback string }{
		{"evals", "P2-01: no fail-first case for ALP-002"},
		{"impl", "E01: failed case"},
	} {
		prompt := nextRead(t, f.logDir, "worker-"+phase.name+"-2.prompt")
		if !strings.Contains(prompt, phase.feedback) {
			t.Errorf("%s retry prompt lacks real verification feedback %q", phase.name, phase.feedback)
		}
	}
	if prompt := nextRead(t, f.logDir, "worker-impl-2.prompt"); !strings.Contains(prompt, "--- FAIL: Test_ALP_002_NewValue") {
		t.Error("impl retry prompt lacks the Go test failure output")
	}
}

package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"slices"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/runner"
)

// Hooks are tracked in the fixture repository. The fake worker and its
// observations live outside the fixture repository.
type drillFixture struct{ root, observations, base string }

func newDrillFixture(t *testing.T) drillFixture {
	t.Helper()
	root, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	observations, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	agent := filepath.Join(observations, "agent.sh")
	wrapper := filepath.Join(observations, "wrapper.sh")
	nextWrite(t, root, ".fireweed/setup.sh", `#!/bin/sh
mkdir -p .fireweed/bin
if [ -d .git ]; then
  printf '%s\n%s\n' "$PWD" "$FIREWEED_BIN" > "$1/input.setup"
  printf 'INPUT_SETUP' > .fireweed/bin/fireweed
  exit 0
fi
printf '%s\n%s\n' "$PWD" "$FIREWEED_BIN" > "$1/setup"
git symbolic-ref -q HEAD > "$1/setup.branch" || :
git rev-parse HEAD > "$1/setup.head"
printf 'SETUP_BINARY_SECRET' > .fireweed/bin/fireweed
if [ -f "$1/setup-fail" ]; then echo SETUP_FAILURE >&2; exit 17; fi
if [ -f "$1/cleanup-fail" ]; then chmod 0500 "$(dirname "$PWD")"; fi
`)
	nextWrite(t, observations, "wrapper.sh", `#!/bin/bash
printf '%s\n' "$FIREWEED_BIN" >> "$1/wrapper.bin"
shift
exec "$@"
`)
	nextWrite(t, observations, "agent.sh", `#!/bin/bash
count=0
[ ! -f "$1/count" ] || count=$(cat "$1/count")
count=$((count + 1))
printf '%s' "$count" > "$1/count"
: > "$1/invoked-$count"
printf '%s' "$PWD" > "$1/agent.cwd"
printf '%s' "$FIREWEED_BIN" > "$1/agent.bin"
printf '%s' "$2" > "$1/prompt-$count"
if [ "$count" -eq 1 ]; then
  find . -path './.git' -prune -o ! -type d -print | sort > "$1/input.files"
  git for-each-ref --format='%(refname)' > "$1/input.refs"
  git rev-list --count HEAD > "$1/input.history"
  if git cat-file -e "$(cat "$1/base")^{commit}" 2>/dev/null; then echo present > "$1/input.base"; else echo absent > "$1/input.base"; fi
  if [ -f "$1/saved-attempt" ]; then
    if git cat-file -e "$(cat "$1/saved-attempt")^{commit}" 2>/dev/null; then echo present; else echo absent; fi > "$1/input.saved-attempt"
  fi
  if [ -e .git/objects/info/alternates ]; then echo present > "$1/input.alternates"; else echo absent > "$1/input.alternates"; fi
  if [ -f .fireweed/bin/fireweed ]; then cp .fireweed/bin/fireweed "$1/input.setup-output"; fi
fi
mode=success
[ ! -f "$1/mode" ] || mode=$(cat "$1/mode")
mkdir -p packages/alpha/internal
case "$mode" in
  timeout) printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; sleep 10;;
  usage_always) printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'RATE LIMIT' >&2; exit 9;;
  usage_once) if [ "$count" -eq 1 ]; then printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'RATE LIMIT' >&2; exit 9; fi; [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" = PARTIAL_WORK ] || { echo 'prior retry work missing' >&2; exit 8; }; printf 'new' > packages/alpha/internal/value.go;;
  agent_retry|mixed_exhaust) if [ "$count" -eq 1 ]; then printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'WORKER_FAILURE' >&2; exit 7; fi; if [ "$mode" = agent_retry ]; then [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" = PARTIAL_WORK ] || { echo 'prior retry work missing' >&2; exit 8; }; printf 'new' > packages/alpha/internal/value.go; fi;;
  agent_fail_always) printf 'PARTIAL_WORK' > packages/alpha/internal/value.go; echo 'WORKER_FAILURE' >&2; exit 7;;
  verify_retry) if [ "$count" -eq 1 ]; then printf 'wrong' > packages/alpha/internal/value.go; exit 0; fi; [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" = wrong ] || { echo 'prior retry work missing' >&2; exit 8; }; printf 'new' > packages/alpha/internal/value.go;;
  wrong|mixed_exhaust) printf 'wrong' > packages/alpha/internal/value.go;;
  *) printf 'new' > packages/alpha/internal/value.go;;
esac
stray=''
case "$mode" in
  change_api) printf 'corrupted' > packages/alpha/api.go;;
  delete_eval) rm packages/alpha/evals/alpha_test.go;;
  add_foreign) stray=packages/core/internal/foreign.go;;
  add_other) stray=packages/alpha/extra.txt;;
  change_preserved) printf 'bad' > packages/alpha/internal/retained.go;;
  add_project_protected) stray=packages/alpha/internal/project-new.txt;;
  change_protected_dependency_interface) printf 'bad' > packages/core/api.go;;
  change_setup|delete_setup)
    [ "$(cat .fireweed/bin/fireweed)" = INPUT_SETUP ] || exit 12
    if [ "$mode" = change_setup ]; then printf 'bad' > .fireweed/bin/fireweed; else rm .fireweed/bin/fireweed; fi;;
  self_check)
    # Read the advertised shell command from a fenced block, inline code, or a bare line.
    command=$(awk '
      /^[[:space:]]*\140\140\140/ {
        fenced = !fenced
        if (!fenced && block ~ /[.]fireweed\/eval[.]sh/) { print block; exit }
        block = ""
        next
      }
      fenced { block = block $0 "\n"; next }
      /[.]fireweed\/eval[.]sh/ {
        if (match($0, /\140[^\140]*[.]fireweed\/eval[.]sh[^\140]*\140/))
          print substr($0, RSTART + 1, RLENGTH - 2)
        else
          print
        exit
      }
    ' "$1/prompt-$count")
    [ -n "$command" ] || { echo 'prompt omitted the local eval command' >&2; exit 12; }
    printf 'wrong' > packages/alpha/internal/value.go
    if output=$(sh -c "$command" 2>&1); then echo 'local eval passed the wrong candidate' >&2; exit 12; fi
    case "$output" in
      *'own verification failure'*) printf 'new' > packages/alpha/internal/value.go;;
      *) printf '%s\n' "$output" >&2; exit 12;;
    esac
    sh -c "$command" || exit 12;;
esac
if [ -n "$stray" ]; then
  mkdir -p "$(dirname "$stray")" && printf 'stray' > "$stray" || exit 12
  printf '%s\n' "$stray" > "$1/worker.artifact"
fi
cat > DRILL_NOTES.md <<'NOTES'
## Contradicted Decisions
- ALP-001: ordered preference may be contradicted.

## Gaps
- ALP-001, module alpha: ordering needs a spec.
- no spec, module aardvark: caller expectation may need a spec.
NOTES
printf 'worker observation'
`)
	nextWrite(t, root, ".fireweed/eval.sh", `#!/bin/sh
report=$3
observations=$4
root=$(pwd -P)
if [ "$(cd "$2" && pwd -P)" != "$root/packages/$1" ]; then echo invalid module directory >&2; exit 25; fi
if [ ! -f go.mod ] || [ ! -f packages/core/internal/secret.txt ]; then echo build inputs missing >&2; exit 25; fi
if [ -d .git ]; then
  [ "$1" = alpha ] && [ "$(cat .fireweed/bin/fireweed)" = INPUT_SETUP ] || exit 25
  case "$report" in ''|*'{report}'*) echo unexpanded local report path >&2; exit 25;; esac
  mkdir -p "$(dirname "$report")" || exit 25
  report_file="$(cd "$(dirname "$report")" && pwd -P)/$(basename "$report")"
  case "$report_file" in
    "$root/"*) ;;
    *) echo local report outside input workspace >&2; exit 25;;
  esac
  printf '%s\n' "${report_file#"$root"/}" > "$observations/local-eval.report"
  printf '%s\n' "$PWD" >> "$observations/local-eval.cwd"
else
  printf '%s\n' "$1" >> "$observations/eval.order"
  printf '%s' "$PWD" > "$observations/eval.cwd"
  if [ ! -f packages/aardvark/internal/secret.txt ] || [ ! -f packages/zeta/internal/secret.txt ]; then echo dependent internals missing >&2; exit 25; fi
  for observation in "$observations/worker.artifact" "$observations/local-eval.report"; do
    if [ -f "$observation" ] && [ -e "$(cat "$observation")" ]; then
      echo "unexpected file in verifier: $(cat "$observation")" >&2
      exit 25
    fi
  done
fi
case "$1" in alpha) id=ALP-001;; aardvark) id=AAR-001;; zeta) id=ZET-001;; *) exit 26;; esac
if [ "$1" = alpha ] && [ -f "$observations/omit-own-coverage" ]; then printf '<testsuite><testcase name="ALP-001_case"/></testsuite>' > "$report"; exit 0; fi
if [ "$1" = aardvark ] && [ -f "$observations/omit-dependent-coverage" ]; then printf '<testsuite><testcase name="AAR-001_case"/></testsuite>' > "$report"; exit 0; fi
fail=''
if [ "$1" = alpha ] && [ "$(cat packages/alpha/internal/value.go 2>/dev/null)" != new ]; then fail='own verification failure'; fi
if [ "$1" != alpha ] && [ -f "$observations/fail-dependents" ]; then fail="dependent $1 failure"; fi
if [ -z "$fail" ]; then printf '<testsuite><testcase name="%s_case" classname="fixture"/></testsuite>' "$id" > "$report"; exit $?; fi
printf '<testsuite><testcase name="%s_case" classname="fixture"><failure message="%s"/></testcase></testsuite>' "$id" "$fail" > "$report" || exit 25
printf '%s\n' "$fail" >&2
exit 1
`)
	nextWrite(t, root, ".gitignore", ".fireweed/bin/\n.fireweed/drills/\n.fireweed/runs/\n")
	nextWrite(t, root, "SYSTEM.md", "# Fixture guidance\nSYSTEM_RETAINED\n")
	nextWrite(t, root, "AGENTS.md", "AGENT_INSTRUCTION_RETAINED\n")
	nextWrite(t, root, "go.mod", "module example.invalid/drill-fixture\n\ngo 1.22\n")
	nextWrite(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  preserve: [internal/retained.go]
  disposable: [internal/**]
hooks:
  setup: [sh, .fireweed/setup.sh, %q]
  eval: [sh, .fireweed/eval.sh, "{module}", "{module_dir}", "{report}", %q]
agent:
  command: [bash, %q, %q, "{prompt}"]
  wrap: [bash, %q, %q]
  model: fixture-model
attempts: 2
timeout: 2s
limit_patterns: [RATE LIMIT]
limit_wait: 2ms
limit_budget: 4ms
verify_dependents: false
`, observations, observations, agent, observations, wrapper, observations))
	for _, module := range []struct{ name, prefix, deps string }{
		{"foundation", "FOU", "[]"}, {"core", "COR", "[foundation]"}, {"alpha", "ALP", "[core]"},
		{"aardvark", "AAR", "[alpha]"}, {"zeta", "ZET", "[aardvark]"},
	} {
		nextWrite(t, root, "packages/"+module.name+"/SPEC.md", fmt.Sprintf("---\nname: %s\nid_prefix: %s\ndepends_on: %s\n---\n# %s\n\n## Specs\n\n### %s-001\n%s MUST preserve its value.\n\n## Decisions\n\n[ordering] Prefer ordered values.\n", module.name, module.prefix, module.deps, module.name, module.prefix, module.name))
		nextWrite(t, root, "packages/"+module.name+"/api.go", "package "+module.name+"\n// INTERFACE_"+module.name+"\n")
		nextWrite(t, root, "packages/"+module.name+"/evals/"+module.name+"_test.go", "package evals\n// EVAL_"+module.name+"\n")
		nextWrite(t, root, "packages/"+module.name+"/internal/secret.txt", "INTERNAL_SECRET_"+module.name+"\n")
	}
	nextWrite(t, root, "packages/alpha/internal/value.go", "OLD_IMPLEMENTATION_SECRET")
	nextWrite(t, root, "packages/alpha/internal/retained.go", "RETAINED_OVERLAP")
	nextWrite(t, root, "packages/alpha/other.txt", "RETAINED_OTHER")
	nextGit(t, root, "init", "-q")
	nextGit(t, root, "checkout", "-q", "-b", "main")
	nextGit(t, root, "config", "user.name", "Fixture")
	nextGit(t, root, "config", "user.email", "fixture@example.invalid")
	nextCommit(t, root, "drill fixture")
	base := nextGit(t, root, "rev-parse", "HEAD")
	nextWrite(t, observations, "base", base)
	return drillFixture{root, observations, base}
}

func (f drillFixture) run(ctx context.Context, timeout time.Duration) (runner.DrillResult, error) {
	return runner.Drill(ctx, runner.DrillRequest{Root: f.root, Module: "alpha", Bin: "/fixture/new-controller", Timeout: timeout})
}

exec
/bin/zsh -lc "rg -n 'func .*Evaluate|func .*RequirePass|func .*hook|module_dir|timeout' packages/checks/internal/engine --glob '*.go' | head -90; sed -n '1130,1170p' packages/workflow/internal/lifecycle/stub.go; sed -n '1,125p' packages/workflow/internal/lifecycle/gc.go; rg -n 'Validate.*[Cc]onfig|limit_wait|limit_budget' packages/project; git ls-files -s | head; .fireweed/bin/fireweed help verify | head -45" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "sed -n '430,610p' packages/runner/evals/next_test.go; sed -n '318,451p' packages/runner/evals/audit_workspace_test.go; cat packages/runner/evals/retry_failure_details_test.go | tail -125; sed -n '255,320p;345,410p;620,647p;802,860p' packages/runner/evals/drill_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 684ms:
	blocked, blockErr := runner.Next(context.Background(), request)
	if blockErr == nil || !strings.Contains(blockErr.Error(), "P2-01: no fail-first case for ALP-002") || blocked.State.Change.Slug != "alpha" || blocked.State.Phase != "evals" {
		t.Fatalf("fixture alpha did not block: %#v, %v", blocked, blockErr)
	}
	if err := os.Remove(filepath.Join(f.root, "changes/0003-gamma.md")); err != nil {
		t.Fatal(err)
	}
	nextCommit(t, f.root, "leave only waiting beta")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	invocations := runInvocations(t, f)
	eligible, err := workflow.Eligible(f.root)
	if err != nil || len(eligible) != 0 {
		t.Fatalf("fixture entirely waiting queue = %#v, %v", eligible, err)
	}
	empty, err := runner.Next(context.Background(), request)
	if err != nil || !empty.Empty || len(empty.Steps) != 0 {
		t.Errorf("Next with entirely waiting queue = %#v, %v", empty, err)
	}
	if nextGit(t, f.root, "rev-parse", "HEAD") != before || strings.Join(runInvocations(t, f), "\n") != strings.Join(invocations, "\n") {
		t.Error("entirely waiting Next changed HEAD or invoked an agent")
	}
}

func Test_RUN_013_RUN_014_CompleteSpecChange(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err != nil {
		t.Fatalf("Next: %#v, %v", r, err)
	}
	if r.Empty || r.State.Phase != "applied" || r.State.Change.Slug != "next" || r.State.Run == "" || r.State.ClaimOID != "" || r.State.Snapshot.Root != f.root {
		t.Errorf("applied result = %#v", r)
	}
	if len(r.Steps) != 3 {
		t.Fatalf("steps = %#v", r.Steps)
	}
	for i, want := range []struct{ phase, role, stdout string }{{"evals", "worker", "worker-evals"}, {"audit", "audit", "auditor"}, {"impl", "worker", "worker-impl"}} {
		got := r.Steps[i]
		if got.Phase != want.phase || got.Role != want.role || got.Agent.Stdout != want.stdout || got.Agent.ExitCode != 0 {
			t.Errorf("step %d = %#v", i, got)
		}
	}
	lines := strings.Split(strings.TrimSpace(nextRead(t, f.logDir, "invocations")), "\n")
	if len(lines) != 3 || lines[0] != "evals|"+f.root+"|/fixture/next-bin" || !strings.HasPrefix(lines[1], "audit|") || !strings.HasSuffix(lines[1], "|/fixture/next-bin") || lines[2] != "impl|"+f.root+"|/fixture/next-bin" {
		t.Errorf("invocations = %#v", lines)
	}
	assertContent := func(prompt, category, content string) {
		t.Helper()
		if !strings.Contains(strings.Join(strings.Fields(prompt), " "), strings.Join(strings.Fields(content), " ")) {
			t.Errorf("%s omitted content %q", category, content)
		}
	}
	assertPlaybook := func(prompt, phase string) {
		t.Helper()
		playbook, err := assets.Playbook(phase)
		if err != nil {
			t.Fatalf("load %s playbook: %v", phase, err)
		}
		// Check the public playbook's instructions without requiring its headings
		// or the context command's presentation and whitespace.
		for _, line := range strings.Split(playbook, "\n") {
			line = strings.TrimSpace(line)
			if line != "" && !strings.HasPrefix(line, "#") {
				assertContent(prompt, phase+" playbook", line)
			}
		}
	}
	for _, phase := range []string{"evals", "impl"} {
		prompt := nextRead(t, f.logDir, "worker-"+phase+".prompt")
		for _, want := range []struct{ category, content string }{
			{"change", "One new value is needed."},
			{"change interface description", "Existing Value function."},
			{"added spec", "Alpha MUST return the new value."},
			// This retained spec occurs only in the projected spec, not the change.
			{"projected spec", "Alpha MUST return its value."},
			{"interface contents", "func Value() string { return internal.Value() }"},
			{"SYSTEM", "SYSTEM_CONTEXT_MARKER"},
			{"other implementation contents", "FULL_INTERNAL_SECRET"},
			{"other eval contents", "func Test_ALP_001_Value"},
		} {
			assertContent(prompt, phase+" worker "+want.category, want.content)
		}
		assertPlaybook(prompt, phase)
		if phase == "impl" {
			assertContent(prompt, "impl worker added eval contents", "EVAL_CONTENT_MARKER")
		}
	}
	audit := nextRead(t, f.logDir, "audit.prompt")
	for _, marker := range []string{nextChange, "One new value is needed.", "Alpha MUST return the new value.", "Alpha MUST return its value.", "ALP-002", "EVAL_CONTENT_MARKER", "func Test_ALP_001_Value"} {
		assertContent(audit, "audit context", marker)
	}
	assertPlaybook(audit, "audit")
	for _, secret := range []string{"FULL_INTERNAL_SECRET", "SYSTEM_CONTEXT_MARKER"} {
		if strings.Contains(audit, secret) {
			t.Errorf("audit context exposed %q", secret)
		}
	}
	if _, err := os.Stat(filepath.Join(f.root, nextChange)); !os.IsNotExist(err) {
		t.Errorf("pending change remains: %v", err)
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	for _, marker := range []string{"phase: applied", "### ALP-002", "No findings", "- evals ", "- audit ", "- impl ", "agent=bash"} {
		if !strings.Contains(archived, marker) {
			t.Errorf("archive omitted %q", marker)
		}
	}
	if !strings.Contains(nextRead(t, f.root, "packages/alpha/SPEC.md"), "Alpha MUST return the new value.") {
		t.Error("SPEC not applied")
	}
	if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("claim retained: %s", refs)
	}
	if count := nextGit(t, f.root, "rev-list", "--count", before+"..HEAD"); count != "3" {
		t.Errorf("phase commit count = %s", count)
	}
	if old := nextGit(t, f.root, "show", "HEAD~2:packages/alpha/internal/value.go"); !strings.Contains(old, `return "old"`) {
		t.Errorf("evals commit changed implementation: %q", old)
	}
	if first := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~2"); !strings.Contains(first, "evals") {
		t.Errorf("first phase commit = %q", first)
	}
	if second := nextGit(t, f.root, "show", "--format=%s", "--no-patch", "HEAD~1"); !strings.Contains(second, "audit") {
		t.Errorf("second phase commit = %q", second)
	}
	if applied := nextGit(t, f.root, "show", "HEAD:changes/archive/0001-next.md"); !strings.Contains(applied, "phase: applied") {
		t.Errorf("third commit did not archive applied state: %q", applied)
	}
	failReport := nextRead(t, f.root, ".fireweed/runs/"+r.State.Run+"/evals/alpha.json")
	passReport := nextRead(t, f.root, ".fireweed/runs/"+r.State.Run+"/impl/alpha.json")
	if !strings.Contains(failReport, "ALP-002") || !strings.Contains(failReport, "fail") || !strings.Contains(passReport, "ALP-002") || !strings.Contains(passReport, "pass") {
		t.Errorf("phase reports omitted genuine fail-first/pass: evals=%q impl=%q", failReport, passReport)
	}
	if _, err := workflow.Current(f.root, "next"); err == nil {
		t.Error("applied change remains current")
	}
}

func Test_RUN_015_InvocationFailureRetainsHeldPhase(t *testing.T) {
	f := newNextFixture(t, true, true, "exit")
	before := nextGit(t, f.root, "rev-parse", "HEAD")
	r, err := f.run(false)
	if err == nil || len(r.Steps) != 1 {
		t.Errorf("invocation failure: error=%v steps=%d, want 1", err, len(r.Steps))
	}
	for i, step := range r.Steps {
		if step.Phase != "evals" || step.Role != "worker" || step.Agent.Stdout != "partial worker output" || step.Agent.ExitCode != 9 || step.Agent.Stderr != "worker failure" {
			t.Errorf("captured invocation %d = %#v", i, step)
		}
	}
	if got := len(nextInvocations(t, f)); got != 1 {
		t.Errorf("agent calls = %d, want 1", got)
	}
	state, currentErr := workflow.Current(f.root, "next")
	if currentErr != nil || state.Phase != "evals" || state.Run == "" || r.State.Phase != "evals" || r.State.Run != state.Run || nextGit(t, f.root, "rev-parse", "HEAD") != before {
		t.Errorf("invocation failure advanced phase: currentErr=%v currentPhase=%q resultPhase=%q runPreserved=%t", currentErr, state.Phase, r.State.Phase, r.State.Run == state.Run)
	}
	if claim := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims"); claim != "refs/fireweed/claims/next" || state.ClaimOID == "" {
		t.Errorf("invocation failure lost claim: ref=%q claimOIDPresent=%t", claim, state.ClaimOID != "")
	}
}

func nextAssertExhaustedBlock(t *testing.T, f nextFixture, r runner.NextResult, err error, phase, before string, wantPhaseCommits int) string {
	t.Helper()
	phaseHead := strings.TrimSpace(nextRead(t, f.logDir, phase+"-phase-head"))
	if finalHead := strings.TrimSpace(nextRead(t, f.logDir, phase+"-final-head")); finalHead != phaseHead {
		t.Errorf("fixture %s HEAD changed across attempts: first=%s final=%s", phase, phaseHead, finalHead)
	}
	if claim := strings.TrimSpace(nextRead(t, f.logDir, phase+"-claim")); claim != "refs/fireweed/claims/next" {
		t.Errorf("fixture lacked held claim at %s: %q", phase, claim)
	}
	if got := nextGit(t, f.root, "rev-list", "--count", before+".."+phaseHead); got != fmt.Sprint(wantPhaseCommits) {
		t.Errorf("fixture %s phase commits = %s, want %d", phase, got, wantPhaseCommits)
	}
	held := nextRead(t, f.logDir, phase+"-held-change")
	if !strings.Contains(held, "phase: "+phase) || !strings.Contains(held, "run: "+r.State.Run) || !strings.Contains(held, "### ALP-002") {
		t.Error("fixture held phase lacked expected phase, run or assigned ID")
	}
	if parents := strings.Fields(nextGit(t, f.root, "rev-list", "--parents", "-n", "1", "HEAD")); len(parents) != 2 || parents[1] != phaseHead {
		t.Errorf("block commit parents = %d, want held %s HEAD", len(parents)-1, phase)
	}
	if paths := nextGit(t, f.root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); paths != nextChange {
		t.Fatalf("two-change audit lifecycle: applied=%d error=%v", len(r.Applied), err)
	}
	var cwds []string
	for _, line := range runInvocations(t, f) {
		parts := strings.Split(line, "|")
		if len(parts) == 4 && parts[1] == "audit" {
			cwds = append(cwds, parts[2])
		}
	}
	if len(cwds) != 2 {
		t.Fatalf("audit invocation roots = %#v", cwds)
	}
	if cwds[0] == cwds[1] || cwds[0] == f.root || cwds[1] == f.root {
		t.Errorf("audits shared a workspace or used held checkout: %#v", cwds)
	}
	for _, cwd := range cwds {
		if _, err := os.Stat(cwd); !os.IsNotExist(err) {
			t.Errorf("audit workspace %q remains: %v", cwd, err)
		}
	}
}

func Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically(t *testing.T) {
	for _, tc := range []struct{ name, action string }{
		{"change outside Audit", `sed 's/One new value is needed./Altered motivation./' changes/0001-next.md > "$1/altered-change"; cp "$1/altered-change" changes/0001-next.md`},
		{"projected spec", `printf '\nAltered spec\n' >> packages/alpha/SPEC.md`},
		{"eval", `printf '\n// altered eval\n' >> packages/alpha/evals/next_test.go`},
		{"added file", `printf 'extra\n' > unexpected.txt`},
		{"deleted file", `rm packages/alpha/evals/existing_test.go`},
		{"committed invalid edit", `printf '\nAltered spec\n' >> packages/alpha/SPEC.md; git add -A; git -c user.name=Fixture -c user.email=fixture@example.invalid commit -q -m invalid-audit-edit`},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			nextAttempts(t, f, 1)
			spec := nextRead(t, f.root, "packages/alpha/SPEC.md")
			existing := nextRead(t, f.root, "packages/alpha/evals/existing_test.go")
			auditAction(t, f, tc.action+"\n")
			r, err := f.run(false)
			if err == nil || len(r.Steps) != 2 || r.State.Phase != "audit" || r.State.ClaimOID == "" {
				t.Errorf("invalid audit was not rejected while retaining claim: phase=%q claim=%q steps=%d error=%v", r.State.Phase, r.State.ClaimOID, len(r.Steps), err)
			}
			heldBytes, readErr := os.ReadFile(filepath.Join(f.root, nextChange))
			if readErr != nil {
				t.Fatalf("rejected edit removed held change: %v", readErr)
			}
			held := string(heldBytes)
			if strings.Contains(held, "No findings") || strings.Contains(held, "Altered motivation") || !strings.Contains(held, "One new value is needed.") {
				t.Errorf("rejected edit reached held change: %q", held)
			}
			if got := nextRead(t, f.root, "packages/alpha/SPEC.md"); got != spec {
				t.Error("rejected audit edit reached held spec")
			}
			if got := nextRead(t, f.root, "packages/alpha/evals/existing_test.go"); got != existing {
				t.Error("rejected audit edit reached held eval")
			}
			if strings.Contains(nextRead(t, f.root, "packages/alpha/evals/next_test.go"), "altered eval") {
				t.Error("rejected audit edit reached held new eval")
			}
			if _, statErr := os.Stat(filepath.Join(f.root, "unexpected.txt")); !os.IsNotExist(statErr) {
				t.Errorf("rejected file reached held checkout: %v", statErr)
			}
			if got := nextInvocations(t, f); len(got) != 2 {
				t.Errorf("invalid audit advanced phase: %#v", got)
			}
			auditWorkspaceGone(t, f)
		})
	}
}

func Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops(t *testing.T) {
	for _, tc := range []struct {
		name, action, policy string
		timeStop             bool
	}{
		{"agent failure", "printf 'partial audit'; exit 9\n", "", false},
		{"usage stop", "printf 'RATE LIMIT' >&2; exit 9\n", "limit_patterns: [RATE LIMIT]\nlimit_budget: 0s\n", false},
		{"time stop", "sleep 6\n", "", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := newNextFixture(t, true, true, "normal")
			if tc.policy != "" {
				nextWrite(t, f.root, "fireweed.yaml", nextRead(t, f.root, "fireweed.yaml")+tc.policy)
				nextCommit(t, f.root, "audit limit policy")
			}
			if tc.timeStop {
				fastAuditHook(t, f)
			}
			auditAction(t, f, tc.action)
			base := nextGit(t, f.root, "rev-parse", "HEAD")
			var r runner.NextResult
			var err error
			if tc.timeStop {
				result, runErr := runner.Run(context.Background(), runner.RunRequest{Root: f.root, Bin: "/fixture/next-bin", For: 3 * time.Second, Timeout: 10 * time.Second})
				err = runErr
				if len(result.Applied) != 0 {
					t.Errorf("time stop applied %d changes", len(result.Applied))
				}
			} else {
				r, err = f.run(false)
			}
			if err == nil || tc.timeStop && !runner.IsTimeStop(err) || tc.name == "usage stop" && !runner.IsUsageStop(err) {
				t.Errorf("audit failure/stop = %v", err)
			}
			if !tc.timeStop && (len(r.Steps) != 2 || r.State.Phase != "audit") {
				t.Errorf("audit failure result = %#v", r)
			}
			held := nextRead(t, f.root, nextChange)
			if !strings.Contains(held, "No findings") || !strings.Contains(held, "One new value is needed.") || !strings.Contains(held, "phase: audit") {
				t.Errorf("partial Audit not retained in held change: %q", held)
			}
			if got := nextGit(t, f.root, "rev-list", "--count", base+"..HEAD"); got != "1" {
				t.Errorf("audit error advanced phases %s times, want only evals", got)
			}
			if state, stateErr := workflow.Current(f.root, "next"); stateErr != nil || state.Phase != "audit" || state.ClaimOID == "" {
				t.Errorf("audit error released or advanced phase: %#v, %v", state, stateErr)
			}
			if got := nextInvocations(t, f); len(got) != 2 {
				t.Errorf("audit error advanced phase: %#v", got)
			}
			auditWorkspaceGone(t, f)
		})
	}
}

func fastAuditHook(t *testing.T, f nextFixture) {
	t.Helper()
	nextWrite(t, f.root, ".fireweed/eval.sh", `#!/bin/sh
if [ -f packages/alpha/evals/next_test.go ]; then
  printf '<testsuite><testcase name="ALP-001" classname="fixture"/><testcase name="ALP-002" classname="fixture"><failure message="not implemented"/></testcase></testsuite>' > "$2"
  exit 1
fi
exit 30
`)
	nextCommit(t, f.root, "fast audit fixture hook")
	nextWrite(t, f.root, ".fireweed/eval.sh", hook)
	nextCommit(t, f.root, "fixture with dependent JUnit failures")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid retry-detail fixture: %#v", d)
		}
	}
	return f
}

func retryDetailSection(t *testing.T, f nextFixture) (string, string) {
	t.Helper()
	first := nextRead(t, f.logDir, "worker-impl-1.prompt")
	retry := nextRead(t, f.logDir, "worker-impl-2.prompt")
	if !strings.HasPrefix(retry, first) {
		t.Fatal("retry did not retain the original worker context")
	}
	const marker = retryDetailHookMarker + "\n"
	i := strings.LastIndex(retry, marker)
	if i < len(first) {
		t.Fatalf("retry prompt lost dependent hook stderr marker: %q", retry[len(first):])
	}
	return retry[len(first):], retry[i+len(marker):]
}

func Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules(t *testing.T) {
	alpha := `<testsuite>
<testcase name="ALP-001_pass_only" classname="pass.class"/>
<testcase name="ALP-002_first" classname="alpha.first"><failure message="RUN043_ALPHA_FIRST_ASSERTION"/></testcase>
<testcase name="ALP-002_skip_only" classname="skip.class"><skipped/></testcase>
<testcase name="ALP-002_second" classname="alpha.second"><error>RUN043_ALPHA_SECOND_ASSERTION</error></testcase>
</testsuite>`
	dependent := `<testsuite>
<testcase name="AAR-001_dependent" classname="aardvark.dependent"><failure message="RUN043_DEPENDENT_ASSERTION"/></testcase>
<testcase name="AAR-001_pass_only" classname="aardvark.pass"/>
</testsuite>`
	f := retryDetailFixture(t, alpha, dependent)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 4 {
		t.Fatalf("Next with JUnit-only dependent failure = %#v, %v", r, err)
	}
	feedback, section := retryDetailSection(t, f)
	if !strings.Contains(feedback, "E01: failed case") || !strings.Contains(feedback, retryDetailHookMarker) {
		t.Error("retry lost existing verification feedback or hook stderr")
	}
	for _, want := range []string{
		"aardvark", "AAR-001_dependent", "aardvark.dependent", "RUN043_DEPENDENT_ASSERTION",
		"alpha", "ALP-002_first", "alpha.first", "RUN043_ALPHA_FIRST_ASSERTION",
		"ALP-002_second", "alpha.second", "RUN043_ALPHA_SECOND_ASSERTION",
	} {
		if !strings.Contains(section, want) {
			t.Errorf("failed-case section lacks %q: %q", want, section)
		}
	}
	for _, absent := range []string{"ALP-001_pass_only", "ALP-002_skip_only", "AAR-001_pass_only"} {
		if strings.Contains(section, absent) {
			t.Errorf("nonfailed case %q included in detail section", absent)
		}
	}
	previous := -1
	for _, name := range []string{"AAR-001_dependent", "ALP-002_first", "ALP-002_second"} {
		i := strings.Index(section, name)
		if i <= previous {
			t.Errorf("failed cases are not in module and report order: %q", section)
			break
		}
		previous = i
	}
}

func Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions(t *testing.T) {
	var alpha strings.Builder
	alpha.WriteString(`<testsuite><testcase name="ALP-001"/>`)
	for i := 0; i < 12; i++ {
		fill := strings.Repeat("界", 5000)
		fmt.Fprintf(&alpha, `<testcase name="ALP-002_long_%02d" classname="large.case"><failure>RUN044_HEAD_%02d_%s_RUN044_TAIL_%02d</failure></testcase>`, i, i, fill, i)
	}
	alpha.WriteString(`</testsuite>`)
	dependent := `<testsuite><testcase name="AAR-001_first" classname="first.case"><failure>RUN044_DEPENDENT_FIRST</failure></testcase></testsuite>`
	f := retryDetailFixture(t, alpha.String(), dependent)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 4 {
		t.Fatalf("Next with large JUnit failures = %#v, %v", r, err)
	}
	_, section := retryDetailSection(t, f)
	if n := utf8.RuneCountInString(section); n > 16384 {
		t.Errorf("failed-case section has %d Unicode code points, want at most 16384", n)
	}
	if !strings.Contains(section, "RUN044_HEAD_00_") || !strings.Contains(section, strings.Repeat("界", 4000)) {
		t.Error("first failure did not receive a readable Unicode excerpt")
	}
	for i := 0; i < 12; i++ {
		head := fmt.Sprintf("RUN044_HEAD_%02d_", i)
		start := strings.Index(section, head)
		if start < 0 {
			continue // This case was omitted under the section budget.
		}
		fillRunes := 0
		for _, r := range section[start+len(head):] {
			if r != '界' {
				break
			}
			fillRunes++
		}
		if n := utf8.RuneCountInString(head) + fillRunes; n > 4096 {
			t.Errorf("case %d failure excerpt contains at least %d Unicode code points, want at most 4096", i, n)
		}
		if strings.Contains(section, fmt.Sprintf("RUN044_TAIL_%02d", i)) {
			t.Errorf("case %d failure excerpt continued past its first 4096 code points", i)
		}
	}
	if !strings.Contains(section, "RUN044_DEPENDENT_FIRST") || strings.Index(section, "RUN044_DEPENDENT_FIRST") >= strings.Index(section, "RUN044_HEAD_00_") {
		t.Error("lexically first dependent failure was omitted or misordered")
	}
	if strings.Contains(section, "ALP-002_long_11") {
		t.Error("section exceeded its budget instead of omitting later cases")
	}
	lower := strings.ToLower(section)
	if !strings.Contains(lower, "truncat") && !strings.Contains(lower, "shorten") && !strings.Contains(section, "…") {
		t.Error("shortened diagnostic text has no marker")
	}
	if !strings.Contains(lower, "omitt") && !strings.Contains(lower, "remaining") && !strings.Contains(lower, "more cases") && !strings.Contains(lower, "additional cases") {
		t.Error("omitted cases have no marker")
	}
}
func drillNoSetupOrAgent(t *testing.T, f drillFixture) {
	t.Helper()
	for _, path := range []string{"setup", "input.setup", "count"} {
		if _, err := os.Stat(filepath.Join(f.observations, path)); !os.IsNotExist(err) {
			t.Errorf("%s exists: %v", path, err)
		}
	}
}

func drillCleaned(t *testing.T, f drillFixture) {
	t.Helper()
	for _, path := range []string{"setup", "input.setup", "agent.cwd"} {
		b, err := os.ReadFile(filepath.Join(f.observations, path))
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			t.Fatal(err)
		}
		cwd := strings.TrimSpace(strings.SplitN(string(b), "\n", 2)[0])
		if cwd == f.root || cwd == "" {
			t.Errorf("%s cwd = %q", path, cwd)
			continue
		}
		if _, err := os.Stat(cwd); !os.IsNotExist(err) {
			t.Errorf("%s workspace remains: %v", cwd, err)
		}
	}
	if out := nextGit(t, f.root, "worktree", "list", "--porcelain"); strings.Count(out, "worktree ") != 1 {
		t.Errorf("scratch worktree remains: %q", out)
	}
}

func Test_RUN_050_DrillPreflightRejectsWithoutEffects(t *testing.T) {
	for _, mode := range []string{"unknown", "dirty", "claim"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			request := runner.DrillRequest{Root: f.root, Module: "alpha", Bin: "/fixture/new-controller"}
			switch mode {
			case "unknown":
				request.Module = "missing"
			case "dirty":
				nextWrite(t, f.root, "stray.txt", "dirty")
			case "claim":
				nextGit(t, f.root, "update-ref", "refs/fireweed/claims/live", f.base)
			}
			before := nextGit(t, f.root, "for-each-ref", "--format=%(refname) %(objectname)", "refs/fireweed")
			_, err := runner.Drill(context.Background(), request)
			if err == nil {
				t.Fatal("preflight accepted " + mode)
			}
			want := map[string]string{"unknown": "missing", "dirty": "dirty", "claim": "claim"}[mode]
			if !strings.Contains(strings.ToLower(err.Error()), want) {
				t.Errorf("%s preflight returned unrelated error: %v", mode, err)
			}
			if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "for-each-ref", "--format=%(refname) %(objectname)", "refs/fireweed") != before {
				t.Error("preflight changed invoking refs")
			}
			drillNoSetupOrAgent(t, f)
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_051_SetupRunsDetachedAtHeadAndFailureStops(t *testing.T) {
	f := newDrillFixture(t)
func Test_RUN_052_RUN_053_RestrictedInputsAndIndependentGit(t *testing.T) {
	f := newDrillFixture(t)
	nextWrite(t, f.root, "changes/archive/prior.md", "prior change\n")
	nextWrite(t, f.root, ".fireweed/drills/alpha.md", "prior drill report\n")
	nextGit(t, f.root, "add", "-f", ".fireweed/drills/alpha.md")
	nextCommit(t, f.root, "tracked changes and reports")
	f.base = nextGit(t, f.root, "rev-parse", "HEAD")
	nextWrite(t, f.observations, "base", f.base)
	nextGit(t, f.root, "branch", "invoking-only")
	attempt := nextGit(t, f.root, "commit-tree", "HEAD^{tree}", "-m", "saved attempt")
	nextGit(t, f.root, "update-ref", "refs/fireweed/attempts/sentinel", attempt)
	nextWrite(t, f.observations, "saved-attempt", attempt)
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" {
		t.Fatalf("drill = %#v, %v", r, err)
	}
	got := strings.Fields(drillObserved(t, f, "input.files"))
	want := []string{
		"./.gitignore", "./.fireweed/bin/fireweed", "./.fireweed/eval.sh", "./.fireweed/setup.sh",
		"./AGENTS.md", "./SYSTEM.md", "./fireweed.yaml", "./go.mod",
		"./packages/alpha/SPEC.md", "./packages/alpha/api.go", "./packages/alpha/evals/alpha_test.go",
		"./packages/alpha/internal/retained.go", "./packages/alpha/other.txt",
		"./packages/core/SPEC.md", "./packages/core/api.go", "./packages/core/evals/core_test.go", "./packages/core/internal/secret.txt",
		"./packages/foundation/SPEC.md", "./packages/foundation/api.go", "./packages/foundation/evals/foundation_test.go", "./packages/foundation/internal/secret.txt",
	}
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("reconstruction inputs = %q, want %q", got, want)
	}
	for path, want := range map[string]string{
		"input.base": "absent", "input.saved-attempt": "absent", "input.alternates": "absent", "input.history": "1",
	} {
		if got := strings.TrimSpace(drillObserved(t, f, path)); got != want {
			t.Errorf("%s = %q, want %q", path, got, want)
		}
	}
	if refs := drillObserved(t, f, "input.refs"); strings.Contains(refs, "invoking-only") || strings.Contains(refs, "fireweed/") {
		t.Errorf("input Git contains invoking refs: %q", refs)
	}
	agentRoot := drillObserved(t, f, "agent.cwd")
	scratchRoot := strings.SplitN(drillObserved(t, f, "setup"), "\n", 2)[0]
	if agentRoot == f.root || agentRoot == scratchRoot {
		t.Errorf("agent used %q instead of a separate workspace", agentRoot)
	}
	if got, readErr := os.ReadFile(filepath.Join(f.observations, "input.setup")); readErr != nil || string(got) != agentRoot+"\n/fixture/new-controller\n" {
		t.Errorf("input setup = %q, %v, want setup in %s before the worker", got, readErr, agentRoot)
	}
	if got, readErr := os.ReadFile(filepath.Join(f.observations, "input.setup-output")); readErr != nil || string(got) != "INPUT_SETUP" {
		t.Errorf("worker saw setup output %q, %v", got, readErr)
	}
	if got := nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/retained.go"); got != "RETAINED_OVERLAP" {
		t.Errorf("preserve precedence lost: %q", got)
	}
	if out := nextGit(t, f.root, "ls-tree", "-r", "--name-only", r.Ref, "packages/alpha/internal/secret.txt"); out != "" {
		t.Errorf("target disposable secret survived deletion: %q", out)
	}
	for _, module := range []string{"foundation", "core", "aardvark", "zeta"} {
		if got := nextGit(t, f.root, "show", r.Ref+":packages/"+module+"/internal/secret.txt"); got != "INTERNAL_SECRET_"+module {
			t.Errorf("%s verifier internal changed: %q", module, got)
		}
	}
	drillCleaned(t, f)
}

func Test_RUN_054_PromptLetsWorkerRunOwnEval(t *testing.T) {
	f := newDrillFixture(t)
func Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "wrong")
	nextWrite(t, f.observations, "fail-dependents", "yes")
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result != "fail" || len(r.Steps) != 2 {
		t.Errorf("failed verification = %#v, %v", r, err)
	}
	if len(r.Own.Cases) != 1 || r.Own.Cases[0].Result != "fail" || !strings.Contains(r.Own.Cases[0].Failure, "own verification failure") {
		t.Errorf("own failure details lost: %#v", r.Own)
	}
	if len(r.Dependents) != 2 || r.Dependents["aardvark"].Cases[0].Result != "fail" || r.Dependents["zeta"].Cases[0].Result != "fail" {
		t.Errorf("dependent reports lost: %#v", r.Dependents)
	}
	if got := strings.Fields(drillObserved(t, f, "eval.order")); !reflect.DeepEqual(got, []string{"alpha", "aardvark", "zeta", "alpha", "aardvark", "zeta"}) {
		t.Errorf("own and transitive dependent order = %q", got)
	}
	drillCleaned(t, f)

	dependent := newDrillFixture(t)
	nextWrite(t, dependent.observations, "fail-dependents", "yes")
	depResult, depErr := dependent.run(context.Background(), 0)
	if depErr == nil || depResult.Result != "fail" || depResult.Own.Cases[0].Result != "pass" || depResult.Dependents["aardvark"].Cases[0].Result != "fail" || depResult.Dependents["zeta"].Cases[0].Result != "fail" {
		t.Errorf("dependent-only failures passed or lost reports: %#v, %v", depResult, depErr)
	}
	drillCleaned(t, dependent)
}

func Test_RUN_062_CleanupOnSuccessAndFailures(t *testing.T) {
	for _, mode := range []string{"success", "agent_retry", "change_api", "usage_always", "timeout"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			f.mode(t, mode)
			if mode == "timeout" {
				_, _ = f.run(context.Background(), 150*time.Millisecond)
			} else {
				_, _ = f.run(context.Background(), 0)
			}
			for _, path := range []string{"setup", "agent.cwd"} {
				if _, err := os.Stat(filepath.Join(f.observations, path)); err != nil {
					t.Errorf("%s scenario did not reach %s: %v", mode, path, err)
				}
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_062_CleanupFailureCannotPass(t *testing.T) {
	f := newDrillFixture(t)
	nextWrite(t, f.observations, "cleanup-fail", "yes")
	common := nextGit(t, f.root, "rev-parse", "--git-common-dir")
	if !filepath.IsAbs(common) {
		common = filepath.Join(f.root, common)
	}
	parent := filepath.Join(common, "fireweed", "worktrees")
	defer func() {
		_ = os.Chmod(parent, 0755)
		setup, err := os.ReadFile(filepath.Join(f.observations, "setup"))
		if err != nil {
			return
		}
		scratch := strings.TrimSpace(strings.SplitN(string(setup), "\n", 2)[0])
		if _, err := os.Stat(scratch); err == nil {
			nextGit(t, f.root, "worktree", "remove", "--force", scratch)
		}
	}()
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result == "pass" || len(r.Steps) != 1 {
		t.Errorf("cleanup failure reported a passing drill: %#v, %v", r, err)
	}
	setup := strings.SplitN(drillObserved(t, f, "setup"), "\n", 2)[0]
	if _, statErr := os.Stat(setup); statErr != nil {
		t.Errorf("fixture did not make scratch removal fail: %v", statErr)
	}
}

 succeeded in 739ms:
packages/checks/internal/engine/doctor.go:66:func (d *doctorState) fail(name, module, hook, detail string) {
packages/checks/internal/engine/doctor.go:344:		v := map[string]string{"name": d.names[i], "prefix": d.prefixes[i], "package": d.names[i], "go_module": d.goModule, "module_dir": d.dirs[i], "other_dir": d.dirs[1-i]}
packages/checks/internal/engine/doctor.go:394:	return map[string]string{"name": d.names[i], "prefix": d.prefixes[i], "package": d.names[i], "go_module": d.goModule, "module_dir": d.dirs[i], "other_dir": d.dirs[1-i]}
packages/checks/internal/engine/doctor.go:907:	timeout := d.req.Timeout
packages/checks/internal/engine/doctor.go:908:	if timeout <= 0 {
packages/checks/internal/engine/doctor.go:909:		timeout = 3 * time.Minute
packages/checks/internal/engine/doctor.go:911:	p, e := gitstore.Run(d.ctx, gitstore.Request{Args: args, Dir: d.root, Timeout: timeout})
packages/checks/internal/engine/junit.go:193:func RequirePass(report model.Report, specs []contracts.Item) error {
packages/checks/internal/engine/hooks.go:25:var placeholder = regexp.MustCompile(`\{(?:module_dir|module|report)\}`)
packages/checks/internal/engine/hooks.go:31:func hookArgs(command any, values map[string]string) ([]string, error) {
packages/checks/internal/engine/hooks.go:74:func runHook(ctx context.Context, root string, command any, values map[string]string, bin string, timeout time.Duration) (gitstore.Result, error) {
packages/checks/internal/engine/hooks.go:79:	result, err := gitstore.Run(ctx, gitstore.Request{Args: args, Dir: root, Env: map[string]string{"FIREWEED_BIN": bin}, Timeout: timeout})
packages/checks/internal/engine/hooks.go:105:func Evaluate(ctx context.Context, request model.HookRequest) (model.Report, error) {
packages/checks/internal/engine/hooks.go:133:	values := map[string]string{"module": module.Name, "module_dir": module.Dir, "report": reportPath}
packages/checks/internal/engine/boundaries.go:76:		values := map[string]string{"module": hook.module.Name, "module_dir": hook.module.Dir}
	_ = filepath.WalkDir(file(s.Snapshot.Root, m.Dir), func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		rel, _ := filepath.Rel(s.Snapshot.Root, path)
		rel = filepath.ToSlash(rel)
		if rel == m.Dir+"/SPEC.md" {
			return nil
		}
		for _, p := range m.Zones.Interface {
			r := strings.TrimPrefix(rel, m.Dir+"/")
			if contracts.MatchGlob(p, r) {
				return nil
			}
		}
		info, e := d.Info()
		if e != nil {
			return nil
		}
		b.WriteString(fmt.Sprintf("- %s (%d bytes)\n", rel, info.Size()))
		if full {
			data, e := os.ReadFile(path)
			if e == nil {
				b.Write(data)
				b.WriteByte('\n')
			}
		}
		return nil
	})
	return b.String(), nil
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
		if os.IsNotExist(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		for _, entry := range entries {
			name := entry.Name()
			if entry.IsDir() || !commitChangeName.MatchString(name) {
				continue
			}
			slug := strings.TrimSuffix(name[strings.IndexByte(name, '-')+1:], ".md")
			if directory == "changes/archive" {
				archived[slug] = true
			} else {
				delete(archived, slug)
			}
		}
	}
	return archived, nil
}

func gcClaimStale(root string, entry gcRef, registered map[string]bool, common string) bool {
	data, err := gitstore.ReadBlob(root, entry.oid)
	if err != nil {
		return false
	}
	var claim struct {
		Worktree string `json:"worktree"`
	}
	if json.Unmarshal(data, &claim) != nil || !filepath.IsAbs(claim.Worktree) {
		return false
	}
100644 89dbb7995f18c8ad1dcb8329b58aabf25f92b9ad 0	.fireweed/canaries/canaries.yaml
100644 b7121740deae6b6d0057bee4c603e16183504dea 0	.fireweed/hooks.go
100644 a7de391faaa790847617d021ed2e1aa6d6fe8cd3 0	.fireweed/recipe.md
100644 34a9d4b347a7e8d607fe4fcee6b45e76931d47ea 0	.fireweed/setup.sh
100644 fe50ac34d7e01aa8fad735bcdf5bcd6d0865b238 0	.fireweed/templates/module/SPEC.md
100644 267aa31649d5f9932e3ad06e193f72e15b5c993a 0	.fireweed/templates/module/api.go.tmpl
100644 e69de29bb2d1d6434b8b29ae775ad8c2e48c5391 0	.fireweed/templates/module/evals/.gitkeep
100644 e69de29bb2d1d6434b8b29ae775ad8c2e48c5391 0	.fireweed/templates/module/internal/.gitkeep
100644 f4987b04cd86adcff6e3a66a62f0324685db4229 0	.gitignore
100644 700aa760e06daa76e53fc0b021026b38bd650d0e 0	AGENTS.md
U01: help takes no arguments. Use help alone.

exec
/bin/zsh -lc "sed -n '325,452p' packages/runner/evals/run_test.go; sed -n '215,290p' packages/runner/evals/audit_workspace_test.go; sed -n '23,166p' packages/checks/internal/engine/hooks.go; rg -n 'func (Section|ReplaceSection)|func section' packages/contracts/internal/primitives/primitives.go; sed -n '1,110p' packages/runner/evals/agent_logs_test.go; .fireweed/bin/fireweed help | rg -n 'verify|context'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 553ms:
	}
	assertRunApplied(t, f, r.Applied, "alpha")
}

func Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding(t *testing.T) {
	f := newRunFixture(t, 3, true, false)
	first, err := runner.Run(context.Background(), f.request(1, true))
	if err != nil {
		t.Fatalf("first Run: %#v, %v", first, err)
	}
	if first.Empty || !first.Limited {
		t.Errorf("first flags = %#v", first)
	}
	assertRunApplied(t, f, first.Applied, "alpha")
	firstArchive := runArchive(t, f, 0)
	q, err := project.Queue(f.root)
	if err != nil || len(q) != 2 || q[0].Slug != "beta" || q[1].Slug != "gamma" {
		t.Errorf("remaining queue = %#v, %v", q, err)
	}
	second, err := runner.Run(context.Background(), f.request(0, true))
	if err != nil {
		t.Fatalf("continuation: %#v, %v", second, err)
	}
	if !second.Empty || second.Limited {
		t.Errorf("continuation flags = %#v", second)
	}
	assertRunApplied(t, f, second.Applied, "beta", "gamma")
	assertRunAgentLog(t, f, "alpha", "beta", "gamma")
	if archive := runArchive(t, f, 0); archive != firstArchive {
		t.Error("first archive changed during continuation")
	}
	if q, err := project.Queue(f.root); err != nil || len(q) != 0 {
		t.Errorf("remaining queue = %#v, %v", q, err)
	}
}

func Test_RUN_018_InitialEmptyAndExactLimit(t *testing.T) {
	t.Run("initial_empty", func(t *testing.T) {
		f := newRunFixture(t, 0, false, false)
		before := nextGit(t, f.root, "rev-parse", "HEAD")
		r, err := runner.Run(context.Background(), f.request(2, false))
		if err != nil || !r.Empty || r.Limited || len(r.Applied) != 0 {
			t.Errorf("empty Run = %#v, %v", r, err)
		}
		if nextGit(t, f.root, "rev-parse", "HEAD") != before || len(runInvocations(t, f)) != 0 {
			t.Error("empty Run changed repo or invoked agents")
		}
	})
	t.Run("exact_limit", func(t *testing.T) {
		f := newRunFixture(t, 2, false, false)
		r, err := runner.Run(context.Background(), f.request(2, false))
		if err != nil {
			t.Fatalf("exact Run: %#v, %v", r, err)
		}
		if !r.Empty || r.Limited {
			t.Errorf("exact flags = %#v", r)
		}
		assertRunApplied(t, f, r.Applied, "alpha", "beta")
		assertRunAgentLog(t, f, "alpha", "beta")
	})
}

func Test_RUN_020_HeldFailureRetainsEarlierResults(t *testing.T) {
	f := newRunFixture(t, 3, false, true)
	r, err := runner.Run(context.Background(), f.request(0, false))
	if err == nil || !strings.Contains(err.Error(), "17") {
		t.Errorf("held failure cause = %v", err)
	}
	assertRunApplied(t, f, r.Applied, "alpha")
	state, currentErr := workflow.Current(f.root, "beta")
	if currentErr != nil || state.Phase != "evals" || state.Run == "" {
		t.Errorf("held beta = %#v, %v", state, currentErr)
	}
	if _, statErr := os.Stat(filepath.Join(f.root, "changes", "0002-beta.md")); statErr != nil {
		t.Errorf("beta no longer held: %v", statErr)
	}
	if _, statErr := os.Stat(filepath.Join(f.root, "changes", "0003-gamma.md")); statErr != nil {
		t.Errorf("gamma no longer queued: %v", statErr)
	}
	if _, statErr := os.Stat(filepath.Join(f.root, "changes", "archive", "0003-gamma.md")); !os.IsNotExist(statErr) {
		t.Errorf("gamma archived: %v", statErr)
	}
	if strings.Contains(nextRead(t, f.root, "changes/0003-gamma.md"), "phase:") {
		t.Error("gamma was started")
	}
	got := runInvocations(t, f)
	if len(got) != 4 || got[3] != "beta|evals|"+f.root+"|/fixture/run-bin" {
		t.Errorf("failure invocation order = %#v", got)
	}
}

func Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	r, err := runner.Run(context.Background(), f.request(0, false))
	assertRunBlockedAlpha(t, f, err)
	assertRunApplied(t, f, r.Applied, "gamma")
	assertRunBetaWaiting(t, f)
	if q, queueErr := project.Queue(f.root); queueErr != nil || len(q) != 2 || q[0].Slug != "alpha" || q[1].Slug != "beta" {
		t.Errorf("remaining queue = %#v, %v", q, queueErr)
	}
	if got := runInvocations(t, f); len(got) != 5 || got[0] != "alpha|evals|"+f.root+"|/fixture/run-bin" || got[1] != got[0] || got[2] != "gamma|evals|"+f.root+"|/fixture/run-bin" || !strings.HasPrefix(got[3], "gamma|audit|") || !strings.HasSuffix(got[3], "|/fixture/run-bin") || got[4] != "gamma|impl|"+f.root+"|/fixture/run-bin" {
		t.Errorf("blocked continuation agent calls = %#v", got)
	}
}

func Test_RUN_025_RUN_026_PositiveMaxRetainsBlock(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	nextWrite(t, f.root, "changes/0004-later.md", "---\nmodule: gamma\nbases: {}\n---\n\n## Motivation\nLater independent work.\n\n## Interface\n\n## Added\n\n### new:later\nGamma MUST return another value.\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	nextCommit(t, f.root, "later independent work")
	for _, d := range project.Check(f.root) {
		if !d.Warning {
			t.Fatalf("invalid Max fixture: %#v", d)
		}
	}
	r, err := runner.Run(context.Background(), f.request(1, false))
	assertRunBlockedAlpha(t, f, err)
	assertRunApplied(t, f, r.Applied, "gamma")
	assertRunBetaWaiting(t, f)
	if later := nextRead(t, f.root, "changes/0004-later.md"); strings.Contains(later, "run:") || strings.Contains(later, "phase:") {
		t.Errorf("Max started later work: %q", later)
	}
	if got := runInvocations(t, f); len(got) != 5 {
		t.Errorf("Max agent calls = %#v, want exhausted alpha then one applied gamma", got)
	}
}
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	const inserted = "## Audit\nNo findings\nChecked IDs: ALP-002\n\n"
	auditBeforeEdit(t, f, `if [ -f "$1/audit.original" ]; then
  cp changes/0001-next.md "$1/audit.retry-input"
  printf auditor
  exit 0
fi
cp changes/0001-next.md "$1/audit.original"
cp "$1/audit.cwd" "$1/first-audit.cwd"
awk '
  /^## Provenance$/ && !inserted {
    printf "## Audit\nNo findings\nChecked IDs: ALP-002\n\n"
    inserted = 1
  }
  { print }
  END { if (!inserted) exit 8 }
' "$1/audit.original" > changes/0001-next.md || exit 8
cp changes/0001-next.md "$1/audit.candidate"
printf auditor
exit 9
`)
	first, firstErr := f.run(false)
	if firstErr == nil || first.State.Phase != "audit" || first.State.ClaimOID == "" || len(first.Steps) != 2 {
		t.Fatalf("first audit failure = %#v, %v", first, firstErr)
	}
	auditWorkspaceGone(t, f)
	r, err := f.run(false)
	steps := append(first.Steps, r.Steps...)
	if err != nil || r.State.Phase != "applied" || len(steps) != 4 {
		t.Fatalf("audit before Provenance lifecycle = %#v, %v", r, err)
	}
	for i, want := range []struct {
		phase, role string
		exit        int
	}{{"evals", "worker", 0}, {"audit", "audit", 9}, {"audit", "audit", 0}, {"impl", "worker", 0}} {
		if step := steps[i]; step.Phase != want.phase || step.Role != want.role || step.Agent.ExitCode != want.exit {
			t.Errorf("step %d = %#v, want %#v", i, step, want)
		}
	}
	if cwd := nextRead(t, f.logDir, "first-audit.cwd"); cwd == nextRead(t, f.logDir, "audit.cwd") {
		t.Errorf("audit retry reused workspace %q", cwd)
	}
	original := nextRead(t, f.logDir, "audit.original")
	candidate := nextRead(t, f.logDir, "audit.candidate")
	if !strings.Contains(original, "## Provenance\n") || strings.Count(candidate, inserted) != 1 ||
		strings.Replace(candidate, inserted, "", 1) != original ||
		strings.Index(candidate, "## Audit\n") > strings.Index(candidate, "## Provenance\n") {
		t.Errorf("auditor did not insert only the first Audit before Provenance")
	}
	beforeProvenance, hasOriginalProvenance := contracts.Section(original, "Provenance")
	retryInput := nextRead(t, f.logDir, "audit.retry-input")
	retryProvenance, hasRetryProvenance := contracts.Section(retryInput, "Provenance")
	if !hasOriginalProvenance || !hasRetryProvenance || retryProvenance != beforeProvenance {
		t.Errorf("Audit copying changed raw Provenance before successful Done: got %q, want %q", retryProvenance, beforeProvenance)
	}
	if audit, ok := contracts.Section(retryInput, "Audit"); !ok || !strings.Contains(audit, "No findings\nChecked IDs: ALP-002") {
		t.Errorf("failed auditor's Audit was not imported before retry: %q", audit)
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	for _, section := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Decisions"} {
		want, _ := contracts.Section(original, section)
		got, _ := contracts.Section(archived, section)
		if got != want {
			t.Errorf("held %s changed outside Audit: got %q, want %q", section, got, want)
		}
	}
	afterProvenance, hasProvenance := contracts.Section(archived, "Provenance")
	nonblankLines := func(section string) []string {
		var lines []string
		for _, line := range strings.Split(section, "\n") {
			if strings.TrimSpace(line) != "" {
}

var placeholder = regexp.MustCompile(`\{(?:module_dir|module|report)\}`)

func shellQuote(value string) string {
	return "'" + strings.ReplaceAll(value, "'", "'\"'\"'") + "'"
}

func hookArgs(command any, values map[string]string) ([]string, error) {
	replace := func(s string, quote bool) string {
		return placeholder.ReplaceAllStringFunc(s, func(token string) string {
			value := values[strings.Trim(token, "{}")]
			if quote {
				return shellQuote(value)
			}
			return value
		})
	}
	switch v := command.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return nil, fmt.Errorf("empty hook command")
		}
		return []string{"sh", "-c", replace(v, true)}, nil
	case []any:
		args := make([]string, 0, len(v))
		for _, arg := range v {
			s, ok := arg.(string)
			if !ok {
				return nil, fmt.Errorf("hook argv contains non-string")
			}
			args = append(args, replace(s, false))
		}
		if len(args) == 0 || args[0] == "" {
			return nil, fmt.Errorf("empty hook argv")
		}
		return args, nil
	case []string:
		args := make([]string, len(v))
		for i, s := range v {
			args[i] = replace(s, false)
		}
		if len(args) == 0 || args[0] == "" {
			return nil, fmt.Errorf("empty hook argv")
		}
		return args, nil
	default:
		return nil, fmt.Errorf("invalid hook command %T", command)
	}
}

func runHook(ctx context.Context, root string, command any, values map[string]string, bin string, timeout time.Duration) (gitstore.Result, error) {
	args, err := hookArgs(command, values)
	if err != nil {
		return gitstore.Result{}, err
	}
	result, err := gitstore.Run(ctx, gitstore.Request{Args: args, Dir: root, Env: map[string]string{"FIREWEED_BIN": bin}, Timeout: timeout})
	if err != nil {
		var exit *exec.ExitError
		if !errors.As(err, &exit) {
			return result, fmt.Errorf("run hook: %w", err)
		}
	}
	return result, nil
}

func moduleSpecs(module project.Module) ([]contracts.Item, error) {
	section, ok := contracts.Section(module.Document.Body, "Specs")
	if !ok {
		return nil, nil
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		code, _, ok := contracts.ErrorCode(err)
		if !ok {
			return nil, fmt.Errorf("parse module specs: %w", err)
		}
		return nil, fmt.Errorf("%s: %w", code, err)
	}
	return items, nil
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
59:func sectionBounds(markdown, name string) (int, int, bool) {
79:func Section(markdown, name string) (string, bool) {
87:func ReplaceSection(markdown, name, body string) string {
package evals_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/runner"
)

func assertLogFile(t *testing.T, path, root, run, phase string, attempt int, streams ...string) []byte {
	t.Helper()
	want := filepath.Join(root, ".fireweed", "runs", run, fmt.Sprintf("%s-%d.log", phase, attempt))
	if path != want || !filepath.IsAbs(path) {
		t.Errorf("log path = %q, want absolute %q", path, want)
	}
	if path == "" {
		return nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("read agent log %q: %v", path, err)
		return nil
	}
	for _, stream := range streams {
		if !strings.Contains(string(data), stream) {
			t.Errorf("log %q lacks complete stream %q: %q", path, stream, data)
		}
	}
	return data
}

func Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend(t *testing.T) {
	snap := agentFixture(t)
	logRoot := t.TempDir()
	request := request(snap, "worker", "")
	request.LogRoot = logRoot
	request.Run = "fixture-run"
	request.Phase = "evals"
	request.Attempt = 2
	var prior []byte
	for _, tc := range []struct {
		name, command, stdout, stderr string
		cancel                        bool
		wantError                     bool
	}{
		{"failure", `printf 'failure stdout Ω'; printf 'failure stderr β' >&2; exit 7`, "failure stdout Ω", "failure stderr β", false, true},
		{"timeout", `printf 'timeout stdout Ω'; printf 'timeout stderr β' >&2; sleep 5`, "timeout stdout Ω", "timeout stderr β", false, true},
		{"cancellation", `printf 'cancel stdout Ω'; printf 'cancel stderr β' >&2; sleep 5`, "cancel stdout Ω", "cancel stderr β", true, true},
		{"append", `printf 'second stdout Ω'; printf 'second stderr β' >&2`, "second stdout Ω", "second stderr β", false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			command := tc.command
			request.Timeout = 150 * time.Millisecond
			ctx := context.Background()
			if tc.cancel {
				ready := filepath.Join(logRoot, "cancel.ready")
				command = `printf 'cancel stdout Ω'; printf 'cancel stderr β' >&2; printf ready > ` + shellQuote(ready) + `; sleep 5`
				request.Timeout = 3 * time.Second
				var cancel context.CancelFunc
				ctx, cancel = context.WithCancel(ctx)
				go func() {
					for i := 0; i < 200; i++ {
						if _, err := os.Stat(ready); err == nil {
							cancel()
							return
						}
						time.Sleep(10 * time.Millisecond)
					}
					cancel()
				}()
				defer cancel()
			}
			snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", command}}
			request.Snapshot = snap
			result, err := runner.InvokeAgent(ctx, request)
			if (err != nil) != tc.wantError || result.Stdout != tc.stdout || result.Stderr != tc.stderr {
				t.Errorf("invocation = %#v, %v", result, err)
			}
			if tc.cancel && !errors.Is(err, context.Canceled) {
				t.Errorf("cancellation returned %v", err)
			}
			path := result.LogPath
			data := assertLogFile(t, path, logRoot, "fixture-run", "evals", 2, tc.stdout, tc.stderr)
			if len(prior) > 0 && !strings.HasPrefix(string(data), string(prior)) {
				t.Error("repeated invocation did not append after previous log records")
			}
			prior = data
			if tc.name == "append" {
				for _, earlier := range []string{"failure stdout Ω", "failure stderr β", "timeout stdout Ω", "timeout stderr β", "cancel stdout Ω", "cancel stderr β"} {
					if !strings.Contains(string(data), earlier) {
						t.Errorf("append discarded earlier output %q", earlier)
					}
				}
			}
		})
	}
}

func Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput(t *testing.T) {
	f := newNextFixture(t, true, true, "retry-both")
	result, err := f.run(false)
	if err != nil || result.State.Phase != "applied" || len(result.Steps) != 5 {
		t.Fatalf("Next fixture = %#v, %v", result, err)
	}
5:Reading:      list, status, show, modules, context, explain, playbook
7:Lifecycle:    start, verify, done, next, run, release, commits, abandon

codex
The retained APIs provide the lifecycle and process controls the runner needs. I’ll reuse them and keep the new code under `packages/runner/internal`.
exec
/bin/zsh -lc "mkdir -p packages/runner/internal/{support,invoke,next,run,drill}
cat > packages/runner/internal/support/config.go <<'EOF'
package support

import (
	\"context\"
	\"crypto/rand\"
	\"encoding/hex\"
	\"errors\"
	\"fmt\"
	\"regexp\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
)

func Mapping(value any) map[string]any {
	m, _ := value.(map[string]any)
	return m
}

func String(value any) string {
	s, _ := value.(string)
	return s
}

func List(value any) []string {
	switch v := value.(type) {
	case []string:
		return v
	case []any:
		var strings []string
		for _, item := range v {
			if s, ok := item.(string); ok {
				strings = append(strings, s)
			}
		}
		return strings
	}
	return nil
}

func Attempts(config map[string]any) int {
	if n, ok := config[\"attempts\"].(int); ok && n > 0 {
		return n
	}
	if n, ok := config[\"attempts\"].(float64); ok && n > 0 {
		return int(n)
	}
	return 2
}

func Timeout(config map[string]any, requested time.Duration) (time.Duration, error) {
	if requested > 0 {
		return requested, nil
	}
	return duration(config, \"timeout\", 90*time.Minute, false)
}

func duration(config map[string]any, key string, fallback time.Duration, allowZero bool) (time.Duration, error) {
	value, exists := config[key]
	if "'!exists {
		return fallback, nil
	}
	n, err := time.ParseDuration(String(value))
	if err != nil || n < 0 || (!allowZero && n == 0) {
		message := fmt.Errorf("C01: invalid %s duration %v", key, value)
		return 0, contracts.CodedError{Code: "C01", Subject: key, Cause: message}
	}
	return n, nil
}

type UsagePolicy struct {
	Wait, Budget time.Duration
}

func Usage(config map[string]any) (UsagePolicy, error) {
	wait, err := duration(config, "limit_wait", 30*time.Minute, false)
	if err != nil {
		return UsagePolicy{}, err
	}
	budget, err := duration(config, "limit_budget", 4*time.Hour, true)
	return UsagePolicy{Wait: wait, Budget: budget}, err
}

type usageStop struct{ completed, wait, budget time.Duration }

func (e *usageStop) Error() string {
	return fmt.Sprintf("usage limit stop: completed wait %s, next wait %s, budget %s", e.completed, e.wait, e.budget)
}

type TimeStop struct{ Limit time.Duration }

func (e *TimeStop) Error() string { return fmt.Sprintf("run time limit %s reached", e.Limit) }

type CombinedFailure struct{ Causes []error }

func (e *CombinedFailure) Error() string { return errors.Join(e.Causes...).Error() }
func (e *CombinedFailure) Unwrap() []error { return e.Causes }

func IsUsageStop(err error) bool {
	var stop *usageStop
	var combined *CombinedFailure
	return errors.As(err, &stop) && !errors.As(err, &combined)
}

func IsTimeStop(err error) bool {
	var stop *TimeStop
	var combined *CombinedFailure
	return errors.As(err, &stop) && !errors.As(err, &combined)
}

type Budget struct {
	Waited   time.Duration
	Deadline time.Time
	Stop     *TimeStop
}

func (b *Budget) Wait(ctx context.Context, policy UsagePolicy) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if policy.Wait > policy.Budget-b.Waited {
		return &usageStop{b.Waited, policy.Wait, policy.Budget}
	}
	if b.Stop != nil && policy.Wait > time.Until(b.Deadline) {
		return b.Stop
	}
	timer := time.NewTimer(policy.Wait)
	defer timer.Stop()
	select {
	case <-timer.C:
		b.Waited += policy.Wait
		return ctx.Err()
	case <-ctx.Done():
		return ctx.Err()
	}
}

func RunID() (string, error) {
	var id [16]byte
	_, err := rand.Read(id[:])
	return hex.EncodeToString(id[:]), err
}

func AgentCommand(config map[string]any, role string) any {
	if role == "audit" {
		if command := Mapping(config["audit"])["command"]; command != nil {
			return command
		}
	}
	return Mapping(config["agent"])["command"]
}

func AgentName(command any) string {
	if fields := strings.Fields(String(command)); len(fields) > 0 {
		return fields[0]
	}
	if args := List(command); len(args) > 0 {
		return args[0]
	}
	return ""
}

var placeholder = regexp.MustCompile(`'"\\{[a-z_]+\\}"'`)

func ShellQuote(value string) string {
	return "'"'\" + strings.ReplaceAll(value, \"'\", \"'\\\"'\\\"'\") + \"'\"
}

func Expand(s string, values map[string]string, shell bool) string {
	return placeholder.ReplaceAllStringFunc(s, func(token string) string {
		value, ok := values[token[1:len(token)-1]]
		if "'!ok {
			return token
		}
		if shell {
			return ShellQuote(value)
		}
		return value
	})
}

func Command(value any, values map[string]string) ([]string, error) {
	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) != "" {
			return []string{"sh", "-c", Expand(v, values, true)}, nil
		}
	case []any, []string:
		args, err := Argv(value, values)
		if err != nil {
			return nil, err
		}
		if len(args) > 0 && args[0] != "" {
			return args, nil
		}
	}
	return nil, fmt.Errorf("invalid or empty command (%T)", value)
}

func Argv(value any, values map[string]string) ([]string, error) {
	var args []string
	switch v := value.(type) {
	case nil:
	case []string:
		args = v
	case []any:
		for _, item := range v {
			arg, ok := item.(string)
			if !ok {
				return nil, fmt.Errorf("command argv contains non-string %T", item)
			}
			args = append(args, arg)
		}
	default:
		return nil, fmt.Errorf("expected command argv, got %T", value)
	}
	expanded := make([]string, len(args))
	for i, arg := range args {
		expanded[i] = Expand(arg, values, false)
	}
	return expanded, nil
}

func Git(ctx context.Context, root string, args ...string) (string, error) {
	result, err := gitstore.Run(ctx, gitstore.Request{
		Args: append([]string{"git", "-c", "core.hooksPath=/dev/null"}, args...),
		Dir: root, Timeout: 30*time.Second,
		Env: map[string]string{"GIT_OPTIONAL_LOCKS": "0", "GIT_TERMINAL_PROMPT": "0", "GIT_NO_LAZY_FETCH": "1"},
	})
	if err != nil {
		return result.Stdout, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}

func Init(ctx context.Context, root string, commit bool) error {
	if _, err := Git(ctx, root, "init", "--quiet", "--template="); err != nil {
		return err
	}
	if !commit {
		return nil
	}
	if _, err := Git(ctx, root, "add", "--force", "--all"); err != nil {
		return err
	}
	_, err := Git(ctx, root, "-c", "user.name=Fireweed", "-c", "user.email=fireweed@localhost", "commit", "--quiet", "--allow-empty", "-m", "Fireweed reconstruction inputs")
	return err
}
EOF
cat > packages/runner/internal/invoke/agent.go <<'"'EOF'
package invoke

import (
	\"context\"
	\"errors\"
	\"fmt\"
	\"os\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
	\"github.com/lurkshark/fireweed/packages/runner/model\"
)

func Agent(ctx context.Context, request model.AgentRequest) (result model.AgentResult, err error) {
	logPath, err := invocationLog(request)
	if err "'!= nil {
		return result, err
	}
	result.LogPath = logPath
	defer func() { err = errors.Join(err, appendLog(result)) }()

	config := request.Snapshot.Config
	command := support.AgentCommand(config, request.Role)
	wrapper := support.Mapping(config["agent"])["wrap"]
	timeout, err := support.Timeout(config, request.Timeout)
	if err != nil {
		return result, err
	}
	values := map[string]string{"prompt": request.Prompt}
	if containsTemplate(command, "{git_common_dir}") || containsTemplate(wrapper, "{git_common_dir}") {
		repo, inspectErr := gitstore.Inspect(request.Snapshot.Root)
		if inspectErr != nil {
			return result, inspectErr
		}
		values["git_common_dir"] = repo.CommonDir
	}
	if containsTemplate(command, "{prompt_file}") || containsTemplate(wrapper, "{prompt_file}") {
		f, createErr := os.CreateTemp("", "fireweed-prompt-*")
		if createErr != nil {
			return result, createErr
		}
		defer func() { err = errors.Join(err, os.Remove(f.Name())) }()
		_, writeErr := f.WriteString(request.Prompt)
		if fileErr := errors.Join(writeErr, f.Close()); fileErr != nil {
			return result, fileErr
		}
		values["prompt_file"], err = filepath.Abs(f.Name())
		if err != nil {
			return result, err
		}
	}
	args, err := support.Command(command, values)
	if err != nil {
		return result, err
	}
	prefix, err := support.Argv(wrapper, values)
	if err != nil {
		return result, err
	}
	process, err := gitstore.Run(ctx, gitstore.Request{
		Args: append(prefix, args...), Dir: request.Snapshot.Root, Timeout: timeout,
		Env: map[string]string{"FIREWEED_BIN": request.Bin},
	})
	result.Stdout, result.Stderr, result.ExitCode = process.Stdout, process.Stderr, process.ExitCode
	result.UsageLimit = process.ExitCode != 0 && usageMatch(process, support.List(config["limit_patterns"]))
	if err != nil {
		err = fmt.Errorf("invoke %s: %w", request.Role, err)
	}
	return result, err
}

func containsTemplate(command any, token string) bool {
	if strings.Contains(support.String(command), token) {
		return true
	}
	for _, arg := range support.List(command) {
		if strings.Contains(arg, token) {
			return true
		}
	}
	return false
}

func usageMatch(process gitstore.Result, patterns []string) bool {
	out := strings.TrimRight(process.Stdout, "'"\\n\")
	stderr := strings.TrimRight(process.Stderr, \"\\n\")
	if stderr "'!= "" {
		if out != "" {
			out += "'"\\n\"
		}
		out += stderr
	}
	lines := strings.Split(out, \"\\n\")
	if len(lines) > 20 {
		lines = lines[len(lines)-20:]
	}
	tail := strings.ToLower(strings.Join(lines, \"\\n\"))
	for _, pattern := range patterns {
		if strings.Contains(tail, strings.ToLower(pattern)) {
			return true
		}
	}
	return false
}

func invocationLog(request model.AgentRequest) (string, error) {
	root, run, phase, attempt := request.LogRoot, request.Run, request.Phase, request.Attempt
	if root == \"\" {
		root = request.Snapshot.Root
	}
	if run == \"\" {
		var err error
		run, err = support.RunID()
		if err "'!= nil {
			return "", err
		}
	}
	if phase == "" {
		phase = request.Role
	}
	if attempt <= 0 {
		attempt = 1
	}
	if !filepath.IsLocal(run) || filepath.Base(run) != run || (phase != "" && (!filepath.IsLocal(phase) || filepath.Base(phase) != phase)) {
		return "", fmt.Errorf("invalid log run or phase %q / %q", run, phase)
	}
	return filepath.Abs(filepath.Join(root, ".fireweed", "runs", run, fmt.Sprintf("%s-%d.log", phase, attempt)))
}

func appendLog(result model.AgentResult) error {
	if err := os.MkdirAll(filepath.Dir(result.LogPath), 0755); err != nil {
		return fmt.Errorf("create agent log directory: %w", err)
	}
	f, err := os.OpenFile(result.LogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		return fmt.Errorf("open agent log: %w", err)
	}
	record := fmt.Sprintf("stdout (%d bytes)'"\\n%s\\nstderr (%d bytes)\\n%s\\n\", len(result.Stdout), result.Stdout, len(result.Stderr), result.Stderr)
	_, err = f.WriteString(record)
	if err = errors.Join(err, f.Close()); err "'!= nil {
		return fmt.Errorf("write agent log: %w", err)
	}
	return nil
}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 539ms:
exec
/bin/zsh -lc "sed -n '59,105p' packages/runner/evals/audit_workspace_test.go; sed -n '48,104p' packages/contracts/internal/primitives/primitives.go; sed -n '180,220p' packages/checks/internal/engine/junit.go; sed -n '100,195p' packages/workflow/internal/lifecycle/block_stub.go; sed -n '74,110p;225,252p' packages/runner/evals/invoke_test.go; sed -n '185,215p' packages/runner/evals/next_test.go; sed -n '360,432p' packages/runner/evals/next_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 461ms:
func Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	heldOID := nextGit(t, f.root, "rev-parse", "HEAD")
	nextGit(t, f.root, "update-ref", "refs/fireweed/attempts/sentinel", heldOID)
	auditBeforeEdit(t, f, `find . -path './.git' -prune -o -type f -print > "$1/audit.files"
git rev-parse --is-inside-work-tree > "$1/audit.is-git"
git for-each-ref --format='%(refname)' > "$1/audit.refs"
if git cat-file -e '`+heldOID+`^{commit}' 2>/dev/null; then printf present > "$1/audit.held-history"; else printf absent > "$1/audit.held-history"; fi
cp changes/0001-next.md "$1/audit.change"
cp packages/alpha/SPEC.md "$1/audit.spec"
cp packages/alpha/evals/next_test.go "$1/audit.eval"
`)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" {
		t.Fatalf("isolated audit lifecycle = %#v, %v", r, err)
	}
	cwd := nextRead(t, f.logDir, "audit.cwd")
	if cwd == f.root || !filepath.IsAbs(cwd) || nextRead(t, f.logDir, "audit.is-git") != "true\n" {
		t.Errorf("audit root = %q; git status = %q", cwd, nextRead(t, f.logDir, "audit.is-git"))
	}
	files := strings.Fields(nextRead(t, f.logDir, "audit.files"))
	sort.Strings(files)
	want := []string{"./changes/0001-next.md", "./packages/alpha/SPEC.md", "./packages/alpha/evals/existing_test.go", "./packages/alpha/evals/next_test.go"}
	if strings.Join(files, "\n") != strings.Join(want, "\n") {
		t.Errorf("audit workspace files = %q, want %q", files, want)
	}
	if got := nextRead(t, f.logDir, "audit.held-history"); got != "absent" {
		t.Errorf("audit Git can read held commit: %q", got)
	}
	if refs := nextRead(t, f.logDir, "audit.refs"); strings.Contains(refs, "sentinel") || strings.Contains(refs, "fireweed/") {
		t.Errorf("audit Git exposes held refs: %q", refs)
	}
	for _, item := range []struct{ file, marker string }{{"audit.change", "phase: audit"}, {"audit.change", "### ALP-002"}, {"audit.spec", "### ALP-002"}, {"audit.eval", "EVAL_CONTENT_MARKER"}} {
		if !strings.Contains(nextRead(t, f.logDir, item.file), item.marker) {
			t.Errorf("%s omitted %q", item.file, item.marker)
		}
	}
	if strings.Contains(nextRead(t, f.logDir, "audit.change"), "No findings") {
		t.Error("audit input already contained auditor result")
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	config := nextRead(t, f.root, "fireweed.yaml") + "limit_wait: 1ms\nlimit_budget: 50ms\nlimit_patterns: [RATE LIMIT]\n"
	nextWrite(t, f.root, "fireweed.yaml", config)

// lineAt returns a line without its line ending and the offset of the next line.
func lineAt(source string, start int) (string, int) {
	end := strings.IndexByte(source[start:], '\n')
	if end < 0 {
		return strings.TrimSuffix(source[start:], "\r"), len(source)
	}
	end += start
	return strings.TrimSuffix(source[start:end], "\r"), end + 1
}

func sectionBounds(markdown, name string) (int, int, bool) {
	start := -1
	for pos := 0; pos < len(markdown); {
		line, next := lineAt(markdown, pos)
		if strings.HasPrefix(line, "## ") {
			if start >= 0 {
				return start, pos, true
			}
			if line[3:] == name {
				start = next
			}
		}
		pos = next
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

var itemHeading = regexp.MustCompile("^### ([A-Z0-9]+-[0-9]+|new:[a-z0-9-]+)(?: \\[([^\\[\\]]+)\\])?$")

func ParseItems(section string) ([]model.Item, error) {
	var items []model.Item
	seen := make(map[string]bool)
	bodyStart := 0
	for pos := 0; pos < len(section); {
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
}

func Block(request model.BlockRequest) (model.BlockResult, error) {
	var result model.BlockResult
	s, c, err := loadCurrent(request.Root, request.Change)
	result.State = s
	if err != nil {
		return result, err
	}
	if err := validate(s, c); err != nil {
		return result, err
	}
	if s.Head != c.PhaseBase {
		return result, errors.New("P05: HEAD differs from the held phase base")
	}
	if strings.TrimSpace(request.Reason) == "" {
		return result, errors.New("blocked reason is empty")
	}
	switch c.Phase {
	case "evals", "audit", "impl":
	default:
		return result, fmt.Errorf("P01: unsupported phase %s", c.Phase)
	}
	root := s.Snapshot.Root
	doc, err := blockDocument(root, s, c)
	if err != nil {
		return result, err
	}
	// The index already contains staged work. Select only paths that still
	// need staging so an ignored replacement for a staged deletion stays ignored.
	unstaged, err := commitsGit(root, "ls-files", "--modified", "--others", "--exclude-standard", "-z")
	if err != nil {
		return result, err
	}
	var ordinary strings.Builder
	for _, path := range strings.Split(unstaged, "\x00") {
		if path != "" && !isRunArtifact(path) {
			ordinary.WriteString(path)
			ordinary.WriteByte(0)
		}
	}
	if ordinary.Len() != 0 {
		if _, err := blockGit(root, []byte(ordinary.String()), "--literal-pathspecs", "add", "-A", "--pathspec-from-file=-", "--pathspec-file-nul"); err != nil {
			return result, err
		}
	}
	if err := unstageRunArtifacts(root, c.PhaseBase); err != nil {
		return result, err
	}
	tree, err := blockGit(root, nil, "write-tree")
	if err != nil {
		return result, err
	}
	attemptOID, err := blockGit(root, []byte("attempt: "+s.Change.Slug+"\n"), "commit-tree", tree, "-p", c.PhaseBase)
	if err != nil {
		return result, err
	}
	attemptRef, err := nextAttempt(root, s.Change.Slug, attemptOID)
	if err != nil {
		return result, err
	}
	result.AttemptRef, result.AttemptOID = attemptRef, attemptOID
	if _, err := blockGit(root, nil, "restore", "--source="+c.PhaseBase, "--staged", "--worktree", "--", ".", ":(exclude,literal)"+runArtifactsDir); err != nil {
		return result, err
	}
	doc.Frontmatter["blocked"] = request.Reason
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		return result, err
	}
	if err = write(root, s.Change.Path, []byte(rendered)); err != nil {
		return result, err
	}
	head, err := gitstore.Commit(root, "chore("+s.Change.Module+"): block "+s.Change.Slug, []string{s.Change.Path})
	if err != nil {
		return result, err
	}
	if err = gitstore.CASRef(root, ref(s.Change.Slug), s.ClaimOID, ""); err != nil {
		return result, err
	}
	c.PhaseBase = head
	s.Change.Document = doc
	result.State = state(root, s.Change, c, "", head)
	return result, nil
}

func Test_RUN_003_ShellAndArgvExpandOnceAndPreserveValues(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(common.CommonDir, " ") {
		t.Fatalf("fixture common Git directory must contain spaces: %q", common.CommonDir)
	}
	printArgs := agentScript(t, snap.Root, "args.sh", `printf '%s\n' "$1" "$2"`)
	prompt := `a'b ; $(printf hacked) {git_common_dir}`
	for _, tc := range []struct {
		name    string
		command any
	}{
		{"shell", shellQuote(printArgs) + " {prompt} {git_common_dir}"},
		{"argv", []any{printArgs, "{prompt}", "{git_common_dir}"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap.Config["agent"] = map[string]any{"command": tc.command}
			result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", prompt))
			want := prompt + "\n" + common.CommonDir + "\n"
			if err != nil || result.Stdout != want || result.ExitCode != 0 {
				t.Errorf("%s substitution = %#v, %v; want stdout %q", tc.name, result, err, want)
			}
		})
	}
}

func Test_RUN_003_RUN_004_WrapperPrefixForBothRoles(t *testing.T) {
	snap := agentFixture(t)
	common, err := gitstore.Inspect(snap.Root)
	if err != nil {
		t.Fatal(err)
	}
	log := filepath.Join(snap.Root, "wrap.log")
}

func Test_RUN_008_UsageLimitOnlyForNonzeroMatchInLastTwentyLines(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["limit_patterns"] = []any{"RATE LIMIT"}
	for _, tc := range []struct {
		name, command string
		want          bool
	}{
		{"late_stderr_match", `printf normal; printf 'rAtE LiMiT\n' >&2; exit 9`, true},
		{"late_stdout_match", `printf 'rAtE LiMiT\n'; printf normal >&2; exit 9`, true},
		{"successful_match", `printf 'rate limit\n'; exit 0`, false},
		{"outside_last_twenty", `printf 'rate limit\n'; i=0; while [ "$i" -lt 21 ]; do printf 'noise\n'; i=$((i+1)); done; exit 9`, false},
		{"stdout_before_stderr_window", `printf 'rate limit\n'; i=0; while [ "$i" -lt 21 ]; do printf 'noise\n' >&2; i=$((i+1)); done; exit 9`, false},
		{"literal_substring", `printf 'rating limit\n'; exit 9`, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", tc.command}}
			result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", ""))
			if result.UsageLimit != tc.want || (err != nil) != (result.ExitCode != 0) {
				t.Errorf("%s = %#v, %v; want UsageLimit=%v", tc.name, result, err, tc.want)
			}
		})
	}
}

func Test_RUN_009_LaunchFailureReturnsError(t *testing.T) {
	snap := agentFixture(t)
      printf 'package evals\n// STAGED_EVAL_WORK\n' > packages/alpha/evals/staged_test.go
      git add packages/alpha/evals/staged_test.go
      printf '// TRACKED_EVAL_WORK\n' >> packages/alpha/evals/next_test.go
      printf 'package evals\n// UNTRACKED_EVAL_WORK\n' > packages/alpha/evals/untracked_test.go
      ;;
  esac
fi
printf 'worker-%s' "$phase"
`)
	nextWrite(t, logDir, "auditor.sh", `#!/bin/bash
printf '%s' "$2" > "$1/audit.prompt"
printf 'audit|%s|%s\n' "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
printf '%s' "$PWD" > "$1/audit.cwd"
cat >> changes/0001-next.md <<'AUDIT'

## Audit
No findings
Checked IDs: ALP-002
AUDIT
if [ -f "$1/commit-success" ]; then
  git add changes/0001-next.md
  git -c user.name=Fixture -c user.email=fixture@example.invalid commit -q -m 'fixture agent audit commit'
  printf '%s %s\n' 'audit' "$(git rev-parse HEAD)" >> "$1/fixture-commits"
fi
printf 'auditor'
`)
	if workerMode == "commit-success" {
		nextWrite(t, logDir, "commit-success", "enabled\n")
	}
	nextWrite(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	nextWrite(t, root, "go.mod", "module example.invalid/next-fixture\n\ngo 1.22\n")
		if currentErr != nil || state.Change.Slug != "next" || state.Phase != "evals" {
			t.Errorf("selected change = %#v, %v", state, currentErr)
		}
	})
	t.Run("unapproved_archive_bound", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextGit(t, f.root, "update-ref", "refs/remotes/origin/main", "HEAD")
		nextWrite(t, f.root, "changes/archive/0000-prior.md", "---\nmodule: alpha\nphase: applied\nrun: prior\nbases: {}\n---\n\n## Motivation\nPrior work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		nextWrite(t, f.root, "fireweed.yaml", nextRead(t, f.root, "fireweed.yaml")+"authoring:\n  max_unapproved: 1\n")
		nextCommit(t, f.root, "unapproved prior archive")
		for _, d := range project.Check(f.root) {
			if !d.Warning {
				t.Fatalf("bad bound fixture: %#v", d)
			}
		}
		_, err := f.run(false)
		if err == nil || !strings.Contains(err.Error(), "P07") {
			t.Errorf("unapproved archive bound error = %v", err)
		}
		if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
			t.Errorf("P07 claimed change: %s", refs)
		}
		f.noAgent(t)
	})
	t.Run("shared_base_approval", func(t *testing.T) {
		f := newNextFixture(t, true, false, "exit")
		nextGit(t, f.root, "update-ref", "refs/remotes/origin/main", "HEAD")
		_, err := f.run(false)
		if err == nil {
			t.Error("expected worker failure after shared-base approval")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Phase != "evals" {
			t.Errorf("shared-base approval = %#v, %v", state, currentErr)
		}
	})
	t.Run("allow_unmerged", func(t *testing.T) {
		f := newNextFixture(t, true, false, "exit")
		_, err := f.run(true)
		if err == nil {
			t.Error("expected worker failure after allowed start")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Phase != "evals" || state.Run == "" || !strings.Contains(nextRead(t, f.root, nextChange), "ALP-002") {
			t.Errorf("allowed start = %#v, %v", state, currentErr)
		}
	})
}

func Test_RUN_011_WaitingHeadSelectsIndependent(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	request := runner.NextRequest{Root: f.root, Bin: "/fixture/run-bin", Timeout: 20 * time.Second}
	blocked, blockErr := runner.Next(context.Background(), request)
	if blockErr == nil || !strings.Contains(blockErr.Error(), "P2-01: no fail-first case for ALP-002") || blocked.State.Change.Slug != "alpha" || blocked.State.Phase != "evals" {
		t.Fatalf("fixture alpha did not block: %#v, %v", blocked, blockErr)
	}
	eligible, err := workflow.Eligible(f.root)
	if err != nil || len(eligible) != 1 || eligible[0].Slug != "gamma" {
		t.Fatalf("fixture eligible queue = %#v, %v", eligible, err)
	}
	result, err := runner.Next(context.Background(), request)
	if err != nil || result.Empty || result.State.Change.Slug != "gamma" || result.State.Phase != "applied" {
		t.Errorf("Next past waiting beta = %#v, %v", result, err)
	}
	assertRunBetaWaiting(t, f)
}

func Test_RUN_010_EntirelyWaitingQueue(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	request := runner.NextRequest{Root: f.root, Bin: "/fixture/run-bin", Timeout: 20 * time.Second}
	blocked, blockErr := runner.Next(context.Background(), request)
	if blockErr == nil || !strings.Contains(blockErr.Error(), "P2-01: no fail-first case for ALP-002") || blocked.State.Change.Slug != "alpha" || blocked.State.Phase != "evals" {
		t.Fatalf("fixture alpha did not block: %#v, %v", blocked, blockErr)

exec
/bin/zsh -lc "sed -n '100,194p' packages/runner/evals/audit_workspace_test.go; sed -n '207,250p' packages/runner/evals/next_test.go; rg -n 'blocked|returned|Timeout' packages/runner/evals/drill_test.go | tail -25; sed -n '1,70p' packages/contracts/internal/primitives/primitives.go; rg -n 'func ResolveZone|func .*Zone' packages/contracts/internal/primitives/primitives.go packages/checks/internal/engine/zones.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 541ms:
}

func Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	config := nextRead(t, f.root, "fireweed.yaml") + "limit_wait: 1ms\nlimit_budget: 50ms\nlimit_patterns: [RATE LIMIT]\n"
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextCommit(t, f.root, "approved audit retry policy")
	alternate := filepath.Join(f.logDir, "unapproved-auditor.sh")
	unapproved := strings.Replace(config, filepath.Join(f.logDir, "auditor.sh"), alternate, 1)
	if unapproved == config {
		t.Fatal("fixture did not change the audit command")
	}
	nextWrite(t, f.logDir, "unapproved.yaml", unapproved)
	nextWrite(t, f.logDir, "held-root", f.root)
	nextWrite(t, f.logDir, "unapproved-auditor.sh", "#!/bin/sh\nprintf unapproved > \"$1/unapproved-invoked\"\nexit 9\n")
	auditBeforeEdit(t, f, `count=0
[ ! -f "$1/audit-attempts" ] || count=$(cat "$1/audit-attempts")
count=$((count + 1))
printf '%s' "$count" > "$1/audit-attempts"
if [ "$count" -eq 1 ]; then
  printf '%s' "$PWD" > "$1/first-audit.cwd"
  held=$(cat "$1/held-root")
  cp "$1/unapproved.yaml" "$held/fireweed.yaml"
  printf 'RATE LIMIT' >&2
  exit 9
fi
printf approved > "$1/approved-reused"
exit 9
`)
	r, err := f.run(false)
	if err == nil || runner.IsUsageStop(err) || len(r.Steps) != 3 || !r.Steps[1].Agent.UsageLimit || r.Steps[2].Phase != "audit" {
		t.Fatalf("approved audit retry = %#v, %v", r, err)
	}
	if got := nextRead(t, f.logDir, "approved-reused"); got != "approved" {
		t.Errorf("approved auditor was not reused: %q", got)
	}
	if _, statErr := os.Stat(filepath.Join(f.logDir, "unapproved-invoked")); !os.IsNotExist(statErr) {
		t.Errorf("unapproved auditor ran: %v", statErr)
	}
	if got := nextRead(t, f.root, "fireweed.yaml"); got != unapproved {
		t.Error("fixture did not leave the on-disk unapproved config in place")
	}
	if cwd := nextRead(t, f.logDir, "first-audit.cwd"); cwd == f.root || cwd == nextRead(t, f.logDir, "audit.cwd") {
		t.Errorf("audit retry did not use a fresh restricted workspace: first=%q last=%q", cwd, nextRead(t, f.logDir, "audit.cwd"))
	} else if _, statErr := os.Stat(cwd); !os.IsNotExist(statErr) {
		t.Errorf("first audit workspace remains: %v", statErr)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	nextWrite(t, f.logDir, "held-root", f.root)
	nextWrite(t, f.logDir, "auditor.sh", `#!/bin/bash
printf 'audit|%s|%s\n' "$PWD" "$FIREWEED_BIN" >> "$1/invocations"
printf '%s' "$PWD" > "$1/audit.cwd"
count=0
[ ! -f "$1/audit-attempts" ] || count=$(cat "$1/audit-attempts")
count=$((count + 1))
printf '%s' "$count" > "$1/audit-attempts"
if [ "$count" -eq 1 ]; then
  cat >> changes/0001-next.md <<'AUDIT'

## Audit
No findings
Checked IDs: ALP-002
AUDIT
  held=$(cat "$1/held-root")
  printf 'package internal\n' > "$held/packages/alpha/internal/audit-zone-violation.go"
else
  held=$(cat "$1/held-root")
  cp "$held/changes/0001-next.md" "$1/held-before-retry"
  cp changes/0001-next.md "$1/workspace-before-retry"
  rm "$held/packages/alpha/internal/audit-zone-violation.go"
fi
printf auditor
`)
	r, err := f.run(false)
	if err != nil || r.State.Phase != "applied" || len(r.Steps) != 4 || r.Steps[1].Phase != "audit" || r.Steps[2].Phase != "audit" {
		t.Fatalf("audit verification retry = %#v, %v", r, err)
	}
	for _, path := range []string{"held-before-retry", "workspace-before-retry"} {
		got := nextRead(t, f.logDir, path)
		if !strings.Contains(got, "No findings\nChecked IDs: ALP-002") || !strings.Contains(got, "One new value is needed.") {
			t.Errorf("first permitted Audit edit absent before retry in %s: %q", path, got)
		}
	}
	archived := nextRead(t, f.root, "changes/archive/0001-next.md")
	if !strings.Contains(archived, "No findings; checked ALP-002") {
		t.Errorf("retry did not retain Audit: %q", archived)
	}
	auditWorkspaceGone(t, f)
}

func Test_RUN_046_AuditCommitDoesNotEnterHeldGit(t *testing.T) {
  printf '%s %s\n' 'audit' "$(git rev-parse HEAD)" >> "$1/fixture-commits"
fi
printf 'auditor'
`)
	if workerMode == "commit-success" {
		nextWrite(t, logDir, "commit-success", "enabled\n")
	}
	nextWrite(t, root, ".gitignore", ".fireweed/runs/\n.fireweed/go-test.log\n")
	nextWrite(t, root, "go.mod", "module example.invalid/next-fixture\n\ngo 1.22\n")
	nextWrite(t, root, "SYSTEM.md", "# Fixture\nSYSTEM_CONTEXT_MARKER\n")
	nextWrite(t, root, "fireweed.yaml", fmt.Sprintf(`fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}", %q]
  boundaries: [sh, .fireweed/boundaries.sh]
agent:
  command: [bash, %q, %q, %q, "{prompt}"]
audit:
  command: [bash, %q, %q, "{prompt}"]
timeout: 50ms
`, logDir, worker, logDir, workerMode, auditor, logDir))
	nextWrite(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ -f packages/alpha/evals/staged_test.go ] && [ ! -f "$3/evals-observed" ]; then
  git rev-parse HEAD > "$3/evals-reset-head"
  git diff --cached --name-only > "$3/evals-staged"
  git diff --name-only > "$3/evals-unstaged"
  git ls-files --others --exclude-standard > "$3/evals-untracked"
  touch "$3/evals-observed"
fi
go test ./packages/alpha/evals > .fireweed/go-test.log 2>&1
code=$?
if [ "$code" -eq 0 ]; then result=''; else result='<failure message="Go eval failed"/>'; cat .fireweed/go-test.log; fi
printf '<testsuite><testcase name="ALP-001" classname="fixture"/><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" > "$2"
exit "$code"
`)
	nextWrite(t, root, ".fireweed/boundaries.sh", "#!/bin/sh\nexit 0\n")
	nextWrite(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n## Decisions\n")
	nextWrite(t, root, "packages/alpha/api.go", "package alpha\nimport \"example.invalid/next-fixture/packages/alpha/internal\"\nfunc Value() string { return internal.Value() }\n")
229:	return runner.Drill(ctx, runner.DrillRequest{Root: f.root, Module: "alpha", Bin: "/fixture/new-controller", Timeout: timeout})
308:				t.Errorf("%s preflight returned unrelated error: %v", mode, err)
559:func Test_RUN_056_ConfiguredTimeoutBoundsDrillWorker(t *testing.T) {
603:	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
package primitives

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts/model"
	"gopkg.in/yaml.v3"
)

func ParseDocument(source string) (model.Document, error) {
	first, next := lineAt(source, 0)
	if first != "---" || next == len(source) && !strings.HasSuffix(source, "\n") {
		return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("missing YAML frontmatter")}
	}
	for pos := next; pos < len(source); {
		line, after := lineAt(source, pos)
		if line == "---" {
			var frontmatter map[string]any
			if err := yaml.Unmarshal([]byte(source[next:pos]), &frontmatter); err != nil {
				return model.Document{}, model.CodedError{Code: "C03", Cause: fmt.Errorf("invalid YAML frontmatter: %w", err)}
			}
			if frontmatter == nil {
				return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("YAML frontmatter must be a mapping")}
			}
			return model.Document{Frontmatter: frontmatter, Body: source[after:]}, nil
		}
		pos = after
	}
	return model.Document{}, model.CodedError{Code: "C03", Cause: errors.New("unterminated YAML frontmatter")}
}

func RenderDocument(doc model.Document) (string, error) {
	frontmatter := doc.Frontmatter
	if frontmatter == nil {
		return "", errors.New("render YAML frontmatter: nil mapping")
	}
	data, err := yaml.Marshal(frontmatter)
	if err != nil {
		return "", fmt.Errorf("render YAML frontmatter: %w", err)
	}
	return "---\n" + string(data) + "---\n" + doc.Body, nil
}

// lineAt returns a line without its line ending and the offset of the next line.
func lineAt(source string, start int) (string, int) {
	end := strings.IndexByte(source[start:], '\n')
	if end < 0 {
		return strings.TrimSuffix(source[start:], "\r"), len(source)
	}
	end += start
	return strings.TrimSuffix(source[start:end], "\r"), end + 1
}

func sectionBounds(markdown, name string) (int, int, bool) {
	start := -1
	for pos := 0; pos < len(markdown); {
		line, next := lineAt(markdown, pos)
		if strings.HasPrefix(line, "## ") {
			if start >= 0 {
				return start, pos, true
			}
			if line[3:] == name {
				start = next
			}
		}
packages/contracts/internal/primitives/primitives.go:215:func ResolveZone(path string, zones model.Zones) string {
packages/checks/internal/engine/zones.go:40:func phaseBaseHas(request model.ZoneRequest, file string) bool {
packages/checks/internal/engine/zones.go:48:func zonePath(request model.ZoneRequest, file string) error {
packages/checks/internal/engine/zones.go:114:func Zones(request model.ZoneRequest) []project.Diagnostic {

exec
/bin/zsh -lc "cat > packages/runner/internal/support/files.go <<'EOF'
package support

import (
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/project\"
)

type File struct {
	Mode fs.FileMode
	Data string
}

func ReadFile(path string) (File, error) {
	info, err := os.Lstat(path)
	if err "'!= nil {
		return File{}, err
	}
	mode := info.Mode()
	switch {
	case mode.IsRegular():
		data, err := os.ReadFile(path)
		return File{Mode: mode, Data: string(data)}, err
	case mode.Type() == os.ModeSymlink:
		link, err := os.Readlink(path)
		return File{Mode: mode, Data: link}, err
	default:
		return File{}, fmt.Errorf("unsupported file type at %s", path)
	}
}

func WriteFile(path string, file File) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	if file.Mode.Type() == os.ModeSymlink {
		return os.Symlink(file.Data, path)
	}
	if err := os.WriteFile(path, []byte(file.Data), file.Mode.Perm()); err != nil {
		return err
	}
	return os.Chmod(path, file.Mode.Perm())
}

func Walk(root string, visit func(string, fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == ".git" {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		return visit(rel, entry)
	})
}

func Files(root string) (map[string]File, error) {
	files := map[string]File{}
	err := Walk(root, func(rel string, entry fs.DirEntry) error {
		if entry.IsDir() {
			return nil
		}
		file, err := ReadFile(filepath.Join(root, rel))
		if err == nil {
			files[rel] = file
		}
		return err
	})
	return files, err
}

func Within(path, dir string) bool {
	return path == dir || strings.HasPrefix(path, dir+"/")
}

func Module(snapshot project.Snapshot, name string) (project.Module, bool) {
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module, true
		}
	}
	return project.Module{}, false
}
EOF
cat > packages/runner/internal/next/audit.go <<'"'EOF'
package next

import (
	\"context\"
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/assets\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/invoke\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
	\"github.com/lurkshark/fireweed/packages/runner/model\"
	\"github.com/lurkshark/fireweed/packages/workflow\"
)

func audit(ctx context.Context, state workflow.State, request model.AgentRequest, feedback string) (result model.AgentResult, invoked bool, err error) {
	root, err := os.MkdirTemp(\"\", \"fireweed-audit-*\")
	if err "'!= nil {
		return result, false, err
	}
	defer func() { err = errors.Join(err, os.RemoveAll(root)) }()
	held := request.Snapshot.Root
	change := state.Change.Path
	original, err := support.ReadFile(filepath.Join(held, change))
	if err != nil {
		return result, false, err
	}
	if !original.Mode.IsRegular() {
		return result, false, fmt.Errorf("audit change %s is not a regular file", change)
	}
	module, ok := support.Module(request.Snapshot, state.Change.Module)
	if !ok {
		return result, false, fmt.Errorf("audit module %s is missing", state.Change.Module)
	}
	spec, err := contracts.RenderDocument(module.Document)
	if err != nil {
		return result, false, err
	}
	specPath := filepath.ToSlash(filepath.Join(module.Dir, "SPEC.md"))
	before := map[string]support.File{change: original, specPath: {Mode: 0644, Data: spec}}
	var prompt strings.Builder
	fmt.Fprintf(&prompt, "# Change'"\\n\\n%s\\n\\n%s\\n\\n# Projected spec\\n\\n%s\\n\\n# Evals\\n\", change, original.Data, spec)
	err = support.Walk(filepath.Join(held, module.Dir), func(rel string, entry fs.DirEntry) error {
		if entry.IsDir() || contracts.ResolveZone(rel, module.Zones) "'!= "evals" {
			return nil
		}
		path := filepath.ToSlash(filepath.Join(module.Dir, rel))
		file, err := support.ReadFile(filepath.Join(held, path))
		if err != nil {
			return err
		}
		if !file.Mode.IsRegular() {
			return fmt.Errorf("audit eval %s is not a regular file", path)
		}
		before[path] = file
		fmt.Fprintf(&prompt, "'"\\n## %s\\n%s\\n\", path, file.Data)
		return nil
	})
	if err "'!= nil {
		return result, false, err
	}
	playbook, err := assets.Playbook("audit")
	if err != nil {
		return result, false, err
	}
	fmt.Fprintf(&prompt, "'"\\n# Playbook\\n%s%s\", playbook, feedback)
	for path, file := range before {
		if err := support.WriteFile(filepath.Join(root, path), file); err "'!= nil {
			return result, false, err
		}
	}
	if err := support.Init(ctx, root, false); err != nil {
		return result, false, err
	}

	request.Snapshot.Root = root
	request.Prompt = prompt.String()
	result, agentErr := invoke.Agent(ctx, request)
	submissionErr := importAudit(root, held, change, before)
	if submissionErr != nil {
		return result, true, &support.CombinedFailure{Causes: []error{agentErr, submissionErr}}
	}
	return result, true, agentErr
}

func importAudit(root, held, change string, before map[string]support.File) error {
	after, err := support.Files(root)
	if err != nil {
		return err
	}
	if len(after) != len(before) {
		return fmt.Errorf("audit submission added or deleted files")
	}
	for path, original := range before {
		candidate, exists := after[path]
		if !exists || candidate.Mode != original.Mode || (path != change && candidate.Data != original.Data) {
			return fmt.Errorf("audit submission changed restricted file %s", path)
		}
	}
	original, candidate := before[change].Data, after[change].Data
	if original == candidate {
		return nil
	}
	if !onlyAuditChanged(original, candidate) {
		return fmt.Errorf("audit submission edited %s outside its Audit section", change)
	}
	path := filepath.Join(held, change)
	current, err := support.ReadFile(path)
	if err != nil {
		return err
	}
	if current != before[change] {
		return fmt.Errorf("held change %s changed during audit", change)
	}
	return support.WriteFile(path, after[change])
}

// Locate the raw section so the comparison includes the original frontmatter,
// other section bytes, and the position of an existing Audit section.
func auditBounds(raw string) (start, end int, found, duplicate bool) {
	doc, err := contracts.ParseDocument(raw)
	if err != nil {
		return 0, 0, false, true
	}
	pos := len(raw) - len(doc.Body)
	for line := range strings.SplitAfterSeq(doc.Body, "'"\\n\") {
		if strings.HasPrefix(line, \"## \") {
			if found && end == 0 {
				end = pos
			}
			if strings.TrimRight(line, \"\\r\\n\") == \"## Audit\" {
				if found {
					return 0, 0, false, true
				}
				start, found = pos, true
			}
		}
		pos += len(line)
	}
	if end == 0 {
		end = len(raw)
	}
	return
}

func onlyAuditChanged(before, after string) bool {
	bs, be, bFound, bInvalid := auditBounds(before)
	as, ae, aFound, aInvalid := auditBounds(after)
	switch {
	case bInvalid || aInvalid:
		return false
	case bFound && aFound:
		return before[:bs] == after[:as] && before[be:] == after[ae:]
	case bFound:
		return before[:bs]+before[be:] == after
	case aFound:
		// The blank line separating a newly inserted heading belongs to
		// that insertion when the existing bytes end at the boundary.
		for start := as; ; {
			if after[:start]+after[ae:] == before {
				return true
			}
			if start < 2 || after[start-1] "'!= '"'\\n' || after[start-2] "'!= '"'\\n' {
				return false
			}
			start--
		}
	default:
		return false
	}
}
EOF
cat > packages/runner/internal/next/feedback.go <<'EOF'
package next

import (
	\"fmt\"
	\"maps\"
	\"slices\"
	\"strings\"
	\"unicode/utf8\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/workflow\"
)

func verificationFeedback(result workflow.Result, cause error) string {
	var output strings.Builder
	fmt.Fprintf(&output, \"\\n\\n# Last verification output\\n\\n%s\\n\", cause)
	for _, diagnostic := range result.Diagnostics {
		fmt.Fprintf(&output, \"%s %s %s\\n\", diagnostic.Code, diagnostic.Path, diagnostic.Message)
	}
	for _, module := range slices.Sorted(maps.Keys(result.Reports)) {
		report := result.Reports[module]
		if report.Stdout "'!= "" {
			fmt.Fprintf(&output, "'"\\n%s stdout\\n%s\\n\", module, strings.TrimRight(report.Stdout, \"\\n\"))
		}
		if report.Stderr "'!= "" {
			fmt.Fprintf(&output, "'"\\n%s stderr\\n%s\\n\", module, strings.TrimRight(report.Stderr, \"\\n\"))
		}
	}
	output.WriteString(failedCases(result.Reports))
	return output.String()
}

func failedCases(reports map[string]checks.Report) string {
	const maxSection = 16384
	const maxFailure = 4096
	const omitted = \"\\n[Additional failed cases omitted]\\n\"
	const heading = \"\\n# Failed testcase details\\n\"
	var detail strings.Builder
	remaining := maxSection - len(heading)
	for _, module := range slices.Sorted(maps.Keys(reports)) {
		for _, c := range reports[module].Cases {
			if c.Result "'!= "fail" {
				continue
			}
			failure := c.Failure
			for i := range maxFailure {
				_, size := utf8.DecodeRuneInString(failure)
				failure = failure[size:]
				if failure == "" {
					break
				}
				if i == maxFailure-1 {
					failure = c.Failure[:len(c.Failure)-len(failure)] + "'"\\n[Failure text truncated]\"
					break
				}
			}
			if failure == \"\" {
				failure = c.Failure
			}
			entry := fmt.Sprintf(\"\\nModule: %s\\nClassname: %s\\nTestcase: %s\\n%s\\n\", module, c.Classname, c.Name, failure)
			length := utf8.RuneCountInString(entry)
			if length > remaining-len(omitted) {
				return heading + detail.String() + omitted
			}
			remaining -= length
			detail.WriteString(entry)
		}
	}
	if detail.Len() == 0 {
		return \"\"
	}
	return heading + detail.String()
}
EOF
cat > packages/runner/internal/next/next.go <<'EOF'
package next

import (
	\"context\"
	\"errors\"
	\"fmt\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/invoke\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
	\"github.com/lurkshark/fireweed/packages/runner/model\"
	\"github.com/lurkshark/fireweed/packages/workflow\"
)

type approvalRequired struct{ path string }

func (e *approvalRequired) Error() string {
	return fmt.Sprintf(\"P04: change %s requires approval on the local or shared base\", e.path)
}

func IsApprovalRequired(err error) bool {
	var approval *approvalRequired
	return errors.As(err, &approval)
}

func IsUsageStop(err error) bool { return support.IsUsageStop(err) }

type Blocked struct {
	Ref   string
	Cause error
}

func (e *Blocked) Error() string { return fmt.Sprintf(\"verification exhausted, saved attempt %s: %v\", e.Ref, e.Cause) }
func (e *Blocked) Unwrap() error { return e.Cause }

func Run(ctx context.Context, request model.NextRequest) (model.NextResult, error) {
	return Execute(ctx, request, &support.Budget{})
}

func Execute(ctx context.Context, request model.NextRequest, budget *support.Budget) (result model.NextResult, err error) {
	if err := ctx.Err(); err "'!= nil {
		return result, err
	}
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return result, err
	}
	root := repo.Root
	state, err := workflow.Current(root, "")
	if err != nil {
		queue, err := workflow.Eligible(root)
		if err != nil {
			return result, err
		}
		if len(queue) == 0 {
			result.Empty = true
			return result, nil
		}
		selected := queue[0]
		snapshot, err := project.Load(root)
		if err != nil {
			return result, err
		}
		if _, err := support.Usage(snapshot.Config); err != nil {
			return result, err
		}
		if !request.AllowUnmerged && !approved(root, snapshot, selected.Path) {
			return result, &approvalRequired{selected.Path}
		}
		if err := ctx.Err(); err != nil {
			return result, err
		}
		result.State, err = workflow.Start(root, selected.Slug)
		if err != nil {
			return result, err
		}
		state = result.State
	}
	result.State = state
	usage, err := support.Usage(state.Snapshot.Config)
	if err != nil {
		return result, err
	}
	diagnostics := append([]project.Diagnostic(nil), state.Diagnostics...)
	defer func() { result.State.Diagnostics = diagnostics }()
	for result.State.Phase != "applied" {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		state = result.State
		phase := state.Phase
		role := "worker"
		var prompt string
		switch phase {
		case "audit":
			role = "audit"
		case "evals", "impl":
			prompt, err = workflow.Context(root, state.Change.Slug, true)
			if err != nil {
				return result, err
			}
		default:
			return result, fmt.Errorf("P01: unsupported phase %s", phase)
		}
		attempts := support.Attempts(state.Snapshot.Config)
		check := workflow.Request{
			Root: root, Change: state.Change.Slug, Bin: request.Bin, Timeout: request.Timeout,
			Agent: support.AgentName(support.AgentCommand(state.Snapshot.Config, role)),
		}
		var feedback string
		for attempt := 1; attempt <= attempts; {
			if err := ctx.Err(); err != nil {
				return result, err
			}
			agentReq := model.AgentRequest{
				Snapshot: state.Snapshot, Role: role, Prompt: prompt+feedback,
				Bin: request.Bin, Timeout: request.Timeout,
				LogRoot: root, Run: state.Run, Phase: phase, Attempt: attempt,
			}
			var agent model.AgentResult
			var agentErr, recoveryErr error
			invoked := true
			if role == "audit" {
				agent, invoked, agentErr = audit(ctx, state, agentReq, feedback)
			} else {
				agent, agentErr = invoke.Agent(ctx, agentReq)
				var warning *project.Diagnostic
				warning, recoveryErr = recoverHead(state)
				if warning != nil {
					diagnostics = append(diagnostics, *warning)
				}
			}
			if invoked {
				result.Steps = append(result.Steps, model.StepResult{Phase: phase, Role: role, Agent: agent})
			}
			if recoveryErr != nil {
				return result, errors.Join(agentErr, recoveryErr)
			}
			if fresh, currentErr := workflow.Current(root, state.Change.Slug); currentErr == nil {
				result.State = fresh
			}
			var invalid *support.CombinedFailure
			if errors.As(agentErr, &invalid) {
				return result, agentErr
			}
			if ctx.Err() != nil {
				return result, errors.Join(ctx.Err(), agentErr)
			}
			if agent.UsageLimit && !errors.Is(agentErr, context.DeadlineExceeded) && !errors.Is(agentErr, context.Canceled) {
				if err := budget.Wait(ctx, usage); err != nil {
					return result, err
				}
				continue
			}
			if agentErr != nil {
				return result, agentErr
			}
			verified, verifyErr := workflow.Verify(ctx, check)
			if ctx.Err() != nil {
				return result, ctx.Err()
			}
			if verifyErr == nil {
				done, doneErr := workflow.Done(ctx, check)
				if doneErr != nil {
					if ctx.Err() != nil {
						return result, ctx.Err()
					}
					// Done rechecks the phase and its aftermath. A failure before
					// lifecycle mutation can still be repaired by the worker.
					if done.State.Run != state.Run || done.State.Phase != phase ||
						done.State.Head != state.PhaseBase || done.State.ClaimOID != state.ClaimOID {
						result.State = done.State
						return result, doneErr
					}
					verified, verifyErr = done, doneErr
				} else {
					result.State = done.State
					break
				}
			}
			if attempt == attempts {
				block, blockErr := workflow.Block(workflow.BlockRequest{Root: root, Change: state.Change.Slug, Reason: verifyErr.Error()})
				if block.State.Run != "" {
					result.State = block.State
				}
				if blockErr != nil {
					return result, errors.Join(verifyErr, fmt.Errorf("block phase %s: %w", phase, blockErr))
				}
				return result, &Blocked{Ref: block.AttemptRef, Cause: verifyErr}
			}
			feedback = verificationFeedback(verified, verifyErr)
			attempt++
		}
	}
	return result, nil
}

func approved(root string, snapshot project.Snapshot, path string) bool {
	base := support.String(snapshot.Config["base"])
	if base == "" {
		base = "main"
	}
	remote := support.String(snapshot.Config["remote"])
	if remote == "" {
		remote = "origin"
	}
	for _, ref := range []string{"refs/heads/"+base, "refs/remotes/"+remote+"/"+base} {
		if _, err := gitstore.ReadAt(root, ref, path); err == nil {
			return true
		}
	}
	return false
}

func recoverHead(state workflow.State) (*project.Diagnostic, error) {
	root := state.Snapshot.Root
	repo, err := gitstore.Inspect(root)
	if err != nil {
		return nil, err
	}
	if repo.Head == state.PhaseBase {
		return nil, nil
	}
	if _, err := support.Git(context.Background(), root, "reset", "--soft", state.PhaseBase); err != nil {
		return nil, fmt.Errorf("P05: reset %s HEAD: %w", state.Phase, err)
	}
	return &project.Diagnostic{
		Code: "P05", Warning: true, Path: state.Change.Path,
		Message: strings.Join([]string{state.Phase, "worker HEAD reset to the phase base, retaining index and working-tree work"}, " "),
	}, nil
}
EOF
cat > packages/runner/internal/run/run.go <<'"'EOF'
package run

import (
	\"context\"
	\"errors\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/next\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
	\"github.com/lurkshark/fireweed/packages/runner/model\"
)

func IsTimeStop(err error) bool { return support.IsTimeStop(err) }

func Execute(parent context.Context, request model.RunRequest) (result model.RunResult, err error) {
	ctx := parent
	budget := &support.Budget{}
	if request.For > 0 {
		var cancel context.CancelFunc
		budget.Deadline = time.Now().Add(request.For)
		budget.Stop = &support.TimeStop{Limit: request.For}
		ctx, cancel = context.WithDeadlineCause(parent, budget.Deadline, budget.Stop)
		defer cancel()
	}
	var blocked []error
	defer func() {
		var invalid *support.CombinedFailure
		if "'!errors.As(err, &invalid) && ctx.Err() != nil {
			if parent.Err() != nil {
				err = errors.Join(parent.Err(), err)
			} else if context.Cause(ctx) == budget.Stop {
				err = budget.Stop
			}
		}
		if len(blocked) > 0 {
			err = &support.CombinedFailure{Causes: append(blocked, err)}
		}
	}()

	single := model.NextRequest{
		Root: request.Root, Bin: request.Bin, AllowUnmerged: request.AllowUnmerged, Timeout: request.Timeout,
	}
	for request.Max <= 0 || len(result.Applied) < request.Max {
		change, stepErr := next.Execute(ctx, single, budget)
		if stepErr != nil {
			var saved *next.Blocked
			if errors.As(stepErr, &saved) {
				blocked = append(blocked, stepErr)
				continue
			}
			return result, stepErr
		}
		if change.Empty {
			break
		}
		result.Applied = append(result.Applied, change)
	}
	queue, err := project.Queue(request.Root)
	if err != nil {
		return result, err
	}
	result.Empty = len(queue) == 0
	result.Limited = request.Max > 0 && len(result.Applied) == request.Max && !result.Empty
	return result, nil
}
EOF
gofmt -w packages/runner/internal' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 507ms:
exec
/bin/zsh -lc "sed -n '215,245p' packages/contracts/internal/primitives/primitives.go; sed -n '48,112p' packages/checks/internal/engine/zones.go; cat packages/workflow/internal/lifecycle/regenerate.go; sed -n '477,620p' packages/runner/evals/drill_test.go; sed -n '714,805p' packages/runner/evals/drill_test.go; sed -n '1,90p' .gitignore; cat go.mod" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 471ms:
func ResolveZone(path string, zones model.Zones) string {
	if matchesAny(zones.Protected, path) {
		return "protected"
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
		if matchesAny(zone.patterns, path) {
			return zone.name
		}
	}
	return "other"
}

func matchesAny(patterns []string, path string) bool {
	for _, pattern := range patterns {
		if MatchGlob(pattern, path) {
			return true
		}
	}
	return false
func zonePath(request model.ZoneRequest, file string) error {
	file = path.Clean(strings.TrimPrefix(file, "./"))
	if file == "." || file == ".." || strings.HasPrefix(file, "../") || path.IsAbs(file) {
		return fmt.Errorf("invalid path %q", file)
	}
	protected := []string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}
	protected = append(protected, patterns(request.Snapshot.Config["protected"])...)
	if matches(protected, file) || path.Base(file) == "AGENTS.md" || path.Base(file) == "CLAUDE.md" {
		return fmt.Errorf("protected path %s", file)
	}
	if file == request.ChangePath {
		return nil
	}
	if strings.HasPrefix(file, "changes/") {
		return fmt.Errorf("foreign change %s", file)
	}
	if request.Phase == "audit" {
		return fmt.Errorf("audit may edit only current change: %s", file)
	}
	var owner *project.Module
	for i := range request.Snapshot.Modules {
		module := &request.Snapshot.Modules[i]
		if file == module.Dir || strings.HasPrefix(file, strings.TrimSuffix(module.Dir, "/")+"/") {
			if owner == nil || len(module.Dir) > len(owner.Dir) {
				owner = module
			}
		}
	}
	if owner == nil {
		if matches(patterns(request.Snapshot.Config["worker_writable"]), file) && (request.Phase == "evals" || request.Phase == "impl") {
			return nil
		}
		return fmt.Errorf("outside module: %s", file)
	}
	if owner.Name != request.Module {
		return fmt.Errorf("foreign module: %s", file)
	}
	rel := strings.TrimPrefix(file, strings.TrimSuffix(owner.Dir, "/")+"/")
	zone := contracts.ResolveZone(rel, owner.Zones)
	switch zone {
	case "protected", "spec":
		return fmt.Errorf("protected module path %s", file)
	case "evals":
		if request.Phase == "evals" {
			return nil
		}
	case "interface":
		if request.Phase == "evals" && request.InterfaceDeclared {
			return nil
		}
	case "disposable":
		if request.Phase == "impl" || request.Phase == "evals" && !phaseBaseHas(request, file) {
			return nil
		}
	case "append_only":
		if request.Phase == "impl" && !phaseBaseHas(request, file) {
			return nil
		}
	case "preserve", "other":
		if request.Phase == "impl" || zone == "other" && request.Phase == "evals" {
			return nil
		}
	}
	return fmt.Errorf("%s phase cannot edit %s (%s)", request.Phase, file, zone)
}
package lifecycle

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

// removeDisposable uses the approved projection captured before Start changes
// the worktree. Only files owned by the target module are candidates.
func removeDisposable(root string, snapshot project.Snapshot, name string) error {
	m := module(snapshot, name)
	if m.Dir == "" {
		return fmt.Errorf("regenerate: module %q is missing", name)
	}
	protected := []string{".git", ".git/**", ".fireweed", ".fireweed/**", "SYSTEM.md", "fireweed.yaml"}
	protected = append(protected, list(snapshot.Config["protected"])...)
	var disposable []string
	err := filepath.WalkDir(file(root, m.Dir), func(full string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		rel, err := filepath.Rel(root, full)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		for _, other := range snapshot.Modules {
			if other.Name != name && strings.HasPrefix(other.Dir, m.Dir+"/") && (rel == other.Dir || strings.HasPrefix(rel, other.Dir+"/")) {
				if entry.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
		}
		if path.Base(rel) == ".git" || path.Base(rel) == ".fireweed" {
			if entry.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		if path.Base(rel) == "AGENTS.md" || path.Base(rel) == "CLAUDE.md" {
			return nil
		}
		for _, pattern := range protected {
			if contracts.MatchGlob(pattern, rel) {
				return nil
			}
		}
		moduleRel := strings.TrimPrefix(rel, m.Dir+"/")
		if moduleRel != rel && contracts.ResolveZone(moduleRel, m.Zones) == "disposable" {
			disposable = append(disposable, full)
		}
		return nil
	})
	if err != nil {
		return err
	}
	for _, name := range disposable {
		if err := os.Remove(name); err != nil {
			return err
		}
	}
	return nil
}

func documentDate(value any) time.Time {
	switch date := value.(type) {
	case time.Time:
		return date
	case string:
		for _, layout := range []string{time.RFC3339Nano, "2006-01-02"} {
			if parsed, err := time.Parse(layout, date); err == nil {
				return parsed
			}
		}
	}
	return time.Time{}
}

func latestImplementation(root, name string) time.Time {
	entries, err := os.ReadDir(file(root, "changes/archive"))
	if err != nil {
		return time.Time{}
	}
	var latest time.Time
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") {
			continue
		}
		data, err := read(root, "changes/archive/"+entry.Name())
		if err != nil {
			continue
		}
		doc, err := contracts.ParseDocument(string(data))
		if err != nil || str(doc.Frontmatter["module"]) != name || str(doc.Frontmatter["phase"]) != "applied" {
			continue
		}
		provenance, _ := contracts.Section(doc.Body, "Provenance")
		for _, line := range strings.Split(provenance, "\n") {
			fields := strings.Fields(line)
			if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
				continue
			}
			date := documentDate(fields[2])
			if date.After(latest) {
				latest = date
			}
		}
	}
	return latest
}

func recentPassingDrill(root, name string) bool {
	data, err := read(root, ".fireweed/drills/"+name+".md")
	if err != nil {
		return false
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil || str(doc.Frontmatter["result"]) != "pass" {
		return false
	}
	date := documentDate(doc.Frontmatter["date"])
	return !date.IsZero() && date.After(latestImplementation(root, name))
}
func Test_RUN_055_ReconstructionRejectsModifiedOrDeletedInputs(t *testing.T) {
	for _, mode := range []string{"change_api", "delete_eval", "change_preserved", "change_protected_dependency_interface", "change_setup", "delete_setup"} {
		t.Run(mode, func(t *testing.T) {
			f := drillProjectProtectedFixture(t)
			f.mode(t, mode)
			r, err := f.run(context.Background(), 0)
			if err == nil || r.Result != "fail" || len(r.Steps) != 1 || r.Steps[0].ExitCode != 0 || r.Ref != "" {
				t.Errorf("unauthorized %s reconstruction = %#v, %v", mode, r, err)
			}
			if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "--no-optional-locks", "status", "--porcelain", "-uall") != "" {
				t.Error("invalid submission changed invoking checkout")
			}
			if refs := nextGit(t, f.root, "for-each-ref", "--format=%(refname)", "refs/fireweed/drills"); refs != "" {
				t.Errorf("invalid reconstruction retained rebuilt ref: %q", refs)
			}
			if _, statErr := os.Stat(filepath.Join(f.observations, "eval.order")); !os.IsNotExist(statErr) {
				t.Errorf("invalid reconstruction reached evaluation: %v", statErr)
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_055_ReconstructionImportsOnlyTargetDisposableAdditions(t *testing.T) {
	for _, mode := range []string{"add_foreign", "add_other", "add_project_protected"} {
		t.Run(mode, func(t *testing.T) {
			f := drillProjectProtectedFixture(t)
			f.mode(t, mode)
			r, err := f.run(context.Background(), 0)
			if err != nil || r.Result != "pass" || len(r.Steps) != 1 || r.Ref == "" {
				t.Fatalf("drill with stray file: result=%q, err=%v", r.Result, err)
			}
			if got := nextGit(t, f.root, "diff", "--name-only", f.base, r.Ref); got != "packages/alpha/internal/secret.txt\npackages/alpha/internal/value.go" {
				t.Errorf("rebuilt ref changed files outside the expected disposable paths: %q", got)
			}
			if got := nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go"); got != "new" {
				t.Errorf("rebuilt ref lacks submitted code: %q", got)
			}
			if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "--no-optional-locks", "status", "--porcelain", "-uall") != "" {
				t.Error("accepted submission changed invoking checkout")
			}
			drillCleaned(t, f)
		})
	}
}

func Test_RUN_056_DrillRetriesAgentAndVerificationWithinOneBound(t *testing.T) {
	for _, mode := range []string{"agent_retry", "verify_retry"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
			f.mode(t, mode)
			r, err := f.run(context.Background(), 0)
			if err != nil || r.Result != "pass" || len(r.Steps) != 2 {
				t.Fatalf("retry = %#v, %v", r, err)
			}
			if !strings.Contains(drillObserved(t, f, "prompt-2"), map[string]string{"agent_retry": "WORKER_FAILURE", "verify_retry": "own verification failure"}[mode]) {
				t.Errorf("retry prompt lacks prior failure: %q", drillObserved(t, f, "prompt-2"))
			}
			if got := nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go"); got != "new" {
				t.Errorf("retry did not retain rebuilt work: %q", got)
			}
			drillCleaned(t, f)
		})
	}
	mixed := newDrillFixture(t)
	mixed.mode(t, "mixed_exhaust")
	mixedResult, mixedErr := mixed.run(context.Background(), 0)
	if mixedErr == nil || mixedResult.Result != "fail" || len(mixedResult.Steps) != 2 || !strings.Contains(drillObserved(t, mixed, "prompt-2"), "WORKER_FAILURE") {
		t.Errorf("agent and verification failures did not share attempts: %#v, %v", mixedResult, mixedErr)
	}
	drillCleaned(t, mixed)
	f := newDrillFixture(t)
	f.mode(t, "usage_always")
	nextWrite(t, f.root, "fireweed.yaml", strings.Replace(nextRead(t, f.root, "fireweed.yaml"), "limit_budget: 4ms", "limit_budget: 0s", 1))
	nextCommit(t, f.root, "no usage wait")
	r, err := f.run(context.Background(), 0)
	if err == nil || !runner.IsUsageStop(err) || r.Result != "stopped" || len(r.Steps) != 1 {
		t.Errorf("usage bounded attempts = %#v, %v", r, err)
	}
	drillCleaned(t, f)
}

func Test_RUN_056_ConfiguredTimeoutBoundsDrillWorker(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "timeout")
	config := strings.Replace(nextRead(t, f.root, "fireweed.yaml"), "timeout: 2s", "timeout: 200ms", 1)
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextCommit(t, f.root, "short configured worker timeout")
	start := time.Now()
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result == "pass" || len(r.Steps) != 1 || time.Since(start) > 4*time.Second {
		t.Errorf("configured drill timeout = %#v, %v after %s", r, err, time.Since(start))
	}
	drillCleaned(t, f)
}

func Test_RUN_057_UsageDoesNotSpendAttemptAndStopsWhenBudgetExhausted(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "usage_once")
	config := strings.Replace(nextRead(t, f.root, "fireweed.yaml"), "attempts: 2", "attempts: 1", 1)
	config = strings.Replace(config, "limit_wait: 2ms\nlimit_budget: 4ms", "limit_wait: 800ms\nlimit_budget: 800ms", 1)
	nextWrite(t, f.root, "fireweed.yaml", config)
	nextCommit(t, f.root, "one attempt")
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" || len(r.Steps) != 2 || !r.Steps[0].UsageLimit {
		t.Errorf("usage retry = %#v, %v", r, err)
	}
	first, firstErr := os.Stat(filepath.Join(f.observations, "invoked-1"))
	second, secondErr := os.Stat(filepath.Join(f.observations, "invoked-2"))
	if firstErr != nil || secondErr != nil {
		t.Errorf("usage invocation timestamps: first=%v second=%v", firstErr, secondErr)
	} else if gap := second.ModTime().Sub(first.ModTime()); gap < 700*time.Millisecond {
		t.Errorf("configured 800ms usage wait lasted only %s", gap)
	}
	drillCleaned(t, f)

	stopped := newDrillFixture(t)
	stopped.mode(t, "usage_always")
	r, err = stopped.run(context.Background(), 0)
	if err == nil || !runner.IsUsageStop(err) || r.Result != "stopped" || len(r.Steps) != 3 || !r.Steps[2].UsageLimit {
		t.Errorf("usage budget = %#v, %v", r, err)
	}
	drillCleaned(t, stopped)

	cancelled := newDrillFixture(t)
	cancelled.mode(t, "timeout")
	ctx, cancel := context.WithTimeout(context.Background(), 250*time.Millisecond)
	defer cancel()
	r, err = cancelled.run(ctx, 5*time.Second)
	if !errors.Is(err, context.DeadlineExceeded) || r.Result == "pass" {
		t.Errorf("cancelled drill = %#v, %v", r, err)
	}
	drillCleaned(t, cancelled)

	timed := newDrillFixture(t)
	timed.mode(t, "timeout")
	r, err = timed.run(context.Background(), 250*time.Millisecond)
	if err == nil || r.Result == "pass" || len(r.Steps) != 1 {
		t.Errorf("timed drill = %#v, %v", r, err)
	}
	drillCleaned(t, timed)
}

func Test_RUN_058_RUN_059_DrillReturnsAllReportsEvenWhenEvalsFail(t *testing.T) {
func Test_RUN_060_RUN_061_DrillReportAndRetainedRef(t *testing.T) {
	f := newDrillFixture(t)
	const priorMarker = "PRIOR_REVIEWED_DRILL_REPORT_MARKER"
	const reportPath = ".fireweed/drills/alpha.md"
	nextWrite(t, f.root, reportPath, "---\nresult: fail\n---\n\nPrior review: "+priorMarker+"\n")
	nextGit(t, f.root, "add", "-f", reportPath)
	nextGit(t, f.root, "commit", "-q", "-m", "reviewed prior drill report")
	f.base = nextGit(t, f.root, "rev-parse", "HEAD")
	nextWrite(t, f.observations, "base", f.base)
	if prior := nextGit(t, f.root, "show", f.base+":"+reportPath); !strings.Contains(prior, priorMarker) {
		t.Fatalf("committed prior report lacks marker: %q", prior)
	}
	r, err := f.run(context.Background(), 0)
	if err != nil || r.Result != "pass" || r.Module != "alpha" || r.Base != f.base || r.Date == "" || r.Agent == "" || r.Model == "" || r.ReportPath != ".fireweed/drills/alpha.md" || r.Ref != "refs/fireweed/drills/alpha/"+r.Date || len(r.Steps) != 1 {
		t.Fatalf("passing result = %#v, %v", r, err)
	}
	if nextGit(t, f.root, "rev-parse", "HEAD") != f.base || nextGit(t, f.root, "for-each-ref", "--format=%(refname)", r.Ref) != r.Ref || nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go") != "new" || nextGit(t, f.root, "show", "HEAD:packages/alpha/internal/value.go") != "OLD_IMPLEMENTATION_SECRET" {
		t.Error("retained rebuilt ref or invoking checkout wrong")
	}
	report := nextRead(t, f.root, r.ReportPath)
	if strings.Contains(report, priorMarker) {
		t.Errorf("current report retained prior review marker: %q", report)
	}
	if prior := nextGit(t, f.root, "show", f.base+":"+reportPath); !strings.Contains(prior, priorMarker) {
		t.Errorf("reviewed prior report is not readable in Git: %q", prior)
	}
	for _, want := range []string{"date:", "base_commit: " + f.base, "agent:", "model:", "result: pass", "## Own evals", "## Dependent evals", "## Contradicted Decisions", "## Gaps", "ALP-001_case", "AAR-001_case", "ZET-001_case", "ALP-001", "AAR-001", "ZET-001", "pass", "ordered preference may be contradicted", "module alpha", "no spec", "module aardvark"} {
		if !strings.Contains(report, want) {
			t.Errorf("report lacks %q: %q", want, report)
		}
	}
	for _, entry := range []struct{ section, module, testcase, id, result string }{
		{"Own evals", "alpha", "ALP-001_case", "ALP-001", "pass"},
		{"Dependent evals", "aardvark", "AAR-001_case", "AAR-001", "pass"},
		{"Dependent evals", "zeta", "ZET-001_case", "ZET-001", "pass"},
	} {
		if !drillEvalEntry(drillReportSection(report, entry.section), entry.module, entry.testcase, entry.id, entry.result) {
			t.Errorf("%s lacks a complete %s eval entry: %q", entry.section, entry.module, report)
		}
	}
	if !drillReportLine(drillReportSection(report, "Gaps"), "ALP-001", "alpha", "ordering needs a spec") || !drillReportLine(drillReportSection(report, "Gaps"), "no spec", "aardvark", "caller expectation") {
		t.Errorf("gap entries do not each identify nearest spec and target module: %q", report)
	}
	if !strings.Contains(report, "worker observation") {
		t.Errorf("successful worker observation absent from report: %q", report)
	}
	if len(r.Own.Cases) != 1 || len(r.Dependents) != 2 {
		t.Errorf("passing reports = %#v, %#v", r.Own, r.Dependents)
	}
	if drillObserved(t, f, "agent.bin") != "/fixture/new-controller" || strings.TrimSpace(drillObserved(t, f, "wrapper.bin")) != "/fixture/new-controller" {
		t.Error("worker or wrapper did not receive configured binary")
	}
	drillCleaned(t, f)
}

func Test_RUN_060_RUN_061_FailedDrillReportsFailuresAndRetainsCode(t *testing.T) {
	f := newDrillFixture(t)
	f.mode(t, "wrong")
	r, err := f.run(context.Background(), 0)
	if err == nil || r.Result != "fail" || r.Ref == "" || r.ReportPath != ".fireweed/drills/alpha.md" {
		t.Fatalf("failed drill = %#v, %v", r, err)
	}
	if nextGit(t, f.root, "show", r.Ref+":packages/alpha/internal/value.go") != "wrong" {
		t.Error("failed rebuilt code was not retained")
	}
	report := nextRead(t, f.root, r.ReportPath)
	for _, want := range []string{"result: fail", "own verification failure", "ALP-001_case", "ALP-001", "fail", "## Contradicted Decisions", "## Gaps"} {
		if !strings.Contains(report, want) {
			t.Errorf("failure report lacks %q", want)
		}
	}
	if !drillEvalEntry(drillReportSection(report, "Own evals"), "alpha", "ALP-001_case", "ALP-001", "fail") {
		t.Errorf("failed own eval entry lacks module, testcase, spec ID or result: %q", report)
	}
	drillCleaned(t, f)

	agentFailure := newDrillFixture(t)
	agentFailure.mode(t, "agent_fail_always")
	failed, agentErr := agentFailure.run(context.Background(), 0)
	if agentErr == nil || failed.Result != "fail" || len(failed.Steps) != 2 || failed.ReportPath != ".fireweed/drills/alpha.md" {
		t.Errorf("agent exhaustion report = %#v, %v", failed, agentErr)
	}
	if report := nextRead(t, agentFailure.root, ".fireweed/drills/alpha.md"); !strings.Contains(report, "WORKER_FAILURE") || !strings.Contains(report, "result: fail") {
		t.Errorf("agent failure absent from report: %q", report)
	}
	drillCleaned(t, agentFailure)
}

func Test_RUN_062_CleanupOnSuccessAndFailures(t *testing.T) {
	for _, mode := range []string{"success", "agent_retry", "change_api", "usage_always", "timeout"} {
		t.Run(mode, func(t *testing.T) {
			f := newDrillFixture(t)
.fireweed/runs/
.fireweed/bin/
bin/
.fireweed/evidence/
module github.com/lurkshark/fireweed

go 1.23

require gopkg.in/yaml.v3 v3.0.1

exec
/bin/zsh -lc "cat > packages/runner/internal/drill/workspace.go <<'EOF'
package drill

import (
	\"context\"
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path\"
	\"path/filepath\"
	\"slices\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
)

func dependentModules(snapshot project.Snapshot, target string) []project.Module {
	dependent := map[string]bool{target: true}
	for changed := true; changed; {
		changed = false
		for _, module := range snapshot.Modules {
			if dependent[module.Name] {
				continue
			}
			for _, dependency := range module.DependsOn {
				if dependent[dependency] {
					dependent[module.Name] = true
					changed = true
					break
				}
			}
		}
	}
	var modules []project.Module
	for _, module := range snapshot.Modules {
		if module.Name "'!= target && dependent[module.Name] {
			modules = append(modules, module)
		}
	}
	slices.SortFunc(modules, func(a, b project.Module) int { return strings.Compare(a.Name, b.Name) })
	return modules
}

func disposable(snapshot project.Snapshot, target project.Module, name string) bool {
	rel, inside := strings.CutPrefix(name, target.Dir+"/")
	if !inside || contracts.ResolveZone(rel, target.Zones) != "disposable" {
		return false
	}
	for _, module := range snapshot.Modules {
		if module.Name != target.Name && len(module.Dir) > len(target.Dir) && support.Within(name, module.Dir) {
			return false
		}
	}
	switch path.Base(name) {
	case "AGENTS.md", "CLAUDE.md":
		return false
	}
	for _, part := range strings.Split(rel, "/") {
		if part == ".git" || part == ".fireweed" {
			return false
		}
	}
	for _, pattern := range support.List(snapshot.Config["protected"]) {
		if contracts.MatchGlob(pattern, name) {
			return false
		}
	}
	return true
}

func setup(ctx context.Context, root, bin string, timeout time.Duration, config map[string]any) error {
	command := support.Mapping(config["hooks"])["setup"]
	if command == nil {
		return nil
	}
	args, err := support.Command(command, nil)
	if err != nil {
		return fmt.Errorf("drill setup: %w", err)
	}
	process, err := gitstore.Run(ctx, gitstore.Request{
		Args: args, Dir: root, Timeout: timeout,
		Env: map[string]string{"FIREWEED_BIN": bin},
	})
	if err != nil {
		return fmt.Errorf("drill setup failed (exit %d): %w'"\\n%s\\n%s\", process.ExitCode, err, process.Stdout, process.Stderr)
	}
	return nil
}

func prepareInput(ctx context.Context, snapshot project.Snapshot, input string, tracked []string, target project.Module, dependents []project.Module) error {
	for _, name := range tracked {
		if name == \"\" || support.Within(name, \"changes\") || support.Within(name, \".fireweed/drills\") || disposable(snapshot, target, name) {
			continue
		}
		if slices.ContainsFunc(dependents, func(module project.Module) bool { return support.Within(name, module.Dir) }) {
			continue
		}
		file, err := support.ReadFile(filepath.Join(snapshot.Root, name))
		if err "'!= nil {
			return err
		}
		if err := support.WriteFile(filepath.Join(input, name), file); err != nil {
			return err
		}
	}
	return support.Init(ctx, input, true)
}

func removeDisposable(snapshot project.Snapshot, target project.Module, scratch string) error {
	return support.Walk(filepath.Join(scratch, target.Dir), func(rel string, entry fs.DirEntry) error {
		if entry.IsDir() {
			return nil
		}
		name := target.Dir+"/"+rel
		if disposable(snapshot, target, name) {
			return os.Remove(filepath.Join(scratch, name))
		}
		return nil
	})
}

func candidateFiles(root string, baseline map[string]support.File, snapshot project.Snapshot, target project.Module) (map[string]support.File, error) {
	for name, original := range baseline {
		current, err := support.ReadFile(filepath.Join(root, name))
		if err != nil || current != original {
			return nil, fmt.Errorf("drill reconstruction modified or deleted input %s", name)
		}
	}
	candidate := map[string]support.File{}
	err := support.Walk(root, func(name string, entry fs.DirEntry) error {
		if _, retained := baseline[name]; retained || entry.IsDir() || !disposable(snapshot, target, name) {
			return nil
		}
		file, err := support.ReadFile(filepath.Join(root, name))
		if err == nil {
			candidate[name] = file
		}
		return err
	})
	return candidate, err
}

func importCandidate(scratch string, candidate map[string]support.File) error {
	for name, file := range candidate {
		if err := support.WriteFile(filepath.Join(scratch, name), file); err != nil {
			return err
		}
	}
	return nil
}

func saveCandidate(ctx context.Context, repo gitstore.Repository, scratch string, target project.Module, snapshot project.Snapshot, tracked []string, candidate map[string]support.File) (string, error) {
	// Build the retained tree from the original index plus accepted paths.
	if _, err := support.Git(ctx, scratch, "read-tree", repo.Head); err != nil {
		return "", err
	}
	for _, name := range tracked {
		if !disposable(snapshot, target, name) {
			continue
		}
		if _, rebuilt := candidate[name]; !rebuilt {
			if _, err := support.Git(ctx, scratch, "update-index", "--force-remove", "--", name); err != nil {
				return "", err
			}
		}
	}
	if len(candidate) > 0 {
		args := []string{"--literal-pathspecs", "add", "--force", "--"}
		for name := range candidate {
			args = append(args, name)
		}
		if _, err := support.Git(ctx, scratch, args...); err != nil {
			return "", err
		}
	}
	tree, err := support.Git(ctx, scratch, "write-tree")
	if err != nil {
		return "", err
	}
	commit, err := support.Git(ctx, scratch, "commit-tree", strings.TrimSpace(tree), "-p", repo.Head, "-m", "drill("+target.Name+"): retain reconstruction")
	return strings.TrimSpace(commit), err
}
EOF
cat > packages/runner/internal/drill/report.go <<'"'EOF'
package drill

import (
	\"fmt\"
	\"maps\"
	\"os\"
	\"path/filepath\"
	\"slices\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
	\"github.com/lurkshark/fireweed/packages/runner/model\"
)

func reconstructionPrompt(snapshot project.Snapshot, target project.Module, baseline map[string]support.File) (string, error) {
	spec, err := contracts.RenderDocument(target.Document)
	if err "'!= nil {
		return "", err
	}
	var prompt strings.Builder
	fmt.Fprintf(&prompt, "# Rebuild %s'"\\n\\n## Target spec and Decisions\\n\\n%s\\n\", target.Name, spec)
	for _, zone := range []string{\"interface\", \"evals\"} {
		fmt.Fprintf(&prompt, \"\\n## Retained %s\\n\\n\", zone)
		for _, name := range slices.Sorted(maps.Keys(baseline)) {
			rel, inside := strings.CutPrefix(name, target.Dir+\"/\")
			if inside && contracts.ResolveZone(rel, target.Zones) == zone {
				fmt.Fprintln(&prompt, name)
			}
		}
	}
	prompt.WriteString(\"\\nRebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.\\n\")
	command := support.Mapping(snapshot.Config[\"hooks\"])[\"eval\"]
	if override, ok := support.Mapping(target.Document.Frontmatter[\"hooks\"])[\"eval\"]; ok {
		command = override
	}
	if command "'!= nil {
		report := ".fireweed/runs/drill/report.xml"
		args, err := support.Command(command, map[string]string{"module": target.Name, "module_dir": target.Dir, "report": report})
		if err != nil {
			return "", err
		}
		for i, arg := range args {
			args[i] = support.ShellQuote(arg)
		}
		fmt.Fprintf(&prompt, "'"\\nRun the target's eval hook from the workspace root as often as needed.\\n\\n"'```sh'"\\nmkdir -p .fireweed/runs/drill && %s\\n"'```'"\\n\", strings.Join(args, \" \"))
	}
	prompt.WriteString(\"\\nWrite DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or \\\"no spec\\\" and a target module.\\n\")
	return prompt.String(), nil
}

func agentModel(config map[string]any, command any) string {
	if model := support.String(support.Mapping(config[\"agent\"])[\"model\"]); model "'!= "" {
		return model
	}
	args := support.List(command)
	if shell, ok := command.(string); ok {
		args = strings.Fields(shell)
	}
	for i, arg := range args {
		if (arg == "-m" || arg == "--model") && i+1 < len(args) {
			return strings.Trim(args[i+1], "'"\\\"'\")
		}
		if value, ok := strings.CutPrefix(arg, \"--model=\"); ok {
			return strings.Trim(value, \"\\\"'\")
		}
	}
	return \"unknown\"
}

func reportCases(out *strings.Builder, module string, report checks.Report) {
	if len(report.Cases) == 0 {
		fmt.Fprintf(out, \"- %s: no eval cases reported\\n\", module)
	}
	for _, c := range report.Cases {
		ids := strings.Join(c.IDs, \", \")
		if ids == \"\" {
			ids = \"no spec\"
		}
		fmt.Fprintf(out, \"- %s / %s / %s / %s\\n\", module, c.Name, ids, c.Result)
		if c.Failure "'!= "" {
			fmt.Fprintf(out, "  %s'"\\n\", strings.ReplaceAll(strings.TrimRight(c.Failure, \"\\n\"), \"\\n\", \"\\n  \"))
		}
	}
}

func writeReport(root string, result model.DrillResult, started time.Time, notes string, failure error) error {
	var body strings.Builder
	body.WriteString(\"# Drill \" + result.Module + \"\\n\\n## Own evals\\n\\n\")
	reportCases(&body, result.Module, result.Own)
	body.WriteString(\"\\n## Dependent evals\\n\\n\")
	for _, module := range slices.Sorted(maps.Keys(result.Dependents)) {
		reportCases(&body, module, result.Dependents[module])
	}
	for _, heading := range []string{\"Contradicted Decisions\", \"Gaps\"} {
		section, _ := contracts.Section(notes, heading)
		if strings.TrimSpace(section) == \"\" {
			section = \"None reported.\"
		}
		fmt.Fprintf(&body, \"\\n## %s\\n\\n%s\\n\", heading, strings.TrimSpace(section))
	}
	body.WriteString(\"\\n## Agent observations\\n\")
	for i, step := range result.Steps {
		fmt.Fprintf(&body, \"\\n### Invocation %d (exit %d)\\n\\n%s\\n%s\\n\", i+1, step.ExitCode, step.Stdout, step.Stderr)
	}
	if failure "'!= nil {
		fmt.Fprintf(&body, "'"\\n## Drill errors\\n\\n%s\\n\", failure)
	}
	doc, err := contracts.RenderDocument(contracts.Document{
		Frontmatter: map[string]any{
			\"date\": started.UTC().Format(time.RFC3339Nano), \"base_commit\": result.Base,
			\"agent\": result.Agent, \"model\": result.Model, \"result\": result.Result,
		},
		Body: \"\\n\"+body.String(),
	})
	if err "'!= nil {
		return err
	}
	path := filepath.Join(root, result.ReportPath)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(doc), 0644)
}
EOF
cat > packages/runner/internal/drill/drill.go <<'"'EOF'
package drill

import (
	\"context\"
	\"errors\"
	\"fmt\"
	\"os\"
	\"path/filepath\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/invoke\"
	\"github.com/lurkshark/fireweed/packages/runner/internal/support\"
	\"github.com/lurkshark/fireweed/packages/runner/model\"
)

func Run(ctx context.Context, request model.DrillRequest) (result model.DrillResult, err error) {
	started := time.Now().UTC()
	repo, err := gitstore.Inspect(request.Root)
	if err "'!= nil {
		return result, err
	}
	snapshot, err := project.LoadAt(repo.Root, repo.Head)
	if err != nil {
		return result, err
	}
	target, ok := support.Module(snapshot, request.Module)
	if !ok {
		return result, fmt.Errorf("unknown drill module %q", request.Module)
	}
	dirty, err := gitstore.Changes(repo.Root, repo.Head)
	if err != nil {
		return result, err
	}
	if len(dirty) > 0 {
		return result, fmt.Errorf("drill requires a clean checkout, found dirty path %s", dirty[0].Path)
	}
	claims, err := support.Git(ctx, repo.Root, "for-each-ref", "--format=%(refname)", "refs/fireweed/claims/")
	if err != nil {
		return result, err
	}
	if claims != "" {
		return result, fmt.Errorf("drill requires no shared change claims: %s", strings.TrimSpace(claims))
	}

	config := snapshot.Config
	timeout, err := support.Timeout(config, request.Timeout)
	if err != nil {
		return result, err
	}
	usage, err := support.Usage(config)
	if err != nil {
		return result, err
	}
	run, err := support.RunID()
	if err != nil {
		return result, err
	}
	command := support.AgentCommand(config, "worker")
	result = model.DrillResult{
		Module: target.Name, Base: repo.Head, Date: started.Format("20060102T150405.000000000"),
		Agent: support.AgentName(command), Model: agentModel(config, command),
		Result: "fail", Dependents: map[string]checks.Report{},
	}
	var notes string
	defer func() {
		if len(result.Steps) == 0 {
			return
		}
		result.ReportPath = filepath.ToSlash(filepath.Join(".fireweed", "drills", target.Name+".md"))
		if reportErr := writeReport(repo.Root, result, started, notes, err); reportErr != nil {
			result.Result = "fail"
			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("write drill report: %w", reportErr)}}
		}
	}()

	parent := filepath.Join(repo.CommonDir, "fireweed", "worktrees")
	if err := os.MkdirAll(parent, 0755); err != nil {
		return result, err
	}
	scratch, err := os.MkdirTemp(parent, "drill-"+target.Name+"-*")
	if err != nil {
		return result, err
	}
	defer func() {
		_, cleanupErr := support.Git(context.Background(), repo.Root, "worktree", "remove", "--force", scratch)
		if cleanupErr != nil {
			result.Result = "fail"
			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("remove verification worktree: %w", cleanupErr)}}
		}
	}()
	if _, err := support.Git(ctx, repo.Root, "worktree", "add", "--detach", "--quiet", scratch, repo.Head); err != nil {
		// Worktree creation can leave an empty directory before it registers.
		_ = os.Remove(scratch)
		return result, err
	}
	if err := setup(ctx, scratch, request.Bin, timeout, config); err != nil {
		return result, err
	}
	if err := removeDisposable(snapshot, target, scratch); err != nil {
		return result, err
	}
	trackedOutput, err := support.Git(ctx, repo.Root, "ls-tree", "-r", "--name-only", "-z", repo.Head)
	if err != nil {
		return result, err
	}
	tracked := strings.Split(strings.TrimSuffix(trackedOutput, "'"\\x00\"), \"\\x00\")
	dependents := dependentModules(snapshot, target.Name)
	input, err := os.MkdirTemp(\"\", \"fireweed-drill-input-*\")
	if err "'!= nil {
		return result, err
	}
	defer func() {
		if cleanupErr := os.RemoveAll(input); cleanupErr != nil {
			result.Result = "fail"
			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("remove reconstruction workspace: %w", cleanupErr)}}
		}
	}()
	if err := prepareInput(ctx, snapshot, input, tracked, target, dependents); err != nil {
		return result, err
	}
	if err := setup(ctx, input, request.Bin, timeout, config); err != nil {
		return result, err
	}
	baseline, err := support.Files(input)
	if err != nil {
		return result, err
	}
	prompt, err := reconstructionPrompt(snapshot, target, baseline)
	if err != nil {
		return result, err
	}
	inputSnapshot, verifySnapshot := snapshot, snapshot
	inputSnapshot.Root, verifySnapshot.Root = input, scratch

	var budget support.Budget
	var failure error
	var retained string
	attempts := support.Attempts(config)
	for attempt := 1; attempt <= attempts; {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		text := prompt
		if failure != nil {
			text += "'"\\n## Last reconstruction failure\\n\\n\" + failure.Error() + \"\\n\"
		}
		step, agentErr := invoke.Agent(ctx, model.AgentRequest{
			Snapshot: inputSnapshot, Role: \"worker\", Prompt: text, Bin: request.Bin, Timeout: request.Timeout,
			LogRoot: repo.Root, Run: run, Phase: \"drill\", Attempt: attempt,
		})
		result.Steps = append(result.Steps, step)
		if file, notesErr := support.ReadFile(filepath.Join(input, \"DRILL_NOTES.md\")); notesErr == nil && file.Mode.IsRegular() {
			notes = file.Data
		}
		if errors.Is(agentErr, context.Canceled) || errors.Is(agentErr, context.DeadlineExceeded) {
			return result, agentErr
		}
		if ctx.Err() "'!= nil {
			return result, ctx.Err()
		}
		if step.UsageLimit {
			if err := budget.Wait(ctx, usage); err != nil {
				if support.IsUsageStop(err) {
					result.Result = "stopped"
				}
				return result, err
			}
			continue
		}
		if agentErr != nil {
			failure = fmt.Errorf("%w'"\\n%s\\n%s\", agentErr, step.Stdout, step.Stderr)
			attempt++
			continue
		}
		candidate, err := candidateFiles(input, baseline, snapshot, target)
		if err "'!= nil {
			return result, err
		}
		if err := removeDisposable(snapshot, target, scratch); err != nil {
			return result, err
		}
		if err := importCandidate(scratch, candidate); err != nil {
			return result, err
		}
		retained, err = saveCandidate(ctx, repo, scratch, target, snapshot, tracked, candidate)
		if err != nil {
			return result, err
		}
		failure = verify(ctx, request, verifySnapshot, target, dependents, &result)
		if ctx.Err() != nil {
			return result, ctx.Err()
		}
		if failure == nil {
			break
		}
		attempt++
	}
	if retained != "" {
		ref := "refs/fireweed/drills/"+target.Name+"/"+result.Date
		if err := gitstore.CASRef(repo.Root, ref, "", retained); err != nil {
			return result, errors.Join(failure, err)
		}
		result.Ref = ref
	}
	if failure != nil {
		return result, failure
	}
	result.Result = "pass"
	return result, nil
}

func verify(ctx context.Context, request model.DrillRequest, snapshot project.Snapshot, target project.Module, dependents []project.Module, result *model.DrillResult) error {
	var failures []error
	for _, module := range append([]project.Module{target}, dependents...) {
		report, err := checks.Evaluate(ctx, checks.HookRequest{Snapshot: snapshot, Module: module.Name, Bin: request.Bin, Timeout: request.Timeout})
		if module.Name == target.Name {
			result.Own = report
		} else {
			result.Dependents[module.Name] = report
		}
		specs, _ := contracts.Section(module.Document.Body, "Specs")
		items, parseErr := contracts.ParseItems(specs)
		err = errors.Join(err, parseErr, checks.RequireCoverage(report, items), checks.RequirePass(report, items))
		if err != nil {
			var detail strings.Builder
			for _, c := range report.Cases {
				if c.Result == "fail" && c.Failure != "" {
					fmt.Fprintf(&detail, "'"\\n%s: %s\", c.Name, c.Failure)
				}
			}
			failures = append(failures, fmt.Errorf(\"%s verification: %w%s\\n%s%s\", module.Name, err, detail.String(), report.Stdout, report.Stderr))
		}
		if ctx.Err() "'!= nil {
			break
		}
	}
	return errors.Join(failures...)
}
EOF
gofmt -w packages/runner/internal
export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
mkdir -p .fireweed/runs/drill
sh -c '"'go run .fireweed/hooks.go eval '\"'\"'packages/runner'\"'\"' '\"'\"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/.fireweed/runs/drill/report.xml'\"'\"''" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc ".fireweed/bin/fireweed help; git status --short; go version; rg -n 'func LoadAt|func (Check|Queue)' packages/project/internal/{fixture,queue}; rg -n 'P05|RUN-0|usage limit|time limit' packages/runner/SPEC.md | tail -30; ls -l .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 532ms:
fireweed 0.1.0

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
?? packages/runner/internal/
go version go1.27.2 darwin/arm64
packages/project/internal/fixture/load_at.go:17:func LoadAt(root, revision string) (model.Snapshot, error) {
packages/project/internal/queue/queue.go:19:func Queue(root string) ([]model.Change, error) {
packages/project/internal/queue/check.go:15:func Check(root string) []model.Diagnostic {
190:### RUN-039
195:### RUN-040
200:### RUN-041
203:### RUN-042
206:### RUN-043
207:On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to RUN-044.
210:### RUN-044
215:### RUN-045
219:### RUN-046
223:### RUN-047
227:### RUN-048
231:### RUN-049
234:### RUN-050
237:### RUN-051
240:### RUN-052
243:### RUN-053
246:### RUN-054
249:### RUN-055
252:### RUN-056
255:### RUN-057
258:### RUN-058
261:### RUN-059
264:### RUN-060
268:### RUN-061
271:### RUN-062
274:### RUN-063
277:### RUN-064
280:### RUN-065
283:### RUN-066
286:### RUN-067
total 0

codex
The runner implementation is in place, and the retained eval hook is running. I’m reviewing the paths where cancellation, rejected audit edits, or cleanup failures must preserve work or return a failure.
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed verify; printf '"'controller verify status=%s\\n' \""'$?"; sed -n '"'83,155p;155,236p;236,318p' packages/runner/evals/time_limit_test.go; git status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 590ms:
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.
controller verify status=1
func Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied(t *testing.T) {
	f := newRunFixture(t, 2, false, false)
	timeFastHook(t, f)
	timeAgentAfterWork(t, f, `if [ "$module" = alpha ]; then sleep 1; fi
if [ "$module" = beta ] && [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/beta/evals/time-work.txt
  sleep 6
fi
`)
	audit := nextRead(t, f.agents, "auditor.sh")
	nextWrite(t, f.agents, "auditor.sh", strings.Replace(audit, "sleep 0.12", "sleep 1", 1))
	req := f.request(0, false)
	req.For = 10 * time.Second
	start := time.Now()
	r, err := runner.Run(context.Background(), req)
	elapsed := time.Since(start)
	timeStop(t, err, "10s")
	if elapsed > 14*time.Second {
		t.Errorf("total deadline waited %s; budget may have reset after alpha", elapsed)
	}
	if len(r.Applied) == 1 {
		assertRunApplied(t, f, r.Applied, "alpha")
		timeHeldWork(t, f, "beta", "evals", "packages/beta/evals/time-work.txt", "TIME_HELD_WORK", nextGit(t, f.root, "rev-parse", "HEAD"))
	} else {
		t.Errorf("Applied count = %d, want one retained alpha", len(r.Applied))
	}
	if got := runInvocations(t, f); len(got) != 4 || !strings.HasPrefix(got[3], "beta|evals|") {
		t.Errorf("invocations through held beta = %#v", got)
	}
}

func Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	req := f.request(0, false)
	req.For = 0
	r, err := runner.Run(context.Background(), req)
	if err != nil || !r.Empty || runner.IsTimeStop(err) {
		t.Fatalf("zero For run = %#v, %v", r, err)
	}
	assertRunApplied(t, f, r.Applied, "alpha")
}

func Test_RUN_037_RUN_039_AgentDeadlineRetainsHeldWork(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	timeAgentAfterWork(t, f, `if [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/alpha/evals/time-work.txt
  sleep 4
fi
`)
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	original, startErr := workflow.Start(f.root, "alpha")
	if startErr != nil {
		t.Fatalf("start interrupted agent step: %v", startErr)
	}
	req := f.request(0, false)
	req.For = 2 * time.Second
	r, err := runner.Run(context.Background(), req)
	timeStop(t, err, "2s")
	if err != nil && !runner.IsTimeStop(fmt.Errorf("outer: %w", err)) {
		t.Errorf("wrapped time stop classification = %v", err)
	}
	if len(r.Applied) != 0 {
		t.Errorf("interrupted agent applied %d changes", len(r.Applied))
	} else {
		held := timeHeldWork(t, f, "alpha", "evals", "packages/alpha/evals/time-work.txt", "TIME_HELD_WORK", head)
		if held.Run != original.Run || held.ClaimOID != original.ClaimOID || held.Phase != original.Phase || held.Change.Added[0].ID != original.Change.Added[0].ID || held.Head != original.Head {
			t.Errorf("interrupted agent identity changed: held=%#v, original=%#v", held, original)
		}
	}
}


func Test_RUN_037_VerificationDeadlineRetainsHeldWork(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	hook := nextRead(t, f.root, ".fireweed/eval.sh")
	anchor := "if [ \"$phase\" = evals ]; then result="
	addition := `if [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/alpha/evals/time-work.txt
  sleep 4
fi
`
	if !strings.Contains(hook, anchor) {
		t.Fatal("fixture hook insertion point missing")
	}
	nextWrite(t, f.root, ".fireweed/eval.sh", strings.Replace(hook, anchor, addition+anchor, 1))
	nextCommit(t, f.root, "bounded verification interruption")
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	original, startErr := workflow.Start(f.root, "alpha")
	if startErr != nil {
		t.Fatalf("start interrupted verification step: %v", startErr)
	}
	req := f.request(0, false)
	req.For = 2 * time.Second
	r, err := runner.Run(context.Background(), req)
	invocations, invocationErr := os.ReadFile(filepath.Join(f.agents, "invocations"))
	workPath := filepath.Join(f.root, "packages/alpha/evals/added_test.go")
	workInfo, workStatErr := os.Stat(workPath)
	work, workReadErr := os.ReadFile(workPath)
	markerPath := filepath.Join(f.root, "packages/alpha/evals/time-work.txt")
	markerInfo, markerStatErr := os.Stat(markerPath)
	marker, markerReadErr := os.ReadFile(markerPath)
	reportPath := filepath.Join(f.root, ".fireweed", "runs", original.Run, "evals", "alpha.json")
	reportInfo, reportStatErr := os.Stat(reportPath)
	report, reportReadErr := os.ReadFile(reportPath)
	held, heldErr := workflow.Current(f.root, "alpha")
	readGit := func(args ...string) (string, error) {
		command := exec.Command("git", args...)
		command.Dir = f.root
		output, commandErr := command.Output()
		return strings.TrimSpace(string(output)), commandErr
	}
	actualHead, headErr := readGit("rev-parse", "HEAD")
	claimRef, claimErr := readGit("for-each-ref", "--format=%(refname)", "refs/fireweed/claims")
	commits, commitsErr := readGit("rev-list", "--count", head+"..HEAD")
	t.Logf("RUN-037 verification stop: result=%#v error=%v clean_time_stop=%t", r, err, runner.IsTimeStop(err))
	t.Logf("agent invocations: read_error=%v content=%q", invocationErr, invocations)
	t.Logf("agent work %s: exists=%t stat_error=%v read_error=%v content=%q", workPath, workStatErr == nil && workInfo != nil, workStatErr, workReadErr, work)
	t.Logf("hook marker %s: exists=%t stat_error=%v read_error=%v content=%q", markerPath, markerStatErr == nil && markerInfo != nil, markerStatErr, markerReadErr, marker)
	t.Logf("verification report %s: exists=%t stat_error=%v read_error=%v content=%q", reportPath, reportStatErr == nil && reportInfo != nil, reportStatErr, reportReadErr, report)
	t.Logf("held identity: state=%#v read_error=%v original phase=%q run=%q claim=%q assigned=%#v HEAD=%q; actual HEAD=%q error=%v; claim refs=%q error=%v; commits=%q error=%v", held, heldErr, original.Phase, original.Run, original.ClaimOID, original.Change.Added, original.Head, actualHead, headErr, claimRef, claimErr, commits, commitsErr)
	timeStop(t, err, "2s")
	if len(r.Applied) != 0 {
		t.Errorf("interrupted verification applied %d changes", len(r.Applied))
	}
	if markerReadErr != nil || !strings.Contains(string(marker), "TIME_HELD_WORK") {
		t.Errorf("interrupted alpha work missing: marker=%q read_error=%v", marker, markerReadErr)
	}
	if workStatErr != nil {
		t.Errorf("agent work lost before interrupted verification: %v", workStatErr)
	}
	if commitsErr != nil || commits != "0" {
		t.Errorf("interrupted step committed %s times: %v", commits, commitsErr)
	}
	if heldErr != nil {
		t.Errorf("held state: %v", heldErr)
	} else {
		if held.Phase != "evals" || held.Run == "" || held.Head != head || held.ClaimOID == "" || held.Snapshot.Root != f.root || len(held.Change.Added) != 1 || (len(held.Change.Added) == 1 && held.Change.Added[0].ID != "ALP-002") {
			t.Errorf("held identity changed: phase=%q run=%q head=%q claim=%q added=%#v snapshot=%q", held.Phase, held.Run, held.Head, held.ClaimOID, held.Change.Added, held.Snapshot.Root)
		}
		if held.Run != original.Run || held.ClaimOID != original.ClaimOID || held.Phase != original.Phase || len(held.Change.Added) != len(original.Change.Added) || (len(held.Change.Added) == 1 && len(original.Change.Added) == 1 && held.Change.Added[0].ID != original.Change.Added[0].ID) || held.Head != original.Head {
			t.Errorf("interrupted verification identity changed: held=%#v, original=%#v", held, original)
		}
	}
	if headErr != nil || actualHead != head {
		t.Errorf("HEAD = %s, want %s: %v", actualHead, head, headErr)
	}
	if claimErr != nil || claimRef != "refs/fireweed/claims/alpha" {
		t.Errorf("held claim ref = %q: %v", claimRef, claimErr)
	}
}

func Test_RUN_038_UsageWaitPreflightStopsBeforeWaiting(t *testing.T) {
func Test_RUN_038_UsageWaitPreflightStopsBeforeWaiting(t *testing.T) {
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	usagePolicy(t, f.root, "limit_wait: 1m\nlimit_budget: 2m\n")
	runUsageAgent(t, f, "alpha/evals")
	head := nextGit(t, f.root, "rev-parse", "HEAD")
	req := f.request(0, false)
	req.For = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	start := time.Now()
	r, err := runner.Run(ctx, req)
	elapsed := time.Since(start)
	timeStop(t, err, "5s")
	if elapsed >= 3*time.Second {
		t.Errorf("usage wait preflight took %s", elapsed)
	}
	if len(r.Applied) != 0 {
		t.Errorf("preflight applied %d changes", len(r.Applied))
	} else {
		timeHeldWork(t, f, "alpha", "evals", "packages/alpha/evals/usage_test.go", "USAGE_PARTIAL_WORK", head)
	}
}

func Test_RUN_039_ClassifierAndParentCancellation(t *testing.T) {
	if runner.IsTimeStop(nil) || runner.IsTimeStop(errors.New("ordinary")) || runner.IsTimeStop(context.Canceled) {
		t.Error("ordinary, nil or parent cancellation classified as time stop")
	}
	f := newRunFixture(t, 1, false, false)
	timeFastHook(t, f)
	usagePolicy(t, f.root, "limit_wait: 1m\nlimit_budget: 0s\n")
	runUsageAgent(t, f, "alpha/evals")
	req := f.request(0, false)
	_, usageErr := runner.Run(context.Background(), req)
	if !runner.IsUsageStop(usageErr) || runner.IsTimeStop(usageErr) {
		t.Errorf("usage stop classification = %v", usageErr)
	}

	other := newRunFixture(t, 1, false, false)
	timeFastHook(t, other)
	timeAgentAfterWork(t, other, `if [ "$phase" = evals ]; then sleep 8; fi
`)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	parentReq := other.request(0, false)
	parentReq.For = 10 * time.Second
	_, parentErr := runner.Run(ctx, parentReq)
	if !errors.Is(parentErr, context.DeadlineExceeded) || runner.IsTimeStop(parentErr) {
		t.Errorf("parent cancellation = %v", parentErr)
	}
}

func Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop(t *testing.T) {
	f := runBlockedDependencyFixture(t)
	timeFastHook(t, f)
	timeAgentAfterWork(t, f, `if [ "$module" = gamma ] && [ "$phase" = evals ]; then
  printf 'TIME_HELD_WORK\n' > packages/gamma/evals/time-work.txt
  while :; do sleep 1; done
fi
`)
	req := f.request(0, false)
	req.For = 8 * time.Second
	r, err := runner.Run(context.Background(), req)
	if err == nil || !strings.Contains(err.Error(), "P2-01: no fail-first case for ALP-002") || !strings.Contains(err.Error(), "refs/fireweed/attempts/alpha/1") || !strings.Contains(err.Error(), "8s") || runner.IsTimeStop(err) {
		t.Errorf("combined block and time failure: Applied=%d, err=%v", len(r.Applied), err)
	}
	if saved := nextGit(t, f.root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/alpha/1"); saved == "" {
		t.Error("blocked alpha lost saved attempt ref")
	}
	if len(r.Applied) != 0 {
		t.Errorf("blocked run applied %d changes", len(r.Applied))
	} else {
		timeHeldWork(t, f, "gamma", "evals", "packages/gamma/evals/time-work.txt", "TIME_HELD_WORK", nextGit(t, f.root, "rev-parse", "HEAD"))
	}
	assertRunBetaWaiting(t, f)
}
?? packages/runner/internal/

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/drill/workspace.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/audit.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/feedback.go
diff --git a/packages/runner/internal/drill/workspace.go b/packages/runner/internal/drill/workspace.go
index 1d4d8fdee6e026ca3d5c0afa507811862ab75582..91c3753e529df7737b1564ff686d6f32ed6edf0c
--- a/packages/runner/internal/drill/workspace.go
+++ b/packages/runner/internal/drill/workspace.go
@@ -123,15 +123,19 @@
 }
 
 func candidateFiles(root string, baseline map[string]support.File, snapshot project.Snapshot, target project.Module) (map[string]support.File, error) {
-	for name, original := range baseline {
-		current, err := support.ReadFile(filepath.Join(root, name))
-		if err != nil || current != original {
-			return nil, fmt.Errorf("drill reconstruction modified or deleted input %s", name)
-		}
-	}
 	candidate := map[string]support.File{}
+	seen := map[string]bool{}
 	err := support.Walk(root, func(name string, entry fs.DirEntry) error {
-		if _, retained := baseline[name]; retained || entry.IsDir() || !disposable(snapshot, target, name) {
+		original, retained := baseline[name]
+		if retained {
+			current, err := support.ReadFile(filepath.Join(root, name))
+			if err != nil || current != original {
+				return fmt.Errorf("drill reconstruction modified input %s", name)
+			}
+			seen[name] = true
+			return nil
+		}
+		if entry.IsDir() || !disposable(snapshot, target, name) {
 			return nil
 		}
 		file, err := support.ReadFile(filepath.Join(root, name))
@@ -140,7 +144,15 @@
 		}
 		return err
 	})
-	return candidate, err
+	if err != nil {
+		return nil, err
+	}
+	for name := range baseline {
+		if !seen[name] {
+			return nil, fmt.Errorf("drill reconstruction deleted input %s", name)
+		}
+	}
+	return candidate, nil
 }
 
 func importCandidate(scratch string, candidate map[string]support.File) error {
diff --git a/packages/runner/internal/next/audit.go b/packages/runner/internal/next/audit.go
index 3690e308b5e5f3e880481704291aee2e7db361c1..612ea7a9269f7131bd4c9827b1dbf44a60daa18a
--- a/packages/runner/internal/next/audit.go
+++ b/packages/runner/internal/next/audit.go
@@ -127,7 +127,7 @@
 		return 0, 0, false, true
 	}
 	pos := len(raw) - len(doc.Body)
-	for line := range strings.SplitAfterSeq(doc.Body, "\n") {
+	for _, line := range strings.SplitAfter(doc.Body, "\n") {
 		if strings.HasPrefix(line, "## ") {
 			if found && end == 0 {
 				end = pos
@@ -164,7 +164,7 @@
 			if after[:start]+after[ae:] == before {
 				return true
 			}
-			if start < 2 || after[start-1] != '\n' || after[start-2] != '\n' {
+			if start == 0 || after[start-1] != '\n' {
 				return false
 			}
 			start--
diff --git a/packages/runner/internal/next/feedback.go b/packages/runner/internal/next/feedback.go
index 6aff47fe7c774980ca589300dfc4ad339cdcc3d2..7f88faf12bd8d89a53525c3953cac7f60163f027
--- a/packages/runner/internal/next/feedback.go
+++ b/packages/runner/internal/next/feedback.go
@@ -43,19 +43,13 @@
 				continue
 			}
 			failure := c.Failure
-			for i := range maxFailure {
-				_, size := utf8.DecodeRuneInString(failure)
-				failure = failure[size:]
-				if failure == "" {
+			runes := 0
+			for offset := range failure {
+				if runes == maxFailure {
+					failure = failure[:offset] + "\n[Failure text truncated]"
 					break
 				}
-				if i == maxFailure-1 {
-					failure = c.Failure[:len(c.Failure)-len(failure)] + "\n[Failure text truncated]"
-					break
-				}
-			}
-			if failure == "" {
-				failure = c.Failure
+				runes++
 			}
 			entry := fmt.Sprintf("\nModule: %s\nClassname: %s\nTestcase: %s\n%s\n", module, c.Classname, c.Name, failure)
 			length := utf8.RuneCountInString(entry)

exec
/bin/zsh -lc "gofmt -w packages/runner/internal; ls -l .fireweed/runs/drill; ps -Ao pid,etime,comm | rg 'evals.test|fireweed/hooks|/bin/go' | tail -20" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 540ms:
total 0
zsh:1: operation not permitted: ps

exec
/bin/zsh -lc "rg -n 'verify|eval.order|verification' packages/runner/evals/next_test.go | tail -55; rg -n 'Timeout|attempts|selected|hasEval' packages/runner/evals/run_test.go | head -50; sed -n '180,245p' packages/runner/evals/agent_logs_test.go; sed -n '72,158p' packages/gitstore/internal/local/local.go; cat AGENTS.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 510ms:
713:			t.Errorf("%s retry prompt lacks real verification feedback %q", phase.name, phase.feedback)
743:		t.Error("third attempt lacks the second verification's actual Go output")
161:	return runner.RunRequest{Root: f.root, Bin: "/fixture/run-bin", AllowUnmerged: allow, Timeout: 20 * time.Second, Max: max}
270:	const ref = "refs/fireweed/attempts/alpha/1"
		})
	}
}

func Test_RUN_066_InvokeAgentLogContextDefaultsAndExplicitFields(t *testing.T) {
	snap := agentFixture(t)
	snap.Config["agent"] = map[string]any{"command": []any{"/bin/sh", "-c", `printf 'default output'; printf 'default error' >&2`}}
	var previous string
	for i := 0; i < 2; i++ {
		result, err := runner.InvokeAgent(context.Background(), request(snap, "worker", ""))
		if err != nil {
			t.Fatalf("default invocation: %v", err)
		}
		path := result.LogPath
		rel, relErr := filepath.Rel(filepath.Join(snap.Root, ".fireweed", "runs"), path)
		if relErr != nil || strings.HasPrefix(rel, "..") {
			t.Errorf("default log %q outside Snapshot.Root: %v", path, relErr)
			continue
		}
		parts := strings.Split(filepath.ToSlash(rel), "/")
		if len(parts) != 2 || parts[0] == "" || parts[1] != "worker-1.log" {
			t.Errorf("default log relative path = %q", rel)
			continue
		}
		assertLogFile(t, path, snap.Root, parts[0], "worker", 1, "default output", "default error")
		if i == 1 && path == previous {
			t.Error("default run identifier was reused")
		}
		previous = path
	}
	chosenRoot := t.TempDir()
	explicit := request(snap, "worker", "")
	explicit.LogRoot = chosenRoot
	explicit.Run = "chosen-run"
	explicit.Phase = "chosen-phase"
	explicit.Attempt = 3
	result, err := runner.InvokeAgent(context.Background(), explicit)
	if err != nil {
		t.Fatalf("explicit invocation: %v", err)
	}
	assertLogFile(t, result.LogPath, chosenRoot, "chosen-run", "chosen-phase", 3, "default output", "default error")
}
func git(root string, input []byte, env map[string]string, args ...string) ([]byte, error) {
	if env == nil {
		env = make(map[string]string)
	}
	env["GIT_NO_LAZY_FETCH"] = "1"
	env["GIT_TERMINAL_PROMPT"] = "0"
	argv := append([]string{"git", "-C", root, "-c", "core.hooksPath=/dev/null"}, args...)
	result, err := Run(context.Background(), model.Request{Args: argv, Env: env, Input: input, Timeout: gitTimeout})
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
	}
	return []byte(result.Stdout), nil
}

func gitText(root string, args ...string) (string, error) {
	out, err := git(root, nil, nil, args...)
	return strings.TrimSpace(string(out)), err
}

func canonical(path string) (string, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(path)
}

func Inspect(dir string) (model.Repository, error) {
	var repo model.Repository
	root, err := gitText(dir, "rev-parse", "--show-toplevel")
	if err != nil {
		return repo, err
	}
	repo.Root, err = canonical(root)
	if err != nil {
		return repo, err
	}
	common, err := gitText(dir, "rev-parse", "--git-common-dir")
	if err != nil {
		return repo, err
	}
	if !filepath.IsAbs(common) {
		common = filepath.Join(dir, common)
	}
	repo.CommonDir, err = canonical(common)
	if err != nil {
		return repo, err
	}
	repo.Head, err = gitText(dir, "rev-parse", "--verify", "HEAD")
	if err != nil {
		return repo, err
	}
	repo.Branch, _ = gitText(dir, "symbolic-ref", "--quiet", "--short", "HEAD")
	return repo, nil
}

func BaseRef(root, remote, base string) (string, error) {
	if remote == "" {
		remote = "origin"
	}
	if base == "" {
		remoteHead := "refs/remotes/" + remote + "/HEAD"
		resolved, err := gitText(root, "symbolic-ref", "--quiet", remoteHead)
		if err == nil && strings.HasPrefix(resolved, "refs/remotes/"+remote+"/") {
			base = strings.TrimPrefix(resolved, "refs/remotes/"+remote+"/")
		} else {
			base = "main"
		}
	}
	for _, name := range []string{"refs/remotes/" + remote + "/" + base, "refs/heads/" + base} {
		if _, exists, err := Ref(root, name); err != nil {
			return "", err
		} else if exists {
			return name, nil
		}
	}
	return "", fmt.Errorf("base %q not found", base)
}

func Changes(root, base string) ([]model.Change, error) {
	byPath := make(map[string]model.Change)
	for _, args := range [][]string{
		{"diff", "--cached", "--name-status", "-z", "-M", base, "--"},
		{"diff", "--name-status", "-z", "-M", "--"},
	} {
		out, err := git(root, nil, nil, args...)
		if err != nil {
<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step.
<!-- fireweed:end -->

## Rebuild controller

Drive the workflow with the controller `.fireweed/bin/fireweed` (see SYSTEM.md), not `bin/fireweed`, which is the product under test. Put `.fireweed/bin` first on worker PATH and use `GOCACHE=/tmp/fireweed-rebuild-cache`. Never switch controllers while a run or claim is active.

Use `fireweed verify` to check your work. Never commit during a worker step; fireweed owns commits. Read SYSTEM.md (including its scope guardrails) and docs/rebuild/PROGRESS.md before work; docs/rebuild/DECISIONS.md retains design clarifications. Do not consult `rebuild-v1` or deleted code as an implementation oracle.

2026-11-10T19:32:24.371065Z ERROR codex_core::tools::router: error=apply_patch verification failed: invalid patch: multiple operations target /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/audit.go
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/drill/workspace.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/invoke/agent.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/audit.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/next.go
diff --git a/packages/runner/internal/drill/workspace.go b/packages/runner/internal/drill/workspace.go
index 1d4d8fdee6e026ca3d5c0afa507811862ab75582..91d32c036fcc7bf1a77a53309d73e2b3fc580a31
--- a/packages/runner/internal/drill/workspace.go
+++ b/packages/runner/internal/drill/workspace.go
@@ -123,15 +123,19 @@
 }
 
 func candidateFiles(root string, baseline map[string]support.File, snapshot project.Snapshot, target project.Module) (map[string]support.File, error) {
-	for name, original := range baseline {
-		current, err := support.ReadFile(filepath.Join(root, name))
-		if err != nil || current != original {
-			return nil, fmt.Errorf("drill reconstruction modified or deleted input %s", name)
-		}
-	}
 	candidate := map[string]support.File{}
+	seen := map[string]bool{}
 	err := support.Walk(root, func(name string, entry fs.DirEntry) error {
-		if _, retained := baseline[name]; retained || entry.IsDir() || !disposable(snapshot, target, name) {
+		original, retained := baseline[name]
+		if retained {
+			current, err := support.ReadFile(filepath.Join(root, name))
+			if err != nil || current != original {
+				return fmt.Errorf("drill reconstruction modified input %s", name)
+			}
+			seen[name] = true
+			return nil
+		}
+		if entry.IsDir() || !disposable(snapshot, target, name) {
 			return nil
 		}
 		file, err := support.ReadFile(filepath.Join(root, name))
@@ -140,12 +144,39 @@
 		}
 		return err
 	})
-	return candidate, err
+	if err != nil {
+		return nil, err
+	}
+	for name := range baseline {
+		if !seen[name] {
+			return nil, fmt.Errorf("drill reconstruction deleted input %s", name)
+		}
+	}
+	return candidate, nil
 }
 
 func importCandidate(scratch string, candidate map[string]support.File) error {
 	for name, file := range candidate {
-		if err := support.WriteFile(filepath.Join(scratch, name), file); err != nil {
+		destination := filepath.Join(scratch, name)
+		for dir := filepath.Dir(destination); dir != scratch; dir = filepath.Dir(dir) {
+			info, err := os.Lstat(dir)
+			if os.IsNotExist(err) {
+				continue
+			}
+			if err != nil {
+				return err
+			}
+			if !info.IsDir() {
+				return fmt.Errorf("rebuilt path %s crosses a retained file", name)
+			}
+		}
+		if _, err := os.Lstat(destination); !os.IsNotExist(err) {
+			if err != nil {
+				return err
+			}
+			return fmt.Errorf("rebuilt path %s overlaps retained content", name)
+		}
+		if err := support.WriteFile(destination, file); err != nil {
 			return err
 		}
 	}
diff --git a/packages/runner/internal/invoke/agent.go b/packages/runner/internal/invoke/agent.go
index c769769eeba3502f3db5372fc075d65096b5b1ee..85bfd2a14c8d691b6834a9f7baa16e044157bd74
--- a/packages/runner/internal/invoke/agent.go
+++ b/packages/runner/internal/invoke/agent.go
@@ -30,11 +30,23 @@
 	}
 	values := map[string]string{"prompt": request.Prompt}
 	if containsTemplate(command, "{git_common_dir}") || containsTemplate(wrapper, "{git_common_dir}") {
-		repo, inspectErr := gitstore.Inspect(request.Snapshot.Root)
-		if inspectErr != nil {
-			return result, inspectErr
+		root := request.Snapshot.Root
+		dir, gitErr := support.Git(ctx, root, "rev-parse", "--git-common-dir")
+		if gitErr != nil {
+			return result, gitErr
+		}
+		dir = strings.TrimSuffix(dir, "\n")
+		if !filepath.IsAbs(dir) {
+			dir = filepath.Join(root, dir)
+		}
+		dir, err = filepath.Abs(dir)
+		if err != nil {
+			return result, err
+		}
+		values["git_common_dir"], err = filepath.EvalSymlinks(dir)
+		if err != nil {
+			return result, err
 		}
-		values["git_common_dir"] = repo.CommonDir
 	}
 	if containsTemplate(command, "{prompt_file}") || containsTemplate(wrapper, "{prompt_file}") {
 		f, createErr := os.CreateTemp("", "fireweed-prompt-*")
diff --git a/packages/runner/internal/next/audit.go b/packages/runner/internal/next/audit.go
index 3690e308b5e5f3e880481704291aee2e7db361c1..34b41f0442d563ee741c7135ecfffe066739a8af
--- a/packages/runner/internal/next/audit.go
+++ b/packages/runner/internal/next/audit.go
@@ -2,7 +2,6 @@
 
 import (
 	"context"
-	"errors"
 	"fmt"
 	"io/fs"
 	"os"
@@ -22,7 +21,11 @@
 	if err != nil {
 		return result, false, err
 	}
-	defer func() { err = errors.Join(err, os.RemoveAll(root)) }()
+	defer func() {
+		if cleanupErr := os.RemoveAll(root); cleanupErr != nil {
+			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("remove audit workspace: %w", cleanupErr)}}
+		}
+	}()
 	held := request.Snapshot.Root
 	change := state.Change.Path
 	original, err := support.ReadFile(filepath.Join(held, change))
@@ -127,7 +130,7 @@
 		return 0, 0, false, true
 	}
 	pos := len(raw) - len(doc.Body)
-	for line := range strings.SplitAfterSeq(doc.Body, "\n") {
+	for _, line := range strings.SplitAfter(doc.Body, "\n") {
 		if strings.HasPrefix(line, "## ") {
 			if found && end == 0 {
 				end = pos
@@ -164,7 +167,7 @@
 			if after[:start]+after[ae:] == before {
 				return true
 			}
-			if start < 2 || after[start-1] != '\n' || after[start-2] != '\n' {
+			if start == 0 || after[start-1] != '\n' {
 				return false
 			}
 			start--
diff --git a/packages/runner/internal/next/feedback.go b/packages/runner/internal/next/feedback.go
index 6aff47fe7c774980ca589300dfc4ad339cdcc3d2..7f88faf12bd8d89a53525c3953cac7f60163f027
--- a/packages/runner/internal/next/feedback.go
+++ b/packages/runner/internal/next/feedback.go
@@ -43,19 +43,13 @@
 				continue
 			}
 			failure := c.Failure
-			for i := range maxFailure {
-				_, size := utf8.DecodeRuneInString(failure)
-				failure = failure[size:]
-				if failure == "" {
+			runes := 0
+			for offset := range failure {
+				if runes == maxFailure {
+					failure = failure[:offset] + "\n[Failure text truncated]"
 					break
 				}
-				if i == maxFailure-1 {
-					failure = c.Failure[:len(c.Failure)-len(failure)] + "\n[Failure text truncated]"
-					break
-				}
-			}
-			if failure == "" {
-				failure = c.Failure
+				runes++
 			}
 			entry := fmt.Sprintf("\nModule: %s\nClassname: %s\nTestcase: %s\n%s\n", module, c.Classname, c.Name, failure)
 			length := utf8.RuneCountInString(entry)
diff --git a/packages/runner/internal/next/next.go b/packages/runner/internal/next/next.go
index 9732acae9d5a29bd014a63e49b164a25dedc08e2..d732a397ba912212484892155977ea03f755c985
--- a/packages/runner/internal/next/next.go
+++ b/packages/runner/internal/next/next.go
@@ -4,7 +4,6 @@
 	"context"
 	"errors"
 	"fmt"
-	"strings"
 
 	"github.com/lurkshark/fireweed/packages/gitstore"
 	"github.com/lurkshark/fireweed/packages/project"
@@ -183,6 +182,9 @@
 				}
 			}
 			if attempt == attempts {
+				if err := ctx.Err(); err != nil {
+					return result, err
+				}
 				block, blockErr := workflow.Block(workflow.BlockRequest{Root: root, Change: state.Change.Slug, Reason: verifyErr.Error()})
 				if block.State.Run != "" {
 					result.State = block.State
@@ -230,6 +232,6 @@
 	}
 	return &project.Diagnostic{
 		Code: "P05", Warning: true, Path: state.Change.Path,
-		Message: strings.Join([]string{state.Phase, "worker HEAD reset to the phase base, retaining index and working-tree work"}, " "),
+		Message: fmt.Sprintf("%s worker HEAD reset to the phase base, retaining index and working-tree work", state.Phase),
 	}, nil
 }

exec
/bin/zsh -lc 'gofmt -w packages/runner/internal; date -u; ls -lh .fireweed/runs/drill' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 578ms:
Tue Nov 10 19:32:43 UTC 2026
total 0

codex
The controller’s `verify` command requires a held change, which this drill workspace has no claim for. The supplied runner eval hook is the verification path here, and it’s still running.
exec
/bin/zsh -lc "sed -n '254,285p' packages/gitstore/internal/local/local.go; sed -n '1,90p' packages/project/internal/queue/queue.go; ls -lh .fireweed/runs/drill; date -u" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 2291ms:
	return strings.TrimSpace(string(out)), err
}

func ReadBlob(root, oid string) ([]byte, error) {
	return git(root, nil, nil, "cat-file", "blob", oid)
}

func Ref(root, name string) (string, bool, error) {
	if _, err := git(root, nil, nil, "check-ref-format", name); err != nil {
		return "", false, err
	}
	out, err := git(root, nil, nil, "show-ref", "--verify", "--hash", name)
	if err == nil {
		return strings.TrimSpace(string(out)), true, nil
	}
	// Git versions differ on the status returned for a missing valid ref.
	var exit *exec.ExitError
	if errors.As(err, &exit) {
		return "", false, nil
	}
	return "", false, err
}

func CASRef(root, name, old, new string) error {
	const zeroSHA1 = "0000000000000000000000000000000000000000"
	if old == "" {
		old = zeroSHA1
	}
	if new == "" && old == zeroSHA1 {
		// Git's delete mode treats an all-zero old ID as unconditional.
		if strings.ContainsRune(name, '\x00') {
			return errors.New("ref name contains NUL")
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
total 0
Tue Nov 10 19:33:35 UTC 2026

exec
/bin/zsh -lc "sed -n '755,791p;883,920p' packages/workflow/internal/lifecycle/stub.go; sed -n '322,360p' packages/runner/evals/next_test.go; cat packages/runner/evals/approval_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 912ms:
	line := fmt.Sprintf("- %s %s agent=%s fireweed=%s playbook=%x tests=%d failing=%d passing=%d\n", phase, time.Now().UTC().Format(time.RFC3339Nano), agent, assets.Version(), hash, total, failing, passing)
	body, _ := contracts.Section(ch.Document.Body, "Provenance")
	ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Provenance", body+line)
	return ch, nil
}
func commitPhase(s model.State, c claim, oid string, ch project.Change, paths []string, subject, next string) (model.State, error) {
	root := s.Snapshot.Root
	ch.Document.Frontmatter["phase"] = next
	rendered, e := contracts.RenderDocument(ch.Document)
	if e != nil {
		return s, e
	}
	if e = write(root, ch.Path, []byte(rendered)); e != nil {
		return s, e
	}
	head, e := gitstore.Commit(root, subject, paths)
	if e != nil {
		return s, e
	}
	c.Phase = next
	c.PhaseBase = head
	c.Audit = auditState(ch.Document)
	oid, e = storeClaim(root, c, oid)
	if e != nil {
		return s, e
	}
	fresh, e := changeBy(root, ch.Slug)
	if e != nil {
		return s, e
	}
	return state(root, fresh, c, oid, head), nil
}
func auditProcess(s model.State, c claim) (project.Change, string, bool, error) {
	ch := s.Change
	ch.Document.Frontmatter = cloneMap(ch.Document.Frontmatter)
	body, _ := contracts.Section(ch.Document.Body, "Audit")
	all := map[string]string{}
func apply(s model.State, c claim, ch project.Change, oid string) (model.State, error) {
	root := s.Snapshot.Root
	m := module(c.Snapshot, ch.Module)
	specPath := m.Dir + "/SPEC.md"
	original, e := read(root, specPath)
	if e != nil {
		return s, e
	}
	pending, e := read(root, ch.Path)
	if e != nil {
		return s, e
	}
	archive := "changes/archive/" + filepath.Base(ch.Path)
	if _, e = os.Stat(file(root, archive)); e == nil {
		return s, errors.New("archive already exists")
	}
	originals := map[string][]byte{specPath: original, ch.Path: pending, archive: nil}
	oldDoc, e := contracts.ParseDocument(string(original))
	if e != nil {
		return s, e
	}
	decision, _ := contracts.Section(oldDoc.Body, "Decisions")
	removed, _ := contracts.Section(ch.Document.Body, "Removed decisions")
	for _, line := range strings.Split(removed, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			decision = strings.ReplaceAll(decision, line+"\n", "")
			decision = strings.ReplaceAll(decision, line, "")
		}
	}
	added, _ := contracts.Section(ch.Document.Body, "Decisions")
	added = strings.TrimSpace(added)
	if added != "" {
		if !strings.HasSuffix(decision, "\n") {
			decision += "\n"
		}
		decision += fmt.Sprintf("\n- %s: %s\n", ch.Slug, added)
	}
	f.noAgent(t)
}

func Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents(t *testing.T) {
	t.Run("unmerged", func(t *testing.T) {
		f := newNextFixture(t, true, false, "normal")
		before := nextGit(t, f.root, "rev-parse", "HEAD")
		_, err := f.run(false)
		if err == nil || !strings.Contains(strings.ToLower(err.Error()), "approv") {
			t.Errorf("unapproved Next error = %v", err)
		}
		if nextGit(t, f.root, "rev-parse", "HEAD") != before || strings.Contains(nextRead(t, f.root, nextChange), "run:") {
			t.Error("unapproved change was claimed or assigned")
		}
		if out := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); out != "" {
			t.Errorf("unapproved claim = %s", out)
		}
		f.noAgent(t)
	})
	t.Run("start_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextWrite(t, f.root, "stray.txt", "dirty\n")
		_, err := f.run(false)
		if err == nil || !strings.Contains(err.Error(), "P01") {
			t.Errorf("dirty start error = %v", err)
		}
		f.noAgent(t)
	})
	t.Run("first_unblocked", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exit")
		nextWrite(t, f.root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
		nextWrite(t, f.root, "changes/0000-blocked.md", "---\nmodule: beta\nblocked: human review\n---\n\n## Motivation\nBlocked work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		nextCommit(t, f.root, "blocked first change")
		_, err := f.run(false)
		if err == nil {
			t.Error("expected worker failure after eligible start")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Change.Slug != "next" || state.Phase != "evals" {
package evals_test

import (
	"errors"
	"testing"

	"github.com/lurkshark/fireweed/packages/runner"
)

func Test_RUN_041_UnapprovedChangeRequiresApproval(t *testing.T) {
	f := newNextFixture(t, true, false, "normal")
	_, err := f.run(false)
	if err == nil {
		t.Fatal("Next accepted a change absent from the local and shared base")
	}
	if !runner.IsApprovalRequired(err) {
		t.Errorf("IsApprovalRequired(%v) = false", err)
	}
	f.noAgent(t)
}

func Test_RUN_042_OtherNextErrorsDoNotRequireApproval(t *testing.T) {
	if runner.IsApprovalRequired(nil) {
		t.Error("nil error requires approval")
	}
	if runner.IsApprovalRequired(errors.New("requires approval")) {
		t.Error("unrelated error with approval wording requires approval")
	}

	t.Run("start_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextWrite(t, f.root, "stray.txt", "dirty\n")
		_, err := f.run(false)
		if err == nil {
			t.Fatal("expected dirty-worktree start failure")
		}
		if runner.IsApprovalRequired(err) {
			t.Errorf("start failure requires approval: %v", err)
		}
		f.noAgent(t)
	})

	t.Run("agent_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exit")
		result, err := f.run(false)
		if err == nil || len(result.Steps) != 1 || result.Steps[0].Agent.ExitCode != 9 {
			t.Fatalf("expected agent exit 9, got %#v, %v", result, err)
		}
		if runner.IsApprovalRequired(err) {
			t.Errorf("agent failure requires approval: %v", err)
		}
	})
}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/next.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/run/run.go
diff --git a/packages/runner/internal/drill/workspace.go b/packages/runner/internal/drill/workspace.go
index 1d4d8fdee6e026ca3d5c0afa507811862ab75582..91d32c036fcc7bf1a77a53309d73e2b3fc580a31
--- a/packages/runner/internal/drill/workspace.go
+++ b/packages/runner/internal/drill/workspace.go
@@ -123,15 +123,19 @@
 }
 
 func candidateFiles(root string, baseline map[string]support.File, snapshot project.Snapshot, target project.Module) (map[string]support.File, error) {
-	for name, original := range baseline {
-		current, err := support.ReadFile(filepath.Join(root, name))
-		if err != nil || current != original {
-			return nil, fmt.Errorf("drill reconstruction modified or deleted input %s", name)
-		}
-	}
 	candidate := map[string]support.File{}
+	seen := map[string]bool{}
 	err := support.Walk(root, func(name string, entry fs.DirEntry) error {
-		if _, retained := baseline[name]; retained || entry.IsDir() || !disposable(snapshot, target, name) {
+		original, retained := baseline[name]
+		if retained {
+			current, err := support.ReadFile(filepath.Join(root, name))
+			if err != nil || current != original {
+				return fmt.Errorf("drill reconstruction modified input %s", name)
+			}
+			seen[name] = true
+			return nil
+		}
+		if entry.IsDir() || !disposable(snapshot, target, name) {
 			return nil
 		}
 		file, err := support.ReadFile(filepath.Join(root, name))
@@ -140,12 +144,39 @@
 		}
 		return err
 	})
-	return candidate, err
+	if err != nil {
+		return nil, err
+	}
+	for name := range baseline {
+		if !seen[name] {
+			return nil, fmt.Errorf("drill reconstruction deleted input %s", name)
+		}
+	}
+	return candidate, nil
 }
 
 func importCandidate(scratch string, candidate map[string]support.File) error {
 	for name, file := range candidate {
-		if err := support.WriteFile(filepath.Join(scratch, name), file); err != nil {
+		destination := filepath.Join(scratch, name)
+		for dir := filepath.Dir(destination); dir != scratch; dir = filepath.Dir(dir) {
+			info, err := os.Lstat(dir)
+			if os.IsNotExist(err) {
+				continue
+			}
+			if err != nil {
+				return err
+			}
+			if !info.IsDir() {
+				return fmt.Errorf("rebuilt path %s crosses a retained file", name)
+			}
+		}
+		if _, err := os.Lstat(destination); !os.IsNotExist(err) {
+			if err != nil {
+				return err
+			}
+			return fmt.Errorf("rebuilt path %s overlaps retained content", name)
+		}
+		if err := support.WriteFile(destination, file); err != nil {
 			return err
 		}
 	}
diff --git a/packages/runner/internal/invoke/agent.go b/packages/runner/internal/invoke/agent.go
index c769769eeba3502f3db5372fc075d65096b5b1ee..85bfd2a14c8d691b6834a9f7baa16e044157bd74
--- a/packages/runner/internal/invoke/agent.go
+++ b/packages/runner/internal/invoke/agent.go
@@ -30,11 +30,23 @@
 	}
 	values := map[string]string{"prompt": request.Prompt}
 	if containsTemplate(command, "{git_common_dir}") || containsTemplate(wrapper, "{git_common_dir}") {
-		repo, inspectErr := gitstore.Inspect(request.Snapshot.Root)
-		if inspectErr != nil {
-			return result, inspectErr
+		root := request.Snapshot.Root
+		dir, gitErr := support.Git(ctx, root, "rev-parse", "--git-common-dir")
+		if gitErr != nil {
+			return result, gitErr
+		}
+		dir = strings.TrimSuffix(dir, "\n")
+		if !filepath.IsAbs(dir) {
+			dir = filepath.Join(root, dir)
+		}
+		dir, err = filepath.Abs(dir)
+		if err != nil {
+			return result, err
+		}
+		values["git_common_dir"], err = filepath.EvalSymlinks(dir)
+		if err != nil {
+			return result, err
 		}
-		values["git_common_dir"] = repo.CommonDir
 	}
 	if containsTemplate(command, "{prompt_file}") || containsTemplate(wrapper, "{prompt_file}") {
 		f, createErr := os.CreateTemp("", "fireweed-prompt-*")
diff --git a/packages/runner/internal/next/audit.go b/packages/runner/internal/next/audit.go
index 3690e308b5e5f3e880481704291aee2e7db361c1..34b41f0442d563ee741c7135ecfffe066739a8af
--- a/packages/runner/internal/next/audit.go
+++ b/packages/runner/internal/next/audit.go
@@ -2,7 +2,6 @@
 
 import (
 	"context"
-	"errors"
 	"fmt"
 	"io/fs"
 	"os"
@@ -22,7 +21,11 @@
 	if err != nil {
 		return result, false, err
 	}
-	defer func() { err = errors.Join(err, os.RemoveAll(root)) }()
+	defer func() {
+		if cleanupErr := os.RemoveAll(root); cleanupErr != nil {
+			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("remove audit workspace: %w", cleanupErr)}}
+		}
+	}()
 	held := request.Snapshot.Root
 	change := state.Change.Path
 	original, err := support.ReadFile(filepath.Join(held, change))
@@ -127,7 +130,7 @@
 		return 0, 0, false, true
 	}
 	pos := len(raw) - len(doc.Body)
-	for line := range strings.SplitAfterSeq(doc.Body, "\n") {
+	for _, line := range strings.SplitAfter(doc.Body, "\n") {
 		if strings.HasPrefix(line, "## ") {
 			if found && end == 0 {
 				end = pos
@@ -164,7 +167,7 @@
 			if after[:start]+after[ae:] == before {
 				return true
 			}
-			if start < 2 || after[start-1] != '\n' || after[start-2] != '\n' {
+			if start == 0 || after[start-1] != '\n' {
 				return false
 			}
 			start--
diff --git a/packages/runner/internal/next/feedback.go b/packages/runner/internal/next/feedback.go
index 6aff47fe7c774980ca589300dfc4ad339cdcc3d2..7f88faf12bd8d89a53525c3953cac7f60163f027
--- a/packages/runner/internal/next/feedback.go
+++ b/packages/runner/internal/next/feedback.go
@@ -43,19 +43,13 @@
 				continue
 			}
 			failure := c.Failure
-			for i := range maxFailure {
-				_, size := utf8.DecodeRuneInString(failure)
-				failure = failure[size:]
-				if failure == "" {
+			runes := 0
+			for offset := range failure {
+				if runes == maxFailure {
+					failure = failure[:offset] + "\n[Failure text truncated]"
 					break
 				}
-				if i == maxFailure-1 {
-					failure = c.Failure[:len(c.Failure)-len(failure)] + "\n[Failure text truncated]"
-					break
-				}
-			}
-			if failure == "" {
-				failure = c.Failure
+				runes++
 			}
 			entry := fmt.Sprintf("\nModule: %s\nClassname: %s\nTestcase: %s\n%s\n", module, c.Classname, c.Name, failure)
 			length := utf8.RuneCountInString(entry)
diff --git a/packages/runner/internal/next/next.go b/packages/runner/internal/next/next.go
index 9732acae9d5a29bd014a63e49b164a25dedc08e2..58d0621c3e9b4e4680ade5cc3c4761f501bfb73f
--- a/packages/runner/internal/next/next.go
+++ b/packages/runner/internal/next/next.go
@@ -68,8 +68,14 @@
 		if _, err := support.Usage(snapshot.Config); err != nil {
 			return result, err
 		}
-		if !request.AllowUnmerged && !approved(root, snapshot, selected.Path) {
-			return result, &approvalRequired{selected.Path}
+		if !request.AllowUnmerged {
+			ok, err := approved(ctx, root, snapshot, selected.Path)
+			if err != nil {
+				return result, err
+			}
+			if !ok {
+				return result, &approvalRequired{selected.Path}
+			}
 		}
 		if err := ctx.Err(); err != nil {
 			return result, err
@@ -86,12 +92,22 @@
 		return result, err
 	}
 	diagnostics := append([]project.Diagnostic(nil), state.Diagnostics...)
-	defer func() { result.State.Diagnostics = diagnostics }()
+	defer func() {
+		if err != nil && result.State.ClaimOID != "" {
+			if held, currentErr := workflow.Current(root, result.State.Change.Slug); currentErr == nil {
+				result.State = held
+			}
+		}
+		result.State.Diagnostics = diagnostics
+	}()
 	for result.State.Phase != "applied" {
 		if err := ctx.Err(); err != nil {
 			return result, err
 		}
 		state = result.State
+		if state.Head != state.PhaseBase {
+			return result, fmt.Errorf("P05: %s HEAD differs from the held phase base", state.Phase)
+		}
 		phase := state.Phase
 		role := "worker"
 		var prompt string
@@ -139,9 +155,6 @@
 			}
 			if recoveryErr != nil {
 				return result, errors.Join(agentErr, recoveryErr)
-			}
-			if fresh, currentErr := workflow.Current(root, state.Change.Slug); currentErr == nil {
-				result.State = fresh
 			}
 			var invalid *support.CombinedFailure
 			if errors.As(agentErr, &invalid) {
@@ -183,6 +196,9 @@
 				}
 			}
 			if attempt == attempts {
+				if err := ctx.Err(); err != nil {
+					return result, err
+				}
 				block, blockErr := workflow.Block(workflow.BlockRequest{Root: root, Change: state.Change.Slug, Reason: verifyErr.Error()})
 				if block.State.Run != "" {
 					result.State = block.State
@@ -199,7 +215,7 @@
 	return result, nil
 }
 
-func approved(root string, snapshot project.Snapshot, path string) bool {
+func approved(ctx context.Context, root string, snapshot project.Snapshot, path string) (bool, error) {
 	base := support.String(snapshot.Config["base"])
 	if base == "" {
 		base = "main"
@@ -209,20 +225,31 @@
 		remote = "origin"
 	}
 	for _, ref := range []string{"refs/heads/" + base, "refs/remotes/" + remote + "/" + base} {
-		if _, err := gitstore.ReadAt(root, ref, path); err == nil {
-			return true
+		oid, exists, err := gitstore.Ref(root, ref)
+		if err != nil {
+			return false, err
+		}
+		if !exists {
+			continue
+		}
+		entry, err := support.Git(ctx, root, "--literal-pathspecs", "ls-tree", "-z", oid, "--", path)
+		if err != nil {
+			return false, err
 		}
+		if fields := strings.Fields(entry); len(fields) >= 2 && fields[1] == "blob" {
+			return true, nil
+		}
 	}
-	return false
+	return false, nil
 }
 
 func recoverHead(state workflow.State) (*project.Diagnostic, error) {
 	root := state.Snapshot.Root
-	repo, err := gitstore.Inspect(root)
+	head, err := support.Git(context.Background(), root, "rev-parse", "--verify", "HEAD")
 	if err != nil {
 		return nil, err
 	}
-	if repo.Head == state.PhaseBase {
+	if strings.TrimSpace(head) == state.PhaseBase {
 		return nil, nil
 	}
 	if _, err := support.Git(context.Background(), root, "reset", "--soft", state.PhaseBase); err != nil {
@@ -230,6 +257,6 @@
 	}
 	return &project.Diagnostic{
 		Code: "P05", Warning: true, Path: state.Change.Path,
-		Message: strings.Join([]string{state.Phase, "worker HEAD reset to the phase base, retaining index and working-tree work"}, " "),
+		Message: fmt.Sprintf("%s worker HEAD reset to the phase base, retaining index and working-tree work", state.Phase),
 	}, nil
 }
diff --git a/packages/runner/internal/run/run.go b/packages/runner/internal/run/run.go
index 269ee94da4b7807d505061937b14375faece9aa9..5b499b482ee8b97527fa7b3d8da3e85a4ee4f7be
--- a/packages/runner/internal/run/run.go
+++ b/packages/runner/internal/run/run.go
@@ -27,10 +27,16 @@
 	defer func() {
 		var invalid *support.CombinedFailure
 		if !errors.As(err, &invalid) && ctx.Err() != nil {
-			if parent.Err() != nil {
-				err = errors.Join(parent.Err(), err)
-			} else if context.Cause(ctx) == budget.Stop {
-				err = budget.Stop
+			if budget.Stop != nil && context.Cause(ctx) == budget.Stop {
+				if err == nil || errors.Is(err, ctx.Err()) {
+					err = budget.Stop
+				}
+			} else if parent.Err() != nil {
+				if support.IsTimeStop(err) || support.IsUsageStop(err) {
+					err = parent.Err()
+				} else if !errors.Is(err, parent.Err()) {
+					err = errors.Join(parent.Err(), err)
+				}
 			}
 		}
 		if len(blocked) > 0 {

exec
/bin/zsh -lc "cat > DRILL_NOTES.md <<'EOF'
## Contradicted Decisions

None observed.

## Gaps

- RUN-003, module runner. The shell template grammar leaves prequoted placeholders unspecified. The reconstruction shell-quotes each replacement, using the same convention as the eval hook.
- RUN-014, module workflow. Done verifies internally and accepts no prior Verify result. The reconstruction verifies before calling Done, so successful phases can run their eval hooks twice.
EOF
gofmt -w packages/runner/internal; ls -lh .fireweed/runs/drill; date -u" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 578ms:
total 0
Tue Nov 10 19:34:23 UTC 2026

exec
/bin/zsh -lc "stat -f '%SB' -t '%H:%M:%S' .fireweed/runs/drill; wc -l packages/runner/internal/*/*.go; git status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 724ms:
11:29:53
     243 packages/runner/internal/drill/drill.go
     127 packages/runner/internal/drill/report.go
     216 packages/runner/internal/drill/workspace.go
     158 packages/runner/internal/invoke/agent.go
     178 packages/runner/internal/next/audit.go
      67 packages/runner/internal/next/feedback.go
     262 packages/runner/internal/next/next.go
      72 packages/runner/internal/run/run.go
     250 packages/runner/internal/support/config.go
      95 packages/runner/internal/support/files.go
    1668 total
?? DRILL_NOTES.md
?? packages/runner/internal/

 exited 1 in 306795ms:
# github.com/lurkshark/fireweed/packages/runner/internal/next
packages/runner/internal/next/audit.go:130:28: strings.SplitAfterSeq requires go1.24 or later (module is go1.23)
?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
FAIL	github.com/lurkshark/fireweed/packages/runner/internal/next [build failed]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/support	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/runner/evals	303.939s
exit status 1

codex
The first eval run hasn’t produced its JUnit report yet. I’ve added the requested drill notes, including the verification behavior where `Done` reruns the checks that `Next` already passed.
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; go vet ./packages/runner/...; git diff --check' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 1276ms:
exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
import xml.etree.ElementTree as ET
path = Path('.fireweed/runs/drill/report.xml')
root = ET.parse(path).getroot()
cases = root.findall('testcase')
failed = [c for c in cases if c.find('failure') is not None]
print(f'{len(cases)} cases, {len(failed)} failed')
for case in failed:
 print(case.attrib['name'])
 for line in (case.findtext('failure') or '').splitlines():
  if not line.startswith('--- FAIL') and not line.startswith('=== RUN'):
   print(line[:1700])
PY" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 742ms:
163 cases, 2 failed
Test_RUN_062_CleanupFailureCannotPass
    drill_test.go:838: git [worktree remove --force /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_062_CleanupFailureCannotPass2883901273/001/.git/fireweed/worktrees/drill-alpha-3938536518]: exit status 128: fatal: '/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_062_CleanupFailureCannotPass2883901273/001/.git/fireweed/worktrees/drill-alpha-3938536518' is not a working tree
PackageFailure
fireweed: package failed without a failing test (build, setup or collection failure)

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/drill/drill.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/support/config.go
diff --git a/packages/runner/internal/drill/drill.go b/packages/runner/internal/drill/drill.go
index 753639f2c34d0724a7a0729887c2f648a3cec39e..992bfc677393577b030cab978c953e6e79a48563
--- a/packages/runner/internal/drill/drill.go
+++ b/packages/runner/internal/drill/drill.go
@@ -7,6 +7,7 @@
 	"os"
 	"path/filepath"
 	"strings"
+	"syscall"
 	"time"
 
 	"github.com/lurkshark/fireweed/packages/checks"
@@ -87,7 +88,13 @@
 		return result, err
 	}
 	defer func() {
-		_, cleanupErr := support.Git(context.Background(), repo.Root, "worktree", "remove", "--force", scratch)
+		// Git can unregister the worktree even when it cannot remove the
+		// directory. Keep the registration available for cleanup in that case.
+		const writeAndSearch = 3
+		cleanupErr := syscall.Access(parent, writeAndSearch)
+		if cleanupErr == nil {
+			_, cleanupErr = support.Git(context.Background(), repo.Root, "worktree", "remove", "--force", scratch)
+		}
 		if cleanupErr != nil {
 			result.Result = "fail"
 			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("remove verification worktree: %w", cleanupErr)}}
@@ -140,6 +147,18 @@
 	var budget support.Budget
 	var failure error
 	var retained string
+	defer func() {
+		if retained == "" {
+			return
+		}
+		ref := "refs/fireweed/drills/" + target.Name + "/" + result.Date
+		if refErr := gitstore.CASRef(repo.Root, ref, "", retained); refErr != nil {
+			result.Result = "fail"
+			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("retain drill tree: %w", refErr)}}
+			return
+		}
+		result.Ref = ref
+	}()
 	attempts := support.Attempts(config)
 	for attempt := 1; attempt <= attempts; {
 		if err := ctx.Err(); err != nil {
@@ -187,10 +206,11 @@
 		if err := importCandidate(scratch, candidate); err != nil {
 			return result, err
 		}
-		retained, err = saveCandidate(ctx, repo, scratch, target, snapshot, tracked, candidate)
+		commit, err := saveCandidate(ctx, repo, scratch, target, snapshot, tracked, candidate)
 		if err != nil {
 			return result, err
 		}
+		retained = commit
 		failure = verify(ctx, request, verifySnapshot, target, dependents, &result)
 		if ctx.Err() != nil {
 			return result, ctx.Err()
@@ -199,13 +219,6 @@
 			break
 		}
 		attempt++
-	}
-	if retained != "" {
-		ref := "refs/fireweed/drills/" + target.Name + "/" + result.Date
-		if err := gitstore.CASRef(repo.Root, ref, "", retained); err != nil {
-			return result, errors.Join(failure, err)
-		}
-		result.Ref = ref
 	}
 	if failure != nil {
 		return result, failure
diff --git a/packages/runner/internal/drill/workspace.go b/packages/runner/internal/drill/workspace.go
index 1d4d8fdee6e026ca3d5c0afa507811862ab75582..91d32c036fcc7bf1a77a53309d73e2b3fc580a31
--- a/packages/runner/internal/drill/workspace.go
+++ b/packages/runner/internal/drill/workspace.go
@@ -123,15 +123,19 @@
 }
 
 func candidateFiles(root string, baseline map[string]support.File, snapshot project.Snapshot, target project.Module) (map[string]support.File, error) {
-	for name, original := range baseline {
-		current, err := support.ReadFile(filepath.Join(root, name))
-		if err != nil || current != original {
-			return nil, fmt.Errorf("drill reconstruction modified or deleted input %s", name)
-		}
-	}
 	candidate := map[string]support.File{}
+	seen := map[string]bool{}
 	err := support.Walk(root, func(name string, entry fs.DirEntry) error {
-		if _, retained := baseline[name]; retained || entry.IsDir() || !disposable(snapshot, target, name) {
+		original, retained := baseline[name]
+		if retained {
+			current, err := support.ReadFile(filepath.Join(root, name))
+			if err != nil || current != original {
+				return fmt.Errorf("drill reconstruction modified input %s", name)
+			}
+			seen[name] = true
+			return nil
+		}
+		if entry.IsDir() || !disposable(snapshot, target, name) {
 			return nil
 		}
 		file, err := support.ReadFile(filepath.Join(root, name))
@@ -140,12 +144,39 @@
 		}
 		return err
 	})
-	return candidate, err
+	if err != nil {
+		return nil, err
+	}
+	for name := range baseline {
+		if !seen[name] {
+			return nil, fmt.Errorf("drill reconstruction deleted input %s", name)
+		}
+	}
+	return candidate, nil
 }
 
 func importCandidate(scratch string, candidate map[string]support.File) error {
 	for name, file := range candidate {
-		if err := support.WriteFile(filepath.Join(scratch, name), file); err != nil {
+		destination := filepath.Join(scratch, name)
+		for dir := filepath.Dir(destination); dir != scratch; dir = filepath.Dir(dir) {
+			info, err := os.Lstat(dir)
+			if os.IsNotExist(err) {
+				continue
+			}
+			if err != nil {
+				return err
+			}
+			if !info.IsDir() {
+				return fmt.Errorf("rebuilt path %s crosses a retained file", name)
+			}
+		}
+		if _, err := os.Lstat(destination); !os.IsNotExist(err) {
+			if err != nil {
+				return err
+			}
+			return fmt.Errorf("rebuilt path %s overlaps retained content", name)
+		}
+		if err := support.WriteFile(destination, file); err != nil {
 			return err
 		}
 	}
diff --git a/packages/runner/internal/invoke/agent.go b/packages/runner/internal/invoke/agent.go
index c769769eeba3502f3db5372fc075d65096b5b1ee..85bfd2a14c8d691b6834a9f7baa16e044157bd74
--- a/packages/runner/internal/invoke/agent.go
+++ b/packages/runner/internal/invoke/agent.go
@@ -30,11 +30,23 @@
 	}
 	values := map[string]string{"prompt": request.Prompt}
 	if containsTemplate(command, "{git_common_dir}") || containsTemplate(wrapper, "{git_common_dir}") {
-		repo, inspectErr := gitstore.Inspect(request.Snapshot.Root)
-		if inspectErr != nil {
-			return result, inspectErr
+		root := request.Snapshot.Root
+		dir, gitErr := support.Git(ctx, root, "rev-parse", "--git-common-dir")
+		if gitErr != nil {
+			return result, gitErr
+		}
+		dir = strings.TrimSuffix(dir, "\n")
+		if !filepath.IsAbs(dir) {
+			dir = filepath.Join(root, dir)
+		}
+		dir, err = filepath.Abs(dir)
+		if err != nil {
+			return result, err
+		}
+		values["git_common_dir"], err = filepath.EvalSymlinks(dir)
+		if err != nil {
+			return result, err
 		}
-		values["git_common_dir"] = repo.CommonDir
 	}
 	if containsTemplate(command, "{prompt_file}") || containsTemplate(wrapper, "{prompt_file}") {
 		f, createErr := os.CreateTemp("", "fireweed-prompt-*")
diff --git a/packages/runner/internal/next/audit.go b/packages/runner/internal/next/audit.go
index 3690e308b5e5f3e880481704291aee2e7db361c1..34b41f0442d563ee741c7135ecfffe066739a8af
--- a/packages/runner/internal/next/audit.go
+++ b/packages/runner/internal/next/audit.go
@@ -2,7 +2,6 @@
 
 import (
 	"context"
-	"errors"
 	"fmt"
 	"io/fs"
 	"os"
@@ -22,7 +21,11 @@
 	if err != nil {
 		return result, false, err
 	}
-	defer func() { err = errors.Join(err, os.RemoveAll(root)) }()
+	defer func() {
+		if cleanupErr := os.RemoveAll(root); cleanupErr != nil {
+			err = &support.CombinedFailure{Causes: []error{err, fmt.Errorf("remove audit workspace: %w", cleanupErr)}}
+		}
+	}()
 	held := request.Snapshot.Root
 	change := state.Change.Path
 	original, err := support.ReadFile(filepath.Join(held, change))
@@ -127,7 +130,7 @@
 		return 0, 0, false, true
 	}
 	pos := len(raw) - len(doc.Body)
-	for line := range strings.SplitAfterSeq(doc.Body, "\n") {
+	for _, line := range strings.SplitAfter(doc.Body, "\n") {
 		if strings.HasPrefix(line, "## ") {
 			if found && end == 0 {
 				end = pos
@@ -164,7 +167,7 @@
 			if after[:start]+after[ae:] == before {
 				return true
 			}
-			if start < 2 || after[start-1] != '\n' || after[start-2] != '\n' {
+			if start == 0 || after[start-1] != '\n' {
 				return false
 			}
 			start--
diff --git a/packages/runner/internal/next/feedback.go b/packages/runner/internal/next/feedback.go
index 6aff47fe7c774980ca589300dfc4ad339cdcc3d2..7f88faf12bd8d89a53525c3953cac7f60163f027
--- a/packages/runner/internal/next/feedback.go
+++ b/packages/runner/internal/next/feedback.go
@@ -43,19 +43,13 @@
 				continue
 			}
 			failure := c.Failure
-			for i := range maxFailure {
-				_, size := utf8.DecodeRuneInString(failure)
-				failure = failure[size:]
-				if failure == "" {
+			runes := 0
+			for offset := range failure {
+				if runes == maxFailure {
+					failure = failure[:offset] + "\n[Failure text truncated]"
 					break
 				}
-				if i == maxFailure-1 {
-					failure = c.Failure[:len(c.Failure)-len(failure)] + "\n[Failure text truncated]"
-					break
-				}
-			}
-			if failure == "" {
-				failure = c.Failure
+				runes++
 			}
 			entry := fmt.Sprintf("\nModule: %s\nClassname: %s\nTestcase: %s\n%s\n", module, c.Classname, c.Name, failure)
 			length := utf8.RuneCountInString(entry)
diff --git a/packages/runner/internal/next/next.go b/packages/runner/internal/next/next.go
index 9732acae9d5a29bd014a63e49b164a25dedc08e2..58d0621c3e9b4e4680ade5cc3c4761f501bfb73f
--- a/packages/runner/internal/next/next.go
+++ b/packages/runner/internal/next/next.go
@@ -68,8 +68,14 @@
 		if _, err := support.Usage(snapshot.Config); err != nil {
 			return result, err
 		}
-		if !request.AllowUnmerged && !approved(root, snapshot, selected.Path) {
-			return result, &approvalRequired{selected.Path}
+		if !request.AllowUnmerged {
+			ok, err := approved(ctx, root, snapshot, selected.Path)
+			if err != nil {
+				return result, err
+			}
+			if !ok {
+				return result, &approvalRequired{selected.Path}
+			}
 		}
 		if err := ctx.Err(); err != nil {
 			return result, err
@@ -86,12 +92,22 @@
 		return result, err
 	}
 	diagnostics := append([]project.Diagnostic(nil), state.Diagnostics...)
-	defer func() { result.State.Diagnostics = diagnostics }()
+	defer func() {
+		if err != nil && result.State.ClaimOID != "" {
+			if held, currentErr := workflow.Current(root, result.State.Change.Slug); currentErr == nil {
+				result.State = held
+			}
+		}
+		result.State.Diagnostics = diagnostics
+	}()
 	for result.State.Phase != "applied" {
 		if err := ctx.Err(); err != nil {
 			return result, err
 		}
 		state = result.State
+		if state.Head != state.PhaseBase {
+			return result, fmt.Errorf("P05: %s HEAD differs from the held phase base", state.Phase)
+		}
 		phase := state.Phase
 		role := "worker"
 		var prompt string
@@ -139,9 +155,6 @@
 			}
 			if recoveryErr != nil {
 				return result, errors.Join(agentErr, recoveryErr)
-			}
-			if fresh, currentErr := workflow.Current(root, state.Change.Slug); currentErr == nil {
-				result.State = fresh
 			}
 			var invalid *support.CombinedFailure
 			if errors.As(agentErr, &invalid) {
@@ -183,6 +196,9 @@
 				}
 			}
 			if attempt == attempts {
+				if err := ctx.Err(); err != nil {
+					return result, err
+				}
 				block, blockErr := workflow.Block(workflow.BlockRequest{Root: root, Change: state.Change.Slug, Reason: verifyErr.Error()})
 				if block.State.Run != "" {
 					result.State = block.State
@@ -199,7 +215,7 @@
 	return result, nil
 }
 
-func approved(root string, snapshot project.Snapshot, path string) bool {
+func approved(ctx context.Context, root string, snapshot project.Snapshot, path string) (bool, error) {
 	base := support.String(snapshot.Config["base"])
 	if base == "" {
 		base = "main"
@@ -209,20 +225,31 @@
 		remote = "origin"
 	}
 	for _, ref := range []string{"refs/heads/" + base, "refs/remotes/" + remote + "/" + base} {
-		if _, err := gitstore.ReadAt(root, ref, path); err == nil {
-			return true
+		oid, exists, err := gitstore.Ref(root, ref)
+		if err != nil {
+			return false, err
+		}
+		if !exists {
+			continue
+		}
+		entry, err := support.Git(ctx, root, "--literal-pathspecs", "ls-tree", "-z", oid, "--", path)
+		if err != nil {
+			return false, err
 		}
+		if fields := strings.Fields(entry); len(fields) >= 2 && fields[1] == "blob" {
+			return true, nil
+		}
 	}
-	return false
+	return false, nil
 }
 
 func recoverHead(state workflow.State) (*project.Diagnostic, error) {
 	root := state.Snapshot.Root
-	repo, err := gitstore.Inspect(root)
+	head, err := support.Git(context.Background(), root, "rev-parse", "--verify", "HEAD")
 	if err != nil {
 		return nil, err
 	}
-	if repo.Head == state.PhaseBase {
+	if strings.TrimSpace(head) == state.PhaseBase {
 		return nil, nil
 	}
 	if _, err := support.Git(context.Background(), root, "reset", "--soft", state.PhaseBase); err != nil {
@@ -230,6 +257,6 @@
 	}
 	return &project.Diagnostic{
 		Code: "P05", Warning: true, Path: state.Change.Path,
-		Message: strings.Join([]string{state.Phase, "worker HEAD reset to the phase base, retaining index and working-tree work"}, " "),
+		Message: fmt.Sprintf("%s worker HEAD reset to the phase base, retaining index and working-tree work", state.Phase),
 	}, nil
 }
diff --git a/packages/runner/internal/run/run.go b/packages/runner/internal/run/run.go
index 269ee94da4b7807d505061937b14375faece9aa9..5b499b482ee8b97527fa7b3d8da3e85a4ee4f7be
--- a/packages/runner/internal/run/run.go
+++ b/packages/runner/internal/run/run.go
@@ -27,10 +27,16 @@
 	defer func() {
 		var invalid *support.CombinedFailure
 		if !errors.As(err, &invalid) && ctx.Err() != nil {
-			if parent.Err() != nil {
-				err = errors.Join(parent.Err(), err)
-			} else if context.Cause(ctx) == budget.Stop {
-				err = budget.Stop
+			if budget.Stop != nil && context.Cause(ctx) == budget.Stop {
+				if err == nil || errors.Is(err, ctx.Err()) {
+					err = budget.Stop
+				}
+			} else if parent.Err() != nil {
+				if support.IsTimeStop(err) || support.IsUsageStop(err) {
+					err = parent.Err()
+				} else if !errors.Is(err, parent.Err()) {
+					err = errors.Join(parent.Err(), err)
+				}
 			}
 		}
 		if len(blocked) > 0 {
diff --git a/packages/runner/internal/support/config.go b/packages/runner/internal/support/config.go
index 3b15d4a9bee01910d62cfff683fa25ff25824c1e..f57f018f2a6161420e9af5cf1346081995228755
--- a/packages/runner/internal/support/config.go
+++ b/packages/runner/internal/support/config.go
@@ -44,7 +44,7 @@
 	if n, ok := config["attempts"].(int); ok && n > 0 {
 		return n
 	}
-	if n, ok := config["attempts"].(float64); ok && n > 0 {
+	if n, ok := config["attempts"].(float64); ok && int(n) > 0 {
 		return int(n)
 	}
 	return 2

exec
/bin/zsh -lc 'gofmt -w packages/runner/internal; export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; go test ./packages/runner/evals -run '"'"'^Test_RUN_062_CleanupFailureCannotPass$'"' -count=1" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 2769ms:
ok  	github.com/lurkshark/fireweed/packages/runner/evals	1.091s

codex
The first eval run exposed two issues. The Go 1.23 vet check caught a newer standard-library call, and the cleanup fixture showed that Git could unregister a worktree it couldn’t remove. I corrected both and am checking the cleanup case before rerunning the full hook.
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
mkdir -p .fireweed/runs/drill && sh -c '"'go run .fireweed/hooks.go eval '\"'\"'packages/runner'\"'\"' '\"'\"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/.fireweed/runs/drill/report.xml'\"'\"''" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache FIREWEED_BIN="$PWD/.fireweed/bin/fireweed"; go run .fireweed/hooks.go boundaries; go vet ./packages/runner/...; git --no-optional-locks status --short; git diff --check' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 1023ms:
?? DRILL_NOTES.md
?? packages/runner/internal/

codex
The targeted cleanup case passed. The full runner hook is running again, and I’ve started a boundary check and `go vet` against the rebuilt module.
exec
/bin/zsh -lc "date -u; stat -f '%Sm %z' -t '%H:%M:%S' .fireweed/runs/drill/report.xml; git --no-optional-locks status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 600ms:
Tue Nov 10 19:37:21 UTC 2026
11:35:00 25813
?? DRILL_NOTES.md
?? packages/runner/internal/

codex
The boundary check and `go vet` passed. The full eval hook is the remaining check.
exec
/bin/zsh -lc "date -u; stat -f '%Sm %z' -t '%H:%M:%S' .fireweed/runs/drill/report.xml; gofmt -l packages/runner/internal; git --no-optional-locks status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 692ms:
Tue Nov 10 19:41:31 UTC 2026
11:35:00 25813
?? DRILL_NOTES.md
?? packages/runner/internal/

 succeeded in 330957ms:
?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/support	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	321.091s

exec
/bin/zsh -lc "python3 - <<'PY'
import re
from pathlib import Path
from xml.etree import ElementTree as ET
suite = ET.parse('.fireweed/runs/drill/report.xml').getroot()
cases = suite.findall('testcase')
failed = sum(c.find('failure') is not None or c.find('error') is not None for c in cases)
skipped = sum(c.find('skipped') is not None for c in cases)
expected = set(re.findall(r'"'^### (RUN-'"\\d+)', Path('packages/runner/SPEC.md').read_text(), re.M))
covered = {f'RUN-{id}' for c in cases for id in re.findall(r'RUN[_-](\\d+)', c.attrib['name'])}
print(f'{len(cases) - failed - skipped} passed, {failed} failed, {skipped} skipped')
print(f'{len(covered & expected)}/{len(expected)} spec IDs named by evals')
PY
git --no-optional-locks status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 1223ms:
162 passed, 0 failed, 0 skipped
67/67 spec IDs named by evals
?? DRILL_NOTES.md
?? packages/runner/internal/

codex
Rebuilt the runner under `packages/runner/internal`. All 162 retained evals passed, with all 67 runner spec IDs represented. `go vet` and the boundary hook also passed.

The controller’s `fireweed verify` returned P02 because the drill workspace had no held change. The supplied runner hook produced the passing [JUnit report](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/.fireweed/runs/drill/report.xml).

[DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/DRILL_NOTES.md) records two gaps and no contradicted Decisions.
tokens used
175,927


### Attempt 2 (exit 0)
Corrected the runner’s approval refusal to emit `P01` while preserving `IsApprovalRequired`. Recorded the downstream CLI requirement in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/DRILL_NOTES.md).

All 162 runner evals passed, along with `fireweed check` and the boundary hook. The three app regressions still need verification in the full checkout. Retained files are intact, and no commit was made.

Reading additional input from stdin...
2026-11-10T19:51:37.519540Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1c704-1805-7902-8ddb-b0420ddd11cd
--------
user
# Rebuild runner

## Target spec and Decisions

---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
    - workflow
id_prefix: RUN
interface:
    - api.go
    - model/**
name: runner
---

# runner

## Specs

### RUN-001
InvokeAgent MUST select agent.command for the worker role and audit.command for the audit role, falling back to agent.command when audit.command is unset.

### RUN-002
InvokeAgent MUST run the selected command from Snapshot.Root.

### RUN-003
InvokeAgent MUST expand {prompt}, {prompt_file} and {git_common_dir} once in command templates and wrapper argv. Commands MUST accept shell strings or argv, quoting shell substitutions and preserving argv values literally.

### RUN-004
InvokeAgent MUST prepend the configured agent.wrap argv prefix for both worker and audit commands.

### RUN-005
InvokeAgent MUST set the command's FIREWEED_BIN environment variable to the request Bin.

### RUN-006
InvokeAgent MUST retain the command's stdout, stderr and exit code in AgentResult.

### RUN-007
InvokeAgent MUST use a positive request Timeout in preference to the configured timeout, using the configured timeout when the request is unset. It MUST stop the command's process group when that bound expires or its context is cancelled.

### RUN-008
InvokeAgent MUST set UsageLimit true exactly when a nonzero command exit has a configured limit_patterns match in its final 20 output lines.

### RUN-009
InvokeAgent MUST return command-launch, timeout or cancellation errors without discarding captured output.

### RUN-010
Next MUST return Empty true without invoking agents when workflow.Eligible
returns no changes, including an empty or entirely waiting queue.

### RUN-011
When workflow.Eligible returns changes, Next MUST select its first result and
start it through workflow.Start, propagating start failures without invoking
an agent.

### RUN-012
Unless AllowUnmerged is true, Next MUST reject a selected change absent from
both the local and shared base before claiming it or assigning IDs.

### RUN-013
For each phase, Next MUST invoke the worker role for evals and impl from the held checkout and the audit role from its restricted input workspace, using approved agent configuration and request Bin and Timeout.
Worker prompts MUST include full workflow context.
Audit prompts MUST include the change, projected spec, eval contents and audit playbook without implementation contents.

### RUN-014
After each successful agent invocation, Next MUST verify the current phase and
call workflow.Done only when verification passes, continuing until applied.
It MUST return the applied state and the ordered phase invocation results.

### RUN-015
For a non-usage-limit agent invocation failure, Next MUST return an error and
captured invocation results without advancing the phase or releasing its held
claim.

### RUN-016
Run without a positive Max MUST continue single-change execution until no
eligible changes remain or its For bound stops it, returning applied NextResults
in execution order.

### RUN-017
With a positive Max, Run MUST apply at most Max changes per invocation. A later
Run invocation MUST continue the remaining queue without reapplying archived
changes.

### RUN-018
On successful return, Run MUST set Empty when no queued changes remain and set
Limited exactly when a positive Max stopped it with queued work remaining.

### RUN-019
Run MUST supply its Root, Bin, AllowUnmerged and Timeout to single-change
execution for every change it processes.

### RUN-020
If single-change execution fails while the calling worktree still holds the
change, Run MUST return the error with its cause and already-applied results
retained, leaving later changes unstarted.

### RUN-021
Next MUST permit at most the configured attempts count of agent invocations
for a phase after verification failures, defaulting to two when unset. Each
phase MUST have its own attempt allowance.

### RUN-022
When phase verification fails and an attempt remains, Next MUST retain the
earlier work and retry that phase. The retry prompt MUST include the last
verification output.

### RUN-023
For exhausted phase verification with the worktree's approved claim and an
unchanged HEAD, Next MUST save and block the step through the lifecycle,
leaving a clean unclaimed checkout and a state-only block commit.

### RUN-024
After that blocking succeeds, Next MUST return the blocked lifecycle State and
ordered captured invocations without changing the run, assigned IDs or current
phase. Its error MUST retain the last verification cause and identify the saved
attempt ref.

### RUN-025
When single-change execution returns an error after successfully saving and
blocking its change with a clean unclaimed checkout, Run MUST continue processing
eligible changes. Run MUST exclude that blocked change from Applied.

### RUN-026
If a change was blocked during Run, Run MUST return an error when eligible work
is exhausted or Max is reached. That error MUST retain a blocked change's
verification cause and saved attempt ref.

### RUN-027
When an evals or impl worker creates commits, Next MUST soft-reset the held checkout's HEAD to its phase base after invocation and before verification or returning an invocation failure, retaining the step's working-tree and index work.

### RUN-028
For each such recovery, Next MUST return a P05 warning in State.Diagnostics
identifying the phase and that work was retained. It MUST preserve these warnings
across phase advancement and on an invocation-failure return.

### RUN-029
Next MUST read limit_wait and limit_budget as durations, defaulting to 30m and
4h when unset. It MUST reject malformed durations, a nonpositive wait or negative
budget with C01 before claiming a fresh change or invoking an agent.

### RUN-030
Subject to the usage budget and Run's For bound, an AgentResult with UsageLimit
true MUST make Next wait limit_wait and repeat the same phase attempt without
spending a failed attempt. It MUST retain earlier step work and captured
invocations, keeping existing verification feedback.

### RUN-031
When another usage wait would exceed limit_budget, Next MUST return a clean
usage-stop error identifying completed wait, next wait and budget. It MUST retain
the held State, work and ordered Steps without advancing, blocking, committing
or releasing the claim.

### RUN-032
Run MUST account for total usage wait across phases and changes in one
invocation. On a clean usage stop it MUST retain already-applied results and
propagate the stop classification.

### RUN-033
When the calling worktree already holds a queued step, Next MUST resume that
owned State using its approved snapshot without a fresh Start, retaining its
run, phase, assigned IDs and work. Existing approval and unapproved-work bounds
MUST remain enforced for starting another change.

### RUN-034
IsUsageStop MUST recognize clean usage-stop errors through wrapping and return
false for nil or ordinary failures.

### RUN-035
If Run blocked a change before a usage stop, it MUST return failure with the
blocking cause/ref and stop reason retained. IsUsageStop MUST return false for
that combined failure so the caller cannot report it as a clean stop.

### RUN-036
Run MUST treat For zero as unbounded and positive For as one elapsed-time
budget for the whole invocation, including agent and verification time across
phases and changes.

### RUN-037
When its For bound stops execution, Run MUST return a nonnil clean time-stop
error naming the limit, with already-applied results retained. A step interrupted
during agent invocation or verification MUST keep its work, claim, run, phase
and assigned IDs without advancing, blocking, committing or releasing that step.

### RUN-038
When a usage wait fits its usage budget but would pass Run's For deadline,
execution MUST return a clean time stop before waiting. It MUST preserve the
held step's work and identity.

### RUN-039
IsTimeStop MUST recognize clean time-stop errors through wrapping. It MUST
return false for nil, usage stops, ordinary failures and parent-context
cancellation.

### RUN-040
If Run blocked a change before a time stop, it MUST return ordinary failure
retaining the blocking cause/ref and time-stop reason. IsTimeStop MUST return
false for that combined failure.

### RUN-041
Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.

### RUN-042
IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.

### RUN-043
On a verification retry, Next MUST append a detail section identifying the module, classname, testcase name and Case.Failure for failed cases from every reported module, including dependents, subject to RUN-044.
It MUST select cases by lexical module name and then report order, excluding passing and skipped cases.

### RUN-044
The failed-case detail section MUST fit within 16,384 Unicode code points, including labels and truncation markers.
Each Case.Failure excerpt MUST contain at most its first 4,096 code points.
Shortened diagnostic text or omitted cases MUST be marked.

### RUN-045
Next MUST run every automatic audit invocation in a fresh temporary workspace containing only the current change, projected module SPEC.md and module eval files at project-relative paths, plus independent Git metadata.
It MUST omit the held checkout's other files, including implementation, interfaces, hooks and configuration.

### RUN-046
The audit workspace's Git metadata MUST contain none of the held repository's history, refs or saved attempts.
Audit-agent commits within this workspace MUST remain absent from the held checkout's history and refs.

### RUN-047
Next MUST accept only edits to the Audit section of the copied current change.
Edits elsewhere in that change, edits to copied specs or evals, or added or deleted files outside Git metadata MUST reject the submission without copying any workspace edits to the held checkout.

### RUN-048
When audit workspace edits are permitted, Next MUST copy its Audit section back before retrying or returning, including agent failures and time or usage stops.
Copying audit work MUST not itself advance a phase or replace other held change contents.

### RUN-049
Next MUST remove every temporary audit workspace after invocation, including after invalid submissions, agent failures and time or usage stops.

### RUN-050
Drill MUST reject an unknown module, dirty invoking checkout or any live shared change claim before setup or agent invocation, leaving the invoking checkout and refs unchanged.

### RUN-051
Drill MUST create a detached scratch worktree at the invoking HEAD under the common Git directory's fireweed/worktrees directory and run the configured setup hook there with FIREWEED_BIN set to request Bin. Setup failure MUST prevent agent invocation and return its captured failure.

### RUN-052
Drill MUST delete only target-module files classified disposable by the approved zone precedence before reconstruction. It MUST retain every other target-module file and every other module's files in the verification worktree.

### RUN-053
Drill MUST invoke reconstruction in a separate input workspace holding the invoking HEAD's tracked files except the target's disposable files, every direct and transitive dependent module's files, the changes directory and drill reports. Its Git metadata MUST hold one fresh commit and none of the invoking repository's history, refs or saved attempts. Drill MUST run the configured setup hook there before agent invocation.

### RUN-054
The reconstruction prompt MUST include the target spec and Decisions, identify its retained interfaces and evals, and require rebuilding only disposable paths without reading outside the input workspace. It MUST give a command that runs the target's eval hook in that workspace. It MUST request separate Contradicted Decisions and Gaps sections in DRILL_NOTES.md, each gap naming the nearest spec ID or "no spec" and a target module.

### RUN-055
Drill MUST reject an attempt that modified or deleted any input file present after setup, before importing code. It MUST import only the target's rebuilt disposable files into the verification worktree, ignoring other files the worker or its eval runs created, without changing retained files or the invoking implementation.

### RUN-056
Drill MUST use the configured worker command, wrapper and timeout through InvokeAgent, retaining ordered AgentResults in Steps. Non-usage failures and verification retries MUST share the configured attempts bound, retain rebuilt work and include the preceding failure in retry prompts.

### RUN-057
Drill MUST treat usage-limit responses as the same attempt, applying configured limit_wait and limit_budget without spending a failed attempt. An exhausted usage budget MUST return an IsUsageStop error with Result stopped; cancellation or timeout MUST return an error and stop reconstruction.

### RUN-058
After accepted reconstruction, Drill MUST run the target module's eval hook and require passing enforced cases and coverage through checks. It MUST retain the complete returned own report in Own, including failed testcase details when verification fails.

### RUN-059
Drill MUST evaluate every direct and transitive dependent in lexical module-name order regardless of verify_dependents, retaining each report in Dependents. Every dependent MUST satisfy passing enforced cases and coverage for a passing drill, even when own or another dependent's evals fail.

### RUN-060

After reconstruction has been invoked, Drill MUST write the current invocation's report to .fireweed/drills/<module>.md, replacing any previous report there, with date, base_commit, agent, model and result frontmatter and Own evals, Dependent evals, Contradicted Decisions and Gaps sections. Eval entries MUST identify module, testcase, spec IDs and result. Agent observations and verification failures MUST remain visible, with each gap identifying a nearest spec ID or "no spec" and target module.

### RUN-061
After accepted reconstruction, Drill MUST retain the rebuilt verification tree at refs/fireweed/drills/<module>/<date>, returning that ref without moving the invoking HEAD. A passing result MUST require accepted reconstruction and passing own and all dependent evals; failure MUST return a nonnil error with Result fail.

### RUN-062
Drill MUST remove its verification worktree and input workspace on success, agent or verification failure, invalid edits, usage stop and cancellation. Cleanup failure MUST be returned as an error rather than reported as a passing drill.

### RUN-063
InvokeAgent MUST persist the invocation's complete stdout and stderr, including partial output on command failure, timeout or cancellation, at .fireweed/runs/<run>/<phase>-<attempt>.log under its log project root. Repeated invocations at the same path MUST append without discarding earlier records. AgentResult.LogPath MUST identify that file by an absolute path.

### RUN-064
Next MUST assign each invocation the held change's run, current phase and one-based phase attempt, writing its log under the invoking project's .fireweed/runs directory. Audit logs MUST stay in that invoking project, outside the restricted audit input workspace.

### RUN-065
Drill MUST allocate one log run identity per invocation and reuse it across reconstruction retries and usage waits. Its logs MUST use phase drill, the current one-based reconstruction attempt and the invoking project's root, outside scratch and reconstruction input workspaces.

### RUN-066
InvokeAgent MUST accept a log project root, run, phase and positive one-based attempt in AgentRequest. Unset fields MUST default to Snapshot.Root, a fresh run identifier, request Role and attempt 1, respectively.

### RUN-067
InvokeAgent MUST create one temporary file containing the exact request Prompt only when {prompt_file} occurs in the selected command or wrapper templates, and expand that placeholder to its absolute path. The file MUST remain available through command execution and be removed afterward, including after launch failure, nonzero exit, timeout or cancellation.

## Decisions

[0075-bounded-agent-invocation] An unset configured timeout defaults to 90 minutes. limit_patterns are
case-insensitive literal substrings, using the final 20 lines of stdout followed
by stderr with terminal newline separators trimmed.

- agent-invocation-logs: Agent logs append length-labeled stdout and stderr records after each invocation, preserving both streams' bytes even without final newlines. A log write error is returned alongside any agent error while the captured result remains available.


## Retained interfaces and files

packages/runner/SPEC.md
packages/runner/api.go
packages/runner/model/model.go

## Retained evals

packages/runner/evals/.gitkeep
packages/runner/evals/agent_logs_test.go
packages/runner/evals/approval_test.go
packages/runner/evals/audit_workspace_test.go
packages/runner/evals/commit_recovery_test.go
packages/runner/evals/drill_test.go
packages/runner/evals/invoke_test.go
packages/runner/evals/next_test.go
packages/runner/evals/prompt_file_test.go
packages/runner/evals/retry_failure_details_test.go
packages/runner/evals/run_test.go
packages/runner/evals/time_limit_test.go
packages/runner/evals/usage_test.go

Rebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.

Run the target's eval hook from the workspace root as often as needed.

~~~sh
mkdir -p .fireweed/runs/drill && sh -c 'go run .fireweed/hooks.go eval '"'"'packages/runner'"'"' '"'"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/.fireweed/runs/drill/report.xml'"'"''
~~~

Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module.


# Previous verification failed
app eval: E01: failed case "Test_APP_075_NextAllowUnmergedAndUsage" for APP-075
runner stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/support	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	297.856s

runner stderr: 

app stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	270.913s

app stderr: exit status 1

app github.com/lurkshark/fireweed/packages/app/evals Test_APP_075_NextAllowUnmergedAndUsage: === RUN   Test_APP_075_NextAllowUnmergedAndUsage
    next_cli_test.go:236: plain diagnostic stdout="" stderr="P04: P04: change changes/0001-next.md requires approval on the local or shared base. Approve the change on the configured base or use --allow-unmerged for authorized local work, then retry next.\n"; want P01 and action
--- FAIL: Test_APP_075_NextAllowUnmergedAndUsage (4.91s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_077_NextExitClasses: === RUN   Test_APP_077_NextExitClasses
    next_cli_test.go:273: plain diagnostic stdout="" stderr="P04: P04: change changes/0001-next.md requires approval on the local or shared base. Approve the change on the configured base or use --allow-unmerged for authorized local work, then retry next.\n"; want P01 and action
--- FAIL: Test_APP_077_NextExitClasses (1.05s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_082_RunAllowUnmerged: === RUN   Test_APP_082_RunAllowUnmerged
    run_cli_test.go:366: plain diagnostic stdout="" stderr="P04: P04: change changes/0001-alpha.md requires approval on the local or shared base. Approve the change on the configured base or use --allow-unmerged for authorized local work, then retry run.\n"; want P01 and action
--- FAIL: Test_APP_082_RunAllowUnmerged (4.97s)

codex
I’ll read the rebuild guardrails and retained runner contracts, then check the approval diagnostic reported by the previous verification.
exec
/bin/zsh -lc "pwd; rg --files -g 'AGENTS.md' -g 'SYSTEM.md' -g 'PROGRESS.md' -g 'DECISIONS.md' -g 'go.mod' -g 'DRILL_NOTES.md'; ls -la .fireweed; git status --short" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc 'rg --files packages/runner; ls packages' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 542ms:
packages/runner/SPEC.md
packages/runner/api.go
packages/runner/model/model.go
packages/runner/internal/run/run.go
packages/runner/internal/support/files.go
packages/runner/internal/support/config.go
packages/runner/internal/invoke/agent.go
packages/runner/internal/next/next.go
packages/runner/internal/next/feedback.go
packages/runner/internal/next/audit.go
packages/runner/internal/drill/workspace.go
packages/runner/internal/drill/report.go
packages/runner/internal/drill/drill.go
packages/runner/evals/next_test.go
packages/runner/evals/usage_test.go
packages/runner/evals/approval_test.go
packages/runner/evals/prompt_file_test.go
packages/runner/evals/time_limit_test.go
packages/runner/evals/invoke_test.go
packages/runner/evals/retry_failure_details_test.go
packages/runner/evals/agent_logs_test.go
packages/runner/evals/run_test.go
packages/runner/evals/drill_test.go
packages/runner/evals/commit_recovery_test.go
packages/runner/evals/audit_workspace_test.go
assets
checks
contracts
gitstore
project
runner
workflow

 succeeded in 595ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
AGENTS.md
SYSTEM.md
DRILL_NOTES.md
go.mod
docs/rebuild/DECISIONS.md
docs/rebuild/PROGRESS.md
total 24
drwxr-xr-x@  9 joel  staff   288 Nov 10 11:29 .
drwx------@ 18 joel  staff   576 Nov 10 11:34 ..
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:24 bin
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:24 canaries
-rw-r--r--@  1 joel  staff  3993 Nov 10 11:24 hooks.go
-rw-r--r--@  1 joel  staff  1502 Nov 10 11:24 recipe.md
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:29 runs
-rw-r--r--@  1 joel  staff   470 Nov 10 11:24 setup.sh
drwxr-xr-x@  3 joel  staff    96 Nov 10 11:24 templates
?? DRILL_NOTES.md
?? packages/runner/internal/

exec
/bin/zsh -lc "rg -n 'P01|P04|Approval' packages; cat DRILL_NOTES.md; ls -l .fireweed/bin; ls -d changes 2>/dev/null" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "cat SYSTEM.md; cat docs/rebuild/PROGRESS.md; rg -n 'P01|P04|approval|agent-invocation|runner' docs/rebuild/DECISIONS.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 525ms:
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
`runner` and `app` need first native reports and retained refs.

## Next step

Run the owner-approved `runner` drill with the verified controller, commit
the native report, retain the proof ref, and stop for owner review before
`app`. The remaining drill order is `runner`, then `app` in the
[regeneration goal](GOAL.md).

Deadline set by the owner: **2026-11-10 22:51:05 UTC**.
22:| P01–P99, P2-01 | Lifecycle and approval. P01 preconditions; P02 claims; P03 approved state; P04 approval warning; P05 stale base; P06 blocked/Questions; P07 unapproved-work limit; P08 repeated-module warning; P09 audit submission. Existing P2-01 remains missing genuine fail-first evidence. |

 exited 1 in 533ms:
packages/assets/SPEC.md:91:Explain MUST describe the cause and a corrective action for P01 (operational preconditions, including lifecycle, new-change and init), P02 (claims), P03 (unapproved state alteration), P04 (approval warning), P05 (stale bases), P06 (blocked change), P07 (unapproved-work bound), P09 (audit submissions) and Z01 (forbidden phase-zone edits).
packages/assets/internal/content/content.go:106:	"P01":   "A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.",
packages/assets/internal/content/content.go:109:	"P04":   "Warning: approval for this change is not recorded on the base branch. Obtain approval through review before unattended execution.",
packages/assets/evals/assets_test.go:317:		{"P01", `(?i)(precondition|prerequisite)[^.\n]*(missing|unmet|fail)|(missing|unmet|fail)[^.\n]*(precondition|prerequisite)`, `(?i)(complete|satisfy|fix|restore)[^.\n]*(precondition|prerequisite|required|state)`},
packages/assets/evals/assets_test.go:320:		{"P04", `(?i)warn[^.\n]*approv|approv[^.\n]*warn`, `(?i)(review|obtain|record|request)[^.\n]*approv`},
packages/assets/evals/assets_test.go:335:			if tc.code == "P01" {
packages/runner/SPEC.md:201:Next MUST return an error for which IsApprovalRequired reports true when the selected change file is on neither the local nor the shared base and AllowUnmerged is false.
packages/runner/SPEC.md:204:IsApprovalRequired MUST report false for Next errors caused by anything other than a missing approval.
packages/runner/api.go:43:func IsApprovalRequired(err error) bool {
packages/runner/api.go:44:	return next.IsApprovalRequired(err)
packages/runner/evals/next_test.go:325:func Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents(t *testing.T) {
packages/runner/evals/next_test.go:345:		if err == nil || !strings.Contains(err.Error(), "P01") {
packages/workflow/SPEC.md:46:Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.
packages/workflow/SPEC.md:54:change. When Eligible returns no change, Start MUST reject with P01.
packages/workflow/SPEC.md:194:Start MUST prefix dirty-worktree and queue-eligibility rejections with P01.
packages/workflow/SPEC.md:197:Verify and Done MUST prefix unsupported-phase rejections with P01.
packages/workflow/SPEC.md:294:attempt ref with P01 before claiming or writing the change.
packages/workflow/SPEC.md:356:Abandon MUST reject a dirty working tree with P01 before changing files, HEAD or refs.
packages/workflow/SPEC.md:364:Abandon MUST reject an applied archived change with P01.
packages/workflow/SPEC.md:480:An empty or unresolvable range MUST return an error with P01 and the rejected range or resolution cause.
packages/workflow/internal/lifecycle/resume_stub.go:75:		return model.State{}, errors.New("P01: attempt number must be positive")
packages/workflow/internal/lifecycle/resume_stub.go:87:		return model.State{}, errors.New("P01: saved attempt requires a committed phase")
packages/workflow/internal/lifecycle/resume_stub.go:95:		return model.State{}, fmt.Errorf("P01: saved attempt %s does not exist", name)
packages/runner/evals/approval_test.go:10:func Test_RUN_041_UnapprovedChangeRequiresApproval(t *testing.T) {
packages/runner/evals/approval_test.go:16:	if !runner.IsApprovalRequired(err) {
packages/runner/evals/approval_test.go:17:		t.Errorf("IsApprovalRequired(%v) = false", err)
packages/runner/evals/approval_test.go:22:func Test_RUN_042_OtherNextErrorsDoNotRequireApproval(t *testing.T) {
packages/runner/evals/approval_test.go:23:	if runner.IsApprovalRequired(nil) {
packages/runner/evals/approval_test.go:26:	if runner.IsApprovalRequired(errors.New("requires approval")) {
packages/runner/evals/approval_test.go:37:		if runner.IsApprovalRequired(err) {
packages/runner/evals/approval_test.go:49:		if runner.IsApprovalRequired(err) {
packages/workflow/internal/lifecycle/stub.go:306:		return model.State{}, errors.New("P01: start requires clean worktree")
packages/workflow/internal/lifecycle/stub.go:321:		return model.State{}, errors.New("P01: no eligible change")
packages/workflow/internal/lifecycle/stub.go:325:		return model.State{}, errors.New("P01: change is not eligible queue head")
packages/workflow/internal/lifecycle/stub.go:330:		ds = append(ds, diag("P04", "change is not approved on local or shared base", true))
packages/workflow/internal/lifecycle/stub.go:627:		return r, fmt.Errorf("P01: unsupported phase %s", c.Phase)
packages/workflow/internal/lifecycle/stub.go:1030:		return result, fmt.Errorf("P01: unsupported phase %s", c.Phase)
packages/workflow/internal/lifecycle/block_stub.go:121:		return result, fmt.Errorf("P01: unsupported phase %s", c.Phase)
packages/workflow/internal/lifecycle/pr_body_stub.go:31:		return model.PRBodyReport{}, fmt.Errorf("P01: invalid range %q", revisionRange)
packages/workflow/internal/lifecycle/pr_body_stub.go:37:			return model.PRBodyReport{}, fmt.Errorf("P01: resolve range %q: %w", revisionRange, err)
packages/workflow/internal/lifecycle/abandon_stub.go:25:		return "", errors.New("P01: abandon requires clean worktree")
packages/workflow/internal/lifecycle/abandon_stub.go:43:			return "", errors.New("P01: applied archived change cannot be abandoned")
packages/workflow/internal/lifecycle/abandon_stub.go:75:			return "", fmt.Errorf("P01: phase commit %s has no parent", oid)
packages/workflow/evals/abandon_test.go:142:			assertAbandonRefusal(t, root, "manual", "P01")
packages/workflow/evals/abandon_test.go:196:	assertAbandonRefusal(t, root, "manual", "P01")
packages/workflow/evals/pr_body_test.go:339:			if err == nil || !strings.Contains(err.Error(), "P01") || (invalid != "" && !strings.Contains(err.Error(), "not-a-commit") && !strings.Contains(err.Error(), "not-a-range") && !strings.Contains(err.Error(), "resolve")) {
packages/workflow/evals/pr_body_test.go:340:				t.Errorf("PRBody(%q) error = %v, want P01 with rejected range or cause", invalid, err)
packages/runner/internal/next/next.go:20:	return fmt.Sprintf("P04: change %s requires approval on the local or shared base", e.path)
packages/runner/internal/next/next.go:23:func IsApprovalRequired(err error) bool {
packages/runner/internal/next/next.go:123:			return result, fmt.Errorf("P01: unsupported phase %s", phase)
packages/workflow/evals/lifecycle_test.go:331:func Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:333:	if s := start(t, approved); containsCode(s.Diagnostics, "P04") {
packages/workflow/evals/lifecycle_test.go:344:	if s := start(t, shared); containsCode(s.Diagnostics, "P04") {
packages/workflow/evals/lifecycle_test.go:377:				if err != nil || !containsCode(s.Diagnostics, "P04") {
packages/workflow/evals/lifecycle_test.go:381:					if d.Code == "P04" && !d.Warning {
packages/workflow/evals/lifecycle_test.go:382:						t.Errorf("P04 diagnostic is not a warning: %#v", d)
packages/workflow/evals/lifecycle_test.go:394:	if _, err := verify(root, true); err == nil || !strings.HasPrefix(err.Error(), "P01: unsupported phase review") {
packages/workflow/evals/lifecycle_test.go:397:	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P01: unsupported phase review") || head(t, root) != before {
packages/workflow/evals/lifecycle_test.go:417:		if _, err := operation.call(); err == nil || !strings.HasPrefix(err.Error(), "P01: unsupported phase review") {
packages/workflow/evals/lifecycle_test.go:418:			t.Errorf("%s with protected edit = %v, want P01 unsupported phase", operation.name, err)
packages/workflow/evals/lifecycle_test.go:1405:func Test_FLOW_054_StartOperationalPreconditionsHaveP01(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1408:	if _, err := workflow.Start(dirty, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P01: start requires clean worktree") {
packages/workflow/evals/lifecycle_test.go:1414:	if _, err := workflow.Start(queued, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P01: change is not eligible queue head") {
packages/workflow/evals/eligible_test.go:179:			if _, err := workflow.Start(root, target); err == nil || !strings.HasPrefix(err.Error(), "P01") {
packages/workflow/evals/eligible_test.go:180:				t.Errorf("Start(%s) = %v, want P01 rejection", target, err)
packages/workflow/evals/eligible_test.go:188:		if _, err := workflow.Start(root, ""); err == nil || !strings.HasPrefix(err.Error(), "P01") {
packages/workflow/evals/eligible_test.go:189:			t.Errorf("Start(empty) = %v, want P01 rejection", err)
packages/workflow/evals/resume_test.go:87:func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptRestoresEvalsStepAndRetainsCurrentApproval(t *testing.T) {
packages/workflow/evals/resume_test.go:249:func Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval(t *testing.T) {
packages/workflow/evals/resume_test.go:312:		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
packages/workflow/evals/resume_test.go:313:			t.Errorf("dirty resume = %v, want P01", err)
packages/workflow/evals/resume_test.go:322:		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
packages/workflow/evals/resume_test.go:323:			t.Errorf("blocked resume = %v, want P01", err)
packages/workflow/evals/resume_test.go:333:		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
packages/workflow/evals/resume_test.go:334:			t.Errorf("waiting resume = %v, want P01", err)
packages/workflow/evals/resume_test.go:362:		if _, err := workflow.StartFromAttempt(root, "manual", attempt); err == nil || !strings.HasPrefix(err.Error(), "P01") {
packages/workflow/evals/resume_test.go:363:			t.Errorf("attempt %d = %v, want P01", attempt, err)
## Contradicted Decisions

None observed.

## Gaps

- RUN-003, module runner. The shell template grammar leaves prequoted placeholders unspecified. The reconstruction shell-quotes each replacement, using the same convention as the eval hook.
- RUN-014, module workflow. Done verifies internally and accepts no prior Verify result. The reconstruction verifies before calling Done, so successful phases can run their eval hooks twice.
total 13512
-rwxr-xr-x@ 1 joel  staff  6914850 Nov 10 11:24 fireweed

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed verify --help; fireweed context --help; sed -n '"'1,170p' packages/runner/internal/next/next.go; cat packages/runner/evals/approval_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "cat .fireweed/hooks.go; cat .fireweed/setup.sh; cat docs/rebuild/DECISIONS.md; git ls-files 'packages/runner/**' | head -25; ls -a" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 541ms:
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
packages/runner/SPEC.md
packages/runner/api.go
packages/runner/evals/.gitkeep
packages/runner/evals/agent_logs_test.go
packages/runner/evals/approval_test.go
packages/runner/evals/audit_workspace_test.go
packages/runner/evals/commit_recovery_test.go
packages/runner/evals/drill_test.go
packages/runner/evals/invoke_test.go
packages/runner/evals/next_test.go
packages/runner/evals/prompt_file_test.go
packages/runner/evals/retry_failure_details_test.go
packages/runner/evals/run_test.go
packages/runner/evals/time_limit_test.go
packages/runner/evals/usage_test.go
packages/runner/model/model.go
.
..
.fireweed
.git
.gitignore
AGENTS.md
DRILL_NOTES.md
LICENSE
Makefile
README.md
SYSTEM.md
docs
fireweed-design-v1.md
fireweed.yaml
go.mod
go.sum
main.go
packages

 succeeded in 544ms:
U01: Unknown option "--help". Use help to see available commands and options.
U01: Unknown option "--help". Use help to see available commands and options.
package next

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/invoke"
	"github.com/lurkshark/fireweed/packages/runner/internal/support"
	"github.com/lurkshark/fireweed/packages/runner/model"
	"github.com/lurkshark/fireweed/packages/workflow"
)

type approvalRequired struct{ path string }

func (e *approvalRequired) Error() string {
	return fmt.Sprintf("P04: change %s requires approval on the local or shared base", e.path)
}

func IsApprovalRequired(err error) bool {
	var approval *approvalRequired
	return errors.As(err, &approval)
}

func IsUsageStop(err error) bool { return support.IsUsageStop(err) }

type Blocked struct {
	Ref   string
	Cause error
}

func (e *Blocked) Error() string {
	return fmt.Sprintf("verification exhausted, saved attempt %s: %v", e.Ref, e.Cause)
}
func (e *Blocked) Unwrap() error { return e.Cause }

func Run(ctx context.Context, request model.NextRequest) (model.NextResult, error) {
	return Execute(ctx, request, &support.Budget{})
}

func Execute(ctx context.Context, request model.NextRequest, budget *support.Budget) (result model.NextResult, err error) {
	if err := ctx.Err(); err != nil {
		return result, err
	}
	repo, err := gitstore.Inspect(request.Root)
	if err != nil {
		return result, err
	}
	root := repo.Root
	state, err := workflow.Current(root, "")
	if err != nil {
		queue, err := workflow.Eligible(root)
		if err != nil {
			return result, err
		}
		if len(queue) == 0 {
			result.Empty = true
			return result, nil
		}
		selected := queue[0]
		snapshot, err := project.Load(root)
		if err != nil {
			return result, err
		}
		if _, err := support.Usage(snapshot.Config); err != nil {
			return result, err
		}
		if !request.AllowUnmerged {
			ok, err := approved(ctx, root, snapshot, selected.Path)
			if err != nil {
				return result, err
			}
			if !ok {
				return result, &approvalRequired{selected.Path}
			}
		}
		if err := ctx.Err(); err != nil {
			return result, err
		}
		result.State, err = workflow.Start(root, selected.Slug)
		if err != nil {
			return result, err
		}
		state = result.State
	}
	result.State = state
	usage, err := support.Usage(state.Snapshot.Config)
	if err != nil {
		return result, err
	}
	diagnostics := append([]project.Diagnostic(nil), state.Diagnostics...)
	defer func() {
		if err != nil && result.State.ClaimOID != "" {
			if held, currentErr := workflow.Current(root, result.State.Change.Slug); currentErr == nil {
				result.State = held
			}
		}
		result.State.Diagnostics = diagnostics
	}()
	for result.State.Phase != "applied" {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		state = result.State
		if state.Head != state.PhaseBase {
			return result, fmt.Errorf("P05: %s HEAD differs from the held phase base", state.Phase)
		}
		phase := state.Phase
		role := "worker"
		var prompt string
		switch phase {
		case "audit":
			role = "audit"
		case "evals", "impl":
			prompt, err = workflow.Context(root, state.Change.Slug, true)
			if err != nil {
				return result, err
			}
		default:
			return result, fmt.Errorf("P01: unsupported phase %s", phase)
		}
		attempts := support.Attempts(state.Snapshot.Config)
		check := workflow.Request{
			Root: root, Change: state.Change.Slug, Bin: request.Bin, Timeout: request.Timeout,
			Agent: support.AgentName(support.AgentCommand(state.Snapshot.Config, role)),
		}
		var feedback string
		for attempt := 1; attempt <= attempts; {
			if err := ctx.Err(); err != nil {
				return result, err
			}
			agentReq := model.AgentRequest{
				Snapshot: state.Snapshot, Role: role, Prompt: prompt + feedback,
				Bin: request.Bin, Timeout: request.Timeout,
				LogRoot: root, Run: state.Run, Phase: phase, Attempt: attempt,
			}
			var agent model.AgentResult
			var agentErr, recoveryErr error
			invoked := true
			if role == "audit" {
				agent, invoked, agentErr = audit(ctx, state, agentReq, feedback)
			} else {
				agent, agentErr = invoke.Agent(ctx, agentReq)
				var warning *project.Diagnostic
				warning, recoveryErr = recoverHead(state)
				if warning != nil {
					diagnostics = append(diagnostics, *warning)
				}
			}
			if invoked {
				result.Steps = append(result.Steps, model.StepResult{Phase: phase, Role: role, Agent: agent})
			}
			if recoveryErr != nil {
				return result, errors.Join(agentErr, recoveryErr)
			}
			var invalid *support.CombinedFailure
			if errors.As(agentErr, &invalid) {
				return result, agentErr
			}
			if ctx.Err() != nil {
				return result, errors.Join(ctx.Err(), agentErr)
			}
			if agent.UsageLimit && !errors.Is(agentErr, context.DeadlineExceeded) && !errors.Is(agentErr, context.Canceled) {
				if err := budget.Wait(ctx, usage); err != nil {
					return result, err
				}
				continue
package evals_test

import (
	"errors"
	"testing"

	"github.com/lurkshark/fireweed/packages/runner"
)

func Test_RUN_041_UnapprovedChangeRequiresApproval(t *testing.T) {
	f := newNextFixture(t, true, false, "normal")
	_, err := f.run(false)
	if err == nil {
		t.Fatal("Next accepted a change absent from the local and shared base")
	}
	if !runner.IsApprovalRequired(err) {
		t.Errorf("IsApprovalRequired(%v) = false", err)
	}
	f.noAgent(t)
}

func Test_RUN_042_OtherNextErrorsDoNotRequireApproval(t *testing.T) {
	if runner.IsApprovalRequired(nil) {
		t.Error("nil error requires approval")
	}
	if runner.IsApprovalRequired(errors.New("requires approval")) {
		t.Error("unrelated error with approval wording requires approval")
	}

	t.Run("start_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextWrite(t, f.root, "stray.txt", "dirty\n")
		_, err := f.run(false)
		if err == nil {
			t.Fatal("expected dirty-worktree start failure")
		}
		if runner.IsApprovalRequired(err) {
			t.Errorf("start failure requires approval: %v", err)
		}
		f.noAgent(t)
	})

	t.Run("agent_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exit")
		result, err := f.run(false)
		if err == nil || len(result.Steps) != 1 || result.Steps[0].Agent.ExitCode != 9 {
			t.Fatalf("expected agent exit 9, got %#v, %v", result, err)
		}
		if runner.IsApprovalRequired(err) {
			t.Errorf("agent failure requires approval: %v", err)
		}
	})
}

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed help; fireweed context runner; fireweed verify --impl' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "cat fireweed.yaml; sed -n '310,370p' packages/runner/evals/next_test.go; cat go.mod; cat .gitignore; ls -l .fireweed/runs/drill" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 539ms:
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
func Test_RUN_010_AllBlockedQueue(t *testing.T) {
	f := newNextFixture(t, true, true, "normal")
	change := nextRead(t, f.root, nextChange)
	nextWrite(t, f.root, nextChange, strings.Replace(change, "module: alpha\n", "module: alpha\nblocked: human review\n", 1))
	nextCommit(t, f.root, "block queue")
	r, err := f.run(false)
	if err != nil || !r.Empty || len(r.Steps) != 0 {
		t.Errorf("all-blocked Next = %#v, %v", r, err)
	}
	if refs := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); refs != "" {
		t.Errorf("all-blocked queue claimed change: %s", refs)
	}
	f.noAgent(t)
}

func Test_RUN_011_RUN_012_StartAndApprovalPrecedeAgents(t *testing.T) {
	t.Run("unmerged", func(t *testing.T) {
		f := newNextFixture(t, true, false, "normal")
		before := nextGit(t, f.root, "rev-parse", "HEAD")
		_, err := f.run(false)
		if err == nil || !strings.Contains(strings.ToLower(err.Error()), "approv") {
			t.Errorf("unapproved Next error = %v", err)
		}
		if nextGit(t, f.root, "rev-parse", "HEAD") != before || strings.Contains(nextRead(t, f.root, nextChange), "run:") {
			t.Error("unapproved change was claimed or assigned")
		}
		if out := nextGit(t, f.root, "for-each-ref", "refs/fireweed/claims"); out != "" {
			t.Errorf("unapproved claim = %s", out)
		}
		f.noAgent(t)
	})
	t.Run("start_failure", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextWrite(t, f.root, "stray.txt", "dirty\n")
		_, err := f.run(false)
		if err == nil || !strings.Contains(err.Error(), "P01") {
			t.Errorf("dirty start error = %v", err)
		}
		f.noAgent(t)
	})
	t.Run("first_unblocked", func(t *testing.T) {
		f := newNextFixture(t, true, true, "exit")
		nextWrite(t, f.root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
		nextWrite(t, f.root, "changes/0000-blocked.md", "---\nmodule: beta\nblocked: human review\n---\n\n## Motivation\nBlocked work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		nextCommit(t, f.root, "blocked first change")
		_, err := f.run(false)
		if err == nil {
			t.Error("expected worker failure after eligible start")
		}
		state, currentErr := workflow.Current(f.root, "next")
		if currentErr != nil || state.Change.Slug != "next" || state.Phase != "evals" {
			t.Errorf("selected change = %#v, %v", state, currentErr)
		}
	})
	t.Run("unapproved_archive_bound", func(t *testing.T) {
		f := newNextFixture(t, true, true, "normal")
		nextGit(t, f.root, "update-ref", "refs/remotes/origin/main", "HEAD")
		nextWrite(t, f.root, "changes/archive/0000-prior.md", "---\nmodule: alpha\nphase: applied\nrun: prior\nbases: {}\n---\n\n## Motivation\nPrior work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		nextWrite(t, f.root, "fireweed.yaml", nextRead(t, f.root, "fireweed.yaml")+"authoring:\n  max_unapproved: 1\n")
		nextCommit(t, f.root, "unapproved prior archive")
		for _, d := range project.Check(f.root) {
module github.com/lurkshark/fireweed

go 1.23

require gopkg.in/yaml.v3 v3.0.1
.fireweed/runs/
.fireweed/bin/
bin/
.fireweed/evidence/
total 56
-rw-r--r--@ 1 joel  staff  25033 Nov 10 11:41 report.xml

 exited 2 in 587ms:
fireweed 0.1.0

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
U02: change "runner" not found; no valid change slugs. Use a change held by this worktree and retry context.
U01: Unknown option "--impl". Use help to see available commands and options.

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed verify; fireweed explain P01; fireweed explain P04' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "rg -n 'allow-unmerged|approval warning|P04' fireweed-design-v1.md packages/runner/SPEC.md packages/workflow/SPEC.md; sed -n '1,200p' packages/runner/internal/run/run.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 538ms:
packages/workflow/SPEC.md:46:Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.
fireweed-design-v1.md:259:**Approval.** A change file is approved when it is on the base branch, either `origin/<base>` or the local `<base>`, so a developer who commits a change to local `main` can start it without a warning. `start` warns (P04) only when the change file is on neither. `run`, which works unattended, skips such changes unless given `--allow-unmerged`.
fireweed-design-v1.md:261:**Unapproved work is bounded.** An agent working directly on the local base branch could approve its own changes by committing them, so this bound counts against the shared base: an unapproved change is an archived change on HEAD whose change file is not on `origin/<base>` (or on the local `<base>` when no remote ref exists). Pushing or merging is the human review. When HEAD already has `authoring.max_unapproved` of them (default 5), `start`, `next` and `run` refuse another unmerged change, even with `--allow-unmerged`, with diagnostic P07; merging the work into the base branch, which is a human review, resets the count. This holds whoever authored the changes, so an agent that both writes and runs changes still stops for a person at regular intervals. `status` shows the count, and warns (P08) when three or more consecutive applied changes target the same module, which is the usual shape of a regression chain where each change fixes the edge case the last one exposed.
fireweed-design-v1.md:569:| Unlimited unapproved changes with `--allow-unmerged` | Rejected | An agent authoring and running its own changes never stopped for review; `authoring.max_unapproved` forces a merge checkpoint |
package run

import (
	"context"
	"errors"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/runner/internal/next"
	"github.com/lurkshark/fireweed/packages/runner/internal/support"
	"github.com/lurkshark/fireweed/packages/runner/model"
)

func IsTimeStop(err error) bool { return support.IsTimeStop(err) }

func Execute(parent context.Context, request model.RunRequest) (result model.RunResult, err error) {
	ctx := parent
	budget := &support.Budget{}
	if request.For > 0 {
		var cancel context.CancelFunc
		budget.Deadline = time.Now().Add(request.For)
		budget.Stop = &support.TimeStop{Limit: request.For}
		ctx, cancel = context.WithDeadlineCause(parent, budget.Deadline, budget.Stop)
		defer cancel()
	}
	var blocked []error
	defer func() {
		var invalid *support.CombinedFailure
		if !errors.As(err, &invalid) && ctx.Err() != nil {
			if budget.Stop != nil && context.Cause(ctx) == budget.Stop {
				if err == nil || errors.Is(err, ctx.Err()) {
					err = budget.Stop
				}
			} else if parent.Err() != nil {
				if support.IsTimeStop(err) || support.IsUsageStop(err) {
					err = parent.Err()
				} else if !errors.Is(err, parent.Err()) {
					err = errors.Join(parent.Err(), err)
				}
			}
		}
		if len(blocked) > 0 {
			err = &support.CombinedFailure{Causes: append(blocked, err)}
		}
	}()

	single := model.NextRequest{
		Root: request.Root, Bin: request.Bin, AllowUnmerged: request.AllowUnmerged, Timeout: request.Timeout,
	}
	for request.Max <= 0 || len(result.Applied) < request.Max {
		change, stepErr := next.Execute(ctx, single, budget)
		if stepErr != nil {
			var saved *next.Blocked
			if errors.As(stepErr, &saved) {
				blocked = append(blocked, stepErr)
				continue
			}
			return result, stepErr
		}
		if change.Empty {
			break
		}
		result.Applied = append(result.Applied, change)
	}
	queue, err := project.Queue(request.Root)
	if err != nil {
		return result, err
	}
	result.Empty = len(queue) == 0
	result.Limited = request.Max > 0 && len(result.Applied) == request.Max && !result.Empty
	return result, nil
}

 succeeded in 580ms:
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.
A lifecycle or new-change operational precondition is missing or unmet for the requested operation. Init can also fail when existing files conflict with initialization or a filesystem permission or setup precondition prevents writing the project. Resolve the conflicting existing file or path, or fix the filesystem permission or setup issue. Complete the required precondition and retry from the expected state.
Warning: approval for this change is not recorded on the base branch. Obtain approval through review before unattended execution.

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/DRILL_NOTES.md
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/packages/runner/internal/next/next.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 346d1ca7e24205b7aaa41eb00c997be3297da9ea..ad0f22463839fb8e3aa11c1b86a77e12d2076659
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,3 +6,4 @@
 
 - RUN-003, module runner. The shell template grammar leaves prequoted placeholders unspecified. The reconstruction shell-quotes each replacement, using the same convention as the eval hook.
 - RUN-014, module workflow. Done verifies internally and accepts no prior Verify result. The reconstruction verifies before calling Done, so successful phases can run their eval hooks twice.
+- RUN-041, module runner. The spec requires `IsApprovalRequired` to recognize the refusal. The app evals for APP-075, APP-077 and APP-082 also require P01 for the refusal. The reconstruction uses P01 and preserves the typed approval error. P04 is the warning emitted by `workflow.Start`.
diff --git a/packages/runner/internal/next/next.go b/packages/runner/internal/next/next.go
index 58d0621c3e9b4e4680ade5cc3c4761f501bfb73f..a18ed9f28a94ab6501260fd419b60faf2a24760c
--- a/packages/runner/internal/next/next.go
+++ b/packages/runner/internal/next/next.go
@@ -17,7 +17,7 @@
 type approvalRequired struct{ path string }
 
 func (e *approvalRequired) Error() string {
-	return fmt.Sprintf("P04: change %s requires approval on the local or shared base", e.path)
+	return fmt.Sprintf("P01: change %s requires approval on the local or shared base", e.path)
 }
 
 func IsApprovalRequired(err error) bool {

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; export FIREWEED_BIN="$PWD/.fireweed/bin/fireweed"; mkdir -p .fireweed/runs/drill && go run .fireweed/hooks.go eval packages/runner "$PWD/.fireweed/runs/drill/report.xml"' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc "git status --short --untracked-files=all; git diff --exit-code; rg -n 'requires approval on the local or shared base|RUN-041, module runner' packages/runner/internal/next/next.go DRILL_NOTES.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; gofmt -l packages/runner/internal/next/next.go; fireweed modules' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 587ms:
?? DRILL_NOTES.md
?? packages/runner/internal/drill/drill.go
?? packages/runner/internal/drill/report.go
?? packages/runner/internal/drill/workspace.go
?? packages/runner/internal/invoke/agent.go
?? packages/runner/internal/next/audit.go
?? packages/runner/internal/next/feedback.go
?? packages/runner/internal/next/next.go
?? packages/runner/internal/run/run.go
?? packages/runner/internal/support/config.go
?? packages/runner/internal/support/files.go
packages/runner/internal/next/next.go:20:	return fmt.Sprintf("P01: change %s requires approval on the local or shared base", e.path)
DRILL_NOTES.md:9:- RUN-041, module runner. The spec requires `IsApprovalRequired` to recognize the refusal. The app evals for APP-075, APP-077 and APP-082 also require P01 for the refusal. The reconstruction uses P01 and preserves the typed approval error. P04 is the warning emitted by `workflow.Start`.

 succeeded in 593ms:
assets (AST) packages/assets
  depends_on: [contracts]
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]
checks (CHK) packages/checks
  depends_on: [contracts gitstore assets project]
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]
contracts (CON) packages/contracts
  depends_on: []
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]
gitstore (GIT) packages/gitstore
  depends_on: [contracts]
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]
project (PRJ) packages/project
  depends_on: [contracts gitstore assets]
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]
runner (RUN) packages/runner
  depends_on: [contracts gitstore assets project checks workflow]
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]
workflow (FLOW) packages/workflow
  depends_on: [contracts gitstore assets project checks]
  protected: []
  evals: [evals/**]
  interface: [api.go model/**]
  append_only: []
  preserve: []
  disposable: [internal/**]

codex
The approval refusal now uses `P01` and retains its typed error. I recorded the CLI dependency in `DRILL_NOTES.md`. The runner eval hook is running, and the app cases are absent from this drill workspace.
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; export FIREWEED_BIN="$PWD/.fireweed/bin/fireweed"; fireweed check && go run .fireweed/hooks.go boundaries' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 969ms:
Check passed

codex
`fireweed check` and the boundary hook passed. The runner eval hook is still running. Lifecycle verification requires a held change, which this drill workspace does not have.
 succeeded in 314775ms:
?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/support	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	304.642s

exec
/bin/zsh -lc "python3 -c 'import collections, xml.etree.ElementTree as E; cases=E.parse(\".fireweed/runs/drill/report.xml\").findall(\"testcase\"); print(dict(collections.Counter(\"failed\" if c.find(\"failure\") is not None else \"skipped\" if c.find(\"skipped\") is not None else \"passed\" for c in cases)))'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
exec
/bin/zsh -lc 'git diff --exit-code HEAD; git status --short; git diff --cached --exit-code' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218
 succeeded in 720ms:
?? DRILL_NOTES.md
?? packages/runner/internal/

 succeeded in 852ms:
{'passed': 162}

codex
Corrected the runner’s approval refusal to emit `P01` while preserving `IsApprovalRequired`. Recorded the downstream CLI requirement in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-47583218/DRILL_NOTES.md).

All 162 runner evals passed, along with `fireweed check` and the boundary hook. The three app regressions still need verification in the full checkout. Retained files are intact, and no commit was made.
tokens used
51,408

