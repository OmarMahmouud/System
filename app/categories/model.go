package categories

import (
	"System/System"
	"System/exchanges/category"
	"System/models"
)

func store(req category.StoreCategory) models.Category {
	row := setHelper(req)
	System.DB.Create(&row)

	return row
}

func all() []models.Category {
	var rows []models.Category
	System.DB.Find(&rows)
	return rows
}

func GetCategoryByID(id uint) models.Category {
	var row models.Category
	System.DB.Where("id = ?", id).First(&row)
	return row
}
func deleteCategory(id uint) {
	System.DB.Where("id = ?", id).Unscoped().Delete(&models.Category{})
}

func updateModel(request category.UpdatedRequest, row *models.Category) {
	updateHelper(request, row)
	System.DB.Save(&row)
}

func categoryIsUsed(id uint) bool {
	row := models.Course{}
	System.DB.Where("category_id = ?", id).First(&row)
	if row.ID == 0 {
		return false
	}
	return true
}
