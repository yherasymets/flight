package articles

import (
	"example.com/module/models"
	"gorm.io/gorm"
)

//https://gorm.io/docs/create.html

type gormService struct {
	db *gorm.DB
}

func NewGormService(db *gorm.DB) Service {
	return &gormService{db}
}

func (s *gormService) Get(Id uint16) (*models.Article, error) {
	var post models.Article
	if err := s.db.Preload("User").First(&post, Id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *gormService) Create(title, anons, text string, UserId uint16) (*models.Article, error) {
	post := models.Article{
		Title:    title,
		Anons:    anons,
		UserId:   UserId,
		FullText: text,
	}
	if err := s.db.Create(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *gormService) Update(Id uint16, title, anons, text string) (*models.Article, error) {
	post := models.Article{
		Id:       Id,
		Title:    title,
		Anons:    anons,
		FullText: text,
	}
	if err := s.db.Model(&models.Article).Update(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (s *gormService) Delete(Id uint16) error {
	post := models.Article{
		Id: Id,
	}
	if err := s.db.Delete(&post).Error; err != nil {
		return nil
	}
	return nil
}

// func (s *service) List() ([]models.Article, error) {
// 	selectArticle, err := s.db.Query("SELECT * FROM `articles`")
// 	if err != nil {
// 		return nil, err
// 	}

// 	posts := []models.Article{}

// 	for selectArticle.Next() {
// 		var post models.Article
// 		err = selectArticle.Scan(&post.Id, &post.Title, &post.Anons, &post.FullText, &post.UserId)
// 		if err != nil {
// 			return nil, err
// 		}
// 		selectUser, err := s.db.Query(fmt.Sprintf("SELECT * FROM `users` WHERE  `id` = '%d'", post.UserId))
// 		if err != nil {
// 			return nil, err
// 		}
// 		defer selectUser.Close()
// 		if err := selectUser.Scan(&post.User.Id, &post.User.Name, &post.User.UserName, &post.User.Password, &post.User.Password); err != nil {
// 			return nil, err
// 		}
// 		posts = append(posts, post)
// 	}

// 	return posts, nil

// }
