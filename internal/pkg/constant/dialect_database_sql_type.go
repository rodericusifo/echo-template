package constant

type DialectDatabaseSQLType string

var (
	POSTGRES = DialectDatabaseSQLType("postgres")
	MYSQL    = DialectDatabaseSQLType("mysql")
)
