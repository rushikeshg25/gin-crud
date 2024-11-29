package main

import (
	"gin-todo/env"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init(){
	env.InitEnv()
	env.ConnectDb()
}

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() 
}