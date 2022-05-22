package articles

import (
	"database/sql"
	"errors"
	"fmt"

	"example.com/module/models"
)

type Service interface {
	Get(Id uint16) (*models.Article, error)
	Create(title, anons, text string, UserId uint16) (*models.Article, error)
	List() ([]models.Article, error)
	Delete(Id uint16) error
	Update(Id uint16, title, anons, text string) (*models.Article, error)
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
		return nil, err
	}
	defer selectArticle.Close()

	post := models.Article{}

	if err := selectArticle.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText, &post.UserId); err != nil {
		return nil, err
	}

	selectUser, err := s.db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE  `id` = '%d'", post.UserId))
	if err != nil {
		return nil, err
	}
	defer selectUser.Close()
	if err := selectUser.Scan(&post.User.Id, &post.User.Name, &post.User.UserName, &post.User.Password, &post.User.Password); err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *service) Create(title, anons, text string, UserId uint16) (*models.Article, error) {
	if title == "" || anons == "" || text == "" {
		return nil, errors.New("all data must be filled")
	} else if UserId < 1 {
		return nil, errors.New("userid must be filled")
	}
	var id uint16
	insert, err := s.db.Query(fmt.Sprintf(`INSERT INTO 'articles' ('title', 'anons', 'full_text', 'user_id')
		VALUES ('%s', '%s', '%s', '%d') RETURNING id`,
		title, anons, text, UserId))

	if err != nil {
		return nil, err
	}

	defer insert.Close()

	err = insert.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.Article{
		Id:       id,
		Title:    title,
		Anons:    anons,
		FullText: text,
		UserId:   UserId,
	}, nil
}

func (s *service) List() ([]models.Article, error) {
	selectArticle, err := s.db.Query("SELECT * FROM `articles`")
	if err != nil {
		return nil, err
	}

	posts := []models.Article{}

	for selectArticle.Next() {
		var post models.Article
		err = selectArticle.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText, &post.UserId)
		if err != nil {
			return nil, err
		}
		selectUser, err := s.db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE  `id` = '%d'", post.UserId))
		if err != nil {
			return nil, err
		}
		defer selectUser.Close()
		if err := selectUser.Scan(&post.User.Id, &post.User.Name, &post.User.UserName, &post.User.Password, &post.User.Password); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil

}

func (s *service) Delete(Id uint16) error {
	_, err := s.Get(Id)
	if err != nil {
		return err
	}
	res, err := s.db.Query(fmt.Sprintf("DELETE FROM `articles` WHERE `id` = '%d'", Id))
	if err != nil {
		return err
	}
	defer res.Close()
	return nil
}

func (s *service) Update(Id uint16, title, anons, text string) (*models.Article, error) {
	_, err := s.Get(Id)
	if err != nil {
		return nil, err
	}
	res, err := s.db.Query(fmt.Sprintf(`UPDATE 'articles' SET  ('title', 'anons', 'full_text') 
	VALUES ('%s', '%s', '%s') WHERE 'id'='%d'`, title, anons, text, Id))
	if err != nil {
		return nil, err
	}
	defer res.Close()
	return s.Get(Id)
}
