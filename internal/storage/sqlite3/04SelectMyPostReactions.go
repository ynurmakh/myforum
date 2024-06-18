package sqlite3

import (
	"fmt"

	"forum/internal/models"
)

func (s *Sqlite) SelectMyPostReactions(thisUser *models.User) (*[]models.Post, error) {
	query := fmt.Sprintf(`
	SELECT 
	posts.post_id,
	posts.post_title,
	posts.post_content,
	posts.created_time,
	posts.categories_id,
	posts.liked_ids,
	posts.disliked_ids,
	users.user_id,
	users.user_lvl,
	users.user_email,
	users.user_nickname
	FROM posts, json_each(posts.liked_ids)
	JOIN users on posts.user_id = users.user_id
	WHERE value in (%v)
	UNION
	SELECT 
	posts.post_id,
	posts.post_title,
	posts.post_content,
	posts.created_time,
	posts.categories_id,
	posts.liked_ids,
	posts.disliked_ids,
	users.user_id,
	users.user_lvl,
	users.user_email,
	users.user_nickname
	FROM posts, json_each(posts.disliked_ids)
	JOIN users on posts.user_id = users.user_id
	WHERE value in (%v)	
	`, thisUser.User_id, thisUser.User_id)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		var t temporaryStruct

		rows.Scan(
			&post.Post_ID,
			&post.Post_Title,
			&post.Post_Content,
			&post.Created_Time,
			&t.categories_id,
			&t.liked_ids,
			&t.disliked_ids,
			&post.User.User_id,
			&post.User.User_lvl,
			&post.User.User_email,
			&post.User.User_nickname,
		)
		parceTemporaryToPost(&post, &t, &post.Post_Categories, thisUser)
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &posts, nil
}
