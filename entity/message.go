package entity

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Token   string `json:"token" gorm:"type:varchar(100)"`
	Message string `json:"message" gorm:"type:varchar(10000)"`
}

func (g *Message) TableName() string {
	return "message"
}
