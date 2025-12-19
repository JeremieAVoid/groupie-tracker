package groupie

// Écrire 0 dans "aPlacer" permet de suprimé les caractères "aRetirer".
func Remplacer(texte string, aRetirer rune, aPlacer rune) string {
	total := ""
	for i := 0; i < len(texte); i++ {
		if texte[i] != byte(aRetirer) {
			total += string(texte[i])
		} else {
			if aPlacer != 0 {
				total += string(aPlacer)
			}
		}
	}
	return total
}
