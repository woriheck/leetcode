package main

func repeatedNTimes(A []int) int {
	m := map[int]int{}
	for index := 0; index < len(A); index++ {
		if m[A[index]] == 0 {
			m[A[index]]++
		} else {
			return A[index]
		}
	}
	return 0
}

func main() {
}
