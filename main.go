package main

import (
	"fmt"
	"os"
)

func Reverse(a []rune) []rune {
	for i := 0; i < len(a)/2; i++ {
		j := len(a) - i - 1
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func IsVowel(letter string) bool {
	harf, ok := Haruf[letter]
	if !ok {
		return false
	}
	if harf.Tashkeel {
		return true
	}
	return false
}

func GetNextLetter(word string, i int) (Harf, int, bool) {
	wordLen := len(word)
	letter := string(word[i])
	lastLetter := false
	if i == wordLen-1 {
		i++
		lastLetter = true
	} else {
		nextLetter := string(word[i+1])
		if letter == "a" {
			if nextLetter == "a" || nextLetter == "e" || nextLetter == "N" {
				letter = word[i : i+2]
				i += 2
			} else {
				i++
			}
		} else if nextLetter == "h" {
			if letter == "s" || letter == "t" || letter == "k" || letter == "d" || letter == "g" {
				letter = word[i : i+2]
				i += 2
			} else {
				i++
			}
		} else if letter == "'" {
			if nextLetter == "a" || nextLetter == "y" || nextLetter == "w" {
				letter = word[i : i+2]
				i += 2
			} else {
				i++
			}
		} else {
			i++
		}
	}
	if i == wordLen-1 || i == wordLen-2 {
		nextLetter := string(word[i:wordLen])
		if IsVowel(nextLetter) {
			lastLetter = true
		}
	}
	return Haruf[letter], i, lastLetter
}

func Transliterate(englishWords []string) string {
	var arabicWords []rune
	for _, word := range englishWords {
		var lastConnected bool
		i := 0
		wordLen := len(word)
		var harf Harf
		var last bool
		for i < wordLen {
			oldI := i
			harf, i, last = GetNextLetter(word, i)
			var letterToAdd rune
			if oldI == 0 {
				letterToAdd = harf.Initial
			} else if last {
				if !lastConnected {
					letterToAdd = harf.Isolated
				} else {
					letterToAdd = harf.Final
				}
			} else if !lastConnected {
				letterToAdd = harf.Initial
			} else {
				letterToAdd = harf.Median
			}
			if !harf.Tashkeel {

				lastConnected = harf.ConnectNext
			}
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
