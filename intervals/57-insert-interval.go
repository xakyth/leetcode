package main

func insert(intervals [][]int, newInterval []int) [][]int {
	lv, rv := newInterval[0], newInterval[1]

	res := [][]int{}
	inserted := false
	for i := 0; i < len(intervals); i++ {
		cur := []int{intervals[i][0], intervals[i][1]}
		//intersect
		if lv >= intervals[i][0] && lv <= intervals[i][1] {
			j := i
			for ; j < len(intervals) && intervals[j][0] <= rv; j++ {
				if intervals[j][1] >= rv {
					j++
					break
				}
			}
			cur[0] = intervals[i][0]
			cur[1] = max(rv, intervals[j-1][1])
			i += (j - i) - 1
			inserted = true
		} else if lv < intervals[i][0] && (i == 0 || lv > intervals[i-1][1]) {
			cur[0] = lv
			if intervals[i][0] > rv {
				cur[1] = rv
				res = append(res, cur)
				cur = intervals[i]
			} else {
				j := i
				for ; j < len(intervals) && intervals[j][0] <= rv; j++ {
					if intervals[j][1] >= rv {
						j++
						break
					}
				}
				cur[1] = max(rv, intervals[j-1][1])
				i += (j - i) - 1
			}
			inserted = true
		}
		res = append(res, cur)
	}
	if !inserted {
		res = append(res, newInterval)
	}
	return res
}
