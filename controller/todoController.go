package controller

import (
	"fmt"
	"gin-todo/env"
	"gin-todo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func AddTodo(c *gin.Context) {
	var body struct{
		Title string 
	}
	c.Bind(&body)
	fmt.Println(body)
	newTodo:=model.Todo{
		Title: body.Title,
	}
	result:=env.DB.Create(&newTodo)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

    c.Status(http.StatusCreated)
}

func GetTodo(c *gin.Context) {
	// var body struct{
	// 	title string
	// }
}

func UpdateTodo(c *gin.Context) {}

func DeleteTodo(c *gin.Context) {}

func GetTodos(c *gin.Context) {}