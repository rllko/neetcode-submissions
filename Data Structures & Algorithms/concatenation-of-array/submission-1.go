func getConcatenation(nums []int) []int {
    arr := make([]int,len(nums)*2)

    copy(arr,nums)
    i := len(nums)
    for _,num := range nums {
        arr[i] = num
        i++
    }

    return arr
}
