package sqlite3

import (
	"errors"
	"fmt"
	"time"
)

func (s *Sqlite) TieCookieToUser(UserID int, UUID string, DeadTimeSeconds int) (bool, error) {
	var exists struct {
		Cookie   string
		User_ID  int
		DeadTime time.Time
	}
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE cookie=?)", UUID).Scan(&exists.Cookie, &exists.User_ID, &exists.DeadTime)
	if err != nil {
		return false, fmt.Errorf("ошибка при проверке пользователя: %v", err)
	}
	if exists.Cookie == "" {
		return false, errors.New("пользователь не найден")
	}

	deadTime := time.Now().Add(time.Duration(DeadTimeSeconds) * time.Second)

	_, err = s.db.Exec("INSERT INTO cookies (cookie, user_id, deadTime) VALUES (?, ?, ?)", UUID, UserID, deadTime)
	if err != nil {
		return false, fmt.Errorf("ошибка при вставке записи в таблицу cookies: %v", err)
	}

	return true, nil
}
