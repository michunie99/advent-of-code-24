package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	obsticle = '#'
	empty    = '.'
	visited  = 'X'
)

type Map [][]byte
type Guard struct {
	position [2]int
	heading  [2]int
}

func (g *Guard) move(m Map) bool {
	//fmt.Printf("%v", g)
	m[g.position[0]][g.position[1]] = visited
	rows, cols := len(m), len(m[0])
	new_position := [2]int{
		g.position[0] + g.heading[0],
		g.position[1] + g.heading[1],
	}
	switch {
	case new_position[0] < 0 || new_position[1] < 0 || new_position[0] >= rows || new_position[1] >= cols:
		m[g.position[0]][g.position[1]] = visited
		return true
	case m[new_position[0]][new_position[1]] == empty || m[new_position[0]][new_position[1]] == visited:
		g.position = new_position
		return false
	case m[new_position[0]][new_position[1]] == obsticle:
		switch g.heading {
		case [2]int{0, 1}:
			g.heading = [2]int{1, 0}
		case [2]int{1, 0}:
			g.heading = [2]int{0, -1}
		case [2]int{0, -1}:
			g.heading = [2]int{-1, 0}
		case [2]int{-1, 0}:
			g.heading = [2]int{0, 1}
		}
		return false
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file_path := os.Args[1]
	f, err := os.Open(file_path)
	check(err)

	m := make(Map, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr := []byte(scanner.Text())
		m = append(m, arr)
	}

	var guard Guard
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			switch m[i][j] {
			case 'v':
				guard = Guard{[2]int{i, j}, [2]int{1, 0}}
				m[i][j] = empty
			case '>':
				guard = Guard{[2]int{i, j}, [2]int{0, 1}}
				m[i][j] = empty
			case '<':
				guard = Guard{[2]int{i, j}, [2]int{0, -1}}
				m[i][j] = empty
			case '^':
				guard = Guard{[2]int{i, j}, [2]int{-1, 0}}
				m[i][j] = empty
			}
		}
	}

	// fmt.Printf("Guard initial position %v\n", guard)

	for {
		left := guard.move(m)
		if left {
			break
		}
	}

	res := 0
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == visited {
				res += 1
			}
		}
	}

	// for _, l := range m {
	// 	fmt.Println(string(l))
	// }
	//fmt.Println()
	fmt.Printf("Guard visited %d fileds\n", res)
}
