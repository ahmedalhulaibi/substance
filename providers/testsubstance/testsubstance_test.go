package testsubstance

import (
	"testing"

	"github.com/ahmedalhulaibi/substance"
)

func TestGetCurrDbName(t *testing.T) {
	testProvider := testsql{}
	name, _ := testProvider.GetCurrentDatabaseNameFunc("", "")
	if name != "testDatabase" {
		t.Errorf("Expected database name 'testDatabase' got '%s' instead.", name)
	}
}

func TestDescribeDb(t *testing.T) {
	testProvider := testsql{}
	myColumnDesc := []substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberOne",
		TableName:    "TableNumberOne",
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberTwo",
		TableName:    "TableNumberTwo",
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberThree",
		TableName:    "TableNumberThree",
	})
	columnDescResult, _ := testProvider.DescribeDatabaseFunc("", "")
	for i := range columnDescResult {
		if columnDescResult[i] != myColumnDesc[i] {
			t.Errorf("Result does not match case: \nExpected:\n%v\nResult:\n%v", myColumnDesc, columnDescResult)
		}
	}
}

func TestDescribeTable(t *testing.T) {
	testProvider := testsql{}
	myColumnDesc := []substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberOne",
		PropertyType: "int32",
		PropertyName: "UniqueIdOne",
		KeyType:      "p",
		Nullable:     false,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberOne",
		PropertyType: "string",
		PropertyName: "Name",
		KeyType:      "",
		Nullable:     false,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberOne",
		PropertyType: "float64",
		PropertyName: "Salary",
		KeyType:      "",
		Nullable:     true,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberTwo",
		PropertyName: "UniqueIdTwo",
		PropertyType: "int32",
		KeyType:      "",
		Nullable:     false,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberTwo",
		PropertyName: "ForeignIdOne",
		PropertyType: "int32",
		KeyType:      "f",
		Nullable:     false,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberThree",
		PropertyName: "UniqueIdThree",
		PropertyType: "int32",
		KeyType:      "",
		Nullable:     false,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberThree",
		PropertyName: "ForeignIdOne",
		PropertyType: "int32",
		KeyType:      "f",
		Nullable:     false,
	})
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberThree",
		PropertyName: "ForeignIdTwo",
		PropertyType: "int32",
		KeyType:      "f",
		Nullable:     true,
	})
	columnDescResult, _ := testProvider.DescribeTableFunc("", "", "")
	for i := range columnDescResult {
		if columnDescResult[i] != myColumnDesc[i] {
			t.Errorf("Result does not match case: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnDesc[i], columnDescResult[i])
		}
	}
}

func TestDescribeTableRelationship(t *testing.T) {
	testProvider := testsql{}
	myColumnRel := []substance.ColumnRelationship{}
	myColumnRel = append(myColumnRel, substance.ColumnRelationship{
		TableName:           "TableNumberTwo",
		ColumnName:          "ForeignIdOne",
		ReferenceTableName:  "TableNumberOne",
		ReferenceColumnName: "UniqueIdOne",
	})
	myColumnRel = append(myColumnRel, substance.ColumnRelationship{
		TableName:           "TableNumberThree",
		ColumnName:          "ForeignIdOne",
		ReferenceTableName:  "TableNumberOne",
		ReferenceColumnName: "UniqueIdOne",
	})
	myColumnRel = append(myColumnRel, substance.ColumnRelationship{
		TableName:           "TableNumberThree",
		ColumnName:          "ForeignIdTwo",
		ReferenceTableName:  "TableNumberTwo",
		ReferenceColumnName: "UniqueIdTwo",
	})
	columnRelResult, _ := testProvider.DescribeTableRelationshipFunc("", "", "")
	for i := range columnRelResult {
		if columnRelResult[i] != myColumnRel[i] {
			t.Errorf("Result does not match case: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnRel[i], columnRelResult[i])
		}
	}
}

func TestDescribeTableContraints(t *testing.T) {
	testProvider := testsql{}
	myColumnConstraint := []substance.ColumnConstraint{}
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberOne",
		ColumnName:     "UniqueIdOne",
		ConstraintType: "p",
	})
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberTwo",
		ColumnName:     "UniqueIdTwo",
		ConstraintType: "p",
	})
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberTwo",
		ColumnName:     "ForeignIdOne",
		ConstraintType: "f",
	})
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "UniqueIdThree",
		ConstraintType: "p",
	})
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "ForeignIdOne",
		ConstraintType: "u",
	})
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "ForeignIdOne",
		ConstraintType: "f",
	})
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "ForeignIdTwo",
		ConstraintType: "f",
	})
	columnConstraintResult, _ := testProvider.DescribeTableConstraintsFunc("", "", "")
	for i := range columnConstraintResult {
		if columnConstraintResult[i] != myColumnConstraint[i] {
			t.Errorf("Result does not match case: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnConstraint[i], columnConstraintResult[i])
		}
	}
}
