package main

import (
	"fmt"
	"os"
	"strings"
)

func Reverse(a []rune) []rune {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func Transliterate(englishWords []string) string {
	var arabicWords []rune
	for _, word := range englishWords {
		englishLetters := strings.Split(word, "-")
		var lastConnected bool
		wordLen := len(englishLetters)
		for i, letter := range englishLetters {
			harf := letters[letter]
			var letterToAdd rune
			if i == 0 {
				letterToAdd = harf.Initial
			} else if i == wordLen-1 {
				letterToAdd = harf.Final
			} else if !lastConnected {
				letterToAdd = harf.Initial
			} else {
				letterToAdd = harf.Median
			}
			lastConnected = harf.ConnectNext
			arabicWords = append(arabicWords, letterToAdd)
		}
		arabicWords = append(arabicWords, space)
	}
	correctOrder := Reverse(arabicWords)
	return string(correctOrder[1:])
}
func main() {
	englishWords := os.Args[1:]
	fmt.Println(Transliterate(englishWords))
}
