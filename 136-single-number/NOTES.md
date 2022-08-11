```
func singleNumber(nums []int) int {
dict := make(map[int]int)
for _, num := range nums {
if val,ok := dict[num]; ok {
dict[num] = val+1
} else {
dict[num] = 1
}
}
var res int
for k, v := range dict{
if v == 1 {
res = k
}
}
return res
}
```