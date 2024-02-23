package models

import (
	"time"
)
type UserRepository struct {
	Id								int
	Identity					string
	UserIdentity			string
	RepositoryIdentity	string 
	ParentId 					int
	Ext								string
	Filename					string
	CreatedAt 				time.Time 
	UpdatedAt 				time.Time 
	DeletedAt 				time.Time
}

func (UserRepository) TableName() string {
	return "user_repository"
}