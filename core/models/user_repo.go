package models

import "time"
type UserRepository struct {
	Id								int
	Identity					string
	UserIdentity			string
	RepositoryIdentity	string 
	ParentId 					int
	Ext								string
	Filename					string
	CreatedAt 				time.Time `xorm:"created_at"`
	UpdatedAt 				time.Time `xorm:"updated_at"`
	DeletedAt 				time.Time `xorm:"deleted_at"`
}

func (UserRepository) TableName() string {
	return "user_repository"
}