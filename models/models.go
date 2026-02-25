package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Posts     []Post // Has Many: One user has many posts
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"not null"`
	Content   string         `gorm:"not null"`
	UserID    uint           `gorm:"not null"`
	Comments  []Comment      // Has Many: One post has many comments
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除字段
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	PostID    uint   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
