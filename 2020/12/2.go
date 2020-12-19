package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)


func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}


func rotateWaypoint(directionN, directionE, rotations int) (int, int) {
	if rotations < 0 {
		rotations += 4
	}
	rotations %= 4
	switch (rotations) {
		case 0:
			return directionN, directionE
		case 1:
			return -directionE, directionN
		case 2:
			return -directionN, -directionE
		case 3:
			return directionE, -directionN
	}
	return 0, 0
}


func main() {
	x := 0
	y := 0
	waypointN := 1
	waypointE := 10

	scanner := bufio.NewScanner(os.Stdin)
	var line string
	var command rune
	var amount int
	var rotations int
	for scanner.Scan() {
		line = scanner.Text()
		command = rune(line[0])
		amount, _ = strconv.Atoi(line[1:len(line)])
		switch (command) {
			case 'N':
				waypointN += amount
			case 'E':
				waypointE += amount
			case 'W':
				waypointE -= amount
			case 'S':
				waypointN -= amount
			case 'R':
				rotations = amount / 90
				waypointN, waypointE = rotateWaypoint(waypointN, waypointE, rotations)
			case 'L':
				rotations = -amount / 90
				waypointN, waypointE = rotateWaypoint(waypointN, waypointE, rotations)
			case 'F':
				x += waypointE * amount
				y += waypointN * amount
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(abs(x) + abs(y))

}

