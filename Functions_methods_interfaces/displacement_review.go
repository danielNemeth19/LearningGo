package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("acceleration: ")
	acceleration := userInput()

	fmt.Print("velocity: ")
	velocity := userInput()

	fmt.Print("displacement: ")
	displacement := userInput()

	fmt.Print("time (seconds): ")
	time := userInput()

	fn := GenDisplaceFn(acceleration, velocity, displacement)
	fmt.Println(fn(time))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return (a/2)*(t*t) + (v * t) + s
	}
}

func userInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error. Please, enter only integers or floats")
		os.Exit(1)
	}

	return num
}
