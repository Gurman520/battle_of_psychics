// This file describes all the files that need to be read and transferred to htm/template so that the user receives ready-made html pages.
package server

// A slice of files for the game start page
var filesStart = []string{
	"./templates/start/startPage.tmpl",
	"./templates/base.tmpl",
	"./templates/start/start.tmpl",
	"./templates/finish.tmpl",
}

// A slice of files for the hypothesis page
var filesHypotheses = []string{
	"./templates/hypotheses/hypothesesPage.tmpl",
	"./templates/base.tmpl",
	"./templates/hypotheses/hypotheses.tmpl",
	"./templates/finish.tmpl",
}

// A slice of files for the result page
var filesRank = []string{
	"./templates/rating/ratingPage.tmpl",
	"./templates/base.tmpl",
	"./templates/rating/rating.tmpl",
	"./templates/finish.tmpl",
}
