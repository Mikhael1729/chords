package main

var (
	augSymbols = map[string]bool{"aug": true, "+": true}
	majSymbols = map[string]bool{"maj": true, "M": true, "Δ": true}
	minSymbols = map[string]bool{"min": true, "m": true}
	dimSymbols = map[string]bool{"dim": true, "°": true}
)

type ChordsTranslator struct {
	States       map[string]Chord
	InitialState Chord
	CurrentState Chord
}

func NewChordsTranslator() *ChordsTranslator {
	newTranslator := &ChordsTranslator{
		States:       map[string]Chord{},
		CurrentState: Chord{},
		InitialState: Chord{},
	}

	initializeStates(newTranslator)

	return newTranslator
}

func initializeStates(translator *ChordsTranslator) {
	for note := range notesPositions {
		translator.States[note] = Chord{
			Fifth:       Intervals[Fifth][Just],
			Third:       Intervals[Third][Major],
			Fundamental: translator.CurrentState.Fundamental,
		}
	}

  insertSymbolsMapping(translator, majSymbols, Intervals[Third][Major], Intervals[Fifth][Just])
	//insertSymbolsMapping(translator, minSymbols, 3, 7)
	//insertSymbolsMapping(translator, augSymbols, 4, 8)
	//insertSymbolsMapping(translator, dimSymbols, 3, 6)
}

func insertSymbolsMapping(translator *ChordsTranslator, lettersMap map[string]bool, third, fifth Interval) {
	for symbol := range lettersMap {
		translator.States[symbol] = Chord{
			Fifth:       third,
			Third:       fifth,
			Fundamental: translator.CurrentState.Fundamental,
		}
	}
}

func (translator *ChordsTranslator) Process(word []string) {
	for _, word := range word {
		translator.Transition(word)
	}
}

func (translator *ChordsTranslator) Transition(symbol string) {
	if *translator.CurrentState[0] == 0 {
		notePosition, ok := notesPositions[symbol]
		if !ok {
			return
		}

		*translator.CurrentState[0] = notePosition
	}

	translator.CurrentState = translator.States[symbol]
}
