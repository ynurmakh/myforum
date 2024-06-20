package businessrealiz

import (
	"forum/internal/models"
)

func (s *Service) GetPostsForHome(pageNum, onPage int, categories []int, thisUser *models.User) (*[]models.Post, int, error) {
	if pageNum < 1 || onPage < 1 {
		return &[]models.Post{}, 0, nil
	}

	if len(categories) > 0 {
		return s.storage.FilteredSelectLastPostsByCount(onPage*(pageNum-1), onPage, thisUser, categories)
	}

	return s.storage.SelectLastPostsByCount(onPage*(pageNum-1), onPage, thisUser)
}
