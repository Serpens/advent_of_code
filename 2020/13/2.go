package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)


type Bus struct {
	line int
	delay int
}


func mulInv(a, b int) int {
	b0 := b
	x0 := 0
	x1 := 1
	var temp int
	if b <= 1 {
		return 1
	}
	for a > 1 {
		q := a / b

		temp = b
		b = a%b
		a = temp

		temp = x0
		x0 = x1 - q * x0
		x1 = temp
	}
	if x1 < 0 {
		x1 += b0
	}
	return x1
}


func chinese(buses []Bus) int {
	sum := 0
	prod := 1
	for _, bus := range buses {
		prod *= bus.line
	}
	for _, bus :=  range buses {
		p := prod
		sum += bus.delay * mulInv(p, bus.line) * p
	}
	return sum % prod
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	scanner.Scan()
	line = scanner.Text()
	scanner.Scan()
	line = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	splits := strings.Split(line, ",")
	buses := make([]Bus, 0)
	for i, s := range splits {
		if s == "x" {
			continue
		}
		lineNumber, _ := strconv.Atoi(s)
		currentBus := Bus{
			line: lineNumber,
			delay: i,
		}
		buses = append(buses, currentBus)
	}

	fmt.Println(chinese(buses))
}

