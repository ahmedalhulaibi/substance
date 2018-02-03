package testsubstance

import (
	"github.com/ahmedalhulaibi/substance"
)

func init() {
	testPlugin := testsql{}
	substance.Register("test", &testPlugin)
}

type testsql struct {
	name string
}

/*GetCurrentDatabaseName returns currrent database schema name as string*/
func (t testsql) GetCurrentDatabaseNameFunc(dbType string, connectionString string) (string, error) {
	returnValue := "testDatabase"
	var err error
	return returnValue, err
}

/*DescribeDatabase returns tables in database*/
func (t testsql) DescribeDatabaseFunc(dbType string, connectionString string) ([]substance.ColumnDescription, error) {
	columnDesc := []substance.ColumnDescription{}
	columnDesc = append(columnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberOne",
		TableName:    "TableNumberOne",
	})
	columnDesc = append(columnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberTwo",
		TableName:    "TableNumberTwo",
	})
	columnDesc = append(columnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberThree",
		TableName:    "TableNumberThree",
	})
	return columnDesc, nil
}

/*DescribeTable returns columns in database*/
func (t testsql) DescribeTableFunc(dbType string, connectionString string, tableName string) ([]substance.ColumnDescription, error) {
	columnDesc := []substance.ColumnDescription{}
	switch tableName {
	case "TableNumberOne":
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberOne",
			PropertyType: "int32",
			PropertyName: "UniqueIdOne",
			KeyType:      "",
			Nullable:     false,
		})
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberOne",
			PropertyType: "string",
			PropertyName: "Name",
			KeyType:      "",
			Nullable:     false,
		})
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberOne",
			PropertyType: "float64",
			PropertyName: "Salary",
			KeyType:      "",
			Nullable:     true,
		})
	case "TableNumberTwo":
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberTwo",
			PropertyType: "UniqueIdTwo",
			PropertyName: "int32",
			KeyType:      "",
			Nullable:     false,
		})
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberTwo",
			PropertyType: "ForeignIdOne",
			PropertyName: "int32",
			KeyType:      "",
			Nullable:     false,
		})
	case "TableNumberThree":
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberThree",
			PropertyType: "UniqueIdThree",
			PropertyName: "int32",
			KeyType:      "",
			Nullable:     false,
		})
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberThree",
			PropertyType: "ForeignIdOne",
			PropertyName: "int32",
			KeyType:      "",
			Nullable:     false,
		})
		columnDesc = append(columnDesc, substance.ColumnDescription{
			DatabaseName: "testDatabase",
			TableName:    "TableNumberThree",
			PropertyType: "ForeignIdTwo",
			PropertyName: "int32",
			KeyType:      "",
			Nullable:     true,
		})
	}
	return columnDesc, nil
}

/*DescribeTableRelationship returns all foreign column references in database table*/
func (t testsql) DescribeTableRelationshipFunc(dbType string, connectionString string, tableName string) ([]substance.ColumnRelationship, error) {
	columnRel := []substance.ColumnRelationship{}
	switch tableName {
	case "TableNumberOne":
		columnRel = append(columnRel, substance.ColumnRelationship{
			TableName:           "TableNumberTwo",
			ColumnName:          "ForeignIdOne",
			ReferenceTableName:  "TableNumberOne",
			ReferenceColumnName: "UniqueIdOne",
		})
		columnRel = append(columnRel, substance.ColumnRelationship{
			TableName:           "TableNumberThree",
			ColumnName:          "ForeignIdOne",
			ReferenceTableName:  "TableNumberOne",
			ReferenceColumnName: "UniqueIdOne",
		})
	case "TableNumberTwo":
		columnRel = append(columnRel, substance.ColumnRelationship{
			TableName:           "TableNumberThree",
			ColumnName:          "ForeignIdTwo",
			ReferenceTableName:  "TableNumberTwo",
			ReferenceColumnName: "UniqueIdTwo",
		})
	case "TableNumberThree":
	}
	return columnRel, nil
}

func (t testsql) DescribeTableConstraintsFunc(dbType string, connectionString string, tableName string) ([]substance.ColumnConstraint, error) {
	columnConstraint := []substance.ColumnConstraint{}
	switch tableName {
	case "TableNumberOne":
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberOne",
			ColumnName:     "UniqueIdOne",
			ConstraintType: "PRIMARY KEY",
		})
	case "TableNumberTwo":
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberTwo",
			ColumnName:     "UniqueIdTwo",
			ConstraintType: "PRIMARY KEY",
		})
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberTwo",
			ColumnName:     "ForeignIdOne",
			ConstraintType: "FOREIGN KEY",
		})
	case "TableNumberThree":
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberThree",
			ColumnName:     "UniqueIdThree",
			ConstraintType: "PRIMARY KEY",
		})
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberThree",
			ColumnName:     "ForeignIdOne",
			ConstraintType: "UNIQUE",
		})
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberThree",
			ColumnName:     "ForeignIdOne",
			ConstraintType: "FOREIGN KEY",
		})
		columnConstraint = append(columnConstraint, substance.ColumnConstraint{
			TableName:      "TableNumberThree",
			ColumnName:     "ForeignIdTwo",
			ConstraintType: "FOREIGN KEY",
		})
	}
	return columnConstraint, nil
}

func (t testsql) GetGoDataType(sqlType string) (string, error) {
	return "", nil
}
