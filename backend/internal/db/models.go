package db

import "gorm.io/gorm"

type game struct {
	gorm.Model
	GameID string `gorm:"unique"`
	Name   string
	Deck   string
	Player []player
}

type player struct {
	gorm.Model
	GameID    string
	PlayerID  string `gorm:"unique"`
	ClientID  string
	Name      string
	Spectator bool
	HasPicked bool
}

