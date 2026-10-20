func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    mapper := make(map[rune]int)

    for _, letter := range s {
        mapper[letter]++
    }

    for _, letter := range t {
        mapper[letter]--
        if mapper[letter] < 0 {
            return false
        }
    }
    return true
}