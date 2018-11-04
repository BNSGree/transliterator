package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/BNSGree/transliterator/transliterate"
)

func main() {
	var target string
	flag.StringVar(&target, "target", "latin", "What you want your end product to be. (latin or runic)")

	flag.Parse()

	var from transliterate.Alphabet
	var to transliterate.Alphabet
	switch target {
	case "latin":
		from = transliterate.Runic
		to = transliterate.Latin
	case "runic":
		from = transliterate.Latin
		to = transliterate.Runic
	default:
		fmt.Println("Unexpected item in the bagging area.")
		flag.Usage()
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Please place your item in the bagging area.")
		flag.Usage()
		return
	}

	value := strings.Join(flag.Args(), " ")

	fmt.Println(transliterate.Translate(from, to, value))
}
