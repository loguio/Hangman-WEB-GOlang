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
	Title           string
	Letter          string
	URLpendu        string
	NumberOfAttemps int
	SearchWord      string
	Word            string
	WordFind        bool
}

var tabURL = []string{
	"",
	"https://i.goopics.net/bqidl2.png",
	"https://i.goopics.net/rojgcs.png",
	"https://i.goopics.net/igktwf.png",
	"https://i.goopics.net/a8sgek.png",
	"https://i.goopics.net/a9mnww.png",
	"https://i.goopics.net/lif1mq.png",
	"https://i.goopics.net/hy8no6.png",
	"https://i.goopics.net/h377xl.png",
	"https://i.goopics.net/rswcvd.png",
	"https://i.goopics.net/uiwsjx.png"}

var WordFind bool

func main() {

	website()
	fileServer := http.FileServer(http.Dir("assets")) //application du CSS qui sera en fichier statique
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur

}
func website() {

	game.NumberOfAttemps = 10
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		Error404()
	} else {

		list_letter := ""

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			data := Page{"Hangman-Web ", list_letter, "", game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayAnswer), WordFind}
			if r.Method == "POST" {
				if r.FormValue("restart") == "Restart" {
					Restart()
				} else {
					game.Lettre = r.FormValue("letter") //recupere la valeur letter du formulaire (html)
					game.Game()
					if string(game.ArrayAnswer) == string(game.ArrayInit) {
						WordFind = true
					}
					if game.Lettre != "" && game.Lettre != " " {
						list_letter += game.Lettre
						list_letter += ", "
					}
					if game.NumberOfAttemps == 0 {
						data = Page{"Hangman-Web ", list_letter, "https://th.bing.com/th/id/OIP.kq55Q_4YugKpMd67w_YD3wHaFO?pid=ImgDet&rs=1", game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind}
					} else {
						data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer), string(game.ArrayInit), WordFind}
					}
				}
				tmpl.ExecuteTemplate(w, "index", data)
			} else if r.Method == "GET" {
				fmt.Println("pas de POST effectué")
				tmpl.ExecuteTemplate(w, "index", data)
			} else {
				Error501()
			}

		})
	}
}

func Restart() {
	WordFind = false
	game.InitString = hangman.GetRandomWord()
	game.RandomWord = string(game.InitString[rand.Intn(len(game.InitString))])
	game.ArrayAnswer = hangman.InitArray(game.RandomWord)
	game.ArrayInit = []rune(game.RandomWord)
	game.NumberOfAttemps = 10

}

func Error404() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
		tmpl, _ := template.ParseFiles("./templates/error404.gohtml")
		fmt.Println("le serveur est en cours d'éxécution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error404", nil)
	})
	http.ListenAndServe("localhost:3000", nil)

}
func Error501() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("./tempaltes/error501.gohtml")
		fmt.Println("le serveur est en cours d'execution à l'adresse localhost:3000")
		tmpl.ExecuteTemplate(w, "error501", nil)
	})
	http.ListenAndServe("localhost:3000", nil)
}
