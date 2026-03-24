package templates

const SqlcConfig = `
- name: {{.Name.Flat}}_postgres
  schema: /db/{{.Name.Flat}}/postgres_schema.sql
  queries: /db/{{.Name.Flat}}/postgres_queries.sql
  engine: postgresql
  analyzer:
    database: false
  strict_function_checks: false
  strict_order_by: true
  gen:
    go:
      package: outbound
      out: /internal/{{.Name.Flat}}/adapters/outbound
      sql_package: pgx/v5
      sql_driver: github.com/jackc/pgx/v5
      json_tags_id_uppercase: false
      emit_db_tags: true
      emit_json_tags: true
      emit_sql_as_comment: true
      emit_prepared_queries: false
      emit_interface: true
      emit_exact_table_names: false
      emit_methods_with_db_argument: false
      emit_pointers_for_null_types: false
      omit_unused_structs: false
      output_db_file_name: postgres_database.gen.go
      output_files_suffix: .gen.go
      output_models_file_name: postgres_models.gen.go
      output_querier_file_name: postgres_query_interface.gen.go
      query_parameter_limit: 0 # always return a struct, even if its empty`
