package testsubstance

import (
	"reflect"
	"sort"
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

func TestGetGoDataType(t *testing.T) {
	testProvider := testsql{}
	dataType, _ := testProvider.GetGoDataType("")
	expectedDataType := ""
	if dataType != expectedDataType {
		t.Errorf("Result does not match expected result: \nExpected:\n%s\nResult:\n%s", expectedDataType, dataType)
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
	columnDescResultOne, _ := testProvider.DescribeTableFunc("", "", "TableNumberOne")
	columnDescResultTwo, _ := testProvider.DescribeTableFunc("", "", "TableNumberTwo")
	columnDescResultThree, _ := testProvider.DescribeTableFunc("", "", "TableNumberThree")
	columnDescResult := append(columnDescResultOne, columnDescResultTwo...)
	columnDescResult = append(columnDescResult, columnDescResultThree...)

	sort.Slice(myColumnDesc, func(i, j int) bool {
		return myColumnDesc[i].PropertyName < myColumnDesc[j].PropertyName
	})
	sort.Slice(columnDescResult, func(i, j int) bool {
		return columnDescResult[i].PropertyName < columnDescResult[j].PropertyName
	})
	if len(columnDescResult) != len(myColumnDesc) {
		t.Errorf("Result length does not match expected length: \nExpected:\n%v\nResult:\n%v", len(myColumnDesc), len(columnDescResult))
	} else {
		for i := range columnDescResult {
			if !reflect.DeepEqual(columnDescResult[i], myColumnDesc[i]) {
				t.Errorf("Result does not match expected result: \nExpected:\n%v\nResult:\n%v", myColumnDesc, columnDescResult)
			}
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
	columnRelResultOne, _ := testProvider.DescribeTableRelationshipFunc("", "", "TableNumberOne")
	columnRelResultTwo, _ := testProvider.DescribeTableRelationshipFunc("", "", "TableNumberTwo")
	columnRelResultThree, _ := testProvider.DescribeTableRelationshipFunc("", "", "TableNumberThree")
	columnRelResult := append(columnRelResultOne, columnRelResultTwo...)
	columnRelResult = append(columnRelResult, columnRelResultThree...)

	sort.Slice(myColumnRel, func(i, j int) bool {
		return myColumnRel[i].ColumnName < myColumnRel[j].ColumnName
	})
	sort.Slice(columnRelResult, func(i, j int) bool {
		return columnRelResult[i].ColumnName < columnRelResult[j].ColumnName
	})
	if len(columnRelResult) != len(myColumnRel) {
		t.Errorf("Result length does not match expected length: \nExpected:\n%v\nResult:\n%v", len(myColumnRel), len(columnRelResult))
	} else {
		for i := range columnRelResult {
			if !reflect.DeepEqual(columnRelResult[i], myColumnRel[i]) {
				t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnRel[i], columnRelResult[i])
			}
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
	columnConstraintResultOne, _ := testProvider.DescribeTableConstraintsFunc("", "", "TableNumberOne")
	columnConstraintResultTwo, _ := testProvider.DescribeTableConstraintsFunc("", "", "TableNumberTwo")
	columnConstraintResultThree, _ := testProvider.DescribeTableConstraintsFunc("", "", "TableNumberThree")
	columnConstraintResult := append(columnConstraintResultOne, columnConstraintResultTwo...)
	columnConstraintResult = append(columnConstraintResult, columnConstraintResultThree...)

	sort.Slice(myColumnConstraint, func(i, j int) bool {
		return myColumnConstraint[i].ColumnName < myColumnConstraint[j].ColumnName
	})
	sort.Slice(columnConstraintResult, func(i, j int) bool {
		return columnConstraintResult[i].ColumnName < columnConstraintResult[j].ColumnName
	})
	if len(columnConstraintResult) != len(myColumnConstraint) {
		t.Errorf("Result length does not match expected length: \nExpected:\n%v\nResult:\n%v", len(myColumnConstraint), len(columnConstraintResult))
	} else {
		for i := range columnConstraintResult {
			if !reflect.DeepEqual(columnConstraintResult[i], myColumnConstraint[i]) {
				t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnConstraint[i], columnConstraintResult[i])
			}
		}
	}
}
