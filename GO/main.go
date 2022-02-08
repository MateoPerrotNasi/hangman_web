package main

import (
	"text/template"
	"net/http"
	"fmt"
)

type Hangman struct {
	OldUserInput string
	UserInput   string
	Word         string
	HiddenWord   []string // mot a afficher sur l'HTML
	Attempt      int
	Tried        []string
	Difficulty   string
	DiffChoosed bool
	Win          bool
}

//Reprendre le tour du joueur du début

func main() {

	Hangman_data_blank := Hangman{
		DiffChoosed: false,
	}

	menu := template.Must(template.ParseFiles("../HTML/menu.html"))

	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		Hangman_data_1 := Hangman{
			Difficulty:   r.FormValue("difficulty"),
		}
		if Hangman_data_1.Difficulty == "Facile" || Hangman_data_1.Difficulty == "Modéré" || Hangman_data_1.Difficulty == "Difficile" {
			Hangman_data_1.DiffChoosed = true
			Hangman_data_blank.Difficulty = Hangman_data_1.Difficulty	
			fmt.Println("\nDifficultée changée")
		}
		Hangman_data_blank.Difficulty = Hangman_data_1.Difficulty

		fmt.Println("La difficulté est:", Hangman_data_blank.Difficulty)

		menu.Execute(w, Hangman_data_1)
	})

	hangman := template.Must(template.ParseFiles("../HTML/hangman.html"))

	OldInput := "Ne peut pas etre égal à UserInput"

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		if Hangman_data_blank.Word == "" {
			Hangman_data_blank.ReadFile((Hangman_data_blank.Difficulty))
			Hangman_data_blank.CreateHidden(Hangman_data_blank.Word)
			Hangman_data_blank.Reveal(Hangman_data_blank.Word)
		}


		Hangman_data_2 := Hangman{
			Attempt: 10,
			Word: Hangman_data_blank.Word,
			HiddenWord: Hangman_data_blank.HiddenWord,
			OldUserInput: OldInput,
			UserInput: r.FormValue("user_input"),
		}
		fmt.Println("Le mot 2 est:",Hangman_data_2.Word)
		fmt.Println("Le mot caché est:",Hangman_data_2.HiddenWord)

		if Hangman_data_2.OldUserInput != Hangman_data_2.UserInput {
			OldInput = Hangman_data_2.UserInput
			fmt.Println("On vérifie l'input")
			Hangman_data_2.VerifInput()
			Hangman_data_2.Tried = append(Hangman_data_2.Tried, Hangman_data_2.UserInput)
			fmt.Println("OldInput vaut: ",OldInput)
			fmt.Println("OldUserInput vaut: ", Hangman_data_2.OldUserInput)
		} else {
			OldInput = Hangman_data_2.UserInput
			fmt.Println("Old et New sont egaux, on ne vérifie pas")
			fmt.Println("OldInput vaut: ",OldInput)
			fmt.Println("OldUserInput vaut: ", Hangman_data_2.OldUserInput)
		}

		hangman.Execute(w, Hangman_data_2)
	})

	http.ListenAndServe(":8080", nil)
}
