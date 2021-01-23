package main

import (
	"context"
	"path/filepath"

	"github.com/gin"
)
  func main (){
  engine:=gin.Default()
 
  engine.POST("/register",func (context *gin.Context) {
	  fmt.Println("context.FullPath"())
	  user_name:=context.Query("user_name")
	  fmt.Println(user_name)
	  context.Writer.Writer([]byte("用户注册" + user_name))
  })
   
  engine.POST("/login",func (context *gin.Context)  {
	fmt.Println("context.FullPath"()) 
	user_name,exist := context.GetPostForm("user_name")
	password,exist := context.GetPostForm("password")
	if exist{
	    fmt.Println(user_name)
	    fmt.Println(password)
    }
	
	context.Writer.Writer([]byte("用户登录" +user_name))
	  
  })
  engine.DELETE("/user/:user_id",func(c *gin.Context) {
		user_id:=context.Param("user_id")
		context.Writer.Writer([]byte("delete用户ID：" +user_id ))


  })
   

    engine.POST("/login",func (context *gin.Context) {

	fmt.Println("context.FullPath"()) 
	user_name,exist := context.GetPostForm("user_name")
	password,exist := context.GetPostForm("password")
    if exist{
	    fmt.Println(user_name)
	    fmt.Println(password)
    }
	
	context.Writer.Writer([]byte("用户登录" +user_name))

})

	
    engine.POST("/register",func (context *gin.Context) {
	  
	fmt.Println("context.FullPath"()) 
	user_name,exist := context.GetPostForm("user_name")
	password,exist := context.GetPostForm("password")
    if exist{
	    fmt.Println(user_name)
	    fmt.Println(password)
    }
	
	context.Writer.Writer([]byte("用户注册" +user_name))

})
	

    engine.DELETE("/user/:user_id",func(c *gin.Context) {
	user_id:=context.Param("user_id")
	context.Writer.Writer([]byte("delete用户ID：" +user_id ))


})
	

    engine.POST("/idea",func (context *gin.Context) {
	  
	fmt.Println("context.FullPath"()) 
	user_name,exist := context.GetPostForm("user_name")
	password,exist := context.GetPostForm("password")
    if exist{
	    fmt.Println(user_name)
	    fmt.Println(password)
    }
	
	context.Writer.Writer([]byte("用户注册" +user_name))

})
	


  engine.Run()
}