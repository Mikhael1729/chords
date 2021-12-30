package main

var (
	ascendingNotesPositions = map[string]int{
		"C": 1, "C#": 2,
		"D": 3, "D#": 4,
		"E": 5, "F": 6,
		"F#": 7, "G": 8,
		"G#": 9, "A": 10,
		"A#": 11, "B": 12,
	}
	ascendingPositionsNotes = map[int]string{
		1: "C", 2: "C#",
		3: "D", 4: "D#",
		5: "E", 6: "F",
		7: "F#", 8: "G",
		9: "G#", 10: "A",
		11: "A#", 12: "B",
	}
	descendingNotesPositions = map[string]int{
		"B": 1, "Bb": 2,
		"A": 3, "Ab": 4,
		"G": 5, "Gb": 6,
		"F": 7, "E": 8,
		"Eb": 9, "D": 10,
		"Db": 11, "C": 12,
	}
	descendingPositionsNotes = map[int]string{
		1: "B", 2: "Bb",
		3: "A", 4: "Ab",
		5: "G", 6: "Gb",
		7: "F", 8: "E",
		9: "Eb", 10: "D",
		11: "Db", 12: "C",
	}
)
