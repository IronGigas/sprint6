package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Determine(input string) (string, error) {

	if input == "" {
		return "", errors.New("String is empty, cannot be converted")
	}

	//Какие тут еще чеки придумать? не знаю даже...

	val := strings.Split(input, " ")

	for _, word := range val {
		for _, char := range word {
			if char != '.' && char != '-' {
				return morse.ToMorse(input), nil
			}
		}
	}

	return morse.ToText(input), nil

}
