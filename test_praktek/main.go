package main

import (
	"fmt"
	"qoin_test/test_praktek/util/dice_game"
)

var totalPlayer, totalDice int

func main() {
	err := getTotalPlayerAndTotalDice()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	game := dice_game.NewGame(totalPlayer, totalDice)
	game.Start()
}

func getTotalPlayerAndTotalDice() error {
	fmt.Print("Total pemain: ")
	fmt.Scan(&totalPlayer)
	fmt.Print("Total dadu: ")
	fmt.Scan(&totalDice)
	if totalPlayer < 1 {
		return fmt.Errorf("total player kurang dari 1 orang")
	}

	if totalDice < 1 {
		return fmt.Errorf("total dadu kurang dari 1 dadu")
	}
	return nil
}
