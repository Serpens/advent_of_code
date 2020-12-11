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
			for k := 0; k < length; k++ {
				if numbers[i] + numbers[j] + numbers[k] == 2020 {
					fmt.Println(numbers[i] * numbers[j] * numbers[k])
				}
			}
		}
	}
}

