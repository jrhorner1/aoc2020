package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
)

func getParents(bags map[string]int, rules map[string]map[string]int, s string)  {
	for p, cs := range rules {
		for c, _ := range cs {
			if c == s {
				bags[p]++
				getParents(bags, rules, p)
				break
			}
		}
	}
}

func getChildren(rules map[string]map[string]int, s string) (total int) {
	for child, count := range rules[s] {
		total += count * (getChildren(rules, child) + 1)
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("input")
	// dotted indigo bags contain 4 faded black bags, 4 clear cyan bags, 5 vibrant teal bags.
	rules := map[string]map[string]int{}
	for _, r := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		reg1, reg2 := regexp.MustCompile(`\sbags|[.]`), regexp.MustCompile(`\scontain`)
		r = reg1.ReplaceAllString(r, "")
		s := reg2.Split(r, -1)
		parent, children := s[0], s[1]
		rules[parent] = map[string]int{}
		for _, child := range regexp.MustCompile(`(\d+) (\w+ \w+)`).FindAllStringSubmatch(children, -1) {
			rules[parent][child[2]], _ = strconv.Atoi(child[1])
		}
		 
	}
	parents := map[string]int{}
	getParents(parents, rules, "shiny gold")
	fmt.Println("Part 1:", len(parents))
	fmt.Println("Part 2:", getChildren(rules, "shiny gold"))
}