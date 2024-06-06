package businessrealiz

import (
	"forum/internal/business"
	"forum/internal/storage"
)

type Service struct {
	storage storage.StorageInterface
}

func InitService(b storage.StorageInterface) (business.Business, error) {
	// TODO Чтение конифогов
	sqlite := &Service{
		storage: b,
	}

	var inte business.Business
	_ = inte
	inte = sqlite

	// TODO Подключение к БД

	// TODO Возврать Структуры
	return sqlite, nil
}
