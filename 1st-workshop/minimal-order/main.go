package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Order struct {
	OrderID   int
	OrderNo   string
	Total     int64
	CreatedAt time.Time
}

func main() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", "raka", "raka_pass", "127.0.0.1", "workshop_minimal_order")
	fmt.Printf("Connecting to %s\n", dataSourceName)
	sqlDB, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	// Insert new order
	fmt.Println()
	fmt.Println("Starting to insert new order ...")
	newOrder := Order{
		OrderNo:   "0714TYGD83",
		Total:     120000,
		CreatedAt: time.Now(),
	}

	stmt, err := sqlDB.Prepare("insert into orders (order_number, total, created_at) values (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(newOrder.OrderNo, newOrder.Total, newOrder.CreatedAt)
	if err != nil {
		panic(err)
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Printf("lastInsertID: %d\n", lastInsertID)

	// Select last inserted order
	fmt.Println()
	fmt.Println("Starting to select last inserted order ...")
	singleSelectStmt, err := sqlDB.Prepare("select * from orders where order_id = ?")
	if err != nil {
		panic(err)
	}
	defer singleSelectStmt.Close()

	row := singleSelectStmt.QueryRow(lastInsertID)
	var insertedOrder Order
	if err := row.Scan(&insertedOrder.OrderID, &insertedOrder.OrderNo, &insertedOrder.Total, &insertedOrder.CreatedAt); err != nil {
		panic(err)
	}
	fmt.Printf(
		"Last inserted order, orderNo: %s, total: %d, createdAt: %s\n",
		insertedOrder.OrderNo, insertedOrder.Total, insertedOrder.CreatedAt)

	// Select orders
	fmt.Println()
	selectStmt, err := sqlDB.Prepare("select * from orders")
	if err != nil {
		panic(err)
	}
	defer selectStmt.Close()
	rows, err := selectStmt.Query()
	if err != nil {
		panic(err)
	}

	var listOrder []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(&order.OrderID, &order.OrderNo, &order.Total, &order.CreatedAt)
		if err != nil {
			fmt.Printf("Error brow: %v\n", err)
			continue
		}
		listOrder = append(listOrder, order)
	}

	// Iterating order to render each item to standard output
	for _, orderRow := range listOrder {
		fmt.Printf(
			"Order #%d, orderNumber: %s, total %d, createdAt: %s\n",
			orderRow.OrderID, orderRow.OrderNo, orderRow.Total, orderRow.CreatedAt)
	}

	fmt.Println()
	fmt.Println("Udah gitu aja")
}
