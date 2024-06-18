package businessrealiz

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"

	"forum/internal/models"
)

func (s *Service) CreateNewUser(user *models.User, password string) (*models.User, error) {
	// check email
	if err := checkEmailBeforRegistration(user.User_email); err != nil {
		return nil, err
	}

	// check nickname
	if err := checkNickNameBeforRegistration(user.User_nickname); err != nil {
		return nil, err
	}
	// check pass
	if err := checkPasswordBeforRegistration(password); err != nil {
		return nil, err
	}
	// hash pass pass
	hashingPass(&password, s.configs.PasswordHashingSecret)

	return s.storage.InsertNewUser(user.User_email, user.User_nickname, password)
}

func checkEmailBeforRegistration(email string) error {
	if strings.Index(email, "@") < 1 {
		return errors.New("Worng email")
	}
	if len(email) > 50 {
		return errors.New("Too long email")
	}
	return nil
}

func checkNickNameBeforRegistration(nickname string) error {
	if len(nickname) > 50 {
		return errors.New("Nick Name Too long")
	}
	return nil
}

func checkPasswordBeforRegistration(password string) error {
	if len(password) < 8 {
		return errors.New("Password too short")
	}

	return nil
}

func hashingPass(usersPassword *string, secretForHashing string) {
	h := hmac.New(sha256.New, []byte(secretForHashing))
	h.Write([]byte(*usersPassword))
	hashedPass := hex.EncodeToString(h.Sum(nil))
	*usersPassword = hashedPass
}
