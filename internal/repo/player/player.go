package player

import (
	"simple-go-api/internal/model"
	"simple-go-api/internal/provider"

	"gorm.io/gorm/clause"
)

func GetAll() ([]model.Player, error) {
	var players []model.Player
	var err error
	db := provider.GetGormMain()
	err = db.Model(&model.Player{}).Preload(clause.Associations).Preload("Team.League").Find(&players).Error
	return players, err
}
