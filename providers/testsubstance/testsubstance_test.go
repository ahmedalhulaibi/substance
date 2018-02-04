package testsubstance

import (
	"testing"

	"github.com/ahmedalhulaibi/substance"
)

func TestGetCurrDbName(t *testing.T) {
	testProvider := testsql{}
	name, _ := testProvider.GetCurrentDatabaseNameFunc("", "")
	expectedName := "testDatabase"
	if name != expectedName {
		t.Errorf("Result does not match expected result: \nExpected:\n%s\nResult:\n%s", expectedName, name)
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
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberTwo",
		TableName:    "TableNumberTwo",
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		PropertyType: "Table",
		PropertyName: "TableNumberThree",
		TableName:    "TableNumberThree",
	})
	columnDescResult, _ := testProvider.DescribeDatabaseFunc("", "")
	for i := range columnDescResult {
		if columnDescResult[i] != myColumnDesc[i] {
			t.Errorf("Result does not match expected result: \nExpected:\n%v\nResult:\n%v", myColumnDesc, columnDescResult)
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
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberOne",
		PropertyType: "string",
		PropertyName: "Name",
		KeyType:      "",
		Nullable:     false,
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberOne",
		PropertyType: "float64",
		PropertyName: "Salary",
		KeyType:      "",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberTwo",
		PropertyName: "UniqueIdTwo",
		PropertyType: "int32",
		KeyType:      "",
		Nullable:     false,
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberTwo",
		PropertyName: "ForeignIdOne",
		PropertyType: "int32",
		KeyType:      "f",
		Nullable:     false,
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberThree",
		PropertyName: "UniqueIdThree",
		PropertyType: "int32",
		KeyType:      "",
		Nullable:     false,
	}, substance.ColumnDescription{
		DatabaseName: "testDatabase",
		TableName:    "TableNumberThree",
		PropertyName: "ForeignIdOne",
		PropertyType: "int32",
		KeyType:      "f",
		Nullable:     false,
	}, substance.ColumnDescription{
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
			t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnDesc[i], columnDescResult[i])
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
	}, substance.ColumnRelationship{
		TableName:           "TableNumberThree",
		ColumnName:          "ForeignIdOne",
		ReferenceTableName:  "TableNumberOne",
		ReferenceColumnName: "UniqueIdOne",
	}, substance.ColumnRelationship{
		TableName:           "TableNumberThree",
		ColumnName:          "ForeignIdTwo",
		ReferenceTableName:  "TableNumberTwo",
		ReferenceColumnName: "UniqueIdTwo",
	})
	columnRelResult, _ := testProvider.DescribeTableRelationshipFunc("", "", "")
	for i := range columnRelResult {
		if columnRelResult[i] != myColumnRel[i] {
			t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnRel[i], columnRelResult[i])
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
	}, substance.ColumnConstraint{
		TableName:      "TableNumberTwo",
		ColumnName:     "UniqueIdTwo",
		ConstraintType: "p",
	}, substance.ColumnConstraint{
		TableName:      "TableNumberTwo",
		ColumnName:     "ForeignIdOne",
		ConstraintType: "f",
	}, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "UniqueIdThree",
		ConstraintType: "p",
	}, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "ForeignIdOne",
		ConstraintType: "u",
	}, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "ForeignIdOne",
		ConstraintType: "f",
	}, substance.ColumnConstraint{
		TableName:      "TableNumberThree",
		ColumnName:     "ForeignIdTwo",
		ConstraintType: "f",
	})
	columnConstraintResult, _ := testProvider.DescribeTableConstraintsFunc("", "", "")
	for i := range columnConstraintResult {
		if columnConstraintResult[i] != myColumnConstraint[i] {
			t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnConstraint[i], columnConstraintResult[i])
		}
	}
}
