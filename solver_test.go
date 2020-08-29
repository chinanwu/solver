package solver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	wordsList := []string{"aaaa", "aaab"}
	path, dist, err := Solve("aaaa", "aaab", wordsList)
	if err != nil {
		t.Errorf("Error returned from Solve(): %s", err)
	}

	if (len(wordsList) != len(path)) || (dist != 1) {
		t.Errorf("Path returned: %q, expected: %q, dist: %q", path, wordsList, dist)
	}

	for i, v := range wordsList {
		if v != path[i] {
			t.Errorf("Path returned: %q, expected: %q", path, wordsList)
		}
	}
}

func TestSolveSameFromTo(t *testing.T) {
	wordsList := []string{"aaaa"}
	path, dist, err := Solve("aaaa", "aaaa", wordsList)

	if err != nil {
		t.Errorf("Error returned from Solve(): %s", err)
	}

	if (len(wordsList) != len(path)) || (dist != 0) {
		t.Errorf("Path returned: %q, expected: %q, dist: %q", path, wordsList, dist)
	}

	for i, v := range wordsList {
		if v != path[i] {
			t.Errorf("Path returned: %q, expected: %q", path, wordsList)
		}
	}
}

func TestNoSolution(t *testing.T) {
	wordsList := []string{"aaaa", "bbbb"}
	path, dist, err := Solve("aaaa", "bbbb", wordsList)
	if err == nil {
		t.Errorf("Path found for impossible problem: %q, with dist: %q", path, dist)
	}
}

func TestSolveWithDefaultWordsList(t *testing.T) {
	expected := []string{"heat", "head", "held", "hold", "cold"}
	path, dist, err := Solve("heat", "cold", nil)
	if err != nil {
		t.Errorf("Error returned from Solve(): %s", err)
	}

	if (len(expected) != len(path)) || (dist != 4) {
		t.Errorf("Path returned: %q, expected: %q, dist: %q", path, expected, dist)
	}

	for i, v := range expected {
		if v != path[i] {
			t.Errorf("Path returned: %q, expected: %q", path, expected)
		}
	}
}
