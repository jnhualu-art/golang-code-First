package main

import "fmt" // 导入输入输出包

// 判断整数是否为回文数
func palindrome(num int) bool {
	// 边界判断：负数 或 末尾为0的非0数，一定不是回文数
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}
	// 修复：把reverseNum声明在for外面，作用域覆盖整个函数
	reverseNum := 0
	// 修复：Go的for循环 去掉圆括号()，这是核心语法修正
	for ; num > reverseNum; {
		reverseNum = reverseNum*10 + num%10
		num /= 10
	}
	// 偶数位相等 或 奇数位去掉中间数相等，返回true
	return num == reverseNum || num == reverseNum/10
}

func main() {
	var num int // 声明存储输入的变量
	fmt.Println("请输入一个整数，程序将判断是否为回文数：")
	fmt.Scan(&num) // 接收控制台输入，&不能省略
	fmt.Println("输入的数字是：", num)
	fmt.Println("是否为回文数：", palindrome(num))
}