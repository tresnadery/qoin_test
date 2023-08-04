package dice_game

type Player interface {
	Play()
	Dices() []Dice
	AddPoint()
	Point() int
}

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

func (p *player) Play() {
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
