package PFC

var historic map[int]Histo
var index int

type Histo struct {
	Winner    Winner
	SignHuman Sign
	SignAI    Sign
}

type IHistoric interface {
	Save(winner Winner, humanSign Sign, aiSign Sign)
	GetHistoric() map[int]Histo
	GetWinner() Winner
}
type repo struct{}

func NewHistoric() IHistoric {
	historic = make(map[int]Histo)
	index = 1
	return &repo{}
}

func (r repo) Save(winner Winner, human Sign, ai Sign) {
	historic[index] = Histo{
		Winner:    winner,
		SignHuman: human,
		SignAI:    ai,
	}
	index++
}

func (r repo) GetHistoric() map[int]Histo {
	return historic
}

func (r repo) GetWinner() Winner {
	var ai int
	var human int

	for i := 1; i < index; i++ {
		if historic[i].Winner == AI {
			ai++
		} else if historic[i].Winner == Human {
			human++
		}
	}

	if human > ai {
		return Human
	} else if human < ai {
		return AI
	} else {
		return Tie
	}
}
