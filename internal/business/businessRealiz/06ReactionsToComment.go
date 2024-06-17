package businessrealiz

import (
	"errors"

	"forum/internal/models"
)

func (s *Service) ReactionsToComment(commentId int, thisUser *models.User, reactions int) error {
	// chec user
	if thisUser.User_id < 1 {
		return errors.New("err: wrong id on this User")
	}
	// chec comment
	if commentId < 1 {
		return errors.New("err: wrong id to comment")
	}
	// chec react
	if reactions != 1 && reactions != -1 {
		return errors.New("err: wrong react")
	}

	return s.storage.ReactionsToComment(commentId, thisUser, reactions)
}
