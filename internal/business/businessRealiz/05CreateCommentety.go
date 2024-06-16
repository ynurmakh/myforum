package businessrealiz

import (
	"forum/internal/models"
)

func (s *Service)CraeteCommentary(forComment *models.Post, comment *models.Comment) (error){
	err := s.storage.InsertNewComment(forComment, comment)
	if err != nil {
		return  err
	}
	return err

}