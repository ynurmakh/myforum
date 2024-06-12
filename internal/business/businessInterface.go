package business

import (
	"forum/internal/models"
)

type Business interface {
	_Cookie
	_Posts
}

type _Cookie interface {
	// transport запрашивает у service новый uuid если у клиента нет кукиса
	CreateNewCookie() (string, error)
	// transport запрашивает у service прикриплен ли user под данным cookie
	GetUserByCookie(sessionValue string) (*models.User, error)
	// transport запрашивает у service отвязать юзера от этого куки
	DeregisterByCookieValue(sessionValue string) (bool, error)
}

type _Registration interface {
	// transport запрашивает у service создать user
	//  при успешном созданий вернется созданный user
	// UserIsExist Realized here
	CreateNewUser(user models.User, password string) (*models.User, error)
}

type _Login interface {
	LoginByEmailAndPass(email, pass string) (*models.User, error)
}

type _Posts interface {
	// transport запрашивает у service посты для отображения на странице :8080/home
	//  он отравляет сколько постов помещаеться на 1 странице у пользователя и на каком номере страницы
	// пример user на 3 странице, и на одну страницу у него вмещаеться 30 постов
	//  business получив запрос GetPostsForHome(3, 30, []string{}) вернет самые свежие посты с 60 по 90
	// categories []string{} должен содержать по каким категориям отсортировать посты
	//  Пустой categories вернет все посты несмотря на категорий
	// 3 orting
	GetPostsForHome(pageNum, onPage int, categories []int) (*[]models.Post, error)
	// 1
	GetPostByID(Post_ID int) (post *models.Post, reactions []int, err error)
	// 2
	CreatePost(*models.Post) error

	GetCategiries() (*[]models.Categories, error)
}

type Reactions interface {
	LikePost(Userid int, postId int, react int) error
}
