package main

import (
	"gin-todo/controller"
	"gin-todo/initialize"
	"gin-todo/middleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initialize.InitEnv()
	initialize.ConnectDb()
}

func main() {
  r := gin.Default()

  public:=r.Group("/")
  {
    public.POST("/signup",controller.Signup)
    public.POST("/signin",controller.Signin)  
  }

  secure := r.Group("/")
	secure.Use(middleware.AuthMiddleware()) 
	{
		secure.GET("/gettodo", controller.GetTodo)
		secure.POST("/addtodo", controller.AddTodo)
		secure.PUT("/updatetodo", controller.UpdateTodo)
		secure.DELETE("/deletetodo", controller.DeleteTodo)
		secure.POST("/logout", controller.Logout)
	}
  
  r.Run() 
}