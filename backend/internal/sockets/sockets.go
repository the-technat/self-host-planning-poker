package sockets

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/the-technat/self-host-planning-poker/backend/internal/handlers"
	"github.com/zishang520/socket.io/socket"
)

var io *socket.Server

func InitSockets() {
	io = socket.NewServer(nil, nil)

	io.On("connection", func(clients ...any) {
		client := clients[0].(*socket.Socket)
		client.On("join", func(data ...any) {
			log.Infof("join: %v", data)
			gameInfo := handlers.Join(data, io, client)
			var callbackFn func([]interface{}, error)
			if fn, ok := data[len(data)-1].(func([]interface{}, error)); ok {
				callbackFn = fn
				callbackFn([]interface{}{gameInfo}, nil)
			}
		})
		client.On("disconnect", func(...any) {
			log.Infof("disconnect")
			handlers.Disconnect()
		})
		client.On("rename_game", func(data ...any) {
			log.Infof("rename_game: %v", data)
			handlers.RenameGame(data)
		})
		client.On("set_deck", func(data ...any) {
			log.Infof("set_deck: %v", data)
			handlers.SetDeck(data)
		})
		client.On("set_player_name", func(data ...any) {
			log.Infof("set_player_name: %v", data)
			handlers.SetPlayerName(data)
		})
		client.On("pick_card", func(data ...any) {
			log.Infof("pick_card: %v", data)
			handlers.PickCard(data)
		})
		client.On("reveal_cards", func(data ...any) {
			log.Infof("reveal_cards: %v", data)
			handlers.RevealCards(data)
		})
		client.On("end_turn", func(data ...any) {
			log.Infof("end_turn: %v", data)
			handlers.EndTurn(data)
		})
		client.On("set_spectator", func(data ...any) {
			log.Infof("set_spectator: %v", data)
			handlers.SetSpectator(data)
		})
	})
}

func GetServeHandler() http.Handler {
	return io.ServeHandler(nil)
}
