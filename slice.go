package main
import (
    "fmt"
    "sort"
    "os"
    "bufio"
    "strconv"
    "strings"
)


func _convert(s string) string {
    return strings.Trim(s, "\n")
}


func original_solution(){
    my_slice := make([]int, 0, 3)
    for true {
        fmt.Println("Enter an integer!")
        in := bufio.NewReader(os.Stdin)
        inputSting, _ := in.ReadString('\n')
        s := _convert(inputSting)
        
        if s == "X" || s == "x" {
            fmt.Println("Exit signal received")
            break
        }
        
        number, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        if err == nil {
            my_slice = append(my_slice, number)
            sort.Ints(my_slice)
            fmt.Println(my_slice)
        }
    }
}

func bufio_newscanner_solution(){
    my_slice := make([]int, 0, 3)
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Println("Enter an integer!")
        // Scans a line from stdin(console)
        scanner.Scan()
        // Holds the string that scanned
        s := scanner.Text()
        if  s == "X" || s == "x" {
            fmt.Println("Exit signal received")
            break
        }
        number, err := strconv.Atoi(s)
        if err != nil {
            continue
        }
        if err == nil {
            my_slice = append(my_slice, number)
            sort.Ints(my_slice)
            fmt.Println(my_slice)
        }
    }
}


func scanf_solution(){
    var s string
    my_slice := make([]int, 0, 3)
    for s != "x" && s != "X" {
        fmt.Println("Enter an integer!")
        fmt.Scanf("%s", &s)
        number, err := strconv.Atoi(s)
        if err != nil {
            continue
        } else {
            my_slice = append(my_slice, number)
            sort.Ints(my_slice)
            fmt.Println(my_slice)
        }
    }
}


func main(){
    original_solution()
    bufio_newscanner_solution()
    scanf_solution()
}
