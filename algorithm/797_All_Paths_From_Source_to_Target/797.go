package main

var result [][]int

// DFS Depth First Search algorithm
func DFS(v int, path []int, graph [][]int) {
	if v == (len(graph) - 1) {
		// preprend zero to path
		path = append([]int{0}, path...)
		result = append(result, path)
		return
	}

	for _, node := range graph[v] {
		DFS(node, append(path, node), graph)
	}
}

func allPathsSourceTarget(graph [][]int) [][]int {
	result = result[:0]
	DFS(0, []int{}, graph)
	return result
}

func main() {
}
