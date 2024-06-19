package businessrealiz

import "forum/internal/models"

func (s *Service) LoginByEmailAndPass(email, pass, uuid string) (*models.User, error) {
	if err := checkEmailBeforRegistration(email); err != nil {
		return nil, err
	}
	if err := checkPasswordBeforRegistration(pass); err != nil {
		return nil, err
	}
	hashingPass(&pass, s.configs.PasswordHashingSecret)

	user, err := s.storage.GetUserByEmailAndPass(email, pass)
	if err != nil {
		return nil, err
	}

	_, err = s.storage.TieCookieToUser(int(user.User_id), uuid, s.configs.CookieMaxAge)
	if err != nil {
		return nil, err
	}
	return user, nil
}
