package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{firstName}", func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		firstName := vars["firstName"]
		w.Write([]byte(fmt.Sprintf("Hello %s", firstName)))
	})
	http.ListenAndServe(":8080", router)
}
