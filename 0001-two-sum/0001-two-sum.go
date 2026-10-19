func twoSum(nums []int, target int) []int {
    numbers := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if index, found := numbers[complement]; found {
            return[]int{index,i}
        }
        numbers[num] = i
    }
    return nil
}