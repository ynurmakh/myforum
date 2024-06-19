package sqlite3

func (s *Sqlite) TieCookieToUser(UserID int, UUID string, livetime int) (bool, error) {
	query := `
	UPDATE cookies SET user_id = ?
	WHERE cookie = ?	
	`
	_, err := s.db.Exec(query, UserID, UUID)
	if err != nil {
		return false, err
	}
	return true, s.touchCookie(UUID, livetime)
}
