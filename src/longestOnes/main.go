package main

import(
	"fmt"
)

func main(){
	nums := []int{1, 0, 0, 0}
	fmt.Println(longestOnes(nums, 2))
}

func longestOnes(nums []int, k int) int {
    curr := make([]int, 0)
    pupu :=0
    maxElement := 0
    for _, elem:= range nums{
        if elem == 0{
            if pupu !=0 {
                curr = append(curr, pupu)
                if pupu > maxElement{
                    maxElement = pupu
                }
                pupu = 0
                curr = append(curr, 0)
            } else{
                curr = append(curr, 0)
            }
        } else {
            pupu++
        }
    }
    if pupu != 0{
        curr = append(curr, pupu)
    }
    if pupu > maxElement{
        maxElement = pupu
    }
    l := -1
    r := 0
    pupu = 0
    MaxResult := 0
    iter := 0
    for iter < k && iter < len(curr) && r < len(curr){
        if curr[r] == 0{
            iter++
        }
        
        pupu += curr[r]
        if pupu > MaxResult{
            MaxResult = pupu
        }
        if iter < k {
            r++
        }
    }
    if curr[0] == 0{
        l = 0
    } else {
        l = 1
    }
    pupu = pupu + rCheck(curr, r) + k
	if pupu > MaxResult{
		MaxResult = pupu
	}
    // 3 0 0 0 4 0
    //Нужно в цикле пройти
    fmt.Println(l, r, MaxResult)
    //fmt.Println(curr)
    if pupu >= len(nums){
        return len(nums)
    }
    if k == 0{
        return maxElement
    }
    r++
    for r < len(curr){
        if curr[r] == 0 {
            //l --> нужна функция, которая найдет следующий ноль
            pupu = pupu - lCheck(curr, l)
            if curr[l+1] == 0{
                l++
            } else{
                l+=2
            }
            //r -->
            pupu = pupu + rCheck(curr, r)
            if r+1 < len(curr) && curr[r+1] == 0{
                r++
            } else{
                r+=2
            }
            //
            if pupu > MaxResult{
                MaxResult = pupu
            }
        } else {
            r++
        }
        // if r+1<len(curr){
        //     r++
        //     fmt.Println("ya tuta")
        //     pupu = pupu + rCheck(curr, r)
        //     pupu = pupu - lCheck(curr, l)
        //     if curr[l+1] == 0{
        //         l++
        //     } else{
        //         l+=2
        //     }
        //     if pupu > MaxResult{
        //         MaxResult = pupu
        //     }
        // }
    }
    if k == len(nums){
        return k
    }
    return MaxResult
}

func rCheck(mas []int, point int) int{
    if point + 1 < len(mas){
        return mas[point+1]
    }
    return 0
}

func lCheck(mas []int,point int) int{
    if point - 1 >= 0{
        return mas[point-1]
    }
    return 0
}