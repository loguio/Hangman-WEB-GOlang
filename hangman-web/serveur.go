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

	InitString := hangman.GetRandomWord() //met chaque mot de la liste de mot dans un tableau de string
	game.RandomWord = (InitString[rand.Intn(len(InitString))])
	game.ArrayInit = []rune(game.RandomWord)
	game.ArrayAnswer = hangman.InitArray(game.RandomWord)
	game.NumberOfAttemps = 11

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
		fmt.Println(r.Header.Get("User-Agent"))
		data := Page{string(game.RandomWord), "Contenu"}
		tmpl.ExecuteTemplate(w, "index", data)
		game.Lettre = r.FormValue("letter") //recupere la valeur letter du formulaire ( html)
		if r.Method == "POST" {
			game.Game(game.Lettre)
			fmt.Fprintln(w, string(game.ArrayAnswer))
		} else if r.Method == "GET" {
			fmt.Println("GET")
		}
	})

	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:3000", nil)
}

//TODO Une fonction avec des switch pour l'état du pendu
//TODO gestion d'erreur 404(erreur connection serveur) et 501(parti Backend beug)
