package main

import (
	"encoding/base64"
)

type PuzzleField struct {
	sum     byte
	groupId byte
}

type Puzzle struct {
	fields [81]PuzzleField
}

func FromString(s string) (puzzle Puzzle) {
	out, _ := base64.StdEncoding.DecodeString(GridText)

	var fields []byte
	cageData := out[2:164]

	for i := 0; i < len(cageData); i += 2 {
		fields = append(fields, cageData[i+1])
	}

	cageSums := out[164:]
	for i, v := range fields {
		puzzle.fields[i] = PuzzleField{
			sum:     cageSums[v],
			groupId: v,
		}
	}

	return
}
