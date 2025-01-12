package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Example usage
// go run main.go data/input.txt 100 103 101

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func convertNumber(s string) int {
	num, err := strconv.Atoi(s)
	check(err)
	return num
}

func wrap(pos [4]int, rows, cols int) [4]int {
	//	fmt.Println("Before :", pos)
	pos[0] = pos[0] % cols
	if pos[0] < 0 {
		pos[0] += cols
	}
	pos[1] = pos[1] % rows
	if pos[1] < 0 {
		pos[1] += rows
	}
	//	fmt.Println("After: ", pos)
	return pos
}

func printRobots(robots map[int][4]int, rows, cols int) [][]byte {
	grid := make([][]byte, 0)
	for range rows {
		row := make([]byte, cols)
		for i := range len(row) {
			row[i] = '.'
		}
		grid = append(grid, row)
	}
	for _, robot := range robots {
		grid[robot[1]][robot[0]] = 'X'
	}
	//for _, row := range grid {
	//	fmt.Println(string(row))
	//}
	//fmt.Println()
	return grid
}

func isTree(board [][]byte, rows, cols int, thr int) bool {
	visited := make(map[[2]int]int)
	biggestRegion := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				biggestRegion = max(biggestRegion, dfs(board, i, j, rows, cols, visited))
			}
		}
	}
	//fmt.Println(biggestRegion)
	return biggestRegion > thr
}

func dfs(board [][]byte, x, y, rows, cols int, visited map[[2]int]int) int {
	if _, ok := visited[[2]int{x, y}]; ok {
		return 0
	}
	size := 1
	visited[[2]int{x, y}] = 1
	for _, diff := range [][2]int{
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
	} {
		newX, newY := x+diff[0], y+diff[1]
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && board[newX][newY] == 'X' {
			size += dfs(board, newX, newY, rows, cols, visited)
		}
	}
	return size
}

func simulateRobots(robotsInit map[int][4]int, seconds, rows, cols int) map[int][4]int {
	robots := make(map[int][4]int)
	for k, v := range robotsInit {
		robots[k] = v
	}
	for i, robot := range robots {
		robots[i] = wrap(
			[4]int{
				robot[0] + seconds*robot[2],
				robot[1] + seconds*robot[3],
				robot[2],
				robot[3],
			},
			rows, cols,
		)
	}
	return robots
}

func main() {

	f, err := os.Open(os.Args[1])
	check(err)
	seconds, err := strconv.Atoi(os.Args[2])
	check(err)
	rows, err := strconv.Atoi(os.Args[3])
	check(err)
	cols, err := strconv.Atoi(os.Args[4])
	check(err)

	scanner := bufio.NewScanner(f)
	robotStats := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)
	initRobots := make(map[int][4]int)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		parts := robotStats.FindAllStringSubmatch(line, -1)
		if parts == nil {
			continue
		}
		initRobots[i] = [4]int{
			convertNumber(parts[0][1]),
			convertNumber(parts[0][2]),
			convertNumber(parts[0][3]),
			convertNumber(parts[0][4]),
		}
	}
	//fmt.Println(robots)
	robots := simulateRobots(initRobots, seconds, rows, cols)

	// Calculate score
	quadrants := make(map[int]int)
	for _, robot := range robots {
		var quadrant int
		switch {
		case robot[0] >= 0 && robot[0] < cols/2 && robot[1] >= 0 && robot[1] < rows/2:
			quadrant = 0
		case robot[0] > cols/2 && robot[1] >= 0 && robot[1] < rows/2:
			quadrant = 1
		case robot[0] >= 0 && robot[0] < cols/2 && robot[1] > rows/2:
			quadrant = 2
		case robot[0] > cols/2 && robot[1] > rows/2:
			quadrant = 3
		default:
			continue
		}
		quadrants[quadrant] += 1
	}
	res := 1
	for _, cnt := range quadrants {
		res *= cnt
	}

	// Find tree
	thr := int(len(initRobots) / 6)
	for sec := range 20000 {
		robots := simulateRobots(initRobots, sec, rows, cols)
		board := printRobots(robots, rows, cols)
		if isTree(board, rows, cols, thr) {
			for _, row := range board {
				fmt.Println(string(row))
			}
			fmt.Println(sec)
			fmt.Println()
		}
	}
	fmt.Printf("Safety factor is %d\n", res)
}
