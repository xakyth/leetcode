package main

func mergeTriplets(triplets [][]int, target []int) bool {
	t1, t2, t3 := false, false, false
	for _, triplet := range triplets {
		if triplet[0] == target[0] && triplet[1] <= target[1] && triplet[2] <= target[2] {
			t1 = true
		}
		if triplet[0] <= target[0] && triplet[1] == target[1] && triplet[2] <= target[2] {
			t2 = true
		}
		if triplet[0] <= target[0] && triplet[1] <= target[1] && triplet[2] == target[2] {
			t3 = true
		}
	}

	return t1 && t2 && t3
}
