package main

func (h *Hangman) CreateHidden() {
	for i := 0; i < len(h.Word); i++ {
		h.HiddenWord = append(h.HiddenWord, "_")
	}
}