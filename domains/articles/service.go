package articles

import (
	"database/sql"
	"fmt"

	"example.com/module/models"
)

type Service interface {
	Get(Id uint16) (*models.Article, error)
	// Create(title, anons, text string, UserId uint16) error
	// List() ([]models.Article, error)
	// Delete(Id uint16) error
	// Update(Id uint16, title, anons, text string) (*models.Article, error)
}

type service struct {
	db *sql.DB
}

func NewService(db *sql.DB) Service {
	return &service{db}
}

func (s *service) Get(Id uint16) (*models.Article, error) {

	selectArticle, err := s.db.Query(fmt.Sprintf("SELECT * FROM `articles` WHERE  `id` = '%d'", Id))
	if err != nil {
		panic(err)
	}

	post := models.Article{}

	if err := selectArticle.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText, &post.UserId); err != nil {
		return nil, err
	}

	selectUser, err := s.db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE  `id` = '%d'", post.UserId))
	if err != nil {
		panic(err)
	}

	if err := selectUser.Scan(&post.User.Id, &post.User.Name, &post.User.UserName, &post.User.Password, &post.User.Password); err != nil {
		return nil, err
	}
	return &post, nil
}
