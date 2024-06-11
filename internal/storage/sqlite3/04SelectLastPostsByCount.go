package sqlite3

import (
	"fmt"

	"forum/internal/models"
)

func (s *Sqlite) SelectLastPostsByCount(start, onPage int) (*[]models.Post, error) {
	zapros := fmt.Sprintf(`SELECT * FROM posts ORDER BY created_time LIMIT %d OFFSET %d`, onPage, start)

	rows, err := s.db.Query(zapros)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		err := rows.Scan(&post.Post_ID, &post.User_ID, &post.Post_Title, &post.Post_Content, &post.Created_Time)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return &posts, nil
}
