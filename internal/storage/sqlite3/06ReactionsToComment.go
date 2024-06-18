package sqlite3

import (
	"fmt"
	"forum/internal/models"
	"strconv"
	"strings"
)

func (s *Sqlite) ReactionsToComment(commentId int, thisUser *models.User, reactions int) error {
	liked_ids, err := s.commentLikes(commentId)
	if err != nil {
		return err
	}
	disliked_ids, err := s.commentDislikes(commentId)
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

	if reactions == 1 {
		if liked {
			err := s.likeUbratToComment(commentId, thisUser) // снять лайк
			if err != nil {
				return err
			}
			return nil
		}
		if disliked {
			err := s.dislikeUbratToComment(commentId, thisUser) // снять дизлайк
			if err != nil {
				return err
			}
			err = s.likePostavitToComment(commentId, thisUser) // поставить лайк
			if err != nil {
				return err
			}
			return nil
		}
		err := s.likePostavitToComment(commentId, thisUser) // поставить лайк
		if err != nil {
			return err
		}
		return nil
	}
	if reactions == -1 {
		if disliked {
			err := s.dislikeUbratToComment(commentId, thisUser) // снять диз
			if err != nil {
				return err
			}
			return nil
		}
		if liked {
			err := s.likeUbratToComment(commentId, thisUser) // снять лайк
			if err != nil {
				return err
			}
			err = s.dislikePostavitToComment(commentId, thisUser) // поставить диз
			if err != nil {
				return err
			}
			return nil
		}
		err := s.dislikePostavitToComment(commentId, thisUser) // поставить диз
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (s *Sqlite) commentLikes(commentId int) (*map[int64]int, error) {
	query := fmt.Sprintf(`
	SELECT value
	FROM commentaries, json_each(commentaries.liked_ids) 
	WHERE commentaries.commentary_id = %v
	`, commentId)

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

func (s *Sqlite) commentDislikes(commentId int) (*map[int64]int, error) {
	query := fmt.Sprintf(`
	SELECT value
	FROM commentaries, json_each(commentaries.disliked_ids) 
	WHERE commentaries.commentary_id = %v
	`, commentId)

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
		disliked_ids[id] = 1
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &disliked_ids, nil
}

func (s *Sqlite) likePostavitToComment(commentId int, thisUser *models.User) error {
	query := fmt.Sprintf(`UPDATE commentaries SET liked_ids = json_insert(liked_ids, '$[#]', %v) WHERE commentary_id = %v`, thisUser.User_id, commentId)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) likeUbratToComment(commentId int, thisUser *models.User) error {
	query := fmt.Sprintf(`SELECT value FROM commentaries, json_each(commentaries.liked_ids) WHERE commentaries.commentary_id = %v`, commentId)
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

	query = fmt.Sprintf(`UPDATE commentaries SET liked_ids = '%v' WHERE commentary_id = %v`, liked_ids_string, commentId)
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) dislikePostavitToComment(commentId int, thisUser *models.User) error {
	query := fmt.Sprintf(`UPDATE commentaries SET disliked_ids = json_insert(disliked_ids, '$[#]', %v) WHERE commentary_id = %v`, thisUser.User_id, commentId)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sqlite) dislikeUbratToComment(commentId int, thisUser *models.User) error {
	query := fmt.Sprintf(`SELECT value FROM commentaries, json_each(commentaries.disliked_ids) WHERE commentaries.commentary_id = %v`, commentId)
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

	query = fmt.Sprintf(`UPDATE commentaries SET disliked_ids = '%v' WHERE commentary_id = %v`, disliked_ids_string, commentId)
	_, err = s.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
