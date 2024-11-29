package main

import (
	"gin-todo/env"
	"gin-todo/model"
	"log"
)


func init(){
	env.InitEnv()
	env.ConnectDb()
}


func main(){
	env.DB.AutoMigrate(&model.Todo{})
	env.DB.AutoMigrate(&model.User{})
	log.Println("Migrations done")
}