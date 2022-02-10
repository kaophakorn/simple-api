package model

type Team struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	LeagueID int    `json:"-"`
	League   League `gorm:"foreignkey:ID;references:LeagueID"`
}
