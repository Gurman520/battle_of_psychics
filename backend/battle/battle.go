package battle

// A structure describing a competition between psychics
type Battle struct {
	Psychics   [5]Psychic
	UserNumber []int
}

func NewBattle() *Battle {
	return &Battle{
		Psychics: [5]Psychic{{}, {}, {}, {}, {}},
	}
}
