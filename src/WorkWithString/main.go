package main


import(
	"fmt"
)
// func isCircularSentence(sentence string) bool {
//     // встречается пробел - обрабатываем проверяем вокруг пробела значения
//     // конец строки проверяем с начала
//     if sentence[len(sentence)-1] != sentence[0]{
//         return false   
//     }
//     for i:= range sentence{
//         if sentence[i] == ' ' {
//             if sentence[i-1] != sentence[i+1]{
//                 return false
//             } 
//         }
//     }
//     return true
// }
////----------------------------------------------------------------------------------------------------

func main(){

		findKthBit(20, 1048575)

}

// func searchSPrev(n int) string{
//     if n == 1{
//         return "0"
//     }
//     return searchSPrev(n-1)+"1"+reverse(invert(searchSPrev(n-1)))

// }


func findKthBit(n int, k int) byte {
    //var str = searchSPrev(n)
    var str string = "0"

    for i:=1;i<n;i++{
        str += "1" + reverse(invert(str))
    }
	fmt.Println(str)
    if str[k-1] == '0'{
        return 48
    } else {
        return 49
    }
}

//0 -- 1
//0 1 1 -- 2
//011 1 001 -- 3

// func searchSPrev(n int) string{
//     if n == 1{
//         return "0"
//     }
//     return searchSPrev(n-1)+"1"+reverse(invert(searchSPrev(n-1)))
    // красиво - но долго:(
// }


func invert(str string) string{
    tmpstr := ""
    for i:=range str{
        if str[i] == '0'{
            tmpstr += "1"
        } else{
            tmpstr += "0"
        }
    }
    return tmpstr
}
func reverse(str string) string {
    runes := []rune(str)
    length := len(runes)
    for i := 0; i < length/2; i++ {
        runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
    }

    return string(runes)
}