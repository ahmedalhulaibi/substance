package substance

import "database/sql"

/*SubstanceProviderInterface */
type SubstanceProviderInterface interface {
	DatabaseName(dbType string, db *sql.DB) (string, error)
	DescribeDatabase(dbType string, db *sql.DB) ([]ColumnDescription, error)
	DescribeTable(dbType string, db *sql.DB, tableName string) ([]ColumnDescription, error)
	TableRelationships(dbType string, db *sql.DB, tableName string) ([]ColumnRelationship, error)
	TableConstraints(dbType string, db *sql.DB, tableName string) ([]ColumnConstraint, error)
	ToGoDataType(sqlType string) (string, error)
}

/*substance plugin map*/
var substancePlugins = make(map[string]SubstanceProviderInterface)

/*Register registers a sbustance plugin which implements the Substance interface*/
func Register(pluginName string, pluginInterface SubstanceProviderInterface) {
	//fmt.Println(substancePlugins)
	substancePlugins[pluginName] = pluginInterface
}

/*ColumnDescription Structure to store properties of each column in a table */
type ColumnDescription struct {
	DatabaseName string
	TableName    string
	PropertyName string
	PropertyType string
	KeyType      string
	DefaultValue string
	Nullable     bool
}

/*ColumnRelationship Structure to store relationships between tables*/
type ColumnRelationship struct {
	TableName           string
	ColumnName          string
	ReferenceTableName  string
	ReferenceColumnName string
}

/*ColumnConstraint Struct to store column constraint types*/
type ColumnConstraint struct {
	TableName      string
	ColumnName     string
	ConstraintType string
}

/*QueryResult Struct to store results from ExecuteQuery*/
type QueryResult struct {
	Rows     *sql.Rows
	Columns  []string
	Values   []interface{}
	ScanArgs []interface{}
	Err      error
}

/*GetCurrentDatabaseName returns currrent database schema name as string*/
func GetCurrentDatabaseName(dbType string, db *sql.DB) (string, error) {
	return substancePlugins[dbType].DatabaseName(dbType, db)
}

/*DescribeDatabase returns tables in database*/
func DescribeDatabase(dbType string, db *sql.DB) ([]ColumnDescription, error) {
	return substancePlugins[dbType].DescribeDatabase(dbType, db)
}

/*DescribeTable returns columns of a table*/
func DescribeTable(dbType string, db *sql.DB, tableName string) ([]ColumnDescription, error) {
	return substancePlugins[dbType].DescribeTable(dbType, db, tableName)
}

/*DescribeTableRelationship returns all foreign column references in database table*/
func DescribeTableRelationship(dbType string, db *sql.DB, tableName string) ([]ColumnRelationship, error) {
	return substancePlugins[dbType].TableRelationships(dbType, db, tableName)
}

/*DescribeTableConstraints returns all column constraints in a database table*/
func DescribeTableConstraints(dbType string, db *sql.DB, tableName string) ([]ColumnConstraint, error) {
	return substancePlugins[dbType].TableConstraints(dbType, db, tableName)
}

/*ExecuteQuery executes a sql query with one or no tableName, specific to mysqlsubstnace and pgsqlsubstance*/
func ExecuteQuery(dbType string, db *sql.DB, tableName string, query string) QueryResult {
	var (
		rows *sql.Rows
		err  error
	)
	if tableName == "" {
		rows, err = db.Query(query)
	} else {
		rows, err = db.Query(query, tableName)
	}
	if err != nil {
		return QueryResult{Err: err, Rows: rows}
	}

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return QueryResult{Err: err, Rows: rows, Columns: columns}
	}
	// Make a slice for the values
	values := make([]interface{}, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	return QueryResult{Err: err, Rows: rows, Columns: columns, ScanArgs: scanArgs, Values: values}
}
