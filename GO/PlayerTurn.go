package main

import "strings"

func (h *Hangman) PlayerTurn(){
	IsTried := false
	found := false

	strings.ToLower(h.UserInput)
	if len(h.UserInput) > 1 {
		if h.UserInput == h.Word {
			h.HiddenWord = strings.Split(h.Word, "")
			h.Win = true
		} else {
			h.Tried = append(h.Tried, h.UserInput)
			h.Attempt -= 2
		}
	} else {
		if len(h.Tried) > 1 {
			for _, word := range h.Tried {
				if h.UserInput == word {
					IsTried = true
					h.Tried = append(h.Tried, h.UserInput)
					h.Attempt -= 1
				}
			}
		}
		if !IsTried {
			for i, letter := range h.Word {
				if h.UserInput == string(letter) {
					h.HiddenWord[i] = string(letter) 
					found = true
				}
			}
			h.Tried = append(h.Tried, h.UserInput)
			if !found {
				h.Attempt -= 1
			}
		}
	}
}
