package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc        string
		input       [][]string
		wantResults []string
	}{
		{
			desc: "test 1",
			input: [][]string{
				{"A", "B", "C"},
				{"A", "B", "C"},
			},
			wantResults: []string{"A", "B", "C"},
		},
		{
			desc: "test 1",
			input: [][]string{
				{"C", "A", "B"},
				{"A", "C", "B"},
				{"C", "A", "B"},
				{"A", "C", "B"},
				// A: 5, B: 2, C:5
			},
			wantResults: []string{"A", "C", "B"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			gotResults := FindWinner(tC.input)

			assert.Equal(t, tC.wantResults, gotResults)

		})
	}
}
