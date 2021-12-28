package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	hangman "./hangman-classic"
	game "./hangman-classic/test"
)

type Page struct {
	Title            string
	Letter           string
	URLpendu         string
	NumberOfAttemps  int
	SearchWord       string
	Word             string
	WordFind         bool
	Tabletter        []string
	LetterGoodFormat bool
}

var tabURL = []string{
	"",
	"./picture/pendu_10.png",
	"./picture/pendu_9.png",
	"./picture/pendu_8.png",
	"./picture/pendu_7.png",
	"./picture/pendu_6.png",
	"./picture/pendu_5.png",
	"./picture/pendu_4.png",
	"./picture/pendu_3.png",
	"./picture/pendu_2.png",
	"./picture/pendu_1.png",
}

var WordFind bool
var tabletter []string
var list_letter string
var data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat} // actualisation de Data

func Website() {
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println("beug")
		Error404()
	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			data := Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat}
			if r.Method == "POST" {
				fmt.Println("POST")
				if r.FormValue("restart") == "Restart" {
					Restart()
					data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat}
				} else {
					game.Lettre = r.FormValue("letter") //recupere la valeur letter du formulaire (html)
					same := false
					for i := 0; i < len(tabletter); i++ {
						if game.Lettre == tabletter[i] {
							same = true
						}
					}
					if string(game.ArrayAnswer) == string(game.ArrayInit) {
						WordFind = true
					}
					if game.Lettre != "" && game.Lettre != " " && !same {
						tabletter = append(tabletter, game.Lettre)
						list_letter += game.Lettre
						list_letter += ", "
						if game.Game() == false { //si probleme lors de l'execution du programme du hangman
							Error500()
							return
						}
					}
					data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat} // actualisation de Data
				}
				tmpl.ExecuteTemplate(w, "index", data)
			} else if r.Method == "GET" {
				fmt.Println("GET")
			} else {
				Error501()
			}

		})
	}
}

//Fonction qui sert au restart
//Reinstaure des nouvelles valeurs et réinitialise le nombres de vies
func Restart() {
	tabletter = nil
	WordFind = false
	game.InitString = hangman.GetRandomWord()
	game.RandomWord = string(game.InitString[rand.Intn(len(game.InitString))])
	game.ArrayInit = []rune(game.RandomWord)
	game.ArrayAnswer = hangman.InitArray(game.RandomWord)
	game.NumberOfAttemps = 10
}

//Fonction erreur 500
func Error500() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tempaltes/error500.gohtml")
		fmt.Println("le serveur est en cours d'execution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error500", nil)
	})
	http.ListenAndServe("localhost:3000", nil)
}

//Fonction erreur 404
func Error404() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/error404.gohtml")
		fmt.Println("le serveur est en cours d'éxécution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error404", nil)
	})
	http.ListenAndServe("localhost:3000", nil)
}

//Fonction erreur 501
func Error501() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tempaltes/error501.gohtml")
		fmt.Println("le serveur est en cours d'execution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error501", nil)
	})
	http.ListenAndServe("localhost:3000", nil)
}
