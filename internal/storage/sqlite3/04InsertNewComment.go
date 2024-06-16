package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) InsertNewComment(post *models.Post, comment *models.Comment) ([]*models.Comment, error) {
	query := `INSERT INTO commentaries (post_id, user_id, commentray_content) VALUES (?, ?, ?)`

	res, err := s.db.Exec(query, post.Post_ID, comment.User.User_id, comment.Commentraie_Content)
	if err != nil {
		return nil, err
	}

	commentID, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	newComments, err := s.SelectComentByPostID(int(commentID))
	if err != nil {
		return nil, err
	}
	//*post = *newComments

	return newComments, nil
}
