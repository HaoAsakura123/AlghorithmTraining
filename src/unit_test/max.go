package main

import "fmt"


func main(){
	fmt.Println(Max([]int{1,14,16, 32,55}))
}

func Max(numbers []int) int{
	var max int

	for i := range numbers{
		if numbers[i]>max{
			max = numbers[i]
		}
	}
	return max
}