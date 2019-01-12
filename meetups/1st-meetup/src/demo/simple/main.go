package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Starting server ...")
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		w.Write([]byte(fmt.Sprintf("Hello World! %s", url)))
	})
	http.HandleFunc("/gophers", func(w http.ResponseWriter, req *http.Request) {
		list := []string{
			"Ketut",
			"Nyoman",
		}
		w.Write([]byte(fmt.Sprintf("Gophers: %s", strings.Join(list, ", "))))
	})
	http.ListenAndServe(":8080", nil)
}
