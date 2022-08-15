package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopVotedCandidate_Q(t *testing.T) {
	persons := []int{0, 1, 1, 0, 0, 1, 0}
	times := []int{0, 5, 10, 15, 20, 25, 30}
	tests := []struct {
		name       string
		persons    []int
		times      []int
		queryTime  int
		wantPerson int
	}{
		{
			name:       "test - query time 3",
			persons:    persons,
			times:      times,
			queryTime:  3,
			wantPerson: 0,
		},
		{
			name:       "test - query time 12",
			persons:    persons,
			times:      times,
			queryTime:  12,
			wantPerson: 1,
		},
		{
			name:       "test - query time 25",
			persons:    persons,
			times:      times,
			queryTime:  25,
			wantPerson: 1,
		},
		{
			name:       "test - query time 15",
			persons:    persons,
			times:      times,
			queryTime:  15,
			wantPerson: 0,
		},
		{
			name:       "test - query time 24",
			persons:    persons,
			times:      times,
			queryTime:  24,
			wantPerson: 0,
		},
		{
			name:       "test - query time 8",
			persons:    persons,
			times:      times,
			queryTime:  8,
			wantPerson: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewTopVotedCandidate(tt.persons, tt.times)
			gotPerson := sut.Q(tt.queryTime)
			assert.Equal(t, tt.wantPerson, gotPerson)
		})
	}
}
