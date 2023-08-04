package dice_game

type Game interface {
}

type game struct {
	players []Player
}

func NewGame(totalPlayer, totalDice int) Game {
	players := []Player{}

	for count := 0; count < totalPlayer; count++ {
		players = append(players, NewPlayer(totalDice))
	}

	return &game{
		players: players,
	}
}
