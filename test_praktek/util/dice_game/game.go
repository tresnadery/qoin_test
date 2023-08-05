package dice_game

import (
	"fmt"
	"strings"
)

type Game interface {
	Start()
}

type game struct {
	lastActivePlayer       []int
	playerWithHighestPoint []int
	highestPoint           int
	playerActive           int
	players                []Player
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

func (g *game) Start() {
	var round int = 1
	g.highestPoint = g.players[0].Point()
	for {
		fmt.Printf("==================\n")
		fmt.Printf("Giliran %d lempar dadu:\n", round)
		for iPlayer, player := range g.players {
			player.Play()
			player.Info(iPlayer)
		}

		g.playerActive = 0
		g.lastActivePlayer = []int{}
		g.playerWithHighestPoint = []int{}
		fmt.Printf("Setelah evaluasi:\n")
		for iPlayer := range g.players {
			g.evaluateDicePlayer(iPlayer)
		}

		for iPlayer := range g.players {
			g.mergeExtraDice(iPlayer)
		}

		if g.playerActive <= 1 {
			break
		}

		round++
	}
	g.gameResult()
}

func (g *game) evaluateDicePlayer(iPlayer int) {
	var iDice int = 0
	player := g.players[iPlayer]
	for _, dice := range player.Dices() {
		if dice.Number() == 6 {
			player.RemoveDice(iDice)
			player.AddPoint()
			iDice--
		}

		if dice.Number() == 1 {
			if iPlayer+1 == len(g.players) {
				g.players[iPlayer+1-len(g.players)].AddExtraDice(dice)
			} else {
				g.players[iPlayer+1].AddExtraDice(dice)
			}
			player.RemoveDice(iDice)
			iDice--
		}

		iDice++
	}
}

func (g *game) mergeExtraDice(iPlayer int) {
	player := g.players[iPlayer]
	player.MergeExtra()
	if len(player.Dices()) > 0 {
		g.playerActive++
		g.lastActivePlayer = append(g.lastActivePlayer, iPlayer+1)
	}

	if player.Point() >= g.highestPoint {
		g.highestPoint = player.Point()
		g.playerWithHighestPoint = append(g.playerWithHighestPoint, iPlayer+1)
	}
	player.Info(iPlayer)
}

func (g *game) gameResult() {
	fmt.Printf("==================\n")
	if len(g.lastActivePlayer) == 0 {
		fmt.Printf("Game berakhir karena semua pemain tidak memiliki dadu\n")
	} else {
		fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu.\n", g.lastActivePlayer[0])
	}
	if len(g.playerWithHighestPoint) == 1 {
		fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.\n", g.playerWithHighestPoint[0])
	} else if len(g.playerWithHighestPoint) == len(g.players) {
		fmt.Printf("Game seri karena semua pemain mempunyai poin yang sama.\n")
	} else {
		playerWithHighestPoint := []string{}
		for count := 0; count < len(g.playerWithHighestPoint); count++ {
			playerWithHighestPoint = append(playerWithHighestPoint, fmt.Sprintf("pemain#%d", g.playerWithHighestPoint[count]))
		}
		fmt.Printf("Game seri karena %s mempunyai poin yang sama.\n", strings.Join(playerWithHighestPoint, ", "))
	}
}
