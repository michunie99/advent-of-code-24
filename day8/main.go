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

func validatePoint(p [2]int, xMin, xMax, yMin, yMax int) bool {
	return p[0] >= xMin && p[0] < xMax && p[1] >= yMin && p[1] < yMax
}

func findAntinodes(a1 [2]int, a2 [2]int) [][2]int {
	// fmt.Println(a1, a2)
	res := make([][2]int, 2)
	dir := [2]int{a1[0] - a2[0], a1[1] - a2[1]}
	// fmt.Println(dir)
	res[0] = [2]int{a1[0] + dir[0], a1[1] + dir[1]}
	res[1] = [2]int{a2[0] - dir[0], a2[1] - dir[1]}
	// fmt.Println(res)
	return res
}
func bonusAntinodes(a1 [2]int, a2 [2]int, xMin, xMax, yMin, yMax int) [][2]int {
	// fmt.Println(a1, a2)
	res := make([][2]int, 0)
	dir := [2]int{a1[0] - a2[0], a1[1] - a2[1]}

	for i := 0; ; i++ {
		offset := [2]int{dir[0] * i, dir[1] * i}
		antinode := [2]int{a1[0] + offset[0], a1[1] + offset[1]}
		if validatePoint(antinode, xMin, xMax, yMin, yMax) {
			res = append(res, antinode)
		} else {
			break
		}
	}
	for i := 0; ; i++ {
		offset := [2]int{dir[0] * i, dir[1] * i}
		antinode := [2]int{a2[0] - offset[0], a2[1] - offset[1]}
		if validatePoint(antinode, xMin, xMax, yMin, yMax) {
			res = append(res, antinode)
		} else {
			break
		}
	}
	return res
}

func main() {
	file_path := os.Args[1]
	f, err := os.Open(file_path)
	check(err)

	scanner := bufio.NewScanner(f)
	var cols int
	rows := 0

	towers := make(map[rune][][2]int) // NOTE: this type tho
	for scanner.Scan() {
		line := scanner.Text()
		cols = len(line)
		for col, c := range line {
			if c != '.' && c != '#' {
				if _, ok := towers[c]; !ok {
					towers[c] = make([][2]int, 0)
				}
				towers[c] = append(towers[c], [2]int{rows, col})
			}
		}
		rows++
	}

	antinodes := make(map[[2]int]int)
	antinodesAll := make(map[[2]int]int)
	for _, pos := range towers {
		for i, p1 := range pos {
			for _, p2 := range pos[i+1:] {
				pn := findAntinodes(p1, p2)
				if validatePoint(pn[0], 0, rows, 0, cols) {
					antinodes[pn[0]] = 1
				}
				if validatePoint(pn[1], 0, rows, 0, cols) {
					antinodes[pn[1]] = 1
				}
				pnn := bonusAntinodes(p1, p2, 0, rows, 0, cols)
				for _, node := range pnn {
					antinodesAll[node] = 1
				}
			}
		}

	}
	// fmt.Println(towers)
	// fmt.Println(antinodesAll)
	fmt.Printf("There are %d antinodes\n", len(antinodes))
	fmt.Printf("There are %d bonus antinodes\n", len(antinodesAll))
}
