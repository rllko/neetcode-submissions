class Solution {
    /**
     * @param {character[][]} board
     * @return {boolean}
     */
    isValidSudoku(board) {
        let curr_box = 0;
        const rows  = Array.from({length: 9}, () => new Set());
        const cols  = Array.from({length: 9}, () => new Set());
        let arr = Array.from({ length: 9 }, () => new Set());

        for (let i = 0; i < board.length * board[0].length; i++) {
            const box_row_start = Math.floor(curr_box / 3) * 3; 
            const box_col_start = (curr_box % 3) * 3; 
            const cell = i % 9; 

            const cur_row = box_row_start + Math.floor(cell / 3);
            const cur_col = box_col_start + (cell % 3);

            const val = board[cur_row][cur_col];

            if (val === ".") {
                if (i % 9 == 8) curr_box += 1;
                continue;
            }

            if (rows[cur_row].has(val) || cols[cur_col].has(val) || arr[curr_box].has(val)) {
                return false;
            }

            rows[cur_row].add(val);
            cols[cur_col].add(val);
            arr[curr_box].add(val);
            

            if (i % 9 == 8) curr_box += 1;
        }

        return true;
    }
}
