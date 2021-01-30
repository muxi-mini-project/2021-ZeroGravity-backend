package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

  func Login(c *gin.Context) {
	
	var  login  model.User
	
	if err := c.BindJSON(&login); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	u := model.User{
		Account: login.Account,
		Password: login.Password,
	}
	
	// Validate the data.
	if err := u.Validate(); err != nil {
		c.JSON(400, "用户名和密码错误")
		return
	}

}

