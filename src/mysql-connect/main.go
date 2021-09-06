package main

import (
    "database/sql"
	"os"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Args[1]
	password := os.Args[2]
	host := os.Args[3]
	dbname := os.Args[4]
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4", user, password, host, dbname))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	var version string
	err = db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Conntect to:", version)
}