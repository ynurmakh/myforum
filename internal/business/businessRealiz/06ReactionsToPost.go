package businessrealiz

import (
	"errors"

	"forum/internal/models"
)

func (s *Service) ReactionsToPost(post *models.Post, thisUser *models.User, reactions int) error {
	// chec user
	if thisUser.User_id < 1 {
		return errors.New("err: wrong id on this User")
	}
	// chec post
	checkPost, err := s.storage.SelectPostByPostID(int(post.Post_ID), thisUser)
	if err != nil {
		return err
	}
	if checkPost == nil || checkPost.Post_ID < 1 {
		return errors.New("err: wrong id on post")
	}
	// chec react
	if reactions != 1 && reactions != -1 {
		return errors.New("err: wrong react")
	}

	return s.storage.ReactionsToPost(post, thisUser, reactions)
}
