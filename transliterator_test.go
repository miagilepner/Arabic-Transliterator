package main

import "testing"

func TestIsVowel(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"aN", true},
		{"uN", true},
		{"iN", true},
		{"a", true},
		{"u", true},
		{"i", true},
		{"-", true},
		{"'a", false},
		{"'i", false},
	}
	for _, c := range cases {
		got := isVowel(c.in)
		if got != c.want {
			t.Errorf("IsVowel (%s) should be (%t)", c.in, c.want)
		}
	}
}

func TestMultiLetterUnit(t *testing.T) {
	cases := []struct {
		inWord string
		inI    int
		want   string
	}{
		{"refaN", 3, "aN"},
		{"aafaN", 0, "aa"},
		{"aafaN", 3, "aN"},
		{"refan", 3, "a"},
		{"refaan", 3, "aa"},
		{"thaqat", 0, "th"},
		{"tahaqat", 0, "t"},
		{"'awa", 0, "'a"},
		{"'b", 0, "'"},
	}
	for _, c := range cases {
		got, _ := multiLetterUnit(c.inWord, c.inI)
		if got != c.want {
			t.Errorf("MultiLetterUnit (%q, %d) should be (%q)", c.inWord, c.inI, c.want)
		}
	}
}

func TestGetLetterType(t *testing.T) {
	cases := []struct {
		inLastConn bool
		inFirst    bool
		inLast     bool
		inHarf     Harf
		want       rune
	}{
		{false, true, false, Haruf["aa"], '\u0627'},
		{false, false, false, Haruf["th"], '\ufe9b'},
		{false, true, true, Haruf["j"], '\ufe9d'},
		{true, false, true, Haruf["ae"], '\ufef0'},
		{true, false, false, Haruf["h"], '\ufeec'},
	}
	for _, c := range cases {
		got := letterType(c.inLastConn, c.inFirst, c.inLast, c.inHarf)
		if got != c.want {
			t.Errorf("GetLetterType (%t, %t, %t, %v) should be (%c), got (%c)", c.inLastConn, c.inFirst, c.inLast, c.inHarf, c.want, got)
		}
	}
}
