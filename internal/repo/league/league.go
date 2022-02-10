package league

import (
	"simple-go-api/internal/model"
	"simple-go-api/internal/provider"
)

func GetAll() ([]model.League, error) {
	var leagues []model.League
	var err error
	db := provider.GetGormMain()
	err = db.Model(&model.League{}).Find(&leagues).Error
	return leagues, err
}
