package main

import (
	"fmt"
	"html/template"
	"net/http"

	game "./hangman-classic/test"
)

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
	"https://th.bing.com/th/id/R.465dad455068168abb0a4e97ece5dd25?rik=eZV0J8IydjJ6xA&pid=ImgRaw&r=0&sres=1&sresct=1"}

func main() {
	// restart := false
	game.NumberOfAttemps = 10
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	// string(game.ArrayInit)
	type Page struct {
		Title           string
		Letter          string
		URLpendu        string
		NumberOfAttemps int
		SearchWord      string
	}
	list_letter := ""
	tabletter := []string{}
	if err != nil {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			tmpl, err = template.ParseFiles("./templates/error404.gohtml")
			tmpl.ExecuteTemplate(w, "error404", nil)
		})
		http.ListenAndServe("localhost:3000", nil)

	} else {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page
			if string(game.ArrayAnswer) == string(game.ArrayInit) {
				fmt.Println("oue")
			}
			data := Page{"Hangman-Web ", list_letter, "", game.NumberOfAttemps, string(game.ArrayAnswer)}
			if r.Method == "POST" {
				game.Lettre = r.FormValue("letter") //recupere la valeur letter du formulaire (html)
				same := false
				for i := 0; i < len(tabletter); i++ {
					if game.Lettre == tabletter[i] {
						same = true
					}
				}
				if game.Lettre != "" && game.Lettre != " " && !same {
					list_letter += game.Lettre
					list_letter += ", "
					tabletter = append(tabletter, game.Lettre)
					game.Game()
				}
				if game.NumberOfAttemps == 0 {
					data = Page{"Hangman-Web ", list_letter, "https://th.bing.com/th/id/OIP.kq55Q_4YugKpMd67w_YD3wHaFO?pid=ImgDet&rs=1", game.NumberOfAttemps, string(game.ArrayAnswer)}
					game.NumberOfAttemps = 10
				} else {
					data = Page{"Hangman-Web ", list_letter, tabURL[game.NumberOfAttemps], game.NumberOfAttemps, string(game.ArrayAnswer)}
				}
				tmpl.ExecuteTemplate(w, "index", data)
			} else if r.Method == "GET" {
				tmpl.ExecuteTemplate(w, "index", data)
			} else if r.Method == "RESTART" {
				game.NumberOfAttemps = 10
				list_letter = ""
				tmpl.ExecuteTemplate(w, "index", data)
				main()
			} else {
				tmpl, err = template.ParseFiles("./templates/error501.gohtml")
				tmpl.ExecuteTemplate(w, "error501", nil)
			}
		})

		fileServer := http.FileServer(http.Dir("assets")) //application du CSS qui sera en fichier statique
		http.Handle("/static/", http.StripPrefix("/static/", fileServer))

		fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
		http.ListenAndServe("localhost:8080", nil)
	}
}
