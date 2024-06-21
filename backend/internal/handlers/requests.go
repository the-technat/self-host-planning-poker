package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/the-technat/self-host-planning-poker/backend/internal/helpers"
	"github.com/the-technat/self-host-planning-poker/backend/internal/models"
	"github.com/zishang520/socket.io/socket"
)

func CreateGame(c *gin.Context) {
	var game models.Game
	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createNewGame(game)
	c.Data(http.StatusOK, "text/html", helpers.GenerateRandomBytes(8))
}

func Join(data interface{}, io *socket.Server, client *socket.Socket) gameInfo {
	dataList, ok := data.([]interface{})
	if !ok {
		fmt.Println("Input is not of type []interface{}")
		return gameInfo{}
	}
	dataMap, ok := dataList[0].(map[string]interface{})
	if !ok {
		fmt.Println("First element is not a map")
		return gameInfo{}
	}
	log.Infof("join: %v", dataMap)

	gameID := dataMap["game"].(string)
	playerName := dataMap["name"].(string)
	spectator := dataMap["spectator"].(bool)
	playerID := uuid.New().String()
	game := getGameFromUUID(gameID)
	gameInfo, gameState := game.joinGame(client.Conn().Id(), gameID, playerID, playerName, spectator)
	log.Info(gameInfo)
	io.To(socket.Room(gameID)).Emit("state", gameState)
	return gameInfo
}

func Disconnect() {
	// TODO
}

func RenameGame(data []interface{}) {
	// TODO
}

func SetDeck(data []interface{}) {
	// TODO
}

func SetPlayerName(data []interface{}) {
	// TODO
}

func PickCard(data []interface{}) {
	// TODO
}

func RevealCards(data []interface{}) {
	// TODO
}

func EndTurn(data []interface{}) {
	// TODO
}

func SetSpectator(data []interface{}) {
	log.Infof("set_spectator: %v", data)
}
