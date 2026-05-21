func getConcatenation(nums []int) []int {
    arr := make([]int,len(nums)*2)

    copy(arr,nums)
    copy(arr[len(nums):],nums)

    return arr
}