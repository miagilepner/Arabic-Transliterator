package main

import (
	"flag"
	"fmt"
)

func main() {
	boolPtr := flag.Bool("vowels", true, "vowels = true will print vowels")
	flag.Parse()
	fmt.Println(Transliterate(flag.Args(), *boolPtr))
}
