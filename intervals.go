package main

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

func GetNoteFromInterval(note string, interval Interval, classification IntervalClassification) string {
  basePosition := ascendingNotesPositions[note]
  targetPosition := basePosition + interval.GetSemitonesSum()

  if classification == Third {
    baseRawNote := ExtractNoteRawName(note)
    targetRawNote := Ñ
    possibleTargetRawNote := ExtractNoteRawName(ascendingPositionsNotes[targetPosition])

    //targetsTheCorrectNote := 
  }
}

func ExtractNoteRawName(note string) string {
  if len(note) == 2 {
    return note
  }

  return note[0:1]
}
