package main

import "math"

type SemitoneType string

const (
	Diatonic              SemitoneType = "diatonic"
	Chromatic             SemitoneType = "chromatic"
	UndefinedSemitoneType SemitoneType = "undefined_semitone"
)

func GetSemitoneType(source, target string) SemitoneType {
	sourcePosition, _, _ := GetNotePosition(source)
	targetPosition, _, _ := GetNotePosition(target)

	if math.Abs(float64(sourcePosition-targetPosition)) != 1 {
		return UndefinedSemitoneType
	}

	rawSource := ExtractNoteRawName(source)
	rawTarget := ExtractNoteRawName(target)

	if rawSource == rawTarget {
		return Chromatic
	}

	return Diatonic
}
