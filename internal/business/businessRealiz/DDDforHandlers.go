package businessrealiz

import (
	"forum/internal/storage"
)

var m = storage.MainModel{
	DB: storage.OpenDB(),
}

func (s *Service) GetPosts() ([]storage.Post, error) {
	return m.GetPosts()
}

func (s *Service) GetPost(id int) (storage.Post, error) {
	return m.GetPost(id)
}

func (s *Service) CreatePost(user_id, category_id int, title, content string) (int64, error) {
	return m.CreatePost(user_id, category_id, title, content)
}

func (s *Service) Login(email, pass string) (int, error) {
	return m.Login(email, pass)
}

func (s *Service) Exists(id int) (bool, error) {
	return m.Exists(id)
}

func (s *Service) GetUser(id int) (storage.User, error) {
	return m.GetUser(id)
}

func (s *Service) CreateUser(name, email, pass string) (int64, error) {
	return m.CreateUser(name, email, pass)
}
