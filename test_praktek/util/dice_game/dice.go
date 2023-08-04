package dice_game

import (
	"math/rand"
	"time"
)

type Dice interface{}

type dice struct {
	number int
}

func NewDice() Dice {
	return &dice{}
}

func (d *dice) Number() int {
	return d.number
}

func (d *dice) Roll() {
	rand.Seed(time.Now().UnixNano())
	d.number = rand.Intn(6) + 1
}
