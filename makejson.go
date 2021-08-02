package main
import (
    "fmt"
    "os"
    "bufio"
    "encoding/json"
)

func get_map() map[string]string {
    var name string
    var address string
    
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Enter a name:")
    scanner.Scan()
    name = scanner.Text()
    
    fmt.Println("Enter the address:")
    scanner.Scan()
    address = scanner.Text()
    
    nameMap := make(map[string]string)
    nameMap["name"] = name
    nameMap["address"] = address
    return nameMap
}


func map_2_json(m map[string]string) []byte {
    js_m, _ := json.MarshalIndent(m, "", "  ")
    return js_m
}


func main(){
    nameMap := get_map()
    map_as_json := map_2_json(nameMap)
    fmt.Println(string(map_as_json))
}
