package sqlitesubstance

import (
	"database/sql"
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/ahmedalhulaibi/substance"
	/*blank import to load sqlite3 driver*/
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	pgsqlPlugin := sqlite{}
	substance.Register("sqlite3", &pgsqlPlugin)
}

type sqlite struct {
	name string
}

/*GetCurrentDatabaseName returns currrent database schema name as string*/
func (p sqlite) DatabaseName(dbType string, db *sql.DB) (string, error) {
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
				case "file":
					returnValue = path.Base(string(value.([]byte)))
				}
			case string:
				switch queryResult.Columns[i] {
				case "file":
					returnValue = path.Base(string(value.(string)))
				}
			}
		}

	}

	return returnValue, nil
}

/*DescribeDatabase returns tables in database*/
func (p sqlite) DescribeDatabase(dbType string, db *sql.DB) ([]substance.ColumnDescription, error) {
	queryResult := substance.ExecuteQuery(dbType, db, "", DescribeDatabaseQuery)

	if queryResult.Err != nil {
		return nil, queryResult.Err
	}

	//setup array of column descriptions
	columnDesc := []substance.ColumnDescription{}

	//get database name
	databaseName, err := p.DatabaseName(dbType, db)
	if err != nil {
		return nil, err
	}

	newColDesc := substance.ColumnDescription{DatabaseName: databaseName, PropertyType: "Table"}

	for queryResult.Rows.Next() {
		err = queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case []byte:
				switch queryResult.Columns[i] {
				case "name":
					newColDesc.TableName = string(value.([]byte))
					newColDesc.PropertyName = string(value.([]byte))
				}
			case string:
				switch queryResult.Columns[i] {
				case "name":
					newColDesc.TableName = value.(string)
					newColDesc.PropertyName = value.(string)
				}
			}
		}
		columnDesc = append(columnDesc, newColDesc)
	}
	return columnDesc, nil
}

/*DescribeTable returns columns in database*/
func (p sqlite) DescribeTable(dbType string, db *sql.DB, tableName string) ([]substance.ColumnDescription, error) {
	queryResult := substance.ExecuteQuery(dbType, db, tableName, strings.Replace(DescribeTableQuery, "$1", tableName, -1))
	if queryResult.Err != nil {
		return nil, queryResult.Err
	}

	columnDesc := []substance.ColumnDescription{}

	databaseName, err := p.DatabaseName(dbType, db)
	if err != nil {
		return nil, err
	}

	newColDesc := substance.ColumnDescription{DatabaseName: databaseName, TableName: tableName}

	for queryResult.Rows.Next() {
		err := queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case int64:
				switch queryResult.Columns[i] {
				case "notnull":
					newColDesc.Nullable = !(value.(int64) != 0)
				}
			case []byte:
				switch queryResult.Columns[i] {
				case "name":
					newColDesc.PropertyName = string(value.([]byte))
				case "type":
					newColDesc.PropertyType, _ = p.ToGoDataType(string(value.([]byte)))
				case "dflt_value":
					newColDesc.DefaultValue = string(value.([]byte))
				}
			case string:
				switch queryResult.Columns[i] {
				case "name":
					newColDesc.PropertyName = value.(string)
				case "type":
					newColDesc.PropertyType, _ = p.ToGoDataType(value.(string))
				case "dflt_value":
					newColDesc.DefaultValue = value.(string)
				}
			}

		}
		columnDesc = append(columnDesc, newColDesc)

	}
	return columnDesc, nil
}

/*TableRelationships returns all foreign column references in database table*/
func (p sqlite) TableRelationships(dbType string, db *sql.DB, tableName string) ([]substance.ColumnRelationship, error) {
	queryResult := substance.ExecuteQuery(dbType, db, tableName, strings.Replace(DescribeTableRelationshipQuery, "$1", tableName, -1))
	if queryResult.Err != nil {
		return nil, queryResult.Err
	}

	columnRel := []substance.ColumnRelationship{}
	newColRel := substance.ColumnRelationship{}

	for queryResult.Rows.Next() {
		err := queryResult.Rows.Scan(queryResult.ScanArgs...)
		if err != nil {
			return nil, err
		}

		// Print data
		for i, value := range queryResult.Values {
			switch value.(type) {
			case nil:
				//fmt.Println("\t", columns[i], ": NULL")
				err := fmt.Errorf("Null column value found at column: '%s' index: '%d'", queryResult.Columns[i], i)
				return nil, error(err)
			case []byte:
				switch queryResult.Columns[i] {
				case "from":
					newColRel.ColumnName = string(value.([]byte))
				case "table":
					newColRel.ReferenceTableName = string(value.([]byte))
				case "to":
					newColRel.ReferenceColumnName = string(value.([]byte))
				}
			case string:
				switch queryResult.Columns[i] {
				case "from":
					newColRel.ColumnName = value.(string)
				case "table":
					newColRel.ReferenceTableName = value.(string)
				case "to":
					newColRel.ReferenceColumnName = value.(string)
				}
			default:
				//fmt.Println("\t", columns[i], ": ", value)
			}
		}
		newColRel.TableName = tableName
		columnRel = append(columnRel, newColRel)
		//fmt.Println("-----------------------------------")
	}
	return columnRel, nil
}

/*TableConstraints returns an array of ColumnConstraint objects*/
func (p sqlite) TableConstraints(dbType string, db *sql.DB, tableName string) ([]substance.ColumnConstraint, error) {
	columnCon := []substance.ColumnConstraint{}
	newColCon := substance.ColumnConstraint{}

	//getting column relationships to retrieve foreign key constraints
	columnRels, err := p.TableRelationships(dbType, db, tableName)
	if err != nil {
		return nil, err
	}

	//populate foreign keys using column relationships
	for _, columnRel := range columnRels {
		if columnRel.TableName == tableName {
			newColCon.TableName = tableName
			newColCon.ColumnName = columnRel.ColumnName
			newColCon.ConstraintType = "f"
		}
		columnCon = append(columnCon, newColCon)
	}

	indexListResult := substance.ExecuteQuery(dbType, db, tableName, strings.Replace(SQLLiteIndexList, "$1", tableName, -1))
	if indexListResult.Err != nil {
		return nil, indexListResult.Err
	}

	for indexListResult.Rows.Next() {
		var seq int64
		var name string
		var unique int64
		var origin string
		var partial int64
		err = indexListResult.Rows.Scan(&seq, &name, &unique, &origin, &partial)
		if err != nil {
			return nil, err
		}
		if origin == "pk" {
			origin = "p"
		}

		indexInfoResult := substance.ExecuteQuery(dbType, db, "", strings.Replace(SQLLiteIndexInfo, "$1", name, -1))
		for indexInfoResult.Rows.Next() {
			var seqno int64
			var cid int64
			var colName string
			err = indexInfoResult.Rows.Scan(&seqno, &cid, &colName)
			if err != nil {
				return nil, err
			}
			newColCon.ColumnName = colName
			newColCon.ConstraintType = origin
			newColCon.TableName = tableName
			columnCon = append(columnCon, newColCon)
			if unique == 1 && origin != "u" {
				newColCon.ColumnName = colName
				newColCon.ConstraintType = "u"
				newColCon.TableName = tableName
				columnCon = append(columnCon, newColCon)
			}

		}
	}
	return columnCon, nil
}

func (p sqlite) ToGoDataType(sqlType string) (string, error) {
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
