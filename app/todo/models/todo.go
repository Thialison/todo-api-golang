package models

import (
	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `gorm:"column:title;not null"`
	Body   string `gorm:"column:body"`
	Status string `gorm:"column:status;not null"`
}
