package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func features(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Other features")
}

func handleRequest() {

	http.HandleFunc("/features", features)
	http.HandleFunc("/create", create)
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)

}

func main() {
	handleRequest()
}
