package main

import (
	"strings"
)

var (
	notePositions = map[string]int{"C": 1, "C#": 2, "Cb": 3, "D": 4, "D#": 5, "Db": 6, "E": 7, "E#": 8, "Eb": 9, "F": 10, "F#": 11, "Fb": 12, "G": 13, "G#": 14, "Gb": 15, "A": 16, "A#": 17, "Ab": 18, "B": 19, "B#": 20, "Bb": 21}
	positionNotes = map[int]string{1: "C", 2: "C#", 3: "Cb", 4: "D", 5: "D#", 6: "Db", 7: "E", 8: "E#", 9: "Eb", 10: "F", 11: "F#", 12: "Fb", 13: "G", 14: "G#", 15: "Gb", 16: "A", 17: "A#", 18: "Ab", 19: "B", 20: "B#", 21: "Bb"}
	augSymbols    = map[string]bool{"aug": true, "+": true}
	majSymbols    = map[string]bool{"maj": true, "M": true, "Δ": true}
	minSymbols    = map[string]bool{"min": true, "m": true}
	dimSymbols    = map[string]bool{"dim": true, "°": true}
)

type ChordsTranslator struct {
	States       map[string][]*int
	InitialState []*int
	CurrentState []*int
}

func NewChordsTranslator() *ChordsTranslator {
	newTranslator := &ChordsTranslator{
		States:       map[string][]*int{},
		CurrentState: []*int{new(int), new(int), new(int)},
		InitialState: []*int{new(int), new(int), new(int)},
	}

	initializeStates(newTranslator)

	return newTranslator
}

func initializeStates(translator *ChordsTranslator) {
	for note, _ := range notePositions {
		third, fifth := initThirdAndFifth(4, 7)
		translator.States[note] = []*int{translator.CurrentState[0], third, fifth}
	}

	insertSymbolsMapping(translator, majSymbols, 4, 7)
	insertSymbolsMapping(translator, minSymbols, 3, 7)
	insertSymbolsMapping(translator, augSymbols, 4, 8)
	insertSymbolsMapping(translator, dimSymbols, 3, 6)
}

func insertSymbolsMapping(translator *ChordsTranslator, lettersMap map[string]bool, third, fifth int) {
	for symbol, _ := range lettersMap {
		third, fifth := initThirdAndFifth(third, fifth)
		translator.States[symbol] = []*int{translator.CurrentState[0], third, fifth}
	}
}

func initThirdAndFifth(thirdValue, fifthValue int) (*int, *int) {
	third, fifth := new(int), new(int)
	*third = thirdValue
	*fifth = fifthValue

	return third, fifth
}

func (translator *ChordsTranslator) GetStateInNotesFormat() string {
	notes := []string{}
	fundamentalPosition := *translator.CurrentState[0]

	fundamental := getFundamentalName(fundamentalPosition)
	notes = append(notes, fundamental)

	third := getThirdNoteName(fundamentalPosition, *translator.CurrentState[1])
	notes = append(notes, third)

	fifth := getFifthNoteName(fundamentalPosition, *translator.CurrentState[2])
	notes = append(notes, fifth)

	return strings.Join(notes, ",")
}

func getFundamentalName(fundamentalPosition int) string {
	return positionNotes[fundamentalPosition]
}

func getThirdNoteName(fundamentalPosition, thirdNotePosition int) string {
	if thirdNotePosition == 3 {
		return positionNotes[fundamentalPosition+8]
	}

	return positionNotes[fundamentalPosition+6]
}

func getFifthNoteName(fundamentalPosition, fithNotePosition int) string {
	if fithNotePosition == 6 {
		return positionNotes[fundamentalPosition+14]
	}
	if fithNotePosition == 7 {
		return positionNotes[fundamentalPosition+12]
	}

	return positionNotes[fundamentalPosition+13]
}

func (translator *ChordsTranslator) Process(word []string) {
	for _, word := range word {
		translator.Transition(word)
	}
}

func (translator *ChordsTranslator) Transition(symbol string) {
	if *translator.CurrentState[0] == 0 {
		notePosition, ok := notePositions[symbol]
		if !ok {
			return
		}

		*translator.CurrentState[0] = notePosition
	}

	translator.CurrentState = translator.States[symbol]
}
