package businessrealiz

import "forum/internal/models"

func (s *Service) CreatePost(post *models.Post) error {
	return s.storage.InsertNewPost(post)
}
