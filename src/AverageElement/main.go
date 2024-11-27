package main

import (
    "fmt"
)

func main() {

    var number1 int
    _, _ = fmt.Scan(&number1)
    var number2 int
    _, _ = fmt.Scan(&number2)
    var number3 int
    _, _ = fmt.Scan(&number3)
    // 1 и 2 сравниваем, потом сравниваем 3 и полученное
    if number1 >= number2{
        if number1 >= number3{
            if number3 >= number2{
                fmt.Println(number3)
            } else{
                fmt.Println(number2)
            }
        }
    } else if number2 >= number3{
        if number3>=number1{
            fmt.Println(number3)
        } else{
            fmt.Println(number1)
        }
    } else{
		fmt.Println(number2)
	}

}
