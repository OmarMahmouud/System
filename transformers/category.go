package transformers

import (
	"System/exchanges/category"
	"System/models"
)

func TranformCategory(row *models.Category) category.CategoryResponse {
	return category.CategoryResponse{
		Name:        row.Name,
		Description: row.Description,
		ID:          row.ID,
	}
}
func TranformCategories(categories []models.Category) (responses []category.CategoryResponse) {
	for _, category := range categories {
		responses = append(responses, TranformCategory(&category))
	}
	return
}
