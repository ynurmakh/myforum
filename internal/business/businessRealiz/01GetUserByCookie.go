package businessrealiz

import (
	"forum/internal/models"
)

func (s *Service) GetUserByCookie(sessionValue string) (*models.User, error) {
	user, err := s.storage.CheckTheCookie(sessionValue, s.configs.CookieMaxAge)
	return user, err
}
