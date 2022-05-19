package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packed string) (string, error) {
	var prev rune
	var builder strings.Builder
	for _, sym := range packed {
		switch {
		case unicode.IsSpace(sym):
			return "", nil
		case unicode.IsDigit(sym) && unicode.IsLetter(prev):
			{
				if val, err := strconv.Atoi(string(sym)); err == nil && val > 0 {
					_, err = builder.WriteString(strings.Repeat(string(prev), val))
					if err != nil {
						return "", err
					}
				}
			}
		case unicode.IsDigit(sym) && !unicode.IsLetter(prev):
			return "", ErrInvalidString
		default:
			if unicode.IsLetter(prev) {
				if _, err := builder.WriteRune(prev); err != nil {
					return "", err
				}
			}
		}
		prev = sym
	}
	if unicode.IsLetter(prev) {
		if _, err := builder.WriteRune(prev); err != nil {
			return "", err
		}
	}
	return builder.String(), nil
}
