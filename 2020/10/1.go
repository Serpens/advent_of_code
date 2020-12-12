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
	numbers := make([]int, 0)
	for scanner.Scan() {
		line = scanner.Text()
		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	sort.Ints(numbers)
	dif1 := 1
	dif3 := 1
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i+1] - numbers[i] == 1 {
			dif1++
		} else if numbers[i+1] - numbers[i] == 3 {
			dif3++
		}
	}
	fmt.Println(dif1 * dif3)
}

