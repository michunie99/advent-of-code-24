package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile(os.Args[1])
	check(err)

}
