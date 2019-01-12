package task

import (
	"errors"
	"fmt"
	"github.com/gusdecool/backpack/examples/to-do-app/db/connector"
	"github.com/gusdecool/backpack/examples/to-do-app/db/model"
)

func FindAll() ([]model.Task, error) {
	db, err := connector.Connect()
	var tasks []model.Task

	if err != nil {
		return tasks, err
	}

	defer db.Close()
	db.Find(&tasks)

	return tasks, nil
}

func GetOneById(id int) (model.Task, error) {
	db, err := connector.Connect()
	var task model.Task

	if err != nil {
		return task, err
	}

	defer db.Close()

	if db.First(&task, id).RowsAffected == 0 {
		return task, errors.New(fmt.Sprintf("can't find task with id %d", id))
	}

	return task, nil
}

func Create(task *model.Task) (*model.Task, error) {
	db, err := connector.Connect()

	if err != nil {
		return task, err
	}

	defer db.Close()

	if db.NewRecord(task) == false {
		return task, errors.New("primary key not blank")
	}

	db.Create(&task)

	return task, nil
}

func Update(task *model.Task) (*model.Task, error) {
	db, err := connector.Connect()

	if err != nil {
		return task, err
	}

	defer db.Close()
	err = db.Save(task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

func Delete(task *model.Task) error {
	db, err := connector.Connect()

	if err != nil {
		return err
	}

	defer db.Close()

	return db.Delete(task).Error
}