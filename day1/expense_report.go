package main

import (
	"fmt"
	"os"
	"log"
	"strconv"
	"bufio"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func p1(input []int64) {
	values := [2]int64{0, 0}
	for _, i := range input {
		for _, j := range input {
			if i + j == 2020 {
				values[0] = i 
				values[1] = j
				break
			} else {
				continue
			}
		}
		if values[0] != 0 {
			break
		}
	}
	product := values[0] * values[1]
	fmt.Println("Part 1", product)
}

func p2(input []int64) {
	values := [3]int64{0, 0, 0}
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				if i + j + k == 2020 {
					values[0] = i 
					values[1] = j
					values[2] = k
					break
				} else {
					continue
				}
			}
			if values[1] != 0 {
				break
			}
		}
		if values[0] != 0 {
			break
		}
	}
	product := values[0] * values[1] * values[2]
	fmt.Println("Part 2", product)
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	var input []int64 // nil slice
	scanner := bufio.NewScanner(file) // open the file
	for scanner.Scan() { // scan each line
		num, err := strconv.ParseInt(scanner.Text(), 10, 64) // parse the text to int type
		check(err) // check for any parsing errors
		input = append(input, num) // append the value to the slice
	}

	p1(input)
	p2(input)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}