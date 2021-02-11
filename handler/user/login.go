package user

import (
	"github.com/2021-ZeroGravity-backend/model"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)
 type UserModel struct {
	Account        string    `json:"account"`
	Nickname       string    `json:"nickname"`


}



func (u *UserModel) Create() error {
	return DB.Create(&u).Error
}




func (u *UserModel) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}




func Login(c *gin.Context) {
	
	var  login  model.User
	// 获取账号和密码
	if err := c.BindJSON(&login); err != nil {
		c.JSON(400, gin.H{"message": "输入有误，格式错误"})
		return
	}
	u := model.User{
		Account: login.Account,
		Password: login.Password,
	}
	
	// Validate the data.检验账号和密码
	if err := u.Validate(); err != nil {
		c.JSON(400, "用户名和密码错误")
		return
	}
	// 插入到数据库
    if err := u.Create(); err != nil {
		c.JSON(400, "插入数据失败")
		return
	
	}
}

