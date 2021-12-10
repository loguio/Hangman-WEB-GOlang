package hangman

import (
	"fmt"
	"io/ioutil"
)

func Jose(numberOfAttemps int) { //fonction qui affiche la position du pendu par rapport au nombre de vie restante
	var arrayJose []byte
	initNumberOfLife := 9
	pos := initNumberOfLife - numberOfAttemps
	content, err := ioutil.ReadFile("josé.txt")

	if err != nil {
		fmt.Println("[fichier non trouvé]")
	} else {
		for i := 0; i < 71; i++ {
			arrayJose = append(arrayJose, content[i+(71*pos)])

		}
		josé := (string(arrayJose))
		fmt.Println(josé)

	}
}
