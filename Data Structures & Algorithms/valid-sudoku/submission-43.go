func isValidSudoku(board [][]byte) bool {
    cur_box := 0
    
    // 1. Fixed: Change flat map[int]int to an array of 9 maps, tracking byte frequencies
    var rows [9]map[byte]int
    var cols [9]map[byte]int
    var arr  [9]map[byte]int

    // Initialize each map in the arrays
    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]int)
        cols[i] = make(map[byte]int)
        arr[i]  = make(map[byte]int)
    }

    for i := 0; i < len(board) * len(board[0]); i++ {
        box_row_start := int(math.Floor(float64(cur_box)/3) * 3)
        box_col_start := (cur_box % 3) * 3
        cell := i % 9

        cur_row := box_row_start + int(math.Floor(float64(cell)/3))
        cur_col := box_col_start + (cell % 3)

        val := board[cur_row][cur_col]

        // 2. Fixed: Compare byte to byte literal '.' using single quotes instead of double quotes
        if val == '.' {
            if i % 9 == 8 { cur_box += 1 }
            continue
        }

        // 3. Fixed: Check if map integer values are greater than 0 instead of using boolean checks
        if rows[cur_row][val] > 0 || cols[cur_col][val] > 0 || arr[cur_box][val] > 0 {
            return false
        }

        rows[cur_row][val]++
        cols[cur_col][val]++
        arr[cur_box][val]++

        if i % 9 == 8 { cur_box += 1 }
    }
    return true
}