class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        const anagramMap = {};

        for (let word of strs) {
            let sortedWord = word.split("").sort().join("");
            
            if (!anagramMap[sortedWord]) {
                anagramMap[sortedWord] = [];
            }
            
            anagramMap[sortedWord].push(word);
        }

        return Object.values(anagramMap);
    }
}
