package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Repo interface {
	Store(order Order) (int64, error)
	ResolveByOrderID(orderID int64) (Order, error)
}

type MysqlRepo struct {
	sqlDB *sql.DB
}

func NewMysqlRepo(sqlDB *sql.DB) Repo {
	return MysqlRepo{
		sqlDB: sqlDB,
	}
}

const (
	insertQuery             = "INSERT INTO `order` (order_number, full_name, email, total, created_at) VALUES (?,?,?,?,?)"
	insertItemQuery         = "INSERT INTO `item` (order_id, title, count, price) VALUES (?, ?, ?, ?)"
	selectByIDQuery         = "SELECT * FROM `order` WHERE order_id = ?"
	selectItemsByOrderQuery = "SELECT * FROM `item` WHERE order_id = ?"
)

func (repo MysqlRepo) Store(order Order) (int64, error) {
	stmt, err := repo.sqlDB.Prepare(insertQuery)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(order.OrderNumber, order.FullName, order.Email, order.Total, order.CreatedAt)
	if err != nil {
		return 0, err
	}
	lastInsertedID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	err = repo.storeItems(lastInsertedID, order.Items)
	if err != nil {
		return 0, err
	}
	return lastInsertedID, nil
}

func (repo MysqlRepo) storeItems(orderID int64, orderItems []Item) error {
	for _, orderItem := range orderItems {
		stmt, err := repo.sqlDB.Prepare(insertItemQuery)
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(orderID, orderItem.Title, orderItem.Count, orderItem.Price)
		if err != nil {
			return err
		}
	}
	return nil
}

func (repo MysqlRepo) ResolveByOrderID(orderID int64) (Order, error) {
	stmt, err := repo.sqlDB.Prepare(selectByIDQuery)
	if err != nil {
		return Order{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(orderID)

	var orderRow Order
	err = row.Scan(&orderRow.OrderID, &orderRow.OrderNumber, &orderRow.FullName, &orderRow.Email, &orderRow.Total, &orderRow.CreatedAt)
	if err != nil {
		return Order{}, err
	}
	orderItems, err := repo.resolveItems(orderID)
	if err != nil {
		return Order{}, err
	}
	orderRow.Items = orderItems
	return orderRow, nil
}

func (repo MysqlRepo) resolveItems(orderID int64) ([]Item, error) {
	var res []Item
	stmt, err := repo.sqlDB.Prepare(selectItemsByOrderQuery)
	if err != nil {
		return res, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(orderID)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var orderItem Item
		err = rows.Scan(&orderItem.ItemID, &orderItem.OrderID, &orderItem.Title, &orderItem.Count, &orderItem.Price)
		if err != nil {
			return res, err
		}
		res = append(res, orderItem)
	}
	return res, nil
}
