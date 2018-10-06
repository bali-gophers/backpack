package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	ID        int
	OrderNo   string
	Email     string
	SKU       string
	Quantity  int
	CreatedAt time.Time
}

func createMysql() (*sql.DB, error) {
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/order_workshop",
		"raka",
		"raka-is-not-used",
		"localhost")
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return sqlDB, nil
}

const (
	insertQuery = `
		INSERT INTO orders(
			order_number, 
			email, 
			sku,
			quantity,
			created_at) VALUES (?, ?, ?, ?, ?)
	`
)

func main() {
	sqlDB, err := createMysql()
	if err != nil {
		fmt.Println(err)
	}
	_, err = sqlDB.Exec(insertQuery,
		"O12",
		"raka@baligophers.lol",
		"BALI-1",
		12,
		time.Now())
	if err != nil {
		fmt.Println(err)
	}
}
