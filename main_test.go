package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRankTeamsByVotes(t *testing.T) {
	type args struct {
		votes []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				votes: []string{"ABC", "ACB", "ABC", "ACB", "ACB"},
			},
			want: "ACB",
		},
		{
			name: "test2",
			args: args{
				votes: []string{"WXYZ", "XYZW"},
			},
			want: "XWYZ",
		},
		{
			name: "test3",
			args: args{
				votes: []string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"},
			},
			want: "ZMNAGUEDSJYLBOPHRQICWFXTVK",
		},
		{
			name: "test4",
			args: args{
				votes: []string{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RankTeamsByVotes(tt.args.votes)
			assert.Equal(t, tt.want, got)
		})
	}
}
