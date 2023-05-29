package app

type Psychic struct {
	Hypothesis  int
	Reliability Reliability
}

type Reliability struct {
	Correct int
	All     int
	Percent int
}

type Battle struct {
	Psychics [5]Psychic
}

type Session struct {
	Battle Battle
}

func CreateNewSession() *Session {
	return &Session{
		Battle: *NewBattle(),
	}
}

func NewBattle() *Battle {
	return &Battle{
		Psychics: [5]Psychic{{}, {}, {}, {}, {}},
	}
}
