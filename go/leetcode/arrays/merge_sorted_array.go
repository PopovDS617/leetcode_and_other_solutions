package leetcode

import "slices"

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {

	nums1 = nums1[:m]

	nums1 = append(nums1, nums2...)

	slices.Sort(nums1)

}
