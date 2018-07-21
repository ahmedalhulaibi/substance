package sqlitesubstance

/*GetCurrentDatabaseNameQuery used in GetCurrentDatabaseNamefunc*/
var GetCurrentDatabaseNameQuery = `PRAGMA database_list;`

/*DescribeDatabaseQuery used in DescribeDatabaseFunc*/
var DescribeDatabaseQuery = `select name from sqlite_master where name not like 'sqlite%';`

/*DescribeTableQuery used in DescribeTableFunc*/
var DescribeTableQuery = `PRAGMA table_info( $1 );`

/*DescribeTableRelationshipQuery used in DescribeTableRelationshipFunc*/
var DescribeTableRelationshipQuery = `PRAGMA foreign_key_list( $1 );`

/*SQLLiteIndexList used in DescribeTableConstraintsFunc*/
var SQLLiteIndexList = `PRAGMA index_list( $1 );`

/*SQLLiteIndexInfo used in DescribeTableConstraintsFunc*/
var SQLLiteIndexInfo = `PRAGMA index_info( $1 );`
