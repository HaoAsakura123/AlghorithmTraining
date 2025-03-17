package main

import "fmt"

//first

func mergeAlternately(word1 string, word2 string) string {
	n, m := len(word1), len(word2)
	b := make([]byte, 0, n+m) 

	i, j := 0, 0
	for i < n || j < m {
		if i < n {
			b = append(b, word1[i])
			i++
		}
		if j < m {
			b = append(b, word2[j])
			j++
		}
	}

	return string(b)
}

//second


func gcdOfStrings(str1 string, str2 string) string {
    a, b := len(str1), len(str2)
	if a <= b{
		a, b = b, a
	}
	// a > b
	// a = bq + r1 => if r1 / 
	if str1 + str2 != str2+str1{
		return ""
	}
	nod := gcd(a,b)
	return str1[:nod]
}

func gcd(n int, m int) int{
	for n != m{
		n = n-m
		if n < m{
			n, m = m, n
		}
	}
	return n
}

// thirst

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
  
	for _, candy := range candies{
	  if candy > max {
		  max = candy
	  }
	}
	result := make([]bool, 0, len(candies))
	  for _, candy := range candies{
	  if candy >= max - extraCandies {
		  result = append(result, true)
	  } else{
		  result = append(result, false)
	  }
  
	}
  
	return result
  }

//fourth

func reverseVowels(s string) string {
    right:=len(s)-1
    left:=0
    vowels := map[rune] bool{
        'a': true, 'e': true, 'i': true, 'o': true, 'u':true,
        'A': true, 'E': true, 'I': true, 'O': true, 'U':true,        
    }
    runes := []rune(s)
    // while( right > left )
    for i:=0; left<right && i<len(runes);i++{
        if vowels[runes[left]] && vowels[runes[right]] {
            runes[left], runes[right] = runes[right], runes[left]
            left++
            right--
        } else{
            if vowels[runes[left]] == false { left++ }
            if vowels[runes[right]] == false { right-- }
        }
    }
    return string(runes)
}

//fifth

func reverseWords(s string) string {
    words := strings.Fields(s)
    result:= ""
    for i:=len(words) - 1; i>=0;i--{
        if i!=len(words) - 1{
            result+=" "
        }
        result+=words[i]
    }
    return result
}