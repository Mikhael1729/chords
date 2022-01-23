package main

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

func (interval Interval) GetClassification() IntervalClassification {
	sum := interval.GetSemitonesSum()
	if sum >= Intervals[Third][Diminished].GetSemitonesSum() && sum <= Intervals[Third][Augmented].GetSemitonesSum() {
		return Third
	}

	return Fifth
}

func (interval Interval) GetNote(sourceNote string) string {
	sourcePosition := GetNotePosition(sourceNote)
	semitones := interval.GetSemitonesSum()
	targetPosition := normalizeNotePosition(sourcePosition + semitones)

	diatonicSourcePosition := GetNotePosition(string(sourceNote[0]))
	diatonicTargetName := countNotes(diatonicSourcePosition, int(interval.GetClassification()))

	switch getBase(sourcePosition) {
	case NaturalNotesBase:
		return getName(targetPosition, diatonicTargetName, NaturalNotesBase, BemolNotesBase, SharpNotesBase)
	case SharpNotesBase:
		return getName(targetPosition, diatonicTargetName, SharpNotesBase, NaturalNotesBase, DoubleSharpNotesBase)
	case BemolNotesBase:
		return getName(targetPosition, diatonicTargetName, BemolNotesBase, DoubleBemolNotesBase, NaturalNotesBase)
	}

	return ""
}

func countNotes(sourcePosition, times int) string {
	base := int(getBase(sourcePosition))

	var note string
	for i := sourcePosition; times > 0; i++ {
		position := normalizeNotePosition(base + i)

		note = positionsNotes[position]
		if note == "" {
			continue
		}

		times--
	}

	return note
}

func getName(position int, diatonicTargetName string, center, left, right NoteBase) string {
	nameInLeft := GetNoteName(normalizeNotePosition(position + int(left)))
	nameInCenter := GetNoteName(normalizeNotePosition(position + int(center)))
	nameInRight := GetNoteName(normalizeNotePosition(position + int(right)))

	if len(nameInRight) > 0 && string(nameInRight[0]) == diatonicTargetName {
		return nameInRight
	}

	if len(nameInCenter) > 0 && string(nameInCenter[0]) == diatonicTargetName {
		return nameInCenter
	}

	return nameInLeft
}

func SumIntervals(interval1, interval2 Interval) Interval {
	result := Interval{
		ChromaticSemitones: interval1.ChromaticSemitones + interval2.ChromaticSemitones,
		DiatonicSemitones:  interval1.DiatonicSemitones + interval2.DiatonicSemitones,
	}

	return result
}
