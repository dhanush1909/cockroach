// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package tests

var asyncpgBlocklists = blocklistsForVersion{
	{"v21.2", "asyncpgBlocklist21_2", asyncpgBlocklist21_2, "asyncpgIgnoreList21_2", asyncpgIgnoreList21_2},
	{"v22.1", "asyncpgBlocklist22_1", asyncpgBlocklist22_1, "asyncpgIgnoreList22_1", asyncpgIgnoreList22_1},
}

// These are lists of known asyncpg test errors and failures.
// When the asyncpg test suite is run, the results are compared to this list.
// Any failed test that is on this list is reported as FAIL - expected.
// Any failed test that is not on this list is reported as FAIL - unexpected.
//
// Please keep these lists alphabetized for easy diffing.
// After a failed run, an updated version of this blocklist should be available
// in the test log.
var asyncpgBlocklist22_1 = asyncpgBlocklist21_2

var asyncpgBlocklist21_2 = blocklist{}

var asyncpgIgnoreList22_1 = asyncpgIgnoreList21_2

var asyncpgIgnoreList21_2 = blocklist{
	"test_cache_invalidation.TestCacheInvalidation.test_prepare_cache_invalidation_in_pool":               "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_prepare_cache_invalidation_in_transaction":        "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_prepare_cache_invalidation_silent":                "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_prepared_type_cache_invalidation":                 "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_type_cache_invalidation_in_cancelled_transaction": "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_type_cache_invalidation_in_pool":                  "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_type_cache_invalidation_in_transaction":           "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_type_cache_invalidation_on_change_attr":           "unknown",
	"test_cache_invalidation.TestCacheInvalidation.test_type_cache_invalidation_on_drop_type_attr":        "unknown",
	"test_cancellation.TestCancellation.test_cancellation_03":                                             "unknown",
	"test_codecs.TestCodecs.test_array_with_custom_json_text_codec":                                       "unknown",
	"test_codecs.TestCodecs.test_composites_in_arrays":                                                    "unknown",
	"test_codecs.TestCodecs.test_enum":                                                                    "unknown",
	"test_codecs.TestCodecs.test_enum_and_range":                                                          "unknown",
	"test_codecs.TestCodecs.test_enum_function_return":                                                    "unknown",
	"test_codecs.TestCodecs.test_enum_in_array":                                                           "unknown",
	"test_codecs.TestCodecs.test_enum_in_composite":                                                       "unknown",
	"test_codecs.TestCodecs.test_invalid_input":                                                           "unknown",
	"test_codecs.TestCodecs.test_relacl_array_type":                                                       "unknown",
	"test_codecs.TestCodecs.test_timetz_encoding":                                                         "unknown",
	"test_codecs.TestCodecs.test_unhandled_type_fallback":                                                 "unknown",
	"test_codecs.TestCodecs.test_unknown_type_text_fallback":                                              "unknown",
	"test_codecs.TestCodecs.test_void":                                                                    "unknown",
	"test_connect.TestSettings.test_get_settings_01":                                                      "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_basics":                                                  "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_cancellation_explicit":                                   "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_cancellation_on_sink_error":                              "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_cancellation_while_waiting_for_data":                     "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_timeout_1":                                               "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_timeout_2":                                               "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_to_path":                                                 "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_to_path_like":                                            "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_to_sink":                                                 "unknown",
	"test_copy.TestCopyFrom.test_copy_from_query_with_args":                                               "unknown",
	"test_copy.TestCopyFrom.test_copy_from_table_basics":                                                  "unknown",
	"test_copy.TestCopyFrom.test_copy_from_table_large_rows":                                              "unknown",
	"test_copy.TestCopyTo.test_copy_records_to_table_1":                                                   "unknown",
	"test_copy.TestCopyTo.test_copy_records_to_table_async":                                               "unknown",
	"test_copy.TestCopyTo.test_copy_to_table_basics":                                                      "unknown",
	"test_cursor.TestCursor.test_cursor_02":                                                               "unknown",
	"test_cursor.TestCursor.test_cursor_04":                                                               "unknown",
	"test_cursor.TestIterableCursor.test_cursor_iterable_02":                                              "unknown",
	"test_exceptions.TestExceptions.test_exceptions_str":                                                  "unknown",
	"test_exceptions.TestExceptions.test_exceptions_unpacking":                                            "unknown",
	"test_execute.TestExecuteMany.test_executemany_client_failure_after_writes":                           "unknown",
	"test_execute.TestExecuteMany.test_executemany_client_failure_in_transaction":                         "unknown",
	"test_execute.TestExecuteMany.test_executemany_client_server_failure_conflict":                        "unknown",
	"test_execute.TestExecuteMany.test_executemany_prepare":                                               "unknown",
	"test_execute.TestExecuteMany.test_executemany_server_failure":                                        "unknown",
	"test_execute.TestExecuteMany.test_executemany_server_failure_after_writes":                           "unknown",
	"test_execute.TestExecuteMany.test_executemany_server_failure_during_writes":                          "unknown",
	"test_execute.TestExecuteMany.test_executemany_timeout":                                               "unknown",
	"test_execute.TestExecuteMany.test_executemany_timeout_flow_control":                                  "unknown",
	"test_execute.TestExecuteScript.test_execute_script_3":                                                "unknown",
	"test_execute.TestExecuteScript.test_execute_script_check_transactionality":                           "unknown",
	"test_execute.TestExecuteScript.test_execute_script_interrupted_close":                                "unknown",
	"test_introspection.TestIntrospection.test_introspection_no_stmt_cache_01":                            "unknown",
	"test_introspection.TestIntrospection.test_introspection_no_stmt_cache_02":                            "unknown",
	"test_introspection.TestIntrospection.test_introspection_no_stmt_cache_03":                            "unknown",
	"test_introspection.TestIntrospection.test_introspection_on_large_db":                                 "unknown",
	"test_introspection.TestIntrospection.test_introspection_retries_after_cache_bust":                    "unknown",
	"test_introspection.TestIntrospection.test_introspection_sticks_for_ps":                               "unknown",
	"test_listeners.TestListeners.test_listen_01":                                                         "unknown",
	"test_listeners.TestListeners.test_listen_02":                                                         "unknown",
	"test_listeners.TestListeners.test_listen_notletters":                                                 "unknown",
	"test_listeners.TestLogListeners.test_log_listener_01":                                                "unknown",
	"test_listeners.TestLogListeners.test_log_listener_02":                                                "unknown",
	"test_listeners.TestLogListeners.test_log_listener_03":                                                "unknown",
	"test_pool.TestPool.test_pool_handles_query_cancel_in_release":                                        "unknown",
	"test_pool.TestPool.test_pool_handles_task_cancel_in_release":                                         "unknown",
	"test_pool.TestPool.test_pool_init_race":                                                              "unknown",
	"test_pool.TestPool.test_pool_init_run_until_complete":                                                "unknown",
	"test_pool.TestPool.test_pool_remote_close":                                                           "unknown",
	"test_prepare.TestPrepare.test_prepare_06_interrupted_close":                                          "unknown",
	"test_prepare.TestPrepare.test_prepare_08_big_result":                                                 "unknown",
	"test_prepare.TestPrepare.test_prepare_09_raise_error":                                                "unknown",
	"test_prepare.TestPrepare.test_prepare_14_explain":                                                    "unknown",
	"test_prepare.TestPrepare.test_prepare_16_command_result":                                             "unknown",
	"test_prepare.TestPrepare.test_prepare_19_concurrent_calls":                                           "unknown",
	"test_prepare.TestPrepare.test_prepare_22_empty":                                                      "unknown",
	"test_prepare.TestPrepare.test_prepare_28_max_args":                                                   "unknown",
	"test_prepare.TestPrepare.test_prepare_31_pgbouncer_note":                                             "unknown",
	"test_prepare.TestPrepare.test_prepare_statement_invalid":                                             "unknown",
	"test_record.TestRecord.test_record_subclass_01":                                                      "unknown",
	"test_record.TestRecord.test_record_subclass_02":                                                      "unknown",
	"test_record.TestRecord.test_record_subclass_03":                                                      "unknown",
	"test_record.TestRecord.test_record_subclass_04":                                                      "unknown",
	"test_timeout.TestConnectionCommandTimeout.test_command_timeout_01":                                   "unknown",
	"test_timeout.TestTimeout.test_timeout_04":                                                            "unknown",
	"test_timeout.TestTimeout.test_timeout_06":                                                            "unknown",
	"test_utils.TestUtils.test_mogrify_multiple":                                                          "unknown",
	"test_utils.TestUtils.test_mogrify_simple":                                                            "unknown",
}
