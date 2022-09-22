package hangman

func Compare(Letter string, ArrayInit []rune, ArrayAnswer []rune) bool {
	arrayRune := []rune(Letter)
	for i := 0; i < len(ArrayInit); i++ {
		if arrayRune[0] == ArrayInit[i] {
			return true
		}
	}
	return false
}
