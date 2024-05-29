package main

import (
	"strconv"
)

func openLock(deadends []string, target string) int {
	state := "0000"
	if target == state {
		return 0
	}

	used := [10000]bool{}
	dist := [10000]int{}
	blocked := [10000]bool{}
	for _, v := range deadends {
		idx, _ := strconv.Atoi(v)
		blocked[idx] = true
	}

	queue := make([]string, 0)
	queue = append(queue, state)
	for len(queue) > 0 {
		state = queue[0]
		queue = queue[1:]
		stateIdx, _ := strconv.Atoi(state)
		curDist := dist[stateIdx]
		if used[stateIdx] || blocked[stateIdx] {
			continue
		}
		used[stateIdx] = true
		for i, r := range state {
			var temp string
			if r == '9' {
				temp = state[:i] + "0" + state[i+1:]
			} else {
				temp = state[:i] + string(r+1) + state[i+1:]
			}
			queue = append(queue, temp)
			tempIdx, _ := strconv.Atoi(temp)
			if dist[tempIdx] == 0 {
				dist[tempIdx] = curDist + 1
			}

			if r == '0' {
				temp = state[:i] + "9" + state[i+1:]
			} else {
				temp = state[:i] + string(r-1) + state[i+1:]
			}
			queue = append(queue, temp)
			tempIdx, _ = strconv.Atoi(temp)
			if dist[tempIdx] == 0 {
				dist[tempIdx] = curDist + 1
			}
		}
	}
	targetIdx, _ := strconv.Atoi(target)
	if dist[targetIdx] == 0 {
		return -1
	} else {
		return dist[targetIdx]
	}
}
