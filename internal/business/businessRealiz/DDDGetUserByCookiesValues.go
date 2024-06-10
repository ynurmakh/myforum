package businessrealiz

import "forum/internal/models"

func (s *Service) GetUserByCookiesValues(sessionValue string) (*models.User, error) {
	if sessionValue == "123" {
		return &models.User{Email: "rus@mail.ru"}, nil
	}

	return nil, nil
}
