package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"bufio"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func parse() []string {
	file, err := os.Open("input")
	check(err)
	defer file.Close()
	var input []string 
	scanner := bufio.NewScanner(file) 
	for scanner.Scan() { 
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		text := scanner.Text() 
		check(err) 
		input = append(input, text) 
	}
	return input
}

func getFields(in string) []string {
	f := strings.Fields(in)
	ss := strings.FieldsFunc(f[0], func(r rune) bool {
		if r =='-' {
			return true
		}
		return false
	})
	fields := []string{ss[0], ss[1], string(f[1][0]), f[2]} 
	return fields
}

func p1(input []string) {
	valid := 0
	for _, i := range input {
		fields := getFields(i)
		lower, _ := strconv.Atoi(fields[0])
		upper, _ := strconv.Atoi(fields[1])
		count := 0
		for _, j := range fields[3] {
			if string(j) == fields[2] {
				count++
			}
		}
		if count >= lower && count <= upper {
			valid++
		}
	}
	fmt.Println("Part 1:", valid)
}

func p2(input []string) {
	valid := 0
	for _, i := range input {
		fields := getFields(i)
		p1, _ := strconv.Atoi(fields[0])
		p2, _ := strconv.Atoi(fields[1])
		if string(fields[3][p1-1]) == fields[2] || string(fields[3][p2-1]) == fields[2] {
			if !(string(fields[3][p1-1]) == fields[2] && string(fields[3][p2-1]) == fields[2]) {
				valid++
			}
		}
	}
	fmt.Println("Part 2:", valid)
}

func main() {
	input := parse()
	p1(input)
	p2(input)
}