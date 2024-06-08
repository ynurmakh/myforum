package business

import (
	"forum/internal/models"
	"forum/internal/storage"
)

type Business interface {
	GetPosts() ([]storage.Post, error)
	GetPost(id int) (storage.Post, error)
	CreatePost(user_id, category_id int, title, content string) (int64, error)
	Login(email, pass string) (int, error)
	Exists(id int) (bool, error)
	GetUser(id int) (storage.User, error)
	CreateUser(name, email, pass string) (int64, error)
	GetUserByCookiesValues(sessionValue string) (*models.User, error)
	GetNewCookie() (string, error)
}
