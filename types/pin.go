package types

import (
	"errors"
	"strconv"

	"github.com/xoreo/mystery/common"
)

// ErrInvalidPin - An error thrown when the user supplies a non-4-digit pin
var ErrInvalidPin = errors.New("a non-4-digit pin was supplied")

// Pin - The security pin
type Pin struct {
	Pin    int `json:"pin"`
	Length int `json:"length"`
}

// NewPin - Create a new security pin
func NewPin(pin int) (*Pin, error) {
	s := strconv.Itoa(pin)
	if len(s) != 4 {
		return nil, ErrInvalidPin
	}
	newPin := &Pin{
		Pin:    pin,
		Length: common.DefaultPinLength,
	}
	return newPin, nil
}
