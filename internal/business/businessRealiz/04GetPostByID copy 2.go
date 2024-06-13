package businessrealiz

import "forum/internal/models"

func (s *Service) GetCategiries() (*[]models.Categories, error) {
	return s.storage.GetAllCategiries()
}
