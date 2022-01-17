package main

import "math"

const limit = 12

type ChordCalification string

const (
	MinorChord                 ChordCalification = "minor"
	MajorChord                 ChordCalification = "major"
	DiminishedChord            ChordCalification = "diminished"
	AugmentedChord             ChordCalification = "augmented"
	UndefinedChordCalification ChordCalification = "undefined"
)

type IntervalCalification string

const (
	Diminished            IntervalCalification = "D"
	Minor                 IntervalCalification = "m"
	Just                  IntervalCalification = "J"
	Major                 IntervalCalification = "M"
	Augmented             IntervalCalification = "A"
	UndefinedCalification IntervalCalification = "~"
)

type IntervalClassification int

const (
	Third                   IntervalClassification = 3
	Fifth                   IntervalClassification = 5
	UndefinedClassification IntervalClassification = 0
)

type Interval struct {
	ChromaticSemitones int
	DiatonicSemitones  int
}

func (interval Interval) GetSemitonesSum() int {
	return interval.ChromaticSemitones + interval.DiatonicSemitones
}

var Intervals = map[IntervalClassification]map[IntervalCalification]Interval{
	Third: {
		Diminished: Interval{ChromaticSemitones: 0, DiatonicSemitones: 2},
		Minor:      Interval{ChromaticSemitones: 1, DiatonicSemitones: 2},
		Major:      Interval{ChromaticSemitones: 2, DiatonicSemitones: 2},
		Augmented:  Interval{ChromaticSemitones: 3, DiatonicSemitones: 2},
	},
	Fifth: {
		Diminished: Interval{ChromaticSemitones: 2, DiatonicSemitones: 4},
		Just:       Interval{ChromaticSemitones: 3, DiatonicSemitones: 4},
		Augmented:  Interval{ChromaticSemitones: 4, DiatonicSemitones: 4},
	},
}

func GetCalification(classification IntervalClassification, interval Interval) IntervalCalification {
	for calification, currentInterval := range Intervals[classification] {
		if currentInterval == interval {
			return calification
		}
	}

	return IntervalCalification("")
}

func (interval Interval) GetNote(sourceNote string) string {
	sourcePosition, getSourceNote := GetNotePosition2(sourceNote)

	var targetNote string
	for i := sourcePosition; i < limit-1; i++ {
		if interval.GetSemitonesSum() == 0 {
			return targetNote
		}

		normalizedTargePosition := normalizePosition(i + 1)

		if (sourcePosition-i+1)%2 == 0 {
			if interval.DiatonicSemitones > 0 {
				targetNote = getSourceNote(normalizedTargePosition, Diatonic)
				interval.DiatonicSemitones -= 1
				continue
			}
		}

		if interval.ChromaticSemitones > 0 {
			targetNote = getSourceNote(normalizedTargePosition, Chromatic)
			interval.ChromaticSemitones -= 1
			continue
		}

		targetNote = getSourceNote(normalizedTargePosition, Diatonic)
		interval.DiatonicSemitones -= 1
	}

	return ""
}

func normalizePosition(notePosition int) int {
	if notePosition < limit {
		return notePosition
	}

	timesLimitIsContained := int(math.Floor(float64(notePosition) / float64(limit)))
	normalizedPosition := notePosition - limit*timesLimitIsContained

	return normalizedPosition
}

func ExtractNoteRawName(note string) string {
	if len(note) == 1 {
		return note
	}

	return note[0:1]
}

func SumIntervals(interval1, interval2 Interval) Interval {
	result := Interval{
		ChromaticSemitones: interval1.ChromaticSemitones + interval2.ChromaticSemitones,
		DiatonicSemitones:  interval1.DiatonicSemitones + interval2.DiatonicSemitones,
	}

	return result
}
