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