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
	sourcePosition, _, _ := GetNotePosition(sourceNote)

	var targetNote string
	for i := sourcePosition; i < limit; i++ {
		if interval.GetSemitonesSum() == 0 {
			return targetNote
		}

		source := GetNote(i)
		targetNote = GetNote(normalizePosition(i + 1))
		semitoneType := GetSemitoneType(source, targetNote)

		if semitoneType == Chromatic {
			if interval.ChromaticSemitones > 0 {
				interval.ChromaticSemitones -= 1

				continue
			}

			targetNote = GetNoteFromSemitone(source, Diatonic)
			interval.DiatonicSemitones -= 1

			continue
		}

		if interval.DiatonicSemitones > 0 {
			interval.DiatonicSemitones -= 1

			continue
		}

		targetNote = GetNoteFromSemitone(source, Chromatic)
		interval.ChromaticSemitones -= 1
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

func virtualizeNote(targetPosition, targetRawPosition int) int {
	if math.Abs(float64(targetPosition-targetRawPosition)) < limit-1 {
		return targetPosition
	}

	if targetPosition < targetRawPosition {
		return targetPosition + limit
	}

	return targetPosition
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
