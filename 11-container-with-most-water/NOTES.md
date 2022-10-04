func maxArea(height []int) int {
maxArea := 0
for i := 0; i < len(height); i++ {
for j := 1; j < len(height); j++ {
min(height[i], height[j])
m := min(height[i], height[j])
maxArea = max(maxArea, m*(j-i))
}
}
return maxArea
}
​
func min(a, b int) int {
if a < b {
return a
}
return b
}
​
func max(a, b int) int {
if a > b {
return a
}
return b
}