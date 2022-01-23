package main

import (
	"fmt"
)

func main() {
	translator := NewChordsTranslator()
	bSharpnotes := translator.Process("B#", "min")

	fmt.Println(bSharpnotes)
}
