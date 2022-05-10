package articles

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"example.com/module/models"
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

type handlerService struct {
	*sql.DB
}

func NewHandler(db *sql.DB) *handlerService {
	return &handlerService{db}
}

func (s *handlerService) List(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	selectArticle, err := s.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	posts := []models.Article{}

	for selectArticle.Next() {
		var post models.Article
		err = selectArticle.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText, &post.UserId)

		posts = append(posts, post)

	}

	t.ExecuteTemplate(w, "index", posts)
}

func (s *handlerService) New(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("templates/new.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "new", nil)
}

func (s *handlerService) Features(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Other features")

}

func (s *handlerService) Create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	full_text := r.FormValue("full_text")

	if title == "" || anons == "" || full_text == "" {
		fmt.Fprintf(w, "Не все данные заполнены")

	} else {

		insert, err := s.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`, `user_id`) VALUES ('%s', '%s', '%s', '%d')",
			title, anons, full_text, 2))

		if err != nil {
			panic(err)
		}

		defer insert.Close()

		http.Redirect(w, r, "/articles/", http.StatusSeeOther)
	}
}

func (s *handlerService) Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	res, err := s.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE `id` = '%s'", vars["id"]))

	if err != nil {
		panic(err)
	}

	showPost := models.Article{}

	for res.Next() {
		var post models.Article
		err = res.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText, &post.UserId)

		showPost = post

	}

	t.ExecuteTemplate(w, "show", showPost)

}

func (s *handlerService) Delete(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	res, err := s.Query(fmt.Sprintf("DELETE FROM `articles` WHERE `id` = '%s'", vars["id"]))

	if err != nil {
		panic(err)
	}

	defer res.Close()

	http.Redirect(w, r, "/articles/", http.StatusSeeOther)

}
