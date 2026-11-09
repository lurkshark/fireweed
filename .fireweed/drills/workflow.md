---
date: 2026-11-08T23:37:24.098992Z
base_commit: 68234aa118c7424ac497b7f3987adb21ce006909
agent: codex
model: gpt-6.1-sol
result: fail
---

## Own evals
- workflow | Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange | FLOW-089, FLOW-091, FLOW-094 | pass
- workflow | Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange | FLOW-089, FLOW-091, FLOW-094 | fail
  Failure: === RUN   Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange
    abandon_test.go:116: Abandon owned stepped change: reverse phase commit 5597e8361efb362aef7b99d4edc9060a14c570d9: git apply: exit status 128: error: --cached and --3way cannot be used together.
--- FAIL: Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange (1.55s)
- workflow | Test_FLOW_090_AbandonRejectsDirtyWorkingTree | FLOW-090 | pass
- workflow | Test_FLOW_090_AbandonRejectsDirtyWorkingTree/tracked | FLOW-090 | pass
- workflow | Test_FLOW_090_AbandonRejectsDirtyWorkingTree/untracked | FLOW-090 | pass
- workflow | Test_FLOW_091_AbandonRejectsForeignClaim | FLOW-091 | pass
- workflow | Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending | FLOW-092 | fail
  Failure: === RUN   Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending
    abandon_test.go:210: Abandon after git revert: reverse phase commit b54e11d4cbcd176549fa85dbec2c5464b7b696e9: git apply: exit status 128: error: --cached and --3way cannot be used together.
--- FAIL: Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending (1.74s)
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
- workflow | Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim | FLOW-133 | pass
- workflow | Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim/unclaimed_queued_change | FLOW-133 | pass
- workflow | Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim/empty_valid_queue | FLOW-133 | pass
- workflow | Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim/claim_owned_by_another_worktree | FLOW-133 | pass
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
- workflow | Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly | FLOW-128 | pass
- workflow | Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts | FLOW-129 | pass
- workflow | Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate | FLOW-130, FLOW-131 | pass
- workflow | Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/configured | FLOW-130, FLOW-131 | pass
- workflow | Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/default | FLOW-130, FLOW-131 | pass
- workflow | Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning | FLOW-131 | pass
- workflow | Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning/0s | FLOW-131 | pass
- workflow | Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning/-1h | FLOW-131 | pass
- workflow | Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning/invalid | FLOW-131 | pass
- workflow | Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState | FLOW-132 | pass
- workflow | Test_FLOW_001_FLOW_006_FLOW_007_FLOW_008_FLOW_012_FLOW_013_FLOW_014_FLOW_017_FLOW_018_FLOW_019_FLOW_053_StartAssignsLateIDsAndRetainsSnapshotWithoutCommit | FLOW-001, FLOW-006, FLOW-007, FLOW-008, FLOW-012, FLOW-013, FLOW-014, FLOW-017, FLOW-018, FLOW-019, FLOW-053 | pass
- workflow | Test_FLOW_007_StartReservesCurrentAndPreviouslyRemovedIDs | FLOW-007 | pass
- workflow | Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead | FLOW-010, FLOW-011 | pass
- workflow | Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead/tracked | FLOW-010, FLOW-011 | pass
- workflow | Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead/untracked | FLOW-010, FLOW-011 | pass
- workflow | Test_FLOW_001_FLOW_016_FLOW_017_LinkedWorktreeExcludesCompetingOwner | FLOW-001, FLOW-016, FLOW-017 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/default-below | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/default-at | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit/configured-at | FLOW-009, FLOW-015 | pass
- workflow | Test_FLOW_055_VerifyAndDoneRejectUnsupportedHeldPhase | FLOW-055 | pass
- workflow | Test_FLOW_055_UnsupportedPhasePrecedesProtectedEdit | FLOW-055 | pass
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
- workflow | Test_FLOW_023_ReviewTaggedAdditionNeedsNoFailFirstCase | FLOW-023 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/frozen-eval | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/direct | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/all | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/none | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/quick-boundary | FLOW-002, FLOW-003, FLOW-022, FLOW-025, FLOW-026, FLOW-027, FLOW-029, FLOW-030, FLOW-031 | pass
- workflow | Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures | FLOW-134 | pass
- workflow | Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/direct | FLOW-134 | pass
- workflow | Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/all | FLOW-134 | pass
- workflow | Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/none | FLOW-134 | pass
- workflow | Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_pass/all | FLOW-134 | pass
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
- workflow | Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses | FLOW-062 | fail
  Failure: === RUN   Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses
    lifecycle_test.go:1456: boundary cause = X01: 1 boundary violation(s)
--- FAIL: Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses (5.05s)
- workflow | Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics | FLOW-062 | pass
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
- workflow | Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges | FLOW-115, FLOW-124, FLOW-125 | pass
- workflow | Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges/FLOW_124_binary_and_large_artifacts | FLOW-115, FLOW-124, FLOW-125 | pass
- workflow | Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges/FLOW_115_FLOW_125_empty_and_invalid_ranges | FLOW-115, FLOW-124, FLOW-125 | pass
- workflow | Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive | FLOW-080, FLOW-081, FLOW-083 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/missing | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/older_passing | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/newer_failed | FLOW-082 | pass
- workflow | Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate/newer_passing | FLOW-082 | pass
- workflow | Test_FLOW_081_RecordedRegenerateStartRetainsDisposableFilesAndState | FLOW-081 | pass
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
- workflow | Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing | FLOW-059 | pass
- workflow | Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks | FLOW-002 | pass
- workflow | Test_FLOW_061_UnknownLifecycleSelectorsNameValidSlugs | FLOW-061 | pass
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
FAIL
FAIL	github.com/lurkshark/fireweed/packages/workflow/evals	159.246s

  Stderr: exit status 1


## Dependent evals
- app | Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON | APP-097, APP-099 | fail
  Failure: === RUN   Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON
--- FAIL: Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON (1.86s)
- app | Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/plain-queued | APP-097, APP-099 | pass
- app | Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned | APP-097, APP-099 | fail
  Failure: === RUN   Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned
    abandon_cli_test.go:91: abandon: exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"reverse phase commit 68d5ea6e1f37330ee17909bbff4975530e07b3aa: git apply: exit status 128: error: --cached and --3way cannot be used together.. Use an eligible pending change in a clean Git project, resolve any conflicting claim, then retry abandon.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned (1.33s)
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
- app | Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange | APP-074, APP-076, APP-077 | fail
  Failure: === RUN   Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange
    next_cli_test.go:213: evals reports=[]: <nil>
    next_cli_test.go:213: impl reports=[]: <nil>
--- FAIL: Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange (4.39s)
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
- app | Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation | APP-080, APP-081, APP-083, APP-084 | fail
  Failure: === RUN   Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation
    run_cli_test.go:278: alpha evals reports=[]: <nil>
--- FAIL: Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation (4.30s)
- app | Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON | APP-081, APP-084, APP-086 | pass
- app | Test_APP_083_RunPlainMultiple | APP-083 | pass
- app | Test_APP_086_RunJSONLimited | APP-086 | pass
- app | Test_APP_082_RunAllowUnmerged | APP-082 | pass
- app | Test_APP_081_APP_084_RunUsageAndConfiguration | APP-081, APP-084 | pass
- app | Test_APP_084_APP_085_APP_086_RunWrappedFailure | APP-084, APP-085, APP-086 | pass
- app | Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes | APP-084, APP-089, APP-090 | fail
  Failure: === RUN   Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes
    run_for_cli_test.go:118: alpha evals reports=[]: <nil>
--- FAIL: Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes (8.75s)
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
- app | Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume | APP-076, APP-077, APP-079 | fail
  Failure: === RUN   Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume
    usage_cli_test.go:185: eval reports=[]: <nil>
    usage_cli_test.go:191: impl reports=[]: <nil>
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (4.71s)
- app | Test_APP_076_APP_077_NextPlainUsageStop | APP-076, APP-077 | pass
- app | Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume | APP-083, APP-084, APP-086 | fail
  Failure: === RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:264: alpha evals reports=[]: <nil>
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (5.59s)
- app | Test_APP_083_APP_084_RunPlainSharedUsageStop | APP-083, APP-084 | pass
- app | Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop | APP-084, APP-086 | pass
- app | Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation | APP-077, APP-079, APP-084, APP-086 | pass
- app | Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/next | APP-077, APP-079, APP-084, APP-086 | pass
- app | Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/run | APP-077, APP-079, APP-084, APP-086 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	248.480s

  Stderr: exit status 1

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
- runner | Test_RUN_013_RUN_014_CompleteSpecChange | RUN-013, RUN-014 | fail
  Failure: === RUN   Test_RUN_013_RUN_014_CompleteSpecChange
    next_test.go:557: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3447153749/001/.fireweed/runs/581b94296216f4de3d77c943e60ea147/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_013_RUN_014_CompleteSpecChange (4.11s)
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
- runner | Test_RUN_016_UnlimitedQueueOrder | RUN-016 | fail
  Failure: === RUN   Test_RUN_016_UnlimitedQueueOrder
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_016_UnlimitedQueueOrder1554671597/001/.fireweed/runs/afcd4d3ba3fab3b40acd3258025b143a/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_016_UnlimitedQueueOrder (12.96s)
- runner | Test_RUN_016_NegativeMaxIsUnlimited | RUN-016 | fail
  Failure: === RUN   Test_RUN_016_NegativeMaxIsUnlimited
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_016_NegativeMaxIsUnlimited1747265915/001/.fireweed/runs/2367cc802539c73ccee6b330beeab479/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_016_NegativeMaxIsUnlimited (4.20s)
- runner | Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding | RUN-017, RUN-018, RUN-019 | fail
  Failure: === RUN   Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding449182212/001/.fireweed/runs/2e075d5314b45accad54f1d5220b31dd/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding (4.16s)
- runner | Test_RUN_018_InitialEmptyAndExactLimit | RUN-018 | fail
  Failure: === RUN   Test_RUN_018_InitialEmptyAndExactLimit
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit (8.39s)
- runner | Test_RUN_018_InitialEmptyAndExactLimit/initial_empty | RUN-018 | pass
- runner | Test_RUN_018_InitialEmptyAndExactLimit/exact_limit | RUN-018 | fail
  Failure: === RUN   Test_RUN_018_InitialEmptyAndExactLimit/exact_limit
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_018_InitialEmptyAndExactLimitexact_limit2370949650/001/.fireweed/runs/c31cd8502ae1bdc4d9f9ef52869dd65b/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit/exact_limit (8.16s)
- runner | Test_RUN_020_HeldFailureRetainsEarlierResults | RUN-020 | fail
  Failure: === RUN   Test_RUN_020_HeldFailureRetainsEarlierResults
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_020_HeldFailureRetainsEarlierResults3216587672/001/.fireweed/runs/0773b132ee8522ec1dbbe23e7b39600a/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_020_HeldFailureRetainsEarlierResults (4.76s)
- runner | Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent | RUN-016, RUN-025, RUN-026 | fail
  Failure: === RUN   Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent2150866924/001/.fireweed/runs/21aec49c3da8d0d6bb0544551d6a9478/evals/gamma.json: no such file or directory
--- FAIL: Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent (6.32s)
- runner | Test_RUN_025_RUN_026_PositiveMaxRetainsBlock | RUN-025, RUN-026 | fail
  Failure: === RUN   Test_RUN_025_RUN_026_PositiveMaxRetainsBlock
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_025_RUN_026_PositiveMaxRetainsBlock988106862/001/.fireweed/runs/bf1b5ae48aab4b89289ea7543c21e6df/evals/gamma.json: no such file or directory
--- FAIL: Test_RUN_025_RUN_026_PositiveMaxRetainsBlock (6.33s)
- runner | Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied | RUN-036, RUN-037 | fail
  Failure: === RUN   Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied2451242838/001/.fireweed/runs/ef5b92d2095e154e9460df72f9f31bea/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied (10.24s)
- runner | Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded | RUN-036 | fail
  Failure: === RUN   Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded
    run_test.go:182: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded1372614714/001/.fireweed/runs/77f7cafa3a564568c6c08ca5c7af1ade/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded (2.87s)
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
- runner | Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied | RUN-032 | fail
  Failure: === RUN   Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied
    usage_test.go:329: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied2094440256/001/.fireweed/runs/9a05694f991ec8ce3bd5283f3cd97f09/evals/alpha.json: no such file or directory
--- FAIL: Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied (4.94s)
- runner | Test_RUN_034_UsageStopClassifierWrapping | RUN-034 | pass
- runner | Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure | RUN-035 | pass
- runner | Test_RUN_031_CancelledUsageWaitRetainsHeldWork | RUN-031 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/runner/evals	247.865s

  Stderr: exit status 1


## Contradicted Decisions
None found in the retained workflow Decisions.
## Gaps
- FLOW-001 and FLOW-019 · workflow. The claim payload format is left to the implementation. The rebuild stores the canonical worktree, change identity, phase base, approved document and projected snapshot in a JSON blob addressed by the claim ref.
- FLOW-020 and FLOW-075 · workflow. The spec leaves the complete list of worker-editable body sections open. The retained evals establish Decisions and Questions for evals and Decisions for impl, and Audit for audit. The rebuild also treats Questions as editable during impl.
- FLOW-059 · workflow. The spec leaves the audit submission grammar open. The rebuild uses the checked-ID line and quoted finding bullets from the retained evals, with an optional untestable marker.
- FLOW-052 and FLOW-064 · workflow. The spec does not describe how retained run output interacts with staging. Done and Block stage tracked paths and Git-visible untracked paths from a private index, excluding `.fireweed/runs` even when a project's ignore rules omit it.
- FLOW-074 · workflow. The spec leaves conflict resolution open when a saved attempt overlaps later commits. The rebuild applies the saved delta, tries Git's three-way merge when needed, and returns an error for unresolved conflicts.
- FLOW-113 and FLOW-123 · workflow. Saved budget diagnostics have no dedicated module field. The rebuild uses a module name or a path within a module, then the module or spec ID in the diagnostic text. Aggregate diagnostics and warnings without a usable module key stay associated with their selected snapshot.
- FLOW-117 · workflow. The spec leaves file classification open for historical commits that predate loadable project metadata. The rebuild retains their recorded findings and result counts and classifies files for commits with loadable metadata.
- FLOW-128 · workflow. The spec does not set a minimum Git version. GC falls back to line-delimited porcelain output when Git rejects `worktree list --porcelain -z`.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt the workflow implementation under `packages/workflow/internal/lifecycle` and documented the implementation choices and spec gaps in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/DRILL_NOTES.md).

`gofmt` passed. Validation was limited to formatting and source review because `SYSTEM.md` reserves builds and evals for the separate verifier. No commit was created.

Reading additional input from stdin...
2026-11-08T23:39:29.293136Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1bd88-004e-7c71-af4f-0efb196c5f7c
--------
user
# Rebuild workflow

Target spec and Decisions:
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
id_prefix: FLOW
interface:
    - api.go
    - model/**
name: workflow
---

# workflow

## Specs

### FLOW-001
Start MUST atomically create the local claim for the selected queued change.

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before running hooks, checking differences since the held phase base.

### FLOW-003
Done MUST require full verification to succeed before committing a phase transition.

### FLOW-004
Successful impl Done MUST apply Added, Modified and Removed spec edits to the module's SPEC.md.

### FLOW-005
Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook. It MUST list other relevant files by path and size, or include their contents when full is requested.

### FLOW-006
Start MUST include the approved projected project snapshot in its returned State.

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next
available spec ID for its module. IDs of current and previously removed specs
MUST remain reserved, including specs Removed by the starting change.

### FLOW-008
Start MUST assign a unique run identifier when starting an unstarted change.

### FLOW-009
Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.

### FLOW-010
Start MUST reject a worktree containing uncommitted changes.

### FLOW-011
Start MUST reject a requested change unless it is the first change returned by
Eligible. With no requested target, Start MUST start that first eligible
change. When Eligible returns no change, Start MUST reject with P01.

### FLOW-012
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### FLOW-013
Start MUST set phase to evals when starting an unstarted spec change.

### FLOW-014
Start MUST leave HEAD unchanged.

### FLOW-015
Start MUST reject with P07 when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

### FLOW-016
Start MUST reject a competing worktree that attempts to claim the same change.

### FLOW-017
Current MUST resolve only a claim owned by the calling worktree.

### FLOW-018
Start MUST use the current HEAD as the initial phase base of its claim.

### FLOW-019
While a claim is held, Current MUST return the project snapshot approved at Start.

### FLOW-020
Verify MUST reject an unapproved alteration of the active change's author fields before running a hook.

### FLOW-021
Verify MUST reject an unapproved alteration of the active change's recorded runtime state before running a hook.

### FLOW-022
Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec.

### FLOW-023
During evals Verify for spec and fix changes MUST require genuine fail-first
evidence for every enforced change ID that is not already_satisfied.
Review-tagged specs MUST be excluded from fail-first requirements.

### FLOW-024
During impl Verify MUST reject a failing current-module testcase that references a non-quarantined spec.

### FLOW-025
Full impl Verify MUST use verify_dependents with a default of direct to select dependent modules.

### FLOW-026
Full impl Verify MUST reject a failed enforced testcase from a selected dependent module.

### FLOW-027
Full impl Verify MUST require boundary verification to succeed.

### FLOW-028
Verify MUST select hooks from the held approved projected project snapshot.

### FLOW-029
Quick Verify MUST skip dependent-module evals.

### FLOW-030
Quick Verify MUST skip boundary verification.

### FLOW-031
Verify MUST leave HEAD unchanged.

### FLOW-032
Done MUST commit a successfully verified evals phase and advance it to audit when audit is enabled.

### FLOW-033
Done MUST commit a successfully verified evals phase and advance it to impl when audit.enabled is false.

### FLOW-034
Audit Done MUST discard a finding whose spec ID is absent or whose quoted clause is not an exact substring of that spec.

### FLOW-035
On re-audit, Audit Done MUST discard findings for spec IDs that had no retained finding in the preceding audit round.

### FLOW-036
Audit Done MUST require every enforced ID in the change's Added, Modified,
fixes or eval_fix entries to appear in a No findings submission's checked-ID list.

### FLOW-037
Retained testable audit findings within the round limit MUST return the change to evals.

### FLOW-038
An audit return to evals MUST increment and persist audit_rounds.

### FLOW-039
Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.

### FLOW-040
Audit Done MUST block the change when a retained finding marks its quoted clause untestable.

### FLOW-041
Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.

### FLOW-042
When audit blocks a change, Done MUST include the blocking findings in its Questions section.

### FLOW-043
Audit blocking MUST release the calling worktree's held change claim.

### FLOW-044
Successful impl Done MUST apply the change's set values to module metadata.

### FLOW-045
Successful impl Done MUST apply Decisions attributed to the change slug and remove entries quoted in Removed decisions.

### FLOW-046
Each successful Done MUST advance the phase base to its new commit.

### FLOW-047
Successful impl Done MUST release only the exact held claim upon application.

### FLOW-048
Failed Done verification MUST leave the phase unadvanced and create no commit.

### FLOW-049
Impl Done MUST require the applied project to pass static validation before committing.

### FLOW-050
Successful impl Done MUST move the active change to the archive with phase applied and commit it together with the application.

### FLOW-051
Each successful Done MUST append one provenance line containing phase, date,
manual or configured agent, version, playbook hash and result counts.
It MUST preserve existing nonblank provenance lines unchanged and in order.
It MAY normalize blank lines within the Provenance section.

### FLOW-052
Successful Done MUST retain verification reports and process output under the run directory.

### FLOW-053
State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

### FLOW-054
Start MUST prefix dirty-worktree and queue-eligibility rejections with P01.

### FLOW-055
Verify and Done MUST prefix unsupported-phase rejections with P01.

### FLOW-056
Start and Current MUST prefix existing claim ownership or contention failures with P02.

### FLOW-057
Verify and Done MUST prefix unapproved author, runtime or audit-state alteration errors with P03 before running hooks.

### FLOW-058
Start, Verify and Done MUST prefix stale-base rejections with P05, retaining the underlying cause.

### FLOW-059
During audit, Verify MUST apply Done's audit-submission checks. Verify and Done MUST prefix malformed or incomplete submission rejections with P09.

### FLOW-060
Done MUST prefix its existing blocked-change error with P06 while preserving the committed blocking state and claim release.

### FLOW-061
Start MUST distinguish an explicit unknown change selector from queue eligibility. Lifecycle operations MUST return U02 for unknown explicit change selectors, naming the rejected value and listing valid slugs for the operation.

### FLOW-062
Lifecycle operations MUST preserve assigned static, eval, zone and boundary causes when returning their existing failures, except for lifecycle stale-base classification P05.

### FLOW-063
Block MUST require the calling worktree's current claim with unchanged approved
metadata and phase. It MUST reject a changed HEAD rather than blocking a stale
phase base.

### FLOW-064
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior attempt
refs. BlockResult MUST identify that ref and its commit OID.

### FLOW-065
After saving the attempt, Block MUST restore the step's tracked files, untracked
files and index to the phase base, removing the step's edits from the checkout.

### FLOW-066
Block MUST commit only the held change's assigned IDs, run, current phase and
requested blocked reason over the phase base. The block commit MUST preserve
the approved author and already-committed audit content and exclude step work.

### FLOW-067
On successful blocking, Block MUST release the current claim and return the
blocked State with no held ClaimOID. The resulting working tree MUST be clean.

### FLOW-068
Eligible MUST return unblocked queued changes not held back by earlier blocked
changes, in queue order. Empty or entirely waiting queues MUST return an empty
result without an error.

### FLOW-069
Eligible MUST omit a later change targeting the same module as an earlier
blocked change.

### FLOW-070
Eligible MUST omit a later change whose module depends directly or transitively
on an earlier blocked change's module.

### FLOW-071
When a later change has a nonempty Interface section, Eligible MUST omit it if
an earlier blocked change's module depends directly or transitively on its
module. With an empty Interface, that dependency change MUST remain eligible
unless another blocking rule applies.

### FLOW-072
For each candidate, Eligible MUST apply the blocking rules to the graph
projected through that candidate, including preceding queued set edits and its
own set edits.

### FLOW-073
StartFromAttempt MUST resume a committed phase with the same clean-worktree,
eligibility and claim-ownership preconditions as Start. It MUST preserve the
existing run, phase and assigned IDs and leave HEAD unchanged.

### FLOW-074
StartFromAttempt MUST restore the saved step's modified, deleted and newly
added files from refs/fireweed/attempts/<canonical-slug>/<attempt> into the
working tree without reverting later independent changes.

### FLOW-075
While restoring saved worker-editable body sections for the resumed phase,
StartFromAttempt MUST preserve the current change's approved author fields,
runtime state and controller-written provenance.

### FLOW-076
StartFromAttempt MUST reject a non-positive attempt number or missing saved
attempt ref with P01 before claiming or writing the change.

### FLOW-077
For an unstarted change with no Added, Modified or Removed specs, fixes or
eval_fix, and regenerate not true, Start MUST set phase to impl.

### FLOW-078
For an eval_fix change, Verify MUST permit evals-phase edits only in the target
module's evals zone and the current change file. During impl it MUST permit
edits only to that change file.

### FLOW-079
During evals for an eval_fix change, Verify MUST require all target-module
evals to pass rather than fail first, with every listed eval_fix ID covered by
at least one referencing testcase.

### FLOW-080
Start MUST set phase to impl for an unstarted change with regenerate true.

### FLOW-081
Before returning from successful Start of an unstarted regenerate change,
Fireweed MUST delete only the target module's files classified as disposable
by the approved projected zone map, retaining all other files. When claiming a
regenerate change with an already recorded phase, Start MUST retain its existing
disposable files.

### FLOW-082
For a regenerate change, Start MUST emit warning P10 when the module has no
passing drill dated after its latest applied change's implementation Provenance.
A newer passing drill MUST suppress that warning.

### FLOW-083
For a held regenerate change in impl, Context MUST supply the regenerate
playbook in place of the ordinary impl playbook.

### FLOW-084
Release MUST drop only the named claim owned by the calling worktree while
preserving its working files, HEAD and recorded change state.

### FLOW-085
Release MUST reject a missing claim or a claim owned by another worktree with
P02, leaving refs unchanged.

### FLOW-086
Release MUST accept the same unambiguous change path, filename and slug
selectors as Current.

### FLOW-087
Commits MUST return the full object IDs, newest first, of commits reachable
from HEAD that change the selected change document's phase relative to their
first parent, following its pending and archive paths. Selection MUST use file
contents rather than commit messages.

### FLOW-088
Commits MUST resolve pending or archived changes by repository-relative path,
filename or unambiguous slug. It MUST reject missing or ambiguous selectors
with U02.

### FLOW-089
Abandon MUST undo the unapplied change's phase-commit effects identified by Commits and delete its pending change file in one commit, returning that commit's full object ID.

### FLOW-090
Abandon MUST reject a dirty working tree with P01 before changing files, HEAD or refs.

### FLOW-091
Abandon MUST allow an unclaimed change or a claim owned by the calling worktree.
It MUST reject a foreign claim with P02 before changing files, HEAD or refs.
On success it MUST release the selected owned claim.

### FLOW-092
Abandon MUST reject an applied archived change with P01.
After git revert returns that change to the pending queue, Abandon MUST allow dropping it under the same preconditions.

### FLOW-093
Abandon MUST retain saved attempt refs unchanged.

### FLOW-094
The Abandon commit MUST not assign a new phase value.

### FLOW-095
After successful Abandon, its assigned but unapplied spec IDs MUST be available for subsequent Start allocation.

### FLOW-096
Abandon MUST accept the same unambiguous repository-relative path, filename and slug selectors as Commits, rejecting missing or ambiguous selectors with U02.

### FLOW-097
Each Verify that evaluates modules, including through Done, MUST retain a cost snapshot under its existing run/phase directory and return those module costs in Result.Costs.
The snapshot MUST identify the run, phase and UTC recording time, replacing the previous cost snapshot for that run/phase.

### FLOW-098
Each evaluated module's cost MUST report its name, summed JUnit testcase seconds and number of matching testcases per spec ID, counting each testcase once per matching ID.
It MUST report eval and disposable line counts for regular files selected by the held approved projected zone map, counting an unterminated final line and excluding higher-priority zones.

### FLOW-099
Verify and Done MUST return warnings B01 for total evaluated-module seconds above budgets.verify_time (default15m), B02 for each spec's testcase count above budgets.cases_per_spec (default15), and B03 for each module's eval/disposable line ratio above budgets.eval_ratio (default4).
Warnings MUST identify measurement and limit without affecting verification success or failure.
Positive eval lines with zero disposable lines MUST exceed the ratio limit.

### FLOW-100
Costs MUST return the retained run/phase cost snapshots in recording-time order with their module values and budget diagnostics, or an empty slice when none exist.
The retained values MUST reflect the evaluated state even if source files or budget settings later change.

### FLOW-101
Status MUST return every project.Queue entry in queue order with slug, path, module, phase, blocked reason and audit rounds.
Absent phase, reason and rounds MUST appear as queued, empty string and zero respectively.

### FLOW-102
For an unblocked queued change targeting an earlier blocked change's module, Status MUST name that earlier change in WaitingOn with WaitingRule same_module.

### FLOW-103
For an unblocked queued change whose module depends directly or transitively on an earlier blocked change's module, Status MUST name that earlier change with WaitingRule depends_on_blocked.

### FLOW-104
For an unblocked queued change with a nonempty Interface section, Status MUST name an earlier blocked change with WaitingRule interface_dependency when that blocked module depends directly or transitively on the candidate's module.
An empty Interface MUST not trigger this rule.

### FLOW-105
Status MUST determine waiting reasons using the graph projected through each candidate, including preceding and candidate set edits.
Blocked entries and unblocked entries unaffected by the blocking rules MUST have empty WaitingOn and WaitingRule.

### FLOW-106
Status MUST expose the discovered modules with their declared identities, dependency graph and resolved zones.

### FLOW-107
Status MUST list every currently quarantined spec with its module, ID and quarantine reason.

### FLOW-108
Status MUST list every module with boundaries_opt_out and its authored reason.

### FLOW-109
Status MUST report each module's count of applied archived changes whose implementation Provenance date is later than its latest drill report date.
With no drill report, it MUST count all that module's applied archived changes, including zero when there are none.

### FLOW-110
Status MUST count applied changes archived on HEAD whose change file is absent from the locally known configured remote/base ref, falling back to the local base when the remote ref is absent.
The base and remote MUST use the existing configured/default base-ref resolution.

### FLOW-111
Status MUST warn with P08 when the three or more most recently applied changes consecutively target one module, identifying the module and consecutive count.

### FLOW-112
Status MUST report the number of phase-changing commits reachable from HEAD for each applied archived change, matching Commits for that change.

### FLOW-113
Status MUST expose the latest retained ModuleCost for every module with recorded costs, together with the saved budget diagnostics associated with those selected measurements.
Modules without a cost record MUST not receive a fabricated measurement.

### FLOW-114
PRBody MUST render a Markdown description for an explicit Git commit range, identifying its resolved endpoints.
Project configuration, module/spec metadata and change documents MUST come from the range endpoint rather than later HEAD or working-tree edits.

### FLOW-115
PRBody MUST identify each pending or archived change present at the endpoint whose document identity was touched in the range, including its path, module and change kind.
An empty selection MUST be reported explicitly.

### FLOW-116
For every selected change, PRBody MUST include its recorded eval result counts from Provenance and identify its eval-file edits.

### FLOW-117
PRBody MUST identify interface edits, new disposable stubs added in evals, module manifest/other-file edits and worker_writable edits associated with each selected change.

### FLOW-118
PRBody MUST identify preserved-file edits associated with selected changes.

### FLOW-119
PRBody MUST list every review-tagged Added or Modified spec of each selected change with its ID and normative text.
It MUST not list an unrelated pre-existing review spec as changed.

### FLOW-120
PRBody MUST include each selected change's audit_rounds and recorded audit findings, including findings in earlier phase commits that were corrected before the endpoint.

### FLOW-121
PRBody MUST include the selected changes' added and removed Decisions.

### FLOW-122
PRBody MUST list all quarantined endpoint specs with module, ID and reason, every endpoint boundary opt-out with module and reason, and every blocked pending endpoint change with its reason.

### FLOW-123
PRBody MUST show the latest retained measurements and associated B01–B03 diagnostics for modules represented at the endpoint, without fabricating missing measurements.
Saved local cost snapshots MUST remain the measurement source rather than later file counts or recalculated budgets.

### FLOW-124
PRBody MUST list every binary file and every file larger than 256 KB newly added to a disposable zone between the range endpoints, identifying path and byte size.
Each such file MUST produce a B04 warning in Diagnostics and the Markdown body; ordinary smaller text additions MUST not produce B04.

### FLOW-125
An empty or unresolvable range MUST return an error with P01 and the rejected range or resolution cause.
A valid empty range MUST return a report rather than an error.

### FLOW-126
PRBody MUST return an error for which IsEndpointConfigError reports true when the project at the range endpoint fails to load.

### FLOW-127
IsEndpointConfigError MUST report false when PRBody fails because the range cannot be resolved.

### FLOW-128
GC MUST prune refs/fireweed/claims entries whose recorded worktree is missing or is no longer registered with the same Git common directory, while retaining claims for live registered worktrees.

### FLOW-129
GC MUST prune refs/fireweed/attempts/<slug>/* only for changes present in changes/archive, retaining attempt refs for unapplied changes.

### FLOW-130
GC MUST prune refs/fireweed/drills/<module>/<date> when the UTC timestamp encoded in the native drill ref name is older than gc.drill_max_age relative to GCRequest.Now, or current UTC time when Now is zero, retaining refs at or newer than the cutoff.

### FLOW-131
GC MUST read gc.drill_max_age as a positive Go duration, defaulting to 720h (30 days) when unset. An invalid configured value MUST return C01 before any pruning.

### FLOW-132
GC MUST return the sorted full names of only its successfully deleted refs. It MUST leave all other refs, HEAD, index and working files unchanged.

### FLOW-133
For a valid queue, Current(root, "") MUST return an error whose text is `P02: worktree holds no change` when the calling worktree owns none of its change claims.

### FLOW-134
During full impl Verify, mapped testcase failures in valid reports MUST NOT
stop evaluation of the remaining selected modules.
Verify MUST return every completed module's report in Result.Reports.
Cancellation, timeouts, hook execution errors, and invalid or missing reports
MAY stop further evaluation.

## Decisions

[0012-manual-lifecycle] Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.

[0119-verification-costs] Cost snapshots use .fireweed/runs/<run>/<phase>/costs.json beside existing module
reports; repeated verification in that phase replaces its snapshot, as it does
the module reports. Each regular file's line count is newline count plus one
when nonempty content lacks a trailing newline.

[0120-status-report] When several earlier changes hold one entry, report the first in queue order,
with same-module, dependent and interface-dependency precedence. Applied order
uses implementation Provenance dates, newest first with archive path as the
tie-breaker; a drill report date counts regardless of result. Budget diagnostics
from selected snapshots are returned once each, rather than recomputed from
later settings or source edits. PhaseCommits keys are archived change filenames
without .md.

[0123-pr-body-report] Selected changes are identities touched by range commits and present at its end;
removed change documents are absent from that endpoint report. Review context for
a selected change includes its phase commits reachable at the endpoint, so a
range covering only its final phase retains earlier eval/stub/audit information.
Binary follows Git's text/binary classification; 256 KB is262144 bytes, with
strictly larger text files warned. File categories use that phase's declared
metadata; endpoint exceptions use the endpoint's declarations.
Cost measurements use a Cost/Metric Markdown section with one row per measured
module: module, seconds, cases per spec (`ID: count`), eval lines, disposable lines.


Retained interfaces and files:
packages/workflow/SPEC.md
packages/workflow/api.go
packages/workflow/model/model.go

Retained evals:
packages/workflow/evals/.gitkeep
packages/workflow/evals/abandon_test.go
packages/workflow/evals/block_test.go
packages/workflow/evals/commits_test.go
packages/workflow/evals/costs_test.go
packages/workflow/evals/current_absence_test.go
packages/workflow/evals/eligible_test.go
packages/workflow/evals/eval_fix_lifecycle_test.go
packages/workflow/evals/gc_test.go
packages/workflow/evals/lifecycle_test.go
packages/workflow/evals/maintenance_start_test.go
packages/workflow/evals/pr_body_test.go
packages/workflow/evals/regenerate_start_test.go
packages/workflow/evals/release_test.go
packages/workflow/evals/resume_test.go
packages/workflow/evals/selector_audit_validation_test.go
packages/workflow/evals/status_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.

codex
I’ll read the retained workflow API, evals, and dependency interfaces, then rebuild the disposable implementation and run the workflow evals.
exec
/bin/zsh -lc "pwd; rg --files --hidden -g '"'!.git'"' -g '"'!.git/**'"' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 444ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
SYSTEM.md
packages/assets/api.go
packages/assets/model/model.go
packages/checks/api.go
packages/checks/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/project/api.go
packages/project/model/model.go
packages/workflow/SPEC.md
packages/workflow/api.go
packages/workflow/evals/.gitkeep
packages/workflow/evals/abandon_test.go
packages/workflow/evals/block_test.go
packages/workflow/evals/commits_test.go
packages/workflow/evals/costs_test.go
packages/workflow/evals/current_absence_test.go
packages/workflow/evals/eligible_test.go
packages/workflow/evals/eval_fix_lifecycle_test.go
packages/workflow/evals/gc_test.go
packages/workflow/evals/lifecycle_test.go
packages/workflow/evals/maintenance_start_test.go
packages/workflow/evals/pr_body_test.go
packages/workflow/evals/regenerate_start_test.go
packages/workflow/evals/release_test.go
packages/workflow/evals/resume_test.go
packages/workflow/evals/selector_audit_validation_test.go
packages/workflow/evals/status_test.go
packages/workflow/model/model.go

exec
/bin/zsh -lc "rg -n '"'^func |''^type '"' packages/workflow/evals" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc 'cat SYSTEM.md packages/workflow/api.go packages/workflow/model/model.go; ls -la; ls -la packages/workflow' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc 'cat packages/{assets,checks,contracts,gitstore,project}/{api.go,model/model.go}' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 545ms:
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

 succeeded in 553ms:
packages/workflow/evals/costs_test.go:18:func costAt(t *testing.T, root, run, phase string) (workflow.CostSnapshot, bool) {
packages/workflow/evals/costs_test.go:34:func costFor(costs []workflow.ModuleCost, name string) (workflow.ModuleCost, bool) {
packages/workflow/evals/costs_test.go:43:func warningFor(ds []project.Diagnostic, code string) (project.Diagnostic, bool) {
packages/workflow/evals/costs_test.go:52:func requireCostWarning(t *testing.T, ds []project.Diagnostic, code string, terms ...string) {
packages/workflow/evals/costs_test.go:66:func Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot(t *testing.T) {
packages/workflow/evals/costs_test.go:134:func Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts(t *testing.T) {
packages/workflow/evals/costs_test.go:187:func Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure(t *testing.T) {
packages/workflow/evals/costs_test.go:255:func Test_FLOW_099_DefaultBudgetBoundaries(t *testing.T) {
packages/workflow/evals/costs_test.go:314:func Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits(t *testing.T) {
packages/workflow/evals/selector_audit_validation_test.go:13:func Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing(t *testing.T) {
packages/workflow/evals/selector_audit_validation_test.go:71:func Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks(t *testing.T) {
packages/workflow/evals/selector_audit_validation_test.go:115:func Test_FLOW_061_UnknownLifecycleSelectorsNameValidSlugs(t *testing.T) {
packages/workflow/evals/abandon_test.go:15:func abandonGitIdentity(t *testing.T, root string) {
packages/workflow/evals/abandon_test.go:21:func assertAbandonCommit(t *testing.T, root, parent, oid string) {
packages/workflow/evals/abandon_test.go:43:func assertAbandonRefusal(t *testing.T, root, selector, code string) {
packages/workflow/evals/abandon_test.go:66:func Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange(t *testing.T) {
packages/workflow/evals/abandon_test.go:87:func Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange(t *testing.T) {
packages/workflow/evals/abandon_test.go:133:func Test_FLOW_090_AbandonRejectsDirtyWorkingTree(t *testing.T) {
packages/workflow/evals/abandon_test.go:153:func Test_FLOW_091_AbandonRejectsForeignClaim(t *testing.T) {
packages/workflow/evals/abandon_test.go:174:func Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending(t *testing.T) {
packages/workflow/evals/abandon_test.go:218:func Test_FLOW_093_FLOW_095_AbandonRetainsRealAttemptAndFreesAssignedID(t *testing.T) {
packages/workflow/evals/abandon_test.go:267:func Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal(t *testing.T) {
packages/workflow/evals/commits_test.go:17:func commitPhase(t *testing.T, root, path, old, next, message string) string {
packages/workflow/evals/commits_test.go:34:func assertCommits(t *testing.T, root, selector string, want []string) {
packages/workflow/evals/commits_test.go:50:func Test_FLOW_087_FLOW_088_CommitsFollowPendingAndArchivePhaseContents(t *testing.T) {
packages/workflow/evals/commits_test.go:97:func Test_FLOW_087_CommitsUseFirstParentOfMerge(t *testing.T) {
packages/workflow/evals/commits_test.go:126:func Test_FLOW_088_CommitsRejectMissingAndAmbiguousSelectors(t *testing.T) {
packages/workflow/evals/release_test.go:11:func releaseRefs(t *testing.T, root string) string {
packages/workflow/evals/release_test.go:16:func Test_FLOW_084_ReleaseOwnedClaimPreservesDirtyWorkAndRecordedState(t *testing.T) {
packages/workflow/evals/release_test.go:67:func Test_FLOW_085_ReleaseMissingAndForeignClaimsPreserveRefs(t *testing.T) {
packages/workflow/evals/release_test.go:109:func Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors(t *testing.T) {
packages/workflow/evals/maintenance_start_test.go:13:func Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly(t *testing.T) {
packages/workflow/evals/pr_body_test.go:16:func prBody(t *testing.T, root, revisionRange string) workflow.PRBodyReport {
packages/workflow/evals/pr_body_test.go:25:func prRequire(t *testing.T, body string, terms ...string) {
packages/workflow/evals/pr_body_test.go:34:func prReject(t *testing.T, body string, terms ...string) {
packages/workflow/evals/pr_body_test.go:43:func prPhase(t *testing.T, root, phase string) {
packages/workflow/evals/pr_body_test.go:63:func prHistoryFixture(t *testing.T) (root, base, finalPhase, endpoint string) {
packages/workflow/evals/pr_body_test.go:131:func Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint(t *testing.T) {
packages/workflow/evals/pr_body_test.go:194:func prSaveCost(t *testing.T, root, run, phase string, snapshot workflow.CostSnapshot) {
packages/workflow/evals/pr_body_test.go:203:func Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts(t *testing.T) {
packages/workflow/evals/pr_body_test.go:263:func Test_FLOW_126_PRBodyEndpointConfigError(t *testing.T) {
packages/workflow/evals/pr_body_test.go:278:func Test_FLOW_127_PRBodyUnresolvableRangeIsNotEndpointConfigError(t *testing.T) {
packages/workflow/evals/pr_body_test.go:289:func Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges(t *testing.T) {
packages/workflow/evals/status_test.go:15:func status(t *testing.T, root string) workflow.StatusReport {
packages/workflow/evals/status_test.go:24:func statusEntry(t *testing.T, r workflow.StatusReport, numberedSlug string) workflow.StatusEntry {
packages/workflow/evals/status_test.go:35:func waiting(t *testing.T, r workflow.StatusReport, slug, on, rule string) {
packages/workflow/evals/status_test.go:43:func Test_FLOW_101_FLOW_102_FLOW_103_StatusQueueAndForwardWaiting(t *testing.T) {
packages/workflow/evals/status_test.go:88:func Test_FLOW_104_FLOW_105_StatusInterfaceAndProjectedWaiting(t *testing.T) {
packages/workflow/evals/status_test.go:114:func Test_FLOW_106_FLOW_107_FLOW_108_StatusModuleMetadata(t *testing.T) {
packages/workflow/evals/status_test.go:146:func Test_FLOW_109_FLOW_111_StatusProvenanceDriftAndRecentModuleWarning(t *testing.T) {
packages/workflow/evals/status_test.go:170:func Test_FLOW_110_StatusSharedBaseUsesSameArchivedDocumentIdentity(t *testing.T) {
packages/workflow/evals/status_test.go:226:func Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle(t *testing.T) {
packages/workflow/evals/status_test.go:251:func Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings(t *testing.T) {
packages/workflow/evals/status_test.go:320:func containsDiagnostic(ds []project.Diagnostic, want project.Diagnostic) bool {
packages/workflow/evals/current_absence_test.go:11:func Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:28:func write(t *testing.T, root, name, content string) {
packages/workflow/evals/lifecycle_test.go:39:func git(t *testing.T, root string, args ...string) string {
packages/workflow/evals/lifecycle_test.go:51:func head(t *testing.T, root string) string { t.Helper(); return git(t, root, "rev-parse", "HEAD") }
packages/workflow/evals/lifecycle_test.go:52:func commit(t *testing.T, root, message string) {
packages/workflow/evals/lifecycle_test.go:58:func fixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
packages/workflow/evals/lifecycle_test.go:125:func start(t *testing.T, root string) workflow.State {
packages/workflow/evals/lifecycle_test.go:134:func request(root string) workflow.Request {
packages/workflow/evals/lifecycle_test.go:137:func verify(root string, quick bool) (workflow.Result, error) {
packages/workflow/evals/lifecycle_test.go:142:func done(root string) (workflow.Result, error) {
packages/workflow/evals/lifecycle_test.go:146:func containsCode(ds []project.Diagnostic, code string) bool {
packages/workflow/evals/lifecycle_test.go:155:func Test_FLOW_001_FLOW_006_FLOW_007_FLOW_008_FLOW_012_FLOW_013_FLOW_014_FLOW_017_FLOW_018_FLOW_019_FLOW_053_StartAssignsLateIDsAndRetainsSnapshotWithoutCommit(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:205:func Test_FLOW_007_StartReservesCurrentAndPreviouslyRemovedIDs(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:249:func Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:275:func Test_FLOW_001_FLOW_016_FLOW_017_LinkedWorktreeExcludesCompetingOwner(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:320:func Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:379:func Test_FLOW_055_VerifyAndDoneRejectUnsupportedHeldPhase(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:391:func Test_FLOW_055_UnsupportedPhasePrecedesProtectedEdit(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:419:func Test_FLOW_056_FLOW_061_CurrentAndContextIdentifyMissingClaimsAndChanges(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:432:func Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:474:func hookMarker(t *testing.T, root string) string {
packages/workflow/evals/lifecycle_test.go:487:func Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:523:func Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:585:func Test_FLOW_023_AddedIDRequiresMappedFailingCase(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:609:func Test_FLOW_023_ReviewTaggedAdditionNeedsNoFailFirstCase(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:629:func Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:731:func Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:829:func audit(t *testing.T, root, body string) {
packages/workflow/evals/lifecycle_test.go:840:func toAudit(t *testing.T, root string) workflow.State {
packages/workflow/evals/lifecycle_test.go:850:func Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:874:func Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:895:func Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:963:func Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:978:func Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1070:func Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1136:func Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1298:func Test_FLOW_054_StartOperationalPreconditionsHaveP01(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1312:func Test_FLOW_056_StartContentionAndCurrentForeignClaimHaveP02(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1325:func Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1362:func Test_FLOW_059_AuditNoFindingsErrorsHaveP09(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1376:func Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1397:func Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses(t *testing.T) {
packages/workflow/evals/lifecycle_test.go:1495:func Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics(t *testing.T) {
packages/workflow/evals/eligible_test.go:17:func eligibilityBase(t *testing.T, blockedModule string) string {
packages/workflow/evals/eligible_test.go:29:func eligibilityChange(t *testing.T, root string, number int, slug, module, set, iface string) {
packages/workflow/evals/eligible_test.go:34:func checkEligibilityFixture(t *testing.T, root string) {
packages/workflow/evals/eligible_test.go:52:func readyEligibilityFixture(t *testing.T, root string) {
packages/workflow/evals/eligible_test.go:58:func eligibleSlugs(t *testing.T, root string) []string {
packages/workflow/evals/eligible_test.go:74:func assertEligible(t *testing.T, root string, want ...string) {
packages/workflow/evals/eligible_test.go:82:func Test_FLOW_068_EligibleOrderAndEmptyQueues(t *testing.T) {
packages/workflow/evals/eligible_test.go:108:func Test_FLOW_069_EligibleWaitsForSameModule(t *testing.T) {
packages/workflow/evals/eligible_test.go:116:func Test_FLOW_070_EligibleWaitsForDirectAndTransitiveDependents(t *testing.T) {
packages/workflow/evals/eligible_test.go:125:func Test_FLOW_071_EligibleInterfaceChangeWaitsForBlockedDependent(t *testing.T) {
packages/workflow/evals/eligible_test.go:135:func Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet(t *testing.T) {
packages/workflow/evals/eligible_test.go:152:func Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead(t *testing.T) {
packages/workflow/evals/eval_fix_lifecycle_test.go:14:func evalFixFixture(t *testing.T, interfaceBody string) string {
packages/workflow/evals/eval_fix_lifecycle_test.go:54:func correctEvalFix(t *testing.T, root string) {
packages/workflow/evals/eval_fix_lifecycle_test.go:59:func Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits(t *testing.T) {
packages/workflow/evals/eval_fix_lifecycle_test.go:106:func Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals(t *testing.T) {
packages/workflow/evals/eval_fix_lifecycle_test.go:134:func Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks(t *testing.T) {
packages/workflow/evals/eval_fix_lifecycle_test.go:183:func Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit(t *testing.T) {
packages/workflow/evals/resume_test.go:16:func unblockAttempt(t *testing.T, root string) {
packages/workflow/evals/resume_test.go:37:func savedEvalsAttempt(t *testing.T) (string, workflow.State, workflow.BlockResult) {
packages/workflow/evals/resume_test.go:74:func assertNoResumeMutation(t *testing.T, root, beforeHead, beforeChange string) {
packages/workflow/evals/resume_test.go:87:func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptRestoresEvalsStepAndRetainsCurrentApproval(t *testing.T) {
packages/workflow/evals/resume_test.go:149:func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance(t *testing.T) {
packages/workflow/evals/resume_test.go:221:func Test_FLOW_074_StartFromAttemptSelectsRequestedSavedRef(t *testing.T) {
packages/workflow/evals/resume_test.go:249:func Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval(t *testing.T) {
packages/workflow/evals/resume_test.go:307:func Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange(t *testing.T) {
packages/workflow/evals/resume_test.go:357:func Test_FLOW_076_StartFromAttemptRejectsInvalidOrMissingAttemptBeforeMutation(t *testing.T) {
packages/workflow/evals/block_test.go:14:func block(t *testing.T, root, reason string) workflow.BlockResult {
packages/workflow/evals/block_test.go:23:func blockFixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
packages/workflow/evals/block_test.go:31:func readFile(t *testing.T, root, name string) string {
packages/workflow/evals/block_test.go:40:func assertBlockDocumentOnlyAddsReason(t *testing.T, approved, blocked, reason string) {
packages/workflow/evals/block_test.go:66:func assertBlockReleased(t *testing.T, root string, r workflow.BlockResult, reason string) {
packages/workflow/evals/block_test.go:88:func Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState(t *testing.T) {
packages/workflow/evals/block_test.go:167:func Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit(t *testing.T) {
packages/workflow/evals/block_test.go:209:func Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead(t *testing.T) {
packages/workflow/evals/gc_test.go:20:func gcCall(t *testing.T, root string, now time.Time) workflow.GCResult {
packages/workflow/evals/gc_test.go:29:func gcRef(t *testing.T, root, name, oid string) {
packages/workflow/evals/gc_test.go:36:func gcClaim(t *testing.T, root, name, worktree string) {
packages/workflow/evals/gc_test.go:49:func gcPresent(t *testing.T, root, name string, want bool) {
packages/workflow/evals/gc_test.go:57:func gcNames(t *testing.T, root string) []string {
packages/workflow/evals/gc_test.go:66:func gcPruned(t *testing.T, got, want []string) {
packages/workflow/evals/gc_test.go:74:func Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly(t *testing.T) {
packages/workflow/evals/gc_test.go:109:func Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts(t *testing.T) {
packages/workflow/evals/gc_test.go:128:func Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate(t *testing.T) {
packages/workflow/evals/gc_test.go:164:func Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning(t *testing.T) {
packages/workflow/evals/gc_test.go:185:func Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState(t *testing.T) {
packages/workflow/evals/regenerate_start_test.go:18:func regenerateFixture(t *testing.T, configExtra string) string {
packages/workflow/evals/regenerate_start_test.go:44:func archivedImpl(t *testing.T, root, name, module, date string) {
packages/workflow/evals/regenerate_start_test.go:76:func Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive(t *testing.T) {
packages/workflow/evals/regenerate_start_test.go:161:func Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate(t *testing.T) {
packages/workflow/evals/regenerate_start_test.go:201:func recordedRegenerateFixture(t *testing.T) string {
packages/workflow/evals/regenerate_start_test.go:216:func Test_FLOW_081_RecordedRegenerateStartRetainsDisposableFilesAndState(t *testing.T) {
packages/workflow/evals/regenerate_start_test.go:231:func Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook(t *testing.T) {

 succeeded in 557ms:
# Workflow reconstruction

Reconstruct workflow from the supplied retained workflow SPEC (including Decisions), interface/model and public evals. Only the declared contracts, gitstore, assets, project and checks public api.go/model interfaces are provided as dependency inputs.

Write only packages/workflow/internal/** and DRILL_NOTES.md. All supplied inputs are immutable. No go.mod, replacement dependency, hook, manifest, tests or root setup may be added. Do not run local builds, tests, or module setup, including temporary module setup. Missing compilation infrastructure is intentional. Format only the fresh code you write, then return it to the separate outside verifier; that verifier runs full builds, own evals and all transitive dependent evals.

Read only this supplied input tree and the fresh disposable code you write here. Its Git history is input-only. Do not read original/deleted implementation, old history, refs/attempts/proofs/reports/reconstructions, bootstrap or rebuild-v1, outside repositories, verifier paths, binaries, or source caches. Do not search outside this tree, use controller queries as an oracle, run other agents, or contact the network. Do not infer any permission from a visible outside path or environment variable.

Do not commit, change Git refs/history/config, invoke lifecycle mutators, modify retained inputs, or weaken any verification requirement. Return to the outside controller after reconstruction; it owns all transitions and commits. Use the configured workspace-write command and model/effort only.

The current tree's .fireweed/bin is first on PATH and GOCACHE=/tmp/fireweed-rebuild-cache. Do not invoke any existing Fireweed binary to obtain behavior or implementation. No controller/model fallback is permitted. Record contradicted Decisions and targeted gaps in DRILL_NOTES.md, naming the nearest spec ID (or no spec) and owning module. A gap authorizes no correction of retained inputs.
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
total 8
drwx------@    5 joel  staff     160 Nov  8 15:37 .
drwx------@ 5528 joel  staff  176896 Nov  8 15:39 ..
drwxr-xr-x@    9 joel  staff     288 Nov  8 15:39 .git
-rw-r--r--@    1 joel  staff    1950 Nov  8 15:37 SYSTEM.md
drwxr-xr-x@    8 joel  staff     256 Nov  8 15:37 packages
total 56
drwxr-xr-x@  6 joel  staff    192 Nov  8 15:37 .
drwxr-xr-x@  8 joel  staff    256 Nov  8 15:37 ..
-rw-r--r--@  1 joel  staff  23549 Nov  8 15:37 SPEC.md
-rw-r--r--@  1 joel  staff   2289 Nov  8 15:37 api.go
drwxr-xr-x@ 19 joel  staff    608 Nov  8 15:37 evals
drwxr-xr-x@  3 joel  staff     96 Nov  8 15:37 model

exec
/bin/zsh -lc "sed -n '1136,1600p' packages/workflow/evals/lifecycle_test.go; cat packages/workflow/evals/selector_audit_validation_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc "sed -n '1,525p' packages/workflow/evals/lifecycle_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc "sed -n '525,1137p' packages/workflow/evals/lifecycle_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc 'cat packages/workflow/evals/{eligible,eval_fix_lifecycle,maintenance_start,regenerate_start}_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 681ms:
package evals_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const oldClause = "Alpha MUST return the old value."
const newClause = "Alpha MUST return the revised value."

func write(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func git(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.invalid", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.invalid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func head(t *testing.T, root string) string { t.Helper(); return git(t, root, "rev-parse", "HEAD") }
func commit(t *testing.T, root, message string) {
	t.Helper()
	git(t, root, "add", "-A")
	git(t, root, "commit", "-m", message)
}

func fixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
	t.Helper()
	root := t.TempDir()
	var err error
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "SYSTEM.md", "# Fixture system\nLocal lifecycle fixture.\n")
	write(t, root, ".gitignore", ".fireweed/runs/\n")
	write(t, root, "fireweed.yaml", `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
`+configExtra)
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
printf 'fixture-hook-output\n'
module="$1"
report="$2"
if [ "$module" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then result=''; exitcode=0
  else result='<failure message="expected before implementation"/>'; exitcode=1; fi
  if [ -f .fireweed/unrelated ]; then other='<failure message="known unrelated"/>'; exitcode=1
  else other=''; fi
  printf '<testsuite><testcase name="ALP-001" classname="fixture">%s</testcase><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" "$other" > "$report"
  exit "$exitcode"
fi
if [ "$module" = beta ] && [ -f .fireweed/beta-fail ]; then result='<failure/>'; exitcode=1
elif [ "$module" = gamma ] && [ -f .fireweed/gamma-fail ]; then result='<failure/>'; exitcode=1
else result=''; exitcode=0; fi
case "$module" in beta) id=BET-001;; gamma) id=GAM-001;; *) exit 2;; esac
printf '<testsuite><testcase name="%s" classname="fixture">%s</testcase></testsuite>' "$id" "$result" > "$report"
exit "$exitcode"
`)
	write(t, root, ".fireweed/boundaries.sh", `#!/bin/sh
if [ -f .fireweed/boundary-fail ] && [ -f packages/alpha/internal/ready ]; then printf 'alpha: forbidden dependency\n'; exit 1; fi
`)
	write(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\n"+oldClause+"\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n")
	write(t, root, "packages/alpha/api.go", "package alpha\n")
	write(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n")
	write(t, root, "packages/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: [beta]\n---\n# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n")
	if modified == "" {
		modified = "### ALP-001\n" + newClause + "\n"
	}
	base := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: oldClause + "\n"})
	if removed != "" {
		base += "\n  ALP-002: " + contracts.SpecHash(contracts.Item{ID: "ALP-002", Body: "Alpha MUST retain the second behavior.\n"})
	}
	write(t, root, "changes/0001-manual.md", "---\nmodule: alpha\nbases:\n  ALP-001: "+base+"\n"+changeFront+"---\n\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n")
	git(t, root, "init")
	git(t, root, "checkout", "-b", "main")
	commit(t, root, "approved fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid lifecycle fixture: %#v", d)
		}
	}
	return root
}

func start(t *testing.T, root string) workflow.State {
	t.Helper()
	s, err := workflow.Start(root, "manual")
	if err != nil {
		t.Fatalf("Start: %v, diagnostics: %#v", err, s.Diagnostics)
	}
	return s
}

func request(root string) workflow.Request {
	return workflow.Request{Root: root, Change: "manual", Bin: "/tmp/fireweed-fixture", Timeout: 2 * time.Second}
}
func verify(root string, quick bool) (workflow.Result, error) {
	r := request(root)
	r.Quick = quick
	return workflow.Verify(context.Background(), r)
}
func done(root string) (workflow.Result, error) {
	return workflow.Done(context.Background(), request(root))
}

func containsCode(ds []project.Diagnostic, code string) bool {
	for _, d := range ds {
		if d.Code == code {
			return true
		}
	}
	return false
}

func Test_FLOW_001_FLOW_006_FLOW_007_FLOW_008_FLOW_012_FLOW_013_FLOW_014_FLOW_017_FLOW_018_FLOW_019_FLOW_053_StartAssignsLateIDsAndRetainsSnapshotWithoutCommit(t *testing.T) {
	root := fixture(t, "", "", "### new:first\nFirst MUST work.\n\n### new:second\nSecond MUST use new:first.\n", "", "")
	before := head(t, root)
	s := start(t, root+"/.")
	if s.Phase != "evals" || s.Run == "" || s.ClaimOID == "" || s.PhaseBase != before || head(t, root) != before {
		t.Errorf("start state = %#v; head before %s after %s", s, before, head(t, root))
	}
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if strings.Contains(string(b), "new:first") || strings.Contains(string(b), "new:second") || !strings.Contains(string(b), "### ALP-003") || !strings.Contains(string(b), "### ALP-004") || !strings.Contains(string(b), "use ALP-003") {
		t.Errorf("late IDs not assigned throughout change: %s", b)
	}
	if len(s.Change.Added) != 2 || s.Change.Added[0].ID != "ALP-003" || s.Change.Added[1].ID != "ALP-004" {
		t.Errorf("returned Added IDs = %#v", s.Change.Added)
	}
	alphaProjected := ""
	for _, module := range s.Snapshot.Modules {
		if module.Name == "alpha" {
			alphaProjected = module.Document.Body
		}
	}
	if !strings.Contains(alphaProjected, newClause) || !strings.Contains(alphaProjected, "First MUST work.") || !strings.Contains(alphaProjected, "### ALP-003") || !strings.Contains(alphaProjected, "### ALP-004") {
		t.Errorf("returned snapshot does not project the approved change: %q", alphaProjected)
	}
	edited := strings.Replace(string(b), "First MUST work.", "First MUST change after approval.", 1)
	if edited == string(b) {
		t.Fatal("fixture change has no approved Added body to edit")
	}
	write(t, root, "changes/0001-manual.md", edited)
	current, err := workflow.Current(root, "")
	if err != nil || current.Run != s.Run || current.Change.Slug != "manual" || current.Snapshot.Root != root {
		t.Errorf("Current = %#v, %v", current, err)
	}
	currentAlpha := ""
	for _, module := range current.Snapshot.Modules {
		if module.Name == "alpha" {
			currentAlpha = module.Document.Body
		}
	}
	if currentAlpha != alphaProjected {
		t.Errorf("Current snapshot changed after editing the held change: got %q, approved %q", currentAlpha, alphaProjected)
	}
	if s.Snapshot.Root != root || len(s.Snapshot.Modules) != 3 {
		t.Errorf("returned projected snapshot = %#v", s.Snapshot)
	}
	second := fixture(t, "", "", "", "", "")
	if another := start(t, second); another.Run == "" || another.Run == s.Run {
		t.Errorf("separate starts returned runs %q and %q", s.Run, another.Run)
	}
}

func Test_FLOW_007_StartReservesCurrentAndPreviouslyRemovedIDs(t *testing.T) {
	root := fixture(t, "", "", "### new:reserved\nNew new:reserved behavior MUST work.\n", "", "### ALP-002\n")
	write(t, root, "changes/archive/0000-prior.md", `---
module: alpha
phase: applied
---

## Motivation
Retire the earlier third behavior.

## Interface

## Added

## Modified

## Removed
### ALP-003

## Decisions

`)
	change := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(change, "Alpha public value is revised.", "Alpha public value is revised for new:reserved.", 1))
	commit(t, root, "approved reserved-ID fixture")

	s := start(t, root)
	if len(s.Change.Added) != 1 || s.Change.Added[0].ID != "ALP-004" {
		t.Fatalf("Start assigned Added = %#v; want ALP-004 after current ALP-001/002 and archived removal ALP-003", s.Change.Added)
	}
	if len(s.Change.Removed) != 1 || s.Change.Removed[0].ID != "ALP-002" {
		t.Fatalf("Start lost the removal of current ALP-002: %#v", s.Change.Removed)
	}
	recorded := readFile(t, root, "changes/0001-manual.md")
	for _, want := range []string{"### ALP-004", "New ALP-004 behavior MUST work.", "revised for ALP-004.", "### ALP-002"} {
		if !strings.Contains(recorded, want) {
			t.Errorf("recorded change lacks %q: %s", want, recorded)
		}
	}
	if strings.Contains(recorded, "new:reserved") {
		t.Errorf("Start left an Added handle in the recorded change: %s", recorded)
	}
}

func Test_FLOW_010_FLOW_011_RequiresCleanEligibleQueueHead(t *testing.T) {
	clean := fixture(t, "", "", "", "", "")
	if _, err := workflow.Start(clean, "manual"); err != nil {
		t.Errorf("clean eligible queue head rejected: %v", err)
	}
	for _, dirty := range []string{"tracked", "untracked"} {
		t.Run(dirty, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			if dirty == "tracked" {
				write(t, root, "SYSTEM.md", "edited\n")
			} else {
				write(t, root, "stray.txt", "untracked\n")
			}
			if _, err := workflow.Start(root, "manual"); err == nil {
				t.Error("dirty repository claimed")
			}
		})
	}
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/0000-first.md", "---\nmodule: alpha\n---\n## Motivation\nFirst request.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	commit(t, root, "earlier queued change")
	if _, err := workflow.Start(root, "manual"); err == nil {
		t.Error("later queued change claimed")
	}
}

func Test_FLOW_001_FLOW_016_FLOW_017_LinkedWorktreeExcludesCompetingOwner(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	type attempt struct {
		root  string
		state workflow.State
		err   error
	}
	gate := make(chan struct{})
	results := make(chan attempt, 2)
	for _, dir := range []string{root, other} {
		go func(dir string) {
			<-gate
			s, err := workflow.Start(dir, "manual")
			results <- attempt{dir, s, err}
		}(dir)
	}
	close(gate)
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	var attempts [2]attempt
	for i := range attempts {
		select {
		case attempts[i] = <-results:
		case <-timer.C:
			t.Fatal("simultaneous claim attempts did not finish")
		}
	}
	winners := 0
	for _, result := range attempts {
		if result.err == nil {
			winners++
			if got, err := workflow.Current(result.root, "manual"); err != nil || got.ClaimOID != result.state.ClaimOID {
				t.Errorf("winner lost claim: %#v, %v", got, err)
			}
		} else if _, err := workflow.Current(result.root, "manual"); err == nil {
			t.Errorf("losing worktree %s resolved foreign claim", result.root)
		}
	}
	if winners != 1 {
		t.Errorf("simultaneous claims had %d winners: %#v", winners, attempts)
	}
}

func Test_FLOW_009_FLOW_015_ApprovalWarningAndSharedUnapprovedLimit(t *testing.T) {
	approved := fixture(t, "", "", "", "", "")
	if s := start(t, approved); containsCode(s.Diagnostics, "P04") {
		t.Error("local-base approval warned")
	}
	shared := fixture(t, "", "", "", "", "")
	sharedBase := head(t, shared)
	git(t, shared, "checkout", "-b", "feature")
	git(t, shared, "checkout", "main")
	git(t, shared, "rm", "changes/0001-manual.md")
	commit(t, shared, "local base lacks pending change")
	git(t, shared, "update-ref", "refs/remotes/origin/main", sharedBase)
	git(t, shared, "checkout", "feature")
	if s := start(t, shared); containsCode(s.Diagnostics, "P04") {
		t.Error("shared-base approval warned")
	}
	for _, tc := range []struct {
		name   string
		limit  int
		config string
		reject bool
	}{
		{name: "default-below", limit: 4},
		{name: "default-at", limit: 5, reject: true},
		{name: "configured-at", limit: 2, config: "authoring:\n  max_unapproved: 2\n", reject: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, tc.config, "", "", "", "")
			git(t, root, "checkout", "-b", "feature")
			git(t, root, "checkout", "main")
			git(t, root, "rm", "changes/0001-manual.md")
			commit(t, root, "base without pending change")
			git(t, root, "update-ref", "refs/remotes/origin/main", head(t, root))
			git(t, root, "checkout", "feature")
			for i := 0; i < tc.limit; i++ {
				write(t, root, fmt.Sprintf("changes/archive/%04d-prior.md", i+1), fmt.Sprintf("---\nmodule: alpha\nphase: applied\n---\n## Motivation\nPrior %d.\n", i))
				commit(t, root, "archived prior")
			}
			before := head(t, root)
			s, err := workflow.Start(root, "manual")
			if tc.reject {
				if err == nil || !containsCode(s.Diagnostics, "P07") || head(t, root) != before {
					t.Errorf("shared max bound = %#v, %v", s, err)
				}
			}
			if !tc.reject {
				if err != nil || !containsCode(s.Diagnostics, "P04") {
					t.Errorf("unapproved below bound = %#v, %v", s, err)
				}
				for _, d := range s.Diagnostics {
					if d.Code == "P04" && !d.Warning {
						t.Errorf("P04 diagnostic is not a warning: %#v", d)
					}
				}
			}
		})
	}
}

func Test_FLOW_055_VerifyAndDoneRejectUnsupportedHeldPhase(t *testing.T) {
	root := fixture(t, "", "phase: review\n", "", "", "")
	start(t, root)
	before := head(t, root)
	if _, err := verify(root, true); err == nil || !strings.HasPrefix(err.Error(), "P01: unsupported phase review") {
		t.Errorf("unsupported Verify phase = %v", err)
	}
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P01: unsupported phase review") || head(t, root) != before {
		t.Errorf("unsupported Done phase = %v", err)
	}
}

func Test_FLOW_055_UnsupportedPhasePrecedesProtectedEdit(t *testing.T) {
	root := fixture(t, "protected: [packages/alpha/internal/protected.txt]\n", "phase: review\n", "", "", "")
	started := start(t, root)
	if started.Phase != "review" {
		t.Fatalf("held phase = %q, want review", started.Phase)
	}
	before := head(t, root)
	write(t, root, "packages/alpha/internal/protected.txt", "protected worktree edit\n")
	for _, operation := range []struct {
		name string
		call func() (workflow.Result, error)
	}{
		{"Verify", func() (workflow.Result, error) { return verify(root, true) }},
		{"Done", func() (workflow.Result, error) { return done(root) }},
	} {
		if _, err := operation.call(); err == nil || !strings.HasPrefix(err.Error(), "P01: unsupported phase review") {
			t.Errorf("%s with protected edit = %v, want P01 unsupported phase", operation.name, err)
		}
		if got := head(t, root); got != before {
			t.Errorf("rejected %s moved HEAD: %s -> %s", operation.name, before, got)
		}
		current, err := workflow.Current(root, "manual")
		if err != nil || current.Phase != "review" || current.PhaseBase != before || current.ClaimOID != started.ClaimOID {
			t.Errorf("rejected %s changed held phase or claim: %#v, %v", operation.name, current, err)
		}
	}
}

func Test_FLOW_056_FLOW_061_CurrentAndContextIdentifyMissingClaimsAndChanges(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02: change has no claim") {
		t.Errorf("unclaimed Current = %v", err)
	}
	if _, err := workflow.Current(root, "absent"); err == nil || !strings.HasPrefix(err.Error(), `U02: change "absent" not found`) {
		t.Errorf("unknown Current change = %v", err)
	}
	if _, err := workflow.Context(root, "absent", false); err == nil || !strings.HasPrefix(err.Error(), `U02: change "absent" not found`) {
		t.Errorf("unknown Context change = %v", err)
	}
}

func Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08(t *testing.T) {
	staleStart := fixture(t, "", "", "", "", "")
	p := filepath.Join(staleStart, "changes/0001-manual.md")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	write(t, staleStart, "changes/0001-manual.md", regexp.MustCompile(`(?m)^  ALP-001: [0-9a-f]+$`).ReplaceAllString(string(b), "  ALP-001: deadbeef"))
	commit(t, staleStart, "stale approved base")
	before := head(t, staleStart)
	if _, err := workflow.Start(staleStart, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P05:") || !strings.Contains(err.Error(), "stale base for ALP-001") || head(t, staleStart) != before {
		t.Errorf("stale Start = %v", err)
	}

	missingBase := fixture(t, "", "", "", "", "")
	p = filepath.Join(missingBase, "changes/0001-manual.md")
	b, err = os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	write(t, missingBase, "changes/0001-manual.md", regexp.MustCompile(`(?m)^  ALP-001: [0-9a-f]+$`).ReplaceAllString(string(b), "  ALP-001: "))
	commit(t, missingBase, "missing approved base")
	if _, err := workflow.Start(missingBase, "manual"); err == nil || !strings.HasPrefix(err.Error(), "C08: missing base for ALP-001") {
		t.Errorf("missing-base static cause = %v", err)
	}

	root := fixture(t, "", "", "", "", "")
	marker := hookMarker(t, root)
	start(t, root)
	write(t, root, "packages/alpha/SPEC.md", strings.Replace(git(t, root, "show", "HEAD:packages/alpha/SPEC.md"), oldClause, "Alpha MUST return a changed base value.", 1)+"\n")
	before = head(t, root)
	if _, err := verify(root, true); err == nil || !strings.HasPrefix(err.Error(), "P05:") || !strings.Contains(err.Error(), "stale base for ALP-001") {
		t.Errorf("stale Verify = %v", err)
	}
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P05:") || !strings.Contains(err.Error(), "stale base for ALP-001") || head(t, root) != before {
		t.Errorf("stale Done = %v", err)
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Errorf("hook ran before stale-base rejection: %v", err)
	}
}

func hookMarker(t *testing.T, root string) string {
	t.Helper()
	marker := filepath.Join(t.TempDir(), "hook-ran")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, ".fireweed/eval.sh", "printf run > "+fmt.Sprintf("%q", marker)+"\n"+string(b))
	commit(t, root, "track hook observer")
	return marker
}

func Test_FLOW_002_FLOW_020_FLOW_021_RejectsAuthorMetadataAndZonesBeforeHooks(t *testing.T) {
	for _, edit := range []string{"author", "eval-zone"} {
		t.Run(edit, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			write(t, root, "packages/alpha/internal/existing.go", "package internal\n")
			commit(t, root, "existing implementation")
			marker := hookMarker(t, root)
			start(t, root)
			if edit == "author" {
				p := filepath.Join(root, "changes/0001-manual.md")
				b, _ := os.ReadFile(p)
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "Users need a revised alpha value.", "Unapproved changed purpose.", 1))
			} else {
				write(t, root, "packages/alpha/internal/existing.go", "package internal // changed in evals\n")
			}
			if r, err := verify(root, true); err == nil {
				t.Errorf("%s accepted: %#v", edit, r)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Errorf("hook ran before %s rejection: %v", edit, err)
			}
			if edit == "author" {
				p := filepath.Join(root, "changes/0001-manual.md")
				b, _ := os.ReadFile(p)
				write(t, root, "changes/0001-manual.md", strings.Replace(strings.Replace(string(b), "Unapproved changed purpose.", "Users need a revised alpha value.", 1), "phase: evals", "phase: impl", 1))
				if r, err := verify(root, true); err == nil {
					t.Errorf("approved phase state alteration accepted: %#v", r)
				}
				if _, err := os.Stat(marker); !os.IsNotExist(err) {
					t.Errorf("hook ran before phase-state rejection: %v", err)
				}
			}
		})
	}
}

func Test_FLOW_022_FLOW_023_FLOW_031_EvalsRequireCoverageAndPerIDFailFirst(t *testing.T) {
	for _, scenario := range []string{"missing-coverage", "passing-target-unrelated-failure", "already-satisfied", "fixes"} {
		t.Run(scenario, func(t *testing.T) {

 succeeded in 681ms:
		t.Run(scenario, func(t *testing.T) {
			front := ""
			if scenario == "already-satisfied" {
				front = "already_satisfied: [ALP-001]\n"
			}
			if scenario == "fixes" {
				front = "fixes: [ALP-002]\n"
			}
			root := fixture(t, "", front, "", "", "")
			if scenario == "missing-coverage" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				write(t, root, ".fireweed/eval.sh", strings.ReplaceAll(string(b), `name="ALP-001"`, `name="ALP-002"`))
				commit(t, root, "hook omits target")
			}
			if scenario == "already-satisfied" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				script := strings.Replace(string(b), `  exit "$exitcode"`, `  if [ -f packages/alpha/evals/no-coverage ]; then sed 's/name="ALP-002"/name="ALP-001"/g' "$report" > "$report.tmp"; mv "$report.tmp" "$report"; fi
  exit "$exitcode"`, 1)
				if script == string(b) {
					t.Fatal("fixture coverage switch was not installed")
				}
				write(t, root, ".fireweed/eval.sh", script)
				commit(t, root, "hook can omit unrelated current spec")
			}
			if scenario == "passing-target-unrelated-failure" || scenario == "already-satisfied" {
				write(t, root, ".fireweed/unrelated", "yes\n")
				commit(t, root, "known unrelated failing case")
			}
			start(t, root)
			if scenario == "passing-target-unrelated-failure" || scenario == "already-satisfied" {
				write(t, root, "packages/alpha/internal/ready", "ready\n")
			}
			before := head(t, root)
			r, err := verify(root, true)
			wantPass := scenario == "already-satisfied"
			if (err == nil) != wantPass {
				t.Errorf("Verify = %#v, %v; want pass %v", r, err, wantPass)
			}
			if head(t, root) != before {
				t.Error("Verify committed")
			}
			if wantPass && (len(r.Reports["alpha"].Cases) == 0 || r.State.Snapshot.Root != root) {
				t.Errorf("reports or held snapshot absent: %#v", r)
			}
			if scenario == "already-satisfied" {
				write(t, root, "packages/alpha/evals/no-coverage", "omit ALP-002\n")
				missing, err := verify(root, true)
				if len(missing.Reports["alpha"].Cases) == 0 {
					t.Errorf("target module evals did not run: %#v, %v", missing, err)
				}
				if err == nil {
					t.Errorf("evals accepted report without unrelated current spec ALP-002: %#v", missing)
				}
			}
		})
	}
}

func Test_FLOW_023_AddedIDRequiresMappedFailingCase(t *testing.T) {
	root := fixture(t, "", "set:\n  interface: [api.go, extra.go]\n", "### new:third\nThird MUST return a value.\n", "", "")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, _ := os.ReadFile(p)
	script := strings.Replace(string(b), `</testsuite>' "$result" "$other"`, `<testcase name="ALP-003" classname="fixture">%s</testcase></testsuite>' "$result" "$other" "$result"`, 1)
	if script == string(b) {
		t.Fatal("fixture report extension was not installed")
	}
	write(t, root, ".fireweed/eval.sh", script)
	commit(t, root, "report projected added ID")
	s := start(t, root)
	if len(s.Change.Added) != 1 || s.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("assigned Added ID = %#v", s.Change.Added)
	}
	write(t, root, "packages/alpha/extra.go", "package alpha\n")
	r, err := verify(root, true)
	if err != nil {
		t.Fatalf("mapped failing Added case rejected: %#v, %v", r, err)
	}
	if len(r.Reports["alpha"].Cases) != 3 {
		t.Errorf("added ID was not reported: %#v", r.Reports["alpha"])
	}
}

func Test_FLOW_023_ReviewTaggedAdditionNeedsNoFailFirstCase(t *testing.T) {
	root := fixture(t, "", "", "### new:review [review]\nReview judgment MUST be recorded.\n", "", "")
	s := start(t, root)
	if s.Phase != "evals" || len(s.Change.Added) != 1 || s.Change.Added[0].ID != "ALP-003" || s.Change.Added[0].Tag != "review" {
		t.Fatalf("fixture lacks an assigned review-tagged change in evals: %#v", s)
	}
	before := head(t, root)
	r, err := verify(root, true)
	if err != nil {
		t.Fatalf("evals Verify rejected review item without a failing case: %#v, %v", r.Reports["alpha"], err)
	}
	if head(t, root) != before {
		t.Error("Verify committed the review fixture")
	}
	cases := r.Reports["alpha"].Cases
	if len(cases) != 2 || cases[0].Result != "fail" || !reflect.DeepEqual(cases[0].IDs, []string{"ALP-001"}) || cases[1].Result != "pass" || !reflect.DeepEqual(cases[1].IDs, []string{"ALP-002"}) {
		t.Errorf("fixture needs genuine fail-first for enforced ALP-001 and no fabricated review case: %#v", cases)
	}
}

func Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries(t *testing.T) {
	for _, scenario := range []string{"frozen-eval", "direct", "all", "none", "quick-boundary"} {
		t.Run(scenario, func(t *testing.T) {
			config := "audit:\n  enabled: false\n"
			if scenario == "all" {
				config += "verify_dependents: all\n"
			}
			if scenario == "none" {
				config += "verify_dependents: none\n"
			}
			root := fixture(t, config, "", "", "", "")
			if scenario == "frozen-eval" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				script := strings.Replace(string(b), `  exit "$exitcode"`, `  if [ -f packages/alpha/internal/no-coverage ]; then sed 's/name="ALP-001"/name="ALP-002"/g' "$report" > "$report.tmp"; mv "$report.tmp" "$report"; fi
  exit "$exitcode"`, 1)
				write(t, root, ".fireweed/eval.sh", script)
				commit(t, root, "coverage switch fixture")
			}
			switch scenario {
			case "direct":
				write(t, root, ".fireweed/beta-fail", "yes\n")
			case "none":
				write(t, root, ".fireweed/beta-fail", "yes\n")
			case "all":
				write(t, root, ".fireweed/gamma-fail", "yes\n")
			case "quick-boundary":
				write(t, root, ".fireweed/boundary-fail", "yes\n")
			}
			if scenario != "frozen-eval" {
				commit(t, root, "configured failing hook")
			}
			start(t, root)
			if _, err := done(root); err != nil {
				t.Fatalf("evals Done: %v", err)
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.Phase != "impl" {
				t.Fatalf("impl state = %#v, %v", current, err)
			}
			write(t, root, "packages/alpha/internal/ready", "ready\n")
			if scenario == "frozen-eval" {
				write(t, root, "packages/alpha/evals/new_test.go", "package evals\n")
			}
			before := head(t, root)
			quick, quickErr := verify(root, true)
			full, fullErr := verify(root, false)
			if head(t, root) != before {
				t.Error("Verify committed")
			}
			if scenario == "frozen-eval" {
				if quickErr == nil || fullErr == nil {
					t.Errorf("frozen eval edit accepted: %#v %v; %#v %v", quick, quickErr, full, fullErr)
				}
				if err := os.Remove(filepath.Join(root, "packages/alpha/evals/new_test.go")); err != nil {
					t.Fatal(err)
				}
				write(t, root, "packages/alpha/internal/no-coverage", "yes\n")
				if r, err := verify(root, true); err == nil {
					t.Errorf("impl accepted missing current coverage: %#v", r)
				}
			} else {
				if quickErr != nil {
					t.Errorf("quick rejected own passing evals: %v", quickErr)
				}
				wantFullPass := scenario == "none"
				if (fullErr == nil) != wantFullPass {
					t.Errorf("full Verify = %#v, %v; want pass %v", full, fullErr, wantFullPass)
				}
				if scenario == "direct" && len(full.Reports["beta"].Cases) == 0 {
					t.Error("direct dependent report absent")
				}
				if scenario == "direct" && len(full.Reports["gamma"].Cases) != 0 {
					t.Error("default direct verification ran a transitive dependent")
				}
				if scenario == "all" && len(full.Reports["gamma"].Cases) == 0 {
					t.Error("transitive dependent report absent")
				}
				if scenario == "none" && (len(full.Reports["beta"].Cases) != 0 || len(full.Reports["gamma"].Cases) != 0) {
					t.Error("dependent verification ran with verify_dependents: none")
				}
				if scenario == "direct" && len(full.Reports["beta"].Cases) > 0 && full.Reports["beta"].Cases[0].Result != "fail" {
					t.Errorf("selected direct dependent did not report its failing testcase: %#v", full.Reports["beta"])
				}
				if scenario == "all" && len(full.Reports["gamma"].Cases) > 0 && full.Reports["gamma"].Cases[0].Result != "fail" {
					t.Errorf("selected transitive dependent did not report its failing testcase: %#v", full.Reports["gamma"])
				}
				if scenario == "direct" || scenario == "all" || scenario == "quick-boundary" {
					r, err := done(root)
					if err == nil || head(t, root) != before {
						t.Errorf("Done accepted failed full verification or committed: %#v, %v", r, err)
					}
					current, currentErr := workflow.Current(root, "manual")
					if currentErr != nil || current.Phase != "impl" || current.PhaseBase != before {
						t.Errorf("failed full verification advanced phase: %#v, %v", current, currentErr)
					}
				}
			}
		})
	}
}

func Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures(t *testing.T) {
	for _, tc := range []struct {
		name, mode string
		ownPass    bool
		selected   []string
	}{
		{"own fail/direct", "direct", false, []string{"alpha", "beta"}},
		{"own fail/all", "all", false, []string{"alpha", "beta", "gamma"}},
		{"own fail/none", "none", false, []string{"alpha"}},
		{"own pass/all", "all", true, []string{"alpha", "beta", "gamma"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, "audit:\n  enabled: false\nverify_dependents: "+tc.mode+"\n", "", "", "", "")
			invocations := filepath.Join(t.TempDir(), "invocations")
			write(t, root, ".fireweed/eval.sh", fmt.Sprintf("#!/bin/sh\nprintf '%%s\\n' \"$1\" >> %q\n", invocations)+`case "$1" in
  alpha)
    if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
    else failure='<failure message="alpha-failure-sentinel">alpha-failure-sentinel</failure>'; code=1; fi
    printf '<testsuite><testcase name="ALP-001 own" classname="fixture">%s</testcase><testcase name="ALP-002 retained" classname="fixture"/></testsuite>' "$failure" > "$2"
    exit "$code"
    ;;
  beta) id=BET-001; sentinel=beta-failure-sentinel ;;
  gamma) id=GAM-001; sentinel=gamma-failure-sentinel ;;
  *) exit 2 ;;
esac
printf '<testsuite><testcase name="%s dependent" classname="fixture"><failure message="%s">%s</failure></testcase></testsuite>' "$id" "$sentinel" "$sentinel" > "$2"
exit 1
`)
			commit(t, root, "approved mapped-failure reports and invocation observer")
			start(t, root)
			if _, err := done(root); err != nil {
				t.Fatalf("fixture evals Done: %v", err)
			}
			held, err := workflow.Current(root, "manual")
			if err != nil || held.Phase != "impl" {
				t.Fatalf("fixture held impl state = %#v, %v", held, err)
			}
			if tc.ownPass {
				write(t, root, "packages/alpha/internal/ready", "ready\n")
			}
			if err := os.WriteFile(invocations, nil, 0644); err != nil {
				t.Fatal(err)
			}
			before := head(t, root)
			result, err := verify(root, false)
			if err == nil {
				t.Error("full impl Verify accepted mapped testcase failures")
			}
			if len(result.Reports) != len(tc.selected) {
				t.Errorf("returned reports = %#v; want exactly %v", result.Reports, tc.selected)
			}
			wantInvocations := map[string]int{}
			for _, name := range tc.selected {
				wantInvocations[name] = 1
				report, ok := result.Reports[name]
				count, id, caseName := 1, "BET-001", "BET-001 dependent"
				if name == "alpha" {
					count, id, caseName = 2, "ALP-001", "ALP-001 own"
				} else if name == "gamma" {
					id, caseName = "GAM-001", "GAM-001 dependent"
				}
				if !ok || len(report.Cases) != count {
					t.Errorf("%s report = %#v, present %t; want %d mapped cases", name, report, ok, count)
					continue
				}
				c := report.Cases[0]
				wantResult, wantExit := "fail", 1
				if name == "alpha" && tc.ownPass {
					wantResult, wantExit = "pass", 0
				}
				if c.Name != caseName || c.Classname != "fixture" || !reflect.DeepEqual(c.IDs, []string{id}) || c.Result != wantResult || report.HookExit != wantExit {
					t.Errorf("%s mapped case/ordinary hook exit = %#v / %d; want %s %s / %d", name, c, report.HookExit, id, wantResult, wantExit)
				}
				if wantResult == "fail" && !strings.Contains(c.Failure, name+"-failure-sentinel") {
					t.Errorf("%s report lost failure text: %#v", name, c)
				}
				if name == "alpha" {
					retained := report.Cases[1]
					if retained.Name != "ALP-002 retained" || !reflect.DeepEqual(retained.IDs, []string{"ALP-002"}) || retained.Result != "pass" {
						t.Errorf("alpha retained mapped passing case = %#v", retained)
					}
				}
			}
			gotInvocations := map[string]int{}
			for _, name := range strings.Fields(readFile(t, "", invocations)) {
				gotInvocations[name]++
			}
			if !reflect.DeepEqual(gotInvocations, wantInvocations) {
				t.Errorf("hook invocations = %v; want only selected modules %v", gotInvocations, wantInvocations)
			}
			current, currentErr := workflow.Current(root, "manual")
			if head(t, root) != before || currentErr != nil || current.Phase != "impl" || current.PhaseBase != held.PhaseBase || current.ClaimOID != held.ClaimOID || current.Run != held.Run {
				t.Errorf("failed Verify changed HEAD or held impl state: before %#v, after %#v, error %v", held, current, currentErr)
			}
		})
	}
}

func audit(t *testing.T, root, body string) {
	t.Helper()
	p := filepath.Join(root, "changes/0001-manual.md")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	text := contracts.ReplaceSection(string(b), "Audit", body)
	write(t, root, "changes/0001-manual.md", text)
}

func toAudit(t *testing.T, root string) workflow.State {
	t.Helper()
	start(t, root)
	r, err := done(root)
	if err != nil || r.State.Phase != "audit" {
		t.Fatalf("evals Done = %#v, %v", r, err)
	}
	return r.State
}

func Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl(t *testing.T) {
	for _, disabled := range []bool{false, true} {
		t.Run(fmt.Sprint(disabled), func(t *testing.T) {
			config := ""
			want := "audit"
			if disabled {
				config = "audit:\n  enabled: false\n"
				want = "impl"
			}
			root := fixture(t, config, "", "", "", "")
			start(t, root)
			before := head(t, root)
			r, err := done(root)
			if err != nil || r.State.Phase != want || head(t, root) == before || r.State.PhaseBase != head(t, root) {
				t.Errorf("Done = %#v, %v; want %s with commit", r, err, want)
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.Phase != want || current.PhaseBase != head(t, root) {
				t.Errorf("phase base not advanced: %#v, %v", current, err)
			}
		})
	}
}

func Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- ALP-999: \"Missing MUST clause.\" — invalid ID\n- ALP-001: \"Alpha MUST do something else.\" — invalid quote\nChecked IDs: ALP-001\nNo findings\n")
	r, err := done(root)
	if err != nil || r.State.Phase != "impl" || r.State.PhaseBase != head(t, root) {
		t.Errorf("invalid findings were retained: %#v, %v", r, err)
	}
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if strings.Contains(string(b), "ALP-999") || strings.Contains(string(b), "something else") || !strings.Contains(string(b), "No findings") {
		t.Errorf("audit normalization = %s", b)
	}
	missing := fixture(t, "", "", "", "", "")
	toAudit(t, missing)
	audit(t, missing, "No findings\n")
	before := head(t, missing)
	if _, err := done(missing); err == nil || head(t, missing) != before {
		t.Error("No findings without checked IDs accepted or committed")
	}
}

func Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound(t *testing.T) {
	root := fixture(t, "", "", "### new:extra [review]\nExtra MUST be visible.\n", "", "")
	toAudit(t, root)
	finding := "- ALP-001: \"" + newClause + "\" — missing assertion\n"
	for round := 1; round <= 2; round++ {
		body := finding
		if round == 2 {
			body += "- ALP-003: \"Extra MUST be visible.\" — new finding on re-audit\n"
		}
		audit(t, root, body)
		r, err := done(root)
		if err != nil || r.State.Phase != "evals" || r.State.PhaseBase != head(t, root) {
			t.Fatalf("finding round %d = %#v, %v", round, r, err)
		}
		current, err := workflow.Current(root, "manual")
		if err != nil || current.Change.Document.Frontmatter["audit_rounds"] != round {
			t.Errorf("persistent rounds %d = %#v, %v", round, current, err)
		}
		if round == 2 {
			b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
			if strings.Contains(string(b), "new finding on re-audit") {
				t.Error("new re-audit ID retained")
			}
		}
		if _, err := done(root); err != nil {
			t.Fatalf("revised evals Done round %d: %v", round, err)
		}
	}
	audit(t, root, finding)
	beforeBlock := head(t, root)
	_, _ = done(root)
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	questions, hasQuestions := contracts.Section(string(b), "Questions")
	if !strings.Contains(string(b), "blocked:") || !hasQuestions || !strings.Contains(questions, newClause) {
		t.Errorf("third finding did not block with question: %s", b)
	}
	if head(t, root) == beforeBlock || !strings.Contains(string(b), "audit_rounds: 2") {
		t.Error("two-round bound was not committed with persistent count")
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
	configured := fixture(t, "audit:\n  max_rounds: 1\n", "", "", "", "")
	toAudit(t, configured)
	audit(t, configured, finding)
	first, err := done(configured)
	if err != nil || first.State.Phase != "evals" || first.State.PhaseBase != head(t, configured) {
		t.Fatalf("configured first finding = %#v, %v", first, err)
	}
	firstCurrent, err := workflow.Current(configured, "manual")
	if err != nil || firstCurrent.Change.Document.Frontmatter["audit_rounds"] != 1 {
		t.Errorf("configured round count = %#v, %v", firstCurrent, err)
	}
	if _, err := done(configured); err != nil {
		t.Fatalf("configured revised evals Done: %v", err)
	}
	audit(t, configured, finding)
	_, _ = done(configured)
	change, err := os.ReadFile(filepath.Join(configured, "changes/0001-manual.md"))
	configuredQuestions, configuredHasQuestions := contracts.Section(string(change), "Questions")
	if err != nil || !strings.Contains(string(change), "blocked:") || !strings.Contains(string(change), "audit_rounds: 1") || !configuredHasQuestions || !strings.Contains(configuredQuestions, newClause) {
		t.Errorf("configured max_rounds did not block after one return: %s, %v", change, err)
	}
	if _, err := workflow.Current(configured, "manual"); err == nil {
		t.Error("configured bound retained claim")
	}
}

func Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- [untestable] ALP-001: \""+newClause+"\" — cannot observe the value\n")
	_, _ = done(root)
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	questions, hasQuestions := contracts.Section(string(b), "Questions")
	if !strings.Contains(string(b), "blocked:") || !hasQuestions || !strings.Contains(questions, newClause) {
		t.Errorf("untestable finding did not block: %s", b)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
}

func Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nverify_dependents: none\n", "set:\n  interface: [api.go, extra.go]\n  hooks:\n    eval: [sh, .fireweed/eval.sh, \"{module}\", \"{report}\"]\n", "### new:extra [review]\nExtra MUST be visible.\n", "", "### ALP-002\n")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, _ := os.ReadFile(p)
	script := strings.Replace(string(b), `<testcase name="ALP-002" classname="fixture">%s</testcase>`, "", 1)
	script = strings.Replace(script, `"$result" "$other" > "$report"`, `"$result" > "$report"`, 1)
	write(t, root, ".fireweed/eval.sh", script)
	write(t, root, ".fireweed/global-eval.sh", "#!/bin/sh\nexit 1\n")
	configPath := filepath.Join(root, "fireweed.yaml")
	config, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "fireweed.yaml", strings.Replace(string(config), "eval: [sh, .fireweed/eval.sh,", "eval: [sh, .fireweed/global-eval.sh,", 1))
	commit(t, root, "report only projected spec IDs")
	start(t, root)
	evalsHead := head(t, root)
	r, err := done(root)
	if err != nil || r.State.Phase != "impl" || head(t, root) == evalsHead || r.State.PhaseBase != head(t, root) {
		t.Fatalf("evals transition = %#v, %v", r, err)
	}
	implBase := head(t, root)
	otherClaim := []byte(`{"worktree":"independent","change":"other"}`)
	otherOID, err := gitstore.WriteBlob(root, otherClaim)
	if err != nil {
		t.Fatal(err)
	}
	otherRef := "refs/fireweed/claims/other"
	if err := gitstore.CASRef(root, otherRef, "", otherOID); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	r, err = done(root)
	if err != nil || r.State.Phase != "applied" || head(t, root) == implBase || r.State.PhaseBase != head(t, root) {
		t.Fatalf("impl transition = %#v, %v", r, err)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending change remains: %v", err)
	}
	archive, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	spec, err := os.ReadFile(filepath.Join(root, "packages/alpha/SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	appliedHead := head(t, root)
	archiveAtHead, err := gitstore.ReadAt(root, appliedHead, "changes/archive/0001-manual.md")
	if err != nil || string(archiveAtHead) != string(archive) {
		t.Errorf("applied HEAD lacks archived change: %q, %v", archiveAtHead, err)
	}
	specAtHead, err := gitstore.ReadAt(root, appliedHead, "packages/alpha/SPEC.md")
	if err != nil || string(specAtHead) != string(spec) {
		t.Errorf("applied HEAD lacks applied spec: %q, %v", specAtHead, err)
	}
	moduleDoc, err := contracts.ParseDocument(string(spec))
	if err != nil {
		t.Fatal(err)
	}
	wantInterface := []any{"api.go", "extra.go"}
	wantHooks := map[string]any{"eval": []any{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}
	if !reflect.DeepEqual(moduleDoc.Frontmatter["interface"], wantInterface) || !reflect.DeepEqual(moduleDoc.Frontmatter["hooks"], wantHooks) {
		t.Errorf("set values missing from module metadata: %#v", moduleDoc.Frontmatter)
	}
	for _, want := range []string{"phase: applied", "## Provenance"} {
		if !strings.Contains(string(archive), want) {
			t.Errorf("archive misses %q", want)
		}
	}
	for _, want := range []string{newClause, "### ALP-003", "extra.go", "Use the revised policy."} {
		if !strings.Contains(string(spec), want) {
			t.Errorf("applied spec misses %q: %s", want, spec)
		}
	}
	if strings.Contains(string(spec), "### ALP-002") || strings.Contains(string(spec), "Retain old policy.") {
		t.Errorf("removed spec/decision remains: %s", spec)
	}
	if !strings.Contains(string(spec), "- manual: Use the revised policy.") {
		t.Errorf("decision attribution missing: %s", spec)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("applied claim remained held")
	}
	if got, exists, err := gitstore.Ref(root, otherRef); err != nil || !exists || got != otherOID {
		t.Errorf("application changed independent claim ref: %q, %v, %v", got, exists, err)
	}
	if got, err := gitstore.ReadBlob(root, otherOID); err != nil || string(got) != string(otherClaim) {
		t.Errorf("application changed independent claim blob: %q, %v", got, err)
	}
}

func Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	hookPath := filepath.Join(root, ".fireweed/eval.sh")
	hook, err := os.ReadFile(hookPath)
	if err != nil {
		t.Fatal(err)
	}
	inject := `  if [ -f packages/alpha/internal/ready ]; then printf 'invalid spec fixture\n' > packages/alpha/SPEC.md; fi
  exit "$exitcode"`
	patched := strings.Replace(string(hook), `  exit "$exitcode"`, inject, 1)
	if patched == string(hook) {
		t.Fatal("fixture hook injection was not installed")
	}
	write(t, root, ".fireweed/eval.sh", patched)
	commit(t, root, "tracked invalid-spec hook fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	base := head(t, root)
	own, ownErr := verify(root, false)
	if ownErr == nil {
		t.Errorf("impl Verify accepted failing current-module testcase: %#v", own)
	}
	if cases := own.Reports["alpha"].Cases; len(cases) == 0 || cases[0].Result != "fail" || len(cases[0].IDs) == 0 || cases[0].IDs[0] != "ALP-001" {
		t.Errorf("failing current-module testcase absent from Verify report: %#v", own.Reports["alpha"])
	}
	r, err := done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("passing current evals required: %#v, %v", r, err)
	}
	current, currentErr := workflow.Current(root, "manual")
	if currentErr != nil || current.Phase != "impl" || current.PhaseBase != base {
		t.Errorf("failed Done advanced phase: %#v, %v", current, currentErr)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	write(t, root, "packages/alpha/api.go", "package alpha // edited during impl\n")
	r, err = done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("frozen interface edit accepted or committed: %#v, %v", r, err)
	}
	write(t, root, "packages/alpha/api.go", "package alpha\n")
	r, err = done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("failed static validation applied or committed: %#v, %v", r, err)
	}
	invalidSpec, readErr := os.ReadFile(filepath.Join(root, "packages/alpha/SPEC.md"))
	if readErr != nil || string(invalidSpec) != "invalid spec fixture\n" {
		t.Errorf("successful eval hook did not inject invalid spec: %q, %v", invalidSpec, readErr)
	}
	invalid := false
	for _, d := range project.Check(root) {
		invalid = invalid || !d.Warning
	}
	if !invalid {
		t.Error("injected spec passed static validation")
	}
	if _, err := os.Stat(filepath.Join(root, "changes/archive/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("failed static validation created archive: %v", err)
	}
	pending, readErr := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if readErr != nil || !strings.Contains(string(pending), "phase: impl") || strings.Contains(string(pending), "phase: applied") {
		t.Errorf("failed static validation advanced pending phase: %q, %v", pending, readErr)
	}
}

func Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports(t *testing.T) {
	provenanceLines := func(root, path string) []string {

 succeeded in 683ms:
package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// A real Start/Block leaves the first queued change blocked, committed, and
// unclaimed. The other changes are ordinary approved fixture-local queue files.
func eligibilityBase(t *testing.T, blockedModule string) string {
	t.Helper()
	root := blockFixture(t, "", "", "", "", "")
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n\n### DEL-001\nDelta MUST work.\n")
	write(t, root, "changes/0001-manual.md", fmt.Sprintf("---\nmodule: %s\n---\n\n## Motivation\nThe first change blocks.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n", blockedModule))
	commit(t, root, "blocked change fixture")
	checkEligibilityFixture(t, root)
	start(t, root)
	block(t, root, "fixture blocked")
	return root
}

func eligibilityChange(t *testing.T, root string, number int, slug, module, set, iface string) {
	t.Helper()
	write(t, root, fmt.Sprintf("changes/%04d-%s.md", number, slug), fmt.Sprintf("---\nmodule: %s\n%s---\n\n## Motivation\nThe queued change is needed.\n\n## Interface\n%s\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n", module, set, iface))
}

func checkEligibilityFixture(t *testing.T, root string) {
	t.Helper()
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid eligibility fixture: %s %s: %s", d.Code, d.Path, d.Message)
		}
	}
	queue, err := project.Queue(root)
	if err != nil {
		t.Fatalf("fixture Queue: %v", err)
	}
	for _, ch := range queue {
		if _, err := project.Project(root, ch.Slug); err != nil {
			t.Fatalf("fixture projection through %s: %v", ch.Slug, err)
		}
	}
}

func readyEligibilityFixture(t *testing.T, root string) {
	t.Helper()
	checkEligibilityFixture(t, root)
	commit(t, root, "approved eligibility queue")
}

func eligibleSlugs(t *testing.T, root string) []string {
	t.Helper()
	changes, err := workflow.Eligible(root)
	if err != nil {
		t.Fatalf("Eligible: %v", err)
	}
	slugs := make([]string, 0, len(changes))
	for _, ch := range changes {
		if ch.Path == "" || ch.Module == "" {
			t.Fatalf("Eligible returned incomplete public change: %#v", ch)
		}
		slugs = append(slugs, ch.Slug)
	}
	return slugs
}

func assertEligible(t *testing.T, root string, want ...string) {
	t.Helper()
	want = append([]string{}, want...)
	if got := eligibleSlugs(t, root); !reflect.DeepEqual(got, want) {
		t.Errorf("eligible slugs = %v, want %v", got, want)
	}
}

func Test_FLOW_068_EligibleOrderAndEmptyQueues(t *testing.T) {
	t.Run("skips waiting and retains independent order", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		eligibilityChange(t, root, 3, "independent", "delta", "", "")
		eligibilityChange(t, root, 4, "transitive-wait", "gamma", "", "")
		eligibilityChange(t, root, 5, "second-independent", "delta", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "independent", "second-independent")
	})
	t.Run("empty queue", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
			t.Fatal(err)
		}
		readyEligibilityFixture(t, root)
		assertEligible(t, root)
	})
	t.Run("all waiting", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root)
	})
}

func Test_FLOW_069_EligibleWaitsForSameModule(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "same-module", "alpha", "", "")
	eligibilityChange(t, root, 3, "independent", "delta", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "independent")
}

func Test_FLOW_070_EligibleWaitsForDirectAndTransitiveDependents(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "direct", "beta", "", "")
	eligibilityChange(t, root, 3, "transitive", "gamma", "", "")
	eligibilityChange(t, root, 4, "independent", "delta", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "independent")
}

func Test_FLOW_071_EligibleInterfaceChangeWaitsForBlockedDependent(t *testing.T) {
	root := eligibilityBase(t, "gamma")
	eligibilityChange(t, root, 2, "direct-interface", "beta", "", "Public beta API changes.\n")
	eligibilityChange(t, root, 3, "direct-internal", "beta", "", "")
	eligibilityChange(t, root, 4, "transitive-interface", "alpha", "", "Public alpha API changes.\n")
	eligibilityChange(t, root, 5, "transitive-internal", "alpha", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "direct-internal", "transitive-internal")
}

func Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet(t *testing.T) {
	t.Run("preceding set removes dependency", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "detach-beta", "beta", "set:\n  depends_on: []\n", "")
		eligibilityChange(t, root, 3, "gamma-after-detach", "gamma", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "detach-beta", "gamma-after-detach")
	})
	t.Run("own set adds or removes dependency", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "delta-now-dependent", "delta", "set:\n  depends_on: [alpha]\n", "")
		eligibilityChange(t, root, 3, "gamma-now-independent", "gamma", "set:\n  depends_on: []\n", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "gamma-now-independent")
	})
}

func Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead(t *testing.T) {
	makeQueue := func(t *testing.T) string {
		t.Helper()
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		eligibilityChange(t, root, 3, "first-eligible", "delta", "", "")
		eligibilityChange(t, root, 4, "later-eligible", "delta", "", "")
		readyEligibilityFixture(t, root)
		return root
	}
	t.Run("empty target selects first eligible", func(t *testing.T) {
		root := makeQueue(t)
		s, err := workflow.Start(root, "")
		if err != nil || s.Change.Slug != "first-eligible" {
			t.Errorf("Start(empty) selected %q: %v", s.Change.Slug, err)
		}
	})
	t.Run("explicit first eligible", func(t *testing.T) {
		root := makeQueue(t)
		s, err := workflow.Start(root, "first-eligible")
		if err != nil || s.Change.Slug != "first-eligible" {
			t.Errorf("Start(first-eligible) selected %q: %v", s.Change.Slug, err)
		}
	})
	for _, target := range []string{"waiting", "later-eligible"} {
		t.Run("reject "+target, func(t *testing.T) {
			root := makeQueue(t)
			if _, err := workflow.Start(root, target); err == nil || !strings.HasPrefix(err.Error(), "P01") {
				t.Errorf("Start(%s) = %v, want P01 rejection", target, err)
			}
		})
	}
	t.Run("no eligible change", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		readyEligibilityFixture(t, root)
		if _, err := workflow.Start(root, ""); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("Start(empty) = %v, want P01 rejection", err)
		}
	})
}
package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

// The approved hook reports the assertion in alpha_test.go as a real testcase.
// The initially wrong assertion fails although alpha's implementation is ready.
func evalFixFixture(t *testing.T, interfaceBody string) string {
	t.Helper()
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/0001-manual.md", `---
module: alpha
eval_fix: [ALP-001]
---

## Motivation
Correct the wrong ALP-001 eval while retaining the contract and implementation.

## Interface
`+interfaceBody+`
## Added

## Modified

## Removed

## Decisions

`)
	write(t, root, "packages/alpha/internal/ready", "existing implementation\n")
	write(t, root, "packages/alpha/evals/alpha_test.go", "package alpha_test\n\n// wrong assertion\n")
	hook := readFile(t, root, ".fireweed/eval.sh")
	want := "if [ -f packages/alpha/internal/ready ]; then"
	replacement := "if [ -f packages/alpha/internal/ready ] && grep -q 'correct assertion' packages/alpha/evals/alpha_test.go; then"
	if !strings.Contains(hook, want) {
		t.Fatal("fixture hook condition missing")
	}
	write(t, root, ".fireweed/eval.sh", strings.Replace(hook, want, replacement, 1))
	commit(t, root, "approved eval_fix fixture")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid eval_fix fixture: %#v", d)
		}
	}
	return root
}

func correctEvalFix(t *testing.T, root string) {
	t.Helper()
	write(t, root, "packages/alpha/evals/alpha_test.go", "package alpha_test\n\n// correct assertion\n")
}

func Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits(t *testing.T) {
	root := evalFixFixture(t, "\n")
	retained := map[string]string{}
	for _, name := range []string{"packages/alpha/SPEC.md", "packages/alpha/api.go", "packages/alpha/internal/ready"} {
		retained[name] = readFile(t, root, name)
	}
	approved := head(t, root)
	s := start(t, root)
	if s.Phase != "evals" || s.PhaseBase != approved || head(t, root) != approved {
		t.Fatalf("eval_fix Start = %#v, HEAD %s", s, head(t, root))
	}
	correctEvalFix(t, root)
	checked, err := verify(root, true)
	if err != nil || len(checked.Reports["alpha"].Cases) != 2 || head(t, root) != approved {
		t.Fatalf("corrected eval Verify = %#v, %v; HEAD %s", checked, err, head(t, root))
	}
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" || head(t, root) == approved {
		t.Fatalf("eval_fix evals Done = %#v, %v", evals, err)
	}
	if got := readFile(t, root, "packages/alpha/evals/alpha_test.go"); !strings.Contains(got, "correct assertion") {
		t.Errorf("corrected eval was not committed: %q", got)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	impl, err := done(root)
	if err != nil || impl.State.Phase != "impl" {
		t.Fatalf("eval_fix audit Done = %#v, %v", impl, err)
	}
	beforeImpl := head(t, root)
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" || head(t, root) == beforeImpl {
		t.Fatalf("eval_fix impl Done = %#v, %v", applied, err)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending eval_fix change remains: %v", err)
	}
	archive := readFile(t, root, "changes/archive/0001-manual.md")
	if !strings.Contains(archive, "phase: applied") || !strings.Contains(archive, "No findings; checked ALP-001") {
		t.Errorf("archive lacks applied phase or checked ID: %s", archive)
	}
	for name, want := range retained {
		if got := readFile(t, root, name); got != want {
			t.Errorf("eval_fix changed retained %s: %q", name, got)
		}
	}
}

func Test_FLOW_079_EvalFixRejectsFailingAndMissingListedEvals(t *testing.T) {
	for _, scenario := range []string{"failing", "missing-listed-coverage"} {
		t.Run(scenario, func(t *testing.T) {
			root := evalFixFixture(t, "\n")
			if scenario == "missing-listed-coverage" {
				correctEvalFix(t, root)
				hook := readFile(t, root, ".fireweed/eval.sh")
				// The hook remains approved fixture infrastructure; only its report
				// selection is changed before Start.
				write(t, root, ".fireweed/eval.sh", strings.ReplaceAll(hook, `name="ALP-001"`, `name="ALP-002"`))
				commit(t, root, "approved hook omits listed ID")
			}
			start(t, root)
			before := head(t, root)
			r, err := verify(root, true)
			if err == nil || head(t, root) != before || len(r.Reports["alpha"].Cases) == 0 {
				t.Fatalf("%s Verify: err=%v, target cases=%#v, HEAD=%s want %s", scenario, err, r.Reports["alpha"].Cases, head(t, root), before)
			}
			if scenario == "failing" && r.Reports["alpha"].Cases[0].Result != "fail" {
				t.Errorf("fixture did not produce a failing target eval: %#v", r.Reports["alpha"])
			}
			if _, err := done(root); err == nil || head(t, root) != before {
				t.Errorf("%s Done advanced: %v", scenario, err)
			}
		})
	}
}

func Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks(t *testing.T) {
	t.Run("evals interface", func(t *testing.T) {
		root := evalFixFixture(t, "\nThe public alpha interface may be revised.\n")
		marker := hookMarker(t, root)
		start(t, root)
		correctEvalFix(t, root)
		write(t, root, "packages/alpha/api.go", "package alpha // revised public interface\n")
		before := head(t, root)
		if _, err := verify(root, true); err == nil || head(t, root) != before {
			t.Errorf("eval_fix interface Verify = %v", err)
		}
		if _, err := done(root); err == nil || head(t, root) != before {
			t.Errorf("eval_fix interface Done = %v", err)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Errorf("hook ran before evals interface rejection: %v", err)
		}
	})
	t.Run("impl implementation", func(t *testing.T) {
		root := evalFixFixture(t, "\n")
		marker := hookMarker(t, root)
		start(t, root)
		correctEvalFix(t, root)
		evals, err := done(root)
		if err != nil || evals.State.Phase != "audit" {
			t.Fatalf("evals Done = %#v, %v", evals, err)
		}
		audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
		impl, err := done(root)
		if err != nil || impl.State.Phase != "impl" {
			t.Fatalf("audit Done = %#v, %v", impl, err)
		}
		if err := os.Remove(marker); err != nil {
			t.Fatal(err)
		}
		write(t, root, "packages/alpha/internal/ready", "unapproved implementation change\n")
		before := head(t, root)
		if _, err := verify(root, true); err == nil || head(t, root) != before {
			t.Errorf("eval_fix impl Verify = %v", err)
		}
		if _, err := done(root); err == nil || head(t, root) != before {
			t.Errorf("eval_fix impl Done = %v", err)
		}
		if _, err := os.Stat(marker); !os.IsNotExist(err) {
			t.Errorf("hook ran before impl implementation rejection: %v", err)
		}
	})
}

func Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit(t *testing.T) {
	root := evalFixFixture(t, "\n")
	start(t, root)
	correctEvalFix(t, root)
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-002\nNo findings\n")
	before := head(t, root)
	if _, err := done(root); err == nil || head(t, root) != before {
		t.Errorf("incomplete eval_fix audit advanced: %v", err)
	}
}
package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/0001-manual.md", `---
module: alpha
bases: {}
---

## Motivation
Simplify the existing alpha implementation while keeping its behavior.

## Interface

## Added

## Modified

## Removed

## Decisions

`)
	write(t, root, "packages/alpha/internal/ready", "original implementation\n")
	write(t, root, "packages/alpha/evals/alpha_test.go", "package alpha_test\n\n// Retained coverage fixture.\n")
	marker := hookMarker(t, root)
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid maintenance fixture: %#v", d)
		}
	}
	retained := map[string]string{}
	for _, name := range []string{"packages/alpha/SPEC.md", "packages/alpha/api.go", "packages/alpha/evals/alpha_test.go"} {
		b, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			t.Fatal(err)
		}
		retained[name] = string(b)
	}
	before := head(t, root)
	s := start(t, root)
	if s.Phase != "impl" || s.Run == "" || s.ClaimOID == "" || s.PhaseBase != before || head(t, root) != before {
		t.Fatalf("maintenance Start state = %#v; HEAD before %s after %s", s, before, head(t, root))
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Errorf("Start ran an eval hook: %v", err)
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" || current.Run != s.Run || current.ClaimOID != s.ClaimOID {
		t.Fatalf("persisted maintenance state = %#v, %v", current, err)
	}
	pending, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(pending), "phase: impl") {
		t.Fatalf("pending maintenance phase = %q, %v", pending, err)
	}
	write(t, root, "packages/alpha/internal/ready", "simplified implementation\n")
	checked, err := verify(root, false)
	if err != nil {
		t.Fatalf("full maintenance Verify: %v, reports: %#v", err, checked.Reports)
	}
	if len(checked.Reports["alpha"].Cases) != 2 || len(checked.Reports["beta"].Cases) != 1 {
		t.Errorf("full Verify did not run current and direct dependent coverage: %#v", checked.Reports)
	}
	if head(t, root) != before {
		t.Fatal("Verify changed HEAD")
	}
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" {
		t.Fatalf("maintenance Done = %#v, %v", applied, err)
	}
	after := head(t, root)
	if after == before || git(t, root, "rev-parse", "HEAD^") != before {
		t.Errorf("maintenance archive was not one commit: before %s after %s", before, after)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending maintenance change remains: %v", err)
	}
	archive, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(archive), "phase: applied") || !strings.Contains(string(archive), "- impl ") || strings.Contains(string(archive), "- evals ") || strings.Contains(string(archive), "- audit ") {
		t.Errorf("maintenance archive has unexpected phase history: %s", archive)
	}
	for name, want := range retained {
		got, err := os.ReadFile(filepath.Join(root, name))
		if err != nil || string(got) != want {
			t.Errorf("retained %s changed: %q, %v", name, got, err)
		}
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("applied maintenance claim remained held")
	}
}
package evals_test

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// The shared fixture supplies a Modified item by default. Regeneration is a
// maintenance change, so author its complete approved change before committing.
func regenerateFixture(t *testing.T, configExtra string) string {
	t.Helper()
	root := fixture(t, configExtra, "", "", "", "")
	write(t, root, "changes/0001-manual.md", `---
module: alpha
bases: {}
regenerate: true
---

## Motivation
Rebuild alpha's implementation from retained contracts and evals.

## Interface

## Added

## Modified

## Removed

## Decisions

`)
	return root
}

func archivedImpl(t *testing.T, root, name, module, date string) {
	t.Helper()
	playbook, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	hash := sha256.Sum256([]byte(playbook))
	write(t, root, "changes/archive/"+name+".md", fmt.Sprintf(`---
module: %s
bases: {}
run: fixture-%s
phase: applied
---

## Motivation
Earlier implementation change.

## Interface

## Added

## Modified

## Removed

## Decisions

## Provenance
- impl %s agent=manual fireweed=%s playbook=%x tests=1 failing=0 passing=1
`, module, name, date, assets.Version(), hash))
}

func Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive(t *testing.T) {
	root := regenerateFixture(t, "protected: [packages/alpha/internal/global.go]\n")
	change := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(change, "regenerate: true\n", "regenerate: true\nset:\n  preserve: [internal/keep.go]\n", 1))
	spec := readFile(t, root, "packages/alpha/SPEC.md")
	spec = strings.Replace(spec, "depends_on: []\n", "depends_on: []\ninterface: [api.go, internal/contract.go]\nappend_only: [internal/log.go]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", spec)
	files := map[string]string{
		"packages/alpha/api.go":               "package alpha\n",
		"packages/alpha/evals/alpha_test.go":  "package alpha_test\n// retained eval\n",
		"packages/alpha/internal/contract.go": "package alpha\n// retained interface\n",
		"packages/alpha/internal/log.go":      "retained append-only\n",
		"packages/alpha/internal/keep.go":     "retained preserved\n",
		"packages/alpha/internal/global.go":   "retained global protected\n",
		"packages/alpha/manifest.txt":         "retained module manifest\n",
		"packages/beta/internal/other.go":     "another module\n",
	}
	for name, body := range files {
		write(t, root, name, body)
	}
	write(t, root, "packages/alpha/internal/ready", "old implementation\n")
	write(t, root, "packages/alpha/internal/old.go", "old disposable\n")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid regenerate fixture: %#v", d)
		}
	}
	commit(t, root, "approved regenerate fixture")
	before := head(t, root)
	s := start(t, root)
	if s.Phase != "impl" || s.Run == "" || s.ClaimOID == "" || s.PhaseBase != before || head(t, root) != before {
		t.Errorf("regenerate Start phase/claim/base = %q/%q/%q; HEAD before %s after %s", s.Phase, s.ClaimOID, s.PhaseBase, before, head(t, root))
	}
	for _, name := range []string{"packages/alpha/internal/ready", "packages/alpha/internal/old.go"} {
		if _, err := os.Stat(filepath.Join(root, name)); !os.IsNotExist(err) {
			t.Errorf("disposable %s remained: %v", name, err)
		}
	}
	for name, want := range files {
		if got := readFile(t, root, name); got != want {
			t.Errorf("retained %s changed: %q", name, got)
		}
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" || current.Run != s.Run {
		t.Fatalf("held regenerate state = %#v, %v", current, err)
	}
	regenerate, err := assets.Playbook("regenerate")
	if err != nil {
		t.Fatal(err)
	}
	impl, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	context, err := workflow.Context(root, "manual", false)
	if err != nil || !strings.Contains(context, regenerate) || strings.Contains(context, impl) {
		t.Errorf("regenerate Context did not select playbook: %v, %q", err, context)
	}
	if s.Phase != "impl" {
		return // Archive reconstruction requires the newly specified starting phase.
	}
	write(t, root, "packages/alpha/internal/ready", "reconstructed implementation\n")
	checked, err := verify(root, false)
	if err != nil || len(checked.Reports["alpha"].Cases) != 2 || len(checked.Reports["beta"].Cases) != 1 {
		t.Fatalf("checked reconstruction = %#v, %v", checked.Reports, err)
	}
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" {
		t.Fatalf("regenerate Done = %#v, %v", applied, err)
	}
	if head(t, root) == before {
		t.Error("archive did not commit checked reconstruction")
	}
	archive := readFile(t, root, "changes/archive/0001-manual.md")
	if !strings.Contains(archive, "phase: applied") || !strings.Contains(archive, "- impl ") {
		t.Errorf("regenerate archive lacks implementation provenance: %q", archive)
	}
	for name, want := range files {
		if got := readFile(t, root, name); got != want {
			t.Errorf("archived reconstruction changed retained %s: %q", name, got)
		}
	}
}

func Test_FLOW_082_RegenerateWarningUsesLatestTargetImplProvenanceAndReportDate(t *testing.T) {
	for _, tc := range []struct {
		name, date, result string
		warn               bool
	}{
		{name: "missing", warn: true},
		{name: "older passing", date: "2026-06-14T12:00:00Z", result: "pass", warn: true},
		{name: "newer failed", date: "2026-06-16T12:00:00Z", result: "fail", warn: true},
		{name: "newer passing", date: "2026-06-16T12:00:00Z", result: "pass", warn: false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := regenerateFixture(t, "")
			archivedImpl(t, root, "0000-alpha-earlier", "alpha", "2026-06-01T12:00:00Z")
			archivedImpl(t, root, "0001-alpha-latest", "alpha", "2026-06-15T12:00:00Z")
			archivedImpl(t, root, "0002-beta-later", "beta", "2026-06-20T12:00:00Z")
			if tc.date != "" {
				// An ordinary YAML timestamp is a valid authored report date.
				write(t, root, ".fireweed/drills/alpha.md", fmt.Sprintf("---\ndate: %s\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: %s\n---\n\n## Own evals\n\n## Dependent evals\n\n## Contradicted Decisions\n\n## Gaps\n", tc.date, head(t, root), tc.result))
			}
			for _, d := range project.Check(root) {
				if !d.Warning {
					t.Fatalf("invalid date/report fixture: %#v", d)
				}
			}
			commit(t, root, "approved archives and authored report")
			s := start(t, root)
			if got := containsCode(s.Diagnostics, "P10"); got != tc.warn {
				t.Errorf("P10 warning = %t; want %t; diagnostics = %#v", got, tc.warn, s.Diagnostics)
			}
			if tc.warn {
				for _, d := range s.Diagnostics {
					if d.Code == "P10" && !d.Warning {
						t.Errorf("P10 diagnostic is not a warning: %#v", d)
					}
				}
			}
		})
	}
}

func recordedRegenerateFixture(t *testing.T) string {
	t.Helper()
	root := regenerateFixture(t, "")
	change := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(change, "regenerate: true\n", "regenerate: true\nphase: impl\nrun: fixture-recorded\n", 1))
	write(t, root, "packages/alpha/internal/ready", "recorded phase implementation\n")
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid recorded-phase fixture: %#v", d)
		}
	}
	commit(t, root, "approved recorded regenerate phase")
	return root
}

func Test_FLOW_081_RecordedRegenerateStartRetainsDisposableFilesAndState(t *testing.T) {
	root := recordedRegenerateFixture(t)
	s := start(t, root)
	if s.Phase != "impl" || s.Run != "fixture-recorded" {
		t.Fatalf("recorded regenerate state = %#v", s)
	}
	if got := readFile(t, root, "packages/alpha/internal/ready"); got != "recorded phase implementation\n" {
		t.Errorf("recorded phase repeated deletion: %q", got)
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" || current.Run != "fixture-recorded" {
		t.Errorf("held recorded regenerate state = %#v, %v", current, err)
	}
}

func Test_FLOW_083_RecordedRegenerateImplContextUsesRegeneratePlaybook(t *testing.T) {
	root := recordedRegenerateFixture(t)
	s := start(t, root)
	if s.Phase != "impl" || s.Run != "fixture-recorded" {
		t.Fatalf("recorded regenerate state = %#v", s)
	}
	context, err := workflow.Context(root, "manual", false)
	if err != nil {
		t.Fatal(err)
	}
	playbook, err := assets.Playbook("regenerate")
	if err != nil {
		t.Fatal(err)
	}
	ordinary, err := assets.Playbook("impl")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(context, playbook) || strings.Contains(context, ordinary) {
		t.Errorf("recorded regenerate Context selected wrong playbook: %q", context)
	}
}

 succeeded in 689ms:
func Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports(t *testing.T) {
	provenanceLines := func(root, path string) []string {
		body, ok := contracts.Section(readFile(t, root, path), "Provenance")
		if !ok {
			t.Fatal("Provenance section missing")
		}
		var lines []string
		for _, line := range strings.Split(body, "\n") {
			if strings.TrimSpace(line) != "" {
				lines = append(lines, line)
			}
		}
		return lines
	}
	seedProvenance := func(root string) []string {
		playbook, err := assets.Playbook("audit")
		if err != nil {
			t.Fatal(err)
		}
		hash := sha256.Sum256([]byte(playbook))
		first := fmt.Sprintf("- audit 2026-01-01T00:00:00Z agent=manual fireweed=%s playbook=%x tests=0 failing=0 passing=0", assets.Version(), hash)
		second := fmt.Sprintf("- audit 2026-01-02T00:00:00Z agent=prior-agent  fireweed=%s playbook=%x tests=0 failing=0 passing=0", assets.Version(), hash)
		change := readFile(t, root, pendingCommitChange)
		write(t, root, pendingCommitChange, contracts.ReplaceSection(change, "Provenance", "\n"+first+"\n\n \t\n"+second+"\n\n"))
		return []string{first, second}
	}
	assertAppend := func(root, path string, previous []string) []string {
		lines := provenanceLines(root, path)
		if len(lines) != len(previous)+1 {
			t.Fatalf("Done provenance lines = %q; want unchanged prefix %q and exactly one appended entry", lines, previous)
		}
		if !reflect.DeepEqual(lines[:len(previous)], previous) {
			t.Errorf("Done changed nonblank provenance prefix: got %q, want %q", lines[:len(previous)], previous)
		}
		return lines
	}
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	previous := seedProvenance(root)
	write(t, root, "packages/alpha/internal/context-note.txt", "Full context content marker.\n")
	commit(t, root, "context file")
	state := start(t, root)
	if state.Snapshot.Root != root || len(state.Snapshot.Modules) != 3 {
		t.Errorf("held snapshot absent: %#v", state)
	}
	compact, err := workflow.Context(root, "manual", false)
	if err != nil {
		t.Fatal(err)
	}
	full, err := workflow.Context(root, "manual", true)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Users need a revised alpha value.", "Alpha MUST retain the second behavior.", "# Fixture system", "# Write evals", "package alpha"} {
		if !strings.Contains(compact, want) {
			t.Errorf("compact context misses %q", want)
		}
	}
	if strings.Count(compact, newClause) < 2 {
		t.Errorf("compact context misses revised clause in change and projected spec")
	}
	if !strings.Contains(compact, "packages/alpha/internal/context-note.txt") || !regexp.MustCompile(`(?m)context-note\.txt[^\n]*\b29\b`).MatchString(compact) {
		t.Errorf("compact context omitted file path/size: %q", compact)
	}
	if !strings.Contains(full, "Full context content marker.") {
		t.Error("full context omitted relevant file contents")
	}
	evalsResult, err := done(root)
	if err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	previous = assertAppend(root, pendingCommitChange, previous)
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	implResult, err := done(root)
	if err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	assertAppend(root, archivedCommitChange, previous)
	b, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	provenance, ok := contracts.Section(string(b), "Provenance")
	if !ok || !strings.Contains(provenance, "evals") || !strings.Contains(provenance, "impl") || !strings.Contains(provenance, "manual") || !strings.Contains(provenance, "0.1.0") || !strings.Contains(provenance, "202") {
		t.Errorf("provenance incomplete: %q", provenance)
	}
	for _, phase := range []string{"evals", "impl"} {
		line := ""
		for _, candidate := range strings.Split(provenance, "\n") {
			if strings.HasPrefix(strings.TrimSpace(candidate), "- "+phase+" ") {
				line = strings.TrimSpace(candidate)
				break
			}
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			t.Errorf("%s provenance line misses date: %q", phase, line)
		} else if _, err := time.Parse(time.RFC3339Nano, fields[2]); err != nil {
			t.Errorf("%s provenance line has invalid date %q: %v", phase, fields[2], err)
		}
		if !strings.Contains(line, " agent=manual ") || !strings.Contains(line, " fireweed="+assets.Version()+" ") {
			t.Errorf("%s provenance line misses agent or version: %q", phase, line)
		}
		playbook, err := assets.Playbook(phase)
		if err != nil {
			t.Fatal(err)
		}
		hash := sha256.Sum256([]byte(playbook))
		if line == "" || !strings.Contains(line, "playbook="+hex.EncodeToString(hash[:])) {
			t.Errorf("%s provenance line misses matching playbook hash: %q", phase, line)
		}
		result := evalsResult
		if phase == "impl" {
			result = implResult
		}
		total, passing, failing := 0, 0, 0
		for _, report := range result.Reports {
			for _, c := range report.Cases {
				total++
				switch c.Result {
				case "pass":
					passing++
				case "fail", "error":
					failing++
				}
			}
		}
		if total == 0 || passing == 0 || !strings.Contains(line, fmt.Sprintf("tests=%d failing=%d passing=%d", total, failing, passing)) {
			t.Errorf("%s provenance line has wrong result counts for reports %#v: %q", phase, result.Reports, line)
		}
	}
	runRoot := filepath.Join(root, ".fireweed/runs", state.Run)
	count := 0
	retained := ""
	_ = filepath.Walk(runRoot, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			count++
			b, readErr := os.ReadFile(path)
			if readErr == nil {
				retained += string(b)
			}
		}
		return nil
	})
	if count == 0 || !strings.Contains(retained, "ALP-001") || !strings.Contains(retained, "fixture-hook-output") {
		t.Errorf("reports/output missing under %s (%d files)", runRoot, count)
	}
	configured := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	configuredPrevious := seedProvenance(configured)
	commit(t, configured, "approved prior configured-agent provenance")
	start(t, configured)
	req := request(configured)
	req.Agent = "fixture-agent --model local"
	if _, err := workflow.Done(context.Background(), req); err != nil {
		t.Fatalf("configured-agent evals Done: %v", err)
	}
	assertAppend(configured, pendingCommitChange, configuredPrevious)
	change, err := os.ReadFile(filepath.Join(configured, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(change), "fixture-agent --model local") {
		t.Errorf("configured agent omitted from provenance: %v", err)
	}
}

func Test_FLOW_054_StartOperationalPreconditionsHaveP01(t *testing.T) {
	dirty := fixture(t, "", "", "", "", "")
	write(t, dirty, "stray.txt", "untracked\n")
	if _, err := workflow.Start(dirty, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P01: start requires clean worktree") {
		t.Errorf("dirty Start = %v", err)
	}
	queued := fixture(t, "", "", "", "", "")
	write(t, queued, "changes/0000-first.md", "---\nmodule: alpha\n---\n## Motivation\nFirst request.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	commit(t, queued, "earlier queued change")
	if _, err := workflow.Start(queued, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P01: change is not eligible queue head") {
		t.Errorf("ineligible Start = %v", err)
	}
}

func Test_FLOW_056_StartContentionAndCurrentForeignClaimHaveP02(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	start(t, root)
	if _, err := workflow.Start(other, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02: claim already held:") || !strings.Contains(err.Error(), "update-ref") {
		t.Errorf("contended Start = %v", err)
	}
	if _, err := workflow.Current(other, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02: claim belongs to another worktree") {
		t.Errorf("foreign Current = %v", err)
	}
}

func Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks(t *testing.T) {
	for _, edit := range []string{"author", "runtime", "audit"} {
		t.Run(edit, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			marker := hookMarker(t, root)
			start(t, root)
			p := filepath.Join(root, "changes/0001-manual.md")
			b, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			want := ""
			switch edit {
			case "author":
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "Users need a revised alpha value.", "Unapproved changed purpose.", 1))
				want = "P03: approved author metadata changed"
			case "runtime":
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "phase: evals", "phase: impl", 1))
				want = "P03: approved run or phase changed"
			case "audit":
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "phase: evals", "audit_rounds: 1\nphase: evals", 1))
				want = "P03: approved audit state changed"
			}
			before := head(t, root)
			if _, err := verify(root, true); err == nil || !strings.HasPrefix(err.Error(), want) {
				t.Errorf("%s Verify = %v", edit, err)
			}
			if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), want) || head(t, root) != before {
				t.Errorf("%s Done = %v", edit, err)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Errorf("hook ran before %s rejection: %v", edit, err)
			}
		})
	}
}

func Test_FLOW_059_AuditNoFindingsErrorsHaveP09(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	before := head(t, root)
	audit(t, root, "No findings\n")
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P09: audit requires Checked IDs and No findings") || head(t, root) != before {
		t.Errorf("malformed No findings = %v", err)
	}
	audit(t, root, "Checked IDs: ALP-002\nNo findings\n")
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P09: audit did not check ALP-001") || head(t, root) != before {
		t.Errorf("incomplete No findings = %v", err)
	}
}

func Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- [untestable] ALP-001: \""+newClause+"\" — cannot observe the value\n")
	before := head(t, root)
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P06: change blocked") {
		t.Errorf("blocked Done = %v", err)
	}
	b, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	questions, ok := contracts.Section(string(b), "Questions")
	if head(t, root) == before || !strings.Contains(string(b), "blocked:") || !ok || !strings.Contains(questions, newClause) {
		t.Errorf("blocking state missing from commit: %s", b)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
}

func Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses(t *testing.T) {
	evalRoot := fixture(t, "", "", "", "", "")
	p := filepath.Join(evalRoot, ".fireweed/eval.sh")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	write(t, evalRoot, ".fireweed/eval.sh", strings.ReplaceAll(string(b), `name="ALP-001"`, `name="ALP-002"`))
	commit(t, evalRoot, "hook omits target")
	start(t, evalRoot)
	if _, err := verify(evalRoot, true); err == nil || !strings.HasPrefix(err.Error(), "C10: no coverage for ALP-001") {
		t.Errorf("coverage cause = %v", err)
	}
	before := head(t, evalRoot)
	if _, err := done(evalRoot); err == nil || !strings.HasPrefix(err.Error(), "C10: no coverage for ALP-001") || head(t, evalRoot) != before {
		t.Errorf("Done coverage cause = %v", err)
	}
	missingReport := fixture(t, "", "", "", "", "")
	write(t, missingReport, ".fireweed/eval.sh", "#!/bin/sh\nexit 0\n")
	commit(t, missingReport, "hook omits JUnit report")
	start(t, missingReport)
	if _, err := verify(missingReport, true); err == nil || !strings.HasPrefix(err.Error(), "C11: read JUnit report:") {
		t.Errorf("missing JUnit cause = %v", err)
	}
	failFirst := fixture(t, "", "", "", "", "")
	write(t, failFirst, ".fireweed/unrelated", "yes\n")
	commit(t, failFirst, "unrelated failing case")
	start(t, failFirst)
	write(t, failFirst, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(failFirst, true); err == nil || !strings.HasPrefix(err.Error(), "P2-01: no fail-first case for ALP-001") {
		t.Errorf("fail-first cause = %v", err)
	}
	implFailure := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	start(t, implFailure)
	if _, err := done(implFailure); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	if _, err := verify(implFailure, true); err == nil || !strings.HasPrefix(err.Error(), "E01: failed case") || !strings.Contains(err.Error(), "ALP-001") {
		t.Errorf("impl testcase cause = %v", err)
	}

	zoneRoot := fixture(t, "", "", "", "", "")
	write(t, zoneRoot, "packages/alpha/internal/existing.go", "package internal\n")
	commit(t, zoneRoot, "existing implementation")
	start(t, zoneRoot)
	write(t, zoneRoot, "packages/alpha/internal/existing.go", "package internal // changed in evals\n")
	if _, err := verify(zoneRoot, true); err == nil || !strings.HasPrefix(err.Error(), "Z01:") {
		t.Errorf("zone cause = %v", err)
	}

	boundaryRoot := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	write(t, boundaryRoot, ".fireweed/boundary-fail", "yes\n")
	commit(t, boundaryRoot, "failing boundary fixture")
	start(t, boundaryRoot)
	if _, err := done(boundaryRoot); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, boundaryRoot, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(boundaryRoot, false); err == nil || !strings.HasPrefix(err.Error(), "X01:") || !strings.Contains(err.Error(), "forbidden dependency") {
		t.Errorf("boundary cause = %v", err)
	}
	boundaryHook := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	config, err := os.ReadFile(filepath.Join(boundaryHook, "fireweed.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	write(t, boundaryHook, "fireweed.yaml", strings.Replace(string(config), "boundaries: [sh, .fireweed/boundaries.sh]", "boundaries: [/missing/fireweed-boundary-hook]", 1))
	commit(t, boundaryHook, "unlaunchable boundary hook")
	start(t, boundaryHook)
	if _, err := done(boundaryHook); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, boundaryHook, "packages/alpha/internal/ready", "ready\n")
	snapshot, err := project.Load(boundaryHook)
	if err != nil {
		t.Fatal(err)
	}
	req := request(boundaryHook)
	boundaryDiagnostics, boundaryErr := checks.Boundaries(context.Background(), checks.BoundaryRequest{
		Snapshot: snapshot, Bin: req.Bin, Timeout: req.Timeout,
	})
	if boundaryErr == nil || !strings.Contains(boundaryErr.Error(), "/missing/fireweed-boundary-hook") ||
		len(boundaryDiagnostics) != 1 || boundaryDiagnostics[0].Code != "X02" ||
		!strings.Contains(boundaryDiagnostics[0].Message, "/missing/fireweed-boundary-hook") {
		t.Fatalf("unlaunchable boundary fixture = diagnostics %#v, error %v; want X02 naming missing executable", boundaryDiagnostics, boundaryErr)
	}
	want := boundaryErr.Error()
	if !strings.HasPrefix(want, "X02: ") {
		want = "X02: " + want
	}
	if strings.Count(want, "X02:") != 1 {
		t.Fatalf("boundary cause has repeated X02: %q", want)
	}
	if _, err := verify(boundaryHook, false); err == nil || err.Error() != want {
		t.Errorf("boundary hook cause = %v, want exactly %q", err, want)
	}
}

func Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	config := readFile(t, root, "fireweed.yaml")
	withoutHook := strings.Replace(config, "  boundaries: [sh, .fireweed/boundaries.sh]\n", "", 1)
	if withoutHook == config {
		t.Fatal("fixture did not remove the required boundary hook")
	}
	write(t, root, "fireweed.yaml", withoutHook)
	commit(t, root, "approved fixture without boundary hook")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" {
		t.Fatalf("impl state = %#v, %v", current, err)
	}
	req := request(root)
	boundaryDiagnostics, boundaryErr := checks.Boundaries(context.Background(), checks.BoundaryRequest{
		Snapshot: current.Snapshot, Bin: req.Bin, Timeout: req.Timeout,
	})
	if boundaryErr == nil || strings.TrimSpace(boundaryErr.Error()) == "" {
		t.Fatalf("checks boundary cause = %v; want nonempty cause", boundaryErr)
	}
	wantModules := map[string]bool{"alpha": true, "beta": true, "gamma": true}
	if len(boundaryDiagnostics) != len(wantModules) {
		t.Fatalf("boundary diagnostics = %#v; want one per affected module", boundaryDiagnostics)
	}
	for _, d := range boundaryDiagnostics {
		message := strings.ToLower(d.Message)
		if d.Code != "X02" || d.Warning || !strings.Contains(message, "missing") ||
			!strings.Contains(message, "boundar") || !strings.Contains(message, "hook") {
			t.Fatalf("unexpected boundary diagnostic: %#v", d)
		}
		matched := false
		for module := range wantModules {
			if strings.Contains(d.Message, module) {
				delete(wantModules, module)
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf("boundary diagnostic has no remaining affected module: %#v", d)
		}
	}
	if len(wantModules) != 0 {
		t.Fatalf("boundary diagnostics omitted modules: %v", wantModules)
	}
	result, err := verify(root, false)
	want := boundaryErr.Error()
	if !strings.HasPrefix(want, "X02: ") {
		want = "X02: " + want
	}
	if strings.Count(want, "X02:") != 1 {
		t.Fatalf("checks boundary cause has repeated X02: %q", want)
	}
	if err == nil || err.Error() != want || strings.Count(err.Error(), "X02:") != 1 {
		t.Errorf("full Verify cause = %v; want exactly %q", err, want)
	}
	if !reflect.DeepEqual(result.Diagnostics, boundaryDiagnostics) {
		t.Errorf("full Verify diagnostics = %#v; want %#v", result.Diagnostics, boundaryDiagnostics)
	}
}
package evals_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	base := head(t, root)
	current, err := workflow.Current(root, "manual")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ body, want string }{
		{"No findings\n", "P09: audit requires Checked IDs and No findings"},
		{"Checked IDs: ALP-002\nNo findings\n", "P09: audit did not check ALP-001"},
	} {
		audit(t, root, tc.body)
		before := readFile(t, root, pendingCommitChange)
		for _, quick := range []bool{true, false} {
			result, err := verify(root, quick)
			if err == nil || !strings.HasPrefix(err.Error(), tc.want) {
				t.Errorf("audit Verify(quick=%t, %q) = phase %q, error %v; want %s", quick, tc.body, result.State.Phase, err, tc.want)
			}
		}
		if got := readFile(t, root, pendingCommitChange); got != before {
			t.Errorf("rejected Verify changed audit submission: %q", got)
		}
		if head(t, root) != base {
			t.Error("rejected Verify committed")
		}
		if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
			t.Errorf("rejected Verify changed claim or phase: %#v, %v", now, err)
		}
	}

	// Findings that Done can process must also pass Verify without being filtered or written.
	audit(t, root, "- ALP-001: \""+newClause+"\" — missing assertion\n- ALP-999: \"Unknown MUST work.\" — invalid finding\n")
	beforeFinding := readFile(t, root, pendingCommitChange)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("finding audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != beforeFinding || head(t, root) != base {
		t.Error("finding audit Verify normalized submission or committed")
	}

	// A complete submission is accepted by Verify without Done's normalization.
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	before := readFile(t, root, pendingCommitChange)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("complete audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != before || head(t, root) != base {
		t.Error("successful audit Verify wrote the submission or committed")
	}
	if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
		t.Errorf("successful Verify changed claim or phase: %#v, %v", now, err)
	}
	if result, err := done(root); err != nil || result.State.Phase != "impl" {
		t.Errorf("complete audit Done = phase %q, error %v", result.State.Phase, err)
	}
}

func Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	marker := hookMarker(t, root)
	auditState := toAudit(t, root)
	if auditState.Phase != "audit" {
		t.Fatalf("held phase = %q, want audit", auditState.Phase)
	}
	if err := os.Remove(marker); err != nil {
		t.Fatalf("clear evals Done hook marker: %v", err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	submission := readFile(t, root, pendingCommitChange)
	forbidden := "packages/alpha/evals/audit_zone_test.go"
	write(t, root, forbidden, "package alpha_test\n")
	base := head(t, root)

	for _, operation := range []struct {
		name string
		call func() (workflow.Result, error)
	}{
		{"quick Verify", func() (workflow.Result, error) { return verify(root, true) }},
		{"full Verify", func() (workflow.Result, error) { return verify(root, false) }},
		{"Done", func() (workflow.Result, error) { return done(root) }},
	} {
		result, err := operation.call()
		if err == nil || !strings.HasPrefix(err.Error(), "Z01:") || !strings.Contains(err.Error(), forbidden) {
			t.Errorf("%s = phase %q, error %v; want Z01 naming %s", operation.name, result.State.Phase, err, forbidden)
		}
		if got := head(t, root); got != base {
			t.Errorf("%s advanced HEAD: %s -> %s", operation.name, base, got)
		}
		if got := readFile(t, root, pendingCommitChange); got != submission {
			t.Errorf("%s changed the valid audit submission", operation.name)
		}
		current, currentErr := workflow.Current(root, "manual")
		if currentErr != nil || current.Phase != "audit" || current.PhaseBase != base || current.ClaimOID != auditState.ClaimOID {
			t.Errorf("%s changed held audit phase or claim: %#v, %v", operation.name, current, currentErr)
		}
		if _, markerErr := os.Stat(marker); !os.IsNotExist(markerErr) {
			t.Errorf("%s ran the eval hook before zone rejection: %v", operation.name, markerErr)
		}
	}
}

func Test_FLOW_061_UnknownLifecycleSelectorsNameValidSlugs(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	request := workflow.Request{Root: root, Change: "ghost"}
	operations := []struct {
		name string
		call func(string) error
	}{
		{"Start", func(s string) error { _, err := workflow.Start(root, s); return err }},
		{"StartFromAttempt", func(s string) error { _, err := workflow.StartFromAttempt(root, s, 1); return err }},
		{"Current", func(s string) error { _, err := workflow.Current(root, s); return err }},
		{"Context", func(s string) error { _, err := workflow.Context(root, s, false); return err }},
		{"Verify", func(s string) error {
			request.Change = s
			_, err := workflow.Verify(context.Background(), request)
			return err
		}},
		{"Done", func(s string) error {
			request.Change = s
			_, err := workflow.Done(context.Background(), request)
			return err
		}},
		{"Release", func(s string) error { return workflow.Release(root, s) }},
		{"Commits", func(s string) error { _, err := workflow.Commits(root, s); return err }},
		{"Abandon", func(s string) error { _, err := workflow.Abandon(root, s); return err }},
	}
	before := head(t, root)
	for _, selector := range []string{"ghost", "0009-ghost.md"} {
		for _, op := range operations {
			err := op.call(selector)
			if err == nil || !strings.HasPrefix(err.Error(), "U02:") || !strings.Contains(err.Error(), selector) || !strings.Contains(err.Error(), "manual") {
				t.Errorf("%s(%q) = %v; want U02 with rejected selector and valid slug manual", op.name, selector, err)
			}
		}
	}
	if head(t, root) != before || git(t, root, "status", "--porcelain=v1", "--untracked-files=all") != "" {
		t.Error("unknown selector changed HEAD or working tree")
	}

	// Commits and Abandon also look up archived documents.
	archive := filepath.Join(root, archivedCommitChange)
	if err := os.MkdirAll(filepath.Dir(archive), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", pendingCommitChange, archivedCommitChange)
	commit(t, root, "archive fixture change")
	for _, op := range operations[len(operations)-2:] {
		err := op.call("ghost")
		if err == nil || !strings.HasPrefix(err.Error(), "U02:") || !strings.Contains(err.Error(), "ghost") || !strings.Contains(err.Error(), "manual") {
			t.Errorf("archived %s(ghost) = %v; want U02 with archive slug manual", op.name, err)
		}
	}
}

exec
/bin/zsh -lc "sed -n '525,829p' packages/workflow/evals/lifecycle_test.go; sed -n '829,1135p' packages/workflow/evals/lifecycle_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 481ms:
		t.Run(scenario, func(t *testing.T) {
			front := ""
			if scenario == "already-satisfied" {
				front = "already_satisfied: [ALP-001]\n"
			}
			if scenario == "fixes" {
				front = "fixes: [ALP-002]\n"
			}
			root := fixture(t, "", front, "", "", "")
			if scenario == "missing-coverage" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				write(t, root, ".fireweed/eval.sh", strings.ReplaceAll(string(b), `name="ALP-001"`, `name="ALP-002"`))
				commit(t, root, "hook omits target")
			}
			if scenario == "already-satisfied" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				script := strings.Replace(string(b), `  exit "$exitcode"`, `  if [ -f packages/alpha/evals/no-coverage ]; then sed 's/name="ALP-002"/name="ALP-001"/g' "$report" > "$report.tmp"; mv "$report.tmp" "$report"; fi
  exit "$exitcode"`, 1)
				if script == string(b) {
					t.Fatal("fixture coverage switch was not installed")
				}
				write(t, root, ".fireweed/eval.sh", script)
				commit(t, root, "hook can omit unrelated current spec")
			}
			if scenario == "passing-target-unrelated-failure" || scenario == "already-satisfied" {
				write(t, root, ".fireweed/unrelated", "yes\n")
				commit(t, root, "known unrelated failing case")
			}
			start(t, root)
			if scenario == "passing-target-unrelated-failure" || scenario == "already-satisfied" {
				write(t, root, "packages/alpha/internal/ready", "ready\n")
			}
			before := head(t, root)
			r, err := verify(root, true)
			wantPass := scenario == "already-satisfied"
			if (err == nil) != wantPass {
				t.Errorf("Verify = %#v, %v; want pass %v", r, err, wantPass)
			}
			if head(t, root) != before {
				t.Error("Verify committed")
			}
			if wantPass && (len(r.Reports["alpha"].Cases) == 0 || r.State.Snapshot.Root != root) {
				t.Errorf("reports or held snapshot absent: %#v", r)
			}
			if scenario == "already-satisfied" {
				write(t, root, "packages/alpha/evals/no-coverage", "omit ALP-002\n")
				missing, err := verify(root, true)
				if len(missing.Reports["alpha"].Cases) == 0 {
					t.Errorf("target module evals did not run: %#v, %v", missing, err)
				}
				if err == nil {
					t.Errorf("evals accepted report without unrelated current spec ALP-002: %#v", missing)
				}
			}
		})
	}
}

func Test_FLOW_023_AddedIDRequiresMappedFailingCase(t *testing.T) {
	root := fixture(t, "", "set:\n  interface: [api.go, extra.go]\n", "### new:third\nThird MUST return a value.\n", "", "")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, _ := os.ReadFile(p)
	script := strings.Replace(string(b), `</testsuite>' "$result" "$other"`, `<testcase name="ALP-003" classname="fixture">%s</testcase></testsuite>' "$result" "$other" "$result"`, 1)
	if script == string(b) {
		t.Fatal("fixture report extension was not installed")
	}
	write(t, root, ".fireweed/eval.sh", script)
	commit(t, root, "report projected added ID")
	s := start(t, root)
	if len(s.Change.Added) != 1 || s.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("assigned Added ID = %#v", s.Change.Added)
	}
	write(t, root, "packages/alpha/extra.go", "package alpha\n")
	r, err := verify(root, true)
	if err != nil {
		t.Fatalf("mapped failing Added case rejected: %#v, %v", r, err)
	}
	if len(r.Reports["alpha"].Cases) != 3 {
		t.Errorf("added ID was not reported: %#v", r.Reports["alpha"])
	}
}

func Test_FLOW_023_ReviewTaggedAdditionNeedsNoFailFirstCase(t *testing.T) {
	root := fixture(t, "", "", "### new:review [review]\nReview judgment MUST be recorded.\n", "", "")
	s := start(t, root)
	if s.Phase != "evals" || len(s.Change.Added) != 1 || s.Change.Added[0].ID != "ALP-003" || s.Change.Added[0].Tag != "review" {
		t.Fatalf("fixture lacks an assigned review-tagged change in evals: %#v", s)
	}
	before := head(t, root)
	r, err := verify(root, true)
	if err != nil {
		t.Fatalf("evals Verify rejected review item without a failing case: %#v, %v", r.Reports["alpha"], err)
	}
	if head(t, root) != before {
		t.Error("Verify committed the review fixture")
	}
	cases := r.Reports["alpha"].Cases
	if len(cases) != 2 || cases[0].Result != "fail" || !reflect.DeepEqual(cases[0].IDs, []string{"ALP-001"}) || cases[1].Result != "pass" || !reflect.DeepEqual(cases[1].IDs, []string{"ALP-002"}) {
		t.Errorf("fixture needs genuine fail-first for enforced ALP-001 and no fabricated review case: %#v", cases)
	}
}

func Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries(t *testing.T) {
	for _, scenario := range []string{"frozen-eval", "direct", "all", "none", "quick-boundary"} {
		t.Run(scenario, func(t *testing.T) {
			config := "audit:\n  enabled: false\n"
			if scenario == "all" {
				config += "verify_dependents: all\n"
			}
			if scenario == "none" {
				config += "verify_dependents: none\n"
			}
			root := fixture(t, config, "", "", "", "")
			if scenario == "frozen-eval" {
				p := filepath.Join(root, ".fireweed/eval.sh")
				b, _ := os.ReadFile(p)
				script := strings.Replace(string(b), `  exit "$exitcode"`, `  if [ -f packages/alpha/internal/no-coverage ]; then sed 's/name="ALP-001"/name="ALP-002"/g' "$report" > "$report.tmp"; mv "$report.tmp" "$report"; fi
  exit "$exitcode"`, 1)
				write(t, root, ".fireweed/eval.sh", script)
				commit(t, root, "coverage switch fixture")
			}
			switch scenario {
			case "direct":
				write(t, root, ".fireweed/beta-fail", "yes\n")
			case "none":
				write(t, root, ".fireweed/beta-fail", "yes\n")
			case "all":
				write(t, root, ".fireweed/gamma-fail", "yes\n")
			case "quick-boundary":
				write(t, root, ".fireweed/boundary-fail", "yes\n")
			}
			if scenario != "frozen-eval" {
				commit(t, root, "configured failing hook")
			}
			start(t, root)
			if _, err := done(root); err != nil {
				t.Fatalf("evals Done: %v", err)
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.Phase != "impl" {
				t.Fatalf("impl state = %#v, %v", current, err)
			}
			write(t, root, "packages/alpha/internal/ready", "ready\n")
			if scenario == "frozen-eval" {
				write(t, root, "packages/alpha/evals/new_test.go", "package evals\n")
			}
			before := head(t, root)
			quick, quickErr := verify(root, true)
			full, fullErr := verify(root, false)
			if head(t, root) != before {
				t.Error("Verify committed")
			}
			if scenario == "frozen-eval" {
				if quickErr == nil || fullErr == nil {
					t.Errorf("frozen eval edit accepted: %#v %v; %#v %v", quick, quickErr, full, fullErr)
				}
				if err := os.Remove(filepath.Join(root, "packages/alpha/evals/new_test.go")); err != nil {
					t.Fatal(err)
				}
				write(t, root, "packages/alpha/internal/no-coverage", "yes\n")
				if r, err := verify(root, true); err == nil {
					t.Errorf("impl accepted missing current coverage: %#v", r)
				}
			} else {
				if quickErr != nil {
					t.Errorf("quick rejected own passing evals: %v", quickErr)
				}
				wantFullPass := scenario == "none"
				if (fullErr == nil) != wantFullPass {
					t.Errorf("full Verify = %#v, %v; want pass %v", full, fullErr, wantFullPass)
				}
				if scenario == "direct" && len(full.Reports["beta"].Cases) == 0 {
					t.Error("direct dependent report absent")
				}
				if scenario == "direct" && len(full.Reports["gamma"].Cases) != 0 {
					t.Error("default direct verification ran a transitive dependent")
				}
				if scenario == "all" && len(full.Reports["gamma"].Cases) == 0 {
					t.Error("transitive dependent report absent")
				}
				if scenario == "none" && (len(full.Reports["beta"].Cases) != 0 || len(full.Reports["gamma"].Cases) != 0) {
					t.Error("dependent verification ran with verify_dependents: none")
				}
				if scenario == "direct" && len(full.Reports["beta"].Cases) > 0 && full.Reports["beta"].Cases[0].Result != "fail" {
					t.Errorf("selected direct dependent did not report its failing testcase: %#v", full.Reports["beta"])
				}
				if scenario == "all" && len(full.Reports["gamma"].Cases) > 0 && full.Reports["gamma"].Cases[0].Result != "fail" {
					t.Errorf("selected transitive dependent did not report its failing testcase: %#v", full.Reports["gamma"])
				}
				if scenario == "direct" || scenario == "all" || scenario == "quick-boundary" {
					r, err := done(root)
					if err == nil || head(t, root) != before {
						t.Errorf("Done accepted failed full verification or committed: %#v, %v", r, err)
					}
					current, currentErr := workflow.Current(root, "manual")
					if currentErr != nil || current.Phase != "impl" || current.PhaseBase != before {
						t.Errorf("failed full verification advanced phase: %#v, %v", current, currentErr)
					}
				}
			}
		})
	}
}

func Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures(t *testing.T) {
	for _, tc := range []struct {
		name, mode string
		ownPass    bool
		selected   []string
	}{
		{"own fail/direct", "direct", false, []string{"alpha", "beta"}},
		{"own fail/all", "all", false, []string{"alpha", "beta", "gamma"}},
		{"own fail/none", "none", false, []string{"alpha"}},
		{"own pass/all", "all", true, []string{"alpha", "beta", "gamma"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, "audit:\n  enabled: false\nverify_dependents: "+tc.mode+"\n", "", "", "", "")
			invocations := filepath.Join(t.TempDir(), "invocations")
			write(t, root, ".fireweed/eval.sh", fmt.Sprintf("#!/bin/sh\nprintf '%%s\\n' \"$1\" >> %q\n", invocations)+`case "$1" in
  alpha)
    if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
    else failure='<failure message="alpha-failure-sentinel">alpha-failure-sentinel</failure>'; code=1; fi
    printf '<testsuite><testcase name="ALP-001 own" classname="fixture">%s</testcase><testcase name="ALP-002 retained" classname="fixture"/></testsuite>' "$failure" > "$2"
    exit "$code"
    ;;
  beta) id=BET-001; sentinel=beta-failure-sentinel ;;
  gamma) id=GAM-001; sentinel=gamma-failure-sentinel ;;
  *) exit 2 ;;
esac
printf '<testsuite><testcase name="%s dependent" classname="fixture"><failure message="%s">%s</failure></testcase></testsuite>' "$id" "$sentinel" "$sentinel" > "$2"
exit 1
`)
			commit(t, root, "approved mapped-failure reports and invocation observer")
			start(t, root)
			if _, err := done(root); err != nil {
				t.Fatalf("fixture evals Done: %v", err)
			}
			held, err := workflow.Current(root, "manual")
			if err != nil || held.Phase != "impl" {
				t.Fatalf("fixture held impl state = %#v, %v", held, err)
			}
			if tc.ownPass {
				write(t, root, "packages/alpha/internal/ready", "ready\n")
			}
			if err := os.WriteFile(invocations, nil, 0644); err != nil {
				t.Fatal(err)
			}
			before := head(t, root)
			result, err := verify(root, false)
			if err == nil {
				t.Error("full impl Verify accepted mapped testcase failures")
			}
			if len(result.Reports) != len(tc.selected) {
				t.Errorf("returned reports = %#v; want exactly %v", result.Reports, tc.selected)
			}
			wantInvocations := map[string]int{}
			for _, name := range tc.selected {
				wantInvocations[name] = 1
				report, ok := result.Reports[name]
				count, id, caseName := 1, "BET-001", "BET-001 dependent"
				if name == "alpha" {
					count, id, caseName = 2, "ALP-001", "ALP-001 own"
				} else if name == "gamma" {
					id, caseName = "GAM-001", "GAM-001 dependent"
				}
				if !ok || len(report.Cases) != count {
					t.Errorf("%s report = %#v, present %t; want %d mapped cases", name, report, ok, count)
					continue
				}
				c := report.Cases[0]
				wantResult, wantExit := "fail", 1
				if name == "alpha" && tc.ownPass {
					wantResult, wantExit = "pass", 0
				}
				if c.Name != caseName || c.Classname != "fixture" || !reflect.DeepEqual(c.IDs, []string{id}) || c.Result != wantResult || report.HookExit != wantExit {
					t.Errorf("%s mapped case/ordinary hook exit = %#v / %d; want %s %s / %d", name, c, report.HookExit, id, wantResult, wantExit)
				}
				if wantResult == "fail" && !strings.Contains(c.Failure, name+"-failure-sentinel") {
					t.Errorf("%s report lost failure text: %#v", name, c)
				}
				if name == "alpha" {
					retained := report.Cases[1]
					if retained.Name != "ALP-002 retained" || !reflect.DeepEqual(retained.IDs, []string{"ALP-002"}) || retained.Result != "pass" {
						t.Errorf("alpha retained mapped passing case = %#v", retained)
					}
				}
			}
			gotInvocations := map[string]int{}
			for _, name := range strings.Fields(readFile(t, "", invocations)) {
				gotInvocations[name]++
			}
			if !reflect.DeepEqual(gotInvocations, wantInvocations) {
				t.Errorf("hook invocations = %v; want only selected modules %v", gotInvocations, wantInvocations)
			}
			current, currentErr := workflow.Current(root, "manual")
			if head(t, root) != before || currentErr != nil || current.Phase != "impl" || current.PhaseBase != held.PhaseBase || current.ClaimOID != held.ClaimOID || current.Run != held.Run {
				t.Errorf("failed Verify changed HEAD or held impl state: before %#v, after %#v, error %v", held, current, currentErr)
			}
		})
	}
}

func audit(t *testing.T, root, body string) {
func audit(t *testing.T, root, body string) {
	t.Helper()
	p := filepath.Join(root, "changes/0001-manual.md")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	text := contracts.ReplaceSection(string(b), "Audit", body)
	write(t, root, "changes/0001-manual.md", text)
}

func toAudit(t *testing.T, root string) workflow.State {
	t.Helper()
	start(t, root)
	r, err := done(root)
	if err != nil || r.State.Phase != "audit" {
		t.Fatalf("evals Done = %#v, %v", r, err)
	}
	return r.State
}

func Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl(t *testing.T) {
	for _, disabled := range []bool{false, true} {
		t.Run(fmt.Sprint(disabled), func(t *testing.T) {
			config := ""
			want := "audit"
			if disabled {
				config = "audit:\n  enabled: false\n"
				want = "impl"
			}
			root := fixture(t, config, "", "", "", "")
			start(t, root)
			before := head(t, root)
			r, err := done(root)
			if err != nil || r.State.Phase != want || head(t, root) == before || r.State.PhaseBase != head(t, root) {
				t.Errorf("Done = %#v, %v; want %s with commit", r, err, want)
			}
			current, err := workflow.Current(root, "manual")
			if err != nil || current.Phase != want || current.PhaseBase != head(t, root) {
				t.Errorf("phase base not advanced: %#v, %v", current, err)
			}
		})
	}
}

func Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- ALP-999: \"Missing MUST clause.\" — invalid ID\n- ALP-001: \"Alpha MUST do something else.\" — invalid quote\nChecked IDs: ALP-001\nNo findings\n")
	r, err := done(root)
	if err != nil || r.State.Phase != "impl" || r.State.PhaseBase != head(t, root) {
		t.Errorf("invalid findings were retained: %#v, %v", r, err)
	}
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if strings.Contains(string(b), "ALP-999") || strings.Contains(string(b), "something else") || !strings.Contains(string(b), "No findings") {
		t.Errorf("audit normalization = %s", b)
	}
	missing := fixture(t, "", "", "", "", "")
	toAudit(t, missing)
	audit(t, missing, "No findings\n")
	before := head(t, missing)
	if _, err := done(missing); err == nil || head(t, missing) != before {
		t.Error("No findings without checked IDs accepted or committed")
	}
}

func Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound(t *testing.T) {
	root := fixture(t, "", "", "### new:extra [review]\nExtra MUST be visible.\n", "", "")
	toAudit(t, root)
	finding := "- ALP-001: \"" + newClause + "\" — missing assertion\n"
	for round := 1; round <= 2; round++ {
		body := finding
		if round == 2 {
			body += "- ALP-003: \"Extra MUST be visible.\" — new finding on re-audit\n"
		}
		audit(t, root, body)
		r, err := done(root)
		if err != nil || r.State.Phase != "evals" || r.State.PhaseBase != head(t, root) {
			t.Fatalf("finding round %d = %#v, %v", round, r, err)
		}
		current, err := workflow.Current(root, "manual")
		if err != nil || current.Change.Document.Frontmatter["audit_rounds"] != round {
			t.Errorf("persistent rounds %d = %#v, %v", round, current, err)
		}
		if round == 2 {
			b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
			if strings.Contains(string(b), "new finding on re-audit") {
				t.Error("new re-audit ID retained")
			}
		}
		if _, err := done(root); err != nil {
			t.Fatalf("revised evals Done round %d: %v", round, err)
		}
	}
	audit(t, root, finding)
	beforeBlock := head(t, root)
	_, _ = done(root)
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	questions, hasQuestions := contracts.Section(string(b), "Questions")
	if !strings.Contains(string(b), "blocked:") || !hasQuestions || !strings.Contains(questions, newClause) {
		t.Errorf("third finding did not block with question: %s", b)
	}
	if head(t, root) == beforeBlock || !strings.Contains(string(b), "audit_rounds: 2") {
		t.Error("two-round bound was not committed with persistent count")
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
	configured := fixture(t, "audit:\n  max_rounds: 1\n", "", "", "", "")
	toAudit(t, configured)
	audit(t, configured, finding)
	first, err := done(configured)
	if err != nil || first.State.Phase != "evals" || first.State.PhaseBase != head(t, configured) {
		t.Fatalf("configured first finding = %#v, %v", first, err)
	}
	firstCurrent, err := workflow.Current(configured, "manual")
	if err != nil || firstCurrent.Change.Document.Frontmatter["audit_rounds"] != 1 {
		t.Errorf("configured round count = %#v, %v", firstCurrent, err)
	}
	if _, err := done(configured); err != nil {
		t.Fatalf("configured revised evals Done: %v", err)
	}
	audit(t, configured, finding)
	_, _ = done(configured)
	change, err := os.ReadFile(filepath.Join(configured, "changes/0001-manual.md"))
	configuredQuestions, configuredHasQuestions := contracts.Section(string(change), "Questions")
	if err != nil || !strings.Contains(string(change), "blocked:") || !strings.Contains(string(change), "audit_rounds: 1") || !configuredHasQuestions || !strings.Contains(configuredQuestions, newClause) {
		t.Errorf("configured max_rounds did not block after one return: %s, %v", change, err)
	}
	if _, err := workflow.Current(configured, "manual"); err == nil {
		t.Error("configured bound retained claim")
	}
}

func Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- [untestable] ALP-001: \""+newClause+"\" — cannot observe the value\n")
	_, _ = done(root)
	b, _ := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	questions, hasQuestions := contracts.Section(string(b), "Questions")
	if !strings.Contains(string(b), "blocked:") || !hasQuestions || !strings.Contains(questions, newClause) {
		t.Errorf("untestable finding did not block: %s", b)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
}

func Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nverify_dependents: none\n", "set:\n  interface: [api.go, extra.go]\n  hooks:\n    eval: [sh, .fireweed/eval.sh, \"{module}\", \"{report}\"]\n", "### new:extra [review]\nExtra MUST be visible.\n", "", "### ALP-002\n")
	p := filepath.Join(root, ".fireweed/eval.sh")
	b, _ := os.ReadFile(p)
	script := strings.Replace(string(b), `<testcase name="ALP-002" classname="fixture">%s</testcase>`, "", 1)
	script = strings.Replace(script, `"$result" "$other" > "$report"`, `"$result" > "$report"`, 1)
	write(t, root, ".fireweed/eval.sh", script)
	write(t, root, ".fireweed/global-eval.sh", "#!/bin/sh\nexit 1\n")
	configPath := filepath.Join(root, "fireweed.yaml")
	config, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "fireweed.yaml", strings.Replace(string(config), "eval: [sh, .fireweed/eval.sh,", "eval: [sh, .fireweed/global-eval.sh,", 1))
	commit(t, root, "report only projected spec IDs")
	start(t, root)
	evalsHead := head(t, root)
	r, err := done(root)
	if err != nil || r.State.Phase != "impl" || head(t, root) == evalsHead || r.State.PhaseBase != head(t, root) {
		t.Fatalf("evals transition = %#v, %v", r, err)
	}
	implBase := head(t, root)
	otherClaim := []byte(`{"worktree":"independent","change":"other"}`)
	otherOID, err := gitstore.WriteBlob(root, otherClaim)
	if err != nil {
		t.Fatal(err)
	}
	otherRef := "refs/fireweed/claims/other"
	if err := gitstore.CASRef(root, otherRef, "", otherOID); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	r, err = done(root)
	if err != nil || r.State.Phase != "applied" || head(t, root) == implBase || r.State.PhaseBase != head(t, root) {
		t.Fatalf("impl transition = %#v, %v", r, err)
	}
	if _, err := os.Stat(filepath.Join(root, "changes/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("pending change remains: %v", err)
	}
	archive, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	spec, err := os.ReadFile(filepath.Join(root, "packages/alpha/SPEC.md"))
	if err != nil {
		t.Fatal(err)
	}
	appliedHead := head(t, root)
	archiveAtHead, err := gitstore.ReadAt(root, appliedHead, "changes/archive/0001-manual.md")
	if err != nil || string(archiveAtHead) != string(archive) {
		t.Errorf("applied HEAD lacks archived change: %q, %v", archiveAtHead, err)
	}
	specAtHead, err := gitstore.ReadAt(root, appliedHead, "packages/alpha/SPEC.md")
	if err != nil || string(specAtHead) != string(spec) {
		t.Errorf("applied HEAD lacks applied spec: %q, %v", specAtHead, err)
	}
	moduleDoc, err := contracts.ParseDocument(string(spec))
	if err != nil {
		t.Fatal(err)
	}
	wantInterface := []any{"api.go", "extra.go"}
	wantHooks := map[string]any{"eval": []any{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}
	if !reflect.DeepEqual(moduleDoc.Frontmatter["interface"], wantInterface) || !reflect.DeepEqual(moduleDoc.Frontmatter["hooks"], wantHooks) {
		t.Errorf("set values missing from module metadata: %#v", moduleDoc.Frontmatter)
	}
	for _, want := range []string{"phase: applied", "## Provenance"} {
		if !strings.Contains(string(archive), want) {
			t.Errorf("archive misses %q", want)
		}
	}
	for _, want := range []string{newClause, "### ALP-003", "extra.go", "Use the revised policy."} {
		if !strings.Contains(string(spec), want) {
			t.Errorf("applied spec misses %q: %s", want, spec)
		}
	}
	if strings.Contains(string(spec), "### ALP-002") || strings.Contains(string(spec), "Retain old policy.") {
		t.Errorf("removed spec/decision remains: %s", spec)
	}
	if !strings.Contains(string(spec), "- manual: Use the revised policy.") {
		t.Errorf("decision attribution missing: %s", spec)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("applied claim remained held")
	}
	if got, exists, err := gitstore.Ref(root, otherRef); err != nil || !exists || got != otherOID {
		t.Errorf("application changed independent claim ref: %q, %v, %v", got, exists, err)
	}
	if got, err := gitstore.ReadBlob(root, otherOID); err != nil || string(got) != string(otherClaim) {
		t.Errorf("application changed independent claim blob: %q, %v", got, err)
	}
}

func Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	hookPath := filepath.Join(root, ".fireweed/eval.sh")
	hook, err := os.ReadFile(hookPath)
	if err != nil {
		t.Fatal(err)
	}
	inject := `  if [ -f packages/alpha/internal/ready ]; then printf 'invalid spec fixture\n' > packages/alpha/SPEC.md; fi
  exit "$exitcode"`
	patched := strings.Replace(string(hook), `  exit "$exitcode"`, inject, 1)
	if patched == string(hook) {
		t.Fatal("fixture hook injection was not installed")
	}
	write(t, root, ".fireweed/eval.sh", patched)
	commit(t, root, "tracked invalid-spec hook fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	base := head(t, root)
	own, ownErr := verify(root, false)
	if ownErr == nil {
		t.Errorf("impl Verify accepted failing current-module testcase: %#v", own)
	}
	if cases := own.Reports["alpha"].Cases; len(cases) == 0 || cases[0].Result != "fail" || len(cases[0].IDs) == 0 || cases[0].IDs[0] != "ALP-001" {
		t.Errorf("failing current-module testcase absent from Verify report: %#v", own.Reports["alpha"])
	}
	r, err := done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("passing current evals required: %#v, %v", r, err)
	}
	current, currentErr := workflow.Current(root, "manual")
	if currentErr != nil || current.Phase != "impl" || current.PhaseBase != base {
		t.Errorf("failed Done advanced phase: %#v, %v", current, currentErr)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	write(t, root, "packages/alpha/api.go", "package alpha // edited during impl\n")
	r, err = done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("frozen interface edit accepted or committed: %#v, %v", r, err)
	}
	write(t, root, "packages/alpha/api.go", "package alpha\n")
	r, err = done(root)
	if err == nil || head(t, root) != base {
		t.Errorf("failed static validation applied or committed: %#v, %v", r, err)
	}
	invalidSpec, readErr := os.ReadFile(filepath.Join(root, "packages/alpha/SPEC.md"))
	if readErr != nil || string(invalidSpec) != "invalid spec fixture\n" {
		t.Errorf("successful eval hook did not inject invalid spec: %q, %v", invalidSpec, readErr)
	}
	invalid := false
	for _, d := range project.Check(root) {
		invalid = invalid || !d.Warning
	}
	if !invalid {
		t.Error("injected spec passed static validation")
	}
	if _, err := os.Stat(filepath.Join(root, "changes/archive/0001-manual.md")); !os.IsNotExist(err) {
		t.Errorf("failed static validation created archive: %v", err)
	}
	pending, readErr := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if readErr != nil || !strings.Contains(string(pending), "phase: impl") || strings.Contains(string(pending), "phase: applied") {
		t.Errorf("failed static validation advanced pending phase: %q, %v", pending, readErr)
	}
}


exec
/bin/zsh -lc "sed -n '1136,1580p' packages/workflow/evals/lifecycle_test.go; cat packages/workflow/evals/selector_audit_validation_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 459ms:
func Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports(t *testing.T) {
	provenanceLines := func(root, path string) []string {
		body, ok := contracts.Section(readFile(t, root, path), "Provenance")
		if !ok {
			t.Fatal("Provenance section missing")
		}
		var lines []string
		for _, line := range strings.Split(body, "\n") {
			if strings.TrimSpace(line) != "" {
				lines = append(lines, line)
			}
		}
		return lines
	}
	seedProvenance := func(root string) []string {
		playbook, err := assets.Playbook("audit")
		if err != nil {
			t.Fatal(err)
		}
		hash := sha256.Sum256([]byte(playbook))
		first := fmt.Sprintf("- audit 2026-01-01T00:00:00Z agent=manual fireweed=%s playbook=%x tests=0 failing=0 passing=0", assets.Version(), hash)
		second := fmt.Sprintf("- audit 2026-01-02T00:00:00Z agent=prior-agent  fireweed=%s playbook=%x tests=0 failing=0 passing=0", assets.Version(), hash)
		change := readFile(t, root, pendingCommitChange)
		write(t, root, pendingCommitChange, contracts.ReplaceSection(change, "Provenance", "\n"+first+"\n\n \t\n"+second+"\n\n"))
		return []string{first, second}
	}
	assertAppend := func(root, path string, previous []string) []string {
		lines := provenanceLines(root, path)
		if len(lines) != len(previous)+1 {
			t.Fatalf("Done provenance lines = %q; want unchanged prefix %q and exactly one appended entry", lines, previous)
		}
		if !reflect.DeepEqual(lines[:len(previous)], previous) {
			t.Errorf("Done changed nonblank provenance prefix: got %q, want %q", lines[:len(previous)], previous)
		}
		return lines
	}
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	previous := seedProvenance(root)
	write(t, root, "packages/alpha/internal/context-note.txt", "Full context content marker.\n")
	commit(t, root, "context file")
	state := start(t, root)
	if state.Snapshot.Root != root || len(state.Snapshot.Modules) != 3 {
		t.Errorf("held snapshot absent: %#v", state)
	}
	compact, err := workflow.Context(root, "manual", false)
	if err != nil {
		t.Fatal(err)
	}
	full, err := workflow.Context(root, "manual", true)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"Users need a revised alpha value.", "Alpha MUST retain the second behavior.", "# Fixture system", "# Write evals", "package alpha"} {
		if !strings.Contains(compact, want) {
			t.Errorf("compact context misses %q", want)
		}
	}
	if strings.Count(compact, newClause) < 2 {
		t.Errorf("compact context misses revised clause in change and projected spec")
	}
	if !strings.Contains(compact, "packages/alpha/internal/context-note.txt") || !regexp.MustCompile(`(?m)context-note\.txt[^\n]*\b29\b`).MatchString(compact) {
		t.Errorf("compact context omitted file path/size: %q", compact)
	}
	if !strings.Contains(full, "Full context content marker.") {
		t.Error("full context omitted relevant file contents")
	}
	evalsResult, err := done(root)
	if err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	previous = assertAppend(root, pendingCommitChange, previous)
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	implResult, err := done(root)
	if err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	assertAppend(root, archivedCommitChange, previous)
	b, err := os.ReadFile(filepath.Join(root, "changes/archive/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	provenance, ok := contracts.Section(string(b), "Provenance")
	if !ok || !strings.Contains(provenance, "evals") || !strings.Contains(provenance, "impl") || !strings.Contains(provenance, "manual") || !strings.Contains(provenance, "0.1.0") || !strings.Contains(provenance, "202") {
		t.Errorf("provenance incomplete: %q", provenance)
	}
	for _, phase := range []string{"evals", "impl"} {
		line := ""
		for _, candidate := range strings.Split(provenance, "\n") {
			if strings.HasPrefix(strings.TrimSpace(candidate), "- "+phase+" ") {
				line = strings.TrimSpace(candidate)
				break
			}
		}
		fields := strings.Fields(line)
		if len(fields) < 3 {
			t.Errorf("%s provenance line misses date: %q", phase, line)
		} else if _, err := time.Parse(time.RFC3339Nano, fields[2]); err != nil {
			t.Errorf("%s provenance line has invalid date %q: %v", phase, fields[2], err)
		}
		if !strings.Contains(line, " agent=manual ") || !strings.Contains(line, " fireweed="+assets.Version()+" ") {
			t.Errorf("%s provenance line misses agent or version: %q", phase, line)
		}
		playbook, err := assets.Playbook(phase)
		if err != nil {
			t.Fatal(err)
		}
		hash := sha256.Sum256([]byte(playbook))
		if line == "" || !strings.Contains(line, "playbook="+hex.EncodeToString(hash[:])) {
			t.Errorf("%s provenance line misses matching playbook hash: %q", phase, line)
		}
		result := evalsResult
		if phase == "impl" {
			result = implResult
		}
		total, passing, failing := 0, 0, 0
		for _, report := range result.Reports {
			for _, c := range report.Cases {
				total++
				switch c.Result {
				case "pass":
					passing++
				case "fail", "error":
					failing++
				}
			}
		}
		if total == 0 || passing == 0 || !strings.Contains(line, fmt.Sprintf("tests=%d failing=%d passing=%d", total, failing, passing)) {
			t.Errorf("%s provenance line has wrong result counts for reports %#v: %q", phase, result.Reports, line)
		}
	}
	runRoot := filepath.Join(root, ".fireweed/runs", state.Run)
	count := 0
	retained := ""
	_ = filepath.Walk(runRoot, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			count++
			b, readErr := os.ReadFile(path)
			if readErr == nil {
				retained += string(b)
			}
		}
		return nil
	})
	if count == 0 || !strings.Contains(retained, "ALP-001") || !strings.Contains(retained, "fixture-hook-output") {
		t.Errorf("reports/output missing under %s (%d files)", runRoot, count)
	}
	configured := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	configuredPrevious := seedProvenance(configured)
	commit(t, configured, "approved prior configured-agent provenance")
	start(t, configured)
	req := request(configured)
	req.Agent = "fixture-agent --model local"
	if _, err := workflow.Done(context.Background(), req); err != nil {
		t.Fatalf("configured-agent evals Done: %v", err)
	}
	assertAppend(configured, pendingCommitChange, configuredPrevious)
	change, err := os.ReadFile(filepath.Join(configured, "changes/0001-manual.md"))
	if err != nil || !strings.Contains(string(change), "fixture-agent --model local") {
		t.Errorf("configured agent omitted from provenance: %v", err)
	}
}

func Test_FLOW_054_StartOperationalPreconditionsHaveP01(t *testing.T) {
	dirty := fixture(t, "", "", "", "", "")
	write(t, dirty, "stray.txt", "untracked\n")
	if _, err := workflow.Start(dirty, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P01: start requires clean worktree") {
		t.Errorf("dirty Start = %v", err)
	}
	queued := fixture(t, "", "", "", "", "")
	write(t, queued, "changes/0000-first.md", "---\nmodule: alpha\n---\n## Motivation\nFirst request.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
	commit(t, queued, "earlier queued change")
	if _, err := workflow.Start(queued, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P01: change is not eligible queue head") {
		t.Errorf("ineligible Start = %v", err)
	}
}

func Test_FLOW_056_StartContentionAndCurrentForeignClaimHaveP02(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	start(t, root)
	if _, err := workflow.Start(other, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02: claim already held:") || !strings.Contains(err.Error(), "update-ref") {
		t.Errorf("contended Start = %v", err)
	}
	if _, err := workflow.Current(other, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02: claim belongs to another worktree") {
		t.Errorf("foreign Current = %v", err)
	}
}

func Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks(t *testing.T) {
	for _, edit := range []string{"author", "runtime", "audit"} {
		t.Run(edit, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			marker := hookMarker(t, root)
			start(t, root)
			p := filepath.Join(root, "changes/0001-manual.md")
			b, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			want := ""
			switch edit {
			case "author":
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "Users need a revised alpha value.", "Unapproved changed purpose.", 1))
				want = "P03: approved author metadata changed"
			case "runtime":
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "phase: evals", "phase: impl", 1))
				want = "P03: approved run or phase changed"
			case "audit":
				write(t, root, "changes/0001-manual.md", strings.Replace(string(b), "phase: evals", "audit_rounds: 1\nphase: evals", 1))
				want = "P03: approved audit state changed"
			}
			before := head(t, root)
			if _, err := verify(root, true); err == nil || !strings.HasPrefix(err.Error(), want) {
				t.Errorf("%s Verify = %v", edit, err)
			}
			if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), want) || head(t, root) != before {
				t.Errorf("%s Done = %v", edit, err)
			}
			if _, err := os.Stat(marker); !os.IsNotExist(err) {
				t.Errorf("hook ran before %s rejection: %v", edit, err)
			}
		})
	}
}

func Test_FLOW_059_AuditNoFindingsErrorsHaveP09(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	before := head(t, root)
	audit(t, root, "No findings\n")
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P09: audit requires Checked IDs and No findings") || head(t, root) != before {
		t.Errorf("malformed No findings = %v", err)
	}
	audit(t, root, "Checked IDs: ALP-002\nNo findings\n")
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P09: audit did not check ALP-001") || head(t, root) != before {
		t.Errorf("incomplete No findings = %v", err)
	}
}

func Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	audit(t, root, "- [untestable] ALP-001: \""+newClause+"\" — cannot observe the value\n")
	before := head(t, root)
	if _, err := done(root); err == nil || !strings.HasPrefix(err.Error(), "P06: change blocked") {
		t.Errorf("blocked Done = %v", err)
	}
	b, err := os.ReadFile(filepath.Join(root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	questions, ok := contracts.Section(string(b), "Questions")
	if head(t, root) == before || !strings.Contains(string(b), "blocked:") || !ok || !strings.Contains(questions, newClause) {
		t.Errorf("blocking state missing from commit: %s", b)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("blocked claim remained held")
	}
}

func Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses(t *testing.T) {
	evalRoot := fixture(t, "", "", "", "", "")
	p := filepath.Join(evalRoot, ".fireweed/eval.sh")
	b, err := os.ReadFile(p)
	if err != nil {
		t.Fatal(err)
	}
	write(t, evalRoot, ".fireweed/eval.sh", strings.ReplaceAll(string(b), `name="ALP-001"`, `name="ALP-002"`))
	commit(t, evalRoot, "hook omits target")
	start(t, evalRoot)
	if _, err := verify(evalRoot, true); err == nil || !strings.HasPrefix(err.Error(), "C10: no coverage for ALP-001") {
		t.Errorf("coverage cause = %v", err)
	}
	before := head(t, evalRoot)
	if _, err := done(evalRoot); err == nil || !strings.HasPrefix(err.Error(), "C10: no coverage for ALP-001") || head(t, evalRoot) != before {
		t.Errorf("Done coverage cause = %v", err)
	}
	missingReport := fixture(t, "", "", "", "", "")
	write(t, missingReport, ".fireweed/eval.sh", "#!/bin/sh\nexit 0\n")
	commit(t, missingReport, "hook omits JUnit report")
	start(t, missingReport)
	if _, err := verify(missingReport, true); err == nil || !strings.HasPrefix(err.Error(), "C11: read JUnit report:") {
		t.Errorf("missing JUnit cause = %v", err)
	}
	failFirst := fixture(t, "", "", "", "", "")
	write(t, failFirst, ".fireweed/unrelated", "yes\n")
	commit(t, failFirst, "unrelated failing case")
	start(t, failFirst)
	write(t, failFirst, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(failFirst, true); err == nil || !strings.HasPrefix(err.Error(), "P2-01: no fail-first case for ALP-001") {
		t.Errorf("fail-first cause = %v", err)
	}
	implFailure := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	start(t, implFailure)
	if _, err := done(implFailure); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	if _, err := verify(implFailure, true); err == nil || !strings.HasPrefix(err.Error(), "E01: failed case") || !strings.Contains(err.Error(), "ALP-001") {
		t.Errorf("impl testcase cause = %v", err)
	}

	zoneRoot := fixture(t, "", "", "", "", "")
	write(t, zoneRoot, "packages/alpha/internal/existing.go", "package internal\n")
	commit(t, zoneRoot, "existing implementation")
	start(t, zoneRoot)
	write(t, zoneRoot, "packages/alpha/internal/existing.go", "package internal // changed in evals\n")
	if _, err := verify(zoneRoot, true); err == nil || !strings.HasPrefix(err.Error(), "Z01:") {
		t.Errorf("zone cause = %v", err)
	}

	boundaryRoot := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	write(t, boundaryRoot, ".fireweed/boundary-fail", "yes\n")
	commit(t, boundaryRoot, "failing boundary fixture")
	start(t, boundaryRoot)
	if _, err := done(boundaryRoot); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, boundaryRoot, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(boundaryRoot, false); err == nil || !strings.HasPrefix(err.Error(), "X01:") || !strings.Contains(err.Error(), "forbidden dependency") {
		t.Errorf("boundary cause = %v", err)
	}
	boundaryHook := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	config, err := os.ReadFile(filepath.Join(boundaryHook, "fireweed.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	write(t, boundaryHook, "fireweed.yaml", strings.Replace(string(config), "boundaries: [sh, .fireweed/boundaries.sh]", "boundaries: [/missing/fireweed-boundary-hook]", 1))
	commit(t, boundaryHook, "unlaunchable boundary hook")
	start(t, boundaryHook)
	if _, err := done(boundaryHook); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, boundaryHook, "packages/alpha/internal/ready", "ready\n")
	snapshot, err := project.Load(boundaryHook)
	if err != nil {
		t.Fatal(err)
	}
	req := request(boundaryHook)
	boundaryDiagnostics, boundaryErr := checks.Boundaries(context.Background(), checks.BoundaryRequest{
		Snapshot: snapshot, Bin: req.Bin, Timeout: req.Timeout,
	})
	if boundaryErr == nil || !strings.Contains(boundaryErr.Error(), "/missing/fireweed-boundary-hook") ||
		len(boundaryDiagnostics) != 1 || boundaryDiagnostics[0].Code != "X02" ||
		!strings.Contains(boundaryDiagnostics[0].Message, "/missing/fireweed-boundary-hook") {
		t.Fatalf("unlaunchable boundary fixture = diagnostics %#v, error %v; want X02 naming missing executable", boundaryDiagnostics, boundaryErr)
	}
	want := boundaryErr.Error()
	if !strings.HasPrefix(want, "X02: ") {
		want = "X02: " + want
	}
	if strings.Count(want, "X02:") != 1 {
		t.Fatalf("boundary cause has repeated X02: %q", want)
	}
	if _, err := verify(boundaryHook, false); err == nil || err.Error() != want {
		t.Errorf("boundary hook cause = %v, want exactly %q", err, want)
	}
}

func Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	config := readFile(t, root, "fireweed.yaml")
	withoutHook := strings.Replace(config, "  boundaries: [sh, .fireweed/boundaries.sh]\n", "", 1)
	if withoutHook == config {
		t.Fatal("fixture did not remove the required boundary hook")
	}
	write(t, root, "fireweed.yaml", withoutHook)
	commit(t, root, "approved fixture without boundary hook")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	current, err := workflow.Current(root, "manual")
	if err != nil || current.Phase != "impl" {
		t.Fatalf("impl state = %#v, %v", current, err)
	}
	req := request(root)
	boundaryDiagnostics, boundaryErr := checks.Boundaries(context.Background(), checks.BoundaryRequest{
		Snapshot: current.Snapshot, Bin: req.Bin, Timeout: req.Timeout,
	})
	if boundaryErr == nil || strings.TrimSpace(boundaryErr.Error()) == "" {
		t.Fatalf("checks boundary cause = %v; want nonempty cause", boundaryErr)
	}
	wantModules := map[string]bool{"alpha": true, "beta": true, "gamma": true}
	if len(boundaryDiagnostics) != len(wantModules) {
		t.Fatalf("boundary diagnostics = %#v; want one per affected module", boundaryDiagnostics)
	}
	for _, d := range boundaryDiagnostics {
		message := strings.ToLower(d.Message)
		if d.Code != "X02" || d.Warning || !strings.Contains(message, "missing") ||
			!strings.Contains(message, "boundar") || !strings.Contains(message, "hook") {
			t.Fatalf("unexpected boundary diagnostic: %#v", d)
		}
		matched := false
		for module := range wantModules {
			if strings.Contains(d.Message, module) {
				delete(wantModules, module)
				matched = true
				break
			}
		}
		if !matched {
			t.Fatalf("boundary diagnostic has no remaining affected module: %#v", d)
		}
	}
	if len(wantModules) != 0 {
		t.Fatalf("boundary diagnostics omitted modules: %v", wantModules)
	}
	result, err := verify(root, false)
	want := boundaryErr.Error()
	if !strings.HasPrefix(want, "X02: ") {
		want = "X02: " + want
	}
	if strings.Count(want, "X02:") != 1 {
		t.Fatalf("checks boundary cause has repeated X02: %q", want)
	}
	if err == nil || err.Error() != want || strings.Count(err.Error(), "X02:") != 1 {
		t.Errorf("full Verify cause = %v; want exactly %q", err, want)
	}
	if !reflect.DeepEqual(result.Diagnostics, boundaryDiagnostics) {
		t.Errorf("full Verify diagnostics = %#v; want %#v", result.Diagnostics, boundaryDiagnostics)
	}
}
package evals_test

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	toAudit(t, root)
	base := head(t, root)
	current, err := workflow.Current(root, "manual")
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct{ body, want string }{
		{"No findings\n", "P09: audit requires Checked IDs and No findings"},
		{"Checked IDs: ALP-002\nNo findings\n", "P09: audit did not check ALP-001"},
	} {
		audit(t, root, tc.body)
		before := readFile(t, root, pendingCommitChange)
		for _, quick := range []bool{true, false} {
			result, err := verify(root, quick)
			if err == nil || !strings.HasPrefix(err.Error(), tc.want) {
				t.Errorf("audit Verify(quick=%t, %q) = phase %q, error %v; want %s", quick, tc.body, result.State.Phase, err, tc.want)
			}
		}
		if got := readFile(t, root, pendingCommitChange); got != before {
			t.Errorf("rejected Verify changed audit submission: %q", got)
		}
		if head(t, root) != base {
			t.Error("rejected Verify committed")
		}
		if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
			t.Errorf("rejected Verify changed claim or phase: %#v, %v", now, err)
		}
	}

	// Findings that Done can process must also pass Verify without being filtered or written.
	audit(t, root, "- ALP-001: \""+newClause+"\" — missing assertion\n- ALP-999: \"Unknown MUST work.\" — invalid finding\n")
	beforeFinding := readFile(t, root, pendingCommitChange)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("finding audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != beforeFinding || head(t, root) != base {
		t.Error("finding audit Verify normalized submission or committed")
	}

	// A complete submission is accepted by Verify without Done's normalization.
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	before := readFile(t, root, pendingCommitChange)
	if result, err := verify(root, true); err != nil || result.State.Phase != "audit" {
		t.Errorf("complete audit Verify = phase %q, error %v", result.State.Phase, err)
	}
	if got := readFile(t, root, pendingCommitChange); got != before || head(t, root) != base {
		t.Error("successful audit Verify wrote the submission or committed")
	}
	if now, err := workflow.Current(root, "manual"); err != nil || now.ClaimOID != current.ClaimOID || now.Phase != "audit" || now.PhaseBase != base {
		t.Errorf("successful Verify changed claim or phase: %#v, %v", now, err)
	}
	if result, err := done(root); err != nil || result.State.Phase != "impl" {
		t.Errorf("complete audit Done = phase %q, error %v", result.State.Phase, err)
	}
}

func Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	marker := hookMarker(t, root)
	auditState := toAudit(t, root)
	if auditState.Phase != "audit" {
		t.Fatalf("held phase = %q, want audit", auditState.Phase)
	}
	if err := os.Remove(marker); err != nil {
		t.Fatalf("clear evals Done hook marker: %v", err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	submission := readFile(t, root, pendingCommitChange)
	forbidden := "packages/alpha/evals/audit_zone_test.go"
	write(t, root, forbidden, "package alpha_test\n")
	base := head(t, root)

	for _, operation := range []struct {
		name string
		call func() (workflow.Result, error)
	}{
		{"quick Verify", func() (workflow.Result, error) { return verify(root, true) }},
		{"full Verify", func() (workflow.Result, error) { return verify(root, false) }},
		{"Done", func() (workflow.Result, error) { return done(root) }},
	} {
		result, err := operation.call()
		if err == nil || !strings.HasPrefix(err.Error(), "Z01:") || !strings.Contains(err.Error(), forbidden) {
			t.Errorf("%s = phase %q, error %v; want Z01 naming %s", operation.name, result.State.Phase, err, forbidden)
		}
		if got := head(t, root); got != base {
			t.Errorf("%s advanced HEAD: %s -> %s", operation.name, base, got)
		}
		if got := readFile(t, root, pendingCommitChange); got != submission {
			t.Errorf("%s changed the valid audit submission", operation.name)
		}
		current, currentErr := workflow.Current(root, "manual")
		if currentErr != nil || current.Phase != "audit" || current.PhaseBase != base || current.ClaimOID != auditState.ClaimOID {
			t.Errorf("%s changed held audit phase or claim: %#v, %v", operation.name, current, currentErr)
		}
		if _, markerErr := os.Stat(marker); !os.IsNotExist(markerErr) {
			t.Errorf("%s ran the eval hook before zone rejection: %v", operation.name, markerErr)
		}
	}
}

func Test_FLOW_061_UnknownLifecycleSelectorsNameValidSlugs(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	request := workflow.Request{Root: root, Change: "ghost"}
	operations := []struct {
		name string
		call func(string) error
	}{
		{"Start", func(s string) error { _, err := workflow.Start(root, s); return err }},
		{"StartFromAttempt", func(s string) error { _, err := workflow.StartFromAttempt(root, s, 1); return err }},
		{"Current", func(s string) error { _, err := workflow.Current(root, s); return err }},
		{"Context", func(s string) error { _, err := workflow.Context(root, s, false); return err }},
		{"Verify", func(s string) error {
			request.Change = s
			_, err := workflow.Verify(context.Background(), request)
			return err
		}},
		{"Done", func(s string) error {
			request.Change = s
			_, err := workflow.Done(context.Background(), request)
			return err
		}},
		{"Release", func(s string) error { return workflow.Release(root, s) }},
		{"Commits", func(s string) error { _, err := workflow.Commits(root, s); return err }},
		{"Abandon", func(s string) error { _, err := workflow.Abandon(root, s); return err }},
	}
	before := head(t, root)
	for _, selector := range []string{"ghost", "0009-ghost.md"} {
		for _, op := range operations {
			err := op.call(selector)
			if err == nil || !strings.HasPrefix(err.Error(), "U02:") || !strings.Contains(err.Error(), selector) || !strings.Contains(err.Error(), "manual") {
				t.Errorf("%s(%q) = %v; want U02 with rejected selector and valid slug manual", op.name, selector, err)
			}
		}
	}
	if head(t, root) != before || git(t, root, "status", "--porcelain=v1", "--untracked-files=all") != "" {
		t.Error("unknown selector changed HEAD or working tree")
	}

	// Commits and Abandon also look up archived documents.
	archive := filepath.Join(root, archivedCommitChange)
	if err := os.MkdirAll(filepath.Dir(archive), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", pendingCommitChange, archivedCommitChange)
	commit(t, root, "archive fixture change")
	for _, op := range operations[len(operations)-2:] {
		err := op.call("ghost")
		if err == nil || !strings.HasPrefix(err.Error(), "U02:") || !strings.Contains(err.Error(), "ghost") || !strings.Contains(err.Error(), "manual") {
			t.Errorf("archived %s(ghost) = %v; want U02 with archive slug manual", op.name, err)
		}
	}
}

exec
/bin/zsh -lc 'cat packages/workflow/evals/{block,resume,gc}_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 458ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func block(t *testing.T, root, reason string) workflow.BlockResult {
	t.Helper()
	r, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: reason})
	if err != nil {
		t.Fatalf("Block: %v", err)
	}
	return r
}

func blockFixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
	t.Helper()
	root := fixture(t, configExtra, changeFront, added, modified, removed)
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
	return root
}

func readFile(t *testing.T, root, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func assertBlockDocumentOnlyAddsReason(t *testing.T, approved, blocked, reason string) {
	t.Helper()
	before, err := contracts.ParseDocument(approved)
	if err != nil {
		t.Fatalf("parse approved change: %v", err)
	}
	after, err := contracts.ParseDocument(blocked)
	if err != nil {
		t.Fatalf("parse blocked change: %v", err)
	}
	if _, exists := before.Frontmatter["blocked"]; exists {
		t.Fatal("fixture change was already blocked")
	}
	want := make(map[string]any, len(before.Frontmatter)+1)
	for key, value := range before.Frontmatter {
		want[key] = value
	}
	want["blocked"] = reason
	if !reflect.DeepEqual(after.Frontmatter, want) {
		t.Errorf("block changed approved frontmatter: got %#v, want %#v", after.Frontmatter, want)
	}
	if strings.TrimSpace(after.Body) != strings.TrimSpace(before.Body) {
		t.Errorf("block changed approved Markdown body: got %q, want %q", after.Body, before.Body)
	}
}

func assertBlockReleased(t *testing.T, root string, r workflow.BlockResult, reason string) {
	t.Helper()
	if r.State.Phase == "" || r.State.ClaimOID != "" || r.State.Head != head(t, root) || r.State.PhaseBase != head(t, root) {
		t.Errorf("returned blocked state = %#v", r.State)
	}
	if got := r.State.Change.Document.Frontmatter["blocked"]; got != reason {
		t.Errorf("returned blocked reason = %#v, want %q", got, reason)
	}
	if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
		t.Errorf("claim still exists: %s", got)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("Current still found a held claim")
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Errorf("checkout is not clean: %s", got)
	}
	if index, tree := git(t, root, "write-tree"), git(t, root, "rev-parse", "HEAD^{tree}"); index != tree {
		t.Errorf("index tree %s differs from blocked HEAD tree %s", index, tree)
	}
}

func Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState(t *testing.T) {
	root := blockFixture(t, "", "", "### new:extra\nExtra MUST be checked.\n", "", "")
	write(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	write(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	commit(t, root, "fixture eval files")
	started := start(t, root)
	if started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("invalid held evals fixture: %#v", started)
	}
	approved := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(approved, "### ALP-003") || !strings.Contains(approved, started.Run) {
		t.Fatalf("assigned IDs/run absent from fixture: %s", approved)
	}
	base := head(t, root)
	prior := "refs/fireweed/attempts/manual/1"
	git(t, root, "update-ref", prior, base)
	write(t, root, "packages/alpha/evals/current_test.go", "staged modified eval\n")
	git(t, root, "add", "packages/alpha/evals/current_test.go")
	if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/evals/staged_test.go", "staged new eval\n")
	git(t, root, "add", "packages/alpha/evals/staged_test.go")
	write(t, root, "packages/alpha/evals/untracked_test.go", "untracked eval\n")
	if got := git(t, root, "status", "--porcelain", "-uall"); !strings.Contains(got, "M  packages/alpha/evals/current_test.go") || !strings.Contains(got, " D packages/alpha/evals/retired_test.go") || !strings.Contains(got, "A  packages/alpha/evals/staged_test.go") || !strings.Contains(got, "?? packages/alpha/evals/untracked_test.go") {
		t.Fatalf("fixture lacks staged, deleted or untracked work: %s", got)
	}
	if got := git(t, root, "show", ":packages/alpha/evals/current_test.go"); got != "staged modified eval" {
		t.Fatalf("fixture index was not staged: %q", got)
	}
	reason := "verification attempts exhausted"
	r := block(t, root, reason)
	if r.State.Phase != "evals" || r.State.Run != started.Run || len(r.State.Change.Added) != 1 || r.State.Change.Added[0].ID != "ALP-003" {
		t.Errorf("returned state lost held evals phase, run or assigned ID: %#v", r.State)
	}
	wantRef := "refs/fireweed/attempts/manual/2"
	if r.AttemptRef != wantRef || r.AttemptOID == "" || git(t, root, "rev-parse", wantRef) != r.AttemptOID || git(t, root, "cat-file", "-t", r.AttemptOID) != "commit" {
		t.Errorf("attempt ref/OID = %q %q, want %q at commit", r.AttemptRef, r.AttemptOID, wantRef)
	}
	if got := git(t, root, "rev-parse", prior); got != base {
		t.Errorf("prior attempt ref changed: %s, want %s", got, base)
	}
	for path, want := range map[string]string{
		"packages/alpha/evals/current_test.go":   "staged modified eval",
		"packages/alpha/evals/staged_test.go":    "staged new eval",
		"packages/alpha/evals/untracked_test.go": "untracked eval",
	} {
		if got := git(t, root, "show", r.AttemptOID+":"+path); got != want {
			t.Errorf("saved %s = %q, want %q", path, got, want)
		}
	}
	if got := git(t, root, "ls-tree", "--name-only", r.AttemptOID, "packages/alpha/evals/retired_test.go"); got != "" {
		t.Errorf("saved attempt retained deleted file: %s", got)
	}
	if got := git(t, root, "rev-parse", "HEAD^"); got != base {
		t.Errorf("block parent = %s, want phase base %s", got, base)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != "changes/0001-manual.md" {
		t.Errorf("block commit contains step work: %s", got)
	}
	blocked := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(blocked, "### ALP-003") || !strings.Contains(blocked, started.Run) || !strings.Contains(blocked, "phase: evals") || !strings.Contains(blocked, "blocked: "+reason) || !strings.Contains(blocked, "Users need a revised alpha value.") {
		t.Errorf("block commit lost assigned/author state: %s", blocked)
	}
	assertBlockDocumentOnlyAddsReason(t, approved, blocked, reason)
	if got := readFile(t, root, "packages/alpha/evals/current_test.go"); got != "original eval\n" {
		t.Errorf("tracked checkout not restored: %q", got)
	}
	if got := readFile(t, root, "packages/alpha/evals/retired_test.go"); got != "retired eval\n" {
		t.Errorf("deleted checkout not restored: %q", got)
	}
	for _, path := range []string{"packages/alpha/evals/staged_test.go", "packages/alpha/evals/untracked_test.go"} {
		if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
			t.Errorf("step file %s remains after block: %v", path, err)
		}
	}
	assertBlockReleased(t, root, r, reason)
}

func Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit(t *testing.T) {
	root := blockFixture(t, "", "", "", "", "")
	start(t, root)
	write(t, root, "packages/alpha/evals/committed_test.go", "committed eval\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done: %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("audit Done: %#v, %v", audited, err)
	}
	base := head(t, root)
	committedChange := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(committedChange, "No findings; checked ALP-001") || readFile(t, root, "packages/alpha/evals/committed_test.go") != "committed eval\n" {
		t.Fatal("fixture lacks committed eval/audit content")
	}
	write(t, root, "packages/alpha/internal/attempt.go", "unfinished implementation\n")
	r := block(t, root, "implementation attempts exhausted")
	if r.State.Phase != "impl" || git(t, root, "rev-parse", "HEAD^") != base {
		t.Errorf("impl block base/phase = %s, %#v", base, r.State)
	}
	if got := git(t, root, "show", r.AttemptOID+":packages/alpha/internal/attempt.go"); got != "unfinished implementation" {
		t.Errorf("impl attempt work = %q", got)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != "changes/0001-manual.md" {
		t.Errorf("impl block contains work: %s", got)
	}
	if got := readFile(t, root, "changes/0001-manual.md"); !strings.Contains(got, "No findings; checked ALP-001") || !strings.Contains(got, "blocked: implementation attempts exhausted") {
		t.Errorf("impl block lost committed audit state: %s", got)
	}
	assertBlockDocumentOnlyAddsReason(t, committedChange, readFile(t, root, "changes/0001-manual.md"), "implementation attempts exhausted")
	if got := readFile(t, root, "packages/alpha/evals/committed_test.go"); got != "committed eval\n" {
		t.Errorf("impl block lost committed eval: %q", got)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/alpha/internal/attempt.go")); !os.IsNotExist(err) {
		t.Errorf("impl work remains: %v", err)
	}
	assertBlockReleased(t, root, r, "implementation attempts exhausted")
}

func Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead(t *testing.T) {
	t.Run("no claim", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		before := head(t, root)
		if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("Block accepted change without a claim")
		}
		if head(t, root) != before || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
			t.Error("unclaimed Block changed HEAD or attempts")
		}
	})
	t.Run("foreign claim", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		started := start(t, root)
		before := head(t, other)
		if _, err := workflow.Block(workflow.BlockRequest{Root: other, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("foreign worktree blocked owner's claim")
		}
		if head(t, other) != before {
			t.Error("foreign Block moved HEAD")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
			t.Errorf("owner's claim changed: %#v, %v", current, err)
		}
	})
	for _, edit := range []string{"author", "phase", "audit"} {
		t.Run(edit, func(t *testing.T) {
			root := blockFixture(t, "", "", "", "", "")
			started := start(t, root)
			path := "changes/0001-manual.md"
			before := readFile(t, root, path)
			switch edit {
			case "author":
				write(t, root, path, strings.Replace(before, "Users need a revised alpha value.", "Unapproved purpose.", 1))
			case "phase":
				write(t, root, path, strings.Replace(before, "phase: evals", "phase: impl", 1))
			case "audit":
				write(t, root, path, strings.Replace(before, "phase: evals", "audit_rounds: 1\nphase: evals", 1))
			}
			if before == readFile(t, root, path) {
				t.Fatal("fixture did not alter approved state")
			}
			base := head(t, root)
			if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
				t.Errorf("Block accepted changed %s", edit)
			}
			if head(t, root) != base || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
				t.Errorf("rejected Block committed %s work", edit)
			}
			if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
				t.Errorf("rejected Block released claim: %#v, %v", current, err)
			}
		})
	}
	t.Run("changed HEAD", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		started := start(t, root)
		write(t, root, "packages/alpha/evals/committed_test.go", "premature commit\n")
		git(t, root, "add", "packages/alpha/evals/committed_test.go")
		git(t, root, "commit", "-m", "fixture head advance")
		advanced := head(t, root)
		if advanced == started.PhaseBase {
			t.Fatal("fixture did not advance HEAD")
		}
		if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("Block accepted changed HEAD")
		}
		if head(t, root) != advanced || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
			t.Error("stale-HEAD Block changed refs or HEAD")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
			t.Errorf("stale-HEAD Block released claim: %#v, %v", current, err)
		}
	})
}
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// These fixtures use the public lifecycle to create a saved commit and release
// its claim. An author then commits the unblock before any resume request.
func unblockAttempt(t *testing.T, root string) {
	t.Helper()
	path := "changes/0001-manual.md"
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	if _, blocked := doc.Frontmatter["blocked"]; !blocked {
		t.Fatal("fixture has no committed block to remove")
	}
	delete(doc.Frontmatter, "blocked")
	doc.Body = contracts.ReplaceSection(doc.Body, "Motivation", "\nAuthor approved a revised purpose.\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Interface", "\nAuthor approved the current interface.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	commit(t, root, "author unblocks saved attempt")
}

func savedEvalsAttempt(t *testing.T) (string, workflow.State, workflow.BlockResult) {
	t.Helper()
	root := blockFixture(t, "", "", "### new:extra\nExtra MUST be checked.\n", "", "")
	write(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	write(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	commit(t, root, "fixture eval files")
	started := start(t, root)
	if started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("invalid assigned evals fixture: %#v", started)
	}
	write(t, root, "packages/alpha/evals/current_test.go", "saved modified eval\n")
	if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/evals/new_test.go", "saved new eval\n")
	path := "changes/0001-manual.md"
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", "\nSaved worker decision.\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Questions", "\nSaved worker question.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	saved := block(t, root, "eval attempts exhausted")
	if saved.AttemptRef != "refs/fireweed/attempts/manual/1" || saved.AttemptOID == "" {
		t.Fatalf("Block did not save attempt 1: %#v", saved)
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Fatalf("Block left checkout dirty: %s", got)
	}
	return root, started, saved
}

func assertNoResumeMutation(t *testing.T, root, beforeHead, beforeChange string) {
	t.Helper()
	if head(t, root) != beforeHead || readFile(t, root, "changes/0001-manual.md") != beforeChange {
		t.Error("rejected resume changed HEAD or change document")
	}
	if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
		t.Errorf("rejected resume claimed change: %s", got)
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Errorf("rejected resume dirtied checkout: %s", got)
	}
}

func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptRestoresEvalsStepAndRetainsCurrentApproval(t *testing.T) {
	root, started, saved := savedEvalsAttempt(t)
	write(t, root, "later-independent.txt", "later independent commit\n")
	commit(t, root, "later independent work")
	unblockAttempt(t, root)
	beforeHead := head(t, root)
	before := readFile(t, root, "changes/0001-manual.md")
	if strings.Contains(before, "Saved worker decision.") || strings.Contains(before, "Saved worker question.") {
		t.Fatal("blocked fixture leaked saved worker body")
	}
	got, err := workflow.StartFromAttempt(root, "manual", 1)
	if err != nil {
		t.Fatalf("StartFromAttempt(%s): %v", saved.AttemptRef, err)
	}
	if got.Run != started.Run || got.Phase != started.Phase || len(got.Change.Added) != 1 || got.Change.Added[0].ID != started.Change.Added[0].ID {
		t.Errorf("resume changed run, phase or assigned ID: %#v", got)
	}
	if got.Head != beforeHead || got.PhaseBase != beforeHead || head(t, root) != beforeHead {
		t.Errorf("resume changed HEAD or phase base: state=%#v, HEAD=%s", got, head(t, root))
	}
	current, currentErr := workflow.Current(root, "manual")
	if got.ClaimOID == "" || currentErr != nil || current.ClaimOID != got.ClaimOID || got.Snapshot.Root != root {
		t.Errorf("resume did not return an owned held claim: state=%#v, Current=%#v, %v", got, current, currentErr)
	}
	if text := readFile(t, root, "packages/alpha/evals/current_test.go"); text != "saved modified eval\n" {
		t.Errorf("modified eval = %q", text)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/alpha/evals/retired_test.go")); !os.IsNotExist(err) {
		t.Errorf("saved deletion was not restored: %v", err)
	}
	if text := readFile(t, root, "packages/alpha/evals/new_test.go"); text != "saved new eval\n" {
		t.Errorf("new eval = %q", text)
	}
	if text := readFile(t, root, "later-independent.txt"); text != "later independent commit\n" {
		t.Errorf("later independent work = %q", text)
	}
	resumed, err := contracts.ParseDocument(readFile(t, root, "changes/0001-manual.md"))
	if err != nil {
		t.Fatal(err)
	}
	approved, err := contracts.ParseDocument(before)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Removed decisions"} {
		want, _ := contracts.Section(approved.Body, name)
		actual, _ := contracts.Section(resumed.Body, name)
		if actual != want {
			t.Errorf("current approved %s changed: got %q, want %q", name, actual, want)
		}
	}
	if !reflect.DeepEqual(resumed.Frontmatter, approved.Frontmatter) {
		t.Errorf("current runtime/author frontmatter changed: got %#v, want %#v", resumed.Frontmatter, approved.Frontmatter)
	}
	for name, want := range map[string]string{"Decisions": "Saved worker decision.", "Questions": "Saved worker question."} {
		body, ok := contracts.Section(resumed.Body, name)
		if !ok || !strings.Contains(body, want) {
			t.Errorf("saved worker %s absent: %q", name, body)
		}
	}
}

func Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance(t *testing.T) {
	root := blockFixture(t, "", "", "", "", "")
	start(t, root)
	write(t, root, "packages/alpha/evals/committed_test.go", "committed eval\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done: %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	impl, err := done(root)
	if err != nil || impl.State.Phase != "impl" {
		t.Fatalf("audit Done: %#v, %v", impl, err)
	}
	write(t, root, "packages/alpha/internal/attempt.go", "saved implementation\n")
	path := "changes/0001-manual.md"
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", "\nSaved impl decision.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	saved := block(t, root, "impl attempts exhausted")
	if saved.AttemptRef != "refs/fireweed/attempts/manual/1" {
		t.Fatalf("wrong impl attempt ref: %s", saved.AttemptRef)
	}
	unblockAttempt(t, root)
	beforeHead := head(t, root)
	before, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	provenance, hasProvenance := contracts.Section(before.Body, "Provenance")
	auditBody, hasAudit := contracts.Section(before.Body, "Audit")
	if !hasProvenance || strings.TrimSpace(provenance) == "" || !hasAudit || !strings.Contains(auditBody, "No findings; checked ALP-001") {
		t.Fatalf("fixture lacks committed controller provenance or audit: %#v", before)
	}
	got, err := workflow.StartFromAttempt(root, "manual", 1)
	if err != nil {
		t.Fatalf("impl StartFromAttempt: %v", err)
	}
	if got.Run != impl.State.Run || got.Phase != "impl" || got.Head != beforeHead || head(t, root) != beforeHead || got.ClaimOID == "" {
		t.Errorf("impl resume lost run/phase/claim or changed HEAD: %#v", got)
	}
	if text := readFile(t, root, "packages/alpha/evals/committed_test.go"); text != "committed eval\n" {
		t.Errorf("committed eval changed: %q", text)
	}
	if text := readFile(t, root, "packages/alpha/internal/attempt.go"); text != "saved implementation\n" {
		t.Errorf("impl step work = %q", text)
	}
	resumed, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"Audit", "Provenance", "Motivation", "Interface"} {
		want, _ := contracts.Section(before.Body, name)
		actual, _ := contracts.Section(resumed.Body, name)
		if actual != want {
			t.Errorf("committed/current %s changed: got %q, want %q", name, actual, want)
		}
	}
	if !reflect.DeepEqual(resumed.Frontmatter, before.Frontmatter) {
		t.Errorf("impl resume changed runtime/author frontmatter: %#v, want %#v", resumed.Frontmatter, before.Frontmatter)
	}
	if body, _ := contracts.Section(resumed.Body, "Decisions"); !strings.Contains(body, "Saved impl decision.") {
		t.Errorf("saved impl decision missing: %q", body)
	}
}

func Test_FLOW_074_StartFromAttemptSelectsRequestedSavedRef(t *testing.T) {
	root, _, first := savedEvalsAttempt(t)
	unblockAttempt(t, root)
	start(t, root)
	write(t, root, "packages/alpha/evals/current_test.go", "second modified eval\n")
	write(t, root, "packages/alpha/evals/new_test.go", "second new eval\n")
	second := block(t, root, "second eval attempt exhausted")
	if second.AttemptRef != "refs/fireweed/attempts/manual/2" || second.AttemptOID == "" || second.AttemptOID == first.AttemptOID {
		t.Fatalf("Block did not save a distinct attempt 2: first=%#v, second=%#v", first, second)
	}
	if got := git(t, root, "show", first.AttemptRef+":packages/alpha/evals/current_test.go"); got != "saved modified eval" {
		t.Fatalf("attempt 1 content = %q", got)
	}
	if got := git(t, root, "show", second.AttemptRef+":packages/alpha/evals/current_test.go"); got != "second modified eval" {
		t.Fatalf("attempt 2 content = %q", got)
	}
	unblockAttempt(t, root)
	if _, err := workflow.StartFromAttempt(root, "manual", 2); err != nil {
		t.Fatalf("StartFromAttempt(%s): %v", second.AttemptRef, err)
	}
	if got := readFile(t, root, "packages/alpha/evals/current_test.go"); got != "second modified eval\n" {
		t.Errorf("requested attempt 2 modified eval = %q", got)
	}
	if got := readFile(t, root, "packages/alpha/evals/new_test.go"); got != "second new eval\n" {
		t.Errorf("requested attempt 2 new eval = %q", got)
	}
}

func Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval(t *testing.T) {
	root := blockFixture(t, "", "", "", "", "")
	auditState := toAudit(t, root)
	path := "changes/0001-manual.md"
	committed, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	committedAudit, _ := contracts.Section(committed.Body, "Audit")
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	worker, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	savedAudit, ok := contracts.Section(worker.Body, "Audit")
	if !ok || savedAudit == committedAudit || !strings.Contains(savedAudit, "No findings") {
		t.Fatalf("fixture lacks distinct worker Audit body: saved=%q, committed=%q", savedAudit, committedAudit)
	}
	saved := block(t, root, "audit attempt exhausted")
	if saved.AttemptRef != "refs/fireweed/attempts/manual/1" || saved.AttemptOID == "" {
		t.Fatalf("Block did not save audit attempt: %#v", saved)
	}
	unblockAttempt(t, root)
	before, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	currentAudit, _ := contracts.Section(before.Body, "Audit")
	currentProvenance, hasProvenance := contracts.Section(before.Body, "Provenance")
	if currentAudit == savedAudit || !hasProvenance || strings.TrimSpace(currentProvenance) == "" {
		t.Fatalf("fixture lacks separate committed Audit or controller provenance: Audit=%q, Provenance=%q", currentAudit, currentProvenance)
	}
	got, err := workflow.StartFromAttempt(root, "manual", 1)
	if err != nil {
		t.Fatalf("audit StartFromAttempt(%s): %v", saved.AttemptRef, err)
	}
	if got.Phase != auditState.Phase || got.Run != auditState.Run {
		t.Errorf("audit resume changed phase or run: %#v", got)
	}
	resumed, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	if body, _ := contracts.Section(resumed.Body, "Audit"); body != savedAudit {
		t.Errorf("restored Audit = %q, want saved worker body %q", body, savedAudit)
	}
	for _, name := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Removed decisions", "Provenance"} {
		want, _ := contracts.Section(before.Body, name)
		actual, _ := contracts.Section(resumed.Body, name)
		if actual != want {
			t.Errorf("current %s changed: got %q, want %q", name, actual, want)
		}
	}
	if !reflect.DeepEqual(resumed.Frontmatter, before.Frontmatter) {
		t.Errorf("audit resume changed current author/runtime frontmatter: got %#v, want %#v", resumed.Frontmatter, before.Frontmatter)
	}
}

func Test_FLOW_073_StartFromAttemptRequiresCleanEligibleUnblockedAndUnclaimedChange(t *testing.T) {
	t.Run("dirty", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		unblockAttempt(t, root)
		write(t, root, "uncommitted.txt", "dirty\n")
		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("dirty resume = %v, want P01", err)
		}
		if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
			t.Errorf("dirty resume created claim: %s", got)
		}
	})
	t.Run("still blocked", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		beforeHead, before := head(t, root), readFile(t, root, "changes/0001-manual.md")
		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("blocked resume = %v, want P01", err)
		}
		assertNoResumeMutation(t, root, beforeHead, before)
	})
	t.Run("not first eligible", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		unblockAttempt(t, root)
		write(t, root, "changes/0000-first.md", "---\nmodule: alpha\n---\n## Motivation\nEarlier approved work.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n")
		commit(t, root, "earlier eligible change")
		beforeHead, before := head(t, root), readFile(t, root, "changes/0001-manual.md")
		if _, err := workflow.StartFromAttempt(root, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("waiting resume = %v, want P01", err)
		}
		assertNoResumeMutation(t, root, beforeHead, before)
	})
	t.Run("foreign claim", func(t *testing.T) {
		root, _, _ := savedEvalsAttempt(t)
		unblockAttempt(t, root)
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		owner := start(t, root)
		beforeHead, before := head(t, other), readFile(t, other, "changes/0001-manual.md")
		if _, err := workflow.StartFromAttempt(other, "manual", 1); err == nil || !strings.HasPrefix(err.Error(), "P02") {
			t.Errorf("foreign claim resume = %v, want P02", err)
		}
		if head(t, other) != beforeHead || readFile(t, other, "changes/0001-manual.md") != before {
			t.Error("foreign resume changed other worktree")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != owner.ClaimOID {
			t.Errorf("foreign resume changed owner's claim: %#v, %v", current, err)
		}
	})
}

func Test_FLOW_076_StartFromAttemptRejectsInvalidOrMissingAttemptBeforeMutation(t *testing.T) {
	root, _, _ := savedEvalsAttempt(t)
	unblockAttempt(t, root)
	beforeHead, before := head(t, root), readFile(t, root, "changes/0001-manual.md")
	for _, attempt := range []int{0, -1, 2} {
		if _, err := workflow.StartFromAttempt(root, "manual", attempt); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("attempt %d = %v, want P01", attempt, err)
		}
		assertNoResumeMutation(t, root, beforeHead, before)
	}
}
package evals_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow"
)

var gcNow = time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)

func gcCall(t *testing.T, root string, now time.Time) workflow.GCResult {
	t.Helper()
	r, err := workflow.GC(context.Background(), workflow.GCRequest{Root: root, Now: now})
	if err != nil {
		t.Fatalf("GC: %v", err)
	}
	return r
}

func gcRef(t *testing.T, root, name, oid string) {
	t.Helper()
	if err := gitstore.CASRef(root, name, "", oid); err != nil {
		t.Fatalf("create %s: %v", name, err)
	}
}

func gcClaim(t *testing.T, root, name, worktree string) {
	t.Helper()
	data, err := json.Marshal(map[string]string{"worktree": worktree, "change": name})
	if err != nil {
		t.Fatal(err)
	}
	oid, err := gitstore.WriteBlob(root, data)
	if err != nil {
		t.Fatal(err)
	}
	gcRef(t, root, "refs/fireweed/claims/"+name, oid)
}

func gcPresent(t *testing.T, root, name string, want bool) {
	t.Helper()
	_, found, err := gitstore.Ref(root, name)
	if err != nil || found != want {
		t.Errorf("ref %s present=%t, error=%v; want present=%t", name, found, err, want)
	}
}

func gcNames(t *testing.T, root string) []string {
	t.Helper()
	out := git(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed")
	if out == "" {
		return nil
	}
	return strings.Split(out, "\n")
}

func gcPruned(t *testing.T, got, want []string) {
	t.Helper()
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GC Pruned = %v, want sorted %v", got, want)
	}
}

func Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	linked := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", linked, "HEAD")
	gcClaim(t, root, "main-live", root)
	gcClaim(t, root, "linked-live", linked)
	missing := filepath.Join(t.TempDir(), "removed")
	gcClaim(t, root, "missing", missing)
	unregistered := filepath.Join(t.TempDir(), "unregistered")
	if err := os.Mkdir(unregistered, 0o755); err != nil {
		t.Fatal(err)
	}
	gcClaim(t, root, "unregistered", unregistered)
	foreign, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	git(t, foreign, "init")
	if got := git(t, foreign, "worktree", "list", "--porcelain"); !strings.Contains(got, "worktree "+foreign) {
		t.Fatalf("foreign repository does not register its worktree: %s", got)
	}
	gcClaim(t, root, "foreign", foreign)
	gcRef(t, root, "refs/fireweed/other/keep", head(t, root))

	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, []string{"refs/fireweed/claims/missing", "refs/fireweed/claims/unregistered", "refs/fireweed/claims/foreign"})
	for _, name := range []string{"missing", "unregistered", "foreign"} {
		gcPresent(t, root, "refs/fireweed/claims/"+name, false)
	}
	for _, name := range []string{"main-live", "linked-live"} {
		gcPresent(t, root, "refs/fireweed/claims/"+name, true)
	}
	gcPresent(t, root, "refs/fireweed/other/keep", true)
}

func Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
	write(t, root, "changes/0003-pending.md", "---\nmodule: alpha\n---\n")
	oid := head(t, root)
	archived := []string{"refs/fireweed/attempts/finished/1", "refs/fireweed/attempts/finished/2"}
	for _, name := range append(archived, "refs/fireweed/attempts/pending/1", "refs/fireweed/attempts/unknown/1", "refs/fireweed/attempts/finished/not-an-attempt") {
		gcRef(t, root, name, oid)
	}
	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, archived)
	for _, name := range archived {
		gcPresent(t, root, name, false)
	}
	for _, name := range []string{"refs/fireweed/attempts/pending/1", "refs/fireweed/attempts/unknown/1", "refs/fireweed/attempts/finished/not-an-attempt"} {
		gcPresent(t, root, name, true)
	}
}

func Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate(t *testing.T) {
	for _, tc := range []struct {
		name, config      string
		old, equal, newer string
	}{
		{"configured", "gc:\n  drill_max_age: 24h\n", "20261103T115959.999999999", "20261103T120000.000000000", "20261103T120000.000000001"},
		{"default", "", "20261005T115959.999999999", "20261005T120000.000000000", "20261005T120000.000000001"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, tc.config, "", "", "", "")
			oid := head(t, root)
			prefix := "refs/fireweed/drills/alpha/"
			for _, date := range []string{tc.old, tc.equal, tc.newer} {
				gcRef(t, root, prefix+date, oid)
			}
			gcRef(t, root, prefix+"unrecognized", oid)
			gcRef(t, root, "refs/fireweed/drills/alpha/not-a-date/extra", oid)
			r := gcCall(t, root, gcNow)
			gcPruned(t, r.Pruned, []string{prefix + tc.old})
			gcPresent(t, root, prefix+tc.old, false)
			for _, name := range []string{prefix + tc.equal, prefix + tc.newer, prefix + "unrecognized", "refs/fireweed/drills/alpha/not-a-date/extra"} {
				gcPresent(t, root, name, true)
			}
		})
	}
	// With zero Now, a native date far in the past is expired against any
	// current clock, while a future date stays live.
	root := fixture(t, "", "", "", "", "")
	old := "refs/fireweed/drills/alpha/20200101T000000.000000000"
	future := "refs/fireweed/drills/alpha/20990101T000000.000000000"
	gcRef(t, root, old, head(t, root))
	gcRef(t, root, future, head(t, root))
	gcPruned(t, gcCall(t, root, time.Time{}).Pruned, []string{old})
	gcPresent(t, root, future, true)
}

func Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning(t *testing.T) {
	for _, value := range []string{"0s", "-1h", "invalid"} {
		t.Run(value, func(t *testing.T) {
			root := fixture(t, "gc:\n  drill_max_age: "+value+"\n", "", "", "", "")
			claim := "refs/fireweed/claims/stale"
			old := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			gcClaim(t, root, "stale", filepath.Join(t.TempDir(), "gone"))
			gcRef(t, root, old, head(t, root))
			before := gcNames(t, root)
			r, err := workflow.GC(context.Background(), workflow.GCRequest{Root: root, Now: gcNow})
			if err == nil || !strings.HasPrefix(err.Error(), "C01:") {
				t.Errorf("invalid drill_max_age %q: result=%#v error=%v; want C01", value, r, err)
			}
			if got := gcNames(t, root); !reflect.DeepEqual(got, before) {
				t.Errorf("invalid setting pruned refs: before=%v after=%v", before, got)
			}
			gcPresent(t, root, claim, true)
		})
	}
}

func Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
	write(t, root, "working.txt", "untracked work\n")
	write(t, root, "SYSTEM.md", "staged work\n")
	git(t, root, "add", "SYSTEM.md")
	gcClaim(t, root, "z-stale", filepath.Join(t.TempDir(), "missing"))
	gcRef(t, root, "refs/fireweed/attempts/finished/1", head(t, root))
	gcRef(t, root, "refs/fireweed/drills/alpha/20200101T000000.000000000", head(t, root))
	gcRef(t, root, "refs/fireweed/other/keep", head(t, root))
	want := []string{"refs/fireweed/claims/z-stale", "refs/fireweed/attempts/finished/1", "refs/fireweed/drills/alpha/20200101T000000.000000000"}
	beforeHead := head(t, root)
	beforeIndex := git(t, root, "write-tree")
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeWork := readFile(t, root, "working.txt")
	beforeStaged := readFile(t, root, "SYSTEM.md")
	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, want)
	for _, name := range want {
		gcPresent(t, root, name, false)
	}
	gcPresent(t, root, "refs/fireweed/other/keep", true)
	if got := head(t, root); got != beforeHead {
		t.Errorf("GC moved HEAD: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "write-tree"); got != beforeIndex {
		t.Errorf("GC changed index tree: %s -> %s", beforeIndex, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("GC changed worktree status: %q -> %q", beforeStatus, got)
	}
	if readFile(t, root, "working.txt") != beforeWork || readFile(t, root, "SYSTEM.md") != beforeStaged {
		t.Error("GC changed working files")
	}
	if second := gcCall(t, root, gcNow); len(second.Pruned) != 0 {
		t.Errorf("second GC reported refs it did not delete: %v", second.Pruned)
	}
}

exec
/bin/zsh -lc "cat packages/workflow/evals/{commits,abandon}_test.go; sed -n '1,330p' packages/workflow/evals/status_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 475ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

const pendingCommitChange = "changes/0001-manual.md"
const archivedCommitChange = "changes/archive/0001-manual.md"

func commitPhase(t *testing.T, root, path, old, next, message string) string {
	t.Helper()
	before := readFile(t, root, path)
	var after string
	if old == "" {
		after = strings.Replace(before, "module: alpha\n", "module: alpha\nphase: "+next+"\n", 1)
	} else {
		after = strings.Replace(before, "phase: "+old+"\n", "phase: "+next+"\n", 1)
	}
	if after == before {
		t.Fatalf("fixture phase %q not found in %s", old, path)
	}
	write(t, root, path, after)
	commit(t, root, message)
	return head(t, root)
}

func assertCommits(t *testing.T, root, selector string, want []string) {
	t.Helper()
	got, err := workflow.Commits(root, selector)
	if err != nil {
		t.Fatalf("Commits(%q): %v", selector, err)
	}
	for _, oid := range got {
		if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) {
			t.Errorf("Commits(%q) returned non-full SHA-1 OID %q", selector, oid)
		}
	}
	if !(len(got) == 0 && len(want) == 0) && !reflect.DeepEqual(got, want) {
		t.Errorf("Commits(%q) = %v, want newest-first %v", selector, got, want)
	}
}

func Test_FLOW_087_FLOW_088_CommitsFollowPendingAndArchivePhaseContents(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	author := head(t, root)
	write(t, root, "unrelated.txt", "unrelated work\n")
	commit(t, root, "phase: evals")
	unrelated := head(t, root)
	evals := commitPhase(t, root, pendingCommitChange, "", "evals", "same message")
	change := readFile(t, root, pendingCommitChange)
	write(t, root, pendingCommitChange, strings.Replace(change, "Users need a revised alpha value.", "Users need a revised alpha value with detail.", 1))
	commit(t, root, "phase: audit")
	noPhaseChange := head(t, root)
	audit := commitPhase(t, root, pendingCommitChange, "evals", "audit", "same message")
	impl := commitPhase(t, root, pendingCommitChange, "audit", "impl", "same message")
	git(t, root, "checkout", "-b", "side")
	unreachable := commitPhase(t, root, pendingCommitChange, "impl", "applied", "same message")
	git(t, root, "checkout", "main")
	got, err := workflow.Commits(root, pendingCommitChange)
	if err != nil {
		t.Fatalf("Commits after side branch: %v", err)
	}
	for _, oid := range got {
		if oid == unreachable {
			t.Errorf("Commits included unreachable phase change %s", unreachable)
		}
	}
	for _, selector := range []string{pendingCommitChange, "0001-manual.md", "manual"} {
		assertCommits(t, root, selector, []string{impl, audit, evals})
	}

	if err := os.MkdirAll(filepath.Join(root, "changes", "archive"), 0o755); err != nil {
		t.Fatalf("create archive directory: %v", err)
	}
	git(t, root, "mv", pendingCommitChange, archivedCommitChange)
	archived := readFile(t, root, archivedCommitChange)
	write(t, root, archivedCommitChange, strings.Replace(archived, "phase: impl\n", "phase: applied\n", 1))
	commit(t, root, "same message")
	applied := head(t, root)
	for _, selector := range []string{archivedCommitChange, "0001-manual.md", "manual"} {
		assertCommits(t, root, selector, []string{applied, impl, audit, evals})
	}
	for _, excluded := range []string{author, unrelated, noPhaseChange} {
		if excluded == evals || excluded == audit || excluded == impl || excluded == applied {
			t.Fatalf("fixture OID collision: %s", excluded)
		}
	}
}

func Test_FLOW_087_CommitsUseFirstParentOfMerge(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	evals := commitPhase(t, root, pendingCommitChange, "", "evals", "same message")
	git(t, root, "branch", "alternate", evals)
	mainPhase := commitPhase(t, root, pendingCommitChange, "evals", "impl", "same message")
	git(t, root, "checkout", "alternate")
	branchPhase := commitPhase(t, root, pendingCommitChange, "evals", "audit", "same message")
	git(t, root, "checkout", "main")
	git(t, root, "merge", "--no-ff", "-s", "ours", "-m", "phase: applied", "alternate")
	merge := head(t, root)
	got, err := workflow.Commits(root, "manual")
	if err != nil {
		t.Fatalf("Commits after merge: %v", err)
	}
	want := map[string]bool{evals: true, mainPhase: true, branchPhase: true}
	if len(got) != len(want) {
		t.Errorf("Commits after merge = %v, want three phase changes", got)
	}
	for _, oid := range got {
		if !want[oid] {
			t.Errorf("Commits included %s; merge %s has no phase difference from its first parent", oid, merge)
		}
		delete(want, oid)
	}
	for oid := range want {
		t.Errorf("Commits omitted reachable phase change %s", oid)
	}
}

func Test_FLOW_088_CommitsRejectMissingAndAmbiguousSelectors(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	for _, selector := range []string{"absent", "changes/absent.md"} {
		if _, err := workflow.Commits(root, selector); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
			t.Errorf("Commits(%q) error = %v, want U02", selector, err)
		}
	}
	write(t, root, "changes/0002-manual.md", readFile(t, root, pendingCommitChange))
	commit(t, root, "same message")
	if _, err := workflow.Commits(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Errorf("ambiguous slug error = %v, want U02", err)
	}
	assertCommits(t, root, pendingCommitChange, []string{})
}
package evals_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func abandonGitIdentity(t *testing.T, root string) {
	t.Helper()
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
}

func assertAbandonCommit(t *testing.T, root, parent, oid string) {
	t.Helper()
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) || oid != head(t, root) {
		t.Errorf("Abandon returned %q; HEAD is %s", oid, head(t, root))
	}
	if got := git(t, root, "rev-parse", "HEAD^"); got != parent {
		t.Errorf("Abandon parent = %s, want %s (one new commit)", got, parent)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != "" {
		t.Errorf("Abandon left dirty worktree: %s", got)
	}
	for _, path := range []string{pendingCommitChange, archivedCommitChange} {
		if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
			t.Errorf("change remains at %s: %v", path, err)
		}
	}
	// FLOW-094: the only change-document edit is deletion; no new phase is assigned.
	if diff := git(t, root, "diff", "--unified=0", parent, oid, "--", "changes"); strings.Contains(diff, "+phase:") {
		t.Errorf("Abandon assigned a phase in its commit: %s", diff)
	}
}

func assertAbandonRefusal(t *testing.T, root, selector, code string) {
	t.Helper()
	beforeHead := head(t, root)
	beforeTree := git(t, root, "rev-parse", "HEAD^{tree}")
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeRefs := releaseRefs(t, root)
	if oid, err := workflow.Abandon(root, selector); err == nil || !strings.HasPrefix(err.Error(), code+":") || oid != "" {
		t.Errorf("Abandon(%q) = %q, %v; want %s refusal without OID", selector, oid, err, code)
	}
	if got := head(t, root); got != beforeHead {
		t.Errorf("rejected Abandon moved HEAD: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "rev-parse", "HEAD^{tree}"); got != beforeTree {
		t.Errorf("rejected Abandon changed committed tree: %s -> %s", beforeTree, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("rejected Abandon changed worktree status: %q -> %q", beforeStatus, got)
	}
	if got := releaseRefs(t, root); got != beforeRefs {
		t.Errorf("rejected Abandon changed refs: %q -> %q", beforeRefs, got)
	}
}

func Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	parent := head(t, root)
	beforeSpec := readFile(t, root, "packages/alpha/SPEC.md")
	if commits, err := workflow.Commits(root, "manual"); err != nil || len(commits) != 0 {
		t.Fatalf("queued fixture has phase commits: %v, %v", commits, err)
	}
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon unclaimed queued change: %v", err)
	}
	assertAbandonCommit(t, root, parent, oid)
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-status", "-r", oid); got != "D\t"+pendingCommitChange {
		t.Errorf("queued Abandon effects = %q, want change-file deletion only", got)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != beforeSpec {
		t.Errorf("queued Abandon changed applied spec: %q", got)
	}
}

func Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	initialSpec := readFile(t, root, "packages/alpha/SPEC.md")
	started := start(t, root)
	write(t, root, "packages/alpha/evals/step_test.go", "committed eval effect\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("fixture evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("fixture audit Done = %#v, %v", audited, err)
	}
	phaseCommits, err := workflow.Commits(root, "manual")
	if err != nil || len(phaseCommits) != 2 {
		t.Fatalf("fixture phase commits = %v, %v; want evals and audit", phaseCommits, err)
	}
	if got := readFile(t, root, "packages/alpha/evals/step_test.go"); got != "committed eval effect\n" {
		t.Fatalf("fixture lacks committed phase effect: %q", got)
	}
	before := head(t, root)
	current, err := workflow.Current(root, "manual")
	if err != nil || current.ClaimOID == "" || current.Run != started.Run {
		t.Fatalf("fixture lost owned claim: %#v, %v", current, err)
	}
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon owned stepped change: %v", err)
	}
	assertAbandonCommit(t, root, before, oid)
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-status", "-r", oid); got != "D\t"+pendingCommitChange+"\nD\tpackages/alpha/evals/step_test.go" {
		t.Errorf("stepped Abandon effects = %q; want pending change and committed eval removed", got)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Errorf("stepped Abandon changed applied spec: %q", got)
	}
	if claim := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); claim != "" {
		t.Errorf("owned claim remains: %s", claim)
	}
	if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Errorf("released claim is still current: %v", err)
	}
}

func Test_FLOW_090_AbandonRejectsDirtyWorkingTree(t *testing.T) {
	for _, kind := range []string{"tracked", "untracked"} {
		t.Run(kind, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			if kind == "tracked" {
				write(t, root, "SYSTEM.md", "dirty tracked content\n")
			} else {
				write(t, root, "stray.txt", "dirty untracked content\n")
			}
			assertAbandonRefusal(t, root, "manual", "P01")
			if kind == "tracked" && readFile(t, root, "SYSTEM.md") != "dirty tracked content\n" {
				t.Error("dirty tracked content changed")
			}
			if kind == "untracked" && readFile(t, root, "stray.txt") != "dirty untracked content\n" {
				t.Error("dirty untracked content changed")
			}
		})
	}
}

func Test_FLOW_091_AbandonRejectsForeignClaim(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	started := start(t, root)
	// The linked worktree is clean and sees the same pending change at HEAD.
	if status := git(t, other, "status", "--porcelain=v1", "--untracked-files=all"); status != "" {
		t.Fatalf("foreign worktree is dirty: %s", status)
	}
	ownerHead := head(t, root)
	ownerChange := readFile(t, root, pendingCommitChange)
	assertAbandonRefusal(t, other, "manual", "P02")
	if head(t, root) != ownerHead || readFile(t, root, pendingCommitChange) != ownerChange {
		t.Error("foreign Abandon changed owner's worktree")
	}
	if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
		t.Errorf("foreign Abandon replaced owned claim: %#v, %v", current, err)
	}
}

func Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	initialSpec := readFile(t, root, "packages/alpha/SPEC.md")
	start(t, root)
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("fixture evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("fixture audit Done = %#v, %v", audited, err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	applied, err := done(root)
	if err != nil || applied.State.Phase != "applied" {
		t.Fatalf("fixture impl Done = %#v, %v", applied, err)
	}
	if _, err := os.Stat(filepath.Join(root, archivedCommitChange)); err != nil {
		t.Fatalf("fixture lacks applied archive: %v", err)
	}
	assertAbandonRefusal(t, root, "manual", "P01")
	git(t, root, "revert", "--no-edit", "HEAD")
	if _, err := os.Stat(filepath.Join(root, pendingCommitChange)); err != nil {
		t.Fatalf("git revert did not restore pending change: %v", err)
	}
	if _, err := os.Stat(filepath.Join(root, archivedCommitChange)); !os.IsNotExist(err) {
		t.Fatalf("git revert retained applied archive: %v", err)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Fatalf("git revert failed to restore original spec: %q", got)
	}
	before := head(t, root)
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon after git revert: %v", err)
	}
	assertAbandonCommit(t, root, before, oid)
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Errorf("Abandon after revert changed original spec: %q", got)
	}
}

func Test_FLOW_093_FLOW_095_AbandonRetainsRealAttemptAndFreesAssignedID(t *testing.T) {
	root, started, saved := savedEvalsAttempt(t)
	if len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("fixture did not assign ALP-003: %#v", started.Change.Added)
	}
	beforeAttempt := git(t, root, "rev-parse", saved.AttemptRef)
	if beforeAttempt != saved.AttemptOID {
		t.Fatalf("Block attempt ref %s points to %s, want %s", saved.AttemptRef, beforeAttempt, saved.AttemptOID)
	}
	parent := head(t, root)
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon blocked unclaimed change: %v", err)
	}
	assertAbandonCommit(t, root, parent, oid)
	if got := git(t, root, "rev-parse", saved.AttemptRef); got != beforeAttempt {
		t.Errorf("saved attempt ref changed: %s -> %s", beforeAttempt, got)
	}
	if got := git(t, root, "cat-file", "-t", saved.AttemptOID); got != "commit" {
		t.Errorf("saved attempt object disappeared: %s", got)
	}
	path, err := project.NewChange(root, "alpha", "replacement")
	if err != nil {
		t.Fatalf("NewChange after Abandon: %v", err)
	}
	doc, err := contracts.ParseDocument(readFile(t, root, path))
	if err != nil {
		t.Fatal(err)
	}
	doc.Body = contracts.ReplaceSection(doc.Body, "Motivation", "\nUsers need a replacement alpha behavior.\n")
	doc.Body = contracts.ReplaceSection(doc.Body, "Added", "\n### new:replacement\nReplacement MUST be allocated.\n")
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, path, rendered)
	commit(t, root, "queue replacement change")
	replacement, err := workflow.Start(root, "replacement")
	if err != nil {
		t.Fatalf("Start replacement: %v", err)
	}
	if len(replacement.Change.Added) != 1 || replacement.Change.Added[0].ID != "ALP-003" {
		t.Errorf("replacement assigned IDs = %#v, want reused ALP-003", replacement.Change.Added)
	}
	if got := git(t, root, "rev-parse", saved.AttemptRef); got != beforeAttempt {
		t.Errorf("subsequent Start changed saved attempt ref: %s -> %s", beforeAttempt, got)
	}
}

func Test_FLOW_096_AbandonSelectorsAndMissingAmbiguousRefusal(t *testing.T) {
	for _, selector := range []string{pendingCommitChange, "0001-manual.md", "manual"} {
		t.Run(selector, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			abandonGitIdentity(t, root)
			if commits, err := workflow.Commits(root, selector); err != nil || len(commits) != 0 {
				t.Fatalf("Commits selector %q = %v, %v", selector, commits, err)
			}
			parent := head(t, root)
			oid, err := workflow.Abandon(root, selector)
			if err != nil {
				t.Fatalf("Abandon selector %q: %v", selector, err)
			}
			assertAbandonCommit(t, root, parent, oid)
		})
	}
	root := fixture(t, "", "", "", "", "")
	for _, selector := range []string{"absent", "changes/absent.md"} {
		assertAbandonRefusal(t, root, selector, "U02")
	}
	write(t, root, "changes/0002-manual.md", readFile(t, root, pendingCommitChange))
	commit(t, root, "queue second manual change")
	if _, err := workflow.Commits(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Fatalf("fixture slug is not ambiguous to Commits: %v", err)
	}
	assertAbandonRefusal(t, root, "manual", "U02")
}
package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func status(t *testing.T, root string) workflow.StatusReport {
	t.Helper()
	r, err := workflow.Status(root)
	if err != nil {
		t.Fatalf("Status: %v", err)
	}
	return r
}

func statusEntry(t *testing.T, r workflow.StatusReport, numberedSlug string) workflow.StatusEntry {
	t.Helper()
	for _, entry := range r.Queue {
		if strings.TrimSuffix(filepath.Base(entry.Path), ".md") == numberedSlug {
			return entry
		}
	}
	t.Fatalf("Status queue has no %s: %#v", numberedSlug, r.Queue)
	return workflow.StatusEntry{}
}

func waiting(t *testing.T, r workflow.StatusReport, slug, on, rule string) {
	t.Helper()
	e := statusEntry(t, r, slug)
	if e.WaitingOn != on || e.WaitingRule != rule {
		t.Errorf("%s waiting = %q/%q, want %q/%q", slug, e.WaitingOn, e.WaitingRule, on, rule)
	}
}

func Test_FLOW_101_FLOW_102_FLOW_103_StatusQueueAndForwardWaiting(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "same", "alpha", "", "")
	eligibilityChange(t, root, 3, "direct", "beta", "", "")
	eligibilityChange(t, root, 4, "transitive", "gamma", "", "")
	eligibilityChange(t, root, 5, "free", "delta", "", "")
	eligibilityChange(t, root, 6, "also-blocked", "beta", "", "")
	readyEligibilityFixture(t, root)
	freeFile := readFile(t, root, "changes/0005-free.md")
	write(t, root, "changes/0005-free.md", strings.Replace(freeFile, "module: delta\n", "module: delta\nphase: audit\naudit_rounds: 2\n", 1))
	blockedFile := readFile(t, root, "changes/0006-also-blocked.md")
	write(t, root, "changes/0006-also-blocked.md", strings.Replace(blockedFile, "module: beta\n", "module: beta\nblocked: separately blocked\n", 1))
	r := status(t, root)
	q, err := project.Queue(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Queue) != len(q) {
		t.Fatalf("queue length = %d, want %d", len(r.Queue), len(q))
	}
	for i, ch := range q {
		e := r.Queue[i]
		if e.Slug != ch.Slug || e.Path != ch.Path || e.Module != ch.Module {
			t.Errorf("queue[%d] identity = %#v, want %#v", i, e, ch)
		}
	}
	blocked := statusEntry(t, r, "0001-manual")
	if blocked.Phase != "impl" || blocked.Blocked != "fixture blocked" || blocked.AuditRounds != 0 {
		t.Errorf("blocked state = %#v", blocked)
	}
	free := statusEntry(t, r, "0005-free")
	if free.Phase != "audit" || free.Blocked != "" || free.AuditRounds != 2 {
		t.Errorf("authored state = %#v", free)
	}
	if same := statusEntry(t, r, "0002-same"); same.Phase != "queued" || same.Blocked != "" || same.AuditRounds != 0 {
		t.Errorf("default state = %#v", same)
	}
	waiting(t, r, "0001-manual", "", "")
	waiting(t, r, "0002-same", "0001-manual", "same_module")
	waiting(t, r, "0003-direct", "0001-manual", "depends_on_blocked")
	waiting(t, r, "0004-transitive", "0001-manual", "depends_on_blocked")
	waiting(t, r, "0005-free", "", "")
	waiting(t, r, "0006-also-blocked", "", "")
}

func Test_FLOW_104_FLOW_105_StatusInterfaceAndProjectedWaiting(t *testing.T) {
	root := eligibilityBase(t, "gamma")
	eligibilityChange(t, root, 2, "interface-direct", "beta", "", "Beta API changes.\n")
	eligibilityChange(t, root, 3, "internal-direct", "beta", "", "")
	eligibilityChange(t, root, 4, "interface-transitive", "alpha", "", "Alpha API changes.\n")
	eligibilityChange(t, root, 5, "internal-transitive", "alpha", "", "")
	readyEligibilityFixture(t, root)
	r := status(t, root)
	waiting(t, r, "0002-interface-direct", "0001-manual", "interface_dependency")
	waiting(t, r, "0003-internal-direct", "", "")
	waiting(t, r, "0004-interface-transitive", "0001-manual", "interface_dependency")
	waiting(t, r, "0005-internal-transitive", "", "")

	projected := eligibilityBase(t, "alpha")
	eligibilityChange(t, projected, 2, "own-add", "delta", "set:\n  depends_on: [alpha]\n", "")
	eligibilityChange(t, projected, 3, "detach-beta", "beta", "set:\n  depends_on: []\n", "")
	eligibilityChange(t, projected, 4, "after-detach", "gamma", "", "")
	eligibilityChange(t, projected, 5, "own-remove", "delta", "set:\n  depends_on: []\n", "")
	readyEligibilityFixture(t, projected)
	r = status(t, projected)
	waiting(t, r, "0002-own-add", "0001-manual", "depends_on_blocked")
	waiting(t, r, "0003-detach-beta", "", "")
	waiting(t, r, "0004-after-detach", "", "")
	waiting(t, r, "0005-own-remove", "", "")
}

func Test_FLOW_106_FLOW_107_FLOW_108_StatusModuleMetadata(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\nboundaries_opt_out: approved fixture exception\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [quarantine: intermittent fixture]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n\n### DEL-001 [quarantine: unavailable service]\nDelta MUST work.\n")
	commit(t, root, "module metadata")
	snapshot, err := project.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	r := status(t, root)
	if !reflect.DeepEqual(r.Modules, snapshot.Modules) {
		t.Errorf("modules = %#v, want discovered %#v", r.Modules, snapshot.Modules)
	}
	byName := map[string]project.Module{}
	for _, m := range r.Modules {
		byName[m.Name] = m
	}
	if len(byName) != 4 || byName["alpha"].IDPrefix != "ALP" || byName["alpha"].Dir != "packages/alpha" || !reflect.DeepEqual(byName["gamma"].DependsOn, []string{"beta"}) || !reflect.DeepEqual(byName["beta"].Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(byName["alpha"].Zones.Disposable, []string{"internal/**"}) {
		t.Errorf("discovered identity/dependencies/zones = %#v", byName)
	}
	wantQ := []workflow.Quarantine{{Module: "alpha", ID: "ALP-002", Reason: "intermittent fixture"}, {Module: "delta", ID: "DEL-001", Reason: "unavailable service"}}
	if !reflect.DeepEqual(r.Quarantines, wantQ) {
		t.Errorf("quarantines = %#v, want %#v", r.Quarantines, wantQ)
	}
	if !reflect.DeepEqual(r.BoundaryOptOuts, []workflow.BoundaryOptOut{{Module: "alpha", Reason: "approved fixture exception"}}) {
		t.Errorf("opt-outs = %#v", r.BoundaryOptOuts)
	}
}

func Test_FLOW_109_FLOW_111_StatusProvenanceDriftAndRecentModuleWarning(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n")
	archivedImpl(t, root, "0002-alpha-old", "alpha", "2026-06-01T12:00:00Z")
	archivedImpl(t, root, "0003-beta-old", "beta", "2026-06-02T12:00:00Z")
	archivedImpl(t, root, "0003-gamma-old", "gamma", "2026-06-03T12:00:00Z")
	archivedImpl(t, root, "0004-alpha-new", "alpha", "2026-06-11T12:00:00Z")
	archivedImpl(t, root, "0005-alpha-newer", "alpha", "2026-06-12T12:00:00Z")
	archivedImpl(t, root, "0006-alpha-newest", "alpha", "2026-06-13T12:00:00Z")
	write(t, root, ".fireweed/drills/alpha.md", fmt.Sprintf("---\ndate: 2026-06-10T12:00:00Z\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: fail\n---\n", head(t, root)))
	write(t, root, ".fireweed/drills/beta.md", fmt.Sprintf("---\ndate: 2026-06-14T12:00:00Z\nbase_commit: %s\nagent: fixture\nmodel: fixture\nresult: fail\n---\n", head(t, root)))
	commit(t, root, "authored archives and failed drill")
	r := status(t, root)
	for name, want := range map[string]int{"alpha": 3, "beta": 0, "gamma": 1, "delta": 0} {
		if got, ok := r.DriftSinceDrill[name]; !ok || got != want {
			t.Errorf("drift %s = %d, present %t; want %d", name, got, ok, want)
		}
	}
	d, ok := warningFor(r.Diagnostics, "P08")
	if !ok || !d.Warning || !strings.Contains(d.Message, "alpha") || !strings.Contains(d.Message, "3") {
		t.Errorf("P08 warning = %#v", r.Diagnostics)
	}
}

func Test_FLOW_110_StatusSharedBaseUsesSameArchivedDocumentIdentity(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	pendingBase := head(t, root)
	git(t, root, "checkout", "-b", "feature")
	git(t, root, "checkout", "main")
	if err := os.MkdirAll(filepath.Join(root, "changes/archive"), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", "changes/0001-manual.md", "changes/archive/0001-manual.md")
	mainArchive := readFile(t, root, "changes/archive/0001-manual.md")
	write(t, root, "changes/archive/0001-manual.md", strings.Replace(mainArchive, "module: alpha\n", "module: alpha\nphase: applied\n", 1))
	commit(t, root, "base has applied manual document")
	approved := head(t, root)
	git(t, root, "checkout", "feature")
	if err := os.MkdirAll(filepath.Join(root, "changes/archive"), 0755); err != nil {
		t.Fatal(err)
	}
	git(t, root, "mv", "changes/0001-manual.md", "changes/archive/0001-manual.md")
	change := readFile(t, root, "changes/archive/0001-manual.md")
	write(t, root, "changes/archive/0001-manual.md", strings.Replace(change, "module: alpha\n", "module: alpha\nphase: applied\n", 1))
	archivedImpl(t, root, "0002-extra", "beta", "2026-06-10T12:00:00Z")
	commit(t, root, "one approved and one new archive")
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("local base unapproved = %d, want 1", got)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", head(t, root))
	if got := status(t, root).Unapproved; got != 0 {
		t.Errorf("remote base unapproved = %d, want 0", got)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", approved)
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("shared base unapproved = %d, want 1", got)
	}
	git(t, root, "update-ref", "refs/remotes/origin/main", pendingBase)
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("pending base identity unapproved = %d, want 1", got)
	}
	config := readFile(t, root, "fireweed.yaml")
	write(t, root, "fireweed.yaml", strings.Replace(config, "base: main\n", "base: stable\nremote: upstream\n", 1))
	git(t, root, "update-ref", "refs/heads/stable", head(t, root))
	git(t, root, "update-ref", "refs/remotes/upstream/stable", approved)
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("configured remote/base unapproved = %d, want 1", got)
	}
	git(t, root, "update-ref", "-d", "refs/remotes/upstream/stable")
	if got := status(t, root).Unapproved; got != 0 {
		t.Errorf("configured local fallback unapproved = %d, want 0", got)
	}
	write(t, root, "fireweed.yaml", strings.Replace(config, "base: main\n", "", 1))
	git(t, root, "update-ref", "refs/remotes/origin/stable", approved)
	git(t, root, "symbolic-ref", "refs/remotes/origin/HEAD", "refs/remotes/origin/stable")
	if got := status(t, root).Unapproved; got != 1 {
		t.Errorf("default base via origin/HEAD unapproved = %d, want 1", got)
	}
}

func Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	if _, err := done(root); err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	write(t, root, "author-note.txt", "no phase transition\n")
	commit(t, root, "author note after application")
	commits, err := workflow.Commits(root, "0001-manual.md")
	if err != nil {
		t.Fatal(err)
	}
	if len(commits) != 2 {
		t.Fatalf("real phase commits = %d, want 2", len(commits))
	}
	r := status(t, root)
	if got := r.PhaseCommits["0001-manual"]; got != len(commits) {
		t.Errorf("phase commits = %#v, want %d for 0001-manual", r.PhaseCommits, len(commits))
	}
}

func Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  eval_ratio: 1\n", "", "", "", "")
	write(t, root, "packages/alpha/evals/weight.txt", "one\ntwo\n")
	commit(t, root, "cost fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(root, false); err != nil {
		t.Fatal(err)
	}
	snapshots, err := workflow.Costs(root)
	if err != nil {
		t.Fatal(err)
	}
	latest := map[string]workflow.ModuleCost{}
	var latestAlphaB03, olderAlphaB03 project.Diagnostic
	var hasLatestAlphaB03, hasOlderAlphaB03 bool
	for _, s := range snapshots {
		for _, c := range s.Modules {
			latest[c.Module] = c
			if c.Module == "alpha" {
				olderAlphaB03, hasOlderAlphaB03 = latestAlphaB03, hasLatestAlphaB03
				latestAlphaB03, hasLatestAlphaB03 = warningFor(s.Diagnostics, "B03")
			}
		}
	}
	if !hasOlderAlphaB03 || !hasLatestAlphaB03 || olderAlphaB03 == latestAlphaB03 {
		t.Fatalf("cost snapshots need distinct older and latest alpha B03 warnings: %#v", snapshots)
	}
	config := readFile(t, root, "fireweed.yaml")
	write(t, root, "fireweed.yaml", strings.Replace(config, "eval_ratio: 1", "eval_ratio: 100", 1))
	r := status(t, root)
	if len(r.Costs) != 2 {
		t.Errorf("cost modules = %#v, want only measured alpha and beta", r.Costs)
	}
	for _, c := range r.Costs {
		if !reflect.DeepEqual(c, latest[c.Module]) {
			t.Errorf("latest %s cost = %#v, want %#v", c.Module, c, latest[c.Module])
		}
	}
	for _, module := range []string{"alpha", "beta"} {
		want := latest[module]
		got, ok := costFor(r.Costs, module)
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("recorded %s cost = %#v (present %t), want %#v", module, got, ok, want)
		}
	}
	if _, ok := costFor(r.Costs, "gamma"); ok {
		t.Errorf("unmeasured gamma has cost: %#v", r.Costs)
	}
	if d, ok := warningFor(r.Diagnostics, "B03"); !ok || !d.Warning || d != latestAlphaB03 {
		t.Errorf("selected B03 warning = %#v, want %#v", r.Diagnostics, latestAlphaB03)
	}
	if containsDiagnostic(r.Diagnostics, olderAlphaB03) {
		t.Errorf("older B03 warning returned: %#v", r.Diagnostics)
	}
	count := 0
	for _, d := range r.Diagnostics {
		if d.Code == "B03" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("selected B03 warning count = %d, want 1: %#v", count, r.Diagnostics)
	}
}

func containsDiagnostic(ds []project.Diagnostic, want project.Diagnostic) bool {
	for _, d := range ds {
		if d == want {
			return true
		}
	}
	return false
}

exec
/bin/zsh -lc "cat packages/workflow/evals/pr_body_test.go; sed -n '130,315p' packages/workflow/evals/costs_test.go; cat packages/workflow/evals/release_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 470ms:
package evals_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func prBody(t *testing.T, root, revisionRange string) workflow.PRBodyReport {
	t.Helper()
	r, err := workflow.PRBody(root, revisionRange)
	if err != nil {
		t.Fatalf("PRBody(%q): %v", revisionRange, err)
	}
	return r
}

func prRequire(t *testing.T, body string, terms ...string) {
	t.Helper()
	for _, term := range terms {
		if !strings.Contains(body, term) {
			t.Errorf("PR body lacks %q:\n%s", term, body)
		}
	}
}

func prReject(t *testing.T, body string, terms ...string) {
	t.Helper()
	for _, term := range terms {
		if strings.Contains(body, term) {
			t.Errorf("PR body unexpectedly contains %q:\n%s", term, body)
		}
	}
}

func prPhase(t *testing.T, root, phase string) {
	t.Helper()
	path := "changes/0001-manual.md"
	before := readFile(t, root, path)
	after := before
	for _, old := range []string{"evals", "audit", "impl"} {
		if strings.Contains(after, "phase: "+old+"\n") {
			after = strings.Replace(after, "phase: "+old+"\n", "phase: "+phase+"\n", 1)
			break
		}
	}
	if after == before {
		t.Fatalf("fixture phase missing in %s", path)
	}
	write(t, root, path, after)
}

// The earlier phase commits stay reachable when the requested range contains
// only the final implementation commit. The final commit also archives the
// selected identity, as a real Done does.
func prHistoryFixture(t *testing.T) (root, base, finalPhase, endpoint string) {
	t.Helper()
	root = fixture(t, "worker_writable: [workspace.lock]\n", "", "### ALP-003 [review]\nAdded review clause MUST be inspected.\n", "", "")
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\npreserve: [internal/keep.go]\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [review]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	manual := readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "### ALP-001\n", "### ALP-001 [review]\n", 1)
	manual = strings.Replace(manual, "module: alpha\n", "module: alpha\nphase: evals\nrun: pr-fixture\naudit_rounds: 0\n", 1)
	write(t, root, "changes/0001-manual.md", manual)
	write(t, root, "packages/alpha/internal/keep.go", "package alpha // preserved baseline\n")
	write(t, root, "packages/alpha/internal/existing.go", "package alpha // existing disposable\n")
	write(t, root, "changes/0003-outside.md", "---\nmodule: gamma\n---\n\n## Motivation\nOutside range change.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n")
	write(t, root, "changes/0004-removed.md", "---\nmodule: beta\n---\n\n## Motivation\nRemoved identity.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n")
	commit(t, root, "fixture phase base")
	base = head(t, root)

	// Evals: an interface edit, one new disposable stub, an eval file, a
	// module manifest, an ordinary module file and a shared writable lockfile.
	write(t, root, "packages/alpha/api.go", "package alpha // RANGE INTERFACE EDIT\n")
	write(t, root, "packages/alpha/evals/report_test.go", "package alpha_test // RANGE EVAL EDIT\n")
	write(t, root, "packages/alpha/internal/new_stub.go", "package alpha // RANGE STUB ADDITION\n")
	write(t, root, "packages/alpha/go.mod", "module fixture.invalid/alpha\n")
	write(t, root, "packages/alpha/other.txt", "RANGE OTHER EDIT\n")
	write(t, root, "workspace.lock", "RANGE SHARED EDIT\n")
	write(t, root, "changes/0002-pending.md", "---\nmodule: beta\neval_fix: [BET-001]\nphase: audit\nrun: beta-pr-fixture\naudit_rounds: 0\n---\n\n## Motivation\nSelected pending identity.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Provenance\n- evals 2026-10-30T12:00:00Z tests=7 failing=2 passing=5\n")
	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: evals\n", "phase: audit\n", 1)
	manual += "\n## Provenance\n- evals 2026-10-30T12:00:00Z tests=11 failing=3 passing=8\n"
	write(t, root, "changes/0001-manual.md", manual)
	commit(t, root, "evals phase with selected pending change")

	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: audit\n", "phase: evals\n", 1)
	manual = strings.Replace(manual, "audit_rounds: 0\n", "audit_rounds: 1\n", 1)
	manual += "\n## Audit\n- ALP-001: \"" + newClause + "\" — PRIOR AUDIT FINDING missing assertion\n"
	write(t, root, "changes/0001-manual.md", manual)
	commit(t, root, "audit finding round one")

	write(t, root, "packages/alpha/evals/report_test.go", "package alpha_test // CORRECTED EVAL EDIT\n")
	prPhase(t, root, "audit")
	commit(t, root, "corrected eval phase")

	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: audit\n", "phase: impl\n", 1)
	manual = strings.Replace(manual, "\n## Audit\n- ALP-001: \""+newClause+"\" — PRIOR AUDIT FINDING missing assertion\n", "\n## Audit\nNo findings\nChecked IDs: ALP-001\n", 1)
	write(t, root, "changes/0001-manual.md", manual)
	commit(t, root, "clean audit phase")
	finalPhase = head(t, root)

	write(t, root, "packages/alpha/internal/keep.go", "package alpha // PRESERVED IMPLEMENTATION EDIT\n")
	write(t, root, "packages/alpha/internal/existing.go", "package alpha // IMPLEMENTATION EDIT\n")
	manual = readFile(t, root, "changes/0001-manual.md")
	manual = strings.Replace(manual, "phase: impl\n", "phase: applied\n", 1)
	manual += "- impl 2026-10-30T12:30:00Z tests=13 failing=0 passing=13\n"
	write(t, root, "changes/archive/0001-manual.md", manual)
	if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(root, "changes/0004-removed.md")); err != nil {
		t.Fatal(err)
	}
	commit(t, root, "implementation and archive")
	endpoint = head(t, root)
	return
}

func Test_FLOW_114_FLOW_115_FLOW_116_FLOW_117_FLOW_118_FLOW_119_FLOW_120_FLOW_121_PRBodySelectedHistoryAndEndpoint(t *testing.T) {
	root, base, finalPhase, endpoint := prHistoryFixture(t)
	// Later committed and dirty edits must not change the requested endpoint.
	write(t, root, "fireweed.yaml", strings.Replace(readFile(t, root, "fireweed.yaml"), "modules: [\"packages/*\"]", "modules: [] # LATER HEAD CONFIG MARKER", 1))
	write(t, root, "packages/alpha/SPEC.md", readFile(t, root, "packages/alpha/SPEC.md")+"\nLATER HEAD SPEC MARKER\n")
	commit(t, root, "later unrelated HEAD")
	write(t, root, "fireweed.yaml", readFile(t, root, "fireweed.yaml")+"\n# DIRTY WORKTREE CONFIG MARKER\n")
	write(t, root, "packages/alpha/SPEC.md", readFile(t, root, "packages/alpha/SPEC.md")+"\nDIRTY WORKTREE SPEC MARKER\n")
	full := base + ".." + endpoint
	partial := finalPhase + ".." + endpoint
	t.Run("FLOW_114_explicit_endpoints", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, base[:12], endpoint[:12])
		prReject(t, r.Body, "LATER HEAD CONFIG MARKER", "LATER HEAD SPEC MARKER", "DIRTY WORKTREE CONFIG MARKER", "DIRTY WORKTREE SPEC MARKER")
	})
	t.Run("FLOW_115_selected_identities", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "changes/archive/0001-manual.md", "alpha", "changes/0002-pending.md", "beta", "eval_fix")
		prReject(t, r.Body, "0003-outside", "0004-removed")
		p := prBody(t, root, partial)
		prRequire(t, p.Body, "changes/archive/0001-manual.md")
		prReject(t, p.Body, "0002-pending.md")
	})
	t.Run("FLOW_116_provenance_and_evals", func(t *testing.T) {
		fullReport := prBody(t, root, full)
		prRequire(t, fullReport.Body, "tests=7", "failing=2", "passing=5")
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "tests=11", "failing=3", "passing=8", "report_test.go")
		}
	})
	t.Run("FLOW_117_phase_file_categories", func(t *testing.T) {
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "api.go", "new_stub.go", "go.mod", "other.txt", "workspace.lock")
		}
	})
	t.Run("FLOW_118_preserved_edits", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "keep.go")
	})
	t.Run("FLOW_119_review_specs", func(t *testing.T) {
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "ALP-003", "Added review clause MUST be inspected.", "ALP-001", newClause)
			prReject(t, r.Body, "Alpha MUST retain the second behavior.")
		}
	})
	t.Run("FLOW_120_prior_corrected_audit", func(t *testing.T) {
		for _, rr := range []string{full, partial} {
			r := prBody(t, root, rr)
			prRequire(t, r.Body, "PRIOR AUDIT FINDING")
			if !strings.Contains(r.Body, "audit_rounds") && !strings.Contains(strings.ToLower(r.Body), "audit rounds") {
				t.Errorf("PR body omits audit round count: %s", r.Body)
			}
		}
	})
	t.Run("FLOW_121_decisions", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "Use the revised policy.", "Retain old policy.")
	})
}

func prSaveCost(t *testing.T, root, run, phase string, snapshot workflow.CostSnapshot) {
	t.Helper()
	b, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, filepath.Join(".fireweed/runs", run, phase, "costs.json"), string(b))
}

func Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\nboundaries_opt_out: ENDPOINT OPT OUT REASON\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [quarantine: ENDPOINT QUARANTINE REASON]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	manual := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(manual, "module: alpha\n", "module: alpha\nblocked: ENDPOINT BLOCK REASON\n", 1))
	commit(t, root, "endpoint exceptions")
	endpoint := head(t, root)
	old := workflow.CostSnapshot{Run: "old", Phase: "evals", RecordedAt: time.Date(2026, 10, 29, 0, 0, 0, 0, time.UTC), Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 99, CasesPerSpec: map[string]int{"ALP-001": 99}, EvalLines: 99, DisposableLines: 1}}, Diagnostics: []project.Diagnostic{{Code: "B01", Message: "OLD B01", Warning: true}}}
	latest := workflow.CostSnapshot{Run: "new", Phase: "impl", RecordedAt: time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC), Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 4.25, CasesPerSpec: map[string]int{"ALP-001": 5}, EvalLines: 17, DisposableLines: 3}, {Module: "beta", Seconds: 2.5, CasesPerSpec: map[string]int{"BET-001": 2}, EvalLines: 8, DisposableLines: 4}}, Diagnostics: []project.Diagnostic{{Code: "B01", Message: "SAVED B01 time", Warning: true}, {Code: "B02", Message: "ALP-001: SAVED B02 cases", Warning: true}, {Code: "B03", Message: "alpha: SAVED B03 ratio", Warning: true}}}
	prSaveCost(t, root, "old", "evals", old)
	prSaveCost(t, root, "new", "impl", latest)
	write(t, root, "fireweed.yaml", readFile(t, root, "fireweed.yaml")+"\nbudgets:\n  cases_per_spec: 1000\n  eval_ratio: 1000\n")
	write(t, root, "packages/alpha/evals/later.txt", strings.Repeat("later\n", 50))
	t.Run("FLOW_122_endpoint_exceptions", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "alpha", "ALP-002", "ENDPOINT QUARANTINE REASON", "ENDPOINT OPT OUT REASON", "0001-manual", "ENDPOINT BLOCK REASON")
	})
	t.Run("FLOW_123_latest_saved_costs", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "SAVED B01", "SAVED B02", "SAVED B03")
		prReject(t, r.Body, "OLD B01")
		section := regexp.MustCompile(`(?ms)^#{1,6} (?:Cost|Costs|Metric|Metrics)\s*\n(.*?)(?:^#{1,6} |\z)`).FindStringSubmatch(r.Body)
		if len(section) != 2 {
			t.Errorf("PR body lacks a Cost/Metric section: %s", r.Body)
		} else {
			rows := regexp.MustCompile(`(?m)^\|\s*(?:alpha|beta|gamma)\s*\|[^\n]*$`).FindAllString(section[1], -1)
			want := map[string]string{
				"alpha": "alpha|4.25|ALP-001: 5|17|3",
				"beta":  "beta|2.5|BET-001: 2|8|4",
			}
			seen := map[string]bool{}
			for _, row := range rows {
				cells := strings.Split(strings.Trim(strings.TrimSpace(row), "|"), "|")
				for i := range cells {
					cells[i] = strings.TrimSpace(cells[i])
				}
				module := cells[0]
				if seen[module] || strings.Join(cells, "|") != want[module] {
					t.Errorf("unexpected or stale cost measurement row %q", row)
				}
				seen[module] = true
			}
			for module := range want {
				if !seen[module] {
					t.Errorf("missing saved %s cost measurement row: %s", module, section[1])
				}
			}
		}
		for _, code := range []string{"B01", "B02", "B03"} {
			if !containsCode(r.Diagnostics, code) {
				t.Errorf("PRBody diagnostics lack saved %s: %#v", code, r.Diagnostics)
			}
		}
	})
}

func Test_FLOW_126_PRBodyEndpointConfigError(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	write(t, root, "fireweed.yaml", "[invalid endpoint configuration\n")
	commit(t, root, "invalid endpoint configuration")

	_, err := workflow.PRBody(root, base+".."+head(t, root))
	if err == nil {
		t.Fatal("PRBody accepted malformed endpoint configuration")
	}
	if !workflow.IsEndpointConfigError(err) {
		t.Errorf("IsEndpointConfigError(%v) = false, want true", err)
	}
}

func Test_FLOW_127_PRBodyUnresolvableRangeIsNotEndpointConfigError(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	_, err := workflow.PRBody(root, head(t, root)+"..not-a-commit")
	if err == nil {
		t.Fatal("PRBody accepted an unresolvable range")
	}
	if workflow.IsEndpointConfigError(err) {
		t.Errorf("IsEndpointConfigError(%v) = true, want false", err)
	}
}

func Test_FLOW_115_FLOW_124_FLOW_125_PRBodyDisposableArtifactBoundaryAndRanges(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	for path, data := range map[string][]byte{
		"packages/alpha/internal/binary.dat": {0, 1, 2, 3},
		"packages/alpha/internal/large.txt":  []byte(strings.Repeat("L", 262145)),
		"packages/alpha/internal/exact.txt":  []byte(strings.Repeat("E", 262144)),
		"packages/alpha/internal/small.txt":  []byte("ordinary text\n"),
	} {
		full := filepath.Join(root, path)
		if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, data, 0644); err != nil {
			t.Fatal(err)
		}
	}
	commit(t, root, "new disposable artifacts")
	endpoint := head(t, root)
	t.Run("FLOW_124_binary_and_large_artifacts", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "binary.dat", "large.txt", "262145", "B04")
		if !regexp.MustCompile("(?m)(binary\\.dat[^\\n]*\\b4\\b|\\b4\\b[^\\n]*binary\\.dat)").MatchString(r.Body) {
			t.Errorf("binary artifact lacks its 4-byte size: %s", r.Body)
		}
		count := 0
		for _, d := range r.Diagnostics {
			if d.Code == "B04" {
				count++
				if !d.Warning || !(strings.Contains(d.Message, "binary.dat") || strings.Contains(d.Message, "large.txt")) || strings.Contains(d.Message, "exact.txt") || strings.Contains(d.Message, "small.txt") {
					t.Errorf("unexpected B04 diagnostic: %#v", d)
				}
			}
		}
		if count != 2 {
			t.Errorf("B04 count = %d, want 2: %#v", count, r.Diagnostics)
		}
	})
	t.Run("FLOW_115_FLOW_125_empty_and_invalid_ranges", func(t *testing.T) {
		empty := prBody(t, root, endpoint+".."+endpoint)
		// FLOW-115: the empty statement must identify the change selection;
		// an unrelated "None" in another report section is insufficient.
		if !regexp.MustCompile(`(?i)\b(?:no (?:selected )?changes?|none selected|(?:selected )?changes\s*:?\s+none)\b`).MatchString(empty.Body) {
			t.Errorf("empty range did not explicitly report no selected changes: %q", empty.Body)
		}
		if containsCode(empty.Diagnostics, "B04") {
			t.Errorf("empty range yielded B04: %#v", empty.Diagnostics)
		}
		for _, invalid := range []string{"", "not-a-commit.." + endpoint, base + "..not-a-commit", "not-a-range"} {
			_, err := workflow.PRBody(root, invalid)
			if err == nil || !strings.Contains(err.Error(), "P01") || (invalid != "" && !strings.Contains(err.Error(), "not-a-commit") && !strings.Contains(err.Error(), "not-a-range") && !strings.Contains(err.Error(), "resolve")) {
				t.Errorf("PRBody(%q) error = %v, want P01 with rejected range or cause", invalid, err)
			}
		}
	})
}
		t.Errorf("FLOW-097 full impl Done omitted evaluated dependent: %#v", implDone.Costs)
	}
}

func Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "set:\n  evals: [evals/**, internal/overlap.txt]\n", "", "", "")
	path := filepath.Join(root, "packages/alpha/SPEC.md")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/SPEC.md", strings.Replace(string(b), "depends_on: []\n", "depends_on: []\nprotected: [internal/protected.txt]\npreserve: [internal/preserve.txt]\n", 1))
	write(t, root, "packages/alpha/evals/one.txt", "a\nb")
	write(t, root, "packages/alpha/internal/overlap.txt", "x\n")
	write(t, root, "packages/alpha/internal/code.txt", "p\nq")
	write(t, root, "packages/alpha/internal/preserve.txt", "r\n")
	write(t, root, "packages/alpha/internal/protected.txt", "s\n")
	if err := os.Symlink("code.txt", filepath.Join(root, "packages/alpha/internal/link.txt")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/beta/evals/one.txt", "z")
	write(t, root, "packages/beta/internal/code.txt", "i\n")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  printf '<testsuite><testcase name="ALP-001 ALP-002" classname="ALP-001" time="1.25">%s</testcase><testcase name="ALP-001 second" time="0.5"/></testsuite>' "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="2"/></testsuite>' > "$2"
`)
	commit(t, root, "held zone and multi-ID cost fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	result, err := verify(root, false)
	if err != nil {
		t.Fatalf("full impl Verify: %v", err)
	}
	if cases := result.Reports["alpha"].Cases; len(cases) != 2 || !reflect.DeepEqual(cases[0].IDs, []string{"ALP-001", "ALP-002"}) || !reflect.DeepEqual(cases[1].IDs, []string{"ALP-001"}) || result.Reports["alpha"].Seconds != 1.75 || result.Reports["beta"].Seconds != 2 {
		t.Fatalf("cost fixture reports have unexpected cases or time: %#v", result.Reports)
	}
	if len(result.Costs) != 2 {
		t.Errorf("FLOW-098 costs for evaluated modules = %#v; want alpha and beta", result.Costs)
	}
	alpha, ok := costFor(result.Costs, "alpha")
	if !ok || alpha.Seconds != 1.75 || !reflect.DeepEqual(alpha.CasesPerSpec, map[string]int{"ALP-001": 2, "ALP-002": 1}) || alpha.EvalLines != 3 || alpha.DisposableLines != 3 {
		t.Errorf("FLOW-098 alpha cost = %#v, present %v; want seconds 1.75, counts 2/1, eval 3, disposable 3", alpha, ok)
	}
	beta, ok := costFor(result.Costs, "beta")
	if !ok || beta.Seconds != 2 || !reflect.DeepEqual(beta.CasesPerSpec, map[string]int{"BET-001": 1}) || beta.EvalLines != 1 || beta.DisposableLines != 1 {
		t.Errorf("FLOW-098 beta cost = %#v, present %v", beta, ok)
	}
}

func Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  verify_time: 3s\n  cases_per_spec: 1\n  eval_ratio: 1\n", "", "", "", "")
	write(t, root, "packages/alpha/evals/weight.txt", "a\nb\n")
	write(t, root, ".fireweed/beta-fail", "yes\n")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  printf '<testsuite><testcase name="ALP-001 first" time="1">%s</testcase><testcase name="ALP-001 second" time="1"/><testcase name="ALP-002" time="1"/></testsuite>' "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="1"><failure/></testcase></testsuite>' > "$2"
exit 1
`)
	commit(t, root, "custom budget cost fixture")
	start(t, root)
	verified, err := verify(root, true)
	if err != nil {
		t.Fatalf("warning made evals Verify fail: %v", err)
	}
	if d, ok := warningFor(verified.Diagnostics, "B01"); ok {
		t.Errorf("FLOW-099 B01 at custom 3s threshold = %#v", d)
	}
	requireCostWarning(t, verified.Diagnostics, "B02", "ALP-001", "2", "1")
	requireCostWarning(t, verified.Diagnostics, "B03", "alpha", "1")
	if cost, ok := costFor(verified.Costs, "alpha"); !ok || cost.EvalLines != 2 || cost.DisposableLines != 0 {
		t.Errorf("FLOW-099 zero-disposable measurement = %#v", verified.Costs)
	}
	evalsDone, err := done(root)
	if err != nil {
		t.Fatalf("warning made evals Done fail: %v", err)
	}
	if d, ok := warningFor(evalsDone.Diagnostics, "B01"); ok {
		t.Errorf("FLOW-099 evals Done B01 at custom 3s threshold = %#v", d)
	}
	for _, code := range []string{"B02", "B03"} {
		requireCostWarning(t, evalsDone.Diagnostics, code)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	before := head(t, root)
	failed, err := done(root)
	if err == nil {
		t.Error("fixture expected dependent testcase failure from impl Done")
	}
	if head(t, root) != before {
		t.Error("failed impl Done committed")
	}
	if len(failed.Reports["beta"].Cases) != 1 {
		t.Errorf("failure fixture did not evaluate beta: %#v", failed.Reports)
	}
	for _, code := range []string{"B01", "B02", "B03"} {
		requireCostWarning(t, failed.Diagnostics, code)
	}
	requireCostWarning(t, failed.Diagnostics, "B01", "4", "3")
	if snapshot, ok := costAt(t, root, failed.State.Run, "impl"); ok {
		if !reflect.DeepEqual(snapshot.Modules, failed.Costs) {
			t.Errorf("FLOW-099 failed Done snapshot costs = %#v; result %#v", snapshot.Modules, failed.Costs)
		}
		for _, code := range []string{"B01", "B02", "B03"} {
			retained, found := warningFor(snapshot.Diagnostics, code)
			returned, present := warningFor(failed.Diagnostics, code)
			if !found || !present || retained != returned {
				t.Errorf("FLOW-099 %s failed Done retained warning = %#v; result %#v", code, retained, returned)
			}
		}
	}
}

func Test_FLOW_099_DefaultBudgetBoundaries(t *testing.T) {
	for _, tc := range []struct {
		name  string
		cases int
		secs  int
		lines int
		warn  bool
	}{{"equal", 15, 900, 4, false}, {"above", 16, 901, 5, true}} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			write(t, root, "packages/alpha/evals/weight.txt", strings.Repeat("line\n", tc.lines))
			write(t, root, "packages/alpha/internal/code.txt", "line\n")
			var cases strings.Builder
			for i := 0; i < tc.cases; i++ {
				seconds, failure := 0, ""
				if i == 0 {
					seconds, failure = tc.secs, "<failure/>"
				}
				fmt.Fprintf(&cases, `<testcase name="ALP-001 case %d" time="%d">%s</testcase>`, i, seconds, failure)
			}
			cases.WriteString(`<testcase name="ALP-002" time="0"/>`)
			write(t, root, ".fireweed/eval.sh", "#!/bin/sh\ncat > \"$2\" <<'XML'\n<testsuite>"+cases.String()+"</testsuite>\nXML\nexit 1\n")
			commit(t, root, "default budget boundary fixture")
			start(t, root)
			result, err := verify(root, true)
			if err != nil {
				t.Fatalf("boundary Verify: %v", err)
			}
			if report := result.Reports["alpha"]; len(report.Cases) != tc.cases+1 || report.Seconds != float64(tc.secs) {
				t.Fatalf("budget fixture JUnit = %#v; want %d cases and %ds", report, tc.cases+1, tc.secs)
			}
			for _, code := range []string{"B01", "B02", "B03"} {
				d, ok := warningFor(result.Diagnostics, code)
				if ok != tc.warn || ok && !d.Warning {
					t.Errorf("FLOW-099 %s at defaults (secs %d, cases %d, ratio %d) = %#v; want warning %v", code, tc.secs, tc.cases, tc.lines, result.Diagnostics, tc.warn)
				}
			}
			if tc.warn {
				requireCostWarning(t, result.Diagnostics, "B01")
				if d, _ := warningFor(result.Diagnostics, "B01"); !strings.Contains(d.Message, "901") && !strings.Contains(d.Message, "15m1s") {
					t.Errorf("FLOW-099 B01 warning omits measured 901s: %q", d.Message)
				}
				if d, _ := warningFor(result.Diagnostics, "B01"); !strings.Contains(d.Message, "900") && !strings.Contains(d.Message, "15m") {
					t.Errorf("FLOW-099 B01 warning omits default 15m limit: %q", d.Message)
				}
				requireCostWarning(t, result.Diagnostics, "B02", "ALP-001", "16", "15")
				requireCostWarning(t, result.Diagnostics, "B03", "alpha", "5", "4")
				completed, err := done(root)
				if err != nil {
					t.Fatalf("default warnings made evals Done fail: %v", err)
				}
				for _, code := range []string{"B01", "B02", "B03"} {
					requireCostWarning(t, completed.Diagnostics, code)
				}
			}
		})
	}
}

func Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  eval_ratio: 1\n", "", "", "", "")
package evals_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func releaseRefs(t *testing.T, root string) string {
	t.Helper()
	return git(t, root, "for-each-ref", "--format=%(refname) %(objectname)")
}

func Test_FLOW_084_ReleaseOwnedClaimPreservesDirtyWorkAndRecordedState(t *testing.T) {
	root := fixture(t, "", "", "### new:release-work\nRelease work MUST remain.\n", "", "")
	started := start(t, root)
	if started.Run == "" || started.Phase != "evals" || started.ClaimOID == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("Start did not assign runtime state and ID: %#v", started)
	}
	changePath := "changes/0001-manual.md"
	change := readFile(t, root, changePath)
	if !strings.Contains(change, "run: "+started.Run) || !strings.Contains(change, "phase: "+started.Phase) || !strings.Contains(change, "### ALP-003") {
		t.Fatalf("Start did not record runtime state and assigned ID: %q", change)
	}
	write(t, root, changePath, strings.Replace(change, "Use the revised policy.", "Use the revised policy.\nWorker draft remains.", 1))
	write(t, root, "packages/alpha/internal/worker.txt", "unfinished worker work\n")
	write(t, root, "worker-untracked.txt", "untracked worker work\n")
	git(t, root, "update-ref", "refs/fireweed/fixture-unrelated", head(t, root))
	beforeHead := head(t, root)
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeRefs := releaseRefs(t, root)
	beforeChange := readFile(t, root, changePath)
	beforeTracked := readFile(t, root, "packages/alpha/internal/worker.txt")
	beforeUntracked := readFile(t, root, "worker-untracked.txt")
	if err := workflow.Release(root, "manual"); err != nil {
		t.Fatalf("Release owned claim: %v", err)
	}
	if got := head(t, root); got != beforeHead {
		t.Errorf("HEAD moved: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("worktree status changed: %q -> %q", beforeStatus, got)
	}
	for path, want := range map[string]string{
		changePath:                           beforeChange,
		"packages/alpha/internal/worker.txt": beforeTracked,
		"worker-untracked.txt":               beforeUntracked,
	} {
		if got := readFile(t, root, path); got != want {
			t.Errorf("%s changed: %q -> %q", path, want, got)
		}
	}
	if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
		t.Errorf("released claim remains current: %v", err)
	}
	wantRefs := strings.Replace(beforeRefs, "refs/fireweed/claims/manual "+started.ClaimOID+"\n", "", 1)
	if wantRefs == beforeRefs {
		t.Fatalf("Start claim absent from refs: %q", beforeRefs)
	}
	if got := releaseRefs(t, root); got != wantRefs {
		t.Errorf("Release changed refs beyond owned claim: got %q, want %q", got, wantRefs)
	}
}

func Test_FLOW_085_ReleaseMissingAndForeignClaimsPreserveRefs(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	before := releaseRefs(t, root)
	if err := workflow.Release(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
		t.Errorf("missing claim Release = %v, want P02", err)
	}
	if got := releaseRefs(t, root); got != before {
		t.Errorf("missing claim changed refs: %q -> %q", before, got)
	}

	other := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
	owner := start(t, root)
	write(t, root, "owner-work.txt", "owner's unfinished work\n")
	write(t, other, "foreign-work.txt", "other worktree's unfinished work\n")
	ownerHead, foreignHead := head(t, root), head(t, other)
	ownerStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	foreignStatus := git(t, other, "status", "--porcelain=v1", "--untracked-files=all")
	ownerChange := readFile(t, root, "changes/0001-manual.md")
	foreignChange := readFile(t, other, "changes/0001-manual.md")
	before = releaseRefs(t, root)
	if err := workflow.Release(other, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
		t.Errorf("foreign claim Release = %v, want P02", err)
	}
	if got := releaseRefs(t, root); got != before {
		t.Errorf("foreign claim changed refs: %q -> %q", before, got)
	}
	if head(t, root) != ownerHead || head(t, other) != foreignHead ||
		git(t, root, "status", "--porcelain=v1", "--untracked-files=all") != ownerStatus ||
		git(t, other, "status", "--porcelain=v1", "--untracked-files=all") != foreignStatus ||
		readFile(t, root, "changes/0001-manual.md") != ownerChange ||
		readFile(t, other, "changes/0001-manual.md") != foreignChange ||
		readFile(t, root, "owner-work.txt") != "owner's unfinished work\n" ||
		readFile(t, other, "foreign-work.txt") != "other worktree's unfinished work\n" {
		t.Error("foreign Release changed HEAD, status or working files")
	}
	current, err := workflow.Current(root, "manual")
	if err != nil || current.ClaimOID != owner.ClaimOID || current.Run != owner.Run {
		t.Errorf("owner's exact claim changed: current = %#v, error = %v", current, err)
	}
}

func Test_FLOW_086_ReleaseAcceptsCurrentChangeSelectors(t *testing.T) {
	for _, selector := range []struct {
		name, value string
	}{
		{"path", "changes/0001-manual.md"},
		{"filename", "0001-manual.md"},
		{"slug", "manual"},
	} {
		t.Run(selector.name, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			started := start(t, root)
			selected := selector.value
			current, err := workflow.Current(root, selected)
			if err != nil || current.ClaimOID != started.ClaimOID {
				t.Fatalf("Current selector %q = %#v, %v", selected, current, err)
			}
			if err := workflow.Release(root, selected); err != nil {
				t.Fatalf("Release selector %q: %v", selected, err)
			}
			if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "P02:") {
				t.Errorf("selector %q left claim current: %v", selected, err)
			}
		})
	}
}

exec
/bin/zsh -lc "rg -n 'project.Project|BaseRef|fixes:|already_satisfied|Audit|worker_writable|P2-' packages/workflow/evals/eligible_test.go packages/workflow/SPEC.md; cat packages/workflow/evals/current_absence_test.go; sed -n '66,134p' packages/workflow/evals/costs_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 479ms:
packages/workflow/evals/eligible_test.go:46:		if _, err := project.Project(root, ch.Slug); err != nil {
packages/workflow/SPEC.md:91:evidence for every enforced change ID that is not already_satisfied.
packages/workflow/SPEC.md:125:Audit Done MUST discard a finding whose spec ID is absent or whose quoted clause is not an exact substring of that spec.
packages/workflow/SPEC.md:128:On re-audit, Audit Done MUST discard findings for spec IDs that had no retained finding in the preceding audit round.
packages/workflow/SPEC.md:131:Audit Done MUST require every enforced ID in the change's Added, Modified,
packages/workflow/SPEC.md:141:Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.
packages/workflow/SPEC.md:144:Audit Done MUST block the change when a retained finding marks its quoted clause untestable.
packages/workflow/SPEC.md:147:Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.
packages/workflow/SPEC.md:153:Audit blocking MUST release the calling worktree's held change claim.
packages/workflow/SPEC.md:438:PRBody MUST identify interface edits, new disposable stubs added in evals, module manifest/other-file edits and worker_writable edits associated with each selected change.
package evals_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lurkshark/fireweed/packages/workflow"
)

func Test_FLOW_133_CurrentEmptySelectorWithoutOwnedClaim(t *testing.T) {
	assertAbsent := func(t *testing.T, root string) {
		t.Helper()
		_, err := workflow.Current(root, "")
		if err == nil || err.Error() != "P02: worktree holds no change" {
			t.Errorf("Current(root, empty) error = %v; want exact P02 absence signal", err)
		}
	}

	t.Run("unclaimed queued change", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		assertAbsent(t, root)
	})

	t.Run("empty valid queue", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
			t.Fatal(err)
		}
		assertAbsent(t, root)
	})

	t.Run("claim owned by another worktree", func(t *testing.T) {
		root := fixture(t, "", "", "", "", "")
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		start(t, root)
		assertAbsent(t, other)
	})
}
func Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  if [ -f packages/alpha/evals/second ]; then seconds=0.75
  else seconds=0.25; fi
  printf '<testsuite><testcase name="ALP-001" time="%s">%s</testcase><testcase name="ALP-002" time="0.1"/></testsuite>' "$seconds" "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="0.2"/></testsuite>' > "$2"
`)
	commit(t, root, "deterministic cost report")
	state := start(t, root)
	before := time.Now().UTC()
	first, err := verify(root, true)
	if err != nil {
		t.Fatalf("first Verify: %v", err)
	}
	if cost, ok := costFor(first.Costs, "alpha"); !ok || math.Abs(cost.Seconds-0.35) > 1e-9 {
		t.Errorf("FLOW-097 Verify costs = %#v; want alpha at 0.35 seconds", first.Costs)
	}
	savedFirst, okFirst := costAt(t, root, state.Run, "evals")
	if okFirst {
		if savedFirst.Run != state.Run || savedFirst.Phase != "evals" || savedFirst.RecordedAt.IsZero() || savedFirst.RecordedAt.Location() != time.UTC || savedFirst.RecordedAt.Before(before) || savedFirst.RecordedAt.After(time.Now().UTC()) || !reflect.DeepEqual(savedFirst.Modules, first.Costs) {
			t.Errorf("FLOW-097 saved evals identity/time/modules = %#v; returned costs %#v", savedFirst, first.Costs)
		}
	}
	write(t, root, "packages/alpha/evals/second", "second\n")
	second, err := verify(root, true)
	if err != nil {
		t.Fatalf("second Verify: %v", err)
	}
	if cost, ok := costFor(second.Costs, "alpha"); !ok || math.Abs(cost.Seconds-0.85) > 1e-9 {
		t.Errorf("FLOW-097 replacement Verify costs = %#v; want alpha at 0.85 seconds", second.Costs)
	}
	savedSecond, okSecond := costAt(t, root, state.Run, "evals")
	if okSecond && (savedSecond.Run != state.Run || savedSecond.Phase != "evals" || !reflect.DeepEqual(savedSecond.Modules, second.Costs) || okFirst && savedSecond.RecordedAt.Before(savedFirst.RecordedAt)) {
		t.Errorf("FLOW-097 replacement snapshot = %#v; first %#v, returned %#v", savedSecond, savedFirst, second.Costs)
	}
	retained, err := workflow.Costs(root)
	if err != nil || len(retained) != 1 || retained[0].Phase != "evals" || !reflect.DeepEqual(retained[0].Modules, second.Costs) {
		t.Errorf("FLOW-097 repeated Verify retained snapshots = %#v, %v; want one replaced evals snapshot", retained, err)
	}
	evalsDone, err := done(root)
	if err != nil {
		t.Fatalf("evals Done: %v", err)
	}
	if cost, ok := costFor(evalsDone.Costs, "alpha"); !ok || math.Abs(cost.Seconds-0.85) > 1e-9 {
		t.Errorf("FLOW-097 evals Done costs = %#v", evalsDone.Costs)
	}
	if saved, ok := costAt(t, root, state.Run, "evals"); ok && !reflect.DeepEqual(saved.Modules, evalsDone.Costs) {
		t.Errorf("FLOW-097 evals Done snapshot modules = %#v; result %#v", saved.Modules, evalsDone.Costs)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	implDone, err := done(root)
	if err != nil {
		t.Fatalf("impl Done: %v", err)
	}
	if saved, ok := costAt(t, root, state.Run, "impl"); ok && (saved.Run != state.Run || saved.Phase != "impl" || saved.RecordedAt.Location() != time.UTC || !reflect.DeepEqual(saved.Modules, implDone.Costs)) {
		t.Errorf("FLOW-097 impl Done snapshot = %#v; result costs %#v", saved, implDone.Costs)
	}
	if _, ok := costFor(implDone.Costs, "beta"); !ok {
		t.Errorf("FLOW-097 full impl Done omitted evaluated dependent: %#v", implDone.Costs)
	}
}

func Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts(t *testing.T) {

codex
The drill’s SYSTEM.md reserves builds and evals for a separate verifier. I’ll use the retained evals to guide the rebuild and format and review the new source locally.
exec
/bin/zsh -lc "mkdir -p packages/workflow/internal/lifecycle
cat > packages/workflow/internal/lifecycle/common.go <<'EOF'
package lifecycle

import (
	\"bytes\"
	\"context\"
	\"encoding/json\"
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"os/exec\"
	\"path\"
	\"path/filepath\"
	\"regexp\"
	\"sort\"
	\"strconv\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
)

const (
	pendingDir = \"changes\"
	archiveDir = \"changes/archive\"
	claimPrefix = \"refs/fireweed/claims/\"
	attemptPrefix = \"refs/fireweed/attempts/\"
	drillPrefix = \"refs/fireweed/drills/\"
)

func git(ctx context.Context, root string, env []string, input []byte, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, \"git\", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdin = bytes.NewReader(input)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err "'!= nil {
		if ctx.Err() != nil {
			err = ctx.Err()
		}
		return nil, fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(stderr.String()))
	}
	return out, nil
}

func gitOutput(root string, args ...string) (string, error) {
	out, err := git(context.Background(), root, nil, nil, args...)
	return strings.TrimSpace(string(out)), err
}

func repository(root string) (gitstore.Repository, error) {
	repo, err := gitstore.Inspect(root)
	if err == nil {
		repo.Root, err = filepath.EvalSymlinks(repo.Root)
	}
	if err == nil {
		repo.Root, err = filepath.Abs(repo.Root)
	}
	return repo, err
}

func coded(code, subject, message string) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: errors.New(code + ": " + message)}
}

func withCode(code, subject string, err error) error {
	if err == nil {
		return nil
	}
	if strings.HasPrefix(err.Error(), code+":") {
		return err
	}
	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s: %w", code, err)}
}

func diagnosticError(d project.Diagnostic) error {
	if d.Code == "" {
		return errors.New(d.Message)
	}
	return withCode(d.Code, d.Path, errors.New(d.Message))
}

func firstError(ds []project.Diagnostic) error {
	for _, d := range ds {
		if !d.Warning {
			return diagnosticError(d)
		}
	}
	return nil
}

func staleError(err error) error {
	if err != nil && strings.Contains(err.Error(), "stale base") {
		return withCode("P05", "", err)
	}
	return err
}

func stringValue(v any) string {
	s, _ := v.(string)
	return s
}

func intValue(v any, fallback int) int {
	switch v := v.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case json.Number:
		if n, err := v.Int64(); err == nil {
			return int(n)
		}
	case string:
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return fallback
}

func stringList(v any) []string {
	switch v := v.(type) {
	case []string:
		return v
	case []any:
		result := make([]string, 0, len(v))
		for _, entry := range v {
			if s, ok := entry.(string); ok {
				result = append(result, s)
			}
		}
		return result
	}
	return nil
}

func configGroup(config map[string]any, key string) map[string]any {
	group, _ := config[key].(map[string]any)
	return group
}

func field(c project.Change, key string) string { return stringValue(c.Document.Frontmatter[key]) }
func section(doc contracts.Document, name string) string {
	s, _ := contracts.Section(doc.Body, name)
	return s
}

func moduleNamed(s project.Snapshot, name string) (project.Module, error) {
	for _, m := range s.Modules {
		if m.Name == name {
			return m, nil
		}
	}
	return project.Module{}, fmt.Errorf("module %q not found", name)
}

func specs(m project.Module) ([]contracts.Item, error) {
	return contracts.ParseItems(section(m.Document, "Specs"))
}

func dependsOn(s project.Snapshot, from, target string, transitive bool) bool {
	seen := make(map[string]bool)
	var visit func(string) bool
	visit = func(name string) bool {
		if seen[name] {
			return false
		}
		seen[name] = true
		m, _ := moduleNamed(s, name)
		for _, dependency := range m.DependsOn {
			if dependency == target || transitive && visit(dependency) {
				return true
			}
		}
		return false
	}
	return visit(from)
}

func relativeTo(dir, file string) (string, bool) {
	if dir == "" || dir == "." {
		return file, true
	}
	rel, ok := strings.CutPrefix(file, strings.TrimSuffix(dir, "/")+"/")
	return rel, ok
}

func matchesAny(patterns []string, name string) bool {
	for _, pattern := range patterns {
		if contracts.MatchGlob(pattern, name) {
			return true
		}
	}
	return false
}

func zone(s project.Snapshot, m project.Module, name string) string {
	if matchesAny(stringList(s.Config["protected"]), name) {
		return "protected"
	}
	relative, inside := relativeTo(m.Dir, name)
	if inside {
		return contracts.ResolveZone(relative, m.Zones)
	}
	return ""
}

func belongsTo(s project.Snapshot, m project.Module, name string) bool {
	if _, inside := relativeTo(m.Dir, name); !inside {
		return false
	}
	for _, other := range s.Modules {
		if other.Dir != m.Dir && len(other.Dir) > len(m.Dir) {
			if _, inside := relativeTo(other.Dir, name); inside {
				return false
			}
		}
	}
	return true
}

func moduleFiles(s project.Snapshot, m project.Module, visit func(string, fs.DirEntry) error) error {
	dir := filepath.Join(s.Root, filepath.FromSlash(m.Dir))
	return filepath.WalkDir(dir, func(file string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(s.Root, file)
		if err != nil {
			return err
		}
		name := filepath.ToSlash(relative)
		if entry.IsDir() {
			if file != dir && (entry.Name() == ".git" || entry.Name() == ".fireweed" || !belongsTo(s, m, name)) {
				return filepath.SkipDir
			}
			return nil
		}
		return visit(name, entry)
	})
}

var numberedName = regexp.MustCompile(`''^([0-9]+)-(.+)$`)

func identity(file string) project.Change {
	stem := strings.TrimSuffix(path.Base(file), ".md")
	change := project.Change{Path: file, Slug: stem}
	if parts := numberedName.FindStringSubmatch(stem); parts != nil {
		change.Number, _ = strconv.Atoi(parts[1])
		change.Slug = parts[2]
	}
	return change
}

func changePaths(root, dir string) ([]project.Change, error) {
	entries, err := os.ReadDir(filepath.Join(root, dir))
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var result []project.Change
	for _, entry := range entries {
		if !entry.IsDir() && path.Ext(entry.Name()) == ".md" {
			result = append(result, identity(path.Join(dir, entry.Name())))
		}
	}
	return result, nil
}

func pathsAt(root, revision, dir string) ([]project.Change, error) {
	out, err := git(context.Background(), root, nil, nil, "ls-tree", "-z", "--name-only", revision, dir+"/")
	if err != nil {
		return nil, err
	}
	var result []project.Change
	for _, name := range splitNUL(out) {
		if path.Dir(name) == dir && path.Ext(name) == ".md" {
			result = append(result, identity(name))
		}
	}
	return result, nil
}

func splitNUL(out []byte) []string {
	return strings.Split(strings.TrimSuffix(string(out), "'"\\x00\"), \"\\x00\")[:nulFields(out)]
}

func nulFields(out []byte) int {
	if len(out) == 0 {
		return 0
	}
	count := bytes.Count(out, []byte{0})
	if out[len(out)-1] "'!= 0 {
		count++
	}
	return count
}

func selectChange(queue []project.Change, selector string) (project.Change, error) {
	var matches []project.Change
	slugs := map[string]bool{}
	clean := filepath.ToSlash(filepath.Clean(selector))
	for _, ch := range queue {
		slugs[ch.Slug] = true
		name := path.Base(ch.Path)
		if clean == ch.Path || selector == name || selector == ch.Slug || selector == strings.TrimSuffix(name, ".md") {
			matches = append(matches, ch)
		}
	}
	if len(matches) == 1 {
		return matches[0], nil
	}
	reason := "not found"
	if len(matches) > 1 {
		reason = "ambiguous"
	}
	return project.Change{}, coded("U02", selector, fmt.Sprintf("change %q %s (valid slugs: %s)", selector, reason, strings.Join(sortedKeys(slugs), ", ")))
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func readChange(ch project.Change, data []byte) (project.Change, error) {
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return ch, err
	}
	ch.Document = doc
	ch.Module = stringValue(doc.Frontmatter["module"])
	for _, edit := range []struct{
		name string
		items *[]contracts.Item
	}{{"Added", &ch.Added}, {"Modified", &ch.Modified}, {"Removed", &ch.Removed}} {
		*edit.items, err = contracts.ParseItems(section(doc, edit.name))
		if err != nil {
			return ch, err
		}
	}
	return ch, nil
}

func writeFile(file string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(file), ".fireweed-write-")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	mode := fs.FileMode(0644)
	if info, err := os.Stat(file); err == nil {
		mode = info.Mode().Perm()
	}
	if err := tmp.Chmod(mode); err != nil {
		tmp.Close()
		return err
	}
	_, writeErr := tmp.Write(data)
	if err := errors.Join(writeErr, tmp.Close()); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), file)
}

func writeJSON(file string, data any) error {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(file, append(b, '"'\\n'))
}

func requireClean(root, action string) error {
	status, err := gitOutput(root, \"status\", \"--porcelain=v1\", \"--untracked-files=all\")
	if err == nil && status "'!= "" {
		err = coded("P01", root, action+" requires clean worktree")
	}
	return err
}

func parseDate(v any) time.Time {
	if date, ok := v.(time.Time); ok {
		return date.UTC()
	}
	for _, layout := range []string{time.RFC3339Nano, time.DateOnly} {
		if date, err := time.Parse(layout, stringValue(v)); err == nil {
			return date.UTC()
		}
	}
	return time.Time{}
}

func implementationDate(ch project.Change) time.Time {
	var latest time.Time
	for _, line := range strings.Split(section(ch.Document, "Provenance"), "'"\\n\") {
		fields := strings.Fields(line)
		if len(fields) >= 3 && fields[0] == \"-\" && fields[1] == \"impl\" {
			if date := parseDate(fields[2]); date.After(latest) {
				latest = date
			}
		}
	}
	return latest
}

func drillReport(root, module string) (time.Time, bool, error) {
	data, err := os.ReadFile(filepath.Join(root, \".fireweed\", \"drills\", module+\".md\"))
	if errors.Is(err, fs.ErrNotExist) {
		return time.Time{}, false, nil
	}
	if err "'!= nil {
		return time.Time{}, false, err
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return time.Time{}, false, err
	}
	return parseDate(doc.Frontmatter["date"]), stringValue(doc.Frontmatter["result"]) == "pass", nil
}
EOF
cat > packages/workflow/internal/lifecycle/queue.go <<'"'EOF'
package lifecycle

import (
	\"fmt\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

func waitingQueue(root string, queue []project.Change) ([]model.StatusEntry, error) {
	entries := make([]model.StatusEntry, 0, len(queue))
	var blocked []project.Change
	for _, ch := range queue {
		entry := model.StatusEntry{
			Slug: ch.Slug, Path: ch.Path, Module: ch.Module,
			Phase: field(ch, \"phase\"), Blocked: field(ch, \"blocked\"),
			AuditRounds: intValue(ch.Document.Frontmatter[\"audit_rounds\"], 0),
		}
		if entry.Phase == \"\" {
			entry.Phase = \"queued\"
		}
		if entry.Blocked "'!= "" {
			blocked = append(blocked, ch)
		} else if len(blocked) > 0 {
			projection, err := project.Project(root, ch.Slug)
			if err != nil {
				return nil, err
			}
			interfaceEdit := strings.TrimSpace(section(ch.Document, "Interface")) != ""
			for _, earlier := range blocked {
				switch {
				case earlier.Module == ch.Module:
					entry.WaitingRule = "same_module"
				case dependsOn(projection, ch.Module, earlier.Module, true):
					entry.WaitingRule = "depends_on_blocked"
				case interfaceEdit && dependsOn(projection, earlier.Module, ch.Module, true):
					entry.WaitingRule = "interface_dependency"
				}
				if entry.WaitingRule != "" {
					entry.WaitingOn = strings.TrimSuffix(earlier.Path[strings.LastIndex(earlier.Path, "/")+1:], ".md")
					break
				}
			}
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func Eligible(root string) ([]project.Change, error) {
	repo, err := repository(root)
	if err != nil {
		return nil, err
	}
	queue, err := project.Queue(repo.Root)
	if err != nil {
		return nil, err
	}
	return eligible(repo.Root, queue)
}

func eligible(root string, queue []project.Change) ([]project.Change, error) {
	entries, err := waitingQueue(root, queue)
	if err != nil {
		return nil, err
	}
	result := make([]project.Change, 0, len(queue))
	for i, e := range entries {
		if e.Blocked == "" && e.WaitingRule == "" {
			result = append(result, queue[i])
		}
	}
	return result, nil
}

func checkStale(root string, ch project.Change) error {
	stale, err := project.StaleSpecs(root, ch.Slug)
	if err != nil {
		return staleError(err)
	}
	for _, spec := range stale {
		if !spec.Recovered {
			return coded("P05", ch.Path, fmt.Sprintf("stale base for %s: recorded %s, current %s", spec.ID, spec.RecordedHash, spec.Current))
		}
	}
	return nil
}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 692ms:
exec
/bin/zsh -lc "cat > packages/workflow/internal/lifecycle/claim.go <<'EOF'
package lifecycle

import (
	\"context\"
	\"crypto/rand\"
	\"encoding/hex\"
	\"encoding/json\"
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path\"
	\"path/filepath\"
	\"reflect\"
	\"regexp\"
	\"strconv\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

type claim struct {
	Worktree string           "'`json:"worktree"`
	Change   string           `json:"change"`
	Path     string           `json:"path"`
	Base     string           `json:"phase_base"`
	Approved string           `json:"approved"`
	Snapshot project.Snapshot `json:"snapshot"`
}

func (c claim) ref() string { return claimPrefix+c.Change }

func (c claim) blob(root string) (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return gitstore.WriteBlob(root, b)
}

func readClaim(root, oid string) (claim, error) {
	var c claim
	b, err := gitstore.ReadBlob(root, oid)
	if err != nil {
		return c, err
	}
	decoder := json.NewDecoder(strings.NewReader(string(b)))
	decoder.UseNumber()
	err = decoder.Decode(&c)
	c.Snapshot.Config = normalizeMap(c.Snapshot.Config)
	for i := range c.Snapshot.Modules {
		doc := &c.Snapshot.Modules[i].Document
		doc.Frontmatter = normalizeMap(doc.Frontmatter)
	}
	return c, err
}

func normalizeMap(m map[string]any) map[string]any {
	for key, value := range m {
		m[key] = normalizeNumber(value)
	}
	return m
}

func normalizeNumber(value any) any {
	switch v := value.(type) {
	case json.Number:
		if n, err := strconv.Atoi(string(v)); err == nil {
			return n
		}
		if n, err := v.Float64(); err == nil {
			return n
		}
	case map[string]any:
		return normalizeMap(v)
	case []any:
		for i := range v {
			v[i] = normalizeNumber(v[i])
		}
	}
	return value
}

func heldClaim(repo gitstore.Repository, selector string) (claim, string, error) {
	root := repo.Root
	var empty claim
	if selector == "" {
		refs, err := listRefs(context.Background(), root, claimPrefix)
		if err != nil {
			return empty, "", err
		}
		for _, ref := range refs {
			c, err := readClaim(root, ref.oid)
			if err == nil && c.Worktree == root && c.Path != "" && c.ref() == ref.name {
				return c, ref.oid, nil
			}
		}
		return empty, "", coded("P02", root, "worktree holds no change")
	}
	queue, err := changePaths(root, pendingDir)
	if err != nil {
		return empty, "", err
	}
	ch, err := selectChange(queue, selector)
	if err != nil {
		// A deleted held document still belongs to its claim. Verify can then
		// classify the missing document as a change to approved state.
		if c, oid, claimErr := heldClaim(repo, ""); claimErr == nil {
			found := false
			for _, entry := range queue {
				found = found || entry.Path == c.Path
			}
			if !found {
				if _, matchErr := selectChange([]project.Change{identity(c.Path)}, selector); matchErr == nil {
					return c, oid, nil
				}
			}
		}
		return empty, "", err
	}
	oid, exists, err := gitstore.Ref(root, claimPrefix+ch.Slug)
	if err != nil {
		return empty, "", err
	}
	if !exists {
		return empty, "", coded("P02", ch.Path, "change has no claim")
	}
	c, err := readClaim(root, oid)
	if err != nil {
		return empty, "", withCode("P02", ch.Path, err)
	}
	if c.Worktree != root {
		return empty, "", coded("P02", ch.Path, "claim belongs to another worktree")
	}
	if c.Path != ch.Path || c.Change != ch.Slug {
		return empty, "", coded("P02", ch.Path, "claim belongs to another change")
	}
	return c, oid, nil
}

func (c claim) state(repo gitstore.Repository, oid string) (model.State, error) {
	approved, err := readChange(identity(c.Path), []byte(c.Approved))
	if err != nil {
		return model.State{}, err
	}
	c.Snapshot.Root = repo.Root
	return model.State{
		Change: approved, Snapshot: c.Snapshot, Run: field(approved, "run"),
		Phase: field(approved, "phase"), Head: repo.Head, PhaseBase: c.Base, ClaimOID: oid,
	}, nil
}

func Current(root, selector string) (model.State, error) {
	repo, err := repository(root)
	if err != nil {
		return model.State{}, err
	}
	c, oid, err := heldClaim(repo, selector)
	if err != nil {
		return model.State{}, err
	}
	s, err := c.state(repo, oid)
	if err != nil {
		return s, err
	}
	if data, err := os.ReadFile(filepath.Join(repo.Root, c.Path)); err == nil {
		if working, err := readChange(s.Change, data); err == nil {
			s.Change = working
		}
	}
	return s, nil
}

func Release(root, selector string) error {
	repo, err := repository(root)
	if err != nil {
		return err
	}
	c, oid, err := heldClaim(repo, selector)
	if err != nil {
		return err
	}
	return withCode("P02", c.Path, gitstore.CASRef(repo.Root, c.ref(), oid, ""))
}

func Start(root, selector string) (model.State, error) {
	return start(root, selector, nil)
}

func StartFromAttempt(root, selector string, attempt int) (model.State, error) {
	return start(root, selector, &attempt)
}

func start(root, selector string, attempt *int) (s model.State, err error) {
	repo, err := repository(root)
	if err != nil {
		return s, err
	}
	root = repo.Root
	if err = requireClean(root, "start"); err != nil {
		// A second Start from the owning checkout is an ownership error.
		if c, _, claimErr := heldClaim(repo, selector); claimErr == nil {
			return s, coded("P02", c.Path, "claim already held by this worktree")
		}
		return s, err
	}
	queue, err := project.Queue(root)
	if err != nil {
		return s, staleError(err)
	}
	var chosen project.Change
	if selector != "" {
		chosen, err = selectChange(queue, selector)
		if err != nil {
			return s, err
		}
	}
	ready, err := eligible(root, queue)
	if err != nil {
		return s, staleError(err)
	}
	if len(ready) == 0 || selector != "" && chosen.Path != ready[0].Path {
		return s, coded("P01", selector, "change is not eligible queue head")
	}
	chosen = ready[0]
	var savedOID string
	if attempt != nil {
		if *attempt <= 0 {
			return s, coded("P01", chosen.Path, "attempt number must be positive")
		}
		ref := attemptPrefix+chosen.Slug+"/"+strconv.Itoa(*attempt)
		var exists bool
		savedOID, exists, err = gitstore.Ref(root, ref)
		if err != nil {
			return s, err
		}
		if !exists {
			return s, coded("P01", chosen.Path, "saved attempt not found: "+ref)
		}
		if field(chosen, "phase") == "" || field(chosen, "run") == "" {
			return s, coded("P01", chosen.Path, "resuming an attempt requires a recorded run and phase")
		}
	}
	s.Snapshot, err = project.Project(root, chosen.Slug)
	if err != nil {
		return s, staleError(err)
	}
	s.Snapshot.Root = root
	for _, d := range project.Check(root) {
		if !d.Warning && (d.Path == "" || d.Path == chosen.Path || !strings.HasPrefix(d.Path, "changes/")) {
			return s, staleError(diagnosticError(d))
		}
	}
	if err = checkStale(root, chosen); err != nil {
		return s, err
	}

	shared, local, err := baseRefs(root, s.Snapshot.Config)
	if err != nil {
		return s, err
	}
	applied, err := archivedAt(root, repo.Head)
	if err != nil {
		return s, err
	}
	unapproved, err := unapprovedCount(root, shared, applied)
	if err != nil {
		return s, err
	}
	limit := intValue(configGroup(s.Snapshot.Config, "authoring")["max_unapproved"], 5)
	if unapproved >= limit {
		message := fmt.Sprintf("%d unapproved archived changes reaches the limit of %d", unapproved, limit)
		s.Diagnostics = append(s.Diagnostics, project.Diagnostic{Code: "P07", Path: chosen.Path, Message: message})
		return s, coded("P07", chosen.Path, message)
	}
	if !existsAt(root, local, chosen.Path) && !existsAt(root, shared, chosen.Path) {
		s.Diagnostics = append(s.Diagnostics, project.Diagnostic{
			Code: "P04", Path: chosen.Path, Warning: true, Message: "change is absent from the local and shared bases",
		})
	}

	file := filepath.Join(root, chosen.Path)
	original, err := os.ReadFile(file)
	if err != nil {
		return s, err
	}
	if attempt == nil {
		assigned, assignErr := assignIDs(root, chosen, s.Snapshot, queue, applied, string(original))
		if assignErr != nil {
			return s, assignErr
		}
		chosen, err = readChange(chosen, []byte(assigned))
		if err != nil {
			return s, err
		}
	}
	front := chosen.Document.Frontmatter
	phase := field(chosen, "phase")
	unstarted := phase == ""
	regenerate := front["regenerate"] == true
	if regenerate {
		var latest time.Time
		for _, archive := range applied {
			if archive.Module == chosen.Module {
				if date := implementationDate(archive); date.After(latest) {
					latest = date
				}
			}
		}
		date, passing, reportErr := drillReport(root, chosen.Module)
		if reportErr != nil {
			return s, reportErr
		}
		if !passing || !date.After(latest) {
			s.Diagnostics = append(s.Diagnostics, project.Diagnostic{
				Code: "P10", Path: chosen.Path, Warning: true, Message: "module "+chosen.Module+" has no passing drill newer than its latest implementation",
			})
		}
	}
	if unstarted {
		phase = "evals"
		if regenerate || !hasSpecWork(chosen) {
			phase = "impl"
		}
		front["phase"] = phase
	}
	if field(chosen, "run") == "" {
		var id [16]byte
		if _, err = rand.Read(id[:]); err != nil {
			return s, err
		}
		front["run"] = hex.EncodeToString(id[:])
	}
	if err = validateRun(field(chosen, "run")); err != nil {
		return s, err
	}
	approved, err := contracts.RenderDocument(chosen.Document)
	if err != nil {
		return s, err
	}
	if attempt != nil {
		approved = string(original)
	}

	// Reproject with real IDs before storing the claim. The original projection
	// may use provisional IDs, so replace its spec section from the preceding
	// projection and the approved change'"'s assigned edits.
	if err = installAssignedSpecs(root, queue, chosen, &s.Snapshot); err "'!= nil {
		return s, staleError(err)
	}
	c := claim{Worktree: root, Change: chosen.Slug, Path: chosen.Path, Base: repo.Head, Approved: approved, Snapshot: s.Snapshot}
	oid, err := c.blob(root)
	if err != nil {
		return s, err
	}
	if err = gitstore.CASRef(root, c.ref(), "", oid); err != nil {
		return s, withCode("P02", c.Path, fmt.Errorf("claim already held: %w", err))
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, gitstore.CASRef(root, c.ref(), oid, ""))
		}
	}()
	if repoNow, inspectErr := repository(root); inspectErr != nil {
		return s, inspectErr
	} else if repoNow.Head != repo.Head {
		return s, coded("P05", c.Path, "stale phase base: HEAD changed during Start")
	}
	diagnostics := s.Diagnostics
	s, err = c.state(repo, oid)
	if err != nil {
		return s, err
	}
	s.Diagnostics = diagnostics
	working := approved
	if attempt != nil {
		working, err = resumeDocument(root, savedOID, s.Change)
		if err != nil {
			return s, err
		}
	}
	if working != string(original) {
		if err = writeFile(file, []byte(working)); err != nil {
			return s, err
		}
		defer func() {
			if err != nil {
				err = errors.Join(err, writeFile(file, original))
			}
		}()
	}
	if attempt != nil {
		s.Change, err = readChange(s.Change, []byte(working))
		if err == nil {
			err = restoreAttempt(root, savedOID, chosen.Path)
		}
	} else if unstarted && regenerate {
		err = removeDisposable(s.Snapshot, chosen.Module)
	}
	return s, err
}

func validateRun(run string) error {
	if run == "" || run == "." || run == ".." || strings.ContainsAny(run, "/'"\\\\\\x00\") {
		return coded(\"P01\", run, fmt.Sprintf(\"invalid run identifier %q\", run))
	}
	return nil
}

func hasSpecWork(c project.Change) bool {
	return len(c.Added)+len(c.Modified)+len(c.Removed)+len(stringList(c.Document.Frontmatter[\"fixes\"]))+len(stringList(c.Document.Frontmatter[\"eval_fix\"])) > 0
}

var addedHandle = regexp.MustCompile("'`new:[[:alnum:]_][[:alnum:]_.-]*`)

func assignIDs(root string, ch project.Change, projection project.Snapshot, queue, archives []project.Change, source string) (string, error) {
	hasHandles := false
	for _, item := range ch.Added {
		hasHandles = hasHandles || strings.HasPrefix(item.ID, "new:")
	}
	if !hasHandles {
		return source, nil
	}
	applied, err := project.Load(root)
	if err != nil {
		return "", err
	}
	m, err := moduleNamed(applied, ch.Module)
	if err != nil {
		return "", err
	}
	before, err := specs(m)
	if err != nil {
		return "", err
	}
	afterModule, err := moduleNamed(projection, ch.Module)
	if err != nil {
		return "", err
	}
	after, err := specs(afterModule)
	if err != nil {
		return "", err
	}
	reserved := map[string]bool{}
	for _, list := range [][]contracts.Item{before, after, ch.Added, ch.Removed} {
		for _, item := range list {
			reserved[item.ID] = true
		}
	}
	for _, list := range [][]project.Change{archives, queue} {
		for _, other := range list {
			if other.Module == ch.Module && other.Path != ch.Path {
				for _, item := range other.Added {
					reserved[item.ID] = true
				}
				for _, item := range other.Removed {
					reserved[item.ID] = true
				}
			}
		}
	}
	replacements := map[string]string{}
	next := 1
	for _, item := range ch.Added {
		if !strings.HasPrefix(item.ID, "new:") {
			continue
		}
		for {
			id := fmt.Sprintf("%s-%03d", m.IDPrefix, next)
			next++
			if !reserved[id] {
				replacements[item.ID] = id
				reserved[id] = true
				break
			}
		}
	}
	return addedHandle.ReplaceAllStringFunc(source, func(handle string) string {
		if id, ok := replacements[handle]; ok {
			return id
		}
		// A prose reference can be followed by a sentence period.
		if id, ok := replacements[strings.TrimRight(handle, ".")]; ok {
			return id+handle[len(strings.TrimRight(handle, ".")):]
		}
		return handle
	}), nil
}

func installAssignedSpecs(root string, queue []project.Change, ch project.Change, projected *project.Snapshot) error {
	if len(ch.Added)+len(ch.Modified)+len(ch.Removed) == 0 {
		return nil
	}
	var prior string
	for _, other := range queue {
		if other.Path == ch.Path {
			break
		}
		prior = other.Slug
	}
	load := project.Load
	if prior != "" {
		load = func(root string) (project.Snapshot, error) { return project.Project(root, prior) }
	}
	before, err := load(root)
	if err != nil {
		return err
	}
	base, err := moduleNamed(before, ch.Module)
	if err != nil {
		return err
	}
	items, err := specs(base)
	if err != nil {
		return err
	}
	body := renderSpecEdits(items, ch)
	for i := range projected.Modules {
		if projected.Modules[i].Name == ch.Module {
			projected.Modules[i].Document.Body = contracts.ReplaceSection(projected.Modules[i].Document.Body, "Specs", body)
			return nil
		}
	}
	return fmt.Errorf("module %q missing from projected snapshot", ch.Module)
}

func renderSpecEdits(items []contracts.Item, ch project.Change) string {
	modified := map[string]contracts.Item{}
	removed := map[string]bool{}
	for _, edit := range ch.Modified {
		modified[edit.ID] = edit
	}
	for _, edit := range ch.Removed {
		removed[edit.ID] = true
	}
	var rendered []string
	for _, item := range items {
		if !removed[item.ID] {
			if edit, ok := modified[item.ID]; ok {
				item = edit
			}
			rendered = append(rendered, strings.TrimSpace(contracts.RenderItem(item)))
		}
	}
	for _, item := range ch.Added {
		rendered = append(rendered, strings.TrimSpace(contracts.RenderItem(item)))
	}
	return "'"\\n\"+strings.Join(rendered, \"\\n\\n\")+\"\\n\"
}

func removeDisposable(s project.Snapshot, name string) error {
	m, err := moduleNamed(s, name)
	if err "'!= nil {
		return err
	}
	var files []string
	err = moduleFiles(s, m, func(name string, entry fs.DirEntry) error {
		if zone(s, m, name) == "disposable" {
			files = append(files, name)
		}
		return nil
	})
	if err != nil || len(files) == 0 {
		return err
	}
	backup, err := os.MkdirTemp(s.Root, ".fireweed-regenerate-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(backup)
	for i, name := range files {
		if err = os.Rename(filepath.Join(s.Root, name), filepath.Join(backup, strconv.Itoa(i))); err != nil {
			for j, moved := range files[:i] {
				err = errors.Join(err, os.Rename(filepath.Join(backup, strconv.Itoa(j)), filepath.Join(s.Root, moved)))
			}
			return err
		}
	}
	return nil
}

func approvedWorking(s model.State) (project.Change, error) {
	ch := s.Change
	file := filepath.Join(s.Snapshot.Root, ch.Path)
	info, err := os.Lstat(file)
	if err != nil || !info.Mode().IsRegular() {
		return ch, coded("P03", ch.Path, "approved author metadata changed: change document is missing or is not a regular file")
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return ch, withCode("P03", ch.Path, err)
	}
	working, err := readChange(ch, data)
	if err != nil {
		return ch, withCode("P03", ch.Path, fmt.Errorf("approved author metadata changed: %w", err))
	}
	a, b := ch.Document.Frontmatter, working.Document.Frontmatter
	for _, key := range []string{"run", "phase"} {
		if !reflect.DeepEqual(a[key], b[key]) {
			return ch, coded("P03", ch.Path, "approved run or phase changed")
		}
	}
	for _, key := range []string{"audit_rounds", "blocked"} {
		if !reflect.DeepEqual(a[key], b[key]) {
			return ch, coded("P03", ch.Path, "approved audit state changed")
		}
	}
	for _, front := range []map[string]any{a, b} {
		for _, key := range []string{"run", "phase", "audit_rounds", "blocked"} {
			// Compare runtime fields separately without mutating either document.
			_ = front[key]
		}
	}
	authorFront := func(front map[string]any) map[string]any {
		copy := make(map[string]any)
		for key, value := range front {
			switch key {
			case "run", "phase", "audit_rounds", "blocked":
			default:
				copy[key] = value
			}
		}
		return copy
	}
	if !reflect.DeepEqual(authorFront(a), authorFront(b)) {
		return ch, coded("P03", ch.Path, "approved author metadata changed")
	}
	for _, name := range []string{"Provenance", "Audit"} {
		if name == "Audit" && s.Phase == "audit" {
			continue
		}
		if section(ch.Document, name) != section(working.Document, name) {
			return ch, coded("P03", ch.Path, "approved audit state changed: "+name)
		}
	}
	if immutableBody(ch.Document.Body, s.Phase) != immutableBody(working.Document.Body, s.Phase) {
		return ch, coded("P03", ch.Path, "approved author metadata changed")
	}
	return working, nil
}

func editableSection(name, phase string) bool {
	switch name {
	case "Decisions", "Questions":
		return phase == "evals" || phase == "impl"
	case "Audit":
		return phase == "audit"
	}
	return false
}

func immutableBody(body, phase string) string {
	var lines []string
	keep := true
	for _, line := range strings.Split(body, "'"\\n\") {
		if name, ok := strings.CutPrefix(line, \"## \"); ok {
			name = strings.TrimSpace(name)
			keep = "'!editableSection(name, phase) && name != "Audit" && name != "Provenance"
		}
		if keep && strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, "'"\\n\")
}

func archivePath(changePath string) string {
	return path.Join(archiveDir, path.Base(changePath))
}
EOF
python3 - <<'PY'
from pathlib import Path
p = Path('packages/workflow/internal/lifecycle/claim.go')
s = p.read_text().replace('\"strings\"\\n', '\"strings\"\\n\\t\"time\"\\n', 1)
start = s.index('\\tfor _, front := range []map[string]any{a, b}')
end = s.index('\\tauthorFront :=', start)
p.write_text(s[:start] + s[end:])
PY" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 597ms:
exec
/bin/zsh -lc "cat > packages/workflow/internal/lifecycle/verify.go <<'EOF'
package lifecycle

import (
	\"context\"
	\"errors\"
	\"fmt\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

type verification struct {
	repo gitstore.Repository
	claim claim
	working project.Change
	audit auditSubmission
}

func Verify(ctx context.Context, request model.Request) (model.Result, error) {
	result, _, err := verify(ctx, request)
	return result, err
}

func verify(ctx context.Context, request model.Request) (result model.Result, step verification, err error) {
	step.repo, err = repository(request.Root)
	if err "'!= nil {
		return result, step, err
	}
	root := step.repo.Root
	var oid string
	step.claim, oid, err = heldClaim(step.repo, request.Change)
	if err != nil {
		return result, step, err
	}
	result.State, err = step.claim.state(step.repo, oid)
	if err != nil {
		return result, step, err
	}
	state := result.State
	switch state.Phase {
	case "evals", "audit", "impl":
	default:
		return result, step, coded("P01", state.Change.Path, "unsupported phase "+state.Phase)
	}
	if state.Head != state.PhaseBase {
		return result, step, coded("P05", state.Change.Path, fmt.Sprintf("stale phase base: HEAD is %s, expected %s", state.Head, state.PhaseBase))
	}
	step.working, err = approvedWorking(state)
	if err != nil {
		return result, step, err
	}
	if err = checkStale(root, step.working); err != nil {
		return result, step, err
	}
	result.State.Change = step.working
	m, err := moduleNamed(state.Snapshot, state.Change.Module)
	if err != nil {
		return result, step, err
	}
	changes, err := gitstore.Changes(root, state.PhaseBase)
	if err != nil {
		return result, step, err
	}
	// An index entry may differ even when its working copy matches the base.
	index, err := diffPaths(ctx, root, "diff", "--cached", "--name-status", "-z", "--no-renames", state.PhaseBase, "--")
	if err != nil {
		return result, step, err
	}
	known := map[string]bool{}
	for _, change := range changes {
		known[change.Path] = true
		known[change.OldPath] = true
	}
	for _, change := range index {
		if !known[change.Path] {
			changes = append(changes, change)
		}
	}
	evalFix := stringList(state.Change.Document.Frontmatter["eval_fix"])
	if len(evalFix) != 0 {
		for _, change := range changes {
			for _, file := range []string{change.Path, change.OldPath} {
				if file == "" || file == state.Change.Path || state.Phase == "evals" && belongsTo(state.Snapshot, m, file) && zone(state.Snapshot, m, file) == "evals" {
					continue
				}
				result.Diagnostics = append(result.Diagnostics, project.Diagnostic{
					Code: "Z01", Path: file, Message: fmt.Sprintf("%s is forbidden during eval_fix %s", file, state.Phase),
				})
			}
		}
	}
	result.Diagnostics = append(result.Diagnostics, checks.Zones(checks.ZoneRequest{
		Snapshot: state.Snapshot, Module: m.Name, Phase: state.Phase,
		ChangePath: state.Change.Path, Base: state.PhaseBase, Changes: changes,
		InterfaceDeclared: strings.TrimSpace(section(state.Change.Document, "Interface")) != "",
	})...)
	if err = firstError(result.Diagnostics); err != nil {
		return result, step, err
	}

	items, err := specs(m)
	if err != nil {
		return result, step, err
	}
	if state.Phase == "audit" {
		step.audit, err = checkAudit(state, step.working, items)
		return result, step, err
	}
	if err = validateRun(state.Run); err != nil {
		return result, step, err
	}
	runDir := filepath.Join(root, ".fireweed", "runs", state.Run, state.Phase)
	result.Reports = make(map[string]checks.Report)
	defer func() {
		if len(result.Costs) == 0 {
			return
		}
		diagnostics := costWarnings(state.Snapshot, result.Costs)
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		snapshot := newCostSnapshot(state, result.Costs, diagnostics)
		err = combine(err, writeJSON(filepath.Join(runDir, "costs.json"), snapshot))
	}()
	fullImpl := state.Phase == "impl" && !request.Quick
	modules := []project.Module{m}
	if fullImpl {
		mode := stringValue(state.Snapshot.Config["verify_dependents"])
		if mode == "" {
			mode = "direct"
		}
		for _, dependent := range state.Snapshot.Modules {
			if dependent.Name != m.Name && mode != "none" && dependsOn(state.Snapshot, dependent.Name, m.Name, mode == "all") {
				modules = append(modules, dependent)
			}
		}
	}

	var verificationErr error
	for i, module := range modules {
		moduleItems := items
		if i != 0 {
			moduleItems, err = specs(module)
			if err != nil {
				return result, step, combine(verificationErr, err)
			}
		}
		report, evalErr := checks.Evaluate(ctx, checks.HookRequest{
			Snapshot: state.Snapshot, Module: module.Name, Bin: request.Bin, Timeout: request.Timeout,
		})
		if evalErr != nil {
			return result, step, combine(verificationErr, evalErr)
		}
		result.Reports[module.Name] = report
		cost, costErr := moduleCost(state.Snapshot, module, report)
		if costErr != nil {
			return result, step, combine(verificationErr, costErr)
		}
		result.Costs = append(result.Costs, cost)
		if err = writeJSON(filepath.Join(runDir, module.Name, "report.json"), report); err != nil {
			return result, step, combine(verificationErr, err)
		}
		var reportErr error
		if i == 0 {
			reportErr = checks.RequireCoverage(report, enforcedSpecs(moduleItems))
		}
		if reportErr == nil {
			switch {
			case state.Phase == "evals" && len(evalFix) != 0:
				for _, c := range report.Cases {
					if c.Result == "fail" || c.Result == "error" {
						reportErr = coded("E01", module.Name, fmt.Sprintf("failed case %s: %s", c.Name, c.Failure))
						break
					}
				}
				if reportErr == nil {
					listed := make([]contracts.Item, 0, len(evalFix))
					for _, id := range evalFix {
						listed = append(listed, contracts.Item{ID: id})
					}
					reportErr = checks.RequireCoverage(report, listed)
				}
			case state.Phase == "evals":
				reportErr = checks.RequireFailFirst(report, changeIDs(state.Change, moduleItems, false), stringList(state.Change.Document.Frontmatter["already_satisfied"]))
			case i == 0:
				reportErr = checks.RequirePass(report, nonQuarantinedSpecs(moduleItems))
			default:
				reportErr = checks.RequirePass(report, enforcedSpecs(moduleItems))
			}
		}
		verificationErr = combine(verificationErr, reportErr)
	}
	if fullImpl {
		diagnostics, boundaryErr := checks.Boundaries(ctx, checks.BoundaryRequest{
			Snapshot: state.Snapshot, Bin: request.Bin, Timeout: request.Timeout,
		})
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		if boundaryErr == nil {
			boundaryErr = firstError(diagnostics)
		} else {
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && diagnostic.Code != "" {
					boundaryErr = withCode(diagnostic.Code, diagnostic.Path, boundaryErr)
					break
				}
			}
		}
		verificationErr = combine(verificationErr, boundaryErr)
	}
	return result, step, verificationErr
}

func combine(first, second error) error {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}
	return errors.Join(first, second)
}

func quarantineReason(item contracts.Item) (string, bool) {
	tag := strings.TrimSpace(item.Tag)
	if tag == "quarantine" {
		return "", true
	}
	reason, ok := strings.CutPrefix(tag, "quarantine:")
	return strings.TrimSpace(reason), ok
}

func nonQuarantinedSpecs(items []contracts.Item) []contracts.Item {
	var result []contracts.Item
	for _, item := range items {
		if _, quarantined := quarantineReason(item); !quarantined {
			result = append(result, item)
		}
	}
	return result
}

func enforcedSpecs(items []contracts.Item) []contracts.Item {
	var result []contracts.Item
	for _, item := range items {
		if _, quarantined := quarantineReason(item); !quarantined && item.Tag != "review" {
			result = append(result, item)
		}
	}
	return result
}

func changeIDs(ch project.Change, items []contracts.Item, audit bool) []string {
	enforced := map[string]bool{}
	for _, item := range enforcedSpecs(items) {
		enforced[item.ID] = true
	}
	selected := map[string]bool{}
	for _, edits := range [][]contracts.Item{ch.Added, ch.Modified} {
		for _, item := range edits {
			if enforced[item.ID] {
				selected[item.ID] = true
			}
		}
	}
	for _, id := range stringList(ch.Document.Frontmatter["fixes"]) {
		if enforced[id] {
			selected[id] = true
		}
	}
	if audit {
		for _, id := range stringList(ch.Document.Frontmatter["eval_fix"]) {
			if enforced[id] {
				selected[id] = true
			}
		}
	}
	return sortedKeys(selected)
}
EOF
cat > packages/workflow/internal/lifecycle/audit.go <<'"'EOF'
package lifecycle

import (
	\"fmt\"
	\"regexp\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

type auditFinding struct {
	id string
	clause string
	line string
	untestable bool
}

type auditSubmission struct {
	findings []auditFinding
	checked []string
}

var findingLine = regexp.MustCompile("'`''^[-*]'"\\s+(?:(\\[untestable\\])\\s+)?(["'^'"\\s:]+):\\s*\"(.+?)\"\\s*(.*)"'$`)

func parseFinding(line string) (auditFinding, bool) {
	parts := findingLine.FindStringSubmatch(line)
	if parts == nil {
		return auditFinding{}, false
	}
	return auditFinding{id: parts[2], clause: parts[3], line: line, untestable: parts[1] != ""}, true
}

func checkAudit(state model.State, ch project.Change, items []contracts.Item) (auditSubmission, error) {
	var result auditSubmission
	specBodies := map[string]string{}
	for _, item := range items {
		specBodies[item.ID] = item.Body
	}
	priorIDs := map[string]bool{}
	reaudit := intValue(ch.Document.Frontmatter["audit_rounds"], 0) > 0
	if reaudit {
		for _, line := range strings.Split(section(state.Change.Document, "Audit"), "'"\\n\") {
			if finding, ok := parseFinding(strings.TrimSpace(line)); ok {
				priorIDs[finding.id] = true
			}
		}
	}
	checked, noFindings := false, false
	ids := map[string]bool{}
	addChecked := func(value string) {
		checked = true
		for _, id := range strings.Fields(strings.ReplaceAll(value, \",\", \" \")) {
			ids[strings.Trim(id, \""'`")] = true
		}
	}
	for _, line := range strings.Split(section(ch.Document, "Audit"), "'"\\n\") {
		line = strings.TrimSpace(line)
		if line == \"\" {
			continue
		}
		switch {
		case line == \"No findings\":
			noFindings = true
		case strings.HasPrefix(line, \"Checked IDs:\"):
			addChecked(strings.TrimPrefix(line, \"Checked IDs:\"))
		case strings.HasPrefix(line, \"No findings; checked\"):
			noFindings = true
			addChecked(strings.TrimPrefix(line, \"No findings; checked\"))
		default:
			finding, ok := parseFinding(line)
			if "'!ok {
				return result, coded("P09", ch.Path, fmt.Sprintf("malformed audit submission line %q", line))
			}
			if body, ok := specBodies[finding.id]; ok && strings.Contains(body, finding.clause) && (!reaudit || priorIDs[finding.id]) {
				result.findings = append(result.findings, finding)
			}
		}
	}
	if len(result.findings) != 0 {
		return result, nil
	}
	if !checked || !noFindings {
		return result, coded("P09", ch.Path, "audit requires Checked IDs and No findings")
	}
	for _, id := range changeIDs(ch, items, true) {
		if !ids[id] {
			return result, coded("P09", ch.Path, "audit did not check "+id)
		}
	}
	result.checked = sortedKeys(ids)
	return result, nil
}

func (a auditSubmission) body() string {
	if len(a.findings) == 0 {
		return "'"\\nNo findings; checked \"+strings.Join(a.checked, \", \")+\"\\n\"
	}
	var lines []string
	for _, f := range a.findings {
		lines = append(lines, f.line)
	}
	return \"\\n\"+strings.Join(lines, \"\\n\")+\"\\n\"
}

func (a auditSubmission) blocked(rounds, max int) string {
	for _, f := range a.findings {
		if f.untestable {
			return \"audit retained an untestable clause\"
		}
	}
	if len(a.findings) "'!= 0 && rounds >= max {
		return fmt.Sprintf("audit findings exceed the limit of %d rounds", max)
	}
	return ""
}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 532ms:
codex
The claim, queue, and verification code is in place. I’m now connecting phase transitions, saved attempts, and reporting, with Git ref updates designed to preserve claim ownership.
exec
/bin/zsh -lc "cat > packages/workflow/internal/lifecycle/git.go <<'EOF'
package lifecycle

import (
	\"context\"
	\"fmt\"
	\"os\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/gitstore\"
)

type gitRef struct {
	name string
	oid string
}

func listRefs(ctx context.Context, root string, prefixes ...string) ([]gitRef, error) {
	args := append([]string{\"for-each-ref\", \"--format=%(refname) %(objectname)\"}, prefixes...)
	out, err := git(ctx, root, nil, nil, args...)
	if err "'!= nil {
		return nil, err
	}
	var refs []gitRef
	for _, line := range strings.Split(string(out), "'"\\n\") {
		if name, oid, ok := strings.Cut(line, \" \"); ok {
			refs = append(refs, gitRef{name, oid})
		}
	}
	return refs, nil
}

func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
	out, err := git(ctx, root, nil, nil, args...)
	if err "'!= nil {
		return nil, err
	}
	parts := splitNUL(out)
	var changes []gitstore.Change
	for len(parts) > 0 {
		if len(parts) < 2 || parts[0] == "" {
			return nil, fmt.Errorf("invalid git name-status output")
		}
		change := gitstore.Change{Status: parts[0], Path: parts[1]}
		parts = parts[2:]
		if change.Status[0] == '"'R' || change.Status[0] == 'C' {
			if len(parts) == 0 {
				return nil, fmt.Errorf(\"missing destination in git name-status output\")
			}
			change.OldPath, change.Path = change.Path, parts[0]
			parts = parts[1:]
		}
		changes = append(changes, change)
	}
	return changes, nil
}

type gitIndex struct {
	root string
	env []string
	ctx context.Context
}

func temporaryIndex(ctx context.Context, root, base string) (gitIndex, func(), error) {
	dir, err := os.MkdirTemp(\"\", \"fireweed-index-\")
	if err "'!= nil {
		return gitIndex{}, nil, err
	}
	clean := func() { os.RemoveAll(dir) }
	idx := gitIndex{root: root, env: []string{"GIT_INDEX_FILE="+filepath.Join(dir, "index")}, ctx: ctx}
	if _, err := idx.run(nil, "read-tree", base); err != nil {
		clean()
		return gitIndex{}, nil, err
	}
	return idx, clean, nil
}

func (idx gitIndex) run(input []byte, args ...string) ([]byte, error) {
	return git(idx.ctx, idx.root, idx.env, input, args...)
}

func (idx gitIndex) stageWorktree() error {
	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
	return err
}

func (idx gitIndex) stageDocument(name string, data []byte) error {
	oid, err := gitstore.WriteBlob(idx.root, data)
	if err != nil {
		return err
	}
	_, err = idx.run(nil, "update-index", "--add", "--cacheinfo", "100644", oid, name)
	return err
}

func (idx gitIndex) commit(parent, message string) (string, error) {
	out, err := idx.run(nil, "write-tree")
	if err != nil {
		return "", err
	}
	tree := strings.TrimSpace(string(out))
	env := append([]string(nil), idx.env...)
	for _, identity := range []struct{key, value, author, committer string}{
		{"user.name", "Fireweed", "GIT_AUTHOR_NAME", "GIT_COMMITTER_NAME"},
		{"user.email", "fireweed@localhost", "GIT_AUTHOR_EMAIL", "GIT_COMMITTER_EMAIL"},
	} {
		if value, _ := gitOutput(idx.root, "config", "--get", identity.key); value == "" {
			for _, key := range []string{identity.author, identity.committer} {
				if os.Getenv(key) == "" {
					env = append(env, key+"="+identity.value)
				}
			}
		}
	}
	out, err = git(idx.ctx, idx.root, env, []byte(message+"'"\\n\"), \"commit-tree\", tree, \"-p\", parent)
	return strings.TrimSpace(string(out)), err
}

func (idx gitIndex) checkout(revision string) error {
	_, err := idx.run(nil, \"read-tree\", \"--reset\", \"-u\", revision)
	return err
}

func replaceIndex(root, revision string) error {
	_, err := gitOutput(root, \"read-tree\", \"--reset\", revision)
	return err
}

func commitRefs(root, oldHead, newHead, extra string) error {
	commands := fmt.Sprintf(\"update HEAD %s %s\\n%s\", newHead, oldHead, extra)
	_, err := git(context.Background(), root, nil, []byte(commands), \"update-ref\", \"--stdin\")
	return err
}

func claimUpdate(ref, oldOID, newOID string) string {
	if newOID == \"\" {
		return fmt.Sprintf(\"delete %s %s\\n\", ref, oldOID)
	}
	return fmt.Sprintf(\"update %s %s %s\\n\", ref, newOID, oldOID)
}

func transitionError(root, base, name string, err error) error {
	if head, inspectErr := gitOutput(root, \"rev-parse\", \"HEAD\"); inspectErr == nil && head "'!= base {
		return withCode("P05", name, fmt.Errorf("stale phase base: %w", err))
	}
	return withCode("P02", name, err)
}
EOF
cat > packages/workflow/internal/lifecycle/done.go <<'"'EOF'
package lifecycle

import (
	\"context\"
	\"crypto/sha256\"
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path\"
	\"path/filepath\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/assets\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

func Done(ctx context.Context, request model.Request) (result model.Result, err error) {
	request.Quick = false
	result, step, err := verify(ctx, request)
	if err "'!= nil {
		return result, err
	}
	state := result.State
	root := step.repo.Root
	before, err := step.claim.state(step.repo, state.ClaimOID)
	if err != nil {
		return result, err
	}
	ch, err := approvedWorking(before)
	if err != nil {
		return result, err
	}
	if state.Phase == "audit" {
		ch = step.working
	}
	next, blocked := "impl", ""
	switch state.Phase {
	case "evals":
		if configGroup(state.Snapshot.Config, "audit")["enabled"] != false {
			next = "audit"
		}
	case "audit":
		ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Audit", step.audit.body())
		rounds := intValue(ch.Document.Frontmatter["audit_rounds"], 0)
		blocked = step.audit.blocked(rounds, intValue(configGroup(state.Snapshot.Config, "audit")["max_rounds"], 2))
		if blocked != "" {
			next = "audit"
			ch.Document.Frontmatter["blocked"] = blocked
			questions := strings.TrimRight(section(ch.Document, "Questions"), "'"\\n\")
			ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, \"Questions\", questions+\"\\n\"+step.audit.body())
		} else if len(step.audit.findings) "'!= 0 {
			next = "evals"
			ch.Document.Frontmatter["audit_rounds"] = rounds+1
		}
	case "impl":
		next = "applied"
		// Check the checkout again after hooks, before replacing any manifest.
		diagnostics := project.Check(root)
		if staticErr := firstError(diagnostics); staticErr != nil {
			result.Diagnostics = append(result.Diagnostics, diagnostics...)
			return result, staleError(staticErr)
		}
	}
	ch.Document.Frontmatter["phase"] = next
	playbook, err := assets.Playbook(playbookName(state))
	if err != nil {
		return result, err
	}
	agent := strings.TrimSpace(request.Agent)
	if agent == "" {
		agent = "manual"
	}
	agent = strings.NewReplacer("'"\\r\", \" \", \"\\n\", \" \").Replace(agent)
	tests, failing, passing := 0, 0, 0
	for _, report := range result.Reports {
		for _, c := range report.Cases {
			tests++
			switch c.Result {
			case \"fail\", \"error\":
				failing++
			case \"pass\":
				passing++
			}
		}
	}
	line := fmt.Sprintf(\"- %s %s agent=%s fireweed=%s playbook=%x tests=%d failing=%d passing=%d\",
		state.Phase, time.Now().UTC().Format(time.RFC3339Nano), agent, assets.Version(),
		sha256.Sum256([]byte(playbook)), tests, failing, passing)
	ch.Document.Body = appendProvenance(ch.Document.Body, line)
	data, err := contracts.RenderDocument(ch.Document)
	if err "'!= nil {
		return result, err
	}

	files := fileEdits{root: root}
	published := false
	defer func() {
		if err != nil && !published {
			err = combine(err, files.rollback())
		}
	}()
	if next == "applied" {
		if err = applyToModule(root, ch, &files); err != nil {
			return result, err
		}
		pending := ch.Path
		ch.Path = archivePath(ch.Path)
		if _, err = os.Lstat(filepath.Join(root, ch.Path)); err == nil {
			return result, fmt.Errorf("archive already exists: %s", ch.Path)
		} else if !errors.Is(err, fs.ErrNotExist) {
			return result, err
		}
		if err = files.remove(pending); err != nil {
			return result, err
		}
	}
	if err = files.write(ch.Path, []byte(data)); err != nil {
		return result, err
	}
	if next == "applied" {
		diagnostics := project.Check(root)
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		if err = firstError(diagnostics); err != nil {
			return result, staleError(err)
		}
	}
	idx, cleanup, err := temporaryIndex(ctx, root, state.PhaseBase)
	if err != nil {
		return result, err
	}
	defer cleanup()
	if err = idx.stageWorktree(); err != nil {
		return result, err
	}
	newHead, err := idx.commit(state.PhaseBase, "fireweed: "+state.Change.Slug+" "+state.Phase)
	if err != nil {
		return result, err
	}
	var newClaim string
	if next != "applied" && blocked == "" {
		updated := step.claim
		updated.Base, updated.Approved = newHead, data
		newClaim, err = updated.blob(root)
		if err != nil {
			return result, err
		}
	}
	if err = commitRefs(root, state.PhaseBase, newHead, claimUpdate(step.claim.ref(), state.ClaimOID, newClaim)); err != nil {
		return result, transitionError(root, state.PhaseBase, ch.Path, err)
	}
	published = true
	result.State.Change, result.State.Phase, result.State.ClaimOID = ch, next, newClaim
	result.State.Head, result.State.PhaseBase = newHead, newHead
	if err = replaceIndex(root, newHead); err != nil {
		return result, err
	}
	if blocked != "" {
		return result, coded("P06", ch.Path, "change blocked: "+blocked)
	}
	return result, nil
}

func playbookName(s model.State) string {
	if s.Phase == "impl" && s.Change.Document.Frontmatter["regenerate"] == true {
		return "regenerate"
	}
	return s.Phase
}

func appendProvenance(body, line string) string {
	prior, _ := contracts.Section(body, "Provenance")
	var lines []string
	for _, entry := range strings.Split(prior, "'"\\n\") {
		if strings.TrimSpace(entry) "'!= "" {
			lines = append(lines, entry)
		}
	}
	return contracts.ReplaceSection(body, "Provenance", "'"\\n\"+strings.Join(append(lines, line), \"\\n\")+\"\\n\")
}

func applyToModule(root string, ch project.Change, files *fileEdits) error {
	snapshot, err := project.Load(root)
	if err "'!= nil {
		return err
	}
	m, err := moduleNamed(snapshot, ch.Module)
	if err != nil {
		return err
	}
	doc := m.Document
	changed := false
	for key, value := range configGroup(ch.Document.Frontmatter, "set") {
		doc.Frontmatter[key] = value
		changed = true
	}
	if len(ch.Added)+len(ch.Modified)+len(ch.Removed) != 0 {
		items, err := specs(m)
		if err != nil {
			return err
		}
		doc.Body = contracts.ReplaceSection(doc.Body, "Specs", renderSpecEdits(items, ch))
		changed = true
	}
	added := strings.TrimSpace(section(ch.Document, "Decisions"))
	removed := strings.TrimSpace(section(ch.Document, "Removed decisions"))
	if added != "" || removed != "" {
		doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", applyDecisions(section(doc, "Decisions"), added, removed, ch.Slug))
		changed = true
	}
	if !changed {
		return nil
	}
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		return err
	}
	return files.write(path.Join(m.Dir, "SPEC.md"), []byte(rendered))
}

func decisionText(line string) string {
	line = strings.TrimSpace(line)
	line = strings.TrimPrefix(line, "- ")
	line = strings.TrimPrefix(line, "> ")
	line = strings.Trim(line, "'"\\\"\")
	if strings.HasPrefix(line, \"[\") {
		if _, rest, ok := strings.Cut(line, \"] \"); ok {
			return strings.TrimSpace(rest)
		}
	}
	return line
}

func applyDecisions(prior, added, removed, slug string) string {
	removals := map[string]bool{}
	for _, line := range strings.Split(removed, \"\\n\") {
		if text := decisionText(line); text "'!= "" {
			removals[text] = true
		}
	}
	var kept []string
	for _, line := range strings.Split(strings.Trim(prior, "'"\\n\"), \"\\n\") {
		text := decisionText(line)
		_, unattributed, attributed := strings.Cut(text, \": \")
		if "'!removals[text] && !(strings.HasPrefix(strings.TrimSpace(line), "- ") && attributed && removals[unattributed]) {
			kept = append(kept, line)
		}
	}
	for _, line := range strings.Split(added, "'"\\n\") {
		if text := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), \"- \")); text "'!= "" {
			kept = append(kept, "- "+slug+": "+text)
		}
	}
	return "'"\\n\"+strings.Trim(strings.Join(kept, \"\\n\"), \"\\n\")+\"\\n\"
}

type fileEdits struct {
	root string
	undo []func() error
}

func (files *fileEdits) replace(name string, data []byte, remove bool) error {
	file := filepath.Join(files.root, name)
	before, err := os.ReadFile(file)
	exists := err == nil
	if err "'!= nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if remove {
		err = os.Remove(file)
	} else {
		err = writeFile(file, data)
	}
	if err == nil {
		files.undo = append(files.undo, func() error {
			if exists {
				return writeFile(file, before)
			}
			return os.Remove(file)
		})
	}
	return err
}

func (files *fileEdits) write(name string, data []byte) error { return files.replace(name, data, false) }
func (files *fileEdits) remove(name string) error { return files.replace(name, nil, true) }

func (files *fileEdits) rollback() error {
	var err error
	for i := len(files.undo)-1; i >= 0; i-- {
		err = combine(err, files.undo[i]())
	}
	return err
}
EOF
cat > packages/workflow/internal/lifecycle/attempt.go <<'"'EOF'
package lifecycle

import (
	\"context\"
	\"fmt\"
	\"strconv\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

func Block(request model.BlockRequest) (result model.BlockResult, err error) {
	repo, err := repository(request.Root)
	if err "'!= nil {
		return result, err
	}
	root := repo.Root
	c, oid, err := heldClaim(repo, request.Change)
	if err != nil {
		return result, err
	}
	state, err := c.state(repo, oid)
	if err != nil {
		return result, err
	}
	if state.Head != state.PhaseBase {
		return result, coded("P05", c.Path, fmt.Sprintf("stale phase base: HEAD is %s, expected %s", state.Head, state.PhaseBase))
	}
	if _, err = approvedWorking(state); err != nil {
		return result, err
	}
	if strings.TrimSpace(request.Reason) == "" {
		return result, coded("P01", c.Path, "block requires a reason")
	}
	ctx := context.Background()
	attempt, cleanAttempt, err := temporaryIndex(ctx, root, state.PhaseBase)
	if err != nil {
		return result, err
	}
	defer cleanAttempt()
	if err = attempt.stageWorktree(); err != nil {
		return result, err
	}
	attemptOID, err := attempt.commit(state.PhaseBase, "fireweed: save "+c.Change+" "+state.Phase+" attempt")
	if err != nil {
		return result, err
	}
	prefix := attemptPrefix+c.Change+"/"
	refs, err := listRefs(ctx, root, prefix)
	if err != nil {
		return result, err
	}
	used := map[string]bool{}
	for _, ref := range refs {
		used[ref.name] = true
	}
	n := 1
	for used[prefix+strconv.Itoa(n)] {
		n++
	}
	ref := prefix+strconv.Itoa(n)
	state.Change.Document.Frontmatter["blocked"] = request.Reason
	approved, err := contracts.RenderDocument(state.Change.Document)
	if err != nil {
		return result, err
	}
	blocked, cleanBlocked, err := temporaryIndex(ctx, root, state.PhaseBase)
	if err != nil {
		return result, err
	}
	defer cleanBlocked()
	if err = blocked.stageDocument(c.Path, []byte(approved)); err != nil {
		return result, err
	}
	newHead, err := blocked.commit(state.PhaseBase, "fireweed: block "+c.Change)
	if err != nil {
		return result, err
	}
	// The attempt index describes the working tree, including new files.
	// Switching it to the block tree restores precisely the captured step.
	if err = attempt.checkout(newHead); err != nil {
		return result, err
	}
	updates := claimUpdate(c.ref(), oid, "")+fmt.Sprintf("create %s %s'"\\n\", ref, attemptOID)
	if err = commitRefs(root, state.PhaseBase, newHead, updates); err "'!= nil {
		return result, combine(transitionError(root, state.PhaseBase, c.Path, err), attempt.checkout(attemptOID))
	}
	state.ClaimOID, state.Head, state.PhaseBase = "", newHead, newHead
	result = model.BlockResult{State: state, AttemptRef: ref, AttemptOID: attemptOID}
	return result, replaceIndex(root, newHead)
}

func resumeDocument(root, oid string, current project.Change) (string, error) {
	raw, err := gitstore.ReadAt(root, oid, current.Path)
	if err != nil {
		return "", withCode("P01", current.Path, err)
	}
	saved, err := readChange(current, raw)
	if err != nil {
		return "", withCode("P01", current.Path, err)
	}
	phase := field(current, "phase")
	if field(saved, "phase") != phase || field(saved, "run") != field(current, "run") {
		return "", coded("P01", current.Path, "saved attempt belongs to a different run or phase")
	}
	for _, name := range []string{"Decisions", "Questions", "Audit"} {
		if editableSection(name, phase) {
			if body, ok := contracts.Section(saved.Document.Body, name); ok {
				current.Document.Body = contracts.ReplaceSection(current.Document.Body, name, body)
			}
		}
	}
	return contracts.RenderDocument(current.Document)
}

func restoreAttempt(root, oid, file string) error {
	ctx := context.Background()
	patch, err := git(ctx, root, nil, nil, "diff", "--binary", "--full-index", "--no-renames", oid+"''^", oid, "--", ".", ":(exclude,top,literal)"+file)
	if err != nil {
		return withCode("P01", file, err)
	}
	if len(patch) == 0 {
		return nil
	}
	_, err = git(ctx, root, nil, patch, "apply", "--whitespace=nowarn", "-")
	return withCode("P01", file, err)
}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 528ms:
exec
/bin/zsh -lc "cat > packages/workflow/internal/lifecycle/history.go <<'EOF'
package lifecycle

import (
	\"context\"
	\"fmt\"
	\"os\"
	\"path\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
)

type revision struct {
	oid string
	parent string
}

func revisions(root, rangeExpr string) ([]revision, error) {
	out, err := gitOutput(root, \"rev-list\", \"--date-order\", \"--parents\", rangeExpr, \"--\")
	if err "'!= nil {
		return nil, err
	}
	var result []revision
	for _, line := range strings.Split(out, "'"\\n\") {
		fields := strings.Fields(line)
		if len(fields) "'!= 0 {
			rev := revision{oid: fields[0]}
			if len(fields) > 1 {
				rev.parent = fields[1]
			}
			result = append(result, rev)
		}
	}
	return result, nil
}

func resolveRevision(root, name string) (string, error) {
	return gitOutput(root, "rev-parse", "--verify", "--end-of-options", name+"''^{commit}")
}

func committedChanges(root, rev, dir string) ([]project.Change, error) {
	paths, err := pathsAt(root, rev, dir)
	if err != nil {
		return nil, err
	}
	for i, ch := range paths {
		raw, err := gitstore.ReadAt(root, rev, ch.Path)
		if err != nil {
			return nil, err
		}
		paths[i], err = readChange(ch, raw)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", ch.Path, err)
		}
	}
	return paths, nil
}

func archivedAt(root, rev string) ([]project.Change, error) {
	changes, err := committedChanges(root, rev, archiveDir)
	if err != nil {
		return nil, err
	}
	var applied []project.Change
	for _, ch := range changes {
		if field(ch, "phase") == "applied" {
			applied = append(applied, ch)
		}
	}
	return applied, nil
}

func existsAt(root, rev, name string) bool {
	if rev == "" {
		return false
	}
	_, err := gitOutput(root, "cat-file", "-e", rev+":"+name)
	return err == nil
}

func baseRefs(root string, config map[string]any) (shared, local string, err error) {
	remote := stringValue(config["remote"])
	base := stringValue(config["base"])
	if remote == "" {
		remote = "origin"
	}
	resolved, err := gitstore.BaseRef(root, remote, base)
	if err != nil {
		return "", "", err
	}
	if base == "" {
		remotePrefix := "refs/remotes/"+remote+"/"
		if symbolic, symErr := gitOutput(root, "symbolic-ref", "--quiet", remotePrefix+"HEAD"); symErr == nil {
			base = strings.TrimPrefix(symbolic, remotePrefix)
		} else {
			base = strings.TrimPrefix(resolved, "refs/heads/")
			base = strings.TrimPrefix(base, remotePrefix)
			base = strings.TrimPrefix(base, remote+"/")
		}
	}
	base = strings.TrimPrefix(base, "refs/heads/")
	if oid, resolveErr := resolveRevision(root, "refs/heads/"+base); resolveErr == nil {
		local = oid
	}
	if oid, resolveErr := resolveRevision(root, "refs/remotes/"+remote+"/"+base); resolveErr == nil {
		shared = oid
	} else if local != "" {
		shared = local
	} else {
		shared, err = resolveRevision(root, resolved)
	}
	return shared, local, err
}

func unapprovedCount(root, shared string, archives []project.Change) (int, error) {
	approved := map[string]bool{}
	for _, dir := range []string{pendingDir, archiveDir} {
		changes, err := pathsAt(root, shared, dir)
		if err != nil {
			return 0, err
		}
		for _, ch := range changes {
			approved[path.Base(ch.Path)] = true
		}
	}
	count := 0
	for _, ch := range archives {
		if !approved[path.Base(ch.Path)] {
			count++
		}
	}
	return count, nil
}

func historySelection(root, selector string) (project.Change, error) {
	var changes []project.Change
	for _, dir := range []string{pendingDir, archiveDir} {
		paths, err := changePaths(root, dir)
		if err != nil {
			return project.Change{}, err
		}
		changes = append(changes, paths...)
	}
	return selectChange(changes, selector)
}

func documentAt(root, rev, filename string) (project.Change, error) {
	if rev == "" {
		return project.Change{}, nil
	}
	for _, dir := range []string{archiveDir, pendingDir} {
		name := path.Join(dir, filename)
		if existsAt(root, rev, name) {
			raw, err := gitstore.ReadAt(root, rev, name)
			if err != nil {
				return project.Change{}, err
			}
			return readChange(identity(name), raw)
		}
	}
	return project.Change{}, nil
}

func phaseAt(root, rev, filename string) (string, error) {
	if rev == "" {
		return "", nil
	}
	for _, dir := range []string{archiveDir, pendingDir} {
		name := path.Join(dir, filename)
		if existsAt(root, rev, name) {
			raw, err := gitstore.ReadAt(root, rev, name)
			if err != nil {
				return "", err
			}
			doc, err := contracts.ParseDocument(string(raw))
			if err != nil {
				return "", err
			}
			return stringValue(doc.Frontmatter["phase"]), nil
		}
	}
	return "", nil
}

func revisionDiff(ctx context.Context, root string, rev revision) ([]gitstore.Change, error) {
	args := []string{"diff-tree", "--root", "--no-commit-id", "--name-status", "-z", "-r", "--no-renames"}
	if rev.parent != "" {
		args = append(args, rev.parent)
	}
	return diffPaths(ctx, root, append(args, rev.oid, "--")...)
}

func phaseCommits(root, endpoint, filename string) ([]revision, error) {
	revs, err := revisions(root, endpoint)
	if err != nil {
		return nil, err
	}
	var selected []revision
	phases := map[string]string{"": ""}
	phase := func(oid string) (string, error) {
		if phase, ok := phases[oid]; ok {
			return phase, nil
		}
		phase, err := phaseAt(root, oid, filename)
		phases[oid] = phase
		return phase, err
	}
	for _, rev := range revs {
		next, err := phase(rev.oid)
		if err != nil {
			return nil, err
		}
		if next == "" {
			continue
		}
		previous, err := phase(rev.parent)
		if err != nil {
			return nil, err
		}
		if next != previous {
			selected = append(selected, rev)
		}
	}
	return selected, nil
}

func Commits(root, selector string) ([]string, error) {
	repo, err := repository(root)
	if err != nil {
		return nil, err
	}
	ch, err := historySelection(repo.Root, selector)
	if err != nil {
		return nil, err
	}
	revs, err := phaseCommits(repo.Root, repo.Head, path.Base(ch.Path))
	if err != nil {
		return nil, err
	}
	result := make([]string, 0, len(revs))
	for _, rev := range revs {
		result = append(result, rev.oid)
	}
	return result, nil
}

func Abandon(root, selector string) (string, error) {
	repo, err := repository(root)
	if err != nil {
		return "", err
	}
	root = repo.Root
	ch, err := historySelection(root, selector)
	if err != nil {
		return "", err
	}
	if err := requireClean(root, "abandon"); err != nil {
		return "", err
	}
	if path.Dir(ch.Path) == archiveDir {
		return "", coded("P01", ch.Path, "cannot abandon an archived change")
	}
	if raw, err := os.ReadFile(filepath.Join(root, ch.Path)); err != nil {
		return "", err
	} else if ch, err = readChange(ch, raw); err != nil {
		return "", err
	} else if field(ch, "phase") == "applied" {
		return "", coded("P01", ch.Path, "cannot abandon an applied change")
	}

	claimRef := claimPrefix+ch.Slug
	oid, exists, err := gitstore.Ref(root, claimRef)
	if err != nil {
		return "", err
	}
	if exists {
		c, err := readClaim(root, oid)
		if err != nil {
			return "", withCode("P02", ch.Path, err)
		}
		if c.Worktree != root || c.Path != ch.Path {
			return "", coded("P02", ch.Path, "claim belongs to another worktree or change")
		}
	}
	revs, err := phaseCommits(root, repo.Head, path.Base(ch.Path))
	if err != nil {
		return "", err
	}
	ctx := context.Background()
	idx, cleanup, err := temporaryIndex(ctx, root, repo.Head)
	if err != nil {
		return "", err
	}
	defer cleanup()
	for _, rev := range revs {
		args := []string{"diff-tree", "--root", "--no-commit-id", "-r", "-p", "--binary", "--full-index", "--no-renames"}
		if rev.parent != "" {
			args = append(args, rev.parent)
		}
		args = append(args, rev.oid, "--", ".", ":(exclude,top,literal)"+ch.Path, ":(exclude,top,literal)"+archivePath(ch.Path))
		patch, err := git(ctx, root, nil, nil, args...)
		if err != nil {
			return "", err
		}
		if len(patch) != 0 {
			if _, err := idx.run(patch, "apply", "--cached", "--3way", "--reverse", "--whitespace=nowarn", "-"); err != nil {
				return "", fmt.Errorf("reverse phase commit %s: %w", rev.oid, err)
			}
		}
	}
	if _, err := idx.run(nil, "update-index", "--force-remove", "--", ch.Path); err != nil {
		return "", err
	}
	newHead, err := idx.commit(repo.Head, "fireweed: abandon "+ch.Slug)
	if err != nil {
		return "", err
	}
	checkout, cleanCheckout, err := temporaryIndex(ctx, root, repo.Head)
	if err != nil {
		return "", err
	}
	defer cleanCheckout()
	if err := checkout.checkout(newHead); err != nil {
		return "", err
	}
	// An unclaimed change remains unclaimed through publication.
	claimCommand := "verify "+claimRef+"'"\\n\"
	if exists {
		claimCommand = claimUpdate(claimRef, oid, \"\")
	}
	if err := commitRefs(root, repo.Head, newHead, claimCommand); err "'!= nil {
		return "", combine(transitionError(root, repo.Head, ch.Path, err), checkout.checkout(repo.Head))
	}
	return newHead, replaceIndex(root, newHead)
}
EOF
cat > packages/workflow/internal/lifecycle/costs.go <<'"'EOF'
package lifecycle

import (
	\"bytes\"
	\"encoding/json\"
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"math\"
	\"os\"
	\"path/filepath\"
	\"sort\"
	\"strconv\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/checks\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

func moduleCost(snapshot project.Snapshot, m project.Module, report checks.Report) (model.ModuleCost, error) {
	cost := model.ModuleCost{Module: m.Name, CasesPerSpec: map[string]int{}}
	for _, c := range report.Cases {
		cost.Seconds += c.Seconds
		seen := map[string]bool{}
		for _, id := range c.IDs {
			if "'!seen[id] {
				cost.CasesPerSpec[id]++
				seen[id] = true
			}
		}
	}
	err := moduleFiles(snapshot, m, func(name string, entry fs.DirEntry) error {
		if !entry.Type().IsRegular() {
			return nil
		}
		switch zone(snapshot, m, name) {
		case "evals", "disposable":
			data, err := os.ReadFile(filepath.Join(snapshot.Root, name))
			if err != nil {
				return err
			}
			lines := bytes.Count(data, []byte{'"'\\n'})
			if len(data) > 0 && data[len(data)-1] "'!= '"'\\n' {
				lines++
			}
			if zone(snapshot, m, name) == \"evals\" {
				cost.EvalLines += lines
			} else {
				cost.DisposableLines += lines
			}
		}
		return nil
	})
	return cost, err
}

func newCostSnapshot(state model.State, modules []model.ModuleCost, ds []project.Diagnostic) model.CostSnapshot {
	return model.CostSnapshot{Run: state.Run, Phase: state.Phase, RecordedAt: time.Now().UTC(), Modules: modules, Diagnostics: ds}
}

func costWarnings(snapshot project.Snapshot, modules []model.ModuleCost) []project.Diagnostic {
	budgets := configGroup(snapshot.Config, \"budgets\")
	timeLimit := 15*time.Minute
	if value, err := time.ParseDuration(stringValue(budgets[\"verify_time\"])); err == nil {
		timeLimit = value
	}
	caseLimit := intValue(budgets[\"cases_per_spec\"], 15)
	ratioLimit := float64(4)
	if value, err := strconv.ParseFloat(fmt.Sprint(budgets[\"eval_ratio\"]), 64); err == nil {
		ratioLimit = value
	}
	var ds []project.Diagnostic
	warn := func(code, path, message string) {
		ds = append(ds, project.Diagnostic{Code: code, Path: path, Warning: true, Message: message})
	}
	var seconds float64
	for _, m := range modules {
		seconds += m.Seconds
	}
	if seconds > timeLimit.Seconds() {
		warn(\"B01\", \"\", fmt.Sprintf(\"evaluated modules took %gs, exceeding the %gs verify_time limit\", seconds, timeLimit.Seconds()))
	}
	for _, m := range modules {
		for _, id := range sortedKeys(m.CasesPerSpec) {
			if count := m.CasesPerSpec[id]; count > caseLimit {
				warn(\"B02\", m.Module, fmt.Sprintf(\"%s has %d testcases, exceeding the limit of %d\", id, count, caseLimit))
			}
		}
		ratio := float64(0)
		if m.DisposableLines > 0 {
			ratio = float64(m.EvalLines)/float64(m.DisposableLines)
		} else if m.EvalLines > 0 {
			ratio = math.Inf(1)
		}
		if ratio > ratioLimit {
			warn(\"B03\", m.Module, fmt.Sprintf(\"%s eval/disposable line ratio %g (%d/%d) exceeds the limit of %g\", m.Module, ratio, m.EvalLines, m.DisposableLines, ratioLimit))
		}
	}
	return ds
}

func Costs(root string) ([]model.CostSnapshot, error) {
	repo, err := repository(root)
	if err "'!= nil {
		return nil, err
	}
	files, err := filepath.Glob(filepath.Join(repo.Root, ".fireweed", "runs", "*", "*", "costs.json"))
	if err != nil {
		return nil, err
	}
	result := make([]model.CostSnapshot, 0, len(files))
	for _, file := range files {
		data, err := os.ReadFile(file)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		var snapshot model.CostSnapshot
		if err := json.Unmarshal(data, &snapshot); err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		result = append(result, snapshot)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].RecordedAt.Before(result[j].RecordedAt)
	})
	return result, nil
}

func latestCosts(root string, modules []project.Module) ([]model.ModuleCost, []project.Diagnostic, error) {
	snapshots, err := Costs(root)
	if err != nil {
		return nil, nil, err
	}
	known := map[string]project.Module{}
	for _, m := range modules {
		known[m.Name] = m
	}
	costs := map[string]model.ModuleCost{}
	selected := map[string]int{}
	for i, snapshot := range snapshots {
		for _, cost := range snapshot.Modules {
			if _, ok := known[cost.Module]; ok {
				costs[cost.Module] = cost
				selected[cost.Module] = i
			}
		}
	}
	var values []model.ModuleCost
	for _, module := range sortedKeys(costs) {
		values = append(values, costs[module])
	}
	var ds []project.Diagnostic
	seen := map[project.Diagnostic]bool{}
	for i, snapshot := range snapshots {
		for _, diagnostic := range snapshot.Diagnostics {
			if seen[diagnostic] || diagnostic.Code != "B01" && diagnostic.Code != "B02" && diagnostic.Code != "B03" {
				continue
			}
			for _, cost := range snapshot.Modules {
				index, ok := selected[cost.Module]
				if !ok || index != i || !costDiagnosticMatches(diagnostic, cost, known[cost.Module]) {
					continue
				}
				ds = append(ds, diagnostic)
				seen[diagnostic] = true
				break
			}
		}
	}
	return values, ds, nil
}

func costDiagnosticMatches(d project.Diagnostic, cost model.ModuleCost, module project.Module) bool {
	switch d.Code {
	case "B01":
		return true
	case "B02", "B03":
		if d.Path != "" {
			return d.Path == cost.Module || d.Path == module.Dir
		}
		if d.Code == "B03" {
			return strings.Contains(d.Message, cost.Module)
		}
		for id := range cost.CasesPerSpec {
			if strings.Contains(d.Message, id) {
				return true
			}
		}
	}
	return false
}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 528ms:
exec
/bin/zsh -lc "cat > packages/workflow/internal/lifecycle/status.go <<'EOF'
package lifecycle

import (
	\"context\"
	\"fmt\"
	\"path\"
	\"path/filepath\"
	\"sort\"
	\"strconv\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

func exceptions(snapshot project.Snapshot) ([]model.Quarantine, []model.BoundaryOptOut, error) {
	var quarantines []model.Quarantine
	var optOuts []model.BoundaryOptOut
	for _, m := range snapshot.Modules {
		items, err := specs(m)
		if err "'!= nil {
			return nil, nil, err
		}
		for _, item := range items {
			if reason, ok := quarantineReason(item); ok {
				quarantines = append(quarantines, model.Quarantine{Module: m.Name, ID: item.ID, Reason: reason})
			}
		}
		if reason := stringValue(m.Document.Frontmatter["boundaries_opt_out"]); reason != "" {
			optOuts = append(optOuts, model.BoundaryOptOut{Module: m.Name, Reason: reason})
		}
	}
	return quarantines, optOuts, nil
}

func Status(root string) (report model.StatusReport, err error) {
	repo, err := repository(root)
	if err != nil {
		return report, err
	}
	root = repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return report, err
	}
	report.Modules = snapshot.Modules
	queue, err := project.Queue(root)
	if err != nil {
		return report, err
	}
	report.Queue, err = waitingQueue(root, queue)
	if err != nil {
		return report, err
	}
	report.Quarantines, report.BoundaryOptOuts, err = exceptions(snapshot)
	if err != nil {
		return report, err
	}
	archives, err := archivedAt(root, repo.Head)
	if err != nil {
		return report, err
	}
	shared, _, err := baseRefs(root, snapshot.Config)
	if err != nil {
		return report, err
	}
	report.Unapproved, err = unapprovedCount(root, shared, archives)
	if err != nil {
		return report, err
	}
	report.DriftSinceDrill = map[string]int{}
	dates := map[string]time.Time{}
	for _, m := range snapshot.Modules {
		date, _, err := drillReport(root, m.Name)
		if err != nil {
			return report, err
		}
		dates[m.Name] = date
		report.DriftSinceDrill[m.Name] = 0
	}
	report.PhaseCommits = map[string]int{}
	type application struct {
		change project.Change
		date time.Time
	}
	var applications []application
	for _, ch := range archives {
		date := implementationDate(ch)
		if drill, ok := dates[ch.Module]; ok && (drill.IsZero() || date.After(drill)) {
			report.DriftSinceDrill[ch.Module]++
		}
		filename := path.Base(ch.Path)
		commits, err := phaseCommits(root, repo.Head, filename)
		if err != nil {
			return report, err
		}
		report.PhaseCommits[strings.TrimSuffix(filename, ".md")] = len(commits)
		applications = append(applications, application{ch, date})
	}
	sort.Slice(applications, func(i, j int) bool {
		if applications[i].date.Equal(applications[j].date) {
			return applications[i].change.Path < applications[j].change.Path
		}
		return applications[i].date.After(applications[j].date)
	})
	if len(applications) != 0 {
		module := applications[0].change.Module
		consecutive := 0
		for _, application := range applications {
			if application.change.Module != module {
				break
			}
			consecutive++
		}
		if consecutive >= 3 {
			report.Diagnostics = append(report.Diagnostics, project.Diagnostic{
				Code: "P08", Warning: true,
				Message: fmt.Sprintf("%d consecutive recently applied changes target %s", consecutive, module),
			})
		}
	}
	costs, ds, err := latestCosts(root, snapshot.Modules)
	report.Costs = costs
	report.Diagnostics = append(report.Diagnostics, ds...)
	return report, err
}

func GC(ctx context.Context, request model.GCRequest) (result model.GCResult, err error) {
	repo, err := repository(request.Root)
	if err != nil {
		return result, err
	}
	root := repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return result, err
	}
	age := 720*time.Hour
	if value, configured := configGroup(snapshot.Config, "gc")["drill_max_age"]; configured {
		age, err = time.ParseDuration(stringValue(value))
		if err != nil || age <= 0 {
			return result, coded("C01", "gc.drill_max_age", fmt.Sprintf("gc.drill_max_age must be a positive Go duration, got %v", value))
		}
	}
	now := request.Now
	if now.IsZero() {
		now = time.Now()
	}
	cutoff := now.UTC().Add(-age)
	archives, err := changePaths(root, archiveDir)
	if err != nil {
		return result, err
	}
	archived := map[string]bool{}
	for _, ch := range archives {
		archived[ch.Slug] = true
	}
	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
	if err != nil {
		return result, err
	}
	live := map[string]bool{}
	common := canonicalPath(repo.CommonDir, root)
	for _, line := range splitNUL(out) {
		if name, ok := strings.CutPrefix(line, "worktree "); ok {
			if other, err := repository(name); err == nil && canonicalPath(other.CommonDir, other.Root) == common {
				live[other.Root] = true
			}
		}
	}
	refs, err := listRefs(ctx, root, claimPrefix, attemptPrefix, drillPrefix)
	if err != nil {
		return result, err
	}
	for _, ref := range refs {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		prune := false
		switch {
		case strings.HasPrefix(ref.name, claimPrefix):
			c, err := readClaim(root, ref.oid)
			prune = err == nil && c.Worktree != "" && !live[canonicalPath(c.Worktree, root)]
		case strings.HasPrefix(ref.name, attemptPrefix):
			slug, n, ok := strings.Cut(strings.TrimPrefix(ref.name, attemptPrefix), "/")
			attempt, numberErr := strconv.Atoi(n)
			prune = ok && numberErr == nil && attempt > 0 && strconv.Itoa(attempt) == n && archived[slug]
		case strings.HasPrefix(ref.name, drillPrefix):
			module, stamp, ok := strings.Cut(strings.TrimPrefix(ref.name, drillPrefix), "/")
			date, dateErr := time.Parse("20060102T150405.000000000", stamp)
			prune = ok && module != "" && dateErr == nil && date.UTC().Before(cutoff)
		}
		if prune {
			if _, deleteErr := git(ctx, root, nil, nil, "update-ref", "-d", ref.name, ref.oid); deleteErr == nil {
				result.Pruned = append(result.Pruned, ref.name)
			} else if current, found, refErr := refStatus(root, ref.name); refErr != nil || found && current == ref.oid {
				err = combine(err, deleteErr)
			}
		}
	}
	sort.Strings(result.Pruned)
	return result, err
}

func canonicalPath(name, root string) string {
	if !filepath.IsAbs(name) {
		name = filepath.Join(root, name)
	}
	if resolved, err := filepath.EvalSymlinks(name); err == nil {
		return resolved
	}
	return filepath.Clean(name)
}
EOF
python3 - <<'"'PY'
from pathlib import Path
p = Path('packages/workflow/internal/lifecycle/status.go')
s = p.read_text().replace('\"github.com/lurkshark/fireweed/packages/project\"', '\"github.com/lurkshark/fireweed/packages/gitstore\"\\n\\t\"github.com/lurkshark/fireweed/packages/project\"').replace('refStatus(root, ref.name)', 'gitstore.Ref(root, ref.name)')
p.write_text(s)
PY
cat > packages/workflow/internal/lifecycle/context.go <<'EOF'
package lifecycle

import (
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/assets\"
	\"github.com/lurkshark/fireweed/packages/contracts\"
)

func Context(root, selector string, full bool) (string, error) {
	s, err := Current(root, selector)
	if err "'!= nil {
		return "", err
	}
	root = s.Snapshot.Root
	target, err := moduleNamed(s.Snapshot, s.Change.Module)
	if err != nil {
		return "", err
	}
	playbook, err := assets.Playbook(playbookName(s))
	if err != nil {
		return "", err
	}
	spec, err := contracts.RenderDocument(target.Document)
	if err != nil {
		return "", err
	}
	var body strings.Builder
	appendContent := func(name string, data []byte) {
		fmt.Fprintf(&body, "'"\\n## %s\\n\\n%s\\n\", name, data)
	}
	for _, name := range []string{s.Change.Path, \"SYSTEM.md\"} {
		data, err := os.ReadFile(filepath.Join(root, name))
		if err "'!= nil {
			return "", err
		}
		appendContent(name, data)
	}
	appendContent(path.Join(target.Dir, "SPEC.md")+" (projected)", []byte(spec))
	body.WriteString("'"\\n\"+playbook+\"\\n\")
	for _, m := range s.Snapshot.Modules {
		if m.Name "'!= target.Name && !dependsOn(s.Snapshot, target.Name, m.Name, true) {
			continue
		}
		err := moduleFiles(s.Snapshot, m, func(name string, entry fs.DirEntry) error {
			if !entry.Type().IsRegular() || name == path.Join(target.Dir, "SPEC.md") || name == s.Change.Path || name == "SYSTEM.md" {
				return nil
			}
			if full || zone(s.Snapshot, m, name) == "interface" {
				data, err := os.ReadFile(filepath.Join(root, name))
				if err != nil {
					return err
				}
				appendContent(name, data)
			} else {
				info, err := entry.Info()
				if err != nil {
					return err
				}
				fmt.Fprintf(&body, "'"\\n- %s (%d bytes)\\n\", name, info.Size())
			}
			return nil
		})
		if err "'!= nil {
			return "", err
		}
	}
	return strings.TrimSpace(body.String())+"'"\\n\", nil
}
EOF
cat > packages/workflow/internal/lifecycle/pr.go <<'EOF'
package lifecycle

import (
	\"context\"
	\"errors\"
	\"fmt\"
	\"path\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project\"
	\"github.com/lurkshark/fireweed/packages/workflow/model\"
)

type endpointConfigError struct{ cause error }

func (e endpointConfigError) Error() string { return e.cause.Error() }
func (e endpointConfigError) Unwrap() error { return e.cause }

func IsEndpointConfigError(err error) bool {
	var endpoint endpointConfigError
	return errors.As(err, &endpoint)
}

func PRBody(root, revisionRange string) (report model.PRBodyReport, err error) {
	repo, err := repository(root)
	if err "'!= nil {
		return report, err
	}
	root = repo.Root
	left, right, ok := strings.Cut(strings.TrimSpace(revisionRange), "..")
	if !ok || left == "" || right == "" || strings.HasPrefix(right, ".") || strings.Contains(right, "..") {
		return report, coded("P01", revisionRange, fmt.Sprintf("expected an explicit commit range, got %q", revisionRange))
	}
	base, err := resolveRevision(root, left)
	if err != nil {
		return report, withCode("P01", revisionRange, fmt.Errorf("resolve range %q: %w", revisionRange, err))
	}
	endpoint, err := resolveRevision(root, right)
	if err != nil {
		return report, withCode("P01", revisionRange, fmt.Errorf("resolve range %q: %w", revisionRange, err))
	}
	snapshot, err := project.LoadAt(root, endpoint)
	if err != nil {
		return report, endpointConfigError{fmt.Errorf("project at %s: %w", endpoint, err)}
	}
	pending, err := committedChanges(root, endpoint, pendingDir)
	if err != nil {
		return report, err
	}
	archives, err := committedChanges(root, endpoint, archiveDir)
	if err != nil {
		return report, err
	}
	revs, err := revisions(root, base+".."+endpoint)
	if err != nil {
		return report, err
	}
	diffs := map[string][]gitstore.Change{}
	touched := map[string][]revision{}
	for _, rev := range revs {
		diff, err := revisionDiff(context.Background(), root, rev)
		if err != nil {
			return report, err
		}
		diffs[rev.oid] = diff
		seen := map[string]bool{}
		for _, change := range diff {
			dir, filename := path.Split(change.Path)
			if (dir == pendingDir+"/" || dir == archiveDir+"/") && path.Ext(filename) == ".md" && !seen[filename] {
				touched[filename] = append(touched[filename], rev)
				seen[filename] = true
			}
		}
	}

	var body strings.Builder
	fmt.Fprintf(&body, "# Fireweed changes'"\\n\\nRange "'`%s..%s`'"\\n\\n## Selected changes\\n\", base, endpoint)
	selected := 0
	for _, ch := range append(pending, archives...) {
		filename := path.Base(ch.Path)
		if len(touched[filename]) == 0 {
			continue
		}
		selected++
		if err := renderChangeHistory(&body, root, endpoint, ch, snapshot, touched[filename], diffs); err "'!= nil {
			return report, err
		}
	}
	if selected == 0 {
		body.WriteString("'"\\nNo selected changes.\\n\")
	}

	body.WriteString(\"\\n## Endpoint exceptions\\n\")
	quarantines, optOuts, err := exceptions(snapshot)
	if err "'!= nil {
		return report, err
	}
	for _, q := range quarantines {
		fmt.Fprintf(&body, "'"\\n- Quarantine in %s, %s — %s\", q.Module, q.ID, q.Reason)
	}
	for _, o := range optOuts {
		fmt.Fprintf(&body, \"\\n- Boundary opt-out in %s — %s\", o.Module, o.Reason)
	}
	for _, ch := range pending {
		if reason := field(ch, \"blocked\"); reason "'!= "" {
			fmt.Fprintf(&body, "'"\\n- Blocked "'`%s` (%s) — %s", ch.Path, ch.Module, reason)
		}
	}
	body.WriteString("'"\\n\\n## Costs\\n\\n| Module | Seconds | Cases per spec | Eval lines | Disposable lines |\\n| --- | ---: | --- | ---: | ---: |\\n\")
	costs, diagnostics, err := latestCosts(root, snapshot.Modules)
	if err "'!= nil {
		return report, err
	}
	for _, cost := range costs {
		var cases []string
		for _, id := range sortedKeys(cost.CasesPerSpec) {
			cases = append(cases, fmt.Sprintf("%s: %d", id, cost.CasesPerSpec[id]))
		}
		fmt.Fprintf(&body, "| %s | %g | %s | %d | %d |'"\\n\", cost.Module, cost.Seconds, strings.Join(cases, \", \"), cost.EvalLines, cost.DisposableLines)
	}
	artifacts, err := artifactWarnings(root, base, endpoint, snapshot)
	if err "'!= nil {
		return report, err
	}
	report.Diagnostics = append(diagnostics, artifacts...)
	if len(report.Diagnostics) != 0 {
		body.WriteString("'"\\n## Diagnostics\\n\")
		for _, d := range report.Diagnostics {
			fmt.Fprintf(&body, \"\\n- %s — %s\", d.Code, d.Message)
		}
	}
	report.Body = strings.TrimSpace(body.String())+\"\\n\"
	return report, nil
}

func changeKind(ch project.Change) string {
	switch {
	case ch.Document.Frontmatter[\"regenerate\"] == true:
		return \"regenerate\"
	case len(stringList(ch.Document.Frontmatter[\"eval_fix\"])) "'!= 0:
		return "eval_fix"
	case len(stringList(ch.Document.Frontmatter["fixes"])) != 0:
		return "fix"
	case hasSpecWork(ch):
		return "spec"
	default:
		return "maintenance"
	}
}

func renderChangeHistory(body *strings.Builder, root, endpoint string, ch project.Change, snapshot project.Snapshot, touched []revision, diffs map[string][]gitstore.Change) error {
	filename := path.Base(ch.Path)
	rounds := intValue(ch.Document.Frontmatter["audit_rounds"], 0)
	fmt.Fprintf(body, "'"\\n### "'`%s`'"\\n\\nModule **%s** · Kind **%s** · audit_rounds **%d**\\n\", ch.Path, ch.Module, changeKind(ch), rounds)
	history, err := phaseCommits(root, endpoint, filename)
	if err "'!= nil {
		return err
	}
	seen := map[string]bool{}
	categories := map[string]map[string]bool{}
	findings := map[string]bool{}
	provenance := map[string]bool{}
	var resultLines []string
	record := func(doc project.Change) {
		for _, line := range strings.Split(section(doc.Document, "Audit"), "'"\\n\") {
			line = strings.TrimSpace(line)
			if _, ok := parseFinding(line); ok {
				findings[line] = true
			}
		}
		for _, line := range strings.Split(section(doc.Document, \"Provenance\"), \"\\n\") {
			line = strings.TrimSpace(line)
			if strings.Contains(line, \"tests=\") && "'!provenance[line] {
				provenance[line] = true
				resultLines = append(resultLines, line)
			}
		}
	}
	record(ch)
	for _, rev := range append(history, touched...) {
		if seen[rev.oid] {
			continue
		}
		seen[rev.oid] = true
		doc, err := documentAt(root, rev.oid, filename)
		if err != nil {
			return err
		}
		if doc.Path == "" {
			continue
		}
		record(doc)
		phase, err := phaseAt(root, rev.parent, filename)
		if err != nil {
			return err
		}
		if phase == "" {
			switch field(doc, "phase") {
			case "audit":
				phase = "evals"
			case "applied":
				phase = "impl"
			case "impl":
				phase = "evals"
				if !hasSpecWork(doc) || doc.Document.Frontmatter["regenerate"] == true {
					phase = "impl"
				}
			}
		}
		metadata, err := project.LoadAt(root, rev.oid)
		if err != nil {
			// History can reach a commit before the first valid project
			// configuration. Its document findings still belong in the report.
			continue
		}
		m, err := moduleNamed(metadata, doc.Module)
		if err != nil {
			continue
		}
		if path.Dir(doc.Path) == pendingDir {
			applyZoneSet(&m, doc)
		}
		changes, ok := diffs[rev.oid]
		if !ok {
			changes, err = revisionDiff(context.Background(), root, rev)
			if err != nil {
				return err
			}
			diffs[rev.oid] = changes
		}
		for _, change := range changes {
			if category := changeCategory(metadata, m, phase, change); category != "" {
				if categories[category] == nil {
					categories[category] = map[string]bool{}
				}
				categories[category][change.Path] = true
			}
		}
	}
	if len(resultLines) != 0 {
		fmt.Fprintf(body, "'"\\n**Recorded eval results**\\n\\n%s\\n\", strings.Join(resultLines, \"\\n\"))
	}
	for _, category := range []string{\"Eval edits\", \"Interface edits\", \"Disposable stubs added in evals\", \"Module manifests and other files\", \"Worker-writable edits\", \"Preserved edits\", \"Disposable implementation edits\"} {
		if files := categories[category]; len(files) "'!= 0 {
			fmt.Fprintf(body, "'"\\n**%s**\\n\\n\", category)
			for _, file := range sortedKeys(files) {
				fmt.Fprintf(body, \"- "'`%s`'"\\n\", file)
			}
		}
	}
	for _, group := range []struct {
		name string
		modified bool
	}{{\"Added\", false}, {\"Modified\", true}} {
		items := ch.Added
		if group.modified {
			items = ch.Modified
		}
		for _, item := range items {
			if item.Tag == \"review\" {
				fmt.Fprintf(body, \"\\n- Review %s **%s** — %s\\n\", strings.ToLower(group.name), item.ID, strings.TrimSpace(item.Body))
			}
		}
	}
	if len(findings) "'!= 0 {
		fmt.Fprintf(body, "'"\\n**Audit findings**\\n\\n%s\\n\", strings.Join(sortedKeys(findings), \"\\n\"))
	}
	for _, label := range []string{\"Decisions\", \"Removed decisions\"} {
		if lines := strings.TrimSpace(section(ch.Document, label)); lines "'!= "" {
			fmt.Fprintf(body, "'"\\n**%s**\\n\\n%s\\n\", label, lines)
		}
	}
	return nil
}

func applyZoneSet(m *project.Module, ch project.Change) {
	set := configGroup(ch.Document.Frontmatter, \"set\")
	for key, patterns := range map[string]*[]string{
		\"protected\": &m.Zones.Protected, \"evals\": &m.Zones.Evals,
		\"interface\": &m.Zones.Interface, \"append_only\": &m.Zones.AppendOnly,
		\"preserve\": &m.Zones.Preserve, \"disposable\": &m.Zones.Disposable,
	} {
		if value, ok := set[key]; ok {
			*patterns = stringList(value)
		}
	}
}

func changeCategory(snapshot project.Snapshot, module project.Module, phase string, ch gitstore.Change) string {
	if matchesAny(stringList(snapshot.Config[\"worker_writable\"]), ch.Path) && zone(snapshot, module, ch.Path) "'!= "protected" {
		return "Worker-writable edits"
	}
	if !belongsTo(snapshot, module, ch.Path) {
		return ""
	}
	switch zone(snapshot, module, ch.Path) {
	case "evals":
		return "Eval edits"
	case "interface":
		return "Interface edits"
	case "preserve":
		return "Preserved edits"
	case "disposable":
		if phase == "evals" && strings.HasPrefix(ch.Status, "A") {
			return "Disposable stubs added in evals"
		}
		return "Disposable implementation edits"
	default:
		return "Module manifests and other files"
	}
}

func artifactWarnings(root, base, endpoint string, snapshot project.Snapshot) ([]project.Diagnostic, error) {
	out, err := git(context.Background(), root, []string{"GIT_ATTR_SOURCE="+endpoint}, nil,
		"diff", "--numstat", "-z", "--diff-filter=A", "--no-renames", "--no-ext-diff", "--no-textconv", base, endpoint, "--")
	if err != nil {
		return nil, err
	}
	var warnings []project.Diagnostic
	for _, record := range splitNUL(out) {
		fields := strings.SplitN(record, "'"\\t\", 3)
		if len(fields) "'!= 3 {
			return nil, fmt.Errorf("invalid Git numstat output")
		}
		name := fields[2]
		for _, m := range snapshot.Modules {
			if !belongsTo(snapshot, m, name) || zone(snapshot, m, name) != "disposable" {
				continue
			}
			size, err := gitOutput(root, "cat-file", "-s", endpoint+":"+name)
			if err != nil {
				return nil, err
			}
			bytes := intValue(size, -1)
			binary := fields[0] == "-" && fields[1] == "-"
			if binary || bytes > 262144 {
				description := "oversized"
				if binary {
					description = "binary"
				}
				warnings = append(warnings, project.Diagnostic{Code: "B04", Path: name, Warning: true,
					Message: fmt.Sprintf("%s disposable addition %s (%d bytes)", description, name, bytes)})
			}
			break
		}
	}
	return warnings, nil
}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 580ms:
exec
/bin/zsh -lc "gofmt -w packages/workflow/internal/lifecycle/*.go; rg -n '"'^func |''^type '"' packages/workflow/internal/lifecycle; rg -n 'GIT_|\\.fireweed|checks\\.Require|StaleSpecs|BaseRef' packages/workflow/internal/lifecycle" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 630ms:
packages/workflow/internal/lifecycle/verify.go:17:type verification struct {
packages/workflow/internal/lifecycle/verify.go:24:func Verify(ctx context.Context, request model.Request) (model.Result, error) {
packages/workflow/internal/lifecycle/verify.go:29:func verify(ctx context.Context, request model.Request) (result model.Result, step verification, err error) {
packages/workflow/internal/lifecycle/verify.go:216:func combine(first, second error) error {
packages/workflow/internal/lifecycle/verify.go:226:func quarantineReason(item contracts.Item) (string, bool) {
packages/workflow/internal/lifecycle/verify.go:235:func nonQuarantinedSpecs(items []contracts.Item) []contracts.Item {
packages/workflow/internal/lifecycle/verify.go:245:func enforcedSpecs(items []contracts.Item) []contracts.Item {
packages/workflow/internal/lifecycle/verify.go:255:func changeIDs(ch project.Change, items []contracts.Item, audit bool) []string {
packages/workflow/internal/lifecycle/queue.go:11:func waitingQueue(root string, queue []project.Change) ([]model.StatusEntry, error) {
packages/workflow/internal/lifecycle/queue.go:51:func Eligible(root string) ([]project.Change, error) {
packages/workflow/internal/lifecycle/queue.go:63:func eligible(root string, queue []project.Change) ([]project.Change, error) {
packages/workflow/internal/lifecycle/queue.go:77:func checkStale(root string, ch project.Change) error {
packages/workflow/internal/lifecycle/pr.go:15:type endpointConfigError struct{ cause error }
packages/workflow/internal/lifecycle/pr.go:17:func (e endpointConfigError) Error() string { return e.cause.Error() }
packages/workflow/internal/lifecycle/pr.go:18:func (e endpointConfigError) Unwrap() error { return e.cause }
packages/workflow/internal/lifecycle/pr.go:20:func IsEndpointConfigError(err error) bool {
packages/workflow/internal/lifecycle/pr.go:25:func PRBody(root, revisionRange string) (report model.PRBodyReport, err error) {
packages/workflow/internal/lifecycle/pr.go:137:func changeKind(ch project.Change) string {
packages/workflow/internal/lifecycle/pr.go:152:func renderChangeHistory(body *strings.Builder, root, endpoint string, ch project.Change, snapshot project.Snapshot, touched []revision, diffs map[string][]gitstore.Change) error {
packages/workflow/internal/lifecycle/pr.go:277:func applyZoneSet(m *project.Module, ch project.Change) {
packages/workflow/internal/lifecycle/pr.go:290:func changeCategory(snapshot project.Snapshot, module project.Module, phase string, ch gitstore.Change) string {
packages/workflow/internal/lifecycle/pr.go:314:func artifactWarnings(root, base, endpoint string, snapshot project.Snapshot) ([]project.Diagnostic, error) {
packages/workflow/internal/lifecycle/attempt.go:15:func Block(request model.BlockRequest) (result model.BlockResult, err error) {
packages/workflow/internal/lifecycle/attempt.go:96:func resumeDocument(root, oid string, current project.Change) (string, error) {
packages/workflow/internal/lifecycle/attempt.go:119:func restoreAttempt(root, oid, file string) error {
packages/workflow/internal/lifecycle/claim.go:26:type claim struct {
packages/workflow/internal/lifecycle/claim.go:35:func (c claim) ref() string { return claimPrefix + c.Change }
packages/workflow/internal/lifecycle/claim.go:37:func (c claim) blob(root string) (string, error) {
packages/workflow/internal/lifecycle/claim.go:45:func readClaim(root, oid string) (claim, error) {
packages/workflow/internal/lifecycle/claim.go:62:func normalizeMap(m map[string]any) map[string]any {
packages/workflow/internal/lifecycle/claim.go:69:func normalizeNumber(value any) any {
packages/workflow/internal/lifecycle/claim.go:88:func heldClaim(repo gitstore.Repository, selector string) (claim, string, error) {
packages/workflow/internal/lifecycle/claim.go:145:func (c claim) state(repo gitstore.Repository, oid string) (model.State, error) {
packages/workflow/internal/lifecycle/claim.go:157:func Current(root, selector string) (model.State, error) {
packages/workflow/internal/lifecycle/claim.go:178:func Release(root, selector string) error {
packages/workflow/internal/lifecycle/claim.go:190:func Start(root, selector string) (model.State, error) {
packages/workflow/internal/lifecycle/claim.go:194:func StartFromAttempt(root, selector string, attempt int) (model.State, error) {
packages/workflow/internal/lifecycle/claim.go:198:func start(root, selector string, attempt *int) (s model.State, err error) {
packages/workflow/internal/lifecycle/claim.go:407:func validateRun(run string) error {
packages/workflow/internal/lifecycle/claim.go:414:func hasSpecWork(c project.Change) bool {
packages/workflow/internal/lifecycle/claim.go:420:func assignIDs(root string, ch project.Change, projection project.Snapshot, queue, archives []project.Change, source string) (string, error) {
packages/workflow/internal/lifecycle/claim.go:494:func installAssignedSpecs(root string, queue []project.Change, ch project.Change, projected *project.Snapshot) error {
packages/workflow/internal/lifecycle/claim.go:531:func renderSpecEdits(items []contracts.Item, ch project.Change) string {
packages/workflow/internal/lifecycle/claim.go:555:func removeDisposable(s project.Snapshot, name string) error {
packages/workflow/internal/lifecycle/claim.go:586:func approvedWorking(s model.State) (project.Change, error) {
packages/workflow/internal/lifecycle/claim.go:640:func editableSection(name, phase string) bool {
packages/workflow/internal/lifecycle/claim.go:650:func immutableBody(body, phase string) string {
packages/workflow/internal/lifecycle/claim.go:665:func archivePath(changePath string) string {
packages/workflow/internal/lifecycle/git.go:13:type gitRef struct {
packages/workflow/internal/lifecycle/git.go:18:func listRefs(ctx context.Context, root string, prefixes ...string) ([]gitRef, error) {
packages/workflow/internal/lifecycle/git.go:33:func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
packages/workflow/internal/lifecycle/git.go:58:type gitIndex struct {
packages/workflow/internal/lifecycle/git.go:64:func temporaryIndex(ctx context.Context, root, base string) (gitIndex, func(), error) {
packages/workflow/internal/lifecycle/git.go:78:func (idx gitIndex) run(input []byte, args ...string) ([]byte, error) {
packages/workflow/internal/lifecycle/git.go:82:func (idx gitIndex) stageWorktree() error {
packages/workflow/internal/lifecycle/git.go:87:func (idx gitIndex) stageDocument(name string, data []byte) error {
packages/workflow/internal/lifecycle/git.go:96:func (idx gitIndex) commit(parent, message string) (string, error) {
packages/workflow/internal/lifecycle/git.go:119:func (idx gitIndex) checkout(revision string) error {
packages/workflow/internal/lifecycle/git.go:124:func replaceIndex(root, revision string) error {
packages/workflow/internal/lifecycle/git.go:129:func commitRefs(root, oldHead, newHead, extra string) error {
packages/workflow/internal/lifecycle/git.go:135:func claimUpdate(ref, oldOID, newOID string) string {
packages/workflow/internal/lifecycle/git.go:142:func transitionError(root, base, name string, err error) error {
packages/workflow/internal/lifecycle/costs.go:22:func moduleCost(snapshot project.Snapshot, m project.Module, report checks.Report) (model.ModuleCost, error) {
packages/workflow/internal/lifecycle/costs.go:59:func newCostSnapshot(state model.State, modules []model.ModuleCost, ds []project.Diagnostic) model.CostSnapshot {
packages/workflow/internal/lifecycle/costs.go:63:func costWarnings(snapshot project.Snapshot, modules []model.ModuleCost) []project.Diagnostic {
packages/workflow/internal/lifecycle/costs.go:104:func Costs(root string) ([]model.CostSnapshot, error) {
packages/workflow/internal/lifecycle/costs.go:134:func latestCosts(root string, modules []project.Module) ([]model.ModuleCost, []project.Diagnostic, error) {
packages/workflow/internal/lifecycle/costs.go:178:func costDiagnosticMatches(d project.Diagnostic, cost model.ModuleCost, module project.Module) bool {
packages/workflow/internal/lifecycle/status.go:18:func exceptions(snapshot project.Snapshot) ([]model.Quarantine, []model.BoundaryOptOut, error) {
packages/workflow/internal/lifecycle/status.go:38:func Status(root string) (report model.StatusReport, err error) {
packages/workflow/internal/lifecycle/status.go:130:func GC(ctx context.Context, request model.GCRequest) (result model.GCResult, err error) {
packages/workflow/internal/lifecycle/status.go:207:func canonicalPath(name, root string) string {
packages/workflow/internal/lifecycle/done.go:21:func Done(ctx context.Context, request model.Request) (result model.Result, err error) {
packages/workflow/internal/lifecycle/done.go:167:func playbookName(s model.State) string {
packages/workflow/internal/lifecycle/done.go:174:func appendProvenance(body, line string) string {
packages/workflow/internal/lifecycle/done.go:185:func applyToModule(root string, ch project.Change, files *fileEdits) error {
packages/workflow/internal/lifecycle/done.go:224:func decisionText(line string) string {
packages/workflow/internal/lifecycle/done.go:237:func applyDecisions(prior, added, removed, slug string) string {
packages/workflow/internal/lifecycle/done.go:260:type fileEdits struct {
packages/workflow/internal/lifecycle/done.go:265:func (files *fileEdits) replace(name string, data []byte, remove bool) error {
packages/workflow/internal/lifecycle/done.go:288:func (files *fileEdits) write(name string, data []byte) error {
packages/workflow/internal/lifecycle/done.go:291:func (files *fileEdits) remove(name string) error { return files.replace(name, nil, true) }
packages/workflow/internal/lifecycle/done.go:293:func (files *fileEdits) rollback() error {
packages/workflow/internal/lifecycle/audit.go:13:type auditFinding struct {
packages/workflow/internal/lifecycle/audit.go:20:type auditSubmission struct {
packages/workflow/internal/lifecycle/audit.go:27:func parseFinding(line string) (auditFinding, bool) {
packages/workflow/internal/lifecycle/audit.go:35:func checkAudit(state model.State, ch project.Change, items []contracts.Item) (auditSubmission, error) {
packages/workflow/internal/lifecycle/audit.go:96:func (a auditSubmission) body() string {
packages/workflow/internal/lifecycle/audit.go:107:func (a auditSubmission) blocked(rounds, max int) string {
packages/workflow/internal/lifecycle/context.go:15:func Context(root, selector string, full bool) (string, error) {
packages/workflow/internal/lifecycle/common.go:33:func git(ctx context.Context, root string, env []string, input []byte, args ...string) ([]byte, error) {
packages/workflow/internal/lifecycle/common.go:50:func gitOutput(root string, args ...string) (string, error) {
packages/workflow/internal/lifecycle/common.go:55:func repository(root string) (gitstore.Repository, error) {
packages/workflow/internal/lifecycle/common.go:66:func coded(code, subject, message string) error {
packages/workflow/internal/lifecycle/common.go:70:func withCode(code, subject string, err error) error {
packages/workflow/internal/lifecycle/common.go:80:func diagnosticError(d project.Diagnostic) error {
packages/workflow/internal/lifecycle/common.go:87:func firstError(ds []project.Diagnostic) error {
packages/workflow/internal/lifecycle/common.go:96:func staleError(err error) error {
packages/workflow/internal/lifecycle/common.go:103:func stringValue(v any) string {
packages/workflow/internal/lifecycle/common.go:108:func intValue(v any, fallback int) int {
packages/workflow/internal/lifecycle/common.go:126:func stringList(v any) []string {
packages/workflow/internal/lifecycle/common.go:142:func configGroup(config map[string]any, key string) map[string]any {
packages/workflow/internal/lifecycle/common.go:147:func field(c project.Change, key string) string { return stringValue(c.Document.Frontmatter[key]) }
packages/workflow/internal/lifecycle/common.go:148:func section(doc contracts.Document, name string) string {
packages/workflow/internal/lifecycle/common.go:153:func moduleNamed(s project.Snapshot, name string) (project.Module, error) {
packages/workflow/internal/lifecycle/common.go:162:func specs(m project.Module) ([]contracts.Item, error) {
packages/workflow/internal/lifecycle/common.go:166:func dependsOn(s project.Snapshot, from, target string, transitive bool) bool {
packages/workflow/internal/lifecycle/common.go:185:func relativeTo(dir, file string) (string, bool) {
packages/workflow/internal/lifecycle/common.go:193:func matchesAny(patterns []string, name string) bool {
packages/workflow/internal/lifecycle/common.go:202:func zone(s project.Snapshot, m project.Module, name string) string {
packages/workflow/internal/lifecycle/common.go:213:func belongsTo(s project.Snapshot, m project.Module, name string) bool {
packages/workflow/internal/lifecycle/common.go:227:func moduleFiles(s project.Snapshot, m project.Module, visit func(string, fs.DirEntry) error) error {
packages/workflow/internal/lifecycle/common.go:250:func identity(file string) project.Change {
packages/workflow/internal/lifecycle/common.go:260:func changePaths(root, dir string) ([]project.Change, error) {
packages/workflow/internal/lifecycle/common.go:277:func pathsAt(root, revision, dir string) ([]project.Change, error) {
packages/workflow/internal/lifecycle/common.go:291:func splitNUL(out []byte) []string {
packages/workflow/internal/lifecycle/common.go:295:func nulFields(out []byte) int {
packages/workflow/internal/lifecycle/common.go:306:func selectChange(queue []project.Change, selector string) (project.Change, error) {
packages/workflow/internal/lifecycle/common.go:327:func sortedKeys[V any](m map[string]V) []string {
packages/workflow/internal/lifecycle/common.go:336:func readChange(ch project.Change, data []byte) (project.Change, error) {
packages/workflow/internal/lifecycle/common.go:355:func writeFile(file string, data []byte) error {
packages/workflow/internal/lifecycle/common.go:379:func writeJSON(file string, data any) error {
packages/workflow/internal/lifecycle/common.go:387:func requireClean(root, action string) error {
packages/workflow/internal/lifecycle/common.go:395:func parseDate(v any) time.Time {
packages/workflow/internal/lifecycle/common.go:407:func implementationDate(ch project.Change) time.Time {
packages/workflow/internal/lifecycle/common.go:420:func drillReport(root, module string) (time.Time, bool, error) {
packages/workflow/internal/lifecycle/history.go:16:type revision struct {
packages/workflow/internal/lifecycle/history.go:21:func revisions(root, rangeExpr string) ([]revision, error) {
packages/workflow/internal/lifecycle/history.go:40:func resolveRevision(root, name string) (string, error) {
packages/workflow/internal/lifecycle/history.go:44:func committedChanges(root, rev, dir string) ([]project.Change, error) {
packages/workflow/internal/lifecycle/history.go:62:func archivedAt(root, rev string) ([]project.Change, error) {
packages/workflow/internal/lifecycle/history.go:76:func existsAt(root, rev, name string) bool {
packages/workflow/internal/lifecycle/history.go:84:func baseRefs(root string, config map[string]any) (shared, local string, err error) {
packages/workflow/internal/lifecycle/history.go:118:func unapprovedCount(root, shared string, archives []project.Change) (int, error) {
packages/workflow/internal/lifecycle/history.go:138:func historySelection(root, selector string) (project.Change, error) {
packages/workflow/internal/lifecycle/history.go:150:func documentAt(root, rev, filename string) (project.Change, error) {
packages/workflow/internal/lifecycle/history.go:167:func phaseAt(root, rev, filename string) (string, error) {
packages/workflow/internal/lifecycle/history.go:188:func revisionDiff(ctx context.Context, root string, rev revision) ([]gitstore.Change, error) {
packages/workflow/internal/lifecycle/history.go:196:func phaseCommits(root, endpoint, filename string) ([]revision, error) {
packages/workflow/internal/lifecycle/history.go:230:func Commits(root, selector string) ([]string, error) {
packages/workflow/internal/lifecycle/history.go:250:func Abandon(root, selector string) (string, error) {
packages/workflow/internal/lifecycle/verify.go:117:	runDir := filepath.Join(root, ".fireweed", "runs", state.Run, state.Phase)
packages/workflow/internal/lifecycle/verify.go:168:			reportErr = checks.RequireCoverage(report, enforcedSpecs(moduleItems))
packages/workflow/internal/lifecycle/verify.go:184:					reportErr = checks.RequireCoverage(report, listed)
packages/workflow/internal/lifecycle/verify.go:187:				reportErr = checks.RequireFailFirst(report, changeIDs(state.Change, moduleItems, false), stringList(state.Change.Document.Frontmatter["already_satisfied"]))
packages/workflow/internal/lifecycle/verify.go:189:				reportErr = checks.RequirePass(report, nonQuarantinedSpecs(moduleItems))
packages/workflow/internal/lifecycle/verify.go:191:				reportErr = checks.RequirePass(report, enforcedSpecs(moduleItems))
packages/workflow/internal/lifecycle/history.go:90:	resolved, err := gitstore.BaseRef(root, remote, base)
packages/workflow/internal/lifecycle/claim.go:570:	backup, err := os.MkdirTemp(s.Root, ".fireweed-regenerate-")
packages/workflow/internal/lifecycle/common.go:239:			if file != dir && (entry.Name() == ".git" || entry.Name() == ".fireweed" || !belongsTo(s, m, name)) {
packages/workflow/internal/lifecycle/common.go:359:	tmp, err := os.CreateTemp(filepath.Dir(file), ".fireweed-write-")
packages/workflow/internal/lifecycle/common.go:421:	data, err := os.ReadFile(filepath.Join(root, ".fireweed", "drills", module+".md"))
packages/workflow/internal/lifecycle/pr.go:315:	out, err := git(context.Background(), root, []string{"GIT_ATTR_SOURCE=" + endpoint}, nil,
packages/workflow/internal/lifecycle/costs.go:109:	files, err := filepath.Glob(filepath.Join(repo.Root, ".fireweed", "runs", "*", "*", "costs.json"))
packages/workflow/internal/lifecycle/queue.go:78:	stale, err := project.StaleSpecs(root, ch.Slug)
packages/workflow/internal/lifecycle/git.go:70:	idx := gitIndex{root: root, env: []string{"GIT_INDEX_FILE=" + filepath.Join(dir, "index")}, ctx: ctx}
packages/workflow/internal/lifecycle/git.go:83:	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
packages/workflow/internal/lifecycle/git.go:104:		{"user.name", "Fireweed", "GIT_AUTHOR_NAME", "GIT_COMMITTER_NAME"},
packages/workflow/internal/lifecycle/git.go:105:		{"user.email", "fireweed@localhost", "GIT_AUTHOR_EMAIL", "GIT_COMMITTER_EMAIL"},

exec
/bin/zsh -lc "sed -n '198,407p' packages/workflow/internal/lifecycle/claim.go; cat packages/workflow/evals/eligible_test.go; sed -n '1,80p' packages/workflow/evals/costs_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 544ms:
func start(root, selector string, attempt *int) (s model.State, err error) {
	repo, err := repository(root)
	if err != nil {
		return s, err
	}
	root = repo.Root
	if err = requireClean(root, "start"); err != nil {
		// A second Start from the owning checkout is an ownership error.
		if c, _, claimErr := heldClaim(repo, selector); claimErr == nil {
			return s, coded("P02", c.Path, "claim already held by this worktree")
		}
		return s, err
	}
	queue, err := project.Queue(root)
	if err != nil {
		return s, staleError(err)
	}
	var chosen project.Change
	if selector != "" {
		chosen, err = selectChange(queue, selector)
		if err != nil {
			return s, err
		}
	}
	ready, err := eligible(root, queue)
	if err != nil {
		return s, staleError(err)
	}
	if len(ready) == 0 || selector != "" && chosen.Path != ready[0].Path {
		return s, coded("P01", selector, "change is not eligible queue head")
	}
	chosen = ready[0]
	var savedOID string
	if attempt != nil {
		if *attempt <= 0 {
			return s, coded("P01", chosen.Path, "attempt number must be positive")
		}
		ref := attemptPrefix + chosen.Slug + "/" + strconv.Itoa(*attempt)
		var exists bool
		savedOID, exists, err = gitstore.Ref(root, ref)
		if err != nil {
			return s, err
		}
		if !exists {
			return s, coded("P01", chosen.Path, "saved attempt not found: "+ref)
		}
		if field(chosen, "phase") == "" || field(chosen, "run") == "" {
			return s, coded("P01", chosen.Path, "resuming an attempt requires a recorded run and phase")
		}
	}
	s.Snapshot, err = project.Project(root, chosen.Slug)
	if err != nil {
		return s, staleError(err)
	}
	s.Snapshot.Root = root
	for _, d := range project.Check(root) {
		if !d.Warning && (d.Path == "" || d.Path == chosen.Path || !strings.HasPrefix(d.Path, "changes/")) {
			return s, staleError(diagnosticError(d))
		}
	}
	if err = checkStale(root, chosen); err != nil {
		return s, err
	}

	shared, local, err := baseRefs(root, s.Snapshot.Config)
	if err != nil {
		return s, err
	}
	applied, err := archivedAt(root, repo.Head)
	if err != nil {
		return s, err
	}
	unapproved, err := unapprovedCount(root, shared, applied)
	if err != nil {
		return s, err
	}
	limit := intValue(configGroup(s.Snapshot.Config, "authoring")["max_unapproved"], 5)
	if unapproved >= limit {
		message := fmt.Sprintf("%d unapproved archived changes reaches the limit of %d", unapproved, limit)
		s.Diagnostics = append(s.Diagnostics, project.Diagnostic{Code: "P07", Path: chosen.Path, Message: message})
		return s, coded("P07", chosen.Path, message)
	}
	if !existsAt(root, local, chosen.Path) && !existsAt(root, shared, chosen.Path) {
		s.Diagnostics = append(s.Diagnostics, project.Diagnostic{
			Code: "P04", Path: chosen.Path, Warning: true, Message: "change is absent from the local and shared bases",
		})
	}

	file := filepath.Join(root, chosen.Path)
	original, err := os.ReadFile(file)
	if err != nil {
		return s, err
	}
	if attempt == nil {
		assigned, assignErr := assignIDs(root, chosen, s.Snapshot, queue, applied, string(original))
		if assignErr != nil {
			return s, assignErr
		}
		chosen, err = readChange(chosen, []byte(assigned))
		if err != nil {
			return s, err
		}
	}
	front := chosen.Document.Frontmatter
	phase := field(chosen, "phase")
	unstarted := phase == ""
	regenerate := front["regenerate"] == true
	if regenerate {
		var latest time.Time
		for _, archive := range applied {
			if archive.Module == chosen.Module {
				if date := implementationDate(archive); date.After(latest) {
					latest = date
				}
			}
		}
		date, passing, reportErr := drillReport(root, chosen.Module)
		if reportErr != nil {
			return s, reportErr
		}
		if !passing || !date.After(latest) {
			s.Diagnostics = append(s.Diagnostics, project.Diagnostic{
				Code: "P10", Path: chosen.Path, Warning: true, Message: "module " + chosen.Module + " has no passing drill newer than its latest implementation",
			})
		}
	}
	if unstarted {
		phase = "evals"
		if regenerate || !hasSpecWork(chosen) {
			phase = "impl"
		}
		front["phase"] = phase
	}
	if field(chosen, "run") == "" {
		var id [16]byte
		if _, err = rand.Read(id[:]); err != nil {
			return s, err
		}
		front["run"] = hex.EncodeToString(id[:])
	}
	if err = validateRun(field(chosen, "run")); err != nil {
		return s, err
	}
	approved, err := contracts.RenderDocument(chosen.Document)
	if err != nil {
		return s, err
	}
	if attempt != nil {
		approved = string(original)
	}

	// Reproject with real IDs before storing the claim. The original projection
	// may use provisional IDs, so replace its spec section from the preceding
	// projection and the approved change's assigned edits.
	if err = installAssignedSpecs(root, queue, chosen, &s.Snapshot); err != nil {
		return s, staleError(err)
	}
	c := claim{Worktree: root, Change: chosen.Slug, Path: chosen.Path, Base: repo.Head, Approved: approved, Snapshot: s.Snapshot}
	oid, err := c.blob(root)
	if err != nil {
		return s, err
	}
	if err = gitstore.CASRef(root, c.ref(), "", oid); err != nil {
		return s, withCode("P02", c.Path, fmt.Errorf("claim already held: %w", err))
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, gitstore.CASRef(root, c.ref(), oid, ""))
		}
	}()
	if repoNow, inspectErr := repository(root); inspectErr != nil {
		return s, inspectErr
	} else if repoNow.Head != repo.Head {
		return s, coded("P05", c.Path, "stale phase base: HEAD changed during Start")
	}
	diagnostics := s.Diagnostics
	s, err = c.state(repo, oid)
	if err != nil {
		return s, err
	}
	s.Diagnostics = diagnostics
	working := approved
	if attempt != nil {
		working, err = resumeDocument(root, savedOID, s.Change)
		if err != nil {
			return s, err
		}
	}
	if working != string(original) {
		if err = writeFile(file, []byte(working)); err != nil {
			return s, err
		}
		defer func() {
			if err != nil {
				err = errors.Join(err, writeFile(file, original))
			}
		}()
	}
	if attempt != nil {
		s.Change, err = readChange(s.Change, []byte(working))
		if err == nil {
			err = restoreAttempt(root, savedOID, chosen.Path)
		}
	} else if unstarted && regenerate {
		err = removeDisposable(s.Snapshot, chosen.Module)
	}
	return s, err
}

func validateRun(run string) error {
package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

// A real Start/Block leaves the first queued change blocked, committed, and
// unclaimed. The other changes are ordinary approved fixture-local queue files.
func eligibilityBase(t *testing.T, blockedModule string) string {
	t.Helper()
	root := blockFixture(t, "", "", "", "", "")
	write(t, root, "packages/delta/SPEC.md", "---\nname: delta\nid_prefix: DEL\ndepends_on: []\n---\n# Delta\n\n## Specs\n\n### DEL-001\nDelta MUST work.\n")
	write(t, root, "changes/0001-manual.md", fmt.Sprintf("---\nmodule: %s\n---\n\n## Motivation\nThe first change blocks.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n", blockedModule))
	commit(t, root, "blocked change fixture")
	checkEligibilityFixture(t, root)
	start(t, root)
	block(t, root, "fixture blocked")
	return root
}

func eligibilityChange(t *testing.T, root string, number int, slug, module, set, iface string) {
	t.Helper()
	write(t, root, fmt.Sprintf("changes/%04d-%s.md", number, slug), fmt.Sprintf("---\nmodule: %s\n%s---\n\n## Motivation\nThe queued change is needed.\n\n## Interface\n%s\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n", module, set, iface))
}

func checkEligibilityFixture(t *testing.T, root string) {
	t.Helper()
	for _, d := range project.Check(root) {
		if !d.Warning {
			t.Fatalf("invalid eligibility fixture: %s %s: %s", d.Code, d.Path, d.Message)
		}
	}
	queue, err := project.Queue(root)
	if err != nil {
		t.Fatalf("fixture Queue: %v", err)
	}
	for _, ch := range queue {
		if _, err := project.Project(root, ch.Slug); err != nil {
			t.Fatalf("fixture projection through %s: %v", ch.Slug, err)
		}
	}
}

func readyEligibilityFixture(t *testing.T, root string) {
	t.Helper()
	checkEligibilityFixture(t, root)
	commit(t, root, "approved eligibility queue")
}

func eligibleSlugs(t *testing.T, root string) []string {
	t.Helper()
	changes, err := workflow.Eligible(root)
	if err != nil {
		t.Fatalf("Eligible: %v", err)
	}
	slugs := make([]string, 0, len(changes))
	for _, ch := range changes {
		if ch.Path == "" || ch.Module == "" {
			t.Fatalf("Eligible returned incomplete public change: %#v", ch)
		}
		slugs = append(slugs, ch.Slug)
	}
	return slugs
}

func assertEligible(t *testing.T, root string, want ...string) {
	t.Helper()
	want = append([]string{}, want...)
	if got := eligibleSlugs(t, root); !reflect.DeepEqual(got, want) {
		t.Errorf("eligible slugs = %v, want %v", got, want)
	}
}

func Test_FLOW_068_EligibleOrderAndEmptyQueues(t *testing.T) {
	t.Run("skips waiting and retains independent order", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		eligibilityChange(t, root, 3, "independent", "delta", "", "")
		eligibilityChange(t, root, 4, "transitive-wait", "gamma", "", "")
		eligibilityChange(t, root, 5, "second-independent", "delta", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "independent", "second-independent")
	})
	t.Run("empty queue", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		if err := os.Remove(filepath.Join(root, "changes/0001-manual.md")); err != nil {
			t.Fatal(err)
		}
		readyEligibilityFixture(t, root)
		assertEligible(t, root)
	})
	t.Run("all waiting", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root)
	})
}

func Test_FLOW_069_EligibleWaitsForSameModule(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "same-module", "alpha", "", "")
	eligibilityChange(t, root, 3, "independent", "delta", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "independent")
}

func Test_FLOW_070_EligibleWaitsForDirectAndTransitiveDependents(t *testing.T) {
	root := eligibilityBase(t, "alpha")
	eligibilityChange(t, root, 2, "direct", "beta", "", "")
	eligibilityChange(t, root, 3, "transitive", "gamma", "", "")
	eligibilityChange(t, root, 4, "independent", "delta", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "independent")
}

func Test_FLOW_071_EligibleInterfaceChangeWaitsForBlockedDependent(t *testing.T) {
	root := eligibilityBase(t, "gamma")
	eligibilityChange(t, root, 2, "direct-interface", "beta", "", "Public beta API changes.\n")
	eligibilityChange(t, root, 3, "direct-internal", "beta", "", "")
	eligibilityChange(t, root, 4, "transitive-interface", "alpha", "", "Public alpha API changes.\n")
	eligibilityChange(t, root, 5, "transitive-internal", "alpha", "", "")
	readyEligibilityFixture(t, root)
	assertEligible(t, root, "direct-internal", "transitive-internal")
}

func Test_FLOW_072_EligibleUsesPrecedingAndOwnProjectedSet(t *testing.T) {
	t.Run("preceding set removes dependency", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "detach-beta", "beta", "set:\n  depends_on: []\n", "")
		eligibilityChange(t, root, 3, "gamma-after-detach", "gamma", "", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "detach-beta", "gamma-after-detach")
	})
	t.Run("own set adds or removes dependency", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "delta-now-dependent", "delta", "set:\n  depends_on: [alpha]\n", "")
		eligibilityChange(t, root, 3, "gamma-now-independent", "gamma", "set:\n  depends_on: []\n", "")
		readyEligibilityFixture(t, root)
		assertEligible(t, root, "gamma-now-independent")
	})
}

func Test_FLOW_011_StartSelectsFirstEligiblePastWaitingHead(t *testing.T) {
	makeQueue := func(t *testing.T) string {
		t.Helper()
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		eligibilityChange(t, root, 3, "first-eligible", "delta", "", "")
		eligibilityChange(t, root, 4, "later-eligible", "delta", "", "")
		readyEligibilityFixture(t, root)
		return root
	}
	t.Run("empty target selects first eligible", func(t *testing.T) {
		root := makeQueue(t)
		s, err := workflow.Start(root, "")
		if err != nil || s.Change.Slug != "first-eligible" {
			t.Errorf("Start(empty) selected %q: %v", s.Change.Slug, err)
		}
	})
	t.Run("explicit first eligible", func(t *testing.T) {
		root := makeQueue(t)
		s, err := workflow.Start(root, "first-eligible")
		if err != nil || s.Change.Slug != "first-eligible" {
			t.Errorf("Start(first-eligible) selected %q: %v", s.Change.Slug, err)
		}
	})
	for _, target := range []string{"waiting", "later-eligible"} {
		t.Run("reject "+target, func(t *testing.T) {
			root := makeQueue(t)
			if _, err := workflow.Start(root, target); err == nil || !strings.HasPrefix(err.Error(), "P01") {
				t.Errorf("Start(%s) = %v, want P01 rejection", target, err)
			}
		})
	}
	t.Run("no eligible change", func(t *testing.T) {
		root := eligibilityBase(t, "alpha")
		eligibilityChange(t, root, 2, "waiting", "beta", "", "")
		readyEligibilityFixture(t, root)
		if _, err := workflow.Start(root, ""); err == nil || !strings.HasPrefix(err.Error(), "P01") {
			t.Errorf("Start(empty) = %v, want P01 rejection", err)
		}
	})
}
package evals_test

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func costAt(t *testing.T, root, run, phase string) (workflow.CostSnapshot, bool) {
	t.Helper()
	path := filepath.Join(root, ".fireweed", "runs", run, phase, "costs.json")
	b, err := os.ReadFile(path)
	if err != nil {
		t.Errorf("cost snapshot %s: %v", path, err)
		return workflow.CostSnapshot{}, false
	}
	var snapshot workflow.CostSnapshot
	if err := json.Unmarshal(b, &snapshot); err != nil {
		t.Errorf("decode cost snapshot %s: %v", path, err)
		return workflow.CostSnapshot{}, false
	}
	return snapshot, true
}

func costFor(costs []workflow.ModuleCost, name string) (workflow.ModuleCost, bool) {
	for _, cost := range costs {
		if cost.Module == name {
			return cost, true
		}
	}
	return workflow.ModuleCost{}, false
}

func warningFor(ds []project.Diagnostic, code string) (project.Diagnostic, bool) {
	for _, d := range ds {
		if d.Code == code {
			return d, true
		}
	}
	return project.Diagnostic{}, false
}

func requireCostWarning(t *testing.T, ds []project.Diagnostic, code string, terms ...string) {
	t.Helper()
	d, ok := warningFor(ds, code)
	if !ok || !d.Warning {
		t.Errorf("%s warning absent from %#v", code, ds)
		return
	}
	for _, term := range terms {
		if !strings.Contains(d.Message, term) {
			t.Errorf("%s warning %q does not identify %q", code, d.Message, term)
		}
	}
}

func Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\n", "", "", "", "")
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
if [ "$1" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then failure=''; code=0
  else failure='<failure/>'; code=1; fi
  if [ -f packages/alpha/evals/second ]; then seconds=0.75
  else seconds=0.25; fi
  printf '<testsuite><testcase name="ALP-001" time="%s">%s</testcase><testcase name="ALP-002" time="0.1"/></testsuite>' "$seconds" "$failure" > "$2"
  exit "$code"
fi
printf '<testsuite><testcase name="BET-001" time="0.2"/></testsuite>' > "$2"
`)
	commit(t, root, "deterministic cost report")
	state := start(t, root)

codex
The transition and reporting code is in place. I’m reviewing failure paths now, especially whether a failed transition leaves the checkout intact and whether reports use the held snapshot or requested endpoint.
exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
base=Path('packages/workflow/internal/lifecycle')
p=base/'common.go'; s=p.read_text().replace('\"ls-tree\", \"-z\", \"--name-only\"', '\"ls-tree\", \"-r\", \"-z\", \"--name-only\"')
a=s.index('func splitNUL'); b=s.index('func selectChange', a)
s=s[:a]+'''func splitNUL(out []byte) []string {
	if len(out) == 0 {
		return nil
	}
	return strings.Split(strings.TrimSuffix(string(out), \"\\\\x00\"), \"\\\\x00\")
}

'''+s[b:]; p.write_text(s)
p=base/'queue.go'; s=p.read_text().replace('\"fmt\"\\n', '\"fmt\"\\n\\t\"path\"\\n').replace('earlier.Path[strings.LastIndex(earlier.Path, \"/\")+1:]', 'path.Base(earlier.Path)'); p.write_text(s)
p=base/'claim.go'; s=p.read_text()
s=s.replace('''		// A second Start from the owning checkout is an ownership error.
		if c, _, claimErr := heldClaim(repo, selector); claimErr == nil {
			return s, coded(\"P02\", c.Path, \"claim already held by this worktree\")
		}
''','')
s=s.replace('''	s.Snapshot, err = project.Project(root, chosen.Slug)
	if err "'!= nil {
		return s, staleError(err)
	}
	s.Snapshot.Root = root
'"''','')
start=s.index('\\tshared, local, err := baseRefs')
s=s[:start]+'''	s.Snapshot, err = project.Project(root, chosen.Slug)
	if err "'!= nil {
		return s, staleError(err)
	}
	s.Snapshot.Root = root

'"'''+s[start:]
s=s.replace('''	if attempt == nil {
		assigned, assignErr := assignIDs(root, chosen, s.Snapshot, queue, applied, string(original))''','''	assigned := string(original)
	if attempt == nil {
		var assignErr error
		assigned, assignErr = assignIDs(root, chosen, queue, applied, assigned)''')
s=s.replace('''	if err = installAssignedSpecs(root, queue, chosen, &s.Snapshot); err "'!= nil {
		return s, staleError(err)
	}'"''','''	if assigned "'!= string(original) {
		if err = installAssignedSpecs(root, queue, chosen, &s.Snapshot); err != nil {
			return s, staleError(err)
		}
	}'"''')
s=s.replace('func assignIDs(root string, ch project.Change, projection project.Snapshot, queue, archives []project.Change, source string)', 'func assignIDs(root string, ch project.Change, queue, archives []project.Change, source string)')
a=s.index('\\tafterModule, err := moduleNamed(projection'); b=s.index('\\treserved :=',a); s=s[:a]+s[b:]
s=s.replace('{before, after, ch.Added, ch.Removed}', '{before, ch.Added, ch.Removed}')
s=s.replace('''			for j, moved := range files[:i] {
				err = errors.Join(err, os.Rename(filepath.Join(backup, strconv.Itoa(j)), filepath.Join(s.Root, moved)))
			}
			return err''','''			for j, moved := range files[:i] {
				err = combine(err, os.Rename(filepath.Join(backup, strconv.Itoa(j)), filepath.Join(s.Root, moved)))
			}
			return err''')
p.write_text(s)
p=base/'costs.go'; s=p.read_text().replace('''		switch zone(snapshot, m, name) {
''','''		kind := zone(snapshot, m, name)
		switch kind {
''').replace('if zone(snapshot, m, name) == \"evals\" {','if kind == \"evals\" {'); p.write_text(s)
p=base/'pr.go'; s=p.read_text().replace('renderChangeHistory(&body, root, endpoint, ch, snapshot, touched[filename], diffs)', 'renderChangeHistory(&body, root, endpoint, ch, touched[filename], diffs)').replace('ch project.Change, snapshot project.Snapshot, touched []revision', 'ch project.Change, touched []revision'); p.write_text(s)
PY
gofmt -w packages/workflow/internal/lifecycle/*.go
rg -n 'type State|func verify|func stageWorktree|func readClaim|func assignIDs' packages/workflow/{api.go,model/model.go,internal/lifecycle/*.go}" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 1433ms:
packages/workflow/api.go:11:type State = model.State
packages/workflow/model/model.go:10:type State struct {
packages/workflow/internal/lifecycle/verify.go:29:func verify(ctx context.Context, request model.Request) (result model.Result, step verification, err error) {
packages/workflow/internal/lifecycle/claim.go:45:func readClaim(root, oid string) (claim, error) {
packages/workflow/internal/lifecycle/claim.go:421:func assignIDs(root string, ch project.Change, queue, archives []project.Change, source string) (string, error) {

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
base = Path('packages/workflow/internal/lifecycle')
p = base/'queue.go'; s = p.read_text().replace('''	for _, spec := range stale {
		if "'!spec.Recovered {
			return coded("P05", ch.Path, fmt.Sprintf("stale base for %s: recorded %s, current %s", spec.ID, spec.RecordedHash, spec.Current))
		}
	}
'"''','''	if len(stale) "'!= 0 {
		spec := stale[0]
		return coded("P05", ch.Path, fmt.Sprintf("stale base for %s (recorded %s)", spec.ID, spec.RecordedHash))
	}
'"'''); p.write_text(s)
p = base/'claim.go'; s = p.read_text().replace('''			s.Change = working
''','''			s.Change.Document = workerDocument(s.Change.Document, working.Document, s.Phase)
''',1)
i=s.index('func immutableBody')
s=s[:i]+'''func workerDocument(approved, working contracts.Document, phase string) contracts.Document {
	for _, name := range []string{\"Decisions\", \"Questions\", \"Audit\"} {
		if "'!editableSection(name, phase) {
			continue
		}
		if body, ok := contracts.Section(working.Body, name); ok {
			approved.Body = contracts.ReplaceSection(approved.Body, name, body)
		}
	}
	return approved
}

'"'''+s[i:]; p.write_text(s)
p = base/'attempt.go'; s = p.read_text(); i=s.index('\\tfor _, name := range []string{\"Decisions\"'); j=s.index('\\n}', i); s=s[:i]+'''	return contracts.RenderDocument(workerDocument(current.Document, saved.Document, phase))'''+s[j:]; p.write_text(s)
p = base/'costs.go'; s = p.read_text().replace('modules []project.Module) ([]model.ModuleCost', 'modules []project.Module, restrict bool) ([]model.ModuleCost').replace('if _, ok := known[cost.Module]; ok {','if _, ok := known[cost.Module]; ok || "'!restrict {'"'); p.write_text(s)
p = base/'status.go'; s = p.read_text().replace('latestCosts(root, snapshot.Modules)', 'latestCosts(root, snapshot.Modules, false)'); p.write_text(s)
p = base/'pr.go'; s = p.read_text().replace('latestCosts(root, snapshot.Modules)', 'latestCosts(root, snapshot.Modules, true)')
s=s.replace('''	left, right, ok := strings.Cut(strings.TrimSpace(revisionRange), \"..\")
	if "'!ok || left == "" || right == "" || strings.HasPrefix(right, ".") || strings.Contains(right, "..") {'"''','''	left, right, ok := strings.Cut(strings.TrimSpace(revisionRange), \"..\")
	separator := \"..\"
	if strings.HasPrefix(right, \".\") {
		separator = \"...\"
		right = strings.TrimPrefix(right, \".\")
	}
	if "'!ok || left == "" || right == "" || strings.HasPrefix(right, ".") || strings.Contains(right, "..") {'"''')
s=s.replace('revisions(root, base+\"..\"+endpoint)', 'revisions(root, base+separator+endpoint)').replace('Range "'`%s..%s`'"', 'Range "'`%s%s%s`'"').replace('changes\\\\n\", base, endpoint)', 'changes\\\\n\", base, separator, endpoint)')
s=s.replace('''		if path.Dir(doc.Path) == pendingDir {
			applyZoneSet(&m, doc)
		}''','''		if path.Dir(doc.Path) == pendingDir {
			queue, err := committedChanges(root, rev.oid, pendingDir)
			if err "'!= nil {
				return err
			}
			for _, queued := range queue {
				if queued.Module == m.Name {
					applyZoneSet(&m, queued)
				}
				if queued.Path == doc.Path {
					break
				}
			}
		}'"''')
p.write_text(s)
p = base/'audit.go'; s = p.read_text().replace('''(["'^'"\\\\s:]+):\\\\s*\"(.+?)\"\\\\s*(.*)"'$'"''','''(["'^'"\\\\s:]+):\\\\s*\"(.+)\"(?:\\\\s*(?:—|–|-)\\\\s*.*)?"'$'"'''); p.write_text(s)
PY
rg -n 'Range|findingLine|Recovered|latestCosts' packages/workflow/internal/lifecycle
gofmt -w packages/workflow/internal/lifecycle/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 580ms:
packages/workflow/internal/lifecycle/costs.go:135:func latestCosts(root string, modules []project.Module, restrict bool) ([]model.ModuleCost, []project.Diagnostic, error) {
packages/workflow/internal/lifecycle/status.go:124:	costs, ds, err := latestCosts(root, snapshot.Modules, false)
packages/workflow/internal/lifecycle/pr.go:25:func PRBody(root, revisionRange string) (report model.PRBodyReport, err error) {
packages/workflow/internal/lifecycle/pr.go:31:	left, right, ok := strings.Cut(strings.TrimSpace(revisionRange), "..")
packages/workflow/internal/lifecycle/pr.go:38:		return report, coded("P01", revisionRange, fmt.Sprintf("expected an explicit commit range, got %q", revisionRange))
packages/workflow/internal/lifecycle/pr.go:42:		return report, withCode("P01", revisionRange, fmt.Errorf("resolve range %q: %w", revisionRange, err))
packages/workflow/internal/lifecycle/pr.go:46:		return report, withCode("P01", revisionRange, fmt.Errorf("resolve range %q: %w", revisionRange, err))
packages/workflow/internal/lifecycle/pr.go:83:	fmt.Fprintf(&body, "# Fireweed changes\n\nRange `%s%s%s`\n\n## Selected changes\n", base, separator, endpoint)
packages/workflow/internal/lifecycle/pr.go:116:	costs, diagnostics, err := latestCosts(root, snapshot.Modules, true)
packages/workflow/internal/lifecycle/audit.go:25:var findingLine = regexp.MustCompile(`^[-*]\s+(?:(\[untestable\])\s+)?([^\s:]+):\s*"(.+)"(?:\s*(?:—|–|-)\s*.*)?$`)
packages/workflow/internal/lifecycle/audit.go:28:	parts := findingLine.FindStringSubmatch(line)

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
base=Path('packages/workflow/internal/lifecycle')
p=base/'verify.go'; s=p.read_text()
a=s.index('\\tchanges, err := gitstore.Changes'); b=s.index('\\n\\titems, err := specs(m)',a)
segment=s[a:b]
s=s[:a]+'''	result.Diagnostics, err = verifyZones(ctx, state, m)
	if err "'!= nil {
		return result, step, err
	}
	evalFix := stringList(state.Change.Document.Frontmatter["eval_fix"])
'"'''+s[b:]
segment=segment.replace('return result, step, err', 'return nil, err')
segment=segment.replace('''	known := map[string]bool{}''','''	known := map[string]bool{}''').replace('''	evalFix := stringList''','''	phaseChanges := changes[:0]
	for _, change := range changes {
		if "'!strings.HasPrefix(change.Path, ".fireweed/runs/") || change.OldPath != "" {
			phaseChanges = append(phaseChanges, change)
		}
	}
	changes = phaseChanges
	var diagnostics []project.Diagnostic
	evalFix := stringList'"''')
segment=segment.replace('result.Diagnostics = append(result.Diagnostics,', 'diagnostics = append(diagnostics,')
segment=segment[:segment.index('\\tif err = firstError')]+'''	return diagnostics, firstError(diagnostics)
'''
i=s.index('func combine'); s=s[:i]+'''func verifyZones(ctx context.Context, state model.State, m project.Module) ([]project.Diagnostic, error) {
	root := state.Snapshot.Root
'''+segment+'''}

'''+s[i:]
s=s.replace('''		if costErr "'!= nil {
			return result, step, combine(verificationErr, costErr)
		}
		result.Costs = append(result.Costs, cost)
		if err = writeJSON(filepath.Join(runDir, module.Name, "report.json"), report); err != nil {
			return result, step, combine(verificationErr, err)
		}'"''','''		if costErr == nil {
			result.Costs = append(result.Costs, cost)
		}
		verificationErr = combine(verificationErr, costErr)
		verificationErr = combine(verificationErr, writeJSON(filepath.Join(runDir, module.Name, \"report.json\"), report))''')
s=s.replace('''			result = append(result, item)
		}
	}
	return result
}

func enforcedSpecs''','''			// Review items do not require coverage but a mapped failure
			// in the current module still needs to be reported.
			item.Tag = \"\"
			result = append(result, item)
		}
	}
	return result
}

func enforcedSpecs''',1)
p.write_text(s)
p=base/'done.go'; s=p.read_text()
needle='''	ch.Document.Frontmatter[\"phase\"] = next
'''
s=s.replace(needle,'''	if state.Phase "'!= "audit" {
		m, err := moduleNamed(state.Snapshot, ch.Module)
		if err != nil {
			return result, err
		}
		diagnostics, err := verifyZones(ctx, before, m)
		if err != nil {
			result.Diagnostics = append(result.Diagnostics, diagnostics...)
			return result, err
		}
	}
	ch.Document.Frontmatter["phase"] = next
'"''',1)
p.write_text(s)
p=base/'common.go'; s=p.read_text().replace('''	if strings.HasPrefix(err.Error(), code+\":\") {
		return err
	}
	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf(\"%s: %w\", code, err)}''','''	if "'!strings.HasPrefix(err.Error(), code+":") {
		return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s: %w", code, err)}
	}
	if existing, _, ok := contracts.ErrorCode(err); ok && existing == code {
		return err
	}
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}'"'''); p.write_text(s)
p=base/'claim.go'; s=p.read_text().replace('''	if err = gitstore.CASRef(root, c.ref(), \"\", oid); err "'!= nil {'"''','''	if err = requireClean(root, \"start\"); err "'!= nil {
		return s, err
	}
	if err = gitstore.CASRef(root, c.ref(), "", oid); err != nil {'"''',1)
p.write_text(s)
PY
gofmt -w packages/workflow/internal/lifecycle/*.go
sed -n '35,280p' packages/workflow/internal/lifecycle/verify.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 593ms:
	var oid string
	step.claim, oid, err = heldClaim(step.repo, request.Change)
	if err != nil {
		return result, step, err
	}
	result.State, err = step.claim.state(step.repo, oid)
	if err != nil {
		return result, step, err
	}
	state := result.State
	switch state.Phase {
	case "evals", "audit", "impl":
	default:
		return result, step, coded("P01", state.Change.Path, "unsupported phase "+state.Phase)
	}
	if state.Head != state.PhaseBase {
		return result, step, coded("P05", state.Change.Path, fmt.Sprintf("stale phase base: HEAD is %s, expected %s", state.Head, state.PhaseBase))
	}
	step.working, err = approvedWorking(state)
	if err != nil {
		return result, step, err
	}
	if err = checkStale(root, step.working); err != nil {
		return result, step, err
	}
	result.State.Change = step.working
	m, err := moduleNamed(state.Snapshot, state.Change.Module)
	if err != nil {
		return result, step, err
	}
	result.Diagnostics, err = verifyZones(ctx, state, m)
	if err != nil {
		return result, step, err
	}
	evalFix := stringList(state.Change.Document.Frontmatter["eval_fix"])

	items, err := specs(m)
	if err != nil {
		return result, step, err
	}
	if state.Phase == "audit" {
		step.audit, err = checkAudit(state, step.working, items)
		return result, step, err
	}
	if err = validateRun(state.Run); err != nil {
		return result, step, err
	}
	runDir := filepath.Join(root, ".fireweed", "runs", state.Run, state.Phase)
	result.Reports = make(map[string]checks.Report)
	defer func() {
		if len(result.Costs) == 0 {
			return
		}
		diagnostics := costWarnings(state.Snapshot, result.Costs)
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		snapshot := newCostSnapshot(state, result.Costs, diagnostics)
		err = combine(err, writeJSON(filepath.Join(runDir, "costs.json"), snapshot))
	}()
	fullImpl := state.Phase == "impl" && !request.Quick
	modules := []project.Module{m}
	if fullImpl {
		mode := stringValue(state.Snapshot.Config["verify_dependents"])
		if mode == "" {
			mode = "direct"
		}
		for _, dependent := range state.Snapshot.Modules {
			if dependent.Name != m.Name && mode != "none" && dependsOn(state.Snapshot, dependent.Name, m.Name, mode == "all") {
				modules = append(modules, dependent)
			}
		}
	}

	var verificationErr error
	for i, module := range modules {
		moduleItems := items
		if i != 0 {
			moduleItems, err = specs(module)
			if err != nil {
				return result, step, combine(verificationErr, err)
			}
		}
		report, evalErr := checks.Evaluate(ctx, checks.HookRequest{
			Snapshot: state.Snapshot, Module: module.Name, Bin: request.Bin, Timeout: request.Timeout,
		})
		if evalErr != nil {
			return result, step, combine(verificationErr, evalErr)
		}
		result.Reports[module.Name] = report
		cost, costErr := moduleCost(state.Snapshot, module, report)
		if costErr == nil {
			result.Costs = append(result.Costs, cost)
		}
		verificationErr = combine(verificationErr, costErr)
		verificationErr = combine(verificationErr, writeJSON(filepath.Join(runDir, module.Name, "report.json"), report))
		var reportErr error
		if i == 0 {
			reportErr = checks.RequireCoverage(report, enforcedSpecs(moduleItems))
		}
		if reportErr == nil {
			switch {
			case state.Phase == "evals" && len(evalFix) != 0:
				for _, c := range report.Cases {
					if c.Result == "fail" || c.Result == "error" {
						reportErr = coded("E01", module.Name, fmt.Sprintf("failed case %s: %s", c.Name, c.Failure))
						break
					}
				}
				if reportErr == nil {
					listed := make([]contracts.Item, 0, len(evalFix))
					for _, id := range evalFix {
						listed = append(listed, contracts.Item{ID: id})
					}
					reportErr = checks.RequireCoverage(report, listed)
				}
			case state.Phase == "evals":
				reportErr = checks.RequireFailFirst(report, changeIDs(state.Change, moduleItems, false), stringList(state.Change.Document.Frontmatter["already_satisfied"]))
			case i == 0:
				reportErr = checks.RequirePass(report, nonQuarantinedSpecs(moduleItems))
			default:
				reportErr = checks.RequirePass(report, enforcedSpecs(moduleItems))
			}
		}
		verificationErr = combine(verificationErr, reportErr)
	}
	if fullImpl {
		diagnostics, boundaryErr := checks.Boundaries(ctx, checks.BoundaryRequest{
			Snapshot: state.Snapshot, Bin: request.Bin, Timeout: request.Timeout,
		})
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		if boundaryErr == nil {
			boundaryErr = firstError(diagnostics)
		} else {
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && diagnostic.Code != "" {
					boundaryErr = withCode(diagnostic.Code, diagnostic.Path, boundaryErr)
					break
				}
			}
		}
		verificationErr = combine(verificationErr, boundaryErr)
	}
	return result, step, verificationErr
}

func verifyZones(ctx context.Context, state model.State, m project.Module) ([]project.Diagnostic, error) {
	root := state.Snapshot.Root
	changes, err := gitstore.Changes(root, state.PhaseBase)
	if err != nil {
		return nil, err
	}
	// An index entry may differ even when its working copy matches the base.
	index, err := diffPaths(ctx, root, "diff", "--cached", "--name-status", "-z", "--no-renames", state.PhaseBase, "--")
	if err != nil {
		return nil, err
	}
	known := map[string]bool{}
	for _, change := range changes {
		known[change.Path] = true
		known[change.OldPath] = true
	}
	for _, change := range index {
		if !known[change.Path] {
			changes = append(changes, change)
		}
	}
	phaseChanges := changes[:0]
	for _, change := range changes {
		if !strings.HasPrefix(change.Path, ".fireweed/runs/") || change.OldPath != "" {
			phaseChanges = append(phaseChanges, change)
		}
	}
	changes = phaseChanges
	var diagnostics []project.Diagnostic
	evalFix := stringList(state.Change.Document.Frontmatter["eval_fix"])
	if len(evalFix) != 0 {
		for _, change := range changes {
			for _, file := range []string{change.Path, change.OldPath} {
				if file == "" || file == state.Change.Path || state.Phase == "evals" && belongsTo(state.Snapshot, m, file) && zone(state.Snapshot, m, file) == "evals" {
					continue
				}
				diagnostics = append(diagnostics, project.Diagnostic{
					Code: "Z01", Path: file, Message: fmt.Sprintf("%s is forbidden during eval_fix %s", file, state.Phase),
				})
			}
		}
	}
	diagnostics = append(diagnostics, checks.Zones(checks.ZoneRequest{
		Snapshot: state.Snapshot, Module: m.Name, Phase: state.Phase,
		ChangePath: state.Change.Path, Base: state.PhaseBase, Changes: changes,
		InterfaceDeclared: strings.TrimSpace(section(state.Change.Document, "Interface")) != "",
	})...)
	return diagnostics, firstError(diagnostics)
}

func combine(first, second error) error {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}
	return errors.Join(first, second)
}

func quarantineReason(item contracts.Item) (string, bool) {
	tag := strings.TrimSpace(item.Tag)
	if tag == "quarantine" {
		return "", true
	}
	reason, ok := strings.CutPrefix(tag, "quarantine:")
	return strings.TrimSpace(reason), ok
}

func nonQuarantinedSpecs(items []contracts.Item) []contracts.Item {
	var result []contracts.Item
	for _, item := range items {
		if _, quarantined := quarantineReason(item); !quarantined {
			// Review items do not require coverage but a mapped failure
			// in the current module still needs to be reported.
			item.Tag = ""
			result = append(result, item)
		}
	}
	return result
}

func enforcedSpecs(items []contracts.Item) []contracts.Item {
	var result []contracts.Item
	for _, item := range items {
		if _, quarantined := quarantineReason(item); !quarantined && item.Tag != "review" {
			result = append(result, item)
		}
	}
	return result
}

func changeIDs(ch project.Change, items []contracts.Item, audit bool) []string {
	enforced := map[string]bool{}
	for _, item := range enforcedSpecs(items) {
		enforced[item.ID] = true
	}
	selected := map[string]bool{}
	for _, edits := range [][]contracts.Item{ch.Added, ch.Modified} {
		for _, item := range edits {
			if enforced[item.ID] {
				selected[item.ID] = true

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
base=Path('packages/workflow/internal/lifecycle')
p=base/'claim.go'; s=p.read_text()
s=s.replace('''	root = repo.Root
	if err = requireClean(root, \"start\"); err "'!= nil {'"''','''	root = repo.Root
	if selector "'!= "" {
		paths, err := changePaths(root, pendingDir)
		if err != nil {
			return s, err
		}
		if _, err := selectChange(paths, selector); err != nil {
			return s, err
		}
	}
	if err = requireClean(root, "start"); err != nil {'"''',1)
p.write_text(s)
p=base/'verify.go'; s=p.read_text()
s=s.replace('''					if c.Result == \"fail\" || c.Result == \"error\" {
						reportErr = coded(\"E01\", module.Name, fmt.Sprintf(\"failed case %s: %s\", c.Name, c.Failure))''','''					if c.Result "'!= "pass" {
						reportErr = coded("E01", module.Name, fmt.Sprintf("failed case %s (%s): %s", c.Name, c.Result, c.Failure))'"''')
s=s.replace('''		if "'!strings.HasPrefix(change.Path, ".fireweed/runs/") || change.OldPath != "" {'"''','''		if "'!strings.HasPrefix(change.Path, ".fireweed/runs/") || change.OldPath != "" || existsAt(root, state.PhaseBase, change.Path) {'"''')
p.write_text(s)
p=base/'attempt.go'; s=p.read_text()
s=s.replace('restoreAttempt(root, oid, file string)', 'restoreAttempt(root, base, oid, file string)')
s=s.replace('''	_, err = git(ctx, root, nil, patch, \"apply\", \"--whitespace=nowarn\", \"-\")
	return withCode(\"P01\", file, err)''','''	if _, err = git(ctx, root, nil, patch, \"apply\", \"--whitespace=nowarn\", \"-\"); err == nil {
		return nil
	}

	// A later commit can edit the same file without changing the saved
	// hunk. Merge through a private index, then apply the resulting delta
	// to the checkout with the real index left untouched.
	idx, cleanup, err := temporaryIndex(ctx, root, base)
	if err "'!= nil {
		return withCode("P01", file, err)
	}
	defer cleanup()
	if _, err = idx.run(patch, "apply", "--cached", "--3way", "--whitespace=nowarn", "-"); err != nil {
		return withCode("P01", file, err)
	}
	tree, err := idx.run(nil, "write-tree")
	if err != nil {
		return withCode("P01", file, err)
	}
	merged, err := git(ctx, root, nil, nil, "diff", "--binary", "--full-index", "--no-renames", base, strings.TrimSpace(string(tree)), "--")
	if err != nil || len(merged) == 0 {
		return withCode("P01", file, err)
	}
	_, err = git(ctx, root, nil, merged, "apply", "--whitespace=nowarn", "-")
	return withCode("P01", file, err)'"''')
p.write_text(s)
p=base/'claim.go'; s=p.read_text().replace('restoreAttempt(root, savedOID, chosen.Path)', 'restoreAttempt(root, s.PhaseBase, savedOID, chosen.Path)'); p.write_text(s)
PY
gofmt -w packages/workflow/internal/lifecycle/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 569ms:
codex
The review caught two edge cases I’ve addressed. Resume can merge saved edits with later independent changes to the same file, and Done checks the phase zones again after hooks before committing. I’ll record the remaining specification gaps in DRILL_NOTES.md.
exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
p = Path('packages/workflow/internal/lifecycle/done.go')
s = p.read_text()
a = s.index('func decisionText(')
b = s.index('type fileEdits', a)
s = s[:a]+'''func decisionEntries(body string) []string {
	var entries, lines []string
	flush := func() {
		if len(lines) "'!= 0 {
			entries = append(entries, strings.Join(lines, "'"\\\\n\"))
			lines = nil
		}
	}
	for _, line := range strings.Split(body, \"\\\\n\") {
		trimmed := strings.TrimSpace(line)
		if trimmed == \"\" {
			flush()
			continue
		}
		if strings.HasPrefix(line, \"- \") || strings.HasPrefix(line, \"* \") || strings.HasPrefix(line, \"[\") {
			flush()
		}
		lines = append(lines, line)
	}
	flush()
	return entries
}

func decisionText(entry string) string {
	var lines []string
	for _, line := range strings.Split(entry, \"\\\\n\") {
		text := strings.TrimSpace(line)
		text = strings.TrimPrefix(text, \"- \")
		text = strings.TrimPrefix(text, \"* \")
		text = strings.TrimPrefix(text, \"> \")
		lines = append(lines, text)
	}
	return strings.Trim(strings.Join(lines, \"\\\\n\"), \"\\\\\"\")
}

func applyDecisions(prior, added, removed, slug string) string {
	removals := map[string]bool{}
	for _, entry := range decisionEntries(removed) {
		removals[decisionText(entry)] = true
	}
	var kept []string
	for _, entry := range decisionEntries(prior) {
		text := decisionText(entry)
		if removals[text] {
			continue
		}
		if strings.HasPrefix(entry, \"[\") {
			if _, rest, ok := strings.Cut(text, \"] \"); ok && removals[rest] {
				continue
			}
		}
		if strings.HasPrefix(entry, \"- \") {
			if _, rest, ok := strings.Cut(text, \": \"); ok && removals[rest] {
				continue
			}
		}
		kept = append(kept, entry)
	}
	for _, entry := range decisionEntries(added) {
		lines := strings.Split(decisionText(entry), \"\\\\n\")
		kept = append(kept, \"- \"+slug+\": \"+strings.Join(lines, \"\\\\n  \"))
	}
	return \"\\\\n\"+strings.Join(kept, \"\\\\n\\\\n\")+\"\\\\n\"
}

'''+s[b:]
p.write_text(s)
PY
cat > DRILL_NOTES.md <<'EOF'
## Contradicted Decisions

None found in the retained workflow Decisions.

## Gaps

- FLOW-001 and FLOW-019 · workflow. The claim payload has no specified schema. The rebuild stores the canonical worktree, change identity, phase base, approved document and projected snapshot in a JSON blob addressed by the claim ref.
- FLOW-020 and FLOW-075 · workflow. The spec does not enumerate worker-editable body sections. The retained evals establish Decisions and Questions for evals and Decisions for impl, and Audit for audit. The rebuild also treats Questions as editable during impl.
- FLOW-059 · workflow. The spec gives no full audit submission grammar. The rebuild uses the checked-ID line and quoted finding bullets from the retained evals, with an optional untestable marker.
- FLOW-074 · workflow. The spec does not define how to resolve overlapping edits in a saved attempt and later commits. The rebuild applies the saved delta, tries Git's three-way merge when needed, and returns an error for unresolved conflicts.
- FLOW-117 · workflow. The spec does not define file classification for historical commits that predate loadable project metadata. The rebuild retains their recorded findings and result counts and classifies files for commits with loadable metadata.
EOF
gofmt -w packages/workflow/internal/lifecycle/*.go
gofmt -l packages/workflow/internal/lifecycle/*.go
rg --files packages/workflow/internal DRILL_NOTES.md | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 678ms:
DRILL_NOTES.md
packages/workflow/internal/lifecycle/attempt.go
packages/workflow/internal/lifecycle/audit.go
packages/workflow/internal/lifecycle/claim.go
packages/workflow/internal/lifecycle/common.go
packages/workflow/internal/lifecycle/context.go
packages/workflow/internal/lifecycle/costs.go
packages/workflow/internal/lifecycle/done.go
packages/workflow/internal/lifecycle/git.go
packages/workflow/internal/lifecycle/history.go
packages/workflow/internal/lifecycle/pr.go
packages/workflow/internal/lifecycle/queue.go
packages/workflow/internal/lifecycle/status.go
packages/workflow/internal/lifecycle/verify.go

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
base=Path('packages/workflow/internal/lifecycle')
p=base/'claim.go'; s=p.read_text()
s=s.replace('''		if section(ch.Document, name) "'!= section(working.Document, name) {'"''','''		if "'!reflect.DeepEqual(nonblankLines(section(ch.Document, name)), nonblankLines(section(working.Document, name))) {'"''')
a=s.index('func immutableBody')
b=s.index('func archivePath',a)
s=s[:a]+'''func immutableBody(body, phase string) string {
	var sections []string
	var lines []string
	keep := true
	flush := func() {
		for len(lines) "'!= 0 && strings.TrimSpace(lines[0]) == "" {
			lines = lines[1:]
		}
		for len(lines) != 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
			lines = lines[:len(lines)-1]
		}
		if len(lines) != 0 {
			sections = append(sections, strings.Join(lines, "'"\\\\n\"))
		}
		lines = nil
	}
	for _, line := range strings.Split(body, \"\\\\n\") {
		if name, ok := strings.CutPrefix(line, \"## \"); ok {
			flush()
			name = strings.TrimSpace(name)
			keep = "'!editableSection(name, phase) && name != "Audit" && name != "Provenance"
		}
		if keep {
			lines = append(lines, line)
		}
	}
	flush()
	return strings.Join(sections, "'"\\\\n\\\\n\")
}

'''+s[b:]
p.write_text(s)
p=base/'done.go'; s=p.read_text()
a=s.index('func appendProvenance')
b=s.index('func applyToModule',a)
s=s[:a]+'''func nonblankLines(body string) []string {
	var lines []string
	for _, line := range strings.Split(body, \"\\\\n\") {
		if strings.TrimSpace(line) "'!= "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func appendProvenance(body, line string) string {
	prior, _ := contracts.Section(body, "Provenance")
	return contracts.ReplaceSection(body, "Provenance", "'"\\\\n\"+strings.Join(append(nonblankLines(prior), line), \"\\\\n\")+\"\\\\n\")
}

'''+s[b:]
p.write_text(s)
PY
gofmt -w packages/workflow/internal/lifecycle/*.go
gofmt -l packages/workflow/internal/lifecycle/*.go
wc -l DRILL_NOTES.md packages/workflow/internal/lifecycle/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 581ms:
      11 DRILL_NOTES.md
     146 packages/workflow/internal/lifecycle/attempt.go
     117 packages/workflow/internal/lifecycle/audit.go
     699 packages/workflow/internal/lifecycle/claim.go
     428 packages/workflow/internal/lifecycle/common.go
      74 packages/workflow/internal/lifecycle/context.go
     197 packages/workflow/internal/lifecycle/costs.go
     343 packages/workflow/internal/lifecycle/done.go
     147 packages/workflow/internal/lifecycle/git.go
     338 packages/workflow/internal/lifecycle/history.go
     365 packages/workflow/internal/lifecycle/pr.go
      88 packages/workflow/internal/lifecycle/queue.go
     215 packages/workflow/internal/lifecycle/status.go
     297 packages/workflow/internal/lifecycle/verify.go
    3465 total

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
p=Path('packages/workflow/internal/lifecycle/claim.go')
s=p.read_text().replace('if field(chosen, \"run\") == \"\" {\\n\\t\\tvar id [16]byte', 'if unstarted || field(chosen, \"run\") == \"\" {\\n\\t\\tvar id [16]byte')
p.write_text(s)
p=Path('DRILL_NOTES.md')
s=p.read_text().replace('The claim payload has no specified schema.', 'The claim payload format is left to the implementation.').replace('The spec does not enumerate worker-editable body sections.', 'The spec leaves the complete list of worker-editable body sections open.').replace('The spec gives no full audit submission grammar.', 'The spec leaves the audit submission grammar open.').replace('The spec does not define how to resolve overlapping edits in a saved attempt and later commits.', 'The spec leaves conflict resolution open when a saved attempt overlaps later commits.').replace('The spec does not define file classification for historical commits that predate loadable project metadata.', 'The spec leaves file classification open for historical commits that predate loadable project metadata.')
p.write_text(s)
PY
gofmt -w packages/workflow/internal/lifecycle/claim.go
gofmt -l packages/workflow/internal/lifecycle/*.go
rg -n '"'^## '"' DRILL_NOTES.md
rg -n '"'^func (Start|StartFromAttempt|Eligible|Current|Commits|Abandon|Release|Verify|Done|Block|Context|Costs|Status|PRBody|IsEndpointConfigError|GC)'"\\(' packages/workflow/internal/lifecycle" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 599ms:
1:## Contradicted Decisions
5:## Gaps
packages/workflow/internal/lifecycle/done.go:21:func Done(ctx context.Context, request model.Request) (result model.Result, err error) {
packages/workflow/internal/lifecycle/verify.go:24:func Verify(ctx context.Context, request model.Request) (model.Result, error) {
packages/workflow/internal/lifecycle/status.go:38:func Status(root string) (report model.StatusReport, err error) {
packages/workflow/internal/lifecycle/status.go:130:func GC(ctx context.Context, request model.GCRequest) (result model.GCResult, err error) {
packages/workflow/internal/lifecycle/history.go:230:func Commits(root, selector string) ([]string, error) {
packages/workflow/internal/lifecycle/history.go:250:func Abandon(root, selector string) (string, error) {
packages/workflow/internal/lifecycle/claim.go:157:func Current(root, selector string) (model.State, error) {
packages/workflow/internal/lifecycle/claim.go:178:func Release(root, selector string) error {
packages/workflow/internal/lifecycle/claim.go:190:func Start(root, selector string) (model.State, error) {
packages/workflow/internal/lifecycle/claim.go:194:func StartFromAttempt(root, selector string, attempt int) (model.State, error) {
packages/workflow/internal/lifecycle/context.go:15:func Context(root, selector string, full bool) (string, error) {
packages/workflow/internal/lifecycle/pr.go:20:func IsEndpointConfigError(err error) bool {
packages/workflow/internal/lifecycle/pr.go:25:func PRBody(root, revisionRange string) (report model.PRBodyReport, err error) {
packages/workflow/internal/lifecycle/costs.go:105:func Costs(root string) ([]model.CostSnapshot, error) {
packages/workflow/internal/lifecycle/attempt.go:15:func Block(request model.BlockRequest) (result model.BlockResult, err error) {
packages/workflow/internal/lifecycle/queue.go:52:func Eligible(root string) ([]project.Change, error) {

exec
/bin/zsh -lc "python3 - <<'PY'
from pathlib import Path
p=Path('packages/workflow/internal/lifecycle/verify.go')
s=p.read_text().replace('reportErr = checks.RequirePass(report, nonQuarantinedSpecs(moduleItems))', 'reportErr = requireMappedPass(report, nonQuarantinedSpecs(moduleItems))').replace('reportErr = checks.RequirePass(report, enforcedSpecs(moduleItems))', 'reportErr = requireMappedPass(report, enforcedSpecs(moduleItems))')
a=s.index('func quarantineReason')
s=s[:a]+'''func requireMappedPass(report checks.Report, items []contracts.Item) error {
	selected := make(map[string]bool, len(items))
	for _, item := range items {
		selected[item.ID] = true
	}
	var failures checks.Report
	for _, testcase := range report.Cases {
		if testcase.Result "'!= "fail" && testcase.Result != "error" {
			continue
		}
		var ids []string
		for _, id := range testcase.IDs {
			if selected[id] {
				ids = append(ids, id)
			}
		}
		if len(ids) != 0 {
			testcase.IDs = ids
			failures.Cases = append(failures.Cases, testcase)
		}
	}
	if len(failures.Cases) == 0 {
		return nil
	}
	return checks.RequirePass(failures, items)
}

'"'''+s[a:]
p.write_text(s)
PY
gofmt -w packages/workflow/internal/lifecycle/verify.go
gofmt -l packages/workflow/internal/lifecycle/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 553ms:
codex
Rebuilt the workflow implementation under `packages/workflow/internal/lifecycle` and documented the implementation choices and spec gaps in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/DRILL_NOTES.md).

`gofmt` passed. Validation was limited to formatting and source review because `SYSTEM.md` reserves builds and evals for the separate verifier. No commit was created.
tokens used
207,742


### Attempt 2 (exit 0)
Updated the disposable workflow code to keep retained run output out of Done and Block commits, support GC on older Git versions, and preserve saved budget warnings that point to files within a module.

Recorded the spec gaps in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/DRILL_NOTES.md). Retained inputs were unchanged, and no commit was made.

Formatting checks passed. The workflow, app, and runner evals still need to run in the external verifier.

Reading additional input from stdin...
2026-11-09T00:06:13.982860Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1bda0-7c9a-79d0-a12c-ae167e5d309c
--------
user
# Rebuild workflow

Target spec and Decisions:
---
depends_on:
    - contracts
    - gitstore
    - assets
    - project
    - checks
id_prefix: FLOW
interface:
    - api.go
    - model/**
name: workflow
---

# workflow

## Specs

### FLOW-001
Start MUST atomically create the local claim for the selected queued change.

### FLOW-002
Verify MUST reject edits forbidden by the approved projected phase zones before running hooks, checking differences since the held phase base.

### FLOW-003
Done MUST require full verification to succeed before committing a phase transition.

### FLOW-004
Successful impl Done MUST apply Added, Modified and Removed spec edits to the module's SPEC.md.

### FLOW-005
Context MUST include the change, projected spec, interface contents, SYSTEM and matching playbook. It MUST list other relevant files by path and size, or include their contents when full is requested.

### FLOW-006
Start MUST include the approved projected project snapshot in its returned State.

### FLOW-007
Start MUST replace each local Added handle throughout the change with the next
available spec ID for its module. IDs of current and previously removed specs
MUST remain reserved, including specs Removed by the starting change.

### FLOW-008
Start MUST assign a unique run identifier when starting an unstarted change.

### FLOW-009
Start MUST report P04 as a warning when neither the local base nor the shared base contains the change file.

### FLOW-010
Start MUST reject a worktree containing uncommitted changes.

### FLOW-011
Start MUST reject a requested change unless it is the first change returned by
Eligible. With no requested target, Start MUST start that first eligible
change. When Eligible returns no change, Start MUST reject with P01.

### FLOW-012
Start MUST return the canonical absolute worktree root in State.Snapshot.Root.

### FLOW-013
Start MUST set phase to evals when starting an unstarted spec change.

### FLOW-014
Start MUST leave HEAD unchanged.

### FLOW-015
Start MUST reject with P07 when the shared-base unapproved archive count reaches authoring.max_unapproved, whose default is 5.

### FLOW-016
Start MUST reject a competing worktree that attempts to claim the same change.

### FLOW-017
Current MUST resolve only a claim owned by the calling worktree.

### FLOW-018
Start MUST use the current HEAD as the initial phase base of its claim.

### FLOW-019
While a claim is held, Current MUST return the project snapshot approved at Start.

### FLOW-020
Verify MUST reject an unapproved alteration of the active change's author fields before running a hook.

### FLOW-021
Verify MUST reject an unapproved alteration of the active change's recorded runtime state before running a hook.

### FLOW-022
Whenever Verify runs the target module's evals it MUST require coverage of every enforced current spec.

### FLOW-023
During evals Verify for spec and fix changes MUST require genuine fail-first
evidence for every enforced change ID that is not already_satisfied.
Review-tagged specs MUST be excluded from fail-first requirements.

### FLOW-024
During impl Verify MUST reject a failing current-module testcase that references a non-quarantined spec.

### FLOW-025
Full impl Verify MUST use verify_dependents with a default of direct to select dependent modules.

### FLOW-026
Full impl Verify MUST reject a failed enforced testcase from a selected dependent module.

### FLOW-027
Full impl Verify MUST require boundary verification to succeed.

### FLOW-028
Verify MUST select hooks from the held approved projected project snapshot.

### FLOW-029
Quick Verify MUST skip dependent-module evals.

### FLOW-030
Quick Verify MUST skip boundary verification.

### FLOW-031
Verify MUST leave HEAD unchanged.

### FLOW-032
Done MUST commit a successfully verified evals phase and advance it to audit when audit is enabled.

### FLOW-033
Done MUST commit a successfully verified evals phase and advance it to impl when audit.enabled is false.

### FLOW-034
Audit Done MUST discard a finding whose spec ID is absent or whose quoted clause is not an exact substring of that spec.

### FLOW-035
On re-audit, Audit Done MUST discard findings for spec IDs that had no retained finding in the preceding audit round.

### FLOW-036
Audit Done MUST require every enforced ID in the change's Added, Modified,
fixes or eval_fix entries to appear in a No findings submission's checked-ID list.

### FLOW-037
Retained testable audit findings within the round limit MUST return the change to evals.

### FLOW-038
An audit return to evals MUST increment and persist audit_rounds.

### FLOW-039
Audit Done MUST block the change when retained findings would exceed audit.max_rounds, whose default is 2.

### FLOW-040
Audit Done MUST block the change when a retained finding marks its quoted clause untestable.

### FLOW-041
Audit Done MUST require No findings with a complete checked-ID list before advancing to impl.

### FLOW-042
When audit blocks a change, Done MUST include the blocking findings in its Questions section.

### FLOW-043
Audit blocking MUST release the calling worktree's held change claim.

### FLOW-044
Successful impl Done MUST apply the change's set values to module metadata.

### FLOW-045
Successful impl Done MUST apply Decisions attributed to the change slug and remove entries quoted in Removed decisions.

### FLOW-046
Each successful Done MUST advance the phase base to its new commit.

### FLOW-047
Successful impl Done MUST release only the exact held claim upon application.

### FLOW-048
Failed Done verification MUST leave the phase unadvanced and create no commit.

### FLOW-049
Impl Done MUST require the applied project to pass static validation before committing.

### FLOW-050
Successful impl Done MUST move the active change to the archive with phase applied and commit it together with the application.

### FLOW-051
Each successful Done MUST append one provenance line containing phase, date,
manual or configured agent, version, playbook hash and result counts.
It MUST preserve existing nonblank provenance lines unchanged and in order.
It MAY normalize blank lines within the Provenance section.

### FLOW-052
Successful Done MUST retain verification reports and process output under the run directory.

### FLOW-053
State MUST expose the held projected snapshot so composing commands use the approved graph during hooks.

### FLOW-054
Start MUST prefix dirty-worktree and queue-eligibility rejections with P01.

### FLOW-055
Verify and Done MUST prefix unsupported-phase rejections with P01.

### FLOW-056
Start and Current MUST prefix existing claim ownership or contention failures with P02.

### FLOW-057
Verify and Done MUST prefix unapproved author, runtime or audit-state alteration errors with P03 before running hooks.

### FLOW-058
Start, Verify and Done MUST prefix stale-base rejections with P05, retaining the underlying cause.

### FLOW-059
During audit, Verify MUST apply Done's audit-submission checks. Verify and Done MUST prefix malformed or incomplete submission rejections with P09.

### FLOW-060
Done MUST prefix its existing blocked-change error with P06 while preserving the committed blocking state and claim release.

### FLOW-061
Start MUST distinguish an explicit unknown change selector from queue eligibility. Lifecycle operations MUST return U02 for unknown explicit change selectors, naming the rejected value and listing valid slugs for the operation.

### FLOW-062
Lifecycle operations MUST preserve assigned static, eval, zone and boundary causes when returning their existing failures, except for lifecycle stale-base classification P05.

### FLOW-063
Block MUST require the calling worktree's current claim with unchanged approved
metadata and phase. It MUST reject a changed HEAD rather than blocking a stale
phase base.

### FLOW-064
Block MUST save the held step's tracked and untracked work as a commit at the
next unused positive refs/fireweed/attempts/<slug>/<n>, preserving prior attempt
refs. BlockResult MUST identify that ref and its commit OID.

### FLOW-065
After saving the attempt, Block MUST restore the step's tracked files, untracked
files and index to the phase base, removing the step's edits from the checkout.

### FLOW-066
Block MUST commit only the held change's assigned IDs, run, current phase and
requested blocked reason over the phase base. The block commit MUST preserve
the approved author and already-committed audit content and exclude step work.

### FLOW-067
On successful blocking, Block MUST release the current claim and return the
blocked State with no held ClaimOID. The resulting working tree MUST be clean.

### FLOW-068
Eligible MUST return unblocked queued changes not held back by earlier blocked
changes, in queue order. Empty or entirely waiting queues MUST return an empty
result without an error.

### FLOW-069
Eligible MUST omit a later change targeting the same module as an earlier
blocked change.

### FLOW-070
Eligible MUST omit a later change whose module depends directly or transitively
on an earlier blocked change's module.

### FLOW-071
When a later change has a nonempty Interface section, Eligible MUST omit it if
an earlier blocked change's module depends directly or transitively on its
module. With an empty Interface, that dependency change MUST remain eligible
unless another blocking rule applies.

### FLOW-072
For each candidate, Eligible MUST apply the blocking rules to the graph
projected through that candidate, including preceding queued set edits and its
own set edits.

### FLOW-073
StartFromAttempt MUST resume a committed phase with the same clean-worktree,
eligibility and claim-ownership preconditions as Start. It MUST preserve the
existing run, phase and assigned IDs and leave HEAD unchanged.

### FLOW-074
StartFromAttempt MUST restore the saved step's modified, deleted and newly
added files from refs/fireweed/attempts/<canonical-slug>/<attempt> into the
working tree without reverting later independent changes.

### FLOW-075
While restoring saved worker-editable body sections for the resumed phase,
StartFromAttempt MUST preserve the current change's approved author fields,
runtime state and controller-written provenance.

### FLOW-076
StartFromAttempt MUST reject a non-positive attempt number or missing saved
attempt ref with P01 before claiming or writing the change.

### FLOW-077
For an unstarted change with no Added, Modified or Removed specs, fixes or
eval_fix, and regenerate not true, Start MUST set phase to impl.

### FLOW-078
For an eval_fix change, Verify MUST permit evals-phase edits only in the target
module's evals zone and the current change file. During impl it MUST permit
edits only to that change file.

### FLOW-079
During evals for an eval_fix change, Verify MUST require all target-module
evals to pass rather than fail first, with every listed eval_fix ID covered by
at least one referencing testcase.

### FLOW-080
Start MUST set phase to impl for an unstarted change with regenerate true.

### FLOW-081
Before returning from successful Start of an unstarted regenerate change,
Fireweed MUST delete only the target module's files classified as disposable
by the approved projected zone map, retaining all other files. When claiming a
regenerate change with an already recorded phase, Start MUST retain its existing
disposable files.

### FLOW-082
For a regenerate change, Start MUST emit warning P10 when the module has no
passing drill dated after its latest applied change's implementation Provenance.
A newer passing drill MUST suppress that warning.

### FLOW-083
For a held regenerate change in impl, Context MUST supply the regenerate
playbook in place of the ordinary impl playbook.

### FLOW-084
Release MUST drop only the named claim owned by the calling worktree while
preserving its working files, HEAD and recorded change state.

### FLOW-085
Release MUST reject a missing claim or a claim owned by another worktree with
P02, leaving refs unchanged.

### FLOW-086
Release MUST accept the same unambiguous change path, filename and slug
selectors as Current.

### FLOW-087
Commits MUST return the full object IDs, newest first, of commits reachable
from HEAD that change the selected change document's phase relative to their
first parent, following its pending and archive paths. Selection MUST use file
contents rather than commit messages.

### FLOW-088
Commits MUST resolve pending or archived changes by repository-relative path,
filename or unambiguous slug. It MUST reject missing or ambiguous selectors
with U02.

### FLOW-089
Abandon MUST undo the unapplied change's phase-commit effects identified by Commits and delete its pending change file in one commit, returning that commit's full object ID.

### FLOW-090
Abandon MUST reject a dirty working tree with P01 before changing files, HEAD or refs.

### FLOW-091
Abandon MUST allow an unclaimed change or a claim owned by the calling worktree.
It MUST reject a foreign claim with P02 before changing files, HEAD or refs.
On success it MUST release the selected owned claim.

### FLOW-092
Abandon MUST reject an applied archived change with P01.
After git revert returns that change to the pending queue, Abandon MUST allow dropping it under the same preconditions.

### FLOW-093
Abandon MUST retain saved attempt refs unchanged.

### FLOW-094
The Abandon commit MUST not assign a new phase value.

### FLOW-095
After successful Abandon, its assigned but unapplied spec IDs MUST be available for subsequent Start allocation.

### FLOW-096
Abandon MUST accept the same unambiguous repository-relative path, filename and slug selectors as Commits, rejecting missing or ambiguous selectors with U02.

### FLOW-097
Each Verify that evaluates modules, including through Done, MUST retain a cost snapshot under its existing run/phase directory and return those module costs in Result.Costs.
The snapshot MUST identify the run, phase and UTC recording time, replacing the previous cost snapshot for that run/phase.

### FLOW-098
Each evaluated module's cost MUST report its name, summed JUnit testcase seconds and number of matching testcases per spec ID, counting each testcase once per matching ID.
It MUST report eval and disposable line counts for regular files selected by the held approved projected zone map, counting an unterminated final line and excluding higher-priority zones.

### FLOW-099
Verify and Done MUST return warnings B01 for total evaluated-module seconds above budgets.verify_time (default15m), B02 for each spec's testcase count above budgets.cases_per_spec (default15), and B03 for each module's eval/disposable line ratio above budgets.eval_ratio (default4).
Warnings MUST identify measurement and limit without affecting verification success or failure.
Positive eval lines with zero disposable lines MUST exceed the ratio limit.

### FLOW-100
Costs MUST return the retained run/phase cost snapshots in recording-time order with their module values and budget diagnostics, or an empty slice when none exist.
The retained values MUST reflect the evaluated state even if source files or budget settings later change.

### FLOW-101
Status MUST return every project.Queue entry in queue order with slug, path, module, phase, blocked reason and audit rounds.
Absent phase, reason and rounds MUST appear as queued, empty string and zero respectively.

### FLOW-102
For an unblocked queued change targeting an earlier blocked change's module, Status MUST name that earlier change in WaitingOn with WaitingRule same_module.

### FLOW-103
For an unblocked queued change whose module depends directly or transitively on an earlier blocked change's module, Status MUST name that earlier change with WaitingRule depends_on_blocked.

### FLOW-104
For an unblocked queued change with a nonempty Interface section, Status MUST name an earlier blocked change with WaitingRule interface_dependency when that blocked module depends directly or transitively on the candidate's module.
An empty Interface MUST not trigger this rule.

### FLOW-105
Status MUST determine waiting reasons using the graph projected through each candidate, including preceding and candidate set edits.
Blocked entries and unblocked entries unaffected by the blocking rules MUST have empty WaitingOn and WaitingRule.

### FLOW-106
Status MUST expose the discovered modules with their declared identities, dependency graph and resolved zones.

### FLOW-107
Status MUST list every currently quarantined spec with its module, ID and quarantine reason.

### FLOW-108
Status MUST list every module with boundaries_opt_out and its authored reason.

### FLOW-109
Status MUST report each module's count of applied archived changes whose implementation Provenance date is later than its latest drill report date.
With no drill report, it MUST count all that module's applied archived changes, including zero when there are none.

### FLOW-110
Status MUST count applied changes archived on HEAD whose change file is absent from the locally known configured remote/base ref, falling back to the local base when the remote ref is absent.
The base and remote MUST use the existing configured/default base-ref resolution.

### FLOW-111
Status MUST warn with P08 when the three or more most recently applied changes consecutively target one module, identifying the module and consecutive count.

### FLOW-112
Status MUST report the number of phase-changing commits reachable from HEAD for each applied archived change, matching Commits for that change.

### FLOW-113
Status MUST expose the latest retained ModuleCost for every module with recorded costs, together with the saved budget diagnostics associated with those selected measurements.
Modules without a cost record MUST not receive a fabricated measurement.

### FLOW-114
PRBody MUST render a Markdown description for an explicit Git commit range, identifying its resolved endpoints.
Project configuration, module/spec metadata and change documents MUST come from the range endpoint rather than later HEAD or working-tree edits.

### FLOW-115
PRBody MUST identify each pending or archived change present at the endpoint whose document identity was touched in the range, including its path, module and change kind.
An empty selection MUST be reported explicitly.

### FLOW-116
For every selected change, PRBody MUST include its recorded eval result counts from Provenance and identify its eval-file edits.

### FLOW-117
PRBody MUST identify interface edits, new disposable stubs added in evals, module manifest/other-file edits and worker_writable edits associated with each selected change.

### FLOW-118
PRBody MUST identify preserved-file edits associated with selected changes.

### FLOW-119
PRBody MUST list every review-tagged Added or Modified spec of each selected change with its ID and normative text.
It MUST not list an unrelated pre-existing review spec as changed.

### FLOW-120
PRBody MUST include each selected change's audit_rounds and recorded audit findings, including findings in earlier phase commits that were corrected before the endpoint.

### FLOW-121
PRBody MUST include the selected changes' added and removed Decisions.

### FLOW-122
PRBody MUST list all quarantined endpoint specs with module, ID and reason, every endpoint boundary opt-out with module and reason, and every blocked pending endpoint change with its reason.

### FLOW-123
PRBody MUST show the latest retained measurements and associated B01–B03 diagnostics for modules represented at the endpoint, without fabricating missing measurements.
Saved local cost snapshots MUST remain the measurement source rather than later file counts or recalculated budgets.

### FLOW-124
PRBody MUST list every binary file and every file larger than 256 KB newly added to a disposable zone between the range endpoints, identifying path and byte size.
Each such file MUST produce a B04 warning in Diagnostics and the Markdown body; ordinary smaller text additions MUST not produce B04.

### FLOW-125
An empty or unresolvable range MUST return an error with P01 and the rejected range or resolution cause.
A valid empty range MUST return a report rather than an error.

### FLOW-126
PRBody MUST return an error for which IsEndpointConfigError reports true when the project at the range endpoint fails to load.

### FLOW-127
IsEndpointConfigError MUST report false when PRBody fails because the range cannot be resolved.

### FLOW-128
GC MUST prune refs/fireweed/claims entries whose recorded worktree is missing or is no longer registered with the same Git common directory, while retaining claims for live registered worktrees.

### FLOW-129
GC MUST prune refs/fireweed/attempts/<slug>/* only for changes present in changes/archive, retaining attempt refs for unapplied changes.

### FLOW-130
GC MUST prune refs/fireweed/drills/<module>/<date> when the UTC timestamp encoded in the native drill ref name is older than gc.drill_max_age relative to GCRequest.Now, or current UTC time when Now is zero, retaining refs at or newer than the cutoff.

### FLOW-131
GC MUST read gc.drill_max_age as a positive Go duration, defaulting to 720h (30 days) when unset. An invalid configured value MUST return C01 before any pruning.

### FLOW-132
GC MUST return the sorted full names of only its successfully deleted refs. It MUST leave all other refs, HEAD, index and working files unchanged.

### FLOW-133
For a valid queue, Current(root, "") MUST return an error whose text is `P02: worktree holds no change` when the calling worktree owns none of its change claims.

### FLOW-134
During full impl Verify, mapped testcase failures in valid reports MUST NOT
stop evaluation of the remaining selected modules.
Verify MUST return every completed module's report in Result.Reports.
Cancellation, timeouts, hook execution errors, and invalid or missing reports
MAY stop further evaluation.

## Decisions

[0012-manual-lifecycle] Provenance lines begin `- <phase>` and use `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` for the reports verified by that Done call.

[0119-verification-costs] Cost snapshots use .fireweed/runs/<run>/<phase>/costs.json beside existing module
reports; repeated verification in that phase replaces its snapshot, as it does
the module reports. Each regular file's line count is newline count plus one
when nonempty content lacks a trailing newline.

[0120-status-report] When several earlier changes hold one entry, report the first in queue order,
with same-module, dependent and interface-dependency precedence. Applied order
uses implementation Provenance dates, newest first with archive path as the
tie-breaker; a drill report date counts regardless of result. Budget diagnostics
from selected snapshots are returned once each, rather than recomputed from
later settings or source edits. PhaseCommits keys are archived change filenames
without .md.

[0123-pr-body-report] Selected changes are identities touched by range commits and present at its end;
removed change documents are absent from that endpoint report. Review context for
a selected change includes its phase commits reachable at the endpoint, so a
range covering only its final phase retains earlier eval/stub/audit information.
Binary follows Git's text/binary classification; 256 KB is262144 bytes, with
strictly larger text files warned. File categories use that phase's declared
metadata; endpoint exceptions use the endpoint's declarations.
Cost measurements use a Cost/Metric Markdown section with one row per measured
module: module, seconds, cases per spec (`ID: count`), eval lines, disposable lines.


Retained interfaces and files:
packages/workflow/SPEC.md
packages/workflow/api.go
packages/workflow/model/model.go

Retained evals:
packages/workflow/evals/.gitkeep
packages/workflow/evals/abandon_test.go
packages/workflow/evals/block_test.go
packages/workflow/evals/commits_test.go
packages/workflow/evals/costs_test.go
packages/workflow/evals/current_absence_test.go
packages/workflow/evals/eligible_test.go
packages/workflow/evals/eval_fix_lifecycle_test.go
packages/workflow/evals/gc_test.go
packages/workflow/evals/lifecycle_test.go
packages/workflow/evals/maintenance_start_test.go
packages/workflow/evals/pr_body_test.go
packages/workflow/evals/regenerate_start_test.go
packages/workflow/evals/release_test.go
packages/workflow/evals/resume_test.go
packages/workflow/evals/selector_audit_validation_test.go
packages/workflow/evals/status_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.


# Previous verification failed
workflow eval: E01: failed case "Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries" for FLOW-002
app eval: E01: failed case "Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks" for APP-010
runner eval: E01: failed case "Test_RUN_010_EntirelyWaitingQueue" for RUN-010
workflow stdout: ?   	github.com/lurkshark/fireweed/packages/workflow	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/workflow/evals	118.120s

workflow stderr: exit status 1

workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange: === RUN   Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange
    abandon_test.go:95: fixture evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"f69fed5ed89f35b84fedb2c4de422c4b"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange456515819/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"f69fed5ed89f35b84fedb2c4de422c4b", Phase:"evals", Head:"cf5fbe7fac79e7690b48bc7088c375ec8ba5f313", PhaseBase:"cf5fbe7fac79e7690b48bc7088c375ec8ba5f313", ClaimOID:"1e67ddac253eef0c806b37f0e168c360a752dab9", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:1, DisposableLines:0}}, Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"B03", Message:"alpha eval/disposable line ratio +Inf (1/0) exceeds the limit of 4", Path:"alpha", Warning:true}}}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending: === RUN   Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending
    abandon_test.go:181: fixture evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"f2d6921b5373347439221bee0e5af985"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending3560597012/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"f2d6921b5373347439221bee0e5af985", Phase:"evals", Head:"ce66e8be39901d4917368805c84a9a2317521f0d", PhaseBase:"ce66e8be39901d4917368805c84a9a2317521f0d", ClaimOID:"1b73490ea91804d7285be5bc6a080efe2aeda886", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_092_AbandonAppliedRefusalThenGitRevertToPending (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit: === RUN   Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit
    block_test.go:173: evals Done: model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"aaa26425bf2067b461459b89bbd88de3"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit1297880482/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"aaa26425bf2067b461459b89bbd88de3", Phase:"evals", Head:"4ac86b930197b354bd57b87760fdf2d6f40f9173", PhaseBase:"4ac86b930197b354bd57b87760fdf2d6f40f9173", ClaimOID:"d2d91e0f0d306f9be53e393c69e5904e49dcd39b", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:1, DisposableLines:0}}, Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"B03", Message:"alpha eval/disposable line ratio +Inf (1/0) exceeds the limit of 4", Path:"alpha", Warning:true}}}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot: === RUN   Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot
    costs_test.go:113: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_097_VerifyAndDoneReplaceRunPhaseCostSnapshot (0.85s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts: === RUN   Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts
    costs_test.go:164: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_098_ModuleJUnitCostsAndHeldZoneLineCounts (0.57s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure: === RUN   Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure
    costs_test.go:217: warning made evals Done fail: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_099_CustomWarningsSurviveSuccessAndFailure (0.69s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_099_DefaultBudgetBoundaries: === RUN   Test_FLOW_099_DefaultBudgetBoundaries
--- FAIL: Test_FLOW_099_DefaultBudgetBoundaries (1.18s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_099_DefaultBudgetBoundaries/above: === RUN   Test_FLOW_099_DefaultBudgetBoundaries/above
    costs_test.go:304: default warnings made evals Done fail: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_099_DefaultBudgetBoundaries/above (0.70s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits: === RUN   Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits
    costs_test.go:325: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_100_CostsEmptyOrderedAndRetainedAfterEdits (0.61s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits: === RUN   Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits
    eval_fix_lifecycle_test.go:77: eval_fix evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"eval_fix":[]interface {}{"ALP-001"}, "module":"alpha", "phase":"evals", "run":"a5159ec276b8cba36f384e7ba8d98be8"}, Body:"\n## Motivation\nCorrect the wrong ALP-001 eval while retaining the contract and implementation.\n\n## Interface\n\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"}, Added:[]model.Item(nil), Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWi3820717810/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the old value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"a5159ec276b8cba36f384e7ba8d98be8", Phase:"evals", Head:"6af316a2d6371c6742bc56bab2bffff5f45761b6", PhaseBase:"6af316a2d6371c6742bc56bab2bffff5f45761b6", ClaimOID:"5965affeff192d6b6323ab4d258c7493758587ce", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"pass", Failure:"", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:3, DisposableLines:1}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_078_FLOW_079_FLOW_036_EvalFixCorrectsEvalAndArchivesWithoutContractOrImplementationEdits (0.72s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks: === RUN   Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks
--- FAIL: Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks (1.31s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/impl_implementation: === RUN   Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/impl_implementation
    eval_fix_lifecycle_test.go:159: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"eval_fix":[]interface {}{"ALP-001"}, "module":"alpha", "phase":"evals", "run":"1fca3ee8b4baf9e79eb203b916de652e"}, Body:"\n## Motivation\nCorrect the wrong ALP-001 eval while retaining the contract and implementation.\n\n## Interface\n\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"}, Added:[]model.Item(nil), Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBefor1333953650/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the old value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"1fca3ee8b4baf9e79eb203b916de652e", Phase:"evals", Head:"568cbd1decbea8733f07c2c4041cf00d6ab46380", PhaseBase:"568cbd1decbea8733f07c2c4041cf00d6ab46380", ClaimOID:"bfad0b1e5349b64149c9e091e6425cd23faa4569", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"pass", Failure:"", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:3, DisposableLines:1}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_078_EvalFixRejectsInterfaceAndImplementationEditsBeforeHooks/impl_implementation (0.59s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit: === RUN   Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit
    eval_fix_lifecycle_test.go:189: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"eval_fix":[]interface {}{"ALP-001"}, "module":"alpha", "phase":"evals", "run":"cf24e0e8058c48040d60ee2729e384f6"}, Body:"\n## Motivation\nCorrect the wrong ALP-001 eval while retaining the contract and implementation.\n\n## Interface\n\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"}, Added:[]model.Item(nil), Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit2193656263/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the old value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"cf24e0e8058c48040d60ee2729e384f6", Phase:"evals", Head:"c57ea1e47b10108bb7e190689ce7caf4196ca393", PhaseBase:"c57ea1e47b10108bb7e190689ce7caf4196ca393", ClaimOID:"64636200676185115c1f2ae80a7cbf5d20bae445", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"pass", Failure:"", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:3, DisposableLines:1}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_036_EvalFixAuditRequiresListedIDBeforeCommit (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly: === RUN   Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly
    gc_test.go:98: GC: git worktree: exit status 129: error: unknown switch `z'
        usage: git worktree add [<options>] <path> [<commit-ish>]
           or: git worktree list [<options>]
           or: git worktree lock [<options>] <path>
           or: git worktree move <worktree> <new-path>
           or: git worktree prune [<options>]
           or: git worktree remove [<options>] <worktree>
           or: git worktree unlock <path>
        
            --porcelain           machine-readable output
--- FAIL: Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly (0.34s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts: === RUN   Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts
    gc_test.go:118: GC: git worktree: exit status 129: error: unknown switch `z'
        usage: git worktree add [<options>] <path> [<commit-ish>]
           or: git worktree list [<options>]
           or: git worktree lock [<options>] <path>
           or: git worktree move <worktree> <new-path>
           or: git worktree prune [<options>]
           or: git worktree remove [<options>] <worktree>
           or: git worktree unlock <path>
        
            --porcelain           machine-readable output
--- FAIL: Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts (0.21s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate: === RUN   Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate
    gc_test.go:160: GC: git worktree: exit status 129: error: unknown switch `z'
        usage: git worktree add [<options>] <path> [<commit-ish>]
           or: git worktree list [<options>]
           or: git worktree lock [<options>] <path>
           or: git worktree move <worktree> <new-path>
           or: git worktree prune [<options>]
           or: git worktree remove [<options>] <worktree>
           or: git worktree unlock <path>
        
            --porcelain           machine-readable output
--- FAIL: Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate (0.62s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/configured: === RUN   Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/configured
    gc_test.go:145: GC: git worktree: exit status 129: error: unknown switch `z'
        usage: git worktree add [<options>] <path> [<commit-ish>]
           or: git worktree list [<options>]
           or: git worktree lock [<options>] <path>
           or: git worktree move <worktree> <new-path>
           or: git worktree prune [<options>]
           or: git worktree remove [<options>] <worktree>
           or: git worktree unlock <path>
        
            --porcelain           machine-readable output
--- FAIL: Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/configured (0.21s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/default: === RUN   Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/default
    gc_test.go:145: GC: git worktree: exit status 129: error: unknown switch `z'
        usage: git worktree add [<options>] <path> [<commit-ish>]
           or: git worktree list [<options>]
           or: git worktree lock [<options>] <path>
           or: git worktree move <worktree> <new-path>
           or: git worktree prune [<options>]
           or: git worktree remove [<options>] <worktree>
           or: git worktree unlock <path>
        
            --porcelain           machine-readable output
--- FAIL: Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate/default (0.21s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState: === RUN   Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState
    gc_test.go:201: GC: git worktree: exit status 129: error: unknown switch `z'
        usage: git worktree add [<options>] <path> [<commit-ish>]
           or: git worktree list [<options>]
           or: git worktree lock [<options>] <path>
           or: git worktree move <worktree> <new-path>
           or: git worktree prune [<options>]
           or: git worktree remove [<options>] <worktree>
           or: git worktree unlock <path>
        
            --porcelain           machine-readable output
--- FAIL: Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState (0.26s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries: === RUN   Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries
--- FAIL: Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries (2.75s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/frozen-eval: === RUN   Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/frozen-eval
    lifecycle_test.go:663: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/frozen-eval (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/direct: === RUN   Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/direct
    lifecycle_test.go:663: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/direct (0.54s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/all: === RUN   Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/all
    lifecycle_test.go:663: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/all (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/none: === RUN   Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/none
    lifecycle_test.go:663: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/none (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/quick-boundary: === RUN   Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/quick-boundary
    lifecycle_test.go:663: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_002_FLOW_003_FLOW_022_FLOW_025_FLOW_026_FLOW_027_FLOW_029_FLOW_030_FLOW_031_ImplFreezesEvalsAndRunsDependentsAndBoundaries/quick-boundary (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures: === RUN   Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures
--- FAIL: Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures (2.22s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/direct: === RUN   Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/direct
    lifecycle_test.go:762: fixture evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/direct (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/all: === RUN   Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/all
    lifecycle_test.go:762: fixture evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/all (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/none: === RUN   Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/none
    lifecycle_test.go:762: fixture evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_fail/none (0.56s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_pass/all: === RUN   Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_pass/all
    lifecycle_test.go:762: fixture evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_134_FullImplVerifyReturnsAllSelectedReportsAfterMappedFailures/own_pass/all (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl: === RUN   Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl
--- FAIL: Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl (1.27s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/false: === RUN   Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/false
    lifecycle_test.go:864: Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"d132dd18f73173b4bbea3c0c23d3ba10"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisab405965659/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"d132dd18f73173b4bbea3c0c23d3ba10", Phase:"evals", Head:"cab5a5a5a4a885d05b21d1c01b051582155510e8", PhaseBase:"cab5a5a5a4a885d05b21d1c01b051582155510e8", ClaimOID:"552f0fc7648c88da0624c6e60df2ff626a60c8ef", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"; want audit with commit
    lifecycle_test.go:868: phase base not advanced: model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"d132dd18f73173b4bbea3c0c23d3ba10"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisab405965659/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"d132dd18f73173b4bbea3c0c23d3ba10", Phase:"evals", Head:"cab5a5a5a4a885d05b21d1c01b051582155510e8", PhaseBase:"cab5a5a5a4a885d05b21d1c01b051582155510e8", ClaimOID:"552f0fc7648c88da0624c6e60df2ff626a60c8ef", Diagnostics:[]model.Diagnostic(nil)}, <nil>
--- FAIL: Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/false (0.64s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/true: === RUN   Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/true
    lifecycle_test.go:864: Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"099a28648eb0a5c1f34016205fc4ca37"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisab892295005/001", Config:map[string]interface {}{"audit":map[string]interface {}{"enabled":false}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"099a28648eb0a5c1f34016205fc4ca37", Phase:"evals", Head:"c9a63530889b47eae6783c6254b615af1ee6ee86", PhaseBase:"c9a63530889b47eae6783c6254b615af1ee6ee86", ClaimOID:"52ac806bd28e7acc1eb0236e436c40d982fbae1a", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"; want impl with commit
    lifecycle_test.go:868: phase base not advanced: model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"099a28648eb0a5c1f34016205fc4ca37"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisab892295005/001", Config:map[string]interface {}{"audit":map[string]interface {}{"enabled":false}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"099a28648eb0a5c1f34016205fc4ca37", Phase:"evals", Head:"c9a63530889b47eae6783c6254b615af1ee6ee86", PhaseBase:"c9a63530889b47eae6783c6254b615af1ee6ee86", ClaimOID:"52ac806bd28e7acc1eb0236e436c40d982fbae1a", Diagnostics:[]model.Diagnostic(nil)}, <nil>
--- FAIL: Test_FLOW_032_FLOW_033_FLOW_046_EvalsDoneUsesDefaultAuditOrDisabledImpl/true (0.63s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings: === RUN   Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings
    lifecycle_test.go:876: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"26d29241959f6d935fe437ac6a4ed016"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactCl572132980/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"26d29241959f6d935fe437ac6a4ed016", Phase:"evals", Head:"3c0efa6c1dc58bd4a79106524803c7222533b92a", PhaseBase:"3c0efa6c1dc58bd4a79106524803c7222533b92a", ClaimOID:"4d5e16abe460a4ffd5c8be4502750e8140a96222", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings (0.52s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound: === RUN   Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound
    lifecycle_test.go:897: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"c43055448ad6fdda91643f6c109000cb"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n### ALP-003 [review]\nExtra MUST be visible.\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item{model.Item{ID:"ALP-003", Tag:"review", Body:"Extra MUST be visible.\n\n"}}, Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_1714148278/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n### ALP-003 [review]\nExtra MUST be visible.\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"c43055448ad6fdda91643f6c109000cb", Phase:"evals", Head:"abb3a637438980ba87711ddddfd1615ef56f5b81", PhaseBase:"abb3a637438980ba87711ddddfd1615ef56f5b81", ClaimOID:"97441cea3abbceecd2581310050ba93659f50ff4", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound (0.52s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion: === RUN   Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion
    lifecycle_test.go:965: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"858e80e7b6980381bb98827e9b2a8323"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuest2790254811/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"858e80e7b6980381bb98827e9b2a8323", Phase:"evals", Head:"b7ac65b5888fdb5d00bcdb147b7011ce6ece0df9", PhaseBase:"b7ac65b5888fdb5d00bcdb147b7011ce6ece0df9", ClaimOID:"7f8f23aed40aaf2bf4207c83b86f410f79c924ce", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion (0.52s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives: === RUN   Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives
    lifecycle_test.go:997: evals transition = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676", "ALP-002":"321a28df"}, "module":"alpha", "phase":"evals", "run":"d2149c8587d198417f69d402a7aa182b", "set":map[string]interface {}{"hooks":map[string]interface {}{"eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "interface":[]interface {}{"api.go", "extra.go"}}}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n### ALP-003 [review]\nExtra MUST be visible.\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n### ALP-002\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item{model.Item{ID:"ALP-003", Tag:"review", Body:"Extra MUST be visible.\n\n"}}, Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"\n"}}}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_1536689560/001", Config:map[string]interface {}{"audit":map[string]interface {}{"enabled":false}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/global-eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "verify_dependents":"none", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go", "extra.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "hooks":map[string]interface {}{"eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "id_prefix":"ALP", "interface":[]interface {}{"api.go", "extra.go"}, "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-003 [review]\nExtra MUST be visible.\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"d2149c8587d198417f69d402a7aa182b", Phase:"evals", Head:"d6473aa04886e92c59ca9b2d5ebb5ff192fefc5a", PhaseBase:"d6473aa04886e92c59ca9b2d5ebb5ff192fefc5a", ClaimOID:"8f60be180433e492d1276e0a9e5c5d79ab486a8f", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives (0.56s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead: === RUN   Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead
    lifecycle_test.go:1087: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead (0.64s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports: === RUN   Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports
    lifecycle_test.go:1204: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports (0.71s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_059_AuditNoFindingsErrorsHaveP09: === RUN   Test_FLOW_059_AuditNoFindingsErrorsHaveP09
    lifecycle_test.go:1364: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"bf2d26c9ac755ea07625d6f43464701a"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_059_AuditNoFindingsErrorsHaveP091024908416/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"bf2d26c9ac755ea07625d6f43464701a", Phase:"evals", Head:"39f9590c2d6b8d03397093bdbe3eae4b1a68c7d9", PhaseBase:"39f9590c2d6b8d03397093bdbe3eae4b1a68c7d9", ClaimOID:"4c949645d11787d65900fb39d2d28e6d19cb25e7", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_059_AuditNoFindingsErrorsHaveP09 (0.51s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim: === RUN   Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim
    lifecycle_test.go:1378: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"07955e244057b92e884532e207a73025"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim1236049266/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"07955e244057b92e884532e207a73025", Phase:"evals", Head:"39f9590c2d6b8d03397093bdbe3eae4b1a68c7d9", PhaseBase:"39f9590c2d6b8d03397093bdbe3eae4b1a68c7d9", ClaimOID:"8e070a445448d8b1594177103ce75f7122ce0d8e", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim (0.52s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses: === RUN   Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses
    lifecycle_test.go:1432: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses (2.18s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics: === RUN   Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics
    lifecycle_test.go:1506: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_062_MissingRequiredBoundaryHookKeepsCauseAndAllDiagnostics (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly: === RUN   Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly
    maintenance_start_test.go:79: maintenance Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"impl", "run":"4f37998908db05ab92b114cae292b1a4"}, Body:"\n## Motivation\nSimplify the existing alpha implementation while keeping its behavior.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"}, Added:[]model.Item(nil), Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly24199835/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the old value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"4f37998908db05ab92b114cae292b1a4", Phase:"impl", Head:"0a4bbee3c2eed04bdcc57655c11103dc3c90a161", PhaseBase:"0a4bbee3c2eed04bdcc57655c11103dc3c90a161", ClaimOID:"b09b6f49250282a36db97864e9438e8c51a5cd40", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"pass", Failure:"", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}, "beta":model.Report{Cases:[]model.Case{model.Case{Name:"BET-001", Classname:"fixture", IDs:[]string{"BET-001"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:3, DisposableLines:1}, model.ModuleCost{Module:"beta", Seconds:0, CasesPerSpec:map[string]int{"BET-001":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_077_MaintenanceStartsInImplAndArchivesDirectly (0.81s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive: === RUN   Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive
    regenerate_start_test.go:145: regenerate Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"impl", "regenerate":true, "run":"20c7c9d75e2f744568d2bf01aaf7df08", "set":map[string]interface {}{"preserve":[]interface {}{"internal/keep.go"}}}, Body:"\n## Motivation\nRebuild alpha's implementation from retained contracts and evals.\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n\n"}, Added:[]model.Item(nil), Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDispos1149867764/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "protected":[]interface {}{"packages/alpha/internal/global.go"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go", "internal/contract.go"}, AppendOnly:[]string{"internal/log.go"}, Preserve:[]string{"internal/keep.go"}, Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"append_only":[]interface {}{"internal/log.go"}, "depends_on":[]interface {}{}, "id_prefix":"ALP", "interface":[]interface {}{"api.go", "internal/contract.go"}, "name":"alpha", "preserve":[]interface {}{"internal/keep.go"}}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the old value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"20c7c9d75e2f744568d2bf01aaf7df08", Phase:"impl", Head:"c9cdbf73d9543d296d7feda7650aa449aad061ef", PhaseBase:"c9cdbf73d9543d296d7feda7650aa449aad061ef", ClaimOID:"e4436cf4b35555430bd18bdfb2013effc4633ff2", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"pass", Failure:"", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}, "beta":model.Report{Cases:[]model.Case{model.Case{Name:"BET-001", Classname:"fixture", IDs:[]string{"BET-001"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:0, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:2, DisposableLines:1}, model.ModuleCost{Module:"beta", Seconds:0, CasesPerSpec:map[string]int{"BET-001":1}, EvalLines:0, DisposableLines:1}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_080_FLOW_081_FLOW_083_RegenerateDeletesProjectedDisposableAndRebuildsThroughArchive (0.89s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance: === RUN   Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance
    resume_test.go:155: evals Done: model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"412a39e25c1bdfee6348b6c35a2eb83b"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommi2995521250/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"412a39e25c1bdfee6348b6c35a2eb83b", Phase:"evals", Head:"bbd10338ba1c0bac89a4de8eb6a5402670e42129", PhaseBase:"bbd10338ba1c0bac89a4de8eb6a5402670e42129", ClaimOID:"5a5516e88ea0a305bbd63ee435bda729a5b8cc7b", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:1, DisposableLines:0}}, Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"B03", Message:"alpha eval/disposable line ratio +Inf (1/0) exceeds the limit of 4", Path:"alpha", Warning:true}}}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_073_FLOW_074_FLOW_075_StartFromAttemptImplRetainsCommittedEvalsAuditAndProvenance (0.55s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval: === RUN   Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval
    resume_test.go:251: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"a9078c445f0237623ccc8b2ba663c1ef"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrent469598512/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"a9078c445f0237623ccc8b2ba663c1ef", Phase:"evals", Head:"70a261c1f324a2cec2d071bbbfeccb12e62759ee", PhaseBase:"70a261c1f324a2cec2d071bbbfeccb12e62759ee", ClaimOID:"b16a56a2a169e77fa617b05da270edf21c58e1f5", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_075_StartFromAttemptRestoresAuditBodyAndRetainsCurrentApproval (0.54s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing: === RUN   Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing
    selector_audit_validation_test.go:15: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"51d7849220f68ea2d1b7a0f2a8cdd23f"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing123097556/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"51d7849220f68ea2d1b7a0f2a8cdd23f", Phase:"evals", Head:"fe89e0d1e631d30e6bf98aba5439fbbd12640525", PhaseBase:"fe89e0d1e631d30e6bf98aba5439fbbd12640525", ClaimOID:"6c675e3767eda16322467c6bfa8ed886822e3d42", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing (0.51s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks: === RUN   Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks
    selector_audit_validation_test.go:74: evals Done = model.Result{State:model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{"ALP-001":"4305e676"}, "module":"alpha", "phase":"evals", "run":"411e3f304e726d67e5e37b3ad075c31f"}, Body:"\n## Motivation\nUsers need a revised alpha value.\n\n## Interface\nAlpha public value is revised.\n\n## Added\n\n## Modified\n### ALP-001\nAlpha MUST return the revised value.\n\n## Removed\n\n## Decisions\nUse the revised policy.\n\n## Removed decisions\nRetain old policy.\n"}, Added:[]model.Item(nil), Modified:[]model.Item{model.Item{ID:"ALP-001", Tag:"", Body:"Alpha MUST return the revised value.\n\n"}}, Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks2723400656/001", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return the revised value.\n\n### ALP-002\nAlpha MUST retain the second behavior.\n\n## Decisions\nRetain old policy.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{"beta"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"beta"}, "id_prefix":"GAM", "name":"gamma"}, Body:"# Gamma\n\n## Specs\n\n### GAM-001\nGamma MUST follow beta.\n"}}}}, Run:"411e3f304e726d67e5e37b3ad075c31f", Phase:"evals", Head:"3980d2007bf731a071f66970b034b2cbde2c6068", PhaseBase:"3980d2007bf731a071f66970b034b2cbde2c6068", ClaimOID:"149d5b39993dcbd9326d44fdddbb375950eee9df", Diagnostics:[]model.Diagnostic(nil)}, Reports:map[string]model.Report{"alpha":model.Report{Cases:[]model.Case{model.Case{Name:"ALP-001", Classname:"fixture", IDs:[]string{"ALP-001"}, Result:"fail", Failure:"expected before implementation", Seconds:0}, model.Case{Name:"ALP-002", Classname:"fixture", IDs:[]string{"ALP-002"}, Result:"pass", Failure:"", Seconds:0}}, Seconds:0, HookExit:1, Stdout:"fixture-hook-output\n", Stderr:""}}, Costs:[]model.ModuleCost{model.ModuleCost{Module:"alpha", Seconds:0, CasesPerSpec:map[string]int{"ALP-001":1, "ALP-002":1}, EvalLines:0, DisposableLines:0}}, Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_002_AuditRejectsForbiddenEditBeforeHooks (0.54s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle: === RUN   Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle
    status_test.go:230: evals Done: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_112_StatusPhaseCommitsMatchRealLifecycle (0.51s)
workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings: === RUN   Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings
    status_test.go:257: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings (0.55s)

app stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	238.124s

app stderr: exit status 1

app github.com/lurkshark/fireweed/packages/app/evals Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON: === RUN   Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON
--- FAIL: Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON (1.58s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned: === RUN   Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned
    abandon_cli_test.go:75: fixture evals Done: phase="evals" err=git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_APP_097_APP_099_AbandonNestedRootPlainAndJSON/json-owned (1.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_014_MaintenanceCommandsDispatch: === RUN   Test_APP_014_MaintenanceCommandsDispatch
--- FAIL: Test_APP_014_MaintenanceCommandsDispatch (0.72s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_014_MaintenanceCommandsDispatch/gc: === RUN   Test_APP_014_MaintenanceCommandsDispatch/gc
    app_test.go:175: gc plain: status=1 stdout="" stderr="P01: git worktree: exit status 129: error: unknown switch `z'\nusage: git worktree add [<options>] <path> [<commit-ish>]\n   or: git worktree list [<options>]\n   or: git worktree lock [<options>] <path>\n   or: git worktree move <worktree> <new-path>\n   or: git worktree prune [<options>]\n   or: git worktree remove [<options>] <worktree>\n   or: git worktree unlock <path>\n\n    --porcelain           machine-readable output. Correct the reported project configuration or Git state, then retry gc.\n"
--- FAIL: Test_APP_014_MaintenanceCommandsDispatch/gc (0.35s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty: === RUN   Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty
--- FAIL: Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty (1.29s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/plain: === RUN   Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/plain
    gc_cli_test.go:111: gc: exit=1 stdout="" stderr="P01: git worktree: exit status 129: error: unknown switch `z'\nusage: git worktree add [<options>] <path> [<commit-ish>]\n   or: git worktree list [<options>]\n   or: git worktree lock [<options>] <path>\n   or: git worktree move <worktree> <new-path>\n   or: git worktree prune [<options>]\n   or: git worktree remove [<options>] <worktree>\n   or: git worktree unlock <path>\n\n    --porcelain           machine-readable output. Correct the reported project configuration or Git state, then retry gc.\n"
--- FAIL: Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/plain (0.62s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/JSON: === RUN   Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/JSON
    gc_cli_test.go:111: gc: exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"git worktree: exit status 129: error: unknown switch `z'\\nusage: git worktree add [\\u003coptions\\u003e] \\u003cpath\\u003e [\\u003ccommit-ish\\u003e]\\n   or: git worktree list [\\u003coptions\\u003e]\\n   or: git worktree lock [\\u003coptions\\u003e] \\u003cpath\\u003e\\n   or: git worktree move \\u003cworktree\\u003e \\u003cnew-path\\u003e\\n   or: git worktree prune [\\u003coptions\\u003e]\\n   or: git worktree remove [\\u003coptions\\u003e] \\u003cworktree\\u003e\\n   or: git worktree unlock \\u003cpath\\u003e\\n\\n    --porcelain           machine-readable output. Correct the reported project configuration or Git state, then retry gc.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_150_APP_151_GCNearestProjectSortedAndEmpty/JSON (0.38s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks: === RUN   Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks
    lifecycle_cli_test.go:423: evals done: 1 "" "P01: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the held change and correct its checks, then retry done.\n"
--- FAIL: Test_APP_010_APP_029_APP_030_APP_031_VerifyDoneAuditAndFullChecks (1.48s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits: === RUN   Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits
    lifecycle_cli_test.go:485: evals done failed
--- FAIL: Test_APP_010_APP_029_APP_030_APP_031_VerifyRejectsFalseFailFirstAndFrozenEdits (1.24s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit: === RUN   Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit
    lifecycle_cli_test.go:523: fixture audit phase: model.State{Change:model.Change{Number:1, Slug:"manual", Path:"changes/0001-manual.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"module":"alpha", "phase":"evals", "run":"193f46d6b85b0c1b573f427a3e47b2e7", "set":map[string]interface {}{"interface":[]interface {}{"api.go", "extra.go"}}}, Body:"\n## Motivation\nUsers need an additional value.\n\n## Interface\nAlpha exposes the additional value.\n\n## Added\n### ALP-002\nAlpha MUST return the additional value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the additional value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit3751400817/002", Config:map[string]interface {}{"base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}, "preserve":[]interface {}{"data/**"}}, "modules":[]interface {}{"packages/*"}, "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go", "extra.go"}, AppendOnly:[]string(nil), Preserve:[]string{"data/**"}, Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "interface":[]interface {}{"api.go", "extra.go"}, "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST retain the existing value.\n\n### ALP-002\nAlpha MUST return the additional value.\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string{"data/**"}, Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# Beta\n\n## Specs\n\n### BET-001\nBeta MUST follow alpha.\n"}}}}, Run:"193f46d6b85b0c1b573f427a3e47b2e7", Phase:"evals", Head:"6493c90692da0ecd8262d8e606240f5cf475ccb0", PhaseBase:"6493c90692da0ecd8262d8e606240f5cf475ccb0", ClaimOID:"97403dad41e4bbd2b23cd4b6489eac222ee06e4a", Diagnostics:[]model.Diagnostic(nil)}, git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_APP_010_APP_029_APP_031_DoneReportsBlockedAudit (0.87s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange: === RUN   Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange
    next_cli_test.go:167: next exit=1 stdout="" stderr="P01: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry next.\n"
--- FAIL: Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange (2.68s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_075_NextAllowUnmergedAndUsage: === RUN   Test_APP_075_NextAllowUnmergedAndUsage
    next_cli_test.go:242: allowed next exit=1 stdout="" stderr="P01: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry next.\n"
--- FAIL: Test_APP_075_NextAllowUnmergedAndUsage (2.72s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_079_NextJSONSuccessAndFailure: === RUN   Test_APP_079_NextJSONSuccessAndFailure
    next_cli_test.go:323: JSON applied exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\\n.fireweed/runs\\nhint: Use -f if you really want to add them.\\nhint: Turn this message off by running\\nhint: \\\"git config advice.addIgnoredFile false\\\". Use the reported cause to correct the change or project checks, then retry next.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_079_NextJSONSuccessAndFailure (5.39s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation: === RUN   Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation
    run_cli_test.go:257: run exit=1 stdout="" stderr="P01: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run.\n"
--- FAIL: Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation (5.28s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON: === RUN   Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON
    run_cli_test.go:318: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\\n.fireweed/runs\\nhint: Use -f if you really want to add them.\\nhint: Turn this message off by running\\nhint: \\\"git config advice.addIgnoredFile false\\\". Use the reported cause to correct the change or project checks, then retry run.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON (4.88s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_083_RunPlainMultiple: === RUN   Test_APP_083_RunPlainMultiple
    run_cli_test.go:342: run exit=1 stdout="" stderr="P01: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run.\n"
--- FAIL: Test_APP_083_RunPlainMultiple (3.59s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_086_RunJSONLimited: === RUN   Test_APP_086_RunJSONLimited
    run_cli_test.go:354: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\\n.fireweed/runs\\nhint: Use -f if you really want to add them.\\nhint: Turn this message off by running\\nhint: \\\"git config advice.addIgnoredFile false\\\". Use the reported cause to correct the change or project checks, then retry run.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_086_RunJSONLimited (3.57s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_082_RunAllowUnmerged: === RUN   Test_APP_082_RunAllowUnmerged
    run_cli_test.go:371: run exit=1 stdout="" stderr="P01: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run.\n"
--- FAIL: Test_APP_082_RunAllowUnmerged (3.87s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes: === RUN   Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes
    run_for_cli_test.go:95: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\\n.fireweed/runs\\nhint: Use -f if you really want to add them.\\nhint: Turn this message off by running\\nhint: \\\"git config advice.addIgnoredFile false\\\". Use the reported cause to correct the change or project checks, then retry run.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes (3.05s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_084_RunForPriorBlockStillFails: === RUN   Test_APP_084_RunForPriorBlockStillFails
    run_for_cli_test.go:231: mixed failure lacks "refs/fireweed/attempts/alpha/1": "verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."
    run_for_cli_test.go:231: mixed failure lacks "8s": "verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."
    run_for_cli_test.go:234: time-stop reason="verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."; want 8s and time limit
    run_for_cli_test.go:237: Alpha block state="---\nbases: {}\nmodule: alpha\nphase: evals\nrun: 51ae6d94aa58ea693747bd4e6377cb30\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    run_for_cli_test.go:241: missing Alpha saved attempt refs/fireweed/attempts/alpha/1
    run_for_cli_test.go:245: stop phase commits=0; want Alpha's block only
    run_for_cli_test.go:247: held beta state: P02: change has no claim
--- FAIL: Test_APP_084_RunForPriorBlockStillFails (2.72s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport: === RUN   Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport
    status_cli_test.go:247: rich public Status fixture lacks B03 warning: [{Code:P08 Message:3 consecutive recently applied changes target alpha Path: Warning:true} {Code:B01 Message:beta fixture time budget exceeded; review cost Path:packages/beta/evals/weight_test.go Warning:true}]
--- FAIL: Test_APP_118_APP_119_APP_120_APP_121_APP_122_APP_123_APP_124_APP_125_StatusRichReport (0.76s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume: === RUN   Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume
    usage_cli_test.go:158: owned resume exit=1 stderr="" stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\\n.fireweed/runs\\nhint: Use -f if you really want to add them.\\nhint: Turn this message off by running\\nhint: \\\"git config advice.addIgnoredFile false\\\". Use the reported cause to correct the change or project checks, then retry next.\",\"warning\":false}]}\n"
    usage_cli_test.go:166: resumed next data={Change: Run: Phase: Empty:false}: <nil>
    usage_cli_test.go:168: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume4134328538/002/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (3.22s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume: === RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:242: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\\n.fireweed/runs\\nhint: Use -f if you really want to add them.\\nhint: Turn this message off by running\\nhint: \\\"git config advice.addIgnoredFile false\\\". Use the reported cause to correct the change or project checks, then retry run.\",\"warning\":false}]}\n" stderr=""
    usage_cli_test.go:244: clean stop data=null; want object: <nil>
    usage_cli_test.go:260: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume1963589911/002/changes/archive/0001-alpha.md: no such file or directory
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (3.02s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_083_APP_084_RunPlainSharedUsageStop: === RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:321: plain run stop exit=1 stdout="" stderr="P01: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run.\n"
    usage_cli_test.go:323: plain run lacks alpha: ""
    usage_cli_test.go:323: plain run lacks applied count 1: ""
    usage_cli_test.go:325: plain run lacks resume guidance: ""
    usage_cli_test.go:327: stop reason "" lacks "completed"
    usage_cli_test.go:327: stop reason "" lacks "next"
    usage_cli_test.go:327: stop reason "" lacks "budget"
    usage_cli_test.go:327: stop reason "" lacks "5ms"
    usage_cli_test.go:327: stop reason "" lacks "5ms"
    usage_cli_test.go:329: alpha not applied: stat /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_RunPlainSharedUsageStop2572866543/002/changes/archive/0001-alpha.md: no such file or directory
    usage_cli_test.go:331: held beta state: P02: change has no claim
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (2.96s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop: === RUN   Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop
    usage_cli_test.go:378: combined failure lacks "refs/fireweed/attempts/alpha/1": "verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."
    usage_cli_test.go:378: combined failure lacks "5ms": "verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."
    usage_cli_test.go:378: combined failure lacks "0s": "verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."
    usage_cli_test.go:378: combined failure lacks "budget": "verify evals: C10: no coverage for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:\n.fireweed/runs\nhint: Use -f if you really want to add them.\nhint: Turn this message off by running\nhint: \"git config advice.addIgnoredFile false\". Use the reported cause to correct the change or project checks, then retry run."
    usage_cli_test.go:383: alpha block state="---\nbases: {}\nmodule: alpha\nphase: evals\nrun: a7b5be6a48babcaf2f8ee0d021332aef\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    usage_cli_test.go:387: missing saved attempt refs/fireweed/attempts/alpha/1
    usage_cli_test.go:391: held beta state: P02: change has no claim
--- FAIL: Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop (2.39s)

runner stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/runner/evals	164.735s

runner stderr: exit status 1

runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput: === RUN   Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput
    agent_logs_test.go:109: Next fixture = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"f448abda63d4286433ee4268574d8f70"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/002", "retry-both", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"f448abda63d4286433ee4268574d8f70", Phase:"evals", Head:"c1e196987ed814d9548fd8a50e9aedcf49cb8ff2", PhaseBase:"c1e196987ed814d9548fd8a50e9aedcf49cb8ff2", ClaimOID:"ee99515ff1de27ebb39b0ee4248598937c6f48bc", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/001/.fireweed/runs/f448abda63d4286433ee4268574d8f70/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit2338870879/001/.fireweed/runs/f448abda63d4286433ee4268574d8f70/evals-2.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput (3.30s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit: === RUN   Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit
    audit_workspace_test.go:73: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"93e133c150da751e9bf0ffea9544b70c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"93e133c150da751e9bf0ffea9544b70c", Phase:"evals", Head:"ab53560a9f46a21a47b2b02e10582a4e85a1ae0e", PhaseBase:"ab53560a9f46a21a47b2b02e10582a4e85a1ae0e", ClaimOID:"df8f5e1e5a06efeb1793e38df668133a84c1240a", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndInde2299132067/001/.fireweed/runs/93e133c150da751e9bf0ffea9544b70c/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_013_RUN_045_RUN_046_AuditUsesOnlyProjectedInputsAndIndependentGit (2.51s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration: === RUN   Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration
    audit_workspace_test.go:131: approved audit retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"aed874a22305e9bf6d0e9255be16dab3"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/002"}}, "limit_budget":"50ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"1ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"aed874a22305e9bf6d0e9255be16dab3", Phase:"evals", Head:"7524a2b04b0fbe37f785518a00624f06d00c669d", PhaseBase:"7524a2b04b0fbe37f785518a00624f06d00c669d", ClaimOID:"c5d52669cef9f9cc7bfc9955cd085b4fab0751d5", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration3352302336/001/.fireweed/runs/aed874a22305e9bf6d0e9255be16dab3/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration (2.26s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry: === RUN   Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry
    audit_workspace_test.go:179: audit verification retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"986544fe2416fc284bfce22c8f5244d3"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"986544fe2416fc284bfce22c8f5244d3", Phase:"evals", Head:"fd11aedb264b5e236d9494f8ac3ea93aa93ec3cf", PhaseBase:"fd11aedb264b5e236d9494f8ac3ea93aa93ec3cf", ClaimOID:"5466545230a020488b6c277c0d8ba9c08cb746b5", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1534857786/001/.fireweed/runs/986544fe2416fc284bfce22c8f5244d3/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry (2.26s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_046_AuditCommitDoesNotEnterHeldGit: === RUN   Test_RUN_046_AuditCommitDoesNotEnterHeldGit
    audit_workspace_test.go:201: audit commit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"b95c0fe39dd8ee22abac5d3c66a4aced"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/002", "commit-success", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"b95c0fe39dd8ee22abac5d3c66a4aced", Phase:"evals", Head:"b3d18fc8b5dba7bc22918d0691c476c0ecf9c88b", PhaseBase:"b3d18fc8b5dba7bc22918d0691c476c0ecf9c88b", ClaimOID:"f4045e0396c0b60e0ad9b5018a0d580d3afb1f99", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"evals agent created commits; soft-reset HEAD to the held phase base and retained index and worktree work", Path:"", Warning:true}}}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit2335691163/001/.fireweed/runs/b95c0fe39dd8ee22abac5d3c66a4aced/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_046_AuditCommitDoesNotEnterHeldGit (2.62s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported: === RUN   Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported
    audit_workspace_test.go:243: first audit failure = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"cb0fb87d58fbffd0c6b77f4fdc573cd3"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"cb0fb87d58fbffd0c6b77f4fdc573cd3", Phase:"evals", Head:"32f2a64fbee3c9c1729bc36d708a1e6257a34f12", PhaseBase:"32f2a64fbee3c9c1729bc36d708a1e6257a34f12", ClaimOID:"543709e4427692488ff23b0bd8e75dbb19d4d66d", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported336624775/001/.fireweed/runs/cb0fb87d58fbffd0c6b77f4fdc573cd3/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported (2.29s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces: === RUN   Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces
    audit_workspace_test.go:318: two-change audit lifecycle: applied=0 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces (1.38s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically (14.31s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/change_outside_Audit: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/change_outside_Audit
    audit_workspace_test.go:357: invalid audit was not rejected while retaining claim: phase="evals" claim="d742bb8ba59361df2ecc20d407461216e52fe043" steps=1 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:380: invalid audit advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallycha1878867573/001|/fixture/next-bin"}
    audit_workspace_test.go:382: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallycha1878867573/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/change_outside_Audit (2.29s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/projected_spec: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/projected_spec
    audit_workspace_test.go:357: invalid audit was not rejected while retaining claim: phase="evals" claim="a243b22aca18a22f470ae602963c41e126a67532" steps=1 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:380: invalid audit advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallypro3336382681/001|/fixture/next-bin"}
    audit_workspace_test.go:382: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallypro3336382681/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/projected_spec (2.29s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/eval: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/eval
    audit_workspace_test.go:357: invalid audit was not rejected while retaining claim: phase="evals" claim="44b640dc88d4741b406c81660124bef3004bed1b" steps=1 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:380: invalid audit advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallyeva1784908863/001|/fixture/next-bin"}
    audit_workspace_test.go:382: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallyeva1784908863/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/eval (2.25s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/added_file: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/added_file
    audit_workspace_test.go:357: invalid audit was not rejected while retaining claim: phase="evals" claim="21d8bfdaebb4082fe1f8dfff9b2e39d43c28bf2d" steps=1 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:380: invalid audit advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallyadd1638700356/001|/fixture/next-bin"}
    audit_workspace_test.go:382: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallyadd1638700356/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/added_file (2.29s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/deleted_file: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/deleted_file
    audit_workspace_test.go:357: invalid audit was not rejected while retaining claim: phase="evals" claim="eba43eb79acaef9f891c17f65d42837a5a667a96" steps=1 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:380: invalid audit advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallydel3863848751/001|/fixture/next-bin"}
    audit_workspace_test.go:382: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallydel3863848751/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/deleted_file (2.73s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/committed_invalid_edit: === RUN   Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/committed_invalid_edit
    audit_workspace_test.go:357: invalid audit was not rejected while retaining claim: phase="evals" claim="c8c98d87e7b37447f93ca94833bd39b6a65b086b" steps=1 error=complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:380: invalid audit advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallycom3441896801/001|/fixture/next-bin"}
    audit_workspace_test.go:382: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomicallycom3441896801/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_047_InvalidAuditWorkspaceEditsAreRejectedAtomically/committed_invalid_edit (2.46s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops: === RUN   Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops
--- FAIL: Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops (6.26s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/agent_failure: === RUN   Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/agent_failure
    audit_workspace_test.go:422: audit failure result = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"de2a4e4a6a2398f2e47e3225a3350f92"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"de2a4e4a6a2398f2e47e3225a3350f92", Phase:"evals", Head:"9f7a1fc0240a7bf7372131e10c25fd1559b2a644", PhaseBase:"9f7a1fc0240a7bf7372131e10c25fd1559b2a644", ClaimOID:"55b7f50c33c3dd2deca2e51c9a89a22364f3f30e", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/001/.fireweed/runs/de2a4e4a6a2398f2e47e3225a3350f92/evals-1.log"}}}}
    audit_workspace_test.go:426: partial Audit not retained in held change: "---\nbases: {}\nmodule: alpha\nphase: evals\nrun: de2a4e4a6a2398f2e47e3225a3350f92\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    audit_workspace_test.go:429: audit error advanced phases 0 times, want only evals
    audit_workspace_test.go:432: audit error released or advanced phase: model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"de2a4e4a6a2398f2e47e3225a3350f92"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"de2a4e4a6a2398f2e47e3225a3350f92", Phase:"evals", Head:"9f7a1fc0240a7bf7372131e10c25fd1559b2a644", PhaseBase:"9f7a1fc0240a7bf7372131e10c25fd1559b2a644", ClaimOID:"55b7f50c33c3dd2deca2e51c9a89a22364f3f30e", Diagnostics:[]model.Diagnostic(nil)}, <nil>
    audit_workspace_test.go:435: audit error advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/001|/fixture/next-bin"}
    audit_workspace_test.go:437: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsa2169341574/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/agent_failure (2.35s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/usage_stop: === RUN   Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/usage_stop
    audit_workspace_test.go:419: audit failure/stop = complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:422: audit failure result = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"fea03c99cea8145c435bf5cda0a40c86"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002"}}, "limit_budget":"0s", "limit_patterns":[]interface {}{"RATE LIMIT"}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"fea03c99cea8145c435bf5cda0a40c86", Phase:"evals", Head:"336d730234887d260a8c3e301e7818d20fb70ef6", PhaseBase:"336d730234887d260a8c3e301e7818d20fb70ef6", ClaimOID:"4658148cbacf19a4e2154f086d74fd524a4bfe27", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/001/.fireweed/runs/fea03c99cea8145c435bf5cda0a40c86/evals-1.log"}}}}
    audit_workspace_test.go:426: partial Audit not retained in held change: "---\nbases: {}\nmodule: alpha\nphase: evals\nrun: fea03c99cea8145c435bf5cda0a40c86\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    audit_workspace_test.go:429: audit error advanced phases 0 times, want only evals
    audit_workspace_test.go:432: audit error released or advanced phase: model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"fea03c99cea8145c435bf5cda0a40c86"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002"}}, "limit_budget":"0s", "limit_patterns":[]interface {}{"RATE LIMIT"}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"fea03c99cea8145c435bf5cda0a40c86", Phase:"evals", Head:"336d730234887d260a8c3e301e7818d20fb70ef6", PhaseBase:"336d730234887d260a8c3e301e7818d20fb70ef6", ClaimOID:"4658148cbacf19a4e2154f086d74fd524a4bfe27", Diagnostics:[]model.Diagnostic(nil)}, <nil>
    audit_workspace_test.go:435: audit error advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/001|/fixture/next-bin"}
    audit_workspace_test.go:437: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopsu1091459233/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/usage_stop (2.42s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/time_stop: === RUN   Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/time_stop
    audit_workspace_test.go:419: audit failure/stop = complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    audit_workspace_test.go:426: partial Audit not retained in held change: "---\nbases: {}\nmodule: alpha\nphase: evals\nrun: ccc9c9e774392c087144774d7f0f165a\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    audit_workspace_test.go:429: audit error advanced phases 0 times, want only evals
    audit_workspace_test.go:432: audit error released or advanced phase: model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"ccc9c9e774392c087144774d7f0f165a"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"ccc9c9e774392c087144774d7f0f165a", Phase:"evals", Head:"fdd965bfcf07def16adcb9c4cbd0ca74f8561d6d", PhaseBase:"fdd965bfcf07def16adcb9c4cbd0ca74f8561d6d", ClaimOID:"825fd220dcb21cc2bafa1605d1e2933c562007e2", Diagnostics:[]model.Diagnostic(nil)}, <nil>
    audit_workspace_test.go:435: audit error advanced phase: []string{"evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/001|/fixture/next-bin"}
    audit_workspace_test.go:437: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStopst1694774439/002/audit.cwd: no such file or directory
--- FAIL: Test_RUN_048_RUN_049_PermittedAuditWorkSurvivesFailureAndStops/time_stop (1.49s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle: === RUN   Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle
    commit_recovery_test.go:38: Next after fixture agent commits: complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"; state=model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"8dce1acf046ae6216f9e431c64c8e7b3"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/002", "commit-success", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"8dce1acf046ae6216f9e431c64c8e7b3", Phase:"evals", Head:"a0bfec043d856528e0c86ff98bf90d352d64ee41", PhaseBase:"a0bfec043d856528e0c86ff98bf90d352d64ee41", ClaimOID:"009f7669ac75f5d172c17f0a1b48538390086859", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"evals agent created commits; soft-reset HEAD to the held phase base and retained index and worktree work", Path:"", Warning:true}}}; steps=[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2365137999/001/.fireweed/runs/8dce1acf046ae6216f9e431c64c8e7b3/evals-1.log"}}}
--- FAIL: Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle (2.49s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_011_WaitingHeadSelectsIndependent: === RUN   Test_RUN_011_WaitingHeadSelectsIndependent
    next_test.go:418: fixture eligible queue = []model.Change{model.Change{Number:1, Slug:"alpha", Path:"changes/0001-alpha.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"de282eb2506a0b52a3f38046c2750d9e"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, model.Change{Number:2, Slug:"beta", Path:"changes/0002-beta.md", Module:"beta", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"beta"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\nbeta MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"new:value", Tag:"", Body:"beta MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, model.Change{Number:3, Slug:"gamma", Path:"changes/0003-gamma.md", Module:"gamma", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"gamma"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\ngamma MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"new:value", Tag:"", Body:"gamma MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}}, <nil>
--- FAIL: Test_RUN_011_WaitingHeadSelectsIndependent (2.60s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_010_EntirelyWaitingQueue: === RUN   Test_RUN_010_EntirelyWaitingQueue
    next_test.go:442: fixture entirely waiting queue = []model.Change{model.Change{Number:1, Slug:"alpha", Path:"changes/0001-alpha.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"928a0d76070a7a8e41b86f70dbb838dc"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, model.Change{Number:2, Slug:"beta", Path:"changes/0002-beta.md", Module:"beta", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"beta"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### new:value\nbeta MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"new:value", Tag:"", Body:"beta MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}}, <nil>
--- FAIL: Test_RUN_010_EntirelyWaitingQueue (2.22s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_013_RUN_014_CompleteSpecChange: === RUN   Test_RUN_013_RUN_014_CompleteSpecChange
    next_test.go:458: Next: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"9048f8898d7805e029d10f202737e982"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/002", "normal", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"9048f8898d7805e029d10f202737e982", Phase:"evals", Head:"12822109464d0dd07f8648bb2839d49405c6c874", PhaseBase:"12822109464d0dd07f8648bb2839d49405c6c874", ClaimOID:"700e9f96d5d95ebba5a679d0a9494f84eb3022d3", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange2303103209/001/.fireweed/runs/9048f8898d7805e029d10f202737e982/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_013_RUN_014_CompleteSpecChange (3.10s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks: === RUN   Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks
    next_test.go:661: block commit parents = 0, want held evals HEAD
    next_test.go:661: block commit paths = "", want state only
    next_test.go:661: blocked checkout dirty: "M changes/0001-next.md"
    next_test.go:661: blocked claim retained: "refs/fireweed/claims/next"
    next_test.go:661: blocked change still held
    next_test.go:661: block document lost phase, run, assigned ID or reason
    next_test.go:661: returned block state: phase="evals" runPresent=true assignedID=true claimReleased=false headMatches=true snapshotMatches=true blocked=""
    next_test.go:661: exhaustion error does not identify refs/fireweed/attempts/next/1: verify evals: P2-01: no fail-first case for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    next_test.go:661: saved attempt ref absent: refs/fireweed/attempts/next/1
--- FAIL: Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks (2.19s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase: === RUN   Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase
    next_test.go:673: Next after retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"8e914f1989dfa4453a4aa6c2b5f810bb"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/002", "retry-both", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"8e914f1989dfa4453a4aa6c2b5f810bb", Phase:"evals", Head:"ba89f08fd103ec062c6c93b3519073e3a2f8cc72", PhaseBase:"ba89f08fd103ec062c6c93b3519073e3a2f8cc72", ClaimOID:"6868b472cb5ef5f23ccb16a65fa96c1cf5800a60", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/001/.fireweed/runs/8e914f1989dfa4453a4aa6c2b5f810bb/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1539577955/001/.fireweed/runs/8e914f1989dfa4453a4aa6c2b5f810bb/evals-2.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase (3.07s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_021_ConfiguredThreeAttempts: === RUN   Test_RUN_021_ConfiguredThreeAttempts
    next_test.go:726: Next with three attempts = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"fb7ac9921d34a04d807bdb8d04a9af6b"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/002", "retry-third", "{prompt}"}}, "attempts":3, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"fb7ac9921d34a04d807bdb8d04a9af6b", Phase:"evals", Head:"50ace2707234939876c9586dba02a749940ea29d", PhaseBase:"50ace2707234939876c9586dba02a749940ea29d", ClaimOID:"eb4ecb4ad70e035db4e1d361ad277b05a3453375", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts3384741248/001/.fireweed/runs/fb7ac9921d34a04d807bdb8d04a9af6b/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_021_ConfiguredThreeAttempts (2.22s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks: === RUN   Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks
    next_test.go:752: exhaustion cause = complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    next_test.go:755: exhaustion steps = 1, want 4
--- FAIL: Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks (2.37s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules: === RUN   Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules
    retry_failure_details_test.go:90: Next with JUnit-only dependent failure = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"2a5091ee4310dac57bb2256a056a1555"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/002", "retry-impl", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"aardvark", IDPrefix:"AAR", Dir:"packages/aardvark", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"AAR", "name":"aardvark"}, Body:"# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"2a5091ee4310dac57bb2256a056a1555", Phase:"evals", Head:"c30c073028bd3432d174d50706c51bb8048cf646", PhaseBase:"c30c073028bd3432d174d50706c51bb8048cf646", ClaimOID:"820d63e8894dedea6d3318f16d225a18937752c2", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModul3184044944/001/.fireweed/runs/2a5091ee4310dac57bb2256a056a1555/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_043_RetryIncludesJUnitOnlyCasesFromOwnAndDependentModules (1.36s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions: === RUN   Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions
    retry_failure_details_test.go:133: Next with large JUnit failures = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"99ce769b1c1ef54e8b9a8f97eae361ba"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/002", "retry-impl", "{prompt}"}}, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"aardvark", IDPrefix:"AAR", Dir:"packages/aardvark", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"AAR", "name":"aardvark"}, Body:"# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"99ce769b1c1ef54e8b9a8f97eae361ba", Phase:"evals", Head:"5058217e399f954e48243cb53340ab1b7cf77308", PhaseBase:"5058217e399f954e48243cb53340ab1b7cf77308", ClaimOID:"e82768d876962dfcea8a86026ba476d082d098b5", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1632849064/001/.fireweed/runs/99ce769b1c1ef54e8b9a8f97eae361ba/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions (1.39s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_016_UnlimitedQueueOrder: === RUN   Test_RUN_016_UnlimitedQueueOrder
    run_test.go:305: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_016_UnlimitedQueueOrder (2.29s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_016_NegativeMaxIsUnlimited: === RUN   Test_RUN_016_NegativeMaxIsUnlimited
    run_test.go:321: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_016_NegativeMaxIsUnlimited (2.18s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding: === RUN   Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding
    run_test.go:333: first Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding (2.24s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_018_InitialEmptyAndExactLimit: === RUN   Test_RUN_018_InitialEmptyAndExactLimit
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit (2.42s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_018_InitialEmptyAndExactLimit/exact_limit: === RUN   Test_RUN_018_InitialEmptyAndExactLimit/exact_limit
    run_test.go:377: exact Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit/exact_limit (2.17s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_020_HeldFailureRetainsEarlierResults: === RUN   Test_RUN_020_HeldFailureRetainsEarlierResults
    run_test.go:391: held failure cause = complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    run_test.go:182: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_020_HeldFailureRetainsEarlierResults (2.39s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent: === RUN   Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent
    run_test.go:419: Run blocking error = verify evals: P2-01: no fail-first case for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false", want verification cause and refs/fireweed/attempts/alpha/1
    run_test.go:419: missing saved attempt refs/fireweed/attempts/alpha/1
    run_test.go:419: alpha block state = "---\nbases: {}\nmodule: alpha\nphase: evals\nrun: 9bde24d01a3f7804f4300f07bbd10f31\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    run_test.go:419: blocked run retained claim "refs/fireweed/claims/alpha"
    run_test.go:419: blocked run left dirty checkout "M changes/0001-alpha.md"
    run_test.go:182: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent (2.09s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_025_RUN_026_PositiveMaxRetainsBlock: === RUN   Test_RUN_025_RUN_026_PositiveMaxRetainsBlock
    run_test.go:440: Run blocking error = verify evals: P2-01: no fail-first case for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false", want verification cause and refs/fireweed/attempts/alpha/1
    run_test.go:440: missing saved attempt refs/fireweed/attempts/alpha/1
    run_test.go:440: alpha block state = "---\nbases: {}\nmodule: alpha\nphase: evals\nrun: d1ba6ba889e8e87d04b26680c6ef4f3d\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    run_test.go:440: blocked run retained claim "refs/fireweed/claims/alpha"
    run_test.go:440: blocked run left dirty checkout "M changes/0001-alpha.md"
    run_test.go:182: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_025_RUN_026_PositiveMaxRetainsBlock (2.02s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied: === RUN   Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied
    time_limit_test.go:99: clean time stop = complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false", want classified nonnil error naming 10s
    time_limit_test.go:107: Applied count = 0, want one retained alpha
    time_limit_test.go:110: invocations through held beta = []string{"alpha|evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied3911881878/001|/fixture/run-bin"}
--- FAIL: Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied (2.45s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded: === RUN   Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded
    time_limit_test.go:121: zero For run = model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded (1.40s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop: === RUN   Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop
    time_limit_test.go:300: combined block and time failure: Applied=0, err=verify evals: P2-01: no fail-first case for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    time_limit_test.go:303: blocked alpha lost saved attempt ref
    time_limit_test.go:308: held state: P02: change has no claim
--- FAIL: Test_RUN_040_EarlierBlockRetainsCauseBeforeTimeStop (1.59s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep: === RUN   Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep
    usage_test.go:204: owned resume = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"f37d13acd75399005d7fa2163950412b"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/002", "normal", "{prompt}"}}, "attempts":1, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"f37d13acd75399005d7fa2163950412b", Phase:"evals", Head:"2f10b3017767468b32f4c49ec70cc04be39b1708", PhaseBase:"2f10b3017767468b32f4c49ec70cc04be39b1708", ClaimOID:"3152b02df548457043a810cfc1bfbeb60c17cbde", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/001/.fireweed/runs/f37d13acd75399005d7fa2163950412b/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    usage_test.go:206: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep4254209542/001/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep (3.39s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_030_UsageRetryPreservesVerificationFeedback: === RUN   Test_RUN_030_UsageRetryPreservesVerificationFeedback
    usage_test.go:286: verification and usage retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"3fcccec3781b058d74346566bd3013c8"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/002", "normal", "{prompt}"}}, "attempts":2, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/002"}}, "limit_budget":"20ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"3fcccec3781b058d74346566bd3013c8", Phase:"evals", Head:"6779f87a5eb49731e741f593c73212577af62ef3", PhaseBase:"6779f87a5eb49731e741f593c73212577af62ef3", ClaimOID:"0c789223d4468348aa21aa8c80b06cc00b91a7f9", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"first verification", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/001/.fireweed/runs/3fcccec3781b058d74346566bd3013c8/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"limited evals", Stderr:"RATE LIMIT", ExitCode:9, UsageLimit:true, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/001/.fireweed/runs/3fcccec3781b058d74346566bd3013c8/evals-2.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback3045127824/001/.fireweed/runs/3fcccec3781b058d74346566bd3013c8/evals-2.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_030_UsageRetryPreservesVerificationFeedback (3.31s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases: === RUN   Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases
    usage_test.go:305: cross-phase budget = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"evals", "run":"98f8dd30acb4318f19dd1a75e254ac6f"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/001", Config:map[string]interface {}{"agent":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/002", "normal", "{prompt}"}}, "attempts":1, "audit":map[string]interface {}{"command":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/002", "{prompt}"}}, "base":"main", "fireweed":">=0.1, <0.2", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n## Decisions\n"}}}}, Run:"98f8dd30acb4318f19dd1a75e254ac6f", Phase:"evals", Head:"2cc8ef381ad662eaf5a67e794ebc2de858e09192", PhaseBase:"2cc8ef381ad662eaf5a67e794ebc2de858e09192", ClaimOID:"264b6165267fe5cf179554a68d7495516453f658", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"limited evals", Stderr:"RATE LIMIT", ExitCode:9, UsageLimit:true, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/001/.fireweed/runs/98f8dd30acb4318f19dd1a75e254ac6f/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases3145253386/001/.fireweed/runs/98f8dd30acb4318f19dd1a75e254ac6f/evals-1.log"}}}}, complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
--- FAIL: Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases (2.47s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied: === RUN   Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied
    usage_test.go:327: Run stop = complete evals: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    usage_test.go:329: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied (2.47s)
runner github.com/lurkshark/fireweed/packages/runner/evals Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure: === RUN   Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure
    usage_test.go:369: combined block and usage failure = model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, verify evals: P2-01: no fail-first case for ALP-002; block exhausted step: git add: exit status 1: The following paths are ignored by one of your .gitignore files:
        .fireweed/runs
        hint: Use -f if you really want to add them.
        hint: Turn this message off by running
        hint: "git config advice.addIgnoredFile false"
    usage_test.go:372: blocked alpha lost its saved attempt ref
    usage_test.go:375: alpha block state = "---\nbases: {}\nmodule: alpha\nphase: evals\nrun: 36fc5858da538a14176c99151f55fa1d\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n"
    usage_test.go:378: held state: P02: change has no claim
--- FAIL: Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure (2.09s)

codex
I’ll check the retained workflow interfaces and evals first, then rebuild within the disposable paths and verify the failures shown in the log.
exec
/bin/zsh -lc "ls -a; rg --files --hidden -g '"'!.git/**'"' -g 'AGENTS.md' -g 'DRILL_NOTES.md' -g '*fireweed*y*' -g 'go.mod' -g 'api.go' -g 'model.go' -g 'SPEC.md' -g 'SYSTEM.md'; rg --files packages/workflow" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 664ms:
.
..
.git
DRILL_NOTES.md
SYSTEM.md
packages
SYSTEM.md
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/assets/api.go
packages/assets/model/model.go
packages/workflow/SPEC.md
packages/workflow/api.go
packages/workflow/model/model.go
packages/project/api.go
packages/project/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/checks/api.go
packages/checks/model/model.go
DRILL_NOTES.md
packages/workflow/SPEC.md
packages/workflow/api.go
packages/workflow/model/model.go
packages/workflow/internal/lifecycle/done.go
packages/workflow/internal/lifecycle/verify.go
packages/workflow/internal/lifecycle/status.go
packages/workflow/internal/lifecycle/history.go
packages/workflow/internal/lifecycle/claim.go
packages/workflow/internal/lifecycle/common.go
packages/workflow/internal/lifecycle/context.go
packages/workflow/internal/lifecycle/pr.go
packages/workflow/internal/lifecycle/costs.go
packages/workflow/internal/lifecycle/attempt.go
packages/workflow/internal/lifecycle/queue.go
packages/workflow/internal/lifecycle/git.go
packages/workflow/internal/lifecycle/audit.go
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
packages/workflow/evals/block_test.go
packages/workflow/evals/regenerate_start_test.go
packages/workflow/evals/commits_test.go
packages/workflow/evals/release_test.go

exec
/bin/zsh -lc "cat packages/workflow/api.go; cat packages/gitstore/api.go; cat DRILL_NOTES.md; ls -la packages; rg --files --hidden -g '"'!.git/**'"' | wc -l; rg --files --hidden -g '"'!.git/**'"' -g '"'!packages/workflow/**'"'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 557ms:
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
## Contradicted Decisions

None found in the retained workflow Decisions.

## Gaps

- FLOW-001 and FLOW-019 · workflow. The claim payload format is left to the implementation. The rebuild stores the canonical worktree, change identity, phase base, approved document and projected snapshot in a JSON blob addressed by the claim ref.
- FLOW-020 and FLOW-075 · workflow. The spec leaves the complete list of worker-editable body sections open. The retained evals establish Decisions and Questions for evals and Decisions for impl, and Audit for audit. The rebuild also treats Questions as editable during impl.
- FLOW-059 · workflow. The spec leaves the audit submission grammar open. The rebuild uses the checked-ID line and quoted finding bullets from the retained evals, with an optional untestable marker.
- FLOW-074 · workflow. The spec leaves conflict resolution open when a saved attempt overlaps later commits. The rebuild applies the saved delta, tries Git's three-way merge when needed, and returns an error for unresolved conflicts.
- FLOW-117 · workflow. The spec leaves file classification open for historical commits that predate loadable project metadata. The rebuild retains their recorded findings and result counts and classifies files for commits with loadable metadata.
total 0
drwxr-xr-x@ 8 joel  staff  256 Nov  8 15:37 .
drwx------@ 6 joel  staff  192 Nov  8 15:54 ..
drwxr-xr-x@ 4 joel  staff  128 Nov  8 15:37 assets
drwxr-xr-x@ 4 joel  staff  128 Nov  8 15:37 checks
drwxr-xr-x@ 4 joel  staff  128 Nov  8 15:37 contracts
drwxr-xr-x@ 4 joel  staff  128 Nov  8 15:37 gitstore
drwxr-xr-x@ 4 joel  staff  128 Nov  8 15:37 project
drwxr-xr-x@ 7 joel  staff  224 Nov  8 15:41 workflow
      45
SYSTEM.md
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/assets/api.go
packages/assets/model/model.go
packages/project/api.go
packages/project/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/checks/api.go
packages/checks/model/model.go
DRILL_NOTES.md

exec
/bin/zsh -lc "rg -n 'gitignore|run[s/]|B03|costs.json|func Test_FLOW_12[89]|func Test_FLOW_13[012]' packages/workflow/evals/{lifecycle_test,status_test,costs_test,gc_test}.go; ls -la" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc "rg -n 'gitstore.Commit|\\badd\\b|worktree.*porcelain|func GC|func Status|latest' packages/workflow/internal/lifecycle; cat packages/gitstore/model/model.go; cat SYSTEM.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 582ms:
packages/workflow/evals/status_test.go:268:	var latestAlphaB03, olderAlphaB03 project.Diagnostic
packages/workflow/evals/status_test.go:269:	var hasLatestAlphaB03, hasOlderAlphaB03 bool
packages/workflow/evals/status_test.go:274:				olderAlphaB03, hasOlderAlphaB03 = latestAlphaB03, hasLatestAlphaB03
packages/workflow/evals/status_test.go:275:				latestAlphaB03, hasLatestAlphaB03 = warningFor(s.Diagnostics, "B03")
packages/workflow/evals/status_test.go:279:	if !hasOlderAlphaB03 || !hasLatestAlphaB03 || olderAlphaB03 == latestAlphaB03 {
packages/workflow/evals/status_test.go:280:		t.Fatalf("cost snapshots need distinct older and latest alpha B03 warnings: %#v", snapshots)
packages/workflow/evals/status_test.go:303:	if d, ok := warningFor(r.Diagnostics, "B03"); !ok || !d.Warning || d != latestAlphaB03 {
packages/workflow/evals/status_test.go:304:		t.Errorf("selected B03 warning = %#v, want %#v", r.Diagnostics, latestAlphaB03)
packages/workflow/evals/status_test.go:306:	if containsDiagnostic(r.Diagnostics, olderAlphaB03) {
packages/workflow/evals/status_test.go:307:		t.Errorf("older B03 warning returned: %#v", r.Diagnostics)
packages/workflow/evals/status_test.go:311:		if d.Code == "B03" {
packages/workflow/evals/status_test.go:316:		t.Errorf("selected B03 warning count = %d, want 1: %#v", count, r.Diagnostics)
packages/workflow/evals/gc_test.go:74:func Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly(t *testing.T) {
packages/workflow/evals/gc_test.go:109:func Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts(t *testing.T) {
packages/workflow/evals/gc_test.go:128:func Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate(t *testing.T) {
packages/workflow/evals/gc_test.go:164:func Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning(t *testing.T) {
packages/workflow/evals/gc_test.go:185:func Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState(t *testing.T) {
packages/workflow/evals/costs_test.go:20:	path := filepath.Join(root, ".fireweed", "runs", run, phase, "costs.json")
packages/workflow/evals/costs_test.go:211:	requireCostWarning(t, verified.Diagnostics, "B03", "alpha", "1")
packages/workflow/evals/costs_test.go:222:	for _, code := range []string{"B02", "B03"} {
packages/workflow/evals/costs_test.go:237:	for _, code := range []string{"B01", "B02", "B03"} {
packages/workflow/evals/costs_test.go:245:		for _, code := range []string{"B01", "B02", "B03"} {
packages/workflow/evals/costs_test.go:286:			for _, code := range []string{"B01", "B02", "B03"} {
packages/workflow/evals/costs_test.go:301:				requireCostWarning(t, result.Diagnostics, "B03", "alpha", "5", "4")
packages/workflow/evals/costs_test.go:306:				for _, code := range []string{"B01", "B02", "B03"} {
packages/workflow/evals/costs_test.go:346:			requireCostWarning(t, snapshot.Diagnostics, "B03")
packages/workflow/evals/lifecycle_test.go:67:	write(t, root, ".gitignore", ".fireweed/runs/\n")
packages/workflow/evals/lifecycle_test.go:201:		t.Errorf("separate starts returned runs %q and %q", s.Run, another.Run)
packages/workflow/evals/lifecycle_test.go:1266:	runRoot := filepath.Join(root, ".fireweed/runs", state.Run)
total 16
drwx------@    6 joel  staff     192 Nov  8 15:54 .
drwx------@ 5528 joel  staff  176896 Nov  8 16:06 ..
drwxr-xr-x@    9 joel  staff     288 Nov  8 16:06 .git
-rw-r--r--@    1 joel  staff    1292 Nov  8 15:55 DRILL_NOTES.md
-rw-r--r--@    1 joel  staff    1950 Nov  8 15:37 SYSTEM.md
drwxr-xr-x@    8 joel  staff     256 Nov  8 15:37 packages

 succeeded in 583ms:
packages/workflow/internal/lifecycle/status.go:38:func Status(root string) (report model.StatusReport, err error) {
packages/workflow/internal/lifecycle/status.go:124:	costs, ds, err := latestCosts(root, snapshot.Modules, false)
packages/workflow/internal/lifecycle/status.go:130:func GC(ctx context.Context, request model.GCRequest) (result model.GCResult, err error) {
packages/workflow/internal/lifecycle/status.go:160:	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
packages/workflow/internal/lifecycle/pr.go:116:	costs, diagnostics, err := latestCosts(root, snapshot.Modules, true)
packages/workflow/internal/lifecycle/common.go:403:	var latest time.Time
packages/workflow/internal/lifecycle/common.go:407:			if date := parseDate(fields[2]); date.After(latest) {
packages/workflow/internal/lifecycle/common.go:408:				latest = date
packages/workflow/internal/lifecycle/common.go:412:	return latest
packages/workflow/internal/lifecycle/costs.go:135:func latestCosts(root string, modules []project.Module, restrict bool) ([]model.ModuleCost, []project.Diagnostic, error) {
packages/workflow/internal/lifecycle/git.go:83:	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
packages/workflow/internal/lifecycle/git.go:92:	_, err = idx.run(nil, "update-index", "--add", "--cacheinfo", "100644", oid, name)
packages/workflow/internal/lifecycle/claim.go:314:		var latest time.Time
packages/workflow/internal/lifecycle/claim.go:317:				if date := implementationDate(archive); date.After(latest) {
packages/workflow/internal/lifecycle/claim.go:318:					latest = date
packages/workflow/internal/lifecycle/claim.go:326:		if !passing || !date.After(latest) {
packages/workflow/internal/lifecycle/claim.go:328:				Code: "P10", Path: chosen.Path, Warning: true, Message: "module " + chosen.Module + " has no passing drill newer than its latest implementation",
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
# Workflow reconstruction

Reconstruct workflow from the supplied retained workflow SPEC (including Decisions), interface/model and public evals. Only the declared contracts, gitstore, assets, project and checks public api.go/model interfaces are provided as dependency inputs.

Write only packages/workflow/internal/** and DRILL_NOTES.md. All supplied inputs are immutable. No go.mod, replacement dependency, hook, manifest, tests or root setup may be added. Do not run local builds, tests, or module setup, including temporary module setup. Missing compilation infrastructure is intentional. Format only the fresh code you write, then return it to the separate outside verifier; that verifier runs full builds, own evals and all transitive dependent evals.

Read only this supplied input tree and the fresh disposable code you write here. Its Git history is input-only. Do not read original/deleted implementation, old history, refs/attempts/proofs/reports/reconstructions, bootstrap or rebuild-v1, outside repositories, verifier paths, binaries, or source caches. Do not search outside this tree, use controller queries as an oracle, run other agents, or contact the network. Do not infer any permission from a visible outside path or environment variable.

Do not commit, change Git refs/history/config, invoke lifecycle mutators, modify retained inputs, or weaken any verification requirement. Return to the outside controller after reconstruction; it owns all transitions and commits. Use the configured workspace-write command and model/effort only.

The current tree's .fireweed/bin is first on PATH and GOCACHE=/tmp/fireweed-rebuild-cache. Do not invoke any existing Fireweed binary to obtain behavior or implementation. No controller/model fallback is permitted. Record contradicted Decisions and targeted gaps in DRILL_NOTES.md, naming the nearest spec ID (or no spec) and owning module. A gap authorizes no correction of retained inputs.

exec
/bin/zsh -lc 'cat packages/workflow/internal/lifecycle/attempt.go; cat packages/workflow/internal/lifecycle/done.go; cat packages/workflow/evals/gc_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc 'cat packages/workflow/internal/lifecycle/git.go; cat packages/workflow/internal/lifecycle/status.go; cat packages/workflow/internal/lifecycle/costs.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 581ms:
package lifecycle

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

type gitRef struct {
	name string
	oid  string
}

func listRefs(ctx context.Context, root string, prefixes ...string) ([]gitRef, error) {
	args := append([]string{"for-each-ref", "--format=%(refname) %(objectname)"}, prefixes...)
	out, err := git(ctx, root, nil, nil, args...)
	if err != nil {
		return nil, err
	}
	var refs []gitRef
	for _, line := range strings.Split(string(out), "\n") {
		if name, oid, ok := strings.Cut(line, " "); ok {
			refs = append(refs, gitRef{name, oid})
		}
	}
	return refs, nil
}

func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
	out, err := git(ctx, root, nil, nil, args...)
	if err != nil {
		return nil, err
	}
	parts := splitNUL(out)
	var changes []gitstore.Change
	for len(parts) > 0 {
		if len(parts) < 2 || parts[0] == "" {
			return nil, fmt.Errorf("invalid git name-status output")
		}
		change := gitstore.Change{Status: parts[0], Path: parts[1]}
		parts = parts[2:]
		if change.Status[0] == 'R' || change.Status[0] == 'C' {
			if len(parts) == 0 {
				return nil, fmt.Errorf("missing destination in git name-status output")
			}
			change.OldPath, change.Path = change.Path, parts[0]
			parts = parts[1:]
		}
		changes = append(changes, change)
	}
	return changes, nil
}

type gitIndex struct {
	root string
	env  []string
	ctx  context.Context
}

func temporaryIndex(ctx context.Context, root, base string) (gitIndex, func(), error) {
	dir, err := os.MkdirTemp("", "fireweed-index-")
	if err != nil {
		return gitIndex{}, nil, err
	}
	clean := func() { os.RemoveAll(dir) }
	idx := gitIndex{root: root, env: []string{"GIT_INDEX_FILE=" + filepath.Join(dir, "index")}, ctx: ctx}
	if _, err := idx.run(nil, "read-tree", base); err != nil {
		clean()
		return gitIndex{}, nil, err
	}
	return idx, clean, nil
}

func (idx gitIndex) run(input []byte, args ...string) ([]byte, error) {
	return git(idx.ctx, idx.root, idx.env, input, args...)
}

func (idx gitIndex) stageWorktree() error {
	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
	return err
}

func (idx gitIndex) stageDocument(name string, data []byte) error {
	oid, err := gitstore.WriteBlob(idx.root, data)
	if err != nil {
		return err
	}
	_, err = idx.run(nil, "update-index", "--add", "--cacheinfo", "100644", oid, name)
	return err
}

func (idx gitIndex) commit(parent, message string) (string, error) {
	out, err := idx.run(nil, "write-tree")
	if err != nil {
		return "", err
	}
	tree := strings.TrimSpace(string(out))
	env := append([]string(nil), idx.env...)
	for _, identity := range []struct{ key, value, author, committer string }{
		{"user.name", "Fireweed", "GIT_AUTHOR_NAME", "GIT_COMMITTER_NAME"},
		{"user.email", "fireweed@localhost", "GIT_AUTHOR_EMAIL", "GIT_COMMITTER_EMAIL"},
	} {
		if value, _ := gitOutput(idx.root, "config", "--get", identity.key); value == "" {
			for _, key := range []string{identity.author, identity.committer} {
				if os.Getenv(key) == "" {
					env = append(env, key+"="+identity.value)
				}
			}
		}
	}
	out, err = git(idx.ctx, idx.root, env, []byte(message+"\n"), "commit-tree", tree, "-p", parent)
	return strings.TrimSpace(string(out)), err
}

func (idx gitIndex) checkout(revision string) error {
	_, err := idx.run(nil, "read-tree", "--reset", "-u", revision)
	return err
}

func replaceIndex(root, revision string) error {
	_, err := gitOutput(root, "read-tree", "--reset", revision)
	return err
}

func commitRefs(root, oldHead, newHead, extra string) error {
	commands := fmt.Sprintf("update HEAD %s %s\n%s", newHead, oldHead, extra)
	_, err := git(context.Background(), root, nil, []byte(commands), "update-ref", "--stdin")
	return err
}

func claimUpdate(ref, oldOID, newOID string) string {
	if newOID == "" {
		return fmt.Sprintf("delete %s %s\n", ref, oldOID)
	}
	return fmt.Sprintf("update %s %s %s\n", ref, newOID, oldOID)
}

func transitionError(root, base, name string, err error) error {
	if head, inspectErr := gitOutput(root, "rev-parse", "HEAD"); inspectErr == nil && head != base {
		return withCode("P05", name, fmt.Errorf("stale phase base: %w", err))
	}
	return withCode("P02", name, err)
}
package lifecycle

import (
	"context"
	"fmt"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func exceptions(snapshot project.Snapshot) ([]model.Quarantine, []model.BoundaryOptOut, error) {
	var quarantines []model.Quarantine
	var optOuts []model.BoundaryOptOut
	for _, m := range snapshot.Modules {
		items, err := specs(m)
		if err != nil {
			return nil, nil, err
		}
		for _, item := range items {
			if reason, ok := quarantineReason(item); ok {
				quarantines = append(quarantines, model.Quarantine{Module: m.Name, ID: item.ID, Reason: reason})
			}
		}
		if reason := stringValue(m.Document.Frontmatter["boundaries_opt_out"]); reason != "" {
			optOuts = append(optOuts, model.BoundaryOptOut{Module: m.Name, Reason: reason})
		}
	}
	return quarantines, optOuts, nil
}

func Status(root string) (report model.StatusReport, err error) {
	repo, err := repository(root)
	if err != nil {
		return report, err
	}
	root = repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return report, err
	}
	report.Modules = snapshot.Modules
	queue, err := project.Queue(root)
	if err != nil {
		return report, err
	}
	report.Queue, err = waitingQueue(root, queue)
	if err != nil {
		return report, err
	}
	report.Quarantines, report.BoundaryOptOuts, err = exceptions(snapshot)
	if err != nil {
		return report, err
	}
	archives, err := archivedAt(root, repo.Head)
	if err != nil {
		return report, err
	}
	shared, _, err := baseRefs(root, snapshot.Config)
	if err != nil {
		return report, err
	}
	report.Unapproved, err = unapprovedCount(root, shared, archives)
	if err != nil {
		return report, err
	}
	report.DriftSinceDrill = map[string]int{}
	dates := map[string]time.Time{}
	for _, m := range snapshot.Modules {
		date, _, err := drillReport(root, m.Name)
		if err != nil {
			return report, err
		}
		dates[m.Name] = date
		report.DriftSinceDrill[m.Name] = 0
	}
	report.PhaseCommits = map[string]int{}
	type application struct {
		change project.Change
		date   time.Time
	}
	var applications []application
	for _, ch := range archives {
		date := implementationDate(ch)
		if drill, ok := dates[ch.Module]; ok && (drill.IsZero() || date.After(drill)) {
			report.DriftSinceDrill[ch.Module]++
		}
		filename := path.Base(ch.Path)
		commits, err := phaseCommits(root, repo.Head, filename)
		if err != nil {
			return report, err
		}
		report.PhaseCommits[strings.TrimSuffix(filename, ".md")] = len(commits)
		applications = append(applications, application{ch, date})
	}
	sort.Slice(applications, func(i, j int) bool {
		if applications[i].date.Equal(applications[j].date) {
			return applications[i].change.Path < applications[j].change.Path
		}
		return applications[i].date.After(applications[j].date)
	})
	if len(applications) != 0 {
		module := applications[0].change.Module
		consecutive := 0
		for _, application := range applications {
			if application.change.Module != module {
				break
			}
			consecutive++
		}
		if consecutive >= 3 {
			report.Diagnostics = append(report.Diagnostics, project.Diagnostic{
				Code: "P08", Warning: true,
				Message: fmt.Sprintf("%d consecutive recently applied changes target %s", consecutive, module),
			})
		}
	}
	costs, ds, err := latestCosts(root, snapshot.Modules, false)
	report.Costs = costs
	report.Diagnostics = append(report.Diagnostics, ds...)
	return report, err
}

func GC(ctx context.Context, request model.GCRequest) (result model.GCResult, err error) {
	repo, err := repository(request.Root)
	if err != nil {
		return result, err
	}
	root := repo.Root
	snapshot, err := project.Load(root)
	if err != nil {
		return result, err
	}
	age := 720 * time.Hour
	if value, configured := configGroup(snapshot.Config, "gc")["drill_max_age"]; configured {
		age, err = time.ParseDuration(stringValue(value))
		if err != nil || age <= 0 {
			return result, coded("C01", "gc.drill_max_age", fmt.Sprintf("gc.drill_max_age must be a positive Go duration, got %v", value))
		}
	}
	now := request.Now
	if now.IsZero() {
		now = time.Now()
	}
	cutoff := now.UTC().Add(-age)
	archives, err := changePaths(root, archiveDir)
	if err != nil {
		return result, err
	}
	archived := map[string]bool{}
	for _, ch := range archives {
		archived[ch.Slug] = true
	}
	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
	if err != nil {
		return result, err
	}
	live := map[string]bool{}
	common := canonicalPath(repo.CommonDir, root)
	for _, line := range splitNUL(out) {
		if name, ok := strings.CutPrefix(line, "worktree "); ok {
			if other, err := repository(name); err == nil && canonicalPath(other.CommonDir, other.Root) == common {
				live[other.Root] = true
			}
		}
	}
	refs, err := listRefs(ctx, root, claimPrefix, attemptPrefix, drillPrefix)
	if err != nil {
		return result, err
	}
	for _, ref := range refs {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		prune := false
		switch {
		case strings.HasPrefix(ref.name, claimPrefix):
			c, err := readClaim(root, ref.oid)
			prune = err == nil && c.Worktree != "" && !live[canonicalPath(c.Worktree, root)]
		case strings.HasPrefix(ref.name, attemptPrefix):
			slug, n, ok := strings.Cut(strings.TrimPrefix(ref.name, attemptPrefix), "/")
			attempt, numberErr := strconv.Atoi(n)
			prune = ok && numberErr == nil && attempt > 0 && strconv.Itoa(attempt) == n && archived[slug]
		case strings.HasPrefix(ref.name, drillPrefix):
			module, stamp, ok := strings.Cut(strings.TrimPrefix(ref.name, drillPrefix), "/")
			date, dateErr := time.Parse("20060102T150405.000000000", stamp)
			prune = ok && module != "" && dateErr == nil && date.UTC().Before(cutoff)
		}
		if prune {
			if _, deleteErr := git(ctx, root, nil, nil, "update-ref", "-d", ref.name, ref.oid); deleteErr == nil {
				result.Pruned = append(result.Pruned, ref.name)
			} else if current, found, refErr := gitstore.Ref(root, ref.name); refErr != nil || found && current == ref.oid {
				err = combine(err, deleteErr)
			}
		}
	}
	sort.Strings(result.Pruned)
	return result, err
}

func canonicalPath(name, root string) string {
	if !filepath.IsAbs(name) {
		name = filepath.Join(root, name)
	}
	if resolved, err := filepath.EvalSymlinks(name); err == nil {
		return resolved
	}
	return filepath.Clean(name)
}
package lifecycle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func moduleCost(snapshot project.Snapshot, m project.Module, report checks.Report) (model.ModuleCost, error) {
	cost := model.ModuleCost{Module: m.Name, CasesPerSpec: map[string]int{}}
	for _, c := range report.Cases {
		cost.Seconds += c.Seconds
		seen := map[string]bool{}
		for _, id := range c.IDs {
			if !seen[id] {
				cost.CasesPerSpec[id]++
				seen[id] = true
			}
		}
	}
	err := moduleFiles(snapshot, m, func(name string, entry fs.DirEntry) error {
		if !entry.Type().IsRegular() {
			return nil
		}
		kind := zone(snapshot, m, name)
		switch kind {
		case "evals", "disposable":
			data, err := os.ReadFile(filepath.Join(snapshot.Root, name))
			if err != nil {
				return err
			}
			lines := bytes.Count(data, []byte{'\n'})
			if len(data) > 0 && data[len(data)-1] != '\n' {
				lines++
			}
			if kind == "evals" {
				cost.EvalLines += lines
			} else {
				cost.DisposableLines += lines
			}
		}
		return nil
	})
	return cost, err
}

func newCostSnapshot(state model.State, modules []model.ModuleCost, ds []project.Diagnostic) model.CostSnapshot {
	return model.CostSnapshot{Run: state.Run, Phase: state.Phase, RecordedAt: time.Now().UTC(), Modules: modules, Diagnostics: ds}
}

func costWarnings(snapshot project.Snapshot, modules []model.ModuleCost) []project.Diagnostic {
	budgets := configGroup(snapshot.Config, "budgets")
	timeLimit := 15 * time.Minute
	if value, err := time.ParseDuration(stringValue(budgets["verify_time"])); err == nil {
		timeLimit = value
	}
	caseLimit := intValue(budgets["cases_per_spec"], 15)
	ratioLimit := float64(4)
	if value, err := strconv.ParseFloat(fmt.Sprint(budgets["eval_ratio"]), 64); err == nil {
		ratioLimit = value
	}
	var ds []project.Diagnostic
	warn := func(code, path, message string) {
		ds = append(ds, project.Diagnostic{Code: code, Path: path, Warning: true, Message: message})
	}
	var seconds float64
	for _, m := range modules {
		seconds += m.Seconds
	}
	if seconds > timeLimit.Seconds() {
		warn("B01", "", fmt.Sprintf("evaluated modules took %gs, exceeding the %gs verify_time limit", seconds, timeLimit.Seconds()))
	}
	for _, m := range modules {
		for _, id := range sortedKeys(m.CasesPerSpec) {
			if count := m.CasesPerSpec[id]; count > caseLimit {
				warn("B02", m.Module, fmt.Sprintf("%s has %d testcases, exceeding the limit of %d", id, count, caseLimit))
			}
		}
		ratio := float64(0)
		if m.DisposableLines > 0 {
			ratio = float64(m.EvalLines) / float64(m.DisposableLines)
		} else if m.EvalLines > 0 {
			ratio = math.Inf(1)
		}
		if ratio > ratioLimit {
			warn("B03", m.Module, fmt.Sprintf("%s eval/disposable line ratio %g (%d/%d) exceeds the limit of %g", m.Module, ratio, m.EvalLines, m.DisposableLines, ratioLimit))
		}
	}
	return ds
}

func Costs(root string) ([]model.CostSnapshot, error) {
	repo, err := repository(root)
	if err != nil {
		return nil, err
	}
	files, err := filepath.Glob(filepath.Join(repo.Root, ".fireweed", "runs", "*", "*", "costs.json"))
	if err != nil {
		return nil, err
	}
	result := make([]model.CostSnapshot, 0, len(files))
	for _, file := range files {
		data, err := os.ReadFile(file)
		if errors.Is(err, fs.ErrNotExist) {
			continue
		}
		if err != nil {
			return nil, err
		}
		var snapshot model.CostSnapshot
		if err := json.Unmarshal(data, &snapshot); err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		result = append(result, snapshot)
	}
	sort.SliceStable(result, func(i, j int) bool {
		return result[i].RecordedAt.Before(result[j].RecordedAt)
	})
	return result, nil
}

func latestCosts(root string, modules []project.Module, restrict bool) ([]model.ModuleCost, []project.Diagnostic, error) {
	snapshots, err := Costs(root)
	if err != nil {
		return nil, nil, err
	}
	known := map[string]project.Module{}
	for _, m := range modules {
		known[m.Name] = m
	}
	costs := map[string]model.ModuleCost{}
	selected := map[string]int{}
	for i, snapshot := range snapshots {
		for _, cost := range snapshot.Modules {
			if _, ok := known[cost.Module]; ok || !restrict {
				costs[cost.Module] = cost
				selected[cost.Module] = i
			}
		}
	}
	var values []model.ModuleCost
	for _, module := range sortedKeys(costs) {
		values = append(values, costs[module])
	}
	var ds []project.Diagnostic
	seen := map[project.Diagnostic]bool{}
	for i, snapshot := range snapshots {
		for _, diagnostic := range snapshot.Diagnostics {
			if seen[diagnostic] || diagnostic.Code != "B01" && diagnostic.Code != "B02" && diagnostic.Code != "B03" {
				continue
			}
			for _, cost := range snapshot.Modules {
				index, ok := selected[cost.Module]
				if !ok || index != i || !costDiagnosticMatches(diagnostic, cost, known[cost.Module]) {
					continue
				}
				ds = append(ds, diagnostic)
				seen[diagnostic] = true
				break
			}
		}
	}
	return values, ds, nil
}

func costDiagnosticMatches(d project.Diagnostic, cost model.ModuleCost, module project.Module) bool {
	switch d.Code {
	case "B01":
		return true
	case "B02", "B03":
		if d.Path != "" {
			return d.Path == cost.Module || d.Path == module.Dir
		}
		if d.Code == "B03" {
			return strings.Contains(d.Message, cost.Module)
		}
		for id := range cost.CasesPerSpec {
			if strings.Contains(d.Message, id) {
				return true
			}
		}
	}
	return false
}

 succeeded in 581ms:
package lifecycle

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Block(request model.BlockRequest) (result model.BlockResult, err error) {
	repo, err := repository(request.Root)
	if err != nil {
		return result, err
	}
	root := repo.Root
	c, oid, err := heldClaim(repo, request.Change)
	if err != nil {
		return result, err
	}
	state, err := c.state(repo, oid)
	if err != nil {
		return result, err
	}
	if state.Head != state.PhaseBase {
		return result, coded("P05", c.Path, fmt.Sprintf("stale phase base: HEAD is %s, expected %s", state.Head, state.PhaseBase))
	}
	if _, err = approvedWorking(state); err != nil {
		return result, err
	}
	if strings.TrimSpace(request.Reason) == "" {
		return result, coded("P01", c.Path, "block requires a reason")
	}
	ctx := context.Background()
	attempt, cleanAttempt, err := temporaryIndex(ctx, root, state.PhaseBase)
	if err != nil {
		return result, err
	}
	defer cleanAttempt()
	if err = attempt.stageWorktree(); err != nil {
		return result, err
	}
	attemptOID, err := attempt.commit(state.PhaseBase, "fireweed: save "+c.Change+" "+state.Phase+" attempt")
	if err != nil {
		return result, err
	}
	prefix := attemptPrefix + c.Change + "/"
	refs, err := listRefs(ctx, root, prefix)
	if err != nil {
		return result, err
	}
	used := map[string]bool{}
	for _, ref := range refs {
		used[ref.name] = true
	}
	n := 1
	for used[prefix+strconv.Itoa(n)] {
		n++
	}
	ref := prefix + strconv.Itoa(n)
	state.Change.Document.Frontmatter["blocked"] = request.Reason
	approved, err := contracts.RenderDocument(state.Change.Document)
	if err != nil {
		return result, err
	}
	blocked, cleanBlocked, err := temporaryIndex(ctx, root, state.PhaseBase)
	if err != nil {
		return result, err
	}
	defer cleanBlocked()
	if err = blocked.stageDocument(c.Path, []byte(approved)); err != nil {
		return result, err
	}
	newHead, err := blocked.commit(state.PhaseBase, "fireweed: block "+c.Change)
	if err != nil {
		return result, err
	}
	// The attempt index describes the working tree, including new files.
	// Switching it to the block tree restores precisely the captured step.
	if err = attempt.checkout(newHead); err != nil {
		return result, err
	}
	updates := claimUpdate(c.ref(), oid, "") + fmt.Sprintf("create %s %s\n", ref, attemptOID)
	if err = commitRefs(root, state.PhaseBase, newHead, updates); err != nil {
		return result, combine(transitionError(root, state.PhaseBase, c.Path, err), attempt.checkout(attemptOID))
	}
	state.ClaimOID, state.Head, state.PhaseBase = "", newHead, newHead
	result = model.BlockResult{State: state, AttemptRef: ref, AttemptOID: attemptOID}
	return result, replaceIndex(root, newHead)
}

func resumeDocument(root, oid string, current project.Change) (string, error) {
	raw, err := gitstore.ReadAt(root, oid, current.Path)
	if err != nil {
		return "", withCode("P01", current.Path, err)
	}
	saved, err := readChange(current, raw)
	if err != nil {
		return "", withCode("P01", current.Path, err)
	}
	phase := field(current, "phase")
	if field(saved, "phase") != phase || field(saved, "run") != field(current, "run") {
		return "", coded("P01", current.Path, "saved attempt belongs to a different run or phase")
	}
	return contracts.RenderDocument(workerDocument(current.Document, saved.Document, phase))
}

func restoreAttempt(root, base, oid, file string) error {
	ctx := context.Background()
	patch, err := git(ctx, root, nil, nil, "diff", "--binary", "--full-index", "--no-renames", oid+"^", oid, "--", ".", ":(exclude,top,literal)"+file)
	if err != nil {
		return withCode("P01", file, err)
	}
	if len(patch) == 0 {
		return nil
	}
	if _, err = git(ctx, root, nil, patch, "apply", "--whitespace=nowarn", "-"); err == nil {
		return nil
	}

	// A later commit can edit the same file without changing the saved
	// hunk. Merge through a private index, then apply the resulting delta
	// to the checkout with the real index left untouched.
	idx, cleanup, err := temporaryIndex(ctx, root, base)
	if err != nil {
		return withCode("P01", file, err)
	}
	defer cleanup()
	if _, err = idx.run(patch, "apply", "--cached", "--3way", "--whitespace=nowarn", "-"); err != nil {
		return withCode("P01", file, err)
	}
	tree, err := idx.run(nil, "write-tree")
	if err != nil {
		return withCode("P01", file, err)
	}
	merged, err := git(ctx, root, nil, nil, "diff", "--binary", "--full-index", "--no-renames", base, strings.TrimSpace(string(tree)), "--")
	if err != nil || len(merged) == 0 {
		return withCode("P01", file, err)
	}
	_, err = git(ctx, root, nil, merged, "apply", "--whitespace=nowarn", "-")
	return withCode("P01", file, err)
}
package lifecycle

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/assets"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func Done(ctx context.Context, request model.Request) (result model.Result, err error) {
	request.Quick = false
	result, step, err := verify(ctx, request)
	if err != nil {
		return result, err
	}
	state := result.State
	root := step.repo.Root
	before, err := step.claim.state(step.repo, state.ClaimOID)
	if err != nil {
		return result, err
	}
	ch, err := approvedWorking(before)
	if err != nil {
		return result, err
	}
	if state.Phase == "audit" {
		ch = step.working
	}
	next, blocked := "impl", ""
	switch state.Phase {
	case "evals":
		if configGroup(state.Snapshot.Config, "audit")["enabled"] != false {
			next = "audit"
		}
	case "audit":
		ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Audit", step.audit.body())
		rounds := intValue(ch.Document.Frontmatter["audit_rounds"], 0)
		blocked = step.audit.blocked(rounds, intValue(configGroup(state.Snapshot.Config, "audit")["max_rounds"], 2))
		if blocked != "" {
			next = "audit"
			ch.Document.Frontmatter["blocked"] = blocked
			questions := strings.TrimRight(section(ch.Document, "Questions"), "\n")
			ch.Document.Body = contracts.ReplaceSection(ch.Document.Body, "Questions", questions+"\n"+step.audit.body())
		} else if len(step.audit.findings) != 0 {
			next = "evals"
			ch.Document.Frontmatter["audit_rounds"] = rounds + 1
		}
	case "impl":
		next = "applied"
		// Check the checkout again after hooks, before replacing any manifest.
		diagnostics := project.Check(root)
		if staticErr := firstError(diagnostics); staticErr != nil {
			result.Diagnostics = append(result.Diagnostics, diagnostics...)
			return result, staleError(staticErr)
		}
	}
	if state.Phase != "audit" {
		m, err := moduleNamed(state.Snapshot, ch.Module)
		if err != nil {
			return result, err
		}
		diagnostics, err := verifyZones(ctx, before, m)
		if err != nil {
			result.Diagnostics = append(result.Diagnostics, diagnostics...)
			return result, err
		}
	}
	ch.Document.Frontmatter["phase"] = next
	playbook, err := assets.Playbook(playbookName(state))
	if err != nil {
		return result, err
	}
	agent := strings.TrimSpace(request.Agent)
	if agent == "" {
		agent = "manual"
	}
	agent = strings.NewReplacer("\r", " ", "\n", " ").Replace(agent)
	tests, failing, passing := 0, 0, 0
	for _, report := range result.Reports {
		for _, c := range report.Cases {
			tests++
			switch c.Result {
			case "fail", "error":
				failing++
			case "pass":
				passing++
			}
		}
	}
	line := fmt.Sprintf("- %s %s agent=%s fireweed=%s playbook=%x tests=%d failing=%d passing=%d",
		state.Phase, time.Now().UTC().Format(time.RFC3339Nano), agent, assets.Version(),
		sha256.Sum256([]byte(playbook)), tests, failing, passing)
	ch.Document.Body = appendProvenance(ch.Document.Body, line)
	data, err := contracts.RenderDocument(ch.Document)
	if err != nil {
		return result, err
	}

	files := fileEdits{root: root}
	published := false
	defer func() {
		if err != nil && !published {
			err = combine(err, files.rollback())
		}
	}()
	if next == "applied" {
		if err = applyToModule(root, ch, &files); err != nil {
			return result, err
		}
		pending := ch.Path
		ch.Path = archivePath(ch.Path)
		if _, err = os.Lstat(filepath.Join(root, ch.Path)); err == nil {
			return result, fmt.Errorf("archive already exists: %s", ch.Path)
		} else if !errors.Is(err, fs.ErrNotExist) {
			return result, err
		}
		if err = files.remove(pending); err != nil {
			return result, err
		}
	}
	if err = files.write(ch.Path, []byte(data)); err != nil {
		return result, err
	}
	if next == "applied" {
		diagnostics := project.Check(root)
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		if err = firstError(diagnostics); err != nil {
			return result, staleError(err)
		}
	}
	idx, cleanup, err := temporaryIndex(ctx, root, state.PhaseBase)
	if err != nil {
		return result, err
	}
	defer cleanup()
	if err = idx.stageWorktree(); err != nil {
		return result, err
	}
	newHead, err := idx.commit(state.PhaseBase, "fireweed: "+state.Change.Slug+" "+state.Phase)
	if err != nil {
		return result, err
	}
	var newClaim string
	if next != "applied" && blocked == "" {
		updated := step.claim
		updated.Base, updated.Approved = newHead, data
		newClaim, err = updated.blob(root)
		if err != nil {
			return result, err
		}
	}
	if err = commitRefs(root, state.PhaseBase, newHead, claimUpdate(step.claim.ref(), state.ClaimOID, newClaim)); err != nil {
		return result, transitionError(root, state.PhaseBase, ch.Path, err)
	}
	published = true
	result.State.Change, result.State.Phase, result.State.ClaimOID = ch, next, newClaim
	result.State.Head, result.State.PhaseBase = newHead, newHead
	if err = replaceIndex(root, newHead); err != nil {
		return result, err
	}
	if blocked != "" {
		return result, coded("P06", ch.Path, "change blocked: "+blocked)
	}
	return result, nil
}

func playbookName(s model.State) string {
	if s.Phase == "impl" && s.Change.Document.Frontmatter["regenerate"] == true {
		return "regenerate"
	}
	return s.Phase
}

func nonblankLines(body string) []string {
	var lines []string
	for _, line := range strings.Split(body, "\n") {
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func appendProvenance(body, line string) string {
	prior, _ := contracts.Section(body, "Provenance")
	return contracts.ReplaceSection(body, "Provenance", "\n"+strings.Join(append(nonblankLines(prior), line), "\n")+"\n")
}

func applyToModule(root string, ch project.Change, files *fileEdits) error {
	snapshot, err := project.Load(root)
	if err != nil {
		return err
	}
	m, err := moduleNamed(snapshot, ch.Module)
	if err != nil {
		return err
	}
	doc := m.Document
	changed := false
	for key, value := range configGroup(ch.Document.Frontmatter, "set") {
		doc.Frontmatter[key] = value
		changed = true
	}
	if len(ch.Added)+len(ch.Modified)+len(ch.Removed) != 0 {
		items, err := specs(m)
		if err != nil {
			return err
		}
		doc.Body = contracts.ReplaceSection(doc.Body, "Specs", renderSpecEdits(items, ch))
		changed = true
	}
	added := strings.TrimSpace(section(ch.Document, "Decisions"))
	removed := strings.TrimSpace(section(ch.Document, "Removed decisions"))
	if added != "" || removed != "" {
		doc.Body = contracts.ReplaceSection(doc.Body, "Decisions", applyDecisions(section(doc, "Decisions"), added, removed, ch.Slug))
		changed = true
	}
	if !changed {
		return nil
	}
	rendered, err := contracts.RenderDocument(doc)
	if err != nil {
		return err
	}
	return files.write(path.Join(m.Dir, "SPEC.md"), []byte(rendered))
}

func decisionEntries(body string) []string {
	var entries, lines []string
	flush := func() {
		if len(lines) != 0 {
			entries = append(entries, strings.Join(lines, "\n"))
			lines = nil
		}
	}
	for _, line := range strings.Split(body, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			flush()
			continue
		}
		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") || strings.HasPrefix(line, "[") {
			flush()
		}
		lines = append(lines, line)
	}
	flush()
	return entries
}

func decisionText(entry string) string {
	var lines []string
	for _, line := range strings.Split(entry, "\n") {
		text := strings.TrimSpace(line)
		text = strings.TrimPrefix(text, "- ")
		text = strings.TrimPrefix(text, "* ")
		text = strings.TrimPrefix(text, "> ")
		lines = append(lines, text)
	}
	return strings.Trim(strings.Join(lines, "\n"), "\"")
}

func applyDecisions(prior, added, removed, slug string) string {
	removals := map[string]bool{}
	for _, entry := range decisionEntries(removed) {
		removals[decisionText(entry)] = true
	}
	var kept []string
	for _, entry := range decisionEntries(prior) {
		text := decisionText(entry)
		if removals[text] {
			continue
		}
		if strings.HasPrefix(entry, "[") {
			if _, rest, ok := strings.Cut(text, "] "); ok && removals[rest] {
				continue
			}
		}
		if strings.HasPrefix(entry, "- ") {
			if _, rest, ok := strings.Cut(text, ": "); ok && removals[rest] {
				continue
			}
		}
		kept = append(kept, entry)
	}
	for _, entry := range decisionEntries(added) {
		lines := strings.Split(decisionText(entry), "\n")
		kept = append(kept, "- "+slug+": "+strings.Join(lines, "\n  "))
	}
	return "\n" + strings.Join(kept, "\n\n") + "\n"
}

type fileEdits struct {
	root string
	undo []func() error
}

func (files *fileEdits) replace(name string, data []byte, remove bool) error {
	file := filepath.Join(files.root, name)
	before, err := os.ReadFile(file)
	exists := err == nil
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}
	if remove {
		err = os.Remove(file)
	} else {
		err = writeFile(file, data)
	}
	if err == nil {
		files.undo = append(files.undo, func() error {
			if exists {
				return writeFile(file, before)
			}
			return os.Remove(file)
		})
	}
	return err
}

func (files *fileEdits) write(name string, data []byte) error {
	return files.replace(name, data, false)
}
func (files *fileEdits) remove(name string) error { return files.replace(name, nil, true) }

func (files *fileEdits) rollback() error {
	var err error
	for i := len(files.undo) - 1; i >= 0; i-- {
		err = combine(err, files.undo[i]())
	}
	return err
}
package evals_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/workflow"
)

var gcNow = time.Date(2026, 11, 4, 12, 0, 0, 0, time.UTC)

func gcCall(t *testing.T, root string, now time.Time) workflow.GCResult {
	t.Helper()
	r, err := workflow.GC(context.Background(), workflow.GCRequest{Root: root, Now: now})
	if err != nil {
		t.Fatalf("GC: %v", err)
	}
	return r
}

func gcRef(t *testing.T, root, name, oid string) {
	t.Helper()
	if err := gitstore.CASRef(root, name, "", oid); err != nil {
		t.Fatalf("create %s: %v", name, err)
	}
}

func gcClaim(t *testing.T, root, name, worktree string) {
	t.Helper()
	data, err := json.Marshal(map[string]string{"worktree": worktree, "change": name})
	if err != nil {
		t.Fatal(err)
	}
	oid, err := gitstore.WriteBlob(root, data)
	if err != nil {
		t.Fatal(err)
	}
	gcRef(t, root, "refs/fireweed/claims/"+name, oid)
}

func gcPresent(t *testing.T, root, name string, want bool) {
	t.Helper()
	_, found, err := gitstore.Ref(root, name)
	if err != nil || found != want {
		t.Errorf("ref %s present=%t, error=%v; want present=%t", name, found, err, want)
	}
}

func gcNames(t *testing.T, root string) []string {
	t.Helper()
	out := git(t, root, "for-each-ref", "--format=%(refname)", "refs/fireweed")
	if out == "" {
		return nil
	}
	return strings.Split(out, "\n")
}

func gcPruned(t *testing.T, got, want []string) {
	t.Helper()
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("GC Pruned = %v, want sorted %v", got, want)
	}
}

func Test_FLOW_128_GCPrunesMissingAndUnregisteredClaimsOnly(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	linked := filepath.Join(t.TempDir(), "linked")
	git(t, root, "worktree", "add", "-b", "linked", linked, "HEAD")
	gcClaim(t, root, "main-live", root)
	gcClaim(t, root, "linked-live", linked)
	missing := filepath.Join(t.TempDir(), "removed")
	gcClaim(t, root, "missing", missing)
	unregistered := filepath.Join(t.TempDir(), "unregistered")
	if err := os.Mkdir(unregistered, 0o755); err != nil {
		t.Fatal(err)
	}
	gcClaim(t, root, "unregistered", unregistered)
	foreign, err := filepath.EvalSymlinks(t.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	git(t, foreign, "init")
	if got := git(t, foreign, "worktree", "list", "--porcelain"); !strings.Contains(got, "worktree "+foreign) {
		t.Fatalf("foreign repository does not register its worktree: %s", got)
	}
	gcClaim(t, root, "foreign", foreign)
	gcRef(t, root, "refs/fireweed/other/keep", head(t, root))

	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, []string{"refs/fireweed/claims/missing", "refs/fireweed/claims/unregistered", "refs/fireweed/claims/foreign"})
	for _, name := range []string{"missing", "unregistered", "foreign"} {
		gcPresent(t, root, "refs/fireweed/claims/"+name, false)
	}
	for _, name := range []string{"main-live", "linked-live"} {
		gcPresent(t, root, "refs/fireweed/claims/"+name, true)
	}
	gcPresent(t, root, "refs/fireweed/other/keep", true)
}

func Test_FLOW_129_GCPrunesOnlyArchivedChangeAttempts(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
	write(t, root, "changes/0003-pending.md", "---\nmodule: alpha\n---\n")
	oid := head(t, root)
	archived := []string{"refs/fireweed/attempts/finished/1", "refs/fireweed/attempts/finished/2"}
	for _, name := range append(archived, "refs/fireweed/attempts/pending/1", "refs/fireweed/attempts/unknown/1", "refs/fireweed/attempts/finished/not-an-attempt") {
		gcRef(t, root, name, oid)
	}
	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, archived)
	for _, name := range archived {
		gcPresent(t, root, name, false)
	}
	for _, name := range []string{"refs/fireweed/attempts/pending/1", "refs/fireweed/attempts/unknown/1", "refs/fireweed/attempts/finished/not-an-attempt"} {
		gcPresent(t, root, name, true)
	}
}

func Test_FLOW_130_FLOW_131_GCDrillAgeUsesConfiguredDurationAndNativeUTCDate(t *testing.T) {
	for _, tc := range []struct {
		name, config      string
		old, equal, newer string
	}{
		{"configured", "gc:\n  drill_max_age: 24h\n", "20261103T115959.999999999", "20261103T120000.000000000", "20261103T120000.000000001"},
		{"default", "", "20261005T115959.999999999", "20261005T120000.000000000", "20261005T120000.000000001"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t, tc.config, "", "", "", "")
			oid := head(t, root)
			prefix := "refs/fireweed/drills/alpha/"
			for _, date := range []string{tc.old, tc.equal, tc.newer} {
				gcRef(t, root, prefix+date, oid)
			}
			gcRef(t, root, prefix+"unrecognized", oid)
			gcRef(t, root, "refs/fireweed/drills/alpha/not-a-date/extra", oid)
			r := gcCall(t, root, gcNow)
			gcPruned(t, r.Pruned, []string{prefix + tc.old})
			gcPresent(t, root, prefix+tc.old, false)
			for _, name := range []string{prefix + tc.equal, prefix + tc.newer, prefix + "unrecognized", "refs/fireweed/drills/alpha/not-a-date/extra"} {
				gcPresent(t, root, name, true)
			}
		})
	}
	// With zero Now, a native date far in the past is expired against any
	// current clock, while a future date stays live.
	root := fixture(t, "", "", "", "", "")
	old := "refs/fireweed/drills/alpha/20200101T000000.000000000"
	future := "refs/fireweed/drills/alpha/20990101T000000.000000000"
	gcRef(t, root, old, head(t, root))
	gcRef(t, root, future, head(t, root))
	gcPruned(t, gcCall(t, root, time.Time{}).Pruned, []string{old})
	gcPresent(t, root, future, true)
}

func Test_FLOW_131_GCInvalidDurationFailsBeforeAnyPruning(t *testing.T) {
	for _, value := range []string{"0s", "-1h", "invalid"} {
		t.Run(value, func(t *testing.T) {
			root := fixture(t, "gc:\n  drill_max_age: "+value+"\n", "", "", "", "")
			claim := "refs/fireweed/claims/stale"
			old := "refs/fireweed/drills/alpha/20200101T000000.000000000"
			gcClaim(t, root, "stale", filepath.Join(t.TempDir(), "gone"))
			gcRef(t, root, old, head(t, root))
			before := gcNames(t, root)
			r, err := workflow.GC(context.Background(), workflow.GCRequest{Root: root, Now: gcNow})
			if err == nil || !strings.HasPrefix(err.Error(), "C01:") {
				t.Errorf("invalid drill_max_age %q: result=%#v error=%v; want C01", value, r, err)
			}
			if got := gcNames(t, root); !reflect.DeepEqual(got, before) {
				t.Errorf("invalid setting pruned refs: before=%v after=%v", before, got)
			}
			gcPresent(t, root, claim, true)
		})
	}
}

func Test_FLOW_132_GCReportsOnlyDeletedRefsAndPreservesRepositoryState(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	write(t, root, "changes/archive/0002-finished.md", "---\nmodule: alpha\nphase: applied\n---\n")
	write(t, root, "working.txt", "untracked work\n")
	write(t, root, "SYSTEM.md", "staged work\n")
	git(t, root, "add", "SYSTEM.md")
	gcClaim(t, root, "z-stale", filepath.Join(t.TempDir(), "missing"))
	gcRef(t, root, "refs/fireweed/attempts/finished/1", head(t, root))
	gcRef(t, root, "refs/fireweed/drills/alpha/20200101T000000.000000000", head(t, root))
	gcRef(t, root, "refs/fireweed/other/keep", head(t, root))
	want := []string{"refs/fireweed/claims/z-stale", "refs/fireweed/attempts/finished/1", "refs/fireweed/drills/alpha/20200101T000000.000000000"}
	beforeHead := head(t, root)
	beforeIndex := git(t, root, "write-tree")
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeWork := readFile(t, root, "working.txt")
	beforeStaged := readFile(t, root, "SYSTEM.md")
	r := gcCall(t, root, gcNow)
	gcPruned(t, r.Pruned, want)
	for _, name := range want {
		gcPresent(t, root, name, false)
	}
	gcPresent(t, root, "refs/fireweed/other/keep", true)
	if got := head(t, root); got != beforeHead {
		t.Errorf("GC moved HEAD: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "write-tree"); got != beforeIndex {
		t.Errorf("GC changed index tree: %s -> %s", beforeIndex, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("GC changed worktree status: %q -> %q", beforeStatus, got)
	}
	if readFile(t, root, "working.txt") != beforeWork || readFile(t, root, "SYSTEM.md") != beforeStaged {
		t.Error("GC changed working files")
	}
	if second := gcCall(t, root, gcNow); len(second.Pruned) != 0 {
		t.Errorf("second GC reported refs it did not delete: %v", second.Pruned)
	}
}

exec
/bin/zsh -lc "cat packages/workflow/model/model.go; sed -n '1,150p' packages/workflow/internal/lifecycle/verify.go; rg -n 'stageWorktree|GIT_INDEX_FILE|:(exclude' packages/workflow/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc "rg -n -A25 -B4 'func (git\\(|gitOutput|splitNUL|repository|writeFile)' packages/workflow/internal/lifecycle/common.go; rg -n 'cost|diagnostic' packages/workflow/evals/pr_body_test.go; sed -n '245,325p' packages/workflow/evals/status_test.go; rg -n 'func (verifyZones|verify\\()|\\.fireweed/runs' packages/workflow/internal/lifecycle/verify.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 exited 2 in 686ms:
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
package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/lurkshark/fireweed/packages/checks"
	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type verification struct {
	repo    gitstore.Repository
	claim   claim
	working project.Change
	audit   auditSubmission
}

func Verify(ctx context.Context, request model.Request) (model.Result, error) {
	result, _, err := verify(ctx, request)
	return result, err
}

func verify(ctx context.Context, request model.Request) (result model.Result, step verification, err error) {
	step.repo, err = repository(request.Root)
	if err != nil {
		return result, step, err
	}
	root := step.repo.Root
	var oid string
	step.claim, oid, err = heldClaim(step.repo, request.Change)
	if err != nil {
		return result, step, err
	}
	result.State, err = step.claim.state(step.repo, oid)
	if err != nil {
		return result, step, err
	}
	state := result.State
	switch state.Phase {
	case "evals", "audit", "impl":
	default:
		return result, step, coded("P01", state.Change.Path, "unsupported phase "+state.Phase)
	}
	if state.Head != state.PhaseBase {
		return result, step, coded("P05", state.Change.Path, fmt.Sprintf("stale phase base: HEAD is %s, expected %s", state.Head, state.PhaseBase))
	}
	step.working, err = approvedWorking(state)
	if err != nil {
		return result, step, err
	}
	if err = checkStale(root, step.working); err != nil {
		return result, step, err
	}
	result.State.Change = step.working
	m, err := moduleNamed(state.Snapshot, state.Change.Module)
	if err != nil {
		return result, step, err
	}
	result.Diagnostics, err = verifyZones(ctx, state, m)
	if err != nil {
		return result, step, err
	}
	evalFix := stringList(state.Change.Document.Frontmatter["eval_fix"])

	items, err := specs(m)
	if err != nil {
		return result, step, err
	}
	if state.Phase == "audit" {
		step.audit, err = checkAudit(state, step.working, items)
		return result, step, err
	}
	if err = validateRun(state.Run); err != nil {
		return result, step, err
	}
	runDir := filepath.Join(root, ".fireweed", "runs", state.Run, state.Phase)
	result.Reports = make(map[string]checks.Report)
	defer func() {
		if len(result.Costs) == 0 {
			return
		}
		diagnostics := costWarnings(state.Snapshot, result.Costs)
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		snapshot := newCostSnapshot(state, result.Costs, diagnostics)
		err = combine(err, writeJSON(filepath.Join(runDir, "costs.json"), snapshot))
	}()
	fullImpl := state.Phase == "impl" && !request.Quick
	modules := []project.Module{m}
	if fullImpl {
		mode := stringValue(state.Snapshot.Config["verify_dependents"])
		if mode == "" {
			mode = "direct"
		}
		for _, dependent := range state.Snapshot.Modules {
			if dependent.Name != m.Name && mode != "none" && dependsOn(state.Snapshot, dependent.Name, m.Name, mode == "all") {
				modules = append(modules, dependent)
			}
		}
	}

	var verificationErr error
	for i, module := range modules {
		moduleItems := items
		if i != 0 {
			moduleItems, err = specs(module)
			if err != nil {
				return result, step, combine(verificationErr, err)
			}
		}
		report, evalErr := checks.Evaluate(ctx, checks.HookRequest{
			Snapshot: state.Snapshot, Module: module.Name, Bin: request.Bin, Timeout: request.Timeout,
		})
		if evalErr != nil {
			return result, step, combine(verificationErr, evalErr)
		}
		result.Reports[module.Name] = report
		cost, costErr := moduleCost(state.Snapshot, module, report)
		if costErr == nil {
			result.Costs = append(result.Costs, cost)
		}
		verificationErr = combine(verificationErr, costErr)
		verificationErr = combine(verificationErr, writeJSON(filepath.Join(runDir, module.Name, "report.json"), report))
		var reportErr error
		if i == 0 {
			reportErr = checks.RequireCoverage(report, enforcedSpecs(moduleItems))
		}
		if reportErr == nil {
			switch {
			case state.Phase == "evals" && len(evalFix) != 0:
				for _, c := range report.Cases {
					if c.Result != "pass" {
						reportErr = coded("E01", module.Name, fmt.Sprintf("failed case %s (%s): %s", c.Name, c.Result, c.Failure))
						break
					}
				}
				if reportErr == nil {
					listed := make([]contracts.Item, 0, len(evalFix))
					for _, id := range evalFix {
						listed = append(listed, contracts.Item{ID: id})
					}
					reportErr = checks.RequireCoverage(report, listed)
				}
			case state.Phase == "evals":
				reportErr = checks.RequireFailFirst(report, changeIDs(state.Change, moduleItems, false), stringList(state.Change.Document.Frontmatter["already_satisfied"]))
rg: regex parse error:
    (?:stageWorktree|GIT_INDEX_FILE|:(exclude)
    ^
error: unclosed group

 succeeded in 699ms:
29-	attemptPrefix = "refs/fireweed/attempts/"
30-	drillPrefix   = "refs/fireweed/drills/"
31-)
32-
33:func git(ctx context.Context, root string, env []string, input []byte, args ...string) ([]byte, error) {
34-	cmd := exec.CommandContext(ctx, "git", args...)
35-	cmd.Dir = root
36-	cmd.Env = append(os.Environ(), env...)
37-	cmd.Stdin = bytes.NewReader(input)
38-	var stderr bytes.Buffer
39-	cmd.Stderr = &stderr
40-	out, err := cmd.Output()
41-	if err != nil {
42-		if ctx.Err() != nil {
43-			err = ctx.Err()
44-		}
45-		return nil, fmt.Errorf("git %s: %w: %s", args[0], err, strings.TrimSpace(stderr.String()))
46-	}
47-	return out, nil
48-}
49-
50:func gitOutput(root string, args ...string) (string, error) {
51-	out, err := git(context.Background(), root, nil, nil, args...)
52-	return strings.TrimSpace(string(out)), err
53-}
54-
55:func repository(root string) (gitstore.Repository, error) {
56-	repo, err := gitstore.Inspect(root)
57-	if err == nil {
58-		repo.Root, err = filepath.EvalSymlinks(repo.Root)
59-	}
60-	if err == nil {
61-		repo.Root, err = filepath.Abs(repo.Root)
62-	}
63-	return repo, err
64-}
65-
66-func coded(code, subject, message string) error {
67-	return contracts.CodedError{Code: code, Subject: subject, Cause: errors.New(code + ": " + message)}
68-}
69-
70-func withCode(code, subject string, err error) error {
71-	if err == nil {
72-		return nil
73-	}
74-	if !strings.HasPrefix(err.Error(), code+":") {
75-		return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s: %w", code, err)}
76-	}
77-	if existing, _, ok := contracts.ErrorCode(err); ok && existing == code {
78-		return err
79-	}
80-	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
--
290-	}
291-	return result, nil
292-}
293-
294:func splitNUL(out []byte) []string {
295-	if len(out) == 0 {
296-		return nil
297-	}
298-	return strings.Split(strings.TrimSuffix(string(out), "\x00"), "\x00")
299-}
300-
301-func selectChange(queue []project.Change, selector string) (project.Change, error) {
302-	var matches []project.Change
303-	slugs := map[string]bool{}
304-	clean := filepath.ToSlash(filepath.Clean(selector))
305-	for _, ch := range queue {
306-		slugs[ch.Slug] = true
307-		name := path.Base(ch.Path)
308-		if clean == ch.Path || selector == name || selector == ch.Slug || selector == strings.TrimSuffix(name, ".md") {
309-			matches = append(matches, ch)
310-		}
311-	}
312-	if len(matches) == 1 {
313-		return matches[0], nil
314-	}
315-	reason := "not found"
316-	if len(matches) > 1 {
317-		reason = "ambiguous"
318-	}
319-	return project.Change{}, coded("U02", selector, fmt.Sprintf("change %q %s (valid slugs: %s)", selector, reason, strings.Join(sortedKeys(slugs), ", ")))
--
346-	}
347-	return ch, nil
348-}
349-
350:func writeFile(file string, data []byte) error {
351-	if err := os.MkdirAll(filepath.Dir(file), 0755); err != nil {
352-		return err
353-	}
354-	tmp, err := os.CreateTemp(filepath.Dir(file), ".fireweed-write-")
355-	if err != nil {
356-		return err
357-	}
358-	defer os.Remove(tmp.Name())
359-	mode := fs.FileMode(0644)
360-	if info, err := os.Stat(file); err == nil {
361-		mode = info.Mode().Perm()
362-	}
363-	if err := tmp.Chmod(mode); err != nil {
364-		tmp.Close()
365-		return err
366-	}
367-	_, writeErr := tmp.Write(data)
368-	if err := errors.Join(writeErr, tmp.Close()); err != nil {
369-		return err
370-	}
371-	return os.Rename(tmp.Name(), file)
372-}
373-
374-func writeJSON(file string, data any) error {
375-	b, err := json.MarshalIndent(data, "", "  ")
200:	write(t, root, filepath.Join(".fireweed/runs", run, phase, "costs.json"), string(b))
224:	t.Run("FLOW_123_latest_saved_costs", func(t *testing.T) {
245:					t.Errorf("unexpected or stale cost measurement row %q", row)
251:					t.Errorf("missing saved %s cost measurement row: %s", module, section[1])
257:				t.Errorf("PRBody diagnostics lack saved %s: %#v", code, r.Diagnostics)
319:					t.Errorf("unexpected B04 diagnostic: %#v", d)
	r := status(t, root)
	if got := r.PhaseCommits["0001-manual"]; got != len(commits) {
		t.Errorf("phase commits = %#v, want %d for 0001-manual", r.PhaseCommits, len(commits))
	}
}

func Test_FLOW_113_StatusLatestRecordedCostsAndSavedWarnings(t *testing.T) {
	root := fixture(t, "audit:\n  enabled: false\nbudgets:\n  eval_ratio: 1\n", "", "", "", "")
	write(t, root, "packages/alpha/evals/weight.txt", "one\ntwo\n")
	commit(t, root, "cost fixture")
	start(t, root)
	if _, err := done(root); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/internal/ready", "ready\n")
	if _, err := verify(root, false); err != nil {
		t.Fatal(err)
	}
	snapshots, err := workflow.Costs(root)
	if err != nil {
		t.Fatal(err)
	}
	latest := map[string]workflow.ModuleCost{}
	var latestAlphaB03, olderAlphaB03 project.Diagnostic
	var hasLatestAlphaB03, hasOlderAlphaB03 bool
	for _, s := range snapshots {
		for _, c := range s.Modules {
			latest[c.Module] = c
			if c.Module == "alpha" {
				olderAlphaB03, hasOlderAlphaB03 = latestAlphaB03, hasLatestAlphaB03
				latestAlphaB03, hasLatestAlphaB03 = warningFor(s.Diagnostics, "B03")
			}
		}
	}
	if !hasOlderAlphaB03 || !hasLatestAlphaB03 || olderAlphaB03 == latestAlphaB03 {
		t.Fatalf("cost snapshots need distinct older and latest alpha B03 warnings: %#v", snapshots)
	}
	config := readFile(t, root, "fireweed.yaml")
	write(t, root, "fireweed.yaml", strings.Replace(config, "eval_ratio: 1", "eval_ratio: 100", 1))
	r := status(t, root)
	if len(r.Costs) != 2 {
		t.Errorf("cost modules = %#v, want only measured alpha and beta", r.Costs)
	}
	for _, c := range r.Costs {
		if !reflect.DeepEqual(c, latest[c.Module]) {
			t.Errorf("latest %s cost = %#v, want %#v", c.Module, c, latest[c.Module])
		}
	}
	for _, module := range []string{"alpha", "beta"} {
		want := latest[module]
		got, ok := costFor(r.Costs, module)
		if !ok || !reflect.DeepEqual(got, want) {
			t.Errorf("recorded %s cost = %#v (present %t), want %#v", module, got, ok, want)
		}
	}
	if _, ok := costFor(r.Costs, "gamma"); ok {
		t.Errorf("unmeasured gamma has cost: %#v", r.Costs)
	}
	if d, ok := warningFor(r.Diagnostics, "B03"); !ok || !d.Warning || d != latestAlphaB03 {
		t.Errorf("selected B03 warning = %#v, want %#v", r.Diagnostics, latestAlphaB03)
	}
	if containsDiagnostic(r.Diagnostics, olderAlphaB03) {
		t.Errorf("older B03 warning returned: %#v", r.Diagnostics)
	}
	count := 0
	for _, d := range r.Diagnostics {
		if d.Code == "B03" {
			count++
		}
	}
	if count != 1 {
		t.Errorf("selected B03 warning count = %d, want 1: %#v", count, r.Diagnostics)
	}
}

func containsDiagnostic(ds []project.Diagnostic, want project.Diagnostic) bool {
	for _, d := range ds {
		if d == want {
			return true
		}
	}
29:func verify(ctx context.Context, request model.Request) (result model.Result, step verification, err error) {
179:func verifyZones(ctx context.Context, state model.State, m project.Module) ([]project.Diagnostic, error) {
202:		if !strings.HasPrefix(change.Path, ".fireweed/runs/") || change.OldPath != "" || existsAt(root, state.PhaseBase, change.Path) {

exec
/bin/zsh -lc "sed -n '179,250p' packages/workflow/internal/lifecycle/verify.go; sed -n '185,275p' packages/workflow/evals/pr_body_test.go; rg -n 'stageWorktree|GIT_INDEX_FILE|exclude,top' packages/workflow/internal; cat packages/project/model/model.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 559ms:
func verifyZones(ctx context.Context, state model.State, m project.Module) ([]project.Diagnostic, error) {
	root := state.Snapshot.Root
	changes, err := gitstore.Changes(root, state.PhaseBase)
	if err != nil {
		return nil, err
	}
	// An index entry may differ even when its working copy matches the base.
	index, err := diffPaths(ctx, root, "diff", "--cached", "--name-status", "-z", "--no-renames", state.PhaseBase, "--")
	if err != nil {
		return nil, err
	}
	known := map[string]bool{}
	for _, change := range changes {
		known[change.Path] = true
		known[change.OldPath] = true
	}
	for _, change := range index {
		if !known[change.Path] {
			changes = append(changes, change)
		}
	}
	phaseChanges := changes[:0]
	for _, change := range changes {
		if !strings.HasPrefix(change.Path, ".fireweed/runs/") || change.OldPath != "" || existsAt(root, state.PhaseBase, change.Path) {
			phaseChanges = append(phaseChanges, change)
		}
	}
	changes = phaseChanges
	var diagnostics []project.Diagnostic
	evalFix := stringList(state.Change.Document.Frontmatter["eval_fix"])
	if len(evalFix) != 0 {
		for _, change := range changes {
			for _, file := range []string{change.Path, change.OldPath} {
				if file == "" || file == state.Change.Path || state.Phase == "evals" && belongsTo(state.Snapshot, m, file) && zone(state.Snapshot, m, file) == "evals" {
					continue
				}
				diagnostics = append(diagnostics, project.Diagnostic{
					Code: "Z01", Path: file, Message: fmt.Sprintf("%s is forbidden during eval_fix %s", file, state.Phase),
				})
			}
		}
	}
	diagnostics = append(diagnostics, checks.Zones(checks.ZoneRequest{
		Snapshot: state.Snapshot, Module: m.Name, Phase: state.Phase,
		ChangePath: state.Change.Path, Base: state.PhaseBase, Changes: changes,
		InterfaceDeclared: strings.TrimSpace(section(state.Change.Document, "Interface")) != "",
	})...)
	return diagnostics, firstError(diagnostics)
}

func combine(first, second error) error {
	if first == nil {
		return second
	}
	if second == nil {
		return first
	}
	return errors.Join(first, second)
}

func requireMappedPass(report checks.Report, items []contracts.Item) error {
	selected := make(map[string]bool, len(items))
	for _, item := range items {
		selected[item.ID] = true
	}
	var failures checks.Report
	for _, testcase := range report.Cases {
		if testcase.Result != "fail" && testcase.Result != "error" {
			continue
		}
		var ids []string
		for _, id := range testcase.IDs {
			}
		}
	})
	t.Run("FLOW_121_decisions", func(t *testing.T) {
		r := prBody(t, root, full)
		prRequire(t, r.Body, "Use the revised policy.", "Retain old policy.")
	})
}

func prSaveCost(t *testing.T, root, run, phase string, snapshot workflow.CostSnapshot) {
	t.Helper()
	b, err := json.Marshal(snapshot)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, filepath.Join(".fireweed/runs", run, phase, "costs.json"), string(b))
}

func Test_FLOW_122_FLOW_123_PRBodyEndpointExceptionsAndRetainedCosts(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	alpha := readFile(t, root, "packages/alpha/SPEC.md")
	alpha = strings.Replace(alpha, "depends_on: []\n", "depends_on: []\nboundaries_opt_out: ENDPOINT OPT OUT REASON\n", 1)
	alpha = strings.Replace(alpha, "### ALP-002\n", "### ALP-002 [quarantine: ENDPOINT QUARANTINE REASON]\n", 1)
	write(t, root, "packages/alpha/SPEC.md", alpha)
	manual := readFile(t, root, "changes/0001-manual.md")
	write(t, root, "changes/0001-manual.md", strings.Replace(manual, "module: alpha\n", "module: alpha\nblocked: ENDPOINT BLOCK REASON\n", 1))
	commit(t, root, "endpoint exceptions")
	endpoint := head(t, root)
	old := workflow.CostSnapshot{Run: "old", Phase: "evals", RecordedAt: time.Date(2026, 10, 29, 0, 0, 0, 0, time.UTC), Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 99, CasesPerSpec: map[string]int{"ALP-001": 99}, EvalLines: 99, DisposableLines: 1}}, Diagnostics: []project.Diagnostic{{Code: "B01", Message: "OLD B01", Warning: true}}}
	latest := workflow.CostSnapshot{Run: "new", Phase: "impl", RecordedAt: time.Date(2026, 10, 30, 0, 0, 0, 0, time.UTC), Modules: []workflow.ModuleCost{{Module: "alpha", Seconds: 4.25, CasesPerSpec: map[string]int{"ALP-001": 5}, EvalLines: 17, DisposableLines: 3}, {Module: "beta", Seconds: 2.5, CasesPerSpec: map[string]int{"BET-001": 2}, EvalLines: 8, DisposableLines: 4}}, Diagnostics: []project.Diagnostic{{Code: "B01", Message: "SAVED B01 time", Warning: true}, {Code: "B02", Message: "ALP-001: SAVED B02 cases", Warning: true}, {Code: "B03", Message: "alpha: SAVED B03 ratio", Warning: true}}}
	prSaveCost(t, root, "old", "evals", old)
	prSaveCost(t, root, "new", "impl", latest)
	write(t, root, "fireweed.yaml", readFile(t, root, "fireweed.yaml")+"\nbudgets:\n  cases_per_spec: 1000\n  eval_ratio: 1000\n")
	write(t, root, "packages/alpha/evals/later.txt", strings.Repeat("later\n", 50))
	t.Run("FLOW_122_endpoint_exceptions", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "alpha", "ALP-002", "ENDPOINT QUARANTINE REASON", "ENDPOINT OPT OUT REASON", "0001-manual", "ENDPOINT BLOCK REASON")
	})
	t.Run("FLOW_123_latest_saved_costs", func(t *testing.T) {
		r := prBody(t, root, base+".."+endpoint)
		prRequire(t, r.Body, "SAVED B01", "SAVED B02", "SAVED B03")
		prReject(t, r.Body, "OLD B01")
		section := regexp.MustCompile(`(?ms)^#{1,6} (?:Cost|Costs|Metric|Metrics)\s*\n(.*?)(?:^#{1,6} |\z)`).FindStringSubmatch(r.Body)
		if len(section) != 2 {
			t.Errorf("PR body lacks a Cost/Metric section: %s", r.Body)
		} else {
			rows := regexp.MustCompile(`(?m)^\|\s*(?:alpha|beta|gamma)\s*\|[^\n]*$`).FindAllString(section[1], -1)
			want := map[string]string{
				"alpha": "alpha|4.25|ALP-001: 5|17|3",
				"beta":  "beta|2.5|BET-001: 2|8|4",
			}
			seen := map[string]bool{}
			for _, row := range rows {
				cells := strings.Split(strings.Trim(strings.TrimSpace(row), "|"), "|")
				for i := range cells {
					cells[i] = strings.TrimSpace(cells[i])
				}
				module := cells[0]
				if seen[module] || strings.Join(cells, "|") != want[module] {
					t.Errorf("unexpected or stale cost measurement row %q", row)
				}
				seen[module] = true
			}
			for module := range want {
				if !seen[module] {
					t.Errorf("missing saved %s cost measurement row: %s", module, section[1])
				}
			}
		}
		for _, code := range []string{"B01", "B02", "B03"} {
			if !containsCode(r.Diagnostics, code) {
				t.Errorf("PRBody diagnostics lack saved %s: %#v", code, r.Diagnostics)
			}
		}
	})
}

func Test_FLOW_126_PRBodyEndpointConfigError(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	base := head(t, root)
	write(t, root, "fireweed.yaml", "[invalid endpoint configuration\n")
	commit(t, root, "invalid endpoint configuration")

	_, err := workflow.PRBody(root, base+".."+head(t, root))
	if err == nil {
		t.Fatal("PRBody accepted malformed endpoint configuration")
	}
	if !workflow.IsEndpointConfigError(err) {
		t.Errorf("IsEndpointConfigError(%v) = false, want true", err)
	}
packages/workflow/internal/lifecycle/done.go:147:	if err = idx.stageWorktree(); err != nil {
packages/workflow/internal/lifecycle/attempt.go:44:	if err = attempt.stageWorktree(); err != nil {
packages/workflow/internal/lifecycle/attempt.go:114:	patch, err := git(ctx, root, nil, nil, "diff", "--binary", "--full-index", "--no-renames", oid+"^", oid, "--", ".", ":(exclude,top,literal)"+file)
packages/workflow/internal/lifecycle/git.go:70:	idx := gitIndex{root: root, env: []string{"GIT_INDEX_FILE=" + filepath.Join(dir, "index")}, ctx: ctx}
packages/workflow/internal/lifecycle/git.go:82:func (idx gitIndex) stageWorktree() error {
packages/workflow/internal/lifecycle/git.go:83:	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
packages/workflow/internal/lifecycle/history.go:303:		args = append(args, rev.oid, "--", ".", ":(exclude,top,literal)"+ch.Path, ":(exclude,top,literal)"+archivePath(ch.Path))
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

codex
The repeated Done and Block failures point to the same staging command: Git is rejecting the ignored `.fireweed/runs` directory. GC also uses a flag the verifier’s Git doesn’t support. I’ll address both in workflow and check how Status matches saved warnings to modules.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/packages/workflow/internal/lifecycle/git.go
diff --git a/packages/workflow/internal/lifecycle/git.go b/packages/workflow/internal/lifecycle/git.go
index bddbd3a3e0cc97817da080b329fff28bc310eb31..3db167c7ec28aed7e6790ae2cd7f13a429d7987c
--- a/packages/workflow/internal/lifecycle/git.go
+++ b/packages/workflow/internal/lifecycle/git.go
@@ -1,10 +1,14 @@
 package lifecycle
 
 import (
+	"bytes"
 	"context"
+	"errors"
 	"fmt"
 	"os"
+	"os/exec"
 	"path/filepath"
+	"strconv"
 	"strings"
 
 	"github.com/lurkshark/fireweed/packages/gitstore"
@@ -30,6 +34,42 @@
 	return refs, nil
 }
 
+func registeredWorktrees(ctx context.Context, root string) ([]string, error) {
+	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	var records []string
+	quoted := false
+	if err == nil {
+		records = splitNUL(out)
+	} else {
+		var exit *exec.ExitError
+		if !errors.As(err, &exit) || exit.ExitCode() != 129 {
+			return nil, err
+		}
+		// Older Git versions support porcelain records but lack -z.
+		out, err = git(ctx, root, nil, nil, "worktree", "list", "--porcelain")
+		if err != nil {
+			return nil, err
+		}
+		records = strings.Split(string(out), "\n")
+		quoted = true
+	}
+	var paths []string
+	for _, record := range records {
+		name, ok := strings.CutPrefix(record, "worktree ")
+		if !ok {
+			continue
+		}
+		if quoted && strings.HasPrefix(name, `"`) {
+			name, err = strconv.Unquote(name)
+			if err != nil {
+				return nil, fmt.Errorf("invalid quoted worktree path: %w", err)
+			}
+		}
+		paths = append(paths, name)
+	}
+	return paths, nil
+}
+
 func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
 	out, err := git(ctx, root, nil, nil, args...)
 	if err != nil {
@@ -80,7 +120,24 @@
 }
 
 func (idx gitIndex) stageWorktree() error {
-	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
+	// Include tracked deletions and non-ignored new files. Giving older Git
+	// a broad add pathspec with an exclusion can still hit ignored directories.
+	out, err := idx.run(nil, "ls-files", "--cached", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return err
+	}
+	var paths bytes.Buffer
+	for _, name := range splitNUL(out) {
+		if name == ".fireweed/runs" || strings.HasPrefix(name, ".fireweed/runs/") {
+			continue
+		}
+		paths.WriteString(name)
+		paths.WriteByte(0)
+	}
+	if paths.Len() == 0 {
+		return nil
+	}
+	_, err = idx.run(paths.Bytes(), "--literal-pathspecs", "add", "-A", "--pathspec-from-file=-", "--pathspec-file-nul")
 	return err
 }
 

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/packages/workflow/internal/lifecycle/status.go
diff --git a/packages/workflow/internal/lifecycle/git.go b/packages/workflow/internal/lifecycle/git.go
index bddbd3a3e0cc97817da080b329fff28bc310eb31..3db167c7ec28aed7e6790ae2cd7f13a429d7987c
--- a/packages/workflow/internal/lifecycle/git.go
+++ b/packages/workflow/internal/lifecycle/git.go
@@ -1,10 +1,14 @@
 package lifecycle
 
 import (
+	"bytes"
 	"context"
+	"errors"
 	"fmt"
 	"os"
+	"os/exec"
 	"path/filepath"
+	"strconv"
 	"strings"
 
 	"github.com/lurkshark/fireweed/packages/gitstore"
@@ -30,6 +34,42 @@
 	return refs, nil
 }
 
+func registeredWorktrees(ctx context.Context, root string) ([]string, error) {
+	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	var records []string
+	quoted := false
+	if err == nil {
+		records = splitNUL(out)
+	} else {
+		var exit *exec.ExitError
+		if !errors.As(err, &exit) || exit.ExitCode() != 129 {
+			return nil, err
+		}
+		// Older Git versions support porcelain records but lack -z.
+		out, err = git(ctx, root, nil, nil, "worktree", "list", "--porcelain")
+		if err != nil {
+			return nil, err
+		}
+		records = strings.Split(string(out), "\n")
+		quoted = true
+	}
+	var paths []string
+	for _, record := range records {
+		name, ok := strings.CutPrefix(record, "worktree ")
+		if !ok {
+			continue
+		}
+		if quoted && strings.HasPrefix(name, `"`) {
+			name, err = strconv.Unquote(name)
+			if err != nil {
+				return nil, fmt.Errorf("invalid quoted worktree path: %w", err)
+			}
+		}
+		paths = append(paths, name)
+	}
+	return paths, nil
+}
+
 func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
 	out, err := git(ctx, root, nil, nil, args...)
 	if err != nil {
@@ -80,7 +120,24 @@
 }
 
 func (idx gitIndex) stageWorktree() error {
-	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
+	// Include tracked deletions and non-ignored new files. Giving older Git
+	// a broad add pathspec with an exclusion can still hit ignored directories.
+	out, err := idx.run(nil, "ls-files", "--cached", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return err
+	}
+	var paths bytes.Buffer
+	for _, name := range splitNUL(out) {
+		if name == ".fireweed/runs" || strings.HasPrefix(name, ".fireweed/runs/") {
+			continue
+		}
+		paths.WriteString(name)
+		paths.WriteByte(0)
+	}
+	if paths.Len() == 0 {
+		return nil
+	}
+	_, err = idx.run(paths.Bytes(), "--literal-pathspecs", "add", "-A", "--pathspec-from-file=-", "--pathspec-file-nul")
 	return err
 }
 
diff --git a/packages/workflow/internal/lifecycle/status.go b/packages/workflow/internal/lifecycle/status.go
index 0acd60803368b6ef9fd7454eb9f6196d7b602a28..1fa84461b966a90ede920db86a62a638c7f90946
--- a/packages/workflow/internal/lifecycle/status.go
+++ b/packages/workflow/internal/lifecycle/status.go
@@ -157,17 +157,16 @@
 	for _, ch := range archives {
 		archived[ch.Slug] = true
 	}
-	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	worktrees, err := registeredWorktrees(ctx, root)
 	if err != nil {
 		return result, err
 	}
 	live := map[string]bool{}
 	common := canonicalPath(repo.CommonDir, root)
-	for _, line := range splitNUL(out) {
-		if name, ok := strings.CutPrefix(line, "worktree "); ok {
-			if other, err := repository(name); err == nil && canonicalPath(other.CommonDir, other.Root) == common {
-				live[other.Root] = true
-			}
+	for _, name := range worktrees {
+		if other, err := repository(name); err == nil &&
+			other.Root == canonicalPath(name, root) && canonicalPath(other.CommonDir, other.Root) == common {
+			live[other.Root] = true
 		}
 	}
 	refs, err := listRefs(ctx, root, claimPrefix, attemptPrefix, drillPrefix)

exec
/bin/zsh -lc "rg -n -A35 -B6 'func (belongsTo|ownerModule|zone|moduleFiles)' packages/workflow/internal/lifecycle/common.go; sed -n '145,182p' packages/workflow/internal/lifecycle/verify.go; sed -n '20,100p' packages/workflow/evals/lifecycle_test.go; sed -n '1,150p' packages/workflow/internal/lifecycle/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 672ms:
199-			return true
200-		}
201-	}
202-	return false
203-}
204-
205:func zone(s project.Snapshot, m project.Module, name string) string {
206-	if matchesAny(stringList(s.Config["protected"]), name) {
207-		return "protected"
208-	}
209-	relative, inside := relativeTo(m.Dir, name)
210-	if inside {
211-		return contracts.ResolveZone(relative, m.Zones)
212-	}
213-	return ""
214-}
215-
216:func belongsTo(s project.Snapshot, m project.Module, name string) bool {
217-	if _, inside := relativeTo(m.Dir, name); !inside {
218-		return false
219-	}
220-	for _, other := range s.Modules {
221-		if other.Dir != m.Dir && len(other.Dir) > len(m.Dir) {
222-			if _, inside := relativeTo(other.Dir, name); inside {
223-				return false
224-			}
225-		}
226-	}
227-	return true
228-}
229-
230:func moduleFiles(s project.Snapshot, m project.Module, visit func(string, fs.DirEntry) error) error {
231-	dir := filepath.Join(s.Root, filepath.FromSlash(m.Dir))
232-	return filepath.WalkDir(dir, func(file string, entry fs.DirEntry, err error) error {
233-		if err != nil {
234-			return err
235-		}
236-		relative, err := filepath.Rel(s.Root, file)
237-		if err != nil {
238-			return err
239-		}
240-		name := filepath.ToSlash(relative)
241-		if entry.IsDir() {
242-			if file != dir && (entry.Name() == ".git" || entry.Name() == ".fireweed" || !belongsTo(s, m, name)) {
243-				return filepath.SkipDir
244-			}
245-			return nil
246-		}
247-		return visit(name, entry)
248-	})
249-}
250-
251-var numberedName = regexp.MustCompile(`^([0-9]+)-(.+)$`)
252-
253-func identity(file string) project.Change {
254-	stem := strings.TrimSuffix(path.Base(file), ".md")
255-	change := project.Change{Path: file, Slug: stem}
256-	if parts := numberedName.FindStringSubmatch(stem); parts != nil {
257-		change.Number, _ = strconv.Atoi(parts[1])
258-		change.Slug = parts[2]
259-	}
260-	return change
261-}
262-
263-func changePaths(root, dir string) ([]project.Change, error) {
264-	entries, err := os.ReadDir(filepath.Join(root, dir))
265-	if errors.Is(err, fs.ErrNotExist) {
						listed = append(listed, contracts.Item{ID: id})
					}
					reportErr = checks.RequireCoverage(report, listed)
				}
			case state.Phase == "evals":
				reportErr = checks.RequireFailFirst(report, changeIDs(state.Change, moduleItems, false), stringList(state.Change.Document.Frontmatter["already_satisfied"]))
			case i == 0:
				reportErr = requireMappedPass(report, nonQuarantinedSpecs(moduleItems))
			default:
				reportErr = requireMappedPass(report, enforcedSpecs(moduleItems))
			}
		}
		verificationErr = combine(verificationErr, reportErr)
	}
	if fullImpl {
		diagnostics, boundaryErr := checks.Boundaries(ctx, checks.BoundaryRequest{
			Snapshot: state.Snapshot, Bin: request.Bin, Timeout: request.Timeout,
		})
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		if boundaryErr == nil {
			boundaryErr = firstError(diagnostics)
		} else {
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && diagnostic.Code != "" {
					boundaryErr = withCode(diagnostic.Code, diagnostic.Path, boundaryErr)
					break
				}
			}
		}
		verificationErr = combine(verificationErr, boundaryErr)
	}
	return result, step, verificationErr
}

func verifyZones(ctx context.Context, state model.State, m project.Module) ([]project.Diagnostic, error) {
	root := state.Snapshot.Root
	changes, err := gitstore.Changes(root, state.PhaseBase)
	if err != nil {
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

const oldClause = "Alpha MUST return the old value."
const newClause = "Alpha MUST return the revised value."

func write(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}

func git(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	cmd.Env = append(os.Environ(), "GIT_CONFIG_NOSYSTEM=1", "GIT_AUTHOR_NAME=Fixture", "GIT_AUTHOR_EMAIL=fixture@example.invalid", "GIT_COMMITTER_NAME=Fixture", "GIT_COMMITTER_EMAIL=fixture@example.invalid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func head(t *testing.T, root string) string { t.Helper(); return git(t, root, "rev-parse", "HEAD") }
func commit(t *testing.T, root, message string) {
	t.Helper()
	git(t, root, "add", "-A")
	git(t, root, "commit", "-m", message)
}

func fixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
	t.Helper()
	root := t.TempDir()
	var err error
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	write(t, root, "SYSTEM.md", "# Fixture system\nLocal lifecycle fixture.\n")
	write(t, root, ".gitignore", ".fireweed/runs/\n")
	write(t, root, "fireweed.yaml", `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
hooks:
  eval: [sh, .fireweed/eval.sh, "{module}", "{report}"]
  boundaries: [sh, .fireweed/boundaries.sh]
`+configExtra)
	write(t, root, ".fireweed/eval.sh", `#!/bin/sh
printf 'fixture-hook-output\n'
module="$1"
report="$2"
if [ "$module" = alpha ]; then
  if [ -f packages/alpha/internal/ready ]; then result=''; exitcode=0
  else result='<failure message="expected before implementation"/>'; exitcode=1; fi
  if [ -f .fireweed/unrelated ]; then other='<failure message="known unrelated"/>'; exitcode=1
  else other=''; fi
  printf '<testsuite><testcase name="ALP-001" classname="fixture">%s</testcase><testcase name="ALP-002" classname="fixture">%s</testcase></testsuite>' "$result" "$other" > "$report"
  exit "$exitcode"
fi
if [ "$module" = beta ] && [ -f .fireweed/beta-fail ]; then result='<failure/>'; exitcode=1
elif [ "$module" = gamma ] && [ -f .fireweed/gamma-fail ]; then result='<failure/>'; exitcode=1
else result=''; exitcode=0; fi
case "$module" in beta) id=BET-001;; gamma) id=GAM-001;; *) exit 2;; esac
printf '<testsuite><testcase name="%s" classname="fixture">%s</testcase></testsuite>' "$id" "$result" > "$report"
exit "$exitcode"
`)
	write(t, root, ".fireweed/boundaries.sh", `#!/bin/sh
if [ -f .fireweed/boundary-fail ] && [ -f packages/alpha/internal/ready ]; then printf 'alpha: forbidden dependency\n'; exit 1; fi
package lifecycle

import (
	"fmt"
	"path"
	"strings"

	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

func waitingQueue(root string, queue []project.Change) ([]model.StatusEntry, error) {
	entries := make([]model.StatusEntry, 0, len(queue))
	var blocked []project.Change
	for _, ch := range queue {
		entry := model.StatusEntry{
			Slug: ch.Slug, Path: ch.Path, Module: ch.Module,
			Phase: field(ch, "phase"), Blocked: field(ch, "blocked"),
			AuditRounds: intValue(ch.Document.Frontmatter["audit_rounds"], 0),
		}
		if entry.Phase == "" {
			entry.Phase = "queued"
		}
		if entry.Blocked != "" {
			blocked = append(blocked, ch)
		} else if len(blocked) > 0 {
			projection, err := project.Project(root, ch.Slug)
			if err != nil {
				return nil, err
			}
			interfaceEdit := strings.TrimSpace(section(ch.Document, "Interface")) != ""
			for _, earlier := range blocked {
				switch {
				case earlier.Module == ch.Module:
					entry.WaitingRule = "same_module"
				case dependsOn(projection, ch.Module, earlier.Module, true):
					entry.WaitingRule = "depends_on_blocked"
				case interfaceEdit && dependsOn(projection, earlier.Module, ch.Module, true):
					entry.WaitingRule = "interface_dependency"
				}
				if entry.WaitingRule != "" {
					entry.WaitingOn = strings.TrimSuffix(path.Base(earlier.Path), ".md")
					break
				}
			}
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func Eligible(root string) ([]project.Change, error) {
	repo, err := repository(root)
	if err != nil {
		return nil, err
	}
	queue, err := project.Queue(repo.Root)
	if err != nil {
		return nil, err
	}
	return eligible(repo.Root, queue)
}

func eligible(root string, queue []project.Change) ([]project.Change, error) {
	entries, err := waitingQueue(root, queue)
	if err != nil {
		return nil, err
	}
	result := make([]project.Change, 0, len(queue))
	for i, e := range entries {
		if e.Blocked == "" && e.WaitingRule == "" {
			result = append(result, queue[i])
		}
	}
	return result, nil
}

func checkStale(root string, ch project.Change) error {
	stale, err := project.StaleSpecs(root, ch.Slug)
	if err != nil {
		return staleError(err)
	}
	if len(stale) != 0 {
		spec := stale[0]
		return coded("P05", ch.Path, fmt.Sprintf("stale base for %s (recorded %s)", spec.ID, spec.RecordedHash))
	}
	return nil
}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/packages/workflow/internal/lifecycle/costs.go
diff --git a/packages/workflow/internal/lifecycle/costs.go b/packages/workflow/internal/lifecycle/costs.go
index 9b49be791650857642f6e8dbeda3453a953e7425..cc3e7ca72652dc4b352a8cdf657dea2b3452e09d
--- a/packages/workflow/internal/lifecycle/costs.go
+++ b/packages/workflow/internal/lifecycle/costs.go
@@ -13,6 +13,7 @@
 	"strconv"
 	"strings"
 	"time"
+	"unicode"
 
 	"github.com/lurkshark/fireweed/packages/checks"
 	"github.com/lurkshark/fireweed/packages/project"
@@ -157,41 +158,74 @@
 	}
 	var ds []project.Diagnostic
 	seen := map[project.Diagnostic]bool{}
+	selectedSnapshots := make(map[int]bool)
+	for _, index := range selected {
+		selectedSnapshots[index] = true
+	}
 	for i, snapshot := range snapshots {
+		if !selectedSnapshots[i] {
+			continue
+		}
 		for _, diagnostic := range snapshot.Diagnostics {
 			if seen[diagnostic] || diagnostic.Code != "B01" && diagnostic.Code != "B02" && diagnostic.Code != "B03" {
 				continue
 			}
-			for _, cost := range snapshot.Modules {
-				index, ok := selected[cost.Module]
-				if !ok || index != i || !costDiagnosticMatches(diagnostic, cost, known[cost.Module]) {
+			if module := costDiagnosticModule(diagnostic, snapshot.Modules, known); module != "" {
+				if index, ok := selected[module]; !ok || index != i {
 					continue
 				}
-				ds = append(ds, diagnostic)
-				seen[diagnostic] = true
-				break
 			}
+			ds = append(ds, diagnostic)
+			seen[diagnostic] = true
 		}
 	}
 	return values, ds, nil
 }
 
-func costDiagnosticMatches(d project.Diagnostic, cost model.ModuleCost, module project.Module) bool {
-	switch d.Code {
-	case "B01":
-		return true
-	case "B02", "B03":
-		if d.Path != "" {
-			return d.Path == cost.Module || d.Path == module.Dir
+// An empty module means the warning is associated with the snapshot as a
+// whole, either because it is aggregate or because it has no usable module key.
+func costDiagnosticModule(d project.Diagnostic, costs []model.ModuleCost, known map[string]project.Module) string {
+	if d.Code == "B01" {
+		return ""
+	}
+	if d.Path != "" {
+		file := filepath.ToSlash(filepath.Clean(d.Path))
+		var owner string
+		var longest int
+		for _, cost := range costs {
+			if file == cost.Module {
+				return cost.Module
+			}
+			dir := known[cost.Module].Dir
+			if len(dir) > longest && (file == dir || strings.HasPrefix(file, dir+"/")) {
+				owner, longest = cost.Module, len(dir)
+			}
+		}
+		if owner != "" {
+			return owner
 		}
-		if d.Code == "B03" {
-			return strings.Contains(d.Message, cost.Module)
+	}
+	// Saved warnings can also carry the module or spec only in the message.
+	tokens := make(map[string]bool)
+	for _, token := range strings.FieldsFunc(d.Message, func(r rune) bool {
+		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_' && r != '.'
+	}) {
+		tokens[strings.Trim(token, ".")] = true
+	}
+	var owner string
+	for _, cost := range costs {
+		match := d.Code == "B03" && tokens[cost.Module]
+		if d.Code == "B02" {
+			for id := range cost.CasesPerSpec {
+				match = match || tokens[id]
+			}
 		}
-		for id := range cost.CasesPerSpec {
-			if strings.Contains(d.Message, id) {
-				return true
+		if match {
+			if owner != "" && owner != cost.Module {
+				return ""
 			}
+			owner = cost.Module
 		}
 	}
-	return false
+	return owner
 }
diff --git a/packages/workflow/internal/lifecycle/git.go b/packages/workflow/internal/lifecycle/git.go
index bddbd3a3e0cc97817da080b329fff28bc310eb31..3db167c7ec28aed7e6790ae2cd7f13a429d7987c
--- a/packages/workflow/internal/lifecycle/git.go
+++ b/packages/workflow/internal/lifecycle/git.go
@@ -1,10 +1,14 @@
 package lifecycle
 
 import (
+	"bytes"
 	"context"
+	"errors"
 	"fmt"
 	"os"
+	"os/exec"
 	"path/filepath"
+	"strconv"
 	"strings"
 
 	"github.com/lurkshark/fireweed/packages/gitstore"
@@ -30,6 +34,42 @@
 	return refs, nil
 }
 
+func registeredWorktrees(ctx context.Context, root string) ([]string, error) {
+	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	var records []string
+	quoted := false
+	if err == nil {
+		records = splitNUL(out)
+	} else {
+		var exit *exec.ExitError
+		if !errors.As(err, &exit) || exit.ExitCode() != 129 {
+			return nil, err
+		}
+		// Older Git versions support porcelain records but lack -z.
+		out, err = git(ctx, root, nil, nil, "worktree", "list", "--porcelain")
+		if err != nil {
+			return nil, err
+		}
+		records = strings.Split(string(out), "\n")
+		quoted = true
+	}
+	var paths []string
+	for _, record := range records {
+		name, ok := strings.CutPrefix(record, "worktree ")
+		if !ok {
+			continue
+		}
+		if quoted && strings.HasPrefix(name, `"`) {
+			name, err = strconv.Unquote(name)
+			if err != nil {
+				return nil, fmt.Errorf("invalid quoted worktree path: %w", err)
+			}
+		}
+		paths = append(paths, name)
+	}
+	return paths, nil
+}
+
 func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
 	out, err := git(ctx, root, nil, nil, args...)
 	if err != nil {
@@ -80,7 +120,24 @@
 }
 
 func (idx gitIndex) stageWorktree() error {
-	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
+	// Include tracked deletions and non-ignored new files. Giving older Git
+	// a broad add pathspec with an exclusion can still hit ignored directories.
+	out, err := idx.run(nil, "ls-files", "--cached", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return err
+	}
+	var paths bytes.Buffer
+	for _, name := range splitNUL(out) {
+		if name == ".fireweed/runs" || strings.HasPrefix(name, ".fireweed/runs/") {
+			continue
+		}
+		paths.WriteString(name)
+		paths.WriteByte(0)
+	}
+	if paths.Len() == 0 {
+		return nil
+	}
+	_, err = idx.run(paths.Bytes(), "--literal-pathspecs", "add", "-A", "--pathspec-from-file=-", "--pathspec-file-nul")
 	return err
 }
 
diff --git a/packages/workflow/internal/lifecycle/status.go b/packages/workflow/internal/lifecycle/status.go
index 0acd60803368b6ef9fd7454eb9f6196d7b602a28..1fa84461b966a90ede920db86a62a638c7f90946
--- a/packages/workflow/internal/lifecycle/status.go
+++ b/packages/workflow/internal/lifecycle/status.go
@@ -157,17 +157,16 @@
 	for _, ch := range archives {
 		archived[ch.Slug] = true
 	}
-	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	worktrees, err := registeredWorktrees(ctx, root)
 	if err != nil {
 		return result, err
 	}
 	live := map[string]bool{}
 	common := canonicalPath(repo.CommonDir, root)
-	for _, line := range splitNUL(out) {
-		if name, ok := strings.CutPrefix(line, "worktree "); ok {
-			if other, err := repository(name); err == nil && canonicalPath(other.CommonDir, other.Root) == common {
-				live[other.Root] = true
-			}
+	for _, name := range worktrees {
+		if other, err := repository(name); err == nil &&
+			other.Root == canonicalPath(name, root) && canonicalPath(other.CommonDir, other.Root) == common {
+			live[other.Root] = true
 		}
 	}
 	refs, err := listRefs(ctx, root, claimPrefix, attemptPrefix, drillPrefix)

exec
/bin/zsh -lc "sed -n '1,230p' packages/workflow/internal/lifecycle/claim.go; sed -n '160,210p' packages/workflow/internal/lifecycle/common.go; cat packages/checks/api.go; cat packages/project/api.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc "sed -n '1,250p' packages/workflow/evals/block_test.go; sed -n '1,150p' packages/workflow/evals/abandon_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 716ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func block(t *testing.T, root, reason string) workflow.BlockResult {
	t.Helper()
	r, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: reason})
	if err != nil {
		t.Fatalf("Block: %v", err)
	}
	return r
}

func blockFixture(t *testing.T, configExtra, changeFront, added, modified, removed string) string {
	t.Helper()
	root := fixture(t, configExtra, changeFront, added, modified, removed)
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
	return root
}

func readFile(t *testing.T, root, name string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func assertBlockDocumentOnlyAddsReason(t *testing.T, approved, blocked, reason string) {
	t.Helper()
	before, err := contracts.ParseDocument(approved)
	if err != nil {
		t.Fatalf("parse approved change: %v", err)
	}
	after, err := contracts.ParseDocument(blocked)
	if err != nil {
		t.Fatalf("parse blocked change: %v", err)
	}
	if _, exists := before.Frontmatter["blocked"]; exists {
		t.Fatal("fixture change was already blocked")
	}
	want := make(map[string]any, len(before.Frontmatter)+1)
	for key, value := range before.Frontmatter {
		want[key] = value
	}
	want["blocked"] = reason
	if !reflect.DeepEqual(after.Frontmatter, want) {
		t.Errorf("block changed approved frontmatter: got %#v, want %#v", after.Frontmatter, want)
	}
	if strings.TrimSpace(after.Body) != strings.TrimSpace(before.Body) {
		t.Errorf("block changed approved Markdown body: got %q, want %q", after.Body, before.Body)
	}
}

func assertBlockReleased(t *testing.T, root string, r workflow.BlockResult, reason string) {
	t.Helper()
	if r.State.Phase == "" || r.State.ClaimOID != "" || r.State.Head != head(t, root) || r.State.PhaseBase != head(t, root) {
		t.Errorf("returned blocked state = %#v", r.State)
	}
	if got := r.State.Change.Document.Frontmatter["blocked"]; got != reason {
		t.Errorf("returned blocked reason = %#v, want %q", got, reason)
	}
	if got := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); got != "" {
		t.Errorf("claim still exists: %s", got)
	}
	if _, err := workflow.Current(root, "manual"); err == nil {
		t.Error("Current still found a held claim")
	}
	if got := git(t, root, "status", "--porcelain", "-uall"); got != "" {
		t.Errorf("checkout is not clean: %s", got)
	}
	if index, tree := git(t, root, "write-tree"), git(t, root, "rev-parse", "HEAD^{tree}"); index != tree {
		t.Errorf("index tree %s differs from blocked HEAD tree %s", index, tree)
	}
}

func Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState(t *testing.T) {
	root := blockFixture(t, "", "", "### new:extra\nExtra MUST be checked.\n", "", "")
	write(t, root, "packages/alpha/evals/current_test.go", "original eval\n")
	write(t, root, "packages/alpha/evals/retired_test.go", "retired eval\n")
	commit(t, root, "fixture eval files")
	started := start(t, root)
	if started.Phase != "evals" || started.Run == "" || len(started.Change.Added) != 1 || started.Change.Added[0].ID != "ALP-003" {
		t.Fatalf("invalid held evals fixture: %#v", started)
	}
	approved := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(approved, "### ALP-003") || !strings.Contains(approved, started.Run) {
		t.Fatalf("assigned IDs/run absent from fixture: %s", approved)
	}
	base := head(t, root)
	prior := "refs/fireweed/attempts/manual/1"
	git(t, root, "update-ref", prior, base)
	write(t, root, "packages/alpha/evals/current_test.go", "staged modified eval\n")
	git(t, root, "add", "packages/alpha/evals/current_test.go")
	if err := os.Remove(filepath.Join(root, "packages/alpha/evals/retired_test.go")); err != nil {
		t.Fatal(err)
	}
	write(t, root, "packages/alpha/evals/staged_test.go", "staged new eval\n")
	git(t, root, "add", "packages/alpha/evals/staged_test.go")
	write(t, root, "packages/alpha/evals/untracked_test.go", "untracked eval\n")
	if got := git(t, root, "status", "--porcelain", "-uall"); !strings.Contains(got, "M  packages/alpha/evals/current_test.go") || !strings.Contains(got, " D packages/alpha/evals/retired_test.go") || !strings.Contains(got, "A  packages/alpha/evals/staged_test.go") || !strings.Contains(got, "?? packages/alpha/evals/untracked_test.go") {
		t.Fatalf("fixture lacks staged, deleted or untracked work: %s", got)
	}
	if got := git(t, root, "show", ":packages/alpha/evals/current_test.go"); got != "staged modified eval" {
		t.Fatalf("fixture index was not staged: %q", got)
	}
	reason := "verification attempts exhausted"
	r := block(t, root, reason)
	if r.State.Phase != "evals" || r.State.Run != started.Run || len(r.State.Change.Added) != 1 || r.State.Change.Added[0].ID != "ALP-003" {
		t.Errorf("returned state lost held evals phase, run or assigned ID: %#v", r.State)
	}
	wantRef := "refs/fireweed/attempts/manual/2"
	if r.AttemptRef != wantRef || r.AttemptOID == "" || git(t, root, "rev-parse", wantRef) != r.AttemptOID || git(t, root, "cat-file", "-t", r.AttemptOID) != "commit" {
		t.Errorf("attempt ref/OID = %q %q, want %q at commit", r.AttemptRef, r.AttemptOID, wantRef)
	}
	if got := git(t, root, "rev-parse", prior); got != base {
		t.Errorf("prior attempt ref changed: %s, want %s", got, base)
	}
	for path, want := range map[string]string{
		"packages/alpha/evals/current_test.go":   "staged modified eval",
		"packages/alpha/evals/staged_test.go":    "staged new eval",
		"packages/alpha/evals/untracked_test.go": "untracked eval",
	} {
		if got := git(t, root, "show", r.AttemptOID+":"+path); got != want {
			t.Errorf("saved %s = %q, want %q", path, got, want)
		}
	}
	if got := git(t, root, "ls-tree", "--name-only", r.AttemptOID, "packages/alpha/evals/retired_test.go"); got != "" {
		t.Errorf("saved attempt retained deleted file: %s", got)
	}
	if got := git(t, root, "rev-parse", "HEAD^"); got != base {
		t.Errorf("block parent = %s, want phase base %s", got, base)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != "changes/0001-manual.md" {
		t.Errorf("block commit contains step work: %s", got)
	}
	blocked := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(blocked, "### ALP-003") || !strings.Contains(blocked, started.Run) || !strings.Contains(blocked, "phase: evals") || !strings.Contains(blocked, "blocked: "+reason) || !strings.Contains(blocked, "Users need a revised alpha value.") {
		t.Errorf("block commit lost assigned/author state: %s", blocked)
	}
	assertBlockDocumentOnlyAddsReason(t, approved, blocked, reason)
	if got := readFile(t, root, "packages/alpha/evals/current_test.go"); got != "original eval\n" {
		t.Errorf("tracked checkout not restored: %q", got)
	}
	if got := readFile(t, root, "packages/alpha/evals/retired_test.go"); got != "retired eval\n" {
		t.Errorf("deleted checkout not restored: %q", got)
	}
	for _, path := range []string{"packages/alpha/evals/staged_test.go", "packages/alpha/evals/untracked_test.go"} {
		if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
			t.Errorf("step file %s remains after block: %v", path, err)
		}
	}
	assertBlockReleased(t, root, r, reason)
}

func Test_FLOW_066_BlockImplPreservesCommittedEvalsAndAudit(t *testing.T) {
	root := blockFixture(t, "", "", "", "", "")
	start(t, root)
	write(t, root, "packages/alpha/evals/committed_test.go", "committed eval\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("evals Done: %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("audit Done: %#v, %v", audited, err)
	}
	base := head(t, root)
	committedChange := readFile(t, root, "changes/0001-manual.md")
	if !strings.Contains(committedChange, "No findings; checked ALP-001") || readFile(t, root, "packages/alpha/evals/committed_test.go") != "committed eval\n" {
		t.Fatal("fixture lacks committed eval/audit content")
	}
	write(t, root, "packages/alpha/internal/attempt.go", "unfinished implementation\n")
	r := block(t, root, "implementation attempts exhausted")
	if r.State.Phase != "impl" || git(t, root, "rev-parse", "HEAD^") != base {
		t.Errorf("impl block base/phase = %s, %#v", base, r.State)
	}
	if got := git(t, root, "show", r.AttemptOID+":packages/alpha/internal/attempt.go"); got != "unfinished implementation" {
		t.Errorf("impl attempt work = %q", got)
	}
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-only", "-r", "HEAD"); got != "changes/0001-manual.md" {
		t.Errorf("impl block contains work: %s", got)
	}
	if got := readFile(t, root, "changes/0001-manual.md"); !strings.Contains(got, "No findings; checked ALP-001") || !strings.Contains(got, "blocked: implementation attempts exhausted") {
		t.Errorf("impl block lost committed audit state: %s", got)
	}
	assertBlockDocumentOnlyAddsReason(t, committedChange, readFile(t, root, "changes/0001-manual.md"), "implementation attempts exhausted")
	if got := readFile(t, root, "packages/alpha/evals/committed_test.go"); got != "committed eval\n" {
		t.Errorf("impl block lost committed eval: %q", got)
	}
	if _, err := os.Stat(filepath.Join(root, "packages/alpha/internal/attempt.go")); !os.IsNotExist(err) {
		t.Errorf("impl work remains: %v", err)
	}
	assertBlockReleased(t, root, r, "implementation attempts exhausted")
}

func Test_FLOW_063_BlockRejectsForeignClaimApprovedEditsAndChangedHead(t *testing.T) {
	t.Run("no claim", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		before := head(t, root)
		if _, err := workflow.Block(workflow.BlockRequest{Root: root, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("Block accepted change without a claim")
		}
		if head(t, root) != before || git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/attempts/manual") != "" {
			t.Error("unclaimed Block changed HEAD or attempts")
		}
	})
	t.Run("foreign claim", func(t *testing.T) {
		root := blockFixture(t, "", "", "", "", "")
		other := filepath.Join(t.TempDir(), "linked")
		git(t, root, "worktree", "add", "-b", "linked", other, "HEAD")
		started := start(t, root)
		before := head(t, other)
		if _, err := workflow.Block(workflow.BlockRequest{Root: other, Change: "manual", Reason: "exhausted"}); err == nil {
			t.Error("foreign worktree blocked owner's claim")
		}
		if head(t, other) != before {
			t.Error("foreign Block moved HEAD")
		}
		if current, err := workflow.Current(root, "manual"); err != nil || current.ClaimOID != started.ClaimOID {
			t.Errorf("owner's claim changed: %#v, %v", current, err)
		}
	})
	for _, edit := range []string{"author", "phase", "audit"} {
		t.Run(edit, func(t *testing.T) {
			root := blockFixture(t, "", "", "", "", "")
			started := start(t, root)
			path := "changes/0001-manual.md"
			before := readFile(t, root, path)
			switch edit {
			case "author":
				write(t, root, path, strings.Replace(before, "Users need a revised alpha value.", "Unapproved purpose.", 1))
			case "phase":
				write(t, root, path, strings.Replace(before, "phase: evals", "phase: impl", 1))
			case "audit":
				write(t, root, path, strings.Replace(before, "phase: evals", "audit_rounds: 1\nphase: evals", 1))
			}
			if before == readFile(t, root, path) {
package evals_test

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow"
)

func abandonGitIdentity(t *testing.T, root string) {
	t.Helper()
	git(t, root, "config", "user.name", "Fixture")
	git(t, root, "config", "user.email", "fixture@example.invalid")
}

func assertAbandonCommit(t *testing.T, root, parent, oid string) {
	t.Helper()
	if !regexp.MustCompile(`^[0-9a-f]{40}$`).MatchString(oid) || oid != head(t, root) {
		t.Errorf("Abandon returned %q; HEAD is %s", oid, head(t, root))
	}
	if got := git(t, root, "rev-parse", "HEAD^"); got != parent {
		t.Errorf("Abandon parent = %s, want %s (one new commit)", got, parent)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != "" {
		t.Errorf("Abandon left dirty worktree: %s", got)
	}
	for _, path := range []string{pendingCommitChange, archivedCommitChange} {
		if _, err := os.Stat(filepath.Join(root, path)); !os.IsNotExist(err) {
			t.Errorf("change remains at %s: %v", path, err)
		}
	}
	// FLOW-094: the only change-document edit is deletion; no new phase is assigned.
	if diff := git(t, root, "diff", "--unified=0", parent, oid, "--", "changes"); strings.Contains(diff, "+phase:") {
		t.Errorf("Abandon assigned a phase in its commit: %s", diff)
	}
}

func assertAbandonRefusal(t *testing.T, root, selector, code string) {
	t.Helper()
	beforeHead := head(t, root)
	beforeTree := git(t, root, "rev-parse", "HEAD^{tree}")
	beforeStatus := git(t, root, "status", "--porcelain=v1", "--untracked-files=all")
	beforeRefs := releaseRefs(t, root)
	if oid, err := workflow.Abandon(root, selector); err == nil || !strings.HasPrefix(err.Error(), code+":") || oid != "" {
		t.Errorf("Abandon(%q) = %q, %v; want %s refusal without OID", selector, oid, err, code)
	}
	if got := head(t, root); got != beforeHead {
		t.Errorf("rejected Abandon moved HEAD: %s -> %s", beforeHead, got)
	}
	if got := git(t, root, "rev-parse", "HEAD^{tree}"); got != beforeTree {
		t.Errorf("rejected Abandon changed committed tree: %s -> %s", beforeTree, got)
	}
	if got := git(t, root, "status", "--porcelain=v1", "--untracked-files=all"); got != beforeStatus {
		t.Errorf("rejected Abandon changed worktree status: %q -> %q", beforeStatus, got)
	}
	if got := releaseRefs(t, root); got != beforeRefs {
		t.Errorf("rejected Abandon changed refs: %q -> %q", beforeRefs, got)
	}
}

func Test_FLOW_089_FLOW_091_FLOW_094_AbandonUnclaimedQueuedChange(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	parent := head(t, root)
	beforeSpec := readFile(t, root, "packages/alpha/SPEC.md")
	if commits, err := workflow.Commits(root, "manual"); err != nil || len(commits) != 0 {
		t.Fatalf("queued fixture has phase commits: %v, %v", commits, err)
	}
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon unclaimed queued change: %v", err)
	}
	assertAbandonCommit(t, root, parent, oid)
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-status", "-r", oid); got != "D\t"+pendingCommitChange {
		t.Errorf("queued Abandon effects = %q, want change-file deletion only", got)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != beforeSpec {
		t.Errorf("queued Abandon changed applied spec: %q", got)
	}
}

func Test_FLOW_089_FLOW_091_FLOW_094_AbandonCleanOwnedSteppedChange(t *testing.T) {
	root := fixture(t, "", "", "", "", "")
	abandonGitIdentity(t, root)
	initialSpec := readFile(t, root, "packages/alpha/SPEC.md")
	started := start(t, root)
	write(t, root, "packages/alpha/evals/step_test.go", "committed eval effect\n")
	evals, err := done(root)
	if err != nil || evals.State.Phase != "audit" {
		t.Fatalf("fixture evals Done = %#v, %v", evals, err)
	}
	audit(t, root, "Checked IDs: ALP-001\nNo findings\n")
	audited, err := done(root)
	if err != nil || audited.State.Phase != "impl" {
		t.Fatalf("fixture audit Done = %#v, %v", audited, err)
	}
	phaseCommits, err := workflow.Commits(root, "manual")
	if err != nil || len(phaseCommits) != 2 {
		t.Fatalf("fixture phase commits = %v, %v; want evals and audit", phaseCommits, err)
	}
	if got := readFile(t, root, "packages/alpha/evals/step_test.go"); got != "committed eval effect\n" {
		t.Fatalf("fixture lacks committed phase effect: %q", got)
	}
	before := head(t, root)
	current, err := workflow.Current(root, "manual")
	if err != nil || current.ClaimOID == "" || current.Run != started.Run {
		t.Fatalf("fixture lost owned claim: %#v, %v", current, err)
	}
	oid, err := workflow.Abandon(root, "manual")
	if err != nil {
		t.Fatalf("Abandon owned stepped change: %v", err)
	}
	assertAbandonCommit(t, root, before, oid)
	if got := git(t, root, "diff-tree", "--no-commit-id", "--name-status", "-r", oid); got != "D\t"+pendingCommitChange+"\nD\tpackages/alpha/evals/step_test.go" {
		t.Errorf("stepped Abandon effects = %q; want pending change and committed eval removed", got)
	}
	if got := readFile(t, root, "packages/alpha/SPEC.md"); got != initialSpec {
		t.Errorf("stepped Abandon changed applied spec: %q", got)
	}
	if claim := git(t, root, "for-each-ref", "--format=%(objectname)", "refs/fireweed/claims/manual"); claim != "" {
		t.Errorf("owned claim remains: %s", claim)
	}
	if _, err := workflow.Current(root, "manual"); err == nil || !strings.HasPrefix(err.Error(), "U02:") {
		t.Errorf("released claim is still current: %v", err)
	}
}

func Test_FLOW_090_AbandonRejectsDirtyWorkingTree(t *testing.T) {
	for _, kind := range []string{"tracked", "untracked"} {
		t.Run(kind, func(t *testing.T) {
			root := fixture(t, "", "", "", "", "")
			if kind == "tracked" {
				write(t, root, "SYSTEM.md", "dirty tracked content\n")
			} else {
				write(t, root, "stray.txt", "dirty untracked content\n")
			}
			assertAbandonRefusal(t, root, "manual", "P01")
			if kind == "tracked" && readFile(t, root, "SYSTEM.md") != "dirty tracked content\n" {
				t.Error("dirty tracked content changed")
			}
			if kind == "untracked" && readFile(t, root, "stray.txt") != "dirty untracked content\n" {
				t.Error("dirty untracked content changed")
			}
		})
	}

 succeeded in 720ms:
package lifecycle

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project"
	"github.com/lurkshark/fireweed/packages/workflow/model"
)

type claim struct {
	Worktree string           `json:"worktree"`
	Change   string           `json:"change"`
	Path     string           `json:"path"`
	Base     string           `json:"phase_base"`
	Approved string           `json:"approved"`
	Snapshot project.Snapshot `json:"snapshot"`
}

func (c claim) ref() string { return claimPrefix + c.Change }

func (c claim) blob(root string) (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", err
	}
	return gitstore.WriteBlob(root, b)
}

func readClaim(root, oid string) (claim, error) {
	var c claim
	b, err := gitstore.ReadBlob(root, oid)
	if err != nil {
		return c, err
	}
	decoder := json.NewDecoder(strings.NewReader(string(b)))
	decoder.UseNumber()
	err = decoder.Decode(&c)
	c.Snapshot.Config = normalizeMap(c.Snapshot.Config)
	for i := range c.Snapshot.Modules {
		doc := &c.Snapshot.Modules[i].Document
		doc.Frontmatter = normalizeMap(doc.Frontmatter)
	}
	return c, err
}

func normalizeMap(m map[string]any) map[string]any {
	for key, value := range m {
		m[key] = normalizeNumber(value)
	}
	return m
}

func normalizeNumber(value any) any {
	switch v := value.(type) {
	case json.Number:
		if n, err := strconv.Atoi(string(v)); err == nil {
			return n
		}
		if n, err := v.Float64(); err == nil {
			return n
		}
	case map[string]any:
		return normalizeMap(v)
	case []any:
		for i := range v {
			v[i] = normalizeNumber(v[i])
		}
	}
	return value
}

func heldClaim(repo gitstore.Repository, selector string) (claim, string, error) {
	root := repo.Root
	var empty claim
	if selector == "" {
		refs, err := listRefs(context.Background(), root, claimPrefix)
		if err != nil {
			return empty, "", err
		}
		for _, ref := range refs {
			c, err := readClaim(root, ref.oid)
			if err == nil && c.Worktree == root && c.Path != "" && c.ref() == ref.name {
				return c, ref.oid, nil
			}
		}
		return empty, "", coded("P02", root, "worktree holds no change")
	}
	queue, err := changePaths(root, pendingDir)
	if err != nil {
		return empty, "", err
	}
	ch, err := selectChange(queue, selector)
	if err != nil {
		// A deleted held document still belongs to its claim. Verify can then
		// classify the missing document as a change to approved state.
		if c, oid, claimErr := heldClaim(repo, ""); claimErr == nil {
			found := false
			for _, entry := range queue {
				found = found || entry.Path == c.Path
			}
			if !found {
				if _, matchErr := selectChange([]project.Change{identity(c.Path)}, selector); matchErr == nil {
					return c, oid, nil
				}
			}
		}
		return empty, "", err
	}
	oid, exists, err := gitstore.Ref(root, claimPrefix+ch.Slug)
	if err != nil {
		return empty, "", err
	}
	if !exists {
		return empty, "", coded("P02", ch.Path, "change has no claim")
	}
	c, err := readClaim(root, oid)
	if err != nil {
		return empty, "", withCode("P02", ch.Path, err)
	}
	if c.Worktree != root {
		return empty, "", coded("P02", ch.Path, "claim belongs to another worktree")
	}
	if c.Path != ch.Path || c.Change != ch.Slug {
		return empty, "", coded("P02", ch.Path, "claim belongs to another change")
	}
	return c, oid, nil
}

func (c claim) state(repo gitstore.Repository, oid string) (model.State, error) {
	approved, err := readChange(identity(c.Path), []byte(c.Approved))
	if err != nil {
		return model.State{}, err
	}
	c.Snapshot.Root = repo.Root
	return model.State{
		Change: approved, Snapshot: c.Snapshot, Run: field(approved, "run"),
		Phase: field(approved, "phase"), Head: repo.Head, PhaseBase: c.Base, ClaimOID: oid,
	}, nil
}

func Current(root, selector string) (model.State, error) {
	repo, err := repository(root)
	if err != nil {
		return model.State{}, err
	}
	c, oid, err := heldClaim(repo, selector)
	if err != nil {
		return model.State{}, err
	}
	s, err := c.state(repo, oid)
	if err != nil {
		return s, err
	}
	if data, err := os.ReadFile(filepath.Join(repo.Root, c.Path)); err == nil {
		if working, err := readChange(s.Change, data); err == nil {
			s.Change.Document = workerDocument(s.Change.Document, working.Document, s.Phase)
		}
	}
	return s, nil
}

func Release(root, selector string) error {
	repo, err := repository(root)
	if err != nil {
		return err
	}
	c, oid, err := heldClaim(repo, selector)
	if err != nil {
		return err
	}
	return withCode("P02", c.Path, gitstore.CASRef(repo.Root, c.ref(), oid, ""))
}

func Start(root, selector string) (model.State, error) {
	return start(root, selector, nil)
}

func StartFromAttempt(root, selector string, attempt int) (model.State, error) {
	return start(root, selector, &attempt)
}

func start(root, selector string, attempt *int) (s model.State, err error) {
	repo, err := repository(root)
	if err != nil {
		return s, err
	}
	root = repo.Root
	if selector != "" {
		paths, err := changePaths(root, pendingDir)
		if err != nil {
			return s, err
		}
		if _, err := selectChange(paths, selector); err != nil {
			return s, err
		}
	}
	if err = requireClean(root, "start"); err != nil {
		return s, err
	}
	queue, err := project.Queue(root)
	if err != nil {
		return s, staleError(err)
	}
	var chosen project.Change
	if selector != "" {
		chosen, err = selectChange(queue, selector)
		if err != nil {
			return s, err
		}
	}
	ready, err := eligible(root, queue)
	if err != nil {
		return s, staleError(err)
	}
		}
	}
	return project.Module{}, fmt.Errorf("module %q not found", name)
}

func specs(m project.Module) ([]contracts.Item, error) {
	return contracts.ParseItems(section(m.Document, "Specs"))
}

func dependsOn(s project.Snapshot, from, target string, transitive bool) bool {
	seen := make(map[string]bool)
	var visit func(string) bool
	visit = func(name string) bool {
		if seen[name] {
			return false
		}
		seen[name] = true
		m, _ := moduleNamed(s, name)
		for _, dependency := range m.DependsOn {
			if dependency == target || transitive && visit(dependency) {
				return true
			}
		}
		return false
	}
	return visit(from)
}

func relativeTo(dir, file string) (string, bool) {
	if dir == "" || dir == "." {
		return file, true
	}
	rel, ok := strings.CutPrefix(file, strings.TrimSuffix(dir, "/")+"/")
	return rel, ok
}

func matchesAny(patterns []string, name string) bool {
	for _, pattern := range patterns {
		if contracts.MatchGlob(pattern, name) {
			return true
		}
	}
	return false
}

func zone(s project.Snapshot, m project.Module, name string) string {
	if matchesAny(stringList(s.Config["protected"]), name) {
		return "protected"
	}
	relative, inside := relativeTo(m.Dir, name)
	if inside {
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

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index 59654da4bd7e1f88a8de1d63a7da1c39ddcb260e..16aee96526618bd1d98a915047d1318406fba2c0
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -7,5 +7,8 @@
 - FLOW-001 and FLOW-019 · workflow. The claim payload format is left to the implementation. The rebuild stores the canonical worktree, change identity, phase base, approved document and projected snapshot in a JSON blob addressed by the claim ref.
 - FLOW-020 and FLOW-075 · workflow. The spec leaves the complete list of worker-editable body sections open. The retained evals establish Decisions and Questions for evals and Decisions for impl, and Audit for audit. The rebuild also treats Questions as editable during impl.
 - FLOW-059 · workflow. The spec leaves the audit submission grammar open. The rebuild uses the checked-ID line and quoted finding bullets from the retained evals, with an optional untestable marker.
+- FLOW-052 and FLOW-064 · workflow. The spec does not describe how retained run output interacts with staging. Done and Block stage tracked paths and Git-visible untracked paths from a private index, excluding `.fireweed/runs` even when a project's ignore rules omit it.
 - FLOW-074 · workflow. The spec leaves conflict resolution open when a saved attempt overlaps later commits. The rebuild applies the saved delta, tries Git's three-way merge when needed, and returns an error for unresolved conflicts.
+- FLOW-113 and FLOW-123 · workflow. Saved budget diagnostics have no dedicated module field. The rebuild uses a module name or a path within a module, then the module or spec ID in the diagnostic text. Aggregate diagnostics and warnings without a usable module key stay associated with their selected snapshot.
 - FLOW-117 · workflow. The spec leaves file classification open for historical commits that predate loadable project metadata. The rebuild retains their recorded findings and result counts and classifies files for commits with loadable metadata.
+- FLOW-128 · workflow. The spec does not set a minimum Git version. GC falls back to line-delimited porcelain output when Git rejects `worktree list --porcelain -z`.
diff --git a/packages/workflow/internal/lifecycle/costs.go b/packages/workflow/internal/lifecycle/costs.go
index 9b49be791650857642f6e8dbeda3453a953e7425..cc3e7ca72652dc4b352a8cdf657dea2b3452e09d
--- a/packages/workflow/internal/lifecycle/costs.go
+++ b/packages/workflow/internal/lifecycle/costs.go
@@ -13,6 +13,7 @@
 	"strconv"
 	"strings"
 	"time"
+	"unicode"
 
 	"github.com/lurkshark/fireweed/packages/checks"
 	"github.com/lurkshark/fireweed/packages/project"
@@ -157,41 +158,74 @@
 	}
 	var ds []project.Diagnostic
 	seen := map[project.Diagnostic]bool{}
+	selectedSnapshots := make(map[int]bool)
+	for _, index := range selected {
+		selectedSnapshots[index] = true
+	}
 	for i, snapshot := range snapshots {
+		if !selectedSnapshots[i] {
+			continue
+		}
 		for _, diagnostic := range snapshot.Diagnostics {
 			if seen[diagnostic] || diagnostic.Code != "B01" && diagnostic.Code != "B02" && diagnostic.Code != "B03" {
 				continue
 			}
-			for _, cost := range snapshot.Modules {
-				index, ok := selected[cost.Module]
-				if !ok || index != i || !costDiagnosticMatches(diagnostic, cost, known[cost.Module]) {
+			if module := costDiagnosticModule(diagnostic, snapshot.Modules, known); module != "" {
+				if index, ok := selected[module]; !ok || index != i {
 					continue
 				}
-				ds = append(ds, diagnostic)
-				seen[diagnostic] = true
-				break
 			}
+			ds = append(ds, diagnostic)
+			seen[diagnostic] = true
 		}
 	}
 	return values, ds, nil
 }
 
-func costDiagnosticMatches(d project.Diagnostic, cost model.ModuleCost, module project.Module) bool {
-	switch d.Code {
-	case "B01":
-		return true
-	case "B02", "B03":
-		if d.Path != "" {
-			return d.Path == cost.Module || d.Path == module.Dir
+// An empty module means the warning is associated with the snapshot as a
+// whole, either because it is aggregate or because it has no usable module key.
+func costDiagnosticModule(d project.Diagnostic, costs []model.ModuleCost, known map[string]project.Module) string {
+	if d.Code == "B01" {
+		return ""
+	}
+	if d.Path != "" {
+		file := filepath.ToSlash(filepath.Clean(d.Path))
+		var owner string
+		var longest int
+		for _, cost := range costs {
+			if file == cost.Module {
+				return cost.Module
+			}
+			dir := known[cost.Module].Dir
+			if len(dir) > longest && (file == dir || strings.HasPrefix(file, dir+"/")) {
+				owner, longest = cost.Module, len(dir)
+			}
+		}
+		if owner != "" {
+			return owner
 		}
-		if d.Code == "B03" {
-			return strings.Contains(d.Message, cost.Module)
+	}
+	// Saved warnings can also carry the module or spec only in the message.
+	tokens := make(map[string]bool)
+	for _, token := range strings.FieldsFunc(d.Message, func(r rune) bool {
+		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_' && r != '.'
+	}) {
+		tokens[strings.Trim(token, ".")] = true
+	}
+	var owner string
+	for _, cost := range costs {
+		match := d.Code == "B03" && tokens[cost.Module]
+		if d.Code == "B02" {
+			for id := range cost.CasesPerSpec {
+				match = match || tokens[id]
+			}
 		}
-		for id := range cost.CasesPerSpec {
-			if strings.Contains(d.Message, id) {
-				return true
+		if match {
+			if owner != "" && owner != cost.Module {
+				return ""
 			}
+			owner = cost.Module
 		}
 	}
-	return false
+	return owner
 }
diff --git a/packages/workflow/internal/lifecycle/git.go b/packages/workflow/internal/lifecycle/git.go
index bddbd3a3e0cc97817da080b329fff28bc310eb31..3db167c7ec28aed7e6790ae2cd7f13a429d7987c
--- a/packages/workflow/internal/lifecycle/git.go
+++ b/packages/workflow/internal/lifecycle/git.go
@@ -1,10 +1,14 @@
 package lifecycle
 
 import (
+	"bytes"
 	"context"
+	"errors"
 	"fmt"
 	"os"
+	"os/exec"
 	"path/filepath"
+	"strconv"
 	"strings"
 
 	"github.com/lurkshark/fireweed/packages/gitstore"
@@ -30,6 +34,42 @@
 	return refs, nil
 }
 
+func registeredWorktrees(ctx context.Context, root string) ([]string, error) {
+	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	var records []string
+	quoted := false
+	if err == nil {
+		records = splitNUL(out)
+	} else {
+		var exit *exec.ExitError
+		if !errors.As(err, &exit) || exit.ExitCode() != 129 {
+			return nil, err
+		}
+		// Older Git versions support porcelain records but lack -z.
+		out, err = git(ctx, root, nil, nil, "worktree", "list", "--porcelain")
+		if err != nil {
+			return nil, err
+		}
+		records = strings.Split(string(out), "\n")
+		quoted = true
+	}
+	var paths []string
+	for _, record := range records {
+		name, ok := strings.CutPrefix(record, "worktree ")
+		if !ok {
+			continue
+		}
+		if quoted && strings.HasPrefix(name, `"`) {
+			name, err = strconv.Unquote(name)
+			if err != nil {
+				return nil, fmt.Errorf("invalid quoted worktree path: %w", err)
+			}
+		}
+		paths = append(paths, name)
+	}
+	return paths, nil
+}
+
 func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
 	out, err := git(ctx, root, nil, nil, args...)
 	if err != nil {
@@ -80,7 +120,24 @@
 }
 
 func (idx gitIndex) stageWorktree() error {
-	_, err := idx.run(nil, "add", "-A", "--", ".", ":(exclude,top).fireweed/runs")
+	// Include tracked deletions and non-ignored new files. Giving older Git
+	// a broad add pathspec with an exclusion can still hit ignored directories.
+	out, err := idx.run(nil, "ls-files", "--cached", "--others", "--exclude-standard", "-z")
+	if err != nil {
+		return err
+	}
+	var paths bytes.Buffer
+	for _, name := range splitNUL(out) {
+		if name == ".fireweed/runs" || strings.HasPrefix(name, ".fireweed/runs/") {
+			continue
+		}
+		paths.WriteString(name)
+		paths.WriteByte(0)
+	}
+	if paths.Len() == 0 {
+		return nil
+	}
+	_, err = idx.run(paths.Bytes(), "--literal-pathspecs", "add", "-A", "--pathspec-from-file=-", "--pathspec-file-nul")
 	return err
 }
 
diff --git a/packages/workflow/internal/lifecycle/status.go b/packages/workflow/internal/lifecycle/status.go
index 0acd60803368b6ef9fd7454eb9f6196d7b602a28..1fa84461b966a90ede920db86a62a638c7f90946
--- a/packages/workflow/internal/lifecycle/status.go
+++ b/packages/workflow/internal/lifecycle/status.go
@@ -157,17 +157,16 @@
 	for _, ch := range archives {
 		archived[ch.Slug] = true
 	}
-	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
+	worktrees, err := registeredWorktrees(ctx, root)
 	if err != nil {
 		return result, err
 	}
 	live := map[string]bool{}
 	common := canonicalPath(repo.CommonDir, root)
-	for _, line := range splitNUL(out) {
-		if name, ok := strings.CutPrefix(line, "worktree "); ok {
-			if other, err := repository(name); err == nil && canonicalPath(other.CommonDir, other.Root) == common {
-				live[other.Root] = true
-			}
+	for _, name := range worktrees {
+		if other, err := repository(name); err == nil &&
+			other.Root == canonicalPath(name, root) && canonicalPath(other.CommonDir, other.Root) == common {
+			live[other.Root] = true
 		}
 	}
 	refs, err := listRefs(ctx, root, claimPrefix, attemptPrefix, drillPrefix)

exec
/bin/zsh -lc 'gofmt -w packages/workflow/internal/lifecycle/git.go packages/workflow/internal/lifecycle/status.go packages/workflow/internal/lifecycle/costs.go && gofmt -l packages/workflow/internal/lifecycle/git.go packages/workflow/internal/lifecycle/status.go packages/workflow/internal/lifecycle/costs.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 552ms:
exec
/bin/zsh -lc "sed -n '1,160p' packages/workflow/internal/lifecycle/git.go; sed -n '146,202p' packages/workflow/internal/lifecycle/status.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
exec
/bin/zsh -lc "sed -n '133,260p' packages/workflow/internal/lifecycle/costs.go; cat DRILL_NOTES.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728
 succeeded in 559ms:
package lifecycle

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/gitstore"
)

type gitRef struct {
	name string
	oid  string
}

func listRefs(ctx context.Context, root string, prefixes ...string) ([]gitRef, error) {
	args := append([]string{"for-each-ref", "--format=%(refname) %(objectname)"}, prefixes...)
	out, err := git(ctx, root, nil, nil, args...)
	if err != nil {
		return nil, err
	}
	var refs []gitRef
	for _, line := range strings.Split(string(out), "\n") {
		if name, oid, ok := strings.Cut(line, " "); ok {
			refs = append(refs, gitRef{name, oid})
		}
	}
	return refs, nil
}

func registeredWorktrees(ctx context.Context, root string) ([]string, error) {
	out, err := git(ctx, root, nil, nil, "worktree", "list", "--porcelain", "-z")
	var records []string
	quoted := false
	if err == nil {
		records = splitNUL(out)
	} else {
		var exit *exec.ExitError
		if !errors.As(err, &exit) || exit.ExitCode() != 129 {
			return nil, err
		}
		// Older Git versions support porcelain records but lack -z.
		out, err = git(ctx, root, nil, nil, "worktree", "list", "--porcelain")
		if err != nil {
			return nil, err
		}
		records = strings.Split(string(out), "\n")
		quoted = true
	}
	var paths []string
	for _, record := range records {
		name, ok := strings.CutPrefix(record, "worktree ")
		if !ok {
			continue
		}
		if quoted && strings.HasPrefix(name, `"`) {
			name, err = strconv.Unquote(name)
			if err != nil {
				return nil, fmt.Errorf("invalid quoted worktree path: %w", err)
			}
		}
		paths = append(paths, name)
	}
	return paths, nil
}

func diffPaths(ctx context.Context, root string, args ...string) ([]gitstore.Change, error) {
	out, err := git(ctx, root, nil, nil, args...)
	if err != nil {
		return nil, err
	}
	parts := splitNUL(out)
	var changes []gitstore.Change
	for len(parts) > 0 {
		if len(parts) < 2 || parts[0] == "" {
			return nil, fmt.Errorf("invalid git name-status output")
		}
		change := gitstore.Change{Status: parts[0], Path: parts[1]}
		parts = parts[2:]
		if change.Status[0] == 'R' || change.Status[0] == 'C' {
			if len(parts) == 0 {
				return nil, fmt.Errorf("missing destination in git name-status output")
			}
			change.OldPath, change.Path = change.Path, parts[0]
			parts = parts[1:]
		}
		changes = append(changes, change)
	}
	return changes, nil
}

type gitIndex struct {
	root string
	env  []string
	ctx  context.Context
}

func temporaryIndex(ctx context.Context, root, base string) (gitIndex, func(), error) {
	dir, err := os.MkdirTemp("", "fireweed-index-")
	if err != nil {
		return gitIndex{}, nil, err
	}
	clean := func() { os.RemoveAll(dir) }
	idx := gitIndex{root: root, env: []string{"GIT_INDEX_FILE=" + filepath.Join(dir, "index")}, ctx: ctx}
	if _, err := idx.run(nil, "read-tree", base); err != nil {
		clean()
		return gitIndex{}, nil, err
	}
	return idx, clean, nil
}

func (idx gitIndex) run(input []byte, args ...string) ([]byte, error) {
	return git(idx.ctx, idx.root, idx.env, input, args...)
}

func (idx gitIndex) stageWorktree() error {
	// Include tracked deletions and non-ignored new files. Giving older Git
	// a broad add pathspec with an exclusion can still hit ignored directories.
	out, err := idx.run(nil, "ls-files", "--cached", "--others", "--exclude-standard", "-z")
	if err != nil {
		return err
	}
	var paths bytes.Buffer
	for _, name := range splitNUL(out) {
		if name == ".fireweed/runs" || strings.HasPrefix(name, ".fireweed/runs/") {
			continue
		}
		paths.WriteString(name)
		paths.WriteByte(0)
	}
	if paths.Len() == 0 {
		return nil
	}
	_, err = idx.run(paths.Bytes(), "--literal-pathspecs", "add", "-A", "--pathspec-from-file=-", "--pathspec-file-nul")
	return err
}

func (idx gitIndex) stageDocument(name string, data []byte) error {
	oid, err := gitstore.WriteBlob(idx.root, data)
	if err != nil {
		return err
	}
	_, err = idx.run(nil, "update-index", "--add", "--cacheinfo", "100644", oid, name)
	return err
}

func (idx gitIndex) commit(parent, message string) (string, error) {
	out, err := idx.run(nil, "write-tree")
	if err != nil {
		return "", err
	}
	tree := strings.TrimSpace(string(out))
	env := append([]string(nil), idx.env...)
	for _, identity := range []struct{ key, value, author, committer string }{
	}
	now := request.Now
	if now.IsZero() {
		now = time.Now()
	}
	cutoff := now.UTC().Add(-age)
	archives, err := changePaths(root, archiveDir)
	if err != nil {
		return result, err
	}
	archived := map[string]bool{}
	for _, ch := range archives {
		archived[ch.Slug] = true
	}
	worktrees, err := registeredWorktrees(ctx, root)
	if err != nil {
		return result, err
	}
	live := map[string]bool{}
	common := canonicalPath(repo.CommonDir, root)
	for _, name := range worktrees {
		if other, err := repository(name); err == nil &&
			other.Root == canonicalPath(name, root) && canonicalPath(other.CommonDir, other.Root) == common {
			live[other.Root] = true
		}
	}
	refs, err := listRefs(ctx, root, claimPrefix, attemptPrefix, drillPrefix)
	if err != nil {
		return result, err
	}
	for _, ref := range refs {
		if err := ctx.Err(); err != nil {
			return result, err
		}
		prune := false
		switch {
		case strings.HasPrefix(ref.name, claimPrefix):
			c, err := readClaim(root, ref.oid)
			prune = err == nil && c.Worktree != "" && !live[canonicalPath(c.Worktree, root)]
		case strings.HasPrefix(ref.name, attemptPrefix):
			slug, n, ok := strings.Cut(strings.TrimPrefix(ref.name, attemptPrefix), "/")
			attempt, numberErr := strconv.Atoi(n)
			prune = ok && numberErr == nil && attempt > 0 && strconv.Itoa(attempt) == n && archived[slug]
		case strings.HasPrefix(ref.name, drillPrefix):
			module, stamp, ok := strings.Cut(strings.TrimPrefix(ref.name, drillPrefix), "/")
			date, dateErr := time.Parse("20060102T150405.000000000", stamp)
			prune = ok && module != "" && dateErr == nil && date.UTC().Before(cutoff)
		}
		if prune {
			if _, deleteErr := git(ctx, root, nil, nil, "update-ref", "-d", ref.name, ref.oid); deleteErr == nil {
				result.Pruned = append(result.Pruned, ref.name)
			} else if current, found, refErr := gitstore.Ref(root, ref.name); refErr != nil || found && current == ref.oid {
				err = combine(err, deleteErr)
			}
		}
	}
	sort.Strings(result.Pruned)

 succeeded in 559ms:
	return result, nil
}

func latestCosts(root string, modules []project.Module, restrict bool) ([]model.ModuleCost, []project.Diagnostic, error) {
	snapshots, err := Costs(root)
	if err != nil {
		return nil, nil, err
	}
	known := map[string]project.Module{}
	for _, m := range modules {
		known[m.Name] = m
	}
	costs := map[string]model.ModuleCost{}
	selected := map[string]int{}
	for i, snapshot := range snapshots {
		for _, cost := range snapshot.Modules {
			if _, ok := known[cost.Module]; ok || !restrict {
				costs[cost.Module] = cost
				selected[cost.Module] = i
			}
		}
	}
	var values []model.ModuleCost
	for _, module := range sortedKeys(costs) {
		values = append(values, costs[module])
	}
	var ds []project.Diagnostic
	seen := map[project.Diagnostic]bool{}
	selectedSnapshots := make(map[int]bool)
	for _, index := range selected {
		selectedSnapshots[index] = true
	}
	for i, snapshot := range snapshots {
		if !selectedSnapshots[i] {
			continue
		}
		for _, diagnostic := range snapshot.Diagnostics {
			if seen[diagnostic] || diagnostic.Code != "B01" && diagnostic.Code != "B02" && diagnostic.Code != "B03" {
				continue
			}
			if module := costDiagnosticModule(diagnostic, snapshot.Modules, known); module != "" {
				if index, ok := selected[module]; !ok || index != i {
					continue
				}
			}
			ds = append(ds, diagnostic)
			seen[diagnostic] = true
		}
	}
	return values, ds, nil
}

// An empty module means the warning is associated with the snapshot as a
// whole, either because it is aggregate or because it has no usable module key.
func costDiagnosticModule(d project.Diagnostic, costs []model.ModuleCost, known map[string]project.Module) string {
	if d.Code == "B01" {
		return ""
	}
	if d.Path != "" {
		file := filepath.ToSlash(filepath.Clean(d.Path))
		var owner string
		var longest int
		for _, cost := range costs {
			if file == cost.Module {
				return cost.Module
			}
			dir := known[cost.Module].Dir
			if len(dir) > longest && (file == dir || strings.HasPrefix(file, dir+"/")) {
				owner, longest = cost.Module, len(dir)
			}
		}
		if owner != "" {
			return owner
		}
	}
	// Saved warnings can also carry the module or spec only in the message.
	tokens := make(map[string]bool)
	for _, token := range strings.FieldsFunc(d.Message, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '_' && r != '.'
	}) {
		tokens[strings.Trim(token, ".")] = true
	}
	var owner string
	for _, cost := range costs {
		match := d.Code == "B03" && tokens[cost.Module]
		if d.Code == "B02" {
			for id := range cost.CasesPerSpec {
				match = match || tokens[id]
			}
		}
		if match {
			if owner != "" && owner != cost.Module {
				return ""
			}
			owner = cost.Module
		}
	}
	return owner
}
## Contradicted Decisions

None found in the retained workflow Decisions.

## Gaps

- FLOW-001 and FLOW-019 · workflow. The claim payload format is left to the implementation. The rebuild stores the canonical worktree, change identity, phase base, approved document and projected snapshot in a JSON blob addressed by the claim ref.
- FLOW-020 and FLOW-075 · workflow. The spec leaves the complete list of worker-editable body sections open. The retained evals establish Decisions and Questions for evals and Decisions for impl, and Audit for audit. The rebuild also treats Questions as editable during impl.
- FLOW-059 · workflow. The spec leaves the audit submission grammar open. The rebuild uses the checked-ID line and quoted finding bullets from the retained evals, with an optional untestable marker.
- FLOW-052 and FLOW-064 · workflow. The spec does not describe how retained run output interacts with staging. Done and Block stage tracked paths and Git-visible untracked paths from a private index, excluding `.fireweed/runs` even when a project's ignore rules omit it.
- FLOW-074 · workflow. The spec leaves conflict resolution open when a saved attempt overlaps later commits. The rebuild applies the saved delta, tries Git's three-way merge when needed, and returns an error for unresolved conflicts.
- FLOW-113 and FLOW-123 · workflow. Saved budget diagnostics have no dedicated module field. The rebuild uses a module name or a path within a module, then the module or spec ID in the diagnostic text. Aggregate diagnostics and warnings without a usable module key stay associated with their selected snapshot.
- FLOW-117 · workflow. The spec leaves file classification open for historical commits that predate loadable project metadata. The rebuild retains their recorded findings and result counts and classifies files for commits with loadable metadata.
- FLOW-128 · workflow. The spec does not set a minimum Git version. GC falls back to line-delimited porcelain output when Git rejects `worktree list --porcelain -z`.

codex
Updated the disposable workflow code to keep retained run output out of Done and Block commits, support GC on older Git versions, and preserve saved budget warnings that point to files within a module.

Recorded the spec gaps in [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1251601728/DRILL_NOTES.md). Retained inputs were unchanged, and no commit was made.

Formatting checks passed. The workflow, app, and runner evals still need to run in the external verifier.
tokens used
150,444


## Verification failures
workflow eval: E01: failed case "Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses" for FLOW-062
app eval: E01: failed case "Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange" for APP-074
runner eval: E01: failed case "Test_RUN_013_RUN_014_CompleteSpecChange" for RUN-013

## Supervised review

The single authorized invocation used both configured attempts and ended with
22 failing rows across workflow, app and runner. The native report above is
retained byte for byte. The review records the candidate and preparation,
input and trace checks, all final failures, the retained native ref and
cleanup. WF4 remains open.
[Review and evidence](workflow-native-retry-20261108T233724/review.md).
