package sqlite3

import (
	"forum/internal/models"
)

func (s *Sqlite) GetAllCategiries() (*[]models.Categories, error) {
	rows, err := s.db.Query(`SELECT * FROM categories_name`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categoris := []models.Categories{}

	for rows.Next() {
		var category models.Categories
		err := rows.Scan(&category.Category_id, &category.Category_name)
		if err != nil {
			return nil, err
		}

		categoris = append(categoris, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &categoris, nil
}
