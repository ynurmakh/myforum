package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) InsertNewUser(email, nickname, password string) (*models.User, error) {
	query := `
	INSERT INTO users (user_email, user_nickname,hashed_password) VALUES (?,?,?)
	`
	res, err := s.db.Exec(query, email, nickname, password)
	if err != nil {
		if err.Error() == "UNIQUE constraint failed: users.user_email" {
			return nil, models.CreateUser_NotUniqEmail
		}
		if err.Error() == "UNIQUE constraint failed: users.user_nickname" {
			return nil, models.CreateUser_NotUniqNickName
		}
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &models.User{
		User_id:       id,
		User_lvl:      1,
		User_email:    email,
		User_nickname: nickname,
	}, nil
}
