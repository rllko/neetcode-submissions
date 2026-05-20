func replaceElements(arr []int) []int {
    // get biggest array to the
    for i,_ := range arr{
        biggest := -1
        for j:= i+1; j < len(arr); j++ {
            if(arr[j] > biggest){
                biggest = arr[j] 
            }
        }
        arr[i] = biggest
    }

    
    arr[len(arr)-1] = -1

    return arr;
}
