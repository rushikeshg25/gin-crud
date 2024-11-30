package model

import "time"


type User struct {
	ID string
	Username string 
	Password string
	Todos []Todo
	CreatedAt time.Time
	UpdatedAt time.Time
}