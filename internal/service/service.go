package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectEncoding(input string) (string, error) {

	if input == "" {
		return "", errors.New("String is empty, cannot be converted")
	}

	f := func(r rune) bool {
		return r == '.' || r == '-'
	}

	if strings.ContainsFunc(input, f) {
		return morse.ToText(input), nil
	}

	f2 := func(r rune) bool {
		return r >= 'а' && r <= 'я' || r >= 'А' && r <= 'Я'
	}

	if strings.ContainsFunc(input, f2) {
		return morse.ToMorse(input), nil
	}

	return "", errors.New("Usupported text")

}
