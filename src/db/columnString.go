package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	if len(os.Args) < 6 {
		fmt.Println(fmt.Sprintf("Usage: go run %s [host] [user] [password] [database] [tableName]\n", os.Args[0]))
		os.Exit(1)
	}
	host := os.Args[1]
	user := os.Args[2]
	password := os.Args[3]
	database := os.Args[4]
	tableName := os.Args[5]

	connString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&connection+timeout=30", user, password, host, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Connect to Database failed:", err.Error())
		os.Exit(1)
	}

	query := fmt.Sprintf("SELECT TOP 1 * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columnNames, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Join(columnNames, ",")
	fmt.Println(result)
}
