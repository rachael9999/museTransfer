package models

import "time"

type User struct {
	Id          int
	Identity    string
	Name        string 
	Password    string
	Email       string
	CreatedAt   time.Time 
	UpdatedAt   time.Time 
	DeletedAt   time.Time 
}

func (table User) TableName() string {
	return "user"
}

