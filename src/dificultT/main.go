package main

import (
	"fmt"
)
// tps://coderun.yandex.ru/selections/backend/problems/diversity-improvement/ не пошла :(
func findMaxIndex(sliceOfSlices [][]int, iskl int) int {
	maxLen := 0
	maxIndex := -1

	for i, slice := range sliceOfSlices {
		if len(slice) > maxLen && (iskl == -1 || iskl != i) {
			maxLen = len(slice)
			maxIndex = i
		}
	}

	return maxIndex
}

func main() {

	//1 2 3 : 1
	//4 5 6 : 2
	//7 8 9 : 3
	matrix := make(map[int][]int)
	matrixResult := make([][]int, 0)
	var number int
	_, _ = fmt.Scan(&number)
	maxLen := 0
	for i := 0; i < number; i++ {
		var key, value int
		_, _ = fmt.Scan(&value, &key)
		matrix[key] = append(matrix[key], value)
		if len(matrix[key]) > maxLen {
			maxLen = len(matrix[key])
		}
	}
	for _, values := range matrix {
		matrixResult = append(matrixResult, values)
	}
	k := -1
	for i := 0; i < number; i++ {
		k = findMaxIndex(matrixResult, k)
		fmt.Println(matrixResult[k][0])
		matrixResult[k] = matrixResult[k][1:]
	}
	// for i:=0;i<maxLen;i++{
	// 		for index:=range matrix{
	// 			if len(matrix[index])>0{
	// 				//fmt.Println(matrix[index][0])

	// 				matrix[index] = matrix[index][1:]
	// 			}
	// 		}
	// }

}
