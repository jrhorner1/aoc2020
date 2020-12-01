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

func main() {
	file, err := os.Open("../input")
	check(err)
	defer file.Close()

	var input []int64 // nil slice
	scanner := bufio.NewScanner(file) // open the file
	for scanner.Scan() { // scan each line
		num, err := strconv.ParseInt(scanner.Text(), 10, 64) // parse the text to int type
		check(err) // check for any parsing errors
		input = append(input, num) // append the value to the slice
	}

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
	fmt.Println(product)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}