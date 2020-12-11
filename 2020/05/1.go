package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var line string


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
	max_id := 0
	for scanner.Scan() {
		line = scanner.Text()
		row_part := line[0:7]
		column_part := line[7:10]
		current_id := 8 * decode(row_part, 7, 'B') + decode(column_part, 3, 'R')
		if current_id > max_id {
			max_id = current_id
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Println(max_id)
}

