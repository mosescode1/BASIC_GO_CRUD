package entity

import (
	"gorm.io/gorm"
)


type Todo struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed" gorm:"default:false;not null"`

}
