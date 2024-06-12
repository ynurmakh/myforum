package sqlite3

import (
	"fmt"
	"forum/internal/models"
)

func (s *Sqlite) InsertNewPost(post *models.Post) error {
	queru := `INSERT INTO posts (user_id, post_title, post_content) VALUES (?, ?, ?)`

	res, err := s.db.Exec(queru, post.User.User_id, post.Post_Title, post.Post_Content)
	if err != nil {
		return err
	}

	fmt.Println(err)
	fmt.Println(res)

	return nil
}
