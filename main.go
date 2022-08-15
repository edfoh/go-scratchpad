package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
}

/*
Online Election https://leetcode.com/problems/online-election/

You are given two integer arrays persons and times. In an election, the ith vote was cast for persons[i] at time times[i].

For each query at a time t, find the person that was leading the election at time t. Votes cast at time t will count towards our query. In the case of a tie, the most recent vote (among tied candidates) wins.

Implement the TopVotedCandidate class:

TopVotedCandidate(int[] persons, int[] times) Initializes the object with the persons and times arrays.
int q(int t) Returns the number of the person that was leading the election at time t according to the mentioned rules.

persons: [0, 1, 1, 0, 0, 1, 0]. times: [0, 5, 10, 15, 20, 25, 30]
*/
type TopVotedCandidate struct {
	times                []int
	leadingPersonPerTime []int
	totalVotesByPerson   map[int]int //key: person, calue number of votes
}

func NewTopVotedCandidate(persons []int, times []int) *TopVotedCandidate {
	currentLeader := persons[0]
	leadingPersonPerTime := make([]int, len(times))
	totalVotesByPerson := make(map[int]int)

	// assume lens are same
	for i, person := range persons {
		if _, ok := totalVotesByPerson[person]; !ok {
			totalVotesByPerson[person] = 0
		}
		totalVotesByPerson[person] += 1 // increase persons vote by 1

		// find current leader
		if totalVotesByPerson[person] >= totalVotesByPerson[currentLeader] {
			currentLeader = person
		}
		leadingPersonPerTime[i] = currentLeader
	}

	return &TopVotedCandidate{
		leadingPersonPerTime: leadingPersonPerTime,
		totalVotesByPerson:   totalVotesByPerson,
		times:                times,
	}
}

// Q Returns the number of the person that was leading the election at time t according to the mentioned rules.
func (c *TopVotedCandidate) Q(t int) int {
	timeIndex := c.binarySearchTimes(c.times, 0, len(c.times)-1, t)
	if timeIndex != -1 {
		return c.leadingPersonPerTime[timeIndex]
	}
	return -1
}

func (c *TopVotedCandidate) binarySearchTimes(times []int, low int, high int, t int) int {
	if low <= high {
		mid := (low + high) / 2

		midVal := times[mid]

		if midVal == t {
			return mid
		}

		// when low == high, mid is same value as either. If midVal check fails, we default to the prev value
		if low == high {
			return min(low - 1)
		}

		if t < midVal {
			return c.binarySearchTimes(times, low, mid-1, t)
		}
		return c.binarySearchTimes(times, mid+1, high, t)
	}
	return -1
}

func min(x int) int {
	if x < 0 {
		return 0
	}
	return x
}
