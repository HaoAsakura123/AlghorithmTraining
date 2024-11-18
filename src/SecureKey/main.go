package main

import (
	"fmt"
)

func main() {    
    //НОК = x*y/НОД(x,y)

    //1 - k*m*NOD = NOK => NOK/NOD = k*m   (if NOK/NOD != 0 res = 0)
    //2 - for 0 to sqrt(k*m) if k*m/i == 0 => res +=2
	var numberNOD int64
	_, _ = fmt.Scan(&numberNOD)
	var numberNOK int64
	_, _ = fmt.Scan(&numberNOK)
	if numberNOK % numberNOD != 0{
		fmt.Println(0)
		return
	}
	var km = int64(numberNOK / numberNOD)
	res := 0
	for i:=1; int64(i*i)<=km;i++{
		if km % int64(i) == 0{

			if NOD(int64(i) * numberNOD, int64(km / int64(i))* numberNOD) == numberNOD && (int64(i) * numberNOD) != (int64(km / int64(i)) * numberNOD){
				res+=2
				//fmt.Println(int64(i) * numberNOD, int64(km / int64(i))* numberNOD)
			} else if NOD(int64(i) * numberNOD, int64(km / int64(i))* numberNOD) == numberNOD{
				res++
			}
			//println(int64(i) * numberNOD, int64(km / int64(i)) * numberNOD, NOD(int64(i) * numberNOD, int64(km / int64(i))* numberNOD))
			//println(int64(i), int64(km / int64(i)))
		}
	}
	fmt.Println(res)
}



func NOD(x int64, y int64) int64 {
    for y != 0 {
        x, y = y, x % y
    }
    return x
}
// 1 1000 +
// 2 500 -
// 4 250 -
// 5 200 - 
// 8 125 +
// 