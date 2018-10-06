package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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
		"%s:%s@tcp(%s:3306)/order_workshop?parseTime=true",
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

	router := mux.NewRouter()
	router.HandleFunc("/order/{id}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		orderID := vars["id"]
		row := sqlDB.QueryRow(
			"SELECT * FROM orders WHERE id=?",
			orderID)
		var order Order
		err = row.Scan(
			&order.ID,
			&order.OrderNo,
			&order.Email,
			&order.SKU,
			&order.Quantity,
			&order.CreatedAt)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		plainInfo := fmt.Sprintf("%s => %s", order.Email, order.OrderNo)
		w.Write([]byte(plainInfo))

	})
	http.ListenAndServe(":8080", router)

}
