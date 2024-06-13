package sqlite3

import (
	"fmt"

	"forum/internal/models"
)

func (s *Sqlite) SelectLastPostsByCount(start, onPage int) (*[]models.Post, error) {
	zapros := fmt.Sprintf(`SELECT 
	posts.post_id,
	posts.user_id,
	posts.post_title,
	posts.post_content,
	posts.created_time,
	posts.categories_id,
	posts.liked_ids,
	posts.disliked_ids,
	users.user_email,
	users.user_nickname
	FROM posts JOIN users ON posts.user_id = users.user_id ORDER BY posts.created_time LIMIT %d OFFSET %d`, onPage, start)

	rows, err := s.db.Query(zapros)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		var i struct {
			_1 interface{}
			_2 interface{}
			_3 interface{}
		}
		// err := rows.Scan(&post.Post_ID, &post.User.User_id, &post.Post_Title, &post.Post_Content, &post.Created_Time, &post.User.User_email, &post.User.User_nickname)
		err := rows.Scan(
			&post.Post_ID,
			&post.User.User_id,
			&post.Post_Title,
			&post.Post_Content,
			&post.Created_Time,
			&i._1,
			&i._3,
			&i._3,
			&post.User.User_email,
			&post.User.User_nickname,
		)
		if err != nil {
			return nil, err
		}
		fmt.Println(post)

		fmt.Println(i)

		posts = append(posts, post)
	}

	return &posts, nil
}

/*


1
2
3
4
5



*/
