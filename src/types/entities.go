package types

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null" json:"name"`
	Role string `gorm:"type:varchar(50);not null" json:"role"`
}

type Team struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"name"`
	Striker  string `gorm:"type:varchar(255);not null" json:"striker"`
	Defender string `gorm:"type:varchar(255);not null" json:"defender"`
}

type Table struct {
	gorm.Model
	TeamName     string `gorm:"type:varchar(255);not null" json:"teamname"`
	Played       int    `json:"played"`
	Wins         int    `json:"wins"`
	Draws        int    `json:"draws"`
	Losses       int    `json:"losses"`
	GoalsFor     int    `json:"goalsfor"`
	GoalsAgainst int    `json:"goalsagainst"`
	Points       int    `json:"points"`
}

type Fixture struct {
    gorm.Model
    Home  string `gorm:"type:varchar(255);not null" json:"home"`
    Away  string `gorm:"type:varchar(255);not null" json:"away"`
    Date  string `gorm:"type:varchar(255);not null" json:"date"`
    Week  int    `json:"week"`
}
