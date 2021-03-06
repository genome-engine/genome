package identifier

var alphabet = map[byte]int{
	'A': 1, 'a': 2, 'B': 3, 'b': 4, 'C': 5, 'c': 6,
	'D': 7, 'd': 8, 'E': 9, 'e': 10, 'F': 11, 'f': 12,
	'G': 13, 'g': 14, 'H': 15, 'h': 16, 'I': 17, 'i': 18,
	'J': 19, 'j': 20, 'K': 21, 'k': 22, 'L': 23, 'l': 24,
	'M': 25, 'm': 26, 'N': 27, 'n': 28, 'O': 29, 'o': 30,
	'P': 31, 'p': 32, 'Q': 33, 'q': 34, 'R': 35, 'r': 36,
	'S': 37, 's': 38, 'T': 39, 't': 40, 'U': 41, 'u': 42,
	'V': 43, 'v': 44, 'W': 45, 'w': 46, 'X': 47, 'x': 48,
	'Y': 49, 'y': 50, 'Z': 51, 'z': 52, '0': 53, '1': 54,
	'2': 55, '3': 56, '4': 57, '5': 58, '6': 59, '7': 60,
	'8': 61, '9': 62, '_': 63,
}

func GenerateID(objectName string) int {
	if objectName == "" {
		return 0
	}
	var control, last int

	for _, letter := range objectName {
		if num, ok := alphabet[byte(letter)]; ok {
			control += num + last
			last = num
		}
	}

	return control
}
