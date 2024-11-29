package migrations

import (
	"gin-todo/env"
	"gin-todo/model"
)


func init(){
	env.InitEnv()
	env.ConnectDb()
}


func main(){
	env.DB.AutoMigrate(&model.Todo{})
}