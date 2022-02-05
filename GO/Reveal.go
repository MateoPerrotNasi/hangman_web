package main

import "math/rand"

func (h *Hangman) Reveal() {
	var randomLetter int
	number := len(h.Word)/2 - 1
	if number > 0 {
		for i := 1; i <= number; i++ {
			randomLetter = rand.Intn(len(h.Word))
			for i, letter := range h.Word {
				if i == randomLetter && h.HiddenWord[i] == "_" {
					h.HiddenWord[i] = string(letter)
				}
			}
		}
	}
}