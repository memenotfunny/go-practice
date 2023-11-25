package go_practice

import (
	"slices"
	"testing"
)

/**
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.



Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func TestTs(t *testing.T) {
	ts := twoSum([]int{2, 7, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 15, 1, 1, 1}, 11+15)
	t.Log(ts)
	if !slices.Equal([]int{2, 17}, ts) {
		t.Error("failed test")
	}
}

func twoSum(nums []int, target int) []int {
	rc := make(chan []int, 1)
	check(nums[1:], nums[0], target, 0, rc)
	return <-rc
}

func check(nums []int, toCheck int, target int, toCheckPos int, rc chan []int) {
	for iteratedPos, iteratedNum := range nums {
		if toCheck+iteratedNum == target {
			select {
			case rc <- []int{toCheckPos, toCheckPos + iteratedPos + 1}:
			default:
			}
			break
		}

		newSlice := nums[iteratedPos+1:]
		if len(newSlice) > 0 {
			go check(newSlice, iteratedNum, target, toCheckPos+iteratedPos+1, rc)
		}
	}
}
