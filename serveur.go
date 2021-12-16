package main

import (
	"fmt"
	"net/http"
)

func main() {

	picturesFolder := http.FileServer(http.Dir("pictures"))
	http.Handle("/picture/", http.StripPrefix("/picture/", picturesFolder))

	Website()
	fileServer := http.FileServer(http.Dir("assets")) //application du CSS qui sera en fichier statique
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur

}
