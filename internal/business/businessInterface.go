package business

import (
	"forum/internal/models"
)

type Business interface {
	GetUserByCookiesValues(sessionValue string) (*models.User, error)
	GetNewCookie() (string, error)
	_Cookie
	_Posts
}

type _Cookie interface {
	// transport запрашивает у service новый uuid если у клиента нет кукиса
	CreateNewCookie() (string, error)
	// transport запрашивает у service прикриплен ли user под данным cookie
	// GetUserByCookie(sessionValue string) (*models.User, error)
	// transport запрашивает у service
	// DeregisterByCookieValue(sessionValue string) (bool, error)
}

type _Registration interface {
	// transport запрашивает у service существет ли такой user
	//  Проверка по email и/или nickname производится тут
	UserIsExist(models.User) (bool, error)
	// transport запрашивает у service создать user
	//  при успешном созданий вернется созданный user
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
	GetPostsForHome(pageNum, onPage int, categories []string) (*[]models.Post, error)

	//
	GetPost(post_id int) (*models.Post, error)
}
