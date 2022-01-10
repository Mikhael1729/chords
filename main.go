package main

import (
	"fmt"
)

func main() {
	translator := NewChordsTranslator()

	// fmt.Println(translator.Process([]string{"B", "min"}))
	fmt.Println(translator.Process([]string{"B#"}))
	// fmt.Println(translator.Process([]string{"B#", "min"}))
	// fmt.Println(translator.Process([]string{"F"}))
	// fmt.Println(translator.Process([]string{"E"}))
}
