package main

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	result := [][]int{{r0, c0}}

	// NESW start from E
	direction := 1
	dy := []int{-1, 0, 1, 0}
	dx := []int{0, 1, 0, -1}

	y := r0
	x := c0

	step := 0
	i := 0

	for R*C > len(result) {

		// pattern 1E1S 2W2N 3E3S...
		if i%2 == 0 {
			step++
		}

		for j := 0; j < step; j++ {
			// move direction
			y += dy[direction]
			x += dx[direction]

			// check is within scope
			if y >= 0 && y < R && x >= 0 && x < C {
				result = append(result, []int{y, x})
			}
		}

		// change direction
		direction = (direction + 1) % 4

		i++
	}

	return result
}

func main() {
}
