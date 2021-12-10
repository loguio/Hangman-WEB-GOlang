package hangman

func Compare(letter string, arrayInit []rune, arrayAnswer []rune) bool {
	arrayRune := []rune(letter)
	for i := 0; i < len(arrayInit); i++ {
		if arrayRune[0] == arrayInit[i] {
			return true
		}
	}
	return false
}
