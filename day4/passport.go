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
		p2 := []bool{false, false, false, false, false, false, false}
		for _, kv := range ppp { 
			switch kv[0] {
			case "byr": // birth year - four digits; at least 1920 and at most 2002
				p1++
				year, _ := strconv.Atoi(kv[1])
				if !( year < 1920 || year > 2002 ) {
					p2[0] = true
				}
			case "iyr": // issue year - four digits; at least 2010 and at most 2020
				p1++
				year, _ := strconv.Atoi(kv[1])
				if !( year < 2010 || year > 2020 ) {
					p2[1] = true
				}
			case "eyr": // expiration year - four digits; at least 2020 and at most 2030
				p1++
				year, _ := strconv.Atoi(kv[1])
				if !( year < 2020 || year > 2030 ) {
					p2[2] = true
				}
			case "hgt": // height - a number followed by either cm or in:
						// If cm, the number must be at least 150 and at most 193.
						// If in, the number must be at least 59 and at most 76.
				p1++
				value := kv[1][:len(kv[1])-2] 
				height, _ := strconv.Atoi(value)
				unit := kv[1][len(kv[1])-2:]
				switch unit {
				case "cm": 
					if !( height < 150 || height > 193 ) {
						p2[3] = true
					}
				case "in":
					if !( height < 59 || height > 76 ) {
						p2[3] = true
					}
				}
			case "hcl": // hair color - a # followed by exactly six characters 0-9 or a-f.
				p1++
				match, _ := regexp.MatchString("^#[0-9a-f]{6}$", kv[1])
				if match {
					p2[4] = true
				}
			case "ecl": // eye color - exactly one of: amb blu brn gry grn hzl oth.
				p1++
				match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", kv[1])
				if match {
					p2[5] = true
				}
			case "pid": // passport ID - a nine-digit number, including leading zeroes.
				p1++
				match, _ := regexp.MatchString("^[0-9]{9}$", kv[1])
				if match {
					p2[6] = true
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
		check := 0
		for _, i := range p2 {
			if i {
				check++
			}
		}
		if check == 7 {
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