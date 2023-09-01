package entity

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	Token          string    `json:"token" gorm:"type:varchar(100);index,unique"`
	ExpirationTime time.Time `json:"expirationTime" gorm:"type:datetime"`
	Count          int       `json:"count" gorm:"type:int(8)"`
}

func (g *Token) TableName() string {
	return "token"
}
