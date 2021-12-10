package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	tmpl, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		fmt.Println("erreur ouverture template")
		return
	}
	type Page struct {
		Title  string
		Letter string
	}
	//str := []string{""}
	i := 0
	list_letter := ""
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { //crée une page

		fmt.Println(r.Header.Get("User-Agent"))
		data := Page{"Hangman-Web ", list_letter}

		if r.Method == "POST" {
			i++
			lettre := r.FormValue("letter") //recupere la valeur letter du formulaire ( html)
			list_letter += lettre
			list_letter += ", "
			data = Page{"Hangman-Web ", list_letter}
			fmt.Println(lettre)
		} else if r.Method == "GET" {
			fmt.Println("pas de POST effectué")
		}
		tmpl.ExecuteTemplate(w, "index", data)
	})

	fileServer := http.FileServer(http.Dir("assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:3000", nil)
}

//TODO faire 3 partie pour chaque demande dans la base de donnée, une partiesi c'est GET , une autre si c'est POST et une autre qui differe des autres parties ( != GET || Post)
