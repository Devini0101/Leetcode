func containsDuplicate(nums []int) bool {
    maper := make(map[int]bool)
    for _, value  := range nums {
        _, ok := maper[value]
        //already exists
        if ok {
            return true
        }
        maper[value] = true
    }
    return false
}