package note

import "errors"

//
// ErrNoteNotFound is returned when a note is not found.
//
var ErrNoteNotFound = errors.New(`Error note not found`)
