package main

import (
	"fmt"
)

var (
	notesPositions = map[string]int{
		"C": 0, "D": 2, "E": 4, "F": 5,
		"G": 7, "A": 9, "B": 11,
	}
	positionsNotes = map[int]string{
		0: "C", 2: "D", 4: "E", 5: "F",
		7: "G", 9: "A", 11: "B",
	}
	bemolsNotesPositions = map[string]int{
		"Db": 1, "Eb": 3, "Fb": 4, "Gb": 6,
		"Ab": 8, "Bb": 10, "Cb": 11,
	}
	bemolsPositionsNotes = map[int]string{
		1: "Db", 3: "Eb", 4: "Fb", 6: "Gb",
		8: "Ab", 10: "Bb", 11: "Cb",
	}
	sharpsNotesPositions = map[string]int{
		"B#": 0, "C#": 1, "D#": 3, "E#": 5,
		"F#": 6, "G#": 8, "A#": 10,
	}
	sharpsPositionsNotes = map[int]string{
		0: "B#", 1: "C#", 3: "D#", 5: "E#",
		6: "F#", 8: "G#", 10: "A#",
	}
	doubleBemolsNotesPositions = map[string]int{
		"Dbb": 0, "Ebb": 2, "Fbb": 3, "Gbb": 5,
		"Abb": 7, "Bbb": 9, "Cbb": 10,
	}
	doubleBemolsPositionsNotes = map[int]string{
		0: "Dbb", 2: "Ebb", 3: "Fbb", 5: "Gbb",
		7: "Abb", 9: "Bbb", 10: "Cbb",
	}
	doubleSharpsNotesPositions = map[string]int{
		"Cx": 1, "Dx": 2, "Ex": 4, "Fx": 6,
		"Gx": 7, "Ax": 9, "Bx": 11,
	}
	doubleSharpsPositionsNotes = map[int]string{
		1: "Cx", 2: "Dx", 4: "Ex", 6: "Fx",
		7: "Gx", 9: "Ax", 11: "Bx",
	}
)

func GetNotePosition(note string) (int, *map[int]string, error) {
	position, ok := notesPositions[note]
	if ok {
		return position, &positionsNotes, nil
	}

	position, ok = bemolsNotesPositions[note]
	if ok {
		return position, &bemolsPositionsNotes, nil
	}

	position, ok = sharpsNotesPositions[note]
	if ok {
		return position, &sharpsPositionsNotes, nil
	}

	position, ok = doubleBemolsNotesPositions[note]
	if ok {
		return position, &doubleBemolsPositionsNotes, nil
	}

	position, ok = doubleSharpsNotesPositions[note]
	if ok {
		return position, &doubleSharpsPositionsNotes, nil
	}

	return position, nil, fmt.Errorf(fmt.Sprintf("Invalid note %v", note))
}

func GetNote(position int) string {
	note := GetNaturalNote(position)
	if note != "" {
		return note
	}

	note = GetBemolNote(position)
	if note != "" {
		return note
	}

	note = GetSharpNote(position)
	if note != "" {
		return note
	}

	return ""
}

func GetNaturalNote(position int) string {
	note, ok := positionsNotes[position]
	if ok {
		return note
	}

	return ""
}

func GetBemolNote(position int) string {
	note, ok := bemolsPositionsNotes[position]
	if ok {
		return note
	}

	note, ok = doubleBemolsPositionsNotes[position]
	if ok {
		return note
	}

	return ""
}

func GetSharpNote(position int) string {
	note, ok := sharpsPositionsNotes[position]
	if ok {
		return note
	}

	note, ok = doubleSharpsPositionsNotes[position]
	if ok {
		return note
	}

	return ""
}
