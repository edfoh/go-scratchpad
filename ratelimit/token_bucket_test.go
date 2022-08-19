package ratelimit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var fakeNow = time.Date(2022, 01, 01, 12, 0, 0, 0, time.UTC)

func TestTokenBucket(t *testing.T) {
	testCases := []struct {
		desc                string
		capacity            int
		refillRatePerSecond float64
		fakeTimeAhead       []time.Time
		takeN               []int
		wantSuccess         []bool
	}{
		{
			desc:                "takeN within capacity works",
			capacity:            10,
			refillRatePerSecond: 1,
			fakeTimeAhead:       []time.Time{fakeNow},
			takeN:               []int{5},
			wantSuccess:         []bool{true},
		},
		{
			desc:                "takeN over capacity should return false",
			capacity:            10,
			refillRatePerSecond: 1,
			fakeTimeAhead:       []time.Time{fakeNow},
			takeN:               []int{11},
			wantSuccess:         []bool{false},
		},
		{
			desc:                "takeN twice in the same time region",
			capacity:            10,
			refillRatePerSecond: 1,
			fakeTimeAhead:       []time.Time{fakeNow, fakeNow.Add(5 * time.Second)},
			takeN:               []int{50, 20},
			wantSuccess:         []bool{false, false},
		},
		{
			desc:                "takeN, refill, takeN again should work",
			capacity:            10,
			refillRatePerSecond: 1,
			fakeTimeAhead:       []time.Time{fakeNow, fakeNow.Add(10 * time.Second)},
			takeN:               []int{10, 10},
			wantSuccess:         []bool{true, true},
		},
		{
			desc:                "takeN, refill, takeN again over capacity should return false",
			capacity:            10,
			refillRatePerSecond: 1,
			fakeTimeAhead:       []time.Time{fakeNow, fakeNow.Add(10 * time.Second)},
			takeN:               []int{10, 11},
			wantSuccess:         []bool{true, false},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			teardown := fakeTimeSetup()
			defer teardown()

			bucket := NewTokenBucket(tC.capacity, tC.refillRatePerSecond)

			for index, fakeTime := range tC.fakeTimeAhead {
				setFakeTime(fakeTime)

				gotSuccess := bucket.TakeN(tC.takeN[index])
				assert.Equal(t, tC.wantSuccess[index], gotSuccess)
			}
		})
	}
}

func setFakeTime(t time.Time) {
	fakeNow = t
}

func fakeTimeSetup() func() {
	now = func() time.Time {
		return fakeNow
	}

	return func() { now = time.Now }
}
