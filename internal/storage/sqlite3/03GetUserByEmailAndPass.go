package sqlite3

import "forum/internal/models"

func (s *Sqlite) GetUserByEmailAndPass(email string, hashed_password string) (*models.User, error) {
	return nil, nil
}
