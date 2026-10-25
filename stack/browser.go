package main

import (
	"fmt"
)

// Stack 用切片实现字符串栈，切片末尾是栈顶。
// 栈遵循“后进先出”（LIFO）：最后入栈的页面最先出栈。
// 零值即可使用，不需要提前分配空间。
type Stack struct {
	items []string
}

func (s *Stack) Push(page string) {
	s.items = append(s.items, page)
}

// Pop 弹出栈顶元素。空栈返回 false，避免切片下标越界。
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	last := len(s.items) - 1
	tmp := s.items[last]
	s.items[last] = "" // 清除底层数组对已弹出字符串的引用。
	s.items = s.items[:last]
	return tmp, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Clear() {
	s.items = nil
}

// Browser 使用两个栈模拟单个浏览器标签页的历史记录。
// 请用 NewBrowser 指定初始页面。
//
// 当前页面单独保存，不在任何一个栈中：
//   - backStack 的栈顶是下一次“后退”要访问的页面。
//   - forwardStack 的栈顶是下一次“前进”要访问的页面。
//
// 例如依次访问 首页、A、B 后：
// 后退栈：[首页 A]，当前页面：B，前进栈：[]。
// 后退一次后：
// 后退栈：[首页]，当前页面：A，前进栈：[B]。
type Browser struct {
	current      string
	backStack    Stack
	forwardStack Stack
}

func NewBrowser(homepage string) *Browser {
	return &Browser{current: homepage}
}

// Visit 访问一个新页面：当前页面入后退栈，再更新当前页面。
// 必须清空前进栈：从旧页面打开新页面后，原来的前进路径就失效了。
func (b *Browser) Visit(page string) {
	b.backStack.Push(b.current)
	b.current = page
	b.forwardStack.Clear()
}

// Back 后退一步。没有可后退的页面时返回 false，浏览器状态不变。
func (b *Browser) Back() bool {
	page, ok := b.backStack.Pop()
	if !ok {
		return false
	}
	b.forwardStack.Push(b.current)
	b.current = page
	return true
}

// Forward 前进一步，与 Back 的操作对称。
// 注意：前进不能调用 Visit，否则会错误地清空剩余的前进历史。
func (b *Browser) Forward() bool {
	page, ok := b.forwardStack.Pop()
	if !ok {
		return false
	}
	b.backStack.Push(b.current)
	b.current = page
	return true
}

func (b *Browser) Current() string {
	return b.current
}

// PrintState 仅用于课堂演示：两个栈均从左到右按“栈底 -> 栈顶”显示。
func (b *Browser) PrintState(action string) {
	fmt.Printf("%s\n  后退栈：%v | 当前：%s | 前进栈：%v\n\n",
		action, b.backStack.items, b.Current(), b.forwardStack.items)
}

// 运行：go run ./stack
// 使用切片作为栈时，Push 均摊 O(1)，Pop 为 O(1)。
// 因此访问、前进、后退均摊 O(1)，历史记录共占 O(n) 空间。
func main() {

	// n := runtime.Callers(1, pcs)
	// frames := runtime.CallersFrames(pcs[:n])

	browser := NewBrowser("首页")
	browser.PrintState("打开浏览器（栈的最右侧是栈顶）")

	browser.Visit("页面A")
	browser.PrintState("访问页面A")
	browser.Visit("页面B")
	browser.PrintState("访问页面B")

	ok := browser.Back()
	browser.PrintState(fmt.Sprintf("后退，成功：%t", ok)) // 当前：页面A
	ok = browser.Forward()
	browser.PrintState(fmt.Sprintf("前进，成功：%t", ok)) // 当前：页面B

	// 连续后退：前进栈将依次压入页面B、页面A，栈顶是页面A。
	ok = browser.Back()
	browser.PrintState(fmt.Sprintf("后退，成功：%t", ok)) // 当前：页面A
	ok = browser.Back()
	browser.PrintState(fmt.Sprintf("后退，成功：%t", ok)) // 当前：首页
	ok = browser.Back()
	browser.PrintState(fmt.Sprintf("已到首页，再后退，成功：%t", ok)) // false，状态不变

	// 前进栈原本是 [页面B 页面A]，访问页面C后整个前进栈被清空。
	browser.Visit("页面C")
	browser.PrintState("后退后访问页面C：清空前进历史")
	ok = browser.Forward()
	browser.PrintState(fmt.Sprintf("尝试前进，成功：%t", ok)) // false，状态不变
}
