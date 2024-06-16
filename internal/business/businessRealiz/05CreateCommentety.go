package businessrealiz

import (
	"forum/internal/models"
)

func (s *Service)CraeteCommentary(forComment *models.Post, comment *models.Comment) (*[]models.Comment, error){
	sl, err := s.storage.InsertNewComment(forComment, comment)
	if err != nil {
		return nil, err
	}
	return sl, err

}