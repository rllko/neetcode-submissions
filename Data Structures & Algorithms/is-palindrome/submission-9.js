class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        s = s.replace(/[ .!@#$%^&*()_+\-=?,:\\''\\/]/gi, ""); 
        for (let i = 0, j = s.length-1; i <= j; j--, i++) {


            if (s[i].toLocaleLowerCase() != s[j].toLocaleLowerCase()) {
                return false;
            }

            
        }
        return true;
    }
}
