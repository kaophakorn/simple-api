package model

import "time"

type Player struct {
	ID         int       `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Middlename string    `json:"middlename"`
	Lastname   string    `json:"lastname"`
	Pos        string    `json:"position"`
	TeamID     int       `json:"-"`
	Team       Team      `gorm:"foreignkey:ID;references:TeamID"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
