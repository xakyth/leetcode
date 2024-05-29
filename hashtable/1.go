package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int][]int)
	for idx, v := range nums {
		seen[v] = append(seen[v], idx)
	}
	for k := range seen {
		if k == target-k {
			if len(seen[k]) >= 2 {
				return []int{seen[k][0], seen[k][1]}
			}
		} else {
			if len(seen[target-k]) > 0 {
				return []int{seen[k][0], seen[target-k][0]}
			}
		}

	}
	return nil
}
