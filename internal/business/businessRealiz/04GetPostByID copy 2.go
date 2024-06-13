package businessrealiz

import "forum/internal/models"

func (s *Service) GetCategiries() (*[]models.Category, error) {
	return s.storage.GetAllCategiries()
}
