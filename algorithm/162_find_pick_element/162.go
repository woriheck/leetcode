package main

func findPeakIndex(nums []int, start int, end int) int {
	r := end - start
	center := start + (r / 2) + (r % 2)

	if start == end {
		return start
	}

	if nums[center] < nums[center-1] {
		return findPeakIndex(nums, start, center-1)
	} else if center == end {
		return center
	} else if nums[center] < nums[center+1] {
		return findPeakIndex(nums, center+1, end)
	} else {
		return center
	}
}

func findPeakElement(nums []int) int {
	return findPeakIndex(nums, 0, len(nums)-1)
}
