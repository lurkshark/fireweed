---
date: 2026-11-10T16:44:27.736325Z
base_commit: 73e269df5e7ca04be4451d5bc37390712c623483
agent: codex
model: gpt-6.1-sol
result: pass
---

## Own evals
- project | Test_PRJ_057_StampInvalidBasesHasCode | PRJ-057 | pass
- project | Test_PRJ_058_MissingPendingChangeHasSubject | PRJ-058 | pass
- project | Test_PRJ_058_MissingPendingChangeHasSubject/Stamp | PRJ-058 | pass
- project | Test_PRJ_058_MissingPendingChangeHasSubject/Project | PRJ-058 | pass
- project | Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath | PRJ-059, PRJ-068 | pass
- project | Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath/malformed_frontmatter | PRJ-059, PRJ-068 | pass
- project | Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath/malformed_spec_heading | PRJ-059, PRJ-068 | pass
- project | Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath/duplicate_spec_id | PRJ-059, PRJ-068 | pass
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
- project | Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause | PRJ-067 | pass
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
- project | Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs | PRJ-004 | pass
- project | Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs/Project | PRJ-004 | pass
- project | Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs/Stamp | PRJ-004 | pass
- project | Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs | PRJ-066 | pass
- project | Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs/StaleSpecs | PRJ-066 | pass
- project | Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs/StampMissing | PRJ-066 | pass
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
- project | Test_PRJ_069_RegularFileAtChangesReportsQueueCause | PRJ-069 | pass
- project | Test_PRJ_045_CheckPendingModuleTargetCause | PRJ-045 | pass
- project | Test_PRJ_013_CheckDuplicateWithinEditSection | PRJ-013 | pass
  Stdout: ?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.348s


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
ok  	github.com/lurkshark/fireweed/packages/app/evals	260.853s

- checks | Test_CHK_005_GlobalAndModuleBoundaryHooks | CHK-005 | pass
- checks | Test_CHK_041_ExactBaselineOnly | CHK-041 | pass
- checks | Test_CHK_039_CHK_040_OptOutSuppressesOnlyItsModuleAndWarns | CHK-039, CHK-040 | pass
- checks | Test_CHK_044_CHK_045_MissingHookRequiredVersusOptional | CHK-044, CHK-045 | pass
- checks | Test_CHK_041_CHK_042_NormalNonzeroExitCanUseBaselineAndOptOut | CHK-041, CHK-042 | pass
- checks | Test_CHK_043_RequiredGlobalHookRejectsUnattributedNonzeroExit | CHK-043 | pass
- checks | Test_CHK_043_SilentNonzeroOverrideIsIndependentOfGlobalViolation | CHK-043 | pass
- checks | Test_CHK_043_LaunchCancellationAndTimeoutStayErrors | CHK-043 | pass
- checks | Test_CHK_081_DoctorProbeLogsShareFreshRunAndPreserveStreams | CHK-081 | pass
- checks | Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput | CHK-081 | pass
- checks | Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput/failure | CHK-081 | pass
- checks | Test_CHK_081_DoctorPersistsFailedAndTimedOutProbeOutput/timeout | CHK-081 | pass
- checks | Test_CHK_081_DoctorPersistsCancelledProbeOutput | CHK-081 | pass
- checks | Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary | CHK-059 | pass
- checks | Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/declared_test | CHK-059 | pass
- checks | Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/subtest | CHK-059 | pass
- checks | Test_CHK_059_DoctorMatchesSuppliedGoMissingIDCanary/unrelated_test_with_shared_prefix | CHK-059 | pass
- checks | Test_CHK_082_CHK_083_DoctorExpandsSelectedCommandsAndWrapperOnce | CHK-082, CHK-083 | pass
- checks | Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures | CHK-083 | pass
- checks | Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/launch | CHK-083 | pass
- checks | Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/nonzero | CHK-083 | pass
- checks | Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/timeout | CHK-083 | pass
- checks | Test_CHK_083_DoctorRemovesPromptFileOnProbeFailures/cancel | CHK-083 | pass
- checks | Test_CHK_083_DoctorDoesNotCreatePromptFileWithoutSelectedPlaceholder | CHK-083 | pass
- checks | Test_CHK_082_CHK_083_DoctorWrapperOnlyFileWithAuditFallback | CHK-082, CHK-083 | pass
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
- checks | Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports | CHK-034 | pass
- checks | Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports/unmapped_PackageFailure | CHK-034 | pass
- checks | Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports/malformed_XML | CHK-034 | pass
- checks | Test_CHK_034_EvaluateRetainsHookOutputOnRejectedReports/missing_report | CHK-034 | pass
- checks | Test_CHK_034_EvaluateRetainsOutputOnHookTimeout | CHK-034 | pass
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
ok  	github.com/lurkshark/fireweed/packages/checks/evals	266.575s

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
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	276.471s

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
- workflow | Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState/ignored | FLOW-063, FLOW-064, FLOW-065, FLOW-066, FLOW-067 | pass
- workflow | Test_FLOW_063_FLOW_064_FLOW_065_FLOW_066_FLOW_067_BlockSavesHeldEvalsWorkAndCommitsOnlyState/unignored | FLOW-063, FLOW-064, FLOW-065, FLOW-066, FLOW-067 | pass
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
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal | FLOW-032, FLOW-033, FLOW-046, FLOW-135 | pass
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal/default_audit/ignored | FLOW-032, FLOW-033, FLOW-046, FLOW-135 | pass
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal/default_audit/unignored | FLOW-032, FLOW-033, FLOW-046, FLOW-135 | pass
- workflow | Test_FLOW_032_FLOW_033_FLOW_046_FLOW_135_EvalsDoneAdvancesAndKeepsRunArtifactsLocal/disabled_audit/ignored | FLOW-032, FLOW-033, FLOW-046, FLOW-135 | pass
- workflow | Test_FLOW_034_FLOW_036_FLOW_041_FLOW_046_AuditAcceptsOnlyExactClauseFindingsAndCheckedNoFindings | FLOW-034, FLOW-036, FLOW-041, FLOW-046 | pass
- workflow | Test_FLOW_035_FLOW_037_FLOW_038_FLOW_039_FLOW_042_FLOW_043_FLOW_046_AuditFindingsConvergeAndPersistTwoRoundBound | FLOW-035, FLOW-037, FLOW-038, FLOW-039, FLOW-042, FLOW-043, FLOW-046 | pass
- workflow | Test_FLOW_040_FLOW_042_FLOW_043_UntestableFindingBlocksWithQuestion | FLOW-040, FLOW-042, FLOW-043 | pass
- workflow | Test_FLOW_004_FLOW_028_FLOW_044_FLOW_045_FLOW_046_FLOW_047_FLOW_050_ImplDoneAppliesSpecsSetAndDecisionsAndArchives | FLOW-004, FLOW-028, FLOW-044, FLOW-045, FLOW-046, FLOW-047, FLOW-050 | pass
- workflow | Test_FLOW_003_FLOW_024_FLOW_048_FLOW_049_FailedDoneKeepsPhaseAndHead | FLOW-003, FLOW-024, FLOW-048, FLOW-049 | pass
- workflow | Test_FLOW_005_FLOW_051_FLOW_052_FLOW_053_ContextAndProvenanceAndRetainedReports | FLOW-005, FLOW-051, FLOW-052, FLOW-053 | pass
- workflow | Test_FLOW_052_DoneRetainsCostModuleReportBesideCostSnapshot | FLOW-052 | pass
- workflow | Test_FLOW_052_DoneRetainsCostModuleReportBesideCostSnapshot/costs | FLOW-052 | pass
- workflow | Test_FLOW_052_DoneRetainsCostModuleReportBesideCostSnapshot/CoStS | FLOW-052 | pass
- workflow | Test_FLOW_054_StartOperationalPreconditionsHaveP01 | FLOW-054 | pass
- workflow | Test_FLOW_056_StartContentionAndCurrentForeignClaimHaveP02 | FLOW-056 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks | FLOW-057 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/author | FLOW-057 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/runtime | FLOW-057 | pass
- workflow | Test_FLOW_057_VerifyAndDoneRejectApprovedStateEditsBeforeHooks/audit | FLOW-057 | pass
- workflow | Test_FLOW_059_AuditNoFindingsErrorsHaveP09 | FLOW-059 | pass
- workflow | Test_FLOW_060_BlockedDoneKeepsCommitQuestionsAndReleasesClaim | FLOW-060 | pass
- workflow | Test_FLOW_062_LifecyclePreservesEvalZoneAndBoundaryCauses | FLOW-062 | pass
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
PASS
ok  	github.com/lurkshark/fireweed/packages/workflow/evals	135.181s


## Contradicted Decisions
None. Spec history uses the current archive files, returns section contents
without their headings, and returns only the matching ID for an ID-filtered
entry, as specified by 0126-spec-history-query.
## Gaps
- **PRJ-001 / project.** The design calls for a warning about SPEC.md files
  outside configured module globs. The project spec assigns no diagnostic for
  those files. Discovery stays within the configured globs.
- **PRJ-004 and PRJ-032 / project.** The specs leave open whether a malformed
  document after the selected change invalidates a partial projection. A
  partial projection validates queue identities across the queue and parses
  documents through the selected change.
- **PRJ-004, PRJ-058 and PRJ-066 / project.** APP-038 expects the CLI to
  render U02 once. Pending selector errors carry U02 as metadata and provide
  the rejected selector and queue guidance in the cause text.
- **PRJ-016 / project.** The retained eval requires two C08 diagnostics for an
  absent edit target, with generic and section-specific wording. The spec
  requires C08 and leaves the number and wording of diagnostics open.
- **PRJ-021 and PRJ-041 / project.** APP-037 expects the phrase `missing module`
  for an absent dependency. The project specs require an explanatory cause
  and C02 without prescribing the phrase.
- **PRJ-023 and PRJ-027 / project.** The filename format leaves numbering
  beyond 9999 implicit. The queue accepts at least four decimal digits so a
  change created after 9999 remains discoverable.
- **PRJ-024 / project.** APP-019 expects C13 for duplicate slugs in the queue
  and archive. The project spec requires rejection and file preservation and
  leaves the code open.
- **PRJ-031 and PRJ-044 / project.** APP-037 and the APP-116/APP-117 eval
  expect the phrase `ambiguous change slug`. The project specs require
  rejection and C13 and leave the wording open.
- **PRJ-035 / project.** The spec leaves validation timing open when a later
  change repairs a graph made invalid by an earlier change. Each projected
  graph is validated as its change is applied.
- **PRJ-040 and PRJ-056 / project.** APP-037 expects the CLI to render C01
  once. Load exposes C01 as error metadata and Check carries it separately
  from the message. SpecHistory adds the textual prefix required by PRJ-056.
- **PRJ-060 and PRJ-061 / project.** The specs leave the stale result for a
  removed current item and the Git history traversal unspecified. An item with
  a recorded base and no current counterpart has empty Current text. Recovery
  follows HEAD's first-parent history at the queued path.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt `project` within its disposable paths. The target eval hook passed all 134 cases, and `go vet ./packages/project/...` passed. `fireweed verify` returned P02 because the drill workspace has no held change.

[DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/DRILL_NOTES.md) records six spec gaps and no contradicted Decisions. It documents the choices to validate the graph after each applied change and use first-parent Git history for stale-base recovery.

Reading additional input from stdin...
2026-11-10T16:44:29.726076Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1c658-c57c-7c62-9b5b-fbda5d170171
--------
user
# Rebuild project

## Target spec and Decisions

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

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module with an error beginning
`unknown module ` and naming the rejected module.

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
An unknown module or an ID absent from both current module Specs and matching
archive edits MUST return an error beginning `U02: ` and naming the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists
in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce an error beginning `C01: `.
Malformed archived change documents or malformed Added/Modified item syntax
MUST produce an error beginning `C03: ` and identifying the archive path.

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
An unknown selector MUST return U02 naming the requested change and listing valid pending change slugs; an ambiguous slug MUST return C13.

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

### PRJ-069
When `changes` is a regular file, Queue MUST return an error identifying `changes`. Check MUST report a non-warning C03 diagnostic identifying `changes`.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.


## Retained interfaces and files

packages/project/SPEC.md
packages/project/api.go
packages/project/model/model.go

## Retained evals

packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/selector_guidance_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go

Rebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.

Run the target's eval hook from the workspace root as often as needed.

~~~sh
mkdir -p .fireweed/runs/drill && sh -c 'go run .fireweed/hooks.go eval '"'"'packages/project'"'"' '"'"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/.fireweed/runs/drill/report.xml'"'"''
~~~

Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module.

codex
I’ll read the rebuild instructions and retained interfaces, then implement the project package within its disposable paths and run the target eval hook.
exec
/bin/zsh -lc 'pwd; git status --short; ls -la; cat SYSTEM.md; cat docs/rebuild/PROGRESS.md; cat docs/rebuild/DECISIONS.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc "rg --files --hidden -g '"'!.git/**'"' packages .fireweed | sort; cat fireweed.yaml; cat go.mod; cat packages/project/api.go; cat packages/project/model/model.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 492ms:
.fireweed/canaries/canaries.yaml
.fireweed/hooks.go
.fireweed/recipe.md
.fireweed/setup.sh
.fireweed/templates/module/SPEC.md
.fireweed/templates/module/api.go.tmpl
.fireweed/templates/module/evals/.gitkeep
.fireweed/templates/module/internal/.gitkeep
packages/assets/SPEC.md
packages/assets/api.go
packages/assets/evals/.gitkeep
packages/assets/evals/additional_recipes_test.go
packages/assets/evals/assets_test.go
packages/assets/evals/decision_notes_playbooks_test.go
packages/assets/evals/eval_fix_playbooks_test.go
packages/assets/evals/go_recipe_id_guidance_test.go
packages/assets/evals/go_recipe_test.go
packages/assets/evals/history_explanations_test.go
packages/assets/evals/mixed_eval_fix_explanation_test.go
packages/assets/evals/pr_body_artifact_explanation_test.go
packages/assets/evals/regenerate_warning_explanation_test.go
packages/assets/evals/status_warning_explanations_test.go
packages/assets/evals/stub_handoff_test.go
packages/assets/internal/.gitkeep
packages/assets/internal/content/content.go
packages/assets/internal/recipe/files/SPEC.md
packages/assets/internal/recipe/files/api.go.tmpl
packages/assets/internal/recipe/files/canaries.yaml
packages/assets/internal/recipe/files/go.md
packages/assets/internal/recipe/files/hooks.go
packages/assets/internal/recipe/files/python-uv-SPEC.md
packages/assets/internal/recipe/files/python-uv-canaries.yaml
packages/assets/internal/recipe/files/python-uv-hooks.py
packages/assets/internal/recipe/files/python-uv-init.py.tmpl
packages/assets/internal/recipe/files/python-uv-pyproject.toml.tmpl
packages/assets/internal/recipe/files/python-uv.md
packages/assets/internal/recipe/files/typescript-pnpm-SPEC.md
packages/assets/internal/recipe/files/typescript-pnpm-canaries.yaml
packages/assets/internal/recipe/files/typescript-pnpm-hooks.mjs
packages/assets/internal/recipe/files/typescript-pnpm-index.ts.tmpl
packages/assets/internal/recipe/files/typescript-pnpm-package.json.tmpl
packages/assets/internal/recipe/files/typescript-pnpm.md
packages/assets/internal/recipe/recipe.go
packages/assets/model/model.go
packages/contracts/SPEC.md
packages/contracts/api.go
packages/contracts/evals/.gitkeep
packages/contracts/evals/contracts_test.go
packages/contracts/internal/.gitkeep
packages/contracts/internal/primitives/coded_error.go
packages/contracts/internal/primitives/primitives.go
packages/contracts/model/model.go
packages/gitstore/SPEC.md
packages/gitstore/api.go
packages/gitstore/evals/.gitkeep
packages/gitstore/evals/gitstore_test.go
packages/gitstore/internal/local/local.go
packages/gitstore/model/model.go
packages/project/SPEC.md
packages/project/api.go
packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/selector_guidance_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go
packages/project/model/model.go
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
module github.com/lurkshark/fireweed

go 1.23

require gopkg.in/yaml.v3 v3.0.1
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

 succeeded in 498ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
total 280
drwx------@   17 joel  staff     544 Nov 10 08:44 .
drwx------@ 6449 joel  staff  206368 Nov 10 08:44 ..
drwxr-xr-x@    8 joel  staff     256 Nov 10 08:44 .fireweed
drwxr-xr-x@    9 joel  staff     288 Nov 10 08:44 .git
-rw-r--r--@    1 joel  staff      56 Nov 10 08:44 .gitignore
-rw-r--r--@    1 joel  staff     859 Nov 10 08:44 AGENTS.md
-rw-r--r--@    1 joel  staff   35149 Nov 10 08:44 LICENSE
-rw-r--r--@    1 joel  staff     100 Nov 10 08:44 Makefile
-rw-r--r--@    1 joel  staff    3399 Nov 10 08:44 README.md
-rw-r--r--@    1 joel  staff    8073 Nov 10 08:44 SYSTEM.md
drwxr-xr-x@    3 joel  staff      96 Nov 10 08:44 docs
-rw-r--r--@    1 joel  staff   61761 Nov 10 08:44 fireweed-design-v1.md
-rw-r--r--@    1 joel  staff    1289 Nov 10 08:44 fireweed.yaml
-rw-r--r--@    1 joel  staff      79 Nov 10 08:44 go.mod
-rw-r--r--@    1 joel  staff     151 Nov 10 08:44 go.sum
-rw-r--r--@    1 joel  staff     146 Nov 10 08:44 main.go
drwxr-xr-x@    6 joel  staff     192 Nov 10 08:44 packages
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

The owner approved the `0197` worker and restricted audit contexts for OpenAI
Codex and granted standing authorization for the configured change lifecycle.
[SYSTEM.md](../../SYSTEM.md#agent-configuration) records the scope and destination.

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

**Project proof renewal awaits the PRJ-069 audit.** The verified controller ran
`fireweed drill project` at `d4e8d95`. All 133 project cases and 790 dependent
cases passed after the configured retry. One app case failed. The
[native report](../../.fireweed/drills/project.md) is committed in `dfbb207`.
Fireweed cleaned up both scratch locations and retained the reconstruction
`917fb02dba3b227f5c57ed422455e7e5881735c7` at
`refs/fireweed/drills/project/20261110T144729.916008000`. Its changes are
confined to disposable project paths.

`Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors` requires a C03
error from `project.Check` when `changes` is a regular file. The reconstruction
classified the queue read failure as C13. The same focused app eval passed
against the current product. The project queue and Check specs leave this
diagnostic implicit.

The owner approved making C03 explicit. Fireweed started
[0201-queue-path-diagnostic](../../changes/0201-queue-path-diagnostic.md)
and assigned `PRJ-069`. The project eval passed with the other 133 cases.
Audit found that the change also promises C03 on the public `Queue` error,
whose classification is private today. The eval asserts the path on `Queue`
and the C03 diagnostic on `Check`. App derives C03 from that diagnostic.

Both auditor submissions omitted the required dash before the finding's
explanation. The controller rejected them with P09 and blocked `0201` in
`d5ea3be`. It saved the audit work at
`refs/fireweed/attempts/queue-path-diagnostic/1` (`dc223bd`) and released the
claim. Audit rounds remain at zero.

The owner approved narrowing `PRJ-069` to the behavior the existing eval
asserts. `Queue` rejects a regular-file `changes` path and identifies it in
the error. `Check` reports a non-warning C03 diagnostic identifying the same
path. Author commit `fde499d` made the revision and removed `blocked` using
the documented resume procedure, preserving the assigned ID, run and audit
phase. `fireweed check` and the focused `PRJ-069` and `APP-116`/`APP-117` evals
passed. Automatic approval review rejected the lifecycle command before it
ran because it required specific authorization to send repository context to
OpenAI Codex. No claim is active. The saved attempt and failed drill ref remain
available.

The report lists no contradicted Decisions. Its other observations cover
diagnostic codes and wording relied on by app, duplicate C08 messages,
change numbering past 9999, intermediate graph validation, and stale-view
and history behavior. It also identifies the design's warning for specs
outside configured globs as an unassigned contract.

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

Reassess reuse if an exact historical proof is recovered from a backup.
Checks needs current proof of the `0197` CHK-043 change. `runner` and `app`
need first native reports and retained refs.

## Next step

Obtain explicit authorization for the configured Codex audit, implementation
and project reconstruction contexts. Resume `0201` at audit with the verified
controller, then renew `project` after Fireweed applies it. Commit the native
report, retain the proof ref and stop for owner review before `checks`. The
remaining drill order is `project`, `checks`, `runner`, then `app` in the
[regeneration goal](GOAL.md).

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

exec
/bin/zsh -lc "ls -la .fireweed .fireweed/bin changes; cat AGENTS.md; git log -1 --oneline; git ls-files 'packages/project/*'; export PATH=\""'$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; if [ -x .fireweed/bin/fireweed ]; then fireweed playbook author-change; else printf '"'%s\\n' 'Controller absent'; fi; rg -n 'func |"'^type |''^const '"' packages/{contracts,gitstore,assets}/{api.go,model/model.go}" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'cat packages/project/evals/{coded_errors,mixed_eval_fix,queue,selector_guidance,static_codes}_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'cat packages/project/evals/{project,revision_snapshot,spec_history,stale}_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc "rg -n 'func |project|go test' .fireweed/hooks.go; rg -n 'module_defaults|"'^#.*(Changes|Project|Module|Spec)|Provenance|outside .*globs|bases:'"' fireweed-design-v1.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 628ms:
package evals_test

import (
	"errors"
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
			_, err := project.NewChange(root, tc.module, tc.slug)
			if err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			} else if tc.id == "PRJ-025" && (!strings.HasPrefix(err.Error(), "unknown module ") || !strings.Contains(err.Error(), tc.module)) {
				t.Errorf("unknown module error = %q; want prefix and rejected module %q", err, tc.module)
			}
		})
	}
}

func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
	root := fixture(t)
	const queuedPath = "changes/0001-blocked.md"
	if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(queuedPath)), 0755); err != nil {
		t.Fatal(err)
	}
	path, err := project.NewChange(root, "alpha", "blocked")
	if err == nil {
		t.Fatal("NewChange accepted a queued path occupied by a directory")
	}
	if path != "" {
		t.Errorf("rejected NewChange path = %q; want empty", path)
	}
	if !strings.HasPrefix(err.Error(), "create "+queuedPath+":") {
		t.Errorf("NewChange error = %q; want repository-relative create prefix", err)
	}
	var pathErr *os.PathError
	if !errors.As(err, &pathErr) || pathErr.Path != filepath.Join(root, filepath.FromSlash(queuedPath)) || pathErr.Err == nil {
		t.Errorf("NewChange error = %v; want retained filesystem cause for %s", err, queuedPath)
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
		if err == nil || !strings.HasPrefix(err.Error(), "U02: ") || !strings.Contains(err.Error(), tc.rejected) {
			t.Errorf("lookup module %q ID %q = %v; want U02: naming %q", tc.module, tc.id, err, tc.rejected)
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
			if err == nil || !strings.HasPrefix(err.Error(), "C01: ") {
				t.Errorf("invalid config = %v; want C01: prefix", err)
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
			if err == nil || !strings.HasPrefix(err.Error(), "C03: ") || !strings.Contains(err.Error(), path) {
				t.Errorf("malformed archive = %v; want C03: naming %s", err, path)
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

 succeeded in 633ms:
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

func Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath(t *testing.T) {
	for _, tc := range []struct {
		name, spec, code string
	}{
		{"malformed_frontmatter", "---\nname: [broken\n---\n# Alpha\n", "C03"},
		{"malformed_spec_heading", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001 [broken\nAlpha MUST work.\n", "C03"},
		{"duplicate_spec_id", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n\n### ALP-001\nAlpha MUST also work.\n", "C05"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			const path = "packages/alpha/SPEC.md"
			put(t, root, path, tc.spec)
			_, err := project.Load(root)
			if err == nil {
				t.Fatal("Load accepted malformed module SPEC.md")
			}
			if code, _, ok := contracts.ErrorCode(err); !ok || code != tc.code {
				t.Errorf("Load error code = (%q, %v); want (%s, true): %v", code, ok, tc.code, err)
			}
			found := false
			for _, diagnostic := range project.Check(root) {
				if !diagnostic.Warning && diagnostic.Code == tc.code && diagnostic.Path == path {
					found = true
				}
			}
			if !found {
				t.Errorf("Check diagnostics = %#v; want %s identifying %s", project.Check(root), tc.code, path)
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
			if tc.name == "missing_base" {
				for _, variant := range []struct{ name, front string }{
					{"null", "bases:\n  ALP-001: null\n"},
					{"empty_string", "bases:\n  ALP-001: \"\"\n"},
				} {
					for _, edit := range []string{"Modified", "Removed"} {
						modified, removed := "### ALP-001\nAlpha MUST change.\n", ""
						if edit == "Removed" {
							modified, removed = "", "### ALP-001\n"
						}
						queued(t, root, "0001-check.md", variant.front, "Needed.\n", "", modified, removed)
						found := false
						for _, diagnostic := range project.Check(root) {
							if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "missing base for ALP-001") {
								found = true
							}
						}
						if !found {
							t.Errorf("Check did not report C08 missing base for %s with %s entry", edit, variant.name)
						}
					}
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
	"errors"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if err == nil {
		t.Fatal("unknown pending selector was accepted")
	}
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	var coded contracts.CodedError
	if !errors.As(err, &coded) || coded.Cause == nil {
		t.Errorf("lookup error = %v; want a retained U02 cause", err)
	}
	message := err.Error()
	for _, want := range []string{selector, "first", "second"} {
		if !strings.Contains(message, want) {
			t.Errorf("lookup error %q lacks rejected selector or pending slug %q", message, want)
		}
	}
	for _, unwanted := range []string{"0001-first.md", "0002-second.md", "archived"} {
		if strings.Contains(message, unwanted) {
			t.Errorf("lookup guidance %q includes queued filename or archived slug %q", message, unwanted)
		}
	}
}

func emptyPendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Fatalf("empty queue error = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	message := strings.ToLower(err.Error())
	if !strings.Contains(message, selector) {
		t.Errorf("empty queue error %q lacks rejected selector", message)
	}
	if !strings.Contains(message, "no pending") && !strings.Contains(message, "no valid pending") && !strings.Contains(message, "none") && !strings.Contains(message, "empty") && !strings.Contains(message, "[]") {
		t.Errorf("empty queue error %q lacks empty-set guidance", message)
	}
}

func Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"Project", "missing", func(selector string) error { _, err := project.Project(root, selector); return err }},
		{"Stamp", "0099-absent.md", func(selector string) error { return project.Stamp(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.Project(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
}

func Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"StaleSpecs", "missing", func(selector string) error { _, err := project.StaleSpecs(root, selector); return err }},
		{"StampMissing", "changes/0099-absent.md", func(selector string) error { return project.StampMissing(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.StaleSpecs(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
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

func Test_PRJ_069_RegularFileAtChangesReportsQueueCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	const queuePath = "changes"
	put(t, root, queuePath, "not a directory")

	_, cause := project.Queue(root)
	if cause == nil || !strings.Contains(cause.Error(), queuePath) {
		t.Fatalf("Queue error = %v; want an error identifying %q", cause, queuePath)
	}
	diagnostics := project.Check(root)
	for _, diagnostic := range diagnostics {
		if !diagnostic.Warning && diagnostic.Code == "C03" && diagnostic.Path == queuePath && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check diagnostics = %#v; want a C03 error at %q for Queue cause %q", diagnostics, queuePath, cause)
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

 succeeded in 634ms:
1:// Protected, project-owned Go hook. No third-party dependencies.
20:func main() {
59:## Modules and the zone map
78:Any zone field left out takes its value from `module_defaults` in `fireweed.yaml`, so a module using the recommended layout needs only `name`, `id_prefix` and `depends_on`. A Go module might instead declare `interface: [auth.go]`, `evals: ["*_eval_test.go"]` and `disposable: [internal/**]`.
103:## Specs
108:## Specs
150:bases:
202:**Body sections.** Motivation is required. Interface describes interface changes in prose; the evals phase turns them into code, which keeps change files independent of language. Added and Modified contain complete spec items; Removed contains headings only. Decisions is filled in by the worker. Questions is written by an agent that blocks the change because the spec is ambiguous. Removed decisions quotes module Decisions to drop. Provenance is written by fireweed at each step. Rejected alternatives and Notes are free prose.
322:**Provenance.** Each `done` appends one line to the change file's `## Provenance` section: the step, the date, the configured agent command (including its model) or `manual` when a person did the work, the fireweed version, the hash of the playbook used, and result counts such as tests added, failing before and passing after. Logs stay local in `.fireweed/runs/`; this record is committed and kept in the archive with the change.
339:Drills run only on demand, because scheduling would need an external trigger. `fireweed status` shows, per module, how many changes have been applied since its last drill, counting the module's archived changes whose implementation Provenance date is later than the drill report's date, so drift is visible without scheduling anything. The count uses dates rather than commit ancestry so it survives squash and rebase merges; the report's base commit is recorded for information only.
356:| `boundaries` | Runs across the project and prints one violation per line, prefixed with the module name: imports of modules outside `depends_on`, and evals that use anything beyond the interface | Yes, for modules with disposable globs (below) |
417:fireweed lets the user pick any of these. The per-step commits matter until review: they are what proves the implementation step left the evals alone, and `fireweed ci` on the PR branch checks them while they still exist. After a merge, fireweed reads only file contents (change files, the archive, Provenance and `bases` hashes), never a recorded commit SHA, so squash and rebase lose nothing it needs later.
478:module_defaults:
563:| Drill drift counted by commit ancestry | Replaced | Squash and rebase merges orphan the recorded base commit; Provenance dates survive any merge strategy |

 succeeded in 1079ms:
ls: changes: No such file or directory
.fireweed:
total 24
drwxr-xr-x@  8 joel  staff   256 Nov 10 08:44 .
drwx------@ 17 joel  staff   544 Nov 10 08:44 ..
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:44 bin
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:44 canaries
-rw-r--r--@  1 joel  staff  3993 Nov 10 08:44 hooks.go
-rw-r--r--@  1 joel  staff  1502 Nov 10 08:44 recipe.md
-rw-r--r--@  1 joel  staff   470 Nov 10 08:44 setup.sh
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:44 templates

.fireweed/bin:
total 13512
drwxr-xr-x@ 3 joel  staff       96 Nov 10 08:44 .
drwxr-xr-x@ 8 joel  staff      256 Nov 10 08:44 ..
-rwxr-xr-x@ 1 joel  staff  6914850 Nov 10 08:44 fireweed
<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step.
<!-- fireweed:end -->

## Rebuild controller

Drive the workflow with the controller `.fireweed/bin/fireweed` (see SYSTEM.md), not `bin/fireweed`, which is the product under test. Put `.fireweed/bin` first on worker PATH and use `GOCACHE=/tmp/fireweed-rebuild-cache`. Never switch controllers while a run or claim is active.

Use `fireweed verify` to check your work. Never commit during a worker step; fireweed owns commits. Read SYSTEM.md (including its scope guardrails) and docs/rebuild/PROGRESS.md before work; docs/rebuild/DECISIONS.md retains design clarifications. Do not consult `rebuild-v1` or deleted code as an implementation oracle.
b024bda reconstruction inputs
packages/project/SPEC.md
packages/project/api.go
packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/selector_guidance_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go
packages/project/model/model.go
# Author change
Read the design, SYSTEM.md and the module spec before writing a change. Give its Motivation a user-visible need or a clause of the design or SYSTEM.md; a probe showing a theoretical failure is insufficient. State each observable behavior at the module interface and keep the change focused. If a requirement costs more than it is worth, narrow or remove the spec and use already_satisfied for behavior the current implementation already meets. Do not commit during a worker step.
packages/contracts/model/model.go:4:type Document struct {
packages/contracts/model/model.go:10:type Item struct {
packages/contracts/model/model.go:17:type CodedError struct {
packages/contracts/model/model.go:23:func (e CodedError) Error() string {
packages/contracts/model/model.go:30:func (e CodedError) Unwrap() error { return e.Cause }
packages/contracts/model/model.go:33:type Zones struct {
packages/gitstore/model/model.go:5:type Repository struct {
packages/gitstore/model/model.go:12:type Change struct {
packages/gitstore/model/model.go:18:type Request struct {
packages/gitstore/model/model.go:28:type Result struct {
packages/assets/model/model.go:6:type Recipe struct {
packages/gitstore/api.go:10:type Repository = model.Repository
packages/gitstore/api.go:11:type Change = model.Change
packages/gitstore/api.go:12:type Request = model.Request
packages/gitstore/api.go:13:type Result = model.Result
packages/gitstore/api.go:15:func Inspect(dir string) (Repository, error)             { return local.Inspect(dir) }
packages/gitstore/api.go:16:func BaseRef(root, remote, base string) (string, error)  { return local.BaseRef(root, remote, base) }
packages/gitstore/api.go:17:func Changes(root, base string) ([]Change, error)        { return local.Changes(root, base) }
packages/gitstore/api.go:18:func ReadAt(root, revision, path string) ([]byte, error) { return local.ReadAt(root, revision, path) }
packages/gitstore/api.go:19:func Commit(root, subject string, paths []string) (string, error) {
packages/gitstore/api.go:22:func WriteBlob(root string, data []byte) (string, error) { return local.WriteBlob(root, data) }
packages/gitstore/api.go:23:func ReadBlob(root, oid string) ([]byte, error)          { return local.ReadBlob(root, oid) }
packages/gitstore/api.go:24:func Ref(root, name string) (string, bool, error)        { return local.Ref(root, name) }
packages/gitstore/api.go:25:func CASRef(root, name, old, new string) error           { return local.CASRef(root, name, old, new) }
packages/gitstore/api.go:28:func Run(ctx context.Context, request Request) (Result, error) { return local.Run(ctx, request) }
packages/contracts/api.go:8:type Document = model.Document
packages/contracts/api.go:9:type Item = model.Item
packages/contracts/api.go:10:type Zones = model.Zones
packages/contracts/api.go:11:type CodedError = model.CodedError
packages/contracts/api.go:14:func ErrorCode(err error) (code, subject string, ok bool) { return primitives.ErrorCode(err) }
packages/contracts/api.go:16:func ParseDocument(source string) (Document, error) { return primitives.ParseDocument(source) }
packages/contracts/api.go:17:func RenderDocument(doc Document) (string, error)   { return primitives.RenderDocument(doc) }
packages/contracts/api.go:20:func Section(markdown, name string) (string, bool) { return primitives.Section(markdown, name) }
packages/contracts/api.go:23:func ReplaceSection(markdown, name, body string) string {
packages/contracts/api.go:27:func ParseItems(section string) ([]Item, error) { return primitives.ParseItems(section) }
packages/contracts/api.go:28:func RenderItem(item Item) string               { return primitives.RenderItem(item) }
packages/contracts/api.go:29:func SpecHash(item Item) string                 { return primitives.SpecHash(item) }
packages/contracts/api.go:31:func MatchesID(id, testcaseName, classname string) bool {
packages/contracts/api.go:35:func MatchGlob(pattern, path string) bool { return primitives.MatchGlob(pattern, path) }
packages/contracts/api.go:38:func ResolveZone(path string, zones Zones) string { return primitives.ResolveZone(path, zones) }
packages/assets/api.go:7:type Recipe = model.Recipe
packages/assets/api.go:9:func Version() string { return content.Version() }
packages/assets/api.go:11:func Help() string { return content.Help() }
packages/assets/api.go:13:func Playbook(name string) (string, error) { return content.Playbook(name) }
packages/assets/api.go:15:func Explain(code string) (string, error) { return content.Explain(code) }
packages/assets/api.go:17:func BuiltinRecipe(name string) (Recipe, error) { return recipe.BuiltinRecipe(name) }

exec
/bin/zsh -lc 'cat packages/project/evals/static_codes_test.go packages/project/evals/selector_guidance_test.go packages/project/evals/project_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc "cat packages/project/evals/revision_snapshot_test.go; sed -n '200,430p' packages/project/evals/queue_test.go; cat packages/contracts/model/model.go; cat packages/contracts/internal/primitives/{coded_error,primitives}.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc "sed -n '59,206p;440,540p' fireweed-design-v1.md; cat packages/gitstore/model/model.go; rg -n 'func (Inspect|Run|ReadAt)|func git' packages/gitstore/internal/local/local.go; cat .fireweed/hooks.go; export PATH=\""'$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed context project; fireweed verify --help' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 567ms:
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

func Test_PRJ_069_RegularFileAtChangesReportsQueueCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	const queuePath = "changes"
	put(t, root, queuePath, "not a directory")

	_, cause := project.Queue(root)
	if cause == nil || !strings.Contains(cause.Error(), queuePath) {
		t.Fatalf("Queue error = %v; want an error identifying %q", cause, queuePath)
	}
	diagnostics := project.Check(root)
	for _, diagnostic := range diagnostics {
		if !diagnostic.Warning && diagnostic.Code == "C03" && diagnostic.Path == queuePath && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check diagnostics = %#v; want a C03 error at %q for Queue cause %q", diagnostics, queuePath, cause)
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
	"errors"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if err == nil {
		t.Fatal("unknown pending selector was accepted")
	}
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	var coded contracts.CodedError
	if !errors.As(err, &coded) || coded.Cause == nil {
		t.Errorf("lookup error = %v; want a retained U02 cause", err)
	}
	message := err.Error()
	for _, want := range []string{selector, "first", "second"} {
		if !strings.Contains(message, want) {
			t.Errorf("lookup error %q lacks rejected selector or pending slug %q", message, want)
		}
	}
	for _, unwanted := range []string{"0001-first.md", "0002-second.md", "archived"} {
		if strings.Contains(message, unwanted) {
			t.Errorf("lookup guidance %q includes queued filename or archived slug %q", message, unwanted)
		}
	}
}

func emptyPendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Fatalf("empty queue error = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	message := strings.ToLower(err.Error())
	if !strings.Contains(message, selector) {
		t.Errorf("empty queue error %q lacks rejected selector", message)
	}
	if !strings.Contains(message, "no pending") && !strings.Contains(message, "no valid pending") && !strings.Contains(message, "none") && !strings.Contains(message, "empty") && !strings.Contains(message, "[]") {
		t.Errorf("empty queue error %q lacks empty-set guidance", message)
	}
}

func Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"Project", "missing", func(selector string) error { _, err := project.Project(root, selector); return err }},
		{"Stamp", "0099-absent.md", func(selector string) error { return project.Stamp(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.Project(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
}

func Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"StaleSpecs", "missing", func(selector string) error { _, err := project.StaleSpecs(root, selector); return err }},
		{"StampMissing", "changes/0099-absent.md", func(selector string) error { return project.StampMissing(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.StaleSpecs(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
}
package evals_test

import (
	"errors"
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
			_, err := project.NewChange(root, tc.module, tc.slug)
			if err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			} else if tc.id == "PRJ-025" && (!strings.HasPrefix(err.Error(), "unknown module ") || !strings.Contains(err.Error(), tc.module)) {
				t.Errorf("unknown module error = %q; want prefix and rejected module %q", err, tc.module)
			}
		})
	}
}

func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
	root := fixture(t)
	const queuedPath = "changes/0001-blocked.md"
	if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(queuedPath)), 0755); err != nil {
		t.Fatal(err)
	}
	path, err := project.NewChange(root, "alpha", "blocked")
	if err == nil {
		t.Fatal("NewChange accepted a queued path occupied by a directory")
	}
	if path != "" {
		t.Errorf("rejected NewChange path = %q; want empty", path)
	}
	if !strings.HasPrefix(err.Error(), "create "+queuedPath+":") {
		t.Errorf("NewChange error = %q; want repository-relative create prefix", err)
	}
	var pathErr *os.PathError
	if !errors.As(err, &pathErr) || pathErr.Path != filepath.Join(root, filepath.FromSlash(queuedPath)) || pathErr.Err == nil {
		t.Errorf("NewChange error = %v; want retained filesystem cause for %s", err, queuedPath)
	}
}

 succeeded in 572ms:
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
			if tc.name == "missing_base" {
				for _, variant := range []struct{ name, front string }{
					{"null", "bases:\n  ALP-001: null\n"},
					{"empty_string", "bases:\n  ALP-001: \"\"\n"},
				} {
					for _, edit := range []string{"Modified", "Removed"} {
						modified, removed := "### ALP-001\nAlpha MUST change.\n", ""
						if edit == "Removed" {
							modified, removed = "", "### ALP-001\n"
						}
						queued(t, root, "0001-check.md", variant.front, "Needed.\n", "", modified, removed)
						found := false
						for _, diagnostic := range project.Check(root) {
							if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "missing base for ALP-001") {
								found = true
							}
						}
						if !found {
							t.Errorf("Check did not report C08 missing base for %s with %s entry", edit, variant.name)
						}
					}
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

func MatchGlob(pattern, path string) bool {
	parts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")
	var match func(int, int) bool
	match = func(pi, si int) bool {
		if pi == len(parts) {
			return si == len(pathParts)
		}
		if parts[pi] == "**" {
			return match(pi+1, si) || si < len(pathParts) && match(pi, si+1)
		}
		return si < len(pathParts) && matchComponent(parts[pi], pathParts[si]) && match(pi+1, si+1)
	}
	return match(0, 0)
}

func matchComponent(pattern, name string) bool {
	pi, ni := 0, 0
	star, resume := -1, 0
	for ni < len(name) {
		if pi < len(pattern) && pattern[pi] == name[ni] && pattern[pi] != '*' {
			pi++
			ni++
		} else if pi < len(pattern) && pattern[pi] == '*' {
			star = pi
			pi++
			resume = ni
		} else if star >= 0 {
			pi = star + 1
			resume++
			ni = resume
		} else {
			return false
		}
	}
	for pi < len(pattern) && pattern[pi] == '*' {
		pi++
	}
	return pi == len(pattern)
}

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
}

 exited 2 in 625ms:
## Modules and the zone map

A module is a directory containing a `SPEC.md`, and its frontmatter declares a zone map: which files belong to which zone. The zone map, not the directory layout, is the contract fireweed enforces.

Internally, fireweed resolves every path in the repository to a module (or none) and a zone. Almost every check works on that mapping plus a git diff, which is why the core needs no language knowledge. fireweed looks for `SPEC.md` files only under the `modules` globs in `fireweed.yaml`, and `check` warns about any `SPEC.md` outside them.

```yaml
---
name: auth
id_prefix: AUTH
depends_on: [clock]
interface: [src/index.ts]        # globs, relative to the module directory
evals: [evals/**]
disposable: [src/**]
preserve: [src/migrations/**]
append_only: [src/migrations/**]
---
```

Any zone field left out takes its value from `module_defaults` in `fireweed.yaml`, so a module using the recommended layout needs only `name`, `id_prefix` and `depends_on`. A Go module might instead declare `interface: [auth.go]`, `evals: ["*_eval_test.go"]` and `disposable: [internal/**]`.

| Zone | Default | Evals phase | Impl phase | Regeneration |
| --- | --- | --- | --- | --- |
| Interface | `src/index.ts` (per recipe) | May edit | Read only | Kept |
| Evals | `evals/**` | May edit | Read only | Kept |
| Disposable | `src/**` | New stub files only | May edit | Deleted, rebuilt |
| Preserved | none | Read only | May edit | Kept |
| Append-only | none | Read only | May add, never edit or delete | Kept |
| SPEC.md | `SPEC.md` | Read only | Read only | Kept |
| Other module files (manifests) | everything else in the module | May edit | May edit | Kept |
| Protected | `fireweed.yaml`, `SYSTEM.md`, `.fireweed/**`, every `AGENTS.md` and `CLAUDE.md`, configured globs | Read only | Read only | Kept |

When a file matches several zones, the strictest wins, in this order: protected, `SPEC.md`, evals, interface, append-only, preserved, disposable, other module files. So an interface file is read-only in the implementation step even though it is also preserved, and a preserved file is never deleted on regeneration even when a disposable glob matches it.

**Invariants `fireweed check` enforces:**

- Every path belongs to at most one module.
- Interface files and append-only files are always treated as preserved, whether or not `preserve` lists them.
- The evals zone and the disposable zone do not overlap.
- `depends_on` names existing modules, never the module itself, and has no cycles.
- `id_prefix` is unique across modules.

Files outside every module belong to humans. Workers may not touch them, except paths listed under `worker_writable` in `fireweed.yaml`, such as a workspace lockfile that changes when a module adds a dependency. Recipes set this list.

## Specs

Specs stay in Markdown: each item is a `### PREFIX-NNN` heading under a single `## Specs` section, with a body that states one observable behavior using MUST, MUST NOT or SHOULD. Only fireweed edits the `## Specs` section, by applying changes.

```markdown
## Specs

### AUTH-014
A login MUST be rejected for 15 minutes after 5 consecutive failed attempts for the same account.

### AUTH-020 [review]
Error messages MUST NOT reveal whether an account exists.

### AUTH-021 [quarantine: flaky clock in CI, see change fake-clock]
A session MUST expire 24 hours after its last use.
```

**Tags describe two separate things:**

- **How a spec is verified.** By default, by evals. `[review]` marks a spec that no deterministic test can check; it needs no evals. The PR body lists every review-tagged spec a change adds or modifies, and a human judges it during review.
- **Whether it is currently enforced.** `[quarantine: <reason>]` suspends enforcement and requires a reason. Every quarantined spec is listed by `fireweed status` and in every PR body until the tag is removed.

**IDs.** IDs are assigned by fireweed when a change starts, never chosen by authors (see Change files). A removed ID is never reused; fireweed tracks removed IDs through the archive.

**Spec shape.** Each item states one observable behavior at the module's interface. `check` warns (C17) when an item has more than `spec_limits.max_clauses` normative keywords (MUST, MUST NOT, SHOULD; default 3) or more than `spec_limits.max_words` words (default 80), because every extra clause is something an eval must assert and an audit can dispute. Properties of how code is written, such as purity, avoiding aliased slices, or which system calls run, are not observable behavior and are not specs.

**Matching IDs in test names.** A test covers a spec when the spec's ID appears in the JUnit testcase name or classname. Matching ignores case, treats `-` and `_` as equal, and allows the separator to be missing, so `AUTH-014`, `AUTH_014`, `AUTH014` and Python's `test_auth_014` all match. The ID must not be preceded by a letter or digit or followed by a digit, so `AUTH0145` and `OAUTH014` do not match. This lets languages such as Go, whose function names cannot contain a hyphen, follow the convention.

**System-wide behavior lives in composition-root modules.** There is no system pseudo-module. A root module such as `app` wires other modules together; its interface is the product's external surface (an HTTP API, a CLI), its `depends_on` lists the modules it composes, and its evals are end-to-end tests. It is ordinary disposable code that follows every rule. A product with several entry points has several root modules.

**SYSTEM.md is a plain document.** It holds the product description, project decisions (language, frameworks, deployment) and conventions. It has no spec IDs, workers cannot edit it, and humans edit it directly. Cross-cutting rules such as "never log personal data" go here as conventions, checked by review rather than by evals. A convention is not repeated in specs: `check` warns (C18) when the same normalized sentence appears in three or more spec items across the project, because that usually means a convention has been copied into specs, where every copy gains evals and audit findings. The audit does not report findings against conventions.

**Decisions.** A worker records any behavioral choice the spec did not cover in the change file's `## Decisions` section. When the change is applied, fireweed appends those decisions to a `## Decisions` section of the module's `SPEC.md`, which only fireweed edits, tagging each with the change's slug. Regeneration reads them, so a rebuilt module makes the same choices. A later change can drop decisions that a spec now covers or that no longer hold by quoting them in a `## Removed decisions` section, and drill reports flag decisions the rebuilt code contradicts.

Because every later worker and every regeneration reads them, decisions stay short: each is one or two sentences about product behavior. Notes about tests, verification runs, recovery or process belong in the change's Notes, which are archived with the change but never copied into `SPEC.md`. `check` warns (C19) when a module's Decisions exceed `spec_limits.max_decision_words` (default 300), which is the signal to drop stale ones through Removed decisions or turn lasting ones into specs.

**Spec history.** `fireweed show <module> <ID> --history` lists the archived changes that added, modified or retagged a spec, newest first, with each change's Motivation, Decisions and provenance. It answers why a spec exists without digging through git history.

## Change files

A change file describes how one module's specs should change. The queue is the set of files in `changes/`, run in order of number and then slug, and merging a change file into the base branch is the developer's approval.

```markdown
---
module: auth
set:
  append_only: [src/migrations/**]
bases:
  AUTH-003: 9f2c1a07
---

## Motivation
Accounts are being brute-forced; lock them after repeated failures.

## Interface
Add a `lockedUntil` field to the login error response.

## Added

### new:lockout
A login MUST be rejected for 15 minutes after 5 consecutive failed attempts for the same account.

## Modified

### AUTH-003
A failed login MUST return error code `invalid_credentials` and, when locked, `lockedUntil`.

## Removed

## Decisions

## Rejected alternatives
IP-based throttling: shared IPs would lock out whole offices.
```

**Ordering and numbering.** Files are named `NNNN-slug.md`. Numbers only order the queue and need not be unique; changes that share a number run in slug order. Two changes written in parallel without knowing about each other have no meaningful order anyway. Nothing refers to a change by its number.

**New spec IDs are assigned late.** New specs use a placeholder handle such as `new:lockout`. When the change starts, fireweed replaces each handle with the next free ID for the module, for example `AUTH-014`. Implementation runs one change at a time in base-branch order, so two authors working in parallel can never claim the same ID. Handles are `new:<slug>`, with the slug limited to lowercase letters, digits and hyphens and unique within the change. Other specs and the Interface section refer to a handle by the same string, and `start` rewrites every occurrence in the change file. A change may not use another change's handle: `check` reports an error, and the authoring playbook combines such changes into one. Evals never see handles, because they are written after `start`.

**Stale changes are caught.** `bases` records a short hash of each spec text the change modifies or removes, as it stood in the projected spec when the change was written. If another change has since altered that spec, `fireweed check` reports the change as stale and it must be rewritten. This prevents one change silently overwriting another's edit. Each hash is the first 8 hex characters of a SHA-256 over the spec item, tag included, with line endings and trailing whitespace normalized. Including the tag means that adding `[quarantine]` counts as an edit. A Modified or Removed spec with no `bases` entry is a `check` error, and the authoring playbook runs `stamp`.

**Fixing a stale change.** `show --stale <change>` prints, for each stale spec, a diff from the text the change was written against to the current projected text, so the author sees what moved underneath it. fireweed finds the old text by projecting the spec at the commit that last changed the change file's `bases` and checking that its hash matches; when no such commit exists it shows the current text and the recorded hash. `check --fix` stamps missing `bases` entries but never overwrites a stale one, because that would silently discard the conflict.

**Frontmatter fields:**

| Field | Set by | Meaning |
| --- | --- | --- |
| `module` | Author | The one module this change targets |
| `set` | Author | New values for module frontmatter: `depends_on`, `interface`, `evals`, `disposable`, `preserve`, `append_only`, `boundaries_opt_out` and per-module hook overrides. `name` and `id_prefix` cannot be changed |
| `bases` | fireweed, when the author runs `fireweed stamp` | Hashes of the spec versions this change was written against |
| `fixes` | Author | Spec IDs whose behavior is already specified correctly but whose evals missed a bug. The evals phase must add a failing test reproducing it |
| `already_satisfied` | Author | New or modified spec IDs whose evals may pass before implementation, such as clarifications of existing behavior or narrowed specs |
| `eval_fix` | Author | Spec IDs whose existing evals are wrong (for example a test helper rejects correct behavior) while the spec is right. The evals step corrects them and they must pass (see Kinds of change) |
| `regenerate` | Author | `true` rebuilds the module from scratch (see Regeneration) |
| `run` | fireweed (on start) | Unique ID of the run implementing this change, which also locates its logs |
| `phase` | fireweed | The next step to do: evals, audit or impl, and applied once archived |
| `audit_rounds` | fireweed | How many times the audit has sent this change back to evals |
| `blocked` | fireweed or agent | Why the change is stuck. Removing it unblocks the change |

**Body sections.** Motivation is required. Interface describes interface changes in prose; the evals phase turns them into code, which keeps change files independent of language. Added and Modified contain complete spec items; Removed contains headings only. Decisions is filled in by the worker. Questions is written by an agent that blocks the change because the spec is ambiguous. Removed decisions quotes module Decisions to drop. Provenance is written by fireweed at each step. Rejected alternatives and Notes are free prose.

**Kinds of change**, derived from the content:

- **Spec change:** has Added, Modified or Removed specs. Runs evals, audit and implementation.
| Checks | `ci [<range>]` | Every deterministic check in one command, for any CI system |
| Checks | `pr-body <range>` | A pull request description for the range: changes with their evals, interface and stub edits, preserved-file edits, review-tagged specs, audit findings and rounds, Decisions, quarantines, boundary opt-outs, blocked changes, budget warnings and binary or large files added to disposable zones |
| Lifecycle | `start`, `verify [--quick]`, `done` | Drive one change through its steps |
| Lifecycle | `next`, `run [--max N] [--for 6h]` | Run the agent through the queue |
| Lifecycle | `release <change>`, `commits <change>`, `abandon <change>` | Drop a local claim; list a change's commits on the current branch; revert and remove an unapplied change |
| Regeneration | `drill <module>` | Rebuild a module in a scratch worktree and report gaps |
| Maintenance | `gc`, `migrate`, `agents-md --write` | Prune refs, upgrade file formats, refresh the `AGENTS.md` block |

**What makes it easy for agents:**

- `--json` on every read command, with a stable schema.
- Stable diagnostic codes (for example C10, P2-01, H03), each with `fireweed explain <code>`.
- No prompts, ever. Missing input is a usage error that lists the valid options.
- Distinct exit codes: 0 for success or an empty queue, 1 for failed checks or any change blocked during the run, 2 for usage or configuration errors, 3 for a clean early stop on a time or usage limit.
- `context` replaces a dozen file reads with one call, and stays small as modules grow: files other than the change, spec, interface, `SYSTEM.md` and playbook are listed by path and size, and the agent reads the ones it needs.
- The same embedded playbook text drives unattended steps and interactive authoring, so every agent gets identical instructions.

## Distribution and init

fireweed ships as one Go binary that never touches the network, pinned per project, and set up by an agent following a recipe.

**Packaging.** goreleaser builds for GitHub releases and a Homebrew tap, alongside `go install` and an install script that takes an explicit version. Package managers handle updates; there is no self-update.

**No network access.** The binary reads and writes files and runs git, the project's hooks and the agent CLI. It never downloads recipes, checks for updates or sends telemetry. The agent CLIs do their own networking. Network features may come later but none are planned.

**Versioning.**

- `fireweed.yaml` pins a range of binary versions, such as `fireweed: ">=0.4, <0.5"`. The binary refuses to run outside it, so a laptop, CI and an agent's sandbox can't disagree about the rules.
- File formats have their own `version`. A binary reads the current and the previous format, and `fireweed migrate` rewrites files forward.

**Configuration.**

```yaml
fireweed: ">=0.4, <0.5"
version: 1
base: main
remote: origin
modules: ["packages/*"]
module_defaults:
  interface: [src/index.ts]
  evals: [evals/**]
  disposable: [src/**]
hooks:
  eval: "pnpm --filter ./{module_dir} test -- --reporter=junit --outputFile={report}"
  boundaries: "pnpm depcruise --config .fireweed/depcruise.cjs packages"
  setup: "pnpm install --frozen-lockfile"
protected: [".github/**"]
worker_writable: ["pnpm-lock.yaml"]
agent:
  command: ["claude", "-p", "{prompt}", "--permission-mode", "acceptEdits"]
  wrap: []          # optional sandbox prefix, e.g. a docker run command
audit:
  enabled: true
  command: ["codex", "exec", "{prompt}"]   # unset: uses agent.command, doctor warns
  max_rounds: 2          # audit-findings round trips before the change blocks
authoring:
  max_unapproved: 5      # applied changes not on base before start/run refuse more
spec_limits:
  max_clauses: 3         # normative keywords per spec item (C17 warning)
  max_words: 80
  max_decision_words: 300   # per module Decisions section (C19 warning)
budgets:
  verify_time: 15m       # B01 warning
  cases_per_spec: 15     # B02 warning
  eval_ratio: 4          # B03 warning: eval lines per disposable line
attempts: 2
timeout: 90m
verify_dependents: direct   # all | direct | none
limit_patterns: ["usage limit", "rate limit"]   # matched against the last 20 lines
limit_wait: 30m
limit_budget: 4h
```

The agent is any CLI that takes a prompt. Sandboxing is the user's choice: `wrap` is a command prefix, and fireweed's zone checks are the real enforcement either way.

**First run.** The expected flow starts with a developer telling their agent "set up fireweed for a TypeScript project":

1. The agent runs `fireweed init --recipe typescript-pnpm --agent claude`. This writes the config, `SYSTEM.md`, the `AGENTS.md` block, the recipe copy and an `app` composition-root module (`--no-app` skips it), then prints the recipe's instructions.
2. The agent completes the glue by following the recipe: workspace config, eval reporter, boundary checker.
3. The agent runs `fireweed doctor`, which plants canaries in throwaway modules and confirms every hook catches them. Failures name the hook that misbehaved.
4. With doctor passing, the developer fills in `SYSTEM.md` and authors the first change.

**Playbooks** for authoring changes, writing evals, auditing, implementing and regenerating are embedded in the binary, because they describe the protocol the binary verifies and must match it. Two rules in them exist to keep agents from escalating. The authoring playbook requires each change's Motivation to name a user-visible need or a clause of the design or `SYSTEM.md`; a probe showing that some failure is theoretically possible is not enough on its own. It also tells authors to narrow or remove specs, with `already_satisfied`, when a requirement turns out to cost more than it is worth. The audit playbook limits findings to spec clauses no test asserts, as described in Audit step. A project can override one in `.fireweed/playbooks/`; each override records the hash of the embedded playbook it was based on, and doctor warns when an upgrade changes that playbook underneath it.

**AGENTS.md.** `init` writes a short block between markers telling any agent that the project uses fireweed and pointing it at `fireweed playbook author-change` and `fireweed context`. `fireweed agents-md --write` regenerates only that block, and doctor warns when it is out of date.

**CI.** `fireweed ci` runs every deterministic check: `check`, the history checks on the range, coverage for every module, the boundary hooks, and the doctor canaries (skip them with `--no-doctor`), but not doctor's agent check. Adding it to any CI system takes one line, so fireweed needs no GitHub Action or other CI integration.

## Room for existing codebases

Adopting fireweed in an existing codebase is out of scope for now, and the design avoids choices that would rule it out later.

- **Disposability is opt-in per glob.** An existing module would start with `disposable: []`, keeping all its code, and grow the list as evals cover more.
- **Boundary checks follow disposability.** Modules with nothing disposable only get warnings, and a baseline file holds existing violations so only new ones fail.
- **Spec states can grow.** Separating verification from enforcement in tags leaves room for a "described, not yet verified" state without new grammar.
- **Module ownership can widen.** Every check goes through a single path-to-zone mapping, so an `owns:` glob list (a module claiming files spread across directories) or a `root:` field (a spec stored away from its code) can be added without redesign.
- **A spec-extraction playbook** that writes specs from an existing module would be the main addition.

## Alternatives considered

| Option | Decision | Why |
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
23:func Run(ctx context.Context, request model.Request) (model.Result, error) {
72:func git(root string, input []byte, env map[string]string, args ...string) ([]byte, error) {
86:func gitText(root string, args ...string) (string, error) {
99:func Inspect(dir string) (model.Repository, error) {
206:func ReadAt(root, revision, path string) ([]byte, error) {
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
U02: change "project" not found; no valid change slugs. Use a change held by this worktree and retry context.
U01: Unknown option "--help". Use help to see available commands and options.

2026-11-10T16:44:50.992333Z  WARN codex_sandboxing::violation: recorded sandbox violation: resource=filesystem backend=seatbelt reason=policy_denied path=unknown
exec
/bin/zsh -lc "sed -n '207,390p' packages/project/evals/project_test.go; sed -n '195,400p' packages/project/evals/queue_test.go; cat packages/project/evals/revision_snapshot_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc "cat packages/contracts/model/model.go; sed -n '1,310p' packages/contracts/internal/primitives/primitives.go; sed -n '1,125p;195,230p' packages/gitstore/internal/local/local.go; rg -n 'normative|outside|boundaries_opt_out' fireweed-design-v1.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 517ms:
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
			_, err := project.NewChange(root, tc.module, tc.slug)
			if err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			} else if tc.id == "PRJ-025" && (!strings.HasPrefix(err.Error(), "unknown module ") || !strings.Contains(err.Error(), tc.module)) {
				t.Errorf("unknown module error = %q; want prefix and rejected module %q", err, tc.module)
			}
		})
	}
}

func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
	root := fixture(t)
	const queuedPath = "changes/0001-blocked.md"
	if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(queuedPath)), 0755); err != nil {
		t.Fatal(err)
	}
	path, err := project.NewChange(root, "alpha", "blocked")
	if err == nil {
		t.Fatal("NewChange accepted a queued path occupied by a directory")
	}
	if path != "" {
		t.Errorf("rejected NewChange path = %q; want empty", path)
	}
	if !strings.HasPrefix(err.Error(), "create "+queuedPath+":") {
		t.Errorf("NewChange error = %q; want repository-relative create prefix", err)
	}
	var pathErr *os.PathError
	if !errors.As(err, &pathErr) || pathErr.Path != filepath.Join(root, filepath.FromSlash(queuedPath)) || pathErr.Err == nil {
		t.Errorf("NewChange error = %v; want retained filesystem cause for %s", err, queuedPath)
	}
}
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
			if tc.name == "missing_base" {
				for _, variant := range []struct{ name, front string }{
					{"null", "bases:\n  ALP-001: null\n"},
					{"empty_string", "bases:\n  ALP-001: \"\"\n"},
				} {
					for _, edit := range []string{"Modified", "Removed"} {
						modified, removed := "### ALP-001\nAlpha MUST change.\n", ""
						if edit == "Removed" {
							modified, removed = "", "### ALP-001\n"
						}
						queued(t, root, "0001-check.md", variant.front, "Needed.\n", "", modified, removed)
						found := false
						for _, diagnostic := range project.Check(root) {
							if !diagnostic.Warning && diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "missing base for ALP-001") {
								found = true
							}
						}
						if !found {
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

 succeeded in 520ms:
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

func MatchGlob(pattern, path string) bool {
	parts := strings.Split(pattern, "/")
	pathParts := strings.Split(path, "/")
	var match func(int, int) bool
	match = func(pi, si int) bool {
		if pi == len(parts) {
			return si == len(pathParts)
		}
		if parts[pi] == "**" {
			return match(pi+1, si) || si < len(pathParts) && match(pi, si+1)
		}
		return si < len(pathParts) && matchComponent(parts[pi], pathParts[si]) && match(pi+1, si+1)
	}
	return match(0, 0)
}

func matchComponent(pattern, name string) bool {
	pi, ni := 0, 0
	star, resume := -1, 0
	for ni < len(name) {
		if pi < len(pattern) && pattern[pi] == name[ni] && pattern[pi] != '*' {
			pi++
			ni++
		} else if pi < len(pattern) && pattern[pi] == '*' {
			star = pi
			pi++
			resume = ni
		} else if star >= 0 {
			pi = star + 1
			resume++
			ni = resume
		} else {
			return false
		}
	}
	for pi < len(pattern) && pattern[pi] == '*' {
		pi++
	}
	return pi == len(pattern)
}

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
	for path := range byPath {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	changes := make([]model.Change, 0, len(paths))
	for _, path := range paths {
		changes = append(changes, byPath[path])
	}
	return changes, nil
}

func ReadAt(root, revision, path string) ([]byte, error) {
	return git(root, nil, nil, "show", revision+":"+path)
}

func Commit(root, subject string, paths []string) (string, error) {
	if len(paths) == 0 {
		return "", errors.New("no paths selected")
	}
	old, err := gitText(root, "rev-parse", "HEAD")
	if err != nil {
		return "", err
	}
	file, err := os.CreateTemp("", "fireweed-index-*")
	if err != nil {
		return "", err
	}
	index := file.Name()
	_ = file.Close()
	defer os.Remove(index)
	env := map[string]string{"GIT_INDEX_FILE": index}
	if _, err := git(root, nil, env, "read-tree", "HEAD"); err != nil {
		return "", err
	}
	args := append([]string{"add", "-A", "--"}, paths...)
	if _, err := git(root, nil, env, args...); err != nil {
63:Internally, fireweed resolves every path in the repository to a module (or none) and a zone. Almost every check works on that mapping plus a git diff, which is why the core needs no language knowledge. fireweed looks for `SPEC.md` files only under the `modules` globs in `fireweed.yaml`, and `check` warns about any `SPEC.md` outside them.
101:Files outside every module belong to humans. Workers may not touch them, except paths listed under `worker_writable` in `fireweed.yaml`, such as a workspace lockfile that changes when a module adds a dependency. Recipes set this list.
127:**Spec shape.** Each item states one observable behavior at the module's interface. `check` warns (C17) when an item has more than `spec_limits.max_clauses` normative keywords (MUST, MUST NOT, SHOULD; default 3) or more than `spec_limits.max_words` words (default 80), because every extra clause is something an eval must assert and an audit can dispute. Properties of how code is written, such as purity, avoiding aliased slices, or which system calls run, are not observable behavior and are not specs.
191:| `set` | Author | New values for module frontmatter: `depends_on`, `interface`, `evals`, `disposable`, `preserve`, `append_only`, `boundaries_opt_out` and per-module hook overrides. `name` and `id_prefix` cannot be changed |
356:| `boundaries` | Runs across the project and prints one violation per line, prefixed with the module name: imports of modules outside `depends_on`, and evals that use anything beyond the interface | Yes, for modules with disposable globs (below) |
371:- Opting out is explicit: `boundaries_opt_out: "<reason>"` in the module's frontmatter. `doctor`, `status` and PR bodies report every opt-out.
398:**Worktrees.** fireweed creates a worktree only when work must not touch the user's checkout: `drill` and `doctor`. They go under `$GIT_COMMON_DIR/fireweed/worktrees/`, outside the working tree, run the `setup` hook, and are removed afterwards. `start`, `next` and `run` work in the current checkout and require a clean tree; to keep working during a run, create a worktree and run fireweed there. The agent `wrap` command receives `{git_common_dir}`, so a sandbox wrapper can mount git metadata for a linked worktree.
467:- `fireweed.yaml` pins a range of binary versions, such as `fireweed: ">=0.4, <0.5"`. The binary refuses to run outside it, so a laptop, CI and an agent's sandbox can't disagree about the rules.
498:  max_clauses: 3         # normative keywords per spec item (C17 warning)

exec
/bin/zsh -lc "rg -n 'C0[1-9]|C12|C13' packages/assets/internal/content/content.go; sed -n '59,105p' fireweed-design-v1.md; cat packages/project/evals/spec_history_test.go | sed -n '1,115p'; git status --short; export PATH=\""'$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed verify; go env GOMODCACHE GOPROXY' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 601ms:
79:	"C01":   "The project configuration is invalid: fireweed.yaml is missing or has an invalid required setting. Edit the configuration file to correct the reported field, then run fireweed check again.",
80:	"C02":   "The module graph is invalid because of an empty module name, a duplicate module name, an empty module prefix, a duplicate module prefix, or a dependency that names an unknown module or creates a cycle. Correct module names and prefixes in SPEC.md so they are nonempty and unique. Correct the module dependency in SPEC.md so it names an existing module and the graph is acyclic.",
81:	"C03":   "A document or spec item is unreadable or malformed. Repair the document syntax or restore the missing file, then check it again.",
82:	"C04":   "The spec item is non-normative because its body lacks MUST or SHOULD wording. Rewrite the spec with a MUST or SHOULD requirement that states the expected behavior.",
83:	"C05":   "A spec ID uses a foreign module prefix or duplicates another ID, or a change has a duplicate edit identity. Change a foreign ID to use the owning module prefix, assign a unique unused ID for a duplicate spec ID, or remove or combine the repeated edit.",
84:	"C06":   "The spec has an invalid review or quarantine tag. Correct the tag to use a valid review or quarantine form, or remove it.",
85:	"C07":   "The change has a missing Motivation section or its body is empty, or the pending change module target is missing or invalid. Add a Motivation explaining the user-visible or design need, and set the pending change module to an existing module. An eval_fix change mixed with fixes or spec edits is also invalid. Separate eval-fix work from fixes or spec changes.",
86:	"C08":   "An edit base is missing or stale, or its target ID is absent from the preceding projection. Stamp or update the base against the preceding projection, and correct or remove an edit whose target is absent.",
87:	"C09":   "A change references an undeclared handle. Declare the handle in Added, or correct or remove the reference.",
90:	"C12":   "Module metadata or a zone definition is invalid or missing, or local recipe metadata is invalid. Correct the metadata field or define a valid zone in the module specification; repair invalid local recipe metadata in recipe.yaml.",
91:	"C13":   "A duplicate queue slug or an unparseable numeric queue prefix makes the queue identity invalid; tied numeric positions are allowed. Choose an unused slug with a valid numeric prefix.",
## Modules and the zone map

A module is a directory containing a `SPEC.md`, and its frontmatter declares a zone map: which files belong to which zone. The zone map, not the directory layout, is the contract fireweed enforces.

Internally, fireweed resolves every path in the repository to a module (or none) and a zone. Almost every check works on that mapping plus a git diff, which is why the core needs no language knowledge. fireweed looks for `SPEC.md` files only under the `modules` globs in `fireweed.yaml`, and `check` warns about any `SPEC.md` outside them.

```yaml
---
name: auth
id_prefix: AUTH
depends_on: [clock]
interface: [src/index.ts]        # globs, relative to the module directory
evals: [evals/**]
disposable: [src/**]
preserve: [src/migrations/**]
append_only: [src/migrations/**]
---
```

Any zone field left out takes its value from `module_defaults` in `fireweed.yaml`, so a module using the recommended layout needs only `name`, `id_prefix` and `depends_on`. A Go module might instead declare `interface: [auth.go]`, `evals: ["*_eval_test.go"]` and `disposable: [internal/**]`.

| Zone | Default | Evals phase | Impl phase | Regeneration |
| --- | --- | --- | --- | --- |
| Interface | `src/index.ts` (per recipe) | May edit | Read only | Kept |
| Evals | `evals/**` | May edit | Read only | Kept |
| Disposable | `src/**` | New stub files only | May edit | Deleted, rebuilt |
| Preserved | none | Read only | May edit | Kept |
| Append-only | none | Read only | May add, never edit or delete | Kept |
| SPEC.md | `SPEC.md` | Read only | Read only | Kept |
| Other module files (manifests) | everything else in the module | May edit | May edit | Kept |
| Protected | `fireweed.yaml`, `SYSTEM.md`, `.fireweed/**`, every `AGENTS.md` and `CLAUDE.md`, configured globs | Read only | Read only | Kept |

When a file matches several zones, the strictest wins, in this order: protected, `SPEC.md`, evals, interface, append-only, preserved, disposable, other module files. So an interface file is read-only in the implementation step even though it is also preserved, and a preserved file is never deleted on regeneration even when a disposable glob matches it.

**Invariants `fireweed check` enforces:**

- Every path belongs to at most one module.
- Interface files and append-only files are always treated as preserved, whether or not `preserve` lists them.
- The evals zone and the disposable zone do not overlap.
- `depends_on` names existing modules, never the module itself, and has no cycles.
- `id_prefix` is unique across modules.

Files outside every module belong to humans. Workers may not touch them, except paths listed under `worker_writable` in `fireweed.yaml`, such as a workspace lockfile that changes when a module adds a dependency. Recipes set this list.

## Specs

Specs stay in Markdown: each item is a `### PREFIX-NNN` heading under a single `## Specs` section, with a body that states one observable behavior using MUST, MUST NOT or SHOULD. Only fireweed edits the `## Specs` section, by applying changes.
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
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.
/Users/joel/go/pkg/mod
https://proxy.golang.org,direct

codex
The controller reports that this reconstruction workspace has no held change, so `context` and `verify` cannot run through the normal lifecycle. I’ll use the supplied eval hook to verify the rebuild.
exec
/bin/zsh -lc "mkdir -p packages/project/internal/{files,fixture,queue,history}
cat > packages/project/internal/files/source.go <<'EOF'
// Package files reads project documents from the worktree or a recorded Git tree.
package files

import (
	\"context\"
	\"fmt\"
	\"io/fs\"
	\"os\"
	\"path\"
	\"sort\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
)

type Source interface {
	ReadFile(string) ([]byte, error)
	ReadDir(string) ([]Entry, error)
}

type Entry struct {
	Name string
	Mode fs.FileMode
}

type Worktree struct{ Root string }

func (w Worktree) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(os.DirFS(w.Root), name)
}

func (w Worktree) ReadDir(name string) ([]Entry, error) {
	entries, err := fs.ReadDir(os.DirFS(w.Root), name)
	if err "'!= nil {
		return nil, err
	}
	result := make([]Entry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, Entry{entry.Name(), entry.Type()})
	}
	return result, nil
}

type Revision struct {
	Root string
	commit string
	dirs map[string][]Entry
	modes map[string]fs.FileMode
}

func Git(root string, args ...string) (string, error) {
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Dir: root,
		Args: append([]string{"git", "--no-optional-locks"}, args...),
		Env: map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
		Timeout: 30*time.Second,
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}

func OpenRevision(root, revision string) (*Revision, error) {
	output, err := Git(root, "rev-parse", "--verify", "--end-of-options", revision+"''^{commit}")
	if err != nil {
		return nil, err
	}
	commit := strings.TrimSpace(output)
	output, err = Git(root, "ls-tree", "--full-tree", "-r", "-t", "-z", commit)
	if err != nil {
		return nil, err
	}
	tree := &Revision{
		Root: root, commit: commit,
		dirs: map[string][]Entry{".": nil},
		modes: make(map[string]fs.FileMode),
	}
	for _, record := range strings.Split(strings.TrimSuffix(output, "'"\\x00\"), \"\\x00\") {
		if record == \"\" {
			continue
		}
		header, name, _ := strings.Cut(record, \"\\t\")
		mode, _, _ := strings.Cut(header, \" \")
		var kind fs.FileMode
		switch mode {
		case \"040000\":
			kind = fs.ModeDir
			tree.dirs[name] = nil
		case \"120000\":
			kind = fs.ModeSymlink
		case \"100644\", \"100755\":
		default:
			kind = fs.ModeIrregular
		}
		parent := path.Dir(name)
		tree.dirs[parent] = append(tree.dirs[parent], Entry{path.Base(name), kind})
		tree.modes[name] = kind
	}
	return tree, nil
}

func (r *Revision) ReadFile(name string) ([]byte, error) {
	mode, exists := r.modes[name]
	if "'!exists {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrNotExist}
	}
	if !mode.IsRegular() {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrInvalid}
	}
	return gitstore.ReadAt(r.Root, r.commit, name)
}

func (r *Revision) ReadDir(name string) ([]Entry, error) {
	entries, exists := r.dirs[name]
	if !exists {
		cause := fs.ErrNotExist
		if _, present := r.modes[name]; present {
			cause = fs.ErrInvalid
		}
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: cause}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })
	return entries, nil
}

type PathError struct {
	Path string
	Err error
}

func (e PathError) Error() string { return e.Path + ": " + e.Err.Error() }
func (e PathError) Unwrap() error { return e.Err }

func Code(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
}

func At(name, code string, err error) error {
	return PathError{Path: name, Err: Code(code, name, err)}
}

// Syntax preserves the parser'"'s classification and subject.
func Syntax(name string, err error) error {
	return PathError{Path: name, Err: err}
}

func ReadDocument(source Source, name string) (contracts.Document, error) {
	data, err := source.ReadFile(name)
	if err "'!= nil {
		return contracts.Document{}, At(name, "C03", err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return contracts.Document{}, Syntax(name, err)
	}
	return doc, nil
}

func Items(doc contracts.Document, section string) ([]contracts.Item, error) {
	body, _ := contracts.Section(doc.Body, section)
	return contracts.ParseItems(body)
}
EOF
cat > packages/project/internal/files/changes.go <<'"'EOF'
package files

import (
	\"errors\"
	\"fmt\"
	\"io/fs\"
	\"path\"
	\"regexp\"
	\"sort\"
	\"strconv\"
)

const Pending = \"changes\"
const Archive = \"changes/archive\"

var Slug = regexp.MustCompile("'`''^[a-z0-9-]+$`)
var changeName = regexp.MustCompile(`''^([0-9]{4,})-([a-z0-9-]+)'"\\.md"'$`)

type ChangeFile struct {
	Number int
	Slug string
	Path string
}

func Changes(source Source, dir string) ([]ChangeFile, error) {
	entries, err := source.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, At(dir, "C03", err)
	}
	var result []ChangeFile
	for _, entry := range entries {
		match := changeName.FindStringSubmatch(entry.Name)
		if match == nil || entry.Mode.IsDir() {
			continue
		}
		name := path.Join(dir, entry.Name)
		if !entry.Mode.IsRegular() {
			return nil, At(name, "C03", fmt.Errorf("change is not a regular file"))
		}
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, At(name, "C13", fmt.Errorf("invalid change number: %w", err))
		}
		result = append(result, ChangeFile{number, match[2], name})
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if a.Number != b.Number {
			return a.Number < b.Number
		}
		if a.Slug != b.Slug {
			return a.Slug < b.Slug
		}
		return a.Path < b.Path
	})
	return result, nil
}
EOF
cat > packages/project/internal/fixture/load.go <<'"'EOF'
package fixture

import (
	\"bytes\"
	\"fmt\"
	\"io\"
	\"path\"
	\"path/filepath\"
	\"slices\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
	\"gopkg.in/yaml.v3\"
)

const configPath = \"fireweed.yaml\"

func Load(root string) (model.Snapshot, error) {
	absolute, err := filepath.Abs(root)
	if err "'!= nil {
		return model.Snapshot{}, configError(err)
	}
	return LoadFrom(absolute, files.Worktree{Root: absolute})
}

func LoadAt(root, revision string) (model.Snapshot, error) {
	top, err := files.Git(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return model.Snapshot{}, err
	}
	canonical, err := filepath.EvalSymlinks(strings.TrimSuffix(top, "'"\\n\"))
	if err "'!= nil {
		return model.Snapshot{}, err
	}
	source, err := files.OpenRevision(canonical, revision)
	if err != nil {
		return model.Snapshot{}, err
	}
	return LoadFrom(canonical, source)
}

func configError(err error) error {
	return files.Code("C01", configPath, fmt.Errorf("C01: %w", files.PathError{Path: configPath, Err: err}))
}

func LoadFrom(root string, source files.Source) (model.Snapshot, error) {
	data, err := source.ReadFile(configPath)
	if err != nil {
		return model.Snapshot{}, configError(err)
	}
	var config map[string]any
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&config); err != nil {
		return model.Snapshot{}, configError(err)
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		if err == nil {
			err = fmt.Errorf("expected one YAML document")
		}
		return model.Snapshot{}, configError(err)
	}
	patterns, err := stringList(config["modules"])
	if err != nil {
		return model.Snapshot{}, configError(fmt.Errorf("modules: %w", err))
	}
	for i, pattern := range patterns {
		clean := path.Clean(pattern)
		if pattern == "" || path.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, "../") {
			return model.Snapshot{}, configError(fmt.Errorf("invalid module glob %q", pattern))
		}
		patterns[i] = clean
	}
	defaults, err := Defaults(config)
	if err != nil {
		return model.Snapshot{}, files.At(configPath, "C12", err)
	}
	dirs, err := moduleDirs(source, patterns)
	if err != nil {
		return model.Snapshot{}, configError(err)
	}
	snapshot := model.Snapshot{Root: root, Config: config}
	for _, dir := range dirs {
		name := path.Join(dir, "SPEC.md")
		doc, err := files.ReadDocument(source, name)
		if err != nil {
			return model.Snapshot{}, err
		}
		if _, err := files.Items(doc, "Specs"); err != nil {
			return model.Snapshot{}, files.Syntax(name, err)
		}
		module, err := DecodeModule(dir, doc, defaults)
		if err != nil {
			return model.Snapshot{}, files.At(name, "C12", err)
		}
		snapshot.Modules = append(snapshot.Modules, module)
	}
	return snapshot, nil
}

func moduleDirs(source files.Source, globs []string) ([]string, error) {
	// Walk only branches that can match a configured glob. A standalone **
	// can either match no directory or stay active while descending.
	matches := make(map[string]bool)
	for _, glob := range globs {
		parts := strings.Split(glob, "/")
		if glob == "." {
			parts = nil
		}
		visited := make(map[struct{dir string; index int}]bool)
		var walk func(string, int) error
		walk = func(dir string, index int) error {
			state := struct{dir string; index int}{dir, index}
			if visited[state] {
				return nil
			}
			visited[state] = true
			if index == len(parts) {
				matches[dir] = true
				return nil
			}
			part := parts[index]
			if part == "**" {
				if err := walk(dir, index+1); err != nil {
					return err
				}
			}
			children, err := source.ReadDir(dir)
			if err != nil {
				return files.PathError{Path: dir, Err: err}
			}
			for _, child := range children {
				if !child.Mode.IsDir() || child.Name == ".git" {
					continue
				}
				next := index + 1
				if part == "**" {
					next = index
				} else if !contracts.MatchGlob(part, child.Name) {
					continue
				}
				if err := walk(path.Join(dir, child.Name), next); err != nil {
					return err
				}
			}
			return nil
		}
		if err := walk(".", 0); err != nil {
			return nil, err
		}
	}
	dirs := make([]string, 0, len(matches))
	for dir := range matches {
		dirs = append(dirs, dir)
	}
	slices.Sort(dirs)
	return dirs, nil
}
EOF
cat > packages/project/internal/fixture/metadata.go <<'"'EOF'
package fixture

import (
	\"fmt\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func stringList(value any) ([]string, error) {
	list, ok := value.([]any)
	if "'!ok {
		return nil, fmt.Errorf("expected a list of strings")
	}
	result := make([]string, len(list))
	for i, value := range list {
		if result[i], ok = value.(string); !ok {
			return nil, fmt.Errorf("expected a string at list position %d", i+1)
		}
	}
	return result, nil
}

func Defaults(config map[string]any) (map[string]any, error) {
	raw, present := config["module_defaults"]
	if !present {
		return nil, nil
	}
	defaults, ok := raw.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("module_defaults must be a mapping")
	}
	if _, err := zones(defaults, nil); err != nil {
		return nil, fmt.Errorf("module_defaults: %w", err)
	}
	return defaults, nil
}

func zones(fields, defaults map[string]any) (contracts.Zones, error) {
	var result contracts.Zones
	for _, zone := range []struct {
		key string
		patterns *[]string
	}{
		{"protected", &result.Protected},
		{"evals", &result.Evals},
		{"interface", &result.Interface},
		{"append_only", &result.AppendOnly},
		{"preserve", &result.Preserve},
		{"disposable", &result.Disposable},
	} {
		value, present := fields[zone.key]
		if !present {
			value, present = defaults[zone.key]
		}
		if present {
			list, err := stringList(value)
			if err != nil {
				return result, fmt.Errorf("%s: %w", zone.key, err)
			}
			*zone.patterns = list
		}
	}
	return result, nil
}

func DecodeModule(dir string, doc contracts.Document, defaults map[string]any) (model.Module, error) {
	module := model.Module{Dir: dir, Document: doc}
	fields := doc.Frontmatter
	for _, field := range []struct {
		key string
		target *string
	}{{"name", &module.Name}, {"id_prefix", &module.IDPrefix}, {"boundaries_opt_out", new(string)}} {
		if value, present := fields[field.key]; present {
			var ok bool
			if *field.target, ok = value.(string); !ok {
				return module, fmt.Errorf("%s must be a string", field.key)
			}
		}
	}
	if hooks, present := fields["hooks"]; present {
		if _, ok := hooks.(map[string]any); !ok {
			return module, fmt.Errorf("hooks must be a mapping")
		}
	}
	var err error
	if deps, present := fields["depends_on"]; present {
		module.DependsOn, err = stringList(deps)
		if err != nil {
			return module, fmt.Errorf("depends_on: %w", err)
		}
	}
	module.Zones, err = zones(fields, defaults)
	return module, err
}

func ModuleIndex(modules []model.Module, name string) int {
	for i, module := range modules {
		if module.Name == name {
			return i
		}
	}
	return -1
}
EOF
cat > packages/project/internal/fixture/graph.go <<'"'EOF'
package fixture

import (
	\"fmt\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func ValidateGraph(modules []model.Module) error {
	names := make(map[string]model.Module, len(modules))
	prefixes := make(map[string]string, len(modules))
	for _, module := range modules {
		switch {
		case strings.TrimSpace(module.Name) == \"\":
			return fmt.Errorf(\"empty module name in %q\", module.Dir)
		case strings.TrimSpace(module.IDPrefix) == \"\":
			return fmt.Errorf(\"empty ID prefix for module %q\", module.Name)
		}
		if _, present := names[module.Name]; present {
			return fmt.Errorf(\"duplicate module name %q\", module.Name)
		}
		if other, present := prefixes[module.IDPrefix]; present {
			return fmt.Errorf(\"duplicate ID prefix %q for modules %q and %q\", module.IDPrefix, other, module.Name)
		}
		names[module.Name] = module
		prefixes[module.IDPrefix] = module.Name
	}
	for _, module := range modules {
		for _, dep := range module.DependsOn {
			if dep == module.Name {
				return fmt.Errorf(\"module %q has a self dependency\", module.Name)
			}
			if _, present := names[dep]; "'!present {
				return fmt.Errorf("module %q depends on unknown module %q", module.Name, dep)
			}
		}
	}
	active := make(map[string]bool)
	done := make(map[string]bool)
	var visit func(string, []string) error
	visit = func(name string, trail []string) error {
		if active[name] {
			return fmt.Errorf("module dependency cycle: %s", strings.Join(append(trail, name), " -> "))
		}
		if done[name] {
			return nil
		}
		active[name] = true
		for _, dep := range names[name].DependsOn {
			if err := visit(dep, append(trail, name)); err != nil {
				return err
			}
		}
		delete(active, name)
		done[name] = true
		return nil
	}
	for _, module := range modules {
		if err := visit(module.Name, nil); err != nil {
			return err
		}
	}
	return nil
}
EOF
cat > packages/project/internal/fixture/new.go <<'"'EOF'
package fixture

import (
	\"fmt\"
	\"math\"
	\"os\"
	\"path/filepath\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
)

const skeleton = \"\\n## Motivation\\n\\n## Interface\\n\\n## Added\\n\\n## Modified\\n\\n## Removed\\n\\n## Decisions\\n\"

func NewChange(root, module, slug string) (string, error) {
	snapshot, err := Load(root)
	if err "'!= nil {
		return "", err
	}
	if ModuleIndex(snapshot.Modules, module) < 0 {
		return "", fmt.Errorf("unknown module %q", module)
	}
	if !files.Slug.MatchString(slug) {
		return "", fmt.Errorf("invalid change slug %q: use lowercase letters, digits and hyphens", slug)
	}
	source := files.Worktree{Root: snapshot.Root}
	maximum := 0
	for _, dir := range []string{files.Pending, files.Archive} {
		changes, err := files.Changes(source, dir)
		if err != nil {
			if dir == files.Pending {
				return "", fmt.Errorf("create changes/%04d-%s.md: %w", maximum+1, slug, err)
			}
			return "", err
		}
		for _, change := range changes {
			if change.Slug == slug {
				return "", fmt.Errorf("change slug %q already exists at %s", slug, change.Path)
			}
			maximum = max(maximum, change.Number)
		}
	}
	if maximum == math.MaxInt {
		return "", fmt.Errorf("cannot number change %q after %d", slug, maximum)
	}
	relative := fmt.Sprintf("changes/%04d-%s.md", maximum+1, slug)
	name := filepath.Join(snapshot.Root, filepath.FromSlash(relative))
	content, err := contracts.RenderDocument(contracts.Document{
		Frontmatter: map[string]any{"module": module},
		Body: skeleton,
	})
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	_, writeErr := file.WriteString(content)
	closeErr := file.Close()
	if writeErr == nil {
		writeErr = closeErr
	}
	if writeErr != nil {
		_ = os.Remove(name)
		return "", fmt.Errorf("create %s: %w", relative, writeErr)
	}
	return relative, nil
}
EOF
gofmt -w packages/project/internal/{files,fixture}' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 503ms:
exec
/bin/zsh -lc "cat > packages/project/internal/queue/queue.go <<'EOF'
package queue

import (
	\"fmt\"
	\"path\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func Queue(root string) ([]model.Change, error) {
	_, changes, err := load(root)
	return changes, err
}

func load(root string) (model.Snapshot, []model.Change, error) {
	snapshot, err := fixture.Load(root)
	if err "'!= nil {
		return model.Snapshot{}, nil, err
	}
	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules)
	return snapshot, changes, err
}

func read(source files.Source, modules []model.Module) ([]model.Change, error) {
	pending, err := files.Changes(source, files.Pending)
	if err != nil {
		return nil, err
	}
	slugs := make(map[string]string)
	for _, entry := range pending {
		if other, present := slugs[entry.Slug]; present {
			return nil, files.At(entry.Path, "C13", fmt.Errorf("duplicate change slug %q at %s and %s", entry.Slug, other, entry.Path))
		}
		slugs[entry.Slug] = entry.Path
	}
	changes := make([]model.Change, 0, len(pending))
	for _, entry := range pending {
		doc, err := files.ReadDocument(source, entry.Path)
		if err != nil {
			return nil, err
		}
		module, _ := doc.Frontmatter["module"].(string)
		if module == "" || fixture.ModuleIndex(modules, module) < 0 {
			return nil, files.At(entry.Path, "C07", fmt.Errorf("invalid pending change module %v", doc.Frontmatter["module"]))
		}
		change := model.Change{Number: entry.Number, Slug: entry.Slug, Path: entry.Path, Module: module, Document: doc}
		for _, section := range []struct {
			name string
			items *[]contracts.Item
		}{{"Added", &change.Added}, {"Modified", &change.Modified}, {"Removed", &change.Removed}} {
			items, err := files.Items(doc, section.name)
			if err != nil {
				return nil, files.Syntax(entry.Path, err)
			}
			*section.items = items
		}
		seen := make(map[string]bool)
		for _, items := range [][]contracts.Item{change.Added, change.Modified, change.Removed} {
			for _, item := range items {
				if seen[item.ID] {
					return nil, files.At(entry.Path, "C05", fmt.Errorf("duplicate edit identity %s", item.ID))
				}
				seen[item.ID] = true
			}
		}
		changes = append(changes, change)
	}
	return changes, nil
}

func resolve(changes []model.Change, selector string) (int, error) {
	slugs := make([]string, 0, len(changes))
	for i, change := range changes {
		if selector == change.Slug || selector == path.Base(change.Path) || path.Clean(selector) == change.Path {
			return i, nil
		}
		slugs = append(slugs, change.Slug)
	}
	guidance := "no pending changes"
	if len(slugs) != 0 {
		guidance = "valid pending change slugs: " + strings.Join(slugs, ", ")
	}
	return -1, files.Code("U02", selector, fmt.Errorf("U02: change %q not found; %s", selector, guidance))
}
EOF
cat > packages/project/internal/queue/project.go <<'"'EOF'
package queue

import (
	\"fmt\"
	\"maps\"
	\"slices\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func Project(root, through string) (model.Snapshot, error) {
	snapshot, changes, err := load(root)
	if err "'!= nil {
		return model.Snapshot{}, err
	}
	if through != "" {
		index, err := resolve(changes, through)
		if err != nil {
			return model.Snapshot{}, err
		}
		changes = changes[:index+1]
	}
	projected, err := project(snapshot, changes)
	if err != nil {
		return model.Snapshot{}, err
	}
	return projected.snapshot, nil
}

type projection struct {
	snapshot model.Snapshot
	items [][]contracts.Item
	defaults map[string]any
}

func newProjection(snapshot model.Snapshot) (projection, error) {
	if err := fixture.ValidateGraph(snapshot.Modules); err != nil {
		return projection{}, files.Code("C02", "", err)
	}
	p := projection{snapshot: snapshot, items: make([][]contracts.Item, len(snapshot.Modules))}
	p.defaults, _ = snapshot.Config["module_defaults"].(map[string]any)
	for i, module := range snapshot.Modules {
		items, err := files.Items(module.Document, "Specs")
		if err != nil {
			return projection{}, files.Syntax(modulePath(module), err)
		}
		p.items[i] = items
	}
	return p, nil
}

func project(snapshot model.Snapshot, changes []model.Change) (projection, error) {
	p, err := newProjection(snapshot)
	if err != nil {
		return projection{}, err
	}
	for _, change := range changes {
		if err := p.apply(change); err != nil {
			return projection{}, err
		}
	}
	return p, nil
}

func modulePath(module model.Module) string {
	if module.Dir == "." {
		return "SPEC.md"
	}
	return module.Dir + "/SPEC.md"
}

// stableIndex excludes handles. They are rendered as written, but they never
// become available as targets of later edits.
func stableIndex(items []contracts.Item) map[string]int {
	index := make(map[string]int)
	for i, item := range items {
		if !isHandle(item.ID) {
			index[item.ID] = i
		}
	}
	return index
}

func isHandle(id string) bool { return strings.HasPrefix(id, "new:") }

func (p *projection) apply(change model.Change) error {
	i := fixture.ModuleIndex(p.snapshot.Modules, change.Module)
	original := p.snapshot.Modules[i]
	module := original
	set, present := change.Document.Frontmatter["set"]
	if present {
		values, ok := set.(map[string]any)
		if !ok {
			return files.At(change.Path, "C12", fmt.Errorf("set must be a mapping"))
		}
		doc := module.Document
		doc.Frontmatter = maps.Clone(doc.Frontmatter)
		for _, field := range slices.Sorted(maps.Keys(values)) {
			switch field {
			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only", "boundaries_opt_out", "hooks":
				doc.Frontmatter[field] = values[field]
			default:
				return files.At(change.Path, "C12", fmt.Errorf("invalid set field %q", field))
			}
		}
		var err error
		module, err = fixture.DecodeModule(module.Dir, doc, p.defaults)
		if err != nil {
			return files.At(change.Path, "C12", fmt.Errorf("set: %w", err))
		}
	}

	previous := p.items[i]
	items := previous
	if len(change.Added)+len(change.Modified)+len(change.Removed) > 0 {
		var err error
		items, err = applyEdits(previous, change)
		if err != nil {
			return err
		}
		module.Document.Body = renderSpecs(module.Document.Body, previous, items)
	}

	p.snapshot.Modules[i] = module
	if present {
		if err := fixture.ValidateGraph(p.snapshot.Modules); err != nil {
			p.snapshot.Modules[i] = original
			return files.At(change.Path, "C02", err)
		}
	}
	p.items[i] = items
	return nil
}

func applyEdits(previous []contracts.Item, change model.Change) ([]contracts.Item, error) {
	index := stableIndex(previous)
	items := slices.Clone(previous)
	for _, item := range change.Modified {
		position, ok := index[item.ID]
		if !ok {
			return nil, absentEdit(change.Path, "Modified ", item.ID)
		}
		items[position] = item
	}
	removed := make(map[string]bool)
	for _, item := range change.Removed {
		if _, ok := index[item.ID]; !ok {
			return nil, absentEdit(change.Path, "Removed ", item.ID)
		}
		removed[item.ID] = true
	}
	items = slices.DeleteFunc(items, func(item contracts.Item) bool { return removed[item.ID] })
	for _, item := range change.Added {
		if _, exists := index[item.ID]; exists {
			return nil, files.At(change.Path, "C05", fmt.Errorf("duplicate spec ID %s", item.ID))
		}
		items = append(items, item)
	}
	return items, nil
}

func absentEdit(name, kind, id string) error {
	return files.At(name, "C08", fmt.Errorf("%sID %s absent from preceding projection", kind, id))
}

func renderSpecs(body string, previous, items []contracts.Item) string {
	section, present := contracts.Section(body, "Specs")
	if len(previous) != 0 {
		// Keep prose and whitespace before the first item in the section.
		start := 0
		for _, line := range strings.SplitAfter(section, "'"\\n\") {
			if strings.HasPrefix(line, \"### \") {
				break
			}
			start += len(line)
		}
		section = section[:start]
	}
	if "'!present {
		section = "'"\\n\"
	}
	var rendered strings.Builder
	rendered.WriteString(section)
	needsNewline := section "'!= "" && !strings.HasSuffix(section, "'"\\n\")
	for _, item := range items {
		if needsNewline {
			rendered.WriteByte('\\n')
		}
		text := contracts.RenderItem(item)
		rendered.WriteString(text)
		needsNewline = "'!strings.HasSuffix(text, "'"\\n\")
	}
	if needsNewline {
		rendered.WriteByte('\\n')
	}
	return contracts.ReplaceSection(body, \"Specs\", rendered.String())
}
EOF
cat > packages/project/internal/queue/bases.go <<'EOF'
package queue

import (
	\"fmt\"
	\"maps\"
	\"os\"
	\"path/filepath\"
	\"slices\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func preceding(root, selector string) (projection, model.Change, error) {
	snapshot, changes, err := load(root)
	if err "'!= nil {
		return projection{}, model.Change{}, err
	}
	index, err := resolve(changes, selector)
	if err != nil {
		return projection{}, model.Change{}, err
	}
	p, err := project(snapshot, changes[:index])
	return p, changes[index], err
}

func bases(change model.Change) (map[string]any, error) {
	value, present := change.Document.Frontmatter["bases"]
	if !present {
		return nil, nil
	}
	mapping, ok := value.(map[string]any)
	if !ok {
		return nil, files.At(change.Path, "C08", fmt.Errorf("bases must be a mapping"))
	}
	return mapping, nil
}

func recordedHash(value any) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

func baseEdits(change model.Change) []contracts.Item {
	return slices.Concat(change.Modified, change.Removed)
}

func (p projection) stableItems(module string) map[string]contracts.Item {
	i := fixture.ModuleIndex(p.snapshot.Modules, module)
	result := make(map[string]contracts.Item)
	if i >= 0 {
		for _, item := range p.items[i] {
			if !isHandle(item.ID) {
				result[item.ID] = item
			}
		}
	}
	return result
}

func Stamp(root, selector string) error { return stamp(root, selector, false) }
func StampMissing(root, selector string) error { return stamp(root, selector, true) }

func stamp(root, selector string, missingOnly bool) error {
	p, change, err := preceding(root, selector)
	if err != nil {
		return err
	}
	existing, err := bases(change)
	if err != nil {
		return err
	}
	stamped := make(map[string]any, len(existing))
	maps.Copy(stamped, existing)
	items := p.stableItems(change.Module)
	changed := false
	for _, edit := range baseEdits(change) {
		item, present := items[edit.ID]
		if !present {
			return absentEdit(change.Path, "", edit.ID)
		}
		value, stampedAlready := existing[edit.ID]
		if missingOnly && stampedAlready {
			continue
		}
		hash := contracts.SpecHash(item)
		if old, ok := value.(string); !ok || old != hash {
			stamped[edit.ID] = hash
			changed = true
		}
	}
	if !changed {
		return nil
	}
	change.Document.Frontmatter["bases"] = stamped
	text, err := contracts.RenderDocument(change.Document)
	if err != nil {
		return err
	}
	return writeStamp(filepath.Join(p.snapshot.Root, filepath.FromSlash(change.Path)), text)
}

func writeStamp(name, content string) error {
	info, err := os.Stat(name)
	if err != nil {
		return err
	}
	temp, err := os.CreateTemp(filepath.Dir(name), ".stamp-*")
	if err != nil {
		return err
	}
	defer os.Remove(temp.Name())
	err = temp.Chmod(info.Mode().Perm())
	if err == nil {
		_, err = temp.WriteString(content)
	}
	closeErr := temp.Close()
	if err != nil {
		return err
	}
	if closeErr != nil {
		return closeErr
	}
	return os.Rename(temp.Name(), name)
}
EOF
gofmt -w packages/project/internal/queue' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 494ms:
exec
/bin/zsh -lc "cat > packages/project/internal/queue/check.go <<'EOF'
package queue

import (
	\"errors\"
	\"fmt\"
	\"regexp\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

var normative = regexp.MustCompile("'`'"\\b(?:MUST|SHOULD)\\b"'`)
var handleReference = regexp.MustCompile(`new:[a-z0-9-]+`)

func Check(root string) []model.Diagnostic {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return []model.Diagnostic{diagnostic(err)}
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return []model.Diagnostic{diagnostic(err)}
	}
	var diagnostics []model.Diagnostic
	for i, module := range snapshot.Modules {
		for _, item := range p.items[i] {
			name := modulePath(module)
			if !owns(module, item.ID) {
				diagnostics = append(diagnostics, report(name, "C05", "foreign spec ID "+item.ID))
			}
			diagnostics = append(diagnostics, checkText(name, item)...)
		}
	}
	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules)
	if err != nil {
		return append(diagnostics, diagnostic(err))
	}
	for _, change := range changes {
		i := fixture.ModuleIndex(p.snapshot.Modules, change.Module)
		diagnostics = append(diagnostics, checkChange(change, p.snapshot.Modules[i])...)
		diagnostics = append(diagnostics, p.checkBases(change)...)
		if err := p.apply(change); err != nil {
			diagnostics = append(diagnostics, diagnostic(err))
			break
		}
	}
	return diagnostics
}

func report(path, code, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Path: path, Message: message}
}

func diagnostic(err error) model.Diagnostic {
	code, subject, _ := contracts.ErrorCode(err)
	var at files.PathError
	if errors.As(err, &at) {
		subject = at.Path
	}
	return report(subject, code, err.Error())
}

func owns(module model.Module, id string) bool {
	prefix, _, found := strings.Cut(id, "-")
	return found && prefix == module.IDPrefix
}

func checkText(path string, item contracts.Item) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	if !normative.MatchString(item.Body) {
		diagnostics = append(diagnostics, report(path, "C04", "non-normative spec body for "+item.ID))
	}
	if !validTag(item.Tag) {
		diagnostics = append(diagnostics, report(path, "C06", "invalid review/quarantine tag for "+item.ID))
	}
	return diagnostics
}

func validTag(tag string) bool {
	if tag == "" || tag == "review" {
		return true
	}
	reason, ok := strings.CutPrefix(tag, "quarantine:")
	return ok && strings.TrimSpace(reason) != ""
}

func checkChange(change model.Change, module model.Module) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	name := change.Path
	motivation, _ := contracts.Section(change.Document.Body, "Motivation")
	if strings.TrimSpace(motivation) == "" {
		diagnostics = append(diagnostics, report(name, "C07", "missing Motivation"))
	}
	_, evalFix := change.Document.Frontmatter["eval_fix"]
	_, fixes := change.Document.Frontmatter["fixes"]
	if evalFix && (fixes || len(change.Added)+len(change.Modified)+len(change.Removed) > 0) {
		diagnostics = append(diagnostics, report(name, "C07", "eval_fix cannot be combined with fixes or spec edits"))
	}

	declared := make(map[string]bool)
	for _, item := range change.Added {
		if isHandle(item.ID) {
			declared[item.ID] = true
		}
	}
	for _, items := range [][]contracts.Item{change.Added, change.Modified, change.Removed} {
		for _, item := range items {
			if !isHandle(item.ID) && !owns(module, item.ID) {
				diagnostics = append(diagnostics, report(name, "C05", "foreign module ID "+item.ID))
			}
		}
	}
	for _, items := range [][]contracts.Item{change.Added, change.Modified} {
		for _, item := range items {
			diagnostics = append(diagnostics, checkText(name, item)...)
		}
	}
	for _, item := range change.Removed {
		if !validTag(item.Tag) {
			diagnostics = append(diagnostics, report(name, "C06", "invalid review/quarantine tag for "+item.ID))
		}
	}

	// Rendering exposes frontmatter references as well as references in prose.
	text, err := contracts.RenderDocument(change.Document)
	if err != nil {
		diagnostics = append(diagnostics, diagnostic(files.At(name, "C03", err)))
	}
	for _, handle := range handleReference.FindAllString(text, -1) {
		if !declared[handle] {
			diagnostics = append(diagnostics, report(name, "C09", "reference to undeclared handle "+handle))
			declared[handle] = true
		}
	}
	return diagnostics
}

func (p projection) checkBases(change model.Change) []model.Diagnostic {
	mapping, err := bases(change)
	if err != nil {
		return []model.Diagnostic{diagnostic(err)}
	}
	items := p.stableItems(change.Module)
	var diagnostics []model.Diagnostic
	for _, edit := range baseEdits(change) {
		item, present := items[edit.ID]
		hash := recordedHash(mapping[edit.ID])
		var message string
		switch {
		case !present:
			message = fmt.Sprintf("ID %s absent from preceding projection", edit.ID)
		case hash == "":
			message = "missing base for " + edit.ID
		case hash != contracts.SpecHash(item):
			message = "stale base for " + edit.ID
		}
		if message != "" {
			diagnostics = append(diagnostics, report(change.Path, "C08", message))
		}
	}
	return diagnostics
}
EOF
cat > packages/project/internal/queue/stale.go <<'"'EOF'
package queue

import (
	\"reflect\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/gitstore\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
	p, change, err := preceding(root, selector)
	if err "'!= nil {
		return nil, err
	}
	mapping, err := bases(change)
	if err != nil {
		return nil, err
	}
	current := p.stableItems(change.Module)
	stale := make([]model.StaleSpec, 0)
	for _, edit := range baseEdits(change) {
		value, recorded := mapping[edit.ID]
		if !recorded {
			continue
		}
		hash := recordedHash(value)
		entry := model.StaleSpec{ID: edit.ID, RecordedHash: hash}
		if item, present := current[edit.ID]; present {
			if contracts.SpecHash(item) == hash {
				continue
			}
			entry.Current = contracts.RenderItem(item)
		}
		stale = append(stale, entry)
	}
	if len(stale) == 0 {
		return stale, nil
	}
	historical := recoverItems(p.snapshot.Root, change)
	for i := range stale {
		entry := &stale[i]
		if item, ok := historical[entry.ID]; ok && contracts.SpecHash(item) == entry.RecordedHash {
			entry.Before = contracts.RenderItem(item)
			entry.Recovered = true
		}
	}
	return stale, nil
}

func recoverItems(root string, change model.Change) map[string]contracts.Item {
	commit := lastBasesCommit(root, change.Path)
	if commit == "" {
		return nil
	}
	source, err := files.OpenRevision(root, commit)
	if err != nil {
		return nil
	}
	snapshot, err := fixture.LoadFrom(root, source)
	if err != nil {
		return nil
	}
	changes, err := read(source, snapshot.Modules)
	if err != nil {
		return nil
	}
	index, err := resolve(changes, change.Path)
	if err != nil {
		return nil
	}
	p, err := project(snapshot, changes[:index])
	if err != nil {
		return nil
	}
	return p.stableItems(changes[index].Module)
}

func lastBasesCommit(root, name string) string {
	output, err := files.Git(root, "log", "--first-parent", "--format=%H", "HEAD", "--", ":(literal)"+name)
	if err != nil {
		return ""
	}
	for _, commit := range strings.Fields(output) {
		current, ok := readBasesAt(root, commit, name)
		if !ok {
			return ""
		}
		previous, _ := readBasesAt(root, commit+"''^1", name)
		if !reflect.DeepEqual(current, previous) {
			return commit
		}
	}
	return ""
}

func readBasesAt(root, commit, name string) (map[string]any, bool) {
	data, err := gitstore.ReadAt(root, commit, name)
	if err != nil {
		return nil, false
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return nil, false
	}
	mapping, err := bases(model.Change{Path: name, Document: doc})
	return mapping, err == nil
}
EOF
cat > packages/project/internal/history/history.go <<'"'EOF'
package history

import (
	\"fmt\"
	\"path\"
	\"slices\"
	\"sort\"
	\"strings\"
	\"time\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project/internal/files\"
	\"github.com/lurkshark/fireweed/packages/project/internal/fixture\"
	\"github.com/lurkshark/fireweed/packages/project/model\"
)

func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
	snapshot, err := fixture.Load(root)
	if err "'!= nil {
		return nil, err
	}
	index := fixture.ModuleIndex(snapshot.Modules, module)
	if index < 0 {
		return nil, unknown("module", module)
	}
	source := files.Worktree{Root: snapshot.Root}
	archive, err := source.ReadDir(files.Archive)
	if err != nil {
		// Share the change-file convention and missing-directory handling.
		if _, err := files.Changes(source, files.Archive); err != nil {
			return nil, archiveError(err)
		}
	}

	entries := make([]model.SpecHistoryEntry, 0)
	for _, file := range archive {
		if file.Mode.IsDir() || !strings.HasSuffix(file.Name, ".md") {
			continue
		}
		name := path.Join(files.Archive, file.Name)
		doc, err := files.ReadDocument(source, name)
		if err != nil {
			return nil, archiveError(err)
		}
		target, ok := doc.Frontmatter["module"].(string)
		if !ok || target == "" {
			return nil, archiveError(files.At(name, "C03", fmt.Errorf("archive module must be a nonempty string")))
		}
		ids, err := affectedIDs(doc, id)
		if err != nil {
			return nil, archiveError(files.Syntax(name, err))
		}
		if target != module || len(ids) == 0 {
			continue
		}
		entry := model.SpecHistoryEntry{Path: name, Module: target, IDs: ids}
		entry.Motivation, _ = contracts.Section(doc.Body, "Motivation")
		entry.Decisions, _ = contracts.Section(doc.Body, "Decisions")
		entry.RemovedDecisions, _ = contracts.Section(doc.Body, "Removed decisions")
		entry.Provenance, _ = contracts.Section(doc.Body, "Provenance")
		entry.ImplementationDate = implementationDate(entry.Provenance)
		entries = append(entries, entry)
	}
	if id != "" && len(entries) == 0 {
		items, err := files.Items(snapshot.Modules[index].Document, "Specs")
		if err != nil {
			return nil, err
		}
		if !slices.ContainsFunc(items, func(item contracts.Item) bool { return item.ID == id }) {
			return nil, unknown("spec ID", id)
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if a.ImplementationDate.Equal(b.ImplementationDate) {
			return a.Path < b.Path
		}
		return a.ImplementationDate.After(b.ImplementationDate)
	})
	return entries, nil
}

func affectedIDs(doc contracts.Document, id string) ([]string, error) {
	seen := make(map[string]bool)
	for _, section := range []string{"Added", "Modified"} {
		items, err := files.Items(doc, section)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			if id == "" || item.ID == id {
				seen[item.ID] = true
			}
		}
	}
	ids := make([]string, 0, len(seen))
	for affected := range seen {
		ids = append(ids, affected)
	}
	slices.Sort(ids)
	return ids, nil
}

func implementationDate(provenance string) time.Time {
	var latest time.Time
	for _, line := range strings.Split(provenance, "'"\\n\") {
		fields := strings.Fields(line)
		if len(fields) < 3 || fields[0] "'!= "-" || fields[1] != "impl" {
			continue
		}
		date, err := time.Parse(time.RFC3339, fields[2])
		if err == nil && (latest.IsZero() || date.After(latest)) {
			latest = date
		}
	}
	return latest
}

func archiveError(err error) error {
	return files.Code("C03", files.Archive, fmt.Errorf("C03: %w", err))
}

func unknown(kind, name string) error {
	return files.Code("U02", name, fmt.Errorf("U02: unknown %s %q", kind, name))
}
EOF
gofmt -w packages/project/internal
export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache GOPROXY=off
mkdir -p .fireweed/runs/drill && sh -c '"'go run .fireweed/hooks.go eval '\"'\"'packages/project'\"'\"' '\"'\"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/.fireweed/runs/drill/report.xml'\"'\"''" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 4666ms:
?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.398s

exec
/bin/zsh -lc "git status --short --ignored; git diff --check; if [ -f .fireweed/runs/drill/report.xml ]; then python3 -c 'import xml.etree.ElementTree as ET; t=ET.parse(\".fireweed/runs/drill/report.xml\"); cases=t.findall(\"testcase\"); print(len(cases), \"cases\", len(t.findall(\"testcase/failure\")), \"failures\"); [print(c.attrib[\"name\"], c.find(\"failure\").text) for c in cases if c.find(\"failure\") is not None]'; fi; rg -n 'C01|C07|C12|version' packages/project/evals" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 768ms:
?? packages/project/internal/
!! .fireweed/bin/
!! .fireweed/runs/
134 cases 0 failures
packages/project/evals/static_codes_test.go:34:				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:38:			t.Errorf("Check did not report C01 for configuration cause %q", cause)
packages/project/evals/static_codes_test.go:154:				if !diagnostic.Warning && diagnostic.Code == "C12" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:158:			t.Errorf("Check did not report C12 for %s cause %q", tc.source, cause)
packages/project/evals/static_codes_test.go:206:		if !diagnostic.Warning && diagnostic.Code == "C07" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:210:	t.Errorf("Check did not report C07 for pending module cause %q", cause)
packages/project/evals/mixed_eval_fix_test.go:15:		wantC07            bool
packages/project/evals/mixed_eval_fix_test.go:36:			if tc.wantC07 {
packages/project/evals/mixed_eval_fix_test.go:37:				if len(errors) != 1 || errors[0].Code != "C07" {
packages/project/evals/mixed_eval_fix_test.go:38:					t.Errorf("mixed change error diagnostics = %#v; want only error C07", errors)
packages/project/evals/project_test.go:17:version: 1
packages/project/evals/project_test.go:57:	if got.Config["version"] != 1 {
packages/project/evals/project_test.go:58:		t.Errorf("decoded Config version = %#v", got.Config["version"])
packages/project/evals/queue_test.go:297:		{"PRJ-008", "missing_motivation", "", "", "\n", "", "", "", "C07", "missing Motivation"},
packages/project/evals/revision_snapshot_test.go:36:	put(t, root, "fireweed.yaml", `version: 1
packages/project/evals/revision_snapshot_test.go:51:	put(t, root, "fireweed.yaml", "version: 2\nmodules: [\"components/*\"]\nmodule_defaults:\n  evals: [later/**]\n")
packages/project/evals/revision_snapshot_test.go:56:	put(t, root, "fireweed.yaml", "version: 3\nmodules: [\"dirty/*\"]\n")
packages/project/evals/revision_snapshot_test.go:68:	if got.Config["version"] != 1 || !reflect.DeepEqual(got.Config["modules"], []any{"packages/*"}) {
packages/project/evals/spec_history_test.go:210:			if err == nil || !strings.HasPrefix(err.Error(), "C01: ") {
packages/project/evals/spec_history_test.go:211:				t.Errorf("invalid config = %v; want C01: prefix", err)

codex
The target hook passed all 134 cases. I’ll review the paths those evals exercise for edge cases in historical recovery, queue selection and preservation of change text before the final run.
exec
/bin/zsh -lc "wc -l packages/project/internal/*/*.go; git status --short; sed -n '1,150p' packages/project/internal/files/source.go; rg -n 'PRJ-(00[4-9]|02[3-9]|03[0-6]|05[1-6]|06[0-6])' packages/project/SPEC.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 502ms:
      60 packages/project/internal/files/changes.go
     165 packages/project/internal/files/source.go
      65 packages/project/internal/fixture/graph.go
     165 packages/project/internal/fixture/load.go
     104 packages/project/internal/fixture/metadata.go
      72 packages/project/internal/fixture/new.go
     125 packages/project/internal/history/history.go
     129 packages/project/internal/queue/bases.go
     165 packages/project/internal/queue/check.go
     201 packages/project/internal/queue/project.go
      88 packages/project/internal/queue/queue.go
     111 packages/project/internal/queue/stale.go
    1450 total
?? packages/project/internal/
// Package files reads project documents from the worktree or a recorded Git tree.
package files

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
)

type Source interface {
	ReadFile(string) ([]byte, error)
	ReadDir(string) ([]Entry, error)
}

type Entry struct {
	Name string
	Mode fs.FileMode
}

type Worktree struct{ Root string }

func (w Worktree) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(os.DirFS(w.Root), name)
}

func (w Worktree) ReadDir(name string) ([]Entry, error) {
	entries, err := fs.ReadDir(os.DirFS(w.Root), name)
	if err != nil {
		return nil, err
	}
	result := make([]Entry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, Entry{entry.Name(), entry.Type()})
	}
	return result, nil
}

type Revision struct {
	Root   string
	commit string
	dirs   map[string][]Entry
	modes  map[string]fs.FileMode
}

func Git(root string, args ...string) (string, error) {
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Dir:     root,
		Args:    append([]string{"git", "--no-optional-locks"}, args...),
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
		Timeout: 30 * time.Second,
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}

func OpenRevision(root, revision string) (*Revision, error) {
	output, err := Git(root, "rev-parse", "--verify", "--end-of-options", revision+"^{commit}")
	if err != nil {
		return nil, err
	}
	commit := strings.TrimSpace(output)
	output, err = Git(root, "ls-tree", "--full-tree", "-r", "-t", "-z", commit)
	if err != nil {
		return nil, err
	}
	tree := &Revision{
		Root: root, commit: commit,
		dirs:  map[string][]Entry{".": nil},
		modes: make(map[string]fs.FileMode),
	}
	for _, record := range strings.Split(strings.TrimSuffix(output, "\x00"), "\x00") {
		if record == "" {
			continue
		}
		header, name, _ := strings.Cut(record, "\t")
		mode, _, _ := strings.Cut(header, " ")
		var kind fs.FileMode
		switch mode {
		case "040000":
			kind = fs.ModeDir
			tree.dirs[name] = nil
		case "120000":
			kind = fs.ModeSymlink
		case "100644", "100755":
		default:
			kind = fs.ModeIrregular
		}
		parent := path.Dir(name)
		tree.dirs[parent] = append(tree.dirs[parent], Entry{path.Base(name), kind})
		tree.modes[name] = kind
	}
	return tree, nil
}

func (r *Revision) ReadFile(name string) ([]byte, error) {
	mode, exists := r.modes[name]
	if !exists {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrNotExist}
	}
	if !mode.IsRegular() {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrInvalid}
	}
	return gitstore.ReadAt(r.Root, r.commit, name)
}

func (r *Revision) ReadDir(name string) ([]Entry, error) {
	entries, exists := r.dirs[name]
	if !exists {
		cause := fs.ErrNotExist
		if _, present := r.modes[name]; present {
			cause = fs.ErrInvalid
		}
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: cause}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })
	return entries, nil
}

type PathError struct {
	Path string
	Err  error
}

func (e PathError) Error() string { return e.Path + ": " + e.Err.Error() }
func (e PathError) Unwrap() error { return e.Err }

func Code(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
}

func At(name, code string, err error) error {
	return PathError{Path: name, Err: Code(code, name, err)}
}

// Syntax preserves the parser's classification and subject.
func Syntax(name string, err error) error {
	return PathError{Path: name, Err: err}
}

func ReadDocument(source Source, name string) (contracts.Document, error) {
26:### PRJ-004
29:### PRJ-005
32:### PRJ-006
35:### PRJ-007
38:### PRJ-008
41:### PRJ-009
86:### PRJ-023
89:### PRJ-024
92:### PRJ-025
96:### PRJ-026
99:### PRJ-027
102:### PRJ-028
105:### PRJ-029
108:### PRJ-030
111:### PRJ-031
114:### PRJ-032
117:### PRJ-033
120:### PRJ-034
123:### PRJ-035
126:### PRJ-036
172:### PRJ-051
176:### PRJ-052
179:### PRJ-053
183:### PRJ-054
187:### PRJ-055
194:### PRJ-056
208:### PRJ-060
212:### PRJ-061
216:### PRJ-062
220:### PRJ-063
224:### PRJ-064
228:### PRJ-065
231:### PRJ-066

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/fixture/load.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/fixture/metadata.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/queue/bases.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/queue/check.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/queue/project.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/queue/queue.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/queue/stale.go
diff --git a/packages/project/internal/fixture/load.go b/packages/project/internal/fixture/load.go
index 01a45db4ba2d9e0a6faad4ff3228c72bf8471ba0..6f404bd75bc9abaf8492dc0aaa6ae2139668ea0a
--- a/packages/project/internal/fixture/load.go
+++ b/packages/project/internal/fixture/load.go
@@ -72,7 +72,7 @@
 		}
 		patterns[i] = clean
 	}
-	defaults, err := Defaults(config)
+	defaults, err := moduleDefaults(config)
 	if err != nil {
 		return model.Snapshot{}, files.At(configPath, "C12", err)
 	}
diff --git a/packages/project/internal/fixture/metadata.go b/packages/project/internal/fixture/metadata.go
index a248952ebcc823a0bc45282a8b78727c10e6554b..5bb697c0a2e9902f399d5b886a2b88e0f43713f4
--- a/packages/project/internal/fixture/metadata.go
+++ b/packages/project/internal/fixture/metadata.go
@@ -21,7 +21,7 @@
 	return result, nil
 }
 
-func Defaults(config map[string]any) (map[string]any, error) {
+func moduleDefaults(config map[string]any) (map[string]any, error) {
 	raw, present := config["module_defaults"]
 	if !present {
 		return nil, nil
@@ -70,7 +70,7 @@
 	for _, field := range []struct {
 		key    string
 		target *string
-	}{{"name", &module.Name}, {"id_prefix", &module.IDPrefix}, {"boundaries_opt_out", new(string)}} {
+	}{{"name", &module.Name}, {"id_prefix", &module.IDPrefix}} {
 		if value, present := fields[field.key]; present {
 			var ok bool
 			if *field.target, ok = value.(string); !ok {
@@ -78,6 +78,11 @@
 			}
 		}
 	}
+	if reason, present := fields["boundaries_opt_out"]; present {
+		if _, ok := reason.(string); !ok {
+			return module, fmt.Errorf("boundaries_opt_out must be a string")
+		}
+	}
 	if hooks, present := fields["hooks"]; present {
 		if _, ok := hooks.(map[string]any); !ok {
 			return module, fmt.Errorf("hooks must be a mapping")
diff --git a/packages/project/internal/queue/bases.go b/packages/project/internal/queue/bases.go
index ecdfa65e8d416bf01f7fbb600b5071b9af4eacf4..a99dcb7656bf124f3ea53feb7f0714e5ddf7e88c
--- a/packages/project/internal/queue/bases.go
+++ b/packages/project/internal/queue/bases.go
@@ -14,14 +14,11 @@
 )
 
 func preceding(root, selector string) (projection, model.Change, error) {
-	snapshot, changes, err := load(root)
-	if err != nil {
-		return projection{}, model.Change{}, err
-	}
-	index, err := resolve(changes, selector)
+	snapshot, changes, err := load(root, &selector)
 	if err != nil {
 		return projection{}, model.Change{}, err
 	}
+	index := len(changes) - 1
 	p, err := project(snapshot, changes[:index])
 	return p, changes[index], err
 }
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
index 24498cbfb696120ec212adb5a706cb5b6c15dd24..f3b900fb1ce066fc2d4d31c9c416b759f7161638
--- a/packages/project/internal/queue/check.go
+++ b/packages/project/internal/queue/check.go
@@ -34,7 +34,7 @@
 			diagnostics = append(diagnostics, checkText(name, item)...)
 		}
 	}
-	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules)
+	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules, nil)
 	if err != nil {
 		return append(diagnostics, diagnostic(err))
 	}
@@ -148,13 +148,15 @@
 	for _, edit := range baseEdits(change) {
 		item, present := items[edit.ID]
 		hash := recordedHash(mapping[edit.ID])
+		if !present {
+			message := fmt.Sprintf("ID %s absent from preceding projection", edit.ID)
+			diagnostics = append(diagnostics, report(change.Path, "C08", message))
+		}
 		var message string
 		switch {
-		case !present:
-			message = fmt.Sprintf("ID %s absent from preceding projection", edit.ID)
 		case hash == "":
 			message = "missing base for " + edit.ID
-		case hash != contracts.SpecHash(item):
+		case present && hash != contracts.SpecHash(item):
 			message = "stale base for " + edit.ID
 		}
 		if message != "" {
diff --git a/packages/project/internal/queue/project.go b/packages/project/internal/queue/project.go
index de0cc02bdfbe32cb3eccd0d9e9e57a87276d0644..a3951766dcf37c03fed48340885290221de6258b
--- a/packages/project/internal/queue/project.go
+++ b/packages/project/internal/queue/project.go
@@ -3,6 +3,7 @@
 import (
 	"fmt"
 	"maps"
+	"path"
 	"slices"
 	"strings"
 
@@ -13,17 +14,14 @@
 )
 
 func Project(root, through string) (model.Snapshot, error) {
-	snapshot, changes, err := load(root)
+	var selector *string
+	if through != "" {
+		selector = &through
+	}
+	snapshot, changes, err := load(root, selector)
 	if err != nil {
 		return model.Snapshot{}, err
 	}
-	if through != "" {
-		index, err := resolve(changes, through)
-		if err != nil {
-			return model.Snapshot{}, err
-		}
-		changes = changes[:index+1]
-	}
 	projected, err := project(snapshot, changes)
 	if err != nil {
 		return model.Snapshot{}, err
@@ -67,10 +65,7 @@
 }
 
 func modulePath(module model.Module) string {
-	if module.Dir == "." {
-		return "SPEC.md"
-	}
-	return module.Dir + "/SPEC.md"
+	return path.Join(module.Dir, "SPEC.md")
 }
 
 // stableIndex excludes handles. They are rendered as written, but they never
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 25fac70c3d7f55f2126dd7be5093c6722f6e5457..5499d3540781d979958633a7f211ce63fe5c2a03
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -12,20 +12,22 @@
 )
 
 func Queue(root string) ([]model.Change, error) {
-	_, changes, err := load(root)
+	_, changes, err := load(root, nil)
 	return changes, err
 }
 
-func load(root string) (model.Snapshot, []model.Change, error) {
+func load(root string, through *string) (model.Snapshot, []model.Change, error) {
 	snapshot, err := fixture.Load(root)
 	if err != nil {
 		return model.Snapshot{}, nil, err
 	}
-	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules)
+	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules, through)
 	return snapshot, changes, err
 }
 
-func read(source files.Source, modules []model.Module) ([]model.Change, error) {
+// A nil selector reads the whole queue. For a selected prefix, validate all
+// queue identities but only parse documents through the selected change.
+func read(source files.Source, modules []model.Module, through *string) ([]model.Change, error) {
 	pending, err := files.Changes(source, files.Pending)
 	if err != nil {
 		return nil, err
@@ -37,6 +39,13 @@
 		}
 		slugs[entry.Slug] = entry.Path
 	}
+	if through != nil {
+		index, err := resolve(pending, *through)
+		if err != nil {
+			return nil, err
+		}
+		pending = pending[:index+1]
+	}
 	changes := make([]model.Change, 0, len(pending))
 	for _, entry := range pending {
 		doc, err := files.ReadDocument(source, entry.Path)
@@ -72,7 +81,7 @@
 	return changes, nil
 }
 
-func resolve(changes []model.Change, selector string) (int, error) {
+func resolve(changes []files.ChangeFile, selector string) (int, error) {
 	slugs := make([]string, 0, len(changes))
 	for i, change := range changes {
 		if selector == change.Slug || selector == path.Base(change.Path) || path.Clean(selector) == change.Path {
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
index bd85c9f852ede074dc8bae739772a7fb0a3f9f21..5425ac35bb1148fa0d1c5f266b89f92eb3049667
--- a/packages/project/internal/queue/stale.go
+++ b/packages/project/internal/queue/stale.go
@@ -64,14 +64,11 @@
 	if err != nil {
 		return nil
 	}
-	changes, err := read(source, snapshot.Modules)
-	if err != nil {
-		return nil
-	}
-	index, err := resolve(changes, change.Path)
+	changes, err := read(source, snapshot.Modules, &change.Path)
 	if err != nil {
 		return nil
 	}
+	index := len(changes) - 1
 	p, err := project(snapshot, changes[:index])
 	if err != nil {
 		return nil
@@ -80,7 +77,7 @@
 }
 
 func lastBasesCommit(root, name string) string {
-	output, err := files.Git(root, "log", "--first-parent", "--format=%H", "HEAD", "--", ":(literal)"+name)
+	output, err := files.Git(root, "log", "--first-parent", "--full-history", "--format=%H", "HEAD", "--", ":(literal)"+name)
 	if err != nil {
 		return ""
 	}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/history/history.go
diff --git a/packages/project/internal/fixture/load.go b/packages/project/internal/fixture/load.go
index 01a45db4ba2d9e0a6faad4ff3228c72bf8471ba0..6f404bd75bc9abaf8492dc0aaa6ae2139668ea0a
--- a/packages/project/internal/fixture/load.go
+++ b/packages/project/internal/fixture/load.go
@@ -72,7 +72,7 @@
 		}
 		patterns[i] = clean
 	}
-	defaults, err := Defaults(config)
+	defaults, err := moduleDefaults(config)
 	if err != nil {
 		return model.Snapshot{}, files.At(configPath, "C12", err)
 	}
diff --git a/packages/project/internal/fixture/metadata.go b/packages/project/internal/fixture/metadata.go
index a248952ebcc823a0bc45282a8b78727c10e6554b..5bb697c0a2e9902f399d5b886a2b88e0f43713f4
--- a/packages/project/internal/fixture/metadata.go
+++ b/packages/project/internal/fixture/metadata.go
@@ -21,7 +21,7 @@
 	return result, nil
 }
 
-func Defaults(config map[string]any) (map[string]any, error) {
+func moduleDefaults(config map[string]any) (map[string]any, error) {
 	raw, present := config["module_defaults"]
 	if !present {
 		return nil, nil
@@ -70,7 +70,7 @@
 	for _, field := range []struct {
 		key    string
 		target *string
-	}{{"name", &module.Name}, {"id_prefix", &module.IDPrefix}, {"boundaries_opt_out", new(string)}} {
+	}{{"name", &module.Name}, {"id_prefix", &module.IDPrefix}} {
 		if value, present := fields[field.key]; present {
 			var ok bool
 			if *field.target, ok = value.(string); !ok {
@@ -78,6 +78,11 @@
 			}
 		}
 	}
+	if reason, present := fields["boundaries_opt_out"]; present {
+		if _, ok := reason.(string); !ok {
+			return module, fmt.Errorf("boundaries_opt_out must be a string")
+		}
+	}
 	if hooks, present := fields["hooks"]; present {
 		if _, ok := hooks.(map[string]any); !ok {
 			return module, fmt.Errorf("hooks must be a mapping")
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
index 4ce4f09f8b2e577586962bbd7a1cde9ff5c840b7..fb4085c766fd1020d4fe967e2311913629aee700
--- a/packages/project/internal/history/history.go
+++ b/packages/project/internal/history/history.go
@@ -1,7 +1,9 @@
 package history
 
 import (
+	"errors"
 	"fmt"
+	"io/fs"
 	"path"
 	"slices"
 	"sort"
@@ -17,6 +19,9 @@
 func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
 	snapshot, err := fixture.Load(root)
 	if err != nil {
+		if code, subject, _ := contracts.ErrorCode(err); code == "C12" && subject == "fireweed.yaml" {
+			return nil, files.Code("C01", subject, fmt.Errorf("C01: %w", err))
+		}
 		return nil, err
 	}
 	index := fixture.ModuleIndex(snapshot.Modules, module)
@@ -25,14 +30,12 @@
 	}
 	source := files.Worktree{Root: snapshot.Root}
 	archive, err := source.ReadDir(files.Archive)
-	if err != nil {
-		// Share the change-file convention and missing-directory handling.
-		if _, err := files.Changes(source, files.Archive); err != nil {
-			return nil, archiveError(err)
-		}
+	if err != nil && !errors.Is(err, fs.ErrNotExist) {
+		return nil, archiveError(files.At(files.Archive, "C03", err))
 	}
 
 	entries := make([]model.SpecHistoryEntry, 0)
+	dated := make(map[string]bool)
 	for _, file := range archive {
 		if file.Mode.IsDir() || !strings.HasSuffix(file.Name, ".md") {
 			continue
@@ -58,7 +61,7 @@
 		entry.Decisions, _ = contracts.Section(doc.Body, "Decisions")
 		entry.RemovedDecisions, _ = contracts.Section(doc.Body, "Removed decisions")
 		entry.Provenance, _ = contracts.Section(doc.Body, "Provenance")
-		entry.ImplementationDate = implementationDate(entry.Provenance)
+		entry.ImplementationDate, dated[name] = implementationDate(entry.Provenance)
 		entries = append(entries, entry)
 	}
 	if id != "" && len(entries) == 0 {
@@ -72,6 +75,9 @@
 	}
 	sort.Slice(entries, func(i, j int) bool {
 		a, b := entries[i], entries[j]
+		if dated[a.Path] != dated[b.Path] {
+			return dated[a.Path]
+		}
 		if a.ImplementationDate.Equal(b.ImplementationDate) {
 			return a.Path < b.Path
 		}
@@ -101,19 +107,21 @@
 	return ids, nil
 }
 
-func implementationDate(provenance string) time.Time {
+func implementationDate(provenance string) (time.Time, bool) {
 	var latest time.Time
+	dated := false
 	for _, line := range strings.Split(provenance, "\n") {
 		fields := strings.Fields(line)
 		if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
 			continue
 		}
 		date, err := time.Parse(time.RFC3339, fields[2])
-		if err == nil && (latest.IsZero() || date.After(latest)) {
+		if err == nil && (!dated || date.After(latest)) {
 			latest = date
+			dated = true
 		}
 	}
-	return latest
+	return latest, dated
 }
 
 func archiveError(err error) error {
diff --git a/packages/project/internal/queue/bases.go b/packages/project/internal/queue/bases.go
index ecdfa65e8d416bf01f7fbb600b5071b9af4eacf4..a99dcb7656bf124f3ea53feb7f0714e5ddf7e88c
--- a/packages/project/internal/queue/bases.go
+++ b/packages/project/internal/queue/bases.go
@@ -14,14 +14,11 @@
 )
 
 func preceding(root, selector string) (projection, model.Change, error) {
-	snapshot, changes, err := load(root)
-	if err != nil {
-		return projection{}, model.Change{}, err
-	}
-	index, err := resolve(changes, selector)
+	snapshot, changes, err := load(root, &selector)
 	if err != nil {
 		return projection{}, model.Change{}, err
 	}
+	index := len(changes) - 1
 	p, err := project(snapshot, changes[:index])
 	return p, changes[index], err
 }
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
index 24498cbfb696120ec212adb5a706cb5b6c15dd24..f3b900fb1ce066fc2d4d31c9c416b759f7161638
--- a/packages/project/internal/queue/check.go
+++ b/packages/project/internal/queue/check.go
@@ -34,7 +34,7 @@
 			diagnostics = append(diagnostics, checkText(name, item)...)
 		}
 	}
-	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules)
+	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules, nil)
 	if err != nil {
 		return append(diagnostics, diagnostic(err))
 	}
@@ -148,13 +148,15 @@
 	for _, edit := range baseEdits(change) {
 		item, present := items[edit.ID]
 		hash := recordedHash(mapping[edit.ID])
+		if !present {
+			message := fmt.Sprintf("ID %s absent from preceding projection", edit.ID)
+			diagnostics = append(diagnostics, report(change.Path, "C08", message))
+		}
 		var message string
 		switch {
-		case !present:
-			message = fmt.Sprintf("ID %s absent from preceding projection", edit.ID)
 		case hash == "":
 			message = "missing base for " + edit.ID
-		case hash != contracts.SpecHash(item):
+		case present && hash != contracts.SpecHash(item):
 			message = "stale base for " + edit.ID
 		}
 		if message != "" {
diff --git a/packages/project/internal/queue/project.go b/packages/project/internal/queue/project.go
index de0cc02bdfbe32cb3eccd0d9e9e57a87276d0644..a3951766dcf37c03fed48340885290221de6258b
--- a/packages/project/internal/queue/project.go
+++ b/packages/project/internal/queue/project.go
@@ -3,6 +3,7 @@
 import (
 	"fmt"
 	"maps"
+	"path"
 	"slices"
 	"strings"
 
@@ -13,17 +14,14 @@
 )
 
 func Project(root, through string) (model.Snapshot, error) {
-	snapshot, changes, err := load(root)
+	var selector *string
+	if through != "" {
+		selector = &through
+	}
+	snapshot, changes, err := load(root, selector)
 	if err != nil {
 		return model.Snapshot{}, err
 	}
-	if through != "" {
-		index, err := resolve(changes, through)
-		if err != nil {
-			return model.Snapshot{}, err
-		}
-		changes = changes[:index+1]
-	}
 	projected, err := project(snapshot, changes)
 	if err != nil {
 		return model.Snapshot{}, err
@@ -67,10 +65,7 @@
 }
 
 func modulePath(module model.Module) string {
-	if module.Dir == "." {
-		return "SPEC.md"
-	}
-	return module.Dir + "/SPEC.md"
+	return path.Join(module.Dir, "SPEC.md")
 }
 
 // stableIndex excludes handles. They are rendered as written, but they never
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 25fac70c3d7f55f2126dd7be5093c6722f6e5457..5499d3540781d979958633a7f211ce63fe5c2a03
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -12,20 +12,22 @@
 )
 
 func Queue(root string) ([]model.Change, error) {
-	_, changes, err := load(root)
+	_, changes, err := load(root, nil)
 	return changes, err
 }
 
-func load(root string) (model.Snapshot, []model.Change, error) {
+func load(root string, through *string) (model.Snapshot, []model.Change, error) {
 	snapshot, err := fixture.Load(root)
 	if err != nil {
 		return model.Snapshot{}, nil, err
 	}
-	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules)
+	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules, through)
 	return snapshot, changes, err
 }
 
-func read(source files.Source, modules []model.Module) ([]model.Change, error) {
+// A nil selector reads the whole queue. For a selected prefix, validate all
+// queue identities but only parse documents through the selected change.
+func read(source files.Source, modules []model.Module, through *string) ([]model.Change, error) {
 	pending, err := files.Changes(source, files.Pending)
 	if err != nil {
 		return nil, err
@@ -37,6 +39,13 @@
 		}
 		slugs[entry.Slug] = entry.Path
 	}
+	if through != nil {
+		index, err := resolve(pending, *through)
+		if err != nil {
+			return nil, err
+		}
+		pending = pending[:index+1]
+	}
 	changes := make([]model.Change, 0, len(pending))
 	for _, entry := range pending {
 		doc, err := files.ReadDocument(source, entry.Path)
@@ -72,7 +81,7 @@
 	return changes, nil
 }
 
-func resolve(changes []model.Change, selector string) (int, error) {
+func resolve(changes []files.ChangeFile, selector string) (int, error) {
 	slugs := make([]string, 0, len(changes))
 	for i, change := range changes {
 		if selector == change.Slug || selector == path.Base(change.Path) || path.Clean(selector) == change.Path {
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
index bd85c9f852ede074dc8bae739772a7fb0a3f9f21..5425ac35bb1148fa0d1c5f266b89f92eb3049667
--- a/packages/project/internal/queue/stale.go
+++ b/packages/project/internal/queue/stale.go
@@ -64,14 +64,11 @@
 	if err != nil {
 		return nil
 	}
-	changes, err := read(source, snapshot.Modules)
-	if err != nil {
-		return nil
-	}
-	index, err := resolve(changes, change.Path)
+	changes, err := read(source, snapshot.Modules, &change.Path)
 	if err != nil {
 		return nil
 	}
+	index := len(changes) - 1
 	p, err := project(snapshot, changes[:index])
 	if err != nil {
 		return nil
@@ -80,7 +77,7 @@
 }
 
 func lastBasesCommit(root, name string) string {
-	output, err := files.Git(root, "log", "--first-parent", "--format=%H", "HEAD", "--", ":(literal)"+name)
+	output, err := files.Git(root, "log", "--first-parent", "--full-history", "--format=%H", "HEAD", "--", ":(literal)"+name)
 	if err != nil {
 		return ""
 	}

exec
/bin/zsh -lc "gofmt -w packages/project/internal
cat > DRILL_NOTES.md <<'EOF'
## Contradicted Decisions

None. Spec history uses the current archive files, returns section contents
without their headings, and returns only the matching ID for an ID-filtered
entry, as specified by 0126-spec-history-query.

## Gaps

- **PRJ-001 / project.** The design calls for a warning about SPEC.md files
  outside configured module globs. The project spec assigns no diagnostic for
  those files. Discovery stays within the configured globs.
- **PRJ-004 and PRJ-032 / project.** The specs leave open whether a malformed
  document after the selected change invalidates a partial projection. A
  partial projection validates queue identities across the queue and parses
  documents through the selected change.
- **PRJ-016 / project.** The retained eval requires two C08 diagnostics for an
  absent edit target, with generic and section-specific wording. The spec
  requires C08 and leaves the number and wording of diagnostics open.
- **PRJ-023 and PRJ-027 / project.** The filename format leaves numbering
  beyond 9999 implicit. The queue accepts at least four decimal digits so a
  change created after 9999 remains discoverable.
- **PRJ-035 / project.** The spec leaves validation timing open when a later
  change repairs a graph made invalid by an earlier change. Each projected
  graph is validated as its change is applied.
- **PRJ-060 and PRJ-061 / project.** The specs leave the stale result for a
  removed current item and the Git history traversal unspecified. An item with
  a recorded base and no current counterpart has empty Current text. Recovery
  follows HEAD's first-parent history at the queued path.
EOF
git diff --check; git status --short --ignored" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 697ms:
?? DRILL_NOTES.md
?? packages/project/internal/
!! .fireweed/bin/
!! .fireweed/runs/

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache GOPROXY=off
go vet ./packages/project/...' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache GOPROXY=off
fireweed context project
fireweed verify
git diff HEAD --check
git status --short --untracked-files=all' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache GOPROXY=off
mkdir -p .fireweed/runs/drill && sh -c '"'go run .fireweed/hooks.go eval '\"'\"'packages/project'\"'\"' '\"'\"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/.fireweed/runs/drill/report.xml'\"'\"''
status="'$?
python3 -c '"'import xml.etree.ElementTree as ET; root=ET.parse(\".fireweed/runs/drill/report.xml\"); print(f\"{len(root.findall('\"'testcase'\"'))} cases, {len(root.findall('\"'testcase/failure'\"'))} failures, {len(root.findall('\"'testcase/skipped'\"'))} skipped\")'
exit \""'$status"' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 878ms:
U02: change "project" not found; no valid change slugs. Use a change held by this worktree and retry context.
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.
?? DRILL_NOTES.md
?? packages/project/internal/files/changes.go
?? packages/project/internal/files/source.go
?? packages/project/internal/fixture/graph.go
?? packages/project/internal/fixture/load.go
?? packages/project/internal/fixture/metadata.go
?? packages/project/internal/fixture/new.go
?? packages/project/internal/history/history.go
?? packages/project/internal/queue/bases.go
?? packages/project/internal/queue/check.go
?? packages/project/internal/queue/project.go
?? packages/project/internal/queue/queue.go
?? packages/project/internal/queue/stale.go

 succeeded in 1166ms:
 exited 1 in 4944ms:
?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.757s
zsh:3: read-only variable: status

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache GOPROXY=off
mkdir -p .fireweed/runs/drill &&
sh -c '"'go run .fireweed/hooks.go eval '\"'\"'packages/project'\"'\"' '\"'\"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/.fireweed/runs/drill/report.xml'\"'\"'' &&
python3 - <<'PY'
import xml.etree.ElementTree as ET
root = ET.parse(\".fireweed/runs/drill/report.xml\")
cases = root.findall(\"testcase\")
failures = root.findall(\"testcase/failure\")
skipped = root.findall(\"testcase/skipped\")
print(f\"{len(cases)} cases, {len(failures)} failures, {len(skipped)} skipped\")
assert cases and not failures
PY" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 2081ms:
?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	(cached)
134 cases, 0 failures, 0 skipped

codex
Rebuilt `project` within its disposable paths. The target eval hook passed all 134 cases, and `go vet ./packages/project/...` passed. `fireweed verify` returned P02 because the drill workspace has no held change.

[DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/DRILL_NOTES.md) records six spec gaps and no contradicted Decisions. It documents the choices to validate the graph after each applied change and use first-parent Git history for stale-base recovery.
tokens used
141,235


### Attempt 2 (exit 0)
Updated the project rebuild to return C13 for duplicate slugs, match the graph and queue wording expected by app, and supply C01 and U02 separately from the CLI message text. Spec history retains its required textual prefixes.

All 134 project eval cases passed, along with focused checks of the public diagnostic behavior. The app suite is absent from this workspace, so the downstream failures could not be rerun here. `fireweed verify` returned P02 because the worktree holds no change.

Changes are confined to project disposable paths and [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/DRILL_NOTES.md).

Reading additional input from stdin...
2026-11-10T17:09:39.406229Z  WARN codex_rollout::list: state db discrepancy during find_thread_path_by_id_str_in_subdir: falling_back
OpenAI Codex v0.167.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
model: gpt-6.1-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a1c66f-cf3d-7db0-892e-f887113210e5
--------
user
# Rebuild project

## Target spec and Decisions

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

### PRJ-024
NewChange MUST reject a slug already present in the pending queue or archive without overwriting existing files.

### PRJ-025
NewChange MUST reject an unknown target module with an error beginning
`unknown module ` and naming the rejected module.

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
An unknown module or an ID absent from both current module Specs and matching
archive edits MUST return an error beginning `U02: ` and naming the rejected name.
A current ID with no matching archive edits MUST return an empty slice.
An ID in matching archive edits MUST remain queryable when it no longer exists
in current Specs.

### PRJ-056
Missing or malformed fireweed.yaml MUST produce an error beginning `C01: `.
Malformed archived change documents or malformed Added/Modified item syntax
MUST produce an error beginning `C03: ` and identifying the archive path.

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
An unknown selector MUST return U02 naming the requested change and listing valid pending change slugs; an ambiguous slug MUST return C13.

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

### PRJ-069
When `changes` is a regular file, Queue MUST return an error identifying `changes`. Check MUST report a non-warning C03 diagnostic identifying `changes`.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.


## Retained interfaces and files

packages/project/SPEC.md
packages/project/api.go
packages/project/model/model.go

## Retained evals

packages/project/evals/.gitkeep
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/selector_guidance_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go

Rebuild only the target's disposable paths. Keep all retained inputs intact and do not read outside the input workspace. Do not commit.

Run the target's eval hook from the workspace root as often as needed.

~~~sh
mkdir -p .fireweed/runs/drill && sh -c 'go run .fireweed/hooks.go eval '"'"'packages/project'"'"' '"'"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/.fireweed/runs/drill/report.xml'"'"''
~~~

Write DRILL_NOTES.md with separate ## Contradicted Decisions and ## Gaps sections. Every gap must name the nearest spec ID or "no spec" and a target module.


# Previous verification failed
app eval: E01: failed case "Test_APP_019_NewChangeErrorsPreserveFiles" for APP-019
project stdout: ?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.611s

project stderr: 

app stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	301.205s

app stderr: exit status 1

app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses: === RUN   Test_APP_037_StaticCauses
--- FAIL: Test_APP_037_StaticCauses (1.25s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/configuration: === RUN   Test_APP_037_StaticCauses/configuration
    diagnostic_causes_test.go:105: duplicate C01 prefix in "C01: C01: fireweed.yaml: yaml: line 1: did not find expected ',' or ']'. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.\n"
--- FAIL: Test_APP_037_StaticCauses/configuration (0.29s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/graph: === RUN   Test_APP_037_StaticCauses/graph
    diagnostic_causes_test.go:105: missing C02 cause "missing module" in plain output "C02: module \"beta\" depends on unknown module \"missing\". Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.\n"
    diagnostic_causes_test.go:105: missing C02 cause "missing module" in {"schema":1,"data":null,"diagnostics":[{"code":"C02","message":"module \"beta\" depends on unknown module \"missing\". Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/graph (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/queue-identity: === RUN   Test_APP_037_StaticCauses/queue-identity
    diagnostic_causes_test.go:105: missing C13 cause "ambiguous change slug" in plain output "C13: changes/0002-repeat.md: duplicate change slug \"repeat\" at changes/0001-repeat.md and changes/0002-repeat.md\n"
    diagnostic_causes_test.go:105: missing C13 cause "ambiguous change slug" in {"schema":1,"data":null,"diagnostics":[{"code":"C13","message":"changes/0002-repeat.md: duplicate change slug \"repeat\" at changes/0001-repeat.md and changes/0002-repeat.md","path":"changes/0002-repeat.md","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/queue-identity (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses: === RUN   Test_APP_038_LookupCauses
--- FAIL: Test_APP_038_LookupCauses (0.66s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses/stamp-change: === RUN   Test_APP_038_LookupCauses/stamp-change
    diagnostic_causes_test.go:156: duplicate U02 prefix in "U02: U02: change \"absent\" not found; no pending changes\n"
--- FAIL: Test_APP_038_LookupCauses/stamp-change (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses/coverage-change: === RUN   Test_APP_038_LookupCauses/coverage-change
    diagnostic_causes_test.go:156: duplicate U02 prefix in "U02: U02: change \"missing\" not found; no pending changes. Use a valid queued change alias.\n"
--- FAIL: Test_APP_038_LookupCauses/coverage-change (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors: === RUN   Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors
    list_cli_test.go:193: public Queue fixture cause: changes/0002-repeat.md: duplicate change slug "repeat" at changes/0001-repeat.md and changes/0002-repeat.md
--- FAIL: Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors (0.00s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles (0.69s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain
    project_commands_test.go:397: diagnostic="P01: change slug \"repeat\" already exists at changes/0003-repeat.md. Use a valid module and an unused lowercase slug.\n"; want corrective C13
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain
    project_commands_test.go:397: diagnostic="P01: change slug \"repeat\" already exists at changes/archive/0003-repeat.md. Use a valid module and an unused lowercase slug.\n"; want corrective C13
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON (0.01s)

checks stdout: ?   	github.com/lurkshark/fireweed/packages/checks	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/internal/engine	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/checks/evals	256.617s

checks stderr: 

runner stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	294.385s

runner stderr: 

workflow stdout: ?   	github.com/lurkshark/fireweed/packages/workflow	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/workflow/evals	135.768s

workflow stderr: 

codex
I’ll read the rebuild guardrails, inspect the project implementation, and trace the reported app failures before changing any disposable files.
exec
/bin/zsh -lc 'pwd; git status --short; ls -la; ls -la .fireweed; cat SYSTEM.md; cat docs/rebuild/PROGRESS.md; cat docs/rebuild/DECISIONS.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 455ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
?? DRILL_NOTES.md
?? packages/project/internal/
total 288
drwx------@   18 joel  staff     576 Nov 10 08:51 .
drwx------@ 6450 joel  staff  206400 Nov 10 09:09 ..
drwxr-xr-x@    9 joel  staff     288 Nov 10 08:49 .fireweed
drwxr-xr-x@    9 joel  staff     288 Nov 10 08:44 .git
-rw-r--r--@    1 joel  staff      56 Nov 10 08:44 .gitignore
-rw-r--r--@    1 joel  staff     859 Nov 10 08:44 AGENTS.md
-rw-r--r--@    1 joel  staff    1632 Nov 10 08:51 DRILL_NOTES.md
-rw-r--r--@    1 joel  staff   35149 Nov 10 08:44 LICENSE
-rw-r--r--@    1 joel  staff     100 Nov 10 08:44 Makefile
-rw-r--r--@    1 joel  staff    3399 Nov 10 08:44 README.md
-rw-r--r--@    1 joel  staff    8073 Nov 10 08:44 SYSTEM.md
drwxr-xr-x@    3 joel  staff      96 Nov 10 08:44 docs
-rw-r--r--@    1 joel  staff   61761 Nov 10 08:44 fireweed-design-v1.md
-rw-r--r--@    1 joel  staff    1289 Nov 10 08:44 fireweed.yaml
-rw-r--r--@    1 joel  staff      79 Nov 10 08:44 go.mod
-rw-r--r--@    1 joel  staff     151 Nov 10 08:44 go.sum
-rw-r--r--@    1 joel  staff     146 Nov 10 08:44 main.go
drwxr-xr-x@    6 joel  staff     192 Nov 10 08:44 packages
total 24
drwxr-xr-x@  9 joel  staff   288 Nov 10 08:49 .
drwx------@ 18 joel  staff   576 Nov 10 08:51 ..
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:44 bin
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:44 canaries
-rw-r--r--@  1 joel  staff  3993 Nov 10 08:44 hooks.go
-rw-r--r--@  1 joel  staff  1502 Nov 10 08:44 recipe.md
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:49 runs
-rw-r--r--@  1 joel  staff   470 Nov 10 08:44 setup.sh
drwxr-xr-x@  3 joel  staff    96 Nov 10 08:44 templates
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

The owner approved the `0197` worker and restricted audit contexts for OpenAI
Codex and granted standing authorization for the configured change lifecycle.
[SYSTEM.md](../../SYSTEM.md#agent-configuration) records the scope and destination.

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

**Project proof renewal awaits the PRJ-069 audit.** The verified controller ran
`fireweed drill project` at `d4e8d95`. All 133 project cases and 790 dependent
cases passed after the configured retry. One app case failed. The
[native report](../../.fireweed/drills/project.md) is committed in `dfbb207`.
Fireweed cleaned up both scratch locations and retained the reconstruction
`917fb02dba3b227f5c57ed422455e7e5881735c7` at
`refs/fireweed/drills/project/20261110T144729.916008000`. Its changes are
confined to disposable project paths.

`Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors` requires a C03
error from `project.Check` when `changes` is a regular file. The reconstruction
classified the queue read failure as C13. The same focused app eval passed
against the current product. The project queue and Check specs leave this
diagnostic implicit.

The owner approved making C03 explicit. Fireweed started
[0201-queue-path-diagnostic](../../changes/0201-queue-path-diagnostic.md)
and assigned `PRJ-069`. The project eval passed with the other 133 cases.
Audit found that the change also promises C03 on the public `Queue` error,
whose classification is private today. The eval asserts the path on `Queue`
and the C03 diagnostic on `Check`. App derives C03 from that diagnostic.

Both auditor submissions omitted the required dash before the finding's
explanation. The controller rejected them with P09 and blocked `0201` in
`d5ea3be`. It saved the audit work at
`refs/fireweed/attempts/queue-path-diagnostic/1` (`dc223bd`) and released the
claim. Audit rounds remain at zero.

The owner approved narrowing `PRJ-069` to the behavior the existing eval
asserts. `Queue` rejects a regular-file `changes` path and identifies it in
the error. `Check` reports a non-warning C03 diagnostic identifying the same
path. Author commit `fde499d` made the revision and removed `blocked` using
the documented resume procedure, preserving the assigned ID, run and audit
phase. `fireweed check` and the focused `PRJ-069` and `APP-116`/`APP-117` evals
passed. Automatic approval review rejected the lifecycle command before it
ran because it required specific authorization to send repository context to
OpenAI Codex. No claim is active. The saved attempt and failed drill ref remain
available.

The report lists no contradicted Decisions. Its other observations cover
diagnostic codes and wording relied on by app, duplicate C08 messages,
change numbering past 9999, intermediate graph validation, and stale-view
and history behavior. It also identifies the design's warning for specs
outside configured globs as an unassigned contract.

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

Reassess reuse if an exact historical proof is recovered from a backup.
Checks needs current proof of the `0197` CHK-043 change. `runner` and `app`
need first native reports and retained refs.

## Next step

Obtain explicit authorization for the configured Codex audit, implementation
and project reconstruction contexts. Resume `0201` at audit with the verified
controller, then renew `project` after Fireweed applies it. Commit the native
report, retain the proof ref and stop for owner review before `checks`. The
remaining drill order is `project`, `checks`, `runner`, then `app` in the
[regeneration goal](GOAL.md).

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

exec
/bin/zsh -lc 'rg --files packages/project packages/app packages/contracts .fireweed | sort; git status --short --untracked-files=all; git log -1 --oneline; cat AGENTS.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed playbook author-change; fireweed --help; ls -la .fireweed/runs/drill; cat fireweed.yaml; cat DRILL_NOTES.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 519ms:
# Author change
Read the design, SYSTEM.md and the module spec before writing a change. Give its Motivation a user-visible need or a clause of the design or SYSTEM.md; a probe showing a theoretical failure is insufficient. State each observable behavior at the module interface and keep the change focused. If a requirement costs more than it is worth, narrow or remove the spec and use already_satisfied for behavior the current implementation already meets. Do not commit during a worker step.
U01: Unknown option "--help". Use help to see available commands and options.
total 48
drwxr-xr-x@ 3 joel  staff     96 Nov 10 08:49 .
drwxr-xr-x@ 3 joel  staff     96 Nov 10 08:49 ..
-rw-r--r--@ 1 joel  staff  20635 Nov 10 08:52 report.xml
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
## Contradicted Decisions

None. Spec history uses the current archive files, returns section contents
without their headings, and returns only the matching ID for an ID-filtered
entry, as specified by 0126-spec-history-query.

## Gaps

- **PRJ-001 / project.** The design calls for a warning about SPEC.md files
  outside configured module globs. The project spec assigns no diagnostic for
  those files. Discovery stays within the configured globs.
- **PRJ-004 and PRJ-032 / project.** The specs leave open whether a malformed
  document after the selected change invalidates a partial projection. A
  partial projection validates queue identities across the queue and parses
  documents through the selected change.
- **PRJ-016 / project.** The retained eval requires two C08 diagnostics for an
  absent edit target, with generic and section-specific wording. The spec
  requires C08 and leaves the number and wording of diagnostics open.
- **PRJ-023 and PRJ-027 / project.** The filename format leaves numbering
  beyond 9999 implicit. The queue accepts at least four decimal digits so a
  change created after 9999 remains discoverable.
- **PRJ-035 / project.** The spec leaves validation timing open when a later
  change repairs a graph made invalid by an earlier change. Each projected
  graph is validated as its change is applied.
- **PRJ-060 and PRJ-061 / project.** The specs leave the stale result for a
  removed current item and the Git history traversal unspecified. An item with
  a recorded base and no current counterpart has empty Current text. Recovery
  follows HEAD's first-parent history at the queued path.

 succeeded in 532ms:
rg: packages/app: No such file or directory (os error 2)
.fireweed/canaries/canaries.yaml
.fireweed/hooks.go
.fireweed/recipe.md
.fireweed/setup.sh
.fireweed/templates/module/SPEC.md
.fireweed/templates/module/api.go.tmpl
packages/contracts/SPEC.md
packages/contracts/api.go
packages/contracts/evals/contracts_test.go
packages/contracts/internal/primitives/coded_error.go
packages/contracts/internal/primitives/primitives.go
packages/contracts/model/model.go
packages/project/SPEC.md
packages/project/api.go
packages/project/evals/coded_errors_test.go
packages/project/evals/mixed_eval_fix_test.go
packages/project/evals/project_test.go
packages/project/evals/queue_test.go
packages/project/evals/revision_snapshot_test.go
packages/project/evals/selector_guidance_test.go
packages/project/evals/spec_history_test.go
packages/project/evals/stale_test.go
packages/project/evals/static_codes_test.go
packages/project/internal/files/changes.go
packages/project/internal/files/source.go
packages/project/internal/fixture/graph.go
packages/project/internal/fixture/load.go
packages/project/internal/fixture/metadata.go
packages/project/internal/fixture/new.go
packages/project/internal/history/history.go
packages/project/internal/queue/bases.go
packages/project/internal/queue/check.go
packages/project/internal/queue/project.go
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go
packages/project/model/model.go
?? DRILL_NOTES.md
?? packages/project/internal/files/changes.go
?? packages/project/internal/files/source.go
?? packages/project/internal/fixture/graph.go
?? packages/project/internal/fixture/load.go
?? packages/project/internal/fixture/metadata.go
?? packages/project/internal/fixture/new.go
?? packages/project/internal/history/history.go
?? packages/project/internal/queue/bases.go
?? packages/project/internal/queue/check.go
?? packages/project/internal/queue/project.go
?? packages/project/internal/queue/queue.go
?? packages/project/internal/queue/stale.go
b024bda reconstruction inputs
<!-- fireweed:start -->
This project uses fireweed. Specs and change files are the source of truth.
Run `fireweed playbook author-change` before authoring changes and
`fireweed context <change>` before each step.
<!-- fireweed:end -->

## Rebuild controller

Drive the workflow with the controller `.fireweed/bin/fireweed` (see SYSTEM.md), not `bin/fireweed`, which is the product under test. Put `.fireweed/bin` first on worker PATH and use `GOCACHE=/tmp/fireweed-rebuild-cache`. Never switch controllers while a run or claim is active.

Use `fireweed verify` to check your work. Never commit during a worker step; fireweed owns commits. Read SYSTEM.md (including its scope guardrails) and docs/rebuild/PROGRESS.md before work; docs/rebuild/DECISIONS.md retains design clarifications. Do not consult `rebuild-v1` or deleted code as an implementation oracle.

exec
/bin/zsh -lc 'cat packages/project/internal/fixture/{graph,load,new}.go; cat packages/project/internal/queue/{queue,check}.go; cat packages/contracts/internal/primitives/coded_error.go; cat packages/contracts/api.go; cat packages/project/api.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed help context; fireweed help verify; fireweed context 0201-queue-path-diagnostic; fireweed help drill' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc "rg -n 'func Test|unknown module|missing module|duplicate change slug|ambiguous change slug|C01|U02|C13|already exists' packages/project/evals packages/contracts/evals fireweed-design-v1.md | head -190; ls packages; git ls-files | wc -l; cat .fireweed/hooks.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 555ms:
package fixture

import (
	"fmt"
	"strings"

	"github.com/lurkshark/fireweed/packages/project/model"
)

func ValidateGraph(modules []model.Module) error {
	names := make(map[string]model.Module, len(modules))
	prefixes := make(map[string]string, len(modules))
	for _, module := range modules {
		switch {
		case strings.TrimSpace(module.Name) == "":
			return fmt.Errorf("empty module name in %q", module.Dir)
		case strings.TrimSpace(module.IDPrefix) == "":
			return fmt.Errorf("empty ID prefix for module %q", module.Name)
		}
		if _, present := names[module.Name]; present {
			return fmt.Errorf("duplicate module name %q", module.Name)
		}
		if other, present := prefixes[module.IDPrefix]; present {
			return fmt.Errorf("duplicate ID prefix %q for modules %q and %q", module.IDPrefix, other, module.Name)
		}
		names[module.Name] = module
		prefixes[module.IDPrefix] = module.Name
	}
	for _, module := range modules {
		for _, dep := range module.DependsOn {
			if dep == module.Name {
				return fmt.Errorf("module %q has a self dependency", module.Name)
			}
			if _, present := names[dep]; !present {
				return fmt.Errorf("module %q depends on unknown module %q", module.Name, dep)
			}
		}
	}
	active := make(map[string]bool)
	done := make(map[string]bool)
	var visit func(string, []string) error
	visit = func(name string, trail []string) error {
		if active[name] {
			return fmt.Errorf("module dependency cycle: %s", strings.Join(append(trail, name), " -> "))
		}
		if done[name] {
			return nil
		}
		active[name] = true
		for _, dep := range names[name].DependsOn {
			if err := visit(dep, append(trail, name)); err != nil {
				return err
			}
		}
		delete(active, name)
		done[name] = true
		return nil
	}
	for _, module := range modules {
		if err := visit(module.Name, nil); err != nil {
			return err
		}
	}
	return nil
}
package fixture

import (
	"bytes"
	"fmt"
	"io"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/files"
	"github.com/lurkshark/fireweed/packages/project/model"
	"gopkg.in/yaml.v3"
)

const configPath = "fireweed.yaml"

func Load(root string) (model.Snapshot, error) {
	absolute, err := filepath.Abs(root)
	if err != nil {
		return model.Snapshot{}, configError(err)
	}
	return LoadFrom(absolute, files.Worktree{Root: absolute})
}

func LoadAt(root, revision string) (model.Snapshot, error) {
	top, err := files.Git(root, "rev-parse", "--show-toplevel")
	if err != nil {
		return model.Snapshot{}, err
	}
	canonical, err := filepath.EvalSymlinks(strings.TrimSuffix(top, "\n"))
	if err != nil {
		return model.Snapshot{}, err
	}
	source, err := files.OpenRevision(canonical, revision)
	if err != nil {
		return model.Snapshot{}, err
	}
	return LoadFrom(canonical, source)
}

func configError(err error) error {
	return files.Code("C01", configPath, fmt.Errorf("C01: %w", files.PathError{Path: configPath, Err: err}))
}

func LoadFrom(root string, source files.Source) (model.Snapshot, error) {
	data, err := source.ReadFile(configPath)
	if err != nil {
		return model.Snapshot{}, configError(err)
	}
	var config map[string]any
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(&config); err != nil {
		return model.Snapshot{}, configError(err)
	}
	if err := decoder.Decode(new(any)); err != io.EOF {
		if err == nil {
			err = fmt.Errorf("expected one YAML document")
		}
		return model.Snapshot{}, configError(err)
	}
	patterns, err := stringList(config["modules"])
	if err != nil {
		return model.Snapshot{}, configError(fmt.Errorf("modules: %w", err))
	}
	for i, pattern := range patterns {
		clean := path.Clean(pattern)
		if pattern == "" || path.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, "../") {
			return model.Snapshot{}, configError(fmt.Errorf("invalid module glob %q", pattern))
		}
		patterns[i] = clean
	}
	defaults, err := moduleDefaults(config)
	if err != nil {
		return model.Snapshot{}, files.At(configPath, "C12", err)
	}
	dirs, err := moduleDirs(source, patterns)
	if err != nil {
		return model.Snapshot{}, configError(err)
	}
	snapshot := model.Snapshot{Root: root, Config: config}
	for _, dir := range dirs {
		name := path.Join(dir, "SPEC.md")
		doc, err := files.ReadDocument(source, name)
		if err != nil {
			return model.Snapshot{}, err
		}
		if _, err := files.Items(doc, "Specs"); err != nil {
			return model.Snapshot{}, files.Syntax(name, err)
		}
		module, err := DecodeModule(dir, doc, defaults)
		if err != nil {
			return model.Snapshot{}, files.At(name, "C12", err)
		}
		snapshot.Modules = append(snapshot.Modules, module)
	}
	return snapshot, nil
}

func moduleDirs(source files.Source, globs []string) ([]string, error) {
	// Walk only branches that can match a configured glob. A standalone **
	// can either match no directory or stay active while descending.
	matches := make(map[string]bool)
	for _, glob := range globs {
		parts := strings.Split(glob, "/")
		if glob == "." {
			parts = nil
		}
		visited := make(map[struct {
			dir   string
			index int
		}]bool)
		var walk func(string, int) error
		walk = func(dir string, index int) error {
			state := struct {
				dir   string
				index int
			}{dir, index}
			if visited[state] {
				return nil
			}
			visited[state] = true
			if index == len(parts) {
				matches[dir] = true
				return nil
			}
			part := parts[index]
			if part == "**" {
				if err := walk(dir, index+1); err != nil {
					return err
				}
			}
			children, err := source.ReadDir(dir)
			if err != nil {
				return files.PathError{Path: dir, Err: err}
			}
			for _, child := range children {
				if !child.Mode.IsDir() || child.Name == ".git" {
					continue
				}
				next := index + 1
				if part == "**" {
					next = index
				} else if !contracts.MatchGlob(part, child.Name) {
					continue
				}
				if err := walk(path.Join(dir, child.Name), next); err != nil {
					return err
				}
			}
			return nil
		}
		if err := walk(".", 0); err != nil {
			return nil, err
		}
	}
	dirs := make([]string, 0, len(matches))
	for dir := range matches {
		dirs = append(dirs, dir)
	}
	slices.Sort(dirs)
	return dirs, nil
}
package fixture

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/files"
)

const skeleton = "\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"

func NewChange(root, module, slug string) (string, error) {
	snapshot, err := Load(root)
	if err != nil {
		return "", err
	}
	if ModuleIndex(snapshot.Modules, module) < 0 {
		return "", fmt.Errorf("unknown module %q", module)
	}
	if !files.Slug.MatchString(slug) {
		return "", fmt.Errorf("invalid change slug %q: use lowercase letters, digits and hyphens", slug)
	}
	source := files.Worktree{Root: snapshot.Root}
	maximum := 0
	for _, dir := range []string{files.Pending, files.Archive} {
		changes, err := files.Changes(source, dir)
		if err != nil {
			if dir == files.Pending {
				return "", fmt.Errorf("create changes/%04d-%s.md: %w", maximum+1, slug, err)
			}
			return "", err
		}
		for _, change := range changes {
			if change.Slug == slug {
				return "", fmt.Errorf("change slug %q already exists at %s", slug, change.Path)
			}
			maximum = max(maximum, change.Number)
		}
	}
	if maximum == math.MaxInt {
		return "", fmt.Errorf("cannot number change %q after %d", slug, maximum)
	}
	relative := fmt.Sprintf("changes/%04d-%s.md", maximum+1, slug)
	name := filepath.Join(snapshot.Root, filepath.FromSlash(relative))
	content, err := contracts.RenderDocument(contracts.Document{
		Frontmatter: map[string]any{"module": module},
		Body:        skeleton,
	})
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Dir(name), 0755); err != nil {
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", relative, err)
	}
	_, writeErr := file.WriteString(content)
	closeErr := file.Close()
	if writeErr == nil {
		writeErr = closeErr
	}
	if writeErr != nil {
		_ = os.Remove(name)
		return "", fmt.Errorf("create %s: %w", relative, writeErr)
	}
	return relative, nil
}
package queue

import (
	"fmt"
	"path"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/files"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func Queue(root string) ([]model.Change, error) {
	_, changes, err := load(root, nil)
	return changes, err
}

func load(root string, through *string) (model.Snapshot, []model.Change, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return model.Snapshot{}, nil, err
	}
	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules, through)
	return snapshot, changes, err
}

// A nil selector reads the whole queue. For a selected prefix, validate all
// queue identities but only parse documents through the selected change.
func read(source files.Source, modules []model.Module, through *string) ([]model.Change, error) {
	pending, err := files.Changes(source, files.Pending)
	if err != nil {
		return nil, err
	}
	slugs := make(map[string]string)
	for _, entry := range pending {
		if other, present := slugs[entry.Slug]; present {
			return nil, files.At(entry.Path, "C13", fmt.Errorf("duplicate change slug %q at %s and %s", entry.Slug, other, entry.Path))
		}
		slugs[entry.Slug] = entry.Path
	}
	if through != nil {
		index, err := resolve(pending, *through)
		if err != nil {
			return nil, err
		}
		pending = pending[:index+1]
	}
	changes := make([]model.Change, 0, len(pending))
	for _, entry := range pending {
		doc, err := files.ReadDocument(source, entry.Path)
		if err != nil {
			return nil, err
		}
		module, _ := doc.Frontmatter["module"].(string)
		if module == "" || fixture.ModuleIndex(modules, module) < 0 {
			return nil, files.At(entry.Path, "C07", fmt.Errorf("invalid pending change module %v", doc.Frontmatter["module"]))
		}
		change := model.Change{Number: entry.Number, Slug: entry.Slug, Path: entry.Path, Module: module, Document: doc}
		for _, section := range []struct {
			name  string
			items *[]contracts.Item
		}{{"Added", &change.Added}, {"Modified", &change.Modified}, {"Removed", &change.Removed}} {
			items, err := files.Items(doc, section.name)
			if err != nil {
				return nil, files.Syntax(entry.Path, err)
			}
			*section.items = items
		}
		seen := make(map[string]bool)
		for _, items := range [][]contracts.Item{change.Added, change.Modified, change.Removed} {
			for _, item := range items {
				if seen[item.ID] {
					return nil, files.At(entry.Path, "C05", fmt.Errorf("duplicate edit identity %s", item.ID))
				}
				seen[item.ID] = true
			}
		}
		changes = append(changes, change)
	}
	return changes, nil
}

func resolve(changes []files.ChangeFile, selector string) (int, error) {
	slugs := make([]string, 0, len(changes))
	for i, change := range changes {
		if selector == change.Slug || selector == path.Base(change.Path) || path.Clean(selector) == change.Path {
			return i, nil
		}
		slugs = append(slugs, change.Slug)
	}
	guidance := "no pending changes"
	if len(slugs) != 0 {
		guidance = "valid pending change slugs: " + strings.Join(slugs, ", ")
	}
	return -1, files.Code("U02", selector, fmt.Errorf("U02: change %q not found; %s", selector, guidance))
}
package queue

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/files"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var normative = regexp.MustCompile(`\b(?:MUST|SHOULD)\b`)
var handleReference = regexp.MustCompile(`new:[a-z0-9-]+`)

func Check(root string) []model.Diagnostic {
	snapshot, err := fixture.Load(root)
	if err != nil {
		return []model.Diagnostic{diagnostic(err)}
	}
	p, err := newProjection(snapshot)
	if err != nil {
		return []model.Diagnostic{diagnostic(err)}
	}
	var diagnostics []model.Diagnostic
	for i, module := range snapshot.Modules {
		for _, item := range p.items[i] {
			name := modulePath(module)
			if !owns(module, item.ID) {
				diagnostics = append(diagnostics, report(name, "C05", "foreign spec ID "+item.ID))
			}
			diagnostics = append(diagnostics, checkText(name, item)...)
		}
	}
	changes, err := read(files.Worktree{Root: snapshot.Root}, snapshot.Modules, nil)
	if err != nil {
		return append(diagnostics, diagnostic(err))
	}
	for _, change := range changes {
		i := fixture.ModuleIndex(p.snapshot.Modules, change.Module)
		diagnostics = append(diagnostics, checkChange(change, p.snapshot.Modules[i])...)
		diagnostics = append(diagnostics, p.checkBases(change)...)
		if err := p.apply(change); err != nil {
			diagnostics = append(diagnostics, diagnostic(err))
			break
		}
	}
	return diagnostics
}

func report(path, code, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Path: path, Message: message}
}

func diagnostic(err error) model.Diagnostic {
	code, subject, _ := contracts.ErrorCode(err)
	var at files.PathError
	if errors.As(err, &at) {
		subject = at.Path
	}
	return report(subject, code, err.Error())
}

func owns(module model.Module, id string) bool {
	prefix, _, found := strings.Cut(id, "-")
	return found && prefix == module.IDPrefix
}

func checkText(path string, item contracts.Item) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	if !normative.MatchString(item.Body) {
		diagnostics = append(diagnostics, report(path, "C04", "non-normative spec body for "+item.ID))
	}
	if !validTag(item.Tag) {
		diagnostics = append(diagnostics, report(path, "C06", "invalid review/quarantine tag for "+item.ID))
	}
	return diagnostics
}

func validTag(tag string) bool {
	if tag == "" || tag == "review" {
		return true
	}
	reason, ok := strings.CutPrefix(tag, "quarantine:")
	return ok && strings.TrimSpace(reason) != ""
}

func checkChange(change model.Change, module model.Module) []model.Diagnostic {
	var diagnostics []model.Diagnostic
	name := change.Path
	motivation, _ := contracts.Section(change.Document.Body, "Motivation")
	if strings.TrimSpace(motivation) == "" {
		diagnostics = append(diagnostics, report(name, "C07", "missing Motivation"))
	}
	_, evalFix := change.Document.Frontmatter["eval_fix"]
	_, fixes := change.Document.Frontmatter["fixes"]
	if evalFix && (fixes || len(change.Added)+len(change.Modified)+len(change.Removed) > 0) {
		diagnostics = append(diagnostics, report(name, "C07", "eval_fix cannot be combined with fixes or spec edits"))
	}

	declared := make(map[string]bool)
	for _, item := range change.Added {
		if isHandle(item.ID) {
			declared[item.ID] = true
		}
	}
	for _, items := range [][]contracts.Item{change.Added, change.Modified, change.Removed} {
		for _, item := range items {
			if !isHandle(item.ID) && !owns(module, item.ID) {
				diagnostics = append(diagnostics, report(name, "C05", "foreign module ID "+item.ID))
			}
		}
	}
	for _, items := range [][]contracts.Item{change.Added, change.Modified} {
		for _, item := range items {
			diagnostics = append(diagnostics, checkText(name, item)...)
		}
	}
	for _, item := range change.Removed {
		if !validTag(item.Tag) {
			diagnostics = append(diagnostics, report(name, "C06", "invalid review/quarantine tag for "+item.ID))
		}
	}

	// Rendering exposes frontmatter references as well as references in prose.
	text, err := contracts.RenderDocument(change.Document)
	if err != nil {
		diagnostics = append(diagnostics, diagnostic(files.At(name, "C03", err)))
	}
	for _, handle := range handleReference.FindAllString(text, -1) {
		if !declared[handle] {
			diagnostics = append(diagnostics, report(name, "C09", "reference to undeclared handle "+handle))
			declared[handle] = true
		}
	}
	return diagnostics
}

func (p projection) checkBases(change model.Change) []model.Diagnostic {
	mapping, err := bases(change)
	if err != nil {
		return []model.Diagnostic{diagnostic(err)}
	}
	items := p.stableItems(change.Module)
	var diagnostics []model.Diagnostic
	for _, edit := range baseEdits(change) {
		item, present := items[edit.ID]
		hash := recordedHash(mapping[edit.ID])
		if !present {
			message := fmt.Sprintf("ID %s absent from preceding projection", edit.ID)
			diagnostics = append(diagnostics, report(change.Path, "C08", message))
		}
		var message string
		switch {
		case hash == "":
			message = "missing base for " + edit.ID
		case present && hash != contracts.SpecHash(item):
			message = "stale base for " + edit.ID
		}
		if message != "" {
			diagnostics = append(diagnostics, report(change.Path, "C08", message))
		}
	}
	return diagnostics
}
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

 succeeded in 562ms:
packages/project/evals/mixed_eval_fix_test.go:10:func Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits(t *testing.T) {
packages/project/evals/stale_test.go:62:func Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder(t *testing.T) {
packages/project/evals/stale_test.go:88:func Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash(t *testing.T) {
packages/project/evals/stale_test.go:135:func Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem(t *testing.T) {
packages/project/evals/stale_test.go:149:func Test_PRJ_063_StaleSpecsEmptyAndReadOnly(t *testing.T) {
packages/project/evals/stale_test.go:182:func Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent(t *testing.T) {
packages/project/evals/stale_test.go:221:func Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically(t *testing.T) {
packages/project/evals/stale_test.go:243:func Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures(t *testing.T) {
packages/project/evals/stale_test.go:272:			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
packages/project/evals/stale_test.go:273:				t.Errorf("unknown selector = %v; want U02 with subject %q", err, target)
packages/project/evals/stale_test.go:288:			if code, _, ok := contracts.ErrorCode(err); !ok || code != "C13" {
packages/project/evals/stale_test.go:289:				t.Errorf("ambiguous slug = %v; want C13", err)
packages/project/evals/spec_history_test.go:61:func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
packages/project/evals/spec_history_test.go:89:func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
packages/project/evals/spec_history_test.go:117:func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
packages/project/evals/spec_history_test.go:139:func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
packages/project/evals/spec_history_test.go:173:func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
packages/project/evals/spec_history_test.go:189:		if err == nil || !strings.HasPrefix(err.Error(), "U02: ") || !strings.Contains(err.Error(), tc.rejected) {
packages/project/evals/spec_history_test.go:190:			t.Errorf("lookup module %q ID %q = %v; want U02: naming %q", tc.module, tc.id, err, tc.rejected)
packages/project/evals/spec_history_test.go:195:func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
packages/project/evals/spec_history_test.go:210:			if err == nil || !strings.HasPrefix(err.Error(), "C01: ") {
packages/project/evals/spec_history_test.go:211:				t.Errorf("invalid config = %v; want C01: prefix", err)
packages/contracts/evals/contracts_test.go:15:func Test_CON_001_ParseDocument(t *testing.T) {
packages/contracts/evals/contracts_test.go:29:func Test_CON_007_RejectInvalidFrontmatter(t *testing.T) {
packages/contracts/evals/contracts_test.go:43:func Test_CON_008_RenderRoundTrip(t *testing.T) {
packages/contracts/evals/contracts_test.go:61:func Test_CON_008_YAMLRepresentationsAndNilFrontmatter(t *testing.T) {
packages/contracts/evals/contracts_test.go:91:func Test_CON_002_CON_009_Sections(t *testing.T) {
packages/contracts/evals/contracts_test.go:105:func Test_CON_009_CON_002_AppendSection(t *testing.T) {
packages/contracts/evals/contracts_test.go:116:func Test_CON_002_CON_009_FirstDuplicateSection(t *testing.T) {
packages/contracts/evals/contracts_test.go:127:func Test_CON_010_ParseItems(t *testing.T) {
packages/contracts/evals/contracts_test.go:142:func Test_CON_011_CON_012_RejectInvalidItems(t *testing.T) {
packages/contracts/evals/contracts_test.go:155:func Test_CON_006_CON_010_ParseHandlesInSourceOrder(t *testing.T) {
packages/contracts/evals/contracts_test.go:171:func Test_CON_010_AssignedIDGrammar(t *testing.T) {
packages/contracts/evals/contracts_test.go:199:func Test_CON_015_CON_012_RejectInvalidOrDuplicateHandles(t *testing.T) {
packages/contracts/evals/contracts_test.go:215:func Test_CON_013_RenderItem(t *testing.T) {
packages/contracts/evals/contracts_test.go:229:func Test_CON_003_CanonicalHash(t *testing.T) {
packages/contracts/evals/contracts_test.go:242:func Test_CON_003_IDAndTagAffectHash(t *testing.T) {
packages/contracts/evals/contracts_test.go:254:func Test_CON_004_PositiveMatches(t *testing.T) {
packages/contracts/evals/contracts_test.go:267:func Test_CON_004_Boundaries(t *testing.T) {
packages/contracts/evals/contracts_test.go:279:func Test_CON_005_MatchGlob(t *testing.T) {
packages/contracts/evals/contracts_test.go:297:func Test_CON_014_ResolveZone(t *testing.T) {
packages/contracts/evals/contracts_test.go:319:func Test_CON_016_ErrorCode(t *testing.T) {
packages/contracts/evals/contracts_test.go:348:func Test_CON_017_ParseItemsDuplicateCode(t *testing.T) {
packages/contracts/evals/contracts_test.go:372:func Test_CON_018_ParseItemsMalformedCode(t *testing.T) {
packages/contracts/evals/contracts_test.go:396:func Test_CON_019_ParseDocumentFrontmatterCode(t *testing.T) {
packages/project/evals/project_test.go:47:func Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults(t *testing.T) {
packages/project/evals/project_test.go:88:func Test_PRJ_018_LoadInvalidDocuments(t *testing.T) {
packages/project/evals/project_test.go:118:func Test_PRJ_002_Graph(t *testing.T) {
packages/project/evals/project_test.go:156:func Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering(t *testing.T) {
packages/project/evals/project_test.go:193:func Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
packages/project/evals/project_test.go:215:func Test_PRJ_025_PRJ_026_NewChangeInvalidArguments(t *testing.T) {
packages/project/evals/project_test.go:228:			} else if tc.id == "PRJ-025" && (!strings.HasPrefix(err.Error(), "unknown module ") || !strings.Contains(err.Error(), tc.module)) {
packages/project/evals/project_test.go:229:				t.Errorf("unknown module error = %q; want prefix and rejected module %q", err, tc.module)
packages/project/evals/project_test.go:235:func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
packages/project/evals/coded_errors_test.go:10:func Test_PRJ_057_StampInvalidBasesHasCode(t *testing.T) {
packages/project/evals/coded_errors_test.go:22:func Test_PRJ_058_MissingPendingChangeHasSubject(t *testing.T) {
packages/project/evals/coded_errors_test.go:37:			if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != target {
packages/project/evals/coded_errors_test.go:38:				t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, target, err)
packages/project/evals/coded_errors_test.go:44:func Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:62:func Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:98:func Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:127:func Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:149:func Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:158:	missingSpec := revisionCommit(t, root, "missing module document")
packages/project/evals/static_codes_test.go:12:func Test_PRJ_040_CheckConfigurationCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:34:				if !diagnostic.Warning && diagnostic.Code == "C01" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:38:			t.Errorf("Check did not report C01 for configuration cause %q", cause)
packages/project/evals/static_codes_test.go:43:func Test_PRJ_041_CheckGraphCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:78:func Test_PRJ_042_CheckDocumentAndItemCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:120:func Test_PRJ_043_CheckMetadataAndSetCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:163:func Test_PRJ_044_CheckQueueIdentityCause(t *testing.T) {
packages/project/evals/static_codes_test.go:172:		if !diagnostic.Warning && diagnostic.Code == "C13" && strings.Contains(diagnostic.Message, cause.Error()) {
packages/project/evals/static_codes_test.go:176:	t.Errorf("Check did not report C13 for queue identity cause %q", cause)
packages/project/evals/static_codes_test.go:179:func Test_PRJ_069_RegularFileAtChangesReportsQueueCause(t *testing.T) {
packages/project/evals/static_codes_test.go:198:func Test_PRJ_045_CheckPendingModuleTargetCause(t *testing.T) {
packages/project/evals/static_codes_test.go:213:func Test_PRJ_013_CheckDuplicateWithinEditSection(t *testing.T) {
packages/project/evals/queue_test.go:54:func Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive(t *testing.T) {
packages/project/evals/queue_test.go:89:func Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments(t *testing.T) {
packages/project/evals/queue_test.go:104:func Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug(t *testing.T) {
packages/project/evals/queue_test.go:144:func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
packages/project/evals/queue_test.go:172:func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
packages/project/evals/queue_test.go:214:func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
packages/project/evals/queue_test.go:256:func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
packages/project/evals/queue_test.go:280:func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
packages/project/evals/queue_test.go:292:func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
packages/project/evals/selector_guidance_test.go:17:	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
packages/project/evals/selector_guidance_test.go:18:		t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
packages/project/evals/selector_guidance_test.go:22:		t.Errorf("lookup error = %v; want a retained U02 cause", err)
packages/project/evals/selector_guidance_test.go:39:	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
packages/project/evals/selector_guidance_test.go:40:		t.Fatalf("empty queue error = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
packages/project/evals/selector_guidance_test.go:51:func Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs(t *testing.T) {
packages/project/evals/selector_guidance_test.go:73:func Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs(t *testing.T) {
assets
contracts
gitstore
project
      86
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

 exited 2 in 601ms:
U01: help takes no arguments. Use help alone.
U01: help takes no arguments. Use help alone.
U02: change "0201-queue-path-diagnostic" not found; no valid change slugs. Use a change held by this worktree and retry context.
U01: help takes no arguments. Use help alone.

exec
/bin/zsh -lc 'cat packages/project/internal/files/{changes,source}.go; cat packages/project/internal/history/history.go; cat packages/contracts/model/model.go; cat packages/project/evals/{selector_guidance,static_codes}_test.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed help; git ls-files packages | cut -d/ -f2 | sort -u; rg -n '"'diagnostic|unknown module|ambiguous change slug|new change|C13|U02' fireweed-design-v1.md | head -100" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 520ms:
package files

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"sort"
	"strconv"
)

const Pending = "changes"
const Archive = "changes/archive"

var Slug = regexp.MustCompile(`^[a-z0-9-]+$`)
var changeName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

type ChangeFile struct {
	Number int
	Slug   string
	Path   string
}

func Changes(source Source, dir string) ([]ChangeFile, error) {
	entries, err := source.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, At(dir, "C03", err)
	}
	var result []ChangeFile
	for _, entry := range entries {
		match := changeName.FindStringSubmatch(entry.Name)
		if match == nil || entry.Mode.IsDir() {
			continue
		}
		name := path.Join(dir, entry.Name)
		if !entry.Mode.IsRegular() {
			return nil, At(name, "C03", fmt.Errorf("change is not a regular file"))
		}
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, At(name, "C13", fmt.Errorf("invalid change number: %w", err))
		}
		result = append(result, ChangeFile{number, match[2], name})
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if a.Number != b.Number {
			return a.Number < b.Number
		}
		if a.Slug != b.Slug {
			return a.Slug < b.Slug
		}
		return a.Path < b.Path
	})
	return result, nil
}
// Package files reads project documents from the worktree or a recorded Git tree.
package files

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
)

type Source interface {
	ReadFile(string) ([]byte, error)
	ReadDir(string) ([]Entry, error)
}

type Entry struct {
	Name string
	Mode fs.FileMode
}

type Worktree struct{ Root string }

func (w Worktree) ReadFile(name string) ([]byte, error) {
	return fs.ReadFile(os.DirFS(w.Root), name)
}

func (w Worktree) ReadDir(name string) ([]Entry, error) {
	entries, err := fs.ReadDir(os.DirFS(w.Root), name)
	if err != nil {
		return nil, err
	}
	result := make([]Entry, 0, len(entries))
	for _, entry := range entries {
		result = append(result, Entry{entry.Name(), entry.Type()})
	}
	return result, nil
}

type Revision struct {
	Root   string
	commit string
	dirs   map[string][]Entry
	modes  map[string]fs.FileMode
}

func Git(root string, args ...string) (string, error) {
	result, err := gitstore.Run(context.Background(), gitstore.Request{
		Dir:     root,
		Args:    append([]string{"git", "--no-optional-locks"}, args...),
		Env:     map[string]string{"GIT_NO_LAZY_FETCH": "1", "GIT_TERMINAL_PROMPT": "0"},
		Timeout: 30 * time.Second,
	})
	if err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(result.Stderr))
	}
	return result.Stdout, nil
}

func OpenRevision(root, revision string) (*Revision, error) {
	output, err := Git(root, "rev-parse", "--verify", "--end-of-options", revision+"^{commit}")
	if err != nil {
		return nil, err
	}
	commit := strings.TrimSpace(output)
	output, err = Git(root, "ls-tree", "--full-tree", "-r", "-t", "-z", commit)
	if err != nil {
		return nil, err
	}
	tree := &Revision{
		Root: root, commit: commit,
		dirs:  map[string][]Entry{".": nil},
		modes: make(map[string]fs.FileMode),
	}
	for _, record := range strings.Split(strings.TrimSuffix(output, "\x00"), "\x00") {
		if record == "" {
			continue
		}
		header, name, _ := strings.Cut(record, "\t")
		mode, _, _ := strings.Cut(header, " ")
		var kind fs.FileMode
		switch mode {
		case "040000":
			kind = fs.ModeDir
			tree.dirs[name] = nil
		case "120000":
			kind = fs.ModeSymlink
		case "100644", "100755":
		default:
			kind = fs.ModeIrregular
		}
		parent := path.Dir(name)
		tree.dirs[parent] = append(tree.dirs[parent], Entry{path.Base(name), kind})
		tree.modes[name] = kind
	}
	return tree, nil
}

func (r *Revision) ReadFile(name string) ([]byte, error) {
	mode, exists := r.modes[name]
	if !exists {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrNotExist}
	}
	if !mode.IsRegular() {
		return nil, &fs.PathError{Op: "read", Path: name, Err: fs.ErrInvalid}
	}
	return gitstore.ReadAt(r.Root, r.commit, name)
}

func (r *Revision) ReadDir(name string) ([]Entry, error) {
	entries, exists := r.dirs[name]
	if !exists {
		cause := fs.ErrNotExist
		if _, present := r.modes[name]; present {
			cause = fs.ErrInvalid
		}
		return nil, &fs.PathError{Op: "readdir", Path: name, Err: cause}
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name < entries[j].Name })
	return entries, nil
}

type PathError struct {
	Path string
	Err  error
}

func (e PathError) Error() string { return e.Path + ": " + e.Err.Error() }
func (e PathError) Unwrap() error { return e.Err }

func Code(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
}

func At(name, code string, err error) error {
	return PathError{Path: name, Err: Code(code, name, err)}
}

// Syntax preserves the parser's classification and subject.
func Syntax(name string, err error) error {
	return PathError{Path: name, Err: err}
}

func ReadDocument(source Source, name string) (contracts.Document, error) {
	data, err := source.ReadFile(name)
	if err != nil {
		return contracts.Document{}, At(name, "C03", err)
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return contracts.Document{}, Syntax(name, err)
	}
	return doc, nil
}

func Items(doc contracts.Document, section string) ([]contracts.Item, error) {
	body, _ := contracts.Section(doc.Body, section)
	return contracts.ParseItems(body)
}
package history

import (
	"errors"
	"fmt"
	"io/fs"
	"path"
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/files"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
	snapshot, err := fixture.Load(root)
	if err != nil {
		if code, subject, _ := contracts.ErrorCode(err); code == "C12" && subject == "fireweed.yaml" {
			return nil, files.Code("C01", subject, fmt.Errorf("C01: %w", err))
		}
		return nil, err
	}
	index := fixture.ModuleIndex(snapshot.Modules, module)
	if index < 0 {
		return nil, unknown("module", module)
	}
	source := files.Worktree{Root: snapshot.Root}
	archive, err := source.ReadDir(files.Archive)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, archiveError(files.At(files.Archive, "C03", err))
	}

	entries := make([]model.SpecHistoryEntry, 0)
	dated := make(map[string]bool)
	for _, file := range archive {
		if file.Mode.IsDir() || !strings.HasSuffix(file.Name, ".md") {
			continue
		}
		name := path.Join(files.Archive, file.Name)
		doc, err := files.ReadDocument(source, name)
		if err != nil {
			return nil, archiveError(err)
		}
		target, ok := doc.Frontmatter["module"].(string)
		if !ok || target == "" {
			return nil, archiveError(files.At(name, "C03", fmt.Errorf("archive module must be a nonempty string")))
		}
		ids, err := affectedIDs(doc, id)
		if err != nil {
			return nil, archiveError(files.Syntax(name, err))
		}
		if target != module || len(ids) == 0 {
			continue
		}
		entry := model.SpecHistoryEntry{Path: name, Module: target, IDs: ids}
		entry.Motivation, _ = contracts.Section(doc.Body, "Motivation")
		entry.Decisions, _ = contracts.Section(doc.Body, "Decisions")
		entry.RemovedDecisions, _ = contracts.Section(doc.Body, "Removed decisions")
		entry.Provenance, _ = contracts.Section(doc.Body, "Provenance")
		entry.ImplementationDate, dated[name] = implementationDate(entry.Provenance)
		entries = append(entries, entry)
	}
	if id != "" && len(entries) == 0 {
		items, err := files.Items(snapshot.Modules[index].Document, "Specs")
		if err != nil {
			return nil, err
		}
		if !slices.ContainsFunc(items, func(item contracts.Item) bool { return item.ID == id }) {
			return nil, unknown("spec ID", id)
		}
	}
	sort.Slice(entries, func(i, j int) bool {
		a, b := entries[i], entries[j]
		if dated[a.Path] != dated[b.Path] {
			return dated[a.Path]
		}
		if a.ImplementationDate.Equal(b.ImplementationDate) {
			return a.Path < b.Path
		}
		return a.ImplementationDate.After(b.ImplementationDate)
	})
	return entries, nil
}

func affectedIDs(doc contracts.Document, id string) ([]string, error) {
	seen := make(map[string]bool)
	for _, section := range []string{"Added", "Modified"} {
		items, err := files.Items(doc, section)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			if id == "" || item.ID == id {
				seen[item.ID] = true
			}
		}
	}
	ids := make([]string, 0, len(seen))
	for affected := range seen {
		ids = append(ids, affected)
	}
	slices.Sort(ids)
	return ids, nil
}

func implementationDate(provenance string) (time.Time, bool) {
	var latest time.Time
	dated := false
	for _, line := range strings.Split(provenance, "\n") {
		fields := strings.Fields(line)
		if len(fields) < 3 || fields[0] != "-" || fields[1] != "impl" {
			continue
		}
		date, err := time.Parse(time.RFC3339, fields[2])
		if err == nil && (!dated || date.After(latest)) {
			latest = date
			dated = true
		}
	}
	return latest, dated
}

func archiveError(err error) error {
	return files.Code("C03", files.Archive, fmt.Errorf("C03: %w", err))
}

func unknown(kind, name string) error {
	return files.Code("U02", name, fmt.Errorf("U02: unknown %s %q", kind, name))
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
package evals_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project"
)

func pendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if err == nil {
		t.Fatal("unknown pending selector was accepted")
	}
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Errorf("error code = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	var coded contracts.CodedError
	if !errors.As(err, &coded) || coded.Cause == nil {
		t.Errorf("lookup error = %v; want a retained U02 cause", err)
	}
	message := err.Error()
	for _, want := range []string{selector, "first", "second"} {
		if !strings.Contains(message, want) {
			t.Errorf("lookup error %q lacks rejected selector or pending slug %q", message, want)
		}
	}
	for _, unwanted := range []string{"0001-first.md", "0002-second.md", "archived"} {
		if strings.Contains(message, unwanted) {
			t.Errorf("lookup guidance %q includes queued filename or archived slug %q", message, unwanted)
		}
	}
}

func emptyPendingSelectorGuidance(t *testing.T, err error, selector string) {
	t.Helper()
	if code, subject, ok := contracts.ErrorCode(err); !ok || code != "U02" || subject != selector {
		t.Fatalf("empty queue error = (%q, %q, %v); want (U02, %q, true): %v", code, subject, ok, selector, err)
	}
	message := strings.ToLower(err.Error())
	if !strings.Contains(message, selector) {
		t.Errorf("empty queue error %q lacks rejected selector", message)
	}
	if !strings.Contains(message, "no pending") && !strings.Contains(message, "no valid pending") && !strings.Contains(message, "none") && !strings.Contains(message, "empty") && !strings.Contains(message, "[]") {
		t.Errorf("empty queue error %q lacks empty-set guidance", message)
	}
}

func Test_PRJ_004_UnknownProjectAndStampSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"Project", "missing", func(selector string) error { _, err := project.Project(root, selector); return err }},
		{"Stamp", "0099-absent.md", func(selector string) error { return project.Stamp(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.Project(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
}

func Test_PRJ_066_UnknownStaleSelectorsListPendingSlugs(t *testing.T) {
	root := fixture(t)
	queued(t, root, "0001-first.md", "", "First.\n", "", "", "")
	queued(t, root, "0002-second.md", "", "Second.\n", "", "", "")
	put(t, root, "changes/archive/0003-archived.md", "archived material")
	for _, operation := range []struct {
		name, selector string
		call           func(string) error
	}{
		{"StaleSpecs", "missing", func(selector string) error { _, err := project.StaleSpecs(root, selector); return err }},
		{"StampMissing", "changes/0099-absent.md", func(selector string) error { return project.StampMissing(root, selector) }},
	} {
		t.Run(operation.name, func(t *testing.T) {
			pendingSelectorGuidance(t, operation.call(operation.selector), operation.selector)
		})
	}

	emptyRoot := fixture(t)
	_, err := project.StaleSpecs(emptyRoot, "missing")
	emptyPendingSelectorGuidance(t, err, "missing")
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

func Test_PRJ_069_RegularFileAtChangesReportsQueueCause(t *testing.T) {
	root := fixture(t)
	put(t, root, "packages/alpha/SPEC.md", "---\nname: alpha\nid_prefix: ALP\ndepends_on: []\n---\n# Alpha\n\n## Specs\n\n### ALP-001\nAlpha MUST work.\n")
	const queuePath = "changes"
	put(t, root, queuePath, "not a directory")

	_, cause := project.Queue(root)
	if cause == nil || !strings.Contains(cause.Error(), queuePath) {
		t.Fatalf("Queue error = %v; want an error identifying %q", cause, queuePath)
	}
	diagnostics := project.Check(root)
	for _, diagnostic := range diagnostics {
		if !diagnostic.Warning && diagnostic.Code == "C03" && diagnostic.Path == queuePath && strings.Contains(diagnostic.Message, cause.Error()) {
			return
		}
	}
	t.Errorf("Check diagnostics = %#v; want a C03 error at %q for Queue cause %q", diagnostics, queuePath, cause)
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

 succeeded in 527ms:
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
assets
contracts
gitstore
project
28:- **Easy for agents to drive.** Every command is non-interactive, has machine-readable output, and has stable diagnostic codes.
261:**Unapproved work is bounded.** An agent working directly on the local base branch could approve its own changes by committing them, so this bound counts against the shared base: an unapproved change is an archived change on HEAD whose change file is not on `origin/<base>` (or on the local `<base>` when no remote ref exists). Pushing or merging is the human review. When HEAD already has `authoring.max_unapproved` of them (default 5), `start`, `next` and `run` refuse another unmerged change, even with `--allow-unmerged`, with diagnostic P07; merging the work into the base branch, which is a human review, resets the count. This holds whoever authored the changes, so an agent that both writes and runs changes still stops for a person at regular intervals. `status` shows the count, and warns (P08) when three or more consecutive applied changes target the same module, which is the usual shape of a regression chain where each change fixes the edge case the last one exposed.
436:| Reading | `explain <code>`, `playbook <name>` | What a diagnostic means and how to fix it; a playbook's text |
451:- Stable diagnostic codes (for example C10, P2-01, H03), each with `fireweed explain <code>`.

exec
/bin/zsh -lc "sed -n '105,156p' packages/project/evals/project_test.go; sed -n '100,145p' packages/project/evals/queue_test.go; rg -n '\"C01\"|U02:|C01:' packages/project/internal; ls -l .fireweed/bin; git status --short --ignored | head -45; cat go.mod; cat .fireweed/setup.sh" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 490ms:
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
packages/project/internal/queue/queue.go:96:	return -1, files.Code("U02", selector, fmt.Errorf("U02: change %q not found; %s", selector, guidance))
packages/project/internal/fixture/load.go:45:	return files.Code("C01", configPath, fmt.Errorf("C01: %w", files.PathError{Path: configPath, Err: err}))
packages/project/internal/history/history.go:23:			return nil, files.Code("C01", subject, fmt.Errorf("C01: %w", err))
packages/project/internal/history/history.go:132:	return files.Code("U02", name, fmt.Errorf("U02: unknown %s %q", kind, name))
total 13512
-rwxr-xr-x@ 1 joel  staff  6914850 Nov 10 08:44 fireweed
?? DRILL_NOTES.md
?? packages/project/internal/
!! .fireweed/bin/
!! .fireweed/runs/
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
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed explain C13; fireweed verify; rg -n '"'new-change|duplicate slug|ambiguous|configuration' fireweed-design-v1.md | head -75; sed -n '190,265p' packages/project/evals/project_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 544ms:
A duplicate queue slug or an unparseable numeric queue prefix makes the queue identity invalid; tied numeric positions are allowed. Choose an unused slug with a valid numeric prefix.
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.
202:**Body sections.** Motivation is required. Interface describes interface changes in prose; the evals phase turns them into code, which keeps change files independent of language. Added and Modified contain complete spec items; Removed contains headings only. Decisions is filled in by the worker. Questions is written by an agent that blocks the change because the spec is ambiguous. Removed decisions quotes module Decisions to drop. Provenance is written by fireweed at each step. Rejected alternatives and Notes are free prose.
303:- **The agent blocks the change** because the spec is ambiguous or contradicts existing code. It writes a `## Questions` section and a `blocked` reason. The author answers by editing the change and removing `blocked`.
430:| Authoring | `new-change <module> <slug>` | Create a change file with the next number |
453:- Distinct exit codes: 0 for success or an empty queue, 1 for failed checks or any change blocked during the run, 2 for usage or configuration errors, 3 for a clean early stop on a time or usage limit.
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
			_, err := project.NewChange(root, tc.module, tc.slug)
			if err == nil {
				t.Errorf("accepted module %q and slug %q", tc.module, tc.slug)
			} else if tc.id == "PRJ-025" && (!strings.HasPrefix(err.Error(), "unknown module ") || !strings.Contains(err.Error(), tc.module)) {
				t.Errorf("unknown module error = %q; want prefix and rejected module %q", err, tc.module)
			}
		})
	}
}

func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
	root := fixture(t)
	const queuedPath = "changes/0001-blocked.md"
	if err := os.MkdirAll(filepath.Join(root, filepath.FromSlash(queuedPath)), 0755); err != nil {
		t.Fatal(err)
	}
	path, err := project.NewChange(root, "alpha", "blocked")
	if err == nil {
		t.Fatal("NewChange accepted a queued path occupied by a directory")
	}
	if path != "" {
		t.Errorf("rejected NewChange path = %q; want empty", path)
	}
	if !strings.HasPrefix(err.Error(), "create "+queuedPath+":") {
		t.Errorf("NewChange error = %q; want repository-relative create prefix", err)
	}
	var pathErr *os.PathError
	if !errors.As(err, &pathErr) || pathErr.Path != filepath.Join(root, filepath.FromSlash(queuedPath)) || pathErr.Err == nil {
		t.Errorf("NewChange error = %v; want retained filesystem cause for %s", err, queuedPath)
	}
}

codex
The reported failures point to four project error paths. I’ll separate diagnostic codes from cause text where the CLI adds its own prefix, align the graph and queue messages with the app expectations, and return C13 for duplicate slugs. This workspace contains the project suite and its dependencies, so I can verify those public behaviors locally.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/fixture/graph.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/fixture/load.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/fixture/new.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/history/history.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/graph.go b/packages/project/internal/fixture/graph.go
index 27073bab6c53a0743906371713b2a121a425a88d..1b1c741b33afffed0bcf3d63bf4ef65a878e6cae
--- a/packages/project/internal/fixture/graph.go
+++ b/packages/project/internal/fixture/graph.go
@@ -32,7 +32,7 @@
 				return fmt.Errorf("module %q has a self dependency", module.Name)
 			}
 			if _, present := names[dep]; !present {
-				return fmt.Errorf("module %q depends on unknown module %q", module.Name, dep)
+				return fmt.Errorf("module %q depends on missing module %q", module.Name, dep)
 			}
 		}
 	}
diff --git a/packages/project/internal/fixture/load.go b/packages/project/internal/fixture/load.go
index 6f404bd75bc9abaf8492dc0aaa6ae2139668ea0a..b282b4dcbef7366d61de6c313db25c69d3fe67e4
--- a/packages/project/internal/fixture/load.go
+++ b/packages/project/internal/fixture/load.go
@@ -42,7 +42,7 @@
 }
 
 func configError(err error) error {
-	return files.Code("C01", configPath, fmt.Errorf("C01: %w", files.PathError{Path: configPath, Err: err}))
+	return files.At(configPath, "C01", err)
 }
 
 func LoadFrom(root string, source files.Source) (model.Snapshot, error) {
diff --git a/packages/project/internal/fixture/new.go b/packages/project/internal/fixture/new.go
index 2905b629a863834636f01e3643da3b200c176392..37a7454e5becfaadc60413a63dad8a1b8f707010
--- a/packages/project/internal/fixture/new.go
+++ b/packages/project/internal/fixture/new.go
@@ -35,7 +35,7 @@
 		}
 		for _, change := range changes {
 			if change.Slug == slug {
-				return "", fmt.Errorf("change slug %q already exists at %s", slug, change.Path)
+				return "", files.Code("C13", slug, fmt.Errorf("duplicate change slug %q: already exists at %s", slug, change.Path))
 			}
 			maximum = max(maximum, change.Number)
 		}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
index fb4085c766fd1020d4fe967e2311913629aee700..42562647fceed5b848cb1f38258845b7f0e3111e
--- a/packages/project/internal/history/history.go
+++ b/packages/project/internal/history/history.go
@@ -19,7 +19,7 @@
 func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
 	snapshot, err := fixture.Load(root)
 	if err != nil {
-		if code, subject, _ := contracts.ErrorCode(err); code == "C12" && subject == "fireweed.yaml" {
+		if code, subject, _ := contracts.ErrorCode(err); code == "C01" || (code == "C12" && subject == "fireweed.yaml") {
 			return nil, files.Code("C01", subject, fmt.Errorf("C01: %w", err))
 		}
 		return nil, err
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 5499d3540781d979958633a7f211ce63fe5c2a03..48e5296f7a4fcb410d6d23102ce0331962a4e0fc
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -35,7 +35,7 @@
 	slugs := make(map[string]string)
 	for _, entry := range pending {
 		if other, present := slugs[entry.Slug]; present {
-			return nil, files.At(entry.Path, "C13", fmt.Errorf("duplicate change slug %q at %s and %s", entry.Slug, other, entry.Path))
+			return nil, files.At(entry.Path, "C13", fmt.Errorf("ambiguous change slug %q at %s and %s", entry.Slug, other, entry.Path))
 		}
 		slugs[entry.Slug] = entry.Path
 	}
@@ -93,5 +93,5 @@
 	if len(slugs) != 0 {
 		guidance = "valid pending change slugs: " + strings.Join(slugs, ", ")
 	}
-	return -1, files.Code("U02", selector, fmt.Errorf("U02: change %q not found; %s", selector, guidance))
+	return -1, files.Code("U02", selector, fmt.Errorf("change %q not found; %s", selector, guidance))
 }

exec
/bin/zsh -lc 'gofmt -l packages/project/internal/fixture/{graph,load,new}.go packages/project/internal/history/history.go packages/project/internal/queue/queue.go; git diff --check; git status --short --untracked-files=all' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; mkdir -p .fireweed/runs/drill && sh -c '"'go run .fireweed/hooks.go eval '\"'\"'packages/project'\"'\"' '\"'\"'/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/.fireweed/runs/drill/report.xml'\"'\"''" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 560ms:
?? DRILL_NOTES.md
?? packages/project/internal/files/changes.go
?? packages/project/internal/files/source.go
?? packages/project/internal/fixture/graph.go
?? packages/project/internal/fixture/load.go
?? packages/project/internal/fixture/metadata.go
?? packages/project/internal/fixture/new.go
?? packages/project/internal/history/history.go
?? packages/project/internal/queue/bases.go
?? packages/project/internal/queue/check.go
?? packages/project/internal/queue/project.go
?? packages/project/internal/queue/queue.go
?? packages/project/internal/queue/stale.go

 succeeded in 3997ms:
?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/files	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.395s

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index f4618d12284e8a7eb873b8edc01bb280fa3b16ef..58546e605b84bef2ec1c7ab1b72fb53b01a5b5e3
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -13,15 +13,30 @@
   document after the selected change invalidates a partial projection. A
   partial projection validates queue identities across the queue and parses
   documents through the selected change.
+- **PRJ-004, PRJ-058 and PRJ-066 / project.** APP-038 expects the CLI to
+  render U02 once. Pending selector errors carry U02 as metadata and provide
+  the rejected selector and queue guidance in the cause text.
 - **PRJ-016 / project.** The retained eval requires two C08 diagnostics for an
   absent edit target, with generic and section-specific wording. The spec
   requires C08 and leaves the number and wording of diagnostics open.
+- **PRJ-021 and PRJ-041 / project.** APP-037 expects the phrase `missing module`
+  for an absent dependency. The project specs require an explanatory cause
+  and C02 without prescribing the phrase.
 - **PRJ-023 and PRJ-027 / project.** The filename format leaves numbering
   beyond 9999 implicit. The queue accepts at least four decimal digits so a
   change created after 9999 remains discoverable.
+- **PRJ-024 / project.** APP-019 expects C13 for duplicate slugs in the queue
+  and archive. The project spec requires rejection and file preservation and
+  leaves the code open.
+- **PRJ-031 and PRJ-044 / project.** APP-037 and the APP-116/APP-117 eval
+  expect the phrase `ambiguous change slug`. The project specs require
+  rejection and C13 and leave the wording open.
 - **PRJ-035 / project.** The spec leaves validation timing open when a later
   change repairs a graph made invalid by an earlier change. Each projected
   graph is validated as its change is applied.
+- **PRJ-040 and PRJ-056 / project.** APP-037 expects the CLI to render C01
+  once. Load exposes C01 as error metadata and Check carries it separately
+  from the message. SpecHistory adds the textual prefix required by PRJ-056.
 - **PRJ-060 and PRJ-061 / project.** The specs leave the stale result for a
   removed current item and the Git history traversal unspecified. An item with
   a recorded base and no current counterpart has empty Current text. Recovery
diff --git a/packages/project/internal/fixture/graph.go b/packages/project/internal/fixture/graph.go
index 27073bab6c53a0743906371713b2a121a425a88d..1b1c741b33afffed0bcf3d63bf4ef65a878e6cae
--- a/packages/project/internal/fixture/graph.go
+++ b/packages/project/internal/fixture/graph.go
@@ -32,7 +32,7 @@
 				return fmt.Errorf("module %q has a self dependency", module.Name)
 			}
 			if _, present := names[dep]; !present {
-				return fmt.Errorf("module %q depends on unknown module %q", module.Name, dep)
+				return fmt.Errorf("module %q depends on missing module %q", module.Name, dep)
 			}
 		}
 	}
diff --git a/packages/project/internal/fixture/load.go b/packages/project/internal/fixture/load.go
index 6f404bd75bc9abaf8492dc0aaa6ae2139668ea0a..b282b4dcbef7366d61de6c313db25c69d3fe67e4
--- a/packages/project/internal/fixture/load.go
+++ b/packages/project/internal/fixture/load.go
@@ -42,7 +42,7 @@
 }
 
 func configError(err error) error {
-	return files.Code("C01", configPath, fmt.Errorf("C01: %w", files.PathError{Path: configPath, Err: err}))
+	return files.At(configPath, "C01", err)
 }
 
 func LoadFrom(root string, source files.Source) (model.Snapshot, error) {
diff --git a/packages/project/internal/fixture/new.go b/packages/project/internal/fixture/new.go
index 2905b629a863834636f01e3643da3b200c176392..37a7454e5becfaadc60413a63dad8a1b8f707010
--- a/packages/project/internal/fixture/new.go
+++ b/packages/project/internal/fixture/new.go
@@ -35,7 +35,7 @@
 		}
 		for _, change := range changes {
 			if change.Slug == slug {
-				return "", fmt.Errorf("change slug %q already exists at %s", slug, change.Path)
+				return "", files.Code("C13", slug, fmt.Errorf("duplicate change slug %q: already exists at %s", slug, change.Path))
 			}
 			maximum = max(maximum, change.Number)
 		}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
index fb4085c766fd1020d4fe967e2311913629aee700..42562647fceed5b848cb1f38258845b7f0e3111e
--- a/packages/project/internal/history/history.go
+++ b/packages/project/internal/history/history.go
@@ -19,7 +19,7 @@
 func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
 	snapshot, err := fixture.Load(root)
 	if err != nil {
-		if code, subject, _ := contracts.ErrorCode(err); code == "C12" && subject == "fireweed.yaml" {
+		if code, subject, _ := contracts.ErrorCode(err); code == "C01" || (code == "C12" && subject == "fireweed.yaml") {
 			return nil, files.Code("C01", subject, fmt.Errorf("C01: %w", err))
 		}
 		return nil, err
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index 5499d3540781d979958633a7f211ce63fe5c2a03..48e5296f7a4fcb410d6d23102ce0331962a4e0fc
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -35,7 +35,7 @@
 	slugs := make(map[string]string)
 	for _, entry := range pending {
 		if other, present := slugs[entry.Slug]; present {
-			return nil, files.At(entry.Path, "C13", fmt.Errorf("duplicate change slug %q at %s and %s", entry.Slug, other, entry.Path))
+			return nil, files.At(entry.Path, "C13", fmt.Errorf("ambiguous change slug %q at %s and %s", entry.Slug, other, entry.Path))
 		}
 		slugs[entry.Slug] = entry.Path
 	}
@@ -93,5 +93,5 @@
 	if len(slugs) != 0 {
 		guidance = "valid pending change slugs: " + strings.Join(slugs, ", ")
 	}
-	return -1, files.Code("U02", selector, fmt.Errorf("U02: change %q not found; %s", selector, guidance))
+	return -1, files.Code("U02", selector, fmt.Errorf("change %q not found; %s", selector, guidance))
 }

exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache; fireweed verify; git diff --check; git diff --exit-code -- packages/project/SPEC.md packages/project/api.go packages/project/model packages/project/evals; git status --short --untracked-files=all' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 564ms:
P02: worktree holds no change. Use the held change and correct its checks, then retry verify.
?? DRILL_NOTES.md
?? packages/project/internal/files/changes.go
?? packages/project/internal/files/source.go
?? packages/project/internal/fixture/graph.go
?? packages/project/internal/fixture/load.go
?? packages/project/internal/fixture/metadata.go
?? packages/project/internal/fixture/new.go
?? packages/project/internal/history/history.go
?? packages/project/internal/queue/bases.go
?? packages/project/internal/queue/check.go
?? packages/project/internal/queue/project.go
?? packages/project/internal/queue/queue.go
?? packages/project/internal/queue/stale.go

2026-11-10T17:11:53.218230Z ERROR codex_core::tools::router: error=exec_command failed: CreateProcess { message: "Rejected(\"`/bin/zsh -lc 'set -eu\\nexport PATH=\\\"$PWD/.fireweed/bin:$PATH\\\" GOCACHE=/tmp/fireweed-rebuild-cache\\nprobe_dir=$(mktemp -d \\\"$PWD/.fireweed/runs/drill/diagnostics.XXXXXX\\\")\\ntrap '\\\"'rm -rf \\\\\\\"\\\"'$probe_dir\\\"'\\\"' EXIT\\ncat > \\\\\\\"\\\"'$probe_dir/probe.go\\\" <<'\\\"'EOF'\\npackage main\\n\\nimport (\\n\\t\\\\\\\"fmt\\\\\\\"\\n\\t\\\\\\\"os\\\\\\\"\\n\\t\\\\\\\"…4052 chars truncated…\\\".fireweed/runs/drill/report.xml\\\\\\\")\\ncases = report.findall(\\\\\\\".//testcase\\\\\\\")\\nfailures = sum(case.find(\\\\\\\"failure\\\\\\\") is not None for case in cases)\\nskips = sum(case.find(\\\\\\\"skipped\\\\\\\") is not None for case in cases)\\nprint(f\\\\\\\"Project report: {len(cases)} cases, {failures} failures, {skips} skipped\\\\\\\")\\nPY\\n\\\"` rejected: rm -f style commands are not permitted. Use a safer approach\")" }
exec
/bin/zsh -lc 'export PATH="$PWD/.fireweed/bin:$PATH" GOCACHE=/tmp/fireweed-rebuild-cache
python3 - <<'"'PY'
import pathlib
import subprocess
import tempfile
import xml.etree.ElementTree as ET

drill = pathlib.Path(\".fireweed/runs/drill\")
cases = ET.parse(drill / \"report.xml\").findall(\".//testcase\")
failures = sum(case.find(\"failure\") is not None for case in cases)
skips = sum(case.find(\"skipped\") is not None for case in cases)
print(f\"Project report: {len(cases)} cases, {failures} failures, {skips} skipped\", flush=True)

with tempfile.TemporaryDirectory(prefix=\"diagnostics-\", dir=drill) as scratch:
    path = pathlib.Path(scratch) / \"probe.go\"
    path.write_text(r'''
package main

import (
	\"fmt\"
	\"os\"
	\"path/filepath\"
	\"strings\"

	\"github.com/lurkshark/fireweed/packages/contracts\"
	\"github.com/lurkshark/fireweed/packages/project\"
)

func expect(ok bool, message string, values ...any) {
	if "'!ok {
		panic(fmt.Sprintf(message, values...))
	}
}

func write(root, name, contents string) {
	target := filepath.Join(root, name)
	if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(target, []byte(contents), 0644); err != nil {
		panic(err)
	}
}

func checkCode(err error, want string) {
	code, _, _ := contracts.ErrorCode(err)
	expect(code == want, "wanted %s, got %q from %v", want, code, err)
}

func main() {
	root := filepath.Join(os.Args[1], "fixture")
	write(root, "fireweed.yaml", "modules: [broken'"\\n\")
	_, err := project.Load(root)
	checkCode(err, \"C01\")
	expect(strings.HasPrefix(err.Error(), \"fireweed.yaml:\"), \"Load configuration error: %v\", err)
	diagnostics := project.Check(root)
	expect(len(diagnostics) == 1 && diagnostics[0].Code == \"C01\" && diagnostics[0].Message == err.Error(), \"configuration diagnostics: %+v\", diagnostics)
	_, err = project.SpecHistory(root, \"alpha\", \"\")
	expect(err "'!= nil && strings.HasPrefix(err.Error(), "C01: fireweed.yaml:"), "history configuration error: %v", err)

	write(root, "fireweed.yaml", "modules: [packages/*]'"\\n\")
	const module = \"---\\nname: alpha\\nid_prefix: ALP\\n%s---\\n## Specs\\n\"
	write(root, \"packages/alpha/SPEC.md\", fmt.Sprintf(module, \"depends_on: [absent]\\n\"))
	diagnostics = project.Check(root)
	expect(len(diagnostics) == 1 && diagnostics[0].Code == \"C02\" && strings.Contains(diagnostics[0].Message, \"missing module\"), \"graph diagnostics: %+v\", diagnostics)
	write(root, \"packages/alpha/SPEC.md\", fmt.Sprintf(module, \"\"))

	err = project.Stamp(root, \"absent\")
	checkCode(err, \"U02\")
	expect(strings.Contains(err.Error(), \"absent\") && strings.Contains(err.Error(), \"no pending changes\") && "'!strings.Contains(err.Error(), "U02:"), "selector error: %v", err)

	const change = "---'"\\nmodule: alpha\\n---\\n## Motivation\\nRequired.\\n\"
	for _, name := range []string{\"changes/0001-repeat.md\", \"changes/archive/0002-archived.md\"} {
		write(root, name, change)
		slug := strings.TrimSuffix(filepath.Base(name)[5:], \".md\")
		_, err := project.NewChange(root, \"alpha\", slug)
		checkCode(err, \"C13\")
		expect(strings.Contains(err.Error(), name), \"duplicate slug error: %v\", err)
		data, err := os.ReadFile(filepath.Join(root, name))
		expect(err == nil && string(data) == change, \"duplicate slug altered %s\", name)
	}
	write(root, \"changes/0003-repeat.md\", change)
	_, err = project.Queue(root)
	checkCode(err, \"C13\")
	expect(strings.Contains(err.Error(), \"ambiguous change slug\"), \"queue error: %v\", err)
	diagnostics = project.Check(root)
	expect(len(diagnostics) == 1 && diagnostics[0].Code == \"C13\" && diagnostics[0].Message == err.Error(), \"queue diagnostics: %+v\", diagnostics)
	fmt.Println(\"Public diagnostic checks passed\")
}
''')
    subprocess.run([\"go\", \"run\", str(path), scratch], check=True)
PY" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908
 succeeded in 1742ms:
Project report: 134 cases, 0 failures, 0 skipped
Public diagnostic checks passed

codex
Updated the project rebuild to return C13 for duplicate slugs, match the graph and queue wording expected by app, and supply C01 and U02 separately from the CLI message text. Spec history retains its required textual prefixes.

All 134 project eval cases passed, along with focused checks of the public diagnostic behavior. The app suite is absent from this workspace, so the downstream failures could not be rerun here. `fireweed verify` returned P02 because the worktree holds no change.

Changes are confined to project disposable paths and [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-4230050908/DRILL_NOTES.md).
tokens used
72,699

