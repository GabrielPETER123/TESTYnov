package main

import "github.com/01-edu/z01"

func main()  {
	var startASCIIValue int = 97

	for i := 0; i < 26; i++ {
		z01.PrintRune(rune(startASCIIValue + i))
	}
	z01.PrintRune('\n')
}