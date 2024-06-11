package sqlite3

import (
	"errors"
	"fmt"
)

func (s *Sqlite) KillCookie(UUID string) (bool, error) {
	// Удаление cookie из базы данных
	result, err := s.db.Exec("DELETE FROM cookies WHERE cookie=?", UUID)
	if err != nil {
		return false, fmt.Errorf("ошибка при удалении cookie: %v", err)
	}

	// Проверка на количество затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("ошибка при получении количества затронутых строк: %v", err)
	}

	if rowsAffected == 0 {
		return false, errors.New("cookie с указанным UUID не найден")
	}

	return true, nil
}
