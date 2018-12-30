package main

func minDeletionSize(A []string) int {
	result := 0
	for i := 0; i < len(A[0]); i++ {
		for j := 0; j < len(A)-1; j++ {
			if A[j][i] > A[j+1][i] {
				result++
				break
			}
		}
	}
	return result
}

func main() {
	minDeletionSize([]string{"cba", "daf", "ghi"})
}
