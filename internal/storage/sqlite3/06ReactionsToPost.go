package sqlite3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"forum/internal/models"
)

func (s *Sqlite) ReactionsToPost(post *models.Post, thisUser *models.User, reactions int) error {
	liked_ids, err := s.postLikes(post)
	if err != nil {
		return err
	}
	disliked_ids, err := s.postDislikes(post)
	if err != nil {
		return err
	}

	var liked bool = false
	var disliked bool = false

	if _, finded := (*liked_ids)[thisUser.User_id]; finded {
		liked = true
	}
	if _, finded := (*disliked_ids)[thisUser.User_id]; finded {
		disliked = true
	}

	if liked && disliked {
		return errors.New("err: both reactins")
	}

	if reactions == 1 {
		if liked {
			err := s.likeUbratToPost(post, thisUser) // снять лайк
			if err != nil {
				return err
			}
			return nil
		}
		if disliked {
			err := s.dislikeUbratToPost(post, thisUser) // снять дизлайк
			if err != nil {
				return err
			}
			err = s.likePostavitToPost(post, thisUser) // поставить лайк
			if err != nil {
				return err
			}
			return nil
		}
		err := s.likePostavitToPost(post, thisUser) // поставить лайк
		if err != nil {
			return err
		}
		return nil
	}
	if reactions == -1 {
		if disliked {
			err := s.dislikeUbratToPost(post, thisUser) // снять диз
			if err != nil {
				return err
			}
			return nil
		}
		if liked {
			err := s.likeUbratToPost(post, thisUser) // снять лайк
			if err != nil {
				return err
			}
			err = s.dislikePostavitToPost(post, thisUser) // поставить диз
			if err != nil {
				return err
			}
			return nil
		}
		err := s.dislikePostavitToPost(post, thisUser) // поставить диз
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *Sqlite) postLikes(post *models.Post) (*map[int64]int, error) {
	query := fmt.Sprintf(`
	SELECT value
	FROM posts, json_each(posts.liked_ids) 
	WHERE posts.post_id = %v
	`, post.Post_ID)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	liked_ids := make(map[int64]int, 0)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		liked_ids[id] = 1
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &liked_ids, nil
}

func (s *Sqlite) postDislikes(post *models.Post) (*map[int64]int, error) {
	query := fmt.Sprintf(`
	SELECT value
	FROM posts, json_each(posts.disliked_ids) 
	WHERE posts.post_id = %v
	`, post.Post_ID)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	disliked_ids := make(map[int64]int, 0)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		disliked_ids[id] = -1
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &disliked_ids, nil
}

func (s *Sqlite) likePostavitToPost(post *models.Post, thisUser *models.User) error {
	query := fmt.Sprintf(`UPDATE posts SET liked_ids = json_insert(liked_ids, '$[#]', %v) WHERE post_id = %v`, thisUser.User_id, post.Post_ID)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) likeUbratToPost(post *models.Post, thisUser *models.User) error {
	query := fmt.Sprintf(`SELECT value FROM posts, json_each(posts.liked_ids) WHERE posts.post_id = %v`, post.Post_ID)
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	liked_ids := make(map[int64]int)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
		liked_ids[id] = 1
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	delete(liked_ids, int64(thisUser.User_id))

	liked_ids_after := make([]string, 0, len(liked_ids))

	for id := range liked_ids {
		liked_ids_after = append(liked_ids_after, strconv.Itoa(int(id)))
	}

	liked_ids_string := fmt.Sprintf("[%v]", strings.Join(liked_ids_after, ", "))

	query = fmt.Sprintf(`UPDATE posts SET liked_ids = '%v' WHERE post_id = %v`, liked_ids_string, post.Post_ID)
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) dislikePostavitToPost(post *models.Post, thisUser *models.User) error {
	query := fmt.Sprintf(`UPDATE posts SET disliked_ids = json_insert(disliked_ids, '$[#]', %v) WHERE post_id = %v`, thisUser.User_id, post.Post_ID)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) dislikeUbratToPost(post *models.Post, thisUser *models.User) error {
	query := fmt.Sprintf(`SELECT value FROM posts, json_each(posts.disliked_ids) WHERE posts.post_id = %v`, post.Post_ID)
	rows, err := s.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	disliked_ids := make(map[int64]int)
	for rows.Next() {
		var id int64
		err := rows.Scan(&id)
		if err != nil {
			return err
		}
		disliked_ids[id] = 1
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	delete(disliked_ids, int64(thisUser.User_id))

	disliked_ids_after := make([]string, 0, len(disliked_ids))

	for id := range disliked_ids {
		disliked_ids_after = append(disliked_ids_after, strconv.Itoa(int(id)))
	}

	disliked_ids_string := fmt.Sprintf("[%v]", strings.Join(disliked_ids_after, ", "))

	query = fmt.Sprintf(`UPDATE posts SET disliked_ids = '%v' WHERE post_id = %v`, disliked_ids_string, post.Post_ID)
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
