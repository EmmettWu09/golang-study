package main

import (
	"fmt"
	"runtime/debug"
)

// 运行：go run ./stack/callstack
// 调用顺序是 main -> a -> b，返回顺序是 b -> a -> main。
func main() {
	fmt.Println("1. main 开始")
	a()
	fmt.Println("5. 回到 main，程序结束")
}

func a() {
	fmt.Println("2. 进入 a")
	b()
	fmt.Println("4. 回到 a")
}

func b() {
	fmt.Println("3. 进入 b，打印当前调用栈：")

	// debug.Stack() 返回当前 goroutine 的调用栈，包含函数名、文件和行号。
	// 用 %s 将返回的字节切片作为文字打印。
	// 跳过打印工具自身的 runtime/debug.Stack 后，依次看到 b、a、main。
	fmt.Printf("%s\n", debug.Stack())

	// 打印调用栈不会中断程序，b 结束后仍会返回 a。
}
