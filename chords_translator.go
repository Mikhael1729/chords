package main

var (
	augSymbols = map[string]bool{"aug": true, "+": true}
	majSymbols = map[string]bool{"maj": true, "M": true, "Δ": true}
	minSymbols = map[string]bool{"min": true, "m": true}
	dimSymbols = map[string]bool{"dim": true, "°": true}
)

type ChordOperation struct {
	ThirdOperation Interval
	FifthOperation Interval
}

type ChordsTranslator struct {
	States       map[string]map[ChordCalification]ChordOperation
	InitialState Chord
	CurrentState Chord
}

func NewChordsTranslator() *ChordsTranslator {
	fundamental := -1
	newTranslator := &ChordsTranslator{
		States: map[string]map[ChordCalification]ChordOperation{},
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
		translator.States[note] = make(map[ChordCalification]ChordOperation)
		translator.States[note][UndefinedChordCalification] = ChordOperation{
			ThirdOperation: Intervals[Third][Major],
			FifthOperation: Intervals[Fifth][Just],
		}
	}

	insertSymbolsMapping(translator, MajorChord, majSymbols, ChordOperation{Interval{0, 0}, Interval{0, 0}})
	insertSymbolsMapping(translator, MajorChord, minSymbols, ChordOperation{Interval{-1, 0}, Interval{0, 0}})
	insertSymbolsMapping(translator, MajorChord, augSymbols, ChordOperation{Interval{0, 0}, Interval{1, 0}})
	insertSymbolsMapping(translator, MajorChord, dimSymbols, ChordOperation{Interval{-1, 0}, Interval{-1, 0}})
}

func insertSymbolsMapping(translator *ChordsTranslator, chordCalifiation ChordCalification, lettersMap map[string]bool, operation ChordOperation) {
	for symbol := range lettersMap {
		translator.States[symbol] = make(map[ChordCalification]ChordOperation)
		translator.States[symbol][chordCalifiation] = operation
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

		translator.CurrentState.SumOperation(translator.States[symbol][UndefinedChordCalification])
	}

	chordCalification := translator.CurrentState.GetCalification()

	translator.CurrentState.SumOperation(translator.States[symbol][chordCalification])
}
