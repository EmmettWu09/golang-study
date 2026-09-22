package main

import "fmt"

func main() {
	// 创建 int 切片：长度为 3，容量为 5，元素默认值为 0
	s := make([]int, 3, 5)

	fmt.Printf("切片变量的地址：%p\n", &s)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40 // 这里会报错，索引越界，因为切片的长度是 3，不能直接访问索引为 3 的元素
	s = append(s, 40)

	fmt.Println(s)      // [10 20 30 40]
	fmt.Println(len(s)) // 4，当前元素数量
	fmt.Println(cap(s)) // 5，底层数组的容量
	fmt.Println("------------")
	fmt.Printf("切片变量的地址：%p\n", &s)
	s = append(s, 40)
	s = append(s, 40)
	s = append(s, 40)
	s = append(s, 40)
	s = append(s, 40)
	fmt.Printf("切片变量的地址：%p\n", &s)
}
