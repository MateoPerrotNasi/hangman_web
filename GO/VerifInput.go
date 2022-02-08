package main

import "fmt"

func (h *Hangman) VerifInput() {
    if len(h.UserInput) < 2 {
        for i := 0; i < len(h.Word); i++ {
            if h.UserInput == string(h.Word[i]) {
                h.HiddenWord[i] = h.UserInput
                fmt.Println("Lettre trouvée")
            }
        }
    } else if h.UserInput == h.Word {
        fmt.Println("SUUUUUUUU")
        for i := 0; i < len(h.Word); i++ {
            if h.UserInput == string(h.Word[i]) {
                h.HiddenWord[i] = h.UserInput
            }
        }
        h.Win = true
    } else {
        fmt.Println("La lettre n'est pas présente")
        h.Attempt -= 1
    }
}