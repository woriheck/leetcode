package main

func flipAndInvertImage(A [][]int) [][]int {
	var image [][]int
	image = make([][]int, len(A))

	for i := 0; i < len(A); i++ {
		image[i] = make([]int, len(A[i]))
		k := 0
		for j := len(A[i]) - 1; j >= 0; j-- {
			image[i][k] = A[i][j] ^ 1
			k++
		}
	}
	return image
}

func main() {
	flipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	})
}
