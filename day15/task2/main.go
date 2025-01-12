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

func drawMap(walls, boxesLeft, boxesRight ObjectMap, robot Pos, rows, cols int) {
	drawing := make([][]byte, rows)
	for i := range rows {
		drawing[i] = make([]byte, cols*2)
		for j := range drawing[i] {
			drawing[i][j] = '.'
		}
	}
	for k := range walls {
		drawing[k[0]][k[1]] = '#'
	}
	for k := range boxesLeft {
		drawing[k[0]][k[1]] = '['
	}
	for k := range boxesRight {
		drawing[k[0]][k[1]] = ']'
	}
	drawing[robot[0]][robot[1]] = '@'

	for _, line := range drawing {
		fmt.Println(string(line))
	}
}

// func (b *WideBox) moveBox(off [2]int, wall ObjectMap, boxes BoxesMap) bool {
// 	newBox := WideBox{b[0] + off[0], b[1] + off[1], b[2] + off[0], b[3] + off[1]}

// 	switch {
// 	case wall[newPos]:
// 		return false
// 	case boxes[newPos]:
// 		return newPos.moveBox(off, wall, boxes)
// 	default:
// 		boxes[newPos] = true
// 		return true
// 	}
// }

func moveBox(pos Pos, boxesLeft, boxesRight ObjectMap, off Pos, walls ObjectMap) bool {
	var moved bool
	var lBox, rBox Pos
	switch {
	case boxesLeft[pos]:
		lBox = pos
		rBox = Pos{pos[0], pos[1] + 1}
	case boxesRight[pos]:
		rBox = pos
		lBox = Pos{pos[0], pos[1] - 1}
	}
	newlBox := Pos{lBox[0] + off[0], lBox[1] + off[1]}
	newrBox := Pos{rBox[0] + off[0], rBox[1] + off[1]}
	switch {
	case walls[newlBox] || walls[newrBox]:
		return false
	case boxesLeft[newlBox] ||
		boxesLeft[newrBox] ||
		boxesLeft[newlBox] ||
		boxesLeft[newrBox]:
		moved = moveBox(lBox, boxesLeft, boxesRight, off, walls)
	default:
		moved = true
	}
	if moved {
		delete(boxesLeft, lBox)
		delete(boxesRight, rBox)
		boxesLeft[newlBox] = true
		boxesRight[newrBox] = true
	}
}

func deleteBox(boxesLeft, boxesRight ObjectMap, newPos Pos) {

}

func main() {
	data, err := os.ReadFile(os.Args[1])
	check(err)

	var robot Pos
	boxesLeft := make(ObjectMap)
	boxesRight := make(ObjectMap)
	walls := make(ObjectMap)
	commands := make(Instructions, 0, len(data))

	var rows, cols int
	for i, line := range bytes.Fields(data) {
		p_commands := len(commands)
		for j, b := range line {
			pos := [2]int{i, j * 2}
			switch b {
			case '#':
				walls[pos] = true
				walls[[2]int{pos[0], pos[1] + 1}] = true
			case 'O':
				boxesLeft[pos] = true
				boxesRight[[2]int{pos[0], pos[1] + 1}] = true
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
		case boxesLeft[newPos]:
			moved := moveBox(newPos, boxesLeft, boxesRight, off, walls)
			if moved {
				robot = newPos
				deleteBox(boxesLeft, boxesRight, newPos)
			}
		case boxesRight[newPos]:
			moved := moveBox(newPos, boxesLeft, boxesRight, off, walls)
			if moved {
				robot = newPos
				deleteBox(boxesLeft, boxesRight, newPos)
			}
		default:
			robot = newPos
		}
	}
	drawMap(walls, boxesLeft, boxesRight, robot, rows, cols)

	// Calculate GPS
	// gps := 0
	// for k := range boxes {
	// 	gps += 100*k[0] + k[1]
	// }
	// fmt.Printf("GPS score for this map is  %d\n", gps)

}
