package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	r1 := s4(nums1, nums2)
	if r1 != -1 {
		return r1
	}
	return s4(nums2, nums1)
}

func s4(nums1, nums2 []int) float64 {
	l1, r1 := 0, len(nums1)
	target := (len(nums1) + len(nums2)) / 2
	singleMedian := (len(nums1)+len(nums2))%2 == 1
	for l1 < r1 {
		m1 := (r1-l1)/2 + l1
		l2, r2 := 0, len(nums2)
		for l2 < r2 {
			m2 := (r2-l2)/2 + l2
			if nums2[m2] > nums1[m1] {
				r2 = m2
			} else {
				l2 = m2 + 1
			}
		}
		if m1+r2 == target {
			if !singleMedian {
				if r2-1 < 0 || (r2-1 >= 0 && m1-1 >= 0 && nums2[r2-1] < nums1[m1-1]) {
					return float64(nums1[m1]+nums1[m1-1]) / 2
				} else {
					return float64(nums1[m1]+nums2[r2-1]) / 2
				}
			} else {
				return float64(nums1[m1])
			}
		} else if m1+r2 > target {
			if target-m1 >= 0 && target-m1 < len(nums2) && nums2[target-m1] == nums1[m1] {
				if singleMedian {
					return float64(nums1[m1])
				} else {
					return float64(nums1[m1]+nums2[target-m1-1]) / 2
				}
			}
			r1 = m1
		} else {
			l1 = m1 + 1
		}
	}
	return -1
}
