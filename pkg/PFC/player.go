package PFC

import "math/rand"

type IPlayer interface {
	AiPlay() Sign
}

type player struct{}

func NewPlayer() IPlayer {
	return &player{}
}

func (p player) AiPlay() Sign {
	sign := rand.Intn(2)
	switch sign {
	case 0:
		return Rock
	case 1:
		return Cisor
	case 2:
		return Paper
	}
	return Rock
}
