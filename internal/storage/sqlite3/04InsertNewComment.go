package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) InsertNewComment(post *models.Post, comment *models.Comment) (error) {
	query := `INSERT INTO commentaries (post_id, user_id, commentray_content) VALUES (?, ?, ?)`

	res, err := s.db.Exec(query, post.Post_ID, comment.User.User_id, comment.Commentraie_Content)
	if err != nil {
		return err
	}

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}

	newPost, err := s.SelectPostByPostID(int(post.Post_ID), &comment.User)
	if err != nil {
		return err
	}
	*post = *newPost

	return nil
}
