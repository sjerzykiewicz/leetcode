func mergeArrays(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	if nums1[0] <= nums2[0] {
		return append([]int{nums1[0]}, mergeArrays(nums1[1:], nums2)...)
	}
	return append([]int{nums2[0]}, mergeArrays(nums1, nums2[1:])...)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	mergedArrays := mergeArrays(nums1, nums2)
	l := len(mergedArrays)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		return float64(mergedArrays[l/2-1]+mergedArrays[l/2]) / 2.0
	} else {
		return float64(mergedArrays[l/2])
	}
}