package main

import (

    "fmt"

)

func main() {
    var number int
    _, _ = fmt.Scan(&number)
 
    myMap := make(map[int]int)
    var tmp int
    for i:=0; i<number; i++{
    _, _ = fmt.Scan(&tmp)
    myMap[tmp] = myMap[tmp] + 1
    }

    fmt.Println(len(myMap))
    

}
