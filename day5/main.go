package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type IntSet map[int]int
type Rules map[int]IntSet

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (rules Rules) validatePage(line string) (int, []int) {
	parts := strings.Split(line, ",")
	order := make([]int, 0, len(line))
	for _, p := range parts {
		num, err := strconv.Atoi(p)
		check(err)
		order = append(order, num)
	}
	contained := make(IntSet)
	for i := len(order) - 1; i >= 0; i-- {
		for r := range rules[order[i]] {
			for c := range contained {
				if c == r {
					fmt.Println(order)
					return 0, order
				}
			}
		}
		contained[order[i]] = i
	}
	// fmt.Println(order)
	return order[len(order)/2], order
}

func (rules Rules) addRule(line string) {
	ints := strings.Split(line, "|")
	key, err := strconv.Atoi(ints[1])
	check(err)
	val, err := strconv.Atoi(ints[0])
	check(err)
	_, ok := rules[key]
	if !ok {
		rules[key] = make(IntSet)
	}
	rules[key][val] = 0
}

func (rules Rules) fixOrder(order []int) int {
	fmt.Println(order)
	contained := make(IntSet)
	for i := len(order) - 1; i >= 0; i-- {
		for r := range rules[order[i]] {
			for c, j := range contained {
				if c == r {
					tmp := order[i]
					order[i] = order[j]
					order[j] = tmp
					return rules.fixOrder(order) // NOTE: Very ugly but works
				}
			}
		}
		contained[order[i]] = i
	}
	return order[len(order)/2]
}

func main() {
	file_path := os.Args[1]
	f, err := os.Open(file_path)
	check(err)

	rules := make(Rules)
	var validRule = regexp.MustCompile(`\d+\|\d+`) // NOTE: would be nice to use groups to capture the digits :/
	var validOrder = regexp.MustCompile(`(,?\d)+`)

	res := 0
	scanner := bufio.NewScanner(f)
	toCorrect := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case validRule.MatchString(line):
			rules.addRule(line)
		case validOrder.MatchString(line):
			middle, order := rules.validatePage(line)
			res += middle
			if middle == 0 {
				toCorrect = append(toCorrect, order)
			}
		}

	}

	corr_res := 0
	for _, order := range toCorrect {
		// fmt.Println(order)
		corr_res += rules.fixOrder(order)
		// fmt.Println(order)
		// fmt.Println()
	}
	// fmt.Println()
	// fmt.Printf("%v\n", rules)
	// fmt.Println()
	fmt.Printf("Sum of middle pages is %d\n", res)
	fmt.Printf("Sum of corrected middle pages is %d\n", corr_res)
}
