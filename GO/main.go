package main

import (
    "text/template"

	"net/http"
)

type Hangman struct {
	User_input string
	Word       string
	HiddenWord []string // mot a afficher sur l'HTML
	Attempt    int
	Tried      []string
	Difficulty string
	Diff_choosed bool
}

//Reprendre le tour du joueur du début

func main() {
	menu := template.Must(template.ParseFiles("../HTML/menu.html"))
	hangman := template.Must(template.ParseFiles("../HTML/hangman.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			menu.Execute(w, nil)
			return
		}
		hangman_data := Hangman{
			Difficulty: r.FormValue("difficulty"),
			Diff_choosed: true,
		}
		
		menu.Execute(w, hangman_data)

		word_to_find := hangman_data.ReadFile(hangman_data.Difficulty)
		hangman_data = Hangman{
			Word: word_to_find,
			User_input: r.FormValue("user_input"),
			Attempt: 10,
		}
		hangman_data.CreateHidden()
		hangman_data.Reveal()
		hangman.Execute(w, hangman_data)
		
	})

	http.ListenAndServe(":8080", nil)
}