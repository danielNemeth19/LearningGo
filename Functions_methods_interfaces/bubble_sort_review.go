package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbersToSort := getInput()
	BubbleSort(numbersToSort)
	fmt.Println("Sorted numbers:", numbersToSort)
}

func BubbleSort(numbers []int) {
	for anchor := len(numbers); anchor > 1; anchor-- {
		for traveller := 0; traveller < anchor-1; traveller++ {
			if numbers[traveller] > numbers[traveller+1] {
				Swap(numbers, traveller)
			}
		}
	}
}

func Swap(numbers []int, position int) {
	numbers[position], numbers[position+1] = numbers[position+1], numbers[position]
}

func getInput() []int {
	fmt.Println("Enter numbers below with a space in between. For example:- 10 73 87 45.")
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)

	var numbers = make([]int, 0)

	if scanner.Scan() {
		line := scanner.Text()

		inputNumbers := strings.Split(line, " ")

		for _, value := range inputNumbers {
			intValue, e := strconv.Atoi(value)
			check(e)
			numbers = append(numbers, intValue)
		}
	}

	return numbers
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
