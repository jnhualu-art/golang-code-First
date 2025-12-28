package main

import "fmt"

// removeDuplicates 原地删除有序数组的重复项
func removeDuplicates(nums []int) int {
	// 边界处理：空数组 或 只有1个元素，直接返回原长度
	if len(nums) <= 1 {
		return len(nums)
	}
	// 慢指针：指向去重后数组的「最后一个有效元素」，初始为0
	slow := 0
	// 快指针：遍历数组，寻找和有效元素不同的新元素，初始为1
	for fast := 1; fast < len(nums); fast++ {
		// 找到不重复的元素
		if nums[fast] != nums[slow] {
			slow++ // 慢指针后移一位，指向新的填充位置
			nums[slow] = nums[fast] // 原地覆盖，写入不重复元素
		}
		// 如果相等，就是重复元素，快指针直接跳过即可
	}
	// 返回有效长度：慢指针是下标，从0开始，所以+1
	return slow + 1
}

func main() {
	// 包含所有边界场景，无遗漏
	testCases := [][]int{
		{},                  // 空数组
		{1},                 // 单个元素
		{1, 1, 2},           // 题目指定用例 ✅
		{0, 0, 1, 1, 1, 2},  // 连续重复项
		{1, 2, 3, 4, 5},     // 无任何重复
		{2, 2, 2, 2},        // 全是重复元素
	}

	// 遍历测试+友好打印，格式清晰
	for _, nums := range testCases {
		length := removeDuplicates(nums)
		fmt.Printf("原数组: %v  去重后长度: %2d  有效元素: %v\n", nums, length, nums[:length])
	}
}