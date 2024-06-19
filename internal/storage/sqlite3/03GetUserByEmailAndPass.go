package sqlite3

import "forum/internal/models"

func (s *Sqlite) GetUserByEmailAndPass(email string, hashed_password string) (*models.User, error) {
	query := `
	SELECT user_id,user_lvl,user_email,user_nickname FROM users WHERE (user_email, hashed_password) = (?,?)
	`
	row := s.db.QueryRow(query, email, hashed_password)
	user := &models.User{}
	err := row.Scan(&user.User_id, &user.User_lvl, &user.User_email, &user.User_nickname)
	if err != nil {
		return nil, err
	}

	return user, nil
}
