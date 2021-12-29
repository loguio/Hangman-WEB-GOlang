package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("assets")) //Envoie des fichiers aux serveurs (CSS, sons, images)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
	Website()

	fmt.Println("le serveur est en cours d'éxécution a l'adresse localhost:3000")
	http.ListenAndServe("localhost:3000", nil) //lancement du serveur

}
