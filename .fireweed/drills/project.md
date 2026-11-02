---
date: 2026-11-02T00:01:16.03212Z
base_commit: fed193a3f8d2c86dab33aef69972543e41d45e9f
agent: codex
model: gpt-6-sol
result: fail
---

## Own evals
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
- project | Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder | PRJ-060 | pass
- project | Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash | PRJ-061, PRJ-062 | pass
- project | Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem | PRJ-062 | pass
- project | Test_PRJ_063_StaleSpecsEmptyAndReadOnly | PRJ-063 | pass
- project | Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent | PRJ-064 | pass
- project | Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically | PRJ-065 | pass
- project | Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically/nonmapping | PRJ-065 | pass
- project | Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically/absent_modified_after_valid | PRJ-065 | pass
- project | Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically/absent_removed_after_valid | PRJ-065 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/target | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/0001-target.md | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/changes/0001-target.md | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StaleSpecs_unknown | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StampMissing_unknown | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StaleSpecs_ambiguous | PRJ-066 | pass
- project | Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures/StampMissing_ambiguous | PRJ-066 | pass
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
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.165s


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
- app | Test_APP_107_APP_108_CheckCommitsFailures | APP-107, APP-108 | fail
  Failure: === RUN   Test_APP_107_APP_108_CheckCommitsFailures
    check_commits_cli_test.go:269: coded operational fixture: want malformed recorded SPEC, got C03 packages/alpha/SPEC.md: malformed recorded SPEC: C03 packages/alpha/SPEC.md: invalid YAML frontmatter: yaml: line 1: did not find expected ',' or ']'
--- FAIL: Test_APP_107_APP_108_CheckCommitsFailures (0.35s)
- app | Test_APP_107_APP_108_CheckCommitsFailures/extra | APP-107, APP-108 | pass
- app | Test_APP_107_APP_108_CheckCommitsFailures/option | APP-107, APP-108 | pass
- app | Test_APP_107_APP_108_CheckCommitsFailures/bad-range | APP-107, APP-108 | pass
- app | Test_APP_147_APP_148_CheckFixAcrossQueuePreservesRecordedBases | APP-147, APP-148 | pass
- app | Test_APP_147_APP_148_CheckFixMissingOnlySuccessAndNoOp | APP-147, APP-148 | pass
- app | Test_APP_148_APP_149_CheckFixRejectsBeforeMutation | APP-148, APP-149 | pass
- app | Test_APP_109_APP_110_APP_112_APP_113_CIHistoryAndAllModules | APP-109, APP-110, APP-112, APP-113 | pass
- app | Test_APP_110_APP_112_APP_113_CIEvalAndBoundaryFailures | APP-110, APP-112, APP-113 | pass
- app | Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration | APP-109, APP-112, APP-113 | fail
  Failure: === RUN   Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration
    ci_cli_test.go:190: CI C05 diagnostic lacks corrective action: "C05 packages/beta/SPEC.md: duplicate item ID: BET-001; remove or rename the conflicting spec ID. Correct fireweed.yaml and module declarations, then retry ci."
--- FAIL: Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration (0.26s)
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
- app | Test_APP_038_LookupCauses | APP-038 | fail
  Failure: === RUN   Test_APP_038_LookupCauses
--- FAIL: Test_APP_038_LookupCauses (0.55s)
- app | Test_APP_038_LookupCauses/author-module | APP-038 | fail
  Failure: === RUN   Test_APP_038_LookupCauses/author-module
    diagnostic_causes_test.go:156: missing U02 cause "unknown module" in plain output "U02: U02 unknown: module unknown not found. Use a valid module and an unused lowercase slug.\n"
    diagnostic_causes_test.go:156: missing U02 cause "unknown module" in {"schema":1,"data":null,"diagnostics":[{"code":"U02","message":"U02 unknown: module unknown not found. Use a valid module and an unused lowercase slug.","warning":false}]}
--- FAIL: Test_APP_038_LookupCauses/author-module (0.25s)
- app | Test_APP_038_LookupCauses/stamp-change | APP-038 | pass
- app | Test_APP_038_LookupCauses/show-module | APP-038 | pass
- app | Test_APP_038_LookupCauses/show-item | APP-038 | pass
- app | Test_APP_038_LookupCauses/show-ambiguous-item | APP-038 | pass
- app | Test_APP_038_LookupCauses/coverage-module | APP-038 | pass
- app | Test_APP_038_LookupCauses/coverage-change | APP-038 | pass
- app | Test_APP_039_BlockedNewChangeOutputPath | APP-039 | fail
  Failure: === RUN   Test_APP_039_BlockedNewChangeOutputPath
    diagnostic_causes_test.go:183: lost create-file cause or correction: "P01: P01 changes/0003-blocked.md: create change file changes/0003-blocked.md: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_039_BlockedNewChangeOutputPath2663529851/002/changes/0003-blocked.md: file exists. Use a valid module and an unused lowercase slug.\n"
    diagnostic_causes_test.go:183: lost create-file cause or correction: "P01 changes/0003-blocked.md: create change file changes/0003-blocked.md: open /var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_APP_039_BlockedNewChangeOutputPath2663529851/003/changes/0003-blocked.md: file exists. Use a valid module and an unused lowercase slug."
--- FAIL: Test_APP_039_BlockedNewChangeOutputPath (0.49s)
- app | Test_APP_040_ForwardedCheckCode | APP-040 | pass
- app | Test_APP_062_APP_063_APP_064_APP_065_APP_066_APP_067_DoctorCLI | APP-062, APP-063, APP-064, APP-065, APP-066, APP-067 | pass
- app | Test_APP_062_DoctorUsesCloserProject | APP-062 | pass
- app | Test_APP_065_APP_066_APP_067_DoctorFailures | APP-065, APP-066, APP-067 | fail
  Failure: === RUN   Test_APP_065_APP_066_APP_067_DoctorFailures
    doctor_setup_test.go:272: APP-067 static diagnostic=map[code:"C05" message:"C05 packages/alpha/SPEC.md: duplicate item ID: ALP-001; remove or rename the conflicting spec ID" warning:false]
--- FAIL: Test_APP_065_APP_066_APP_067_DoctorFailures (0.77s)
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
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures (0.56s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain
    show_history_cli_test.go:296: status=1; want 2; stdout="" stderr="P01: U02 absent-module: U02 unknown module absent-module. Use a valid archive and correct the reported cause, then retry show --history.\n"
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: U02 absent-module: U02 unknown module absent-module. Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and U02 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain (0.25s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON
    show_history_cli_test.go:296: status=1; want 2; stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"U02 absent-module: U02 unknown module absent-module. Use a valid archive and correct the reported cause, then retry show --history.\",\"warning\":false}]}\n" stderr=""
    show_history_cli_test.go:296: diagnostic code="P01"; want "U02"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON (0.01s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain
    show_history_cli_test.go:296: status=1; want 2; stdout="" stderr="P01: U02 ALP-999: U02 unknown spec ID ALP-999. Use a valid archive and correct the reported cause, then retry show --history.\n"
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: U02 ALP-999: U02 unknown spec ID ALP-999. Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and U02 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain (0.01s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON
    show_history_cli_test.go:296: status=1; want 2; stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"U02 ALP-999: U02 unknown spec ID ALP-999. Use a valid archive and correct the reported cause, then retry show --history.\",\"warning\":false}]}\n" stderr=""
    show_history_cli_test.go:296: diagnostic code="P01"; want "U02"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON (0.01s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/missing-config/plain | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/missing-config/JSON | APP-135 | pass
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain
    show_history_cli_test.go:296: status=1; want 2; stdout="" stderr="P01: C01 fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Use a valid archive and correct the reported cause, then retry show --history.\n"
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: C01 fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and C01 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain (0.01s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON
    show_history_cli_test.go:296: status=1; want 2; stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"C01 fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Use a valid archive and correct the reported cause, then retry show --history.\",\"warning\":false}]}\n" stderr=""
    show_history_cli_test.go:296: diagnostic code="P01"; want "C01"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON (0.01s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: C03 changes/archive/0001-invalid.md: malformed item heading: \"### ALP-001 [broken\". Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and C03 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain (0.01s)
- app | Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON | APP-135 | fail
  Failure: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON
    show_history_cli_test.go:296: diagnostic code="P01"; want "C03"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON (0.01s)
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
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	223.541s

  Stderr: exit status 1

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
ok  	github.com/lurkshark/fireweed/packages/checks/evals	223.923s

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
ok  	github.com/lurkshark/fireweed/packages/runner/evals	233.976s

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
- workflow | Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08 | FLOW-058, FLOW-062 | fail
  Failure: === RUN   Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08
    lifecycle_test.go:382: missing-base static cause = P05: C08: stale base for ALP-001
--- FAIL: Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08 (0.80s)
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
FAIL
FAIL	github.com/lurkshark/fireweed/packages/workflow/evals	92.600s

  Stderr: exit status 1


## Contradicted Decisions
None identified.
## Gaps
- no spec (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite cannot run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt the disposable project internals for loading, queueing, projection, stamping, checks, revision snapshots, spec history, and stale-base recovery. Added [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/DRILL_NOTES.md).

`gofmt` passed. The Go evals could not run in this supplied workspace because it has no `go.mod` or dependency implementations; the notes record that verification gap. Retained inputs were not edited, and nothing was committed.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1998f-7059-79c2-a660-8f8f2f14920b
--------
user
# Rebuild project

Target spec and Decisions:
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
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

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
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash with diagnostic C08.

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

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module.

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens.

### PRJ-027
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### PRJ-028
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### PRJ-029
Queue MUST reject a malformed pending change document.

### PRJ-030
Queue MUST reject a pending change naming an unknown module.

### PRJ-031
Queue MUST reject pending changes that share a slug.

### PRJ-032
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### PRJ-033
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### PRJ-034
Project MUST reject identity or unknown fields in set.

### PRJ-035
Project MUST reject an invalid projected module graph.

### PRJ-036
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

### PRJ-037
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### PRJ-038
Stamp MUST retain other frontmatter values and human section contents.

### PRJ-039
Stamp MUST succeed on an Added-only change without assigning spec IDs.

### PRJ-040
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### PRJ-041
Check MUST report invalid current or projected module graphs with diagnostic C02.

### PRJ-042
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### PRJ-043
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### PRJ-044
Check MUST report pending queue identity errors with diagnostic C13.

### PRJ-045
Check MUST report invalid pending change module targets with diagnostic C07.

### PRJ-046
Check MUST report error C07 for a pending change containing eval_fix together
with fixes or any Added, Modified or Removed spec edits.

### PRJ-047
LoadAt MUST return the selected Git revision's project configuration and modules, using its recorded fireweed.yaml module globs and module SPEC.md documents rather than working-tree contents.

### PRJ-048
LoadAt MUST apply the selected revision's module_defaults to omitted zone fields while preserving explicitly empty module zone lists.

### PRJ-049
LoadAt MUST set Snapshot.Root to the canonical repository root and Module.Dir to each module's repository-relative directory.

### PRJ-050
LoadAt MUST return an error when the revision cannot be resolved, its recorded fireweed.yaml is missing or malformed, or a matched recorded module directory lacks a readable SPEC.md.

### PRJ-051
For a nonempty ID, SpecHistory MUST return the target module's archived changes whose Added or Modified items name that ID, including tag-only modifications.
Pending changes and archives belonging to other modules or unrelated spec IDs MUST not appear in that selection.

### PRJ-052
Every history entry MUST identify its archive path and module, the matching IDs, its Motivation, added and removed Decisions, and its complete Provenance section.

### PRJ-053
SpecHistory MUST order entries by their latest valid implementation Provenance timestamp, newest first, with missing or invalid implementation dates after dated entries.
Equal-date or undated entries MUST use archive-path lexical order as a total-order fallback.

### PRJ-054
With an empty ID filter, SpecHistory MUST return each target-module archive containing Added or Modified spec edits once, with its unique affected IDs in lexical order.
A module with no matching archives MUST return an empty slice.

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching archive edits MUST return U02 with the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce a C01 error.
Malformed archived change documents or malformed Added/Modified item syntax MUST produce a C03 error identifying the archive path.

### PRJ-057
Stamp MUST return an error whose contracts ErrorCode is C08 when the change's bases frontmatter is not a mapping.

### PRJ-058
Stamp MUST return an error whose contracts ErrorCode is U02, with the requested change as subject, when no pending change matches it.

### PRJ-059
Load MUST return an error preserving the contracts ErrorCode C03 or C05 when a module SPEC.md fails to parse.

### PRJ-060
StaleSpecs MUST return the target change's Modified and Removed items whose recorded bases differ from the corresponding preceding projected spec hashes, in edit order, with ID, RecordedHash and Current text.
Items without a bases entry MUST be omitted.

### PRJ-061
StaleSpecs MUST recover Before by projecting the spec immediately before the target change at the Git commit that last changed that change file's bases mapping.
It MUST set Recovered true only when the recovered item's normalized SpecHash equals RecordedHash.

### PRJ-062
When matching stamped history cannot be recovered, StaleSpecs MUST return the stale item's Current text and RecordedHash with Recovered false and empty Before.
Unavailable history MUST NOT prevent reporting other stale items.

### PRJ-063
StaleSpecs MUST leave repository files, index and refs unchanged.
It MUST return an empty slice when no recorded bases are stale.

### PRJ-064
StampMissing MUST add canonical preceding-projection hashes only for Modified or Removed IDs absent from bases, preserving every existing entry including stale hashes.
It MUST retain other frontmatter values and human section contents.

### PRJ-065
StampMissing MUST reject a nonmapping bases value or an edited ID absent from the preceding projection with C08 and preserve the target file on rejection.

### PRJ-066
StaleSpecs and StampMissing MUST accept a unique pending slug, queued basename or repository-relative queued path.
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.


Retained interfaces and files:
packages/project/SPEC.md
packages/project/api.go
packages/project/model/model.go

Retained evals:
packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.

codex
I’ll inspect the retained project API, model, spec, and evals, then rebuild only disposable project files. I’ll run the relevant tests and record any spec gaps or contradictions in `DRILL_NOTES.md`.
exec
/bin/zsh -lc 'git status --short; ls -la packages/project' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "pwd; rg --files -g 'AGENTS.md' -g 'go.mod' -g 'packages/project/**' -g 'packages/contracts/**' -g 'packages/gitstore/**' -g 'packages/assets/**' -g 'DRILL_NOTES.md'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 489ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
packages/gitstore/api.go
packages/gitstore/model/model.go
packages/assets/api.go
packages/assets/model/model.go
packages/project/SPEC.md
packages/project/api.go
packages/project/model/model.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go
packages/project/evals/project_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/queue_test.go
packages/project/evals/.gitkeep
packages/project/evals/revision_snapshot_test.go
packages/contracts/api.go
packages/contracts/model/model.go

 succeeded in 491ms:
total 32
drwxr-xr-x@  6 joel  staff    192 Nov  1 16:01 .
drwxr-xr-x@  6 joel  staff    192 Nov  1 16:01 ..
-rw-r--r--@  1 joel  staff  10038 Nov  1 16:01 SPEC.md
-rw-r--r--@  1 joel  staff   1468 Nov  1 16:01 api.go
drwxr-xr-x@ 11 joel  staff    352 Nov  1 16:01 evals
drwxr-xr-x@  3 joel  staff     96 Nov  1 16:01 model

exec
/bin/zsh -lc 'cat packages/project/api.go packages/project/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/contracts/api.go packages/contracts/model/model.go packages/gitstore/api.go packages/gitstore/model/model.go packages/assets/api.go packages/assets/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/project_test.go packages/project/evals/queue_test.go packages/project/evals/coded_errors_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/revision_snapshot_test.go packages/project/evals/spec_history_test.go packages/project/evals/stale_test.go packages/project/evals/static_codes_test.go packages/project/evals/mixed_eval_fix_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "ls -la; rg --files -g 'go.work' -g 'go.mod' -g 'AGENTS.md' -g 'fireweed.yaml'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 660ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const fixtureConfig = `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
  preserve: [data/**]
`

func put(t *testing.T, root, name, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	put(t, root, "fireweed.yaml", fixtureConfig)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n")
	return root
}

func Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults(t *testing.T) {
	root := fixture(t)
	put(t, root, "other/ignored/SPEC.md", "not a document")
	got, err := project.Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got.Root != root {
		t.Errorf("Root = %q; want %q", got.Root, root)
	}
	if got.Config["version"] != 1 {
		t.Errorf("decoded Config version = %#v", got.Config["version"])
	}
	byName := map[string]project.Module{}
	for _, m := range got.Modules {
		byName[m.Name] = m
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("modules = %#v; want exactly alpha and beta", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Dir != "packages/alpha" || len(alpha.DependsOn) != 0 {
		t.Errorf("alpha identity = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || beta.Dir != "packages/beta" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) {
		t.Errorf("beta identity/dependencies = %#v", beta)
	}
	if alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "### ALP-001\nAlpha behavior.") {
		t.Errorf("alpha Document = %#v", alpha.Document)
	}
	if !reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) {
		t.Errorf("inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone lost: %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_018_LoadInvalidDocuments(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing_config", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_config", func(t *testing.T, root string) { put(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"missing_module_document", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module_document", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			if _, err := project.Load(root); err == nil {
				t.Error("Load accepted missing or malformed project document")
			}
		})
	}
}

func Test_PRJ_002_Graph(t *testing.T) {
	module := func(name, prefix string, deps ...string) project.Module {
		return project.Module{Name: name, IDPrefix: prefix, DependsOn: deps}
	}
	for _, tc := range []struct {
		id      string
		name    string
		modules []project.Module
		valid   bool
		explain []string
	}{
		{"PRJ-002", "acyclic", []project.Module{module("alpha", "ALP"), module("beta", "BET", "alpha"), module("gamma", "GAM", "beta")}, true, nil},
		{"PRJ-019", "empty_name", []project.Module{module("", "ALP")}, false, []string{"empty", "name"}},
		{"PRJ-020", "empty_prefix", []project.Module{module("alpha", "")}, false, []string{"empty", "prefix"}},
		{"PRJ-019", "duplicate_name", []project.Module{module("alpha", "ALP"), module("alpha", "BET")}, false, []string{"duplicate", "alpha"}},
		{"PRJ-020", "duplicate_prefix", []project.Module{module("alpha", "ALP"), module("beta", "ALP")}, false, []string{"duplicate", "alp"}},
		{"PRJ-021", "missing_dependency", []project.Module{module("alpha", "ALP", "missing")}, false, []string{"depend", "missing"}},
		{"PRJ-021", "self_dependency", []project.Module{module("alpha", "ALP", "alpha")}, false, []string{"self", "alpha"}},
		{"PRJ-022", "cycle", []project.Module{module("alpha", "ALP", "beta"), module("beta", "BET", "alpha")}, false, []string{"cycl"}},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			err := project.ValidateGraph(tc.modules)
			if tc.valid && err != nil {
				t.Errorf("valid graph rejected: %v", err)
			}
			if !tc.valid && (err == nil || strings.TrimSpace(err.Error()) == "") {
				t.Errorf("invalid graph needs an explanatory error; got %v", err)
			} else if !tc.valid {
				for _, part := range tc.explain {
					if !strings.Contains(strings.ToLower(err.Error()), part) {
						t.Errorf("error %q does not explain %s rejection (missing %q)", err, tc.name, part)
					}
				}
			}
		})
	}
}

func Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0003-waiting.md", "pending marker")
	put(t, root, "changes/archive/0007-finished.md", "archived marker")
	rel, err := project.NewChange(root, "alpha", "new-feature")
	if err != nil {
		t.Fatalf("NewChange: %v", err)
	}
	if rel != "changes/0008-new-feature.md" {
		t.Errorf("relative path = %q; want changes/0008-new-feature.md", rel)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("read new change: %v", err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("new change document: %v", err)
	}
	if doc.Frontmatter["module"] != "alpha" {
		t.Errorf("new change module = %#v", doc.Frontmatter["module"])
	}
	for _, field := range []string{"phase", "run", "blocked"} {
		if value, present := doc.Frontmatter[field]; present {
			t.Errorf("queued change has %s = %#v", field, value)
		}
	}
	for _, heading := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Decisions"} {
		if _, ok := contracts.Section(doc.Body, heading); !ok {
			t.Errorf("missing %s section", heading)
		}
	}
	if regexp.MustCompile(`(?m)^### [A-Z]+-[0-9]+\b`).Match(data) {
		t.Errorf("skeleton assigned a spec ID: %s", data)
	}
}

func Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
	for _, archived := range []bool{false, true} {
		name := "pending"
		path := "changes/0002-repeat.md"
		if archived {
			name = "archived"
			path = "changes/archive/0002-repeat.md"
		}
		t.Run(name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, path, "original marker")
			if _, err := project.NewChange(root, "alpha", "repeat"); err == nil {
				t.Error("duplicate slug accepted")
			}
			data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
			if err != nil || string(data) != "original marker" {
				t.Errorf("existing change overwritten: %q, %v", data, err)
			}
		})
	}
}

func Test_PRJ_025_PRJ_026_NewChangeInvalidArguments(t *testing.T) {
	for _, tc := range []struct{ id, name, module, slug string }{
		{"PRJ-025", "unknown_module", "missing", "valid"},
		{"PRJ-026", "uppercase_slug", "alpha", "Bad"},
		{"PRJ-026", "underscore_slug", "alpha", "bad_name"},
		{"PRJ-026", "slash_slug", "alpha", "../bad"},
		{"PRJ-026", "empty_slug", "alpha", ""},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			if _, err := project.NewChange(root, tc.module, tc.slug); err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			}
		})
	}
}
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func queued(t *testing.T, root, filename, front, motivation, added, modified, removed string) {
	t.Helper()
	put(t, root, "changes/"+filename, "---\nmodule: alpha\n"+front+"---\n\n## Motivation\n"+motivation+"\n## Interface\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\n\n## Notes\nHuman note.\n")
}

func moduleByName(t *testing.T, snapshot project.Snapshot, name string) project.Module {
	t.Helper()
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module
		}
	}
	t.Fatalf("module %q absent from projection", name)
	return project.Module{}
}

func specItems(t *testing.T, module project.Module) []contracts.Item {
	t.Helper()
	section, ok := contracts.Section(module.Document.Body, "Specs")
	if !ok {
		t.Fatalf("Specs section absent from %s", module.Name)
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		t.Fatalf("parse projected %s specs: %v", module.Name, err)
	}
	return items
}

func findItem(t *testing.T, items []contracts.Item, id string) contracts.Item {
	t.Helper()
	for _, item := range items {
		if item.ID == id {
			return item
		}
	}
	t.Fatalf("spec %s absent from %#v", id, items)
	return contracts.Item{}
}

func Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0002-zed.md", "set:\n  interface: [next.go]\n", "A queue item is needed.\n", "### new:feature\nFeature MUST work.\n", "### ALP-001\nAlpha MUST change.\n", "### ALP-002\n")
	queued(t, root, "0002-ace.md", "", "Another queue item is needed.\n", "", "", "")
	queued(t, root, "0001-first.md", "", "The first queue item is needed.\n", "", "", "")
	put(t, root, "changes/archive/0000-old.md", "archived material is not a pending document")
	changes, err := project.Queue(root)
	if err != nil {
		t.Fatalf("Queue: %v", err)
	}
	if len(changes) != 3 {
		t.Fatalf("queue = %#v; want three pending changes", changes)
	}
	for i, want := range []struct {
		number     int
		slug, path string
	}{
		{1, "first", "changes/0001-first.md"},
		{2, "ace", "changes/0002-ace.md"},
		{2, "zed", "changes/0002-zed.md"},
	} {
		got := changes[i]
		if got.Number != want.number || got.Slug != want.slug || got.Path != want.path || got.Module != "alpha" {
			t.Errorf("change %d = %#v; want %#v with module alpha", i, got, want)
		}
	}
	last := changes[2]
	if !strings.Contains(last.Document.Body, "A queue item is needed.") || last.Document.Frontmatter["module"] != "alpha" || last.Document.Frontmatter["set"] == nil {
		t.Errorf("change document/frontmatter lost: %#v", last.Document)
	}
	if len(last.Added) != 1 || last.Added[0].ID != "new:feature" || !strings.Contains(last.Added[0].Body, "Feature MUST work.") || len(last.Modified) != 1 || last.Modified[0].ID != "ALP-001" || !strings.Contains(last.Modified[0].Body, "Alpha MUST change.") || len(last.Removed) != 1 || last.Removed[0].ID != "ALP-002" {
		t.Errorf("parsed edits = added %#v, modified %#v, removed %#v", last.Added, last.Modified, last.Removed)
	}
}

func Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments(t *testing.T) {
	for _, tc := range []struct{ id, name, body string }{
		{"PRJ-029", "malformed", "---\nmodule: [broken\n---\n"},
		{"PRJ-030", "unknown_module", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "changes/0001-bad.md", tc.body)
			if _, err := project.Queue(root); err == nil {
				t.Error("Queue accepted invalid change")
			}
		})
	}
}

func Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:first\nFirst MUST work.\n", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:second\nSecond MUST work.\n", "", "")
	for _, target := range []string{"first", "0001-first.md", "changes/0001-first.md"} {
		projected, err := project.Project(root, target)
		if err != nil {
			t.Errorf("Project(%q): %v", target, err)
			continue
		}
		items := specItems(t, moduleByName(t, projected, "alpha"))
		if len(items) != 2 || findItem(t, items, "new:first").ID == "" {
			t.Errorf("Project(%q) did not stop at first change: %#v", target, items)
		}
		if err := project.Stamp(root, target); err != nil {
			t.Errorf("Stamp(%q): %v", target, err)
		}
	}
	if _, err := project.Project(root, "missing"); err == nil {
		t.Error("Project accepted missing target")
	}
	if err := project.Stamp(root, "missing"); err == nil {
		t.Error("Stamp accepted missing target")
	}
	put(t, root, "changes/archive/0000-archived.md", "archived document")
	if _, err := project.Project(root, "changes/archive/0000-archived.md"); err == nil {
		t.Error("Project accepted archived target")
	}
	queued(t, root, "0003-first.md", "", "Another first.\n", "", "", "")
	if _, err := project.Queue(root); err == nil {
		t.Error("Queue accepted ambiguous slug")
	}
	if _, err := project.Project(root, "first"); err == nil {
		t.Error("Project accepted ambiguous slug")
	}
	if err := project.Stamp(root, "first"); err == nil {
		t.Error("Stamp accepted ambiguous slug")
	}
}

func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:shared\nFirst change MUST retain its own text.\n", "### ALP-001 [review]\nAlpha MUST have revised full text.\n", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:shared\nSecond change MUST retain different text.\n", "", "### ALP-001\n")
	first, err := project.Project(root, "first")
	if err != nil {
		t.Fatalf("Project first: %v", err)
	}
	firstItems := specItems(t, moduleByName(t, first, "alpha"))
	if len(firstItems) != 2 || findItem(t, firstItems, "ALP-001").Tag != "review" || findItem(t, firstItems, "ALP-001").Body != "Alpha MUST have revised full text.\n\n" || findItem(t, firstItems, "new:shared").Body != "First change MUST retain its own text.\n\n" {
		t.Errorf("first projection = %#v", firstItems)
	}
	all, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project whole queue: %v", err)
	}
	section, ok := contracts.Section(moduleByName(t, all, "alpha").Document.Body, "Specs")
	if !ok {
		t.Fatal("Specs section absent from whole-queue projection")
	}
	if strings.Count(section, "### new:shared\n") != 2 ||
		!strings.Contains(section, "### new:shared\nFirst change MUST retain its own text.\n") ||
		!strings.Contains(section, "### new:shared\nSecond change MUST retain different text.\n") ||
		strings.Contains(section, "ALP-001") {
		t.Errorf("whole-queue projection lost local handles or retained removed ID: %q", section)
	}
}

func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
	queued(t, root, "0001-settings.md", "set:\n  depends_on: [beta]\n  interface: [next.go]\n  evals: [checks/**]\n  disposable: [generated/**]\n  preserve: [records/**]\n  append_only: [migrations/**]\n  boundaries_opt_out: reviewed boundary\n  hooks:\n    test:\n      command: [go, test, ./...]\n", "Settings are needed.\n", "", "", "")
	got, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project settings: %v", err)
	}
	alpha := moduleByName(t, got, "alpha")
	wantLists := map[string][]string{
		"depends_on": {"beta"}, "interface": {"next.go"}, "evals": {"checks/**"},
		"disposable": {"generated/**"}, "preserve": {"records/**"}, "append_only": {"migrations/**"},
	}
	for key, want := range wantLists {
		value, ok := alpha.Document.Frontmatter[key].([]any)
		if !ok || len(value) != 1 || value[0] != want[0] {
			t.Errorf("set %s = %#v; want %q", key, alpha.Document.Frontmatter[key], want)
		}
	}
	if !reflect.DeepEqual(alpha.DependsOn, []string{"beta"}) || !reflect.DeepEqual(alpha.Zones.Interface, []string{"next.go"}) || !reflect.DeepEqual(alpha.Zones.Evals, []string{"checks/**"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"generated/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"records/**"}) || !reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"migrations/**"}) {
		t.Errorf("projected module fields not updated: %#v", alpha)
	}
	wantHooks := map[string]any{"test": map[string]any{"command": []any{"go", "test", "./..."}}}
	if alpha.Document.Frontmatter["boundaries_opt_out"] != "reviewed boundary" || !reflect.DeepEqual(alpha.Document.Frontmatter["hooks"], wantHooks) {
		t.Errorf("projected scalar/map settings lost: %#v", alpha.Document.Frontmatter)
	}
	for _, tc := range []struct{ id, name, settings string }{
		{"PRJ-034", "identity_name", "name: renamed"},
		{"PRJ-034", "identity_prefix", "id_prefix: NEW"},
		{"PRJ-034", "unknown_field", "surprise: true"},
		{"PRJ-035", "invalid_graph", "depends_on: [missing]"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			badRoot := fixture(t)
			queued(t, badRoot, "0001-invalid.md", "set:\n  "+tc.settings+"\n", "Needed.\n", "", "", "")
			if _, err := project.Project(badRoot, ""); err == nil {
				t.Errorf("Project accepted %s", tc.name)
			}
		})
	}
}

func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-revise.md", "", "Earlier change.\n", "", "### ALP-001 [review]\nAlpha MUST reflect the earlier revision.\n", "")
	queued(t, root, "0002-target.md", "custom: author-value\n", "Keep this human motivation.\n", "", "### ALP-001\nAlpha MUST reflect the later revision.\n", "")
	wantHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Tag: "review", Body: "Alpha MUST reflect the earlier revision.\n"})
	if err := project.Stamp(root, "changes/0002-target.md"); err != nil {
		t.Fatalf("Stamp: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(root, "changes/0002-target.md"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("stamped document: %v", err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok || bases["ALP-001"] != wantHash {
		t.Errorf("bases = %#v; want preceding hash %s", doc.Frontmatter["bases"], wantHash)
	}
	if doc.Frontmatter["custom"] != "author-value" || !strings.Contains(doc.Body, "Keep this human motivation.") || !strings.Contains(doc.Body, "Alpha MUST reflect the later revision.") || !strings.Contains(doc.Body, "Human note.") {
		t.Errorf("Stamp lost author content: %#v", doc)
	}
	queued(t, root, "0003-remove.md", "", "Remove the item.\n", "", "", "### ALP-001\n")
	if err := project.Stamp(root, "0003-remove.md"); err != nil {
		t.Fatalf("Stamp Removed: %v", err)
	}
	removedBytes, err := os.ReadFile(filepath.Join(root, "changes/0003-remove.md"))
	if err != nil {
		t.Fatal(err)
	}
	removedDoc, err := contracts.ParseDocument(string(removedBytes))
	if err != nil {
		t.Fatal(err)
	}
	removedWant := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST reflect the later revision.\n"})
	removedBases, ok := removedDoc.Frontmatter["bases"].(map[string]any)
	if !ok || removedBases["ALP-001"] != removedWant {
		t.Errorf("Removed base = %#v; want %s", removedDoc.Frontmatter["bases"], removedWant)
	}
}

func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-added.md", "", "Added item.\n", "### new:thing\nThing MUST work.\n", "", "")
	if err := project.Stamp(root, "added"); err != nil {
		t.Fatalf("Stamp Added-only: %v", err)
	}
	bytes, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bytes), "### new:thing") || strings.Contains(string(bytes), "### ALP-002") {
		t.Errorf("Stamp assigned an ID to Added-only change: %s", bytes)
	}
	queued(t, root, "0002-invalid.md", "", "Invalid edit.\n", "", "### ALP-999\nMissing item MUST change.\n", "")
	if err := project.Stamp(root, "invalid"); err == nil {
		t.Error("Stamp accepted Modified ID absent from preceding projection")
	}
	removeRoot := fixture(t)
	queued(t, removeRoot, "0001-invalid-remove.md", "", "Invalid removal.\n", "", "", "### ALP-998\n")
	if err := project.Stamp(removeRoot, "invalid-remove"); err == nil {
		t.Error("Stamp accepted Removed ID absent from preceding projection")
	}
}

func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
	validRoot := fixture(t)
	put(t, validRoot, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	queued(t, validRoot, "0001-valid.md", "bases:\n  ALP-001: "+contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})+"\n", "A user needs the change.\n", "### new:local [review]\nA local addition MUST work.\n", "### ALP-001 [quarantine: flaky fixture]\nAlpha MUST work with revision.\n", "")
	queued(t, validRoot, "0002-also-valid.md", "", "Another user need.\n", "### new:local\nA second local addition MUST work.\n", "", "")
	for _, diagnostic := range project.Check(validRoot) {
		if !diagnostic.Warning {
			t.Errorf("valid spec/change received error diagnostic: %#v", diagnostic)
		}
	}
}

func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
	for _, tc := range []struct {
		id, name, spec, front, motivation, added, modified, removed string
		code, cause                                                 string
	}{
		{"PRJ-008", "missing_motivation", "", "", "\n", "", "", "", "C07", "missing Motivation"},
		{"PRJ-009", "nonnormative_module", "### ALP-001\nAlpha works.\n", "", "Needed.\n", "", "", "", "C04", "non-normative spec body"},
		{"PRJ-009", "nonnormative_change", "", "", "Needed.\n", "### new:item\nAn observation.\n", "", "", "C04", "non-normative spec body"},
		{"PRJ-010", "invalid_review_tag", "", "", "Needed.\n", "### new:item [review: reason]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-010", "invalid_quarantine_tag", "", "", "Needed.\n", "### new:item [quarantine]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-011", "duplicate_module_id", "### ALP-001\nAlpha MUST work.\n\n### ALP-001\nDuplicate MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "duplicate item ID"},
		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
		{"PRJ-013", "duplicate_edit_identity", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "### ALP-001\n", "C05", "duplicate edit identity ALP-001"},
		{"PRJ-014", "undeclared_handle", "", "", "Needed.\n", "### new:local\nLocal MUST work.\n", "### ALP-001\nAlpha MUST refer to new:foreign.\n", "", "C09", "reference to undeclared handle new:foreign"},
		{"PRJ-015", "missing_base", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "missing base for ALP-001"},
		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
		{"PRJ-016", "absent_modified", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "### ALP-999\nAbsent MUST change.\n", "", "C08", "ID ALP-999 absent from preceding projection"},
		{"PRJ-016", "absent_removed", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "", "### ALP-999\n", "C08", "ID ALP-999 absent from preceding projection"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			front := tc.front
			baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
			if tc.name == "duplicate_edit_identity" || tc.name == "undeclared_handle" {
				front = "bases:\n  ALP-001: " + baseHash + "\n"
			}
			if tc.name == "foreign_change_id" {
				put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
				front = "bases:\n  BET-001: " + contracts.SpecHash(contracts.Item{ID: "BET-001", Body: "Beta MUST work.\n"}) + "\n"
			}
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n"+func() string {
				if tc.spec != "" {
					return tc.spec
				}
				return "### ALP-001\nAlpha MUST work.\n"
			}())
			if tc.name == "undeclared_handle" {
				queued(t, root, "0000-other.md", "", "Earlier local handle.\n", "### new:foreign\nEarlier item MUST work.\n", "", "")
			}
			queued(t, root, "0001-check.md", front, tc.motivation, tc.added, tc.modified, tc.removed)
			diagnostics := project.Check(root)
			found := false
			foundCause := false
			foundApplyCause := false
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
					t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
				}
				if !diagnostic.Warning && strings.TrimSpace(diagnostic.Code) != "" && strings.TrimSpace(diagnostic.Message) != "" {
					found = true
				}
				matchesCause := strings.Contains(diagnostic.Message, tc.cause)
				if tc.id == "PRJ-016" {
					matchesCause = diagnostic.Message == tc.cause
				}
				if tc.code != "" && !diagnostic.Warning && diagnostic.Code == tc.code && matchesCause {
					foundCause = true
				}
				if tc.name == "absent_modified" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Modified ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
				if tc.name == "absent_removed" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Removed ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
			}
			if !found {
				t.Errorf("Check did not report error diagnostic with nonempty code and message: %#v", diagnostics)
			}
			if tc.code != "" && !foundCause {
				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
			}
			if tc.id == "PRJ-016" && !foundApplyCause {
				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
			}
			if tc.name == "missing_base" || tc.name == "stale_base" {
				queued(t, root, "0001-check.md", front, "Needed.\n", "", "", "### ALP-001\n")
				removedDiagnostics := project.Check(root)
				found = false
				for _, diagnostic := range removedDiagnostics {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message for Removed entry: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, tc.cause) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not report C08 for Removed entry with %s (%s): %#v", tc.name, tc.cause, removedDiagnostics)
				}
			}
			if tc.name == "nonnormative_change" || tc.name == "invalid_review_tag" {
				front, added, modified, want := "", "", "", ""
				code := "C04"
				if tc.name == "nonnormative_change" {
					front = "bases:\n  ALP-001: " + baseHash + "\n"
					modified = "### ALP-001\nAn observation.\n"
					want = "non-normative spec body"
				} else {
					code = "C06"
					added = "### new:item [other: reason]\nItem MUST work.\n"
					want = "invalid review/quarantine tag"
				}
				queued(t, root, "0001-check.md", front, "Needed.\n", added, modified, "")
				found = false
				for _, diagnostic := range project.Check(root) {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == code && strings.Contains(diagnostic.Message, want) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not reject %s: wanted error diagnostic %s with message containing %q", tc.name, code, want)
				}
			}
		})
	}
}
package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_057_StampInvalidBasesHasCode(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-invalid-bases.md", "bases: [ALP-001]\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "")
	err := project.Stamp(root, "invalid-bases")
	if err == nil {
		t.Fatal("Stamp accepted bases that are not a mapping")
	}
	if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
		t.Errorf("Stamp error code = (%q, %v); want (C08, true): %v", code, ok, err)
	}
}

func Test_PRJ_058_MissingPendingChangeHasSubject(t *testing.T) {
	root := fixture(t)
	const target = "changes/0099-absent.md"
	for _, operation := range []struct {
		name string
		call func() error
	}{
		{"Stamp", func() error { return project.Stamp(root, target) }},
		{"Project", func() error { _, err := project.Project(root, target); return err }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			err := operation.call()
			if err == nil {
				t.Fatal("missing pending change was accepted")
			}
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
				t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, target, err)
			}
		})
	}
}

func Test_PRJ_059_LoadPreservesModuleParseCode(t *testing.T) {
	for _, tc := range []struct {
		name, spec, code string
	}{
		{"malformed_frontmatter", "---\nname: [broken\n---\n# Alpha\n", "C03"},
		{"malformed_spec_heading", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001 [broken\nAlpha MUST work.\n", "C03"},
		{"duplicate_spec_id", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n\n### ALP-001\nAlpha MUST also work.\n", "C05"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", tc.spec)
			_, err := project.Load(root)
			if err == nil {
				t.Fatal("Load accepted malformed module SPEC.md")
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != tc.code {
				t.Errorf("Load error code = (%q, %v); want (%s, true): %v", code, ok, tc.code, err)
			}
		})
	}
}

 succeeded in 668ms:
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

 exited 1 in 669ms:
total 8
drwx------@    5 joel  staff     160 Nov  1 16:01 .
drwx------@ 6843 joel  staff  218976 Nov  1 16:01 ..
drwxr-xr-x@    9 joel  staff     288 Nov  1 16:01 .git
-rw-r--r--@    1 joel  staff    2965 Nov  1 16:01 SYSTEM.md
drwxr-xr-x@    6 joel  staff     192 Nov  1 16:01 packages

 succeeded in 680ms:
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

 succeeded in 675ms:
package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func revisionGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func revisionCommit(t *testing.T, root, subject string) string {
	t.Helper()
	revisionGit(t, root, "add", "-A")
	revisionGit(t, root, "-c", "user.name=Fireweed Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-qm", subject)
	return revisionGit(t, root, "rev-parse", "HEAD")
}

func historicalProject(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "fireweed.yaml", `version: 1
modules: ["packages/*"]
module_defaults:
  protected: [guard/**]
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
  preserve: [data/**]
  disposable: [internal/**]
`)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nHistorical alpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nHistorical beta behavior.\n")
	put(t, root, "extras/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: []\n---\n# Gamma\n\n## Specs\n\n### GAM-001\nHistorical gamma behavior.\n")
	recorded := revisionCommit(t, root, "record original project")

	put(t, root, "fireweed.yaml", "version: 2\nmodules: [\"components/*\"]\nmodule_defaults:\n  evals: [later/**]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: later-alpha\nid_prefix: LAL\n---\n# Later\n\n## Specs\n")
	put(t, root, "components/next/SPEC.md", "---\nname: next\nid_prefix: NXT\n---\n# Next\n\n## Specs\n")
	revisionCommit(t, root, "change config glob and modules")

	put(t, root, "fireweed.yaml", "version: 3\nmodules: [\"dirty/*\"]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: dirty-alpha\nid_prefix: DAL\n---\n# Dirty\n\n## Specs\n")
	put(t, root, "dirty/only/SPEC.md", "---\nname: dirty\nid_prefix: DRT\n---\n# Dirty\n\n## Specs\n")
	return root, recorded
}

func Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Config["version"] != 1 || !reflect.DeepEqual(got.Config["modules"], []any{"packages/*"}) {
		t.Errorf("recorded configuration = %#v", got.Config)
	}
	wantDefaults := map[string]any{
		"protected":   []any{"guard/**"},
		"evals":       []any{"evals/**"},
		"interface":   []any{"api.go"},
		"append_only": []any{"ledger/**"},
		"preserve":    []any{"data/**"},
		"disposable":  []any{"internal/**"},
	}
	if !reflect.DeepEqual(got.Config["module_defaults"], wantDefaults) {
		t.Errorf("recorded module_defaults = %#v; want %#v", got.Config["module_defaults"], wantDefaults)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("recorded modules = %#v; want alpha and beta only", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "Historical alpha behavior.") {
		t.Errorf("recorded alpha document = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) || !strings.Contains(beta.Document.Body, "Historical beta behavior.") {
		t.Errorf("recorded beta document = %#v", beta)
	}
}

func Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 {
		t.Fatalf("recorded modules = %#v", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if !reflect.DeepEqual(alpha.Zones.Protected, []string{"guard/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) ||
		!reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"ledger/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) {
		t.Errorf("recorded inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone = %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("recorded default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs(t *testing.T) {
	root, recorded := historicalProject(t)
	canonical, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := project.LoadAt(root+string(filepath.Separator)+".", recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Root != canonical {
		t.Errorf("Root = %q; want canonical %q", got.Root, canonical)
	}
	dirs := make(map[string]string)
	for _, module := range got.Modules {
		dirs[module.Name] = module.Dir
	}
	if !reflect.DeepEqual(dirs, map[string]string{"alpha": "packages/alpha", "beta": "packages/beta"}) {
		t.Errorf("module directories = %#v", dirs)
	}
}

func Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs(t *testing.T) {
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "README.md", "repository without fireweed.yaml\n")
	missingConfig := revisionCommit(t, root, "missing config")
	put(t, root, "fireweed.yaml", "modules: [broken\n")
	malformedConfig := revisionCommit(t, root, "malformed config")
	put(t, root, "fireweed.yaml", "modules: [\"modules/*\"]\n")
	put(t, root, "modules/ghost/README.md", "directory without SPEC.md\n")
	missingSpec := revisionCommit(t, root, "missing module document")
	for _, tc := range []struct{ name, revision string }{
		{"invalid_revision", "not-a-revision"},
		{"missing_recorded_config", missingConfig},
		{"malformed_recorded_config", malformedConfig},
		{"missing_recorded_SPEC", missingSpec},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := project.LoadAt(root, tc.revision); err == nil {
				t.Errorf("LoadAt accepted %s (%s)", tc.name, tc.revision)
			}
		})
	}
}
package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

func historyArchive(t *testing.T, root, path, module, motivation, added, modified, removed, decisions, removedDecisions, provenance string) {
	t.Helper()
	put(t, root, path, fmt.Sprintf("---\nmodule: %s\nphase: applied\n---\n\n## Motivation\n%s## Added\n%s## Modified\n%s## Removed\n%s## Decisions\n%s## Removed decisions\n%s## Provenance\n%s", module, motivation, added, modified, removed, decisions, removedDecisions, provenance))
}

func historyProject(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	revisionGit(t, root, "init", "-q")
	historyArchive(t, root, "changes/archive/0009-added.md", "alpha",
		"Why the behavior exists.\nSecond motivation line.\n\n",
		"### ALP-004\nAnother behavior.\n### ALP-002\nRetained behavior.\n", "", "",
		"Choose the stable form.\nKeep this second decision.\n\n",
		"Prior choice was dropped.\n\n",
		"- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n")
	revisionCommit(t, root, "archive newest implementation date first")
	historyArchive(t, root, "changes/archive/0001-retag.md", "alpha", "Retag the behavior.\n", "",
		"### ALP-002 [review]\nRetained behavior.\n", "", "", "", "- impl 2024-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0003-equal.md", "alpha", "First equal date.\n", "",
		"### ALP-002\nEqual date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0004-equal.md", "alpha", "Second equal date.\n", "",
		"### ALP-002\nAnother equal date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0005-bad-date.md", "alpha", "Invalid date.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- impl not-a-date agent=fixture\n")
	historyArchive(t, root, "changes/archive/0006-no-impl.md", "alpha", "No implementation line.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- evals 2030-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0010-other-id.md", "alpha", "Other specs.\n",
		"### ALP-005\nFifth behavior.\n### ALP-003\nThird behavior.\n",
		"### ALP-003 [review]\nThird behavior retagged.\n", "", "", "", "")
	historyArchive(t, root, "changes/archive/0011-removed-only.md", "alpha", "Only a removal.\n", "", "",
		"### ALP-002\n", "", "", "- impl 2031-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0012-beta.md", "beta", "Beta history.\n",
		"### BET-001\nBeta behavior.\n", "", "", "", "", "- impl 2032-01-01T00:00:00Z agent=fixture\n")
	queued(t, root, "0013-pending.md", "", "Not archived.\n", "### ALP-002\nPending behavior.\n", "", "")
	revisionCommit(t, root, "archive older dates and pending change later")
	return root
}

func historyPaths(entries []project.SpecHistoryEntry) []string {
	paths := make([]string, len(entries))
	for i, entry := range entries {
		paths[i] = entry.Path
	}
	return paths
}

func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	want := map[string]bool{
		"changes/archive/0009-added.md":    true,
		"changes/archive/0001-retag.md":    true,
		"changes/archive/0003-equal.md":    true,
		"changes/archive/0004-equal.md":    true,
		"changes/archive/0005-bad-date.md": true,
		"changes/archive/0006-no-impl.md":  true,
	}
	if len(entries) != len(want) {
		t.Fatalf("selected archives = %#v; want paths %#v", historyPaths(entries), want)
	}
	for _, entry := range entries {
		if !want[entry.Path] || entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, []string{"ALP-002"}) {
			t.Errorf("unexpected path/module/filtered IDs = %q, %q, %#v", entry.Path, entry.Module, entry.IDs)
		}
		delete(want, entry.Path)
	}
	if len(want) != 0 {
		t.Errorf("missing archives: %#v", want)
	}
}

func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("archive order/selection = %#v; want %#v", got, wantPaths)
	}
	wantDates := []time.Time{
		time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		{}, {},
	}
	for i, entry := range entries {
		if !entry.ImplementationDate.Equal(wantDates[i]) {
			t.Errorf("%s implementation date = %s; want %s", entry.Path, entry.ImplementationDate, wantDates[i])
		}
	}
}

func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("missing archived ALP-002 history")
	}
	want := project.SpecHistoryEntry{
		Path: "changes/archive/0009-added.md", Module: "alpha", IDs: []string{"ALP-002"},
		Motivation:         "Why the behavior exists.\nSecond motivation line.\n\n",
		Decisions:          "Choose the stable form.\nKeep this second decision.\n\n",
		RemovedDecisions:   "Prior choice was dropped.\n\n",
		Provenance:         "- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n",
		ImplementationDate: time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(entries[0], want) {
		t.Errorf("entry = %#v; want %#v", entries[0], want)
	}
}

func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "")
	if err != nil {
		t.Fatalf("SpecHistory module-wide: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
		"changes/archive/0010-other-id.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("module-wide archive order/selection = %#v; want %#v", got, wantPaths)
	}
	for _, entry := range entries {
		wantIDs := []string{"ALP-002"}
		switch entry.Path {
		case "changes/archive/0009-added.md":
			wantIDs = []string{"ALP-002", "ALP-004"}
		case "changes/archive/0010-other-id.md":
			wantIDs = []string{"ALP-003", "ALP-005"}
		}
		if entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, wantIDs) {
			t.Errorf("%s module/unique IDs = %q, %#v; want alpha, %#v", entry.Path, entry.Module, entry.IDs, wantIDs)
		}
	}
	emptyRoot := fixture(t)
	empty, err := project.SpecHistory(emptyRoot, "beta", "")
	if err != nil || len(empty) != 0 {
		t.Errorf("module without archives = %#v, %v; want empty slice", empty, err)
	}
}

func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
	root := historyProject(t)
	archived, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil || len(archived) == 0 {
		t.Errorf("archived ID absent from current Specs = %#v, %v", archived, err)
	}
	currentOnly, err := project.SpecHistory(root, "alpha", "ALP-001")
	if err != nil || len(currentOnly) != 0 {
		t.Errorf("current ID with no archive edits = %#v, %v; want empty slice", currentOnly, err)
	}
	for _, tc := range []struct{ module, id, rejected string }{
		{"absent-module", "", "absent-module"},
		{"alpha", "ALP-999", "ALP-999"},
		{"alpha", "BET-001", "BET-001"},
	} {
		_, err := project.SpecHistory(root, tc.module, tc.id)
		if err == nil || !strings.HasPrefix(err.Error(), "U02") || !strings.Contains(err.Error(), tc.rejected) {
			t.Errorf("lookup module %q ID %q = %v; want U02 naming %q", tc.module, tc.id, err, tc.rejected)
		}
	}
}

func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
	for _, tc := range []struct{ name, config string }{
		{"missing", ""},
		{"malformed", "modules: [broken\n"},
	} {
		t.Run("config_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			if tc.name == "missing" {
				if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
			} else {
				put(t, root, "fireweed.yaml", tc.config)
			}
			_, err := project.SpecHistory(root, "alpha", "")
			if err == nil || !strings.HasPrefix(err.Error(), "C01") {
				t.Errorf("invalid config = %v; want C01", err)
			}
		})
	}
	for _, tc := range []struct{ name, document string }{
		{"document", "---\nmodule: [broken\n---\n## Added\n### ALP-002\nBehavior.\n"},
		{"added", "---\nmodule: alpha\n---\n## Added\n### ALP-002 [broken\nBehavior.\n"},
		{"modified", "---\nmodule: alpha\n---\n## Modified\n### ALP-002 [broken\nBehavior.\n"},
	} {
		t.Run("archive_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			path := "changes/archive/0001-invalid.md"
			put(t, root, path, tc.document)
			_, err := project.SpecHistory(root, "alpha", "ALP-002")
			if err == nil || !strings.HasPrefix(err.Error(), "C03") || !strings.Contains(err.Error(), path) {
				t.Errorf("malformed archive = %v; want C03 naming %s", err, path)
			}
		})
	}
}
package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func staleSpec(t *testing.T, root, through, id string) contracts.Item {
	t.Helper()
	snapshot, err := project.Project(root, through)
	if err != nil {
		t.Fatal(err)
	}
	return findItem(t, specItems(t, moduleByName(t, snapshot, "alpha")), id)
}

func staleBaseItem(t *testing.T, root, id string) contracts.Item {
	t.Helper()
	snapshot, err := project.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return findItem(t, specItems(t, moduleByName(t, snapshot, "alpha")), id)
}

func staleRead(t *testing.T, root, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func staleBases(t *testing.T, root, path string) (map[string]any, contracts.Document) {
	t.Helper()
	doc, err := contracts.ParseDocument(string(staleRead(t, root, path)))
	if err != nil {
		t.Fatal(err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok {
		t.Fatalf("bases = %#v, want mapping", doc.Frontmatter["bases"])
	}
	return bases, doc
}

func staleFixture(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nOne MUST work.\n\n### ALP-002\nTwo MUST work.\n\n### ALP-003\nThree MUST work.\n\n### ALP-004\nFour MUST work.\n")
	return root
}

func Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder(t *testing.T) {
	root := staleFixture(t)
	oldOne := contracts.SpecHash(staleSpec(t, root, "", "ALP-001"))
	oldTwo := contracts.SpecHash(staleSpec(t, root, "", "ALP-002"))
	oldThree := contracts.SpecHash(staleSpec(t, root, "", "ALP-003"))
	queued(t, root, "0001-earlier.md", "", "Earlier edits.\n", "", "### ALP-001 [review]\nOne MUST change.\n\n### ALP-002 [review]\nTwo MUST change.\n\n### ALP-003 [review]\nThree MUST change.\n", "")
	queued(t, root, "0002-target.md", "bases:\n  ALP-002: "+oldTwo+"\n  ALP-001: "+oldOne+"\n  ALP-003: "+oldThree+"\n", "Later edits.\n", "", "### ALP-002\nTwo MUST change again.\n\n### ALP-001\nOne MUST change again.\n\n### ALP-004\nFour MUST change without a stamp.\n", "### ALP-003\n")
	got, err := project.StaleSpecs(root, "target")
	if err != nil {
		t.Fatal(err)
	}
	want := []struct{ id, hash, current string }{
		{"ALP-002", oldTwo, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-002"))},
		{"ALP-001", oldOne, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-001"))},
		{"ALP-003", oldThree, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-003"))},
	}
	if len(got) != len(want) {
		t.Fatalf("stale entries = %#v; want %d recorded entries", got, len(want))
	}
	for i, entry := range got {
		if entry.ID != want[i].id || entry.RecordedHash != want[i].hash || entry.Current != want[i].current {
			t.Errorf("entry %d = %#v; want ID %s, hash %s, Current %q", i, entry, want[i].id, want[i].hash, want[i].current)
		}
	}
}

func Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash(t *testing.T) {
	root := staleFixture(t)
	revisionGit(t, root, "init", "-q")
	queued(t, root, "0001-earlier.md", "", "Earlier.\n", "", "### ALP-001 [review]\nOne MUST use old text.  \n\n### ALP-002 [review]\nTwo MUST use old text.\n", "")
	queued(t, root, "0002-target.md", "", "Target.\n", "", "### ALP-001\nOne MUST be edited.\n", "### ALP-002\n")
	revisionCommit(t, root, "initial unstamped queue")
	before := staleSpec(t, root, "earlier", "ALP-001")
	if before.Tag != "review" || !strings.Contains(before.Body, "old text.  ") {
		t.Fatalf("fixture lost tag or trailing spaces: %#v", before)
	}
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	path := "changes/0002-target.md"
	bases, doc := staleBases(t, root, path)
	if bases["ALP-001"] != contracts.SpecHash(before) {
		t.Fatalf("stamped hash %v differs from normalized tagged hash %s", bases["ALP-001"], contracts.SpecHash(before))
	}
	// A second recorded hash is intentionally wrong. Recovery must be per item.
	bases["ALP-002"] = "deadbeef"
	text, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	put(t, root, path, text)
	revisionCommit(t, root, "last bases mapping change")
	queued(t, root, "0001-earlier.md", "", "Earlier.\n", "", "### ALP-001 [quarantine: changed]\nOne MUST use current text.\n\n### ALP-002\nTwo MUST use current text.\n", "")
	// This later target-file commit changes only prose, not the bases mapping.
	put(t, root, path, strings.Replace(string(staleRead(t, root, path)), "Target.", "Target with later prose.", 1))
	revisionCommit(t, root, "later unrelated change to target file")
	currentOne := staleSpec(t, root, "earlier", "ALP-001")
	currentTwo := staleSpec(t, root, "earlier", "ALP-002")
	got, err := project.StaleSpecs(root, "target")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("stale entries = %#v; want two", got)
	}
	if got[0].ID != "ALP-001" || got[0].RecordedHash != contracts.SpecHash(before) || !got[0].Recovered || got[0].Before != contracts.RenderItem(before) || got[0].Current != contracts.RenderItem(currentOne) {
		t.Errorf("recovered tagged item = %#v; want historical pre-target item %#v and current %#v", got[0], before, currentOne)
	}
	if got[1].ID != "ALP-002" || got[1].RecordedHash != "deadbeef" || got[1].Recovered || got[1].Before != "" || got[1].Current != contracts.RenderItem(currentTwo) {
		t.Errorf("mismatched historical hash = %#v; want current text without Before", got[1])
	}
}

func Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem(t *testing.T) {
	root := staleFixture(t) // No Git repository or recorded target history.
	queued(t, root, "0001-target.md", "bases:\n  ALP-001: deadbeef\n  ALP-002: 12345678\n", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
	got, err := project.StaleSpecs(root, "target")
	if err != nil || len(got) != 2 {
		t.Fatalf("unavailable history = %#v, %v; want both stale items", got, err)
	}
	for i, want := range []struct{ id, hash string }{{"ALP-001", "deadbeef"}, {"ALP-002", "12345678"}} {
		if got[i].ID != want.id || got[i].RecordedHash != want.hash || got[i].Recovered || got[i].Before != "" || got[i].Current != contracts.RenderItem(staleBaseItem(t, root, want.id)) {
			t.Errorf("entry %d = %#v; want unrecovered %s", i, got[i], want.id)
		}
	}
}

func Test_PRJ_063_StaleSpecsEmptyAndReadOnly(t *testing.T) {
	root := staleFixture(t)
	revisionGit(t, root, "init", "-q")
	queued(t, root, "0001-target.md", "", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "")
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	revisionCommit(t, root, "stamped target")
	path := "changes/0001-target.md"
	readState := func() (target, index []byte, head, refs string) {
		target = staleRead(t, root, path)
		index = staleRead(t, root, ".git/index")
		head = revisionGit(t, root, "rev-parse", "HEAD")
		refs = revisionGit(t, root, "show-ref")
		return
	}
	target, index, head, refs := readState()
	got, err := project.StaleSpecs(root, "target")
	if err != nil || len(got) != 0 {
		t.Fatalf("matching base = %#v, %v; want empty slice", got, err)
	}
	put(t, root, "packages/alpha/SPEC.md", strings.Replace(string(staleRead(t, root, "packages/alpha/SPEC.md")), "One MUST work.", "One MUST now work differently.", 1))
	spec := staleRead(t, root, "packages/alpha/SPEC.md")
	got, err = project.StaleSpecs(root, "target")
	if err != nil || len(got) != 1 {
		t.Fatalf("stale base = %#v, %v; want one", got, err)
	}
	afterTarget, afterIndex, afterHead, afterRefs := readState()
	if !bytes.Equal(target, afterTarget) || !bytes.Equal(spec, staleRead(t, root, "packages/alpha/SPEC.md")) || !bytes.Equal(index, afterIndex) || head != afterHead || refs != afterRefs {
		t.Error("StaleSpecs changed a worktree file, index or ref")
	}
}

func Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent(t *testing.T) {
	root := staleFixture(t)
	queued(t, root, "0001-target.md", "custom: keep-me\nbases:\n  ALP-001: deadbeef\n  ALP-002: "+contracts.SpecHash(staleBaseItem(t, root, "ALP-002"))+"\n  ALP-999: extra\n", "Keep this motivation.\n", "", "### ALP-001\nOne MUST change.\n\n### ALP-002\nTwo MUST change.\n\n### ALP-003\nThree MUST change.\n", "### ALP-004\n")
	path := "changes/0001-target.md"
	_, original := staleBases(t, root, path)
	beforeCheck := project.Check(root)
	if err := project.StampMissing(root, "target"); err != nil {
		t.Fatal(err)
	}
	bases, after := staleBases(t, root, path)
	want := map[string]any{
		"ALP-001": "deadbeef", "ALP-002": contracts.SpecHash(staleBaseItem(t, root, "ALP-002")),
		"ALP-003": contracts.SpecHash(staleBaseItem(t, root, "ALP-003")),
		"ALP-004": contracts.SpecHash(staleBaseItem(t, root, "ALP-004")), "ALP-999": "extra",
	}
	if !reflect.DeepEqual(bases, want) || after.Frontmatter["custom"] != "keep-me" || after.Body != original.Body {
		t.Errorf("StampMissing lost bases/frontmatter/human sections: bases=%#v, body changed=%v", bases, after.Body != original.Body)
	}
	if len(beforeCheck) == 0 {
		t.Fatal("fixture lacked Check diagnostics before stamping")
	}
	staleFound := false
	for _, diagnostic := range project.Check(root) {
		if diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "stale base for ALP-001") {
			staleFound = true
		}
	}
	if !staleFound {
		t.Error("ordinary Check stopped reporting the preserved stale base")
	}
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	stamped, _ := staleBases(t, root, path)
	if stamped["ALP-001"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-001")) {
		t.Errorf("ordinary Stamp stopped replacing a stale base: %#v", stamped)
	}
}

func Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically(t *testing.T) {
	for _, tc := range []struct{ name, front, modified, removed string }{
		{"nonmapping", "bases: [ALP-001]\n", "### ALP-001\nOne MUST change.\n", ""},
		{"absent_modified_after_valid", "", "### ALP-001\nOne MUST change.\n\n### ALP-999\nMissing MUST change.\n", ""},
		{"absent_removed_after_valid", "", "### ALP-001\nOne MUST change.\n", "### ALP-998\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := staleFixture(t)
			path := "changes/0001-target.md"
			queued(t, root, "0001-target.md", tc.front, "Needed.\n", "", tc.modified, tc.removed)
			before := staleRead(t, root, path)
			err := project.StampMissing(root, "target")
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
				t.Errorf("StampMissing error = %v; want C08", err)
			}
			if !bytes.Equal(before, staleRead(t, root, path)) {
				t.Error("StampMissing rewrote target after rejection")
			}
		})
	}
}

func Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures(t *testing.T) {
	for _, alias := range []string{"target", "0001-target.md", "changes/0001-target.md"} {
		t.Run(alias, func(t *testing.T) {
			root := staleFixture(t)
			queued(t, root, "0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
			got, err := project.StaleSpecs(root, alias)
			if err != nil || len(got) != 1 || got[0].ID != "ALP-001" {
				t.Errorf("StaleSpecs(%q) = %#v, %v", alias, got, err)
			}
			if err := project.StampMissing(root, alias); err != nil {
				t.Errorf("StampMissing(%q): %v", alias, err)
			}
			bases, _ := staleBases(t, root, "changes/0001-target.md")
			if bases["ALP-002"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-002")) || bases["ALP-001"] != "deadbeef" {
				t.Errorf("StampMissing(%q) bases = %#v", alias, bases)
			}
		})
	}
	root := staleFixture(t)
	for _, operation := range []struct {
		name string
		call func(string) error
	}{
		{"StaleSpecs", func(target string) error { _, err := project.StaleSpecs(root, target); return err }},
		{"StampMissing", func(target string) error { return project.StampMissing(root, target) }},
	} {
		t.Run(operation.name+"_unknown", func(t *testing.T) {
			const target = "changes/0099-absent.md"
			err := operation.call(target)
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
				t.Errorf("unknown selector = %v; want U02 with subject %q", err, target)
			}
		})
	}
	queued(t, root, "0001-target.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-target.md", "", "Second.\n", "", "", "")
	for _, operation := range []struct {
		name string
		call func(string) error
	}{
		{"StaleSpecs", func(target string) error { _, err := project.StaleSpecs(root, target); return err }},
		{"StampMissing", func(target string) error { return project.StampMissing(root, target) }},
	} {
		t.Run(operation.name+"_ambiguous", func(t *testing.T) {
			err := operation.call("target")
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C13" {
				t.Errorf("ambiguous slug = %v; want C13", err)
			}
		})
	}
}
package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_040_CheckConfigurationCauses(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed", func(t *testing.T, root string) {
			put(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			_, cause := project.Load(root)
			if cause == nil {
				t.Fatal("Load accepted invalid configuration")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C01 for configuration cause %q", cause)
		})
	}
}

func Test_PRJ_041_CheckGraphCauses(t *testing.T) {
	t.Run("current", func(t *testing.T) {
		root := fixture(t)
		put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		snapshot, err := project.Load(root)
		if err != nil {
			t.Fatal(err)
		}
		cause := project.ValidateGraph(snapshot.Modules)
		if cause == nil {
			t.Fatal("ValidateGraph accepted missing dependency")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for current graph cause %q", cause)
	})
	t.Run("projected", func(t *testing.T) {
		root := fixture(t)
		queued(t, root, "0001-invalid.md", "set:\n  depends_on: [missing]\n", "Needed.\n", "", "", "")
		_, cause := project.Project(root, "")
		if cause == nil {
			t.Fatal("Project accepted invalid projected graph")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for projected graph cause %q", cause)
	})
}

func Test_PRJ_042_CheckDocumentAndItemCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"missing_module", "load", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
		{"malformed_change", "queue", func(t *testing.T, root string) {
			put(t, root, "changes/0001-bad.md", "---\nmodule: [broken\n---\n")
		}},
		{"malformed_item", "queue", func(t *testing.T, root string) {
			queued(t, root, "0001-bad-item.md", "", "Needed.\n", "### new:item [broken\nItem MUST work.\n", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Queue(root)
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid document or item")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C03" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C03 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_043_CheckMetadataAndSetCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"metadata_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [alpha]\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n")
		}},
		{"zone_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: 7\n---\n# Alpha\n\n## Specs\n")
		}},
		{"identity_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  name: renamed\n", "Needed.\n", "", "", "")
		}},
		{"unknown_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  surprise: true\n", "Needed.\n", "", "", "")
		}},
		{"set_value_type", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  interface: 7\n", "Needed.\n", "", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Project(root, "")
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid metadata or set")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C12" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C12 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_044_CheckQueueIdentityCause(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-repeat.md", "", "Needed.\n", "", "", "")
	queued(t, root, "0002-repeat.md", "", "Needed.\n", "", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate slug")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C13" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C13 for queue identity cause %q", cause)
}

func Test_PRJ_045_CheckPendingModuleTargetCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0001-bad.md", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted unknown change module")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C07" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C07 for pending module cause %q", cause)
}

func Test_PRJ_013_CheckDuplicateWithinEditSection(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-duplicate.md", "", "Needed.\n", "### new:local\nLocal MUST work.\n\n### new:local\nDuplicate MUST fail.\n", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate Added identity")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C05" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C05 for duplicate Added identity cause %q", cause)
}
package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits(t *testing.T) {
	const added = "### new:feature\nThe feature MUST work.\n"
	baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
	for _, tc := range []struct {
		name, front, added, modified, removed string
		wantC07            bool
	}{
		{"eval_fix_with_fixes", "eval_fix: [ALP-001]\nfixes: [ALP-001]\n", "", "", "", true},
		{"eval_fix_with_added", "eval_fix: [ALP-001]\n", added, "", "", true},
		{"eval_fix_with_modified", "eval_fix: [ALP-001]\nbases:\n  ALP-001: " + baseHash + "\n", "", "### ALP-001\nAlpha MUST change.\n", "", true},
		{"eval_fix_with_removed", "eval_fix: [ALP-001]\nbases:\n  ALP-001: " + baseHash + "\n", "", "", "### ALP-001\n", true},
		{"eval_fix_only", "eval_fix: [ALP-001]\n", "", "", "", false},
		{"fixes_only", "fixes: [ALP-001]\n", "", "", "", false},
		{"added_only", "", added, "", "", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
			queued(t, root, "0001-check.md", tc.front, "Users need the change.\n", tc.added, tc.modified, tc.removed)

			var errors []project.Diagnostic
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning {
					errors = append(errors, diagnostic)
				}
			}
			if tc.wantC07 {
				if len(errors) != 1 || errors[0].Code != "C07" {
					t.Errorf("mixed change error diagnostics = %#v; want only error C07", errors)
				}
			} else if len(errors) != 0 {
				t.Errorf("valid change error diagnostics = %#v; want none", errors)
			}
		})
	}
}

exec
/bin/zsh -lc 'cat packages/project/evals/spec_history_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat SYSTEM.md packages/project/SPEC.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/queue_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/stale_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/revision_snapshot_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 685ms:
package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func revisionGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func revisionCommit(t *testing.T, root, subject string) string {
	t.Helper()
	revisionGit(t, root, "add", "-A")
	revisionGit(t, root, "-c", "user.name=Fireweed Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-qm", subject)
	return revisionGit(t, root, "rev-parse", "HEAD")
}

func historicalProject(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "fireweed.yaml", `version: 1
modules: ["packages/*"]
module_defaults:
  protected: [guard/**]
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
  preserve: [data/**]
  disposable: [internal/**]
`)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nHistorical alpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nHistorical beta behavior.\n")
	put(t, root, "extras/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: []\n---\n# Gamma\n\n## Specs\n\n### GAM-001\nHistorical gamma behavior.\n")
	recorded := revisionCommit(t, root, "record original project")

	put(t, root, "fireweed.yaml", "version: 2\nmodules: [\"components/*\"]\nmodule_defaults:\n  evals: [later/**]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: later-alpha\nid_prefix: LAL\n---\n# Later\n\n## Specs\n")
	put(t, root, "components/next/SPEC.md", "---\nname: next\nid_prefix: NXT\n---\n# Next\n\n## Specs\n")
	revisionCommit(t, root, "change config glob and modules")

	put(t, root, "fireweed.yaml", "version: 3\nmodules: [\"dirty/*\"]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: dirty-alpha\nid_prefix: DAL\n---\n# Dirty\n\n## Specs\n")
	put(t, root, "dirty/only/SPEC.md", "---\nname: dirty\nid_prefix: DRT\n---\n# Dirty\n\n## Specs\n")
	return root, recorded
}

func Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Config["version"] != 1 || !reflect.DeepEqual(got.Config["modules"], []any{"packages/*"}) {
		t.Errorf("recorded configuration = %#v", got.Config)
	}
	wantDefaults := map[string]any{
		"protected":   []any{"guard/**"},
		"evals":       []any{"evals/**"},
		"interface":   []any{"api.go"},
		"append_only": []any{"ledger/**"},
		"preserve":    []any{"data/**"},
		"disposable":  []any{"internal/**"},
	}
	if !reflect.DeepEqual(got.Config["module_defaults"], wantDefaults) {
		t.Errorf("recorded module_defaults = %#v; want %#v", got.Config["module_defaults"], wantDefaults)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("recorded modules = %#v; want alpha and beta only", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "Historical alpha behavior.") {
		t.Errorf("recorded alpha document = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) || !strings.Contains(beta.Document.Body, "Historical beta behavior.") {
		t.Errorf("recorded beta document = %#v", beta)
	}
}

func Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 {
		t.Fatalf("recorded modules = %#v", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if !reflect.DeepEqual(alpha.Zones.Protected, []string{"guard/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) ||
		!reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"ledger/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) {
		t.Errorf("recorded inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone = %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("recorded default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs(t *testing.T) {
	root, recorded := historicalProject(t)
	canonical, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := project.LoadAt(root+string(filepath.Separator)+".", recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Root != canonical {
		t.Errorf("Root = %q; want canonical %q", got.Root, canonical)
	}
	dirs := make(map[string]string)
	for _, module := range got.Modules {
		dirs[module.Name] = module.Dir
	}
	if !reflect.DeepEqual(dirs, map[string]string{"alpha": "packages/alpha", "beta": "packages/beta"}) {
		t.Errorf("module directories = %#v", dirs)
	}
}

func Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs(t *testing.T) {
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "README.md", "repository without fireweed.yaml\n")
	missingConfig := revisionCommit(t, root, "missing config")
	put(t, root, "fireweed.yaml", "modules: [broken\n")
	malformedConfig := revisionCommit(t, root, "malformed config")
	put(t, root, "fireweed.yaml", "modules: [\"modules/*\"]\n")
	put(t, root, "modules/ghost/README.md", "directory without SPEC.md\n")
	missingSpec := revisionCommit(t, root, "missing module document")
	for _, tc := range []struct{ name, revision string }{
		{"invalid_revision", "not-a-revision"},
		{"missing_recorded_config", missingConfig},
		{"malformed_recorded_config", malformedConfig},
		{"missing_recorded_SPEC", missingSpec},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := project.LoadAt(root, tc.revision); err == nil {
				t.Errorf("LoadAt accepted %s (%s)", tc.name, tc.revision)
			}
		})
	}
}

 succeeded in 694ms:
package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

func historyArchive(t *testing.T, root, path, module, motivation, added, modified, removed, decisions, removedDecisions, provenance string) {
	t.Helper()
	put(t, root, path, fmt.Sprintf("---\nmodule: %s\nphase: applied\n---\n\n## Motivation\n%s## Added\n%s## Modified\n%s## Removed\n%s## Decisions\n%s## Removed decisions\n%s## Provenance\n%s", module, motivation, added, modified, removed, decisions, removedDecisions, provenance))
}

func historyProject(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	revisionGit(t, root, "init", "-q")
	historyArchive(t, root, "changes/archive/0009-added.md", "alpha",
		"Why the behavior exists.\nSecond motivation line.\n\n",
		"### ALP-004\nAnother behavior.\n### ALP-002\nRetained behavior.\n", "", "",
		"Choose the stable form.\nKeep this second decision.\n\n",
		"Prior choice was dropped.\n\n",
		"- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n")
	revisionCommit(t, root, "archive newest implementation date first")
	historyArchive(t, root, "changes/archive/0001-retag.md", "alpha", "Retag the behavior.\n", "",
		"### ALP-002 [review]\nRetained behavior.\n", "", "", "", "- impl 2024-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0003-equal.md", "alpha", "First equal date.\n", "",
		"### ALP-002\nEqual date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0004-equal.md", "alpha", "Second equal date.\n", "",
		"### ALP-002\nAnother equal date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0005-bad-date.md", "alpha", "Invalid date.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- impl not-a-date agent=fixture\n")
	historyArchive(t, root, "changes/archive/0006-no-impl.md", "alpha", "No implementation line.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- evals 2030-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0010-other-id.md", "alpha", "Other specs.\n",
		"### ALP-005\nFifth behavior.\n### ALP-003\nThird behavior.\n",
		"### ALP-003 [review]\nThird behavior retagged.\n", "", "", "", "")
	historyArchive(t, root, "changes/archive/0011-removed-only.md", "alpha", "Only a removal.\n", "", "",
		"### ALP-002\n", "", "", "- impl 2031-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0012-beta.md", "beta", "Beta history.\n",
		"### BET-001\nBeta behavior.\n", "", "", "", "", "- impl 2032-01-01T00:00:00Z agent=fixture\n")
	queued(t, root, "0013-pending.md", "", "Not archived.\n", "### ALP-002\nPending behavior.\n", "", "")
	revisionCommit(t, root, "archive older dates and pending change later")
	return root
}

func historyPaths(entries []project.SpecHistoryEntry) []string {
	paths := make([]string, len(entries))
	for i, entry := range entries {
		paths[i] = entry.Path
	}
	return paths
}

func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	want := map[string]bool{
		"changes/archive/0009-added.md":    true,
		"changes/archive/0001-retag.md":    true,
		"changes/archive/0003-equal.md":    true,
		"changes/archive/0004-equal.md":    true,
		"changes/archive/0005-bad-date.md": true,
		"changes/archive/0006-no-impl.md":  true,
	}
	if len(entries) != len(want) {
		t.Fatalf("selected archives = %#v; want paths %#v", historyPaths(entries), want)
	}
	for _, entry := range entries {
		if !want[entry.Path] || entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, []string{"ALP-002"}) {
			t.Errorf("unexpected path/module/filtered IDs = %q, %q, %#v", entry.Path, entry.Module, entry.IDs)
		}
		delete(want, entry.Path)
	}
	if len(want) != 0 {
		t.Errorf("missing archives: %#v", want)
	}
}

func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("archive order/selection = %#v; want %#v", got, wantPaths)
	}
	wantDates := []time.Time{
		time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		{}, {},
	}
	for i, entry := range entries {
		if !entry.ImplementationDate.Equal(wantDates[i]) {
			t.Errorf("%s implementation date = %s; want %s", entry.Path, entry.ImplementationDate, wantDates[i])
		}
	}
}

func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("missing archived ALP-002 history")
	}
	want := project.SpecHistoryEntry{
		Path: "changes/archive/0009-added.md", Module: "alpha", IDs: []string{"ALP-002"},
		Motivation:         "Why the behavior exists.\nSecond motivation line.\n\n",
		Decisions:          "Choose the stable form.\nKeep this second decision.\n\n",
		RemovedDecisions:   "Prior choice was dropped.\n\n",
		Provenance:         "- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n",
		ImplementationDate: time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(entries[0], want) {
		t.Errorf("entry = %#v; want %#v", entries[0], want)
	}
}

func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "")
	if err != nil {
		t.Fatalf("SpecHistory module-wide: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
		"changes/archive/0010-other-id.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("module-wide archive order/selection = %#v; want %#v", got, wantPaths)
	}
	for _, entry := range entries {
		wantIDs := []string{"ALP-002"}
		switch entry.Path {
		case "changes/archive/0009-added.md":
			wantIDs = []string{"ALP-002", "ALP-004"}
		case "changes/archive/0010-other-id.md":
			wantIDs = []string{"ALP-003", "ALP-005"}
		}
		if entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, wantIDs) {
			t.Errorf("%s module/unique IDs = %q, %#v; want alpha, %#v", entry.Path, entry.Module, entry.IDs, wantIDs)
		}
	}
	emptyRoot := fixture(t)
	empty, err := project.SpecHistory(emptyRoot, "beta", "")
	if err != nil || len(empty) != 0 {
		t.Errorf("module without archives = %#v, %v; want empty slice", empty, err)
	}
}

func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
	root := historyProject(t)
	archived, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil || len(archived) == 0 {
		t.Errorf("archived ID absent from current Specs = %#v, %v", archived, err)
	}
	currentOnly, err := project.SpecHistory(root, "alpha", "ALP-001")
	if err != nil || len(currentOnly) != 0 {
		t.Errorf("current ID with no archive edits = %#v, %v; want empty slice", currentOnly, err)
	}
	for _, tc := range []struct{ module, id, rejected string }{
		{"absent-module", "", "absent-module"},
		{"alpha", "ALP-999", "ALP-999"},
		{"alpha", "BET-001", "BET-001"},
	} {
		_, err := project.SpecHistory(root, tc.module, tc.id)
		if err == nil || !strings.HasPrefix(err.Error(), "U02") || !strings.Contains(err.Error(), tc.rejected) {
			t.Errorf("lookup module %q ID %q = %v; want U02 naming %q", tc.module, tc.id, err, tc.rejected)
		}
	}
}

func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
	for _, tc := range []struct{ name, config string }{
		{"missing", ""},
		{"malformed", "modules: [broken\n"},
	} {
		t.Run("config_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			if tc.name == "missing" {
				if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
			} else {
				put(t, root, "fireweed.yaml", tc.config)
			}
			_, err := project.SpecHistory(root, "alpha", "")
			if err == nil || !strings.HasPrefix(err.Error(), "C01") {
				t.Errorf("invalid config = %v; want C01", err)
			}
		})
	}
	for _, tc := range []struct{ name, document string }{
		{"document", "---\nmodule: [broken\n---\n## Added\n### ALP-002\nBehavior.\n"},
		{"added", "---\nmodule: alpha\n---\n## Added\n### ALP-002 [broken\nBehavior.\n"},
		{"modified", "---\nmodule: alpha\n---\n## Modified\n### ALP-002 [broken\nBehavior.\n"},
	} {
		t.Run("archive_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			path := "changes/archive/0001-invalid.md"
			put(t, root, path, tc.document)
			_, err := project.SpecHistory(root, "alpha", "ALP-002")
			if err == nil || !strings.HasPrefix(err.Error(), "C03") || !strings.Contains(err.Error(), path) {
				t.Errorf("malformed archive = %v; want C03 naming %s", err, path)
			}
		})
	}
}

 succeeded in 694ms:
# Independent project reconstruction policy

This policy governs the independent, input-only project reconstruction workspace.
Rebuild project solely from its retained SPEC.md, Decisions, public interfaces,
model types and evals, plus supplied contracts, gitstore and assets interfaces.
The Go module path is github.com/lurkshark/fireweed, as shown by those interfaces.

## Retained repository conventions

The product never contacts the network or runs fetch or push. Read operations
never run hooks or agents and never change files, the index or refs; output is
deterministic. These are repository conventions, not new module spec items.
Retained api.go facades and model types remain interfaces; behavior belongs
under internal/. Never import the public project parent from its own internal
implementation; shared data belongs in model or contracts. Keep dependencies
minimal and add no native objects, assembly, cgo or embedded executables.

## Input isolation and verification

Read only the supplied input workspace and your newly written files. Do not read
outside it, old implementation, source history, saved attempts, other-module
internals, product binaries, bootstrap, external checkouts or prior drill output.
Do not query the invoking executable for project behavior, help, recipes,
playbooks, explanations or other content. Do not run fireweed authoring or
lifecycle commands in this input workspace. Do not commit or invoke other agents.

Retained inputs must remain byte-unchanged. Write reconstruction only under
packages/project/internal/**, and write the requested DRILL_NOTES.md with
Contradicted Decisions and Gaps sections. Do not create root manifests, copy
working dependency implementations, or repair dependent code. Remove temporary
probe files before finishing. Every behavioral gap names a nearest spec ID (or
no spec) and its target module; do not author fixes from gaps.

This input deliberately lacks full build manifests, dependency implementations
and verification hooks. Supplied dependency facades can reference absent internal
packages; do not search for those implementations. The separate native verifier
supplies the full project, actual dependencies and dependents and runs project's
own evals plus every transitive dependent's enforced evals and coverage. Inability
to run that full suite here is an input-only verification limitation, not a reason
to copy forbidden material or change retained inputs. Describe it accurately if
relevant; the native verifier will provide actual results.

Use ordinary local Go and Git tools for bounded owned-fixture probes when useful.
Keep the current input tree's .fireweed/bin first on worker PATH and use
GOCACHE=/tmp/fireweed-rebuild-cache. No controller binary is supplied as an input.
The orchestration config selects exactly codex exec --ephemeral
--ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high"
-s workspace-write {prompt}; do not substitute model or effort.
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
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

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
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash with diagnostic C08.

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

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module.

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens.

### PRJ-027
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### PRJ-028
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### PRJ-029
Queue MUST reject a malformed pending change document.

### PRJ-030
Queue MUST reject a pending change naming an unknown module.

### PRJ-031
Queue MUST reject pending changes that share a slug.

### PRJ-032
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### PRJ-033
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### PRJ-034
Project MUST reject identity or unknown fields in set.

### PRJ-035
Project MUST reject an invalid projected module graph.

### PRJ-036
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

### PRJ-037
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### PRJ-038
Stamp MUST retain other frontmatter values and human section contents.

### PRJ-039
Stamp MUST succeed on an Added-only change without assigning spec IDs.

### PRJ-040
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### PRJ-041
Check MUST report invalid current or projected module graphs with diagnostic C02.

### PRJ-042
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### PRJ-043
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### PRJ-044
Check MUST report pending queue identity errors with diagnostic C13.

### PRJ-045
Check MUST report invalid pending change module targets with diagnostic C07.

### PRJ-046
Check MUST report error C07 for a pending change containing eval_fix together
with fixes or any Added, Modified or Removed spec edits.

### PRJ-047
LoadAt MUST return the selected Git revision's project configuration and modules, using its recorded fireweed.yaml module globs and module SPEC.md documents rather than working-tree contents.

### PRJ-048
LoadAt MUST apply the selected revision's module_defaults to omitted zone fields while preserving explicitly empty module zone lists.

### PRJ-049
LoadAt MUST set Snapshot.Root to the canonical repository root and Module.Dir to each module's repository-relative directory.

### PRJ-050
LoadAt MUST return an error when the revision cannot be resolved, its recorded fireweed.yaml is missing or malformed, or a matched recorded module directory lacks a readable SPEC.md.

### PRJ-051
For a nonempty ID, SpecHistory MUST return the target module's archived changes whose Added or Modified items name that ID, including tag-only modifications.
Pending changes and archives belonging to other modules or unrelated spec IDs MUST not appear in that selection.

### PRJ-052
Every history entry MUST identify its archive path and module, the matching IDs, its Motivation, added and removed Decisions, and its complete Provenance section.

### PRJ-053
SpecHistory MUST order entries by their latest valid implementation Provenance timestamp, newest first, with missing or invalid implementation dates after dated entries.
Equal-date or undated entries MUST use archive-path lexical order as a total-order fallback.

### PRJ-054
With an empty ID filter, SpecHistory MUST return each target-module archive containing Added or Modified spec edits once, with its unique affected IDs in lexical order.
A module with no matching archives MUST return an empty slice.

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching archive edits MUST return U02 with the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce a C01 error.
Malformed archived change documents or malformed Added/Modified item syntax MUST produce a C03 error identifying the archive path.

### PRJ-057
Stamp MUST return an error whose contracts ErrorCode is C08 when the change's bases frontmatter is not a mapping.

### PRJ-058
Stamp MUST return an error whose contracts ErrorCode is U02, with the requested change as subject, when no pending change matches it.

### PRJ-059
Load MUST return an error preserving the contracts ErrorCode C03 or C05 when a module SPEC.md fails to parse.

### PRJ-060
StaleSpecs MUST return the target change's Modified and Removed items whose recorded bases differ from the corresponding preceding projected spec hashes, in edit order, with ID, RecordedHash and Current text.
Items without a bases entry MUST be omitted.

### PRJ-061
StaleSpecs MUST recover Before by projecting the spec immediately before the target change at the Git commit that last changed that change file's bases mapping.
It MUST set Recovered true only when the recovered item's normalized SpecHash equals RecordedHash.

### PRJ-062
When matching stamped history cannot be recovered, StaleSpecs MUST return the stale item's Current text and RecordedHash with Recovered false and empty Before.
Unavailable history MUST NOT prevent reporting other stale items.

### PRJ-063
StaleSpecs MUST leave repository files, index and refs unchanged.
It MUST return an empty slice when no recorded bases are stale.

### PRJ-064
StampMissing MUST add canonical preceding-projection hashes only for Modified or Removed IDs absent from bases, preserving every existing entry including stale hashes.
It MUST retain other frontmatter values and human section contents.

### PRJ-065
StampMissing MUST reject a nonmapping bases value or an edited ID absent from the preceding projection with C08 and preserve the target file on rejection.

### PRJ-066
StaleSpecs and StampMissing MUST accept a unique pending slug, queued basename or repository-relative queued path.
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.

 succeeded in 694ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func queued(t *testing.T, root, filename, front, motivation, added, modified, removed string) {
	t.Helper()
	put(t, root, "changes/"+filename, "---\nmodule: alpha\n"+front+"---\n\n## Motivation\n"+motivation+"\n## Interface\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\n\n## Notes\nHuman note.\n")
}

func moduleByName(t *testing.T, snapshot project.Snapshot, name string) project.Module {
	t.Helper()
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module
		}
	}
	t.Fatalf("module %q absent from projection", name)
	return project.Module{}
}

func specItems(t *testing.T, module project.Module) []contracts.Item {
	t.Helper()
	section, ok := contracts.Section(module.Document.Body, "Specs")
	if !ok {
		t.Fatalf("Specs section absent from %s", module.Name)
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		t.Fatalf("parse projected %s specs: %v", module.Name, err)
	}
	return items
}

func findItem(t *testing.T, items []contracts.Item, id string) contracts.Item {
	t.Helper()
	for _, item := range items {
		if item.ID == id {
			return item
		}
	}
	t.Fatalf("spec %s absent from %#v", id, items)
	return contracts.Item{}
}

func Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0002-zed.md", "set:\n  interface: [next.go]\n", "A queue item is needed.\n", "### new:feature\nFeature MUST work.\n", "### ALP-001\nAlpha MUST change.\n", "### ALP-002\n")
	queued(t, root, "0002-ace.md", "", "Another queue item is needed.\n", "", "", "")
	queued(t, root, "0001-first.md", "", "The first queue item is needed.\n", "", "", "")
	put(t, root, "changes/archive/0000-old.md", "archived material is not a pending document")
	changes, err := project.Queue(root)
	if err != nil {
		t.Fatalf("Queue: %v", err)
	}
	if len(changes) != 3 {
		t.Fatalf("queue = %#v; want three pending changes", changes)
	}
	for i, want := range []struct {
		number     int
		slug, path string
	}{
		{1, "first", "changes/0001-first.md"},
		{2, "ace", "changes/0002-ace.md"},
		{2, "zed", "changes/0002-zed.md"},
	} {
		got := changes[i]
		if got.Number != want.number || got.Slug != want.slug || got.Path != want.path || got.Module != "alpha" {
			t.Errorf("change %d = %#v; want %#v with module alpha", i, got, want)
		}
	}
	last := changes[2]
	if !strings.Contains(last.Document.Body, "A queue item is needed.") || last.Document.Frontmatter["module"] != "alpha" || last.Document.Frontmatter["set"] == nil {
		t.Errorf("change document/frontmatter lost: %#v", last.Document)
	}
	if len(last.Added) != 1 || last.Added[0].ID != "new:feature" || !strings.Contains(last.Added[0].Body, "Feature MUST work.") || len(last.Modified) != 1 || last.Modified[0].ID != "ALP-001" || !strings.Contains(last.Modified[0].Body, "Alpha MUST change.") || len(last.Removed) != 1 || last.Removed[0].ID != "ALP-002" {
		t.Errorf("parsed edits = added %#v, modified %#v, removed %#v", last.Added, last.Modified, last.Removed)
	}
}

func Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments(t *testing.T) {
	for _, tc := range []struct{ id, name, body string }{
		{"PRJ-029", "malformed", "---\nmodule: [broken\n---\n"},
		{"PRJ-030", "unknown_module", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "changes/0001-bad.md", tc.body)
			if _, err := project.Queue(root); err == nil {
				t.Error("Queue accepted invalid change")
			}
		})
	}
}

func Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:first\nFirst MUST work.\n", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:second\nSecond MUST work.\n", "", "")
	for _, target := range []string{"first", "0001-first.md", "changes/0001-first.md"} {
		projected, err := project.Project(root, target)
		if err != nil {
			t.Errorf("Project(%q): %v", target, err)
			continue
		}
		items := specItems(t, moduleByName(t, projected, "alpha"))
		if len(items) != 2 || findItem(t, items, "new:first").ID == "" {
			t.Errorf("Project(%q) did not stop at first change: %#v", target, items)
		}
		if err := project.Stamp(root, target); err != nil {
			t.Errorf("Stamp(%q): %v", target, err)
		}
	}
	if _, err := project.Project(root, "missing"); err == nil {
		t.Error("Project accepted missing target")
	}
	if err := project.Stamp(root, "missing"); err == nil {
		t.Error("Stamp accepted missing target")
	}
	put(t, root, "changes/archive/0000-archived.md", "archived document")
	if _, err := project.Project(root, "changes/archive/0000-archived.md"); err == nil {
		t.Error("Project accepted archived target")
	}
	queued(t, root, "0003-first.md", "", "Another first.\n", "", "", "")
	if _, err := project.Queue(root); err == nil {
		t.Error("Queue accepted ambiguous slug")
	}
	if _, err := project.Project(root, "first"); err == nil {
		t.Error("Project accepted ambiguous slug")
	}
	if err := project.Stamp(root, "first"); err == nil {
		t.Error("Stamp accepted ambiguous slug")
	}
}

func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:shared\nFirst change MUST retain its own text.\n", "### ALP-001 [review]\nAlpha MUST have revised full text.\n", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:shared\nSecond change MUST retain different text.\n", "", "### ALP-001\n")
	first, err := project.Project(root, "first")
	if err != nil {
		t.Fatalf("Project first: %v", err)
	}
	firstItems := specItems(t, moduleByName(t, first, "alpha"))
	if len(firstItems) != 2 || findItem(t, firstItems, "ALP-001").Tag != "review" || findItem(t, firstItems, "ALP-001").Body != "Alpha MUST have revised full text.\n\n" || findItem(t, firstItems, "new:shared").Body != "First change MUST retain its own text.\n\n" {
		t.Errorf("first projection = %#v", firstItems)
	}
	all, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project whole queue: %v", err)
	}
	section, ok := contracts.Section(moduleByName(t, all, "alpha").Document.Body, "Specs")
	if !ok {
		t.Fatal("Specs section absent from whole-queue projection")
	}
	if strings.Count(section, "### new:shared\n") != 2 ||
		!strings.Contains(section, "### new:shared\nFirst change MUST retain its own text.\n") ||
		!strings.Contains(section, "### new:shared\nSecond change MUST retain different text.\n") ||
		strings.Contains(section, "ALP-001") {
		t.Errorf("whole-queue projection lost local handles or retained removed ID: %q", section)
	}
}

func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
	queued(t, root, "0001-settings.md", "set:\n  depends_on: [beta]\n  interface: [next.go]\n  evals: [checks/**]\n  disposable: [generated/**]\n  preserve: [records/**]\n  append_only: [migrations/**]\n  boundaries_opt_out: reviewed boundary\n  hooks:\n    test:\n      command: [go, test, ./...]\n", "Settings are needed.\n", "", "", "")
	got, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project settings: %v", err)
	}
	alpha := moduleByName(t, got, "alpha")
	wantLists := map[string][]string{
		"depends_on": {"beta"}, "interface": {"next.go"}, "evals": {"checks/**"},
		"disposable": {"generated/**"}, "preserve": {"records/**"}, "append_only": {"migrations/**"},
	}
	for key, want := range wantLists {
		value, ok := alpha.Document.Frontmatter[key].([]any)
		if !ok || len(value) != 1 || value[0] != want[0] {
			t.Errorf("set %s = %#v; want %q", key, alpha.Document.Frontmatter[key], want)
		}
	}
	if !reflect.DeepEqual(alpha.DependsOn, []string{"beta"}) || !reflect.DeepEqual(alpha.Zones.Interface, []string{"next.go"}) || !reflect.DeepEqual(alpha.Zones.Evals, []string{"checks/**"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"generated/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"records/**"}) || !reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"migrations/**"}) {
		t.Errorf("projected module fields not updated: %#v", alpha)
	}
	wantHooks := map[string]any{"test": map[string]any{"command": []any{"go", "test", "./..."}}}
	if alpha.Document.Frontmatter["boundaries_opt_out"] != "reviewed boundary" || !reflect.DeepEqual(alpha.Document.Frontmatter["hooks"], wantHooks) {
		t.Errorf("projected scalar/map settings lost: %#v", alpha.Document.Frontmatter)
	}
	for _, tc := range []struct{ id, name, settings string }{
		{"PRJ-034", "identity_name", "name: renamed"},
		{"PRJ-034", "identity_prefix", "id_prefix: NEW"},
		{"PRJ-034", "unknown_field", "surprise: true"},
		{"PRJ-035", "invalid_graph", "depends_on: [missing]"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			badRoot := fixture(t)
			queued(t, badRoot, "0001-invalid.md", "set:\n  "+tc.settings+"\n", "Needed.\n", "", "", "")
			if _, err := project.Project(badRoot, ""); err == nil {
				t.Errorf("Project accepted %s", tc.name)
			}
		})
	}
}

func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-revise.md", "", "Earlier change.\n", "", "### ALP-001 [review]\nAlpha MUST reflect the earlier revision.\n", "")
	queued(t, root, "0002-target.md", "custom: author-value\n", "Keep this human motivation.\n", "", "### ALP-001\nAlpha MUST reflect the later revision.\n", "")
	wantHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Tag: "review", Body: "Alpha MUST reflect the earlier revision.\n"})
	if err := project.Stamp(root, "changes/0002-target.md"); err != nil {
		t.Fatalf("Stamp: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(root, "changes/0002-target.md"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("stamped document: %v", err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok || bases["ALP-001"] != wantHash {
		t.Errorf("bases = %#v; want preceding hash %s", doc.Frontmatter["bases"], wantHash)
	}
	if doc.Frontmatter["custom"] != "author-value" || !strings.Contains(doc.Body, "Keep this human motivation.") || !strings.Contains(doc.Body, "Alpha MUST reflect the later revision.") || !strings.Contains(doc.Body, "Human note.") {
		t.Errorf("Stamp lost author content: %#v", doc)
	}
	queued(t, root, "0003-remove.md", "", "Remove the item.\n", "", "", "### ALP-001\n")
	if err := project.Stamp(root, "0003-remove.md"); err != nil {
		t.Fatalf("Stamp Removed: %v", err)
	}
	removedBytes, err := os.ReadFile(filepath.Join(root, "changes/0003-remove.md"))
	if err != nil {
		t.Fatal(err)
	}
	removedDoc, err := contracts.ParseDocument(string(removedBytes))
	if err != nil {
		t.Fatal(err)
	}
	removedWant := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST reflect the later revision.\n"})
	removedBases, ok := removedDoc.Frontmatter["bases"].(map[string]any)
	if !ok || removedBases["ALP-001"] != removedWant {
		t.Errorf("Removed base = %#v; want %s", removedDoc.Frontmatter["bases"], removedWant)
	}
}

func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-added.md", "", "Added item.\n", "### new:thing\nThing MUST work.\n", "", "")
	if err := project.Stamp(root, "added"); err != nil {
		t.Fatalf("Stamp Added-only: %v", err)
	}
	bytes, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bytes), "### new:thing") || strings.Contains(string(bytes), "### ALP-002") {
		t.Errorf("Stamp assigned an ID to Added-only change: %s", bytes)
	}
	queued(t, root, "0002-invalid.md", "", "Invalid edit.\n", "", "### ALP-999\nMissing item MUST change.\n", "")
	if err := project.Stamp(root, "invalid"); err == nil {
		t.Error("Stamp accepted Modified ID absent from preceding projection")
	}
	removeRoot := fixture(t)
	queued(t, removeRoot, "0001-invalid-remove.md", "", "Invalid removal.\n", "", "", "### ALP-998\n")
	if err := project.Stamp(removeRoot, "invalid-remove"); err == nil {
		t.Error("Stamp accepted Removed ID absent from preceding projection")
	}
}

func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
	validRoot := fixture(t)
	put(t, validRoot, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	queued(t, validRoot, "0001-valid.md", "bases:\n  ALP-001: "+contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})+"\n", "A user needs the change.\n", "### new:local [review]\nA local addition MUST work.\n", "### ALP-001 [quarantine: flaky fixture]\nAlpha MUST work with revision.\n", "")
	queued(t, validRoot, "0002-also-valid.md", "", "Another user need.\n", "### new:local\nA second local addition MUST work.\n", "", "")
	for _, diagnostic := range project.Check(validRoot) {
		if !diagnostic.Warning {
			t.Errorf("valid spec/change received error diagnostic: %#v", diagnostic)
		}
	}
}

func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
	for _, tc := range []struct {
		id, name, spec, front, motivation, added, modified, removed string
		code, cause                                                 string
	}{
		{"PRJ-008", "missing_motivation", "", "", "\n", "", "", "", "C07", "missing Motivation"},
		{"PRJ-009", "nonnormative_module", "### ALP-001\nAlpha works.\n", "", "Needed.\n", "", "", "", "C04", "non-normative spec body"},
		{"PRJ-009", "nonnormative_change", "", "", "Needed.\n", "### new:item\nAn observation.\n", "", "", "C04", "non-normative spec body"},
		{"PRJ-010", "invalid_review_tag", "", "", "Needed.\n", "### new:item [review: reason]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-010", "invalid_quarantine_tag", "", "", "Needed.\n", "### new:item [quarantine]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-011", "duplicate_module_id", "### ALP-001\nAlpha MUST work.\n\n### ALP-001\nDuplicate MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "duplicate item ID"},
		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
		{"PRJ-013", "duplicate_edit_identity", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "### ALP-001\n", "C05", "duplicate edit identity ALP-001"},
		{"PRJ-014", "undeclared_handle", "", "", "Needed.\n", "### new:local\nLocal MUST work.\n", "### ALP-001\nAlpha MUST refer to new:foreign.\n", "", "C09", "reference to undeclared handle new:foreign"},
		{"PRJ-015", "missing_base", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "missing base for ALP-001"},
		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
		{"PRJ-016", "absent_modified", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "### ALP-999\nAbsent MUST change.\n", "", "C08", "ID ALP-999 absent from preceding projection"},
		{"PRJ-016", "absent_removed", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "", "### ALP-999\n", "C08", "ID ALP-999 absent from preceding projection"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			front := tc.front
			baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
			if tc.name == "duplicate_edit_identity" || tc.name == "undeclared_handle" {
				front = "bases:\n  ALP-001: " + baseHash + "\n"
			}
			if tc.name == "foreign_change_id" {
				put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
				front = "bases:\n  BET-001: " + contracts.SpecHash(contracts.Item{ID: "BET-001", Body: "Beta MUST work.\n"}) + "\n"
			}
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n"+func() string {
				if tc.spec != "" {
					return tc.spec
				}
				return "### ALP-001\nAlpha MUST work.\n"
			}())
			if tc.name == "undeclared_handle" {
				queued(t, root, "0000-other.md", "", "Earlier local handle.\n", "### new:foreign\nEarlier item MUST work.\n", "", "")
			}
			queued(t, root, "0001-check.md", front, tc.motivation, tc.added, tc.modified, tc.removed)
			diagnostics := project.Check(root)
			found := false
			foundCause := false
			foundApplyCause := false
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
					t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
				}
				if !diagnostic.Warning && strings.TrimSpace(diagnostic.Code) != "" && strings.TrimSpace(diagnostic.Message) != "" {
					found = true
				}
				matchesCause := strings.Contains(diagnostic.Message, tc.cause)
				if tc.id == "PRJ-016" {
					matchesCause = diagnostic.Message == tc.cause
				}
				if tc.code != "" && !diagnostic.Warning && diagnostic.Code == tc.code && matchesCause {
					foundCause = true
				}
				if tc.name == "absent_modified" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Modified ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
				if tc.name == "absent_removed" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Removed ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
			}
			if !found {
				t.Errorf("Check did not report error diagnostic with nonempty code and message: %#v", diagnostics)
			}
			if tc.code != "" && !foundCause {
				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
			}
			if tc.id == "PRJ-016" && !foundApplyCause {
				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
			}
			if tc.name == "missing_base" || tc.name == "stale_base" {
				queued(t, root, "0001-check.md", front, "Needed.\n", "", "", "### ALP-001\n")
				removedDiagnostics := project.Check(root)
				found = false
				for _, diagnostic := range removedDiagnostics {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message for Removed entry: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, tc.cause) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not report C08 for Removed entry with %s (%s): %#v", tc.name, tc.cause, removedDiagnostics)
				}
			}
			if tc.name == "nonnormative_change" || tc.name == "invalid_review_tag" {
				front, added, modified, want := "", "", "", ""
				code := "C04"
				if tc.name == "nonnormative_change" {
					front = "bases:\n  ALP-001: " + baseHash + "\n"
					modified = "### ALP-001\nAn observation.\n"
					want = "non-normative spec body"
				} else {
					code = "C06"
					added = "### new:item [other: reason]\nItem MUST work.\n"
					want = "invalid review/quarantine tag"
				}
				queued(t, root, "0001-check.md", front, "Needed.\n", added, modified, "")
				found = false
				for _, diagnostic := range project.Check(root) {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == code && strings.Contains(diagnostic.Message, want) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not reject %s: wanted error diagnostic %s with message containing %q", tc.name, code, want)
				}
			}
		})
	}
}

 succeeded in 698ms:
package evals_test

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func staleSpec(t *testing.T, root, through, id string) contracts.Item {
	t.Helper()
	snapshot, err := project.Project(root, through)
	if err != nil {
		t.Fatal(err)
	}
	return findItem(t, specItems(t, moduleByName(t, snapshot, "alpha")), id)
}

func staleBaseItem(t *testing.T, root, id string) contracts.Item {
	t.Helper()
	snapshot, err := project.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	return findItem(t, specItems(t, moduleByName(t, snapshot, "alpha")), id)
}

func staleRead(t *testing.T, root, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func staleBases(t *testing.T, root, path string) (map[string]any, contracts.Document) {
	t.Helper()
	doc, err := contracts.ParseDocument(string(staleRead(t, root, path)))
	if err != nil {
		t.Fatal(err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok {
		t.Fatalf("bases = %#v, want mapping", doc.Frontmatter["bases"])
	}
	return bases, doc
}

func staleFixture(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nOne MUST work.\n\n### ALP-002\nTwo MUST work.\n\n### ALP-003\nThree MUST work.\n\n### ALP-004\nFour MUST work.\n")
	return root
}

func Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder(t *testing.T) {
	root := staleFixture(t)
	oldOne := contracts.SpecHash(staleSpec(t, root, "", "ALP-001"))
	oldTwo := contracts.SpecHash(staleSpec(t, root, "", "ALP-002"))
	oldThree := contracts.SpecHash(staleSpec(t, root, "", "ALP-003"))
	queued(t, root, "0001-earlier.md", "", "Earlier edits.\n", "", "### ALP-001 [review]\nOne MUST change.\n\n### ALP-002 [review]\nTwo MUST change.\n\n### ALP-003 [review]\nThree MUST change.\n", "")
	queued(t, root, "0002-target.md", "bases:\n  ALP-002: "+oldTwo+"\n  ALP-001: "+oldOne+"\n  ALP-003: "+oldThree+"\n", "Later edits.\n", "", "### ALP-002\nTwo MUST change again.\n\n### ALP-001\nOne MUST change again.\n\n### ALP-004\nFour MUST change without a stamp.\n", "### ALP-003\n")
	got, err := project.StaleSpecs(root, "target")
	if err != nil {
		t.Fatal(err)
	}
	want := []struct{ id, hash, current string }{
		{"ALP-002", oldTwo, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-002"))},
		{"ALP-001", oldOne, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-001"))},
		{"ALP-003", oldThree, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-003"))},
	}
	if len(got) != len(want) {
		t.Fatalf("stale entries = %#v; want %d recorded entries", got, len(want))
	}
	for i, entry := range got {
		if entry.ID != want[i].id || entry.RecordedHash != want[i].hash || entry.Current != want[i].current {
			t.Errorf("entry %d = %#v; want ID %s, hash %s, Current %q", i, entry, want[i].id, want[i].hash, want[i].current)
		}
	}
}

func Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash(t *testing.T) {
	root := staleFixture(t)
	revisionGit(t, root, "init", "-q")
	queued(t, root, "0001-earlier.md", "", "Earlier.\n", "", "### ALP-001 [review]\nOne MUST use old text.  \n\n### ALP-002 [review]\nTwo MUST use old text.\n", "")
	queued(t, root, "0002-target.md", "", "Target.\n", "", "### ALP-001\nOne MUST be edited.\n", "### ALP-002\n")
	revisionCommit(t, root, "initial unstamped queue")
	before := staleSpec(t, root, "earlier", "ALP-001")
	if before.Tag != "review" || !strings.Contains(before.Body, "old text.  ") {
		t.Fatalf("fixture lost tag or trailing spaces: %#v", before)
	}
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	path := "changes/0002-target.md"
	bases, doc := staleBases(t, root, path)
	if bases["ALP-001"] != contracts.SpecHash(before) {
		t.Fatalf("stamped hash %v differs from normalized tagged hash %s", bases["ALP-001"], contracts.SpecHash(before))
	}
	// A second recorded hash is intentionally wrong. Recovery must be per item.
	bases["ALP-002"] = "deadbeef"
	text, err := contracts.RenderDocument(doc)
	if err != nil {
		t.Fatal(err)
	}
	put(t, root, path, text)
	revisionCommit(t, root, "last bases mapping change")
	queued(t, root, "0001-earlier.md", "", "Earlier.\n", "", "### ALP-001 [quarantine: changed]\nOne MUST use current text.\n\n### ALP-002\nTwo MUST use current text.\n", "")
	// This later target-file commit changes only prose, not the bases mapping.
	put(t, root, path, strings.Replace(string(staleRead(t, root, path)), "Target.", "Target with later prose.", 1))
	revisionCommit(t, root, "later unrelated change to target file")
	currentOne := staleSpec(t, root, "earlier", "ALP-001")
	currentTwo := staleSpec(t, root, "earlier", "ALP-002")
	got, err := project.StaleSpecs(root, "target")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 {
		t.Fatalf("stale entries = %#v; want two", got)
	}
	if got[0].ID != "ALP-001" || got[0].RecordedHash != contracts.SpecHash(before) || !got[0].Recovered || got[0].Before != contracts.RenderItem(before) || got[0].Current != contracts.RenderItem(currentOne) {
		t.Errorf("recovered tagged item = %#v; want historical pre-target item %#v and current %#v", got[0], before, currentOne)
	}
	if got[1].ID != "ALP-002" || got[1].RecordedHash != "deadbeef" || got[1].Recovered || got[1].Before != "" || got[1].Current != contracts.RenderItem(currentTwo) {
		t.Errorf("mismatched historical hash = %#v; want current text without Before", got[1])
	}
}

func Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem(t *testing.T) {
	root := staleFixture(t) // No Git repository or recorded target history.
	queued(t, root, "0001-target.md", "bases:\n  ALP-001: deadbeef\n  ALP-002: 12345678\n", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
	got, err := project.StaleSpecs(root, "target")
	if err != nil || len(got) != 2 {
		t.Fatalf("unavailable history = %#v, %v; want both stale items", got, err)
	}
	for i, want := range []struct{ id, hash string }{{"ALP-001", "deadbeef"}, {"ALP-002", "12345678"}} {
		if got[i].ID != want.id || got[i].RecordedHash != want.hash || got[i].Recovered || got[i].Before != "" || got[i].Current != contracts.RenderItem(staleBaseItem(t, root, want.id)) {
			t.Errorf("entry %d = %#v; want unrecovered %s", i, got[i], want.id)
		}
	}
}

func Test_PRJ_063_StaleSpecsEmptyAndReadOnly(t *testing.T) {
	root := staleFixture(t)
	revisionGit(t, root, "init", "-q")
	queued(t, root, "0001-target.md", "", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "")
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	revisionCommit(t, root, "stamped target")
	path := "changes/0001-target.md"
	readState := func() (target, index []byte, head, refs string) {
		target = staleRead(t, root, path)
		index = staleRead(t, root, ".git/index")
		head = revisionGit(t, root, "rev-parse", "HEAD")
		refs = revisionGit(t, root, "show-ref")
		return
	}
	target, index, head, refs := readState()
	got, err := project.StaleSpecs(root, "target")
	if err != nil || len(got) != 0 {
		t.Fatalf("matching base = %#v, %v; want empty slice", got, err)
	}
	put(t, root, "packages/alpha/SPEC.md", strings.Replace(string(staleRead(t, root, "packages/alpha/SPEC.md")), "One MUST work.", "One MUST now work differently.", 1))
	spec := staleRead(t, root, "packages/alpha/SPEC.md")
	got, err = project.StaleSpecs(root, "target")
	if err != nil || len(got) != 1 {
		t.Fatalf("stale base = %#v, %v; want one", got, err)
	}
	afterTarget, afterIndex, afterHead, afterRefs := readState()
	if !bytes.Equal(target, afterTarget) || !bytes.Equal(spec, staleRead(t, root, "packages/alpha/SPEC.md")) || !bytes.Equal(index, afterIndex) || head != afterHead || refs != afterRefs {
		t.Error("StaleSpecs changed a worktree file, index or ref")
	}
}

func Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent(t *testing.T) {
	root := staleFixture(t)
	queued(t, root, "0001-target.md", "custom: keep-me\nbases:\n  ALP-001: deadbeef\n  ALP-002: "+contracts.SpecHash(staleBaseItem(t, root, "ALP-002"))+"\n  ALP-999: extra\n", "Keep this motivation.\n", "", "### ALP-001\nOne MUST change.\n\n### ALP-002\nTwo MUST change.\n\n### ALP-003\nThree MUST change.\n", "### ALP-004\n")
	path := "changes/0001-target.md"
	_, original := staleBases(t, root, path)
	beforeCheck := project.Check(root)
	if err := project.StampMissing(root, "target"); err != nil {
		t.Fatal(err)
	}
	bases, after := staleBases(t, root, path)
	want := map[string]any{
		"ALP-001": "deadbeef", "ALP-002": contracts.SpecHash(staleBaseItem(t, root, "ALP-002")),
		"ALP-003": contracts.SpecHash(staleBaseItem(t, root, "ALP-003")),
		"ALP-004": contracts.SpecHash(staleBaseItem(t, root, "ALP-004")), "ALP-999": "extra",
	}
	if !reflect.DeepEqual(bases, want) || after.Frontmatter["custom"] != "keep-me" || after.Body != original.Body {
		t.Errorf("StampMissing lost bases/frontmatter/human sections: bases=%#v, body changed=%v", bases, after.Body != original.Body)
	}
	if len(beforeCheck) == 0 {
		t.Fatal("fixture lacked Check diagnostics before stamping")
	}
	staleFound := false
	for _, diagnostic := range project.Check(root) {
		if diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "stale base for ALP-001") {
			staleFound = true
		}
	}
	if !staleFound {
		t.Error("ordinary Check stopped reporting the preserved stale base")
	}
	if err := project.Stamp(root, "target"); err != nil {
		t.Fatal(err)
	}
	stamped, _ := staleBases(t, root, path)
	if stamped["ALP-001"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-001")) {
		t.Errorf("ordinary Stamp stopped replacing a stale base: %#v", stamped)
	}
}

func Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically(t *testing.T) {
	for _, tc := range []struct{ name, front, modified, removed string }{
		{"nonmapping", "bases: [ALP-001]\n", "### ALP-001\nOne MUST change.\n", ""},
		{"absent_modified_after_valid", "", "### ALP-001\nOne MUST change.\n\n### ALP-999\nMissing MUST change.\n", ""},
		{"absent_removed_after_valid", "", "### ALP-001\nOne MUST change.\n", "### ALP-998\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := staleFixture(t)
			path := "changes/0001-target.md"
			queued(t, root, "0001-target.md", tc.front, "Needed.\n", "", tc.modified, tc.removed)
			before := staleRead(t, root, path)
			err := project.StampMissing(root, "target")
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
				t.Errorf("StampMissing error = %v; want C08", err)
			}
			if !bytes.Equal(before, staleRead(t, root, path)) {
				t.Error("StampMissing rewrote target after rejection")
			}
		})
	}
}

func Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures(t *testing.T) {
	for _, alias := range []string{"target", "0001-target.md", "changes/0001-target.md"} {
		t.Run(alias, func(t *testing.T) {
			root := staleFixture(t)
			queued(t, root, "0001-target.md", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nOne MUST change.\n", "### ALP-002\n")
			got, err := project.StaleSpecs(root, alias)
			if err != nil || len(got) != 1 || got[0].ID != "ALP-001" {
				t.Errorf("StaleSpecs(%q) = %#v, %v", alias, got, err)
			}
			if err := project.StampMissing(root, alias); err != nil {
				t.Errorf("StampMissing(%q): %v", alias, err)
			}
			bases, _ := staleBases(t, root, "changes/0001-target.md")
			if bases["ALP-002"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-002")) || bases["ALP-001"] != "deadbeef" {
				t.Errorf("StampMissing(%q) bases = %#v", alias, bases)
			}
		})
	}
	root := staleFixture(t)
	for _, operation := range []struct {
		name string
		call func(string) error
	}{
		{"StaleSpecs", func(target string) error { _, err := project.StaleSpecs(root, target); return err }},
		{"StampMissing", func(target string) error { return project.StampMissing(root, target) }},
	} {
		t.Run(operation.name+"_unknown", func(t *testing.T) {
			const target = "changes/0099-absent.md"
			err := operation.call(target)
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
				t.Errorf("unknown selector = %v; want U02 with subject %q", err, target)
			}
		})
	}
	queued(t, root, "0001-target.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-target.md", "", "Second.\n", "", "", "")
	for _, operation := range []struct {
		name string
		call func(string) error
	}{
		{"StaleSpecs", func(target string) error { _, err := project.StaleSpecs(root, target); return err }},
		{"StampMissing", func(target string) error { return project.StampMissing(root, target) }},
	} {
		t.Run(operation.name+"_ambiguous", func(t *testing.T) {
			err := operation.call("target")
			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C13" {
				t.Errorf("ambiguous slug = %v; want C13", err)
			}
		})
	}
}

codex
The retained API delegates to three missing internal packages: `fixture`, `queue`, and `history`. The supplied workspace has no build manifest or dependency implementations, so I can check the implementation against the retained contracts and evals here, but the full Go suite will need the separate verifier.
exec
/bin/zsh -lc "sed -n '115,370p' packages/project/evals/queue_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/spec_history_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/evals/static_codes_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 536ms:
		if len(items) != 2 || findItem(t, items, "new:first").ID == "" {
			t.Errorf("Project(%q) did not stop at first change: %#v", target, items)
		}
		if err := project.Stamp(root, target); err != nil {
			t.Errorf("Stamp(%q): %v", target, err)
		}
	}
	if _, err := project.Project(root, "missing"); err == nil {
		t.Error("Project accepted missing target")
	}
	if err := project.Stamp(root, "missing"); err == nil {
		t.Error("Stamp accepted missing target")
	}
	put(t, root, "changes/archive/0000-archived.md", "archived document")
	if _, err := project.Project(root, "changes/archive/0000-archived.md"); err == nil {
		t.Error("Project accepted archived target")
	}
	queued(t, root, "0003-first.md", "", "Another first.\n", "", "", "")
	if _, err := project.Queue(root); err == nil {
		t.Error("Queue accepted ambiguous slug")
	}
	if _, err := project.Project(root, "first"); err == nil {
		t.Error("Project accepted ambiguous slug")
	}
	if err := project.Stamp(root, "first"); err == nil {
		t.Error("Stamp accepted ambiguous slug")
	}
}

func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:shared\nFirst change MUST retain its own text.\n", "### ALP-001 [review]\nAlpha MUST have revised full text.\n", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:shared\nSecond change MUST retain different text.\n", "", "### ALP-001\n")
	first, err := project.Project(root, "first")
	if err != nil {
		t.Fatalf("Project first: %v", err)
	}
	firstItems := specItems(t, moduleByName(t, first, "alpha"))
	if len(firstItems) != 2 || findItem(t, firstItems, "ALP-001").Tag != "review" || findItem(t, firstItems, "ALP-001").Body != "Alpha MUST have revised full text.\n\n" || findItem(t, firstItems, "new:shared").Body != "First change MUST retain its own text.\n\n" {
		t.Errorf("first projection = %#v", firstItems)
	}
	all, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project whole queue: %v", err)
	}
	section, ok := contracts.Section(moduleByName(t, all, "alpha").Document.Body, "Specs")
	if !ok {
		t.Fatal("Specs section absent from whole-queue projection")
	}
	if strings.Count(section, "### new:shared\n") != 2 ||
		!strings.Contains(section, "### new:shared\nFirst change MUST retain its own text.\n") ||
		!strings.Contains(section, "### new:shared\nSecond change MUST retain different text.\n") ||
		strings.Contains(section, "ALP-001") {
		t.Errorf("whole-queue projection lost local handles or retained removed ID: %q", section)
	}
}

func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
	queued(t, root, "0001-settings.md", "set:\n  depends_on: [beta]\n  interface: [next.go]\n  evals: [checks/**]\n  disposable: [generated/**]\n  preserve: [records/**]\n  append_only: [migrations/**]\n  boundaries_opt_out: reviewed boundary\n  hooks:\n    test:\n      command: [go, test, ./...]\n", "Settings are needed.\n", "", "", "")
	got, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project settings: %v", err)
	}
	alpha := moduleByName(t, got, "alpha")
	wantLists := map[string][]string{
		"depends_on": {"beta"}, "interface": {"next.go"}, "evals": {"checks/**"},
		"disposable": {"generated/**"}, "preserve": {"records/**"}, "append_only": {"migrations/**"},
	}
	for key, want := range wantLists {
		value, ok := alpha.Document.Frontmatter[key].([]any)
		if !ok || len(value) != 1 || value[0] != want[0] {
			t.Errorf("set %s = %#v; want %q", key, alpha.Document.Frontmatter[key], want)
		}
	}
	if !reflect.DeepEqual(alpha.DependsOn, []string{"beta"}) || !reflect.DeepEqual(alpha.Zones.Interface, []string{"next.go"}) || !reflect.DeepEqual(alpha.Zones.Evals, []string{"checks/**"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"generated/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"records/**"}) || !reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"migrations/**"}) {
		t.Errorf("projected module fields not updated: %#v", alpha)
	}
	wantHooks := map[string]any{"test": map[string]any{"command": []any{"go", "test", "./..."}}}
	if alpha.Document.Frontmatter["boundaries_opt_out"] != "reviewed boundary" || !reflect.DeepEqual(alpha.Document.Frontmatter["hooks"], wantHooks) {
		t.Errorf("projected scalar/map settings lost: %#v", alpha.Document.Frontmatter)
	}
	for _, tc := range []struct{ id, name, settings string }{
		{"PRJ-034", "identity_name", "name: renamed"},
		{"PRJ-034", "identity_prefix", "id_prefix: NEW"},
		{"PRJ-034", "unknown_field", "surprise: true"},
		{"PRJ-035", "invalid_graph", "depends_on: [missing]"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			badRoot := fixture(t)
			queued(t, badRoot, "0001-invalid.md", "set:\n  "+tc.settings+"\n", "Needed.\n", "", "", "")
			if _, err := project.Project(badRoot, ""); err == nil {
				t.Errorf("Project accepted %s", tc.name)
			}
		})
	}
}

func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-revise.md", "", "Earlier change.\n", "", "### ALP-001 [review]\nAlpha MUST reflect the earlier revision.\n", "")
	queued(t, root, "0002-target.md", "custom: author-value\n", "Keep this human motivation.\n", "", "### ALP-001\nAlpha MUST reflect the later revision.\n", "")
	wantHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Tag: "review", Body: "Alpha MUST reflect the earlier revision.\n"})
	if err := project.Stamp(root, "changes/0002-target.md"); err != nil {
		t.Fatalf("Stamp: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(root, "changes/0002-target.md"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("stamped document: %v", err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok || bases["ALP-001"] != wantHash {
		t.Errorf("bases = %#v; want preceding hash %s", doc.Frontmatter["bases"], wantHash)
	}
	if doc.Frontmatter["custom"] != "author-value" || !strings.Contains(doc.Body, "Keep this human motivation.") || !strings.Contains(doc.Body, "Alpha MUST reflect the later revision.") || !strings.Contains(doc.Body, "Human note.") {
		t.Errorf("Stamp lost author content: %#v", doc)
	}
	queued(t, root, "0003-remove.md", "", "Remove the item.\n", "", "", "### ALP-001\n")
	if err := project.Stamp(root, "0003-remove.md"); err != nil {
		t.Fatalf("Stamp Removed: %v", err)
	}
	removedBytes, err := os.ReadFile(filepath.Join(root, "changes/0003-remove.md"))
	if err != nil {
		t.Fatal(err)
	}
	removedDoc, err := contracts.ParseDocument(string(removedBytes))
	if err != nil {
		t.Fatal(err)
	}
	removedWant := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST reflect the later revision.\n"})
	removedBases, ok := removedDoc.Frontmatter["bases"].(map[string]any)
	if !ok || removedBases["ALP-001"] != removedWant {
		t.Errorf("Removed base = %#v; want %s", removedDoc.Frontmatter["bases"], removedWant)
	}
}

func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-added.md", "", "Added item.\n", "### new:thing\nThing MUST work.\n", "", "")
	if err := project.Stamp(root, "added"); err != nil {
		t.Fatalf("Stamp Added-only: %v", err)
	}
	bytes, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bytes), "### new:thing") || strings.Contains(string(bytes), "### ALP-002") {
		t.Errorf("Stamp assigned an ID to Added-only change: %s", bytes)
	}
	queued(t, root, "0002-invalid.md", "", "Invalid edit.\n", "", "### ALP-999\nMissing item MUST change.\n", "")
	if err := project.Stamp(root, "invalid"); err == nil {
		t.Error("Stamp accepted Modified ID absent from preceding projection")
	}
	removeRoot := fixture(t)
	queued(t, removeRoot, "0001-invalid-remove.md", "", "Invalid removal.\n", "", "", "### ALP-998\n")
	if err := project.Stamp(removeRoot, "invalid-remove"); err == nil {
		t.Error("Stamp accepted Removed ID absent from preceding projection")
	}
}

func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
	validRoot := fixture(t)
	put(t, validRoot, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	queued(t, validRoot, "0001-valid.md", "bases:\n  ALP-001: "+contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})+"\n", "A user needs the change.\n", "### new:local [review]\nA local addition MUST work.\n", "### ALP-001 [quarantine: flaky fixture]\nAlpha MUST work with revision.\n", "")
	queued(t, validRoot, "0002-also-valid.md", "", "Another user need.\n", "### new:local\nA second local addition MUST work.\n", "", "")
	for _, diagnostic := range project.Check(validRoot) {
		if !diagnostic.Warning {
			t.Errorf("valid spec/change received error diagnostic: %#v", diagnostic)
		}
	}
}

func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
	for _, tc := range []struct {
		id, name, spec, front, motivation, added, modified, removed string
		code, cause                                                 string
	}{
		{"PRJ-008", "missing_motivation", "", "", "\n", "", "", "", "C07", "missing Motivation"},
		{"PRJ-009", "nonnormative_module", "### ALP-001\nAlpha works.\n", "", "Needed.\n", "", "", "", "C04", "non-normative spec body"},
		{"PRJ-009", "nonnormative_change", "", "", "Needed.\n", "### new:item\nAn observation.\n", "", "", "C04", "non-normative spec body"},
		{"PRJ-010", "invalid_review_tag", "", "", "Needed.\n", "### new:item [review: reason]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-010", "invalid_quarantine_tag", "", "", "Needed.\n", "### new:item [quarantine]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-011", "duplicate_module_id", "### ALP-001\nAlpha MUST work.\n\n### ALP-001\nDuplicate MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "duplicate item ID"},
		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
		{"PRJ-013", "duplicate_edit_identity", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "### ALP-001\n", "C05", "duplicate edit identity ALP-001"},
		{"PRJ-014", "undeclared_handle", "", "", "Needed.\n", "### new:local\nLocal MUST work.\n", "### ALP-001\nAlpha MUST refer to new:foreign.\n", "", "C09", "reference to undeclared handle new:foreign"},
		{"PRJ-015", "missing_base", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "missing base for ALP-001"},
		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
		{"PRJ-016", "absent_modified", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "### ALP-999\nAbsent MUST change.\n", "", "C08", "ID ALP-999 absent from preceding projection"},
		{"PRJ-016", "absent_removed", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "", "### ALP-999\n", "C08", "ID ALP-999 absent from preceding projection"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			front := tc.front
			baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
			if tc.name == "duplicate_edit_identity" || tc.name == "undeclared_handle" {
				front = "bases:\n  ALP-001: " + baseHash + "\n"
			}
			if tc.name == "foreign_change_id" {
				put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
				front = "bases:\n  BET-001: " + contracts.SpecHash(contracts.Item{ID: "BET-001", Body: "Beta MUST work.\n"}) + "\n"
			}
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n"+func() string {
				if tc.spec != "" {
					return tc.spec
				}
				return "### ALP-001\nAlpha MUST work.\n"
			}())
			if tc.name == "undeclared_handle" {
				queued(t, root, "0000-other.md", "", "Earlier local handle.\n", "### new:foreign\nEarlier item MUST work.\n", "", "")
			}
			queued(t, root, "0001-check.md", front, tc.motivation, tc.added, tc.modified, tc.removed)
			diagnostics := project.Check(root)
			found := false
			foundCause := false
			foundApplyCause := false
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
					t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
				}
				if !diagnostic.Warning && strings.TrimSpace(diagnostic.Code) != "" && strings.TrimSpace(diagnostic.Message) != "" {
					found = true
				}
				matchesCause := strings.Contains(diagnostic.Message, tc.cause)
				if tc.id == "PRJ-016" {
					matchesCause = diagnostic.Message == tc.cause
				}
				if tc.code != "" && !diagnostic.Warning && diagnostic.Code == tc.code && matchesCause {
					foundCause = true
				}
				if tc.name == "absent_modified" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Modified ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
				if tc.name == "absent_removed" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Removed ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
			}
			if !found {
				t.Errorf("Check did not report error diagnostic with nonempty code and message: %#v", diagnostics)
			}
			if tc.code != "" && !foundCause {
				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
			}
			if tc.id == "PRJ-016" && !foundApplyCause {
				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
			}
			if tc.name == "missing_base" || tc.name == "stale_base" {
				queued(t, root, "0001-check.md", front, "Needed.\n", "", "", "### ALP-001\n")
				removedDiagnostics := project.Check(root)
				found = false

 succeeded in 536ms:
package evals_test

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/lurkshark/fireweed/packages/project"
)

func historyArchive(t *testing.T, root, path, module, motivation, added, modified, removed, decisions, removedDecisions, provenance string) {
	t.Helper()
	put(t, root, path, fmt.Sprintf("---\nmodule: %s\nphase: applied\n---\n\n## Motivation\n%s## Added\n%s## Modified\n%s## Removed\n%s## Decisions\n%s## Removed decisions\n%s## Provenance\n%s", module, motivation, added, modified, removed, decisions, removedDecisions, provenance))
}

func historyProject(t *testing.T) string {
	t.Helper()
	root := fixture(t)
	revisionGit(t, root, "init", "-q")
	historyArchive(t, root, "changes/archive/0009-added.md", "alpha",
		"Why the behavior exists.\nSecond motivation line.\n\n",
		"### ALP-004\nAnother behavior.\n### ALP-002\nRetained behavior.\n", "", "",
		"Choose the stable form.\nKeep this second decision.\n\n",
		"Prior choice was dropped.\n\n",
		"- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n")
	revisionCommit(t, root, "archive newest implementation date first")
	historyArchive(t, root, "changes/archive/0001-retag.md", "alpha", "Retag the behavior.\n", "",
		"### ALP-002 [review]\nRetained behavior.\n", "", "", "", "- impl 2024-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0003-equal.md", "alpha", "First equal date.\n", "",
		"### ALP-002\nEqual date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0004-equal.md", "alpha", "Second equal date.\n", "",
		"### ALP-002\nAnother equal date revision.\n", "", "", "", "- impl 2025-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0005-bad-date.md", "alpha", "Invalid date.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- impl not-a-date agent=fixture\n")
	historyArchive(t, root, "changes/archive/0006-no-impl.md", "alpha", "No implementation line.\n", "",
		"### ALP-002\nUndated revision.\n", "", "", "", "- evals 2030-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0010-other-id.md", "alpha", "Other specs.\n",
		"### ALP-005\nFifth behavior.\n### ALP-003\nThird behavior.\n",
		"### ALP-003 [review]\nThird behavior retagged.\n", "", "", "", "")
	historyArchive(t, root, "changes/archive/0011-removed-only.md", "alpha", "Only a removal.\n", "", "",
		"### ALP-002\n", "", "", "- impl 2031-01-01T00:00:00Z agent=fixture\n")
	historyArchive(t, root, "changes/archive/0012-beta.md", "beta", "Beta history.\n",
		"### BET-001\nBeta behavior.\n", "", "", "", "", "- impl 2032-01-01T00:00:00Z agent=fixture\n")
	queued(t, root, "0013-pending.md", "", "Not archived.\n", "### ALP-002\nPending behavior.\n", "", "")
	revisionCommit(t, root, "archive older dates and pending change later")
	return root
}

func historyPaths(entries []project.SpecHistoryEntry) []string {
	paths := make([]string, len(entries))
	for i, entry := range entries {
		paths[i] = entry.Path
	}
	return paths
}

func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	want := map[string]bool{
		"changes/archive/0009-added.md":    true,
		"changes/archive/0001-retag.md":    true,
		"changes/archive/0003-equal.md":    true,
		"changes/archive/0004-equal.md":    true,
		"changes/archive/0005-bad-date.md": true,
		"changes/archive/0006-no-impl.md":  true,
	}
	if len(entries) != len(want) {
		t.Fatalf("selected archives = %#v; want paths %#v", historyPaths(entries), want)
	}
	for _, entry := range entries {
		if !want[entry.Path] || entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, []string{"ALP-002"}) {
			t.Errorf("unexpected path/module/filtered IDs = %q, %q, %#v", entry.Path, entry.Module, entry.IDs)
		}
		delete(want, entry.Path)
	}
	if len(want) != 0 {
		t.Errorf("missing archives: %#v", want)
	}
}

func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("archive order/selection = %#v; want %#v", got, wantPaths)
	}
	wantDates := []time.Time{
		time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		{}, {},
	}
	for i, entry := range entries {
		if !entry.ImplementationDate.Equal(wantDates[i]) {
			t.Errorf("%s implementation date = %s; want %s", entry.Path, entry.ImplementationDate, wantDates[i])
		}
	}
}

func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil {
		t.Fatalf("SpecHistory: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("missing archived ALP-002 history")
	}
	want := project.SpecHistoryEntry{
		Path: "changes/archive/0009-added.md", Module: "alpha", IDs: []string{"ALP-002"},
		Motivation:         "Why the behavior exists.\nSecond motivation line.\n\n",
		Decisions:          "Choose the stable form.\nKeep this second decision.\n\n",
		RemovedDecisions:   "Prior choice was dropped.\n\n",
		Provenance:         "- evals 2026-01-01T00:00:00Z agent=fixture tests=2\n- impl 2023-01-01T00:00:00Z agent=fixture\n- impl not-a-date agent=fixture\n- impl 2026-06-01T12:30:00Z agent=fixture\n",
		ImplementationDate: time.Date(2026, 6, 1, 12, 30, 0, 0, time.UTC),
	}
	if !reflect.DeepEqual(entries[0], want) {
		t.Errorf("entry = %#v; want %#v", entries[0], want)
	}
}

func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
	root := historyProject(t)
	entries, err := project.SpecHistory(root, "alpha", "")
	if err != nil {
		t.Fatalf("SpecHistory module-wide: %v", err)
	}
	wantPaths := []string{
		"changes/archive/0009-added.md", "changes/archive/0003-equal.md",
		"changes/archive/0004-equal.md", "changes/archive/0001-retag.md",
		"changes/archive/0005-bad-date.md", "changes/archive/0006-no-impl.md",
		"changes/archive/0010-other-id.md",
	}
	if got := historyPaths(entries); !reflect.DeepEqual(got, wantPaths) {
		t.Fatalf("module-wide archive order/selection = %#v; want %#v", got, wantPaths)
	}
	for _, entry := range entries {
		wantIDs := []string{"ALP-002"}
		switch entry.Path {
		case "changes/archive/0009-added.md":
			wantIDs = []string{"ALP-002", "ALP-004"}
		case "changes/archive/0010-other-id.md":
			wantIDs = []string{"ALP-003", "ALP-005"}
		}
		if entry.Module != "alpha" || !reflect.DeepEqual(entry.IDs, wantIDs) {
			t.Errorf("%s module/unique IDs = %q, %#v; want alpha, %#v", entry.Path, entry.Module, entry.IDs, wantIDs)
		}
	}
	emptyRoot := fixture(t)
	empty, err := project.SpecHistory(emptyRoot, "beta", "")
	if err != nil || len(empty) != 0 {
		t.Errorf("module without archives = %#v, %v; want empty slice", empty, err)
	}
}

func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
	root := historyProject(t)
	archived, err := project.SpecHistory(root, "alpha", "ALP-002")
	if err != nil || len(archived) == 0 {
		t.Errorf("archived ID absent from current Specs = %#v, %v", archived, err)
	}
	currentOnly, err := project.SpecHistory(root, "alpha", "ALP-001")
	if err != nil || len(currentOnly) != 0 {
		t.Errorf("current ID with no archive edits = %#v, %v; want empty slice", currentOnly, err)
	}
	for _, tc := range []struct{ module, id, rejected string }{
		{"absent-module", "", "absent-module"},
		{"alpha", "ALP-999", "ALP-999"},
		{"alpha", "BET-001", "BET-001"},
	} {
		_, err := project.SpecHistory(root, tc.module, tc.id)
		if err == nil || !strings.HasPrefix(err.Error(), "U02") || !strings.Contains(err.Error(), tc.rejected) {
			t.Errorf("lookup module %q ID %q = %v; want U02 naming %q", tc.module, tc.id, err, tc.rejected)
		}
	}
}

func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
	for _, tc := range []struct{ name, config string }{
		{"missing", ""},
		{"malformed", "modules: [broken\n"},
	} {
		t.Run("config_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			if tc.name == "missing" {
				if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
					t.Fatal(err)
				}
			} else {
				put(t, root, "fireweed.yaml", tc.config)
			}
			_, err := project.SpecHistory(root, "alpha", "")
			if err == nil || !strings.HasPrefix(err.Error(), "C01") {
				t.Errorf("invalid config = %v; want C01", err)
			}
		})
	}
	for _, tc := range []struct{ name, document string }{
		{"document", "---\nmodule: [broken\n---\n## Added\n### ALP-002\nBehavior.\n"},
		{"added", "---\nmodule: alpha\n---\n## Added\n### ALP-002 [broken\nBehavior.\n"},
		{"modified", "---\nmodule: alpha\n---\n## Modified\n### ALP-002 [broken\nBehavior.\n"},
	} {
		t.Run("archive_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			path := "changes/archive/0001-invalid.md"
			put(t, root, path, tc.document)
			_, err := project.SpecHistory(root, "alpha", "ALP-002")
			if err == nil || !strings.HasPrefix(err.Error(), "C03") || !strings.Contains(err.Error(), path) {
				t.Errorf("malformed archive = %v; want C03 naming %s", err, path)
			}
		})
	}
}

 succeeded in 536ms:
package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_040_CheckConfigurationCauses(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed", func(t *testing.T, root string) {
			put(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			_, cause := project.Load(root)
			if cause == nil {
				t.Fatal("Load accepted invalid configuration")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C01 for configuration cause %q", cause)
		})
	}
}

func Test_PRJ_041_CheckGraphCauses(t *testing.T) {
	t.Run("current", func(t *testing.T) {
		root := fixture(t)
		put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		snapshot, err := project.Load(root)
		if err != nil {
			t.Fatal(err)
		}
		cause := project.ValidateGraph(snapshot.Modules)
		if cause == nil {
			t.Fatal("ValidateGraph accepted missing dependency")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for current graph cause %q", cause)
	})
	t.Run("projected", func(t *testing.T) {
		root := fixture(t)
		queued(t, root, "0001-invalid.md", "set:\n  depends_on: [missing]\n", "Needed.\n", "", "", "")
		_, cause := project.Project(root, "")
		if cause == nil {
			t.Fatal("Project accepted invalid projected graph")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for projected graph cause %q", cause)
	})
}

func Test_PRJ_042_CheckDocumentAndItemCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"missing_module", "load", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
		{"malformed_change", "queue", func(t *testing.T, root string) {
			put(t, root, "changes/0001-bad.md", "---\nmodule: [broken\n---\n")
		}},
		{"malformed_item", "queue", func(t *testing.T, root string) {
			queued(t, root, "0001-bad-item.md", "", "Needed.\n", "### new:item [broken\nItem MUST work.\n", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Queue(root)
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid document or item")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C03" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C03 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_043_CheckMetadataAndSetCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"metadata_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [alpha]\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n")
		}},
		{"zone_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: 7\n---\n# Alpha\n\n## Specs\n")
		}},
		{"identity_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  name: renamed\n", "Needed.\n", "", "", "")
		}},
		{"unknown_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  surprise: true\n", "Needed.\n", "", "", "")
		}},
		{"set_value_type", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  interface: 7\n", "Needed.\n", "", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Project(root, "")
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid metadata or set")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C12" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C12 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_044_CheckQueueIdentityCause(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-repeat.md", "", "Needed.\n", "", "", "")
	queued(t, root, "0002-repeat.md", "", "Needed.\n", "", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate slug")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C13" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C13 for queue identity cause %q", cause)
}

func Test_PRJ_045_CheckPendingModuleTargetCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0001-bad.md", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted unknown change module")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C07" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C07 for pending module cause %q", cause)
}

func Test_PRJ_013_CheckDuplicateWithinEditSection(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-duplicate.md", "", "Needed.\n", "### new:local\nLocal MUST work.\n\n### new:local\nDuplicate MUST fail.\n", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate Added identity")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C05" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C05 for duplicate Added identity cause %q", cause)
}

exec
/bin/zsh -lc "mkdir -p packages/project/internal/fixture packages/project/internal/queue packages/project/internal/history
cat > packages/project/internal/fixture/fixture.go <<'EOF'
package fixture

import (
 \"fmt\"
 \"os\"
 \"path/filepath\"
 \"regexp\"
 \"sort\"
 \"strconv\"
 \"strings\"

 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/gitstore\"
 \"github.com/lurkshark/fireweed/packages/project/model\"
)

func Coded(code, subject string, err error) error { return contracts.CodedError{Code:code, Subject:subject, Cause:fmt.Errorf(\"%s %s: %w\",code,subject,err)} }
func Code(err error) string { if c,_,ok:=contracts.ErrorCode(err); ok {return c}; return \"C03\" }
func Canonical(root string) (string,error) { p,e:=filepath.Abs(root); if e"'!=nil{return "",e}; return filepath.EvalSymlinks(p) }
func ReadConfig(root string)(map[string]any,error) { b,e:=os.ReadFile(filepath.Join(root,"fireweed.yaml")); if e!=nil{return nil,Coded("C01","fireweed.yaml",e)}; return DecodeConfig(b) }
func DecodeConfig(b []byte)(map[string]any,error) { doc,e:=contracts.ParseDocument("---'"\\n\"+string(b)+\"\\n---\\n\"); if e"'!=nil{return nil,Coded("C01","fireweed.yaml",e)}; c:=doc.Frontmatter; if _,e:=Patterns(c); e!=nil{return nil,Coded("C01","fireweed.yaml",e)}; return c,nil }
func Strings(v any)([]string,error) { a,ok:=v.([]any); if !ok{return nil,fmt.Errorf("expected list of strings, got %T",v)}; out:=make([]string,0,len(a)); for _,x:=range a {s,ok:=x.(string);if !ok{return nil,fmt.Errorf("expected string in list, got %T",x)};out=append(out,s)};return out,nil }
func Patterns(c map[string]any)([]string,error) { v,ok:=c["modules"];if !ok{return nil,fmt.Errorf("missing modules globs")}; return Strings(v) }
func Defaults(c map[string]any)(map[string]any,error) {v,ok:=c["module_defaults"];if !ok{return map[string]any{},nil};m,ok:=v.(map[string]any);if !ok{return nil,fmt.Errorf("module_defaults must be a mapping")};return m,nil}
var zoneKeys=[]string{"protected","evals","interface","append_only","preserve","disposable"}
func Metadata(doc contracts.Document,defaults map[string]any,dir string)(model.Module,error) { f:=doc.Frontmatter; name,ok:=f["name"].(string);if !ok{return model.Module{},Coded("C12",dir,fmt.Errorf("name must be a string"))};prefix,ok:=f["id_prefix"].(string);if !ok{return model.Module{},Coded("C12",dir,fmt.Errorf("id_prefix must be a string"))}; m:=model.Module{Name:name,IDPrefix:prefix,Dir:dir,Document:doc}; if v,ok:=f["depends_on"];ok {d,e:=Strings(v);if e!=nil{return m,Coded("C12",dir,fmt.Errorf("depends_on: %w",e))};m.DependsOn=d}; for _,key:=range zoneKeys {v,ok:=f[key];if !ok {v,ok=defaults[key]};if !ok{continue};list,e:=Strings(v);if e!=nil{return m,Coded("C12",dir,fmt.Errorf("%s: %w",key,e))};switch key {case "protected":m.Zones.Protected=list;case "evals":m.Zones.Evals=list;case "interface":m.Zones.Interface=list;case "append_only":m.Zones.AppendOnly=list;case "preserve":m.Zones.Preserve=list;case "disposable":m.Zones.Disposable=list} };return m,nil }
func ValidateGraph(modules []model.Module) error {names:=map[string]bool{};prefixes:=map[string]bool{};for _,m:=range modules {if strings.TrimSpace(m.Name)==""{return fmt.Errorf("empty module name")};if names[m.Name]{return fmt.Errorf("duplicate module name %s",m.Name)};names[m.Name]=true;if strings.TrimSpace(m.IDPrefix)==""{return fmt.Errorf("empty module prefix")};if prefixes[m.IDPrefix]{return fmt.Errorf("duplicate module prefix %s",m.IDPrefix)};prefixes[m.IDPrefix]=true};for _,m:=range modules{for _,d:=range m.DependsOn{if d==m.Name{return fmt.Errorf("self dependency %s",m.Name)};if !names[d]{return fmt.Errorf("module %s depends on missing dependency %s",m.Name,d)}}};edges:=map[string][]string{};for _,m:=range modules{edges[m.Name]=m.DependsOn};state:=map[string]int{};var visit func(string)error;visit=func(n string)error{if state[n]==1{return fmt.Errorf("dependency cycle involving %s",n)};if state[n]==2{return nil};state[n]=1;for _,d:=range edges[n]{if e:=visit(d);e!=nil{return e}};state[n]=2;return nil};for _,m:=range modules{if e:=visit(m.Name);e!=nil{return e}};return nil}
func Load(root string)(model.Snapshot,error){canon,e:=Canonical(root);if e!=nil{return model.Snapshot{},Coded("C01",root,e)};config,e:=ReadConfig(canon);if e!=nil{return model.Snapshot{},e};patterns,_:=Patterns(config);defaults,e:=Defaults(config);if e!=nil{return model.Snapshot{},Coded("C01","fireweed.yaml",e)};dirs:=map[string]bool{};for _,p:=range patterns{matches,e:=filepath.Glob(filepath.Join(canon,filepath.FromSlash(p)));if e!=nil{return model.Snapshot{},Coded("C01","fireweed.yaml",e)};for _,match:=range matches{info,e:=os.Stat(match);if e==nil&&info.IsDir(){rel,_:=filepath.Rel(canon,match);dirs[filepath.ToSlash(rel)]=true}}};return loadModules(canon,config,defaults,dirs,func(path string)([]byte,error){return os.ReadFile(filepath.Join(canon,filepath.FromSlash(path)))})}
func loadModules(root string,config,defaults map[string]any,dirs map[string]bool,read func(string)([]byte,error))(model.Snapshot,error){s:=model.Snapshot{Root:root,Config:config,Modules:[]model.Module{}};keys:=make([]string,0,len(dirs));for d:=range dirs{keys=append(keys,d)};sort.Strings(keys);for _,dir:=range keys{path:=dir+"/SPEC.md";b,e:=read(path);if e!=nil{return s,Coded("C03",path,e)};doc,e:=contracts.ParseDocument(string(b));if e!=nil{return s,Coded("C03",path,e)};if section,ok:=contracts.Section(doc.Body,"Specs");ok{items,e:=contracts.ParseItems(section);if e!=nil{return s,Coded("C03",path,e)};seen:=map[string]bool{};for _,it:=range items{if seen[it.ID]{return s,Coded("C05",path,fmt.Errorf("duplicate item ID %s",it.ID))};seen[it.ID]=true}};m,e:=Metadata(doc,defaults,dir);if e!=nil{return s,e};s.Modules=append(s.Modules,m)};return s,nil}
func LoadAt(root,revision string)(model.Snapshot,error){canon,e:=Canonical(root);if e!=nil{return model.Snapshot{},e};res,e:=gitstore.RunContextPlaceholder();_ = res;_ = e;return loadAt(canon,revision)}
func loadAt(root,revision string)(model.Snapshot,error){b,e:=gitstore.ReadAt(root,revision,"fireweed.yaml");if e!=nil{return model.Snapshot{},Coded("C01","fireweed.yaml",e)};config,e:=DecodeConfig(b);if e!=nil{return model.Snapshot{},e};patterns,_:=Patterns(config);defaults,e:=Defaults(config);if e!=nil{return model.Snapshot{},Coded("C01","fireweed.yaml",e)};paths,e:=TreePaths(root,revision);if e!=nil{return model.Snapshot{},e};dirs:=map[string]bool{};for _,path:=range paths{parts:=strings.Split(path,"/");for i:=1;i<len(parts);i++{dir:=strings.Join(parts[:i],"/");for _,p:=range patterns{if contracts.MatchGlob(p,dir){dirs[dir]=true}}}};return loadModules(root,config,defaults,dirs,func(path string)([]byte,error){return gitstore.ReadAt(root,revision,path)})}
func TreePaths(root,revision string)([]string,error){r,e:=gitstore.Run(context.Background(),gitstore.Request{Dir:root,Args:[]string{"git","ls-tree","-r","--name-only",revision}});if e!=nil{return nil,e};if r.ExitCode!=0{return nil,fmt.Errorf("git ls-tree %s: %s",revision,r.Stderr)};return strings.Fields(r.Stdout),nil}
var slugRE=regexp.MustCompile(`''^[a-z0-9-]+$`)
var filenameRE=regexp.MustCompile(`''^([0-9]{4,})-(.+)'"\\.md"'$`)
func ChangeName(name string)(int,string,bool){m:=filenameRE.FindStringSubmatch(name);if m==nil{return 0,"",false};n,_:=strconv.Atoi(m[1]);return n,m[2],true}
func NewChange(root,module,slug string)(string,error){if !slugRE.MatchString(slug){return "",fmt.Errorf("invalid slug %q",slug)};s,e:=Load(root);if e!=nil{return "",e};found:=false;for _,m:=range s.Modules{if m.Name==module{found=true}};if !found{return "",fmt.Errorf("unknown module %s",module)};max:=0;for _,dir:=range []string{"changes","changes/archive"}{entries,e:=os.ReadDir(filepath.Join(s.Root,dir));if e!=nil&& !os.IsNotExist(e){return "",e};for _,ent:=range entries{n,sl,ok:=ChangeName(ent.Name());if !ok{continue};if sl==slug{return "",fmt.Errorf("duplicate change slug %s",slug)};if n>max{max=n}}};rel:=fmt.Sprintf("changes/%04d-%s.md",max+1,slug);if e:=os.MkdirAll(filepath.Join(s.Root,"changes"),0755);e!=nil{return "",e};body:="---'"\\nmodule: \"+module+\"\\n---\\n\\n## Motivation\\n\\n## Interface\\n\\n## Added\\n\\n## Modified\\n\\n## Removed\\n\\n## Decisions\\n\";f,e:=os.OpenFile(filepath.Join(s.Root,filepath.FromSlash(rel)),os.O_WRONLY|os.O_CREATE|os.O_EXCL,0644);if e"'!=nil{return "",e};defer f.Close();_,e=f.WriteString(body);return rel,e}
EOF
python3 - <<'"'PY'
p='packages/project/internal/fixture/fixture.go'
s=open(p).read().replace(' \"fmt\"',' \"context\"\\n \"fmt\"').replace('res,e:=gitstore.RunContextPlaceholder();_ = res;_ = e;return loadAt(canon,revision)','return loadAt(canon,revision)')
open(p,'w').write(s)
PY" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 620ms:
exec
/bin/zsh -lc "cat > packages/project/internal/queue/queue.go <<'EOF'
package queue

import (
 \"fmt\"
 \"os\"
 \"path/filepath\"
 \"regexp\"
 \"sort\"
 \"strings\"

 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
 \"github.com/lurkshark/fireweed/packages/project/model\"
)

func parseChange(path string,b []byte)(model.Change,error){doc,e:=contracts.ParseDocument(string(b));if e"'!=nil{return model.Change{},fixture.Coded("C03",path,e)};n,slug,_:=fixture.ChangeName(filepath.Base(path));c:=model.Change{Number:n,Slug:slug,Path:path,Document:doc};m,ok:=doc.Frontmatter["module"].(string);if !ok{return c,fixture.Coded("C07",path,fmt.Errorf("invalid module target"))};c.Module=m;for _,part:=range []struct{name string;dest *[]contracts.Item}{{"Added",&c.Added},{"Modified",&c.Modified},{"Removed",&c.Removed}}{body,_:=contracts.Section(doc.Body,part.name);items,e:=contracts.ParseItems(body);if e!=nil{return c,fixture.Coded("C03",path,e)};*part.dest=items};seen:=map[string]bool{};for _,it:=range allEdits(c){if seen[it.ID]{return c,fixture.Coded("C05",path,fmt.Errorf("duplicate edit identity %s",it.ID))};seen[it.ID]=true};return c,nil}
func allEdits(c model.Change)[]contracts.Item{out:=append([]contracts.Item{},c.Added...);out=append(out,c.Modified...);out=append(out,c.Removed...);return out}
func queue(root string,s model.Snapshot,allowDuplicates bool)([]model.Change,error){ents,e:=os.ReadDir(filepath.Join(s.Root,"changes"));if os.IsNotExist(e){return []model.Change{},nil};if e!=nil{return nil,fixture.Coded("C03","changes",e)};known:=map[string]bool{};for _,m:=range s.Modules{known[m.Name]=true};out:=[]model.Change{};slugs:=map[string]bool{};for _,ent:=range ents{if ent.IsDir(){continue};_,slug,ok:=fixture.ChangeName(ent.Name());if !ok{continue};path:="changes/"+ent.Name();if slugs[slug]&&!allowDuplicates{return nil,fixture.Coded("C13",path,fmt.Errorf("duplicate pending slug %s",slug))};slugs[slug]=true;b,e:=os.ReadFile(filepath.Join(s.Root,filepath.FromSlash(path)));if e!=nil{return nil,fixture.Coded("C03",path,e)};c,e:=parseChange(path,b);if e!=nil{return nil,e};if !known[c.Module]{return nil,fixture.Coded("C07",path,fmt.Errorf("unknown change module %s",c.Module))};out=append(out,c)};sort.Slice(out,func(i,j int)bool{if out[i].Number!=out[j].Number{return out[i].Number<out[j].Number};return out[i].Slug<out[j].Slug});return out,nil}
func Queue(root string)([]model.Change,error){s,e:=fixture.Load(root);if e!=nil{return nil,e};return queue(root,s,false)}
func resolve(changes []model.Change,target string)(int,error){matches:=[]int{};for i,c:=range changes{if target==c.Slug||target==filepath.Base(c.Path)||target==c.Path{matches=append(matches,i)}};if len(matches)==0{return -1,fixture.Coded("U02",target,fmt.Errorf("unknown change %s",target))};if len(matches)>1{return -1,fixture.Coded("C13",target,fmt.Errorf("ambiguous change %s",target))};return matches[0],nil}
func Project(root,through string)(model.Snapshot,error){s,e:=fixture.Load(root);if e!=nil{return s,e};changes,e:=queue(root,s,false);if e!=nil{return s,e};limit:=len(changes);if through!=""{idx,e:=resolve(changes,through);if e!=nil{return s,e};limit=idx+1};for _,c:=range changes[:limit]{if e:=Apply(&s,c);e!=nil{return s,e}};if e:=fixture.ValidateGraph(s.Modules);e!=nil{return s,fixture.Coded("C02","graph",e)};return s,nil}
func preceding(root,target string)(model.Snapshot,model.Change,error){s,e:=fixture.Load(root);if e!=nil{return s,model.Change{},e};changes,e:=queue(root,s,true);if e!=nil{return s,model.Change{},e};idx,e:=resolve(changes,target);if e!=nil{return s,model.Change{},e};for _,c:=range changes[:idx]{if e:=Apply(&s,c);e!=nil{return s,model.Change{},e}};return s,changes[idx],nil}
func findModule(s *model.Snapshot,name string)*model.Module{for i:=range s.Modules{if s.Modules[i].Name==name{return &s.Modules[i]}};return nil}
func SpecItems(m model.Module)([]contracts.Item,error){body,_:=contracts.Section(m.Document.Body,"Specs");return contracts.ParseItems(body)}
func Before(s model.Snapshot,c model.Change)(map[string]contracts.Item,error){m:=findModule(&s,c.Module);if m==nil{return nil,fixture.Coded("C07",c.Path,fmt.Errorf("unknown change module %s",c.Module))};items,e:=SpecItems(*m);if e!=nil{return nil,fixture.Coded("C03",m.Dir+"/SPEC.md",e)};out:=map[string]contracts.Item{};for _,it:=range items{out[it.ID]=it};return out,nil}
func Apply(s *model.Snapshot,c model.Change)error{m:=findModule(s,c.Module);if m==nil{return fixture.Coded("C07",c.Path,fmt.Errorf("unknown change module %s",c.Module))};if e:=applySet(m,c);e!=nil{return e};items,e:=SpecItems(*m);if e!=nil{return fixture.Coded("C03",m.Dir+"/SPEC.md",e)};positions:=map[string]int{};for i,it:=range items{positions[it.ID]=i};for _,it:=range c.Modified{idx,ok:=positions[it.ID];if !ok{return fixture.Coded("C08",c.Path,fmt.Errorf("Modified ID %s absent from preceding projection",it.ID))};items[idx]=it};remove:=map[string]bool{};for _,it:=range c.Removed{if _,ok:=positions[it.ID];!ok{return fixture.Coded("C08",c.Path,fmt.Errorf("Removed ID %s absent from preceding projection",it.ID))};remove[it.ID]=true};kept:=[]contracts.Item{};for _,it:=range items{if !remove[it.ID]{kept=append(kept,it)}};kept=append(kept,c.Added...);var body strings.Builder;for _,it:=range kept{body.WriteString(contracts.RenderItem(it));body.WriteString("'"\\n\")};m.Document.Body=contracts.ReplaceSection(m.Document.Body,\"Specs\",body.String());return nil}
var listKeys=map[string]bool{\"depends_on\":true,\"interface\":true,\"evals\":true,\"disposable\":true,\"preserve\":true,\"append_only\":true}
func applySet(m *model.Module,c model.Change)error{raw,exists:=c.Document.Frontmatter[\"set\"];if "'!exists{return nil};set,ok:=raw.(map[string]any);if !ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set must be a mapping"))};for k,v:=range set{if listKeys[k]{if _,e:=fixture.Strings(v);e!=nil{return fixture.Coded("C12",c.Path,fmt.Errorf("set %s: %w",k,e))}}else{switch k{case "boundaries_opt_out":if _,ok:=v.(string);!ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set %s must be a string",k))};case "hooks":if _,ok:=v.(map[string]any);!ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set hooks must be a mapping"))};default:return fixture.Coded("C12",c.Path,fmt.Errorf("invalid set key %s",k))}};m.Document.Frontmatter[k]=v};defaults:=map[string]any{};updated,e:=fixture.Metadata(m.Document,defaults,m.Dir);if e!=nil{return e};*m=updated;return nil}
func bases(c model.Change)(map[string]any,error){v,ok:=c.Document.Frontmatter["bases"];if !ok{return map[string]any{},nil};m,ok:=v.(map[string]any);if !ok{return nil,fixture.Coded("C08",c.Path,fmt.Errorf("bases must be a mapping"))};return m,nil}
func Stamp(root,target string)error{return stamp(root,target,false)}
func StampMissing(root,target string)error{return stamp(root,target,true)}
func stamp(root,target string,missingOnly bool)error{s,c,e:=preceding(root,target);if e!=nil{return e};old,e:=Before(s,c);if e!=nil{return e};b,e:=bases(c);if e!=nil{return e};for _,part:=range []struct{name string;items []contracts.Item}{{"Modified",c.Modified},{"Removed",c.Removed}}{for _,it:=range part.items{prior,ok:=old[it.ID];if !ok{return fixture.Coded("C08",c.Path,fmt.Errorf("%s ID %s absent from preceding projection",part.name,it.ID))};if _,exists:=b[it.ID];!exists||!missingOnly{b[it.ID]=contracts.SpecHash(prior)}}};c.Document.Frontmatter["bases"]=b;rendered,e:=contracts.RenderDocument(c.Document);if e!=nil{return fixture.Coded("C03",c.Path,e)};path:=filepath.Join(s.Root,filepath.FromSlash(c.Path));info,e:=os.Stat(path);if e!=nil{return e};return os.WriteFile(path,[]byte(rendered),info.Mode().Perm())}
var handleRE=regexp.MustCompile(`'"\\bnew:[a-zA-Z0-9_-]+\\b"'`)
var normRE=regexp.MustCompile(`'"\\b(MUST|SHALL|SHOULD|MAY|MUST NOT|SHALL NOT|SHOULD NOT)\\b"'`)
func Check(root string)[]model.Diagnostic{out:=[]model.Diagnostic{};add:=func(code,msg,path string){out=append(out,model.Diagnostic{Code:code,Message:msg,Path:path})};s,e:=fixture.Load(root);if e!=nil{add(fixture.Code(e),e.Error(),"");return out};if e:=fixture.ValidateGraph(s.Modules);e!=nil{add("C02",e.Error(),"")};prefixes:=map[string]string{};for _,m:=range s.Modules{prefixes[m.IDPrefix]=m.Name};checkItem:=func(it contracts.Item,path,module string,needBody bool){if strings.Contains(it.ID,"-"){prefix:=strings.SplitN(it.ID,"-",2)[0];if owner,ok:=prefixes[prefix];ok&&owner!=module{add("C05",fmt.Sprintf("foreign spec ID %s belongs to %s",it.ID,owner),path)}};if needBody&&!normRE.MatchString(it.Body){add("C04","non-normative spec body for "+it.ID,path)};if it.Tag!=""&&it.Tag!="review"&&!strings.HasPrefix(it.Tag,"quarantine: "){add("C06","invalid review/quarantine tag for "+it.ID,path)};if strings.HasPrefix(it.Tag,"quarantine: ")&&strings.TrimSpace(strings.TrimPrefix(it.Tag,"quarantine: "))==""{add("C06","invalid review/quarantine tag for "+it.ID,path)}};for _,m:=range s.Modules{path:=m.Dir+"/SPEC.md";items,e:=SpecItems(m);if e!=nil{add("C03",e.Error(),path);continue};seen:=map[string]bool{};for _,it:=range items{if seen[it.ID]{add("C05","duplicate item ID "+it.ID,path)};seen[it.ID]=true;checkItem(it,path,m.Name,true)}};changes,e:=queue(root,s,false);if e!=nil{add(fixture.Code(e),e.Error(),"");return out};for _,c:=range changes{mot,_:=contracts.Section(c.Document.Body,"Motivation");if strings.TrimSpace(mot)==""{add("C07","missing Motivation in "+c.Path,c.Path)};if v,ok:=c.Document.Frontmatter["eval_fix"];ok&&nonempty(v){if v,ok:=c.Document.Frontmatter["fixes"];ok&&nonempty(v)||len(allEdits(c))>0{add("C07","eval_fix cannot be combined with fixes or spec edits",c.Path);continue}};prior,e:=Before(s,c);if e!=nil{add(fixture.Code(e),e.Error(),c.Path);continue};b,e:=bases(c);if e!=nil{add("C08",e.Error(),c.Path);continue};added:=map[string]bool{};for _,it:=range c.Added{if strings.HasPrefix(it.ID,"new:"){added[it.ID]=true}};for _,it:=range allEdits(c){checkItem(it,c.Path,c.Module,true);for _,ref:=range handleRE.FindAllString(it.Body,-1){if !added[ref]{add("C09","reference to undeclared handle "+ref,c.Path)}}};for _,part:=range []struct{name string;items []contracts.Item}{{"Modified",c.Modified},{"Removed",c.Removed}}{for _,it:=range part.items{p,ok:=prior[it.ID];if !ok{add("C08",fmt.Sprintf("ID %s absent from preceding projection",it.ID),c.Path);add("C08",fmt.Sprintf("%s ID %s absent from preceding projection",part.name,it.ID),c.Path);continue};v,ok:=b[it.ID];if !ok||v==""{add("C08","missing base for "+it.ID,c.Path)}else if v!=contracts.SpecHash(p){add("C08","stale base for "+it.ID,c.Path)}}};if e:=Apply(&s,c);e!=nil{add(fixture.Code(e),e.Error(),c.Path);continue};if e:=fixture.ValidateGraph(s.Modules);e!=nil{add("C02",e.Error(),c.Path)}};return out}
func nonempty(v any)bool{switch x:=v.(type){case []any:return len(x)>0;case string:return strings.TrimSpace(x)!="";default:return v!=nil}}
EOF' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 516ms:
codex
The loader, graph validation, queue parsing, projection, stamping, and diagnostic checks are in place. I’m finishing archive history and stale-base recovery, then I’ll review the code for preservation and error-code edge cases.
exec
/bin/zsh -lc "cat > packages/project/internal/queue/stale.go <<'EOF'
package queue

import (
 \"context\"
 \"fmt\"
 \"strings\"

 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/gitstore\"
 \"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
 \"github.com/lurkshark/fireweed/packages/project/model\"
)

func StaleSpecs(root,target string)([]model.StaleSpec,error){s,c,e:=preceding(root,target);if e"'!=nil{return nil,e};prior,e:=Before(s,c);if e!=nil{return nil,e};b,e:=bases(c);if e!=nil{return nil,e};out:=[]model.StaleSpec{};historical:=map[string]contracts.Item{};historyLoaded:=false;for _,it:=range append(append([]contracts.Item{},c.Modified...),c.Removed...){raw,ok:=b[it.ID];if !ok{continue};recorded,ok:=raw.(string);if !ok{recorded=fmt.Sprint(raw)};current,ok:=prior[it.ID];if !ok{continue};if recorded==contracts.SpecHash(current){continue};entry:=model.StaleSpec{ID:it.ID,RecordedHash:recorded,Current:contracts.RenderItem(current)};if !historyLoaded{historical=recoverBefore(s.Root,c);historyLoaded=true};if old,ok:=historical[it.ID];ok&&contracts.SpecHash(old)==recorded{entry.Before=contracts.RenderItem(old);entry.Recovered=true};out=append(out,entry)};return out,nil}
func recoverBefore(root string,c model.Change)map[string]contracts.Item{out:=map[string]contracts.Item{};r,e:=gitstore.Run(context.Background(),gitstore.Request{Dir:root,Args:[]string{"git","log","--format=%H","--",c.Path}});if e!=nil||r.ExitCode!=0{return out};commits:=strings.Fields(r.Stdout);var chosen string;var priorBases map[string]any;for _,rev:=range commits{b,e:=gitstore.ReadAt(root,rev,c.Path);if e!=nil{continue};doc,e:=contracts.ParseDocument(string(b));if e!=nil{continue};value,ok:=doc.Frontmatter["bases"].(map[string]any);if !ok{value=map[string]any{}};if chosen==""{priorBases=value;chosen=rev;continue};if !equalBases(priorBases,value){break};chosen=rev};if chosen==""{return out};hist,e:=fixture.LoadAt(root,chosen);if e!=nil{return out};files,e:=fixture.TreePaths(root,chosen);if e!=nil{return out};changes:=[]model.Change{};for _,path:=range files{if !strings.HasPrefix(path,"changes/")||strings.HasPrefix(path,"changes/archive/"){continue};if _,_,ok:=fixture.ChangeName(strings.TrimPrefix(path,"changes/"));!ok{continue};b,e:=gitstore.ReadAt(root,chosen,path);if e!=nil{return out};item,e:=parseChange(path,b);if e!=nil{return out};changes=append(changes,item)};sortChanges(changes);index:=-1;for i,item:=range changes{if item.Path==c.Path{index=i;break}};if index<0{return out};for _,item:=range changes[:index]{if e:=Apply(&hist,item);e!=nil{return out}};m:=findModule(&hist,c.Module);if m==nil{return out};items,e:=SpecItems(*m);if e!=nil{return out};for _,it:=range items{out[it.ID]=it};return out}
func equalBases(a,b map[string]any)bool{if len(a)!=len(b){return false};for k,v:=range a{if b[k]!=v{return false}};return true}
func sortChanges(cs []model.Change){for i:=1;i<len(cs);i++{for j:=i;j>0&&(cs[j].Number<cs[j-1].Number||(cs[j].Number==cs[j-1].Number&&cs[j].Slug<cs[j-1].Slug));j--{cs[j],cs[j-1]=cs[j-1],cs[j]}}}
EOF
cat > packages/project/internal/history/history.go <<'"'EOF'
package history

import (
 \"fmt\"
 \"os\"
 \"path/filepath\"
 \"sort\"
 \"strings\"
 \"time\"

 \"github.com/lurkshark/fireweed/packages/contracts\"
 \"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
 \"github.com/lurkshark/fireweed/packages/project/model\"
)

func SpecHistory(root,module,id string)([]model.SpecHistoryEntry,error){s,e:=fixture.Load(root);if e"'!=nil{return nil,e};var target *model.Module;for i:=range s.Modules{if s.Modules[i].Name==module{target=&s.Modules[i];break}};if target==nil{return nil,fixture.Coded("U02",module,fmt.Errorf("U02 unknown module %s",module))};paths,e:=os.ReadDir(filepath.Join(s.Root,"changes/archive"));if os.IsNotExist(e){paths=nil}else if e!=nil{return nil,fixture.Coded("C03","changes/archive",e)};out:=[]model.SpecHistoryEntry{};found:=false;for _,ent:=range paths{if ent.IsDir()||!strings.HasSuffix(ent.Name(),".md"){continue};path:="changes/archive/"+ent.Name();b,e:=os.ReadFile(filepath.Join(s.Root,filepath.FromSlash(path)));if e!=nil{return nil,fixture.Coded("C03",path,e)};doc,e:=contracts.ParseDocument(string(b));if e!=nil{return nil,fixture.Coded("C03",path,e)};mod,ok:=doc.Frontmatter["module"].(string);if !ok{return nil,fixture.Coded("C03",path,fmt.Errorf("invalid archive module"))};addedText,_:=contracts.Section(doc.Body,"Added");modifiedText,_:=contracts.Section(doc.Body,"Modified");added,e:=contracts.ParseItems(addedText);if e!=nil{return nil,fixture.Coded("C03",path,e)};modified,e:=contracts.ParseItems(modifiedText);if e!=nil{return nil,fixture.Coded("C03",path,e)};if mod!=module{continue};ids:=map[string]bool{};for _,it:=range append(added,modified...){ids[it.ID]=true};if id!=""{if !ids[id]{continue};found=true;ids=map[string]bool{id:true}};if len(ids)==0{continue};list:=make([]string,0,len(ids));for k:=range ids{list=append(list,k)};sort.Strings(list);mot,_:=contracts.Section(doc.Body,"Motivation");dec,_:=contracts.Section(doc.Body,"Decisions");removed,_:=contracts.Section(doc.Body,"Removed decisions");prov,_:=contracts.Section(doc.Body,"Provenance");out=append(out,model.SpecHistoryEntry{Path:path,Module:module,IDs:list,Motivation:mot,Decisions:dec,RemovedDecisions:removed,Provenance:prov,ImplementationDate:implDate(prov)})};if id!=""&&!found{body,_:=contracts.Section(target.Document.Body,"Specs");items,e:=contracts.ParseItems(body);if e!=nil{return nil,fixture.Coded("C03",target.Dir+"/SPEC.md",e)};current:=false;for _,it:=range items{if it.ID==id{current=true}};if !current{return nil,fixture.Coded("U02",id,fmt.Errorf("U02 unknown spec ID %s",id))}};sort.Slice(out,func(i,j int)bool{a,b:=out[i],out[j];if !a.ImplementationDate.Equal(b.ImplementationDate){if a.ImplementationDate.IsZero(){return false};if b.ImplementationDate.IsZero(){return true};return a.ImplementationDate.After(b.ImplementationDate)};return a.Path<b.Path});return out,nil}
func implDate(prov string)time.Time{var latest time.Time;for _,line:=range strings.Split(prov,"'"\\n\"){fields:=strings.Fields(line);if len(fields)<3||fields[0]"'!="-"||fields[1]!="impl"{continue};t,e:=time.Parse(time.RFC3339Nano,fields[2]);if e==nil&&t.After(latest){latest=t}};return latest}
EOF
python3 - <<'"'PY'
p='packages/project/internal/queue/queue.go';s=open(p).read();s=s.replace('for _,it:=range allEdits(c){checkItem(it,c.Path,c.Module,true);for _,ref:=range handleRE.FindAllString(it.Body,-1){','for _,it:=range append(append([]contracts.Item{},c.Added...),c.Modified...){checkItem(it,c.Path,c.Module,true)};for _,it:=range c.Removed{checkItem(it,c.Path,c.Module,false)};for _,it:=range allEdits(c){for _,ref:=range handleRE.FindAllString(it.Body,-1){');open(p,'w').write(s)
PY
gofmt -w packages/project/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 734ms:
exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/project/internal/queue/queue.go';s=open(p).read();start=s.index('func applySet(');end=s.index('\\nfunc bases(',start);s=s[:start]+'''func applySet(m *model.Module, c model.Change) error {
 raw,exists:=c.Document.Frontmatter[\"set\"]
 if "'!exists{return nil}
 set,ok:=raw.(map[string]any)
 if !ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set must be a mapping"))}
 for k,v:=range set{
  if listKeys[k]{if _,e:=fixture.Strings(v);e!=nil{return fixture.Coded("C12",c.Path,fmt.Errorf("set %s: %w",k,e))}}
  else {switch k{case "boundaries_opt_out":if _,ok:=v.(string);!ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set %s must be a string",k))};case "hooks":if _,ok:=v.(map[string]any);!ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set hooks must be a mapping"))};default:return fixture.Coded("C12",c.Path,fmt.Errorf("invalid set key %s",k))}}
 }
 for k,v:=range set{
  m.Document.Frontmatter[k]=v
  if listKeys[k]{list,_:=fixture.Strings(v);switch k{case "depends_on":m.DependsOn=list;case "interface":m.Zones.Interface=list;case "evals":m.Zones.Evals=list;case "disposable":m.Zones.Disposable=list;case "preserve":m.Zones.Preserve=list;case "append_only":m.Zones.AppendOnly=list}}
 }
 return nil
}'"''' + s[end:];open(p,'w').write(s)
p='packages/project/internal/fixture/fixture.go';s=open(p).read();s=s.replace('return s, Coded(\"C03\", path, e)\\n\\t\\t}', 'return s, Coded(Code(e), path, e)\\n\\t\\t}', 1) if False else s;open(p,'w').write(s)
PY
gofmt -w packages/project/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 exited 2 in 652ms:
packages/project/internal/queue/queue.go:253:3: expected statement, found 'else'
packages/project/internal/queue/queue.go:255:2: expected declaration, found 'for'

exec
/bin/zsh -lc "sed -n '240,265p' packages/project/internal/queue/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 454ms:
	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
	return nil
}

var listKeys = map[string]bool{"depends_on": true, "interface": true, "evals": true, "disposable": true, "preserve": true, "append_only": true}

func applySet(m *model.Module, c model.Change) error {
 raw,exists:=c.Document.Frontmatter["set"]
 if !exists{return nil}
 set,ok:=raw.(map[string]any)
 if !ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set must be a mapping"))}
 for k,v:=range set{
  if listKeys[k]{if _,e:=fixture.Strings(v);e!=nil{return fixture.Coded("C12",c.Path,fmt.Errorf("set %s: %w",k,e))}}
  else {switch k{case "boundaries_opt_out":if _,ok:=v.(string);!ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set %s must be a string",k))};case "hooks":if _,ok:=v.(map[string]any);!ok{return fixture.Coded("C12",c.Path,fmt.Errorf("set hooks must be a mapping"))};default:return fixture.Coded("C12",c.Path,fmt.Errorf("invalid set key %s",k))}}
 }
 for k,v:=range set{
  m.Document.Frontmatter[k]=v
  if listKeys[k]{list,_:=fixture.Strings(v);switch k{case "depends_on":m.DependsOn=list;case "interface":m.Zones.Interface=list;case "evals":m.Zones.Evals=list;case "disposable":m.Zones.Disposable=list;case "preserve":m.Zones.Preserve=list;case "append_only":m.Zones.AppendOnly=list}}
 }
 return nil
}
func bases(c model.Change) (map[string]any, error) {
	v, ok := c.Document.Frontmatter["bases"]
	if !ok {
		return map[string]any{}, nil
	}

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/project/internal/queue/queue.go';s=open(p).read().replace('}}\\n  else {switch k','}} else {switch k');open(p,'w').write(s)
PY
gofmt -w packages/project/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 617ms:
exec
/bin/zsh -lc "sed -n '355,450p' packages/project/evals/queue_test.go
rg -n 'C05|C08|Stamp' packages/project/evals/coded_errors_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 464ms:
					foundApplyCause = true
				}
			}
			if !found {
				t.Errorf("Check did not report error diagnostic with nonempty code and message: %#v", diagnostics)
			}
			if tc.code != "" && !foundCause {
				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
			}
			if tc.id == "PRJ-016" && !foundApplyCause {
				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
			}
			if tc.name == "missing_base" || tc.name == "stale_base" {
				queued(t, root, "0001-check.md", front, "Needed.\n", "", "", "### ALP-001\n")
				removedDiagnostics := project.Check(root)
				found = false
				for _, diagnostic := range removedDiagnostics {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message for Removed entry: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, tc.cause) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not report C08 for Removed entry with %s (%s): %#v", tc.name, tc.cause, removedDiagnostics)
				}
			}
			if tc.name == "nonnormative_change" || tc.name == "invalid_review_tag" {
				front, added, modified, want := "", "", "", ""
				code := "C04"
				if tc.name == "nonnormative_change" {
					front = "bases:\n  ALP-001: " + baseHash + "\n"
					modified = "### ALP-001\nAn observation.\n"
					want = "non-normative spec body"
				} else {
					code = "C06"
					added = "### new:item [other: reason]\nItem MUST work.\n"
					want = "invalid review/quarantine tag"
				}
				queued(t, root, "0001-check.md", front, "Needed.\n", added, modified, "")
				found = false
				for _, diagnostic := range project.Check(root) {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == code && strings.Contains(diagnostic.Message, want) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not reject %s: wanted error diagnostic %s with message containing %q", tc.name, code, want)
				}
			}
		})
	}
}
10:func Test_PRJ_057_StampInvalidBasesHasCode(t *testing.T) {
13:	err := project.Stamp(root, "invalid-bases")
15:		t.Fatal("Stamp accepted bases that are not a mapping")
17:	if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
18:		t.Errorf("Stamp error code = (%q, %v); want (C08, true): %v", code, ok, err)
29:		{"Stamp", func() error { return project.Stamp(root, target) }},
50:		{"duplicate_spec_id", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n\n### ALP-001\nAlpha MUST also work.\n", "C05"},

exec
/bin/zsh -lc "rg -n -C 3 'ParseItems|ParseDocument|loadAt|TreePaths' packages/project/internal/fixture/fixture.go
rg -n 'func Check|checkItem|for _, it := range allEdits' packages/project/internal/queue/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 454ms:
39-	return DecodeConfig(b)
40-}
41-func DecodeConfig(b []byte) (map[string]any, error) {
42:	doc, e := contracts.ParseDocument("---\n" + string(b) + "\n---\n")
43-	if e != nil {
44-		return nil, Coded("C01", "fireweed.yaml", e)
45-	}
--
233-		if e != nil {
234-			return s, Coded("C03", path, e)
235-		}
236:		doc, e := contracts.ParseDocument(string(b))
237-		if e != nil {
238-			return s, Coded("C03", path, e)
239-		}
240-		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
241:			items, e := contracts.ParseItems(section)
242-			if e != nil {
243-				return s, Coded("C03", path, e)
244-			}
--
263-	if e != nil {
264-		return model.Snapshot{}, e
265-	}
266:	return loadAt(canon, revision)
267-}
268:func loadAt(root, revision string) (model.Snapshot, error) {
269-	b, e := gitstore.ReadAt(root, revision, "fireweed.yaml")
270-	if e != nil {
271-		return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
--
279-	if e != nil {
280-		return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
281-	}
282:	paths, e := TreePaths(root, revision)
283-	if e != nil {
284-		return model.Snapshot{}, e
285-	}
--
297-	}
298-	return loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
299-}
300:func TreePaths(root, revision string) ([]string, error) {
301-	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "ls-tree", "-r", "--name-only", revision}})
302-	if e != nil {
303-		return nil, e
40:	for _, it := range allEdits(c) {
353:func Check(root string) []model.Diagnostic {
370:	checkItem := func(it contracts.Item, path, module string, needBody bool) {
400:			checkItem(it, path, m.Name, true)
436:			checkItem(it, c.Path, c.Module, true)
439:			checkItem(it, c.Path, c.Module, false)
441:		for _, it := range allEdits(c) {

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/project/internal/fixture/fixture.go';s=open(p).read();s=s.replace('return s, Coded(\"C03\", path, e)\\n\\t\\t\\t}\\n\\t\\t\\tseen', 'return s, Coded(Code(e), path, e)\\n\\t\\t\\t}\\n\\t\\t\\tseen');open(p,'w').write(s)
p='packages/project/internal/queue/queue.go';s=open(p).read();s=s.replace('body.WriteString(contracts.RenderItem(it))\\n\\t\\tbody.WriteString(\"\\\\n\")','rendered := contracts.RenderItem(it)\\n\\t\\tbody.WriteString(rendered)\\n\\t\\tif "'!strings.HasSuffix(rendered, "'"\\\\n\\\\n\") { body.WriteString(\"\\\\n\") }');open(p,'w').write(s)
PY
gofmt -w packages/project/internal
rg -n -C 2 'body.WriteString|return s, Coded\\(Code' packages/project/internal/queue/queue.go packages/project/internal/fixture/fixture.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 682ms:
packages/project/internal/fixture/fixture.go-241-			items, e := contracts.ParseItems(section)
packages/project/internal/fixture/fixture.go-242-			if e != nil {
packages/project/internal/fixture/fixture.go:243:				return s, Coded(Code(e), path, e)
packages/project/internal/fixture/fixture.go-244-			}
packages/project/internal/fixture/fixture.go-245-			seen := map[string]bool{}
--
packages/project/internal/queue/queue.go-236-	for _, it := range kept {
packages/project/internal/queue/queue.go-237-		rendered := contracts.RenderItem(it)
packages/project/internal/queue/queue.go:238:		body.WriteString(rendered)
packages/project/internal/queue/queue.go-239-		if !strings.HasSuffix(rendered, "\n\n") {
packages/project/internal/queue/queue.go:240:			body.WriteString("\n")
packages/project/internal/queue/queue.go-241-		}
packages/project/internal/queue/queue.go-242-	}

exec
/bin/zsh -lc "sed -n '180,230p' packages/project/internal/fixture/fixture.go
sed -n '1,80p' packages/project/internal/queue/stale.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 454ms:
				return e
			}
		}
		state[n] = 2
		return nil
	}
	for _, m := range modules {
		if e := visit(m.Name); e != nil {
			return e
		}
	}
	return nil
}
func Load(root string) (model.Snapshot, error) {
	canon, e := Canonical(root)
	if e != nil {
		return model.Snapshot{}, Coded("C01", root, e)
	}
	config, e := ReadConfig(canon)
	if e != nil {
		return model.Snapshot{}, e
	}
	patterns, _ := Patterns(config)
	defaults, e := Defaults(config)
	if e != nil {
		return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
	}
	dirs := map[string]bool{}
	for _, p := range patterns {
		matches, e := filepath.Glob(filepath.Join(canon, filepath.FromSlash(p)))
		if e != nil {
			return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
		}
		for _, match := range matches {
			info, e := os.Stat(match)
			if e == nil && info.IsDir() {
				rel, _ := filepath.Rel(canon, match)
				dirs[filepath.ToSlash(rel)] = true
			}
		}
	}
	return loadModules(canon, config, defaults, dirs, func(path string) ([]byte, error) { return os.ReadFile(filepath.Join(canon, filepath.FromSlash(path))) })
}
func loadModules(root string, config, defaults map[string]any, dirs map[string]bool, read func(string) ([]byte, error)) (model.Snapshot, error) {
	s := model.Snapshot{Root: root, Config: config, Modules: []model.Module{}}
	keys := make([]string, 0, len(dirs))
	for d := range dirs {
		keys = append(keys, d)
	}
	sort.Strings(keys)
	for _, dir := range keys {
package queue

import (
	"context"
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func StaleSpecs(root, target string) ([]model.StaleSpec, error) {
	s, c, e := preceding(root, target)
	if e != nil {
		return nil, e
	}
	prior, e := Before(s, c)
	if e != nil {
		return nil, e
	}
	b, e := bases(c)
	if e != nil {
		return nil, e
	}
	out := []model.StaleSpec{}
	historical := map[string]contracts.Item{}
	historyLoaded := false
	for _, it := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
		raw, ok := b[it.ID]
		if !ok {
			continue
		}
		recorded, ok := raw.(string)
		if !ok {
			recorded = fmt.Sprint(raw)
		}
		current, ok := prior[it.ID]
		if !ok {
			continue
		}
		if recorded == contracts.SpecHash(current) {
			continue
		}
		entry := model.StaleSpec{ID: it.ID, RecordedHash: recorded, Current: contracts.RenderItem(current)}
		if !historyLoaded {
			historical = recoverBefore(s.Root, c)
			historyLoaded = true
		}
		if old, ok := historical[it.ID]; ok && contracts.SpecHash(old) == recorded {
			entry.Before = contracts.RenderItem(old)
			entry.Recovered = true
		}
		out = append(out, entry)
	}
	return out, nil
}
func recoverBefore(root string, c model.Change) map[string]contracts.Item {
	out := map[string]contracts.Item{}
	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "log", "--format=%H", "--", c.Path}})
	if e != nil || r.ExitCode != 0 {
		return out
	}
	commits := strings.Fields(r.Stdout)
	var chosen string
	var priorBases map[string]any
	for _, rev := range commits {
		b, e := gitstore.ReadAt(root, rev, c.Path)
		if e != nil {
			continue
		}
		doc, e := contracts.ParseDocument(string(b))
		if e != nil {
			continue
		}
		value, ok := doc.Frontmatter["bases"].(map[string]any)
		if !ok {
			value = map[string]any{}
		}

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/project/internal/fixture/fixture.go';s=open(p).read().replace('"'^([0-9]{4,})-(.+)'"\\\\.md"'$'"','"'^([0-9]{4,})-([a-z0-9-]+)'"\\\\.md"'$'"');open(p,'w').write(s)
p='packages/project/internal/queue/queue.go';s=open(p).read();s=s.replace('if v, ok := c.Document.Frontmatter[\"eval_fix\"]; ok && nonempty(v) {\\n\\t\\t\\tif v, ok := c.Document.Frontmatter[\"fixes\"]; ok && nonempty(v) || len(allEdits(c)) > 0 {','if _, ok := c.Document.Frontmatter[\"eval_fix\"]; ok {\\n\\t\\t\\tif _, hasFixes := c.Document.Frontmatter[\"fixes\"]; hasFixes || len(allEdits(c)) > 0 {');start=s.find('\\t\\tfor _, it := range allEdits(c) {\\n\\t\\t\\tfor _, ref := range handleRE.FindAllString(it.Body, -1)');
if start"'!=-1:
 end=s.find('"'\\n\\t\\tfor _, part := range []struct',start)
 s=s[:start]+'''\\t\\tfor _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
            if "'!added[ref] { add("C09", "reference to undeclared handle "+ref, c.Path) }
        }'"'''+s[end:]
open(p,'w').write(s)
PY
gofmt -w packages/project/internal
rg -n 'nonempty|eval_fix|handleRE.Find' packages/project/internal/queue/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 658ms:
416:		if _, ok := c.Document.Frontmatter["eval_fix"]; ok {
418:				add("C07", "eval_fix cannot be combined with fixes or spec edits", c.Path)
444:		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
478:func nonempty(v any) bool {

codex
The read paths now use Git data for revision snapshots and historical stale-base recovery; `StaleSpecs` itself does not write repository state. I’m doing a final check of the implementation and recording the verification limits in the requested notes.
exec
/bin/zsh -lc 'PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache go test ./packages/project/...' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "git status --short; rg --files packages/project/internal; rg -n 'RunContextPlaceholder|TODO|FIXME' packages/project/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 exited 1 in 538ms:
?? packages/project/internal/
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go
packages/project/internal/history/history.go
packages/project/internal/fixture/fixture.go

 exited 1 in 559ms:
# ./packages/project/...
pattern ./packages/project/...: directory prefix packages/project does not contain main module or its selected dependencies
FAIL	./packages/project/... [setup failed]
FAIL

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/project/internal/queue/queue.go';s=open(p).read();s=s.replace('if e := fixture.ValidateGraph(s.Modules); e "'!= nil {'"\\n\\t\\t\\tadd(\"C02\", e.Error(), c.Path)\\n\\t\\t}', 'if e := fixture.ValidateGraph(s.Modules); e "'!= nil {'"\\n\\t\\t\\tadd(\"C02\", fixture.Coded(\"C02\", \"graph\", e).Error(), c.Path)\\n\\t\\t}');open(p,'w').write(s)
PY
gofmt -w packages/project/internal
rg -n -A 3 'ValidateGraph\\(s.Modules\\)' packages/project/internal/queue/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 711ms:
146:	if e := fixture.ValidateGraph(s.Modules); e != nil {
147-		return s, fixture.Coded("C02", "graph", e)
148-	}
149-	return s, nil
--
366:	if e := fixture.ValidateGraph(s.Modules); e != nil {
367-		add("C02", e.Error(), "")
368-	}
369-	prefixes := map[string]string{}
--
472:		if e := fixture.ValidateGraph(s.Modules); e != nil {
473-			add("C02", fixture.Coded("C02", "graph", e).Error(), c.Path)
474-		}
475-	}

exec
/bin/zsh -lc "python3 - <<'PY'
p='packages/project/internal/queue/stale.go';s=open(p).read().replace('\"fmt\"\\n', '\"fmt\"\\n \"reflect\"\\n').replace('if b[k] "'!= v {'"','if "'!reflect.DeepEqual(b[k], v) {'"');open(p,'w').write(s)
p='packages/project/internal/queue/queue.go';s=open(p).read();start=s.find('\\nfunc nonempty(v any) bool {');
if start"'!=-1:s=s[:start]+'"'\\n';open(p,'w').write(s)
PY
gofmt -w packages/project/internal
cat > DRILL_NOTES.md <<'EOF'
## Contradicted Decisions

None identified.

## Gaps

- PRJ-007 (project): The supplied input has no "'`go.mod` or dependency implementations, so the native Go eval suite could not run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.
EOF
git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 665ms:
?? DRILL_NOTES.md
?? packages/project/internal/

exec
/bin/zsh -lc "gofmt -l packages/project/internal/fixture/fixture.go packages/project/internal/queue/queue.go packages/project/internal/queue/stale.go packages/project/internal/history/history.go
rg -n 'func (Load|LoadAt|ValidateGraph|NewChange|Queue|Project|Stamp|StampMissing|StaleSpecs|Check|SpecHistory)' packages/project/internal" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 491ms:
packages/project/internal/queue/queue.go:102:func Queue(root string) ([]model.Change, error) {
packages/project/internal/queue/queue.go:124:func Project(root, through string) (model.Snapshot, error) {
packages/project/internal/queue/queue.go:311:func Stamp(root, target string) error        { return stamp(root, target, false) }
packages/project/internal/queue/queue.go:312:func StampMissing(root, target string) error { return stamp(root, target, true) }
packages/project/internal/queue/queue.go:356:func Check(root string) []model.Diagnostic {
packages/project/internal/queue/stale.go:15:func StaleSpecs(root, target string) ([]model.StaleSpec, error) {
packages/project/internal/history/history.go:16:func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
packages/project/internal/fixture/fixture.go:135:func ValidateGraph(modules []model.Module) error {
packages/project/internal/fixture/fixture.go:193:func Load(root string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:261:func LoadAt(root, revision string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:322:func NewChange(root, module, slug string) (string, error) {

codex
Rebuilt the disposable project internals for loading, queueing, projection, stamping, checks, revision snapshots, spec history, and stale-base recovery. Added [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/DRILL_NOTES.md).

`gofmt` passed. The Go evals could not run in this supplied workspace because it has no `go.mod` or dependency implementations; the notes record that verification gap. Retained inputs were not edited, and nothing was committed.
tokens used
72,808


### Attempt 2 (exit 0)
Rebuilt the disposable project implementation to address the reported `Load` root mismatch, C05 duplicate and foreign-ID diagnostics, and `NewChange` error codes. Updated [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/DRILL_NOTES.md). Retained inputs were untouched; nothing was committed.

`gofmt` passed. I could not run the Go evals here because the supplied workspace has no `go.mod` or dependency implementations, so the fixes remain unverified by the native suite.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1999f-dc19-7600-8fda-cdc767d6f765
--------
user
# Rebuild project

Target spec and Decisions:
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
Project and Stamp MUST resolve a change argument from its unique slug, queued basename or repository-relative queued path, rejecting missing or ambiguous targets.

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
Check MUST reject Modified or Removed entries whose bases hash is missing or differs from the preceding projected spec's canonical hash with diagnostic C08.

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

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module.

### PRJ-026
NewChange MUST reject an empty slug or a slug containing characters other than lowercase letters, digits and hyphens.

### PRJ-027
Queue MUST return pending NNNN-slug.md changes in numeric-then-slug order, allowing tied numbers and excluding archives.

### PRJ-028
Queue MUST return each pending change's module, frontmatter and parsed Added, Modified and Removed items.

### PRJ-029
Queue MUST reject a malformed pending change document.

### PRJ-030
Queue MUST reject a pending change naming an unknown module.

### PRJ-031
Queue MUST reject pending changes that share a slug.

### PRJ-032
Project MUST apply the queue through the named change inclusively, or the whole queue when through is empty.

### PRJ-033
Project MUST apply set values for depends_on, interface, evals, disposable, preserve, append_only, boundaries_opt_out and hooks to the projected module.

### PRJ-034
Project MUST reject identity or unknown fields in set.

### PRJ-035
Project MUST reject an invalid projected module graph.

### PRJ-036
Project MUST keep Added handles local to their declaring change, allowing different changes to reuse handle text.

### PRJ-037
Stamp MUST reject Modified or Removed IDs absent from the projection preceding the target change.

### PRJ-038
Stamp MUST retain other frontmatter values and human section contents.

### PRJ-039
Stamp MUST succeed on an Added-only change without assigning spec IDs.

### PRJ-040
Check MUST report configuration discovery or decoding failures with diagnostic C01.

### PRJ-041
Check MUST report invalid current or projected module graphs with diagnostic C02.

### PRJ-042
Check MUST report unreadable or malformed module/change documents and malformed item syntax with diagnostic C03, retaining C05 for duplicate identities.

### PRJ-043
Check MUST report invalid module metadata or zone field types, and invalid set keys or value types, with diagnostic C12.

### PRJ-044
Check MUST report pending queue identity errors with diagnostic C13.

### PRJ-045
Check MUST report invalid pending change module targets with diagnostic C07.

### PRJ-046
Check MUST report error C07 for a pending change containing eval_fix together
with fixes or any Added, Modified or Removed spec edits.

### PRJ-047
LoadAt MUST return the selected Git revision's project configuration and modules, using its recorded fireweed.yaml module globs and module SPEC.md documents rather than working-tree contents.

### PRJ-048
LoadAt MUST apply the selected revision's module_defaults to omitted zone fields while preserving explicitly empty module zone lists.

### PRJ-049
LoadAt MUST set Snapshot.Root to the canonical repository root and Module.Dir to each module's repository-relative directory.

### PRJ-050
LoadAt MUST return an error when the revision cannot be resolved, its recorded fireweed.yaml is missing or malformed, or a matched recorded module directory lacks a readable SPEC.md.

### PRJ-051
For a nonempty ID, SpecHistory MUST return the target module's archived changes whose Added or Modified items name that ID, including tag-only modifications.
Pending changes and archives belonging to other modules or unrelated spec IDs MUST not appear in that selection.

### PRJ-052
Every history entry MUST identify its archive path and module, the matching IDs, its Motivation, added and removed Decisions, and its complete Provenance section.

### PRJ-053
SpecHistory MUST order entries by their latest valid implementation Provenance timestamp, newest first, with missing or invalid implementation dates after dated entries.
Equal-date or undated entries MUST use archive-path lexical order as a total-order fallback.

### PRJ-054
With an empty ID filter, SpecHistory MUST return each target-module archive containing Added or Modified spec edits once, with its unique affected IDs in lexical order.
A module with no matching archives MUST return an empty slice.

### PRJ-055
An unknown module or an ID absent from both current module Specs and matching archive edits MUST return U02 with the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce a C01 error.
Malformed archived change documents or malformed Added/Modified item syntax MUST produce a C03 error identifying the archive path.

### PRJ-057
Stamp MUST return an error whose contracts ErrorCode is C08 when the change's bases frontmatter is not a mapping.

### PRJ-058
Stamp MUST return an error whose contracts ErrorCode is U02, with the requested change as subject, when no pending change matches it.

### PRJ-059
Load MUST return an error preserving the contracts ErrorCode C03 or C05 when a module SPEC.md fails to parse.

### PRJ-060
StaleSpecs MUST return the target change's Modified and Removed items whose recorded bases differ from the corresponding preceding projected spec hashes, in edit order, with ID, RecordedHash and Current text.
Items without a bases entry MUST be omitted.

### PRJ-061
StaleSpecs MUST recover Before by projecting the spec immediately before the target change at the Git commit that last changed that change file's bases mapping.
It MUST set Recovered true only when the recovered item's normalized SpecHash equals RecordedHash.

### PRJ-062
When matching stamped history cannot be recovered, StaleSpecs MUST return the stale item's Current text and RecordedHash with Recovered false and empty Before.
Unavailable history MUST NOT prevent reporting other stale items.

### PRJ-063
StaleSpecs MUST leave repository files, index and refs unchanged.
It MUST return an empty slice when no recorded bases are stale.

### PRJ-064
StampMissing MUST add canonical preceding-projection hashes only for Modified or Removed IDs absent from bases, preserving every existing entry including stale hashes.
It MUST retain other frontmatter values and human section contents.

### PRJ-065
StampMissing MUST reject a nonmapping bases value or an edited ID absent from the preceding projection with C08 and preserve the target file on rejection.

### PRJ-066
StaleSpecs and StampMissing MUST accept a unique pending slug, queued basename or repository-relative queued path.
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.


Retained interfaces and files:
packages/project/SPEC.md
packages/project/api.go
packages/project/model/model.go

Retained evals:
packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go

Rebuild only target disposable paths. Do not edit retained inputs. Do not read outside the supplied input workspace, external trees, old code, source history, binaries, saved attempts, or dependent internal implementation. Do not repair dependent internals. Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module. Do not commit.


# Previous verification failed
project eval: E01: failed case "Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults" for PRJ-001
app eval: E01: failed case "Test_APP_019_NewChangeErrorsPreserveFiles" for APP-019
workflow eval: E01: failed case "Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08" for FLOW-058
project stdout: ?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/project/evals	1.866s

project stderr: exit status 1

project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults: === RUN   Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults
    project_test.go:54: Root = "/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults2319314990/001"; want "/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults2319314990/001"
--- FAIL: Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults (0.00s)
project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics: === RUN   Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics
--- FAIL: Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics (0.03s)
project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_change_id: === RUN   Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_change_id
    queue_test.go:362: Check did not report C05 for foreign module ID: []model.Diagnostic{model.Diagnostic{Code:"C05", Message:"foreign spec ID BET-001 belongs to beta", Path:"changes/0001-check.md", Warning:false}, model.Diagnostic{Code:"C08", Message:"ID BET-001 absent from preceding projection", Path:"changes/0001-check.md", Warning:false}, model.Diagnostic{Code:"C08", Message:"Modified ID BET-001 absent from preceding projection", Path:"changes/0001-check.md", Warning:false}, model.Diagnostic{Code:"C08", Message:"C08 changes/0001-check.md: Modified ID BET-001 absent from preceding projection", Path:"changes/0001-check.md", Warning:false}}
--- FAIL: Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_change_id (0.00s)
project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_013_CheckDuplicateWithinEditSection: === RUN   Test_PRJ_013_CheckDuplicateWithinEditSection
    static_codes_test.go:206: Check did not report C05 for duplicate Added identity cause "C03 changes/0001-duplicate.md: duplicate item ID: new:local"
--- FAIL: Test_PRJ_013_CheckDuplicateWithinEditSection (0.00s)

app stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	237.575s

app stderr: exit status 1

app github.com/lurkshark/fireweed/packages/app/evals Test_APP_107_APP_108_CheckCommitsFailures: === RUN   Test_APP_107_APP_108_CheckCommitsFailures
    check_commits_cli_test.go:269: coded operational fixture: want malformed recorded SPEC, got C03 packages/alpha/SPEC.md: invalid YAML frontmatter: yaml: line 1: did not find expected ',' or ']'
--- FAIL: Test_APP_107_APP_108_CheckCommitsFailures (0.35s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration: === RUN   Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration
    ci_cli_test.go:190: CI C05 diagnostic lacks corrective action: "C05 packages/beta/SPEC.md: duplicate item ID: BET-001. Correct fireweed.yaml and module declarations, then retry ci."
--- FAIL: Test_APP_109_APP_112_APP_113_CIStaticUsageAndConfiguration (0.26s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses: === RUN   Test_APP_037_StaticCauses
--- FAIL: Test_APP_037_StaticCauses (0.64s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/graph: === RUN   Test_APP_037_StaticCauses/graph
    diagnostic_causes_test.go:105: missing C02 cause "missing module" in plain output "C02: module beta depends on missing dependency missing. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.\n"
    diagnostic_causes_test.go:105: missing C02 cause "missing module" in {"schema":1,"data":null,"diagnostics":[{"code":"C02","message":"module beta depends on missing dependency missing. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/graph (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/identity: === RUN   Test_APP_037_StaticCauses/identity
    diagnostic_causes_test.go:105: missing C05 cause "duplicate item ID" in plain output "C03: C03 changes/0001-first.md: duplicate item ID: new:one\n"
    diagnostic_causes_test.go:105: missing C05 cause "duplicate item ID" in {"schema":1,"data":null,"diagnostics":[{"code":"C03","message":"C03 changes/0001-first.md: duplicate item ID: new:one","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/identity (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/queue-identity: === RUN   Test_APP_037_StaticCauses/queue-identity
    diagnostic_causes_test.go:105: missing C13 cause "ambiguous change slug" in plain output "C13: C13 changes/0002-repeat.md: duplicate pending slug repeat\n"
    diagnostic_causes_test.go:105: missing C13 cause "ambiguous change slug" in {"schema":1,"data":null,"diagnostics":[{"code":"C13","message":"C13 changes/0002-repeat.md: duplicate pending slug repeat","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/queue-identity (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses: === RUN   Test_APP_038_LookupCauses
--- FAIL: Test_APP_038_LookupCauses (0.55s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses/stamp-change: === RUN   Test_APP_038_LookupCauses/stamp-change
    diagnostic_causes_test.go:156: missing U02 cause "not found" in plain output "U02: U02 absent: unknown change absent\n"
    diagnostic_causes_test.go:156: missing U02 cause "not found" in {"schema":1,"data":null,"diagnostics":[{"code":"U02","message":"U02 absent: unknown change absent","warning":false}]}
--- FAIL: Test_APP_038_LookupCauses/stamp-change (0.02s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses/coverage-change: === RUN   Test_APP_038_LookupCauses/coverage-change
    diagnostic_causes_test.go:156: missing U02 cause "not found" in plain output "U02: U02 missing: unknown change missing. Use a valid queued change alias.\n"
    diagnostic_causes_test.go:156: missing U02 cause "not found" in {"schema":1,"data":null,"diagnostics":[{"code":"U02","message":"U02 missing: unknown change missing. Use a valid queued change alias.","warning":false}]}
--- FAIL: Test_APP_038_LookupCauses/coverage-change (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_039_BlockedNewChangeOutputPath: === RUN   Test_APP_039_BlockedNewChangeOutputPath
    diagnostic_causes_test.go:183: lost create-file cause or correction: "P01: duplicate change slug blocked. Use a valid module and an unused lowercase slug.\n"
    diagnostic_causes_test.go:183: lost create-file cause or correction: "duplicate change slug blocked. Use a valid module and an unused lowercase slug."
--- FAIL: Test_APP_039_BlockedNewChangeOutputPath (0.47s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_065_APP_066_APP_067_DoctorFailures: === RUN   Test_APP_065_APP_066_APP_067_DoctorFailures
    doctor_setup_test.go:272: APP-067 static diagnostic=map[code:"C05" message:"C05 packages/alpha/SPEC.md: duplicate item ID: ALP-001" warning:false]
--- FAIL: Test_APP_065_APP_066_APP_067_DoctorFailures (0.77s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors: === RUN   Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors
    list_cli_test.go:193: public Queue fixture cause: C13 changes/0002-repeat.md: duplicate pending slug repeat
--- FAIL: Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors (0.00s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles (0.54s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain
    project_commands_test.go:397: diagnostic="P01: invalid slug \"Bad\". Use a valid module and an unused lowercase slug.\n"; want corrective U01
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "U01"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain
    project_commands_test.go:397: diagnostic="P01: duplicate change slug repeat. Use a valid module and an unused lowercase slug.\n"; want corrective C13
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain
    project_commands_test.go:397: diagnostic="P01: duplicate change slug repeat. Use a valid module and an unused lowercase slug.\n"; want corrective C13
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures (0.53s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain
    show_history_cli_test.go:296: status=1; want 2; stdout="" stderr="P01: U02 absent-module: U02 unknown module absent-module. Use a valid archive and correct the reported cause, then retry show --history.\n"
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: U02 absent-module: U02 unknown module absent-module. Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and U02 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/plain (0.24s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON
    show_history_cli_test.go:296: status=1; want 2; stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"U02 absent-module: U02 unknown module absent-module. Use a valid archive and correct the reported cause, then retry show --history.\",\"warning\":false}]}\n" stderr=""
    show_history_cli_test.go:296: diagnostic code="P01"; want "U02"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-module/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain
    show_history_cli_test.go:296: status=1; want 2; stdout="" stderr="P01: U02 ALP-999: U02 unknown spec ID ALP-999. Use a valid archive and correct the reported cause, then retry show --history.\n"
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: U02 ALP-999: U02 unknown spec ID ALP-999. Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and U02 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON
    show_history_cli_test.go:296: status=1; want 2; stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"U02 ALP-999: U02 unknown spec ID ALP-999. Use a valid archive and correct the reported cause, then retry show --history.\",\"warning\":false}]}\n" stderr=""
    show_history_cli_test.go:296: diagnostic code="P01"; want "U02"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/unknown-ID/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain
    show_history_cli_test.go:296: status=1; want 2; stdout="" stderr="P01: C01 fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Use a valid archive and correct the reported cause, then retry show --history.\n"
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: C01 fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and C01 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON
    show_history_cli_test.go:296: status=1; want 2; stdout="{\"schema\":1,\"data\":null,\"diagnostics\":[{\"code\":\"P01\",\"message\":\"C01 fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Use a valid archive and correct the reported cause, then retry show --history.\",\"warning\":false}]}\n" stderr=""
    show_history_cli_test.go:296: diagnostic code="P01"; want "C01"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-config/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain
    show_history_cli_test.go:296: plain failure stdout="" stderr="P01: C03 changes/archive/0001-invalid.md: malformed item heading: \"### ALP-001 [broken\". Use a valid archive and correct the reported cause, then retry show --history.\n"; want empty stdout and C03 diagnostic
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON: === RUN   Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON
    show_history_cli_test.go:296: diagnostic code="P01"; want "C03"
--- FAIL: Test_APP_135_ShowHistoryLookupAndArchiveFailures/malformed-archive/JSON (0.01s)

checks stdout: ?   	github.com/lurkshark/fireweed/packages/checks	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/internal/engine	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/checks/evals	225.150s

checks stderr: 

runner stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	235.144s

runner stderr: 

workflow stdout: ?   	github.com/lurkshark/fireweed/packages/workflow	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/workflow/evals	103.825s

workflow stderr: exit status 1

workflow github.com/lurkshark/fireweed/packages/workflow/evals Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08: === RUN   Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08
    lifecycle_test.go:382: missing-base static cause = P05: C08: stale base for ALP-001
--- FAIL: Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08 (0.81s)

codex
I’ll inspect the retained project interfaces and tests, then rebuild the disposable project implementation and run the relevant evals. I’ll keep changes inside the supplied workspace and leave retained files untouched.
exec
/bin/zsh -lc 'git status --short && rg --files packages/project | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "pwd && rg --files -g 'AGENTS.md' -g 'packages/project/**' -g 'packages/contracts/**' -g 'fireweed.yaml' -g 'DRILL_NOTES.md' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 502ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
DRILL_NOTES.md
packages/contracts/api.go
packages/contracts/model/model.go
packages/project/SPEC.md
packages/project/api.go
packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go
packages/project/internal/fixture/fixture.go
packages/project/internal/history/history.go
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go
packages/project/model/model.go

 succeeded in 503ms:
?? DRILL_NOTES.md
?? packages/project/internal/
packages/project/SPEC.md
packages/project/api.go
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go
packages/project/internal/fixture/fixture.go
packages/project/internal/history/history.go
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go
packages/project/model/model.go

exec
/bin/zsh -lc 'cat packages/project/api.go packages/project/model/model.go packages/contracts/api.go packages/contracts/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'wc -l packages/project/internal/*/*.go packages/project/evals/*.go DRILL_NOTES.md && cat DRILL_NOTES.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "rg -n 'func |type ' packages/project/internal -g '*.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 564ms:
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

 succeeded in 566ms:
packages/project/internal/queue/queue.go:16:func parseChange(path string, b []byte) (model.Change, error) {
packages/project/internal/queue/queue.go:48:func allEdits(c model.Change) []contracts.Item {
packages/project/internal/queue/queue.go:54:func queue(root string, s model.Snapshot, allowDuplicates bool) ([]model.Change, error) {
packages/project/internal/queue/queue.go:102:func Queue(root string) ([]model.Change, error) {
packages/project/internal/queue/queue.go:109:func resolve(changes []model.Change, target string) (int, error) {
packages/project/internal/queue/queue.go:124:func Project(root, through string) (model.Snapshot, error) {
packages/project/internal/queue/queue.go:151:func preceding(root, target string) (model.Snapshot, model.Change, error) {
packages/project/internal/queue/queue.go:171:func findModule(s *model.Snapshot, name string) *model.Module {
packages/project/internal/queue/queue.go:179:func SpecItems(m model.Module) ([]contracts.Item, error) {
packages/project/internal/queue/queue.go:183:func Before(s model.Snapshot, c model.Change) (map[string]contracts.Item, error) {
packages/project/internal/queue/queue.go:198:func Apply(s *model.Snapshot, c model.Change) error {
packages/project/internal/queue/queue.go:249:func applySet(m *model.Module, c model.Change) error {
packages/project/internal/queue/queue.go:300:func bases(c model.Change) (map[string]any, error) {
packages/project/internal/queue/queue.go:311:func Stamp(root, target string) error        { return stamp(root, target, false) }
packages/project/internal/queue/queue.go:312:func StampMissing(root, target string) error { return stamp(root, target, true) }
packages/project/internal/queue/queue.go:313:func stamp(root, target string, missingOnly bool) error {
packages/project/internal/queue/queue.go:356:func Check(root string) []model.Diagnostic {
packages/project/internal/queue/stale.go:15:func StaleSpecs(root, target string) ([]model.StaleSpec, error) {
packages/project/internal/queue/stale.go:60:func recoverBefore(root string, c model.Change) map[string]contracts.Item {
packages/project/internal/queue/stale.go:150:func equalBases(a, b map[string]any) bool {
packages/project/internal/queue/stale.go:161:func sortChanges(cs []model.Change) {
packages/project/internal/history/history.go:16:func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
packages/project/internal/history/history.go:125:func implDate(prov string) time.Time {
packages/project/internal/fixture/fixture.go:18:func Coded(code, subject string, err error) error {
packages/project/internal/fixture/fixture.go:21:func Code(err error) string {
packages/project/internal/fixture/fixture.go:27:func Canonical(root string) (string, error) {
packages/project/internal/fixture/fixture.go:34:func ReadConfig(root string) (map[string]any, error) {
packages/project/internal/fixture/fixture.go:41:func DecodeConfig(b []byte) (map[string]any, error) {
packages/project/internal/fixture/fixture.go:52:func Strings(v any) ([]string, error) {
packages/project/internal/fixture/fixture.go:67:func Patterns(c map[string]any) ([]string, error) {
packages/project/internal/fixture/fixture.go:74:func Defaults(c map[string]any) (map[string]any, error) {
packages/project/internal/fixture/fixture.go:88:func Metadata(doc contracts.Document, defaults map[string]any, dir string) (model.Module, error) {
packages/project/internal/fixture/fixture.go:135:func ValidateGraph(modules []model.Module) error {
packages/project/internal/fixture/fixture.go:193:func Load(root string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:223:func loadModules(root string, config, defaults map[string]any, dirs map[string]bool, read func(string) ([]byte, error)) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:261:func LoadAt(root, revision string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:268:func loadAt(root, revision string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:300:func TreePaths(root, revision string) ([]string, error) {
packages/project/internal/fixture/fixture.go:314:func ChangeName(name string) (int, string, bool) {
packages/project/internal/fixture/fixture.go:322:func NewChange(root, module, slug string) (string, error) {

 succeeded in 568ms:
     370 packages/project/internal/fixture/fixture.go
     138 packages/project/internal/history/history.go
     477 packages/project/internal/queue/queue.go
     167 packages/project/internal/queue/stale.go
      64 packages/project/evals/coded_errors_test.go
      45 packages/project/evals/mixed_eval_fix_test.go
     229 packages/project/evals/project_test.go
     411 packages/project/evals/queue_test.go
     171 packages/project/evals/revision_snapshot_test.go
     230 packages/project/evals/spec_history_test.go
     293 packages/project/evals/stale_test.go
     207 packages/project/evals/static_codes_test.go
       7 DRILL_NOTES.md
    2809 total
## Contradicted Decisions

None identified.

## Gaps

- PRJ-007 (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite could not run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.

exec
/bin/zsh -lc 'cat packages/project/internal/queue/queue.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/internal/fixture/fixture.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "ls -la && rg --files -g 'go.mod' -g 'go.work' -g 'AGENTS.md' -g 'packages/contracts/internal/**' | head -60" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 624ms:
package queue

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func parseChange(path string, b []byte) (model.Change, error) {
	doc, e := contracts.ParseDocument(string(b))
	if e != nil {
		return model.Change{}, fixture.Coded("C03", path, e)
	}
	n, slug, _ := fixture.ChangeName(filepath.Base(path))
	c := model.Change{Number: n, Slug: slug, Path: path, Document: doc}
	m, ok := doc.Frontmatter["module"].(string)
	if !ok {
		return c, fixture.Coded("C07", path, fmt.Errorf("invalid module target"))
	}
	c.Module = m
	for _, part := range []struct {
		name string
		dest *[]contracts.Item
	}{{"Added", &c.Added}, {"Modified", &c.Modified}, {"Removed", &c.Removed}} {
		body, _ := contracts.Section(doc.Body, part.name)
		items, e := contracts.ParseItems(body)
		if e != nil {
			return c, fixture.Coded("C03", path, e)
		}
		*part.dest = items
	}
	seen := map[string]bool{}
	for _, it := range allEdits(c) {
		if seen[it.ID] {
			return c, fixture.Coded("C05", path, fmt.Errorf("duplicate edit identity %s", it.ID))
		}
		seen[it.ID] = true
	}
	return c, nil
}
func allEdits(c model.Change) []contracts.Item {
	out := append([]contracts.Item{}, c.Added...)
	out = append(out, c.Modified...)
	out = append(out, c.Removed...)
	return out
}
func queue(root string, s model.Snapshot, allowDuplicates bool) ([]model.Change, error) {
	ents, e := os.ReadDir(filepath.Join(s.Root, "changes"))
	if os.IsNotExist(e) {
		return []model.Change{}, nil
	}
	if e != nil {
		return nil, fixture.Coded("C03", "changes", e)
	}
	known := map[string]bool{}
	for _, m := range s.Modules {
		known[m.Name] = true
	}
	out := []model.Change{}
	slugs := map[string]bool{}
	for _, ent := range ents {
		if ent.IsDir() {
			continue
		}
		_, slug, ok := fixture.ChangeName(ent.Name())
		if !ok {
			continue
		}
		path := "changes/" + ent.Name()
		if slugs[slug] && !allowDuplicates {
			return nil, fixture.Coded("C13", path, fmt.Errorf("duplicate pending slug %s", slug))
		}
		slugs[slug] = true
		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
		if e != nil {
			return nil, fixture.Coded("C03", path, e)
		}
		c, e := parseChange(path, b)
		if e != nil {
			return nil, e
		}
		if !known[c.Module] {
			return nil, fixture.Coded("C07", path, fmt.Errorf("unknown change module %s", c.Module))
		}
		out = append(out, c)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Number != out[j].Number {
			return out[i].Number < out[j].Number
		}
		return out[i].Slug < out[j].Slug
	})
	return out, nil
}
func Queue(root string) ([]model.Change, error) {
	s, e := fixture.Load(root)
	if e != nil {
		return nil, e
	}
	return queue(root, s, false)
}
func resolve(changes []model.Change, target string) (int, error) {
	matches := []int{}
	for i, c := range changes {
		if target == c.Slug || target == filepath.Base(c.Path) || target == c.Path {
			matches = append(matches, i)
		}
	}
	if len(matches) == 0 {
		return -1, fixture.Coded("U02", target, fmt.Errorf("unknown change %s", target))
	}
	if len(matches) > 1 {
		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change %s", target))
	}
	return matches[0], nil
}
func Project(root, through string) (model.Snapshot, error) {
	s, e := fixture.Load(root)
	if e != nil {
		return s, e
	}
	changes, e := queue(root, s, false)
	if e != nil {
		return s, e
	}
	limit := len(changes)
	if through != "" {
		idx, e := resolve(changes, through)
		if e != nil {
			return s, e
		}
		limit = idx + 1
	}
	for _, c := range changes[:limit] {
		if e := Apply(&s, c); e != nil {
			return s, e
		}
	}
	if e := fixture.ValidateGraph(s.Modules); e != nil {
		return s, fixture.Coded("C02", "graph", e)
	}
	return s, nil
}
func preceding(root, target string) (model.Snapshot, model.Change, error) {
	s, e := fixture.Load(root)
	if e != nil {
		return s, model.Change{}, e
	}
	changes, e := queue(root, s, true)
	if e != nil {
		return s, model.Change{}, e
	}
	idx, e := resolve(changes, target)
	if e != nil {
		return s, model.Change{}, e
	}
	for _, c := range changes[:idx] {
		if e := Apply(&s, c); e != nil {
			return s, model.Change{}, e
		}
	}
	return s, changes[idx], nil
}
func findModule(s *model.Snapshot, name string) *model.Module {
	for i := range s.Modules {
		if s.Modules[i].Name == name {
			return &s.Modules[i]
		}
	}
	return nil
}
func SpecItems(m model.Module) ([]contracts.Item, error) {
	body, _ := contracts.Section(m.Document.Body, "Specs")
	return contracts.ParseItems(body)
}
func Before(s model.Snapshot, c model.Change) (map[string]contracts.Item, error) {
	m := findModule(&s, c.Module)
	if m == nil {
		return nil, fixture.Coded("C07", c.Path, fmt.Errorf("unknown change module %s", c.Module))
	}
	items, e := SpecItems(*m)
	if e != nil {
		return nil, fixture.Coded("C03", m.Dir+"/SPEC.md", e)
	}
	out := map[string]contracts.Item{}
	for _, it := range items {
		out[it.ID] = it
	}
	return out, nil
}
func Apply(s *model.Snapshot, c model.Change) error {
	m := findModule(s, c.Module)
	if m == nil {
		return fixture.Coded("C07", c.Path, fmt.Errorf("unknown change module %s", c.Module))
	}
	if e := applySet(m, c); e != nil {
		return e
	}
	items, e := SpecItems(*m)
	if e != nil {
		return fixture.Coded("C03", m.Dir+"/SPEC.md", e)
	}
	positions := map[string]int{}
	for i, it := range items {
		positions[it.ID] = i
	}
	for _, it := range c.Modified {
		idx, ok := positions[it.ID]
		if !ok {
			return fixture.Coded("C08", c.Path, fmt.Errorf("Modified ID %s absent from preceding projection", it.ID))
		}
		items[idx] = it
	}
	remove := map[string]bool{}
	for _, it := range c.Removed {
		if _, ok := positions[it.ID]; !ok {
			return fixture.Coded("C08", c.Path, fmt.Errorf("Removed ID %s absent from preceding projection", it.ID))
		}
		remove[it.ID] = true
	}
	kept := []contracts.Item{}
	for _, it := range items {
		if !remove[it.ID] {
			kept = append(kept, it)
		}
	}
	kept = append(kept, c.Added...)
	var body strings.Builder
	for _, it := range kept {
		rendered := contracts.RenderItem(it)
		body.WriteString(rendered)
		if !strings.HasSuffix(rendered, "\n\n") {
			body.WriteString("\n")
		}
	}
	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
	return nil
}

var listKeys = map[string]bool{"depends_on": true, "interface": true, "evals": true, "disposable": true, "preserve": true, "append_only": true}

func applySet(m *model.Module, c model.Change) error {
	raw, exists := c.Document.Frontmatter["set"]
	if !exists {
		return nil
	}
	set, ok := raw.(map[string]any)
	if !ok {
		return fixture.Coded("C12", c.Path, fmt.Errorf("set must be a mapping"))
	}
	for k, v := range set {
		if listKeys[k] {
			if _, e := fixture.Strings(v); e != nil {
				return fixture.Coded("C12", c.Path, fmt.Errorf("set %s: %w", k, e))
			}
		} else {
			switch k {
			case "boundaries_opt_out":
				if _, ok := v.(string); !ok {
					return fixture.Coded("C12", c.Path, fmt.Errorf("set %s must be a string", k))
				}
			case "hooks":
				if _, ok := v.(map[string]any); !ok {
					return fixture.Coded("C12", c.Path, fmt.Errorf("set hooks must be a mapping"))
				}
			default:
				return fixture.Coded("C12", c.Path, fmt.Errorf("invalid set key %s", k))
			}
		}
	}
	for k, v := range set {
		m.Document.Frontmatter[k] = v
		if listKeys[k] {
			list, _ := fixture.Strings(v)
			switch k {
			case "depends_on":
				m.DependsOn = list
			case "interface":
				m.Zones.Interface = list
			case "evals":
				m.Zones.Evals = list
			case "disposable":
				m.Zones.Disposable = list
			case "preserve":
				m.Zones.Preserve = list
			case "append_only":
				m.Zones.AppendOnly = list
			}
		}
	}
	return nil
}
func bases(c model.Change) (map[string]any, error) {
	v, ok := c.Document.Frontmatter["bases"]
	if !ok {
		return map[string]any{}, nil
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, fixture.Coded("C08", c.Path, fmt.Errorf("bases must be a mapping"))
	}
	return m, nil
}
func Stamp(root, target string) error        { return stamp(root, target, false) }
func StampMissing(root, target string) error { return stamp(root, target, true) }
func stamp(root, target string, missingOnly bool) error {
	s, c, e := preceding(root, target)
	if e != nil {
		return e
	}
	old, e := Before(s, c)
	if e != nil {
		return e
	}
	b, e := bases(c)
	if e != nil {
		return e
	}
	for _, part := range []struct {
		name  string
		items []contracts.Item
	}{{"Modified", c.Modified}, {"Removed", c.Removed}} {
		for _, it := range part.items {
			prior, ok := old[it.ID]
			if !ok {
				return fixture.Coded("C08", c.Path, fmt.Errorf("%s ID %s absent from preceding projection", part.name, it.ID))
			}
			if _, exists := b[it.ID]; !exists || !missingOnly {
				b[it.ID] = contracts.SpecHash(prior)
			}
		}
	}
	c.Document.Frontmatter["bases"] = b
	rendered, e := contracts.RenderDocument(c.Document)
	if e != nil {
		return fixture.Coded("C03", c.Path, e)
	}
	path := filepath.Join(s.Root, filepath.FromSlash(c.Path))
	info, e := os.Stat(path)
	if e != nil {
		return e
	}
	return os.WriteFile(path, []byte(rendered), info.Mode().Perm())
}

var handleRE = regexp.MustCompile(`\bnew:[a-zA-Z0-9_-]+\b`)
var normRE = regexp.MustCompile(`\b(MUST|SHALL|SHOULD|MAY|MUST NOT|SHALL NOT|SHOULD NOT)\b`)

func Check(root string) []model.Diagnostic {
	out := []model.Diagnostic{}
	add := func(code, msg, path string) {
		out = append(out, model.Diagnostic{Code: code, Message: msg, Path: path})
	}
	s, e := fixture.Load(root)
	if e != nil {
		add(fixture.Code(e), e.Error(), "")
		return out
	}
	if e := fixture.ValidateGraph(s.Modules); e != nil {
		add("C02", e.Error(), "")
	}
	prefixes := map[string]string{}
	for _, m := range s.Modules {
		prefixes[m.IDPrefix] = m.Name
	}
	checkItem := func(it contracts.Item, path, module string, needBody bool) {
		if strings.Contains(it.ID, "-") {
			prefix := strings.SplitN(it.ID, "-", 2)[0]
			if owner, ok := prefixes[prefix]; ok && owner != module {
				add("C05", fmt.Sprintf("foreign spec ID %s belongs to %s", it.ID, owner), path)
			}
		}
		if needBody && !normRE.MatchString(it.Body) {
			add("C04", "non-normative spec body for "+it.ID, path)
		}
		if it.Tag != "" && it.Tag != "review" && !strings.HasPrefix(it.Tag, "quarantine: ") {
			add("C06", "invalid review/quarantine tag for "+it.ID, path)
		}
		if strings.HasPrefix(it.Tag, "quarantine: ") && strings.TrimSpace(strings.TrimPrefix(it.Tag, "quarantine: ")) == "" {
			add("C06", "invalid review/quarantine tag for "+it.ID, path)
		}
	}
	for _, m := range s.Modules {
		path := m.Dir + "/SPEC.md"
		items, e := SpecItems(m)
		if e != nil {
			add("C03", e.Error(), path)
			continue
		}
		seen := map[string]bool{}
		for _, it := range items {
			if seen[it.ID] {
				add("C05", "duplicate item ID "+it.ID, path)
			}
			seen[it.ID] = true
			checkItem(it, path, m.Name, true)
		}
	}
	changes, e := queue(root, s, false)
	if e != nil {
		add(fixture.Code(e), e.Error(), "")
		return out
	}
	for _, c := range changes {
		mot, _ := contracts.Section(c.Document.Body, "Motivation")
		if strings.TrimSpace(mot) == "" {
			add("C07", "missing Motivation in "+c.Path, c.Path)
		}
		if _, ok := c.Document.Frontmatter["eval_fix"]; ok {
			if _, hasFixes := c.Document.Frontmatter["fixes"]; hasFixes || len(allEdits(c)) > 0 {
				add("C07", "eval_fix cannot be combined with fixes or spec edits", c.Path)
				continue
			}
		}
		prior, e := Before(s, c)
		if e != nil {
			add(fixture.Code(e), e.Error(), c.Path)
			continue
		}
		b, e := bases(c)
		if e != nil {
			add("C08", e.Error(), c.Path)
			continue
		}
		added := map[string]bool{}
		for _, it := range c.Added {
			if strings.HasPrefix(it.ID, "new:") {
				added[it.ID] = true
			}
		}
		for _, it := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
			checkItem(it, c.Path, c.Module, true)
		}
		for _, it := range c.Removed {
			checkItem(it, c.Path, c.Module, false)
		}
		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
			if !added[ref] {
				add("C09", "reference to undeclared handle "+ref, c.Path)
			}
		}
		for _, part := range []struct {
			name  string
			items []contracts.Item
		}{{"Modified", c.Modified}, {"Removed", c.Removed}} {
			for _, it := range part.items {
				p, ok := prior[it.ID]
				if !ok {
					add("C08", fmt.Sprintf("ID %s absent from preceding projection", it.ID), c.Path)
					add("C08", fmt.Sprintf("%s ID %s absent from preceding projection", part.name, it.ID), c.Path)
					continue
				}
				v, ok := b[it.ID]
				if !ok || v == "" {
					add("C08", "missing base for "+it.ID, c.Path)
				} else if v != contracts.SpecHash(p) {
					add("C08", "stale base for "+it.ID, c.Path)
				}
			}
		}
		if e := Apply(&s, c); e != nil {
			add(fixture.Code(e), e.Error(), c.Path)
			continue
		}
		if e := fixture.ValidateGraph(s.Modules); e != nil {
			add("C02", fixture.Coded("C02", "graph", e).Error(), c.Path)
		}
	}
	return out
}

 succeeded in 627ms:
package fixture

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func Coded(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s %s: %w", code, subject, err)}
}
func Code(err error) string {
	if c, _, ok := contracts.ErrorCode(err); ok {
		return c
	}
	return "C03"
}
func Canonical(root string) (string, error) {
	p, e := filepath.Abs(root)
	if e != nil {
		return "", e
	}
	return filepath.EvalSymlinks(p)
}
func ReadConfig(root string) (map[string]any, error) {
	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
	if e != nil {
		return nil, Coded("C01", "fireweed.yaml", e)
	}
	return DecodeConfig(b)
}
func DecodeConfig(b []byte) (map[string]any, error) {
	doc, e := contracts.ParseDocument("---\n" + string(b) + "\n---\n")
	if e != nil {
		return nil, Coded("C01", "fireweed.yaml", e)
	}
	c := doc.Frontmatter
	if _, e := Patterns(c); e != nil {
		return nil, Coded("C01", "fireweed.yaml", e)
	}
	return c, nil
}
func Strings(v any) ([]string, error) {
	a, ok := v.([]any)
	if !ok {
		return nil, fmt.Errorf("expected list of strings, got %T", v)
	}
	out := make([]string, 0, len(a))
	for _, x := range a {
		s, ok := x.(string)
		if !ok {
			return nil, fmt.Errorf("expected string in list, got %T", x)
		}
		out = append(out, s)
	}
	return out, nil
}
func Patterns(c map[string]any) ([]string, error) {
	v, ok := c["modules"]
	if !ok {
		return nil, fmt.Errorf("missing modules globs")
	}
	return Strings(v)
}
func Defaults(c map[string]any) (map[string]any, error) {
	v, ok := c["module_defaults"]
	if !ok {
		return map[string]any{}, nil
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("module_defaults must be a mapping")
	}
	return m, nil
}

var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}

func Metadata(doc contracts.Document, defaults map[string]any, dir string) (model.Module, error) {
	f := doc.Frontmatter
	name, ok := f["name"].(string)
	if !ok {
		return model.Module{}, Coded("C12", dir, fmt.Errorf("name must be a string"))
	}
	prefix, ok := f["id_prefix"].(string)
	if !ok {
		return model.Module{}, Coded("C12", dir, fmt.Errorf("id_prefix must be a string"))
	}
	m := model.Module{Name: name, IDPrefix: prefix, Dir: dir, Document: doc}
	if v, ok := f["depends_on"]; ok {
		d, e := Strings(v)
		if e != nil {
			return m, Coded("C12", dir, fmt.Errorf("depends_on: %w", e))
		}
		m.DependsOn = d
	}
	for _, key := range zoneKeys {
		v, ok := f[key]
		if !ok {
			v, ok = defaults[key]
		}
		if !ok {
			continue
		}
		list, e := Strings(v)
		if e != nil {
			return m, Coded("C12", dir, fmt.Errorf("%s: %w", key, e))
		}
		switch key {
		case "protected":
			m.Zones.Protected = list
		case "evals":
			m.Zones.Evals = list
		case "interface":
			m.Zones.Interface = list
		case "append_only":
			m.Zones.AppendOnly = list
		case "preserve":
			m.Zones.Preserve = list
		case "disposable":
			m.Zones.Disposable = list
		}
	}
	return m, nil
}
func ValidateGraph(modules []model.Module) error {
	names := map[string]bool{}
	prefixes := map[string]bool{}
	for _, m := range modules {
		if strings.TrimSpace(m.Name) == "" {
			return fmt.Errorf("empty module name")
		}
		if names[m.Name] {
			return fmt.Errorf("duplicate module name %s", m.Name)
		}
		names[m.Name] = true
		if strings.TrimSpace(m.IDPrefix) == "" {
			return fmt.Errorf("empty module prefix")
		}
		if prefixes[m.IDPrefix] {
			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
		}
		prefixes[m.IDPrefix] = true
	}
	for _, m := range modules {
		for _, d := range m.DependsOn {
			if d == m.Name {
				return fmt.Errorf("self dependency %s", m.Name)
			}
			if !names[d] {
				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
			}
		}
	}
	edges := map[string][]string{}
	for _, m := range modules {
		edges[m.Name] = m.DependsOn
	}
	state := map[string]int{}
	var visit func(string) error
	visit = func(n string) error {
		if state[n] == 1 {
			return fmt.Errorf("dependency cycle involving %s", n)
		}
		if state[n] == 2 {
			return nil
		}
		state[n] = 1
		for _, d := range edges[n] {
			if e := visit(d); e != nil {
				return e
			}
		}
		state[n] = 2
		return nil
	}
	for _, m := range modules {
		if e := visit(m.Name); e != nil {
			return e
		}
	}
	return nil
}
func Load(root string) (model.Snapshot, error) {
	canon, e := Canonical(root)
	if e != nil {
		return model.Snapshot{}, Coded("C01", root, e)
	}
	config, e := ReadConfig(canon)
	if e != nil {
		return model.Snapshot{}, e
	}
	patterns, _ := Patterns(config)
	defaults, e := Defaults(config)
	if e != nil {
		return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
	}
	dirs := map[string]bool{}
	for _, p := range patterns {
		matches, e := filepath.Glob(filepath.Join(canon, filepath.FromSlash(p)))
		if e != nil {
			return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
		}
		for _, match := range matches {
			info, e := os.Stat(match)
			if e == nil && info.IsDir() {
				rel, _ := filepath.Rel(canon, match)
				dirs[filepath.ToSlash(rel)] = true
			}
		}
	}
	return loadModules(canon, config, defaults, dirs, func(path string) ([]byte, error) { return os.ReadFile(filepath.Join(canon, filepath.FromSlash(path))) })
}
func loadModules(root string, config, defaults map[string]any, dirs map[string]bool, read func(string) ([]byte, error)) (model.Snapshot, error) {
	s := model.Snapshot{Root: root, Config: config, Modules: []model.Module{}}
	keys := make([]string, 0, len(dirs))
	for d := range dirs {
		keys = append(keys, d)
	}
	sort.Strings(keys)
	for _, dir := range keys {
		path := dir + "/SPEC.md"
		b, e := read(path)
		if e != nil {
			return s, Coded("C03", path, e)
		}
		doc, e := contracts.ParseDocument(string(b))
		if e != nil {
			return s, Coded("C03", path, e)
		}
		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
			items, e := contracts.ParseItems(section)
			if e != nil {
				return s, Coded(Code(e), path, e)
			}
			seen := map[string]bool{}
			for _, it := range items {
				if seen[it.ID] {
					return s, Coded("C05", path, fmt.Errorf("duplicate item ID %s", it.ID))
				}
				seen[it.ID] = true
			}
		}
		m, e := Metadata(doc, defaults, dir)
		if e != nil {
			return s, e
		}
		s.Modules = append(s.Modules, m)
	}
	return s, nil
}
func LoadAt(root, revision string) (model.Snapshot, error) {
	canon, e := Canonical(root)
	if e != nil {
		return model.Snapshot{}, e
	}
	return loadAt(canon, revision)
}
func loadAt(root, revision string) (model.Snapshot, error) {
	b, e := gitstore.ReadAt(root, revision, "fireweed.yaml")
	if e != nil {
		return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
	}
	config, e := DecodeConfig(b)
	if e != nil {
		return model.Snapshot{}, e
	}
	patterns, _ := Patterns(config)
	defaults, e := Defaults(config)
	if e != nil {
		return model.Snapshot{}, Coded("C01", "fireweed.yaml", e)
	}
	paths, e := TreePaths(root, revision)
	if e != nil {
		return model.Snapshot{}, e
	}
	dirs := map[string]bool{}
	for _, path := range paths {
		parts := strings.Split(path, "/")
		for i := 1; i < len(parts); i++ {
			dir := strings.Join(parts[:i], "/")
			for _, p := range patterns {
				if contracts.MatchGlob(p, dir) {
					dirs[dir] = true
				}
			}
		}
	}
	return loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
}
func TreePaths(root, revision string) ([]string, error) {
	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "ls-tree", "-r", "--name-only", revision}})
	if e != nil {
		return nil, e
	}
	if r.ExitCode != 0 {
		return nil, fmt.Errorf("git ls-tree %s: %s", revision, r.Stderr)
	}
	return strings.Fields(r.Stdout), nil
}

var slugRE = regexp.MustCompile(`^[a-z0-9-]+$`)
var filenameRE = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

func ChangeName(name string) (int, string, bool) {
	m := filenameRE.FindStringSubmatch(name)
	if m == nil {
		return 0, "", false
	}
	n, _ := strconv.Atoi(m[1])
	return n, m[2], true
}
func NewChange(root, module, slug string) (string, error) {
	if !slugRE.MatchString(slug) {
		return "", fmt.Errorf("invalid slug %q", slug)
	}
	s, e := Load(root)
	if e != nil {
		return "", e
	}
	found := false
	for _, m := range s.Modules {
		if m.Name == module {
			found = true
		}
	}
	if !found {
		return "", fmt.Errorf("unknown module %s", module)
	}
	max := 0
	for _, dir := range []string{"changes", "changes/archive"} {
		entries, e := os.ReadDir(filepath.Join(s.Root, dir))
		if e != nil && !os.IsNotExist(e) {
			return "", e
		}
		for _, ent := range entries {
			n, sl, ok := ChangeName(ent.Name())
			if !ok {
				continue
			}
			if sl == slug {
				return "", fmt.Errorf("duplicate change slug %s", slug)
			}
			if n > max {
				max = n
			}
		}
	}
	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
	if e := os.MkdirAll(filepath.Join(s.Root, "changes"), 0755); e != nil {
		return "", e
	}
	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if e != nil {
		return "", e
	}
	defer f.Close()
	_, e = f.WriteString(body)
	return rel, e
}

 succeeded in 638ms:
total 16
drwx------@    6 joel  staff     192 Nov  1 16:05 .
drwx------@ 6898 joel  staff  220736 Nov  1 16:19 ..
drwxr-xr-x@    9 joel  staff     288 Nov  1 16:05 .git
-rw-r--r--@    1 joel  staff     288 Nov  1 16:05 DRILL_NOTES.md
-rw-r--r--@    1 joel  staff    2965 Nov  1 16:01 SYSTEM.md
drwxr-xr-x@    6 joel  staff     192 Nov  1 16:01 packages

exec
/bin/zsh -lc "sed -n '1,430p' packages/project/evals/queue_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "sed -n '1,210p' packages/project/evals/coded_errors_test.go && sed -n '1,210p' packages/project/evals/revision_snapshot_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "sed -n '1,230p' packages/project/evals/project_test.go && sed -n '1,225p' packages/project/evals/static_codes_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 576ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func queued(t *testing.T, root, filename, front, motivation, added, modified, removed string) {
	t.Helper()
	put(t, root, "changes/"+filename, "---\nmodule: alpha\n"+front+"---\n\n## Motivation\n"+motivation+"\n## Interface\n\n## Added\n"+added+"\n## Modified\n"+modified+"\n## Removed\n"+removed+"\n## Decisions\n\n## Notes\nHuman note.\n")
}

func moduleByName(t *testing.T, snapshot project.Snapshot, name string) project.Module {
	t.Helper()
	for _, module := range snapshot.Modules {
		if module.Name == name {
			return module
		}
	}
	t.Fatalf("module %q absent from projection", name)
	return project.Module{}
}

func specItems(t *testing.T, module project.Module) []contracts.Item {
	t.Helper()
	section, ok := contracts.Section(module.Document.Body, "Specs")
	if !ok {
		t.Fatalf("Specs section absent from %s", module.Name)
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		t.Fatalf("parse projected %s specs: %v", module.Name, err)
	}
	return items
}

func findItem(t *testing.T, items []contracts.Item, id string) contracts.Item {
	t.Helper()
	for _, item := range items {
		if item.ID == id {
			return item
		}
	}
	t.Fatalf("spec %s absent from %#v", id, items)
	return contracts.Item{}
}

func Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0002-zed.md", "set:\n  interface: [next.go]\n", "A queue item is needed.\n", "### new:feature\nFeature MUST work.\n", "### ALP-001\nAlpha MUST change.\n", "### ALP-002\n")
	queued(t, root, "0002-ace.md", "", "Another queue item is needed.\n", "", "", "")
	queued(t, root, "0001-first.md", "", "The first queue item is needed.\n", "", "", "")
	put(t, root, "changes/archive/0000-old.md", "archived material is not a pending document")
	changes, err := project.Queue(root)
	if err != nil {
		t.Fatalf("Queue: %v", err)
	}
	if len(changes) != 3 {
		t.Fatalf("queue = %#v; want three pending changes", changes)
	}
	for i, want := range []struct {
		number     int
		slug, path string
	}{
		{1, "first", "changes/0001-first.md"},
		{2, "ace", "changes/0002-ace.md"},
		{2, "zed", "changes/0002-zed.md"},
	} {
		got := changes[i]
		if got.Number != want.number || got.Slug != want.slug || got.Path != want.path || got.Module != "alpha" {
			t.Errorf("change %d = %#v; want %#v with module alpha", i, got, want)
		}
	}
	last := changes[2]
	if !strings.Contains(last.Document.Body, "A queue item is needed.") || last.Document.Frontmatter["module"] != "alpha" || last.Document.Frontmatter["set"] == nil {
		t.Errorf("change document/frontmatter lost: %#v", last.Document)
	}
	if len(last.Added) != 1 || last.Added[0].ID != "new:feature" || !strings.Contains(last.Added[0].Body, "Feature MUST work.") || len(last.Modified) != 1 || last.Modified[0].ID != "ALP-001" || !strings.Contains(last.Modified[0].Body, "Alpha MUST change.") || len(last.Removed) != 1 || last.Removed[0].ID != "ALP-002" {
		t.Errorf("parsed edits = added %#v, modified %#v, removed %#v", last.Added, last.Modified, last.Removed)
	}
}

func Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments(t *testing.T) {
	for _, tc := range []struct{ id, name, body string }{
		{"PRJ-029", "malformed", "---\nmodule: [broken\n---\n"},
		{"PRJ-030", "unknown_module", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "changes/0001-bad.md", tc.body)
			if _, err := project.Queue(root); err == nil {
				t.Error("Queue accepted invalid change")
			}
		})
	}
}

func Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:first\nFirst MUST work.\n", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:second\nSecond MUST work.\n", "", "")
	for _, target := range []string{"first", "0001-first.md", "changes/0001-first.md"} {
		projected, err := project.Project(root, target)
		if err != nil {
			t.Errorf("Project(%q): %v", target, err)
			continue
		}
		items := specItems(t, moduleByName(t, projected, "alpha"))
		if len(items) != 2 || findItem(t, items, "new:first").ID == "" {
			t.Errorf("Project(%q) did not stop at first change: %#v", target, items)
		}
		if err := project.Stamp(root, target); err != nil {
			t.Errorf("Stamp(%q): %v", target, err)
		}
	}
	if _, err := project.Project(root, "missing"); err == nil {
		t.Error("Project accepted missing target")
	}
	if err := project.Stamp(root, "missing"); err == nil {
		t.Error("Stamp accepted missing target")
	}
	put(t, root, "changes/archive/0000-archived.md", "archived document")
	if _, err := project.Project(root, "changes/archive/0000-archived.md"); err == nil {
		t.Error("Project accepted archived target")
	}
	queued(t, root, "0003-first.md", "", "Another first.\n", "", "", "")
	if _, err := project.Queue(root); err == nil {
		t.Error("Queue accepted ambiguous slug")
	}
	if _, err := project.Project(root, "first"); err == nil {
		t.Error("Project accepted ambiguous slug")
	}
	if err := project.Stamp(root, "first"); err == nil {
		t.Error("Stamp accepted ambiguous slug")
	}
}

func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "### new:shared\nFirst change MUST retain its own text.\n", "### ALP-001 [review]\nAlpha MUST have revised full text.\n", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "### new:shared\nSecond change MUST retain different text.\n", "", "### ALP-001\n")
	first, err := project.Project(root, "first")
	if err != nil {
		t.Fatalf("Project first: %v", err)
	}
	firstItems := specItems(t, moduleByName(t, first, "alpha"))
	if len(firstItems) != 2 || findItem(t, firstItems, "ALP-001").Tag != "review" || findItem(t, firstItems, "ALP-001").Body != "Alpha MUST have revised full text.\n\n" || findItem(t, firstItems, "new:shared").Body != "First change MUST retain its own text.\n\n" {
		t.Errorf("first projection = %#v", firstItems)
	}
	all, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project whole queue: %v", err)
	}
	section, ok := contracts.Section(moduleByName(t, all, "alpha").Document.Body, "Specs")
	if !ok {
		t.Fatal("Specs section absent from whole-queue projection")
	}
	if strings.Count(section, "### new:shared\n") != 2 ||
		!strings.Contains(section, "### new:shared\nFirst change MUST retain its own text.\n") ||
		!strings.Contains(section, "### new:shared\nSecond change MUST retain different text.\n") ||
		strings.Contains(section, "ALP-001") {
		t.Errorf("whole-queue projection lost local handles or retained removed ID: %q", section)
	}
}

func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: []\n---\n# Beta\n\n## Specs\n")
	queued(t, root, "0001-settings.md", "set:\n  depends_on: [beta]\n  interface: [next.go]\n  evals: [checks/**]\n  disposable: [generated/**]\n  preserve: [records/**]\n  append_only: [migrations/**]\n  boundaries_opt_out: reviewed boundary\n  hooks:\n    test:\n      command: [go, test, ./...]\n", "Settings are needed.\n", "", "", "")
	got, err := project.Project(root, "")
	if err != nil {
		t.Fatalf("Project settings: %v", err)
	}
	alpha := moduleByName(t, got, "alpha")
	wantLists := map[string][]string{
		"depends_on": {"beta"}, "interface": {"next.go"}, "evals": {"checks/**"},
		"disposable": {"generated/**"}, "preserve": {"records/**"}, "append_only": {"migrations/**"},
	}
	for key, want := range wantLists {
		value, ok := alpha.Document.Frontmatter[key].([]any)
		if !ok || len(value) != 1 || value[0] != want[0] {
			t.Errorf("set %s = %#v; want %q", key, alpha.Document.Frontmatter[key], want)
		}
	}
	if !reflect.DeepEqual(alpha.DependsOn, []string{"beta"}) || !reflect.DeepEqual(alpha.Zones.Interface, []string{"next.go"}) || !reflect.DeepEqual(alpha.Zones.Evals, []string{"checks/**"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"generated/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"records/**"}) || !reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"migrations/**"}) {
		t.Errorf("projected module fields not updated: %#v", alpha)
	}
	wantHooks := map[string]any{"test": map[string]any{"command": []any{"go", "test", "./..."}}}
	if alpha.Document.Frontmatter["boundaries_opt_out"] != "reviewed boundary" || !reflect.DeepEqual(alpha.Document.Frontmatter["hooks"], wantHooks) {
		t.Errorf("projected scalar/map settings lost: %#v", alpha.Document.Frontmatter)
	}
	for _, tc := range []struct{ id, name, settings string }{
		{"PRJ-034", "identity_name", "name: renamed"},
		{"PRJ-034", "identity_prefix", "id_prefix: NEW"},
		{"PRJ-034", "unknown_field", "surprise: true"},
		{"PRJ-035", "invalid_graph", "depends_on: [missing]"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			badRoot := fixture(t)
			queued(t, badRoot, "0001-invalid.md", "set:\n  "+tc.settings+"\n", "Needed.\n", "", "", "")
			if _, err := project.Project(badRoot, ""); err == nil {
				t.Errorf("Project accepted %s", tc.name)
			}
		})
	}
}

func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-revise.md", "", "Earlier change.\n", "", "### ALP-001 [review]\nAlpha MUST reflect the earlier revision.\n", "")
	queued(t, root, "0002-target.md", "custom: author-value\n", "Keep this human motivation.\n", "", "### ALP-001\nAlpha MUST reflect the later revision.\n", "")
	wantHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Tag: "review", Body: "Alpha MUST reflect the earlier revision.\n"})
	if err := project.Stamp(root, "changes/0002-target.md"); err != nil {
		t.Fatalf("Stamp: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(root, "changes/0002-target.md"))
	if err != nil {
		t.Fatal(err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("stamped document: %v", err)
	}
	bases, ok := doc.Frontmatter["bases"].(map[string]any)
	if !ok || bases["ALP-001"] != wantHash {
		t.Errorf("bases = %#v; want preceding hash %s", doc.Frontmatter["bases"], wantHash)
	}
	if doc.Frontmatter["custom"] != "author-value" || !strings.Contains(doc.Body, "Keep this human motivation.") || !strings.Contains(doc.Body, "Alpha MUST reflect the later revision.") || !strings.Contains(doc.Body, "Human note.") {
		t.Errorf("Stamp lost author content: %#v", doc)
	}
	queued(t, root, "0003-remove.md", "", "Remove the item.\n", "", "", "### ALP-001\n")
	if err := project.Stamp(root, "0003-remove.md"); err != nil {
		t.Fatalf("Stamp Removed: %v", err)
	}
	removedBytes, err := os.ReadFile(filepath.Join(root, "changes/0003-remove.md"))
	if err != nil {
		t.Fatal(err)
	}
	removedDoc, err := contracts.ParseDocument(string(removedBytes))
	if err != nil {
		t.Fatal(err)
	}
	removedWant := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST reflect the later revision.\n"})
	removedBases, ok := removedDoc.Frontmatter["bases"].(map[string]any)
	if !ok || removedBases["ALP-001"] != removedWant {
		t.Errorf("Removed base = %#v; want %s", removedDoc.Frontmatter["bases"], removedWant)
	}
}

func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-added.md", "", "Added item.\n", "### new:thing\nThing MUST work.\n", "", "")
	if err := project.Stamp(root, "added"); err != nil {
		t.Fatalf("Stamp Added-only: %v", err)
	}
	bytes, err := os.ReadFile(filepath.Join(root, "changes/0001-added.md"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(bytes), "### new:thing") || strings.Contains(string(bytes), "### ALP-002") {
		t.Errorf("Stamp assigned an ID to Added-only change: %s", bytes)
	}
	queued(t, root, "0002-invalid.md", "", "Invalid edit.\n", "", "### ALP-999\nMissing item MUST change.\n", "")
	if err := project.Stamp(root, "invalid"); err == nil {
		t.Error("Stamp accepted Modified ID absent from preceding projection")
	}
	removeRoot := fixture(t)
	queued(t, removeRoot, "0001-invalid-remove.md", "", "Invalid removal.\n", "", "", "### ALP-998\n")
	if err := project.Stamp(removeRoot, "invalid-remove"); err == nil {
		t.Error("Stamp accepted Removed ID absent from preceding projection")
	}
}

func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
	validRoot := fixture(t)
	put(t, validRoot, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	queued(t, validRoot, "0001-valid.md", "bases:\n  ALP-001: "+contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})+"\n", "A user needs the change.\n", "### new:local [review]\nA local addition MUST work.\n", "### ALP-001 [quarantine: flaky fixture]\nAlpha MUST work with revision.\n", "")
	queued(t, validRoot, "0002-also-valid.md", "", "Another user need.\n", "### new:local\nA second local addition MUST work.\n", "", "")
	for _, diagnostic := range project.Check(validRoot) {
		if !diagnostic.Warning {
			t.Errorf("valid spec/change received error diagnostic: %#v", diagnostic)
		}
	}
}

func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
	for _, tc := range []struct {
		id, name, spec, front, motivation, added, modified, removed string
		code, cause                                                 string
	}{
		{"PRJ-008", "missing_motivation", "", "", "\n", "", "", "", "C07", "missing Motivation"},
		{"PRJ-009", "nonnormative_module", "### ALP-001\nAlpha works.\n", "", "Needed.\n", "", "", "", "C04", "non-normative spec body"},
		{"PRJ-009", "nonnormative_change", "", "", "Needed.\n", "### new:item\nAn observation.\n", "", "", "C04", "non-normative spec body"},
		{"PRJ-010", "invalid_review_tag", "", "", "Needed.\n", "### new:item [review: reason]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-010", "invalid_quarantine_tag", "", "", "Needed.\n", "### new:item [quarantine]\nItem MUST work.\n", "", "", "C06", "invalid review/quarantine tag"},
		{"PRJ-011", "duplicate_module_id", "### ALP-001\nAlpha MUST work.\n\n### ALP-001\nDuplicate MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "duplicate item ID"},
		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
		{"PRJ-013", "duplicate_edit_identity", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "### ALP-001\n", "C05", "duplicate edit identity ALP-001"},
		{"PRJ-014", "undeclared_handle", "", "", "Needed.\n", "### new:local\nLocal MUST work.\n", "### ALP-001\nAlpha MUST refer to new:foreign.\n", "", "C09", "reference to undeclared handle new:foreign"},
		{"PRJ-015", "missing_base", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "missing base for ALP-001"},
		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
		{"PRJ-016", "absent_modified", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "### ALP-999\nAbsent MUST change.\n", "", "C08", "ID ALP-999 absent from preceding projection"},
		{"PRJ-016", "absent_removed", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "", "### ALP-999\n", "C08", "ID ALP-999 absent from preceding projection"},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			front := tc.front
			baseHash := contracts.SpecHash(contracts.Item{ID: "ALP-001", Body: "Alpha MUST work.\n"})
			if tc.name == "duplicate_edit_identity" || tc.name == "undeclared_handle" {
				front = "bases:\n  ALP-001: " + baseHash + "\n"
			}
			if tc.name == "foreign_change_id" {
				put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nBeta MUST work.\n")
				front = "bases:\n  BET-001: " + contracts.SpecHash(contracts.Item{ID: "BET-001", Body: "Beta MUST work.\n"}) + "\n"
			}
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n"+func() string {
				if tc.spec != "" {
					return tc.spec
				}
				return "### ALP-001\nAlpha MUST work.\n"
			}())
			if tc.name == "undeclared_handle" {
				queued(t, root, "0000-other.md", "", "Earlier local handle.\n", "### new:foreign\nEarlier item MUST work.\n", "", "")
			}
			queued(t, root, "0001-check.md", front, tc.motivation, tc.added, tc.modified, tc.removed)
			diagnostics := project.Check(root)
			found := false
			foundCause := false
			foundApplyCause := false
			for _, diagnostic := range diagnostics {
				if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
					t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
				}
				if !diagnostic.Warning && strings.TrimSpace(diagnostic.Code) != "" && strings.TrimSpace(diagnostic.Message) != "" {
					found = true
				}
				matchesCause := strings.Contains(diagnostic.Message, tc.cause)
				if tc.id == "PRJ-016" {
					matchesCause = diagnostic.Message == tc.cause
				}
				if tc.code != "" && !diagnostic.Warning && diagnostic.Code == tc.code && matchesCause {
					foundCause = true
				}
				if tc.name == "absent_modified" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Modified ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
				if tc.name == "absent_removed" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Removed ID ALP-999 absent from preceding projection") {
					foundApplyCause = true
				}
			}
			if !found {
				t.Errorf("Check did not report error diagnostic with nonempty code and message: %#v", diagnostics)
			}
			if tc.code != "" && !foundCause {
				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
			}
			if tc.id == "PRJ-016" && !foundApplyCause {
				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
			}
			if tc.name == "missing_base" || tc.name == "stale_base" {
				queued(t, root, "0001-check.md", front, "Needed.\n", "", "", "### ALP-001\n")
				removedDiagnostics := project.Check(root)
				found = false
				for _, diagnostic := range removedDiagnostics {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message for Removed entry: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, tc.cause) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not report C08 for Removed entry with %s (%s): %#v", tc.name, tc.cause, removedDiagnostics)
				}
			}
			if tc.name == "nonnormative_change" || tc.name == "invalid_review_tag" {
				front, added, modified, want := "", "", "", ""
				code := "C04"
				if tc.name == "nonnormative_change" {
					front = "bases:\n  ALP-001: " + baseHash + "\n"
					modified = "### ALP-001\nAn observation.\n"
					want = "non-normative spec body"
				} else {
					code = "C06"
					added = "### new:item [other: reason]\nItem MUST work.\n"
					want = "invalid review/quarantine tag"
				}
				queued(t, root, "0001-check.md", front, "Needed.\n", added, modified, "")
				found = false
				for _, diagnostic := range project.Check(root) {
					if !diagnostic.Warning && (strings.TrimSpace(diagnostic.Code) == "" || strings.TrimSpace(diagnostic.Message) == "") {
						t.Errorf("Check returned error diagnostic without code and message: %#v", diagnostic)
					}
					if !diagnostic.Warning && diagnostic.Code == code && strings.Contains(diagnostic.Message, want) {
						found = true
					}
				}
				if !found {
					t.Errorf("Check did not reject %s: wanted error diagnostic %s with message containing %q", tc.name, code, want)
				}
			}
		})
	}
}

 succeeded in 578ms:
package evals_test

import (
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_057_StampInvalidBasesHasCode(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-invalid-bases.md", "bases: [ALP-001]\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "")
	err := project.Stamp(root, "invalid-bases")
	if err == nil {
		t.Fatal("Stamp accepted bases that are not a mapping")
	}
	if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
		t.Errorf("Stamp error code = (%q, %v); want (C08, true): %v", code, ok, err)
	}
}

func Test_PRJ_058_MissingPendingChangeHasSubject(t *testing.T) {
	root := fixture(t)
	const target = "changes/0099-absent.md"
	for _, operation := range []struct {
		name string
		call func() error
	}{
		{"Stamp", func() error { return project.Stamp(root, target) }},
		{"Project", func() error { _, err := project.Project(root, target); return err }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			err := operation.call()
			if err == nil {
				t.Fatal("missing pending change was accepted")
			}
			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
				t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, target, err)
			}
		})
	}
}

func Test_PRJ_059_LoadPreservesModuleParseCode(t *testing.T) {
	for _, tc := range []struct {
		name, spec, code string
	}{
		{"malformed_frontmatter", "---\nname: [broken\n---\n# Alpha\n", "C03"},
		{"malformed_spec_heading", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001 [broken\nAlpha MUST work.\n", "C03"},
		{"duplicate_spec_id", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n\n### ALP-001\nAlpha MUST also work.\n", "C05"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, "packages/alpha/SPEC.md", tc.spec)
			_, err := project.Load(root)
			if err == nil {
				t.Fatal("Load accepted malformed module SPEC.md")
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != tc.code {
				t.Errorf("Load error code = (%q, %v); want (%s, true): %v", code, ok, tc.code, err)
			}
		})
	}
}
package evals_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func revisionGit(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	cmd.Env = append(os.Environ(), "GIT_CONFIG_GLOBAL=/dev/null", "GIT_CONFIG_NOSYSTEM=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %v: %v: %s", args, err, out)
	}
	return strings.TrimSpace(string(out))
}

func revisionCommit(t *testing.T, root, subject string) string {
	t.Helper()
	revisionGit(t, root, "add", "-A")
	revisionGit(t, root, "-c", "user.name=Fireweed Fixture", "-c", "user.email=fixture@example.invalid", "commit", "-qm", subject)
	return revisionGit(t, root, "rev-parse", "HEAD")
}

func historicalProject(t *testing.T) (string, string) {
	t.Helper()
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "fireweed.yaml", `version: 1
modules: ["packages/*"]
module_defaults:
  protected: [guard/**]
  evals: [evals/**]
  interface: [api.go]
  append_only: [ledger/**]
  preserve: [data/**]
  disposable: [internal/**]
`)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nHistorical alpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n\n### BET-001\nHistorical beta behavior.\n")
	put(t, root, "extras/gamma/SPEC.md", "---\nname: gamma\nid_prefix: GAM\ndepends_on: []\n---\n# Gamma\n\n## Specs\n\n### GAM-001\nHistorical gamma behavior.\n")
	recorded := revisionCommit(t, root, "record original project")

	put(t, root, "fireweed.yaml", "version: 2\nmodules: [\"components/*\"]\nmodule_defaults:\n  evals: [later/**]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: later-alpha\nid_prefix: LAL\n---\n# Later\n\n## Specs\n")
	put(t, root, "components/next/SPEC.md", "---\nname: next\nid_prefix: NXT\n---\n# Next\n\n## Specs\n")
	revisionCommit(t, root, "change config glob and modules")

	put(t, root, "fireweed.yaml", "version: 3\nmodules: [\"dirty/*\"]\n")
	put(t, root, "packages/alpha/SPEC.md", "---\nname: dirty-alpha\nid_prefix: DAL\n---\n# Dirty\n\n## Specs\n")
	put(t, root, "dirty/only/SPEC.md", "---\nname: dirty\nid_prefix: DRT\n---\n# Dirty\n\n## Specs\n")
	return root, recorded
}

func Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Config["version"] != 1 || !reflect.DeepEqual(got.Config["modules"], []any{"packages/*"}) {
		t.Errorf("recorded configuration = %#v", got.Config)
	}
	wantDefaults := map[string]any{
		"protected":   []any{"guard/**"},
		"evals":       []any{"evals/**"},
		"interface":   []any{"api.go"},
		"append_only": []any{"ledger/**"},
		"preserve":    []any{"data/**"},
		"disposable":  []any{"internal/**"},
	}
	if !reflect.DeepEqual(got.Config["module_defaults"], wantDefaults) {
		t.Errorf("recorded module_defaults = %#v; want %#v", got.Config["module_defaults"], wantDefaults)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("recorded modules = %#v; want alpha and beta only", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "Historical alpha behavior.") {
		t.Errorf("recorded alpha document = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) || !strings.Contains(beta.Document.Body, "Historical beta behavior.") {
		t.Errorf("recorded beta document = %#v", beta)
	}
}

func Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones(t *testing.T) {
	root, recorded := historicalProject(t)
	got, err := project.LoadAt(root, recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	byName := make(map[string]project.Module)
	for _, module := range got.Modules {
		byName[module.Name] = module
	}
	if len(got.Modules) != 2 {
		t.Fatalf("recorded modules = %#v", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if !reflect.DeepEqual(alpha.Zones.Protected, []string{"guard/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) ||
		!reflect.DeepEqual(alpha.Zones.AppendOnly, []string{"ledger/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) ||
		!reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) {
		t.Errorf("recorded inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone = %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("recorded default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs(t *testing.T) {
	root, recorded := historicalProject(t)
	canonical, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	got, err := project.LoadAt(root+string(filepath.Separator)+".", recorded)
	if err != nil {
		t.Fatalf("LoadAt(%s): %v", recorded, err)
	}
	if got.Root != canonical {
		t.Errorf("Root = %q; want canonical %q", got.Root, canonical)
	}
	dirs := make(map[string]string)
	for _, module := range got.Modules {
		dirs[module.Name] = module.Dir
	}
	if !reflect.DeepEqual(dirs, map[string]string{"alpha": "packages/alpha", "beta": "packages/beta"}) {
		t.Errorf("module directories = %#v", dirs)
	}
}

func Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs(t *testing.T) {
	root := t.TempDir()
	revisionGit(t, root, "init", "-q")
	put(t, root, "README.md", "repository without fireweed.yaml\n")
	missingConfig := revisionCommit(t, root, "missing config")
	put(t, root, "fireweed.yaml", "modules: [broken\n")
	malformedConfig := revisionCommit(t, root, "malformed config")
	put(t, root, "fireweed.yaml", "modules: [\"modules/*\"]\n")
	put(t, root, "modules/ghost/README.md", "directory without SPEC.md\n")
	missingSpec := revisionCommit(t, root, "missing module document")
	for _, tc := range []struct{ name, revision string }{
		{"invalid_revision", "not-a-revision"},
		{"missing_recorded_config", missingConfig},
		{"malformed_recorded_config", malformedConfig},
		{"missing_recorded_SPEC", missingSpec},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := project.LoadAt(root, tc.revision); err == nil {
				t.Errorf("LoadAt accepted %s (%s)", tc.name, tc.revision)
			}
		})
	}
}

 succeeded in 577ms:
package evals_test

import (
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

const fixtureConfig = `fireweed: ">=0.1, <0.2"
version: 1
base: main
modules: ["packages/*"]
module_defaults:
  interface: [api.go]
  evals: [evals/**]
  disposable: [internal/**]
  preserve: [data/**]
`

func put(t *testing.T, root, name, body string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(name))
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0644); err != nil {
		t.Fatal(err)
	}
}

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	put(t, root, "fireweed.yaml", fixtureConfig)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha behavior.\n")
	put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [alpha]\n---\n# Beta\n\n## Specs\n")
	return root
}

func Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults(t *testing.T) {
	root := fixture(t)
	put(t, root, "other/ignored/SPEC.md", "not a document")
	got, err := project.Load(root)
	if err != nil {
		t.Fatalf("Load: %v", err)
	}
	if got.Root != root {
		t.Errorf("Root = %q; want %q", got.Root, root)
	}
	if got.Config["version"] != 1 {
		t.Errorf("decoded Config version = %#v", got.Config["version"])
	}
	byName := map[string]project.Module{}
	for _, m := range got.Modules {
		byName[m.Name] = m
	}
	if len(got.Modules) != 2 || len(byName) != 2 {
		t.Fatalf("modules = %#v; want exactly alpha and beta", got.Modules)
	}
	alpha, beta := byName["alpha"], byName["beta"]
	if alpha.IDPrefix != "ALP" || alpha.Dir != "packages/alpha" || len(alpha.DependsOn) != 0 {
		t.Errorf("alpha identity = %#v", alpha)
	}
	if beta.IDPrefix != "BET" || beta.Dir != "packages/beta" || !reflect.DeepEqual(beta.DependsOn, []string{"alpha"}) {
		t.Errorf("beta identity/dependencies = %#v", beta)
	}
	if alpha.Document.Frontmatter["name"] != "alpha" || !strings.Contains(alpha.Document.Body, "### ALP-001\nAlpha behavior.") {
		t.Errorf("alpha Document = %#v", alpha.Document)
	}
	if !reflect.DeepEqual(alpha.Zones.Interface, []string{"api.go"}) || !reflect.DeepEqual(alpha.Zones.Disposable, []string{"internal/**"}) || !reflect.DeepEqual(alpha.Zones.Preserve, []string{"data/**"}) {
		t.Errorf("inherited zones = %#v", alpha.Zones)
	}
	if alpha.Zones.Evals == nil || len(alpha.Zones.Evals) != 0 {
		t.Errorf("explicit empty evals zone lost: %#v", alpha.Zones.Evals)
	}
	if !reflect.DeepEqual(beta.Zones.Evals, []string{"evals/**"}) {
		t.Errorf("default evals zone = %#v", beta.Zones.Evals)
	}
}

func Test_PRJ_018_LoadInvalidDocuments(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing_config", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_config", func(t *testing.T, root string) { put(t, root, "fireweed.yaml", "modules: [broken\n") }},
		{"missing_module_document", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module_document", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			if _, err := project.Load(root); err == nil {
				t.Error("Load accepted missing or malformed project document")
			}
		})
	}
}

func Test_PRJ_002_Graph(t *testing.T) {
	module := func(name, prefix string, deps ...string) project.Module {
		return project.Module{Name: name, IDPrefix: prefix, DependsOn: deps}
	}
	for _, tc := range []struct {
		id      string
		name    string
		modules []project.Module
		valid   bool
		explain []string
	}{
		{"PRJ-002", "acyclic", []project.Module{module("alpha", "ALP"), module("beta", "BET", "alpha"), module("gamma", "GAM", "beta")}, true, nil},
		{"PRJ-019", "empty_name", []project.Module{module("", "ALP")}, false, []string{"empty", "name"}},
		{"PRJ-020", "empty_prefix", []project.Module{module("alpha", "")}, false, []string{"empty", "prefix"}},
		{"PRJ-019", "duplicate_name", []project.Module{module("alpha", "ALP"), module("alpha", "BET")}, false, []string{"duplicate", "alpha"}},
		{"PRJ-020", "duplicate_prefix", []project.Module{module("alpha", "ALP"), module("beta", "ALP")}, false, []string{"duplicate", "alp"}},
		{"PRJ-021", "missing_dependency", []project.Module{module("alpha", "ALP", "missing")}, false, []string{"depend", "missing"}},
		{"PRJ-021", "self_dependency", []project.Module{module("alpha", "ALP", "alpha")}, false, []string{"self", "alpha"}},
		{"PRJ-022", "cycle", []project.Module{module("alpha", "ALP", "beta"), module("beta", "BET", "alpha")}, false, []string{"cycl"}},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			err := project.ValidateGraph(tc.modules)
			if tc.valid && err != nil {
				t.Errorf("valid graph rejected: %v", err)
			}
			if !tc.valid && (err == nil || strings.TrimSpace(err.Error()) == "") {
				t.Errorf("invalid graph needs an explanatory error; got %v", err)
			} else if !tc.valid {
				for _, part := range tc.explain {
					if !strings.Contains(strings.ToLower(err.Error()), part) {
						t.Errorf("error %q does not explain %s rejection (missing %q)", err, tc.name, part)
					}
				}
			}
		})
	}
}

func Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0003-waiting.md", "pending marker")
	put(t, root, "changes/archive/0007-finished.md", "archived marker")
	rel, err := project.NewChange(root, "alpha", "new-feature")
	if err != nil {
		t.Fatalf("NewChange: %v", err)
	}
	if rel != "changes/0008-new-feature.md" {
		t.Errorf("relative path = %q; want changes/0008-new-feature.md", rel)
	}
	data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(rel)))
	if err != nil {
		t.Fatalf("read new change: %v", err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		t.Fatalf("new change document: %v", err)
	}
	if doc.Frontmatter["module"] != "alpha" {
		t.Errorf("new change module = %#v", doc.Frontmatter["module"])
	}
	for _, field := range []string{"phase", "run", "blocked"} {
		if value, present := doc.Frontmatter[field]; present {
			t.Errorf("queued change has %s = %#v", field, value)
		}
	}
	for _, heading := range []string{"Motivation", "Interface", "Added", "Modified", "Removed", "Decisions"} {
		if _, ok := contracts.Section(doc.Body, heading); !ok {
			t.Errorf("missing %s section", heading)
		}
	}
	if regexp.MustCompile(`(?m)^### [A-Z]+-[0-9]+\b`).Match(data) {
		t.Errorf("skeleton assigned a spec ID: %s", data)
	}
}

func Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
	for _, archived := range []bool{false, true} {
		name := "pending"
		path := "changes/0002-repeat.md"
		if archived {
			name = "archived"
			path = "changes/archive/0002-repeat.md"
		}
		t.Run(name, func(t *testing.T) {
			root := fixture(t)
			put(t, root, path, "original marker")
			if _, err := project.NewChange(root, "alpha", "repeat"); err == nil {
				t.Error("duplicate slug accepted")
			}
			data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
			if err != nil || string(data) != "original marker" {
				t.Errorf("existing change overwritten: %q, %v", data, err)
			}
		})
	}
}

func Test_PRJ_025_PRJ_026_NewChangeInvalidArguments(t *testing.T) {
	for _, tc := range []struct{ id, name, module, slug string }{
		{"PRJ-025", "unknown_module", "missing", "valid"},
		{"PRJ-026", "uppercase_slug", "alpha", "Bad"},
		{"PRJ-026", "underscore_slug", "alpha", "bad_name"},
		{"PRJ-026", "slash_slug", "alpha", "../bad"},
		{"PRJ-026", "empty_slug", "alpha", ""},
	} {
		t.Run(tc.id+"_"+tc.name, func(t *testing.T) {
			root := fixture(t)
			if _, err := project.NewChange(root, tc.module, tc.slug); err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			}
		})
	}
}
package evals_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/project"
)

func Test_PRJ_040_CheckConfigurationCauses(t *testing.T) {
	for _, tc := range []struct {
		name string
		edit func(*testing.T, string)
	}{
		{"missing", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "fireweed.yaml")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed", func(t *testing.T, root string) {
			put(t, root, "fireweed.yaml", "modules: [broken\n")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			_, cause := project.Load(root)
			if cause == nil {
				t.Fatal("Load accepted invalid configuration")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C01 for configuration cause %q", cause)
		})
	}
}

func Test_PRJ_041_CheckGraphCauses(t *testing.T) {
	t.Run("current", func(t *testing.T) {
		root := fixture(t)
		put(t, root, "packages/beta/SPEC.md", "---\nname: beta\nid_prefix: BET\ndepends_on: [missing]\n---\n# Beta\n\n## Specs\n")
		snapshot, err := project.Load(root)
		if err != nil {
			t.Fatal(err)
		}
		cause := project.ValidateGraph(snapshot.Modules)
		if cause == nil {
			t.Fatal("ValidateGraph accepted missing dependency")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for current graph cause %q", cause)
	})
	t.Run("projected", func(t *testing.T) {
		root := fixture(t)
		queued(t, root, "0001-invalid.md", "set:\n  depends_on: [missing]\n", "Needed.\n", "", "", "")
		_, cause := project.Project(root, "")
		if cause == nil {
			t.Fatal("Project accepted invalid projected graph")
		}
		for _, diagnostic := range project.Check(root) {
			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
				return
			}
		}
		t.Errorf("Check did not report C02 for projected graph cause %q", cause)
	})
}

func Test_PRJ_042_CheckDocumentAndItemCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"missing_module", "load", func(t *testing.T, root string) {
			if err := os.Remove(filepath.Join(root, "packages/alpha/SPEC.md")); err != nil {
				t.Fatal(err)
			}
		}},
		{"malformed_module", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [broken\n---\n# Alpha\n")
		}},
		{"malformed_change", "queue", func(t *testing.T, root string) {
			put(t, root, "changes/0001-bad.md", "---\nmodule: [broken\n---\n")
		}},
		{"malformed_item", "queue", func(t *testing.T, root string) {
			queued(t, root, "0001-bad-item.md", "", "Needed.\n", "### new:item [broken\nItem MUST work.\n", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Queue(root)
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid document or item")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C03" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C03 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_043_CheckMetadataAndSetCauses(t *testing.T) {
	for _, tc := range []struct {
		name, source string
		edit         func(*testing.T, string)
	}{
		{"metadata_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: [alpha]\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n")
		}},
		{"zone_type", "load", func(t *testing.T, root string) {
			put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\nevals: 7\n---\n# Alpha\n\n## Specs\n")
		}},
		{"identity_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  name: renamed\n", "Needed.\n", "", "", "")
		}},
		{"unknown_set_key", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  surprise: true\n", "Needed.\n", "", "", "")
		}},
		{"set_value_type", "project", func(t *testing.T, root string) {
			queued(t, root, "0001-invalid.md", "set:\n  interface: 7\n", "Needed.\n", "", "", "")
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.edit(t, root)
			var cause error
			if tc.source == "load" {
				_, cause = project.Load(root)
			} else {
				_, cause = project.Project(root, "")
			}
			if cause == nil {
				t.Fatal("public operation accepted invalid metadata or set")
			}
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == "C12" && strings.Contains(diagnostic.Message, cause.Error()) {
					return
				}
			}
			t.Errorf("Check did not report C12 for %s cause %q", tc.source, cause)
		})
	}
}

func Test_PRJ_044_CheckQueueIdentityCause(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-repeat.md", "", "Needed.\n", "", "", "")
	queued(t, root, "0002-repeat.md", "", "Needed.\n", "", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate slug")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C13" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C13 for queue identity cause %q", cause)
}

func Test_PRJ_045_CheckPendingModuleTargetCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "changes/0001-bad.md", "---\nmodule: missing\n---\n## Motivation\nNeeded.\n")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted unknown change module")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C07" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C07 for pending module cause %q", cause)
}

func Test_PRJ_013_CheckDuplicateWithinEditSection(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-duplicate.md", "", "Needed.\n", "### new:local\nLocal MUST work.\n\n### new:local\nDuplicate MUST fail.\n", "", "")
	_, cause := project.Queue(root)
	if cause == nil {
		t.Fatal("Queue accepted duplicate Added identity")
	}
	for _, diagnostic := range project.Check(root) {
		if !diagnostic.Warning && diagnostic.Code == "C05" && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check did not report C05 for duplicate Added identity cause %q", cause)
}

codex
I found a small set of causes behind the reported failures: working-tree `Load` resolves a symlink the test expects it to retain, duplicate item parse errors lose their C05 code, foreign change IDs use the wrong diagnostic wording, and `NewChange` returns uncoded argument errors. I’m fixing those in the disposable implementation.
exec
/bin/zsh -lc "rg -n 'NewChange|C13|ambiguous|C05|C03|U02|C08|cause|missing module' packages/project/evals/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat SYSTEM.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'cat packages/project/internal/history/history.go && cat packages/project/internal/queue/stale.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 566ms:
# Independent project reconstruction policy

This policy governs the independent, input-only project reconstruction workspace.
Rebuild project solely from its retained SPEC.md, Decisions, public interfaces,
model types and evals, plus supplied contracts, gitstore and assets interfaces.
The Go module path is github.com/lurkshark/fireweed, as shown by those interfaces.

## Retained repository conventions

The product never contacts the network or runs fetch or push. Read operations
never run hooks or agents and never change files, the index or refs; output is
deterministic. These are repository conventions, not new module spec items.
Retained api.go facades and model types remain interfaces; behavior belongs
under internal/. Never import the public project parent from its own internal
implementation; shared data belongs in model or contracts. Keep dependencies
minimal and add no native objects, assembly, cgo or embedded executables.

## Input isolation and verification

Read only the supplied input workspace and your newly written files. Do not read
outside it, old implementation, source history, saved attempts, other-module
internals, product binaries, bootstrap, external checkouts or prior drill output.
Do not query the invoking executable for project behavior, help, recipes,
playbooks, explanations or other content. Do not run fireweed authoring or
lifecycle commands in this input workspace. Do not commit or invoke other agents.

Retained inputs must remain byte-unchanged. Write reconstruction only under
packages/project/internal/**, and write the requested DRILL_NOTES.md with
Contradicted Decisions and Gaps sections. Do not create root manifests, copy
working dependency implementations, or repair dependent code. Remove temporary
probe files before finishing. Every behavioral gap names a nearest spec ID (or
no spec) and its target module; do not author fixes from gaps.

This input deliberately lacks full build manifests, dependency implementations
and verification hooks. Supplied dependency facades can reference absent internal
packages; do not search for those implementations. The separate native verifier
supplies the full project, actual dependencies and dependents and runs project's
own evals plus every transitive dependent's enforced evals and coverage. Inability
to run that full suite here is an input-only verification limitation, not a reason
to copy forbidden material or change retained inputs. Describe it accurately if
relevant; the native verifier will provide actual results.

Use ordinary local Go and Git tools for bounded owned-fixture probes when useful.
Keep the current input tree's .fireweed/bin first on worker PATH and use
GOCACHE=/tmp/fireweed-rebuild-cache. No controller binary is supplied as an input.
The orchestration config selects exactly codex exec --ephemeral
--ignore-user-config -m gpt-6-sol -c model_reasoning_effort="high"
-s workspace-write {prompt}; do not substitute model or effort.

 succeeded in 571ms:
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
	s, e := fixture.Load(root)
	if e != nil {
		return nil, e
	}
	var target *model.Module
	for i := range s.Modules {
		if s.Modules[i].Name == module {
			target = &s.Modules[i]
			break
		}
	}
	if target == nil {
		return nil, fixture.Coded("U02", module, fmt.Errorf("U02 unknown module %s", module))
	}
	paths, e := os.ReadDir(filepath.Join(s.Root, "changes/archive"))
	if os.IsNotExist(e) {
		paths = nil
	} else if e != nil {
		return nil, fixture.Coded("C03", "changes/archive", e)
	}
	out := []model.SpecHistoryEntry{}
	found := false
	for _, ent := range paths {
		if ent.IsDir() || !strings.HasSuffix(ent.Name(), ".md") {
			continue
		}
		path := "changes/archive/" + ent.Name()
		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
		if e != nil {
			return nil, fixture.Coded("C03", path, e)
		}
		doc, e := contracts.ParseDocument(string(b))
		if e != nil {
			return nil, fixture.Coded("C03", path, e)
		}
		mod, ok := doc.Frontmatter["module"].(string)
		if !ok {
			return nil, fixture.Coded("C03", path, fmt.Errorf("invalid archive module"))
		}
		addedText, _ := contracts.Section(doc.Body, "Added")
		modifiedText, _ := contracts.Section(doc.Body, "Modified")
		added, e := contracts.ParseItems(addedText)
		if e != nil {
			return nil, fixture.Coded("C03", path, e)
		}
		modified, e := contracts.ParseItems(modifiedText)
		if e != nil {
			return nil, fixture.Coded("C03", path, e)
		}
		if mod != module {
			continue
		}
		ids := map[string]bool{}
		for _, it := range append(added, modified...) {
			ids[it.ID] = true
		}
		if id != "" {
			if !ids[id] {
				continue
			}
			found = true
			ids = map[string]bool{id: true}
		}
		if len(ids) == 0 {
			continue
		}
		list := make([]string, 0, len(ids))
		for k := range ids {
			list = append(list, k)
		}
		sort.Strings(list)
		mot, _ := contracts.Section(doc.Body, "Motivation")
		dec, _ := contracts.Section(doc.Body, "Decisions")
		removed, _ := contracts.Section(doc.Body, "Removed decisions")
		prov, _ := contracts.Section(doc.Body, "Provenance")
		out = append(out, model.SpecHistoryEntry{Path: path, Module: module, IDs: list, Motivation: mot, Decisions: dec, RemovedDecisions: removed, Provenance: prov, ImplementationDate: implDate(prov)})
	}
	if id != "" && !found {
		body, _ := contracts.Section(target.Document.Body, "Specs")
		items, e := contracts.ParseItems(body)
		if e != nil {
			return nil, fixture.Coded("C03", target.Dir+"/SPEC.md", e)
		}
		current := false
		for _, it := range items {
			if it.ID == id {
				current = true
			}
		}
		if !current {
			return nil, fixture.Coded("U02", id, fmt.Errorf("U02 unknown spec ID %s", id))
		}
	}
	sort.Slice(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if !a.ImplementationDate.Equal(b.ImplementationDate) {
			if a.ImplementationDate.IsZero() {
				return false
			}
			if b.ImplementationDate.IsZero() {
				return true
			}
			return a.ImplementationDate.After(b.ImplementationDate)
		}
		return a.Path < b.Path
	})
	return out, nil
}
func implDate(prov string) time.Time {
	var latest time.Time
	for _, line := range strings.Split(prov, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
			continue
		}
		t, e := time.Parse(time.RFC3339Nano, fields[2])
		if e == nil && t.After(latest) {
			latest = t
		}
	}
	return latest
}
package queue

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func StaleSpecs(root, target string) ([]model.StaleSpec, error) {
	s, c, e := preceding(root, target)
	if e != nil {
		return nil, e
	}
	prior, e := Before(s, c)
	if e != nil {
		return nil, e
	}
	b, e := bases(c)
	if e != nil {
		return nil, e
	}
	out := []model.StaleSpec{}
	historical := map[string]contracts.Item{}
	historyLoaded := false
	for _, it := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
		raw, ok := b[it.ID]
		if !ok {
			continue
		}
		recorded, ok := raw.(string)
		if !ok {
			recorded = fmt.Sprint(raw)
		}
		current, ok := prior[it.ID]
		if !ok {
			continue
		}
		if recorded == contracts.SpecHash(current) {
			continue
		}
		entry := model.StaleSpec{ID: it.ID, RecordedHash: recorded, Current: contracts.RenderItem(current)}
		if !historyLoaded {
			historical = recoverBefore(s.Root, c)
			historyLoaded = true
		}
		if old, ok := historical[it.ID]; ok && contracts.SpecHash(old) == recorded {
			entry.Before = contracts.RenderItem(old)
			entry.Recovered = true
		}
		out = append(out, entry)
	}
	return out, nil
}
func recoverBefore(root string, c model.Change) map[string]contracts.Item {
	out := map[string]contracts.Item{}
	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "log", "--format=%H", "--", c.Path}})
	if e != nil || r.ExitCode != 0 {
		return out
	}
	commits := strings.Fields(r.Stdout)
	var chosen string
	var priorBases map[string]any
	for _, rev := range commits {
		b, e := gitstore.ReadAt(root, rev, c.Path)
		if e != nil {
			continue
		}
		doc, e := contracts.ParseDocument(string(b))
		if e != nil {
			continue
		}
		value, ok := doc.Frontmatter["bases"].(map[string]any)
		if !ok {
			value = map[string]any{}
		}
		if chosen == "" {
			priorBases = value
			chosen = rev
			continue
		}
		if !equalBases(priorBases, value) {
			break
		}
		chosen = rev
	}
	if chosen == "" {
		return out
	}
	hist, e := fixture.LoadAt(root, chosen)
	if e != nil {
		return out
	}
	files, e := fixture.TreePaths(root, chosen)
	if e != nil {
		return out
	}
	changes := []model.Change{}
	for _, path := range files {
		if !strings.HasPrefix(path, "changes/") || strings.HasPrefix(path, "changes/archive/") {
			continue
		}
		if _, _, ok := fixture.ChangeName(strings.TrimPrefix(path, "changes/")); !ok {
			continue
		}
		b, e := gitstore.ReadAt(root, chosen, path)
		if e != nil {
			return out
		}
		item, e := parseChange(path, b)
		if e != nil {
			return out
		}
		changes = append(changes, item)
	}
	sortChanges(changes)
	index := -1
	for i, item := range changes {
		if item.Path == c.Path {
			index = i
			break
		}
	}
	if index < 0 {
		return out
	}
	for _, item := range changes[:index] {
		if e := Apply(&hist, item); e != nil {
			return out
		}
	}
	m := findModule(&hist, c.Module)
	if m == nil {
		return out
	}
	items, e := SpecItems(*m)
	if e != nil {
		return out
	}
	for _, it := range items {
		out[it.ID] = it
	}
	return out
}
func equalBases(a, b map[string]any) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if !reflect.DeepEqual(b[k], v) {
			return false
		}
	}
	return true
}
func sortChanges(cs []model.Change) {
	for i := 1; i < len(cs); i++ {
		for j := i; j > 0 && (cs[j].Number < cs[j-1].Number || (cs[j].Number == cs[j-1].Number && cs[j].Slug < cs[j-1].Slug)); j-- {
			cs[j], cs[j-1] = cs[j-1], cs[j]
		}
	}
}

 succeeded in 579ms:
packages/project/evals/revision_snapshot_test.go:158:	missingSpec := revisionCommit(t, root, "missing module document")
packages/project/evals/coded_errors_test.go:17:	if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
packages/project/evals/coded_errors_test.go:18:		t.Errorf("Stamp error code = (%q, %v); want (C08, true): %v", code, ok, err)
packages/project/evals/coded_errors_test.go:37:			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
packages/project/evals/coded_errors_test.go:38:				t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, target, err)
packages/project/evals/coded_errors_test.go:48:		{"malformed_frontmatter", "---\nname: [broken\n---\n# Alpha\n", "C03"},
packages/project/evals/coded_errors_test.go:49:		{"malformed_spec_heading", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001 [broken\nAlpha MUST work.\n", "C03"},
packages/project/evals/coded_errors_test.go:50:		{"duplicate_spec_id", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n\n### ALP-001\nAlpha MUST also work.\n", "C05"},
packages/project/evals/project_test.go:155:func Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering(t *testing.T) {
packages/project/evals/project_test.go:159:	rel, err := project.NewChange(root, "alpha", "new-feature")
packages/project/evals/project_test.go:161:		t.Fatalf("NewChange: %v", err)
packages/project/evals/project_test.go:192:func Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
packages/project/evals/project_test.go:203:			if _, err := project.NewChange(root, "alpha", "repeat"); err == nil {
packages/project/evals/project_test.go:214:func Test_PRJ_025_PRJ_026_NewChangeInvalidArguments(t *testing.T) {
packages/project/evals/project_test.go:224:			if _, err := project.NewChange(root, tc.module, tc.slug); err == nil {
packages/project/evals/static_codes_test.go:29:			_, cause := project.Load(root)
packages/project/evals/static_codes_test.go:30:			if cause == nil {
packages/project/evals/static_codes_test.go:34:				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:38:			t.Errorf("Check did not report C01 for configuration cause %q", cause)
packages/project/evals/static_codes_test.go:51:		cause := project.ValidateGraph(snapshot.Modules)
packages/project/evals/static_codes_test.go:52:		if cause == nil {
packages/project/evals/static_codes_test.go:56:			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:60:		t.Errorf("Check did not report C02 for current graph cause %q", cause)
packages/project/evals/static_codes_test.go:65:		_, cause := project.Project(root, "")
packages/project/evals/static_codes_test.go:66:		if cause == nil {
packages/project/evals/static_codes_test.go:70:			if !diagnostic.Warning && diagnostic.Code == "C02" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:74:		t.Errorf("Check did not report C02 for projected graph cause %q", cause)
packages/project/evals/static_codes_test.go:101:			var cause error
packages/project/evals/static_codes_test.go:103:				_, cause = project.Load(root)
packages/project/evals/static_codes_test.go:105:				_, cause = project.Queue(root)
packages/project/evals/static_codes_test.go:107:			if cause == nil {
packages/project/evals/static_codes_test.go:111:				if !diagnostic.Warning && diagnostic.Code == "C03" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:115:			t.Errorf("Check did not report C03 for %s cause %q", tc.source, cause)
packages/project/evals/static_codes_test.go:144:			var cause error
packages/project/evals/static_codes_test.go:146:				_, cause = project.Load(root)
packages/project/evals/static_codes_test.go:148:				_, cause = project.Project(root, "")
packages/project/evals/static_codes_test.go:150:			if cause == nil {
packages/project/evals/static_codes_test.go:154:				if !diagnostic.Warning && diagnostic.Code == "C12" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:158:			t.Errorf("Check did not report C12 for %s cause %q", tc.source, cause)
packages/project/evals/static_codes_test.go:167:	_, cause := project.Queue(root)
packages/project/evals/static_codes_test.go:168:	if cause == nil {
packages/project/evals/static_codes_test.go:172:		if !diagnostic.Warning && diagnostic.Code == "C13" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:176:	t.Errorf("Check did not report C13 for queue identity cause %q", cause)
packages/project/evals/static_codes_test.go:182:	_, cause := project.Queue(root)
packages/project/evals/static_codes_test.go:183:	if cause == nil {
packages/project/evals/static_codes_test.go:187:		if !diagnostic.Warning && diagnostic.Code == "C07" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:191:	t.Errorf("Check did not report C07 for pending module cause %q", cause)
packages/project/evals/static_codes_test.go:197:	_, cause := project.Queue(root)
packages/project/evals/static_codes_test.go:198:	if cause == nil {
packages/project/evals/static_codes_test.go:202:		if !diagnostic.Warning && diagnostic.Code == "C05" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:206:	t.Errorf("Check did not report C05 for duplicate Added identity cause %q", cause)
packages/project/evals/stale_test.go:205:		if diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "stale base for ALP-001") {
packages/project/evals/stale_test.go:233:			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C08" {
packages/project/evals/stale_test.go:234:				t.Errorf("StampMissing error = %v; want C08", err)
packages/project/evals/stale_test.go:272:			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
packages/project/evals/stale_test.go:273:				t.Errorf("unknown selector = %v; want U02 with subject %q", err, target)
packages/project/evals/stale_test.go:286:		t.Run(operation.name+"_ambiguous", func(t *testing.T) {
packages/project/evals/stale_test.go:288:			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C13" {
packages/project/evals/stale_test.go:289:				t.Errorf("ambiguous slug = %v; want C13", err)
packages/project/evals/spec_history_test.go:189:		if err == nil || !strings.HasPrefix(err.Error(), "U02") || !strings.Contains(err.Error(), tc.rejected) {
packages/project/evals/spec_history_test.go:190:			t.Errorf("lookup module %q ID %q = %v; want U02 naming %q", tc.module, tc.id, err, tc.rejected)
packages/project/evals/spec_history_test.go:225:			if err == nil || !strings.HasPrefix(err.Error(), "C03") || !strings.Contains(err.Error(), path) {
packages/project/evals/spec_history_test.go:226:				t.Errorf("malformed archive = %v; want C03 naming %s", err, path)
packages/project/evals/queue_test.go:134:		t.Error("Queue accepted ambiguous slug")
packages/project/evals/queue_test.go:137:		t.Error("Project accepted ambiguous slug")
packages/project/evals/queue_test.go:140:		t.Error("Stamp accepted ambiguous slug")
packages/project/evals/queue_test.go:295:		code, cause                                                 string
packages/project/evals/queue_test.go:302:		{"PRJ-011", "duplicate_module_id", "### ALP-001\nAlpha MUST work.\n\n### ALP-001\nDuplicate MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "duplicate item ID"},
packages/project/evals/queue_test.go:303:		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
packages/project/evals/queue_test.go:304:		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
packages/project/evals/queue_test.go:305:		{"PRJ-013", "duplicate_edit_identity", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "### ALP-001\n", "C05", "duplicate edit identity ALP-001"},
packages/project/evals/queue_test.go:307:		{"PRJ-015", "missing_base", "", "", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "missing base for ALP-001"},
packages/project/evals/queue_test.go:308:		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
packages/project/evals/queue_test.go:309:		{"PRJ-016", "absent_modified", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "### ALP-999\nAbsent MUST change.\n", "", "C08", "ID ALP-999 absent from preceding projection"},
packages/project/evals/queue_test.go:310:		{"PRJ-016", "absent_removed", "", "bases:\n  ALP-999: deadbeef\n", "Needed.\n", "", "", "### ALP-999\n", "C08", "ID ALP-999 absent from preceding projection"},
packages/project/evals/queue_test.go:344:				matchesCause := strings.Contains(diagnostic.Message, tc.cause)
packages/project/evals/queue_test.go:346:					matchesCause = diagnostic.Message == tc.cause
packages/project/evals/queue_test.go:351:				if tc.name == "absent_modified" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Modified ID ALP-999 absent from preceding projection") {
packages/project/evals/queue_test.go:354:				if tc.name == "absent_removed" && !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "Removed ID ALP-999 absent from preceding projection") {
packages/project/evals/queue_test.go:362:				t.Errorf("Check did not report %s for %s: %#v", tc.code, tc.cause, diagnostics)
packages/project/evals/queue_test.go:365:				t.Errorf("Check did not report C08 for absent %s during projection application: %#v", tc.name, diagnostics)
packages/project/evals/queue_test.go:375:					if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, tc.cause) {
packages/project/evals/queue_test.go:380:					t.Errorf("Check did not report C08 for Removed entry with %s (%s): %#v", tc.name, tc.cause, removedDiagnostics)

2026-11-02T00:19:50.839643Z ERROR codex_core::tools::router: error=apply_patch verification failed: Failed to find expected lines in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/fixture/fixture.go:
			return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..dba2740d49c1c07a0fa2b0bc5459aff5d07e7cc2
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -31,6 +31,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +164,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +198,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +247,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..4df1784383a64ca27233d4c33f6caa52ee0d2b1e
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -31,6 +31,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +164,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +198,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +247,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {
@@ -321,7 +328,7 @@
 }
 func NewChange(root, module, slug string) (string, error) {
 	if !slugRE.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
 	}
 	s, e := Load(root)
 	if e != nil {
@@ -334,7 +341,7 @@
 		}
 	}
 	if !found {
-		return "", fmt.Errorf("unknown module %s", module)
+		return "", Coded("U02", module, fmt.Errorf("module %s not found", module))
 	}
 	max := 0
 	for _, dir := range []string{"changes", "changes/archive"} {
@@ -343,12 +350,15 @@
 			return "", e
 		}
 		for _, ent := range entries {
+			if ent.IsDir() {
+				continue
+			}
 			n, sl, ok := ChangeName(ent.Name())
 			if !ok {
 				continue
 			}
 			if sl == slug {
-				return "", fmt.Errorf("duplicate change slug %s", slug)
+				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
 			}
 			if n > max {
 				max = n
@@ -362,7 +372,7 @@
 	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
 	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
 	if e != nil {
-		return "", e
+		return "", Coded("P01", rel, fmt.Errorf("create change file %s: %w", rel, e))
 	}
 	defer f.Close()
 	_, e = f.WriteString(body)

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..4df1784383a64ca27233d4c33f6caa52ee0d2b1e
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -31,6 +31,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +164,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +198,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +247,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {
@@ -321,7 +328,7 @@
 }
 func NewChange(root, module, slug string) (string, error) {
 	if !slugRE.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
 	}
 	s, e := Load(root)
 	if e != nil {
@@ -334,7 +341,7 @@
 		}
 	}
 	if !found {
-		return "", fmt.Errorf("unknown module %s", module)
+		return "", Coded("U02", module, fmt.Errorf("module %s not found", module))
 	}
 	max := 0
 	for _, dir := range []string{"changes", "changes/archive"} {
@@ -343,12 +350,15 @@
 			return "", e
 		}
 		for _, ent := range entries {
+			if ent.IsDir() {
+				continue
+			}
 			n, sl, ok := ChangeName(ent.Name())
 			if !ok {
 				continue
 			}
 			if sl == slug {
-				return "", fmt.Errorf("duplicate change slug %s", slug)
+				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
 			}
 			if n > max {
 				max = n
@@ -362,7 +372,7 @@
 	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
 	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
 	if e != nil {
-		return "", e
+		return "", Coded("P01", rel, fmt.Errorf("create change file %s: %w", rel, e))
 	}
 	defer f.Close()
 	_, e = f.WriteString(body)
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 51709337a73d6eeef1736960282553ef1d6cfa6d..d94a7b34a50289d1d9eb401465bf9d42feca3efa
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -32,7 +32,7 @@
 		body, _ := contracts.Section(doc.Body, part.name)
 		items, e := contracts.ParseItems(body)
 		if e != nil {
-			return c, fixture.Coded("C03", path, e)
+			return c, fixture.Coded(fixture.ItemErrorCode(e), path, e)
 		}
 		*part.dest = items
 	}
@@ -75,7 +75,7 @@
 		}
 		path := "changes/" + ent.Name()
 		if slugs[slug] && !allowDuplicates {
-			return nil, fixture.Coded("C13", path, fmt.Errorf("duplicate pending slug %s", slug))
+			return nil, fixture.Coded("C13", path, fmt.Errorf("ambiguous change slug %s: duplicate pending slug", slug))
 		}
 		slugs[slug] = true
 		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
@@ -114,10 +114,10 @@
 		}
 	}
 	if len(matches) == 0 {
-		return -1, fixture.Coded("U02", target, fmt.Errorf("unknown change %s", target))
+		return -1, fixture.Coded("U02", target, fmt.Errorf("change %s not found", target))
 	}
 	if len(matches) > 1 {
-		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change %s", target))
+		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change slug %s", target))
 	}
 	return matches[0], nil
 }
@@ -370,11 +370,15 @@
 	for _, m := range s.Modules {
 		prefixes[m.IDPrefix] = m.Name
 	}
-	checkItem := func(it contracts.Item, path, module string, needBody bool) {
+	checkItem := func(it contracts.Item, path, module string, needBody, change bool) {
 		if strings.Contains(it.ID, "-") {
 			prefix := strings.SplitN(it.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module {
-				add("C05", fmt.Sprintf("foreign spec ID %s belongs to %s", it.ID, owner), path)
+				kind := "spec"
+				if change {
+					kind = "module"
+				}
+				add("C05", fmt.Sprintf("foreign %s ID %s belongs to %s", kind, it.ID, owner), path)
 			}
 		}
 		if needBody && !normRE.MatchString(it.Body) {
@@ -400,7 +404,7 @@
 				add("C05", "duplicate item ID "+it.ID, path)
 			}
 			seen[it.ID] = true
-			checkItem(it, path, m.Name, true)
+			checkItem(it, path, m.Name, true, false)
 		}
 	}
 	changes, e := queue(root, s, false)
@@ -436,10 +440,10 @@
 			}
 		}
 		for _, it := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
-			checkItem(it, c.Path, c.Module, true)
+			checkItem(it, c.Path, c.Module, true, true)
 		}
 		for _, it := range c.Removed {
-			checkItem(it, c.Path, c.Module, false)
+			checkItem(it, c.Path, c.Module, false, true)
 		}
 		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
 			if !added[ref] {

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..112779bafa1126b95921e9a7f711874c2db65c96
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -16,6 +16,9 @@
 )
 
 func Coded(code, subject string, err error) error {
+	if code == "C05" {
+		err = fmt.Errorf("%w; remove or rename the conflicting spec ID", err)
+	}
 	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s %s: %w", code, subject, err)}
 }
 func Code(err error) string {
@@ -31,6 +34,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +167,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +201,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +250,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {
@@ -321,7 +331,7 @@
 }
 func NewChange(root, module, slug string) (string, error) {
 	if !slugRE.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
 	}
 	s, e := Load(root)
 	if e != nil {
@@ -334,7 +344,7 @@
 		}
 	}
 	if !found {
-		return "", fmt.Errorf("unknown module %s", module)
+		return "", Coded("U02", module, fmt.Errorf("module %s not found", module))
 	}
 	max := 0
 	for _, dir := range []string{"changes", "changes/archive"} {
@@ -343,12 +353,15 @@
 			return "", e
 		}
 		for _, ent := range entries {
+			if ent.IsDir() {
+				continue
+			}
 			n, sl, ok := ChangeName(ent.Name())
 			if !ok {
 				continue
 			}
 			if sl == slug {
-				return "", fmt.Errorf("duplicate change slug %s", slug)
+				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
 			}
 			if n > max {
 				max = n
@@ -362,7 +375,7 @@
 	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
 	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
 	if e != nil {
-		return "", e
+		return "", Coded("P01", rel, fmt.Errorf("create change file %s: %w", rel, e))
 	}
 	defer f.Close()
 	_, e = f.WriteString(body)
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 51709337a73d6eeef1736960282553ef1d6cfa6d..d94a7b34a50289d1d9eb401465bf9d42feca3efa
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -32,7 +32,7 @@
 		body, _ := contracts.Section(doc.Body, part.name)
 		items, e := contracts.ParseItems(body)
 		if e != nil {
-			return c, fixture.Coded("C03", path, e)
+			return c, fixture.Coded(fixture.ItemErrorCode(e), path, e)
 		}
 		*part.dest = items
 	}
@@ -75,7 +75,7 @@
 		}
 		path := "changes/" + ent.Name()
 		if slugs[slug] && !allowDuplicates {
-			return nil, fixture.Coded("C13", path, fmt.Errorf("duplicate pending slug %s", slug))
+			return nil, fixture.Coded("C13", path, fmt.Errorf("ambiguous change slug %s: duplicate pending slug", slug))
 		}
 		slugs[slug] = true
 		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
@@ -114,10 +114,10 @@
 		}
 	}
 	if len(matches) == 0 {
-		return -1, fixture.Coded("U02", target, fmt.Errorf("unknown change %s", target))
+		return -1, fixture.Coded("U02", target, fmt.Errorf("change %s not found", target))
 	}
 	if len(matches) > 1 {
-		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change %s", target))
+		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change slug %s", target))
 	}
 	return matches[0], nil
 }
@@ -370,11 +370,15 @@
 	for _, m := range s.Modules {
 		prefixes[m.IDPrefix] = m.Name
 	}
-	checkItem := func(it contracts.Item, path, module string, needBody bool) {
+	checkItem := func(it contracts.Item, path, module string, needBody, change bool) {
 		if strings.Contains(it.ID, "-") {
 			prefix := strings.SplitN(it.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module {
-				add("C05", fmt.Sprintf("foreign spec ID %s belongs to %s", it.ID, owner), path)
+				kind := "spec"
+				if change {
+					kind = "module"
+				}
+				add("C05", fmt.Sprintf("foreign %s ID %s belongs to %s", kind, it.ID, owner), path)
 			}
 		}
 		if needBody && !normRE.MatchString(it.Body) {
@@ -400,7 +404,7 @@
 				add("C05", "duplicate item ID "+it.ID, path)
 			}
 			seen[it.ID] = true
-			checkItem(it, path, m.Name, true)
+			checkItem(it, path, m.Name, true, false)
 		}
 	}
 	changes, e := queue(root, s, false)
@@ -436,10 +440,10 @@
 			}
 		}
 		for _, it := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
-			checkItem(it, c.Path, c.Module, true)
+			checkItem(it, c.Path, c.Module, true, true)
 		}
 		for _, it := range c.Removed {
-			checkItem(it, c.Path, c.Module, false)
+			checkItem(it, c.Path, c.Module, false, true)
 		}
 		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
 			if !added[ref] {

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..125da22d5c2ce63ed5644a77d06c154705c948ee
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -2,6 +2,7 @@
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"os"
 	"path/filepath"
@@ -16,6 +17,9 @@
 )
 
 func Coded(code, subject string, err error) error {
+	if code == "C05" {
+		err = fmt.Errorf("%w; remove or rename the conflicting spec ID", err)
+	}
 	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s %s: %w", code, subject, err)}
 }
 func Code(err error) string {
@@ -31,6 +35,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +168,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +202,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +251,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {
@@ -295,7 +306,18 @@
 			}
 		}
 	}
-	return loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
+	snapshot, e := loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
+	if e != nil {
+		if code, subject, ok := contracts.ErrorCode(e); ok && code == "C03" && strings.HasSuffix(subject, "/SPEC.md") {
+			cause := "malformed recorded SPEC"
+			if errors.Is(e, os.ErrNotExist) {
+				cause = "missing recorded SPEC"
+			}
+			return snapshot, Coded("C03", subject, fmt.Errorf("%s %s: %w", cause, subject, e))
+		}
+		return snapshot, e
+	}
+	return snapshot, nil
 }
 func TreePaths(root, revision string) ([]string, error) {
 	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "ls-tree", "-r", "--name-only", revision}})
@@ -321,7 +343,7 @@
 }
 func NewChange(root, module, slug string) (string, error) {
 	if !slugRE.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
 	}
 	s, e := Load(root)
 	if e != nil {
@@ -334,7 +356,7 @@
 		}
 	}
 	if !found {
-		return "", fmt.Errorf("unknown module %s", module)
+		return "", Coded("U02", module, fmt.Errorf("module %s not found", module))
 	}
 	max := 0
 	for _, dir := range []string{"changes", "changes/archive"} {
@@ -343,12 +365,15 @@
 			return "", e
 		}
 		for _, ent := range entries {
+			if ent.IsDir() {
+				continue
+			}
 			n, sl, ok := ChangeName(ent.Name())
 			if !ok {
 				continue
 			}
 			if sl == slug {
-				return "", fmt.Errorf("duplicate change slug %s", slug)
+				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
 			}
 			if n > max {
 				max = n
@@ -362,7 +387,7 @@
 	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
 	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
 	if e != nil {
-		return "", e
+		return "", Coded("P01", rel, fmt.Errorf("create change file %s: %w", rel, e))
 	}
 	defer f.Close()
 	_, e = f.WriteString(body)
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 51709337a73d6eeef1736960282553ef1d6cfa6d..d94a7b34a50289d1d9eb401465bf9d42feca3efa
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -32,7 +32,7 @@
 		body, _ := contracts.Section(doc.Body, part.name)
 		items, e := contracts.ParseItems(body)
 		if e != nil {
-			return c, fixture.Coded("C03", path, e)
+			return c, fixture.Coded(fixture.ItemErrorCode(e), path, e)
 		}
 		*part.dest = items
 	}
@@ -75,7 +75,7 @@
 		}
 		path := "changes/" + ent.Name()
 		if slugs[slug] && !allowDuplicates {
-			return nil, fixture.Coded("C13", path, fmt.Errorf("duplicate pending slug %s", slug))
+			return nil, fixture.Coded("C13", path, fmt.Errorf("ambiguous change slug %s: duplicate pending slug", slug))
 		}
 		slugs[slug] = true
 		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
@@ -114,10 +114,10 @@
 		}
 	}
 	if len(matches) == 0 {
-		return -1, fixture.Coded("U02", target, fmt.Errorf("unknown change %s", target))
+		return -1, fixture.Coded("U02", target, fmt.Errorf("change %s not found", target))
 	}
 	if len(matches) > 1 {
-		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change %s", target))
+		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change slug %s", target))
 	}
 	return matches[0], nil
 }
@@ -370,11 +370,15 @@
 	for _, m := range s.Modules {
 		prefixes[m.IDPrefix] = m.Name
 	}
-	checkItem := func(it contracts.Item, path, module string, needBody bool) {
+	checkItem := func(it contracts.Item, path, module string, needBody, change bool) {
 		if strings.Contains(it.ID, "-") {
 			prefix := strings.SplitN(it.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module {
-				add("C05", fmt.Sprintf("foreign spec ID %s belongs to %s", it.ID, owner), path)
+				kind := "spec"
+				if change {
+					kind = "module"
+				}
+				add("C05", fmt.Sprintf("foreign %s ID %s belongs to %s", kind, it.ID, owner), path)
 			}
 		}
 		if needBody && !normRE.MatchString(it.Body) {
@@ -400,7 +404,7 @@
 				add("C05", "duplicate item ID "+it.ID, path)
 			}
 			seen[it.ID] = true
-			checkItem(it, path, m.Name, true)
+			checkItem(it, path, m.Name, true, false)
 		}
 	}
 	changes, e := queue(root, s, false)
@@ -436,10 +440,10 @@
 			}
 		}
 		for _, it := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
-			checkItem(it, c.Path, c.Module, true)
+			checkItem(it, c.Path, c.Module, true, true)
 		}
 		for _, it := range c.Removed {
-			checkItem(it, c.Path, c.Module, false)
+			checkItem(it, c.Path, c.Module, false, true)
 		}
 		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
 			if !added[ref] {

codex
The reported project failures now have targeted fixes. I also added stable codes for `NewChange` argument and duplicate-slug errors, plus clearer error causes used by callers. I’ll format the files and review the diff; this input workspace has no build manifest or dependency internals, so the native Go evals cannot run here.
exec
/bin/zsh -lc 'gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/queue/queue.go && git diff --no-index -- /dev/null packages/project/internal/fixture/fixture.go | head -20' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc "git status --short && rg -n 'func Coded|func Canonical|func Absolute|ItemErrorCode|missing module|func Load\\(|func NewChange|invalid slug|duplicate change slug|malformed recorded SPEC|ambiguous change slug|change .*not found|foreign .* ID' packages/project/internal/fixture/fixture.go packages/project/internal/queue/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 573ms:
?? DRILL_NOTES.md
?? packages/project/internal/
packages/project/internal/fixture/fixture.go:19:func Coded(code, subject string, err error) error {
packages/project/internal/fixture/fixture.go:31:func Canonical(root string) (string, error) {
packages/project/internal/fixture/fixture.go:38:func Absolute(root string) (string, error) { return filepath.Abs(root) }
packages/project/internal/fixture/fixture.go:39:func ItemErrorCode(err error) string {
packages/project/internal/fixture/fixture.go:81:		return nil, fmt.Errorf("missing modules globs")
packages/project/internal/fixture/fixture.go:171:				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
packages/project/internal/fixture/fixture.go:204:func Load(root string) (model.Snapshot, error) {
packages/project/internal/fixture/fixture.go:254:				return s, Coded(ItemErrorCode(e), path, e)
packages/project/internal/fixture/fixture.go:312:			cause := "malformed recorded SPEC"
packages/project/internal/fixture/fixture.go:344:func NewChange(root, module, slug string) (string, error) {
packages/project/internal/fixture/fixture.go:346:		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
packages/project/internal/fixture/fixture.go:376:				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
packages/project/internal/queue/queue.go:35:			return c, fixture.Coded(fixture.ItemErrorCode(e), path, e)
packages/project/internal/queue/queue.go:78:			return nil, fixture.Coded("C13", path, fmt.Errorf("ambiguous change slug %s: duplicate pending slug", slug))
packages/project/internal/queue/queue.go:117:		return -1, fixture.Coded("U02", target, fmt.Errorf("change %s not found", target))
packages/project/internal/queue/queue.go:120:		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change slug %s", target))
packages/project/internal/queue/queue.go:381:				add("C05", fmt.Sprintf("foreign %s ID %s belongs to %s", kind, it.ID, owner), path)

 succeeded in 599ms:
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000..125da22
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,395 @@
+package fixture
+
+import (
+	"context"
+	"errors"
+	"fmt"
+	"os"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index eb02d7bb5e00b2c3c9d74ed7b1f0c23de9f1e6fe..97dce1fea789c930735224a459f15e1e09dbd939
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -4,4 +4,4 @@
 
 ## Gaps
 
-- PRJ-007 (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite could not run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.
+- no spec (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite cannot run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..125da22d5c2ce63ed5644a77d06c154705c948ee
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -2,6 +2,7 @@
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"os"
 	"path/filepath"
@@ -16,6 +17,9 @@
 )
 
 func Coded(code, subject string, err error) error {
+	if code == "C05" {
+		err = fmt.Errorf("%w; remove or rename the conflicting spec ID", err)
+	}
 	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s %s: %w", code, subject, err)}
 }
 func Code(err error) string {
@@ -31,6 +35,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +168,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +202,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +251,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {
@@ -295,7 +306,18 @@
 			}
 		}
 	}
-	return loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
+	snapshot, e := loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
+	if e != nil {
+		if code, subject, ok := contracts.ErrorCode(e); ok && code == "C03" && strings.HasSuffix(subject, "/SPEC.md") {
+			cause := "malformed recorded SPEC"
+			if errors.Is(e, os.ErrNotExist) {
+				cause = "missing recorded SPEC"
+			}
+			return snapshot, Coded("C03", subject, fmt.Errorf("%s %s: %w", cause, subject, e))
+		}
+		return snapshot, e
+	}
+	return snapshot, nil
 }
 func TreePaths(root, revision string) ([]string, error) {
 	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "ls-tree", "-r", "--name-only", revision}})
@@ -321,7 +343,7 @@
 }
 func NewChange(root, module, slug string) (string, error) {
 	if !slugRE.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
 	}
 	s, e := Load(root)
 	if e != nil {
@@ -334,7 +356,7 @@
 		}
 	}
 	if !found {
-		return "", fmt.Errorf("unknown module %s", module)
+		return "", Coded("U02", module, fmt.Errorf("module %s not found", module))
 	}
 	max := 0
 	for _, dir := range []string{"changes", "changes/archive"} {
@@ -343,12 +365,15 @@
 			return "", e
 		}
 		for _, ent := range entries {
+			if ent.IsDir() {
+				continue
+			}
 			n, sl, ok := ChangeName(ent.Name())
 			if !ok {
 				continue
 			}
 			if sl == slug {
-				return "", fmt.Errorf("duplicate change slug %s", slug)
+				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
 			}
 			if n > max {
 				max = n
@@ -362,7 +387,7 @@
 	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
 	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
 	if e != nil {
-		return "", e
+		return "", Coded("P01", rel, fmt.Errorf("create change file %s: %w", rel, e))
 	}
 	defer f.Close()
 	_, e = f.WriteString(body)
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 51709337a73d6eeef1736960282553ef1d6cfa6d..d94a7b34a50289d1d9eb401465bf9d42feca3efa
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -32,7 +32,7 @@
 		body, _ := contracts.Section(doc.Body, part.name)
 		items, e := contracts.ParseItems(body)
 		if e != nil {
-			return c, fixture.Coded("C03", path, e)
+			return c, fixture.Coded(fixture.ItemErrorCode(e), path, e)
 		}
 		*part.dest = items
 	}
@@ -75,7 +75,7 @@
 		}
 		path := "changes/" + ent.Name()
 		if slugs[slug] && !allowDuplicates {
-			return nil, fixture.Coded("C13", path, fmt.Errorf("duplicate pending slug %s", slug))
+			return nil, fixture.Coded("C13", path, fmt.Errorf("ambiguous change slug %s: duplicate pending slug", slug))
 		}
 		slugs[slug] = true
 		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
@@ -114,10 +114,10 @@
 		}
 	}
 	if len(matches) == 0 {
-		return -1, fixture.Coded("U02", target, fmt.Errorf("unknown change %s", target))
+		return -1, fixture.Coded("U02", target, fmt.Errorf("change %s not found", target))
 	}
 	if len(matches) > 1 {
-		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change %s", target))
+		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change slug %s", target))
 	}
 	return matches[0], nil
 }
@@ -370,11 +370,15 @@
 	for _, m := range s.Modules {
 		prefixes[m.IDPrefix] = m.Name
 	}
-	checkItem := func(it contracts.Item, path, module string, needBody bool) {
+	checkItem := func(it contracts.Item, path, module string, needBody, change bool) {
 		if strings.Contains(it.ID, "-") {
 			prefix := strings.SplitN(it.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module {
-				add("C05", fmt.Sprintf("foreign spec ID %s belongs to %s", it.ID, owner), path)
+				kind := "spec"
+				if change {
+					kind = "module"
+				}
+				add("C05", fmt.Sprintf("foreign %s ID %s belongs to %s", kind, it.ID, owner), path)
 			}
 		}
 		if needBody && !normRE.MatchString(it.Body) {
@@ -400,7 +404,7 @@
 				add("C05", "duplicate item ID "+it.ID, path)
 			}
 			seen[it.ID] = true
-			checkItem(it, path, m.Name, true)
+			checkItem(it, path, m.Name, true, false)
 		}
 	}
 	changes, e := queue(root, s, false)
@@ -436,10 +440,10 @@
 			}
 		}
 		for _, it := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
-			checkItem(it, c.Path, c.Module, true)
+			checkItem(it, c.Path, c.Module, true, true)
 		}
 		for _, it := range c.Removed {
-			checkItem(it, c.Path, c.Module, false)
+			checkItem(it, c.Path, c.Module, false, true)
 		}
 		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
 			if !added[ref] {

exec
/bin/zsh -lc 'gofmt -l packages/project/internal/fixture/fixture.go packages/project/internal/queue/queue.go packages/project/internal/queue/stale.go packages/project/internal/history/history.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'git status --short && cat DRILL_NOTES.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
exec
/bin/zsh -lc 'go test ./packages/project/...' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574
 succeeded in 580ms:
 succeeded in 592ms:
?? DRILL_NOTES.md
?? packages/project/internal/
## Contradicted Decisions

None identified.

## Gaps

- no spec (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite cannot run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.

 exited 1 in 616ms:
FAIL	./packages/project/... [setup failed]
# ./packages/project/...
pattern ./packages/project/...: directory prefix packages/project does not contain main module or its selected dependencies
FAIL

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/packages/project/internal/fixture/fixture.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index eb02d7bb5e00b2c3c9d74ed7b1f0c23de9f1e6fe..97dce1fea789c930735224a459f15e1e09dbd939
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -4,4 +4,4 @@
 
 ## Gaps
 
-- PRJ-007 (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite could not run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.
+- no spec (project): The supplied input has no `go.mod` or dependency implementations, so the native Go eval suite cannot run here. `gofmt` completed; the separate native verifier must run the retained project and dependent evals.
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index 660d0f35bed549919450d433f6841f16f8edd42c..540816a8ae6963ce4348daa7847c00943117fa5a
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -2,6 +2,7 @@
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"os"
 	"path/filepath"
@@ -16,6 +17,9 @@
 )
 
 func Coded(code, subject string, err error) error {
+	if code == "C05" {
+		err = fmt.Errorf("%w; remove or rename the conflicting spec ID", err)
+	}
 	return contracts.CodedError{Code: code, Subject: subject, Cause: fmt.Errorf("%s %s: %w", code, subject, err)}
 }
 func Code(err error) string {
@@ -31,6 +35,13 @@
 	}
 	return filepath.EvalSymlinks(p)
 }
+func Absolute(root string) (string, error) { return filepath.Abs(root) }
+func ItemErrorCode(err error) string {
+	if strings.Contains(err.Error(), "duplicate item ID") {
+		return "C05"
+	}
+	return Code(err)
+}
 func ReadConfig(root string) (map[string]any, error) {
 	b, e := os.ReadFile(filepath.Join(root, "fireweed.yaml"))
 	if e != nil {
@@ -157,7 +168,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if !names[d] {
-				return fmt.Errorf("module %s depends on missing dependency %s", m.Name, d)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, d)
 			}
 		}
 	}
@@ -191,7 +202,7 @@
 	return nil
 }
 func Load(root string) (model.Snapshot, error) {
-	canon, e := Canonical(root)
+	canon, e := Absolute(root)
 	if e != nil {
 		return model.Snapshot{}, Coded("C01", root, e)
 	}
@@ -240,7 +251,7 @@
 		if section, ok := contracts.Section(doc.Body, "Specs"); ok {
 			items, e := contracts.ParseItems(section)
 			if e != nil {
-				return s, Coded(Code(e), path, e)
+				return s, Coded(ItemErrorCode(e), path, e)
 			}
 			seen := map[string]bool{}
 			for _, it := range items {
@@ -295,7 +306,18 @@
 			}
 		}
 	}
-	return loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
+	snapshot, e := loadModules(root, config, defaults, dirs, func(path string) ([]byte, error) { return gitstore.ReadAt(root, revision, path) })
+	if e != nil {
+		if code, subject, ok := contracts.ErrorCode(e); ok && code == "C03" && strings.HasSuffix(subject, "/SPEC.md") {
+			cause := "malformed recorded SPEC"
+			if errors.Is(e, os.ErrNotExist) {
+				cause = "missing recorded SPEC"
+			}
+			return snapshot, Coded("C03", subject, fmt.Errorf("%s: %w", cause, e))
+		}
+		return snapshot, e
+	}
+	return snapshot, nil
 }
 func TreePaths(root, revision string) ([]string, error) {
 	r, e := gitstore.Run(context.Background(), gitstore.Request{Dir: root, Args: []string{"git", "ls-tree", "-r", "--name-only", revision}})
@@ -321,7 +343,7 @@
 }
 func NewChange(root, module, slug string) (string, error) {
 	if !slugRE.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Coded("U01", slug, fmt.Errorf("invalid slug %q; use lowercase letters, digits and hyphens", slug))
 	}
 	s, e := Load(root)
 	if e != nil {
@@ -334,7 +356,7 @@
 		}
 	}
 	if !found {
-		return "", fmt.Errorf("unknown module %s", module)
+		return "", Coded("U02", module, fmt.Errorf("module %s not found", module))
 	}
 	max := 0
 	for _, dir := range []string{"changes", "changes/archive"} {
@@ -343,12 +365,15 @@
 			return "", e
 		}
 		for _, ent := range entries {
+			if ent.IsDir() {
+				continue
+			}
 			n, sl, ok := ChangeName(ent.Name())
 			if !ok {
 				continue
 			}
 			if sl == slug {
-				return "", fmt.Errorf("duplicate change slug %s", slug)
+				return "", Coded("C13", slug, fmt.Errorf("duplicate change slug %s; choose an unused slug", slug))
 			}
 			if n > max {
 				max = n
@@ -362,7 +387,7 @@
 	body := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
 	f, e := os.OpenFile(filepath.Join(s.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
 	if e != nil {
-		return "", e
+		return "", Coded("P01", rel, fmt.Errorf("create change file %s: %w", rel, e))
 	}
 	defer f.Close()
 	_, e = f.WriteString(body)
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 51709337a73d6eeef1736960282553ef1d6cfa6d..d94a7b34a50289d1d9eb401465bf9d42feca3efa
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -32,7 +32,7 @@
 		body, _ := contracts.Section(doc.Body, part.name)
 		items, e := contracts.ParseItems(body)
 		if e != nil {
-			return c, fixture.Coded("C03", path, e)
+			return c, fixture.Coded(fixture.ItemErrorCode(e), path, e)
 		}
 		*part.dest = items
 	}
@@ -75,7 +75,7 @@
 		}
 		path := "changes/" + ent.Name()
 		if slugs[slug] && !allowDuplicates {
-			return nil, fixture.Coded("C13", path, fmt.Errorf("duplicate pending slug %s", slug))
+			return nil, fixture.Coded("C13", path, fmt.Errorf("ambiguous change slug %s: duplicate pending slug", slug))
 		}
 		slugs[slug] = true
 		b, e := os.ReadFile(filepath.Join(s.Root, filepath.FromSlash(path)))
@@ -114,10 +114,10 @@
 		}
 	}
 	if len(matches) == 0 {
-		return -1, fixture.Coded("U02", target, fmt.Errorf("unknown change %s", target))
+		return -1, fixture.Coded("U02", target, fmt.Errorf("change %s not found", target))
 	}
 	if len(matches) > 1 {
-		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change %s", target))
+		return -1, fixture.Coded("C13", target, fmt.Errorf("ambiguous change slug %s", target))
 	}
 	return matches[0], nil
 }
@@ -370,11 +370,15 @@
 	for _, m := range s.Modules {
 		prefixes[m.IDPrefix] = m.Name
 	}
-	checkItem := func(it contracts.Item, path, module string, needBody bool) {
+	checkItem := func(it contracts.Item, path, module string, needBody, change bool) {
 		if strings.Contains(it.ID, "-") {
 			prefix := strings.SplitN(it.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module {
-				add("C05", fmt.Sprintf("foreign spec ID %s belongs to %s", it.ID, owner), path)
+				kind := "spec"
+				if change {
+					kind = "module"
+				}
+				add("C05", fmt.Sprintf("foreign %s ID %s belongs to %s", kind, it.ID, owner), path)
 			}
 		}
 		if needBody && !normRE.MatchString(it.Body) {
@@ -400,7 +404,7 @@
 				add("C05", "duplicate item ID "+it.ID, path)
 			}
 			seen[it.ID] = true
-			checkItem(it, path, m.Name, true)
+			checkItem(it, path, m.Name, true, false)
 		}
 	}
 	changes, e := queue(root, s, false)
@@ -436,10 +440,10 @@
 			}
 		}
 		for _, it := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
-			checkItem(it, c.Path, c.Module, true)
+			checkItem(it, c.Path, c.Module, true, true)
 		}
 		for _, it := range c.Removed {
-			checkItem(it, c.Path, c.Module, false)
+			checkItem(it, c.Path, c.Module, false, true)
 		}
 		for _, ref := range handleRE.FindAllString(c.Document.Body, -1) {
 			if !added[ref] {

codex
Rebuilt the disposable project implementation to address the reported `Load` root mismatch, C05 duplicate and foreign-ID diagnostics, and `NewChange` error codes. Updated [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-1212345574/DRILL_NOTES.md). Retained inputs were untouched; nothing was committed.

`gofmt` passed. I could not run the Go evals here because the supplied workspace has no `go.mod` or dependency implementations, so the fixes remain unverified by the native suite.
tokens used
58,542


## Verification failures
app eval: E01: failed case "Test_APP_038_LookupCauses" for APP-038
workflow eval: E01: failed case "Test_FLOW_058_FLOW_062_StaleLifecycleBaseKeepsCauseAndMissingBaseKeepsC08" for FLOW-058
