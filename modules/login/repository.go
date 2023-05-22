package login

import (
	"gorm.io/gorm"
	// "time"
)

type Repository struct {
	DB *gorm.DB
}

func (repo Repository) CheckUsernameAndPassword(username, password string) (*User, error) {
	var users User
	result := repo.DB.Where("username = ?", username).Where("password = ?", password).First(&users)
	return &users, result.Error
}