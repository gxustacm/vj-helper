package model

import (
	"testing"
)

func TestFindProblemList(t *testing.T) {
	data, err := findProblemList([]string{"math"})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}
}

func TestFindProblemListByRating(t *testing.T) {
	data, err := findProblemListByCondition([]string{"math", "greedy"}, 800, 1400)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}
}

func TestFindOneProblemByRating(t *testing.T) {
	data, err := FindOneProblemByCondition(ProblemCondition{Tags: []string{"math", "greedy"}, Low: 800, High: 1400})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}
}

func TestFindSomeProblemByRating(t *testing.T) {
	data, err := FindSomeProblemByCondition([]ProblemCondition{ProblemCondition{Tags: []string{"math", "greedy"}, Low: 800, High: 1400}, ProblemCondition{Tags: []string{"math", "greedy"}, Low: 800, High: 1600}})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(data)
	}
}
