package twosum

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[n]; ok {
			return []int{j, i}
		}
		m[target-nums[i]] = i
	}
	return nil
}
