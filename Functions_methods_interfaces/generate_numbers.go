package main

import (
    "fmt"
    "math/rand"
    "time"
    "strings"
    "strconv"
)


func generate_number() int {
    rand.Seed(time.Now().UnixNano())
    min := -100
    max := 500
    number := rand.Intn(max - min + 1) + min
    return number
}
        
func main() {
    numbers := make([]string, 0, 100)
    for i:= 0; i < 100; i++ {
        number := generate_number()
        s := strconv.Itoa(number)
        numbers = append(numbers, s)
    }
    output := strings.Join(numbers, `,`)
    fmt.Println(output)
}
