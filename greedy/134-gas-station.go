package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < n; i++ {
		if gas[i] < cost[i] {
			continue
		}
		gasLeft, idx := gas[i], i
		for j := i + 1; j <= n || j <= i; j++ {
			if gasLeft < cost[idx] {
				break
			}
			if j == n {
				j = 0
			}
			if j == i {
				return i
			}
			gasLeft = gasLeft - cost[idx] + gas[j]
			idx = j
		}
		if idx < i {
			break
		}
		i = idx
	}

	return -1
}
