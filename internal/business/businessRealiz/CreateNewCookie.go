package businessrealiz

import (
	"fmt"

	"github.com/google/uuid"
)

func (s *Service) CreateNewCookie() (string, error) {
	newUuid := uuid.New().String()

	fmt.Println(s.storage.CheckTheCookie(newUuid, s.configs.CookieMaxAge))
	return newUuid, nil
}
