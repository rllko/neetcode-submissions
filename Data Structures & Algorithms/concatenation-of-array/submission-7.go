func getConcatenation(nums []int) []int {
    l := len(nums)
    arr := make([]int, 0, l*2)

    for i :=0 ; i < 2; i ++ {
        for _,num := range nums {
            arr = append(arr,num) 
        }
    }

    return arr
}
