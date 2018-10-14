package main

import (
	"fmt"
)

// DFS Depth First Search algorithm
func DFS(v int, visited []bool, rooms [][]int) {
	for _, room := range rooms[v] {
		if visited[room] == false {
			visited[room] = true
			fmt.Println(room)
			DFS(room, visited, rooms)
		}
	}
}

func canVisitAllRooms(rooms [][]int) bool {
	visited := make([]bool, len(rooms))
	visited[0] = true
	DFS(0, visited, rooms)

	for _, visit := range visited {
		if visit == false {
			return false
		}
	}
	return true
}

func main() {
}
