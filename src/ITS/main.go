package main

import(
	"fmt"
)

func main(){
	ints := []int{4,10,5,1,2,3,-9}
	// ints = append(ints, 20)
	// ints = append(ints, 100)
	// ints = append(ints, 10)
	// ints = append(ints, 12)
	// ints = append(ints, 5)
	// ints = append(ints, 13)
	fmt.Println(increasingTriplet(ints))
}

func increasingTriplet(nums []int) bool {
    mapa := make(map[int] []int)

    for i, elem := range nums{
        mapa[elem] = append(mapa[elem], i)
    }
	if len(mapa)<3{
		return false
	}
	fmt.Println(mapa)
    return solveTriplets(mapa) 
}

func solveTriplets(myMap map[int][]int) bool{
	var solve bool = false
	for i, elem:= range myMap{
		for _, k:= range elem{
			if up(myMap, i, k) == true && down(myMap, i, k) == true{
				return true
			}
		}
	}
	return solve
}
// id in map[qwe] = qwe  || i = qwe
// elem = []int
// element = qwe
// id = in []int
func up(myMap map[int][]int, element int, id int) bool{
	
	for i, elem:= range myMap{
		if i <= element{
			continue
		}
		for _, k:= range elem{
			if k > id{
				return true
			}
		}
	}
	return false
}

func down(myMap map[int][]int, element int, id int) bool{
	for i, elem:= range myMap{
		if i < element{
			for _, k:= range elem{
				if k < id{
					return true
				}
			}
		}
	}
	return false
}

// 4 3 1 2 5 1