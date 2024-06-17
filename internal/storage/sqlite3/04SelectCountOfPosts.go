package sqlite3

func (s *Sqlite) SelectCountOfPosts() (int, error) {
	row := s.db.QueryRow(`SELECT COUNT(*) FROM posts`)
	var res int64
	err := row.Scan(&res)
	if err != nil {
		return 0, err
	}

	return int(res), nil
}
