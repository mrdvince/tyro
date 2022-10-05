func twoSum(nums []int, target int) []int {
    values := []int{}
    dict := map[int]int{}
    
    for idx, num := range nums {
        missingNum := target - num
        
        if val, ok := dict[missingNum]; ok {
            return []int{idx,val} 
        }
        dict[num] = idx
    }
    
    return values
    
    
}