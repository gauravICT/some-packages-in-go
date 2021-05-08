package main

import (
	"fmt"
	"unicode"
)

func main() {
	ch := '\u0041'
	fmt.Println(unicode.IsLetter(ch)) //testing for a letter
	ch1 := '1'
	fmt.Println(unicode.IsDigit(ch1)) //testing for a digit
}
