---
agent: opencode
base_commit: f282faee412a591912b8f34dba1c66f947a44b5e
date: "2026-11-18T07:14:30.75524Z"
model: opencode/inkling-2-small-free#high
result: fail
---

## Own evals
- gitstore / Test_GIT_001_InspectWorktreesAndDetachedHead / GIT-001 / pass
- gitstore / Test_GIT_002_BaseRefPreferenceDefaultsAndMissing / GIT-002 / pass
- gitstore / Test_GIT_003_GIT_006_GIT_007_ChangesReadAtAndSelectedCommit / GIT-003, GIT-006, GIT-007 / pass
- gitstore / Test_GIT_003_UnstagedTrackedRename / GIT-003 / pass
- gitstore / Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS / GIT-004, GIT-008, GIT-009, GIT-010 / pass
- gitstore / Test_GIT_004_GIT_008_GIT_009_GIT_010_BlobsAndRefCAS/StaleUpdatePreservesNewerRef / GIT-004, GIT-008, GIT-009, GIT-010 / pass
- gitstore / Test_GIT_010_DeleteWithAbsentExpectedRef / GIT-010 / pass
- gitstore / Test_GIT_010_DeleteWithAbsentExpectedRef/Empty / GIT-010 / pass
- gitstore / Test_GIT_010_DeleteWithAbsentExpectedRef/ZeroSHA1 / GIT-010 / pass
- gitstore / Test_GIT_005_RunRequiresExecutableForGitSubcommand / GIT-005 / pass
- gitstore / Test_GIT_011_RunLaunchFailure / GIT-011 / pass
- gitstore / Test_GIT_005_GIT_011_RunCaptureExitTimeoutAndGroupCancel / GIT-005, GIT-011 / pass
- gitstore / Test_GIT_005_GIT_011_RunCaptureExitTimeoutAndGroupCancel/TimeoutTerminatesGroup / GIT-005, GIT-011 / pass
- gitstore / Test_GIT_005_GIT_011_RunCaptureExitTimeoutAndGroupCancel/ContextCancellationTerminatesGroup / GIT-005, GIT-011 / pass
- gitstore / Test_GIT_005_RunCancelWhileShellLaunchesChild / GIT-005 / pass

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
- app / Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip / APP-111, APP-112, APP-113 / fail

=== RUN   Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip
    ci_cli_test.go:206: default doctor CI: status=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"C11\",\"message\":\"read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-2613054885/report.xml: no such file or directory. Correct the reported issue and run fireweed ci again.\",\"path\":\"alpha\",\"warning\":false},{\"code\":\"X02\",\"message\":\"global boundary hook: exit status 127 without a module-attributed violation. Correct the reported issue and run fireweed ci again.\",\"warning\":false},{\"code\":\"C16\",\"message\":\"AGENTS.md managed block is missing or stale\",\"path\":\"AGENTS.md\",\"warning\":true},{\"code\":\"C14\",\"message\":\"setup setup: hook sh .fireweed/setup.sh: \\u003cnil\\u003e; exit 1;  . Correct the reported issue and run fireweed ci again.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip (0.42s)

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
- app / Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI / APP-062, APP-063, APP-064, APP-065, APP-066, APP-067 / fail

=== RUN   Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI
    doctor_setup_test.go:125: APP-062/066 doctor status=1 stdout="FAIL baseline global eval doctor3204c34f12 eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-3048936237/report.xml: no such file or directory\nFAIL failing_eval global eval doctor3204c34f12 skipped after failed baseline\nFAIL missing_id global eval doctor3204c34f12 skipped after failed baseline\nFAIL baseline global boundaries doctor3204c34f12 boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation  false}]\nFAIL undeclared global boundaries doctor3204c34f12 skipped after failed baseline\nFAIL internal_eval global boundaries doctor36c7417442 skipped after failed baseline\n" stderr="C14: baseline global eval eval: eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-3048936237/report.xml: no such file or directory\nC14: failing_eval global eval eval: skipped after failed baseline\nC14: missing_id global eval eval: skipped after failed baseline\nC14: baseline global boundaries boundary hook: boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation  false}]\nC14: undeclared global boundaries boundary hook: skipped after failed baseline\nC14: internal_eval global boundaries boundary hook: skipped after failed baseline\n"
--- FAIL: Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI (0.82s)

- app / Test_APP_062_DoctorUsesCloserProject / APP-062 / fail

=== RUN   Test_APP_062_DoctorUsesCloserProject
    doctor_setup_test.go:245: APP-062 closer doctor status=1 output="FAIL baseline global eval doctorc5201995ac eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-2631893846/report.xml: no such file or directory\nFAIL failing_eval global eval doctorc5201995ac skipped after failed baseline\nFAIL missing_id global eval doctorc5201995ac skipped after failed baseline\nFAIL baseline global boundaries doctorc5201995ac boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctorc5201995ac boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctorc5201995ac boundary hook: exit status 127 without a module-attributed violation  false}]\nFAIL undeclared global boundaries doctorc5201995ac skipped after failed baseline\nFAIL internal_eval global boundaries doctor56c0df8ca5 skipped after failed baseline\nC14: baseline global eval eval: eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-2631893846/report.xml: no such file or directory\nC14: failing_eval global eval eval: skipped after failed baseline\nC14: missing_id global eval eval: skipped after failed baseline\nC14: baseline global boundaries boundary hook: boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctorc5201995ac boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctorc5201995ac boundary hook: exit status 127 without a module-attributed violation  false}]\nC14: undeclared global boundaries boundary hook: skipped after failed baseline\nC14: internal_eval global boundaries boundary hook: skipped after failed baseline\n"
--- FAIL: Test_APP_062_DoctorUsesCloserProject (0.89s)

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
    next_cli_test.go:172: next exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \"ALP-002\" for ALP-002. Correct the reported problem and use fireweed next to resume.\n"
--- FAIL: Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange (2.44s)

- app / Test_APP_075_NextAllowUnmergedAndUsage / APP-075 / fail

=== RUN   Test_APP_075_NextAllowUnmergedAndUsage
    next_cli_test.go:247: allowed next exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \"ALP-002\" for ALP-002. Correct the reported problem and use fireweed next to resume.\nP04: warning: changes/0001-next.md: change is absent from the local and shared base\n"
--- FAIL: Test_APP_075_NextAllowUnmergedAndUsage (2.62s)

- app / Test_APP_076_NextPlainEmptyQueue / APP-076 / pass
- app / Test_APP_077_NextExitClasses / APP-077 / pass
- app / Test_APP_078_NextWrappedCauseAndAction / APP-078 / pass
- app / Test_APP_079_NextJSONSuccessAndFailure / APP-079 / fail

=== RUN   Test_APP_079_NextJSONSuccessAndFailure
    next_cli_test.go:328: JSON applied exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \\\"ALP-002\\\" for ALP-002. Correct the reported problem and use fireweed next to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_079_NextJSONSuccessAndFailure (2.60s)

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
    run_cli_test.go:262: run exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation (3.39s)

- app / Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON / APP-081, APP-084, APP-086 / fail

=== RUN   Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON
    run_cli_test.go:323: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON (3.43s)

- app / Test_APP_083_RunPlainMultiple / APP-083 / fail

=== RUN   Test_APP_083_RunPlainMultiple
    run_cli_test.go:347: run exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_083_RunPlainMultiple (3.38s)

- app / Test_APP_086_RunJSONLimited / APP-086 / fail

=== RUN   Test_APP_086_RunJSONLimited
    run_cli_test.go:359: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_086_RunJSONLimited (3.37s)

- app / Test_APP_082_RunAllowUnmerged / APP-082 / fail

=== RUN   Test_APP_082_RunAllowUnmerged
    run_cli_test.go:376: run exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_082_RunAllowUnmerged (2.63s)

- app / Test_APP_081_APP_084_RunUsageAndConfiguration / APP-081, APP-084 / pass
- app / Test_APP_084_APP_085_APP_086_RunWrappedFailure / APP-084, APP-085, APP-086 / pass
- app / Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes / APP-084, APP-089, APP-090 / fail

=== RUN   Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes
    run_for_cli_test.go:95: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes (3.53s)

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
    usage_cli_test.go:168: owned resume exit=1 stderr="" stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \\\"ALP-002\\\" for ALP-002. Correct the reported problem and use fireweed next to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n"
    usage_cli_test.go:176: resumed next data={Change: Run: Phase: Empty:false}: <nil>
    usage_cli_test.go:178: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume1248919961/002/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (2.68s)

- app / Test_APP_076_APP_077_NextPlainUsageStop / APP-076, APP-077 / pass
- app / Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume / APP-083, APP-084, APP-086 / fail

=== RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:252: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
    usage_cli_test.go:254: clean stop data=null; want object: <nil>
    usage_cli_test.go:270: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume2831879836/002/changes/archive/0001-alpha.md: no such file or directory
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (3.19s)

- app / Test_APP_083_APP_084_RunPlainSharedUsageStop / APP-083, APP-084 / fail

=== RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:331: plain run stop exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\n"
    usage_cli_test.go:333: plain run lacks alpha: ""
    usage_cli_test.go:333: plain run lacks applied count 1: ""
    usage_cli_test.go:335: plain run lacks resume guidance: ""
    usage_cli_test.go:337: stop reason "" lacks a labeled completed wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled next wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled budget of 5ms
    usage_cli_test.go:339: alpha not applied: stat /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_RunPlainSharedUsageStop2262373806/002/changes/archive/0001-alpha.md: no such file or directory
    usage_cli_test.go:341: held beta state: change has no claim: beta
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (3.26s)

- app / Test_APP_084_APP_086_EarlierBlockPrecedesUsageStop / APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/next / APP-077, APP-079, APP-084, APP-086 / pass
- app / Test_APP_077_APP_079_APP_084_APP_086_InvalidUsagePolicyBeforeMutation/run / APP-077, APP-079, APP-084, APP-086 / pass
- checks / Test_CHK_005_GlobalAndModuleBoundaryHooks / CHK-005 / pass
- checks / Test_CHK_041_ExactBaselineOnly / CHK-041 / pass
- checks / Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns / CHK-039, CHK-040 / pass
- checks / Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional / CHK-044, CHK-045 / pass
- checks / Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut / CHK-041, CHK-042 / pass
- checks / Test_CHK_043_RequiredGlobalHookRejectsUnattributedNonzeroExit / CHK-043 / pass
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
- checks / Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary / CHK-059 / fail

=== RUN   Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary
--- FAIL: Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary (2.11s)

- checks / Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test / CHK-059 / fail

=== RUN   Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test
    doctor_missing_id_test.go:66: Doctor missed the supplied Go canary: checks=[{Name:setup Module: Passed:false Detail:hook go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_missing_id_test.go:76: Doctor did not report passing missing_id recovery: checks=[{Name:setup Module: Passed:false Detail:hook go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
--- FAIL: Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test (0.50s)

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
- checks / Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe / CHK-050, CHK-051, CHK-052, CHK-053, CHK-055, CHK-056, CHK-057, CHK-058, CHK-059, CHK-060, CHK-061, CHK-067 / fail

=== RUN   Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe
    doctor_test.go:345: real Go recipe doctor failed: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing baseline outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing undeclared outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing internal_eval outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing failing_eval outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing missing_id outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing recovery outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:359: clean recovery not reported after each canary: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:368: two distinct throwaway modules absent: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
--- FAIL: Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe (0.55s)

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
- project / Test_PRJ_086_ResolveExplicitWorkerAlias / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/select_listed_alias / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/unknown_alias / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/malformed_mapping / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/empty_shell / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/blank_shell / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/empty_argv / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/empty_executable / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/nonstring_argument / PRJ-086 / pass
- project / Test_PRJ_086_ResolveExplicitWorkerAlias/invalid_command_type / PRJ-086 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault / PRJ-087 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault/select_named_default / PRJ-087 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault/missing_default / PRJ-087 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault/empty_default / PRJ-087 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault/nonstring_default / PRJ-087 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault/unlisted_default / PRJ-087 / pass
- project / Test_PRJ_087_ResolveNamedWorkerDefault/invalid_default_command / PRJ-087 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/select_independent_audit_alias / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/explicitly_name_worker_default / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/inherit_without_audit_section / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/inherit_without_audit_reference / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/empty_audit_reference / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/null_audit_reference / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/nonstring_audit_reference / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/unlisted_audit_reference / PRJ-088 / pass
- project / Test_PRJ_088_ResolveAuditAliasOrInheritWorkerDefault/invalid_audit_command / PRJ-088 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/fixes/scalar / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/fixes/null / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/fixes/nonstring_member / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/fixes/empty_member / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/eval_fix/scalar / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/eval_fix/null / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/eval_fix/nonstring_member / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/eval_fix/empty_member / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/already_satisfied/scalar / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/already_satisfied/null / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/already_satisfied/nonstring_member / PRJ-079 / pass
- project / Test_PRJ_079_CheckRejectsMalformedChangeTargetLists/already_satisfied/empty_member / PRJ-079 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/unknown / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/foreign / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/added_earlier / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/added_later / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/current / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/empty / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/eval_fix/unknown / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/eval_fix/foreign / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/eval_fix/added_earlier / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/eval_fix/added_later / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/eval_fix/current / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/eval_fix/empty / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/removed_in_same_change / PRJ-080 / pass
- project / Test_PRJ_080_CheckFixTargetsAgainstTheProjectionThroughTheirChange/fixes/repeated_target_added_in_same_change / PRJ-080 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/unknown / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/existing_without_edit / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/removed / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/handle_from_earlier_change / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/modified / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/local_handle_before_start / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/assigned_added_id / PRJ-081 / pass
- project / Test_PRJ_081_CheckAlreadySatisfiedTargetsOnlySameChangeAdditionsAndModifications/empty / PRJ-081 / pass
- project / Test_PRJ_057_StampInvalidBasesHasCode / PRJ-057 / pass
- project / Test_PRJ_058_MissingPendingChangeHasSubject / PRJ-058 / pass
- project / Test_PRJ_058_MissingPendingChangeHasSubject/Stamp / PRJ-058 / pass
- project / Test_PRJ_058_MissingPendingChangeHasSubject/Project / PRJ-058 / pass
- project / Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath / PRJ-059, PRJ-068 / pass
- project / Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath/malformed_frontmatter / PRJ-059, PRJ-068 / pass
- project / Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath/malformed_spec_heading / PRJ-059, PRJ-068 / pass
- project / Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath/duplicate_spec_id / PRJ-059, PRJ-068 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/init_range / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/full_equality / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/strict_lower_bound / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/inclusive_bounds / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/strict_upper_patch / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/numeric_order_with_leading_zero / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/large_decimal_component / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/surrounding_whitespace / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/omitted / PRJ-082 / pass
- project / Test_PRJ_082_CheckCompatibilityAcceptsMatchingPins/omitted_patch_is_zero / PRJ-082 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/empty / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/nonstring / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/double_equals / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/missing_operator / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/missing_minor / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/extra_component / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/negative_component / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/prerelease / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/empty_comparison / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/invalid_after_exclusion / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/strict_upper_endpoint / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/strict_lower_endpoint / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/second_comparison_excludes / PRJ-083 / pass
- project / Test_PRJ_083_CheckCompatibilityReportsInvalidAndIncompatiblePins/documented_range / PRJ-083 / pass
- project / Test_PRJ_084_LoadAndCheckUseLiveBinaryPin / PRJ-084 / pass
- project / Test_PRJ_084_LoadAndCheckUseLiveBinaryPin/omitted / PRJ-084 / pass
- project / Test_PRJ_084_LoadAndCheckUseLiveBinaryPin/compatible / PRJ-084 / pass
- project / Test_PRJ_084_LoadAndCheckUseLiveBinaryPin/incompatible / PRJ-084 / pass
- project / Test_PRJ_084_LoadAndCheckUseLiveBinaryPin/malformed / PRJ-084 / pass
- project / Test_PRJ_084_LoadAndCheckUseLiveBinaryPin/null / PRJ-084 / pass
- project / Test_PRJ_047_LoadAtReadsRecordedBinaryPin / PRJ-047 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/explicit_initial / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/omitted / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/compatible_pin_with_initial / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/future / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/compatible_pin_with_future / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/zero / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/string / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/null / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/float / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/boolean / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/sequence / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/mapping / PRJ-085 / pass
- project / Test_PRJ_085_CheckCompatibilityProjectFormat/excluding_pin_with_initial / PRJ-085 / pass
- project / Test_PRJ_085_LoadAndCheckRejectUnsupportedLiveFormat / PRJ-085 / pass
- project / Test_PRJ_047_LoadAtReadsUnsupportedRecordedFormat / PRJ-047 / pass
- project / Test_PRJ_014_CheckHandlesInChangeProse / PRJ-014 / pass
- project / Test_PRJ_014_CheckHandlesInChangeProse/undeclared_in_motivation / PRJ-014 / pass
- project / Test_PRJ_014_CheckHandlesInChangeProse/undeclared_in_interface / PRJ-014 / pass
- project / Test_PRJ_014_CheckHandlesInChangeProse/declared_in_added / PRJ-014 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_fixes / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_added / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_modified / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_with_removed / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/eval_fix_only / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/fixes_only / PRJ-046 / pass
- project / Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits/added_only / PRJ-046 / pass
- project / Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults / PRJ-001, PRJ-017 / pass
- project / Test_PRJ_001_LoadRootModuleForDotGlob / PRJ-001 / pass
- project / Test_PRJ_018_LoadInvalidDocuments / PRJ-018 / pass
- project / Test_PRJ_018_LoadInvalidDocuments/missing_config / PRJ-018 / pass
- project / Test_PRJ_018_LoadInvalidDocuments/malformed_config / PRJ-018 / pass
- project / Test_PRJ_018_LoadInvalidDocuments/missing_module_document / PRJ-018 / pass
- project / Test_PRJ_018_LoadInvalidDocuments/malformed_module_document / PRJ-018 / pass
- project / Test_PRJ_002_Graph / PRJ-002 / pass
- project / Test_PRJ_002_Graph/PRJ-002_acyclic / PRJ-002 / pass
- project / Test_PRJ_002_Graph/PRJ-019_empty_name / PRJ-002, PRJ-019 / pass
- project / Test_PRJ_002_Graph/PRJ-020_empty_prefix / PRJ-002, PRJ-020 / pass
- project / Test_PRJ_002_Graph/PRJ-019_duplicate_name / PRJ-002, PRJ-019 / pass
- project / Test_PRJ_002_Graph/PRJ-020_duplicate_prefix / PRJ-002, PRJ-020 / pass
- project / Test_PRJ_002_Graph/PRJ-021_missing_dependency / PRJ-002, PRJ-021 / pass
- project / Test_PRJ_002_Graph/PRJ-021_self_dependency / PRJ-002, PRJ-021 / pass
- project / Test_PRJ_002_Graph/PRJ-022_cycle / PRJ-002, PRJ-022 / pass
- project / Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering / PRJ-003, PRJ-023 / pass
- project / Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite / PRJ-024 / pass
- project / Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite/pending / PRJ-024 / pass
- project / Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite/archived / PRJ-024 / pass
- project / Test_PRJ_025_NewChangeRejectsUnknownModuleWithoutWritingFiles / PRJ-025 / pass
- project / Test_PRJ_026_NewChangeRejectsInvalidSlugs / PRJ-026 / pass
- project / Test_PRJ_026_NewChangeRejectsInvalidSlugs/uppercase_slug / PRJ-026 / pass
- project / Test_PRJ_026_NewChangeRejectsInvalidSlugs/underscore_slug / PRJ-026 / pass
- project / Test_PRJ_026_NewChangeRejectsInvalidSlugs/slash_slug / PRJ-026 / pass
- project / Test_PRJ_026_NewChangeRejectsInvalidSlugs/empty_slug / PRJ-026 / pass
- project / Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause / PRJ-067 / pass
- project / Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive / PRJ-027, PRJ-028 / pass
- project / Test_PRJ_029_QueueAndProjectPreservePendingParseDiagnostics / PRJ-029 / pass
- project / Test_PRJ_029_QueueAndProjectPreservePendingParseDiagnostics/malformed_frontmatter / PRJ-029 / pass
- project / Test_PRJ_029_QueueAndProjectPreservePendingParseDiagnostics/malformed_item_heading / PRJ-029 / pass
- project / Test_PRJ_029_QueueAndProjectPreservePendingParseDiagnostics/duplicate_added_identity / PRJ-029 / pass
- project / Test_PRJ_029_QueueAndProjectPreservePendingParseDiagnostics/duplicate_modified_identity / PRJ-029 / pass
- project / Test_PRJ_029_QueueAndProjectPreservePendingParseDiagnostics/duplicate_removed_identity / PRJ-029 / pass
- project / Test_PRJ_030_QueueRejectsUnknownModule / PRJ-030 / pass
- project / Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug / PRJ-004, PRJ-031, PRJ-032 / pass
- project / Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles / PRJ-005, PRJ-032, PRJ-036 / pass
- project / Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections / PRJ-033, PRJ-034, PRJ-035 / pass
- project / Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-034_identity_name / PRJ-033, PRJ-034, PRJ-035 / pass
- project / Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-034_identity_prefix / PRJ-033, PRJ-034, PRJ-035 / pass
- project / Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-034_unknown_field / PRJ-033, PRJ-034, PRJ-035 / pass
- project / Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections/PRJ-035_invalid_graph / PRJ-033, PRJ-034, PRJ-035 / pass
- project / Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation / PRJ-006, PRJ-038 / pass
- project / Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID / PRJ-037, PRJ-039 / pass
- project / Test_PRJ_007_CheckValidAndDiagnostics / PRJ-007 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics / PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-008_missing_motivation / PRJ-008, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-009_nonnormative_module / PRJ-009, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-009_nonnormative_change / PRJ-009, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-010_invalid_review_tag / PRJ-010, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-010_invalid_quarantine_tag / PRJ-010, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-011_PRJ-068_duplicate_module_id / PRJ-011, PRJ-015, PRJ-016, PRJ-068 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_module_id / PRJ-012, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_change_id / PRJ-012, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-013_duplicate_edit_identity / PRJ-013, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-014_undeclared_handle / PRJ-014, PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-015_missing_base / PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-015_stale_base / PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-016_absent_modified / PRJ-015, PRJ-016 / pass
- project / Test_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-016_absent_removed / PRJ-015, PRJ-016 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings/item_body / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings/leading_prose / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings/prose_only / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings/second_item_body / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings/empty / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedSectionContainsOnlyHeadings/headings_with_whitespace / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedPreservesHeadingIdentityAndBaseDiagnostics / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedPreservesHeadingIdentityAndBaseDiagnostics/malformed_heading / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedPreservesHeadingIdentityAndBaseDiagnostics/duplicate_identity / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedPreservesHeadingIdentityAndBaseDiagnostics/absent_target / PRJ-078 / pass
- project / Test_PRJ_078_CheckRemovedPreservesHeadingIdentityAndBaseDiagnostics/stale_base / PRJ-078 / pass
- project / Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments / PRJ-047 / pass
- project / Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones / PRJ-048 / pass
- project / Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs / PRJ-049 / pass
- project / Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/invalid_revision / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/missing_recorded_config / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/malformed_recorded_config / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs/missing_recorded_SPEC / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtPreservesRecordedSpecParseDiagnostics / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtPreservesRecordedSpecParseDiagnostics/malformed_frontmatter / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtPreservesRecordedSpecParseDiagnostics/malformed_item_heading / PRJ-050 / pass
- project / Test_PRJ_050_LoadAtPreservesRecordedSpecParseDiagnostics/duplicate_item_identity / PRJ-050 / pass
- project / Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs / PRJ-004 / pass
- project / Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs/Project / PRJ-004 / pass
- project / Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs/Stamp / PRJ-004 / pass
- project / Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs / PRJ-066 / pass
- project / Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs/StaleSpecs / PRJ-066 / pass
- project / Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs/StampMissing / PRJ-066 / pass
- project / Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications / PRJ-051 / pass
- project / Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate / PRJ-053 / pass
- project / Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata / PRJ-052 / pass
- project / Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule / PRJ-054 / pass
- project / Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec / PRJ-055 / pass
- project / Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors / PRJ-056 / pass
- project / Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/config_missing / PRJ-056 / pass
- project / Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/config_malformed / PRJ-056 / pass
- project / Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/archive_document / PRJ-056 / pass
- project / Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/archive_added / PRJ-056 / pass
- project / Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors/archive_modified / PRJ-056 / pass
- project / Test_PRJ_071_LoadSpecLimits / PRJ-071 / pass
- project / Test_PRJ_071_LoadSpecLimits/omitted / PRJ-071 / pass
- project / Test_PRJ_071_LoadSpecLimits/empty_mapping / PRJ-071 / pass
- project / Test_PRJ_071_LoadSpecLimits/partial / PRJ-071 / pass
- project / Test_PRJ_071_LoadSpecLimits/configured / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01 / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/scalar_section / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/null_section / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/sequence_section / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/zero_clauses / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/negative_words / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/quoted_decision_words / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/floating_point_clauses / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/boolean_words / PRJ-071 / pass
- project / Test_PRJ_071_LoadRejectsInvalidSpecLimitsWithC01/null_decision_words / PRJ-071 / pass
- project / Test_PRJ_072_CheckWarnsForOversizedCurrentAndPendingItems / PRJ-072 / pass
- project / Test_PRJ_072_CheckWarnsForOversizedCurrentAndPendingItems/current / PRJ-072 / pass
- project / Test_PRJ_073_CheckWarnsForNormalizedSentencesAcrossProjectedItems / PRJ-073 / pass
- project / Test_PRJ_073_CheckCountsDistinctItemsInTheFinalProjection / PRJ-073 / pass
- project / Test_PRJ_073_CheckCountsDistinctItemsInTheFinalProjection/empty_queue / PRJ-073 / pass
- project / Test_PRJ_074_CheckWarnsOnlyForModulesOverTheDecisionsLimit / PRJ-074 / pass
- project / Test_PRJ_070_CheckWarnsForTrackedSpecsOutsideModuleGlobs / PRJ-070 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/absent / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/duplicate_with_repeated_id / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/duplicate_with_atx_closing_hashes / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/single_empty / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/single_populated / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/single_atx_heading_with_crlf / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/single_with_backtick_example / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/single_with_tilde_example / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/fenced_example_only / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/nonmatching_headings_only / PRJ-077 / pass
- project / Test_PRJ_077_CheckRequiresExactlyOneSpecsSectionPerModule/pending_change_without_specs / PRJ-077 / pass
- project / Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder / PRJ-060 / pass
- project / Test_PRJ_060_StaleSpecsRejectsNonmappingBasesAtQueuedPath / PRJ-060 / pass
- project / Test_PRJ_060_StaleSpecsRejectsNonmappingBasesAtQueuedPath/sequence / PRJ-060 / pass
- project / Test_PRJ_060_StaleSpecsRejectsNonmappingBasesAtQueuedPath/scalar / PRJ-060 / pass
- project / Test_PRJ_060_StaleSpecsRejectsNonmappingBasesAtQueuedPath/null / PRJ-060 / pass
- project / Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash / PRJ-061, PRJ-062 / pass
- project / Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem / PRJ-062 / pass
- project / Test_PRJ_063_StaleSpecsEmptyAndReadOnly / PRJ-063 / pass
- project / Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent / PRJ-064 / pass
- project / Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically / PRJ-065 / pass
- project / Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically/nonmapping / PRJ-065 / pass
- project / Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically/absent_modified_after_valid / PRJ-065 / pass
- project / Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically/absent_removed_after_valid / PRJ-065 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/target / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/0001-target.md / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/changes/0001-target.md / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StaleSpecs_unknown / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StampMissing_unknown / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StaleSpecs_ambiguous / PRJ-066 / pass
- project / Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StampMissing_ambiguous / PRJ-066 / pass
- project / Test_PRJ_040_CheckConfigurationCauses / PRJ-040 / pass
- project / Test_PRJ_040_CheckConfigurationCauses/missing / PRJ-040 / pass
- project / Test_PRJ_040_CheckConfigurationCauses/malformed / PRJ-040 / pass
- project / Test_PRJ_041_CheckGraphCauses / PRJ-041 / pass
- project / Test_PRJ_041_CheckGraphCauses/current / PRJ-041 / pass
- project / Test_PRJ_041_CheckGraphCauses/projected / PRJ-041 / pass
- project / Test_PRJ_042_CheckDocumentAndItemCauses / PRJ-042 / pass
- project / Test_PRJ_042_CheckDocumentAndItemCauses/missing_module / PRJ-042 / pass
- project / Test_PRJ_042_CheckDocumentAndItemCauses/malformed_module / PRJ-042 / pass
- project / Test_PRJ_042_CheckDocumentAndItemCauses/malformed_change / PRJ-042 / pass
- project / Test_PRJ_042_CheckDocumentAndItemCauses/malformed_item / PRJ-042 / pass
- project / Test_PRJ_043_CheckMetadataAndSetCauses / PRJ-043 / pass
- project / Test_PRJ_043_CheckMetadataAndSetCauses/metadata_type / PRJ-043 / pass
- project / Test_PRJ_043_CheckMetadataAndSetCauses/zone_type / PRJ-043 / pass
- project / Test_PRJ_043_CheckMetadataAndSetCauses/identity_set_key / PRJ-043 / pass
- project / Test_PRJ_043_CheckMetadataAndSetCauses/unknown_set_key / PRJ-043 / pass
- project / Test_PRJ_043_CheckMetadataAndSetCauses/set_value_type / PRJ-043 / pass
- project / Test_PRJ_044_CheckQueueIdentityCause / PRJ-044 / pass
- project / Test_PRJ_069_RegularFileAtChangesReportsQueueCause / PRJ-069 / pass
- project / Test_PRJ_045_CheckPendingModuleTargetCause / PRJ-045 / pass
- project / Test_PRJ_013_CheckDuplicateWithinEditSection / PRJ-013 / pass
- project / Test_PRJ_075_CheckPossibleEvalDisposablePaths / PRJ-075 / pass
- project / Test_PRJ_075_CheckPossibleEvalDisposablePaths/identical / PRJ-075 / pass
- project / Test_PRJ_075_CheckPossibleEvalDisposablePaths/recursive_and_suffix / PRJ-075 / pass
- project / Test_PRJ_075_CheckPossibleEvalDisposablePaths/zero_recursive_components / PRJ-075 / pass
- project / Test_PRJ_075_CheckPossibleEvalDisposablePaths/disjoint_siblings / PRJ-075 / pass
- project / Test_PRJ_075_CheckPossibleEvalDisposablePaths/single_component_wildcard / PRJ-075 / pass
- project / Test_PRJ_076_CheckResolvedZonesThroughTheQueue / PRJ-076 / pass
- project / Test_PRJ_076_CheckResolvedZonesThroughTheQueue/current_overlap_cleared_by_first_change / PRJ-076 / pass
- project / Test_PRJ_076_CheckResolvedZonesThroughTheQueue/explicit_empty_lists_override_defaults / PRJ-076 / pass
- project / Test_PRJ_076_CheckResolvedZonesThroughTheQueue/set_edits_create_overlap_at_tip_and_intermediate_projection / PRJ-076 / pass
- project / Test_PRJ_076_CheckResolvedZonesThroughTheQueue/set_edits_create_overlap_at_tip_and_intermediate_projection/tip / PRJ-076 / pass
- project / Test_PRJ_076_CheckResolvedZonesThroughTheQueue/set_edits_create_overlap_at_tip_and_intermediate_projection/intermediate / PRJ-076 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/failure / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/timeout / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/cancellation / RUN-063 / pass
- runner / Test_RUN_063_InvokeAgentPersistsStreamsOnFailureTimeoutCancellationAndAppend/append / RUN-063 / pass
- runner / Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput / RUN-064 / fail

=== RUN   Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput
    agent_logs_test.go:109: Next fixture = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074216-5e949a32f53bedf7ce71ffd8"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:16.780134Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:17.15557Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074216-5e949a32f53bedf7ce71ffd8", Phase:"impl", Head:"d920b7f6089b1f8e590e1ecb79fb113967820cef", PhaseBase:"d920b7f6089b1f8e590e1ecb79fb113967820cef", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput (1.99s)

- runner / Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries / RUN-065 / pass
- runner / Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries/usage_once / RUN-065 / pass
- runner / Test_RUN_065_DrillLogRunIsStableAcrossUsageAndReconstructionRetries/agent_retry / RUN-065 / pass
- runner / Test_RUN_066_InvokeAgentLogContextDefaultsAndExplicitFields / RUN-066 / pass
- runner / Test_RUN_041_UnapprovedChangeRequiresApproval / RUN-041 / pass
- runner / Test_RUN_041_RunPreservesApprovalRefusalCode / RUN-041 / pass
- runner / Test_RUN_042_OtherNextErrorsDoNotRequireApproval / RUN-042 / pass
- runner / Test_RUN_042_OtherNextErrorsDoNotRequireApproval/start_failure / RUN-042 / pass
- runner / Test_RUN_042_OtherNextErrorsDoNotRequireApproval/agent_failure / RUN-042 / pass
- runner / Test_RUN_021_AuditCorrectionWithinAllowance / RUN-021 / fail

=== RUN   Test_RUN_021_AuditCorrectionWithinAllowance
--- FAIL: Test_RUN_021_AuditCorrectionWithinAllowance (6.39s)

- runner / Test_RUN_021_AuditCorrectionWithinAllowance/correction / RUN-021 / fail

=== RUN   Test_RUN_021_AuditCorrectionWithinAllowance/correction
    audit_return_attempts_test.go:123: audit correction ended in "impl": impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_AuditCorrectionWithinAllowance/correction (3.17s)

- runner / Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry / RUN-021 / fail

=== RUN   Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry
    audit_return_attempts_test.go:123: audit correction ended in "impl": impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry (3.22s)

- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals / RUN-021, RUN-070 / fail

=== RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals (5.46s)

- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/later_return / RUN-021, RUN-070 / pass
- runner / Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry / RUN-021, RUN-070 / fail

=== RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry
    audit_return_attempts_test.go:154: phases captured=[evals audit evals audit] invoked=[evals audit evals audit], want [evals evals audit]
    audit_return_attempts_test.go:166: phase commits before block = 4, want 2
    audit_return_attempts_test.go:187: phase base did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T074230-23913473b684f9a0e3d7803d\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:30.982366Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:31.363375Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T07:42:31.962282Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:32.342388Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2"
    audit_return_attempts_test.go:189: blocked change did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nblocked: evals attempt allowance exhausted after 2 attempts\nmodule: alpha\nphase: evals\nrun: 20261118T074230-23913473b684f9a0e3d7803d\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:30.982366Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:31.363375Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T07:42:31.962282Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:32.342388Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2\n\n"
    audit_return_attempts_test.go:194: returned state lost the latest audit finding: "\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2\n\n"
    audit_return_attempts_test.go:197: saved attempt did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T074230-23913473b684f9a0e3d7803d\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:30.982366Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:31.363375Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T07:42:31.962282Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:32.342388Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2"
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry (2.64s)

- runner / Test_RUN_021_AuditRoundLimitCanStopBeforeEvalsAllowance / RUN-021 / pass
- runner / Test_RUN_070_RunContinuesIndependentWorkAfterAuditReturnBlock / RUN-070 / pass
- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit / RUN-013, RUN-045, RUN-046, RUN-071 / fail

=== RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit (4.88s)

- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override / RUN-013, RUN-045, RUN-046, RUN-071 / fail

=== RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074239-d8e068a30b313b955a9fdbd0"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:40.239572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:40.652163Z agent=bash fireweed=0.2.1 playbook=a1eef0b302afb0055dd630de8b87b5ff467a00f40d16783ba79821ee228a5cf6 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074239-d8e068a30b313b955a9fdbd0", Phase:"impl", Head:"5f46ed8bcd504cd9ba5d37d4f360d923d60d5caf", PhaseBase:"5f46ed8bcd504cd9ba5d37d4f360d923d60d5caf", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override (2.45s)

- runner / Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback / RUN-013, RUN-045, RUN-046, RUN-071 / fail

=== RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074242-b5609f79c0a71407e0d82918"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:42.660238Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:43.075258Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074242-b5609f79c0a71407e0d82918", Phase:"impl", Head:"6478a2461bec24bccb26d8a9d9100ba6ea6da968", PhaseBase:"6478a2461bec24bccb26d8a9d9100ba6ea6da968", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback (2.42s)

- runner / Test_RUN_013_AuditRetryUsesApprovedAgentConfiguration / RUN-013 / pass
- runner / Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry / RUN-048 / fail

=== RUN   Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry
    audit_workspace_test.go:210: audit verification retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074245-41fd0f1873bdfaa0915f8e2c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:46.196144Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:46.735827Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074245-41fd0f1873bdfaa0915f8e2c", Phase:"impl", Head:"6f938f2978c3826be5c55c4973d5de9f9a9b25f3", PhaseBase:"6f938f2978c3826be5c55c4973d5de9f9a9b25f3", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/audit-2.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry (2.54s)

- runner / Test_RUN_046_AuditCommitDoesNotEnterHeldGit / RUN-046 / fail

=== RUN   Test_RUN_046_AuditCommitDoesNotEnterHeldGit
    audit_workspace_test.go:232: audit commit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074248-ef9e8b131e0a5c24c66480fe"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:48.892307Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:49.317526Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074248-ef9e8b131e0a5c24c66480fe", Phase:"impl", Head:"5bd546b8147a182190eb3eb4b26c14b4f5815719", PhaseBase:"5bd546b8147a182190eb3eb4b26c14b4f5815719", ClaimOID:"", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_046_AuditCommitDoesNotEnterHeldGit (2.69s)

- runner / Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported / RUN-047, RUN-048 / fail

=== RUN   Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported
    audit_workspace_test.go:280: audit before Provenance lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074250-a0af3dd5b3ca84e1a25bf268"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Audit\n\nNo findings; checked ALP-002\n\n## Provenance\n\n- evals 2026-11-18T07:42:51.422456Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:51.980842Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074250-a0af3dd5b3ca84e1a25bf268", Phase:"impl", Head:"0c69be8c3798128dad2239b8e631094df6feea50", PhaseBase:"0c69be8c3798128dad2239b8e631094df6feea50", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001/.fireweed/runs/20261118T074250-a0af3dd5b3ca84e1a25bf268/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001/.fireweed/runs/20261118T074250-a0af3dd5b3ca84e1a25bf268/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001/.fireweed/runs/20261118T074250-a0af3dd5b3ca84e1a25bf268/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported (2.55s)

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
- runner / Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle / RUN-027, RUN-028 / fail

=== RUN   Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle
    commit_recovery_test.go:38: Next after fixture agent commits: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002; state=model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074310-1276a3290904e5d5d6a6a845"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:43:11.182746Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:43:11.596746Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074310-1276a3290904e5d5d6a6a845", Phase:"impl", Head:"e37f814e581328035e80bd367856bbd06badd3a7", PhaseBase:"e37f814e581328035e80bd367856bbd06badd3a7", ClaimOID:"", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}; steps=[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/impl-2.log"}}}
--- FAIL: Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle (2.66s)

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
    next_test.go:428: Next blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002")
    next_test.go:431: fixture alpha did not block: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"alpha", Path:"changes/0001-alpha.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"", "module":"alpha", "phase":"impl", "run":"20261118T074404-02d870a893c1d79ea5ba4251"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:04.99338Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:05.528544Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002", "exhaust-alpha", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# alpha\n\n## Specs\n\n### ALP-001\nalpha MUST return a value.\n\n### ALP-002\nalpha MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# beta\n\n## Specs\n\n### BET-001\nbeta MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"GAM", "name":"gamma"}, Body:"# gamma\n\n## Specs\n\n### GAM-001\ngamma MUST return a value.\n\n## Decisions\n"}}}}, Run:"20261118T074404-02d870a893c1d79ea5ba4251", Phase:"impl", Head:"d1c34b42cc6670a148d16b2a46c39e8f7ff4f372", PhaseBase:"d1c34b42cc6670a148d16b2a46c39e8f7ff4f372", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-no-edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"alpha-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
--- FAIL: Test_RUN_011_WaitingHeadSelectsIndependent (2.33s)

- runner / Test_RUN_010_EntirelyWaitingQueue / RUN-010 / fail

=== RUN   Test_RUN_010_EntirelyWaitingQueue
    next_test.go:449: Next blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002")
    next_test.go:452: fixture alpha did not block: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"alpha", Path:"changes/0001-alpha.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"", "module":"alpha", "phase":"impl", "run":"20261118T074406-f95fc7385dd40d14e6c424bf"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:07.334363Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:07.854304Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002", "exhaust-alpha", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# alpha\n\n## Specs\n\n### ALP-001\nalpha MUST return a value.\n\n### ALP-002\nalpha MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# beta\n\n## Specs\n\n### BET-001\nbeta MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"GAM", "name":"gamma"}, Body:"# gamma\n\n## Specs\n\n### GAM-001\ngamma MUST return a value.\n\n## Decisions\n"}}}}, Run:"20261118T074406-f95fc7385dd40d14e6c424bf", Phase:"impl", Head:"86a0d0046636bf39356a94f4f29894fb36aa11b6", PhaseBase:"86a0d0046636bf39356a94f4f29894fb36aa11b6", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-no-edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"alpha-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
--- FAIL: Test_RUN_010_EntirelyWaitingQueue (2.30s)

- runner / Test_RUN_013_RUN_014_CompleteSpecChange / RUN-013, RUN-014 / fail

=== RUN   Test_RUN_013_RUN_014_CompleteSpecChange
    next_test.go:478: Next: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074409-e109af499ffad963aa5ac6ae"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:09.790834Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:10.163755Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074409-e109af499ffad963aa5ac6ae", Phase:"impl", Head:"fbbb7c9cd3bd239785d2f0b034a68abf992b1e58", PhaseBase:"fbbb7c9cd3bd239785d2f0b034a68abf992b1e58", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_013_RUN_014_CompleteSpecChange (2.39s)

- runner / Test_RUN_015_InvocationFailureRetainsHeldPhase / RUN-015 / pass
- runner / Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks / RUN-023, RUN-024 / fail

=== RUN   Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks
    next_test.go:678: Next error code = "E01" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002)
    next_test.go:681: evals exhaustion cause = impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
    next_test.go:684: evals exhaustion steps = 4, want 2
    next_test.go:688: captured invocation 1 = model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/001/.fireweed/runs/20261118T074412-10877f23019fb888d8c42713/audit-1.log"}}
    next_test.go:688: captured invocation 2 = model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"no edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/001/.fireweed/runs/20261118T074412-10877f23019fb888d8c42713/impl-1.log"}}
    next_test.go:688: captured invocation 3 = model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"no edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/001/.fireweed/runs/20261118T074412-10877f23019fb888d8c42713/impl-2.log"}}
    next_test.go:692: agent calls = 4, want 2
    next_test.go:694: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/002/evals-final-head: no such file or directory
--- FAIL: Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks (1.79s)

- runner / Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase / RUN-021, RUN-022 / fail

=== RUN   Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase
    next_test.go:706: Next after retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074413-e80248025c074b7568278683"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:14.307437Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:14.686686Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074413-e80248025c074b7568278683", Phase:"impl", Head:"e9d5dadb411b19ea59287a95fbb817086961e416", PhaseBase:"e9d5dadb411b19ea59287a95fbb817086961e416", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase (1.99s)

- runner / Test_RUN_021_ConfiguredThreeAttempts / RUN-021 / fail

=== RUN   Test_RUN_021_ConfiguredThreeAttempts
    next_test.go:756: Next with three attempts = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 3 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074415-d3bc668f21f3067c2f7bd67d"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:16.523918Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:16.905793Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002", "retry-third", "{prompt}"}}, "default":"fixture-worker"}, "attempts":3, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074415-d3bc668f21f3067c2f7bd67d", Phase:"impl", Head:"3a5205442a6486d14f3f251fb8b8413fb4707938", PhaseBase:"3a5205442a6486d14f3f251fb8b8413fb4707938", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-2", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/impl-2.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/impl-3.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_ConfiguredThreeAttempts (2.34s)

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
- runner / Test_RUN_016_UnlimitedQueueOrder / RUN-016 / fail

=== RUN   Test_RUN_016_UnlimitedQueueOrder
    run_test.go:313: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/beta/1): failed case "BET-002" for BET-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002
--- FAIL: Test_RUN_016_UnlimitedQueueOrder (7.12s)

- runner / Test_RUN_016_NegativeMaxIsUnlimited / RUN-016 / fail

=== RUN   Test_RUN_016_NegativeMaxIsUnlimited
    run_test.go:329: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
--- FAIL: Test_RUN_016_NegativeMaxIsUnlimited (2.49s)

- runner / Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding / RUN-017, RUN-018, RUN-019 / fail

=== RUN   Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding
    run_test.go:341: first Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/beta/1): failed case "BET-002" for BET-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002
--- FAIL: Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding (7.14s)

- runner / Test_RUN_018_InitialEmptyAndExactLimit / RUN-018 / fail

=== RUN   Test_RUN_018_InitialEmptyAndExactLimit
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit (4.96s)

- runner / Test_RUN_018_InitialEmptyAndExactLimit/initial_empty / RUN-018 / pass
- runner / Test_RUN_018_InitialEmptyAndExactLimit/exact_limit / RUN-018 / fail

=== RUN   Test_RUN_018_InitialEmptyAndExactLimit/exact_limit
    run_test.go:385: exact Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/beta/1): failed case "BET-002" for BET-002
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit/exact_limit (4.78s)

- runner / Test_RUN_020_HeldFailureRetainsEarlierResults / RUN-020 / fail

=== RUN   Test_RUN_020_HeldFailureRetainsEarlierResults
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_020_HeldFailureRetainsEarlierResults (2.79s)

- runner / Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent / RUN-016, RUN-025, RUN-026 / fail

=== RUN   Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent
    run_test.go:427: Run blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002)
    run_test.go:427: Run blocking error = impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002, want verification cause and refs/fireweed/attempts/alpha/1
    run_test.go:427: alpha block state = "---\nbases: {}\nblocked: 'impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"'\nmodule: alpha\nphase: impl\nrun: 20261118T074451-e91a7f0322043a6ec6a63f9d\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:52.217632Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:52.750661Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent (4.74s)

- runner / Test_RUN_025_RUN_026_PositiveMaxRetainsBlock / RUN-025, RUN-026 / fail

=== RUN   Test_RUN_025_RUN_026_PositiveMaxRetainsBlock
    run_test.go:448: Run blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002)
    run_test.go:448: Run blocking error = impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002, want verification cause and refs/fireweed/attempts/alpha/1
    run_test.go:448: alpha block state = "---\nbases: {}\nblocked: 'impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"'\nmodule: alpha\nphase: impl\nrun: 20261118T074456-4627920bb90d5fe811e6e4e3\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:56.99838Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:57.529869Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_025_RUN_026_PositiveMaxRetainsBlock (4.78s)

- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt / RUN-014 / pass
- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt/pass / RUN-014 / pass
- runner / Test_RUN_014_NextRunsFullVerificationOncePerAttempt/dependent_retry / RUN-014 / pass
- runner / Test_RUN_072_RetainsStartWarning / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning
--- FAIL: Test_RUN_072_RetainsStartWarning (10.71s)

- runner / Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery
    start_warnings_test.go:81: Next: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery (2.85s)

- runner / Test_RUN_072_RetainsStartWarning/verification_exhaustion / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/invocation_failure / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/usage_stop / RUN-072 / pass
- runner / Test_RUN_072_RetainsStartWarning/Run_applied_result / RUN-072 / fail

=== RUN   Test_RUN_072_RetainsStartWarning/Run_applied_result
    start_warnings_test.go:147: expected Run to apply one queued change, got 0 applied, empty=false, limited=false, err=impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_072_RetainsStartWarning/Run_applied_result (2.46s)

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
- runner / Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep / RUN-030, RUN-031, RUN-033 / fail

=== RUN   Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep
    usage_test.go:206: owned resume = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 1 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074636-2705e1c8f4aaa1254bb3c9c3"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:46:36.859721Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:46:37.233548Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":1, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074636-2705e1c8f4aaa1254bb3c9c3", Phase:"impl", Head:"69bc14f4c1fb6f5ed776cd1250621821f266693f", PhaseBase:"69bc14f4c1fb6f5ed776cd1250621821f266693f", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/.fireweed/runs/20261118T074636-2705e1c8f4aaa1254bb3c9c3/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/.fireweed/runs/20261118T074636-2705e1c8f4aaa1254bb3c9c3/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/.fireweed/runs/20261118T074636-2705e1c8f4aaa1254bb3c9c3/impl-1.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
    usage_test.go:208: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep (2.39s)

- runner / Test_RUN_033_OwnedResumeUsesApprovedUsageBudget / RUN-033 / pass
- runner / Test_RUN_030_UsageRetryPreservesVerificationFeedback / RUN-030 / fail

=== RUN   Test_RUN_030_UsageRetryPreservesVerificationFeedback
    usage_test.go:288: verification and usage retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074639-26311bead8c44ed47d206aca"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:46:39.649858Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:46:40.023355Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":2, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002"}}, "limit_budget":"20ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074639-26311bead8c44ed47d206aca", Phase:"impl", Head:"93e110f016f309f23e1ec864747796a5b74cd431", PhaseBase:"93e110f016f309f23e1ec864747796a5b74cd431", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"first verification", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_030_UsageRetryPreservesVerificationFeedback (2.18s)

- runner / Test_RUN_030_RUN_031_WaitBudgetSharedAcrossNextPhases / RUN-030, RUN-031 / pass
- runner / Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied / RUN-032 / fail

=== RUN   Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied
    usage_test.go:329: Run stop = impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        usage limit stopped execution: completed wait 5ms, next wait 5ms, budget 5ms
    usage_test.go:331: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied (2.85s)

- runner / Test_RUN_034_UsageStopClassifierWrapping / RUN-034 / pass
- runner / Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure / RUN-035 / fail

=== RUN   Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure
    usage_test.go:371: combined block and usage error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        usage limit stopped execution: completed wait 0s, next wait 5ms, budget 0s)
    usage_test.go:374: combined block and usage failure = model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        usage limit stopped execution: completed wait 0s, next wait 5ms, budget 0s
    usage_test.go:380: alpha block state = "---\nbases: {}\nblocked: 'impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"'\nmodule: alpha\nphase: impl\nrun: 20261118T074646-a18b30c12b89b65703d935d5\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:46:46.805926Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:46:47.331203Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"
--- FAIL: Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure (2.81s)

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

## Contradicted Decisions
None observed. Decision [0010-manual-lifecycle-git] ("Change.Status uses the
leading letter of Git's name-status code for tracked changes; untracked paths
use `?`") is consistent with the implementation and all eval evidence:
trung `R100` reports `R`, untracked paths report `?`, and the unstaged-rename
eval passes with `OldPath` retained.

## Gaps
- GIT-005 (target: gitstore): The spec does not say whether a supplied
  `Request.Env` replaces the process environment or merges with it. The
  implementation treats a non-nil map as the complete environment (nil
  inherits the parent process). Fixing change would tighten GIT-005 wording
  or the `Request` doc comment.
- GIT-005 (target: gitstore): The spec says the launched process group is
  terminated on timeout/cancellation, but names no bound on how long `Run`
  may keep waiting on pipe EOF after the direct child exits. A grandchild
  forked in the same instant the SIGKILL lands can survive the group kill
  (observed in practice) while holding inherited stdout/stderr pipes;
  `exec` then blocks in `Wait` until it exits (~5s in the fixture), which
  breaks prompt-return expectations (eval: "terminated Run did not return
  promptly"). The implementation uses `Cmd.WaitDelay = 200ms` to force-close
  the pipes shortly after the child dies. Fixing change belongs in gitstore
  (either spec text or a documented `WaitDelay` behavior).
- GIT-010 (target: gitstore): The spec says "exact current object matches the
  expected old value" but never defines empty or all-zero object IDs as
  expected-absence markers. The evals require both "" and a 40-zero SHA-1 to
  succeed as `old` when the ref is absent, and to fail when it exists. The
  implementation normalizes "" and any all-zero string to absence (and zero
  `new` to deletion), with the zero-length derived from
  `git rev-parse --show-object-format`. Fixing change belongs in gitstore.
- GIT-003 (target: gitstore): The spec requires "once per reported Path" and
  that renames reported by either tracked diff retain both names, but does
  not define a merge rule for a path reported by both the staged and
  unstaged diffs (e.g. staged `M` plus worktree `M`, or staged `R` plus
  worktree `M`). The implementation merges by path, prefers entries
  carrying `OldPath` (rename names) and keeps the first status otherwise.
  Fixing change belongs in gitstore.
- GIT-007 (target: gitstore): The spec does not state behavior for an unborn
  HEAD (no parent commit), for an empty path selection, or for a commit
  whose tree is identical to its parent. The implementation supports unborn
  HEAD (no `-p`), rejects an empty path list, and creates the commit tree
  regardless of identical content. Fixing change belongs in gitstore.
- GIT-001 (target: gitstore): "Canonical absolute" directories are not
  defined with respect to symlinks. The implementation calls
  `filepath.EvalSymlinks` on the worktree root and the common Git dir
  (matters on macOS, where `/var` resolves under `/private`). Fixing change
  belongs in gitstore.
- GIT-002 (target: gitstore): The spec does not define "when known" for a
  dangling or unreadable `refs/remotes/<remote>/HEAD` symref (target ref
  missing). The implementation treats it as unknown and falls back to
  `main`. Fixing change belongs in gitstore.
- GIT-006 (target: gitstore): The spec does not constrain revision or path
  syntax; paths containing `:` cannot be addressed with `<rev>:<path>`
  plumbing. No observable eval conflict; fix would be wording in gitstore
  or a documented path filter.
- GIT-001 / GIT-002 / GIT-009 (target: gitstore): The module shells out to
  the `git` binary found on PATH for all Git operations and assumes 2.30+
  Apple Git compatibility (no `--path-format=absolute`; relative
  `--git-common-dir` results resolved against the invocation directory).
  The spec does not name the required git version or plumbing surface.
  Fixing change belongs in gitstore.

All gaps above carry their nearest spec ID (no truly spec-free gaps were
encountered) and the same fixing-change target, `gitstore`.

## Agent observations

### Invocation 1
Log: .fireweed/runs/20261118T071430-896ccac70870a7030cc46943/drill-1.log

All my git primitives check out. Now let me write the implementation.
Apple git 2.30 doesn't recognize `--path-format=absolute` — it echoes it verbatim. I'll resolve the common dir path manually instead.
Almost everything passes — one flaky case in `Test_GIT_005_RunCancelWhileShellLaunchesChild`. Let me reproduce it in isolation to understand the process-group behavior.
Let me stress-test the fixture path to understand the 5s case empirically.
Let me instrument the watcher temporarily to see the kill result and process group in failing iterations.
All kills report success, yet sometimes the child outlives by ~5s. Let me capture per-iteration process snapshots (pid, pgid, state) to see group membership at cancel time.
The ps instrumentation made it pass 40/40 — confirming the slow case is a cancel-timing race around when `sleep` is forked/exec'd. Let me capture the process state right at cancel in the racy window.
My pgrep pattern was likely a literal string on macOS (no ERE), making the samples meaningless. Let me verify and fix the sampling.
pgrep works with ERE. The empty samples mean nothing matching existed after cancel, yet Wait blocked ~5s. I need per-pid state sampling. Let me capture the killed pid from Run and probe it during the wait.
The killed pid is gone 30ms after cancel, output fully captured, yet Run waits ~5s. Something still holds the pipes. Let me check for a surviving `sleep` process specifically.
A `sleep 5` grandchild (pid 9761) survived the group kill. I need its pgid.
Theory keeps failing against data. Let me dump the entire process group at cancel+30ms for slow iterations.
The group is empty 30ms after cancel, output fully captured, yet Wait blocks to 5.0s. I need the group membership *at kill time*. Let me have the watcher itself snapshot the group right after the kill syscall.
The snap at kill time is empty — meaning bash's group contained nothing at kill time, yet `sleep 5` survives somewhere. Let me test what group bash puts `sleep` in.
Let me re-add the global `sleep 5` probe to see where the surviving sleep actually lives.
All processes are dead 30ms after cancel and output is fully captured, yet `Wait` blocks ~5s. Let me read Go's exec source to understand what `Wait` actually blocks on.
Let me add fine-grained timestamps inside Run to see exactly where it blocks.
Now the mechanism is clear: bash forks `sleep`, bash is SIGKILLed at t=3ms, but the just-forked `sleep` (pid 21449, pgid 21448) survives and orphaned to launchd (ppid 1), holding the inherited pipe write-ends — so Go's `awaitGoroutines` blocks ~5s on pipe EOF. `Cmd.WaitDelay` exists for exactly this (go.mod = go 1.23). Let me check its exact error semantics.
`cmd.WaitDelay` is the documented escape hatch for exactly this case (grandchildren holding inherited pipes — Go issue 23019). With `WaitDelay` set, `Wait` force-closes the parent pipes shortly after the child exits, the copiers unwind, and the `ExitError` stays the reported error. Since bash writes all marker output *before* spawning `sleep`, captured output stays complete. Let me remove the debug instrumentation and set it.
120/120 fast iterations, capture intact. Removing the debug test and re-running the official eval:
All evals pass and repeated runs are stable. Let me check the workspace state and gofmt before writing the notes.
Done. Summary:

**Rebuilt** (disposable paths only, `packages/gitstore/internal/local/`):
- `git.go` — internal git binary runner (`-C root`, hermetic env, stdin support, exit-code reporting).
- `local.go` — Inspect (canonical worktree/common-Git dirs via `EvalSymlinks`, HEAD, branch empty on detached HEAD, error for non-repos), BaseRef (remote ref preference → local branch fallback → error on missing; origin default; remote-HEAD branch default), Changes (staged `--cached` + unstaged `-M -z` diffs against base plus `?`-marked untracked paths, merged once per path, rename `OldPath` retained from either diff), ReadAt, WriteBlob (`hash-object -w --stdin`), ReadBlob, Ref, CASRef (exact-match CAS via `update-ref` with oldvalue pin; empty/all-zero `old` = expected absence; empty/all-zero `new` = delete; mismatch leaves ref untouched), Commit (temp `GIT_INDEX_FILE` from HEAD + `git add -A -- <paths>` + `write-tree` + `commit-tree` + `update-ref HEAD`, leaving index/worktree untouched).
- `run.go` — Run with complete argv (exec.Command, never prepends git), requested Dir, supplied environment (non-nil map replaces, nil inherits), stdin, captured streams, process-group launch via `Setpgid` + group SIGKILL on timeout/cancellation, errors wrapping `context.DeadlineExceeded`/`context.Canceled` appropriately, `*exec.ExitError` on nonzero exit (launch failures wrapped so they don't match), process errors exclude captured streams.

**Notable bug fixed during rebuild:** Apple Git 2.30 lacks `--path-format=absolute` (echoes it verbatim) — resolved by resolving `--git-common-dir` manually. Also, `Cmd.WaitDelay = 200ms` was required: a grandchild forked in the same instant as group SIGKILL can survive it while holding inherited pipes, which blocked `exec`'s `Wait` on pipe EOF for the full child lifetime (eval flake `RunCancelWhileShellLaunchesChild`); `WaitDelay` bounds that, and output capture stays exact since bash writes all markers before spawning the grandchild.

**Eval result:** `go run .fireweed/hooks.go eval packages/gitstore .fireweed/drill-eval.xml` → PASS (all 17 cases), stable across 4 consecutive runs. Report at `.fireweed/drill-eval.xml`.

**DRILL_NOTES.md** written with `## Contradicted Decisions` (none; decision 0010 consistent) and `## Gaps` (9 items: GIT-005 ×2, GIT-010, GIT-003, GIT-007, GIT-001 ×2, GIT-002, GIT-006, plus a host/git-version note; each names its nearest spec ID and targets `gitstore` for the fix).

Interface (`api.go`, `model/**`) and evals were not modified.


## Failures
app: failed case "Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI" for APP-062
Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip: === RUN   Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip
    ci_cli_test.go:206: default doctor CI: status=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"C11\",\"message\":\"read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-2613054885/report.xml: no such file or directory. Correct the reported issue and run fireweed ci again.\",\"path\":\"alpha\",\"warning\":false},{\"code\":\"X02\",\"message\":\"global boundary hook: exit status 127 without a module-attributed violation. Correct the reported issue and run fireweed ci again.\",\"warning\":false},{\"code\":\"C16\",\"message\":\"AGENTS.md managed block is missing or stale\",\"path\":\"AGENTS.md\",\"warning\":true},{\"code\":\"C14\",\"message\":\"setup setup: hook sh .fireweed/setup.sh: \\u003cnil\\u003e; exit 1;  . Correct the reported issue and run fireweed ci again.\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_111_APP_112_APP_113_CIDoctorCanariesNoAgentAndSkip (0.42s)
Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI: === RUN   Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI
    doctor_setup_test.go:125: APP-062/066 doctor status=1 stdout="FAIL baseline global eval doctor3204c34f12 eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-3048936237/report.xml: no such file or directory\nFAIL failing_eval global eval doctor3204c34f12 skipped after failed baseline\nFAIL missing_id global eval doctor3204c34f12 skipped after failed baseline\nFAIL baseline global boundaries doctor3204c34f12 boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation  false}]\nFAIL undeclared global boundaries doctor3204c34f12 skipped after failed baseline\nFAIL internal_eval global boundaries doctor36c7417442 skipped after failed baseline\n" stderr="C14: baseline global eval eval: eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-3048936237/report.xml: no such file or directory\nC14: failing_eval global eval eval: skipped after failed baseline\nC14: missing_id global eval eval: skipped after failed baseline\nC14: baseline global boundaries boundary hook: boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctor3204c34f12 boundary hook: exit status 127 without a module-attributed violation  false}]\nC14: undeclared global boundaries boundary hook: skipped after failed baseline\nC14: internal_eval global boundaries boundary hook: skipped after failed baseline\n"
--- FAIL: Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI (0.82s)
Test_APP_062_DoctorUsesCloserProject: === RUN   Test_APP_062_DoctorUsesCloserProject
    doctor_setup_test.go:245: APP-062 closer doctor status=1 output="FAIL baseline global eval doctorc5201995ac eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-2631893846/report.xml: no such file or directory\nFAIL failing_eval global eval doctorc5201995ac skipped after failed baseline\nFAIL missing_id global eval doctorc5201995ac skipped after failed baseline\nFAIL baseline global boundaries doctorc5201995ac boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctorc5201995ac boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctorc5201995ac boundary hook: exit status 127 without a module-attributed violation  false}]\nFAIL undeclared global boundaries doctorc5201995ac skipped after failed baseline\nFAIL internal_eval global boundaries doctor56c0df8ca5 skipped after failed baseline\nC14: baseline global eval eval: eval: read JUnit report: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-junit-2631893846/report.xml: no such file or directory\nC14: failing_eval global eval eval: skipped after failed baseline\nC14: missing_id global eval eval: skipped after failed baseline\nC14: baseline global boundaries boundary hook: boundaries: global boundary hook: exit status 127 without a module-attributed violation\ndoctorc5201995ac boundary hook: exit status 127 without a module-attributed violation: [{X02 global boundary hook: exit status 127 without a module-attributed violation  false} {X02 doctorc5201995ac boundary hook: exit status 127 without a module-attributed violation  false}]\nC14: undeclared global boundaries boundary hook: skipped after failed baseline\nC14: internal_eval global boundaries boundary hook: skipped after failed baseline\n"
--- FAIL: Test_APP_062_DoctorUsesCloserProject (0.89s)
Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange: === RUN   Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange
    next_cli_test.go:172: next exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \"ALP-002\" for ALP-002. Correct the reported problem and use fireweed next to resume.\n"
--- FAIL: Test_APP_074_APP_076_APP_077_NextDiscoversRootAndAppliesChange (2.44s)
Test_APP_075_NextAllowUnmergedAndUsage: === RUN   Test_APP_075_NextAllowUnmergedAndUsage
    next_cli_test.go:247: allowed next exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \"ALP-002\" for ALP-002. Correct the reported problem and use fireweed next to resume.\nP04: warning: changes/0001-next.md: change is absent from the local and shared base\n"
--- FAIL: Test_APP_075_NextAllowUnmergedAndUsage (2.62s)
Test_APP_079_NextJSONSuccessAndFailure: === RUN   Test_APP_079_NextJSONSuccessAndFailure
    next_cli_test.go:328: JSON applied exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \\\"ALP-002\\\" for ALP-002. Correct the reported problem and use fireweed next to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_079_NextJSONSuccessAndFailure (2.60s)
Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation: === RUN   Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation
    run_cli_test.go:262: run exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_080_APP_081_APP_083_APP_084_RunBoundedContinuation (3.39s)
Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON: === RUN   Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON
    run_cli_test.go:323: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_081_APP_084_APP_086_RunUnlimitedAndJSON (3.43s)
Test_APP_083_RunPlainMultiple: === RUN   Test_APP_083_RunPlainMultiple
    run_cli_test.go:347: run exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_083_RunPlainMultiple (3.38s)
Test_APP_086_RunJSONLimited: === RUN   Test_APP_086_RunJSONLimited
    run_cli_test.go:359: run exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_086_RunJSONLimited (3.37s)
Test_APP_082_RunAllowUnmerged: === RUN   Test_APP_082_RunAllowUnmerged
    run_cli_test.go:376: run exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002. Correct the reported problem and use fireweed run to resume.\n"
--- FAIL: Test_APP_082_RunAllowUnmerged (2.63s)
Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes: === RUN   Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes
    run_for_cli_test.go:95: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
--- FAIL: Test_APP_089_APP_090_APP_084_RunForJSONRetainsAppliedAndResumes (3.53s)
Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume: === RUN   Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume
    usage_cli_test.go:168: owned resume exit=1 stderr="" stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case \\\"ALP-002\\\" for ALP-002. Correct the reported problem and use fireweed next to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n"
    usage_cli_test.go:176: resumed next data={Change: Run: Phase: Empty:false}: <nil>
    usage_cli_test.go:178: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume1248919961/002/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_APP_076_APP_077_APP_079_NextJSONUsageStopAndOwnedResume (2.68s)
Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume: === RUN   Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume
    usage_cli_test.go:252: run stop exit=1 stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"E01\",\"message\":\"impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \\\"ALP-002\\\" for ALP-002\\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\",\"path\":\"ALP-002\",\"warning\":false}]}\n" stderr=""
    usage_cli_test.go:254: clean stop data=null; want object: <nil>
    usage_cli_test.go:270: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume2831879836/002/changes/archive/0001-alpha.md: no such file or directory
--- FAIL: Test_APP_083_APP_084_APP_086_RunJSONSharedUsageStopAndResume (3.19s)
Test_APP_083_APP_084_RunPlainSharedUsageStop: === RUN   Test_APP_083_APP_084_RunPlainSharedUsageStop
    usage_cli_test.go:331: plain run stop exit=1 stdout="" stderr="E01: ALP-002: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case \"ALP-002\" for ALP-002\nevals verification exhausted (saved attempt refs/fireweed/attempts/beta/1): foreign module: packages/alpha/internal/value.go. Correct the reported problem and use fireweed run to resume.\n"
    usage_cli_test.go:333: plain run lacks alpha: ""
    usage_cli_test.go:333: plain run lacks applied count 1: ""
    usage_cli_test.go:335: plain run lacks resume guidance: ""
    usage_cli_test.go:337: stop reason "" lacks a labeled completed wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled next wait of 5ms
    usage_cli_test.go:337: stop reason "" lacks a labeled budget of 5ms
    usage_cli_test.go:339: alpha not applied: stat /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_083_APP_084_RunPlainSharedUsageStop2262373806/002/changes/archive/0001-alpha.md: no such file or directory
    usage_cli_test.go:341: held beta state: change has no claim: beta
--- FAIL: Test_APP_083_APP_084_RunPlainSharedUsageStop (3.26s)
?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	163.537s

exit status 1

checks: failed case "Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe" for CHK-050
Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary: === RUN   Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary
--- FAIL: Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary (2.11s)
Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test: === RUN   Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test
    doctor_missing_id_test.go:66: Doctor missed the supplied Go canary: checks=[{Name:setup Module: Passed:false Detail:hook go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_missing_id_test.go:76: Doctor did not report passing missing_id recovery: checks=[{Name:setup Module: Passed:false Detail:hook go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
--- FAIL: Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test (0.50s)
Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe: === RUN   Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe
    doctor_test.go:345: real Go recipe doctor failed: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing baseline outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing undeclared outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing internal_eval outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing failing_eval outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing missing_id outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:349: missing passing recovery outcome: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:359: clean recovery not reported after each canary: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
    doctor_test.go:368: two distinct throwaway modules absent: checks=[{Name:setup Module: Passed:false Detail:hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
        }] diagnostics=[{Code:C14 Message:setup setup: hook printf 'setup:%s\n' "$PWD" >> "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK73629533/002/probes.log"; test -z "$(git symbolic-ref -q HEAD)" && test -f .fireweed/doctor-untracked && go mod download: <nil>; exit 127;  sh: go: command not found
         Path: Warning:false}]
--- FAIL: Test_CHK_050_CHK_051_CHK_052_CHK_053_CHK_055_CHK_056_CHK_057_CHK_058_CHK_059_CHK_060_CHK_061_CHK_067_RealGoRecipe (0.55s)
?   	github.com/lurkshark/fireweed/packages/checks	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/internal/engine	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/checks/evals	243.049s

exit status 1

runner: failed case "Test_RUN_010_EntirelyWaitingQueue" for RUN-010
Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput: === RUN   Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput
    agent_logs_test.go:109: Next fixture = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074216-5e949a32f53bedf7ce71ffd8"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:16.780134Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:17.15557Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074216-5e949a32f53bedf7ce71ffd8", Phase:"impl", Head:"d920b7f6089b1f8e590e1ecb79fb113967820cef", PhaseBase:"d920b7f6089b1f8e590e1ecb79fb113967820cef", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAudit3707438102/001/.fireweed/runs/20261118T074216-5e949a32f53bedf7ce71ffd8/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_064_NextLogsHeldRunPhaseAndOneBasedAttemptsOutsideAuditInput (1.99s)
Test_RUN_021_AuditCorrectionWithinAllowance: === RUN   Test_RUN_021_AuditCorrectionWithinAllowance
--- FAIL: Test_RUN_021_AuditCorrectionWithinAllowance (6.39s)
Test_RUN_021_AuditCorrectionWithinAllowance/correction: === RUN   Test_RUN_021_AuditCorrectionWithinAllowance/correction
    audit_return_attempts_test.go:123: audit correction ended in "impl": impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_AuditCorrectionWithinAllowance/correction (3.17s)
Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry: === RUN   Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry
    audit_return_attempts_test.go:123: audit correction ended in "impl": impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_AuditCorrectionWithinAllowance/usage_retry (3.22s)
Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals: === RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals (5.46s)
Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry: === RUN   Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry
    audit_return_attempts_test.go:154: phases captured=[evals audit evals audit] invoked=[evals audit evals audit], want [evals evals audit]
    audit_return_attempts_test.go:166: phase commits before block = 4, want 2
    audit_return_attempts_test.go:187: phase base did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T074230-23913473b684f9a0e3d7803d\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:30.982366Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:31.363375Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T07:42:31.962282Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:32.342388Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2"
    audit_return_attempts_test.go:189: blocked change did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nblocked: evals attempt allowance exhausted after 2 attempts\nmodule: alpha\nphase: evals\nrun: 20261118T074230-23913473b684f9a0e3d7803d\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:30.982366Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:31.363375Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T07:42:31.962282Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:32.342388Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2\n\n"
    audit_return_attempts_test.go:194: returned state lost the latest audit finding: "\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2\n\n"
    audit_return_attempts_test.go:197: saved attempt did not retain the held evals identity and latest audit finding: "---\naudit_rounds: 2\nbases: {}\nmodule: alpha\nphase: evals\nrun: 20261118T074230-23913473b684f9a0e3d7803d\n---\n\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\nAlpha MUST return the new value on subsequent calls.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:30.982366Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:31.363375Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n- evals 2026-11-18T07:42:31.962282Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:32.342388Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\n- ALP-002: \"Alpha MUST return the new value on subsequent calls.\" — missing an assertion for subsequent calls in audit 2"
--- FAIL: Test_RUN_021_RUN_070_AuditReturnExhaustionSavesAndBlocksEvals/earlier_verification_retry (2.64s)
Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit: === RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit (4.88s)
Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override: === RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074239-d8e068a30b313b955a9fdbd0"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:40.239572Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:40.652163Z agent=bash fireweed=0.2.1 playbook=a1eef0b302afb0055dd630de8b87b5ff467a00f40d16783ba79821ee228a5cf6 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074239-d8e068a30b313b955a9fdbd0", Phase:"impl", Head:"5f46ed8bcd504cd9ba5d37d4f360d923d60d5caf", PhaseBase:"5f46ed8bcd504cd9ba5d37d4f360d923d60d5caf", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2730508749/001/.fireweed/runs/20261118T074239-d8e068a30b313b955a9fdbd0/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/project_override (2.45s)
Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback: === RUN   Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback
    audit_workspace_test.go:78: isolated audit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074242-b5609f79c0a71407e0d82918"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:42.660238Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:43.075258Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074242-b5609f79c0a71407e0d82918", Phase:"impl", Head:"6478a2461bec24bccb26d8a9d9100ba6ea6da968", PhaseBase:"6478a2461bec24bccb26d8a9d9100ba6ea6da968", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInput2776841683/001/.fireweed/runs/20261118T074242-b5609f79c0a71407e0d82918/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_013_RUN_045_RUN_046_RUN_071_AuditUsesOnlyProjectedInputsAndIndependentGit/embedded_fallback (2.42s)
Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry: === RUN   Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry
    audit_workspace_test.go:210: audit verification retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074245-41fd0f1873bdfaa0915f8e2c"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:46.196144Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:46.735827Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074245-41fd0f1873bdfaa0915f8e2c", Phase:"impl", Head:"6f938f2978c3826be5c55c4973d5de9f9a9b25f3", PhaseBase:"6f938f2978c3826be5c55c4973d5de9f9a9b25f3", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/audit-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/audit-2.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificati1906816077/001/.fireweed/runs/20261118T074245-41fd0f1873bdfaa0915f8e2c/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_048_PermittedAuditEditReachesHeldChangeBeforeVerificationRetry (2.54s)
Test_RUN_046_AuditCommitDoesNotEnterHeldGit: === RUN   Test_RUN_046_AuditCommitDoesNotEnterHeldGit
    audit_workspace_test.go:232: audit commit lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074248-ef9e8b131e0a5c24c66480fe"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:42:48.892307Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:49.317526Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074248-ef9e8b131e0a5c24c66480fe", Phase:"impl", Head:"5bd546b8147a182190eb3eb4b26c14b4f5815719", PhaseBase:"5bd546b8147a182190eb3eb4b26c14b4f5815719", ClaimOID:"", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_046_AuditCommitDoesNotEnterHeldGit1237886595/001/.fireweed/runs/20261118T074248-ef9e8b131e0a5c24c66480fe/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_046_AuditCommitDoesNotEnterHeldGit (2.69s)
Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported: === RUN   Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported
    audit_workspace_test.go:280: audit before Provenance lifecycle = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074250-a0af3dd5b3ca84e1a25bf268"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Audit\n\nNo findings; checked ALP-002\n\n## Provenance\n\n- evals 2026-11-18T07:42:51.422456Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:42:51.980842Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074250-a0af3dd5b3ca84e1a25bf268", Phase:"impl", Head:"0c69be8c3798128dad2239b8e631094df6feea50", PhaseBase:"0c69be8c3798128dad2239b8e631094df6feea50", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001/.fireweed/runs/20261118T074250-a0af3dd5b3ca84e1a25bf268/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001/.fireweed/runs/20261118T074250-a0af3dd5b3ca84e1a25bf268/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported1796399985/001/.fireweed/runs/20261118T074250-a0af3dd5b3ca84e1a25bf268/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_047_RUN_048_FirstAuditBeforeProvenanceIsImported (2.55s)
Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle: === RUN   Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle
    commit_recovery_test.go:38: Next after fixture agent commits: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002; state=model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074310-1276a3290904e5d5d6a6a845"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:43:11.182746Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:43:11.596746Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002", "commit-success", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074310-1276a3290904e5d5d6a6a845", Phase:"impl", Head:"e37f814e581328035e80bd367856bbd06badd3a7", PhaseBase:"e37f814e581328035e80bd367856bbd06badd3a7", ClaimOID:"", Diagnostics:[]model.Diagnostic{model.Diagnostic{Code:"P05", Message:"soft reset evals worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}, model.Diagnostic{Code:"P05", Message:"soft reset impl worker commits to the phase base and retained the work", Path:"changes/0001-next.md", Warning:true}}}; steps=[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle655327431/001/.fireweed/runs/20261118T074310-1276a3290904e5d5d6a6a845/impl-2.log"}}}
--- FAIL: Test_RUN_027_RUN_028_RecoversAgentCommitsAcrossLifecycle (2.66s)
Test_RUN_011_WaitingHeadSelectsIndependent: === RUN   Test_RUN_011_WaitingHeadSelectsIndependent
    next_test.go:428: Next blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002")
    next_test.go:431: fixture alpha did not block: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"alpha", Path:"changes/0001-alpha.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"", "module":"alpha", "phase":"impl", "run":"20261118T074404-02d870a893c1d79ea5ba4251"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:04.99338Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:05.528544Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/002", "exhaust-alpha", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# alpha\n\n## Specs\n\n### ALP-001\nalpha MUST return a value.\n\n### ALP-002\nalpha MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# beta\n\n## Specs\n\n### BET-001\nbeta MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"GAM", "name":"gamma"}, Body:"# gamma\n\n## Specs\n\n### GAM-001\ngamma MUST return a value.\n\n## Decisions\n"}}}}, Run:"20261118T074404-02d870a893c1d79ea5ba4251", Phase:"impl", Head:"d1c34b42cc6670a148d16b2a46c39e8f7ff4f372", PhaseBase:"d1c34b42cc6670a148d16b2a46c39e8f7ff4f372", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-no-edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"alpha-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_011_WaitingHeadSelectsIndependent1075997610/001/.fireweed/runs/20261118T074404-02d870a893c1d79ea5ba4251/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
--- FAIL: Test_RUN_011_WaitingHeadSelectsIndependent (2.33s)
Test_RUN_010_EntirelyWaitingQueue: === RUN   Test_RUN_010_EntirelyWaitingQueue
    next_test.go:449: Next blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002")
    next_test.go:452: fixture alpha did not block: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"alpha", Path:"changes/0001-alpha.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"", "module":"alpha", "phase":"impl", "run":"20261118T074406-f95fc7385dd40d14e6c424bf"}, Body:"\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:07.334363Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:07.854304Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/002", "exhaust-alpha", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# alpha\n\n## Specs\n\n### ALP-001\nalpha MUST return a value.\n\n### ALP-002\nalpha MUST return the new value.\n\n## Decisions\n"}}, model.Module{Name:"beta", IDPrefix:"BET", Dir:"packages/beta", DependsOn:[]string{"alpha"}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{"alpha"}, "id_prefix":"BET", "name":"beta"}, Body:"# beta\n\n## Specs\n\n### BET-001\nbeta MUST return a value.\n\n## Decisions\n"}}, model.Module{Name:"gamma", IDPrefix:"GAM", Dir:"packages/gamma", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"GAM", "name":"gamma"}, Body:"# gamma\n\n## Specs\n\n### GAM-001\ngamma MUST return a value.\n\n## Decisions\n"}}}}, Run:"20261118T074406-f95fc7385dd40d14e6c424bf", Phase:"impl", Head:"86a0d0046636bf39356a94f4f29894fb36aa11b6", PhaseBase:"86a0d0046636bf39356a94f4f29894fb36aa11b6", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-no-edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"alpha-audit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"alpha-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_010_EntirelyWaitingQueue2938442048/001/.fireweed/runs/20261118T074406-f95fc7385dd40d14e6c424bf/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
--- FAIL: Test_RUN_010_EntirelyWaitingQueue (2.30s)
Test_RUN_013_RUN_014_CompleteSpecChange: === RUN   Test_RUN_013_RUN_014_CompleteSpecChange
    next_test.go:478: Next: model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074409-e109af499ffad963aa5ac6ae"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:09.790834Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:10.163755Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074409-e109af499ffad963aa5ac6ae", Phase:"impl", Head:"fbbb7c9cd3bd239785d2f0b034a68abf992b1e58", PhaseBase:"fbbb7c9cd3bd239785d2f0b034a68abf992b1e58", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_013_RUN_014_CompleteSpecChange1185415938/001/.fireweed/runs/20261118T074409-e109af499ffad963aa5ac6ae/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_013_RUN_014_CompleteSpecChange (2.39s)
Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks: === RUN   Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks
    next_test.go:678: Next error code = "E01" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002)
    next_test.go:681: evals exhaustion cause = impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
    next_test.go:684: evals exhaustion steps = 4, want 2
    next_test.go:688: captured invocation 1 = model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/001/.fireweed/runs/20261118T074412-10877f23019fb888d8c42713/audit-1.log"}}
    next_test.go:688: captured invocation 2 = model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"no edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/001/.fireweed/runs/20261118T074412-10877f23019fb888d8c42713/impl-1.log"}}
    next_test.go:688: captured invocation 3 = model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"no edit", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/001/.fireweed/runs/20261118T074412-10877f23019fb888d8c42713/impl-2.log"}}
    next_test.go:692: agent calls = 4, want 2
    next_test.go:694: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks3950714332/002/evals-final-head: no such file or directory
--- FAIL: Test_RUN_023_RUN_024_DefaultEvalsExhaustionBlocks (1.79s)
Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase: === RUN   Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase
    next_test.go:706: Next after retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074413-e80248025c074b7568278683"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:14.307437Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:14.686686Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002", "retry-both", "{prompt}"}}, "default":"fixture-worker"}, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074413-e80248025c074b7568278683", Phase:"impl", Head:"e9d5dadb411b19ea59287a95fbb817086961e416", PhaseBase:"e9d5dadb411b19ea59287a95fbb817086961e416", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase2175373443/001/.fireweed/runs/20261118T074413-e80248025c074b7568278683/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_RUN_022_DefaultRetriesRetainWorkAndFeedbackPerPhase (1.99s)
Test_RUN_021_ConfiguredThreeAttempts: === RUN   Test_RUN_021_ConfiguredThreeAttempts
    next_test.go:756: Next with three attempts = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 3 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074415-d3bc668f21f3067c2f7bd67d"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:16.523918Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:16.905793Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002", "retry-third", "{prompt}"}}, "default":"fixture-worker"}, "attempts":3, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/002"}}, "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074415-d3bc668f21f3067c2f7bd67d", Phase:"impl", Head:"3a5205442a6486d14f3f251fb8b8413fb4707938", PhaseBase:"3a5205442a6486d14f3f251fb8b8413fb4707938", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-1", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl-2", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/impl-2.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_021_ConfiguredThreeAttempts2795209233/001/.fireweed/runs/20261118T074415-d3bc668f21f3067c2f7bd67d/impl-3.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_021_ConfiguredThreeAttempts (2.34s)
Test_RUN_016_UnlimitedQueueOrder: === RUN   Test_RUN_016_UnlimitedQueueOrder
    run_test.go:313: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/beta/1): failed case "BET-002" for BET-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002
--- FAIL: Test_RUN_016_UnlimitedQueueOrder (7.12s)
Test_RUN_016_NegativeMaxIsUnlimited: === RUN   Test_RUN_016_NegativeMaxIsUnlimited
    run_test.go:329: Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
--- FAIL: Test_RUN_016_NegativeMaxIsUnlimited (2.49s)
Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding: === RUN   Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding
    run_test.go:341: first Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/beta/1): failed case "BET-002" for BET-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002
--- FAIL: Test_RUN_017_RUN_018_RUN_019_BoundedContinuationAndForwarding (7.14s)
Test_RUN_018_InitialEmptyAndExactLimit: === RUN   Test_RUN_018_InitialEmptyAndExactLimit
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit (4.96s)
Test_RUN_018_InitialEmptyAndExactLimit/exact_limit: === RUN   Test_RUN_018_InitialEmptyAndExactLimit/exact_limit
    run_test.go:385: exact Run: model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        impl verification exhausted (saved attempt refs/fireweed/attempts/beta/1): failed case "BET-002" for BET-002
--- FAIL: Test_RUN_018_InitialEmptyAndExactLimit/exact_limit (4.78s)
Test_RUN_020_HeldFailureRetainsEarlierResults: === RUN   Test_RUN_020_HeldFailureRetainsEarlierResults
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_020_HeldFailureRetainsEarlierResults (2.79s)
Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent: === RUN   Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent
    run_test.go:427: Run blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002)
    run_test.go:427: Run blocking error = impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002, want verification cause and refs/fireweed/attempts/alpha/1
    run_test.go:427: alpha block state = "---\nbases: {}\nblocked: 'impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"'\nmodule: alpha\nphase: impl\nrun: 20261118T074451-e91a7f0322043a6ec6a63f9d\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:52.217632Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:52.750661Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_016_RUN_025_RUN_026_BlockThenApplyIndependent (4.74s)
Test_RUN_025_RUN_026_PositiveMaxRetainsBlock: === RUN   Test_RUN_025_RUN_026_PositiveMaxRetainsBlock
    run_test.go:448: Run blocking error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002)
    run_test.go:448: Run blocking error = impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        impl verification exhausted (saved attempt refs/fireweed/attempts/gamma/1): failed case "GAM-002" for GAM-002, want verification cause and refs/fireweed/attempts/alpha/1
    run_test.go:448: alpha block state = "---\nbases: {}\nblocked: 'impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"'\nmodule: alpha\nphase: impl\nrun: 20261118T074456-4627920bb90d5fe811e6e4e3\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:44:56.99838Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:44:57.529869Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"
    run_test.go:187: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_025_RUN_026_PositiveMaxRetainsBlock (4.78s)
Test_RUN_072_RetainsStartWarning: === RUN   Test_RUN_072_RetainsStartWarning
--- FAIL: Test_RUN_072_RetainsStartWarning (10.71s)
Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery: === RUN   Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery
    start_warnings_test.go:81: Next: impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_072_RetainsStartWarning/applied_with_worker_recovery (2.85s)
Test_RUN_072_RetainsStartWarning/Run_applied_result: === RUN   Test_RUN_072_RetainsStartWarning/Run_applied_result
    start_warnings_test.go:147: expected Run to apply one queued change, got 0 applied, empty=false, limited=false, err=impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_072_RetainsStartWarning/Run_applied_result (2.46s)
Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep: === RUN   Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep
    usage_test.go:206: owned resume = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 1 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074636-2705e1c8f4aaa1254bb3c9c3"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:46:36.859721Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:46:37.233548Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":1, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/002"}}, "limit_budget":"5ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074636-2705e1c8f4aaa1254bb3c9c3", Phase:"impl", Head:"69bc14f4c1fb6f5ed776cd1250621821f266693f", PhaseBase:"69bc14f4c1fb6f5ed776cd1250621821f266693f", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"worker-evals", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/.fireweed/runs/20261118T074636-2705e1c8f4aaa1254bb3c9c3/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/.fireweed/runs/20261118T074636-2705e1c8f4aaa1254bb3c9c3/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/.fireweed/runs/20261118T074636-2705e1c8f4aaa1254bb3c9c3/impl-1.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
    usage_test.go:208: open /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep1570118501/001/changes/archive/0001-next.md: no such file or directory
--- FAIL: Test_RUN_030_RUN_031_RUN_033_ShortBudgetResumeOwnedStep (2.39s)
Test_RUN_030_UsageRetryPreservesVerificationFeedback: === RUN   Test_RUN_030_UsageRetryPreservesVerificationFeedback
    usage_test.go:288: verification and usage retry = model.NextResult{State:model.State{Change:model.Change{Number:1, Slug:"next", Path:"changes/0001-next.md", Module:"alpha", Document:model.Document{Frontmatter:map[string]interface {}{"bases":map[string]interface {}{}, "blocked":"impl verification failed after 2 attempts: failed case \"ALP-002_fixture_value\" for ALP-002", "module":"alpha", "phase":"impl", "run":"20261118T074639-26311bead8c44ed47d206aca"}, Body:"\n## Motivation\nOne new value is needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:46:39.649858Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:46:40.023355Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"}, Added:[]model.Item{model.Item{ID:"ALP-002", Tag:"", Body:"Alpha MUST return the new value.\n\n"}}, Modified:[]model.Item(nil), Removed:[]model.Item(nil)}, Snapshot:model.Snapshot{Root:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001", Config:map[string]interface {}{"agent":map[string]interface {}{"commands":map[string]interface {}{"fixture-auditor":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002/auditor.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002", "{prompt}"}, "fixture-worker":[]interface {}{"bash", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002/worker.sh", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002", "normal", "{prompt}"}}, "default":"fixture-worker"}, "attempts":2, "audit":map[string]interface {}{"agent":"fixture-auditor"}, "base":"main", "fireweed":"=0.2.1", "hooks":map[string]interface {}{"boundaries":[]interface {}{"sh", ".fireweed/boundaries.sh"}, "eval":[]interface {}{"sh", ".fireweed/eval.sh", "{module}", "{report}", "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/002"}}, "limit_budget":"20ms", "limit_patterns":[]interface {}{"RATE LIMIT"}, "limit_wait":"5ms", "module_defaults":map[string]interface {}{"disposable":[]interface {}{"internal/**"}, "evals":[]interface {}{"evals/**"}, "interface":[]interface {}{"api.go"}}, "modules":[]interface {}{"packages/*"}, "spec_limits":map[string]interface {}{"max_clauses":3, "max_decision_words":300, "max_words":80}, "timeout":"50ms", "version":1}, Modules:[]model.Module{model.Module{Name:"alpha", IDPrefix:"ALP", Dir:"packages/alpha", DependsOn:[]string{}, Zones:model.Zones{Protected:[]string(nil), Evals:[]string{"evals/**"}, Interface:[]string{"api.go"}, AppendOnly:[]string(nil), Preserve:[]string(nil), Disposable:[]string{"internal/**"}}, Document:model.Document{Frontmatter:map[string]interface {}{"depends_on":[]interface {}{}, "id_prefix":"ALP", "name":"alpha"}, Body:"# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST return its value.\n\n### ALP-002\nAlpha MUST return the new value.\n\n## Decisions\n"}}}}, Run:"20261118T074639-26311bead8c44ed47d206aca", Phase:"impl", Head:"93e110f016f309f23e1ec864747796a5b74cd431", PhaseBase:"93e110f016f309f23e1ec864747796a5b74cd431", ClaimOID:"", Diagnostics:[]model.Diagnostic(nil)}, Empty:false, Steps:[]model.StepResult{model.StepResult{Phase:"evals", Role:"worker", Agent:model.AgentResult{Stdout:"first verification", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/evals-1.log"}}, model.StepResult{Phase:"audit", Role:"audit", Agent:model.AgentResult{Stdout:"auditor", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/audit-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/impl-1.log"}}, model.StepResult{Phase:"impl", Role:"worker", Agent:model.AgentResult{Stdout:"worker-impl", Stderr:"", ExitCode:0, UsageLimit:false, LogPath:"/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_RUN_030_UsageRetryPreservesVerificationFeedback2129357536/001/.fireweed/runs/20261118T074639-26311bead8c44ed47d206aca/impl-2.log"}}}}, impl verification exhausted (saved attempt refs/fireweed/attempts/next/1): failed case "ALP-002_fixture_value" for ALP-002
--- FAIL: Test_RUN_030_UsageRetryPreservesVerificationFeedback (2.18s)
Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied: === RUN   Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied
    usage_test.go:329: Run stop = impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): failed case "ALP-002" for ALP-002
        usage limit stopped execution: completed wait 5ms, next wait 5ms, budget 5ms
    usage_test.go:331: Applied count = 0, want 1: []model.NextResult(nil)
--- FAIL: Test_RUN_032_RunSharesWaitAcrossChangesAndRetainsApplied (2.85s)
Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure: === RUN   Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure
    usage_test.go:371: combined block and usage error code = "C11" (present: true), want P2-01 (error: impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        usage limit stopped execution: completed wait 0s, next wait 5ms, budget 0s)
    usage_test.go:374: combined block and usage failure = model.RunResult{Applied:[]model.NextResult(nil), Empty:false, Limited:false}, impl verification exhausted (saved attempt refs/fireweed/attempts/alpha/1): testcase "BET-002" references unknown spec ID "BET-002"
        usage limit stopped execution: completed wait 0s, next wait 5ms, budget 0s
    usage_test.go:380: alpha block state = "---\nbases: {}\nblocked: 'impl verification failed after 2 attempts: testcase \"BET-002\" references unknown spec ID \"BET-002\"'\nmodule: alpha\nphase: impl\nrun: 20261118T074646-a18b30c12b89b65703d935d5\n---\n\n## Motivation\nNew value needed.\n\n## Interface\nExisting Value function.\n\n## Added\n\n### ALP-002\nalpha MUST return the new value.\n\n## Modified\n\n## Removed\n\n## Decisions\n## Provenance\n\n- evals 2026-11-18T07:46:46.805926Z agent=bash fireweed=0.2.1 playbook=d2fd7216b035a9614c390d254dd9195ac2c65ae1ed18f07ac93c77b9732d4156 tests=2 failing=1 passing=1\n- audit 2026-11-18T07:46:47.331203Z agent=bash fireweed=0.2.1 playbook=af82dfbbb21a41a13c693f48de95ff7540e3b1a20f18c8f6a1a65e0106b7dc96 tests=0 failing=0 passing=0\n\n## Audit\n\nNo findings; checked ALP-002\n\n"
--- FAIL: Test_RUN_035_EarlierBlockMakesLaterUsageStopFailure (2.81s)
?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/policy	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/workspace	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/runner/evals	275.627s

exit status 1

