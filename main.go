package main

import (
	"database/sql"
	"net/http"

	"example.com/module/articles"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	db, err := sql.Open("mysql", "yara24:password@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}

	articleHandler := articles.NewHandler(db)

	articleResurce := mux.NewRouter()

	articleResurce.HandleFunc("/", articleHandler.List).Methods("GET")
	articleResurce.HandleFunc("/post/{id:[0-9]+}", articleHandler.Show).Methods("GET")
	articleResurce.HandleFunc("/create", articleHandler.New).Methods("GET")
	articleResurce.HandleFunc("/save_article", articleHandler.Create).Methods("POST")
	articleResurce.HandleFunc("/delete_article/{id:[0-9]+}", articleHandler.Delete).Methods("GET")

	http.Handle("/", articleResurce)

	http.HandleFunc("/features", articleHandler.Features)

	http.ListenAndServe(":8080", nil)

}
