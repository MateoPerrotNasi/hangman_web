package main

import (
	"net/http"
	"strconv"
	"text/template"
	"fmt"
)

type Hangman struct {
	UserInput   string
	Word        string
	HiddenWord  []string
	Attempt     int
	Tried       []string
	Difficulty  string
	DiffChoosed bool
	Win         bool
	Link string
}

func main() {

	Hangman_data_blank := Hangman{
		DiffChoosed: false,
		Attempt:     11,
		Word:        "",
	}

	menu := template.Must(template.ParseFiles("../HTML/menu.html"))

	http.HandleFunc("/menu", func(w http.ResponseWriter, r *http.Request) {
		Hangman_data_1 := Hangman{
			Difficulty: r.FormValue("difficulty"),
		}
		if Hangman_data_1.Difficulty == "Facile" || Hangman_data_1.Difficulty == "Modéré" || Hangman_data_1.Difficulty == "Difficile" {
			Hangman_data_1.DiffChoosed = true
			Hangman_data_blank.Difficulty = Hangman_data_1.Difficulty
		}
		Hangman_data_blank.Difficulty = Hangman_data_1.Difficulty

		menu.Execute(w, Hangman_data_1)
	})

	css := http.FileServer(http.Dir("../CSS/"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	images := http.FileServer(http.Dir("../images/"))
	http.Handle("/images/", http.StripPrefix("/images/", images))

	hangman := template.Must(template.ParseFiles("../HTML/hangman.html"))

	victory := template.Must(template.ParseFiles("../HTML/win.html"))

	defeat := template.Must(template.ParseFiles("../HTML/loose.html"))

	http.HandleFunc("/hangman", func(w http.ResponseWriter, r *http.Request) {
		if Hangman_data_blank.Word == "" {
			Hangman_data_blank.ReadFile()
			Hangman_data_blank.CreateHidden()
			Hangman_data_blank.Reveal()
		}

		Hangman_data_blank = Hangman{
			UserInput:  r.FormValue("UserInput"),
			Difficulty: Hangman_data_blank.Difficulty,
			Word:       Hangman_data_blank.Word,
			HiddenWord: Hangman_data_blank.HiddenWord,
			Tried:      Hangman_data_blank.Tried,
			Attempt:    Hangman_data_blank.Attempt,
			Link: strconv.Itoa(Hangman_data_blank.Attempt),
			Win:        Hangman_data_blank.Win,
		}

		Hangman_data_blank.PlayerTurn()
		Hangman_data_blank.IsWin()

		if Hangman_data_blank.Win {
			http.Redirect(w, r, "/win", http.StatusSeeOther)
		} else {
			if Hangman_data_blank.IsLoose() {
				http.Redirect(w, r, "/loose", http.StatusSeeOther)
			} else {
				hangman.Execute(w, Hangman_data_blank)
			}
		}
	})

	http.HandleFunc("/win", func(w http.ResponseWriter, r *http.Request) {
		victory.Execute(w, Hangman_data_blank)
	})

	http.HandleFunc("/loose", func(w http.ResponseWriter, r *http.Request) {
		defeat.Execute(w, Hangman_data_blank)
	})

	fmt.Println("Serveur fonctionnel")
	http.ListenAndServe(":8080", nil)
}
