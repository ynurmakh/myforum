package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) SelectPostByPostID(PostID int, thisUser *models.User) (*models.Post, error) {
	categories, err := s.GetAllCategiries()
	if err != nil {
		return nil, err
	}

	query := `SELECT
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
FROM posts
JOIN users ON posts.user_id = users.user_id
WHERE posts.post_id = ?`

	row := s.db.QueryRow(query, PostID)

	var post models.Post
	temporary := &temporaryStruct{}
	err = row.Scan(
		&post.Post_ID,
		&post.Post_Title,
		&post.Post_Content,
		&post.Created_Time,
		&temporary.categories_id,
		&temporary.liked_ids,
		&temporary.disliked_ids,
		&post.User.User_id,
		&post.User.User_lvl,
		&post.User.User_email,
		&post.User.User_nickname,
	)
	if err != nil {
		return nil, err
	}

	parceTemporaryToPost(&post, temporary, categories, thisUser)

	comments, err := s.SelectComentByPostID(int(post.Post_ID), thisUser)
	if err != nil {
		return nil, err
	}
	post.Comments = *comments

	return &post, nil
}
