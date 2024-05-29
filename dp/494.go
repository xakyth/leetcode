package main

func findTargetSumWays(nums []int, target int) int {
	row := make(map[int]int)
	row[0] = 1
	for _, num := range nums {
		nextRow := make(map[int]int)
		for sum, cnt := range row {
			nextRow[sum-num] += cnt
			nextRow[sum+num] += cnt
		}
		row = nextRow
	}
	return row[target]
}
