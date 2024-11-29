package main

import (
	"gin-todo/initialize"
	"gin-todo/model"
	"log"
)


func init(){
	initialize.InitEnv()
	initialize.ConnectDb()
}


func main(){
	initialize.DB.AutoMigrate(&model.Todo{})
	initialize.DB.AutoMigrate(&model.User{})
	log.Println("Migrations done")
}