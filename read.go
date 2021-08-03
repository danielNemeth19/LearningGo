package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)


type Name struct  {
    fname string
    lname string 
}


func set_name_parts (full_name []string) (string, string) {
    if len(full_name) != 2 {
        fmt.Printf("Line should have two parts: got %d\n", len(full_name))
        os.Exit(1)
    }
    
    for _, v := range(full_name) {
        if invalid_field_name(v) {
            fmt.Printf("Name part should be less than 21 chars: got %d (%s)\n", len(v), v)
            os.Exit(1)
        }
    } 
    fname := full_name[0]
    lname := full_name[1]
    return fname, lname
}


func invalid_field_name(f string) bool {
    if len(f) > 20 {
        return true
    }
    return false
}


func get_path() string {
    var file_path string
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Please specify the absolute path to the names file:")
    scanner.Scan()
    file_path = scanner.Text()
    return file_path
}


func create_name_slice(fp string) []Name {
    name_slice := make([]Name, 0, 5)
    file, err := os.Open(fp)
    if err != nil {
        fmt.Printf("Problem with opening file: %v", err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        full_name := strings.Fields(scanner.Text())
        fname, lname := set_name_parts(full_name)
        name_obj := Name{fname: fname, lname:lname}
        name_slice = append(name_slice, name_obj)
    }
    return name_slice
}


func main(){
    file_path := get_path()
    name_slice := create_name_slice(file_path)
    for i, v := range name_slice {
        fmt.Printf("Name#: %d First name: %s --> Last name: %s\n", i, v.fname, v.lname)
    }
}

