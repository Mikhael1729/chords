package main

import (
  "strings"
)

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
	augSymbols  = []string{"aug", "+"}
	majSymbosl  = []string{"maj", "M", "Δ"}
	minSymbols  = []string{"min", "m"}
	dimSymbols  = []string{"dim", "°"}
	states      = []string{"Empty", "Major", "Minor", "Augmented"}
)

type ChordsTranslator struct {
	States       map[string]bool
	InitialState string
	CurrentState string
	Adjacency    map[string]string
}

func NewChordsTranslator() *ChordsTranslator {
	return &ChordsTranslator{}
}

func (translator *ChordsTranslator) Transition(symbol string) {
}

func getNextNoteOfInterval(note string, interval int) {
  basePosition := notes[note]
  nextPosition := 
  nextNote := 
}

func fin
