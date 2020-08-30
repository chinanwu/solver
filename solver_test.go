package solver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	wordsList := []string{"aaaa", "aaab"}
	path, dist, err := Solve("aaaa", "aaab", wordsList, 4)
	if err != nil {
		t.Errorf("error returned from Solve(): %s", err)
	}

	if (len(wordsList) != len(path)) || (dist != 1) {
		t.Errorf("path returned: %q, expected: %q, dist: %q", path, wordsList, dist)
	}

	for i, v := range wordsList {
		if v != path[i] {
			t.Errorf("path returned: %q, expected: %q", path, wordsList)
		}
	}
}

func TestSolveSameFromTo(t *testing.T) {
	wordsList := []string{"aaaa"}
	path, dist, err := Solve("aaaa", "aaaa", wordsList, 4)

	if err != nil {
		t.Errorf("error returned from Solve(): %s", err)
	}

	if (len(wordsList) != len(path)) || (dist != 0) {
		t.Errorf("path returned: %q, expected: %q, dist: %q", path, wordsList, dist)
	}

	for i, v := range wordsList {
		if v != path[i] {
			t.Errorf("path returned: %q, expected: %q", path, wordsList)
		}
	}
}

func TestNoSolution(t *testing.T) {
	wordsList := []string{"aaaa", "bbbb"}
	path, dist, err := Solve("aaaa", "bbbb", wordsList, 4)
	if err == nil {
		t.Errorf("Path found for impossible problem: %q, with dist: %q", path, dist)
	}

	if err.Error() != "unable to find a solution going from: aaaa, to: bbbb" {
		t.Errorf("wrong error thrown during finding of an impossible problem: %q", err.Error())
	}
}

func TestSolveWithEmptyWordList(t *testing.T) {
	path, dist, err := Solve("heat", "cold", []string{}, 4)
	if err == nil {
		t.Errorf("Path found for impossible problem: %q, with dist: %q", path, dist)
	}

	if err.Error() != "unable to find the string heat in the list" {
		t.Errorf("wrong error thrown during finding of an impossible problem: %q", err.Error())
	}
}

func TestSolveWithNilWordList(t *testing.T) {
	path, dist, err := Solve("heat", "cold", nil, 4)
	if err == nil {
		t.Errorf("Path found for impossible problem: %q, with dist: %q", path, dist)
	}

	if err.Error() != "invalid array provided for words list" {
		t.Errorf("wrong error thrown during finding of an impossible problem: %q", err.Error())
	}
}

func TestSolveWithShortWords(t *testing.T) {
	expected := []string{"a", "b"}
	path, dist, err := Solve("a", "b", []string{"a", "b", "c"}, 1)

	if err != nil {
		t.Errorf("error returned from Solve(): %s", err)
	}

	if (len(expected) != len(path)) || (dist != 1) {
		t.Errorf("path returned: %q, expected: %q, dist: %q", path, expected, dist)
	}

	for i, v := range expected {
		if v != path[i] {
			t.Errorf("path returned: %q, expected: %q", path, expected)
		}
	}
}

func TestSolveWithLongWords(t *testing.T) {
	expected := []string{"aaaaa", "aaaab"}
	path, dist, err := Solve("aaaaa", "aaaab", expected, 5)

	if err != nil {
		t.Errorf("error returned from Solve(): %s", err)
	}

	if (len(expected) != len(path)) || (dist != 1) {
		t.Errorf("path returned: %q, expected: %q, dist: %q", path, expected, dist)
	}

	for i, v := range expected {
		if v != path[i] {
			t.Errorf("path returned: %q, expected: %q", path, expected)
		}
	}
}
