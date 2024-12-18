package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findTrails(topoMap [][]byte, pos [2]int, rows int, cols int, visited map[[2]int]int) int {
	if _, ok := visited[pos]; ok {
		return 0
	}
	currHeight := topoMap[pos[0]][pos[1]]
	visited[pos] = 1
	if topoMap[pos[0]][pos[1]] == '9' {
		return 1
	}
	res := 0
	for _, diff := range [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	} {
		row, col := pos[0]+diff[0], pos[1]+diff[1]
		if row >= 0 && row < rows && col >= 0 && col < cols {
			if topoMap[row][col]-currHeight == 1 {
				// fmt.Printf("Point: (%d, %d), height: %d\n", row, col, topoMap[row][col])
				res += findTrails(topoMap, [2]int{row, col}, rows, cols, visited)
			}
		}
	}
	return res
}

func main() {
	file_path := os.Args[1]
	data, err := os.ReadFile(file_path)
	check(err)

	topoMap := make([][]byte, 0)
	for _, line := range strings.Fields(string(data)) {
		bLine := make([]byte, len(line))
		copy(bLine, []byte(line))
		topoMap = append(topoMap, bLine)
	}

	res := 0
	rows := len(topoMap)
	for i := range len(topoMap) {
		cols := len(topoMap[i])
		for j := range len(topoMap[i]) {
			if topoMap[i][j] == '0' {
				visited := make(map[[2]int]int)
				score := findTrails(topoMap, [2]int{i, j}, rows, cols, visited)
				res += score
			}
		}
	}
	fmt.Printf("Total score is %d\n", res)
}
