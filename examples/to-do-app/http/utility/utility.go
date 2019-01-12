package utility

import "net/http"

func HandleSuccessResponse(message []byte, response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	response.Write(message)
}

func HandleSuccessEmptyResponse(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusNoContent)
}

func HandleErrorResponse(err error, response http.ResponseWriter) {
	response.WriteHeader(http.StatusBadRequest)
	response.Write([]byte(err.Error()))
}
