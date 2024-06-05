package main

func commonChars(words []string) []string {
	count := [26]int{}
	for _, char := range words[0] {
		count[char-'a']++
	}
	for i := 1; i < len(words); i++ {
		countCurr := [26]int{}
		for _, char := range words[i] {
			countCurr[char-'a']++
		}
		for j := range 26 {
			count[j] = min(count[j], countCurr[j])
		}
	}

	res := []string{}
	for ch, cnt := range count {
		for i := 0; i < cnt; i++ {
			res = append(res, string(ch+'a'))
		}
	}
	return res
}
