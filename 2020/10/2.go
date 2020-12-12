package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
	"strconv"
	"sort"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var line string
	numbers := make([]int, 1)
	numbers[0] = 0
	for scanner.Scan() {
		line = scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1] + 3)
	ways := make([]int, len(numbers))
	ways[0] = 1

	for i := 1; i < len(numbers); i++ {
		for j := 1; i - j >= 0 && j <= 3; j++ {
			if numbers[i] - numbers[i -j] > 3 {
				break
			}
			ways[i] += ways[i - j]
		}
	}
	fmt.Println(ways[len(ways)-1])
}

