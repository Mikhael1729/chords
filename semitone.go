package main

type SemitoneType string

const (
	Diatonic              SemitoneType = "diatonic"
	Chromatic             SemitoneType = "chromatic"
	UndefinedSemitoneType SemitoneType = "undefined_semitone"
)

func GetSemitoneType(sourcePosition, targetPosition int) SemitoneType {
	sourceName := GetNoteName(sourcePosition)
	targetName := GetNoteName(targetPosition)

	if sourceName[0] != targetName[0] {
		return Diatonic
	}

	return Chromatic
}
