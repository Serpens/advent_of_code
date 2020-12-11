package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var line string

func main() {
	tree_map := make([][31]bool, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		var tree_line_bool [31]bool
		for i, char := range line {
			tree_line_bool[i] = (char == rune('#'))
		}
		tree_map = append(tree_map, tree_line_bool)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	var propositions_right = [...]int {1, 3, 5, 7, 1}
	var propositions_down = [...]int {1, 1, 1, 1, 2}
	result := 1
	for i := 0; i < 5; i++ {
		x := 0
		right := propositions_right[i]
		down := propositions_down[i]
		tree_counter := 0
		for y := 0; y < len(tree_map); y += down {
			if tree_map[y][x] {
				tree_counter++
			}
			x = (x + right) % 31
		}
		result *= tree_counter
	}
	fmt.Println(result)
}

