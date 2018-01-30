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
	return columnDesc, nil
}

/*DescribeTable returns columns in database*/
func (t testsql) DescribeTableFunc(dbType string, connectionString string, tableName string) ([]substance.ColumnDescription, error) {
	columnDesc := []substance.ColumnDescription{}
	return columnDesc, nil
}

/*DescribeTableRelationship returns all foreign column references in database table*/
func (t testsql) DescribeTableRelationshipFunc(dbType string, connectionString string, tableName string) ([]substance.ColumnRelationship, error) {
	columnDesc := []substance.ColumnRelationship{}
	return columnDesc, nil
}

func (t testsql) DescribeTableConstraintsFunc(dbType string, connectionString string, tableName string) ([]substance.ColumnConstraint, error) {
	columnDesc := []substance.ColumnConstraint{}
	return columnDesc, nil
}

func (t testsql) GetGoDataType(sqlType string) (string, error) {
	return "", nil
}
