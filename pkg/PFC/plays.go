package PFC

import "log"

type IPlay interface {
	Play(sign Sign)
}

type play struct{}

func NewGame() IPlay {
	return &play{}
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
	historic := NewHistoric()
	
	aiSign := player.AiPlay()
	winner := referee.Match(myPlay, aiSign)
	historic.Save(winner, myPlay, aiSign)

	if winner == Human {
		log.Println("You win")
	} else if winner == AI {
		log.Println("Computer win")
	} else {
		log.Println("It's a tie")
	}
}
