package main

func sortArrayByParity(A []int) []int {
	for current := 0; current < len(A); current++ {
		for s := current + 1; s < len(A); s++ {
			if A[s]%2 == 0 {
				A[current], A[s] = A[s], A[current]
			}
		}
	}
	return A
}

func main() {
	sortArrayByParity([]int{3, 1, 2, 4})
}
