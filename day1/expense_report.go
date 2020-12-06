package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	entries := []int{}
	for _, i := range strings.Fields(string(input)) {
		entry, _ := strconv.Atoi(i)
		entries = append(entries, entry)
	}
	for i, x := range entries {
		for j, y := range entries[i+1:] {
			if x + y == 2020 { fmt.Println("Part 1:", x * y) }
			for _, z := range entries[j+1:] {
				if x + y + z == 2020 { fmt.Println("Part 2:", x * y * z) }
			}
		}  
	}
}