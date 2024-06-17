package business

import (
	"forum/internal/models"
)

type Business interface {
	_1Cookie
	_2Registration
	_3Login
	_4Posts
	_5Commentaries
	_6Reactions
}

type _1Cookie interface {
	// transport запрашивает у service новый uuid если у клиента нет кукиса
	// NOT REALIZED
	CreateNewCookie() (string, error)
	// transport запрашивает у service прикриплен ли user под данным cookie
	// NOT REALIZED
	GetUserByCookie(sessionValue string) (*models.User, error)
	// transport запрашивает у service отвязать юзера от этого куки
	// NOT REALIZED
	DeregisterByCookieValue(sessionValue string) (bool, error)
}

type _2Registration interface {
	// transport запрашивает у service создать user
	//  при успешном созданий вернется созданный user
	// UserIsExist Realized here
	// NOT REALIZED
	CreateNewUser(user models.User, password string) (*models.User, error)
}

type _3Login interface {
	// NOT REALIZED
	LoginByEmailAndPass(email, pass string) (*models.User, error)
}

type _4Posts interface {
	// transport запрашивает у service посты для отображения на странице :8080/home
	//  он отравляет сколько постов помещаеться на 1 странице у пользователя и на каком номере страницы
	// пример user на 3 странице, и на одну страницу у него вмещаеться 30 постов
	//  business получив запрос GetPostsForHome(3, 30, []string{}) вернет самые свежие посты с 60 по 90
	// categories []string{} должен содержать по каким категориям отсортировать посты
	//  Пустой categories вернет все посты несмотря на категорий
	// READY TO USE
	GetPostsForHome(pageNum, onPage int, categories []int, thisUser *models.User) (*[]models.Post, error)
	// READY TO USE
	GetPostByID(Post_ID int, thisUser *models.User) (post *models.Post, err error)
	// READY TO USE
	CreatePost(post *models.Post, categiresNum []int) error
	// REAY TO USE
	GetCategiries() (*[]models.Category, error)
	// Получить количество всех постов для пагинаций
	GetCountOfPosts() (int, error)
}

type _5Commentaries interface {
	// REAY TO USE
	CraeteCommentary(forComment *models.Post, comment *models.Comment) error
}

type _6Reactions interface {
	// READY TO USE
	ReactionsToPost(post *models.Post, thisUser *models.User, reactions int) error
	// READY TO USE
	ReactionsToComment(commentId int, thisUser *models.User, reactions int) error
}
