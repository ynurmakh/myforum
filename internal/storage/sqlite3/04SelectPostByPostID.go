package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) SelectPostByPostID(PostID int) (*models.Post, error) {
	query := `SELECT posts.post_id, posts.post_title, posts.post_content, posts.created_time, users.user_id, users.user_lvl, users.user_email, users.user_nickname
	FROM posts
	JOIN users ON posts.user_id = users.user_id
	WHERE posts.post_id = ?`
	row := s.db.QueryRow(query, PostID)

	var post models.Post
	err := row.Scan(&post.Post_ID, &post.Post_Title, &post.Post_Content, &post.Created_Time, &post.User.User_id, &post.User.User_lvl, &post.User.User_email, &post.User.User_nickname)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
