package main

import (
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

func numberDigits(num int) int {
	if num == 0 {
		return 1
	}
	i := 0
	for num != 0 {
		num /= 10
		i++
	}
	return i
}

func main() {
	file_path := os.Args[1]
	data, err := os.ReadFile(file_path)
	check(err)
	blinks, err := strconv.Atoi(os.Args[2])
	check(err)
	stones := make(map[int]int)
	for _, numStr := range strings.Fields(string(data)) {
		num, err := strconv.Atoi(numStr)
		check(err)
		stones[num]++
	}

	for range blinks {
		// fmt.Println(stones)
		newStones := make(map[int]int)
		for k, v := range stones {
			nDigits := numberDigits(k)
			switch nDigits % 2 {
			case 0:
				base := int(math.Pow(float64(10), float64(nDigits/2)))
				left, right := k/base, k%base
				newStones[left] += v
				newStones[right] += v
			case 1:
				if k == 0 {
					newStones[1] += v
				} else {
					newStones[k*2024] += v
				}
			}
		}
		stones = newStones
	}
	res := 0
	for _, v := range stones {
		res += v
	}
	fmt.Printf("%d blinks: %d\n", blinks, res)
}
