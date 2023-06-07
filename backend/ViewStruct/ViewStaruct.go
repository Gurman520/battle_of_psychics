// This file describes the structures that are used to display information on html pages
package viewstruct

// Start page
type ViewDataStartPage struct {
	Title string
}

// Hypothesis Page
type ViewDataHypothesesPage struct {
	Title      string
	Hypothesis []int
}

// Results page
type ViewDataRankPage struct {
	Title       string
	History     map[int][]int
	Reliability []int
}
