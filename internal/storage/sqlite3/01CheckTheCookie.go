package sqlite3

import (
	"database/sql"
	"fmt"
	"time"

	"forum/internal/models"
)

func (s *Sqlite) CheckTheCookie(cookie string, expireTime int) (*models.User, error) {
	query := `
	SELECT 
    cookie, 
    user_id, 
    liveTime, 
    last_call
FROM 
    cookies
WHERE 
    cookie = ?;

	`
	row := s.db.QueryRow(query, cookie)

	var tStruct struct {
		cookie   string
		user_id  sql.NullInt64
		livetime int64
		lastcall time.Time
	}
	err := row.Scan(
		&tStruct.cookie,
		&tStruct.user_id,
		&tStruct.livetime,
		&tStruct.lastcall,
	)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	exp := tStruct.lastcall.Add(time.Duration(tStruct.livetime) * time.Second)

	fmt.Printf("now: %v\nexp: %v \n\n", now, exp)

	if now.Before(exp) {
		fmt.Println("Yes, продлеваю")
		s.touchCookie(cookie, expireTime)
	} else {
		fmt.Println("no")
	}

	return nil, nil
}

func (s *Sqlite) touchCookie(cookie string, livetime int) error {
	query := `
	UPDATE cookies
	SET last_call = datetime(CURRENT_TIMESTAMP) , livetime = (?)
	WHERE cookies.cookie = (?)
	`
	res, err := s.db.Exec(query, livetime, cookie)
	fmt.Println(res, err)
	return nil
}
