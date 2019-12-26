package peak

func search(nums []int, start int, end int) int {
	if start == end {
		return start
	}

	center := (start + end) / 2

	// Look to the left
	if nums[center] > nums[center+1] {
		return search(nums, start, center)
	}

	// Look to the right
	return search(nums, center+1, end)
}

func findPeakElement(nums []int) int {
	return search(nums, 0, len(nums)-1)
}
