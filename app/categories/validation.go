package categories

import (
	"System/System"
	"System/exchanges/category"
	"System/models"
	"github.com/gin-gonic/gin"
)

func jsonValidate(g *gin.Context, request interface{}) bool {
	if err := g.ShouldBindJSON(request); err != nil {
		System.BadRequest(g, err)
		return false
	}
	return true
}

func categoryIsFound(c *gin.Context, id uint) (bool, *models.Category) {
	row := GetCategoryByID(id)
	if row.ID == 0 {
		System.NotFound(c, "category is Not Found")
		return false, nil
	}
	return true, &row
}

func storeValidate(c *gin.Context, request *category.StoreCategory) bool {
	if check, err := category.ValidStore(c.Request, request); !check {
		System.BadRequest(c, err)
		return false
	}
	return true
}

func GetCategoryByidValidate(c *gin.Context, request *category.IdRequest) (bool, *models.Category) {
	if !jsonValidate(c, request) {
		return false, nil
	}
	if found, row := categoryIsFound(c, request.ID); found {
		return true, row
	}
	return false, nil
}

func DeleteCategoryByidValidate(c *gin.Context, request *category.IdRequest) (bool, *models.Category) {
	if !jsonValidate(c, request) {
		return false, nil
	}
	if used := categoryIsUsed(request.ID); used {
		System.BadRequest(c, "Can't Delete This Category it's Used")
		return false, nil
	}
	if found, row := categoryIsFound(c, request.ID); !found {
		return true, row
	}
	return false, nil
}

func updateValidate(c *gin.Context, request *category.UpdatedRequest) (bool, *models.Category) {
	if check, err := category.ValidUbdate(c.Request, request); !check {
		System.BadRequest(c, err)
		return false, nil
	}
	if found, row := categoryIsFound(c, request.ID); found {
		return true, row
	}
	return false, nil
}
