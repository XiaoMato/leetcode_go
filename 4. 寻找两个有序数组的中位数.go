package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2, paased, shouldPaased int

	var n int = len(nums1) + len(nums2)
	shouldPaased = n/2 + 1

	var lastOne, lastTwo int

	for paased < shouldPaased {
		println(lastOne, lastTwo)
		if l1 >= len(nums1) {
			lastTwo, lastOne = lastOne, nums2[l2]
			l2++
			paased++
			continue
		}
		if l2 >= len(nums2) {
			lastTwo, lastOne = lastOne, nums1[l1]
			l1++
			paased++
			continue
		}
		if nums1[l1] <= nums2[l2] {
			lastTwo, lastOne = lastOne, nums1[l1]
			l1++
			paased++
			continue
		}
		if nums1[l1] > nums2[l2] {
			lastTwo, lastOne = lastOne, nums2[l2]
			l2++
			paased++
			continue
		}
	}

	if n%2 == 1 {
		return float64(lastOne)
	}
	return float64(lastOne+lastTwo) / 2
}
