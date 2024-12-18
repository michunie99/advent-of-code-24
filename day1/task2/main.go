package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file_path := os.Args[1]
	f, err := os.Open(file_path)
	check(err)

	left := make(map[int]int)
	right := make(map[int]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) >= 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			check(err1)
			check(err2)
			left[num1]++
			right[num2]++
		}
	}
	check(scanner.Err())
	res := 0
	for key, cnt := range left {
		res += (key * right[key]) * cnt
	}
	fmt.Printf("Simularity score is: %d\n", res)
}
