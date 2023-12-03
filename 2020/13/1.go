package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	scanner.Scan()
	line = scanner.Text()
	earliestTime, _ := strconv.Atoi(line)
	scanner.Scan()
	line = scanner.Text()

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	splits := strings.Split(line, ",")
	var earliestLine int
	waitingTime := 10000000000
	for _, s := range splits {
		if s == "x" {
			continue
		}
		bus, _ := strconv.Atoi(s)
		currentWaitingTime := bus - earliestTime % bus
		if bus == currentWaitingTime {
			fmt.Println(0)
			return
		}
		if currentWaitingTime < waitingTime {
			waitingTime = currentWaitingTime
			earliestLine = bus
		}
	}

	fmt.Println(earliestLine * waitingTime)
}

