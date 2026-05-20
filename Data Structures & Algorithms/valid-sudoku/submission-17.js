class Solution {
    /**
     * @param {character[][]} board
     * @return {boolean}
     */
    isValidSudoku(board) {
        let curr_box = 0
        let arr = Array.from({length: 9}, () => new Set());

        for(let i=0; i < board.length * board[0].length ; i++){
            const box_row_start = Math.floor(curr_box / 3) * 3;  // 0, 3, or 6
            const box_col_start = (curr_box % 3) * 3;            // 0, 3, or 6
            const cell = i % 9;                                  // 0..8 within this box
            const cur_row = box_row_start + Math.floor(cell / 3);
            const cur_col = box_col_start + (cell % 3);

            const val = board[cur_row][cur_col];

            if(arr[curr_box].has(val)){
                return false;
            }

            if(val != "."){
                arr[curr_box].add(val);
            }
            
            if(i%9 == 8)
                curr_box += 1
        }


        for(let col=0; col< board.length; col++){
            let seen_col = []
            let seen_row = []

            for(let row=0;row < board[0].length; row++) {
                if(seen_col.includes(board[row][col])){
                    return false;
                }

                if(board[row][col] != "."){
                    seen_col = [...seen_col,board[row][col]]
                }

                if(seen_row.includes(board[col][row])){
                    return false;
                }

                if(board[col][row] != "."){
                    seen_row = [...seen_row,board[col][row]]
                }
            } 
        }
        return true;
    }
}
