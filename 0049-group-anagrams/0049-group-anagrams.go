func groupAnagrams(strs []string) [][]string {
	mapper := make(map[string][]string)

	for _, val := range strs {
		//slice of unicodes of the str
		word := []byte(val)
		slices.Sort(word)
		sorted := string(word)
		mapper[sorted] = append(mapper[sorted], val)
	}

	var arr [][]string

	for _, vals := range mapper {
		arr = append(arr, vals)
	}

	return arr
}