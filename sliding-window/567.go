package main

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	s1cnt := [26]int{}
	for _, r := range s1 {
		s1cnt[r-'a']++
	}
	s2cnt := [26]int{}
	k := len(s1)
	for i := 0; i < k; i++ {
		s2cnt[s2[i]-'a']++
	}
	if check(s1cnt, s2cnt) {
		return true
	}
	for i := 1; i <= len(s2)-k; i++ {
		s2cnt[s2[i-1]-'a']--
		s2cnt[s2[i+k-1]-'a']++
		if check(s1cnt, s2cnt) {
			return true
		}
	}
	return false
}

func check(s1cnt, s2cnt [26]int) bool {
	for i := range s1cnt {
		if s1cnt[i] > 0 && s1cnt[i] != s2cnt[i] {
			return false
		}
	}
	return true
}
