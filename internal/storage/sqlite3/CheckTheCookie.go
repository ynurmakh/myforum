package sqlite3

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"forum/internal/models"
)

func (s *Sqlite) CheckTheCookie(cookie string, expireTime int) (*models.User, error) {
	var userID int
	var deadTime time.Time

	// Проверка наличия cookie и получения информации о сроке действия
	err := s.db.QueryRow("SELECT user_id, deadTime FROM cookies WHERE cookie=?", cookie).Scan(&userID, &deadTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("cookie не найден")
		}
		return nil, fmt.Errorf("ошибка при проверке cookie: %v", err)
	}

	// Проверка, не истек ли срок действия cookie
	if time.Now().After(deadTime) {
		return nil, errors.New("срок действия cookie истек")
	}

	// Получение информации о пользователе
	var user models.User
	err = s.db.QueryRow("SELECT id, email FROM users WHERE id=?", userID).Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении информации о пользователе: %v", err)
	}

	return &user, nil
}
