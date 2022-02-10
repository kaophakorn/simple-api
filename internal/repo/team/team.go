package team

import (
	"simple-go-api/internal/model"
	"simple-go-api/internal/provider"

	"gorm.io/gorm/clause"
)

func GetAll() ([]model.Team, error) {
	var teams []model.Team
	var err error
	db := provider.GetGormMain()
	err = db.Model(&model.Team{}).Preload(clause.Associations).Find(&teams).Error
	return teams, err
}
