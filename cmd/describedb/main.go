package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/ahmedalhulaibi/substance/substancegen"

	"github.com/ahmedalhulaibi/substance"
	_ "github.com/ahmedalhulaibi/substance/providers/mysqlsubstance"
	_ "github.com/ahmedalhulaibi/substance/providers/pgsqlsubstance"
	_ "github.com/ahmedalhulaibi/substance/providers/sqlitesubstance"
	_ "github.com/lib/pq"
)

func main() {
	dbtype := flag.String("db", "", "Database driver name.\nSupported databases types:\n\t- mysql\n\t- postgres \n\t- sqlite3\n")
	connString := flag.String("cnstr", "", "Connection string to connect to database.")

	flag.Parse()

	db, err := sql.Open(*dbtype, *connString)
	if err != nil {
		log.Fatalf("Error opening database: %s\n", err.Error())
	}

	done := make(chan bool)

	go func() {
		results, err := substance.DescribeDatabase(*dbtype, db)
		fmt.Fprintf(os.Stderr, "Results: %+v\n", results)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error describing db: %v", err)
			done <- false
		}
		if len(results) > 0 {
			fmt.Fprintf(os.Stderr, "Database: %v", results[0].DatabaseName)
		}
		var tables []string
		for idx := range results {
			fmt.Fprintf(os.Stderr, "Table: %s\n", results[idx].TableName)
			tables = append(tables, results[idx].TableName)
		}
		fmt.Fprintln(os.Stderr, "=====================")

		tableObjects := substancegen.GetObjectTypesFunc(*dbtype, db, tables)

		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "\t")
		if err := enc.Encode(tableObjects); err != nil {
			fmt.Fprintf(os.Stderr, "Error marhsalling to json: %v", err)
			done <- false
		}
		done <- true
	}()

	ticker := time.NewTicker(500 * time.Millisecond)
	PrintMemUsage()

Loop:
	for {
		select {
		case <-ticker.C:
			PrintMemUsage()
		case exitCode := <-done:
			if !exitCode {
				os.Exit(1)
			}
			os.Exit(0)
			break Loop
		}

	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Fprintf(os.Stderr, "Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Fprintf(os.Stderr, "\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Fprintf(os.Stderr, "\tSys = %v MiB", bToMb(m.Sys))
	fmt.Fprintf(os.Stderr, "\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
