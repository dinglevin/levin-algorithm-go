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

