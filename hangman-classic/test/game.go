package game

import (
	"fmt"
	"math/rand"

	hangman ".."
)

var InitString []string = hangman.GetRandomWord()
var RandomWord string = string(InitString[rand.Intn(len(InitString))])
var Lettre string
var ArrayAnswer []rune = hangman.InitArray(RandomWord)
var ArrayInit []rune = []rune(RandomWord)
var NumberOfAttemps int = 10

func Game() {
	if InitString == nil { //si le fichier n'est pas trouvé et donc il n'y a pas de mot disponible pour le jeu, arret du jeu
		fmt.Println("erreur obtention mot aléatoire")
		return
	}
	hangman.Compare(Lettre, ArrayInit, ArrayAnswer) //regarde si la lettre est contenu dans le mot
	compare := hangman.Compare(Lettre, ArrayInit, ArrayAnswer)
	if compare == false { //si la lettre n'est pas contenue
		NumberOfAttemps--
		fmt.Println(NumberOfAttemps)
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

//TODO renvoye le nb de vie actualisé
