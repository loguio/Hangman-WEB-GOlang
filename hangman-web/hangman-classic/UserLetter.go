package hangman

import (
	"fmt"
)

func UserLetter() string { // fonction qui demande une lettre au joueur et la vérifie
	var Lettre string
	if Lettre == "é" || Lettre == "à" || Lettre == "â" || Lettre == "è" || Lettre == "ê" || Lettre == "-" || Lettre == "ç" || Lettre == "ù" { // caractère spéciaux
		return Lettre
	}
	if len(Lettre) != 1 { //si le joueur rentre plus d'une lettre
		fmt.Println("veuillez mettre qu'une seule lettre")
		UserLetter()
	} else {
		runeArrayLetter := []rune(Lettre)
		if runeArrayLetter[0] >= 'a' && runeArrayLetter[0] <= 'z' {
			runeArrayLetter[0] -= 32
		}
		if !(runeArrayLetter[0] >= 65 && runeArrayLetter[0] <= 90) {
			fmt.Println("Veuillez rentrer un caractère valide")
			UserLetter()
		} else {
			return Lettre
		}
	}
	return Lettre
}
