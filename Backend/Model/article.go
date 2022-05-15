package Model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string   `gorm:"type:varchar(100);not null;" json:"title"`
	Category    Category `gorm:"type:varchar(100);not null;" json:"category"`
	Cid         int      `gorm:"type int" json:"cid"`
	Description string   `gorm:"type:varchar(200)" json:"description"`
	//Content     string   `gorm:"type:longtext" json:"content"`
	Img string `gorm:"type:varchar(100)" json:"img"`
}
