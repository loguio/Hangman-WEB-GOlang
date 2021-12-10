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
	i := 0
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			tmpl, err = template.ParseFiles("./templates/error404.gohtml")
			tmpl.ExecuteTemplate(w, "error404", nil)
		})
		http.ListenAndServe("localhost:3000", nil)

	} else {
		type Page struct {
			Title    string
			Letter   string
			URLpendu string
		}
		InitString := hangman.GetRandomWord() //met chaque mot de la liste de mot dans un tableau de string
		game.RandomWord = (InitString[rand.Intn(len(InitString))])
		game.ArrayInit = []rune(game.RandomWord)
		game.ArrayAnswer = hangman.InitArray(game.RandomWord)
		game.NumberOfAttemps = 11
		list_letter := ""
		var tabURL = []string{
			"",
			"https://i.goopics.net/uiwsjx.png",
			"https://i.goopics.net/rswcvd.png",
			"https://i.goopics.net/h377xl.png",
			"https://i.goopics.net/hy8no6.png",
			"https://i.goopics.net/lif1mq.png",
			"https://i.goopics.net/a9mnww.png",
			"https://i.goopics.net/a8sgek.png",
			"https://i.goopics.net/igktwf.png",
			"https://i.goopics.net/rojgcs.png",
			"https://i.goopics.net/bqidl2.png"}
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
			data := Page{"Hangman-Web ", list_letter, ""}
			if r.Method == "POST" {
				lettre := r.FormValue("letter") //recupere la valeur letter du formulaire ( html)
				if lettre != "" && lettre != " " {
					list_letter += lettre
					list_letter += ", "
				}
				if i > 10 {
					data = Page{"Hangman-Web ", list_letter, "https://th.bing.com/th/id/OIP.kq55Q_4YugKpMd67w_YD3wHaFO?pid=ImgDet&rs=1"}
				} else {
					data = Page{"Hangman-Web ", list_letter, tabURL[i]}
				}
				i++
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
