package businessrealiz

import "forum/internal/models"

func (s *Service) LoginByEmailAndPass(email, pass string) (*models.User, error) {
	if err := checkEmailBeforRegistration(email); err != nil {
		return nil, err
	}
	if err := checkPasswordBeforRegistration(pass); err != nil {
		return nil, err
	}
	hashingPass(&pass, s.configs.PasswordHashingSecret)

	return s.storage.GetUserByEmailAndPass(email, pass)
}
