package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string `gorm:"type:varchar(255);not null" json:"title" binding:"required"`
	Description string `gorm:"type:text" json:"description"`
	Content     string `gorm:"type:text;not null" json:"content" binding:"required"`
	AuthorID    uint   `gorm:"not null;index:idx_author_id"` // 会成为数据库的author_id字段,成为实际起作用的外键
	Author      User   `gorm:"foreignKey:AuthorID"`          // 用于关联查询，不会出现在数据库字段中
}
