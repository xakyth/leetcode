package main

func isInterleave(s1, s2, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	if len(s1) == 0 {
		return s2 == s3
	} else if len(s2) == 0 {
		return s1 == s3
	}
	hm := map[[2]int]bool{}
	if len(s1) > 0 && s1[0] == s3[0] {
		hm[[2]int{0, -1}] = true
	}
	if len(s2) > 0 && s2[0] == s3[0] {
		hm[[2]int{-1, 0}] = true
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		key := [2]int{i, j}
		valid, exist := hm[key]
		if !exist {
			if i >= 0 && s1[i] == s3[i+j+1] {
				valid = dfs(i-1, j)
			}
			if j >= 0 && s2[j] == s3[i+j+1] {
				valid = valid || dfs(i, j-1)
			}
			hm[key] = valid
		}
		return valid
	}
	return dfs(len(s1)-1, len(s2)-1)
}
