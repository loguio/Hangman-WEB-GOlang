package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	hangman "./hangman-classic"
	game "./hangman-classic/test"
)

func main() {
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println("erreur ouverture template")
		return
	}
	type Page struct {
		Title   string
		Content string
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
		fmt.Println(r.Header.Get("User-Agent"))
		data := Page{"Titre 1 ", "Contenu"}
		tmpl.ExecuteTemplate(w, "index", data)
		game.Lettre = r.FormValue("letter")   //recupere la valeur letter du formulaire ( html)
		InitString := hangman.GetRandomWord() //met chaque mot de la liste de mot dans un tableau de string
		game.RandomWord = (InitString[rand.Intn(len(InitString))])
		game.ArrayInit = hangman.InitArray(game.RandomWord)
		game.ArrayAnswer = hangman.InitArray(game.RandomWord)
		if r.Method == "POST" {
			//numberOfAttemps := game.Game(game.Lettre)
			//fmt.Println(numberOfAttemps)
			fmt.Println(game.Lettre)
		} else if r.Method == "GET" {
			fmt.Println("pas de POST effectué")
		}
	})

	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:3000", nil)
}

//TODO faire 3 partie pour chaque demande dans la base de donnée, une partiesi c'est GET , une autre si c'est POST et une autre qui differe des autres parties ( != GET || Post)
//TODO Une fonction avec des switch pour l'état du pendu
//TODO Image pour pendu
