package main

import (
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
	"fmt"
)

func (h *Hangman) ReadFile(difficulty string) {
	rand.Seed(time.Now().UnixNano())

	randomNumber := 0
	if difficulty == "Facile"{
		min := 1
		max := 12
		randomNumber = min+rand.Intn(max-min+1)
	} else if difficulty == "Modéré" {
		min := 13
		max := 25
		randomNumber = min+rand.Intn(max-min+1)
	} else if difficulty == "Difficile"{
		min := 26
		max := 36
		randomNumber = min+rand.Intn(max-min+1)
	} else {
		fmt.Println("Il y a un problème, difficulté vaut: ",)
	}

	data, err := ioutil.ReadFile("../Dictionnaire/words.txt")
	content := string(data)
	words := strings.Split(content, "\n")
	if err != nil {
		fmt.Println(err)

	} else {
		for i := 0; i < randomNumber; i++ {
			h.Word = words[i]
		}
	}
}