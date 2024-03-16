package models

import (
	"time"

)

type ShareBasic struct {

	Id          				int
	Identity    				string
	UserIdentity        string 
	ParentId    				int
	RepositoryIdentity  string
	UserRepositoryIdentity string
	ExpireTime					int
	Name								string
	ClickNum						int
	CreatedAt   				time.Time 
	UpdatedAt   				time.Time 
	DeletedAt   				time.Time 
}

func (table ShareBasic) TableName() string {
	return "share"
}