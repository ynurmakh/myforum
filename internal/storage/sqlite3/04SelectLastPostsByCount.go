package sqlite3

import (
	"fmt"
	"forum/internal/models"
)

func (s *Sqlite) SelectLastPostsByCount(start, onPage int) (*[]models.Post, error) {
	zapros := fmt.Sprintf(`SELECT posts.*, users.user_email, users.user_nickname FROM posts JOIN users ON posts.user_id = users.user_id ORDER BY posts.created_time LIMIT %d OFFSET %d`, onPage, start)

	rows, err := s.db.Query(zapros)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		err := rows.Scan(&post.Post_ID, &post.User.User_id, &post.Post_Title, &post.Post_Content, &post.Created_Time, &post.User.User_email, &post.User.User_nickname)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}
