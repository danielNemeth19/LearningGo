package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]string)

	fmt.Println("Enter your name:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	m["name"] = scanner.Text()

	fmt.Println("Enter your address (on a single line):")
	scanner.Scan()
	m["address"] = scanner.Text()

	jval, _ := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(jval))
}
