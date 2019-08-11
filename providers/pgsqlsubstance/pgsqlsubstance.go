package pgsqlsubstance

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ahmedalhulaibi/substance"
	/*blank import to load postgres driver*/
	_ "github.com/lib/pq"
)

func init() {
	pgsqlPlugin := pgsql{}
	substance.Register("postgres", &pgsqlPlugin)
}

type pgsql struct {
	name string
}

/*GetCurrentDatabaseName returns currrent database schema name as string*/
func (p pgsql) DatabaseName(dbType string, db *sql.DB) (string, error) {
	returnValue := "placeholder"

	queryResult := substance.ExecuteQuery(dbType, db, "", GetCurrentDatabaseNameQuery)
	if queryResult.Err != nil {
		return "", queryResult.Err
	}

	for queryResult.Rows.Next() {
		err := queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return "", err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case []byte:
				switch queryResult.Columns[i] {
				case "current_database":
					returnValue = string(value.([]byte))
				}
			}
		}

	}

	return returnValue, nil
}

/*DescribeDatabase returns tables in database*/
func (p pgsql) DescribeDatabase(dbType string, db *sql.DB) ([]*substance.ColumnDescription, error) {
	queryResult := substance.ExecuteQuery(dbType, db, "", DescribeDatabaseQuery)

	if queryResult.Err != nil {
		return nil, queryResult.Err
	}

	//setup array of column descriptions
	columnDesc := []*substance.ColumnDescription{}

	//get database name
	databaseName, err := p.DatabaseName(dbType, db)
	if err != nil {
		return nil, err
	}

	for queryResult.Rows.Next() {
		newColDesc := &substance.ColumnDescription{DatabaseName: databaseName, PropertyType: "Table"}
		err = queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case []byte:
				switch queryResult.Columns[i] {
				case "tablename":
					newColDesc.TableName = string(value.([]byte))
					newColDesc.PropertyName = string(value.([]byte))
				}
			}
		}
		columnDesc = append(columnDesc, newColDesc)
	}
	return columnDesc, nil
}

/*DescribeTable returns columns in database*/
func (p pgsql) DescribeTable(dbType string, db *sql.DB, tableName string) ([]*substance.ColumnDescription, error) {
	queryResult := substance.ExecuteQuery(dbType, db, tableName, DescribeTableQuery)

	if queryResult.Err != nil {
		return nil, queryResult.Err
	}

	columnDesc := []*substance.ColumnDescription{}

	databaseName, err := p.DatabaseName(dbType, db)
	if err != nil {
		return nil, err
	}

	for queryResult.Rows.Next() {
		newColDesc := &substance.ColumnDescription{DatabaseName: databaseName, TableName: tableName}
		err = queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case bool:
				switch queryResult.Columns[i] {
				case "isNotNull":
					newColDesc.Nullable = !value.(bool)
				}
			case []byte:
				switch queryResult.Columns[i] {
				case "Field":
					newColDesc.PropertyName = string(value.([]byte))
				case "Type":
					newColDesc.PropertyType, _ = p.ToGoDataType(string(value.([]byte)))
				}
			}
		}
		columnDesc = append(columnDesc, newColDesc)

	}
	return columnDesc, nil
}

/*TableRelationships returns all foreign column references in database table*/
func (p pgsql) TableRelationships(dbType string, db *sql.DB, tableName string) ([]*substance.ColumnRelationship, error) {
	queryResult := substance.ExecuteQuery(dbType, db, tableName, DescribeTableRelationshipQuery)
	if queryResult.Err != nil {
		return nil, queryResult.Err
	}

	columnTableDesc, err := substance.DescribeTable(dbType, db, tableName)
	if err != nil {
		return nil, err
	}
	columnDesc := []*substance.ColumnRelationship{}

	for queryResult.Rows.Next() {
		newColDesc := &substance.ColumnRelationship{}
		err = queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {

			switch value.(type) {
			case string:

				switch queryResult.Columns[i] {
				case "table_name":
					newColDesc.TableName = string(value.(string))
				case "column":
					newColDesc.ColumnName = string(value.(string))
				}
			case []byte:

				switch queryResult.Columns[i] {
				case "ref_table":
					newColDesc.ReferenceTableName = string(value.([]byte))
					columnTableDesc, err = substance.DescribeTable(dbType, db, newColDesc.ReferenceTableName)
					if err != nil {
						return nil, err
					}
				case "ref_columnNum":
					//this gets returned as {1} a reference to the column number in the table
					//this has to be replaced with the column name

					refColumnNumStr := strings.Replace(strings.Replace(string(value.([]byte)), "{", "", -1), "}", "", -1)

					refColumnNum, err := strconv.Atoi(refColumnNumStr)
					if err != nil {
						return nil, err
					}

					newColDesc.ReferenceColumnName = columnTableDesc[refColumnNum-1].PropertyName

				}
			}
		}
		columnDesc = append(columnDesc, newColDesc)
	}
	return columnDesc, nil
}

/*TableConstraints returns an array of ColumnConstraint objects*/
func (p pgsql) TableConstraints(dbType string, db *sql.DB, tableName string) ([]*substance.ColumnConstraint, error) {
	queryResult := substance.ExecuteQuery(dbType, db, tableName, DescribeTableConstraintsQuery)
	if queryResult.Err != nil {
		return nil, queryResult.Err
	}
	columnDesc := []*substance.ColumnConstraint{}

	for queryResult.Rows.Next() {
		newColDesc := &substance.ColumnConstraint{}
		err := queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case string:

				switch queryResult.Columns[i] {
				case "table_name":
					newColDesc.TableName = string(value.(string))
				case "column":
					newColDesc.ColumnName = string(value.(string))
				case "contype":
					newColDesc.ConstraintType = string(value.(string))
				}
			default:

			}
		}
		columnDesc = append(columnDesc, newColDesc)
	}
	return columnDesc, nil
}

func (p pgsql) ToGoDataType(sqlType string) (string, error) {
	if regexDataTypePatterns == nil {
		regexDataTypePatterns["bit.*"] = "int64"
		regexDataTypePatterns["bool.*|tinyint\\(1\\)"] = "bool"
		regexDataTypePatterns["tinyint.*"] = "int8"
		regexDataTypePatterns["unsigned\\stinyint.*"] = "uint8"
		regexDataTypePatterns["smallint.*"] = "int16"
		regexDataTypePatterns["unsigned\\ssmallint.*"] = "uint16"
		regexDataTypePatterns["(mediumint.*|int.*)"] = "int32"
		regexDataTypePatterns["unsigned\\s(mediumint.*|int.*)"] = "uint32"
		regexDataTypePatterns["bigint.*"] = "int64"
		regexDataTypePatterns["unsigned\\sbigint.*"] = "uint64"
		regexDataTypePatterns["(unsigned\\s){0,1}(double.*|float.*|dec.*)"] = "float64"
		regexDataTypePatterns["varchar.*|date.*|time.*|year.*|char.*|.*text.*|enum.*|set.*|.*blob.*|.*binary.*"] = "string"
	}

	for pattern, value := range regexDataTypePatterns {
		match, err := regexp.MatchString(pattern, sqlType)
		if match && err == nil {
			result := value
			return result, nil
		}
	}
	err := fmt.Errorf("No match found for data type %s", sqlType)
	fmt.Println(err)
	return sqlType, err
}
