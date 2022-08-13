```
func max(a, b int) int {
if a > b {
return a
}
return b
}
​
func maxProfit(prices []int) int {
if len(prices) <= 1 {
return 0
}
left_ptr, right_ptr := 0, 1
max_profit := 0
for right_ptr < len(prices) {
if prices[right_ptr] > prices[left_ptr]{
max_profit = max(max_profit, prices[right_ptr] - prices[left_ptr])
} else {
left_ptr = right_ptr
}
right_ptr +=1
}
return max_profit
}
​
​
```