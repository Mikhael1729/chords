package main

import "strings"

type Chord struct {
	Fundamental *int
	Third       Interval
	Fifth       Interval
}

func (chord Chord) GetCalification() IntervalCalification {
	if chord.IsMajor() {
		return Major
	}

	return Minor
}

func (chord Chord) IsMajor() bool {
	thirdIsMajor := chord.Third == Intervals[Third][Major]
	fifthIsJust := chord.Fifth == Intervals[Fifth][Just]

	return thirdIsMajor && fifthIsJust
}

func (chord *Chord) ToString() string {
	notes := []string{}
	//fundamentalPosition := *chord.Fundamental

	fundamental := chord.getFundamentalName()
	notes = append(notes, fundamental)

	third := chord.getThirdNoteName()
	notes = append(notes, third)

	fifth := chord.getFifthNote()
	notes = append(notes, fifth)

	return strings.Join(notes, ",")
}

func (chord Chord) getFifthNote() string {
	fundamental := chord.getFundamentalName()
	fifthCalification := GetCalification(Fifth, chord.Fifth)
	chordClassification := Intervals[Fifth][fifthCalification]
	fifth := GetNoteFromInterval(fundamental, Fifth, chordClassification)

	return fifth
}

func (chord Chord) getThirdNoteName() string {
	fundamental := chord.getFundamentalName()
	thirdCalification := GetCalification(Third, chord.Third)
	chordClassification := Intervals[Third][thirdCalification]
	third := GetNoteFromInterval(fundamental, Third, chordClassification)

	return third
}

func (chord Chord) getFundamentalName() string {
	return positionsNotes[*chord.Fundamental]
}

func getFifthNoteName(fundamentalPosition, fithNotePosition int) string {
	if fithNotePosition == 6 {
		return positionsNotes[fundamentalPosition+14]
	}
	if fithNotePosition == 7 {
		return positionsNotes[fundamentalPosition+12]
	}

	return positionsNotes[fundamentalPosition+13]
}
