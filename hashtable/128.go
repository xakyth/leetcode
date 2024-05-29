package main

func longestConsecutive(nums []int) int {
	numsMap := map[int]struct{}{}

	for _, n := range nums {
		numsMap[n] = struct{}{}
	}

	longestSeq := 0
	for k := range numsMap {
		cnt := 0
		delete(numsMap, k)
		for i := k + 1; true; i++ {
			if _, ok := numsMap[i]; ok {
				delete(numsMap, i)
			} else {
				cnt += i - k
				break
			}
		}
		for i := k - 1; true; i-- {
			if _, ok := numsMap[i]; ok {
				delete(numsMap, i)
			} else {
				cnt += k - i - 1
				break
			}
		}
		if cnt > longestSeq {
			longestSeq = cnt
		}
	}
	return longestSeq
}
