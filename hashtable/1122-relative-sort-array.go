package main

func relativeSortArray(arr1 []int, arr2 []int) []int {
	cnt := make([]int, 1001)
	for _, v := range arr1 {
		cnt[v]++
	}
	res := []int{}

	for _, v := range arr2 {
		c := cnt[v]
		cnt[v] = 0
		for i := 0; i < c; i++ {
			res = append(res, v)
		}
	}
	for i, c := range cnt {
		for j := 0; j < c; j++ {
			res = append(res, i)
		}
	}

	return res
}
