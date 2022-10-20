package categories

import (
	"System/exchanges/category"
	"System/models"
)

func setHelper(req category.StoreCategory) models.Category {
	return models.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}

func updateHelper(req category.UpdatedRequest, category *models.Category) {
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.Description != "" {
		category.Description = req.Description
	}
}
