package main

import "fmt"

// 原地修改，带返回值，无任何额外数组
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits // 无进位，返回修改后的原数组
		}
		digits[i] = 0 // 有进位，原地置0
	}
	// 全9情况：原地修改+扩容
	digits[0] = 1
	// 扩容一位，末尾补0
	return append(digits, 0)
}

func main() {
	// 测试用例
	testCases := [][]int{
		{1, 2, 3}, // [1 2 4]
		{9, 9, 9}, // [1 0 0 0]
		{9},       // [1 0]
	}
	// 遍历测试用例
	for _, testCase := range testCases {
		res := plusOne(testCase)
		// 打印结果
		fmt.Printf("输入: %v  , 这个大整数加一后输出: ", testCase)
		fmt.Println(res)
		// 另一种写法
		fmt.Printf("输入: %v  ，这个大整数加一后输出: %v\n", testCase, res)
	}
}