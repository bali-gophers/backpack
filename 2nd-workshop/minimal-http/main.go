package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	fmt.Println("Starting server ...")

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello %s", "Gopher!")
	})

	http.HandleFunc("/helloHtml", func(w http.ResponseWriter, req *http.Request) {
		t := template.Must(template.ParseFiles("hello.html"))
		t.Execute(w, nil)
	})

	http.HandleFunc("/helloVars", func(w http.ResponseWriter, req *http.Request) {
		type PageVar struct {
			Title   string
			Message string
		}
		pageVar := PageVar{"Fundamental Go Workshop", "Send `hello gopher` to the world"}
		t := template.Must(template.ParseFiles("helloVars.html"))

		t.Execute(w, pageVar)
	})
	http.ListenAndServe(":8080", nil)
}
