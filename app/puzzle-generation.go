package main

import (
	"encoding/base64"
)

type Puzzle struct {
	cages [9]*Cage
}

type Cage struct {
	id     int
	fields [9]*PuzzleField
}

type PuzzleField struct {
	sum    int
	group  Group
	cageId int
	value  int
}

type Group struct {
	id int
	// combinations [][]int
}

func FromString(s string) *Puzzle {
	var puzzle = Puzzle{}
	out, _ := base64.StdEncoding.DecodeString(GridText)

	var fields []byte
	cageData := out[2:164]

	for i := 0; i < len(cageData); i += 2 {
		fields = append(fields, cageData[i+1])
	}

	cageSums := out[164:]

	for i := 0; i < len(fields); i += 3 {
		if i%3 == 0 {
			fieldIndex := (i%27 - i%9) / 3
			cageId := (i-i%9)/9 - (i % 27 / 9) + i%9/3
			cage := puzzle.cages[cageId]
			if cage == nil {
				cage = &Cage{}
			}
			cage.id = cageId

			cage.fields[fieldIndex] = &PuzzleField{
				group: Group{
					id: int(fields[i]),
				},
				cageId: cageId,
				sum:    int(cageSums[fields[i]]),
			}
			cage.fields[fieldIndex+1] = &PuzzleField{
				group: Group{
					id: int(fields[i+1]),
				},
				cageId: cageId,
				sum:    int(cageSums[fields[i+1]]),
			}
			cage.fields[fieldIndex+2] = &PuzzleField{
				group: Group{
					id: int(fields[i+2]),
				},
				cageId: cageId,
				sum:    int(cageSums[fields[i+2]]),
			}
			puzzle.cages[cageId] = cage
		}
	}

	return &puzzle
}
