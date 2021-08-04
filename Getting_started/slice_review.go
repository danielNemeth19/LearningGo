package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numberStr string
	numbers := make([]int, 0, 3)

	for numberStr != "x" && numberStr != "X" {
		fmt.Print("Enter an integer (or 'x' to exit) - ")
		fmt.Scanf("%s", &numberStr)
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			continue
		} else {
			numbers = append(numbers, number)
			sort.Slice(numbers, func(i, j int) bool {
				return numbers[i] < numbers[j]
			})
			fmt.Printf("Sorted slice - %v\n", numbers)
		}
	}
}
