package generator

import (
	"log"
	"time"
)

// removed I, i and l to avoid unnecessary confusion
var customAlphabet = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

// Translates a uint to a string using the provided alphabet.
func utoa(n uint64, alphabet []byte) string {
	log.Println(n) // debugging
	bits := []byte{}
	for n > 0 {
		bits = append(bits, alphabet[n%uint64(len(alphabet)-1)])
		n = n / uint64(len(alphabet))
	}
	// reverse the rune order
	l := len(bits)
	for i := 0; i < l/2; i++ {
		bits[i], bits[l-1-i] = bits[l-1-i], bits[i]
	}
	return string(bits)
}

func ShortUrl() string {
	u := time.Now().UnixNano()

	return utoa(uint64(u), customAlphabet)
}
