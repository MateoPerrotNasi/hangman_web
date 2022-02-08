package main

import (
	"math/rand"
	"fmt"
)

func (h *Hangman) Reveal(Word string) {
	var randomLetter int
	number := len(Word)/2 - 1
	if number > 0 {
		for i := 1; i <= number; i++ {
			randomLetter = rand.Intn(len(Word))
			for i, letter := range Word {
				if i == randomLetter && h.HiddenWord[i] == "_" {
					h.HiddenWord[i] = string(letter)
				}
			}
		}
	}
	fmt.Println("Reveal Hiddenword:", h.HiddenWord)
}