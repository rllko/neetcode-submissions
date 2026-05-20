
func isAnagram(s string, t string) bool {
    sStr := strings.Split(s,"");
    tStr := strings.Split(t,"");

    sort.Strings(sStr)
    sort.Strings(tStr)

    sSorted:= strings.Join(sStr,"")
    tSorted := strings.Join(tStr,"");

    return sSorted == tSorted
}
