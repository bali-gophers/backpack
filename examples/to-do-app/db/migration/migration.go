package migration

import (
	"github.com/gusdecool/backpack/examples/to-do-app/db/connector"
	"github.com/gusdecool/backpack/examples/to-do-app/db/model"
)

func Migrate() error {
	db, err := connector.Connect()

	if err != nil {
		return err
	}

	defer db.Close()
	db.AutoMigrate(&model.Task{})

	return nil
}