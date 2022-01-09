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
	fundamental := -1
	newTranslator := &ChordsTranslator{
		States: map[string]Chord{},
		CurrentState: Chord{
			Fundamental: &fundamental,
			Third:       Interval{},
			Fifth:       Interval{},
		},
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
	insertSymbolsMapping(translator, minSymbols, Intervals[Third][Minor], Intervals[Fifth][Just])
	insertSymbolsMapping(translator, augSymbols, Intervals[Third][Major], Intervals[Fifth][Augmented])
	insertSymbolsMapping(translator, dimSymbols, Intervals[Third][Minor], Intervals[Fifth][Diminished])
}

func insertSymbolsMapping(translator *ChordsTranslator, lettersMap map[string]bool, third, fifth Interval) {
	for symbol := range lettersMap {
		translator.States[symbol] = Chord{
			Fifth:       fifth,
			Third:       third,
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
	if *translator.CurrentState.Fundamental == -1 {
		notePosition, validSymbol := notesPositions[symbol]
		if !validSymbol {
			return
		}

		*translator.CurrentState.Fundamental = notePosition
	}

	translator.CurrentState = translator.States[symbol]
}
