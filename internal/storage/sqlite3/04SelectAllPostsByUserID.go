package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) SelectAllPostsByUserID(thisUser *models.User) (*[]models.Post, error) {
	zapros := `
SELECT DISTINCT
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
FROM posts, json_each(posts.categories_id)
JOIN users ON posts.user_id = users.user_id
WHERE posts.user_id = ?
ORDER BY posts.created_time DESC
`

	rows, err := s.db.Query(zapros, thisUser.User_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// берем вс е возможные катгорий чтоб потом по ней восстановить post.Categories
	categories, err := s.GetAllCategiries()
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		temporary := &temporaryStruct{}
		err := rows.Scan(
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

		posts = append(posts, post)
	}

	return &posts, nil
}
