package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
)


func check(number int, previous []int) bool {
	for _, n1 := range previous {
		for _, n2 := range previous {
			if n1 != n2 && n1 + n2 == number {
				return true
			}
		}
	}
	return false
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	numbers := make([]int, 0)
	for scanner.Scan() {
		line = scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for i := 25; i < len(numbers); i++ {
		if !check(numbers[i], numbers[i-25:i]) {
			fmt.Println(numbers[i])
			return
		}
	}
}

