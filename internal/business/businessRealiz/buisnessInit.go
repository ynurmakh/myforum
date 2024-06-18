package businessrealiz

import (
	"encoding/json"
	"os"

	"forum/internal/business"
	"forum/internal/storage"
	"forum/internal/storage/sqlite3"
)

type Service struct {
	storage *sqlite3.Sqlite
	configs *ConfigType
}

type ConfigType struct {
	CookieMaxAge          int    `json: "CookieMaxAge"`
	PasswordHashingSecret string `json: "PasswordHashingSecret"`
}

func InitService(b storage.StorageInterface) (business.Business, error) {
	config, err := configsParce()
	if err != nil {
		return nil, err
	}

	service := &Service{
		storage: b.(*sqlite3.Sqlite),
		configs: config,
	}

	// TODO Подключение к БД

	// TODO Возврать Структуры
	return service, nil
}

func configsParce() (*ConfigType, error) {
	config := &ConfigType{}

	file, err := os.ReadFile("internal/business/configs.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
