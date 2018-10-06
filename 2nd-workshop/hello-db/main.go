package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// How to use:
// - go get github.com/go-sql-driver/mysql
// - go run main.go

func main() {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/hello_db",
		"raka",
		"raka-is-not-used",
		"localhost")
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
	}
	if err := sqlDB.Ping(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Pong!")
	}
}
