package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func validLine(target int, numbers []int) bool {
	mask := 0
	max_mask := int(math.Pow(float64(2), float64(len(numbers))))
	for {
		sum := 0
		div := 1
		for i := range len(numbers) {
			// div := int(math.Pow(float64(2), float64(i)))
			switch (mask / div) % 2 {
			case 0:
				sum += numbers[i]
			case 1:
				sum *= numbers[i]
			}
			div *= 2
		}
		if sum == target {
			return true
		}
		mask++
		if mask > max_mask {
			break
		}
	}
	return false
}

func numberDigits(num int) int {
	tmp := num
	res := 0
	for tmp > 0 {
		tmp /= 10
		res++
	}
	return res
}

func validLineCorrected(target int, numbers []int) bool {
	mask := 0
	max_mask := int(math.Pow(float64(3), float64(len(numbers))))
	for {
		sum := 0
		div := 1
		for i := range len(numbers) {
			// div := int(math.Pow(float64(3), float64(i)))
			switch (mask / div) % 3 {
			case 0:
				sum += numbers[i]
			case 1:
				sum *= numbers[i]
			case 2:
				digits := numberDigits(numbers[i])
				sum *= int(math.Pow(float64(10), float64(digits)))
				sum += numbers[i]
			}
			div *= 3
		}
		if sum == target {
			return true
		}
		mask++
		if mask > max_mask {
			break
		}
	}
	return false
}

func main() {
	file_path := os.Args[1]
	f, err := os.Open(file_path)
	check(err)

	scanner := bufio.NewScanner(f)
	res1 := 0
	res2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		target, err := strconv.Atoi(parts[0])
		check(err)
		numbers := make([]int, 0)
		for _, p := range strings.Fields(parts[1]) {
			n, err := strconv.Atoi(p)
			check(err)
			numbers = append(numbers, n)
		}
		if validLine(target, numbers) {
			res1 += target
		}
		if validLineCorrected(target, numbers) {
			res2 += target
		}

	}

	fmt.Printf("Sum of corrected line with 2 operators is %d\n", res1)
	fmt.Printf("Sum of corrected line with 3 operators is %d\n", res2)
}
