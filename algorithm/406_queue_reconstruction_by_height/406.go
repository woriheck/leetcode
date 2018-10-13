package main

func reconstructQueue(people [][]int) [][]int {
	result := make([][]int, len(people))

	for len(people) > 0 {
		lowest := people[0][0]
		lIndex := 0
		for index := 1; index < len(people); index++ {
			height := people[index][0]
			if lowest > height {
				lowest = height
				lIndex = index
			}
		}

		position := people[lIndex][1]
		for i := 0; i < len(result); i++ {

			if len(result[i]) == 0 && position == 0 {
				result[i] = people[lIndex]
			}

			if len(result[i]) == 0 || result[i][0] >= people[lIndex][0] {
				position--
			}

		}
		people = append(people[:lIndex], people[lIndex+1:]...)
	}

	return result
}

func main() {
	reconstructQueue([][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	})
}
