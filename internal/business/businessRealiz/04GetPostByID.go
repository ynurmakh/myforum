package businessrealiz

import "forum/internal/models"

func (s *Service) GetPostByID(Post_ID int) (*models.Post, error) {
	return s.storage.SelectPostByPostID(Post_ID)
}
