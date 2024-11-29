package controller

import (
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
	var body struct{
		Title string
	}
	c.Bind(&body)
	Result:=env.DB.Where("title = ?", body.Title).Find(&model.Todo{})
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
	result:=env.DB.First(&todo, "title = ?", body.OldTitle)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	todo.Title=body.NewTitle
	result=env.DB.Save(&todo)
	if result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": result.Error.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func DeleteTodo(c *gin.Context) {}

func GetTodos(c *gin.Context) {
	var body struct{
		Title string
	}
	c.Bind(&body)
	Result:=env.DB.Where("title <> ?", body.Title).Find(&model.Todo{})
	if Result.Error!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": Result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Result)
}