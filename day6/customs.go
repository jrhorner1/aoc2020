package main

import (
	"fmt"
	"../utils"
	"strings"
	"regexp"
)

func p1(answers []string) {
	count := 0
	q := [26]string{ "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, i := range answers {
		for _, j := range q {
			match, _ := regexp.MatchString(j, i)
			if match { count++ }
		}
	}
	fmt.Println("Part 1:",count)
}

func p2(answers []string) {
	count := 0
	for _, i := range answers {
		split := strings.Split(i, "\n")
		q := map[string]bool{ "a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true, "h": true, "i": true, "j": true, "k": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "r": true, "s": true, "t": true, "u": true, "v": true, "w": true, "x": true, "y": true, "z": true}
		for _, j := range split {
			for k, _ := range q {
				match, _ := regexp.MatchString(k, j)
				if !match { q[k] = match }
			}
		}
		for _, v := range q {
			if v { count++}
		}
	}
	fmt.Println("Part 2:",count)
}

func main() {
	input := utils.ReadFile("input")
	answers := strings.Split(input, "\n\n")
	p1(answers)
	p2(answers)
}