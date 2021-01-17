package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second

	l.length++
}

func (l *linkedList) delete(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next

		l.length--

		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}

		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next

	l.length--
}

func (l linkedList) listData() {
	head := l.head

	for l.length != 0 {
		fmt.Printf("%d ", head.data)

		head = head.next

		l.length--
	}

	fmt.Printf("\n") // Line break
}

func main() {
	myLinkedList := linkedList{}

	nodeOne := &node{data: 48}
	nodeTwo := &node{data: 38}
	nodeThree := &node{data: 28}

	myLinkedList.prepend(nodeOne)
	myLinkedList.prepend(nodeTwo)
	myLinkedList.prepend(nodeThree)

	myLinkedList.listData()

	myLinkedList.delete(48)

	myLinkedList.listData()
}
