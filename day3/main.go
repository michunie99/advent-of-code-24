package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f_name := os.Args[1]
	data, err := os.ReadFile(f_name) // Hope not to big
	check(err)
	res1, res2 := 0, 0
	execute := true
	// var r = regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	var r = regexp.MustCompile(`mul\((\d*),(\d*)\)|don't\(\)|do\(\)`)
	for _, expr := range r.FindAllStringSubmatch(string(data), -1) {
		switch expr[0] {
		case "do()":
			execute = true
		case "don't()":
			execute = false
		default:
			num1, err := strconv.Atoi(expr[1])
			check(err)
			num2, err := strconv.Atoi(expr[2])
			check(err)
			res1 += num1 * num2
			if execute {
				res2 += num1 * num2
			}

		}
	}
	fmt.Printf("Results is: %d\n", res1)
	fmt.Printf("Results with adjustments is: %d\n", res2)
}
