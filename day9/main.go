package main

import (
	"fmt"
	"os"
	"strconv"
)

type DisckFile struct {
	fid   int
	start int
	size  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func calculateCheckSum(disk []int) int {
	res := 0
	for i, v := range disk {
		if v != -1 {
			res += i * v
		}
	}
	return res
}

func defragment(disk []int) []int {
	defrag := make([]int, len(disk))
	copy(defrag, disk)
	for i, j := 0, len(disk)-1; i <= j; {
		if defrag[j] == -1 {
			j--
			continue
		}
		if defrag[i] != -1 {
			i++
			continue
		}
		tmp := defrag[i]
		defrag[i] = defrag[j]
		defrag[j] = tmp
	}
	return defrag
}

func defragmentContinous(files []DisckFile, freeSpaces []DisckFile) []int {
	total_size := files[len(files)-1].start + files[len(files)-1].size
	for i := len(files) - 1; i >= 0; i-- {
		for j := 0; j < len(freeSpaces); j++ {
			if freeSpaces[j].size >= files[i].size && freeSpaces[j].start < files[i].start {
				files[i].start = freeSpaces[j].start
				freeSpaces[j].size = freeSpaces[j].size - files[i].size
				freeSpaces[j].start = freeSpaces[j].start + files[i].size
				break
			}
		}
	}
	disk := make([]int, total_size)
	for i := 0; i < len(disk); i++ {
		disk[i] = -1
	}
	for _, f := range files {
		for i := range f.size {
			disk[f.start+i] = f.fid
		}
	}
	for _, fs := range freeSpaces {
		for i := range fs.size {
			disk[fs.start+i] = fs.fid
		}
	}
	return disk
}

func main() {
	file_path := os.Args[1]
	data, err := os.ReadFile(file_path)
	check(err)
	disk := make([]int, 0, len(data))
	free_spaces := make([]DisckFile, 0)
	files := make([]DisckFile, 0)
	id := 0
	for i := range len(data) {
		size, err := strconv.Atoi(string(data[i]))
		if err != nil {
			continue
		}
		var curr_id int
		switch i % 2 {
		case 0: // datablock
			curr_id = id
			files = append(files, DisckFile{id, len(disk), size})
			id++
		case 1: // freespace
			curr_id = -1
			free_spaces = append(free_spaces, DisckFile{-1, len(disk), size})
		}
		for range size {
			disk = append(disk, curr_id)
		}
	}

	defrag1 := defragment(disk)
	checkSum1 := calculateCheckSum(defrag1)
	fmt.Printf("Checksum of the defragmented drive is %d\n", checkSum1)

	defrag2 := defragmentContinous(files, free_spaces)
	checkSum2 := calculateCheckSum(defrag2)
	fmt.Printf("Checksum of the defragmented drive second aproach is %d\n", checkSum2)
}
