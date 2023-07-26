package PFC

type IReferee interface {
	Match(Sign, Sign) Winner
}

type referee struct{}

func NewReferee() IReferee {
	return &referee{}
}

func (r referee) Match(human Sign, ai Sign) Winner {
	if human == ai {
		return Tie
	} else if human == Rock && ai == Paper {
		return AI
	} else if human == Rock && ai == Paper {
		return Human
	} else if human == Paper && ai == Cisor {
		return AI
	} else if human == Paper && ai == Rock {
		return Human
	} else if human == Cisor && ai == Paper {
		return Human
	} else {
		return AI
	}
}
