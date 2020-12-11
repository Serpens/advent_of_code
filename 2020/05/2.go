package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var line string
const seats = 1024

func decode(encoded string, length int, onechar rune) int {
	result := 0
	current_digit := 2 << (length - 2)
	for _, c := range encoded {
		if c == onechar {
			result += current_digit
		}
		current_digit /= 2
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var taken [seats]bool
	for i := 0; i < seats; i ++ {
		taken[i] = false
	}
	for scanner.Scan() {
		line = scanner.Text()
		row_part := line[0:7]
		column_part := line[7:10]
		current_id := 8 * decode(row_part, 7, 'B') + decode(column_part, 3, 'R')
		taken[current_id] = true
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for i := 1; i < seats; i++ {
		if !taken[i] && taken[i-1] && taken[i+1] {
			fmt.Println(i)
			break
		}
	}
}

