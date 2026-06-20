func isValid(s string) bool {
    chars := map[rune]rune{
        ')':'(',
        '}':'{', 
        ']':'[',
    }

    stack := make([]rune,0)
    
    for _,c := range s {
        if len(stack) == 0 {
            stack = append(stack,c);
        } else {
            top := stack[len(stack)-1]

            if top == chars[c] {
                stack = stack[:len(stack)-1]
            }else {
                stack = append(stack,c);
            }
        }
    }

    return len(stack) == 0;
    
}
