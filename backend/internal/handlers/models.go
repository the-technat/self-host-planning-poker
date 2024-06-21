package handlers

type PlayerData struct {
	Name      string
	Spectator bool
	Game      string
	HasPicked bool
}

type Session struct {
	PlayerID string
	GameID   string
}

type game struct {
	UUID string `json:"uuid,omitempty"`
	Name string `json:"name" binding:"required"`
	Deck string `json:"deck" binding:"required"`
}


type gameInfo struct {
	Name string
	Deck string
	Revealed bool
}

type gameState struct {
	Players []PlayerData
}
