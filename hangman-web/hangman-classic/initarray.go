package hangman

import (
	"math/rand"
	"time"
)

func InitArray(randomWord string) []rune { // créer le tableau de réponse avec des lettres déjà mise
	arrayInit := []rune(randomWord)
	arrayAnswer := []rune(nil)
	for i := 0; i < len(arrayInit); i++ {
		arrayAnswer = append(arrayAnswer, '_')
	}
	letterReveal := (len(arrayInit) / 2) - 1
	for i := 0; i < letterReveal; i++ {
		rand.Seed(time.Now().UnixNano())
		pos := rand.Intn(len(arrayInit))
		letter := arrayInit[pos]
		arrayAnswer[pos] = letter
		for j := 0; j < len(arrayInit); j++ {
			if letter == arrayInit[j] {
				arrayAnswer[j] = letter
			}
		}
	}
	return arrayAnswer
}
