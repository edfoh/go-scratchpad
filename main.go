package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

func RankTeamsByVotes(votes []string) string {
	voting := NewVoting(votes)
	return voting.Sort()
}

type Voting struct {
	teams     []rune
	teamVotes map[rune][]int
}

func NewVoting(votes []string) *Voting {
	if len(votes) == 0 {
		return new(Voting)
	}
	teams := []rune(votes[0])

	teamVotes := map[rune][]int{}
	for _, vote := range votes {
		for position, team := range vote {
			if _, ok := teamVotes[team]; !ok {
				teamVotes[team] = make([]int, len(teams))
			}
			teamVotes[team][position] += 1
		}
	}

	return &Voting{
		teams:     teams,
		teamVotes: teamVotes,
	}
}

func (v *Voting) Sort() string {
	sort.Sort(v)
	fmt.Printf("%+v", v)
	return string(v.teams)
}

// Len is the number of elements in the collection.
func (v *Voting) Len() int {
	return len(v.teams)
}

func (v *Voting) Less(i int, j int) bool {
	iTeam := v.teams[i]
	jTeam := v.teams[j]
	iScores := v.teamVotes[iTeam]
	jScores := v.teamVotes[jTeam]
	for i := 0; i < len(iScores); i++ {
		if iScores[i] > jScores[i] {
			return true
		} else if iScores[i] < jScores[i] {
			return false
		}
	}
	return iTeam < jTeam
}

func (v *Voting) Swap(i int, j int) {
	iTeam := v.teams[i]
	v.teams[i] = v.teams[j]
	v.teams[j] = iTeam
}
