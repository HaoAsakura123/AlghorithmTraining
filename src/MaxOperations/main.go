// func maxOperations(nums []int, k int) int {
//     sort.Ints(nums)
//     result := 0
//     l:=0
//     r:=len(nums)-1

//     for l < r{
//         if nums[l] + nums[r] == k {
//             nums = append(nums[:r], nums[r+1:]...)
//             nums = append(nums[:l], nums[l+1:]...)
//             r = r - 2
//             result++
//         } else if nums[l] + nums[r]<k{
//             l++
//         } else {
//             r--
//         }
//     }
//     return result
// }