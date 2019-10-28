package categories

import (
	. "ant-forum/handler/v1"
	"ant-forum/model"
	"ant-forum/pkg/errno"
	"github.com/gin-gonic/gin"
	"strconv"
)

// @Summary 用分类id获取单个分类信息
// @Description 用分类id获取单个分类信息
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path string true "Id"
// @Success 200 {object} model.CategoryInfo "{"code":0,"message":"OK","data":{"id":0,"category_name":"前端"}}"
// @Router /v1/categories/{id} [get]
func GetCategoryById(c *gin.Context) {
	categoryId, _ := strconv.Atoi(c.Param("id"))
	// Get the category by the `id` from the database.
	category, err := model.GetCategoryById(uint64(categoryId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, &model.CategoryInfo{Id: category.Id, CategoryName: category.CategoryName})
}