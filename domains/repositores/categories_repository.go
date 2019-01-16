package repositores

import "github.com/wmsuke/WordsOfWisdom/domains/models"

// CategoryRepository is
type CategoryRepository struct {
}

// NewCategoryRepository ...
func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{}
}

// GetByID ...
func (m CategoryRepository) GetByID(id int) *models.Categories {
	// var category = models.Categories{ID: id}
	// has, _ := engine.Get(&category)
	// if has {
	// 	return &category
	// }

	return nil
}
