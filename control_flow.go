package main

import "fmt" // 导入输出包，打印内容用

func main() {
	// 定义目标数组
	nums := [21]int{3, 7, 9, 2, 5, 8, 11, 4, 6, 17, 6, 4, 11, 8, 5, 2, 9, 7, 3, 10, 10}

	// 数组转切片
	snums := nums[:]

	// 1. 定义map统计每个元素出现的次数
	countMap := make(map[int]int)

	// 2. 遍历切片统计次数
	for _, v := range snums {
		countMap[v]++
	}

	// ===== 第一步：只打印【所有元素的出现次数】，纯统计，无多余内容 =====
	fmt.Println("===== 每个元素的出现个数 =====")
	for k, v := range countMap {
		fmt.Printf("元素 %d 出现了 %d 次\n", k, v)
	}

	// ===== 第二步：单独查找+输出【唯一只出现1次的元素】，单独排版，醒目 =====
	fmt.Println("\n===== 唯一只出现一次的元素 ======") // \n 是换行符，隔开两段内容更美观
	for k, v := range countMap {
		if v == 1 {
			fmt.Printf("✅ 唯一元素：%d\n", k)
		}
	}
}