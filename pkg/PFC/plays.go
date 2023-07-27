package PFC

import (
	"fmt"
	"log"
)

var histo IHistoric

type IPlay interface {
	Play(sign Sign)
	GetHistoric()
	End()
}

type play struct {
	historic IHistoric
}

func NewGame(historic IHistoric) IPlay {
	return &play{historic: historic}
}

type Sign string

const (
	Rock  Sign = "Rock"
	Paper      = "Paper"
	Cisor      = "Cisor"
)

type Winner string

const (
	Human Winner = "Human"
	AI           = "AI"
	Tie          = "Tie"
)

func (p play) Play(myPlay Sign) {

	referee := NewReferee()
	player := NewPlayer()

	aiSign := player.AiPlay()
	winner := referee.Match(myPlay, aiSign)
	p.historic.Save(winner, myPlay, aiSign)

	if winner == Human {
		log.Println("You win")
	} else if winner == AI {
		log.Println("Computer win")
	} else {
		log.Println("It's a tie")
	}
}

func (p play) GetHistoric() {
	fmt.Println(p.historic.GetHistoric())
}

func (p play) End() {
	winner := p.historic.GetWinner()
	if winner == Human {
		log.Println("You win the game")
	} else if winner == AI {
		log.Println("Computer win the game")
	} else {
		log.Println("It's a tie")
	}
}
