package controller

import "github.com/gin-gonic/gin"

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
}


func Logout(c *gin.Context) {}