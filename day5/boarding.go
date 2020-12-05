package main

import (
	"fmt"
	"../utils"
	"sort"
)

type BoardingPass struct {
	Value string
	Row int
	Column int
	Id int
}

func getrange(value rune, v1 int, v2 int) (int, int) {
	switch string(value) {
	case "F", "L":
		v2 = (v2 + v1) / 2
	case "B", "R":
		v1 = (v1 + v2 + 1) / 2
	}
	return v1, v2
}

func findseat(bp BoardingPass) BoardingPass {
	r1, r2 := 0, 127
	c1, c2 := 0, 7
	for i, value := range bp.Value {
		switch i {
		case 0, 1, 2, 3, 4, 5:
			r1, r2 = getrange(value, r1, r2)
		case 6:
			if string(value) == "F" {
				bp.Row = r1
			} else if string(value) == "B" {
				bp.Row = r2
			}
		case 7, 8:
			c1, c2 = getrange(value, c1, c2)
		case 9:
			if string(value) == "L" {
				bp.Column = c1
			} else if string(value) == "R" {
				bp.Column = c2
			}
		}
	}
	bp.Id = (bp.Row * 8) + bp.Column
	return bp
}

func p1(bps []BoardingPass) {
	maxId := 0
	for _, bp := range bps {
		// fmt.Println(bp)
		if bp.Id > maxId {
			maxId = bp.Id
		}
	}
	fmt.Println("Part 1:", maxId)
}

type ByRowCol []BoardingPass
func (a ByRowCol) Len() int { return len(a) }
func (a ByRowCol) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRowCol) Less(i, j int) bool { 
	if a[i].Row < a[j].Row {
		return true
	}
	if a[i].Row > a[j].Row {
		return false
	}
	return a[i].Column < a[j].Column 
}

func p2(bps []BoardingPass) {
	var my_bp BoardingPass
	sort.Sort(ByRowCol(bps))
	for i, bp := range bps {
		if i+2 > len(bps) {
			break
		}
		if bps[i+1].Id != bp.Id + 1 && bps[i+1].Id == bp.Id + 2 {
			my_bp.Id = bp.Id + 1
			my_bp.Row = bp.Row
			my_bp.Column = bp.Column + 1
		}
	}
	fmt.Println("Part 2:", my_bp.Id, my_bp)
}

func main() {
	input := utils.OpenFile()
	var bps []BoardingPass
	for _, i := range input {
		pass := BoardingPass{i, 0, 0, 0}
		bps = append(bps, pass)
	}
	for i, bp := range bps {
		bps[i] = findseat(bp)
	}
	p1(bps)
	p2(bps)
}