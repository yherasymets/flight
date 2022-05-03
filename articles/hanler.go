package articles

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Handler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	New(w http.ResponseWriter, r *http.Request)
	Show(w http.ResponseWriter, r *http.Request)
	Features(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type service struct {
	*sql.DB
}

func NewHandler(db *sql.DB) *service {
	return &service{db}
}

func (s *service) List(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	res, err := s.Query("SELECT * FROM `articles`")

	if err != nil {
		panic(err)
	}

	posts := []Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		posts = append(posts, post)

	}

	t.ExecuteTemplate(w, "index", posts)
}

func (s *service) New(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}

func (s *service) Features(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Other features")

}

func (s *service) Create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {

		insert, err := s.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))

		if err != nil {
			panic(err)
		}

		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func (s *service) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	res, err := s.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))

	if err != nil {
		panic(err)
	}

	showPost := Article{}

	for res.Next() {
		var post Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText)

		showPost = post

	}

	t.ExecuteTemplate(w, "show", showPost)

}

func (s *service) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, err := s.Query(fmt.Sprintf("DELETE FROM `articles` WHERE `id` = '%s'", vars["id"]))

	if err != nil {
		panic(err)
	}

	defer res.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
