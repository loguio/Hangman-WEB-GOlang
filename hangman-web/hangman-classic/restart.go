package hangman

import "fmt"

func Restart() bool { //fonction pour rejouer ou non
	var reponse string
	fmt.Println("Voulez vous rejouer ? Y pour rejouer , N pour arreter")
	fmt.Scan(&reponse)
	if reponse == "Y" {
		return true
	} else if reponse == "N" {
		return false
	} else {
		fmt.Println("Veuillez rentrer soit Y pour rejouer, soit N pour arreter")
		fmt.Scan(&reponse)
	}
	return false
}
