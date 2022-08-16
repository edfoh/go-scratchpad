package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	testCases := []struct {
		desc      string
		nums      []int
		k         int
		wantValue int
	}{
		{
			desc:      "input 1",
			nums:      []int{3, 2, 1, 5, 6, 4},
			k:         2,
			wantValue: 5,
		},
		{
			desc:      "input 2",
			nums:      []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:         4,
			wantValue: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			gotValue := FindKthLargest(tC.nums, tC.k)
			assert.Equal(t, tC.wantValue, gotValue)
		})
	}
}
