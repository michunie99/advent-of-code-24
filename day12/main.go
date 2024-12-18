package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findPerimiter(filed [][]byte, pos [2]int, rows, cols int, visited map[[2]int]int) (int, int, int) {
	visited[pos] = 1
	region := filed[pos[0]][pos[1]]
	perimiters := 0
	same := 0
	corners := 0
	area := 1
	for _, v := range [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	} {
		newPos := [2]int{pos[0] + v[0], pos[1] + v[1]}
		if newPos[0] >= 0 && newPos[0] < rows && newPos[1] >= 0 && newPos[1] < cols {
			neighbour := filed[newPos[0]][newPos[1]]
			if neighbour != region {
				perimiters++
			} else {
				same += 1
				if _, ok := visited[newPos]; !ok {
					fArea, fPerimiter, fCorners := findPerimiter(filed, newPos, rows, cols, visited)
					corners += fCorners
					perimiters += fPerimiter
					area += fArea
				}
			}
		} else {
			perimiters++
		}
	}

	return area, perimiters, corners
}

func main() {
	file_path := os.Args[1]
	f, err := os.Open(file_path)
	check(err)
	field := make([][]byte, 0)
	scanner := bufio.NewScanner(f)
	var rows, cols int
	for scanner.Scan() {
		line := []byte(scanner.Text())
		cols = len(line)
		field = append(field, line)
	}
	rows = len(field)

	visited := make(map[[2]int]int)
	res1 := 0
	res2 := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pos := [2]int{i, j}
			if _, ok := visited[pos]; !ok {
				area, primiter, corners := findPerimiter(field, pos, rows, cols, visited)
				res1 += area * primiter
				res2 += area * corners
				// fmt.Printf("%c, %d, %d\n", field[pos[0]][pos[1]], area, primiter)
			}
		}
	}
	fmt.Printf("Initial cost of the fence is %d\n", res1)
	fmt.Printf("Reduced cost of the fence is %d\n", res2)

}
