package services

import (
	"gorm.io/gorm"
	"namaz-vakitleri/models"
)

type ContactService struct {
	DB *gorm.DB
}

func NewContactService(db *gorm.DB) *ContactService {
	return &ContactService{DB: db}
}

func (s *ContactService) Create(contact models.Contact) error {
	return s.DB.Create(&contact).Error
}
