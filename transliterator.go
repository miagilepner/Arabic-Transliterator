package main

import "fmt"

func isVowel(letter string) bool {
	harf, ok := Haruf[letter]
	if !ok {
		return false
	}
	if harf.Tashkeel {
		return true
	}
	return false
}

func multiLetterUnit(word string, i int) (string, int) {
	letter := string(word[i])
	nextLetter := string(word[i+1])
	lengthOfUnit := 1
	switch letter {
	case "a":
		switch nextLetter {
		case "a", "e", "N":
			lengthOfUnit = 2
		}
	case "i":
		switch nextLetter {
		case "N":
			lengthOfUnit = 2
		}
	case "u":
		switch nextLetter {
		case "N":
			lengthOfUnit = 2
		}
	case "s", "t", "k", "d", "g":
		switch nextLetter {
		case "h":
			lengthOfUnit = 2
		}
	case "'":
		switch nextLetter {
		case "a", "y", "w":
			lengthOfUnit = 2
		}
	case "~":
		// the next two characters must be "aa"
		if nextLetter != "a" ||
			len(word) < i+2 ||
			string(word[i+2]) != "a" {
			panic(fmt.Sprintf("invalid input %s, ~ must be followed by 'aa'", word))
		}
		lengthOfUnit = 3
	}
	return word[i : i+lengthOfUnit], i + lengthOfUnit
}

func nextLetter(word string, i int) (Harf, int, bool) {
	wordLen := len(word)
	letter := string(word[i])
	lastLetter := false
	if i == wordLen-1 {
		i++
		lastLetter = true
	} else {
		letter, i = multiLetterUnit(word, i)
	}
	if i == wordLen-1 || i == wordLen-2 {
		nextLetter := string(word[i:wordLen])
		if isVowel(nextLetter) {
			lastLetter = true
		}
	}
	return Haruf[letter], i, lastLetter
}

func letterType(lastConnected bool, first bool, last bool, harf Harf) rune {
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

func Transliterate(englishWords []string, vowels bool) string {
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
			harf, i, last = nextLetter(word, i)
			lettersToAdd = append(lettersToAdd, letterType(lastConnected, first, last, harf))

			//don't get the next letter if we're at the end
			if i < wordLen {
				nextHarf, nextI, _ := nextLetter(word, i)

				// if the next letter is a vowel, add that as well
				if nextHarf.Tashkeel {
					if vowels {
						lettersToAdd = append(lettersToAdd, nextHarf.Initial)
					}
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
