type NumArray struct {
    Nums []int
}


func Constructor(nums []int) NumArray {
    cumulative := make([]int, len(nums)) 
    // accumulate
    // [-2,0,3,-5,2,-1] -> [-2 -2 1 -4 -2 -3]

    sum := 0
    for idx := range nums {
        sum += nums[idx]
        cumulative[idx] = sum
    }    
    return NumArray{Nums:cumulative}
}


func (this *NumArray) SumRange(left int, right int) int {
    // out of range guard
    if left == 0 {
        return this.Nums[right]
    }
    return this.Nums[right] - this.Nums[left-1] 
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
