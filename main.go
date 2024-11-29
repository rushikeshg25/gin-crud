package main

import (
	"gin-todo/controller"
	"gin-todo/env"

	"github.com/gin-gonic/gin"
)

func init(){
	env.InitEnv()
	env.ConnectDb()
}

func main() {
  r := gin.Default()
  r.GET("/", controller.GetTodos)
  r.GET("/gettodo", controller.GetTodo)
  r.POST("/addtodo",controller.AddTodo)
  r.PUT("/updatetodo",controller.UpdateTodo)
  r.DELETE("/deletetodo",controller.DeleteTodo)
  r.Run() 
}