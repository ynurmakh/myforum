package businessrealiz

import (
	"forum/internal/models"
)

func (s *Service) CreatePost(post *models.Post, cats []int) error {
	return s.storage.InsertNewPost(post, cats)
}
