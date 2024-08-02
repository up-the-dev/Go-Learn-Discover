// https://leetcode.com/problems/two-sum/

// here we used map to store elements in slice and its index
package main

func twoSum(nums []int, target int) []int {
	elementIndexMap := make(map[int]int)
	for indexOfVal, val := range nums {

		val2 := target - val
		value, ok := elementIndexMap[val2]
		if ok == true {
			return []int{indexOfVal, value}
			break
		}
		elementIndexMap[val] = indexOfVal
	}

	return []int{-1, -1}
}
