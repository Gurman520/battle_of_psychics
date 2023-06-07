package server

import "errors"

var ErrForbiden = errors.New("forbidden")
var ErrNotValid = errors.New("not Valid Session")
