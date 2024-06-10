package sqlite3

import (
	"fmt"
	"log"
	"time"

	"forum/internal/models"
)

func (s *Sqlite) CheckTheCookie(cookie string, expireTime int) (*models.User, error) {
	// selct

	row := s.db.QueryRow(`SELECT * FROM cookies WHERE cookie = ?`, cookie)

	var cookieRow struct {
		Cookie   string    `db:"cookie"`
		User_ID  int       `db:"user_id"`
		DeadTime time.Time `db:"deadTime"`
	}

	err := row.Scan(&cookieRow.Cookie, &cookieRow.User_ID, &cookieRow.DeadTime)
	if err != nil {
		return nil, err
	}

	fmt.Println(cookieRow)
	log.Fatalln(cookieRow)

	// insert

	return nil, nil
}
