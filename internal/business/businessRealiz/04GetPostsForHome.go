package businessrealiz

import (
	"forum/internal/models"
)

func (s *Service) GetPostsForHome(pageNum, onPage int, categories []int, thisUser *models.User) (*[]models.Post, error) {
	if pageNum < 1 || onPage < 1 {
		return &[]models.Post{}, nil
	}

	if len(categories) > 0 {
		return s.storage.FilteredSelectLastPostsByCount(onPage*(pageNum-1), onPage, thisUser, categories)
	}

	posts, err := s.storage.SelectLastPostsByCount(onPage*(pageNum-1), onPage, thisUser)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
