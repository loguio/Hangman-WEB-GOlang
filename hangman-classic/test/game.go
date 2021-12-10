package game

import (
	"fmt"

	hangman ".."
)

var Lettre string
var ArrayAnswer []rune
var ArrayInit []rune
var RandomWord string
var NumberOfAttemps int

func Game(Lettre string) {
	InitString := hangman.GetRandomWord() //met chaque mot de la liste de mot dans un tableau de string
	if InitString == nil {                //si le fichier n'est pas trouvé et donc il n'y a pas de mot disponible pour le jeu, arret du jeu
		return
	}
	fmt.Println(string(ArrayAnswer))                //affiche les lettres déjà trouver dans le tableau
	hangman.Compare(Lettre, ArrayInit, ArrayAnswer) //regarde si la lettre est contenu dans le mot
	compare := hangman.Compare(Lettre, ArrayInit, ArrayAnswer)
	if compare == false { //si la lettre n'est pas contenue
		NumberOfAttemps--
	} else if compare == true { //si la lettre est contenue dans le mot, remplacer dans le tableau de réponse le caractère _ par la lettre mise par le joueur
		arrayRune := []rune(Lettre)
		pos := []int(nil)
		for i := 0; i < len(ArrayInit); i++ {
			if arrayRune[0] == ArrayInit[i] {
				pos = append(pos, i)
			}
			for i := 0; i < len(pos); i++ {
				ArrayAnswer[pos[i]] = ArrayInit[pos[i]]
			}
		}
	}

}

//TODO revoir le programme hangman-classic avec le bon cheminement de la lettrre et tout les trucs qu'ils faut return ...
