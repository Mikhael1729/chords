package main

import "strings"

type ChordCalification string

const (
	MinorChord                 ChordCalification = "minor"
	MajorChord                 ChordCalification = "major"
	DiminishedChord            ChordCalification = "diminished"
	AugmentedChord             ChordCalification = "augmented"
	UndefinedChordCalification ChordCalification = "undefined"
)

type ChordOperation struct {
	ThirdOperation Interval
	FifthOperation Interval
}

type Chord struct {
	Fundamental int
	Third       Interval
	Fifth       Interval
}

func (chord *Chord) GetCalification() ChordCalification {
	isMajor := chord.Third == Intervals[Third][Major] && chord.Fifth == Intervals[Fifth][Just]
	if isMajor {
		return MajorChord
	}

	isMinor := chord.Third == Intervals[Third][Minor] && chord.Fifth == Intervals[Fifth][Just]
	if isMinor {
		return MinorChord
	}

	isDiminished := chord.Third == Intervals[Third][Minor] && chord.Fifth == Intervals[Fifth][Diminished]
	if isDiminished {
		return DiminishedChord
	}

	isAugmented := chord.Third == Intervals[Third][Major] && chord.Fifth == Intervals[Fifth][Augmented]
	if isAugmented {
		return AugmentedChord
	}

	return MinorChord
}

func (chord *Chord) SumOperation(operation ChordOperation) {
	chord.Third = SumIntervals(chord.Third, operation.ThirdOperation)
	chord.Fifth = SumIntervals(chord.Fifth, operation.FifthOperation)
}

func (chord *Chord) ToString() string {
	notes := []string{}
	fundamental := GetNoteName(chord.Fundamental)

	notes = append(notes, fundamental)

	third := chord.Third.GetNote(fundamental)
	notes = append(notes, third)

	fifth := chord.Fifth.GetNote(fundamental)
	notes = append(notes, fifth)

	return strings.Join(notes, ", ")
}
