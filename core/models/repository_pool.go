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
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt time.Time 
}

func (table RepositoryPool) TableName() string{
	return "repository_pool"
}