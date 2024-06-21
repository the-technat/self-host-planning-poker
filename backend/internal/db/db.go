package db

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	db, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&game{}, &player{})
	if err != nil {
		panic("failed to migrate database")
	}
}

func CreateGame(uuid string, name string, deck string) bool {
	tx := db.Begin()
	if err := tx.Create(&game{GameID: uuid, Name: name, Deck: deck}).Error; err != nil {
		log.Errorf("failed to create game %s with error: %s", uuid, err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func JoinGame(clientID string, gameID string, playerID string, playerName string, spectator bool) bool {
	tx := db.Begin()
	if err := tx.Create(&player{ClientID: clientID, GameID: gameID, PlayerID: playerID, Name: playerName, Spectator: spectator, HasPicked: false}).Error; err != nil {
		log.Errorf("failed to join game %s with error: %s", gameID, err)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetPlayersByGameID(gameID string) []player {
	var players []player
	db.Where("game_id = ?", gameID).Find(&players)
	return players
}
