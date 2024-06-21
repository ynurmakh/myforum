package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) GetCategiriesByI(categoriesId []int) (*[]models.Category, error) {
	if len(categoriesId) == 0 {
		return &[]models.Category{}, nil
	}

	// Создаем placeholders для IN-условия
	placeholders := make([]string, len(categoriesId))
	args := make([]interface{}, len(categoriesId))
	for i, id := range categoriesId {
		placeholders[i] = "?"
		args[i] = id
	}

	return &[]models.Category{}, nil
}
