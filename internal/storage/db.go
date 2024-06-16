package storage

import (
	"forum/internal/models"
)

type StorageInterface interface {
	_1Cookies
	// UserRegistration
	// UserLogin
	_4Posts
	_5Categories
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
	InsertNewPost(post *models.Post) error

	//
	SelectAllPostsByCategory()

	//
	SelectAllPostsByUserID()

	// READY TO USE
	SelectLastPostsByCount(start, end int, thisUser *models.User) (*[]models.Post, error)

	//
	SelectPostByPostID(PostId int) (*models.Post, error)

	SelectComentByPostID(PostId int) ([]*models.Comment, error)
}

type _5Categories interface {
	GetCategiriesByID([]int) (*[]models.Category, error)

	// READY TO USE Возвращает из базы все категорий
	GetAllCategiries() (*[]models.Category, error)
}
