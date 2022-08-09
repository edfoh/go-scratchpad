package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeFunc(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SomeFunc(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
		})
	}
}
