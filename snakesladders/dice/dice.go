package dice

import (
	"math/rand"
	"time"
)

type Dice interface {
	Roll() int
}

type RandomDice struct{}

func (RandomDice) Roll() int {
	rand.Seed(time.Now().UnixMicro())
	return rand.Intn(6) + 1
}
