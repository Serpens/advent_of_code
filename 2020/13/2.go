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

	var allOnTime bool
	for i := 0; true; i += buses[0].line {
		allOnTime = true
		for j := 1; j < len(buses); j++ {
			if buses[j].line - i % buses[j].line != buses[j].delay {
				allOnTime = false
				break
			}
		}
		if allOnTime {
			fmt.Println(i)
			return
		}
	}

}

