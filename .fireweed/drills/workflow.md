---
agent: opencode
base_commit: ab2967192247ef0ff830610b0a6eae2f3b2c09ce
date: "2026-11-18T10:37:53.627932Z"
model: opencode/inkling-2-small-free#high
result: fail
---

## Own evals
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
- workflow / Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound / FLOW-035, FLOW-037, FLOW-038, FLOW-039, FLOW-042, FLOW-043, FLOW-046 / pass
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
- workflow / Test_FLOW_059_AuditVerifyChecksSubmissionWithoutAdvancing / FLOW-059 / pass
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
- app / Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning / APP-009, APP-027, APP-028, APP-040 / fail

=== RUN   Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning
    lifecycle_cli_test.go:260: approval warning missing: {"schema":1,"data":"# Context for manual\n\n## Change\n\n---\nmodule: alpha\nphase: evals\nrun: 20261118T115451-08fc9741\nset:\n    interface:\n        - api.go\n        - extra.go\n---\n\n## Motivation\nUsers need an additional value.\n\n## Interface\nAlpha exposes the additional value.\n\n## Added\n### ALP-002\nAlpha MUST return the additional value.\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Held phase\n\nevals\n\n## Projected spec\n\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST retain the existing value.\n### ALP-002\nAlpha MUST return the additional value.\n\n\n## Public interface\n\n### packages/alpha/api.go\n\n```\npackage alpha\n```\n\n## SYSTEM.md\n\n# Fixture\nLifecycle CLI fixture.\n\n\n## Playbook\n\nPlaybook (evals):\n\n# Write evals\nRead the change, projected spec and public interface. Add tests for every assigned added, modified or fixed spec ID in the evals zone.\n\nExercise the module's public behavior that consumers depend on through its public interface. Build ordinary dependency state through public operations. Assert structured diagnostic codes and subjects when the public contract defines them. When the type affects callers, inspect typed values before encoding. Assert exact wording or layout where the public spec requires it.\n\nRun the tests against the current implementation and record genuine fail-first evidence: at least one test for each ID must fail, unless the change marks it already_satisfied. Keep implementation out of this step and do not commit.\n\nKeep temporary stub behavior in disposable files that remain implementation-writable. Implementation must be able to replace each stub without editing frozen interface files. Before completing evals, identify the implementation-writable files for each missing behavior; create a disposable stub file if needed.\n\nFor eval_fix, correct only evals, retain coverage of listed IDs, and require all module evals to pass; fail-first is not required.\n\nDecisions contain only previously unspecified behavioral choices. Write one or two sentences per choice in Decisions. Put verification, debugging and recovery records in Notes.\nDecision: When a lookup matches more than one item, return an ambiguity error.\nNote: Verification passed after rerunning the module tests.\n\n## Other relevant files\n\n(none)\n\n## Declaring a block\n\nIf you cannot continue, declare a block on the change: set `blocked: \u003creason\u003e` in the change's YAML frontmatter (a short blocked reason) and add questions to its `## Questions` section.\nFireweed will save your work and stop the step when you run `done` again with the declared blocked reason and your questions in place.\n\n","diagnostics":[{"code":"P05","message":"base ref refs/remotes/origin/main is not an ancestor of HEAD; the queue may be stale","warning":true},{"code":"P04","message":"change changes/0001-manual.md is absent from both the local base and the shared base; the queue may be unapproved","warning":true}]}
        , <nil>
--- FAIL: Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning (1.55s)

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
- app / Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange / APP-074, APP-076, APP-077 / fail

=== RUN   Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange
    next_cli_test.go:172: next exit=1 stdout="" stderr="P09: changes/0001-next.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed next to resume.\n"
--- FAIL: Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange (2.43s)

- app / Test_APP_075_NextAllowUnmergedAndUsage / APP-075 / fail

=== RUN   Test_APP_075_NextAllowUnmergedAndUsage
    next_cli_test.go:247: allowed next exit=1 stdout="" stderr="P09: changes/0001-next.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed next to resume.\nP04: warning: change changes/0001-next.md is absent from both the local base and the shared base; the queue may be unapproved\n"
--- FAIL: Test_APP_075_NextAllowUnmergedAndUsage (2.38s)

- app / Test_APP_076_NextPlainEmptyQueue / APP-076 / pass
- app / Test_APP_077_NextExitClasses / APP-077 / pass
- app / Test_APP_078_NextWrappedCauseAndAction / APP-078 / pass
- app / Test_APP_079_NextJSONSuccessAndFailure / APP-079 / fail

=== RUN   Test_APP_079_NextJSONSuccessAndFailure
    next_cli_test.go:328: JSON applied exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed next to resume.\",\"path\":\"changes/0001-next.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_079_NextJSONSuccessAndFailure (2.43s)

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
- app / Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation / APP-080, APP-081, APP-083, APP-084 / fail

=== RUN   Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation
    run_cli_test.go:262: run exit=1 stdout="" stderr="P09: changes/0001-alpha.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation (2.36s)

- app / Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON / APP-081, APP-084, APP-086 / fail

=== RUN   Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON
    run_cli_test.go:323: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\",\"path\":\"changes/0001-alpha.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON (2.25s)

- app / Test_APP_083_RunPlainMultiple / APP-083 / fail

=== RUN   Test_APP_083_RunPlainMultiple
    run_cli_test.go:347: run exit=1 stdout="" stderr="P09: changes/0001-alpha.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_083_RunPlainMultiple (2.19s)

- app / Test_APP_086_RunJSONLimited / APP-086 / fail

=== RUN   Test_APP_086_RunJSONLimited
    run_cli_test.go:359: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\",\"path\":\"changes/0001-alpha.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_086_RunJSONLimited (2.21s)

- app / Test_APP_082_RunAllowUnmerged / APP-082 / fail

=== RUN   Test_APP_082_RunAllowUnmerged
    run_cli_test.go:376: run exit=1 stdout="" stderr="P09: changes/0001-alpha.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_082_RunAllowUnmerged (2.37s)

- app / Test_APP_081_APP_084_RunUsageAndConfiguration / APP-081, APP-084 / pass
- app / Test_APP_084_APP_085_APP_086_RunWrappedFailure / APP-084, APP-085, APP-086 / pass
- app / Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes / APP-084, APP-089, APP-090 / fail

=== RUN   Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes
    run_for_cli_test.go:95: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\",\"path\":\"changes/0001-alpha.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes (2.25s)

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
    usage_cli_test.go:168: owned resume exit=1 stderr="" stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P03\",\"message\":\"approved author metadata changed\\napproved author metadata changed. Correct the reported problem and use fireweed next to resume.\",\"warning\":false}]}\n"
    usage_cli_test.go:176: resumed next data={Change: Run: Phase: Empty:false}: <nil>
    usage_cli_test.go:178: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume1645599774/002/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (2.71s)

- app / Test_APP_076_APP_077_NextPlainUsageStop / APP-076, APP-077 / pass
- app / Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume / APP-083, APP-084, APP-086 / fail

=== RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:252: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P03\",\"message\":\"approved author metadata changed\\napproved author metadata changed. Correct the reported problem and use fireweed run to resume.\",\"warning\":false}]}\n" stderr=""
    usage_cli_test.go:254: clean stop data=null; want object: <nil>
    usage_cli_test.go:270: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume1254103917/002/changes/archive/0001-alpha.md: no such file or directory
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (2.26s)

- app / Test_APP_083_APP_084_RunPlainSharedUsageStop / APP-083, APP-084 / fail

=== RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:331: plain run stop exit=1 stdout="" stderr="P03: approved author metadata changed\napproved author metadata changed. Correct the reported problem and use fireweed run to resume.\n"
    usage_cli_test.go:333: plain run lacks alpha: ""
    usage_cli_test.go:333: plain run lacks applied count 1: ""
    usage_cli_test.go:335: plain run lacks resume guidance: ""
    usage_cli_test.go:337: stop reason "" lacks a labeled completed wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled next wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled budget of 5ms
    usage_cli_test.go:339: alpha not applied: stat /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_RunPlainSharedUsageStop1206435393/002/changes/archive/0001-alpha.md: no such file or directory
    usage_cli_test.go:341: held beta state: change has no claim
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (2.29s)

- app / Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop / APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/next / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/run / APP-077, APP-079, APP-084, APP-086 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/failure / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/timeout / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/cancellation / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/append / RUN-063 / pass
- runner / Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput / RUN-064 / fail

=== RUN   Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput
    agent_logs_test.go:109: Next fixture = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115640-d4528e9c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:42.089591Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115640-d4528e9c", Phase:"audit", Head:"56056bee559713e1c7a77bc702b66252ee9cbe0a", PhaseBase:"56056bee559713e1c7a77bc702b66252ee9cbe0a", ClaimOID:"adb70e179f65896e4533aa6b403ea2ad2854b510", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/evals-2.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput (2.43s)

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
- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals / RUN-021, RUN-070 / fail

=== RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals (6.85s)

- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return / RUN-021, RUN-070 / fail

=== RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return
    audit_return_attempts_test.go:187: phase base did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115655-82145afd\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:56.978572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:57.358136Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T11:56:58.276811Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:58.653465Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2"
    audit_return_attempts_test.go:189: blocked change did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nblocked: evals attempt allowance exhausted after 2 attempts\nmodule: alpha\nphase: evals\nrun: 20261118T115655-82145afd\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:56.978572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:57.358136Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T11:56:58.276811Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:58.653465Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2\n"
    audit_return_attempts_test.go:194: returned state lost the latest audit finding: "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2\n"
    audit_return_attempts_test.go:197: saved attempt did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115655-82145afd\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:56.978572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:57.358136Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T11:56:58.276811Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:58.653465Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2"
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return (3.79s)

- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry / RUN-021, RUN-070 / fail

=== RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry
    audit_return_attempts_test.go:187: phase base did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 1\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115659-fb252365\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:01.363451Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:57:01.740402Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1"
    audit_return_attempts_test.go:189: blocked change did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 1\nbases: {}\nblocked: evals attempt allowance exhausted after 2 attempts\nmodule: alpha\nphase: evals\nrun: 20261118T115659-fb252365\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:01.363451Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:57:01.740402Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1\n"
    audit_return_attempts_test.go:194: returned state lost the latest audit finding: "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1\n"
    audit_return_attempts_test.go:197: saved attempt did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 1\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115659-fb252365\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:01.363451Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:57:01.740402Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1"
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry (3.06s)

- runner / Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance / RUN-021 / fail

=== RUN   Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance
    audit_return_attempts_test.go:235: returned Questions = "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2\n", want finding "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2"
--- FAIL: Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance (3.35s)

- runner / Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock / RUN-070 / fail

=== RUN   Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock
    audit_return_attempts_test.go:282: Run invocations = [alpha/evals alpha/audit gamma/evals gamma/audit], want [alpha/evals alpha/audit gamma/evals gamma/audit gamma/impl]
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock (3.33s)

- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit / RUN-013, RUN-045, RUN-046, RUN-071 / fail

=== RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit (3.83s)

- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override / RUN-013, RUN-045, RUN-046, RUN-071 / fail

=== RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115709-7dc275b7"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:10.507326Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115709-7dc275b7", Phase:"audit", Head:"f461e4d00bfffb623267c3c6c97da0f3679338b5", PhaseBase:"f461e4d00bfffb623267c3c6c97da0f3679338b5", ClaimOID:"cc3b15c6935e04ea7ad9e4f33bfbcda5649ad0ed", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001/.fireweed/runs/20261118T115709-7dc275b7/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001/.fireweed/runs/20261118T115709-7dc275b7/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001/.fireweed/runs/20261118T115709-7dc275b7/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override (1.88s)

- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback / RUN-013, RUN-045, RUN-046, RUN-071 / fail

=== RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115711-b660c6a2"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:12.471105Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115711-b660c6a2", Phase:"audit", Head:"7323354131ddd4c24cb0df4b263d800d6347495c", PhaseBase:"7323354131ddd4c24cb0df4b263d800d6347495c", ClaimOID:"759753cccc294e9e93bdf892cf1269a8196df413", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001/.fireweed/runs/20261118T115711-b660c6a2/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001/.fireweed/runs/20261118T115711-b660c6a2/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001/.fireweed/runs/20261118T115711-b660c6a2/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback (1.95s)

- runner / Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration / RUN-013 / pass
- runner / Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry / RUN-048 / fail

=== RUN   Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry
    audit_workspace_test.go:210: audit verification retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115714-c675223e"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:15.936364Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115714-c675223e", Phase:"audit", Head:"3e5a1a06e52c9036c094a9f4c47293b969dae9a2", PhaseBase:"3e5a1a06e52c9036c094a9f4c47293b969dae9a2", ClaimOID:"c89bbd02ad7d758ef1cd27018d56f76fc4162bbd", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001/.fireweed/runs/20261118T115714-c675223e/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001/.fireweed/runs/20261118T115714-c675223e/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001/.fireweed/runs/20261118T115714-c675223e/audit-2.log"}}}}, approved author metadata changed
        approved author metadata changed
--- FAIL: Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry (1.92s)

- runner / Test_RUN_046_AuditCommitDoesNotEnterHeldGit / RUN-046 / fail

=== RUN   Test_RUN_046_AuditCommitDoesNotEnterHeldGit
    audit_workspace_test.go:232: audit commit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115716-3d1ea161"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:18.025281Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115716-3d1ea161", Phase:"audit", Head:"1eeb3ace855fe55c20f204b4fea2bbb3200ec7c5", PhaseBase:"1eeb3ace855fe55c20f204b4fea2bbb3200ec7c5", ClaimOID:"a26a8105f0a62a326c136f7e31ab1e44263bd18b", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001/.fireweed/runs/20261118T115716-3d1ea161/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001/.fireweed/runs/20261118T115716-3d1ea161/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001/.fireweed/runs/20261118T115716-3d1ea161/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_046_AuditCommitDoesNotEnterHeldGit (2.01s)

- runner / Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported / RUN-047, RUN-048 / pass
- runner / Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces / RUN-045, RUN-049 / fail

=== RUN   Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces
    audit_workspace_test.go:349: two-change audit lifecycle: applied=0 error=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces (1.58s)

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
- runner / Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle / RUN-027, RUN-028 / fail

=== RUN   Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle
    commit_recovery_test.go:38: Next after fixture agent commits: invalid audit workspace submission: multiple Audit sections; state=model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115741-e0350e6e"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:42.46816Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115741-e0350e6e", Phase:"audit", Head:"34061ccd438c34f76c9be4f2cd0bb3ec4c64dfaa", PhaseBase:"34061ccd438c34f76c9be4f2cd0bb3ec4c64dfaa", ClaimOID:"4c95caf4b13541874037db1ee85298696f5de4f4", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}; steps=[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001/.fireweed/runs/20261118T115741-e0350e6e/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001/.fireweed/runs/20261118T115741-e0350e6e/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001/.fireweed/runs/20261118T115741-e0350e6e/audit-2.log"}}}
--- FAIL: Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle (2.02s)

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
- runner / Test_RUN_011_WaitingHeadSelectsIndependent / RUN-011 / fail

=== RUN   Test_RUN_011_WaitingHeadSelectsIndependent
    next_test.go:439: Next past waiting beta = model.NextResult{State:model.State{Change:model.Change{Number:3, Slug:"gamma", Path:"changes/0003-gamma.md", Module:"gamma", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"gamma", "phase":"audit", "run":"20261118T115836-1e976187"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### GAM-002\ngamma MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:37.488028Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"GAM-002", Tag:"", Body:"gamma MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002", "exhaust-alpha", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# alpha\n\n## Specs\n\n### ALP-001\nalpha MUST return a value.\n\n### ALP-002\nalpha MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# beta\n\n## Specs\n\n### BET-001\nbeta MUST return a value.\n\n### new:value\nbeta MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"GAM", "name":"gamma"}, Body:"# gamma\n\n## Specs\n\n### GAM-001\ngamma MUST return a value.\n\n### GAM-002\ngamma MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115836-1e976187", Phase:"audit", Head:"a398e1a11a635acbab9fb2b0f81131155fecbf56", PhaseBase:"a398e1a11a635acbab9fb2b0f81131155fecbf56", ClaimOID:"32f05f98693a31fb1c62ee05a0b9467d54a69dc7", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"gamma-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001/.fireweed/runs/20261118T115836-1e976187/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"gamma-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001/.fireweed/runs/20261118T115836-1e976187/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"gamma-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001/.fireweed/runs/20261118T115836-1e976187/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_011_WaitingHeadSelectsIndependent (4.26s)

- runner / Test_RUN_010_EntirelyWaitingQueue / RUN-010 / pass
- runner / Test_RUN_013_RUN_014_CompleteSpecChange / RUN-013, RUN-014 / fail

=== RUN   Test_RUN_013_RUN_014_CompleteSpecChange
    next_test.go:478: Next: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115840-c66fa482"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:41.464273Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115840-c66fa482", Phase:"audit", Head:"8445c2f6551ab73e0222b8070acf66ba7b090aa0", PhaseBase:"8445c2f6551ab73e0222b8070acf66ba7b090aa0", ClaimOID:"6b661f6406349c4882aad6b6c26afa0733a524f3", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001/.fireweed/runs/20261118T115840-c66fa482/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001/.fireweed/runs/20261118T115840-c66fa482/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001/.fireweed/runs/20261118T115840-c66fa482/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_013_RUN_014_CompleteSpecChange (1.78s)

- runner / Test_RUN_015_InvocationFailureRetainsHeldPhase / RUN-015 / pass
- runner / Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks / RUN-023, RUN-024 / pass
- runner / Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase / RUN-021, RUN-022 / fail

=== RUN   Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase
    next_test.go:706: Next after retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115844-26dd119d"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:46.382795Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115844-26dd119d", Phase:"audit", Head:"11e75f65a0bff9584ca3b6ca5930940eba5be056", PhaseBase:"11e75f65a0bff9584ca3b6ca5930940eba5be056", ClaimOID:"f595c105c28c0624bd9b493393d29715c7af5e79", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/evals-2.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase (2.38s)

- runner / Test_RUN_021_ConfiguredThreeAttempts / RUN-021 / fail

=== RUN   Test_RUN_021_ConfiguredThreeAttempts
    next_test.go:756: Next with three attempts = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115847-696d7484"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:48.196608Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002", "retry-third", "{prompt}"}}, "default":"fixture-worker"}, "attempts":3, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115847-696d7484", Phase:"audit", Head:"b103d2920519dcbb0ad68d7006d53c4e8e047466", PhaseBase:"b103d2920519dcbb0ad68d7006d53c4e8e047466", ClaimOID:"6251742ef5374c870cac6764fcd034fab1b9a25f", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001/.fireweed/runs/20261118T115847-696d7484/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001/.fireweed/runs/20261118T115847-696d7484/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001/.fireweed/runs/20261118T115847-696d7484/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_021_ConfiguredThreeAttempts (1.81s)

- runner / Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks / RUN-023, RUN-024 / fail

=== RUN   Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks
    next_test.go:782: Next error code = "P09" (present: true), want E01 (error: invalid audit workspace submission: multiple Audit sections)
    next_test.go:785: exhaustion cause = invalid audit workspace submission: multiple Audit sections, want testcase "ALP-002_fixture_value"
    next_test.go:788: exhaustion steps = 3, want 4
--- FAIL: Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks (1.75s)

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
- runner / Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures / RUN-022, RUN-043 / fail

=== RUN   Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures
    retry_failure_details_test.go:90: Next with JUnit-only dependent failure = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115853-98a799c7"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:53.705851Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002", "retry-impl", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"aardvark", IDPrefix:"AAR", Dir:"packages/aardvark", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"AAR", "name":"aardvark"}, Body:"# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115853-98a799c7", Phase:"audit", Head:"2ffe183e390a2d1fef7ddb0ea8bd85fefbdfa557", PhaseBase:"2ffe183e390a2d1fef7ddb0ea8bd85fefbdfa557", ClaimOID:"1a42cc8ec6c0e3382a1a0418b54579f1ddbeb0c3", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001/.fireweed/runs/20261118T115853-98a799c7/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001/.fireweed/runs/20261118T115853-98a799c7/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001/.fireweed/runs/20261118T115853-98a799c7/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures (1.32s)

- runner / Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions / RUN-044 / fail

=== RUN   Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions
    retry_failure_details_test.go:133: Next with large JUnit failures = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115854-5bd1031c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:55.022646Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002", "retry-impl", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"aardvark", IDPrefix:"AAR", Dir:"packages/aardvark", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"AAR", "name":"aardvark"}, Body:"# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115854-5bd1031c", Phase:"audit", Head:"5cbbf3e822229ef8e09d7dd343cdf9a43a23a364", PhaseBase:"5cbbf3e822229ef8e09d7dd343cdf9a43a23a364", ClaimOID:"3ec38545acb8f8ea994d0d331053278d98024344", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001/.fireweed/runs/20261118T115854-5bd1031c/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001/.fireweed/runs/20261118T115854-5bd1031c/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001/.fireweed/runs/20261118T115854-5bd1031c/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions (1.31s)

- runner / Test_RUN_016_UnlimitedQueueOrder / RUN-016 / fail

=== RUN   Test_RUN_016_UnlimitedQueueOrder
    run_test.go:313: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_016_UnlimitedQueueOrder (2.05s)

- runner / Test_RUN_016_NegativeMaxIsUnlimited / RUN-016 / fail

=== RUN   Test_RUN_016_NegativeMaxIsUnlimited
    run_test.go:329: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_016_NegativeMaxIsUnlimited (2.17s)

- runner / Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding / RUN-017, RUN-018, RUN-019 / fail

=== RUN   Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding
    run_test.go:341: first Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding (2.09s)

- runner / Test_RUN_018_InitialEmptyAndExactLimit / RUN-018 / fail

=== RUN   Test_RUN_018_InitialEmptyAndExactLimit
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit (2.24s)

- runner / Test_RUN_018_InitialEmptyAndExactLimit/initial_empty / RUN-018 / pass
- runner / Test_RUN_018_InitialEmptyAndExactLimit/exact_limit / RUN-018 / fail

=== RUN   Test_RUN_018_InitialEmptyAndExactLimit/exact_limit
    run_test.go:385: exact Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit/exact_limit (2.05s)

- runner / Test_RUN_020_HeldFailureRetainsEarlierResults / RUN-020 / fail

=== RUN   Test_RUN_020_HeldFailureRetainsEarlierResults
    run_test.go:399: held failure cause = invalid audit workspace submission: multiple Audit sections
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_020_HeldFailureRetainsEarlierResults (2.05s)

- runner / Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent / RUN-016, RUN-025, RUN-026 / fail

=== RUN   Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent
    run_test.go:427: blocked run retained claim "refs/fireweed/claims/gamma"
    run_test.go:427: blocked run left dirty checkout "M changes/0003-gamma.md"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent (3.76s)

- runner / Test_RUN_025_RUN_026_PositiveMaxRetainsBlock / RUN-025, RUN-026 / fail

=== RUN   Test_RUN_025_RUN_026_PositiveMaxRetainsBlock
    run_test.go:448: blocked run retained claim "refs/fireweed/claims/gamma"
    run_test.go:448: blocked run left dirty checkout "M changes/0003-gamma.md"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_025_RUN_026_PositiveMaxRetainsBlock (3.80s)

- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt / RUN-014 / fail

=== RUN   Test_RUN_014_NextRunsFullVerificationOncePerAttempt
--- FAIL: Test_RUN_014_NextRunsFullVerificationOncePerAttempt (2.60s)

- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass / RUN-014 / fail

=== RUN   Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass
    single_verification_test.go:28: Next returned phase "audit", empty=false, claim="bcb4961b4b67576f2c6c7d99546c6f4c1852b042", error=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass (1.30s)

- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry / RUN-014 / fail

=== RUN   Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry
    single_verification_test.go:28: Next returned phase "audit", empty=false, claim="6f2e214a587b1a603141cf93b0764933a69d32bd", error=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry (1.30s)

- runner / Test_RUN_072_RetainsStartWarning / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning
--- FAIL: Test_RUN_072_RetainsStartWarning (9.40s)

- runner / Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery
    start_warnings_test.go:81: Next: invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery (1.95s)

- runner / Test_RUN_072_RetainsStartWarning/verification_exhaustion / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning/verification_exhaustion
    start_warnings_test.go:94: expected implementation verification exhaustion, got 3 steps and error invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_072_RetainsStartWarning/verification_exhaustion (1.71s)

- runner / Test_RUN_072_RetainsStartWarning/invocation_failure / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/usage_stop / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/Run_applied_result / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning/Run_applied_result
    start_warnings_test.go:147: expected Run to apply one queued change, got 0 applied, empty=false, limited=false, err=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_072_RetainsStartWarning/Run_applied_result (1.77s)

- runner / Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied / RUN-036, RUN-037 / fail

=== RUN   Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied
    time_limit_test.go:102: clean time stop = invalid audit workspace submission: multiple Audit sections, want classified nonnil error naming 30s
    time_limit_test.go:110: Applied count = 0, want one retained alpha
    time_limit_test.go:113: invocations through held beta = []string{"alpha|evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied4101557207/001|/fixture/run-bin", "alpha|audit|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-audit-4292775755|/fixture/run-bin", "alpha|audit|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-audit-1937299072|/fixture/run-bin"}
--- FAIL: Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied (1.70s)

- runner / Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded / RUN-036 / fail

=== RUN   Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded
    time_limit_test.go:124: zero For run = model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded (1.65s)

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
- runner / Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep / RUN-030, RUN-031, RUN-033 / fail

=== RUN   Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep
    usage_test.go:206: owned resume = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T120017-906e0c1f"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T12:00:18.964416Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":1, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T120017-906e0c1f", Phase:"audit", Head:"3fc09cc39246a298c8cfaef7b75350491f667fa9", PhaseBase:"3fc09cc39246a298c8cfaef7b75350491f667fa9", ClaimOID:"12c84ebe057f4fe0f71587616593d4102459f9af", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001/.fireweed/runs/20261118T120017-906e0c1f/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001/.fireweed/runs/20261118T120017-906e0c1f/audit-1.log"}}}}, approved author metadata changed
        approved author metadata changed
    usage_test.go:208: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep (2.71s)

- runner / Test_RUN_033_OwnedResumeUsesApprovedUsageBudget / RUN-033 / pass
- runner / Test_RUN_030_UsageRetryPreservesVerificationFeedback / RUN-030 / fail

=== RUN   Test_RUN_030_UsageRetryPreservesVerificationFeedback
    usage_test.go:288: verification and usage retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T120020-d2bc8ac6"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T12:00:22.50461Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":2, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002"}}, "limit_budget":"20ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T120020-d2bc8ac6", Phase:"audit", Head:"afe4dd55bdc7e13ec7bc3a2ba096e59b460f902f", PhaseBase:"afe4dd55bdc7e13ec7bc3a2ba096e59b460f902f", ClaimOID:"366e096fbf691826e976b1d901dddf983b7dd666", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"first verification", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"limited evals", Stderr:"RATE LIMIT", ExitCode:9, UsageLimit:true, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/evals-2.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/evals-2.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_030_UsageRetryPreservesVerificationFeedback (2.47s)

- runner / Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases / RUN-030, RUN-031 / fail

=== RUN   Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases
    usage_test.go:307: cross-phase budget = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T120023-5c757d46"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T12:00:24.354395Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":1, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T120023-5c757d46", Phase:"audit", Head:"0ac4a26f37d75de550a850f2f94c9ce9c5adce43", PhaseBase:"0ac4a26f37d75de550a850f2f94c9ce9c5adce43", ClaimOID:"2d3859751eda2a2506b39253de39513db8f9315a", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"limited evals", Stderr:"RATE LIMIT", ExitCode:9, UsageLimit:true, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001/.fireweed/runs/20261118T120023-5c757d46/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001/.fireweed/runs/20261118T120023-5c757d46/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001/.fireweed/runs/20261118T120023-5c757d46/audit-1.log"}}}}, approved author metadata changed
        approved author metadata changed
--- FAIL: Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases (1.86s)

- runner / Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied / RUN-032 / fail

=== RUN   Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied
    usage_test.go:329: Run stop = invalid audit workspace submission: multiple Audit sections
    usage_test.go:331: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied (2.13s)

- runner / Test_RUN_034_UsageStopClassifierWrapping / RUN-034 / pass
- runner / Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure / RUN-035 / pass
- runner / Test_RUN_031_CancelledUsageWaitRetainsHeldWork / RUN-031 / pass

## Contradicted Decisions
None. The four retained Decisions remain satisfied by this build:

- [0012-manual-lifecycle] Provenance lines begin `- <phase>` and carry
  `playbook=<full SHA-256>` plus `tests=N failing=N passing=N` from the reports
  verified by that Done call (encoded in `provenanceLine`/`provenanceCounts`).
- [0119-verification-costs] Cost snapshots are retained at
  `.fireweed/runs/<run>/<phase>/costs.json` beside module reports (module
  reports named `costs` case-insensitively are written to
  `<phase>/reports/<module>.json`); repeated verification replaces the snapshot;
  line counting is newline count plus one for a nonempty unterminated final
  line (`lineCount`).
- [0120-status-report] PhaseCommits are keyed by archived change filename
  without `.md`; budget diagnostics from selected snapshots are returned once
  each (`costDiagnostics` dedups); the first queue-order blocker is reported
  with same-module/dependent/interface-dependency precedence; `Status`
  applied order is implementation Provenance dates, newest first with archive
  path as tie-breaker (`appliedOrder`).
- [0123-pr-body-report] Review context for a selected change includes the phase
  commits reachable at the endpoint regardless of the requested range
  (`selectedPhaseHistory`); selection is identities touched in the range
  present at the endpoint; 256 KB means 262144 bytes with strictly larger text
  files warned; the Cost/Metric section carries one row per measured module
  (module, seconds, cases per spec `ID: count`, eval lines, disposable lines).

## Gaps
- The "enforced" qualifier in "every enforced current spec" (FLOW-022) has no
  definition in this module; enforcement is supplied by the checks engine
  (review/quarantine tags exempt). A spec clarification (or explicit reference)
  would remove the ambiguity. Nearest spec: FLOW-022. Module for a fixing
  change: workflow.
- Static validation in impl Done (FLOW-049) requires the applied project to
  pass, but the pending change being applied has stale recorded bases at that
  moment. The implementation validates with the pending change set aside; the
  spec does not state that ordering or the rollback semantics. Nearest spec:
  FLOW-049/050. Module for a fixing change: workflow.
- Audit-phase verification evaluates no modules (no hooks, zero provenance
  counts). That is required by the evals (audit fixtures fail with failing
  ALP-001 testcases) but FLOW-003/FLOW-059 do not state it. Nearest spec:
  FLOW-003/059. Module for a fixing change: workflow.
- FLOW-039 does not say which `audit_rounds` value a blocked commit records.
  This build persists `audit.max_rounds` on round-limit blocking and increments
  on untestable blocking (FLOW-040 path); the evals only assert the two-round
  default bound and the configured bound. Nearest spec: FLOW-039/040. Module
  for a fixing change: workflow.
- FLOW-110's rule for the unapproved archive count does not define how a
  same-identity pending document affects the comparison. This build treats an
  applied archived change as present at the base when the same filename/number
  exists at either the archive or the pending path, as the eval requires.
  Nearest spec: FLOW-110. Module for a fixing change: workflow.
- FLOW-119 requires review-tagged Added/Modified specs of selected changes and
  forbids unrelated pre-existing review specs; "whole change" kind names
  (`spec`, `fix`, `eval_fix`, `regenerate`, `maintenance`) are used in the PR
  body but are not defined by the spec. Nearest spec: FLOW-114/119. Module for
  a fixing change: workflow.
- The archive-absence rule for FLOW-129 matches attempt refs by slug of
  archived change files; the spec does not describe retention of attempts for
  archived-but-renamed identities, though no eval exercises it. Nearest spec:
  FLOW-129. Module for a fixing change: workflow.
- FLOW-036's "enforced ID" set for No-findings submissions is unspecified in
  the spec; this build excludes review/quarantine-tagged targets (consistent
  with FLOW-023). Nearest spec: FLOW-036. Module for a fixing change: workflow.
- Retained evaluation artifacts that a hook writes during verification were
  observed to break zone checking (errors like `protected path
  .fireweed/runs/...` produced "protected" diagnostics for git-unignored
  artifacts). This build excludes `.fireweed/runs/**` from zone diagnostics for
  exactly that reason (FLOW-135/FLOW-052). The exclusion is a reconstruction
  decision; the spec explicitly requires only that commits exclude the
  artifacts. Nearest spec: FLOW-002/135. Module for a fixing change: checks or
  workflow.
- `gitstore.Run` returns the process error for nonzero exits (exit 1 yields a
  non-nil error alongside `ExitCode: 1`), so code that probes exit status must
  read `Result.ExitCode` first (as `isAncestor` now does). This behavior is
  not documented in the gitstore interface and every nonzero-exit probe needs
  the same care. Nearest spec: no spec. Module for a fixing change: gitstore.
- The `git worktree list --porcelain` format prints paths unescaped, so a
  worktree path containing a newline corrupts naive line parsing; this build
  reassembles multiline paths (GC and worktree registration). No spec or eval
  requires this robustness, but FLOW-128 relies on worktree liveness matching.
  Nearest spec: FLOW-128. Module for a fixing change: workflow (or gitstore).
- FLOW-051 counts skipped testcases under `failing`; the spec does not define
  how `skip` results are categorized in `tests/failing/passing`. Nearest spec:
  FLOW-051. Module for a fixing change: workflow.

## Agent observations

### Invocation 1
Log: .fireweed/runs/20261118T103753-f61e9868c7fb5a043c45d27f/drill-1.log

This is a large reconstruction task. Let me first verify a couple of serialization details that the held-claim design depends on, then build the `internal/lifecycle` package incrementally.
The fixtures are large; the key design decisions are made (tagged-JSON held snapshot, claims as JSON with the snapshot embedded). Now I'll build the `internal/lifecycle` package. Starting with the shared core file.
The ops.go file needs a clean rewrite with the claim creation factored out. Let me rewrite it and add the missing helpers.
Let me write the costs machinery first since Verify depends on it.
Everything builds. Now let's run the eval suite and start fixing failures.
Let me temporarily add debug output to see the exact hash mismatch.
Let me check what `project.Check` actually reports at each step of impl Done with a temporary debug test.
The applied validation must not see the pending change. I'll temporarily set the pending change aside during static validation, then restore it on failure.
Several fixes to apply. First, remove my debug instrumentation and the debug test, then fix the static validation ordering — impl Done must archive+validate-then-commit with rollback on failure.
All 218 evals pass. Now let me check for DRILL_NOTES requirements.
All evals pass. Let me do a final holistic review of the worktree state and then write DRILL_NOTES.md.
Let me do a final integrity check: confirm the hook still passes from a clean state, and that only disposable paths were touched.
Done. The workflow module is reconstructed and the eval hook reports **PASS** (218 testcases, verified with `-count=1`).

## What was rebuilt

`packages/workflow/internal/**` (the disposable zone) — a new `internal/lifecycle` package implementing the full controller:

- **Claims**: stored as JSON blobs at `refs/fireweed/claims/<slug>` carrying worktree, run, phase, phase base and the *approved snapshot* (config, module metadata, zone map, projected spec documents) encoded with a tagged-scalar codec so `int`/`float64` survive round-trips exactly. `Approved`/`Current` render the approved state with zero live parsing, satisfying FLOW-019/139/140.
- **Lifecycle**: `Start` (eligibility, ID reservation across current/archived removals, disposable regenerate deletion projected from set-edits, claim CAS), `StartFromAttempt` (pre-mutation attempt-ref validation, worker-section restore preserving author/runtime/Provenance, later-independent-edit preservation), `Verify` (phase gate → stale-base P05 → zone checks → approval-metadata P03 → module evals with coverage/fail-first/pass rules → dependent selection `verify_dependents` → boundary verification → costs), `Done` (verification-failure marking per FLOW-136, worker-declared blocks per FLOW-137, phase transitions with provenance, audit normalization/max_rounds blocking, impl application with static validation rollback-before-commit, archive move, claim release), `Block` (attempt refs via temp index excluding `.fireweed/runs`, restore-to-phase-base), plus `Eligible`, `Commits`, `Abandon`, `Release`, `Context`, `Costs`, `Status`, `PRBody`, `GC`, with the required diagnostic-code contract (P01/P02/P03/P04/P05/P06/P07/P09/P10, Z01/X01/X02, C08/C10/C11, B01–B04, U02).

## Verification

- `go vet` clean; `go test -count=1 ./packages/workflow/evals/` passes all 218 spec evals; `go run .fireweed/hooks.go eval packages/workflow .fireweed/drill-eval.xml` → `ok`.
- Only disposable paths added: `packages/workflow/internal/lifecycle/*.go` (+ `DRILL_NOTES.md` and the hook's untracked `.fireweed/drill-eval.xml`). `api.go`, `model/**`, `SPEC.md` and `evals/**` are untouched.

## Notes

- `DRILL_NOTES.md` records the two required sections: **Contradicted Decisions** (none — the four Decisions are honored) and **Gaps** (spec holes found while reconstructing, each with nearest spec ID and the target module for a fixing change).
- Two discovered environment quirks worth knowing: `gitstore.Run` returns non-nil `err` for *any* nonzero exit (exit-status probes must read `Result.ExitCode` first), and `git worktree list --porcelain` prints paths unescaped (newline-safe parsing used in GC).


## Failures
app: failed case "Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning" for APP-009
Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning: === RUN   Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning
    lifecycle_cli_test.go:260: approval warning missing: {"schema":1,"data":"# Context for manual\n\n## Change\n\n---\nmodule: alpha\nphase: evals\nrun: 20261118T115451-08fc9741\nset:\n    interface:\n        - api.go\n        - extra.go\n---\n\n## Motivation\nUsers need an additional value.\n\n## Interface\nAlpha exposes the additional value.\n\n## Added\n### ALP-002\nAlpha MUST return the additional value.\n\n## Modified\n\n## Removed\n\n## Decisions\n\n## Held phase\n\nevals\n\n## Projected spec\n\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST retain the existing value.\n### ALP-002\nAlpha MUST return the additional value.\n\n\n## Public interface\n\n### packages/alpha/api.go\n\n```\npackage alpha\n```\n\n## SYSTEM.md\n\n# Fixture\nLifecycle CLI fixture.\n\n\n## Playbook\n\nPlaybook (evals):\n\n# Write evals\nRead the change, projected spec and public interface. Add tests for every assigned added, modified or fixed spec ID in the evals zone.\n\nExercise the module's public behavior that consumers depend on through its public interface. Build ordinary dependency state through public operations. Assert structured diagnostic codes and subjects when the public contract defines them. When the type affects callers, inspect typed values before encoding. Assert exact wording or layout where the public spec requires it.\n\nRun the tests against the current implementation and record genuine fail-first evidence: at least one test for each ID must fail, unless the change marks it already_satisfied. Keep implementation out of this step and do not commit.\n\nKeep temporary stub behavior in disposable files that remain implementation-writable. Implementation must be able to replace each stub without editing frozen interface files. Before completing evals, identify the implementation-writable files for each missing behavior; create a disposable stub file if needed.\n\nFor eval_fix, correct only evals, retain coverage of listed IDs, and require all module evals to pass; fail-first is not required.\n\nDecisions contain only previously unspecified behavioral choices. Write one or two sentences per choice in Decisions. Put verification, debugging and recovery records in Notes.\nDecision: When a lookup matches more than one item, return an ambiguity error.\nNote: Verification passed after rerunning the module tests.\n\n## Other relevant files\n\n(none)\n\n## Declaring a block\n\nIf you cannot continue, declare a block on the change: set `blocked: \u003creason\u003e` in the change's YAML frontmatter (a short blocked reason) and add questions to its `## Questions` section.\nFireweed will save your work and stop the step when you run `done` again with the declared blocked reason and your questions in place.\n\n","diagnostics":[{"code":"P05","message":"base ref refs/remotes/origin/main is not an ancestor of HEAD; the queue may be stale","warning":true},{"code":"P04","message":"change changes/0001-manual.md is absent from both the local base and the shared base; the queue may be unapproved","warning":true}]}
        , <nil>
--- FAIL: Test_APP_009_APP_027_APP_028_APP_040_StartPreservesApprovalWarning (1.55s)
Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange: === RUN   Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange
    next_cli_test.go:172: next exit=1 stdout="" stderr="P09: changes/0001-next.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed next to resume.\n"
--- FAIL: Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange (2.43s)
Test_APP_075_NextAllowUnmergedAndUsage: === RUN   Test_APP_075_NextAllowUnmergedAndUsage
    next_cli_test.go:247: allowed next exit=1 stdout="" stderr="P09: changes/0001-next.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed next to resume.\nP04: warning: change changes/0001-next.md is absent from both the local base and the shared base; the queue may be unapproved\n"
--- FAIL: Test_APP_075_NextAllowUnmergedAndUsage (2.38s)
Test_APP_079_NextJSONSuccessAndFailure: === RUN   Test_APP_079_NextJSONSuccessAndFailure
    next_cli_test.go:328: JSON applied exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed next to resume.\",\"path\":\"changes/0001-next.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_079_NextJSONSuccessAndFailure (2.43s)
Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation: === RUN   Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation
    run_cli_test.go:262: run exit=1 stdout="" stderr="P09: changes/0001-alpha.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation (2.36s)
Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON: === RUN   Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON
    run_cli_test.go:323: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\",\"path\":\"changes/0001-alpha.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON (2.25s)
Test_APP_083_RunPlainMultiple: === RUN   Test_APP_083_RunPlainMultiple
    run_cli_test.go:347: run exit=1 stdout="" stderr="P09: changes/0001-alpha.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_083_RunPlainMultiple (2.19s)
Test_APP_086_RunJSONLimited: === RUN   Test_APP_086_RunJSONLimited
    run_cli_test.go:359: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\",\"path\":\"changes/0001-alpha.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_086_RunJSONLimited (2.21s)
Test_APP_082_RunAllowUnmerged: === RUN   Test_APP_082_RunAllowUnmerged
    run_cli_test.go:376: run exit=1 stdout="" stderr="P09: changes/0001-alpha.md: invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_082_RunAllowUnmerged (2.37s)
Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes: === RUN   Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes
    run_for_cli_test.go:95: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P09\",\"message\":\"invalid audit workspace submission: multiple Audit sections. Correct the reported problem and use fireweed run to resume.\",\"path\":\"changes/0001-alpha.md\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes (2.25s)
Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume: === RUN   Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume
    usage_cli_test.go:168: owned resume exit=1 stderr="" stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P03\",\"message\":\"approved author metadata changed\\napproved author metadata changed. Correct the reported problem and use fireweed next to resume.\",\"warning\":false}]}\n"
    usage_cli_test.go:176: resumed next data={Change: Run: Phase: Empty:false}: <nil>
    usage_cli_test.go:178: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume1645599774/002/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (2.71s)
Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume: === RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:252: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P03\",\"message\":\"approved author metadata changed\\napproved author metadata changed. Correct the reported problem and use fireweed run to resume.\",\"warning\":false}]}\n" stderr=""
    usage_cli_test.go:254: clean stop data=null; want object: <nil>
    usage_cli_test.go:270: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume1254103917/002/changes/archive/0001-alpha.md: no such file or directory
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (2.26s)
Test_APP_083_APP_084_RunPlainSharedUsageStop: === RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:331: plain run stop exit=1 stdout="" stderr="P03: approved author metadata changed\napproved author metadata changed. Correct the reported problem and use fireweed run to resume.\n"
    usage_cli_test.go:333: plain run lacks alpha: ""
    usage_cli_test.go:333: plain run lacks applied count 1: ""
    usage_cli_test.go:335: plain run lacks resume guidance: ""
    usage_cli_test.go:337: stop reason "" lacks a labeled completed wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled next wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled budget of 5ms
    usage_cli_test.go:339: alpha not applied: stat /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_RunPlainSharedUsageStop1206435393/002/changes/archive/0001-alpha.md: no such file or directory
    usage_cli_test.go:341: held beta state: change has no claim
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (2.29s)
?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	236.670s

exit status 1

runner: failed case "Test_RUN_011_WaitingHeadSelectsIndependent" for RUN-011
Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput: === RUN   Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput
    agent_logs_test.go:109: Next fixture = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115640-d4528e9c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:42.089591Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115640-d4528e9c", Phase:"audit", Head:"56056bee559713e1c7a77bc702b66252ee9cbe0a", PhaseBase:"56056bee559713e1c7a77bc702b66252ee9cbe0a", ClaimOID:"adb70e179f65896e4533aa6b403ea2ad2854b510", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/evals-2.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit4135844620/001/.fireweed/runs/20261118T115640-d4528e9c/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput (2.43s)
Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals: === RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals (6.85s)
Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return: === RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return
    audit_return_attempts_test.go:187: phase base did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115655-82145afd\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:56.978572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:57.358136Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T11:56:58.276811Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:58.653465Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2"
    audit_return_attempts_test.go:189: blocked change did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nblocked: evals attempt allowance exhausted after 2 attempts\nmodule: alpha\nphase: evals\nrun: 20261118T115655-82145afd\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:56.978572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:57.358136Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T11:56:58.276811Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:58.653465Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2\n"
    audit_return_attempts_test.go:194: returned state lost the latest audit finding: "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2\n"
    audit_return_attempts_test.go:197: saved attempt did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115655-82145afd\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:56:56.978572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:57.358136Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T11:56:58.276811Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:56:58.653465Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2"
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return (3.79s)
Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry: === RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry
    audit_return_attempts_test.go:187: phase base did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 1\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115659-fb252365\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:01.363451Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:57:01.740402Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1"
    audit_return_attempts_test.go:189: blocked change did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 1\nbases: {}\nblocked: evals attempt allowance exhausted after 2 attempts\nmodule: alpha\nphase: evals\nrun: 20261118T115659-fb252365\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:01.363451Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:57:01.740402Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1\n"
    audit_return_attempts_test.go:194: returned state lost the latest audit finding: "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1\n"
    audit_return_attempts_test.go:197: saved attempt did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 1\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T115659-fb252365\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:01.363451Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T11:57:01.740402Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n## Audit\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 1"
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry (3.06s)
Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance: === RUN   Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance
    audit_return_attempts_test.go:235: returned Questions = "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — — missing an assertion for subsequent calls in audit 2\n", want finding "- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2"
--- FAIL: Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance (3.35s)
Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock: === RUN   Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock
    audit_return_attempts_test.go:282: Run invocations = [alpha/evals alpha/audit gamma/evals gamma/audit], want [alpha/evals alpha/audit gamma/evals gamma/audit gamma/impl]
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock (3.33s)
Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit: === RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit (3.83s)
Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override: === RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115709-7dc275b7"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:10.507326Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115709-7dc275b7", Phase:"audit", Head:"f461e4d00bfffb623267c3c6c97da0f3679338b5", PhaseBase:"f461e4d00bfffb623267c3c6c97da0f3679338b5", ClaimOID:"cc3b15c6935e04ea7ad9e4f33bfbcda5649ad0ed", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001/.fireweed/runs/20261118T115709-7dc275b7/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001/.fireweed/runs/20261118T115709-7dc275b7/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput1227372604/001/.fireweed/runs/20261118T115709-7dc275b7/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override (1.88s)
Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback: === RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115711-b660c6a2"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:12.471105Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115711-b660c6a2", Phase:"audit", Head:"7323354131ddd4c24cb0df4b263d800d6347495c", PhaseBase:"7323354131ddd4c24cb0df4b263d800d6347495c", ClaimOID:"759753cccc294e9e93bdf892cf1269a8196df413", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001/.fireweed/runs/20261118T115711-b660c6a2/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001/.fireweed/runs/20261118T115711-b660c6a2/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2162148760/001/.fireweed/runs/20261118T115711-b660c6a2/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback (1.95s)
Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry: === RUN   Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry
    audit_workspace_test.go:210: audit verification retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115714-c675223e"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:15.936364Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115714-c675223e", Phase:"audit", Head:"3e5a1a06e52c9036c094a9f4c47293b969dae9a2", PhaseBase:"3e5a1a06e52c9036c094a9f4c47293b969dae9a2", ClaimOID:"c89bbd02ad7d758ef1cd27018d56f76fc4162bbd", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001/.fireweed/runs/20261118T115714-c675223e/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001/.fireweed/runs/20261118T115714-c675223e/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1410505929/001/.fireweed/runs/20261118T115714-c675223e/audit-2.log"}}}}, approved author metadata changed
        approved author metadata changed
--- FAIL: Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry (1.92s)
Test_RUN_046_AuditCommitDoesNotEnterHeldGit: === RUN   Test_RUN_046_AuditCommitDoesNotEnterHeldGit
    audit_workspace_test.go:232: audit commit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115716-3d1ea161"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:18.025281Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115716-3d1ea161", Phase:"audit", Head:"1eeb3ace855fe55c20f204b4fea2bbb3200ec7c5", PhaseBase:"1eeb3ace855fe55c20f204b4fea2bbb3200ec7c5", ClaimOID:"a26a8105f0a62a326c136f7e31ab1e44263bd18b", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001/.fireweed/runs/20261118T115716-3d1ea161/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001/.fireweed/runs/20261118T115716-3d1ea161/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit3283934207/001/.fireweed/runs/20261118T115716-3d1ea161/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_046_AuditCommitDoesNotEnterHeldGit (2.01s)
Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces: === RUN   Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces
    audit_workspace_test.go:349: two-change audit lifecycle: applied=0 error=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_045_RUN_049_SeparateAuditsUseFreshWorkspaces (1.58s)
Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle: === RUN   Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle
    commit_recovery_test.go:38: Next after fixture agent commits: invalid audit workspace submission: multiple Audit sections; state=model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115741-e0350e6e"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:57:42.46816Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115741-e0350e6e", Phase:"audit", Head:"34061ccd438c34f76c9be4f2cd0bb3ec4c64dfaa", PhaseBase:"34061ccd438c34f76c9be4f2cd0bb3ec4c64dfaa", ClaimOID:"4c95caf4b13541874037db1ee85298696f5de4f4", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}; steps=[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001/.fireweed/runs/20261118T115741-e0350e6e/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001/.fireweed/runs/20261118T115741-e0350e6e/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle2461722236/001/.fireweed/runs/20261118T115741-e0350e6e/audit-2.log"}}}
--- FAIL: Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle (2.02s)
Test_RUN_011_WaitingHeadSelectsIndependent: === RUN   Test_RUN_011_WaitingHeadSelectsIndependent
    next_test.go:439: Next past waiting beta = model.NextResult{State:model.State{Change:model.Change{Number:3, Slug:"gamma", Path:"changes/0003-gamma.md", Module:"gamma", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"gamma", "phase":"audit", "run":"20261118T115836-1e976187"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### GAM-002\ngamma MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:37.488028Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"GAM-002", Tag:"", Body:"gamma MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/002", "exhaust-alpha", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# alpha\n\n## Specs\n\n### ALP-001\nalpha MUST return a value.\n\n### ALP-002\nalpha MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# beta\n\n## Specs\n\n### BET-001\nbeta MUST return a value.\n\n### new:value\nbeta MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"GAM", "name":"gamma"}, Body:"# gamma\n\n## Specs\n\n### GAM-001\ngamma MUST return a value.\n\n### GAM-002\ngamma MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115836-1e976187", Phase:"audit", Head:"a398e1a11a635acbab9fb2b0f81131155fecbf56", PhaseBase:"a398e1a11a635acbab9fb2b0f81131155fecbf56", ClaimOID:"32f05f98693a31fb1c62ee05a0b9467d54a69dc7", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"gamma-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001/.fireweed/runs/20261118T115836-1e976187/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"gamma-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001/.fireweed/runs/20261118T115836-1e976187/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"gamma-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1539770518/001/.fireweed/runs/20261118T115836-1e976187/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_011_WaitingHeadSelectsIndependent (4.26s)
Test_RUN_013_RUN_014_CompleteSpecChange: === RUN   Test_RUN_013_RUN_014_CompleteSpecChange
    next_test.go:478: Next: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115840-c66fa482"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:41.464273Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115840-c66fa482", Phase:"audit", Head:"8445c2f6551ab73e0222b8070acf66ba7b090aa0", PhaseBase:"8445c2f6551ab73e0222b8070acf66ba7b090aa0", ClaimOID:"6b661f6406349c4882aad6b6c26afa0733a524f3", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001/.fireweed/runs/20261118T115840-c66fa482/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001/.fireweed/runs/20261118T115840-c66fa482/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange3492556644/001/.fireweed/runs/20261118T115840-c66fa482/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_013_RUN_014_CompleteSpecChange (1.78s)
Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase: === RUN   Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase
    next_test.go:706: Next after retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115844-26dd119d"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:46.382795Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115844-26dd119d", Phase:"audit", Head:"11e75f65a0bff9584ca3b6ca5930940eba5be056", PhaseBase:"11e75f65a0bff9584ca3b6ca5930940eba5be056", ClaimOID:"f595c105c28c0624bd9b493393d29715c7af5e79", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/evals-2.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase1839902448/001/.fireweed/runs/20261118T115844-26dd119d/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase (2.38s)
Test_RUN_021_ConfiguredThreeAttempts: === RUN   Test_RUN_021_ConfiguredThreeAttempts
    next_test.go:756: Next with three attempts = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115847-696d7484"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:48.196608Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002", "retry-third", "{prompt}"}}, "default":"fixture-worker"}, "attempts":3, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115847-696d7484", Phase:"audit", Head:"b103d2920519dcbb0ad68d7006d53c4e8e047466", PhaseBase:"b103d2920519dcbb0ad68d7006d53c4e8e047466", ClaimOID:"6251742ef5374c870cac6764fcd034fab1b9a25f", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001/.fireweed/runs/20261118T115847-696d7484/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001/.fireweed/runs/20261118T115847-696d7484/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2955317746/001/.fireweed/runs/20261118T115847-696d7484/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_021_ConfiguredThreeAttempts (1.81s)
Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks: === RUN   Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks
    next_test.go:782: Next error code = "P09" (present: true), want E01 (error: invalid audit workspace submission: multiple Audit sections)
    next_test.go:785: exhaustion cause = invalid audit workspace submission: multiple Audit sections, want testcase "ALP-002_fixture_value"
    next_test.go:788: exhaustion steps = 3, want 4
--- FAIL: Test_RUN_023_RUN_024_ImplExhaustionSavesWorkAndBlocks (1.75s)
Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures: === RUN   Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures
    retry_failure_details_test.go:90: Next with JUnit-only dependent failure = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115853-98a799c7"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:53.705851Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002", "retry-impl", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"aardvark", IDPrefix:"AAR", Dir:"packages/aardvark", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"AAR", "name":"aardvark"}, Body:"# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115853-98a799c7", Phase:"audit", Head:"2ffe183e390a2d1fef7ddb0ea8bd85fefbdfa557", PhaseBase:"2ffe183e390a2d1fef7ddb0ea8bd85fefbdfa557", ClaimOID:"1a42cc8ec6c0e3382a1a0418b54579f1ddbeb0c3", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001/.fireweed/runs/20261118T115853-98a799c7/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001/.fireweed/runs/20261118T115853-98a799c7/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFail1961838356/001/.fireweed/runs/20261118T115853-98a799c7/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_022_RUN_043_RetryRetainsContextAndIncludesJUnitOnlyFailures (1.32s)
Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions: === RUN   Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions
    retry_failure_details_test.go:133: Next with large JUnit failures = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T115854-5bd1031c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T11:58:55.022646Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002", "retry-impl", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"aardvark", IDPrefix:"AAR", Dir:"packages/aardvark", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"AAR", "name":"aardvark"}, Body:"# Aardvark\n\n## Specs\n\n### AAR-001\nAardvark MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T115854-5bd1031c", Phase:"audit", Head:"5cbbf3e822229ef8e09d7dd343cdf9a43a23a364", PhaseBase:"5cbbf3e822229ef8e09d7dd343cdf9a43a23a364", ClaimOID:"3ec38545acb8f8ea994d0d331053278d98024344", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001/.fireweed/runs/20261118T115854-5bd1031c/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001/.fireweed/runs/20261118T115854-5bd1031c/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions1659426337/001/.fireweed/runs/20261118T115854-5bd1031c/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_044_RetryBoundsUnicodeFailureExcerptsAndMarksOmissions (1.31s)
Test_RUN_016_UnlimitedQueueOrder: === RUN   Test_RUN_016_UnlimitedQueueOrder
    run_test.go:313: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_016_UnlimitedQueueOrder (2.05s)
Test_RUN_016_NegativeMaxIsUnlimited: === RUN   Test_RUN_016_NegativeMaxIsUnlimited
    run_test.go:329: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_016_NegativeMaxIsUnlimited (2.17s)
Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding: === RUN   Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding
    run_test.go:341: first Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding (2.09s)
Test_RUN_018_InitialEmptyAndExactLimit: === RUN   Test_RUN_018_InitialEmptyAndExactLimit
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit (2.24s)
Test_RUN_018_InitialEmptyAndExactLimit/exact_limit: === RUN   Test_RUN_018_InitialEmptyAndExactLimit/exact_limit
    run_test.go:385: exact Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit/exact_limit (2.05s)
Test_RUN_020_HeldFailureRetainsEarlierResults: === RUN   Test_RUN_020_HeldFailureRetainsEarlierResults
    run_test.go:399: held failure cause = invalid audit workspace submission: multiple Audit sections
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_020_HeldFailureRetainsEarlierResults (2.05s)
Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent: === RUN   Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent
    run_test.go:427: blocked run retained claim "refs/fireweed/claims/gamma"
    run_test.go:427: blocked run left dirty checkout "M changes/0003-gamma.md"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent (3.76s)
Test_RUN_025_RUN_026_PositiveMaxRetainsBlock: === RUN   Test_RUN_025_RUN_026_PositiveMaxRetainsBlock
    run_test.go:448: blocked run retained claim "refs/fireweed/claims/gamma"
    run_test.go:448: blocked run left dirty checkout "M changes/0003-gamma.md"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_025_RUN_026_PositiveMaxRetainsBlock (3.80s)
Test_RUN_014_NextRunsFullVerificationOncePerAttempt: === RUN   Test_RUN_014_NextRunsFullVerificationOncePerAttempt
--- FAIL: Test_RUN_014_NextRunsFullVerificationOncePerAttempt (2.60s)
Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass: === RUN   Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass
    single_verification_test.go:28: Next returned phase "audit", empty=false, claim="bcb4961b4b67576f2c6c7d99546c6f4c1852b042", error=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass (1.30s)
Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry: === RUN   Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry
    single_verification_test.go:28: Next returned phase "audit", empty=false, claim="6f2e214a587b1a603141cf93b0764933a69d32bd", error=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry (1.30s)
Test_RUN_072_RetainsStartWarning: === RUN   Test_RUN_072_RetainsStartWarning
--- FAIL: Test_RUN_072_RetainsStartWarning (9.40s)
Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery: === RUN   Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery
    start_warnings_test.go:81: Next: invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery (1.95s)
Test_RUN_072_RetainsStartWarning/verification_exhaustion: === RUN   Test_RUN_072_RetainsStartWarning/verification_exhaustion
    start_warnings_test.go:94: expected implementation verification exhaustion, got 3 steps and error invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_072_RetainsStartWarning/verification_exhaustion (1.71s)
Test_RUN_072_RetainsStartWarning/Run_applied_result: === RUN   Test_RUN_072_RetainsStartWarning/Run_applied_result
    start_warnings_test.go:147: expected Run to apply one queued change, got 0 applied, empty=false, limited=false, err=invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_072_RetainsStartWarning/Run_applied_result (1.77s)
Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied: === RUN   Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied
    time_limit_test.go:102: clean time stop = invalid audit workspace submission: multiple Audit sections, want classified nonnil error naming 30s
    time_limit_test.go:110: Applied count = 0, want one retained alpha
    time_limit_test.go:113: invocations through held beta = []string{"alpha|evals|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied4101557207/001|/fixture/run-bin", "alpha|audit|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-audit-4292775755|/fixture/run-bin", "alpha|audit|/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-audit-1937299072|/fixture/run-bin"}
--- FAIL: Test_RUN_036_RUN_037_OneBudgetAcrossChangesRetainsApplied (1.70s)
Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded: === RUN   Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded
    time_limit_test.go:124: zero For run = model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_036_ZeroForKeepsOrdinaryRunUnbounded (1.65s)
Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep: === RUN   Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep
    usage_test.go:206: owned resume = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T120017-906e0c1f"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T12:00:18.964416Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":1, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T120017-906e0c1f", Phase:"audit", Head:"3fc09cc39246a298c8cfaef7b75350491f667fa9", PhaseBase:"3fc09cc39246a298c8cfaef7b75350491f667fa9", ClaimOID:"12c84ebe057f4fe0f71587616593d4102459f9af", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001/.fireweed/runs/20261118T120017-906e0c1f/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001/.fireweed/runs/20261118T120017-906e0c1f/audit-1.log"}}}}, approved author metadata changed
        approved author metadata changed
    usage_test.go:208: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep2801722550/001/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep (2.71s)
Test_RUN_030_UsageRetryPreservesVerificationFeedback: === RUN   Test_RUN_030_UsageRetryPreservesVerificationFeedback
    usage_test.go:288: verification and usage retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T120020-d2bc8ac6"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T12:00:22.50461Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":2, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/002"}}, "limit_budget":"20ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T120020-d2bc8ac6", Phase:"audit", Head:"afe4dd55bdc7e13ec7bc3a2ba096e59b460f902f", PhaseBase:"afe4dd55bdc7e13ec7bc3a2ba096e59b460f902f", ClaimOID:"366e096fbf691826e976b1d901dddf983b7dd666", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"first verification", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"limited evals", Stderr:"RATE LIMIT", ExitCode:9, UsageLimit:true, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/evals-2.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/evals-2.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2225491363/001/.fireweed/runs/20261118T120020-d2bc8ac6/audit-2.log"}}}}, invalid audit workspace submission: multiple Audit sections
--- FAIL: Test_RUN_030_UsageRetryPreservesVerificationFeedback (2.47s)
Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases: === RUN   Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases
    usage_test.go:307: cross-phase budget = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "module":"alpha", "phase":"audit", "run":"20261118T120023-5c757d46"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T12:00:24.354395Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":1, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T120023-5c757d46", Phase:"audit", Head:"0ac4a26f37d75de550a850f2f94c9ce9c5adce43", PhaseBase:"0ac4a26f37d75de550a850f2f94c9ce9c5adce43", ClaimOID:"2d3859751eda2a2506b39253de39513db8f9315a", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"limited evals", Stderr:"RATE LIMIT", ExitCode:9, UsageLimit:true, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001/.fireweed/runs/20261118T120023-5c757d46/evals-1.log"}}, model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001/.fireweed/runs/20261118T120023-5c757d46/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases1377148109/001/.fireweed/runs/20261118T120023-5c757d46/audit-1.log"}}}}, approved author metadata changed
        approved author metadata changed
--- FAIL: Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases (1.86s)
Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied: === RUN   Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied
    usage_test.go:329: Run stop = invalid audit workspace submission: multiple Audit sections
    usage_test.go:331: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied (2.13s)
?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/policy	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/workspace	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/runner/evals	232.660s

exit status 1

