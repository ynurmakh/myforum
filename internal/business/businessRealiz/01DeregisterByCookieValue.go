package businessrealiz

func (s *Service) DeregisterByCookieValue(sessionValue string) (bool, error) {
	_, err := s.storage.KillCookie(sessionValue)
	return false, err
}
