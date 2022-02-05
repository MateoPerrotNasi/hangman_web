package main

func (h *Hangman) PlayerTurn(user_input string) {
	for _, word := range h.Tried {
		if user_input == word {
			h.Attempt -= 1
		}
	}
}