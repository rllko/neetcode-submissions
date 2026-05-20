func replaceElements(arr []int) []int {
    // get biggest array to the
    n := len(arr)
    ans := make([]int,n)

    for i,_ := range arr{
        biggest := -1
        for j:= i+1; j < n; j++ {
            if(arr[j] > biggest){
                biggest = arr[j] 
            }
        }
        ans[i] = biggest
    }

    
    ans[len(arr)-1] = -1
    return ans;
}
