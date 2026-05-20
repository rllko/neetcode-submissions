class Solution {
    /**
     * @param {string[]} strs
     * @return {string[][]}
     */
    groupAnagrams(strs) {
        const anagramMap = {}; // This will hold our groups

        for (let word of strs) {
            // 1. Sort the word to use as a key
            let sortedWord = word.split("").sort().join("");
            
            // 2. If this sorted word isn't in our map yet, create an empty array
            if (!anagramMap[sortedWord]) {
                anagramMap[sortedWord] = [];
            }
            
            // 3. Push the original word into the correct group
            anagramMap[sortedWord].push(word);
        }

        // 4. Return just the grouped arrays using Object.values()
        return Object.values(anagramMap);
    }
}
