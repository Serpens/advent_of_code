package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)

const directionOrder string = "ESWN"


func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}


func main() {
	x := 0
	y := 0
	direction := 0

	scanner := bufio.NewScanner(os.Stdin)
	var line string
	var command rune
	var amount int
	for scanner.Scan() {
		line = scanner.Text()
		command = rune(line[0])
		amount, _ = strconv.Atoi(line[1:len(line)])
		if command == 'F' {
			command = rune(directionOrder[direction])
		}
		switch (command) {
			case 'N':
				x += amount
			case 'E':
				y += amount
			case 'W':
				y -= amount
			case 'S':
				x -= amount
			case 'R':
				direction += amount / 90
				direction %= 4
			case 'L':
				direction -= amount / 90
				if direction < 0 {
					direction += 4
				}
				direction %= 4
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(abs(x) + abs(y))

}

