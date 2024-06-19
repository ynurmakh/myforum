package businessrealiz

import "forum/internal/models"

func (s *Service) GetUserByCookie(sessionValue string) (*models.User, error) {
	return s.storage.CheckTheCookie(sessionValue, s.configs.CookieMaxAge)
}
