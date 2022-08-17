package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		exp     string
		wantRes int
	}{
		{
			name:    "expr1",
			exp:     "3+2*2",
			wantRes: 7,
		},
		{
			name:    "expr2",
			exp:     "3 / 2",
			wantRes: 1,
		},
		{
			name:    "expr3",
			exp:     "3+5 / 2",
			wantRes: 5,
		},
		{
			name:    "expr4",
			exp:     "3*5-6/2",
			wantRes: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := Calculate(tt.exp)
			assert.Equal(t, tt.wantRes, gotRes)
		})
	}
}
