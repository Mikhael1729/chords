package main

import (
	"fmt"
)

func main() {
	translator := NewChordsTranslator()
	bSharpMinor := translator.Process("B#", "min")

	fmt.Println(bSharpMinor)
}
