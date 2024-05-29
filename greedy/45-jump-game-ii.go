package main

func jump(nums []int) int {
	steps, cur, next := 0, 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[cur]+cur < i {
			cur = next
			steps++
			i--
		} else {
			if nums[i]+i > nums[next]+next {
				next = i
			}
			if i == len(nums)-1 {
				steps++
			}
		}
	}
	return steps
}
