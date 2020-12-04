package main

import (
	"fmt"
	"../utils"
	"strings"
	"strconv"
	"regexp"
)

func parse(input string) [][]string {
	var passport [][]string
	fields := strings.Fields(input)
	for _, field := range fields {
		kv := strings.Split(field, ":")
		passport = append(passport, kv)
	}
	return passport
}

func validate(passports []string) (int, int) {
	v1 := 0
	v2 := 0
	for _, passport := range passports {
		ppp := parse(passport)
		p1 := 0
		p2 := true
		for _, kv := range ppp {
			switch kv[0] {
			case "byr": // birth year - four digits; at least 1920 and at most 2002
				p1++
				year, _ := strconv.Atoi(kv[1])
				if year < 1920 || year > 2002 {
					p2 = false
				}
			case "iyr": // issue year - four digits; at least 2010 and at most 2020
				p1++
				year, _ := strconv.Atoi(kv[1])
				if year < 2010 || year > 2020 {
					p2 = false
				}
			case "eyr": // expiration year - four digits; at least 2020 and at most 2030
				p1++
				year, _ := strconv.Atoi(kv[1])
				if year < 2020 || year > 2030 {
					p2 = false
				}
			case "hgt": // height - a number followed by either cm or in:
						// If cm, the number must be at least 150 and at most 193.
						// If in, the number must be at least 59 and at most 76.
				p1++
				hgt, _ := strconv.Atoi(kv[1][:len(kv[1])-2])
				unit := kv[1][len(kv[1])-2:]
				switch unit {
				case "cm": 
					if hgt < 150 || hgt > 193 {
						p2 = false
					}
				case "in":
					if hgt < 59 || hgt > 76 {
						p2 = false
					}
				default:
					p2 = false
				}
			case "hcl": // hair color - a # followed by exactly six characters 0-9 or a-f.
				p1++
				match, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, kv[1])
				if !match {
					p2 = false
				}
			case "ecl": // eye color - exactly one of: amb blu brn gry grn hzl oth.
				p1++
				match, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, kv[1])
				if !match {
					p2 = false
				}
			case "pid": // passport ID - a nine-digit number, including leading zeroes.
				p1++
				match, _ := regexp.MatchString(`^[0-9]{9}$`, kv[1])
				if !match {
					p2 = false
				}
			case "cid": // country ID - ignored, missing or not
				continue
			default: // undefined fields
				continue
			}
		}
		if p1 == 7 {
			v1++
		} 
		if p2 {
			v2++
		}
	}
	return v1, v2
}

func main() {
	input := utils.ReadFile("input")
	passports := strings.Split(strings.Replace(input, " ", "\n", -1), "\n\n")
	valid, verified := validate(passports)
	fmt.Println("Part 1:", valid)
	fmt.Println("Part 2:", verified)
}