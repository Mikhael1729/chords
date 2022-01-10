package main

import "strings"

type Chord struct {
	Fundamental *int
	Third       Interval
	Fifth       Interval
}

func (chord *Chord) GetCalification() ChordCalification {
	if chord.IsMajor() {
		return MajorChord
	}

	return MinorChord
}

func (chord *Chord) SumOperation(operation ChordOperation) {
	chord.Third = SumIntervals(chord.Third, operation.ThirdOperation)
	chord.Fifth = SumIntervals(chord.Fifth, operation.FifthOperation)
}

func (chord *Chord) IsMajor() bool {
	thirdIsMajor := chord.Third == Intervals[Third][Major]
	fifthIsJust := chord.Fifth == Intervals[Fifth][Just]

	return thirdIsMajor && fifthIsJust
}

func (chord *Chord) ToString() string {
	notes := []string{}

	fundamental := chord.getFundamentalName()
	notes = append(notes, fundamental)

	third := chord.getNote(Third, chord.Third)
	notes = append(notes, third)

	fifth := chord.getNote(Fifth, chord.Fifth)
	notes = append(notes, fifth)

	return strings.Join(notes, ",")
}

func (chord *Chord) getNote(classification IntervalClassification, interval Interval) string {
	return GetNoteFromInterval(
		chord.getFundamentalName(),
		classification,
		Intervals[classification][GetCalification(classification, interval)],
	)
}

func (chord *Chord) getFundamentalName() string {
	return positionsNotes[*chord.Fundamental]
}
