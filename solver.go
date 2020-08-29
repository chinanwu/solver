package solver

import (
	"errors"
	"github.com/yourbasic/graph"
	"io/ioutil"
	"strings"
)

// Params:
// 		from: string, the word to start from
//		to: string, the word to get to
//		wordsList: []string, a list of string that can be used as steps to get to 'to' from 'from'
// Returns:
//		The solution, the number of words it takes, or an error
func Solve(from, to string, wordsList []string) ([]string, int64, error){
	var arr []string
	if wordsList == nil {
		words, err := ioutil.ReadFile("./words.txt")

		if err != nil {
			return nil, -1, err
		}

		// 3997 len
		arr = strings.Split(string(words), " ")
	} else {
		arr = wordsList
	}

	max := len(arr)

	// Build graph of words
	g := graph.New(max)
	for i, elem := range arr {
		for ci, comp := range arr {
			if numOfDiffChars(elem, comp) == 1 {
				g.AddBothCost(i, ci, 1)
			}
		}
	}

	fromI, err := findInd(from, arr)
	if err != nil {
		return nil, -1, err
	}

	toI, err := findInd(to, arr)
	if err != nil {
		return nil, -1, err
	}

	pathI, dist := graph.ShortestPath(g, fromI, toI)
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
	return -1, errors.New("Unable to find the string " + finding + " in the list")
}

func numOfDiffChars(a, b string) int {
	result := 0
	for i := 0; i < 4; i++ {
		if a[i] != b[i] {
			result++
		}
	}
	return result
}