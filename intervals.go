package main

import "math"

const limit = 11

type IntervalCalification string

const (
	Diminished IntervalCalification = "D"
	Minor      IntervalCalification = "m"
	Just       IntervalCalification = "J"
	Major      IntervalCalification = "M"
	Augmented  IntervalCalification = "A"
)

type IntervalClassification int

const (
	Third IntervalClassification = 3
	Fifth IntervalClassification = 5
)

type Interval struct {
	ChromaticSemitones int
	DiatonicSemitones  int
}

func (interval Interval) GetSemitonesSum() int {
	return interval.ChromaticSemitones + interval.DiatonicSemitones
}

var Intervals = map[IntervalClassification]map[IntervalCalification]Interval{
	Third: map[IntervalCalification]Interval{
		Diminished: Interval{ChromaticSemitones: 0, DiatonicSemitones: 2},
		Minor:      Interval{ChromaticSemitones: 1, DiatonicSemitones: 2},
		Major:      Interval{ChromaticSemitones: 2, DiatonicSemitones: 2},
		Augmented:  Interval{ChromaticSemitones: 3, DiatonicSemitones: 2},
	},
	Fifth: map[IntervalCalification]Interval{
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
		return getThirdFromInterval(note, interval)
	}

	return ""
}

func getThirdFromInterval(sourceNote string, interval Interval) string {
	semitonesSum := interval.GetSemitonesSum()

	targetPosition := notesPositions[sourceNote] + semitonesSum
	targetNote := positionsNotes[targetPosition]
	targetName := ExtractNoteRawName(targetNote)

	rawSourceNote := ExtractNoteRawName(sourceNote)

	targetRawPosition := notesPositions[rawSourceNote] + semitonesSum
	targetRawNote := positionsNotes[targetRawPosition]
	targetRawName := ExtractNoteRawName(targetRawNote)

	if targetName != targetRawName {
		if targetPosition < targetRawPosition {
      return missingBemolsPositionsNotes[targetPosition]
    }

    return missingSharpsPositionsNotes[targetPosition]
	}

	return positionsNotes[targetPosition]
}

func normalizeNotePosition(notePosition int) int {
	if notePosition <= limit {
		return notePosition
	}

	timesLimitIsContained := int(math.Floor(float64(notePosition) / float64(limit)))
	normalizedPosition := notePosition - limit*timesLimitIsContained

	return normalizedPosition
}

func noteNamesAreEqual(firstNote, secondNote string) bool {
	if firstNote[0] == secondNote[0] {
		return true
	}

	return false
}

func ExtractNoteRawName(note string) string {
	if len(note) == 1 {
		return note
	}

	return note[0:1]
}
