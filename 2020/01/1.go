package main

import "fmt"

const length = 200
var numbers [200] int

func main() {
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &numbers[i])
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if numbers[i] + numbers[j] == 2020 {
				fmt.Println(numbers[i] * numbers[j])
			}
		}
	}
}

