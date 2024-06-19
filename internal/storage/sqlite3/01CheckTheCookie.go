package sqlite3

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
	"time"
)

func (s *Sqlite) CheckTheCookie(cookie string, livetime int) (*models.User, error) {
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
		// Если такой строчки в бд нет то создаем такую стрчку с нуль юзером
		err := s.insertCookie(cookie, livetime)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	now := time.Now().UTC()
	exp := tStruct.lastcall.Add(time.Duration(tStruct.livetime) * time.Second)
	// Если время жизни не просрочено вернуть юзера, продлить
	if now.Before(exp) {
		err := s.touchCookie(cookie, livetime)
		if err != nil {
			return nil, err
		}

		// Если юзер под этим куки закреплен то вернуть этого юзера, если нет то вернуть нуль
		if tStruct.user_id.Valid {
			user, err := s.getUserByUserId(tStruct.user_id.Int64)
			if err != nil {
				return nil, nil
			}
			fmt.Println(user, "<<<")
			return user, nil

		}
		return nil, nil
	}

	// если просрочено сбросить юзера вернуть нуль, продлить
	err = s.sbrosCookies(cookie)
	if err != nil {
		return nil, err
	}
	// продение кукиса
	err = s.touchCookie(cookie, livetime)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *Sqlite) touchCookie(cookie string, livetime int) error {
	query := `
	UPDATE cookies
	SET last_call = datetime(CURRENT_TIMESTAMP) , livetime = (?)
	WHERE cookies.cookie = (?)
	`
	_, err := s.db.Exec(query, livetime, cookie)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) insertCookie(cookie string, livetime int) error {
	query := `INSERT INTO cookies (cookie, livetime) VALUES (? ,?)`
	_, err := s.db.Exec(query, cookie, livetime)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) getUserByUserId(id int64) (*models.User, error) {
	query := `
	SELECT user_lvl, user_email,user_nickname
	FROM users 
	WHERE users.user_id = (?)
	`
	row := s.db.QueryRow(query, id)
	user := &models.User{}
	err := row.Scan(&user.User_lvl, &user.User_email, &user.User_nickname)
	if err != nil {
		return nil, err
	}
	user.User_id = id
	return user, nil
}

func (s *Sqlite) sbrosCookies(cookies string) error {
	query := `UPDATE cookies
	SET user_id = NULL
	WHERE cookie = (?)`
	_, err := s.db.Exec(query, cookies)
	if err != nil {
		return err
	}
	return nil
}
