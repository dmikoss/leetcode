func twoSum(nums []int, target int) []int {
	t := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				t = append(t, i, j)
				return t
			}
		}
	}
	return t
}