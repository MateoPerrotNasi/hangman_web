package main

import "fmt"

func (h *Hangman) CreateHidden(Word string) {
	for i := 0; i < len(Word); i++ {
		h.HiddenWord = append(h.HiddenWord, "_")
	}
	fmt.Println("Create Hiddenword:",h.HiddenWord)
}