package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"image"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	in := strings.Split(strings.TrimSpace(string(input)), "\n")
	seats := map[image.Point]int{}
	for x := range in {
		for y := range in[x] {
			switch string(in[x][y]) {
			case "#" : seats[ image.Point{ x , y } ] = 2
			case "L" : seats[ image.Point{ x , y } ] = 1
			case "." : seats[ image.Point{ x , y } ] = 0
			}
		}
	}
	adjacentSeats := []image.Point{ 
		{-1, -1}, // top left
		{-1,  0}, // top
		{-1,  1}, // top right
		{ 0, -1}, // left
		{ 0,  1}, // right
		{ 1, -1}, // bottom left
		{ 1,  0}, // bottom
		{ 1,  1}, // bottom right
	}
	occupiedSeats := 0
	for loop := true; loop; {
		occupiedSeats = 0 // reset occupiedSeats for each iteration
		occupied, vacant := 2, 1
		new_seats := map[image.Point]int{}
		for seat, status := range seats {
			adjacentOccupied := 0
			for _, adjacentSeat := range adjacentSeats {
				if seats[seat.Add(adjacentSeat)] == occupied { adjacentOccupied++ }
			}
			if status == occupied && adjacentOccupied >= 4 {
				status = vacant
			} else if status == vacant && adjacentOccupied == 0 || status == occupied {
				status = occupied
				occupiedSeats++
			}
			new_seats[seat] = status
			loop = false || seats[seat] != new_seats[seat] 
		}
		seats = new_seats
	}
	fmt.Println(seats)
	fmt.Println("Part 1:", occupiedSeats)
	// fmt.Println("Part 2:", input)
}