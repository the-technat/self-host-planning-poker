package handlers

import (
	"github.com/google/uuid"
	"github.com/the-technat/self-host-planning-poker/backend/internal/db"
	"github.com/the-technat/self-host-planning-poker/backend/internal/models"
)

func createNewGame(game models.Game) string {
	game.UUID = uuid.New().String()
	ok := db.CreateGame(game.UUID, game.Name, game.Deck)
	if !ok {
		return ""
	}
	return game.UUID
}

func getGameFromUUID(uuid string) game {
	return game{
		UUID: uuid,
		Name: "Test Game",
		Deck: "DEFAULT",
	}
}

func (g game) joinGame(clientID string, gameID string, playerID string, playerName string, spectator bool) (gameInfo, gameState) {
	var gameState gameState
	db.JoinGame(clientID, gameID, playerID, playerName, spectator)
	gameInfo := gameInfo{
		Name:     g.Name,
		Deck:     g.Deck,
		Revealed: false,
	}
	players := db.GetPlayersByGameID(gameID)
	for _, player := range players {
		gameState.Players = append(gameState.Players, PlayerData{
			Name:      player.Name,
			Spectator: player.Spectator,
			Game:      player.GameID,
			HasPicked: player.HasPicked,
		})
	}
	return gameInfo, gameState
}
