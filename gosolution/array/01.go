package array

/**
  解题思路:
   1. 用一个map保存我们需要的数 key 是我们需要的数，value 是该数在nums中的下标
   2. 遍历nums 数组，如果map中存在该数，则返回该数的下标，以及当前的i

https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0001.Two-Sum/
*/

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	length := len(nums)
	for i := 0; i < length; i++ {
		another := target - nums[i]
		if index, ok := m[another]; ok {
			return []int{index, i}
		}
		m[nums[i]] = i
	}

	return nil

}

func s2(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	result := []int{0, 0}
	m := make(map[int]int)

	for left <= right {
		if v, ok := m[target-nums[left]]; ok {
			result[0] = left
			result[1] = v
			break
		}

		m[nums[left]] = left
		if v, ok := m[target-nums[right]]; ok {
			result[0] = right
			result[1] = v
			break
		}
		m[nums[right]] = right
		right--
		left++
	}
	return result
}

// 26 - 2 = 24
// m[24] = nil
// m[2] = 0

// 26 - 7 = 19
// m[19] = nil
// m[7] = 1

// 26 - 11 = 15
// m[15] = nil
// m[11] = 2

// 26 - 15 = 11
// m[11] = 2
