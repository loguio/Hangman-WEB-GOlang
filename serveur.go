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
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			tmpl, err = template.ParseFiles("./templates/error404.gohtml")
			tmpl.ExecuteTemplate(w, "error404", nil)
		})
		http.ListenAndServe("localhost:3000", nil)

	} else {
		type Page struct {
			Title  string
			Letter string
		}
		InitString := hangman.GetRandomWord() //met chaque mot de la liste de mot dans un tableau de string
		game.RandomWord = (InitString[rand.Intn(len(InitString))])
		game.ArrayInit = []rune(game.RandomWord)
		game.ArrayAnswer = hangman.InitArray(game.RandomWord)
		game.NumberOfAttemps = 11
		list_letter := ""

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			fmt.Println(r.Header.Get("User-Agent"))
			data := Page{"Hangman-Web ", list_letter}
			if r.Method == "POST" {
				lettre := r.FormValue("letter") //recupere la valeur letter du formulaire ( html)
				list_letter += lettre
				list_letter += ", "
				data = Page{"Hangman-Web ", list_letter}
				fmt.Println(lettre)
				tmpl.ExecuteTemplate(w, "index", data)
			} else if r.Method == "GET" {
				fmt.Println("pas de POST effectué")
				tmpl.ExecuteTemplate(w, "index", data)
			} else {
				tmpl, err = template.ParseFiles("./templates/error501.gohtml")
				tmpl.ExecuteTemplate(w, "error501", nil)
			}

		})

		fileServer := http.FileServer(http.Dir("assets"))
		http.Handle("/static/", http.StripPrefix("/static/", fileServer))

		http.ListenAndServe("localhost:3000", nil)
	}

}

//TODO faire 3 partie pour chaque demande dans la base de donnée, une partiesi c'est GET , une autre si c'est POST et une autre qui differe des autres parties ( != GET || Post)
