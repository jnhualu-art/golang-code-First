package main

import "fmt"

// 判断有效的括号
func isValid(s string) bool {
    // 用map存括号的匹配规则：右括号为key，左括号为value
    bracketMap := map[byte]byte{')':'(', ']':'[', '}':'{'}
    // 定义栈，存左括号
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        char := s[i]
        // 如果是右括号
        if val, ok := bracketMap[char]; ok {
            // 栈为空 或 栈顶的左括号不匹配，直接返回false
            if len(stack) == 0 || stack[len(stack)-1] != val {
                return false
            }
            // 匹配成功，弹出栈顶元素
            stack = stack[:len(stack)-1]
        } else {
            // 如果是左括号，入栈
            stack = append(stack, char)
        }
    }
    // 最终栈必须为空，才算完全匹配
    return len(stack) == 0
}

func main() {
    // 测试用例，可随意修改测试任意字符串
	testCases := []string{
		"()", "()[]{}", "(]", "([)]", "{[]}", "", "((()))", "((()){",
	}
	// 遍历测试用例，打印结果
	for _, s := range testCases {
		if isValid(s) {
			fmt.Printf("字符串: %-10s  有效!\n", s)
		}else {
			fmt.Printf("字符串: %-10s  无效!\n", s)
		}
	}
}