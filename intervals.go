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

func GetNoteFromInterval(note string, classification IntervalClassification, interval Interval) string {
	if classification == Third {
		return getNoteFromInterval(note, Third, interval)
	}

	if classification == Fifth {
		return getNoteFromInterval(note, Fifth, interval)
	}

	return ""
}

func getNoteFromInterval(sourceNote string, classification IntervalClassification, interval Interval) string {
	var intervalCalification IntervalCalification

	if classification == Third {
		intervalCalification = Major
	} else {
		intervalCalification = Just
	}

	semitonesSum := interval.GetSemitonesSum()

	sourcePosition, _, err := GetNotePosition(sourceNote)
	if err != nil {
		panic(err)
	}

	targetPosition := normalizePosition(sourcePosition + semitonesSum)
	targetNote := positionsNotes[targetPosition]
	targetName := ExtractNoteRawName(targetNote)

	rawSourceNote := ExtractNoteRawName(sourceNote)

	targetRawPosition := normalizePosition(notesPositions[rawSourceNote] + Intervals[classification][intervalCalification].GetSemitonesSum())
	targetRawNote := positionsNotes[targetRawPosition]
	targetRawName := ExtractNoteRawName(targetRawNote)

	if targetName != targetRawName {
		virtualizedTarget := virtualizeNote(targetPosition, targetRawPosition)
		virtualizedRawTarget := virtualizeNote(targetRawPosition, targetPosition)

		if virtualizedTarget < virtualizedRawTarget {
			return missingBemolsPositionsNotes[targetPosition]
		}

		return missingSharpsPositionsNotes[targetPosition]
	}

	return positionsNotes[targetPosition]
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
