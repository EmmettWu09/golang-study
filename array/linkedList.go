package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

type LinkedListNode struct {
	key  string // book tile key
	Book Book
	Next *LinkedListNode
	Prev *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Tail *LinkedListNode
	Size int // 0
}

func (ll *LinkedList) Add(book Book) {
	newNode := &LinkedListNode{Book: book}

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
	} else {
		ll.Tail.Next = newNode
		newNode.Prev = ll.Tail
		ll.Tail = newNode
	}

	ll.Size++
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func (ll *LinkedList) Find(title string) *Book {
	current := ll.Head
	for current != nil {
		if current.Book.Title == title {
			return &current.Book
		}
		current = current.Next
	}
	return nil
}

func main2() {
	bk := Book{
		Title:  "Go Programming",
		Author: "John Doe",
		Price:  29.99,
	}
	bk2 := Book{
		Title:  "Advanced Go",
		Author: "John Doe2",
		Price:  59.99,
	}
	fmt.Printf("Book: %+v\n", bk)
	fmt.Printf("Book: %+v\n", bk2)

	ll := &LinkedList{}
	ll.Add(bk)
	ll.Add(bk2)
}
