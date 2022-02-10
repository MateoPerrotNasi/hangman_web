package main

func (h *Hangman) IsWin() {
	for i, letter := range h.Word {
		if h.HiddenWord[i] != string(letter) {
			return
		}
	}

	h.Win = true
}
