package main

import (
	"bufio"
	"fmt"
	"os"
)

type Cord struct {
	x int
	y int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FindSequence(board [][]byte, i, j int, sequence []byte, direction []int, visited map[Cord]struct{}) int {
	if len(sequence) == 0 {
		// visited[Cord{i, j}] = struct{}{}
		return 1
	}
	res := 0
	if row, col := i+direction[0], j+direction[1]; row >= 0 && col >= 0 &&
		row < len(board) && col < len(board[row]) &&
		board[row][col] == sequence[0] {
		if vis := FindSequence(board, row, col, sequence[1:], direction, visited); vis != 0 {
			// visited[Cord{row, col}] = struct{}{}
			res += vis
		}
	}
	return res
}

func main() {
	f_name := os.Args[1]
	f, err := os.Open(f_name)
	check(err)

	arr := make([][]byte, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr = append(arr, []byte(scanner.Text()))
	}

	// visited := make(map[Cordintes]struct{})
	sequence := []byte("MAS")
	res := 0

	visited := make(map[Cord]struct{})
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			//fmt.Printf("%c", arr[i][j])
			if arr[i][j] == 'X' {
				// p_res := res
				res += FindSequence(arr, i, j, sequence, []int{0, 1}, visited)
				res += FindSequence(arr, i, j, sequence, []int{0, -1}, visited)
				res += FindSequence(arr, i, j, sequence, []int{1, 0}, visited)
				res += FindSequence(arr, i, j, sequence, []int{-1, 0}, visited)
				res += FindSequence(arr, i, j, sequence, []int{1, 1}, visited)
				res += FindSequence(arr, i, j, sequence, []int{-1, 1}, visited)
				res += FindSequence(arr, i, j, sequence, []int{1, -1}, visited)
				res += FindSequence(arr, i, j, sequence, []int{-1, -1}, visited)
				// if p_res < res {
				// 	visited[Cord{i, j}] = struct{}{}
				// }
			}
		}
	}

	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr[i]); j++ {
	// 		if _, ok := visited[Cord{i, j}]; !ok {
	// 			fmt.Printf(".")
	// 		} else {
	// 			fmt.Printf("%c", arr[i][j])
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Printf("Found %d substings of XMAS in the input\n", res)
}
