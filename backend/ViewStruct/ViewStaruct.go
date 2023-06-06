package viewstruct

type ViewDataStartPage struct {
	Title string
}

type ViewDataHypothesesPage struct {
	Title      string
	Hypothesis []int
}

type ViewDataRankPage struct {
	Title       string
	History     map[int][]int
	Reliability []int
}
