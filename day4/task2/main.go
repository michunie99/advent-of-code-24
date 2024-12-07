package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckWindow(arr [][]byte, i, j int) int {
	window := make([]byte, 9)
	for ii := -1; ii <= 1; ii++ {
		for jj := -1; jj <= 1; jj++ {
			window = append(window, arr[i+ii][j+jj])
		}
	}

	var regex = regexp.MustCompile(`(M.S.A.M.S)|(M.M.A.S.S)|(S.S.A.M.M)|(S.M.A.S.M)`)
	// fmt.Println(string(window))
	if regex.MatchString(string(window)) {
		return 1
	}
	return 0
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
	res := 0

	for i := 1; i < len(arr)-1; i++ {
		for j := 1; j < len(arr[i])-1; j++ {
			res += CheckWindow(arr, i, j)
		}
	}

	fmt.Printf("Found %d occurences of XMAS in the input\n", res)
}
