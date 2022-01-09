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

func scanString() string {
	var str string
	fmt.Scanln(&str)

	return str
}

/*

C# -->

https://es.m.wikipedia.org/wiki/Tr%C3%ADada_(m%C3%BAsica)
https://app.diagrams.net/#G1NXb9pOKQtFDPPHiPULvM8NJJrnOWTkNB

*/
