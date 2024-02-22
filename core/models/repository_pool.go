package models

import "time"

type RepositoryPool struct {
	Id        int
	Identity  string
	Hash      string
	Filename      string
	Ext       string
	Size      int64
	Path      string
	CreatedAt time.Time `xorm:"created_at"`
	UpdatedAt time.Time `xorm:"updated_at"`
	DeletedAt time.Time `xorm:"deleted_at"`
}

func (table RepositoryPool) TableName() string{
	return "repository_pool"
}