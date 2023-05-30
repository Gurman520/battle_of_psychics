package app

// A structure describing a psychic
type Psychic struct {
	Hypothesis  int
	Reliability Reliability
}

// A structure describing the reliability of psychic predictions
type Reliability struct {
	Correct float32
	All     float32
	Percent int
}

// A structure describing a competition between psychics
type Battle struct {
	Psychics [5]Psychic
}

// Structure describing the session
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
