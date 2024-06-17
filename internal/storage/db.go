package storage

import (
	"forum/internal/models"
)

type StorageInterface interface {
	_1Cookies
	_2UserRegistration
	_3UserLogin
	_4Posts
	_5Categories
	_6Reactions
}

type _1Cookies interface {
	// Service отправляет storage UUID и userEmail и storage привязывает за UUID данного user`а
	TieCookieToUser(UserID int, UUID string, DeadTimeSeconds int) (bool, error)
	// Service отправляет storage UUID кукиса и storage возврашает какой юзер закреплен за данным UUID. Если UUID не сущ. в БД то создаст его и вернет nil,nil. Данный метод автоматический продливает жизннь UUID до time.Now() + expireTime
	CheckTheCookie(cookie string, expireTime int) (*models.User, error)
	// Service отправляет storage UUID кукиса и storage открепляет user`а за данным UUID
	KillCookie(UUID string) (bool, error)
}

type _2UserRegistration interface {
	// Service отправляет email и storage возвращает ture если сущ. юзер с такой почтой, или false если нет.
	IsExistByEmail(email string) (bool, error)
	// Service отправляет NickName и storage возвращает ture если сущ. юзер с такой NickName, или false если нет.
	IsExistByNickName(email string) (bool, error)
	// Service отправляет email, NickName и password в storage, после storage возвращает true/false
	InsertNewUserByEmailAndPass(email, nickname, password string) (bool, error)
}

type _3UserLogin interface {
	// Сервис отдает БД email юзера.
	//  Если пароль который отправил сервис являееться "GUARANTEED" то сторона storage его не проверяет.
	// А если пароль не "GUARANTEED", то storage проверит совпадение пароля и вернят *User только если верно, иначе nil, nil
	GetUserByEmailAndPass(email string, hashed_password string) (*models.User, error)
}

type _4Posts interface {
	// Отправляеться userID(Владелец поста), заголовог поста, контент поста и катогорий массивом стрингов, в ответ жду номер post_id который был успешно создан
	InsertNewPost(post *models.Post, cats []int) error

	InsertNewComment(post *models.Post, comment *models.Comment) error
	//
	SelectAllPostsByCategory()

	//
	SelectAllPostsByUserID()

	// READY TO USE
	SelectLastPostsByCount(start, end int, thisUser *models.User) (*[]models.Post, error)

	//
	SelectPostByPostID(PostId int, thisUser *models.User) (*models.Post, error)

	// InsertNewComment(forPost int) (commentaries *models.Comment, error)
	SelectComentByPostID(PostId int) (*[]models.Comment, error)
}

type _5Categories interface {
	GetCategiriesByID([]int) (*[]models.Category, error)

	// READY TO USE Возвращает из базы все категорий
	GetAllCategiries() (*[]models.Category, error)
}

type _6Reactions interface {
	// 1
	// NOT REALIZED
	ReactionsToPost(post *models.Post, thisUser *models.User, reactions int) error
	// 2
	// NOT REALIZED
	ReactionsToComment(commentId int, thisUser *models.User, reactions int) error
}
