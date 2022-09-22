package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	hangman "./hangman-classic"
	game "./hangman-classic/test"
)

type Page struct { //type Page :=> donnée envoyé a la page HTML  et utilisé lors de l'affichage grâce au package "html/template"
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

var tabURL = []string{ // tableau avec les adresses pour les images de l'état du pendu
	"",
	"./static/pictures/pendu_10.png",
	"./static/pictures/pendu_9.png",
	"./static/pictures/pendu_8.png",
	"./static/pictures/pendu_7.png",
	"./static/pictures/pendu_6.png",
	"./static/pictures/pendu_5.png",
	"./static/pictures/pendu_4.png",
	"./static/pictures/pendu_3.png",
	"./static/pictures/pendu_2.png",
	"./static/pictures/pendu_1.png",
}

var WordFind bool
var tabletter []string
var list_letter string
var data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat}

func Website() {
	Error404()
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println("beug")
		Error500()
	} else {
		http.HandleFunc("/Hangman-Web", func(w http.ResponseWriter, r *http.Request) { //crée une page
			if r.Method == "POST" {
				if r.FormValue("restart") == "Restart" {
					Restart()
					data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat} //actualisation de la variable data
				} else {
					game.Lettre = r.FormValue("letter") //recupere la valeur letter du formulaire (html)
					same := false
					for i := 0; i < len(tabletter); i++ { //regarde si la lettre a déjà été utlisé
						if game.Lettre == tabletter[i] {
							same = true
						}
					}
					if string(game.ArrayAnswer) == string(game.ArrayInit) { //tets si le mot est trouvé
						WordFind = true
					}
					if game.Lettre != "" && game.Lettre != " " && !same { //si les lettres sont identiques
						tabletter = append(tabletter, game.Lettre)
						list_letter += game.Lettre
						list_letter += ", "
						if game.Game() == false { //si probleme lors de l'execution du programme du hangman
							Error500() // execute le code d'erreur 500
							return
						}
					}
					data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind, tabletter, game.LetterGoodFormat} // actualisation de data
				}
			} else if r.Method == "GET" {
				fmt.Println(r.Method)
			} else {
				fmt.Println(r.Method)
				//si méthode du serveur différent de POST et GET
				Error501() //execution du code d'erreur 501
			}
			tmpl.ExecuteTemplate(w, "index", data) //execution de la template "index" avec les données
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
	http.HandleFunc("/Hangman-Web", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/error500.gohtml")
		fmt.Println("le serveur est en cours d'execution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error500", nil)
	})
}

//Fonction erreur 404
//erreur d'URL
func Error404() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/error404.gohtml")
		tmpl.ExecuteTemplate(w, "error404", nil)
	})
}

//Fonction erreur 501
func Error501() {
	http.HandleFunc("/Hangman-Web", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./templates/error501.gohtml")
		fmt.Println("le serveur est en cours d'execution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error501", nil)
	})
}
