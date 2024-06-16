package sqlite3

import (
	"forum/internal/models"
	"strconv"
	"strings"
)

func (s *Sqlite) InsertNewPost(post *models.Post) error {
	categStrArr := make([]string, 0, len(post.Post_Categories))
	for _, cat := range post.Post_Categories {
		categStrArr = append(categStrArr, strconv.Itoa(cat.Category_id))
	}

	categStr := "[" + strings.Join(categStrArr, ", ") + "]"

	queru := `INSERT INTO posts (user_id, post_title, post_content, categories_id) VALUES (?, ?, ?, ?)`
	res, err := s.db.Exec(queru, post.User.User_id, post.Post_Title, post.Post_Content, categStr)
	if err != nil {
		return err
	}

	postid, err := res.LastInsertId()
	if err != nil {
		return err
	}
	newpost, err := s.SelectPostByPostID(int(postid), &post.User)
	if err != nil {
		return err
	}
	*post = *newpost

	return nil
}
