package main

import(
	"fmt"
	"strconv"
)


func main(){

	chars := []byte{'a','a','b','b','c','c','c'}
	fmt.Println(compress(chars))
}


func compress(chars []byte) int {
    var cur byte = byte(chars[0])
    leng := 0
    result := make([]byte, 0)
    for _, elem:= range chars{
        if cur == byte(elem){
            leng++
        } else{
            result = append(result, byte(cur))
            if leng>1{
                bs := []byte(strconv.Itoa(leng))
                for _, chs := range bs{
                   result = append(result, byte(chs)) 
                }
			leng = 1
            }
            cur = byte(elem)
        }
    }
	result = append(result, byte(cur))
	if leng>1{
		bs := []byte(strconv.Itoa(leng))
		for _, chs := range bs{
		   result = append(result, byte(chs)) 
		}
	}
	//fmt.Println(string(result))
	copy(chars, result)
	//fmt.Println(chars)
    return len(result)
}