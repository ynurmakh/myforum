package businessrealiz

import (
	"errors"

	"forum/internal/models"
)

func (s *Service) GetOnlyMyPosts(thisUser *models.User) (*[]models.Post, error) {
	if thisUser == nil || thisUser.User_id < 1 {
		return nil, errors.New("err := wrong user_id")
	}

	return s.storage.SelectAllPostsByUserID(thisUser)
}
