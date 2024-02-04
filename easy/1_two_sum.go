func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    var remaining int
    for i := 0; i < len(nums); i++ {
        remaining = target - nums[i]

        if val, ok := seen[remaining]; ok {
            return []int{i, val}
        }

        seen[nums[i]] = i
    }
    return []int{}
}
