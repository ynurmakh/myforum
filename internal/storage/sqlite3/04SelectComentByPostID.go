package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) SelectComentByPostID(PostId int) (*[]models.Comment, error) {
	query := `
	SELECT 
	commentaries.commentary_id,
	commentaries.commentray_content,
	commentaries.created_time,
	users.user_id,
	users.user_lvl,
	users.user_email,
	users.user_nickname
	FROM commentaries
	JOIN posts ON commentaries.post_id = posts.post_id
    JOIN users ON commentaries.user_id = users.user_id
    WHERE commentaries.post_id = ?
	ORDER BY commentaries.created_time DESC;`
	rows, err := s.db.Query(query, PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.Comment_Id,
			&comment.Commentraie_Content,
			&comment.Commentarie_Date,
			&comment.User.User_id,
			&comment.User.User_lvl,
			&comment.User.User_email,
			&comment.User.User_nickname)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &comments, nil
}
