package main

import (
	"database/sql"
	"net/http"

	"example.com/module/domains/articles"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("mysql", "yara24:password@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	articleHandler := articles.NewHandler(db)

	r := mux.NewRouter()

	articleResurce := r.PathPrefix("/articles").Subrouter()

	articleResurce.HandleFunc("/", articleHandler.List).Methods("GET")
	articleResurce.HandleFunc("/{id:[0-9]+}", articleHandler.Show).Methods("GET")
	articleResurce.HandleFunc("/new", articleHandler.New).Methods("GET")
	articleResurce.HandleFunc("/create", articleHandler.Create).Methods("POST")
	articleResurce.HandleFunc("/{id:[0-9]+}/delete", articleHandler.Delete).Methods("POST")

	http.Handle("/", articleResurce)

	http.HandleFunc("/features", articleHandler.Features)

	http.ListenAndServe(":8080", nil)

}
