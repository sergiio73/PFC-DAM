package models

import "errors"

type Game struct {
	ID          int
	Title       string
	Description string
	Date        string
	MasterName  string
	MaxPlayers  int
	Location    string
}

func CreateGame(game Game) error {
	game.ID = NextGameID
	NextGameID++
	Games = append(Games, game)
	return nil
}

func GetAllGames() ([]Game, error) {
	gamesCopy := []Game{}
	for _, game := range Games {
		gamesCopy = append(gamesCopy, game)
	}
	return gamesCopy, nil
}

func GetGameByID(id int) (Game, error) {
	for _, game := range Games {
		if game.ID == id {
			return game, nil
		}
	}
	return Game{}, errors.New("partida no encontrada")
}
