package service

import (
	"errors"
	"strings"
)

// StringService provides some services for strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// New returns the default implementation of StringService.
func New() StringService { return stringService{} }

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when the string to uppercase is empty.
var ErrEmpty = errors.New("Empty String")
