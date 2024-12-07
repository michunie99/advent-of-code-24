package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type ResultsSequence []int

func (rs ResultsSequence) IsSafe() int {
	var sum int
	for i := 1; i < len(rs); i++ {
		diff := rs[i-1] - rs[i]
		if diff > 3 || diff < -3 || diff == 0 || (diff > 0 && sum < 0) || (diff < 0 && sum > 0) {
			return i - 1
		}
		sum += diff
	}
	return -1
}

func (rs ResultsSequence) IsSafeCorrected() bool {
	if i := rs.IsSafe(); i != -1 {
		if slices.Delete(slices.Clone(rs), i, i+1).IsSafe() != -1 &&
			slices.Delete(slices.Clone(rs), i+1, i+2).IsSafe() != -1 && // NOTE: possibly can caouse to delete outside of the array I guess
			slices.Delete(slices.Clone(rs), 0, 1).IsSafe() != -1 {
			return false
		}
	}
	return true
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
	res1, res2 := 0, 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		var nums ResultsSequence
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			check(err)
			nums = append(nums, num)
		}
		if nums.IsSafe() == -1 {
			res1++
		}
		if nums.IsSafeCorrected() {
			res2++
		}
	}
	fmt.Printf("Detected %d safe lines\n", res1)
	fmt.Printf("Detected %d safe lines after correction\n", res2)
}
