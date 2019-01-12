package connector

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root@/go-todo-app?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		return nil, err
	}

	db.SingularTable(true)

	return db, err
}
