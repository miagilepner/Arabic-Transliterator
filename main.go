package main

import (
	"fmt"
	"os"
)

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

func TwoLetterUnit(word string, i int) (string, int) {
	letter := string(word[i])
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
	return letter, i
}

func GetNextLetter(word string, i int) (Harf, int, bool) {
	wordLen := len(word)
	var letter string
	lastLetter := false
	if i == wordLen-1 {
		i++
		lastLetter = true
	} else {
		letter, i = TwoLetterUnit(word, i)
	}
	if i == wordLen-1 || i == wordLen-2 {
		nextLetter := string(word[i:wordLen])
		if IsVowel(nextLetter) {
			lastLetter = true
		}
	}
	return Haruf[letter], i, lastLetter
}

func GetLetterType(lastConnected bool, first bool, last bool, harf Harf) rune {
	var letterToAdd rune
	if first && !last {
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
	return letterToAdd
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
			var lettersToAdd []rune
			first := i == 0
			harf, i, last = GetNextLetter(word, i)
			lettersToAdd = append(lettersToAdd, GetLetterType(lastConnected, first, last, harf))

			//don't get the next letter if we're at the end
			if i < wordLen-1 {
				nextHarf, nextI, _ := GetNextLetter(word, i)

				// if the next letter is a vowel, add that as well
				if nextHarf.Tashkeel {
					lettersToAdd = append(lettersToAdd, nextHarf.Initial)
					i = nextI
				}
			}
			lastConnected = harf.ConnectNext
			arabicWords = append(lettersToAdd, arabicWords...)
		}
		arabicWords = append([]rune{space}, arabicWords...)
	}
	return string(arabicWords)
}
func main() {
	englishWords := os.Args[1:]
	fmt.Println(Transliterate(englishWords))
}
