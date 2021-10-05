package fo_errors

import (
	"errors"
)

var ErrEmptyIssue = errors.New("empty issue")
var ErrHttpAddrWebsite = errors.New("HTTP address for website not set")
var ErrDBConnection = errors.New("DB connection string not set")

// New create error with custom string
func New(e string) error {
	return errors.New(e)
}
