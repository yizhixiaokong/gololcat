package main

import (
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
)

func TestPrint(t *testing.T) {
	var phrases []string
	for range 5 {
		phrases = append(phrases, gofakeit.HackerPhrase())
	}

	output := strings.Join(phrases[:], " \n")

	LolcatPrint([]rune(output))
}
