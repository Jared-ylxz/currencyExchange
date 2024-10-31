package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text;not null"`
	Content     string `gorm:"type:text;not null"`
	AuthorID    uint   `gorm:"not null;index"`      // 会成为数据库的author_id字段,成为实际起作用的外键
	Author      User   `gorm:"foreignKey:AuthorID"` // 用于关联查询，不会出现在数据库字段中
}