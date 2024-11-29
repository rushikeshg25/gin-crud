package controller

import (
	"gin-todo/initialize"
	"gin-todo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func AddTodo(c *gin.Context) {
	var body struct{
		Title string 
	}
	c.Bind(&body)
	newTodo:=model.Todo{
		Title: body.Title,
	}
	result:=initialize.DB.Create(&newTodo)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

    c.Status(http.StatusCreated)
}

func GetTodo(c *gin.Context) {
	var body struct{
		Title string
	}
	c.Bind(&body)
	Result:=initialize.DB.Where("title = ?", body.Title).Find(&model.Todo{})
	if Result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": Result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Result)
}

func UpdateTodo(c *gin.Context) {
	var body struct{
		OldTitle string
		NewTitle string
	}
	c.Bind(&body)
	var todo model.Todo
	result:=initialize.DB.First(&todo, "title = ?", body.OldTitle)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	todo.Title=body.NewTitle
	result=initialize.DB.Save(&todo)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTodo(c *gin.Context) {
	var body struct{
		Title string
	}
	c.Bind(&body)
	result:=initialize.DB.Where("title = ?", body.Title).Delete(&model.Todo{})
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func GetTodos(c *gin.Context) {
	var body struct{
		Title string
	}
	c.Bind(&body)
	var todos []model.Todo
	Result:=initialize.DB.Where("title <> ?", body.Title).Find(&todos)
	if Result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": Result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, todos)
}