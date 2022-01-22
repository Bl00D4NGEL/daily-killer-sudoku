package main

import (
	"fmt"
)

const GridText = "AZoAAAAAAAAAAQABAAMABAACAAIABgAHAAgAAQADAAMABAAFAAUABgAHAAgACAAJAAkADQAKAAoABgALAA8AFAAQAAwADQAOAA4AEgALAA8AFAAQAAwAEQAVAA4AEgASABMAFAAQAAwAEQAVABoAFgAWABMAFwAXABgAGAAZABoAGwAbABwAHQAdAB8AGAAZABoAHgAeABwAHQAfAB8AIAAgACALFg0MCAYJDBEKCQ8QCw0FDgkODQ8RBgkWAxEQCRMJCBA="

const SolutionText = "AJoGAwIJBwQBBQgBCAkGAwUHBAIFBAcBAggJBgMDCQEIBAYCBwUCBgQFCQcDCAEHBQgCAQMGCQQEAgUDBgkIAQcJBwMECAEFAgYIAQYHBQIEAwk="

/**
 DKS.puzzle = new DKS.Puzzle(
	{
		"id": 23190,
		"date": "2022-01-19",
		"difficulty": 2,
		"board_base64": "AZoAAAAAAAAAAQABAAMABAACAAIABgAHAAgAAQADAAMABAAFAAUABgAHAAgACAAJAAkADQAKAAoABgALAA8AFAAQAAwADQAOAA4AEgALAA8AFAAQAAwAEQAVAA4AEgASABMAFAAQAAwAEQAVABoAFgAWABMAFwAXABgAGAAZABoAGwAbABwAHQAdAB8AGAAZABoAHgAeABwAHQAfAB8AIAAgACALFg0MCAYJDBEKCQ8QCw0FDgkODQ8RBgkWAxEQCRMJCBA=",
		"solution_base64": "AJoGAwIJBwQBBQgBCAkGAwUHBAIFBAcBAggJBgMDCQEIBAYCBwUCBgQFCQcDCAEHBQgCAQMGCQQEAgUDBgkIAQcJBwMECAEFAgYIAQYHBQIEAwk=",
		"puzzle_type": 1
	});
*/

type GroupedFields struct {
	groupId int
	fields  []*PuzzleField
}

func main() {
	puzzle := FromString(GridText)

	// Map for group id -> fields associated with that group
	var puzzleGroups = make(map[int]*GroupedFields)
	for _, cage := range puzzle.cages {
		for _, field := range cage.fields {
			val, ok := puzzleGroups[field.group.id]
			if !ok {
				val = &GroupedFields{}
			}
			val.fields = append(val.fields, field)
			val.groupId = field.group.id
			puzzleGroups[field.group.id] = val
		}
	}

	// Determine number by checking if grid contains only one "ungrouped" field
	for cageId, cage := range puzzle.cages {
		ungroupedFields := getUngroupedFields(cageId, cage, puzzleGroups)

		if len(ungroupedFields) != 1 {
			continue
		}

		var groupsInCageSum = make(map[int]int)
		for _, field := range cage.fields {
			if areAllFieldsOfGroupInCage(field.group.id, cageId, puzzleGroups[field.group.id].fields) {
				groupsInCageSum[field.group.id] = field.sum
			}
		}
		cageSum := 0
		for _, sum := range groupsInCageSum {
			cageSum += int(sum)
		}
		ungroupedField := ungroupedFields[0]
		ungroupedField.value = 45 - cageSum
	}

	// If groups have only one missing value fill it with the sum of other fields
	for _, group := range puzzleGroups {
		hasValue := false
		for _, field := range group.fields {
			if field.value != 0 {
				hasValue = true
				break
			}
		}

		if hasValue {
			fieldSum := 0
			var unsetFields []*PuzzleField
			for _, field := range group.fields {
				if field.value != 0 {
					fieldSum += field.value
				} else {
					unsetFields = append(unsetFields, field)
				}
			}
			if len(unsetFields) == 1 {
				unsetFields[0].value = unsetFields[0].sum - fieldSum
			}
		}
	}

	// TODO: set possible combinations and remove those that aren't possible (e.g. due to cage already containing number)
	// knownCombinations := make(map[string][][]int)
	// for _, cage := range puzzle.cages {
	// 	for _, field := range cage.fields {
	// 		field.group.combinations = GetPossibilities(field.sum, len(puzzleGroups[field.group.id].fields), knownCombinations).combinations
	// 	}
	// }

	PrintPuzzle(puzzle)
}

func getUngroupedFields(cageId int, cage *Cage, puzzleGroups map[int]*GroupedFields) []*PuzzleField {
	var ungroupedFieldsGroupIds []int
	for _, field := range cage.fields {
		for _, v := range puzzleGroups[field.group.id].fields {
			if v.cageId != cageId {
				ungroupedFieldsGroupIds = append(ungroupedFieldsGroupIds, v.group.id)
			}
		}
	}

	var fields []*PuzzleField
	for _, groupId := range ungroupedFieldsGroupIds {
		for _, field := range puzzleGroups[groupId].fields {
			if field.cageId == cageId {
				fields = append(fields, field)
			}
		}
	}
	return fields
}

func areAllFieldsOfGroupInCage(groupId int, cageId int, fields []*PuzzleField) bool {
	for _, v := range fields {
		if v.cageId != cageId {
			return false
		}
	}

	return true
}

func PrintPuzzle(p *Puzzle) {
	var fields [81]PuzzleField
	for i := 0; i < 81; i++ {
		if i%3 == 0 {
			fieldIndex := (i%27 - i%9) / 3
			cageId := (i-i%9)/9 - (i % 27 / 9) + i%9/3
			fields[i] = *p.cages[cageId].fields[fieldIndex]
			fields[i+1] = *p.cages[cageId].fields[fieldIndex+1]
			fields[i+2] = *p.cages[cageId].fields[fieldIndex+2]
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Print("|")
		for j := 0; j < 9; j++ {
			field := fields[i*9+j]
			if field.value != 0 {
				fmt.Printf("   %2d    |", field.value)
			} else {
				fmt.Printf(
					" %2d (%2d) |",
					field.sum,
					field.group.id,
				)
			}
		}
		fmt.Println("")
	}
}
