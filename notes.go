package main

import "math"

type NoteBase int

const (
	NaturalNotesBase     = 0
	SharpNotesBase       = 24
	DoubleSharpNotesBase = 48
	BemolNotesBase       = -24
	DoubleBemolNotesBase = -48
)

const (
	NotesQuantity = 12
	Separation    = 24
)

var (
	notesPositions = map[string]int{
		"Dbb": -48, "Ebb": -50, "Fbb": -51, "Gbb": -53, "Abb": -55, "Bbb": -57, "Cbb": -58,
		"Db": -25, "Eb": -27, "Fb": -28, "Gb": -30, "Ab": -32, "Bb": -34, "Cb": -35,
		"C": 0, "D": 2, "E": 4, "F": 5, "G": 7, "A": 9, "B": 11,
		"B#": 24, "C#": 25, "D#": 27, "E#": 29, "F#": 30, "G#": 32, "A#": 34,
		"Bx": 49, "Cx": 50, "Dx": 52, "Ex": 54, "Fx": 55, "Gx": 57, "Ax": 59,
	}

	positionsNotes = map[int]string{
		-48: "Dbb", -50: "Ebb", -51: "Fbb", -53: "Gbb", -55: "Abb", -57: "Bbb", -58: "Cbb",
		-25: "Db", -27: "Eb", -28: "Fb", -30: "Gb", -32: "Ab", -34: "Bb", -35: "Cb",
		0: "C", 2: "D", 4: "E", 5: "F", 7: "G", 9: "A", 11: "B",
		24: "B#", 25: "C#", 27: "D#", 29: "E#", 30: "F#", 32: "G#", 34: "A#",
		49: "Bx", 50: "Cx", 52: "Dx", 54: "Ex", 55: "Fx", 57: "Gx", 59: "Ax",
	}
)

func GetNotePosition(note string) int {
	return notesPositions[note]
}

func GetNoteName(position int) string {
	return positionsNotes[position]
}

func NormalizeNotePosition(position int) int {
	base := int(getBase(position))
	isNegative := position < 0
	position = int(math.Abs(float64(position)))
	base = int(math.Abs(float64(base)))

	normalizedQuantity := base + NotesQuantity

	if position >= base && position < normalizedQuantity {
		return multiplyIf(isNegative, position, -1)
	}

	times := int(math.Floor(float64(position) / float64(normalizedQuantity)))
	normalizedPosition := position + base
	normalized := normalizedPosition - normalizedQuantity*times

	return multiplyIf(isNegative, normalized, -1)
}

func getBase(position int) NoteBase {
	if position >= DoubleSharpNotesBase {
		return DoubleSharpNotesBase
	}

	if position >= SharpNotesBase {
		return SharpNotesBase
	}

	if position >= NaturalNotesBase {
		return NaturalNotesBase
	}

	if position >= BemolNotesBase {
		return BemolNotesBase
	}

	return DoubleBemolNotesBase
}

func multiplyIf(condition bool, multiplicand, multiplier int) int {
	if !condition {
		return multiplicand
	}

	return multiplicand * multiplier
}
