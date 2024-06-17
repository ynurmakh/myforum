package businessrealiz

import (
	"errors"
	"strings"

	"forum/internal/models"
)

func (s *Service) CreateNewUser(user models.User, password string) (*models.User, error) {
	// check email
	if !checkEmailBeforRegistration(user.User_email) {
		return nil, errors.New("Worng email")
	}

	// check nickname
	// check pass
	// hash pass pass

	return nil, nil
}

func checkEmailBeforRegistration(email string) bool {
	if strings.Index(email, "@") < 1 {
		return false
	}
	if len(email) > 50 {
		return false
	}
	return true
}
