package utils

import (
	"os"
	"log"
	"bufio"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ReadFile(filename string) string {
	buffer, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimSpace(string(buffer))
}

func OpenFile() []string {
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