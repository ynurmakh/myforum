package sqlite3

import (
	"fmt"
	"strings"

	"forum/internal/models"
)

func (s *Sqlite) GetCategiriesByID(categoriesId []int) (*[]models.Category, error) {
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

	query := fmt.Sprintf("SELECT category_id, category_name FROM categories WHERE category_id IN (%s)", strings.Join(placeholders, ","))
	fmt.Println(query)

	return &[]models.Category{}, nil
}
