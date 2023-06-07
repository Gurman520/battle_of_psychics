package battle

// A structure describing a competition between psychics
type Battle struct {
	Psychics   [5]Psychic
	UserNumber []int
}

// The function creates a new battle
func NewBattle() *Battle {
	b := &Battle{
		Psychics:   [5]Psychic{{make([]int, 0), 0}, {make([]int, 0), 0}, {make([]int, 0), 0}, {make([]int, 0), 0}, {make([]int, 0), 0}},
		UserNumber: []int{},
	}
	return b
}
