package service

import (
	"errors"
	"strings"
)

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// StringService provides operations on strings.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// stringService is a concrete implementation of StringService
type stringService struct{}

func NewStringService() StringService {
	return &stringService{}
}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}
