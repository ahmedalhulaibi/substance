package pgsqlsubstance

import (
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/ahmedalhulaibi/substance"
)

func TestGetCurrDbName(t *testing.T) {
	db, err := sql.Open("postgres", os.Getenv("SUBSTANCE_PGSQL"))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	pgsqlProvider := pgsql{}
	nameExpected := "postgres"

	nameResult, err := pgsqlProvider.DatabaseName("postgres", db)
	fmt.Println(os.Getenv("SUBSTANCE_PGSQL"))
	if nameResult != nameExpected {
		t.Errorf("Expected '%s' as database name but got '%s'.", nameExpected, nameResult)
	}
	if err != nil {
		t.Error(err)
	}
}
func TestGetGoDataType(t *testing.T) {
	pgsqlProvider := pgsql{}
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
		"REAL",
		"POINT",
		"NUMERIC",
		"NUMERIC(5,4)",
		"JSONB",
		"JSON",
	}
	for _, sqlDatatype := range sqlDatatypes {
		goDataType, err := pgsqlProvider.ToGoDataType(strings.ToLower(sqlDatatype))
		if err != nil {
			t.Error(err)
		} else if testing.Verbose() {
			fmt.Printf("SQL Type: %s = Go Type: %s\n", sqlDatatype, goDataType)
		}
	}
}
func TestDescribeDb(t *testing.T) {
	db, err := sql.Open("postgres", os.Getenv("SUBSTANCE_PGSQL"))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	pgsqlProvider := pgsql{}
	myColumnDesc := []*substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "Table",
		PropertyName: "persons",
		TableName:    "persons",
	}, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "Table",
		PropertyName: "orders",
		TableName:    "orders",
	}, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "Table",
		PropertyName: "antiorders",
		TableName:    "antiorders",
	})

	columnDescResult, err := pgsqlProvider.DescribeDatabase("postgres", db)
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

func TestDescribeTable(t *testing.T) {
	db, err := sql.Open("postgres", os.Getenv("SUBSTANCE_PGSQL"))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	pgsqlProvider := pgsql{}
	myColumnDesc := []*substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "int32",
		PropertyName: "id",
		TableName:    "persons",
		Nullable:     false,
	}, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "string",
		PropertyName: "lastname",
		TableName:    "persons",
		Nullable:     true,
	}, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "string",
		PropertyName: "firstname",
		TableName:    "persons",
		Nullable:     true,
	}, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "string",
		PropertyName: "address",
		TableName:    "persons",
		Nullable:     true,
	}, &substance.ColumnDescription{
		DatabaseName: "postgres",
		PropertyType: "string",
		PropertyName: "city",
		TableName:    "persons",
		Nullable:     true,
	})
	columnDescResult, err := pgsqlProvider.DescribeTable("postgres", db, "persons")
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
	db, err := sql.Open("postgres", os.Getenv("SUBSTANCE_PGSQL"))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	pgsqlProvider := pgsql{}
	myColumnRel := []*substance.ColumnRelationship{}
	myColumnRel = append(myColumnRel, &substance.ColumnRelationship{
		TableName:           "orders",
		ColumnName:          "personid",
		ReferenceTableName:  "persons",
		ReferenceColumnName: "id",
	})
	columnRelResult, err := pgsqlProvider.TableRelationships("postgres", db, "orders")
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
	db, err := sql.Open("postgres", os.Getenv("SUBSTANCE_PGSQL"))
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()

	pgsqlProvider := pgsql{}
	myColumnConstraint := []*substance.ColumnConstraint{}
	myColumnConstraint = append(myColumnConstraint, &substance.ColumnConstraint{
		TableName:      "antiorders",
		ColumnName:     "antiorderid",
		ConstraintType: "p",
	}, &substance.ColumnConstraint{
		TableName:      "antiorders",
		ColumnName:     "personid",
		ConstraintType: "u",
	}, &substance.ColumnConstraint{
		TableName:      "antiorders",
		ColumnName:     "personid",
		ConstraintType: "f",
	})
	columnConstraintResult, err := pgsqlProvider.TableConstraints("postgres", db, "antiorders")
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
