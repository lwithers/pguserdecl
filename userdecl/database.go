package userdecl

// TODO:
//  here we'll iterate through the list of databases from PostgresQL. For each:
//   - if it doesn't exist
//     - emit code to create it and its schemas
//   - if it does exist
//     - test its ownership
//     - check for any schemas that need to be created
//     - dispatch to table-level tests

const (
	queryDBs = `SELECT pg_database.datname, pg_user.usename FROM pg_database
		INNER JOIN pg_user on (pg_database.datdba = pg_user.usesysid)`

	// NB: also finds a bunch of system ones
	querySchemas = `SELECT nspname FROM pg_namespace`
)
