package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// Project structure
type Project struct {
	gorm.Model
	Title    string `gorm:"unique" json:"title"`
	Archived bool   `json:"archived"`
	Tasks    []Task `gorm:"ForeignKey:ProjectID" json:"tasks"`
}

// Task structure
type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Priority  string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline  *time.Time `gorm:"default:null" json:"deadline"`
	Done      bool       `json:"done"`
	ProjectID uint       `json:"project_id"`
}

//User model
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Token sturcture
type Token struct {
	UserID string
	jwt.StandardClaims
}

// DBMigrate database migration
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, &Task{}, &User{})
	db.Model(&Task{}).AddForeignKey("project_id", "projects(id)", "CASCADE", "CASCADE")
	return db
}
