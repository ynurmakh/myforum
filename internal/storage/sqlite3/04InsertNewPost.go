package sqlite3

import (
	"database/sql"
	"strconv"
	"strings"

	"forum/internal/models"
)

func (s *Sqlite) InsertNewPost(post *models.Post, cats []int) error {
	var (
		res sql.Result
		err error
	)
	if len(cats) >= 1 {
		categStrArr := make([]string, 0, len(cats))
		for _, cat := range cats {
			categStrArr = append(categStrArr, strconv.Itoa(cat))
		}

		categStr := "[" + strings.Join(categStrArr, ", ") + "]"

		queru := `INSERT INTO posts (user_id, post_title, post_content, categories_id) VALUES (?, ?, ?, ?)`
		res, err = s.db.Exec(queru, post.User.User_id, post.Post_Title, post.Post_Content, categStr)
	}
	if len(cats) < 1 {
		queru := `INSERT INTO posts (user_id, post_title, post_content) VALUES (?, ?, ?)`
		res, err = s.db.Exec(queru, post.User.User_id, post.Post_Title, post.Post_Content)
	}
	if err != nil {
		return err
	}

	postid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	newpost, err := s.SelectPostByPostID(int(postid), &post.User)
	if err != nil {
		return err
	}
	*post = *newpost

	return nil
}
