package controller

import (
	"gin-todo/initialize"
	"gin-todo/model"
	"gin-todo/utils"

	"github.com/gin-gonic/gin"
)

type reqBody struct{
	Username string
	Password string
}

func Signin(c *gin.Context) {
	var body reqBody
	c.Bind(&body)
	
}


func Signup(c *gin.Context) {
	var body reqBody
	c.Bind(&body)
	newUser:=model.User{
		Username: body.Username,
		Password: body.Password,
	}
	result:=initialize.DB.Create(&newUser)
	if result.Error!=nil{
		c.JSON(400, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	token,err:=utils.GenerateToken(newUser.ID)
}


func Logout(c *gin.Context) {}