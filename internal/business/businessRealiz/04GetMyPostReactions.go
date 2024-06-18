package businessrealiz

import "forum/internal/models"

func (s *Service) GetMyPostReactions(thisUser *models.User) (*[]models.Post, error) {
	return s.storage.SelectMyPostReactions(thisUser)
}
