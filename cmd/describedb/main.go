package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ahmedalhulaibi/substance/substancegen"

	"github.com/ahmedalhulaibi/substance"
	_ "github.com/ahmedalhulaibi/substance/providers/mysqlsubstance"
	_ "github.com/ahmedalhulaibi/substance/providers/pgsqlsubstance"
	_ "github.com/ahmedalhulaibi/substance/providers/sqlitesubstance"
)

func main() {
	dbtype := flag.String("db", "", "Database driver name.\nSupported databases types:\n\t- mysql\n\t- postgres \n\t- sqlite3\n")
	connString := flag.String("cnstr", "", "Connection string to connect to database.")
	flag.Parse()
	results, err := substance.DescribeDatabase(*dbtype, *connString)
	if err != nil {
		log.Fatalf("Error describing db: %v", err)
	}
	if len(results) > 0 {
		fmt.Println("Database: ", results[0].DatabaseName)
	}
	var tables []string
	for _, result := range results {
		fmt.Printf("Table: %s\n", result.TableName)
		tables = append(tables, result.TableName)
	}
	fmt.Println("=====================")

	tableObjects := substancegen.GetObjectTypesFunc(*dbtype, *connString, tables)

	jsonB, err := json.Marshal(tableObjects)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonB))
		err = ioutil.WriteFile("substance-objects.json", jsonB, 0644)
	}
}
