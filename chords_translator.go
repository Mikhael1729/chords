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
	Operations   map[string]map[ChordCalification]ChordOperation
	InitialState Chord
	CurrentState Chord
}

func NewChordsTranslator() *ChordsTranslator {
	newTranslator := &ChordsTranslator{
		Operations:   map[string]map[ChordCalification]ChordOperation{},
		CurrentState: Chord{Fundamental: -1},
		InitialState: Chord{Fundamental: -1},
	}

	initOperations(newTranslator)

	return newTranslator
}

func (translator *ChordsTranslator) Process(word []string) {
	// Reset the automaton before processing any word
	if translator.CurrentState.Fundamental != translator.InitialState.Fundamental {
		translator.CurrentState = translator.InitialState
	}

	for _, word := range word {
		translator.Transition(word)
	}
}

func (translator *ChordsTranslator) Transition(symbol string) {
	if translator.CurrentState.Fundamental == translator.InitialState.Fundamental {
		notePosition, validSymbol := notesPositions[symbol]
		if !validSymbol {
			return
		}

		translator.CurrentState.Fundamental = notePosition
		translator.CurrentState.SumOperation(translator.Operations[symbol][UndefinedChordCalification])
	}

	chordCalification := translator.CurrentState.GetCalification()

	translator.CurrentState.SumOperation(translator.Operations[symbol][chordCalification])
}

func initOperations(translator *ChordsTranslator) {
	for note := range notesPositions {
		translator.Operations[note] = make(map[ChordCalification]ChordOperation)
		translator.Operations[note][UndefinedChordCalification] = ChordOperation{
			ThirdOperation: Intervals[Third][Major],
			FifthOperation: Intervals[Fifth][Just],
		}
	}

	insertTransition(translator, MajorChord, majSymbols, ChordOperation{Interval{0, 0}, Interval{0, 0}})
	insertTransition(translator, MajorChord, minSymbols, ChordOperation{Interval{-1, 0}, Interval{0, 0}})
	insertTransition(translator, MajorChord, augSymbols, ChordOperation{Interval{0, 0}, Interval{1, 0}})
	insertTransition(translator, MajorChord, dimSymbols, ChordOperation{Interval{-1, 0}, Interval{-1, 0}})
}

func insertTransition(translator *ChordsTranslator, chordCalifiation ChordCalification, lettersMap map[string]bool, operation ChordOperation) {
	for symbol := range lettersMap {
		translator.Operations[symbol] = make(map[ChordCalification]ChordOperation)
		translator.Operations[symbol][chordCalifiation] = operation
	}
}
