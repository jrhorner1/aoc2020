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

func validate(passports []string) int  {
	valid := 0
	for _, passport := range passports {
		ppp := parse(passport)
		fmt.Println(ppp)
		rfc := true
		for _, kv := range ppp {
			switch kv[0] {
			case "byr": // birth year - four digits; at least 1920 and at most 2002
				status := true
				year, _ := strconv.Atoi(kv[1])
				if year < 1920 || year > 2002 {
					rfc = false
					status = false
				}
				fmt.Println("byr", status)
			case "iyr": // issue year - four digits; at least 2010 and at most 2020
				status := true
				year, _ := strconv.Atoi(kv[1])
				if year < 2010 || year > 2020 {
					rfc = false
					status = false
				}
				fmt.Println("iyr", status)
			case "eyr": // expiration year - four digits; at least 2020 and at most 2030
				status := true
				year, _ := strconv.Atoi(kv[1])
				if year < 2020 || year > 2030 {
					rfc = false
					status = false
				}
				fmt.Println("eyr", status)
			case "hgt": // height - a number followed by either cm or in:
						// If cm, the number must be at least 150 and at most 193.
						// If in, the number must be at least 59 and at most 76.
				status := true
				hgt, _ := strconv.Atoi(kv[1][:len(kv[1])-2])
				unit := kv[1][len(kv[1])-2:]
				switch unit {
				case "cm": 
					if hgt < 150 || hgt > 193 {
						rfc = false
						status = false
					}
				case "in":
					if hgt < 59 || hgt > 76 {
						rfc = false
						status = false
					}
				default:
					rfc = false
					status = false
				}
				fmt.Println("hgt", status)
			case "hcl": // hair color - a # followed by exactly six characters 0-9 or a-f.
			status := true
				match, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, kv[1])
				if !match {
					rfc = false
					status = false
				}
				fmt.Println("hcl", status)
			case "ecl": // eye color - exactly one of: amb blu brn gry grn hzl oth.
			status := true
				match, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, kv[1])
				if !match {
					rfc = false
					status = false
				}
				fmt.Println("ecl", status)
			case "pid": // passport ID - a nine-digit number, including leading zeroes.
			status := true
				match, _ := regexp.MatchString(`^[0-9]{9}$`, kv[1])
				if !match {
					rfc = false
					status = false
				}
				fmt.Println("pid", status)
			case "cid": // country ID - ignored, missing or not
				continue
			default: // undefined fields
				continue
			}
		}
		if rfc {
			valid++
		}
	}
	return valid
}

func p1(passports []string) {
	valid := 0
	for _, passport := range passports {
		ppp := parse(passport)
		rfc := 0
		for _, kv := range ppp {
			switch kv[0] {
			case "byr": // birth year
				rfc++ 
			case "iyr": // issue year
				rfc++ 
			case "eyr": // expiration year
				rfc++ 
			case "hgt": // height
				rfc++ 
			case "hcl": // hair color
				rfc++ 
			case "ecl": // eye color
				rfc++ 
			case "pid": // passport ID
				rfc++ 
			case "cid":
				continue
			default:
				continue
			}
		}
		if rfc == 7 {
			valid++
		} 
	}
	fmt.Println("Part 1:", valid)
}

func p2(passports []string) {
	valid := validate(passports)
	fmt.Println("Part 2:", valid)
}

func main() {
	input := utils.ReadFile("input")
	passports := strings.Split(strings.Replace(input, " ", "\n", -1), "\n\n")
	p1(passports)
	p2(passports)
}