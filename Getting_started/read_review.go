package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type person struct {
	fname string
	lname string
}

func main() {
	var s []person
	var f string
	fmt.Printf("Enter the File Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	f = scanner.Text()
	data, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	scanner_file := bufio.NewScanner(data)
	for scanner_file.Scan() {
		line := scanner_file.Text()
		full_name := strings.Fields(line)
		var p person
		p.fname = full_name[0]
		p.lname = full_name[1]
		s = append(s, p)
	}
	for _, n := range s {
		fmt.Println("First Name:", n.fname, ",", "Last Name:", n.lname)
	}
}
