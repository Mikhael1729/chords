package main

import (
	"fmt"
)

func main() {
	chord := []string{"C#", "min"}

	translator := NewChordsTranslator()
	translator.Process(chord)

	fmt.Println(translator.CurrentState.ToString())
}
