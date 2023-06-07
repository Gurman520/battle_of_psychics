// The file describes all the errors that are used in the project
package server

import "errors"

var ErrForbiden = errors.New("forbidden")
var ErrNotValid = errors.New("not Valid Session")
