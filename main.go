package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

var (
	maxWinners = 3
)

func FindWinner(votes [][]string) []string {
	voting := NewVoting(maxWinners)
	for _, v := range votes {
		voting.AddVote(v)
	}

	projectVotes := voting.FindWinner()

	var projects []string
	for _, pv := range projectVotes {
		projects = append(projects, pv.projectName)
	}
	return projects
}

type Voting struct {
	size         int
	projectVotes map[string][]int
}

func NewVoting(size int) *Voting {
	return &Voting{
		size:         size,
		projectVotes: map[string][]int{},
	}
}

func (v *Voting) AddVote(votes []string) {
	for index, vote := range votes {
		if _, ok := v.projectVotes[vote]; !ok {
			v.projectVotes[vote] = make([]int, v.size)
		}
		v.projectVotes[vote][index] += 1
	}
}

func (v *Voting) FindWinner() []ProjectVote {
	result := map[string]int{}

	for projectName, votes := range v.projectVotes {
		score := calculateScore(votes)
		if _, ok := result[projectName]; !ok {
			result[projectName] = 0
		}
		result[projectName] = score
	}

	var projectVotes []ProjectVote
	for k, v := range result {
		projectVotes = append(projectVotes, ProjectVote{k, v})
	}

	sort.Slice(projectVotes, func(i, j int) bool {
		left := projectVotes[i]
		right := projectVotes[j]
		return left.score > right.score
	})

	return projectVotes
}

type ProjectVote struct {
	projectName string
	score       int
}

func calculateScore(score []int) int {
	result := 0
	weight := 3
	for _, s := range score {
		result += s * weight
		weight -= 1
	}
	return result
}

// type Score struct {
// 	score []int
// }

// // Len is the number of elements in the collection.
// func (s *Score) Len() int {
// 	return len(s.score)
// }

// func (s *Score) Less(i int, j int) bool {
// 	iScore := s.score[i]
// 	jScore := s.score[j]
// 	for _, v := range v {

// 	}
// }

// // Swap swaps the elements with indexes i and j.
// func (s *Score) Swap(i int, j int) {
// 	s.score[i], s.score[j] = s.score[j], s.score[i]
// }
