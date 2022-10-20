package categories

import (
	"System/exchanges/category"
	"System/models"
)

func storeService(request category.StoreCategory) models.Category {
	return store(request)
}

func showAllService() []models.Category {
	return all()
}

func deleteService(row *models.Category) {
	deleteCategory(row.ID)
}

func updateService(req category.UpdatedRequest, row *models.Category) {
	updateModel(req, row)
}
