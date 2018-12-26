package main

func singleNonDuplicate(nums []int) int {
	ans := nums[0]

	for index := 0; index < len(nums)-1; index++ {

		if nums[index+1] != ans {
			break
		}

		for nums[index] == ans {
			index++
		}
		ans = nums[index]
		index--
	}
	return ans
}

func main() {
	singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
}
