package main

import(
	"fmt"
)

func main(){
	height := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(height))
}

//two pointers 
func maxArea(height []int) int {
	max:=0
    l:=0
    r:=len(height)-1
	for l < r {
        if minHeight(height[l], height[r]) * (r-l) >= max{
            max = minHeight(height[l], height[r]) * (r-l)
        }
        if height[l]<height[r]{
            l++
        } else{
            r--
        }
	}
    //max = minHeight(height[i], height[j]) * (j-i)
	return max
}

func minHeight(a int, b int) int{
	if a >= b{
		return b
	} else{
		return a
	}
}