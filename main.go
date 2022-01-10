package main

import (
	"fmt"
)

func main() {
	chord := []string{"B", "min"}

	translator := NewChordsTranslator()
	translator.Process(chord)

	fmt.Println(translator.CurrentState.ToString())
}
