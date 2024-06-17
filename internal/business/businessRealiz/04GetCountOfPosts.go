package businessrealiz

func (s *Service) GetCountOfPosts() (int, error) {
	return s.storage.SelectCountOfPosts()
}
