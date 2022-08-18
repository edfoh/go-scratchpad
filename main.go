package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
In a special ranking system, each voter gives a rank from highest to lowest to all teams participated in the competition.

The ordering of teams is decided by who received the most position-one votes. If two or more teams tie in the first position, we consider the second position to resolve the conflict, if they tie again, we continue this process until the ties are resolved. If two or more teams are still tied after considering all positions, we rank them alphabetically based on their team letter.

Given an array of strings votes which is the votes of all voters in the ranking systems. Sort all teams according to the ranking system described above.

Return a string of all teams sorted by the ranking system.

Example 1:

Input: votes = ["ABC","ACB","ABC","ACB","ACB"]
Output: "ACB"
Explanation: Team A was ranked first place by 5 voters. No other team was voted as first place so team A is the first team.
Team B was ranked second by 2 voters and was ranked third by 3 voters.
Team C was ranked second by 3 voters and was ranked third by 2 voters.
As most of the voters ranked C second, team C is the second team and team B is the third.

https://leetcode.com/problems/rank-teams-by-votes/
*/
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
