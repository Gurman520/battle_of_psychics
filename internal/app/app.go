package app

type Psychic struct {
	Hypothesis  int
	Reliability Reliability
}

type Reliability struct {
	Correct float32
	All     float32
	Percent int
}

type Battle struct {
	Psychics [5]Psychic
}

type Session struct {
	Battle Battle
	Result bool
}

func CreateNewSession() *Session {
	return &Session{
		Battle: *NewBattle(),
		Result: false,
	}
}

func NewBattle() *Battle {
	return &Battle{
		Psychics: [5]Psychic{{}, {}, {}, {}, {}},
	}
}
