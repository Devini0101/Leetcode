func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	mapper := make(map[string][]string)

	for _, val := range strs {
		//slice of unicodes of the str
		word := []rune(val)

		sort.Slice(word, func(i, j int) bool {
			return word[i] < word[j]
		})
		sorted := string(word)
		mapper[sorted] = append(mapper[sorted], val)
	}

	var arr [][]string

	for _, vals := range mapper {
		arr = append(arr, vals)
	}

	return arr
}