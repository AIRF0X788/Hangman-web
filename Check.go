package piscine

// Vérification
func Check(a int) string {
	w := ""
	if a != 10 {
		if Read("killjose.txt")[9] == byte(13) {
			w = (string(Read("killjose.txt")[(9-a)*79 : (9-a)*79+77]))
		} else {
			w = (string(Read("killjose.txt")[(9-a)*71 : (9-a)*71+70]))
		}
	}
	return w
}
