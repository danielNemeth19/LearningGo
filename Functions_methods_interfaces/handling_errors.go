package main
import (
    "fmt"
    "os"
)

/* Error interface
 * Many Go programs return error interface objects to indicate errors
 * type error interface {
 *      Error() string
 * }
 * Correct operation: error == nil
 * Incorrect operation: Error() prints error message
 */
func main(){
    _, err := os.Open("no/file/like/this.txt")
    if err != nil {
        // err is type error and Println will call the Error method of err
        fmt.Println(err)
        fmt.Printf("%s, %s\n", err.Error())
    }
}
