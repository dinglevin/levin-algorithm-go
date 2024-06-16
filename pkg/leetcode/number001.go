/**
 * 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
 * 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
 * 你可以按任意顺序返回答案。
 *
 * 示例 1：
 * 输入：nums = [2,7,11,15], target = 9
 * 输出：[0,1]
 * 解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
 *
 * 示例 2：
 * 输入：nums = [3,2,4], target = 6
 * 输出：[1,2]
 *
 * 示例 3：
 * 输入：nums = [3,3], target = 6
 * 输出：[0,1]
 *
 * 提示：
 * 2 <= nums.length <= 10^4
 * -10^9 <= nums[i] <= 10^9
 * -10^9 <= target <= 10^9
 * 只会存在一个有效答案
 *
 * 进阶：你可以想出一个时间复杂度小于 O(n2) 的算法吗？
 */
package leetcode

import "sort"

func twoSumV1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int {i, j}
			}
		}
	}

	return []int {0, 0}
}

func twoSumV2(nums []int, target int) []int {
	m1 := make(map[int]int)
	for i, v := range nums {
		_, ok := m1[target - v]
		if ok {
			return []int {i, m1[target - v]}
		} else {
			m1[v] = i
		}
	}

	return []int{0, 0}
}

func twoSumV3(nums []int, target int) []int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	left := 0
	right := len(nums) - 1
	for left < right {
		sum := sortedNums[left] + sortedNums[right]
		if sum == target {
			return []int{
				findIndex(nums, sortedNums[left], false),
				findIndex(nums, sortedNums[right], true),
			}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{0, 0}
}

func findIndex(nums []int, value int, reverse bool) int {
	if reverse {
		for i := len(nums) - 1; i >= 0; i-- {
			if nums[i] == value {
				return i
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			if nums[i] == value {
				return i
			}
		}
	}
	return -1
}

