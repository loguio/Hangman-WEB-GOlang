package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GetRandomWord() []string { //fonction qui choisi un mot aléatoire dans une des 3 listes
	var tableauMot []string
	min := 1
	max := 3
	rand.Seed(time.Now().UnixNano())
	whichWord := (rand.Intn(max-min+1) + min) //choisi un nombre aléatoire entre 1 et 3( pour les listes de mots )
	words := "words" + strconv.Itoa(whichWord) + ".txt"

	f, err := os.Open(words)

	if err != nil { //gestion de l'erreur
		fmt.Println("Fichier introuvable")
		return tableauMot
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		rand.Seed(time.Now().UnixNano())
		tableauMot = append(tableauMot, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Fichier introuvable")
	}
	return tableauMot
}
