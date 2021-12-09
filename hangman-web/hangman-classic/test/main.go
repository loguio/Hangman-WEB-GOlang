package game

import (
	"fmt"

	hangman ".."
)

const (
	message_erreur = "erreur lors de l'ouverture d'un fichier "
)

var Lettre string
var ArrayAnswer []rune
var ArrayInit []rune
var RandomWord string

func Game(Lettre string) {
	fmt.Print(Lettre)
	numberOfAttemps := 10                 //nombre d'essais
	InitString := hangman.GetRandomWord() //met chaque mot de la liste de mot dans un tableau de string
	if InitString == nil {                //si le fichier n'est pas trouvé et donc il n'y a pas de mot disponible pour le jeu, arret du jeu
		return
	}
	for string(ArrayInit) != string(ArrayAnswer) { //répéter tant que le tableau de rune Initial(donc avec le mot en entier), n'est pas le même que le tableau de réponse du joueur
		fmt.Println(string(ArrayAnswer))                              //affiche les lettres déjà trouver dans le tableau
		letter := hangman.UserLetter()                                //demande une lettre au joueur
		hangman.Compare(letter, ArrayInit, ArrayAnswer)               //regarde si la lettre est contenu dans le mot
		if hangman.Compare(letter, ArrayInit, ArrayAnswer) == false { //si la lettre n'est pas contenue
			numberOfAttemps--
			if numberOfAttemps > 0 {
				hangman.Jose(numberOfAttemps) //affichage de la position du pendu en fonction du nmbre de tentatives restantes
			}
		} else if hangman.Compare(letter, ArrayInit, ArrayAnswer) == true { //si la lettre est contenue dans le mot, remplacer dans le tableau de réponse le caractère _ par la lettre mise par le joueur
			arrayRune := []rune(letter)
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
}

//TODO revoir le programme hangman-classic avec le bon cheminement de la lettrre et tout les trucs qu'ils faut return ...
