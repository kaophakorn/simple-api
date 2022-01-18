package model

import (
	"time"
)

type CovidStat struct {
	ID         int        `gorm:"primaryKey" json:"-"`
	StatDate   time.Time  `json:"-"`
	TotalDeath uint       `json:"-"`
	CreatedAt  *time.Time `json:"-"`
	UpdatedAt  *time.Time `json:"-"`
}
