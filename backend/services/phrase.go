package services

import (
	"gorm.io/gorm"
	"namaz-vakitleri/models"
)

type PhraseService struct {
	DB *gorm.DB
}

func NewPhraseService(db *gorm.DB) *PhraseService {
	return &PhraseService{DB: db}
}

func (s *PhraseService) GetPhrases() ([]models.Phrases, int64, error) {
	var phrases []models.Phrases
	var total int64
	err := s.DB.Model(&models.Phrases{}).Count(&total).Error
	if err != nil {
		return []models.Phrases{}, 0, err
	}

	err = s.DB.Find(&phrases).Error
	if err != nil {
		return []models.Phrases{}, 0, err
	}

	return phrases, total, nil
}

func (s *PhraseService) GetPhraseByID(id int) (models.Phrases, error) {
	var phrase models.Phrases

	err := s.DB.Where("id", id).Find(&phrase).Error
	if err != nil {
		return models.Phrases{}, err
	}

	return phrase, nil
}
