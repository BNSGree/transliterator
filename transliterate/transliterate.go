package transliterate

import (
	"strings"
)

const (
	latin = "abcdefghijklmnopqrstuvwxyz"
	runic = "ᚨᛒᚲᛞᛖᚠᚷᚺᛁᛃᚴᛚᛗᚾᛟᛈᛩᚱᛋᛏᚢᚡᚹᛪᚤᛉ"
)

type Alphabet int

const (
	Latin Alphabet = iota
	Runic
)

var alphabets = map[Alphabet]string{
	Latin: latin,
	Runic: runic,
}

func transliterateRune(from Alphabet, to Alphabet, r rune) rune {
	i := 0
	for _, r2 := range alphabets[from] {
		if r2 == r {
			return []rune(alphabets[to])[i]
		}
		i++
	}
	return r
}

func Translate(from Alphabet, to Alphabet, value string) string {
	return strings.Map(func(r rune) rune { return transliterateRune(from, to, r) }, value)
}
