package main
import "fmt"


func main(){
    var idMap map[string]int

    // with the help of make a map can be created
    idMap = make(map[string]int)
    fmt.Printf("%T\n", idMap)

    // map can be created as a map literal as well
    idMapLiteral := map[string]int {"Joe": 123}
    fmt.Printf("Adding Joe with value: %d\n", idMapLiteral["Joe"])
    
    // trying to access a key that doesn't exists returns 0
    fmt.Printf("Value of non-existent key is always: %d\n", idMapLiteral["NoSuchKey"])
    
    // adding new key-value pair
    idMapLiteral["Sally"] = 345
    fmt.Printf("Adding Sally with value: %d\n", idMapLiteral["Sally"])
    
    // Two value assignment can be used to test existence of the NoSuchKey
    // id will be the value of the key, p is a boolean
    id, p := idMapLiteral["Joe"]
    fmt.Printf("Is Joe present in the map: %t\n", p)
    fmt.Printf("Value of Joe: %d\n", id)
    
    // deleting a key-value pair
    delete(idMapLiteral, "Joe")
    _, p2 := idMapLiteral["Joe"]
    fmt.Printf("After delete - is Joe present in the map: %t\n", p2)
        
    // len can be used to return number of keys
    idMapLiteral["Kate"] = 987
    fmt.Println(len(idMapLiteral))
    
    // Iterating through a map - use a for loop with the range keyword (two-value assignment with range)
    fmt.Printf("Iteration of idMapLiteral:\n")
    for key, value := range idMapLiteral {
        fmt.Printf("\t%s, %d\n", key, value)
    }
}
