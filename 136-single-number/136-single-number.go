func singleNumber(nums []int) int {
    xor := 0
    for _, num := range nums {
        xor = xor^num
    }
    return xor
}