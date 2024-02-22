package models

import "time"

type Code struct {
	Value      string
	Expiration time.Time
}

func (table Code) TableName() string {
	return "code"
}