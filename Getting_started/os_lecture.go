package main
import (
    "fmt"
    "os"
    "log"
)

func check_error(e error) {
    if e != nil {
        log.Fatal(e)
    }
}


func example_read(){
    f, err := os.Open("arrays.go")
    check_error(err)
    // os.Read reads and fills byte array - returns # of bytes read
    barr := make([]byte, 10)
    nb, err := f.Read(barr)
    check_error(err)
    f.Close()
    fmt.Printf("Read %d bytes: %q\n", nb, barr)
}


func example_write(){
    outfile, err := os.Create("outfile.txt")
    check_error(err)
    barr_out := []byte{100, 101}
    nb_out, err_out := outfile.Write(barr_out)
    check_error(err_out)
    fmt.Printf("%d\n", nb_out)
}


func example_writestring(){
    outfile, err := os.OpenFile("outfile.txt", os.O_APPEND|os.O_WRONLY, 0755)
    check_error(err)
    nb_out, err := outfile.WriteString(" hello")
    check_error(err)
    fmt.Printf("%d\n", nb_out)
   
    f, err := os.Open("outfile.txt")
    check_error(err)
    barr := make([]byte, 8)
    nb, err := f.Read(barr)
    check_error(err)
    f.Close()
    fmt.Printf("Read %d bytes: %q\n", nb, barr)
}


func cleanup(){
    os.Remove("outfile.txt")
}


func main(){
    example_read()
    example_write()
    example_writestring()
    cleanup()
}
