package hangman

func UserLetter(Lettre string) bool { // fonction qui demande une lettre au joueur et la vérifie
	SpecialChar := []string{"é", "è", "ê", "à", "â", "-", "ù", "û", "ç"}
	for i := 0; i < len(SpecialChar); i++ {
		if Lettre == SpecialChar[i] {
			return true
		}
	}
	runeArrayLetter := []rune(Lettre)
	if runeArrayLetter[0] >= 'a' && runeArrayLetter[0] <= 'z' {
		runeArrayLetter[0] -= 32
	}
	if runeArrayLetter[0] >= 'A' && runeArrayLetter[0] <= 'Z' {
		return true
	} else {
		return false
	}
}
