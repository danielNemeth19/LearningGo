package main

import (
	"fmt"
	"time"
)

// mimic "slow" retrieval of value
func delayValue(v *int) int {
	time.Sleep(time.Millisecond * 5)
	return *v
}

func isEven(v *int) {
	n := delayValue(v)
	if n % 2 == 0 {
		fmt.Println("v is even", n)
	} else {
		fmt.Println("v is odd", n)
	}
	*v = *v * 2
}

func isPositive(v *int) {
	n := delayValue(v)
	fmt.Println(*v, n)
	if n >= 0 {
		fmt.Println("v is positive", n)
	} else {
		fmt.Println("v is negative", n)
	}
	*v = *v + 1
}

func main() {
	v := 10
	
	fmt.Println(v)

	go isEven(&v)
	go isPositive(&v)

	time.Sleep(time.Second)
	fmt.Println(v)
}
