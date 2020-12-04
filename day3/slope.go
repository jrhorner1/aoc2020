package main

import (
	"fmt"
	"../utils"
)

func checkSlope(mov []int, input []string) int {
	trees := 0
	pos := []int{0, 0} // x, y
	for y, i := range input {
		if y != pos[1] {
			continue
		}
		if string(i[pos[0]]) == "#" {
			trees++
		}
		if pos[0] + mov[0] >= len(i) {
			pos[0] += mov[0] - len(i)
		} else {
			pos[0] += mov[0]
		}
		pos[1] += mov[1]
	}
	return trees
}

func p1(input []string) {
	slope := []int{3, 1}
	trees := checkSlope(slope, input)
	fmt.Println("Part 1:", trees)
}

func p2(input []string) {
	slopes := [][]int{
		[]int{1, 1},
		[]int{3, 1},
		[]int{5, 1},
		[]int{7, 1},
		[]int{1, 2}}
	trees := []int{0, 0, 0, 0, 0}
	product := 0
	for i, slope := range slopes {
		trees[i] = checkSlope(slope, input)
		if i == 0 {
			product = trees[i]
		} else {
			product = product * trees[i]
		}
	}
	fmt.Println("Part 2:", product)
}

func main() {
	input := utils.OpenFile()
	p1(input)
	p2(input)
}