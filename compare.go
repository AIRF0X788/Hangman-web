package piscine

type Sauvegarde struct {
	Attempts int
	Pick     string
	UdScore  []rune
}

var Repetition []string

// Vérification
func Compare(u []rune, a int, p string, letter string) ([]rune, int) {

	d := false

	if len(letter) == 1 { // Letter
		for v := range p {
			if p[v] == letter[0] && u[v] == '_' {
				u[v] = rune(p[v])
				d = true
			}
		}
	}

	if letter == p { // word
		u = []rune(p)
		d = true
	}

	repeat := 0
	for v := range Repetition {
		if Repetition[v] == ToUpper(letter) {
			repeat++
			d = true
		}
	}

	if d {
	} else {
		a--
		if len(letter) > 1 && a > 0 {
			a--
		}
	}

	Repetition = append(Repetition, letter)
	return u, a
}
