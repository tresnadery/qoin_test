package dice_game

type Player interface{}

type player struct {
	dices []Dice
	point int
}

func NewPlayer(diceCount int) Player {
	var dices []Dice
	for count := 0; count < diceCount; count++ {
		dices = append(dices, NewDice())
	}
	return &player{
		dices: dices,
	}
}
