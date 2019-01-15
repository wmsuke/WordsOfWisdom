package services

import (
	"github.com/wmsuke/WordsOfWisdom/domains/models"
)

// Category
type Category struct {
}

// NewCategory
func NewCategory() Category {
	return Category{}
}

// Get
func (c Category) Get(n int) interface{} {
	repo := models.NewCategoryRepository()
	user := repo.GetByID(n)
	return user
}
