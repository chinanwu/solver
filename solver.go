package solver

import (
	"errors"
	"github.com/yourbasic/graph"
)

// Params:
// 		from: string, the word to start from. Needs to have the same length as to
//		to: string, the word to get to. Needs to have the same length as from
//		wordsList: []string, a list of string that can be used as steps to get to 'to' from 'from', MUST be the same length as them
//			- Solver expects the wordsList to only include words that are of len 4
// Returns:
//		The solution, the number of words it takes, or an error
func Solve(from, to string, arr []string, stringLen int) ([]string, int64, error) {
	if len(from) != len(to) || len(from) != stringLen || len(to) != stringLen {
		return nil, -1, errors.New("words provided must both be the same length")
	}

	if arr == nil {
		return nil, -1, errors.New("invalid array provided for words list")
	}

	max := len(arr)

	// Build graph of words that are one of from one another
	g := graph.New(max)
	for i, elem := range arr {
		for ci, comp := range arr {
			diff, err := numOfDiffChars(elem, comp, stringLen)
			if err != nil {
				return nil, -1, err
			}

			if diff == 1 {
				g.AddBothCost(i, ci, 1)
			}
		}
	}

	// This is due to the inputs expected from ShortestPath,
	// and the fact that graph is not a <key, val> graph where
	// the key can be a string
	fromI, err := findInd(from, arr)
	if err != nil {
		return nil, -1, err
	}

	toI, err := findInd(to, arr)
	if err != nil {
		return nil, -1, err
	}

	pathI, dist := graph.ShortestPath(g, fromI, toI)

	if dist == -1 {
		return nil, dist, errors.New("unable to find a solution going from: " + from + ", to: " + to)
	}

	if dist == 0 {
		if from == to { //ease of mind check
			return []string{from}, 0, nil
		}
	}

	var path []string
	for _, node := range pathI {
		path = append(path, arr[node])
	}

	return path, dist, nil
}

func findInd(finding string, slice []string) (int, error) {
	for i, elem := range slice {
		if elem == finding {
			return i, nil
		}
	}
	return -1, errors.New("unable to find the string " + finding + " in the list")
}

func numOfDiffChars(a, b string, stringLen int) (int, error) {
	if len(a) != stringLen || len(b) != stringLen {
		return -1, errors.New("Strings " + a + " and " + b + " have different lengths")
	}

	result := 0
	for i := 0; i < stringLen; i++ {
		if a[i] != b[i] {
			result++
		}
	}
	return result, nil
}
