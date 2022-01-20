package main

import "fmt"

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
func main() {
	puzzle := FromString(GridText)

	for i := 0; i < 9; i++ {
		fmt.Print("|")
		for j := 0; j < 9; j++ {
			fmt.Printf(
				" %2d (%2d) |",
				puzzle.fields[i*9+j].sum,
				puzzle.fields[i*9+j].groupId,
			)
		}
		fmt.Println("")
	}
}
