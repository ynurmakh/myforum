package businessrealiz

import (
	"fmt"

	"forum/internal/models"
)

func (s *Service) GetPostsForHome(pageNum, onPage int, categories []int) (*[]models.Post, error) {
	if pageNum < 1 || onPage < 1 {
		return &[]models.Post{}, nil
	}

	if len(categories) > 0 {
		allcategories, err := s.storage.GetAllCategiries()
		if err != nil {
			return nil, err
		}
		_ = allcategories
		// to do филтрация по категриям и возврать постов с учетом пагинаций и вмещаемости
		return &[]models.Post{
			{
				Post_Title: "with categories not realized yet",
			},
		}, nil
	}

	fmt.Println(onPage*(pageNum-1), onPage)
	posts, err := s.storage.SelectLastPostsByCount(onPage*(pageNum-1), onPage)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
