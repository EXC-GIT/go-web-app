package services

import (
	"errors"
	"sample-api/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	// Check if email already exists
	var existingUser models.User
	if err := s.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return models.User{}, errors.New("email already exists")
	}

	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
