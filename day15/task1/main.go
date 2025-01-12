package main

import (
	"bytes"
	"fmt"
	"os"
)

type Pos [2]int
type ObjectMap map[Pos]bool
type Instructions []byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func drawMap(walls, boxes ObjectMap, robot Pos, rows, cols int) {
	drawing := make([][]byte, rows)
	for i := range rows {
		drawing[i] = make([]byte, cols)
		for j := range drawing[i] {
			drawing[i][j] = '.'
		}
	}
	for k := range walls {
		drawing[k[0]][k[1]] = '#'
	}
	for k := range boxes {
		drawing[k[0]][k[1]] = 'O'
	}
	drawing[robot[0]][robot[1]] = '@'

	for _, line := range drawing {
		fmt.Println(string(line))
	}
}

func (p *Pos) moveBox(off [2]int, wall, boxes ObjectMap) bool {
	newPos := Pos{p[0] + off[0], p[1] + off[1]}
	switch {
	case wall[newPos]:
		return false
	case boxes[newPos]:
		return newPos.moveBox(off, wall, boxes)
	default:
		boxes[newPos] = true
		return true
	}
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	check(err)

	var robot Pos
	boxes := make(ObjectMap)
	walls := make(ObjectMap)
	commands := make(Instructions, 0, len(data))

	var rows, cols int
	for i, line := range bytes.Fields(data) {
		p_commands := len(commands)
		for j, b := range line {
			pos := [2]int{i, j}
			switch b {
			case '#':
				walls[pos] = true
			case 'O':
				boxes[pos] = true
			case '@':
				robot = pos
			case '\n':
			case '.':
			default:
				commands = append(commands, b)
			}
		}
		if len(commands) == p_commands {
			cols = max(len(line), cols)
			rows = i
		}
	}
	rows++

	for _, c := range commands {
		// fmt.Println(string(c))
		// drawMap(walls, boxes, robot, 10)
		var off [2]int
		switch c {
		case '>':
			off = [2]int{0, 1}
		case 'v':
			off = [2]int{1, 0}
		case '^':
			off = [2]int{-1, 0}
		case '<':
			off = [2]int{0, -1}
		}
		newPos := Pos{robot[0] + off[0], robot[1] + off[1]}
		switch {
		case walls[newPos]:
			break
		case boxes[newPos]:
			moved := newPos.moveBox(off, walls, boxes)
			if moved {
				robot = newPos
				delete(boxes, newPos)
			}
		default:
			robot = newPos
		}
		// if i == 10 {
		// 	break
		// }
	}
	//fmt.Println(boxes, walls, robot, string(commands))
	drawMap(walls, boxes, robot, rows, cols)

	// Calculate GPS
	gps := 0
	for k := range boxes {
		gps += 100*k[0] + k[1]
	}
	fmt.Printf("GPS score for this map is  %d\n", gps)

}
