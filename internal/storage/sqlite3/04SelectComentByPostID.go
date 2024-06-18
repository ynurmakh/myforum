package sqlite3

import (
	"forum/internal/models"
	"log"
	"strconv"
	"strings"
)

type temporaryforComment struct {
	liked_ids    string
	disliked_ids string
}

func (s *Sqlite) SelectComentByPostID(PostId int, thisUser *models.User) (*[]models.Comment, error) {
	query := `
	SELECT 
	commentaries.commentary_id,
	commentaries.commentray_content,
	commentaries.created_time,
    commentaries.liked_ids,
    commentaries.disliked_ids,
	users.user_id,
	users.user_lvl,
	users.user_email,
	users.user_nickname
	FROM commentaries
	JOIN posts ON commentaries.post_id = posts.post_id
    JOIN users ON commentaries.user_id = users.user_id
    WHERE commentaries.post_id = ?
	ORDER BY commentaries.created_time DESC;`
	rows, err := s.db.Query(query, PostId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.Comment
	for rows.Next() {
		var t temporaryforComment
		var comment models.Comment
		err := rows.Scan(
			&comment.Comment_Id,
			&comment.Commentraie_Content,
			&comment.Commentarie_Date,
			&t.liked_ids,
			&t.disliked_ids,
			&comment.User.User_id,
			&comment.User.User_lvl,
			&comment.User.User_email,
			&comment.User.User_nickname)
		if err != nil {
			return nil, err
		}

		parceTemporaryToComment(&comment, &t, thisUser)
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &comments, nil
}

func parceTemporaryToComment(comment *models.Comment, temporary *temporaryforComment, thisUser *models.User) error {
	// Теперь reactions парсим
	// Удаляем запятые
	temporary.liked_ids, temporary.disliked_ids = strings.Trim(temporary.liked_ids, "[]"), strings.Trim(temporary.disliked_ids, "[]")
	liked_idsStr, disliked_idsStr := strings.Split(temporary.liked_ids, ","), strings.Split(temporary.disliked_ids, ",")

	// Удаляем пробелы
	for i := 0; i < len(liked_idsStr); i++ {
		liked_idsStr[i] = strings.Trim(liked_idsStr[i], " ")
	}
	for i := 0; i < len(disliked_idsStr); i++ {
		disliked_idsStr[i] = strings.Trim(disliked_idsStr[i], " ")
	}

	liked_idsInt, disliked_idsInt := make([]int, 0, len(liked_idsStr)), make([]int, 0, len(disliked_idsStr))

	if len(strings.Trim(strings.Trim(temporary.liked_ids, ","), " ")) == 0 {
		comment.Reactions.Likes = 0
	} else {

		for _, idStr := range liked_idsStr {
			liked_idInt, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("err: post %v has not convertable id '%v'", comment.Comment_Id, idStr)
				continue
			}
			liked_idsInt = append(liked_idsInt, liked_idInt)
		}
		comment.Reactions.Likes = len(liked_idsInt)
	}

	if len(strings.Trim(strings.Trim(temporary.disliked_ids, ","), " ")) == 0 {
		comment.Reactions.Dislikes = 0
	} else {
		for _, idStr := range disliked_idsStr {
			disliked_idInt, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("err: post %v has not convertable id '%v'", comment.Comment_Id, idStr)
				continue
			}
			disliked_idsInt = append(disliked_idsInt, disliked_idInt)
		}
		comment.Reactions.Dislikes = len(disliked_idsInt)
	}

	if thisUser == nil {
		return nil
	}
	ThisUserLiked := false
	ThisUserDisLiked := false

	for _, liked := range liked_idsInt {
		if liked == int(thisUser.User_id) {
			ThisUserLiked = true
			break
		}
	}

	for _, disliked := range disliked_idsInt {
		if disliked == int(thisUser.User_id) {
			ThisUserDisLiked = true
			break
		}
	}

	if ThisUserLiked && ThisUserDisLiked {
		comment.Reactions.ReactByThisUser = 0
	} else if ThisUserLiked {
		comment.Reactions.ReactByThisUser = 1
	} else if ThisUserDisLiked {
		comment.Reactions.ReactByThisUser = -1
	}

	return nil
}
