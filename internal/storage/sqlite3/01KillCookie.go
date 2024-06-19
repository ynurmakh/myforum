package sqlite3

func (s *Sqlite) KillCookie(UUID string) (bool, error) {
	return true, s.sbrosCookies(UUID)
}
