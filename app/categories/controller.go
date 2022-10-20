package categories

import (
	"System/System"
	"System/exchanges/category"
	"System/transformers"
	"github.com/gin-gonic/gin"
)

func storeControl(c *gin.Context) {
	var req category.StoreCategory
	if valid := storeValidate(c, &req); !valid {
		return
	}
	row := storeService(req)
	res := transformers.TranformCategory(&row)
	System.Created(c, res)
}
func getAllControl(c *gin.Context) {
	categories := showAllService()
	rows := transformers.TranformCategories(categories)
	System.Ok(c, rows)
}

func ShowByIdControl(c *gin.Context) {
	var request category.IdRequest
	valid, row := GetCategoryByidValidate(c, &request)
	if !valid {
		return
	}
	res := transformers.TranformCategory(row)
	System.Ok(c, res)
}

func deleteControl(c *gin.Context) {
	var request category.IdRequest
	valid, row := DeleteCategoryByidValidate(c, &request)
	if !valid {
		return
	}
	deleteService(row)
	System.Deleted(c)
}

func updatedControl(c *gin.Context) {
	var request category.UpdatedRequest
	valid, row := updateValidate(c, &request)
	if !valid {
		return
	}
	updateService(request, row)
	res := transformers.TranformCategory(row)
	System.Ok(c, res)
}
