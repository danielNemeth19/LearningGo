package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func _convert(s string) string {
	return strings.Trim(strings.ToLower(s), "\n")
}

func _check_if_complies(mys string) bool {
	flag := false
	if strings.HasPrefix(mys, "i") && strings.HasSuffix(mys, "n") && strings.ContainsAny(mys, "a") {
		flag = true
	}
	return flag
}

func main() {
	fmt.Printf("Enter a string:\n")

	in := bufio.NewReader(os.Stdin)
	inputSting, err := in.ReadString('\n')
	if err != nil {
		fmt.Printf("%T, %v\n", err, err)
	}

	s := _convert(inputSting)
	if _check_if_complies(s) {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not found!\n")
	}
}
