func twoSum(numbers []int, target int) []int {
    m := map[int]int{}

    for i,_ := range numbers {
        diff := target - numbers[i];

        if _, ok := m[diff]; ok {
            return []int{m[diff]+1,i+1}
        }

        m[numbers[i]] = i;  
    }
    return []int{}
}
