package main

import (
	"errors"
	"fmt"
	"strings"
)

//go:generate stringer -type Output -linecomment

const (
	OutputText Output = iota //text
	OutputJson               //json
)

type Output uint8

func (f Output) Type() string {
	return "string"
}

var ErrInvalidOutput = errors.New("invalid output")

func (f *Output) Set(s string) error {
	switch strings.ToLower(s) {
	case OutputText.String(), "t":
		*f = OutputText
	case OutputJson.String(), "j":
		*f = OutputJson
	default:
		return fmt.Errorf("%s: %w", s, ErrInvalidOutput)
	}
	return nil
}
