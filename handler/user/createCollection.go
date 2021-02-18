package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"github.com/gin-gonic/gin"
)

// CreateCollection  is used to add a collection record of idea 新增收藏记录
func CreateCollection(c *gin.Context) {
	var increaseCollection model.CollectionModel
	if err := c.BindJSON(&increaseCollection); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
}
