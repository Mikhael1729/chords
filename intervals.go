package main

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
	limit := int(getBase(GetNotePosition(sourceNote))) + NotesQuantity - 1

	var targetNote string
	for i := GetNotePosition(sourceNote); i < limit; i++ {
		targetPosition := NormalizeNotePosition(i + 1)
		targetNote = GetNoteName(targetPosition)

		if interval.GetSemitonesSum() == 0 {
			return targetNote
		}

		semitoneType := GetSemitoneType(i, targetPosition)
		if semitoneType == Chromatic {
			interval.ChromaticSemitones -= 1
			continue
		}

		interval.DiatonicSemitones -= 1
	}

	return targetNote
}

func SumIntervals(interval1, interval2 Interval) Interval {
	result := Interval{
		ChromaticSemitones: interval1.ChromaticSemitones + interval2.ChromaticSemitones,
		DiatonicSemitones:  interval1.DiatonicSemitones + interval2.DiatonicSemitones,
	}

	return result
}
