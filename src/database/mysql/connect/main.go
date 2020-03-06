package main

import (
	"os"
	"fmt"
	"log"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

func main() {
	if len(os.Args) != 6 {
		fmt.Println("Usage: connect [account] [password] [host] [port] [database]")
		os.Exit(1)
	}
	account, password, host, port, database := os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5]
	connectString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", account, password, host, port, database)
	log.Println(connectString)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Connect Successful!")
	}
}