package main

import (
	"bufio"
	"fmt"
	//	"io"
	"container/heap"
	"math"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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

	h1, h2 := &IntHeap{}, &IntHeap{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) >= 2 {
			num1, err1 := strconv.Atoi(parts[0])
			num2, err2 := strconv.Atoi(parts[1])
			// Check for errors
			check(err1)
			check(err2)
			heap.Push(h1, num1)
			heap.Push(h2, num2)
		}
	}
	check(scanner.Err())

	dist := 0.0
	for h1.Len() != 0 && h2.Len() != 0 {
		n1 := heap.Pop(h1).(int)
		n2 := heap.Pop(h2).(int)
		dist += math.Abs(float64(n1 - n2))
	}
	fmt.Printf("Distance is: %d\n", int(dist))
}
