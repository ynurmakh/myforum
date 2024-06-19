package businessrealiz

import (
	"github.com/google/uuid"
)

func (s *Service) CreateNewCookie() (string, error) {
	newUuid := uuid.New().String()
	_, err := s.storage.CheckTheCookie(newUuid, s.configs.CookieMaxAge)
	if err != nil {
		return "", err
	}
	return newUuid, nil
}
