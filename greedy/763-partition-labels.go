package main

func partitionLabels(s string) []int {
	res := []int{}
	hm := make([]int, 26)
	for i, r := range s {
		hm[r-'a'] = i
	}
	for i := 0; i < len(s); i++ {
		r := hm[s[i]-'a']
		for j := i + 1; j <= r; j++ {
			if hm[s[j]-'a'] > r {
				r = hm[s[j]-'a']
			}
		}
		res = append(res, r-i+1)
		i = r
	}
	return res
}
