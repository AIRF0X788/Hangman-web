package piscine

func Accent(a string) string {
	b := []rune(a)
	w := []rune{}
	for v := range b {
		if b[v] > 231 && b[v] < 236 {
			w = append(w, 'e')
		} else if b[v] > 223 && b[v] < 231 {
			w = append(w, 'a')
		} else if b[v] > 235 && b[v] < 240 {
			w = append(w, 'i')
		} else if b[v] == 231 {
			w = append(w, 'c')
		} else if b[v] > 241 && b[v] < 247 {
			w = append(w, 'o')
		} else if b[v] > 248 && b[v] < 253 {
			w = append(w, 'u')
		} else {
			w = append(w, b[v])
		}
	}
	return string(w)
}
