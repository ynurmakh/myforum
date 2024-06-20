package sqlite3

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"forum/internal/models"
)

type temporaryStruct struct {
	categories_id string
	liked_ids     string
	disliked_ids  string
}

func (s *Sqlite) SelectLastPostsByCount(start, onPage int, thisUser *models.User) (*[]models.Post, int, error) {
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
ORDER BY posts.created_time DESC
LIMIT ? OFFSET ?
`

	rows, err := s.db.Query(zapros, onPage, start)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	// берем вс е возможные катгорий чтоб потом по ней восстановить post.Categories
	categories, err := s.GetAllCategiries()
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}

		parceTemporaryToPost(&post, temporary, categories, thisUser)

		posts = append(posts, post)
	}

	countRow := s.db.QueryRow(`
	SELECT COUNT(DISTINCT post_id)  
	FROM posts, json_each(posts.categories_id)
	JOIN users ON posts.user_id = users.user_id
	`)
	var count int64
	err = countRow.Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return &posts, int(count), nil
}

func (s *Sqlite) FilteredSelectLastPostsByCount(start, onPage int, thisUser *models.User, categoriesInt []int) (*[]models.Post, int, error) {
	categoriesStr := make([]string, 0, len(categoriesInt))
	for _, cat := range categoriesInt {
		categoriesStr = append(categoriesStr, strconv.Itoa(cat))
	}
	catsJoined := strings.Join(categoriesStr, ",")

	zapros := fmt.Sprintf(`
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
WHERE value IN (%s)
ORDER BY posts.created_time DESC
LIMIT %v OFFSET %v
`, catsJoined, onPage, start)

	rows, err := s.db.Query(zapros)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	countRow := s.db.QueryRow(`
	SELECT COUNT(DISTINCT post_id)  
	FROM posts, json_each(posts.categories_id)
	JOIN users ON posts.user_id = users.user_id
	WHERE value IN (?)	
	`, catsJoined)
	var count int64
	err = countRow.Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	// берем вс е возможные катгорий чтоб потом по ней восстановить post.Categories
	categories, err := s.GetAllCategiries()
	if err != nil {
		return nil, 0, err
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
			return nil, 0, err
		}

		parceTemporaryToPost(&post, temporary, categories, thisUser)

		posts = append(posts, post)
	}

	return &posts, int(count), nil
}

func parceTemporaryToPost(post *models.Post, temporary *temporaryStruct, categories *[]models.Category, thisUser *models.User) error {
	// чистим ковычки
	temporary.categories_id = strings.Trim(temporary.categories_id, "[]")

	// Если категорий пустые то делаем его -1 что означает категорию "Other"
	if strings.Trim(temporary.categories_id, " ") == "" {
		temporary.categories_id = "-1"
	}

	idsString := strings.Split(temporary.categories_id, ",")
	idsInteger := make([]int, 0, len(idsString))

	for _, idString := range idsString {
		idInt, err := strconv.Atoi(strings.Trim(idString, " "))
		if err != nil {
			log.Printf("err: post id %b has illegal format of categories '%s'", post.Post_ID, idString)
			continue
		}
		idsInteger = append(idsInteger, idInt)
	}

	for i := 0; i < len(idsInteger); i++ {
		for _, category := range *categories {
			if category.Category_id == idsInteger[i] {
				post.Post_Categories = append(post.Post_Categories, category)
				break
			}
		}
	}

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
		post.Reactions.Likes = 0
	} else {

		for _, idStr := range liked_idsStr {
			liked_idInt, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("err: post %v has not convertable id '%v'", post.Post_ID, idStr)
				continue
			}
			liked_idsInt = append(liked_idsInt, liked_idInt)
		}
		post.Reactions.Likes = len(liked_idsInt)
	}

	if len(strings.Trim(strings.Trim(temporary.disliked_ids, ","), " ")) == 0 {
		post.Reactions.Dislikes = 0
	} else {
		for _, idStr := range disliked_idsStr {
			disliked_idInt, err := strconv.Atoi(idStr)
			if err != nil {
				log.Printf("err: post %v has not convertable id '%v'", post.Post_ID, idStr)
				continue
			}
			disliked_idsInt = append(disliked_idsInt, disliked_idInt)
		}
		post.Reactions.Dislikes = len(disliked_idsInt)
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
		post.Reactions.ReactByThisUser = 0
	} else if ThisUserLiked {
		post.Reactions.ReactByThisUser = 1
	} else if ThisUserDisLiked {
		post.Reactions.ReactByThisUser = -1
	}

	return nil
}
