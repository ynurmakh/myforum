package sqlite3

import (
	"errors"
	"fmt"

	"forum/internal/models"
)

func (s *Sqlite) ReactionsToPost(post *models.Post, thisUser *models.User, reactions int) error {
	liked_ids, err := s.postLikes(post, thisUser, reactions)
	if err != nil {
		return err
	}

	disliked_ids, err := s.postDislikes(post, thisUser, reactions)
	if err != nil {
		return err
	}

	var liked bool
	var disliked bool
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
			// снять лайк
		}
		if disliked {
			// снять дизлайк
			// поставить лайк
		}

	}
	if reactions == -1 && !liked {
		// поставить лайк
	}

	return nil
}

func (s *Sqlite) postLikes(post *models.Post, thisUser *models.User, reactions int) (*map[int64]int, error) {
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

	return &liked_ids, nil
}

func (s *Sqlite) postDislikes(post *models.Post, thisUser *models.User, reactions int) (*map[int64]int, error) {
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

	return &disliked_ids, nil
}
