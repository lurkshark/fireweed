---
date: 2026-11-02T03:16:47.104747Z
base_commit: 3f95497a164b045bf374da4a1498011bb0bc14c2
agent: codex
model: gpt-6-sol
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
ok  	github.com/lurkshark/fireweed/packages/project/evals	2.033s


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
ok  	github.com/lurkshark/fireweed/packages/app/evals	230.929s

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
ok  	github.com/lurkshark/fireweed/packages/checks/evals	227.181s

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
ok  	github.com/lurkshark/fireweed/packages/runner/evals	246.224s

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
PASS
ok  	github.com/lurkshark/fireweed/packages/workflow/evals	103.464s


## Contradicted Decisions
None identified in the retained project inputs.
## Gaps
- no spec (project): The supplied input omits the Go module manifest and dependent implementations, so the project and transitive app evals cannot run locally. The separate native verifier must compile and run them.
## Agent observations

### Attempt 1 (exit 0)
Rebuilt the project implementation under [packages/project/internal](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal) and wrote [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/DRILL_NOTES.md). Retained inputs were not edited, and nothing was committed.

`gofmt` and `git diff --check` passed. The retained evals could not run here because this input workspace has no Go module manifest or dependency implementations; the native verifier must run them.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a19a42-71b3-7cf2-8772-03a11411f94e
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
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

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
I’ll inspect the retained project API, model, specs, and evals, then rebuild only the disposable project implementation. I’ll run the retained evals and write the requested drill notes without changing retained inputs.
exec
/bin/zsh -lc "pwd && rg --files -g 'AGENTS.md' -g 'packages/project/**' -g 'go.mod' -g 'fireweed.yaml' -g 'DRILL_NOTES.md' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 510ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
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
packages/project/model/model.go

exec
/bin/zsh -lc 'rg --files | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'cat packages/project/api.go packages/project/model/model.go packages/project/SPEC.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'cat packages/project/evals/*.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'find . -name AGENTS.md -print' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 728ms:
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
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

## Decisions

[0126-spec-history-query] History reads the project's current archive files, as the archive-backed show
view; it does not infer chronology from commit ancestry or filenames. Entry
metadata are Markdown section contents without their level-two heading; a
nonempty ID filter returns that matching ID in each entry's IDs field.

 succeeded in 731ms:
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

 succeeded in 733ms:
 succeeded in 739ms:
SYSTEM.md
packages/assets/api.go
packages/assets/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
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
packages/project/model/model.go

exec
/bin/zsh -lc "rg -n '"'^func Test|''^func fixture|''^func queued|''^const fixtureConfig'"' packages/project/evals/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'cat SYSTEM.md packages/contracts/api.go packages/contracts/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'cat packages/gitstore/api.go packages/gitstore/model/model.go packages/assets/api.go packages/assets/model/model.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'ls -la; ls -la packages/project; ls -la packages/project/internal 2>/dev/null; ls -la go.* 2>/dev/null' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 621ms:
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

 succeeded in 627ms:
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

 succeeded in 628ms:
packages/project/evals/mixed_eval_fix_test.go:10:func Test_PRJ_046_CheckEvalFixCannotMixWithFixesOrSpecEdits(t *testing.T) {
packages/project/evals/stale_test.go:62:func Test_PRJ_060_StaleSpecsSelectsRecordedModifiedAndRemovedInEditOrder(t *testing.T) {
packages/project/evals/stale_test.go:88:func Test_PRJ_061_PRJ_062_HistoryUsesLastBasesChangeAndValidatesEachHash(t *testing.T) {
packages/project/evals/stale_test.go:135:func Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem(t *testing.T) {
packages/project/evals/stale_test.go:149:func Test_PRJ_063_StaleSpecsEmptyAndReadOnly(t *testing.T) {
packages/project/evals/stale_test.go:182:func Test_PRJ_064_StampMissingPreservesExistingBasesAndHumanContent(t *testing.T) {
packages/project/evals/stale_test.go:221:func Test_PRJ_065_StampMissingRejectsInvalidBasesAndAbsentEditsAtomically(t *testing.T) {
packages/project/evals/stale_test.go:243:func Test_PRJ_066_StaleQueriesResolveAliasesAndCodedLookupFailures(t *testing.T) {
packages/project/evals/queue_test.go:14:func queued(t *testing.T, root, filename, front, motivation, added, modified, removed string) {
packages/project/evals/queue_test.go:54:func Test_PRJ_027_PRJ_028_QueueOrderContentAndArchive(t *testing.T) {
packages/project/evals/queue_test.go:89:func Test_PRJ_029_PRJ_030_QueueRejectsInvalidDocuments(t *testing.T) {
packages/project/evals/queue_test.go:104:func Test_PRJ_004_PRJ_031_PRJ_032_TargetAliasesAndAmbiguousSlug(t *testing.T) {
packages/project/evals/queue_test.go:144:func Test_PRJ_005_PRJ_032_PRJ_036_ProjectEditsAndLocalHandles(t *testing.T) {
packages/project/evals/queue_test.go:172:func Test_PRJ_033_PRJ_034_PRJ_035_SetFieldsAndRejections(t *testing.T) {
packages/project/evals/queue_test.go:214:func Test_PRJ_006_PRJ_038_StampPrecedingProjectionAndPreservation(t *testing.T) {
packages/project/evals/queue_test.go:256:func Test_PRJ_037_PRJ_039_StampAddedOnlyAndAbsentID(t *testing.T) {
packages/project/evals/queue_test.go:280:func Test_PRJ_007_CheckValidAndDiagnostics(t *testing.T) {
packages/project/evals/queue_test.go:292:func Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics(t *testing.T) {
packages/project/evals/project_test.go:16:const fixtureConfig = `fireweed: ">=0.1, <0.2"
packages/project/evals/project_test.go:38:func fixture(t *testing.T) string {
packages/project/evals/project_test.go:47:func Test_PRJ_001_PRJ_017_LoadConfiguredModulesAndDefaults(t *testing.T) {
packages/project/evals/project_test.go:88:func Test_PRJ_018_LoadInvalidDocuments(t *testing.T) {
packages/project/evals/project_test.go:118:func Test_PRJ_002_Graph(t *testing.T) {
packages/project/evals/project_test.go:156:func Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering(t *testing.T) {
packages/project/evals/project_test.go:193:func Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
packages/project/evals/project_test.go:215:func Test_PRJ_025_PRJ_026_NewChangeInvalidArguments(t *testing.T) {
packages/project/evals/project_test.go:235:func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
packages/project/evals/spec_history_test.go:61:func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
packages/project/evals/spec_history_test.go:89:func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
packages/project/evals/spec_history_test.go:117:func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
packages/project/evals/spec_history_test.go:139:func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
packages/project/evals/spec_history_test.go:173:func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
packages/project/evals/spec_history_test.go:195:func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
packages/project/evals/static_codes_test.go:12:func Test_PRJ_040_CheckConfigurationCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:43:func Test_PRJ_041_CheckGraphCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:78:func Test_PRJ_042_CheckDocumentAndItemCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:120:func Test_PRJ_043_CheckMetadataAndSetCauses(t *testing.T) {
packages/project/evals/static_codes_test.go:163:func Test_PRJ_044_CheckQueueIdentityCause(t *testing.T) {
packages/project/evals/static_codes_test.go:179:func Test_PRJ_045_CheckPendingModuleTargetCause(t *testing.T) {
packages/project/evals/static_codes_test.go:194:func Test_PRJ_013_CheckDuplicateWithinEditSection(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:62:func Test_PRJ_047_LoadAtRecordedConfigGlobsAndDocuments(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:98:func Test_PRJ_048_LoadAtRecordedDefaultsAndExplicitEmptyZones(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:127:func Test_PRJ_049_LoadAtCanonicalRootAndRelativeModuleDirs(t *testing.T) {
packages/project/evals/revision_snapshot_test.go:149:func Test_PRJ_050_LoadAtRejectsInvalidRecordedInputs(t *testing.T) {
packages/project/evals/coded_errors_test.go:10:func Test_PRJ_057_StampInvalidBasesHasCode(t *testing.T) {
packages/project/evals/coded_errors_test.go:22:func Test_PRJ_058_MissingPendingChangeHasSubject(t *testing.T) {
packages/project/evals/coded_errors_test.go:44:func Test_PRJ_059_PRJ_068_LoadPreservesModuleParseCodeAndCheckPath(t *testing.T) {

 exited 1 in 632ms:
total 8
drwx------@    5 joel  staff     160 Nov  1 19:16 .
drwx------@ 6898 joel  staff  220736 Nov  1 19:16 ..
drwxr-xr-x@    9 joel  staff     288 Nov  1 19:16 .git
-rw-r--r--@    1 joel  staff    2965 Nov  1 19:16 SYSTEM.md
drwxr-xr-x@    6 joel  staff     192 Nov  1 19:16 packages
total 32
drwxr-xr-x@  6 joel  staff    192 Nov  1 19:16 .
drwxr-xr-x@  6 joel  staff    192 Nov  1 19:16 ..
-rw-r--r--@  1 joel  staff  10696 Nov  1 19:16 SPEC.md
-rw-r--r--@  1 joel  staff   1468 Nov  1 19:16 api.go
drwxr-xr-x@ 11 joel  staff    352 Nov  1 19:16 evals
drwxr-xr-x@  3 joel  staff     96 Nov  1 19:16 model
zsh:1: no matches found: go.*

exec
/bin/zsh -lc "sed -n '1,280p' packages/project/evals/project_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "sed -n '1,340p' packages/project/evals/queue_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "sed -n '1,230p' packages/project/evals/spec_history_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "sed -n '1,195p' packages/project/evals/revision_snapshot_test.go; sed -n '1,100p' packages/project/evals/stale_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 649ms:
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

 succeeded in 651ms:
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

 succeeded in 654ms:
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

 succeeded in 649ms:
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

exec
/bin/zsh -lc "sed -n '205,340p' packages/project/evals/queue_test.go; sed -n '1,65p' packages/project/evals/spec_history_test.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 455ms:
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

codex
The retained inputs define the public shapes and expected diagnostics. The workspace intentionally lacks the dependency implementations and build manifest, so I’ll implement under `packages/project/internal` and use the available interfaces, then make the bounded checks this input permits.
exec
/bin/zsh -lc 'mkdir -p packages/project/internal/fixture packages/project/internal/queue packages/project/internal/history' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 543ms:
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..ff6aa39d0958d76095a9b51bd3fe3c54bf7630b6
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,149 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		result, e := gitstore.RunGitTree(base, revision)
+		_ = result; _ = e
+	}
+	_ = path.Clean
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..7d4345fb4638a24a682d0b4446af359e3702defa
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,152 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output(); if e != nil { return snap, fmt.Errorf("read revision %s tree: %w", revision, e) }
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") { if file != "" { parts := strings.Split(file, "/"); for i := 1; i < len(parts); i++ { allDirs[strings.Join(parts[:i], "/")] = true } } }
+		for dir := range allDirs { for _, glob := range globs { if contracts.MatchGlob(glob, dir) { dirs = append(dirs, dir); break } } }
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..7d4345fb4638a24a682d0b4446af359e3702defa
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,152 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output(); if e != nil { return snap, fmt.Errorf("read revision %s tree: %w", revision, e) }
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") { if file != "" { parts := strings.Split(file, "/"); for i := 1; i < len(parts); i++ { allDirs[strings.Join(parts[:i], "/")] = true } } }
+		for dir := range allDirs { for _, glob := range globs { if contracts.MatchGlob(glob, dir) { dirs = append(dirs, dir); break } } }
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..5628d1ff809f24aea163f136051bbffba248e334
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,145 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string); if !ok || module == "" { return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel)) }; c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		for _, item := range items { if seen[item.ID] { return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID)) }; seen[item.ID] = true }
+		switch field { case "Added": c.Added = items; case "Modified": c.Modified = items; case "Removed": c.Removed = items }
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	modules := map[string]bool{}; for _, m := range snap.Modules { modules[m.Name] = true }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes")); if os.IsNotExist(err) { return []model.Change{}, nil }; if err != nil { return nil, fixture.Code("C03", "changes", err) }
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() { continue }; match := pendingName.FindStringSubmatch(entry.Name()); if match == nil { continue }
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] { return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2])) }; slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data, rel, num, match[2]); if e != nil { return nil, e }
+		if !modules[c.Module] { return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module)) }
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i,j int) bool { if changes[i].Number != changes[j].Number { return changes[i].Number < changes[j].Number }; return changes[i].Slug < changes[j].Slug })
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match { if idx >= 0 { return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector)) }; idx = i }
+	}
+	if idx < 0 { return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector)) }
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m); if err != nil { return contracts.Item{}, false }
+	for _, item := range items { if item.ID == id { return item, true } }; return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root); if err != nil { return snap, err }
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return snap, fixture.Code("C02", "", err) }
+	for i := 0; i < target; i++ { if err := Apply(&snap, changes[i]); err != nil { return snap, err } }
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, err }
+	limit := len(changes)
+	if through != "" { idx, e := Resolve(changes, through); if e != nil { return model.Snapshot{}, e }; limit = idx+1 }
+	snap, err := Before(root, changes, limit); if err != nil { return snap, err }
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1; for n, m := range snap.Modules { if m.Name == change.Module { i = n; break } }; if i < 0 { return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module)) }
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any); if !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping")) }
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err)) }
+			case "boundaries_opt_out": if _, ok := v.(string); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			case "hooks": if _, ok := v.(map[string]any); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			default: return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults); if err != nil { return err }; m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m); if err != nil { return fixture.Code("C03", path.Join(m.Dir,"SPEC.md"), err) }
+	for _, edit := range change.Modified { found := false; for j := range items { if items[j].ID == edit.ID { items[j] = edit; found = true; break } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) } }
+	for _, edit := range change.Removed { found := false; next := items[:0]; for _, item := range items { if item.ID == edit.ID { found = true } else { next = append(next, item) } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }; items = next }
+	items = append(items, change.Added...)
+	var body strings.Builder; for _, item := range items { body.WriteString(contracts.RenderItem(item)); if !strings.HasSuffix(body.String(), "\n") { body.WriteByte('\n') }; if !strings.HasSuffix(body.String(), "\n\n") { body.WriteByte('\n') } }
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return fixture.Code("C02", change.Path, err) }
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	idx, err := Resolve(changes, selector); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	snap, err := Before(root, changes, idx); return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]; if !exists { return map[string]any{}, nil }
+	m, ok := v.(map[string]any); if !ok { return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping")) }
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector); if err != nil { return err }
+	b, err := bases(c.Document); if err != nil { return err }
+	var m model.Module; for _, module := range snap.Modules { if module.Name == c.Module { m = module; break } }
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID); if !ok { return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }
+		if _, exists := b[edit.ID]; missingOnly && exists { continue }
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document); if err != nil { return err }
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/history/history.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..7d4345fb4638a24a682d0b4446af359e3702defa
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,152 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output(); if e != nil { return snap, fmt.Errorf("read revision %s tree: %w", revision, e) }
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") { if file != "" { parts := strings.Split(file, "/"); for i := 1; i < len(parts); i++ { allDirs[strings.Join(parts[:i], "/")] = true } } }
+		for dir := range allDirs { for _, glob := range globs { if contracts.MatchGlob(glob, dir) { dirs = append(dirs, dir); break } } }
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..5628d1ff809f24aea163f136051bbffba248e334
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,145 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string); if !ok || module == "" { return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel)) }; c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		for _, item := range items { if seen[item.ID] { return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID)) }; seen[item.ID] = true }
+		switch field { case "Added": c.Added = items; case "Modified": c.Modified = items; case "Removed": c.Removed = items }
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	modules := map[string]bool{}; for _, m := range snap.Modules { modules[m.Name] = true }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes")); if os.IsNotExist(err) { return []model.Change{}, nil }; if err != nil { return nil, fixture.Code("C03", "changes", err) }
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() { continue }; match := pendingName.FindStringSubmatch(entry.Name()); if match == nil { continue }
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] { return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2])) }; slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data, rel, num, match[2]); if e != nil { return nil, e }
+		if !modules[c.Module] { return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module)) }
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i,j int) bool { if changes[i].Number != changes[j].Number { return changes[i].Number < changes[j].Number }; return changes[i].Slug < changes[j].Slug })
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match { if idx >= 0 { return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector)) }; idx = i }
+	}
+	if idx < 0 { return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector)) }
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m); if err != nil { return contracts.Item{}, false }
+	for _, item := range items { if item.ID == id { return item, true } }; return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root); if err != nil { return snap, err }
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return snap, fixture.Code("C02", "", err) }
+	for i := 0; i < target; i++ { if err := Apply(&snap, changes[i]); err != nil { return snap, err } }
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, err }
+	limit := len(changes)
+	if through != "" { idx, e := Resolve(changes, through); if e != nil { return model.Snapshot{}, e }; limit = idx+1 }
+	snap, err := Before(root, changes, limit); if err != nil { return snap, err }
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1; for n, m := range snap.Modules { if m.Name == change.Module { i = n; break } }; if i < 0 { return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module)) }
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any); if !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping")) }
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err)) }
+			case "boundaries_opt_out": if _, ok := v.(string); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			case "hooks": if _, ok := v.(map[string]any); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			default: return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults); if err != nil { return err }; m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m); if err != nil { return fixture.Code("C03", path.Join(m.Dir,"SPEC.md"), err) }
+	for _, edit := range change.Modified { found := false; for j := range items { if items[j].ID == edit.ID { items[j] = edit; found = true; break } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) } }
+	for _, edit := range change.Removed { found := false; next := items[:0]; for _, item := range items { if item.ID == edit.ID { found = true } else { next = append(next, item) } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }; items = next }
+	items = append(items, change.Added...)
+	var body strings.Builder; for _, item := range items { body.WriteString(contracts.RenderItem(item)); if !strings.HasSuffix(body.String(), "\n") { body.WriteByte('\n') }; if !strings.HasSuffix(body.String(), "\n\n") { body.WriteByte('\n') } }
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return fixture.Code("C02", change.Path, err) }
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	idx, err := Resolve(changes, selector); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	snap, err := Before(root, changes, idx); return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]; if !exists { return map[string]any{}, nil }
+	m, ok := v.(map[string]any); if !ok { return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping")) }
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector); if err != nil { return err }
+	b, err := bases(c.Document); if err != nil { return err }
+	var m model.Module; for _, module := range snap.Modules { if module.Name == c.Module { m = module; break } }
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID); if !ok { return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }
+		if _, exists := b[edit.ID]; missingOnly && exists { continue }
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document); if err != nil { return err }
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/check.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..7d4345fb4638a24a682d0b4446af359e3702defa
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,152 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output(); if e != nil { return snap, fmt.Errorf("read revision %s tree: %w", revision, e) }
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") { if file != "" { parts := strings.Split(file, "/"); for i := 1; i < len(parts); i++ { allDirs[strings.Join(parts[:i], "/")] = true } } }
+		for dir := range allDirs { for _, glob := range globs { if contracts.MatchGlob(glob, dir) { dirs = append(dirs, dir); break } } }
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..20bbdbc0f628c9f942702446029a0f3001512689
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,66 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err); if !ok { code = fallback }
+	if file == "" && strings.HasSuffix(subject, ".md") { file = subject }
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) { out = append(out, model.Diagnostic{Code: code, Message: message, Path: file}) }
+	snap, err := fixture.Load(root); if err != nil { return []model.Diagnostic{diagnostic(err,"C03", "")} }
+	if err = fixture.ValidateGraph(snap.Modules); err != nil { add("C02",err.Error(),"") }
+	prefixes := map[string]string{}; for _, m := range snap.Modules { prefixes[m.IDPrefix] = m.Name }
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) { add("C04",fmt.Sprintf("non-normative spec body for %s",item.ID),file) }
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag,"quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag,"quarantine:")) != "") { add("C06",fmt.Sprintf("invalid review/quarantine tag for %s",item.ID),file) }
+		if !strings.HasPrefix(item.ID,"new:") { prefix := strings.SplitN(item.ID,"-",2)[0]; if owner, ok := prefixes[prefix]; ok && owner != module.Name { add("C05",fmt.Sprintf("foreign module ID %s belongs to %s",item.ID,owner),file) } }
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir,"SPEC.md")
+		section, _ := contracts.Section(m.Document.Body,"Specs")
+		items, e := contracts.ParseItems(section); if e != nil { out = append(out, diagnostic(e,"C03",file)); continue }
+		seen := map[string]bool{}
+		for _, item := range items { if seen[item.ID] { add("C05",fmt.Sprintf("duplicate item ID %s",item.ID),file) }; seen[item.ID] = true; validateItem(item,m,file,true) }
+	}
+	changes, err := Queue(root); if err != nil { out = append(out, diagnostic(err,"C03", "")); return out }
+	for _, c := range changes {
+		var m model.Module; for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+		motivation, _ := contracts.Section(c.Document.Body,"Motivation")
+		if strings.TrimSpace(motivation) == "" { add("C07","missing Motivation",c.Path) }
+		if evalFix, exists := c.Document.Frontmatter["eval_fix"]; exists && evalFix != nil {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 { add("C07","eval_fix cannot be mixed with fixes or spec edits",c.Path); continue }
+		}
+		declared := map[string]bool{}; for _, item := range c.Added { if strings.HasPrefix(item.ID,"new:") { declared[item.ID] = true } }
+		for _, item := range append(append([]contracts.Item{},c.Added...),c.Modified...) { validateItem(item,m,c.Path,true) }
+		for _, item := range c.Removed { validateItem(item,m,c.Path,false) }
+		for _, item := range append(append(append([]contracts.Item{},c.Added...),c.Modified...),c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body,-1) { if !declared[handle] { add("C09",fmt.Sprintf("reference to undeclared handle %s",handle),c.Path) } }
+		}
+		b, e := bases(c.Document); if e != nil { out = append(out, diagnostic(e,"C08",c.Path)) } else {
+			for _, item := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+				prior, ok := itemByID(m,item.ID); if !ok { add("C08",fmt.Sprintf("ID %s absent from preceding projection",item.ID),c.Path); continue }
+				recorded, present := b[item.ID]; hash, _ := recorded.(string)
+				if !present || hash == "" { add("C08",fmt.Sprintf("missing base for %s",item.ID),c.Path) } else if hash != contracts.SpecHash(prior) { add("C08",fmt.Sprintf("stale base for %s",item.ID),c.Path) }
+			}
+		}
+		if e = Apply(&snap,c); e != nil { out = append(out, diagnostic(e,"C03",c.Path)); break }
+		for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..5628d1ff809f24aea163f136051bbffba248e334
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,145 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string); if !ok || module == "" { return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel)) }; c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		for _, item := range items { if seen[item.ID] { return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID)) }; seen[item.ID] = true }
+		switch field { case "Added": c.Added = items; case "Modified": c.Modified = items; case "Removed": c.Removed = items }
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	modules := map[string]bool{}; for _, m := range snap.Modules { modules[m.Name] = true }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes")); if os.IsNotExist(err) { return []model.Change{}, nil }; if err != nil { return nil, fixture.Code("C03", "changes", err) }
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() { continue }; match := pendingName.FindStringSubmatch(entry.Name()); if match == nil { continue }
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] { return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2])) }; slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data, rel, num, match[2]); if e != nil { return nil, e }
+		if !modules[c.Module] { return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module)) }
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i,j int) bool { if changes[i].Number != changes[j].Number { return changes[i].Number < changes[j].Number }; return changes[i].Slug < changes[j].Slug })
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match { if idx >= 0 { return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector)) }; idx = i }
+	}
+	if idx < 0 { return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector)) }
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m); if err != nil { return contracts.Item{}, false }
+	for _, item := range items { if item.ID == id { return item, true } }; return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root); if err != nil { return snap, err }
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return snap, fixture.Code("C02", "", err) }
+	for i := 0; i < target; i++ { if err := Apply(&snap, changes[i]); err != nil { return snap, err } }
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, err }
+	limit := len(changes)
+	if through != "" { idx, e := Resolve(changes, through); if e != nil { return model.Snapshot{}, e }; limit = idx+1 }
+	snap, err := Before(root, changes, limit); if err != nil { return snap, err }
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1; for n, m := range snap.Modules { if m.Name == change.Module { i = n; break } }; if i < 0 { return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module)) }
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any); if !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping")) }
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err)) }
+			case "boundaries_opt_out": if _, ok := v.(string); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			case "hooks": if _, ok := v.(map[string]any); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			default: return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults); if err != nil { return err }; m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m); if err != nil { return fixture.Code("C03", path.Join(m.Dir,"SPEC.md"), err) }
+	for _, edit := range change.Modified { found := false; for j := range items { if items[j].ID == edit.ID { items[j] = edit; found = true; break } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) } }
+	for _, edit := range change.Removed { found := false; next := items[:0]; for _, item := range items { if item.ID == edit.ID { found = true } else { next = append(next, item) } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }; items = next }
+	items = append(items, change.Added...)
+	var body strings.Builder; for _, item := range items { body.WriteString(contracts.RenderItem(item)); if !strings.HasSuffix(body.String(), "\n") { body.WriteByte('\n') }; if !strings.HasSuffix(body.String(), "\n\n") { body.WriteByte('\n') } }
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return fixture.Code("C02", change.Path, err) }
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	idx, err := Resolve(changes, selector); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	snap, err := Before(root, changes, idx); return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]; if !exists { return map[string]any{}, nil }
+	m, ok := v.(map[string]any); if !ok { return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping")) }
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector); if err != nil { return err }
+	b, err := bases(c.Document); if err != nil { return err }
+	var m model.Module; for _, module := range snap.Modules { if module.Name == c.Module { m = module; break } }
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID); if !ok { return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }
+		if _, exists := b[edit.ID]; missingOnly && exists { continue }
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document); if err != nil { return err }
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }

codex
The loader, queue, projection, stamping, static checks, and archive history are now in place. I’m finishing the stale base recovery path, which needs to read the selected Git commit without touching the working tree, index, or refs.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/stale.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..7d4345fb4638a24a682d0b4446af359e3702defa
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,152 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output(); if e != nil { return snap, fmt.Errorf("read revision %s tree: %w", revision, e) }
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") { if file != "" { parts := strings.Split(file, "/"); for i := 1; i < len(parts); i++ { allDirs[strings.Join(parts[:i], "/")] = true } } }
+		for dir := range allDirs { for _, glob := range globs { if contracts.MatchGlob(glob, dir) { dirs = append(dirs, dir); break } } }
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..20bbdbc0f628c9f942702446029a0f3001512689
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,66 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err); if !ok { code = fallback }
+	if file == "" && strings.HasSuffix(subject, ".md") { file = subject }
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) { out = append(out, model.Diagnostic{Code: code, Message: message, Path: file}) }
+	snap, err := fixture.Load(root); if err != nil { return []model.Diagnostic{diagnostic(err,"C03", "")} }
+	if err = fixture.ValidateGraph(snap.Modules); err != nil { add("C02",err.Error(),"") }
+	prefixes := map[string]string{}; for _, m := range snap.Modules { prefixes[m.IDPrefix] = m.Name }
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) { add("C04",fmt.Sprintf("non-normative spec body for %s",item.ID),file) }
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag,"quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag,"quarantine:")) != "") { add("C06",fmt.Sprintf("invalid review/quarantine tag for %s",item.ID),file) }
+		if !strings.HasPrefix(item.ID,"new:") { prefix := strings.SplitN(item.ID,"-",2)[0]; if owner, ok := prefixes[prefix]; ok && owner != module.Name { add("C05",fmt.Sprintf("foreign module ID %s belongs to %s",item.ID,owner),file) } }
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir,"SPEC.md")
+		section, _ := contracts.Section(m.Document.Body,"Specs")
+		items, e := contracts.ParseItems(section); if e != nil { out = append(out, diagnostic(e,"C03",file)); continue }
+		seen := map[string]bool{}
+		for _, item := range items { if seen[item.ID] { add("C05",fmt.Sprintf("duplicate item ID %s",item.ID),file) }; seen[item.ID] = true; validateItem(item,m,file,true) }
+	}
+	changes, err := Queue(root); if err != nil { out = append(out, diagnostic(err,"C03", "")); return out }
+	for _, c := range changes {
+		var m model.Module; for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+		motivation, _ := contracts.Section(c.Document.Body,"Motivation")
+		if strings.TrimSpace(motivation) == "" { add("C07","missing Motivation",c.Path) }
+		if evalFix, exists := c.Document.Frontmatter["eval_fix"]; exists && evalFix != nil {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 { add("C07","eval_fix cannot be mixed with fixes or spec edits",c.Path); continue }
+		}
+		declared := map[string]bool{}; for _, item := range c.Added { if strings.HasPrefix(item.ID,"new:") { declared[item.ID] = true } }
+		for _, item := range append(append([]contracts.Item{},c.Added...),c.Modified...) { validateItem(item,m,c.Path,true) }
+		for _, item := range c.Removed { validateItem(item,m,c.Path,false) }
+		for _, item := range append(append(append([]contracts.Item{},c.Added...),c.Modified...),c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body,-1) { if !declared[handle] { add("C09",fmt.Sprintf("reference to undeclared handle %s",handle),c.Path) } }
+		}
+		b, e := bases(c.Document); if e != nil { out = append(out, diagnostic(e,"C08",c.Path)) } else {
+			for _, item := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+				prior, ok := itemByID(m,item.ID); if !ok { add("C08",fmt.Sprintf("ID %s absent from preceding projection",item.ID),c.Path); continue }
+				recorded, present := b[item.ID]; hash, _ := recorded.(string)
+				if !present || hash == "" { add("C08",fmt.Sprintf("missing base for %s",item.ID),c.Path) } else if hash != contracts.SpecHash(prior) { add("C08",fmt.Sprintf("stale base for %s",item.ID),c.Path) }
+			}
+		}
+		if e = Apply(&snap,c); e != nil { out = append(out, diagnostic(e,"C03",c.Path)); break }
+		for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..5628d1ff809f24aea163f136051bbffba248e334
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,145 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string); if !ok || module == "" { return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel)) }; c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section); if err != nil { return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		for _, item := range items { if seen[item.ID] { return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID)) }; seen[item.ID] = true }
+		switch field { case "Added": c.Added = items; case "Modified": c.Modified = items; case "Removed": c.Removed = items }
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	modules := map[string]bool{}; for _, m := range snap.Modules { modules[m.Name] = true }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes")); if os.IsNotExist(err) { return []model.Change{}, nil }; if err != nil { return nil, fixture.Code("C03", "changes", err) }
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() { continue }; match := pendingName.FindStringSubmatch(entry.Name()); if match == nil { continue }
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] { return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2])) }; slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data, rel, num, match[2]); if e != nil { return nil, e }
+		if !modules[c.Module] { return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module)) }
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i,j int) bool { if changes[i].Number != changes[j].Number { return changes[i].Number < changes[j].Number }; return changes[i].Slug < changes[j].Slug })
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match { if idx >= 0 { return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector)) }; idx = i }
+	}
+	if idx < 0 { return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector)) }
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m); if err != nil { return contracts.Item{}, false }
+	for _, item := range items { if item.ID == id { return item, true } }; return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root); if err != nil { return snap, err }
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return snap, fixture.Code("C02", "", err) }
+	for i := 0; i < target; i++ { if err := Apply(&snap, changes[i]); err != nil { return snap, err } }
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, err }
+	limit := len(changes)
+	if through != "" { idx, e := Resolve(changes, through); if e != nil { return model.Snapshot{}, e }; limit = idx+1 }
+	snap, err := Before(root, changes, limit); if err != nil { return snap, err }
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1; for n, m := range snap.Modules { if m.Name == change.Module { i = n; break } }; if i < 0 { return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module)) }
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any); if !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping")) }
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err)) }
+			case "boundaries_opt_out": if _, ok := v.(string); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			case "hooks": if _, ok := v.(map[string]any); !ok { return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key)) }
+			default: return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults); if err != nil { return err }; m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m); if err != nil { return fixture.Code("C03", path.Join(m.Dir,"SPEC.md"), err) }
+	for _, edit := range change.Modified { found := false; for j := range items { if items[j].ID == edit.ID { items[j] = edit; found = true; break } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) } }
+	for _, edit := range change.Removed { found := false; next := items[:0]; for _, item := range items { if item.ID == edit.ID { found = true } else { next = append(next, item) } }; if !found { return fixture.Code("C08", change.Path, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }; items = next }
+	items = append(items, change.Added...)
+	var body strings.Builder; for _, item := range items { body.WriteString(contracts.RenderItem(item)); if !strings.HasSuffix(body.String(), "\n") { body.WriteByte('\n') }; if !strings.HasSuffix(body.String(), "\n\n") { body.WriteByte('\n') } }
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil { return fixture.Code("C02", change.Path, err) }
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	idx, err := Resolve(changes, selector); if err != nil { return model.Snapshot{}, model.Change{}, err }
+	snap, err := Before(root, changes, idx); return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]; if !exists { return map[string]any{}, nil }
+	m, ok := v.(map[string]any); if !ok { return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping")) }
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector); if err != nil { return err }
+	b, err := bases(c.Document); if err != nil { return err }
+	var m model.Module; for _, module := range snap.Modules { if module.Name == c.Module { m = module; break } }
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID); if !ok { return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID)) }
+		if _, exists := b[edit.ID]; missingOnly && exists { continue }
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document); if err != nil { return err }
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

exec
/bin/zsh -lc "sed -n '340,390p' packages/project/evals/queue_test.go; sed -n '174,216p' packages/project/evals/queue_test.go; gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/queue/*.go packages/project/internal/history/history.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 573ms:
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

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..7d4345fb4638a24a682d0b4446af359e3702defa
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,152 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil { return "", err }
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err)) }
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil { return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err)) }
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok { return nil, fmt.Errorf("expected list of strings, got %T", value) }
+	out := make([]string, 0, len(list))
+	for _, v := range list { s, ok := v.(string); if !ok { return nil, fmt.Errorf("expected string, got %T", v) }; out = append(out, s) }
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string); if !ok { return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"])) }; m.Name = name
+	prefix, ok := front["id_prefix"].(string); if !ok { return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"])) }; m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present { deps, err := Strings(v); if err != nil { return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err)) }; m.DependsOn = deps }
+	get := func(key string) ([]string, error) { v, present := front[key]; if !present { v, present = defaults[key] }; if !present { return nil, nil }; list, err := Strings(v); if err != nil { return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err)) }; return list, nil }
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil { return m, err }
+	if m.Zones.Evals, err = get("evals"); err != nil { return m, err }
+	if m.Zones.Interface, err = get("interface"); err != nil { return m, err }
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil { return m, err }
+	if m.Zones.Preserve, err = get("preserve"); err != nil { return m, err }
+	if m.Zones.Disposable, err = get("disposable"); err != nil { return m, err }
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok { return nil }
+	items, err := contracts.ParseItems(section)
+	if err != nil { return Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+	seen := map[string]bool{}
+	for _, item := range items { if seen[item.ID] { return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID)) }; seen[item.ID] = true }
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root); if err != nil { return snap, err }; snap.Root = base
+	var configBytes []byte
+	if revision == "" { configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml")) } else { configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml") }
+	if err != nil { return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err)) }
+	snap.Config, err = Config(configBytes); if err != nil { return snap, err }
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil { return snap, Code("C01", "modules", e) }
+			for _, match := range matches { info, e := os.Stat(match); if e == nil && info.IsDir() { rel, _ := filepath.Rel(base, match); dirs = append(dirs, filepath.ToSlash(rel)) } }
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output(); if e != nil { return snap, fmt.Errorf("read revision %s tree: %w", revision, e) }
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") { if file != "" { parts := strings.Split(file, "/"); for i := 1; i < len(parts); i++ { allDirs[strings.Join(parts[:i], "/")] = true } } }
+		for dir := range allDirs { for _, glob := range globs { if contracts.MatchGlob(glob, dir) { dirs = append(dirs, dir); break } } }
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] { continue }; seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" { data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel))) } else { data, err = gitstore.ReadAt(base, revision, rel) }
+		if err != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e)) }
+		if e = specItems(doc, rel); e != nil { return snap, e }
+		module, e := ValidateMetadata(doc.Frontmatter, defaults); if e != nil { return snap, e }
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i,j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error) { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" { return fmt.Errorf("empty module name") }; if m.IDPrefix == "" { return fmt.Errorf("empty module prefix") }
+		if _, ok := byName[m.Name]; ok { return fmt.Errorf("duplicate module name %s", m.Name) }; if prefixes[m.IDPrefix] { return fmt.Errorf("duplicate module prefix %s", m.IDPrefix) }
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules { for _, dep := range m.DependsOn { if dep == m.Name { return fmt.Errorf("self dependency %s", m.Name) }; if _, ok := byName[dep]; !ok { return fmt.Errorf("module %s dependency %s is missing", m.Name, dep) } } }
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error { if state[name] == 1 { return fmt.Errorf("dependency cycle at %s", name) }; if state[name] == 2 { return nil }; state[name] = 1; for _, dep := range byName[name].DependsOn { if err := visit(dep); err != nil { return err } }; state[name] = 2; return nil }
+	for _, m := range modules { if err := visit(m.Name); err != nil { return err } }
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root); if err != nil { return "", err }
+	found := false; for _, m := range snap.Modules { if m.Name == module { found = true; break } }; if !found { return "", fmt.Errorf("unknown module %s", module) }
+	if !slugName.MatchString(slug) { return "", fmt.Errorf("invalid slug %q", slug) }
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder)); if e != nil && !os.IsNotExist(e) { return "", e }
+		for _, entry := range entries { parts := queuedName.FindStringSubmatch(entry.Name()); if parts == nil { continue }; if parts[2] == slug { return "", fmt.Errorf("duplicate slug %s", slug) }; n, _ := strconv.Atoi(parts[1]); if n > max { max = n } }
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content); closeErr := file.Close(); if err != nil { return "", fmt.Errorf("create %s: %w", rel, err) }; if closeErr != nil { return "", fmt.Errorf("create %s: %w", rel, closeErr) }
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..20bbdbc0f628c9f942702446029a0f3001512689
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,66 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err); if !ok { code = fallback }
+	if file == "" && strings.HasSuffix(subject, ".md") { file = subject }
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) { out = append(out, model.Diagnostic{Code: code, Message: message, Path: file}) }
+	snap, err := fixture.Load(root); if err != nil { return []model.Diagnostic{diagnostic(err,"C03", "")} }
+	if err = fixture.ValidateGraph(snap.Modules); err != nil { add("C02",err.Error(),"") }
+	prefixes := map[string]string{}; for _, m := range snap.Modules { prefixes[m.IDPrefix] = m.Name }
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) { add("C04",fmt.Sprintf("non-normative spec body for %s",item.ID),file) }
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag,"quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag,"quarantine:")) != "") { add("C06",fmt.Sprintf("invalid review/quarantine tag for %s",item.ID),file) }
+		if !strings.HasPrefix(item.ID,"new:") { prefix := strings.SplitN(item.ID,"-",2)[0]; if owner, ok := prefixes[prefix]; ok && owner != module.Name { add("C05",fmt.Sprintf("foreign module ID %s belongs to %s",item.ID,owner),file) } }
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir,"SPEC.md")
+		section, _ := contracts.Section(m.Document.Body,"Specs")
+		items, e := contracts.ParseItems(section); if e != nil { out = append(out, diagnostic(e,"C03",file)); continue }
+		seen := map[string]bool{}
+		for _, item := range items { if seen[item.ID] { add("C05",fmt.Sprintf("duplicate item ID %s",item.ID),file) }; seen[item.ID] = true; validateItem(item,m,file,true) }
+	}
+	changes, err := Queue(root); if err != nil { out = append(out, diagnostic(err,"C03", "")); return out }
+	for _, c := range changes {
+		var m model.Module; for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+		motivation, _ := contracts.Section(c.Document.Body,"Motivation")
+		if strings.TrimSpace(motivation) == "" { add("C07","missing Motivation",c.Path) }
+		if evalFix, exists := c.Document.Frontmatter["eval_fix"]; exists && evalFix != nil {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 { add("C07","eval_fix cannot be mixed with fixes or spec edits",c.Path); continue }
+		}
+		declared := map[string]bool{}; for _, item := range c.Added { if strings.HasPrefix(item.ID,"new:") { declared[item.ID] = true } }
+		for _, item := range append(append([]contracts.Item{},c.Added...),c.Modified...) { validateItem(item,m,c.Path,true) }
+		for _, item := range c.Removed { validateItem(item,m,c.Path,false) }
+		for _, item := range append(append(append([]contracts.Item{},c.Added...),c.Modified...),c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body,-1) { if !declared[handle] { add("C09",fmt.Sprintf("reference to undeclared handle %s",handle),c.Path) } }
+		}
+		b, e := bases(c.Document); if e != nil { out = append(out, diagnostic(e,"C08",c.Path)) } else {
+			for _, item := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+				prior, ok := itemByID(m,item.ID); if !ok { add("C08",fmt.Sprintf("ID %s absent from preceding projection",item.ID),c.Path); continue }
+				recorded, present := b[item.ID]; hash, _ := recorded.(string)
+				if !present || hash == "" { add("C08",fmt.Sprintf("missing base for %s",item.ID),c.Path) } else if hash != contracts.SpecHash(prior) { add("C08",fmt.Sprintf("stale base for %s",item.ID),c.Path) }
+			}
+		}
+		if e = Apply(&snap,c); e != nil { out = append(out, diagnostic(e,"C03",c.Path)); break }
+		for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..f95e27c4308c91ff69949accb09b11ad4876e7aa
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,336 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..419904825ec27d2528187a024e4b1200b114c5a1
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,348 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, err
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults, _ := snap.Config["module_defaults"].(map[string]any)
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() { continue }
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..20bbdbc0f628c9f942702446029a0f3001512689
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,66 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err); if !ok { code = fallback }
+	if file == "" && strings.HasSuffix(subject, ".md") { file = subject }
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) { out = append(out, model.Diagnostic{Code: code, Message: message, Path: file}) }
+	snap, err := fixture.Load(root); if err != nil { return []model.Diagnostic{diagnostic(err,"C03", "")} }
+	if err = fixture.ValidateGraph(snap.Modules); err != nil { add("C02",err.Error(),"") }
+	prefixes := map[string]string{}; for _, m := range snap.Modules { prefixes[m.IDPrefix] = m.Name }
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) { add("C04",fmt.Sprintf("non-normative spec body for %s",item.ID),file) }
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag,"quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag,"quarantine:")) != "") { add("C06",fmt.Sprintf("invalid review/quarantine tag for %s",item.ID),file) }
+		if !strings.HasPrefix(item.ID,"new:") { prefix := strings.SplitN(item.ID,"-",2)[0]; if owner, ok := prefixes[prefix]; ok && owner != module.Name { add("C05",fmt.Sprintf("foreign module ID %s belongs to %s",item.ID,owner),file) } }
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir,"SPEC.md")
+		section, _ := contracts.Section(m.Document.Body,"Specs")
+		items, e := contracts.ParseItems(section); if e != nil { out = append(out, diagnostic(e,"C03",file)); continue }
+		seen := map[string]bool{}
+		for _, item := range items { if seen[item.ID] { add("C05",fmt.Sprintf("duplicate item ID %s",item.ID),file) }; seen[item.ID] = true; validateItem(item,m,file,true) }
+	}
+	changes, err := Queue(root); if err != nil { out = append(out, diagnostic(err,"C03", "")); return out }
+	for _, c := range changes {
+		var m model.Module; for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+		motivation, _ := contracts.Section(c.Document.Body,"Motivation")
+		if strings.TrimSpace(motivation) == "" { add("C07","missing Motivation",c.Path) }
+		if evalFix, exists := c.Document.Frontmatter["eval_fix"]; exists && evalFix != nil {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 { add("C07","eval_fix cannot be mixed with fixes or spec edits",c.Path); continue }
+		}
+		declared := map[string]bool{}; for _, item := range c.Added { if strings.HasPrefix(item.ID,"new:") { declared[item.ID] = true } }
+		for _, item := range append(append([]contracts.Item{},c.Added...),c.Modified...) { validateItem(item,m,c.Path,true) }
+		for _, item := range c.Removed { validateItem(item,m,c.Path,false) }
+		for _, item := range append(append(append([]contracts.Item{},c.Added...),c.Modified...),c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body,-1) { if !declared[handle] { add("C09",fmt.Sprintf("reference to undeclared handle %s",handle),c.Path) } }
+		}
+		b, e := bases(c.Document); if e != nil { out = append(out, diagnostic(e,"C08",c.Path)) } else {
+			for _, item := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+				prior, ok := itemByID(m,item.ID); if !ok { add("C08",fmt.Sprintf("ID %s absent from preceding projection",item.ID),c.Path); continue }
+				recorded, present := b[item.ID]; hash, _ := recorded.(string)
+				if !present || hash == "" { add("C08",fmt.Sprintf("missing base for %s",item.ID),c.Path) } else if hash != contracts.SpecHash(prior) { add("C08",fmt.Sprintf("stale base for %s",item.ID),c.Path) }
+			}
+		}
+		if e = Apply(&snap,c); e != nil { out = append(out, diagnostic(e,"C03",c.Path)); break }
+		for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..f95e27c4308c91ff69949accb09b11ad4876e7aa
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,336 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

exec
/bin/zsh -lc 'GOCACHE=/tmp/fireweed-rebuild-cache go test ./packages/project/evals' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/queue/*.go packages/project/internal/history/history.go && git diff --check' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 543ms:
?? packages/project/internal/

 succeeded in 561ms:
 exited 1 in 585ms:
go: cannot find main module, but found .git/config in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
	to create a module there, run:
	go mod init

exec
/bin/zsh -lc "sed -n '1,280p' packages/project/internal/fixture/fixture.go; sed -n '1,360p' packages/project/internal/queue/queue.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 521ms:
package fixture

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func Code(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
}

func canonical(root string) (string, error) {
	absolute, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(absolute)
}

func Config(data []byte) (map[string]any, error) {
	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
	if err != nil {
		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
	}
	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
	}
	return doc.Frontmatter, nil
}

func Strings(value any) ([]string, error) {
	list, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("expected list of strings, got %T", value)
	}
	out := make([]string, 0, len(list))
	for _, v := range list {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("expected string, got %T", v)
		}
		out = append(out, s)
	}
	return out, nil
}

var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}

func ZoneKeys() []string { return zoneKeys }

func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
	m := model.Module{Document: contracts.Document{Frontmatter: front}}
	name, ok := front["name"].(string)
	if !ok {
		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
	}
	m.Name = name
	prefix, ok := front["id_prefix"].(string)
	if !ok {
		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
	}
	m.IDPrefix = prefix
	if v, present := front["depends_on"]; present {
		deps, err := Strings(v)
		if err != nil {
			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
		}
		m.DependsOn = deps
	}
	get := func(key string) ([]string, error) {
		v, present := front[key]
		if !present {
			v, present = defaults[key]
		}
		if !present {
			return nil, nil
		}
		list, err := Strings(v)
		if err != nil {
			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
		}
		return list, nil
	}
	var err error
	if m.Zones.Protected, err = get("protected"); err != nil {
		return m, err
	}
	if m.Zones.Evals, err = get("evals"); err != nil {
		return m, err
	}
	if m.Zones.Interface, err = get("interface"); err != nil {
		return m, err
	}
	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
		return m, err
	}
	if m.Zones.Preserve, err = get("preserve"); err != nil {
		return m, err
	}
	if m.Zones.Disposable, err = get("disposable"); err != nil {
		return m, err
	}
	return m, nil
}

func specItems(doc contracts.Document, rel string) error {
	section, ok := contracts.Section(doc.Body, "Specs")
	if !ok {
		return nil
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
	}
	seen := map[string]bool{}
	for _, item := range items {
		if seen[item.ID] {
			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
		}
		seen[item.ID] = true
	}
	return nil
}

func load(root, revision string) (model.Snapshot, error) {
	var snap model.Snapshot
	base, err := canonical(root)
	if err != nil {
		return snap, err
	}
	snap.Root = base
	var configBytes []byte
	if revision == "" {
		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
	} else {
		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
	}
	if err != nil {
		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
	}
	snap.Config, err = Config(configBytes)
	if err != nil {
		return snap, err
	}
	globs, _ := Strings(snap.Config["modules"])
	defaults, _ := snap.Config["module_defaults"].(map[string]any)
	var dirs []string
	if revision == "" {
		for _, glob := range globs {
			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
			if e != nil {
				return snap, Code("C01", "modules", e)
			}
			for _, match := range matches {
				info, e := os.Stat(match)
				if e == nil && info.IsDir() {
					rel, _ := filepath.Rel(base, match)
					dirs = append(dirs, filepath.ToSlash(rel))
				}
			}
		}
	} else {
		// Tree names come from the selected revision; working tree directories are irrelevant.
		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
		out, e := cmd.Output()
		if e != nil {
			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
		}
		allDirs := map[string]bool{}
		for _, file := range strings.Split(string(out), "\n") {
			if file != "" {
				parts := strings.Split(file, "/")
				for i := 1; i < len(parts); i++ {
					allDirs[strings.Join(parts[:i], "/")] = true
				}
			}
		}
		for dir := range allDirs {
			for _, glob := range globs {
				if contracts.MatchGlob(glob, dir) {
					dirs = append(dirs, dir)
					break
				}
			}
		}
	}
	seen := map[string]bool{}
	for _, dir := range dirs {
		if seen[dir] {
			continue
		}
		seen[dir] = true
		rel := path.Join(dir, "SPEC.md")
		var data []byte
		if revision == "" {
			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
		} else {
			data, err = gitstore.ReadAt(base, revision, rel)
		}
		if err != nil {
			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
		}
		doc, e := contracts.ParseDocument(string(data))
		if e != nil {
			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
		}
		if e = specItems(doc, rel); e != nil {
			return snap, e
		}
		module, e := ValidateMetadata(doc.Frontmatter, defaults)
		if e != nil {
			return snap, e
		}
		module.Dir, module.Document = dir, doc
		snap.Modules = append(snap.Modules, module)
	}
	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
	return snap, nil
}

func Load(root string) (model.Snapshot, error)             { return load(root, "") }
func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }

func ValidateGraph(modules []model.Module) error {
	byName, prefixes := map[string]model.Module{}, map[string]bool{}
	for _, m := range modules {
		if m.Name == "" {
			return fmt.Errorf("empty module name")
		}
		if m.IDPrefix == "" {
			return fmt.Errorf("empty module prefix")
		}
		if _, ok := byName[m.Name]; ok {
			return fmt.Errorf("duplicate module name %s", m.Name)
		}
		if prefixes[m.IDPrefix] {
			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
		}
		byName[m.Name], prefixes[m.IDPrefix] = m, true
	}
	for _, m := range modules {
		for _, dep := range m.DependsOn {
			if dep == m.Name {
				return fmt.Errorf("self dependency %s", m.Name)
			}
			if _, ok := byName[dep]; !ok {
				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
			}
		}
	}
	state := map[string]int{}
	var visit func(string) error
	visit = func(name string) error {
		if state[name] == 1 {
			return fmt.Errorf("dependency cycle at %s", name)
		}
		if state[name] == 2 {
			return nil
		}
		state[name] = 1
		for _, dep := range byName[name].DependsOn {
			if err := visit(dep); err != nil {
				return err
			}
		}
		state[name] = 2
		return nil
	}
	for _, m := range modules {
package queue

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)

func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
	c := model.Change{Number: number, Slug: slug, Path: rel}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
	}
	c.Document = doc
	module, ok := doc.Frontmatter["module"].(string)
	if !ok || module == "" {
		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
	}
	c.Module = module
	seen := map[string]bool{}
	for _, field := range []string{"Added", "Modified", "Removed"} {
		section, _ := contracts.Section(doc.Body, field)
		items, err := contracts.ParseItems(section)
		if err != nil {
			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
		}
		for _, item := range items {
			if seen[item.ID] {
				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
			}
			seen[item.ID] = true
		}
		switch field {
		case "Added":
			c.Added = items
		case "Modified":
			c.Modified = items
		case "Removed":
			c.Removed = items
		}
	}
	return c, nil
}

func Queue(root string) ([]model.Change, error) {
	snap, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	modules := map[string]bool{}
	for _, m := range snap.Modules {
		modules[m.Name] = true
	}
	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
	if os.IsNotExist(err) {
		return []model.Change{}, nil
	}
	if err != nil {
		return nil, fixture.Code("C03", "changes", err)
	}
	var changes []model.Change
	slugs := map[string]bool{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		match := pendingName.FindStringSubmatch(entry.Name())
		if match == nil {
			continue
		}
		rel := path.Join("changes", entry.Name())
		if slugs[match[2]] {
			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
		}
		slugs[match[2]] = true
		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
		if e != nil {
			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
		}
		num, _ := strconv.Atoi(match[1])
		c, e := ParseChange(data, rel, num, match[2])
		if e != nil {
			return nil, e
		}
		if !modules[c.Module] {
			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
		}
		changes = append(changes, c)
	}
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].Number != changes[j].Number {
			return changes[i].Number < changes[j].Number
		}
		return changes[i].Slug < changes[j].Slug
	})
	return changes, nil
}

func Resolve(changes []model.Change, selector string) (int, error) {
	idx := -1
	for i, c := range changes {
		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
		if match {
			if idx >= 0 {
				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
			}
			idx = i
		}
	}
	if idx < 0 {
		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
	}
	return idx, nil
}

func itemList(m model.Module) ([]contracts.Item, error) {
	section, _ := contracts.Section(m.Document.Body, "Specs")
	return contracts.ParseItems(section)
}

func itemByID(m model.Module, id string) (contracts.Item, bool) {
	items, err := itemList(m)
	if err != nil {
		return contracts.Item{}, false
	}
	for _, item := range items {
		if item.ID == id {
			return item, true
		}
	}
	return contracts.Item{}, false
}

func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
	snap, err := fixture.Load(root)
	if err != nil {
		return snap, err
	}
	if err := fixture.ValidateGraph(snap.Modules); err != nil {
		return snap, fixture.Code("C02", "", err)
	}
	for i := 0; i < target; i++ {
		if err := Apply(&snap, changes[i]); err != nil {
			return snap, err
		}
	}
	return snap, nil
}

func Project(root, through string) (model.Snapshot, error) {
	changes, err := Queue(root)
	if err != nil {
		return model.Snapshot{}, err
	}
	limit := len(changes)
	if through != "" {
		idx, e := Resolve(changes, through)
		if e != nil {
			return model.Snapshot{}, e
		}
		limit = idx + 1
	}
	snap, err := Before(root, changes, limit)
	if err != nil {
		return snap, err
	}
	return snap, nil
}

func Apply(snap *model.Snapshot, change model.Change) error {
	i := -1
	for n, m := range snap.Modules {
		if m.Name == change.Module {
			i = n
			break
		}
	}
	if i < 0 {
		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
	}
	m := snap.Modules[i]
	if value, ok := change.Document.Frontmatter["set"]; ok {
		settings, ok := value.(map[string]any)
		if !ok {
			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
		}
		for key, v := range settings {
			switch key {
			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
				if _, err := fixture.Strings(v); err != nil {
					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
				}
			case "boundaries_opt_out":
				if _, ok := v.(string); !ok {
					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
				}
			case "hooks":
				if _, ok := v.(map[string]any); !ok {
					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
				}
			default:
				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
			}
			m.Document.Frontmatter[key] = v
		}
		defaults, _ := snap.Config["module_defaults"].(map[string]any)
		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
		if err != nil {
			return err
		}
		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
	}
	items, err := itemList(m)
	if err != nil {
		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
	}
	for _, edit := range change.Modified {
		found := false
		for j := range items {
			if items[j].ID == edit.ID {
				items[j] = edit
				found = true
				break
			}
		}
		if !found {
			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
		}
	}
	for _, edit := range change.Removed {
		found := false
		next := items[:0]
		for _, item := range items {
			if item.ID == edit.ID {
				found = true
			} else {
				next = append(next, item)
			}
		}
		if !found {
			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
		}
		items = next
	}
	items = append(items, change.Added...)
	var body strings.Builder
	for _, item := range items {
		body.WriteString(contracts.RenderItem(item))
		if !strings.HasSuffix(body.String(), "\n") {
			body.WriteByte('\n')
		}
		if !strings.HasSuffix(body.String(), "\n\n") {
			body.WriteByte('\n')
		}
	}
	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
	snap.Modules[i] = m
	if err := fixture.ValidateGraph(snap.Modules); err != nil {
		return fixture.Code("C02", change.Path, err)
	}
	return nil
}

func Target(root, selector string) (model.Snapshot, model.Change, error) {
	changes, err := Queue(root)
	if err != nil {
		return model.Snapshot{}, model.Change{}, err
	}
	idx, err := Resolve(changes, selector)
	if err != nil {
		return model.Snapshot{}, model.Change{}, err
	}
	snap, err := Before(root, changes, idx)
	return snap, changes[idx], err
}

func bases(doc contracts.Document) (map[string]any, error) {
	v, exists := doc.Frontmatter["bases"]
	if !exists {
		return map[string]any{}, nil
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
	}
	return m, nil
}

func stamp(root, selector string, missingOnly bool) error {
	snap, c, err := Target(root, selector)
	if err != nil {
		return err
	}
	b, err := bases(c.Document)
	if err != nil {
		return err
	}
	var m model.Module
	for _, module := range snap.Modules {
		if module.Name == c.Module {
			m = module
			break
		}
	}
	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
		item, ok := itemByID(m, edit.ID)
		if !ok {
			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
		}
		if _, exists := b[edit.ID]; missingOnly && exists {
			continue
		}
		b[edit.ID] = contracts.SpecHash(item)
	}
	c.Document.Frontmatter["bases"] = b
	text, err := contracts.RenderDocument(c.Document)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
}

func Stamp(root, selector string) error        { return stamp(root, selector, false) }
func StampMissing(root, selector string) error { return stamp(root, selector, true) }

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..4363c90e5ee018cd7c173382fa44c69252e11dbd
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,356 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, err
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok { return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping")) }
+		defaults = mapping
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..20bbdbc0f628c9f942702446029a0f3001512689
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,66 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err); if !ok { code = fallback }
+	if file == "" && strings.HasSuffix(subject, ".md") { file = subject }
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) { out = append(out, model.Diagnostic{Code: code, Message: message, Path: file}) }
+	snap, err := fixture.Load(root); if err != nil { return []model.Diagnostic{diagnostic(err,"C03", "")} }
+	if err = fixture.ValidateGraph(snap.Modules); err != nil { add("C02",err.Error(),"") }
+	prefixes := map[string]string{}; for _, m := range snap.Modules { prefixes[m.IDPrefix] = m.Name }
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) { add("C04",fmt.Sprintf("non-normative spec body for %s",item.ID),file) }
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag,"quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag,"quarantine:")) != "") { add("C06",fmt.Sprintf("invalid review/quarantine tag for %s",item.ID),file) }
+		if !strings.HasPrefix(item.ID,"new:") { prefix := strings.SplitN(item.ID,"-",2)[0]; if owner, ok := prefixes[prefix]; ok && owner != module.Name { add("C05",fmt.Sprintf("foreign module ID %s belongs to %s",item.ID,owner),file) } }
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir,"SPEC.md")
+		section, _ := contracts.Section(m.Document.Body,"Specs")
+		items, e := contracts.ParseItems(section); if e != nil { out = append(out, diagnostic(e,"C03",file)); continue }
+		seen := map[string]bool{}
+		for _, item := range items { if seen[item.ID] { add("C05",fmt.Sprintf("duplicate item ID %s",item.ID),file) }; seen[item.ID] = true; validateItem(item,m,file,true) }
+	}
+	changes, err := Queue(root); if err != nil { out = append(out, diagnostic(err,"C03", "")); return out }
+	for _, c := range changes {
+		var m model.Module; for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+		motivation, _ := contracts.Section(c.Document.Body,"Motivation")
+		if strings.TrimSpace(motivation) == "" { add("C07","missing Motivation",c.Path) }
+		if evalFix, exists := c.Document.Frontmatter["eval_fix"]; exists && evalFix != nil {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 { add("C07","eval_fix cannot be mixed with fixes or spec edits",c.Path); continue }
+		}
+		declared := map[string]bool{}; for _, item := range c.Added { if strings.HasPrefix(item.ID,"new:") { declared[item.ID] = true } }
+		for _, item := range append(append([]contracts.Item{},c.Added...),c.Modified...) { validateItem(item,m,c.Path,true) }
+		for _, item := range c.Removed { validateItem(item,m,c.Path,false) }
+		for _, item := range append(append(append([]contracts.Item{},c.Added...),c.Modified...),c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body,-1) { if !declared[handle] { add("C09",fmt.Sprintf("reference to undeclared handle %s",handle),c.Path) } }
+		}
+		b, e := bases(c.Document); if e != nil { out = append(out, diagnostic(e,"C08",c.Path)) } else {
+			for _, item := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+				prior, ok := itemByID(m,item.ID); if !ok { add("C08",fmt.Sprintf("ID %s absent from preceding projection",item.ID),c.Path); continue }
+				recorded, present := b[item.ID]; hash, _ := recorded.(string)
+				if !present || hash == "" { add("C08",fmt.Sprintf("missing base for %s",item.ID),c.Path) } else if hash != contracts.SpecHash(prior) { add("C08",fmt.Sprintf("stale base for %s",item.ID),c.Path) }
+			}
+		}
+		if e = Apply(&snap,c); e != nil { out = append(out, diagnostic(e,"C03",c.Path)); break }
+		for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..f95e27c4308c91ff69949accb09b11ad4876e7aa
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,336 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		for key, v := range settings {
+			switch key {
+			case "depends_on", "protected", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..4363c90e5ee018cd7c173382fa44c69252e11dbd
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,356 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, err
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok { return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping")) }
+		defaults = mapping
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..20bbdbc0f628c9f942702446029a0f3001512689
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,66 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err); if !ok { code = fallback }
+	if file == "" && strings.HasSuffix(subject, ".md") { file = subject }
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) { out = append(out, model.Diagnostic{Code: code, Message: message, Path: file}) }
+	snap, err := fixture.Load(root); if err != nil { return []model.Diagnostic{diagnostic(err,"C03", "")} }
+	if err = fixture.ValidateGraph(snap.Modules); err != nil { add("C02",err.Error(),"") }
+	prefixes := map[string]string{}; for _, m := range snap.Modules { prefixes[m.IDPrefix] = m.Name }
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) { add("C04",fmt.Sprintf("non-normative spec body for %s",item.ID),file) }
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag,"quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag,"quarantine:")) != "") { add("C06",fmt.Sprintf("invalid review/quarantine tag for %s",item.ID),file) }
+		if !strings.HasPrefix(item.ID,"new:") { prefix := strings.SplitN(item.ID,"-",2)[0]; if owner, ok := prefixes[prefix]; ok && owner != module.Name { add("C05",fmt.Sprintf("foreign module ID %s belongs to %s",item.ID,owner),file) } }
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir,"SPEC.md")
+		section, _ := contracts.Section(m.Document.Body,"Specs")
+		items, e := contracts.ParseItems(section); if e != nil { out = append(out, diagnostic(e,"C03",file)); continue }
+		seen := map[string]bool{}
+		for _, item := range items { if seen[item.ID] { add("C05",fmt.Sprintf("duplicate item ID %s",item.ID),file) }; seen[item.ID] = true; validateItem(item,m,file,true) }
+	}
+	changes, err := Queue(root); if err != nil { out = append(out, diagnostic(err,"C03", "")); return out }
+	for _, c := range changes {
+		var m model.Module; for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+		motivation, _ := contracts.Section(c.Document.Body,"Motivation")
+		if strings.TrimSpace(motivation) == "" { add("C07","missing Motivation",c.Path) }
+		if evalFix, exists := c.Document.Frontmatter["eval_fix"]; exists && evalFix != nil {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 { add("C07","eval_fix cannot be mixed with fixes or spec edits",c.Path); continue }
+		}
+		declared := map[string]bool{}; for _, item := range c.Added { if strings.HasPrefix(item.ID,"new:") { declared[item.ID] = true } }
+		for _, item := range append(append([]contracts.Item{},c.Added...),c.Modified...) { validateItem(item,m,c.Path,true) }
+		for _, item := range c.Removed { validateItem(item,m,c.Path,false) }
+		for _, item := range append(append(append([]contracts.Item{},c.Added...),c.Modified...),c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body,-1) { if !declared[handle] { add("C09",fmt.Sprintf("reference to undeclared handle %s",handle),c.Path) } }
+		}
+		b, e := bases(c.Document); if e != nil { out = append(out, diagnostic(e,"C08",c.Path)) } else {
+			for _, item := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+				prior, ok := itemByID(m,item.ID); if !ok { add("C08",fmt.Sprintf("ID %s absent from preceding projection",item.ID),c.Path); continue }
+				recorded, present := b[item.ID]; hash, _ := recorded.(string)
+				if !present || hash == "" { add("C08",fmt.Sprintf("missing base for %s",item.ID),c.Path) } else if hash != contracts.SpecHash(prior) { add("C08",fmt.Sprintf("stale base for %s",item.ID),c.Path) }
+			}
+		}
+		if e = Apply(&snap,c); e != nil { out = append(out, diagnostic(e,"C03",c.Path)); break }
+		for _, candidate := range snap.Modules { if candidate.Name == c.Module { m = candidate; break } }
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..cd0b4571879d6c30e946b1ccfa07cbcc2cedf992
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,340 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings { keys = append(keys, key) }
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/check.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..4363c90e5ee018cd7c173382fa44c69252e11dbd
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,356 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, err
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok { return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping")) }
+		defaults = mapping
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..cd0b4571879d6c30e946b1ccfa07cbcc2cedf992
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,340 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings { keys = append(keys, key) }
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..bc3c6e6182c863c0c55c8ecee7bbc809adfd1d5e
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,356 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, err
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok { return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping")) }
+		defaults = mapping
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..3df172c33c480268878a0e0dc63da4dd4c7a1414
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,340 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings { keys = append(keys, key) }
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..d234b5d536ef3804f833fabd2b318ec09860adeb
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,359 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present { if _, ok := value.(map[string]any); !ok { return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping")) } }
+	if value, present := front["boundaries_opt_out"]; present { if _, ok := value.(string); !ok { return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string")) } }
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, err
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok { return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping")) }
+		defaults = mapping
+		for _, key := range zoneKeys { if value, exists := defaults[key]; exists { if _, e := Strings(value); e != nil { return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e)) } } }
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..3df172c33c480268878a0e0dc63da4dd4c7a1414
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,340 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings { keys = append(keys, key) }
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

exec
/bin/zsh -lc 'gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/queue/*.go packages/project/internal/history/history.go && git diff --check && git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "rg -n 'RunGitTree|TODO|panic\\(' packages/project/internal || true" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 539ms:
 succeeded in 564ms:
?? packages/project/internal/

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..b09f8ae0fb1654f76beedde490ff508dcfd78c26
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,375 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present {
+		if _, ok := value.(map[string]any); !ok {
+			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
+		}
+	}
+	if value, present := front["boundaries_opt_out"]; present {
+		if _, ok := value.(string); !ok {
+			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
+		}
+	}
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok {
+			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
+		}
+		defaults = mapping
+		for _, key := range zoneKeys {
+			if value, exists := defaults[key]; exists {
+				if _, e := Strings(value); e != nil {
+					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
+				}
+			}
+		}
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..dacefc9436bbb86626a01367f69a32f4bf4160cb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,49 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root); if err != nil { return nil, err }
+	var current model.Module; found := false
+	for _, m := range snap.Modules { if m.Name == module { current, found = m, true; break } }
+	if !found { return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module)) }
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive")); if os.IsNotExist(err) { entries = nil } else if err != nil { return nil, err }
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") { continue }
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel))); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		doc, e := contracts.ParseDocument(string(data)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+		var ids []string; seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name)); if e != nil { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e)) }
+			for _, item := range items { if !seen[item.ID] { seen[item.ID] = true; ids = append(ids, item.ID) } }
+		}
+		archiveModule, _ := doc.Frontmatter["module"].(string)
+		if archiveModule != module || len(ids) == 0 { continue }
+		if id != "" { if !seen[id] { continue }; ids = []string{id} } else { sort.Strings(ids) }
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body,"Motivation"), Decisions: section(doc.Body,"Decisions"), RemovedDecisions: section(doc.Body,"Removed decisions"), Provenance: section(doc.Body,"Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") { fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-")); if len(fields) < 2 || fields[0] != "impl" { continue }; date, e := time.Parse(time.RFC3339Nano, fields[1]); if e == nil && date.After(entry.ImplementationDate) { entry.ImplementationDate = date } }
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false; body, _ := contracts.Section(current.Document.Body, "Specs"); items, _ := contracts.ParseItems(body); for _, item := range items { if item.ID == id { known = true; break } }
+		if !known { return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id)) }
+	}
+	sort.Slice(result, func(i,j int) bool { a,b := result[i], result[j]; if !a.ImplementationDate.Equal(b.ImplementationDate) { return a.ImplementationDate.After(b.ImplementationDate) }; return a.Path < b.Path })
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..3df172c33c480268878a0e0dc63da4dd4c7a1414
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,340 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings { keys = append(keys, key) }
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/history/history.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..b09f8ae0fb1654f76beedde490ff508dcfd78c26
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,375 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present {
+		if _, ok := value.(map[string]any); !ok {
+			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
+		}
+	}
+	if value, present := front["boundaries_opt_out"]; present {
+		if _, ok := value.(string); !ok {
+			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
+		}
+	}
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok {
+			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
+		}
+		defaults = mapping
+		for _, key := range zoneKeys {
+			if value, exists := defaults[key]; exists {
+				if _, e := Strings(value); e != nil {
+					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
+				}
+			}
+		}
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..7e0b7a8a3cf7fb413e8cee04e99e21a3e299f211
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,117 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	var current model.Module
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			current, found = m, true
+			break
+		}
+	}
+	if !found {
+		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive"))
+	if os.IsNotExist(err) {
+		entries = nil
+	} else if err != nil {
+		return nil, err
+	}
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
+			continue
+		}
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		archiveModule, ok := doc.Frontmatter["module"].(string)
+		if !ok || archiveModule == "" { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: invalid module target", rel)) }
+		var ids []string
+		seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name))
+			if e != nil {
+				return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+			}
+			for _, item := range items {
+				if !seen[item.ID] {
+					seen[item.ID] = true
+					ids = append(ids, item.ID)
+				}
+			}
+		}
+		if archiveModule != module || len(ids) == 0 {
+			continue
+		}
+		if id != "" {
+			if !seen[id] {
+				continue
+			}
+			ids = []string{id}
+		} else {
+			sort.Strings(ids)
+		}
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body, "Motivation"), Decisions: section(doc.Body, "Decisions"), RemovedDecisions: section(doc.Body, "Removed decisions"), Provenance: section(doc.Body, "Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") {
+			fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-"))
+			if len(fields) < 2 || fields[0] != "impl" {
+				continue
+			}
+			date, e := time.Parse(time.RFC3339Nano, fields[1])
+			if e == nil && date.After(entry.ImplementationDate) {
+				entry.ImplementationDate = date
+			}
+		}
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false
+		body, _ := contracts.Section(current.Document.Body, "Specs")
+		items, _ := contracts.ParseItems(body)
+		for _, item := range items {
+			if item.ID == id {
+				known = true
+				break
+			}
+		}
+		if !known {
+			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
+		}
+	}
+	sort.Slice(result, func(i, j int) bool {
+		a, b := result[i], result[j]
+		if !a.ImplementationDate.Equal(b.ImplementationDate) {
+			return a.ImplementationDate.After(b.ImplementationDate)
+		}
+		return a.Path < b.Path
+	})
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..3df172c33c480268878a0e0dc63da4dd4c7a1414
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,340 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings { keys = append(keys, key) }
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..daf19b2ba7b8d622a14a79fb201ff965ced07db8
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,376 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present {
+		if _, ok := value.(map[string]any); !ok {
+			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
+		}
+	}
+	if value, present := front["boundaries_opt_out"]; present {
+		if _, ok := value.(string); !ok {
+			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
+		}
+	}
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok {
+			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
+		}
+		defaults = mapping
+		for _, key := range zoneKeys {
+			if value, exists := defaults[key]; exists {
+				if _, e := Strings(value); e != nil {
+					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
+				}
+			}
+		}
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..7e0b7a8a3cf7fb413e8cee04e99e21a3e299f211
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,117 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	var current model.Module
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			current, found = m, true
+			break
+		}
+	}
+	if !found {
+		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive"))
+	if os.IsNotExist(err) {
+		entries = nil
+	} else if err != nil {
+		return nil, err
+	}
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
+			continue
+		}
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		archiveModule, ok := doc.Frontmatter["module"].(string)
+		if !ok || archiveModule == "" { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: invalid module target", rel)) }
+		var ids []string
+		seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name))
+			if e != nil {
+				return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+			}
+			for _, item := range items {
+				if !seen[item.ID] {
+					seen[item.ID] = true
+					ids = append(ids, item.ID)
+				}
+			}
+		}
+		if archiveModule != module || len(ids) == 0 {
+			continue
+		}
+		if id != "" {
+			if !seen[id] {
+				continue
+			}
+			ids = []string{id}
+		} else {
+			sort.Strings(ids)
+		}
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body, "Motivation"), Decisions: section(doc.Body, "Decisions"), RemovedDecisions: section(doc.Body, "Removed decisions"), Provenance: section(doc.Body, "Provenance")}
+		for _, line := range strings.Split(entry.Provenance, "\n") {
+			fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-"))
+			if len(fields) < 2 || fields[0] != "impl" {
+				continue
+			}
+			date, e := time.Parse(time.RFC3339Nano, fields[1])
+			if e == nil && date.After(entry.ImplementationDate) {
+				entry.ImplementationDate = date
+			}
+		}
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false
+		body, _ := contracts.Section(current.Document.Body, "Specs")
+		items, _ := contracts.ParseItems(body)
+		for _, item := range items {
+			if item.ID == id {
+				known = true
+				break
+			}
+		}
+		if !known {
+			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
+		}
+	}
+	sort.Slice(result, func(i, j int) bool {
+		a, b := result[i], result[j]
+		if !a.ImplementationDate.Equal(b.ImplementationDate) {
+			return a.ImplementationDate.After(b.ImplementationDate)
+		}
+		return a.Path < b.Path
+	})
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..7273b695a8af9b8e36147197fb379478b7df7513
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,343 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return c, fixture.Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings {
+			keys = append(keys, key)
+		}
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/history/history.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..daf19b2ba7b8d622a14a79fb201ff965ced07db8
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,376 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present {
+		if _, ok := value.(map[string]any); !ok {
+			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
+		}
+	}
+	if value, present := front["boundaries_opt_out"]; present {
+		if _, ok := value.(string); !ok {
+			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
+		}
+	}
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok {
+			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
+		}
+		defaults = mapping
+		for _, key := range zoneKeys {
+			if value, exists := defaults[key]; exists {
+				if _, e := Strings(value); e != nil {
+					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
+				}
+			}
+		}
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..1756ff137d1f6e751479db47be7e57bddb39b1bb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,118 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	var current model.Module
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			current, found = m, true
+			break
+		}
+	}
+	if !found {
+		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive"))
+	if os.IsNotExist(err) {
+		entries = nil
+	} else if err != nil {
+		return nil, err
+	}
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
+			continue
+		}
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		archiveModule, ok := doc.Frontmatter["module"].(string)
+		if !ok || archiveModule == "" { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: invalid module target", rel)) }
+		var ids []string
+		seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name))
+			if e != nil {
+				return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+			}
+			for _, item := range items {
+				if !seen[item.ID] {
+					seen[item.ID] = true
+					ids = append(ids, item.ID)
+				}
+			}
+		}
+		if archiveModule != module || len(ids) == 0 {
+			continue
+		}
+		if id != "" {
+			if !seen[id] {
+				continue
+			}
+			ids = []string{id}
+		} else {
+			sort.Strings(ids)
+		}
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body, "Motivation"), Decisions: section(doc.Body, "Decisions"), RemovedDecisions: section(doc.Body, "Removed decisions"), Provenance: section(doc.Body, "Provenance")}
+		if entry.RemovedDecisions == "" { entry.RemovedDecisions = section(doc.Body, "Removed Decisions") }
+		for _, line := range strings.Split(entry.Provenance, "\n") {
+			fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-"))
+			if len(fields) < 2 || fields[0] != "impl" {
+				continue
+			}
+			date, e := time.Parse(time.RFC3339Nano, fields[1])
+			if e == nil && date.After(entry.ImplementationDate) {
+				entry.ImplementationDate = date
+			}
+		}
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false
+		body, _ := contracts.Section(current.Document.Body, "Specs")
+		items, _ := contracts.ParseItems(body)
+		for _, item := range items {
+			if item.ID == id {
+				known = true
+				break
+			}
+		}
+		if !known {
+			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
+		}
+	}
+	sort.Slice(result, func(i, j int) bool {
+		a, b := result[i], result[j]
+		if !a.ImplementationDate.Equal(b.ImplementationDate) {
+			return a.ImplementationDate.After(b.ImplementationDate)
+		}
+		return a.Path < b.Path
+	})
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..7273b695a8af9b8e36147197fb379478b7df7513
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,343 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return c, fixture.Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings {
+			keys = append(keys, key)
+		}
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

codex
The implementation now covers the retained API routes, including coded lookup and parse errors, graph checks, archive ordering, and Git based stale text recovery. I’m doing a final formatting and scope check, then I’ll record the input-only verification limit in `DRILL_NOTES.md`.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..ec32bf8825cb42a234d4a05ad600041fd4bfa159
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,9 @@
+# Project rebuild notes
+
+## Contradicted Decisions
+
+None identified in the retained project inputs.
+
+## Gaps
+
+- no spec (project): The native Go eval suite could not run in this input-only workspace because the build manifest and dependency implementations are deliberately absent. The separate native verifier must compile and run the retained evals.
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..daf19b2ba7b8d622a14a79fb201ff965ced07db8
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,376 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present {
+		if _, ok := value.(map[string]any); !ok {
+			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
+		}
+	}
+	if value, present := front["boundaries_opt_out"]; present {
+		if _, ok := value.(string); !ok {
+			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
+		}
+	}
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	base, err := canonical(root)
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok {
+			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
+		}
+		defaults = mapping
+		for _, key := range zoneKeys {
+			if value, exists := defaults[key]; exists {
+				if _, e := Strings(value); e != nil {
+					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
+				}
+			}
+		}
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..1756ff137d1f6e751479db47be7e57bddb39b1bb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,118 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	var current model.Module
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			current, found = m, true
+			break
+		}
+	}
+	if !found {
+		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive"))
+	if os.IsNotExist(err) {
+		entries = nil
+	} else if err != nil {
+		return nil, err
+	}
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
+			continue
+		}
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		archiveModule, ok := doc.Frontmatter["module"].(string)
+		if !ok || archiveModule == "" { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: invalid module target", rel)) }
+		var ids []string
+		seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name))
+			if e != nil {
+				return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+			}
+			for _, item := range items {
+				if !seen[item.ID] {
+					seen[item.ID] = true
+					ids = append(ids, item.ID)
+				}
+			}
+		}
+		if archiveModule != module || len(ids) == 0 {
+			continue
+		}
+		if id != "" {
+			if !seen[id] {
+				continue
+			}
+			ids = []string{id}
+		} else {
+			sort.Strings(ids)
+		}
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body, "Motivation"), Decisions: section(doc.Body, "Decisions"), RemovedDecisions: section(doc.Body, "Removed decisions"), Provenance: section(doc.Body, "Provenance")}
+		if entry.RemovedDecisions == "" { entry.RemovedDecisions = section(doc.Body, "Removed Decisions") }
+		for _, line := range strings.Split(entry.Provenance, "\n") {
+			fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-"))
+			if len(fields) < 2 || fields[0] != "impl" {
+				continue
+			}
+			date, e := time.Parse(time.RFC3339Nano, fields[1])
+			if e == nil && date.After(entry.ImplementationDate) {
+				entry.ImplementationDate = date
+			}
+		}
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false
+		body, _ := contracts.Section(current.Document.Body, "Specs")
+		items, _ := contracts.ParseItems(body)
+		for _, item := range items {
+			if item.ID == id {
+				known = true
+				break
+			}
+		}
+		if !known {
+			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
+		}
+	}
+	sort.Slice(result, func(i, j int) bool {
+		a, b := result[i], result[j]
+		if !a.ImplementationDate.Equal(b.ImplementationDate) {
+			return a.ImplementationDate.After(b.ImplementationDate)
+		}
+		return a.Path < b.Path
+	})
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..7273b695a8af9b8e36147197fb379478b7df7513
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,343 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return c, fixture.Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings {
+			keys = append(keys, key)
+		}
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

exec
/bin/zsh -lc 'gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/queue/*.go packages/project/internal/history/history.go && git diff --check && git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'rg --files packages/project/internal | sort' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 530ms:
packages/project/internal/fixture/fixture.go
packages/project/internal/history/history.go
packages/project/internal/queue/check.go
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go

 succeeded in 559ms:
?? DRILL_NOTES.md
?? packages/project/internal/

exec
/bin/zsh -lc "sed -n '390,445p' packages/project/evals/queue_test.go; wc -l packages/project/internal/**/*.go DRILL_NOTES.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 471ms:
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
     378 packages/project/internal/fixture/fixture.go
     122 packages/project/internal/history/history.go
     148 packages/project/internal/queue/check.go
     345 packages/project/internal/queue/queue.go
     164 packages/project/internal/queue/stale.go
       9 DRILL_NOTES.md
    1166 total

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
new file mode 100644
index 0000000000000000000000000000000000000000..ec32bf8825cb42a234d4a05ad600041fd4bfa159
--- /dev/null
+++ b/DRILL_NOTES.md
@@ -0,0 +1,9 @@
+# Project rebuild notes
+
+## Contradicted Decisions
+
+None identified in the retained project inputs.
+
+## Gaps
+
+- no spec (project): The native Go eval suite could not run in this input-only workspace because the build manifest and dependency implementations are deliberately absent. The separate native verifier must compile and run the retained evals.
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
new file mode 100644
index 0000000000000000000000000000000000000000..490df267562a6b4872e404087b347a653dd76f8b
--- /dev/null
+++ b/packages/project/internal/fixture/fixture.go
@@ -0,0 +1,380 @@
+package fixture
+
+import (
+	"fmt"
+	"os"
+	"os/exec"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func Code(code, subject string, err error) error {
+	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
+}
+
+func canonical(root string) (string, error) {
+	absolute, err := filepath.Abs(root)
+	if err != nil {
+		return "", err
+	}
+	return filepath.EvalSymlinks(absolute)
+}
+
+func Config(data []byte) (map[string]any, error) {
+	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
+	if err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+	}
+	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+	}
+	return doc.Frontmatter, nil
+}
+
+func Strings(value any) ([]string, error) {
+	list, ok := value.([]any)
+	if !ok {
+		return nil, fmt.Errorf("expected list of strings, got %T", value)
+	}
+	out := make([]string, 0, len(list))
+	for _, v := range list {
+		s, ok := v.(string)
+		if !ok {
+			return nil, fmt.Errorf("expected string, got %T", v)
+		}
+		out = append(out, s)
+	}
+	return out, nil
+}
+
+var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}
+
+func ZoneKeys() []string { return zoneKeys }
+
+func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
+	m := model.Module{Document: contracts.Document{Frontmatter: front}}
+	name, ok := front["name"].(string)
+	if !ok {
+		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
+	}
+	m.Name = name
+	prefix, ok := front["id_prefix"].(string)
+	if !ok {
+		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
+	}
+	m.IDPrefix = prefix
+	if value, present := front["hooks"]; present {
+		if _, ok := value.(map[string]any); !ok {
+			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
+		}
+	}
+	if value, present := front["boundaries_opt_out"]; present {
+		if _, ok := value.(string); !ok {
+			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
+		}
+	}
+	if v, present := front["depends_on"]; present {
+		deps, err := Strings(v)
+		if err != nil {
+			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
+		}
+		m.DependsOn = deps
+	}
+	get := func(key string) ([]string, error) {
+		v, present := front[key]
+		if !present {
+			v, present = defaults[key]
+		}
+		if !present {
+			return nil, nil
+		}
+		list, err := Strings(v)
+		if err != nil {
+			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
+		}
+		return list, nil
+	}
+	var err error
+	if m.Zones.Protected, err = get("protected"); err != nil {
+		return m, err
+	}
+	if m.Zones.Evals, err = get("evals"); err != nil {
+		return m, err
+	}
+	if m.Zones.Interface, err = get("interface"); err != nil {
+		return m, err
+	}
+	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
+		return m, err
+	}
+	if m.Zones.Preserve, err = get("preserve"); err != nil {
+		return m, err
+	}
+	if m.Zones.Disposable, err = get("disposable"); err != nil {
+		return m, err
+	}
+	return m, nil
+}
+
+func specItems(doc contracts.Document, rel string) error {
+	section, ok := contracts.Section(doc.Body, "Specs")
+	if !ok {
+		return nil
+	}
+	items, err := contracts.ParseItems(section)
+	if err != nil {
+		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
+			return Code("C05", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	seen := map[string]bool{}
+	for _, item := range items {
+		if seen[item.ID] {
+			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
+		}
+		seen[item.ID] = true
+	}
+	return nil
+}
+
+func load(root, revision string) (model.Snapshot, error) {
+	var snap model.Snapshot
+	var base string
+	var err error
+	if revision == "" { base, err = filepath.Abs(root) } else { base, err = canonical(root) }
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+	}
+	snap.Root = base
+	var configBytes []byte
+	if revision == "" {
+		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
+	} else {
+		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
+	}
+	if err != nil {
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+	}
+	snap.Config, err = Config(configBytes)
+	if err != nil {
+		return snap, err
+	}
+	globs, _ := Strings(snap.Config["modules"])
+	defaults := map[string]any{}
+	if value, present := snap.Config["module_defaults"]; present {
+		mapping, ok := value.(map[string]any)
+		if !ok {
+			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
+		}
+		defaults = mapping
+		for _, key := range zoneKeys {
+			if value, exists := defaults[key]; exists {
+				if _, e := Strings(value); e != nil {
+					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
+				}
+			}
+		}
+	}
+	var dirs []string
+	if revision == "" {
+		for _, glob := range globs {
+			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
+			if e != nil {
+				return snap, Code("C01", "modules", e)
+			}
+			for _, match := range matches {
+				info, e := os.Stat(match)
+				if e == nil && info.IsDir() {
+					rel, _ := filepath.Rel(base, match)
+					dirs = append(dirs, filepath.ToSlash(rel))
+				}
+			}
+		}
+	} else {
+		// Tree names come from the selected revision; working tree directories are irrelevant.
+		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
+		out, e := cmd.Output()
+		if e != nil {
+			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
+		}
+		allDirs := map[string]bool{}
+		for _, file := range strings.Split(string(out), "\n") {
+			if file != "" {
+				parts := strings.Split(file, "/")
+				for i := 1; i < len(parts); i++ {
+					allDirs[strings.Join(parts[:i], "/")] = true
+				}
+			}
+		}
+		for dir := range allDirs {
+			for _, glob := range globs {
+				if contracts.MatchGlob(glob, dir) {
+					dirs = append(dirs, dir)
+					break
+				}
+			}
+		}
+	}
+	sort.Strings(dirs)
+	seen := map[string]bool{}
+	for _, dir := range dirs {
+		if seen[dir] {
+			continue
+		}
+		seen[dir] = true
+		rel := path.Join(dir, "SPEC.md")
+		var data []byte
+		if revision == "" {
+			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
+		} else {
+			data, err = gitstore.ReadAt(base, revision, rel)
+		}
+		if err != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		if e = specItems(doc, rel); e != nil {
+			return snap, e
+		}
+		module, e := ValidateMetadata(doc.Frontmatter, defaults)
+		if e != nil {
+			return snap, e
+		}
+		module.Dir, module.Document = dir, doc
+		snap.Modules = append(snap.Modules, module)
+	}
+	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
+	return snap, nil
+}
+
+func Load(root string) (model.Snapshot, error)             { return load(root, "") }
+func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }
+
+func ValidateGraph(modules []model.Module) error {
+	byName, prefixes := map[string]model.Module{}, map[string]bool{}
+	for _, m := range modules {
+		if m.Name == "" {
+			return fmt.Errorf("empty module name")
+		}
+		if m.IDPrefix == "" {
+			return fmt.Errorf("empty module prefix")
+		}
+		if _, ok := byName[m.Name]; ok {
+			return fmt.Errorf("duplicate module name %s", m.Name)
+		}
+		if prefixes[m.IDPrefix] {
+			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
+		}
+		byName[m.Name], prefixes[m.IDPrefix] = m, true
+	}
+	for _, m := range modules {
+		for _, dep := range m.DependsOn {
+			if dep == m.Name {
+				return fmt.Errorf("self dependency %s", m.Name)
+			}
+			if _, ok := byName[dep]; !ok {
+				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+			}
+		}
+	}
+	state := map[string]int{}
+	var visit func(string) error
+	visit = func(name string) error {
+		if state[name] == 1 {
+			return fmt.Errorf("dependency cycle at %s", name)
+		}
+		if state[name] == 2 {
+			return nil
+		}
+		state[name] = 1
+		for _, dep := range byName[name].DependsOn {
+			if err := visit(dep); err != nil {
+				return err
+			}
+		}
+		state[name] = 2
+		return nil
+	}
+	for _, m := range modules {
+		if err := visit(m.Name); err != nil {
+			return err
+		}
+	}
+	return nil
+}
+
+var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)
+
+func NewChange(root, module, slug string) (string, error) {
+	snap, err := Load(root)
+	if err != nil {
+		return "", err
+	}
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			found = true
+			break
+		}
+	}
+	if !found {
+		return "", fmt.Errorf("unknown module %s", module)
+	}
+	if !slugName.MatchString(slug) {
+		return "", fmt.Errorf("invalid slug %q", slug)
+	}
+	max := 0
+	for _, folder := range []string{"changes", "changes/archive"} {
+		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
+		if e != nil && !os.IsNotExist(e) {
+			return "", e
+		}
+		for _, entry := range entries {
+			if entry.IsDir() {
+				continue
+			}
+			parts := queuedName.FindStringSubmatch(entry.Name())
+			if parts == nil {
+				continue
+			}
+			if parts[2] == slug {
+				return "", fmt.Errorf("duplicate slug %s", slug)
+			}
+			n, _ := strconv.Atoi(parts[1])
+			if n > max {
+				max = n
+			}
+		}
+	}
+	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
+	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
+	_, err = file.WriteString(content)
+	closeErr := file.Close()
+	if err != nil {
+		return "", fmt.Errorf("create %s: %w", rel, err)
+	}
+	if closeErr != nil {
+		return "", fmt.Errorf("create %s: %w", rel, closeErr)
+	}
+	return rel, nil
+}
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
new file mode 100644
index 0000000000000000000000000000000000000000..1756ff137d1f6e751479db47be7e57bddb39b1bb
--- /dev/null
+++ b/packages/project/internal/history/history.go
@@ -0,0 +1,118 @@
+package history
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"sort"
+	"strings"
+	"time"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func section(body, name string) string { value, _ := contracts.Section(body, name); return value }
+
+func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	var current model.Module
+	found := false
+	for _, m := range snap.Modules {
+		if m.Name == module {
+			current, found = m, true
+			break
+		}
+	}
+	if !found {
+		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive"))
+	if os.IsNotExist(err) {
+		entries = nil
+	} else if err != nil {
+		return nil, err
+	}
+	result := []model.SpecHistoryEntry{}
+	for _, file := range entries {
+		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
+			continue
+		}
+		rel := path.Join("changes/archive", file.Name())
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		doc, e := contracts.ParseDocument(string(data))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+		}
+		archiveModule, ok := doc.Frontmatter["module"].(string)
+		if !ok || archiveModule == "" { return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: invalid module target", rel)) }
+		var ids []string
+		seen := map[string]bool{}
+		for _, name := range []string{"Added", "Modified"} {
+			items, e := contracts.ParseItems(section(doc.Body, name))
+			if e != nil {
+				return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
+			}
+			for _, item := range items {
+				if !seen[item.ID] {
+					seen[item.ID] = true
+					ids = append(ids, item.ID)
+				}
+			}
+		}
+		if archiveModule != module || len(ids) == 0 {
+			continue
+		}
+		if id != "" {
+			if !seen[id] {
+				continue
+			}
+			ids = []string{id}
+		} else {
+			sort.Strings(ids)
+		}
+		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body, "Motivation"), Decisions: section(doc.Body, "Decisions"), RemovedDecisions: section(doc.Body, "Removed decisions"), Provenance: section(doc.Body, "Provenance")}
+		if entry.RemovedDecisions == "" { entry.RemovedDecisions = section(doc.Body, "Removed Decisions") }
+		for _, line := range strings.Split(entry.Provenance, "\n") {
+			fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-"))
+			if len(fields) < 2 || fields[0] != "impl" {
+				continue
+			}
+			date, e := time.Parse(time.RFC3339Nano, fields[1])
+			if e == nil && date.After(entry.ImplementationDate) {
+				entry.ImplementationDate = date
+			}
+		}
+		result = append(result, entry)
+	}
+	if id != "" && len(result) == 0 {
+		known := false
+		body, _ := contracts.Section(current.Document.Body, "Specs")
+		items, _ := contracts.ParseItems(body)
+		for _, item := range items {
+			if item.ID == id {
+				known = true
+				break
+			}
+		}
+		if !known {
+			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
+		}
+	}
+	sort.Slice(result, func(i, j int) bool {
+		a, b := result[i], result[j]
+		if !a.ImplementationDate.Equal(b.ImplementationDate) {
+			return a.ImplementationDate.After(b.ImplementationDate)
+		}
+		return a.Path < b.Path
+	})
+	return result, nil
+}
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
new file mode 100644
index 0000000000000000000000000000000000000000..73d3b04b5426054ab5e9706b22e79e120e3549e6
--- /dev/null
+++ b/packages/project/internal/queue/check.go
@@ -0,0 +1,148 @@
+package queue
+
+import (
+	"fmt"
+	"path"
+	"regexp"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
+var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)
+
+func diagnostic(err error, fallback, file string) model.Diagnostic {
+	code, subject, ok := contracts.ErrorCode(err)
+	if !ok {
+		code = fallback
+	}
+	if file == "" && strings.HasSuffix(subject, ".md") {
+		file = subject
+	}
+	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
+}
+
+func Check(root string) []model.Diagnostic {
+	var out []model.Diagnostic
+	add := func(code, message, file string) {
+		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
+	}
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return []model.Diagnostic{diagnostic(err, "C03", "")}
+	}
+	if err = fixture.ValidateGraph(snap.Modules); err != nil {
+		add("C02", err.Error(), "")
+	}
+	prefixes := map[string]string{}
+	for _, m := range snap.Modules {
+		prefixes[m.IDPrefix] = m.Name
+	}
+	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
+		if requireNormative && !normative.MatchString(item.Body) {
+			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
+		}
+		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
+			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
+		}
+		if !strings.HasPrefix(item.ID, "new:") {
+			prefix := strings.SplitN(item.ID, "-", 2)[0]
+			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
+				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+			}
+		}
+	}
+	for _, m := range snap.Modules {
+		file := path.Join(m.Dir, "SPEC.md")
+		section, _ := contracts.Section(m.Document.Body, "Specs")
+		items, e := contracts.ParseItems(section)
+		if e != nil {
+			out = append(out, diagnostic(e, "C03", file))
+			continue
+		}
+		seen := map[string]bool{}
+		for _, item := range items {
+			if seen[item.ID] {
+				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
+			}
+			seen[item.ID] = true
+			validateItem(item, m, file, true)
+		}
+	}
+	changes, err := Queue(root)
+	if err != nil {
+		out = append(out, diagnostic(err, "C03", ""))
+		return out
+	}
+	for _, c := range changes {
+		var m model.Module
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
+		if strings.TrimSpace(motivation) == "" {
+			add("C07", "missing Motivation", c.Path)
+		}
+		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
+			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
+				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
+				continue
+			}
+		}
+		declared := map[string]bool{}
+		for _, item := range c.Added {
+			if strings.HasPrefix(item.ID, "new:") {
+				declared[item.ID] = true
+			}
+		}
+		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
+			validateItem(item, m, c.Path, true)
+		}
+		for _, item := range c.Removed {
+			validateItem(item, m, c.Path, false)
+		}
+		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
+			for _, handle := range handles.FindAllString(item.Body, -1) {
+				if !declared[handle] {
+					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
+				}
+			}
+		}
+		b, e := bases(c.Document)
+		if e != nil {
+			out = append(out, diagnostic(e, "C08", c.Path))
+		} else {
+			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+				prior, ok := itemByID(m, item.ID)
+				if !ok {
+					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
+					continue
+				}
+				recorded, present := b[item.ID]
+				hash, _ := recorded.(string)
+				if !present || hash == "" {
+					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
+				} else if hash != contracts.SpecHash(prior) {
+					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
+				}
+			}
+		}
+		if e = Apply(&snap, c); e != nil {
+			out = append(out, diagnostic(e, "C03", c.Path))
+			break
+		}
+		for _, candidate := range snap.Modules {
+			if candidate.Name == c.Module {
+				m = candidate
+				break
+			}
+		}
+	}
+	return out
+}
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
new file mode 100644
index 0000000000000000000000000000000000000000..7273b695a8af9b8e36147197fb379478b7df7513
--- /dev/null
+++ b/packages/project/internal/queue/queue.go
@@ -0,0 +1,343 @@
+package queue
+
+import (
+	"fmt"
+	"os"
+	"path"
+	"path/filepath"
+	"regexp"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
+
+func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
+	c := model.Change{Number: number, Slug: slug, Path: rel}
+	doc, err := contracts.ParseDocument(string(data))
+	if err != nil {
+		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+	}
+	c.Document = doc
+	module, ok := doc.Frontmatter["module"].(string)
+	if !ok || module == "" {
+		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
+	}
+	c.Module = module
+	seen := map[string]bool{}
+	for _, field := range []string{"Added", "Modified", "Removed"} {
+		section, _ := contracts.Section(doc.Body, field)
+		items, err := contracts.ParseItems(section)
+		if err != nil {
+			if strings.Contains(strings.ToLower(err.Error()), "duplicate") { return c, fixture.Code("C05", rel, fmt.Errorf("%s: %w", rel, err)) }
+			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
+		}
+		for _, item := range items {
+			if seen[item.ID] {
+				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
+			}
+			seen[item.ID] = true
+		}
+		switch field {
+		case "Added":
+			c.Added = items
+		case "Modified":
+			c.Modified = items
+		case "Removed":
+			c.Removed = items
+		}
+	}
+	return c, nil
+}
+
+func Queue(root string) ([]model.Change, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return nil, err
+	}
+	modules := map[string]bool{}
+	for _, m := range snap.Modules {
+		modules[m.Name] = true
+	}
+	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
+	if os.IsNotExist(err) {
+		return []model.Change{}, nil
+	}
+	if err != nil {
+		return nil, fixture.Code("C03", "changes", err)
+	}
+	var changes []model.Change
+	slugs := map[string]bool{}
+	for _, entry := range entries {
+		if entry.IsDir() {
+			continue
+		}
+		match := pendingName.FindStringSubmatch(entry.Name())
+		if match == nil {
+			continue
+		}
+		rel := path.Join("changes", entry.Name())
+		if slugs[match[2]] {
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+		}
+		slugs[match[2]] = true
+		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
+		if e != nil {
+			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
+		}
+		num, _ := strconv.Atoi(match[1])
+		c, e := ParseChange(data, rel, num, match[2])
+		if e != nil {
+			return nil, e
+		}
+		if !modules[c.Module] {
+			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
+		}
+		changes = append(changes, c)
+	}
+	sort.Slice(changes, func(i, j int) bool {
+		if changes[i].Number != changes[j].Number {
+			return changes[i].Number < changes[j].Number
+		}
+		return changes[i].Slug < changes[j].Slug
+	})
+	return changes, nil
+}
+
+func Resolve(changes []model.Change, selector string) (int, error) {
+	idx := -1
+	for i, c := range changes {
+		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
+		if match {
+			if idx >= 0 {
+				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
+			}
+			idx = i
+		}
+	}
+	if idx < 0 {
+		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+	}
+	return idx, nil
+}
+
+func itemList(m model.Module) ([]contracts.Item, error) {
+	section, _ := contracts.Section(m.Document.Body, "Specs")
+	return contracts.ParseItems(section)
+}
+
+func itemByID(m model.Module, id string) (contracts.Item, bool) {
+	items, err := itemList(m)
+	if err != nil {
+		return contracts.Item{}, false
+	}
+	for _, item := range items {
+		if item.ID == id {
+			return item, true
+		}
+	}
+	return contracts.Item{}, false
+}
+
+func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
+	snap, err := fixture.Load(root)
+	if err != nil {
+		return snap, err
+	}
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return snap, fixture.Code("C02", "", err)
+	}
+	for i := 0; i < target; i++ {
+		if err := Apply(&snap, changes[i]); err != nil {
+			return snap, err
+		}
+	}
+	return snap, nil
+}
+
+func Project(root, through string) (model.Snapshot, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, err
+	}
+	limit := len(changes)
+	if through != "" {
+		idx, e := Resolve(changes, through)
+		if e != nil {
+			return model.Snapshot{}, e
+		}
+		limit = idx + 1
+	}
+	snap, err := Before(root, changes, limit)
+	if err != nil {
+		return snap, err
+	}
+	return snap, nil
+}
+
+func Apply(snap *model.Snapshot, change model.Change) error {
+	i := -1
+	for n, m := range snap.Modules {
+		if m.Name == change.Module {
+			i = n
+			break
+		}
+	}
+	if i < 0 {
+		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
+	}
+	m := snap.Modules[i]
+	if value, ok := change.Document.Frontmatter["set"]; ok {
+		settings, ok := value.(map[string]any)
+		if !ok {
+			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+		}
+		keys := make([]string, 0, len(settings))
+		for key := range settings {
+			keys = append(keys, key)
+		}
+		sort.Strings(keys)
+		for _, key := range keys {
+			v := settings[key]
+			switch key {
+			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
+				if _, err := fixture.Strings(v); err != nil {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
+				}
+			case "boundaries_opt_out":
+				if _, ok := v.(string); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			case "hooks":
+				if _, ok := v.(map[string]any); !ok {
+					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
+				}
+			default:
+				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
+			}
+			m.Document.Frontmatter[key] = v
+		}
+		defaults, _ := snap.Config["module_defaults"].(map[string]any)
+		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
+		if err != nil {
+			return err
+		}
+		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
+	}
+	items, err := itemList(m)
+	if err != nil {
+		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
+	}
+	for _, edit := range change.Modified {
+		found := false
+		for j := range items {
+			if items[j].ID == edit.ID {
+				items[j] = edit
+				found = true
+				break
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
+		}
+	}
+	for _, edit := range change.Removed {
+		found := false
+		next := items[:0]
+		for _, item := range items {
+			if item.ID == edit.ID {
+				found = true
+			} else {
+				next = append(next, item)
+			}
+		}
+		if !found {
+			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
+		}
+		items = next
+	}
+	items = append(items, change.Added...)
+	var body strings.Builder
+	for _, item := range items {
+		body.WriteString(contracts.RenderItem(item))
+		if !strings.HasSuffix(body.String(), "\n") {
+			body.WriteByte('\n')
+		}
+		if !strings.HasSuffix(body.String(), "\n\n") {
+			body.WriteByte('\n')
+		}
+	}
+	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
+	snap.Modules[i] = m
+	if err := fixture.ValidateGraph(snap.Modules); err != nil {
+		return fixture.Code("C02", change.Path, err)
+	}
+	return nil
+}
+
+func Target(root, selector string) (model.Snapshot, model.Change, error) {
+	changes, err := Queue(root)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	idx, err := Resolve(changes, selector)
+	if err != nil {
+		return model.Snapshot{}, model.Change{}, err
+	}
+	snap, err := Before(root, changes, idx)
+	return snap, changes[idx], err
+}
+
+func bases(doc contracts.Document) (map[string]any, error) {
+	v, exists := doc.Frontmatter["bases"]
+	if !exists {
+		return map[string]any{}, nil
+	}
+	m, ok := v.(map[string]any)
+	if !ok {
+		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
+	}
+	return m, nil
+}
+
+func stamp(root, selector string, missingOnly bool) error {
+	snap, c, err := Target(root, selector)
+	if err != nil {
+		return err
+	}
+	b, err := bases(c.Document)
+	if err != nil {
+		return err
+	}
+	var m model.Module
+	for _, module := range snap.Modules {
+		if module.Name == c.Module {
+			m = module
+			break
+		}
+	}
+	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
+		item, ok := itemByID(m, edit.ID)
+		if !ok {
+			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
+		}
+		if _, exists := b[edit.ID]; missingOnly && exists {
+			continue
+		}
+		b[edit.ID] = contracts.SpecHash(item)
+	}
+	c.Document.Frontmatter["bases"] = b
+	text, err := contracts.RenderDocument(c.Document)
+	if err != nil {
+		return err
+	}
+	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
+}
+
+func Stamp(root, selector string) error        { return stamp(root, selector, false) }
+func StampMissing(root, selector string) error { return stamp(root, selector, true) }
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
new file mode 100644
index 0000000000000000000000000000000000000000..a57adfe9e57a91e0bbf3aebfbebe5e855d846993
--- /dev/null
+++ b/packages/project/internal/queue/stale.go
@@ -0,0 +1,72 @@
+package queue
+
+import (
+	"os/exec"
+	"path"
+	"reflect"
+	"sort"
+	"strconv"
+	"strings"
+
+	"github.com/lurkshark/fireweed/packages/contracts"
+	"github.com/lurkshark/fireweed/packages/gitstore"
+	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
+	"github.com/lurkshark/fireweed/packages/project/model"
+)
+
+func gitLines(root string, args ...string) ([]string, error) {
+	cmd := exec.Command("git", append([]string{"-C",root},args...)...)
+	out, err := cmd.Output(); if err != nil { return nil, err }
+	var lines []string; for _, line := range strings.Split(strings.TrimSpace(string(out)),"\n") { if line != "" { lines = append(lines,line) } }; return lines,nil
+}
+
+func recordedBases(root, rev, file string) map[string]any {
+	data, err := gitstore.ReadAt(root, rev, file); if err != nil { return nil }
+	doc, err := contracts.ParseDocument(string(data)); if err != nil { return nil }
+	b, _ := doc.Frontmatter["bases"].(map[string]any); return b
+}
+
+func stampedRevision(root, file string) string {
+	commits, err := gitLines(root,"rev-list","HEAD","--",file); if err != nil { return "" }
+	for _, commit := range commits {
+		current := recordedBases(root,commit,file)
+		previous := recordedBases(root,commit+"^",file)
+		if !reflect.DeepEqual(current,previous) { return commit }
+	}
+	return ""
+}
+
+func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
+	snap, err := fixture.LoadAt(root,revision); if err != nil { return snap,err }
+	files, err := gitLines(root,"ls-tree","-r","--name-only",revision,"--","changes"); if err != nil { return snap,err }
+	var changes []model.Change
+	for _, file := range files {
+		if path.Dir(file) != "changes" { continue }
+		match := pendingName.FindStringSubmatch(path.Base(file)); if match == nil { continue }
+		data, e := gitstore.ReadAt(root,revision,file); if e != nil { return snap,e }
+		num, _ := strconv.Atoi(match[1]); c, e := ParseChange(data,file,num,match[2]); if e != nil { return snap,e }; changes = append(changes,c)
+	}
+	sort.Slice(changes,func(i,j int)bool{if changes[i].Number != changes[j].Number{return changes[i].Number<changes[j].Number};return changes[i].Slug<changes[j].Slug})
+	for _, c := range changes { if c.Path == target { return snap,nil }; if err := Apply(&snap,c); err != nil { return snap,err } }
+	return snap, nil
+}
+
+func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
+	snap, c, err := Target(root,selector); if err != nil { return nil,err }
+	b, err := bases(c.Document); if err != nil { return nil,err }
+	var module model.Module; for _, m := range snap.Modules { if m.Name == c.Module { module = m; break } }
+	result := []model.StaleSpec{}
+	for _, edit := range append(append([]contracts.Item{},c.Modified...),c.Removed...) {
+		recorded, exists := b[edit.ID]; if !exists { continue }
+		hash, _ := recorded.(string)
+		prior, found := itemByID(module,edit.ID); if !found { continue }
+		if hash == contracts.SpecHash(prior) { continue }
+		result = append(result,model.StaleSpec{ID:edit.ID,RecordedHash:hash,Current:contracts.RenderItem(prior)})
+	}
+	if len(result) == 0 { return result,nil }
+	revision := stampedRevision(snap.Root,c.Path); if revision == "" { return result,nil }
+	historical, err := projectedBeforeAt(snap.Root,revision,c.Path); if err != nil { return result,nil }
+	var historicalModule model.Module; for _, m := range historical.Modules { if m.Name == c.Module { historicalModule = m; break } }
+	for i := range result { prior, ok := itemByID(historicalModule,result[i].ID); if ok && contracts.SpecHash(prior) == result[i].RecordedHash { result[i].Before = contracts.RenderItem(prior); result[i].Recovered = true } }
+	return result,nil
+}

exec
/bin/zsh -lc 'gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/queue/*.go packages/project/internal/history/history.go && git diff --check && git status --short' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 503ms:
?? DRILL_NOTES.md
?? packages/project/internal/

codex
Rebuilt the project implementation under [packages/project/internal](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal) and wrote [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/DRILL_NOTES.md). Retained inputs were not edited, and nothing was committed.

`gofmt` and `git diff --check` passed. The retained evals could not run here because this input workspace has no Go module manifest or dependency implementations; the native verifier must run them.
tokens used
79,299


### Attempt 2 (exit 0)
Updated the disposable project implementation to address the reported failures: foreign spec ID wording, stale numeric base hashes, handle references in change prose, and the error codes and messages used by the app. Updated [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/DRILL_NOTES.md).

`gofmt` passed. I could not run the Go evals here because the supplied workspace has no Go module manifest or dependent implementations; the native verifier still needs to confirm the fixes.

Reading additional input from stdin...
OpenAI Codex v0.165.0
--------
workdir: /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
model: gpt-6-sol
provider: openai
approval: never
sandbox: workspace-write [workdir, /tmp, $TMPDIR]
reasoning effort: high
reasoning summaries: none
session id: 01a19a55-80f3-7ad3-ad0c-81616739ba42
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
An unknown selector MUST return U02 naming the requested change; an ambiguous slug MUST return C13.

### PRJ-067
When NewChange cannot create its queued file, it MUST return an error beginning
with `create <queued path>:` and retaining the filesystem cause, where the queued
path is repository-relative.

### PRJ-068
When Check rejects syntax or duplicate identities in a module SPEC.md, the
failure diagnostic MUST identify that document's repository-relative path in
Diagnostic.Path.

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
project eval: E01: failed case "Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_module_id" for PRJ-012
app eval: E01: failed case "Test_APP_019_NewChangeErrorsPreserveFiles" for APP-019
project stdout: ?   	github.com/lurkshark/fireweed/packages/project	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/fixture	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/history	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/internal/queue	[no test files]
?   	github.com/lurkshark/fireweed/packages/project/model	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/project/evals	1.928s

project stderr: exit status 1

project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics: === RUN   Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics
--- FAIL: Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics (0.03s)
project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_module_id: === RUN   Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_module_id
    queue_test.go:362: Check did not report C05 for foreign spec ID: []model.Diagnostic{model.Diagnostic{Code:"C05", Message:"foreign module ID BET-001 belongs to beta", Path:"packages/alpha/SPEC.md", Warning:false}}
--- FAIL: Test_PRJ_014_PRJ_015_PRJ_016_CheckDiagnostics/PRJ-012_foreign_module_id (0.00s)
project github.com/lurkshark/fireweed/packages/project/evals Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem: === RUN   Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem
    stale_test.go:144: entry 1 = model.StaleSpec{ID:"ALP-002", RecordedHash:"", Before:"", Current:"### ALP-002\nTwo MUST work.\n\n", Recovered:false}; want unrecovered ALP-002
--- FAIL: Test_PRJ_062_UnavailableHistoryStillReportsEveryStaleItem (0.01s)

app stdout: ?   	github.com/lurkshark/fireweed/packages/app	[no test files]
?   	github.com/lurkshark/fireweed/packages/app/internal/entry	[no test files]
FAIL
FAIL	github.com/lurkshark/fireweed/packages/app/evals	222.883s

app stderr: exit status 1

app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses: === RUN   Test_APP_037_StaticCauses
--- FAIL: Test_APP_037_StaticCauses (0.68s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/configuration: === RUN   Test_APP_037_StaticCauses/configuration
    diagnostic_causes_test.go:105: duplicate C01 prefix in "C01: C01: fireweed.yaml: invalid YAML frontmatter: yaml: line 2: did not find expected ',' or ']'. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.\n"
--- FAIL: Test_APP_037_StaticCauses/configuration (0.27s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/graph: === RUN   Test_APP_037_StaticCauses/graph
    diagnostic_causes_test.go:105: missing C02 cause "missing module" in plain output "C02: module beta dependency missing is missing. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.\n"
    diagnostic_causes_test.go:105: missing C02 cause "missing module" in {"schema":1,"data":null,"diagnostics":[{"code":"C02","message":"module beta dependency missing is missing. Provide a valid fireweed.yaml and correct module declarations in SPEC.md, then run modules again.","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/graph (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/handle: === RUN   Test_APP_037_StaticCauses/handle
    diagnostic_causes_test.go:105: status=0; want 1; stdout="Check passed\n" stderr=""
    diagnostic_causes_test.go:105: plain diagnostic="Check passed\n"; want code and message
    diagnostic_causes_test.go:105: missing C09 cause "undeclared handle" in plain output "Check passed\n"
    diagnostic_causes_test.go:105: status=0; want 1; stdout="{\"schema\":1,\"data\":\"Check passed\\n\",\"diagnostics\":[]}\n" stderr=""
    diagnostic_causes_test.go:105: data="Check passed\n"; want null
    diagnostic_causes_test.go:105: diagnostics=[]; want nonempty array: <nil>
--- FAIL: Test_APP_037_StaticCauses/handle (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/metadata: === RUN   Test_APP_037_StaticCauses/metadata
    diagnostic_causes_test.go:105: missing C12 cause "set must be a mapping" in plain output "C12: invalid set mapping\n"
    diagnostic_causes_test.go:105: missing C12 cause "set must be a mapping" in {"schema":1,"data":null,"diagnostics":[{"code":"C12","message":"invalid set mapping","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/metadata (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_037_StaticCauses/queue-identity: === RUN   Test_APP_037_StaticCauses/queue-identity
    diagnostic_causes_test.go:105: missing C13 cause "ambiguous change slug" in plain output "C13: duplicate pending slug repeat\n"
    diagnostic_causes_test.go:105: missing C13 cause "ambiguous change slug" in {"schema":1,"data":null,"diagnostics":[{"code":"C13","message":"duplicate pending slug repeat","warning":false}]}
--- FAIL: Test_APP_037_StaticCauses/queue-identity (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses: === RUN   Test_APP_038_LookupCauses
--- FAIL: Test_APP_038_LookupCauses (0.55s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses/stamp-change: === RUN   Test_APP_038_LookupCauses/stamp-change
    diagnostic_causes_test.go:156: duplicate U02 prefix in "U02: U02: pending change absent not found\n"
--- FAIL: Test_APP_038_LookupCauses/stamp-change (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_038_LookupCauses/coverage-change: === RUN   Test_APP_038_LookupCauses/coverage-change
    diagnostic_causes_test.go:156: duplicate U02 prefix in "U02: U02: pending change missing not found. Use a valid queued change alias.\n"
--- FAIL: Test_APP_038_LookupCauses/coverage-change (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors: === RUN   Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors
    list_cli_test.go:193: public Queue fixture cause: duplicate pending slug repeat
--- FAIL: Test_APP_116_APP_117_ListUsageConfigurationAndQueueErrors (0.00s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles (0.55s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain
    project_commands_test.go:397: diagnostic="P01: invalid slug \"Bad\". Use a valid module and an unused lowercase slug.\n"; want corrective U01
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "U01"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/invalid-slug/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain
    project_commands_test.go:397: diagnostic="P01: duplicate slug repeat. Use a valid module and an unused lowercase slug.\n"; want corrective C13
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-pending/JSON (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain
    project_commands_test.go:397: diagnostic="P01: duplicate slug repeat. Use a valid module and an unused lowercase slug.\n"; want corrective C13
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/plain (0.01s)
app github.com/lurkshark/fireweed/packages/app/evals Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON: === RUN   Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON
    project_commands_test.go:393: diagnostic code="P01"; want "C13"
--- FAIL: Test_APP_019_NewChangeErrorsPreserveFiles/duplicate-archived/JSON (0.01s)

checks stdout: ?   	github.com/lurkshark/fireweed/packages/checks	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/internal/engine	[no test files]
?   	github.com/lurkshark/fireweed/packages/checks/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/checks/evals	223.485s

checks stderr: 

runner stdout: ?   	github.com/lurkshark/fireweed/packages/runner	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/drill	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/invoke	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/next	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/internal/run	[no test files]
?   	github.com/lurkshark/fireweed/packages/runner/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/runner/evals	232.998s

runner stderr: 

workflow stdout: ?   	github.com/lurkshark/fireweed/packages/workflow	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/internal/lifecycle	[no test files]
?   	github.com/lurkshark/fireweed/packages/workflow/model	[no test files]
PASS
ok  	github.com/lurkshark/fireweed/packages/workflow/evals	91.708s

workflow stderr: 

codex
I’ll inspect the retained project API and evals, then fix the disposable project implementation and run the focused tests. I’ll also write the required drill notes.
exec
/bin/zsh -lc "pwd && rg --files -g 'AGENTS.md' -g 'packages/project/**' -g 'packages/contracts/**' -g 'packages/gitstore/**' -g 'packages/assets/**' -g 'DRILL_NOTES.md' | sort" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 437ms:
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
DRILL_NOTES.md
packages/assets/api.go
packages/assets/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
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
packages/project/internal/queue/check.go
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go
packages/project/model/model.go

exec
/bin/zsh -lc "cat 'packages/project/internal/history/history.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/internal/fixture/fixture.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/model/model.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/internal/queue/check.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/internal/queue/stale.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/api.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/internal/queue/queue.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'DRILL_NOTES.md'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 949ms:
package queue

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var normative = regexp.MustCompile(`\b(MUST|SHOULD|MAY|SHALL|REQUIRED|RECOMMENDED|OPTIONAL)\b`)
var handles = regexp.MustCompile(`\bnew:[a-z0-9-]+\b`)

func diagnostic(err error, fallback, file string) model.Diagnostic {
	code, subject, ok := contracts.ErrorCode(err)
	if !ok {
		code = fallback
	}
	if file == "" && strings.HasSuffix(subject, ".md") {
		file = subject
	}
	return model.Diagnostic{Code: code, Message: err.Error(), Path: file}
}

func Check(root string) []model.Diagnostic {
	var out []model.Diagnostic
	add := func(code, message, file string) {
		out = append(out, model.Diagnostic{Code: code, Message: message, Path: file})
	}
	snap, err := fixture.Load(root)
	if err != nil {
		return []model.Diagnostic{diagnostic(err, "C03", "")}
	}
	if err = fixture.ValidateGraph(snap.Modules); err != nil {
		add("C02", err.Error(), "")
	}
	prefixes := map[string]string{}
	for _, m := range snap.Modules {
		prefixes[m.IDPrefix] = m.Name
	}
	validateItem := func(item contracts.Item, module model.Module, file string, requireNormative bool) {
		if requireNormative && !normative.MatchString(item.Body) {
			add("C04", fmt.Sprintf("non-normative spec body for %s", item.ID), file)
		}
		if item.Tag != "" && item.Tag != "review" && !(strings.HasPrefix(item.Tag, "quarantine:") && strings.TrimSpace(strings.TrimPrefix(item.Tag, "quarantine:")) != "") {
			add("C06", fmt.Sprintf("invalid review/quarantine tag for %s", item.ID), file)
		}
		if !strings.HasPrefix(item.ID, "new:") {
			prefix := strings.SplitN(item.ID, "-", 2)[0]
			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
			}
		}
	}
	for _, m := range snap.Modules {
		file := path.Join(m.Dir, "SPEC.md")
		section, _ := contracts.Section(m.Document.Body, "Specs")
		items, e := contracts.ParseItems(section)
		if e != nil {
			out = append(out, diagnostic(e, "C03", file))
			continue
		}
		seen := map[string]bool{}
		for _, item := range items {
			if seen[item.ID] {
				add("C05", fmt.Sprintf("duplicate item ID %s", item.ID), file)
			}
			seen[item.ID] = true
			validateItem(item, m, file, true)
		}
	}
	changes, err := Queue(root)
	if err != nil {
		out = append(out, diagnostic(err, "C03", ""))
		return out
	}
	for _, c := range changes {
		var m model.Module
		for _, candidate := range snap.Modules {
			if candidate.Name == c.Module {
				m = candidate
				break
			}
		}
		motivation, _ := contracts.Section(c.Document.Body, "Motivation")
		if strings.TrimSpace(motivation) == "" {
			add("C07", "missing Motivation", c.Path)
		}
		if _, exists := c.Document.Frontmatter["eval_fix"]; exists {
			if _, fixes := c.Document.Frontmatter["fixes"]; fixes || len(c.Added)+len(c.Modified)+len(c.Removed) > 0 {
				add("C07", "eval_fix cannot be mixed with fixes or spec edits", c.Path)
				continue
			}
		}
		declared := map[string]bool{}
		for _, item := range c.Added {
			if strings.HasPrefix(item.ID, "new:") {
				declared[item.ID] = true
			}
		}
		for _, item := range append(append([]contracts.Item{}, c.Added...), c.Modified...) {
			validateItem(item, m, c.Path, true)
		}
		for _, item := range c.Removed {
			validateItem(item, m, c.Path, false)
		}
		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
			for _, handle := range handles.FindAllString(item.Body, -1) {
				if !declared[handle] {
					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
				}
			}
		}
		b, e := bases(c.Document)
		if e != nil {
			out = append(out, diagnostic(e, "C08", c.Path))
		} else {
			for _, item := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
				prior, ok := itemByID(m, item.ID)
				if !ok {
					add("C08", fmt.Sprintf("ID %s absent from preceding projection", item.ID), c.Path)
					continue
				}
				recorded, present := b[item.ID]
				hash, _ := recorded.(string)
				if !present || hash == "" {
					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
				} else if hash != contracts.SpecHash(prior) {
					add("C08", fmt.Sprintf("stale base for %s", item.ID), c.Path)
				}
			}
		}
		if e = Apply(&snap, c); e != nil {
			out = append(out, diagnostic(e, "C03", c.Path))
			break
		}
		for _, candidate := range snap.Modules {
			if candidate.Name == c.Module {
				m = candidate
				break
			}
		}
	}
	return out
}

 succeeded in 959ms:
package queue

import (
	"os/exec"
	"path"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func gitLines(root string, args ...string) ([]string, error) {
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var lines []string
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}

func recordedBases(root, rev, file string) map[string]any {
	data, err := gitstore.ReadAt(root, rev, file)
	if err != nil {
		return nil
	}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return nil
	}
	b, _ := doc.Frontmatter["bases"].(map[string]any)
	return b
}

func stampedRevision(root, file string) string {
	commits, err := gitLines(root, "rev-list", "HEAD", "--", file)
	if err != nil {
		return ""
	}
	for _, commit := range commits {
		current := recordedBases(root, commit, file)
		previous := recordedBases(root, commit+"^", file)
		if !reflect.DeepEqual(current, previous) {
			return commit
		}
	}
	return ""
}

func projectedBeforeAt(root, revision, target string) (model.Snapshot, error) {
	snap, err := fixture.LoadAt(root, revision)
	if err != nil {
		return snap, err
	}
	files, err := gitLines(root, "ls-tree", "-r", "--name-only", revision, "--", "changes")
	if err != nil {
		return snap, err
	}
	var changes []model.Change
	for _, file := range files {
		if path.Dir(file) != "changes" {
			continue
		}
		match := pendingName.FindStringSubmatch(path.Base(file))
		if match == nil {
			continue
		}
		data, e := gitstore.ReadAt(root, revision, file)
		if e != nil {
			return snap, e
		}
		num, _ := strconv.Atoi(match[1])
		c, e := ParseChange(data, file, num, match[2])
		if e != nil {
			return snap, e
		}
		changes = append(changes, c)
	}
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].Number != changes[j].Number {
			return changes[i].Number < changes[j].Number
		}
		return changes[i].Slug < changes[j].Slug
	})
	for _, c := range changes {
		if c.Path == target {
			return snap, nil
		}
		if err := Apply(&snap, c); err != nil {
			return snap, err
		}
	}
	return snap, nil
}

func StaleSpecs(root, selector string) ([]model.StaleSpec, error) {
	snap, c, err := Target(root, selector)
	if err != nil {
		return nil, err
	}
	b, err := bases(c.Document)
	if err != nil {
		return nil, err
	}
	var module model.Module
	for _, m := range snap.Modules {
		if m.Name == c.Module {
			module = m
			break
		}
	}
	result := []model.StaleSpec{}
	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
		recorded, exists := b[edit.ID]
		if !exists {
			continue
		}
		hash, _ := recorded.(string)
		prior, found := itemByID(module, edit.ID)
		if !found {
			continue
		}
		if hash == contracts.SpecHash(prior) {
			continue
		}
		result = append(result, model.StaleSpec{ID: edit.ID, RecordedHash: hash, Current: contracts.RenderItem(prior)})
	}
	if len(result) == 0 {
		return result, nil
	}
	revision := stampedRevision(snap.Root, c.Path)
	if revision == "" {
		return result, nil
	}
	historical, err := projectedBeforeAt(snap.Root, revision, c.Path)
	if err != nil {
		return result, nil
	}
	var historicalModule model.Module
	for _, m := range historical.Modules {
		if m.Name == c.Module {
			historicalModule = m
			break
		}
	}
	for i := range result {
		prior, ok := itemByID(historicalModule, result[i].ID)
		if ok && contracts.SpecHash(prior) == result[i].RecordedHash {
			result[i].Before = contracts.RenderItem(prior)
			result[i].Recovered = true
		}
	}
	return result, nil
}

 succeeded in 965ms:
package fixture

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/gitstore"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func Code(code, subject string, err error) error {
	return contracts.CodedError{Code: code, Subject: subject, Cause: err}
}

func canonical(root string) (string, error) {
	absolute, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	return filepath.EvalSymlinks(absolute)
}

func Config(data []byte) (map[string]any, error) {
	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
	if err != nil {
		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
	}
	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
	}
	return doc.Frontmatter, nil
}

func Strings(value any) ([]string, error) {
	list, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("expected list of strings, got %T", value)
	}
	out := make([]string, 0, len(list))
	for _, v := range list {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("expected string, got %T", v)
		}
		out = append(out, s)
	}
	return out, nil
}

var zoneKeys = []string{"protected", "evals", "interface", "append_only", "preserve", "disposable"}

func ZoneKeys() []string { return zoneKeys }

func ValidateMetadata(front, defaults map[string]any) (model.Module, error) {
	m := model.Module{Document: contracts.Document{Frontmatter: front}}
	name, ok := front["name"].(string)
	if !ok {
		return m, Code("C12", "name", fmt.Errorf("invalid module name type %T", front["name"]))
	}
	m.Name = name
	prefix, ok := front["id_prefix"].(string)
	if !ok {
		return m, Code("C12", "id_prefix", fmt.Errorf("invalid module prefix type %T", front["id_prefix"]))
	}
	m.IDPrefix = prefix
	if value, present := front["hooks"]; present {
		if _, ok := value.(map[string]any); !ok {
			return m, Code("C12", "hooks", fmt.Errorf("invalid hooks: expected mapping"))
		}
	}
	if value, present := front["boundaries_opt_out"]; present {
		if _, ok := value.(string); !ok {
			return m, Code("C12", "boundaries_opt_out", fmt.Errorf("invalid boundaries_opt_out: expected string"))
		}
	}
	if v, present := front["depends_on"]; present {
		deps, err := Strings(v)
		if err != nil {
			return m, Code("C12", "depends_on", fmt.Errorf("invalid depends_on: %w", err))
		}
		m.DependsOn = deps
	}
	get := func(key string) ([]string, error) {
		v, present := front[key]
		if !present {
			v, present = defaults[key]
		}
		if !present {
			return nil, nil
		}
		list, err := Strings(v)
		if err != nil {
			return nil, Code("C12", key, fmt.Errorf("invalid %s: %w", key, err))
		}
		return list, nil
	}
	var err error
	if m.Zones.Protected, err = get("protected"); err != nil {
		return m, err
	}
	if m.Zones.Evals, err = get("evals"); err != nil {
		return m, err
	}
	if m.Zones.Interface, err = get("interface"); err != nil {
		return m, err
	}
	if m.Zones.AppendOnly, err = get("append_only"); err != nil {
		return m, err
	}
	if m.Zones.Preserve, err = get("preserve"); err != nil {
		return m, err
	}
	if m.Zones.Disposable, err = get("disposable"); err != nil {
		return m, err
	}
	return m, nil
}

func specItems(doc contracts.Document, rel string) error {
	section, ok := contracts.Section(doc.Body, "Specs")
	if !ok {
		return nil
	}
	items, err := contracts.ParseItems(section)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			return Code("C05", rel, fmt.Errorf("%s: %w", rel, err))
		}
		return Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
	}
	seen := map[string]bool{}
	for _, item := range items {
		if seen[item.ID] {
			return Code("C05", rel, fmt.Errorf("%s: duplicate item ID %s", rel, item.ID))
		}
		seen[item.ID] = true
	}
	return nil
}

func load(root, revision string) (model.Snapshot, error) {
	var snap model.Snapshot
	var base string
	var err error
	if revision == "" {
		base, err = filepath.Abs(root)
	} else {
		base, err = canonical(root)
	}
	if err != nil {
		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
	}
	snap.Root = base
	var configBytes []byte
	if revision == "" {
		configBytes, err = os.ReadFile(filepath.Join(base, "fireweed.yaml"))
	} else {
		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
	}
	if err != nil {
		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
	}
	snap.Config, err = Config(configBytes)
	if err != nil {
		return snap, err
	}
	globs, _ := Strings(snap.Config["modules"])
	defaults := map[string]any{}
	if value, present := snap.Config["module_defaults"]; present {
		mapping, ok := value.(map[string]any)
		if !ok {
			return snap, Code("C12", "module_defaults", fmt.Errorf("invalid module_defaults: expected mapping"))
		}
		defaults = mapping
		for _, key := range zoneKeys {
			if value, exists := defaults[key]; exists {
				if _, e := Strings(value); e != nil {
					return snap, Code("C12", key, fmt.Errorf("invalid module_defaults %s: %w", key, e))
				}
			}
		}
	}
	var dirs []string
	if revision == "" {
		for _, glob := range globs {
			matches, e := filepath.Glob(filepath.Join(base, filepath.FromSlash(glob)))
			if e != nil {
				return snap, Code("C01", "modules", e)
			}
			for _, match := range matches {
				info, e := os.Stat(match)
				if e == nil && info.IsDir() {
					rel, _ := filepath.Rel(base, match)
					dirs = append(dirs, filepath.ToSlash(rel))
				}
			}
		}
	} else {
		// Tree names come from the selected revision; working tree directories are irrelevant.
		cmd := exec.Command("git", "-C", base, "ls-tree", "-r", "--name-only", revision)
		out, e := cmd.Output()
		if e != nil {
			return snap, fmt.Errorf("read revision %s tree: %w", revision, e)
		}
		allDirs := map[string]bool{}
		for _, file := range strings.Split(string(out), "\n") {
			if file != "" {
				parts := strings.Split(file, "/")
				for i := 1; i < len(parts); i++ {
					allDirs[strings.Join(parts[:i], "/")] = true
				}
			}
		}
		for dir := range allDirs {
			for _, glob := range globs {
				if contracts.MatchGlob(glob, dir) {
					dirs = append(dirs, dir)
					break
				}
			}
		}
	}
	sort.Strings(dirs)
	seen := map[string]bool{}
	for _, dir := range dirs {
		if seen[dir] {
			continue
		}
		seen[dir] = true
		rel := path.Join(dir, "SPEC.md")
		var data []byte
		if revision == "" {
			data, err = os.ReadFile(filepath.Join(base, filepath.FromSlash(rel)))
		} else {
			data, err = gitstore.ReadAt(base, revision, rel)
		}
		if err != nil {
			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
		}
		doc, e := contracts.ParseDocument(string(data))
		if e != nil {
			return snap, Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
		}
		if e = specItems(doc, rel); e != nil {
			return snap, e
		}
		module, e := ValidateMetadata(doc.Frontmatter, defaults)
		if e != nil {
			return snap, e
		}
		module.Dir, module.Document = dir, doc
		snap.Modules = append(snap.Modules, module)
	}
	sort.Slice(snap.Modules, func(i, j int) bool { return snap.Modules[i].Dir < snap.Modules[j].Dir })
	return snap, nil
}

func Load(root string) (model.Snapshot, error)             { return load(root, "") }
func LoadAt(root, revision string) (model.Snapshot, error) { return load(root, revision) }

func ValidateGraph(modules []model.Module) error {
	byName, prefixes := map[string]model.Module{}, map[string]bool{}
	for _, m := range modules {
		if m.Name == "" {
			return fmt.Errorf("empty module name")
		}
		if m.IDPrefix == "" {
			return fmt.Errorf("empty module prefix")
		}
		if _, ok := byName[m.Name]; ok {
			return fmt.Errorf("duplicate module name %s", m.Name)
		}
		if prefixes[m.IDPrefix] {
			return fmt.Errorf("duplicate module prefix %s", m.IDPrefix)
		}
		byName[m.Name], prefixes[m.IDPrefix] = m, true
	}
	for _, m := range modules {
		for _, dep := range m.DependsOn {
			if dep == m.Name {
				return fmt.Errorf("self dependency %s", m.Name)
			}
			if _, ok := byName[dep]; !ok {
				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
			}
		}
	}
	state := map[string]int{}
	var visit func(string) error
	visit = func(name string) error {
		if state[name] == 1 {
			return fmt.Errorf("dependency cycle at %s", name)
		}
		if state[name] == 2 {
			return nil
		}
		state[name] = 1
		for _, dep := range byName[name].DependsOn {
			if err := visit(dep); err != nil {
				return err
			}
		}
		state[name] = 2
		return nil
	}
	for _, m := range modules {
		if err := visit(m.Name); err != nil {
			return err
		}
	}
	return nil
}

var queuedName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)
var slugName = regexp.MustCompile(`^[a-z0-9-]+$`)

func NewChange(root, module, slug string) (string, error) {
	snap, err := Load(root)
	if err != nil {
		return "", err
	}
	found := false
	for _, m := range snap.Modules {
		if m.Name == module {
			found = true
			break
		}
	}
	if !found {
		return "", fmt.Errorf("unknown module %s", module)
	}
	if !slugName.MatchString(slug) {
		return "", fmt.Errorf("invalid slug %q", slug)
	}
	max := 0
	for _, folder := range []string{"changes", "changes/archive"} {
		entries, e := os.ReadDir(filepath.Join(snap.Root, folder))
		if e != nil && !os.IsNotExist(e) {
			return "", e
		}
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			parts := queuedName.FindStringSubmatch(entry.Name())
			if parts == nil {
				continue
			}
			if parts[2] == slug {
				return "", fmt.Errorf("duplicate slug %s", slug)
			}
			n, _ := strconv.Atoi(parts[1])
			if n > max {
				max = n
			}
		}
	}
	rel := fmt.Sprintf("changes/%04d-%s.md", max+1, slug)
	if err := os.MkdirAll(filepath.Join(snap.Root, "changes"), 0755); err != nil {
		return "", fmt.Errorf("create %s: %w", rel, err)
	}
	file, err := os.OpenFile(filepath.Join(snap.Root, filepath.FromSlash(rel)), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return "", fmt.Errorf("create %s: %w", rel, err)
	}
	content := "---\nmodule: " + module + "\n---\n\n## Motivation\n\n## Interface\n\n## Added\n\n## Modified\n\n## Removed\n\n## Decisions\n"
	_, err = file.WriteString(content)
	closeErr := file.Close()
	if err != nil {
		return "", fmt.Errorf("create %s: %w", rel, err)
	}
	if closeErr != nil {
		return "", fmt.Errorf("create %s: %w", rel, closeErr)
	}
	return rel, nil
}

 succeeded in 970ms:
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

 succeeded in 968ms:
# Project rebuild notes

## Contradicted Decisions

None identified in the retained project inputs.

## Gaps

- no spec (project): The native Go eval suite could not run in this input-only workspace because the build manifest and dependency implementations are deliberately absent. The separate native verifier must compile and run the retained evals.

 succeeded in 976ms:
package history

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

func section(body, name string) string { value, _ := contracts.Section(body, name); return value }

func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
	snap, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	var current model.Module
	found := false
	for _, m := range snap.Modules {
		if m.Name == module {
			current, found = m, true
			break
		}
	}
	if !found {
		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
	}
	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes", "archive"))
	if os.IsNotExist(err) {
		entries = nil
	} else if err != nil {
		return nil, err
	}
	result := []model.SpecHistoryEntry{}
	for _, file := range entries {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".md") {
			continue
		}
		rel := path.Join("changes/archive", file.Name())
		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
		if e != nil {
			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
		}
		doc, e := contracts.ParseDocument(string(data))
		if e != nil {
			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
		}
		archiveModule, ok := doc.Frontmatter["module"].(string)
		if !ok || archiveModule == "" {
			return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: invalid module target", rel))
		}
		var ids []string
		seen := map[string]bool{}
		for _, name := range []string{"Added", "Modified"} {
			items, e := contracts.ParseItems(section(doc.Body, name))
			if e != nil {
				return nil, fixture.Code("C03", rel, fmt.Errorf("C03: %s: %w", rel, e))
			}
			for _, item := range items {
				if !seen[item.ID] {
					seen[item.ID] = true
					ids = append(ids, item.ID)
				}
			}
		}
		if archiveModule != module || len(ids) == 0 {
			continue
		}
		if id != "" {
			if !seen[id] {
				continue
			}
			ids = []string{id}
		} else {
			sort.Strings(ids)
		}
		entry := model.SpecHistoryEntry{Path: rel, Module: archiveModule, IDs: ids, Motivation: section(doc.Body, "Motivation"), Decisions: section(doc.Body, "Decisions"), RemovedDecisions: section(doc.Body, "Removed decisions"), Provenance: section(doc.Body, "Provenance")}
		if entry.RemovedDecisions == "" {
			entry.RemovedDecisions = section(doc.Body, "Removed Decisions")
		}
		for _, line := range strings.Split(entry.Provenance, "\n") {
			fields := strings.Fields(strings.TrimPrefix(strings.TrimSpace(line), "-"))
			if len(fields) < 2 || fields[0] != "impl" {
				continue
			}
			date, e := time.Parse(time.RFC3339Nano, fields[1])
			if e == nil && date.After(entry.ImplementationDate) {
				entry.ImplementationDate = date
			}
		}
		result = append(result, entry)
	}
	if id != "" && len(result) == 0 {
		known := false
		body, _ := contracts.Section(current.Document.Body, "Specs")
		items, _ := contracts.ParseItems(body)
		for _, item := range items {
			if item.ID == id {
				known = true
				break
			}
		}
		if !known {
			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
		}
	}
	sort.Slice(result, func(i, j int) bool {
		a, b := result[i], result[j]
		if !a.ImplementationDate.Equal(b.ImplementationDate) {
			return a.ImplementationDate.After(b.ImplementationDate)
		}
		return a.Path < b.Path
	})
	return result, nil
}

 succeeded in 979ms:
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

 succeeded in 976ms:
package queue

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/lurkshark/fireweed/packages/contracts"
	"github.com/lurkshark/fireweed/packages/project/internal/fixture"
	"github.com/lurkshark/fireweed/packages/project/model"
)

var pendingName = regexp.MustCompile(`^([0-9]{4,})-([a-z0-9-]+)\.md$`)

func ParseChange(data []byte, rel string, number int, slug string) (model.Change, error) {
	c := model.Change{Number: number, Slug: slug, Path: rel}
	doc, err := contracts.ParseDocument(string(data))
	if err != nil {
		return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
	}
	c.Document = doc
	module, ok := doc.Frontmatter["module"].(string)
	if !ok || module == "" {
		return c, fixture.Code("C07", rel, fmt.Errorf("%s: invalid module target", rel))
	}
	c.Module = module
	seen := map[string]bool{}
	for _, field := range []string{"Added", "Modified", "Removed"} {
		section, _ := contracts.Section(doc.Body, field)
		items, err := contracts.ParseItems(section)
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return c, fixture.Code("C05", rel, fmt.Errorf("%s: %w", rel, err))
			}
			return c, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, err))
		}
		for _, item := range items {
			if seen[item.ID] {
				return c, fixture.Code("C05", rel, fmt.Errorf("%s: duplicate edit identity %s", rel, item.ID))
			}
			seen[item.ID] = true
		}
		switch field {
		case "Added":
			c.Added = items
		case "Modified":
			c.Modified = items
		case "Removed":
			c.Removed = items
		}
	}
	return c, nil
}

func Queue(root string) ([]model.Change, error) {
	snap, err := fixture.Load(root)
	if err != nil {
		return nil, err
	}
	modules := map[string]bool{}
	for _, m := range snap.Modules {
		modules[m.Name] = true
	}
	entries, err := os.ReadDir(filepath.Join(snap.Root, "changes"))
	if os.IsNotExist(err) {
		return []model.Change{}, nil
	}
	if err != nil {
		return nil, fixture.Code("C03", "changes", err)
	}
	var changes []model.Change
	slugs := map[string]bool{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		match := pendingName.FindStringSubmatch(entry.Name())
		if match == nil {
			continue
		}
		rel := path.Join("changes", entry.Name())
		if slugs[match[2]] {
			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
		}
		slugs[match[2]] = true
		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
		if e != nil {
			return nil, fixture.Code("C03", rel, fmt.Errorf("%s: %w", rel, e))
		}
		num, _ := strconv.Atoi(match[1])
		c, e := ParseChange(data, rel, num, match[2])
		if e != nil {
			return nil, e
		}
		if !modules[c.Module] {
			return nil, fixture.Code("C07", rel, fmt.Errorf("%s: unknown module %s", rel, c.Module))
		}
		changes = append(changes, c)
	}
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].Number != changes[j].Number {
			return changes[i].Number < changes[j].Number
		}
		return changes[i].Slug < changes[j].Slug
	})
	return changes, nil
}

func Resolve(changes []model.Change, selector string) (int, error) {
	idx := -1
	for i, c := range changes {
		match := selector == c.Slug || selector == path.Base(c.Path) || selector == c.Path
		if match {
			if idx >= 0 {
				return -1, fixture.Code("C13", selector, fmt.Errorf("ambiguous pending change %s", selector))
			}
			idx = i
		}
	}
	if idx < 0 {
		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
	}
	return idx, nil
}

func itemList(m model.Module) ([]contracts.Item, error) {
	section, _ := contracts.Section(m.Document.Body, "Specs")
	return contracts.ParseItems(section)
}

func itemByID(m model.Module, id string) (contracts.Item, bool) {
	items, err := itemList(m)
	if err != nil {
		return contracts.Item{}, false
	}
	for _, item := range items {
		if item.ID == id {
			return item, true
		}
	}
	return contracts.Item{}, false
}

func Before(root string, changes []model.Change, target int) (model.Snapshot, error) {
	snap, err := fixture.Load(root)
	if err != nil {
		return snap, err
	}
	if err := fixture.ValidateGraph(snap.Modules); err != nil {
		return snap, fixture.Code("C02", "", err)
	}
	for i := 0; i < target; i++ {
		if err := Apply(&snap, changes[i]); err != nil {
			return snap, err
		}
	}
	return snap, nil
}

func Project(root, through string) (model.Snapshot, error) {
	changes, err := Queue(root)
	if err != nil {
		return model.Snapshot{}, err
	}
	limit := len(changes)
	if through != "" {
		idx, e := Resolve(changes, through)
		if e != nil {
			return model.Snapshot{}, e
		}
		limit = idx + 1
	}
	snap, err := Before(root, changes, limit)
	if err != nil {
		return snap, err
	}
	return snap, nil
}

func Apply(snap *model.Snapshot, change model.Change) error {
	i := -1
	for n, m := range snap.Modules {
		if m.Name == change.Module {
			i = n
			break
		}
	}
	if i < 0 {
		return fixture.Code("C07", change.Path, fmt.Errorf("unknown module %s", change.Module))
	}
	m := snap.Modules[i]
	if value, ok := change.Document.Frontmatter["set"]; ok {
		settings, ok := value.(map[string]any)
		if !ok {
			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
		}
		keys := make([]string, 0, len(settings))
		for key := range settings {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			v := settings[key]
			switch key {
			case "depends_on", "interface", "evals", "disposable", "preserve", "append_only":
				if _, err := fixture.Strings(v); err != nil {
					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s: %w", key, err))
				}
			case "boundaries_opt_out":
				if _, ok := v.(string); !ok {
					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
				}
			case "hooks":
				if _, ok := v.(map[string]any); !ok {
					return fixture.Code("C12", change.Path, fmt.Errorf("invalid set %s", key))
				}
			default:
				return fixture.Code("C12", change.Path, fmt.Errorf("invalid set field %s", key))
			}
			m.Document.Frontmatter[key] = v
		}
		defaults, _ := snap.Config["module_defaults"].(map[string]any)
		updated, err := fixture.ValidateMetadata(m.Document.Frontmatter, defaults)
		if err != nil {
			return err
		}
		m.DependsOn, m.Zones = updated.DependsOn, updated.Zones
	}
	items, err := itemList(m)
	if err != nil {
		return fixture.Code("C03", path.Join(m.Dir, "SPEC.md"), err)
	}
	for _, edit := range change.Modified {
		found := false
		for j := range items {
			if items[j].ID == edit.ID {
				items[j] = edit
				found = true
				break
			}
		}
		if !found {
			return fixture.Code("C08", change.Path, fmt.Errorf("Modified ID %s absent from preceding projection", edit.ID))
		}
	}
	for _, edit := range change.Removed {
		found := false
		next := items[:0]
		for _, item := range items {
			if item.ID == edit.ID {
				found = true
			} else {
				next = append(next, item)
			}
		}
		if !found {
			return fixture.Code("C08", change.Path, fmt.Errorf("Removed ID %s absent from preceding projection", edit.ID))
		}
		items = next
	}
	items = append(items, change.Added...)
	var body strings.Builder
	for _, item := range items {
		body.WriteString(contracts.RenderItem(item))
		if !strings.HasSuffix(body.String(), "\n") {
			body.WriteByte('\n')
		}
		if !strings.HasSuffix(body.String(), "\n\n") {
			body.WriteByte('\n')
		}
	}
	m.Document.Body = contracts.ReplaceSection(m.Document.Body, "Specs", body.String())
	snap.Modules[i] = m
	if err := fixture.ValidateGraph(snap.Modules); err != nil {
		return fixture.Code("C02", change.Path, err)
	}
	return nil
}

func Target(root, selector string) (model.Snapshot, model.Change, error) {
	changes, err := Queue(root)
	if err != nil {
		return model.Snapshot{}, model.Change{}, err
	}
	idx, err := Resolve(changes, selector)
	if err != nil {
		return model.Snapshot{}, model.Change{}, err
	}
	snap, err := Before(root, changes, idx)
	return snap, changes[idx], err
}

func bases(doc contracts.Document) (map[string]any, error) {
	v, exists := doc.Frontmatter["bases"]
	if !exists {
		return map[string]any{}, nil
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, fixture.Code("C08", "bases", fmt.Errorf("bases is not a mapping"))
	}
	return m, nil
}

func stamp(root, selector string, missingOnly bool) error {
	snap, c, err := Target(root, selector)
	if err != nil {
		return err
	}
	b, err := bases(c.Document)
	if err != nil {
		return err
	}
	var m model.Module
	for _, module := range snap.Modules {
		if module.Name == c.Module {
			m = module
			break
		}
	}
	for _, edit := range append(append([]contracts.Item{}, c.Modified...), c.Removed...) {
		item, ok := itemByID(m, edit.ID)
		if !ok {
			return fixture.Code("C08", edit.ID, fmt.Errorf("ID %s absent from preceding projection", edit.ID))
		}
		if _, exists := b[edit.ID]; missingOnly && exists {
			continue
		}
		b[edit.ID] = contracts.SpecHash(item)
	}
	c.Document.Frontmatter["bases"] = b
	text, err := contracts.RenderDocument(c.Document)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(snap.Root, filepath.FromSlash(c.Path)), []byte(text), 0644)
}

func Stamp(root, selector string) error        { return stamp(root, selector, false) }
func StampMissing(root, selector string) error { return stamp(root, selector, true) }

exec
/bin/zsh -lc "cat 'packages/project/evals/stale_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/evals/static_codes_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/evals/coded_errors_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/app/evals/project_commands_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/app/evals/diagnostic_causes_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/contracts/model/model.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/contracts/api.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/project/evals/queue_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "cat 'packages/app/evals/list_cli_test.go'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 exited 1 in 1091ms:
cat: packages/app/evals/list_cli_test.go: No such file or directory

 succeeded in 1095ms:
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

 succeeded in 1096ms:
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

 succeeded in 1107ms:
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

 succeeded in 1112ms:
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

 exited 1 in 1112ms:
cat: packages/app/evals/diagnostic_causes_test.go: No such file or directory

 exited 1 in 1117ms:
cat: packages/app/evals/project_commands_test.go: No such file or directory

 succeeded in 1121ms:
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

 succeeded in 1137ms:
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

exec
/bin/zsh -lc 'rg --files | sort | head -120' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 452ms:
DRILL_NOTES.md
SYSTEM.md
packages/assets/api.go
packages/assets/model/model.go
packages/contracts/api.go
packages/contracts/model/model.go
packages/gitstore/api.go
packages/gitstore/model/model.go
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
packages/project/internal/queue/check.go
packages/project/internal/queue/queue.go
packages/project/internal/queue/stale.go
packages/project/model/model.go

exec
/bin/zsh -lc 'cat SYSTEM.md' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 450ms:
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

codex
The reported failures point to a few precise issues: diagnostic wording and codes, a handle reference that Check misses, and a numeric bases value that StaleSpecs loses. I’m patching those paths in `packages/project/internal` without changing the retained interfaces or evals.
exec
/bin/zsh -lc "rg -n 'NewChange|SpecHistory|C01:|U02:|stale|handle|foreign' packages/project/evals/*.go" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 468ms:
packages/project/evals/project_test.go:156:func Test_PRJ_003_PRJ_023_NewChangeSkeletonAndNumbering(t *testing.T) {
packages/project/evals/project_test.go:160:	rel, err := project.NewChange(root, "alpha", "new-feature")
packages/project/evals/project_test.go:162:		t.Fatalf("NewChange: %v", err)
packages/project/evals/project_test.go:193:func Test_PRJ_024_NewChangeRejectsDuplicatesWithoutOverwrite(t *testing.T) {
packages/project/evals/project_test.go:204:			if _, err := project.NewChange(root, "alpha", "repeat"); err == nil {
packages/project/evals/project_test.go:215:func Test_PRJ_025_PRJ_026_NewChangeInvalidArguments(t *testing.T) {
packages/project/evals/project_test.go:225:			_, err := project.NewChange(root, tc.module, tc.slug)
packages/project/evals/project_test.go:235:func Test_PRJ_067_NewChangeBlockedQueuedPathPreservesCause(t *testing.T) {
packages/project/evals/project_test.go:241:	path, err := project.NewChange(root, "alpha", "blocked")
packages/project/evals/project_test.go:243:		t.Fatal("NewChange accepted a queued path occupied by a directory")
packages/project/evals/project_test.go:246:		t.Errorf("rejected NewChange path = %q; want empty", path)
packages/project/evals/project_test.go:249:		t.Errorf("NewChange error = %q; want repository-relative create prefix", err)
packages/project/evals/project_test.go:253:		t.Errorf("NewChange error = %v; want retained filesystem cause for %s", err, queuedPath)
packages/project/evals/stale_test.go:15:func staleSpec(t *testing.T, root, through, id string) contracts.Item {
packages/project/evals/stale_test.go:24:func staleBaseItem(t *testing.T, root, id string) contracts.Item {
packages/project/evals/stale_test.go:33:func staleRead(t *testing.T, root, path string) []byte {
packages/project/evals/stale_test.go:42:func staleBases(t *testing.T, root, path string) (map[string]any, contracts.Document) {
packages/project/evals/stale_test.go:44:	doc, err := contracts.ParseDocument(string(staleRead(t, root, path)))
packages/project/evals/stale_test.go:55:func staleFixture(t *testing.T) string {
packages/project/evals/stale_test.go:63:	root := staleFixture(t)
packages/project/evals/stale_test.go:64:	oldOne := contracts.SpecHash(staleSpec(t, root, "", "ALP-001"))
packages/project/evals/stale_test.go:65:	oldTwo := contracts.SpecHash(staleSpec(t, root, "", "ALP-002"))
packages/project/evals/stale_test.go:66:	oldThree := contracts.SpecHash(staleSpec(t, root, "", "ALP-003"))
packages/project/evals/stale_test.go:74:		{"ALP-002", oldTwo, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-002"))},
packages/project/evals/stale_test.go:75:		{"ALP-001", oldOne, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-001"))},
packages/project/evals/stale_test.go:76:		{"ALP-003", oldThree, contracts.RenderItem(staleSpec(t, root, "earlier", "ALP-003"))},
packages/project/evals/stale_test.go:79:		t.Fatalf("stale entries = %#v; want %d recorded entries", got, len(want))
packages/project/evals/stale_test.go:89:	root := staleFixture(t)
packages/project/evals/stale_test.go:94:	before := staleSpec(t, root, "earlier", "ALP-001")
packages/project/evals/stale_test.go:102:	bases, doc := staleBases(t, root, path)
packages/project/evals/stale_test.go:116:	put(t, root, path, strings.Replace(string(staleRead(t, root, path)), "Target.", "Target with later prose.", 1))
packages/project/evals/stale_test.go:118:	currentOne := staleSpec(t, root, "earlier", "ALP-001")
packages/project/evals/stale_test.go:119:	currentTwo := staleSpec(t, root, "earlier", "ALP-002")
packages/project/evals/stale_test.go:125:		t.Fatalf("stale entries = %#v; want two", got)
packages/project/evals/stale_test.go:136:	root := staleFixture(t) // No Git repository or recorded target history.
packages/project/evals/stale_test.go:140:		t.Fatalf("unavailable history = %#v, %v; want both stale items", got, err)
packages/project/evals/stale_test.go:143:		if got[i].ID != want.id || got[i].RecordedHash != want.hash || got[i].Recovered || got[i].Before != "" || got[i].Current != contracts.RenderItem(staleBaseItem(t, root, want.id)) {
packages/project/evals/stale_test.go:150:	root := staleFixture(t)
packages/project/evals/stale_test.go:159:		target = staleRead(t, root, path)
packages/project/evals/stale_test.go:160:		index = staleRead(t, root, ".git/index")
packages/project/evals/stale_test.go:170:	put(t, root, "packages/alpha/SPEC.md", strings.Replace(string(staleRead(t, root, "packages/alpha/SPEC.md")), "One MUST work.", "One MUST now work differently.", 1))
packages/project/evals/stale_test.go:171:	spec := staleRead(t, root, "packages/alpha/SPEC.md")
packages/project/evals/stale_test.go:174:		t.Fatalf("stale base = %#v, %v; want one", got, err)
packages/project/evals/stale_test.go:177:	if !bytes.Equal(target, afterTarget) || !bytes.Equal(spec, staleRead(t, root, "packages/alpha/SPEC.md")) || !bytes.Equal(index, afterIndex) || head != afterHead || refs != afterRefs {
packages/project/evals/stale_test.go:183:	root := staleFixture(t)
packages/project/evals/stale_test.go:184:	queued(t, root, "0001-target.md", "custom: keep-me\nbases:\n  ALP-001: deadbeef\n  ALP-002: "+contracts.SpecHash(staleBaseItem(t, root, "ALP-002"))+"\n  ALP-999: extra\n", "Keep this motivation.\n", "", "### ALP-001\nOne MUST change.\n\n### ALP-002\nTwo MUST change.\n\n### ALP-003\nThree MUST change.\n", "### ALP-004\n")
packages/project/evals/stale_test.go:186:	_, original := staleBases(t, root, path)
packages/project/evals/stale_test.go:191:	bases, after := staleBases(t, root, path)
packages/project/evals/stale_test.go:193:		"ALP-001": "deadbeef", "ALP-002": contracts.SpecHash(staleBaseItem(t, root, "ALP-002")),
packages/project/evals/stale_test.go:194:		"ALP-003": contracts.SpecHash(staleBaseItem(t, root, "ALP-003")),
packages/project/evals/stale_test.go:195:		"ALP-004": contracts.SpecHash(staleBaseItem(t, root, "ALP-004")), "ALP-999": "extra",
packages/project/evals/stale_test.go:203:	staleFound := false
packages/project/evals/stale_test.go:205:		if diagnostic.Code == "C08" && strings.Contains(diagnostic.Message, "stale base for ALP-001") {
packages/project/evals/stale_test.go:206:			staleFound = true
packages/project/evals/stale_test.go:209:	if !staleFound {
packages/project/evals/stale_test.go:210:		t.Error("ordinary Check stopped reporting the preserved stale base")
packages/project/evals/stale_test.go:215:	stamped, _ := staleBases(t, root, path)
packages/project/evals/stale_test.go:216:	if stamped["ALP-001"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-001")) {
packages/project/evals/stale_test.go:217:		t.Errorf("ordinary Stamp stopped replacing a stale base: %#v", stamped)
packages/project/evals/stale_test.go:228:			root := staleFixture(t)
packages/project/evals/stale_test.go:231:			before := staleRead(t, root, path)
packages/project/evals/stale_test.go:236:			if !bytes.Equal(before, staleRead(t, root, path)) {
packages/project/evals/stale_test.go:246:			root := staleFixture(t)
packages/project/evals/stale_test.go:255:			bases, _ := staleBases(t, root, "changes/0001-target.md")
packages/project/evals/stale_test.go:256:			if bases["ALP-002"] != contracts.SpecHash(staleBaseItem(t, root, "ALP-002")) || bases["ALP-001"] != "deadbeef" {
packages/project/evals/stale_test.go:261:	root := staleFixture(t)
packages/project/evals/spec_history_test.go:53:func historyPaths(entries []project.SpecHistoryEntry) []string {
packages/project/evals/spec_history_test.go:61:func Test_PRJ_051_SpecHistorySelectsArchivedAddsAndTagOnlyModifications(t *testing.T) {
packages/project/evals/spec_history_test.go:63:	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
packages/project/evals/spec_history_test.go:65:		t.Fatalf("SpecHistory: %v", err)
packages/project/evals/spec_history_test.go:89:func Test_PRJ_053_SpecHistoryOrdersByLatestValidImplementationDate(t *testing.T) {
packages/project/evals/spec_history_test.go:91:	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
packages/project/evals/spec_history_test.go:93:		t.Fatalf("SpecHistory: %v", err)
packages/project/evals/spec_history_test.go:117:func Test_PRJ_052_SpecHistoryPreservesCompleteEntryMetadata(t *testing.T) {
packages/project/evals/spec_history_test.go:119:	entries, err := project.SpecHistory(root, "alpha", "ALP-002")
packages/project/evals/spec_history_test.go:121:		t.Fatalf("SpecHistory: %v", err)
packages/project/evals/spec_history_test.go:126:	want := project.SpecHistoryEntry{
packages/project/evals/spec_history_test.go:139:func Test_PRJ_054_SpecHistoryModuleWideUniqueSortedIDsAndEmptyModule(t *testing.T) {
packages/project/evals/spec_history_test.go:141:	entries, err := project.SpecHistory(root, "alpha", "")
packages/project/evals/spec_history_test.go:143:		t.Fatalf("SpecHistory module-wide: %v", err)
packages/project/evals/spec_history_test.go:167:	empty, err := project.SpecHistory(emptyRoot, "beta", "")
packages/project/evals/spec_history_test.go:173:func Test_PRJ_055_SpecHistoryLookupNamesAndRemovedCurrentSpec(t *testing.T) {
packages/project/evals/spec_history_test.go:175:	archived, err := project.SpecHistory(root, "alpha", "ALP-002")
packages/project/evals/spec_history_test.go:179:	currentOnly, err := project.SpecHistory(root, "alpha", "ALP-001")
packages/project/evals/spec_history_test.go:188:		_, err := project.SpecHistory(root, tc.module, tc.id)
packages/project/evals/spec_history_test.go:189:		if err == nil || !strings.HasPrefix(err.Error(), "U02: ") || !strings.Contains(err.Error(), tc.rejected) {
packages/project/evals/spec_history_test.go:190:			t.Errorf("lookup module %q ID %q = %v; want U02: naming %q", tc.module, tc.id, err, tc.rejected)
packages/project/evals/spec_history_test.go:195:func Test_PRJ_056_SpecHistoryConfigurationAndArchiveSyntaxErrors(t *testing.T) {
packages/project/evals/spec_history_test.go:209:			_, err := project.SpecHistory(root, "alpha", "")
packages/project/evals/spec_history_test.go:210:			if err == nil || !strings.HasPrefix(err.Error(), "C01: ") {
packages/project/evals/spec_history_test.go:211:				t.Errorf("invalid config = %v; want C01: prefix", err)
packages/project/evals/spec_history_test.go:224:			_, err := project.SpecHistory(root, "alpha", "ALP-002")
packages/project/evals/queue_test.go:168:		t.Errorf("whole-queue projection lost local handles or retained removed ID: %q", section)
packages/project/evals/queue_test.go:303:		{"PRJ-012", "foreign_module_id", "### BET-001\nForeign item MUST fail.\n", "", "Needed.\n", "", "", "", "C05", "foreign spec ID"},
packages/project/evals/queue_test.go:304:		{"PRJ-012", "foreign_change_id", "", "", "Needed.\n", "", "### BET-001\nForeign edit MUST fail.\n", "", "C05", "foreign module ID"},
packages/project/evals/queue_test.go:306:		{"PRJ-014", "undeclared_handle", "", "", "Needed.\n", "### new:local\nLocal MUST work.\n", "### ALP-001\nAlpha MUST refer to new:foreign.\n", "", "C09", "reference to undeclared handle new:foreign"},
packages/project/evals/queue_test.go:308:		{"PRJ-015", "stale_base", "", "bases:\n  ALP-001: deadbeef\n", "Needed.\n", "", "### ALP-001\nAlpha MUST change.\n", "", "C08", "stale base for ALP-001"},
packages/project/evals/queue_test.go:316:			if tc.name == "duplicate_edit_identity" || tc.name == "undeclared_handle" {
packages/project/evals/queue_test.go:319:			if tc.name == "foreign_change_id" {
packages/project/evals/queue_test.go:329:			if tc.name == "undeclared_handle" {
packages/project/evals/queue_test.go:330:				queued(t, root, "0000-other.md", "", "Earlier local handle.\n", "### new:foreign\nEarlier item MUST work.\n", "", "")
packages/project/evals/queue_test.go:367:			if tc.name == "missing_base" || tc.name == "stale_base" {

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/fixture/fixture.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/check.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/stale.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index ea0d2d0cffdc2e2e750bc597be0f3aa98d4288fc..89201112649ffd471580238aa383794adafef62b
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -31,10 +31,10 @@
 func Config(data []byte) (map[string]any, error) {
 	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
 	if err != nil {
-		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("fireweed.yaml: %w", err))
 	}
 	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
-		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("invalid modules: %w", err))
 	}
 	return doc.Frontmatter, nil
 }
@@ -156,7 +156,7 @@
 		base, err = canonical(root)
 	}
 	if err != nil {
-		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("discover project root: %w", err))
 	}
 	snap.Root = base
 	var configBytes []byte
@@ -166,7 +166,7 @@
 		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
 	}
 	if err != nil {
-		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("read fireweed.yaml: %w", err))
 	}
 	snap.Config, err = Config(configBytes)
 	if err != nil {
@@ -289,7 +289,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if _, ok := byName[dep]; !ok {
-				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, dep)
 			}
 		}
 	}
@@ -338,7 +338,7 @@
 		return "", fmt.Errorf("unknown module %s", module)
 	}
 	if !slugName.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Code("U01", slug, fmt.Errorf("invalid slug %q", slug))
 	}
 	max := 0
 	for _, folder := range []string{"changes", "changes/archive"} {
@@ -355,7 +355,7 @@
 				continue
 			}
 			if parts[2] == slug {
-				return "", fmt.Errorf("duplicate slug %s", slug)
+				return "", Code("C13", slug, fmt.Errorf("duplicate slug %s", slug))
 			}
 			n, _ := strconv.Atoi(parts[1])
 			if n > max {
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
index 73d3b04b5426054ab5e9706b22e79e120e3549e6..8d5c0f7ec7efd793d277338d2b8ce50c8ea6883f
--- a/packages/project/internal/queue/check.go
+++ b/packages/project/internal/queue/check.go
@@ -51,7 +51,11 @@
 		if !strings.HasPrefix(item.ID, "new:") {
 			prefix := strings.SplitN(item.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
-				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+				kind := "foreign module ID"
+				if strings.HasSuffix(file, "/SPEC.md") {
+					kind = "foreign spec ID"
+				}
+				add("C05", fmt.Sprintf("%s %s belongs to %s", kind, item.ID, owner), file)
 			}
 		}
 	}
@@ -107,11 +111,9 @@
 		for _, item := range c.Removed {
 			validateItem(item, m, c.Path, false)
 		}
-		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
-			for _, handle := range handles.FindAllString(item.Body, -1) {
-				if !declared[handle] {
-					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
-				}
+		for _, handle := range handles.FindAllString(c.Document.Body, -1) {
+			if !declared[handle] {
+				add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
 			}
 		}
 		b, e := bases(c.Document)
@@ -125,7 +127,7 @@
 					continue
 				}
 				recorded, present := b[item.ID]
-				hash, _ := recorded.(string)
+				hash := baseText(recorded)
 				if !present || hash == "" {
 					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
 				} else if hash != contracts.SpecHash(prior) {
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index cf40c2f0b0eeb244322844609e1ddebc2705d802..7fddf6bd8df42441e0f8aa7f7a4f27011ad39a1d
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -85,7 +85,7 @@
 		}
 		rel := path.Join("changes", entry.Name())
 		if slugs[match[2]] {
-			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("ambiguous change slug %s", match[2]))
 		}
 		slugs[match[2]] = true
 		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
@@ -123,7 +123,7 @@
 		}
 	}
 	if idx < 0 {
-		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+		return -1, fixture.Code("U02", selector, fmt.Errorf("pending change %s not found", selector))
 	}
 	return idx, nil
 }
@@ -197,7 +197,7 @@
 	if value, ok := change.Document.Frontmatter["set"]; ok {
 		settings, ok := value.(map[string]any)
 		if !ok {
-			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+			return fixture.Code("C12", change.Path, fmt.Errorf("set must be a mapping"))
 		}
 		keys := make([]string, 0, len(settings))
 		for key := range settings {
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
index c4aeff2e817d277f26d0cead3fb58677d2aead4a..29e492b6bf70616a72a21b0b3697425dde4d792e
--- a/packages/project/internal/queue/stale.go
+++ b/packages/project/internal/queue/stale.go
@@ -125,7 +125,7 @@
 		if !exists {
 			continue
 		}
-		hash, _ := recorded.(string)
+		hash := baseText(recorded)
 		prior, found := itemByID(module, edit.ID)
 		if !found {
 			continue

apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/history/history.go
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/packages/project/internal/queue/queue.go
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index ea0d2d0cffdc2e2e750bc597be0f3aa98d4288fc..89201112649ffd471580238aa383794adafef62b
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -31,10 +31,10 @@
 func Config(data []byte) (map[string]any, error) {
 	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
 	if err != nil {
-		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("fireweed.yaml: %w", err))
 	}
 	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
-		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("invalid modules: %w", err))
 	}
 	return doc.Frontmatter, nil
 }
@@ -156,7 +156,7 @@
 		base, err = canonical(root)
 	}
 	if err != nil {
-		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("discover project root: %w", err))
 	}
 	snap.Root = base
 	var configBytes []byte
@@ -166,7 +166,7 @@
 		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
 	}
 	if err != nil {
-		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("read fireweed.yaml: %w", err))
 	}
 	snap.Config, err = Config(configBytes)
 	if err != nil {
@@ -289,7 +289,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if _, ok := byName[dep]; !ok {
-				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, dep)
 			}
 		}
 	}
@@ -338,7 +338,7 @@
 		return "", fmt.Errorf("unknown module %s", module)
 	}
 	if !slugName.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Code("U01", slug, fmt.Errorf("invalid slug %q", slug))
 	}
 	max := 0
 	for _, folder := range []string{"changes", "changes/archive"} {
@@ -355,7 +355,7 @@
 				continue
 			}
 			if parts[2] == slug {
-				return "", fmt.Errorf("duplicate slug %s", slug)
+				return "", Code("C13", slug, fmt.Errorf("duplicate slug %s", slug))
 			}
 			n, _ := strconv.Atoi(parts[1])
 			if n > max {
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
index e6e683bc26bda6beb082ca3f247962c648c46db4..2154265f1f6b2c03b563d3bb0b9a3ee131f1bf93
--- a/packages/project/internal/history/history.go
+++ b/packages/project/internal/history/history.go
@@ -19,6 +19,9 @@
 func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
 	snap, err := fixture.Load(root)
 	if err != nil {
+		if code, _, ok := contracts.ErrorCode(err); ok && code == "C01" {
+			return nil, fixture.Code("C01", "fireweed.yaml", fmt.Errorf("C01: %w", err))
+		}
 		return nil, err
 	}
 	var current model.Module
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
index 73d3b04b5426054ab5e9706b22e79e120e3549e6..8d5c0f7ec7efd793d277338d2b8ce50c8ea6883f
--- a/packages/project/internal/queue/check.go
+++ b/packages/project/internal/queue/check.go
@@ -51,7 +51,11 @@
 		if !strings.HasPrefix(item.ID, "new:") {
 			prefix := strings.SplitN(item.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
-				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+				kind := "foreign module ID"
+				if strings.HasSuffix(file, "/SPEC.md") {
+					kind = "foreign spec ID"
+				}
+				add("C05", fmt.Sprintf("%s %s belongs to %s", kind, item.ID, owner), file)
 			}
 		}
 	}
@@ -107,11 +111,9 @@
 		for _, item := range c.Removed {
 			validateItem(item, m, c.Path, false)
 		}
-		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
-			for _, handle := range handles.FindAllString(item.Body, -1) {
-				if !declared[handle] {
-					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
-				}
+		for _, handle := range handles.FindAllString(c.Document.Body, -1) {
+			if !declared[handle] {
+				add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
 			}
 		}
 		b, e := bases(c.Document)
@@ -125,7 +127,7 @@
 					continue
 				}
 				recorded, present := b[item.ID]
-				hash, _ := recorded.(string)
+				hash := baseText(recorded)
 				if !present || hash == "" {
 					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
 				} else if hash != contracts.SpecHash(prior) {
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index cf40c2f0b0eeb244322844609e1ddebc2705d802..a6236c7d0e822c55bdb378d5866a9708262c7c1b
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -85,7 +85,7 @@
 		}
 		rel := path.Join("changes", entry.Name())
 		if slugs[match[2]] {
-			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("ambiguous change slug %s", match[2]))
 		}
 		slugs[match[2]] = true
 		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
@@ -123,7 +123,7 @@
 		}
 	}
 	if idx < 0 {
-		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+		return -1, fixture.Code("U02", selector, fmt.Errorf("pending change %s not found", selector))
 	}
 	return idx, nil
 }
@@ -197,7 +197,7 @@
 	if value, ok := change.Document.Frontmatter["set"]; ok {
 		settings, ok := value.(map[string]any)
 		if !ok {
-			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+			return fixture.Code("C12", change.Path, fmt.Errorf("set must be a mapping"))
 		}
 		keys := make([]string, 0, len(settings))
 		for key := range settings {
@@ -307,6 +307,16 @@
 	return m, nil
 }
 
+// YAML may decode an unquoted decimal hash as a number. Its textual value is
+// still the recorded base, so compare and report it rather than treating it as
+// an absent hash.
+func baseText(value any) string {
+	if value == nil {
+		return ""
+	}
+	return fmt.Sprint(value)
+}
+
 func stamp(root, selector string, missingOnly bool) error {
 	snap, c, err := Target(root, selector)
 	if err != nil {
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
index c4aeff2e817d277f26d0cead3fb58677d2aead4a..29e492b6bf70616a72a21b0b3697425dde4d792e
--- a/packages/project/internal/queue/stale.go
+++ b/packages/project/internal/queue/stale.go
@@ -125,7 +125,7 @@
 		if !exists {
 			continue
 		}
-		hash, _ := recorded.(string)
+		hash := baseText(recorded)
 		prior, found := itemByID(module, edit.ID)
 		if !found {
 			continue

exec
/bin/zsh -lc 'gofmt -w packages/project/internal/fixture/fixture.go packages/project/internal/history/history.go packages/project/internal/queue/queue.go packages/project/internal/queue/check.go packages/project/internal/queue/stale.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 succeeded in 449ms:
codex
The focused fixes are in place. The reported app failures also depend on project error codes and cause text, so I adjusted those at the project boundary. I’m checking the edited files now; this workspace does not include the module manifest or dependency implementations needed to run the Go eval suite here.
apply patch
patch: completed
/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/DRILL_NOTES.md
diff --git a/DRILL_NOTES.md b/DRILL_NOTES.md
index ec32bf8825cb42a234d4a05ad600041fd4bfa159..b0c49b065dec7cbe7fc713def9da1e3a1296ae50
--- a/DRILL_NOTES.md
+++ b/DRILL_NOTES.md
@@ -6,4 +6,4 @@
 
 ## Gaps
 
-- no spec (project): The native Go eval suite could not run in this input-only workspace because the build manifest and dependency implementations are deliberately absent. The separate native verifier must compile and run the retained evals.
+- no spec (project): The supplied input omits the Go module manifest and dependent implementations, so the project and transitive app evals cannot run locally. The separate native verifier must compile and run them.
diff --git a/packages/project/internal/fixture/fixture.go b/packages/project/internal/fixture/fixture.go
index ea0d2d0cffdc2e2e750bc597be0f3aa98d4288fc..89201112649ffd471580238aa383794adafef62b
--- a/packages/project/internal/fixture/fixture.go
+++ b/packages/project/internal/fixture/fixture.go
@@ -31,10 +31,10 @@
 func Config(data []byte) (map[string]any, error) {
 	doc, err := contracts.ParseDocument("---\n" + string(data) + "\n---\n")
 	if err != nil {
-		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: fireweed.yaml: %w", err))
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("fireweed.yaml: %w", err))
 	}
 	if _, err = Strings(doc.Frontmatter["modules"]); err != nil {
-		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("C01: invalid modules: %w", err))
+		return nil, Code("C01", "fireweed.yaml", fmt.Errorf("invalid modules: %w", err))
 	}
 	return doc.Frontmatter, nil
 }
@@ -156,7 +156,7 @@
 		base, err = canonical(root)
 	}
 	if err != nil {
-		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: discover project root: %w", err))
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("discover project root: %w", err))
 	}
 	snap.Root = base
 	var configBytes []byte
@@ -166,7 +166,7 @@
 		configBytes, err = gitstore.ReadAt(base, revision, "fireweed.yaml")
 	}
 	if err != nil {
-		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("C01: read fireweed.yaml: %w", err))
+		return snap, Code("C01", "fireweed.yaml", fmt.Errorf("read fireweed.yaml: %w", err))
 	}
 	snap.Config, err = Config(configBytes)
 	if err != nil {
@@ -289,7 +289,7 @@
 				return fmt.Errorf("self dependency %s", m.Name)
 			}
 			if _, ok := byName[dep]; !ok {
-				return fmt.Errorf("module %s dependency %s is missing", m.Name, dep)
+				return fmt.Errorf("module %s depends on missing module %s", m.Name, dep)
 			}
 		}
 	}
@@ -338,7 +338,7 @@
 		return "", fmt.Errorf("unknown module %s", module)
 	}
 	if !slugName.MatchString(slug) {
-		return "", fmt.Errorf("invalid slug %q", slug)
+		return "", Code("U01", slug, fmt.Errorf("invalid slug %q", slug))
 	}
 	max := 0
 	for _, folder := range []string{"changes", "changes/archive"} {
@@ -355,7 +355,7 @@
 				continue
 			}
 			if parts[2] == slug {
-				return "", fmt.Errorf("duplicate slug %s", slug)
+				return "", Code("C13", slug, fmt.Errorf("duplicate slug %s", slug))
 			}
 			n, _ := strconv.Atoi(parts[1])
 			if n > max {
diff --git a/packages/project/internal/history/history.go b/packages/project/internal/history/history.go
index e6e683bc26bda6beb082ca3f247962c648c46db4..2154265f1f6b2c03b563d3bb0b9a3ee131f1bf93
--- a/packages/project/internal/history/history.go
+++ b/packages/project/internal/history/history.go
@@ -19,6 +19,9 @@
 func SpecHistory(root, module, id string) ([]model.SpecHistoryEntry, error) {
 	snap, err := fixture.Load(root)
 	if err != nil {
+		if code, _, ok := contracts.ErrorCode(err); ok && code == "C01" {
+			return nil, fixture.Code("C01", "fireweed.yaml", fmt.Errorf("C01: %w", err))
+		}
 		return nil, err
 	}
 	var current model.Module
diff --git a/packages/project/internal/queue/check.go b/packages/project/internal/queue/check.go
index 73d3b04b5426054ab5e9706b22e79e120e3549e6..8d5c0f7ec7efd793d277338d2b8ce50c8ea6883f
--- a/packages/project/internal/queue/check.go
+++ b/packages/project/internal/queue/check.go
@@ -51,7 +51,11 @@
 		if !strings.HasPrefix(item.ID, "new:") {
 			prefix := strings.SplitN(item.ID, "-", 2)[0]
 			if owner, ok := prefixes[prefix]; ok && owner != module.Name {
-				add("C05", fmt.Sprintf("foreign module ID %s belongs to %s", item.ID, owner), file)
+				kind := "foreign module ID"
+				if strings.HasSuffix(file, "/SPEC.md") {
+					kind = "foreign spec ID"
+				}
+				add("C05", fmt.Sprintf("%s %s belongs to %s", kind, item.ID, owner), file)
 			}
 		}
 	}
@@ -107,11 +111,9 @@
 		for _, item := range c.Removed {
 			validateItem(item, m, c.Path, false)
 		}
-		for _, item := range append(append(append([]contracts.Item{}, c.Added...), c.Modified...), c.Removed...) {
-			for _, handle := range handles.FindAllString(item.Body, -1) {
-				if !declared[handle] {
-					add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
-				}
+		for _, handle := range handles.FindAllString(c.Document.Body, -1) {
+			if !declared[handle] {
+				add("C09", fmt.Sprintf("reference to undeclared handle %s", handle), c.Path)
 			}
 		}
 		b, e := bases(c.Document)
@@ -125,7 +127,7 @@
 					continue
 				}
 				recorded, present := b[item.ID]
-				hash, _ := recorded.(string)
+				hash := baseText(recorded)
 				if !present || hash == "" {
 					add("C08", fmt.Sprintf("missing base for %s", item.ID), c.Path)
 				} else if hash != contracts.SpecHash(prior) {
diff --git a/packages/project/internal/queue/queue.go b/packages/project/internal/queue/queue.go
index cf40c2f0b0eeb244322844609e1ddebc2705d802..a6236c7d0e822c55bdb378d5866a9708262c7c1b
--- a/packages/project/internal/queue/queue.go
+++ b/packages/project/internal/queue/queue.go
@@ -85,7 +85,7 @@
 		}
 		rel := path.Join("changes", entry.Name())
 		if slugs[match[2]] {
-			return nil, fixture.Code("C13", match[2], fmt.Errorf("duplicate pending slug %s", match[2]))
+			return nil, fixture.Code("C13", match[2], fmt.Errorf("ambiguous change slug %s", match[2]))
 		}
 		slugs[match[2]] = true
 		data, e := os.ReadFile(filepath.Join(snap.Root, filepath.FromSlash(rel)))
@@ -123,7 +123,7 @@
 		}
 	}
 	if idx < 0 {
-		return -1, fixture.Code("U02", selector, fmt.Errorf("U02: pending change %s not found", selector))
+		return -1, fixture.Code("U02", selector, fmt.Errorf("pending change %s not found", selector))
 	}
 	return idx, nil
 }
@@ -197,7 +197,7 @@
 	if value, ok := change.Document.Frontmatter["set"]; ok {
 		settings, ok := value.(map[string]any)
 		if !ok {
-			return fixture.Code("C12", change.Path, fmt.Errorf("invalid set mapping"))
+			return fixture.Code("C12", change.Path, fmt.Errorf("set must be a mapping"))
 		}
 		keys := make([]string, 0, len(settings))
 		for key := range settings {
@@ -307,6 +307,16 @@
 	return m, nil
 }
 
+// YAML may decode an unquoted decimal hash as a number. Its textual value is
+// still the recorded base, so compare and report it rather than treating it as
+// an absent hash.
+func baseText(value any) string {
+	if value == nil {
+		return ""
+	}
+	return fmt.Sprint(value)
+}
+
 func stamp(root, selector string, missingOnly bool) error {
 	snap, c, err := Target(root, selector)
 	if err != nil {
diff --git a/packages/project/internal/queue/stale.go b/packages/project/internal/queue/stale.go
index c4aeff2e817d277f26d0cead3fb58677d2aead4a..29e492b6bf70616a72a21b0b3697425dde4d792e
--- a/packages/project/internal/queue/stale.go
+++ b/packages/project/internal/queue/stale.go
@@ -125,7 +125,7 @@
 		if !exists {
 			continue
 		}
-		hash, _ := recorded.(string)
+		hash := baseText(recorded)
 		prior, found := itemByID(module, edit.ID)
 		if !found {
 			continue

exec
/bin/zsh -lc "rg --files -g 'go.mod' -g 'go.work' -g 'packages/contracts/internal/**' -g 'packages/gitstore/internal/**'" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc "rg -n 'C01:|U02:|foreign spec ID|foreign module ID|ambiguous change slug|set must be a mapping|baseText|FindAllString\\(c.Document.Body|invalid slug|duplicate slug|missing module' packages/project/internal DRILL_NOTES.md" in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
exec
/bin/zsh -lc 'gofmt -l packages/project/internal/fixture/fixture.go packages/project/internal/history/history.go packages/project/internal/queue/queue.go packages/project/internal/queue/check.go packages/project/internal/queue/stale.go' in /private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501
 exited 1 in 560ms:
 succeeded in 560ms:
 succeeded in 562ms:
packages/project/internal/queue/queue.go:88:			return nil, fixture.Code("C13", match[2], fmt.Errorf("ambiguous change slug %s", match[2]))
packages/project/internal/queue/queue.go:200:			return fixture.Code("C12", change.Path, fmt.Errorf("set must be a mapping"))
packages/project/internal/queue/queue.go:313:func baseText(value any) string {
packages/project/internal/queue/stale.go:128:		hash := baseText(recorded)
packages/project/internal/history/history.go:23:			return nil, fixture.Code("C01", "fireweed.yaml", fmt.Errorf("C01: %w", err))
packages/project/internal/history/history.go:36:		return nil, fixture.Code("U02", module, fmt.Errorf("U02: unknown module %s", module))
packages/project/internal/history/history.go:114:			return nil, fixture.Code("U02", id, fmt.Errorf("U02: unknown spec %s", id))
packages/project/internal/queue/check.go:54:				kind := "foreign module ID"
packages/project/internal/queue/check.go:56:					kind = "foreign spec ID"
packages/project/internal/queue/check.go:114:		for _, handle := range handles.FindAllString(c.Document.Body, -1) {
packages/project/internal/queue/check.go:130:				hash := baseText(recorded)
packages/project/internal/fixture/fixture.go:292:				return fmt.Errorf("module %s depends on missing module %s", m.Name, dep)
packages/project/internal/fixture/fixture.go:341:		return "", Code("U01", slug, fmt.Errorf("invalid slug %q", slug))
packages/project/internal/fixture/fixture.go:358:				return "", Code("C13", slug, fmt.Errorf("duplicate slug %s", slug))

codex
Updated the disposable project implementation to address the reported failures: foreign spec ID wording, stale numeric base hashes, handle references in change prose, and the error codes and messages used by the app. Updated [DRILL_NOTES.md](/private/var/folders/yh/crwgnf9d41911rdbn5g8q3v00000gn/T/fireweed-drill-input-2744460501/DRILL_NOTES.md).

`gofmt` passed. I could not run the Go evals here because the supplied workspace has no Go module manifest or dependent implementations; the native verifier still needs to confirm the fixes.
tokens used
46,026

