package main
import (
    "fmt"
    "encoding/json"
)


type Person struct  {
    Name string
    Addr string
    Phone string
}

//using struct tags we can set the names of the struct's fields when marshaling to and from json
type Bird struct {
     Species string `json:"birdType"`
     Description string `json:"what it does"`
}


func main(){
    p1 := Person{Name: "Joe", Addr:"Washington", Phone: "+4470258"}
    fmt.Println(p1.Name, p1.Addr, p1.Phone)
    person_barr, _ := json.Marshal(p1)
    fmt.Println(string(person_barr))
    
    pigeon := Bird{Species: "Pigeon", Description: "Likes to eat seed"}
    fmt.Println(pigeon.Species, pigeon.Description)
    bird_barr, _ := json.Marshal(pigeon)
    fmt.Println(string(bird_barr))

    var p2 Person
    json.Unmarshal(person_barr, &p2)
    fmt.Println(p2.Name, p2.Addr, p2.Phone)

    var pigeon2 Bird
    json.Unmarshal(bird_barr, &pigeon2)
    fmt.Println(pigeon2.Species, pigeon2.Description)
}
