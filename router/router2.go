package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin"
)

func main  () { 
	engine := gin.Default()
	//http://localhost:8080/user?user_id=?
	engine.Handle( httpMethod:"GET",relativePath:"/user",func(c *gin.Context) {
	  path:=context.FullPath()
	  fmt.Println(path)
	  name:=context.DefaultQuery(key:"name",)
	  fmt.Println(name)
	  //输出
	  context.Writer.Writer([]byte("user," + name)  )
    })

	engine.Run()
	
}
    //post
    //http://localhost:8080/login?user_id=?
engine.Handle( httpMethod:"POST",relativePath:"/login",func(c *gin.Context) {
	fmt.Println(context.FullPath())
	user_name:=context.PostForm(key:"user_name")
	password:=context.PostForm(key:"password")
	fmt.Println(user_name)
	fmt.Println(password)
	context.Writer.Writer([]byte(user_name + "登录"))
    engine.Run()
}
