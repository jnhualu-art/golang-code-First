package main

import (
    "fmt"
    "strings" 
)

// 最长公共前缀
func longestCommonPrefix(str []string) string {
	// 特殊情况处理
    if len(str) == 0 {
        return ""
    }
	// 取第一个字符串作为初始前缀
    prefix := str[0]
	// 遍历剩余字符串
    for i := 1; i < len(str); i++ {
		// 循环条件：只要【当前字符串】不以【基准前缀】开头，就执行循环体
        for !strings.HasPrefix(str[i], prefix) {
			// 缩短基准前缀：把前缀的【最后一个字符切掉】
            prefix = prefix[:len(prefix)-1]
			// 临界值：如果前缀被切成空字符串了，说明没有公共前缀，直接返回空
            if prefix == "" {
                return ""
            }
        }
    }
    return prefix
}

func main() {
    // 测试用例
	testCases := [][]string{
		{"flower", "flow", "flight"},	// 输出: "fl"
		{"dog", "racecar", "car"},		// 输出: ""
		{"interspecies", "interstellar", "interstate"}, // 输出: "inters"
		{"throne", "throne"},			// 输出: "throne"
		{"throne", "dungeon"},			// 输出: ""
		{},								// 输出: ""
		{"a"},							// 输出: "a"
	}
	// 遍历测试用例，打印结果
	for _, strs := range testCases {
		result := longestCommonPrefix(strs)
		fmt.Printf("输入: %-20v  最长公共前缀: \"%s\"\n", strs, result)
	}						
}