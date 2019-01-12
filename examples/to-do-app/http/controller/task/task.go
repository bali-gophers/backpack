package task

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gusdecool/backpack/examples/to-do-app/db/model"
	repo "github.com/gusdecool/backpack/examples/to-do-app/db/repository/task"
	"github.com/gusdecool/backpack/examples/to-do-app/http/utility"
	"net/http"
	"strconv"
	"time"
)

func List(response http.ResponseWriter, request *http.Request) {
	tasks, err := repo.FindAll()

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	tasksByte, err := json.Marshal(tasks)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	utility.HandleSuccessResponse(tasksByte, response)
}

func Create(response http.ResponseWriter, request *http.Request) {
	taskModel, err := decode(request)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskModel.CreatedAt = time.Now()
	_, err = repo.Create(&taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskByte, err := json.Marshal(taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	utility.HandleSuccessResponse(taskByte, response)
}

func Update(response http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	taskId, err := strconv.Atoi(parameters["id"])

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskModel, err := repo.GetOneById(taskId)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	requestUpdateModel, err := decode(request)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskModel.Name = requestUpdateModel.Name
	_, err = repo.Update(&taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskByte, err := json.Marshal(taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	utility.HandleSuccessResponse(taskByte, response)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	taskId, err := strconv.Atoi(parameters["id"])

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	taskModel, err := repo.GetOneById(taskId)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	err = repo.Delete(&taskModel)

	if err != nil {
		utility.HandleErrorResponse(err, response)
		return
	}

	utility.HandleSuccessEmptyResponse(response)
}

func decode(request *http.Request) (model.Task, error) {
	var taskModel model.Task
	err := json.NewDecoder(request.Body).Decode(&taskModel)

	return taskModel, err
}