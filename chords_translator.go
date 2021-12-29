package main

import (
  "strings"
)


/* JUST NUMBERS 

The state is going to have:

[<1 - 21>, S, S]

The transition function will work as a map of symbols
*/

var (
	notePositions       = map[string]int{
    "C": 1, "C#": 2, "Cb": 3, "D": 4, "D#": 5, "Db": 6, "E": 7, "E#": 8, "Eb": 9, "F": 10, "F#": 11,
    "Fb": 12, "G": 13, "G#": 14, "Gb": 15, "A": 16, "A#": 17, "Ab": 18, "B": 19, "B#": 20, "Bb": 21,
  }
  positionNotes = map[int]string{
    1: "C", 2: "C#", 3: "Cb", 4: "D", 5: "D#", 6: "Db", 7: "E", 8: "E#", 9: "Eb", 10: "F", 11: "F#",
    12: "Fb", 13: "G", 14: "G#", 15: "Gb", 16: "A", 17: "A#", 18: "Ab", 19: "B", 20: "B#", 21: "Bb",
  }
	alterations = []string{"#", "x", "b", "bb"}
	augSymbols  = map[string]bool{ "aug": true, "+": true }
	majSymbosl  = map[string]bool{ "maj": true, "M": true, "Δ": true }
  minSymbols  = map[string]bool{"min": true, "m": true}
  dimSymbols  = map[string]bool{"dim": true, "°": true}
	states      = []string{"Empty", "Major", "Minor", "Augmented"}
)

type ChordsTranslator struct {
	States       map[string]bool
	InitialState string
	CurrentState []int
	Adjacency    map[string][]int
}

func (translator *ChordsTranslator) React(symbol string) {
  if position := notePositions[symbol]; position != 0 {
    translator.CurrentState = []int{position, 4, 7}
    return
  }

  if ok := minSymbols[symbol] {
  }
}

func NewChordsTranslator() *ChordsTranslator {
	return &ChordsTranslator{}
}

func createAdjacency() {
  symbolAndStates := map[string][]int{}

  for note, _ := range notePositions {
    symbolAndStates[note] = []int{positionNotes[note], 4, 7}
  }

  for _, symbol := range majSymbols {
    symbolAndStates[symbol] = []int{-1, 4, }
  }
}

func getKeys() {
  notes := []string
  for key, value := range notePositions {

  }
}

func (translator *ChordsTranslator) Transition(symbol string) {

}
