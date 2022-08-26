type NumArray struct {
    Nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{Nums: nums}
}


func (this *NumArray) SumRange(left int, right int) int {
    result := 0
    for _, val := range this.Nums[left:right+1] {
        result += val
    }
    return result
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
