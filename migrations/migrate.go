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
	log.Println("Migrations done")
}