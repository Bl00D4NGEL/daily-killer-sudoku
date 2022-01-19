package main

import (
	"encoding/base64"
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
func main() {
	out, _ := base64.StdEncoding.DecodeString(GridText)

	var fields []byte
	cageData := out[2:164]
	for i := 0; i < len(cageData); i += 2 {
		fields = append(fields, cageData[i+1])
	}

	cageSums := out[164:]

	for i := 0; i < 9; i++ {
		fmt.Print("|")
		for j := 0; j < 9; j++ {
			fmt.Printf(
				" %2d (%2d) |",
				cageSums[fields[i*9+j]],
				fields[i*9+j],
			)
		}
		fmt.Println("")
	}
}
