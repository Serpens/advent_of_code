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


func sumUntilTarget(target int, slice []int) ([]int, bool) {
	sum := 0
	for i := 0; i < len(slice); i++  {
		sum += slice[i]
		if sum == target {
			return slice[0:i+1], true
		} else if sum > target {
			return slice[0:i+1], false
		}
	}
	return slice, false
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

	var target int
	for i := 25; i < len(numbers); i++ {
		if !check(numbers[i], numbers[i-25:i]) {
			target = numbers[i]
			break
		}
	}


	var subset []int
	var equal bool
	for i := 0; i < len(numbers); i++ {
		subset, equal = sumUntilTarget(target, numbers[i:])
		if equal {
			break
		}
	}

	min := target
	max := 0
	for _, i := range subset {
		if min > i {
			min = i
		} else if max < i {
			max = i
		}
	}

	fmt.Println(min + max)
}

