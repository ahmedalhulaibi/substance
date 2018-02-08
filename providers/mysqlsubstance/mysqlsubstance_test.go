package mysqlsubstance

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ahmedalhulaibi/substance"
)

func TestGetCurrDbName(t *testing.T) {
	mysqlProvider := mysql{}
	nameExpected := "delivery"
	nameResult, err := mysqlProvider.GetCurrentDatabaseNameFunc("mysql", "travis@tcp(127.0.0.1:3306)/delivery")
	if nameResult != nameExpected {
		t.Errorf("Expected '%s' as database name but got '%s'.", nameExpected, nameResult)
	}
	if err != nil {
		t.Error(err)
	}
}
func TestGetGoDataType(t *testing.T) {
	mysqlProvider := mysql{}
	sqlDatatypes := []string{"BIT",
		"BIT(9)",
		"BOOL",
		"BOOLEAN",
		"TINYINT",
		"TINYINT(1)",
		"SMALLINT",
		"SMALLINT(4)",
		"MEDIUMINT",
		"MEDIUMINT(9)",
		"INT",
		"INT(10)",
		"INTEGER",
		"INTEGER(11)",
		"BIGINT",
		"BIGINT(40)",
		"DECIMAL",
		"DECIMAL(65)",
		"DECIMAL(66)",
		"DEC",
		"DEC(65)",
		"DEC(66)",
		"DEC(65,30)",
		"DEC(65,31)",
		"DEC(66,31)",
		"FLOAT",
		"FLOAT(5,2)",
		"DOUBLE",
		"DOUBLE(5,2)",
		"DOUBLE PRECISION",
		"DOUBLE PRECISION(5,2)",
		"UNSIGNED TINYINT",
		"UNSIGNED TINYINT(1)",
		"UNSIGNED SMALLINT",
		"UNSIGNED SMALLINT(4)",
		"UNSIGNED MEDIUMINT",
		"UNSIGNED MEDIUMINT(9)",
		"UNSIGNED INT",
		"UNSIGNED INT(10)",
		"UNSIGNED INTEGER",
		"UNSIGNED INTEGER(11)",
		"UNSIGNED BIGINT",
		"UNSIGNED BIGINT(40)",
		"UNSIGNED DECIMAL",
		"UNSIGNED DECIMAL(65)",
		"UNSIGNED DECIMAL(66)",
		"UNSIGNED DEC",
		"UNSIGNED DEC(65)",
		"UNSIGNED DEC(66)",
		"UNSIGNED DEC(65,30)",
		"UNSIGNED DEC(65,31)",
		"UNSIGNED DEC(66,31)",
		"UNSIGNED FLOAT",
		"UNSIGNED FLOAT(5,2)",
		"UNSIGNED DOUBLE",
		"UNSIGNED DOUBLE(5,2)",
		"UNSIGNED DOUBLE PRECISION",
		"UNSIGNED DOUBLE PRECISION(5,2)",
	}
	for _, sqlDatatype := range sqlDatatypes {
		goDataType, err := mysqlProvider.GetGoDataType(strings.ToLower(sqlDatatype))
		if err != nil {
			t.Error(err)
		} else if testing.Verbose() {
			fmt.Printf("SQL Type: %s = Go Type: %s\n", sqlDatatype, goDataType)
		}
	}
}
func TestDescribeDb(t *testing.T) {
	mysqlProvider := mysql{}
	myColumnDesc := []substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "Table",
		PropertyName: "AntiOrders",
		TableName:    "AntiOrders",
	}, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "Table",
		PropertyName: "Orders",
		TableName:    "Orders",
	}, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "Table",
		PropertyName: "Persons",
		TableName:    "Persons",
	})
	columnDescResult, err := mysqlProvider.DescribeDatabaseFunc("mysql", "travis@tcp(127.0.0.1:3306)/delivery")
	if err != nil {
		t.Error(err)
	}
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

func TestDescribeTable(t *testing.T) {
	mysqlProvider := mysql{}
	myColumnDesc := []substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "int32",
		PropertyName: "ID",
		TableName:    "Persons",
		Nullable:     false,
		KeyType:      "PRI",
	}, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "string",
		PropertyName: "LastName",
		TableName:    "Persons",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "string",
		PropertyName: "FirstName",
		TableName:    "Persons",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "string",
		PropertyName: "Address",
		TableName:    "Persons",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "delivery",
		PropertyType: "string",
		PropertyName: "City",
		TableName:    "Persons",
		Nullable:     true,
	})
	columnDescResult, err := mysqlProvider.DescribeTableFunc("mysql", "travis@tcp(127.0.0.1:3306)/delivery", "Persons")
	if err != nil {
		t.Error(err)
	}
	sort.Slice(myColumnDesc, func(i, j int) bool {
		return myColumnDesc[i].PropertyName < myColumnDesc[j].PropertyName
	})
	sort.Slice(columnDescResult, func(i, j int) bool {
		return columnDescResult[i].PropertyName < columnDescResult[j].PropertyName
	})
	if len(columnDescResult) != len(myColumnDesc) {
		t.Errorf("Result length does not match expected length: \nExpected:\n%v\nResult:\n%v", len(myColumnDesc), len(columnDescResult))
		t.Errorf("Result does not match expected result: \nExpected:\n%v\nResult:\n%v", myColumnDesc, columnDescResult)
	} else {
		for i := range columnDescResult {
			if !reflect.DeepEqual(columnDescResult[i], myColumnDesc[i]) {
				t.Errorf("Result does not match expected result: \nExpected:\n%v\nResult:\n%v", myColumnDesc, columnDescResult)
			}
		}
	}
}

func TestDescribeTableRelationship(t *testing.T) {
	mysqlProvider := mysql{}
	myColumnRel := []substance.ColumnRelationship{}
	myColumnRel = append(myColumnRel, substance.ColumnRelationship{
		TableName:           "AntiOrders",
		ColumnName:          "PersonID",
		ReferenceTableName:  "Persons",
		ReferenceColumnName: "ID",
	}, substance.ColumnRelationship{
		TableName:           "Orders",
		ColumnName:          "PersonID",
		ReferenceTableName:  "Persons",
		ReferenceColumnName: "ID",
	})
	columnRelResult, err := mysqlProvider.DescribeTableRelationshipFunc("mysql", "travis@tcp(127.0.0.1:3306)/delivery", "Persons")
	if err != nil {
		t.Error(err)
	}
	sort.Slice(myColumnRel, func(i, j int) bool {
		return myColumnRel[i].ColumnName < myColumnRel[j].ColumnName
	})
	sort.Slice(columnRelResult, func(i, j int) bool {
		return columnRelResult[i].ColumnName < columnRelResult[j].ColumnName
	})
	if len(columnRelResult) != len(myColumnRel) {
		t.Errorf("Result length does not match expected length: \nExpected:\n%v\nResult:\n%v", len(myColumnRel), len(columnRelResult))
		t.Errorf("Result does not match expected result: \nExpected:\n%v\nResult:\n%v", myColumnRel, columnRelResult)
	} else {
		for i := range columnRelResult {
			if !reflect.DeepEqual(columnRelResult[i], myColumnRel[i]) {
				t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnRel[i], columnRelResult[i])
			}
		}
	}
}

func TestDescribeTableContraints(t *testing.T) {
	mysqlProvider := mysql{}
	myColumnConstraint := []substance.ColumnConstraint{}
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "AntiOrderID",
		ConstraintType: "PRIMARY KEY",
	}, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "AntiOrderID",
		ConstraintType: "UNIQUE",
	}, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "PersonID",
		ConstraintType: "FOREIGN KEY",
	}, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "PersonID",
		ConstraintType: "UNIQUE",
	})
	columnConstraintResult, err := mysqlProvider.DescribeTableConstraintsFunc("mysql", "travis@tcp(127.0.0.1:3306)/delivery", "AntiOrders")
	if err != nil {
		t.Error(err)
	}
	sort.Slice(myColumnConstraint, func(i, j int) bool {
		return myColumnConstraint[i].ColumnName < myColumnConstraint[j].ColumnName
	})
	sort.Slice(columnConstraintResult, func(i, j int) bool {
		return columnConstraintResult[i].ColumnName < columnConstraintResult[j].ColumnName
	})
	if len(columnConstraintResult) != len(myColumnConstraint) {
		t.Errorf("Result length does not match expected length: \nExpected:\n%v\nResult:\n%v", len(myColumnConstraint), len(columnConstraintResult))
		t.Errorf("Result does not match expected result: \nExpected:\n%v\nResult:\n%v", myColumnConstraint, columnConstraintResult)
	} else {
		for i := range columnConstraintResult {
			if !reflect.DeepEqual(columnConstraintResult[i], myColumnConstraint[i]) {
				t.Errorf("Result does not match expected result: \nExpected:\n\t%v\nResult:\n%v\n\n", myColumnConstraint[i], columnConstraintResult[i])
			}
		}
	}
}
