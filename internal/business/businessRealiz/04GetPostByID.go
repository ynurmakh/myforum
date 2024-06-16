package businessrealiz

import "forum/internal/models"

func (s *Service) GetPostByID(Post_ID int, thisUser *models.User) (*models.Post, error) {
	return s.storage.SelectPostByPostID(Post_ID, thisUser)
}
