package main

import(
	"math"
)

func main(){



}

func getMaximumXor(nums []int, maximumBit int) []int {

	kumSum := 0
	maxValue:=int(math.Pow(2, float64(maximumBit))- 1)
	for _, r:=range nums{
		kumSum = kumSum ^ r
	}
	var result []int
	for i:=range nums{
		result = append(result, kumSum^maxValue)
		kumSum = kumSum ^ nums[len(nums) - 1 - i]
	}
	return result
	}
	
	//111111
	//110101
	//------
	//001010