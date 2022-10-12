func rob(nums []int) int {
    loot1, loot2 := 0, 0
    for _, loot := range nums {
        treasure_chest := max(loot + loot1, loot2)
        loot1 = loot2
        loot2 = treasure_chest
    }
    return loot2
} 

func max(a int, b int) int {
    if a < b {
        return b
    }
    return a
}
