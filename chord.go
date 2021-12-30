package main

import "strings"

type Chord struct {
	Fundamental *int
	Third       Interval
	Fifth       Interval
}

func (chord Chord) IsMajor() bool {
	thirdIsMajor := chord.Third == Intervals[Third][Major]
	fifthIsJust := chord.Fifth == Intervals[Fifth][Just]

  return thirdIsMajor && fifthIsJust
}

func (chord *Chord) ToString() string {
	notes := []string{}
	fundamentalPosition := *chord.Fundamental

	fundamental := chord.getFundamentalName()
	notes = append(notes, fundamental)

	third := chord.getThirdNoteName()
	notes = append(notes, third)

	fifth := getFifthNoteName(fundamentalPosition, chord.Fifth)
	notes = append(notes, fifth)

	return strings.Join(notes, ",")
}

func (chord Chord) getFundamentalName() string {
	return ascendingPositionsNotes[*chord.Fundamental]
}

func (chord Chord) getThirdNoteName() string {
	if thirdNotePosition == 3 {
		return ascendingPositionsNotes[fundamentalPosition+8]
	}

	return ascendingPositionsNotes[fundamentalPosition+6]
}

func getFifthNoteName(fundamentalPosition, fithNotePosition int) string {
	if fithNotePosition == 6 {
		return ascendingPositionsNotes[fundamentalPosition+14]
	}
	if fithNotePosition == 7 {
		return ascendingPositionsNotes[fundamentalPosition+12]
	}

	return ascendingPositionsNotes[fundamentalPosition+13]
}
