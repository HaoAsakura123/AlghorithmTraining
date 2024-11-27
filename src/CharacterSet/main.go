// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSpace(input)
// 	dictionary, _ := reader.ReadString('\n')
// 	dictionary = strings.TrimSpace(dictionary)
// 	dictionaryMap := make(map[int]int)
// 	for i := range dictionary {
// 		dictionaryMap[int(dictionary[i])-97] = 1
// 	}
// 	minLenght := len(input) + 1;
// 	mapsSlice := []map[int]int{}
// 	currentMap := make(map[int]int)
// 	mapsSlice = append(mapsSlice, make(map[int]int))

// 	for i := range input {
// 		newMap := make(map[int]int)
// 		for key, value := range currentMap {
// 			newMap[key] = value
// 		}
// 		charKey := int(input[i]) - 97
// 		newMap[charKey] += 1
// 		mapsSlice = append(mapsSlice, newMap)
// 		currentMap = newMap
// 	}
// 	for i:=0; i<len(mapsSlice)-1; i++{
// 		for j:=i+1; j<len(mapsSlice); j++{
// 			flag := 1
// 			for index, _ := range dictionaryMap{
// 				if mapsSlice[j][index] <= mapsSlice[i][index]{
// 					flag = 0
// 				}
// 			}
// 			if flag == 1 && j-i<minLenght{
// 				minLenght = j-i
// 			}
// 		}	
// 	}
// 	if(minLenght>len(input) || len(input) > 100 || len(dictionary) > 26 ){
// 		fmt.Println(0)
// 	} else {
// 		fmt.Println(minLenght)
// 	}
// }

package main

import (
	"fmt"
)

func findMinSubstringLength(s, c string) int {
	targetSet := make(map[rune]bool)
	for _, ch := range c {
		targetSet[ch] = true
	}
	n := len(s)
	minLength := n + 1 

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			
			substringSet := make(map[rune]bool)
			for k := i; k <= j; k++ {
				substringSet[rune(s[k])] = true
			}

			if len(substringSet) == len(targetSet) {
				match := true
				for key := range targetSet {
					if !substringSet[key] {
						match = false
						break
					}
				}

			
				if match {
					minLength = min(minLength, j-i+1)
				}
			}
		}
	}


	if minLength == n+1 {
		return 0
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	var s, c string
	fmt.Scanln(&s)
	fmt.Scanln(&c)

	fmt.Println(findMinSubstringLength(s, c))
}
