package dice_game

import (
	"fmt"
	"strings"
)

type Player interface {
	Play()
	Dices() []Dice
	AddPoint()
	Point() int
	RemoveDice(index int)
	PrintAllDiceNumber() string
	Info(index int)
	AddExtraDice(dice Dice)
	MergeExtra()
}

type player struct {
	dices      []Dice
	point      int
	extraDices []Dice
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

func (p *player) Play() {
	p.extraDices = []Dice{}
	for _, dice := range p.dices {
		dice.Roll()
	}
}

func (p *player) Dices() []Dice {
	return p.dices
}

func (p *player) AddPoint() {
	p.point++
}

func (p *player) Point() int {
	return p.point
}

func (p *player) RemoveDice(index int) {
	if index == 0 {
		p.dices = p.dices[index+1:]
		return
	}

	if index+1 == len(p.dices) {
		p.dices = p.dices[:index]
		return
	}

	p.dices = append(p.dices[:index], p.dices[index+1:]...)
}

func (p *player) PrintAllDiceNumber() string {
	var numbers []string
	for _, dice := range p.dices {
		numbers = append(numbers, fmt.Sprintf("%d", dice.Number()))
	}
	return strings.Join(numbers, ",")
}

func (p *player) Info(index int) {
	if len(p.Dices()) == 0 {
		fmt.Printf("Pemain #%d (%d): _ (Berhenti bermain karena tidak memiliki dadu)\n", (index + 1), p.Point())
		return
	}
	fmt.Printf("Pemain #%d (%d): %s\n", (index + 1), p.Point(), p.PrintAllDiceNumber())
}

func (p *player) AddExtraDice(dice Dice) {
	if len(p.dices) > 0 {
		p.extraDices = append(p.extraDices, dice)
	}
}

func (p *player) MergeExtra() {
	p.dices = append(p.dices, p.extraDices...)
}
