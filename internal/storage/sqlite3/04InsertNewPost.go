package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) InsertNewPost(post *models.Post) error {
	queru := `INSERT INTO posts (user_id, post_title, post_content) VALUES (?, ?, ?)`

	res, err := s.db.Exec(queru, post.User.User_id, post.Post_Title, post.Post_Content)
	if err != nil {
		return err
	}

	postid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	newpost, err := s.SelectPostByPostID(int(postid))
	if err != nil {
		return err
	}
	*post = *newpost

	return nil
}
