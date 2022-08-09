func containsDuplicate(nums []int) bool {
    set := make(map[int]bool)
    for _, num := range nums {
        if in := set[num]; in {
            return true
        } else {
            set[num] = true
        }
    }
    return false
}