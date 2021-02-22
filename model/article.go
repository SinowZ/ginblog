package model

import (
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}
