package sqlitesubstance

import (
	"database/sql"
	"os"
	"reflect"
	"sort"
	"testing"

	/*blank import to load sqlite3 driver*/

	"github.com/ahmedalhulaibi/substance"
	_ "github.com/mattn/go-sqlite3"
)

func TestGetCurrDbName(t *testing.T) {
	os.Remove("./test.db")
	// db, err := sql.Open("sqlite3", "./test1.db")
	// if err != nil {
	// 	t.Errorf(err.Error())
	// }
	// defer db.Close()
	// sqlStmt := `
	// create table foo (id integer not null primary key, name text);
	// delete from foo;
	// `
	// _, err = db.Exec(sqlStmt)
	// if err != nil {
	// 	t.Errorf("%q: %s\n", err, sqlStmt)
	// 	return
	// }

	sqliteProvider := sqlite{}
	nameExpected := "test.db"
	nameResult, err := sqliteProvider.GetCurrentDatabaseNameFunc("sqlite3", "./test.db")
	t.Logf("Expected '%s' as database name but got '%s'.", nameExpected, nameResult)
	if nameResult != nameExpected {
		t.Errorf("Expected '%s' as database name but got '%s'.", nameExpected, nameResult)
	}
	if err != nil {
		t.Error(err)
	}
}

func TestDescribeDb(t *testing.T) {
	os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE Persons (ID int PRIMARY KEY,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));
	CREATE TABLE Orders (OrderID int UNIQUE NOT NULL,OrderNumber int NOT NULL,PersonID int DEFAULT NULL,PRIMARY KEY (OrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	CREATE TABLE AntiOrders (AntiOrderID int UNIQUE NOT NULL,AntiOrderNumber int NOT NULL,PersonID int UNIQUE DEFAULT NULL,PRIMARY KEY (AntiOrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		t.Errorf("%q: %s\n", err, sqlStmt)
		return
	}

	sqliteProvider := sqlite{}
	myColumnDesc := []substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "Table",
		PropertyName: "Persons",
		TableName:    "Persons",
	}, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "Table",
		PropertyName: "Orders",
		TableName:    "Orders",
	}, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "Table",
		PropertyName: "AntiOrders",
		TableName:    "AntiOrders",
	})

	columnDescResult, err := sqliteProvider.DescribeDatabaseFunc("sqlite3", "./test.db")
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
	os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE Persons (ID int PRIMARY KEY NOT NULL,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));
	CREATE TABLE Orders (OrderID int UNIQUE NOT NULL,OrderNumber int NOT NULL,PersonID int DEFAULT NULL,PRIMARY KEY (OrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	CREATE TABLE AntiOrders (AntiOrderID int UNIQUE NOT NULL,AntiOrderNumber int NOT NULL,PersonID int UNIQUE DEFAULT NULL,PRIMARY KEY (AntiOrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		t.Errorf("%q: %s\n", err, sqlStmt)
		return
	}

	sqliteProvider := sqlite{}
	myColumnDesc := []substance.ColumnDescription{}
	myColumnDesc = append(myColumnDesc, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "int32",
		PropertyName: "ID",
		TableName:    "Persons",
		Nullable:     false,
	}, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "string",
		PropertyName: "LastName",
		TableName:    "Persons",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "string",
		PropertyName: "FirstName",
		TableName:    "Persons",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "string",
		PropertyName: "Address",
		TableName:    "Persons",
		Nullable:     true,
	}, substance.ColumnDescription{
		DatabaseName: "test.db",
		PropertyType: "string",
		PropertyName: "City",
		TableName:    "Persons",
		Nullable:     true,
	})
	columnDescResult, err := sqliteProvider.DescribeTableFunc("sqlite3", "./test.db", "Persons")
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
	os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE Persons (ID int PRIMARY KEY NOT NULL,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));
	CREATE TABLE Orders (OrderID int UNIQUE NOT NULL,OrderNumber int NOT NULL,PersonID int DEFAULT NULL,PRIMARY KEY (OrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	CREATE TABLE AntiOrders (AntiOrderID int UNIQUE NOT NULL,AntiOrderNumber int NOT NULL,PersonID int UNIQUE DEFAULT NULL,PRIMARY KEY (AntiOrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		t.Errorf("%q: %s\n", err, sqlStmt)
		return
	}

	sqliteProvider := sqlite{}
	myColumnRel := []substance.ColumnRelationship{}
	myColumnRel = append(myColumnRel, substance.ColumnRelationship{
		TableName:           "Orders",
		ColumnName:          "PersonID",
		ReferenceTableName:  "Persons",
		ReferenceColumnName: "ID",
	})
	columnRelResult, err := sqliteProvider.DescribeTableRelationshipFunc("sqlite3", "./test.db", "Orders")
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
	os.Remove("./test.db")

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer db.Close()
	sqlStmt := `
	CREATE TABLE Persons (ID int PRIMARY KEY NOT NULL,LastName varchar(255),FirstName varchar(255),Address varchar(255),City varchar(255));
	CREATE TABLE Orders (OrderID int UNIQUE NOT NULL,OrderNumber int NOT NULL,PersonID int DEFAULT NULL,PRIMARY KEY (OrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	CREATE TABLE AntiOrders (AntiOrderID int UNIQUE NOT NULL,AntiOrderNumber int NOT NULL,PersonID int UNIQUE DEFAULT NULL,PRIMARY KEY (AntiOrderID),FOREIGN KEY (PersonID) REFERENCES Persons(ID));
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		t.Errorf("%q: %s\n", err, sqlStmt)
		return
	}

	sqliteProvider := sqlite{}
	myColumnConstraint := []substance.ColumnConstraint{}
	myColumnConstraint = append(myColumnConstraint, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "AntiOrderID",
		ConstraintType: "pk",
	}, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "AntiOrderID",
		ConstraintType: "u",
	}, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "PersonID",
		ConstraintType: "f",
	}, substance.ColumnConstraint{
		TableName:      "AntiOrders",
		ColumnName:     "PersonID",
		ConstraintType: "u",
	})
	columnConstraintResult, err := sqliteProvider.DescribeTableConstraintsFunc("sqlite3", "./test.db", "AntiOrders")
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
