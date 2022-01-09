package main

var (
	notesPositions = map[string]int{
		"C": 0, "C#": 1,
		"D": 2, "D#": 3,
		"E": 4, "F": 5,
		"F#": 6, "G": 7,
		"G#": 8, "A": 9,
		"A#": 10, "B": 11,
	}
	positionsNotes = map[int]string{
		0: "C", 1: "C#",
		2: "D", 3: "D#",
		4: "E", 5: "F",
		6: "F#", 7: "G",
		8: "G#", 9: "A",
		10: "A#", 11: "B",
	}
  missingSharpsNotesPositions = map[string]int{
    "B#": 0,
    "E#": 5,
  }
  missingSharpsPositionsNotes = map[int]string{
    0: "B#",
    5: "E#",
  }
  missingBemolsNotesPositions = map[string]int{
    "Db": 1,
    "Eb": 3,
    "Fb": 4,
    "Gb": 6,
    "Ab": 8,
    "Bb": 10,
    "Cb": 11,
  }
  missingBemolsPositionsNotes = map[int]string{
    1: "Db",
    3: "Eb", 
    4: "Fb", 
    6: "Gb", 
    8: "Ab", 
    10: "Bb", 
    11: "Cb", 
  }
)
