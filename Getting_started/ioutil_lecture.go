package main
import (
    "fmt"
    "io/ioutil"
    "os"
    )


func main(){
    // data is []byte (byte array) filled with contents of entire file
    // open/close operations are built-in to ReadFile function
    data, err := ioutil.ReadFile("arrays.go")
    if err != nil {
        fmt.Println("Problem with opening file")
    } else {
        fmt.Println(len(data))
    }
    
    // writes []byte to file, unix-style permission bytes
    is_save_ok := ioutil.WriteFile("arrays_dump.go", data, 0777)
    if is_save_ok == nil {
        fmt.Println("Dump file saved..now removing")
        os.Remove("arrays_dump.go")
    }
}
