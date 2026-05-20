
func groupAnagrams(strs []string) [][]string {
    anagramMap := make(map[string][]string);

    for _,word := range strs {

        wordRune := []rune(word);
        sort.Slice(wordRune, func(i,j int) bool {
            return wordRune[i] < wordRune[j];
        })
        sortedWord := string(wordRune);
    
        anagramMap[sortedWord] = append(anagramMap[sortedWord],word);
    }

    result := make([][]string,0, len(anagramMap))
    for _,group := range anagramMap {
        result = append(result,group);
    }

    return result
}
